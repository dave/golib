package x86

import (
	"debug/elf"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"log"
)

// Append 4 bytes to s and create a R_CALL relocation targeting t to fill them in.
func addcall(ctxt *ld.Link, s *sym.Symbol, t *sym.Symbol) {
	s.Attr |= sym.AttrReachable
	i := s.Size
	s.Size += 4
	s.Grow(s.Size)
	r := s.AddRel()
	r.Sym = t
	r.Off = int32(i)
	r.Type = objabi.R_CALL
	r.Siz = 4
}

func gentext(ctxt *ld.Link) {
	if ctxt.DynlinkingGo() {

	} else {
		switch ctxt.BuildMode {
		case ld.BuildModeCArchive:
			if !ctxt.IsELF {
				return
			}
		case ld.BuildModePIE, ld.BuildModeCShared, ld.BuildModePlugin:

		default:
			return
		}
	}

	thunks := make([]*sym.Symbol, 0, 7+len(ctxt.Textp))
	for _, r := range [...]struct {
		name string
		num  uint8
	}{
		{"ax", 0},
		{"cx", 1},
		{"dx", 2},
		{"bx", 3},

		{"bp", 5},
		{"si", 6},
		{"di", 7},
	} {
		thunkfunc := ctxt.Syms.Lookup("__x86.get_pc_thunk."+r.name, 0)
		thunkfunc.Type = sym.STEXT
		thunkfunc.Attr |= sym.AttrLocal
		thunkfunc.Attr |= sym.AttrReachable
		o := func(op ...uint8) {
			for _, op1 := range op {
				thunkfunc.AddUint8(op1)
			}
		}

		o(0x8b, 0x04+r.num<<3, 0x24)

		o(0xc3)

		thunks = append(thunks, thunkfunc)
	}
	ctxt.Textp = append(thunks, ctxt.Textp...)

	addmoduledata := ctxt.Syms.Lookup("runtime.addmoduledata", 0)
	if addmoduledata.Type == sym.STEXT && ctxt.BuildMode != ld.BuildModePlugin {

		return
	}

	addmoduledata.Attr |= sym.AttrReachable

	initfunc := ctxt.Syms.Lookup("go.link.addmoduledata", 0)
	initfunc.Type = sym.STEXT
	initfunc.Attr |= sym.AttrLocal
	initfunc.Attr |= sym.AttrReachable
	o := func(op ...uint8) {
		for _, op1 := range op {
			initfunc.AddUint8(op1)
		}
	}

	o(0x53)

	o(0xe8)
	addcall(ctxt, initfunc, ctxt.Syms.Lookup("__x86.get_pc_thunk.cx", 0))

	o(0x8d, 0x81)
	initfunc.AddPCRelPlus(ctxt.Arch, ctxt.Moduledata, 6)

	o(0x8d, 0x99)
	i := initfunc.Size
	initfunc.Size += 4
	initfunc.Grow(initfunc.Size)
	r := initfunc.AddRel()
	r.Sym = ctxt.Syms.Lookup("_GLOBAL_OFFSET_TABLE_", 0)
	r.Off = int32(i)
	r.Type = objabi.R_PCREL
	r.Add = 12
	r.Siz = 4

	o(0xe8)
	addcall(ctxt, initfunc, addmoduledata)

	o(0x5b)

	o(0xc3)

	if ctxt.BuildMode == ld.BuildModePlugin {
		ctxt.Textp = append(ctxt.Textp, addmoduledata)
	}
	ctxt.Textp = append(ctxt.Textp, initfunc)
	initarray_entry := ctxt.Syms.Lookup("go.link.addmoduledatainit", 0)
	initarray_entry.Attr |= sym.AttrReachable
	initarray_entry.Attr |= sym.AttrLocal
	initarray_entry.Type = sym.SINITARR
	initarray_entry.AddAddr(ctxt.Arch, initfunc)
}

func (psess *PackageSession) adddynrel(ctxt *ld.Link, s *sym.Symbol, r *sym.Reloc) bool {
	targ := r.Sym

	switch r.Type {
	default:
		if r.Type >= 256 {
			psess.ld.
				Errorf(s, "unexpected relocation type %d (%s)", r.Type, psess.sym.RelocName(ctxt.Arch, r.Type))
			return false
		}

	case 256 + objabi.RelocType(elf.R_386_PC32):
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_386_PC32 relocation for dynamic symbol %s", targ.Name)
		}

		if (targ.Type == 0 || targ.Type == sym.SXREF) && !targ.Attr.VisibilityHidden() {
			psess.ld.
				Errorf(s, "unknown symbol %s in pcrel", targ.Name)
		}
		r.Type = objabi.R_PCREL
		r.Add += 4
		return true

	case 256 + objabi.RelocType(elf.R_386_PLT32):
		r.Type = objabi.R_PCREL
		r.Add += 4
		if targ.Type == sym.SDYNIMPORT {
			psess.
				addpltsym(ctxt, targ)
			r.Sym = ctxt.Syms.Lookup(".plt", 0)
			r.Add += int64(targ.Plt)
		}

		return true

	case 256 + objabi.RelocType(elf.R_386_GOT32), 256 + objabi.RelocType(elf.R_386_GOT32X):
		if targ.Type != sym.SDYNIMPORT {

			if r.Off >= 2 && s.P[r.Off-2] == 0x8b {

				s.P[r.Off-2] = 0x8d

				r.Type = objabi.R_GOTOFF
				return true
			}

			if r.Off >= 2 && s.P[r.Off-2] == 0xff && s.P[r.Off-1] == 0xb3 {

				s.P[r.Off-2] = 0x36

				s.P[r.Off-1] = 0x68
				r.Type = objabi.R_ADDR
				return true
			}
			psess.ld.
				Errorf(s, "unexpected GOT reloc for non-dynamic symbol %s", targ.Name)
			return false
		}
		psess.
			addgotsym(ctxt, targ)
		r.Type = objabi.R_CONST
		r.Sym = nil
		r.Add += int64(targ.Got)
		return true

	case 256 + objabi.RelocType(elf.R_386_GOTOFF):
		r.Type = objabi.R_GOTOFF
		return true

	case 256 + objabi.RelocType(elf.R_386_GOTPC):
		r.Type = objabi.R_PCREL
		r.Sym = ctxt.Syms.Lookup(".got", 0)
		r.Add += 4
		return true

	case 256 + objabi.RelocType(elf.R_386_32):
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_386_32 relocation for dynamic symbol %s", targ.Name)
		}
		r.Type = objabi.R_ADDR
		return true

	case 512 + ld.MACHO_GENERIC_RELOC_VANILLA*2 + 0:
		r.Type = objabi.R_ADDR
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected reloc for dynamic symbol %s", targ.Name)
		}
		return true

	case 512 + ld.MACHO_GENERIC_RELOC_VANILLA*2 + 1:
		if targ.Type == sym.SDYNIMPORT {
			psess.
				addpltsym(ctxt, targ)
			r.Sym = ctxt.Syms.Lookup(".plt", 0)
			r.Add = int64(targ.Plt)
			r.Type = objabi.R_PCREL
			return true
		}

		r.Type = objabi.R_PCREL
		return true

	case 512 + ld.MACHO_FAKE_GOTPCREL:
		if targ.Type != sym.SDYNIMPORT {

			if r.Off < 2 || s.P[r.Off-2] != 0x8b {
				psess.ld.
					Errorf(s, "unexpected GOT reloc for non-dynamic symbol %s", targ.Name)
				return false
			}

			s.P[r.Off-2] = 0x8d
			r.Type = objabi.R_PCREL
			return true
		}
		psess.
			addgotsym(ctxt, targ)
		r.Sym = ctxt.Syms.Lookup(".got", 0)
		r.Add += int64(targ.Got)
		r.Type = objabi.R_PCREL
		return true
	}

	if targ.Type != sym.SDYNIMPORT {
		return true
	}
	switch r.Type {
	case objabi.R_CALL,
		objabi.R_PCREL:
		if ctxt.LinkMode == ld.LinkExternal {

			return true
		}
		psess.
			addpltsym(ctxt, targ)
		r.Sym = ctxt.Syms.Lookup(".plt", 0)
		r.Add = int64(targ.Plt)
		return true

	case objabi.R_ADDR:
		if s.Type != sym.SDATA {
			break
		}
		if ctxt.IsELF {
			psess.ld.
				Adddynsym(ctxt, targ)
			rel := ctxt.Syms.Lookup(".rel", 0)
			rel.AddAddrPlus(ctxt.Arch, s, int64(r.Off))
			rel.AddUint32(ctxt.Arch, ld.ELF32_R_INFO(uint32(targ.Dynid), uint32(elf.R_386_32)))
			r.Type = objabi.R_CONST
			r.Sym = nil
			return true
		}

		if ctxt.HeadType == objabi.Hdarwin && s.Size == int64(ctxt.Arch.PtrSize) && r.Off == 0 {
			psess.ld.
				Adddynsym(ctxt, targ)

			got := ctxt.Syms.Lookup(".got", 0)
			s.Type = got.Type
			s.Attr |= sym.AttrSubSymbol
			s.Outer = got
			s.Sub = got.Sub
			got.Sub = s
			s.Value = got.Size
			got.AddUint32(ctxt.Arch, 0)
			ctxt.Syms.Lookup(".linkedit.got", 0).AddUint32(ctxt.Arch, uint32(targ.Dynid))
			r.Type = 256
			return true
		}
	}

	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	ctxt.Out.Write32(uint32(sectoff))

	elfsym := r.Xsym.ElfsymForReloc()
	switch r.Type {
	default:
		return false
	case objabi.R_ADDR:
		if r.Siz == 4 {
			ctxt.Out.Write32(uint32(elf.R_386_32) | uint32(elfsym)<<8)
		} else {
			return false
		}
	case objabi.R_GOTPCREL:
		if r.Siz == 4 {
			ctxt.Out.Write32(uint32(elf.R_386_GOTPC))
			if r.Xsym.Name != "_GLOBAL_OFFSET_TABLE_" {
				ctxt.Out.Write32(uint32(sectoff))
				ctxt.Out.Write32(uint32(elf.R_386_GOT32) | uint32(elfsym)<<8)
			}
		} else {
			return false
		}
	case objabi.R_CALL:
		if r.Siz == 4 {
			if r.Xsym.Type == sym.SDYNIMPORT {
				ctxt.Out.Write32(uint32(elf.R_386_PLT32) | uint32(elfsym)<<8)
			} else {
				ctxt.Out.Write32(uint32(elf.R_386_PC32) | uint32(elfsym)<<8)
			}
		} else {
			return false
		}
	case objabi.R_PCREL:
		if r.Siz == 4 {
			ctxt.Out.Write32(uint32(elf.R_386_PC32) | uint32(elfsym)<<8)
		} else {
			return false
		}
	case objabi.R_TLS_LE:
		if r.Siz == 4 {
			ctxt.Out.Write32(uint32(elf.R_386_TLS_LE) | uint32(elfsym)<<8)
		} else {
			return false
		}
	case objabi.R_TLS_IE:
		if r.Siz == 4 {
			ctxt.Out.Write32(uint32(elf.R_386_GOTPC))
			ctxt.Out.Write32(uint32(sectoff))
			ctxt.Out.Write32(uint32(elf.R_386_TLS_GOTIE) | uint32(elfsym)<<8)
		} else {
			return false
		}
	}

	return true
}

func (psess *PackageSession) machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	var v uint32

	rs := r.Xsym

	if rs.Type == sym.SHOSTOBJ || r.Type == objabi.R_CALL {
		if rs.Dynid < 0 {
			psess.ld.
				Errorf(s, "reloc %d (%s) to non-macho symbol %s type=%d (%s)", r.Type, psess.sym.RelocName(arch, r.Type), rs.Name, rs.Type, rs.Type)
			return false
		}

		v = uint32(rs.Dynid)
		v |= 1 << 27
	} else {
		v = uint32(rs.Sect.Extnum)
		if v == 0 {
			psess.ld.
				Errorf(s, "reloc %d (%s) to symbol %s in non-macho section %s type=%d (%s)", r.Type, psess.sym.RelocName(arch, r.Type), rs.Name, rs.Sect.Name, rs.Type, rs.Type)
			return false
		}
	}

	switch r.Type {
	default:
		return false
	case objabi.R_ADDR:
		v |= ld.MACHO_GENERIC_RELOC_VANILLA << 28
	case objabi.R_CALL,
		objabi.R_PCREL:
		v |= 1 << 24
		v |= ld.MACHO_GENERIC_RELOC_VANILLA << 28
	}

	switch r.Siz {
	default:
		return false
	case 1:
		v |= 0 << 25
	case 2:
		v |= 1 << 25
	case 4:
		v |= 2 << 25
	case 8:
		v |= 3 << 25
	}

	out.Write32(uint32(sectoff))
	out.Write32(v)
	return true
}

func (psess *PackageSession) pereloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	var v uint32

	rs := r.Xsym

	if rs.Dynid < 0 {
		psess.ld.
			Errorf(s, "reloc %d (%s) to non-coff symbol %s type=%d (%s)", r.Type, psess.sym.RelocName(arch, r.Type), rs.Name, rs.Type, rs.Type)
		return false
	}

	out.Write32(uint32(sectoff))
	out.Write32(uint32(rs.Dynid))

	switch r.Type {
	default:
		return false

	case objabi.R_DWARFSECREF:
		v = ld.IMAGE_REL_I386_SECREL

	case objabi.R_ADDR:
		v = ld.IMAGE_REL_I386_DIR32

	case objabi.R_CALL,
		objabi.R_PCREL:
		v = ld.IMAGE_REL_I386_REL32
	}

	out.Write16(uint16(v))

	return true
}

func (psess *PackageSession) archreloc(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val *int64) bool {
	if ctxt.LinkMode == ld.LinkExternal {
		return false
	}
	switch r.Type {
	case objabi.R_CONST:
		*val = r.Add
		return true
	case objabi.R_GOTOFF:
		*val = psess.ld.Symaddr(r.Sym) + r.Add - psess.ld.Symaddr(ctxt.Syms.Lookup(".got", 0))
		return true
	}

	return false
}

func archrelocvariant(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	log.Fatalf("unexpected relocation variant")
	return t
}

func elfsetupplt(ctxt *ld.Link) {
	plt := ctxt.Syms.Lookup(".plt", 0)
	got := ctxt.Syms.Lookup(".got.plt", 0)
	if plt.Size == 0 {

		plt.AddUint8(0xff)

		plt.AddUint8(0x35)
		plt.AddAddrPlus(ctxt.Arch, got, 4)

		plt.AddUint8(0xff)

		plt.AddUint8(0x25)
		plt.AddAddrPlus(ctxt.Arch, got, 8)

		plt.AddUint32(ctxt.Arch, 0)

		got.AddAddrPlus(ctxt.Arch, ctxt.Syms.Lookup(".dynamic", 0), 0)

		got.AddUint32(ctxt.Arch, 0)
		got.AddUint32(ctxt.Arch, 0)
	}
}

func (psess *PackageSession) addpltsym(ctxt *ld.Link, s *sym.Symbol) {
	if s.Plt >= 0 {
		return
	}
	psess.ld.
		Adddynsym(ctxt, s)

	if ctxt.IsELF {
		plt := ctxt.Syms.Lookup(".plt", 0)
		got := ctxt.Syms.Lookup(".got.plt", 0)
		rel := ctxt.Syms.Lookup(".rel.plt", 0)
		if plt.Size == 0 {
			elfsetupplt(ctxt)
		}

		plt.AddUint8(0xff)

		plt.AddUint8(0x25)
		plt.AddAddrPlus(ctxt.Arch, got, got.Size)

		got.AddAddrPlus(ctxt.Arch, plt, plt.Size)

		plt.AddUint8(0x68)

		plt.AddUint32(ctxt.Arch, uint32(rel.Size))

		plt.AddUint8(0xe9)

		plt.AddUint32(ctxt.Arch, uint32(-(plt.Size + 4)))

		rel.AddAddrPlus(ctxt.Arch, got, got.Size-4)

		rel.AddUint32(ctxt.Arch, ld.ELF32_R_INFO(uint32(s.Dynid), uint32(elf.R_386_JMP_SLOT)))

		s.Plt = int32(plt.Size - 16)
	} else if ctxt.HeadType == objabi.Hdarwin {

		plt := ctxt.Syms.Lookup(".plt", 0)
		psess.
			addgotsym(ctxt, s)

		ctxt.Syms.Lookup(".linkedit.plt", 0).AddUint32(ctxt.Arch, uint32(s.Dynid))

		s.Plt = int32(plt.Size)

		plt.AddUint8(0xff)
		plt.AddUint8(0x25)
		plt.AddAddrPlus(ctxt.Arch, ctxt.Syms.Lookup(".got", 0), int64(s.Got))
	} else {
		psess.ld.
			Errorf(s, "addpltsym: unsupported binary format")
	}
}

func (psess *PackageSession) addgotsym(ctxt *ld.Link, s *sym.Symbol) {
	if s.Got >= 0 {
		return
	}
	psess.ld.
		Adddynsym(ctxt, s)
	got := ctxt.Syms.Lookup(".got", 0)
	s.Got = int32(got.Size)
	got.AddUint32(ctxt.Arch, 0)

	if ctxt.IsELF {
		rel := ctxt.Syms.Lookup(".rel", 0)
		rel.AddAddrPlus(ctxt.Arch, got, int64(s.Got))
		rel.AddUint32(ctxt.Arch, ld.ELF32_R_INFO(uint32(s.Dynid), uint32(elf.R_386_GLOB_DAT)))
	} else if ctxt.HeadType == objabi.Hdarwin {
		ctxt.Syms.Lookup(".linkedit.got", 0).AddUint32(ctxt.Arch, uint32(s.Dynid))
	} else {
		psess.ld.
			Errorf(s, "addgotsym: unsupported binary format")
	}
}

func (psess *PackageSession) asmb(ctxt *ld.Link) {
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f asmb\n", psess.ld.Cputime())
	}

	if ctxt.IsELF {
		psess.ld.
			Asmbelfsetup()
	}

	sect := psess.ld.Segtext.Sections[0]
	ctxt.Out.SeekSet(psess.ld, int64(sect.Vaddr-psess.ld.Segtext.Vaddr+psess.ld.Segtext.Fileoff))
	psess.ld.
		CodeblkPad(ctxt, int64(sect.Vaddr), int64(sect.Length), []byte{0xCC})
	for _, sect = range psess.ld.Segtext.Sections[1:] {
		ctxt.Out.SeekSet(psess.ld, int64(sect.Vaddr-psess.ld.Segtext.Vaddr+psess.ld.Segtext.Fileoff))
		psess.ld.
			Datblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	}

	if psess.ld.Segrodata.Filelen > 0 {
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f rodatblk\n", psess.ld.Cputime())
		}

		ctxt.Out.SeekSet(psess.ld, int64(psess.ld.Segrodata.Fileoff))
		psess.ld.
			Datblk(ctxt, int64(psess.ld.Segrodata.Vaddr), int64(psess.ld.Segrodata.Filelen))
	}
	if psess.ld.Segrelrodata.Filelen > 0 {
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f relrodatblk\n", psess.ld.Cputime())
		}
		ctxt.Out.SeekSet(psess.ld, int64(psess.ld.Segrelrodata.Fileoff))
		psess.ld.
			Datblk(ctxt, int64(psess.ld.Segrelrodata.Vaddr), int64(psess.ld.Segrelrodata.Filelen))
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f datblk\n", psess.ld.Cputime())
	}

	ctxt.Out.SeekSet(psess.ld, int64(psess.ld.Segdata.Fileoff))
	psess.ld.
		Datblk(ctxt, int64(psess.ld.Segdata.Vaddr), int64(psess.ld.Segdata.Filelen))

	ctxt.Out.SeekSet(psess.ld, int64(psess.ld.Segdwarf.Fileoff))
	psess.ld.
		Dwarfblk(ctxt, int64(psess.ld.Segdwarf.Vaddr), int64(psess.ld.Segdwarf.Filelen))

	machlink := uint32(0)
	if ctxt.HeadType == objabi.Hdarwin {
		machlink = uint32(psess.ld.Domacholink(ctxt))
	}
	psess.ld.
		Symsize = 0
	psess.ld.
		Spsize = 0
	psess.ld.
		Lcsize = 0
	symo := uint32(0)
	if !*psess.ld.FlagS {

		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f sym\n", psess.ld.Cputime())
		}
		switch ctxt.HeadType {
		default:
			if ctxt.IsELF {
				symo = uint32(psess.ld.Segdwarf.Fileoff + psess.ld.Segdwarf.Filelen)
				symo = uint32(ld.Rnd(int64(symo), int64(*psess.ld.FlagRound)))
			}

		case objabi.Hplan9:
			symo = uint32(psess.ld.Segdata.Fileoff + psess.ld.Segdata.Filelen)

		case objabi.Hdarwin:
			symo = uint32(psess.ld.Segdwarf.Fileoff + uint64(ld.Rnd(int64(psess.ld.Segdwarf.Filelen), int64(*psess.ld.FlagRound))) + uint64(machlink))

		case objabi.Hwindows:
			symo = uint32(psess.ld.Segdwarf.Fileoff + psess.ld.Segdwarf.Filelen)
			symo = uint32(ld.Rnd(int64(symo), psess.ld.PEFILEALIGN))
		}

		ctxt.Out.SeekSet(psess.ld, int64(symo))
		switch ctxt.HeadType {
		default:
			if ctxt.IsELF {
				if ctxt.Debugvlog != 0 {
					ctxt.Logf("%5.2f elfsym\n", psess.ld.Cputime())
				}
				psess.ld.
					Asmelfsym(ctxt)
				ctxt.Out.Flush(psess.ld)
				ctxt.Out.Write(psess.ld.Elfstrdat)

				if ctxt.LinkMode == ld.LinkExternal {
					psess.ld.
						Elfemitreloc(ctxt)
				}
			}

		case objabi.Hplan9:
			psess.ld.
				Asmplan9sym(ctxt)
			ctxt.Out.Flush(psess.ld)

			sym := ctxt.Syms.Lookup("pclntab", 0)
			if sym != nil {
				psess.ld.
					Lcsize = int32(len(sym.P))
				ctxt.Out.Write(sym.P)
				ctxt.Out.Flush(psess.ld)
			}

		case objabi.Hwindows:
			if ctxt.Debugvlog != 0 {
				ctxt.Logf("%5.2f dwarf\n", psess.ld.Cputime())
			}

		case objabi.Hdarwin:
			if ctxt.LinkMode == ld.LinkExternal {
				psess.ld.
					Machoemitreloc(ctxt)
			}
		}
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f headr\n", psess.ld.Cputime())
	}
	ctxt.Out.SeekSet(psess.ld, 0)
	switch ctxt.HeadType {
	default:
	case objabi.Hplan9:
		magic := int32(4*11*11 + 7)

		ctxt.Out.Write32b(uint32(magic))
		ctxt.Out.Write32b(uint32(psess.ld.Segtext.Filelen))
		ctxt.Out.Write32b(uint32(psess.ld.Segdata.Filelen))
		ctxt.Out.Write32b(uint32(psess.ld.Segdata.Length - psess.ld.Segdata.Filelen))
		ctxt.Out.Write32b(uint32(psess.ld.Symsize))
		ctxt.Out.Write32b(uint32(psess.ld.Entryvalue(ctxt)))
		ctxt.Out.Write32b(uint32(psess.ld.Spsize))
		ctxt.Out.Write32b(uint32(psess.ld.Lcsize))

	case objabi.Hdarwin:
		psess.ld.
			Asmbmacho(ctxt)

	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd,
		objabi.Hnacl:
		psess.ld.
			Asmbelf(ctxt, int64(symo))

	case objabi.Hwindows:
		psess.ld.
			Asmbpe(ctxt)
	}

	ctxt.Out.Flush(psess.ld)
}

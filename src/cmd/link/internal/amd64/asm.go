package amd64

import (
	"debug/elf"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"log"
)

func PADDR(x uint32) uint32 {
	return x &^ 0x80000000
}

func Addcall(ctxt *ld.Link, s *sym.Symbol, t *sym.Symbol) int64 {
	s.Attr |= sym.AttrReachable
	i := s.Size
	s.Size += 4
	s.Grow(s.Size)
	r := s.AddRel()
	r.Sym = t
	r.Off = int32(i)
	r.Type = objabi.R_CALL
	r.Siz = 4
	return i + int64(r.Siz)
}

func gentext(ctxt *ld.Link) {
	if !ctxt.DynlinkingGo() {
		return
	}
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

	o(0x48, 0x8d, 0x3d)
	initfunc.AddPCRelPlus(ctxt.Arch, ctxt.Moduledata, 0)

	o(0xe8)
	Addcall(ctxt, initfunc, addmoduledata)

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

	case 256 + objabi.RelocType(elf.R_X86_64_PC32):
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_X86_64_PC32 relocation for dynamic symbol %s", targ.Name)
		}

		if (targ.Type == 0 || targ.Type == sym.SXREF) && !targ.Attr.VisibilityHidden() {
			psess.ld.
				Errorf(s, "unknown symbol %s in pcrel", targ.Name)
		}
		r.Type = objabi.R_PCREL
		r.Add += 4
		return true

	case 256 + objabi.RelocType(elf.R_X86_64_PC64):
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_X86_64_PC64 relocation for dynamic symbol %s", targ.Name)
		}
		if targ.Type == 0 || targ.Type == sym.SXREF {
			psess.ld.
				Errorf(s, "unknown symbol %s in pcrel", targ.Name)
		}
		r.Type = objabi.R_PCREL
		r.Add += 8
		return true

	case 256 + objabi.RelocType(elf.R_X86_64_PLT32):
		r.Type = objabi.R_PCREL
		r.Add += 4
		if targ.Type == sym.SDYNIMPORT {
			psess.
				addpltsym(ctxt, targ)
			r.Sym = ctxt.Syms.Lookup(".plt", 0)
			r.Add += int64(targ.Plt)
		}

		return true

	case 256 + objabi.RelocType(elf.R_X86_64_GOTPCREL), 256 + objabi.RelocType(elf.R_X86_64_GOTPCRELX), 256 + objabi.RelocType(elf.R_X86_64_REX_GOTPCRELX):
		if targ.Type != sym.SDYNIMPORT {

			if r.Off >= 2 && s.P[r.Off-2] == 0x8b {

				s.P[r.Off-2] = 0x8d

				r.Type = objabi.R_PCREL
				r.Add += 4
				return true
			}
		}
		psess.
			addgotsym(ctxt, targ)

		r.Type = objabi.R_PCREL
		r.Sym = ctxt.Syms.Lookup(".got", 0)
		r.Add += 4
		r.Add += int64(targ.Got)
		return true

	case 256 + objabi.RelocType(elf.R_X86_64_64):
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_X86_64_64 relocation for dynamic symbol %s", targ.Name)
		}
		r.Type = objabi.R_ADDR
		return true

	case 512 + ld.MACHO_X86_64_RELOC_UNSIGNED*2 + 0,
		512 + ld.MACHO_X86_64_RELOC_SIGNED*2 + 0,
		512 + ld.MACHO_X86_64_RELOC_BRANCH*2 + 0:

		r.Type = objabi.R_ADDR

		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected reloc for dynamic symbol %s", targ.Name)
		}
		return true

	case 512 + ld.MACHO_X86_64_RELOC_BRANCH*2 + 1:
		if targ.Type == sym.SDYNIMPORT {
			psess.
				addpltsym(ctxt, targ)
			r.Sym = ctxt.Syms.Lookup(".plt", 0)
			r.Add = int64(targ.Plt)
			r.Type = objabi.R_PCREL
			return true
		}
		fallthrough

	case 512 + ld.MACHO_X86_64_RELOC_UNSIGNED*2 + 1,
		512 + ld.MACHO_X86_64_RELOC_SIGNED*2 + 1,
		512 + ld.MACHO_X86_64_RELOC_SIGNED_1*2 + 1,
		512 + ld.MACHO_X86_64_RELOC_SIGNED_2*2 + 1,
		512 + ld.MACHO_X86_64_RELOC_SIGNED_4*2 + 1:
		r.Type = objabi.R_PCREL

		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected pc-relative reloc for dynamic symbol %s", targ.Name)
		}
		return true

	case 512 + ld.MACHO_X86_64_RELOC_GOT_LOAD*2 + 1:
		if targ.Type != sym.SDYNIMPORT {

			if r.Off < 2 || s.P[r.Off-2] != 0x8b {
				psess.ld.
					Errorf(s, "unexpected GOT_LOAD reloc for non-dynamic symbol %s", targ.Name)
				return false
			}

			s.P[r.Off-2] = 0x8d
			r.Type = objabi.R_PCREL
			return true
		}
		fallthrough

	case 512 + ld.MACHO_X86_64_RELOC_GOT*2 + 1:
		if targ.Type != sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected GOT reloc for non-dynamic symbol %s", targ.Name)
		}
		psess.
			addgotsym(ctxt, targ)
		r.Type = objabi.R_PCREL
		r.Sym = ctxt.Syms.Lookup(".got", 0)
		r.Add += int64(targ.Got)
		return true
	}

	switch r.Type {
	case objabi.R_CALL,
		objabi.R_PCREL:
		if targ.Type != sym.SDYNIMPORT {

			return true
		}
		if ctxt.LinkMode == ld.LinkExternal {

			return true
		}
		psess.
			addpltsym(ctxt, targ)
		r.Sym = ctxt.Syms.Lookup(".plt", 0)
		r.Add = int64(targ.Plt)
		return true

	case objabi.R_ADDR:
		if s.Type == sym.STEXT && ctxt.IsELF {
			if ctxt.HeadType == objabi.Hsolaris {
				psess.
					addpltsym(ctxt, targ)
				r.Sym = ctxt.Syms.Lookup(".plt", 0)
				r.Add += int64(targ.Plt)
				return true
			}
			psess.
				addgotsym(ctxt, targ)

			r.Sym = ctxt.Syms.Lookup(".got", 0)
			r.Add += int64(targ.Got)
			return true
		}

		if ctxt.BuildMode == ld.BuildModePIE && ctxt.LinkMode == ld.LinkInternal {

			switch s.Name {
			case ".dynsym", ".rela", ".rela.plt", ".got.plt", ".dynamic":
				return false
			}
		} else {

			if s.Type != sym.SDATA && s.Type != sym.SRODATA {
				break
			}
		}

		if ctxt.IsELF {
			psess.ld.
				Adddynsym(ctxt, targ)
			rela := ctxt.Syms.Lookup(".rela", 0)
			rela.AddAddrPlus(ctxt.Arch, s, int64(r.Off))
			if r.Siz == 8 {
				rela.AddUint64(ctxt.Arch, ld.ELF64_R_INFO(uint32(targ.Dynid), uint32(elf.R_X86_64_64)))
			} else {

				rela.AddUint64(ctxt.Arch, ld.ELF64_R_INFO(uint32(targ.Dynid), uint32(elf.R_X86_64_32)))
			}
			rela.AddUint64(ctxt.Arch, uint64(r.Add))
			r.Type = 256
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
			got.AddUint64(ctxt.Arch, 0)
			ctxt.Syms.Lookup(".linkedit.got", 0).AddUint32(ctxt.Arch, uint32(targ.Dynid))
			r.Type = 256
			return true
		}
	}

	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	ctxt.Out.Write64(uint64(sectoff))

	elfsym := r.Xsym.ElfsymForReloc()
	switch r.Type {
	default:
		return false
	case objabi.R_ADDR:
		if r.Siz == 4 {
			ctxt.Out.Write64(uint64(elf.R_X86_64_32) | uint64(elfsym)<<32)
		} else if r.Siz == 8 {
			ctxt.Out.Write64(uint64(elf.R_X86_64_64) | uint64(elfsym)<<32)
		} else {
			return false
		}
	case objabi.R_TLS_LE:
		if r.Siz == 4 {
			ctxt.Out.Write64(uint64(elf.R_X86_64_TPOFF32) | uint64(elfsym)<<32)
		} else {
			return false
		}
	case objabi.R_TLS_IE:
		if r.Siz == 4 {
			ctxt.Out.Write64(uint64(elf.R_X86_64_GOTTPOFF) | uint64(elfsym)<<32)
		} else {
			return false
		}
	case objabi.R_CALL:
		if r.Siz == 4 {
			if r.Xsym.Type == sym.SDYNIMPORT {
				if ctxt.DynlinkingGo() {
					ctxt.Out.Write64(uint64(elf.R_X86_64_PLT32) | uint64(elfsym)<<32)
				} else {
					ctxt.Out.Write64(uint64(elf.R_X86_64_GOTPCREL) | uint64(elfsym)<<32)
				}
			} else {
				ctxt.Out.Write64(uint64(elf.R_X86_64_PC32) | uint64(elfsym)<<32)
			}
		} else {
			return false
		}
	case objabi.R_PCREL:
		if r.Siz == 4 {
			if r.Xsym.Type == sym.SDYNIMPORT && r.Xsym.ElfType == elf.STT_FUNC {
				ctxt.Out.Write64(uint64(elf.R_X86_64_PLT32) | uint64(elfsym)<<32)
			} else {
				ctxt.Out.Write64(uint64(elf.R_X86_64_PC32) | uint64(elfsym)<<32)
			}
		} else {
			return false
		}
	case objabi.R_GOTPCREL:
		if r.Siz == 4 {
			ctxt.Out.Write64(uint64(elf.R_X86_64_GOTPCREL) | uint64(elfsym)<<32)
		} else {
			return false
		}
	}

	ctxt.Out.Write64(uint64(r.Xadd))
	return true
}

func (psess *PackageSession) machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	var v uint32

	rs := r.Xsym

	if rs.Type == sym.SHOSTOBJ || r.Type == objabi.R_PCREL || r.Type == objabi.R_GOTPCREL || r.Type == objabi.R_CALL {
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
		v |= ld.MACHO_X86_64_RELOC_UNSIGNED << 28

	case objabi.R_CALL:
		v |= 1 << 24
		v |= ld.MACHO_X86_64_RELOC_BRANCH << 28

	case objabi.R_PCREL:
		v |= 1 << 24
		v |= ld.MACHO_X86_64_RELOC_SIGNED << 28
	case objabi.R_GOTPCREL:
		v |= 1 << 24
		v |= ld.MACHO_X86_64_RELOC_GOT_LOAD << 28
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
		v = ld.IMAGE_REL_AMD64_SECREL

	case objabi.R_ADDR:
		if r.Siz == 8 {
			v = ld.IMAGE_REL_AMD64_ADDR64
		} else {
			v = ld.IMAGE_REL_AMD64_ADDR32
		}

	case objabi.R_CALL,
		objabi.R_PCREL:
		v = ld.IMAGE_REL_AMD64_REL32
	}

	out.Write16(uint16(v))

	return true
}

func archreloc(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val *int64) bool {
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
		plt.AddPCRelPlus(ctxt.Arch, got, 8)

		plt.AddUint8(0xff)

		plt.AddUint8(0x25)
		plt.AddPCRelPlus(ctxt.Arch, got, 16)

		plt.AddUint32(ctxt.Arch, 0x00401f0f)

		got.AddAddrPlus(ctxt.Arch, ctxt.Syms.Lookup(".dynamic", 0), 0)

		got.AddUint64(ctxt.Arch, 0)
		got.AddUint64(ctxt.Arch, 0)
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
		rela := ctxt.Syms.Lookup(".rela.plt", 0)
		if plt.Size == 0 {
			elfsetupplt(ctxt)
		}

		plt.AddUint8(0xff)

		plt.AddUint8(0x25)
		plt.AddPCRelPlus(ctxt.Arch, got, got.Size)

		got.AddAddrPlus(ctxt.Arch, plt, plt.Size)

		plt.AddUint8(0x68)

		plt.AddUint32(ctxt.Arch, uint32((got.Size-24-8)/8))

		plt.AddUint8(0xe9)

		plt.AddUint32(ctxt.Arch, uint32(-(plt.Size + 4)))

		rela.AddAddrPlus(ctxt.Arch, got, got.Size-8)

		rela.AddUint64(ctxt.Arch, ld.ELF64_R_INFO(uint32(s.Dynid), uint32(elf.R_X86_64_JMP_SLOT)))
		rela.AddUint64(ctxt.Arch, 0)

		s.Plt = int32(plt.Size - 16)
	} else if ctxt.HeadType == objabi.Hdarwin {
		psess.
			addgotsym(ctxt, s)
		plt := ctxt.Syms.Lookup(".plt", 0)

		ctxt.Syms.Lookup(".linkedit.plt", 0).AddUint32(ctxt.Arch, uint32(s.Dynid))

		s.Plt = int32(plt.Size)

		plt.AddUint8(0xff)
		plt.AddUint8(0x25)
		plt.AddPCRelPlus(ctxt.Arch, ctxt.Syms.Lookup(".got", 0), int64(s.Got))
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
	got.AddUint64(ctxt.Arch, 0)

	if ctxt.IsELF {
		rela := ctxt.Syms.Lookup(".rela", 0)
		rela.AddAddrPlus(ctxt.Arch, got, int64(s.Got))
		rela.AddUint64(ctxt.Arch, ld.ELF64_R_INFO(uint32(s.Dynid), uint32(elf.R_X86_64_GLOB_DAT)))
		rela.AddUint64(ctxt.Arch, 0)
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

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f codeblk\n", psess.ld.Cputime())
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

	machlink := int64(0)
	if ctxt.HeadType == objabi.Hdarwin {
		machlink = psess.ld.Domacholink(ctxt)
	}

	switch ctxt.HeadType {
	default:
		psess.ld.
			Errorf(nil, "unknown header type %v", ctxt.HeadType)
		fallthrough

	case objabi.Hplan9:
		break

	case objabi.Hdarwin:
		psess.ld.
			Flag8 = true

	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd,
		objabi.Hdragonfly,
		objabi.Hsolaris:
		psess.ld.
			Flag8 = true

	case objabi.Hnacl,
		objabi.Hwindows:
		break
	}
	psess.ld.
		Symsize = 0
	psess.ld.
		Spsize = 0
	psess.ld.
		Lcsize = 0
	symo := int64(0)
	if !*psess.ld.FlagS {
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f sym\n", psess.ld.Cputime())
		}
		switch ctxt.HeadType {
		default:
		case objabi.Hplan9:
			*psess.ld.FlagS = true
			symo = int64(psess.ld.Segdata.Fileoff + psess.ld.Segdata.Filelen)

		case objabi.Hdarwin:
			symo = int64(psess.ld.Segdwarf.Fileoff + uint64(ld.Rnd(int64(psess.ld.Segdwarf.Filelen), int64(*psess.ld.FlagRound))) + uint64(machlink))

		case objabi.Hlinux,
			objabi.Hfreebsd,
			objabi.Hnetbsd,
			objabi.Hopenbsd,
			objabi.Hdragonfly,
			objabi.Hsolaris,
			objabi.Hnacl:
			symo = int64(psess.ld.Segdwarf.Fileoff + psess.ld.Segdwarf.Filelen)
			symo = ld.Rnd(symo, int64(*psess.ld.FlagRound))

		case objabi.Hwindows:
			symo = int64(psess.ld.Segdwarf.Fileoff + psess.ld.Segdwarf.Filelen)
			symo = ld.Rnd(symo, psess.ld.PEFILEALIGN)
		}

		ctxt.Out.SeekSet(psess.ld, symo)
		switch ctxt.HeadType {
		default:
			if ctxt.IsELF {
				ctxt.Out.SeekSet(psess.ld, symo)
				psess.ld.
					Asmelfsym(ctxt)
				ctxt.Out.Flush(psess.ld)
				ctxt.Out.Write(psess.ld.Elfstrdat)

				if ctxt.Debugvlog != 0 {
					ctxt.Logf("%5.2f dwarf\n", psess.ld.Cputime())
				}

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
		magic := int32(4*26*26 + 7)

		magic |= 0x00008000
		ctxt.Out.Write32b(uint32(magic))
		ctxt.Out.Write32b(uint32(psess.ld.Segtext.Filelen))
		ctxt.Out.Write32b(uint32(psess.ld.Segdata.Filelen))
		ctxt.Out.Write32b(uint32(psess.ld.Segdata.Length - psess.ld.Segdata.Filelen))
		ctxt.Out.Write32b(uint32(psess.ld.Symsize))
		vl := psess.ld.Entryvalue(ctxt)
		ctxt.Out.Write32b(PADDR(uint32(vl)))
		ctxt.Out.Write32b(uint32(psess.ld.Spsize))
		ctxt.Out.Write32b(uint32(psess.ld.Lcsize))
		ctxt.Out.Write64b(uint64(vl))

	case objabi.Hdarwin:
		psess.ld.
			Asmbmacho(ctxt)

	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd,
		objabi.Hdragonfly,
		objabi.Hsolaris,
		objabi.Hnacl:
		psess.ld.
			Asmbelf(ctxt, symo)

	case objabi.Hwindows:
		psess.ld.
			Asmbpe(ctxt)
	}

	ctxt.Out.Flush(psess.ld)
}

func tlsIEtoLE(s *sym.Symbol, off, size int) {

	if off < 3 {
		log.Fatal("R_X86_64_GOTTPOFF reloc not preceded by MOVQ or ADDQ instruction")
	}
	op := s.P[off-3 : off]
	reg := op[2] >> 3

	if op[1] == 0x8b || reg == 4 {

		if op[0] == 0x4c {
			op[0] = 0x49
		} else if size == 4 && op[0] == 0x44 {
			op[0] = 0x41
		}
		if op[1] == 0x8b {
			op[1] = 0xc7
		} else {
			op[1] = 0x81
		}
		op[2] = 0xc0 | reg
	} else {

		log.Fatalf("expected TLS IE op to be MOVQ, got %v", op)
	}
}

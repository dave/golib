package s390x

import (
	"debug/elf"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/sym"
)

// gentext generates assembly to append the local moduledata to the global
// moduledata linked list at initialization time. This is only done if the runtime
// is in a different module.
//
// <go.link.addmoduledata>:
// 	larl  %r2, <local.moduledata>
// 	jg    <runtime.addmoduledata@plt>
//	undef
//
// The job of appending the moduledata is delegated to runtime.addmoduledata.
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

	initfunc.AddUint8(0xc0)
	initfunc.AddUint8(0x20)
	lmd := initfunc.AddRel()
	lmd.Off = int32(initfunc.Size)
	lmd.Siz = 4
	lmd.Sym = ctxt.Moduledata
	lmd.Type = objabi.R_PCREL
	lmd.Variant = sym.RV_390_DBL
	lmd.Add = 2 + int64(lmd.Siz)
	initfunc.AddUint32(ctxt.Arch, 0)

	initfunc.AddUint8(0xc0)
	initfunc.AddUint8(0xf4)
	rel := initfunc.AddRel()
	rel.Off = int32(initfunc.Size)
	rel.Siz = 4
	rel.Sym = ctxt.Syms.Lookup("runtime.addmoduledata", 0)
	rel.Type = objabi.R_CALL
	rel.Variant = sym.RV_390_DBL
	rel.Add = 2 + int64(rel.Siz)
	initfunc.AddUint32(ctxt.Arch, 0)

	initfunc.AddUint32(ctxt.Arch, 0)
	if ctxt.BuildMode == ld.BuildModePlugin {
		ctxt.Textp = append(ctxt.Textp, addmoduledata)
	}
	ctxt.Textp = append(ctxt.Textp, initfunc)
	initarray_entry := ctxt.Syms.Lookup("go.link.addmoduledatainit", 0)
	initarray_entry.Attr |= sym.AttrLocal
	initarray_entry.Attr |= sym.AttrReachable
	initarray_entry.Type = sym.SINITARR
	initarray_entry.AddAddr(ctxt.Arch, initfunc)
}

func (psess *PackageSession) adddynrel(ctxt *ld.Link, s *sym.Symbol, r *sym.Reloc) bool {
	targ := r.Sym

	switch r.Type {
	default:
		if r.Type >= 256 {
			psess.ld.
				Errorf(s, "unexpected relocation type %d", r.Type)
			return false
		}

	case 256 + objabi.RelocType(elf.R_390_12),
		256 + objabi.RelocType(elf.R_390_GOT12):
		psess.ld.
			Errorf(s, "s390x 12-bit relocations have not been implemented (relocation type %d)", r.Type-256)
		return false

	case 256 + objabi.RelocType(elf.R_390_8),
		256 + objabi.RelocType(elf.R_390_16),
		256 + objabi.RelocType(elf.R_390_32),
		256 + objabi.RelocType(elf.R_390_64):
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_390_nn relocation for dynamic symbol %s", targ.Name)
		}
		r.Type = objabi.R_ADDR
		return true

	case 256 + objabi.RelocType(elf.R_390_PC16),
		256 + objabi.RelocType(elf.R_390_PC32),
		256 + objabi.RelocType(elf.R_390_PC64):
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_390_PCnn relocation for dynamic symbol %s", targ.Name)
		}

		if (targ.Type == 0 || targ.Type == sym.SXREF) && !targ.Attr.VisibilityHidden() {
			psess.ld.
				Errorf(s, "unknown symbol %s in pcrel", targ.Name)
		}
		r.Type = objabi.R_PCREL
		r.Add += int64(r.Siz)
		return true

	case 256 + objabi.RelocType(elf.R_390_GOT16),
		256 + objabi.RelocType(elf.R_390_GOT32),
		256 + objabi.RelocType(elf.R_390_GOT64):
		psess.ld.
			Errorf(s, "unimplemented S390x relocation: %v", r.Type-256)
		return true

	case 256 + objabi.RelocType(elf.R_390_PLT16DBL),
		256 + objabi.RelocType(elf.R_390_PLT32DBL):
		r.Type = objabi.R_PCREL
		r.Variant = sym.RV_390_DBL
		r.Add += int64(r.Siz)
		if targ.Type == sym.SDYNIMPORT {
			psess.
				addpltsym(ctxt, targ)
			r.Sym = ctxt.Syms.Lookup(".plt", 0)
			r.Add += int64(targ.Plt)
		}
		return true

	case 256 + objabi.RelocType(elf.R_390_PLT32),
		256 + objabi.RelocType(elf.R_390_PLT64):
		r.Type = objabi.R_PCREL
		r.Add += int64(r.Siz)
		if targ.Type == sym.SDYNIMPORT {
			psess.
				addpltsym(ctxt, targ)
			r.Sym = ctxt.Syms.Lookup(".plt", 0)
			r.Add += int64(targ.Plt)
		}
		return true

	case 256 + objabi.RelocType(elf.R_390_COPY):
		psess.ld.
			Errorf(s, "unimplemented S390x relocation: %v", r.Type-256)
		return false

	case 256 + objabi.RelocType(elf.R_390_GLOB_DAT):
		psess.ld.
			Errorf(s, "unimplemented S390x relocation: %v", r.Type-256)
		return false

	case 256 + objabi.RelocType(elf.R_390_JMP_SLOT):
		psess.ld.
			Errorf(s, "unimplemented S390x relocation: %v", r.Type-256)
		return false

	case 256 + objabi.RelocType(elf.R_390_RELATIVE):
		psess.ld.
			Errorf(s, "unimplemented S390x relocation: %v", r.Type-256)
		return false

	case 256 + objabi.RelocType(elf.R_390_GOTOFF):
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_390_GOTOFF relocation for dynamic symbol %s", targ.Name)
		}
		r.Type = objabi.R_GOTOFF
		return true

	case 256 + objabi.RelocType(elf.R_390_GOTPC):
		r.Type = objabi.R_PCREL
		r.Sym = ctxt.Syms.Lookup(".got", 0)
		r.Add += int64(r.Siz)
		return true

	case 256 + objabi.RelocType(elf.R_390_PC16DBL),
		256 + objabi.RelocType(elf.R_390_PC32DBL):
		r.Type = objabi.R_PCREL
		r.Variant = sym.RV_390_DBL
		r.Add += int64(r.Siz)
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_390_PCnnDBL relocation for dynamic symbol %s", targ.Name)
		}
		return true

	case 256 + objabi.RelocType(elf.R_390_GOTPCDBL):
		r.Type = objabi.R_PCREL
		r.Variant = sym.RV_390_DBL
		r.Sym = ctxt.Syms.Lookup(".got", 0)
		r.Add += int64(r.Siz)
		return true

	case 256 + objabi.RelocType(elf.R_390_GOTENT):
		psess.
			addgotsym(ctxt, targ)

		r.Type = objabi.R_PCREL
		r.Variant = sym.RV_390_DBL
		r.Sym = ctxt.Syms.Lookup(".got", 0)
		r.Add += int64(targ.Got)
		r.Add += int64(r.Siz)
		return true
	}

	if targ.Type != sym.SDYNIMPORT {
		return true
	}

	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	ctxt.Out.Write64(uint64(sectoff))

	elfsym := r.Xsym.ElfsymForReloc()
	switch r.Type {
	default:
		return false
	case objabi.R_TLS_LE:
		switch r.Siz {
		default:
			return false
		case 4:

			ctxt.Out.Write64(uint64(elf.R_390_TLS_LE32) | uint64(elfsym)<<32)
		case 8:

			ctxt.Out.Write64(uint64(elf.R_390_TLS_LE64) | uint64(elfsym)<<32)
		}
	case objabi.R_TLS_IE:
		switch r.Siz {
		default:
			return false
		case 4:
			ctxt.Out.Write64(uint64(elf.R_390_TLS_IEENT) | uint64(elfsym)<<32)
		}
	case objabi.R_ADDR:
		switch r.Siz {
		default:
			return false
		case 4:
			ctxt.Out.Write64(uint64(elf.R_390_32) | uint64(elfsym)<<32)
		case 8:
			ctxt.Out.Write64(uint64(elf.R_390_64) | uint64(elfsym)<<32)
		}
	case objabi.R_GOTPCREL:
		if r.Siz == 4 {
			ctxt.Out.Write64(uint64(elf.R_390_GOTENT) | uint64(elfsym)<<32)
		} else {
			return false
		}
	case objabi.R_PCREL, objabi.R_PCRELDBL, objabi.R_CALL:
		elfrel := elf.R_390_NONE
		isdbl := r.Variant&sym.RV_TYPE_MASK == sym.RV_390_DBL

		switch r.Type {
		case objabi.R_PCRELDBL, objabi.R_CALL:
			isdbl = true
		}
		if r.Xsym.Type == sym.SDYNIMPORT && (r.Xsym.ElfType == elf.STT_FUNC || r.Type == objabi.R_CALL) {
			if isdbl {
				switch r.Siz {
				case 2:
					elfrel = elf.R_390_PLT16DBL
				case 4:
					elfrel = elf.R_390_PLT32DBL
				}
			} else {
				switch r.Siz {
				case 4:
					elfrel = elf.R_390_PLT32
				case 8:
					elfrel = elf.R_390_PLT64
				}
			}
		} else {
			if isdbl {
				switch r.Siz {
				case 2:
					elfrel = elf.R_390_PC16DBL
				case 4:
					elfrel = elf.R_390_PC32DBL
				}
			} else {
				switch r.Siz {
				case 2:
					elfrel = elf.R_390_PC16
				case 4:
					elfrel = elf.R_390_PC32
				case 8:
					elfrel = elf.R_390_PC64
				}
			}
		}
		if elfrel == elf.R_390_NONE {
			return false
		}
		ctxt.Out.Write64(uint64(elfrel) | uint64(elfsym)<<32)
	}

	ctxt.Out.Write64(uint64(r.Xadd))
	return true
}

func elfsetupplt(ctxt *ld.Link) {
	plt := ctxt.Syms.Lookup(".plt", 0)
	got := ctxt.Syms.Lookup(".got", 0)
	if plt.Size == 0 {

		plt.AddUint8(0xe3)
		plt.AddUint8(0x10)
		plt.AddUint8(0xf0)
		plt.AddUint8(0x38)
		plt.AddUint8(0x00)
		plt.AddUint8(0x24)

		plt.AddUint8(0xc0)
		plt.AddUint8(0x10)
		plt.AddPCRelPlus(ctxt.Arch, got, 6)

		plt.AddUint8(0xd2)
		plt.AddUint8(0x07)
		plt.AddUint8(0xf0)
		plt.AddUint8(0x30)
		plt.AddUint8(0x10)
		plt.AddUint8(0x08)

		plt.AddUint8(0xe3)
		plt.AddUint8(0x10)
		plt.AddUint8(0x10)
		plt.AddUint8(0x10)
		plt.AddUint8(0x00)
		plt.AddUint8(0x04)

		plt.AddUint8(0x07)
		plt.AddUint8(0xf1)

		plt.AddUint8(0x07)
		plt.AddUint8(0x00)

		plt.AddUint8(0x07)
		plt.AddUint8(0x00)

		plt.AddUint8(0x07)
		plt.AddUint8(0x00)

		got.AddAddrPlus(ctxt.Arch, ctxt.Syms.Lookup(".dynamic", 0), 0)

		got.AddUint64(ctxt.Arch, 0)
		got.AddUint64(ctxt.Arch, 0)
	}
}

func machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	return false
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

func (psess *PackageSession) archrelocvariant(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	switch r.Variant & sym.RV_TYPE_MASK {
	default:
		psess.ld.
			Errorf(s, "unexpected relocation variant %d", r.Variant)
		return t

	case sym.RV_NONE:
		return t

	case sym.RV_390_DBL:
		if (t & 1) != 0 {
			psess.ld.
				Errorf(s, "%s+%v is not 2-byte aligned", r.Sym.Name, r.Sym.Value)
		}
		return t >> 1
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
		got := ctxt.Syms.Lookup(".got", 0)
		rela := ctxt.Syms.Lookup(".rela.plt", 0)
		if plt.Size == 0 {
			elfsetupplt(ctxt)
		}

		plt.AddUint8(0xc0)
		plt.AddUint8(0x10)
		plt.AddPCRelPlus(ctxt.Arch, got, got.Size+6)

		got.AddAddrPlus(ctxt.Arch, plt, plt.Size+8)

		plt.AddUint8(0xe3)
		plt.AddUint8(0x10)
		plt.AddUint8(0x10)
		plt.AddUint8(0x00)
		plt.AddUint8(0x00)
		plt.AddUint8(0x04)

		plt.AddUint8(0x07)
		plt.AddUint8(0xf1)

		plt.AddUint8(0x0d)
		plt.AddUint8(0x10)

		plt.AddUint8(0xe3)
		plt.AddUint8(0x10)
		plt.AddUint8(0x10)
		plt.AddUint8(0x0c)
		plt.AddUint8(0x00)
		plt.AddUint8(0x14)

		plt.AddUint8(0xc0)
		plt.AddUint8(0xf4)

		plt.AddUint32(ctxt.Arch, uint32(-((plt.Size - 2) >> 1)))

		plt.AddUint32(ctxt.Arch, uint32(rela.Size))

		rela.AddAddrPlus(ctxt.Arch, got, got.Size-8)

		rela.AddUint64(ctxt.Arch, ld.ELF64_R_INFO(uint32(s.Dynid), uint32(elf.R_390_JMP_SLOT)))
		rela.AddUint64(ctxt.Arch, 0)

		s.Plt = int32(plt.Size - 32)

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
		rela.AddUint64(ctxt.Arch, ld.ELF64_R_INFO(uint32(s.Dynid), uint32(elf.R_390_GLOB_DAT)))
		rela.AddUint64(ctxt.Arch, 0)
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
		Codeblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
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
			ctxt.Logf("%5.2f rodatblk\n", psess.ld.Cputime())
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
	psess.ld.
		Symsize = 0
	psess.ld.
		Lcsize = 0
	symo := uint32(0)
	if !*psess.ld.FlagS {
		if !ctxt.IsELF {
			psess.ld.
				Errorf(nil, "unsupported executable format")
		}
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f sym\n", psess.ld.Cputime())
		}
		symo = uint32(psess.ld.Segdwarf.Fileoff + psess.ld.Segdwarf.Filelen)
		symo = uint32(ld.Rnd(int64(symo), int64(*psess.ld.FlagRound)))

		ctxt.Out.SeekSet(psess.ld, int64(symo))
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f elfsym\n", psess.ld.Cputime())
		}
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

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f header\n", psess.ld.Cputime())
	}
	ctxt.Out.SeekSet(psess.ld, 0)
	switch ctxt.HeadType {
	default:
		psess.ld.
			Errorf(nil, "unsupported operating system")
	case objabi.Hlinux:
		psess.ld.
			Asmbelf(ctxt, int64(symo))
	}

	ctxt.Out.Flush(psess.ld)
	if *psess.ld.FlagC {
		fmt.Printf("textsize=%d\n", psess.ld.Segtext.Filelen)
		fmt.Printf("datsize=%d\n", psess.ld.Segdata.Filelen)
		fmt.Printf("bsssize=%d\n", psess.ld.Segdata.Length-psess.ld.Segdata.Filelen)
		fmt.Printf("symsize=%d\n", psess.ld.Symsize)
		fmt.Printf("lcsize=%d\n", psess.ld.Lcsize)
		fmt.Printf("total=%d\n", psess.ld.Segtext.Filelen+psess.ld.Segdata.Length+uint64(psess.ld.Symsize)+uint64(psess.ld.Lcsize))
	}
}

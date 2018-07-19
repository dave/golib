package mips64

import (
	"debug/elf"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"log"
)

func gentext(ctxt *ld.Link) {}

func adddynrel(ctxt *ld.Link, s *sym.Symbol, r *sym.Reloc) bool {
	log.Fatalf("adddynrel not implemented")
	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {

	ctxt.Out.Write64(uint64(sectoff))

	elfsym := r.Xsym.ElfsymForReloc()
	ctxt.Out.Write32(uint32(elfsym))
	ctxt.Out.Write8(0)
	ctxt.Out.Write8(0)
	ctxt.Out.Write8(0)
	switch r.Type {
	default:
		return false
	case objabi.R_ADDR:
		switch r.Siz {
		case 4:
			ctxt.Out.Write8(uint8(elf.R_MIPS_32))
		case 8:
			ctxt.Out.Write8(uint8(elf.R_MIPS_64))
		default:
			return false
		}
	case objabi.R_ADDRMIPS:
		ctxt.Out.Write8(uint8(elf.R_MIPS_LO16))
	case objabi.R_ADDRMIPSU:
		ctxt.Out.Write8(uint8(elf.R_MIPS_HI16))
	case objabi.R_ADDRMIPSTLS:
		ctxt.Out.Write8(uint8(elf.R_MIPS_TLS_TPREL_LO16))
	case objabi.R_CALLMIPS,
		objabi.R_JMPMIPS:
		ctxt.Out.Write8(uint8(elf.R_MIPS_26))
	}
	ctxt.Out.Write64(uint64(r.Xadd))

	return true
}

func elfsetupplt(ctxt *ld.Link) {
	return
}

func machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	return false
}

func (psess *PackageSession) archreloc(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val *int64) bool {
	if ctxt.LinkMode == ld.LinkExternal {
		switch r.Type {
		default:
			return false
		case objabi.R_ADDRMIPS,
			objabi.R_ADDRMIPSU:
			r.Done = false

			rs := r.Sym
			r.Xadd = r.Add
			for rs.Outer != nil {
				r.Xadd += psess.ld.Symaddr(rs) - psess.ld.Symaddr(rs.Outer)
				rs = rs.Outer
			}

			if rs.Type != sym.SHOSTOBJ && rs.Type != sym.SDYNIMPORT && rs.Sect == nil {
				psess.ld.
					Errorf(s, "missing section for %s", rs.Name)
			}
			r.Xsym = rs

			return true
		case objabi.R_ADDRMIPSTLS,
			objabi.R_CALLMIPS,
			objabi.R_JMPMIPS:
			r.Done = false
			r.Xsym = r.Sym
			r.Xadd = r.Add
			return true
		}
	}

	switch r.Type {
	case objabi.R_CONST:
		*val = r.Add
		return true
	case objabi.R_GOTOFF:
		*val = psess.ld.Symaddr(r.Sym) + r.Add - psess.ld.Symaddr(ctxt.Syms.Lookup(".got", 0))
		return true
	case objabi.R_ADDRMIPS,
		objabi.R_ADDRMIPSU:
		t := psess.ld.Symaddr(r.Sym) + r.Add
		o1 := ctxt.Arch.ByteOrder.Uint32(s.P[r.Off:])
		if r.Type == objabi.R_ADDRMIPS {
			*val = int64(o1&0xffff0000 | uint32(t)&0xffff)
		} else {
			*val = int64(o1&0xffff0000 | uint32((t+1<<15)>>16)&0xffff)
		}
		return true
	case objabi.R_ADDRMIPSTLS:

		t := psess.ld.Symaddr(r.Sym) + r.Add - 0x7000
		if t < -32768 || t >= 32678 {
			psess.ld.
				Errorf(s, "TLS offset out of range %d", t)
		}
		o1 := ctxt.Arch.ByteOrder.Uint32(s.P[r.Off:])
		*val = int64(o1&0xffff0000 | uint32(t)&0xffff)
		return true
	case objabi.R_CALLMIPS,
		objabi.R_JMPMIPS:

		t := psess.ld.Symaddr(r.Sym) + r.Add
		o1 := ctxt.Arch.ByteOrder.Uint32(s.P[r.Off:])
		*val = int64(o1&0xfc000000 | uint32(t>>2)&^0xfc000000)
		return true
	}

	return false
}

func archrelocvariant(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	return -1
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
		}
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f header\n", psess.ld.Cputime())
	}
	ctxt.Out.SeekSet(psess.ld, 0)
	switch ctxt.HeadType {
	default:
	case objabi.Hplan9:
		magic := uint32(4*18*18 + 7)
		if ctxt.Arch == psess.sys.ArchMIPS64LE {
			magic = uint32(4*26*26 + 7)
		}
		ctxt.Out.Write32(magic)
		ctxt.Out.Write32(uint32(psess.ld.Segtext.Filelen))
		ctxt.Out.Write32(uint32(psess.ld.Segdata.Filelen))
		ctxt.Out.Write32(uint32(psess.ld.Segdata.Length - psess.ld.Segdata.Filelen))
		ctxt.Out.Write32(uint32(psess.ld.Symsize))
		ctxt.Out.Write32(uint32(psess.ld.Entryvalue(ctxt)))
		ctxt.Out.Write32(0)
		ctxt.Out.Write32(uint32(psess.ld.Lcsize))

	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd,
		objabi.Hnacl:
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

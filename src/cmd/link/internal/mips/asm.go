// Inferno utils/5l/asm.c
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/5l/asm.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2016 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package mips

import (
	"debug/elf"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"log"
)

func gentext(ctxt *ld.Link) {
	return
}

func adddynrel(ctxt *ld.Link, s *sym.Symbol, r *sym.Reloc) bool {
	log.Fatalf("adddynrel not implemented")
	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	ctxt.Out.Write32(uint32(sectoff))

	elfsym := r.Xsym.ElfsymForReloc()
	switch r.Type {
	default:
		return false
	case objabi.R_ADDR:
		if r.Siz != 4 {
			return false
		}
		ctxt.Out.Write32(uint32(elf.R_MIPS_32) | uint32(elfsym)<<8)
	case objabi.R_ADDRMIPS:
		ctxt.Out.Write32(uint32(elf.R_MIPS_LO16) | uint32(elfsym)<<8)
	case objabi.R_ADDRMIPSU:
		ctxt.Out.Write32(uint32(elf.R_MIPS_HI16) | uint32(elfsym)<<8)
	case objabi.R_ADDRMIPSTLS:
		ctxt.Out.Write32(uint32(elf.R_MIPS_TLS_TPREL_LO16) | uint32(elfsym)<<8)
	case objabi.R_CALLMIPS, objabi.R_JMPMIPS:
		ctxt.Out.Write32(uint32(elf.R_MIPS_26) | uint32(elfsym)<<8)
	}

	return true
}

func elfsetupplt(ctxt *ld.Link) {
	return
}

func machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	return false
}

func applyrel(arch *sys.Arch, r *sym.Reloc, s *sym.Symbol, val *int64, t int64) {
	o := arch.ByteOrder.Uint32(s.P[r.Off:])
	switch r.Type {
	case objabi.R_ADDRMIPS, objabi.R_ADDRMIPSTLS:
		*val = int64(o&0xffff0000 | uint32(t)&0xffff)
	case objabi.R_ADDRMIPSU:
		*val = int64(o&0xffff0000 | uint32((t+(1<<15))>>16)&0xffff)
	case objabi.R_CALLMIPS, objabi.R_JMPMIPS:
		*val = int64(o&0xfc000000 | uint32(t>>2)&^0xfc000000)
	}
}

func (pstate *PackageState) archreloc(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val *int64) bool {
	if ctxt.LinkMode == ld.LinkExternal {
		switch r.Type {
		default:
			return false
		case objabi.R_ADDRMIPS, objabi.R_ADDRMIPSU:
			r.Done = false

			// set up addend for eventual relocation via outer symbol.
			rs := r.Sym
			r.Xadd = r.Add
			for rs.Outer != nil {
				r.Xadd += pstate.ld.Symaddr(rs) - pstate.ld.Symaddr(rs.Outer)
				rs = rs.Outer
			}

			if rs.Type != sym.SHOSTOBJ && rs.Type != sym.SDYNIMPORT && rs.Sect == nil {
				pstate.ld.Errorf(s, "missing section for %s", rs.Name)
			}
			r.Xsym = rs
			applyrel(ctxt.Arch, r, s, val, r.Xadd)
			return true
		case objabi.R_ADDRMIPSTLS, objabi.R_CALLMIPS, objabi.R_JMPMIPS:
			r.Done = false
			r.Xsym = r.Sym
			r.Xadd = r.Add
			applyrel(ctxt.Arch, r, s, val, r.Add)
			return true
		}
	}

	switch r.Type {
	case objabi.R_CONST:
		*val = r.Add
		return true
	case objabi.R_GOTOFF:
		*val = pstate.ld.Symaddr(r.Sym) + r.Add - pstate.ld.Symaddr(ctxt.Syms.Lookup(".got", 0))
		return true
	case objabi.R_ADDRMIPS, objabi.R_ADDRMIPSU:
		t := pstate.ld.Symaddr(r.Sym) + r.Add
		applyrel(ctxt.Arch, r, s, val, t)
		return true
	case objabi.R_CALLMIPS, objabi.R_JMPMIPS:
		t := pstate.ld.Symaddr(r.Sym) + r.Add

		if t&3 != 0 {
			pstate.ld.Errorf(s, "direct call is not aligned: %s %x", r.Sym.Name, t)
		}

		// check if target address is in the same 256 MB region as the next instruction
		if (s.Value+int64(r.Off)+4)&0xf0000000 != (t & 0xf0000000) {
			pstate.ld.Errorf(s, "direct call too far: %s %x", r.Sym.Name, t)
		}

		applyrel(ctxt.Arch, r, s, val, t)
		return true
	case objabi.R_ADDRMIPSTLS:
		// thread pointer is at 0x7000 offset from the start of TLS data area
		t := pstate.ld.Symaddr(r.Sym) + r.Add - 0x7000
		if t < -32768 || t >= 32678 {
			pstate.ld.Errorf(s, "TLS offset out of range %d", t)
		}
		applyrel(ctxt.Arch, r, s, val, t)
		return true
	}

	return false
}

func archrelocvariant(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	return -1
}

func (pstate *PackageState) asmb(ctxt *ld.Link) {
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f asmb\n", pstate.ld.Cputime())
	}

	if ctxt.IsELF {
		pstate.ld.Asmbelfsetup()
	}

	sect := pstate.ld.Segtext.Sections[0]
	ctxt.Out.SeekSet(pstate.ld, int64(sect.Vaddr-pstate.ld.Segtext.Vaddr+pstate.ld.Segtext.Fileoff))
	pstate.ld.Codeblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	for _, sect = range pstate.ld.Segtext.Sections[1:] {
		ctxt.Out.SeekSet(pstate.ld, int64(sect.Vaddr-pstate.ld.Segtext.Vaddr+pstate.ld.Segtext.Fileoff))
		pstate.ld.Datblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	}

	if pstate.ld.Segrodata.Filelen > 0 {
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f rodatblk\n", pstate.ld.Cputime())
		}

		ctxt.Out.SeekSet(pstate.ld, int64(pstate.ld.Segrodata.Fileoff))
		pstate.ld.Datblk(ctxt, int64(pstate.ld.Segrodata.Vaddr), int64(pstate.ld.Segrodata.Filelen))
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f datblk\n", pstate.ld.Cputime())
	}

	ctxt.Out.SeekSet(pstate.ld, int64(pstate.ld.Segdata.Fileoff))
	pstate.ld.Datblk(ctxt, int64(pstate.ld.Segdata.Vaddr), int64(pstate.ld.Segdata.Filelen))

	ctxt.Out.SeekSet(pstate.ld, int64(pstate.ld.Segdwarf.Fileoff))
	pstate.ld.Dwarfblk(ctxt, int64(pstate.ld.Segdwarf.Vaddr), int64(pstate.ld.Segdwarf.Filelen))

	/* output symbol table */
	pstate.ld.Symsize = 0

	pstate.ld.Lcsize = 0
	symo := uint32(0)
	if !*pstate.ld.FlagS {
		if !ctxt.IsELF {
			pstate.ld.Errorf(nil, "unsupported executable format")
		}
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f sym\n", pstate.ld.Cputime())
		}
		symo = uint32(pstate.ld.Segdwarf.Fileoff + pstate.ld.Segdwarf.Filelen)
		symo = uint32(ld.Rnd(int64(symo), int64(*pstate.ld.FlagRound)))

		ctxt.Out.SeekSet(pstate.ld, int64(symo))
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f elfsym\n", pstate.ld.Cputime())
		}
		pstate.ld.Asmelfsym(ctxt)
		ctxt.Out.Flush(pstate.ld)
		ctxt.Out.Write(pstate.ld.Elfstrdat)

		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f dwarf\n", pstate.ld.Cputime())
		}

		if ctxt.LinkMode == ld.LinkExternal {
			pstate.ld.Elfemitreloc(ctxt)
		}
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f header\n", pstate.ld.Cputime())
	}

	ctxt.Out.SeekSet(pstate.ld, 0)
	switch ctxt.HeadType {
	default:
		pstate.ld.Errorf(nil, "unsupported operating system")
	case objabi.Hlinux:
		pstate.ld.Asmbelf(ctxt, int64(symo))
	}

	ctxt.Out.Flush(pstate.ld)
	if *pstate.ld.FlagC {
		fmt.Printf("textsize=%d\n", pstate.ld.Segtext.Filelen)
		fmt.Printf("datsize=%d\n", pstate.ld.Segdata.Filelen)
		fmt.Printf("bsssize=%d\n", pstate.ld.Segdata.Length-pstate.ld.Segdata.Filelen)
		fmt.Printf("symsize=%d\n", pstate.ld.Symsize)
		fmt.Printf("lcsize=%d\n", pstate.ld.Lcsize)
		fmt.Printf("total=%d\n", pstate.ld.Segtext.Filelen+pstate.ld.Segdata.Length+uint64(pstate.ld.Symsize)+uint64(pstate.ld.Lcsize))
	}
}

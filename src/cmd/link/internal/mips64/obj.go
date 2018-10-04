// Inferno utils/5l/obj.c
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/5l/obj.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
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

package mips64

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
)

func (pstate *PackageState) Init() (*sys.Arch, ld.Arch) {
	arch := pstate.sys.ArchMIPS64
	if pstate.objabi.GOARCH == "mips64le" {
		arch = pstate.sys.ArchMIPS64LE
	}

	theArch := ld.Arch{
		Funcalign:        funcAlign,
		Maxalign:         maxAlign,
		Minalign:         minAlign,
		Dwarfregsp:       dwarfRegSP,
		Dwarfreglr:       dwarfRegLR,
		Adddynrel:        adddynrel,
		Archinit:         pstate.archinit,
		Archreloc:        pstate.archreloc,
		Archrelocvariant: archrelocvariant,
		Asmb:             pstate.asmb,
		Elfreloc1:        elfreloc1,
		Elfsetupplt:      elfsetupplt,
		Gentext:          gentext,
		Machoreloc1:      machoreloc1,

		Linuxdynld:     "/lib64/ld64.so.1",
		Freebsddynld:   "XXX",
		Openbsddynld:   "XXX",
		Netbsddynld:    "XXX",
		Dragonflydynld: "XXX",
		Solarisdynld:   "XXX",
	}

	return arch, theArch
}

func (pstate *PackageState) archinit(ctxt *ld.Link) {
	switch ctxt.HeadType {
	default:
		pstate.ld.Exitf("unknown -H option: %v", ctxt.HeadType)

	case objabi.Hplan9: /* plan 9 */
		pstate.ld.HEADR = 32

		if *pstate.ld.FlagTextAddr == -1 {
			*pstate.ld.FlagTextAddr = 16*1024 + int64(pstate.ld.HEADR)
		}
		if *pstate.ld.FlagDataAddr == -1 {
			*pstate.ld.FlagDataAddr = 0
		}
		if *pstate.ld.FlagRound == -1 {
			*pstate.ld.FlagRound = 16 * 1024
		}

	case objabi.Hlinux: /* mips64 elf */
		pstate.ld.Elfinit(ctxt)
		pstate.ld.HEADR = ld.ELFRESERVE
		if *pstate.ld.FlagTextAddr == -1 {
			*pstate.ld.FlagTextAddr = 0x10000 + int64(pstate.ld.HEADR)
		}
		if *pstate.ld.FlagDataAddr == -1 {
			*pstate.ld.FlagDataAddr = 0
		}
		if *pstate.ld.FlagRound == -1 {
			*pstate.ld.FlagRound = 0x10000
		}

	case objabi.Hnacl:
		pstate.ld.Elfinit(ctxt)
		pstate.ld.HEADR = 0x10000
		pstate.ld.Funcalign = 16
		if *pstate.ld.FlagTextAddr == -1 {
			*pstate.ld.FlagTextAddr = 0x20000
		}
		if *pstate.ld.FlagDataAddr == -1 {
			*pstate.ld.FlagDataAddr = 0
		}
		if *pstate.ld.FlagRound == -1 {
			*pstate.ld.FlagRound = 0x10000
		}
	}

	if *pstate.ld.FlagDataAddr != 0 && *pstate.ld.FlagRound != 0 {
		fmt.Printf("warning: -D0x%x is ignored because of -R0x%x\n", uint64(*pstate.ld.FlagDataAddr), uint32(*pstate.ld.FlagRound))
	}
}

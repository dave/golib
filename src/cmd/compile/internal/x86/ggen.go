// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x86

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/x86"
)

func (pstate *PackageState) zerorange(pp *gc.Progs, p *obj.Prog, off, cnt int64, ax *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	if *ax == 0 {
		p = pp.Appendpp(pstate.gc, p, x86.AMOVL, obj.TYPE_CONST, 0, 0, obj.TYPE_REG, x86.REG_AX, 0)
		*ax = 1
	}

	if cnt <= int64(4*pstate.gc.Widthreg) {
		for i := int64(0); i < cnt; i += int64(pstate.gc.Widthreg) {
			p = pp.Appendpp(pstate.gc, p, x86.AMOVL, obj.TYPE_REG, x86.REG_AX, 0, obj.TYPE_MEM, x86.REG_SP, off+i)
		}
	} else if !pstate.gc.Nacl && cnt <= int64(128*pstate.gc.Widthreg) {
		p = pp.Appendpp(pstate.gc, p, x86.ALEAL, obj.TYPE_MEM, x86.REG_SP, off, obj.TYPE_REG, x86.REG_DI, 0)
		p = pp.Appendpp(pstate.gc, p, obj.ADUFFZERO, obj.TYPE_NONE, 0, 0, obj.TYPE_ADDR, 0, 1*(128-cnt/int64(pstate.gc.Widthreg)))
		p.To.Sym = pstate.gc.Duffzero
	} else {
		p = pp.Appendpp(pstate.gc, p, x86.AMOVL, obj.TYPE_CONST, 0, cnt/int64(pstate.gc.Widthreg), obj.TYPE_REG, x86.REG_CX, 0)
		p = pp.Appendpp(pstate.gc, p, x86.ALEAL, obj.TYPE_MEM, x86.REG_SP, off, obj.TYPE_REG, x86.REG_DI, 0)
		p = pp.Appendpp(pstate.gc, p, x86.AREP, obj.TYPE_NONE, 0, 0, obj.TYPE_NONE, 0, 0)
		p = pp.Appendpp(pstate.gc, p, x86.ASTOSL, obj.TYPE_NONE, 0, 0, obj.TYPE_NONE, 0, 0)
	}

	return p
}

func (pstate *PackageState) zeroAuto(pp *gc.Progs, n *gc.Node) {
	// Note: this code must not clobber any registers.
	sym := n.Sym.Linksym(pstate.types)
	size := n.Type.Size(pstate.types)
	for i := int64(0); i < size; i += 4 {
		p := pp.Prog(pstate.gc, x86.AMOVL)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 0
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_AUTO
		p.To.Reg = x86.REG_SP
		p.To.Offset = n.Xoffset + i
		p.To.Sym = sym
	}
}

func (pstate *PackageState) ginsnop(pp *gc.Progs) {
	p := pp.Prog(pstate.gc, x86.AXCHGL)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = x86.REG_AX
	p.To.Type = obj.TYPE_REG
	p.To.Reg = x86.REG_AX
}

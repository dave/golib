// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arm

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/arm"
)

func (pstate *PackageState) zerorange(pp *gc.Progs, p *obj.Prog, off, cnt int64, r0 *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	if *r0 == 0 {
		p = pp.Appendpp(pstate.gc, p, arm.AMOVW, obj.TYPE_CONST, 0, 0, obj.TYPE_REG, arm.REG_R0, 0)
		*r0 = 1
	}

	if cnt < int64(4*pstate.gc.Widthptr) {
		for i := int64(0); i < cnt; i += int64(pstate.gc.Widthptr) {
			p = pp.Appendpp(pstate.gc, p, arm.AMOVW, obj.TYPE_REG, arm.REG_R0, 0, obj.TYPE_MEM, arm.REGSP, 4+off+i)
		}
	} else if !pstate.gc.Nacl && (cnt <= int64(128*pstate.gc.Widthptr)) {
		p = pp.Appendpp(pstate.gc, p, arm.AADD, obj.TYPE_CONST, 0, 4+off, obj.TYPE_REG, arm.REG_R1, 0)
		p.Reg = arm.REGSP
		p = pp.Appendpp(pstate.gc, p, obj.ADUFFZERO, obj.TYPE_NONE, 0, 0, obj.TYPE_MEM, 0, 0)
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = pstate.gc.Duffzero
		p.To.Offset = 4 * (128 - cnt/int64(pstate.gc.Widthptr))
	} else {
		p = pp.Appendpp(pstate.gc, p, arm.AADD, obj.TYPE_CONST, 0, 4+off, obj.TYPE_REG, arm.REG_R1, 0)
		p.Reg = arm.REGSP
		p = pp.Appendpp(pstate.gc, p, arm.AADD, obj.TYPE_CONST, 0, cnt, obj.TYPE_REG, arm.REG_R2, 0)
		p.Reg = arm.REG_R1
		p = pp.Appendpp(pstate.gc, p, arm.AMOVW, obj.TYPE_REG, arm.REG_R0, 0, obj.TYPE_MEM, arm.REG_R1, 4)
		p1 := p
		p.Scond |= arm.C_PBIT
		p = pp.Appendpp(pstate.gc, p, arm.ACMP, obj.TYPE_REG, arm.REG_R1, 0, obj.TYPE_NONE, 0, 0)
		p.Reg = arm.REG_R2
		p = pp.Appendpp(pstate.gc, p, arm.ABNE, obj.TYPE_NONE, 0, 0, obj.TYPE_BRANCH, 0, 0)
		pstate.gc.Patch(p, p1)
	}

	return p
}

func (pstate *PackageState) zeroAuto(pp *gc.Progs, n *gc.Node) {
	// Note: this code must not clobber any registers.
	sym := n.Sym.Linksym(pstate.types)
	size := n.Type.Size(pstate.types)
	p := pp.Prog(pstate.gc, arm.AMOVW)
	p.From.Type = obj.TYPE_CONST
	p.From.Offset = 0
	p.To.Type = obj.TYPE_REG
	p.To.Reg = arm.REGTMP
	for i := int64(0); i < size; i += 4 {
		p := pp.Prog(pstate.gc, arm.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = arm.REGTMP
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_AUTO
		p.To.Reg = arm.REGSP
		p.To.Offset = n.Xoffset + i
		p.To.Sym = sym
	}
}

func (pstate *PackageState) ginsnop(pp *gc.Progs) {
	p := pp.Prog(pstate.gc, arm.AAND)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = arm.REG_R0
	p.To.Type = obj.TYPE_REG
	p.To.Reg = arm.REG_R0
	p.Scond = arm.C_SCOND_EQ
}

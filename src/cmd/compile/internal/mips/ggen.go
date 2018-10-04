// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mips

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/mips"
)

// TODO(mips): implement DUFFZERO
func (pstate *PackageState) zerorange(pp *gc.Progs, p *obj.Prog, off, cnt int64, _ *uint32) *obj.Prog {

	if cnt == 0 {
		return p
	}
	if cnt < int64(4*pstate.gc.Widthptr) {
		for i := int64(0); i < cnt; i += int64(pstate.gc.Widthptr) {
			p = pp.Appendpp(pstate.gc, p, mips.AMOVW, obj.TYPE_REG, mips.REGZERO, 0, obj.TYPE_MEM, mips.REGSP, pstate.gc.Ctxt.FixedFrameSize()+off+i)
		}
	} else {
		//fmt.Printf("zerorange frame:%v, lo: %v, hi:%v \n", frame ,lo, hi)
		//	ADD 	$(FIXED_FRAME+frame+lo-4), SP, r1
		//	ADD 	$cnt, r1, r2
		// loop:
		//	MOVW	R0, (Widthptr)r1
		//	ADD 	$Widthptr, r1
		//	BNE		r1, r2, loop
		p = pp.Appendpp(pstate.gc, p, mips.AADD, obj.TYPE_CONST, 0, pstate.gc.Ctxt.FixedFrameSize()+off-4, obj.TYPE_REG, mips.REGRT1, 0)
		p.Reg = mips.REGSP
		p = pp.Appendpp(pstate.gc, p, mips.AADD, obj.TYPE_CONST, 0, cnt, obj.TYPE_REG, mips.REGRT2, 0)
		p.Reg = mips.REGRT1
		p = pp.Appendpp(pstate.gc, p, mips.AMOVW, obj.TYPE_REG, mips.REGZERO, 0, obj.TYPE_MEM, mips.REGRT1, int64(pstate.gc.Widthptr))
		p1 := p
		p = pp.Appendpp(pstate.gc, p, mips.AADD, obj.TYPE_CONST, 0, int64(pstate.gc.Widthptr), obj.TYPE_REG, mips.REGRT1, 0)
		p = pp.Appendpp(pstate.gc, p, mips.ABNE, obj.TYPE_REG, mips.REGRT1, 0, obj.TYPE_BRANCH, 0, 0)
		p.Reg = mips.REGRT2
		pstate.gc.Patch(p, p1)
	}

	return p
}

func (pstate *PackageState) zeroAuto(pp *gc.Progs, n *gc.Node) {
	// Note: this code must not clobber any registers.
	sym := n.Sym.Linksym(pstate.types)
	size := n.Type.Size(pstate.types)
	for i := int64(0); i < size; i += 4 {
		p := pp.Prog(pstate.gc, mips.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = mips.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_AUTO
		p.To.Reg = mips.REGSP
		p.To.Offset = n.Xoffset + i
		p.To.Sym = sym
	}
}

func (pstate *PackageState) ginsnop(pp *gc.Progs) {
	p := pp.Prog(pstate.gc, mips.ANOR)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = mips.REG_R0
	p.To.Type = obj.TYPE_REG
	p.To.Reg = mips.REG_R0
}

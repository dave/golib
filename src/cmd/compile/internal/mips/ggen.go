package mips

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/mips"
)

// TODO(mips): implement DUFFZERO
func (psess *PackageSession) zerorange(pp *gc.Progs, p *obj.Prog, off, cnt int64, _ *uint32) *obj.Prog {

	if cnt == 0 {
		return p
	}
	if cnt < int64(4*psess.gc.Widthptr) {
		for i := int64(0); i < cnt; i += int64(psess.gc.Widthptr) {
			p = pp.Appendpp(psess.gc, p, mips.AMOVW, obj.TYPE_REG, mips.REGZERO, 0, obj.TYPE_MEM, mips.REGSP, psess.gc.Ctxt.FixedFrameSize()+off+i)
		}
	} else {

		p = pp.Appendpp(psess.gc, p, mips.AADD, obj.TYPE_CONST, 0, psess.gc.Ctxt.FixedFrameSize()+off-4, obj.TYPE_REG, mips.REGRT1, 0)
		p.Reg = mips.REGSP
		p = pp.Appendpp(psess.gc, p, mips.AADD, obj.TYPE_CONST, 0, cnt, obj.TYPE_REG, mips.REGRT2, 0)
		p.Reg = mips.REGRT1
		p = pp.Appendpp(psess.gc, p, mips.AMOVW, obj.TYPE_REG, mips.REGZERO, 0, obj.TYPE_MEM, mips.REGRT1, int64(psess.gc.Widthptr))
		p1 := p
		p = pp.Appendpp(psess.gc, p, mips.AADD, obj.TYPE_CONST, 0, int64(psess.gc.Widthptr), obj.TYPE_REG, mips.REGRT1, 0)
		p = pp.Appendpp(psess.gc, p, mips.ABNE, obj.TYPE_REG, mips.REGRT1, 0, obj.TYPE_BRANCH, 0, 0)
		p.Reg = mips.REGRT2
		psess.gc.
			Patch(p, p1)
	}

	return p
}

func (psess *PackageSession) zeroAuto(pp *gc.Progs, n *gc.Node) {

	sym := n.Sym.Linksym(psess.types)
	size := n.Type.Size(psess.types)
	for i := int64(0); i < size; i += 4 {
		p := pp.Prog(psess.gc, mips.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = mips.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_AUTO
		p.To.Reg = mips.REGSP
		p.To.Offset = n.Xoffset + i
		p.To.Sym = sym
	}
}

func (psess *PackageSession) ginsnop(pp *gc.Progs) {
	p := pp.Prog(psess.gc, mips.ANOR)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = mips.REG_R0
	p.To.Type = obj.TYPE_REG
	p.To.Reg = mips.REG_R0
}

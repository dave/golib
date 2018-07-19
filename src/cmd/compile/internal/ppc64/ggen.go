package ppc64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/ppc64"
)

func (psess *PackageSession) zerorange(pp *gc.Progs, p *obj.Prog, off, cnt int64, _ *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	if cnt < int64(4*psess.gc.Widthptr) {
		for i := int64(0); i < cnt; i += int64(psess.gc.Widthptr) {
			p = pp.Appendpp(psess.gc, p, ppc64.AMOVD, obj.TYPE_REG, ppc64.REGZERO, 0, obj.TYPE_MEM, ppc64.REGSP, psess.gc.Ctxt.FixedFrameSize()+off+i)
		}
	} else if cnt <= int64(128*psess.gc.Widthptr) {
		p = pp.Appendpp(psess.gc, p, ppc64.AADD, obj.TYPE_CONST, 0, psess.gc.Ctxt.FixedFrameSize()+off-8, obj.TYPE_REG, ppc64.REGRT1, 0)
		p.Reg = ppc64.REGSP
		p = pp.Appendpp(psess.gc, p, obj.ADUFFZERO, obj.TYPE_NONE, 0, 0, obj.TYPE_MEM, 0, 0)
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = psess.gc.Duffzero
		p.To.Offset = 4 * (128 - cnt/int64(psess.gc.Widthptr))
	} else {
		p = pp.Appendpp(psess.gc, p, ppc64.AMOVD, obj.TYPE_CONST, 0, psess.gc.Ctxt.FixedFrameSize()+off-8, obj.TYPE_REG, ppc64.REGTMP, 0)
		p = pp.Appendpp(psess.gc, p, ppc64.AADD, obj.TYPE_REG, ppc64.REGTMP, 0, obj.TYPE_REG, ppc64.REGRT1, 0)
		p.Reg = ppc64.REGSP
		p = pp.Appendpp(psess.gc, p, ppc64.AMOVD, obj.TYPE_CONST, 0, cnt, obj.TYPE_REG, ppc64.REGTMP, 0)
		p = pp.Appendpp(psess.gc, p, ppc64.AADD, obj.TYPE_REG, ppc64.REGTMP, 0, obj.TYPE_REG, ppc64.REGRT2, 0)
		p.Reg = ppc64.REGRT1
		p = pp.Appendpp(psess.gc, p, ppc64.AMOVDU, obj.TYPE_REG, ppc64.REGZERO, 0, obj.TYPE_MEM, ppc64.REGRT1, int64(psess.gc.Widthptr))
		p1 := p
		p = pp.Appendpp(psess.gc, p, ppc64.ACMP, obj.TYPE_REG, ppc64.REGRT1, 0, obj.TYPE_REG, ppc64.REGRT2, 0)
		p = pp.Appendpp(psess.gc, p, ppc64.ABNE, obj.TYPE_NONE, 0, 0, obj.TYPE_BRANCH, 0, 0)
		psess.gc.
			Patch(p, p1)
	}

	return p
}

func (psess *PackageSession) zeroAuto(pp *gc.Progs, n *gc.Node) {

	sym := n.Sym.Linksym(psess.types)
	size := n.Type.Size(psess.types)
	for i := int64(0); i < size; i += 8 {
		p := pp.Prog(psess.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = ppc64.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_AUTO
		p.To.Reg = ppc64.REGSP
		p.To.Offset = n.Xoffset + i
		p.To.Sym = sym
	}
}

func (psess *PackageSession) ginsnop(pp *gc.Progs) {
	p := pp.Prog(psess.gc, ppc64.AOR)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = ppc64.REG_R0
	p.To.Type = obj.TYPE_REG
	p.To.Reg = ppc64.REG_R0
}

func (psess *PackageSession) ginsnop2(pp *gc.Progs) {
	psess.
		ginsnop(pp)
	if psess.gc.Ctxt.Flag_shared {
		p := pp.Prog(psess.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_MEM
		p.From.Offset = 24
		p.From.Reg = ppc64.REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REG_R2
	} else {
		psess.
			ginsnop(pp)
	}
}

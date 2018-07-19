package arm64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/arm64"
)

func padframe(frame int64) int64 {

	if frame != 0 && frame%16 != 8 {
		frame += 8
	}
	return frame
}

func (psess *PackageSession) zerorange(pp *gc.Progs, p *obj.Prog, off, cnt int64, _ *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	if cnt < int64(4*psess.gc.Widthptr) {
		for i := int64(0); i < cnt; i += int64(psess.gc.Widthptr) {
			p = pp.Appendpp(psess.gc, p, arm64.AMOVD, obj.TYPE_REG, arm64.REGZERO, 0, obj.TYPE_MEM, arm64.REGSP, 8+off+i)
		}
	} else if cnt <= int64(128*psess.gc.Widthptr) && !psess.darwin {
		if cnt%(2*int64(psess.gc.Widthptr)) != 0 {
			p = pp.Appendpp(psess.gc, p, arm64.AMOVD, obj.TYPE_REG, arm64.REGZERO, 0, obj.TYPE_MEM, arm64.REGSP, 8+off)
			off += int64(psess.gc.Widthptr)
			cnt -= int64(psess.gc.Widthptr)
		}
		p = pp.Appendpp(psess.gc, p, arm64.AMOVD, obj.TYPE_REG, arm64.REGSP, 0, obj.TYPE_REG, arm64.REGRT1, 0)
		p = pp.Appendpp(psess.gc, p, arm64.AADD, obj.TYPE_CONST, 0, 8+off, obj.TYPE_REG, arm64.REGRT1, 0)
		p.Reg = arm64.REGRT1
		p = pp.Appendpp(psess.gc, p, obj.ADUFFZERO, obj.TYPE_NONE, 0, 0, obj.TYPE_MEM, 0, 0)
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = psess.gc.Duffzero
		p.To.Offset = 4 * (64 - cnt/(2*int64(psess.gc.Widthptr)))
	} else {
		p = pp.Appendpp(psess.gc, p, arm64.AMOVD, obj.TYPE_CONST, 0, 8+off-8, obj.TYPE_REG, arm64.REGTMP, 0)
		p = pp.Appendpp(psess.gc, p, arm64.AMOVD, obj.TYPE_REG, arm64.REGSP, 0, obj.TYPE_REG, arm64.REGRT1, 0)
		p = pp.Appendpp(psess.gc, p, arm64.AADD, obj.TYPE_REG, arm64.REGTMP, 0, obj.TYPE_REG, arm64.REGRT1, 0)
		p.Reg = arm64.REGRT1
		p = pp.Appendpp(psess.gc, p, arm64.AMOVD, obj.TYPE_CONST, 0, cnt, obj.TYPE_REG, arm64.REGTMP, 0)
		p = pp.Appendpp(psess.gc, p, arm64.AADD, obj.TYPE_REG, arm64.REGTMP, 0, obj.TYPE_REG, arm64.REGRT2, 0)
		p.Reg = arm64.REGRT1
		p = pp.Appendpp(psess.gc, p, arm64.AMOVD, obj.TYPE_REG, arm64.REGZERO, 0, obj.TYPE_MEM, arm64.REGRT1, int64(psess.gc.Widthptr))
		p.Scond = arm64.C_XPRE
		p1 := p
		p = pp.Appendpp(psess.gc, p, arm64.ACMP, obj.TYPE_REG, arm64.REGRT1, 0, obj.TYPE_NONE, 0, 0)
		p.Reg = arm64.REGRT2
		p = pp.Appendpp(psess.gc, p, arm64.ABNE, obj.TYPE_NONE, 0, 0, obj.TYPE_BRANCH, 0, 0)
		psess.gc.
			Patch(p, p1)
	}

	return p
}

func (psess *PackageSession) zeroAuto(pp *gc.Progs, n *gc.Node) {

	sym := n.Sym.Linksym(psess.types)
	size := n.Type.Size(psess.types)
	for i := int64(0); i < size; i += 8 {
		p := pp.Prog(psess.gc, arm64.AMOVD)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = arm64.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_AUTO
		p.To.Reg = arm64.REGSP
		p.To.Offset = n.Xoffset + i
		p.To.Sym = sym
	}
}

func (psess *PackageSession) ginsnop(pp *gc.Progs) {
	p := pp.Prog(psess.gc, arm64.AHINT)
	p.From.Type = obj.TYPE_CONST
}

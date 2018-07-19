package arm

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/arm"
)

func (psess *PackageSession) zerorange(pp *gc.Progs, p *obj.Prog, off, cnt int64, r0 *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	if *r0 == 0 {
		p = pp.Appendpp(psess.gc, p, arm.AMOVW, obj.TYPE_CONST, 0, 0, obj.TYPE_REG, arm.REG_R0, 0)
		*r0 = 1
	}

	if cnt < int64(4*psess.gc.Widthptr) {
		for i := int64(0); i < cnt; i += int64(psess.gc.Widthptr) {
			p = pp.Appendpp(psess.gc, p, arm.AMOVW, obj.TYPE_REG, arm.REG_R0, 0, obj.TYPE_MEM, arm.REGSP, 4+off+i)
		}
	} else if !psess.gc.Nacl && (cnt <= int64(128*psess.gc.Widthptr)) {
		p = pp.Appendpp(psess.gc, p, arm.AADD, obj.TYPE_CONST, 0, 4+off, obj.TYPE_REG, arm.REG_R1, 0)
		p.Reg = arm.REGSP
		p = pp.Appendpp(psess.gc, p, obj.ADUFFZERO, obj.TYPE_NONE, 0, 0, obj.TYPE_MEM, 0, 0)
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = psess.gc.Duffzero
		p.To.Offset = 4 * (128 - cnt/int64(psess.gc.Widthptr))
	} else {
		p = pp.Appendpp(psess.gc, p, arm.AADD, obj.TYPE_CONST, 0, 4+off, obj.TYPE_REG, arm.REG_R1, 0)
		p.Reg = arm.REGSP
		p = pp.Appendpp(psess.gc, p, arm.AADD, obj.TYPE_CONST, 0, cnt, obj.TYPE_REG, arm.REG_R2, 0)
		p.Reg = arm.REG_R1
		p = pp.Appendpp(psess.gc, p, arm.AMOVW, obj.TYPE_REG, arm.REG_R0, 0, obj.TYPE_MEM, arm.REG_R1, 4)
		p1 := p
		p.Scond |= arm.C_PBIT
		p = pp.Appendpp(psess.gc, p, arm.ACMP, obj.TYPE_REG, arm.REG_R1, 0, obj.TYPE_NONE, 0, 0)
		p.Reg = arm.REG_R2
		p = pp.Appendpp(psess.gc, p, arm.ABNE, obj.TYPE_NONE, 0, 0, obj.TYPE_BRANCH, 0, 0)
		psess.gc.
			Patch(p, p1)
	}

	return p
}

func (psess *PackageSession) zeroAuto(pp *gc.Progs, n *gc.Node) {

	sym := n.Sym.Linksym(psess.types)
	size := n.Type.Size(psess.types)
	p := pp.Prog(psess.gc, arm.AMOVW)
	p.From.Type = obj.TYPE_CONST
	p.From.Offset = 0
	p.To.Type = obj.TYPE_REG
	p.To.Reg = arm.REGTMP
	for i := int64(0); i < size; i += 4 {
		p := pp.Prog(psess.gc, arm.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = arm.REGTMP
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_AUTO
		p.To.Reg = arm.REGSP
		p.To.Offset = n.Xoffset + i
		p.To.Sym = sym
	}
}

func (psess *PackageSession) ginsnop(pp *gc.Progs) {
	p := pp.Prog(psess.gc, arm.AAND)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = arm.REG_R0
	p.To.Type = obj.TYPE_REG
	p.To.Reg = arm.REG_R0
	p.Scond = arm.C_SCOND_EQ
}

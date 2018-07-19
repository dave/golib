package s390x

import (
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"

	"math"
)

func progedit(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {
	p.From.Class = 0
	p.To.Class = 0

	c := ctxtz{ctxt: ctxt, newprog: newprog}

	switch p.As {
	case ABR, ABL, obj.ARET, obj.ADUFFZERO, obj.ADUFFCOPY:
		if p.To.Sym != nil {
			p.To.Type = obj.TYPE_BRANCH
		}
	}

	switch p.As {
	case AFMOVS:
		if p.From.Type == obj.TYPE_FCONST {
			f32 := float32(p.From.Val.(float64))
			if math.Float32bits(f32) == 0 {
				break
			}
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = ctxt.Float32Sym(f32)
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}

	case AFMOVD:
		if p.From.Type == obj.TYPE_FCONST {
			f64 := p.From.Val.(float64)
			if math.Float64bits(f64) == 0 {
				break
			}
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = ctxt.Float64Sym(f64)
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}

	case AMOVD:
		if p.From.Type == obj.TYPE_CONST {
			val := p.From.Offset
			if int64(int32(val)) != val &&
				int64(uint32(val)) != val &&
				int64(uint64(val)&(0xffffffff<<32)) != val {
				p.From.Type = obj.TYPE_MEM
				p.From.Sym = ctxt.Int64Sym(p.From.Offset)
				p.From.Name = obj.NAME_EXTERN
				p.From.Offset = 0
			}
		}
	}

	switch p.As {
	case ASUBC:
		if p.From.Type == obj.TYPE_CONST && isint32(-p.From.Offset) {
			p.From.Offset = -p.From.Offset
			p.As = AADDC
		}

	case ASUB:
		if p.From.Type == obj.TYPE_CONST && isint32(-p.From.Offset) {
			p.From.Offset = -p.From.Offset
			p.As = AADD
		}
	}

	if c.ctxt.Flag_dynlink {
		c.rewriteToUseGot(p)
	}
}

// Rewrite p, if necessary, to access global data via the global offset table.
func (c *ctxtz) rewriteToUseGot(p *obj.Prog) {

	if p.As == AEXRL {
		return
	}

	if p.From.Type == obj.TYPE_ADDR && p.From.Name == obj.NAME_EXTERN && !p.From.Sym.Local() {

		if p.To.Type != obj.TYPE_REG || p.As != AMOVD {
			c.ctxt.Diag("do not know how to handle LEA-type insn to non-register in %v with -dynlink", p)
		}
		p.From.Type = obj.TYPE_MEM
		p.From.Name = obj.NAME_GOTREF
		q := p
		if p.From.Offset != 0 {
			target := p.To.Reg
			if target == REG_R0 {

				p.To.Reg = REGTMP2
			}
			q = obj.Appendp(q, c.newprog)
			q.As = AMOVD
			q.From.Type = obj.TYPE_ADDR
			q.From.Offset = p.From.Offset
			q.From.Reg = p.To.Reg
			q.To.Type = obj.TYPE_REG
			q.To.Reg = target
			p.From.Offset = 0
		}
	}
	if p.GetFrom3() != nil && p.GetFrom3().Name == obj.NAME_EXTERN {
		c.ctxt.Diag("don't know how to handle %v with -dynlink", p)
	}
	var source *obj.Addr

	if p.From.Name == obj.NAME_EXTERN && !p.From.Sym.Local() {
		if p.To.Name == obj.NAME_EXTERN && !p.To.Sym.Local() {
			c.ctxt.Diag("cannot handle NAME_EXTERN on both sides in %v with -dynlink", p)
		}
		source = &p.From
	} else if p.To.Name == obj.NAME_EXTERN && !p.To.Sym.Local() {
		source = &p.To
	} else {
		return
	}
	if p.As == obj.ATEXT || p.As == obj.AFUNCDATA || p.As == obj.ACALL || p.As == obj.ARET || p.As == obj.AJMP {
		return
	}
	if source.Sym.Type == objabi.STLSBSS {
		return
	}
	if source.Type != obj.TYPE_MEM {
		c.ctxt.Diag("don't know how to handle %v with -dynlink", p)
	}
	p1 := obj.Appendp(p, c.newprog)
	p2 := obj.Appendp(p1, c.newprog)

	p1.As = AMOVD
	p1.From.Type = obj.TYPE_MEM
	p1.From.Sym = source.Sym
	p1.From.Name = obj.NAME_GOTREF
	p1.To.Type = obj.TYPE_REG
	p1.To.Reg = REGTMP2

	p2.As = p.As
	p2.From = p.From
	p2.To = p.To
	if p.From.Name == obj.NAME_EXTERN {
		p2.From.Reg = REGTMP2
		p2.From.Name = obj.NAME_NONE
		p2.From.Sym = nil
	} else if p.To.Name == obj.NAME_EXTERN {
		p2.To.Reg = REGTMP2
		p2.To.Name = obj.NAME_NONE
		p2.To.Sym = nil
	} else {
		return
	}
	obj.Nopout(p)
}

func preprocess(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {

	if cursym.Func.Text == nil || cursym.Func.Text.Link == nil {
		return
	}

	c := ctxtz{ctxt: ctxt, cursym: cursym, newprog: newprog}

	p := c.cursym.Func.Text
	textstksiz := p.To.Offset
	if textstksiz == -8 {

		p.From.Sym.Set(obj.AttrNoFrame, true)
		textstksiz = 0
	}
	if textstksiz%8 != 0 {
		c.ctxt.Diag("frame size %d not a multiple of 8", textstksiz)
	}
	if p.From.Sym.NoFrame() {
		if textstksiz != 0 {
			c.ctxt.Diag("NOFRAME functions must have a frame size of 0, not %d", textstksiz)
		}
	}

	c.cursym.Func.Args = p.To.Val.(int32)
	c.cursym.Func.Locals = int32(textstksiz)

	var q *obj.Prog
	for p := c.cursym.Func.Text; p != nil; p = p.Link {
		switch p.As {
		case obj.ATEXT:
			q = p
			p.Mark |= LEAF

		case ABL, ABCL:
			q = p
			c.cursym.Func.Text.Mark &^= LEAF
			fallthrough

		case ABC,
			ABEQ,
			ABGE,
			ABGT,
			ABLE,
			ABLT,
			ABLEU,
			ABLTU,
			ABNE,
			ABR,
			ABVC,
			ABVS,
			ACMPBEQ,
			ACMPBGE,
			ACMPBGT,
			ACMPBLE,
			ACMPBLT,
			ACMPBNE,
			ACMPUBEQ,
			ACMPUBGE,
			ACMPUBGT,
			ACMPUBLE,
			ACMPUBLT,
			ACMPUBNE:
			q = p
			p.Mark |= BRANCH
			if p.Pcond != nil {
				q := p.Pcond
				for q.As == obj.ANOP {
					q = q.Link
					p.Pcond = q
				}
			}

		case obj.ANOP:
			q.Link = p.Link
			p.Link.Mark |= p.Mark

		default:
			q = p
		}
	}

	autosize := int32(0)
	var pLast *obj.Prog
	var pPre *obj.Prog
	var pPreempt *obj.Prog
	wasSplit := false
	for p := c.cursym.Func.Text; p != nil; p = p.Link {
		pLast = p
		switch p.As {
		case obj.ATEXT:
			autosize = int32(textstksiz)

			if p.Mark&LEAF != 0 && autosize == 0 {

				p.From.Sym.Set(obj.AttrNoFrame, true)
			}

			if !p.From.Sym.NoFrame() {

				autosize += int32(c.ctxt.FixedFrameSize())
			}

			if p.Mark&LEAF != 0 && autosize < objabi.StackSmall {

				p.From.Sym.Set(obj.AttrNoSplit, true)
			}

			p.To.Offset = int64(autosize)

			q := p

			if !p.From.Sym.NoSplit() {
				p, pPreempt = c.stacksplitPre(p, autosize)
				pPre = p
				wasSplit = true
			}

			if autosize != 0 {

				q = obj.Appendp(p, c.newprog)
				q.As = AMOVD
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_LR
				q.To.Type = obj.TYPE_MEM
				q.To.Reg = REGSP
				q.To.Offset = int64(-autosize)

				q = obj.Appendp(q, c.newprog)
				q.As = AMOVD
				q.From.Type = obj.TYPE_ADDR
				q.From.Offset = int64(-autosize)
				q.From.Reg = REGSP
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REGSP
				q.Spadj = autosize
			} else if c.cursym.Func.Text.Mark&LEAF == 0 {

				c.cursym.Func.Text.Mark |= LEAF
			}

			if c.cursym.Func.Text.Mark&LEAF != 0 {
				c.cursym.Set(obj.AttrLeaf, true)
				break
			}

			if c.cursym.Func.Text.From.Sym.Wrapper() {

				q = obj.Appendp(q, c.newprog)

				q.As = AMOVD
				q.From.Type = obj.TYPE_MEM
				q.From.Reg = REGG
				q.From.Offset = 4 * int64(c.ctxt.Arch.PtrSize)
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R3

				q = obj.Appendp(q, c.newprog)
				q.As = ACMP
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R3
				q.To.Type = obj.TYPE_CONST
				q.To.Offset = 0

				q = obj.Appendp(q, c.newprog)
				q.As = ABEQ
				q.To.Type = obj.TYPE_BRANCH
				p1 := q

				q = obj.Appendp(q, c.newprog)
				q.As = AMOVD
				q.From.Type = obj.TYPE_MEM
				q.From.Reg = REG_R3
				q.From.Offset = 0
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R4

				q = obj.Appendp(q, c.newprog)
				q.As = AADD
				q.From.Type = obj.TYPE_CONST
				q.From.Offset = int64(autosize) + c.ctxt.FixedFrameSize()
				q.Reg = REGSP
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R5

				q = obj.Appendp(q, c.newprog)
				q.As = ACMP
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R4
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R5

				q = obj.Appendp(q, c.newprog)
				q.As = ABNE
				q.To.Type = obj.TYPE_BRANCH
				p2 := q

				q = obj.Appendp(q, c.newprog)
				q.As = AADD
				q.From.Type = obj.TYPE_CONST
				q.From.Offset = c.ctxt.FixedFrameSize()
				q.Reg = REGSP
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R6

				q = obj.Appendp(q, c.newprog)
				q.As = AMOVD
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R6
				q.To.Type = obj.TYPE_MEM
				q.To.Reg = REG_R3
				q.To.Offset = 0

				q = obj.Appendp(q, c.newprog)

				q.As = obj.ANOP
				p1.Pcond = q
				p2.Pcond = q
			}

		case obj.ARET:
			retTarget := p.To.Sym

			if c.cursym.Func.Text.Mark&LEAF != 0 {
				if autosize == 0 {
					p.As = ABR
					p.From = obj.Addr{}
					if retTarget == nil {
						p.To.Type = obj.TYPE_REG
						p.To.Reg = REG_LR
					} else {
						p.To.Type = obj.TYPE_BRANCH
						p.To.Sym = retTarget
					}
					p.Mark |= BRANCH
					break
				}

				p.As = AADD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = int64(autosize)
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REGSP
				p.Spadj = -autosize

				q = obj.Appendp(p, c.newprog)
				q.As = ABR
				q.From = obj.Addr{}
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_LR
				q.Mark |= BRANCH
				q.Spadj = autosize
				break
			}

			p.As = AMOVD
			p.From.Type = obj.TYPE_MEM
			p.From.Reg = REGSP
			p.From.Offset = 0
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REG_LR

			q = p

			if autosize != 0 {
				q = obj.Appendp(q, c.newprog)
				q.As = AADD
				q.From.Type = obj.TYPE_CONST
				q.From.Offset = int64(autosize)
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REGSP
				q.Spadj = -autosize
			}

			q = obj.Appendp(q, c.newprog)
			q.As = ABR
			q.From = obj.Addr{}
			if retTarget == nil {
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_LR
			} else {
				q.To.Type = obj.TYPE_BRANCH
				q.To.Sym = retTarget
			}
			q.Mark |= BRANCH
			q.Spadj = autosize

		case AADD:
			if p.To.Type == obj.TYPE_REG && p.To.Reg == REGSP && p.From.Type == obj.TYPE_CONST {
				p.Spadj = int32(-p.From.Offset)
			}

		case obj.AGETCALLERPC:
			if cursym.Leaf() {

				p.As = AMOVD
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_LR
			} else {

				p.As = AMOVD
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REGSP
			}
		}
	}
	if wasSplit {
		c.stacksplitPost(pLast, pPre, pPreempt, autosize)
	}
}

func (c *ctxtz) stacksplitPre(p *obj.Prog, framesize int32) (*obj.Prog, *obj.Prog) {
	var q *obj.Prog

	p = obj.Appendp(p, c.newprog)

	p.As = AMOVD
	p.From.Type = obj.TYPE_MEM
	p.From.Reg = REGG
	p.From.Offset = 2 * int64(c.ctxt.Arch.PtrSize)
	if c.cursym.CFunc() {
		p.From.Offset = 3 * int64(c.ctxt.Arch.PtrSize)
	}
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_R3

	q = nil
	if framesize <= objabi.StackSmall {

		p = obj.Appendp(p, c.newprog)

		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R3
		p.Reg = REGSP
		p.As = ACMPUBGE
		p.To.Type = obj.TYPE_BRANCH

	} else if framesize <= objabi.StackBig {

		p = obj.Appendp(p, c.newprog)

		p.As = AADD
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = -(int64(framesize) - objabi.StackSmall)
		p.Reg = REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R4

		p = obj.Appendp(p, c.newprog)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R3
		p.Reg = REG_R4
		p.As = ACMPUBGE
		p.To.Type = obj.TYPE_BRANCH

	} else {

		p = obj.Appendp(p, c.newprog)

		p.As = ACMP
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R3
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = objabi.StackPreempt

		p = obj.Appendp(p, c.newprog)
		q = p
		p.As = ABEQ
		p.To.Type = obj.TYPE_BRANCH

		p = obj.Appendp(p, c.newprog)
		p.As = AADD
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = objabi.StackGuard
		p.Reg = REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R4

		p = obj.Appendp(p, c.newprog)
		p.As = ASUB
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R3
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R4

		p = obj.Appendp(p, c.newprog)
		p.As = AMOVD
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = int64(framesize) + objabi.StackGuard - objabi.StackSmall
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REGTMP

		p = obj.Appendp(p, c.newprog)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REGTMP
		p.Reg = REG_R4
		p.As = ACMPUBGE
		p.To.Type = obj.TYPE_BRANCH
	}

	return p, q
}

func (c *ctxtz) stacksplitPost(p *obj.Prog, pPre *obj.Prog, pPreempt *obj.Prog, framesize int32) *obj.Prog {

	spfix := obj.Appendp(p, c.newprog)
	spfix.As = obj.ANOP
	spfix.Spadj = -framesize

	pcdata := c.ctxt.EmitEntryLiveness(c.cursym, spfix, c.newprog)

	p = obj.Appendp(pcdata, c.newprog)
	pPre.Pcond = p
	p.As = AMOVD
	p.From.Type = obj.TYPE_REG
	p.From.Reg = REG_LR
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_R5
	if pPreempt != nil {
		pPreempt.Pcond = p
	}

	p = obj.Appendp(p, c.newprog)

	p.As = ABL
	p.To.Type = obj.TYPE_BRANCH
	if c.cursym.CFunc() {
		p.To.Sym = c.ctxt.Lookup("runtime.morestackc")
	} else if !c.cursym.Func.Text.From.Sym.NeedCtxt() {
		p.To.Sym = c.ctxt.Lookup("runtime.morestack_noctxt")
	} else {
		p.To.Sym = c.ctxt.Lookup("runtime.morestack")
	}

	p = obj.Appendp(p, c.newprog)

	p.As = ABR
	p.To.Type = obj.TYPE_BRANCH
	p.Pcond = c.cursym.Func.Text.Link
	return p
}

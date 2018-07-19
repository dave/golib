package arm64

import (
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"

	"math"
)

func (c *ctxt7) stacksplit(p *obj.Prog, framesize int32) *obj.Prog {

	p = obj.Appendp(p, c.newprog)

	p.As = AMOVD
	p.From.Type = obj.TYPE_MEM
	p.From.Reg = REGG
	p.From.Offset = 2 * int64(c.ctxt.Arch.PtrSize)
	if c.cursym.CFunc() {
		p.From.Offset = 3 * int64(c.ctxt.Arch.PtrSize)
	}
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_R1

	q := (*obj.Prog)(nil)
	if framesize <= objabi.StackSmall {

		p = obj.Appendp(p, c.newprog)

		p.As = AMOVD
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2

		p = obj.Appendp(p, c.newprog)
		p.As = ACMP
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R1
		p.Reg = REG_R2
	} else if framesize <= objabi.StackBig {

		p = obj.Appendp(p, c.newprog)

		p.As = ASUB
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = int64(framesize) - objabi.StackSmall
		p.Reg = REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2

		p = obj.Appendp(p, c.newprog)
		p.As = ACMP
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R1
		p.Reg = REG_R2
	} else {

		p = obj.Appendp(p, c.newprog)

		p.As = ACMP
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = objabi.StackPreempt
		p.Reg = REG_R1

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
		p.To.Reg = REG_R2

		p = obj.Appendp(p, c.newprog)
		p.As = ASUB
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2

		p = obj.Appendp(p, c.newprog)
		p.As = AMOVD
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = int64(framesize) + (objabi.StackGuard - objabi.StackSmall)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R3

		p = obj.Appendp(p, c.newprog)
		p.As = ACMP
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R3
		p.Reg = REG_R2
	}

	bls := obj.Appendp(p, c.newprog)
	bls.As = ABLS
	bls.To.Type = obj.TYPE_BRANCH

	var last *obj.Prog
	for last = c.cursym.Func.Text; last.Link != nil; last = last.Link {
	}

	spfix := obj.Appendp(last, c.newprog)
	spfix.As = obj.ANOP
	spfix.Spadj = -framesize

	pcdata := c.ctxt.EmitEntryLiveness(c.cursym, spfix, c.newprog)

	movlr := obj.Appendp(pcdata, c.newprog)
	movlr.As = AMOVD
	movlr.From.Type = obj.TYPE_REG
	movlr.From.Reg = REGLINK
	movlr.To.Type = obj.TYPE_REG
	movlr.To.Reg = REG_R3
	if q != nil {
		q.Pcond = movlr
	}
	bls.Pcond = movlr

	debug := movlr
	if false {
		debug = obj.Appendp(debug, c.newprog)
		debug.As = AMOVD
		debug.From.Type = obj.TYPE_CONST
		debug.From.Offset = int64(framesize)
		debug.To.Type = obj.TYPE_REG
		debug.To.Reg = REGTMP
	}

	call := obj.Appendp(debug, c.newprog)
	call.As = ABL
	call.To.Type = obj.TYPE_BRANCH
	morestack := "runtime.morestack"
	switch {
	case c.cursym.CFunc():
		morestack = "runtime.morestackc"
	case !c.cursym.Func.Text.From.Sym.NeedCtxt():
		morestack = "runtime.morestack_noctxt"
	}
	call.To.Sym = c.ctxt.Lookup(morestack)

	jmp := obj.Appendp(call, c.newprog)
	jmp.As = AB
	jmp.To.Type = obj.TYPE_BRANCH
	jmp.Pcond = c.cursym.Func.Text.Link
	jmp.Spadj = +framesize

	return bls
}

func (psess *PackageSession) progedit(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {
	c := ctxt7{ctxt: ctxt, newprog: newprog}

	p.From.Class = 0
	p.To.Class = 0

	if p.From.Type == obj.TYPE_CONST && p.From.Offset == 0 {
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REGZERO
	}
	if p.To.Type == obj.TYPE_CONST && p.To.Offset == 0 {
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REGZERO
	}

	switch p.As {
	case AB,
		ABL,
		obj.ARET,
		obj.ADUFFZERO,
		obj.ADUFFCOPY:
		if p.To.Sym != nil {
			p.To.Type = obj.TYPE_BRANCH
		}
		break
	}

	switch p.As {
	case AFMOVS:
		if p.From.Type == obj.TYPE_FCONST {
			f32 := float32(p.From.Val.(float64))
			if math.Float32bits(f32) == 0 {
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGZERO
				break
			}
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = c.ctxt.Float32Sym(f32)
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}

	case AFMOVD:
		if p.From.Type == obj.TYPE_FCONST {
			f64 := p.From.Val.(float64)
			if math.Float64bits(f64) == 0 {
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGZERO
				break
			}
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = c.ctxt.Float64Sym(f64)
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}

		break
	}

	switch p.As {
	case AADD, ASUB, ACMP, ACMN:
		if p.From.Type == obj.TYPE_CONST && p.From.Offset < 0 && p.From.Offset != -1<<63 {
			p.From.Offset = -p.From.Offset
			p.As = psess.complements[p.As]
		}
	case AADDW, ASUBW, ACMPW, ACMNW:
		if p.From.Type == obj.TYPE_CONST && p.From.Offset < 0 && int32(p.From.Offset) != -1<<31 {
			p.From.Offset = -p.From.Offset
			p.As = psess.complements[p.As]
		}
	}

	switch p.As {
	case AANDW, AORRW, AEORW, AANDSW, ATSTW:
		if p.From.Type == obj.TYPE_CONST {
			v := p.From.Offset & 0xffffffff
			p.From.Offset = v | v<<32
		}
	}

	if c.ctxt.Flag_dynlink {
		c.rewriteToUseGot(p)
	}
}

// Rewrite p, if necessary, to access global data via the global offset table.
func (c *ctxt7) rewriteToUseGot(p *obj.Prog) {
	if p.As == obj.ADUFFCOPY || p.As == obj.ADUFFZERO {
		//     ADUFFxxx $offset
		// becomes
		//     MOVD runtime.duffxxx@GOT, REGTMP
		//     ADD $offset, REGTMP
		//     CALL REGTMP
		var sym *obj.LSym
		if p.As == obj.ADUFFZERO {
			sym = c.ctxt.Lookup("runtime.duffzero")
		} else {
			sym = c.ctxt.Lookup("runtime.duffcopy")
		}
		offset := p.To.Offset
		p.As = AMOVD
		p.From.Type = obj.TYPE_MEM
		p.From.Name = obj.NAME_GOTREF
		p.From.Sym = sym
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REGTMP
		p.To.Name = obj.NAME_NONE
		p.To.Offset = 0
		p.To.Sym = nil
		p1 := obj.Appendp(p, c.newprog)
		p1.As = AADD
		p1.From.Type = obj.TYPE_CONST
		p1.From.Offset = offset
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = REGTMP
		p2 := obj.Appendp(p1, c.newprog)
		p2.As = obj.ACALL
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = REGTMP
	}

	if p.From.Type == obj.TYPE_ADDR && p.From.Name == obj.NAME_EXTERN && !p.From.Sym.Local() {

		if p.As != AMOVD {
			c.ctxt.Diag("do not know how to handle TYPE_ADDR in %v with -dynlink", p)
		}
		if p.To.Type != obj.TYPE_REG {
			c.ctxt.Diag("do not know how to handle LEAQ-type insn to non-register in %v with -dynlink", p)
		}
		p.From.Type = obj.TYPE_MEM
		p.From.Name = obj.NAME_GOTREF
		if p.From.Offset != 0 {
			q := obj.Appendp(p, c.newprog)
			q.As = AADD
			q.From.Type = obj.TYPE_CONST
			q.From.Offset = p.From.Offset
			q.To = p.To
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
	p1.To.Reg = REGTMP

	p2.As = p.As
	p2.From = p.From
	p2.To = p.To
	if p.From.Name == obj.NAME_EXTERN {
		p2.From.Reg = REGTMP
		p2.From.Name = obj.NAME_NONE
		p2.From.Sym = nil
	} else if p.To.Name == obj.NAME_EXTERN {
		p2.To.Reg = REGTMP
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

	c := ctxt7{ctxt: ctxt, newprog: newprog, cursym: cursym}

	p := c.cursym.Func.Text
	textstksiz := p.To.Offset
	if textstksiz == -8 {

		p.From.Sym.Set(obj.AttrNoFrame, true)
		textstksiz = 0
	}
	if textstksiz < 0 {
		c.ctxt.Diag("negative frame size %d - did you mean NOFRAME?", textstksiz)
	}
	if p.From.Sym.NoFrame() {
		if textstksiz != 0 {
			c.ctxt.Diag("NOFRAME functions must have a frame size of 0, not %d", textstksiz)
		}
	}

	c.cursym.Func.Args = p.To.Val.(int32)
	c.cursym.Func.Locals = int32(textstksiz)

	q := (*obj.Prog)(nil)
	var q1 *obj.Prog
	for p := c.cursym.Func.Text; p != nil; p = p.Link {
		switch p.As {
		case obj.ATEXT:
			p.Mark |= LEAF

		case obj.ARET:
			break

		case obj.ANOP:
			if p.Link != nil {
				q1 = p.Link
				q.Link = q1
				q1.Mark |= p.Mark
			}
			continue

		case ABL,
			obj.ADUFFZERO,
			obj.ADUFFCOPY:
			c.cursym.Func.Text.Mark &^= LEAF
			fallthrough

		case ACBNZ,
			ACBZ,
			ACBNZW,
			ACBZW,
			ATBZ,
			ATBNZ,
			AB,
			ABEQ,
			ABNE,
			ABCS,
			ABHS,
			ABCC,
			ABLO,
			ABMI,
			ABPL,
			ABVS,
			ABVC,
			ABHI,
			ABLS,
			ABGE,
			ABLT,
			ABGT,
			ABLE,
			AADR,
			AADRP:
			q1 = p.Pcond

			if q1 != nil {
				for q1.As == obj.ANOP {
					q1 = q1.Link
					p.Pcond = q1
				}
			}

			break
		}

		q = p
	}

	var retjmp *obj.LSym
	for p := c.cursym.Func.Text; p != nil; p = p.Link {
		o := p.As
		switch o {
		case obj.ATEXT:
			c.cursym.Func.Text = p
			c.autosize = int32(textstksiz)

			if p.Mark&LEAF != 0 && c.autosize == 0 {

				p.From.Sym.Set(obj.AttrNoFrame, true)
			}

			if !p.From.Sym.NoFrame() {

				c.autosize += 8
			}

			if c.autosize != 0 && c.autosize&(16-1) != 0 {

				if c.autosize == 8 {
					c.autosize += 8
					c.cursym.Func.Locals += 8
				} else {
					c.ctxt.Diag("%v: unaligned frame size %d - must be 8 mod 16 (or 0)", p, c.autosize-8)
				}
			}
			if c.autosize == 0 && c.cursym.Func.Text.Mark&LEAF == 0 {
				if c.ctxt.Debugvlog {
					c.ctxt.Logf("save suppressed in: %s\n", c.cursym.Func.Text.From.Sym.Name)
				}
				c.cursym.Func.Text.Mark |= LEAF
			}

			p.To.Offset = int64(c.autosize) - 8

			if cursym.Func.Text.Mark&LEAF != 0 {
				cursym.Set(obj.AttrLeaf, true)
				if p.From.Sym.NoFrame() {
					break
				}
			}

			if !p.From.Sym.NoSplit() {
				p = c.stacksplit(p, c.autosize)
			}

			aoffset := c.autosize
			if aoffset > 0xF0 {
				aoffset = 0xF0
			}

			q = p
			if c.autosize > aoffset {

				q = obj.Appendp(q, c.newprog)
				q.Pos = p.Pos
				q.As = ASUB
				q.From.Type = obj.TYPE_CONST
				q.From.Offset = int64(c.autosize)
				q.Reg = REGSP
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REGTMP

				q = obj.Appendp(q, c.newprog)
				q.Pos = p.Pos
				q.As = AMOVD
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REGLINK
				q.To.Type = obj.TYPE_MEM
				q.To.Reg = REGTMP

				q1 = obj.Appendp(q, c.newprog)
				q1.Pos = p.Pos
				q1.As = AMOVD
				q1.From.Type = obj.TYPE_REG
				q1.From.Reg = REGTMP
				q1.To.Type = obj.TYPE_REG
				q1.To.Reg = REGSP
				q1.Spadj = c.autosize
			} else {

				q1 = obj.Appendp(q, c.newprog)
				q1.As = AMOVD
				q1.Pos = p.Pos
				q1.From.Type = obj.TYPE_REG
				q1.From.Reg = REGLINK
				q1.To.Type = obj.TYPE_MEM
				q1.Scond = C_XPRE
				q1.To.Offset = int64(-aoffset)
				q1.To.Reg = REGSP
				q1.Spadj = aoffset
			}

			if c.cursym.Func.Text.From.Sym.Wrapper() {

				q = q1

				q = obj.Appendp(q, c.newprog)
				q.As = AMOVD
				q.From.Type = obj.TYPE_MEM
				q.From.Reg = REGG
				q.From.Offset = 4 * int64(c.ctxt.Arch.PtrSize)
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R1

				cbnz := obj.Appendp(q, c.newprog)
				cbnz.As = ACBNZ
				cbnz.From.Type = obj.TYPE_REG
				cbnz.From.Reg = REG_R1
				cbnz.To.Type = obj.TYPE_BRANCH

				end := obj.Appendp(cbnz, c.newprog)
				end.As = obj.ANOP

				// find the end of the function
				var last *obj.Prog
				for last = end; last.Link != nil; last = last.Link {
				}

				mov := obj.Appendp(last, c.newprog)
				mov.As = AMOVD
				mov.From.Type = obj.TYPE_MEM
				mov.From.Reg = REG_R1
				mov.From.Offset = 0
				mov.To.Type = obj.TYPE_REG
				mov.To.Reg = REG_R2

				cbnz.Pcond = mov

				q = obj.Appendp(mov, c.newprog)
				q.As = AADD
				q.From.Type = obj.TYPE_CONST
				q.From.Offset = int64(c.autosize) + 8
				q.Reg = REGSP
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R3

				q = obj.Appendp(q, c.newprog)
				q.As = ACMP
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R2
				q.Reg = REG_R3

				q = obj.Appendp(q, c.newprog)
				q.As = ABNE
				q.To.Type = obj.TYPE_BRANCH
				q.Pcond = end

				q = obj.Appendp(q, c.newprog)
				q.As = AADD
				q.From.Type = obj.TYPE_CONST
				q.From.Offset = 8
				q.Reg = REGSP
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R4

				q = obj.Appendp(q, c.newprog)
				q.As = AMOVD
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R4
				q.To.Type = obj.TYPE_MEM
				q.To.Reg = REG_R1
				q.To.Offset = 0

				q = obj.Appendp(q, c.newprog)
				q.As = AB
				q.To.Type = obj.TYPE_BRANCH
				q.Pcond = end
			}

		case obj.ARET:
			nocache(p)
			if p.From.Type == obj.TYPE_CONST {
				c.ctxt.Diag("using BECOME (%v) is not supported!", p)
				break
			}

			retjmp = p.To.Sym
			p.To = obj.Addr{}
			if c.cursym.Func.Text.Mark&LEAF != 0 {
				if c.autosize != 0 {
					p.As = AADD
					p.From.Type = obj.TYPE_CONST
					p.From.Offset = int64(c.autosize)
					p.To.Type = obj.TYPE_REG
					p.To.Reg = REGSP
					p.Spadj = -c.autosize
				}
			} else {

				aoffset := c.autosize

				if aoffset > 0xF0 {
					aoffset = 0xF0
				}
				p.As = AMOVD
				p.From.Type = obj.TYPE_MEM
				p.Scond = C_XPOST
				p.From.Offset = int64(aoffset)
				p.From.Reg = REGSP
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REGLINK
				p.Spadj = -aoffset
				if c.autosize > aoffset {
					q = newprog()
					q.As = AADD
					q.From.Type = obj.TYPE_CONST
					q.From.Offset = int64(c.autosize) - int64(aoffset)
					q.To.Type = obj.TYPE_REG
					q.To.Reg = REGSP
					q.Link = p.Link
					q.Spadj = int32(-q.From.Offset)
					q.Pos = p.Pos
					p.Link = q
					p = q
				}
			}

			if p.As != obj.ARET {
				q = newprog()
				q.Pos = p.Pos
				q.Link = p.Link
				p.Link = q
				p = q
			}

			if retjmp != nil {
				p.As = AB
				p.To.Type = obj.TYPE_BRANCH
				p.To.Sym = retjmp
				p.Spadj = +c.autosize
				break
			}

			p.As = obj.ARET
			p.To.Type = obj.TYPE_MEM
			p.To.Offset = 0
			p.To.Reg = REGLINK
			p.Spadj = +c.autosize

		case AADD, ASUB:
			if p.To.Type == obj.TYPE_REG && p.To.Reg == REGSP && p.From.Type == obj.TYPE_CONST {
				if p.As == AADD {
					p.Spadj = int32(-p.From.Offset)
				} else {
					p.Spadj = int32(+p.From.Offset)
				}
			}
			break

		case obj.AGETCALLERPC:
			if cursym.Leaf() {

				p.As = AMOVD
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGLINK
			} else {

				p.As = AMOVD
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REGSP
			}
		}
	}
}

func nocache(p *obj.Prog) {
	p.Optab = 0
	p.From.Class = 0
	p.To.Class = 0
}

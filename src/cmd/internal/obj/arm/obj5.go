package arm

import (
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
)

func (psess *PackageSession) progedit(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {
	p.From.Class = 0
	p.To.Class = 0

	c := ctxt5{ctxt: ctxt, newprog: newprog}

	switch p.As {
	case AB, ABL, obj.ADUFFZERO, obj.ADUFFCOPY:
		if p.To.Type == obj.TYPE_MEM && (p.To.Name == obj.NAME_EXTERN || p.To.Name == obj.NAME_STATIC) && p.To.Sym != nil {
			p.To.Type = obj.TYPE_BRANCH
		}
	}

	switch p.As {

	case AMRC:
		if p.To.Offset&0xffff0fff == 0xee1d0f70 {

			if p.To.Offset&0xf000 != 0 {
				ctxt.Diag("%v: TLS MRC instruction must write to R0 as it might get translated into a BL instruction", p.Line(psess.obj))
			}

			if psess.objabi.GOARM < 7 {

				if psess.progedit_tlsfallback == nil {
					psess.
						progedit_tlsfallback = ctxt.Lookup("runtime.read_tls_fallback")
				}

				p.As = AMOVW

				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGLINK
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REGTMP

				p = obj.Appendp(p, newprog)

				p.As = ABL
				p.To.Type = obj.TYPE_BRANCH
				p.To.Sym = psess.progedit_tlsfallback
				p.To.Offset = 0

				p = obj.Appendp(p, newprog)

				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGTMP
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REGLINK
				break
			}
		}

		p.As = AWORD
	}

	switch p.As {
	case AMOVF:
		if p.From.Type == obj.TYPE_FCONST && c.chipfloat5(psess, p.From.Val.(float64)) < 0 && (c.chipzero5(psess, p.From.Val.(float64)) < 0 || p.Scond&C_SCOND != C_SCOND_NONE) {
			f32 := float32(p.From.Val.(float64))
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = ctxt.Float32Sym(f32)
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}

	case AMOVD:
		if p.From.Type == obj.TYPE_FCONST && c.chipfloat5(psess, p.From.Val.(float64)) < 0 && (c.chipzero5(psess, p.From.Val.(float64)) < 0 || p.Scond&C_SCOND != C_SCOND_NONE) {
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = ctxt.Float64Sym(p.From.Val.(float64))
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}
	}

	if ctxt.Flag_dynlink {
		c.rewriteToUseGot(p)
	}
}

// Rewrite p, if necessary, to access global data via the global offset table.
func (c *ctxt5) rewriteToUseGot(p *obj.Prog) {
	if p.As == obj.ADUFFCOPY || p.As == obj.ADUFFZERO {
		//     ADUFFxxx $offset
		// becomes
		//     MOVW runtime.duffxxx@GOT, R9
		//     ADD $offset, R9
		//     CALL (R9)
		var sym *obj.LSym
		if p.As == obj.ADUFFZERO {
			sym = c.ctxt.Lookup("runtime.duffzero")
		} else {
			sym = c.ctxt.Lookup("runtime.duffcopy")
		}
		offset := p.To.Offset
		p.As = AMOVW
		p.From.Type = obj.TYPE_MEM
		p.From.Name = obj.NAME_GOTREF
		p.From.Sym = sym
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R9
		p.To.Name = obj.NAME_NONE
		p.To.Offset = 0
		p.To.Sym = nil
		p1 := obj.Appendp(p, c.newprog)
		p1.As = AADD
		p1.From.Type = obj.TYPE_CONST
		p1.From.Offset = offset
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = REG_R9
		p2 := obj.Appendp(p1, c.newprog)
		p2.As = obj.ACALL
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = REG_R9
		return
	}

	if p.From.Type == obj.TYPE_ADDR && p.From.Name == obj.NAME_EXTERN && !p.From.Sym.Local() {

		if p.As != AMOVW {
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

	p1.As = AMOVW
	p1.From.Type = obj.TYPE_MEM
	p1.From.Sym = source.Sym
	p1.From.Name = obj.NAME_GOTREF
	p1.To.Type = obj.TYPE_REG
	p1.To.Reg = REG_R9

	p2.As = p.As
	p2.From = p.From
	p2.To = p.To
	if p.From.Name == obj.NAME_EXTERN {
		p2.From.Reg = REG_R9
		p2.From.Name = obj.NAME_NONE
		p2.From.Sym = nil
	} else if p.To.Name == obj.NAME_EXTERN {
		p2.To.Reg = REG_R9
		p2.To.Name = obj.NAME_NONE
		p2.To.Sym = nil
	} else {
		return
	}
	obj.Nopout(p)
}

// Prog.mark
const (
	FOLL  = 1 << 0
	LABEL = 1 << 1
	LEAF  = 1 << 2
)

func (psess *PackageSession) preprocess(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	autosize := int32(0)

	if cursym.Func.Text == nil || cursym.Func.Text.Link == nil {
		return
	}

	c := ctxt5{ctxt: ctxt, cursym: cursym, newprog: newprog}

	p := c.cursym.Func.Text
	autoffset := int32(p.To.Offset)
	if autoffset == -4 {

		p.From.Sym.Set(obj.AttrNoFrame, true)
		autoffset = 0
	}
	if autoffset < 0 || autoffset%4 != 0 {
		c.ctxt.Diag("frame size %d not 0 or a positive multiple of 4", autoffset)
	}
	if p.From.Sym.NoFrame() {
		if autoffset != 0 {
			c.ctxt.Diag("NOFRAME functions must have a frame size of 0, not %d", autoffset)
		}
	}

	cursym.Func.Locals = autoffset
	cursym.Func.Args = p.To.Val.(int32)

	/*
	 * find leaf subroutines
	 * strip NOPs
	 * expand RET
	 * expand BECOME pseudo
	 */
	var q1 *obj.Prog
	var q *obj.Prog
	for p := cursym.Func.Text; p != nil; p = p.Link {
		switch p.As {
		case obj.ATEXT:
			p.Mark |= LEAF

		case obj.ARET:
			break

		case ADIV, ADIVU, AMOD, AMODU:
			q = p
			cursym.Func.Text.Mark &^= LEAF
			continue

		case obj.ANOP:
			q1 = p.Link
			q.Link = q1
			if q1 != nil {
				q1.Mark |= p.Mark
			}
			continue

		case ABL,
			ABX,
			obj.ADUFFZERO,
			obj.ADUFFCOPY:
			cursym.Func.Text.Mark &^= LEAF
			fallthrough

		case AB,
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
			ABLE:
			q1 = p.Pcond
			if q1 != nil {
				for q1.As == obj.ANOP {
					q1 = q1.Link
					p.Pcond = q1
				}
			}
		}

		q = p
	}

	var q2 *obj.Prog
	for p := cursym.Func.Text; p != nil; p = p.Link {
		o := p.As
		switch o {
		case obj.ATEXT:
			autosize = autoffset

			if p.Mark&LEAF != 0 && autosize == 0 {

				p.From.Sym.Set(obj.AttrNoFrame, true)
			}

			if !p.From.Sym.NoFrame() {

				autosize += 4
			}

			if autosize == 0 && cursym.Func.Text.Mark&LEAF == 0 {

				if ctxt.Debugvlog {
					ctxt.Logf("save suppressed in: %s\n", cursym.Name)
				}

				cursym.Func.Text.Mark |= LEAF
			}

			p.To.Offset = int64(autosize) - 4

			if cursym.Func.Text.Mark&LEAF != 0 {
				cursym.Set(obj.AttrLeaf, true)
				if p.From.Sym.NoFrame() {
					break
				}
			}

			if !p.From.Sym.NoSplit() {
				p = c.stacksplit(p, autosize)
			}

			p = obj.Appendp(p, c.newprog)

			p.As = AMOVW
			p.Scond |= C_WBIT
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REGLINK
			p.To.Type = obj.TYPE_MEM
			p.To.Offset = int64(-autosize)
			p.To.Reg = REGSP
			p.Spadj = autosize

			if cursym.Func.Text.From.Sym.Wrapper() {

				p = obj.Appendp(p, newprog)
				p.As = AMOVW
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REGG
				p.From.Offset = 4 * int64(ctxt.Arch.PtrSize)
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_R1

				p = obj.Appendp(p, newprog)
				p.As = ACMP
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = 0
				p.Reg = REG_R1

				bne := obj.Appendp(p, newprog)
				bne.As = ABNE
				bne.To.Type = obj.TYPE_BRANCH

				end := obj.Appendp(bne, newprog)
				end.As = obj.ANOP

				// find end of function
				var last *obj.Prog
				for last = end; last.Link != nil; last = last.Link {
				}

				mov := obj.Appendp(last, newprog)
				mov.As = AMOVW
				mov.From.Type = obj.TYPE_MEM
				mov.From.Reg = REG_R1
				mov.From.Offset = 0
				mov.To.Type = obj.TYPE_REG
				mov.To.Reg = REG_R2

				bne.Pcond = mov

				p = obj.Appendp(mov, newprog)
				p.As = AADD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = int64(autosize) + 4
				p.Reg = REG_R13
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_R3

				p = obj.Appendp(p, newprog)
				p.As = ACMP
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_R2
				p.Reg = REG_R3

				p = obj.Appendp(p, newprog)
				p.As = ABNE
				p.To.Type = obj.TYPE_BRANCH
				p.Pcond = end

				p = obj.Appendp(p, newprog)
				p.As = AADD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = 4
				p.Reg = REG_R13
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_R4

				p = obj.Appendp(p, newprog)
				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_R4
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = REG_R1
				p.To.Offset = 0

				p = obj.Appendp(p, newprog)
				p.As = AB
				p.To.Type = obj.TYPE_BRANCH
				p.Pcond = end

				p = end
			}

		case obj.ARET:
			nocache(p)
			if cursym.Func.Text.Mark&LEAF != 0 {
				if autosize == 0 {
					p.As = AB
					p.From = obj.Addr{}
					if p.To.Sym != nil {
						p.To.Type = obj.TYPE_BRANCH
					} else {
						p.To.Type = obj.TYPE_MEM
						p.To.Offset = 0
						p.To.Reg = REGLINK
					}

					break
				}
			}

			p.As = AMOVW
			p.Scond |= C_PBIT
			p.From.Type = obj.TYPE_MEM
			p.From.Offset = int64(autosize)
			p.From.Reg = REGSP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REGPC

			if p.To.Sym != nil {
				p.To.Reg = REGLINK
				q2 = obj.Appendp(p, newprog)
				q2.As = AB
				q2.To.Type = obj.TYPE_BRANCH
				q2.To.Sym = p.To.Sym
				p.To.Sym = nil
				p = q2
			}

		case AADD:
			if p.From.Type == obj.TYPE_CONST && p.From.Reg == 0 && p.To.Type == obj.TYPE_REG && p.To.Reg == REGSP {
				p.Spadj = int32(-p.From.Offset)
			}

		case ASUB:
			if p.From.Type == obj.TYPE_CONST && p.From.Reg == 0 && p.To.Type == obj.TYPE_REG && p.To.Reg == REGSP {
				p.Spadj = int32(p.From.Offset)
			}

		case ADIV, ADIVU, AMOD, AMODU:
			if cursym.Func.Text.From.Sym.NoSplit() {
				ctxt.Diag("cannot divide in NOSPLIT function")
			}
			const debugdivmod = false
			if debugdivmod {
				break
			}
			if p.From.Type != obj.TYPE_REG {
				break
			}
			if p.To.Type != obj.TYPE_REG {
				break
			}

			q1 := *p
			if q1.Reg == REGTMP || q1.Reg == 0 && q1.To.Reg == REGTMP {
				ctxt.Diag("div already using REGTMP: %v", p)
			}

			p.As = AMOVW
			p.Pos = q1.Pos
			p.From.Type = obj.TYPE_MEM
			p.From.Reg = REGG
			p.From.Offset = 6 * 4
			p.Reg = 0
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REGTMP

			p = obj.Appendp(p, newprog)
			p.As = AMOVW
			p.Pos = q1.Pos
			p.From.Type = obj.TYPE_REG
			p.From.Reg = q1.From.Reg
			p.To.Type = obj.TYPE_MEM
			p.To.Reg = REGTMP
			p.To.Offset = 8 * 4

			p = obj.Appendp(p, newprog)
			p.As = AMOVW
			p.Pos = q1.Pos
			p.From.Type = obj.TYPE_REG
			p.From.Reg = q1.Reg
			if q1.Reg == 0 {
				p.From.Reg = q1.To.Reg
			}
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REG_R8
			p.To.Offset = 0

			p = obj.Appendp(p, newprog)
			p.As = ABL
			p.Pos = q1.Pos
			p.To.Type = obj.TYPE_BRANCH
			switch o {
			case ADIV:
				p.To.Sym = psess.symdiv
			case ADIVU:
				p.To.Sym = psess.symdivu
			case AMOD:
				p.To.Sym = psess.symmod
			case AMODU:
				p.To.Sym = psess.symmodu
			}

			p = obj.Appendp(p, newprog)
			p.As = AMOVW
			p.Pos = q1.Pos
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REGTMP
			p.From.Offset = 0
			p.To.Type = obj.TYPE_REG
			p.To.Reg = q1.To.Reg

		case AMOVW:
			if (p.Scond&C_WBIT != 0) && p.To.Type == obj.TYPE_MEM && p.To.Reg == REGSP {
				p.Spadj = int32(-p.To.Offset)
			}
			if (p.Scond&C_PBIT != 0) && p.From.Type == obj.TYPE_MEM && p.From.Reg == REGSP && p.To.Reg != REGPC {
				p.Spadj = int32(-p.From.Offset)
			}
			if p.From.Type == obj.TYPE_ADDR && p.From.Reg == REGSP && p.To.Type == obj.TYPE_REG && p.To.Reg == REGSP {
				p.Spadj = int32(-p.From.Offset)
			}

		case obj.AGETCALLERPC:
			if cursym.Leaf() {

				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGLINK
			} else {

				p.As = AMOVW
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REGSP
			}
		}
	}
}

func (c *ctxt5) stacksplit(p *obj.Prog, framesize int32) *obj.Prog {

	p = obj.Appendp(p, c.newprog)

	p.As = AMOVW
	p.From.Type = obj.TYPE_MEM
	p.From.Reg = REGG
	p.From.Offset = 2 * int64(c.ctxt.Arch.PtrSize)
	if c.cursym.CFunc() {
		p.From.Offset = 3 * int64(c.ctxt.Arch.PtrSize)
	}
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_R1

	if framesize <= objabi.StackSmall {

		p = obj.Appendp(p, c.newprog)

		p.As = ACMP
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R1
		p.Reg = REGSP
	} else if framesize <= objabi.StackBig {

		p = obj.Appendp(p, c.newprog)

		p.As = AMOVW
		p.From.Type = obj.TYPE_ADDR
		p.From.Reg = REGSP
		p.From.Offset = -(int64(framesize) - objabi.StackSmall)
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
		p.From.Offset = int64(uint32(objabi.StackPreempt & (1<<32 - 1)))
		p.Reg = REG_R1

		p = obj.Appendp(p, c.newprog)
		p.As = AMOVW
		p.From.Type = obj.TYPE_ADDR
		p.From.Reg = REGSP
		p.From.Offset = objabi.StackGuard
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2
		p.Scond = C_SCOND_NE

		p = obj.Appendp(p, c.newprog)
		p.As = ASUB
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2
		p.Scond = C_SCOND_NE

		p = obj.Appendp(p, c.newprog)
		p.As = AMOVW
		p.From.Type = obj.TYPE_ADDR
		p.From.Offset = int64(framesize) + (objabi.StackGuard - objabi.StackSmall)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R3
		p.Scond = C_SCOND_NE

		p = obj.Appendp(p, c.newprog)
		p.As = ACMP
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R3
		p.Reg = REG_R2
		p.Scond = C_SCOND_NE
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

	movw := obj.Appendp(pcdata, c.newprog)
	movw.As = AMOVW
	movw.From.Type = obj.TYPE_REG
	movw.From.Reg = REGLINK
	movw.To.Type = obj.TYPE_REG
	movw.To.Reg = REG_R3

	bls.Pcond = movw

	call := obj.Appendp(movw, c.newprog)
	call.As = obj.ACALL
	call.To.Type = obj.TYPE_BRANCH
	morestack := "runtime.morestack"
	switch {
	case c.cursym.CFunc():
		morestack = "runtime.morestackc"
	case !c.cursym.Func.Text.From.Sym.NeedCtxt():
		morestack = "runtime.morestack_noctxt"
	}
	call.To.Sym = c.ctxt.Lookup(morestack)

	b := obj.Appendp(call, c.newprog)
	b.As = obj.AJMP
	b.To.Type = obj.TYPE_BRANCH
	b.Pcond = c.cursym.Func.Text.Link
	b.Spadj = +framesize

	return bls
}

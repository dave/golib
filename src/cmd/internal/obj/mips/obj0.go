package mips

import (
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"math"
)

func progedit(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {
	c := ctxt0{ctxt: ctxt, newprog: newprog}

	p.From.Class = 0
	p.To.Class = 0

	switch p.As {
	case AJMP,
		AJAL,
		ARET,
		obj.ADUFFZERO,
		obj.ADUFFCOPY:
		if p.To.Sym != nil {
			p.To.Type = obj.TYPE_BRANCH
		}
	}

	switch p.As {
	case AMOVF:
		if p.From.Type == obj.TYPE_FCONST {
			f32 := float32(p.From.Val.(float64))
			if math.Float32bits(f32) == 0 {
				p.As = AMOVW
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGZERO
				break
			}
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = ctxt.Float32Sym(f32)
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}

	case AMOVD:
		if p.From.Type == obj.TYPE_FCONST {
			f64 := p.From.Val.(float64)
			if math.Float64bits(f64) == 0 && c.ctxt.Arch.Family == sys.MIPS64 {
				p.As = AMOVV
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGZERO
				break
			}
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = ctxt.Float64Sym(f64)
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}

	case AMOVV:
		if p.From.Type == obj.TYPE_CONST && p.From.Name == obj.NAME_NONE && p.From.Reg == 0 && int64(int32(p.From.Offset)) != p.From.Offset {
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = ctxt.Int64Sym(p.From.Offset)
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}
	}

	switch p.As {
	case ASUB:
		if p.From.Type == obj.TYPE_CONST {
			p.From.Offset = -p.From.Offset
			p.As = AADD
		}

	case ASUBU:
		if p.From.Type == obj.TYPE_CONST {
			p.From.Offset = -p.From.Offset
			p.As = AADDU
		}

	case ASUBV:
		if p.From.Type == obj.TYPE_CONST {
			p.From.Offset = -p.From.Offset
			p.As = AADDV
		}

	case ASUBVU:
		if p.From.Type == obj.TYPE_CONST {
			p.From.Offset = -p.From.Offset
			p.As = AADDVU
		}
	}
}

func (psess *PackageSession) preprocess(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {

	c := ctxt0{ctxt: ctxt, newprog: newprog, cursym: cursym}

	nosched := true

	if c.cursym.Func.Text == nil || c.cursym.Func.Text.Link == nil {
		return
	}

	p := c.cursym.Func.Text
	textstksiz := p.To.Offset
	if textstksiz == -ctxt.FixedFrameSize() {

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

	var q *obj.Prog
	var q1 *obj.Prog
	for p := c.cursym.Func.Text; p != nil; p = p.Link {
		switch p.As {

		case obj.ATEXT:
			q = p

			p.Mark |= LABEL | LEAF | SYNC
			if p.Link != nil {
				p.Link.Mark |= LABEL
			}

		case AMOVW,
			AMOVV:
			q = p
			if p.To.Type == obj.TYPE_REG && p.To.Reg >= REG_SPECIAL {
				p.Mark |= LABEL | SYNC
				break
			}
			if p.From.Type == obj.TYPE_REG && p.From.Reg >= REG_SPECIAL {
				p.Mark |= LABEL | SYNC
			}

		case ASYSCALL,
			AWORD,
			ATLBWR,
			ATLBWI,
			ATLBP,
			ATLBR:
			q = p
			p.Mark |= LABEL | SYNC

		case ANOR:
			q = p
			if p.To.Type == obj.TYPE_REG {
				if p.To.Reg == REGZERO {
					p.Mark |= LABEL | SYNC
				}
			}

		case ABGEZAL,
			ABLTZAL,
			AJAL,
			obj.ADUFFZERO,
			obj.ADUFFCOPY:
			c.cursym.Func.Text.Mark &^= LEAF
			fallthrough

		case AJMP,
			ABEQ,
			ABGEZ,
			ABGTZ,
			ABLEZ,
			ABLTZ,
			ABNE,
			ABFPT, ABFPF:
			if p.As == ABFPT || p.As == ABFPF {

				p.Mark |= SYNC
			} else {
				p.Mark |= BRANCH
			}
			q = p
			q1 = p.Pcond
			if q1 != nil {
				for q1.As == obj.ANOP {
					q1 = q1.Link
					p.Pcond = q1
				}

				if q1.Mark&LEAF == 0 {
					q1.Mark |= LABEL
				}
			}

			q1 = p.Link
			if q1 != nil {
				q1.Mark |= LABEL
			}
			continue

		case ARET:
			q = p
			if p.Link != nil {
				p.Link.Mark |= LABEL
			}
			continue

		case obj.ANOP:
			q1 = p.Link
			q.Link = q1
			q1.Mark |= p.Mark
			continue

		default:
			q = p
			continue
		}
	}

	var mov, add obj.As
	if c.ctxt.Arch.Family == sys.MIPS64 {
		add = AADDV
		mov = AMOVV
	} else {
		add = AADDU
		mov = AMOVW
	}

	autosize := int32(0)
	var p1 *obj.Prog
	var p2 *obj.Prog
	for p := c.cursym.Func.Text; p != nil; p = p.Link {
		o := p.As
		switch o {
		case obj.ATEXT:
			autosize = int32(textstksiz)

			if p.Mark&LEAF != 0 && autosize == 0 {

				p.From.Sym.Set(obj.AttrNoFrame, true)
			}

			if !p.From.Sym.NoFrame() {

				autosize += int32(c.ctxt.FixedFrameSize())
			}

			if autosize&4 != 0 && c.ctxt.Arch.Family == sys.MIPS64 {
				autosize += 4
			}

			if autosize == 0 && c.cursym.Func.Text.Mark&LEAF == 0 {
				if c.cursym.Func.Text.From.Sym.NoSplit() {
					if ctxt.Debugvlog {
						ctxt.Logf("save suppressed in: %s\n", c.cursym.Name)
					}

					c.cursym.Func.Text.Mark |= LEAF
				}
			}

			p.To.Offset = int64(autosize) - ctxt.FixedFrameSize()

			if c.cursym.Func.Text.Mark&LEAF != 0 {
				c.cursym.Set(obj.AttrLeaf, true)
				if p.From.Sym.NoFrame() {
					break
				}
			}

			if !p.From.Sym.NoSplit() {
				p = c.stacksplit(p, autosize)
			}

			q = p

			if autosize != 0 {

				q = obj.Appendp(q, newprog)
				q.As = mov
				q.Pos = p.Pos
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REGLINK
				q.To.Type = obj.TYPE_MEM
				q.To.Offset = int64(-autosize)
				q.To.Reg = REGSP

				q = obj.Appendp(q, newprog)
				q.As = add
				q.Pos = p.Pos
				q.From.Type = obj.TYPE_CONST
				q.From.Offset = int64(-autosize)
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REGSP
				q.Spadj = +autosize
			}

			if c.cursym.Func.Text.From.Sym.Wrapper() && c.cursym.Func.Text.Mark&LEAF == 0 {

				q = obj.Appendp(q, newprog)

				q.As = mov
				q.From.Type = obj.TYPE_MEM
				q.From.Reg = REGG
				q.From.Offset = 4 * int64(c.ctxt.Arch.PtrSize)
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R1

				q = obj.Appendp(q, newprog)
				q.As = ABEQ
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R1
				q.To.Type = obj.TYPE_BRANCH
				q.Mark |= BRANCH
				p1 = q

				q = obj.Appendp(q, newprog)
				q.As = mov
				q.From.Type = obj.TYPE_MEM
				q.From.Reg = REG_R1
				q.From.Offset = 0
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R2

				q = obj.Appendp(q, newprog)
				q.As = add
				q.From.Type = obj.TYPE_CONST
				q.From.Offset = int64(autosize) + ctxt.FixedFrameSize()
				q.Reg = REGSP
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R3

				q = obj.Appendp(q, newprog)
				q.As = ABNE
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R2
				q.Reg = REG_R3
				q.To.Type = obj.TYPE_BRANCH
				q.Mark |= BRANCH
				p2 = q

				q = obj.Appendp(q, newprog)
				q.As = add
				q.From.Type = obj.TYPE_CONST
				q.From.Offset = ctxt.FixedFrameSize()
				q.Reg = REGSP
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REG_R2

				q = obj.Appendp(q, newprog)
				q.As = mov
				q.From.Type = obj.TYPE_REG
				q.From.Reg = REG_R2
				q.To.Type = obj.TYPE_MEM
				q.To.Reg = REG_R1
				q.To.Offset = 0

				q = obj.Appendp(q, newprog)

				q.As = obj.ANOP
				p1.Pcond = q
				p2.Pcond = q
			}

		case ARET:
			if p.From.Type == obj.TYPE_CONST {
				ctxt.Diag("using BECOME (%v) is not supported!", p)
				break
			}

			retSym := p.To.Sym
			p.To.Name = obj.NAME_NONE
			p.To.Sym = nil

			if c.cursym.Func.Text.Mark&LEAF != 0 {
				if autosize == 0 {
					p.As = AJMP
					p.From = obj.Addr{}
					if retSym != nil {
						p.To.Type = obj.TYPE_BRANCH
						p.To.Name = obj.NAME_EXTERN
						p.To.Sym = retSym
					} else {
						p.To.Type = obj.TYPE_MEM
						p.To.Reg = REGLINK
						p.To.Offset = 0
					}
					p.Mark |= BRANCH
					break
				}

				p.As = add
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = int64(autosize)
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REGSP
				p.Spadj = -autosize

				q = c.newprog()
				q.As = AJMP
				q.Pos = p.Pos
				q.To.Type = obj.TYPE_MEM
				q.To.Offset = 0
				q.To.Reg = REGLINK
				q.Mark |= BRANCH
				q.Spadj = +autosize

				q.Link = p.Link
				p.Link = q
				break
			}

			p.As = mov
			p.From.Type = obj.TYPE_MEM
			p.From.Offset = 0
			p.From.Reg = REGSP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REGLINK

			if autosize != 0 {
				q = c.newprog()
				q.As = add
				q.Pos = p.Pos
				q.From.Type = obj.TYPE_CONST
				q.From.Offset = int64(autosize)
				q.To.Type = obj.TYPE_REG
				q.To.Reg = REGSP
				q.Spadj = -autosize

				q.Link = p.Link
				p.Link = q
			}

			q1 = c.newprog()
			q1.As = AJMP
			q1.Pos = p.Pos
			if retSym != nil {
				q1.To.Type = obj.TYPE_BRANCH
				q1.To.Name = obj.NAME_EXTERN
				q1.To.Sym = retSym
			} else {
				q1.To.Type = obj.TYPE_MEM
				q1.To.Offset = 0
				q1.To.Reg = REGLINK
			}
			q1.Mark |= BRANCH
			q1.Spadj = +autosize

			q1.Link = q.Link
			q.Link = q1

		case AADD,
			AADDU,
			AADDV,
			AADDVU:
			if p.To.Type == obj.TYPE_REG && p.To.Reg == REGSP && p.From.Type == obj.TYPE_CONST {
				p.Spadj = int32(-p.From.Offset)
			}

		case obj.AGETCALLERPC:
			if cursym.Leaf() {

				p.As = mov
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REGLINK
			} else {

				p.As = mov
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REGSP
			}
		}
	}

	if c.ctxt.Arch.Family == sys.MIPS {

		for p = c.cursym.Func.Text; p != nil; p = p1 {
			p1 = p.Link

			if p.As != AMOVD {
				continue
			}
			if p.From.Type != obj.TYPE_MEM && p.To.Type != obj.TYPE_MEM {
				continue
			}

			p.As = AMOVF
			q = c.newprog()
			*q = *p
			q.Link = p.Link
			p.Link = q
			p1 = q.Link

			var addrOff int64
			if c.ctxt.Arch.ByteOrder == binary.BigEndian {
				addrOff = 4
			}
			if p.From.Type == obj.TYPE_MEM {
				reg := REG_F0 + (p.To.Reg-REG_F0)&^1
				p.To.Reg = reg
				q.To.Reg = reg + 1
				p.From.Offset += addrOff
				q.From.Offset += 4 - addrOff
			} else if p.To.Type == obj.TYPE_MEM {
				reg := REG_F0 + (p.From.Reg-REG_F0)&^1
				p.From.Reg = reg
				q.From.Reg = reg + 1
				p.To.Offset += addrOff
				q.To.Offset += 4 - addrOff
			}
		}
	}

	if nosched {

		for p = c.cursym.Func.Text; p != nil; p = p.Link {
			if p.Mark&BRANCH != 0 {
				c.addnop(p)
			}
		}
		return
	}

	q = nil
	q1 = c.cursym.Func.Text
	o := 0
	for p = c.cursym.Func.Text; p != nil; p = p1 {
		p1 = p.Link
		o++
		if p.Mark&NOSCHED != 0 {
			if q1 != p {
				c.sched(psess, q1, q)
			}
			for ; p != nil; p = p.Link {
				if p.Mark&NOSCHED == 0 {
					break
				}
				q = p
			}
			p1 = p
			q1 = p
			o = 0
			continue
		}
		if p.Mark&(LABEL|SYNC) != 0 {
			if q1 != p {
				c.sched(psess, q1, q)
			}
			q1 = p
			o = 1
		}
		if p.Mark&(BRANCH|SYNC) != 0 {
			c.sched(psess, q1, p)
			q1 = p1
			o = 0
		}
		if o >= NSCHED {
			c.sched(psess, q1, p)
			q1 = p1
			o = 0
		}
		q = p
	}
}

func (c *ctxt0) stacksplit(p *obj.Prog, framesize int32) *obj.Prog {
	var mov, add, sub obj.As

	if c.ctxt.Arch.Family == sys.MIPS64 {
		add = AADDV
		mov = AMOVV
		sub = ASUBVU
	} else {
		add = AADDU
		mov = AMOVW
		sub = ASUBU
	}

	p = obj.Appendp(p, c.newprog)

	p.As = mov
	p.From.Type = obj.TYPE_MEM
	p.From.Reg = REGG
	p.From.Offset = 2 * int64(c.ctxt.Arch.PtrSize)
	if c.cursym.CFunc() {
		p.From.Offset = 3 * int64(c.ctxt.Arch.PtrSize)
	}
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_R1

	var q *obj.Prog
	if framesize <= objabi.StackSmall {

		p = obj.Appendp(p, c.newprog)

		p.As = ASGTU
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REGSP
		p.Reg = REG_R1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R1
	} else if framesize <= objabi.StackBig {

		p = obj.Appendp(p, c.newprog)

		p.As = add
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = -(int64(framesize) - objabi.StackSmall)
		p.Reg = REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2

		p = obj.Appendp(p, c.newprog)
		p.As = ASGTU
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R2
		p.Reg = REG_R1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R1
	} else {

		p = obj.Appendp(p, c.newprog)

		p.As = mov
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = objabi.StackPreempt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2

		p = obj.Appendp(p, c.newprog)
		q = p
		p.As = ABEQ
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R1
		p.Reg = REG_R2
		p.To.Type = obj.TYPE_BRANCH
		p.Mark |= BRANCH

		p = obj.Appendp(p, c.newprog)
		p.As = add
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = objabi.StackGuard
		p.Reg = REGSP
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2

		p = obj.Appendp(p, c.newprog)
		p.As = sub
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R2

		p = obj.Appendp(p, c.newprog)
		p.As = mov
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = int64(framesize) + objabi.StackGuard - objabi.StackSmall
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R1

		p = obj.Appendp(p, c.newprog)
		p.As = ASGTU
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R2
		p.Reg = REG_R1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_R1
	}

	p = obj.Appendp(p, c.newprog)
	q1 := p

	p.As = ABNE
	p.From.Type = obj.TYPE_REG
	p.From.Reg = REG_R1
	p.To.Type = obj.TYPE_BRANCH
	p.Mark |= BRANCH

	p = obj.Appendp(p, c.newprog)

	p.As = mov
	p.From.Type = obj.TYPE_REG
	p.From.Reg = REGLINK
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_R3
	if q != nil {
		q.Pcond = p
		p.Mark |= LABEL
	}

	p = c.ctxt.EmitEntryLiveness(c.cursym, p, c.newprog)

	p = obj.Appendp(p, c.newprog)

	p.As = AJAL
	p.To.Type = obj.TYPE_BRANCH
	if c.cursym.CFunc() {
		p.To.Sym = c.ctxt.Lookup("runtime.morestackc")
	} else if !c.cursym.Func.Text.From.Sym.NeedCtxt() {
		p.To.Sym = c.ctxt.Lookup("runtime.morestack_noctxt")
	} else {
		p.To.Sym = c.ctxt.Lookup("runtime.morestack")
	}
	p.Mark |= BRANCH

	p = obj.Appendp(p, c.newprog)

	p.As = AJMP
	p.To.Type = obj.TYPE_BRANCH
	p.Pcond = c.cursym.Func.Text.Link
	p.Mark |= BRANCH

	p = obj.Appendp(p, c.newprog)

	p.As = obj.ANOP
	q1.Pcond = p

	return p
}

func (c *ctxt0) addnop(p *obj.Prog) {
	q := c.newprog()
	q.As = ANOOP
	q.Pos = p.Pos
	q.Link = p.Link
	p.Link = q
}

const (
	E_HILO  = 1 << 0
	E_FCR   = 1 << 1
	E_MCR   = 1 << 2
	E_MEM   = 1 << 3
	E_MEMSP = 1 << 4 /* uses offset and size */
	E_MEMSB = 1 << 5 /* uses offset and size */
	ANYMEM  = E_MEM | E_MEMSP | E_MEMSB
	//DELAY = LOAD|BRANCH|FCMP
	DELAY = BRANCH /* only schedule branch */
)

type Dep struct {
	ireg uint32
	freg uint32
	cc   uint32
}

type Sch struct {
	p       obj.Prog
	set     Dep
	used    Dep
	soffset int32
	size    uint8
	nop     uint8
	comp    bool
}

func (c *ctxt0) sched(psess *PackageSession, p0, pe *obj.Prog) {
	var sch [NSCHED]Sch

	s := sch[:]
	for p := p0; ; p = p.Link {
		s[0].p = *p
		c.markregused(psess, &s[0])
		if p == pe {
			break
		}
		s = s[1:]
	}
	se := s

	for i := cap(sch) - cap(se); i >= 0; i-- {
		s = sch[i:]
		if s[0].p.Mark&DELAY == 0 {
			continue
		}
		if -cap(s) < -cap(se) {
			if !conflict(&s[0], &s[1]) {
				continue
			}
		}

		var t []Sch
		var j int
		for j = cap(sch) - cap(s) - 1; j >= 0; j-- {
			t = sch[j:]
			if t[0].comp {
				if s[0].p.Mark&BRANCH != 0 {
					continue
				}
			}
			if t[0].p.Mark&DELAY != 0 {
				if -cap(s) >= -cap(se) || conflict(&t[0], &s[1]) {
					continue
				}
			}
			for u := t[1:]; -cap(u) <= -cap(s); u = u[1:] {
				if c.depend(&u[0], &t[0]) {
					continue
				}
			}
			goto out2
		}

		if s[0].p.Mark&BRANCH != 0 {
			s[0].nop = 1
		}
		continue

	out2:

		stmp := t[0]
		copy(t[:i-j], t[1:i-j+1])
		s[0] = stmp

		if t[i-j-1].p.Mark&BRANCH != 0 {

			t[i-j-1].p.Spadj += t[i-j].p.Spadj
			t[i-j].p.Spadj = 0
		}

		i--
	}

	/*
	 * put it all back
	 */
	var p *obj.Prog
	var q *obj.Prog
	for s, p = sch[:], p0; -cap(s) <= -cap(se); s, p = s[1:], q {
		q = p.Link
		if q != s[0].p.Link {
			*p = s[0].p
			p.Link = q
		}
		for s[0].nop != 0 {
			s[0].nop--
			c.addnop(p)
		}
	}
}

func (c *ctxt0) markregused(psess *PackageSession, s *Sch) {
	p := &s.p
	s.comp = c.compound(psess, p)
	s.nop = 0
	if s.comp {
		s.set.ireg |= 1 << (REGTMP - REG_R0)
		s.used.ireg |= 1 << (REGTMP - REG_R0)
	}

	ar := 0
	ad := 0
	ld := 0
	sz := 20

	switch p.As {
	case obj.ATEXT:
		c.autosize = int32(p.To.Offset + 8)
		ad = 1

	case AJAL:
		r := p.Reg
		if r == 0 {
			r = REGLINK
		}
		s.set.ireg |= 1 << uint(r-REG_R0)
		ar = 1
		ad = 1

	case ABGEZAL,
		ABLTZAL:
		s.set.ireg |= 1 << (REGLINK - REG_R0)
		fallthrough
	case ABEQ,
		ABGEZ,
		ABGTZ,
		ABLEZ,
		ABLTZ,
		ABNE:
		ar = 1
		ad = 1

	case ABFPT,
		ABFPF:
		ad = 1
		s.used.cc |= E_FCR

	case ACMPEQD,
		ACMPEQF,
		ACMPGED,
		ACMPGEF,
		ACMPGTD,
		ACMPGTF:
		ar = 1
		s.set.cc |= E_FCR
		p.Mark |= FCMP

	case AJMP:
		ar = 1
		ad = 1

	case AMOVB,
		AMOVBU:
		sz = 1
		ld = 1

	case AMOVH,
		AMOVHU:
		sz = 2
		ld = 1

	case AMOVF,
		AMOVW,
		AMOVWL,
		AMOVWR:
		sz = 4
		ld = 1

	case AMOVD,
		AMOVV,
		AMOVVL,
		AMOVVR:
		sz = 8
		ld = 1

	case ADIV,
		ADIVU,
		AMUL,
		AMULU,
		AREM,
		AREMU,
		ADIVV,
		ADIVVU,
		AMULV,
		AMULVU,
		AREMV,
		AREMVU:
		s.set.cc = E_HILO
		fallthrough
	case AADD,
		AADDU,
		AADDV,
		AADDVU,
		AAND,
		ANOR,
		AOR,
		ASGT,
		ASGTU,
		ASLL,
		ASRA,
		ASRL,
		ASLLV,
		ASRAV,
		ASRLV,
		ASUB,
		ASUBU,
		ASUBV,
		ASUBVU,
		AXOR,

		AADDD,
		AADDF,
		AADDW,
		ASUBD,
		ASUBF,
		ASUBW,
		AMULF,
		AMULD,
		AMULW,
		ADIVF,
		ADIVD,
		ADIVW:
		if p.Reg == 0 {
			if p.To.Type == obj.TYPE_REG {
				p.Reg = p.To.Reg
			}

		}
	}

	cls := int(p.To.Class)
	if cls == 0 {
		cls = c.aclass(&p.To) + 1
		p.To.Class = int8(cls)
	}
	cls--
	switch cls {
	default:
		fmt.Printf("unknown class %d %v\n", cls, p)

	case C_ZCON,
		C_SCON,
		C_ADD0CON,
		C_AND0CON,
		C_ADDCON,
		C_ANDCON,
		C_UCON,
		C_LCON,
		C_NONE,
		C_SBRA,
		C_LBRA,
		C_ADDR,
		C_TEXTSIZE:
		break

	case C_HI,
		C_LO:
		s.set.cc |= E_HILO

	case C_FCREG:
		s.set.cc |= E_FCR

	case C_MREG:
		s.set.cc |= E_MCR

	case C_ZOREG,
		C_SOREG,
		C_LOREG:
		cls = int(p.To.Reg)
		s.used.ireg |= 1 << uint(cls-REG_R0)
		if ad != 0 {
			break
		}
		s.size = uint8(sz)
		s.soffset = c.regoff(&p.To)

		m := uint32(ANYMEM)
		if cls == REGSB {
			m = E_MEMSB
		}
		if cls == REGSP {
			m = E_MEMSP
		}

		if ar != 0 {
			s.used.cc |= m
		} else {
			s.set.cc |= m
		}

	case C_SACON,
		C_LACON:
		s.used.ireg |= 1 << (REGSP - REG_R0)

	case C_SECON,
		C_LECON:
		s.used.ireg |= 1 << (REGSB - REG_R0)

	case C_REG:
		if ar != 0 {
			s.used.ireg |= 1 << uint(p.To.Reg-REG_R0)
		} else {
			s.set.ireg |= 1 << uint(p.To.Reg-REG_R0)
		}

	case C_FREG:
		if ar != 0 {
			s.used.freg |= 1 << uint(p.To.Reg-REG_F0)
		} else {
			s.set.freg |= 1 << uint(p.To.Reg-REG_F0)
		}
		if ld != 0 && p.From.Type == obj.TYPE_REG {
			p.Mark |= LOAD
		}

	case C_SAUTO,
		C_LAUTO:
		s.used.ireg |= 1 << (REGSP - REG_R0)
		if ad != 0 {
			break
		}
		s.size = uint8(sz)
		s.soffset = c.regoff(&p.To)

		if ar != 0 {
			s.used.cc |= E_MEMSP
		} else {
			s.set.cc |= E_MEMSP
		}

	case C_SEXT,
		C_LEXT:
		s.used.ireg |= 1 << (REGSB - REG_R0)
		if ad != 0 {
			break
		}
		s.size = uint8(sz)
		s.soffset = c.regoff(&p.To)

		if ar != 0 {
			s.used.cc |= E_MEMSB
		} else {
			s.set.cc |= E_MEMSB
		}
	}

	cls = int(p.From.Class)
	if cls == 0 {
		cls = c.aclass(&p.From) + 1
		p.From.Class = int8(cls)
	}
	cls--
	switch cls {
	default:
		fmt.Printf("unknown class %d %v\n", cls, p)

	case C_ZCON,
		C_SCON,
		C_ADD0CON,
		C_AND0CON,
		C_ADDCON,
		C_ANDCON,
		C_UCON,
		C_LCON,
		C_NONE,
		C_SBRA,
		C_LBRA,
		C_ADDR,
		C_TEXTSIZE:
		break

	case C_HI,
		C_LO:
		s.used.cc |= E_HILO

	case C_FCREG:
		s.used.cc |= E_FCR

	case C_MREG:
		s.used.cc |= E_MCR

	case C_ZOREG,
		C_SOREG,
		C_LOREG:
		cls = int(p.From.Reg)
		s.used.ireg |= 1 << uint(cls-REG_R0)
		if ld != 0 {
			p.Mark |= LOAD
		}
		s.size = uint8(sz)
		s.soffset = c.regoff(&p.From)

		m := uint32(ANYMEM)
		if cls == REGSB {
			m = E_MEMSB
		}
		if cls == REGSP {
			m = E_MEMSP
		}

		s.used.cc |= m

	case C_SACON,
		C_LACON:
		cls = int(p.From.Reg)
		if cls == 0 {
			cls = REGSP
		}
		s.used.ireg |= 1 << uint(cls-REG_R0)

	case C_SECON,
		C_LECON:
		s.used.ireg |= 1 << (REGSB - REG_R0)

	case C_REG:
		s.used.ireg |= 1 << uint(p.From.Reg-REG_R0)

	case C_FREG:
		s.used.freg |= 1 << uint(p.From.Reg-REG_F0)
		if ld != 0 && p.To.Type == obj.TYPE_REG {
			p.Mark |= LOAD
		}

	case C_SAUTO,
		C_LAUTO:
		s.used.ireg |= 1 << (REGSP - REG_R0)
		if ld != 0 {
			p.Mark |= LOAD
		}
		if ad != 0 {
			break
		}
		s.size = uint8(sz)
		s.soffset = c.regoff(&p.From)

		s.used.cc |= E_MEMSP

	case C_SEXT:
	case C_LEXT:
		s.used.ireg |= 1 << (REGSB - REG_R0)
		if ld != 0 {
			p.Mark |= LOAD
		}
		if ad != 0 {
			break
		}
		s.size = uint8(sz)
		s.soffset = c.regoff(&p.From)

		s.used.cc |= E_MEMSB
	}

	cls = int(p.Reg)
	if cls != 0 {
		if REG_F0 <= cls && cls <= REG_F31 {
			s.used.freg |= 1 << uint(cls-REG_F0)
		} else {
			s.used.ireg |= 1 << uint(cls-REG_R0)
		}
	}
	s.set.ireg &^= (1 << (REGZERO - REG_R0))
}

/*
 * test to see if two instructions can be
 * interchanged without changing semantics
 */
func (c *ctxt0) depend(sa, sb *Sch) bool {
	if sa.set.ireg&(sb.set.ireg|sb.used.ireg) != 0 {
		return true
	}
	if sb.set.ireg&sa.used.ireg != 0 {
		return true
	}

	if sa.set.freg&(sb.set.freg|sb.used.freg) != 0 {
		return true
	}
	if sb.set.freg&sa.used.freg != 0 {
		return true
	}

	if sa.used.cc&sb.used.cc&E_MEM != 0 {
		if sa.p.Reg == sb.p.Reg {
			if c.regoff(&sa.p.From) == c.regoff(&sb.p.From) {
				return true
			}
		}
	}

	x := (sa.set.cc & (sb.set.cc | sb.used.cc)) | (sb.set.cc & sa.used.cc)
	if x != 0 {

		if x != E_MEMSP && x != E_MEMSB {
			return true
		}
		x = sa.set.cc | sb.set.cc | sa.used.cc | sb.used.cc
		if x&E_MEM != 0 {
			return true
		}
		if offoverlap(sa, sb) {
			return true
		}
	}

	return false
}

func offoverlap(sa, sb *Sch) bool {
	if sa.soffset < sb.soffset {
		if sa.soffset+int32(sa.size) > sb.soffset {
			return true
		}
		return false
	}
	if sb.soffset+int32(sb.size) > sa.soffset {
		return true
	}
	return false
}

/*
 * test 2 adjacent instructions
 * and find out if inserted instructions
 * are desired to prevent stalls.
 */
func conflict(sa, sb *Sch) bool {
	if sa.set.ireg&sb.used.ireg != 0 {
		return true
	}
	if sa.set.freg&sb.used.freg != 0 {
		return true
	}
	if sa.set.cc&sb.used.cc != 0 {
		return true
	}
	return false
}

func (c *ctxt0) compound(psess *PackageSession, p *obj.Prog) bool {
	o := c.oplook(psess, p)
	if o.size != 4 {
		return true
	}
	if p.To.Type == obj.TYPE_REG && p.To.Reg == REGSB {
		return true
	}
	return false
}

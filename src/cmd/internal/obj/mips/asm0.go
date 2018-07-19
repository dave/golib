package mips

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"log"
	"sort"
)

// ctxt0 holds state while assembling a single function.
// Each function gets a fresh ctxt0.
// This allows for multiple functions to be safely concurrently assembled.
type ctxt0 struct {
	ctxt       *obj.Link
	newprog    obj.ProgAlloc
	cursym     *obj.LSym
	autosize   int32
	instoffset int64
	pc         int64
}

const (
	mips64FuncAlign = 8
)

const (
	r0iszero = 1
)

type Optab struct {
	as     obj.As
	a1     uint8
	a2     uint8
	a3     uint8
	type_  int8
	size   int8
	param  int16
	family sys.ArchFamily // 0 means both sys.MIPS and sys.MIPS64
}

func (psess *PackageSession) span0(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	p := cursym.Func.Text
	if p == nil || p.Link == nil {
		return
	}

	c := ctxt0{ctxt: ctxt, newprog: newprog, cursym: cursym, autosize: int32(p.To.Offset + ctxt.FixedFrameSize())}

	if psess.oprange[AOR&obj.AMask] == nil {
		c.ctxt.Diag("mips ops not initialized, call mips.buildop first")
	}

	pc := int64(0)
	p.Pc = pc

	var m int
	var o *Optab
	for p = p.Link; p != nil; p = p.Link {
		p.Pc = pc
		o = c.oplook(psess, p)
		m = int(o.size)
		if m == 0 {
			if p.As != obj.ANOP && p.As != obj.AFUNCDATA && p.As != obj.APCDATA {
				c.ctxt.Diag("zero-width instruction\n%v", p)
			}
			continue
		}

		pc += int64(m)
	}

	c.cursym.Size = pc

	bflag := 1

	var otxt int64
	var q *obj.Prog
	for bflag != 0 {
		bflag = 0
		pc = 0
		for p = c.cursym.Func.Text.Link; p != nil; p = p.Link {
			p.Pc = pc
			o = c.oplook(psess, p)

			if o.type_ == 6 && p.Pcond != nil {
				otxt = p.Pcond.Pc - pc
				if otxt < -(1<<17)+10 || otxt >= (1<<17)-10 {
					q = c.newprog()
					q.Link = p.Link
					p.Link = q
					q.As = AJMP
					q.Pos = p.Pos
					q.To.Type = obj.TYPE_BRANCH
					q.Pcond = p.Pcond
					p.Pcond = q
					q = c.newprog()
					q.Link = p.Link
					p.Link = q
					q.As = AJMP
					q.Pos = p.Pos
					q.To.Type = obj.TYPE_BRANCH
					q.Pcond = q.Link.Link

					c.addnop(p.Link)
					c.addnop(p)
					bflag = 1
				}
			}

			m = int(o.size)
			if m == 0 {
				if p.As != obj.ANOP && p.As != obj.AFUNCDATA && p.As != obj.APCDATA {
					c.ctxt.Diag("zero-width instruction\n%v", p)
				}
				continue
			}

			pc += int64(m)
		}

		c.cursym.Size = pc
	}
	if c.ctxt.Arch.Family == sys.MIPS64 {
		pc += -pc & (mips64FuncAlign - 1)
	}
	c.cursym.Size = pc

	c.cursym.Grow(c.cursym.Size)

	bp := c.cursym.P
	var i int32
	var out [4]uint32
	for p := c.cursym.Func.Text.Link; p != nil; p = p.Link {
		c.pc = p.Pc
		o = c.oplook(psess, p)
		if int(o.size) > 4*len(out) {
			log.Fatalf("out array in span0 is too small, need at least %d for %v", o.size/4, p)
		}
		c.asmout(p, o, out[:])
		for i = 0; i < int32(o.size/4); i++ {
			c.ctxt.Arch.ByteOrder.PutUint32(bp, out[i])
			bp = bp[4:]
		}
	}
}

func isint32(v int64) bool {
	return int64(int32(v)) == v
}

func isuint32(v uint64) bool {
	return uint64(uint32(v)) == v
}

func (c *ctxt0) aclass(a *obj.Addr) int {
	switch a.Type {
	case obj.TYPE_NONE:
		return C_NONE

	case obj.TYPE_REG:
		if REG_R0 <= a.Reg && a.Reg <= REG_R31 {
			return C_REG
		}
		if REG_F0 <= a.Reg && a.Reg <= REG_F31 {
			return C_FREG
		}
		if REG_M0 <= a.Reg && a.Reg <= REG_M31 {
			return C_MREG
		}
		if REG_FCR0 <= a.Reg && a.Reg <= REG_FCR31 {
			return C_FCREG
		}
		if a.Reg == REG_LO {
			return C_LO
		}
		if a.Reg == REG_HI {
			return C_HI
		}
		return C_GOK

	case obj.TYPE_MEM:
		switch a.Name {
		case obj.NAME_EXTERN,
			obj.NAME_STATIC:
			if a.Sym == nil {
				break
			}
			c.instoffset = a.Offset
			if a.Sym != nil {
				if a.Sym.Type == objabi.STLSBSS {
					return C_TLS
				}
				return C_ADDR
			}
			return C_LEXT

		case obj.NAME_AUTO:
			if a.Reg == REGSP {

				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset
			if c.instoffset >= -BIG && c.instoffset < BIG {
				return C_SAUTO
			}
			return C_LAUTO

		case obj.NAME_PARAM:
			if a.Reg == REGSP {

				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset + c.ctxt.FixedFrameSize()
			if c.instoffset >= -BIG && c.instoffset < BIG {
				return C_SAUTO
			}
			return C_LAUTO

		case obj.NAME_NONE:
			c.instoffset = a.Offset
			if c.instoffset == 0 {
				return C_ZOREG
			}
			if c.instoffset >= -BIG && c.instoffset < BIG {
				return C_SOREG
			}
			return C_LOREG
		}

		return C_GOK

	case obj.TYPE_TEXTSIZE:
		return C_TEXTSIZE

	case obj.TYPE_CONST,
		obj.TYPE_ADDR:
		switch a.Name {
		case obj.NAME_NONE:
			c.instoffset = a.Offset
			if a.Reg != 0 {
				if -BIG <= c.instoffset && c.instoffset <= BIG {
					return C_SACON
				}
				if isint32(c.instoffset) {
					return C_LACON
				}
				return C_DACON
			}

		case obj.NAME_EXTERN,
			obj.NAME_STATIC:
			s := a.Sym
			if s == nil {
				return C_GOK
			}

			c.instoffset = a.Offset
			if s.Type == objabi.STLSBSS {
				return C_STCON
			}
			return C_LECON

		case obj.NAME_AUTO:
			if a.Reg == REGSP {

				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset
			if c.instoffset >= -BIG && c.instoffset < BIG {
				return C_SACON
			}
			return C_LACON

		case obj.NAME_PARAM:
			if a.Reg == REGSP {

				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset + c.ctxt.FixedFrameSize()
			if c.instoffset >= -BIG && c.instoffset < BIG {
				return C_SACON
			}
			return C_LACON

		default:
			return C_GOK
		}

		if c.instoffset >= 0 {
			if c.instoffset == 0 {
				return C_ZCON
			}
			if c.instoffset <= 0x7fff {
				return C_SCON
			}
			if c.instoffset <= 0xffff {
				return C_ANDCON
			}
			if c.instoffset&0xffff == 0 && isuint32(uint64(c.instoffset)) {
				return C_UCON
			}
			if isint32(c.instoffset) || isuint32(uint64(c.instoffset)) {
				return C_LCON
			}
			return C_LCON
		}

		if c.instoffset >= -0x8000 {
			return C_ADDCON
		}
		if c.instoffset&0xffff == 0 && isint32(c.instoffset) {
			return C_UCON
		}
		if isint32(c.instoffset) {
			return C_LCON
		}
		return C_LCON

	case obj.TYPE_BRANCH:
		return C_SBRA
	}

	return C_GOK
}

func prasm(p *obj.Prog) {
	fmt.Printf("%v\n", p)
}

func (c *ctxt0) oplook(psess *PackageSession, p *obj.Prog) *Optab {
	if psess.oprange[AOR&obj.AMask] == nil {
		c.ctxt.Diag("mips ops not initialized, call mips.buildop first")
	}

	a1 := int(p.Optab)
	if a1 != 0 {
		return &psess.optab[a1-1]
	}
	a1 = int(p.From.Class)
	if a1 == 0 {
		a1 = c.aclass(&p.From) + 1
		p.From.Class = int8(a1)
	}

	a1--
	a3 := int(p.To.Class)
	if a3 == 0 {
		a3 = c.aclass(&p.To) + 1
		p.To.Class = int8(a3)
	}

	a3--
	a2 := C_NONE
	if p.Reg != 0 {
		a2 = C_REG
	}

	ops := psess.oprange[p.As&obj.AMask]
	c1 := &psess.xcmp[a1]
	c3 := &psess.xcmp[a3]
	for i := range ops {
		op := &ops[i]
		if int(op.a2) == a2 && c1[op.a1] && c3[op.a3] && (op.family == 0 || c.ctxt.Arch.Family == op.family) {
			p.Optab = uint16(cap(psess.optab) - cap(ops) + i + 1)
			return op
		}
	}

	c.ctxt.Diag("illegal combination %v %v %v %v", p.As, psess.DRconv(a1), psess.DRconv(a2), psess.DRconv(a3))
	prasm(p)

	return &Optab{obj.AUNDEF, C_NONE, C_NONE, C_NONE, 49, 4, 0, 0}
}

func cmp(a int, b int) bool {
	if a == b {
		return true
	}
	switch a {
	case C_LCON:
		if b == C_ZCON || b == C_SCON || b == C_UCON || b == C_ADDCON || b == C_ANDCON {
			return true
		}

	case C_ADD0CON:
		if b == C_ADDCON {
			return true
		}
		fallthrough

	case C_ADDCON:
		if b == C_ZCON || b == C_SCON {
			return true
		}

	case C_AND0CON:
		if b == C_ANDCON {
			return true
		}
		fallthrough

	case C_ANDCON:
		if b == C_ZCON || b == C_SCON {
			return true
		}

	case C_UCON:
		if b == C_ZCON {
			return true
		}

	case C_SCON:
		if b == C_ZCON {
			return true
		}

	case C_LACON:
		if b == C_SACON {
			return true
		}

	case C_LBRA:
		if b == C_SBRA {
			return true
		}

	case C_LEXT:
		if b == C_SEXT {
			return true
		}

	case C_LAUTO:
		if b == C_SAUTO {
			return true
		}

	case C_REG:
		if b == C_ZCON {
			return r0iszero != 0
		}

	case C_LOREG:
		if b == C_ZOREG || b == C_SOREG {
			return true
		}

	case C_SOREG:
		if b == C_ZOREG {
			return true
		}
	}

	return false
}

type ocmp []Optab

func (x ocmp) Len() int {
	return len(x)
}

func (x ocmp) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x ocmp) Less(i, j int) bool {
	p1 := &x[i]
	p2 := &x[j]
	n := int(p1.as) - int(p2.as)
	if n != 0 {
		return n < 0
	}
	n = int(p1.a1) - int(p2.a1)
	if n != 0 {
		return n < 0
	}
	n = int(p1.a2) - int(p2.a2)
	if n != 0 {
		return n < 0
	}
	n = int(p1.a3) - int(p2.a3)
	if n != 0 {
		return n < 0
	}
	return false
}

func (psess *PackageSession) opset(a, b0 obj.As) {
	psess.
		oprange[a&obj.AMask] = psess.oprange[b0]
}

func (psess *PackageSession) buildop(ctxt *obj.Link) {
	if psess.oprange[AOR&obj.AMask] != nil {

		return
	}

	var n int

	for i := 0; i < C_NCLASS; i++ {
		for n = 0; n < C_NCLASS; n++ {
			if cmp(n, i) {
				psess.
					xcmp[i][n] = true
			}
		}
	}
	for n = 0; psess.optab[n].as != obj.AXXX; n++ {
	}
	sort.Sort(ocmp(psess.optab[:n]))
	for i := 0; i < n; i++ {
		r := psess.optab[i].as
		r0 := r & obj.AMask
		start := i
		for psess.optab[i].as == r {
			i++
		}
		psess.
			oprange[r0] = psess.optab[start:i]
		i--

		switch r {
		default:
			ctxt.Diag("unknown op in build: %v", r)
			ctxt.DiagFlush()
			log.Fatalf("bad code")

		case AABSF:
			psess.
				opset(AMOVFD, r0)
			psess.
				opset(AMOVDF, r0)
			psess.
				opset(AMOVWF, r0)
			psess.
				opset(AMOVFW, r0)
			psess.
				opset(AMOVWD, r0)
			psess.
				opset(AMOVDW, r0)
			psess.
				opset(ANEGF, r0)
			psess.
				opset(ANEGD, r0)
			psess.
				opset(AABSD, r0)
			psess.
				opset(ATRUNCDW, r0)
			psess.
				opset(ATRUNCFW, r0)
			psess.
				opset(ASQRTF, r0)
			psess.
				opset(ASQRTD, r0)

		case AMOVVF:
			psess.
				opset(AMOVVD, r0)
			psess.
				opset(AMOVFV, r0)
			psess.
				opset(AMOVDV, r0)
			psess.
				opset(ATRUNCDV, r0)
			psess.
				opset(ATRUNCFV, r0)

		case AADD:
			psess.
				opset(ASGT, r0)
			psess.
				opset(ASGTU, r0)
			psess.
				opset(AADDU, r0)

		case AADDV:
			psess.
				opset(AADDVU, r0)

		case AADDF:
			psess.
				opset(ADIVF, r0)
			psess.
				opset(ADIVD, r0)
			psess.
				opset(AMULF, r0)
			psess.
				opset(AMULD, r0)
			psess.
				opset(ASUBF, r0)
			psess.
				opset(ASUBD, r0)
			psess.
				opset(AADDD, r0)

		case AAND:
			psess.
				opset(AOR, r0)
			psess.
				opset(AXOR, r0)

		case ABEQ:
			psess.
				opset(ABNE, r0)

		case ABLEZ:
			psess.
				opset(ABGEZ, r0)
			psess.
				opset(ABGEZAL, r0)
			psess.
				opset(ABLTZ, r0)
			psess.
				opset(ABLTZAL, r0)
			psess.
				opset(ABGTZ, r0)

		case AMOVB:
			psess.
				opset(AMOVH, r0)

		case AMOVBU:
			psess.
				opset(AMOVHU, r0)

		case AMUL:
			psess.
				opset(AREM, r0)
			psess.
				opset(AREMU, r0)
			psess.
				opset(ADIVU, r0)
			psess.
				opset(AMULU, r0)
			psess.
				opset(ADIV, r0)

		case AMULV:
			psess.
				opset(ADIVV, r0)
			psess.
				opset(ADIVVU, r0)
			psess.
				opset(AMULVU, r0)
			psess.
				opset(AREMV, r0)
			psess.
				opset(AREMVU, r0)

		case ASLL:
			psess.
				opset(ASRL, r0)
			psess.
				opset(ASRA, r0)

		case ASLLV:
			psess.
				opset(ASRAV, r0)
			psess.
				opset(ASRLV, r0)

		case ASUB:
			psess.
				opset(ASUBU, r0)
			psess.
				opset(ANOR, r0)

		case ASUBV:
			psess.
				opset(ASUBVU, r0)

		case ASYSCALL:
			psess.
				opset(ASYNC, r0)
			psess.
				opset(ANOOP, r0)
			psess.
				opset(ATLBP, r0)
			psess.
				opset(ATLBR, r0)
			psess.
				opset(ATLBWI, r0)
			psess.
				opset(ATLBWR, r0)

		case ACMPEQF:
			psess.
				opset(ACMPGTF, r0)
			psess.
				opset(ACMPGTD, r0)
			psess.
				opset(ACMPGEF, r0)
			psess.
				opset(ACMPGED, r0)
			psess.
				opset(ACMPEQD, r0)

		case ABFPT:
			psess.
				opset(ABFPF, r0)

		case AMOVWL:
			psess.
				opset(AMOVWR, r0)

		case AMOVVL:
			psess.
				opset(AMOVVR, r0)

		case AMOVW,
			AMOVD,
			AMOVF,
			AMOVV,
			ABREAK,
			ARFE,
			AJAL,
			AJMP,
			AMOVWU,
			ALL,
			ALLV,
			ASC,
			ASCV,
			ANEGW,
			ANEGV,
			AWORD,
			obj.ANOP,
			obj.ATEXT,
			obj.AUNDEF,
			obj.AFUNCDATA,
			obj.APCDATA,
			obj.ADUFFZERO,
			obj.ADUFFCOPY:
			break

		case ACMOVN:
			psess.
				opset(ACMOVZ, r0)

		case ACMOVT:
			psess.
				opset(ACMOVF, r0)

		case ACLO:
			psess.
				opset(ACLZ, r0)

		case ATEQ:
			psess.
				opset(ATNE, r0)
		}
	}
}

func OP(x uint32, y uint32) uint32 {
	return x<<3 | y<<0
}

func SP(x uint32, y uint32) uint32 {
	return x<<29 | y<<26
}

func BCOND(x uint32, y uint32) uint32 {
	return x<<19 | y<<16
}

func MMU(x uint32, y uint32) uint32 {
	return SP(2, 0) | 16<<21 | x<<3 | y<<0
}

func FPF(x uint32, y uint32) uint32 {
	return SP(2, 1) | 16<<21 | x<<3 | y<<0
}

func FPD(x uint32, y uint32) uint32 {
	return SP(2, 1) | 17<<21 | x<<3 | y<<0
}

func FPW(x uint32, y uint32) uint32 {
	return SP(2, 1) | 20<<21 | x<<3 | y<<0
}

func FPV(x uint32, y uint32) uint32 {
	return SP(2, 1) | 21<<21 | x<<3 | y<<0
}

func OP_RRR(op uint32, r1 uint32, r2 uint32, r3 uint32) uint32 {
	return op | (r1&31)<<16 | (r2&31)<<21 | (r3&31)<<11
}

func OP_IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32 {
	return op | i&0xFFFF | (r2&31)<<21 | (r3&31)<<16
}

func OP_SRR(op uint32, s uint32, r2 uint32, r3 uint32) uint32 {
	return op | (s&31)<<6 | (r2&31)<<16 | (r3&31)<<11
}

func OP_FRRR(op uint32, r1 uint32, r2 uint32, r3 uint32) uint32 {
	return op | (r1&31)<<16 | (r2&31)<<11 | (r3&31)<<6
}

func OP_JMP(op uint32, i uint32) uint32 {
	return op | i&0x3FFFFFF
}

func (c *ctxt0) asmout(p *obj.Prog, o *Optab, out []uint32) {
	o1 := uint32(0)
	o2 := uint32(0)
	o3 := uint32(0)
	o4 := uint32(0)

	add := AADDU

	if c.ctxt.Arch.Family == sys.MIPS64 {
		add = AADDVU
	}
	switch o.type_ {
	default:
		c.ctxt.Diag("unknown type %d %v", o.type_)
		prasm(p)

	case 0:
		break

	case 1:
		a := AOR
		if p.As == AMOVW && c.ctxt.Arch.Family == sys.MIPS64 {
			a = AADDU
		}
		o1 = OP_RRR(c.oprrr(a), uint32(REGZERO), uint32(p.From.Reg), uint32(p.To.Reg))

	case 2:
		r := int(p.Reg)
		if p.As == ANEGW || p.As == ANEGV {
			r = REGZERO
		}
		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = OP_RRR(c.oprrr(p.As), uint32(p.From.Reg), uint32(r), uint32(p.To.Reg))

	case 3:
		v := c.regoff(&p.From)

		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		a := add
		if o.a1 == C_ANDCON {
			a = AOR
		}

		o1 = OP_IRR(c.opirr(a), uint32(v), uint32(r), uint32(p.To.Reg))

	case 4:
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}

		o1 = OP_IRR(c.opirr(p.As), uint32(v), uint32(r), uint32(p.To.Reg))

	case 5:
		o1 = c.oprrr(p.As)

	case 6:
		v := int32(0)
		if p.Pcond == nil {
			v = int32(-4) >> 2
		} else {
			v = int32(p.Pcond.Pc-p.Pc-4) >> 2
		}
		if (v<<16)>>16 != v {
			c.ctxt.Diag("short branch too far\n%v", p)
		}
		o1 = OP_IRR(c.opirr(p.As), uint32(v), uint32(p.From.Reg), uint32(p.Reg))

		o2 = 0

	case 7:
		r := int(p.To.Reg)
		if r == 0 {
			r = int(o.param)
		}
		v := c.regoff(&p.To)
		o1 = OP_IRR(c.opirr(p.As), uint32(v), uint32(r), uint32(p.From.Reg))

	case 8:
		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		v := c.regoff(&p.From)
		o1 = OP_IRR(c.opirr(-p.As), uint32(v), uint32(r), uint32(p.To.Reg))

	case 9:
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = OP_RRR(c.oprrr(p.As), uint32(r), uint32(p.From.Reg), uint32(p.To.Reg))

	case 10:
		v := c.regoff(&p.From)
		a := AOR
		if v < 0 {
			a = AADDU
		}
		o1 = OP_IRR(c.opirr(a), uint32(v), uint32(0), uint32(REGTMP))
		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		o2 = OP_RRR(c.oprrr(p.As), uint32(REGTMP), uint32(r), uint32(p.To.Reg))

	case 11:
		v := int32(0)
		if c.aclass(&p.To) == C_SBRA && p.To.Sym == nil && p.As == AJMP {

			if p.Pcond == nil {
				v = int32(-4) >> 2
			} else {
				v = int32(p.Pcond.Pc-p.Pc-4) >> 2
			}
			if (v<<16)>>16 == v {
				o1 = OP_IRR(c.opirr(ABEQ), uint32(v), uint32(REGZERO), uint32(REGZERO))
				break
			}
		}
		if p.Pcond == nil {
			v = int32(p.Pc) >> 2
		} else {
			v = int32(p.Pcond.Pc) >> 2
		}
		o1 = OP_JMP(c.opirr(p.As), uint32(v))
		if p.To.Sym == nil {
			p.To.Sym = c.cursym.Func.Text.From.Sym
			p.To.Offset = p.Pcond.Pc
		}
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 4
		rel.Sym = p.To.Sym
		rel.Add = p.To.Offset
		if p.As == AJAL {
			rel.Type = objabi.R_CALLMIPS
		} else {
			rel.Type = objabi.R_JMPMIPS
		}

	case 12:
		v := 16
		if p.As == AMOVB {
			v = 24
		}
		o1 = OP_SRR(c.opirr(ASLL), uint32(v), uint32(p.From.Reg), uint32(p.To.Reg))
		o2 = OP_SRR(c.opirr(ASRA), uint32(v), uint32(p.To.Reg), uint32(p.To.Reg))

	case 13:
		if p.As == AMOVBU {
			o1 = OP_IRR(c.opirr(AAND), uint32(0xff), uint32(p.From.Reg), uint32(p.To.Reg))
		} else {
			o1 = OP_IRR(c.opirr(AAND), uint32(0xffff), uint32(p.From.Reg), uint32(p.To.Reg))
		}

	case 14:
		o1 = OP_SRR(c.opirr(-ASLLV), uint32(0), uint32(p.From.Reg), uint32(p.To.Reg))
		o2 = OP_SRR(c.opirr(-ASRLV), uint32(0), uint32(p.To.Reg), uint32(p.To.Reg))

	case 15:
		v := c.regoff(&p.From)
		r := int(p.Reg)
		if r == 0 {
			r = REGZERO
		}

		o1 = OP_IRR(c.opirr(p.As), (uint32(v)&0x3FF)<<6, uint32(p.Reg), uint32(p.To.Reg))

	case 16:
		v := c.regoff(&p.From)
		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}

		if v >= 32 && vshift(p.As) {
			o1 = OP_SRR(c.opirr(-p.As), uint32(v-32), uint32(r), uint32(p.To.Reg))
		} else {
			o1 = OP_SRR(c.opirr(p.As), uint32(v), uint32(r), uint32(p.To.Reg))
		}

	case 17:
		o1 = OP_RRR(c.oprrr(p.As), uint32(REGZERO), uint32(p.From.Reg), uint32(p.To.Reg))

	case 18:
		r := int(p.Reg)
		if r == 0 {
			r = int(o.param)
		}
		o1 = OP_RRR(c.oprrr(p.As), uint32(0), uint32(p.To.Reg), uint32(r))
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 0
		rel.Type = objabi.R_CALLIND

	case 19:
		v := c.regoff(&p.From)
		o1 = OP_IRR(c.opirr(ALUI), uint32(v>>16), uint32(REGZERO), uint32(p.To.Reg))
		o2 = OP_IRR(c.opirr(AOR), uint32(v), uint32(p.To.Reg), uint32(p.To.Reg))

	case 20:
		a := OP(2, 0)
		if p.From.Reg == REG_LO {
			a = OP(2, 2)
		}
		o1 = OP_RRR(a, uint32(REGZERO), uint32(REGZERO), uint32(p.To.Reg))

	case 21:
		a := OP(2, 1)
		if p.To.Reg == REG_LO {
			a = OP(2, 3)
		}
		o1 = OP_RRR(a, uint32(REGZERO), uint32(p.From.Reg), uint32(REGZERO))

	case 22:
		if p.To.Reg != 0 {
			r := int(p.Reg)
			if r == 0 {
				r = int(p.To.Reg)
			}
			a := SP(3, 4) | 2
			o1 = OP_RRR(a, uint32(p.From.Reg), uint32(r), uint32(p.To.Reg))
		} else {
			o1 = OP_RRR(c.oprrr(p.As), uint32(p.From.Reg), uint32(p.Reg), uint32(REGZERO))
		}

	case 23:
		v := c.regoff(&p.From)
		o1 = OP_IRR(c.opirr(ALUI), uint32(v>>16), uint32(REGZERO), uint32(REGTMP))
		o2 = OP_IRR(c.opirr(AOR), uint32(v), uint32(REGTMP), uint32(REGTMP))
		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		o3 = OP_RRR(c.oprrr(p.As), uint32(REGTMP), uint32(r), uint32(p.To.Reg))

	case 24:
		v := c.regoff(&p.From)
		o1 = OP_IRR(c.opirr(ALUI), uint32(v>>16), uint32(REGZERO), uint32(p.To.Reg))

	case 25:
		v := c.regoff(&p.From)
		o1 = OP_IRR(c.opirr(ALUI), uint32(v>>16), uint32(REGZERO), uint32(REGTMP))
		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		o2 = OP_RRR(c.oprrr(p.As), uint32(REGTMP), uint32(r), uint32(p.To.Reg))

	case 26:
		v := c.regoff(&p.From)
		o1 = OP_IRR(c.opirr(ALUI), uint32(v>>16), uint32(REGZERO), uint32(REGTMP))
		o2 = OP_IRR(c.opirr(AOR), uint32(v), uint32(REGTMP), uint32(REGTMP))
		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		o3 = OP_RRR(c.oprrr(add), uint32(REGTMP), uint32(r), uint32(p.To.Reg))

	case 27:
		v := c.regoff(&p.From)
		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		a := -AMOVF
		if p.As == AMOVD {
			a = -AMOVD
		}
		switch o.size {
		case 12:
			o1 = OP_IRR(c.opirr(ALUI), uint32((v+1<<15)>>16), uint32(REGZERO), uint32(REGTMP))
			o2 = OP_RRR(c.oprrr(add), uint32(r), uint32(REGTMP), uint32(REGTMP))
			o3 = OP_IRR(c.opirr(a), uint32(v), uint32(REGTMP), uint32(p.To.Reg))

		case 4:
			o1 = OP_IRR(c.opirr(a), uint32(v), uint32(r), uint32(p.To.Reg))
		}

	case 28:
		v := c.regoff(&p.To)
		r := int(p.To.Reg)
		if r == 0 {
			r = int(o.param)
		}
		a := AMOVF
		if p.As == AMOVD {
			a = AMOVD
		}
		switch o.size {
		case 12:
			o1 = OP_IRR(c.opirr(ALUI), uint32((v+1<<15)>>16), uint32(REGZERO), uint32(REGTMP))
			o2 = OP_RRR(c.oprrr(add), uint32(r), uint32(REGTMP), uint32(REGTMP))
			o3 = OP_IRR(c.opirr(a), uint32(v), uint32(REGTMP), uint32(p.From.Reg))

		case 4:
			o1 = OP_IRR(c.opirr(a), uint32(v), uint32(r), uint32(p.From.Reg))
		}

	case 30:
		a := SP(2, 1) | (4 << 21)
		o1 = OP_RRR(a, uint32(p.From.Reg), uint32(0), uint32(p.To.Reg))

	case 31:
		a := SP(2, 1) | (0 << 21)
		o1 = OP_RRR(a, uint32(p.To.Reg), uint32(0), uint32(p.From.Reg))

	case 32:
		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = OP_FRRR(c.oprrr(p.As), uint32(p.From.Reg), uint32(r), uint32(p.To.Reg))

	case 33:
		o1 = OP_FRRR(c.oprrr(p.As), uint32(0), uint32(p.From.Reg), uint32(p.To.Reg))

	case 34:
		v := c.regoff(&p.From)
		a := AADDU
		if o.a1 == C_ANDCON {
			a = AOR
		}
		o1 = OP_IRR(c.opirr(a), uint32(v), uint32(0), uint32(REGTMP))
		o2 = OP_RRR(SP(2, 1)|(4<<21), uint32(REGTMP), uint32(0), uint32(p.To.Reg))

	case 35:
		v := c.regoff(&p.To)
		r := int(p.To.Reg)
		if r == 0 {
			r = int(o.param)
		}
		o1 = OP_IRR(c.opirr(ALUI), uint32((v+1<<15)>>16), uint32(REGZERO), uint32(REGTMP))
		o2 = OP_RRR(c.oprrr(add), uint32(r), uint32(REGTMP), uint32(REGTMP))
		o3 = OP_IRR(c.opirr(p.As), uint32(v), uint32(REGTMP), uint32(p.From.Reg))

	case 36:
		v := c.regoff(&p.From)
		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		o1 = OP_IRR(c.opirr(ALUI), uint32((v+1<<15)>>16), uint32(REGZERO), uint32(REGTMP))
		o2 = OP_RRR(c.oprrr(add), uint32(r), uint32(REGTMP), uint32(REGTMP))
		o3 = OP_IRR(c.opirr(-p.As), uint32(v), uint32(REGTMP), uint32(p.To.Reg))

	case 37:
		a := SP(2, 0) | (4 << 21)
		if p.As == AMOVV {
			a = SP(2, 0) | (5 << 21)
		}
		o1 = OP_RRR(a, uint32(p.From.Reg), uint32(0), uint32(p.To.Reg))

	case 38:
		a := SP(2, 0) | (0 << 21)
		if p.As == AMOVV {
			a = SP(2, 0) | (1 << 21)
		}
		o1 = OP_RRR(a, uint32(p.To.Reg), uint32(0), uint32(p.From.Reg))

	case 40:
		o1 = uint32(c.regoff(&p.From))

	case 41:
		o1 = OP_RRR(SP(2, 1)|(2<<21), uint32(REGZERO), uint32(0), uint32(p.To.Reg))
		o2 = OP_RRR(SP(2, 1)|(6<<21), uint32(p.From.Reg), uint32(0), uint32(p.To.Reg))

	case 42:
		o1 = OP_RRR(SP(2, 1)|(2<<21), uint32(p.To.Reg), uint32(0), uint32(p.From.Reg))

	case 47:
		a := SP(2, 1) | (5 << 21)
		o1 = OP_RRR(a, uint32(p.From.Reg), uint32(0), uint32(p.To.Reg))

	case 48:
		a := SP(2, 1) | (1 << 21)
		o1 = OP_RRR(a, uint32(p.To.Reg), uint32(0), uint32(p.From.Reg))

	case 49:
		o1 = 52

	case 50:
		o1 = OP_IRR(c.opirr(ALUI), uint32(0), uint32(REGZERO), uint32(REGTMP))
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 4
		rel.Sym = p.To.Sym
		rel.Add = p.To.Offset
		rel.Type = objabi.R_ADDRMIPSU
		o2 = OP_IRR(c.opirr(p.As), uint32(0), uint32(REGTMP), uint32(p.From.Reg))
		rel2 := obj.Addrel(c.cursym)
		rel2.Off = int32(c.pc + 4)
		rel2.Siz = 4
		rel2.Sym = p.To.Sym
		rel2.Add = p.To.Offset
		rel2.Type = objabi.R_ADDRMIPS

		if o.size == 12 {
			o3 = o2
			o2 = OP_RRR(c.oprrr(AADDVU), uint32(REGSB), uint32(REGTMP), uint32(REGTMP))
			rel2.Off += 4
		}

	case 51:
		o1 = OP_IRR(c.opirr(ALUI), uint32(0), uint32(REGZERO), uint32(REGTMP))
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 4
		rel.Sym = p.From.Sym
		rel.Add = p.From.Offset
		rel.Type = objabi.R_ADDRMIPSU
		o2 = OP_IRR(c.opirr(-p.As), uint32(0), uint32(REGTMP), uint32(p.To.Reg))
		rel2 := obj.Addrel(c.cursym)
		rel2.Off = int32(c.pc + 4)
		rel2.Siz = 4
		rel2.Sym = p.From.Sym
		rel2.Add = p.From.Offset
		rel2.Type = objabi.R_ADDRMIPS

		if o.size == 12 {
			o3 = o2
			o2 = OP_RRR(c.oprrr(AADDVU), uint32(REGSB), uint32(REGTMP), uint32(REGTMP))
			rel2.Off += 4
		}

	case 52:
		o1 = OP_IRR(c.opirr(ALUI), uint32(0), uint32(REGZERO), uint32(p.To.Reg))
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 4
		rel.Sym = p.From.Sym
		rel.Add = p.From.Offset
		rel.Type = objabi.R_ADDRMIPSU
		o2 = OP_IRR(c.opirr(add), uint32(0), uint32(p.To.Reg), uint32(p.To.Reg))
		rel2 := obj.Addrel(c.cursym)
		rel2.Off = int32(c.pc + 4)
		rel2.Siz = 4
		rel2.Sym = p.From.Sym
		rel2.Add = p.From.Offset
		rel2.Type = objabi.R_ADDRMIPS

		if o.size == 12 {
			o3 = o2
			o2 = OP_RRR(c.oprrr(AADDVU), uint32(REGSB), uint32(p.To.Reg), uint32(p.To.Reg))
			rel2.Off += 4
		}

	case 53:

		o1 = (037<<26 + 073) | (29 << 11) | (3 << 16)
		o2 = OP_IRR(c.opirr(p.As), uint32(0), uint32(REG_R3), uint32(p.From.Reg))
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc + 4)
		rel.Siz = 4
		rel.Sym = p.To.Sym
		rel.Add = p.To.Offset
		rel.Type = objabi.R_ADDRMIPSTLS

	case 54:

		o1 = (037<<26 + 073) | (29 << 11) | (3 << 16)
		o2 = OP_IRR(c.opirr(-p.As), uint32(0), uint32(REG_R3), uint32(p.To.Reg))
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc + 4)
		rel.Siz = 4
		rel.Sym = p.From.Sym
		rel.Add = p.From.Offset
		rel.Type = objabi.R_ADDRMIPSTLS

	case 55:

		o1 = (037<<26 + 073) | (29 << 11) | (3 << 16)
		o2 = OP_IRR(c.opirr(add), uint32(0), uint32(REG_R3), uint32(p.To.Reg))
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc + 4)
		rel.Siz = 4
		rel.Sym = p.From.Sym
		rel.Add = p.From.Offset
		rel.Type = objabi.R_ADDRMIPSTLS
	}

	out[0] = o1
	out[1] = o2
	out[2] = o3
	out[3] = o4
}

func (c *ctxt0) vregoff(a *obj.Addr) int64 {
	c.instoffset = 0
	c.aclass(a)
	return c.instoffset
}

func (c *ctxt0) regoff(a *obj.Addr) int32 {
	return int32(c.vregoff(a))
}

func (c *ctxt0) oprrr(a obj.As) uint32 {
	switch a {
	case AADD:
		return OP(4, 0)
	case AADDU:
		return OP(4, 1)
	case ASGT:
		return OP(5, 2)
	case ASGTU:
		return OP(5, 3)
	case AAND:
		return OP(4, 4)
	case AOR:
		return OP(4, 5)
	case AXOR:
		return OP(4, 6)
	case ASUB:
		return OP(4, 2)
	case ASUBU, ANEGW:
		return OP(4, 3)
	case ANOR:
		return OP(4, 7)
	case ASLL:
		return OP(0, 4)
	case ASRL:
		return OP(0, 6)
	case ASRA:
		return OP(0, 7)
	case ASLLV:
		return OP(2, 4)
	case ASRLV:
		return OP(2, 6)
	case ASRAV:
		return OP(2, 7)
	case AADDV:
		return OP(5, 4)
	case AADDVU:
		return OP(5, 5)
	case ASUBV:
		return OP(5, 6)
	case ASUBVU, ANEGV:
		return OP(5, 7)
	case AREM,
		ADIV:
		return OP(3, 2)
	case AREMU,
		ADIVU:
		return OP(3, 3)
	case AMUL:
		return OP(3, 0)
	case AMULU:
		return OP(3, 1)
	case AREMV,
		ADIVV:
		return OP(3, 6)
	case AREMVU,
		ADIVVU:
		return OP(3, 7)
	case AMULV:
		return OP(3, 4)
	case AMULVU:
		return OP(3, 5)

	case AJMP:
		return OP(1, 0)
	case AJAL:
		return OP(1, 1)

	case ABREAK:
		return OP(1, 5)
	case ASYSCALL:
		return OP(1, 4)
	case ATLBP:
		return MMU(1, 0)
	case ATLBR:
		return MMU(0, 1)
	case ATLBWI:
		return MMU(0, 2)
	case ATLBWR:
		return MMU(0, 6)
	case ARFE:
		return MMU(2, 0)

	case ADIVF:
		return FPF(0, 3)
	case ADIVD:
		return FPD(0, 3)
	case AMULF:
		return FPF(0, 2)
	case AMULD:
		return FPD(0, 2)
	case ASUBF:
		return FPF(0, 1)
	case ASUBD:
		return FPD(0, 1)
	case AADDF:
		return FPF(0, 0)
	case AADDD:
		return FPD(0, 0)
	case ATRUNCFV:
		return FPF(1, 1)
	case ATRUNCDV:
		return FPD(1, 1)
	case ATRUNCFW:
		return FPF(1, 5)
	case ATRUNCDW:
		return FPD(1, 5)
	case AMOVFV:
		return FPF(4, 5)
	case AMOVDV:
		return FPD(4, 5)
	case AMOVVF:
		return FPV(4, 0)
	case AMOVVD:
		return FPV(4, 1)
	case AMOVFW:
		return FPF(4, 4)
	case AMOVDW:
		return FPD(4, 4)
	case AMOVWF:
		return FPW(4, 0)
	case AMOVDF:
		return FPD(4, 0)
	case AMOVWD:
		return FPW(4, 1)
	case AMOVFD:
		return FPF(4, 1)
	case AABSF:
		return FPF(0, 5)
	case AABSD:
		return FPD(0, 5)
	case AMOVF:
		return FPF(0, 6)
	case AMOVD:
		return FPD(0, 6)
	case ANEGF:
		return FPF(0, 7)
	case ANEGD:
		return FPD(0, 7)
	case ACMPEQF:
		return FPF(6, 2)
	case ACMPEQD:
		return FPD(6, 2)
	case ACMPGTF:
		return FPF(7, 4)
	case ACMPGTD:
		return FPD(7, 4)
	case ACMPGEF:
		return FPF(7, 6)
	case ACMPGED:
		return FPD(7, 6)

	case ASQRTF:
		return FPF(0, 4)
	case ASQRTD:
		return FPD(0, 4)

	case ASYNC:
		return OP(1, 7)
	case ANOOP:
		return 0

	case ACMOVN:
		return OP(1, 3)
	case ACMOVZ:
		return OP(1, 2)
	case ACMOVT:
		return OP(0, 1) | (1 << 16)
	case ACMOVF:
		return OP(0, 1) | (0 << 16)
	case ACLO:
		return SP(3, 4) | OP(4, 1)
	case ACLZ:
		return SP(3, 4) | OP(4, 0)
	}

	if a < 0 {
		c.ctxt.Diag("bad rrr opcode -%v", -a)
	} else {
		c.ctxt.Diag("bad rrr opcode %v", a)
	}
	return 0
}

func (c *ctxt0) opirr(a obj.As) uint32 {
	switch a {
	case AADD:
		return SP(1, 0)
	case AADDU:
		return SP(1, 1)
	case ASGT:
		return SP(1, 2)
	case ASGTU:
		return SP(1, 3)
	case AAND:
		return SP(1, 4)
	case AOR:
		return SP(1, 5)
	case AXOR:
		return SP(1, 6)
	case ALUI:
		return SP(1, 7)
	case ASLL:
		return OP(0, 0)
	case ASRL:
		return OP(0, 2)
	case ASRA:
		return OP(0, 3)
	case AADDV:
		return SP(3, 0)
	case AADDVU:
		return SP(3, 1)

	case AJMP:
		return SP(0, 2)
	case AJAL,
		obj.ADUFFZERO,
		obj.ADUFFCOPY:
		return SP(0, 3)
	case ABEQ:
		return SP(0, 4)
	case -ABEQ:
		return SP(2, 4)
	case ABNE:
		return SP(0, 5)
	case -ABNE:
		return SP(2, 5)
	case ABGEZ:
		return SP(0, 1) | BCOND(0, 1)
	case -ABGEZ:
		return SP(0, 1) | BCOND(0, 3)
	case ABGEZAL:
		return SP(0, 1) | BCOND(2, 1)
	case -ABGEZAL:
		return SP(0, 1) | BCOND(2, 3)
	case ABGTZ:
		return SP(0, 7)
	case -ABGTZ:
		return SP(2, 7)
	case ABLEZ:
		return SP(0, 6)
	case -ABLEZ:
		return SP(2, 6)
	case ABLTZ:
		return SP(0, 1) | BCOND(0, 0)
	case -ABLTZ:
		return SP(0, 1) | BCOND(0, 2)
	case ABLTZAL:
		return SP(0, 1) | BCOND(2, 0)
	case -ABLTZAL:
		return SP(0, 1) | BCOND(2, 2)
	case ABFPT:
		return SP(2, 1) | (257 << 16)
	case -ABFPT:
		return SP(2, 1) | (259 << 16)
	case ABFPF:
		return SP(2, 1) | (256 << 16)
	case -ABFPF:
		return SP(2, 1) | (258 << 16)

	case AMOVB,
		AMOVBU:
		return SP(5, 0)
	case AMOVH,
		AMOVHU:
		return SP(5, 1)
	case AMOVW,
		AMOVWU:
		return SP(5, 3)
	case AMOVV:
		return SP(7, 7)
	case AMOVF:
		return SP(7, 1)
	case AMOVD:
		return SP(7, 5)
	case AMOVWL:
		return SP(5, 2)
	case AMOVWR:
		return SP(5, 6)
	case AMOVVL:
		return SP(5, 4)
	case AMOVVR:
		return SP(5, 5)

	case ABREAK:
		return SP(5, 7)

	case -AMOVWL:
		return SP(4, 2)
	case -AMOVWR:
		return SP(4, 6)
	case -AMOVVL:
		return SP(3, 2)
	case -AMOVVR:
		return SP(3, 3)
	case -AMOVB:
		return SP(4, 0)
	case -AMOVBU:
		return SP(4, 4)
	case -AMOVH:
		return SP(4, 1)
	case -AMOVHU:
		return SP(4, 5)
	case -AMOVW:
		return SP(4, 3)
	case -AMOVWU:
		return SP(4, 7)
	case -AMOVV:
		return SP(6, 7)
	case -AMOVF:
		return SP(6, 1)
	case -AMOVD:
		return SP(6, 5)

	case ASLLV:
		return OP(7, 0)
	case ASRLV:
		return OP(7, 2)
	case ASRAV:
		return OP(7, 3)
	case -ASLLV:
		return OP(7, 4)
	case -ASRLV:
		return OP(7, 6)
	case -ASRAV:
		return OP(7, 7)

	case ATEQ:
		return OP(6, 4)
	case ATNE:
		return OP(6, 6)
	case -ALL:
		return SP(6, 0)
	case -ALLV:
		return SP(6, 4)
	case ASC:
		return SP(7, 0)
	case ASCV:
		return SP(7, 4)
	}

	if a < 0 {
		c.ctxt.Diag("bad irr opcode -%v", -a)
	} else {
		c.ctxt.Diag("bad irr opcode %v", a)
	}
	return 0
}

func vshift(a obj.As) bool {
	switch a {
	case ASLLV,
		ASRLV,
		ASRAV:
		return true
	}
	return false
}

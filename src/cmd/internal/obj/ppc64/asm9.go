package ppc64

import (
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"log"
	"sort"
)

// ctxt9 holds state while assembling a single function.
// Each function gets a fresh ctxt9.
// This allows for multiple functions to be safely concurrently assembled.
type ctxt9 struct {
	ctxt       *obj.Link
	newprog    obj.ProgAlloc
	cursym     *obj.LSym
	autosize   int32
	instoffset int64
	pc         int64
}

const (
	funcAlign = 16
)

const (
	r0iszero = 1
)

type Optab struct {
	as    obj.As // Opcode
	a1    uint8
	a2    uint8
	a3    uint8
	a4    uint8
	type_ int8 // cases in asmout below. E.g., 44 = st r,(ra+rb); 45 = ld (ra+rb), r
	size  int8
	param int16
}

// This optab contains a list of opcodes with the operand
// combinations that are implemented. Not all opcodes are in this
// table, but are added later in buildop by calling opset for those
// opcodes which allow the same operand combinations as an opcode
// already in the table.
//
// The type field in the Optabl identifies the case in asmout where
// the instruction word is assembled.

func (psess *PackageSession) span9(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	p := cursym.Func.Text
	if p == nil || p.Link == nil {
		return
	}

	if psess.oprange[AANDN&obj.AMask] == nil {
		ctxt.Diag("ppc64 ops not initialized, call ppc64.buildop first")
	}

	c := ctxt9{ctxt: ctxt, newprog: newprog, cursym: cursym, autosize: int32(p.To.Offset)}

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

			if (o.type_ == 16 || o.type_ == 17) && p.Pcond != nil {
				otxt = p.Pcond.Pc - pc
				if otxt < -(1<<15)+10 || otxt >= (1<<15)-10 {
					q = c.newprog()
					q.Link = p.Link
					p.Link = q
					q.As = ABR
					q.To.Type = obj.TYPE_BRANCH
					q.Pcond = p.Pcond
					p.Pcond = q
					q = c.newprog()
					q.Link = p.Link
					p.Link = q
					q.As = ABR
					q.To.Type = obj.TYPE_BRANCH
					q.Pcond = q.Link.Link

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

	pc += -pc & (funcAlign - 1)
	c.cursym.Size = pc

	c.cursym.Grow(c.cursym.Size)

	bp := c.cursym.P
	var i int32
	var out [6]uint32
	for p := c.cursym.Func.Text.Link; p != nil; p = p.Link {
		c.pc = p.Pc
		o = c.oplook(psess, p)
		if int(o.size) > 4*len(out) {
			log.Fatalf("out array in span9 is too small, need at least %d for %v", o.size/4, p)
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

func (c *ctxt9) aclass(a *obj.Addr) int {
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
		if REG_V0 <= a.Reg && a.Reg <= REG_V31 {
			return C_VREG
		}
		if REG_VS0 <= a.Reg && a.Reg <= REG_VS63 {
			return C_VSREG
		}
		if REG_CR0 <= a.Reg && a.Reg <= REG_CR7 || a.Reg == REG_CR {
			return C_CREG
		}
		if REG_SPR0 <= a.Reg && a.Reg <= REG_SPR0+1023 {
			switch a.Reg {
			case REG_LR:
				return C_LR

			case REG_XER:
				return C_XER

			case REG_CTR:
				return C_CTR
			}

			return C_SPR
		}

		if REG_DCR0 <= a.Reg && a.Reg <= REG_DCR0+1023 {
			return C_SPR
		}
		if a.Reg == REG_FPSCR {
			return C_FPSCR
		}
		if a.Reg == REG_MSR {
			return C_MSR
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
					if c.ctxt.Flag_shared {
						return C_TLS_IE
					} else {
						return C_TLS_LE
					}
				}
				return C_ADDR
			}
			return C_LEXT

		case obj.NAME_GOTREF:
			return C_GOTADDR

		case obj.NAME_AUTO:
			c.instoffset = int64(c.autosize) + a.Offset
			if c.instoffset >= -BIG && c.instoffset < BIG {
				return C_SAUTO
			}
			return C_LAUTO

		case obj.NAME_PARAM:
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

			return C_LCON

		case obj.NAME_AUTO:
			c.instoffset = int64(c.autosize) + a.Offset
			if c.instoffset >= -BIG && c.instoffset < BIG {
				return C_SACON
			}
			return C_LACON

		case obj.NAME_PARAM:
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
			return C_DCON
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
		return C_DCON

	case obj.TYPE_BRANCH:
		if a.Sym != nil && c.ctxt.Flag_dynlink {
			return C_LBRAPIC
		}
		return C_SBRA
	}

	return C_GOK
}

func prasm(p *obj.Prog) {
	fmt.Printf("%v\n", p)
}

func (c *ctxt9) oplook(psess *PackageSession, p *obj.Prog) *Optab {
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
	a3 := C_NONE + 1
	if p.GetFrom3() != nil {
		a3 = int(p.GetFrom3().Class)
		if a3 == 0 {
			a3 = c.aclass(p.GetFrom3()) + 1
			p.GetFrom3().Class = int8(a3)
		}
	}

	a3--
	a4 := int(p.To.Class)
	if a4 == 0 {
		a4 = c.aclass(&p.To) + 1
		p.To.Class = int8(a4)
	}

	a4--
	a2 := C_NONE
	if p.Reg != 0 {
		if REG_R0 <= p.Reg && p.Reg <= REG_R31 {
			a2 = C_REG
		} else if REG_V0 <= p.Reg && p.Reg <= REG_V31 {
			a2 = C_VREG
		} else if REG_VS0 <= p.Reg && p.Reg <= REG_VS63 {
			a2 = C_VSREG
		} else if REG_F0 <= p.Reg && p.Reg <= REG_F31 {
			a2 = C_FREG
		}
	}

	ops := psess.oprange[p.As&obj.AMask]
	c1 := &psess.xcmp[a1]
	c3 := &psess.xcmp[a3]
	c4 := &psess.xcmp[a4]
	for i := range ops {
		op := &ops[i]
		if int(op.a2) == a2 && c1[op.a1] && c3[op.a3] && c4[op.a4] {
			p.Optab = uint16(cap(psess.optab) - cap(ops) + i + 1)
			return op
		}
	}

	c.ctxt.Diag("illegal combination %v %v %v %v %v", p.As, psess.DRconv(a1), psess.DRconv(a2), psess.DRconv(a3), psess.DRconv(a4))
	prasm(p)
	if ops == nil {
		ops = psess.optab
	}
	return &ops[0]
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

	case C_ADDCON:
		if b == C_ZCON || b == C_SCON {
			return true
		}

	case C_ANDCON:
		if b == C_ZCON || b == C_SCON {
			return true
		}

	case C_SPR:
		if b == C_LR || b == C_XER || b == C_CTR {
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

	case C_ANY:
		return true
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

// Used when sorting the optab. Sorting is
// done in a way so that the best choice of
// opcode/operand combination is considered first.
func (x ocmp) Less(i, j int) bool {
	p1 := &x[i]
	p2 := &x[j]
	n := int(p1.as) - int(p2.as)

	if n != 0 {
		return n < 0
	}

	n = int(p1.size) - int(p2.size)
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
	n = int(p1.a4) - int(p2.a4)
	if n != 0 {
		return n < 0
	}
	return false
}

// Add an entry to the opcode table for
// a new opcode b0 with the same operand combinations
// as opcode a.
func (psess *PackageSession) opset(a, b0 obj.As) {
	psess.
		oprange[a&obj.AMask] = psess.oprange[b0]
}

// Build the opcode table
func (psess *PackageSession) buildop(ctxt *obj.Link) {
	if psess.oprange[AANDN&obj.AMask] != nil {

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
			log.Fatalf("instruction missing from switch in asm9.go:buildop: %v", r)

		case ADCBF:
			psess.
				opset(ADCBI, r0)
			psess.
				opset(ADCBST, r0)
			psess.
				opset(ADCBT, r0)
			psess.
				opset(ADCBTST, r0)
			psess.
				opset(ADCBZ, r0)
			psess.
				opset(AICBI, r0)

		case AECOWX:
			psess.
				opset(ASTWCCC, r0)
			psess.
				opset(ASTBCCC, r0)
			psess.
				opset(ASTDCCC, r0)

		case AREM:
			psess.
				opset(AREMCC, r0)
			psess.
				opset(AREMV, r0)
			psess.
				opset(AREMVCC, r0)

		case AREMU:
			psess.
				opset(AREMU, r0)
			psess.
				opset(AREMUCC, r0)
			psess.
				opset(AREMUV, r0)
			psess.
				opset(AREMUVCC, r0)

		case AREMD:
			psess.
				opset(AREMDCC, r0)
			psess.
				opset(AREMDV, r0)
			psess.
				opset(AREMDVCC, r0)

		case AREMDU:
			psess.
				opset(AREMDU, r0)
			psess.
				opset(AREMDUCC, r0)
			psess.
				opset(AREMDUV, r0)
			psess.
				opset(AREMDUVCC, r0)

		case ADIVW:
			psess.
				opset(AMULHW, r0)
			psess.
				opset(AMULHWCC, r0)
			psess.
				opset(AMULHWU, r0)
			psess.
				opset(AMULHWUCC, r0)
			psess.
				opset(AMULLWCC, r0)
			psess.
				opset(AMULLWVCC, r0)
			psess.
				opset(AMULLWV, r0)
			psess.
				opset(ADIVWCC, r0)
			psess.
				opset(ADIVWV, r0)
			psess.
				opset(ADIVWVCC, r0)
			psess.
				opset(ADIVWU, r0)
			psess.
				opset(ADIVWUCC, r0)
			psess.
				opset(ADIVWUV, r0)
			psess.
				opset(ADIVWUVCC, r0)
			psess.
				opset(AADDCC, r0)
			psess.
				opset(AADDCV, r0)
			psess.
				opset(AADDCVCC, r0)
			psess.
				opset(AADDV, r0)
			psess.
				opset(AADDVCC, r0)
			psess.
				opset(AADDE, r0)
			psess.
				opset(AADDECC, r0)
			psess.
				opset(AADDEV, r0)
			psess.
				opset(AADDEVCC, r0)
			psess.
				opset(ACRAND, r0)
			psess.
				opset(ACRANDN, r0)
			psess.
				opset(ACREQV, r0)
			psess.
				opset(ACRNAND, r0)
			psess.
				opset(ACRNOR, r0)
			psess.
				opset(ACROR, r0)
			psess.
				opset(ACRORN, r0)
			psess.
				opset(ACRXOR, r0)
			psess.
				opset(AMULHD, r0)
			psess.
				opset(AMULHDCC, r0)
			psess.
				opset(AMULHDU, r0)
			psess.
				opset(AMULHDUCC, r0)
			psess.
				opset(AMULLD, r0)
			psess.
				opset(AMULLDCC, r0)
			psess.
				opset(AMULLDVCC, r0)
			psess.
				opset(AMULLDV, r0)
			psess.
				opset(ADIVD, r0)
			psess.
				opset(ADIVDCC, r0)
			psess.
				opset(ADIVDE, r0)
			psess.
				opset(ADIVDEU, r0)
			psess.
				opset(ADIVDECC, r0)
			psess.
				opset(ADIVDEUCC, r0)
			psess.
				opset(ADIVDVCC, r0)
			psess.
				opset(ADIVDV, r0)
			psess.
				opset(ADIVDU, r0)
			psess.
				opset(ADIVDUCC, r0)
			psess.
				opset(ADIVDUVCC, r0)
			psess.
				opset(ADIVDUCC, r0)

		case APOPCNTD:
			psess.
				opset(APOPCNTW, r0)
			psess.
				opset(APOPCNTB, r0)

		case ACOPY:
			psess.
				opset(APASTECC, r0)

		case AMADDHD:
			psess.
				opset(AMADDHDU, r0)
			psess.
				opset(AMADDLD, r0)

		case AMOVBZ:
			psess.
				opset(AMOVH, r0)
			psess.
				opset(AMOVHZ, r0)

		case AMOVBZU:
			psess.
				opset(AMOVHU, r0)
			psess.
				opset(AMOVHZU, r0)
			psess.
				opset(AMOVWU, r0)
			psess.
				opset(AMOVWZU, r0)
			psess.
				opset(AMOVDU, r0)
			psess.
				opset(AMOVMW, r0)

		case ALV:
			psess.
				opset(ALVEBX, r0)
			psess.
				opset(ALVEHX, r0)
			psess.
				opset(ALVEWX, r0)
			psess.
				opset(ALVX, r0)
			psess.
				opset(ALVXL, r0)
			psess.
				opset(ALVSL, r0)
			psess.
				opset(ALVSR, r0)

		case ASTV:
			psess.
				opset(ASTVEBX, r0)
			psess.
				opset(ASTVEHX, r0)
			psess.
				opset(ASTVEWX, r0)
			psess.
				opset(ASTVX, r0)
			psess.
				opset(ASTVXL, r0)

		case AVAND:
			psess.
				opset(AVAND, r0)
			psess.
				opset(AVANDC, r0)
			psess.
				opset(AVNAND, r0)

		case AVOR:
			psess.
				opset(AVOR, r0)
			psess.
				opset(AVORC, r0)
			psess.
				opset(AVXOR, r0)
			psess.
				opset(AVNOR, r0)
			psess.
				opset(AVEQV, r0)

		case AVADDUM:
			psess.
				opset(AVADDUBM, r0)
			psess.
				opset(AVADDUHM, r0)
			psess.
				opset(AVADDUWM, r0)
			psess.
				opset(AVADDUDM, r0)
			psess.
				opset(AVADDUQM, r0)

		case AVADDCU:
			psess.
				opset(AVADDCUQ, r0)
			psess.
				opset(AVADDCUW, r0)

		case AVADDUS:
			psess.
				opset(AVADDUBS, r0)
			psess.
				opset(AVADDUHS, r0)
			psess.
				opset(AVADDUWS, r0)

		case AVADDSS:
			psess.
				opset(AVADDSBS, r0)
			psess.
				opset(AVADDSHS, r0)
			psess.
				opset(AVADDSWS, r0)

		case AVADDE:
			psess.
				opset(AVADDEUQM, r0)
			psess.
				opset(AVADDECUQ, r0)

		case AVSUBUM:
			psess.
				opset(AVSUBUBM, r0)
			psess.
				opset(AVSUBUHM, r0)
			psess.
				opset(AVSUBUWM, r0)
			psess.
				opset(AVSUBUDM, r0)
			psess.
				opset(AVSUBUQM, r0)

		case AVSUBCU:
			psess.
				opset(AVSUBCUQ, r0)
			psess.
				opset(AVSUBCUW, r0)

		case AVSUBUS:
			psess.
				opset(AVSUBUBS, r0)
			psess.
				opset(AVSUBUHS, r0)
			psess.
				opset(AVSUBUWS, r0)

		case AVSUBSS:
			psess.
				opset(AVSUBSBS, r0)
			psess.
				opset(AVSUBSHS, r0)
			psess.
				opset(AVSUBSWS, r0)

		case AVSUBE:
			psess.
				opset(AVSUBEUQM, r0)
			psess.
				opset(AVSUBECUQ, r0)

		case AVMULESB:
			psess.
				opset(AVMULOSB, r0)
			psess.
				opset(AVMULEUB, r0)
			psess.
				opset(AVMULOUB, r0)
			psess.
				opset(AVMULESH, r0)
			psess.
				opset(AVMULOSH, r0)
			psess.
				opset(AVMULEUH, r0)
			psess.
				opset(AVMULOUH, r0)
			psess.
				opset(AVMULESW, r0)
			psess.
				opset(AVMULOSW, r0)
			psess.
				opset(AVMULEUW, r0)
			psess.
				opset(AVMULOUW, r0)
			psess.
				opset(AVMULUWM, r0)
		case AVPMSUM:
			psess.
				opset(AVPMSUMB, r0)
			psess.
				opset(AVPMSUMH, r0)
			psess.
				opset(AVPMSUMW, r0)
			psess.
				opset(AVPMSUMD, r0)

		case AVR:
			psess.
				opset(AVRLB, r0)
			psess.
				opset(AVRLH, r0)
			psess.
				opset(AVRLW, r0)
			psess.
				opset(AVRLD, r0)

		case AVS:
			psess.
				opset(AVSLB, r0)
			psess.
				opset(AVSLH, r0)
			psess.
				opset(AVSLW, r0)
			psess.
				opset(AVSL, r0)
			psess.
				opset(AVSLO, r0)
			psess.
				opset(AVSRB, r0)
			psess.
				opset(AVSRH, r0)
			psess.
				opset(AVSRW, r0)
			psess.
				opset(AVSR, r0)
			psess.
				opset(AVSRO, r0)
			psess.
				opset(AVSLD, r0)
			psess.
				opset(AVSRD, r0)

		case AVSA:
			psess.
				opset(AVSRAB, r0)
			psess.
				opset(AVSRAH, r0)
			psess.
				opset(AVSRAW, r0)
			psess.
				opset(AVSRAD, r0)

		case AVSOI:
			psess.
				opset(AVSLDOI, r0)

		case AVCLZ:
			psess.
				opset(AVCLZB, r0)
			psess.
				opset(AVCLZH, r0)
			psess.
				opset(AVCLZW, r0)
			psess.
				opset(AVCLZD, r0)

		case AVPOPCNT:
			psess.
				opset(AVPOPCNTB, r0)
			psess.
				opset(AVPOPCNTH, r0)
			psess.
				opset(AVPOPCNTW, r0)
			psess.
				opset(AVPOPCNTD, r0)

		case AVCMPEQ:
			psess.
				opset(AVCMPEQUB, r0)
			psess.
				opset(AVCMPEQUBCC, r0)
			psess.
				opset(AVCMPEQUH, r0)
			psess.
				opset(AVCMPEQUHCC, r0)
			psess.
				opset(AVCMPEQUW, r0)
			psess.
				opset(AVCMPEQUWCC, r0)
			psess.
				opset(AVCMPEQUD, r0)
			psess.
				opset(AVCMPEQUDCC, r0)

		case AVCMPGT:
			psess.
				opset(AVCMPGTUB, r0)
			psess.
				opset(AVCMPGTUBCC, r0)
			psess.
				opset(AVCMPGTUH, r0)
			psess.
				opset(AVCMPGTUHCC, r0)
			psess.
				opset(AVCMPGTUW, r0)
			psess.
				opset(AVCMPGTUWCC, r0)
			psess.
				opset(AVCMPGTUD, r0)
			psess.
				opset(AVCMPGTUDCC, r0)
			psess.
				opset(AVCMPGTSB, r0)
			psess.
				opset(AVCMPGTSBCC, r0)
			psess.
				opset(AVCMPGTSH, r0)
			psess.
				opset(AVCMPGTSHCC, r0)
			psess.
				opset(AVCMPGTSW, r0)
			psess.
				opset(AVCMPGTSWCC, r0)
			psess.
				opset(AVCMPGTSD, r0)
			psess.
				opset(AVCMPGTSDCC, r0)

		case AVCMPNEZB:
			psess.
				opset(AVCMPNEZBCC, r0)

		case AVPERM:
			psess.
				opset(AVPERM, r0)

		case AVBPERMQ:
			psess.
				opset(AVBPERMD, r0)

		case AVSEL:
			psess.
				opset(AVSEL, r0)

		case AVSPLT:
			psess.
				opset(AVSPLTB, r0)
			psess.
				opset(AVSPLTH, r0)
			psess.
				opset(AVSPLTW, r0)

		case AVSPLTI:
			psess.
				opset(AVSPLTISB, r0)
			psess.
				opset(AVSPLTISH, r0)
			psess.
				opset(AVSPLTISW, r0)

		case AVCIPH:
			psess.
				opset(AVCIPHER, r0)
			psess.
				opset(AVCIPHERLAST, r0)

		case AVNCIPH:
			psess.
				opset(AVNCIPHER, r0)
			psess.
				opset(AVNCIPHERLAST, r0)

		case AVSBOX:
			psess.
				opset(AVSBOX, r0)

		case AVSHASIGMA:
			psess.
				opset(AVSHASIGMAW, r0)
			psess.
				opset(AVSHASIGMAD, r0)

		case ALXV:
			psess.
				opset(ALXVD2X, r0)
			psess.
				opset(ALXVDSX, r0)
			psess.
				opset(ALXVW4X, r0)

		case ASTXV:
			psess.
				opset(ASTXVD2X, r0)
			psess.
				opset(ASTXVW4X, r0)

		case ALXS:
			psess.
				opset(ALXSDX, r0)

		case ASTXS:
			psess.
				opset(ASTXSDX, r0)

		case ALXSI:
			psess.
				opset(ALXSIWAX, r0)
			psess.
				opset(ALXSIWZX, r0)

		case ASTXSI:
			psess.
				opset(ASTXSIWX, r0)

		case AMFVSR:
			psess.
				opset(AMFVSRD, r0)
			psess.
				opset(AMFFPRD, r0)
			psess.
				opset(AMFVRD, r0)
			psess.
				opset(AMFVSRWZ, r0)
			psess.
				opset(AMFVSRLD, r0)

		case AMTVSR:
			psess.
				opset(AMTVSRD, r0)
			psess.
				opset(AMTFPRD, r0)
			psess.
				opset(AMTVRD, r0)
			psess.
				opset(AMTVSRWA, r0)
			psess.
				opset(AMTVSRWZ, r0)
			psess.
				opset(AMTVSRDD, r0)
			psess.
				opset(AMTVSRWS, r0)

		case AXXLAND:
			psess.
				opset(AXXLANDQ, r0)
			psess.
				opset(AXXLANDC, r0)
			psess.
				opset(AXXLEQV, r0)
			psess.
				opset(AXXLNAND, r0)

		case AXXLOR:
			psess.
				opset(AXXLORC, r0)
			psess.
				opset(AXXLNOR, r0)
			psess.
				opset(AXXLORQ, r0)
			psess.
				opset(AXXLXOR, r0)

		case AXXSEL:
			psess.
				opset(AXXSEL, r0)

		case AXXMRG:
			psess.
				opset(AXXMRGHW, r0)
			psess.
				opset(AXXMRGLW, r0)

		case AXXSPLT:
			psess.
				opset(AXXSPLTW, r0)

		case AXXPERM:
			psess.
				opset(AXXPERMDI, r0)

		case AXXSI:
			psess.
				opset(AXXSLDWI, r0)

		case AXSCV:
			psess.
				opset(AXSCVDPSP, r0)
			psess.
				opset(AXSCVSPDP, r0)
			psess.
				opset(AXSCVDPSPN, r0)
			psess.
				opset(AXSCVSPDPN, r0)

		case AXVCV:
			psess.
				opset(AXVCVDPSP, r0)
			psess.
				opset(AXVCVSPDP, r0)

		case AXSCVX:
			psess.
				opset(AXSCVDPSXDS, r0)
			psess.
				opset(AXSCVDPSXWS, r0)
			psess.
				opset(AXSCVDPUXDS, r0)
			psess.
				opset(AXSCVDPUXWS, r0)

		case AXSCVXP:
			psess.
				opset(AXSCVSXDDP, r0)
			psess.
				opset(AXSCVUXDDP, r0)
			psess.
				opset(AXSCVSXDSP, r0)
			psess.
				opset(AXSCVUXDSP, r0)

		case AXVCVX:
			psess.
				opset(AXVCVDPSXDS, r0)
			psess.
				opset(AXVCVDPSXWS, r0)
			psess.
				opset(AXVCVDPUXDS, r0)
			psess.
				opset(AXVCVDPUXWS, r0)
			psess.
				opset(AXVCVSPSXDS, r0)
			psess.
				opset(AXVCVSPSXWS, r0)
			psess.
				opset(AXVCVSPUXDS, r0)
			psess.
				opset(AXVCVSPUXWS, r0)

		case AXVCVXP:
			psess.
				opset(AXVCVSXDDP, r0)
			psess.
				opset(AXVCVSXWDP, r0)
			psess.
				opset(AXVCVUXDDP, r0)
			psess.
				opset(AXVCVUXWDP, r0)
			psess.
				opset(AXVCVSXDSP, r0)
			psess.
				opset(AXVCVSXWSP, r0)
			psess.
				opset(AXVCVUXDSP, r0)
			psess.
				opset(AXVCVUXWSP, r0)

		case AAND:
			psess.
				opset(AANDN, r0)
			psess.
				opset(AANDNCC, r0)
			psess.
				opset(AEQV, r0)
			psess.
				opset(AEQVCC, r0)
			psess.
				opset(ANAND, r0)
			psess.
				opset(ANANDCC, r0)
			psess.
				opset(ANOR, r0)
			psess.
				opset(ANORCC, r0)
			psess.
				opset(AORCC, r0)
			psess.
				opset(AORN, r0)
			psess.
				opset(AORNCC, r0)
			psess.
				opset(AXORCC, r0)

		case AADDME:
			psess.
				opset(AADDMECC, r0)
			psess.
				opset(AADDMEV, r0)
			psess.
				opset(AADDMEVCC, r0)
			psess.
				opset(AADDZE, r0)
			psess.
				opset(AADDZECC, r0)
			psess.
				opset(AADDZEV, r0)
			psess.
				opset(AADDZEVCC, r0)
			psess.
				opset(ASUBME, r0)
			psess.
				opset(ASUBMECC, r0)
			psess.
				opset(ASUBMEV, r0)
			psess.
				opset(ASUBMEVCC, r0)
			psess.
				opset(ASUBZE, r0)
			psess.
				opset(ASUBZECC, r0)
			psess.
				opset(ASUBZEV, r0)
			psess.
				opset(ASUBZEVCC, r0)

		case AADDC:
			psess.
				opset(AADDCCC, r0)

		case ABEQ:
			psess.
				opset(ABGE, r0)
			psess.
				opset(ABGT, r0)
			psess.
				opset(ABLE, r0)
			psess.
				opset(ABLT, r0)
			psess.
				opset(ABNE, r0)
			psess.
				opset(ABVC, r0)
			psess.
				opset(ABVS, r0)

		case ABR:
			psess.
				opset(ABL, r0)

		case ABC:
			psess.
				opset(ABCL, r0)

		case AEXTSB:
			psess.
				opset(AEXTSBCC, r0)
			psess.
				opset(AEXTSH, r0)
			psess.
				opset(AEXTSHCC, r0)
			psess.
				opset(ACNTLZW, r0)
			psess.
				opset(ACNTLZWCC, r0)
			psess.
				opset(ACNTLZD, r0)
			psess.
				opset(AEXTSW, r0)
			psess.
				opset(AEXTSWCC, r0)
			psess.
				opset(ACNTLZDCC, r0)

		case AFABS:
			psess.
				opset(AFABSCC, r0)
			psess.
				opset(AFNABS, r0)
			psess.
				opset(AFNABSCC, r0)
			psess.
				opset(AFNEG, r0)
			psess.
				opset(AFNEGCC, r0)
			psess.
				opset(AFRSP, r0)
			psess.
				opset(AFRSPCC, r0)
			psess.
				opset(AFCTIW, r0)
			psess.
				opset(AFCTIWCC, r0)
			psess.
				opset(AFCTIWZ, r0)
			psess.
				opset(AFCTIWZCC, r0)
			psess.
				opset(AFCTID, r0)
			psess.
				opset(AFCTIDCC, r0)
			psess.
				opset(AFCTIDZ, r0)
			psess.
				opset(AFCTIDZCC, r0)
			psess.
				opset(AFCFID, r0)
			psess.
				opset(AFCFIDCC, r0)
			psess.
				opset(AFCFIDU, r0)
			psess.
				opset(AFCFIDUCC, r0)
			psess.
				opset(AFCFIDS, r0)
			psess.
				opset(AFCFIDSCC, r0)
			psess.
				opset(AFRES, r0)
			psess.
				opset(AFRESCC, r0)
			psess.
				opset(AFRIM, r0)
			psess.
				opset(AFRIMCC, r0)
			psess.
				opset(AFRIP, r0)
			psess.
				opset(AFRIPCC, r0)
			psess.
				opset(AFRIZ, r0)
			psess.
				opset(AFRIZCC, r0)
			psess.
				opset(AFRIN, r0)
			psess.
				opset(AFRINCC, r0)
			psess.
				opset(AFRSQRTE, r0)
			psess.
				opset(AFRSQRTECC, r0)
			psess.
				opset(AFSQRT, r0)
			psess.
				opset(AFSQRTCC, r0)
			psess.
				opset(AFSQRTS, r0)
			psess.
				opset(AFSQRTSCC, r0)

		case AFADD:
			psess.
				opset(AFADDS, r0)
			psess.
				opset(AFADDCC, r0)
			psess.
				opset(AFADDSCC, r0)
			psess.
				opset(AFCPSGN, r0)
			psess.
				opset(AFCPSGNCC, r0)
			psess.
				opset(AFDIV, r0)
			psess.
				opset(AFDIVS, r0)
			psess.
				opset(AFDIVCC, r0)
			psess.
				opset(AFDIVSCC, r0)
			psess.
				opset(AFSUB, r0)
			psess.
				opset(AFSUBS, r0)
			psess.
				opset(AFSUBCC, r0)
			psess.
				opset(AFSUBSCC, r0)

		case AFMADD:
			psess.
				opset(AFMADDCC, r0)
			psess.
				opset(AFMADDS, r0)
			psess.
				opset(AFMADDSCC, r0)
			psess.
				opset(AFMSUB, r0)
			psess.
				opset(AFMSUBCC, r0)
			psess.
				opset(AFMSUBS, r0)
			psess.
				opset(AFMSUBSCC, r0)
			psess.
				opset(AFNMADD, r0)
			psess.
				opset(AFNMADDCC, r0)
			psess.
				opset(AFNMADDS, r0)
			psess.
				opset(AFNMADDSCC, r0)
			psess.
				opset(AFNMSUB, r0)
			psess.
				opset(AFNMSUBCC, r0)
			psess.
				opset(AFNMSUBS, r0)
			psess.
				opset(AFNMSUBSCC, r0)
			psess.
				opset(AFSEL, r0)
			psess.
				opset(AFSELCC, r0)

		case AFMUL:
			psess.
				opset(AFMULS, r0)
			psess.
				opset(AFMULCC, r0)
			psess.
				opset(AFMULSCC, r0)

		case AFCMPO:
			psess.
				opset(AFCMPU, r0)

		case AISEL:
			psess.
				opset(AISEL, r0)

		case AMTFSB0:
			psess.
				opset(AMTFSB0CC, r0)
			psess.
				opset(AMTFSB1, r0)
			psess.
				opset(AMTFSB1CC, r0)

		case ANEG:
			psess.
				opset(ANEGCC, r0)
			psess.
				opset(ANEGV, r0)
			psess.
				opset(ANEGVCC, r0)

		case AOR:
			psess.
				opset(AXOR, r0)

		case AORIS:
			psess.
				opset(AXORIS, r0)

		case ASLW:
			psess.
				opset(ASLWCC, r0)
			psess.
				opset(ASRW, r0)
			psess.
				opset(ASRWCC, r0)
			psess.
				opset(AROTLW, r0)

		case ASLD:
			psess.
				opset(ASLDCC, r0)
			psess.
				opset(ASRD, r0)
			psess.
				opset(ASRDCC, r0)
			psess.
				opset(AROTL, r0)

		case ASRAW:
			psess.
				opset(ASRAWCC, r0)

		case ASRAD:
			psess.
				opset(ASRADCC, r0)

		case ASUB:
			psess.
				opset(ASUB, r0)
			psess.
				opset(ASUBCC, r0)
			psess.
				opset(ASUBV, r0)
			psess.
				opset(ASUBVCC, r0)
			psess.
				opset(ASUBCCC, r0)
			psess.
				opset(ASUBCV, r0)
			psess.
				opset(ASUBCVCC, r0)
			psess.
				opset(ASUBE, r0)
			psess.
				opset(ASUBECC, r0)
			psess.
				opset(ASUBEV, r0)
			psess.
				opset(ASUBEVCC, r0)

		case ASYNC:
			psess.
				opset(AISYNC, r0)
			psess.
				opset(ALWSYNC, r0)
			psess.
				opset(APTESYNC, r0)
			psess.
				opset(ATLBSYNC, r0)

		case ARLWMI:
			psess.
				opset(ARLWMICC, r0)
			psess.
				opset(ARLWNM, r0)
			psess.
				opset(ARLWNMCC, r0)

		case ARLDMI:
			psess.
				opset(ARLDMICC, r0)
			psess.
				opset(ARLDIMI, r0)
			psess.
				opset(ARLDIMICC, r0)

		case ARLDC:
			psess.
				opset(ARLDCCC, r0)

		case ARLDCL:
			psess.
				opset(ARLDCR, r0)
			psess.
				opset(ARLDCLCC, r0)
			psess.
				opset(ARLDCRCC, r0)

		case ARLDICL:
			psess.
				opset(ARLDICLCC, r0)
			psess.
				opset(ARLDICR, r0)
			psess.
				opset(ARLDICRCC, r0)

		case AFMOVD:
			psess.
				opset(AFMOVDCC, r0)
			psess.
				opset(AFMOVDU, r0)
			psess.
				opset(AFMOVS, r0)
			psess.
				opset(AFMOVSU, r0)

		case ALDAR:
			psess.
				opset(ALBAR, r0)
			psess.
				opset(ALHAR, r0)
			psess.
				opset(ALWAR, r0)

		case ASYSCALL:
			psess.
				opset(ARFI, r0)
			psess.
				opset(ARFCI, r0)
			psess.
				opset(ARFID, r0)
			psess.
				opset(AHRFID, r0)

		case AMOVHBR:
			psess.
				opset(AMOVWBR, r0)
			psess.
				opset(AMOVDBR, r0)

		case ASLBMFEE:
			psess.
				opset(ASLBMFEV, r0)

		case ATW:
			psess.
				opset(ATD, r0)

		case ATLBIE:
			psess.
				opset(ASLBIE, r0)
			psess.
				opset(ATLBIEL, r0)

		case AEIEIO:
			psess.
				opset(ASLBIA, r0)

		case ACMP:
			psess.
				opset(ACMPW, r0)

		case ACMPU:
			psess.
				opset(ACMPWU, r0)

		case ACMPB:
			psess.
				opset(ACMPB, r0)

		case AFTDIV:
			psess.
				opset(AFTDIV, r0)

		case AFTSQRT:
			psess.
				opset(AFTSQRT, r0)

		case AADD,
			AADDIS,
			AANDCC,
			AANDISCC,
			AFMOVSX,
			AFMOVSZ,
			ALSW,
			AMOVW,

			AMOVWZ,
			AMOVD,
			AMOVB,
			AMOVBU,
			AMOVFL,
			AMULLW,

			ASUBC,
			ASTSW,
			ASLBMTE,
			AWORD,
			ADWORD,
			ADARN,
			ALDMX,
			AVMSUMUDM,
			AADDEX,
			ACMPEQB,
			AECIWX,
			obj.ANOP,
			obj.ATEXT,
			obj.AUNDEF,
			obj.AFUNCDATA,
			obj.APCDATA,
			obj.ADUFFZERO,
			obj.ADUFFCOPY:
			break
		}
	}
}

func OPVXX1(o uint32, xo uint32, oe uint32) uint32 {
	return o<<26 | xo<<1 | oe<<11
}

func OPVXX2(o uint32, xo uint32, oe uint32) uint32 {
	return o<<26 | xo<<2 | oe<<11
}

func OPVXX3(o uint32, xo uint32, oe uint32) uint32 {
	return o<<26 | xo<<3 | oe<<11
}

func OPVXX4(o uint32, xo uint32, oe uint32) uint32 {
	return o<<26 | xo<<4 | oe<<11
}

func OPVX(o uint32, xo uint32, oe uint32, rc uint32) uint32 {
	return o<<26 | xo | oe<<11 | rc&1
}

func OPVC(o uint32, xo uint32, oe uint32, rc uint32) uint32 {
	return o<<26 | xo | oe<<11 | (rc&1)<<10
}

func OPVCC(o uint32, xo uint32, oe uint32, rc uint32) uint32 {
	return o<<26 | xo<<1 | oe<<10 | rc&1
}

func OPCC(o uint32, xo uint32, rc uint32) uint32 {
	return OPVCC(o, xo, 0, rc)
}

func OP(o uint32, xo uint32) uint32 {
	return OPVCC(o, xo, 0, 0)
}

/* the order is dest, a/s, b/imm for both arithmetic and logical operations */
func AOP_RRR(op uint32, d uint32, a uint32, b uint32) uint32 {
	return op | (d&31)<<21 | (a&31)<<16 | (b&31)<<11
}

/* VX-form 2-register operands, r/none/r */
func AOP_RR(op uint32, d uint32, a uint32) uint32 {
	return op | (d&31)<<21 | (a&31)<<11
}

/* VA-form 4-register operands */
func AOP_RRRR(op uint32, d uint32, a uint32, b uint32, c uint32) uint32 {
	return op | (d&31)<<21 | (a&31)<<16 | (b&31)<<11 | (c&31)<<6
}

func AOP_IRR(op uint32, d uint32, a uint32, simm uint32) uint32 {
	return op | (d&31)<<21 | (a&31)<<16 | simm&0xFFFF
}

/* VX-form 2-register + UIM operands */
func AOP_VIRR(op uint32, d uint32, a uint32, simm uint32) uint32 {
	return op | (d&31)<<21 | (simm&0xFFFF)<<16 | (a&31)<<11
}

/* VX-form 2-register + ST + SIX operands */
func AOP_IIRR(op uint32, d uint32, a uint32, sbit uint32, simm uint32) uint32 {
	return op | (d&31)<<21 | (a&31)<<16 | (sbit&1)<<15 | (simm&0xF)<<11
}

/* VA-form 3-register + SHB operands */
func AOP_IRRR(op uint32, d uint32, a uint32, b uint32, simm uint32) uint32 {
	return op | (d&31)<<21 | (a&31)<<16 | (b&31)<<11 | (simm&0xF)<<6
}

/* VX-form 1-register + SIM operands */
func AOP_IR(op uint32, d uint32, simm uint32) uint32 {
	return op | (d&31)<<21 | (simm&31)<<16
}

/* XX1-form 3-register operands, 1 VSR operand */
func AOP_XX1(op uint32, d uint32, a uint32, b uint32) uint32 {

	r := d - REG_VS0
	return op | (r&31)<<21 | (a&31)<<16 | (b&31)<<11 | (r&32)>>5
}

/* XX2-form 3-register operands, 2 VSR operands */
func AOP_XX2(op uint32, d uint32, a uint32, b uint32) uint32 {
	xt := d - REG_VS0
	xb := b - REG_VS0
	return op | (xt&31)<<21 | (a&3)<<16 | (xb&31)<<11 | (xb&32)>>4 | (xt&32)>>5
}

/* XX3-form 3 VSR operands */
func AOP_XX3(op uint32, d uint32, a uint32, b uint32) uint32 {
	xt := d - REG_VS0
	xa := a - REG_VS0
	xb := b - REG_VS0
	return op | (xt&31)<<21 | (xa&31)<<16 | (xb&31)<<11 | (xa&32)>>3 | (xb&32)>>4 | (xt&32)>>5
}

/* XX3-form 3 VSR operands + immediate */
func AOP_XX3I(op uint32, d uint32, a uint32, b uint32, c uint32) uint32 {
	xt := d - REG_VS0
	xa := a - REG_VS0
	xb := b - REG_VS0
	return op | (xt&31)<<21 | (xa&31)<<16 | (xb&31)<<11 | (c&3)<<8 | (xa&32)>>3 | (xb&32)>>4 | (xt&32)>>5
}

/* XX4-form, 4 VSR operands */
func AOP_XX4(op uint32, d uint32, a uint32, b uint32, c uint32) uint32 {
	xt := d - REG_VS0
	xa := a - REG_VS0
	xb := b - REG_VS0
	xc := c - REG_VS0
	return op | (xt&31)<<21 | (xa&31)<<16 | (xb&31)<<11 | (xc&31)<<6 | (xc&32)>>2 | (xa&32)>>3 | (xb&32)>>4 | (xt&32)>>5
}

/* Z23-form, 3-register operands + CY field */
func AOP_Z23I(op uint32, d uint32, a uint32, b uint32, c uint32) uint32 {
	return op | (d&31)<<21 | (a&31)<<16 | (b&31)<<11 | (c&3)<<7
}

/* X-form, 3-register operands + EH field */
func AOP_RRRI(op uint32, d uint32, a uint32, b uint32, c uint32) uint32 {
	return op | (d&31)<<21 | (a&31)<<16 | (b&31)<<11 | (c & 1)
}

func LOP_RRR(op uint32, a uint32, s uint32, b uint32) uint32 {
	return op | (s&31)<<21 | (a&31)<<16 | (b&31)<<11
}

func LOP_IRR(op uint32, a uint32, s uint32, uimm uint32) uint32 {
	return op | (s&31)<<21 | (a&31)<<16 | uimm&0xFFFF
}

func OP_BR(op uint32, li uint32, aa uint32) uint32 {
	return op | li&0x03FFFFFC | aa<<1
}

func OP_BC(op uint32, bo uint32, bi uint32, bd uint32, aa uint32) uint32 {
	return op | (bo&0x1F)<<21 | (bi&0x1F)<<16 | bd&0xFFFC | aa<<1
}

func OP_BCR(op uint32, bo uint32, bi uint32) uint32 {
	return op | (bo&0x1F)<<21 | (bi&0x1F)<<16
}

func OP_RLW(op uint32, a uint32, s uint32, sh uint32, mb uint32, me uint32) uint32 {
	return op | (s&31)<<21 | (a&31)<<16 | (sh&31)<<11 | (mb&31)<<6 | (me&31)<<1
}

func AOP_RLDIC(op uint32, a uint32, s uint32, sh uint32, m uint32) uint32 {
	return op | (s&31)<<21 | (a&31)<<16 | (sh&31)<<11 | ((sh&32)>>5)<<1 | (m&31)<<6 | ((m&32)>>5)<<5
}

func AOP_ISEL(op uint32, t uint32, a uint32, b uint32, bc uint32) uint32 {
	return op | (t&31)<<21 | (a&31)<<16 | (b&31)<<11 | (bc&0x1F)<<6
}

const (
	/* each rhs is OPVCC(_, _, _, _) */
	OP_ADD    = 31<<26 | 266<<1 | 0<<10 | 0
	OP_ADDI   = 14<<26 | 0<<1 | 0<<10 | 0
	OP_ADDIS  = 15<<26 | 0<<1 | 0<<10 | 0
	OP_ANDI   = 28<<26 | 0<<1 | 0<<10 | 0
	OP_EXTSB  = 31<<26 | 954<<1 | 0<<10 | 0
	OP_EXTSH  = 31<<26 | 922<<1 | 0<<10 | 0
	OP_EXTSW  = 31<<26 | 986<<1 | 0<<10 | 0
	OP_ISEL   = 31<<26 | 15<<1 | 0<<10 | 0
	OP_MCRF   = 19<<26 | 0<<1 | 0<<10 | 0
	OP_MCRFS  = 63<<26 | 64<<1 | 0<<10 | 0
	OP_MCRXR  = 31<<26 | 512<<1 | 0<<10 | 0
	OP_MFCR   = 31<<26 | 19<<1 | 0<<10 | 0
	OP_MFFS   = 63<<26 | 583<<1 | 0<<10 | 0
	OP_MFMSR  = 31<<26 | 83<<1 | 0<<10 | 0
	OP_MFSPR  = 31<<26 | 339<<1 | 0<<10 | 0
	OP_MFSR   = 31<<26 | 595<<1 | 0<<10 | 0
	OP_MFSRIN = 31<<26 | 659<<1 | 0<<10 | 0
	OP_MTCRF  = 31<<26 | 144<<1 | 0<<10 | 0
	OP_MTFSF  = 63<<26 | 711<<1 | 0<<10 | 0
	OP_MTFSFI = 63<<26 | 134<<1 | 0<<10 | 0
	OP_MTMSR  = 31<<26 | 146<<1 | 0<<10 | 0
	OP_MTMSRD = 31<<26 | 178<<1 | 0<<10 | 0
	OP_MTSPR  = 31<<26 | 467<<1 | 0<<10 | 0
	OP_MTSR   = 31<<26 | 210<<1 | 0<<10 | 0
	OP_MTSRIN = 31<<26 | 242<<1 | 0<<10 | 0
	OP_MULLW  = 31<<26 | 235<<1 | 0<<10 | 0
	OP_MULLD  = 31<<26 | 233<<1 | 0<<10 | 0
	OP_OR     = 31<<26 | 444<<1 | 0<<10 | 0
	OP_ORI    = 24<<26 | 0<<1 | 0<<10 | 0
	OP_ORIS   = 25<<26 | 0<<1 | 0<<10 | 0
	OP_RLWINM = 21<<26 | 0<<1 | 0<<10 | 0
	OP_RLWNM  = 23<<26 | 0<<1 | 0<<10 | 0
	OP_SUBF   = 31<<26 | 40<<1 | 0<<10 | 0
	OP_RLDIC  = 30<<26 | 4<<1 | 0<<10 | 0
	OP_RLDICR = 30<<26 | 2<<1 | 0<<10 | 0
	OP_RLDICL = 30<<26 | 0<<1 | 0<<10 | 0
	OP_RLDCL  = 30<<26 | 8<<1 | 0<<10 | 0
)

func oclass(a *obj.Addr) int {
	return int(a.Class) - 1
}

const (
	D_FORM = iota
	DS_FORM
)

// This function determines when a non-indexed load or store is D or
// DS form for use in finding the size of the offset field in the instruction.
// The size is needed when setting the offset value in the instruction
// and when generating relocation for that field.
// DS form instructions include: ld, ldu, lwa, std, stdu.  All other
// loads and stores with an offset field are D form.  This function should
// only be called with the same opcodes as are handled by opstore and opload.
func (c *ctxt9) opform(insn uint32) int {
	switch insn {
	default:
		c.ctxt.Diag("bad insn in loadform: %x", insn)
	case OPVCC(58, 0, 0, 0),
		OPVCC(58, 0, 0, 1),
		OPVCC(58, 0, 0, 0) | 1<<1,
		OPVCC(62, 0, 0, 0),
		OPVCC(62, 0, 0, 1):
		return DS_FORM
	case OP_ADDI,
		OPVCC(32, 0, 0, 0),
		OPVCC(33, 0, 0, 0),
		OPVCC(34, 0, 0, 0),
		OPVCC(35, 0, 0, 0),
		OPVCC(40, 0, 0, 0),
		OPVCC(41, 0, 0, 0),
		OPVCC(42, 0, 0, 0),
		OPVCC(43, 0, 0, 0),
		OPVCC(46, 0, 0, 0),
		OPVCC(48, 0, 0, 0),
		OPVCC(49, 0, 0, 0),
		OPVCC(50, 0, 0, 0),
		OPVCC(51, 0, 0, 0),
		OPVCC(36, 0, 0, 0),
		OPVCC(37, 0, 0, 0),
		OPVCC(38, 0, 0, 0),
		OPVCC(39, 0, 0, 0),
		OPVCC(44, 0, 0, 0),
		OPVCC(45, 0, 0, 0),
		OPVCC(47, 0, 0, 0),
		OPVCC(52, 0, 0, 0),
		OPVCC(53, 0, 0, 0),
		OPVCC(54, 0, 0, 0),
		OPVCC(55, 0, 0, 0):
		return D_FORM
	}
	return 0
}

// Encode instructions and create relocation for accessing s+d according to the
// instruction op with source or destination (as appropriate) register reg.
func (c *ctxt9) symbolAccess(s *obj.LSym, d int64, reg int16, op uint32) (o1, o2 uint32) {
	var base uint32
	form := c.opform(op)
	if c.ctxt.Flag_shared {
		base = REG_R2
	} else {
		base = REG_R0
	}
	o1 = AOP_IRR(OP_ADDIS, REGTMP, base, 0)
	o2 = AOP_IRR(op, uint32(reg), REGTMP, 0)
	rel := obj.Addrel(c.cursym)
	rel.Off = int32(c.pc)
	rel.Siz = 8
	rel.Sym = s
	rel.Add = d
	if c.ctxt.Flag_shared {
		switch form {
		case D_FORM:
			rel.Type = objabi.R_ADDRPOWER_TOCREL
		case DS_FORM:
			rel.Type = objabi.R_ADDRPOWER_TOCREL_DS
		}

	} else {
		switch form {
		case D_FORM:
			rel.Type = objabi.R_ADDRPOWER
		case DS_FORM:
			rel.Type = objabi.R_ADDRPOWER_DS
		}
	}
	return
}

/*
 * 32-bit masks
 */
func getmask(m []byte, v uint32) bool {
	m[1] = 0
	m[0] = m[1]
	if v != ^uint32(0) && v&(1<<31) != 0 && v&1 != 0 {
		if getmask(m, ^v) {
			i := int(m[0])
			m[0] = m[1] + 1
			m[1] = byte(i - 1)
			return true
		}

		return false
	}

	for i := 0; i < 32; i++ {
		if v&(1<<uint(31-i)) != 0 {
			m[0] = byte(i)
			for {
				m[1] = byte(i)
				i++
				if i >= 32 || v&(1<<uint(31-i)) == 0 {
					break
				}
			}

			for ; i < 32; i++ {
				if v&(1<<uint(31-i)) != 0 {
					return false
				}
			}
			return true
		}
	}

	return false
}

func (c *ctxt9) maskgen(p *obj.Prog, m []byte, v uint32) {
	if !getmask(m, v) {
		c.ctxt.Diag("cannot generate mask #%x\n%v", v, p)
	}
}

/*
 * 64-bit masks (rldic etc)
 */
func getmask64(m []byte, v uint64) bool {
	m[1] = 0
	m[0] = m[1]
	for i := 0; i < 64; i++ {
		if v&(uint64(1)<<uint(63-i)) != 0 {
			m[0] = byte(i)
			for {
				m[1] = byte(i)
				i++
				if i >= 64 || v&(uint64(1)<<uint(63-i)) == 0 {
					break
				}
			}

			for ; i < 64; i++ {
				if v&(uint64(1)<<uint(63-i)) != 0 {
					return false
				}
			}
			return true
		}
	}

	return false
}

func (c *ctxt9) maskgen64(p *obj.Prog, m []byte, v uint64) {
	if !getmask64(m, v) {
		c.ctxt.Diag("cannot generate mask #%x\n%v", v, p)
	}
}

func loadu32(r int, d int64) uint32 {
	v := int32(d >> 16)
	if isuint32(uint64(d)) {
		return LOP_IRR(OP_ORIS, uint32(r), REGZERO, uint32(v))
	}
	return AOP_IRR(OP_ADDIS, uint32(r), REGZERO, uint32(v))
}

func high16adjusted(d int32) uint16 {
	if d&0x8000 != 0 {
		return uint16((d >> 16) + 1)
	}
	return uint16(d >> 16)
}

func (c *ctxt9) asmout(p *obj.Prog, o *Optab, out []uint32) {
	o1 := uint32(0)
	o2 := uint32(0)
	o3 := uint32(0)
	o4 := uint32(0)
	o5 := uint32(0)

	switch o.type_ {
	default:
		c.ctxt.Diag("unknown type %d", o.type_)
		prasm(p)

	case 0:
		break

	case 1:
		if p.To.Reg == REGZERO && p.From.Type == obj.TYPE_CONST {
			v := c.regoff(&p.From)
			if r0iszero != 0 && v != 0 {

				c.ctxt.Diag("literal operation on R0\n%v", p)
			}

			o1 = LOP_IRR(OP_ADDI, REGZERO, REGZERO, uint32(v))
			break
		}

		o1 = LOP_RRR(OP_OR, uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.From.Reg))

	case 2:
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), uint32(p.From.Reg))

	case 3:
		d := c.vregoff(&p.From)

		v := int32(d)
		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		if r0iszero != 0 && p.To.Reg == 0 && (r != 0 || v != 0) {
			c.ctxt.Diag("literal operation on R0\n%v", p)
		}
		a := OP_ADDI
		if o.a1 == C_UCON {
			if d&0xffff != 0 {
				log.Fatalf("invalid handling of %v", p)
			}

			v >>= 16
			if r == REGZERO && isuint32(uint64(d)) {
				o1 = LOP_IRR(OP_ORIS, uint32(p.To.Reg), REGZERO, uint32(v))
				break
			}

			a = OP_ADDIS
		} else if int64(int16(d)) != d {

			if o.a1 == C_ANDCON {

				if r == 0 || r == REGZERO {
					o1 = LOP_IRR(uint32(OP_ORI), uint32(p.To.Reg), uint32(0), uint32(v))
					break
				}

			} else if o.a1 != C_ADDCON {
				log.Fatalf("invalid handling of %v", p)
			}
		}

		o1 = AOP_IRR(uint32(a), uint32(p.To.Reg), uint32(r), uint32(v))

	case 4:
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		if r0iszero != 0 && p.To.Reg == 0 {
			c.ctxt.Diag("literal operation on R0\n%v", p)
		}
		if int32(int16(v)) != v {
			log.Fatalf("mishandled instruction %v", p)
		}
		o1 = AOP_IRR(c.opirr(p.As), uint32(p.To.Reg), uint32(r), uint32(v))

	case 5:
		o1 = c.oprrr(p.As)

	case 6:
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}

		switch p.As {
		case AROTL:
			o1 = AOP_RLDIC(OP_RLDCL, uint32(p.To.Reg), uint32(r), uint32(p.From.Reg), uint32(0))
		case AROTLW:
			o1 = OP_RLW(OP_RLWNM, uint32(p.To.Reg), uint32(r), uint32(p.From.Reg), 0, 31)
		default:
			o1 = LOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), uint32(p.From.Reg))
		}

	case 7:
		r := int(p.To.Reg)

		if r == 0 {
			r = int(o.param)
		}
		v := c.regoff(&p.To)
		if p.To.Type == obj.TYPE_MEM && p.To.Index != 0 {
			if v != 0 {
				c.ctxt.Diag("illegal indexed instruction\n%v", p)
			}
			if c.ctxt.Flag_shared && r == REG_R13 {
				rel := obj.Addrel(c.cursym)
				rel.Off = int32(c.pc)
				rel.Siz = 4

				rel.Sym = c.ctxt.Lookup("runtime.tls_g")
				rel.Type = objabi.R_POWER_TLS
			}
			o1 = AOP_RRR(c.opstorex(p.As), uint32(p.From.Reg), uint32(p.To.Index), uint32(r))
		} else {
			if int32(int16(v)) != v {
				log.Fatalf("mishandled instruction %v", p)
			}

			inst := c.opstore(p.As)
			if c.opform(inst) == DS_FORM && v&0x3 != 0 {
				log.Fatalf("invalid offset for DS form load/store %v", p)
			}
			o1 = AOP_IRR(inst, uint32(p.From.Reg), uint32(r), uint32(v))
		}

	case 8:
		r := int(p.From.Reg)

		if r == 0 {
			r = int(o.param)
		}
		v := c.regoff(&p.From)
		if p.From.Type == obj.TYPE_MEM && p.From.Index != 0 {
			if v != 0 {
				c.ctxt.Diag("illegal indexed instruction\n%v", p)
			}
			if c.ctxt.Flag_shared && r == REG_R13 {
				rel := obj.Addrel(c.cursym)
				rel.Off = int32(c.pc)
				rel.Siz = 4
				rel.Sym = c.ctxt.Lookup("runtime.tls_g")
				rel.Type = objabi.R_POWER_TLS
			}
			o1 = AOP_RRR(c.oploadx(p.As), uint32(p.To.Reg), uint32(p.From.Index), uint32(r))
		} else {
			if int32(int16(v)) != v {
				log.Fatalf("mishandled instruction %v", p)
			}

			inst := c.opload(p.As)
			if c.opform(inst) == DS_FORM && v&0x3 != 0 {
				log.Fatalf("invalid offset for DS form load/store %v", p)
			}
			o1 = AOP_IRR(inst, uint32(p.To.Reg), uint32(r), uint32(v))
		}

	case 9:
		r := int(p.From.Reg)

		if r == 0 {
			r = int(o.param)
		}
		v := c.regoff(&p.From)
		if p.From.Type == obj.TYPE_MEM && p.From.Index != 0 {
			if v != 0 {
				c.ctxt.Diag("illegal indexed instruction\n%v", p)
			}
			o1 = AOP_RRR(c.oploadx(p.As), uint32(p.To.Reg), uint32(p.From.Index), uint32(r))
		} else {
			o1 = AOP_IRR(c.opload(p.As), uint32(p.To.Reg), uint32(r), uint32(v))
		}
		o2 = LOP_RRR(OP_EXTSB, uint32(p.To.Reg), uint32(p.To.Reg), 0)

	case 10:
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(r))

	case 11:
		v := int32(0)

		if p.Pcond != nil {
			v = int32(p.Pcond.Pc - p.Pc)
			if v&03 != 0 {
				c.ctxt.Diag("odd branch target address\n%v", p)
				v &^= 03
			}

			if v < -(1<<25) || v >= 1<<24 {
				c.ctxt.Diag("branch too far\n%v", p)
			}
		}

		o1 = OP_BR(c.opirr(p.As), uint32(v), 0)
		if p.To.Sym != nil {
			rel := obj.Addrel(c.cursym)
			rel.Off = int32(c.pc)
			rel.Siz = 4
			rel.Sym = p.To.Sym
			v += int32(p.To.Offset)
			if v&03 != 0 {
				c.ctxt.Diag("odd branch target address\n%v", p)
				v &^= 03
			}

			rel.Add = int64(v)
			rel.Type = objabi.R_CALLPOWER
		}
		o2 = 0x60000000

	case 12:
		if p.To.Reg == REGZERO && p.From.Type == obj.TYPE_CONST {
			v := c.regoff(&p.From)
			if r0iszero != 0 && v != 0 {
				c.ctxt.Diag("literal operation on R0\n%v", p)
			}

			o1 = LOP_IRR(OP_ADDI, REGZERO, REGZERO, uint32(v))
			break
		}

		if p.As == AMOVW {
			o1 = LOP_RRR(OP_EXTSW, uint32(p.To.Reg), uint32(p.From.Reg), 0)
		} else {
			o1 = LOP_RRR(OP_EXTSB, uint32(p.To.Reg), uint32(p.From.Reg), 0)
		}

	case 13:
		if p.As == AMOVBZ {
			o1 = OP_RLW(OP_RLWINM, uint32(p.To.Reg), uint32(p.From.Reg), 0, 24, 31)
		} else if p.As == AMOVH {
			o1 = LOP_RRR(OP_EXTSH, uint32(p.To.Reg), uint32(p.From.Reg), 0)
		} else if p.As == AMOVHZ {
			o1 = OP_RLW(OP_RLWINM, uint32(p.To.Reg), uint32(p.From.Reg), 0, 16, 31)
		} else if p.As == AMOVWZ {
			o1 = OP_RLW(OP_RLDIC, uint32(p.To.Reg), uint32(p.From.Reg), 0, 0, 0) | 1<<5
		} else {
			c.ctxt.Diag("internal: bad mov[bhw]z\n%v", p)
		}

	case 14:
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		d := c.vregoff(p.GetFrom3())
		var a int
		switch p.As {

		case ARLDCL, ARLDCLCC:
			var mask [2]uint8
			c.maskgen64(p, mask[:], uint64(d))

			a = int(mask[0])
			if mask[1] != 63 {
				c.ctxt.Diag("invalid mask for rotate: %x (end != bit 63)\n%v", uint64(d), p)
			}
			o1 = LOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), uint32(p.From.Reg))
			o1 |= (uint32(a) & 31) << 6
			if a&0x20 != 0 {
				o1 |= 1 << 5
			}

		case ARLDCR, ARLDCRCC:
			var mask [2]uint8
			c.maskgen64(p, mask[:], uint64(d))

			a = int(mask[1])
			if mask[0] != 0 {
				c.ctxt.Diag("invalid mask for rotate: %x (start != 0)\n%v", uint64(d), p)
			}
			o1 = LOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), uint32(p.From.Reg))
			o1 |= (uint32(a) & 31) << 6
			if a&0x20 != 0 {
				o1 |= 1 << 5
			}

		case ARLDICR, ARLDICRCC:
			me := int(d)
			sh := c.regoff(&p.From)
			o1 = AOP_RLDIC(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), uint32(sh), uint32(me))

		case ARLDICL, ARLDICLCC:
			mb := int(d)
			sh := c.regoff(&p.From)
			o1 = AOP_RLDIC(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), uint32(sh), uint32(mb))

		default:
			c.ctxt.Diag("unexpected op in rldc case\n%v", p)
			a = 0
		}

	case 17,
		16:
		a := 0

		r := int(p.Reg)

		if p.From.Type == obj.TYPE_CONST {
			a = int(c.regoff(&p.From))
		} else if p.From.Type == obj.TYPE_REG {
			if r != 0 {
				c.ctxt.Diag("unexpected register setting for branch with CR: %d\n", r)
			}

			switch p.From.Reg {
			case REG_CR0:
				r = BI_CR0
			case REG_CR1:
				r = BI_CR1
			case REG_CR2:
				r = BI_CR2
			case REG_CR3:
				r = BI_CR3
			case REG_CR4:
				r = BI_CR4
			case REG_CR5:
				r = BI_CR5
			case REG_CR6:
				r = BI_CR6
			case REG_CR7:
				r = BI_CR7
			default:
				c.ctxt.Diag("unrecognized register: expecting CR\n")
			}
		}
		v := int32(0)
		if p.Pcond != nil {
			v = int32(p.Pcond.Pc - p.Pc)
		}
		if v&03 != 0 {
			c.ctxt.Diag("odd branch target address\n%v", p)
			v &^= 03
		}

		if v < -(1<<16) || v >= 1<<15 {
			c.ctxt.Diag("branch too far\n%v", p)
		}
		o1 = OP_BC(c.opirr(p.As), uint32(a), uint32(r), uint32(v), 0)

	case 15:
		var v int32
		if p.As == ABC || p.As == ABCL {
			v = c.regoff(&p.To) & 31
		} else {
			v = 20
		}
		o1 = AOP_RRR(OP_MTSPR, uint32(p.To.Reg), 0, 0) | (REG_LR&0x1f)<<16 | ((REG_LR>>5)&0x1f)<<11
		o2 = OPVCC(19, 16, 0, 0)
		if p.As == ABL || p.As == ABCL {
			o2 |= 1
		}
		o2 = OP_BCR(o2, uint32(v), uint32(p.To.Index))

	case 18:
		var v int32
		if p.As == ABC || p.As == ABCL {
			v = c.regoff(&p.From) & 31
		} else {
			v = 20
		}
		r := int(p.Reg)
		if r == 0 {
			r = 0
		}
		switch oclass(&p.To) {
		case C_CTR:
			o1 = OPVCC(19, 528, 0, 0)

		case C_LR:
			o1 = OPVCC(19, 16, 0, 0)

		default:
			c.ctxt.Diag("bad optab entry (18): %d\n%v", p.To.Class, p)
			v = 0
		}

		if p.As == ABL || p.As == ABCL {
			o1 |= 1
		}
		o1 = OP_BCR(o1, uint32(v), uint32(r))

	case 19:
		d := c.vregoff(&p.From)

		if p.From.Sym == nil {
			o1 = loadu32(int(p.To.Reg), d)
			o2 = LOP_IRR(OP_ORI, uint32(p.To.Reg), uint32(p.To.Reg), uint32(int32(d)))
		} else {
			o1, o2 = c.symbolAccess(p.From.Sym, d, p.To.Reg, OP_ADDI)
		}

	case 20:
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		if p.As == AADD && (r0iszero == 0 && p.Reg == 0 || r0iszero != 0 && p.To.Reg == 0) {
			c.ctxt.Diag("literal operation on R0\n%v", p)
		}
		if p.As == AADDIS {
			o1 = AOP_IRR(c.opirr(p.As), uint32(p.To.Reg), uint32(r), uint32(v))
		} else {
			o1 = AOP_IRR(c.opirr(AADDIS), uint32(p.To.Reg), uint32(r), uint32(v)>>16)
		}

	case 22:
		if p.To.Reg == REGTMP || p.Reg == REGTMP {
			c.ctxt.Diag("can't synthesize large constant\n%v", p)
		}
		d := c.vregoff(&p.From)
		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		if p.From.Sym != nil {
			c.ctxt.Diag("%v is not supported", p)
		}

		if o.size == 8 {
			o1 = LOP_IRR(OP_ORI, REGTMP, REGZERO, uint32(int32(d)))
			o2 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), REGTMP, uint32(r))
		} else {
			o1 = loadu32(REGTMP, d)
			o2 = LOP_IRR(OP_ORI, REGTMP, REGTMP, uint32(int32(d)))
			o3 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), REGTMP, uint32(r))
		}

	case 23:
		if p.To.Reg == REGTMP || p.Reg == REGTMP {
			c.ctxt.Diag("can't synthesize large constant\n%v", p)
		}
		d := c.vregoff(&p.From)
		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}

		if o.size == 8 {
			o1 = LOP_IRR(OP_ADDI, REGZERO, REGTMP, uint32(int32(d)))
			o2 = LOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), REGTMP, uint32(r))
		} else {
			o1 = loadu32(REGTMP, d)
			o2 = LOP_IRR(OP_ORI, REGTMP, REGTMP, uint32(int32(d)))
			o3 = LOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), REGTMP, uint32(r))
		}
		if p.From.Sym != nil {
			c.ctxt.Diag("%v is not supported", p)
		}

	case 25:

		v := c.regoff(&p.From)

		if v < 0 {
			v = 0
		} else if v > 63 {
			v = 63
		}
		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		var a int
		op := uint32(0)
		switch p.As {
		case ASLD, ASLDCC:
			a = int(63 - v)
			op = OP_RLDICR

		case ASRD, ASRDCC:
			a = int(v)
			v = 64 - v
			op = OP_RLDICL
		case AROTL:
			a = int(0)
			op = OP_RLDICL
		default:
			c.ctxt.Diag("unexpected op in sldi case\n%v", p)
			a = 0
			o1 = 0
		}

		o1 = AOP_RLDIC(op, uint32(p.To.Reg), uint32(r), uint32(v), uint32(a))
		if p.As == ASLDCC || p.As == ASRDCC {
			o1 |= 1
		}

	case 26:
		if p.To.Reg == REGTMP {
			c.ctxt.Diag("can't synthesize large constant\n%v", p)
		}
		v := c.regoff(&p.From)
		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		o1 = AOP_IRR(OP_ADDIS, REGTMP, uint32(r), uint32(high16adjusted(v)))
		o2 = AOP_IRR(OP_ADDI, uint32(p.To.Reg), REGTMP, uint32(v))

	case 27:
		v := c.regoff(p.GetFrom3())

		r := int(p.From.Reg)
		o1 = AOP_IRR(c.opirr(p.As), uint32(p.To.Reg), uint32(r), uint32(v))

	case 28:
		if p.To.Reg == REGTMP || p.From.Reg == REGTMP {
			c.ctxt.Diag("can't synthesize large constant\n%v", p)
		}
		v := c.regoff(p.GetFrom3())
		o1 = AOP_IRR(OP_ADDIS, REGTMP, REGZERO, uint32(v)>>16)
		o2 = LOP_IRR(OP_ORI, REGTMP, REGTMP, uint32(v))
		o3 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), REGTMP)
		if p.From.Sym != nil {
			c.ctxt.Diag("%v is not supported", p)
		}

	case 29:
		v := c.regoff(&p.From)

		d := c.vregoff(p.GetFrom3())
		var mask [2]uint8
		c.maskgen64(p, mask[:], uint64(d))
		var a int
		switch p.As {
		case ARLDC, ARLDCCC:
			a = int(mask[0])
			if int32(mask[1]) != (63 - v) {
				c.ctxt.Diag("invalid mask for shift: %x (shift %d)\n%v", uint64(d), v, p)
			}

		case ARLDCL, ARLDCLCC:
			a = int(mask[0])
			if mask[1] != 63 {
				c.ctxt.Diag("invalid mask for shift: %x (shift %d)\n%v", uint64(d), v, p)
			}

		case ARLDCR, ARLDCRCC:
			a = int(mask[1])
			if mask[0] != 0 {
				c.ctxt.Diag("invalid mask for shift: %x (shift %d)\n%v", uint64(d), v, p)
			}

		default:
			c.ctxt.Diag("unexpected op in rldic case\n%v", p)
			a = 0
		}

		o1 = AOP_RRR(c.opirr(p.As), uint32(p.Reg), uint32(p.To.Reg), (uint32(v) & 0x1F))
		o1 |= (uint32(a) & 31) << 6
		if v&0x20 != 0 {
			o1 |= 1 << 1
		}
		if a&0x20 != 0 {
			o1 |= 1 << 5
		}

	case 30:
		v := c.regoff(&p.From)

		d := c.vregoff(p.GetFrom3())

		switch p.As {
		case ARLDMI, ARLDMICC:
			var mask [2]uint8
			c.maskgen64(p, mask[:], uint64(d))
			if int32(mask[1]) != (63 - v) {
				c.ctxt.Diag("invalid mask for shift: %x (shift %d)\n%v", uint64(d), v, p)
			}
			o1 = AOP_RRR(c.opirr(p.As), uint32(p.Reg), uint32(p.To.Reg), (uint32(v) & 0x1F))
			o1 |= (uint32(mask[0]) & 31) << 6
			if v&0x20 != 0 {
				o1 |= 1 << 1
			}
			if mask[0]&0x20 != 0 {
				o1 |= 1 << 5
			}

		case ARLDIMI, ARLDIMICC:
			o1 = AOP_RRR(c.opirr(p.As), uint32(p.Reg), uint32(p.To.Reg), (uint32(v) & 0x1F))
			o1 |= (uint32(d) & 31) << 6
			if d&0x20 != 0 {
				o1 |= 1 << 5
			}
			if v&0x20 != 0 {
				o1 |= 1 << 1
			}
		}

	case 31:
		d := c.vregoff(&p.From)

		if c.ctxt.Arch.ByteOrder == binary.BigEndian {
			o1 = uint32(d >> 32)
			o2 = uint32(d)
		} else {
			o1 = uint32(d)
			o2 = uint32(d >> 32)
		}

		if p.From.Sym != nil {
			rel := obj.Addrel(c.cursym)
			rel.Off = int32(c.pc)
			rel.Siz = 8
			rel.Sym = p.From.Sym
			rel.Add = p.From.Offset
			rel.Type = objabi.R_ADDR
			o2 = 0
			o1 = o2
		}

	case 32:
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), 0) | (uint32(p.From.Reg)&31)<<6

	case 33:
		r := int(p.From.Reg)

		if oclass(&p.From) == C_NONE {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), 0, uint32(r))

	case 34:
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg)) | (uint32(p.GetFrom3().Reg)&31)<<6

	case 35:
		v := c.regoff(&p.To)

		r := int(p.To.Reg)
		if r == 0 {
			r = int(o.param)
		}

		inst := c.opstore(p.As)
		if c.opform(inst) == DS_FORM && v&0x3 != 0 {
			log.Fatalf("invalid offset for DS form load/store %v", p)
		}
		o1 = AOP_IRR(OP_ADDIS, REGTMP, uint32(r), uint32(high16adjusted(v)))
		o2 = AOP_IRR(inst, uint32(p.From.Reg), REGTMP, uint32(v))

	case 36:
		v := c.regoff(&p.From)

		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		o1 = AOP_IRR(OP_ADDIS, REGTMP, uint32(r), uint32(high16adjusted(v)))
		o2 = AOP_IRR(c.opload(p.As), uint32(p.To.Reg), REGTMP, uint32(v))

	case 37:
		v := c.regoff(&p.From)

		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		o1 = AOP_IRR(OP_ADDIS, REGTMP, uint32(r), uint32(high16adjusted(v)))
		o2 = AOP_IRR(c.opload(p.As), uint32(p.To.Reg), REGTMP, uint32(v))
		o3 = LOP_RRR(OP_EXTSB, uint32(p.To.Reg), uint32(p.To.Reg), 0)

	case 40:
		o1 = uint32(c.regoff(&p.From))

	case 41:
		o1 = AOP_RRR(c.opirr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), 0) | (uint32(c.regoff(p.GetFrom3()))&0x7F)<<11

	case 42:
		o1 = AOP_RRR(c.opirr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), 0) | (uint32(c.regoff(p.GetFrom3()))&0x7F)<<11

	case 43:

		if p.To.Type == obj.TYPE_NONE {
			o1 = AOP_RRR(c.oprrr(p.As), 0, uint32(p.From.Index), uint32(p.From.Reg))
		} else {
			th := c.regoff(&p.To)
			o1 = AOP_RRR(c.oprrr(p.As), uint32(th), uint32(p.From.Index), uint32(p.From.Reg))
		}

	case 44:
		o1 = AOP_RRR(c.opstorex(p.As), uint32(p.From.Reg), uint32(p.To.Index), uint32(p.To.Reg))

	case 45:
		switch p.As {

		case ALBAR, ALHAR, ALWAR, ALDAR:
			if p.From3Type() != obj.TYPE_NONE {
				eh := int(c.regoff(p.GetFrom3()))
				if eh > 1 {
					c.ctxt.Diag("illegal EH field\n%v", p)
				}
				o1 = AOP_RRRI(c.oploadx(p.As), uint32(p.To.Reg), uint32(p.From.Index), uint32(p.From.Reg), uint32(eh))
			} else {
				o1 = AOP_RRR(c.oploadx(p.As), uint32(p.To.Reg), uint32(p.From.Index), uint32(p.From.Reg))
			}
		default:
			o1 = AOP_RRR(c.oploadx(p.As), uint32(p.To.Reg), uint32(p.From.Index), uint32(p.From.Reg))
		}
	case 46:
		o1 = c.oprrr(p.As)

	case 47:
		r := int(p.From.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), 0)

	case 48:
		r := int(p.From.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = LOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), 0)

	case 49:
		if p.From.Type != obj.TYPE_REG {
			v := c.regoff(&p.From) & 1
			o1 = AOP_RRR(c.oprrr(p.As), 0, 0, uint32(p.To.Reg)) | uint32(v)<<21
		} else {
			o1 = AOP_RRR(c.oprrr(p.As), 0, 0, uint32(p.From.Reg))
		}

	case 50:
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		v := c.oprrr(p.As)
		t := v & (1<<10 | 1)
		o1 = AOP_RRR(v&^t, REGTMP, uint32(r), uint32(p.From.Reg))
		o2 = AOP_RRR(OP_MULLW, REGTMP, REGTMP, uint32(p.From.Reg))
		o3 = AOP_RRR(OP_SUBF|t, uint32(p.To.Reg), REGTMP, uint32(r))
		if p.As == AREMU {
			o4 = o3

			o3 = OP_RLW(OP_RLDIC, REGTMP, REGTMP, 0, 0, 0) | 1<<5
		}

	case 51:
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		v := c.oprrr(p.As)
		t := v & (1<<10 | 1)
		o1 = AOP_RRR(v&^t, REGTMP, uint32(r), uint32(p.From.Reg))
		o2 = AOP_RRR(OP_MULLD, REGTMP, REGTMP, uint32(p.From.Reg))
		o3 = AOP_RRR(OP_SUBF|t, uint32(p.To.Reg), REGTMP, uint32(r))

	case 52:
		v := c.regoff(&p.From) & 31

		o1 = AOP_RRR(c.oprrr(p.As), uint32(v), 0, 0)

	case 53:
		o1 = AOP_RRR(OP_MFFS, uint32(p.To.Reg), 0, 0)

	case 54:
		if oclass(&p.From) == C_REG {
			if p.As == AMOVD {
				o1 = AOP_RRR(OP_MTMSRD, uint32(p.From.Reg), 0, 0)
			} else {
				o1 = AOP_RRR(OP_MTMSR, uint32(p.From.Reg), 0, 0)
			}
		} else {
			o1 = AOP_RRR(OP_MFMSR, uint32(p.To.Reg), 0, 0)
		}

	case 55:
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), 0, uint32(p.From.Reg))

	case 56:
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.opirr(p.As), uint32(r), uint32(p.To.Reg), uint32(v)&31)
		if (p.As == ASRAD || p.As == ASRADCC) && (v&0x20 != 0) {
			o1 |= 1 << 1
		}

	case 57:
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}

		if v < 0 {
			v = 0
		} else if v > 32 {
			v = 32
		}
		var mask [2]uint8
		switch p.As {
		case AROTLW:
			mask[0], mask[1] = 0, 31
		case ASRW, ASRWCC:
			mask[0], mask[1] = uint8(v), 31
			v = 32 - v
		default:
			mask[0], mask[1] = 0, uint8(31-v)
		}
		o1 = OP_RLW(OP_RLWINM, uint32(p.To.Reg), uint32(r), uint32(v), uint32(mask[0]), uint32(mask[1]))
		if p.As == ASLWCC || p.As == ASRWCC {
			o1 |= 1
		}

	case 58:
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = LOP_IRR(c.opirr(p.As), uint32(p.To.Reg), uint32(r), uint32(v))

	case 59:
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		switch p.As {
		case AOR:
			o1 = LOP_IRR(c.opirr(AORIS), uint32(p.To.Reg), uint32(r), uint32(v)>>16)
		case AXOR:
			o1 = LOP_IRR(c.opirr(AXORIS), uint32(p.To.Reg), uint32(r), uint32(v)>>16)
		case AANDCC:
			o1 = LOP_IRR(c.opirr(AANDISCC), uint32(p.To.Reg), uint32(r), uint32(v)>>16)
		default:
			o1 = LOP_IRR(c.opirr(p.As), uint32(p.To.Reg), uint32(r), uint32(v))
		}

	case 60:
		r := int(c.regoff(&p.From) & 31)

		o1 = AOP_RRR(c.oprrr(p.As), uint32(r), uint32(p.Reg), uint32(p.To.Reg))

	case 61:
		r := int(c.regoff(&p.From) & 31)

		v := c.regoff(&p.To)
		o1 = AOP_IRR(c.opirr(p.As), uint32(r), uint32(p.Reg), uint32(v))

	case 62:
		v := c.regoff(&p.From)

		var mask [2]uint8
		c.maskgen(p, mask[:], uint32(c.regoff(p.GetFrom3())))
		o1 = AOP_RRR(c.opirr(p.As), uint32(p.Reg), uint32(p.To.Reg), uint32(v))
		o1 |= (uint32(mask[0])&31)<<6 | (uint32(mask[1])&31)<<1

	case 63:
		var mask [2]uint8
		c.maskgen(p, mask[:], uint32(c.regoff(p.GetFrom3())))

		o1 = AOP_RRR(c.opirr(p.As), uint32(p.Reg), uint32(p.To.Reg), uint32(p.From.Reg))
		o1 |= (uint32(mask[0])&31)<<6 | (uint32(mask[1])&31)<<1

	case 64:
		var v int32
		if p.From3Type() != obj.TYPE_NONE {
			v = c.regoff(p.GetFrom3()) & 255
		} else {
			v = 255
		}
		o1 = OP_MTFSF | uint32(v)<<17 | uint32(p.From.Reg)<<11

	case 65:
		if p.To.Reg == 0 {
			c.ctxt.Diag("must specify FPSCR(n)\n%v", p)
		}
		o1 = OP_MTFSFI | (uint32(p.To.Reg)&15)<<23 | (uint32(c.regoff(&p.From))&31)<<12

	case 66:
		var r int
		var v int32
		if REG_R0 <= p.From.Reg && p.From.Reg <= REG_R31 {
			r = int(p.From.Reg)
			v = int32(p.To.Reg)
			if REG_DCR0 <= v && v <= REG_DCR0+1023 {
				o1 = OPVCC(31, 451, 0, 0)
			} else {
				o1 = OPVCC(31, 467, 0, 0)
			}
		} else {
			r = int(p.To.Reg)
			v = int32(p.From.Reg)
			if REG_DCR0 <= v && v <= REG_DCR0+1023 {
				o1 = OPVCC(31, 323, 0, 0)
			} else {
				o1 = OPVCC(31, 339, 0, 0)
			}
		}

		o1 = AOP_RRR(o1, uint32(r), 0, 0) | (uint32(v)&0x1f)<<16 | ((uint32(v)>>5)&0x1f)<<11

	case 67:
		if p.From.Type != obj.TYPE_REG || p.From.Reg < REG_CR0 || REG_CR7 < p.From.Reg || p.To.Type != obj.TYPE_REG || p.To.Reg < REG_CR0 || REG_CR7 < p.To.Reg {
			c.ctxt.Diag("illegal CR field number\n%v", p)
		}
		o1 = AOP_RRR(OP_MCRF, ((uint32(p.To.Reg) & 7) << 2), ((uint32(p.From.Reg) & 7) << 2), 0)

	case 68:
		if p.From.Type == obj.TYPE_REG && REG_CR0 <= p.From.Reg && p.From.Reg <= REG_CR7 {
			v := int32(1 << uint(7-(p.To.Reg&7)))
			o1 = AOP_RRR(OP_MFCR, uint32(p.To.Reg), 0, 0) | 1<<20 | uint32(v)<<12
		} else {
			o1 = AOP_RRR(OP_MFCR, uint32(p.To.Reg), 0, 0)
		}

	case 69:
		var v int32
		if p.From3Type() != obj.TYPE_NONE {
			if p.To.Reg != 0 {
				c.ctxt.Diag("can't use both mask and CR(n)\n%v", p)
			}
			v = c.regoff(p.GetFrom3()) & 0xff
		} else {
			if p.To.Reg == 0 {
				v = 0xff
			} else {
				v = 1 << uint(7-(p.To.Reg&7))
			}
		}

		o1 = AOP_RRR(OP_MTCRF, uint32(p.From.Reg), 0, 0) | uint32(v)<<12

	case 70:
		var r int
		if p.Reg == 0 {
			r = 0
		} else {
			r = (int(p.Reg) & 7) << 2
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(r), uint32(p.From.Reg), uint32(p.To.Reg))

	case 71:
		var r int
		if p.Reg == 0 {
			r = 0
		} else {
			r = (int(p.Reg) & 7) << 2
		}
		o1 = AOP_RRR(c.opirr(p.As), uint32(r), uint32(p.From.Reg), 0) | uint32(c.regoff(&p.To))&0xffff

	case 72:
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.From.Reg), 0, uint32(p.To.Reg))

	case 73:
		if p.From.Type != obj.TYPE_REG || p.From.Reg != REG_FPSCR || p.To.Type != obj.TYPE_REG || p.To.Reg < REG_CR0 || REG_CR7 < p.To.Reg {
			c.ctxt.Diag("illegal FPSCR/CR field number\n%v", p)
		}
		o1 = AOP_RRR(OP_MCRFS, ((uint32(p.To.Reg) & 7) << 2), ((0 & 7) << 2), 0)

	case 77:
		if p.From.Type == obj.TYPE_CONST {
			if p.From.Offset > BIG || p.From.Offset < -BIG {
				c.ctxt.Diag("illegal syscall, sysnum too large: %v", p)
			}
			o1 = AOP_IRR(OP_ADDI, REGZERO, REGZERO, uint32(p.From.Offset))
		} else if p.From.Type == obj.TYPE_REG {
			o1 = LOP_RRR(OP_OR, REGZERO, uint32(p.From.Reg), uint32(p.From.Reg))
		} else {
			c.ctxt.Diag("illegal syscall: %v", p)
			o1 = 0x7fe00008
		}

		o2 = c.oprrr(p.As)
		o3 = AOP_RRR(c.oprrr(AXOR), REGZERO, REGZERO, REGZERO)

	case 78:
		o1 = 0

	case 74:
		v := c.vregoff(&p.To)

		inst := c.opstore(p.As)
		if c.opform(inst) == DS_FORM && v&0x3 != 0 {
			log.Fatalf("invalid offset for DS form load/store %v", p)
		}
		o1, o2 = c.symbolAccess(p.To.Sym, v, p.From.Reg, inst)

	case 75:
		v := c.vregoff(&p.From)

		inst := c.opload(p.As)
		if c.opform(inst) == DS_FORM && v&0x3 != 0 {
			log.Fatalf("invalid offset for DS form load/store %v", p)
		}
		o1, o2 = c.symbolAccess(p.From.Sym, v, p.To.Reg, inst)

	case 76:
		v := c.vregoff(&p.From)

		inst := c.opload(p.As)
		if c.opform(inst) == DS_FORM && v&0x3 != 0 {
			log.Fatalf("invalid offset for DS form load/store %v", p)
		}
		o1, o2 = c.symbolAccess(p.From.Sym, v, p.To.Reg, inst)
		o3 = LOP_RRR(OP_EXTSB, uint32(p.To.Reg), uint32(p.To.Reg), 0)

	case 79:
		if p.From.Offset != 0 {
			c.ctxt.Diag("invalid offset against tls var %v", p)
		}
		o1 = AOP_IRR(OP_ADDI, uint32(p.To.Reg), REGZERO, 0)
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 4
		rel.Sym = p.From.Sym
		rel.Type = objabi.R_POWER_TLS_LE

	case 80:
		if p.From.Offset != 0 {
			c.ctxt.Diag("invalid offset against tls var %v", p)
		}
		o1 = AOP_IRR(OP_ADDIS, uint32(p.To.Reg), REG_R2, 0)
		o2 = AOP_IRR(c.opload(AMOVD), uint32(p.To.Reg), uint32(p.To.Reg), 0)
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 8
		rel.Sym = p.From.Sym
		rel.Type = objabi.R_POWER_TLS_IE

	case 81:
		v := c.vregoff(&p.To)
		if v != 0 {
			c.ctxt.Diag("invalid offset against GOT slot %v", p)
		}

		o1 = AOP_IRR(OP_ADDIS, uint32(p.To.Reg), REG_R2, 0)
		o2 = AOP_IRR(c.opload(AMOVD), uint32(p.To.Reg), uint32(p.To.Reg), 0)
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 8
		rel.Sym = p.From.Sym
		rel.Type = objabi.R_ADDRPOWER_GOT
	case 82:
		if p.From.Type == obj.TYPE_REG {

			o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg))
		} else if p.From3Type() == obj.TYPE_CONST {

			six := int(c.regoff(&p.From))
			st := int(c.regoff(p.GetFrom3()))
			o1 = AOP_IIRR(c.opiirr(p.As), uint32(p.To.Reg), uint32(p.Reg), uint32(st), uint32(six))
		} else if p.From3Type() == obj.TYPE_NONE && p.Reg != 0 {

			uim := int(c.regoff(&p.From))
			o1 = AOP_VIRR(c.opirr(p.As), uint32(p.To.Reg), uint32(p.Reg), uint32(uim))
		} else {

			sim := int(c.regoff(&p.From))
			o1 = AOP_IR(c.opirr(p.As), uint32(p.To.Reg), uint32(sim))
		}

	case 83:
		if p.From.Type == obj.TYPE_REG {

			o1 = AOP_RRRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg), uint32(p.GetFrom3().Reg))
		} else if p.From.Type == obj.TYPE_CONST {

			shb := int(c.regoff(&p.From))
			o1 = AOP_IRRR(c.opirrr(p.As), uint32(p.To.Reg), uint32(p.Reg), uint32(p.GetFrom3().Reg), uint32(shb))
		}

	case 84:
		bc := c.vregoff(&p.From)

		o1 = AOP_ISEL(OP_ISEL, uint32(p.To.Reg), uint32(p.Reg), uint32(p.GetFrom3().Reg), uint32(bc))

	case 85:

		o1 = AOP_RR(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg))

	case 86:

		o1 = AOP_XX1(c.opstorex(p.As), uint32(p.From.Reg), uint32(p.To.Index), uint32(p.To.Reg))

	case 87:

		o1 = AOP_XX1(c.oploadx(p.As), uint32(p.To.Reg), uint32(p.From.Index), uint32(p.From.Reg))

	case 88:

		xt := int32(p.To.Reg)
		xs := int32(p.From.Reg)

		if REG_V0 <= xt && xt <= REG_V31 {

			xt = xt + 64
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg))
		} else if REG_F0 <= xt && xt <= REG_F31 {

			xt = xt + 64
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg))
		} else if REG_VS0 <= xt && xt <= REG_VS63 {
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg))
		} else if REG_V0 <= xs && xs <= REG_V31 {

			xs = xs + 64
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), uint32(p.Reg))
		} else if REG_F0 <= xs && xs <= REG_F31 {
			xs = xs + 64
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), uint32(p.Reg))
		} else if REG_VS0 <= xs && xs <= REG_VS63 {
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), uint32(p.Reg))
		}

	case 89:

		uim := int(c.regoff(p.GetFrom3()))
		o1 = AOP_XX2(c.oprrr(p.As), uint32(p.To.Reg), uint32(uim), uint32(p.From.Reg))

	case 90:
		if p.From3Type() == obj.TYPE_NONE {

			o1 = AOP_XX3(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg))
		} else if p.From3Type() == obj.TYPE_CONST {

			dm := int(c.regoff(p.GetFrom3()))
			o1 = AOP_XX3I(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg), uint32(dm))
		}

	case 91:

		o1 = AOP_XX4(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg), uint32(p.GetFrom3().Reg))

	case 92:
		if p.To.Type == obj.TYPE_CONST {

			xf := int32(p.From.Reg)
			if REG_F0 <= xf && xf <= REG_F31 {

				bf := int(c.regoff(&p.To)) << 2
				o1 = AOP_RRR(c.opirr(p.As), uint32(bf), uint32(p.From.Reg), uint32(p.Reg))
			} else {

				l := int(c.regoff(&p.To))
				o1 = AOP_RRR(c.opirr(p.As), uint32(l), uint32(p.From.Reg), uint32(p.Reg))
			}
		} else if p.From3Type() == obj.TYPE_CONST {

			l := int(c.regoff(p.GetFrom3()))
			o1 = AOP_RRR(c.opirr(p.As), uint32(l), uint32(p.To.Reg), uint32(p.From.Reg))
		} else if p.To.Type == obj.TYPE_REG {
			cr := int32(p.To.Reg)
			if REG_CR0 <= cr && cr <= REG_CR7 {

				bf := (int(p.To.Reg) & 7) << 2
				o1 = AOP_RRR(c.opirr(p.As), uint32(bf), uint32(p.From.Reg), uint32(p.Reg))
			} else if p.From.Type == obj.TYPE_CONST {

				l := int(c.regoff(&p.From))
				o1 = AOP_RRR(c.opirr(p.As), uint32(p.To.Reg), uint32(l), uint32(p.Reg))
			} else {
				switch p.As {
				case ACOPY, APASTECC:
					o1 = AOP_RRR(c.opirr(p.As), uint32(1), uint32(p.From.Reg), uint32(p.To.Reg))
				default:

					o1 = AOP_RRR(c.oprrr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), uint32(p.Reg))
				}
			}
		}

	case 93:
		if p.To.Type == obj.TYPE_CONST {

			bf := int(c.regoff(&p.To)) << 2
			o1 = AOP_RR(c.opirr(p.As), uint32(bf), uint32(p.From.Reg))
		} else if p.Reg == 0 {

			o1 = AOP_RRR(c.oprrr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), uint32(p.Reg))
		}

	case 94:

		cy := int(c.regoff(p.GetFrom3()))
		o1 = AOP_Z23I(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg), uint32(cy))
	}

	out[0] = o1
	out[1] = o2
	out[2] = o3
	out[3] = o4
	out[4] = o5
}

func (c *ctxt9) vregoff(a *obj.Addr) int64 {
	c.instoffset = 0
	if a != nil {
		c.aclass(a)
	}
	return c.instoffset
}

func (c *ctxt9) regoff(a *obj.Addr) int32 {
	return int32(c.vregoff(a))
}

func (c *ctxt9) oprrr(a obj.As) uint32 {
	switch a {
	case AADD:
		return OPVCC(31, 266, 0, 0)
	case AADDCC:
		return OPVCC(31, 266, 0, 1)
	case AADDV:
		return OPVCC(31, 266, 1, 0)
	case AADDVCC:
		return OPVCC(31, 266, 1, 1)
	case AADDC:
		return OPVCC(31, 10, 0, 0)
	case AADDCCC:
		return OPVCC(31, 10, 0, 1)
	case AADDCV:
		return OPVCC(31, 10, 1, 0)
	case AADDCVCC:
		return OPVCC(31, 10, 1, 1)
	case AADDE:
		return OPVCC(31, 138, 0, 0)
	case AADDECC:
		return OPVCC(31, 138, 0, 1)
	case AADDEV:
		return OPVCC(31, 138, 1, 0)
	case AADDEVCC:
		return OPVCC(31, 138, 1, 1)
	case AADDME:
		return OPVCC(31, 234, 0, 0)
	case AADDMECC:
		return OPVCC(31, 234, 0, 1)
	case AADDMEV:
		return OPVCC(31, 234, 1, 0)
	case AADDMEVCC:
		return OPVCC(31, 234, 1, 1)
	case AADDZE:
		return OPVCC(31, 202, 0, 0)
	case AADDZECC:
		return OPVCC(31, 202, 0, 1)
	case AADDZEV:
		return OPVCC(31, 202, 1, 0)
	case AADDZEVCC:
		return OPVCC(31, 202, 1, 1)
	case AADDEX:
		return OPVCC(31, 170, 0, 0)

	case AAND:
		return OPVCC(31, 28, 0, 0)
	case AANDCC:
		return OPVCC(31, 28, 0, 1)
	case AANDN:
		return OPVCC(31, 60, 0, 0)
	case AANDNCC:
		return OPVCC(31, 60, 0, 1)

	case ACMP:
		return OPVCC(31, 0, 0, 0) | 1<<21
	case ACMPU:
		return OPVCC(31, 32, 0, 0) | 1<<21
	case ACMPW:
		return OPVCC(31, 0, 0, 0)
	case ACMPWU:
		return OPVCC(31, 32, 0, 0)
	case ACMPB:
		return OPVCC(31, 508, 0, 0)

	case ACNTLZW:
		return OPVCC(31, 26, 0, 0)
	case ACNTLZWCC:
		return OPVCC(31, 26, 0, 1)
	case ACNTLZD:
		return OPVCC(31, 58, 0, 0)
	case ACNTLZDCC:
		return OPVCC(31, 58, 0, 1)

	case ACRAND:
		return OPVCC(19, 257, 0, 0)
	case ACRANDN:
		return OPVCC(19, 129, 0, 0)
	case ACREQV:
		return OPVCC(19, 289, 0, 0)
	case ACRNAND:
		return OPVCC(19, 225, 0, 0)
	case ACRNOR:
		return OPVCC(19, 33, 0, 0)
	case ACROR:
		return OPVCC(19, 449, 0, 0)
	case ACRORN:
		return OPVCC(19, 417, 0, 0)
	case ACRXOR:
		return OPVCC(19, 193, 0, 0)

	case ADCBF:
		return OPVCC(31, 86, 0, 0)
	case ADCBI:
		return OPVCC(31, 470, 0, 0)
	case ADCBST:
		return OPVCC(31, 54, 0, 0)
	case ADCBT:
		return OPVCC(31, 278, 0, 0)
	case ADCBTST:
		return OPVCC(31, 246, 0, 0)
	case ADCBZ:
		return OPVCC(31, 1014, 0, 0)

	case AREM, ADIVW:
		return OPVCC(31, 491, 0, 0)

	case AREMCC, ADIVWCC:
		return OPVCC(31, 491, 0, 1)

	case AREMV, ADIVWV:
		return OPVCC(31, 491, 1, 0)

	case AREMVCC, ADIVWVCC:
		return OPVCC(31, 491, 1, 1)

	case AREMU, ADIVWU:
		return OPVCC(31, 459, 0, 0)

	case AREMUCC, ADIVWUCC:
		return OPVCC(31, 459, 0, 1)

	case AREMUV, ADIVWUV:
		return OPVCC(31, 459, 1, 0)

	case AREMUVCC, ADIVWUVCC:
		return OPVCC(31, 459, 1, 1)

	case AREMD, ADIVD:
		return OPVCC(31, 489, 0, 0)

	case AREMDCC, ADIVDCC:
		return OPVCC(31, 489, 0, 1)

	case ADIVDE:
		return OPVCC(31, 425, 0, 0)

	case ADIVDECC:
		return OPVCC(31, 425, 0, 1)

	case ADIVDEU:
		return OPVCC(31, 393, 0, 0)

	case ADIVDEUCC:
		return OPVCC(31, 393, 0, 1)

	case AREMDV, ADIVDV:
		return OPVCC(31, 489, 1, 0)

	case AREMDVCC, ADIVDVCC:
		return OPVCC(31, 489, 1, 1)

	case AREMDU, ADIVDU:
		return OPVCC(31, 457, 0, 0)

	case AREMDUCC, ADIVDUCC:
		return OPVCC(31, 457, 0, 1)

	case AREMDUV, ADIVDUV:
		return OPVCC(31, 457, 1, 0)

	case AREMDUVCC, ADIVDUVCC:
		return OPVCC(31, 457, 1, 1)

	case AEIEIO:
		return OPVCC(31, 854, 0, 0)

	case AEQV:
		return OPVCC(31, 284, 0, 0)
	case AEQVCC:
		return OPVCC(31, 284, 0, 1)

	case AEXTSB:
		return OPVCC(31, 954, 0, 0)
	case AEXTSBCC:
		return OPVCC(31, 954, 0, 1)
	case AEXTSH:
		return OPVCC(31, 922, 0, 0)
	case AEXTSHCC:
		return OPVCC(31, 922, 0, 1)
	case AEXTSW:
		return OPVCC(31, 986, 0, 0)
	case AEXTSWCC:
		return OPVCC(31, 986, 0, 1)

	case AFABS:
		return OPVCC(63, 264, 0, 0)
	case AFABSCC:
		return OPVCC(63, 264, 0, 1)
	case AFADD:
		return OPVCC(63, 21, 0, 0)
	case AFADDCC:
		return OPVCC(63, 21, 0, 1)
	case AFADDS:
		return OPVCC(59, 21, 0, 0)
	case AFADDSCC:
		return OPVCC(59, 21, 0, 1)
	case AFCMPO:
		return OPVCC(63, 32, 0, 0)
	case AFCMPU:
		return OPVCC(63, 0, 0, 0)
	case AFCFID:
		return OPVCC(63, 846, 0, 0)
	case AFCFIDCC:
		return OPVCC(63, 846, 0, 1)
	case AFCFIDU:
		return OPVCC(63, 974, 0, 0)
	case AFCFIDUCC:
		return OPVCC(63, 974, 0, 1)
	case AFCFIDS:
		return OPVCC(59, 846, 0, 0)
	case AFCFIDSCC:
		return OPVCC(59, 846, 0, 1)
	case AFCTIW:
		return OPVCC(63, 14, 0, 0)
	case AFCTIWCC:
		return OPVCC(63, 14, 0, 1)
	case AFCTIWZ:
		return OPVCC(63, 15, 0, 0)
	case AFCTIWZCC:
		return OPVCC(63, 15, 0, 1)
	case AFCTID:
		return OPVCC(63, 814, 0, 0)
	case AFCTIDCC:
		return OPVCC(63, 814, 0, 1)
	case AFCTIDZ:
		return OPVCC(63, 815, 0, 0)
	case AFCTIDZCC:
		return OPVCC(63, 815, 0, 1)
	case AFDIV:
		return OPVCC(63, 18, 0, 0)
	case AFDIVCC:
		return OPVCC(63, 18, 0, 1)
	case AFDIVS:
		return OPVCC(59, 18, 0, 0)
	case AFDIVSCC:
		return OPVCC(59, 18, 0, 1)
	case AFMADD:
		return OPVCC(63, 29, 0, 0)
	case AFMADDCC:
		return OPVCC(63, 29, 0, 1)
	case AFMADDS:
		return OPVCC(59, 29, 0, 0)
	case AFMADDSCC:
		return OPVCC(59, 29, 0, 1)

	case AFMOVS, AFMOVD:
		return OPVCC(63, 72, 0, 0)
	case AFMOVDCC:
		return OPVCC(63, 72, 0, 1)
	case AFMSUB:
		return OPVCC(63, 28, 0, 0)
	case AFMSUBCC:
		return OPVCC(63, 28, 0, 1)
	case AFMSUBS:
		return OPVCC(59, 28, 0, 0)
	case AFMSUBSCC:
		return OPVCC(59, 28, 0, 1)
	case AFMUL:
		return OPVCC(63, 25, 0, 0)
	case AFMULCC:
		return OPVCC(63, 25, 0, 1)
	case AFMULS:
		return OPVCC(59, 25, 0, 0)
	case AFMULSCC:
		return OPVCC(59, 25, 0, 1)
	case AFNABS:
		return OPVCC(63, 136, 0, 0)
	case AFNABSCC:
		return OPVCC(63, 136, 0, 1)
	case AFNEG:
		return OPVCC(63, 40, 0, 0)
	case AFNEGCC:
		return OPVCC(63, 40, 0, 1)
	case AFNMADD:
		return OPVCC(63, 31, 0, 0)
	case AFNMADDCC:
		return OPVCC(63, 31, 0, 1)
	case AFNMADDS:
		return OPVCC(59, 31, 0, 0)
	case AFNMADDSCC:
		return OPVCC(59, 31, 0, 1)
	case AFNMSUB:
		return OPVCC(63, 30, 0, 0)
	case AFNMSUBCC:
		return OPVCC(63, 30, 0, 1)
	case AFNMSUBS:
		return OPVCC(59, 30, 0, 0)
	case AFNMSUBSCC:
		return OPVCC(59, 30, 0, 1)
	case AFCPSGN:
		return OPVCC(63, 8, 0, 0)
	case AFCPSGNCC:
		return OPVCC(63, 8, 0, 1)
	case AFRES:
		return OPVCC(59, 24, 0, 0)
	case AFRESCC:
		return OPVCC(59, 24, 0, 1)
	case AFRIM:
		return OPVCC(63, 488, 0, 0)
	case AFRIMCC:
		return OPVCC(63, 488, 0, 1)
	case AFRIP:
		return OPVCC(63, 456, 0, 0)
	case AFRIPCC:
		return OPVCC(63, 456, 0, 1)
	case AFRIZ:
		return OPVCC(63, 424, 0, 0)
	case AFRIZCC:
		return OPVCC(63, 424, 0, 1)
	case AFRIN:
		return OPVCC(63, 392, 0, 0)
	case AFRINCC:
		return OPVCC(63, 392, 0, 1)
	case AFRSP:
		return OPVCC(63, 12, 0, 0)
	case AFRSPCC:
		return OPVCC(63, 12, 0, 1)
	case AFRSQRTE:
		return OPVCC(63, 26, 0, 0)
	case AFRSQRTECC:
		return OPVCC(63, 26, 0, 1)
	case AFSEL:
		return OPVCC(63, 23, 0, 0)
	case AFSELCC:
		return OPVCC(63, 23, 0, 1)
	case AFSQRT:
		return OPVCC(63, 22, 0, 0)
	case AFSQRTCC:
		return OPVCC(63, 22, 0, 1)
	case AFSQRTS:
		return OPVCC(59, 22, 0, 0)
	case AFSQRTSCC:
		return OPVCC(59, 22, 0, 1)
	case AFSUB:
		return OPVCC(63, 20, 0, 0)
	case AFSUBCC:
		return OPVCC(63, 20, 0, 1)
	case AFSUBS:
		return OPVCC(59, 20, 0, 0)
	case AFSUBSCC:
		return OPVCC(59, 20, 0, 1)

	case AICBI:
		return OPVCC(31, 982, 0, 0)
	case AISYNC:
		return OPVCC(19, 150, 0, 0)

	case AMTFSB0:
		return OPVCC(63, 70, 0, 0)
	case AMTFSB0CC:
		return OPVCC(63, 70, 0, 1)
	case AMTFSB1:
		return OPVCC(63, 38, 0, 0)
	case AMTFSB1CC:
		return OPVCC(63, 38, 0, 1)

	case AMULHW:
		return OPVCC(31, 75, 0, 0)
	case AMULHWCC:
		return OPVCC(31, 75, 0, 1)
	case AMULHWU:
		return OPVCC(31, 11, 0, 0)
	case AMULHWUCC:
		return OPVCC(31, 11, 0, 1)
	case AMULLW:
		return OPVCC(31, 235, 0, 0)
	case AMULLWCC:
		return OPVCC(31, 235, 0, 1)
	case AMULLWV:
		return OPVCC(31, 235, 1, 0)
	case AMULLWVCC:
		return OPVCC(31, 235, 1, 1)

	case AMULHD:
		return OPVCC(31, 73, 0, 0)
	case AMULHDCC:
		return OPVCC(31, 73, 0, 1)
	case AMULHDU:
		return OPVCC(31, 9, 0, 0)
	case AMULHDUCC:
		return OPVCC(31, 9, 0, 1)
	case AMULLD:
		return OPVCC(31, 233, 0, 0)
	case AMULLDCC:
		return OPVCC(31, 233, 0, 1)
	case AMULLDV:
		return OPVCC(31, 233, 1, 0)
	case AMULLDVCC:
		return OPVCC(31, 233, 1, 1)

	case ANAND:
		return OPVCC(31, 476, 0, 0)
	case ANANDCC:
		return OPVCC(31, 476, 0, 1)
	case ANEG:
		return OPVCC(31, 104, 0, 0)
	case ANEGCC:
		return OPVCC(31, 104, 0, 1)
	case ANEGV:
		return OPVCC(31, 104, 1, 0)
	case ANEGVCC:
		return OPVCC(31, 104, 1, 1)
	case ANOR:
		return OPVCC(31, 124, 0, 0)
	case ANORCC:
		return OPVCC(31, 124, 0, 1)
	case AOR:
		return OPVCC(31, 444, 0, 0)
	case AORCC:
		return OPVCC(31, 444, 0, 1)
	case AORN:
		return OPVCC(31, 412, 0, 0)
	case AORNCC:
		return OPVCC(31, 412, 0, 1)

	case APOPCNTD:
		return OPVCC(31, 506, 0, 0)
	case APOPCNTW:
		return OPVCC(31, 378, 0, 0)
	case APOPCNTB:
		return OPVCC(31, 122, 0, 0)

	case ARFI:
		return OPVCC(19, 50, 0, 0)
	case ARFCI:
		return OPVCC(19, 51, 0, 0)
	case ARFID:
		return OPVCC(19, 18, 0, 0)
	case AHRFID:
		return OPVCC(19, 274, 0, 0)

	case ARLWMI:
		return OPVCC(20, 0, 0, 0)
	case ARLWMICC:
		return OPVCC(20, 0, 0, 1)
	case ARLWNM:
		return OPVCC(23, 0, 0, 0)
	case ARLWNMCC:
		return OPVCC(23, 0, 0, 1)

	case ARLDCL:
		return OPVCC(30, 8, 0, 0)
	case ARLDCR:
		return OPVCC(30, 9, 0, 0)

	case ARLDICL:
		return OPVCC(30, 0, 0, 0)
	case ARLDICLCC:
		return OPVCC(30, 0, 0, 1)
	case ARLDICR:
		return OPVCC(30, 0, 0, 0) | 2<<1
	case ARLDICRCC:
		return OPVCC(30, 0, 0, 1) | 2<<1

	case ASYSCALL:
		return OPVCC(17, 1, 0, 0)

	case ASLW:
		return OPVCC(31, 24, 0, 0)
	case ASLWCC:
		return OPVCC(31, 24, 0, 1)
	case ASLD:
		return OPVCC(31, 27, 0, 0)
	case ASLDCC:
		return OPVCC(31, 27, 0, 1)

	case ASRAW:
		return OPVCC(31, 792, 0, 0)
	case ASRAWCC:
		return OPVCC(31, 792, 0, 1)
	case ASRAD:
		return OPVCC(31, 794, 0, 0)
	case ASRADCC:
		return OPVCC(31, 794, 0, 1)

	case ASRW:
		return OPVCC(31, 536, 0, 0)
	case ASRWCC:
		return OPVCC(31, 536, 0, 1)
	case ASRD:
		return OPVCC(31, 539, 0, 0)
	case ASRDCC:
		return OPVCC(31, 539, 0, 1)

	case ASUB:
		return OPVCC(31, 40, 0, 0)
	case ASUBCC:
		return OPVCC(31, 40, 0, 1)
	case ASUBV:
		return OPVCC(31, 40, 1, 0)
	case ASUBVCC:
		return OPVCC(31, 40, 1, 1)
	case ASUBC:
		return OPVCC(31, 8, 0, 0)
	case ASUBCCC:
		return OPVCC(31, 8, 0, 1)
	case ASUBCV:
		return OPVCC(31, 8, 1, 0)
	case ASUBCVCC:
		return OPVCC(31, 8, 1, 1)
	case ASUBE:
		return OPVCC(31, 136, 0, 0)
	case ASUBECC:
		return OPVCC(31, 136, 0, 1)
	case ASUBEV:
		return OPVCC(31, 136, 1, 0)
	case ASUBEVCC:
		return OPVCC(31, 136, 1, 1)
	case ASUBME:
		return OPVCC(31, 232, 0, 0)
	case ASUBMECC:
		return OPVCC(31, 232, 0, 1)
	case ASUBMEV:
		return OPVCC(31, 232, 1, 0)
	case ASUBMEVCC:
		return OPVCC(31, 232, 1, 1)
	case ASUBZE:
		return OPVCC(31, 200, 0, 0)
	case ASUBZECC:
		return OPVCC(31, 200, 0, 1)
	case ASUBZEV:
		return OPVCC(31, 200, 1, 0)
	case ASUBZEVCC:
		return OPVCC(31, 200, 1, 1)

	case ASYNC:
		return OPVCC(31, 598, 0, 0)
	case ALWSYNC:
		return OPVCC(31, 598, 0, 0) | 1<<21

	case APTESYNC:
		return OPVCC(31, 598, 0, 0) | 2<<21

	case ATLBIE:
		return OPVCC(31, 306, 0, 0)
	case ATLBIEL:
		return OPVCC(31, 274, 0, 0)
	case ATLBSYNC:
		return OPVCC(31, 566, 0, 0)
	case ASLBIA:
		return OPVCC(31, 498, 0, 0)
	case ASLBIE:
		return OPVCC(31, 434, 0, 0)
	case ASLBMFEE:
		return OPVCC(31, 915, 0, 0)
	case ASLBMFEV:
		return OPVCC(31, 851, 0, 0)
	case ASLBMTE:
		return OPVCC(31, 402, 0, 0)

	case ATW:
		return OPVCC(31, 4, 0, 0)
	case ATD:
		return OPVCC(31, 68, 0, 0)

	case AVAND:
		return OPVX(4, 1028, 0, 0)
	case AVANDC:
		return OPVX(4, 1092, 0, 0)
	case AVNAND:
		return OPVX(4, 1412, 0, 0)

	case AVOR:
		return OPVX(4, 1156, 0, 0)
	case AVORC:
		return OPVX(4, 1348, 0, 0)
	case AVNOR:
		return OPVX(4, 1284, 0, 0)
	case AVXOR:
		return OPVX(4, 1220, 0, 0)
	case AVEQV:
		return OPVX(4, 1668, 0, 0)

	case AVADDUBM:
		return OPVX(4, 0, 0, 0)
	case AVADDUHM:
		return OPVX(4, 64, 0, 0)
	case AVADDUWM:
		return OPVX(4, 128, 0, 0)
	case AVADDUDM:
		return OPVX(4, 192, 0, 0)
	case AVADDUQM:
		return OPVX(4, 256, 0, 0)

	case AVADDCUQ:
		return OPVX(4, 320, 0, 0)
	case AVADDCUW:
		return OPVX(4, 384, 0, 0)

	case AVADDUBS:
		return OPVX(4, 512, 0, 0)
	case AVADDUHS:
		return OPVX(4, 576, 0, 0)
	case AVADDUWS:
		return OPVX(4, 640, 0, 0)

	case AVADDSBS:
		return OPVX(4, 768, 0, 0)
	case AVADDSHS:
		return OPVX(4, 832, 0, 0)
	case AVADDSWS:
		return OPVX(4, 896, 0, 0)

	case AVADDEUQM:
		return OPVX(4, 60, 0, 0)
	case AVADDECUQ:
		return OPVX(4, 61, 0, 0)

	case AVMULESB:
		return OPVX(4, 776, 0, 0)
	case AVMULOSB:
		return OPVX(4, 264, 0, 0)
	case AVMULEUB:
		return OPVX(4, 520, 0, 0)
	case AVMULOUB:
		return OPVX(4, 8, 0, 0)
	case AVMULESH:
		return OPVX(4, 840, 0, 0)
	case AVMULOSH:
		return OPVX(4, 328, 0, 0)
	case AVMULEUH:
		return OPVX(4, 584, 0, 0)
	case AVMULOUH:
		return OPVX(4, 72, 0, 0)
	case AVMULESW:
		return OPVX(4, 904, 0, 0)
	case AVMULOSW:
		return OPVX(4, 392, 0, 0)
	case AVMULEUW:
		return OPVX(4, 648, 0, 0)
	case AVMULOUW:
		return OPVX(4, 136, 0, 0)
	case AVMULUWM:
		return OPVX(4, 137, 0, 0)

	case AVPMSUMB:
		return OPVX(4, 1032, 0, 0)
	case AVPMSUMH:
		return OPVX(4, 1096, 0, 0)
	case AVPMSUMW:
		return OPVX(4, 1160, 0, 0)
	case AVPMSUMD:
		return OPVX(4, 1224, 0, 0)

	case AVMSUMUDM:
		return OPVX(4, 35, 0, 0)

	case AVSUBUBM:
		return OPVX(4, 1024, 0, 0)
	case AVSUBUHM:
		return OPVX(4, 1088, 0, 0)
	case AVSUBUWM:
		return OPVX(4, 1152, 0, 0)
	case AVSUBUDM:
		return OPVX(4, 1216, 0, 0)
	case AVSUBUQM:
		return OPVX(4, 1280, 0, 0)

	case AVSUBCUQ:
		return OPVX(4, 1344, 0, 0)
	case AVSUBCUW:
		return OPVX(4, 1408, 0, 0)

	case AVSUBUBS:
		return OPVX(4, 1536, 0, 0)
	case AVSUBUHS:
		return OPVX(4, 1600, 0, 0)
	case AVSUBUWS:
		return OPVX(4, 1664, 0, 0)

	case AVSUBSBS:
		return OPVX(4, 1792, 0, 0)
	case AVSUBSHS:
		return OPVX(4, 1856, 0, 0)
	case AVSUBSWS:
		return OPVX(4, 1920, 0, 0)

	case AVSUBEUQM:
		return OPVX(4, 62, 0, 0)
	case AVSUBECUQ:
		return OPVX(4, 63, 0, 0)

	case AVRLB:
		return OPVX(4, 4, 0, 0)
	case AVRLH:
		return OPVX(4, 68, 0, 0)
	case AVRLW:
		return OPVX(4, 132, 0, 0)
	case AVRLD:
		return OPVX(4, 196, 0, 0)

	case AVSLB:
		return OPVX(4, 260, 0, 0)
	case AVSLH:
		return OPVX(4, 324, 0, 0)
	case AVSLW:
		return OPVX(4, 388, 0, 0)
	case AVSL:
		return OPVX(4, 452, 0, 0)
	case AVSLO:
		return OPVX(4, 1036, 0, 0)
	case AVSRB:
		return OPVX(4, 516, 0, 0)
	case AVSRH:
		return OPVX(4, 580, 0, 0)
	case AVSRW:
		return OPVX(4, 644, 0, 0)
	case AVSR:
		return OPVX(4, 708, 0, 0)
	case AVSRO:
		return OPVX(4, 1100, 0, 0)
	case AVSLD:
		return OPVX(4, 1476, 0, 0)
	case AVSRD:
		return OPVX(4, 1732, 0, 0)

	case AVSRAB:
		return OPVX(4, 772, 0, 0)
	case AVSRAH:
		return OPVX(4, 836, 0, 0)
	case AVSRAW:
		return OPVX(4, 900, 0, 0)
	case AVSRAD:
		return OPVX(4, 964, 0, 0)

	case AVBPERMQ:
		return OPVC(4, 1356, 0, 0)
	case AVBPERMD:
		return OPVC(4, 1484, 0, 0)

	case AVCLZB:
		return OPVX(4, 1794, 0, 0)
	case AVCLZH:
		return OPVX(4, 1858, 0, 0)
	case AVCLZW:
		return OPVX(4, 1922, 0, 0)
	case AVCLZD:
		return OPVX(4, 1986, 0, 0)

	case AVPOPCNTB:
		return OPVX(4, 1795, 0, 0)
	case AVPOPCNTH:
		return OPVX(4, 1859, 0, 0)
	case AVPOPCNTW:
		return OPVX(4, 1923, 0, 0)
	case AVPOPCNTD:
		return OPVX(4, 1987, 0, 0)

	case AVCMPEQUB:
		return OPVC(4, 6, 0, 0)
	case AVCMPEQUBCC:
		return OPVC(4, 6, 0, 1)
	case AVCMPEQUH:
		return OPVC(4, 70, 0, 0)
	case AVCMPEQUHCC:
		return OPVC(4, 70, 0, 1)
	case AVCMPEQUW:
		return OPVC(4, 134, 0, 0)
	case AVCMPEQUWCC:
		return OPVC(4, 134, 0, 1)
	case AVCMPEQUD:
		return OPVC(4, 199, 0, 0)
	case AVCMPEQUDCC:
		return OPVC(4, 199, 0, 1)

	case AVCMPGTUB:
		return OPVC(4, 518, 0, 0)
	case AVCMPGTUBCC:
		return OPVC(4, 518, 0, 1)
	case AVCMPGTUH:
		return OPVC(4, 582, 0, 0)
	case AVCMPGTUHCC:
		return OPVC(4, 582, 0, 1)
	case AVCMPGTUW:
		return OPVC(4, 646, 0, 0)
	case AVCMPGTUWCC:
		return OPVC(4, 646, 0, 1)
	case AVCMPGTUD:
		return OPVC(4, 711, 0, 0)
	case AVCMPGTUDCC:
		return OPVC(4, 711, 0, 1)
	case AVCMPGTSB:
		return OPVC(4, 774, 0, 0)
	case AVCMPGTSBCC:
		return OPVC(4, 774, 0, 1)
	case AVCMPGTSH:
		return OPVC(4, 838, 0, 0)
	case AVCMPGTSHCC:
		return OPVC(4, 838, 0, 1)
	case AVCMPGTSW:
		return OPVC(4, 902, 0, 0)
	case AVCMPGTSWCC:
		return OPVC(4, 902, 0, 1)
	case AVCMPGTSD:
		return OPVC(4, 967, 0, 0)
	case AVCMPGTSDCC:
		return OPVC(4, 967, 0, 1)

	case AVCMPNEZB:
		return OPVC(4, 263, 0, 0)
	case AVCMPNEZBCC:
		return OPVC(4, 263, 0, 1)

	case AVPERM:
		return OPVX(4, 43, 0, 0)

	case AVSEL:
		return OPVX(4, 42, 0, 0)

	case AVCIPHER:
		return OPVX(4, 1288, 0, 0)
	case AVCIPHERLAST:
		return OPVX(4, 1289, 0, 0)
	case AVNCIPHER:
		return OPVX(4, 1352, 0, 0)
	case AVNCIPHERLAST:
		return OPVX(4, 1353, 0, 0)
	case AVSBOX:
		return OPVX(4, 1480, 0, 0)

	case AMFVSRD, AMFVRD, AMFFPRD:
		return OPVXX1(31, 51, 0)
	case AMFVSRWZ:
		return OPVXX1(31, 115, 0)
	case AMFVSRLD:
		return OPVXX1(31, 307, 0)

	case AMTVSRD, AMTFPRD, AMTVRD:
		return OPVXX1(31, 179, 0)
	case AMTVSRWA:
		return OPVXX1(31, 211, 0)
	case AMTVSRWZ:
		return OPVXX1(31, 243, 0)
	case AMTVSRDD:
		return OPVXX1(31, 435, 0)
	case AMTVSRWS:
		return OPVXX1(31, 403, 0)

	case AXXLANDQ:
		return OPVXX3(60, 130, 0)
	case AXXLANDC:
		return OPVXX3(60, 138, 0)
	case AXXLEQV:
		return OPVXX3(60, 186, 0)
	case AXXLNAND:
		return OPVXX3(60, 178, 0)

	case AXXLORC:
		return OPVXX3(60, 170, 0)
	case AXXLNOR:
		return OPVXX3(60, 162, 0)
	case AXXLORQ:
		return OPVXX3(60, 146, 0)
	case AXXLXOR:
		return OPVXX3(60, 154, 0)

	case AXXSEL:
		return OPVXX4(60, 3, 0)

	case AXXMRGHW:
		return OPVXX3(60, 18, 0)
	case AXXMRGLW:
		return OPVXX3(60, 50, 0)

	case AXXSPLTW:
		return OPVXX2(60, 164, 0)

	case AXXPERMDI:
		return OPVXX3(60, 10, 0)

	case AXXSLDWI:
		return OPVXX3(60, 2, 0)

	case AXSCVDPSP:
		return OPVXX2(60, 265, 0)
	case AXSCVSPDP:
		return OPVXX2(60, 329, 0)
	case AXSCVDPSPN:
		return OPVXX2(60, 267, 0)
	case AXSCVSPDPN:
		return OPVXX2(60, 331, 0)

	case AXVCVDPSP:
		return OPVXX2(60, 393, 0)
	case AXVCVSPDP:
		return OPVXX2(60, 457, 0)

	case AXSCVDPSXDS:
		return OPVXX2(60, 344, 0)
	case AXSCVDPSXWS:
		return OPVXX2(60, 88, 0)
	case AXSCVDPUXDS:
		return OPVXX2(60, 328, 0)
	case AXSCVDPUXWS:
		return OPVXX2(60, 72, 0)

	case AXSCVSXDDP:
		return OPVXX2(60, 376, 0)
	case AXSCVUXDDP:
		return OPVXX2(60, 360, 0)
	case AXSCVSXDSP:
		return OPVXX2(60, 312, 0)
	case AXSCVUXDSP:
		return OPVXX2(60, 296, 0)

	case AXVCVDPSXDS:
		return OPVXX2(60, 472, 0)
	case AXVCVDPSXWS:
		return OPVXX2(60, 216, 0)
	case AXVCVDPUXDS:
		return OPVXX2(60, 456, 0)
	case AXVCVDPUXWS:
		return OPVXX2(60, 200, 0)
	case AXVCVSPSXDS:
		return OPVXX2(60, 408, 0)
	case AXVCVSPSXWS:
		return OPVXX2(60, 152, 0)
	case AXVCVSPUXDS:
		return OPVXX2(60, 392, 0)
	case AXVCVSPUXWS:
		return OPVXX2(60, 136, 0)

	case AXVCVSXDDP:
		return OPVXX2(60, 504, 0)
	case AXVCVSXWDP:
		return OPVXX2(60, 248, 0)
	case AXVCVUXDDP:
		return OPVXX2(60, 488, 0)
	case AXVCVUXWDP:
		return OPVXX2(60, 232, 0)
	case AXVCVSXDSP:
		return OPVXX2(60, 440, 0)
	case AXVCVSXWSP:
		return OPVXX2(60, 184, 0)
	case AXVCVUXDSP:
		return OPVXX2(60, 424, 0)
	case AXVCVUXWSP:
		return OPVXX2(60, 168, 0)

	case AMADDHD:
		return OPVX(4, 48, 0, 0)
	case AMADDHDU:
		return OPVX(4, 49, 0, 0)
	case AMADDLD:
		return OPVX(4, 51, 0, 0)

	case AXOR:
		return OPVCC(31, 316, 0, 0)
	case AXORCC:
		return OPVCC(31, 316, 0, 1)
	}

	c.ctxt.Diag("bad r/r, r/r/r or r/r/r/r opcode %v", a)
	return 0
}

func (c *ctxt9) opirrr(a obj.As) uint32 {
	switch a {

	case AVSLDOI:
		return OPVX(4, 44, 0, 0)
	}

	c.ctxt.Diag("bad i/r/r/r opcode %v", a)
	return 0
}

func (c *ctxt9) opiirr(a obj.As) uint32 {
	switch a {

	case AVSHASIGMAW:
		return OPVX(4, 1666, 0, 0)
	case AVSHASIGMAD:
		return OPVX(4, 1730, 0, 0)
	}

	c.ctxt.Diag("bad i/i/r/r opcode %v", a)
	return 0
}

func (c *ctxt9) opirr(a obj.As) uint32 {
	switch a {
	case AADD:
		return OPVCC(14, 0, 0, 0)
	case AADDC:
		return OPVCC(12, 0, 0, 0)
	case AADDCCC:
		return OPVCC(13, 0, 0, 0)
	case AADDIS:
		return OPVCC(15, 0, 0, 0)

	case AANDCC:
		return OPVCC(28, 0, 0, 0)
	case AANDISCC:
		return OPVCC(29, 0, 0, 0)

	case ABR:
		return OPVCC(18, 0, 0, 0)
	case ABL:
		return OPVCC(18, 0, 0, 0) | 1
	case obj.ADUFFZERO:
		return OPVCC(18, 0, 0, 0) | 1
	case obj.ADUFFCOPY:
		return OPVCC(18, 0, 0, 0) | 1
	case ABC:
		return OPVCC(16, 0, 0, 0)
	case ABCL:
		return OPVCC(16, 0, 0, 0) | 1

	case ABEQ:
		return AOP_RRR(16<<26, 12, 2, 0)
	case ABGE:
		return AOP_RRR(16<<26, 4, 0, 0)
	case ABGT:
		return AOP_RRR(16<<26, 12, 1, 0)
	case ABLE:
		return AOP_RRR(16<<26, 4, 1, 0)
	case ABLT:
		return AOP_RRR(16<<26, 12, 0, 0)
	case ABNE:
		return AOP_RRR(16<<26, 4, 2, 0)
	case ABVC:
		return AOP_RRR(16<<26, 4, 3, 0)
	case ABVS:
		return AOP_RRR(16<<26, 12, 3, 0)

	case ACMP:
		return OPVCC(11, 0, 0, 0) | 1<<21
	case ACMPU:
		return OPVCC(10, 0, 0, 0) | 1<<21
	case ACMPW:
		return OPVCC(11, 0, 0, 0)
	case ACMPWU:
		return OPVCC(10, 0, 0, 0)
	case ACMPEQB:
		return OPVCC(31, 224, 0, 0)

	case ALSW:
		return OPVCC(31, 597, 0, 0)

	case ACOPY:
		return OPVCC(31, 774, 0, 0)
	case APASTECC:
		return OPVCC(31, 902, 0, 1)
	case ADARN:
		return OPVCC(31, 755, 0, 0)

	case AMULLW:
		return OPVCC(7, 0, 0, 0)

	case AOR:
		return OPVCC(24, 0, 0, 0)
	case AORIS:
		return OPVCC(25, 0, 0, 0)

	case ARLWMI:
		return OPVCC(20, 0, 0, 0)
	case ARLWMICC:
		return OPVCC(20, 0, 0, 1)
	case ARLDMI:
		return OPVCC(30, 0, 0, 0) | 3<<2
	case ARLDMICC:
		return OPVCC(30, 0, 0, 1) | 3<<2
	case ARLDIMI:
		return OPVCC(30, 0, 0, 0) | 3<<2
	case ARLDIMICC:
		return OPVCC(30, 0, 0, 1) | 3<<2
	case ARLWNM:
		return OPVCC(21, 0, 0, 0)
	case ARLWNMCC:
		return OPVCC(21, 0, 0, 1)

	case ARLDCL:
		return OPVCC(30, 0, 0, 0)
	case ARLDCLCC:
		return OPVCC(30, 0, 0, 1)
	case ARLDCR:
		return OPVCC(30, 1, 0, 0)
	case ARLDCRCC:
		return OPVCC(30, 1, 0, 1)
	case ARLDC:
		return OPVCC(30, 0, 0, 0) | 2<<2
	case ARLDCCC:
		return OPVCC(30, 0, 0, 1) | 2<<2

	case ASRAW:
		return OPVCC(31, 824, 0, 0)
	case ASRAWCC:
		return OPVCC(31, 824, 0, 1)
	case ASRAD:
		return OPVCC(31, (413 << 1), 0, 0)
	case ASRADCC:
		return OPVCC(31, (413 << 1), 0, 1)

	case ASTSW:
		return OPVCC(31, 725, 0, 0)

	case ASUBC:
		return OPVCC(8, 0, 0, 0)

	case ATW:
		return OPVCC(3, 0, 0, 0)
	case ATD:
		return OPVCC(2, 0, 0, 0)

	case AVSPLTB:
		return OPVX(4, 524, 0, 0)
	case AVSPLTH:
		return OPVX(4, 588, 0, 0)
	case AVSPLTW:
		return OPVX(4, 652, 0, 0)

	case AVSPLTISB:
		return OPVX(4, 780, 0, 0)
	case AVSPLTISH:
		return OPVX(4, 844, 0, 0)
	case AVSPLTISW:
		return OPVX(4, 908, 0, 0)

	case AFTDIV:
		return OPVCC(63, 128, 0, 0)
	case AFTSQRT:
		return OPVCC(63, 160, 0, 0)

	case AXOR:
		return OPVCC(26, 0, 0, 0)
	case AXORIS:
		return OPVCC(27, 0, 0, 0)
	}

	c.ctxt.Diag("bad opcode i/r or i/r/r %v", a)
	return 0
}

/*
 * load o(a),d
 */
func (c *ctxt9) opload(a obj.As) uint32 {
	switch a {
	case AMOVD:
		return OPVCC(58, 0, 0, 0)
	case AMOVDU:
		return OPVCC(58, 0, 0, 1)
	case AMOVWZ:
		return OPVCC(32, 0, 0, 0)
	case AMOVWZU:
		return OPVCC(33, 0, 0, 0)
	case AMOVW:
		return OPVCC(58, 0, 0, 0) | 1<<1

	case AMOVB, AMOVBZ:
		return OPVCC(34, 0, 0, 0)

	case AMOVBU, AMOVBZU:
		return OPVCC(35, 0, 0, 0)
	case AFMOVD:
		return OPVCC(50, 0, 0, 0)
	case AFMOVDU:
		return OPVCC(51, 0, 0, 0)
	case AFMOVS:
		return OPVCC(48, 0, 0, 0)
	case AFMOVSU:
		return OPVCC(49, 0, 0, 0)
	case AMOVH:
		return OPVCC(42, 0, 0, 0)
	case AMOVHU:
		return OPVCC(43, 0, 0, 0)
	case AMOVHZ:
		return OPVCC(40, 0, 0, 0)
	case AMOVHZU:
		return OPVCC(41, 0, 0, 0)
	case AMOVMW:
		return OPVCC(46, 0, 0, 0)
	}

	c.ctxt.Diag("bad load opcode %v", a)
	return 0
}

/*
 * indexed load a(b),d
 */
func (c *ctxt9) oploadx(a obj.As) uint32 {
	switch a {
	case AMOVWZ:
		return OPVCC(31, 23, 0, 0)
	case AMOVWZU:
		return OPVCC(31, 55, 0, 0)
	case AMOVW:
		return OPVCC(31, 341, 0, 0)
	case AMOVWU:
		return OPVCC(31, 373, 0, 0)

	case AMOVB, AMOVBZ:
		return OPVCC(31, 87, 0, 0)

	case AMOVBU, AMOVBZU:
		return OPVCC(31, 119, 0, 0)
	case AFMOVD:
		return OPVCC(31, 599, 0, 0)
	case AFMOVDU:
		return OPVCC(31, 631, 0, 0)
	case AFMOVS:
		return OPVCC(31, 535, 0, 0)
	case AFMOVSU:
		return OPVCC(31, 567, 0, 0)
	case AFMOVSX:
		return OPVCC(31, 855, 0, 0)
	case AFMOVSZ:
		return OPVCC(31, 887, 0, 0)
	case AMOVH:
		return OPVCC(31, 343, 0, 0)
	case AMOVHU:
		return OPVCC(31, 375, 0, 0)
	case AMOVHBR:
		return OPVCC(31, 790, 0, 0)
	case AMOVWBR:
		return OPVCC(31, 534, 0, 0)
	case AMOVDBR:
		return OPVCC(31, 532, 0, 0)
	case AMOVHZ:
		return OPVCC(31, 279, 0, 0)
	case AMOVHZU:
		return OPVCC(31, 311, 0, 0)
	case AECIWX:
		return OPVCC(31, 310, 0, 0)
	case ALBAR:
		return OPVCC(31, 52, 0, 0)
	case ALHAR:
		return OPVCC(31, 116, 0, 0)
	case ALWAR:
		return OPVCC(31, 20, 0, 0)
	case ALDAR:
		return OPVCC(31, 84, 0, 0)
	case ALSW:
		return OPVCC(31, 533, 0, 0)
	case AMOVD:
		return OPVCC(31, 21, 0, 0)
	case AMOVDU:
		return OPVCC(31, 53, 0, 0)
	case ALDMX:
		return OPVCC(31, 309, 0, 0)

	case ALVEBX:
		return OPVCC(31, 7, 0, 0)
	case ALVEHX:
		return OPVCC(31, 39, 0, 0)
	case ALVEWX:
		return OPVCC(31, 71, 0, 0)
	case ALVX:
		return OPVCC(31, 103, 0, 0)
	case ALVXL:
		return OPVCC(31, 359, 0, 0)
	case ALVSL:
		return OPVCC(31, 6, 0, 0)
	case ALVSR:
		return OPVCC(31, 38, 0, 0)

	case ALXVD2X:
		return OPVXX1(31, 844, 0)
	case ALXVDSX:
		return OPVXX1(31, 332, 0)
	case ALXVW4X:
		return OPVXX1(31, 780, 0)

	case ALXSDX:
		return OPVXX1(31, 588, 0)

	case ALXSIWAX:
		return OPVXX1(31, 76, 0)
	case ALXSIWZX:
		return OPVXX1(31, 12, 0)

	}

	c.ctxt.Diag("bad loadx opcode %v", a)
	return 0
}

/*
 * store s,o(d)
 */
func (c *ctxt9) opstore(a obj.As) uint32 {
	switch a {
	case AMOVB, AMOVBZ:
		return OPVCC(38, 0, 0, 0)

	case AMOVBU, AMOVBZU:
		return OPVCC(39, 0, 0, 0)
	case AFMOVD:
		return OPVCC(54, 0, 0, 0)
	case AFMOVDU:
		return OPVCC(55, 0, 0, 0)
	case AFMOVS:
		return OPVCC(52, 0, 0, 0)
	case AFMOVSU:
		return OPVCC(53, 0, 0, 0)

	case AMOVHZ, AMOVH:
		return OPVCC(44, 0, 0, 0)

	case AMOVHZU, AMOVHU:
		return OPVCC(45, 0, 0, 0)
	case AMOVMW:
		return OPVCC(47, 0, 0, 0)
	case ASTSW:
		return OPVCC(31, 725, 0, 0)

	case AMOVWZ, AMOVW:
		return OPVCC(36, 0, 0, 0)

	case AMOVWZU, AMOVWU:
		return OPVCC(37, 0, 0, 0)
	case AMOVD:
		return OPVCC(62, 0, 0, 0)
	case AMOVDU:
		return OPVCC(62, 0, 0, 1)
	}

	c.ctxt.Diag("unknown store opcode %v", a)
	return 0
}

/*
 * indexed store s,a(b)
 */
func (c *ctxt9) opstorex(a obj.As) uint32 {
	switch a {
	case AMOVB, AMOVBZ:
		return OPVCC(31, 215, 0, 0)

	case AMOVBU, AMOVBZU:
		return OPVCC(31, 247, 0, 0)
	case AFMOVD:
		return OPVCC(31, 727, 0, 0)
	case AFMOVDU:
		return OPVCC(31, 759, 0, 0)
	case AFMOVS:
		return OPVCC(31, 663, 0, 0)
	case AFMOVSU:
		return OPVCC(31, 695, 0, 0)
	case AFMOVSX:
		return OPVCC(31, 983, 0, 0)

	case AMOVHZ, AMOVH:
		return OPVCC(31, 407, 0, 0)
	case AMOVHBR:
		return OPVCC(31, 918, 0, 0)

	case AMOVHZU, AMOVHU:
		return OPVCC(31, 439, 0, 0)

	case AMOVWZ, AMOVW:
		return OPVCC(31, 151, 0, 0)

	case AMOVWZU, AMOVWU:
		return OPVCC(31, 183, 0, 0)
	case ASTSW:
		return OPVCC(31, 661, 0, 0)
	case AMOVWBR:
		return OPVCC(31, 662, 0, 0)
	case AMOVDBR:
		return OPVCC(31, 660, 0, 0)
	case ASTBCCC:
		return OPVCC(31, 694, 0, 1)
	case ASTWCCC:
		return OPVCC(31, 150, 0, 1)
	case ASTDCCC:
		return OPVCC(31, 214, 0, 1)
	case AECOWX:
		return OPVCC(31, 438, 0, 0)
	case AMOVD:
		return OPVCC(31, 149, 0, 0)
	case AMOVDU:
		return OPVCC(31, 181, 0, 0)

	case ASTVEBX:
		return OPVCC(31, 135, 0, 0)
	case ASTVEHX:
		return OPVCC(31, 167, 0, 0)
	case ASTVEWX:
		return OPVCC(31, 199, 0, 0)
	case ASTVX:
		return OPVCC(31, 231, 0, 0)
	case ASTVXL:
		return OPVCC(31, 487, 0, 0)

	case ASTXVD2X:
		return OPVXX1(31, 972, 0)
	case ASTXVW4X:
		return OPVXX1(31, 908, 0)

	case ASTXSDX:
		return OPVXX1(31, 716, 0)

	case ASTXSIWX:
		return OPVXX1(31, 140, 0)

	}

	c.ctxt.Diag("unknown storex opcode %v", a)
	return 0
}

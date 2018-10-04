// cmd/9l/optab.c, cmd/9l/asmout.c from Vita Nuova.
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2008 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2008 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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

// Instruction layout.

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

func (pstate *PackageState) span9(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	p := cursym.Func.Text
	if p == nil || p.Link == nil { // handle external functions and ELF section symbols
		return
	}

	if pstate.oprange[AANDN&obj.AMask] == nil {
		ctxt.Diag("ppc64 ops not initialized, call ppc64.buildop first")
	}

	c := ctxt9{ctxt: ctxt, newprog: newprog, cursym: cursym, autosize: int32(p.To.Offset)}

	pc := int64(0)
	p.Pc = pc

	var m int
	var o *Optab
	for p = p.Link; p != nil; p = p.Link {
		p.Pc = pc
		o = c.oplook(pstate, p)
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

	/*
	 * if any procedure is large enough to
	 * generate a large SBRA branch, then
	 * generate extra passes putting branches
	 * around jmps to fix. this is rare.
	 */
	bflag := 1

	var otxt int64
	var q *obj.Prog
	for bflag != 0 {
		bflag = 0
		pc = 0
		for p = c.cursym.Func.Text.Link; p != nil; p = p.Link {
			p.Pc = pc
			o = c.oplook(pstate, p)

			// very large conditional branches
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

					//addnop(p->link);
					//addnop(p);
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

	/*
	 * lay out the code, emitting code and data relocations.
	 */

	c.cursym.Grow(c.cursym.Size)

	bp := c.cursym.P
	var i int32
	var out [6]uint32
	for p := c.cursym.Func.Text.Link; p != nil; p = p.Link {
		c.pc = p.Pc
		o = c.oplook(pstate, p)
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
			if a.Sym != nil { // use relocation
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

			/* not sure why this barfs */
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
			if c.instoffset&0xffff == 0 && isuint32(uint64(c.instoffset)) { /* && (instoffset & (1<<31)) == 0) */
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

func (c *ctxt9) oplook(pstate *PackageState, p *obj.Prog) *Optab {
	a1 := int(p.Optab)
	if a1 != 0 {
		return &pstate.optab[a1-1]
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

	//print("oplook %v %d %d %d %d\n", p, a1, a2, a3, a4);
	ops := pstate.oprange[p.As&obj.AMask]
	c1 := &pstate.xcmp[a1]
	c3 := &pstate.xcmp[a3]
	c4 := &pstate.xcmp[a4]
	for i := range ops {
		op := &ops[i]
		if int(op.a2) == a2 && c1[op.a1] && c3[op.a3] && c4[op.a4] {
			p.Optab = uint16(cap(pstate.optab) - cap(ops) + i + 1)
			return op
		}
	}

	c.ctxt.Diag("illegal combination %v %v %v %v %v", p.As, pstate.DRconv(a1), pstate.DRconv(a2), pstate.DRconv(a3), pstate.DRconv(a4))
	prasm(p)
	if ops == nil {
		ops = pstate.optab
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
			return r0iszero != 0 /*TypeKind(100016)*/
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
	// same opcode
	if n != 0 {
		return n < 0
	}
	// Consider those that generate fewer
	// instructions first.
	n = int(p1.size) - int(p2.size)
	if n != 0 {
		return n < 0
	}
	// operand order should match
	// better choices first
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
func (pstate *PackageState) opset(a, b0 obj.As) {
	pstate.oprange[a&obj.AMask] = pstate.oprange[b0]
}

// Build the opcode table
func (pstate *PackageState) buildop(ctxt *obj.Link) {
	if pstate.oprange[AANDN&obj.AMask] != nil {
		// Already initialized; stop now.
		// This happens in the cmd/asm tests,
		// each of which re-initializes the arch.
		return
	}

	var n int

	for i := 0; i < C_NCLASS; i++ {
		for n = 0; n < C_NCLASS; n++ {
			if cmp(n, i) {
				pstate.xcmp[i][n] = true
			}
		}
	}
	for n = 0; pstate.optab[n].as != obj.AXXX; n++ {
	}
	sort.Sort(ocmp(pstate.optab[:n]))
	for i := 0; i < n; i++ {
		r := pstate.optab[i].as
		r0 := r & obj.AMask
		start := i
		for pstate.optab[i].as == r {
			i++
		}
		pstate.oprange[r0] = pstate.optab[start:i]
		i--

		switch r {
		default:
			ctxt.Diag("unknown op in build: %v", r)
			log.Fatalf("instruction missing from switch in asm9.go:buildop: %v", r)

		case ADCBF: /* unary indexed: op (b+a); op (b) */
			pstate.opset(ADCBI, r0)

			pstate.opset(ADCBST, r0)
			pstate.opset(ADCBT, r0)
			pstate.opset(ADCBTST, r0)
			pstate.opset(ADCBZ, r0)
			pstate.opset(AICBI, r0)

		case AECOWX: /* indexed store: op s,(b+a); op s,(b) */
			pstate.opset(ASTWCCC, r0)
			pstate.opset(ASTBCCC, r0)

			pstate.opset(ASTDCCC, r0)

		case AREM: /* macro */
			pstate.opset(AREMCC, r0)

			pstate.opset(AREMV, r0)
			pstate.opset(AREMVCC, r0)

		case AREMU:
			pstate.opset(AREMU, r0)
			pstate.opset(AREMUCC, r0)
			pstate.opset(AREMUV, r0)
			pstate.opset(AREMUVCC, r0)

		case AREMD:
			pstate.opset(AREMDCC, r0)
			pstate.opset(AREMDV, r0)
			pstate.opset(AREMDVCC, r0)

		case AREMDU:
			pstate.opset(AREMDU, r0)
			pstate.opset(AREMDUCC, r0)
			pstate.opset(AREMDUV, r0)
			pstate.opset(AREMDUVCC, r0)

		case ADIVW: /* op Rb[,Ra],Rd */
			pstate.opset(AMULHW, r0)

			pstate.opset(AMULHWCC, r0)
			pstate.opset(AMULHWU, r0)
			pstate.opset(AMULHWUCC, r0)
			pstate.opset(AMULLWCC, r0)
			pstate.opset(AMULLWVCC, r0)
			pstate.opset(AMULLWV, r0)
			pstate.opset(ADIVWCC, r0)
			pstate.opset(ADIVWV, r0)
			pstate.opset(ADIVWVCC, r0)
			pstate.opset(ADIVWU, r0)
			pstate.opset(ADIVWUCC, r0)
			pstate.opset(ADIVWUV, r0)
			pstate.opset(ADIVWUVCC, r0)
			pstate.opset(AADDCC, r0)
			pstate.opset(AADDCV, r0)
			pstate.opset(AADDCVCC, r0)
			pstate.opset(AADDV, r0)
			pstate.opset(AADDVCC, r0)
			pstate.opset(AADDE, r0)
			pstate.opset(AADDECC, r0)
			pstate.opset(AADDEV, r0)
			pstate.opset(AADDEVCC, r0)
			pstate.opset(ACRAND, r0)
			pstate.opset(ACRANDN, r0)
			pstate.opset(ACREQV, r0)
			pstate.opset(ACRNAND, r0)
			pstate.opset(ACRNOR, r0)
			pstate.opset(ACROR, r0)
			pstate.opset(ACRORN, r0)
			pstate.opset(ACRXOR, r0)
			pstate.opset(AMULHD, r0)
			pstate.opset(AMULHDCC, r0)
			pstate.opset(AMULHDU, r0)
			pstate.opset(AMULHDUCC, r0)
			pstate.opset(AMULLD, r0)
			pstate.opset(AMULLDCC, r0)
			pstate.opset(AMULLDVCC, r0)
			pstate.opset(AMULLDV, r0)
			pstate.opset(ADIVD, r0)
			pstate.opset(ADIVDCC, r0)
			pstate.opset(ADIVDE, r0)
			pstate.opset(ADIVDEU, r0)
			pstate.opset(ADIVDECC, r0)
			pstate.opset(ADIVDEUCC, r0)
			pstate.opset(ADIVDVCC, r0)
			pstate.opset(ADIVDV, r0)
			pstate.opset(ADIVDU, r0)
			pstate.opset(ADIVDUCC, r0)
			pstate.opset(ADIVDUVCC, r0)
			pstate.opset(ADIVDUCC, r0)

		case APOPCNTD:
			pstate.opset(APOPCNTW, r0)
			pstate.opset(APOPCNTB, r0)

		case ACOPY: /* copy, paste. */
			pstate.opset(APASTECC, r0)

		case AMADDHD: /* maddhd, maddhdu, maddld */
			pstate.opset(AMADDHDU, r0)
			pstate.opset(AMADDLD, r0)

		case AMOVBZ: /* lbz, stz, rlwm(r/r), lhz, lha, stz, and x variants */
			pstate.opset(AMOVH, r0)
			pstate.opset(AMOVHZ, r0)

		case AMOVBZU: /* lbz[x]u, stb[x]u, lhz[x]u, lha[x]u, sth[u]x, ld[x]u, std[u]x */
			pstate.opset(AMOVHU, r0)

			pstate.opset(AMOVHZU, r0)
			pstate.opset(AMOVWU, r0)
			pstate.opset(AMOVWZU, r0)
			pstate.opset(AMOVDU, r0)
			pstate.opset(AMOVMW, r0)

		case ALV: /* lvebx, lvehx, lvewx, lvx, lvxl, lvsl, lvsr */
			pstate.opset(ALVEBX, r0)
			pstate.opset(ALVEHX, r0)
			pstate.opset(ALVEWX, r0)
			pstate.opset(ALVX, r0)
			pstate.opset(ALVXL, r0)
			pstate.opset(ALVSL, r0)
			pstate.opset(ALVSR, r0)

		case ASTV: /* stvebx, stvehx, stvewx, stvx, stvxl */
			pstate.opset(ASTVEBX, r0)
			pstate.opset(ASTVEHX, r0)
			pstate.opset(ASTVEWX, r0)
			pstate.opset(ASTVX, r0)
			pstate.opset(ASTVXL, r0)

		case AVAND: /* vand, vandc, vnand */
			pstate.opset(AVAND, r0)
			pstate.opset(AVANDC, r0)
			pstate.opset(AVNAND, r0)

		case AVOR: /* vor, vorc, vxor, vnor, veqv */
			pstate.opset(AVOR, r0)
			pstate.opset(AVORC, r0)
			pstate.opset(AVXOR, r0)
			pstate.opset(AVNOR, r0)
			pstate.opset(AVEQV, r0)

		case AVADDUM: /* vaddubm, vadduhm, vadduwm, vaddudm, vadduqm */
			pstate.opset(AVADDUBM, r0)
			pstate.opset(AVADDUHM, r0)
			pstate.opset(AVADDUWM, r0)
			pstate.opset(AVADDUDM, r0)
			pstate.opset(AVADDUQM, r0)

		case AVADDCU: /* vaddcuq, vaddcuw */
			pstate.opset(AVADDCUQ, r0)
			pstate.opset(AVADDCUW, r0)

		case AVADDUS: /* vaddubs, vadduhs, vadduws */
			pstate.opset(AVADDUBS, r0)
			pstate.opset(AVADDUHS, r0)
			pstate.opset(AVADDUWS, r0)

		case AVADDSS: /* vaddsbs, vaddshs, vaddsws */
			pstate.opset(AVADDSBS, r0)
			pstate.opset(AVADDSHS, r0)
			pstate.opset(AVADDSWS, r0)

		case AVADDE: /* vaddeuqm, vaddecuq */
			pstate.opset(AVADDEUQM, r0)
			pstate.opset(AVADDECUQ, r0)

		case AVSUBUM: /* vsububm, vsubuhm, vsubuwm, vsubudm, vsubuqm */
			pstate.opset(AVSUBUBM, r0)
			pstate.opset(AVSUBUHM, r0)
			pstate.opset(AVSUBUWM, r0)
			pstate.opset(AVSUBUDM, r0)
			pstate.opset(AVSUBUQM, r0)

		case AVSUBCU: /* vsubcuq, vsubcuw */
			pstate.opset(AVSUBCUQ, r0)
			pstate.opset(AVSUBCUW, r0)

		case AVSUBUS: /* vsububs, vsubuhs, vsubuws */
			pstate.opset(AVSUBUBS, r0)
			pstate.opset(AVSUBUHS, r0)
			pstate.opset(AVSUBUWS, r0)

		case AVSUBSS: /* vsubsbs, vsubshs, vsubsws */
			pstate.opset(AVSUBSBS, r0)
			pstate.opset(AVSUBSHS, r0)
			pstate.opset(AVSUBSWS, r0)

		case AVSUBE: /* vsubeuqm, vsubecuq */
			pstate.opset(AVSUBEUQM, r0)
			pstate.opset(AVSUBECUQ, r0)

		case AVMULESB: /* vmulesb, vmulosb, vmuleub, vmuloub, vmulosh, vmulouh, vmulesw, vmulosw, vmuleuw, vmulouw, vmuluwm */
			pstate.opset(AVMULOSB, r0)
			pstate.opset(AVMULEUB, r0)
			pstate.opset(AVMULOUB, r0)
			pstate.opset(AVMULESH, r0)
			pstate.opset(AVMULOSH, r0)
			pstate.opset(AVMULEUH, r0)
			pstate.opset(AVMULOUH, r0)
			pstate.opset(AVMULESW, r0)
			pstate.opset(AVMULOSW, r0)
			pstate.opset(AVMULEUW, r0)
			pstate.opset(AVMULOUW, r0)
			pstate.opset(AVMULUWM, r0)
		case AVPMSUM: /* vpmsumb, vpmsumh, vpmsumw, vpmsumd */
			pstate.opset(AVPMSUMB, r0)
			pstate.opset(AVPMSUMH, r0)
			pstate.opset(AVPMSUMW, r0)
			pstate.opset(AVPMSUMD, r0)

		case AVR: /* vrlb, vrlh, vrlw, vrld */
			pstate.opset(AVRLB, r0)
			pstate.opset(AVRLH, r0)
			pstate.opset(AVRLW, r0)
			pstate.opset(AVRLD, r0)

		case AVS: /* vs[l,r], vs[l,r]o, vs[l,r]b, vs[l,r]h, vs[l,r]w, vs[l,r]d */
			pstate.opset(AVSLB, r0)
			pstate.opset(AVSLH, r0)
			pstate.opset(AVSLW, r0)
			pstate.opset(AVSL, r0)
			pstate.opset(AVSLO, r0)
			pstate.opset(AVSRB, r0)
			pstate.opset(AVSRH, r0)
			pstate.opset(AVSRW, r0)
			pstate.opset(AVSR, r0)
			pstate.opset(AVSRO, r0)
			pstate.opset(AVSLD, r0)
			pstate.opset(AVSRD, r0)

		case AVSA: /* vsrab, vsrah, vsraw, vsrad */
			pstate.opset(AVSRAB, r0)
			pstate.opset(AVSRAH, r0)
			pstate.opset(AVSRAW, r0)
			pstate.opset(AVSRAD, r0)

		case AVSOI: /* vsldoi */
			pstate.opset(AVSLDOI, r0)

		case AVCLZ: /* vclzb, vclzh, vclzw, vclzd */
			pstate.opset(AVCLZB, r0)
			pstate.opset(AVCLZH, r0)
			pstate.opset(AVCLZW, r0)
			pstate.opset(AVCLZD, r0)

		case AVPOPCNT: /* vpopcntb, vpopcnth, vpopcntw, vpopcntd */
			pstate.opset(AVPOPCNTB, r0)
			pstate.opset(AVPOPCNTH, r0)
			pstate.opset(AVPOPCNTW, r0)
			pstate.opset(AVPOPCNTD, r0)

		case AVCMPEQ: /* vcmpequb[.], vcmpequh[.], vcmpequw[.], vcmpequd[.] */
			pstate.opset(AVCMPEQUB, r0)
			pstate.opset(AVCMPEQUBCC, r0)
			pstate.opset(AVCMPEQUH, r0)
			pstate.opset(AVCMPEQUHCC, r0)
			pstate.opset(AVCMPEQUW, r0)
			pstate.opset(AVCMPEQUWCC, r0)
			pstate.opset(AVCMPEQUD, r0)
			pstate.opset(AVCMPEQUDCC, r0)

		case AVCMPGT: /* vcmpgt[u,s]b[.], vcmpgt[u,s]h[.], vcmpgt[u,s]w[.], vcmpgt[u,s]d[.] */
			pstate.opset(AVCMPGTUB, r0)
			pstate.opset(AVCMPGTUBCC, r0)
			pstate.opset(AVCMPGTUH, r0)
			pstate.opset(AVCMPGTUHCC, r0)
			pstate.opset(AVCMPGTUW, r0)
			pstate.opset(AVCMPGTUWCC, r0)
			pstate.opset(AVCMPGTUD, r0)
			pstate.opset(AVCMPGTUDCC, r0)
			pstate.opset(AVCMPGTSB, r0)
			pstate.opset(AVCMPGTSBCC, r0)
			pstate.opset(AVCMPGTSH, r0)
			pstate.opset(AVCMPGTSHCC, r0)
			pstate.opset(AVCMPGTSW, r0)
			pstate.opset(AVCMPGTSWCC, r0)
			pstate.opset(AVCMPGTSD, r0)
			pstate.opset(AVCMPGTSDCC, r0)

		case AVCMPNEZB: /* vcmpnezb[.] */
			pstate.opset(AVCMPNEZBCC, r0)

		case AVPERM: /* vperm */
			pstate.opset(AVPERM, r0)

		case AVBPERMQ: /* vbpermq, vbpermd */
			pstate.opset(AVBPERMD, r0)

		case AVSEL: /* vsel */
			pstate.opset(AVSEL, r0)

		case AVSPLT: /* vspltb, vsplth, vspltw */
			pstate.opset(AVSPLTB, r0)
			pstate.opset(AVSPLTH, r0)
			pstate.opset(AVSPLTW, r0)

		case AVSPLTI: /* vspltisb, vspltish, vspltisw */
			pstate.opset(AVSPLTISB, r0)
			pstate.opset(AVSPLTISH, r0)
			pstate.opset(AVSPLTISW, r0)

		case AVCIPH: /* vcipher, vcipherlast */
			pstate.opset(AVCIPHER, r0)
			pstate.opset(AVCIPHERLAST, r0)

		case AVNCIPH: /* vncipher, vncipherlast */
			pstate.opset(AVNCIPHER, r0)
			pstate.opset(AVNCIPHERLAST, r0)

		case AVSBOX: /* vsbox */
			pstate.opset(AVSBOX, r0)

		case AVSHASIGMA: /* vshasigmaw, vshasigmad */
			pstate.opset(AVSHASIGMAW, r0)
			pstate.opset(AVSHASIGMAD, r0)

		case ALXV: /* lxvd2x, lxvdsx, lxvw4x */
			pstate.opset(ALXVD2X, r0)
			pstate.opset(ALXVDSX, r0)
			pstate.opset(ALXVW4X, r0)

		case ASTXV: /* stxvd2x, stxvdsx, stxvw4x */
			pstate.opset(ASTXVD2X, r0)
			pstate.opset(ASTXVW4X, r0)

		case ALXS: /* lxsdx  */
			pstate.opset(ALXSDX, r0)

		case ASTXS: /* stxsdx */
			pstate.opset(ASTXSDX, r0)

		case ALXSI: /* lxsiwax, lxsiwzx  */
			pstate.opset(ALXSIWAX, r0)
			pstate.opset(ALXSIWZX, r0)

		case ASTXSI: /* stxsiwx */
			pstate.opset(ASTXSIWX, r0)

		case AMFVSR: /* mfvsrd, mfvsrwz (and extended mnemonics), mfvsrld */
			pstate.opset(AMFVSRD, r0)
			pstate.opset(AMFFPRD, r0)
			pstate.opset(AMFVRD, r0)
			pstate.opset(AMFVSRWZ, r0)
			pstate.opset(AMFVSRLD, r0)

		case AMTVSR: /* mtvsrd, mtvsrwa, mtvsrwz (and extended mnemonics), mtvsrdd, mtvsrws */
			pstate.opset(AMTVSRD, r0)
			pstate.opset(AMTFPRD, r0)
			pstate.opset(AMTVRD, r0)
			pstate.opset(AMTVSRWA, r0)
			pstate.opset(AMTVSRWZ, r0)
			pstate.opset(AMTVSRDD, r0)
			pstate.opset(AMTVSRWS, r0)

		case AXXLAND: /* xxland, xxlandc, xxleqv, xxlnand */
			pstate.opset(AXXLANDQ, r0)
			pstate.opset(AXXLANDC, r0)
			pstate.opset(AXXLEQV, r0)
			pstate.opset(AXXLNAND, r0)

		case AXXLOR: /* xxlorc, xxlnor, xxlor, xxlxor */
			pstate.opset(AXXLORC, r0)
			pstate.opset(AXXLNOR, r0)
			pstate.opset(AXXLORQ, r0)
			pstate.opset(AXXLXOR, r0)

		case AXXSEL: /* xxsel */
			pstate.opset(AXXSEL, r0)

		case AXXMRG: /* xxmrghw, xxmrglw */
			pstate.opset(AXXMRGHW, r0)
			pstate.opset(AXXMRGLW, r0)

		case AXXSPLT: /* xxspltw */
			pstate.opset(AXXSPLTW, r0)

		case AXXPERM: /* xxpermdi */
			pstate.opset(AXXPERMDI, r0)

		case AXXSI: /* xxsldwi */
			pstate.opset(AXXSLDWI, r0)

		case AXSCV: /* xscvdpsp, xscvspdp, xscvdpspn, xscvspdpn */
			pstate.opset(AXSCVDPSP, r0)
			pstate.opset(AXSCVSPDP, r0)
			pstate.opset(AXSCVDPSPN, r0)
			pstate.opset(AXSCVSPDPN, r0)

		case AXVCV: /* xvcvdpsp, xvcvspdp */
			pstate.opset(AXVCVDPSP, r0)
			pstate.opset(AXVCVSPDP, r0)

		case AXSCVX: /* xscvdpsxds, xscvdpsxws, xscvdpuxds, xscvdpuxws */
			pstate.opset(AXSCVDPSXDS, r0)
			pstate.opset(AXSCVDPSXWS, r0)
			pstate.opset(AXSCVDPUXDS, r0)
			pstate.opset(AXSCVDPUXWS, r0)

		case AXSCVXP: /* xscvsxddp, xscvuxddp, xscvsxdsp, xscvuxdsp */
			pstate.opset(AXSCVSXDDP, r0)
			pstate.opset(AXSCVUXDDP, r0)
			pstate.opset(AXSCVSXDSP, r0)
			pstate.opset(AXSCVUXDSP, r0)

		case AXVCVX: /* xvcvdpsxds, xvcvdpsxws, xvcvdpuxds, xvcvdpuxws, xvcvspsxds, xvcvspsxws, xvcvspuxds, xvcvspuxws */
			pstate.opset(AXVCVDPSXDS, r0)
			pstate.opset(AXVCVDPSXWS, r0)
			pstate.opset(AXVCVDPUXDS, r0)
			pstate.opset(AXVCVDPUXWS, r0)
			pstate.opset(AXVCVSPSXDS, r0)
			pstate.opset(AXVCVSPSXWS, r0)
			pstate.opset(AXVCVSPUXDS, r0)
			pstate.opset(AXVCVSPUXWS, r0)

		case AXVCVXP: /* xvcvsxddp, xvcvsxwdp, xvcvuxddp, xvcvuxwdp, xvcvsxdsp, xvcvsxwsp, xvcvuxdsp, xvcvuxwsp */
			pstate.opset(AXVCVSXDDP, r0)
			pstate.opset(AXVCVSXWDP, r0)
			pstate.opset(AXVCVUXDDP, r0)
			pstate.opset(AXVCVUXWDP, r0)
			pstate.opset(AXVCVSXDSP, r0)
			pstate.opset(AXVCVSXWSP, r0)
			pstate.opset(AXVCVUXDSP, r0)
			pstate.opset(AXVCVUXWSP, r0)

		case AAND: /* logical op Rb,Rs,Ra; no literal */
			pstate.opset(AANDN, r0)
			pstate.opset(AANDNCC, r0)
			pstate.opset(AEQV, r0)
			pstate.opset(AEQVCC, r0)
			pstate.opset(ANAND, r0)
			pstate.opset(ANANDCC, r0)
			pstate.opset(ANOR, r0)
			pstate.opset(ANORCC, r0)
			pstate.opset(AORCC, r0)
			pstate.opset(AORN, r0)
			pstate.opset(AORNCC, r0)
			pstate.opset(AXORCC, r0)

		case AADDME: /* op Ra, Rd */
			pstate.opset(AADDMECC, r0)

			pstate.opset(AADDMEV, r0)
			pstate.opset(AADDMEVCC, r0)
			pstate.opset(AADDZE, r0)
			pstate.opset(AADDZECC, r0)
			pstate.opset(AADDZEV, r0)
			pstate.opset(AADDZEVCC, r0)
			pstate.opset(ASUBME, r0)
			pstate.opset(ASUBMECC, r0)
			pstate.opset(ASUBMEV, r0)
			pstate.opset(ASUBMEVCC, r0)
			pstate.opset(ASUBZE, r0)
			pstate.opset(ASUBZECC, r0)
			pstate.opset(ASUBZEV, r0)
			pstate.opset(ASUBZEVCC, r0)

		case AADDC:
			pstate.opset(AADDCCC, r0)

		case ABEQ:
			pstate.opset(ABGE, r0)
			pstate.opset(ABGT, r0)
			pstate.opset(ABLE, r0)
			pstate.opset(ABLT, r0)
			pstate.opset(ABNE, r0)
			pstate.opset(ABVC, r0)
			pstate.opset(ABVS, r0)

		case ABR:
			pstate.opset(ABL, r0)

		case ABC:
			pstate.opset(ABCL, r0)

		case AEXTSB: /* op Rs, Ra */
			pstate.opset(AEXTSBCC, r0)

			pstate.opset(AEXTSH, r0)
			pstate.opset(AEXTSHCC, r0)
			pstate.opset(ACNTLZW, r0)
			pstate.opset(ACNTLZWCC, r0)
			pstate.opset(ACNTLZD, r0)
			pstate.opset(AEXTSW, r0)
			pstate.opset(AEXTSWCC, r0)
			pstate.opset(ACNTLZDCC, r0)

		case AFABS: /* fop [s,]d */
			pstate.opset(AFABSCC, r0)

			pstate.opset(AFNABS, r0)
			pstate.opset(AFNABSCC, r0)
			pstate.opset(AFNEG, r0)
			pstate.opset(AFNEGCC, r0)
			pstate.opset(AFRSP, r0)
			pstate.opset(AFRSPCC, r0)
			pstate.opset(AFCTIW, r0)
			pstate.opset(AFCTIWCC, r0)
			pstate.opset(AFCTIWZ, r0)
			pstate.opset(AFCTIWZCC, r0)
			pstate.opset(AFCTID, r0)
			pstate.opset(AFCTIDCC, r0)
			pstate.opset(AFCTIDZ, r0)
			pstate.opset(AFCTIDZCC, r0)
			pstate.opset(AFCFID, r0)
			pstate.opset(AFCFIDCC, r0)
			pstate.opset(AFCFIDU, r0)
			pstate.opset(AFCFIDUCC, r0)
			pstate.opset(AFCFIDS, r0)
			pstate.opset(AFCFIDSCC, r0)
			pstate.opset(AFRES, r0)
			pstate.opset(AFRESCC, r0)
			pstate.opset(AFRIM, r0)
			pstate.opset(AFRIMCC, r0)
			pstate.opset(AFRIP, r0)
			pstate.opset(AFRIPCC, r0)
			pstate.opset(AFRIZ, r0)
			pstate.opset(AFRIZCC, r0)
			pstate.opset(AFRIN, r0)
			pstate.opset(AFRINCC, r0)
			pstate.opset(AFRSQRTE, r0)
			pstate.opset(AFRSQRTECC, r0)
			pstate.opset(AFSQRT, r0)
			pstate.opset(AFSQRTCC, r0)
			pstate.opset(AFSQRTS, r0)
			pstate.opset(AFSQRTSCC, r0)

		case AFADD:
			pstate.opset(AFADDS, r0)
			pstate.opset(AFADDCC, r0)
			pstate.opset(AFADDSCC, r0)
			pstate.opset(AFCPSGN, r0)
			pstate.opset(AFCPSGNCC, r0)
			pstate.opset(AFDIV, r0)
			pstate.opset(AFDIVS, r0)
			pstate.opset(AFDIVCC, r0)
			pstate.opset(AFDIVSCC, r0)
			pstate.opset(AFSUB, r0)
			pstate.opset(AFSUBS, r0)
			pstate.opset(AFSUBCC, r0)
			pstate.opset(AFSUBSCC, r0)

		case AFMADD:
			pstate.opset(AFMADDCC, r0)
			pstate.opset(AFMADDS, r0)
			pstate.opset(AFMADDSCC, r0)
			pstate.opset(AFMSUB, r0)
			pstate.opset(AFMSUBCC, r0)
			pstate.opset(AFMSUBS, r0)
			pstate.opset(AFMSUBSCC, r0)
			pstate.opset(AFNMADD, r0)
			pstate.opset(AFNMADDCC, r0)
			pstate.opset(AFNMADDS, r0)
			pstate.opset(AFNMADDSCC, r0)
			pstate.opset(AFNMSUB, r0)
			pstate.opset(AFNMSUBCC, r0)
			pstate.opset(AFNMSUBS, r0)
			pstate.opset(AFNMSUBSCC, r0)
			pstate.opset(AFSEL, r0)
			pstate.opset(AFSELCC, r0)

		case AFMUL:
			pstate.opset(AFMULS, r0)
			pstate.opset(AFMULCC, r0)
			pstate.opset(AFMULSCC, r0)

		case AFCMPO:
			pstate.opset(AFCMPU, r0)

		case AISEL:
			pstate.opset(AISEL, r0)

		case AMTFSB0:
			pstate.opset(AMTFSB0CC, r0)
			pstate.opset(AMTFSB1, r0)
			pstate.opset(AMTFSB1CC, r0)

		case ANEG: /* op [Ra,] Rd */
			pstate.opset(ANEGCC, r0)

			pstate.opset(ANEGV, r0)
			pstate.opset(ANEGVCC, r0)

		case AOR: /* or/xor Rb,Rs,Ra; ori/xori $uimm,Rs,R */
			pstate.opset(AXOR, r0)

		case AORIS: /* oris/xoris $uimm,Rs,Ra */
			pstate.opset(AXORIS, r0)

		case ASLW:
			pstate.opset(ASLWCC, r0)
			pstate.opset(ASRW, r0)
			pstate.opset(ASRWCC, r0)
			pstate.opset(AROTLW, r0)

		case ASLD:
			pstate.opset(ASLDCC, r0)
			pstate.opset(ASRD, r0)
			pstate.opset(ASRDCC, r0)
			pstate.opset(AROTL, r0)

		case ASRAW: /* sraw Rb,Rs,Ra; srawi sh,Rs,Ra */
			pstate.opset(ASRAWCC, r0)

		case ASRAD: /* sraw Rb,Rs,Ra; srawi sh,Rs,Ra */
			pstate.opset(ASRADCC, r0)

		case ASUB: /* SUB Ra,Rb,Rd => subf Rd,ra,rb */
			pstate.opset(ASUB, r0)

			pstate.opset(ASUBCC, r0)
			pstate.opset(ASUBV, r0)
			pstate.opset(ASUBVCC, r0)
			pstate.opset(ASUBCCC, r0)
			pstate.opset(ASUBCV, r0)
			pstate.opset(ASUBCVCC, r0)
			pstate.opset(ASUBE, r0)
			pstate.opset(ASUBECC, r0)
			pstate.opset(ASUBEV, r0)
			pstate.opset(ASUBEVCC, r0)

		case ASYNC:
			pstate.opset(AISYNC, r0)
			pstate.opset(ALWSYNC, r0)
			pstate.opset(APTESYNC, r0)
			pstate.opset(ATLBSYNC, r0)

		case ARLWMI:
			pstate.opset(ARLWMICC, r0)
			pstate.opset(ARLWNM, r0)
			pstate.opset(ARLWNMCC, r0)

		case ARLDMI:
			pstate.opset(ARLDMICC, r0)
			pstate.opset(ARLDIMI, r0)
			pstate.opset(ARLDIMICC, r0)

		case ARLDC:
			pstate.opset(ARLDCCC, r0)

		case ARLDCL:
			pstate.opset(ARLDCR, r0)
			pstate.opset(ARLDCLCC, r0)
			pstate.opset(ARLDCRCC, r0)

		case ARLDICL:
			pstate.opset(ARLDICLCC, r0)
			pstate.opset(ARLDICR, r0)
			pstate.opset(ARLDICRCC, r0)

		case AFMOVD:
			pstate.opset(AFMOVDCC, r0)
			pstate.opset(AFMOVDU, r0)
			pstate.opset(AFMOVS, r0)
			pstate.opset(AFMOVSU, r0)

		case ALDAR:
			pstate.opset(ALBAR, r0)
			pstate.opset(ALHAR, r0)
			pstate.opset(ALWAR, r0)

		case ASYSCALL: /* just the op; flow of control */
			pstate.opset(ARFI, r0)

			pstate.opset(ARFCI, r0)
			pstate.opset(ARFID, r0)
			pstate.opset(AHRFID, r0)

		case AMOVHBR:
			pstate.opset(AMOVWBR, r0)
			pstate.opset(AMOVDBR, r0)

		case ASLBMFEE:
			pstate.opset(ASLBMFEV, r0)

		case ATW:
			pstate.opset(ATD, r0)

		case ATLBIE:
			pstate.opset(ASLBIE, r0)
			pstate.opset(ATLBIEL, r0)

		case AEIEIO:
			pstate.opset(ASLBIA, r0)

		case ACMP:
			pstate.opset(ACMPW, r0)

		case ACMPU:
			pstate.opset(ACMPWU, r0)

		case ACMPB:
			pstate.opset(ACMPB, r0)

		case AFTDIV:
			pstate.opset(AFTDIV, r0)

		case AFTSQRT:
			pstate.opset(AFTSQRT, r0)

		case AADD,
			AADDIS,
			AANDCC, /* and. Rb,Rs,Ra; andi. $uimm,Rs,Ra */
			AANDISCC,
			AFMOVSX,
			AFMOVSZ,
			ALSW,
			AMOVW,
			/* load/store/move word with sign extension; special 32-bit move; move 32-bit literals */
			AMOVWZ, /* load/store/move word with zero extension; move 32-bit literals  */
			AMOVD,  /* load/store/move 64-bit values, including 32-bit literals with/without sign-extension */
			AMOVB,  /* macro: move byte with sign extension */
			AMOVBU, /* macro: move byte with sign extension & update */
			AMOVFL,
			AMULLW,
			/* op $s[,r2],r3; op r1[,r2],r3; no cc/v */
			ASUBC, /* op r1,$s,r3; op r1[,r2],r3 */
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
	/* For the XX-form encodings, we need the VSX register number to be exactly */
	/* between 0-63, so we can properly set the rightmost bits. */
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
	case OPVCC(58, 0, 0, 0), // ld
		OPVCC(58, 0, 0, 1),        // ldu
		OPVCC(58, 0, 0, 0) | 1<<1, // lwa
		OPVCC(62, 0, 0, 0),        // std
		OPVCC(62, 0, 0, 1):        //stdu
		return DS_FORM
	case OP_ADDI, // add
		OPVCC(32, 0, 0, 0), // lwz
		OPVCC(33, 0, 0, 0), // lwzu
		OPVCC(34, 0, 0, 0), // lbz
		OPVCC(35, 0, 0, 0), // lbzu
		OPVCC(40, 0, 0, 0), // lhz
		OPVCC(41, 0, 0, 0), // lhzu
		OPVCC(42, 0, 0, 0), // lha
		OPVCC(43, 0, 0, 0), // lhau
		OPVCC(46, 0, 0, 0), // lmw
		OPVCC(48, 0, 0, 0), // lfs
		OPVCC(49, 0, 0, 0), // lfsu
		OPVCC(50, 0, 0, 0), // lfd
		OPVCC(51, 0, 0, 0), // lfdu
		OPVCC(36, 0, 0, 0), // stw
		OPVCC(37, 0, 0, 0), // stwu
		OPVCC(38, 0, 0, 0), // stb
		OPVCC(39, 0, 0, 0), // stbu
		OPVCC(44, 0, 0, 0), // sth
		OPVCC(45, 0, 0, 0), // sthu
		OPVCC(47, 0, 0, 0), // stmw
		OPVCC(52, 0, 0, 0), // stfs
		OPVCC(53, 0, 0, 0), // stfsu
		OPVCC(54, 0, 0, 0), // stfd
		OPVCC(55, 0, 0, 0): // stfdu
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
	if v != ^uint32(0) && v&(1<<31) != 0 && v&1 != 0 { /* MB > ME */
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

	//print("%v => case %d\n", p, o->type);
	switch o.type_ {
	default:
		c.ctxt.Diag("unknown type %d", o.type_)
		prasm(p)

	case 0: /* pseudo ops */
		break

	case 1: /* mov r1,r2 ==> OR Rs,Rs,Ra */
		if p.To.Reg == REGZERO && p.From.Type == obj.TYPE_CONST {
			v := c.regoff(&p.From)
			if r0iszero != 0 /*TypeKind(100016)*/ && v != 0 {
				//nerrors--;
				c.ctxt.Diag("literal operation on R0\n%v", p)
			}

			o1 = LOP_IRR(OP_ADDI, REGZERO, REGZERO, uint32(v))
			break
		}

		o1 = LOP_RRR(OP_OR, uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.From.Reg))

	case 2: /* int/cr/fp op Rb,[Ra],Rd */
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), uint32(p.From.Reg))

	case 3: /* mov $soreg/addcon/andcon/ucon, r ==> addis/oris/addi/ori $i,reg',r */
		d := c.vregoff(&p.From)

		v := int32(d)
		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		if r0iszero != 0 /*TypeKind(100016)*/ && p.To.Reg == 0 && (r != 0 || v != 0) {
			c.ctxt.Diag("literal operation on R0\n%v", p)
		}
		a := OP_ADDI
		if o.a1 == C_UCON {
			if d&0xffff != 0 {
				log.Fatalf("invalid handling of %v", p)
			}
			// For UCON operands the value is right shifted 16, using ADDIS if the
			// value should be signed, ORIS if unsigned.
			v >>= 16
			if r == REGZERO && isuint32(uint64(d)) {
				o1 = LOP_IRR(OP_ORIS, uint32(p.To.Reg), REGZERO, uint32(v))
				break
			}

			a = OP_ADDIS
		} else if int64(int16(d)) != d {
			// Operand is 16 bit value with sign bit set
			if o.a1 == C_ANDCON {
				// Needs unsigned 16 bit so use ORI
				if r == 0 || r == REGZERO {
					o1 = LOP_IRR(uint32(OP_ORI), uint32(p.To.Reg), uint32(0), uint32(v))
					break
				}
				// With ADDCON, needs signed 16 bit value, fall through to use ADDI
			} else if o.a1 != C_ADDCON {
				log.Fatalf("invalid handling of %v", p)
			}
		}

		o1 = AOP_IRR(uint32(a), uint32(p.To.Reg), uint32(r), uint32(v))

	case 4: /* add/mul $scon,[r1],r2 */
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		if r0iszero != 0 /*TypeKind(100016)*/ && p.To.Reg == 0 {
			c.ctxt.Diag("literal operation on R0\n%v", p)
		}
		if int32(int16(v)) != v {
			log.Fatalf("mishandled instruction %v", p)
		}
		o1 = AOP_IRR(c.opirr(p.As), uint32(p.To.Reg), uint32(r), uint32(v))

	case 5: /* syscall */
		o1 = c.oprrr(p.As)

	case 6: /* logical op Rb,[Rs,]Ra; no literal */
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		// AROTL and AROTLW are extended mnemonics, which map to RLDCL and RLWNM.
		switch p.As {
		case AROTL:
			o1 = AOP_RLDIC(OP_RLDCL, uint32(p.To.Reg), uint32(r), uint32(p.From.Reg), uint32(0))
		case AROTLW:
			o1 = OP_RLW(OP_RLWNM, uint32(p.To.Reg), uint32(r), uint32(p.From.Reg), 0, 31)
		default:
			o1 = LOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), uint32(p.From.Reg))
		}

	case 7: /* mov r, soreg ==> stw o(r) */
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
				// This (and the matching part in the load case
				// below) are the only places in the ppc64 toolchain
				// that knows the name of the tls variable. Possibly
				// we could add some assembly syntax so that the name
				// of the variable does not have to be assumed.
				rel.Sym = c.ctxt.Lookup("runtime.tls_g")
				rel.Type = objabi.R_POWER_TLS
			}
			o1 = AOP_RRR(c.opstorex(p.As), uint32(p.From.Reg), uint32(p.To.Index), uint32(r))
		} else {
			if int32(int16(v)) != v {
				log.Fatalf("mishandled instruction %v", p)
			}
			// Offsets in DS form stores must be a multiple of 4
			inst := c.opstore(p.As)
			if c.opform(inst) == DS_FORM && v&0x3 != 0 {
				log.Fatalf("invalid offset for DS form load/store %v", p)
			}
			o1 = AOP_IRR(inst, uint32(p.From.Reg), uint32(r), uint32(v))
		}

	case 8: /* mov soreg, r ==> lbz/lhz/lwz o(r) */
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
			// Offsets in DS form loads must be a multiple of 4
			inst := c.opload(p.As)
			if c.opform(inst) == DS_FORM && v&0x3 != 0 {
				log.Fatalf("invalid offset for DS form load/store %v", p)
			}
			o1 = AOP_IRR(inst, uint32(p.To.Reg), uint32(r), uint32(v))
		}

	case 9: /* movb soreg, r ==> lbz o(r),r2; extsb r2,r2 */
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

	case 10: /* sub Ra,[Rb],Rd => subf Rd,Ra,Rb */
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(r))

	case 11: /* br/bl lbra */
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
		o2 = 0x60000000 // nop, sometimes overwritten by ld r2, 24(r1) when dynamic linking

	case 12: /* movb r,r (extsb); movw r,r (extsw) */
		if p.To.Reg == REGZERO && p.From.Type == obj.TYPE_CONST {
			v := c.regoff(&p.From)
			if r0iszero != 0 /*TypeKind(100016)*/ && v != 0 {
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

	case 13: /* mov[bhw]z r,r; uses rlwinm not andi. to avoid changing CC */
		if p.As == AMOVBZ {
			o1 = OP_RLW(OP_RLWINM, uint32(p.To.Reg), uint32(p.From.Reg), 0, 24, 31)
		} else if p.As == AMOVH {
			o1 = LOP_RRR(OP_EXTSH, uint32(p.To.Reg), uint32(p.From.Reg), 0)
		} else if p.As == AMOVHZ {
			o1 = OP_RLW(OP_RLWINM, uint32(p.To.Reg), uint32(p.From.Reg), 0, 16, 31)
		} else if p.As == AMOVWZ {
			o1 = OP_RLW(OP_RLDIC, uint32(p.To.Reg), uint32(p.From.Reg), 0, 0, 0) | 1<<5 /* MB=32 */
		} else {
			c.ctxt.Diag("internal: bad mov[bhw]z\n%v", p)
		}

	case 14: /* rldc[lr] Rb,Rs,$mask,Ra -- left, right give different masks */
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		d := c.vregoff(p.GetFrom3())
		var a int
		switch p.As {

		// These opcodes expect a mask operand that has to be converted into the
		// appropriate operand.  The way these were defined, not all valid masks are possible.
		// Left here for compatibility in case they were used or generated.
		case ARLDCL, ARLDCLCC:
			var mask [2]uint8
			c.maskgen64(p, mask[:], uint64(d))

			a = int(mask[0]) /* MB */
			if mask[1] != 63 {
				c.ctxt.Diag("invalid mask for rotate: %x (end != bit 63)\n%v", uint64(d), p)
			}
			o1 = LOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), uint32(p.From.Reg))
			o1 |= (uint32(a) & 31) << 6
			if a&0x20 != 0 {
				o1 |= 1 << 5 /* mb[5] is top bit */
			}

		case ARLDCR, ARLDCRCC:
			var mask [2]uint8
			c.maskgen64(p, mask[:], uint64(d))

			a = int(mask[1]) /* ME */
			if mask[0] != 0 {
				c.ctxt.Diag("invalid mask for rotate: %x (start != 0)\n%v", uint64(d), p)
			}
			o1 = LOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), uint32(p.From.Reg))
			o1 |= (uint32(a) & 31) << 6
			if a&0x20 != 0 {
				o1 |= 1 << 5 /* mb[5] is top bit */
			}

		// These opcodes use a shift count like the ppc64 asm, no mask conversion done
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

	case 17, /* bc bo,bi,lbra (same for now) */
		16: /* bc bo,bi,sbra */
		a := 0

		r := int(p.Reg)

		if p.From.Type == obj.TYPE_CONST {
			a = int(c.regoff(&p.From))
		} else if p.From.Type == obj.TYPE_REG {
			if r != 0 {
				c.ctxt.Diag("unexpected register setting for branch with CR: %d\n", r)
			}
			// BI values for the CR
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

	case 15: /* br/bl (r) => mov r,lr; br/bl (lr) */
		var v int32
		if p.As == ABC || p.As == ABCL {
			v = c.regoff(&p.To) & 31
		} else {
			v = 20 /* unconditional */
		}
		o1 = AOP_RRR(OP_MTSPR, uint32(p.To.Reg), 0, 0) | (REG_LR&0x1f)<<16 | ((REG_LR>>5)&0x1f)<<11
		o2 = OPVCC(19, 16, 0, 0)
		if p.As == ABL || p.As == ABCL {
			o2 |= 1
		}
		o2 = OP_BCR(o2, uint32(v), uint32(p.To.Index))

	case 18: /* br/bl (lr/ctr); bc/bcl bo,bi,(lr/ctr) */
		var v int32
		if p.As == ABC || p.As == ABCL {
			v = c.regoff(&p.From) & 31
		} else {
			v = 20 /* unconditional */
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

	case 19: /* mov $lcon,r ==> cau+or */
		d := c.vregoff(&p.From)

		if p.From.Sym == nil {
			o1 = loadu32(int(p.To.Reg), d)
			o2 = LOP_IRR(OP_ORI, uint32(p.To.Reg), uint32(p.To.Reg), uint32(int32(d)))
		} else {
			o1, o2 = c.symbolAccess(p.From.Sym, d, p.To.Reg, OP_ADDI)
		}

	case 20: /* add $ucon,,r | addis $addcon,r,r */
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		if p.As == AADD && (r0iszero == 0 /*TypeKind(100016)*/ && p.Reg == 0 || r0iszero != 0 /*TypeKind(100016)*/ && p.To.Reg == 0) {
			c.ctxt.Diag("literal operation on R0\n%v", p)
		}
		if p.As == AADDIS {
			o1 = AOP_IRR(c.opirr(p.As), uint32(p.To.Reg), uint32(r), uint32(v))
		} else {
			o1 = AOP_IRR(c.opirr(AADDIS), uint32(p.To.Reg), uint32(r), uint32(v)>>16)
		}

	case 22: /* add $lcon/$andcon,r1,r2 ==> oris+ori+add/ori+add */
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
		// If operand is ANDCON, generate 2 instructions using
		// ORI for unsigned value; with LCON 3 instructions.
		if o.size == 8 {
			o1 = LOP_IRR(OP_ORI, REGTMP, REGZERO, uint32(int32(d)))
			o2 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), REGTMP, uint32(r))
		} else {
			o1 = loadu32(REGTMP, d)
			o2 = LOP_IRR(OP_ORI, REGTMP, REGTMP, uint32(int32(d)))
			o3 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), REGTMP, uint32(r))
		}

	case 23: /* and $lcon/$addcon,r1,r2 ==> oris+ori+and/addi+and */
		if p.To.Reg == REGTMP || p.Reg == REGTMP {
			c.ctxt.Diag("can't synthesize large constant\n%v", p)
		}
		d := c.vregoff(&p.From)
		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}

		// With ADDCON operand, generate 2 instructions using ADDI for signed value,
		// with LCON operand generate 3 instructions.
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
		/* sld[.] $sh,rS,rA -> rldicr[.] $sh,rS,mask(0,63-sh),rA; srd[.] -> rldicl */
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
			o1 |= 1 // Set the condition code bit
		}

	case 26: /* mov $lsext/auto/oreg,,r2 ==> addis+addi */
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

	case 27: /* subc ra,$simm,rd => subfic rd,ra,$simm */
		v := c.regoff(p.GetFrom3())

		r := int(p.From.Reg)
		o1 = AOP_IRR(c.opirr(p.As), uint32(p.To.Reg), uint32(r), uint32(v))

	case 28: /* subc r1,$lcon,r2 ==> cau+or+subfc */
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

	//if(dlm) reloc(&p->from3, p->pc, 0);

	case 29: /* rldic[lr]? $sh,s,$mask,a -- left, right, plain give different masks */
		v := c.regoff(&p.From)

		d := c.vregoff(p.GetFrom3())
		var mask [2]uint8
		c.maskgen64(p, mask[:], uint64(d))
		var a int
		switch p.As {
		case ARLDC, ARLDCCC:
			a = int(mask[0]) /* MB */
			if int32(mask[1]) != (63 - v) {
				c.ctxt.Diag("invalid mask for shift: %x (shift %d)\n%v", uint64(d), v, p)
			}

		case ARLDCL, ARLDCLCC:
			a = int(mask[0]) /* MB */
			if mask[1] != 63 {
				c.ctxt.Diag("invalid mask for shift: %x (shift %d)\n%v", uint64(d), v, p)
			}

		case ARLDCR, ARLDCRCC:
			a = int(mask[1]) /* ME */
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
			o1 |= 1 << 5 /* mb[5] is top bit */
		}

	case 30: /* rldimi $sh,s,$mask,a */
		v := c.regoff(&p.From)

		d := c.vregoff(p.GetFrom3())

		// Original opcodes had mask operands which had to be converted to a shift count as expected by
		// the ppc64 asm.
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
				o1 |= 1 << 5 /* mb[5] is top bit */
			}

		// Opcodes with shift count operands.
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

	case 31: /* dword */
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

	case 32: /* fmul frc,fra,frd */
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), 0) | (uint32(p.From.Reg)&31)<<6

	case 33: /* fabs [frb,]frd; fmr. frb,frd */
		r := int(p.From.Reg)

		if oclass(&p.From) == C_NONE {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), 0, uint32(r))

	case 34: /* FMADDx fra,frb,frc,frt (t=a*c±b) */
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg)) | (uint32(p.GetFrom3().Reg)&31)<<6

	case 35: /* mov r,lext/lauto/loreg ==> cau $(v>>16),sb,r'; store o(r') */
		v := c.regoff(&p.To)

		r := int(p.To.Reg)
		if r == 0 {
			r = int(o.param)
		}
		// Offsets in DS form stores must be a multiple of 4
		inst := c.opstore(p.As)
		if c.opform(inst) == DS_FORM && v&0x3 != 0 {
			log.Fatalf("invalid offset for DS form load/store %v", p)
		}
		o1 = AOP_IRR(OP_ADDIS, REGTMP, uint32(r), uint32(high16adjusted(v)))
		o2 = AOP_IRR(inst, uint32(p.From.Reg), REGTMP, uint32(v))

	case 36: /* mov bz/h/hz lext/lauto/lreg,r ==> lbz/lha/lhz etc */
		v := c.regoff(&p.From)

		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		o1 = AOP_IRR(OP_ADDIS, REGTMP, uint32(r), uint32(high16adjusted(v)))
		o2 = AOP_IRR(c.opload(p.As), uint32(p.To.Reg), REGTMP, uint32(v))

	case 37: /* movb lext/lauto/lreg,r ==> lbz o(reg),r; extsb r */
		v := c.regoff(&p.From)

		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		o1 = AOP_IRR(OP_ADDIS, REGTMP, uint32(r), uint32(high16adjusted(v)))
		o2 = AOP_IRR(c.opload(p.As), uint32(p.To.Reg), REGTMP, uint32(v))
		o3 = LOP_RRR(OP_EXTSB, uint32(p.To.Reg), uint32(p.To.Reg), 0)

	case 40: /* word */
		o1 = uint32(c.regoff(&p.From))

	case 41: /* stswi */
		o1 = AOP_RRR(c.opirr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), 0) | (uint32(c.regoff(p.GetFrom3()))&0x7F)<<11

	case 42: /* lswi */
		o1 = AOP_RRR(c.opirr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), 0) | (uint32(c.regoff(p.GetFrom3()))&0x7F)<<11

	case 43: /* data cache instructions: op (Ra+[Rb]), [th|l] */
		/* TH field for dcbt/dcbtst: */
		/* 0 = Block access - program will soon access EA. */
		/* 8-15 = Stream access - sequence of access (data stream). See section 4.3.2 of the ISA for details. */
		/* 16 = Block access - program will soon make a transient access to EA. */
		/* 17 = Block access - program will not access EA for a long time. */

		/* L field for dcbf: */
		/* 0 = invalidates the block containing EA in all processors. */
		/* 1 = same as 0, but with limited scope (i.e. block in the current processor will not be reused soon). */
		/* 3 = same as 1, but with even more limited scope (i.e. block in the current processor primary cache will not be reused soon). */
		if p.To.Type == obj.TYPE_NONE {
			o1 = AOP_RRR(c.oprrr(p.As), 0, uint32(p.From.Index), uint32(p.From.Reg))
		} else {
			th := c.regoff(&p.To)
			o1 = AOP_RRR(c.oprrr(p.As), uint32(th), uint32(p.From.Index), uint32(p.From.Reg))
		}

	case 44: /* indexed store */
		o1 = AOP_RRR(c.opstorex(p.As), uint32(p.From.Reg), uint32(p.To.Index), uint32(p.To.Reg))

	case 45: /* indexed load */
		switch p.As {
		/* The assembler accepts a 4-operand l*arx instruction. The fourth operand is an Exclusive Access Hint (EH) */
		/* The EH field can be used as a lock acquire/release hint as follows: */
		/* 0 = Atomic Update (fetch-and-operate or similar algorithm) */
		/* 1 = Exclusive Access (lock acquire and release) */
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
	case 46: /* plain op */
		o1 = c.oprrr(p.As)

	case 47: /* op Ra, Rd; also op [Ra,] Rd */
		r := int(p.From.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), 0)

	case 48: /* op Rs, Ra */
		r := int(p.From.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = LOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(r), 0)

	case 49: /* op Rb; op $n, Rb */
		if p.From.Type != obj.TYPE_REG { /* tlbie $L, rB */
			v := c.regoff(&p.From) & 1
			o1 = AOP_RRR(c.oprrr(p.As), 0, 0, uint32(p.To.Reg)) | uint32(v)<<21
		} else {
			o1 = AOP_RRR(c.oprrr(p.As), 0, 0, uint32(p.From.Reg))
		}

	case 50: /* rem[u] r1[,r2],r3 */
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		v := c.oprrr(p.As)
		t := v & (1<<10 | 1) /* OE|Rc */
		o1 = AOP_RRR(v&^t, REGTMP, uint32(r), uint32(p.From.Reg))
		o2 = AOP_RRR(OP_MULLW, REGTMP, REGTMP, uint32(p.From.Reg))
		o3 = AOP_RRR(OP_SUBF|t, uint32(p.To.Reg), REGTMP, uint32(r))
		if p.As == AREMU {
			o4 = o3

			/* Clear top 32 bits */
			o3 = OP_RLW(OP_RLDIC, REGTMP, REGTMP, 0, 0, 0) | 1<<5
		}

	case 51: /* remd[u] r1[,r2],r3 */
		r := int(p.Reg)

		if r == 0 {
			r = int(p.To.Reg)
		}
		v := c.oprrr(p.As)
		t := v & (1<<10 | 1) /* OE|Rc */
		o1 = AOP_RRR(v&^t, REGTMP, uint32(r), uint32(p.From.Reg))
		o2 = AOP_RRR(OP_MULLD, REGTMP, REGTMP, uint32(p.From.Reg))
		o3 = AOP_RRR(OP_SUBF|t, uint32(p.To.Reg), REGTMP, uint32(r))

	case 52: /* mtfsbNx cr(n) */
		v := c.regoff(&p.From) & 31

		o1 = AOP_RRR(c.oprrr(p.As), uint32(v), 0, 0)

	case 53: /* mffsX ,fr1 */
		o1 = AOP_RRR(OP_MFFS, uint32(p.To.Reg), 0, 0)

	case 54: /* mov msr,r1; mov r1, msr*/
		if oclass(&p.From) == C_REG {
			if p.As == AMOVD {
				o1 = AOP_RRR(OP_MTMSRD, uint32(p.From.Reg), 0, 0)
			} else {
				o1 = AOP_RRR(OP_MTMSR, uint32(p.From.Reg), 0, 0)
			}
		} else {
			o1 = AOP_RRR(OP_MFMSR, uint32(p.To.Reg), 0, 0)
		}

	case 55: /* op Rb, Rd */
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), 0, uint32(p.From.Reg))

	case 56: /* sra $sh,[s,]a; srd $sh,[s,]a */
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = AOP_RRR(c.opirr(p.As), uint32(r), uint32(p.To.Reg), uint32(v)&31)
		if (p.As == ASRAD || p.As == ASRADCC) && (v&0x20 != 0) {
			o1 |= 1 << 1 /* mb[5] */
		}

	case 57: /* slw $sh,[s,]a -> rlwinm ... */
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}

		/*
			 * Let user (gs) shoot himself in the foot.
			 * qc has already complained.
			 *
			if(v < 0 || v > 31)
				ctxt->diag("illegal shift %ld\n%v", v, p);
		*/
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
			o1 |= 1 // set the condition code
		}

	case 58: /* logical $andcon,[s],a */
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 = LOP_IRR(c.opirr(p.As), uint32(p.To.Reg), uint32(r), uint32(v))

	case 59: /* or/xor/and $ucon,,r | oris/xoris/andis $addcon,r,r */
		v := c.regoff(&p.From)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		switch p.As {
		case AOR:
			o1 = LOP_IRR(c.opirr(AORIS), uint32(p.To.Reg), uint32(r), uint32(v)>>16) /* oris, xoris, andis. */
		case AXOR:
			o1 = LOP_IRR(c.opirr(AXORIS), uint32(p.To.Reg), uint32(r), uint32(v)>>16)
		case AANDCC:
			o1 = LOP_IRR(c.opirr(AANDISCC), uint32(p.To.Reg), uint32(r), uint32(v)>>16)
		default:
			o1 = LOP_IRR(c.opirr(p.As), uint32(p.To.Reg), uint32(r), uint32(v))
		}

	case 60: /* tw to,a,b */
		r := int(c.regoff(&p.From) & 31)

		o1 = AOP_RRR(c.oprrr(p.As), uint32(r), uint32(p.Reg), uint32(p.To.Reg))

	case 61: /* tw to,a,$simm */
		r := int(c.regoff(&p.From) & 31)

		v := c.regoff(&p.To)
		o1 = AOP_IRR(c.opirr(p.As), uint32(r), uint32(p.Reg), uint32(v))

	case 62: /* rlwmi $sh,s,$mask,a */
		v := c.regoff(&p.From)

		var mask [2]uint8
		c.maskgen(p, mask[:], uint32(c.regoff(p.GetFrom3())))
		o1 = AOP_RRR(c.opirr(p.As), uint32(p.Reg), uint32(p.To.Reg), uint32(v))
		o1 |= (uint32(mask[0])&31)<<6 | (uint32(mask[1])&31)<<1

	case 63: /* rlwmi b,s,$mask,a */
		var mask [2]uint8
		c.maskgen(p, mask[:], uint32(c.regoff(p.GetFrom3())))

		o1 = AOP_RRR(c.opirr(p.As), uint32(p.Reg), uint32(p.To.Reg), uint32(p.From.Reg))
		o1 |= (uint32(mask[0])&31)<<6 | (uint32(mask[1])&31)<<1

	case 64: /* mtfsf fr[, $m] {,fpcsr} */
		var v int32
		if p.From3Type() != obj.TYPE_NONE {
			v = c.regoff(p.GetFrom3()) & 255
		} else {
			v = 255
		}
		o1 = OP_MTFSF | uint32(v)<<17 | uint32(p.From.Reg)<<11

	case 65: /* MOVFL $imm,FPSCR(n) => mtfsfi crfd,imm */
		if p.To.Reg == 0 {
			c.ctxt.Diag("must specify FPSCR(n)\n%v", p)
		}
		o1 = OP_MTFSFI | (uint32(p.To.Reg)&15)<<23 | (uint32(c.regoff(&p.From))&31)<<12

	case 66: /* mov spr,r1; mov r1,spr, also dcr */
		var r int
		var v int32
		if REG_R0 <= p.From.Reg && p.From.Reg <= REG_R31 {
			r = int(p.From.Reg)
			v = int32(p.To.Reg)
			if REG_DCR0 <= v && v <= REG_DCR0+1023 {
				o1 = OPVCC(31, 451, 0, 0) /* mtdcr */
			} else {
				o1 = OPVCC(31, 467, 0, 0) /* mtspr */
			}
		} else {
			r = int(p.To.Reg)
			v = int32(p.From.Reg)
			if REG_DCR0 <= v && v <= REG_DCR0+1023 {
				o1 = OPVCC(31, 323, 0, 0) /* mfdcr */
			} else {
				o1 = OPVCC(31, 339, 0, 0) /* mfspr */
			}
		}

		o1 = AOP_RRR(o1, uint32(r), 0, 0) | (uint32(v)&0x1f)<<16 | ((uint32(v)>>5)&0x1f)<<11

	case 67: /* mcrf crfD,crfS */
		if p.From.Type != obj.TYPE_REG || p.From.Reg < REG_CR0 || REG_CR7 < p.From.Reg || p.To.Type != obj.TYPE_REG || p.To.Reg < REG_CR0 || REG_CR7 < p.To.Reg {
			c.ctxt.Diag("illegal CR field number\n%v", p)
		}
		o1 = AOP_RRR(OP_MCRF, ((uint32(p.To.Reg) & 7) << 2), ((uint32(p.From.Reg) & 7) << 2), 0)

	case 68: /* mfcr rD; mfocrf CRM,rD */
		if p.From.Type == obj.TYPE_REG && REG_CR0 <= p.From.Reg && p.From.Reg <= REG_CR7 {
			v := int32(1 << uint(7-(p.To.Reg&7)))                                 /* CR(n) */
			o1 = AOP_RRR(OP_MFCR, uint32(p.To.Reg), 0, 0) | 1<<20 | uint32(v)<<12 /* new form, mfocrf */
		} else {
			o1 = AOP_RRR(OP_MFCR, uint32(p.To.Reg), 0, 0) /* old form, whole register */
		}

	case 69: /* mtcrf CRM,rS */
		var v int32
		if p.From3Type() != obj.TYPE_NONE {
			if p.To.Reg != 0 {
				c.ctxt.Diag("can't use both mask and CR(n)\n%v", p)
			}
			v = c.regoff(p.GetFrom3()) & 0xff
		} else {
			if p.To.Reg == 0 {
				v = 0xff /* CR */
			} else {
				v = 1 << uint(7-(p.To.Reg&7)) /* CR(n) */
			}
		}

		o1 = AOP_RRR(OP_MTCRF, uint32(p.From.Reg), 0, 0) | uint32(v)<<12

	case 70: /* [f]cmp r,r,cr*/
		var r int
		if p.Reg == 0 {
			r = 0
		} else {
			r = (int(p.Reg) & 7) << 2
		}
		o1 = AOP_RRR(c.oprrr(p.As), uint32(r), uint32(p.From.Reg), uint32(p.To.Reg))

	case 71: /* cmp[l] r,i,cr*/
		var r int
		if p.Reg == 0 {
			r = 0
		} else {
			r = (int(p.Reg) & 7) << 2
		}
		o1 = AOP_RRR(c.opirr(p.As), uint32(r), uint32(p.From.Reg), 0) | uint32(c.regoff(&p.To))&0xffff

	case 72: /* slbmte (Rb+Rs -> slb[Rb]) -> Rs, Rb */
		o1 = AOP_RRR(c.oprrr(p.As), uint32(p.From.Reg), 0, uint32(p.To.Reg))

	case 73: /* mcrfs crfD,crfS */
		if p.From.Type != obj.TYPE_REG || p.From.Reg != REG_FPSCR || p.To.Type != obj.TYPE_REG || p.To.Reg < REG_CR0 || REG_CR7 < p.To.Reg {
			c.ctxt.Diag("illegal FPSCR/CR field number\n%v", p)
		}
		o1 = AOP_RRR(OP_MCRFS, ((uint32(p.To.Reg) & 7) << 2), ((0 & 7) << 2), 0)

	case 77: /* syscall $scon, syscall Rx */
		if p.From.Type == obj.TYPE_CONST {
			if p.From.Offset > BIG || p.From.Offset < -BIG {
				c.ctxt.Diag("illegal syscall, sysnum too large: %v", p)
			}
			o1 = AOP_IRR(OP_ADDI, REGZERO, REGZERO, uint32(p.From.Offset))
		} else if p.From.Type == obj.TYPE_REG {
			o1 = LOP_RRR(OP_OR, REGZERO, uint32(p.From.Reg), uint32(p.From.Reg))
		} else {
			c.ctxt.Diag("illegal syscall: %v", p)
			o1 = 0x7fe00008 // trap always
		}

		o2 = c.oprrr(p.As)
		o3 = AOP_RRR(c.oprrr(AXOR), REGZERO, REGZERO, REGZERO) // XOR R0, R0

	case 78: /* undef */
		o1 = 0 /* "An instruction consisting entirely of binary 0s is guaranteed
		   always to be an illegal instruction."  */

	/* relocation operations */
	case 74:
		v := c.vregoff(&p.To)
		// Offsets in DS form stores must be a multiple of 4
		inst := c.opstore(p.As)
		if c.opform(inst) == DS_FORM && v&0x3 != 0 {
			log.Fatalf("invalid offset for DS form load/store %v", p)
		}
		o1, o2 = c.symbolAccess(p.To.Sym, v, p.From.Reg, inst)

	//if(dlm) reloc(&p->to, p->pc, 1);

	case 75:
		v := c.vregoff(&p.From)
		// Offsets in DS form loads must be a multiple of 4
		inst := c.opload(p.As)
		if c.opform(inst) == DS_FORM && v&0x3 != 0 {
			log.Fatalf("invalid offset for DS form load/store %v", p)
		}
		o1, o2 = c.symbolAccess(p.From.Sym, v, p.To.Reg, inst)

	//if(dlm) reloc(&p->from, p->pc, 1);

	case 76:
		v := c.vregoff(&p.From)
		// Offsets in DS form loads must be a multiple of 4
		inst := c.opload(p.As)
		if c.opform(inst) == DS_FORM && v&0x3 != 0 {
			log.Fatalf("invalid offset for DS form load/store %v", p)
		}
		o1, o2 = c.symbolAccess(p.From.Sym, v, p.To.Reg, inst)
		o3 = LOP_RRR(OP_EXTSB, uint32(p.To.Reg), uint32(p.To.Reg), 0)

	//if(dlm) reloc(&p->from, p->pc, 1);

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
	case 82: /* vector instructions, VX-form and VC-form */
		if p.From.Type == obj.TYPE_REG {
			/* reg reg none OR reg reg reg */
			/* 3-register operand order: VRA, VRB, VRT */
			/* 2-register operand order: VRA, VRT */
			o1 = AOP_RRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg))
		} else if p.From3Type() == obj.TYPE_CONST {
			/* imm imm reg reg */
			/* operand order: SIX, VRA, ST, VRT */
			six := int(c.regoff(&p.From))
			st := int(c.regoff(p.GetFrom3()))
			o1 = AOP_IIRR(c.opiirr(p.As), uint32(p.To.Reg), uint32(p.Reg), uint32(st), uint32(six))
		} else if p.From3Type() == obj.TYPE_NONE && p.Reg != 0 {
			/* imm reg reg */
			/* operand order: UIM, VRB, VRT */
			uim := int(c.regoff(&p.From))
			o1 = AOP_VIRR(c.opirr(p.As), uint32(p.To.Reg), uint32(p.Reg), uint32(uim))
		} else {
			/* imm reg */
			/* operand order: SIM, VRT */
			sim := int(c.regoff(&p.From))
			o1 = AOP_IR(c.opirr(p.As), uint32(p.To.Reg), uint32(sim))
		}

	case 83: /* vector instructions, VA-form */
		if p.From.Type == obj.TYPE_REG {
			/* reg reg reg reg */
			/* 4-register operand order: VRA, VRB, VRC, VRT */
			o1 = AOP_RRRR(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg), uint32(p.GetFrom3().Reg))
		} else if p.From.Type == obj.TYPE_CONST {
			/* imm reg reg reg */
			/* operand order: SHB, VRA, VRB, VRT */
			shb := int(c.regoff(&p.From))
			o1 = AOP_IRRR(c.opirrr(p.As), uint32(p.To.Reg), uint32(p.Reg), uint32(p.GetFrom3().Reg), uint32(shb))
		}

	case 84: // ISEL BC,RA,RB,RT -> isel rt,ra,rb,bc
		bc := c.vregoff(&p.From)

		// rt = To.Reg, ra = p.Reg, rb = p.From3.Reg
		o1 = AOP_ISEL(OP_ISEL, uint32(p.To.Reg), uint32(p.Reg), uint32(p.GetFrom3().Reg), uint32(bc))

	case 85: /* vector instructions, VX-form */
		/* reg none reg */
		/* 2-register operand order: VRB, VRT */
		o1 = AOP_RR(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg))

	case 86: /* VSX indexed store, XX1-form */
		/* reg reg reg */
		/* 3-register operand order: XT, (RB)(RA*1) */
		o1 = AOP_XX1(c.opstorex(p.As), uint32(p.From.Reg), uint32(p.To.Index), uint32(p.To.Reg))

	case 87: /* VSX indexed load, XX1-form */
		/* reg reg reg */
		/* 3-register operand order: (RB)(RA*1), XT */
		o1 = AOP_XX1(c.oploadx(p.As), uint32(p.To.Reg), uint32(p.From.Index), uint32(p.From.Reg))

	case 88: /* VSX instructions, XX1-form */
		/* reg reg none OR reg reg reg */
		/* 3-register operand order: RA, RB, XT */
		/* 2-register operand order: XS, RA or RA, XT */
		xt := int32(p.To.Reg)
		xs := int32(p.From.Reg)
		/* We need to treat the special case of extended mnemonics that may have a FREG/VREG as an argument */
		if REG_V0 <= xt && xt <= REG_V31 {
			/* Convert V0-V31 to VS32-VS63 */
			xt = xt + 64
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg))
		} else if REG_F0 <= xt && xt <= REG_F31 {
			/* Convert F0-F31 to VS0-VS31 */
			xt = xt + 64
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg))
		} else if REG_VS0 <= xt && xt <= REG_VS63 {
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg))
		} else if REG_V0 <= xs && xs <= REG_V31 {
			/* Likewise for XS */
			xs = xs + 64
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), uint32(p.Reg))
		} else if REG_F0 <= xs && xs <= REG_F31 {
			xs = xs + 64
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), uint32(p.Reg))
		} else if REG_VS0 <= xs && xs <= REG_VS63 {
			o1 = AOP_XX1(c.oprrr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), uint32(p.Reg))
		}

	case 89: /* VSX instructions, XX2-form */
		/* reg none reg OR reg imm reg */
		/* 2-register operand order: XB, XT or XB, UIM, XT*/
		uim := int(c.regoff(p.GetFrom3()))
		o1 = AOP_XX2(c.oprrr(p.As), uint32(p.To.Reg), uint32(uim), uint32(p.From.Reg))

	case 90: /* VSX instructions, XX3-form */
		if p.From3Type() == obj.TYPE_NONE {
			/* reg reg reg */
			/* 3-register operand order: XA, XB, XT */
			o1 = AOP_XX3(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg))
		} else if p.From3Type() == obj.TYPE_CONST {
			/* reg reg reg imm */
			/* operand order: XA, XB, DM, XT */
			dm := int(c.regoff(p.GetFrom3()))
			o1 = AOP_XX3I(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg), uint32(dm))
		}

	case 91: /* VSX instructions, XX4-form */
		/* reg reg reg reg */
		/* 3-register operand order: XA, XB, XC, XT */
		o1 = AOP_XX4(c.oprrr(p.As), uint32(p.To.Reg), uint32(p.From.Reg), uint32(p.Reg), uint32(p.GetFrom3().Reg))

	case 92: /* X-form instructions, 3-operands */
		if p.To.Type == obj.TYPE_CONST {
			/* imm reg reg */
			xf := int32(p.From.Reg)
			if REG_F0 <= xf && xf <= REG_F31 {
				/* operand order: FRA, FRB, BF */
				bf := int(c.regoff(&p.To)) << 2
				o1 = AOP_RRR(c.opirr(p.As), uint32(bf), uint32(p.From.Reg), uint32(p.Reg))
			} else {
				/* operand order: RA, RB, L */
				l := int(c.regoff(&p.To))
				o1 = AOP_RRR(c.opirr(p.As), uint32(l), uint32(p.From.Reg), uint32(p.Reg))
			}
		} else if p.From3Type() == obj.TYPE_CONST {
			/* reg reg imm */
			/* operand order: RB, L, RA */
			l := int(c.regoff(p.GetFrom3()))
			o1 = AOP_RRR(c.opirr(p.As), uint32(l), uint32(p.To.Reg), uint32(p.From.Reg))
		} else if p.To.Type == obj.TYPE_REG {
			cr := int32(p.To.Reg)
			if REG_CR0 <= cr && cr <= REG_CR7 {
				/* cr reg reg */
				/* operand order: RA, RB, BF */
				bf := (int(p.To.Reg) & 7) << 2
				o1 = AOP_RRR(c.opirr(p.As), uint32(bf), uint32(p.From.Reg), uint32(p.Reg))
			} else if p.From.Type == obj.TYPE_CONST {
				/* reg imm */
				/* operand order: L, RT */
				l := int(c.regoff(&p.From))
				o1 = AOP_RRR(c.opirr(p.As), uint32(p.To.Reg), uint32(l), uint32(p.Reg))
			} else {
				switch p.As {
				case ACOPY, APASTECC:
					o1 = AOP_RRR(c.opirr(p.As), uint32(1), uint32(p.From.Reg), uint32(p.To.Reg))
				default:
					/* reg reg reg */
					/* operand order: RS, RB, RA */
					o1 = AOP_RRR(c.oprrr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), uint32(p.Reg))
				}
			}
		}

	case 93: /* X-form instructions, 2-operands */
		if p.To.Type == obj.TYPE_CONST {
			/* imm reg */
			/* operand order: FRB, BF */
			bf := int(c.regoff(&p.To)) << 2
			o1 = AOP_RR(c.opirr(p.As), uint32(bf), uint32(p.From.Reg))
		} else if p.Reg == 0 {
			/* popcnt* r,r, X-form */
			/* operand order: RS, RA */
			o1 = AOP_RRR(c.oprrr(p.As), uint32(p.From.Reg), uint32(p.To.Reg), uint32(p.Reg))
		}

	case 94: /* Z23-form instructions, 4-operands */
		/* reg reg reg imm */
		/* operand order: RA, RB, CY, RT */
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
		return OPVCC(31, 170, 0, 0) /* addex - v3.0b */

	case AAND:
		return OPVCC(31, 28, 0, 0)
	case AANDCC:
		return OPVCC(31, 28, 0, 1)
	case AANDN:
		return OPVCC(31, 60, 0, 0)
	case AANDNCC:
		return OPVCC(31, 60, 0, 1)

	case ACMP:
		return OPVCC(31, 0, 0, 0) | 1<<21 /* L=1 */
	case ACMPU:
		return OPVCC(31, 32, 0, 0) | 1<<21
	case ACMPW:
		return OPVCC(31, 0, 0, 0) /* L=0 */
	case ACMPWU:
		return OPVCC(31, 32, 0, 0)
	case ACMPB:
		return OPVCC(31, 508, 0, 0) /* cmpb - v2.05 */

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
		return OPVCC(63, 72, 0, 0) /* load */
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
		return OPVCC(31, 506, 0, 0) /* popcntd - v2.06 */
	case APOPCNTW:
		return OPVCC(31, 378, 0, 0) /* popcntw - v2.06 */
	case APOPCNTB:
		return OPVCC(31, 122, 0, 0) /* popcntb - v2.02 */

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
		return OPVCC(30, 0, 0, 0) | 2<<1 // rldicr
	case ARLDICRCC:
		return OPVCC(30, 0, 0, 1) | 2<<1 // rldicr.

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

	/* Vector (VMX/Altivec) instructions */
	/* ISA 2.03 enables these for PPC970. For POWERx processors, these */
	/* are enabled starting at POWER6 (ISA 2.05). */
	case AVAND:
		return OPVX(4, 1028, 0, 0) /* vand - v2.03 */
	case AVANDC:
		return OPVX(4, 1092, 0, 0) /* vandc - v2.03 */
	case AVNAND:
		return OPVX(4, 1412, 0, 0) /* vnand - v2.07 */

	case AVOR:
		return OPVX(4, 1156, 0, 0) /* vor - v2.03 */
	case AVORC:
		return OPVX(4, 1348, 0, 0) /* vorc - v2.07 */
	case AVNOR:
		return OPVX(4, 1284, 0, 0) /* vnor - v2.03 */
	case AVXOR:
		return OPVX(4, 1220, 0, 0) /* vxor - v2.03 */
	case AVEQV:
		return OPVX(4, 1668, 0, 0) /* veqv - v2.07 */

	case AVADDUBM:
		return OPVX(4, 0, 0, 0) /* vaddubm - v2.03 */
	case AVADDUHM:
		return OPVX(4, 64, 0, 0) /* vadduhm - v2.03 */
	case AVADDUWM:
		return OPVX(4, 128, 0, 0) /* vadduwm - v2.03 */
	case AVADDUDM:
		return OPVX(4, 192, 0, 0) /* vaddudm - v2.07 */
	case AVADDUQM:
		return OPVX(4, 256, 0, 0) /* vadduqm - v2.07 */

	case AVADDCUQ:
		return OPVX(4, 320, 0, 0) /* vaddcuq - v2.07 */
	case AVADDCUW:
		return OPVX(4, 384, 0, 0) /* vaddcuw - v2.03 */

	case AVADDUBS:
		return OPVX(4, 512, 0, 0) /* vaddubs - v2.03 */
	case AVADDUHS:
		return OPVX(4, 576, 0, 0) /* vadduhs - v2.03 */
	case AVADDUWS:
		return OPVX(4, 640, 0, 0) /* vadduws - v2.03 */

	case AVADDSBS:
		return OPVX(4, 768, 0, 0) /* vaddsbs - v2.03 */
	case AVADDSHS:
		return OPVX(4, 832, 0, 0) /* vaddshs - v2.03 */
	case AVADDSWS:
		return OPVX(4, 896, 0, 0) /* vaddsws - v2.03 */

	case AVADDEUQM:
		return OPVX(4, 60, 0, 0) /* vaddeuqm - v2.07 */
	case AVADDECUQ:
		return OPVX(4, 61, 0, 0) /* vaddecuq - v2.07 */

	case AVMULESB:
		return OPVX(4, 776, 0, 0) /* vmulesb - v2.03 */
	case AVMULOSB:
		return OPVX(4, 264, 0, 0) /* vmulosb - v2.03 */
	case AVMULEUB:
		return OPVX(4, 520, 0, 0) /* vmuleub - v2.03 */
	case AVMULOUB:
		return OPVX(4, 8, 0, 0) /* vmuloub - v2.03 */
	case AVMULESH:
		return OPVX(4, 840, 0, 0) /* vmulesh - v2.03 */
	case AVMULOSH:
		return OPVX(4, 328, 0, 0) /* vmulosh - v2.03 */
	case AVMULEUH:
		return OPVX(4, 584, 0, 0) /* vmuleuh - v2.03 */
	case AVMULOUH:
		return OPVX(4, 72, 0, 0) /* vmulouh - v2.03 */
	case AVMULESW:
		return OPVX(4, 904, 0, 0) /* vmulesw - v2.07 */
	case AVMULOSW:
		return OPVX(4, 392, 0, 0) /* vmulosw - v2.07 */
	case AVMULEUW:
		return OPVX(4, 648, 0, 0) /* vmuleuw - v2.07 */
	case AVMULOUW:
		return OPVX(4, 136, 0, 0) /* vmulouw - v2.07 */
	case AVMULUWM:
		return OPVX(4, 137, 0, 0) /* vmuluwm - v2.07 */

	case AVPMSUMB:
		return OPVX(4, 1032, 0, 0) /* vpmsumb - v2.07 */
	case AVPMSUMH:
		return OPVX(4, 1096, 0, 0) /* vpmsumh - v2.07 */
	case AVPMSUMW:
		return OPVX(4, 1160, 0, 0) /* vpmsumw - v2.07 */
	case AVPMSUMD:
		return OPVX(4, 1224, 0, 0) /* vpmsumd - v2.07 */

	case AVMSUMUDM:
		return OPVX(4, 35, 0, 0) /* vmsumudm - v3.00b */

	case AVSUBUBM:
		return OPVX(4, 1024, 0, 0) /* vsububm - v2.03 */
	case AVSUBUHM:
		return OPVX(4, 1088, 0, 0) /* vsubuhm - v2.03 */
	case AVSUBUWM:
		return OPVX(4, 1152, 0, 0) /* vsubuwm - v2.03 */
	case AVSUBUDM:
		return OPVX(4, 1216, 0, 0) /* vsubudm - v2.07 */
	case AVSUBUQM:
		return OPVX(4, 1280, 0, 0) /* vsubuqm - v2.07 */

	case AVSUBCUQ:
		return OPVX(4, 1344, 0, 0) /* vsubcuq - v2.07 */
	case AVSUBCUW:
		return OPVX(4, 1408, 0, 0) /* vsubcuw - v2.03 */

	case AVSUBUBS:
		return OPVX(4, 1536, 0, 0) /* vsububs - v2.03 */
	case AVSUBUHS:
		return OPVX(4, 1600, 0, 0) /* vsubuhs - v2.03 */
	case AVSUBUWS:
		return OPVX(4, 1664, 0, 0) /* vsubuws - v2.03 */

	case AVSUBSBS:
		return OPVX(4, 1792, 0, 0) /* vsubsbs - v2.03 */
	case AVSUBSHS:
		return OPVX(4, 1856, 0, 0) /* vsubshs - v2.03 */
	case AVSUBSWS:
		return OPVX(4, 1920, 0, 0) /* vsubsws - v2.03 */

	case AVSUBEUQM:
		return OPVX(4, 62, 0, 0) /* vsubeuqm - v2.07 */
	case AVSUBECUQ:
		return OPVX(4, 63, 0, 0) /* vsubecuq - v2.07 */

	case AVRLB:
		return OPVX(4, 4, 0, 0) /* vrlb - v2.03 */
	case AVRLH:
		return OPVX(4, 68, 0, 0) /* vrlh - v2.03 */
	case AVRLW:
		return OPVX(4, 132, 0, 0) /* vrlw - v2.03 */
	case AVRLD:
		return OPVX(4, 196, 0, 0) /* vrld - v2.07 */

	case AVSLB:
		return OPVX(4, 260, 0, 0) /* vslh - v2.03 */
	case AVSLH:
		return OPVX(4, 324, 0, 0) /* vslh - v2.03 */
	case AVSLW:
		return OPVX(4, 388, 0, 0) /* vslw - v2.03 */
	case AVSL:
		return OPVX(4, 452, 0, 0) /* vsl - v2.03 */
	case AVSLO:
		return OPVX(4, 1036, 0, 0) /* vsl - v2.03 */
	case AVSRB:
		return OPVX(4, 516, 0, 0) /* vsrb - v2.03 */
	case AVSRH:
		return OPVX(4, 580, 0, 0) /* vsrh - v2.03 */
	case AVSRW:
		return OPVX(4, 644, 0, 0) /* vsrw - v2.03 */
	case AVSR:
		return OPVX(4, 708, 0, 0) /* vsr - v2.03 */
	case AVSRO:
		return OPVX(4, 1100, 0, 0) /* vsro - v2.03 */
	case AVSLD:
		return OPVX(4, 1476, 0, 0) /* vsld - v2.07 */
	case AVSRD:
		return OPVX(4, 1732, 0, 0) /* vsrd - v2.07 */

	case AVSRAB:
		return OPVX(4, 772, 0, 0) /* vsrab - v2.03 */
	case AVSRAH:
		return OPVX(4, 836, 0, 0) /* vsrah - v2.03 */
	case AVSRAW:
		return OPVX(4, 900, 0, 0) /* vsraw - v2.03 */
	case AVSRAD:
		return OPVX(4, 964, 0, 0) /* vsrad - v2.07 */

	case AVBPERMQ:
		return OPVC(4, 1356, 0, 0) /* vbpermq - v2.07 */
	case AVBPERMD:
		return OPVC(4, 1484, 0, 0) /* vbpermd - v3.00 */

	case AVCLZB:
		return OPVX(4, 1794, 0, 0) /* vclzb - v2.07 */
	case AVCLZH:
		return OPVX(4, 1858, 0, 0) /* vclzh - v2.07 */
	case AVCLZW:
		return OPVX(4, 1922, 0, 0) /* vclzw - v2.07 */
	case AVCLZD:
		return OPVX(4, 1986, 0, 0) /* vclzd - v2.07 */

	case AVPOPCNTB:
		return OPVX(4, 1795, 0, 0) /* vpopcntb - v2.07 */
	case AVPOPCNTH:
		return OPVX(4, 1859, 0, 0) /* vpopcnth - v2.07 */
	case AVPOPCNTW:
		return OPVX(4, 1923, 0, 0) /* vpopcntw - v2.07 */
	case AVPOPCNTD:
		return OPVX(4, 1987, 0, 0) /* vpopcntd - v2.07 */

	case AVCMPEQUB:
		return OPVC(4, 6, 0, 0) /* vcmpequb - v2.03 */
	case AVCMPEQUBCC:
		return OPVC(4, 6, 0, 1) /* vcmpequb. - v2.03 */
	case AVCMPEQUH:
		return OPVC(4, 70, 0, 0) /* vcmpequh - v2.03 */
	case AVCMPEQUHCC:
		return OPVC(4, 70, 0, 1) /* vcmpequh. - v2.03 */
	case AVCMPEQUW:
		return OPVC(4, 134, 0, 0) /* vcmpequw - v2.03 */
	case AVCMPEQUWCC:
		return OPVC(4, 134, 0, 1) /* vcmpequw. - v2.03 */
	case AVCMPEQUD:
		return OPVC(4, 199, 0, 0) /* vcmpequd - v2.07 */
	case AVCMPEQUDCC:
		return OPVC(4, 199, 0, 1) /* vcmpequd. - v2.07 */

	case AVCMPGTUB:
		return OPVC(4, 518, 0, 0) /* vcmpgtub - v2.03 */
	case AVCMPGTUBCC:
		return OPVC(4, 518, 0, 1) /* vcmpgtub. - v2.03 */
	case AVCMPGTUH:
		return OPVC(4, 582, 0, 0) /* vcmpgtuh - v2.03 */
	case AVCMPGTUHCC:
		return OPVC(4, 582, 0, 1) /* vcmpgtuh. - v2.03 */
	case AVCMPGTUW:
		return OPVC(4, 646, 0, 0) /* vcmpgtuw - v2.03 */
	case AVCMPGTUWCC:
		return OPVC(4, 646, 0, 1) /* vcmpgtuw. - v2.03 */
	case AVCMPGTUD:
		return OPVC(4, 711, 0, 0) /* vcmpgtud - v2.07 */
	case AVCMPGTUDCC:
		return OPVC(4, 711, 0, 1) /* vcmpgtud. v2.07 */
	case AVCMPGTSB:
		return OPVC(4, 774, 0, 0) /* vcmpgtsb - v2.03 */
	case AVCMPGTSBCC:
		return OPVC(4, 774, 0, 1) /* vcmpgtsb. - v2.03 */
	case AVCMPGTSH:
		return OPVC(4, 838, 0, 0) /* vcmpgtsh - v2.03 */
	case AVCMPGTSHCC:
		return OPVC(4, 838, 0, 1) /* vcmpgtsh. - v2.03 */
	case AVCMPGTSW:
		return OPVC(4, 902, 0, 0) /* vcmpgtsw - v2.03 */
	case AVCMPGTSWCC:
		return OPVC(4, 902, 0, 1) /* vcmpgtsw. - v2.03 */
	case AVCMPGTSD:
		return OPVC(4, 967, 0, 0) /* vcmpgtsd - v2.07 */
	case AVCMPGTSDCC:
		return OPVC(4, 967, 0, 1) /* vcmpgtsd. - v2.07 */

	case AVCMPNEZB:
		return OPVC(4, 263, 0, 0) /* vcmpnezb - v3.00 */
	case AVCMPNEZBCC:
		return OPVC(4, 263, 0, 1) /* vcmpnezb. - v3.00 */

	case AVPERM:
		return OPVX(4, 43, 0, 0) /* vperm - v2.03 */

	case AVSEL:
		return OPVX(4, 42, 0, 0) /* vsel - v2.03 */

	case AVCIPHER:
		return OPVX(4, 1288, 0, 0) /* vcipher - v2.07 */
	case AVCIPHERLAST:
		return OPVX(4, 1289, 0, 0) /* vcipherlast - v2.07 */
	case AVNCIPHER:
		return OPVX(4, 1352, 0, 0) /* vncipher - v2.07 */
	case AVNCIPHERLAST:
		return OPVX(4, 1353, 0, 0) /* vncipherlast - v2.07 */
	case AVSBOX:
		return OPVX(4, 1480, 0, 0) /* vsbox - v2.07 */
	/* End of vector instructions */

	/* Vector scalar (VSX) instructions */
	/* ISA 2.06 enables these for POWER7. */
	case AMFVSRD, AMFVRD, AMFFPRD:
		return OPVXX1(31, 51, 0) /* mfvsrd - v2.07 */
	case AMFVSRWZ:
		return OPVXX1(31, 115, 0) /* mfvsrwz - v2.07 */
	case AMFVSRLD:
		return OPVXX1(31, 307, 0) /* mfvsrld - v3.00 */

	case AMTVSRD, AMTFPRD, AMTVRD:
		return OPVXX1(31, 179, 0) /* mtvsrd - v2.07 */
	case AMTVSRWA:
		return OPVXX1(31, 211, 0) /* mtvsrwa - v2.07 */
	case AMTVSRWZ:
		return OPVXX1(31, 243, 0) /* mtvsrwz - v2.07 */
	case AMTVSRDD:
		return OPVXX1(31, 435, 0) /* mtvsrdd - v3.00 */
	case AMTVSRWS:
		return OPVXX1(31, 403, 0) /* mtvsrws - v3.00 */

	case AXXLANDQ:
		return OPVXX3(60, 130, 0) /* xxland - v2.06 */
	case AXXLANDC:
		return OPVXX3(60, 138, 0) /* xxlandc - v2.06 */
	case AXXLEQV:
		return OPVXX3(60, 186, 0) /* xxleqv - v2.07 */
	case AXXLNAND:
		return OPVXX3(60, 178, 0) /* xxlnand - v2.07 */

	case AXXLORC:
		return OPVXX3(60, 170, 0) /* xxlorc - v2.07 */
	case AXXLNOR:
		return OPVXX3(60, 162, 0) /* xxlnor - v2.06 */
	case AXXLORQ:
		return OPVXX3(60, 146, 0) /* xxlor - v2.06 */
	case AXXLXOR:
		return OPVXX3(60, 154, 0) /* xxlxor - v2.06 */

	case AXXSEL:
		return OPVXX4(60, 3, 0) /* xxsel - v2.06 */

	case AXXMRGHW:
		return OPVXX3(60, 18, 0) /* xxmrghw - v2.06 */
	case AXXMRGLW:
		return OPVXX3(60, 50, 0) /* xxmrglw - v2.06 */

	case AXXSPLTW:
		return OPVXX2(60, 164, 0) /* xxspltw - v2.06 */

	case AXXPERMDI:
		return OPVXX3(60, 10, 0) /* xxpermdi - v2.06 */

	case AXXSLDWI:
		return OPVXX3(60, 2, 0) /* xxsldwi - v2.06 */

	case AXSCVDPSP:
		return OPVXX2(60, 265, 0) /* xscvdpsp - v2.06 */
	case AXSCVSPDP:
		return OPVXX2(60, 329, 0) /* xscvspdp - v2.06 */
	case AXSCVDPSPN:
		return OPVXX2(60, 267, 0) /* xscvdpspn - v2.07 */
	case AXSCVSPDPN:
		return OPVXX2(60, 331, 0) /* xscvspdpn - v2.07 */

	case AXVCVDPSP:
		return OPVXX2(60, 393, 0) /* xvcvdpsp - v2.06 */
	case AXVCVSPDP:
		return OPVXX2(60, 457, 0) /* xvcvspdp - v2.06 */

	case AXSCVDPSXDS:
		return OPVXX2(60, 344, 0) /* xscvdpsxds - v2.06 */
	case AXSCVDPSXWS:
		return OPVXX2(60, 88, 0) /* xscvdpsxws - v2.06 */
	case AXSCVDPUXDS:
		return OPVXX2(60, 328, 0) /* xscvdpuxds - v2.06 */
	case AXSCVDPUXWS:
		return OPVXX2(60, 72, 0) /* xscvdpuxws - v2.06 */

	case AXSCVSXDDP:
		return OPVXX2(60, 376, 0) /* xscvsxddp - v2.06 */
	case AXSCVUXDDP:
		return OPVXX2(60, 360, 0) /* xscvuxddp - v2.06 */
	case AXSCVSXDSP:
		return OPVXX2(60, 312, 0) /* xscvsxdsp - v2.06 */
	case AXSCVUXDSP:
		return OPVXX2(60, 296, 0) /* xscvuxdsp - v2.06 */

	case AXVCVDPSXDS:
		return OPVXX2(60, 472, 0) /* xvcvdpsxds - v2.06 */
	case AXVCVDPSXWS:
		return OPVXX2(60, 216, 0) /* xvcvdpsxws - v2.06 */
	case AXVCVDPUXDS:
		return OPVXX2(60, 456, 0) /* xvcvdpuxds - v2.06 */
	case AXVCVDPUXWS:
		return OPVXX2(60, 200, 0) /* xvcvdpuxws - v2.06 */
	case AXVCVSPSXDS:
		return OPVXX2(60, 408, 0) /* xvcvspsxds - v2.07 */
	case AXVCVSPSXWS:
		return OPVXX2(60, 152, 0) /* xvcvspsxws - v2.07 */
	case AXVCVSPUXDS:
		return OPVXX2(60, 392, 0) /* xvcvspuxds - v2.07 */
	case AXVCVSPUXWS:
		return OPVXX2(60, 136, 0) /* xvcvspuxws - v2.07 */

	case AXVCVSXDDP:
		return OPVXX2(60, 504, 0) /* xvcvsxddp - v2.06 */
	case AXVCVSXWDP:
		return OPVXX2(60, 248, 0) /* xvcvsxwdp - v2.06 */
	case AXVCVUXDDP:
		return OPVXX2(60, 488, 0) /* xvcvuxddp - v2.06 */
	case AXVCVUXWDP:
		return OPVXX2(60, 232, 0) /* xvcvuxwdp - v2.06 */
	case AXVCVSXDSP:
		return OPVXX2(60, 440, 0) /* xvcvsxdsp - v2.06 */
	case AXVCVSXWSP:
		return OPVXX2(60, 184, 0) /* xvcvsxwsp - v2.06 */
	case AXVCVUXDSP:
		return OPVXX2(60, 424, 0) /* xvcvuxdsp - v2.06 */
	case AXVCVUXWSP:
		return OPVXX2(60, 168, 0) /* xvcvuxwsp - v2.06 */
	/* End of VSX instructions */

	case AMADDHD:
		return OPVX(4, 48, 0, 0) /* maddhd - v3.00 */
	case AMADDHDU:
		return OPVX(4, 49, 0, 0) /* maddhdu - v3.00 */
	case AMADDLD:
		return OPVX(4, 51, 0, 0) /* maddld - v3.00 */

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
	/* Vector (VMX/Altivec) instructions */
	/* ISA 2.03 enables these for PPC970. For POWERx processors, these */
	/* are enabled starting at POWER6 (ISA 2.05). */
	case AVSLDOI:
		return OPVX(4, 44, 0, 0) /* vsldoi - v2.03 */
	}

	c.ctxt.Diag("bad i/r/r/r opcode %v", a)
	return 0
}

func (c *ctxt9) opiirr(a obj.As) uint32 {
	switch a {
	/* Vector (VMX/Altivec) instructions */
	/* ISA 2.07 enables these for POWER8 and beyond. */
	case AVSHASIGMAW:
		return OPVX(4, 1666, 0, 0) /* vshasigmaw - v2.07 */
	case AVSHASIGMAD:
		return OPVX(4, 1730, 0, 0) /* vshasigmad - v2.07 */
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
		return OPVCC(15, 0, 0, 0) /* ADDIS */

	case AANDCC:
		return OPVCC(28, 0, 0, 0)
	case AANDISCC:
		return OPVCC(29, 0, 0, 0) /* ANDIS. */

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
		return AOP_RRR(16<<26, 4, 3, 0) // apparently unordered-clear
	case ABVS:
		return AOP_RRR(16<<26, 12, 3, 0) // apparently unordered-set

	case ACMP:
		return OPVCC(11, 0, 0, 0) | 1<<21 /* L=1 */
	case ACMPU:
		return OPVCC(10, 0, 0, 0) | 1<<21
	case ACMPW:
		return OPVCC(11, 0, 0, 0) /* L=0 */
	case ACMPWU:
		return OPVCC(10, 0, 0, 0)
	case ACMPEQB:
		return OPVCC(31, 224, 0, 0) /* cmpeqb - v3.00 */

	case ALSW:
		return OPVCC(31, 597, 0, 0)

	case ACOPY:
		return OPVCC(31, 774, 0, 0) /* copy - v3.00 */
	case APASTECC:
		return OPVCC(31, 902, 0, 1) /* paste. - v3.00 */
	case ADARN:
		return OPVCC(31, 755, 0, 0) /* darn - v3.00 */

	case AMULLW:
		return OPVCC(7, 0, 0, 0)

	case AOR:
		return OPVCC(24, 0, 0, 0)
	case AORIS:
		return OPVCC(25, 0, 0, 0) /* ORIS */

	case ARLWMI:
		return OPVCC(20, 0, 0, 0) /* rlwimi */
	case ARLWMICC:
		return OPVCC(20, 0, 0, 1)
	case ARLDMI:
		return OPVCC(30, 0, 0, 0) | 3<<2 /* rldimi */
	case ARLDMICC:
		return OPVCC(30, 0, 0, 1) | 3<<2
	case ARLDIMI:
		return OPVCC(30, 0, 0, 0) | 3<<2 /* rldimi */
	case ARLDIMICC:
		return OPVCC(30, 0, 0, 1) | 3<<2
	case ARLWNM:
		return OPVCC(21, 0, 0, 0) /* rlwinm */
	case ARLWNMCC:
		return OPVCC(21, 0, 0, 1)

	case ARLDCL:
		return OPVCC(30, 0, 0, 0) /* rldicl */
	case ARLDCLCC:
		return OPVCC(30, 0, 0, 1)
	case ARLDCR:
		return OPVCC(30, 1, 0, 0) /* rldicr */
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

	/* Vector (VMX/Altivec) instructions */
	/* ISA 2.03 enables these for PPC970. For POWERx processors, these */
	/* are enabled starting at POWER6 (ISA 2.05). */
	case AVSPLTB:
		return OPVX(4, 524, 0, 0) /* vspltb - v2.03 */
	case AVSPLTH:
		return OPVX(4, 588, 0, 0) /* vsplth - v2.03 */
	case AVSPLTW:
		return OPVX(4, 652, 0, 0) /* vspltw - v2.03 */

	case AVSPLTISB:
		return OPVX(4, 780, 0, 0) /* vspltisb - v2.03 */
	case AVSPLTISH:
		return OPVX(4, 844, 0, 0) /* vspltish - v2.03 */
	case AVSPLTISW:
		return OPVX(4, 908, 0, 0) /* vspltisw - v2.03 */
	/* End of vector instructions */

	case AFTDIV:
		return OPVCC(63, 128, 0, 0) /* ftdiv - v2.06 */
	case AFTSQRT:
		return OPVCC(63, 160, 0, 0) /* ftsqrt - v2.06 */

	case AXOR:
		return OPVCC(26, 0, 0, 0) /* XORIL */
	case AXORIS:
		return OPVCC(27, 0, 0, 0) /* XORIS */
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
		return OPVCC(58, 0, 0, 0) /* ld */
	case AMOVDU:
		return OPVCC(58, 0, 0, 1) /* ldu */
	case AMOVWZ:
		return OPVCC(32, 0, 0, 0) /* lwz */
	case AMOVWZU:
		return OPVCC(33, 0, 0, 0) /* lwzu */
	case AMOVW:
		return OPVCC(58, 0, 0, 0) | 1<<1 /* lwa */

	/* no AMOVWU */
	case AMOVB, AMOVBZ:
		return OPVCC(34, 0, 0, 0)
	/* load */

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
		return OPVCC(46, 0, 0, 0) /* lmw */
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
		return OPVCC(31, 23, 0, 0) /* lwzx */
	case AMOVWZU:
		return OPVCC(31, 55, 0, 0) /* lwzux */
	case AMOVW:
		return OPVCC(31, 341, 0, 0) /* lwax */
	case AMOVWU:
		return OPVCC(31, 373, 0, 0) /* lwaux */

	case AMOVB, AMOVBZ:
		return OPVCC(31, 87, 0, 0) /* lbzx */

	case AMOVBU, AMOVBZU:
		return OPVCC(31, 119, 0, 0) /* lbzux */
	case AFMOVD:
		return OPVCC(31, 599, 0, 0) /* lfdx */
	case AFMOVDU:
		return OPVCC(31, 631, 0, 0) /*  lfdux */
	case AFMOVS:
		return OPVCC(31, 535, 0, 0) /* lfsx */
	case AFMOVSU:
		return OPVCC(31, 567, 0, 0) /* lfsux */
	case AFMOVSX:
		return OPVCC(31, 855, 0, 0) /* lfiwax - power6, isa 2.05 */
	case AFMOVSZ:
		return OPVCC(31, 887, 0, 0) /* lfiwzx - power7, isa 2.06 */
	case AMOVH:
		return OPVCC(31, 343, 0, 0) /* lhax */
	case AMOVHU:
		return OPVCC(31, 375, 0, 0) /* lhaux */
	case AMOVHBR:
		return OPVCC(31, 790, 0, 0) /* lhbrx */
	case AMOVWBR:
		return OPVCC(31, 534, 0, 0) /* lwbrx */
	case AMOVDBR:
		return OPVCC(31, 532, 0, 0) /* ldbrx */
	case AMOVHZ:
		return OPVCC(31, 279, 0, 0) /* lhzx */
	case AMOVHZU:
		return OPVCC(31, 311, 0, 0) /* lhzux */
	case AECIWX:
		return OPVCC(31, 310, 0, 0) /* eciwx */
	case ALBAR:
		return OPVCC(31, 52, 0, 0) /* lbarx */
	case ALHAR:
		return OPVCC(31, 116, 0, 0) /* lharx */
	case ALWAR:
		return OPVCC(31, 20, 0, 0) /* lwarx */
	case ALDAR:
		return OPVCC(31, 84, 0, 0) /* ldarx */
	case ALSW:
		return OPVCC(31, 533, 0, 0) /* lswx */
	case AMOVD:
		return OPVCC(31, 21, 0, 0) /* ldx */
	case AMOVDU:
		return OPVCC(31, 53, 0, 0) /* ldux */
	case ALDMX:
		return OPVCC(31, 309, 0, 0) /* ldmx */

	/* Vector (VMX/Altivec) instructions */
	/* ISA 2.03 enables these for PPC970. For POWERx processors, these */
	/* are enabled starting at POWER6 (ISA 2.05). */
	case ALVEBX:
		return OPVCC(31, 7, 0, 0) /* lvebx - v2.03 */
	case ALVEHX:
		return OPVCC(31, 39, 0, 0) /* lvehx - v2.03 */
	case ALVEWX:
		return OPVCC(31, 71, 0, 0) /* lvewx - v2.03 */
	case ALVX:
		return OPVCC(31, 103, 0, 0) /* lvx - v2.03 */
	case ALVXL:
		return OPVCC(31, 359, 0, 0) /* lvxl - v2.03 */
	case ALVSL:
		return OPVCC(31, 6, 0, 0) /* lvsl - v2.03 */
	case ALVSR:
		return OPVCC(31, 38, 0, 0) /* lvsr - v2.03 */
	/* End of vector instructions */

	/* Vector scalar (VSX) instructions */
	/* ISA 2.06 enables these for POWER7. */
	case ALXVD2X:
		return OPVXX1(31, 844, 0) /* lxvd2x - v2.06 */
	case ALXVDSX:
		return OPVXX1(31, 332, 0) /* lxvdsx - v2.06 */
	case ALXVW4X:
		return OPVXX1(31, 780, 0) /* lxvw4x - v2.06 */

	case ALXSDX:
		return OPVXX1(31, 588, 0) /* lxsdx - v2.06 */

	case ALXSIWAX:
		return OPVXX1(31, 76, 0) /* lxsiwax - v2.07 */
	case ALXSIWZX:
		return OPVXX1(31, 12, 0) /* lxsiwzx - v2.07 */
		/* End of vector scalar instructions */

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
		return OPVCC(38, 0, 0, 0) /* stb */

	case AMOVBU, AMOVBZU:
		return OPVCC(39, 0, 0, 0) /* stbu */
	case AFMOVD:
		return OPVCC(54, 0, 0, 0) /* stfd */
	case AFMOVDU:
		return OPVCC(55, 0, 0, 0) /* stfdu */
	case AFMOVS:
		return OPVCC(52, 0, 0, 0) /* stfs */
	case AFMOVSU:
		return OPVCC(53, 0, 0, 0) /* stfsu */

	case AMOVHZ, AMOVH:
		return OPVCC(44, 0, 0, 0) /* sth */

	case AMOVHZU, AMOVHU:
		return OPVCC(45, 0, 0, 0) /* sthu */
	case AMOVMW:
		return OPVCC(47, 0, 0, 0) /* stmw */
	case ASTSW:
		return OPVCC(31, 725, 0, 0) /* stswi */

	case AMOVWZ, AMOVW:
		return OPVCC(36, 0, 0, 0) /* stw */

	case AMOVWZU, AMOVWU:
		return OPVCC(37, 0, 0, 0) /* stwu */
	case AMOVD:
		return OPVCC(62, 0, 0, 0) /* std */
	case AMOVDU:
		return OPVCC(62, 0, 0, 1) /* stdu */
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
		return OPVCC(31, 215, 0, 0) /* stbx */

	case AMOVBU, AMOVBZU:
		return OPVCC(31, 247, 0, 0) /* stbux */
	case AFMOVD:
		return OPVCC(31, 727, 0, 0) /* stfdx */
	case AFMOVDU:
		return OPVCC(31, 759, 0, 0) /* stfdux */
	case AFMOVS:
		return OPVCC(31, 663, 0, 0) /* stfsx */
	case AFMOVSU:
		return OPVCC(31, 695, 0, 0) /* stfsux */
	case AFMOVSX:
		return OPVCC(31, 983, 0, 0) /* stfiwx */

	case AMOVHZ, AMOVH:
		return OPVCC(31, 407, 0, 0) /* sthx */
	case AMOVHBR:
		return OPVCC(31, 918, 0, 0) /* sthbrx */

	case AMOVHZU, AMOVHU:
		return OPVCC(31, 439, 0, 0) /* sthux */

	case AMOVWZ, AMOVW:
		return OPVCC(31, 151, 0, 0) /* stwx */

	case AMOVWZU, AMOVWU:
		return OPVCC(31, 183, 0, 0) /* stwux */
	case ASTSW:
		return OPVCC(31, 661, 0, 0) /* stswx */
	case AMOVWBR:
		return OPVCC(31, 662, 0, 0) /* stwbrx */
	case AMOVDBR:
		return OPVCC(31, 660, 0, 0) /* stdbrx */
	case ASTBCCC:
		return OPVCC(31, 694, 0, 1) /* stbcx. */
	case ASTWCCC:
		return OPVCC(31, 150, 0, 1) /* stwcx. */
	case ASTDCCC:
		return OPVCC(31, 214, 0, 1) /* stwdx. */
	case AECOWX:
		return OPVCC(31, 438, 0, 0) /* ecowx */
	case AMOVD:
		return OPVCC(31, 149, 0, 0) /* stdx */
	case AMOVDU:
		return OPVCC(31, 181, 0, 0) /* stdux */

	/* Vector (VMX/Altivec) instructions */
	/* ISA 2.03 enables these for PPC970. For POWERx processors, these */
	/* are enabled starting at POWER6 (ISA 2.05). */
	case ASTVEBX:
		return OPVCC(31, 135, 0, 0) /* stvebx - v2.03 */
	case ASTVEHX:
		return OPVCC(31, 167, 0, 0) /* stvehx - v2.03 */
	case ASTVEWX:
		return OPVCC(31, 199, 0, 0) /* stvewx - v2.03 */
	case ASTVX:
		return OPVCC(31, 231, 0, 0) /* stvx - v2.03 */
	case ASTVXL:
		return OPVCC(31, 487, 0, 0) /* stvxl - v2.03 */
	/* End of vector instructions */

	/* Vector scalar (VSX) instructions */
	/* ISA 2.06 enables these for POWER7. */
	case ASTXVD2X:
		return OPVXX1(31, 972, 0) /* stxvd2x - v2.06 */
	case ASTXVW4X:
		return OPVXX1(31, 908, 0) /* stxvw4x - v2.06 */

	case ASTXSDX:
		return OPVXX1(31, 716, 0) /* stxsdx - v2.06 */

	case ASTXSIWX:
		return OPVXX1(31, 140, 0) /* stxsiwx - v2.07 */
		/* End of vector scalar instructions */

	}

	c.ctxt.Diag("unknown storex opcode %v", a)
	return 0
}

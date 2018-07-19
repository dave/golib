package arm64

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"log"
	"math"
	"sort"
)

// ctxt7 holds state while assembling a single function.
// Each function gets a fresh ctxt7.
// This allows for multiple functions to be safely concurrently assembled.
type ctxt7 struct {
	ctxt       *obj.Link
	newprog    obj.ProgAlloc
	cursym     *obj.LSym
	blitrl     *obj.Prog
	elitrl     *obj.Prog
	autosize   int32
	instoffset int64
	pc         int64
	pool       struct {
		start uint32
		size  uint32
	}
}

const (
	funcAlign = 16
)

const (
	REGFROM = 1
)

type Optab struct {
	as    obj.As
	a1    uint8
	a2    uint8
	a3    uint8
	a4    uint8
	type_ int8
	size  int8
	param int16
	flag  int8
	scond uint16
}

const (
	S32     = 0 << 31
	S64     = 1 << 31
	Sbit    = 1 << 29
	LSL0_32 = 2 << 13
	LSL0_64 = 3 << 13
)

func OPDP2(x uint32) uint32 {
	return 0<<30 | 0<<29 | 0xd6<<21 | x<<10
}

func OPDP3(sf uint32, op54 uint32, op31 uint32, o0 uint32) uint32 {
	return sf<<31 | op54<<29 | 0x1B<<24 | op31<<21 | o0<<15
}

func OPBcc(x uint32) uint32 {
	return 0x2A<<25 | 0<<24 | 0<<4 | x&15
}

func OPBLR(x uint32) uint32 {

	return 0x6B<<25 | 0<<23 | x<<21 | 0x1F<<16 | 0<<10
}

func SYSOP(l uint32, op0 uint32, op1 uint32, crn uint32, crm uint32, op2 uint32, rt uint32) uint32 {
	return 0x354<<22 | l<<21 | op0<<19 | op1<<16 | crn&15<<12 | crm&15<<8 | op2<<5 | rt
}

func SYSHINT(x uint32) uint32 {
	return SYSOP(0, 0, 3, 2, 0, x, 0x1F)
}

func LDSTR12U(sz uint32, v uint32, opc uint32) uint32 {
	return sz<<30 | 7<<27 | v<<26 | 1<<24 | opc<<22
}

func LDSTR9S(sz uint32, v uint32, opc uint32) uint32 {
	return sz<<30 | 7<<27 | v<<26 | 0<<24 | opc<<22
}

func LD2STR(o uint32) uint32 {
	return o &^ (3 << 22)
}

func LDSTX(sz uint32, o2 uint32, l uint32, o1 uint32, o0 uint32) uint32 {
	return sz<<30 | 0x8<<24 | o2<<23 | l<<22 | o1<<21 | o0<<15
}

func FPCMP(m uint32, s uint32, type_ uint32, op uint32, op2 uint32) uint32 {
	return m<<31 | s<<29 | 0x1E<<24 | type_<<22 | 1<<21 | op<<14 | 8<<10 | op2
}

func FPCCMP(m uint32, s uint32, type_ uint32, op uint32) uint32 {
	return m<<31 | s<<29 | 0x1E<<24 | type_<<22 | 1<<21 | 1<<10 | op<<4
}

func FPOP1S(m uint32, s uint32, type_ uint32, op uint32) uint32 {
	return m<<31 | s<<29 | 0x1E<<24 | type_<<22 | 1<<21 | op<<15 | 0x10<<10
}

func FPOP2S(m uint32, s uint32, type_ uint32, op uint32) uint32 {
	return m<<31 | s<<29 | 0x1E<<24 | type_<<22 | 1<<21 | op<<12 | 2<<10
}

func FPOP3S(m uint32, s uint32, type_ uint32, op uint32, op2 uint32) uint32 {
	return m<<31 | s<<29 | 0x1F<<24 | type_<<22 | op<<21 | op2<<15
}

func FPCVTI(sf uint32, s uint32, type_ uint32, rmode uint32, op uint32) uint32 {
	return sf<<31 | s<<29 | 0x1E<<24 | type_<<22 | 1<<21 | rmode<<19 | op<<16 | 0<<10
}

func ADR(p uint32, o uint32, rt uint32) uint32 {
	return p<<31 | (o&3)<<29 | 0x10<<24 | ((o>>2)&0x7FFFF)<<5 | rt&31
}

func OPBIT(x uint32) uint32 {
	return 1<<30 | 0<<29 | 0xD6<<21 | 0<<16 | x<<10
}

const (
	LFROM = 1 << 0
	LTO   = 1 << 1
)

/*
 * valid pstate field values, and value to use in instruction
 */

// the System register values, and value to use in instruction

func (psess *PackageSession) span7(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	p := cursym.Func.Text
	if p == nil || p.Link == nil {
		return
	}

	if psess.oprange[AAND&obj.AMask] == nil {
		ctxt.Diag("arm64 ops not initialized, call arm64.buildop first")
	}

	c := ctxt7{ctxt: ctxt, newprog: newprog, cursym: cursym, autosize: int32(p.To.Offset&0xffffffff) + 8}

	bflag := 1
	pc := int64(0)
	p.Pc = pc
	var m int
	var o *Optab
	for p = p.Link; p != nil; p = p.Link {
		if p.As == ADWORD && (pc&7) != 0 {
			pc += 4
		}
		p.Pc = pc
		o = c.oplook(psess, p)
		m = int(o.size)
		if m == 0 {
			if p.As != obj.ANOP && p.As != obj.AFUNCDATA && p.As != obj.APCDATA {
				c.ctxt.Diag("zero-width instruction\n%v", p)
			}
			continue
		}

		switch o.flag & (LFROM | LTO) {
		case LFROM:
			c.addpool(psess, p, &p.From)

		case LTO:
			c.addpool(psess, p, &p.To)
			break
		}

		if p.As == AB || p.As == obj.ARET || p.As == AERET {
			c.checkpool(p, 0)
		}
		pc += int64(m)
		if c.blitrl != nil {
			c.checkpool(p, 1)
		}
	}

	c.cursym.Size = pc

	for bflag != 0 {
		bflag = 0
		pc = 0
		for p = c.cursym.Func.Text.Link; p != nil; p = p.Link {
			if p.As == ADWORD && (pc&7) != 0 {
				pc += 4
			}
			p.Pc = pc
			o = c.oplook(psess, p)

			if (o.type_ == 7 || o.type_ == 39 || o.type_ == 40) && p.Pcond != nil {
				otxt := p.Pcond.Pc - pc
				var toofar bool
				switch o.type_ {
				case 7, 39:
					toofar = otxt <= -(1<<20)+10 || otxt >= (1<<20)-10
				case 40:
					toofar = otxt <= -(1<<15)+10 || otxt >= (1<<15)-10
				}
				if toofar {
					q := c.newprog()
					q.Link = p.Link
					p.Link = q
					q.As = AB
					q.To.Type = obj.TYPE_BRANCH
					q.Pcond = p.Pcond
					p.Pcond = q
					q = c.newprog()
					q.Link = p.Link
					p.Link = q
					q.As = AB
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
	}

	pc += -pc & (funcAlign - 1)
	c.cursym.Size = pc

	c.cursym.Grow(c.cursym.Size)
	bp := c.cursym.P
	psz := int32(0)
	var i int
	var out [6]uint32
	for p := c.cursym.Func.Text.Link; p != nil; p = p.Link {
		c.pc = p.Pc
		o = c.oplook(psess, p)

		if o.as == ADWORD && psz%8 != 0 {
			bp[3] = 0
			bp[2] = bp[3]
			bp[1] = bp[2]
			bp[0] = bp[1]
			bp = bp[4:]
			psz += 4
		}

		if int(o.size) > 4*len(out) {
			log.Fatalf("out array in span7 is too small, need at least %d for %v", o.size/4, p)
		}
		c.asmout(psess, p, o, out[:])
		for i = 0; i < int(o.size/4); i++ {
			c.ctxt.Arch.ByteOrder.PutUint32(bp, out[i])
			bp = bp[4:]
			psz += 4
		}
	}
}

/*
 * when the first reference to the literal pool threatens
 * to go out of range of a 1Mb PC-relative offset
 * drop the pool now, and branch round it.
 */
func (c *ctxt7) checkpool(p *obj.Prog, skip int) {
	if c.pool.size >= 0xffff0 || !ispcdisp(int32(p.Pc+4+int64(c.pool.size)-int64(c.pool.start)+8)) {
		c.flushpool(p, skip)
	} else if p.Link == nil {
		c.flushpool(p, 2)
	}
}

func (c *ctxt7) flushpool(p *obj.Prog, skip int) {
	if c.blitrl != nil {
		if skip != 0 {
			if c.ctxt.Debugvlog && skip == 1 {
				fmt.Printf("note: flush literal pool at %#x: len=%d ref=%x\n", uint64(p.Pc+4), c.pool.size, c.pool.start)
			}
			q := c.newprog()
			q.As = AB
			q.To.Type = obj.TYPE_BRANCH
			q.Pcond = p.Link
			q.Link = c.blitrl
			q.Pos = p.Pos
			c.blitrl = q
		} else if p.Pc+int64(c.pool.size)-int64(c.pool.start) < maxPCDisp {
			return
		}

		for q := c.blitrl; q != nil; q = q.Link {
			q.Pos = p.Pos
		}

		c.elitrl.Link = p.Link
		p.Link = c.blitrl

		c.blitrl = nil
		c.elitrl = nil
		c.pool.size = 0
		c.pool.start = 0
	}
}

/*
 * MOVD foo(SB), R is actually
 *   MOVD addr, REGTMP
 *   MOVD REGTMP, R
 * where addr is the address of the DWORD containing the address of foo.
 *
 * TODO: hash
 */
func (c *ctxt7) addpool(psess *PackageSession, p *obj.Prog, a *obj.Addr) {
	cls := c.aclass(a)
	lit := c.instoffset
	t := c.newprog()
	t.As = AWORD
	sz := 4

	if a.Type == obj.TYPE_CONST {
		if lit != int64(int32(lit)) && uint64(lit) != uint64(uint32(lit)) {

			t.As = ADWORD
			sz = 8
		}
	} else if p.As == AMOVD && a.Type != obj.TYPE_MEM || cls == C_ADDR || cls == C_VCON || lit != int64(int32(lit)) || uint64(lit) != uint64(uint32(lit)) {

		t.As = ADWORD
		sz = 8
	}

	switch cls {

	default:
		if a.Name != obj.NAME_EXTERN {
			fmt.Printf("addpool: %v in %v shouldn't go to default case\n", psess.DRconv(cls), p)
		}

		t.To.Offset = a.Offset
		t.To.Sym = a.Sym
		t.To.Type = a.Type
		t.To.Name = a.Name

	case C_ADDCON:
		fallthrough

	case C_ZAUTO,
		C_PSAUTO,
		C_PSAUTO_8,
		C_PSAUTO_4,
		C_PPAUTO,
		C_UAUTO4K_8,
		C_UAUTO4K_4,
		C_UAUTO4K_2,
		C_UAUTO4K,
		C_UAUTO8K_8,
		C_UAUTO8K_4,
		C_UAUTO8K,
		C_UAUTO16K_8,
		C_UAUTO16K,
		C_UAUTO32K,
		C_NSAUTO_8,
		C_NSAUTO_4,
		C_NSAUTO,
		C_NPAUTO,
		C_NAUTO4K,
		C_LAUTO,
		C_PPOREG,
		C_PSOREG,
		C_PSOREG_4,
		C_PSOREG_8,
		C_UOREG4K_8,
		C_UOREG4K_4,
		C_UOREG4K_2,
		C_UOREG4K,
		C_UOREG8K_8,
		C_UOREG8K_4,
		C_UOREG8K,
		C_UOREG16K_8,
		C_UOREG16K,
		C_UOREG32K,
		C_NSOREG_8,
		C_NSOREG_4,
		C_NSOREG,
		C_NPOREG,
		C_NOREG4K,
		C_LOREG,
		C_LACON,
		C_LCON,
		C_VCON:
		if a.Name == obj.NAME_EXTERN {
			fmt.Printf("addpool: %v in %v needs reloc\n", psess.DRconv(cls), p)
		}

		t.To.Type = obj.TYPE_CONST
		t.To.Offset = lit
		break
	}

	for q := c.blitrl; q != nil; q = q.Link {
		if q.To == t.To {
			p.Pcond = q
			return
		}
	}

	q := c.newprog()
	*q = *t
	q.Pc = int64(c.pool.size)
	if c.blitrl == nil {
		c.blitrl = q
		c.pool.start = uint32(p.Pc)
	} else {
		c.elitrl.Link = q
	}
	c.elitrl = q
	c.pool.size = -c.pool.size & (funcAlign - 1)
	c.pool.size += uint32(sz)
	p.Pcond = q
}

func (c *ctxt7) regoff(a *obj.Addr) uint32 {
	c.instoffset = 0
	c.aclass(a)
	return uint32(c.instoffset)
}

func isRegShiftOrExt(a *obj.Addr) bool {
	return (a.Index-obj.RBaseARM64)&REG_EXT != 0 || (a.Index-obj.RBaseARM64)&REG_LSL != 0
}

// Maximum PC-relative displacement.
// The actual limit is ±2²⁰, but we are conservative
// to avoid needing to recompute the literal pool flush points
// as span-dependent jumps are enlarged.
const maxPCDisp = 512 * 1024

// ispcdisp reports whether v is a valid PC-relative displacement.
func ispcdisp(v int32) bool {
	return -maxPCDisp < v && v < maxPCDisp && v&3 == 0
}

func isaddcon(v int64) bool {

	if v < 0 {
		return false
	}
	if (v & 0xFFF) == 0 {
		v >>= 12
	}
	return v <= 0xFFF
}

// isbitcon returns whether a constant can be encoded into a logical instruction.
// bitcon has a binary form of repetition of a bit sequence of length 2, 4, 8, 16, 32, or 64,
// which itself is a rotate (w.r.t. the length of the unit) of a sequence of ones.
// special cases: 0 and -1 are not bitcon.
// this function needs to run against virtually all the constants, so it needs to be fast.
// for this reason, bitcon testing and bitcon encoding are separate functions.
func isbitcon(x uint64) bool {
	if x == 1<<64-1 || x == 0 {
		return false
	}

	switch {
	case x != x>>32|x<<32:

	case x != x>>16|x<<48:

		x = uint64(int64(int32(x)))
	case x != x>>8|x<<56:

		x = uint64(int64(int16(x)))
	case x != x>>4|x<<60:

		x = uint64(int64(int8(x)))
	default:

		return true
	}
	return sequenceOfOnes(x) || sequenceOfOnes(^x)
}

// sequenceOfOnes tests whether a constant is a sequence of ones in binary, with leading and trailing zeros
func sequenceOfOnes(x uint64) bool {
	y := x & -x
	y += x
	return (y-1)&y == 0
}

// bitconEncode returns the encoding of a bitcon used in logical instructions
// x is known to be a bitcon
// a bitcon is a sequence of n ones at low bits (i.e. 1<<n-1), right rotated
// by R bits, and repeated with period of 64, 32, 16, 8, 4, or 2.
// it is encoded in logical instructions with 3 bitfields
// N (1 bit) : R (6 bits) : S (6 bits), where
// N=1           -- period=64
// N=0, S=0xxxxx -- period=32
// N=0, S=10xxxx -- period=16
// N=0, S=110xxx -- period=8
// N=0, S=1110xx -- period=4
// N=0, S=11110x -- period=2
// R is the shift amount, low bits of S = n-1
func bitconEncode(x uint64, mode int) uint32 {
	var period uint32

	switch {
	case x != x>>32|x<<32:
		period = 64
	case x != x>>16|x<<48:
		period = 32
		x = uint64(int64(int32(x)))
	case x != x>>8|x<<56:
		period = 16
		x = uint64(int64(int16(x)))
	case x != x>>4|x<<60:
		period = 8
		x = uint64(int64(int8(x)))
	case x != x>>2|x<<62:
		period = 4
		x = uint64(int64(x<<60) >> 60)
	default:
		period = 2
		x = uint64(int64(x<<62) >> 62)
	}
	neg := false
	if int64(x) < 0 {
		x = ^x
		neg = true
	}
	y := x & -x
	s := log2(y)
	n := log2(x+y) - s
	if neg {

		s = n + s
		n = period - n
	}

	N := uint32(0)
	if mode == 64 && period == 64 {
		N = 1
	}
	R := (period - s) & (period - 1) & uint32(mode-1)
	S := (n - 1) | 63&^(period<<1-1)
	return N<<22 | R<<16 | S<<10
}

func log2(x uint64) uint32 {
	if x == 0 {
		panic("log2 of 0")
	}
	n := uint32(0)
	if x >= 1<<32 {
		x >>= 32
		n += 32
	}
	if x >= 1<<16 {
		x >>= 16
		n += 16
	}
	if x >= 1<<8 {
		x >>= 8
		n += 8
	}
	if x >= 1<<4 {
		x >>= 4
		n += 4
	}
	if x >= 1<<2 {
		x >>= 2
		n += 2
	}
	if x >= 1<<1 {
		x >>= 1
		n += 1
	}
	return n
}

func autoclass(l int64) int {
	if l == 0 {
		return C_ZAUTO
	}

	if l < 0 {
		if l >= -256 && (l&7) == 0 {
			return C_NSAUTO_8
		}
		if l >= -256 && (l&3) == 0 {
			return C_NSAUTO_4
		}
		if l >= -256 {
			return C_NSAUTO
		}
		if l >= -512 && (l&7) == 0 {
			return C_NPAUTO
		}
		if l >= -4095 {
			return C_NAUTO4K
		}
		return C_LAUTO
	}

	if l <= 255 {
		if (l & 7) == 0 {
			return C_PSAUTO_8
		}
		if (l & 3) == 0 {
			return C_PSAUTO_4
		}
		return C_PSAUTO
	}
	if l <= 504 && l&7 == 0 {
		return C_PPAUTO
	}
	if l <= 4095 {
		if l&7 == 0 {
			return C_UAUTO4K_8
		}
		if l&3 == 0 {
			return C_UAUTO4K_4
		}
		if l&1 == 0 {
			return C_UAUTO4K_2
		}
		return C_UAUTO4K
	}
	if l <= 8190 {
		if l&7 == 0 {
			return C_UAUTO8K_8
		}
		if l&3 == 0 {
			return C_UAUTO8K_4
		}
		if l&1 == 0 {
			return C_UAUTO8K
		}
	}
	if l <= 16380 {
		if l&7 == 0 {
			return C_UAUTO16K_8
		}
		if l&3 == 0 {
			return C_UAUTO16K
		}
	}
	if l <= 32760 && (l&7) == 0 {
		return C_UAUTO32K
	}
	return C_LAUTO
}

func oregclass(l int64) int {
	return autoclass(l) - C_ZAUTO + C_ZOREG
}

/*
 * given an offset v and a class c (see above)
 * return the offset value to use in the instruction,
 * scaled if necessary
 */
func (c *ctxt7) offsetshift(psess *PackageSession, p *obj.Prog, v int64, cls int) int64 {
	s := 0
	if cls >= C_SEXT1 && cls <= C_SEXT16 {
		s = cls - C_SEXT1
	} else {
		switch cls {
		case C_UAUTO4K, C_UOREG4K, C_ZOREG:
			s = 0
		case C_UAUTO8K, C_UOREG8K:
			s = 1
		case C_UAUTO16K, C_UOREG16K:
			s = 2
		case C_UAUTO32K, C_UOREG32K:
			s = 3
		default:
			c.ctxt.Diag("bad class: %v\n%v", psess.DRconv(cls), p)
		}
	}
	vs := v >> uint(s)
	if vs<<uint(s) != v {
		c.ctxt.Diag("odd offset: %d\n%v", v, p)
	}
	return vs
}

/*
 * if v contains a single 16-bit value aligned
 * on a 16-bit field, and thus suitable for movk/movn,
 * return the field index 0 to 3; otherwise return -1
 */
func movcon(v int64) int {
	for s := 0; s < 64; s += 16 {
		if (uint64(v) &^ (uint64(0xFFFF) << uint(s))) == 0 {
			return s / 16
		}
	}
	return -1
}

func rclass(r int16) int {
	switch {
	case REG_R0 <= r && r <= REG_R30:
		return C_REG
	case r == REGZERO:
		return C_ZCON
	case REG_F0 <= r && r <= REG_F31:
		return C_FREG
	case REG_V0 <= r && r <= REG_V31:
		return C_VREG
	case COND_EQ <= r && r <= COND_NV:
		return C_COND
	case r == REGSP:
		return C_RSP
	case r >= REG_ARNG && r < REG_ELEM:
		return C_ARNG
	case r >= REG_ELEM && r < REG_ELEM_END:
		return C_ELEM
	case r >= REG_UXTB && r < REG_SPECIAL:
		return C_EXTREG
	case r >= REG_SPECIAL:
		return C_SPR
	}
	return C_GOK
}

func (c *ctxt7) aclass(a *obj.Addr) int {
	switch a.Type {
	case obj.TYPE_NONE:
		return C_NONE

	case obj.TYPE_REG:
		return rclass(a.Reg)

	case obj.TYPE_REGREG:
		return C_PAIR

	case obj.TYPE_SHIFT:
		return C_SHIFT

	case obj.TYPE_REGLIST:
		return C_LIST

	case obj.TYPE_MEM:
		switch a.Name {
		case obj.NAME_EXTERN, obj.NAME_STATIC:
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
			if a.Reg == REGSP {

				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset
			return autoclass(c.instoffset)

		case obj.NAME_PARAM:
			if a.Reg == REGSP {

				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset + 8
			return autoclass(c.instoffset)

		case obj.NAME_NONE:
			if a.Index != 0 {
				if a.Offset != 0 {
					if isRegShiftOrExt(a) {

						return C_ROFF
					}
					return C_GOK
				}

				return C_ROFF
			}
			c.instoffset = a.Offset
			return oregclass(c.instoffset)
		}
		return C_GOK

	case obj.TYPE_FCONST:
		return C_FCON

	case obj.TYPE_TEXTSIZE:
		return C_TEXTSIZE

	case obj.TYPE_CONST, obj.TYPE_ADDR:
		switch a.Name {
		case obj.NAME_NONE:
			c.instoffset = a.Offset
			if a.Reg != 0 && a.Reg != REGZERO {
				break
			}
			v := c.instoffset
			if v == 0 {
				return C_ZCON
			}
			if isaddcon(v) {
				if v <= 0xFFF {
					if isbitcon(uint64(v)) {
						return C_ABCON0
					}
					return C_ADDCON0
				}
				if isbitcon(uint64(v)) {
					return C_ABCON
				}
				return C_ADDCON
			}

			t := movcon(v)
			if t >= 0 {
				if isbitcon(uint64(v)) {
					return C_MBCON
				}
				return C_MOVCON
			}

			t = movcon(^v)
			if t >= 0 {
				if isbitcon(uint64(v)) {
					return C_MBCON
				}
				return C_MOVCON
			}

			if isbitcon(uint64(v)) {
				return C_BITCON
			}

			if uint64(v) == uint64(uint32(v)) || v == int64(int32(v)) {
				return C_LCON
			}
			return C_VCON

		case obj.NAME_EXTERN, obj.NAME_STATIC:
			if a.Sym == nil {
				return C_GOK
			}
			if a.Sym.Type == objabi.STLSBSS {
				c.ctxt.Diag("taking address of TLS variable is not supported")
			}
			c.instoffset = a.Offset
			return C_VCONADDR

		case obj.NAME_AUTO:
			if a.Reg == REGSP {

				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset

		case obj.NAME_PARAM:
			if a.Reg == REGSP {

				a.Reg = obj.REG_NONE
			}
			c.instoffset = int64(c.autosize) + a.Offset + 8
		default:
			return C_GOK
		}

		if isaddcon(c.instoffset) {
			return C_AACON
		}
		return C_LACON

	case obj.TYPE_BRANCH:
		return C_SBRA
	}

	return C_GOK
}

func oclass(a *obj.Addr) int {
	return int(a.Class) - 1
}

func (c *ctxt7) oplook(psess *PackageSession, p *obj.Prog) *Optab {
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
		a2 = rclass(p.Reg)
	}

	if false {
		fmt.Printf("oplook %v %d %d %d %d\n", p.As, a1, a2, a3, a4)
		fmt.Printf("\t\t%d %d\n", p.From.Type, p.To.Type)
	}

	ops := psess.oprange[p.As&obj.AMask]
	c1 := &psess.xcmp[a1]
	c2 := &psess.xcmp[a2]
	c3 := &psess.xcmp[a3]
	c4 := &psess.xcmp[a4]
	c5 := &psess.xcmp[p.Scond>>5]
	for i := range ops {
		op := &ops[i]
		if (int(op.a2) == a2 || c2[op.a2]) && c5[op.scond>>5] && c1[op.a1] && c3[op.a3] && c4[op.a4] {
			p.Optab = uint16(cap(psess.optab) - cap(ops) + i + 1)
			return op
		}
	}

	c.ctxt.Diag("illegal combination: %v %v %v %v %v, %d %d", p, psess.DRconv(a1), psess.DRconv(a2), psess.DRconv(a3), psess.DRconv(a4), p.From.Type, p.To.Type)

	return &Optab{obj.AUNDEF, C_NONE, C_NONE, C_NONE, C_NONE, 90, 4, 0, 0, 0}
}

func cmp(a int, b int) bool {
	if a == b {
		return true
	}
	switch a {
	case C_RSP:
		if b == C_REG {
			return true
		}

	case C_REG:
		if b == C_ZCON {
			return true
		}

	case C_ADDCON0:
		if b == C_ZCON || b == C_ABCON0 {
			return true
		}

	case C_ADDCON:
		if b == C_ZCON || b == C_ABCON0 || b == C_ADDCON0 || b == C_ABCON {
			return true
		}

	case C_BITCON:
		if b == C_ABCON0 || b == C_ABCON || b == C_MBCON {
			return true
		}

	case C_MOVCON:
		if b == C_MBCON || b == C_ZCON || b == C_ADDCON0 {
			return true
		}

	case C_LCON:
		if b == C_ZCON || b == C_BITCON || b == C_ADDCON || b == C_ADDCON0 || b == C_ABCON || b == C_ABCON0 || b == C_MBCON || b == C_MOVCON {
			return true
		}

	case C_VCON:
		return cmp(C_LCON, b)

	case C_LACON:
		if b == C_AACON {
			return true
		}

	case C_SEXT2:
		if b == C_SEXT1 {
			return true
		}

	case C_SEXT4:
		if b == C_SEXT1 || b == C_SEXT2 {
			return true
		}

	case C_SEXT8:
		if b >= C_SEXT1 && b <= C_SEXT4 {
			return true
		}

	case C_SEXT16:
		if b >= C_SEXT1 && b <= C_SEXT8 {
			return true
		}

	case C_LEXT:
		if b >= C_SEXT1 && b <= C_SEXT16 {
			return true
		}

	case C_NSAUTO_4:
		if b == C_NSAUTO_8 {
			return true
		}

	case C_NSAUTO:
		switch b {
		case C_NSAUTO_4, C_NSAUTO_8:
			return true
		}

	case C_NPAUTO:
		switch b {
		case C_NSAUTO_8:
			return true
		}

	case C_NAUTO4K:
		switch b {
		case C_NSAUTO_8, C_NSAUTO_4, C_NSAUTO, C_NPAUTO:
			return true
		}

	case C_PSAUTO_8:
		if b == C_ZAUTO {
			return true
		}

	case C_PSAUTO_4:
		switch b {
		case C_ZAUTO, C_PSAUTO_8:
			return true
		}

	case C_PSAUTO:
		switch b {
		case C_ZAUTO, C_PSAUTO_8, C_PSAUTO_4:
			return true
		}

	case C_PPAUTO:
		switch b {
		case C_ZAUTO, C_PSAUTO_8:
			return true
		}

	case C_UAUTO4K:
		switch b {
		case C_ZAUTO, C_PSAUTO, C_PSAUTO_4, C_PSAUTO_8,
			C_PPAUTO, C_UAUTO4K_2, C_UAUTO4K_4, C_UAUTO4K_8:
			return true
		}

	case C_UAUTO8K:
		switch b {
		case C_ZAUTO, C_PSAUTO, C_PSAUTO_4, C_PSAUTO_8, C_PPAUTO,
			C_UAUTO4K_2, C_UAUTO4K_4, C_UAUTO4K_8, C_UAUTO8K_4, C_UAUTO8K_8:
			return true
		}

	case C_UAUTO16K:
		switch b {
		case C_ZAUTO, C_PSAUTO, C_PSAUTO_4, C_PSAUTO_8, C_PPAUTO,
			C_UAUTO4K_4, C_UAUTO4K_8, C_UAUTO8K_4, C_UAUTO8K_8, C_UAUTO16K_8:
			return true
		}

	case C_UAUTO32K:
		switch b {
		case C_ZAUTO, C_PSAUTO, C_PSAUTO_4, C_PSAUTO_8,
			C_PPAUTO, C_UAUTO4K_8, C_UAUTO8K_8, C_UAUTO16K_8:
			return true
		}

	case C_LAUTO:
		switch b {
		case C_ZAUTO, C_NSAUTO, C_NSAUTO_4, C_NSAUTO_8, C_NPAUTO,
			C_NAUTO4K, C_PSAUTO, C_PSAUTO_4, C_PSAUTO_8, C_PPAUTO,
			C_UAUTO4K, C_UAUTO4K_2, C_UAUTO4K_4, C_UAUTO4K_8,
			C_UAUTO8K, C_UAUTO8K_4, C_UAUTO8K_8,
			C_UAUTO16K, C_UAUTO16K_8,
			C_UAUTO32K:
			return true
		}

	case C_NSOREG_4:
		if b == C_NSOREG_8 {
			return true
		}

	case C_NSOREG:
		switch b {
		case C_NSOREG_4, C_NSOREG_8:
			return true
		}

	case C_NPOREG:
		switch b {
		case C_NSOREG_8:
			return true
		}

	case C_NOREG4K:
		switch b {
		case C_NSOREG_8, C_NSOREG_4, C_NSOREG, C_NPOREG:
			return true
		}

	case C_PSOREG_4:
		switch b {
		case C_ZOREG, C_PSOREG_8:
			return true
		}

	case C_PSOREG:
		switch b {
		case C_ZOREG, C_PSOREG_8, C_PSOREG_4:
			return true
		}

	case C_PPOREG:
		switch b {
		case C_ZOREG, C_PSOREG_8:
			return true
		}

	case C_UOREG4K:
		switch b {
		case C_ZOREG, C_PSOREG_4, C_PSOREG_8, C_PSOREG,
			C_PPOREG, C_UOREG4K_2, C_UOREG4K_4, C_UOREG4K_8:
			return true
		}

	case C_UOREG8K:
		switch b {
		case C_ZOREG, C_PSOREG_4, C_PSOREG_8, C_PSOREG,
			C_PPOREG, C_UOREG4K_2, C_UOREG4K_4, C_UOREG4K_8,
			C_UOREG8K_4, C_UOREG8K_8:
			return true
		}

	case C_UOREG16K:
		switch b {
		case C_ZOREG, C_PSOREG_4, C_PSOREG_8, C_PSOREG,
			C_PPOREG, C_UOREG4K_4, C_UOREG4K_8, C_UOREG8K_4,
			C_UOREG8K_8, C_UOREG16K_8:
			return true
		}

	case C_UOREG32K:
		switch b {
		case C_ZOREG, C_PSOREG_4, C_PSOREG_8, C_PSOREG,
			C_PPOREG, C_UOREG4K_8, C_UOREG8K_8, C_UOREG16K_8:
			return true
		}

	case C_LOREG:
		switch b {
		case C_ZOREG, C_NSOREG, C_NSOREG_4, C_NSOREG_8, C_NPOREG,
			C_NOREG4K, C_PSOREG_4, C_PSOREG_8, C_PSOREG, C_PPOREG,
			C_UOREG4K, C_UOREG4K_2, C_UOREG4K_4, C_UOREG4K_8,
			C_UOREG8K, C_UOREG8K_4, C_UOREG8K_8,
			C_UOREG16K, C_UOREG16K_8,
			C_UOREG32K:
			return true
		}

	case C_LBRA:
		if b == C_SBRA {
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
	if p1.as != p2.as {
		return p1.as < p2.as
	}
	if p1.a1 != p2.a1 {
		return p1.a1 < p2.a1
	}
	if p1.a2 != p2.a2 {
		return p1.a2 < p2.a2
	}
	if p1.a3 != p2.a3 {
		return p1.a3 < p2.a3
	}
	if p1.a4 != p2.a4 {
		return p1.a4 < p2.a4
	}
	if p1.scond != p2.scond {
		return p1.scond < p2.scond
	}
	return false
}

func (psess *PackageSession) oprangeset(a obj.As, t []Optab) {
	psess.
		oprange[a&obj.AMask] = t
}

func (psess *PackageSession) buildop(ctxt *obj.Link) {
	if psess.oprange[AAND&obj.AMask] != nil {

		return
	}

	var n int
	for i := 0; i < C_GOK; i++ {
		for n = 0; n < C_GOK; n++ {
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
		start := i
		for psess.optab[i].as == r {
			i++
		}
		t := psess.optab[start:i]
		i--
		psess.
			oprangeset(r, t)
		switch r {
		default:
			ctxt.Diag("unknown op in build: %v", r)
			ctxt.DiagFlush()
			log.Fatalf("bad code")

		case AADD:
			psess.
				oprangeset(AADDS, t)
			psess.
				oprangeset(ASUB, t)
			psess.
				oprangeset(ASUBS, t)
			psess.
				oprangeset(AADDW, t)
			psess.
				oprangeset(AADDSW, t)
			psess.
				oprangeset(ASUBW, t)
			psess.
				oprangeset(ASUBSW, t)

		case AAND:
			psess.
				oprangeset(AANDW, t)
			psess.
				oprangeset(AEOR, t)
			psess.
				oprangeset(AEORW, t)
			psess.
				oprangeset(AORR, t)
			psess.
				oprangeset(AORRW, t)
			psess.
				oprangeset(ABIC, t)
			psess.
				oprangeset(ABICW, t)
			psess.
				oprangeset(AEON, t)
			psess.
				oprangeset(AEONW, t)
			psess.
				oprangeset(AORN, t)
			psess.
				oprangeset(AORNW, t)

		case AANDS:
			psess.
				oprangeset(AANDSW, t)
			psess.
				oprangeset(ABICS, t)
			psess.
				oprangeset(ABICSW, t)

		case ANEG:
			psess.
				oprangeset(ANEGS, t)
			psess.
				oprangeset(ANEGSW, t)
			psess.
				oprangeset(ANEGW, t)

		case AADC:
			psess.
				oprangeset(AADCW, t)
			psess.
				oprangeset(AADCS, t)
			psess.
				oprangeset(AADCSW, t)
			psess.
				oprangeset(ASBC, t)
			psess.
				oprangeset(ASBCW, t)
			psess.
				oprangeset(ASBCS, t)
			psess.
				oprangeset(ASBCSW, t)

		case ANGC:
			psess.
				oprangeset(ANGCW, t)
			psess.
				oprangeset(ANGCS, t)
			psess.
				oprangeset(ANGCSW, t)

		case ACMP:
			psess.
				oprangeset(ACMPW, t)
			psess.
				oprangeset(ACMN, t)
			psess.
				oprangeset(ACMNW, t)

		case ATST:
			psess.
				oprangeset(ATSTW, t)

		case AMVN:
			psess.
				oprangeset(AMVNW, t)

		case AMOVK:
			psess.
				oprangeset(AMOVKW, t)
			psess.
				oprangeset(AMOVN, t)
			psess.
				oprangeset(AMOVNW, t)
			psess.
				oprangeset(AMOVZ, t)
			psess.
				oprangeset(AMOVZW, t)

		case ASWPD:
			psess.
				oprangeset(ASWPB, t)
			psess.
				oprangeset(ASWPH, t)
			psess.
				oprangeset(ASWPW, t)
			psess.
				oprangeset(ALDADDALD, t)
			psess.
				oprangeset(ALDADDALW, t)
			psess.
				oprangeset(ALDADDB, t)
			psess.
				oprangeset(ALDADDH, t)
			psess.
				oprangeset(ALDADDW, t)
			psess.
				oprangeset(ALDADDD, t)
			psess.
				oprangeset(ALDANDB, t)
			psess.
				oprangeset(ALDANDH, t)
			psess.
				oprangeset(ALDANDW, t)
			psess.
				oprangeset(ALDANDD, t)
			psess.
				oprangeset(ALDEORB, t)
			psess.
				oprangeset(ALDEORH, t)
			psess.
				oprangeset(ALDEORW, t)
			psess.
				oprangeset(ALDEORD, t)
			psess.
				oprangeset(ALDORB, t)
			psess.
				oprangeset(ALDORH, t)
			psess.
				oprangeset(ALDORW, t)
			psess.
				oprangeset(ALDORD, t)

		case ABEQ:
			psess.
				oprangeset(ABNE, t)
			psess.
				oprangeset(ABCS, t)
			psess.
				oprangeset(ABHS, t)
			psess.
				oprangeset(ABCC, t)
			psess.
				oprangeset(ABLO, t)
			psess.
				oprangeset(ABMI, t)
			psess.
				oprangeset(ABPL, t)
			psess.
				oprangeset(ABVS, t)
			psess.
				oprangeset(ABVC, t)
			psess.
				oprangeset(ABHI, t)
			psess.
				oprangeset(ABLS, t)
			psess.
				oprangeset(ABGE, t)
			psess.
				oprangeset(ABLT, t)
			psess.
				oprangeset(ABGT, t)
			psess.
				oprangeset(ABLE, t)

		case ALSL:
			psess.
				oprangeset(ALSLW, t)
			psess.
				oprangeset(ALSR, t)
			psess.
				oprangeset(ALSRW, t)
			psess.
				oprangeset(AASR, t)
			psess.
				oprangeset(AASRW, t)
			psess.
				oprangeset(AROR, t)
			psess.
				oprangeset(ARORW, t)

		case ACLS:
			psess.
				oprangeset(ACLSW, t)
			psess.
				oprangeset(ACLZ, t)
			psess.
				oprangeset(ACLZW, t)
			psess.
				oprangeset(ARBIT, t)
			psess.
				oprangeset(ARBITW, t)
			psess.
				oprangeset(AREV, t)
			psess.
				oprangeset(AREVW, t)
			psess.
				oprangeset(AREV16, t)
			psess.
				oprangeset(AREV16W, t)
			psess.
				oprangeset(AREV32, t)

		case ASDIV:
			psess.
				oprangeset(ASDIVW, t)
			psess.
				oprangeset(AUDIV, t)
			psess.
				oprangeset(AUDIVW, t)
			psess.
				oprangeset(ACRC32B, t)
			psess.
				oprangeset(ACRC32CB, t)
			psess.
				oprangeset(ACRC32CH, t)
			psess.
				oprangeset(ACRC32CW, t)
			psess.
				oprangeset(ACRC32CX, t)
			psess.
				oprangeset(ACRC32H, t)
			psess.
				oprangeset(ACRC32W, t)
			psess.
				oprangeset(ACRC32X, t)

		case AMADD:
			psess.
				oprangeset(AMADDW, t)
			psess.
				oprangeset(AMSUB, t)
			psess.
				oprangeset(AMSUBW, t)
			psess.
				oprangeset(ASMADDL, t)
			psess.
				oprangeset(ASMSUBL, t)
			psess.
				oprangeset(AUMADDL, t)
			psess.
				oprangeset(AUMSUBL, t)

		case AREM:
			psess.
				oprangeset(AREMW, t)
			psess.
				oprangeset(AUREM, t)
			psess.
				oprangeset(AUREMW, t)

		case AMUL:
			psess.
				oprangeset(AMULW, t)
			psess.
				oprangeset(AMNEG, t)
			psess.
				oprangeset(AMNEGW, t)
			psess.
				oprangeset(ASMNEGL, t)
			psess.
				oprangeset(ASMULL, t)
			psess.
				oprangeset(ASMULH, t)
			psess.
				oprangeset(AUMNEGL, t)
			psess.
				oprangeset(AUMULH, t)
			psess.
				oprangeset(AUMULL, t)

		case AMOVB:
			psess.
				oprangeset(AMOVBU, t)

		case AMOVH:
			psess.
				oprangeset(AMOVHU, t)

		case AMOVW:
			psess.
				oprangeset(AMOVWU, t)

		case ABFM:
			psess.
				oprangeset(ABFMW, t)
			psess.
				oprangeset(ASBFM, t)
			psess.
				oprangeset(ASBFMW, t)
			psess.
				oprangeset(AUBFM, t)
			psess.
				oprangeset(AUBFMW, t)

		case ABFI:
			psess.
				oprangeset(ABFIW, t)
			psess.
				oprangeset(ABFXIL, t)
			psess.
				oprangeset(ABFXILW, t)
			psess.
				oprangeset(ASBFIZ, t)
			psess.
				oprangeset(ASBFIZW, t)
			psess.
				oprangeset(ASBFX, t)
			psess.
				oprangeset(ASBFXW, t)
			psess.
				oprangeset(AUBFIZ, t)
			psess.
				oprangeset(AUBFIZW, t)
			psess.
				oprangeset(AUBFX, t)
			psess.
				oprangeset(AUBFXW, t)

		case AEXTR:
			psess.
				oprangeset(AEXTRW, t)

		case ASXTB:
			psess.
				oprangeset(ASXTBW, t)
			psess.
				oprangeset(ASXTH, t)
			psess.
				oprangeset(ASXTHW, t)
			psess.
				oprangeset(ASXTW, t)
			psess.
				oprangeset(AUXTB, t)
			psess.
				oprangeset(AUXTH, t)
			psess.
				oprangeset(AUXTW, t)
			psess.
				oprangeset(AUXTBW, t)
			psess.
				oprangeset(AUXTHW, t)

		case ACCMN:
			psess.
				oprangeset(ACCMNW, t)
			psess.
				oprangeset(ACCMP, t)
			psess.
				oprangeset(ACCMPW, t)

		case ACSEL:
			psess.
				oprangeset(ACSELW, t)
			psess.
				oprangeset(ACSINC, t)
			psess.
				oprangeset(ACSINCW, t)
			psess.
				oprangeset(ACSINV, t)
			psess.
				oprangeset(ACSINVW, t)
			psess.
				oprangeset(ACSNEG, t)
			psess.
				oprangeset(ACSNEGW, t)

		case ACINC:
			psess.
				oprangeset(ACINCW, t)
			psess.
				oprangeset(ACINV, t)
			psess.
				oprangeset(ACINVW, t)
			psess.
				oprangeset(ACNEG, t)
			psess.
				oprangeset(ACNEGW, t)

		case ACSET:
			psess.
				oprangeset(ACSETW, t)
			psess.
				oprangeset(ACSETM, t)
			psess.
				oprangeset(ACSETMW, t)

		case AMOVD,
			AMOVBU,
			AB,
			ABL,
			AWORD,
			ADWORD,
			obj.ARET,
			obj.ATEXT,
			ASTP,
			ASTPW,
			ALDP:
			break

		case ALDPW:
			psess.
				oprangeset(ALDPSW, t)

		case AERET:
			psess.
				oprangeset(AWFE, t)
			psess.
				oprangeset(AWFI, t)
			psess.
				oprangeset(AYIELD, t)
			psess.
				oprangeset(ASEV, t)
			psess.
				oprangeset(ASEVL, t)
			psess.
				oprangeset(ADRPS, t)

		case ACBZ:
			psess.
				oprangeset(ACBZW, t)
			psess.
				oprangeset(ACBNZ, t)
			psess.
				oprangeset(ACBNZW, t)

		case ATBZ:
			psess.
				oprangeset(ATBNZ, t)

		case AADR, AADRP:
			break

		case ACLREX:
			break

		case ASVC:
			psess.
				oprangeset(AHVC, t)
			psess.
				oprangeset(AHLT, t)
			psess.
				oprangeset(ASMC, t)
			psess.
				oprangeset(ABRK, t)
			psess.
				oprangeset(ADCPS1, t)
			psess.
				oprangeset(ADCPS2, t)
			psess.
				oprangeset(ADCPS3, t)

		case AFADDS:
			psess.
				oprangeset(AFADDD, t)
			psess.
				oprangeset(AFSUBS, t)
			psess.
				oprangeset(AFSUBD, t)
			psess.
				oprangeset(AFMULS, t)
			psess.
				oprangeset(AFMULD, t)
			psess.
				oprangeset(AFNMULS, t)
			psess.
				oprangeset(AFNMULD, t)
			psess.
				oprangeset(AFDIVS, t)
			psess.
				oprangeset(AFMAXD, t)
			psess.
				oprangeset(AFMAXS, t)
			psess.
				oprangeset(AFMIND, t)
			psess.
				oprangeset(AFMINS, t)
			psess.
				oprangeset(AFMAXNMD, t)
			psess.
				oprangeset(AFMAXNMS, t)
			psess.
				oprangeset(AFMINNMD, t)
			psess.
				oprangeset(AFMINNMS, t)
			psess.
				oprangeset(AFDIVD, t)

		case AFMSUBD:
			psess.
				oprangeset(AFMSUBS, t)
			psess.
				oprangeset(AFMADDS, t)
			psess.
				oprangeset(AFMADDD, t)
			psess.
				oprangeset(AFNMSUBS, t)
			psess.
				oprangeset(AFNMSUBD, t)
			psess.
				oprangeset(AFNMADDS, t)
			psess.
				oprangeset(AFNMADDD, t)

		case AFCVTSD:
			psess.
				oprangeset(AFCVTDS, t)
			psess.
				oprangeset(AFABSD, t)
			psess.
				oprangeset(AFABSS, t)
			psess.
				oprangeset(AFNEGD, t)
			psess.
				oprangeset(AFNEGS, t)
			psess.
				oprangeset(AFSQRTD, t)
			psess.
				oprangeset(AFSQRTS, t)
			psess.
				oprangeset(AFRINTNS, t)
			psess.
				oprangeset(AFRINTND, t)
			psess.
				oprangeset(AFRINTPS, t)
			psess.
				oprangeset(AFRINTPD, t)
			psess.
				oprangeset(AFRINTMS, t)
			psess.
				oprangeset(AFRINTMD, t)
			psess.
				oprangeset(AFRINTZS, t)
			psess.
				oprangeset(AFRINTZD, t)
			psess.
				oprangeset(AFRINTAS, t)
			psess.
				oprangeset(AFRINTAD, t)
			psess.
				oprangeset(AFRINTXS, t)
			psess.
				oprangeset(AFRINTXD, t)
			psess.
				oprangeset(AFRINTIS, t)
			psess.
				oprangeset(AFRINTID, t)
			psess.
				oprangeset(AFCVTDH, t)
			psess.
				oprangeset(AFCVTHS, t)
			psess.
				oprangeset(AFCVTHD, t)
			psess.
				oprangeset(AFCVTSH, t)

		case AFCMPS:
			psess.
				oprangeset(AFCMPD, t)
			psess.
				oprangeset(AFCMPES, t)
			psess.
				oprangeset(AFCMPED, t)

		case AFCCMPS:
			psess.
				oprangeset(AFCCMPD, t)
			psess.
				oprangeset(AFCCMPES, t)
			psess.
				oprangeset(AFCCMPED, t)

		case AFCSELD:
			psess.
				oprangeset(AFCSELS, t)

		case AFMOVS, AFMOVD:
			break

		case AFCVTZSD:
			psess.
				oprangeset(AFCVTZSDW, t)
			psess.
				oprangeset(AFCVTZSS, t)
			psess.
				oprangeset(AFCVTZSSW, t)
			psess.
				oprangeset(AFCVTZUD, t)
			psess.
				oprangeset(AFCVTZUDW, t)
			psess.
				oprangeset(AFCVTZUS, t)
			psess.
				oprangeset(AFCVTZUSW, t)

		case ASCVTFD:
			psess.
				oprangeset(ASCVTFS, t)
			psess.
				oprangeset(ASCVTFWD, t)
			psess.
				oprangeset(ASCVTFWS, t)
			psess.
				oprangeset(AUCVTFD, t)
			psess.
				oprangeset(AUCVTFS, t)
			psess.
				oprangeset(AUCVTFWD, t)
			psess.
				oprangeset(AUCVTFWS, t)

		case ASYS:
			psess.
				oprangeset(AAT, t)
			psess.
				oprangeset(ADC, t)
			psess.
				oprangeset(AIC, t)
			psess.
				oprangeset(ATLBI, t)

		case ASYSL, AHINT:
			break

		case ADMB:
			psess.
				oprangeset(ADSB, t)
			psess.
				oprangeset(AISB, t)

		case AMRS, AMSR:
			break

		case ALDAR:
			psess.
				oprangeset(ALDARW, t)
			psess.
				oprangeset(ALDARB, t)
			psess.
				oprangeset(ALDARH, t)
			fallthrough

		case ALDXR:
			psess.
				oprangeset(ALDXRB, t)
			psess.
				oprangeset(ALDXRH, t)
			psess.
				oprangeset(ALDXRW, t)

		case ALDAXR:
			psess.
				oprangeset(ALDAXRB, t)
			psess.
				oprangeset(ALDAXRH, t)
			psess.
				oprangeset(ALDAXRW, t)

		case ALDXP:
			psess.
				oprangeset(ALDXPW, t)
			psess.
				oprangeset(ALDAXP, t)
			psess.
				oprangeset(ALDAXPW, t)

		case ASTLR:
			psess.
				oprangeset(ASTLRB, t)
			psess.
				oprangeset(ASTLRH, t)
			psess.
				oprangeset(ASTLRW, t)

		case ASTXR:
			psess.
				oprangeset(ASTXRB, t)
			psess.
				oprangeset(ASTXRH, t)
			psess.
				oprangeset(ASTXRW, t)

		case ASTLXR:
			psess.
				oprangeset(ASTLXRB, t)
			psess.
				oprangeset(ASTLXRH, t)
			psess.
				oprangeset(ASTLXRW, t)

		case ASTXP:
			psess.
				oprangeset(ASTLXP, t)
			psess.
				oprangeset(ASTLXPW, t)
			psess.
				oprangeset(ASTXPW, t)

		case AVADDP:
			psess.
				oprangeset(AVAND, t)
			psess.
				oprangeset(AVCMEQ, t)
			psess.
				oprangeset(AVORR, t)
			psess.
				oprangeset(AVEOR, t)

		case AVADD:
			psess.
				oprangeset(AVSUB, t)

		case AAESD:
			psess.
				oprangeset(AAESE, t)
			psess.
				oprangeset(AAESMC, t)
			psess.
				oprangeset(AAESIMC, t)
			psess.
				oprangeset(ASHA1SU1, t)
			psess.
				oprangeset(ASHA256SU0, t)

		case ASHA1C:
			psess.
				oprangeset(ASHA1P, t)
			psess.
				oprangeset(ASHA1M, t)

		case ASHA256H:
			psess.
				oprangeset(ASHA256H2, t)

		case ASHA1SU0:
			psess.
				oprangeset(ASHA256SU1, t)

		case AVADDV:
			psess.
				oprangeset(AVUADDLV, t)

		case AVFMLA:
			psess.
				oprangeset(AVFMLS, t)

		case AVPMULL:
			psess.
				oprangeset(AVPMULL2, t)

		case AVUSHR:
			psess.
				oprangeset(AVSHL, t)
			psess.
				oprangeset(AVSRI, t)

		case AVREV32:
			psess.
				oprangeset(AVRBIT, t)
			psess.
				oprangeset(AVREV64, t)

		case AVZIP1:
			psess.
				oprangeset(AVZIP2, t)

		case ASHA1H,
			AVCNT,
			AVMOV,
			AVLD1,
			AVST1,
			AVTBL,
			AVDUP,
			AVMOVI,
			APRFM,
			AVEXT:
			break

		case obj.ANOP,
			obj.AUNDEF,
			obj.AFUNCDATA,
			obj.APCDATA,
			obj.ADUFFZERO,
			obj.ADUFFCOPY:
			break
		}
	}
}

func (c *ctxt7) chipfloat7(e float64) int {
	ei := math.Float64bits(e)
	l := uint32(int32(ei))
	h := uint32(int32(ei >> 32))

	if l != 0 || h&0xffff != 0 {
		return -1
	}
	h1 := h & 0x7fc00000
	if h1 != 0x40000000 && h1 != 0x3fc00000 {
		return -1
	}
	n := 0

	if h&0x80000000 != 0 {
		n |= 1 << 7
	}

	if h1 == 0x3fc00000 {
		n |= 1 << 6
	}

	n |= int((h >> 16) & 0x3f)

	return n
}

/* form offset parameter to SYS; special register number */
func SYSARG5(op0 int, op1 int, Cn int, Cm int, op2 int) int {
	return op0<<19 | op1<<16 | Cn<<12 | Cm<<8 | op2<<5
}

func SYSARG4(op1 int, Cn int, Cm int, op2 int) int {
	return SYSARG5(0, op1, Cn, Cm, op2)
}

/* checkindex checks if index >= 0 && index <= maxindex */
func (c *ctxt7) checkindex(p *obj.Prog, index, maxindex int) {
	if index < 0 || index > maxindex {
		c.ctxt.Diag("register element index out of range 0 to %d: %v", maxindex, p)
	}
}

/* checkoffset checks whether the immediate offset is valid for VLD1.P and VST1.P */
func (c *ctxt7) checkoffset(p *obj.Prog, as obj.As) {
	var offset, list, n int64
	switch as {
	case AVLD1:
		offset = p.From.Offset
		list = p.To.Offset
	case AVST1:
		offset = p.To.Offset
		list = p.From.Offset
	default:
		c.ctxt.Diag("invalid operation on op %v", p.As)
	}
	opcode := (list >> 12) & 15
	q := (list >> 30) & 1
	if offset == 0 {
		return
	}
	switch opcode {
	case 0x7:
		n = 1
	case 0xa:
		n = 2
	case 0x6:
		n = 3
	case 0x2:
		n = 4
	default:
		c.ctxt.Diag("invalid register numbers in ARM64 register list: %v", p)
	}
	if !(q == 0 && offset == n*8) && !(q == 1 && offset == n*16) {
		c.ctxt.Diag("invalid post-increment offset: %v", p)
	}
}

/* checkShiftAmount checks whether the index shift amount is valid */
/* for load with register offset instructions */
func (c *ctxt7) checkShiftAmount(p *obj.Prog, a *obj.Addr) {
	var amount int16
	amount = (a.Index >> 5) & 7
	switch p.As {
	case AMOVB, AMOVBU:
		if amount != 0 {
			c.ctxt.Diag("invalid index shift amount: %v", p)
		}
	case AMOVH, AMOVHU:
		if amount != 1 && amount != 0 {
			c.ctxt.Diag("invalid index shift amount: %v", p)
		}
	case AMOVW, AMOVWU:
		if amount != 2 && amount != 0 {
			c.ctxt.Diag("invalid index shift amount: %v", p)
		}
	case AMOVD:
		if amount != 3 && amount != 0 {
			c.ctxt.Diag("invalid index shift amount: %v", p)
		}
	default:
		panic("invalid operation")
	}
}

func (c *ctxt7) asmout(psess *PackageSession, p *obj.Prog, o *Optab, out []uint32) {
	o1 := uint32(0)
	o2 := uint32(0)
	o3 := uint32(0)
	o4 := uint32(0)
	o5 := uint32(0)
	if false {
		fmt.Printf("%x: %v\ttype %d\n", uint32(p.Pc), p, o.type_)
	}
	switch o.type_ {
	default:
		c.ctxt.Diag("%v: unknown asm %d", p, o.type_)

	case 0:
		break

	case 1:
		o1 = c.oprrr(p, p.As)

		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		r := int(p.Reg)
		if p.To.Type == obj.TYPE_NONE {
			rt = REGZERO
		}
		if r == 0 {
			r = rt
		}
		o1 |= (uint32(rf&31) << 16) | (uint32(r&31) << 5) | uint32(rt&31)

	case 2:
		o1 = c.opirr(p, p.As)

		rt := int(p.To.Reg)
		if p.To.Type == obj.TYPE_NONE {
			if (o1 & Sbit) == 0 {
				c.ctxt.Diag("ineffective ZR destination\n%v", p)
			}
			rt = REGZERO
		}

		r := int(p.Reg)
		if r == 0 {
			r = rt
		}
		v := int32(c.regoff(&p.From))
		o1 = c.oaddi(p, int32(o1), v, r, rt)

	case 3:
		o1 = c.oprrr(p, p.As)

		amount := (p.From.Offset >> 10) & 63
		is64bit := o1 & (1 << 31)
		if is64bit == 0 && amount >= 32 {
			c.ctxt.Diag("shift amount out of range 0 to 31: %v", p)
		}
		o1 |= uint32(p.From.Offset)
		rt := int(p.To.Reg)
		if p.To.Type == obj.TYPE_NONE {
			rt = REGZERO
		}
		r := int(p.Reg)
		if p.As == AMVN || p.As == AMVNW {
			r = REGZERO
		} else if r == 0 {
			r = rt
		}
		o1 |= (uint32(r&31) << 5) | uint32(rt&31)

	case 4:
		o1 = c.opirr(p, p.As)

		rt := int(p.To.Reg)
		r := int(o.param)
		if r == 0 {
			r = REGZERO
		} else if r == REGFROM {
			r = int(p.From.Reg)
		}
		if r == 0 {
			r = REGSP
		}
		v := int32(c.regoff(&p.From))
		if (v & 0xFFF000) != 0 {
			v >>= 12
			o1 |= 1 << 22
		}

		o1 |= ((uint32(v) & 0xFFF) << 10) | (uint32(r&31) << 5) | uint32(rt&31)

	case 5:
		o1 = c.opbra(p, p.As)

		if p.To.Sym == nil {
			o1 |= uint32(c.brdist(p, 0, 26, 2))
			break
		}

		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 4
		rel.Sym = p.To.Sym
		rel.Add = p.To.Offset
		rel.Type = objabi.R_CALLARM64

	case 6:
		o1 = c.opbrr(p, p.As)

		o1 |= uint32(p.To.Reg&31) << 5
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 0
		rel.Type = objabi.R_CALLIND

	case 7:
		o1 = c.opbra(p, p.As)

		o1 |= uint32(c.brdist(p, 0, 19, 2) << 5)

	case 8:
		rt := int(p.To.Reg)

		rf := int(p.Reg)
		if rf == 0 {
			rf = rt
		}
		v := int32(p.From.Offset)
		switch p.As {
		case AASR:
			o1 = c.opbfm(p, ASBFM, int(v), 63, rf, rt)

		case AASRW:
			o1 = c.opbfm(p, ASBFMW, int(v), 31, rf, rt)

		case ALSL:
			o1 = c.opbfm(p, AUBFM, int((64-v)&63), int(63-v), rf, rt)

		case ALSLW:
			o1 = c.opbfm(p, AUBFMW, int((32-v)&31), int(31-v), rf, rt)

		case ALSR:
			o1 = c.opbfm(p, AUBFM, int(v), 63, rf, rt)

		case ALSRW:
			o1 = c.opbfm(p, AUBFMW, int(v), 31, rf, rt)

		case AROR:
			o1 = c.opextr(p, AEXTR, v, rf, rf, rt)

		case ARORW:
			o1 = c.opextr(p, AEXTRW, v, rf, rf, rt)

		default:
			c.ctxt.Diag("bad shift $con\n%v", p)
			break
		}

	case 9:
		o1 = c.oprrr(p, p.As)

		r := int(p.Reg)
		if r == 0 {
			r = int(p.To.Reg)
		}
		o1 |= (uint32(p.From.Reg&31) << 16) | (uint32(r&31) << 5) | uint32(p.To.Reg&31)

	case 10:
		o1 = c.opimm(p, p.As)

		if p.From.Type != obj.TYPE_NONE {
			o1 |= uint32((p.From.Offset & 0xffff) << 5)
		}

	case 11:
		c.aclass(&p.To)

		o1 = uint32(c.instoffset)
		o2 = uint32(c.instoffset >> 32)
		if p.To.Sym != nil {
			rel := obj.Addrel(c.cursym)
			rel.Off = int32(c.pc)
			rel.Siz = 8
			rel.Sym = p.To.Sym
			rel.Add = p.To.Offset
			rel.Type = objabi.R_ADDR
			o2 = 0
			o1 = o2
		}

	case 12:
		o1 = c.omovlit(p.As, p, &p.From, int(p.To.Reg))

	case 13:
		o1 = c.omovlit(AMOVD, p, &p.From, REGTMP)

		if !(o1 != 0) {
			break
		}
		rt := int(p.To.Reg)
		if p.To.Type == obj.TYPE_NONE {
			rt = REGZERO
		}
		r := int(p.Reg)
		if r == 0 {
			r = rt
		}
		if p.To.Type != obj.TYPE_NONE && (p.To.Reg == REGSP || r == REGSP) {
			o2 = c.opxrrr(p, p.As, false)
			o2 |= REGTMP & 31 << 16
			o2 |= LSL0_64
		} else {
			o2 = c.oprrr(p, p.As)
			o2 |= REGTMP & 31 << 16
		}

		o2 |= uint32(r&31) << 5
		o2 |= uint32(rt & 31)

	case 14:
		if c.aclass(&p.To) == C_ADDR {
			c.ctxt.Diag("address constant needs DWORD\n%v", p)
		}
		o1 = uint32(c.instoffset)
		if p.To.Sym != nil {

			rel := obj.Addrel(c.cursym)

			rel.Off = int32(c.pc)
			rel.Siz = 4
			rel.Sym = p.To.Sym
			rel.Add = p.To.Offset
			rel.Type = objabi.R_ADDR
			o1 = 0
		}

	case 15:
		o1 = c.oprrr(p, p.As)

		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		var r int
		var ra int
		if p.From3Type() == obj.TYPE_REG {
			r = int(p.GetFrom3().Reg)
			ra = int(p.Reg)
			if ra == 0 {
				ra = REGZERO
			}
		} else {
			r = int(p.Reg)
			if r == 0 {
				r = rt
			}
			ra = REGZERO
		}

		o1 |= (uint32(rf&31) << 16) | (uint32(ra&31) << 10) | (uint32(r&31) << 5) | uint32(rt&31)

	case 16:
		o1 = c.oprrr(p, p.As)

		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		r := int(p.Reg)
		if r == 0 {
			r = rt
		}
		o1 |= (uint32(rf&31) << 16) | (uint32(r&31) << 5) | REGTMP&31
		o2 = c.oprrr(p, AMSUBW)
		o2 |= o1 & (1 << 31)
		o2 |= (uint32(rf&31) << 16) | (uint32(r&31) << 10) | (REGTMP & 31 << 5) | uint32(rt&31)

	case 17:
		o1 = c.oprrr(p, p.As)

		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		r := int(p.Reg)
		if p.To.Type == obj.TYPE_NONE {
			rt = REGZERO
		}
		if r == 0 {
			r = REGZERO
		}
		o1 |= (uint32(rf&31) << 16) | (uint32(r&31) << 5) | uint32(rt&31)

	case 18:
		o1 = c.oprrr(p, p.As)

		cond := int(p.From.Reg)
		if cond < COND_EQ || cond > COND_NV {
			c.ctxt.Diag("invalid condition: %v", p)
		} else {
			cond -= COND_EQ
		}

		r := int(p.Reg)
		var rf int
		if r != 0 {
			if p.From3Type() == obj.TYPE_NONE {

				rf = r
				cond ^= 1
			} else {
				rf = int(p.GetFrom3().Reg)
			}
		} else {

			rf = REGZERO
			r = rf
			cond ^= 1
		}

		rt := int(p.To.Reg)
		o1 |= (uint32(rf&31) << 16) | (uint32(cond&15) << 12) | (uint32(r&31) << 5) | uint32(rt&31)

	case 19:
		nzcv := int(p.To.Offset)

		cond := int(p.From.Reg)
		if cond < COND_EQ || cond > COND_NV {
			c.ctxt.Diag("invalid condition\n%v", p)
		} else {
			cond -= COND_EQ
		}
		var rf int
		if p.GetFrom3().Type == obj.TYPE_REG {
			o1 = c.oprrr(p, p.As)
			rf = int(p.GetFrom3().Reg)
		} else {
			o1 = c.opirr(p, p.As)
			rf = int(p.GetFrom3().Offset & 0x1F)
		}

		o1 |= (uint32(rf&31) << 16) | (uint32(cond&15) << 12) | (uint32(p.Reg&31) << 5) | uint32(nzcv)

	case 20:
		v := int32(c.regoff(&p.To))
		sz := int32(1 << uint(movesize(p.As)))

		r := int(p.To.Reg)
		if r == 0 {
			r = int(o.param)
		}
		if v < 0 || v%sz != 0 {
			o1 = c.olsr9s(p, int32(c.opstr9(p, p.As)), v, r, int(p.From.Reg))
		} else {
			v = int32(c.offsetshift(psess, p, int64(v), int(o.a4)))
			o1 = c.olsr12u(p, int32(c.opstr12(p, p.As)), v, r, int(p.From.Reg))
		}

	case 21:
		v := int32(c.regoff(&p.From))
		sz := int32(1 << uint(movesize(p.As)))

		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		if v < 0 || v%sz != 0 {
			o1 = c.olsr9s(p, int32(c.opldr9(p, p.As)), v, r, int(p.To.Reg))
		} else {
			v = int32(c.offsetshift(psess, p, int64(v), int(o.a1)))

			o1 = c.olsr12u(p, int32(c.opldr12(p, p.As)), v, r, int(p.To.Reg))
		}

	case 22:
		v := int32(p.From.Offset)

		if v < -256 || v > 255 {
			c.ctxt.Diag("offset out of range [-255,254]: %v", p)
		}
		o1 = c.opldrpp(p, p.As)
		if o.scond == C_XPOST {
			o1 |= 1 << 10
		} else {
			o1 |= 3 << 10
		}
		o1 |= ((uint32(v) & 0x1FF) << 12) | (uint32(p.From.Reg&31) << 5) | uint32(p.To.Reg&31)

	case 23:
		v := int32(p.To.Offset)

		if v < -256 || v > 255 {
			c.ctxt.Diag("offset out of range [-255,254]: %v", p)
		}
		o1 = LD2STR(c.opldrpp(p, p.As))
		if o.scond == C_XPOST {
			o1 |= 1 << 10
		} else {
			o1 |= 3 << 10
		}
		o1 |= ((uint32(v) & 0x1FF) << 12) | (uint32(p.To.Reg&31) << 5) | uint32(p.From.Reg&31)

	case 24:
		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		s := rf == REGSP || rt == REGSP
		if p.As == AMVN || p.As == AMVNW {
			if s {
				c.ctxt.Diag("illegal SP reference\n%v", p)
			}
			o1 = c.oprrr(p, p.As)
			o1 |= (uint32(rf&31) << 16) | (REGZERO & 31 << 5) | uint32(rt&31)
		} else if s {
			o1 = c.opirr(p, p.As)
			o1 |= (uint32(rf&31) << 5) | uint32(rt&31)
		} else {
			o1 = c.oprrr(p, p.As)
			o1 |= (uint32(rf&31) << 16) | (REGZERO & 31 << 5) | uint32(rt&31)
		}

	case 25:
		o1 = c.oprrr(p, p.As)

		rf := int(p.From.Reg)
		if rf == C_NONE {
			rf = int(p.To.Reg)
		}
		rt := int(p.To.Reg)
		o1 |= (uint32(rf&31) << 16) | (REGZERO & 31 << 5) | uint32(rt&31)

	case 26:
		o1 = c.oprrr(p, p.As)

		o1 |= uint32(p.From.Offset)
		rt := int(p.To.Reg)
		o1 |= (REGZERO & 31 << 5) | uint32(rt&31)

	case 27:
		if (p.From.Reg-obj.RBaseARM64)&REG_EXT != 0 {
			amount := (p.From.Reg >> 5) & 7
			if amount > 4 {
				c.ctxt.Diag("shift amount out of range 0 to 4: %v", p)
			}
			o1 = c.opxrrr(p, p.As, true)
			o1 |= c.encRegShiftOrExt(&p.From, p.From.Reg)
		} else {
			o1 = c.opxrrr(p, p.As, false)
			o1 |= uint32(p.From.Reg&31) << 16
		}
		rt := int(p.To.Reg)
		if p.To.Type == obj.TYPE_NONE {
			rt = REGZERO
		}
		r := int(p.Reg)
		if r == 0 {
			r = rt
		}
		o1 |= (uint32(r&31) << 5) | uint32(rt&31)

	case 28:
		o1 = c.omovlit(AMOVD, p, &p.From, REGTMP)

		if !(o1 != 0) {
			break
		}
		rt := int(p.To.Reg)
		if p.To.Type == obj.TYPE_NONE {
			rt = REGZERO
		}
		r := int(p.Reg)
		if r == 0 {
			r = rt
		}
		o2 = c.oprrr(p, p.As)
		o2 |= REGTMP & 31 << 16
		o2 |= uint32(r&31) << 5
		o2 |= uint32(rt & 31)

	case 29:
		fc := c.aclass(&p.From)
		tc := c.aclass(&p.To)
		if (p.As == AFMOVD || p.As == AFMOVS) && (fc == C_REG || fc == C_ZCON || tc == C_REG || tc == C_ZCON) {

			o1 = FPCVTI(0, 0, 0, 0, 6)
			if p.As == AFMOVD {
				o1 |= 1<<31 | 1<<22
			}
			if fc == C_REG || fc == C_ZCON {
				o1 |= 1 << 16
			}
		} else {
			o1 = c.oprrr(p, p.As)
		}
		o1 |= uint32(p.From.Reg&31)<<5 | uint32(p.To.Reg&31)

	case 30:

		s := movesize(o.as)
		if s < 0 {
			c.ctxt.Diag("unexpected long move, op %v tab %v\n%v", p.As, o.as, p)
		}

		r := int(p.To.Reg)
		if r == 0 {
			r = int(o.param)
		}

		v := int32(c.regoff(&p.To))
		var hi int32
		if v < 0 || (v&((1<<uint(s))-1)) != 0 {

			goto storeusepool
		}

		hi = v - (v & (0xFFF << uint(s)))
		if hi&0xFFF != 0 {
			c.ctxt.Diag("internal: miscalculated offset %d [%d]\n%v", v, s, p)
		}
		if hi&^0xFFF000 != 0 {

			goto storeusepool
		}

		o1 = c.oaddi(p, int32(c.opirr(p, AADD)), hi, r, REGTMP)
		o2 = c.olsr12u(p, int32(c.opstr12(p, p.As)), ((v-hi)>>uint(s))&0xFFF, REGTMP, int(p.From.Reg))
		break

	storeusepool:
		if r == REGTMP || p.From.Reg == REGTMP {
			c.ctxt.Diag("REGTMP used in large offset store: %v", p)
		}
		o1 = c.omovlit(AMOVD, p, &p.To, REGTMP)
		o2 = c.olsxrr(p, int32(c.opstrr(p, p.As, false)), int(p.From.Reg), r, REGTMP)

	case 31:

		s := movesize(o.as)
		if s < 0 {
			c.ctxt.Diag("unexpected long move, op %v tab %v\n%v", p.As, o.as, p)
		}

		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}

		v := int32(c.regoff(&p.From))
		var hi int32
		if v < 0 || (v&((1<<uint(s))-1)) != 0 {

			goto loadusepool
		}

		hi = v - (v & (0xFFF << uint(s)))
		if (hi & 0xFFF) != 0 {
			c.ctxt.Diag("internal: miscalculated offset %d [%d]\n%v", v, s, p)
		}
		if hi&^0xFFF000 != 0 {

			goto loadusepool
		}

		o1 = c.oaddi(p, int32(c.opirr(p, AADD)), hi, r, REGTMP)
		o2 = c.olsr12u(p, int32(c.opldr12(p, p.As)), ((v-hi)>>uint(s))&0xFFF, REGTMP, int(p.To.Reg))
		break

	loadusepool:
		if r == REGTMP || p.From.Reg == REGTMP {
			c.ctxt.Diag("REGTMP used in large offset load: %v", p)
		}
		o1 = c.omovlit(AMOVD, p, &p.From, REGTMP)
		o2 = c.olsxrr(p, int32(c.opldrr(p, p.As, false)), int(p.To.Reg), r, REGTMP)

	case 32:
		o1 = c.omovconst(p.As, p, &p.From, int(p.To.Reg))

	case 33:
		o1 = c.opirr(p, p.As)

		d := p.From.Offset
		s := movcon(d)
		if s < 0 || s >= 4 {
			c.ctxt.Diag("bad constant for MOVK: %#x\n%v", uint64(d), p)
		}
		if (o1&S64) == 0 && s >= 2 {
			c.ctxt.Diag("illegal bit position\n%v", p)
		}
		if ((d >> uint(s*16)) >> 16) != 0 {
			c.ctxt.Diag("requires uimm16\n%v", p)
		}
		rt := int(p.To.Reg)

		o1 |= uint32((((d >> uint(s*16)) & 0xFFFF) << 5) | int64((uint32(s)&3)<<21) | int64(rt&31))

	case 34:
		o1 = c.omovlit(AMOVD, p, &p.From, REGTMP)

		if !(o1 != 0) {
			break
		}
		o2 = c.opxrrr(p, AADD, false)
		o2 |= REGTMP & 31 << 16
		o2 |= LSL0_64
		r := int(p.From.Reg)
		if r == 0 {
			r = int(o.param)
		}
		o2 |= uint32(r&31) << 5
		o2 |= uint32(p.To.Reg & 31)

	case 35:
		o1 = c.oprrr(p, AMRS)

		v := uint32(0)
		for i := 0; i < len(psess.systemreg); i++ {
			if psess.systemreg[i].reg == p.From.Reg {
				v = psess.systemreg[i].enc
				break
			}
		}
		if v == 0 {
			c.ctxt.Diag("illegal system register:\n%v", p)
		}
		if (o1 & (v &^ (3 << 19))) != 0 {
			c.ctxt.Diag("MRS register value overlap\n%v", p)
		}

		o1 |= v
		o1 |= uint32(p.To.Reg & 31)

	case 36:
		o1 = c.oprrr(p, AMSR)

		v := uint32(0)
		for i := 0; i < len(psess.systemreg); i++ {
			if psess.systemreg[i].reg == p.To.Reg {
				v = psess.systemreg[i].enc
				break
			}
		}
		if v == 0 {
			c.ctxt.Diag("illegal system register:\n%v", p)
		}
		if (o1 & (v &^ (3 << 19))) != 0 {
			c.ctxt.Diag("MSR register value overlap\n%v", p)
		}

		o1 |= v
		o1 |= uint32(p.From.Reg & 31)

	case 37:
		if (uint64(p.From.Offset) &^ uint64(0xF)) != 0 {
			c.ctxt.Diag("illegal immediate for PSTATE field\n%v", p)
		}
		o1 = c.opirr(p, AMSR)
		o1 |= uint32((p.From.Offset & 0xF) << 8)
		v := uint32(0)
		for i := 0; i < len(psess.pstatefield); i++ {
			if psess.pstatefield[i].reg == p.To.Reg {
				v = psess.pstatefield[i].enc
				break
			}
		}

		if v == 0 {
			c.ctxt.Diag("illegal PSTATE field for immediate move\n%v", p)
		}
		o1 |= v

	case 38:
		o1 = c.opimm(p, p.As)

		if p.To.Type == obj.TYPE_NONE {
			o1 |= 0xF << 8
		} else {
			o1 |= uint32((p.To.Offset & 0xF) << 8)
		}

	case 39:
		o1 = c.opirr(p, p.As)

		o1 |= uint32(p.From.Reg & 31)
		o1 |= uint32(c.brdist(p, 0, 19, 2) << 5)

	case 40:
		o1 = c.opirr(p, p.As)

		v := int32(p.From.Offset)
		if v < 0 || v > 63 {
			c.ctxt.Diag("illegal bit number\n%v", p)
		}
		o1 |= ((uint32(v) & 0x20) << (31 - 5)) | ((uint32(v) & 0x1F) << 19)
		o1 |= uint32(c.brdist(p, 0, 14, 2) << 5)
		o1 |= uint32(p.Reg & 31)

	case 41:
		o1 = c.op0(p, p.As)

	case 42:
		o1 = c.opbfm(p, p.As, int(p.From.Offset), int(p.GetFrom3().Offset), int(p.Reg), int(p.To.Reg))

	case 43:
		r := int(p.From.Offset)
		s := int(p.GetFrom3().Offset)
		rf := int(p.Reg)
		rt := int(p.To.Reg)
		if rf == 0 {
			rf = rt
		}
		switch p.As {
		case ABFI:
			o1 = c.opbfm(p, ABFM, 64-r, s-1, rf, rt)

		case ABFIW:
			o1 = c.opbfm(p, ABFMW, 32-r, s-1, rf, rt)

		case ABFXIL:
			o1 = c.opbfm(p, ABFM, r, r+s-1, rf, rt)

		case ABFXILW:
			o1 = c.opbfm(p, ABFMW, r, r+s-1, rf, rt)

		case ASBFIZ:
			o1 = c.opbfm(p, ASBFM, 64-r, s-1, rf, rt)

		case ASBFIZW:
			o1 = c.opbfm(p, ASBFMW, 32-r, s-1, rf, rt)

		case ASBFX:
			o1 = c.opbfm(p, ASBFM, r, r+s-1, rf, rt)

		case ASBFXW:
			o1 = c.opbfm(p, ASBFMW, r, r+s-1, rf, rt)

		case AUBFIZ:
			o1 = c.opbfm(p, AUBFM, 64-r, s-1, rf, rt)

		case AUBFIZW:
			o1 = c.opbfm(p, AUBFMW, 32-r, s-1, rf, rt)

		case AUBFX:
			o1 = c.opbfm(p, AUBFM, r, r+s-1, rf, rt)

		case AUBFXW:
			o1 = c.opbfm(p, AUBFMW, r, r+s-1, rf, rt)

		default:
			c.ctxt.Diag("bad bfm alias\n%v", p)
			break
		}

	case 44:
		o1 = c.opextr(p, p.As, int32(p.From.Offset), int(p.GetFrom3().Reg), int(p.Reg), int(p.To.Reg))

	case 45:
		rf := int(p.From.Reg)

		rt := int(p.To.Reg)
		as := p.As
		if rf == REGZERO {
			as = AMOVWU
		}
		switch as {
		case AMOVB, ASXTB:
			o1 = c.opbfm(p, ASBFM, 0, 7, rf, rt)

		case AMOVH, ASXTH:
			o1 = c.opbfm(p, ASBFM, 0, 15, rf, rt)

		case AMOVW, ASXTW:
			o1 = c.opbfm(p, ASBFM, 0, 31, rf, rt)

		case AMOVBU, AUXTB:
			o1 = c.opbfm(p, AUBFM, 0, 7, rf, rt)

		case AMOVHU, AUXTH:
			o1 = c.opbfm(p, AUBFM, 0, 15, rf, rt)

		case AMOVWU:
			o1 = c.oprrr(p, as) | (uint32(rf&31) << 16) | (REGZERO & 31 << 5) | uint32(rt&31)

		case AUXTW:
			o1 = c.opbfm(p, AUBFM, 0, 31, rf, rt)

		case ASXTBW:
			o1 = c.opbfm(p, ASBFMW, 0, 7, rf, rt)

		case ASXTHW:
			o1 = c.opbfm(p, ASBFMW, 0, 15, rf, rt)

		case AUXTBW:
			o1 = c.opbfm(p, AUBFMW, 0, 7, rf, rt)

		case AUXTHW:
			o1 = c.opbfm(p, AUBFMW, 0, 15, rf, rt)

		default:
			c.ctxt.Diag("bad sxt %v", as)
			break
		}

	case 46:
		o1 = c.opbit(p, p.As)

		o1 |= uint32(p.From.Reg&31) << 5
		o1 |= uint32(p.To.Reg & 31)

	case 47:
		rs := p.From.Reg
		rt := p.RegTo2
		rb := p.To.Reg
		switch p.As {
		case ASWPD, ALDADDALD, ALDADDD, ALDANDD, ALDEORD, ALDORD:
			o1 = 3 << 30
		case ASWPW, ALDADDALW, ALDADDW, ALDANDW, ALDEORW, ALDORW:
			o1 = 2 << 30
		case ASWPH, ALDADDH, ALDANDH, ALDEORH, ALDORH:
			o1 = 1 << 30
		case ASWPB, ALDADDB, ALDANDB, ALDEORB, ALDORB:
			o1 = 0 << 30
		default:
			c.ctxt.Diag("illegal instruction: %v\n", p)
		}
		switch p.As {
		case ASWPD, ASWPW, ASWPH, ASWPB:
			o1 |= 0x20 << 10
		case ALDADDALD, ALDADDALW, ALDADDD, ALDADDW, ALDADDH, ALDADDB:
			o1 |= 0x00 << 10
		case ALDANDD, ALDANDW, ALDANDH, ALDANDB:
			o1 |= 0x04 << 10
		case ALDEORD, ALDEORW, ALDEORH, ALDEORB:
			o1 |= 0x08 << 10
		case ALDORD, ALDORW, ALDORH, ALDORB:
			o1 |= 0x0c << 10
		}
		switch p.As {
		case ALDADDALD, ALDADDALW:
			o1 |= 3 << 22
		}
		o1 |= 0x1c1<<21 | uint32(rs&31)<<16 | uint32(rb&31)<<5 | uint32(rt&31)

	case 50:
		o1 = c.opirr(p, p.As)

		if (p.From.Offset &^ int64(SYSARG4(0x7, 0xF, 0xF, 0x7))) != 0 {
			c.ctxt.Diag("illegal SYS argument\n%v", p)
		}
		o1 |= uint32(p.From.Offset)
		if p.To.Type == obj.TYPE_REG {
			o1 |= uint32(p.To.Reg & 31)
		} else if p.Reg != 0 {
			o1 |= uint32(p.Reg & 31)
		} else {
			o1 |= 0x1F
		}

	case 51:
		o1 = c.opirr(p, p.As)

		if p.From.Type == obj.TYPE_CONST {
			o1 |= uint32((p.From.Offset & 0xF) << 8)
		}

	case 52:
		o1 = c.opirr(p, p.As)

		o1 |= uint32((p.From.Offset & 0x7F) << 5)

	case 53:
		a := p.As
		rt := int(p.To.Reg)
		if p.To.Type == obj.TYPE_NONE {
			rt = REGZERO
		}
		r := int(p.Reg)
		if r == 0 {
			r = rt
		}
		mode := 64
		v := uint64(p.From.Offset)
		switch p.As {
		case AANDW, AORRW, AEORW, AANDSW, ATSTW:
			mode = 32
		case ABIC, AORN, AEON, ABICS:
			v = ^v
		case ABICW, AORNW, AEONW, ABICSW:
			v = ^v
			mode = 32
		}
		o1 = c.opirr(p, a)
		o1 |= bitconEncode(v, mode) | uint32(r&31)<<5 | uint32(rt&31)

	case 54:
		o1 = c.oprrr(p, p.As)

		var rf int
		if p.From.Type == obj.TYPE_CONST {
			rf = c.chipfloat7(p.From.Val.(float64))
			if rf < 0 || true {
				c.ctxt.Diag("invalid floating-point immediate\n%v", p)
				rf = 0
			}

			rf |= (1 << 3)
		} else {
			rf = int(p.From.Reg)
		}
		rt := int(p.To.Reg)
		r := int(p.Reg)
		if (o1&(0x1F<<24)) == (0x1E<<24) && (o1&(1<<11)) == 0 {
			r = rf
			rf = 0
		} else if r == 0 {
			r = rt
		}
		o1 |= (uint32(rf&31) << 16) | (uint32(r&31) << 5) | uint32(rt&31)

	case 56:
		o1 = c.oprrr(p, p.As)

		var rf int
		if p.From.Type == obj.TYPE_FCONST {
			o1 |= 8
			rf = 0
		} else {
			rf = int(p.From.Reg)
		}
		rt := int(p.Reg)
		o1 |= uint32(rf&31)<<16 | uint32(rt&31)<<5

	case 57:
		o1 = c.oprrr(p, p.As)

		cond := int(p.From.Reg)
		if cond < COND_EQ || cond > COND_NV {
			c.ctxt.Diag("invalid condition\n%v", p)
		} else {
			cond -= COND_EQ
		}

		nzcv := int(p.To.Offset)
		if nzcv&^0xF != 0 {
			c.ctxt.Diag("implausible condition\n%v", p)
		}
		rf := int(p.Reg)
		if p.GetFrom3() == nil || p.GetFrom3().Reg < REG_F0 || p.GetFrom3().Reg > REG_F31 {
			c.ctxt.Diag("illegal FCCMP\n%v", p)
			break
		}
		rt := int(p.GetFrom3().Reg)
		o1 |= uint32(rf&31)<<16 | uint32(cond&15)<<12 | uint32(rt&31)<<5 | uint32(nzcv)

	case 58:
		o1 = c.opload(p, p.As)

		o1 |= 0x1F << 16
		o1 |= uint32(p.From.Reg&31) << 5
		if p.As == ALDXP || p.As == ALDXPW || p.As == ALDAXP || p.As == ALDAXPW {
			o1 |= uint32(p.To.Offset&31) << 10
		} else {
			o1 |= 0x1F << 10
		}
		o1 |= uint32(p.To.Reg & 31)

	case 59:
		o1 = c.opstore(p, p.As)

		if p.RegTo2 != obj.REG_NONE {
			o1 |= uint32(p.RegTo2&31) << 16
		} else {
			o1 |= 0x1F << 16
		}
		if p.As == ASTXP || p.As == ASTXPW || p.As == ASTLXP || p.As == ASTLXPW {
			o1 |= uint32(p.From.Offset&31) << 10
		}
		o1 |= uint32(p.To.Reg&31)<<5 | uint32(p.From.Reg&31)

	case 60:
		d := c.brdist(p, 12, 21, 0)

		o1 = ADR(1, uint32(d), uint32(p.To.Reg))

	case 61:
		d := c.brdist(p, 0, 21, 0)

		o1 = ADR(0, uint32(d), uint32(p.To.Reg))

	case 62:
		if p.Reg == REGTMP {
			c.ctxt.Diag("cannot use REGTMP as source: %v\n", p)
		}
		o1 = c.omovconst(AMOVD, p, &p.From, REGTMP)

		rt := int(p.To.Reg)
		if p.To.Type == obj.TYPE_NONE {
			rt = REGZERO
		}
		r := int(p.Reg)
		if r == 0 {
			r = rt
		}
		if p.To.Reg == REGSP || r == REGSP {
			o2 = c.opxrrr(p, p.As, false)
			o2 |= REGTMP & 31 << 16
			o2 |= LSL0_64
		} else {
			o2 = c.oprrr(p, p.As)
			o2 |= REGTMP & 31 << 16
		}
		o2 |= uint32(r&31) << 5
		o2 |= uint32(rt & 31)

	case 64:
		o1 = ADR(1, 0, REGTMP)
		o2 = c.opirr(p, AADD) | REGTMP&31<<5 | REGTMP&31
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 8
		rel.Sym = p.To.Sym
		rel.Add = p.To.Offset
		rel.Type = objabi.R_ADDRARM64
		o3 = c.olsr12u(p, int32(c.opstr12(p, p.As)), 0, REGTMP, int(p.From.Reg))

	case 65:
		o1 = ADR(1, 0, REGTMP)
		o2 = c.opirr(p, AADD) | REGTMP&31<<5 | REGTMP&31
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 8
		rel.Sym = p.From.Sym
		rel.Add = p.From.Offset
		rel.Type = objabi.R_ADDRARM64
		o3 = c.olsr12u(p, int32(c.opldr12(p, p.As)), 0, REGTMP, int(p.To.Reg))

	case 66:
		v := int32(c.regoff(&p.From))
		r := int(p.From.Reg)
		if r == obj.REG_NONE {
			r = int(o.param)
		}
		if r == obj.REG_NONE {
			c.ctxt.Diag("invalid ldp source: %v\n", p)
		}
		o1 |= c.opldpstp(p, o, v, uint32(r), uint32(p.To.Reg), uint32(p.To.Offset), 1)

	case 67:
		r := int(p.To.Reg)
		if r == obj.REG_NONE {
			r = int(o.param)
		}
		if r == obj.REG_NONE {
			c.ctxt.Diag("invalid stp destination: %v\n", p)
		}
		v := int32(c.regoff(&p.To))
		o1 = c.opldpstp(p, o, v, uint32(r), uint32(p.From.Reg), uint32(p.From.Offset), 0)

	case 68:
		if p.As == AMOVW {
			c.ctxt.Diag("invalid load of 32-bit address: %v", p)
		}
		o1 = ADR(1, 0, uint32(p.To.Reg))
		o2 = c.opirr(p, AADD) | uint32(p.To.Reg&31)<<5 | uint32(p.To.Reg&31)
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 8
		rel.Sym = p.From.Sym
		rel.Add = p.From.Offset
		rel.Type = objabi.R_ADDRARM64

	case 69:
		o1 = c.opirr(p, AMOVZ)
		o1 |= uint32(p.To.Reg & 31)
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 4
		rel.Sym = p.From.Sym
		rel.Type = objabi.R_ARM64_TLS_LE
		if p.From.Offset != 0 {
			c.ctxt.Diag("invalid offset on MOVW $tlsvar")
		}

	case 70:
		o1 = ADR(1, 0, REGTMP)
		o2 = c.olsr12u(p, int32(c.opldr12(p, AMOVD)), 0, REGTMP, int(p.To.Reg))
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 8
		rel.Sym = p.From.Sym
		rel.Add = 0
		rel.Type = objabi.R_ARM64_TLS_IE
		if p.From.Offset != 0 {
			c.ctxt.Diag("invalid offset on MOVW $tlsvar")
		}

	case 71:
		o1 = ADR(1, 0, REGTMP)
		o2 = c.olsr12u(p, int32(c.opldr12(p, AMOVD)), 0, REGTMP, int(p.To.Reg))
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 8
		rel.Sym = p.From.Sym
		rel.Add = 0
		rel.Type = objabi.R_ARM64_GOTPCREL

	case 72:
		af := int((p.From.Reg >> 5) & 15)
		af3 := int((p.Reg >> 5) & 15)
		at := int((p.To.Reg >> 5) & 15)
		if af != af3 || af != at {
			c.ctxt.Diag("operand mismatch: %v", p)
			break
		}
		o1 = c.oprrr(p, p.As)
		rf := int((p.From.Reg) & 31)
		rt := int((p.To.Reg) & 31)
		r := int((p.Reg) & 31)

		Q := 0
		size := 0
		switch af {
		case ARNG_16B:
			Q = 1
			size = 0
		case ARNG_2D:
			Q = 1
			size = 3
		case ARNG_2S:
			Q = 0
			size = 2
		case ARNG_4H:
			Q = 0
			size = 1
		case ARNG_4S:
			Q = 1
			size = 2
		case ARNG_8B:
			Q = 0
			size = 0
		case ARNG_8H:
			Q = 1
			size = 1
		default:
			c.ctxt.Diag("invalid arrangement: %v", p)
		}

		if (p.As == AVORR || p.As == AVAND || p.As == AVEOR) &&
			(af != ARNG_16B && af != ARNG_8B) {
			c.ctxt.Diag("invalid arrangement: %v", p)
		} else if (p.As == AVFMLA || p.As == AVFMLS) &&
			(af != ARNG_2D && af != ARNG_2S && af != ARNG_4S) {
			c.ctxt.Diag("invalid arrangement: %v", p)
		} else if p.As == AVORR {
			size = 2
		} else if p.As == AVAND || p.As == AVEOR {
			size = 0
		} else if p.As == AVFMLA || p.As == AVFMLS {
			if af == ARNG_2D {
				size = 1
			} else {
				size = 0
			}
		}

		o1 |= (uint32(Q&1) << 30) | (uint32(size&3) << 22) | (uint32(rf&31) << 16) | (uint32(r&31) << 5) | uint32(rt&31)

	case 73:
		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		imm5 := 0
		o1 = 7<<25 | 0xf<<10
		index := int(p.From.Index)
		switch (p.From.Reg >> 5) & 15 {
		case ARNG_B:
			c.checkindex(p, index, 15)
			imm5 |= 1
			imm5 |= index << 1
		case ARNG_H:
			c.checkindex(p, index, 7)
			imm5 |= 2
			imm5 |= index << 2
		case ARNG_S:
			c.checkindex(p, index, 3)
			imm5 |= 4
			imm5 |= index << 3
		case ARNG_D:
			c.checkindex(p, index, 1)
			imm5 |= 8
			imm5 |= index << 4
			o1 |= 1 << 30
		default:
			c.ctxt.Diag("invalid arrangement: %v", p)
		}
		o1 |= (uint32(imm5&0x1f) << 16) | (uint32(rf&31) << 5) | uint32(rt&31)

	case 74:

		r := int(p.From.Reg)
		if r == obj.REG_NONE {
			r = int(o.param)
		}
		if r == obj.REG_NONE {
			c.ctxt.Diag("invalid ldp source: %v", p)
		}
		v := int32(c.regoff(&p.From))

		if v > 0 {
			if v > 4095 {
				c.ctxt.Diag("offset out of range: %v", p)
			}
			o1 = c.oaddi(p, int32(c.opirr(p, AADD)), v, r, REGTMP)
		}
		if v < 0 {
			if v < -4095 {
				c.ctxt.Diag("offset out of range: %v", p)
			}
			o1 = c.oaddi(p, int32(c.opirr(p, ASUB)), -v, r, REGTMP)
		}
		o2 |= c.opldpstp(p, o, 0, uint32(REGTMP), uint32(p.To.Reg), uint32(p.To.Offset), 1)

	case 75:

		r := int(p.From.Reg)
		if r == obj.REG_NONE {
			r = int(o.param)
		}
		if r == obj.REG_NONE {
			c.ctxt.Diag("invalid ldp source: %v", p)
		}
		o1 = c.omovlit(AMOVD, p, &p.From, REGTMP)
		o2 = c.opxrrr(p, AADD, false)
		o2 |= (REGTMP & 31) << 16
		o2 |= uint32(r&31) << 5
		o2 |= uint32(REGTMP & 31)
		o3 |= c.opldpstp(p, o, 0, uint32(REGTMP), uint32(p.To.Reg), uint32(p.To.Offset), 1)

	case 76:

		r := int(p.To.Reg)
		if r == obj.REG_NONE {
			r = int(o.param)
		}
		if r == obj.REG_NONE {
			c.ctxt.Diag("invalid stp destination: %v", p)
		}
		v := int32(c.regoff(&p.To))
		if v > 0 {
			if v > 4095 {
				c.ctxt.Diag("offset out of range: %v", p)
			}
			o1 = c.oaddi(p, int32(c.opirr(p, AADD)), v, r, REGTMP)
		}
		if v < 0 {
			if v < -4095 {
				c.ctxt.Diag("offset out of range: %v", p)
			}
			o1 = c.oaddi(p, int32(c.opirr(p, ASUB)), -v, r, REGTMP)
		}
		o2 |= c.opldpstp(p, o, 0, uint32(REGTMP), uint32(p.From.Reg), uint32(p.From.Offset), 0)

	case 77:

		r := int(p.To.Reg)
		if r == obj.REG_NONE {
			r = int(o.param)
		}
		if r == obj.REG_NONE {
			c.ctxt.Diag("invalid stp destination: %v", p)
		}
		o1 = c.omovlit(AMOVD, p, &p.To, REGTMP)
		o2 = c.opxrrr(p, AADD, false)
		o2 |= REGTMP & 31 << 16
		o2 |= uint32(r&31) << 5
		o2 |= uint32(REGTMP & 31)
		o3 |= c.opldpstp(p, o, 0, uint32(REGTMP), uint32(p.From.Reg), uint32(p.From.Offset), 0)

	case 78:
		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		imm5 := 0
		o1 = 1<<30 | 7<<25 | 7<<10
		index := int(p.To.Index)
		switch (p.To.Reg >> 5) & 15 {
		case ARNG_B:
			c.checkindex(p, index, 15)
			imm5 |= 1
			imm5 |= index << 1
		case ARNG_H:
			c.checkindex(p, index, 7)
			imm5 |= 2
			imm5 |= index << 2
		case ARNG_S:
			c.checkindex(p, index, 3)
			imm5 |= 4
			imm5 |= index << 3
		case ARNG_D:
			c.checkindex(p, index, 1)
			imm5 |= 8
			imm5 |= index << 4
		default:
			c.ctxt.Diag("invalid arrangement: %v", p)
		}
		o1 |= (uint32(imm5&0x1f) << 16) | (uint32(rf&31) << 5) | uint32(rt&31)

	case 79:
		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		o1 = 7<<25 | 1<<10
		var imm5, Q int
		index := int(p.From.Index)
		switch (p.To.Reg >> 5) & 15 {
		case ARNG_16B:
			c.checkindex(p, index, 15)
			Q = 1
			imm5 = 1
			imm5 |= index << 1
		case ARNG_2D:
			c.checkindex(p, index, 1)
			Q = 1
			imm5 = 8
			imm5 |= index << 4
		case ARNG_2S:
			c.checkindex(p, index, 3)
			Q = 0
			imm5 = 4
			imm5 |= index << 3
		case ARNG_4H:
			c.checkindex(p, index, 7)
			Q = 0
			imm5 = 2
			imm5 |= index << 2
		case ARNG_4S:
			c.checkindex(p, index, 3)
			Q = 1
			imm5 = 4
			imm5 |= index << 3
		case ARNG_8B:
			c.checkindex(p, index, 15)
			Q = 0
			imm5 = 1
			imm5 |= index << 1
		case ARNG_8H:
			c.checkindex(p, index, 7)
			Q = 1
			imm5 = 2
			imm5 |= index << 2
		default:
			c.ctxt.Diag("invalid arrangement: %v", p)
		}
		o1 |= (uint32(Q&1) << 30) | (uint32(imm5&0x1f) << 16)
		o1 |= (uint32(rf&31) << 5) | uint32(rt&31)

	case 80:
		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		imm5 := 0
		index := int(p.From.Index)
		switch p.As {
		case AVMOV:
			o1 = 1<<30 | 15<<25 | 1<<10
			switch (p.From.Reg >> 5) & 15 {
			case ARNG_B:
				c.checkindex(p, index, 15)
				imm5 |= 1
				imm5 |= index << 1
			case ARNG_H:
				c.checkindex(p, index, 7)
				imm5 |= 2
				imm5 |= index << 2
			case ARNG_S:
				c.checkindex(p, index, 3)
				imm5 |= 4
				imm5 |= index << 3
			case ARNG_D:
				c.checkindex(p, index, 1)
				imm5 |= 8
				imm5 |= index << 4
			default:
				c.ctxt.Diag("invalid arrangement: %v", p)
			}
		default:
			c.ctxt.Diag("unsupported op %v", p.As)
		}
		o1 |= (uint32(imm5&0x1f) << 16) | (uint32(rf&31) << 5) | uint32(rt&31)

	case 81:
		r := int(p.From.Reg)
		o1 = 3<<26 | 1<<22
		if o.scond == C_XPOST {
			o1 |= 1 << 23
			if p.From.Index == 0 {

				c.checkoffset(p, p.As)
				o1 |= 0x1f << 16
			} else {

				if isRegShiftOrExt(&p.From) {
					c.ctxt.Diag("invalid extended register op: %v\n", p)
				}
				o1 |= uint32(p.From.Index&31) << 16
			}
		}
		o1 |= uint32(p.To.Offset)
		o1 |= uint32(r&31) << 5

	case 82:
		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		o1 = 7<<25 | 3<<10
		var imm5, Q uint32
		switch (p.To.Reg >> 5) & 15 {
		case ARNG_16B:
			Q = 1
			imm5 = 1
		case ARNG_2D:
			Q = 1
			imm5 = 8
		case ARNG_2S:
			Q = 0
			imm5 = 4
		case ARNG_4H:
			Q = 0
			imm5 = 2
		case ARNG_4S:
			Q = 1
			imm5 = 4
		case ARNG_8B:
			Q = 0
			imm5 = 1
		case ARNG_8H:
			Q = 1
			imm5 = 2
		default:
			c.ctxt.Diag("invalid arrangement on VMOV Rn, Vd.<T>: %v\n", p)
		}
		o1 |= (Q & 1 << 30) | (imm5 & 0x1f << 16)
		o1 |= (uint32(rf&31) << 5) | uint32(rt&31)

	case 83:
		af := int((p.From.Reg >> 5) & 15)
		at := int((p.To.Reg >> 5) & 15)
		if af != at {
			c.ctxt.Diag("invalid arrangement: %v\n", p)
		}
		o1 = c.oprrr(p, p.As)
		rf := int((p.From.Reg) & 31)
		rt := int((p.To.Reg) & 31)

		var Q, size uint32
		switch af {
		case ARNG_8B:
			Q = 0
			size = 0
		case ARNG_16B:
			Q = 1
			size = 0
		case ARNG_4H:
			Q = 0
			size = 1
		case ARNG_8H:
			Q = 1
			size = 1
		case ARNG_2S:
			Q = 0
			size = 2
		case ARNG_4S:
			Q = 1
			size = 2
		default:
			c.ctxt.Diag("invalid arrangement: %v\n", p)
		}

		if (p.As == AVMOV || p.As == AVRBIT) && (af != ARNG_16B && af != ARNG_8B) {
			c.ctxt.Diag("invalid arrangement: %v", p)
		}

		if p.As == AVREV32 && (af == ARNG_2S || af == ARNG_4S) {
			c.ctxt.Diag("invalid arrangement: %v", p)
		}

		if p.As == AVMOV {
			o1 |= uint32(rf&31) << 16
		}

		if p.As == AVRBIT {
			size = 1
		}

		o1 |= (Q&1)<<30 | (size&3)<<22 | uint32(rf&31)<<5 | uint32(rt&31)

	case 84:
		r := int(p.To.Reg)
		o1 = 3 << 26
		if o.scond == C_XPOST {
			o1 |= 1 << 23
			if p.To.Index == 0 {

				c.checkoffset(p, p.As)
				o1 |= 0x1f << 16
			} else {

				if isRegShiftOrExt(&p.To) {
					c.ctxt.Diag("invalid extended register: %v\n", p)
				}
				o1 |= uint32(p.To.Index&31) << 16
			}
		}
		o1 |= uint32(p.From.Offset)
		o1 |= uint32(r&31) << 5

	case 85:
		af := int((p.From.Reg >> 5) & 15)
		o1 = c.oprrr(p, p.As)
		rf := int((p.From.Reg) & 31)
		rt := int((p.To.Reg) & 31)
		Q := 0
		size := 0
		switch af {
		case ARNG_8B:
			Q = 0
			size = 0
		case ARNG_16B:
			Q = 1
			size = 0
		case ARNG_4H:
			Q = 0
			size = 1
		case ARNG_8H:
			Q = 1
			size = 1
		case ARNG_4S:
			Q = 1
			size = 2
		default:
			c.ctxt.Diag("invalid arrangement: %v\n", p)
		}
		o1 |= (uint32(Q&1) << 30) | (uint32(size&3) << 22) | (uint32(rf&31) << 5) | uint32(rt&31)

	case 86:
		at := int((p.To.Reg >> 5) & 15)
		r := int(p.From.Offset)
		if r > 255 || r < 0 {
			c.ctxt.Diag("immediate constant out of range: %v\n", p)
		}
		rt := int((p.To.Reg) & 31)
		Q := 0
		switch at {
		case ARNG_8B:
			Q = 0
		case ARNG_16B:
			Q = 1
		default:
			c.ctxt.Diag("invalid arrangement: %v\n", p)
		}
		o1 = 0xf<<24 | 0xe<<12 | 1<<10
		o1 |= (uint32(Q&1) << 30) | (uint32((r>>5)&7) << 16) | (uint32(r&0x1f) << 5) | uint32(rt&31)

	case 87:
		o1 = ADR(1, 0, REGTMP)
		o2 = c.opirr(p, AADD) | REGTMP&31<<5 | REGTMP&31
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 8
		rel.Sym = p.To.Sym
		rel.Add = p.To.Offset
		rel.Type = objabi.R_ADDRARM64
		o3 |= c.opldpstp(p, o, 0, uint32(REGTMP), uint32(p.From.Reg), uint32(p.From.Offset), 0)

	case 88:
		o1 = ADR(1, 0, REGTMP)
		o2 = c.opirr(p, AADD) | REGTMP&31<<5 | REGTMP&31
		rel := obj.Addrel(c.cursym)
		rel.Off = int32(c.pc)
		rel.Siz = 8
		rel.Sym = p.From.Sym
		rel.Add = p.From.Offset
		rel.Type = objabi.R_ADDRARM64
		o3 |= c.opldpstp(p, o, 0, uint32(REGTMP), uint32(p.To.Reg), uint32(p.To.Offset), 1)

	case 89:
		switch p.As {
		case AVADD:
			o1 = 5<<28 | 7<<25 | 7<<21 | 1<<15 | 1<<10

		case AVSUB:
			o1 = 7<<28 | 7<<25 | 7<<21 | 1<<15 | 1<<10

		default:
			c.ctxt.Diag("bad opcode: %v\n", p)
			break
		}

		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		r := int(p.Reg)
		if r == 0 {
			r = rt
		}
		o1 |= (uint32(rf&31) << 16) | (uint32(r&31) << 5) | uint32(rt&31)

	case 90:
		o1 = 0xbea71700

	case 91:
		imm := uint32(p.From.Offset)
		r := p.From.Reg
		v := uint32(0xff)
		if p.To.Type == obj.TYPE_CONST {
			v = uint32(p.To.Offset)
			if v > 31 {
				c.ctxt.Diag("illegal prefetch operation\n%v", p)
			}
		} else {
			for i := 0; i < len(psess.prfopfield); i++ {
				if psess.prfopfield[i].reg == p.To.Reg {
					v = psess.prfopfield[i].enc
					break
				}
			}
			if v == 0xff {
				c.ctxt.Diag("illegal prefetch operation:\n%v", p)
			}
		}

		o1 = c.opldrpp(p, p.As)
		o1 |= (uint32(r&31) << 5) | (uint32((imm>>3)&0xfff) << 10) | (uint32(v & 31))

	case 92:
		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		imm4 := 0
		imm5 := 0
		o1 = 3<<29 | 7<<25 | 1<<10
		index1 := int(p.To.Index)
		index2 := int(p.From.Index)
		if ((p.To.Reg >> 5) & 15) != ((p.From.Reg >> 5) & 15) {
			c.ctxt.Diag("operand mismatch: %v", p)
		}
		switch (p.To.Reg >> 5) & 15 {
		case ARNG_B:
			c.checkindex(p, index1, 15)
			c.checkindex(p, index2, 15)
			imm5 |= 1
			imm5 |= index1 << 1
			imm4 |= index2
		case ARNG_H:
			c.checkindex(p, index1, 7)
			c.checkindex(p, index2, 7)
			imm5 |= 2
			imm5 |= index1 << 2
			imm4 |= index2 << 1
		case ARNG_S:
			c.checkindex(p, index1, 3)
			c.checkindex(p, index2, 3)
			imm5 |= 4
			imm5 |= index1 << 3
			imm4 |= index2 << 2
		case ARNG_D:
			c.checkindex(p, index1, 1)
			c.checkindex(p, index2, 1)
			imm5 |= 8
			imm5 |= index1 << 4
			imm4 |= index2 << 3
		default:
			c.ctxt.Diag("invalid arrangement: %v", p)
		}
		o1 |= (uint32(imm5&0x1f) << 16) | (uint32(imm4&0xf) << 11) | (uint32(rf&31) << 5) | uint32(rt&31)

	case 93:
		af := int((p.From.Reg >> 5) & 15)
		at := int((p.To.Reg >> 5) & 15)
		a := int((p.Reg >> 5) & 15)

		var Q, size uint32
		if p.As == AVPMULL {
			Q = 0
		} else {
			Q = 1
		}

		var fArng int
		switch at {
		case ARNG_8H:
			if Q == 0 {
				fArng = ARNG_8B
			} else {
				fArng = ARNG_16B
			}
			size = 0
		case ARNG_1Q:
			if Q == 0 {
				fArng = ARNG_1D
			} else {
				fArng = ARNG_2D
			}
			size = 3
		default:
			c.ctxt.Diag("invalid arrangement on Vd.<T>: %v", p)
		}

		if af != a || af != fArng {
			c.ctxt.Diag("invalid arrangement: %v", p)
		}

		o1 = c.oprrr(p, p.As)
		rf := int((p.From.Reg) & 31)
		rt := int((p.To.Reg) & 31)
		r := int((p.Reg) & 31)

		o1 |= ((Q & 1) << 30) | ((size & 3) << 22) | (uint32(rf&31) << 16) | (uint32(r&31) << 5) | uint32(rt&31)

	case 94:
		af := int(((p.GetFrom3().Reg) >> 5) & 15)
		at := int((p.To.Reg >> 5) & 15)
		a := int((p.Reg >> 5) & 15)
		index := int(p.From.Offset)

		if af != a || af != at {
			c.ctxt.Diag("invalid arrangement: %v", p)
			break
		}

		var Q uint32
		var b int
		if af == ARNG_8B {
			Q = 0
			b = 7
		} else if af == ARNG_16B {
			Q = 1
			b = 15
		} else {
			c.ctxt.Diag("invalid arrangement, should be 8B or 16B: %v", p)
			break
		}

		if index < 0 || index > b {
			c.ctxt.Diag("illegal offset: %v", p)
		}

		o1 = c.opirr(p, p.As)
		rf := int((p.GetFrom3().Reg) & 31)
		rt := int((p.To.Reg) & 31)
		r := int((p.Reg) & 31)

		o1 |= ((Q & 1) << 30) | (uint32(r&31) << 16) | (uint32(index&15) << 11) | (uint32(rf&31) << 5) | uint32(rt&31)

	case 95:
		at := int((p.To.Reg >> 5) & 15)
		af := int((p.Reg >> 5) & 15)
		shift := int(p.From.Offset)

		if af != at {
			c.ctxt.Diag("invalid arrangement on op Vn.<T>, Vd.<T>: %v", p)
		}

		var Q uint32
		var imax, esize int

		switch af {
		case ARNG_8B, ARNG_4H, ARNG_2S:
			Q = 0
		case ARNG_16B, ARNG_8H, ARNG_4S, ARNG_2D:
			Q = 1
		default:
			c.ctxt.Diag("invalid arrangement on op Vn.<T>, Vd.<T>: %v", p)
		}

		switch af {
		case ARNG_8B, ARNG_16B:
			imax = 15
			esize = 8
		case ARNG_4H, ARNG_8H:
			imax = 31
			esize = 16
		case ARNG_2S, ARNG_4S:
			imax = 63
			esize = 32
		case ARNG_2D:
			imax = 127
			esize = 64
		}

		imm := 0

		switch p.As {
		case AVUSHR, AVSRI:
			imm = esize*2 - shift
			if imm < esize || imm > imax {
				c.ctxt.Diag("shift out of range: %v", p)
			}
		case AVSHL:
			imm = esize + shift
			if imm > imax {
				c.ctxt.Diag("shift out of range: %v", p)
			}
		default:
			c.ctxt.Diag("invalid instruction %v\n", p)
		}

		o1 = c.opirr(p, p.As)
		rt := int((p.To.Reg) & 31)
		rf := int((p.Reg) & 31)

		o1 |= ((Q & 1) << 30) | (uint32(imm&127) << 16) | (uint32(rf&31) << 5) | uint32(rt&31)

	case 96:
		af := int((p.From.Reg >> 5) & 15)
		rt := int((p.From.Reg) & 31)
		rf := int((p.To.Reg) & 31)
		r := int(p.To.Index & 31)
		index := int(p.From.Index)
		offset := int32(c.regoff(&p.To))

		if o.scond == C_XPOST {
			if (p.To.Index != 0) && (offset != 0) {
				c.ctxt.Diag("invalid offset: %v", p)
			}
			if p.To.Index == 0 && offset == 0 {
				c.ctxt.Diag("invalid offset: %v", p)
			}
		}

		if offset != 0 {
			r = 31
		}

		var Q, S, size int
		var opcode uint32
		switch af {
		case ARNG_B:
			c.checkindex(p, index, 15)
			if o.scond == C_XPOST && offset != 0 && offset != 1 {
				c.ctxt.Diag("invalid offset: %v", p)
			}
			Q = index >> 3
			S = (index >> 2) & 1
			size = index & 3
			opcode = 0
		case ARNG_H:
			c.checkindex(p, index, 7)
			if o.scond == C_XPOST && offset != 0 && offset != 2 {
				c.ctxt.Diag("invalid offset: %v", p)
			}
			Q = index >> 2
			S = (index >> 1) & 1
			size = (index & 1) << 1
			opcode = 2
		case ARNG_S:
			c.checkindex(p, index, 3)
			if o.scond == C_XPOST && offset != 0 && offset != 4 {
				c.ctxt.Diag("invalid offset: %v", p)
			}
			Q = index >> 1
			S = index & 1
			size = 0
			opcode = 4
		case ARNG_D:
			c.checkindex(p, index, 1)
			if o.scond == C_XPOST && offset != 0 && offset != 8 {
				c.ctxt.Diag("invalid offset: %v", p)
			}
			Q = index
			S = 0
			size = 1
			opcode = 4
		default:
			c.ctxt.Diag("invalid arrangement: %v", p)
		}

		if o.scond == C_XPOST {
			o1 |= 27 << 23
		} else {
			o1 |= 26 << 23
		}

		o1 |= (uint32(Q&1) << 30) | (uint32(r&31) << 16) | ((opcode & 7) << 13) | (uint32(S&1) << 12) | (uint32(size&3) << 10) | (uint32(rf&31) << 5) | uint32(rt&31)

	case 97:
		at := int((p.To.Reg >> 5) & 15)
		rt := int((p.To.Reg) & 31)
		rf := int((p.From.Reg) & 31)
		r := int(p.From.Index & 31)
		index := int(p.To.Index)
		offset := int32(c.regoff(&p.From))

		if o.scond == C_XPOST {
			if (p.From.Index != 0) && (offset != 0) {
				c.ctxt.Diag("invalid offset: %v", p)
			}
			if p.From.Index == 0 && offset == 0 {
				c.ctxt.Diag("invalid offset: %v", p)
			}
		}

		if offset != 0 {
			r = 31
		}

		Q := 0
		S := 0
		size := 0
		var opcode uint32
		switch at {
		case ARNG_B:
			c.checkindex(p, index, 15)
			if o.scond == C_XPOST && offset != 0 && offset != 1 {
				c.ctxt.Diag("invalid offset: %v", p)
			}
			Q = index >> 3
			S = (index >> 2) & 1
			size = index & 3
			opcode = 0
		case ARNG_H:
			c.checkindex(p, index, 7)
			if o.scond == C_XPOST && offset != 0 && offset != 2 {
				c.ctxt.Diag("invalid offset: %v", p)
			}
			Q = index >> 2
			S = (index >> 1) & 1
			size = (index & 1) << 1
			opcode = 2
		case ARNG_S:
			c.checkindex(p, index, 3)
			if o.scond == C_XPOST && offset != 0 && offset != 4 {
				c.ctxt.Diag("invalid offset: %v", p)
			}
			Q = index >> 1
			S = index & 1
			size = 0
			opcode = 4
		case ARNG_D:
			c.checkindex(p, index, 1)
			if o.scond == C_XPOST && offset != 0 && offset != 8 {
				c.ctxt.Diag("invalid offset: %v", p)
			}
			Q = index
			S = 0
			size = 1
			opcode = 4
		default:
			c.ctxt.Diag("invalid arrangement: %v", p)
		}

		if o.scond == C_XPOST {
			o1 |= 110 << 21
		} else {
			o1 |= 106 << 21
		}

		o1 |= (uint32(Q&1) << 30) | (uint32(r&31) << 16) | ((opcode & 7) << 13) | (uint32(S&1) << 12) | (uint32(size&3) << 10) | (uint32(rf&31) << 5) | uint32(rt&31)

	case 98:
		if isRegShiftOrExt(&p.From) {

			c.checkShiftAmount(p, &p.From)

			o1 = c.opldrr(p, p.As, true)
			o1 |= c.encRegShiftOrExt(&p.From, p.From.Index)
		} else {

			o1 = c.opldrr(p, p.As, false)
			o1 |= uint32(p.From.Index&31) << 16
		}
		o1 |= uint32(p.From.Reg&31) << 5
		rt := int(p.To.Reg)
		o1 |= uint32(rt & 31)

	case 99:
		if isRegShiftOrExt(&p.To) {

			c.checkShiftAmount(p, &p.To)

			o1 = c.opstrr(p, p.As, true)
			o1 |= c.encRegShiftOrExt(&p.To, p.To.Index)
		} else {

			o1 = c.opstrr(p, p.As, false)
			o1 |= uint32(p.To.Index&31) << 16
		}
		o1 |= uint32(p.To.Reg&31) << 5
		rf := int(p.From.Reg)
		o1 |= uint32(rf & 31)

	case 100:
		af := int((p.From.Reg >> 5) & 15)
		at := int((p.To.Reg >> 5) & 15)
		if af != at {
			c.ctxt.Diag("invalid arrangement: %v\n", p)
		}
		var q, len uint32
		switch af {
		case ARNG_8B:
			q = 0
		case ARNG_16B:
			q = 1
		default:
			c.ctxt.Diag("invalid arrangement: %v", p)
		}
		rf := int(p.From.Reg)
		rt := int(p.To.Reg)
		offset := int(p.GetFrom3().Offset)
		opcode := (offset >> 12) & 15
		switch opcode {
		case 0x7:
			len = 0
		case 0xa:
			len = 1
		case 0x6:
			len = 2
		case 0x2:
			len = 3
		default:
			c.ctxt.Diag("invalid register numbers in ARM64 register list: %v", p)
		}
		o1 = q<<30 | 0xe<<24 | len<<13
		o1 |= (uint32(rf&31) << 16) | uint32(offset&31)<<5 | uint32(rt&31)

	}
	out[0] = o1
	out[1] = o2
	out[2] = o3
	out[3] = o4
	out[4] = o5
}

/*
 * basic Rm op Rn -> Rd (using shifted register with 0)
 * also op Rn -> Rt
 * also Rm*Rn op Ra -> Rd
 * also Vm op Vn -> Vd
 */
func (c *ctxt7) oprrr(p *obj.Prog, a obj.As) uint32 {
	switch a {
	case AADC:
		return S64 | 0<<30 | 0<<29 | 0xd0<<21 | 0<<10

	case AADCW:
		return S32 | 0<<30 | 0<<29 | 0xd0<<21 | 0<<10

	case AADCS:
		return S64 | 0<<30 | 1<<29 | 0xd0<<21 | 0<<10

	case AADCSW:
		return S32 | 0<<30 | 1<<29 | 0xd0<<21 | 0<<10

	case ANGC, ASBC:
		return S64 | 1<<30 | 0<<29 | 0xd0<<21 | 0<<10

	case ANGCS, ASBCS:
		return S64 | 1<<30 | 1<<29 | 0xd0<<21 | 0<<10

	case ANGCW, ASBCW:
		return S32 | 1<<30 | 0<<29 | 0xd0<<21 | 0<<10

	case ANGCSW, ASBCSW:
		return S32 | 1<<30 | 1<<29 | 0xd0<<21 | 0<<10

	case AADD:
		return S64 | 0<<30 | 0<<29 | 0x0b<<24 | 0<<22 | 0<<21 | 0<<10

	case AADDW:
		return S32 | 0<<30 | 0<<29 | 0x0b<<24 | 0<<22 | 0<<21 | 0<<10

	case ACMN, AADDS:
		return S64 | 0<<30 | 1<<29 | 0x0b<<24 | 0<<22 | 0<<21 | 0<<10

	case ACMNW, AADDSW:
		return S32 | 0<<30 | 1<<29 | 0x0b<<24 | 0<<22 | 0<<21 | 0<<10

	case ASUB:
		return S64 | 1<<30 | 0<<29 | 0x0b<<24 | 0<<22 | 0<<21 | 0<<10

	case ASUBW:
		return S32 | 1<<30 | 0<<29 | 0x0b<<24 | 0<<22 | 0<<21 | 0<<10

	case ACMP, ASUBS:
		return S64 | 1<<30 | 1<<29 | 0x0b<<24 | 0<<22 | 0<<21 | 0<<10

	case ACMPW, ASUBSW:
		return S32 | 1<<30 | 1<<29 | 0x0b<<24 | 0<<22 | 0<<21 | 0<<10

	case AAND:
		return S64 | 0<<29 | 0xA<<24

	case AANDW:
		return S32 | 0<<29 | 0xA<<24

	case AMOVD, AORR:
		return S64 | 1<<29 | 0xA<<24

	case AMOVWU, AORRW:
		return S32 | 1<<29 | 0xA<<24

	case AEOR:
		return S64 | 2<<29 | 0xA<<24

	case AEORW:
		return S32 | 2<<29 | 0xA<<24

	case AANDS, ATST:
		return S64 | 3<<29 | 0xA<<24

	case AANDSW, ATSTW:
		return S32 | 3<<29 | 0xA<<24

	case ABIC:
		return S64 | 0<<29 | 0xA<<24 | 1<<21

	case ABICW:
		return S32 | 0<<29 | 0xA<<24 | 1<<21

	case ABICS:
		return S64 | 3<<29 | 0xA<<24 | 1<<21

	case ABICSW:
		return S32 | 3<<29 | 0xA<<24 | 1<<21

	case AEON:
		return S64 | 2<<29 | 0xA<<24 | 1<<21

	case AEONW:
		return S32 | 2<<29 | 0xA<<24 | 1<<21

	case AMVN, AORN:
		return S64 | 1<<29 | 0xA<<24 | 1<<21

	case AMVNW, AORNW:
		return S32 | 1<<29 | 0xA<<24 | 1<<21

	case AASR:
		return S64 | OPDP2(10)

	case AASRW:
		return S32 | OPDP2(10)

	case ALSL:
		return S64 | OPDP2(8)

	case ALSLW:
		return S32 | OPDP2(8)

	case ALSR:
		return S64 | OPDP2(9)

	case ALSRW:
		return S32 | OPDP2(9)

	case AROR:
		return S64 | OPDP2(11)

	case ARORW:
		return S32 | OPDP2(11)

	case ACCMN:
		return S64 | 0<<30 | 1<<29 | 0xD2<<21 | 0<<11 | 0<<10 | 0<<4

	case ACCMNW:
		return S32 | 0<<30 | 1<<29 | 0xD2<<21 | 0<<11 | 0<<10 | 0<<4

	case ACCMP:
		return S64 | 1<<30 | 1<<29 | 0xD2<<21 | 0<<11 | 0<<10 | 0<<4

	case ACCMPW:
		return S32 | 1<<30 | 1<<29 | 0xD2<<21 | 0<<11 | 0<<10 | 0<<4

	case ACRC32B:
		return S32 | OPDP2(16)

	case ACRC32H:
		return S32 | OPDP2(17)

	case ACRC32W:
		return S32 | OPDP2(18)

	case ACRC32X:
		return S64 | OPDP2(19)

	case ACRC32CB:
		return S32 | OPDP2(20)

	case ACRC32CH:
		return S32 | OPDP2(21)

	case ACRC32CW:
		return S32 | OPDP2(22)

	case ACRC32CX:
		return S64 | OPDP2(23)

	case ACSEL:
		return S64 | 0<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 0<<10

	case ACSELW:
		return S32 | 0<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 0<<10

	case ACSET:
		return S64 | 0<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 1<<10

	case ACSETW:
		return S32 | 0<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 1<<10

	case ACSETM:
		return S64 | 1<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 0<<10

	case ACSETMW:
		return S32 | 1<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 0<<10

	case ACINC, ACSINC:
		return S64 | 0<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 1<<10

	case ACINCW, ACSINCW:
		return S32 | 0<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 1<<10

	case ACINV, ACSINV:
		return S64 | 1<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 0<<10

	case ACINVW, ACSINVW:
		return S32 | 1<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 0<<10

	case ACNEG, ACSNEG:
		return S64 | 1<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 1<<10

	case ACNEGW, ACSNEGW:
		return S32 | 1<<30 | 0<<29 | 0xD4<<21 | 0<<11 | 1<<10

	case AMUL, AMADD:
		return S64 | 0<<29 | 0x1B<<24 | 0<<21 | 0<<15

	case AMULW, AMADDW:
		return S32 | 0<<29 | 0x1B<<24 | 0<<21 | 0<<15

	case AMNEG, AMSUB:
		return S64 | 0<<29 | 0x1B<<24 | 0<<21 | 1<<15

	case AMNEGW, AMSUBW:
		return S32 | 0<<29 | 0x1B<<24 | 0<<21 | 1<<15

	case AMRS:
		return SYSOP(1, 2, 0, 0, 0, 0, 0)

	case AMSR:
		return SYSOP(0, 2, 0, 0, 0, 0, 0)

	case ANEG:
		return S64 | 1<<30 | 0<<29 | 0xB<<24 | 0<<21

	case ANEGW:
		return S32 | 1<<30 | 0<<29 | 0xB<<24 | 0<<21

	case ANEGS:
		return S64 | 1<<30 | 1<<29 | 0xB<<24 | 0<<21

	case ANEGSW:
		return S32 | 1<<30 | 1<<29 | 0xB<<24 | 0<<21

	case AREM, ASDIV:
		return S64 | OPDP2(3)

	case AREMW, ASDIVW:
		return S32 | OPDP2(3)

	case ASMULL, ASMADDL:
		return OPDP3(1, 0, 1, 0)

	case ASMNEGL, ASMSUBL:
		return OPDP3(1, 0, 1, 1)

	case ASMULH:
		return OPDP3(1, 0, 2, 0)

	case AUMULL, AUMADDL:
		return OPDP3(1, 0, 5, 0)

	case AUMNEGL, AUMSUBL:
		return OPDP3(1, 0, 5, 1)

	case AUMULH:
		return OPDP3(1, 0, 6, 0)

	case AUREM, AUDIV:
		return S64 | OPDP2(2)

	case AUREMW, AUDIVW:
		return S32 | OPDP2(2)

	case AAESE:
		return 0x4E<<24 | 2<<20 | 8<<16 | 4<<12 | 2<<10

	case AAESD:
		return 0x4E<<24 | 2<<20 | 8<<16 | 5<<12 | 2<<10

	case AAESMC:
		return 0x4E<<24 | 2<<20 | 8<<16 | 6<<12 | 2<<10

	case AAESIMC:
		return 0x4E<<24 | 2<<20 | 8<<16 | 7<<12 | 2<<10

	case ASHA1C:
		return 0x5E<<24 | 0<<12

	case ASHA1P:
		return 0x5E<<24 | 1<<12

	case ASHA1M:
		return 0x5E<<24 | 2<<12

	case ASHA1SU0:
		return 0x5E<<24 | 3<<12

	case ASHA256H:
		return 0x5E<<24 | 4<<12

	case ASHA256H2:
		return 0x5E<<24 | 5<<12

	case ASHA256SU1:
		return 0x5E<<24 | 6<<12

	case ASHA1H:
		return 0x5E<<24 | 2<<20 | 8<<16 | 0<<12 | 2<<10

	case ASHA1SU1:
		return 0x5E<<24 | 2<<20 | 8<<16 | 1<<12 | 2<<10

	case ASHA256SU0:
		return 0x5E<<24 | 2<<20 | 8<<16 | 2<<12 | 2<<10

	case AFCVTZSD:
		return FPCVTI(1, 0, 1, 3, 0)

	case AFCVTZSDW:
		return FPCVTI(0, 0, 1, 3, 0)

	case AFCVTZSS:
		return FPCVTI(1, 0, 0, 3, 0)

	case AFCVTZSSW:
		return FPCVTI(0, 0, 0, 3, 0)

	case AFCVTZUD:
		return FPCVTI(1, 0, 1, 3, 1)

	case AFCVTZUDW:
		return FPCVTI(0, 0, 1, 3, 1)

	case AFCVTZUS:
		return FPCVTI(1, 0, 0, 3, 1)

	case AFCVTZUSW:
		return FPCVTI(0, 0, 0, 3, 1)

	case ASCVTFD:
		return FPCVTI(1, 0, 1, 0, 2)

	case ASCVTFS:
		return FPCVTI(1, 0, 0, 0, 2)

	case ASCVTFWD:
		return FPCVTI(0, 0, 1, 0, 2)

	case ASCVTFWS:
		return FPCVTI(0, 0, 0, 0, 2)

	case AUCVTFD:
		return FPCVTI(1, 0, 1, 0, 3)

	case AUCVTFS:
		return FPCVTI(1, 0, 0, 0, 3)

	case AUCVTFWD:
		return FPCVTI(0, 0, 1, 0, 3)

	case AUCVTFWS:
		return FPCVTI(0, 0, 0, 0, 3)

	case AFADDS:
		return FPOP2S(0, 0, 0, 2)

	case AFADDD:
		return FPOP2S(0, 0, 1, 2)

	case AFSUBS:
		return FPOP2S(0, 0, 0, 3)

	case AFSUBD:
		return FPOP2S(0, 0, 1, 3)

	case AFMADDD:
		return FPOP3S(0, 0, 1, 0, 0)

	case AFMADDS:
		return FPOP3S(0, 0, 0, 0, 0)

	case AFMSUBD:
		return FPOP3S(0, 0, 1, 0, 1)

	case AFMSUBS:
		return FPOP3S(0, 0, 0, 0, 1)

	case AFNMADDD:
		return FPOP3S(0, 0, 1, 1, 0)

	case AFNMADDS:
		return FPOP3S(0, 0, 0, 1, 0)

	case AFNMSUBD:
		return FPOP3S(0, 0, 1, 1, 1)

	case AFNMSUBS:
		return FPOP3S(0, 0, 0, 1, 1)

	case AFMULS:
		return FPOP2S(0, 0, 0, 0)

	case AFMULD:
		return FPOP2S(0, 0, 1, 0)

	case AFDIVS:
		return FPOP2S(0, 0, 0, 1)

	case AFDIVD:
		return FPOP2S(0, 0, 1, 1)

	case AFMAXS:
		return FPOP2S(0, 0, 0, 4)

	case AFMINS:
		return FPOP2S(0, 0, 0, 5)

	case AFMAXD:
		return FPOP2S(0, 0, 1, 4)

	case AFMIND:
		return FPOP2S(0, 0, 1, 5)

	case AFMAXNMS:
		return FPOP2S(0, 0, 0, 6)

	case AFMAXNMD:
		return FPOP2S(0, 0, 1, 6)

	case AFMINNMS:
		return FPOP2S(0, 0, 0, 7)

	case AFMINNMD:
		return FPOP2S(0, 0, 1, 7)

	case AFNMULS:
		return FPOP2S(0, 0, 0, 8)

	case AFNMULD:
		return FPOP2S(0, 0, 1, 8)

	case AFCMPS:
		return FPCMP(0, 0, 0, 0, 0)

	case AFCMPD:
		return FPCMP(0, 0, 1, 0, 0)

	case AFCMPES:
		return FPCMP(0, 0, 0, 0, 16)

	case AFCMPED:
		return FPCMP(0, 0, 1, 0, 16)

	case AFCCMPS:
		return FPCCMP(0, 0, 0, 0)

	case AFCCMPD:
		return FPCCMP(0, 0, 1, 0)

	case AFCCMPES:
		return FPCCMP(0, 0, 0, 1)

	case AFCCMPED:
		return FPCCMP(0, 0, 1, 1)

	case AFCSELS:
		return 0x1E<<24 | 0<<22 | 1<<21 | 3<<10

	case AFCSELD:
		return 0x1E<<24 | 1<<22 | 1<<21 | 3<<10

	case AFMOVS:
		return FPOP1S(0, 0, 0, 0)

	case AFABSS:
		return FPOP1S(0, 0, 0, 1)

	case AFNEGS:
		return FPOP1S(0, 0, 0, 2)

	case AFSQRTS:
		return FPOP1S(0, 0, 0, 3)

	case AFCVTSD:
		return FPOP1S(0, 0, 0, 5)

	case AFCVTSH:
		return FPOP1S(0, 0, 0, 7)

	case AFRINTNS:
		return FPOP1S(0, 0, 0, 8)

	case AFRINTPS:
		return FPOP1S(0, 0, 0, 9)

	case AFRINTMS:
		return FPOP1S(0, 0, 0, 10)

	case AFRINTZS:
		return FPOP1S(0, 0, 0, 11)

	case AFRINTAS:
		return FPOP1S(0, 0, 0, 12)

	case AFRINTXS:
		return FPOP1S(0, 0, 0, 14)

	case AFRINTIS:
		return FPOP1S(0, 0, 0, 15)

	case AFMOVD:
		return FPOP1S(0, 0, 1, 0)

	case AFABSD:
		return FPOP1S(0, 0, 1, 1)

	case AFNEGD:
		return FPOP1S(0, 0, 1, 2)

	case AFSQRTD:
		return FPOP1S(0, 0, 1, 3)

	case AFCVTDS:
		return FPOP1S(0, 0, 1, 4)

	case AFCVTDH:
		return FPOP1S(0, 0, 1, 7)

	case AFRINTND:
		return FPOP1S(0, 0, 1, 8)

	case AFRINTPD:
		return FPOP1S(0, 0, 1, 9)

	case AFRINTMD:
		return FPOP1S(0, 0, 1, 10)

	case AFRINTZD:
		return FPOP1S(0, 0, 1, 11)

	case AFRINTAD:
		return FPOP1S(0, 0, 1, 12)

	case AFRINTXD:
		return FPOP1S(0, 0, 1, 14)

	case AFRINTID:
		return FPOP1S(0, 0, 1, 15)

	case AFCVTHS:
		return FPOP1S(0, 0, 3, 4)

	case AFCVTHD:
		return FPOP1S(0, 0, 3, 5)

	case AVADD:
		return 7<<25 | 1<<21 | 1<<15 | 1<<10

	case AVADDP:
		return 7<<25 | 1<<21 | 1<<15 | 15<<10

	case AVAND:
		return 7<<25 | 1<<21 | 7<<10

	case AVCMEQ:
		return 1<<29 | 0x71<<21 | 0x23<<10

	case AVCNT:
		return 0xE<<24 | 0x10<<17 | 5<<12 | 2<<10

	case AVZIP1:
		return 0xE<<24 | 3<<12 | 2<<10

	case AVZIP2:
		return 0xE<<24 | 1<<14 | 3<<12 | 2<<10

	case AVEOR:
		return 1<<29 | 0x71<<21 | 7<<10

	case AVORR:
		return 7<<25 | 5<<21 | 7<<10

	case AVREV32:
		return 11<<26 | 2<<24 | 1<<21 | 1<<11

	case AVREV64:
		return 3<<26 | 2<<24 | 1<<21 | 1<<11

	case AVMOV:
		return 7<<25 | 5<<21 | 7<<10

	case AVADDV:
		return 7<<25 | 3<<20 | 3<<15 | 7<<11

	case AVUADDLV:
		return 1<<29 | 7<<25 | 3<<20 | 7<<11

	case AVFMLA:
		return 7<<25 | 0<<23 | 1<<21 | 3<<14 | 3<<10

	case AVFMLS:
		return 7<<25 | 1<<23 | 1<<21 | 3<<14 | 3<<10

	case AVPMULL, AVPMULL2:
		return 0xE<<24 | 1<<21 | 0x38<<10

	case AVRBIT:
		return 0x2E<<24 | 1<<22 | 0x10<<17 | 5<<12 | 2<<10
	}

	c.ctxt.Diag("%v: bad rrr %d %v", p, a, a)
	return 0
}

/*
 * imm -> Rd
 * imm op Rn -> Rd
 */
func (c *ctxt7) opirr(p *obj.Prog, a obj.As) uint32 {
	switch a {

	case AMOVD, AADD:
		return S64 | 0<<30 | 0<<29 | 0x11<<24

	case ACMN, AADDS:
		return S64 | 0<<30 | 1<<29 | 0x11<<24

	case AMOVW, AADDW:
		return S32 | 0<<30 | 0<<29 | 0x11<<24

	case ACMNW, AADDSW:
		return S32 | 0<<30 | 1<<29 | 0x11<<24

	case ASUB:
		return S64 | 1<<30 | 0<<29 | 0x11<<24

	case ACMP, ASUBS:
		return S64 | 1<<30 | 1<<29 | 0x11<<24

	case ASUBW:
		return S32 | 1<<30 | 0<<29 | 0x11<<24

	case ACMPW, ASUBSW:
		return S32 | 1<<30 | 1<<29 | 0x11<<24

	case AADR:
		return 0<<31 | 0x10<<24

	case AADRP:
		return 1<<31 | 0x10<<24

	case AAND, ABIC:
		return S64 | 0<<29 | 0x24<<23

	case AANDW, ABICW:
		return S32 | 0<<29 | 0x24<<23 | 0<<22

	case AORR, AORN:
		return S64 | 1<<29 | 0x24<<23

	case AORRW, AORNW:
		return S32 | 1<<29 | 0x24<<23 | 0<<22

	case AEOR, AEON:
		return S64 | 2<<29 | 0x24<<23

	case AEORW, AEONW:
		return S32 | 2<<29 | 0x24<<23 | 0<<22

	case AANDS, ABICS, ATST:
		return S64 | 3<<29 | 0x24<<23

	case AANDSW, ABICSW, ATSTW:
		return S32 | 3<<29 | 0x24<<23 | 0<<22

	case AASR:
		return S64 | 0<<29 | 0x26<<23

	case AASRW:
		return S32 | 0<<29 | 0x26<<23 | 0<<22

	case ABFI:
		return S64 | 2<<29 | 0x26<<23 | 1<<22

	case ABFIW:
		return S32 | 2<<29 | 0x26<<23 | 0<<22

	case ABFM:
		return S64 | 1<<29 | 0x26<<23 | 1<<22

	case ABFMW:
		return S32 | 1<<29 | 0x26<<23 | 0<<22

	case ASBFM:
		return S64 | 0<<29 | 0x26<<23 | 1<<22

	case ASBFMW:
		return S32 | 0<<29 | 0x26<<23 | 0<<22

	case AUBFM:
		return S64 | 2<<29 | 0x26<<23 | 1<<22

	case AUBFMW:
		return S32 | 2<<29 | 0x26<<23 | 0<<22

	case ABFXIL:
		return S64 | 1<<29 | 0x26<<23 | 1<<22

	case ABFXILW:
		return S32 | 1<<29 | 0x26<<23 | 0<<22

	case AEXTR:
		return S64 | 0<<29 | 0x27<<23 | 1<<22 | 0<<21

	case AEXTRW:
		return S32 | 0<<29 | 0x27<<23 | 0<<22 | 0<<21

	case ACBNZ:
		return S64 | 0x1A<<25 | 1<<24

	case ACBNZW:
		return S32 | 0x1A<<25 | 1<<24

	case ACBZ:
		return S64 | 0x1A<<25 | 0<<24

	case ACBZW:
		return S32 | 0x1A<<25 | 0<<24

	case ACCMN:
		return S64 | 0<<30 | 1<<29 | 0xD2<<21 | 1<<11 | 0<<10 | 0<<4

	case ACCMNW:
		return S32 | 0<<30 | 1<<29 | 0xD2<<21 | 1<<11 | 0<<10 | 0<<4

	case ACCMP:
		return S64 | 1<<30 | 1<<29 | 0xD2<<21 | 1<<11 | 0<<10 | 0<<4

	case ACCMPW:
		return S32 | 1<<30 | 1<<29 | 0xD2<<21 | 1<<11 | 0<<10 | 0<<4

	case AMOVK:
		return S64 | 3<<29 | 0x25<<23

	case AMOVKW:
		return S32 | 3<<29 | 0x25<<23

	case AMOVN:
		return S64 | 0<<29 | 0x25<<23

	case AMOVNW:
		return S32 | 0<<29 | 0x25<<23

	case AMOVZ:
		return S64 | 2<<29 | 0x25<<23

	case AMOVZW:
		return S32 | 2<<29 | 0x25<<23

	case AMSR:
		return SYSOP(0, 0, 0, 4, 0, 0, 0x1F)

	case AAT,
		ADC,
		AIC,
		ATLBI,
		ASYS:
		return SYSOP(0, 1, 0, 0, 0, 0, 0)

	case ASYSL:
		return SYSOP(1, 1, 0, 0, 0, 0, 0)

	case ATBZ:
		return 0x36 << 24

	case ATBNZ:
		return 0x37 << 24

	case ADSB:
		return SYSOP(0, 0, 3, 3, 0, 4, 0x1F)

	case ADMB:
		return SYSOP(0, 0, 3, 3, 0, 5, 0x1F)

	case AISB:
		return SYSOP(0, 0, 3, 3, 0, 6, 0x1F)

	case AHINT:
		return SYSOP(0, 0, 3, 2, 0, 0, 0x1F)

	case AVEXT:
		return 0x2E<<24 | 0<<23 | 0<<21 | 0<<15

	case AVUSHR:
		return 0x5E<<23 | 1<<10

	case AVSHL:
		return 0x1E<<23 | 21<<10

	case AVSRI:
		return 0x5E<<23 | 17<<10
	}

	c.ctxt.Diag("%v: bad irr %v", p, a)
	return 0
}

func (c *ctxt7) opbit(p *obj.Prog, a obj.As) uint32 {
	switch a {
	case ACLS:
		return S64 | OPBIT(5)

	case ACLSW:
		return S32 | OPBIT(5)

	case ACLZ:
		return S64 | OPBIT(4)

	case ACLZW:
		return S32 | OPBIT(4)

	case ARBIT:
		return S64 | OPBIT(0)

	case ARBITW:
		return S32 | OPBIT(0)

	case AREV:
		return S64 | OPBIT(3)

	case AREVW:
		return S32 | OPBIT(2)

	case AREV16:
		return S64 | OPBIT(1)

	case AREV16W:
		return S32 | OPBIT(1)

	case AREV32:
		return S64 | OPBIT(2)

	default:
		c.ctxt.Diag("bad bit op\n%v", p)
		return 0
	}
}

/*
 * add/subtract sign or zero-extended register
 */
func (c *ctxt7) opxrrr(p *obj.Prog, a obj.As, extend bool) uint32 {
	extension := uint32(0)
	if !extend {
		switch a {
		case AADD, ACMN, AADDS, ASUB, ACMP, ASUBS:
			extension = LSL0_64

		case AADDW, ACMNW, AADDSW, ASUBW, ACMPW, ASUBSW:
			extension = LSL0_32
		}
	}

	switch a {
	case AADD:
		return S64 | 0<<30 | 0<<29 | 0x0b<<24 | 0<<22 | 1<<21 | extension

	case AADDW:
		return S32 | 0<<30 | 0<<29 | 0x0b<<24 | 0<<22 | 1<<21 | extension

	case ACMN, AADDS:
		return S64 | 0<<30 | 1<<29 | 0x0b<<24 | 0<<22 | 1<<21 | extension

	case ACMNW, AADDSW:
		return S32 | 0<<30 | 1<<29 | 0x0b<<24 | 0<<22 | 1<<21 | extension

	case ASUB:
		return S64 | 1<<30 | 0<<29 | 0x0b<<24 | 0<<22 | 1<<21 | extension

	case ASUBW:
		return S32 | 1<<30 | 0<<29 | 0x0b<<24 | 0<<22 | 1<<21 | extension

	case ACMP, ASUBS:
		return S64 | 1<<30 | 1<<29 | 0x0b<<24 | 0<<22 | 1<<21 | extension

	case ACMPW, ASUBSW:
		return S32 | 1<<30 | 1<<29 | 0x0b<<24 | 0<<22 | 1<<21 | extension
	}

	c.ctxt.Diag("bad opxrrr %v\n%v", a, p)
	return 0
}

func (c *ctxt7) opimm(p *obj.Prog, a obj.As) uint32 {
	switch a {
	case ASVC:
		return 0xD4<<24 | 0<<21 | 1

	case AHVC:
		return 0xD4<<24 | 0<<21 | 2

	case ASMC:
		return 0xD4<<24 | 0<<21 | 3

	case ABRK:
		return 0xD4<<24 | 1<<21 | 0

	case AHLT:
		return 0xD4<<24 | 2<<21 | 0

	case ADCPS1:
		return 0xD4<<24 | 5<<21 | 1

	case ADCPS2:
		return 0xD4<<24 | 5<<21 | 2

	case ADCPS3:
		return 0xD4<<24 | 5<<21 | 3

	case ACLREX:
		return SYSOP(0, 0, 3, 3, 0, 2, 0x1F)
	}

	c.ctxt.Diag("%v: bad imm %v", p, a)
	return 0
}

func (c *ctxt7) brdist(p *obj.Prog, preshift int, flen int, shift int) int64 {
	v := int64(0)
	t := int64(0)
	if p.Pcond != nil {
		v = (p.Pcond.Pc >> uint(preshift)) - (c.pc >> uint(preshift))
		if (v & ((1 << uint(shift)) - 1)) != 0 {
			c.ctxt.Diag("misaligned label\n%v", p)
		}
		v >>= uint(shift)
		t = int64(1) << uint(flen-1)
		if v < -t || v >= t {
			c.ctxt.Diag("branch too far %#x vs %#x [%p]\n%v\n%v", v, t, c.blitrl, p, p.Pcond)
			panic("branch too far")
		}
	}

	return v & ((t << 1) - 1)
}

/*
 * pc-relative branches
 */
func (c *ctxt7) opbra(p *obj.Prog, a obj.As) uint32 {
	switch a {
	case ABEQ:
		return OPBcc(0x0)

	case ABNE:
		return OPBcc(0x1)

	case ABCS:
		return OPBcc(0x2)

	case ABHS:
		return OPBcc(0x2)

	case ABCC:
		return OPBcc(0x3)

	case ABLO:
		return OPBcc(0x3)

	case ABMI:
		return OPBcc(0x4)

	case ABPL:
		return OPBcc(0x5)

	case ABVS:
		return OPBcc(0x6)

	case ABVC:
		return OPBcc(0x7)

	case ABHI:
		return OPBcc(0x8)

	case ABLS:
		return OPBcc(0x9)

	case ABGE:
		return OPBcc(0xa)

	case ABLT:
		return OPBcc(0xb)

	case ABGT:
		return OPBcc(0xc)

	case ABLE:
		return OPBcc(0xd)

	case AB:
		return 0<<31 | 5<<26

	case obj.ADUFFZERO, obj.ADUFFCOPY, ABL:
		return 1<<31 | 5<<26
	}

	c.ctxt.Diag("%v: bad bra %v", p, a)
	return 0
}

func (c *ctxt7) opbrr(p *obj.Prog, a obj.As) uint32 {
	switch a {
	case ABL:
		return OPBLR(1)

	case AB:
		return OPBLR(0)

	case obj.ARET:
		return OPBLR(2)
	}

	c.ctxt.Diag("%v: bad brr %v", p, a)
	return 0
}

func (c *ctxt7) op0(p *obj.Prog, a obj.As) uint32 {
	switch a {
	case ADRPS:
		return 0x6B<<25 | 5<<21 | 0x1F<<16 | 0x1F<<5

	case AERET:
		return 0x6B<<25 | 4<<21 | 0x1F<<16 | 0<<10 | 0x1F<<5

	case AYIELD:
		return SYSHINT(1)

	case AWFE:
		return SYSHINT(2)

	case AWFI:
		return SYSHINT(3)

	case ASEV:
		return SYSHINT(4)

	case ASEVL:
		return SYSHINT(5)
	}

	c.ctxt.Diag("%v: bad op0 %v", p, a)
	return 0
}

/*
 * register offset
 */
func (c *ctxt7) opload(p *obj.Prog, a obj.As) uint32 {
	switch a {
	case ALDAR:
		return LDSTX(3, 1, 1, 0, 1) | 0x1F<<10

	case ALDARW:
		return LDSTX(2, 1, 1, 0, 1) | 0x1F<<10

	case ALDARB:
		return LDSTX(0, 1, 1, 0, 1) | 0x1F<<10

	case ALDARH:
		return LDSTX(1, 1, 1, 0, 1) | 0x1F<<10

	case ALDAXP:
		return LDSTX(3, 0, 1, 1, 1)

	case ALDAXPW:
		return LDSTX(2, 0, 1, 1, 1)

	case ALDAXR:
		return LDSTX(3, 0, 1, 0, 1) | 0x1F<<10

	case ALDAXRW:
		return LDSTX(2, 0, 1, 0, 1) | 0x1F<<10

	case ALDAXRB:
		return LDSTX(0, 0, 1, 0, 1) | 0x1F<<10

	case ALDAXRH:
		return LDSTX(1, 0, 1, 0, 1) | 0x1F<<10

	case ALDXR:
		return LDSTX(3, 0, 1, 0, 0) | 0x1F<<10

	case ALDXRB:
		return LDSTX(0, 0, 1, 0, 0) | 0x1F<<10

	case ALDXRH:
		return LDSTX(1, 0, 1, 0, 0) | 0x1F<<10

	case ALDXRW:
		return LDSTX(2, 0, 1, 0, 0) | 0x1F<<10

	case ALDXP:
		return LDSTX(3, 0, 1, 1, 0)

	case ALDXPW:
		return LDSTX(2, 0, 1, 1, 0)

	case AMOVNP:
		return S64 | 0<<30 | 5<<27 | 0<<26 | 0<<23 | 1<<22

	case AMOVNPW:
		return S32 | 0<<30 | 5<<27 | 0<<26 | 0<<23 | 1<<22
	}

	c.ctxt.Diag("bad opload %v\n%v", a, p)
	return 0
}

func (c *ctxt7) opstore(p *obj.Prog, a obj.As) uint32 {
	switch a {
	case ASTLR:
		return LDSTX(3, 1, 0, 0, 1) | 0x1F<<10

	case ASTLRB:
		return LDSTX(0, 1, 0, 0, 1) | 0x1F<<10

	case ASTLRH:
		return LDSTX(1, 1, 0, 0, 1) | 0x1F<<10

	case ASTLP:
		return LDSTX(3, 0, 0, 1, 1)

	case ASTLPW:
		return LDSTX(2, 0, 0, 1, 1)

	case ASTLRW:
		return LDSTX(2, 1, 0, 0, 1) | 0x1F<<10

	case ASTLXP:
		return LDSTX(3, 0, 0, 1, 1)

	case ASTLXPW:
		return LDSTX(2, 0, 0, 1, 1)

	case ASTLXR:
		return LDSTX(3, 0, 0, 0, 1) | 0x1F<<10

	case ASTLXRB:
		return LDSTX(0, 0, 0, 0, 1) | 0x1F<<10

	case ASTLXRH:
		return LDSTX(1, 0, 0, 0, 1) | 0x1F<<10

	case ASTLXRW:
		return LDSTX(2, 0, 0, 0, 1) | 0x1F<<10

	case ASTXR:
		return LDSTX(3, 0, 0, 0, 0) | 0x1F<<10

	case ASTXRB:
		return LDSTX(0, 0, 0, 0, 0) | 0x1F<<10

	case ASTXRH:
		return LDSTX(1, 0, 0, 0, 0) | 0x1F<<10

	case ASTXP:
		return LDSTX(3, 0, 0, 1, 0)

	case ASTXPW:
		return LDSTX(2, 0, 0, 1, 0)

	case ASTXRW:
		return LDSTX(2, 0, 0, 0, 0) | 0x1F<<10

	case AMOVNP:
		return S64 | 0<<30 | 5<<27 | 0<<26 | 0<<23 | 1<<22

	case AMOVNPW:
		return S32 | 0<<30 | 5<<27 | 0<<26 | 0<<23 | 1<<22
	}

	c.ctxt.Diag("bad opstore %v\n%v", a, p)
	return 0
}

/*
 * load/store register (unsigned immediate) C3.3.13
 *	these produce 64-bit values (when there's an option)
 */
func (c *ctxt7) olsr12u(p *obj.Prog, o int32, v int32, b int, r int) uint32 {
	if v < 0 || v >= (1<<12) {
		c.ctxt.Diag("offset out of range: %d\n%v", v, p)
	}
	o |= (v & 0xFFF) << 10
	o |= int32(b&31) << 5
	o |= int32(r & 31)
	return uint32(o)
}

func (c *ctxt7) opldr12(p *obj.Prog, a obj.As) uint32 {
	switch a {
	case AMOVD:
		return LDSTR12U(3, 0, 1)

	case AMOVW:
		return LDSTR12U(2, 0, 2)

	case AMOVWU:
		return LDSTR12U(2, 0, 1)

	case AMOVH:
		return LDSTR12U(1, 0, 2)

	case AMOVHU:
		return LDSTR12U(1, 0, 1)

	case AMOVB:
		return LDSTR12U(0, 0, 2)

	case AMOVBU:
		return LDSTR12U(0, 0, 1)

	case AFMOVS:
		return LDSTR12U(2, 1, 1)

	case AFMOVD:
		return LDSTR12U(3, 1, 1)
	}

	c.ctxt.Diag("bad opldr12 %v\n%v", a, p)
	return 0
}

func (c *ctxt7) opstr12(p *obj.Prog, a obj.As) uint32 {
	return LD2STR(c.opldr12(p, a))
}

/*
 * load/store register (unscaled immediate) C3.3.12
 */
func (c *ctxt7) olsr9s(p *obj.Prog, o int32, v int32, b int, r int) uint32 {
	if v < -256 || v > 255 {
		c.ctxt.Diag("offset out of range: %d\n%v", v, p)
	}
	o |= (v & 0x1FF) << 12
	o |= int32(b&31) << 5
	o |= int32(r & 31)
	return uint32(o)
}

func (c *ctxt7) opldr9(p *obj.Prog, a obj.As) uint32 {
	switch a {
	case AMOVD:
		return LDSTR9S(3, 0, 1)

	case AMOVW:
		return LDSTR9S(2, 0, 2)

	case AMOVWU:
		return LDSTR9S(2, 0, 1)

	case AMOVH:
		return LDSTR9S(1, 0, 2)

	case AMOVHU:
		return LDSTR9S(1, 0, 1)

	case AMOVB:
		return LDSTR9S(0, 0, 2)

	case AMOVBU:
		return LDSTR9S(0, 0, 1)

	case AFMOVS:
		return LDSTR9S(2, 1, 1)

	case AFMOVD:
		return LDSTR9S(3, 1, 1)
	}

	c.ctxt.Diag("bad opldr9 %v\n%v", a, p)
	return 0
}

func (c *ctxt7) opstr9(p *obj.Prog, a obj.As) uint32 {
	return LD2STR(c.opldr9(p, a))
}

func (c *ctxt7) opldrpp(p *obj.Prog, a obj.As) uint32 {
	switch a {
	case AMOVD:
		return 3<<30 | 7<<27 | 0<<26 | 0<<24 | 1<<22

	case AMOVW:
		return 2<<30 | 7<<27 | 0<<26 | 0<<24 | 2<<22

	case AMOVWU:
		return 2<<30 | 7<<27 | 0<<26 | 0<<24 | 1<<22

	case AMOVH:
		return 1<<30 | 7<<27 | 0<<26 | 0<<24 | 2<<22

	case AMOVHU:
		return 1<<30 | 7<<27 | 0<<26 | 0<<24 | 1<<22

	case AMOVB:
		return 0<<30 | 7<<27 | 0<<26 | 0<<24 | 2<<22

	case AMOVBU:
		return 0<<30 | 7<<27 | 0<<26 | 0<<24 | 1<<22

	case AFMOVS:
		return 2<<30 | 7<<27 | 1<<26 | 0<<24 | 1<<22

	case AFMOVD:
		return 3<<30 | 7<<27 | 1<<26 | 0<<24 | 1<<22

	case APRFM:
		return 0xf9<<24 | 2<<22

	}

	c.ctxt.Diag("bad opldr %v\n%v", a, p)
	return 0
}

// olsxrr attaches register operands to a load/store opcode supplied in o.
// The result either encodes a load of r from (r1+r2) or a store of r to (r1+r2).
func (c *ctxt7) olsxrr(p *obj.Prog, o int32, r int, r1 int, r2 int) uint32 {
	o |= int32(r1&31) << 5
	o |= int32(r2&31) << 16
	o |= int32(r & 31)
	return uint32(o)
}

// opldrr returns the ARM64 opcode encoding corresponding to the obj.As opcode
// for load instruction with register offset.
// The offset register can be (Rn)(Rm.UXTW<<2) or (Rn)(Rm<<2) or (Rn)(Rm).
func (c *ctxt7) opldrr(p *obj.Prog, a obj.As, extension bool) uint32 {
	OptionS := uint32(0x1a)
	if extension {
		OptionS = uint32(0)
	}
	switch a {
	case AMOVD:
		return OptionS<<10 | 0x3<<21 | 0x1f<<27
	case AMOVW:
		return OptionS<<10 | 0x5<<21 | 0x17<<27
	case AMOVWU:
		return OptionS<<10 | 0x3<<21 | 0x17<<27
	case AMOVH:
		return OptionS<<10 | 0x5<<21 | 0x0f<<27
	case AMOVHU:
		return OptionS<<10 | 0x3<<21 | 0x0f<<27
	case AMOVB:
		return OptionS<<10 | 0x5<<21 | 0x07<<27
	case AMOVBU:
		return OptionS<<10 | 0x3<<21 | 0x07<<27
	case AFMOVS:
		return OptionS<<10 | 0x3<<21 | 0x17<<27 | 1<<26
	case AFMOVD:
		return OptionS<<10 | 0x3<<21 | 0x1f<<27 | 1<<26
	}
	c.ctxt.Diag("bad opldrr %v\n%v", a, p)
	return 0
}

// opstrr returns the ARM64 opcode encoding corresponding to the obj.As opcode
// for store instruction with register offset.
// The offset register can be (Rn)(Rm.UXTW<<2) or (Rn)(Rm<<2) or (Rn)(Rm).
func (c *ctxt7) opstrr(p *obj.Prog, a obj.As, extension bool) uint32 {
	OptionS := uint32(0x1a)
	if extension {
		OptionS = uint32(0)
	}
	switch a {
	case AMOVD:
		return OptionS<<10 | 0x1<<21 | 0x1f<<27
	case AMOVW, AMOVWU:
		return OptionS<<10 | 0x1<<21 | 0x17<<27
	case AMOVH, AMOVHU:
		return OptionS<<10 | 0x1<<21 | 0x0f<<27
	case AMOVB, AMOVBU:
		return OptionS<<10 | 0x1<<21 | 0x07<<27
	case AFMOVS:
		return OptionS<<10 | 0x1<<21 | 0x17<<27 | 1<<26
	case AFMOVD:
		return OptionS<<10 | 0x1<<21 | 0x1f<<27 | 1<<26
	}
	c.ctxt.Diag("bad opstrr %v\n%v", a, p)
	return 0
}

func (c *ctxt7) oaddi(p *obj.Prog, o1 int32, v int32, r int, rt int) uint32 {
	if (v & 0xFFF000) != 0 {
		if v&0xFFF != 0 {
			c.ctxt.Diag("%v misuses oaddi", p)
		}
		v >>= 12
		o1 |= 1 << 22
	}

	o1 |= ((v & 0xFFF) << 10) | (int32(r&31) << 5) | int32(rt&31)
	return uint32(o1)
}

/*
 * load a a literal value into dr
 */
func (c *ctxt7) omovlit(as obj.As, p *obj.Prog, a *obj.Addr, dr int) uint32 {
	var o1 int32
	if p.Pcond == nil {
		c.aclass(a)
		c.ctxt.Logf("omovlit add %d (%#x)\n", c.instoffset, uint64(c.instoffset))

		o1 = int32(c.opirr(p, AADD))

		v := int32(c.instoffset)
		if v != 0 && (v&0xFFF) == 0 {
			v >>= 12
			o1 |= 1 << 22
		}

		o1 |= ((v & 0xFFF) << 10) | (REGZERO & 31 << 5) | int32(dr&31)
	} else {
		fp, w := 0, 0
		switch as {
		case AFMOVS:
			fp = 1
			w = 0

		case AFMOVD:
			fp = 1
			w = 1

		case AMOVD:
			if p.Pcond.As == ADWORD {
				w = 1
			} else if p.Pcond.To.Offset < 0 {
				w = 2
			} else if p.Pcond.To.Offset >= 0 {
				w = 0
			} else {
				c.ctxt.Diag("invalid operand %v in %v", a, p)
			}

		case AMOVBU, AMOVHU, AMOVWU:
			w = 0

		case AMOVB, AMOVH, AMOVW:
			w = 2

		default:
			c.ctxt.Diag("invalid operation %v in %v", as, p)
		}

		v := int32(c.brdist(p, 0, 19, 2))
		o1 = (int32(w) << 30) | (int32(fp) << 26) | (3 << 27)
		o1 |= (v & 0x7FFFF) << 5
		o1 |= int32(dr & 31)
	}

	return uint32(o1)
}

// load a constant (MOVCON or BITCON) in a into rt
func (c *ctxt7) omovconst(as obj.As, p *obj.Prog, a *obj.Addr, rt int) (o1 uint32) {
	if cls := oclass(a); cls == C_BITCON || cls == C_ABCON || cls == C_ABCON0 {

		mode := 64
		var as1 obj.As
		switch as {
		case AMOVW:
			as1 = AORRW
			mode = 32
		case AMOVD:
			as1 = AORR
		}
		o1 = c.opirr(p, as1)
		o1 |= bitconEncode(uint64(a.Offset), mode) | uint32(REGZERO&31)<<5 | uint32(rt&31)
		return o1
	}

	r := 32
	if as == AMOVD {
		r = 64
	}
	d := a.Offset
	s := movcon(d)
	if s < 0 || s >= r {
		d = ^d
		s = movcon(d)
		if s < 0 || s >= r {
			c.ctxt.Diag("impossible move wide: %#x\n%v", uint64(a.Offset), p)
		}
		if as == AMOVD {
			o1 = c.opirr(p, AMOVN)
		} else {
			o1 = c.opirr(p, AMOVNW)
		}
	} else {
		if as == AMOVD {
			o1 = c.opirr(p, AMOVZ)
		} else {
			o1 = c.opirr(p, AMOVZW)
		}
	}
	o1 |= uint32((((d >> uint(s*16)) & 0xFFFF) << 5) | int64((uint32(s)&3)<<21) | int64(rt&31))
	return o1
}

func (c *ctxt7) opbfm(p *obj.Prog, a obj.As, r int, s int, rf int, rt int) uint32 {
	var b uint32
	o := c.opirr(p, a)
	if (o & (1 << 31)) == 0 {
		b = 32
	} else {
		b = 64
	}
	if r < 0 || uint32(r) >= b {
		c.ctxt.Diag("illegal bit number\n%v", p)
	}
	o |= (uint32(r) & 0x3F) << 16
	if s < 0 || uint32(s) >= b {
		c.ctxt.Diag("illegal bit number\n%v", p)
	}
	o |= (uint32(s) & 0x3F) << 10
	o |= (uint32(rf&31) << 5) | uint32(rt&31)
	return o
}

func (c *ctxt7) opextr(p *obj.Prog, a obj.As, v int32, rn int, rm int, rt int) uint32 {
	var b uint32
	o := c.opirr(p, a)
	if (o & (1 << 31)) != 0 {
		b = 63
	} else {
		b = 31
	}
	if v < 0 || uint32(v) > b {
		c.ctxt.Diag("illegal bit number\n%v", p)
	}
	o |= uint32(v) << 10
	o |= uint32(rn&31) << 5
	o |= uint32(rm&31) << 16
	o |= uint32(rt & 31)
	return o
}

/* genrate instruction encoding for LDP/LDPW/LDPSW/STP/STPW */
func (c *ctxt7) opldpstp(p *obj.Prog, o *Optab, vo int32, rbase, rl, rh, ldp uint32) uint32 {
	var ret uint32
	switch p.As {
	case ALDP, ASTP:
		if vo < -512 || vo > 504 || vo%8 != 0 {
			c.ctxt.Diag("invalid offset %v\n", p)
		}
		vo /= 8
		ret = 2 << 30
	case ALDPW, ASTPW:
		if vo < -256 || vo > 252 || vo%4 != 0 {
			c.ctxt.Diag("invalid offset %v\n", p)
		}
		vo /= 4
		ret = 0
	case ALDPSW:
		if vo < -256 || vo > 252 || vo%4 != 0 {
			c.ctxt.Diag("invalid offset %v\n", p)
		}
		vo /= 4
		ret = 1 << 30
	default:
		c.ctxt.Diag("invalid instruction %v\n", p)
	}
	switch o.scond {
	case C_XPOST:
		ret |= 1 << 23
	case C_XPRE:
		ret |= 3 << 23
	default:
		ret |= 2 << 23
	}
	ret |= 5<<27 | (ldp&1)<<22 | uint32(vo&0x7f)<<15 | (rh&31)<<10 | (rbase&31)<<5 | (rl & 31)
	return ret
}

/*
 * size in log2(bytes)
 */
func movesize(a obj.As) int {
	switch a {
	case AMOVD:
		return 3

	case AMOVW, AMOVWU:
		return 2

	case AMOVH, AMOVHU:
		return 1

	case AMOVB, AMOVBU:
		return 0

	case AFMOVS:
		return 2

	case AFMOVD:
		return 3

	default:
		return -1
	}
}

// rm is the Rm register value, o is the extension, amount is the left shift value.
func roff(rm int16, o uint32, amount int16) uint32 {
	return uint32(rm&31)<<16 | o<<13 | uint32(amount)<<10
}

// encRegShiftOrExt returns the encoding of shifted/extended register, Rx<<n and Rx.UXTW<<n, etc.
func (c *ctxt7) encRegShiftOrExt(a *obj.Addr, r int16) uint32 {
	var num, rm int16
	num = (r >> 5) & 7
	rm = r & 31
	switch {
	case REG_UXTB <= r && r < REG_UXTH:
		return roff(rm, 0, num)
	case REG_UXTH <= r && r < REG_UXTW:
		return roff(rm, 1, num)
	case REG_UXTW <= r && r < REG_UXTX:
		if a.Type == obj.TYPE_MEM {
			if num == 0 {
				return roff(rm, 2, 2)
			} else {
				return roff(rm, 2, 6)
			}
		} else {
			return roff(rm, 2, num)
		}
	case REG_UXTX <= r && r < REG_SXTB:
		return roff(rm, 3, num)
	case REG_SXTB <= r && r < REG_SXTH:
		return roff(rm, 4, num)
	case REG_SXTH <= r && r < REG_SXTW:
		return roff(rm, 5, num)
	case REG_SXTW <= r && r < REG_SXTX:
		if a.Type == obj.TYPE_MEM {
			if num == 0 {
				return roff(rm, 6, 2)
			} else {
				return roff(rm, 6, 6)
			}
		} else {
			return roff(rm, 6, num)
		}
	case REG_SXTX <= r && r < REG_SPECIAL:
		if a.Type == obj.TYPE_MEM {
			if num == 0 {
				return roff(rm, 7, 2)
			} else {
				return roff(rm, 7, 6)
			}
		} else {
			return roff(rm, 7, num)
		}
	case REG_LSL <= r && r < (REG_LSL+1<<8):
		return roff(rm, 3, 6)
	default:
		c.ctxt.Diag("unsupported register extension type.")
	}

	return 0
}

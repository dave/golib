// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s390x

import (
	"math"

	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/s390x"
)

// markMoves marks any MOVXconst ops that need to avoid clobbering flags.
func (pstate *PackageState) ssaMarkMoves(s *gc.SSAGenState, b *ssa.Block) {
	flive := b.FlagsLiveAtEnd
	if b.Control != nil && b.Control.Type.IsFlags(pstate.types) {
		flive = true
	}
	for i := len(b.Values) - 1; i >= 0; i-- {
		v := b.Values[i]
		if flive && v.Op == ssa.OpS390XMOVDconst {
			// The "mark" is any non-nil Aux value.
			v.Aux = v
		}
		if v.Type.IsFlags(pstate.types) {
			flive = false
		}
		for _, a := range v.Args {
			if a.Type.IsFlags(pstate.types) {
				flive = true
			}
		}
	}
}

// loadByType returns the load instruction of the given type.
func (pstate *PackageState) loadByType(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(pstate.types) {
		case 4:
			return s390x.AFMOVS
		case 8:
			return s390x.AFMOVD
		}
	} else {
		switch t.Size(pstate.types) {
		case 1:
			if t.IsSigned() {
				return s390x.AMOVB
			} else {
				return s390x.AMOVBZ
			}
		case 2:
			if t.IsSigned() {
				return s390x.AMOVH
			} else {
				return s390x.AMOVHZ
			}
		case 4:
			if t.IsSigned() {
				return s390x.AMOVW
			} else {
				return s390x.AMOVWZ
			}
		case 8:
			return s390x.AMOVD
		}
	}
	panic("bad load type")
}

// storeByType returns the store instruction of the given type.
func (pstate *PackageState) storeByType(t *types.Type) obj.As {
	width := t.Size(pstate.types)
	if t.IsFloat() {
		switch width {
		case 4:
			return s390x.AFMOVS
		case 8:
			return s390x.AFMOVD
		}
	} else {
		switch width {
		case 1:
			return s390x.AMOVB
		case 2:
			return s390x.AMOVH
		case 4:
			return s390x.AMOVW
		case 8:
			return s390x.AMOVD
		}
	}
	panic("bad store type")
}

// moveByType returns the reg->reg move instruction of the given type.
func (pstate *PackageState) moveByType(t *types.Type) obj.As {
	if t.IsFloat() {
		return s390x.AFMOVD
	} else {
		switch t.Size(pstate.types) {
		case 1:
			if t.IsSigned() {
				return s390x.AMOVB
			} else {
				return s390x.AMOVBZ
			}
		case 2:
			if t.IsSigned() {
				return s390x.AMOVH
			} else {
				return s390x.AMOVHZ
			}
		case 4:
			if t.IsSigned() {
				return s390x.AMOVW
			} else {
				return s390x.AMOVWZ
			}
		case 8:
			return s390x.AMOVD
		}
	}
	panic("bad load type")
}

// opregreg emits instructions for
//     dest := dest(To) op src(From)
// and also returns the created obj.Prog so it
// may be further adjusted (offset, scale, etc).
func (pstate *PackageState) opregreg(s *gc.SSAGenState, op obj.As, dest, src int16) *obj.Prog {
	p := s.Prog(pstate.gc, op)
	p.From.Type = obj.TYPE_REG
	p.To.Type = obj.TYPE_REG
	p.To.Reg = dest
	p.From.Reg = src
	return p
}

// opregregimm emits instructions for
//	dest := src(From) op off
// and also returns the created obj.Prog so it
// may be further adjusted (offset, scale, etc).
func (pstate *PackageState) opregregimm(s *gc.SSAGenState, op obj.As, dest, src int16, off int64) *obj.Prog {
	p := s.Prog(pstate.gc, op)
	p.From.Type = obj.TYPE_CONST
	p.From.Offset = off
	p.Reg = src
	p.To.Reg = dest
	p.To.Type = obj.TYPE_REG
	return p
}

func (pstate *PackageState) ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpS390XSLD, ssa.OpS390XSLW,
		ssa.OpS390XSRD, ssa.OpS390XSRW,
		ssa.OpS390XSRAD, ssa.OpS390XSRAW:
		r := v.Reg(pstate.ssa)
		r1 := v.Args[0].Reg(pstate.ssa)
		r2 := v.Args[1].Reg(pstate.ssa)
		if r2 == s390x.REG_R0 {
			v.Fatalf("cannot use R0 as shift value %s", v.LongString(pstate.ssa))
		}
		p := pstate.opregreg(s, v.Op.Asm(pstate.ssa), r, r2)
		if r != r1 {
			p.Reg = r1
		}
	case ssa.OpS390XADD, ssa.OpS390XADDW,
		ssa.OpS390XSUB, ssa.OpS390XSUBW,
		ssa.OpS390XAND, ssa.OpS390XANDW,
		ssa.OpS390XOR, ssa.OpS390XORW,
		ssa.OpS390XXOR, ssa.OpS390XXORW:
		r := v.Reg(pstate.ssa)
		r1 := v.Args[0].Reg(pstate.ssa)
		r2 := v.Args[1].Reg(pstate.ssa)
		p := pstate.opregreg(s, v.Op.Asm(pstate.ssa), r, r2)
		if r != r1 {
			p.Reg = r1
		}
	// 2-address opcode arithmetic
	case ssa.OpS390XMULLD, ssa.OpS390XMULLW,
		ssa.OpS390XMULHD, ssa.OpS390XMULHDU,
		ssa.OpS390XFADDS, ssa.OpS390XFADD, ssa.OpS390XFSUBS, ssa.OpS390XFSUB,
		ssa.OpS390XFMULS, ssa.OpS390XFMUL, ssa.OpS390XFDIVS, ssa.OpS390XFDIV:
		r := v.Reg(pstate.ssa)
		if r != v.Args[0].Reg(pstate.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(pstate.ssa))
		}
		pstate.opregreg(s, v.Op.Asm(pstate.ssa), r, v.Args[1].Reg(pstate.ssa))
	case ssa.OpS390XFMADD, ssa.OpS390XFMADDS,
		ssa.OpS390XFMSUB, ssa.OpS390XFMSUBS:
		r := v.Reg(pstate.ssa)
		if r != v.Args[0].Reg(pstate.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(pstate.ssa))
		}
		r1 := v.Args[1].Reg(pstate.ssa)
		r2 := v.Args[2].Reg(pstate.ssa)
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.Reg = r2
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XFIDBR:
		switch v.AuxInt {
		case 0, 1, 3, 4, 5, 6, 7:
			pstate.opregregimm(s, v.Op.Asm(pstate.ssa), v.Reg(pstate.ssa), v.Args[0].Reg(pstate.ssa), v.AuxInt)
		default:
			v.Fatalf("invalid FIDBR mask: %v", v.AuxInt)
		}
	case ssa.OpS390XCPSDR:
		p := pstate.opregreg(s, v.Op.Asm(pstate.ssa), v.Reg(pstate.ssa), v.Args[1].Reg(pstate.ssa))
		p.Reg = v.Args[0].Reg(pstate.ssa)
	case ssa.OpS390XDIVD, ssa.OpS390XDIVW,
		ssa.OpS390XDIVDU, ssa.OpS390XDIVWU,
		ssa.OpS390XMODD, ssa.OpS390XMODW,
		ssa.OpS390XMODDU, ssa.OpS390XMODWU:

		// TODO(mundaym): use the temp registers every time like x86 does with AX?
		dividend := v.Args[0].Reg(pstate.ssa)
		divisor := v.Args[1].Reg(pstate.ssa)

		// CPU faults upon signed overflow, which occurs when most
		// negative int is divided by -1.
		var j *obj.Prog
		if v.Op == ssa.OpS390XDIVD || v.Op == ssa.OpS390XDIVW ||
			v.Op == ssa.OpS390XMODD || v.Op == ssa.OpS390XMODW {

			var c *obj.Prog
			c = s.Prog(pstate.gc, s390x.ACMP)
			j = s.Prog(pstate.gc, s390x.ABEQ)

			c.From.Type = obj.TYPE_REG
			c.From.Reg = divisor
			c.To.Type = obj.TYPE_CONST
			c.To.Offset = -1

			j.To.Type = obj.TYPE_BRANCH

		}

		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = divisor
		p.Reg = 0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = dividend

		// signed division, rest of the check for -1 case
		if j != nil {
			j2 := s.Prog(pstate.gc, s390x.ABR)
			j2.To.Type = obj.TYPE_BRANCH

			var n *obj.Prog
			if v.Op == ssa.OpS390XDIVD || v.Op == ssa.OpS390XDIVW {
				// n * -1 = -n
				n = s.Prog(pstate.gc, s390x.ANEG)
				n.To.Type = obj.TYPE_REG
				n.To.Reg = dividend
			} else {
				// n % -1 == 0
				n = s.Prog(pstate.gc, s390x.AXOR)
				n.From.Type = obj.TYPE_REG
				n.From.Reg = dividend
				n.To.Type = obj.TYPE_REG
				n.To.Reg = dividend
			}

			j.To.Val = n
			j2.To.Val = s.Pc()
		}
	case ssa.OpS390XADDconst, ssa.OpS390XADDWconst:
		pstate.opregregimm(s, v.Op.Asm(pstate.ssa), v.Reg(pstate.ssa), v.Args[0].Reg(pstate.ssa), v.AuxInt)
	case ssa.OpS390XMULLDconst, ssa.OpS390XMULLWconst,
		ssa.OpS390XSUBconst, ssa.OpS390XSUBWconst,
		ssa.OpS390XANDconst, ssa.OpS390XANDWconst,
		ssa.OpS390XORconst, ssa.OpS390XORWconst,
		ssa.OpS390XXORconst, ssa.OpS390XXORWconst:
		r := v.Reg(pstate.ssa)
		if r != v.Args[0].Reg(pstate.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(pstate.ssa))
		}
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XSLDconst, ssa.OpS390XSLWconst,
		ssa.OpS390XSRDconst, ssa.OpS390XSRWconst,
		ssa.OpS390XSRADconst, ssa.OpS390XSRAWconst,
		ssa.OpS390XRLLGconst, ssa.OpS390XRLLconst:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		r := v.Reg(pstate.ssa)
		r1 := v.Args[0].Reg(pstate.ssa)
		if r != r1 {
			p.Reg = r1
		}
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XMOVDaddridx:
		r := v.Args[0].Reg(pstate.ssa)
		i := v.Args[1].Reg(pstate.ssa)
		p := s.Prog(pstate.gc, s390x.AMOVD)
		p.From.Scale = 1
		if i == s390x.REGSP {
			r, i = i, r
		}
		p.From.Type = obj.TYPE_ADDR
		p.From.Reg = r
		p.From.Index = i
		pstate.gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)
	case ssa.OpS390XMOVDaddr:
		p := s.Prog(pstate.gc, s390x.AMOVD)
		p.From.Type = obj.TYPE_ADDR
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)
	case ssa.OpS390XCMP, ssa.OpS390XCMPW, ssa.OpS390XCMPU, ssa.OpS390XCMPWU:
		pstate.opregreg(s, v.Op.Asm(pstate.ssa), v.Args[1].Reg(pstate.ssa), v.Args[0].Reg(pstate.ssa))
	case ssa.OpS390XFCMPS, ssa.OpS390XFCMP:
		pstate.opregreg(s, v.Op.Asm(pstate.ssa), v.Args[1].Reg(pstate.ssa), v.Args[0].Reg(pstate.ssa))
	case ssa.OpS390XCMPconst, ssa.OpS390XCMPWconst:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = v.AuxInt
	case ssa.OpS390XCMPUconst, ssa.OpS390XCMPWUconst:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = int64(uint32(v.AuxInt))
	case ssa.OpS390XMOVDconst:
		x := v.Reg(pstate.ssa)
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x
	case ssa.OpS390XFMOVSconst, ssa.OpS390XFMOVDconst:
		x := v.Reg(pstate.ssa)
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_FCONST
		p.From.Val = math.Float64frombits(uint64(v.AuxInt))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x
	case ssa.OpS390XADDWload, ssa.OpS390XADDload,
		ssa.OpS390XMULLWload, ssa.OpS390XMULLDload,
		ssa.OpS390XSUBWload, ssa.OpS390XSUBload,
		ssa.OpS390XANDWload, ssa.OpS390XANDload,
		ssa.OpS390XORWload, ssa.OpS390XORload,
		ssa.OpS390XXORWload, ssa.OpS390XXORload:
		r := v.Reg(pstate.ssa)
		if r != v.Args[0].Reg(pstate.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(pstate.ssa))
		}
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[1].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XMOVDload,
		ssa.OpS390XMOVWZload, ssa.OpS390XMOVHZload, ssa.OpS390XMOVBZload,
		ssa.OpS390XMOVDBRload, ssa.OpS390XMOVWBRload, ssa.OpS390XMOVHBRload,
		ssa.OpS390XMOVBload, ssa.OpS390XMOVHload, ssa.OpS390XMOVWload,
		ssa.OpS390XFMOVSload, ssa.OpS390XFMOVDload:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)
	case ssa.OpS390XMOVBZloadidx, ssa.OpS390XMOVHZloadidx, ssa.OpS390XMOVWZloadidx,
		ssa.OpS390XMOVBloadidx, ssa.OpS390XMOVHloadidx, ssa.OpS390XMOVWloadidx, ssa.OpS390XMOVDloadidx,
		ssa.OpS390XMOVHBRloadidx, ssa.OpS390XMOVWBRloadidx, ssa.OpS390XMOVDBRloadidx,
		ssa.OpS390XFMOVSloadidx, ssa.OpS390XFMOVDloadidx:
		r := v.Args[0].Reg(pstate.ssa)
		i := v.Args[1].Reg(pstate.ssa)
		if i == s390x.REGSP {
			r, i = i, r
		}
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r
		p.From.Scale = 1
		p.From.Index = i
		pstate.gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)
	case ssa.OpS390XMOVBstore, ssa.OpS390XMOVHstore, ssa.OpS390XMOVWstore, ssa.OpS390XMOVDstore,
		ssa.OpS390XMOVHBRstore, ssa.OpS390XMOVWBRstore, ssa.OpS390XMOVDBRstore,
		ssa.OpS390XFMOVSstore, ssa.OpS390XFMOVDstore:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.To, v)
	case ssa.OpS390XMOVBstoreidx, ssa.OpS390XMOVHstoreidx, ssa.OpS390XMOVWstoreidx, ssa.OpS390XMOVDstoreidx,
		ssa.OpS390XMOVHBRstoreidx, ssa.OpS390XMOVWBRstoreidx, ssa.OpS390XMOVDBRstoreidx,
		ssa.OpS390XFMOVSstoreidx, ssa.OpS390XFMOVDstoreidx:
		r := v.Args[0].Reg(pstate.ssa)
		i := v.Args[1].Reg(pstate.ssa)
		if i == s390x.REGSP {
			r, i = i, r
		}
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[2].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = r
		p.To.Scale = 1
		p.To.Index = i
		pstate.gc.AddAux(&p.To, v)
	case ssa.OpS390XMOVDstoreconst, ssa.OpS390XMOVWstoreconst, ssa.OpS390XMOVHstoreconst, ssa.OpS390XMOVBstoreconst:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_CONST
		sc := v.AuxValAndOff(pstate.ssa)
		p.From.Offset = sc.Val()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux2(&p.To, v, sc.Off())
	case ssa.OpS390XMOVBreg, ssa.OpS390XMOVHreg, ssa.OpS390XMOVWreg,
		ssa.OpS390XMOVBZreg, ssa.OpS390XMOVHZreg, ssa.OpS390XMOVWZreg,
		ssa.OpS390XLDGR, ssa.OpS390XLGDR,
		ssa.OpS390XCEFBRA, ssa.OpS390XCDFBRA, ssa.OpS390XCEGBRA, ssa.OpS390XCDGBRA,
		ssa.OpS390XCFEBRA, ssa.OpS390XCFDBRA, ssa.OpS390XCGEBRA, ssa.OpS390XCGDBRA,
		ssa.OpS390XLDEBR, ssa.OpS390XLEDBR,
		ssa.OpS390XFNEG, ssa.OpS390XFNEGS,
		ssa.OpS390XLPDFR, ssa.OpS390XLNDFR:
		pstate.opregreg(s, v.Op.Asm(pstate.ssa), v.Reg(pstate.ssa), v.Args[0].Reg(pstate.ssa))
	case ssa.OpS390XCLEAR:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_CONST
		sc := v.AuxValAndOff(pstate.ssa)
		p.From.Offset = sc.Val()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux2(&p.To, v, sc.Off())
	case ssa.OpCopy, ssa.OpS390XMOVDreg:
		if v.Type.IsMemory(pstate.types) {
			return
		}
		x := v.Args[0].Reg(pstate.ssa)
		y := v.Reg(pstate.ssa)
		if x != y {
			pstate.opregreg(s, pstate.moveByType(v.Type), y, x)
		}
	case ssa.OpS390XMOVDnop:
		if v.Reg(pstate.ssa) != v.Args[0].Reg(pstate.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(pstate.ssa))
		}
	// nothing to do
	case ssa.OpLoadReg:
		if v.Type.IsFlags(pstate.types) {
			v.Fatalf("load flags not implemented: %v", v.LongString(pstate.ssa))
			return
		}
		p := s.Prog(pstate.gc, pstate.loadByType(v.Type))
		pstate.gc.AddrAuto(&p.From, v.Args[0])
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)
	case ssa.OpStoreReg:
		if v.Type.IsFlags(pstate.types) {
			v.Fatalf("store flags not implemented: %v", v.LongString(pstate.ssa))
			return
		}
		p := s.Prog(pstate.gc, pstate.storeByType(v.Type))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddrAuto(&p.To, v)
	case ssa.OpS390XLoweredGetClosurePtr:
		// Closure pointer is R12 (already)
		pstate.gc.CheckLoweredGetClosurePtr(v)
	case ssa.OpS390XLoweredRound32F, ssa.OpS390XLoweredRound64F:
	// input is already rounded
	case ssa.OpS390XLoweredGetG:
		r := v.Reg(pstate.ssa)
		p := s.Prog(pstate.gc, s390x.AMOVD)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = s390x.REGG
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XLoweredGetCallerSP:
		// caller's SP is FixedFrameSize below the address of the first arg
		p := s.Prog(pstate.gc, s390x.AMOVD)
		p.From.Type = obj.TYPE_ADDR
		p.From.Offset = -pstate.gc.Ctxt.FixedFrameSize()
		p.From.Name = obj.NAME_PARAM
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)
	case ssa.OpS390XLoweredGetCallerPC:
		p := s.Prog(pstate.gc, obj.AGETCALLERPC)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)
	case ssa.OpS390XCALLstatic, ssa.OpS390XCALLclosure, ssa.OpS390XCALLinter:
		s.Call(pstate.gc, v)
	case ssa.OpS390XLoweredWB:
		p := s.Prog(pstate.gc, obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = v.Aux.(*obj.LSym)
	case ssa.OpS390XFLOGR, ssa.OpS390XNEG, ssa.OpS390XNEGW,
		ssa.OpS390XMOVWBR, ssa.OpS390XMOVDBR:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)
	case ssa.OpS390XNOT, ssa.OpS390XNOTW:
		v.Fatalf("NOT/NOTW generated %s", v.LongString(pstate.ssa))
	case ssa.OpS390XMOVDEQ, ssa.OpS390XMOVDNE,
		ssa.OpS390XMOVDLT, ssa.OpS390XMOVDLE,
		ssa.OpS390XMOVDGT, ssa.OpS390XMOVDGE,
		ssa.OpS390XMOVDGTnoinv, ssa.OpS390XMOVDGEnoinv:
		r := v.Reg(pstate.ssa)
		if r != v.Args[0].Reg(pstate.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(pstate.ssa))
		}
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XFSQRT:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)
	case ssa.OpS390XInvertFlags:
		v.Fatalf("InvertFlags should never make it to codegen %v", v.LongString(pstate.ssa))
	case ssa.OpS390XFlagEQ, ssa.OpS390XFlagLT, ssa.OpS390XFlagGT:
		v.Fatalf("Flag* ops should never make it to codegen %v", v.LongString(pstate.ssa))
	case ssa.OpS390XAddTupleFirst32, ssa.OpS390XAddTupleFirst64:
		v.Fatalf("AddTupleFirst* should never make it to codegen %v", v.LongString(pstate.ssa))
	case ssa.OpS390XLoweredNilCheck:
		// Issue a load which will fault if the input is nil.
		p := s.Prog(pstate.gc, s390x.AMOVBZ)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = s390x.REGTMP
		if pstate.gc.Debug_checknil != 0 && v.Pos.Line() > 1 { // v.Pos.Line()==1 in generated wrappers
			pstate.gc.Warnl(v.Pos, "generated nil check")
		}
	case ssa.OpS390XMVC:
		vo := v.AuxValAndOff(pstate.ssa)
		p := s.Prog(pstate.gc, s390x.AMVC)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = vo.Val()
		p.SetFrom3(obj.Addr{
			Type:   obj.TYPE_MEM,
			Reg:    v.Args[1].Reg(pstate.ssa),
			Offset: vo.Off(),
		})
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Offset = vo.Off()
	case ssa.OpS390XSTMG2, ssa.OpS390XSTMG3, ssa.OpS390XSTMG4,
		ssa.OpS390XSTM2, ssa.OpS390XSTM3, ssa.OpS390XSTM4:
		for i := 2; i < len(v.Args)-1; i++ {
			if v.Args[i].Reg(pstate.ssa) != v.Args[i-1].Reg(pstate.ssa)+1 {
				v.Fatalf("invalid store multiple %s", v.LongString(pstate.ssa))
			}
		}
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(pstate.ssa)
		p.Reg = v.Args[len(v.Args)-2].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.To, v)
	case ssa.OpS390XLoweredMove:
		// Inputs must be valid pointers to memory,
		// so adjust arg0 and arg1 as part of the expansion.
		// arg2 should be src+size,
		//
		// mvc: MVC  $256, 0(R2), 0(R1)
		//      MOVD $256(R1), R1
		//      MOVD $256(R2), R2
		//      CMP  R2, Rarg2
		//      BNE  mvc
		//      MVC  $rem, 0(R2), 0(R1) // if rem > 0
		// arg2 is the last address to move in the loop + 256
		mvc := s.Prog(pstate.gc, s390x.AMVC)
		mvc.From.Type = obj.TYPE_CONST
		mvc.From.Offset = 256
		mvc.SetFrom3(obj.Addr{Type: obj.TYPE_MEM, Reg: v.Args[1].Reg(pstate.ssa)})
		mvc.To.Type = obj.TYPE_MEM
		mvc.To.Reg = v.Args[0].Reg(pstate.ssa)

		for i := 0; i < 2; i++ {
			movd := s.Prog(pstate.gc, s390x.AMOVD)
			movd.From.Type = obj.TYPE_ADDR
			movd.From.Reg = v.Args[i].Reg(pstate.ssa)
			movd.From.Offset = 256
			movd.To.Type = obj.TYPE_REG
			movd.To.Reg = v.Args[i].Reg(pstate.ssa)
		}

		cmpu := s.Prog(pstate.gc, s390x.ACMPU)
		cmpu.From.Reg = v.Args[1].Reg(pstate.ssa)
		cmpu.From.Type = obj.TYPE_REG
		cmpu.To.Reg = v.Args[2].Reg(pstate.ssa)
		cmpu.To.Type = obj.TYPE_REG

		bne := s.Prog(pstate.gc, s390x.ABLT)
		bne.To.Type = obj.TYPE_BRANCH
		pstate.gc.Patch(bne, mvc)

		if v.AuxInt > 0 {
			mvc := s.Prog(pstate.gc, s390x.AMVC)
			mvc.From.Type = obj.TYPE_CONST
			mvc.From.Offset = v.AuxInt
			mvc.SetFrom3(obj.Addr{Type: obj.TYPE_MEM, Reg: v.Args[1].Reg(pstate.ssa)})
			mvc.To.Type = obj.TYPE_MEM
			mvc.To.Reg = v.Args[0].Reg(pstate.ssa)
		}
	case ssa.OpS390XLoweredZero:
		// Input must be valid pointers to memory,
		// so adjust arg0 as part of the expansion.
		// arg1 should be src+size,
		//
		// clear: CLEAR $256, 0(R1)
		//        MOVD  $256(R1), R1
		//        CMP   R1, Rarg1
		//        BNE   clear
		//        CLEAR $rem, 0(R1) // if rem > 0
		// arg1 is the last address to zero in the loop + 256
		clear := s.Prog(pstate.gc, s390x.ACLEAR)
		clear.From.Type = obj.TYPE_CONST
		clear.From.Offset = 256
		clear.To.Type = obj.TYPE_MEM
		clear.To.Reg = v.Args[0].Reg(pstate.ssa)

		movd := s.Prog(pstate.gc, s390x.AMOVD)
		movd.From.Type = obj.TYPE_ADDR
		movd.From.Reg = v.Args[0].Reg(pstate.ssa)
		movd.From.Offset = 256
		movd.To.Type = obj.TYPE_REG
		movd.To.Reg = v.Args[0].Reg(pstate.ssa)

		cmpu := s.Prog(pstate.gc, s390x.ACMPU)
		cmpu.From.Reg = v.Args[0].Reg(pstate.ssa)
		cmpu.From.Type = obj.TYPE_REG
		cmpu.To.Reg = v.Args[1].Reg(pstate.ssa)
		cmpu.To.Type = obj.TYPE_REG

		bne := s.Prog(pstate.gc, s390x.ABLT)
		bne.To.Type = obj.TYPE_BRANCH
		pstate.gc.Patch(bne, clear)

		if v.AuxInt > 0 {
			clear := s.Prog(pstate.gc, s390x.ACLEAR)
			clear.From.Type = obj.TYPE_CONST
			clear.From.Offset = v.AuxInt
			clear.To.Type = obj.TYPE_MEM
			clear.To.Reg = v.Args[0].Reg(pstate.ssa)
		}
	case ssa.OpS390XMOVWZatomicload, ssa.OpS390XMOVDatomicload:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0(pstate.ssa)
	case ssa.OpS390XMOVWatomicstore, ssa.OpS390XMOVDatomicstore:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.To, v)
	case ssa.OpS390XLAA, ssa.OpS390XLAAG:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.Reg = v.Reg0(pstate.ssa)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.To, v)
	case ssa.OpS390XLoweredAtomicCas32, ssa.OpS390XLoweredAtomicCas64:
		// Convert the flags output of CS{,G} into a bool.
		//    CS{,G} arg1, arg2, arg0
		//    MOVD   $0, ret
		//    BNE    2(PC)
		//    MOVD   $1, ret
		//    NOP (so the BNE has somewhere to land)

		// CS{,G} arg1, arg2, arg0
		cs := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		cs.From.Type = obj.TYPE_REG
		cs.From.Reg = v.Args[1].Reg(pstate.ssa) // old
		cs.Reg = v.Args[2].Reg(pstate.ssa)      // new
		cs.To.Type = obj.TYPE_MEM
		cs.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&cs.To, v)

		// MOVD $0, ret
		movd := s.Prog(pstate.gc, s390x.AMOVD)
		movd.From.Type = obj.TYPE_CONST
		movd.From.Offset = 0
		movd.To.Type = obj.TYPE_REG
		movd.To.Reg = v.Reg0(pstate.ssa)

		// BNE 2(PC)
		bne := s.Prog(pstate.gc, s390x.ABNE)
		bne.To.Type = obj.TYPE_BRANCH

		// MOVD $1, ret
		movd = s.Prog(pstate.gc, s390x.AMOVD)
		movd.From.Type = obj.TYPE_CONST
		movd.From.Offset = 1
		movd.To.Type = obj.TYPE_REG
		movd.To.Reg = v.Reg0(pstate.ssa)

		// NOP (so the BNE has somewhere to land)
		nop := s.Prog(pstate.gc, obj.ANOP)
		pstate.gc.Patch(bne, nop)
	case ssa.OpS390XLoweredAtomicExchange32, ssa.OpS390XLoweredAtomicExchange64:
		// Loop until the CS{,G} succeeds.
		//     MOV{WZ,D} arg0, ret
		// cs: CS{,G}    ret, arg1, arg0
		//     BNE       cs

		// MOV{WZ,D} arg0, ret
		load := s.Prog(pstate.gc, pstate.loadByType(v.Type.FieldType(pstate.types, 0)))
		load.From.Type = obj.TYPE_MEM
		load.From.Reg = v.Args[0].Reg(pstate.ssa)
		load.To.Type = obj.TYPE_REG
		load.To.Reg = v.Reg0(pstate.ssa)
		pstate.gc.AddAux(&load.From, v)

		// CS{,G} ret, arg1, arg0
		cs := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		cs.From.Type = obj.TYPE_REG
		cs.From.Reg = v.Reg0(pstate.ssa)   // old
		cs.Reg = v.Args[1].Reg(pstate.ssa) // new
		cs.To.Type = obj.TYPE_MEM
		cs.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&cs.To, v)

		// BNE cs
		bne := s.Prog(pstate.gc, s390x.ABNE)
		bne.To.Type = obj.TYPE_BRANCH
		pstate.gc.Patch(bne, cs)
	case ssa.OpClobber:
	// TODO: implement for clobberdead experiment. Nop is ok for now.
	default:
		v.Fatalf("genValue not implemented: %s", v.LongString(pstate.ssa))
	}
}

func (pstate *PackageState) ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	switch b.Kind {
	case ssa.BlockPlain:
		if b.Succs[0].Block() != next {
			p := s.Prog(pstate.gc, s390x.ABR)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}
	case ssa.BlockDefer:
		// defer returns in R3:
		// 0 if we should continue executing
		// 1 if we should jump to deferreturn call
		p := s.Prog(pstate.gc, s390x.ACMPW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = s390x.REG_R3
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = 0
		p = s.Prog(pstate.gc, s390x.ABNE)
		p.To.Type = obj.TYPE_BRANCH
		s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[1].Block()})
		if b.Succs[0].Block() != next {
			p := s.Prog(pstate.gc, s390x.ABR)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}
	case ssa.BlockExit:
		s.Prog(pstate.gc, obj.AUNDEF) // tell plive.go that we never reach here
	case ssa.BlockRet:
		s.Prog(pstate.gc, obj.ARET)
	case ssa.BlockRetJmp:
		p := s.Prog(pstate.gc, s390x.ABR)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = b.Aux.(*obj.LSym)
	case ssa.BlockS390XEQ, ssa.BlockS390XNE,
		ssa.BlockS390XLT, ssa.BlockS390XGE,
		ssa.BlockS390XLE, ssa.BlockS390XGT,
		ssa.BlockS390XGEF, ssa.BlockS390XGTF:
		jmp := pstate.blockJump[b.Kind]
		switch next {
		case b.Succs[0].Block():
			s.Br(pstate.gc, jmp.invasm, b.Succs[1].Block())
		case b.Succs[1].Block():
			s.Br(pstate.gc, jmp.asm, b.Succs[0].Block())
		default:
			if b.Likely != ssa.BranchUnlikely {
				s.Br(pstate.gc, jmp.asm, b.Succs[0].Block())
				s.Br(pstate.gc, s390x.ABR, b.Succs[1].Block())
			} else {
				s.Br(pstate.gc, jmp.invasm, b.Succs[1].Block())
				s.Br(pstate.gc, s390x.ABR, b.Succs[0].Block())
			}
		}
	default:
		b.Fatalf("branch not implemented: %s. Control: %s", b.LongString(pstate.ssa), b.Control.LongString(pstate.ssa))
	}
}

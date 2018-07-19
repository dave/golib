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
func (psess *PackageSession) ssaMarkMoves(s *gc.SSAGenState, b *ssa.Block) {
	flive := b.FlagsLiveAtEnd
	if b.Control != nil && b.Control.Type.IsFlags(psess.types) {
		flive = true
	}
	for i := len(b.Values) - 1; i >= 0; i-- {
		v := b.Values[i]
		if flive && v.Op == ssa.OpS390XMOVDconst {

			v.Aux = v
		}
		if v.Type.IsFlags(psess.types) {
			flive = false
		}
		for _, a := range v.Args {
			if a.Type.IsFlags(psess.types) {
				flive = true
			}
		}
	}
}

// loadByType returns the load instruction of the given type.
func (psess *PackageSession) loadByType(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(psess.types) {
		case 4:
			return s390x.AFMOVS
		case 8:
			return s390x.AFMOVD
		}
	} else {
		switch t.Size(psess.types) {
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
func (psess *PackageSession) storeByType(t *types.Type) obj.As {
	width := t.Size(psess.types)
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
func (psess *PackageSession) moveByType(t *types.Type) obj.As {
	if t.IsFloat() {
		return s390x.AFMOVD
	} else {
		switch t.Size(psess.types) {
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
func (psess *PackageSession) opregreg(s *gc.SSAGenState, op obj.As, dest, src int16) *obj.Prog {
	p := s.Prog(psess.gc, op)
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
func (psess *PackageSession) opregregimm(s *gc.SSAGenState, op obj.As, dest, src int16, off int64) *obj.Prog {
	p := s.Prog(psess.gc, op)
	p.From.Type = obj.TYPE_CONST
	p.From.Offset = off
	p.Reg = src
	p.To.Reg = dest
	p.To.Type = obj.TYPE_REG
	return p
}

func (psess *PackageSession) ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpS390XSLD, ssa.OpS390XSLW,
		ssa.OpS390XSRD, ssa.OpS390XSRW,
		ssa.OpS390XSRAD, ssa.OpS390XSRAW:
		r := v.Reg(psess.ssa)
		r1 := v.Args[0].Reg(psess.ssa)
		r2 := v.Args[1].Reg(psess.ssa)
		if r2 == s390x.REG_R0 {
			v.Fatalf("cannot use R0 as shift value %s", v.LongString(psess.ssa))
		}
		p := psess.opregreg(s, v.Op.Asm(psess.ssa), r, r2)
		if r != r1 {
			p.Reg = r1
		}
	case ssa.OpS390XADD, ssa.OpS390XADDW,
		ssa.OpS390XSUB, ssa.OpS390XSUBW,
		ssa.OpS390XAND, ssa.OpS390XANDW,
		ssa.OpS390XOR, ssa.OpS390XORW,
		ssa.OpS390XXOR, ssa.OpS390XXORW:
		r := v.Reg(psess.ssa)
		r1 := v.Args[0].Reg(psess.ssa)
		r2 := v.Args[1].Reg(psess.ssa)
		p := psess.opregreg(s, v.Op.Asm(psess.ssa), r, r2)
		if r != r1 {
			p.Reg = r1
		}

	case ssa.OpS390XMULLD, ssa.OpS390XMULLW,
		ssa.OpS390XMULHD, ssa.OpS390XMULHDU,
		ssa.OpS390XFADDS, ssa.OpS390XFADD, ssa.OpS390XFSUBS, ssa.OpS390XFSUB,
		ssa.OpS390XFMULS, ssa.OpS390XFMUL, ssa.OpS390XFDIVS, ssa.OpS390XFDIV:
		r := v.Reg(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		psess.
			opregreg(s, v.Op.Asm(psess.ssa), r, v.Args[1].Reg(psess.ssa))
	case ssa.OpS390XFMADD, ssa.OpS390XFMADDS,
		ssa.OpS390XFMSUB, ssa.OpS390XFMSUBS:
		r := v.Reg(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		r1 := v.Args[1].Reg(psess.ssa)
		r2 := v.Args[2].Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.Reg = r2
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XFIDBR:
		switch v.AuxInt {
		case 0, 1, 3, 4, 5, 6, 7:
			psess.
				opregregimm(s, v.Op.Asm(psess.ssa), v.Reg(psess.ssa), v.Args[0].Reg(psess.ssa), v.AuxInt)
		default:
			v.Fatalf("invalid FIDBR mask: %v", v.AuxInt)
		}
	case ssa.OpS390XCPSDR:
		p := psess.opregreg(s, v.Op.Asm(psess.ssa), v.Reg(psess.ssa), v.Args[1].Reg(psess.ssa))
		p.Reg = v.Args[0].Reg(psess.ssa)
	case ssa.OpS390XDIVD, ssa.OpS390XDIVW,
		ssa.OpS390XDIVDU, ssa.OpS390XDIVWU,
		ssa.OpS390XMODD, ssa.OpS390XMODW,
		ssa.OpS390XMODDU, ssa.OpS390XMODWU:

		dividend := v.Args[0].Reg(psess.ssa)
		divisor := v.Args[1].Reg(psess.

			// CPU faults upon signed overflow, which occurs when most
			// negative int is divided by -1.
			ssa)

		var j *obj.Prog
		if v.Op == ssa.OpS390XDIVD || v.Op == ssa.OpS390XDIVW ||
			v.Op == ssa.OpS390XMODD || v.Op == ssa.OpS390XMODW {

			var c *obj.Prog
			c = s.Prog(psess.gc, s390x.ACMP)
			j = s.Prog(psess.gc, s390x.ABEQ)

			c.From.Type = obj.TYPE_REG
			c.From.Reg = divisor
			c.To.Type = obj.TYPE_CONST
			c.To.Offset = -1

			j.To.Type = obj.TYPE_BRANCH

		}

		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = divisor
		p.Reg = 0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = dividend

		if j != nil {
			j2 := s.Prog(psess.gc, s390x.ABR)
			j2.To.Type = obj.TYPE_BRANCH

			var n *obj.Prog
			if v.Op == ssa.OpS390XDIVD || v.Op == ssa.OpS390XDIVW {

				n = s.Prog(psess.gc, s390x.ANEG)
				n.To.Type = obj.TYPE_REG
				n.To.Reg = dividend
			} else {

				n = s.Prog(psess.gc, s390x.AXOR)
				n.From.Type = obj.TYPE_REG
				n.From.Reg = dividend
				n.To.Type = obj.TYPE_REG
				n.To.Reg = dividend
			}

			j.To.Val = n
			j2.To.Val = s.Pc()
		}
	case ssa.OpS390XADDconst, ssa.OpS390XADDWconst:
		psess.
			opregregimm(s, v.Op.Asm(psess.ssa), v.Reg(psess.ssa), v.Args[0].Reg(psess.ssa), v.AuxInt)
	case ssa.OpS390XMULLDconst, ssa.OpS390XMULLWconst,
		ssa.OpS390XSUBconst, ssa.OpS390XSUBWconst,
		ssa.OpS390XANDconst, ssa.OpS390XANDWconst,
		ssa.OpS390XORconst, ssa.OpS390XORWconst,
		ssa.OpS390XXORconst, ssa.OpS390XXORWconst:
		r := v.Reg(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XSLDconst, ssa.OpS390XSLWconst,
		ssa.OpS390XSRDconst, ssa.OpS390XSRWconst,
		ssa.OpS390XSRADconst, ssa.OpS390XSRAWconst,
		ssa.OpS390XRLLGconst, ssa.OpS390XRLLconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		r := v.Reg(psess.ssa)
		r1 := v.Args[0].Reg(psess.ssa)
		if r != r1 {
			p.Reg = r1
		}
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XMOVDaddridx:
		r := v.Args[0].Reg(psess.ssa)
		i := v.Args[1].Reg(psess.ssa)
		p := s.Prog(psess.gc, s390x.AMOVD)
		p.From.Scale = 1
		if i == s390x.REGSP {
			r, i = i, r
		}
		p.From.Type = obj.TYPE_ADDR
		p.From.Reg = r
		p.From.Index = i
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpS390XMOVDaddr:
		p := s.Prog(psess.gc, s390x.AMOVD)
		p.From.Type = obj.TYPE_ADDR
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpS390XCMP, ssa.OpS390XCMPW, ssa.OpS390XCMPU, ssa.OpS390XCMPWU:
		psess.
			opregreg(s, v.Op.Asm(psess.ssa), v.Args[1].Reg(psess.ssa), v.Args[0].Reg(psess.ssa))
	case ssa.OpS390XFCMPS, ssa.OpS390XFCMP:
		psess.
			opregreg(s, v.Op.Asm(psess.ssa), v.Args[1].Reg(psess.ssa), v.Args[0].Reg(psess.ssa))
	case ssa.OpS390XCMPconst, ssa.OpS390XCMPWconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = v.AuxInt
	case ssa.OpS390XCMPUconst, ssa.OpS390XCMPWUconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = int64(uint32(v.AuxInt))
	case ssa.OpS390XMOVDconst:
		x := v.Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x
	case ssa.OpS390XFMOVSconst, ssa.OpS390XFMOVDconst:
		x := v.Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
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
		r := v.Reg(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XMOVDload,
		ssa.OpS390XMOVWZload, ssa.OpS390XMOVHZload, ssa.OpS390XMOVBZload,
		ssa.OpS390XMOVDBRload, ssa.OpS390XMOVWBRload, ssa.OpS390XMOVHBRload,
		ssa.OpS390XMOVBload, ssa.OpS390XMOVHload, ssa.OpS390XMOVWload,
		ssa.OpS390XFMOVSload, ssa.OpS390XFMOVDload:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpS390XMOVBZloadidx, ssa.OpS390XMOVHZloadidx, ssa.OpS390XMOVWZloadidx,
		ssa.OpS390XMOVBloadidx, ssa.OpS390XMOVHloadidx, ssa.OpS390XMOVWloadidx, ssa.OpS390XMOVDloadidx,
		ssa.OpS390XMOVHBRloadidx, ssa.OpS390XMOVWBRloadidx, ssa.OpS390XMOVDBRloadidx,
		ssa.OpS390XFMOVSloadidx, ssa.OpS390XFMOVDloadidx:
		r := v.Args[0].Reg(psess.ssa)
		i := v.Args[1].Reg(psess.ssa)
		if i == s390x.REGSP {
			r, i = i, r
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r
		p.From.Scale = 1
		p.From.Index = i
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpS390XMOVBstore, ssa.OpS390XMOVHstore, ssa.OpS390XMOVWstore, ssa.OpS390XMOVDstore,
		ssa.OpS390XMOVHBRstore, ssa.OpS390XMOVWBRstore, ssa.OpS390XMOVDBRstore,
		ssa.OpS390XFMOVSstore, ssa.OpS390XFMOVDstore:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpS390XMOVBstoreidx, ssa.OpS390XMOVHstoreidx, ssa.OpS390XMOVWstoreidx, ssa.OpS390XMOVDstoreidx,
		ssa.OpS390XMOVHBRstoreidx, ssa.OpS390XMOVWBRstoreidx, ssa.OpS390XMOVDBRstoreidx,
		ssa.OpS390XFMOVSstoreidx, ssa.OpS390XFMOVDstoreidx:
		r := v.Args[0].Reg(psess.ssa)
		i := v.Args[1].Reg(psess.ssa)
		if i == s390x.REGSP {
			r, i = i, r
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[2].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = r
		p.To.Scale = 1
		p.To.Index = i
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpS390XMOVDstoreconst, ssa.OpS390XMOVWstoreconst, ssa.OpS390XMOVHstoreconst, ssa.OpS390XMOVBstoreconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		sc := v.AuxValAndOff(psess.ssa)
		p.From.Offset = sc.Val()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux2(&p.To, v, sc.Off())
	case ssa.OpS390XMOVBreg, ssa.OpS390XMOVHreg, ssa.OpS390XMOVWreg,
		ssa.OpS390XMOVBZreg, ssa.OpS390XMOVHZreg, ssa.OpS390XMOVWZreg,
		ssa.OpS390XLDGR, ssa.OpS390XLGDR,
		ssa.OpS390XCEFBRA, ssa.OpS390XCDFBRA, ssa.OpS390XCEGBRA, ssa.OpS390XCDGBRA,
		ssa.OpS390XCFEBRA, ssa.OpS390XCFDBRA, ssa.OpS390XCGEBRA, ssa.OpS390XCGDBRA,
		ssa.OpS390XLDEBR, ssa.OpS390XLEDBR,
		ssa.OpS390XFNEG, ssa.OpS390XFNEGS,
		ssa.OpS390XLPDFR, ssa.OpS390XLNDFR:
		psess.
			opregreg(s, v.Op.Asm(psess.ssa), v.Reg(psess.ssa), v.Args[0].Reg(psess.ssa))
	case ssa.OpS390XCLEAR:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		sc := v.AuxValAndOff(psess.ssa)
		p.From.Offset = sc.Val()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux2(&p.To, v, sc.Off())
	case ssa.OpCopy, ssa.OpS390XMOVDreg:
		if v.Type.IsMemory(psess.types) {
			return
		}
		x := v.Args[0].Reg(psess.ssa)
		y := v.Reg(psess.ssa)
		if x != y {
			psess.
				opregreg(s, psess.moveByType(v.Type), y, x)
		}
	case ssa.OpS390XMOVDnop:
		if v.Reg(psess.ssa) != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}

	case ssa.OpLoadReg:
		if v.Type.IsFlags(psess.types) {
			v.Fatalf("load flags not implemented: %v", v.LongString(psess.ssa))
			return
		}
		p := s.Prog(psess.gc, psess.loadByType(v.Type))
		psess.gc.
			AddrAuto(&p.From, v.Args[0])
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpStoreReg:
		if v.Type.IsFlags(psess.types) {
			v.Fatalf("store flags not implemented: %v", v.LongString(psess.ssa))
			return
		}
		p := s.Prog(psess.gc, psess.storeByType(v.Type))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddrAuto(&p.To, v)
	case ssa.OpS390XLoweredGetClosurePtr:
		psess.gc.
			CheckLoweredGetClosurePtr(v)
	case ssa.OpS390XLoweredRound32F, ssa.OpS390XLoweredRound64F:

	case ssa.OpS390XLoweredGetG:
		r := v.Reg(psess.ssa)
		p := s.Prog(psess.gc, s390x.AMOVD)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = s390x.REGG
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XLoweredGetCallerSP:

		p := s.Prog(psess.gc, s390x.AMOVD)
		p.From.Type = obj.TYPE_ADDR
		p.From.Offset = -psess.gc.Ctxt.FixedFrameSize()
		p.From.Name = obj.NAME_PARAM
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpS390XLoweredGetCallerPC:
		p := s.Prog(psess.gc, obj.AGETCALLERPC)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpS390XCALLstatic, ssa.OpS390XCALLclosure, ssa.OpS390XCALLinter:
		s.Call(psess.gc, v)
	case ssa.OpS390XLoweredWB:
		p := s.Prog(psess.gc, obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = v.Aux.(*obj.LSym)
	case ssa.OpS390XFLOGR, ssa.OpS390XNEG, ssa.OpS390XNEGW,
		ssa.OpS390XMOVWBR, ssa.OpS390XMOVDBR:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpS390XNOT, ssa.OpS390XNOTW:
		v.Fatalf("NOT/NOTW generated %s", v.LongString(psess.ssa))
	case ssa.OpS390XMOVDEQ, ssa.OpS390XMOVDNE,
		ssa.OpS390XMOVDLT, ssa.OpS390XMOVDLE,
		ssa.OpS390XMOVDGT, ssa.OpS390XMOVDGE,
		ssa.OpS390XMOVDGTnoinv, ssa.OpS390XMOVDGEnoinv:
		r := v.Reg(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpS390XFSQRT:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpS390XInvertFlags:
		v.Fatalf("InvertFlags should never make it to codegen %v", v.LongString(psess.ssa))
	case ssa.OpS390XFlagEQ, ssa.OpS390XFlagLT, ssa.OpS390XFlagGT:
		v.Fatalf("Flag* ops should never make it to codegen %v", v.LongString(psess.ssa))
	case ssa.OpS390XAddTupleFirst32, ssa.OpS390XAddTupleFirst64:
		v.Fatalf("AddTupleFirst* should never make it to codegen %v", v.LongString(psess.ssa))
	case ssa.OpS390XLoweredNilCheck:

		p := s.Prog(psess.gc, s390x.AMOVBZ)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = s390x.REGTMP
		if psess.gc.Debug_checknil != 0 && v.Pos.Line() > 1 {
			psess.gc.
				Warnl(v.Pos, "generated nil check")
		}
	case ssa.OpS390XMVC:
		vo := v.AuxValAndOff(psess.ssa)
		p := s.Prog(psess.gc, s390x.AMVC)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = vo.Val()
		p.SetFrom3(obj.Addr{
			Type:   obj.TYPE_MEM,
			Reg:    v.Args[1].Reg(psess.ssa),
			Offset: vo.Off(),
		})
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Offset = vo.Off()
	case ssa.OpS390XSTMG2, ssa.OpS390XSTMG3, ssa.OpS390XSTMG4,
		ssa.OpS390XSTM2, ssa.OpS390XSTM3, ssa.OpS390XSTM4:
		for i := 2; i < len(v.Args)-1; i++ {
			if v.Args[i].Reg(psess.ssa) != v.Args[i-1].Reg(psess.ssa)+1 {
				v.Fatalf("invalid store multiple %s", v.LongString(psess.ssa))
			}
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.Reg = v.Args[len(v.Args)-2].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpS390XLoweredMove:

		mvc := s.Prog(psess.gc, s390x.AMVC)
		mvc.From.Type = obj.TYPE_CONST
		mvc.From.Offset = 256
		mvc.SetFrom3(obj.Addr{Type: obj.TYPE_MEM, Reg: v.Args[1].Reg(psess.ssa)})
		mvc.To.Type = obj.TYPE_MEM
		mvc.To.Reg = v.Args[0].Reg(psess.ssa)

		for i := 0; i < 2; i++ {
			movd := s.Prog(psess.gc, s390x.AMOVD)
			movd.From.Type = obj.TYPE_ADDR
			movd.From.Reg = v.Args[i].Reg(psess.ssa)
			movd.From.Offset = 256
			movd.To.Type = obj.TYPE_REG
			movd.To.Reg = v.Args[i].Reg(psess.ssa)
		}

		cmpu := s.Prog(psess.gc, s390x.ACMPU)
		cmpu.From.Reg = v.Args[1].Reg(psess.ssa)
		cmpu.From.Type = obj.TYPE_REG
		cmpu.To.Reg = v.Args[2].Reg(psess.ssa)
		cmpu.To.Type = obj.TYPE_REG

		bne := s.Prog(psess.gc, s390x.ABLT)
		bne.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(bne, mvc)

		if v.AuxInt > 0 {
			mvc := s.Prog(psess.gc, s390x.AMVC)
			mvc.From.Type = obj.TYPE_CONST
			mvc.From.Offset = v.AuxInt
			mvc.SetFrom3(obj.Addr{Type: obj.TYPE_MEM, Reg: v.Args[1].Reg(psess.ssa)})
			mvc.To.Type = obj.TYPE_MEM
			mvc.To.Reg = v.Args[0].Reg(psess.ssa)
		}
	case ssa.OpS390XLoweredZero:

		clear := s.Prog(psess.gc, s390x.ACLEAR)
		clear.From.Type = obj.TYPE_CONST
		clear.From.Offset = 256
		clear.To.Type = obj.TYPE_MEM
		clear.To.Reg = v.Args[0].Reg(psess.ssa)

		movd := s.Prog(psess.gc, s390x.AMOVD)
		movd.From.Type = obj.TYPE_ADDR
		movd.From.Reg = v.Args[0].Reg(psess.ssa)
		movd.From.Offset = 256
		movd.To.Type = obj.TYPE_REG
		movd.To.Reg = v.Args[0].Reg(psess.ssa)

		cmpu := s.Prog(psess.gc, s390x.ACMPU)
		cmpu.From.Reg = v.Args[0].Reg(psess.ssa)
		cmpu.From.Type = obj.TYPE_REG
		cmpu.To.Reg = v.Args[1].Reg(psess.ssa)
		cmpu.To.Type = obj.TYPE_REG

		bne := s.Prog(psess.gc, s390x.ABLT)
		bne.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(bne, clear)

		if v.AuxInt > 0 {
			clear := s.Prog(psess.gc, s390x.ACLEAR)
			clear.From.Type = obj.TYPE_CONST
			clear.From.Offset = v.AuxInt
			clear.To.Type = obj.TYPE_MEM
			clear.To.Reg = v.Args[0].Reg(psess.ssa)
		}
	case ssa.OpS390XMOVWZatomicload, ssa.OpS390XMOVDatomicload:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0(psess.ssa)
	case ssa.OpS390XMOVWatomicstore, ssa.OpS390XMOVDatomicstore:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpS390XLAA, ssa.OpS390XLAAG:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.Reg = v.Reg0(psess.ssa)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpS390XLoweredAtomicCas32, ssa.OpS390XLoweredAtomicCas64:

		cs := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		cs.From.Type = obj.TYPE_REG
		cs.From.Reg = v.Args[1].Reg(psess.ssa)
		cs.Reg = v.Args[2].Reg(psess.ssa)
		cs.To.Type = obj.TYPE_MEM
		cs.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&cs.To, v)

		movd := s.Prog(psess.gc, s390x.AMOVD)
		movd.From.Type = obj.TYPE_CONST
		movd.From.Offset = 0
		movd.To.Type = obj.TYPE_REG
		movd.To.Reg = v.Reg0(psess.ssa)

		bne := s.Prog(psess.gc, s390x.ABNE)
		bne.To.Type = obj.TYPE_BRANCH

		movd = s.Prog(psess.gc, s390x.AMOVD)
		movd.From.Type = obj.TYPE_CONST
		movd.From.Offset = 1
		movd.To.Type = obj.TYPE_REG
		movd.To.Reg = v.Reg0(psess.ssa)

		nop := s.Prog(psess.gc, obj.ANOP)
		psess.gc.
			Patch(bne, nop)
	case ssa.OpS390XLoweredAtomicExchange32, ssa.OpS390XLoweredAtomicExchange64:

		load := s.Prog(psess.gc, psess.loadByType(v.Type.FieldType(psess.types, 0)))
		load.From.Type = obj.TYPE_MEM
		load.From.Reg = v.Args[0].Reg(psess.ssa)
		load.To.Type = obj.TYPE_REG
		load.To.Reg = v.Reg0(psess.ssa)
		psess.gc.
			AddAux(&load.From, v)

		cs := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		cs.From.Type = obj.TYPE_REG
		cs.From.Reg = v.Reg0(psess.ssa)
		cs.Reg = v.Args[1].Reg(psess.ssa)
		cs.To.Type = obj.TYPE_MEM
		cs.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&cs.To, v)

		bne := s.Prog(psess.gc, s390x.ABNE)
		bne.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(bne, cs)
	case ssa.OpClobber:

	default:
		v.Fatalf("genValue not implemented: %s", v.LongString(psess.ssa))
	}
}

func (psess *PackageSession) ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	switch b.Kind {
	case ssa.BlockPlain:
		if b.Succs[0].Block() != next {
			p := s.Prog(psess.gc, s390x.ABR)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}
	case ssa.BlockDefer:

		p := s.Prog(psess.gc, s390x.ACMPW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = s390x.REG_R3
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = 0
		p = s.Prog(psess.gc, s390x.ABNE)
		p.To.Type = obj.TYPE_BRANCH
		s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[1].Block()})
		if b.Succs[0].Block() != next {
			p := s.Prog(psess.gc, s390x.ABR)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}
	case ssa.BlockExit:
		s.Prog(psess.gc, obj.AUNDEF)
	case ssa.BlockRet:
		s.Prog(psess.gc, obj.ARET)
	case ssa.BlockRetJmp:
		p := s.Prog(psess.gc, s390x.ABR)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = b.Aux.(*obj.LSym)
	case ssa.BlockS390XEQ, ssa.BlockS390XNE,
		ssa.BlockS390XLT, ssa.BlockS390XGE,
		ssa.BlockS390XLE, ssa.BlockS390XGT,
		ssa.BlockS390XGEF, ssa.BlockS390XGTF:
		jmp := psess.blockJump[b.Kind]
		switch next {
		case b.Succs[0].Block():
			s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
		case b.Succs[1].Block():
			s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
		default:
			if b.Likely != ssa.BranchUnlikely {
				s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
				s.Br(psess.gc, s390x.ABR, b.Succs[1].Block())
			} else {
				s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
				s.Br(psess.gc, s390x.ABR, b.Succs[0].Block())
			}
		}
	default:
		b.Fatalf("branch not implemented: %s. Control: %s", b.LongString(psess.ssa), b.Control.LongString(psess.ssa))
	}
}

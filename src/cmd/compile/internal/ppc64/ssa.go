package ppc64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/ppc64"
	"math"
	"strings"
)

// iselOp encodes mapping of comparison operations onto ISEL operands
type iselOp struct {
	cond        int64
	valueIfCond int // if cond is true, the value to return (0 or 1)
}

// Input registers to ISEL used for comparison. Index 0 is zero, 1 is (will be) 1

// markMoves marks any MOVXconst ops that need to avoid clobbering flags.
func ssaMarkMoves(s *gc.SSAGenState, b *ssa.Block) {

}

// loadByType returns the load instruction of the given type.
func (psess *PackageSession) loadByType(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(psess.types) {
		case 4:
			return ppc64.AFMOVS
		case 8:
			return ppc64.AFMOVD
		}
	} else {
		switch t.Size(psess.types) {
		case 1:
			if t.IsSigned() {
				return ppc64.AMOVB
			} else {
				return ppc64.AMOVBZ
			}
		case 2:
			if t.IsSigned() {
				return ppc64.AMOVH
			} else {
				return ppc64.AMOVHZ
			}
		case 4:
			if t.IsSigned() {
				return ppc64.AMOVW
			} else {
				return ppc64.AMOVWZ
			}
		case 8:
			return ppc64.AMOVD
		}
	}
	panic("bad load type")
}

// storeByType returns the store instruction of the given type.
func (psess *PackageSession) storeByType(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(psess.types) {
		case 4:
			return ppc64.AFMOVS
		case 8:
			return ppc64.AFMOVD
		}
	} else {
		switch t.Size(psess.types) {
		case 1:
			return ppc64.AMOVB
		case 2:
			return ppc64.AMOVH
		case 4:
			return ppc64.AMOVW
		case 8:
			return ppc64.AMOVD
		}
	}
	panic("bad store type")
}

func (psess *PackageSession) ssaGenISEL(s *gc.SSAGenState, v *ssa.Value, cr int64, r1, r2 int16) {
	r := v.Reg(psess.ssa)
	p := s.Prog(psess.gc, ppc64.AISEL)
	p.To.Type = obj.TYPE_REG
	p.To.Reg = r
	p.Reg = r1
	p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: r2})
	p.From.Type = obj.TYPE_CONST
	p.From.Offset = cr
}

func (psess *PackageSession) ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpCopy:
		t := v.Type
		if t.IsMemory(psess.types) {
			return
		}
		x := v.Args[0].Reg(psess.ssa)
		y := v.Reg(psess.ssa)
		if x != y {
			rt := obj.TYPE_REG
			op := ppc64.AMOVD

			if t.IsFloat() {
				op = ppc64.AFMOVD
			}
			p := s.Prog(psess.gc, op)
			p.From.Type = rt
			p.From.Reg = x
			p.To.Type = rt
			p.To.Reg = y
		}

	case ssa.OpPPC64LoweredAtomicAnd8,
		ssa.OpPPC64LoweredAtomicOr8:

		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)

		plwsync := s.Prog(psess.gc, ppc64.ALWSYNC)
		plwsync.To.Type = obj.TYPE_NONE
		p := s.Prog(psess.gc, ppc64.ALBAR)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REGTMP
		p1 := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = ppc64.REGTMP
		p2 := s.Prog(psess.gc, ppc64.ASTBCCC)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = ppc64.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = r0
		p2.RegTo2 = ppc64.REGTMP
		p3 := s.Prog(psess.gc, ppc64.ABNE)
		p3.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p3, p)

	case ssa.OpPPC64LoweredAtomicAdd32,
		ssa.OpPPC64LoweredAtomicAdd64:

		ld := ppc64.ALDAR
		st := ppc64.ASTDCCC
		if v.Op == ssa.OpPPC64LoweredAtomicAdd32 {
			ld = ppc64.ALWAR
			st = ppc64.ASTWCCC
		}
		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)
		out := v.Reg0(psess.ssa)

		plwsync := s.Prog(psess.gc, ppc64.ALWSYNC)
		plwsync.To.Type = obj.TYPE_NONE

		p := s.Prog(psess.gc, ld)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = out

		p1 := s.Prog(psess.gc, ppc64.AADD)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.To.Reg = out
		p1.To.Type = obj.TYPE_REG

		p3 := s.Prog(psess.gc, st)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = out
		p3.To.Type = obj.TYPE_MEM
		p3.To.Reg = r0

		p4 := s.Prog(psess.gc, ppc64.ABNE)
		p4.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p4, p)

		if v.Op == ssa.OpPPC64LoweredAtomicAdd32 {
			p5 := s.Prog(psess.gc, ppc64.AMOVWZ)
			p5.To.Type = obj.TYPE_REG
			p5.To.Reg = out
			p5.From.Type = obj.TYPE_REG
			p5.From.Reg = out
		}

	case ssa.OpPPC64LoweredAtomicExchange32,
		ssa.OpPPC64LoweredAtomicExchange64:

		ld := ppc64.ALDAR
		st := ppc64.ASTDCCC
		if v.Op == ssa.OpPPC64LoweredAtomicExchange32 {
			ld = ppc64.ALWAR
			st = ppc64.ASTWCCC
		}
		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)
		out := v.Reg0(psess.ssa)

		plwsync := s.Prog(psess.gc, ppc64.ALWSYNC)
		plwsync.To.Type = obj.TYPE_NONE

		p := s.Prog(psess.gc, ld)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = out

		p1 := s.Prog(psess.gc, st)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.To.Type = obj.TYPE_MEM
		p1.To.Reg = r0

		p2 := s.Prog(psess.gc, ppc64.ABNE)
		p2.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p2, p)

		pisync := s.Prog(psess.gc, ppc64.AISYNC)
		pisync.To.Type = obj.TYPE_NONE

	case ssa.OpPPC64LoweredAtomicLoad32,
		ssa.OpPPC64LoweredAtomicLoad64,
		ssa.OpPPC64LoweredAtomicLoadPtr:

		ld := ppc64.AMOVD
		cmp := ppc64.ACMP
		if v.Op == ssa.OpPPC64LoweredAtomicLoad32 {
			ld = ppc64.AMOVW
			cmp = ppc64.ACMPW
		}
		arg0 := v.Args[0].Reg(psess.ssa)
		out := v.Reg0(psess.ssa)

		psync := s.Prog(psess.gc, ppc64.ASYNC)
		psync.To.Type = obj.TYPE_NONE

		p := s.Prog(psess.gc, ld)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = arg0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = out

		p1 := s.Prog(psess.gc, cmp)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = out
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = out

		p2 := s.Prog(psess.gc, ppc64.ABNE)
		p2.To.Type = obj.TYPE_BRANCH

		pisync := s.Prog(psess.gc, ppc64.AISYNC)
		pisync.To.Type = obj.TYPE_NONE
		psess.gc.
			Patch(p2, pisync)

	case ssa.OpPPC64LoweredAtomicStore32,
		ssa.OpPPC64LoweredAtomicStore64:

		st := ppc64.AMOVD
		if v.Op == ssa.OpPPC64LoweredAtomicStore32 {
			st = ppc64.AMOVW
		}
		arg0 := v.Args[0].Reg(psess.ssa)
		arg1 := v.Args[1].Reg(psess.ssa)

		psync := s.Prog(psess.gc, ppc64.ASYNC)
		psync.To.Type = obj.TYPE_NONE

		p := s.Prog(psess.gc, st)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = arg0
		p.From.Type = obj.TYPE_REG
		p.From.Reg = arg1

	case ssa.OpPPC64LoweredAtomicCas64,
		ssa.OpPPC64LoweredAtomicCas32:

		ld := ppc64.ALDAR
		st := ppc64.ASTDCCC
		cmp := ppc64.ACMP
		if v.Op == ssa.OpPPC64LoweredAtomicCas32 {
			ld = ppc64.ALWAR
			st = ppc64.ASTWCCC
			cmp = ppc64.ACMPW
		}
		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)
		r2 := v.Args[2].Reg(psess.ssa)
		out := v.Reg0(psess.ssa)

		plwsync1 := s.Prog(psess.gc, ppc64.ALWSYNC)
		plwsync1.To.Type = obj.TYPE_NONE

		p := s.Prog(psess.gc, ld)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REGTMP

		p1 := s.Prog(psess.gc, cmp)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.To.Reg = ppc64.REGTMP
		p1.To.Type = obj.TYPE_REG

		p2 := s.Prog(psess.gc, ppc64.ABNE)
		p2.To.Type = obj.TYPE_BRANCH

		p3 := s.Prog(psess.gc, st)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = r2
		p3.To.Type = obj.TYPE_MEM
		p3.To.Reg = r0

		p4 := s.Prog(psess.gc, ppc64.ABNE)
		p4.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p4, p)

		plwsync2 := s.Prog(psess.gc, ppc64.ALWSYNC)
		plwsync2.To.Type = obj.TYPE_NONE

		p5 := s.Prog(psess.gc, ppc64.AMOVD)
		p5.From.Type = obj.TYPE_CONST
		p5.From.Offset = 1
		p5.To.Type = obj.TYPE_REG
		p5.To.Reg = out

		p6 := s.Prog(psess.gc, obj.AJMP)
		p6.To.Type = obj.TYPE_BRANCH

		p7 := s.Prog(psess.gc, ppc64.AMOVD)
		p7.From.Type = obj.TYPE_CONST
		p7.From.Offset = 0
		p7.To.Type = obj.TYPE_REG
		p7.To.Reg = out
		psess.gc.
			Patch(p2, p7)

		p8 := s.Prog(psess.gc, obj.ANOP)
		psess.gc.
			Patch(p6, p8)

	case ssa.OpPPC64LoweredGetClosurePtr:
		psess.gc.
			CheckLoweredGetClosurePtr(v)

	case ssa.OpPPC64LoweredGetCallerSP:

		p := s.Prog(psess.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_ADDR
		p.From.Offset = -psess.gc.Ctxt.FixedFrameSize()
		p.From.Name = obj.NAME_PARAM
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.OpPPC64LoweredGetCallerPC:
		p := s.Prog(psess.gc, obj.AGETCALLERPC)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.OpPPC64LoweredRound32F, ssa.OpPPC64LoweredRound64F:

	case ssa.OpLoadReg:
		loadOp := psess.loadByType(v.Type)
		p := s.Prog(psess.gc, loadOp)
		psess.gc.
			AddrAuto(&p.From, v.Args[0])
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.OpStoreReg:
		storeOp := psess.storeByType(v.Type)
		p := s.Prog(psess.gc, storeOp)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddrAuto(&p.To, v)

	case ssa.OpPPC64DIVD:

		r := v.Reg(psess.ssa)
		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)

		p := s.Prog(psess.gc, ppc64.ACMP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = -1

		pbahead := s.Prog(psess.gc, ppc64.ABEQ)
		pbahead.To.Type = obj.TYPE_BRANCH

		p = s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

		pbover := s.Prog(psess.gc, obj.AJMP)
		pbover.To.Type = obj.TYPE_BRANCH

		p = s.Prog(psess.gc, ppc64.ANEG)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r0
		psess.gc.
			Patch(pbahead, p)

		p = s.Prog(psess.gc, obj.ANOP)
		psess.gc.
			Patch(pbover, p)

	case ssa.OpPPC64DIVW:

		r := v.Reg(psess.ssa)
		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)

		p := s.Prog(psess.gc, ppc64.ACMPW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = -1

		pbahead := s.Prog(psess.gc, ppc64.ABEQ)
		pbahead.To.Type = obj.TYPE_BRANCH

		p = s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

		pbover := s.Prog(psess.gc, obj.AJMP)
		pbover.To.Type = obj.TYPE_BRANCH

		p = s.Prog(psess.gc, ppc64.ANEG)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r0
		psess.gc.
			Patch(pbahead, p)

		p = s.Prog(psess.gc, obj.ANOP)
		psess.gc.
			Patch(pbover, p)

	case ssa.OpPPC64ADD, ssa.OpPPC64FADD, ssa.OpPPC64FADDS, ssa.OpPPC64SUB, ssa.OpPPC64FSUB, ssa.OpPPC64FSUBS,
		ssa.OpPPC64MULLD, ssa.OpPPC64MULLW, ssa.OpPPC64DIVDU, ssa.OpPPC64DIVWU,
		ssa.OpPPC64SRAD, ssa.OpPPC64SRAW, ssa.OpPPC64SRD, ssa.OpPPC64SRW, ssa.OpPPC64SLD, ssa.OpPPC64SLW,
		ssa.OpPPC64ROTL, ssa.OpPPC64ROTLW,
		ssa.OpPPC64MULHD, ssa.OpPPC64MULHW, ssa.OpPPC64MULHDU, ssa.OpPPC64MULHWU,
		ssa.OpPPC64FMUL, ssa.OpPPC64FMULS, ssa.OpPPC64FDIV, ssa.OpPPC64FDIVS, ssa.OpPPC64FCPSGN,
		ssa.OpPPC64AND, ssa.OpPPC64OR, ssa.OpPPC64ANDN, ssa.OpPPC64ORN, ssa.OpPPC64NOR, ssa.OpPPC64XOR, ssa.OpPPC64EQV:
		r := v.Reg(psess.ssa)
		r1 := v.Args[0].Reg(psess.ssa)
		r2 := v.Args[1].Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r2
		p.Reg = r1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

	case ssa.OpPPC64ROTLconst, ssa.OpPPC64ROTLWconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.OpPPC64FMADD, ssa.OpPPC64FMADDS, ssa.OpPPC64FMSUB, ssa.OpPPC64FMSUBS:
		r := v.Reg(psess.ssa)
		r1 := v.Args[0].Reg(psess.ssa)
		r2 := v.Args[1].Reg(psess.ssa)
		r3 := v.Args[2].Reg(psess.ssa)

		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.Reg = r3
		p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: r2})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

	case ssa.OpPPC64MaskIfNotCarry:
		r := v.Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = ppc64.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

	case ssa.OpPPC64ADDconstForCarry:
		r1 := v.Args[0].Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.Reg = r1
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REGTMP

	case ssa.OpPPC64NEG, ssa.OpPPC64FNEG, ssa.OpPPC64FSQRT, ssa.OpPPC64FSQRTS, ssa.OpPPC64FFLOOR, ssa.OpPPC64FTRUNC, ssa.OpPPC64FCEIL, ssa.OpPPC64FCTIDZ, ssa.OpPPC64FCTIWZ, ssa.OpPPC64FCFID, ssa.OpPPC64FCFIDS, ssa.OpPPC64FRSP, ssa.OpPPC64CNTLZD, ssa.OpPPC64CNTLZW, ssa.OpPPC64POPCNTD, ssa.OpPPC64POPCNTW, ssa.OpPPC64POPCNTB, ssa.OpPPC64MFVSRD, ssa.OpPPC64MTVSRD, ssa.OpPPC64FABS, ssa.OpPPC64FNABS, ssa.OpPPC64FROUND:
		r := v.Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)

	case ssa.OpPPC64ADDconst, ssa.OpPPC64ANDconst, ssa.OpPPC64ORconst, ssa.OpPPC64XORconst,
		ssa.OpPPC64SRADconst, ssa.OpPPC64SRAWconst, ssa.OpPPC64SRDconst, ssa.OpPPC64SRWconst, ssa.OpPPC64SLDconst, ssa.OpPPC64SLWconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.Reg = v.Args[0].Reg(psess.ssa)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.OpPPC64ANDCCconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.Reg = v.Args[0].Reg(psess.ssa)

		if v.Aux != nil {
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = psess.gc.AuxOffset(v)
		} else {
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = v.AuxInt
		}

		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REGTMP

	case ssa.OpPPC64MOVDaddr:
		switch v.Aux.(type) {
		default:
			v.Fatalf("aux in MOVDaddr is of unknown type %T", v.Aux)
		case nil:

			if v.AuxInt != 0 || v.Args[0].Reg(psess.ssa) != v.Reg(psess.ssa) {
				p := s.Prog(psess.gc, ppc64.AMOVD)
				p.From.Type = obj.TYPE_ADDR
				p.From.Reg = v.Args[0].Reg(psess.ssa)
				p.From.Offset = v.AuxInt
				p.To.Type = obj.TYPE_REG
				p.To.Reg = v.Reg(psess.ssa)
			}

		case *obj.LSym, *gc.Node:
			p := s.Prog(psess.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_ADDR
			p.From.Reg = v.Args[0].Reg(psess.ssa)
			p.To.Type = obj.TYPE_REG
			p.To.Reg = v.Reg(psess.ssa)
			psess.gc.
				AddAux(&p.From, v)

		}

	case ssa.OpPPC64MOVDconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.OpPPC64FMOVDconst, ssa.OpPPC64FMOVSconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_FCONST
		p.From.Val = math.Float64frombits(uint64(v.AuxInt))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.OpPPC64FCMPU, ssa.OpPPC64CMP, ssa.OpPPC64CMPW, ssa.OpPPC64CMPU, ssa.OpPPC64CMPWU:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Args[1].Reg(psess.ssa)

	case ssa.OpPPC64CMPconst, ssa.OpPPC64CMPUconst, ssa.OpPPC64CMPWconst, ssa.OpPPC64CMPWUconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = v.AuxInt

	case ssa.OpPPC64MOVBreg, ssa.OpPPC64MOVBZreg, ssa.OpPPC64MOVHreg, ssa.OpPPC64MOVHZreg, ssa.OpPPC64MOVWreg, ssa.OpPPC64MOVWZreg:

		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Reg = v.Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG

	case ssa.OpPPC64MOVDload:

		gostring := false
		switch n := v.Aux.(type) {
		case *obj.LSym:
			gostring = strings.HasPrefix(n.Name, "go.string.")
		}
		if gostring {

			p := s.Prog(psess.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_ADDR
			p.From.Reg = v.Args[0].Reg(psess.ssa)
			psess.gc.
				AddAux(&p.From, v)
			p.To.Type = obj.TYPE_REG
			p.To.Reg = v.Reg(psess.ssa)

			p = s.Prog(psess.gc, v.Op.Asm(psess.ssa))
			p.From.Type = obj.TYPE_MEM
			p.From.Reg = v.Reg(psess.ssa)
			p.To.Type = obj.TYPE_REG
			p.To.Reg = v.Reg(psess.ssa)
			break
		}

		fallthrough

	case ssa.OpPPC64MOVWload, ssa.OpPPC64MOVHload, ssa.OpPPC64MOVWZload, ssa.OpPPC64MOVBZload, ssa.OpPPC64MOVHZload:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.OpPPC64MOVDBRload, ssa.OpPPC64MOVWBRload, ssa.OpPPC64MOVHBRload:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.OpPPC64MOVDBRstore, ssa.OpPPC64MOVWBRstore, ssa.OpPPC64MOVHBRstore:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)

	case ssa.OpPPC64FMOVDload, ssa.OpPPC64FMOVSload:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.OpPPC64MOVDstorezero, ssa.OpPPC64MOVWstorezero, ssa.OpPPC64MOVHstorezero, ssa.OpPPC64MOVBstorezero:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = ppc64.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)

	case ssa.OpPPC64MOVDstore, ssa.OpPPC64MOVWstore, ssa.OpPPC64MOVHstore, ssa.OpPPC64MOVBstore:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpPPC64FMOVDstore, ssa.OpPPC64FMOVSstore:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)

	case ssa.OpPPC64Equal,
		ssa.OpPPC64NotEqual,
		ssa.OpPPC64LessThan,
		ssa.OpPPC64FLessThan,
		ssa.OpPPC64LessEqual,
		ssa.OpPPC64GreaterThan,
		ssa.OpPPC64FGreaterThan,
		ssa.OpPPC64GreaterEqual:

		p := s.Prog(psess.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = psess.iselRegs[1]
		iop := psess.iselOps[v.Op]
		psess.
			ssaGenISEL(s, v, iop.cond, psess.iselRegs[iop.valueIfCond], psess.iselRegs[1-iop.valueIfCond])

	case ssa.OpPPC64FLessEqual,
		ssa.OpPPC64FGreaterEqual:

		p := s.Prog(psess.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = psess.iselRegs[1]
		iop := psess.iselOps[v.Op]
		psess.
			ssaGenISEL(s, v, iop.cond, psess.iselRegs[iop.valueIfCond], psess.iselRegs[1-iop.valueIfCond])
		psess.
			ssaGenISEL(s, v, ppc64.C_COND_EQ, psess.iselRegs[1], v.Reg(psess.ssa))

	case ssa.OpPPC64LoweredZero:

		ctr := v.AuxInt / 32

		rem := v.AuxInt % 32

		if ctr > 1 {

			p := s.Prog(psess.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = ctr
			p.To.Type = obj.TYPE_REG
			p.To.Reg = ppc64.REGTMP

			p = s.Prog(psess.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = ppc64.REGTMP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = ppc64.REG_CTR

			// generate 4 MOVDs
			// when this is a loop then the top must be saved
			var top *obj.Prog
			for offset := int64(0); offset < 32; offset += 8 {

				p := s.Prog(psess.gc, ppc64.AMOVD)
				p.From.Type = obj.TYPE_REG
				p.From.Reg = ppc64.REG_R0
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = v.Args[0].Reg(psess.ssa)
				p.To.Offset = offset

				if top == nil {
					top = p
				}
			}

			p = s.Prog(psess.gc, ppc64.AADD)
			p.Reg = v.Args[0].Reg(psess.ssa)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = 32
			p.To.Type = obj.TYPE_REG
			p.To.Reg = v.Args[0].Reg(psess.ssa)

			p = s.Prog(psess.gc, ppc64.ABC)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = ppc64.BO_BCTR
			p.Reg = ppc64.REG_R0
			p.To.Type = obj.TYPE_BRANCH
			psess.gc.
				Patch(p, top)
		}

		if ctr == 1 {
			rem += 32
		}

		offset := int64(0)

		for rem > 0 {
			op, size := ppc64.AMOVB, int64(1)
			switch {
			case rem >= 8:
				op, size = ppc64.AMOVD, 8
			case rem >= 4:
				op, size = ppc64.AMOVW, 4
			case rem >= 2:
				op, size = ppc64.AMOVH, 2
			}
			p := s.Prog(psess.gc, op)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = ppc64.REG_R0
			p.To.Type = obj.TYPE_MEM
			p.To.Reg = v.Args[0].Reg(psess.ssa)
			p.To.Offset = offset
			rem -= size
			offset += size
		}

	case ssa.OpPPC64LoweredMove:

		ctr := v.AuxInt / 32

		rem := v.AuxInt % 32

		dst_reg := v.Args[0].Reg(psess.ssa)
		src_reg := v.Args[1].Reg(psess.ssa)

		useregs := []int16{ppc64.REG_R7, ppc64.REG_R8, ppc64.REG_R9, ppc64.REG_R10}
		offset := int64(0)

		// top of the loop
		var top *obj.Prog

		if ctr > 1 {

			p := s.Prog(psess.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = ctr
			p.To.Type = obj.TYPE_REG
			p.To.Reg = ppc64.REGTMP

			p = s.Prog(psess.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = ppc64.REGTMP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = ppc64.REG_CTR

			for _, rg := range useregs {
				p := s.Prog(psess.gc, ppc64.AMOVD)
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = src_reg
				p.From.Offset = offset
				p.To.Type = obj.TYPE_REG
				p.To.Reg = rg
				if top == nil {
					top = p
				}
				offset += 8
			}

			p = s.Prog(psess.gc, ppc64.AADD)
			p.Reg = src_reg
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = 32
			p.To.Type = obj.TYPE_REG
			p.To.Reg = src_reg

			offset = int64(0)
			for _, rg := range useregs {
				p := s.Prog(psess.gc, ppc64.AMOVD)
				p.From.Type = obj.TYPE_REG
				p.From.Reg = rg
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = dst_reg
				p.To.Offset = offset
				offset += 8
			}

			p = s.Prog(psess.gc, ppc64.AADD)
			p.Reg = dst_reg
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = 32
			p.To.Type = obj.TYPE_REG
			p.To.Reg = dst_reg

			p = s.Prog(psess.gc, ppc64.ABC)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = ppc64.BO_BCTR
			p.Reg = ppc64.REG_R0
			p.To.Type = obj.TYPE_BRANCH
			psess.gc.
				Patch(p, top)

			offset = int64(0)
		}

		if ctr == 1 {
			rem += 32
		}

		for rem > 0 {
			op, size := ppc64.AMOVB, int64(1)
			switch {
			case rem >= 8:
				op, size = ppc64.AMOVD, 8
			case rem >= 4:
				op, size = ppc64.AMOVW, 4
			case rem >= 2:
				op, size = ppc64.AMOVH, 2
			}

			p := s.Prog(psess.gc, op)
			p.To.Type = obj.TYPE_REG
			p.To.Reg = ppc64.REG_R7
			p.From.Type = obj.TYPE_MEM
			p.From.Reg = src_reg
			p.From.Offset = offset

			p = s.Prog(psess.gc, op)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = ppc64.REG_R7
			p.To.Type = obj.TYPE_MEM
			p.To.Reg = dst_reg
			p.To.Offset = offset
			rem -= size
			offset += size
		}

	case ssa.OpPPC64CALLstatic:
		s.Call(psess.gc, v)

	case ssa.OpPPC64CALLclosure, ssa.OpPPC64CALLinter:
		p := s.Prog(psess.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REG_CTR

		if v.Args[0].Reg(psess.ssa) != ppc64.REG_R12 {
			v.Fatalf("Function address for %v should be in R12 %d but is in %d", v.LongString(psess.ssa), ppc64.REG_R12, p.From.Reg)
		}

		pp := s.Call(psess.gc, v)
		pp.To.Reg = ppc64.REG_CTR

		if psess.gc.Ctxt.Flag_shared {

			q := s.Prog(psess.gc, ppc64.AMOVD)
			q.From.Type = obj.TYPE_MEM
			q.From.Offset = 24
			q.From.Reg = ppc64.REGSP
			q.To.Type = obj.TYPE_REG
			q.To.Reg = ppc64.REG_R2
		}

	case ssa.OpPPC64LoweredWB:
		p := s.Prog(psess.gc, obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = v.Aux.(*obj.LSym)

	case ssa.OpPPC64LoweredNilCheck:

		p := s.Prog(psess.gc, ppc64.AMOVBZ)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REGTMP
		if psess.gc.Debug_checknil != 0 && v.Pos.Line() > 1 {
			psess.gc.
				Warnl(v.Pos, "generated nil check")
		}

	case ssa.OpPPC64InvertFlags:
		v.Fatalf("InvertFlags should never make it to codegen %v", v.LongString(psess.ssa))
	case ssa.OpPPC64FlagEQ, ssa.OpPPC64FlagLT, ssa.OpPPC64FlagGT:
		v.Fatalf("Flag* ops should never make it to codegen %v", v.LongString(psess.ssa))
	case ssa.OpClobber:

	default:
		v.Fatalf("genValue not implemented: %s", v.LongString(psess.ssa))
	}
}

func (psess *PackageSession) ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	switch b.Kind {
	case ssa.BlockDefer:

		p := s.Prog(psess.gc, ppc64.ACMP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = ppc64.REG_R3
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REG_R0

		p = s.Prog(psess.gc, ppc64.ABNE)
		p.To.Type = obj.TYPE_BRANCH
		s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[1].Block()})
		if b.Succs[0].Block() != next {
			p := s.Prog(psess.gc, obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}

	case ssa.BlockPlain:
		if b.Succs[0].Block() != next {
			p := s.Prog(psess.gc, obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}
	case ssa.BlockExit:
		s.Prog(psess.gc, obj.AUNDEF)
	case ssa.BlockRet:
		s.Prog(psess.gc, obj.ARET)
	case ssa.BlockRetJmp:
		p := s.Prog(psess.gc, obj.AJMP)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = b.Aux.(*obj.LSym)

	case ssa.BlockPPC64EQ, ssa.BlockPPC64NE,
		ssa.BlockPPC64LT, ssa.BlockPPC64GE,
		ssa.BlockPPC64LE, ssa.BlockPPC64GT,
		ssa.BlockPPC64FLT, ssa.BlockPPC64FGE,
		ssa.BlockPPC64FLE, ssa.BlockPPC64FGT:
		jmp := psess.blockJump[b.Kind]
		switch next {
		case b.Succs[0].Block():
			s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
			if jmp.invasmun {

				s.Br(psess.gc, ppc64.ABVS, b.Succs[1].Block())
			}
		case b.Succs[1].Block():
			s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
			if jmp.asmeq {
				s.Br(psess.gc, ppc64.ABEQ, b.Succs[0].Block())
			}
		default:
			if b.Likely != ssa.BranchUnlikely {
				s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
				if jmp.asmeq {
					s.Br(psess.gc, ppc64.ABEQ, b.Succs[0].Block())
				}
				s.Br(psess.gc, obj.AJMP, b.Succs[1].Block())
			} else {
				s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
				if jmp.invasmun {

					s.Br(psess.gc, ppc64.ABVS, b.Succs[1].Block())
				}
				s.Br(psess.gc, obj.AJMP, b.Succs[0].Block())
			}
		}
	default:
		b.Fatalf("branch not implemented: %s. Control: %s", b.LongString(psess.ssa), b.Control.LongString(psess.ssa))
	}
}

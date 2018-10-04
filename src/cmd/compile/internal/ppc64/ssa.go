// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

// markMoves marks any MOVXconst ops that need to avoid clobbering flags.
func ssaMarkMoves(s *gc.SSAGenState, b *ssa.Block) {
	//	flive := b.FlagsLiveAtEnd
	//	if b.Control != nil && b.Control.Type.IsFlags() {
	//		flive = true
	//	}
	//	for i := len(b.Values) - 1; i >= 0; i-- {
	//		v := b.Values[i]
	//		if flive && (v.Op == v.Op == ssa.OpPPC64MOVDconst) {
	//			// The "mark" is any non-nil Aux value.
	//			v.Aux = v
	//		}
	//		if v.Type.IsFlags() {
	//			flive = false
	//		}
	//		for _, a := range v.Args {
	//			if a.Type.IsFlags() {
	//				flive = true
	//			}
	//		}
	//	}
}

// loadByType returns the load instruction of the given type.
func (pstate *PackageState) loadByType(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(pstate.types) {
		case 4:
			return ppc64.AFMOVS
		case 8:
			return ppc64.AFMOVD
		}
	} else {
		switch t.Size(pstate.types) {
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
func (pstate *PackageState) storeByType(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(pstate.types) {
		case 4:
			return ppc64.AFMOVS
		case 8:
			return ppc64.AFMOVD
		}
	} else {
		switch t.Size(pstate.types) {
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

func (pstate *PackageState) ssaGenISEL(s *gc.SSAGenState, v *ssa.Value, cr int64, r1, r2 int16) {
	r := v.Reg(pstate.ssa)
	p := s.Prog(pstate.gc, ppc64.AISEL)
	p.To.Type = obj.TYPE_REG
	p.To.Reg = r
	p.Reg = r1
	p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: r2})
	p.From.Type = obj.TYPE_CONST
	p.From.Offset = cr
}

func (pstate *PackageState) ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpCopy:
		t := v.Type
		if t.IsMemory(pstate.types) {
			return
		}
		x := v.Args[0].Reg(pstate.ssa)
		y := v.Reg(pstate.ssa)
		if x != y {
			rt := obj.TYPE_REG
			op := ppc64.AMOVD

			if t.IsFloat() {
				op = ppc64.AFMOVD
			}
			p := s.Prog(pstate.gc, op)
			p.From.Type = rt
			p.From.Reg = x
			p.To.Type = rt
			p.To.Reg = y
		}

	case ssa.OpPPC64LoweredAtomicAnd8,
		ssa.OpPPC64LoweredAtomicOr8:
		// LWSYNC
		// LBAR		(Rarg0), Rtmp
		// AND/OR	Rarg1, Rtmp
		// STBCCC	Rtmp, (Rarg0)
		// BNE		-3(PC)
		r0 := v.Args[0].Reg(pstate.ssa)
		r1 := v.Args[1].Reg(pstate.ssa)
		// LWSYNC - Assuming shared data not write-through-required nor
		// caching-inhibited. See Appendix B.2.2.2 in the ISA 2.07b.
		plwsync := s.Prog(pstate.gc, ppc64.ALWSYNC)
		plwsync.To.Type = obj.TYPE_NONE
		p := s.Prog(pstate.gc, ppc64.ALBAR)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REGTMP
		p1 := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = ppc64.REGTMP
		p2 := s.Prog(pstate.gc, ppc64.ASTBCCC)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = ppc64.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = r0
		p2.RegTo2 = ppc64.REGTMP
		p3 := s.Prog(pstate.gc, ppc64.ABNE)
		p3.To.Type = obj.TYPE_BRANCH
		pstate.gc.Patch(p3, p)

	case ssa.OpPPC64LoweredAtomicAdd32,
		ssa.OpPPC64LoweredAtomicAdd64:
		// LWSYNC
		// LDAR/LWAR    (Rarg0), Rout
		// ADD		Rarg1, Rout
		// STDCCC/STWCCC Rout, (Rarg0)
		// BNE         -3(PC)
		// MOVW		Rout,Rout (if Add32)
		ld := ppc64.ALDAR
		st := ppc64.ASTDCCC
		if v.Op == ssa.OpPPC64LoweredAtomicAdd32 {
			ld = ppc64.ALWAR
			st = ppc64.ASTWCCC
		}
		r0 := v.Args[0].Reg(pstate.ssa)
		r1 := v.Args[1].Reg(pstate.ssa)
		out := v.Reg0(pstate.ssa)
		// LWSYNC - Assuming shared data not write-through-required nor
		// caching-inhibited. See Appendix B.2.2.2 in the ISA 2.07b.
		plwsync := s.Prog(pstate.gc, ppc64.ALWSYNC)
		plwsync.To.Type = obj.TYPE_NONE
		// LDAR or LWAR
		p := s.Prog(pstate.gc, ld)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = out
		// ADD reg1,out
		p1 := s.Prog(pstate.gc, ppc64.AADD)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.To.Reg = out
		p1.To.Type = obj.TYPE_REG
		// STDCCC or STWCCC
		p3 := s.Prog(pstate.gc, st)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = out
		p3.To.Type = obj.TYPE_MEM
		p3.To.Reg = r0
		// BNE retry
		p4 := s.Prog(pstate.gc, ppc64.ABNE)
		p4.To.Type = obj.TYPE_BRANCH
		pstate.gc.Patch(p4, p)

		// Ensure a 32 bit result
		if v.Op == ssa.OpPPC64LoweredAtomicAdd32 {
			p5 := s.Prog(pstate.gc, ppc64.AMOVWZ)
			p5.To.Type = obj.TYPE_REG
			p5.To.Reg = out
			p5.From.Type = obj.TYPE_REG
			p5.From.Reg = out
		}

	case ssa.OpPPC64LoweredAtomicExchange32,
		ssa.OpPPC64LoweredAtomicExchange64:
		// LWSYNC
		// LDAR/LWAR    (Rarg0), Rout
		// STDCCC/STWCCC Rout, (Rarg0)
		// BNE         -2(PC)
		// ISYNC
		ld := ppc64.ALDAR
		st := ppc64.ASTDCCC
		if v.Op == ssa.OpPPC64LoweredAtomicExchange32 {
			ld = ppc64.ALWAR
			st = ppc64.ASTWCCC
		}
		r0 := v.Args[0].Reg(pstate.ssa)
		r1 := v.Args[1].Reg(pstate.ssa)
		out := v.Reg0(pstate.ssa)
		// LWSYNC - Assuming shared data not write-through-required nor
		// caching-inhibited. See Appendix B.2.2.2 in the ISA 2.07b.
		plwsync := s.Prog(pstate.gc, ppc64.ALWSYNC)
		plwsync.To.Type = obj.TYPE_NONE
		// LDAR or LWAR
		p := s.Prog(pstate.gc, ld)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = out
		// STDCCC or STWCCC
		p1 := s.Prog(pstate.gc, st)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.To.Type = obj.TYPE_MEM
		p1.To.Reg = r0
		// BNE retry
		p2 := s.Prog(pstate.gc, ppc64.ABNE)
		p2.To.Type = obj.TYPE_BRANCH
		pstate.gc.Patch(p2, p)
		// ISYNC
		pisync := s.Prog(pstate.gc, ppc64.AISYNC)
		pisync.To.Type = obj.TYPE_NONE

	case ssa.OpPPC64LoweredAtomicLoad32,
		ssa.OpPPC64LoweredAtomicLoad64,
		ssa.OpPPC64LoweredAtomicLoadPtr:
		// SYNC
		// MOVD/MOVW (Rarg0), Rout
		// CMP Rout,Rout
		// BNE 1(PC)
		// ISYNC
		ld := ppc64.AMOVD
		cmp := ppc64.ACMP
		if v.Op == ssa.OpPPC64LoweredAtomicLoad32 {
			ld = ppc64.AMOVW
			cmp = ppc64.ACMPW
		}
		arg0 := v.Args[0].Reg(pstate.ssa)
		out := v.Reg0(pstate.ssa)
		// SYNC
		psync := s.Prog(pstate.gc, ppc64.ASYNC)
		psync.To.Type = obj.TYPE_NONE
		// Load
		p := s.Prog(pstate.gc, ld)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = arg0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = out
		// CMP
		p1 := s.Prog(pstate.gc, cmp)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = out
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = out
		// BNE
		p2 := s.Prog(pstate.gc, ppc64.ABNE)
		p2.To.Type = obj.TYPE_BRANCH
		// ISYNC
		pisync := s.Prog(pstate.gc, ppc64.AISYNC)
		pisync.To.Type = obj.TYPE_NONE
		pstate.gc.Patch(p2, pisync)

	case ssa.OpPPC64LoweredAtomicStore32,
		ssa.OpPPC64LoweredAtomicStore64:
		// SYNC
		// MOVD/MOVW arg1,(arg0)
		st := ppc64.AMOVD
		if v.Op == ssa.OpPPC64LoweredAtomicStore32 {
			st = ppc64.AMOVW
		}
		arg0 := v.Args[0].Reg(pstate.ssa)
		arg1 := v.Args[1].Reg(pstate.ssa)
		// SYNC
		psync := s.Prog(pstate.gc, ppc64.ASYNC)
		psync.To.Type = obj.TYPE_NONE
		// Store
		p := s.Prog(pstate.gc, st)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = arg0
		p.From.Type = obj.TYPE_REG
		p.From.Reg = arg1

	case ssa.OpPPC64LoweredAtomicCas64,
		ssa.OpPPC64LoweredAtomicCas32:
		// LWSYNC
		// loop:
		// LDAR        (Rarg0), Rtmp
		// CMP         Rarg1, Rtmp
		// BNE         fail
		// STDCCC      Rarg2, (Rarg0)
		// BNE         loop
		// LWSYNC
		// MOVD        $1, Rout
		// BR          end
		// fail:
		// MOVD        $0, Rout
		// end:
		ld := ppc64.ALDAR
		st := ppc64.ASTDCCC
		cmp := ppc64.ACMP
		if v.Op == ssa.OpPPC64LoweredAtomicCas32 {
			ld = ppc64.ALWAR
			st = ppc64.ASTWCCC
			cmp = ppc64.ACMPW
		}
		r0 := v.Args[0].Reg(pstate.ssa)
		r1 := v.Args[1].Reg(pstate.ssa)
		r2 := v.Args[2].Reg(pstate.ssa)
		out := v.Reg0(pstate.ssa)
		// LWSYNC - Assuming shared data not write-through-required nor
		// caching-inhibited. See Appendix B.2.2.2 in the ISA 2.07b.
		plwsync1 := s.Prog(pstate.gc, ppc64.ALWSYNC)
		plwsync1.To.Type = obj.TYPE_NONE
		// LDAR or LWAR
		p := s.Prog(pstate.gc, ld)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REGTMP
		// CMP reg1,reg2
		p1 := s.Prog(pstate.gc, cmp)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.To.Reg = ppc64.REGTMP
		p1.To.Type = obj.TYPE_REG
		// BNE cas_fail
		p2 := s.Prog(pstate.gc, ppc64.ABNE)
		p2.To.Type = obj.TYPE_BRANCH
		// STDCCC or STWCCC
		p3 := s.Prog(pstate.gc, st)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = r2
		p3.To.Type = obj.TYPE_MEM
		p3.To.Reg = r0
		// BNE retry
		p4 := s.Prog(pstate.gc, ppc64.ABNE)
		p4.To.Type = obj.TYPE_BRANCH
		pstate.gc.Patch(p4, p)
		// LWSYNC - Assuming shared data not write-through-required nor
		// caching-inhibited. See Appendix B.2.1.1 in the ISA 2.07b.
		plwsync2 := s.Prog(pstate.gc, ppc64.ALWSYNC)
		plwsync2.To.Type = obj.TYPE_NONE
		// return true
		p5 := s.Prog(pstate.gc, ppc64.AMOVD)
		p5.From.Type = obj.TYPE_CONST
		p5.From.Offset = 1
		p5.To.Type = obj.TYPE_REG
		p5.To.Reg = out
		// BR done
		p6 := s.Prog(pstate.gc, obj.AJMP)
		p6.To.Type = obj.TYPE_BRANCH
		// return false
		p7 := s.Prog(pstate.gc, ppc64.AMOVD)
		p7.From.Type = obj.TYPE_CONST
		p7.From.Offset = 0
		p7.To.Type = obj.TYPE_REG
		p7.To.Reg = out
		pstate.gc.Patch(p2, p7)
		// done (label)
		p8 := s.Prog(pstate.gc, obj.ANOP)
		pstate.gc.Patch(p6, p8)

	case ssa.OpPPC64LoweredGetClosurePtr:
		// Closure pointer is R11 (already)
		pstate.gc.CheckLoweredGetClosurePtr(v)

	case ssa.OpPPC64LoweredGetCallerSP:
		// caller's SP is FixedFrameSize below the address of the first arg
		p := s.Prog(pstate.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_ADDR
		p.From.Offset = -pstate.gc.Ctxt.FixedFrameSize()
		p.From.Name = obj.NAME_PARAM
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)

	case ssa.OpPPC64LoweredGetCallerPC:
		p := s.Prog(pstate.gc, obj.AGETCALLERPC)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)

	case ssa.OpPPC64LoweredRound32F, ssa.OpPPC64LoweredRound64F:
	// input is already rounded

	case ssa.OpLoadReg:
		loadOp := pstate.loadByType(v.Type)
		p := s.Prog(pstate.gc, loadOp)
		pstate.gc.AddrAuto(&p.From, v.Args[0])
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)

	case ssa.OpStoreReg:
		storeOp := pstate.storeByType(v.Type)
		p := s.Prog(pstate.gc, storeOp)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddrAuto(&p.To, v)

	case ssa.OpPPC64DIVD:
		// For now,
		//
		// cmp arg1, -1
		// be  ahead
		// v = arg0 / arg1
		// b over
		// ahead: v = - arg0
		// over: nop
		r := v.Reg(pstate.ssa)
		r0 := v.Args[0].Reg(pstate.ssa)
		r1 := v.Args[1].Reg(pstate.ssa)

		p := s.Prog(pstate.gc, ppc64.ACMP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = -1

		pbahead := s.Prog(pstate.gc, ppc64.ABEQ)
		pbahead.To.Type = obj.TYPE_BRANCH

		p = s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

		pbover := s.Prog(pstate.gc, obj.AJMP)
		pbover.To.Type = obj.TYPE_BRANCH

		p = s.Prog(pstate.gc, ppc64.ANEG)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r0
		pstate.gc.Patch(pbahead, p)

		p = s.Prog(pstate.gc, obj.ANOP)
		pstate.gc.Patch(pbover, p)

	case ssa.OpPPC64DIVW:
		// word-width version of above
		r := v.Reg(pstate.ssa)
		r0 := v.Args[0].Reg(pstate.ssa)
		r1 := v.Args[1].Reg(pstate.ssa)

		p := s.Prog(pstate.gc, ppc64.ACMPW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = -1

		pbahead := s.Prog(pstate.gc, ppc64.ABEQ)
		pbahead.To.Type = obj.TYPE_BRANCH

		p = s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

		pbover := s.Prog(pstate.gc, obj.AJMP)
		pbover.To.Type = obj.TYPE_BRANCH

		p = s.Prog(pstate.gc, ppc64.ANEG)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r0
		pstate.gc.Patch(pbahead, p)

		p = s.Prog(pstate.gc, obj.ANOP)
		pstate.gc.Patch(pbover, p)

	case ssa.OpPPC64ADD, ssa.OpPPC64FADD, ssa.OpPPC64FADDS, ssa.OpPPC64SUB, ssa.OpPPC64FSUB, ssa.OpPPC64FSUBS,
		ssa.OpPPC64MULLD, ssa.OpPPC64MULLW, ssa.OpPPC64DIVDU, ssa.OpPPC64DIVWU,
		ssa.OpPPC64SRAD, ssa.OpPPC64SRAW, ssa.OpPPC64SRD, ssa.OpPPC64SRW, ssa.OpPPC64SLD, ssa.OpPPC64SLW,
		ssa.OpPPC64ROTL, ssa.OpPPC64ROTLW,
		ssa.OpPPC64MULHD, ssa.OpPPC64MULHW, ssa.OpPPC64MULHDU, ssa.OpPPC64MULHWU,
		ssa.OpPPC64FMUL, ssa.OpPPC64FMULS, ssa.OpPPC64FDIV, ssa.OpPPC64FDIVS, ssa.OpPPC64FCPSGN,
		ssa.OpPPC64AND, ssa.OpPPC64OR, ssa.OpPPC64ANDN, ssa.OpPPC64ORN, ssa.OpPPC64NOR, ssa.OpPPC64XOR, ssa.OpPPC64EQV:
		r := v.Reg(pstate.ssa)
		r1 := v.Args[0].Reg(pstate.ssa)
		r2 := v.Args[1].Reg(pstate.ssa)
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r2
		p.Reg = r1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

	case ssa.OpPPC64ROTLconst, ssa.OpPPC64ROTLWconst:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)

	case ssa.OpPPC64FMADD, ssa.OpPPC64FMADDS, ssa.OpPPC64FMSUB, ssa.OpPPC64FMSUBS:
		r := v.Reg(pstate.ssa)
		r1 := v.Args[0].Reg(pstate.ssa)
		r2 := v.Args[1].Reg(pstate.ssa)
		r3 := v.Args[2].Reg(pstate.ssa)
		// r = r1*r2 ± r3
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.Reg = r3
		p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: r2})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

	case ssa.OpPPC64MaskIfNotCarry:
		r := v.Reg(pstate.ssa)
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = ppc64.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

	case ssa.OpPPC64ADDconstForCarry:
		r1 := v.Args[0].Reg(pstate.ssa)
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.Reg = r1
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REGTMP // Ignored; this is for the carry effect.

	case ssa.OpPPC64NEG, ssa.OpPPC64FNEG, ssa.OpPPC64FSQRT, ssa.OpPPC64FSQRTS, ssa.OpPPC64FFLOOR, ssa.OpPPC64FTRUNC, ssa.OpPPC64FCEIL, ssa.OpPPC64FCTIDZ, ssa.OpPPC64FCTIWZ, ssa.OpPPC64FCFID, ssa.OpPPC64FCFIDS, ssa.OpPPC64FRSP, ssa.OpPPC64CNTLZD, ssa.OpPPC64CNTLZW, ssa.OpPPC64POPCNTD, ssa.OpPPC64POPCNTW, ssa.OpPPC64POPCNTB, ssa.OpPPC64MFVSRD, ssa.OpPPC64MTVSRD, ssa.OpPPC64FABS, ssa.OpPPC64FNABS, ssa.OpPPC64FROUND:
		r := v.Reg(pstate.ssa)
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)

	case ssa.OpPPC64ADDconst, ssa.OpPPC64ANDconst, ssa.OpPPC64ORconst, ssa.OpPPC64XORconst,
		ssa.OpPPC64SRADconst, ssa.OpPPC64SRAWconst, ssa.OpPPC64SRDconst, ssa.OpPPC64SRWconst, ssa.OpPPC64SLDconst, ssa.OpPPC64SLWconst:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.Reg = v.Args[0].Reg(pstate.ssa)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)

	case ssa.OpPPC64ANDCCconst:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.Reg = v.Args[0].Reg(pstate.ssa)

		if v.Aux != nil {
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = pstate.gc.AuxOffset(v)
		} else {
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = v.AuxInt
		}

		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REGTMP // discard result

	case ssa.OpPPC64MOVDaddr:
		switch v.Aux.(type) {
		default:
			v.Fatalf("aux in MOVDaddr is of unknown type %T", v.Aux)
		case nil:
			// If aux offset and aux int are both 0, and the same
			// input and output regs are used, no instruction
			// needs to be generated, since it would just be
			// addi rx, rx, 0.
			if v.AuxInt != 0 || v.Args[0].Reg(pstate.ssa) != v.Reg(pstate.ssa) {
				p := s.Prog(pstate.gc, ppc64.AMOVD)
				p.From.Type = obj.TYPE_ADDR
				p.From.Reg = v.Args[0].Reg(pstate.ssa)
				p.From.Offset = v.AuxInt
				p.To.Type = obj.TYPE_REG
				p.To.Reg = v.Reg(pstate.ssa)
			}

		case *obj.LSym, *gc.Node:
			p := s.Prog(pstate.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_ADDR
			p.From.Reg = v.Args[0].Reg(pstate.ssa)
			p.To.Type = obj.TYPE_REG
			p.To.Reg = v.Reg(pstate.ssa)
			pstate.gc.AddAux(&p.From, v)

		}

	case ssa.OpPPC64MOVDconst:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)

	case ssa.OpPPC64FMOVDconst, ssa.OpPPC64FMOVSconst:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_FCONST
		p.From.Val = math.Float64frombits(uint64(v.AuxInt))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)

	case ssa.OpPPC64FCMPU, ssa.OpPPC64CMP, ssa.OpPPC64CMPW, ssa.OpPPC64CMPU, ssa.OpPPC64CMPWU:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Args[1].Reg(pstate.ssa)

	case ssa.OpPPC64CMPconst, ssa.OpPPC64CMPUconst, ssa.OpPPC64CMPWconst, ssa.OpPPC64CMPWUconst:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = v.AuxInt

	case ssa.OpPPC64MOVBreg, ssa.OpPPC64MOVBZreg, ssa.OpPPC64MOVHreg, ssa.OpPPC64MOVHZreg, ssa.OpPPC64MOVWreg, ssa.OpPPC64MOVWZreg:
		// Shift in register to required size
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Reg = v.Reg(pstate.ssa)
		p.To.Type = obj.TYPE_REG

	case ssa.OpPPC64MOVDload:

		// MOVDload uses a DS instruction which requires the offset value of the data to be a multiple of 4.
		// For offsets known at compile time, a MOVDload won't be selected, but in the case of a go.string,
		// the offset is not known until link time. If the load of a go.string uses relocation for the
		// offset field of the instruction, and if the offset is not aligned to 4, then a link error will occur.
		// To avoid this problem, the full address of the go.string is computed and loaded into the base register,
		// and that base register is used for the MOVDload using a 0 offset. This problem can only occur with
		// go.string types because other types will have proper alignment.

		gostring := false
		switch n := v.Aux.(type) {
		case *obj.LSym:
			gostring = strings.HasPrefix(n.Name, "go.string.")
		}
		if gostring {
			// Generate full addr of the go.string const
			// including AuxInt
			p := s.Prog(pstate.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_ADDR
			p.From.Reg = v.Args[0].Reg(pstate.ssa)
			pstate.gc.AddAux(&p.From, v)
			p.To.Type = obj.TYPE_REG
			p.To.Reg = v.Reg(pstate.ssa)
			// Load go.string using 0 offset
			p = s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
			p.From.Type = obj.TYPE_MEM
			p.From.Reg = v.Reg(pstate.ssa)
			p.To.Type = obj.TYPE_REG
			p.To.Reg = v.Reg(pstate.ssa)
			break
		}
		// Not a go.string, generate a normal load
		fallthrough

	case ssa.OpPPC64MOVWload, ssa.OpPPC64MOVHload, ssa.OpPPC64MOVWZload, ssa.OpPPC64MOVBZload, ssa.OpPPC64MOVHZload:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)

	case ssa.OpPPC64MOVDBRload, ssa.OpPPC64MOVWBRload, ssa.OpPPC64MOVHBRload:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)

	case ssa.OpPPC64MOVDBRstore, ssa.OpPPC64MOVWBRstore, ssa.OpPPC64MOVHBRstore:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(pstate.ssa)

	case ssa.OpPPC64FMOVDload, ssa.OpPPC64FMOVSload:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(pstate.ssa)

	case ssa.OpPPC64MOVDstorezero, ssa.OpPPC64MOVWstorezero, ssa.OpPPC64MOVHstorezero, ssa.OpPPC64MOVBstorezero:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = ppc64.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.To, v)

	case ssa.OpPPC64MOVDstore, ssa.OpPPC64MOVWstore, ssa.OpPPC64MOVHstore, ssa.OpPPC64MOVBstore:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.To, v)
	case ssa.OpPPC64FMOVDstore, ssa.OpPPC64FMOVSstore:
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.To, v)

	case ssa.OpPPC64Equal,
		ssa.OpPPC64NotEqual,
		ssa.OpPPC64LessThan,
		ssa.OpPPC64FLessThan,
		ssa.OpPPC64LessEqual,
		ssa.OpPPC64GreaterThan,
		ssa.OpPPC64FGreaterThan,
		ssa.OpPPC64GreaterEqual:

		// On Power7 or later, can use isel instruction:
		// for a < b, a > b, a = b:
		//   rtmp := 1
		//   isel rt,rtmp,r0,cond // rt is target in ppc asm

		// for  a >= b, a <= b, a != b:
		//   rtmp := 1
		//   isel rt,0,rtmp,!cond // rt is target in ppc asm

		p := s.Prog(pstate.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = pstate.iselRegs[1]
		iop := pstate.iselOps[v.Op]
		pstate.ssaGenISEL(s, v, iop.cond, pstate.iselRegs[iop.valueIfCond], pstate.iselRegs[1-iop.valueIfCond])

	case ssa.OpPPC64FLessEqual, // These include a second branch for EQ -- dealing with NaN prevents REL= to !REL conversion
		ssa.OpPPC64FGreaterEqual:

		p := s.Prog(pstate.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = pstate.iselRegs[1]
		iop := pstate.iselOps[v.Op]
		pstate.ssaGenISEL(s, v, iop.cond, pstate.iselRegs[iop.valueIfCond], pstate.iselRegs[1-iop.valueIfCond])
		pstate.ssaGenISEL(s, v, ppc64.C_COND_EQ, pstate.iselRegs[1], v.Reg(pstate.ssa))

	case ssa.OpPPC64LoweredZero:

		// unaligned data doesn't hurt performance
		// for these instructions on power8 or later

		// for sizes >= 64 generate a loop as follows:

		// set up loop counter in CTR, used by BC
		//	 MOVD len/32,REG_TMP
		//	 MOVD REG_TMP,CTR
		//	 loop:
		//	 MOVD R0,(R3)
		//	 MOVD R0,8(R3)
		//	 MOVD R0,16(R3)
		//	 MOVD R0,24(R3)
		//	 ADD  $32,R3
		//	 BC   16, 0, loop
		//
		// any remainder is done as described below

		// for sizes < 64 bytes, first clear as many doublewords as possible,
		// then handle the remainder
		//	MOVD R0,(R3)
		//	MOVD R0,8(R3)
		// .... etc.
		//
		// the remainder bytes are cleared using one or more
		// of the following instructions with the appropriate
		// offsets depending which instructions are needed
		//
		//	MOVW R0,n1(R3)	4 bytes
		//	MOVH R0,n2(R3)	2 bytes
		//	MOVB R0,n3(R3)	1 byte
		//
		// 7 bytes: MOVW, MOVH, MOVB
		// 6 bytes: MOVW, MOVH
		// 5 bytes: MOVW, MOVB
		// 3 bytes: MOVH, MOVB

		// each loop iteration does 32 bytes
		ctr := v.AuxInt / 32

		// remainder bytes
		rem := v.AuxInt % 32

		// only generate a loop if there is more
		// than 1 iteration.
		if ctr > 1 {
			// Set up CTR loop counter
			p := s.Prog(pstate.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = ctr
			p.To.Type = obj.TYPE_REG
			p.To.Reg = ppc64.REGTMP

			p = s.Prog(pstate.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = ppc64.REGTMP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = ppc64.REG_CTR

			// generate 4 MOVDs
			// when this is a loop then the top must be saved
			var top *obj.Prog
			for offset := int64(0); offset < 32; offset += 8 {
				// This is the top of loop
				p := s.Prog(pstate.gc, ppc64.AMOVD)
				p.From.Type = obj.TYPE_REG
				p.From.Reg = ppc64.REG_R0
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = v.Args[0].Reg(pstate.ssa)
				p.To.Offset = offset
				// Save the top of loop
				if top == nil {
					top = p
				}
			}

			// Increment address for the
			// 4 doublewords just zeroed.
			p = s.Prog(pstate.gc, ppc64.AADD)
			p.Reg = v.Args[0].Reg(pstate.ssa)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = 32
			p.To.Type = obj.TYPE_REG
			p.To.Reg = v.Args[0].Reg(pstate.ssa)

			// Branch back to top of loop
			// based on CTR
			// BC with BO_BCTR generates bdnz
			p = s.Prog(pstate.gc, ppc64.ABC)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = ppc64.BO_BCTR
			p.Reg = ppc64.REG_R0
			p.To.Type = obj.TYPE_BRANCH
			pstate.gc.Patch(p, top)
		}

		// when ctr == 1 the loop was not generated but
		// there are at least 32 bytes to clear, so add
		// that to the remainder to generate the code
		// to clear those doublewords
		if ctr == 1 {
			rem += 32
		}

		// clear the remainder starting at offset zero
		offset := int64(0)

		// first clear as many doublewords as possible
		// then clear remaining sizes as available
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
			p := s.Prog(pstate.gc, op)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = ppc64.REG_R0
			p.To.Type = obj.TYPE_MEM
			p.To.Reg = v.Args[0].Reg(pstate.ssa)
			p.To.Offset = offset
			rem -= size
			offset += size
		}

	case ssa.OpPPC64LoweredMove:

		// This will be used when moving more
		// than 8 bytes.  Moves start with as
		// as many 8 byte moves as possible, then
		// 4, 2, or 1 byte(s) as remaining.  This will
		// work and be efficient for power8 or later.
		// If there are 64 or more bytes, then a
		// loop is generated to move 32 bytes and
		// update the src and dst addresses on each
		// iteration. When < 64 bytes, the appropriate
		// number of moves are generated based on the
		// size.
		// When moving >= 64 bytes a loop is used
		//	MOVD len/32,REG_TMP
		//	MOVD REG_TMP,CTR
		// top:
		//	MOVD (R4),R7
		//	MOVD 8(R4),R8
		//	MOVD 16(R4),R9
		//	MOVD 24(R4),R10
		//	ADD  R4,$32
		//	MOVD R7,(R3)
		//	MOVD R8,8(R3)
		//	MOVD R9,16(R3)
		//	MOVD R10,24(R3)
		//	ADD  R3,$32
		//	BC 16,0,top
		// Bytes not moved by this loop are moved
		// with a combination of the following instructions,
		// starting with the largest sizes and generating as
		// many as needed, using the appropriate offset value.
		//	MOVD  n(R4),R7
		//	MOVD  R7,n(R3)
		//	MOVW  n1(R4),R7
		//	MOVW  R7,n1(R3)
		//	MOVH  n2(R4),R7
		//	MOVH  R7,n2(R3)
		//	MOVB  n3(R4),R7
		//	MOVB  R7,n3(R3)

		// Each loop iteration moves 32 bytes
		ctr := v.AuxInt / 32

		// Remainder after the loop
		rem := v.AuxInt % 32

		dst_reg := v.Args[0].Reg(pstate.ssa)
		src_reg := v.Args[1].Reg(pstate.ssa)

		// The set of registers used here, must match the clobbered reg list
		// in PPC64Ops.go.
		useregs := []int16{ppc64.REG_R7, ppc64.REG_R8, ppc64.REG_R9, ppc64.REG_R10}
		offset := int64(0)

		// top of the loop
		var top *obj.Prog
		// Only generate looping code when loop counter is > 1 for >= 64 bytes
		if ctr > 1 {
			// Set up the CTR
			p := s.Prog(pstate.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = ctr
			p.To.Type = obj.TYPE_REG
			p.To.Reg = ppc64.REGTMP

			p = s.Prog(pstate.gc, ppc64.AMOVD)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = ppc64.REGTMP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = ppc64.REG_CTR

			// Generate all the MOVDs for loads
			// based off the same register, increasing
			// the offset by 8 for each instruction
			for _, rg := range useregs {
				p := s.Prog(pstate.gc, ppc64.AMOVD)
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
			// increment the src_reg for next iteration
			p = s.Prog(pstate.gc, ppc64.AADD)
			p.Reg = src_reg
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = 32
			p.To.Type = obj.TYPE_REG
			p.To.Reg = src_reg

			// generate the MOVDs for stores, based
			// off the same register, using the same
			// offsets as in the loads.
			offset = int64(0)
			for _, rg := range useregs {
				p := s.Prog(pstate.gc, ppc64.AMOVD)
				p.From.Type = obj.TYPE_REG
				p.From.Reg = rg
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = dst_reg
				p.To.Offset = offset
				offset += 8
			}
			// increment the dst_reg for next iteration
			p = s.Prog(pstate.gc, ppc64.AADD)
			p.Reg = dst_reg
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = 32
			p.To.Type = obj.TYPE_REG
			p.To.Reg = dst_reg

			// BC with BO_BCTR generates bdnz to branch on nonzero CTR
			// to loop top.
			p = s.Prog(pstate.gc, ppc64.ABC)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = ppc64.BO_BCTR
			p.Reg = ppc64.REG_R0
			p.To.Type = obj.TYPE_BRANCH
			pstate.gc.Patch(p, top)

			// src_reg and dst_reg were incremented in the loop, so
			// later instructions start with offset 0.
			offset = int64(0)
		}

		// No loop was generated for one iteration, so
		// add 32 bytes to the remainder to move those bytes.
		if ctr == 1 {
			rem += 32
		}

		// Generate all the remaining load and store pairs, starting with
		// as many 8 byte moves as possible, then 4, 2, 1.
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
			// Load
			p := s.Prog(pstate.gc, op)
			p.To.Type = obj.TYPE_REG
			p.To.Reg = ppc64.REG_R7
			p.From.Type = obj.TYPE_MEM
			p.From.Reg = src_reg
			p.From.Offset = offset

			// Store
			p = s.Prog(pstate.gc, op)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = ppc64.REG_R7
			p.To.Type = obj.TYPE_MEM
			p.To.Reg = dst_reg
			p.To.Offset = offset
			rem -= size
			offset += size
		}

	case ssa.OpPPC64CALLstatic:
		s.Call(pstate.gc, v)

	case ssa.OpPPC64CALLclosure, ssa.OpPPC64CALLinter:
		p := s.Prog(pstate.gc, ppc64.AMOVD)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REG_CTR

		if v.Args[0].Reg(pstate.ssa) != ppc64.REG_R12 {
			v.Fatalf("Function address for %v should be in R12 %d but is in %d", v.LongString(pstate.ssa), ppc64.REG_R12, p.From.Reg)
		}

		pp := s.Call(pstate.gc, v)
		pp.To.Reg = ppc64.REG_CTR

		if pstate.gc.Ctxt.Flag_shared {
			// When compiling Go into PIC, the function we just
			// called via pointer might have been implemented in
			// a separate module and so overwritten the TOC
			// pointer in R2; reload it.
			q := s.Prog(pstate.gc, ppc64.AMOVD)
			q.From.Type = obj.TYPE_MEM
			q.From.Offset = 24
			q.From.Reg = ppc64.REGSP
			q.To.Type = obj.TYPE_REG
			q.To.Reg = ppc64.REG_R2
		}

	case ssa.OpPPC64LoweredWB:
		p := s.Prog(pstate.gc, obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = v.Aux.(*obj.LSym)

	case ssa.OpPPC64LoweredNilCheck:
		// Issue a load which will fault if arg is nil.
		p := s.Prog(pstate.gc, ppc64.AMOVBZ)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(pstate.ssa)
		pstate.gc.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REGTMP
		if pstate.gc.Debug_checknil != 0 && v.Pos.Line() > 1 { // v.Pos.Line()==1 in generated wrappers
			pstate.gc.Warnl(v.Pos, "generated nil check")
		}

	case ssa.OpPPC64InvertFlags:
		v.Fatalf("InvertFlags should never make it to codegen %v", v.LongString(pstate.ssa))
	case ssa.OpPPC64FlagEQ, ssa.OpPPC64FlagLT, ssa.OpPPC64FlagGT:
		v.Fatalf("Flag* ops should never make it to codegen %v", v.LongString(pstate.ssa))
	case ssa.OpClobber:
	// TODO: implement for clobberdead experiment. Nop is ok for now.
	default:
		v.Fatalf("genValue not implemented: %s", v.LongString(pstate.ssa))
	}
}

func (pstate *PackageState) ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	switch b.Kind {
	case ssa.BlockDefer:
		// defer returns in R3:
		// 0 if we should continue executing
		// 1 if we should jump to deferreturn call
		p := s.Prog(pstate.gc, ppc64.ACMP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = ppc64.REG_R3
		p.To.Type = obj.TYPE_REG
		p.To.Reg = ppc64.REG_R0

		p = s.Prog(pstate.gc, ppc64.ABNE)
		p.To.Type = obj.TYPE_BRANCH
		s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[1].Block()})
		if b.Succs[0].Block() != next {
			p := s.Prog(pstate.gc, obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}

	case ssa.BlockPlain:
		if b.Succs[0].Block() != next {
			p := s.Prog(pstate.gc, obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}
	case ssa.BlockExit:
		s.Prog(pstate.gc, obj.AUNDEF) // tell plive.go that we never reach here
	case ssa.BlockRet:
		s.Prog(pstate.gc, obj.ARET)
	case ssa.BlockRetJmp:
		p := s.Prog(pstate.gc, obj.AJMP)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = b.Aux.(*obj.LSym)

	case ssa.BlockPPC64EQ, ssa.BlockPPC64NE,
		ssa.BlockPPC64LT, ssa.BlockPPC64GE,
		ssa.BlockPPC64LE, ssa.BlockPPC64GT,
		ssa.BlockPPC64FLT, ssa.BlockPPC64FGE,
		ssa.BlockPPC64FLE, ssa.BlockPPC64FGT:
		jmp := pstate.blockJump[b.Kind]
		switch next {
		case b.Succs[0].Block():
			s.Br(pstate.gc, jmp.invasm, b.Succs[1].Block())
			if jmp.invasmun {
				// TODO: The second branch is probably predict-not-taken since it is for FP unordered
				s.Br(pstate.gc, ppc64.ABVS, b.Succs[1].Block())
			}
		case b.Succs[1].Block():
			s.Br(pstate.gc, jmp.asm, b.Succs[0].Block())
			if jmp.asmeq {
				s.Br(pstate.gc, ppc64.ABEQ, b.Succs[0].Block())
			}
		default:
			if b.Likely != ssa.BranchUnlikely {
				s.Br(pstate.gc, jmp.asm, b.Succs[0].Block())
				if jmp.asmeq {
					s.Br(pstate.gc, ppc64.ABEQ, b.Succs[0].Block())
				}
				s.Br(pstate.gc, obj.AJMP, b.Succs[1].Block())
			} else {
				s.Br(pstate.gc, jmp.invasm, b.Succs[1].Block())
				if jmp.invasmun {
					// TODO: The second branch is probably predict-not-taken since it is for FP unordered
					s.Br(pstate.gc, ppc64.ABVS, b.Succs[1].Block())
				}
				s.Br(pstate.gc, obj.AJMP, b.Succs[0].Block())
			}
		}
	default:
		b.Fatalf("branch not implemented: %s. Control: %s", b.LongString(pstate.ssa), b.Control.LongString(pstate.ssa))
	}
}

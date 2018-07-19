package arm64

import (
	"math"

	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/arm64"
)

// loadByType returns the load instruction of the given type.
func (psess *PackageSession) loadByType(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(psess.types) {
		case 4:
			return arm64.AFMOVS
		case 8:
			return arm64.AFMOVD
		}
	} else {
		switch t.Size(psess.types) {
		case 1:
			if t.IsSigned() {
				return arm64.AMOVB
			} else {
				return arm64.AMOVBU
			}
		case 2:
			if t.IsSigned() {
				return arm64.AMOVH
			} else {
				return arm64.AMOVHU
			}
		case 4:
			if t.IsSigned() {
				return arm64.AMOVW
			} else {
				return arm64.AMOVWU
			}
		case 8:
			return arm64.AMOVD
		}
	}
	panic("bad load type")
}

// storeByType returns the store instruction of the given type.
func (psess *PackageSession) storeByType(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(psess.types) {
		case 4:
			return arm64.AFMOVS
		case 8:
			return arm64.AFMOVD
		}
	} else {
		switch t.Size(psess.types) {
		case 1:
			return arm64.AMOVB
		case 2:
			return arm64.AMOVH
		case 4:
			return arm64.AMOVW
		case 8:
			return arm64.AMOVD
		}
	}
	panic("bad store type")
}

// makeshift encodes a register shifted by a constant, used as an Offset in Prog
func makeshift(reg int16, typ int64, s int64) int64 {
	return int64(reg&31)<<16 | typ | (s&63)<<10
}

// genshift generates a Prog for r = r0 op (r1 shifted by n)
func (psess *PackageSession) genshift(s *gc.SSAGenState, as obj.As, r0, r1, r int16, typ int64, n int64) *obj.Prog {
	p := s.Prog(psess.gc, as)
	p.From.Type = obj.TYPE_SHIFT
	p.From.Offset = makeshift(r1, typ, n)
	p.Reg = r0
	if r != 0 {
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	}
	return p
}

// generate the memory operand for the indexed load/store instructions
func (psess *PackageSession) genIndexedOperand(v *ssa.Value) obj.Addr {

	mop := obj.Addr{Type: obj.TYPE_MEM, Reg: v.Args[0].Reg(psess.ssa)}
	switch v.Op {
	case ssa.OpARM64MOVDloadidx8, ssa.OpARM64MOVDstoreidx8, ssa.OpARM64MOVDstorezeroidx8:
		mop.Index = arm64.REG_LSL | 3<<5 | v.Args[1].Reg(psess.ssa)&31
	case ssa.OpARM64MOVWloadidx4, ssa.OpARM64MOVWUloadidx4, ssa.OpARM64MOVWstoreidx4, ssa.OpARM64MOVWstorezeroidx4:
		mop.Index = arm64.REG_LSL | 2<<5 | v.Args[1].Reg(psess.ssa)&31
	case ssa.OpARM64MOVHloadidx2, ssa.OpARM64MOVHUloadidx2, ssa.OpARM64MOVHstoreidx2, ssa.OpARM64MOVHstorezeroidx2:
		mop.Index = arm64.REG_LSL | 1<<5 | v.Args[1].Reg(psess.ssa)&31
	default:
		mop.Index = v.Args[1].Reg(psess.ssa)
	}
	return mop
}

func (psess *PackageSession) ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpCopy, ssa.OpARM64MOVDreg:
		if v.Type.IsMemory(psess.types) {
			return
		}
		x := v.Args[0].Reg(psess.ssa)
		y := v.Reg(psess.ssa)
		if x == y {
			return
		}
		as := arm64.AMOVD
		if v.Type.IsFloat() {
			switch v.Type.Size(psess.types) {
			case 4:
				as = arm64.AFMOVS
			case 8:
				as = arm64.AFMOVD
			default:
				panic("bad float size")
			}
		}
		p := s.Prog(psess.gc, as)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x
		p.To.Type = obj.TYPE_REG
		p.To.Reg = y
	case ssa.OpARM64MOVDnop:
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
	case ssa.OpARM64ADD,
		ssa.OpARM64SUB,
		ssa.OpARM64AND,
		ssa.OpARM64OR,
		ssa.OpARM64XOR,
		ssa.OpARM64BIC,
		ssa.OpARM64EON,
		ssa.OpARM64ORN,
		ssa.OpARM64MUL,
		ssa.OpARM64MULW,
		ssa.OpARM64MNEG,
		ssa.OpARM64MNEGW,
		ssa.OpARM64MULH,
		ssa.OpARM64UMULH,
		ssa.OpARM64MULL,
		ssa.OpARM64UMULL,
		ssa.OpARM64DIV,
		ssa.OpARM64UDIV,
		ssa.OpARM64DIVW,
		ssa.OpARM64UDIVW,
		ssa.OpARM64MOD,
		ssa.OpARM64UMOD,
		ssa.OpARM64MODW,
		ssa.OpARM64UMODW,
		ssa.OpARM64SLL,
		ssa.OpARM64SRL,
		ssa.OpARM64SRA,
		ssa.OpARM64FADDS,
		ssa.OpARM64FADDD,
		ssa.OpARM64FSUBS,
		ssa.OpARM64FSUBD,
		ssa.OpARM64FMULS,
		ssa.OpARM64FMULD,
		ssa.OpARM64FNMULS,
		ssa.OpARM64FNMULD,
		ssa.OpARM64FDIVS,
		ssa.OpARM64FDIVD:
		r := v.Reg(psess.ssa)
		r1 := v.Args[0].Reg(psess.ssa)
		r2 := v.Args[1].Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r2
		p.Reg = r1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpARM64FMADDS,
		ssa.OpARM64FMADDD,
		ssa.OpARM64FNMADDS,
		ssa.OpARM64FNMADDD,
		ssa.OpARM64FMSUBS,
		ssa.OpARM64FMSUBD,
		ssa.OpARM64FNMSUBS,
		ssa.OpARM64FNMSUBD:
		rt := v.Reg(psess.ssa)
		ra := v.Args[0].Reg(psess.ssa)
		rm := v.Args[1].Reg(psess.ssa)
		rn := v.Args[2].Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.Reg = ra
		p.From.Type = obj.TYPE_REG
		p.From.Reg = rm
		p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: rn})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = rt
	case ssa.OpARM64ADDconst,
		ssa.OpARM64SUBconst,
		ssa.OpARM64ANDconst,
		ssa.OpARM64ORconst,
		ssa.OpARM64XORconst,
		ssa.OpARM64SLLconst,
		ssa.OpARM64SRLconst,
		ssa.OpARM64SRAconst,
		ssa.OpARM64RORconst,
		ssa.OpARM64RORWconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64EXTRconst,
		ssa.OpARM64EXTRWconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: v.Args[0].Reg(psess.ssa)})
		p.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64ADDshiftLL,
		ssa.OpARM64SUBshiftLL,
		ssa.OpARM64ANDshiftLL,
		ssa.OpARM64ORshiftLL,
		ssa.OpARM64XORshiftLL,
		ssa.OpARM64EONshiftLL,
		ssa.OpARM64ORNshiftLL,
		ssa.OpARM64BICshiftLL:
		psess.
			genshift(s, v.Op.Asm(psess.ssa), v.Args[0].Reg(psess.ssa), v.Args[1].Reg(psess.ssa), v.Reg(psess.ssa), arm64.SHIFT_LL, v.AuxInt)
	case ssa.OpARM64ADDshiftRL,
		ssa.OpARM64SUBshiftRL,
		ssa.OpARM64ANDshiftRL,
		ssa.OpARM64ORshiftRL,
		ssa.OpARM64XORshiftRL,
		ssa.OpARM64EONshiftRL,
		ssa.OpARM64ORNshiftRL,
		ssa.OpARM64BICshiftRL:
		psess.
			genshift(s, v.Op.Asm(psess.ssa), v.Args[0].Reg(psess.ssa), v.Args[1].Reg(psess.ssa), v.Reg(psess.ssa), arm64.SHIFT_LR, v.AuxInt)
	case ssa.OpARM64ADDshiftRA,
		ssa.OpARM64SUBshiftRA,
		ssa.OpARM64ANDshiftRA,
		ssa.OpARM64ORshiftRA,
		ssa.OpARM64XORshiftRA,
		ssa.OpARM64EONshiftRA,
		ssa.OpARM64ORNshiftRA,
		ssa.OpARM64BICshiftRA:
		psess.
			genshift(s, v.Op.Asm(psess.ssa), v.Args[0].Reg(psess.ssa), v.Args[1].Reg(psess.ssa), v.Reg(psess.ssa), arm64.SHIFT_AR, v.AuxInt)
	case ssa.OpARM64MOVDconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64FMOVSconst,
		ssa.OpARM64FMOVDconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_FCONST
		p.From.Val = math.Float64frombits(uint64(v.AuxInt))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64CMP,
		ssa.OpARM64CMPW,
		ssa.OpARM64CMN,
		ssa.OpARM64CMNW,
		ssa.OpARM64TST,
		ssa.OpARM64TSTW,
		ssa.OpARM64FCMPS,
		ssa.OpARM64FCMPD:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.Reg = v.Args[0].Reg(psess.ssa)
	case ssa.OpARM64CMPconst,
		ssa.OpARM64CMPWconst,
		ssa.OpARM64CMNconst,
		ssa.OpARM64CMNWconst,
		ssa.OpARM64TSTconst,
		ssa.OpARM64TSTWconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.Reg = v.Args[0].Reg(psess.ssa)
	case ssa.OpARM64CMPshiftLL:
		psess.
			genshift(s, v.Op.Asm(psess.ssa), v.Args[0].Reg(psess.ssa), v.Args[1].Reg(psess.ssa), 0, arm64.SHIFT_LL, v.AuxInt)
	case ssa.OpARM64CMPshiftRL:
		psess.
			genshift(s, v.Op.Asm(psess.ssa), v.Args[0].Reg(psess.ssa), v.Args[1].Reg(psess.ssa), 0, arm64.SHIFT_LR, v.AuxInt)
	case ssa.OpARM64CMPshiftRA:
		psess.
			genshift(s, v.Op.Asm(psess.ssa), v.Args[0].Reg(psess.ssa), v.Args[1].Reg(psess.ssa), 0, arm64.SHIFT_AR, v.AuxInt)
	case ssa.OpARM64MOVDaddr:
		p := s.Prog(psess.gc, arm64.AMOVD)
		p.From.Type = obj.TYPE_ADDR
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

		var wantreg string

		switch v.Aux.(type) {
		default:
			v.Fatalf("aux is of unknown type %T", v.Aux)
		case *obj.LSym:
			wantreg = "SB"
			psess.gc.
				AddAux(&p.From, v)
		case *gc.Node:
			wantreg = "SP"
			psess.gc.
				AddAux(&p.From, v)
		case nil:

			wantreg = "SP"
			p.From.Offset = v.AuxInt
		}
		if reg := v.Args[0].RegName(psess.ssa); reg != wantreg {
			v.Fatalf("bad reg %s for symbol type %T, want %s", reg, v.Aux, wantreg)
		}
	case ssa.OpARM64MOVBload,
		ssa.OpARM64MOVBUload,
		ssa.OpARM64MOVHload,
		ssa.OpARM64MOVHUload,
		ssa.OpARM64MOVWload,
		ssa.OpARM64MOVWUload,
		ssa.OpARM64MOVDload,
		ssa.OpARM64FMOVSload,
		ssa.OpARM64FMOVDload:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64MOVBloadidx,
		ssa.OpARM64MOVBUloadidx,
		ssa.OpARM64MOVHloadidx,
		ssa.OpARM64MOVHUloadidx,
		ssa.OpARM64MOVWloadidx,
		ssa.OpARM64MOVWUloadidx,
		ssa.OpARM64MOVDloadidx,
		ssa.OpARM64MOVHloadidx2,
		ssa.OpARM64MOVHUloadidx2,
		ssa.OpARM64MOVWloadidx4,
		ssa.OpARM64MOVWUloadidx4,
		ssa.OpARM64MOVDloadidx8:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From = psess.genIndexedOperand(v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64LDAR,
		ssa.OpARM64LDARW:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0(psess.ssa)
	case ssa.OpARM64MOVBstore,
		ssa.OpARM64MOVHstore,
		ssa.OpARM64MOVWstore,
		ssa.OpARM64MOVDstore,
		ssa.OpARM64FMOVSstore,
		ssa.OpARM64FMOVDstore,
		ssa.OpARM64STLR,
		ssa.OpARM64STLRW:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpARM64MOVBstoreidx,
		ssa.OpARM64MOVHstoreidx,
		ssa.OpARM64MOVWstoreidx,
		ssa.OpARM64MOVDstoreidx,
		ssa.OpARM64MOVHstoreidx2,
		ssa.OpARM64MOVWstoreidx4,
		ssa.OpARM64MOVDstoreidx8:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.To = psess.genIndexedOperand(v)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[2].Reg(psess.ssa)
	case ssa.OpARM64STP:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REGREG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.From.Offset = int64(v.Args[2].Reg(psess.ssa))
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpARM64MOVBstorezero,
		ssa.OpARM64MOVHstorezero,
		ssa.OpARM64MOVWstorezero,
		ssa.OpARM64MOVDstorezero:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = arm64.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpARM64MOVBstorezeroidx,
		ssa.OpARM64MOVHstorezeroidx,
		ssa.OpARM64MOVWstorezeroidx,
		ssa.OpARM64MOVDstorezeroidx,
		ssa.OpARM64MOVHstorezeroidx2,
		ssa.OpARM64MOVWstorezeroidx4,
		ssa.OpARM64MOVDstorezeroidx8:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.To = psess.genIndexedOperand(v)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = arm64.REGZERO
	case ssa.OpARM64MOVQstorezero:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REGREG
		p.From.Reg = arm64.REGZERO
		p.From.Offset = int64(arm64.REGZERO)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpARM64BFI,
		ssa.OpARM64BFXIL:
		r := v.Reg(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt >> 8
		p.SetFrom3(obj.Addr{Type: obj.TYPE_CONST, Offset: v.AuxInt & 0xff})
		p.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.OpARM64SBFIZ,
		ssa.OpARM64SBFX,
		ssa.OpARM64UBFIZ,
		ssa.OpARM64UBFX:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt >> 8
		p.SetFrom3(obj.Addr{Type: obj.TYPE_CONST, Offset: v.AuxInt & 0xff})
		p.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64LoweredMuluhilo:
		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)
		p := s.Prog(psess.gc, arm64.AUMULH)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0(psess.ssa)
		p1 := s.Prog(psess.gc, arm64.AMUL)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.Reg = r0
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = v.Reg1(psess.ssa)
	case ssa.OpARM64LoweredAtomicExchange64,
		ssa.OpARM64LoweredAtomicExchange32:

		ld := arm64.ALDAXR
		st := arm64.ASTLXR
		if v.Op == ssa.OpARM64LoweredAtomicExchange32 {
			ld = arm64.ALDAXRW
			st = arm64.ASTLXRW
		}
		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)
		out := v.Reg0(psess.ssa)
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
		p1.RegTo2 = arm64.REGTMP
		p2 := s.Prog(psess.gc, arm64.ACBNZ)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = arm64.REGTMP
		p2.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p2, p)
	case ssa.OpARM64LoweredAtomicAdd64,
		ssa.OpARM64LoweredAtomicAdd32:

		ld := arm64.ALDAXR
		st := arm64.ASTLXR
		if v.Op == ssa.OpARM64LoweredAtomicAdd32 {
			ld = arm64.ALDAXRW
			st = arm64.ASTLXRW
		}
		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)
		out := v.Reg0(psess.ssa)
		p := s.Prog(psess.gc, ld)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = out
		p1 := s.Prog(psess.gc, arm64.AADD)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = out
		p2 := s.Prog(psess.gc, st)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = out
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = r0
		p2.RegTo2 = arm64.REGTMP
		p3 := s.Prog(psess.gc, arm64.ACBNZ)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = arm64.REGTMP
		p3.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p3, p)
	case ssa.OpARM64LoweredAtomicAdd64Variant,
		ssa.OpARM64LoweredAtomicAdd32Variant:

		op := arm64.ALDADDALD
		if v.Op == ssa.OpARM64LoweredAtomicAdd32Variant {
			op = arm64.ALDADDALW
		}
		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)
		out := v.Reg0(psess.ssa)
		p := s.Prog(psess.gc, op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r1
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = r0
		p.RegTo2 = out
		p1 := s.Prog(psess.gc, arm64.AADD)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = out
	case ssa.OpARM64LoweredAtomicCas64,
		ssa.OpARM64LoweredAtomicCas32:

		ld := arm64.ALDAXR
		st := arm64.ASTLXR
		cmp := arm64.ACMP
		if v.Op == ssa.OpARM64LoweredAtomicCas32 {
			ld = arm64.ALDAXRW
			st = arm64.ASTLXRW
			cmp = arm64.ACMPW
		}
		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)
		r2 := v.Args[2].Reg(psess.ssa)
		out := v.Reg0(psess.ssa)
		p := s.Prog(psess.gc, ld)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = arm64.REGTMP
		p1 := s.Prog(psess.gc, cmp)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.Reg = arm64.REGTMP
		p2 := s.Prog(psess.gc, arm64.ABNE)
		p2.To.Type = obj.TYPE_BRANCH
		p3 := s.Prog(psess.gc, st)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = r2
		p3.To.Type = obj.TYPE_MEM
		p3.To.Reg = r0
		p3.RegTo2 = arm64.REGTMP
		p4 := s.Prog(psess.gc, arm64.ACBNZ)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = arm64.REGTMP
		p4.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p4, p)
		p5 := s.Prog(psess.gc, arm64.ACSET)
		p5.From.Type = obj.TYPE_REG
		p5.From.Reg = arm64.COND_EQ
		p5.To.Type = obj.TYPE_REG
		p5.To.Reg = out
		psess.gc.
			Patch(p2, p5)
	case ssa.OpARM64LoweredAtomicAnd8,
		ssa.OpARM64LoweredAtomicOr8:

		r0 := v.Args[0].Reg(psess.ssa)
		r1 := v.Args[1].Reg(psess.ssa)
		out := v.Reg0(psess.ssa)
		p := s.Prog(psess.gc, arm64.ALDAXRB)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = out
		p1 := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = r1
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = out
		p2 := s.Prog(psess.gc, arm64.ASTLXRB)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = out
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = r0
		p2.RegTo2 = arm64.REGTMP
		p3 := s.Prog(psess.gc, arm64.ACBNZ)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = arm64.REGTMP
		p3.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p3, p)
	case ssa.OpARM64MOVBreg,
		ssa.OpARM64MOVBUreg,
		ssa.OpARM64MOVHreg,
		ssa.OpARM64MOVHUreg,
		ssa.OpARM64MOVWreg,
		ssa.OpARM64MOVWUreg:
		a := v.Args[0]
		for a.Op == ssa.OpCopy || a.Op == ssa.OpARM64MOVDreg {
			a = a.Args[0]
		}
		if a.Op == ssa.OpLoadReg {
			t := a.Type
			switch {
			case v.Op == ssa.OpARM64MOVBreg && t.Size(psess.types) == 1 && t.IsSigned(),
				v.Op == ssa.OpARM64MOVBUreg && t.Size(psess.types) == 1 && !t.IsSigned(),
				v.Op == ssa.OpARM64MOVHreg && t.Size(psess.types) == 2 && t.IsSigned(),
				v.Op == ssa.OpARM64MOVHUreg && t.Size(psess.types) == 2 && !t.IsSigned(),
				v.Op == ssa.OpARM64MOVWreg && t.Size(psess.types) == 4 && t.IsSigned(),
				v.Op == ssa.OpARM64MOVWUreg && t.Size(psess.types) == 4 && !t.IsSigned():

				if v.Reg(psess.ssa) == v.Args[0].Reg(psess.ssa) {
					return
				}
				p := s.Prog(psess.gc, arm64.AMOVD)
				p.From.Type = obj.TYPE_REG
				p.From.Reg = v.Args[0].Reg(psess.ssa)
				p.To.Type = obj.TYPE_REG
				p.To.Reg = v.Reg(psess.ssa)
				return
			default:
			}
		}
		fallthrough
	case ssa.OpARM64MVN,
		ssa.OpARM64NEG,
		ssa.OpARM64FMOVDfpgp,
		ssa.OpARM64FMOVDgpfp,
		ssa.OpARM64FNEGS,
		ssa.OpARM64FNEGD,
		ssa.OpARM64FSQRTD,
		ssa.OpARM64FCVTZSSW,
		ssa.OpARM64FCVTZSDW,
		ssa.OpARM64FCVTZUSW,
		ssa.OpARM64FCVTZUDW,
		ssa.OpARM64FCVTZSS,
		ssa.OpARM64FCVTZSD,
		ssa.OpARM64FCVTZUS,
		ssa.OpARM64FCVTZUD,
		ssa.OpARM64SCVTFWS,
		ssa.OpARM64SCVTFWD,
		ssa.OpARM64SCVTFS,
		ssa.OpARM64SCVTFD,
		ssa.OpARM64UCVTFWS,
		ssa.OpARM64UCVTFWD,
		ssa.OpARM64UCVTFS,
		ssa.OpARM64UCVTFD,
		ssa.OpARM64FCVTSD,
		ssa.OpARM64FCVTDS,
		ssa.OpARM64REV,
		ssa.OpARM64REVW,
		ssa.OpARM64REV16W,
		ssa.OpARM64RBIT,
		ssa.OpARM64RBITW,
		ssa.OpARM64CLZ,
		ssa.OpARM64CLZW,
		ssa.OpARM64FRINTAD,
		ssa.OpARM64FRINTMD,
		ssa.OpARM64FRINTPD,
		ssa.OpARM64FRINTZD:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64LoweredRound32F, ssa.OpARM64LoweredRound64F:

	case ssa.OpARM64VCNT:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = (v.Args[0].Reg(psess.ssa)-arm64.REG_F0)&31 + arm64.REG_ARNG + ((arm64.ARNG_8B & 15) << 5)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = (v.Reg(psess.ssa)-arm64.REG_F0)&31 + arm64.REG_ARNG + ((arm64.ARNG_8B & 15) << 5)
	case ssa.OpARM64VUADDLV:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = (v.Args[0].Reg(psess.ssa)-arm64.REG_F0)&31 + arm64.REG_ARNG + ((arm64.ARNG_8B & 15) << 5)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa) - arm64.REG_F0 + arm64.REG_V0
	case ssa.OpARM64CSEL, ssa.OpARM64CSEL0:
		r1 := int16(arm64.REGZERO)
		if v.Op != ssa.OpARM64CSEL0 {
			r1 = v.Args[1].Reg(psess.ssa)
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = psess.condBits[v.Aux.(ssa.Op)]
		p.Reg = v.Args[0].Reg(psess.ssa)
		p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: r1})
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64DUFFZERO:

		p := s.Prog(psess.gc, obj.ADUFFZERO)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = psess.gc.Duffzero
		p.To.Offset = v.AuxInt
	case ssa.OpARM64LoweredZero:

		p := s.Prog(psess.gc, arm64.ASTP)
		p.Scond = arm64.C_XPOST
		p.From.Type = obj.TYPE_REGREG
		p.From.Reg = arm64.REGZERO
		p.From.Offset = int64(arm64.REGZERO)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = arm64.REG_R16
		p.To.Offset = 16
		p2 := s.Prog(psess.gc, arm64.ACMP)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = v.Args[1].Reg(psess.ssa)
		p2.Reg = arm64.REG_R16
		p3 := s.Prog(psess.gc, arm64.ABLE)
		p3.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p3, p)
	case ssa.OpARM64DUFFCOPY:
		p := s.Prog(psess.gc, obj.ADUFFCOPY)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = psess.gc.Duffcopy
		p.To.Offset = v.AuxInt
	case ssa.OpARM64LoweredMove:

		p := s.Prog(psess.gc, arm64.AMOVD)
		p.Scond = arm64.C_XPOST
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = arm64.REG_R16
		p.From.Offset = 8
		p.To.Type = obj.TYPE_REG
		p.To.Reg = arm64.REGTMP
		p2 := s.Prog(psess.gc, arm64.AMOVD)
		p2.Scond = arm64.C_XPOST
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = arm64.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = arm64.REG_R17
		p2.To.Offset = 8
		p3 := s.Prog(psess.gc, arm64.ACMP)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = v.Args[2].Reg(psess.ssa)
		p3.Reg = arm64.REG_R16
		p4 := s.Prog(psess.gc, arm64.ABLE)
		p4.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p4, p)
	case ssa.OpARM64CALLstatic, ssa.OpARM64CALLclosure, ssa.OpARM64CALLinter:
		s.Call(psess.gc, v)
	case ssa.OpARM64LoweredWB:
		p := s.Prog(psess.gc, obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = v.Aux.(*obj.LSym)
	case ssa.OpARM64LoweredNilCheck:

		p := s.Prog(psess.gc, arm64.AMOVB)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = arm64.REGTMP
		if psess.gc.Debug_checknil != 0 && v.Pos.Line() > 1 {
			psess.gc.
				Warnl(v.Pos, "generated nil check")
		}
	case ssa.OpARM64Equal,
		ssa.OpARM64NotEqual,
		ssa.OpARM64LessThan,
		ssa.OpARM64LessEqual,
		ssa.OpARM64GreaterThan,
		ssa.OpARM64GreaterEqual,
		ssa.OpARM64LessThanU,
		ssa.OpARM64LessEqualU,
		ssa.OpARM64GreaterThanU,
		ssa.OpARM64GreaterEqualU:

		p := s.Prog(psess.gc, arm64.ACSET)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = psess.condBits[v.Op]
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64LoweredGetClosurePtr:
		psess.gc.
			CheckLoweredGetClosurePtr(v)
	case ssa.OpARM64LoweredGetCallerSP:

		p := s.Prog(psess.gc, arm64.AMOVD)
		p.From.Type = obj.TYPE_ADDR
		p.From.Offset = -psess.gc.Ctxt.FixedFrameSize()
		p.From.Name = obj.NAME_PARAM
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64LoweredGetCallerPC:
		p := s.Prog(psess.gc, obj.AGETCALLERPC)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpARM64FlagEQ,
		ssa.OpARM64FlagLT_ULT,
		ssa.OpARM64FlagLT_UGT,
		ssa.OpARM64FlagGT_ULT,
		ssa.OpARM64FlagGT_UGT:
		v.Fatalf("Flag* ops should never make it to codegen %v", v.LongString(psess.ssa))
	case ssa.OpARM64InvertFlags:
		v.Fatalf("InvertFlags should never make it to codegen %v", v.LongString(psess.ssa))
	case ssa.OpClobber:

	default:
		v.Fatalf("genValue not implemented: %s", v.LongString(psess.ssa))
	}
}

func (psess *PackageSession) ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	switch b.Kind {
	case ssa.BlockPlain:
		if b.Succs[0].Block() != next {
			p := s.Prog(psess.gc, obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}

	case ssa.BlockDefer:

		p := s.Prog(psess.gc, arm64.ACMP)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 0
		p.Reg = arm64.REG_R0
		p = s.Prog(psess.gc, arm64.ABNE)
		p.To.Type = obj.TYPE_BRANCH
		s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[1].Block()})
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
		p := s.Prog(psess.gc, obj.ARET)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = b.Aux.(*obj.LSym)

	case ssa.BlockARM64EQ, ssa.BlockARM64NE,
		ssa.BlockARM64LT, ssa.BlockARM64GE,
		ssa.BlockARM64LE, ssa.BlockARM64GT,
		ssa.BlockARM64ULT, ssa.BlockARM64UGT,
		ssa.BlockARM64ULE, ssa.BlockARM64UGE,
		ssa.BlockARM64Z, ssa.BlockARM64NZ,
		ssa.BlockARM64ZW, ssa.BlockARM64NZW:
		jmp := psess.blockJump[b.Kind]
		var p *obj.Prog
		switch next {
		case b.Succs[0].Block():
			p = s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
		case b.Succs[1].Block():
			p = s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
		default:
			if b.Likely != ssa.BranchUnlikely {
				p = s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
				s.Br(psess.gc, obj.AJMP, b.Succs[1].Block())
			} else {
				p = s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
				s.Br(psess.gc, obj.AJMP, b.Succs[0].Block())
			}
		}
		if !b.Control.Type.IsFlags(psess.types) {
			p.From.Type = obj.TYPE_REG
			p.From.Reg = b.Control.Reg(psess.ssa)
		}
	case ssa.BlockARM64TBZ, ssa.BlockARM64TBNZ:
		jmp := psess.blockJump[b.Kind]
		var p *obj.Prog
		switch next {
		case b.Succs[0].Block():
			p = s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
		case b.Succs[1].Block():
			p = s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
		default:
			if b.Likely != ssa.BranchUnlikely {
				p = s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
				s.Br(psess.gc, obj.AJMP, b.Succs[1].Block())
			} else {
				p = s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
				s.Br(psess.gc, obj.AJMP, b.Succs[0].Block())
			}
		}
		p.From.Offset = b.Aux.(int64)
		p.From.Type = obj.TYPE_CONST
		p.Reg = b.Control.Reg(psess.ssa)

	default:
		b.Fatalf("branch not implemented: %s. Control: %s", b.LongString(psess.ssa), b.Control.LongString(psess.ssa))
	}
}

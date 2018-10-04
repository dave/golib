package ppc64

type PackageState struct {
	gc        *gc.PackageState
	obj       *obj.PackageState
	objabi    *objabi.PackageState
	ppc64     *ppc64.PackageState
	ssa       *ssa.PackageState
	types     *types.PackageState
	blockJump [81]struct {
		asm      obj.As
		invasm   obj.As
		asmeq    bool
		invasmun bool
	}
	iselOps  map[ssa.Op]iselOp
	iselRegs [2]int16
}

func NewPackageState(gc_pstate *gc.PackageState, ssa_pstate *ssa.PackageState, types_pstate *types.PackageState, obj_pstate *obj.PackageState, ppc64_pstate *ppc64.PackageState, objabi_pstate *objabi.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.gc = gc_pstate
	pstate.ssa = ssa_pstate
	pstate.types = types_pstate
	pstate.obj = obj_pstate
	pstate.ppc64 = ppc64_pstate
	pstate.objabi = objabi_pstate
	pstate.iselRegs = [2]int16{ppc64.REG_R0, ppc64.REGTMP}
	pstate.iselOps = map[ssa.Op]iselOp{
		ssa.OpPPC64Equal:         iselOp{cond: ppc64.C_COND_EQ, valueIfCond: 1},
		ssa.OpPPC64NotEqual:      iselOp{cond: ppc64.C_COND_EQ, valueIfCond: 0},
		ssa.OpPPC64LessThan:      iselOp{cond: ppc64.C_COND_LT, valueIfCond: 1},
		ssa.OpPPC64GreaterEqual:  iselOp{cond: ppc64.C_COND_LT, valueIfCond: 0},
		ssa.OpPPC64GreaterThan:   iselOp{cond: ppc64.C_COND_GT, valueIfCond: 1},
		ssa.OpPPC64LessEqual:     iselOp{cond: ppc64.C_COND_GT, valueIfCond: 0},
		ssa.OpPPC64FLessThan:     iselOp{cond: ppc64.C_COND_LT, valueIfCond: 1},
		ssa.OpPPC64FGreaterThan:  iselOp{cond: ppc64.C_COND_GT, valueIfCond: 1},
		ssa.OpPPC64FLessEqual:    iselOp{cond: ppc64.C_COND_LT, valueIfCond: 1}, // 2 comparisons, 2nd is EQ
		ssa.OpPPC64FGreaterEqual: iselOp{cond: ppc64.C_COND_GT, valueIfCond: 1}, // 2 comparisons, 2nd is EQ
	}
	pstate.blockJump = [...]struct {
		asm, invasm     obj.As
		asmeq, invasmun bool
	}{
		ssa.BlockPPC64EQ: {ppc64.ABEQ, ppc64.ABNE, false, false},
		ssa.BlockPPC64NE: {ppc64.ABNE, ppc64.ABEQ, false, false},

		ssa.BlockPPC64LT: {ppc64.ABLT, ppc64.ABGE, false, false},
		ssa.BlockPPC64GE: {ppc64.ABGE, ppc64.ABLT, false, false},
		ssa.BlockPPC64LE: {ppc64.ABLE, ppc64.ABGT, false, false},
		ssa.BlockPPC64GT: {ppc64.ABGT, ppc64.ABLE, false, false},

		// TODO: need to work FP comparisons into block jumps
		ssa.BlockPPC64FLT: {ppc64.ABLT, ppc64.ABGE, false, false},
		ssa.BlockPPC64FGE: {ppc64.ABGT, ppc64.ABLT, true, true}, // GE = GT or EQ; !GE = LT or UN
		ssa.BlockPPC64FLE: {ppc64.ABLT, ppc64.ABGT, true, true}, // LE = LT or EQ; !LE = GT or UN
		ssa.BlockPPC64FGT: {ppc64.ABGT, ppc64.ABLE, false, false},
	}
	return pstate
}

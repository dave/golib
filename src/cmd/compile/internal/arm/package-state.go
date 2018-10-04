package arm

type PackageState struct {
	arm       *arm.PackageState
	gc        *gc.PackageState
	obj       *obj.PackageState
	objabi    *objabi.PackageState
	ssa       *ssa.PackageState
	types     *types.PackageState
	blockJump map[ssa.BlockKind]struct {
		asm    obj.As
		invasm obj.As
	}
	condBits map[ssa.Op]uint8
}

func NewPackageState(gc_pstate *gc.PackageState, obj_pstate *obj.PackageState, arm_pstate *arm.PackageState, ssa_pstate *ssa.PackageState, types_pstate *types.PackageState, objabi_pstate *objabi.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.gc = gc_pstate
	pstate.obj = obj_pstate
	pstate.arm = arm_pstate
	pstate.ssa = ssa_pstate
	pstate.types = types_pstate
	pstate.objabi = objabi_pstate
	pstate.condBits = map[ssa.Op]uint8{
		ssa.OpARMEqual:         arm.C_SCOND_EQ,
		ssa.OpARMNotEqual:      arm.C_SCOND_NE,
		ssa.OpARMLessThan:      arm.C_SCOND_LT,
		ssa.OpARMLessThanU:     arm.C_SCOND_LO,
		ssa.OpARMLessEqual:     arm.C_SCOND_LE,
		ssa.OpARMLessEqualU:    arm.C_SCOND_LS,
		ssa.OpARMGreaterThan:   arm.C_SCOND_GT,
		ssa.OpARMGreaterThanU:  arm.C_SCOND_HI,
		ssa.OpARMGreaterEqual:  arm.C_SCOND_GE,
		ssa.OpARMGreaterEqualU: arm.C_SCOND_HS,
	}
	pstate.blockJump = map[ssa.BlockKind]struct {
		asm, invasm obj.As
	}{
		ssa.BlockARMEQ:  {arm.ABEQ, arm.ABNE},
		ssa.BlockARMNE:  {arm.ABNE, arm.ABEQ},
		ssa.BlockARMLT:  {arm.ABLT, arm.ABGE},
		ssa.BlockARMGE:  {arm.ABGE, arm.ABLT},
		ssa.BlockARMLE:  {arm.ABLE, arm.ABGT},
		ssa.BlockARMGT:  {arm.ABGT, arm.ABLE},
		ssa.BlockARMULT: {arm.ABLO, arm.ABHS},
		ssa.BlockARMUGE: {arm.ABHS, arm.ABLO},
		ssa.BlockARMUGT: {arm.ABHI, arm.ABLS},
		ssa.BlockARMULE: {arm.ABLS, arm.ABHI},
	}
	return pstate
}

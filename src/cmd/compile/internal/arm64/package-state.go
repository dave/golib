package arm64

type PackageState struct {
	arm64     *arm64.PackageState
	gc        *gc.PackageState
	obj       *obj.PackageState
	objabi    *objabi.PackageState
	ssa       *ssa.PackageState
	types     *types.PackageState
	blockJump map[ssa.BlockKind]struct {
		asm    obj.As
		invasm obj.As
	}
	condBits map[ssa.Op]int16
	darwin   bool
}

func NewPackageState(gc_pstate *gc.PackageState, obj_pstate *obj.PackageState, arm64_pstate *arm64.PackageState, objabi_pstate *objabi.PackageState, ssa_pstate *ssa.PackageState, types_pstate *types.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.gc = gc_pstate
	pstate.obj = obj_pstate
	pstate.arm64 = arm64_pstate
	pstate.objabi = objabi_pstate
	pstate.ssa = ssa_pstate
	pstate.types = types_pstate
	pstate.darwin = pstate.objabi.GOOS == "darwin"
	pstate.condBits = map[ssa.Op]int16{
		ssa.OpARM64Equal:         arm64.COND_EQ,
		ssa.OpARM64NotEqual:      arm64.COND_NE,
		ssa.OpARM64LessThan:      arm64.COND_LT,
		ssa.OpARM64LessThanU:     arm64.COND_LO,
		ssa.OpARM64LessEqual:     arm64.COND_LE,
		ssa.OpARM64LessEqualU:    arm64.COND_LS,
		ssa.OpARM64GreaterThan:   arm64.COND_GT,
		ssa.OpARM64GreaterThanU:  arm64.COND_HI,
		ssa.OpARM64GreaterEqual:  arm64.COND_GE,
		ssa.OpARM64GreaterEqualU: arm64.COND_HS,
	}
	pstate.blockJump = map[ssa.BlockKind]struct {
		asm, invasm obj.As
	}{
		ssa.BlockARM64EQ:   {arm64.ABEQ, arm64.ABNE},
		ssa.BlockARM64NE:   {arm64.ABNE, arm64.ABEQ},
		ssa.BlockARM64LT:   {arm64.ABLT, arm64.ABGE},
		ssa.BlockARM64GE:   {arm64.ABGE, arm64.ABLT},
		ssa.BlockARM64LE:   {arm64.ABLE, arm64.ABGT},
		ssa.BlockARM64GT:   {arm64.ABGT, arm64.ABLE},
		ssa.BlockARM64ULT:  {arm64.ABLO, arm64.ABHS},
		ssa.BlockARM64UGE:  {arm64.ABHS, arm64.ABLO},
		ssa.BlockARM64UGT:  {arm64.ABHI, arm64.ABLS},
		ssa.BlockARM64ULE:  {arm64.ABLS, arm64.ABHI},
		ssa.BlockARM64Z:    {arm64.ACBZ, arm64.ACBNZ},
		ssa.BlockARM64NZ:   {arm64.ACBNZ, arm64.ACBZ},
		ssa.BlockARM64ZW:   {arm64.ACBZW, arm64.ACBNZW},
		ssa.BlockARM64NZW:  {arm64.ACBNZW, arm64.ACBZW},
		ssa.BlockARM64TBZ:  {arm64.ATBZ, arm64.ATBNZ},
		ssa.BlockARM64TBNZ: {arm64.ATBNZ, arm64.ATBZ},
	}
	return pstate
}

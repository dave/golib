package s390x

type PackageState struct {
	gc        *gc.PackageState
	obj       *obj.PackageState
	s390x     *s390x.PackageState
	ssa       *ssa.PackageState
	types     *types.PackageState
	blockJump [89]struct {
		asm    obj.As
		invasm obj.As
	}
}

func NewPackageState(gc_pstate *gc.PackageState, s390x_pstate *s390x.PackageState, obj_pstate *obj.PackageState, ssa_pstate *ssa.PackageState, types_pstate *types.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.gc = gc_pstate
	pstate.s390x = s390x_pstate
	pstate.obj = obj_pstate
	pstate.ssa = ssa_pstate
	pstate.types = types_pstate
	pstate.blockJump = [...]struct {
		asm, invasm obj.As
	}{
		ssa.BlockS390XEQ:  {s390x.ABEQ, s390x.ABNE},
		ssa.BlockS390XNE:  {s390x.ABNE, s390x.ABEQ},
		ssa.BlockS390XLT:  {s390x.ABLT, s390x.ABGE},
		ssa.BlockS390XGE:  {s390x.ABGE, s390x.ABLT},
		ssa.BlockS390XLE:  {s390x.ABLE, s390x.ABGT},
		ssa.BlockS390XGT:  {s390x.ABGT, s390x.ABLE},
		ssa.BlockS390XGTF: {s390x.ABGT, s390x.ABLEU},
		ssa.BlockS390XGEF: {s390x.ABGE, s390x.ABLTU},
	}
	return pstate
}

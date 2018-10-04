package mips64

type PackageState struct {
	gc        *gc.PackageState
	mips      *mips.PackageState
	obj       *obj.PackageState
	objabi    *objabi.PackageState
	ssa       *ssa.PackageState
	types     *types.PackageState
	blockJump map[ssa.BlockKind]struct {
		asm    obj.As
		invasm obj.As
	}
}

func NewPackageState(gc_pstate *gc.PackageState, ssa_pstate *ssa.PackageState, mips_pstate *mips.PackageState, objabi_pstate *objabi.PackageState, obj_pstate *obj.PackageState, types_pstate *types.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.gc = gc_pstate
	pstate.ssa = ssa_pstate
	pstate.mips = mips_pstate
	pstate.objabi = objabi_pstate
	pstate.obj = obj_pstate
	pstate.types = types_pstate
	pstate.blockJump = map[ssa.BlockKind]struct {
		asm, invasm obj.As
	}{
		ssa.BlockMIPS64EQ:  {mips.ABEQ, mips.ABNE},
		ssa.BlockMIPS64NE:  {mips.ABNE, mips.ABEQ},
		ssa.BlockMIPS64LTZ: {mips.ABLTZ, mips.ABGEZ},
		ssa.BlockMIPS64GEZ: {mips.ABGEZ, mips.ABLTZ},
		ssa.BlockMIPS64LEZ: {mips.ABLEZ, mips.ABGTZ},
		ssa.BlockMIPS64GTZ: {mips.ABGTZ, mips.ABLEZ},
		ssa.BlockMIPS64FPT: {mips.ABFPT, mips.ABFPF},
		ssa.BlockMIPS64FPF: {mips.ABFPF, mips.ABFPT},
	}
	return pstate
}

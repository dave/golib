package mips

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

func NewPackageState(gc_pstate *gc.PackageState, ssa_pstate *ssa.PackageState, types_pstate *types.PackageState, obj_pstate *obj.PackageState, mips_pstate *mips.PackageState, objabi_pstate *objabi.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.gc = gc_pstate
	pstate.ssa = ssa_pstate
	pstate.types = types_pstate
	pstate.obj = obj_pstate
	pstate.mips = mips_pstate
	pstate.objabi = objabi_pstate
	pstate.blockJump = map[ssa.BlockKind]struct {
		asm, invasm obj.As
	}{
		ssa.BlockMIPSEQ:  {mips.ABEQ, mips.ABNE},
		ssa.BlockMIPSNE:  {mips.ABNE, mips.ABEQ},
		ssa.BlockMIPSLTZ: {mips.ABLTZ, mips.ABGEZ},
		ssa.BlockMIPSGEZ: {mips.ABGEZ, mips.ABLTZ},
		ssa.BlockMIPSLEZ: {mips.ABLEZ, mips.ABGTZ},
		ssa.BlockMIPSGTZ: {mips.ABGTZ, mips.ABLEZ},
		ssa.BlockMIPSFPT: {mips.ABFPT, mips.ABFPF},
		ssa.BlockMIPSFPF: {mips.ABFPF, mips.ABFPT},
	}
	return pstate
}

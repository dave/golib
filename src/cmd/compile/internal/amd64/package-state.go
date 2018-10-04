package amd64

type PackageState struct {
	gc        *gc.PackageState
	obj       *obj.PackageState
	objabi    *objabi.PackageState
	ssa       *ssa.PackageState
	types     *types.PackageState
	x86       *x86.PackageState
	blockJump [29]struct {
		asm    obj.As
		invasm obj.As
	}
	eqfJumps [2][2]gc.FloatingEQNEJump
	isPlan9  bool
	leaptr   obj.As
	nefJumps [2][2]gc.FloatingEQNEJump
}

func NewPackageState(gc_pstate *gc.PackageState, ssa_pstate *ssa.PackageState, types_pstate *types.PackageState, obj_pstate *obj.PackageState, x86_pstate *x86.PackageState, objabi_pstate *objabi.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.gc = gc_pstate
	pstate.ssa = ssa_pstate
	pstate.types = types_pstate
	pstate.obj = obj_pstate
	pstate.x86 = x86_pstate
	pstate.objabi = objabi_pstate
	pstate.blockJump = [...]struct {
		asm, invasm obj.As
	}{
		ssa.BlockAMD64EQ:  {x86.AJEQ, x86.AJNE},
		ssa.BlockAMD64NE:  {x86.AJNE, x86.AJEQ},
		ssa.BlockAMD64LT:  {x86.AJLT, x86.AJGE},
		ssa.BlockAMD64GE:  {x86.AJGE, x86.AJLT},
		ssa.BlockAMD64LE:  {x86.AJLE, x86.AJGT},
		ssa.BlockAMD64GT:  {x86.AJGT, x86.AJLE},
		ssa.BlockAMD64ULT: {x86.AJCS, x86.AJCC},
		ssa.BlockAMD64UGE: {x86.AJCC, x86.AJCS},
		ssa.BlockAMD64UGT: {x86.AJHI, x86.AJLS},
		ssa.BlockAMD64ULE: {x86.AJLS, x86.AJHI},
		ssa.BlockAMD64ORD: {x86.AJPC, x86.AJPS},
		ssa.BlockAMD64NAN: {x86.AJPS, x86.AJPC},
	}
	pstate.eqfJumps = [2][2]gc.FloatingEQNEJump{
		{{Jump: x86.AJNE, Index: 1}, {Jump: x86.AJPS, Index: 1}}, // next == b.Succs[0]
		{{Jump: x86.AJNE, Index: 1}, {Jump: x86.AJPC, Index: 0}}, // next == b.Succs[1]
	}
	pstate.nefJumps = [2][2]gc.FloatingEQNEJump{
		{{Jump: x86.AJNE, Index: 0}, {Jump: x86.AJPC, Index: 1}}, // next == b.Succs[0]
		{{Jump: x86.AJNE, Index: 0}, {Jump: x86.AJPS, Index: 0}}, // next == b.Succs[1]
	}
	pstate.leaptr = x86.ALEAQ
	pstate.isPlan9 = pstate.objabi.GOOS == "plan9"
	return pstate
}

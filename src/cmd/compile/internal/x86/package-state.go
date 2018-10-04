package x86

type PackageState struct {
	gc        *gc.PackageState
	obj       *obj.PackageState
	objabi    *objabi.PackageState
	ssa       *ssa.PackageState
	types     *types.PackageState
	x86       *x86.PackageState
	blockJump [15]struct {
		asm    obj.As
		invasm obj.As
	}
	eqfJumps [2][2]gc.FloatingEQNEJump
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
		ssa.Block386EQ:  {x86.AJEQ, x86.AJNE},
		ssa.Block386NE:  {x86.AJNE, x86.AJEQ},
		ssa.Block386LT:  {x86.AJLT, x86.AJGE},
		ssa.Block386GE:  {x86.AJGE, x86.AJLT},
		ssa.Block386LE:  {x86.AJLE, x86.AJGT},
		ssa.Block386GT:  {x86.AJGT, x86.AJLE},
		ssa.Block386ULT: {x86.AJCS, x86.AJCC},
		ssa.Block386UGE: {x86.AJCC, x86.AJCS},
		ssa.Block386UGT: {x86.AJHI, x86.AJLS},
		ssa.Block386ULE: {x86.AJLS, x86.AJHI},
		ssa.Block386ORD: {x86.AJPC, x86.AJPS},
		ssa.Block386NAN: {x86.AJPS, x86.AJPC},
	}
	pstate.eqfJumps = [2][2]gc.FloatingEQNEJump{
		{{Jump: x86.AJNE, Index: 1}, {Jump: x86.AJPS, Index: 1}}, // next == b.Succs[0]
		{{Jump: x86.AJNE, Index: 1}, {Jump: x86.AJPC, Index: 0}}, // next == b.Succs[1]
	}
	pstate.nefJumps = [2][2]gc.FloatingEQNEJump{
		{{Jump: x86.AJNE, Index: 0}, {Jump: x86.AJPC, Index: 1}}, // next == b.Succs[0]
		{{Jump: x86.AJNE, Index: 0}, {Jump: x86.AJPS, Index: 0}}, // next == b.Succs[1]
	}
	return pstate
}

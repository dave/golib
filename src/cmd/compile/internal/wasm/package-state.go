package wasm

type PackageState struct {
	gc    *gc.PackageState
	obj   *obj.PackageState
	ssa   *ssa.PackageState
	types *types.PackageState
	wasm  *wasm.PackageState
}

func NewPackageState(gc_pstate *gc.PackageState, ssa_pstate *ssa.PackageState, types_pstate *types.PackageState, obj_pstate *obj.PackageState, wasm_pstate *wasm.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.gc = gc_pstate
	pstate.ssa = ssa_pstate
	pstate.types = types_pstate
	pstate.obj = obj_pstate
	pstate.wasm = wasm_pstate
	return pstate
}

package x86

type PackageState struct {
	ld     *ld.PackageState
	objabi *objabi.PackageState
	sym    *sym.PackageState
	sys    *sys.PackageState
}

func NewPackageState(objabi_pstate *objabi.PackageState, sys_pstate *sys.PackageState, ld_pstate *ld.PackageState, sym_pstate *sym.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.objabi = objabi_pstate
	pstate.sys = sys_pstate
	pstate.ld = ld_pstate
	pstate.sym = sym_pstate
	return pstate
}

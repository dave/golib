package loadmacho

type PackageState struct {
	bio    *bio.PackageState
	objabi *objabi.PackageState
	sym    *sym.PackageState
	sys    *sys.PackageState
}

func NewPackageState(bio_pstate *bio.PackageState, objabi_pstate *objabi.PackageState, sys_pstate *sys.PackageState, sym_pstate *sym.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.bio = bio_pstate
	pstate.objabi = objabi_pstate
	pstate.sys = sys_pstate
	pstate.sym = sym_pstate
	return pstate
}

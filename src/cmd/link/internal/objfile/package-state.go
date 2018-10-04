package objfile

type PackageState struct {
	bio      *bio.PackageState
	dwarf    *dwarf.PackageState
	objabi   *objabi.PackageState
	sym      *sym.PackageState
	sys      *sys.PackageState
	emptyPkg []byte
}

func NewPackageState(bio_pstate *bio.PackageState, dwarf_pstate *dwarf.PackageState, objabi_pstate *objabi.PackageState, sys_pstate *sys.PackageState, sym_pstate *sym.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.bio = bio_pstate
	pstate.dwarf = dwarf_pstate
	pstate.objabi = objabi_pstate
	pstate.sys = sys_pstate
	pstate.sym = sym_pstate
	pstate.emptyPkg = []byte("\"\".")
	return pstate
}

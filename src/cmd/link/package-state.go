package main

type PackageState struct {
	amd64  *amd64.PackageState
	arm    *arm.PackageState
	arm64  *arm64.PackageState
	ld     *ld.PackageState
	mips   *mips.PackageState
	mips64 *mips64.PackageState
	objabi *objabi.PackageState
	ppc64  *ppc64.PackageState
	s390x  *s390x.PackageState
	sys    *sys.PackageState
	wasm   *wasm.PackageState
	x86    *x86.PackageState
}

func NewPackageState(objabi_pstate *objabi.PackageState, sys_pstate *sys.PackageState, amd64_pstate *amd64.PackageState, arm_pstate *arm.PackageState, arm64_pstate *arm64.PackageState, ld_pstate *ld.PackageState, mips_pstate *mips.PackageState, mips64_pstate *mips64.PackageState, ppc64_pstate *ppc64.PackageState, s390x_pstate *s390x.PackageState, wasm_pstate *wasm.PackageState, x86_pstate *x86.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.objabi = objabi_pstate
	pstate.sys = sys_pstate
	pstate.amd64 = amd64_pstate
	pstate.arm = arm_pstate
	pstate.arm64 = arm64_pstate
	pstate.ld = ld_pstate
	pstate.mips = mips_pstate
	pstate.mips64 = mips64_pstate
	pstate.ppc64 = ppc64_pstate
	pstate.s390x = s390x_pstate
	pstate.wasm = wasm_pstate
	pstate.x86 = x86_pstate
	return pstate
}

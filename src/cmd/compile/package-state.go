package main

type PackageState struct {
	amd64     *amd64.PackageState
	arm       *arm.PackageState
	arm64     *arm64.PackageState
	gc        *gc.PackageState
	mips      *mips.PackageState
	mips64    *mips64.PackageState
	objabi    *objabi.PackageState
	ppc64     *ppc64.PackageState
	s390x     *s390x.PackageState
	wasm      *wasm.PackageState
	x86       *x86.PackageState
	archInits map[string]func( *gc.Arch)
}

func NewPackageState(amd64_pstate *amd64.PackageState, arm_pstate *arm.PackageState, arm64_pstate *arm64.PackageState, gc_pstate *gc.PackageState, mips_pstate *mips.PackageState, mips64_pstate *mips64.PackageState, ppc64_pstate *ppc64.PackageState, s390x_pstate *s390x.PackageState, wasm_pstate *wasm.PackageState, x86_pstate *x86.PackageState, objabi_pstate *objabi.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.amd64 = amd64_pstate
	pstate.arm = arm_pstate
	pstate.arm64 = arm64_pstate
	pstate.gc = gc_pstate
	pstate.mips = mips_pstate
	pstate.mips64 = mips64_pstate
	pstate.ppc64 = ppc64_pstate
	pstate.s390x = s390x_pstate
	pstate.wasm = wasm_pstate
	pstate.x86 = x86_pstate
	pstate.objabi = objabi_pstate
	pstate.archInits = map[string]func(*gc.Arch){
		"386":      pstate.x86.Init,
		"amd64":    pstate.amd64.Init,
		"amd64p32": pstate.amd64.Init,
		"arm":      pstate.arm.Init,
		"arm64":    pstate.arm64.Init,
		"mips":     pstate.mips.Init,
		"mipsle":   pstate.mips.Init,
		"mips64":   pstate.mips64.Init,
		"mips64le": pstate.mips64.Init,
		"ppc64":    pstate.ppc64.Init,
		"ppc64le":  pstate.ppc64.Init,
		"s390x":    pstate.s390x.Init,
		"wasm":     pstate.wasm.Init,
	}
	return pstate
}

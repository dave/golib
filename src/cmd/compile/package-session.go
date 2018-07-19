package main

import (
	"github.com/dave/golib/src/cmd/compile/internal/amd64"
	"github.com/dave/golib/src/cmd/compile/internal/arm"
	"github.com/dave/golib/src/cmd/compile/internal/arm64"
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/mips"
	"github.com/dave/golib/src/cmd/compile/internal/mips64"
	"github.com/dave/golib/src/cmd/compile/internal/ppc64"
	"github.com/dave/golib/src/cmd/compile/internal/s390x"
	"github.com/dave/golib/src/cmd/compile/internal/wasm"
	"github.com/dave/golib/src/cmd/compile/internal/x86"
	"github.com/dave/golib/src/cmd/internal/objabi"
)

type PackageSession struct {
	amd64  *amd64.PackageSession
	arm    *arm.PackageSession
	arm64  *arm64.PackageSession
	gc     *gc.PackageSession
	mips   *mips.PackageSession
	mips64 *mips64.PackageSession
	objabi *objabi.PackageSession
	ppc64  *ppc64.PackageSession
	s390x  *s390x.PackageSession
	wasm   *wasm.PackageSession
	x86    *x86.PackageSession

	archInits map[string]func(*gc.Arch)
}

func NewPackageSession(amd64_psess *amd64.PackageSession, arm_psess *arm.PackageSession, arm64_psess *arm64.PackageSession, gc_psess *gc.PackageSession, mips_psess *mips.PackageSession, mips64_psess *mips64.PackageSession, ppc64_psess *ppc64.PackageSession, s390x_psess *s390x.PackageSession, wasm_psess *wasm.PackageSession, x86_psess *x86.PackageSession, objabi_psess *objabi.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.amd64 = amd64_psess
	psess.arm = arm_psess
	psess.arm64 = arm64_psess
	psess.gc = gc_psess
	psess.mips = mips_psess
	psess.mips64 = mips64_psess
	psess.ppc64 = ppc64_psess
	psess.s390x = s390x_psess
	psess.wasm = wasm_psess
	psess.x86 = x86_psess
	psess.objabi = objabi_psess
	psess.archInits = map[string]func(*gc.Arch){
		"386":      psess.x86.Init,
		"amd64":    psess.amd64.Init,
		"amd64p32": psess.amd64.Init,
		"arm":      psess.arm.Init,
		"arm64":    psess.arm64.Init,
		"mips":     psess.mips.Init,
		"mipsle":   psess.mips.Init,
		"mips64":   psess.mips64.Init,
		"mips64le": psess.mips64.Init,
		"ppc64":    psess.ppc64.Init,
		"ppc64le":  psess.ppc64.Init,
		"s390x":    psess.s390x.Init,
		"wasm":     psess.wasm.Init,
	}
	return psess
}

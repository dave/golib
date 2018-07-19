package main

import (
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/amd64"
	"github.com/dave/golib/src/cmd/link/internal/arm"
	"github.com/dave/golib/src/cmd/link/internal/arm64"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/mips"
	"github.com/dave/golib/src/cmd/link/internal/mips64"
	"github.com/dave/golib/src/cmd/link/internal/ppc64"
	"github.com/dave/golib/src/cmd/link/internal/s390x"
	"github.com/dave/golib/src/cmd/link/internal/wasm"
	"github.com/dave/golib/src/cmd/link/internal/x86"
)

type PackageSession struct {
	amd64  *amd64.PackageSession
	arm    *arm.PackageSession
	arm64  *arm64.PackageSession
	ld     *ld.PackageSession
	mips   *mips.PackageSession
	mips64 *mips64.PackageSession
	objabi *objabi.PackageSession
	ppc64  *ppc64.PackageSession
	s390x  *s390x.PackageSession
	sys    *sys.PackageSession
	wasm   *wasm.PackageSession
	x86    *x86.PackageSession
}

func NewPackageSession(objabi_psess *objabi.PackageSession, sys_psess *sys.PackageSession, amd64_psess *amd64.PackageSession, arm_psess *arm.PackageSession, arm64_psess *arm64.PackageSession, ld_psess *ld.PackageSession, mips_psess *mips.PackageSession, mips64_psess *mips64.PackageSession, ppc64_psess *ppc64.PackageSession, s390x_psess *s390x.PackageSession, wasm_psess *wasm.PackageSession, x86_psess *x86.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.objabi = objabi_psess
	psess.sys = sys_psess
	psess.amd64 = amd64_psess
	psess.arm = arm_psess
	psess.arm64 = arm64_psess
	psess.ld = ld_psess
	psess.mips = mips_psess
	psess.mips64 = mips64_psess
	psess.ppc64 = ppc64_psess
	psess.s390x = s390x_psess
	psess.wasm = wasm_psess
	psess.x86 = x86_psess
	return psess
}

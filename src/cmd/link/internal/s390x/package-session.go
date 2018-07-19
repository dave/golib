package s390x

import (
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/sym"
)

type PackageSession struct {
	ld     *ld.PackageSession
	objabi *objabi.PackageSession
	sym    *sym.PackageSession
	sys    *sys.PackageSession
}

func NewPackageSession(objabi_psess *objabi.PackageSession, sys_psess *sys.PackageSession, ld_psess *ld.PackageSession, sym_psess *sym.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.objabi = objabi_psess
	psess.sys = sys_psess
	psess.ld = ld_psess
	psess.sym = sym_psess
	return psess
}

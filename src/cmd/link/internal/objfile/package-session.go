package objfile

import (
	"github.com/dave/golib/src/cmd/internal/bio"
	"github.com/dave/golib/src/cmd/internal/dwarf"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/sym"
)

type PackageSession struct {
	bio    *bio.PackageSession
	dwarf  *dwarf.PackageSession
	objabi *objabi.PackageSession
	sym    *sym.PackageSession
	sys    *sys.PackageSession

	emptyPkg []byte
}

func NewPackageSession(bio_psess *bio.PackageSession, dwarf_psess *dwarf.PackageSession, objabi_psess *objabi.PackageSession, sys_psess *sys.PackageSession, sym_psess *sym.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.bio = bio_psess
	psess.dwarf = dwarf_psess
	psess.objabi = objabi_psess
	psess.sys = sys_psess
	psess.sym = sym_psess
	psess.emptyPkg = []byte("\"\".")
	return psess
}

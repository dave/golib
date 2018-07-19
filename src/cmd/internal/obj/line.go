package obj

import (
	"github.com/dave/golib/src/cmd/internal/src"
)

// AddImport adds a package to the list of imported packages.
func (ctxt *Link) AddImport(pkg string) {
	ctxt.Imports = append(ctxt.Imports, pkg)
}

func (psess *PackageSession) linkgetlineFromPos(ctxt *Link, xpos src.XPos) (f string, l int32) {
	pos := ctxt.PosTable.Pos(xpos)
	if !pos.IsKnown() {
		pos = src.Pos{}
	}

	return pos.SymFilename(), int32(pos.RelLine(psess.src))
}

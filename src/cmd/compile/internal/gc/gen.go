package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/src"
	"strconv"
)

func (psess *PackageSession) sysfunc(name string) *obj.LSym {
	return psess.Runtimepkg.Lookup(psess.types, name).Linksym(psess.

		// isParamStackCopy reports whether this is the on-stack copy of a
		// function parameter that moved to the heap.
		types)
}

func (n *Node) isParamStackCopy() bool {
	return n.Op == ONAME && (n.Class() == PPARAM || n.Class() == PPARAMOUT) && n.Name.Param.Heapaddr != nil
}

// isParamHeapCopy reports whether this is the on-heap copy of
// a function parameter that moved to the heap.
func (n *Node) isParamHeapCopy() bool {
	return n.Op == ONAME && n.Class() == PAUTOHEAP && n.Name.Param.Stackcopy != nil
}

// autotmpname returns the name for an autotmp variable numbered n.
func (psess *PackageSession) autotmpname(n int) string {
	// Give each tmp a different name so that they can be registerized.
	// Add a preceding . to avoid clashing with legal names.
	const prefix = ".autotmp_"

	b := []byte(prefix + "      ")[:len(prefix)]
	b = strconv.AppendInt(b, int64(n), 10)
	return psess.types.InternString(b)
}

// make a new Node off the books
func (psess *PackageSession) tempAt(pos src.XPos, curfn *Node, t *types.Type) *Node {
	if curfn == nil {
		psess.
			Fatalf("no curfn for tempname")
	}
	if curfn.Func.Closure != nil && curfn.Op == OCLOSURE {
		Dump("tempname", curfn)
		psess.
			Fatalf("adding tempname to wrong closure function")
	}
	if t == nil {
		psess.
			Fatalf("tempname called with nil type")
	}

	s := &types.Sym{
		Name: psess.autotmpname(len(curfn.Func.Dcl)),
		Pkg:  psess.localpkg,
	}
	n := psess.newnamel(pos, s)
	s.Def = asTypesNode(n)
	n.Type = t
	n.SetClass(PAUTO)
	n.Esc = EscNever
	n.Name.Curfn = curfn
	n.Name.SetUsed(true)
	n.Name.SetAutoTemp(true)
	curfn.Func.Dcl = append(curfn.Func.Dcl, n)
	psess.
		dowidth(t)

	return n.Orig
}

func (psess *PackageSession) temp(t *types.Type) *Node {
	return psess.tempAt(psess.lineno, psess.Curfn, t)
}

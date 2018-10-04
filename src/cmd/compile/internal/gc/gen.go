// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/src"
	"strconv"
)

func (pstate *PackageState) sysfunc(name string) *obj.LSym {
	return pstate.Runtimepkg.Lookup(pstate.types, name).Linksym(pstate.types)
}

// isParamStackCopy reports whether this is the on-stack copy of a
// function parameter that moved to the heap.
func (n *Node) isParamStackCopy() bool {
	return n.Op == ONAME && (n.Class() == PPARAM || n.Class() == PPARAMOUT) && n.Name.Param.Heapaddr != nil
}

// isParamHeapCopy reports whether this is the on-heap copy of
// a function parameter that moved to the heap.
func (n *Node) isParamHeapCopy() bool {
	return n.Op == ONAME && n.Class() == PAUTOHEAP && n.Name.Param.Stackcopy != nil
}

// autotmpname returns the name for an autotmp variable numbered n.
func (pstate *PackageState) autotmpname(n int) string {
	// Give each tmp a different name so that they can be registerized.
	// Add a preceding . to avoid clashing with legal names.
	const prefix = ".autotmp_"
	// Start with a buffer big enough to hold a large n.
	b := []byte(prefix + "      ")[:len(prefix)]
	b = strconv.AppendInt(b, int64(n), 10)
	return pstate.types.InternString(b)
}

// make a new Node off the books
func (pstate *PackageState) tempAt(pos src.XPos, curfn *Node, t *types.Type) *Node {
	if curfn == nil {
		pstate.Fatalf("no curfn for tempname")
	}
	if curfn.Func.Closure != nil && curfn.Op == OCLOSURE {
		Dump("tempname", curfn)
		pstate.Fatalf("adding tempname to wrong closure function")
	}
	if t == nil {
		pstate.Fatalf("tempname called with nil type")
	}

	s := &types.Sym{
		Name: pstate.autotmpname(len(curfn.Func.Dcl)),
		Pkg:  pstate.localpkg,
	}
	n := pstate.newnamel(pos, s)
	s.Def = asTypesNode(n)
	n.Type = t
	n.SetClass(PAUTO)
	n.Esc = EscNever
	n.Name.Curfn = curfn
	n.Name.SetUsed(true)
	n.Name.SetAutoTemp(true)
	curfn.Func.Dcl = append(curfn.Func.Dcl, n)

	pstate.dowidth(t)

	return n.Orig
}

func (pstate *PackageState) temp(t *types.Type) *Node {
	return pstate.tempAt(pstate.lineno, pstate.Curfn, t)
}

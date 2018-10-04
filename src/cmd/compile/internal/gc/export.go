// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/bio"
	"github.com/dave/golib/src/cmd/internal/src"
)

func (pstate *PackageState) exportf(bout *bio.Writer, format string, args ...interface{}) {
	fmt.Fprintf(bout, format, args...)
	if pstate.Debug_export != 0 {
		fmt.Printf(format, args...)
	}
}

// exportsym marks n for export (or reexport).
func (pstate *PackageState) exportsym(n *Node) {
	if n.Sym.OnExportList() {
		return
	}
	n.Sym.SetOnExportList(true)

	if pstate.Debug['E'] != 0 {
		fmt.Printf("export symbol %v\n", n.Sym)
	}

	pstate.exportlist = append(pstate.exportlist, n)
}

func initname(s string) bool {
	return s == "init"
}

func (pstate *PackageState) autoexport(n *Node, ctxt Class) {
	if n.Sym.Pkg != pstate.localpkg {
		return
	}
	if (ctxt != PEXTERN && ctxt != PFUNC) || pstate.dclcontext != PEXTERN {
		return
	}
	if n.Type != nil && n.Type.IsKind(TFUNC) && n.IsMethod(pstate) {
		return
	}

	if types.IsExported(n.Sym.Name) || initname(n.Sym.Name) {
		pstate.exportsym(n)
	}
	if pstate.asmhdr != "" && !n.Sym.Asm() {
		n.Sym.SetAsm(true)
		pstate.asmlist = append(pstate.asmlist, n)
	}
}

// methodbyname sorts types by symbol name.
type methodbyname []*types.Field

func (x methodbyname) Len() int           { return len(x) }
func (x methodbyname) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x methodbyname) Less(i, j int) bool { return x[i].Sym.Name < x[j].Sym.Name }

func (pstate *PackageState) dumpexport(bout *bio.Writer) {
	// The linker also looks for the $$ marker - use char after $$ to distinguish format.
	pstate.exportf(bout, "\n$$B\n") // indicate binary export format
	off := bout.Offset()
	if pstate.flagiexport {
		pstate.iexport(bout.Writer)
	} else {
		pstate.export(bout.Writer, pstate.Debug_export != 0)
	}
	size := bout.Offset() - off
	pstate.exportf(bout, "\n$$\n")

	if pstate.Debug_export != 0 {
		fmt.Printf("export data size = %d bytes\n", size)
	}
}

func (pstate *PackageState) importsym(ipkg *types.Pkg, pos src.XPos, s *types.Sym, op Op) *Node {
	n := asNode(s.Def)
	if n == nil {
		// iimport should have created a stub ONONAME
		// declaration for all imported symbols. The exception
		// is declarations for Runtimepkg, which are populated
		// by loadsys instead.
		if pstate.flagiexport && s.Pkg != pstate.Runtimepkg {
			pstate.Fatalf("missing ONONAME for %v\n", s)
		}

		n = pstate.dclname(s)
		s.Def = asTypesNode(n)
		s.Importdef = ipkg
	}
	if n.Op != ONONAME && n.Op != op {
		pstate.redeclare(pstate.lineno, s, fmt.Sprintf("during import %q", ipkg.Path))
	}
	return n
}

// pkgtype returns the named type declared by symbol s.
// If no such type has been declared yet, a forward declaration is returned.
// ipkg is the package being imported
func (pstate *PackageState) importtype(ipkg *types.Pkg, pos src.XPos, s *types.Sym) *types.Type {
	n := pstate.importsym(ipkg, pos, s, OTYPE)
	if n.Op != OTYPE {
		t := types.New(TFORW)
		t.Sym = s
		t.Nod = asTypesNode(n)

		n.Op = OTYPE
		n.Pos = pos
		n.Type = t
		n.SetClass(PEXTERN)
	}

	t := n.Type
	if t == nil {
		pstate.Fatalf("importtype %v", s)
	}
	return t
}

// importobj declares symbol s as an imported object representable by op.
// ipkg is the package being imported
func (pstate *PackageState) importobj(ipkg *types.Pkg, pos src.XPos, s *types.Sym, op Op, ctxt Class, t *types.Type) *Node {
	n := pstate.importsym(ipkg, pos, s, op)
	if n.Op != ONONAME {
		if n.Op == op && (n.Class() != ctxt || !pstate.eqtype(n.Type, t)) {
			pstate.redeclare(pstate.lineno, s, fmt.Sprintf("during import %q", ipkg.Path))
		}
		return nil
	}

	n.Op = op
	n.Pos = pos
	n.SetClass(ctxt)
	n.Type = t
	return n
}

// importconst declares symbol s as an imported constant with type t and value val.
// ipkg is the package being imported
func (pstate *PackageState) importconst(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type, val Val) {
	n := pstate.importobj(ipkg, pos, s, OLITERAL, PEXTERN, t)
	if n == nil { // TODO: Check that value matches.
		return
	}

	n.SetVal(pstate, val)

	if pstate.Debug['E'] != 0 {
		fmt.Printf("import const %v %L = %v\n", s, t, val)
	}
}

// importfunc declares symbol s as an imported function with type t.
// ipkg is the package being imported
func (pstate *PackageState) importfunc(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type) {
	n := pstate.importobj(ipkg, pos, s, ONAME, PFUNC, t)
	if n == nil {
		return
	}

	n.Func = new(Func)
	t.SetNname(pstate.types, asTypesNode(n))

	if pstate.Debug['E'] != 0 {
		fmt.Printf("import func %v%S\n", s, t)
	}
}

// importvar declares symbol s as an imported variable with type t.
// ipkg is the package being imported
func (pstate *PackageState) importvar(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type) {
	n := pstate.importobj(ipkg, pos, s, ONAME, PEXTERN, t)
	if n == nil {
		return
	}

	if pstate.Debug['E'] != 0 {
		fmt.Printf("import var %v %L\n", s, t)
	}
}

// importalias declares symbol s as an imported type alias with type t.
// ipkg is the package being imported
func (pstate *PackageState) importalias(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type) {
	n := pstate.importobj(ipkg, pos, s, OTYPE, PEXTERN, t)
	if n == nil {
		return
	}

	if pstate.Debug['E'] != 0 {
		fmt.Printf("import type %v = %L\n", s, t)
	}
}

func (pstate *PackageState) dumpasmhdr() {
	b, err := bio.Create(pstate.asmhdr)
	if err != nil {
		pstate.Fatalf("%v", err)
	}
	fmt.Fprintf(b, "// generated by compile -asmhdr from package %s\n\n", pstate.localpkg.Name)
	for _, n := range pstate.asmlist {
		if n.Sym.IsBlank() {
			continue
		}
		switch n.Op {
		case OLITERAL:
			fmt.Fprintf(b, "#define const_%s %#v\n", n.Sym.Name, n.Val())

		case OTYPE:
			t := n.Type
			if !t.IsStruct() || t.StructType(pstate.types).Map != nil || t.IsFuncArgStruct() {
				break
			}
			fmt.Fprintf(b, "#define %s__size %d\n", n.Sym.Name, int(t.Width))
			for _, f := range t.Fields(pstate.types).Slice() {
				if !f.Sym.IsBlank() {
					fmt.Fprintf(b, "#define %s_%s %d\n", n.Sym.Name, f.Sym.Name, int(f.Offset))
				}
			}
		}
	}

	b.Close()
}

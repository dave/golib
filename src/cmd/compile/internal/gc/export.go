package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/bio"
	"github.com/dave/golib/src/cmd/internal/src"
)

// if set, use indexed export data format

// if set, print debugging information about export data

func (psess *PackageSession) exportf(bout *bio.Writer, format string, args ...interface{}) {
	fmt.Fprintf(bout, format, args...)
	if psess.Debug_export != 0 {
		fmt.Printf(format, args...)
	}
}

// exportsym marks n for export (or reexport).
func (psess *PackageSession) exportsym(n *Node) {
	if n.Sym.OnExportList() {
		return
	}
	n.Sym.SetOnExportList(true)

	if psess.Debug['E'] != 0 {
		fmt.Printf("export symbol %v\n", n.Sym)
	}
	psess.
		exportlist = append(psess.exportlist, n)
}

func initname(s string) bool {
	return s == "init"
}

func (psess *PackageSession) autoexport(n *Node, ctxt Class) {
	if n.Sym.Pkg != psess.localpkg {
		return
	}
	if (ctxt != PEXTERN && ctxt != PFUNC) || psess.dclcontext != PEXTERN {
		return
	}
	if n.Type != nil && n.Type.IsKind(TFUNC) && n.IsMethod(psess) {
		return
	}

	if types.IsExported(n.Sym.Name) || initname(n.Sym.Name) {
		psess.
			exportsym(n)
	}
	if psess.asmhdr != "" && !n.Sym.Asm() {
		n.Sym.SetAsm(true)
		psess.
			asmlist = append(psess.asmlist, n)
	}
}

// methodbyname sorts types by symbol name.
type methodbyname []*types.Field

func (x methodbyname) Len() int           { return len(x) }
func (x methodbyname) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x methodbyname) Less(i, j int) bool { return x[i].Sym.Name < x[j].Sym.Name }

func (psess *PackageSession) dumpexport(bout *bio.Writer) {
	psess.
		exportf(bout, "\n$$B\n")
	off := bout.Offset()
	if psess.flagiexport {
		psess.
			iexport(bout.Writer)
	} else {
		psess.
			export(bout.Writer, psess.Debug_export != 0)
	}
	size := bout.Offset() - off
	psess.
		exportf(bout, "\n$$\n")

	if psess.Debug_export != 0 {
		fmt.Printf("export data size = %d bytes\n", size)
	}
}

func (psess *PackageSession) importsym(ipkg *types.Pkg, pos src.XPos, s *types.Sym, op Op) *Node {
	n := asNode(s.Def)
	if n == nil {

		if psess.flagiexport && s.Pkg != psess.Runtimepkg {
			psess.
				Fatalf("missing ONONAME for %v\n", s)
		}

		n = psess.dclname(s)
		s.Def = asTypesNode(n)
		s.Importdef = ipkg
	}
	if n.Op != ONONAME && n.Op != op {
		psess.
			redeclare(psess.lineno, s, fmt.Sprintf("during import %q", ipkg.Path))
	}
	return n
}

// pkgtype returns the named type declared by symbol s.
// If no such type has been declared yet, a forward declaration is returned.
// ipkg is the package being imported
func (psess *PackageSession) importtype(ipkg *types.Pkg, pos src.XPos, s *types.Sym) *types.Type {
	n := psess.importsym(ipkg, pos, s, OTYPE)
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
		psess.
			Fatalf("importtype %v", s)
	}
	return t
}

// importobj declares symbol s as an imported object representable by op.
// ipkg is the package being imported
func (psess *PackageSession) importobj(ipkg *types.Pkg, pos src.XPos, s *types.Sym, op Op, ctxt Class, t *types.Type) *Node {
	n := psess.importsym(ipkg, pos, s, op)
	if n.Op != ONONAME {
		if n.Op == op && (n.Class() != ctxt || !psess.eqtype(n.Type, t)) {
			psess.
				redeclare(psess.lineno, s, fmt.Sprintf("during import %q", ipkg.Path))
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
func (psess *PackageSession) importconst(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type, val Val) {
	n := psess.importobj(ipkg, pos, s, OLITERAL, PEXTERN, t)
	if n == nil {
		return
	}

	n.SetVal(psess, val)

	if psess.Debug['E'] != 0 {
		fmt.Printf("import const %v %L = %v\n", s, t, val)
	}
}

// importfunc declares symbol s as an imported function with type t.
// ipkg is the package being imported
func (psess *PackageSession) importfunc(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type) {
	n := psess.importobj(ipkg, pos, s, ONAME, PFUNC, t)
	if n == nil {
		return
	}

	n.Func = new(Func)
	t.SetNname(psess.types, asTypesNode(n))

	if psess.Debug['E'] != 0 {
		fmt.Printf("import func %v%S\n", s, t)
	}
}

// importvar declares symbol s as an imported variable with type t.
// ipkg is the package being imported
func (psess *PackageSession) importvar(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type) {
	n := psess.importobj(ipkg, pos, s, ONAME, PEXTERN, t)
	if n == nil {
		return
	}

	if psess.Debug['E'] != 0 {
		fmt.Printf("import var %v %L\n", s, t)
	}
}

// importalias declares symbol s as an imported type alias with type t.
// ipkg is the package being imported
func (psess *PackageSession) importalias(ipkg *types.Pkg, pos src.XPos, s *types.Sym, t *types.Type) {
	n := psess.importobj(ipkg, pos, s, OTYPE, PEXTERN, t)
	if n == nil {
		return
	}

	if psess.Debug['E'] != 0 {
		fmt.Printf("import type %v = %L\n", s, t)
	}
}

func (psess *PackageSession) dumpasmhdr() {
	b, err := bio.Create(psess.asmhdr)
	if err != nil {
		psess.
			Fatalf("%v", err)
	}
	fmt.Fprintf(b, "// generated by compile -asmhdr from package %s\n\n", psess.localpkg.Name)
	for _, n := range psess.asmlist {
		if n.Sym.IsBlank() {
			continue
		}
		switch n.Op {
		case OLITERAL:
			fmt.Fprintf(b, "#define const_%s %#v\n", n.Sym.Name, n.Val())

		case OTYPE:
			t := n.Type
			if !t.IsStruct() || t.StructType(psess.types).Map != nil || t.IsFuncArgStruct() {
				break
			}
			fmt.Fprintf(b, "#define %s__size %d\n", n.Sym.Name, int(t.Width))
			for _, f := range t.Fields(psess.types).Slice() {
				if !f.Sym.IsBlank() {
					fmt.Fprintf(b, "#define %s_%s %d\n", n.Sym.Name, f.Sym.Name, int(f.Offset))
				}
			}
		}
	}

	b.Close()
}

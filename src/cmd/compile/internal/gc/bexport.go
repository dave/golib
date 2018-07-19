package gc

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/src"
	"math/big"
	"sort"
	"strings"
)

// If debugFormat is set, each integer and string value is preceded by a marker
// and position information in the encoding. This mechanism permits an importer
// to recognize immediately when it is out of sync. The importer recognizes this
// mode automatically (i.e., it can import export data produced with debugging
// support even if debugFormat is not set at the time of import). This mode will
// lead to massively larger export data (by a factor of 2 to 3) and should only
// be enabled during development and debugging.
//
// NOTE: This flag is the first flag to enable if importing dies because of
// (suspected) format errors, and whenever a change is made to the format.
const debugFormat = false // default: false

// Current export format version. Increase with each format change.
// 6: package height (CL 105038)
// 5: improved position encoding efficiency (issue 20080, CL 41619)
// 4: type name objects support type aliases, uses aliasTag
// 3: Go1.8 encoding (same as version 2, aliasTag defined but never used)
// 2: removed unused bool in ODCL export (compiler only)
// 1: header format change (more regular), export package for _ struct fields
// 0: Go1.7 encoding
const exportVersion = 6

// exportInlined enables the export of inlined function bodies and related
// dependencies. The compiler should work w/o any loss of functionality with
// the flag disabled, but the generated code will lose access to inlined
// function bodies across packages, leading to performance bugs.
// Leave for debugging.
const exportInlined = true // default: true

// trackAllTypes enables cycle tracking for all types, not just named
// types. The existing compiler invariants assume that unnamed types
// that are not completely set up are not used, or else there are spurious
// errors.
// If disabled, only named types are tracked, possibly leading to slightly
// less efficient encoding in rare cases. It also prevents the export of
// some corner-case type declarations (but those were not handled correctly
// with the former textual export format either).
// Note that when a type is only seen once, as many unnamed types are,
// it is less efficient to track it, since we then also record an index for it.
// See CLs 41622 and 41623 for some data and discussion.
// TODO(gri) enable selectively and remove once issues caused by it are fixed
const trackAllTypes = false

type exporter struct {
	out *bufio.Writer

	// object -> index maps, indexed in order of serialization
	strIndex  map[string]int
	pathIndex map[string]int
	pkgIndex  map[*types.Pkg]int
	typIndex  map[*types.Type]int
	funcList  []*Func

	marked map[*types.Type]bool // types already seen by markType

	// position encoding
	posInfoFormat bool
	prevFile      string
	prevLine      int

	// debugging support
	written int // bytes written
	indent  int // for p.trace
	trace   bool
}

// export writes the exportlist for localpkg to out and returns the number of bytes written.
func (psess *PackageSession) export(out *bufio.Writer, trace bool) int {
	p := exporter{
		out:           out,
		strIndex:      map[string]int{"": 0},
		pathIndex:     map[string]int{"": 0},
		pkgIndex:      make(map[*types.Pkg]int),
		typIndex:      make(map[*types.Type]int),
		posInfoFormat: true,
		trace:         trace,
	}

	p.rawStringln(fmt.Sprintf("version %d", exportVersion))
	var debug string
	if debugFormat {
		debug = "debug"
	}
	p.rawStringln(debug)
	p.bool(trackAllTypes)
	p.bool(p.posInfoFormat)

	predecl := psess.predeclared()
	for index, typ := range predecl {
		p.typIndex[typ] = index
	}
	if len(p.typIndex) != len(predecl) {
		psess.
			Fatalf("exporter: duplicate entries in type map?")
	}

	if psess.localpkg.Path != "" {
		psess.
			Fatalf("exporter: local package path not empty: %q", psess.localpkg.Path)
	}
	p.pkg(psess, psess.localpkg)
	if p.trace {
		p.tracef("\n")
	}

	numglobals := len(psess.exportlist)

	if exportInlined {
		p.marked = make(map[*types.Type]bool)
		for _, n := range psess.exportlist {
			sym := n.Sym
			p.markType(psess, asNode(sym.Def).Type)
		}
		p.marked = nil
	}

	objcount := 0
	for _, n := range psess.exportlist[:numglobals] {
		sym := n.Sym

		if strings.Contains(sym.Name, ".") {
			psess.
				Fatalf("exporter: unexpected symbol: %v", sym)
		}

		if sym.Def == nil {
			psess.
				Fatalf("exporter: unknown export symbol: %v", sym)
		}

		if p.trace {
			p.tracef("\n")
		}
		p.obj(psess, sym)
		objcount++
	}

	if p.trace {
		p.tracef("\n")
	}
	p.tag(psess, endTag)

	p.int(objcount)

	if p.trace {
		p.tracef("\n--- compiler-specific export data ---\n[ ")
		if p.indent != 0 {
			psess.
				Fatalf("exporter: incorrect indentation")
		}
	}

	if p.trace {
		p.tracef("\n")
	}

	objcount = 0
	for _, n := range psess.exportlist[numglobals:] {
		sym := n.Sym

		if strings.Contains(sym.Name, ".") {
			psess.
				Fatalf("exporter: unexpected symbol: %v", sym)
		}

		if sym.Def == nil {
			psess.
				Fatalf("exporter: unknown export symbol: %v", sym)
		}

		if p.trace {
			p.tracef("\n")
		}

		if IsAlias(sym) {
			psess.
				Fatalf("exporter: unexpected type alias %v in inlined function body", sym)
		}

		p.obj(psess, sym)
		objcount++
	}

	if p.trace {
		p.tracef("\n")
	}
	p.tag(psess, endTag)

	p.int(objcount)

	if p.trace {
		p.tracef("\n--- inlined function bodies ---\n")
		if p.indent != 0 {
			psess.
				Fatalf("exporter: incorrect indentation")
		}
	}

	objcount = 0
	for i := 0; i < len(p.funcList); i++ {
		if f := p.funcList[i]; f.ExportInline() {

			if p.trace {
				p.tracef("\n----\nfunc { %#v }\n", asNodes(f.Inl.Body))
			}
			p.int(i)
			p.int(int(f.Inl.Cost))
			p.stmtList(psess, asNodes(f.Inl.Body))
			if p.trace {
				p.tracef("\n")
			}
			objcount++
		}
	}

	if p.trace {
		p.tracef("\n")
	}
	p.int(-1)

	p.int(objcount)

	if p.trace {
		p.tracef("\n--- end ---\n")
	}

	return p.written
}

func (p *exporter) pkg(psess *PackageSession, pkg *types.Pkg) {
	if pkg == nil {
		psess.
			Fatalf("exporter: unexpected nil pkg")
	}

	if i, ok := p.pkgIndex[pkg]; ok {
		p.index(psess, 'P', i)
		return
	}

	if p.trace {
		p.tracef("P%d = { ", len(p.pkgIndex))
		defer p.tracef("} ")
	}
	p.pkgIndex[pkg] = len(p.pkgIndex)

	p.tag(psess, packageTag)
	p.string(pkg.Name)
	p.path(pkg.Path)
	p.int(pkg.Height)
}

func (psess *PackageSession) unidealType(typ *types.Type, val Val) *types.Type {

	if typ == nil || typ.IsUntyped(psess.types) {
		typ = psess.untype(val.Ctype(psess))
	}
	return typ
}

// markType recursively visits types reachable from t to identify
// functions whose inline bodies may be needed.
func (p *exporter) markType(psess *PackageSession, t *types.Type) {
	if p.marked[t] {
		return
	}
	p.marked[t] = true

	if t.Sym != nil && t.Etype != TINTER {
		for _, m := range t.Methods().Slice() {
			if types.IsExported(m.Sym.Name) {
				p.markType(psess, m.Type)
			}
		}
	}

	switch t.Etype {
	case TPTR32, TPTR64, TARRAY, TSLICE, TCHAN, TMAP:
		p.markType(psess, t.Elem(psess.types))

	case TSTRUCT:
		for _, f := range t.FieldSlice(psess.types) {
			if types.IsExported(f.Sym.Name) || f.Embedded != 0 {
				p.markType(psess, f.Type)
			}
		}

	case TFUNC:
		psess.
			inlFlood(asNode(t.Nname(psess.types)))

		for _, f := range t.Results(psess.types).FieldSlice(psess.types) {
			p.markType(psess, f.Type)
		}

	case TINTER:
		for _, f := range t.FieldSlice(psess.types) {
			if types.IsExported(f.Sym.Name) {
				p.markType(psess, f.Type)
			}
		}
	}
}

func (p *exporter) obj(psess *PackageSession, sym *types.Sym) {

	n := asNode(sym.Def)
	switch n.Op {
	case OLITERAL:

		n = psess.typecheck(n, Erv)
		if n == nil || n.Op != OLITERAL {
			psess.
				Fatalf("exporter: dumpexportconst: oconst nil: %v", sym)
		}

		p.tag(psess, constTag)
		p.pos(psess, n.Pos)

		p.qualifiedName(psess, sym)
		p.typ(psess, psess.unidealType(n.Type, n.Val()))
		p.value(psess, n.Val())

	case OTYPE:

		t := n.Type
		if t.Etype == TFORW {
			psess.
				Fatalf("exporter: export of incomplete type %v", sym)
		}

		if IsAlias(sym) {
			p.tag(psess, aliasTag)
			p.pos(psess, n.Pos)
			p.qualifiedName(psess, sym)
		} else {
			p.tag(psess, typeTag)
		}
		p.typ(psess, t)

	case ONAME:

		n = psess.typecheck(n, Erv|Ecall)
		if n == nil || n.Type == nil {
			psess.
				Fatalf("exporter: variable/function exported but not defined: %v", sym)
		}

		if n.Type.Etype == TFUNC && n.Class() == PFUNC {

			p.tag(psess, funcTag)
			p.pos(psess, n.Pos)
			p.qualifiedName(psess, sym)

			sig := asNode(sym.Def).Type

			p.paramList(psess, sig.Params(psess.types), true)
			p.paramList(psess, sig.Results(psess.types), true)

			p.funcList = append(p.funcList, asNode(sym.Def).Func)
		} else {

			p.tag(psess, varTag)
			p.pos(psess, n.Pos)
			p.qualifiedName(psess, sym)
			p.typ(psess, asNode(sym.Def).Type)
		}

	default:
		psess.
			Fatalf("exporter: unexpected export symbol: %v %v", n.Op, sym)
	}
}

// deltaNewFile is a magic line delta offset indicating a new file.
// We use -64 because it is rare; see issue 20080 and CL 41619.
// -64 is the smallest int that fits in a single byte as a varint.
const deltaNewFile = -64

func (p *exporter) pos(psess *PackageSession, pos src.XPos) {
	if !p.posInfoFormat {
		return
	}

	file, line := psess.fileLine(pos)
	if file == p.prevFile {

		delta := line - p.prevLine
		p.int(delta)
		if delta == deltaNewFile {
			p.int(-1)
		}
	} else {

		p.int(deltaNewFile)
		p.int(line)
		p.path(file)
		p.prevFile = file
	}
	p.prevLine = line
}

func (p *exporter) path(s string) {
	if i, ok := p.pathIndex[s]; ok {

		p.int(i)
		return
	}
	p.pathIndex[s] = len(p.pathIndex)
	c := strings.Split(s, "/")
	p.int(-len(c))
	for _, x := range c {
		p.string(x)
	}
}

func (psess *PackageSession) fileLine(pos0 src.XPos) (file string, line int) {
	pos := psess.Ctxt.PosTable.Pos(pos0)
	file = pos.Base().AbsFilename()
	line = int(pos.RelLine(psess.src))
	return
}

func (p *exporter) typ(psess *PackageSession, t *types.Type) {
	if t == nil {
		psess.
			Fatalf("exporter: nil type")
	}

	if i, ok := p.typIndex[t]; ok {
		p.index(psess, 'T', i)
		return
	}

	if trackAllTypes {
		if p.trace {
			p.tracef("T%d = {>\n", len(p.typIndex))
			defer p.tracef("<\n} ")
		}
		p.typIndex[t] = len(p.typIndex)
	}

	if tsym := t.Sym; tsym != nil {
		if !trackAllTypes {

			p.typIndex[t] = len(p.typIndex)
		}

		if t.Orig == t {
			psess.
				Fatalf("exporter: predeclared type missing from type map?")
		}

		n := psess.typenod(t)
		if n.Type != t {
			psess.
				Fatalf("exporter: named type definition incorrectly set up")
		}

		p.tag(psess, namedTag)
		p.pos(psess, n.Pos)
		p.qualifiedName(psess, tsym)

		p.typ(psess, t.Orig)

		if t.Orig.IsInterface() {
			return
		}

		// sort methods for reproducible export format
		// TODO(gri) Determine if they are already sorted
		// in which case we can drop this step.
		var methods []*types.Field
		methods = append(methods, t.Methods().Slice()...)
		sort.Sort(methodbyname(methods))
		p.int(len(methods))

		if p.trace && len(methods) > 0 {
			p.tracef("associated methods {>")
		}

		for _, m := range methods {
			if p.trace {
				p.tracef("\n")
			}
			if strings.Contains(m.Sym.Name, ".") {
				psess.
					Fatalf("invalid symbol name: %s (%v)", m.Sym.Name, m.Sym)
			}

			p.pos(psess, m.Pos)
			p.fieldSym(psess, m.Sym, false)

			sig := m.Type
			mfn := asNode(sig.FuncType(psess.types).Nname)

			p.paramList(psess, sig.Recvs(psess.types), true)
			p.paramList(psess, sig.Params(psess.types), true)
			p.paramList(psess, sig.Results(psess.types), true)
			p.bool(m.Nointerface())

			p.funcList = append(p.funcList, mfn.Func)
		}

		if p.trace && len(methods) > 0 {
			p.tracef("<\n} ")
		}

		return
	}

	switch t.Etype {
	case TARRAY:
		if t.IsDDDArray() {
			psess.
				Fatalf("array bounds should be known at export time: %v", t)
		}
		p.tag(psess, arrayTag)
		p.int64(t.NumElem(psess.types))
		p.typ(psess, t.Elem(psess.types))

	case TSLICE:
		p.tag(psess, sliceTag)
		p.typ(psess, t.Elem(psess.types))

	case TDDDFIELD:

		p.tag(psess, dddTag)
		p.typ(psess, t.DDDField(psess.types))

	case TSTRUCT:
		p.tag(psess, structTag)
		p.fieldList(psess, t)

	case TPTR32, TPTR64:
		p.tag(psess, pointerTag)
		p.typ(psess, t.Elem(psess.types))

	case TFUNC:
		p.tag(psess, signatureTag)
		p.paramList(psess, t.Params(psess.types), false)
		p.paramList(psess, t.Results(psess.types), false)

	case TINTER:
		p.tag(psess, interfaceTag)
		p.methodList(psess, t)

	case TMAP:
		p.tag(psess, mapTag)
		p.typ(psess, t.Key(psess.types))
		p.typ(psess, t.Elem(psess.types))

	case TCHAN:
		p.tag(psess, chanTag)
		p.int(int(t.ChanDir(psess.types)))
		p.typ(psess, t.Elem(psess.types))

	default:
		psess.
			Fatalf("exporter: unexpected type: %v (Etype = %d)", t, t.Etype)
	}
}

func (p *exporter) qualifiedName(psess *PackageSession, sym *types.Sym) {
	p.string(sym.Name)
	p.pkg(psess, sym.Pkg)
}

func (p *exporter) fieldList(psess *PackageSession, t *types.Type) {
	if p.trace && t.NumFields(psess.types) > 0 {
		p.tracef("fields {>")
		defer p.tracef("<\n} ")
	}

	p.int(t.NumFields(psess.types))
	for _, f := range t.Fields(psess.types).Slice() {
		if p.trace {
			p.tracef("\n")
		}
		p.field(psess, f)
	}
}

func (p *exporter) field(psess *PackageSession, f *types.Field) {
	p.pos(psess, f.Pos)
	p.fieldName(psess, f)
	p.typ(psess, f.Type)
	p.string(f.Note)
}

func (p *exporter) methodList(psess *PackageSession, t *types.Type) {
	var embeddeds, methods []*types.Field

	for _, m := range t.Methods().Slice() {
		if m.Sym != nil {
			methods = append(methods, m)
		} else {
			embeddeds = append(embeddeds, m)
		}
	}

	if p.trace && len(embeddeds) > 0 {
		p.tracef("embeddeds {>")
	}
	p.int(len(embeddeds))
	for _, m := range embeddeds {
		if p.trace {
			p.tracef("\n")
		}
		p.pos(psess, m.Pos)
		p.typ(psess, m.Type)
	}
	if p.trace && len(embeddeds) > 0 {
		p.tracef("<\n} ")
	}

	if p.trace && len(methods) > 0 {
		p.tracef("methods {>")
	}
	p.int(len(methods))
	for _, m := range methods {
		if p.trace {
			p.tracef("\n")
		}
		p.method(psess, m)
	}
	if p.trace && len(methods) > 0 {
		p.tracef("<\n} ")
	}
}

func (p *exporter) method(psess *PackageSession, m *types.Field) {
	p.pos(psess, m.Pos)
	p.methodName(psess, m.Sym)
	p.paramList(psess, m.Type.Params(psess.types), false)
	p.paramList(psess, m.Type.Results(psess.types), false)
}

func (p *exporter) fieldName(psess *PackageSession, t *types.Field) {
	name := t.Sym.Name
	if t.Embedded != 0 {

		bname := psess.basetypeName(t.Type)
		if name == bname {
			if types.IsExported(name) {
				name = ""
			} else {
				name = "?"
			}
		} else {

			p.string("@")
		}
	}
	p.string(name)
	if name != "" && !types.IsExported(name) {
		p.pkg(psess, t.Sym.Pkg)
	}
}

// methodName is like qualifiedName but it doesn't record the package for exported names.
func (p *exporter) methodName(psess *PackageSession, sym *types.Sym) {
	p.string(sym.Name)
	if !types.IsExported(sym.Name) {
		p.pkg(psess, sym.Pkg)
	}
}

func (psess *PackageSession) basetypeName(t *types.Type) string {
	s := t.Sym
	if s == nil && t.IsPtr() {
		s = t.Elem(psess.types).Sym
	}
	if s != nil {
		return s.Name
	}
	return ""
}

func (p *exporter) paramList(psess *PackageSession, params *types.Type, numbered bool) {
	if !params.IsFuncArgStruct() {
		psess.
			Fatalf("exporter: parameter list expected")
	}

	n := params.NumFields(psess.types)
	if n > 0 && psess.parName(params.Field(psess.types, 0), numbered) == "" {
		n = -n
	}
	p.int(n)
	for _, q := range params.Fields(psess.types).Slice() {
		p.param(psess, q, n, numbered)
	}
}

func (p *exporter) param(psess *PackageSession, q *types.Field, n int, numbered bool) {
	t := q.Type
	if q.Isddd() {

		t = types.NewDDDField(t.Elem(psess.types))
	}
	p.typ(psess, t)
	if n > 0 {
		name := psess.parName(q, numbered)
		if name == "" {

			name = "_"
		}
		p.string(name)
		if name != "_" {

			p.pkg(psess, q.Sym.Pkg)
		}
	}

	p.string(q.Note)
}

func (psess *PackageSession) parName(f *types.Field, numbered bool) string {
	s := psess.origSym(f.Sym)
	if s == nil {
		return ""
	}

	if s.Name == "_" {
		return "_"
	}

	name := s.Name
	if strings.Contains(name, ".") {
		psess.
			Fatalf("invalid symbol name: %s", name)
	}

	if numbered {
		if n := asNode(f.Nname); !strings.Contains(name, "·") && n != nil && n.Name.Vargen > 0 {
			name = fmt.Sprintf("%s·%d", name, n.Name.Vargen)
		}
	} else {
		if i := strings.Index(name, "·"); i > 0 {
			name = name[:i]
		}
	}
	return name
}

func (p *exporter) value(psess *PackageSession, x Val) {
	if p.trace {
		p.tracef("= ")
	}

	switch x := x.U.(type) {
	case bool:
		tag := falseTag
		if x {
			tag = trueTag
		}
		p.tag(psess, tag)

	case *Mpint:
		if psess.minintval[TINT64].Cmp(x) <= 0 && x.Cmp(psess.maxintval[TINT64]) <= 0 {

			p.tag(psess, int64Tag)
			p.int64(x.Int64(psess))
			return
		}

		f := newMpflt()
		f.SetInt(x)
		p.tag(psess, floatTag)
		p.float(psess, f)

	case *Mpflt:
		p.tag(psess, floatTag)
		p.float(psess, x)

	case *Mpcplx:
		p.tag(psess, complexTag)
		p.float(psess, &x.Real)
		p.float(psess, &x.Imag)

	case string:
		p.tag(psess, stringTag)
		p.string(x)

	case *NilVal:

		p.tag(psess, nilTag)

	default:
		psess.
			Fatalf("exporter: unexpected value %v (%T)", x, x)
	}
}

func (p *exporter) float(psess *PackageSession, x *Mpflt) {

	f := &x.Val
	sign := f.Sign()
	if sign == 0 {

		p.int(0)
		return
	}

	// extract exponent such that 0.5 <= m < 1.0
	var m big.Float
	exp := f.MantExp(&m)

	m.SetMantExp(&m, int(m.MinPrec()))
	mant, acc := m.Int(nil)
	if acc != big.Exact {
		psess.
			Fatalf("exporter: internal error")
	}

	p.int(sign)
	p.int(exp)
	p.string(string(mant.Bytes()))
}

// stmtList may emit more (or fewer) than len(list) nodes.
func (p *exporter) stmtList(psess *PackageSession, list Nodes) {
	if p.trace {
		if list.Len() == 0 {
			p.tracef("{}")
		} else {
			p.tracef("{>")
			defer p.tracef("<\n}")
		}
	}

	for _, n := range list.Slice() {
		if p.trace {
			p.tracef("\n")
		}

		p.node(psess, n)
	}

	p.op(OEND)
}

func (p *exporter) node(psess *PackageSession, n *Node) {
	if psess.opprec[n.Op] < 0 {
		p.stmt(psess, n)
	} else {
		p.expr(psess, n)
	}
}

func (p *exporter) exprList(psess *PackageSession, list Nodes) {
	if p.trace {
		if list.Len() == 0 {
			p.tracef("{}")
		} else {
			p.tracef("{>")
			defer p.tracef("<\n}")
		}
	}

	for _, n := range list.Slice() {
		if p.trace {
			p.tracef("\n")
		}
		p.expr(psess, n)
	}

	p.op(OEND)
}

func (p *exporter) elemList(psess *PackageSession, list Nodes) {
	if p.trace {
		p.tracef("[ ")
	}
	p.int(list.Len())
	if p.trace {
		if list.Len() == 0 {
			p.tracef("] {}")
		} else {
			p.tracef("] {>")
			defer p.tracef("<\n}")
		}
	}

	for _, n := range list.Slice() {
		if p.trace {
			p.tracef("\n")
		}
		p.fieldSym(psess, n.Sym, false)
		p.expr(psess, n.Left)
	}
}

func (p *exporter) expr(psess *PackageSession, n *Node) {
	if p.trace {
		p.tracef("( ")
		defer p.tracef(") ")
	}

	for n != nil && n.Implicit() && (n.Op == OIND || n.Op == OADDR) {
		n = n.Left
	}

	switch op := n.Op; op {

	case OPAREN:
		p.expr(psess, n.Left)

	case OLITERAL:
		if n.Val().Ctype(psess) == CTNIL && n.Orig != nil && n.Orig != n {
			p.expr(psess, n.Orig)
			break
		}
		p.op(OLITERAL)
		p.pos(psess, n.Pos)
		p.typ(psess, psess.unidealType(n.Type, n.Val()))
		p.value(psess, n.Val())

	case ONAME:

		if n.isMethodExpression() {
			p.op(OXDOT)
			p.pos(psess, n.Pos)
			p.expr(psess, n.Left)
			p.fieldSym(psess, n.Right.Sym, true)
			break
		}

		p.op(ONAME)
		p.pos(psess, n.Pos)
		p.sym(psess, n)

	case OTYPE:
		p.op(OTYPE)
		p.pos(psess, n.Pos)
		p.typ(psess, n.Type)

	case OPTRLIT:
		p.op(OPTRLIT)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)
		p.bool(n.Implicit())

	case OSTRUCTLIT:
		p.op(OSTRUCTLIT)
		p.pos(psess, n.Pos)
		p.typ(psess, n.Type)
		p.elemList(psess, n.List)

	case OARRAYLIT, OSLICELIT, OMAPLIT:
		p.op(OCOMPLIT)
		p.pos(psess, n.Pos)
		p.typ(psess, n.Type)
		p.exprList(psess, n.List)

	case OKEY:
		p.op(OKEY)
		p.pos(psess, n.Pos)
		p.exprsOrNil(psess, n.Left, n.Right)

	case OXDOT, ODOT, ODOTPTR, ODOTINTER, ODOTMETH:
		p.op(OXDOT)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)
		p.fieldSym(psess, n.Sym, true)

	case ODOTTYPE, ODOTTYPE2:
		p.op(ODOTTYPE)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)
		p.typ(psess, n.Type)

	case OINDEX, OINDEXMAP:
		p.op(OINDEX)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)
		p.expr(psess, n.Right)

	case OSLICE, OSLICESTR, OSLICEARR:
		p.op(OSLICE)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)
		low, high, _ := n.SliceBounds(psess)
		p.exprsOrNil(psess, low, high)

	case OSLICE3, OSLICE3ARR:
		p.op(OSLICE3)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)
		low, high, max := n.SliceBounds(psess)
		p.exprsOrNil(psess, low, high)
		p.expr(psess, max)

	case OCOPY, OCOMPLEX:

		p.op(op)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)
		p.expr(psess, n.Right)
		p.op(OEND)

	case OCONV, OCONVIFACE, OCONVNOP, OARRAYBYTESTR, OARRAYRUNESTR, OSTRARRAYBYTE, OSTRARRAYRUNE, ORUNESTR:
		p.op(OCONV)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)
		p.typ(psess, n.Type)

	case OREAL, OIMAG, OAPPEND, OCAP, OCLOSE, ODELETE, OLEN, OMAKE, ONEW, OPANIC, ORECOVER, OPRINT, OPRINTN:
		p.op(op)
		p.pos(psess, n.Pos)
		if n.Left != nil {
			p.expr(psess, n.Left)
			p.op(OEND)
		} else {
			p.exprList(psess, n.List)
		}

		if op == OAPPEND {
			p.bool(n.Isddd())
		} else if n.Isddd() {
			psess.
				Fatalf("exporter: unexpected '...' with %v call", op)
		}

	case OCALL, OCALLFUNC, OCALLMETH, OCALLINTER, OGETG:
		p.op(OCALL)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)
		p.exprList(psess, n.List)
		p.bool(n.Isddd())

	case OMAKEMAP, OMAKECHAN, OMAKESLICE:
		p.op(op)
		p.pos(psess, n.Pos)
		p.typ(psess, n.Type)
		switch {
		default:

			p.op(OEND)
		case n.List.Len() != 0:
			p.exprList(psess, n.List)
		case n.Right != nil:
			p.expr(psess, n.Left)
			p.expr(psess, n.Right)
			p.op(OEND)
		case n.Left != nil && (n.Op == OMAKESLICE || !n.Left.Type.IsUntyped(psess.types)):
			p.expr(psess, n.Left)
			p.op(OEND)
		}

	case OPLUS, OMINUS, OADDR, OCOM, OIND, ONOT, ORECV:
		p.op(op)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)

	case OADD, OAND, OANDAND, OANDNOT, ODIV, OEQ, OGE, OGT, OLE, OLT,
		OLSH, OMOD, OMUL, ONE, OOR, OOROR, ORSH, OSEND, OSUB, OXOR:
		p.op(op)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)
		p.expr(psess, n.Right)

	case OADDSTR:
		p.op(OADDSTR)
		p.pos(psess, n.Pos)
		p.exprList(psess, n.List)

	case OCMPSTR, OCMPIFACE:
		p.op(n.SubOp(psess))
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)
		p.expr(psess, n.Right)

	case ODCLCONST:

		p.op(ODCLCONST)
		p.pos(psess, n.Pos)

	default:
		psess.
			Fatalf("cannot export %v (%d) node\n"+
				"==> please file an issue and assign to gri@\n", n.Op, int(n.Op))
	}
}

// Caution: stmt will emit more than one node for statement nodes n that have a non-empty
// n.Ninit and where n cannot have a natural init section (such as in "if", "for", etc.).
func (p *exporter) stmt(psess *PackageSession, n *Node) {
	if p.trace {
		p.tracef("( ")
		defer p.tracef(") ")
	}

	if n.Ninit.Len() > 0 && !stmtwithinit(n.Op) {
		if p.trace {
			p.tracef("( /* Ninits */ ")
		}

		for _, n := range n.Ninit.Slice() {
			p.stmt(psess, n)
		}

		if p.trace {
			p.tracef(") ")
		}
	}

	switch op := n.Op; op {
	case ODCL:
		p.op(ODCL)
		p.pos(psess, n.Left.Pos)
		p.sym(psess, n.Left)
		p.typ(psess, n.Left.Type)

	case OAS:

		if n.Right != nil {
			p.op(OAS)
			p.pos(psess, n.Pos)
			p.expr(psess, n.Left)
			p.expr(psess, n.Right)
		}

	case OASOP:
		p.op(OASOP)
		p.pos(psess, n.Pos)
		p.op(n.SubOp(psess))
		p.expr(psess, n.Left)
		if p.bool(!n.Implicit()) {
			p.expr(psess, n.Right)
		}

	case OAS2, OAS2DOTTYPE, OAS2FUNC, OAS2MAPR, OAS2RECV:
		p.op(OAS2)
		p.pos(psess, n.Pos)
		p.exprList(psess, n.List)
		p.exprList(psess, n.Rlist)

	case ORETURN:
		p.op(ORETURN)
		p.pos(psess, n.Pos)
		p.exprList(psess, n.List)

	case OPROC, ODEFER:
		p.op(op)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)

	case OIF:
		p.op(OIF)
		p.pos(psess, n.Pos)
		p.stmtList(psess, n.Ninit)
		p.expr(psess, n.Left)
		p.stmtList(psess, n.Nbody)
		p.stmtList(psess, n.Rlist)

	case OFOR:
		p.op(OFOR)
		p.pos(psess, n.Pos)
		p.stmtList(psess, n.Ninit)
		p.exprsOrNil(psess, n.Left, n.Right)
		p.stmtList(psess, n.Nbody)

	case ORANGE:
		p.op(ORANGE)
		p.pos(psess, n.Pos)
		p.stmtList(psess, n.List)
		p.expr(psess, n.Right)
		p.stmtList(psess, n.Nbody)

	case OSELECT, OSWITCH:
		p.op(op)
		p.pos(psess, n.Pos)
		p.stmtList(psess, n.Ninit)
		p.exprsOrNil(psess, n.Left, nil)
		p.stmtList(psess, n.List)

	case OCASE, OXCASE:
		p.op(OXCASE)
		p.pos(psess, n.Pos)
		p.stmtList(psess, n.List)
		p.stmtList(psess, n.Nbody)

	case OFALL:
		p.op(OFALL)
		p.pos(psess, n.Pos)

	case OBREAK, OCONTINUE:
		p.op(op)
		p.pos(psess, n.Pos)
		p.exprsOrNil(psess, n.Left, nil)

	case OEMPTY:

	case OGOTO, OLABEL:
		p.op(op)
		p.pos(psess, n.Pos)
		p.expr(psess, n.Left)

	default:
		psess.
			Fatalf("exporter: CANNOT EXPORT: %v\nPlease notify gri@\n", n.Op)
	}
}

func (p *exporter) exprsOrNil(psess *PackageSession, a, b *Node) {
	ab := 0
	if a != nil {
		ab |= 1
	}
	if b != nil {
		ab |= 2
	}
	p.int(ab)
	if ab&1 != 0 {
		p.expr(psess, a)
	}
	if ab&2 != 0 {
		p.node(psess, b)
	}
}

func (p *exporter) fieldSym(psess *PackageSession, s *types.Sym, short bool) {
	name := s.Name

	if short {
		if i := strings.LastIndex(name, "."); i >= 0 {
			name = name[i+1:]
		}
	}

	p.string(name)
	if !types.IsExported(name) {
		p.pkg(psess, s.Pkg)
	}
}

// sym must encode the _ (blank) identifier as a single string "_" since
// encoding for some nodes is based on this assumption (e.g. ONAME nodes).
func (p *exporter) sym(psess *PackageSession, n *Node) {
	s := n.Sym
	if s.Pkg != nil {
		if len(s.Name) > 0 && s.Name[0] == '.' {
			psess.
				Fatalf("exporter: exporting synthetic symbol %s", s.Name)
		}
	}

	if p.trace {
		p.tracef("{ SYM ")
		defer p.tracef("} ")
	}

	name := s.Name

	if i := strings.LastIndex(name, "."); i >= 0 {
		name = name[i+1:]
	}

	if strings.Contains(name, "·") && n.Name.Vargen > 0 {
		psess.
			Fatalf("exporter: unexpected · in symbol name")
	}

	if i := n.Name.Vargen; i > 0 {
		name = fmt.Sprintf("%s·%d", name, i)
	}

	p.string(name)
	if name != "_" {
		p.pkg(psess, s.Pkg)
	}

	p.string(s.Linkname)
}

func (p *exporter) bool(b bool) bool {
	if p.trace {
		p.tracef("[")
		defer p.tracef("= %v] ", b)
	}

	x := 0
	if b {
		x = 1
	}
	p.int(x)
	return b
}

func (p *exporter) op(op Op) {
	if p.trace {
		p.tracef("[")
		defer p.tracef("= %v] ", op)
	}

	p.int(int(op))
}

func (p *exporter) index(psess *PackageSession, marker byte, index int) {
	if index < 0 {
		psess.
			Fatalf("exporter: invalid index < 0")
	}
	if debugFormat {
		p.marker('t')
	}
	if p.trace {
		p.tracef("%c%d ", marker, index)
	}
	p.rawInt64(int64(index))
}

func (p *exporter) tag(psess *PackageSession, tag int) {
	if tag >= 0 {
		psess.
			Fatalf("exporter: invalid tag >= 0")
	}
	if debugFormat {
		p.marker('t')
	}
	if p.trace {
		p.tracef("%s ", psess.tagString[-tag])
	}
	p.rawInt64(int64(tag))
}

func (p *exporter) int(x int) {
	p.int64(int64(x))
}

func (p *exporter) int64(x int64) {
	if debugFormat {
		p.marker('i')
	}
	if p.trace {
		p.tracef("%d ", x)
	}
	p.rawInt64(x)
}

func (p *exporter) string(s string) {
	if debugFormat {
		p.marker('s')
	}
	if p.trace {
		p.tracef("%q ", s)
	}

	if i, ok := p.strIndex[s]; ok {
		p.rawInt64(int64(i))
		return
	}

	p.strIndex[s] = len(p.strIndex)
	p.rawInt64(-int64(len(s)))
	for i := 0; i < len(s); i++ {
		p.rawByte(s[i])
	}
}

// marker emits a marker byte and position information which makes
// it easy for a reader to detect if it is "out of sync". Used only
// if debugFormat is set.
func (p *exporter) marker(m byte) {
	p.rawByte(m)

	p.rawInt64(int64(p.written))
}

// rawInt64 should only be used by low-level encoders.
func (p *exporter) rawInt64(x int64) {
	var tmp [binary.MaxVarintLen64]byte
	n := binary.PutVarint(tmp[:], x)
	for i := 0; i < n; i++ {
		p.rawByte(tmp[i])
	}
}

// rawStringln should only be used to emit the initial version string.
func (p *exporter) rawStringln(s string) {
	for i := 0; i < len(s); i++ {
		p.rawByte(s[i])
	}
	p.rawByte('\n')
}

// rawByte is the bottleneck interface to write to p.out.
// rawByte escapes b as follows (any encoding does that
// hides '$'):
//
//	'$'  => '|' 'S'
//	'|'  => '|' '|'
//
// Necessary so other tools can find the end of the
// export data by searching for "$$".
// rawByte should only be used by low-level encoders.
func (p *exporter) rawByte(b byte) {
	switch b {
	case '$':

		b = 'S'
		fallthrough
	case '|':

		p.out.WriteByte('|')
		p.written++
	}
	p.out.WriteByte(b)
	p.written++
}

// tracef is like fmt.Printf but it rewrites the format string
// to take care of indentation.
func (p *exporter) tracef(format string, args ...interface{}) {
	if strings.ContainsAny(format, "<>\n") {
		var buf bytes.Buffer
		for i := 0; i < len(format); i++ {

			ch := format[i]
			switch ch {
			case '>':
				p.indent++
				continue
			case '<':
				p.indent--
				continue
			}
			buf.WriteByte(ch)
			if ch == '\n' {
				for j := p.indent; j > 0; j-- {
					buf.WriteString(".  ")
				}
			}
		}
		format = buf.String()
	}
	fmt.Printf(format, args...)
}

// Tags. Must be < 0.
const (
	// Objects
	packageTag = -(iota + 1)
	constTag
	typeTag
	varTag
	funcTag
	endTag

	// Types
	namedTag
	arrayTag
	sliceTag
	dddTag
	structTag
	pointerTag
	signatureTag
	interfaceTag
	mapTag
	chanTag

	// Values
	falseTag
	trueTag
	int64Tag
	floatTag
	fractionTag // not used by gc
	complexTag
	stringTag
	nilTag
	unknownTag // not used by gc (only appears in packages with errors)

	// Type aliases
	aliasTag
)

// Debugging support.
// (tagString is only used when tracing is enabled)

// untype returns the "pseudo" untyped type for a Ctype (import/export use only).
// (we can't use an pre-initialized array because we must be sure all types are
// set up)
func (psess *PackageSession) untype(ctype Ctype) *types.Type {
	switch ctype {
	case CTINT:
		return psess.types.Idealint
	case CTRUNE:
		return psess.types.Idealrune
	case CTFLT:
		return psess.types.Idealfloat
	case CTCPLX:
		return psess.types.Idealcomplex
	case CTSTR:
		return psess.types.Idealstring
	case CTBOOL:
		return psess.types.Idealbool
	case CTNIL:
		return psess.types.Types[TNIL]
	}
	psess.
		Fatalf("exporter: unknown Ctype")
	return nil
}

// initialized lazily

func (psess *PackageSession) predeclared() []*types.Type {
	if psess.predecl == nil {
		psess.
			predecl = []*types.Type{psess.types.
			Types[TBOOL], psess.types.
			Types[TINT], psess.types.
			Types[TINT8], psess.types.
			Types[TINT16], psess.types.
			Types[TINT32], psess.types.
			Types[TINT64], psess.types.
			Types[TUINT], psess.types.
			Types[TUINT8], psess.types.
			Types[TUINT16], psess.types.
			Types[TUINT32], psess.types.
			Types[TUINT64], psess.types.
			Types[TUINTPTR], psess.types.
			Types[TFLOAT32], psess.types.
			Types[TFLOAT64], psess.types.
			Types[TCOMPLEX64], psess.types.
			Types[TCOMPLEX128], psess.types.
			Types[TSTRING], psess.types.
			Bytetype, psess.types.
			Runetype, psess.types.
			Errortype, psess.
			untype(CTBOOL), psess.
			untype(CTINT), psess.
			untype(CTRUNE), psess.
			untype(CTFLT), psess.
			untype(CTCPLX), psess.
			untype(CTSTR), psess.
			untype(CTNIL), psess.types.
			Types[TUNSAFEPTR], psess.types.
			Types[Txxx], psess.types.
			Types[TANY],
		}
	}
	return psess.predecl
}

package gc

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/src"
	"math/big"
	"strconv"
	"strings"
)

type importer struct {
	in      *bufio.Reader
	imp     *types.Pkg // imported package
	buf     []byte     // reused for reading strings
	version int        // export format version

	// object lists, in order of deserialization
	strList       []string
	pathList      []string
	pkgList       []*types.Pkg
	typList       []*types.Type
	funcList      []*Node // nil entry means already declared
	trackAllTypes bool

	// for delayed type verification
	cmpList []struct{ pt, t *types.Type }

	// position encoding
	posInfoFormat bool
	prevFile      string
	prevLine      int
	posBase       *src.PosBase

	// debugging support
	debugFormat bool
	read        int // bytes read
}

// Import populates imp from the serialized package data read from in.
func (psess *PackageSession) Import(imp *types.Pkg, in *bufio.Reader) {
	psess.
		inimport = true
	defer func() {
		psess.inimport = false
	}()

	p := importer{
		in:       in,
		imp:      imp,
		version:  -1,
		strList:  []string{""},
		pathList: []string{""},
	}

	// read version info
	var versionstr string
	if b := p.rawByte(psess); b == 'c' || b == 'd' {

		if b == 'd' {
			p.debugFormat = true
		}
		p.trackAllTypes = p.rawByte(psess) == 'a'
		p.posInfoFormat = p.bool(psess)
		versionstr = p.string(psess)
		if versionstr == "v1" {
			p.version = 0
		}
	} else {

		versionstr = p.rawStringln(psess, b)
		if s := strings.SplitN(versionstr, " ", 3); len(s) >= 2 && s[0] == "version" {
			if v, err := strconv.Atoi(s[1]); err == nil && v > 0 {
				p.version = v
			}
		}
	}

	switch p.version {

	case 6, 5, 4, 3, 2, 1:
		p.debugFormat = p.rawStringln(psess, p.rawByte(psess)) == "debug"
		p.trackAllTypes = p.bool(psess)
		p.posInfoFormat = p.bool(psess)
	case 0:

	default:
		p.formatErrorf(psess, "unknown export format version %d (%q)", p.version, versionstr)
	}

	p.typList = append(p.typList, psess.predeclared()...)

	p.pkg(psess)

	tcok := psess.typecheckok
	psess.
		typecheckok = true
	psess.
		defercheckwidth()

	objcount := 0
	for {
		tag := p.tagOrIndex(psess)
		if tag == endTag {
			break
		}
		p.obj(psess, tag)
		objcount++
	}

	if count := p.int(psess); count != objcount {
		p.formatErrorf(psess, "got %d objects; want %d", objcount, count)
	}

	objcount = 0
	for {
		tag := p.tagOrIndex(psess)
		if tag == endTag {
			break
		}
		p.obj(psess, tag)
		objcount++
	}

	if count := p.int(psess); count != objcount {
		p.formatErrorf(psess, "got %d objects; want %d", objcount, count)
	}

	if psess.dclcontext != PEXTERN {
		p.formatErrorf(psess, "unexpected context %d", psess.dclcontext)
	}

	objcount = 0
	for i0 := -1; ; {
		i := p.int(psess)
		if i < 0 {
			break
		}

		if i <= i0 {
			p.formatErrorf(psess, "index not increasing: %d <= %d", i, i0)
		}
		i0 = i

		if psess.Curfn != nil {
			p.formatErrorf(psess, "unexpected Curfn %v", psess.Curfn)
		}

		inlCost := p.int(psess)
		if f := p.funcList[i]; f != nil && f.Func.Inl == nil {
			psess.
				funchdr(f)
			body := p.stmtList(psess)
			psess.
				funcbody()
			f.Func.Inl = &Inline{
				Cost: int32(inlCost),
				Body: body,
			}
			psess.
				importlist = append(psess.importlist, f)
			if psess.Debug['E'] > 0 && psess.Debug['m'] > 2 {
				if psess.Debug['m'] > 3 {
					fmt.Printf("inl body for %v: %+v\n", f, asNodes(body))
				} else {
					fmt.Printf("inl body for %v: %v\n", f, asNodes(body))
				}
			}
		} else {
			psess.
				dclcontext = PDISCARD
			p.stmtList(psess)
			psess.
				dclcontext = PEXTERN
		}

		objcount++
	}

	if count := p.int(psess); count != objcount {
		p.formatErrorf(psess, "got %d functions; want %d", objcount, count)
	}

	if psess.dclcontext != PEXTERN {
		p.formatErrorf(psess, "unexpected context %d", psess.dclcontext)
	}

	p.verifyTypes(psess)
	psess.
		typecheckok = tcok
	psess.
		resumecheckwidth()

	if psess.debug_dclstack != 0 {
		psess.
			testdclstack()
	}
}

func (p *importer) formatErrorf(psess *PackageSession, format string, args ...interface{}) {
	if debugFormat {
		psess.
			Fatalf(format, args...)
	}
	psess.
		yyerror("cannot import %q due to version skew - reinstall package (%s)",
			p.imp.Path, fmt.Sprintf(format, args...))
	psess.
		errorexit()
}

func (p *importer) verifyTypes(psess *PackageSession) {
	for _, pair := range p.cmpList {
		pt := pair.pt
		t := pair.t
		if !psess.eqtype(pt.Orig, t) {
			p.formatErrorf(psess, "inconsistent definition for type %v during import\n\t%L (in %q)\n\t%L (in %q)", pt.Sym, pt, pt.Sym.Importdef.Path, t, p.imp.Path)
		}
	}
}

// numImport tracks how often a package with a given name is imported.
// It is used to provide a better error message (by using the package
// path to disambiguate) if a package that appears multiple times with
// the same name appears in an error message.

func (p *importer) pkg(psess *PackageSession) *types.Pkg {

	i := p.tagOrIndex(psess)
	if i >= 0 {
		return p.pkgList[i]
	}

	if i != packageTag {
		p.formatErrorf(psess, "expected package tag, found tag = %d", i)
	}

	name := p.string(psess)
	var path string
	if p.version >= 5 {
		path = p.path(psess)
	} else {
		path = p.string(psess)
	}
	var height int
	if p.version >= 6 {
		height = p.int(psess)
	}

	if name == "" {
		p.formatErrorf(psess, "empty package name for path %q", path)
	}

	if psess.isbadimport(path, true) {
		p.formatErrorf(psess, "bad package path %q for package %s", path, name)
	}

	if (path == "") != (len(p.pkgList) == 0) {
		p.formatErrorf(psess, "package path %q for pkg index %d", path, len(p.pkgList))
	}

	if p.version >= 6 {
		if height < 0 || height >= types.MaxPkgHeight {
			p.formatErrorf(psess, "bad package height %v for package %s", height, name)
		}

		if len(p.pkgList) != 0 && height >= p.imp.Height {
			p.formatErrorf(psess, "package %q (height %d) reexports package %q (height %d)", p.imp.Path, p.imp.Height, path, height)
		}
	}

	pkg := p.imp
	if path != "" {
		pkg = psess.types.NewPkg(path, "")
	}
	if pkg.Name == "" {
		pkg.Name = name
		psess.
			numImport[name]++
	} else if pkg.Name != name {
		psess.
			yyerror("conflicting package names %s and %s for path %q", pkg.Name, name, path)
	}
	if psess.myimportpath != "" && path == psess.myimportpath {
		psess.
			yyerror("import %q: package depends on %q (import cycle)", p.imp.Path, path)
		psess.
			errorexit()
	}
	pkg.Height = height
	p.pkgList = append(p.pkgList, pkg)

	return pkg
}

func (psess *PackageSession) idealType(typ *types.Type) *types.Type {
	switch typ {
	case psess.types.Idealint, psess.types.Idealrune, psess.types.Idealfloat, psess.types.Idealcomplex:

		typ = psess.types.Types[TIDEAL]
	}
	return typ
}

func (p *importer) obj(psess *PackageSession, tag int) {
	switch tag {
	case constTag:
		pos := p.pos(psess)
		sym := p.qualifiedName(psess)
		typ := p.typ(psess)
		val := p.value(psess, typ)
		psess.
			importconst(p.imp, pos, sym, psess.idealType(typ), val)

	case aliasTag:
		pos := p.pos(psess)
		sym := p.qualifiedName(psess)
		typ := p.typ(psess)
		psess.
			importalias(p.imp, pos, sym, typ)

	case typeTag:
		p.typ(psess)

	case varTag:
		pos := p.pos(psess)
		sym := p.qualifiedName(psess)
		typ := p.typ(psess)
		psess.
			importvar(p.imp, pos, sym, typ)

	case funcTag:
		pos := p.pos(psess)
		sym := p.qualifiedName(psess)
		params := p.paramList(psess)
		result := p.paramList(psess)

		sig := psess.functypefield(nil, params, result)
		psess.
			importfunc(p.imp, pos, sym, sig)
		p.funcList = append(p.funcList, asNode(sym.Def))

	default:
		p.formatErrorf(psess, "unexpected object (tag = %d)", tag)
	}
}

func (p *importer) pos(psess *PackageSession) src.XPos {
	if !p.posInfoFormat {
		return psess.src.NoXPos
	}

	file := p.prevFile
	line := p.prevLine
	delta := p.int(psess)
	line += delta
	if p.version >= 5 {
		if delta == deltaNewFile {
			if n := p.int(psess); n >= 0 {

				file = p.path(psess)
				line = n
			}
		}
	} else {
		if delta == 0 {
			if n := p.int(psess); n >= 0 {

				file = p.prevFile[:n] + p.string(psess)
				line = p.int(psess)
			}
		}
	}
	if file != p.prevFile {
		p.prevFile = file
		p.posBase = src.NewFileBase(file, file)
	}
	p.prevLine = line

	pos := src.MakePos(p.posBase, uint(line), 0)
	xpos := psess.Ctxt.PosTable.XPos(pos)
	return xpos
}

func (p *importer) path(psess *PackageSession) string {

	i := p.int(psess)
	if i >= 0 {
		return p.pathList[i]
	}

	a := make([]string, -i)
	for n := range a {
		a[n] = p.string(psess)
	}
	s := strings.Join(a, "/")
	p.pathList = append(p.pathList, s)
	return s
}

func (p *importer) newtyp(etype types.EType) *types.Type {
	t := types.New(etype)
	if p.trackAllTypes {
		p.typList = append(p.typList, t)
	}
	return t
}

// importtype declares that pt, an imported named type, has underlying type t.
func (p *importer) importtype(psess *PackageSession, pt, t *types.Type) {
	if pt.Etype == TFORW {
		psess.
			copytype(psess.typenod(pt), t)
		psess.
			checkwidth(pt)
	} else {

		if p.trackAllTypes {

			p.cmpList = append(p.cmpList, struct{ pt, t *types.Type }{pt, t})
		} else if !psess.eqtype(pt.Orig, t) {
			psess.
				yyerror("inconsistent definition for type %v during import\n\t%L (in %q)\n\t%L (in %q)", pt.Sym, pt, pt.Sym.Importdef.Path, t, p.imp.Path)
		}
	}

	if psess.Debug['E'] != 0 {
		fmt.Printf("import type %v %L\n", pt, t)
	}
}

func (p *importer) typ(psess *PackageSession) *types.Type {

	i := p.tagOrIndex(psess)
	if i >= 0 {
		return p.typList[i]
	}

	// otherwise, i is the type tag (< 0)
	var t *types.Type
	switch i {
	case namedTag:
		pos := p.pos(psess)
		tsym := p.qualifiedName(psess)

		t = psess.importtype(p.imp, pos, tsym)
		p.typList = append(p.typList, t)
		dup := !t.IsKind(types.TFORW)

		t0 := p.typ(psess)
		p.importtype(psess, t, t0)

		if t0.IsInterface() {
			break
		}

		savedContext := psess.dclcontext
		psess.
			dclcontext = PEXTERN

		for i := p.int(psess); i > 0; i-- {
			mpos := p.pos(psess)
			sym := p.fieldSym(psess)

			if !types.IsExported(sym.Name) && sym.Pkg != tsym.Pkg {
				psess.
					Fatalf("imported method name %+v in wrong package %s\n", sym, tsym.Pkg.Name)
			}

			recv := p.paramList(psess)
			params := p.paramList(psess)
			result := p.paramList(psess)
			nointerface := p.bool(psess)

			mt := psess.functypefield(recv[0], params, result)
			oldm := psess.addmethod(sym, mt, false, nointerface)

			if dup {

				n := asNode(oldm.Type.Nname(psess.types))
				p.funcList = append(p.funcList, n)
				continue
			}

			n := psess.newfuncnamel(mpos, psess.methodSym(recv[0].Type, sym))
			n.Type = mt
			n.SetClass(PFUNC)
			psess.
				checkwidth(n.Type)
			p.funcList = append(p.funcList, n)

			mt.SetNname(psess.types, asTypesNode(n))

			if psess.Debug['E'] > 0 {
				fmt.Printf("import [%q] meth %v \n", p.imp.Path, n)
			}
		}
		psess.
			dclcontext = savedContext

	case arrayTag:
		t = p.newtyp(TARRAY)
		bound := p.int64(psess)
		elem := p.typ(psess)
		t.Extra = &types.Array{Elem: elem, Bound: bound}

	case sliceTag:
		t = p.newtyp(TSLICE)
		elem := p.typ(psess)
		t.Extra = types.Slice{Elem: elem}

	case dddTag:
		t = p.newtyp(TDDDFIELD)
		t.Extra = types.DDDField{T: p.typ(psess)}

	case structTag:
		t = p.newtyp(TSTRUCT)
		t.SetFields(psess.types, p.fieldList(psess))
		psess.
			checkwidth(t)

	case pointerTag:
		t = p.newtyp(psess.types.Tptr)
		t.Extra = types.Ptr{Elem: p.typ(psess)}

	case signatureTag:
		t = p.newtyp(TFUNC)
		params := p.paramList(psess)
		result := p.paramList(psess)
		psess.
			functypefield0(t, nil, params, result)

	case interfaceTag:
		if ml := p.methodList(psess); len(ml) == 0 {
			t = psess.types.Types[TINTER]
		} else {
			t = p.newtyp(TINTER)
			t.SetInterface(psess.types, ml)
		}

	case mapTag:
		t = p.newtyp(TMAP)
		mt := t.MapType(psess.types)
		mt.Key = p.typ(psess)
		mt.Elem = p.typ(psess)

	case chanTag:
		t = p.newtyp(TCHAN)
		ct := t.ChanType(psess.types)
		ct.Dir = types.ChanDir(p.int(psess))
		ct.Elem = p.typ(psess)

	default:
		p.formatErrorf(psess, "unexpected type (tag = %d)", i)
	}

	if t == nil {
		p.formatErrorf(psess, "nil type (type tag = %d)", i)
	}

	return t
}

func (p *importer) qualifiedName(psess *PackageSession) *types.Sym {
	name := p.string(psess)
	pkg := p.pkg(psess)
	return pkg.Lookup(psess.types, name)
}

func (p *importer) fieldList(psess *PackageSession) (fields []*types.Field) {
	if n := p.int(psess); n > 0 {
		fields = make([]*types.Field, n)
		for i := range fields {
			fields[i] = p.field(psess)
		}
	}
	return
}

func (p *importer) field(psess *PackageSession) *types.Field {
	pos := p.pos(psess)
	sym, alias := p.fieldName(psess)
	typ := p.typ(psess)
	note := p.string(psess)

	f := types.NewField()
	if sym.Name == "" {

		s := typ.Sym
		if s == nil && typ.IsPtr() {
			s = typ.Elem(psess.types).Sym
		}
		sym = sym.Pkg.Lookup(psess.types, s.Name)
		f.Embedded = 1
	} else if alias {

		f.Embedded = 1
	}

	f.Pos = pos
	f.Sym = sym
	f.Type = typ
	f.Note = note

	return f
}

func (p *importer) methodList(psess *PackageSession) (methods []*types.Field) {
	for n := p.int(psess); n > 0; n-- {
		f := types.NewField()
		f.Pos = p.pos(psess)
		f.Type = p.typ(psess)
		methods = append(methods, f)
	}

	for n := p.int(psess); n > 0; n-- {
		methods = append(methods, p.method(psess))
	}

	return
}

func (p *importer) method(psess *PackageSession) *types.Field {
	pos := p.pos(psess)
	sym := p.methodName(psess)
	params := p.paramList(psess)
	result := p.paramList(psess)

	f := types.NewField()
	f.Pos = pos
	f.Sym = sym
	f.Type = psess.functypefield(psess.fakeRecvField(), params, result)
	return f
}

func (p *importer) fieldName(psess *PackageSession) (*types.Sym, bool) {
	name := p.string(psess)
	if p.version == 0 && name == "_" {

		return psess.builtinpkg.Lookup(psess.types, name), false
	}
	pkg := psess.localpkg
	alias := false
	switch name {
	case "":

	case "?":

		name = ""
		pkg = p.pkg(psess)
	case "@":

		name = p.string(psess)
		alias = true
		fallthrough
	default:
		if !types.IsExported(name) {
			pkg = p.pkg(psess)
		}
	}
	return pkg.Lookup(psess.types, name), alias
}

func (p *importer) methodName(psess *PackageSession) *types.Sym {
	name := p.string(psess)
	if p.version == 0 && name == "_" {

		return psess.builtinpkg.Lookup(psess.types, name)
	}
	pkg := psess.localpkg
	if !types.IsExported(name) {
		pkg = p.pkg(psess)
	}
	return pkg.Lookup(psess.types, name)
}

func (p *importer) paramList(psess *PackageSession) []*types.Field {
	i := p.int(psess)
	if i == 0 {
		return nil
	}

	named := true
	if i < 0 {
		i = -i
		named = false
	}

	fs := make([]*types.Field, i)
	for i := range fs {
		fs[i] = p.param(psess, named)
	}
	return fs
}

func (p *importer) param(psess *PackageSession, named bool) *types.Field {
	f := types.NewField()

	f.Pos = psess.lineno
	f.Type = p.typ(psess)
	if f.Type.Etype == TDDDFIELD {

		f.Type = psess.types.NewSlice(f.Type.DDDField(psess.types))
		f.SetIsddd(true)
	}

	if named {
		name := p.string(psess)
		if name == "" {
			p.formatErrorf(psess, "expected named parameter")
		}

		pkg := psess.localpkg
		if name != "_" {
			pkg = p.pkg(psess)
		}
		f.Sym = pkg.Lookup(psess.types, name)
	}

	f.Note = p.string(psess)

	return f
}

func (p *importer) value(psess *PackageSession, typ *types.Type) (x Val) {
	switch tag := p.tagOrIndex(psess); tag {
	case falseTag:
		x.U = false

	case trueTag:
		x.U = true

	case int64Tag:
		u := new(Mpint)
		u.SetInt64(p.int64(psess))
		u.Rune = typ == psess.types.Idealrune
		x.U = u

	case floatTag:
		f := newMpflt()
		p.float(psess, f)
		if typ == psess.types.Idealint || typ.IsInteger() || typ.IsPtr() || typ.IsUnsafePtr() {

			u := new(Mpint)
			u.SetFloat(f)
			x.U = u
			break
		}
		x.U = f

	case complexTag:
		u := new(Mpcplx)
		p.float(psess, &u.Real)
		p.float(psess, &u.Imag)
		x.U = u

	case stringTag:
		x.U = p.string(psess)

	case unknownTag:
		p.formatErrorf(psess, "unknown constant (importing package with errors)")

	case nilTag:
		x.U = new(NilVal)

	default:
		p.formatErrorf(psess, "unexpected value tag %d", tag)
	}

	if typ.IsUntyped(psess.types) && psess.untype(x.Ctype(psess)) != typ {
		p.formatErrorf(psess, "value %v and type %v don't match", x, typ)
	}

	return
}

func (p *importer) float(psess *PackageSession, x *Mpflt) {
	sign := p.int(psess)
	if sign == 0 {
		x.SetFloat64(0)
		return
	}

	exp := p.int(psess)
	mant := new(big.Int).SetBytes([]byte(p.string(psess)))

	m := x.Val.SetInt(mant)
	m.SetMantExp(m, exp-mant.BitLen())
	if sign < 0 {
		m.Neg(m)
	}
}

func (p *importer) stmtList(psess *PackageSession) []*Node {
	var list []*Node
	for {
		n := p.node(psess)
		if n == nil {
			break
		}

		if n.Op == OBLOCK {
			list = append(list, n.List.Slice()...)
		} else {
			list = append(list, n)
		}
	}
	return list
}

func (p *importer) exprList(psess *PackageSession) []*Node {
	var list []*Node
	for {
		n := p.expr(psess)
		if n == nil {
			break
		}
		list = append(list, n)
	}
	return list
}

func (p *importer) elemList(psess *PackageSession) []*Node {
	c := p.int(psess)
	list := make([]*Node, c)
	for i := range list {
		s := p.fieldSym(psess)
		list[i] = psess.nodSym(OSTRUCTKEY, p.expr(psess), s)
	}
	return list
}

func (p *importer) expr(psess *PackageSession) *Node {
	n := p.node(psess)
	if n != nil && n.Op == OBLOCK {
		psess.
			Fatalf("unexpected block node: %v", n)
	}
	return n
}

func npos(pos src.XPos, n *Node) *Node {
	n.Pos = pos
	return n
}

// TODO(gri) split into expr and stmt
func (p *importer) node(psess *PackageSession) *Node {
	switch op := p.op(psess); op {

	case OLITERAL:
		pos := p.pos(psess)
		typ := p.typ(psess)
		n := npos(pos, psess.nodlit(p.value(psess, typ)))
		n.Type = psess.idealType(typ)
		return n

	case ONAME:
		return npos(p.pos(psess), psess.mkname(p.sym(psess)))

	case OTYPE:
		return npos(p.pos(psess), psess.typenod(p.typ(psess)))

	case OPTRLIT:
		pos := p.pos(psess)
		n := npos(pos, p.expr(psess))
		if !p.bool(psess) {
			if n.Op == OCOMPLIT {

				n.Right = psess.nodl(pos, OIND, n.Right, nil)
				n.Right.SetImplicit(true)
			} else {
				n = psess.nodl(pos, OADDR, n, nil)
			}
		}
		return n

	case OSTRUCTLIT:

		savedlineno := psess.lineno
		psess.
			lineno = p.pos(psess)
		n := psess.nodl(psess.lineno, OCOMPLIT, nil, psess.typenod(p.typ(psess)))
		n.List.Set(p.elemList(psess))
		psess.
			lineno = savedlineno
		return n

	case OCOMPLIT:
		n := psess.nodl(p.pos(psess), OCOMPLIT, nil, psess.typenod(p.typ(psess)))
		n.List.Set(p.exprList(psess))
		return n

	case OKEY:
		pos := p.pos(psess)
		left, right := p.exprsOrNil(psess)
		return psess.nodl(pos, OKEY, left, right)

	case OXDOT:

		return npos(p.pos(psess), psess.nodSym(OXDOT, p.expr(psess), p.fieldSym(psess)))

	case ODOTTYPE:
		n := psess.nodl(p.pos(psess), ODOTTYPE, p.expr(psess), nil)
		n.Type = p.typ(psess)
		return n

	case OINDEX:
		return psess.nodl(p.pos(psess), op, p.expr(psess), p.expr(psess))

	case OSLICE, OSLICE3:
		n := psess.nodl(p.pos(psess), op, p.expr(psess), nil)
		low, high := p.exprsOrNil(psess)
		var max *Node
		if n.Op.IsSlice3(psess) {
			max = p.expr(psess)
		}
		n.SetSliceBounds(psess, low, high, max)
		return n

	case OCONV:
		n := psess.nodl(p.pos(psess), OCONV, p.expr(psess), nil)
		n.Type = p.typ(psess)
		return n

	case OCOPY, OCOMPLEX, OREAL, OIMAG, OAPPEND, OCAP, OCLOSE, ODELETE, OLEN, OMAKE, ONEW, OPANIC, ORECOVER, OPRINT, OPRINTN:
		n := npos(p.pos(psess), psess.builtinCall(op))
		n.List.Set(p.exprList(psess))
		if op == OAPPEND {
			n.SetIsddd(p.bool(psess))
		}
		return n

	case OCALL:
		n := psess.nodl(p.pos(psess), OCALL, p.expr(psess), nil)
		n.List.Set(p.exprList(psess))
		n.SetIsddd(p.bool(psess))
		return n

	case OMAKEMAP, OMAKECHAN, OMAKESLICE:
		n := npos(p.pos(psess), psess.builtinCall(OMAKE))
		n.List.Append(psess.typenod(p.typ(psess)))
		n.List.Append(p.exprList(psess)...)
		return n

	case OPLUS, OMINUS, OADDR, OCOM, OIND, ONOT, ORECV:
		return psess.nodl(p.pos(psess), op, p.expr(psess), nil)

	case OADD, OAND, OANDAND, OANDNOT, ODIV, OEQ, OGE, OGT, OLE, OLT,
		OLSH, OMOD, OMUL, ONE, OOR, OOROR, ORSH, OSEND, OSUB, OXOR:
		return psess.nodl(p.pos(psess), op, p.expr(psess), p.expr(psess))

	case OADDSTR:
		pos := p.pos(psess)
		list := p.exprList(psess)
		x := npos(pos, list[0])
		for _, y := range list[1:] {
			x = psess.nodl(pos, OADD, x, y)
		}
		return x

	case ODCLCONST:

		return psess.nodl(p.pos(psess), OEMPTY, nil, nil)

	case ODCL:
		if p.version < 2 {

			p.bool(psess)
		}
		pos := p.pos(psess)
		lhs := npos(pos, psess.dclname(p.sym(psess)))
		typ := psess.typenod(p.typ(psess))
		return npos(pos, psess.liststmt(psess.variter([]*Node{lhs}, typ, nil)))

	case OAS:
		return psess.nodl(p.pos(psess), OAS, p.expr(psess), p.expr(psess))

	case OASOP:
		n := psess.nodl(p.pos(psess), OASOP, nil, nil)
		n.SetSubOp(psess, p.op(psess))
		n.Left = p.expr(psess)
		if !p.bool(psess) {
			n.Right = psess.nodintconst(1)
			n.SetImplicit(true)
		} else {
			n.Right = p.expr(psess)
		}
		return n

	case OAS2:
		n := psess.nodl(p.pos(psess), OAS2, nil, nil)
		n.List.Set(p.exprList(psess))
		n.Rlist.Set(p.exprList(psess))
		return n

	case ORETURN:
		n := psess.nodl(p.pos(psess), ORETURN, nil, nil)
		n.List.Set(p.exprList(psess))
		return n

	case OPROC, ODEFER:
		return psess.nodl(p.pos(psess), op, p.expr(psess), nil)

	case OIF:
		n := psess.nodl(p.pos(psess), OIF, nil, nil)
		n.Ninit.Set(p.stmtList(psess))
		n.Left = p.expr(psess)
		n.Nbody.Set(p.stmtList(psess))
		n.Rlist.Set(p.stmtList(psess))
		return n

	case OFOR:
		n := psess.nodl(p.pos(psess), OFOR, nil, nil)
		n.Ninit.Set(p.stmtList(psess))
		n.Left, n.Right = p.exprsOrNil(psess)
		n.Nbody.Set(p.stmtList(psess))
		return n

	case ORANGE:
		n := psess.nodl(p.pos(psess), ORANGE, nil, nil)
		n.List.Set(p.stmtList(psess))
		n.Right = p.expr(psess)
		n.Nbody.Set(p.stmtList(psess))
		return n

	case OSELECT, OSWITCH:
		n := psess.nodl(p.pos(psess), op, nil, nil)
		n.Ninit.Set(p.stmtList(psess))
		n.Left, _ = p.exprsOrNil(psess)
		n.List.Set(p.stmtList(psess))
		return n

	case OXCASE:
		n := psess.nodl(p.pos(psess), OXCASE, nil, nil)
		n.List.Set(p.exprList(psess))

		n.Nbody.Set(p.stmtList(psess))
		return n

	case OFALL:
		n := psess.nodl(p.pos(psess), OFALL, nil, nil)
		return n

	case OBREAK, OCONTINUE:
		pos := p.pos(psess)
		left, _ := p.exprsOrNil(psess)
		if left != nil {
			left = psess.newname(left.Sym)
		}
		return psess.nodl(pos, op, left, nil)

	case OGOTO, OLABEL:
		return psess.nodl(p.pos(psess), op, psess.newname(p.expr(psess).Sym), nil)

	case OEND:
		return nil

	default:
		psess.
			Fatalf("cannot import %v (%d) node\n"+
				"==> please file an issue and assign to gri@\n", op, int(op))
		panic("unreachable")
	}
}

func (psess *PackageSession) builtinCall(op Op) *Node {
	return psess.nod(OCALL, psess.mkname(psess.builtinpkg.Lookup(psess.types, psess.goopnames[op])), nil)
}

func (p *importer) exprsOrNil(psess *PackageSession) (a, b *Node) {
	ab := p.int(psess)
	if ab&1 != 0 {
		a = p.expr(psess)
	}
	if ab&2 != 0 {
		b = p.node(psess)
	}
	return
}

func (p *importer) fieldSym(psess *PackageSession) *types.Sym {
	name := p.string(psess)
	pkg := psess.localpkg
	if !types.IsExported(name) {
		pkg = p.pkg(psess)
	}
	return pkg.Lookup(psess.types, name)
}

func (p *importer) sym(psess *PackageSession) *types.Sym {
	name := p.string(psess)
	pkg := psess.localpkg
	if name != "_" {
		pkg = p.pkg(psess)
	}
	linkname := p.string(psess)
	sym := pkg.Lookup(psess.types, name)
	sym.Linkname = linkname
	return sym
}

func (p *importer) bool(psess *PackageSession) bool {
	return p.int(psess) != 0
}

func (p *importer) op(psess *PackageSession) Op {
	return Op(p.int(psess))
}

func (p *importer) tagOrIndex(psess *PackageSession) int {
	if p.debugFormat {
		p.marker(psess, 't')
	}

	return int(p.rawInt64(psess))
}

func (p *importer) int(psess *PackageSession) int {
	x := p.int64(psess)
	if int64(int(x)) != x {
		p.formatErrorf(psess, "exported integer too large")
	}
	return int(x)
}

func (p *importer) int64(psess *PackageSession) int64 {
	if p.debugFormat {
		p.marker(psess, 'i')
	}

	return p.rawInt64(psess)
}

func (p *importer) string(psess *PackageSession) string {
	if p.debugFormat {
		p.marker(psess, 's')
	}

	i := p.rawInt64(psess)
	if i >= 0 {
		return p.strList[i]
	}

	if n := int(-i); n <= cap(p.buf) {
		p.buf = p.buf[:n]
	} else {
		p.buf = make([]byte, n)
	}
	for i := range p.buf {
		p.buf[i] = p.rawByte(psess)
	}
	s := string(p.buf)
	p.strList = append(p.strList, s)
	return s
}

func (p *importer) marker(psess *PackageSession, want byte) {
	if got := p.rawByte(psess); got != want {
		p.formatErrorf(psess, "incorrect marker: got %c; want %c (pos = %d)", got, want, p.read)
	}

	pos := p.read
	if n := int(p.rawInt64(psess)); n != pos {
		p.formatErrorf(psess, "incorrect position: got %d; want %d", n, pos)
	}
}

// rawInt64 should only be used by low-level decoders.
func (p *importer) rawInt64(psess *PackageSession) int64 {
	i, err := binary.ReadVarint(p)
	if err != nil {
		p.formatErrorf(psess, "read error: %v", err)
	}
	return i
}

// rawStringln should only be used to read the initial version string.
func (p *importer) rawStringln(psess *PackageSession, b byte) string {
	p.buf = p.buf[:0]
	for b != '\n' {
		p.buf = append(p.buf, b)
		b = p.rawByte(psess)
	}
	return string(p.buf)
}

// needed for binary.ReadVarint in rawInt64
func (p *importer) ReadByte(psess *PackageSession) (byte, error) {
	return p.rawByte(psess), nil
}

// rawByte is the bottleneck interface for reading from p.in.
// It unescapes '|' 'S' to '$' and '|' '|' to '|'.
// rawByte should only be used by low-level decoders.
func (p *importer) rawByte(psess *PackageSession) byte {
	c, err := p.in.ReadByte()
	p.read++
	if err != nil {
		p.formatErrorf(psess, "read error: %v", err)
	}
	if c == '|' {
		c, err = p.in.ReadByte()
		p.read++
		if err != nil {
			p.formatErrorf(psess, "read error: %v", err)
		}
		switch c {
		case 'S':
			c = '$'
		case '|':

		default:
			p.formatErrorf(psess, "unexpected escape sequence in export data")
		}
	}
	return c
}

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Binary package import.
// See bexport.go for the export data format and how
// to make a format change.

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

// The overall structure of Import is symmetric to Export: For each
// export method in bexport.go there is a matching and symmetric method
// in bimport.go. Changing the export format requires making symmetric
// changes to bimport.go and bexport.go.

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
func (pstate *PackageState) Import(imp *types.Pkg, in *bufio.Reader) {
	pstate.inimport = true
	defer func() { pstate.inimport = false }()

	p := importer{
		in:       in,
		imp:      imp,
		version:  -1,           // unknown version
		strList:  []string{""}, // empty string is mapped to 0
		pathList: []string{""}, // empty path is mapped to 0
	}

	// read version info
	var versionstr string
	if b := p.rawByte(pstate); b == 'c' || b == 'd' {
		// Go1.7 encoding; first byte encodes low-level
		// encoding format (compact vs debug).
		// For backward-compatibility only (avoid problems with
		// old installed packages). Newly compiled packages use
		// the extensible format string.
		// TODO(gri) Remove this support eventually; after Go1.8.
		if b == 'd' {
			p.debugFormat = true
		}
		p.trackAllTypes = p.rawByte(pstate) == 'a'
		p.posInfoFormat = p.bool(pstate)
		versionstr = p.string(pstate)
		if versionstr == "v1" {
			p.version = 0
		}
	} else {
		// Go1.8 extensible encoding
		// read version string and extract version number (ignore anything after the version number)
		versionstr = p.rawStringln(pstate, b)
		if s := strings.SplitN(versionstr, " ", 3); len(s) >= 2 && s[0] == "version" {
			if v, err := strconv.Atoi(s[1]); err == nil && v > 0 {
				p.version = v
			}
		}
	}

	// read version specific flags - extend as necessary
	switch p.version {
	// case 7:
	// 	...
	//	fallthrough
	case 6, 5, 4, 3, 2, 1:
		p.debugFormat = p.rawStringln(pstate, p.rawByte(pstate)) == "debug"
		p.trackAllTypes = p.bool(pstate)
		p.posInfoFormat = p.bool(pstate)
	case 0:
	// Go1.7 encoding format - nothing to do here
	default:
		p.formatErrorf(pstate, "unknown export format version %d (%q)", p.version, versionstr)
	}

	// --- generic export data ---

	// populate typList with predeclared "known" types
	p.typList = append(p.typList, pstate.predeclared()...)

	// read package data
	p.pkg(pstate)

	// defer some type-checking until all types are read in completely
	tcok := pstate.typecheckok
	pstate.typecheckok = true
	pstate.defercheckwidth()

	// read objects

	// phase 1
	objcount := 0
	for {
		tag := p.tagOrIndex(pstate)
		if tag == endTag {
			break
		}
		p.obj(pstate, tag)
		objcount++
	}

	// self-verification
	if count := p.int(pstate); count != objcount {
		p.formatErrorf(pstate, "got %d objects; want %d", objcount, count)
	}

	// --- compiler-specific export data ---

	// read compiler-specific flags

	// phase 2
	objcount = 0
	for {
		tag := p.tagOrIndex(pstate)
		if tag == endTag {
			break
		}
		p.obj(pstate, tag)
		objcount++
	}

	// self-verification
	if count := p.int(pstate); count != objcount {
		p.formatErrorf(pstate, "got %d objects; want %d", objcount, count)
	}

	// read inlineable functions bodies
	if pstate.dclcontext != PEXTERN {
		p.formatErrorf(pstate, "unexpected context %d", pstate.dclcontext)
	}

	objcount = 0
	for i0 := -1; ; {
		i := p.int(pstate) // index of function with inlineable body
		if i < 0 {
			break
		}

		// don't process the same function twice
		if i <= i0 {
			p.formatErrorf(pstate, "index not increasing: %d <= %d", i, i0)
		}
		i0 = i

		if pstate.Curfn != nil {
			p.formatErrorf(pstate, "unexpected Curfn %v", pstate.Curfn)
		}

		// Note: In the original code, funchdr and funcbody are called for
		// all functions (that were not yet imported). Now, we are calling
		// them only for functions with inlineable bodies. funchdr does
		// parameter renaming which doesn't matter if we don't have a body.

		inlCost := p.int(pstate)
		if f := p.funcList[i]; f != nil && f.Func.Inl == nil {
			// function not yet imported - read body and set it
			pstate.funchdr(f)
			body := p.stmtList(pstate)
			pstate.funcbody()
			f.Func.Inl = &Inline{
				Cost: int32(inlCost),
				Body: body,
			}
			pstate.importlist = append(pstate.importlist, f)
			if pstate.Debug['E'] > 0 && pstate.Debug['m'] > 2 {
				if pstate.Debug['m'] > 3 {
					fmt.Printf("inl body for %v: %+v\n", f, asNodes(body))
				} else {
					fmt.Printf("inl body for %v: %v\n", f, asNodes(body))
				}
			}
		} else {
			// function already imported - read body but discard declarations
			pstate.dclcontext = PDISCARD // throw away any declarations
			p.stmtList(pstate)
			pstate.dclcontext = PEXTERN
		}

		objcount++
	}

	// self-verification
	if count := p.int(pstate); count != objcount {
		p.formatErrorf(pstate, "got %d functions; want %d", objcount, count)
	}

	if pstate.dclcontext != PEXTERN {
		p.formatErrorf(pstate, "unexpected context %d", pstate.dclcontext)
	}

	p.verifyTypes(pstate)

	// --- end of export data ---

	pstate.typecheckok = tcok
	pstate.resumecheckwidth()

	if pstate.debug_dclstack != 0 {
		pstate.testdclstack()
	}
}

func (p *importer) formatErrorf(pstate *PackageState, format string, args ...interface{}) {
	if debugFormat {
		pstate.Fatalf(format, args...)
	}

	pstate.yyerror("cannot import %q due to version skew - reinstall package (%s)",
		p.imp.Path, fmt.Sprintf(format, args...))
	pstate.errorexit()
}

func (p *importer) verifyTypes(pstate *PackageState) {
	for _, pair := range p.cmpList {
		pt := pair.pt
		t := pair.t
		if !pstate.eqtype(pt.Orig, t) {
			p.formatErrorf(pstate, "inconsistent definition for type %v during import\n\t%L (in %q)\n\t%L (in %q)", pt.Sym, pt, pt.Sym.Importdef.Path, t, p.imp.Path)
		}
	}
}

func (p *importer) pkg(pstate *PackageState) *types.Pkg {
	// if the package was seen before, i is its index (>= 0)
	i := p.tagOrIndex(pstate)
	if i >= 0 {
		return p.pkgList[i]
	}

	// otherwise, i is the package tag (< 0)
	if i != packageTag {
		p.formatErrorf(pstate, "expected package tag, found tag = %d", i)
	}

	// read package data
	name := p.string(pstate)
	var path string
	if p.version >= 5 {
		path = p.path(pstate)
	} else {
		path = p.string(pstate)
	}
	var height int
	if p.version >= 6 {
		height = p.int(pstate)
	}

	// we should never see an empty package name
	if name == "" {
		p.formatErrorf(pstate, "empty package name for path %q", path)
	}

	// we should never see a bad import path
	if pstate.isbadimport(path, true) {
		p.formatErrorf(pstate, "bad package path %q for package %s", path, name)
	}

	// an empty path denotes the package we are currently importing;
	// it must be the first package we see
	if (path == "") != (len(p.pkgList) == 0) {
		p.formatErrorf(pstate, "package path %q for pkg index %d", path, len(p.pkgList))
	}

	if p.version >= 6 {
		if height < 0 || height >= types.MaxPkgHeight {
			p.formatErrorf(pstate, "bad package height %v for package %s", height, name)
		}

		// reexported packages should always have a lower height than
		// the main package
		if len(p.pkgList) != 0 && height >= p.imp.Height {
			p.formatErrorf(pstate, "package %q (height %d) reexports package %q (height %d)", p.imp.Path, p.imp.Height, path, height)
		}
	}

	// add package to pkgList
	pkg := p.imp
	if path != "" {
		pkg = pstate.types.NewPkg(path, "")
	}
	if pkg.Name == "" {
		pkg.Name = name
		pstate.numImport[name]++
	} else if pkg.Name != name {
		pstate.yyerror("conflicting package names %s and %s for path %q", pkg.Name, name, path)
	}
	if pstate.myimportpath != "" && path == pstate.myimportpath {
		pstate.yyerror("import %q: package depends on %q (import cycle)", p.imp.Path, path)
		pstate.errorexit()
	}
	pkg.Height = height
	p.pkgList = append(p.pkgList, pkg)

	return pkg
}

func (pstate *PackageState) idealType(typ *types.Type) *types.Type {
	switch typ {
	case pstate.types.Idealint, pstate.types.Idealrune, pstate.types.Idealfloat, pstate.types.Idealcomplex:
		// canonicalize ideal types
		typ = pstate.types.Types[TIDEAL]
	}
	return typ
}

func (p *importer) obj(pstate *PackageState, tag int) {
	switch tag {
	case constTag:
		pos := p.pos(pstate)
		sym := p.qualifiedName(pstate)
		typ := p.typ(pstate)
		val := p.value(pstate, typ)
		pstate.importconst(p.imp, pos, sym, pstate.idealType(typ), val)

	case aliasTag:
		pos := p.pos(pstate)
		sym := p.qualifiedName(pstate)
		typ := p.typ(pstate)
		pstate.importalias(p.imp, pos, sym, typ)

	case typeTag:
		p.typ(pstate)

	case varTag:
		pos := p.pos(pstate)
		sym := p.qualifiedName(pstate)
		typ := p.typ(pstate)
		pstate.importvar(p.imp, pos, sym, typ)

	case funcTag:
		pos := p.pos(pstate)
		sym := p.qualifiedName(pstate)
		params := p.paramList(pstate)
		result := p.paramList(pstate)

		sig := pstate.functypefield(nil, params, result)
		pstate.importfunc(p.imp, pos, sym, sig)
		p.funcList = append(p.funcList, asNode(sym.Def))

	default:
		p.formatErrorf(pstate, "unexpected object (tag = %d)", tag)
	}
}

func (p *importer) pos(pstate *PackageState) src.XPos {
	if !p.posInfoFormat {
		return pstate.src.NoXPos
	}

	file := p.prevFile
	line := p.prevLine
	delta := p.int(pstate)
	line += delta
	if p.version >= 5 {
		if delta == deltaNewFile {
			if n := p.int(pstate); n >= 0 {
				// file changed
				file = p.path(pstate)
				line = n
			}
		}
	} else {
		if delta == 0 {
			if n := p.int(pstate); n >= 0 {
				// file changed
				file = p.prevFile[:n] + p.string(pstate)
				line = p.int(pstate)
			}
		}
	}
	if file != p.prevFile {
		p.prevFile = file
		p.posBase = src.NewFileBase(file, file)
	}
	p.prevLine = line

	pos := src.MakePos(p.posBase, uint(line), 0)
	xpos := pstate.Ctxt.PosTable.XPos(pos)
	return xpos
}

func (p *importer) path(pstate *PackageState) string {
	// if the path was seen before, i is its index (>= 0)
	// (the empty string is at index 0)
	i := p.int(pstate)
	if i >= 0 {
		return p.pathList[i]
	}
	// otherwise, i is the negative path length (< 0)
	a := make([]string, -i)
	for n := range a {
		a[n] = p.string(pstate)
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
func (p *importer) importtype(pstate *PackageState, pt, t *types.Type) {
	if pt.Etype == TFORW {
		pstate.copytype(pstate.typenod(pt), t)
		pstate.checkwidth(pt)
	} else {
		// pt.Orig and t must be identical.
		if p.trackAllTypes {
			// If we track all types, t may not be fully set up yet.
			// Collect the types and verify identity later.
			p.cmpList = append(p.cmpList, struct{ pt, t *types.Type }{pt, t})
		} else if !pstate.eqtype(pt.Orig, t) {
			pstate.yyerror("inconsistent definition for type %v during import\n\t%L (in %q)\n\t%L (in %q)", pt.Sym, pt, pt.Sym.Importdef.Path, t, p.imp.Path)
		}
	}

	if pstate.Debug['E'] != 0 {
		fmt.Printf("import type %v %L\n", pt, t)
	}
}

func (p *importer) typ(pstate *PackageState) *types.Type {
	// if the type was seen before, i is its index (>= 0)
	i := p.tagOrIndex(pstate)
	if i >= 0 {
		return p.typList[i]
	}

	// otherwise, i is the type tag (< 0)
	var t *types.Type
	switch i {
	case namedTag:
		pos := p.pos(pstate)
		tsym := p.qualifiedName(pstate)

		t = pstate.importtype(p.imp, pos, tsym)
		p.typList = append(p.typList, t)
		dup := !t.IsKind(types.TFORW) // type already imported

		// read underlying type
		t0 := p.typ(pstate)
		p.importtype(pstate, t, t0)

		// interfaces don't have associated methods
		if t0.IsInterface() {
			break
		}

		// set correct import context (since p.typ() may be called
		// while importing the body of an inlined function)
		savedContext := pstate.dclcontext
		pstate.dclcontext = PEXTERN

		// read associated methods
		for i := p.int(pstate); i > 0; i-- {
			mpos := p.pos(pstate)
			sym := p.fieldSym(pstate)

			// during import unexported method names should be in the type's package
			if !types.IsExported(sym.Name) && sym.Pkg != tsym.Pkg {
				pstate.Fatalf("imported method name %+v in wrong package %s\n", sym, tsym.Pkg.Name)
			}

			recv := p.paramList(pstate) // TODO(gri) do we need a full param list for the receiver?
			params := p.paramList(pstate)
			result := p.paramList(pstate)
			nointerface := p.bool(pstate)

			mt := pstate.functypefield(recv[0], params, result)
			oldm := pstate.addmethod(sym, mt, false, nointerface)

			if dup {
				// An earlier import already declared this type and its methods.
				// Discard the duplicate method declaration.
				n := asNode(oldm.Type.Nname(pstate.types))
				p.funcList = append(p.funcList, n)
				continue
			}

			n := pstate.newfuncnamel(mpos, pstate.methodSym(recv[0].Type, sym))
			n.Type = mt
			n.SetClass(PFUNC)
			pstate.checkwidth(n.Type)
			p.funcList = append(p.funcList, n)

			// (comment from parser.go)
			// inl.C's inlnode in on a dotmeth node expects to find the inlineable body as
			// (dotmeth's type).Nname.Inl, and dotmeth's type has been pulled
			// out by typecheck's lookdot as this $$.ttype. So by providing
			// this back link here we avoid special casing there.
			mt.SetNname(pstate.types, asTypesNode(n))

			if pstate.Debug['E'] > 0 {
				fmt.Printf("import [%q] meth %v \n", p.imp.Path, n)
			}
		}

		pstate.dclcontext = savedContext

	case arrayTag:
		t = p.newtyp(TARRAY)
		bound := p.int64(pstate)
		elem := p.typ(pstate)
		t.Extra = &types.Array{Elem: elem, Bound: bound}

	case sliceTag:
		t = p.newtyp(TSLICE)
		elem := p.typ(pstate)
		t.Extra = types.Slice{Elem: elem}

	case dddTag:
		t = p.newtyp(TDDDFIELD)
		t.Extra = types.DDDField{T: p.typ(pstate)}

	case structTag:
		t = p.newtyp(TSTRUCT)
		t.SetFields(pstate.types, p.fieldList(pstate))
		pstate.checkwidth(t)

	case pointerTag:
		t = p.newtyp(pstate.types.Tptr)
		t.Extra = types.Ptr{Elem: p.typ(pstate)}

	case signatureTag:
		t = p.newtyp(TFUNC)
		params := p.paramList(pstate)
		result := p.paramList(pstate)
		pstate.functypefield0(t, nil, params, result)

	case interfaceTag:
		if ml := p.methodList(pstate); len(ml) == 0 {
			t = pstate.types.Types[TINTER]
		} else {
			t = p.newtyp(TINTER)
			t.SetInterface(pstate.types, ml)
		}

	case mapTag:
		t = p.newtyp(TMAP)
		mt := t.MapType(pstate.types)
		mt.Key = p.typ(pstate)
		mt.Elem = p.typ(pstate)

	case chanTag:
		t = p.newtyp(TCHAN)
		ct := t.ChanType(pstate.types)
		ct.Dir = types.ChanDir(p.int(pstate))
		ct.Elem = p.typ(pstate)

	default:
		p.formatErrorf(pstate, "unexpected type (tag = %d)", i)
	}

	if t == nil {
		p.formatErrorf(pstate, "nil type (type tag = %d)", i)
	}

	return t
}

func (p *importer) qualifiedName(pstate *PackageState) *types.Sym {
	name := p.string(pstate)
	pkg := p.pkg(pstate)
	return pkg.Lookup(pstate.types, name)
}

func (p *importer) fieldList(pstate *PackageState) (fields []*types.Field) {
	if n := p.int(pstate); n > 0 {
		fields = make([]*types.Field, n)
		for i := range fields {
			fields[i] = p.field(pstate)
		}
	}
	return
}

func (p *importer) field(pstate *PackageState) *types.Field {
	pos := p.pos(pstate)
	sym, alias := p.fieldName(pstate)
	typ := p.typ(pstate)
	note := p.string(pstate)

	f := types.NewField()
	if sym.Name == "" {
		// anonymous field: typ must be T or *T and T must be a type name
		s := typ.Sym
		if s == nil && typ.IsPtr() {
			s = typ.Elem(pstate.types).Sym // deref
		}
		sym = sym.Pkg.Lookup(pstate.types, s.Name)
		f.Embedded = 1
	} else if alias {
		// anonymous field: we have an explicit name because it's a type alias
		f.Embedded = 1
	}

	f.Pos = pos
	f.Sym = sym
	f.Type = typ
	f.Note = note

	return f
}

func (p *importer) methodList(pstate *PackageState) (methods []*types.Field) {
	for n := p.int(pstate); n > 0; n-- {
		f := types.NewField()
		f.Pos = p.pos(pstate)
		f.Type = p.typ(pstate)
		methods = append(methods, f)
	}

	for n := p.int(pstate); n > 0; n-- {
		methods = append(methods, p.method(pstate))
	}

	return
}

func (p *importer) method(pstate *PackageState) *types.Field {
	pos := p.pos(pstate)
	sym := p.methodName(pstate)
	params := p.paramList(pstate)
	result := p.paramList(pstate)

	f := types.NewField()
	f.Pos = pos
	f.Sym = sym
	f.Type = pstate.functypefield(pstate.fakeRecvField(), params, result)
	return f
}

func (p *importer) fieldName(pstate *PackageState) (*types.Sym, bool) {
	name := p.string(pstate)
	if p.version == 0 && name == "_" {
		// version 0 didn't export a package for _ field names
		// but used the builtin package instead
		return pstate.builtinpkg.Lookup(pstate.types, name), false
	}
	pkg := pstate.localpkg
	alias := false
	switch name {
	case "":
	// 1) field name matches base type name and is exported: nothing to do
	case "?":
		// 2) field name matches base type name and is not exported: need package
		name = ""
		pkg = p.pkg(pstate)
	case "@":
		// 3) field name doesn't match base type name (alias name): need name and possibly package
		name = p.string(pstate)
		alias = true
		fallthrough
	default:
		if !types.IsExported(name) {
			pkg = p.pkg(pstate)
		}
	}
	return pkg.Lookup(pstate.types, name), alias
}

func (p *importer) methodName(pstate *PackageState) *types.Sym {
	name := p.string(pstate)
	if p.version == 0 && name == "_" {
		// version 0 didn't export a package for _ method names
		// but used the builtin package instead
		return pstate.builtinpkg.Lookup(pstate.types, name)
	}
	pkg := pstate.localpkg
	if !types.IsExported(name) {
		pkg = p.pkg(pstate)
	}
	return pkg.Lookup(pstate.types, name)
}

func (p *importer) paramList(pstate *PackageState) []*types.Field {
	i := p.int(pstate)
	if i == 0 {
		return nil
	}
	// negative length indicates unnamed parameters
	named := true
	if i < 0 {
		i = -i
		named = false
	}
	// i > 0
	fs := make([]*types.Field, i)
	for i := range fs {
		fs[i] = p.param(pstate, named)
	}
	return fs
}

func (p *importer) param(pstate *PackageState, named bool) *types.Field {
	f := types.NewField()
	// TODO(mdempsky): Need param position.
	f.Pos = pstate.lineno
	f.Type = p.typ(pstate)
	if f.Type.Etype == TDDDFIELD {
		// TDDDFIELD indicates wrapped ... slice type
		f.Type = pstate.types.NewSlice(f.Type.DDDField(pstate.types))
		f.SetIsddd(true)
	}

	if named {
		name := p.string(pstate)
		if name == "" {
			p.formatErrorf(pstate, "expected named parameter")
		}
		// TODO(gri) Supply function/method package rather than
		// encoding the package for each parameter repeatedly.
		pkg := pstate.localpkg
		if name != "_" {
			pkg = p.pkg(pstate)
		}
		f.Sym = pkg.Lookup(pstate.types, name)
	}

	// TODO(gri) This is compiler-specific (escape info).
	// Move into compiler-specific section eventually?
	f.Note = p.string(pstate)

	return f
}

func (p *importer) value(pstate *PackageState, typ *types.Type) (x Val) {
	switch tag := p.tagOrIndex(pstate); tag {
	case falseTag:
		x.U = false

	case trueTag:
		x.U = true

	case int64Tag:
		u := new(Mpint)
		u.SetInt64(p.int64(pstate))
		u.Rune = typ == pstate.types.Idealrune
		x.U = u

	case floatTag:
		f := newMpflt()
		p.float(pstate, f)
		if typ == pstate.types.Idealint || typ.IsInteger() || typ.IsPtr() || typ.IsUnsafePtr() {
			// uncommon case: large int encoded as float
			//
			// This happens for unsigned typed integers
			// and (on 64-bit platforms) pointers because
			// of values in the range [2^63, 2^64).
			u := new(Mpint)
			u.SetFloat(f)
			x.U = u
			break
		}
		x.U = f

	case complexTag:
		u := new(Mpcplx)
		p.float(pstate, &u.Real)
		p.float(pstate, &u.Imag)
		x.U = u

	case stringTag:
		x.U = p.string(pstate)

	case unknownTag:
		p.formatErrorf(pstate, "unknown constant (importing package with errors)")

	case nilTag:
		x.U = new(NilVal)

	default:
		p.formatErrorf(pstate, "unexpected value tag %d", tag)
	}

	// verify ideal type
	if typ.IsUntyped(pstate.types) && pstate.untype(x.Ctype(pstate)) != typ {
		p.formatErrorf(pstate, "value %v and type %v don't match", x, typ)
	}

	return
}

func (p *importer) float(pstate *PackageState, x *Mpflt) {
	sign := p.int(pstate)
	if sign == 0 {
		x.SetFloat64(0)
		return
	}

	exp := p.int(pstate)
	mant := new(big.Int).SetBytes([]byte(p.string(pstate)))

	m := x.Val.SetInt(mant)
	m.SetMantExp(m, exp-mant.BitLen())
	if sign < 0 {
		m.Neg(m)
	}
}

// ----------------------------------------------------------------------------
// Inlined function bodies

// Approach: Read nodes and use them to create/declare the same data structures
// as done originally by the (hidden) parser by closely following the parser's
// original code. In other words, "parsing" the import data (which happens to
// be encoded in binary rather textual form) is the best way at the moment to
// re-establish the syntax tree's invariants. At some future point we might be
// able to avoid this round-about way and create the rewritten nodes directly,
// possibly avoiding a lot of duplicate work (name resolution, type checking).
//
// Refined nodes (e.g., ODOTPTR as a refinement of OXDOT) are exported as their
// unrefined nodes (since this is what the importer uses). The respective case
// entries are unreachable in the importer.

func (p *importer) stmtList(pstate *PackageState) []*Node {
	var list []*Node
	for {
		n := p.node(pstate)
		if n == nil {
			break
		}
		// OBLOCK nodes may be created when importing ODCL nodes - unpack them
		if n.Op == OBLOCK {
			list = append(list, n.List.Slice()...)
		} else {
			list = append(list, n)
		}
	}
	return list
}

func (p *importer) exprList(pstate *PackageState) []*Node {
	var list []*Node
	for {
		n := p.expr(pstate)
		if n == nil {
			break
		}
		list = append(list, n)
	}
	return list
}

func (p *importer) elemList(pstate *PackageState) []*Node {
	c := p.int(pstate)
	list := make([]*Node, c)
	for i := range list {
		s := p.fieldSym(pstate)
		list[i] = pstate.nodSym(OSTRUCTKEY, p.expr(pstate), s)
	}
	return list
}

func (p *importer) expr(pstate *PackageState) *Node {
	n := p.node(pstate)
	if n != nil && n.Op == OBLOCK {
		pstate.Fatalf("unexpected block node: %v", n)
	}
	return n
}

func npos(pos src.XPos, n *Node) *Node {
	n.Pos = pos
	return n
}

// TODO(gri) split into expr and stmt
func (p *importer) node(pstate *PackageState) *Node {
	switch op := p.op(pstate); op {
	// expressions
	// case OPAREN:
	// 	unreachable - unpacked by exporter

	// case ODDDARG:
	//	unimplemented

	case OLITERAL:
		pos := p.pos(pstate)
		typ := p.typ(pstate)
		n := npos(pos, pstate.nodlit(p.value(pstate, typ)))
		n.Type = pstate.idealType(typ)
		return n

	case ONAME:
		return npos(p.pos(pstate), pstate.mkname(p.sym(pstate)))

	// case OPACK, ONONAME:
	// 	unreachable - should have been resolved by typechecking

	case OTYPE:
		return npos(p.pos(pstate), pstate.typenod(p.typ(pstate)))

	// case OTARRAY, OTMAP, OTCHAN, OTSTRUCT, OTINTER, OTFUNC:
	//      unreachable - should have been resolved by typechecking

	// case OCLOSURE:
	//	unimplemented

	case OPTRLIT:
		pos := p.pos(pstate)
		n := npos(pos, p.expr(pstate))
		if !p.bool(pstate) /* !implicit, i.e. '&' operator */ {
			if n.Op == OCOMPLIT {
				// Special case for &T{...}: turn into (*T){...}.
				n.Right = pstate.nodl(pos, OIND, n.Right, nil)
				n.Right.SetImplicit(true)
			} else {
				n = pstate.nodl(pos, OADDR, n, nil)
			}
		}
		return n

	case OSTRUCTLIT:
		// TODO(mdempsky): Export position information for OSTRUCTKEY nodes.
		savedlineno := pstate.lineno
		pstate.lineno = p.pos(pstate)
		n := pstate.nodl(pstate.lineno, OCOMPLIT, nil, pstate.typenod(p.typ(pstate)))
		n.List.Set(p.elemList(pstate)) // special handling of field names
		pstate.lineno = savedlineno
		return n

	// case OARRAYLIT, OSLICELIT, OMAPLIT:
	// 	unreachable - mapped to case OCOMPLIT below by exporter

	case OCOMPLIT:
		n := pstate.nodl(p.pos(pstate), OCOMPLIT, nil, pstate.typenod(p.typ(pstate)))
		n.List.Set(p.exprList(pstate))
		return n

	case OKEY:
		pos := p.pos(pstate)
		left, right := p.exprsOrNil(pstate)
		return pstate.nodl(pos, OKEY, left, right)

	// case OSTRUCTKEY:
	//	unreachable - handled in case OSTRUCTLIT by elemList

	// case OCALLPART:
	//	unimplemented

	// case OXDOT, ODOT, ODOTPTR, ODOTINTER, ODOTMETH:
	// 	unreachable - mapped to case OXDOT below by exporter

	case OXDOT:
		// see parser.new_dotname
		return npos(p.pos(pstate), pstate.nodSym(OXDOT, p.expr(pstate), p.fieldSym(pstate)))

	// case ODOTTYPE, ODOTTYPE2:
	// 	unreachable - mapped to case ODOTTYPE below by exporter

	case ODOTTYPE:
		n := pstate.nodl(p.pos(pstate), ODOTTYPE, p.expr(pstate), nil)
		n.Type = p.typ(pstate)
		return n

	// case OINDEX, OINDEXMAP, OSLICE, OSLICESTR, OSLICEARR, OSLICE3, OSLICE3ARR:
	// 	unreachable - mapped to cases below by exporter

	case OINDEX:
		return pstate.nodl(p.pos(pstate), op, p.expr(pstate), p.expr(pstate))

	case OSLICE, OSLICE3:
		n := pstate.nodl(p.pos(pstate), op, p.expr(pstate), nil)
		low, high := p.exprsOrNil(pstate)
		var max *Node
		if n.Op.IsSlice3(pstate) {
			max = p.expr(pstate)
		}
		n.SetSliceBounds(pstate, low, high, max)
		return n

	// case OCONV, OCONVIFACE, OCONVNOP, OARRAYBYTESTR, OARRAYRUNESTR, OSTRARRAYBYTE, OSTRARRAYRUNE, ORUNESTR:
	// 	unreachable - mapped to OCONV case below by exporter

	case OCONV:
		n := pstate.nodl(p.pos(pstate), OCONV, p.expr(pstate), nil)
		n.Type = p.typ(pstate)
		return n

	case OCOPY, OCOMPLEX, OREAL, OIMAG, OAPPEND, OCAP, OCLOSE, ODELETE, OLEN, OMAKE, ONEW, OPANIC, ORECOVER, OPRINT, OPRINTN:
		n := npos(p.pos(pstate), pstate.builtinCall(op))
		n.List.Set(p.exprList(pstate))
		if op == OAPPEND {
			n.SetIsddd(p.bool(pstate))
		}
		return n

	// case OCALL, OCALLFUNC, OCALLMETH, OCALLINTER, OGETG:
	// 	unreachable - mapped to OCALL case below by exporter

	case OCALL:
		n := pstate.nodl(p.pos(pstate), OCALL, p.expr(pstate), nil)
		n.List.Set(p.exprList(pstate))
		n.SetIsddd(p.bool(pstate))
		return n

	case OMAKEMAP, OMAKECHAN, OMAKESLICE:
		n := npos(p.pos(pstate), pstate.builtinCall(OMAKE))
		n.List.Append(pstate.typenod(p.typ(pstate)))
		n.List.Append(p.exprList(pstate)...)
		return n

	// unary expressions
	case OPLUS, OMINUS, OADDR, OCOM, OIND, ONOT, ORECV:
		return pstate.nodl(p.pos(pstate), op, p.expr(pstate), nil)

	// binary expressions
	case OADD, OAND, OANDAND, OANDNOT, ODIV, OEQ, OGE, OGT, OLE, OLT,
		OLSH, OMOD, OMUL, ONE, OOR, OOROR, ORSH, OSEND, OSUB, OXOR:
		return pstate.nodl(p.pos(pstate), op, p.expr(pstate), p.expr(pstate))

	case OADDSTR:
		pos := p.pos(pstate)
		list := p.exprList(pstate)
		x := npos(pos, list[0])
		for _, y := range list[1:] {
			x = pstate.nodl(pos, OADD, x, y)
		}
		return x

	// case OCMPSTR, OCMPIFACE:
	// 	unreachable - mapped to std comparison operators by exporter

	case ODCLCONST:
		// TODO(gri) these should not be exported in the first place
		return pstate.nodl(p.pos(pstate), OEMPTY, nil, nil)

	// --------------------------------------------------------------------
	// statements
	case ODCL:
		if p.version < 2 {
			// versions 0 and 1 exported a bool here but it
			// was always false - simply ignore in this case
			p.bool(pstate)
		}
		pos := p.pos(pstate)
		lhs := npos(pos, pstate.dclname(p.sym(pstate)))
		typ := pstate.typenod(p.typ(pstate))
		return npos(pos, pstate.liststmt(pstate.variter([]*Node{lhs}, typ, nil))) // TODO(gri) avoid list creation

	// case ODCLFIELD:
	//	unimplemented

	// case OAS, OASWB:
	// 	unreachable - mapped to OAS case below by exporter

	case OAS:
		return pstate.nodl(p.pos(pstate), OAS, p.expr(pstate), p.expr(pstate))

	case OASOP:
		n := pstate.nodl(p.pos(pstate), OASOP, nil, nil)
		n.SetSubOp(pstate, p.op(pstate))
		n.Left = p.expr(pstate)
		if !p.bool(pstate) {
			n.Right = pstate.nodintconst(1)
			n.SetImplicit(true)
		} else {
			n.Right = p.expr(pstate)
		}
		return n

	// case OAS2DOTTYPE, OAS2FUNC, OAS2MAPR, OAS2RECV:
	// 	unreachable - mapped to OAS2 case below by exporter

	case OAS2:
		n := pstate.nodl(p.pos(pstate), OAS2, nil, nil)
		n.List.Set(p.exprList(pstate))
		n.Rlist.Set(p.exprList(pstate))
		return n

	case ORETURN:
		n := pstate.nodl(p.pos(pstate), ORETURN, nil, nil)
		n.List.Set(p.exprList(pstate))
		return n

	// case ORETJMP:
	// 	unreachable - generated by compiler for trampolin routines (not exported)

	case OPROC, ODEFER:
		return pstate.nodl(p.pos(pstate), op, p.expr(pstate), nil)

	case OIF:
		n := pstate.nodl(p.pos(pstate), OIF, nil, nil)
		n.Ninit.Set(p.stmtList(pstate))
		n.Left = p.expr(pstate)
		n.Nbody.Set(p.stmtList(pstate))
		n.Rlist.Set(p.stmtList(pstate))
		return n

	case OFOR:
		n := pstate.nodl(p.pos(pstate), OFOR, nil, nil)
		n.Ninit.Set(p.stmtList(pstate))
		n.Left, n.Right = p.exprsOrNil(pstate)
		n.Nbody.Set(p.stmtList(pstate))
		return n

	case ORANGE:
		n := pstate.nodl(p.pos(pstate), ORANGE, nil, nil)
		n.List.Set(p.stmtList(pstate))
		n.Right = p.expr(pstate)
		n.Nbody.Set(p.stmtList(pstate))
		return n

	case OSELECT, OSWITCH:
		n := pstate.nodl(p.pos(pstate), op, nil, nil)
		n.Ninit.Set(p.stmtList(pstate))
		n.Left, _ = p.exprsOrNil(pstate)
		n.List.Set(p.stmtList(pstate))
		return n

	// case OCASE, OXCASE:
	// 	unreachable - mapped to OXCASE case below by exporter

	case OXCASE:
		n := pstate.nodl(p.pos(pstate), OXCASE, nil, nil)
		n.List.Set(p.exprList(pstate))
		// TODO(gri) eventually we must declare variables for type switch
		// statements (type switch statements are not yet exported)
		n.Nbody.Set(p.stmtList(pstate))
		return n

	// case OFALL:
	// 	unreachable - mapped to OXFALL case below by exporter

	case OFALL:
		n := pstate.nodl(p.pos(pstate), OFALL, nil, nil)
		return n

	case OBREAK, OCONTINUE:
		pos := p.pos(pstate)
		left, _ := p.exprsOrNil(pstate)
		if left != nil {
			left = pstate.newname(left.Sym)
		}
		return pstate.nodl(pos, op, left, nil)

	// case OEMPTY:
	// 	unreachable - not emitted by exporter

	case OGOTO, OLABEL:
		return pstate.nodl(p.pos(pstate), op, pstate.newname(p.expr(pstate).Sym), nil)

	case OEND:
		return nil

	default:
		pstate.Fatalf("cannot import %v (%d) node\n"+
			"==> please file an issue and assign to gri@\n", op, int(op))
		panic("unreachable") // satisfy compiler
	}
}

func (pstate *PackageState) builtinCall(op Op) *Node {
	return pstate.nod(OCALL, pstate.mkname(pstate.builtinpkg.Lookup(pstate.types, pstate.goopnames[op])), nil)
}

func (p *importer) exprsOrNil(pstate *PackageState) (a, b *Node) {
	ab := p.int(pstate)
	if ab&1 != 0 {
		a = p.expr(pstate)
	}
	if ab&2 != 0 {
		b = p.node(pstate)
	}
	return
}

func (p *importer) fieldSym(pstate *PackageState) *types.Sym {
	name := p.string(pstate)
	pkg := pstate.localpkg
	if !types.IsExported(name) {
		pkg = p.pkg(pstate)
	}
	return pkg.Lookup(pstate.types, name)
}

func (p *importer) sym(pstate *PackageState) *types.Sym {
	name := p.string(pstate)
	pkg := pstate.localpkg
	if name != "_" {
		pkg = p.pkg(pstate)
	}
	linkname := p.string(pstate)
	sym := pkg.Lookup(pstate.types, name)
	sym.Linkname = linkname
	return sym
}

func (p *importer) bool(pstate *PackageState) bool {
	return p.int(pstate) != 0
}

func (p *importer) op(pstate *PackageState) Op {
	return Op(p.int(pstate))
}

// ----------------------------------------------------------------------------
// Low-level decoders

func (p *importer) tagOrIndex(pstate *PackageState) int {
	if p.debugFormat {
		p.marker(pstate, 't')
	}

	return int(p.rawInt64(pstate))
}

func (p *importer) int(pstate *PackageState) int {
	x := p.int64(pstate)
	if int64(int(x)) != x {
		p.formatErrorf(pstate, "exported integer too large")
	}
	return int(x)
}

func (p *importer) int64(pstate *PackageState) int64 {
	if p.debugFormat {
		p.marker(pstate, 'i')
	}

	return p.rawInt64(pstate)
}

func (p *importer) string(pstate *PackageState) string {
	if p.debugFormat {
		p.marker(pstate, 's')
	}
	// if the string was seen before, i is its index (>= 0)
	// (the empty string is at index 0)
	i := p.rawInt64(pstate)
	if i >= 0 {
		return p.strList[i]
	}
	// otherwise, i is the negative string length (< 0)
	if n := int(-i); n <= cap(p.buf) {
		p.buf = p.buf[:n]
	} else {
		p.buf = make([]byte, n)
	}
	for i := range p.buf {
		p.buf[i] = p.rawByte(pstate)
	}
	s := string(p.buf)
	p.strList = append(p.strList, s)
	return s
}

func (p *importer) marker(pstate *PackageState, want byte) {
	if got := p.rawByte(pstate); got != want {
		p.formatErrorf(pstate, "incorrect marker: got %c; want %c (pos = %d)", got, want, p.read)
	}

	pos := p.read
	if n := int(p.rawInt64(pstate)); n != pos {
		p.formatErrorf(pstate, "incorrect position: got %d; want %d", n, pos)
	}
}

// rawInt64 should only be used by low-level decoders.
func (p *importer) rawInt64(pstate *PackageState) int64 {
	i, err := binary.ReadVarint(p)
	if err != nil {
		p.formatErrorf(pstate, "read error: %v", err)
	}
	return i
}

// rawStringln should only be used to read the initial version string.
func (p *importer) rawStringln(pstate *PackageState, b byte) string {
	p.buf = p.buf[:0]
	for b != '\n' {
		p.buf = append(p.buf, b)
		b = p.rawByte(pstate)
	}
	return string(p.buf)
}

// needed for binary.ReadVarint in rawInt64
func (p *importer) ReadByte(pstate *PackageState) (byte, error) {
	return p.rawByte(pstate), nil
}

// rawByte is the bottleneck interface for reading from p.in.
// It unescapes '|' 'S' to '$' and '|' '|' to '|'.
// rawByte should only be used by low-level decoders.
func (p *importer) rawByte(pstate *PackageState) byte {
	c, err := p.in.ReadByte()
	p.read++
	if err != nil {
		p.formatErrorf(pstate, "read error: %v", err)
	}
	if c == '|' {
		c, err = p.in.ReadByte()
		p.read++
		if err != nil {
			p.formatErrorf(pstate, "read error: %v", err)
		}
		switch c {
		case 'S':
			c = '$'
		case '|':
		// nothing to do
		default:
			p.formatErrorf(pstate, "unexpected escape sequence in export data")
		}
	}
	return c
}

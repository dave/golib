// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Indexed package import.
// See iexport.go for the export data format.

package gc

import (
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/bio"
	"github.com/dave/golib/src/cmd/internal/src"
	"math/big"
	"os"
	"strings"
)

// An iimporterAndOffset identifies an importer and an offset within
// its data section.
type iimporterAndOffset struct {
	p   *iimporter
	off uint64
}

func (pstate *PackageState) expandDecl(n *Node) {
	if n.Op != ONONAME {
		return
	}

	r := importReaderFor(n, pstate.declImporter)
	if r == nil {
		// Can happen if user tries to reference an undeclared name.
		return
	}

	r.doDecl(pstate, n)
}

func (pstate *PackageState) expandInline(fn *Node) {
	if fn.Func.Inl.Body != nil {
		return
	}

	r := importReaderFor(fn, pstate.inlineImporter)
	if r == nil {
		pstate.Fatalf("missing import reader for %v", fn)
	}

	r.doInline(pstate, fn)
}

func importReaderFor(n *Node, importers map[*types.Sym]iimporterAndOffset) *importReader {
	x, ok := importers[n.Sym]
	if !ok {
		return nil
	}

	return x.p.newReader(x.off, n.Sym.Pkg)
}

type intReader struct {
	*bio.Reader
	pkg *types.Pkg
}

func (r *intReader) int64(pstate *PackageState) int64 {
	i, err := binary.ReadVarint(r.Reader)
	if err != nil {
		pstate.yyerror("import %q: read error: %v", r.pkg.Path, err)
		pstate.errorexit()
	}
	return i
}

func (r *intReader) uint64(pstate *PackageState) uint64 {
	i, err := binary.ReadUvarint(r.Reader)
	if err != nil {
		pstate.yyerror("import %q: read error: %v", r.pkg.Path, err)
		pstate.errorexit()
	}
	return i
}

func (pstate *PackageState) iimport(pkg *types.Pkg, in *bio.Reader) {
	ir := &intReader{in, pkg}

	version := ir.uint64(pstate)
	if version != iexportVersion {
		pstate.yyerror("import %q: unknown export format version %d", pkg.Path, version)
		pstate.errorexit()
	}

	sLen := ir.uint64(pstate)
	dLen := ir.uint64(pstate)

	// Map string (and data) section into memory as a single large
	// string. This reduces heap fragmentation and allows
	// returning individual substrings very efficiently.
	data, err := pstate.mapFile(in.File(), in.Offset(), int64(sLen+dLen))
	if err != nil {
		pstate.yyerror("import %q: mapping input: %v", pkg.Path, err)
		pstate.errorexit()
	}
	stringData := data[:sLen]
	declData := data[sLen:]

	in.Seek(int64(sLen+dLen), os.SEEK_CUR)

	p := &iimporter{
		ipkg: pkg,

		pkgCache:     map[uint64]*types.Pkg{},
		posBaseCache: map[uint64]*src.PosBase{},
		typCache:     map[uint64]*types.Type{},

		stringData: stringData,
		declData:   declData,
	}

	for i, pt := range pstate.predeclared() {
		p.typCache[uint64(i)] = pt
	}

	// Declaration index.
	for nPkgs := ir.uint64(pstate); nPkgs > 0; nPkgs-- {
		pkg := p.pkgAt(pstate, ir.uint64(pstate))
		pkgName := p.stringAt(pstate, ir.uint64(pstate))
		pkgHeight := int(ir.uint64(pstate))
		if pkg.Name == "" {
			pkg.Name = pkgName
			pkg.Height = pkgHeight
			pstate.numImport[pkgName]++

			// TODO(mdempsky): This belongs somewhere else.
			pkg.Lookup(pstate.types, "_").Def = asTypesNode(pstate.nblank)
		} else {
			if pkg.Name != pkgName {
				pstate.Fatalf("conflicting package names %v and %v for path %q", pkg.Name, pkgName, pkg.Path)
			}
			if pkg.Height != pkgHeight {
				pstate.Fatalf("conflicting package heights %v and %v for path %q", pkg.Height, pkgHeight, pkg.Path)
			}
		}

		for nSyms := ir.uint64(pstate); nSyms > 0; nSyms-- {
			s := pkg.Lookup(pstate.types, p.stringAt(pstate, ir.uint64(pstate)))
			off := ir.uint64(pstate)

			if _, ok := pstate.declImporter[s]; ok {
				continue
			}
			pstate.declImporter[s] = iimporterAndOffset{p, off}

			// Create stub declaration. If used, this will
			// be overwritten by expandDecl.
			if s.Def != nil {
				pstate.Fatalf("unexpected definition for %v: %v", s, asNode(s.Def))
			}
			s.Def = asTypesNode(npos(pstate.src.NoXPos, pstate.dclname(s)))
		}
	}

	// Inline body index.
	for nPkgs := ir.uint64(pstate); nPkgs > 0; nPkgs-- {
		pkg := p.pkgAt(pstate, ir.uint64(pstate))

		for nSyms := ir.uint64(pstate); nSyms > 0; nSyms-- {
			s := pkg.Lookup(pstate.types, p.stringAt(pstate, ir.uint64(pstate)))
			off := ir.uint64(pstate)

			if _, ok := pstate.inlineImporter[s]; ok {
				continue
			}
			pstate.inlineImporter[s] = iimporterAndOffset{p, off}
		}
	}
}

type iimporter struct {
	ipkg *types.Pkg

	pkgCache     map[uint64]*types.Pkg
	posBaseCache map[uint64]*src.PosBase
	typCache     map[uint64]*types.Type

	stringData string
	declData   string
}

func (p *iimporter) stringAt(pstate *PackageState, off uint64) string {
	var x [binary.MaxVarintLen64]byte
	n := copy(x[:], p.stringData[off:])

	slen, n := binary.Uvarint(x[:n])
	if n <= 0 {
		pstate.Fatalf("varint failed")
	}
	spos := off + uint64(n)
	return p.stringData[spos : spos+slen]
}

func (p *iimporter) posBaseAt(pstate *PackageState, off uint64) *src.PosBase {
	if posBase, ok := p.posBaseCache[off]; ok {
		return posBase
	}

	file := p.stringAt(pstate, off)
	posBase := src.NewFileBase(file, file)
	p.posBaseCache[off] = posBase
	return posBase
}

func (p *iimporter) pkgAt(pstate *PackageState, off uint64) *types.Pkg {
	if pkg, ok := p.pkgCache[off]; ok {
		return pkg
	}

	pkg := p.ipkg
	if pkgPath := p.stringAt(pstate, off); pkgPath != "" {
		pkg = pstate.types.NewPkg(pkgPath, "")
	}
	p.pkgCache[off] = pkg
	return pkg
}

// An importReader keeps state for reading an individual imported
// object (declaration or inline body).
type importReader struct {
	strings.Reader
	p *iimporter

	currPkg  *types.Pkg
	prevBase *src.PosBase
	prevLine int64
}

func (p *iimporter) newReader(off uint64, pkg *types.Pkg) *importReader {
	r := &importReader{
		p:       p,
		currPkg: pkg,
	}
	// (*strings.Reader).Reset wasn't added until Go 1.7, and we
	// need to build with Go 1.4.
	r.Reader = *strings.NewReader(p.declData[off:])
	return r
}

func (r *importReader) string(pstate *PackageState) string {
	return r.p.stringAt(pstate, r.uint64(pstate))
}
func (r *importReader) posBase(pstate *PackageState) *src.PosBase {
	return r.p.posBaseAt(pstate, r.uint64(pstate))
}
func (r *importReader) pkg(pstate *PackageState) *types.Pkg {
	return r.p.pkgAt(pstate, r.uint64(pstate))
}

func (r *importReader) setPkg(pstate *PackageState) {
	r.currPkg = r.pkg(pstate)
}

func (r *importReader) doDecl(pstate *PackageState, n *Node) {
	if n.Op != ONONAME {
		pstate.Fatalf("doDecl: unexpected Op for %v: %v", n.Sym, n.Op)
	}

	tag := r.byte(pstate)
	pos := r.pos(pstate)

	switch tag {
	case 'A':
		typ := r.typ(pstate)

		pstate.importalias(r.p.ipkg, pos, n.Sym, typ)

	case 'C':
		typ, val := r.value(pstate)

		pstate.importconst(r.p.ipkg, pos, n.Sym, typ, val)

	case 'F':
		typ := r.signature(pstate, nil)

		pstate.importfunc(r.p.ipkg, pos, n.Sym, typ)
		r.funcExt(pstate, n)

	case 'T':
		// Types can be recursive. We need to setup a stub
		// declaration before recursing.
		t := pstate.importtype(r.p.ipkg, pos, n.Sym)

		underlying := r.typ(pstate)
		pstate.copytype(pstate.typenod(t), underlying)

		if underlying.IsInterface() {
			break
		}

		ms := make([]*types.Field, r.uint64(pstate))
		for i := range ms {
			mpos := r.pos(pstate)
			msym := r.ident(pstate)
			recv := r.param(pstate)
			mtyp := r.signature(pstate, recv)

			f := types.NewField()
			f.Pos = mpos
			f.Sym = msym
			f.Type = mtyp
			ms[i] = f

			m := pstate.newfuncnamel(mpos, pstate.methodSym(recv.Type, msym))
			m.Type = mtyp
			m.SetClass(PFUNC)

			// (comment from parser.go)
			// inl.C's inlnode in on a dotmeth node expects to find the inlineable body as
			// (dotmeth's type).Nname.Inl, and dotmeth's type has been pulled
			// out by typecheck's lookdot as this $$.ttype. So by providing
			// this back link here we avoid special casing there.
			mtyp.SetNname(pstate.types, asTypesNode(m))
		}
		t.Methods().Set(ms)

		for _, m := range ms {
			r.methExt(pstate, m)
		}

	case 'V':
		typ := r.typ(pstate)

		pstate.importvar(r.p.ipkg, pos, n.Sym, typ)
		r.varExt(pstate, n)

	default:
		pstate.Fatalf("unexpected tag: %v", tag)
	}
}

func (p *importReader) value(pstate *PackageState) (typ *types.Type, v Val) {
	typ = p.typ(pstate)

	switch pstate.constTypeOf(typ) {
	case CTNIL:
		v.U = &NilVal{}
	case CTBOOL:
		v.U = p.bool(pstate)
	case CTSTR:
		v.U = p.string(pstate)
	case CTINT:
		x := new(Mpint)
		x.Rune = typ == pstate.types.Idealrune
		p.mpint(pstate, &x.Val, typ)
		v.U = x
	case CTFLT:
		x := newMpflt()
		p.float(pstate, x, typ)
		v.U = x
	case CTCPLX:
		x := newMpcmplx()
		p.float(pstate, &x.Real, typ)
		p.float(pstate, &x.Imag, typ)
		v.U = x
	}

	typ = pstate.idealType(typ)
	return
}

func (p *importReader) mpint(pstate *PackageState, x *big.Int, typ *types.Type) {
	signed, maxBytes := pstate.intSize(typ)

	maxSmall := 256 - maxBytes
	if signed {
		maxSmall = 256 - 2*maxBytes
	}
	if maxBytes == 1 {
		maxSmall = 256
	}

	n, _ := p.ReadByte()
	if uint(n) < maxSmall {
		v := int64(n)
		if signed {
			v >>= 1
			if n&1 != 0 {
				v = ^v
			}
		}
		x.SetInt64(v)
		return
	}

	v := -n
	if signed {
		v = -(n &^ 1) >> 1
	}
	if v < 1 || uint(v) > maxBytes {
		pstate.Fatalf("weird decoding: %v, %v => %v", n, signed, v)
	}
	b := make([]byte, v)
	p.Read(b)
	x.SetBytes(b)
	if signed && n&1 != 0 {
		x.Neg(x)
	}
}

func (p *importReader) float(pstate *PackageState, x *Mpflt, typ *types.Type) {
	var mant big.Int
	p.mpint(pstate, &mant, typ)
	m := x.Val.SetInt(&mant)
	if m.Sign() == 0 {
		return
	}
	m.SetMantExp(m, int(p.int64(pstate)))
}

func (r *importReader) ident(pstate *PackageState) *types.Sym {
	name := r.string(pstate)
	if name == "" {
		return nil
	}
	pkg := r.currPkg
	if types.IsExported(name) {
		pkg = pstate.localpkg
	}
	return pkg.Lookup(pstate.types, name)
}

func (r *importReader) qualifiedIdent(pstate *PackageState) *types.Sym {
	name := r.string(pstate)
	pkg := r.pkg(pstate)
	return pkg.Lookup(pstate.types, name)
}

func (r *importReader) pos(pstate *PackageState) src.XPos {
	delta := r.int64(pstate)
	if delta != deltaNewFile {
		r.prevLine += delta
	} else if l := r.int64(pstate); l == -1 {
		r.prevLine += deltaNewFile
	} else {
		r.prevBase = r.posBase(pstate)
		r.prevLine = l
	}

	if (r.prevBase == nil || r.prevBase.AbsFilename() == "") && r.prevLine == 0 {
		// TODO(mdempsky): Remove once we reliably write
		// position information for all nodes.
		return pstate.src.NoXPos
	}

	if r.prevBase == nil {
		pstate.Fatalf("missing posbase")
	}
	pos := src.MakePos(r.prevBase, uint(r.prevLine), 0)
	return pstate.Ctxt.PosTable.XPos(pos)
}

func (r *importReader) typ(pstate *PackageState) *types.Type {
	return r.p.typAt(pstate, r.uint64(pstate))
}

func (p *iimporter) typAt(pstate *PackageState, off uint64) *types.Type {
	t, ok := p.typCache[off]
	if !ok {
		if off < predeclReserved {
			pstate.Fatalf("predeclared type missing from cache: %d", off)
		}
		t = p.newReader(off-predeclReserved, nil).typ1(pstate)
		p.typCache[off] = t
	}
	return t
}

func (r *importReader) typ1(pstate *PackageState) *types.Type {
	switch k := r.kind(pstate); k {
	default:
		pstate.Fatalf("unexpected kind tag in %q: %v", r.p.ipkg.Path, k)
		return nil

	case definedType:
		// We might be called from within doInline, in which
		// case Sym.Def can point to declared parameters
		// instead of the top-level types. Also, we don't
		// support inlining functions with local defined
		// types. Therefore, this must be a package-scope
		// type.
		n := asNode(r.qualifiedIdent(pstate).PkgDef(pstate.types))
		if n.Op == ONONAME {
			pstate.expandDecl(n)
		}
		if n.Op != OTYPE {
			pstate.Fatalf("expected OTYPE, got %v: %v, %v", n.Op, n.Sym, n)
		}
		return n.Type
	case pointerType:
		return pstate.types.NewPtr(r.typ(pstate))
	case sliceType:
		return pstate.types.NewSlice(r.typ(pstate))
	case arrayType:
		n := r.uint64(pstate)
		return pstate.types.NewArray(r.typ(pstate), int64(n))
	case chanType:
		dir := types.ChanDir(r.uint64(pstate))
		return pstate.types.NewChan(r.typ(pstate), dir)
	case mapType:
		return pstate.types.NewMap(r.typ(pstate), r.typ(pstate))

	case signatureType:
		r.setPkg(pstate)
		return r.signature(pstate, nil)

	case structType:
		r.setPkg(pstate)

		fs := make([]*types.Field, r.uint64(pstate))
		for i := range fs {
			pos := r.pos(pstate)
			sym := r.ident(pstate)
			typ := r.typ(pstate)
			emb := r.bool(pstate)
			note := r.string(pstate)

			f := types.NewField()
			f.Pos = pos
			f.Sym = sym
			f.Type = typ
			if emb {
				f.Embedded = 1
			}
			f.Note = note
			fs[i] = f
		}

		t := types.New(TSTRUCT)
		t.SetPkg(pstate.types, r.currPkg)
		t.SetFields(pstate.types, fs)
		return t

	case interfaceType:
		r.setPkg(pstate)

		embeddeds := make([]*types.Field, r.uint64(pstate))
		for i := range embeddeds {
			pos := r.pos(pstate)
			typ := r.typ(pstate)

			f := types.NewField()
			f.Pos = pos
			f.Type = typ
			embeddeds[i] = f
		}

		methods := make([]*types.Field, r.uint64(pstate))
		for i := range methods {
			pos := r.pos(pstate)
			sym := r.ident(pstate)
			typ := r.signature(pstate, pstate.fakeRecvField())

			f := types.NewField()
			f.Pos = pos
			f.Sym = sym
			f.Type = typ
			methods[i] = f
		}

		t := types.New(TINTER)
		t.SetPkg(pstate.types, r.currPkg)
		t.SetInterface(pstate.types, append(embeddeds, methods...))
		return t
	}
}

func (r *importReader) kind(pstate *PackageState) itag {
	return itag(r.uint64(pstate))
}

func (r *importReader) signature(pstate *PackageState, recv *types.Field) *types.Type {
	params := r.paramList(pstate)
	results := r.paramList(pstate)
	if n := len(params); n > 0 {
		params[n-1].SetIsddd(r.bool(pstate))
	}
	t := pstate.functypefield(recv, params, results)
	t.SetPkg(pstate.types, r.currPkg)
	return t
}

func (r *importReader) paramList(pstate *PackageState) []*types.Field {
	fs := make([]*types.Field, r.uint64(pstate))
	for i := range fs {
		fs[i] = r.param(pstate)
	}
	return fs
}

func (r *importReader) param(pstate *PackageState) *types.Field {
	f := types.NewField()
	f.Pos = r.pos(pstate)
	f.Sym = r.ident(pstate)
	f.Type = r.typ(pstate)
	return f
}

func (r *importReader) bool(pstate *PackageState) bool {
	return r.uint64(pstate) != 0
}

func (r *importReader) int64(pstate *PackageState) int64 {
	n, err := binary.ReadVarint(r)
	if err != nil {
		pstate.Fatalf("readVarint: %v", err)
	}
	return n
}

func (r *importReader) uint64(pstate *PackageState) uint64 {
	n, err := binary.ReadUvarint(r)
	if err != nil {
		pstate.Fatalf("readVarint: %v", err)
	}
	return n
}

func (r *importReader) byte(pstate *PackageState) byte {
	x, err := r.ReadByte()
	if err != nil {
		pstate.Fatalf("declReader.ReadByte: %v", err)
	}
	return x
}

// Compiler-specific extensions.

func (r *importReader) varExt(pstate *PackageState, n *Node) {
	r.linkname(pstate, n.Sym)
}

func (r *importReader) funcExt(pstate *PackageState, n *Node) {
	r.linkname(pstate, n.Sym)

	// Escape analysis.
	for _, fs := range pstate.types.RecvsParams {
		for _, f := range fs(n.Type).FieldSlice(pstate.types) {
			f.Note = r.string(pstate)
		}
	}

	// Inline body.
	if u := r.uint64(pstate); u > 0 {
		n.Func.Inl = &Inline{
			Cost: int32(u - 1),
		}
	}
}

func (r *importReader) methExt(pstate *PackageState, m *types.Field) {
	if r.bool(pstate) {
		m.SetNointerface(true)
	}
	r.funcExt(pstate, asNode(m.Type.Nname(pstate.types)))
}

func (r *importReader) linkname(pstate *PackageState, s *types.Sym) {
	s.Linkname = r.string(pstate)
}

func (r *importReader) doInline(pstate *PackageState, n *Node) {
	if len(n.Func.Inl.Body) != 0 {
		pstate.Fatalf("%v already has inline body", n)
	}

	pstate.funchdr(n)
	body := r.stmtList(pstate)
	pstate.funcbody()
	if body == nil {
		//
		// Make sure empty body is not interpreted as
		// no inlineable body (see also parser.fnbody)
		// (not doing so can cause significant performance
		// degradation due to unnecessary calls to empty
		// functions).
		body = []*Node{}
	}
	n.Func.Inl.Body = body

	pstate.importlist = append(pstate.importlist, n)

	if pstate.Debug['E'] > 0 && pstate.Debug['m'] > 2 {
		if pstate.Debug['m'] > 3 {
			fmt.Printf("inl body for %v %#v: %+v\n", n, n.Type, asNodes(n.Func.Inl.Body))
		} else {
			fmt.Printf("inl body for %v %#v: %v\n", n, n.Type, asNodes(n.Func.Inl.Body))
		}
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

func (r *importReader) stmtList(pstate *PackageState) []*Node {
	var list []*Node
	for {
		n := r.node(pstate)
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

func (r *importReader) exprList(pstate *PackageState) []*Node {
	var list []*Node
	for {
		n := r.expr(pstate)
		if n == nil {
			break
		}
		list = append(list, n)
	}
	return list
}

func (r *importReader) expr(pstate *PackageState) *Node {
	n := r.node(pstate)
	if n != nil && n.Op == OBLOCK {
		pstate.Fatalf("unexpected block node: %v", n)
	}
	return n
}

// TODO(gri) split into expr and stmt
func (r *importReader) node(pstate *PackageState) *Node {
	switch op := r.op(pstate); op {
	// expressions
	// case OPAREN:
	// 	unreachable - unpacked by exporter

	// case ODDDARG:
	//	unimplemented

	case OLITERAL:
		pos := r.pos(pstate)
		typ, val := r.value(pstate)

		n := npos(pos, pstate.nodlit(val))
		n.Type = typ
		return n

	case ONONAME:
		return pstate.mkname(r.qualifiedIdent(pstate))

	case ONAME:
		return pstate.mkname(r.ident(pstate))

	// case OPACK, ONONAME:
	// 	unreachable - should have been resolved by typechecking

	case OTYPE:
		return pstate.typenod(r.typ(pstate))

	// case OTARRAY, OTMAP, OTCHAN, OTSTRUCT, OTINTER, OTFUNC:
	//      unreachable - should have been resolved by typechecking

	// case OCLOSURE:
	//	unimplemented

	case OPTRLIT:
		pos := r.pos(pstate)
		n := npos(pos, r.expr(pstate))
		if !r.bool(pstate) /* !implicit, i.e. '&' operator */ {
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
		pstate.lineno = r.pos(pstate)
		n := pstate.nodl(pstate.lineno, OCOMPLIT, nil, pstate.typenod(r.typ(pstate)))
		n.List.Set(r.elemList(pstate)) // special handling of field names
		pstate.lineno = savedlineno
		return n

	// case OARRAYLIT, OSLICELIT, OMAPLIT:
	// 	unreachable - mapped to case OCOMPLIT below by exporter

	case OCOMPLIT:
		n := pstate.nodl(r.pos(pstate), OCOMPLIT, nil, pstate.typenod(r.typ(pstate)))
		n.List.Set(r.exprList(pstate))
		return n

	case OKEY:
		pos := r.pos(pstate)
		left, right := r.exprsOrNil(pstate)
		return pstate.nodl(pos, OKEY, left, right)

	// case OSTRUCTKEY:
	//	unreachable - handled in case OSTRUCTLIT by elemList

	// case OCALLPART:
	//	unimplemented

	// case OXDOT, ODOT, ODOTPTR, ODOTINTER, ODOTMETH:
	// 	unreachable - mapped to case OXDOT below by exporter

	case OXDOT:
		// see parser.new_dotname
		return npos(r.pos(pstate), pstate.nodSym(OXDOT, r.expr(pstate), r.ident(pstate)))

	// case ODOTTYPE, ODOTTYPE2:
	// 	unreachable - mapped to case ODOTTYPE below by exporter

	case ODOTTYPE:
		n := pstate.nodl(r.pos(pstate), ODOTTYPE, r.expr(pstate), nil)
		n.Type = r.typ(pstate)
		return n

	// case OINDEX, OINDEXMAP, OSLICE, OSLICESTR, OSLICEARR, OSLICE3, OSLICE3ARR:
	// 	unreachable - mapped to cases below by exporter

	case OINDEX:
		return pstate.nodl(r.pos(pstate), op, r.expr(pstate), r.expr(pstate))

	case OSLICE, OSLICE3:
		n := pstate.nodl(r.pos(pstate), op, r.expr(pstate), nil)
		low, high := r.exprsOrNil(pstate)
		var max *Node
		if n.Op.IsSlice3(pstate) {
			max = r.expr(pstate)
		}
		n.SetSliceBounds(pstate, low, high, max)
		return n

	// case OCONV, OCONVIFACE, OCONVNOP, OARRAYBYTESTR, OARRAYRUNESTR, OSTRARRAYBYTE, OSTRARRAYRUNE, ORUNESTR:
	// 	unreachable - mapped to OCONV case below by exporter

	case OCONV:
		n := pstate.nodl(r.pos(pstate), OCONV, r.expr(pstate), nil)
		n.Type = r.typ(pstate)
		return n

	case OCOPY, OCOMPLEX, OREAL, OIMAG, OAPPEND, OCAP, OCLOSE, ODELETE, OLEN, OMAKE, ONEW, OPANIC, ORECOVER, OPRINT, OPRINTN:
		n := npos(r.pos(pstate), pstate.builtinCall(op))
		n.List.Set(r.exprList(pstate))
		if op == OAPPEND {
			n.SetIsddd(r.bool(pstate))
		}
		return n

	// case OCALL, OCALLFUNC, OCALLMETH, OCALLINTER, OGETG:
	// 	unreachable - mapped to OCALL case below by exporter

	case OCALL:
		n := pstate.nodl(r.pos(pstate), OCALL, r.expr(pstate), nil)
		n.List.Set(r.exprList(pstate))
		n.SetIsddd(r.bool(pstate))
		return n

	case OMAKEMAP, OMAKECHAN, OMAKESLICE:
		n := npos(r.pos(pstate), pstate.builtinCall(OMAKE))
		n.List.Append(pstate.typenod(r.typ(pstate)))
		n.List.Append(r.exprList(pstate)...)
		return n

	// unary expressions
	case OPLUS, OMINUS, OADDR, OCOM, OIND, ONOT, ORECV:
		return pstate.nodl(r.pos(pstate), op, r.expr(pstate), nil)

	// binary expressions
	case OADD, OAND, OANDAND, OANDNOT, ODIV, OEQ, OGE, OGT, OLE, OLT,
		OLSH, OMOD, OMUL, ONE, OOR, OOROR, ORSH, OSEND, OSUB, OXOR:
		return pstate.nodl(r.pos(pstate), op, r.expr(pstate), r.expr(pstate))

	case OADDSTR:
		pos := r.pos(pstate)
		list := r.exprList(pstate)
		x := npos(pos, list[0])
		for _, y := range list[1:] {
			x = pstate.nodl(pos, OADD, x, y)
		}
		return x

	// case OCMPSTR, OCMPIFACE:
	// 	unreachable - mapped to std comparison operators by exporter

	// --------------------------------------------------------------------
	// statements
	case ODCL:
		pos := r.pos(pstate)
		lhs := npos(pos, pstate.dclname(r.ident(pstate)))
		typ := pstate.typenod(r.typ(pstate))
		return npos(pos, pstate.liststmt(pstate.variter([]*Node{lhs}, typ, nil))) // TODO(gri) avoid list creation

	// case ODCLFIELD:
	//	unimplemented

	// case OAS, OASWB:
	// 	unreachable - mapped to OAS case below by exporter

	case OAS:
		return pstate.nodl(r.pos(pstate), OAS, r.expr(pstate), r.expr(pstate))

	case OASOP:
		n := pstate.nodl(r.pos(pstate), OASOP, nil, nil)
		n.SetSubOp(pstate, r.op(pstate))
		n.Left = r.expr(pstate)
		if !r.bool(pstate) {
			n.Right = pstate.nodintconst(1)
			n.SetImplicit(true)
		} else {
			n.Right = r.expr(pstate)
		}
		return n

	// case OAS2DOTTYPE, OAS2FUNC, OAS2MAPR, OAS2RECV:
	// 	unreachable - mapped to OAS2 case below by exporter

	case OAS2:
		n := pstate.nodl(r.pos(pstate), OAS2, nil, nil)
		n.List.Set(r.exprList(pstate))
		n.Rlist.Set(r.exprList(pstate))
		return n

	case ORETURN:
		n := pstate.nodl(r.pos(pstate), ORETURN, nil, nil)
		n.List.Set(r.exprList(pstate))
		return n

	// case ORETJMP:
	// 	unreachable - generated by compiler for trampolin routines (not exported)

	case OPROC, ODEFER:
		return pstate.nodl(r.pos(pstate), op, r.expr(pstate), nil)

	case OIF:
		n := pstate.nodl(r.pos(pstate), OIF, nil, nil)
		n.Ninit.Set(r.stmtList(pstate))
		n.Left = r.expr(pstate)
		n.Nbody.Set(r.stmtList(pstate))
		n.Rlist.Set(r.stmtList(pstate))
		return n

	case OFOR:
		n := pstate.nodl(r.pos(pstate), OFOR, nil, nil)
		n.Ninit.Set(r.stmtList(pstate))
		n.Left, n.Right = r.exprsOrNil(pstate)
		n.Nbody.Set(r.stmtList(pstate))
		return n

	case ORANGE:
		n := pstate.nodl(r.pos(pstate), ORANGE, nil, nil)
		n.List.Set(r.stmtList(pstate))
		n.Right = r.expr(pstate)
		n.Nbody.Set(r.stmtList(pstate))
		return n

	case OSELECT, OSWITCH:
		n := pstate.nodl(r.pos(pstate), op, nil, nil)
		n.Ninit.Set(r.stmtList(pstate))
		n.Left, _ = r.exprsOrNil(pstate)
		n.List.Set(r.stmtList(pstate))
		return n

	// case OCASE, OXCASE:
	// 	unreachable - mapped to OXCASE case below by exporter

	case OXCASE:
		n := pstate.nodl(r.pos(pstate), OXCASE, nil, nil)
		n.List.Set(r.exprList(pstate))
		// TODO(gri) eventually we must declare variables for type switch
		// statements (type switch statements are not yet exported)
		n.Nbody.Set(r.stmtList(pstate))
		return n

	// case OFALL:
	// 	unreachable - mapped to OXFALL case below by exporter

	case OFALL:
		n := pstate.nodl(r.pos(pstate), OFALL, nil, nil)
		return n

	case OBREAK, OCONTINUE:
		pos := r.pos(pstate)
		left, _ := r.exprsOrNil(pstate)
		if left != nil {
			left = pstate.newname(left.Sym)
		}
		return pstate.nodl(pos, op, left, nil)

	// case OEMPTY:
	// 	unreachable - not emitted by exporter

	case OGOTO, OLABEL:
		return pstate.nodl(r.pos(pstate), op, pstate.newname(r.expr(pstate).Sym), nil)

	case OEND:
		return nil

	default:
		pstate.Fatalf("cannot import %v (%d) node\n"+
			"==> please file an issue and assign to gri@\n", op, int(op))
		panic("unreachable") // satisfy compiler
	}
}

func (r *importReader) op(pstate *PackageState) Op {
	return Op(r.uint64(pstate))
}

func (r *importReader) elemList(pstate *PackageState) []*Node {
	c := r.uint64(pstate)
	list := make([]*Node, c)
	for i := range list {
		s := r.ident(pstate)
		list[i] = pstate.nodSym(OSTRUCTKEY, r.expr(pstate), s)
	}
	return list
}

func (r *importReader) exprsOrNil(pstate *PackageState) (a, b *Node) {
	ab := r.uint64(pstate)
	if ab&1 != 0 {
		a = r.expr(pstate)
	}
	if ab&2 != 0 {
		b = r.node(pstate)
	}
	return
}

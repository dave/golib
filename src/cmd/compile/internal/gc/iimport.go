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

// declImporter maps from imported identifiers to an importer
// and offset where that identifier's declaration can be read.

// inlineImporter is like declImporter, but for inline bodies
// for function and method symbols.

func (psess *PackageSession) expandDecl(n *Node) {
	if n.Op != ONONAME {
		return
	}

	r := importReaderFor(n, psess.declImporter)
	if r == nil {

		return
	}

	r.doDecl(psess, n)
}

func (psess *PackageSession) expandInline(fn *Node) {
	if fn.Func.Inl.Body != nil {
		return
	}

	r := importReaderFor(fn, psess.inlineImporter)
	if r == nil {
		psess.
			Fatalf("missing import reader for %v", fn)
	}

	r.doInline(psess, fn)
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

func (r *intReader) int64(psess *PackageSession) int64 {
	i, err := binary.ReadVarint(r.Reader)
	if err != nil {
		psess.
			yyerror("import %q: read error: %v", r.pkg.Path, err)
		psess.
			errorexit()
	}
	return i
}

func (r *intReader) uint64(psess *PackageSession) uint64 {
	i, err := binary.ReadUvarint(r.Reader)
	if err != nil {
		psess.
			yyerror("import %q: read error: %v", r.pkg.Path, err)
		psess.
			errorexit()
	}
	return i
}

func (psess *PackageSession) iimport(pkg *types.Pkg, in *bio.Reader) {
	ir := &intReader{in, pkg}

	version := ir.uint64(psess)
	if version != iexportVersion {
		psess.
			yyerror("import %q: unknown export format version %d", pkg.Path, version)
		psess.
			errorexit()
	}

	sLen := ir.uint64(psess)
	dLen := ir.uint64(psess)

	data, err := psess.mapFile(in.File(), in.Offset(), int64(sLen+dLen))
	if err != nil {
		psess.
			yyerror("import %q: mapping input: %v", pkg.Path, err)
		psess.
			errorexit()
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

	for i, pt := range psess.predeclared() {
		p.typCache[uint64(i)] = pt
	}

	for nPkgs := ir.uint64(psess); nPkgs > 0; nPkgs-- {
		pkg := p.pkgAt(psess, ir.uint64(psess))
		pkgName := p.stringAt(psess, ir.uint64(psess))
		pkgHeight := int(ir.uint64(psess))
		if pkg.Name == "" {
			pkg.Name = pkgName
			pkg.Height = pkgHeight
			psess.
				numImport[pkgName]++

			pkg.Lookup(psess.types, "_").Def = asTypesNode(psess.nblank)
		} else {
			if pkg.Name != pkgName {
				psess.
					Fatalf("conflicting package names %v and %v for path %q", pkg.Name, pkgName, pkg.Path)
			}
			if pkg.Height != pkgHeight {
				psess.
					Fatalf("conflicting package heights %v and %v for path %q", pkg.Height, pkgHeight, pkg.Path)
			}
		}

		for nSyms := ir.uint64(psess); nSyms > 0; nSyms-- {
			s := pkg.Lookup(psess.types, p.stringAt(psess, ir.uint64(psess)))
			off := ir.uint64(psess)

			if _, ok := psess.declImporter[s]; ok {
				continue
			}
			psess.
				declImporter[s] = iimporterAndOffset{p, off}

			if s.Def != nil {
				psess.
					Fatalf("unexpected definition for %v: %v", s, asNode(s.Def))
			}
			s.Def = asTypesNode(npos(psess.src.NoXPos, psess.dclname(s)))
		}
	}

	for nPkgs := ir.uint64(psess); nPkgs > 0; nPkgs-- {
		pkg := p.pkgAt(psess, ir.uint64(psess))

		for nSyms := ir.uint64(psess); nSyms > 0; nSyms-- {
			s := pkg.Lookup(psess.types, p.stringAt(psess, ir.uint64(psess)))
			off := ir.uint64(psess)

			if _, ok := psess.inlineImporter[s]; ok {
				continue
			}
			psess.
				inlineImporter[s] = iimporterAndOffset{p, off}
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

func (p *iimporter) stringAt(psess *PackageSession, off uint64) string {
	var x [binary.MaxVarintLen64]byte
	n := copy(x[:], p.stringData[off:])

	slen, n := binary.Uvarint(x[:n])
	if n <= 0 {
		psess.
			Fatalf("varint failed")
	}
	spos := off + uint64(n)
	return p.stringData[spos : spos+slen]
}

func (p *iimporter) posBaseAt(psess *PackageSession, off uint64) *src.PosBase {
	if posBase, ok := p.posBaseCache[off]; ok {
		return posBase
	}

	file := p.stringAt(psess, off)
	posBase := src.NewFileBase(file, file)
	p.posBaseCache[off] = posBase
	return posBase
}

func (p *iimporter) pkgAt(psess *PackageSession, off uint64) *types.Pkg {
	if pkg, ok := p.pkgCache[off]; ok {
		return pkg
	}

	pkg := p.ipkg
	if pkgPath := p.stringAt(psess, off); pkgPath != "" {
		pkg = psess.types.NewPkg(pkgPath, "")
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

	r.Reader = *strings.NewReader(p.declData[off:])
	return r
}

func (r *importReader) string(psess *PackageSession) string {
	return r.p.stringAt(psess, r.uint64(psess))
}
func (r *importReader) posBase(psess *PackageSession) *src.PosBase {
	return r.p.posBaseAt(psess, r.uint64(psess))
}
func (r *importReader) pkg(psess *PackageSession) *types.Pkg {
	return r.p.pkgAt(psess, r.uint64(psess))
}

func (r *importReader) setPkg(psess *PackageSession) {
	r.currPkg = r.pkg(psess)
}

func (r *importReader) doDecl(psess *PackageSession, n *Node) {
	if n.Op != ONONAME {
		psess.
			Fatalf("doDecl: unexpected Op for %v: %v", n.Sym, n.Op)
	}

	tag := r.byte(psess)
	pos := r.pos(psess)

	switch tag {
	case 'A':
		typ := r.typ(psess)
		psess.
			importalias(r.p.ipkg, pos, n.Sym, typ)

	case 'C':
		typ, val := r.value(psess)
		psess.
			importconst(r.p.ipkg, pos, n.Sym, typ, val)

	case 'F':
		typ := r.signature(psess, nil)
		psess.
			importfunc(r.p.ipkg, pos, n.Sym, typ)
		r.funcExt(psess, n)

	case 'T':

		t := psess.importtype(r.p.ipkg, pos, n.Sym)

		underlying := r.typ(psess)
		psess.
			copytype(psess.typenod(t), underlying)

		if underlying.IsInterface() {
			break
		}

		ms := make([]*types.Field, r.uint64(psess))
		for i := range ms {
			mpos := r.pos(psess)
			msym := r.ident(psess)
			recv := r.param(psess)
			mtyp := r.signature(psess, recv)

			f := types.NewField()
			f.Pos = mpos
			f.Sym = msym
			f.Type = mtyp
			ms[i] = f

			m := psess.newfuncnamel(mpos, psess.methodSym(recv.Type, msym))
			m.Type = mtyp
			m.SetClass(PFUNC)

			mtyp.SetNname(psess.types, asTypesNode(m))
		}
		t.Methods().Set(ms)

		for _, m := range ms {
			r.methExt(psess, m)
		}

	case 'V':
		typ := r.typ(psess)
		psess.
			importvar(r.p.ipkg, pos, n.Sym, typ)
		r.varExt(psess, n)

	default:
		psess.
			Fatalf("unexpected tag: %v", tag)
	}
}

func (p *importReader) value(psess *PackageSession) (typ *types.Type, v Val) {
	typ = p.typ(psess)

	switch psess.constTypeOf(typ) {
	case CTNIL:
		v.U = &NilVal{}
	case CTBOOL:
		v.U = p.bool(psess)
	case CTSTR:
		v.U = p.string(psess)
	case CTINT:
		x := new(Mpint)
		x.Rune = typ == psess.types.Idealrune
		p.mpint(psess, &x.Val, typ)
		v.U = x
	case CTFLT:
		x := newMpflt()
		p.float(psess, x, typ)
		v.U = x
	case CTCPLX:
		x := newMpcmplx()
		p.float(psess, &x.Real, typ)
		p.float(psess, &x.Imag, typ)
		v.U = x
	}

	typ = psess.idealType(typ)
	return
}

func (p *importReader) mpint(psess *PackageSession, x *big.Int, typ *types.Type) {
	signed, maxBytes := psess.intSize(typ)

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
		psess.
			Fatalf("weird decoding: %v, %v => %v", n, signed, v)
	}
	b := make([]byte, v)
	p.Read(b)
	x.SetBytes(b)
	if signed && n&1 != 0 {
		x.Neg(x)
	}
}

func (p *importReader) float(psess *PackageSession, x *Mpflt, typ *types.Type) {
	var mant big.Int
	p.mpint(psess, &mant, typ)
	m := x.Val.SetInt(&mant)
	if m.Sign() == 0 {
		return
	}
	m.SetMantExp(m, int(p.int64(psess)))
}

func (r *importReader) ident(psess *PackageSession) *types.Sym {
	name := r.string(psess)
	if name == "" {
		return nil
	}
	pkg := r.currPkg
	if types.IsExported(name) {
		pkg = psess.localpkg
	}
	return pkg.Lookup(psess.types, name)
}

func (r *importReader) qualifiedIdent(psess *PackageSession) *types.Sym {
	name := r.string(psess)
	pkg := r.pkg(psess)
	return pkg.Lookup(psess.types, name)
}

func (r *importReader) pos(psess *PackageSession) src.XPos {
	delta := r.int64(psess)
	if delta != deltaNewFile {
		r.prevLine += delta
	} else if l := r.int64(psess); l == -1 {
		r.prevLine += deltaNewFile
	} else {
		r.prevBase = r.posBase(psess)
		r.prevLine = l
	}

	if (r.prevBase == nil || r.prevBase.AbsFilename() == "") && r.prevLine == 0 {

		return psess.src.NoXPos
	}

	if r.prevBase == nil {
		psess.
			Fatalf("missing posbase")
	}
	pos := src.MakePos(r.prevBase, uint(r.prevLine), 0)
	return psess.Ctxt.PosTable.XPos(pos)
}

func (r *importReader) typ(psess *PackageSession) *types.Type {
	return r.p.typAt(psess, r.uint64(psess))
}

func (p *iimporter) typAt(psess *PackageSession, off uint64) *types.Type {
	t, ok := p.typCache[off]
	if !ok {
		if off < predeclReserved {
			psess.
				Fatalf("predeclared type missing from cache: %d", off)
		}
		t = p.newReader(off-predeclReserved, nil).typ1(psess)
		p.typCache[off] = t
	}
	return t
}

func (r *importReader) typ1(psess *PackageSession) *types.Type {
	switch k := r.kind(psess); k {
	default:
		psess.
			Fatalf("unexpected kind tag in %q: %v", r.p.ipkg.Path, k)
		return nil

	case definedType:

		n := asNode(r.qualifiedIdent(psess).PkgDef(psess.types))
		if n.Op == ONONAME {
			psess.
				expandDecl(n)
		}
		if n.Op != OTYPE {
			psess.
				Fatalf("expected OTYPE, got %v: %v, %v", n.Op, n.Sym, n)
		}
		return n.Type
	case pointerType:
		return psess.types.NewPtr(r.typ(psess))
	case sliceType:
		return psess.types.NewSlice(r.typ(psess))
	case arrayType:
		n := r.uint64(psess)
		return psess.types.NewArray(r.typ(psess), int64(n))
	case chanType:
		dir := types.ChanDir(r.uint64(psess))
		return psess.types.NewChan(r.typ(psess), dir)
	case mapType:
		return psess.types.NewMap(r.typ(psess), r.typ(psess))

	case signatureType:
		r.setPkg(psess)
		return r.signature(psess, nil)

	case structType:
		r.setPkg(psess)

		fs := make([]*types.Field, r.uint64(psess))
		for i := range fs {
			pos := r.pos(psess)
			sym := r.ident(psess)
			typ := r.typ(psess)
			emb := r.bool(psess)
			note := r.string(psess)

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
		t.SetPkg(psess.types, r.currPkg)
		t.SetFields(psess.types, fs)
		return t

	case interfaceType:
		r.setPkg(psess)

		embeddeds := make([]*types.Field, r.uint64(psess))
		for i := range embeddeds {
			pos := r.pos(psess)
			typ := r.typ(psess)

			f := types.NewField()
			f.Pos = pos
			f.Type = typ
			embeddeds[i] = f
		}

		methods := make([]*types.Field, r.uint64(psess))
		for i := range methods {
			pos := r.pos(psess)
			sym := r.ident(psess)
			typ := r.signature(psess, psess.fakeRecvField())

			f := types.NewField()
			f.Pos = pos
			f.Sym = sym
			f.Type = typ
			methods[i] = f
		}

		t := types.New(TINTER)
		t.SetPkg(psess.types, r.currPkg)
		t.SetInterface(psess.types, append(embeddeds, methods...))
		return t
	}
}

func (r *importReader) kind(psess *PackageSession) itag {
	return itag(r.uint64(psess))
}

func (r *importReader) signature(psess *PackageSession, recv *types.Field) *types.Type {
	params := r.paramList(psess)
	results := r.paramList(psess)
	if n := len(params); n > 0 {
		params[n-1].SetIsddd(r.bool(psess))
	}
	t := psess.functypefield(recv, params, results)
	t.SetPkg(psess.types, r.currPkg)
	return t
}

func (r *importReader) paramList(psess *PackageSession) []*types.Field {
	fs := make([]*types.Field, r.uint64(psess))
	for i := range fs {
		fs[i] = r.param(psess)
	}
	return fs
}

func (r *importReader) param(psess *PackageSession) *types.Field {
	f := types.NewField()
	f.Pos = r.pos(psess)
	f.Sym = r.ident(psess)
	f.Type = r.typ(psess)
	return f
}

func (r *importReader) bool(psess *PackageSession) bool {
	return r.uint64(psess) != 0
}

func (r *importReader) int64(psess *PackageSession) int64 {
	n, err := binary.ReadVarint(r)
	if err != nil {
		psess.
			Fatalf("readVarint: %v", err)
	}
	return n
}

func (r *importReader) uint64(psess *PackageSession) uint64 {
	n, err := binary.ReadUvarint(r)
	if err != nil {
		psess.
			Fatalf("readVarint: %v", err)
	}
	return n
}

func (r *importReader) byte(psess *PackageSession) byte {
	x, err := r.ReadByte()
	if err != nil {
		psess.
			Fatalf("declReader.ReadByte: %v", err)
	}
	return x
}

func (r *importReader) varExt(psess *PackageSession, n *Node) {
	r.linkname(psess, n.Sym)
}

func (r *importReader) funcExt(psess *PackageSession, n *Node) {
	r.linkname(psess, n.Sym)

	for _, fs := range psess.types.RecvsParams {
		for _, f := range fs(n.Type).FieldSlice(psess.types) {
			f.Note = r.string(psess)
		}
	}

	if u := r.uint64(psess); u > 0 {
		n.Func.Inl = &Inline{
			Cost: int32(u - 1),
		}
	}
}

func (r *importReader) methExt(psess *PackageSession, m *types.Field) {
	if r.bool(psess) {
		m.SetNointerface(true)
	}
	r.funcExt(psess, asNode(m.Type.Nname(psess.types)))
}

func (r *importReader) linkname(psess *PackageSession, s *types.Sym) {
	s.Linkname = r.string(psess)
}

func (r *importReader) doInline(psess *PackageSession, n *Node) {
	if len(n.Func.Inl.Body) != 0 {
		psess.
			Fatalf("%v already has inline body", n)
	}
	psess.
		funchdr(n)
	body := r.stmtList(psess)
	psess.
		funcbody()
	if body == nil {

		body = []*Node{}
	}
	n.Func.Inl.Body = body
	psess.
		importlist = append(psess.importlist, n)

	if psess.Debug['E'] > 0 && psess.Debug['m'] > 2 {
		if psess.Debug['m'] > 3 {
			fmt.Printf("inl body for %v %#v: %+v\n", n, n.Type, asNodes(n.Func.Inl.Body))
		} else {
			fmt.Printf("inl body for %v %#v: %v\n", n, n.Type, asNodes(n.Func.Inl.Body))
		}
	}
}

func (r *importReader) stmtList(psess *PackageSession) []*Node {
	var list []*Node
	for {
		n := r.node(psess)
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

func (r *importReader) exprList(psess *PackageSession) []*Node {
	var list []*Node
	for {
		n := r.expr(psess)
		if n == nil {
			break
		}
		list = append(list, n)
	}
	return list
}

func (r *importReader) expr(psess *PackageSession) *Node {
	n := r.node(psess)
	if n != nil && n.Op == OBLOCK {
		psess.
			Fatalf("unexpected block node: %v", n)
	}
	return n
}

// TODO(gri) split into expr and stmt
func (r *importReader) node(psess *PackageSession) *Node {
	switch op := r.op(psess); op {

	case OLITERAL:
		pos := r.pos(psess)
		typ, val := r.value(psess)

		n := npos(pos, psess.nodlit(val))
		n.Type = typ
		return n

	case ONONAME:
		return psess.mkname(r.qualifiedIdent(psess))

	case ONAME:
		return psess.mkname(r.ident(psess))

	case OTYPE:
		return psess.typenod(r.typ(psess))

	case OPTRLIT:
		pos := r.pos(psess)
		n := npos(pos, r.expr(psess))
		if !r.bool(psess) {
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
			lineno = r.pos(psess)
		n := psess.nodl(psess.lineno, OCOMPLIT, nil, psess.typenod(r.typ(psess)))
		n.List.Set(r.elemList(psess))
		psess.
			lineno = savedlineno
		return n

	case OCOMPLIT:
		n := psess.nodl(r.pos(psess), OCOMPLIT, nil, psess.typenod(r.typ(psess)))
		n.List.Set(r.exprList(psess))
		return n

	case OKEY:
		pos := r.pos(psess)
		left, right := r.exprsOrNil(psess)
		return psess.nodl(pos, OKEY, left, right)

	case OXDOT:

		return npos(r.pos(psess), psess.nodSym(OXDOT, r.expr(psess), r.ident(psess)))

	case ODOTTYPE:
		n := psess.nodl(r.pos(psess), ODOTTYPE, r.expr(psess), nil)
		n.Type = r.typ(psess)
		return n

	case OINDEX:
		return psess.nodl(r.pos(psess), op, r.expr(psess), r.expr(psess))

	case OSLICE, OSLICE3:
		n := psess.nodl(r.pos(psess), op, r.expr(psess), nil)
		low, high := r.exprsOrNil(psess)
		var max *Node
		if n.Op.IsSlice3(psess) {
			max = r.expr(psess)
		}
		n.SetSliceBounds(psess, low, high, max)
		return n

	case OCONV:
		n := psess.nodl(r.pos(psess), OCONV, r.expr(psess), nil)
		n.Type = r.typ(psess)
		return n

	case OCOPY, OCOMPLEX, OREAL, OIMAG, OAPPEND, OCAP, OCLOSE, ODELETE, OLEN, OMAKE, ONEW, OPANIC, ORECOVER, OPRINT, OPRINTN:
		n := npos(r.pos(psess), psess.builtinCall(op))
		n.List.Set(r.exprList(psess))
		if op == OAPPEND {
			n.SetIsddd(r.bool(psess))
		}
		return n

	case OCALL:
		n := psess.nodl(r.pos(psess), OCALL, r.expr(psess), nil)
		n.List.Set(r.exprList(psess))
		n.SetIsddd(r.bool(psess))
		return n

	case OMAKEMAP, OMAKECHAN, OMAKESLICE:
		n := npos(r.pos(psess), psess.builtinCall(OMAKE))
		n.List.Append(psess.typenod(r.typ(psess)))
		n.List.Append(r.exprList(psess)...)
		return n

	case OPLUS, OMINUS, OADDR, OCOM, OIND, ONOT, ORECV:
		return psess.nodl(r.pos(psess), op, r.expr(psess), nil)

	case OADD, OAND, OANDAND, OANDNOT, ODIV, OEQ, OGE, OGT, OLE, OLT,
		OLSH, OMOD, OMUL, ONE, OOR, OOROR, ORSH, OSEND, OSUB, OXOR:
		return psess.nodl(r.pos(psess), op, r.expr(psess), r.expr(psess))

	case OADDSTR:
		pos := r.pos(psess)
		list := r.exprList(psess)
		x := npos(pos, list[0])
		for _, y := range list[1:] {
			x = psess.nodl(pos, OADD, x, y)
		}
		return x

	case ODCL:
		pos := r.pos(psess)
		lhs := npos(pos, psess.dclname(r.ident(psess)))
		typ := psess.typenod(r.typ(psess))
		return npos(pos, psess.liststmt(psess.variter([]*Node{lhs}, typ, nil)))

	case OAS:
		return psess.nodl(r.pos(psess), OAS, r.expr(psess), r.expr(psess))

	case OASOP:
		n := psess.nodl(r.pos(psess), OASOP, nil, nil)
		n.SetSubOp(psess, r.op(psess))
		n.Left = r.expr(psess)
		if !r.bool(psess) {
			n.Right = psess.nodintconst(1)
			n.SetImplicit(true)
		} else {
			n.Right = r.expr(psess)
		}
		return n

	case OAS2:
		n := psess.nodl(r.pos(psess), OAS2, nil, nil)
		n.List.Set(r.exprList(psess))
		n.Rlist.Set(r.exprList(psess))
		return n

	case ORETURN:
		n := psess.nodl(r.pos(psess), ORETURN, nil, nil)
		n.List.Set(r.exprList(psess))
		return n

	case OPROC, ODEFER:
		return psess.nodl(r.pos(psess), op, r.expr(psess), nil)

	case OIF:
		n := psess.nodl(r.pos(psess), OIF, nil, nil)
		n.Ninit.Set(r.stmtList(psess))
		n.Left = r.expr(psess)
		n.Nbody.Set(r.stmtList(psess))
		n.Rlist.Set(r.stmtList(psess))
		return n

	case OFOR:
		n := psess.nodl(r.pos(psess), OFOR, nil, nil)
		n.Ninit.Set(r.stmtList(psess))
		n.Left, n.Right = r.exprsOrNil(psess)
		n.Nbody.Set(r.stmtList(psess))
		return n

	case ORANGE:
		n := psess.nodl(r.pos(psess), ORANGE, nil, nil)
		n.List.Set(r.stmtList(psess))
		n.Right = r.expr(psess)
		n.Nbody.Set(r.stmtList(psess))
		return n

	case OSELECT, OSWITCH:
		n := psess.nodl(r.pos(psess), op, nil, nil)
		n.Ninit.Set(r.stmtList(psess))
		n.Left, _ = r.exprsOrNil(psess)
		n.List.Set(r.stmtList(psess))
		return n

	case OXCASE:
		n := psess.nodl(r.pos(psess), OXCASE, nil, nil)
		n.List.Set(r.exprList(psess))

		n.Nbody.Set(r.stmtList(psess))
		return n

	case OFALL:
		n := psess.nodl(r.pos(psess), OFALL, nil, nil)
		return n

	case OBREAK, OCONTINUE:
		pos := r.pos(psess)
		left, _ := r.exprsOrNil(psess)
		if left != nil {
			left = psess.newname(left.Sym)
		}
		return psess.nodl(pos, op, left, nil)

	case OGOTO, OLABEL:
		return psess.nodl(r.pos(psess), op, psess.newname(r.expr(psess).Sym), nil)

	case OEND:
		return nil

	default:
		psess.
			Fatalf("cannot import %v (%d) node\n"+
				"==> please file an issue and assign to gri@\n", op, int(op))
		panic("unreachable")
	}
}

func (r *importReader) op(psess *PackageSession) Op {
	return Op(r.uint64(psess))
}

func (r *importReader) elemList(psess *PackageSession) []*Node {
	c := r.uint64(psess)
	list := make([]*Node, c)
	for i := range list {
		s := r.ident(psess)
		list[i] = psess.nodSym(OSTRUCTKEY, r.expr(psess), s)
	}
	return list
}

func (r *importReader) exprsOrNil(psess *PackageSession) (a, b *Node) {
	ab := r.uint64(psess)
	if ab&1 != 0 {
		a = r.expr(psess)
	}
	if ab&2 != 0 {
		b = r.node(psess)
	}
	return
}

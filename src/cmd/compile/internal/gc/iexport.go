package gc

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/src"
	"go/ast"
	"io"
	"math/big"
	"strings"
)

// Current indexed export format version. Increase with each format change.
// 0: Go1.11 encoding
const iexportVersion = 0

// predeclReserved is the number of type offsets reserved for types
// implicitly declared in the universe block.
const predeclReserved = 32

// An itag distinguishes the kind of type that was written into the
// indexed export format.
type itag uint64

const (
	// Types
	definedType itag = iota
	pointerType
	sliceType
	arrayType
	chanType
	mapType
	signatureType
	structType
	interfaceType
)

func (psess *PackageSession) iexport(out *bufio.Writer) {

	{

		p := &exporter{marked: make(map[*types.Type]bool)}
		for _, n := range psess.exportlist {
			sym := n.Sym
			p.markType(psess, asNode(sym.Def).Type)
		}
	}

	p := iexporter{
		allPkgs:     map[*types.Pkg]bool{},
		stringIndex: map[string]uint64{},
		declIndex:   map[*Node]uint64{},
		inlineIndex: map[*Node]uint64{},
		typIndex:    map[*types.Type]uint64{},
	}

	for i, pt := range psess.predeclared() {
		p.typIndex[pt] = uint64(i)
	}
	if len(p.typIndex) > predeclReserved {
		psess.
			Fatalf("too many predeclared types: %d > %d", len(p.typIndex), predeclReserved)
	}

	for _, n := range psess.exportlist {
		p.pushDecl(psess, n)
	}

	for !p.declTodo.empty() {
		p.doDecl(psess, p.declTodo.popLeft())
	}

	dataLen := uint64(p.data0.Len())
	w := p.newWriter()
	w.writeIndex(psess, p.declIndex, true)
	w.writeIndex(psess, p.inlineIndex, false)
	w.flush()

	// Assemble header.
	var hdr intWriter
	hdr.WriteByte('i')
	hdr.uint64(iexportVersion)
	hdr.uint64(uint64(p.strings.Len()))
	hdr.uint64(dataLen)

	io.Copy(out, &hdr)
	io.Copy(out, &p.strings)
	io.Copy(out, &p.data0)
}

// writeIndex writes out an object index. mainIndex indicates whether
// we're writing out the main index, which is also read by
// non-compiler tools and includes a complete package description
// (i.e., name and height).
func (w *exportWriter) writeIndex(psess *PackageSession, index map[*Node]uint64, mainIndex bool) {

	pkgObjs := map[*types.Pkg][]*Node{}

	if mainIndex {
		pkgObjs[psess.localpkg] = nil
		for pkg := range w.p.allPkgs {
			pkgObjs[pkg] = nil
		}
	}

	for n := range index {
		pkgObjs[n.Sym.Pkg] = append(pkgObjs[n.Sym.Pkg], n)
	}

	var pkgs []*types.Pkg
	for pkg, objs := range pkgObjs {
		pkgs = append(pkgs, pkg)

		obj.SortSlice(objs, func(i, j int) bool {
			return objs[i].Sym.Name < objs[j].Sym.Name
		})
	}

	obj.SortSlice(pkgs, func(i, j int) bool {
		return pkgs[i].Path < pkgs[j].Path
	})

	w.uint64(uint64(len(pkgs)))
	for _, pkg := range pkgs {
		w.string(pkg.Path)
		if mainIndex {
			w.string(pkg.Name)
			w.uint64(uint64(pkg.Height))
		}

		objs := pkgObjs[pkg]
		w.uint64(uint64(len(objs)))
		for _, n := range objs {
			w.string(n.Sym.Name)
			w.uint64(index[n])
		}
	}
}

type iexporter struct {
	// allPkgs tracks all packages that have been referenced by
	// the export data, so we can ensure to include them in the
	// main index.
	allPkgs map[*types.Pkg]bool

	declTodo nodeQueue

	strings     intWriter
	stringIndex map[string]uint64

	data0       intWriter
	declIndex   map[*Node]uint64
	inlineIndex map[*Node]uint64
	typIndex    map[*types.Type]uint64
}

// stringOff returns the offset of s within the string section.
// If not already present, it's added to the end.
func (p *iexporter) stringOff(s string) uint64 {
	off, ok := p.stringIndex[s]
	if !ok {
		off = uint64(p.strings.Len())
		p.stringIndex[s] = off

		p.strings.uint64(uint64(len(s)))
		p.strings.WriteString(s)
	}
	return off
}

// pushDecl adds n to the declaration work queue, if not already present.
func (p *iexporter) pushDecl(psess *PackageSession, n *Node) {
	if n.Sym == nil || asNode(n.Sym.Def) != n && n.Op != OTYPE {
		psess.
			Fatalf("weird Sym: %v, %v", n, n.Sym)
	}

	if n.Sym.Pkg == psess.builtinpkg || n.Sym.Pkg == psess.unsafepkg {
		return
	}

	if _, ok := p.declIndex[n]; ok {
		return
	}

	p.declIndex[n] = ^uint64(0)
	p.declTodo.pushRight(n)
}

// exportWriter handles writing out individual data section chunks.
type exportWriter struct {
	p *iexporter

	data     intWriter
	currPkg  *types.Pkg
	prevFile string
	prevLine int64
}

func (p *iexporter) doDecl(psess *PackageSession, n *Node) {
	w := p.newWriter()
	w.setPkg(psess, n.Sym.Pkg, false)

	switch n.Op {
	case ONAME:
		switch n.Class() {
		case PEXTERN:

			w.tag('V')
			w.pos(psess, n.Pos)
			w.typ(psess, n.Type)
			w.varExt(n)

		case PFUNC:
			if n.IsMethod(psess) {
				psess.
					Fatalf("unexpected method: %v", n)
			}

			w.tag('F')
			w.pos(psess, n.Pos)
			w.signature(psess, n.Type)
			w.funcExt(psess, n)

		default:
			psess.
				Fatalf("unexpected class: %v, %v", n, n.Class())
		}

	case OLITERAL:

		n = psess.typecheck(n, Erv)
		w.tag('C')
		w.pos(psess, n.Pos)
		w.value(psess, n.Type, n.Val())

	case OTYPE:
		if IsAlias(n.Sym) {

			w.tag('A')
			w.pos(psess, n.Pos)
			w.typ(psess, n.Type)
			break
		}

		w.tag('T')
		w.pos(psess, n.Pos)

		underlying := n.Type.Orig
		if underlying == psess.types.Errortype.Orig {

			underlying = psess.types.Errortype
		}
		w.typ(psess, underlying)

		t := n.Type
		if t.IsInterface() {
			break
		}

		ms := t.Methods()
		w.uint64(uint64(ms.Len()))
		for _, m := range ms.Slice() {
			w.pos(psess, m.Pos)
			w.selector(psess, m.Sym)
			w.param(psess, m.Type.Recv(psess.types))
			w.signature(psess, m.Type)
		}

		for _, m := range ms.Slice() {
			w.methExt(psess, m)
		}

	default:
		psess.
			Fatalf("unexpected node: %v", n)
	}

	p.declIndex[n] = w.flush()
}

func (w *exportWriter) tag(tag byte) {
	w.data.WriteByte(tag)
}

func (p *iexporter) doInline(psess *PackageSession, f *Node) {
	w := p.newWriter()
	w.setPkg(psess, psess.fnpkg(f), false)

	w.stmtList(psess, asNodes(f.Func.Inl.Body))

	p.inlineIndex[f] = w.flush()
}

func (w *exportWriter) pos(psess *PackageSession, pos src.XPos) {
	p := psess.Ctxt.PosTable.Pos(pos)
	file := p.Base().AbsFilename()
	line := int64(p.RelLine(psess.src))

	if file == w.prevFile {
		delta := line - w.prevLine
		w.int64(delta)
		if delta == deltaNewFile {
			w.int64(-1)
		}
	} else {
		w.int64(deltaNewFile)
		w.int64(line)
		w.string(file)
		w.prevFile = file
	}
	w.prevLine = line
}

func (w *exportWriter) pkg(pkg *types.Pkg) {

	w.p.allPkgs[pkg] = true

	w.string(pkg.Path)
}

func (w *exportWriter) qualifiedIdent(psess *PackageSession, n *Node) {

	w.p.pushDecl(psess, n)

	s := n.Sym
	w.string(s.Name)
	w.pkg(s.Pkg)
}

func (w *exportWriter) selector(psess *PackageSession, s *types.Sym) {
	if w.currPkg == nil {
		psess.
			Fatalf("missing currPkg")
	}

	name := s.Name
	if i := strings.LastIndex(name, "."); i >= 0 {
		name = name[i+1:]
	} else {
		pkg := w.currPkg
		if types.IsExported(name) {
			pkg = psess.localpkg
		}
		if s.Pkg != pkg {
			psess.
				Fatalf("package mismatch in selector: %v in package %q, but want %q", s, s.Pkg.Path, pkg.Path)
		}
	}

	w.string(name)
}

func (w *exportWriter) typ(psess *PackageSession, t *types.Type) {
	w.data.uint64(w.p.typOff(psess, t))
}

func (p *iexporter) newWriter() *exportWriter {
	return &exportWriter{p: p}
}

func (w *exportWriter) flush() uint64 {
	off := uint64(w.p.data0.Len())
	io.Copy(&w.p.data0, &w.data)
	return off
}

func (p *iexporter) typOff(psess *PackageSession, t *types.Type) uint64 {
	off, ok := p.typIndex[t]
	if !ok {
		w := p.newWriter()
		w.doTyp(psess, t)
		off = predeclReserved + uint64(w.flush())
		p.typIndex[t] = off
	}
	return off
}

func (w *exportWriter) startType(k itag) {
	w.data.uint64(uint64(k))
}

func (w *exportWriter) doTyp(psess *PackageSession, t *types.Type) {
	if t.Sym != nil {
		if t.Sym.Pkg == psess.builtinpkg || t.Sym.Pkg == psess.unsafepkg {
			psess.
				Fatalf("builtin type missing from typIndex: %v", t)
		}

		w.startType(definedType)
		w.qualifiedIdent(psess, psess.typenod(t))
		return
	}

	switch t.Etype {
	case TPTR32, TPTR64:
		w.startType(pointerType)
		w.typ(psess, t.Elem(psess.types))

	case TSLICE:
		w.startType(sliceType)
		w.typ(psess, t.Elem(psess.types))

	case TARRAY:
		w.startType(arrayType)
		w.uint64(uint64(t.NumElem(psess.types)))
		w.typ(psess, t.Elem(psess.types))

	case TCHAN:
		w.startType(chanType)
		w.uint64(uint64(t.ChanDir(psess.types)))
		w.typ(psess, t.Elem(psess.types))

	case TMAP:
		w.startType(mapType)
		w.typ(psess, t.Key(psess.types))
		w.typ(psess, t.Elem(psess.types))

	case TFUNC:
		w.startType(signatureType)
		w.setPkg(psess, t.Pkg(psess.types), true)
		w.signature(psess, t)

	case TSTRUCT:
		w.startType(structType)
		w.setPkg(psess, t.Pkg(psess.types), true)

		w.uint64(uint64(t.NumFields(psess.types)))
		for _, f := range t.FieldSlice(psess.types) {
			w.pos(psess, f.Pos)
			w.selector(psess, f.Sym)
			w.typ(psess, f.Type)
			w.bool(f.Embedded != 0)
			w.string(f.Note)
		}

	case TINTER:
		var embeddeds, methods []*types.Field
		for _, m := range t.Methods().Slice() {
			if m.Sym != nil {
				methods = append(methods, m)
			} else {
				embeddeds = append(embeddeds, m)
			}
		}

		w.startType(interfaceType)
		w.setPkg(psess, t.Pkg(psess.types), true)

		w.uint64(uint64(len(embeddeds)))
		for _, f := range embeddeds {
			w.pos(psess, f.Pos)
			w.typ(psess, f.Type)
		}

		w.uint64(uint64(len(methods)))
		for _, f := range methods {
			w.pos(psess, f.Pos)
			w.selector(psess, f.Sym)
			w.signature(psess, f.Type)
		}

	default:
		psess.
			Fatalf("unexpected type: %v", t)
	}
}

func (w *exportWriter) setPkg(psess *PackageSession, pkg *types.Pkg, write bool) {
	if pkg == nil {

		pkg = psess.localpkg
	}

	if write {
		w.pkg(pkg)
	}

	w.currPkg = pkg
}

func (w *exportWriter) signature(psess *PackageSession, t *types.Type) {
	w.paramList(psess, t.Params(psess.types).FieldSlice(psess.types))
	w.paramList(psess, t.Results(psess.types).FieldSlice(psess.types))
	if n := t.Params(psess.types).NumFields(psess.types); n > 0 {
		w.bool(t.Params(psess.types).Field(psess.types, n-1).Isddd())
	}
}

func (w *exportWriter) paramList(psess *PackageSession, fs []*types.Field) {
	w.uint64(uint64(len(fs)))
	for _, f := range fs {
		w.param(psess, f)
	}
}

func (w *exportWriter) param(psess *PackageSession, f *types.Field) {
	w.pos(psess, f.Pos)
	w.localIdent(psess, psess.origSym(f.Sym), 0)
	w.typ(psess, f.Type)
}

func (psess *PackageSession) constTypeOf(typ *types.Type) Ctype {
	switch typ {
	case psess.types.Idealint, psess.types.Idealrune:
		return CTINT
	case psess.types.Idealfloat:
		return CTFLT
	case psess.types.Idealcomplex:
		return CTCPLX
	}

	switch typ.Etype {
	case TCHAN, TFUNC, TMAP, TNIL, TINTER, TSLICE:
		return CTNIL
	case TBOOL:
		return CTBOOL
	case TSTRING:
		return CTSTR
	case TINT, TINT8, TINT16, TINT32, TINT64,
		TUINT, TUINT8, TUINT16, TUINT32, TUINT64, TUINTPTR,
		TPTR32, TPTR64, TUNSAFEPTR:
		return CTINT
	case TFLOAT32, TFLOAT64:
		return CTFLT
	case TCOMPLEX64, TCOMPLEX128:
		return CTCPLX
	}
	psess.
		Fatalf("unexpected constant type: %v", typ)
	return 0
}

func (w *exportWriter) value(psess *PackageSession, typ *types.Type, v Val) {
	if typ.IsUntyped(psess.types) {
		typ = psess.untype(v.Ctype(psess))
	}
	w.typ(psess, typ)

	switch psess.constTypeOf(typ) {
	case CTNIL:

		_ = v.U.(*NilVal)
	case CTBOOL:
		w.bool(v.U.(bool))
	case CTSTR:
		w.string(v.U.(string))
	case CTINT:
		w.mpint(psess, &v.U.(*Mpint).Val, typ)
	case CTFLT:
		w.mpfloat(psess, &v.U.(*Mpflt).Val, typ)
	case CTCPLX:
		x := v.U.(*Mpcplx)
		w.mpfloat(psess, &x.Real.Val, typ)
		w.mpfloat(psess, &x.Imag.Val, typ)
	}
}

func (psess *PackageSession) intSize(typ *types.Type) (signed bool, maxBytes uint) {
	if typ.IsUntyped(psess.types) {
		return true, Mpprec / 8
	}

	switch typ.Etype {
	case TFLOAT32, TCOMPLEX64:
		return true, 3
	case TFLOAT64, TCOMPLEX128:
		return true, 7
	}

	signed = typ.IsSigned()
	maxBytes = uint(typ.Size(psess.types))

	switch typ.Etype {
	case TINT, TUINT, TUINTPTR:
		maxBytes = 8
	}

	return
}

// mpint exports a multi-precision integer.
//
// For unsigned types, small values are written out as a single
// byte. Larger values are written out as a length-prefixed big-endian
// byte string, where the length prefix is encoded as its complement.
// For example, bytes 0, 1, and 2 directly represent the integer
// values 0, 1, and 2; while bytes 255, 254, and 253 indicate a 1-,
// 2-, and 3-byte big-endian string follow.
//
// Encoding for signed types use the same general approach as for
// unsigned types, except small values use zig-zag encoding and the
// bottom bit of length prefix byte for large values is reserved as a
// sign bit.
//
// The exact boundary between small and large encodings varies
// according to the maximum number of bytes needed to encode a value
// of type typ. As a special case, 8-bit types are always encoded as a
// single byte.
//
// TODO(mdempsky): Is this level of complexity really worthwhile?
func (w *exportWriter) mpint(psess *PackageSession, x *big.Int, typ *types.Type) {
	signed, maxBytes := psess.intSize(typ)

	negative := x.Sign() < 0
	if !signed && negative {
		psess.
			Fatalf("negative unsigned integer; type %v, value %v", typ, x)
	}

	b := x.Bytes()
	if len(b) > 0 && b[0] == 0 {
		psess.
			Fatalf("leading zeros")
	}
	if uint(len(b)) > maxBytes {
		psess.
			Fatalf("bad mpint length: %d > %d (type %v, value %v)", len(b), maxBytes, typ, x)
	}

	maxSmall := 256 - maxBytes
	if signed {
		maxSmall = 256 - 2*maxBytes
	}
	if maxBytes == 1 {
		maxSmall = 256
	}

	if len(b) <= 1 {
		var ux uint
		if len(b) == 1 {
			ux = uint(b[0])
		}
		if signed {
			ux <<= 1
			if negative {
				ux--
			}
		}
		if ux < maxSmall {
			w.data.WriteByte(byte(ux))
			return
		}
	}

	n := 256 - uint(len(b))
	if signed {
		n = 256 - 2*uint(len(b))
		if negative {
			n |= 1
		}
	}
	if n < maxSmall || n >= 256 {
		psess.
			Fatalf("encoding mistake: %d, %v, %v => %d", len(b), signed, negative, n)
	}

	w.data.WriteByte(byte(n))
	w.data.Write(b)
}

// mpfloat exports a multi-precision floating point number.
//
// The number's value is decomposed into mantissa × 2**exponent, where
// mantissa is an integer. The value is written out as mantissa (as a
// multi-precision integer) and then the exponent, except exponent is
// omitted if mantissa is zero.
func (w *exportWriter) mpfloat(psess *PackageSession, f *big.Float, typ *types.Type) {
	if f.IsInf() {
		psess.
			Fatalf("infinite constant")
	}

	// Break into f = mant × 2**exp, with 0.5 <= mant < 1.
	var mant big.Float
	exp := int64(f.MantExp(&mant))

	prec := mant.MinPrec()
	mant.SetMantExp(&mant, int(prec))
	exp -= int64(prec)

	manti, acc := mant.Int(nil)
	if acc != big.Exact {
		psess.
			Fatalf("exporter: internal error")
	}
	w.mpint(psess, manti, typ)
	if manti.Sign() != 0 {
		w.int64(exp)
	}
}

func (w *exportWriter) bool(b bool) bool {
	var x uint64
	if b {
		x = 1
	}
	w.uint64(x)
	return b
}

func (w *exportWriter) int64(x int64)   { w.data.int64(x) }
func (w *exportWriter) uint64(x uint64) { w.data.uint64(x) }
func (w *exportWriter) string(s string) { w.uint64(w.p.stringOff(s)) }

func (w *exportWriter) varExt(n *Node) {
	w.linkname(n.Sym)
}

func (w *exportWriter) funcExt(psess *PackageSession, n *Node) {
	w.linkname(n.Sym)

	for _, fs := range psess.types.RecvsParams {
		for _, f := range fs(n.Type).FieldSlice(psess.types) {
			w.string(f.Note)
		}
	}

	if n.Func.Inl != nil {
		w.uint64(1 + uint64(n.Func.Inl.Cost))
		if n.Func.ExportInline() {
			w.p.doInline(psess, n)
		}
	} else {
		w.uint64(0)
	}
}

func (w *exportWriter) methExt(psess *PackageSession, m *types.Field) {
	w.bool(m.Nointerface())
	w.funcExt(psess, asNode(m.Type.Nname(psess.types)))
}

func (w *exportWriter) linkname(s *types.Sym) {
	w.string(s.Linkname)
}

func (w *exportWriter) stmtList(psess *PackageSession, list Nodes) {
	for _, n := range list.Slice() {
		w.node(psess, n)
	}
	w.op(OEND)
}

func (w *exportWriter) node(psess *PackageSession, n *Node) {
	if psess.opprec[n.Op] < 0 {
		w.stmt(psess, n)
	} else {
		w.expr(psess, n)
	}
}

// Caution: stmt will emit more than one node for statement nodes n that have a non-empty
// n.Ninit and where n cannot have a natural init section (such as in "if", "for", etc.).
func (w *exportWriter) stmt(psess *PackageSession, n *Node) {
	if n.Ninit.Len() > 0 && !stmtwithinit(n.Op) {

		for _, n := range n.Ninit.Slice() {
			w.stmt(psess, n)
		}
	}

	switch op := n.Op; op {
	case ODCL:
		w.op(ODCL)
		w.pos(psess, n.Left.Pos)
		w.localName(psess, n.Left)
		w.typ(psess, n.Left.Type)

	case OAS:

		if n.Right != nil {
			w.op(OAS)
			w.pos(psess, n.Pos)
			w.expr(psess, n.Left)
			w.expr(psess, n.Right)
		}

	case OASOP:
		w.op(OASOP)
		w.pos(psess, n.Pos)
		w.op(n.SubOp(psess))
		w.expr(psess, n.Left)
		if w.bool(!n.Implicit()) {
			w.expr(psess, n.Right)
		}

	case OAS2, OAS2DOTTYPE, OAS2FUNC, OAS2MAPR, OAS2RECV:
		w.op(OAS2)
		w.pos(psess, n.Pos)
		w.exprList(psess, n.List)
		w.exprList(psess, n.Rlist)

	case ORETURN:
		w.op(ORETURN)
		w.pos(psess, n.Pos)
		w.exprList(psess, n.List)

	case OPROC, ODEFER:
		w.op(op)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)

	case OIF:
		w.op(OIF)
		w.pos(psess, n.Pos)
		w.stmtList(psess, n.Ninit)
		w.expr(psess, n.Left)
		w.stmtList(psess, n.Nbody)
		w.stmtList(psess, n.Rlist)

	case OFOR:
		w.op(OFOR)
		w.pos(psess, n.Pos)
		w.stmtList(psess, n.Ninit)
		w.exprsOrNil(psess, n.Left, n.Right)
		w.stmtList(psess, n.Nbody)

	case ORANGE:
		w.op(ORANGE)
		w.pos(psess, n.Pos)
		w.stmtList(psess, n.List)
		w.expr(psess, n.Right)
		w.stmtList(psess, n.Nbody)

	case OSELECT, OSWITCH:
		w.op(op)
		w.pos(psess, n.Pos)
		w.stmtList(psess, n.Ninit)
		w.exprsOrNil(psess, n.Left, nil)
		w.stmtList(psess, n.List)

	case OCASE, OXCASE:
		w.op(OXCASE)
		w.pos(psess, n.Pos)
		w.stmtList(psess, n.List)
		w.stmtList(psess, n.Nbody)

	case OFALL:
		w.op(OFALL)
		w.pos(psess, n.Pos)

	case OBREAK, OCONTINUE:
		w.op(op)
		w.pos(psess, n.Pos)
		w.exprsOrNil(psess, n.Left, nil)

	case OEMPTY:

	case OGOTO, OLABEL:
		w.op(op)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)

	default:
		psess.
			Fatalf("exporter: CANNOT EXPORT: %v\nPlease notify gri@\n", n.Op)
	}
}

func (w *exportWriter) exprList(psess *PackageSession, list Nodes) {
	for _, n := range list.Slice() {
		w.expr(psess, n)
	}
	w.op(OEND)
}

func (w *exportWriter) expr(psess *PackageSession, n *Node) {

	for n.Op == OPAREN || n.Implicit() && (n.Op == OIND || n.Op == OADDR || n.Op == ODOT || n.Op == ODOTPTR) {
		n = n.Left
	}

	switch op := n.Op; op {

	case OLITERAL:
		if n.Val().Ctype(psess) == CTNIL && n.Orig != nil && n.Orig != n {
			w.expr(psess, n.Orig)
			break
		}
		w.op(OLITERAL)
		w.pos(psess, n.Pos)
		w.value(psess, n.Type, n.Val())

	case ONAME:

		if n.isMethodExpression() {
			w.op(OXDOT)
			w.pos(psess, n.Pos)
			w.expr(psess, n.Left)
			w.selector(psess, n.Right.Sym)
			break
		}

		if (n.Class() == PEXTERN || n.Class() == PFUNC) && !n.isBlank() {
			w.op(ONONAME)
			w.qualifiedIdent(psess, n)
			break
		}

		w.op(ONAME)
		w.localName(psess, n)

	case OTYPE:
		w.op(OTYPE)
		w.typ(psess, n.Type)

	case OPTRLIT:
		w.op(OPTRLIT)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)
		w.bool(n.Implicit())

	case OSTRUCTLIT:
		w.op(OSTRUCTLIT)
		w.pos(psess, n.Pos)
		w.typ(psess, n.Type)
		w.elemList(psess, n.List)

	case OARRAYLIT, OSLICELIT, OMAPLIT:
		w.op(OCOMPLIT)
		w.pos(psess, n.Pos)
		w.typ(psess, n.Type)
		w.exprList(psess, n.List)

	case OKEY:
		w.op(OKEY)
		w.pos(psess, n.Pos)
		w.exprsOrNil(psess, n.Left, n.Right)

	case OXDOT, ODOT, ODOTPTR, ODOTINTER, ODOTMETH:
		w.op(OXDOT)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)
		w.selector(psess, n.Sym)

	case ODOTTYPE, ODOTTYPE2:
		w.op(ODOTTYPE)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)
		w.typ(psess, n.Type)

	case OINDEX, OINDEXMAP:
		w.op(OINDEX)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)
		w.expr(psess, n.Right)

	case OSLICE, OSLICESTR, OSLICEARR:
		w.op(OSLICE)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)
		low, high, _ := n.SliceBounds(psess)
		w.exprsOrNil(psess, low, high)

	case OSLICE3, OSLICE3ARR:
		w.op(OSLICE3)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)
		low, high, max := n.SliceBounds(psess)
		w.exprsOrNil(psess, low, high)
		w.expr(psess, max)

	case OCOPY, OCOMPLEX:

		w.op(op)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)
		w.expr(psess, n.Right)
		w.op(OEND)

	case OCONV, OCONVIFACE, OCONVNOP, OARRAYBYTESTR, OARRAYRUNESTR, OSTRARRAYBYTE, OSTRARRAYRUNE, ORUNESTR:
		w.op(OCONV)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)
		w.typ(psess, n.Type)

	case OREAL, OIMAG, OAPPEND, OCAP, OCLOSE, ODELETE, OLEN, OMAKE, ONEW, OPANIC, ORECOVER, OPRINT, OPRINTN:
		w.op(op)
		w.pos(psess, n.Pos)
		if n.Left != nil {
			w.expr(psess, n.Left)
			w.op(OEND)
		} else {
			w.exprList(psess, n.List)
		}

		if op == OAPPEND {
			w.bool(n.Isddd())
		} else if n.Isddd() {
			psess.
				Fatalf("exporter: unexpected '...' with %v call", op)
		}

	case OCALL, OCALLFUNC, OCALLMETH, OCALLINTER, OGETG:
		w.op(OCALL)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)
		w.exprList(psess, n.List)
		w.bool(n.Isddd())

	case OMAKEMAP, OMAKECHAN, OMAKESLICE:
		w.op(op)
		w.pos(psess, n.Pos)
		w.typ(psess, n.Type)
		switch {
		default:

			w.op(OEND)
		case n.List.Len() != 0:
			w.exprList(psess, n.List)
		case n.Right != nil:
			w.expr(psess, n.Left)
			w.expr(psess, n.Right)
			w.op(OEND)
		case n.Left != nil && (n.Op == OMAKESLICE || !n.Left.Type.IsUntyped(psess.types)):
			w.expr(psess, n.Left)
			w.op(OEND)
		}

	case OPLUS, OMINUS, OADDR, OCOM, OIND, ONOT, ORECV:
		w.op(op)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)

	case OADD, OAND, OANDAND, OANDNOT, ODIV, OEQ, OGE, OGT, OLE, OLT,
		OLSH, OMOD, OMUL, ONE, OOR, OOROR, ORSH, OSEND, OSUB, OXOR:
		w.op(op)
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)
		w.expr(psess, n.Right)

	case OADDSTR:
		w.op(OADDSTR)
		w.pos(psess, n.Pos)
		w.exprList(psess, n.List)

	case OCMPSTR, OCMPIFACE:
		w.op(n.SubOp(psess))
		w.pos(psess, n.Pos)
		w.expr(psess, n.Left)
		w.expr(psess, n.Right)

	case ODCLCONST:

	default:
		psess.
			Fatalf("cannot export %v (%d) node\n"+
				"==> please file an issue and assign to gri@\n", n.Op, int(n.Op))
	}
}

func (w *exportWriter) op(op Op) {
	w.uint64(uint64(op))
}

func (w *exportWriter) exprsOrNil(psess *PackageSession, a, b *Node) {
	ab := 0
	if a != nil {
		ab |= 1
	}
	if b != nil {
		ab |= 2
	}
	w.uint64(uint64(ab))
	if ab&1 != 0 {
		w.expr(psess, a)
	}
	if ab&2 != 0 {
		w.node(psess, b)
	}
}

func (w *exportWriter) elemList(psess *PackageSession, list Nodes) {
	w.uint64(uint64(list.Len()))
	for _, n := range list.Slice() {
		w.selector(psess, n.Sym)
		w.expr(psess, n.Left)
	}
}

func (w *exportWriter) localName(psess *PackageSession, n *Node) {
	// Escape analysis happens after inline bodies are saved, but
	// we're using the same ONAME nodes, so we might still see
	// PAUTOHEAP here.
	//
	// Check for Stackcopy to identify PAUTOHEAP that came from
	// PPARAM/PPARAMOUT, because we only want to include vargen in
	// non-param names.
	var v int32
	if n.Class() == PAUTO || (n.Class() == PAUTOHEAP && n.Name.Param.Stackcopy == nil) {
		v = n.Name.Vargen
	}

	w.localIdent(psess, n.Sym, v)
}

func (w *exportWriter) localIdent(psess *PackageSession, s *types.Sym, v int32) {

	if s == nil {
		w.string("")
		return
	}

	name := s.Name
	if name == "_" {
		w.string("_")
		return
	}

	if i := strings.LastIndex(name, "."); i >= 0 {
		psess.
			Fatalf("unexpected dot in identifier: %v", name)
	}

	if v > 0 {
		if strings.Contains(name, "·") {
			psess.
				Fatalf("exporter: unexpected · in symbol name")
		}
		name = fmt.Sprintf("%s·%d", name, v)
	}

	if !ast.IsExported(name) && s.Pkg != w.currPkg {
		psess.
			Fatalf("weird package in name: %v => %v, not %q", s, name, w.currPkg.Path)
	}

	w.string(name)
}

type intWriter struct {
	bytes.Buffer
}

func (w *intWriter) int64(x int64) {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutVarint(buf[:], x)
	w.Write(buf[:n])
}

func (w *intWriter) uint64(x uint64) {
	var buf [binary.MaxVarintLen64]byte
	n := binary.PutUvarint(buf[:], x)
	w.Write(buf[:n])
}

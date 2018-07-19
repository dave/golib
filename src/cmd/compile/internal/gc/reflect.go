package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/gcprog"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"

	"os"
	"sort"
	"strings"
)

type itabEntry struct {
	t, itype *types.Type
	lsym     *obj.LSym // symbol of the itab itself

	// symbols of each method in
	// the itab, sorted by byte offset;
	// filled in by peekitabs
	entries []*obj.LSym
}

type ptabEntry struct {
	s *types.Sym
	t *types.Type
}

// runtime interface and reflection data structures

// protects signatset

type Sig struct {
	name  *types.Sym
	isym  *types.Sym
	tsym  *types.Sym
	type_ *types.Type
	mtype *types.Type
}

// Builds a type representing a Bucket structure for
// the given map type. This type is not visible to users -
// we include only enough information to generate a correct GC
// program for it.
// Make sure this stays in sync with runtime/map.go.
const (
	BUCKETSIZE = 8
	MAXKEYSIZE = 128
	MAXVALSIZE = 128
)

func (psess *PackageSession) structfieldSize() int { return 3 * psess.Widthptr }
func imethodSize() int                             { return 4 + 4 }

func (psess *PackageSession) uncommonSize(t *types.Type) int {
	if t.Sym == nil && len(psess.methods(t)) == 0 {
		return 0
	}
	return 4 + 2 + 2 + 4 + 4
}

func (psess *PackageSession) makefield(name string, t *types.Type) *types.Field {
	f := types.NewField()
	f.Type = t
	f.Sym = (*types.Pkg)(nil).Lookup(psess.types, name)
	return f
}

// bmap makes the map bucket type given the type of the map.
func (psess *PackageSession) bmap(t *types.Type) *types.Type {
	if t.MapType(psess.types).Bucket != nil {
		return t.MapType(psess.types).Bucket
	}

	bucket := types.New(TSTRUCT)
	keytype := t.Key(psess.types)
	valtype := t.Elem(psess.types)
	psess.
		dowidth(keytype)
	psess.
		dowidth(valtype)
	if keytype.Width > MAXKEYSIZE {
		keytype = psess.types.NewPtr(keytype)
	}
	if valtype.Width > MAXVALSIZE {
		valtype = psess.types.NewPtr(valtype)
	}

	field := make([]*types.Field, 0, 5)

	arr := psess.types.NewArray(psess.types.Types[TUINT8], BUCKETSIZE)
	field = append(field, psess.makefield("topbits", arr))

	arr = psess.types.NewArray(keytype, BUCKETSIZE)
	arr.SetNoalg(true)
	keys := psess.makefield("keys", arr)
	field = append(field, keys)

	arr = psess.types.NewArray(valtype, BUCKETSIZE)
	arr.SetNoalg(true)
	values := psess.makefield("values", arr)
	field = append(field, values)

	if int(valtype.Align) > psess.Widthptr || int(keytype.Align) > psess.Widthptr {
		field = append(field, psess.makefield("pad", psess.types.Types[TUINTPTR]))
	}

	otyp := psess.types.NewPtr(bucket)
	if !psess.types.Haspointers(valtype) && !psess.types.Haspointers(keytype) {
		otyp = psess.types.Types[TUINTPTR]
	}
	overflow := psess.makefield("overflow", otyp)
	field = append(field, overflow)

	bucket.SetNoalg(true)
	bucket.SetFields(psess.types, field[:])
	psess.
		dowidth(bucket)

	if !psess.IsComparable(t.Key(psess.types)) {
		psess.
			Fatalf("unsupported map key type for %v", t)
	}
	if BUCKETSIZE < 8 {
		psess.
			Fatalf("bucket size too small for proper alignment")
	}
	if keytype.Align > BUCKETSIZE {
		psess.
			Fatalf("key align too big for %v", t)
	}
	if valtype.Align > BUCKETSIZE {
		psess.
			Fatalf("value align too big for %v", t)
	}
	if keytype.Width > MAXKEYSIZE {
		psess.
			Fatalf("key size to large for %v", t)
	}
	if valtype.Width > MAXVALSIZE {
		psess.
			Fatalf("value size to large for %v", t)
	}
	if t.Key(psess.types).Width > MAXKEYSIZE && !keytype.IsPtr() {
		psess.
			Fatalf("key indirect incorrect for %v", t)
	}
	if t.Elem(psess.types).Width > MAXVALSIZE && !valtype.IsPtr() {
		psess.
			Fatalf("value indirect incorrect for %v", t)
	}
	if keytype.Width%int64(keytype.Align) != 0 {
		psess.
			Fatalf("key size not a multiple of key align for %v", t)
	}
	if valtype.Width%int64(valtype.Align) != 0 {
		psess.
			Fatalf("value size not a multiple of value align for %v", t)
	}
	if bucket.Align%keytype.Align != 0 {
		psess.
			Fatalf("bucket align not multiple of key align %v", t)
	}
	if bucket.Align%valtype.Align != 0 {
		psess.
			Fatalf("bucket align not multiple of value align %v", t)
	}
	if keys.Offset%int64(keytype.Align) != 0 {
		psess.
			Fatalf("bad alignment of keys in bmap for %v", t)
	}
	if values.Offset%int64(valtype.Align) != 0 {
		psess.
			Fatalf("bad alignment of values in bmap for %v", t)
	}

	if overflow.Offset != bucket.Width-int64(psess.Widthptr) {
		psess.
			Fatalf("bad offset of overflow in bmap for %v", t)
	}

	t.MapType(psess.types).Bucket = bucket

	bucket.StructType(psess.types).Map = t
	return bucket
}

// hmap builds a type representing a Hmap structure for the given map type.
// Make sure this stays in sync with runtime/map.go.
func (psess *PackageSession) hmap(t *types.Type) *types.Type {
	if t.MapType(psess.types).Hmap != nil {
		return t.MapType(psess.types).Hmap
	}

	bmap := psess.bmap(t)

	fields := []*types.Field{psess.
		makefield("count", psess.types.Types[TINT]), psess.
		makefield("flags", psess.types.Types[TUINT8]), psess.
		makefield("B", psess.types.Types[TUINT8]), psess.
		makefield("noverflow", psess.types.Types[TUINT16]), psess.
		makefield("hash0", psess.types.Types[TUINT32]), psess.
		makefield("buckets", psess.types.NewPtr(bmap)), psess.
		makefield("oldbuckets", psess.types.NewPtr(bmap)), psess.
		makefield("nevacuate", psess.types.Types[TUINTPTR]), psess.
		makefield("extra", psess.types.Types[TUNSAFEPTR]),
	}

	hmap := types.New(TSTRUCT)
	hmap.SetNoalg(true)
	hmap.SetFields(psess.types, fields)
	psess.
		dowidth(hmap)

	if size := int64(8 + 5*psess.Widthptr); hmap.Width != size {
		psess.
			Fatalf("hmap size not correct: got %d, want %d", hmap.Width, size)
	}

	t.MapType(psess.types).Hmap = hmap
	hmap.StructType(psess.types).Map = t
	return hmap
}

// hiter builds a type representing an Hiter structure for the given map type.
// Make sure this stays in sync with runtime/map.go.
func (psess *PackageSession) hiter(t *types.Type) *types.Type {
	if t.MapType(psess.types).Hiter != nil {
		return t.MapType(psess.types).Hiter
	}

	hmap := psess.hmap(t)
	bmap := psess.bmap(t)

	fields := []*types.Field{psess.
		makefield("key", psess.types.NewPtr(t.Key(psess.types))), psess.
		makefield("val", psess.types.NewPtr(t.Elem(psess.types))), psess.
		makefield("t", psess.types.Types[TUNSAFEPTR]), psess.
		makefield("h", psess.types.NewPtr(hmap)), psess.
		makefield("buckets", psess.types.NewPtr(bmap)), psess.
		makefield("bptr", psess.types.NewPtr(bmap)), psess.
		makefield("overflow", psess.types.Types[TUNSAFEPTR]), psess.
		makefield("oldoverflow", psess.types.Types[TUNSAFEPTR]), psess.
		makefield("startBucket", psess.types.Types[TUINTPTR]), psess.
		makefield("offset", psess.types.Types[TUINT8]), psess.
		makefield("wrapped", psess.types.Types[TBOOL]), psess.
		makefield("B", psess.types.Types[TUINT8]), psess.
		makefield("i", psess.types.Types[TUINT8]), psess.
		makefield("bucket", psess.types.Types[TUINTPTR]), psess.
		makefield("checkBucket", psess.types.Types[TUINTPTR]),
	}

	hiter := types.New(TSTRUCT)
	hiter.SetNoalg(true)
	hiter.SetFields(psess.types, fields)
	psess.
		dowidth(hiter)
	if hiter.Width != int64(12*psess.Widthptr) {
		psess.
			Fatalf("hash_iter size not correct %d %d", hiter.Width, 12*psess.Widthptr)
	}
	t.MapType(psess.types).Hiter = hiter
	hiter.StructType(psess.types).Map = t
	return hiter
}

// f is method type, with receiver.
// return function type, receiver as first argument (or not).
func (psess *PackageSession) methodfunc(f *types.Type, receiver *types.Type) *types.Type {
	var in []*Node
	if receiver != nil {
		d := psess.anonfield(receiver)
		in = append(in, d)
	}

	for _, t := range f.Params(psess.types).Fields(psess.types).Slice() {
		d := psess.anonfield(t.Type)
		d.SetIsddd(t.Isddd())
		in = append(in, d)
	}

	var out []*Node
	for _, t := range f.Results(psess.types).Fields(psess.types).Slice() {
		d := psess.anonfield(t.Type)
		out = append(out, d)
	}

	t := psess.functype(nil, in, out)
	if f.Nname(psess.types) != nil {

		t.SetNname(psess.types, f.Nname(psess.types))
	}

	return t
}

// methods returns the methods of the non-interface type t, sorted by name.
// Generates stub functions as needed.
func (psess *PackageSession) methods(t *types.Type) []*Sig {

	mt := psess.methtype(t)

	if mt == nil {
		return nil
	}
	psess.
		expandmeth(mt)

	it := t

	if !psess.isdirectiface(it) {
		it = psess.types.NewPtr(t)
	}

	// make list of methods for t,
	// generating code if necessary.
	var ms []*Sig
	for _, f := range mt.AllMethods().Slice() {
		if f.Type.Etype != TFUNC || f.Type.Recv(psess.types) == nil {
			psess.
				Fatalf("non-method on %v method %v %v\n", mt, f.Sym, f)
		}
		if f.Type.Recv(psess.types) == nil {
			psess.
				Fatalf("receiver with no type on %v method %v %v\n", mt, f.Sym, f)
		}
		if f.Nointerface() {
			continue
		}

		method := f.Sym
		if method == nil {
			break
		}

		if !psess.isMethodApplicable(t, f) {
			continue
		}

		sig := &Sig{
			name:  method,
			isym:  psess.methodSym(it, method),
			tsym:  psess.methodSym(t, method),
			type_: psess.methodfunc(f.Type, t),
			mtype: psess.methodfunc(f.Type, nil),
		}
		ms = append(ms, sig)

		this := f.Type.Recv(psess.types).Type

		if !sig.isym.Siggen() {
			sig.isym.SetSiggen(true)
			if !psess.eqtype(this, it) {
				psess.
					compiling_wrappers = true
				psess.
					genwrapper(it, f, sig.isym)
				psess.
					compiling_wrappers = false
			}
		}

		if !sig.tsym.Siggen() {
			sig.tsym.SetSiggen(true)
			if !psess.eqtype(this, t) {
				psess.
					compiling_wrappers = true
				psess.
					genwrapper(t, f, sig.tsym)
				psess.
					compiling_wrappers = false
			}
		}
	}

	return ms
}

// imethods returns the methods of the interface type t, sorted by name.
func (psess *PackageSession) imethods(t *types.Type) []*Sig {
	var methods []*Sig
	for _, f := range t.Fields(psess.types).Slice() {
		if f.Type.Etype != TFUNC || f.Sym == nil {
			continue
		}
		if f.Sym.IsBlank() {
			psess.
				Fatalf("unexpected blank symbol in interface method set")
		}
		if n := len(methods); n > 0 {
			last := methods[n-1]
			if !last.name.Less(f.Sym) {
				psess.
					Fatalf("sigcmp vs sortinter %v %v", last.name, f.Sym)
			}
		}

		sig := &Sig{
			name:  f.Sym,
			mtype: f.Type,
			type_: psess.methodfunc(f.Type, nil),
		}
		methods = append(methods, sig)

		isym := psess.methodSym(t, f.Sym)
		if !isym.Siggen() {
			isym.SetSiggen(true)
			psess.
				genwrapper(t, f, isym)
		}
	}

	return methods
}

func (psess *PackageSession) dimportpath(p *types.Pkg) {
	if p.Pathsym != nil {
		return
	}

	if psess.myimportpath == "runtime" && p == psess.Runtimepkg {
		return
	}

	var str string
	if p == psess.localpkg {

		str = psess.myimportpath
	} else {
		str = p.Path
	}

	s := psess.Ctxt.Lookup("type..importpath." + p.Prefix + ".")
	ot := psess.dnameData(s, 0, str, "", nil, false)
	psess.
		ggloblsym(s, int32(ot), obj.DUPOK|obj.RODATA)
	p.Pathsym = s
}

func (psess *PackageSession) dgopkgpath(s *obj.LSym, ot int, pkg *types.Pkg) int {
	if pkg == nil {
		return psess.duintptr(s, ot, 0)
	}

	if pkg == psess.localpkg && psess.myimportpath == "" {

		ns := psess.Ctxt.Lookup("type..importpath.\"\".")
		return psess.dsymptr(s, ot, ns, 0)
	}
	psess.
		dimportpath(pkg)
	return psess.dsymptr(s, ot, pkg.Pathsym, 0)
}

// dgopkgpathOff writes an offset relocation in s at offset ot to the pkg path symbol.
func (psess *PackageSession) dgopkgpathOff(s *obj.LSym, ot int, pkg *types.Pkg) int {
	if pkg == nil {
		return psess.duint32(s, ot, 0)
	}
	if pkg == psess.localpkg && psess.myimportpath == "" {

		ns := psess.Ctxt.Lookup("type..importpath.\"\".")
		return psess.dsymptrOff(s, ot, ns)
	}
	psess.
		dimportpath(pkg)
	return psess.dsymptrOff(s, ot, pkg.Pathsym)
}

// dnameField dumps a reflect.name for a struct field.
func (psess *PackageSession) dnameField(lsym *obj.LSym, ot int, spkg *types.Pkg, ft *types.Field) int {
	if !types.IsExported(ft.Sym.Name) && ft.Sym.Pkg != spkg {
		psess.
			Fatalf("package mismatch for %v", ft.Sym)
	}
	nsym := psess.dname(ft.Sym.Name, ft.Note, nil, types.IsExported(ft.Sym.Name))
	return psess.dsymptr(lsym, ot, nsym, 0)
}

// dnameData writes the contents of a reflect.name into s at offset ot.
func (psess *PackageSession) dnameData(s *obj.LSym, ot int, name, tag string, pkg *types.Pkg, exported bool) int {
	if len(name) > 1<<16-1 {
		psess.
			Fatalf("name too long: %s", name)
	}
	if len(tag) > 1<<16-1 {
		psess.
			Fatalf("tag too long: %s", tag)
	}

	// Encode name and tag. See reflect/type.go for details.
	var bits byte
	l := 1 + 2 + len(name)
	if exported {
		bits |= 1 << 0
	}
	if len(tag) > 0 {
		l += 2 + len(tag)
		bits |= 1 << 1
	}
	if pkg != nil {
		bits |= 1 << 2
	}
	b := make([]byte, l)
	b[0] = bits
	b[1] = uint8(len(name) >> 8)
	b[2] = uint8(len(name))
	copy(b[3:], name)
	if len(tag) > 0 {
		tb := b[3+len(name):]
		tb[0] = uint8(len(tag) >> 8)
		tb[1] = uint8(len(tag))
		copy(tb[2:], tag)
	}

	ot = int(s.WriteBytes(psess.Ctxt, int64(ot), b))

	if pkg != nil {
		ot = psess.dgopkgpathOff(s, ot, pkg)
	}

	return ot
}

// dname creates a reflect.name for a struct field or method.
func (psess *PackageSession) dname(name, tag string, pkg *types.Pkg, exported bool) *obj.LSym {

	sname := "type..namedata."
	if pkg == nil {

		if name == "" {
			if exported {
				sname += "-noname-exported." + tag
			} else {
				sname += "-noname-unexported." + tag
			}
		} else {
			if exported {
				sname += name + "." + tag
			} else {
				sname += name + "-" + tag
			}
		}
	} else {
		sname = fmt.Sprintf("%s\"\".%d", sname, psess.dnameCount)
		psess.
			dnameCount++
	}
	s := psess.Ctxt.Lookup(sname)
	if len(s.P) > 0 {
		return s
	}
	ot := psess.dnameData(s, 0, name, tag, pkg, exported)
	psess.
		ggloblsym(s, int32(ot), obj.DUPOK|obj.RODATA)
	return s
}

// dextratype dumps the fields of a runtime.uncommontype.
// dataAdd is the offset in bytes after the header where the
// backing array of the []method field is written (by dextratypeData).
func (psess *PackageSession) dextratype(lsym *obj.LSym, ot int, t *types.Type, dataAdd int) int {
	m := psess.methods(t)
	if t.Sym == nil && len(m) == 0 {
		return ot
	}
	noff := int(psess.Rnd(int64(ot), int64(psess.Widthptr)))
	if noff != ot {
		psess.
			Fatalf("unexpected alignment in dextratype for %v", t)
	}

	for _, a := range m {
		psess.
			dtypesym(a.type_)
	}

	ot = psess.dgopkgpathOff(lsym, ot, psess.typePkg(t))

	dataAdd += psess.uncommonSize(t)
	mcount := len(m)
	if mcount != int(uint16(mcount)) {
		psess.
			Fatalf("too many methods on %v: %d", t, mcount)
	}
	xcount := sort.Search(mcount, func(i int) bool { return !types.IsExported(m[i].name.Name) })
	if dataAdd != int(uint32(dataAdd)) {
		psess.
			Fatalf("methods are too far away on %v: %d", t, dataAdd)
	}

	ot = psess.duint16(lsym, ot, uint16(mcount))
	ot = psess.duint16(lsym, ot, uint16(xcount))
	ot = psess.duint32(lsym, ot, uint32(dataAdd))
	ot = psess.duint32(lsym, ot, 0)
	return ot
}

func (psess *PackageSession) typePkg(t *types.Type) *types.Pkg {
	tsym := t.Sym
	if tsym == nil {
		switch t.Etype {
		case TARRAY, TSLICE, TPTR32, TPTR64, TCHAN:
			if t.Elem(psess.types) != nil {
				tsym = t.Elem(psess.types).Sym
			}
		}
	}
	if tsym != nil && t != psess.types.Types[t.Etype] && t != psess.types.Errortype {
		return tsym.Pkg
	}
	return nil
}

// dextratypeData dumps the backing array for the []method field of
// runtime.uncommontype.
func (psess *PackageSession) dextratypeData(lsym *obj.LSym, ot int, t *types.Type) int {
	for _, a := range psess.methods(t) {

		exported := types.IsExported(a.name.Name)
		var pkg *types.Pkg
		if !exported && a.name.Pkg != psess.typePkg(t) {
			pkg = a.name.Pkg
		}
		nsym := psess.dname(a.name.Name, "", pkg, exported)

		ot = psess.dsymptrOff(lsym, ot, nsym)
		ot = psess.dmethodptrOff(lsym, ot, psess.dtypesym(a.mtype))
		ot = psess.dmethodptrOff(lsym, ot, a.isym.Linksym(psess.types))
		ot = psess.dmethodptrOff(lsym, ot, a.tsym.Linksym(psess.types))
	}
	return ot
}

func (psess *PackageSession) dmethodptrOff(s *obj.LSym, ot int, x *obj.LSym) int {
	psess.
		duint32(s, ot, 0)
	r := obj.Addrel(s)
	r.Off = int32(ot)
	r.Siz = 4
	r.Sym = x
	r.Type = objabi.R_METHODOFF
	return ot + 4
}

// typeptrdata returns the length in bytes of the prefix of t
// containing pointer data. Anything after this offset is scalar data.
func (psess *PackageSession) typeptrdata(t *types.Type) int64 {
	if !psess.types.Haspointers(t) {
		return 0
	}

	switch t.Etype {
	case TPTR32,
		TPTR64,
		TUNSAFEPTR,
		TFUNC,
		TCHAN,
		TMAP:
		return int64(psess.Widthptr)

	case TSTRING:

		return int64(psess.Widthptr)

	case TINTER:

		return 2 * int64(psess.Widthptr)

	case TSLICE:

		return int64(psess.Widthptr)

	case TARRAY:

		return (t.NumElem(psess.types)-1)*t.Elem(psess.types).Width + psess.typeptrdata(t.Elem(psess.types))

	case TSTRUCT:
		// Find the last field that has pointers.
		var lastPtrField *types.Field
		for _, t1 := range t.Fields(psess.types).Slice() {
			if psess.types.Haspointers(t1.Type) {
				lastPtrField = t1
			}
		}
		return lastPtrField.Offset + psess.typeptrdata(lastPtrField.Type)

	default:
		psess.
			Fatalf("typeptrdata: unexpected type, %v", t)
		return 0
	}
}

// tflag is documented in reflect/type.go.
//
// tflag values must be kept in sync with copies in:
//	cmd/compile/internal/gc/reflect.go
//	cmd/link/internal/ld/decodesym.go
//	reflect/type.go
//	runtime/type.go
const (
	tflagUncommon  = 1 << 0
	tflagExtraStar = 1 << 1
	tflagNamed     = 1 << 2
)

// dcommontype dumps the contents of a reflect.rtype (runtime._type).
func (psess *PackageSession) dcommontype(lsym *obj.LSym, t *types.Type) int {
	sizeofAlg := 2 * psess.Widthptr
	if psess.algarray == nil {
		psess.
			algarray = psess.sysfunc("algarray")
	}
	psess.
		dowidth(t)
	alg := psess.algtype(t)
	var algsym *obj.LSym
	if alg == ASPECIAL || alg == AMEM {
		algsym = psess.dalgsym(t)
	}

	sptrWeak := true
	var sptr *obj.LSym
	if !t.IsPtr() || t.PtrBase != nil {
		tptr := psess.types.NewPtr(t)
		if t.Sym != nil || psess.methods(tptr) != nil {
			sptrWeak = false
		}
		sptr = psess.dtypesym(tptr)
	}

	gcsym, useGCProg, ptrdata := psess.dgcsym(t)

	ot := 0
	ot = psess.duintptr(lsym, ot, uint64(t.Width))
	ot = psess.duintptr(lsym, ot, uint64(ptrdata))
	ot = psess.duint32(lsym, ot, psess.typehash(t))

	var tflag uint8
	if psess.uncommonSize(t) != 0 {
		tflag |= tflagUncommon
	}
	if t.Sym != nil && t.Sym.Name != "" {
		tflag |= tflagNamed
	}

	exported := false
	p := t.LongString(psess.types)

	if !strings.HasPrefix(p, "*") {
		p = "*" + p
		tflag |= tflagExtraStar
		if t.Sym != nil {
			exported = types.IsExported(t.Sym.Name)
		}
	} else {
		if t.Elem(psess.types) != nil && t.Elem(psess.types).Sym != nil {
			exported = types.IsExported(t.Elem(psess.types).Sym.Name)
		}
	}

	ot = psess.duint8(lsym, ot, tflag)

	i := int(t.Align)

	if i == 0 {
		i = 1
	}
	if i&(i-1) != 0 {
		psess.
			Fatalf("invalid alignment %d for %v", t.Align, t)
	}
	ot = psess.duint8(lsym, ot, t.Align)
	ot = psess.duint8(lsym, ot, t.Align)

	i = psess.kinds[t.Etype]
	if !psess.types.Haspointers(t) {
		i |= objabi.KindNoPointers
	}
	if psess.isdirectiface(t) {
		i |= objabi.KindDirectIface
	}
	if useGCProg {
		i |= objabi.KindGCProg
	}
	ot = psess.duint8(lsym, ot, uint8(i))
	if algsym == nil {
		ot = psess.dsymptr(lsym, ot, psess.algarray, int(alg)*sizeofAlg)
	} else {
		ot = psess.dsymptr(lsym, ot, algsym, 0)
	}
	ot = psess.dsymptr(lsym, ot, gcsym, 0)

	nsym := psess.dname(p, "", nil, exported)
	ot = psess.dsymptrOff(lsym, ot, nsym)

	if sptr == nil {
		ot = psess.duint32(lsym, ot, 0)
	} else if sptrWeak {
		ot = psess.dsymptrWeakOff(lsym, ot, sptr)
	} else {
		ot = psess.dsymptrOff(lsym, ot, sptr)
	}

	return ot
}

// typeHasNoAlg returns whether t does not have any associated hash/eq
// algorithms because t, or some component of t, is marked Noalg.
func (psess *PackageSession) typeHasNoAlg(t *types.Type) bool {
	a, bad := psess.algtype1(t)
	return a == ANOEQ && bad.Noalg()
}

func (psess *PackageSession) typesymname(t *types.Type) string {
	name := t.ShortString(psess.types)

	if psess.typeHasNoAlg(t) {
		name = "noalg." + name
	}
	return name
}

// Fake package for runtime type info (headers)
// Don't access directly, use typeLookup below.

// protects typepkg lookups

func (psess *PackageSession) typeLookup(name string) *types.Sym {
	psess.
		typepkgmu.Lock()
	s := psess.typepkg.Lookup(psess.types, name)
	psess.
		typepkgmu.Unlock()
	return s
}

func (psess *PackageSession) typesym(t *types.Type) *types.Sym {
	return psess.typeLookup(psess.typesymname(t))
}

// tracksym returns the symbol for tracking use of field/method f, assumed
// to be a member of struct/interface type t.
func (psess *PackageSession) tracksym(t *types.Type, f *types.Field) *types.Sym {
	return psess.trackpkg.Lookup(psess.types, t.ShortString(psess.types)+"."+f.Sym.Name)
}

func (psess *PackageSession) typesymprefix(prefix string, t *types.Type) *types.Sym {
	p := prefix + "." + t.ShortString(psess.types)
	s := psess.typeLookup(p)

	return s
}

func (psess *PackageSession) typenamesym(t *types.Type) *types.Sym {
	if t == nil || (t.IsPtr() && t.Elem(psess.types) == nil) || t.IsUntyped(psess.types) {
		psess.
			Fatalf("typenamesym %v", t)
	}
	s := psess.typesym(t)
	psess.
		signatsetmu.Lock()
	psess.
		addsignat(t)
	psess.
		signatsetmu.Unlock()
	return s
}

func (psess *PackageSession) typename(t *types.Type) *Node {
	s := psess.typenamesym(t)
	if s.Def == nil {
		n := psess.newnamel(psess.src.NoXPos, s)
		n.Type = psess.types.Types[TUINT8]
		n.SetClass(PEXTERN)
		n.SetTypecheck(1)
		s.Def = asTypesNode(n)
	}

	n := psess.nod(OADDR, asNode(s.Def), nil)
	n.Type = psess.types.NewPtr(asNode(s.Def).Type)
	n.SetAddable(true)
	n.SetTypecheck(1)
	return n
}

func (psess *PackageSession) itabname(t, itype *types.Type) *Node {
	if t == nil || (t.IsPtr() && t.Elem(psess.types) == nil) || t.IsUntyped(psess.types) || !itype.IsInterface() || itype.IsEmptyInterface(psess.types) {
		psess.
			Fatalf("itabname(%v, %v)", t, itype)
	}
	s := psess.itabpkg.Lookup(psess.types, t.ShortString(psess.types)+","+itype.ShortString(psess.types))
	if s.Def == nil {
		n := psess.newname(s)
		n.Type = psess.types.Types[TUINT8]
		n.SetClass(PEXTERN)
		n.SetTypecheck(1)
		s.Def = asTypesNode(n)
		psess.
			itabs = append(psess.itabs, itabEntry{t: t, itype: itype, lsym: s.Linksym(psess.types)})
	}

	n := psess.nod(OADDR, asNode(s.Def), nil)
	n.Type = psess.types.NewPtr(asNode(s.Def).Type)
	n.SetAddable(true)
	n.SetTypecheck(1)
	return n
}

// isreflexive reports whether t has a reflexive equality operator.
// That is, if x==x for all x of type t.
func (psess *PackageSession) isreflexive(t *types.Type) bool {
	switch t.Etype {
	case TBOOL,
		TINT,
		TUINT,
		TINT8,
		TUINT8,
		TINT16,
		TUINT16,
		TINT32,
		TUINT32,
		TINT64,
		TUINT64,
		TUINTPTR,
		TPTR32,
		TPTR64,
		TUNSAFEPTR,
		TSTRING,
		TCHAN:
		return true

	case TFLOAT32,
		TFLOAT64,
		TCOMPLEX64,
		TCOMPLEX128,
		TINTER:
		return false

	case TARRAY:
		return psess.isreflexive(t.Elem(psess.types))

	case TSTRUCT:
		for _, t1 := range t.Fields(psess.types).Slice() {
			if !psess.isreflexive(t1.Type) {
				return false
			}
		}
		return true

	default:
		psess.
			Fatalf("bad type for map key: %v", t)
		return false
	}
}

// needkeyupdate reports whether map updates with t as a key
// need the key to be updated.
func (psess *PackageSession) needkeyupdate(t *types.Type) bool {
	switch t.Etype {
	case TBOOL, TINT, TUINT, TINT8, TUINT8, TINT16, TUINT16, TINT32, TUINT32,
		TINT64, TUINT64, TUINTPTR, TPTR32, TPTR64, TUNSAFEPTR, TCHAN:
		return false

	case TFLOAT32, TFLOAT64, TCOMPLEX64, TCOMPLEX128,
		TINTER,
		TSTRING:
		return true

	case TARRAY:
		return psess.needkeyupdate(t.Elem(psess.types))

	case TSTRUCT:
		for _, t1 := range t.Fields(psess.types).Slice() {
			if psess.needkeyupdate(t1.Type) {
				return true
			}
		}
		return false

	default:
		psess.
			Fatalf("bad type for map key: %v", t)
		return true
	}
}

// formalType replaces byte and rune aliases with real types.
// They've been separate internally to make error messages
// better, but we have to merge them in the reflect tables.
func (psess *PackageSession) formalType(t *types.Type) *types.Type {
	if t == psess.types.Bytetype || t == psess.types.Runetype {
		return psess.types.Types[t.Etype]
	}
	return t
}

func (psess *PackageSession) dtypesym(t *types.Type) *obj.LSym {
	t = psess.formalType(t)
	if t.IsUntyped(psess.types) {
		psess.
			Fatalf("dtypesym %v", t)
	}

	s := psess.typesym(t)
	lsym := s.Linksym(psess.types)
	if s.Siggen() {
		return lsym
	}
	s.SetSiggen(true)

	tbase := t

	if t.IsPtr() && t.Sym == nil && t.Elem(psess.types).Sym != nil {
		tbase = t.Elem(psess.types)
	}
	dupok := 0
	if tbase.Sym == nil {
		dupok = obj.DUPOK
	}

	if psess.myimportpath != "runtime" || (tbase != psess.types.Types[tbase.Etype] && tbase != psess.types.Bytetype && tbase != psess.types.Runetype && tbase != psess.types.Errortype) {

		if tbase.Sym != nil && tbase.Sym.Pkg != psess.localpkg {
			return lsym
		}

		if psess.isforw[tbase.Etype] {
			return lsym
		}
	}

	ot := 0
	switch t.Etype {
	default:
		ot = psess.dcommontype(lsym, t)
		ot = psess.dextratype(lsym, ot, t, 0)

	case TARRAY:

		s1 := psess.dtypesym(t.Elem(psess.types))
		t2 := psess.types.NewSlice(t.Elem(psess.types))
		s2 := psess.dtypesym(t2)
		ot = psess.dcommontype(lsym, t)
		ot = psess.dsymptr(lsym, ot, s1, 0)
		ot = psess.dsymptr(lsym, ot, s2, 0)
		ot = psess.duintptr(lsym, ot, uint64(t.NumElem(psess.types)))
		ot = psess.dextratype(lsym, ot, t, 0)

	case TSLICE:

		s1 := psess.dtypesym(t.Elem(psess.types))
		ot = psess.dcommontype(lsym, t)
		ot = psess.dsymptr(lsym, ot, s1, 0)
		ot = psess.dextratype(lsym, ot, t, 0)

	case TCHAN:

		s1 := psess.dtypesym(t.Elem(psess.types))
		ot = psess.dcommontype(lsym, t)
		ot = psess.dsymptr(lsym, ot, s1, 0)
		ot = psess.duintptr(lsym, ot, uint64(t.ChanDir(psess.types)))
		ot = psess.dextratype(lsym, ot, t, 0)

	case TFUNC:
		for _, t1 := range t.Recvs(psess.types).Fields(psess.types).Slice() {
			psess.
				dtypesym(t1.Type)
		}
		isddd := false
		for _, t1 := range t.Params(psess.types).Fields(psess.types).Slice() {
			isddd = t1.Isddd()
			psess.
				dtypesym(t1.Type)
		}
		for _, t1 := range t.Results(psess.types).Fields(psess.types).Slice() {
			psess.
				dtypesym(t1.Type)
		}

		ot = psess.dcommontype(lsym, t)
		inCount := t.NumRecvs(psess.types) + t.NumParams(psess.types)
		outCount := t.NumResults(psess.types)
		if isddd {
			outCount |= 1 << 15
		}
		ot = psess.duint16(lsym, ot, uint16(inCount))
		ot = psess.duint16(lsym, ot, uint16(outCount))
		if psess.Widthptr == 8 {
			ot += 4
		}

		dataAdd := (inCount + t.NumResults(psess.types)) * psess.Widthptr
		ot = psess.dextratype(lsym, ot, t, dataAdd)

		for _, t1 := range t.Recvs(psess.types).Fields(psess.types).Slice() {
			ot = psess.dsymptr(lsym, ot, psess.dtypesym(t1.Type), 0)
		}
		for _, t1 := range t.Params(psess.types).Fields(psess.types).Slice() {
			ot = psess.dsymptr(lsym, ot, psess.dtypesym(t1.Type), 0)
		}
		for _, t1 := range t.Results(psess.types).Fields(psess.types).Slice() {
			ot = psess.dsymptr(lsym, ot, psess.dtypesym(t1.Type), 0)
		}

	case TINTER:
		m := psess.imethods(t)
		n := len(m)
		for _, a := range m {
			psess.
				dtypesym(a.type_)
		}

		ot = psess.dcommontype(lsym, t)

		var tpkg *types.Pkg
		if t.Sym != nil && t != psess.types.Types[t.Etype] && t != psess.types.Errortype {
			tpkg = t.Sym.Pkg
		}
		ot = psess.dgopkgpath(lsym, ot, tpkg)

		ot = psess.dsymptr(lsym, ot, lsym, ot+3*psess.Widthptr+psess.uncommonSize(t))
		ot = psess.duintptr(lsym, ot, uint64(n))
		ot = psess.duintptr(lsym, ot, uint64(n))
		dataAdd := imethodSize() * n
		ot = psess.dextratype(lsym, ot, t, dataAdd)

		for _, a := range m {

			exported := types.IsExported(a.name.Name)
			var pkg *types.Pkg
			if !exported && a.name.Pkg != tpkg {
				pkg = a.name.Pkg
			}
			nsym := psess.dname(a.name.Name, "", pkg, exported)

			ot = psess.dsymptrOff(lsym, ot, nsym)
			ot = psess.dsymptrOff(lsym, ot, psess.dtypesym(a.type_))
		}

	case TMAP:
		s1 := psess.dtypesym(t.Key(psess.types))
		s2 := psess.dtypesym(t.Elem(psess.types))
		s3 := psess.dtypesym(psess.bmap(t))
		ot = psess.dcommontype(lsym, t)
		ot = psess.dsymptr(lsym, ot, s1, 0)
		ot = psess.dsymptr(lsym, ot, s2, 0)
		ot = psess.dsymptr(lsym, ot, s3, 0)
		if t.Key(psess.types).Width > MAXKEYSIZE {
			ot = psess.duint8(lsym, ot, uint8(psess.Widthptr))
			ot = psess.duint8(lsym, ot, 1)
		} else {
			ot = psess.duint8(lsym, ot, uint8(t.Key(psess.types).Width))
			ot = psess.duint8(lsym, ot, 0)
		}

		if t.Elem(psess.types).Width > MAXVALSIZE {
			ot = psess.duint8(lsym, ot, uint8(psess.Widthptr))
			ot = psess.duint8(lsym, ot, 1)
		} else {
			ot = psess.duint8(lsym, ot, uint8(t.Elem(psess.types).Width))
			ot = psess.duint8(lsym, ot, 0)
		}

		ot = psess.duint16(lsym, ot, uint16(psess.bmap(t).Width))
		ot = psess.duint8(lsym, ot, uint8(obj.Bool2int(psess.isreflexive(t.Key(psess.types)))))
		ot = psess.duint8(lsym, ot, uint8(obj.Bool2int(psess.needkeyupdate(t.Key(psess.types)))))
		ot = psess.dextratype(lsym, ot, t, 0)

	case TPTR32, TPTR64:
		if t.Elem(psess.types).Etype == TANY {

			ot = psess.dcommontype(lsym, t)
			ot = psess.dextratype(lsym, ot, t, 0)

			break
		}

		s1 := psess.dtypesym(t.Elem(psess.types))

		ot = psess.dcommontype(lsym, t)
		ot = psess.dsymptr(lsym, ot, s1, 0)
		ot = psess.dextratype(lsym, ot, t, 0)

	case TSTRUCT:
		fields := t.Fields(psess.types).Slice()
		for _, t1 := range fields {
			psess.
				dtypesym(t1.Type)
		}

		// All non-exported struct field names within a struct
		// type must originate from a single package. By
		// identifying and recording that package within the
		// struct type descriptor, we can omit that
		// information from the field descriptors.
		var spkg *types.Pkg
		for _, f := range fields {
			if !types.IsExported(f.Sym.Name) {
				spkg = f.Sym.Pkg
				break
			}
		}

		ot = psess.dcommontype(lsym, t)
		ot = psess.dgopkgpath(lsym, ot, spkg)
		ot = psess.dsymptr(lsym, ot, lsym, ot+3*psess.Widthptr+psess.uncommonSize(t))
		ot = psess.duintptr(lsym, ot, uint64(len(fields)))
		ot = psess.duintptr(lsym, ot, uint64(len(fields)))

		dataAdd := len(fields) * psess.structfieldSize()
		ot = psess.dextratype(lsym, ot, t, dataAdd)

		for _, f := range fields {

			ot = psess.dnameField(lsym, ot, spkg, f)
			ot = psess.dsymptr(lsym, ot, psess.dtypesym(f.Type), 0)
			offsetAnon := uint64(f.Offset) << 1
			if offsetAnon>>1 != uint64(f.Offset) {
				psess.
					Fatalf("%v: bad field offset for %s", t, f.Sym.Name)
			}
			if f.Embedded != 0 {
				offsetAnon |= 1
			}
			ot = psess.duintptr(lsym, ot, offsetAnon)
		}
	}

	ot = psess.dextratypeData(lsym, ot, t)
	psess.
		ggloblsym(lsym, int32(ot), int16(dupok|obj.RODATA))

	keep := psess.Ctxt.Flag_dynlink
	if !keep && t.Sym == nil {

		switch t.Etype {
		case TPTR32, TPTR64, TARRAY, TCHAN, TFUNC, TMAP, TSLICE, TSTRUCT:
			keep = true
		}
	}

	if psess.typeHasNoAlg(t) {
		keep = false
	}
	lsym.Set(obj.AttrMakeTypelink, keep)

	return lsym
}

// for each itabEntry, gather the methods on
// the concrete type that implement the interface
func (psess *PackageSession) peekitabs() {
	for i := range psess.itabs {
		tab := &psess.itabs[i]
		methods := psess.genfun(tab.t, tab.itype)
		if len(methods) == 0 {
			continue
		}
		tab.entries = methods
	}
}

// for the given concrete type and interface
// type, return the (sorted) set of methods
// on the concrete type that implement the interface
func (psess *PackageSession) genfun(t, it *types.Type) []*obj.LSym {
	if t == nil || it == nil {
		return nil
	}
	sigs := psess.imethods(it)
	methods := psess.methods(t)
	out := make([]*obj.LSym, 0, len(sigs))

	if len(sigs) == 0 {
		return nil
	}

	for _, m := range methods {
		if m.name == sigs[0].name {
			out = append(out, m.isym.Linksym(psess.types))
			sigs = sigs[1:]
			if len(sigs) == 0 {
				break
			}
		}
	}

	if len(sigs) != 0 {
		psess.
			Fatalf("incomplete itab")
	}

	return out
}

// itabsym uses the information gathered in
// peekitabs to de-virtualize interface methods.
// Since this is called by the SSA backend, it shouldn't
// generate additional Nodes, Syms, etc.
func (psess *PackageSession) itabsym(it *obj.LSym, offset int64) *obj.LSym {
	var syms []*obj.LSym
	if it == nil {
		return nil
	}

	for i := range psess.itabs {
		e := &psess.itabs[i]
		if e.lsym == it {
			syms = e.entries
			break
		}
	}
	if syms == nil {
		return nil
	}

	methodnum := int((offset - 2*int64(psess.Widthptr) - 8) / int64(psess.Widthptr))
	if methodnum >= len(syms) {
		return nil
	}
	return syms[methodnum]
}

// addsignat ensures that a runtime type descriptor is emitted for t.
func (psess *PackageSession) addsignat(t *types.Type) {
	psess.
		signatset[t] = struct{}{}
}

func (psess *PackageSession) addsignats(dcls []*Node) {

	for _, n := range dcls {
		if n.Op == OTYPE {
			psess.
				addsignat(n.Type)
		}
	}
}

func (psess *PackageSession) dumpsignats() {

	signats := make([]typeAndStr, len(psess.signatset))
	for len(psess.signatset) > 0 {
		signats = signats[:0]

		for t := range psess.signatset {
			signats = append(signats, typeAndStr{t: t, short: psess.typesymname(t), regular: t.String(psess.types)})
			delete(psess.signatset, t)
		}
		sort.Sort(typesByString(signats))
		for _, ts := range signats {
			t := ts.t
			psess.
				dtypesym(t)
			if t.Sym != nil {
				psess.
					dtypesym(psess.types.NewPtr(t))
			}
		}
	}
}

func (psess *PackageSession) dumptabs() {

	for _, i := range psess.itabs {

		o := psess.dsymptr(i.lsym, 0, psess.dtypesym(i.itype), 0)
		o = psess.dsymptr(i.lsym, o, psess.dtypesym(i.t), 0)
		o = psess.duint32(i.lsym, o, psess.typehash(i.t))
		o += 4
		for _, fn := range psess.genfun(i.t, i.itype) {
			o = psess.dsymptr(i.lsym, o, fn, 0)
		}
		psess.
			ggloblsym(i.lsym, int32(o), int16(obj.DUPOK|obj.RODATA))
		ilink := psess.itablinkpkg.Lookup(psess.types, i.t.ShortString(psess.types)+","+i.itype.ShortString(psess.types)).Linksym(psess.types)
		psess.
			dsymptr(ilink, 0, i.lsym, 0)
		psess.
			ggloblsym(ilink, int32(psess.Widthptr), int16(obj.DUPOK|obj.RODATA))
	}

	if psess.localpkg.Name == "main" && len(psess.ptabs) > 0 {
		ot := 0
		s := psess.Ctxt.Lookup("go.plugin.tabs")
		for _, p := range psess.ptabs {

			nsym := psess.dname(p.s.Name, "", nil, true)
			ot = psess.dsymptrOff(s, ot, nsym)
			ot = psess.dsymptrOff(s, ot, psess.dtypesym(p.t))
		}
		psess.
			ggloblsym(s, int32(ot), int16(obj.RODATA))

		ot = 0
		s = psess.Ctxt.Lookup("go.plugin.exports")
		for _, p := range psess.ptabs {
			ot = psess.dsymptr(s, ot, p.s.Linksym(psess.types), 0)
		}
		psess.
			ggloblsym(s, int32(ot), int16(obj.RODATA))
	}
}

func (psess *PackageSession) dumpimportstrings() {

	for _, p := range psess.types.ImportedPkgList() {
		psess.
			dimportpath(p)
	}
}

func (psess *PackageSession) dumpbasictypes() {

	if psess.myimportpath == "runtime" {
		for i := types.EType(1); i <= TBOOL; i++ {
			psess.
				dtypesym(psess.types.NewPtr(psess.types.Types[i]))
		}
		psess.
			dtypesym(psess.types.NewPtr(psess.types.Types[TSTRING]))
		psess.
			dtypesym(psess.types.NewPtr(psess.types.Types[TUNSAFEPTR]))
		psess.
			dtypesym(psess.types.NewPtr(psess.types.Errortype))
		psess.
			dtypesym(psess.functype(nil, []*Node{psess.anonfield(psess.types.Errortype)}, []*Node{psess.anonfield(psess.types.Types[TSTRING])}))
		psess.
			dimportpath(psess.Runtimepkg)

		if psess.flag_race {
			psess.
				dimportpath(psess.racepkg)
		}
		if psess.flag_msan {
			psess.
				dimportpath(psess.msanpkg)
		}
		psess.
			dimportpath(psess.types.NewPkg("main", ""))
	}
}

type typeAndStr struct {
	t       *types.Type
	short   string
	regular string
}

type typesByString []typeAndStr

func (a typesByString) Len() int { return len(a) }
func (a typesByString) Less(i, j int) bool {
	if a[i].short != a[j].short {
		return a[i].short < a[j].short
	}

	return a[i].regular < a[j].regular
}
func (a typesByString) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (psess *PackageSession) dalgsym(t *types.Type) *obj.LSym {
	var lsym *obj.LSym
	var hashfunc *obj.LSym
	var eqfunc *obj.LSym

	if psess.algtype(t) == AMEM {

		p := fmt.Sprintf(".alg%d", t.Width)

		s := psess.typeLookup(p)
		lsym = s.Linksym(psess.types)
		if s.AlgGen() {
			return lsym
		}
		s.SetAlgGen(true)

		if psess.memhashvarlen == nil {
			psess.
				memhashvarlen = psess.sysfunc("memhash_varlen")
			psess.
				memequalvarlen = psess.sysfunc("memequal_varlen")
		}

		p = fmt.Sprintf(".hashfunc%d", t.Width)

		hashfunc = psess.typeLookup(p).Linksym(psess.types)

		ot := 0
		ot = psess.dsymptr(hashfunc, ot, psess.memhashvarlen, 0)
		ot = psess.duintptr(hashfunc, ot, uint64(t.Width))
		psess.
			ggloblsym(hashfunc, int32(ot), obj.DUPOK|obj.RODATA)

		p = fmt.Sprintf(".eqfunc%d", t.Width)

		eqfunc = psess.typeLookup(p).Linksym(psess.types)

		ot = 0
		ot = psess.dsymptr(eqfunc, ot, psess.memequalvarlen, 0)
		ot = psess.duintptr(eqfunc, ot, uint64(t.Width))
		psess.
			ggloblsym(eqfunc, int32(ot), obj.DUPOK|obj.RODATA)
	} else {

		s := psess.typesymprefix(".alg", t)
		lsym = s.Linksym(psess.types)

		hash := psess.typesymprefix(".hash", t)
		eq := psess.typesymprefix(".eq", t)
		hashfunc = psess.typesymprefix(".hashfunc", t).Linksym(psess.types)
		eqfunc = psess.typesymprefix(".eqfunc", t).Linksym(psess.types)
		psess.
			genhash(hash, t)
		psess.
			geneq(eq, t)
		psess.
			dsymptr(hashfunc, 0, hash.Linksym(psess.types), 0)
		psess.
			ggloblsym(hashfunc, int32(psess.Widthptr), obj.DUPOK|obj.RODATA)
		psess.
			dsymptr(eqfunc, 0, eq.Linksym(psess.types), 0)
		psess.
			ggloblsym(eqfunc, int32(psess.Widthptr), obj.DUPOK|obj.RODATA)
	}

	ot := 0

	ot = psess.dsymptr(lsym, ot, hashfunc, 0)
	ot = psess.dsymptr(lsym, ot, eqfunc, 0)
	psess.
		ggloblsym(lsym, int32(ot), obj.DUPOK|obj.RODATA)
	return lsym
}

// maxPtrmaskBytes is the maximum length of a GC ptrmask bitmap,
// which holds 1-bit entries describing where pointers are in a given type.
// Above this length, the GC information is recorded as a GC program,
// which can express repetition compactly. In either form, the
// information is used by the runtime to initialize the heap bitmap,
// and for large types (like 128 or more words), they are roughly the
// same speed. GC programs are never much larger and often more
// compact. (If large arrays are involved, they can be arbitrarily
// more compact.)
//
// The cutoff must be large enough that any allocation large enough to
// use a GC program is large enough that it does not share heap bitmap
// bytes with any other objects, allowing the GC program execution to
// assume an aligned start and not use atomic operations. In the current
// runtime, this means all malloc size classes larger than the cutoff must
// be multiples of four words. On 32-bit systems that's 16 bytes, and
// all size classes >= 16 bytes are 16-byte aligned, so no real constraint.
// On 64-bit systems, that's 32 bytes, and 32-byte alignment is guaranteed
// for size classes >= 256 bytes. On a 64-bit system, 256 bytes allocated
// is 32 pointers, the bits for which fit in 4 bytes. So maxPtrmaskBytes
// must be >= 4.
//
// We used to use 16 because the GC programs do have some constant overhead
// to get started, and processing 128 pointers seems to be enough to
// amortize that overhead well.
//
// To make sure that the runtime's chansend can call typeBitsBulkBarrier,
// we raised the limit to 2048, so that even 32-bit systems are guaranteed to
// use bitmaps for objects up to 64 kB in size.
//
// Also known to reflect/type.go.
//
const maxPtrmaskBytes = 2048

// dgcsym emits and returns a data symbol containing GC information for type t,
// along with a boolean reporting whether the UseGCProg bit should be set in
// the type kind, and the ptrdata field to record in the reflect type information.
func (psess *PackageSession) dgcsym(t *types.Type) (lsym *obj.LSym, useGCProg bool, ptrdata int64) {
	ptrdata = psess.typeptrdata(t)
	if ptrdata/int64(psess.Widthptr) <= maxPtrmaskBytes*8 {
		lsym = psess.dgcptrmask(t)
		return
	}

	useGCProg = true
	lsym, ptrdata = psess.dgcprog(t)
	return
}

// dgcptrmask emits and returns the symbol containing a pointer mask for type t.
func (psess *PackageSession) dgcptrmask(t *types.Type) *obj.LSym {
	ptrmask := make([]byte, (psess.typeptrdata(t)/int64(psess.Widthptr)+7)/8)
	psess.
		fillptrmask(t, ptrmask)
	p := fmt.Sprintf("gcbits.%x", ptrmask)

	sym := psess.Runtimepkg.Lookup(psess.types, p)
	lsym := sym.Linksym(psess.types)
	if !sym.Uniq() {
		sym.SetUniq(true)
		for i, x := range ptrmask {
			psess.
				duint8(lsym, i, x)
		}
		psess.
			ggloblsym(lsym, int32(len(ptrmask)), obj.DUPOK|obj.RODATA|obj.LOCAL)
	}
	return lsym
}

// fillptrmask fills in ptrmask with 1s corresponding to the
// word offsets in t that hold pointers.
// ptrmask is assumed to fit at least typeptrdata(t)/Widthptr bits.
func (psess *PackageSession) fillptrmask(t *types.Type, ptrmask []byte) {
	for i := range ptrmask {
		ptrmask[i] = 0
	}
	if !psess.types.Haspointers(t) {
		return
	}

	vec := bvalloc(8 * int32(len(ptrmask)))
	psess.
		onebitwalktype1(t, 0, vec)

	nptr := psess.typeptrdata(t) / int64(psess.Widthptr)
	for i := int64(0); i < nptr; i++ {
		if vec.Get(psess, int32(i)) {
			ptrmask[i/8] |= 1 << (uint(i) % 8)
		}
	}
}

// dgcprog emits and returns the symbol containing a GC program for type t
// along with the size of the data described by the program (in the range [typeptrdata(t), t.Width]).
// In practice, the size is typeptrdata(t) except for non-trivial arrays.
// For non-trivial arrays, the program describes the full t.Width size.
func (psess *PackageSession) dgcprog(t *types.Type) (*obj.LSym, int64) {
	psess.
		dowidth(t)
	if t.Width == BADWIDTH {
		psess.
			Fatalf("dgcprog: %v badwidth", t)
	}
	lsym := psess.typesymprefix(".gcprog", t).Linksym(psess.types)
	var p GCProg
	p.init(psess, lsym)
	p.emit(psess, t, 0)
	offset := p.w.BitIndex() * int64(psess.Widthptr)
	p.end(psess)
	if ptrdata := psess.typeptrdata(t); offset < ptrdata || offset > t.Width {
		psess.
			Fatalf("dgcprog: %v: offset=%d but ptrdata=%d size=%d", t, offset, ptrdata, t.Width)
	}
	return lsym, offset
}

type GCProg struct {
	lsym   *obj.LSym
	symoff int
	w      gcprog.Writer
}

// set by -d gcprog

func (p *GCProg) init(psess *PackageSession, lsym *obj.LSym) {
	p.lsym = lsym
	p.symoff = 4
	p.w.Init(p.writeByte)
	if psess.Debug_gcprog > 0 {
		fmt.Fprintf(os.Stderr, "compile: start GCProg for %v\n", lsym)
		p.w.Debug(os.Stderr)
	}
}

func (p *GCProg) writeByte(psess *PackageSession, x byte) {
	p.symoff = psess.duint8(p.lsym, p.symoff, x)
}

func (p *GCProg) end(psess *PackageSession) {
	p.w.End()
	psess.
		duint32(p.lsym, 0, uint32(p.symoff-4))
	psess.
		ggloblsym(p.lsym, int32(p.symoff), obj.DUPOK|obj.RODATA|obj.LOCAL)
	if psess.Debug_gcprog > 0 {
		fmt.Fprintf(os.Stderr, "compile: end GCProg for %v\n", p.lsym)
	}
}

func (p *GCProg) emit(psess *PackageSession, t *types.Type, offset int64) {
	psess.
		dowidth(t)
	if !psess.types.Haspointers(t) {
		return
	}
	if t.Width == int64(psess.Widthptr) {
		p.w.Ptr(offset / int64(psess.Widthptr))
		return
	}
	switch t.Etype {
	default:
		psess.
			Fatalf("GCProg.emit: unexpected type %v", t)

	case TSTRING:
		p.w.Ptr(offset / int64(psess.Widthptr))

	case TINTER:

		p.w.Ptr(offset/int64(psess.Widthptr) + 1)

	case TSLICE:
		p.w.Ptr(offset / int64(psess.Widthptr))

	case TARRAY:
		if t.NumElem(psess.types) == 0 {
			psess.
				Fatalf("GCProg.emit: empty array")
		}

		count := t.NumElem(psess.types)
		elem := t.Elem(psess.types)
		for elem.IsArray() {
			count *= elem.NumElem(psess.types)
			elem = elem.Elem(psess.types)
		}

		if !p.w.ShouldRepeat(elem.Width/int64(psess.Widthptr), count) {

			for i := int64(0); i < count; i++ {
				p.emit(psess, elem, offset+i*elem.Width)
			}
			return
		}
		p.emit(psess, elem, offset)
		p.w.ZeroUntil((offset + elem.Width) / int64(psess.Widthptr))
		p.w.Repeat(elem.Width/int64(psess.Widthptr), count-1)

	case TSTRUCT:
		for _, t1 := range t.Fields(psess.types).Slice() {
			p.emit(psess, t1.Type, offset+t1.Offset)
		}
	}
}

// zeroaddr returns the address of a symbol with at least
// size bytes of zeros.
func (psess *PackageSession) zeroaddr(size int64) *Node {
	if size >= 1<<31 {
		psess.
			Fatalf("map value too big %d", size)
	}
	if psess.zerosize < size {
		psess.
			zerosize = size
	}
	s := psess.mappkg.Lookup(psess.types, "zero")
	if s.Def == nil {
		x := psess.newname(s)
		x.Type = psess.types.Types[TUINT8]
		x.SetClass(PEXTERN)
		x.SetTypecheck(1)
		s.Def = asTypesNode(x)
	}
	z := psess.nod(OADDR, asNode(s.Def), nil)
	z.Type = psess.types.NewPtr(psess.types.Types[TUINT8])
	z.SetAddable(true)
	z.SetTypecheck(1)
	return z
}

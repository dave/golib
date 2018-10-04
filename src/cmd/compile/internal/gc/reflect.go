// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/gcprog"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"os"
	"sort"
	"strings"
	"sync"
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

func (pstate *PackageState) structfieldSize() int { return 3 * pstate.Widthptr } // Sizeof(runtime.structfield{})
func imethodSize() int                            { return 4 + 4 }               // Sizeof(runtime.imethod{})

func (pstate *PackageState) uncommonSize(t *types.Type) int { // Sizeof(runtime.uncommontype{})
	if t.Sym == nil && len(pstate.methods(t)) == 0 {
		return 0
	}
	return 4 + 2 + 2 + 4 + 4
}

func (pstate *PackageState) makefield(name string, t *types.Type) *types.Field {
	f := types.NewField()
	f.Type = t
	f.Sym = (*types.Pkg)(nil).Lookup(pstate.types, name)
	return f
}

// bmap makes the map bucket type given the type of the map.
func (pstate *PackageState) bmap(t *types.Type) *types.Type {
	if t.MapType(pstate.types).Bucket != nil {
		return t.MapType(pstate.types).Bucket
	}

	bucket := types.New(TSTRUCT)
	keytype := t.Key(pstate.types)
	valtype := t.Elem(pstate.types)
	pstate.dowidth(keytype)
	pstate.dowidth(valtype)
	if keytype.Width > MAXKEYSIZE {
		keytype = pstate.types.NewPtr(keytype)
	}
	if valtype.Width > MAXVALSIZE {
		valtype = pstate.types.NewPtr(valtype)
	}

	field := make([]*types.Field, 0, 5)

	// The first field is: uint8 topbits[BUCKETSIZE].
	arr := pstate.types.NewArray(pstate.types.Types[TUINT8], BUCKETSIZE)
	field = append(field, pstate.makefield("topbits", arr))

	arr = pstate.types.NewArray(keytype, BUCKETSIZE)
	arr.SetNoalg(true)
	keys := pstate.makefield("keys", arr)
	field = append(field, keys)

	arr = pstate.types.NewArray(valtype, BUCKETSIZE)
	arr.SetNoalg(true)
	values := pstate.makefield("values", arr)
	field = append(field, values)

	// Make sure the overflow pointer is the last memory in the struct,
	// because the runtime assumes it can use size-ptrSize as the
	// offset of the overflow pointer. We double-check that property
	// below once the offsets and size are computed.
	//
	// BUCKETSIZE is 8, so the struct is aligned to 64 bits to this point.
	// On 32-bit systems, the max alignment is 32-bit, and the
	// overflow pointer will add another 32-bit field, and the struct
	// will end with no padding.
	// On 64-bit systems, the max alignment is 64-bit, and the
	// overflow pointer will add another 64-bit field, and the struct
	// will end with no padding.
	// On nacl/amd64p32, however, the max alignment is 64-bit,
	// but the overflow pointer will add only a 32-bit field,
	// so if the struct needs 64-bit padding (because a key or value does)
	// then it would end with an extra 32-bit padding field.
	// Preempt that by emitting the padding here.
	if int(valtype.Align) > pstate.Widthptr || int(keytype.Align) > pstate.Widthptr {
		field = append(field, pstate.makefield("pad", pstate.types.Types[TUINTPTR]))
	}

	// If keys and values have no pointers, the map implementation
	// can keep a list of overflow pointers on the side so that
	// buckets can be marked as having no pointers.
	// Arrange for the bucket to have no pointers by changing
	// the type of the overflow field to uintptr in this case.
	// See comment on hmap.overflow in runtime/map.go.
	otyp := pstate.types.NewPtr(bucket)
	if !pstate.types.Haspointers(valtype) && !pstate.types.Haspointers(keytype) {
		otyp = pstate.types.Types[TUINTPTR]
	}
	overflow := pstate.makefield("overflow", otyp)
	field = append(field, overflow)

	// link up fields
	bucket.SetNoalg(true)
	bucket.SetFields(pstate.types, field[:])
	pstate.dowidth(bucket)

	// Check invariants that map code depends on.
	if !pstate.IsComparable(t.Key(pstate.types)) {
		pstate.Fatalf("unsupported map key type for %v", t)
	}
	if BUCKETSIZE < 8 {
		pstate.Fatalf("bucket size too small for proper alignment")
	}
	if keytype.Align > BUCKETSIZE {
		pstate.Fatalf("key align too big for %v", t)
	}
	if valtype.Align > BUCKETSIZE {
		pstate.Fatalf("value align too big for %v", t)
	}
	if keytype.Width > MAXKEYSIZE {
		pstate.Fatalf("key size to large for %v", t)
	}
	if valtype.Width > MAXVALSIZE {
		pstate.Fatalf("value size to large for %v", t)
	}
	if t.Key(pstate.types).Width > MAXKEYSIZE && !keytype.IsPtr() {
		pstate.Fatalf("key indirect incorrect for %v", t)
	}
	if t.Elem(pstate.types).Width > MAXVALSIZE && !valtype.IsPtr() {
		pstate.Fatalf("value indirect incorrect for %v", t)
	}
	if keytype.Width%int64(keytype.Align) != 0 {
		pstate.Fatalf("key size not a multiple of key align for %v", t)
	}
	if valtype.Width%int64(valtype.Align) != 0 {
		pstate.Fatalf("value size not a multiple of value align for %v", t)
	}
	if bucket.Align%keytype.Align != 0 {
		pstate.Fatalf("bucket align not multiple of key align %v", t)
	}
	if bucket.Align%valtype.Align != 0 {
		pstate.Fatalf("bucket align not multiple of value align %v", t)
	}
	if keys.Offset%int64(keytype.Align) != 0 {
		pstate.Fatalf("bad alignment of keys in bmap for %v", t)
	}
	if values.Offset%int64(valtype.Align) != 0 {
		pstate.Fatalf("bad alignment of values in bmap for %v", t)
	}

	// Double-check that overflow field is final memory in struct,
	// with no padding at end. See comment above.
	if overflow.Offset != bucket.Width-int64(pstate.Widthptr) {
		pstate.Fatalf("bad offset of overflow in bmap for %v", t)
	}

	t.MapType(pstate.types).Bucket = bucket

	bucket.StructType(pstate.types).Map = t
	return bucket
}

// hmap builds a type representing a Hmap structure for the given map type.
// Make sure this stays in sync with runtime/map.go.
func (pstate *PackageState) hmap(t *types.Type) *types.Type {
	if t.MapType(pstate.types).Hmap != nil {
		return t.MapType(pstate.types).Hmap
	}

	bmap := pstate.bmap(t)

	// build a struct:
	// type hmap struct {
	//    count      int
	//    flags      uint8
	//    B          uint8
	//    noverflow  uint16
	//    hash0      uint32
	//    buckets    *bmap
	//    oldbuckets *bmap
	//    nevacuate  uintptr
	//    extra      unsafe.Pointer // *mapextra
	// }
	// must match runtime/map.go:hmap.
	fields := []*types.Field{
		pstate.makefield("count", pstate.types.Types[TINT]),
		pstate.makefield("flags", pstate.types.Types[TUINT8]),
		pstate.makefield("B", pstate.types.Types[TUINT8]),
		pstate.makefield("noverflow", pstate.types.Types[TUINT16]),
		pstate.makefield("hash0", pstate.types.Types[TUINT32]), // Used in walk.go for OMAKEMAP.
		pstate.makefield("buckets", pstate.types.NewPtr(bmap)), // Used in walk.go for OMAKEMAP.
		pstate.makefield("oldbuckets", pstate.types.NewPtr(bmap)),
		pstate.makefield("nevacuate", pstate.types.Types[TUINTPTR]),
		pstate.makefield("extra", pstate.types.Types[TUNSAFEPTR]),
	}

	hmap := types.New(TSTRUCT)
	hmap.SetNoalg(true)
	hmap.SetFields(pstate.types, fields)
	pstate.dowidth(hmap)

	// The size of hmap should be 48 bytes on 64 bit
	// and 28 bytes on 32 bit platforms.
	if size := int64(8 + 5*pstate.Widthptr); hmap.Width != size {
		pstate.Fatalf("hmap size not correct: got %d, want %d", hmap.Width, size)
	}

	t.MapType(pstate.types).Hmap = hmap
	hmap.StructType(pstate.types).Map = t
	return hmap
}

// hiter builds a type representing an Hiter structure for the given map type.
// Make sure this stays in sync with runtime/map.go.
func (pstate *PackageState) hiter(t *types.Type) *types.Type {
	if t.MapType(pstate.types).Hiter != nil {
		return t.MapType(pstate.types).Hiter
	}

	hmap := pstate.hmap(t)
	bmap := pstate.bmap(t)

	// build a struct:
	// type hiter struct {
	//    key         *Key
	//    val         *Value
	//    t           unsafe.Pointer // *MapType
	//    h           *hmap
	//    buckets     *bmap
	//    bptr        *bmap
	//    overflow    unsafe.Pointer // *[]*bmap
	//    oldoverflow unsafe.Pointer // *[]*bmap
	//    startBucket uintptr
	//    offset      uint8
	//    wrapped     bool
	//    B           uint8
	//    i           uint8
	//    bucket      uintptr
	//    checkBucket uintptr
	// }
	// must match runtime/map.go:hiter.
	fields := []*types.Field{
		pstate.makefield("key", pstate.types.NewPtr(t.Key(pstate.types))),  // Used in range.go for TMAP.
		pstate.makefield("val", pstate.types.NewPtr(t.Elem(pstate.types))), // Used in range.go for TMAP.
		pstate.makefield("t", pstate.types.Types[TUNSAFEPTR]),
		pstate.makefield("h", pstate.types.NewPtr(hmap)),
		pstate.makefield("buckets", pstate.types.NewPtr(bmap)),
		pstate.makefield("bptr", pstate.types.NewPtr(bmap)),
		pstate.makefield("overflow", pstate.types.Types[TUNSAFEPTR]),
		pstate.makefield("oldoverflow", pstate.types.Types[TUNSAFEPTR]),
		pstate.makefield("startBucket", pstate.types.Types[TUINTPTR]),
		pstate.makefield("offset", pstate.types.Types[TUINT8]),
		pstate.makefield("wrapped", pstate.types.Types[TBOOL]),
		pstate.makefield("B", pstate.types.Types[TUINT8]),
		pstate.makefield("i", pstate.types.Types[TUINT8]),
		pstate.makefield("bucket", pstate.types.Types[TUINTPTR]),
		pstate.makefield("checkBucket", pstate.types.Types[TUINTPTR]),
	}

	// build iterator struct holding the above fields
	hiter := types.New(TSTRUCT)
	hiter.SetNoalg(true)
	hiter.SetFields(pstate.types, fields)
	pstate.dowidth(hiter)
	if hiter.Width != int64(12*pstate.Widthptr) {
		pstate.Fatalf("hash_iter size not correct %d %d", hiter.Width, 12*pstate.Widthptr)
	}
	t.MapType(pstate.types).Hiter = hiter
	hiter.StructType(pstate.types).Map = t
	return hiter
}

// f is method type, with receiver.
// return function type, receiver as first argument (or not).
func (pstate *PackageState) methodfunc(f *types.Type, receiver *types.Type) *types.Type {
	var in []*Node
	if receiver != nil {
		d := pstate.anonfield(receiver)
		in = append(in, d)
	}

	for _, t := range f.Params(pstate.types).Fields(pstate.types).Slice() {
		d := pstate.anonfield(t.Type)
		d.SetIsddd(t.Isddd())
		in = append(in, d)
	}

	var out []*Node
	for _, t := range f.Results(pstate.types).Fields(pstate.types).Slice() {
		d := pstate.anonfield(t.Type)
		out = append(out, d)
	}

	t := pstate.functype(nil, in, out)
	if f.Nname(pstate.types) != nil {
		// Link to name of original method function.
		t.SetNname(pstate.types, f.Nname(pstate.types))
	}

	return t
}

// methods returns the methods of the non-interface type t, sorted by name.
// Generates stub functions as needed.
func (pstate *PackageState) methods(t *types.Type) []*Sig {
	// method type
	mt := pstate.methtype(t)

	if mt == nil {
		return nil
	}
	pstate.expandmeth(mt)

	// type stored in interface word
	it := t

	if !pstate.isdirectiface(it) {
		it = pstate.types.NewPtr(t)
	}

	// make list of methods for t,
	// generating code if necessary.
	var ms []*Sig
	for _, f := range mt.AllMethods().Slice() {
		if f.Type.Etype != TFUNC || f.Type.Recv(pstate.types) == nil {
			pstate.Fatalf("non-method on %v method %v %v\n", mt, f.Sym, f)
		}
		if f.Type.Recv(pstate.types) == nil {
			pstate.Fatalf("receiver with no type on %v method %v %v\n", mt, f.Sym, f)
		}
		if f.Nointerface() {
			continue
		}

		method := f.Sym
		if method == nil {
			break
		}

		// get receiver type for this particular method.
		// if pointer receiver but non-pointer t and
		// this is not an embedded pointer inside a struct,
		// method does not apply.
		if !pstate.isMethodApplicable(t, f) {
			continue
		}

		sig := &Sig{
			name:  method,
			isym:  pstate.methodSym(it, method),
			tsym:  pstate.methodSym(t, method),
			type_: pstate.methodfunc(f.Type, t),
			mtype: pstate.methodfunc(f.Type, nil),
		}
		ms = append(ms, sig)

		this := f.Type.Recv(pstate.types).Type

		if !sig.isym.Siggen() {
			sig.isym.SetSiggen(true)
			if !pstate.eqtype(this, it) {
				pstate.compiling_wrappers = true
				pstate.genwrapper(it, f, sig.isym)
				pstate.compiling_wrappers = false
			}
		}

		if !sig.tsym.Siggen() {
			sig.tsym.SetSiggen(true)
			if !pstate.eqtype(this, t) {
				pstate.compiling_wrappers = true
				pstate.genwrapper(t, f, sig.tsym)
				pstate.compiling_wrappers = false
			}
		}
	}

	return ms
}

// imethods returns the methods of the interface type t, sorted by name.
func (pstate *PackageState) imethods(t *types.Type) []*Sig {
	var methods []*Sig
	for _, f := range t.Fields(pstate.types).Slice() {
		if f.Type.Etype != TFUNC || f.Sym == nil {
			continue
		}
		if f.Sym.IsBlank() {
			pstate.Fatalf("unexpected blank symbol in interface method set")
		}
		if n := len(methods); n > 0 {
			last := methods[n-1]
			if !last.name.Less(f.Sym) {
				pstate.Fatalf("sigcmp vs sortinter %v %v", last.name, f.Sym)
			}
		}

		sig := &Sig{
			name:  f.Sym,
			mtype: f.Type,
			type_: pstate.methodfunc(f.Type, nil),
		}
		methods = append(methods, sig)

		// NOTE(rsc): Perhaps an oversight that
		// IfaceType.Method is not in the reflect data.
		// Generate the method body, so that compiled
		// code can refer to it.
		isym := pstate.methodSym(t, f.Sym)
		if !isym.Siggen() {
			isym.SetSiggen(true)
			pstate.genwrapper(t, f, isym)
		}
	}

	return methods
}

func (pstate *PackageState) dimportpath(p *types.Pkg) {
	if p.Pathsym != nil {
		return
	}

	// If we are compiling the runtime package, there are two runtime packages around
	// -- localpkg and Runtimepkg. We don't want to produce import path symbols for
	// both of them, so just produce one for localpkg.
	if pstate.myimportpath == "runtime" && p == pstate.Runtimepkg {
		return
	}

	var str string
	if p == pstate.localpkg {
		// Note: myimportpath != "", or else dgopkgpath won't call dimportpath.
		str = pstate.myimportpath
	} else {
		str = p.Path
	}

	s := pstate.Ctxt.Lookup("type..importpath." + p.Prefix + ".")
	ot := pstate.dnameData(s, 0, str, "", nil, false)
	pstate.ggloblsym(s, int32(ot), obj.DUPOK|obj.RODATA)
	p.Pathsym = s
}

func (pstate *PackageState) dgopkgpath(s *obj.LSym, ot int, pkg *types.Pkg) int {
	if pkg == nil {
		return pstate.duintptr(s, ot, 0)
	}

	if pkg == pstate.localpkg && pstate.myimportpath == "" {
		// If we don't know the full import path of the package being compiled
		// (i.e. -p was not passed on the compiler command line), emit a reference to
		// type..importpath.""., which the linker will rewrite using the correct import path.
		// Every package that imports this one directly defines the symbol.
		// See also https://groups.google.com/forum/#!topic/golang-dev/myb9s53HxGQ.
		ns := pstate.Ctxt.Lookup("type..importpath.\"\".")
		return pstate.dsymptr(s, ot, ns, 0)
	}

	pstate.dimportpath(pkg)
	return pstate.dsymptr(s, ot, pkg.Pathsym, 0)
}

// dgopkgpathOff writes an offset relocation in s at offset ot to the pkg path symbol.
func (pstate *PackageState) dgopkgpathOff(s *obj.LSym, ot int, pkg *types.Pkg) int {
	if pkg == nil {
		return pstate.duint32(s, ot, 0)
	}
	if pkg == pstate.localpkg && pstate.myimportpath == "" {
		// If we don't know the full import path of the package being compiled
		// (i.e. -p was not passed on the compiler command line), emit a reference to
		// type..importpath.""., which the linker will rewrite using the correct import path.
		// Every package that imports this one directly defines the symbol.
		// See also https://groups.google.com/forum/#!topic/golang-dev/myb9s53HxGQ.
		ns := pstate.Ctxt.Lookup("type..importpath.\"\".")
		return pstate.dsymptrOff(s, ot, ns)
	}

	pstate.dimportpath(pkg)
	return pstate.dsymptrOff(s, ot, pkg.Pathsym)
}

// dnameField dumps a reflect.name for a struct field.
func (pstate *PackageState) dnameField(lsym *obj.LSym, ot int, spkg *types.Pkg, ft *types.Field) int {
	if !types.IsExported(ft.Sym.Name) && ft.Sym.Pkg != spkg {
		pstate.Fatalf("package mismatch for %v", ft.Sym)
	}
	nsym := pstate.dname(ft.Sym.Name, ft.Note, nil, types.IsExported(ft.Sym.Name))
	return pstate.dsymptr(lsym, ot, nsym, 0)
}

// dnameData writes the contents of a reflect.name into s at offset ot.
func (pstate *PackageState) dnameData(s *obj.LSym, ot int, name, tag string, pkg *types.Pkg, exported bool) int {
	if len(name) > 1<<16-1 {
		pstate.Fatalf("name too long: %s", name)
	}
	if len(tag) > 1<<16-1 {
		pstate.Fatalf("tag too long: %s", tag)
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

	ot = int(s.WriteBytes(pstate.Ctxt, int64(ot), b))

	if pkg != nil {
		ot = pstate.dgopkgpathOff(s, ot, pkg)
	}

	return ot
}

// dname creates a reflect.name for a struct field or method.
func (pstate *PackageState) dname(name, tag string, pkg *types.Pkg, exported bool) *obj.LSym {
	// Write out data as "type.." to signal two things to the
	// linker, first that when dynamically linking, the symbol
	// should be moved to a relro section, and second that the
	// contents should not be decoded as a type.
	sname := "type..namedata."
	if pkg == nil {
		// In the common case, share data with other packages.
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
		sname = fmt.Sprintf("%s\"\".%d", sname, pstate.dnameCount)
		pstate.dnameCount++
	}
	s := pstate.Ctxt.Lookup(sname)
	if len(s.P) > 0 {
		return s
	}
	ot := pstate.dnameData(s, 0, name, tag, pkg, exported)
	pstate.ggloblsym(s, int32(ot), obj.DUPOK|obj.RODATA)
	return s
}

// dextratype dumps the fields of a runtime.uncommontype.
// dataAdd is the offset in bytes after the header where the
// backing array of the []method field is written (by dextratypeData).
func (pstate *PackageState) dextratype(lsym *obj.LSym, ot int, t *types.Type, dataAdd int) int {
	m := pstate.methods(t)
	if t.Sym == nil && len(m) == 0 {
		return ot
	}
	noff := int(pstate.Rnd(int64(ot), int64(pstate.Widthptr)))
	if noff != ot {
		pstate.Fatalf("unexpected alignment in dextratype for %v", t)
	}

	for _, a := range m {
		pstate.dtypesym(a.type_)
	}

	ot = pstate.dgopkgpathOff(lsym, ot, pstate.typePkg(t))

	dataAdd += pstate.uncommonSize(t)
	mcount := len(m)
	if mcount != int(uint16(mcount)) {
		pstate.Fatalf("too many methods on %v: %d", t, mcount)
	}
	xcount := sort.Search(mcount, func(i int) bool { return !types.IsExported(m[i].name.Name) })
	if dataAdd != int(uint32(dataAdd)) {
		pstate.Fatalf("methods are too far away on %v: %d", t, dataAdd)
	}

	ot = pstate.duint16(lsym, ot, uint16(mcount))
	ot = pstate.duint16(lsym, ot, uint16(xcount))
	ot = pstate.duint32(lsym, ot, uint32(dataAdd))
	ot = pstate.duint32(lsym, ot, 0)
	return ot
}

func (pstate *PackageState) typePkg(t *types.Type) *types.Pkg {
	tsym := t.Sym
	if tsym == nil {
		switch t.Etype {
		case TARRAY, TSLICE, TPTR32, TPTR64, TCHAN:
			if t.Elem(pstate.types) != nil {
				tsym = t.Elem(pstate.types).Sym
			}
		}
	}
	if tsym != nil && t != pstate.types.Types[t.Etype] && t != pstate.types.Errortype {
		return tsym.Pkg
	}
	return nil
}

// dextratypeData dumps the backing array for the []method field of
// runtime.uncommontype.
func (pstate *PackageState) dextratypeData(lsym *obj.LSym, ot int, t *types.Type) int {
	for _, a := range pstate.methods(t) {
		// ../../../../runtime/type.go:/method
		exported := types.IsExported(a.name.Name)
		var pkg *types.Pkg
		if !exported && a.name.Pkg != pstate.typePkg(t) {
			pkg = a.name.Pkg
		}
		nsym := pstate.dname(a.name.Name, "", pkg, exported)

		ot = pstate.dsymptrOff(lsym, ot, nsym)
		ot = pstate.dmethodptrOff(lsym, ot, pstate.dtypesym(a.mtype))
		ot = pstate.dmethodptrOff(lsym, ot, a.isym.Linksym(pstate.types))
		ot = pstate.dmethodptrOff(lsym, ot, a.tsym.Linksym(pstate.types))
	}
	return ot
}

func (pstate *PackageState) dmethodptrOff(s *obj.LSym, ot int, x *obj.LSym) int {
	pstate.duint32(s, ot, 0)
	r := obj.Addrel(s)
	r.Off = int32(ot)
	r.Siz = 4
	r.Sym = x
	r.Type = objabi.R_METHODOFF
	return ot + 4
}

// typeptrdata returns the length in bytes of the prefix of t
// containing pointer data. Anything after this offset is scalar data.
func (pstate *PackageState) typeptrdata(t *types.Type) int64 {
	if !pstate.types.Haspointers(t) {
		return 0
	}

	switch t.Etype {
	case TPTR32,
		TPTR64,
		TUNSAFEPTR,
		TFUNC,
		TCHAN,
		TMAP:
		return int64(pstate.Widthptr)

	case TSTRING:
		// struct { byte *str; intgo len; }
		return int64(pstate.Widthptr)

	case TINTER:
		// struct { Itab *tab;	void *data; } or
		// struct { Type *type; void *data; }
		// Note: see comment in plive.go:onebitwalktype1.
		return 2 * int64(pstate.Widthptr)

	case TSLICE:
		// struct { byte *array; uintgo len; uintgo cap; }
		return int64(pstate.Widthptr)

	case TARRAY:
		// haspointers already eliminated t.NumElem() == 0.
		return (t.NumElem(pstate.types)-1)*t.Elem(pstate.types).Width + pstate.typeptrdata(t.Elem(pstate.types))

	case TSTRUCT:
		// Find the last field that has pointers.
		var lastPtrField *types.Field
		for _, t1 := range t.Fields(pstate.types).Slice() {
			if pstate.types.Haspointers(t1.Type) {
				lastPtrField = t1
			}
		}
		return lastPtrField.Offset + pstate.typeptrdata(lastPtrField.Type)

	default:
		pstate.Fatalf("typeptrdata: unexpected type, %v", t)
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
func (pstate *PackageState) dcommontype(lsym *obj.LSym, t *types.Type) int {
	sizeofAlg := 2 * pstate.Widthptr
	if pstate.algarray == nil {
		pstate.algarray = pstate.sysfunc("algarray")
	}
	pstate.dowidth(t)
	alg := pstate.algtype(t)
	var algsym *obj.LSym
	if alg == ASPECIAL || alg == AMEM {
		algsym = pstate.dalgsym(t)
	}

	sptrWeak := true
	var sptr *obj.LSym
	if !t.IsPtr() || t.PtrBase != nil {
		tptr := pstate.types.NewPtr(t)
		if t.Sym != nil || pstate.methods(tptr) != nil {
			sptrWeak = false
		}
		sptr = pstate.dtypesym(tptr)
	}

	gcsym, useGCProg, ptrdata := pstate.dgcsym(t)

	// ../../../../reflect/type.go:/^type.rtype
	// actual type structure
	//	type rtype struct {
	//		size          uintptr
	//		ptrdata       uintptr
	//		hash          uint32
	//		tflag         tflag
	//		align         uint8
	//		fieldAlign    uint8
	//		kind          uint8
	//		alg           *typeAlg
	//		gcdata        *byte
	//		str           nameOff
	//		ptrToThis     typeOff
	//	}
	ot := 0
	ot = pstate.duintptr(lsym, ot, uint64(t.Width))
	ot = pstate.duintptr(lsym, ot, uint64(ptrdata))
	ot = pstate.duint32(lsym, ot, pstate.typehash(t))

	var tflag uint8
	if pstate.uncommonSize(t) != 0 {
		tflag |= tflagUncommon
	}
	if t.Sym != nil && t.Sym.Name != "" {
		tflag |= tflagNamed
	}

	exported := false
	p := t.LongString(pstate.types)
	// If we're writing out type T,
	// we are very likely to write out type *T as well.
	// Use the string "*T"[1:] for "T", so that the two
	// share storage. This is a cheap way to reduce the
	// amount of space taken up by reflect strings.
	if !strings.HasPrefix(p, "*") {
		p = "*" + p
		tflag |= tflagExtraStar
		if t.Sym != nil {
			exported = types.IsExported(t.Sym.Name)
		}
	} else {
		if t.Elem(pstate.types) != nil && t.Elem(pstate.types).Sym != nil {
			exported = types.IsExported(t.Elem(pstate.types).Sym.Name)
		}
	}

	ot = pstate.duint8(lsym, ot, tflag)

	// runtime (and common sense) expects alignment to be a power of two.
	i := int(t.Align)

	if i == 0 {
		i = 1
	}
	if i&(i-1) != 0 {
		pstate.Fatalf("invalid alignment %d for %v", t.Align, t)
	}
	ot = pstate.duint8(lsym, ot, t.Align) // align
	ot = pstate.duint8(lsym, ot, t.Align) // fieldAlign

	i = pstate.kinds[t.Etype]
	if !pstate.types.Haspointers(t) {
		i |= objabi.KindNoPointers
	}
	if pstate.isdirectiface(t) {
		i |= objabi.KindDirectIface
	}
	if useGCProg {
		i |= objabi.KindGCProg
	}
	ot = pstate.duint8(lsym, ot, uint8(i)) // kind
	if algsym == nil {
		ot = pstate.dsymptr(lsym, ot, pstate.algarray, int(alg)*sizeofAlg)
	} else {
		ot = pstate.dsymptr(lsym, ot, algsym, 0)
	}
	ot = pstate.dsymptr(lsym, ot, gcsym, 0) // gcdata

	nsym := pstate.dname(p, "", nil, exported)
	ot = pstate.dsymptrOff(lsym, ot, nsym) // str
	// ptrToThis
	if sptr == nil {
		ot = pstate.duint32(lsym, ot, 0)
	} else if sptrWeak {
		ot = pstate.dsymptrWeakOff(lsym, ot, sptr)
	} else {
		ot = pstate.dsymptrOff(lsym, ot, sptr)
	}

	return ot
}

// typeHasNoAlg returns whether t does not have any associated hash/eq
// algorithms because t, or some component of t, is marked Noalg.
func (pstate *PackageState) typeHasNoAlg(t *types.Type) bool {
	a, bad := pstate.algtype1(t)
	return a == ANOEQ && bad.Noalg()
}

func (pstate *PackageState) typesymname(t *types.Type) string {
	name := t.ShortString(pstate.types)
	// Use a separate symbol name for Noalg types for #17752.
	if pstate.typeHasNoAlg(t) {
		name = "noalg." + name
	}
	return name
}

func (pstate *PackageState) typeLookup(name string) *types.Sym {
	pstate.typepkgmu.Lock()
	s := pstate.typepkg.Lookup(pstate.types, name)
	pstate.typepkgmu.Unlock()
	return s
}

func (pstate *PackageState) typesym(t *types.Type) *types.Sym {
	return pstate.typeLookup(pstate.typesymname(t))
}

// tracksym returns the symbol for tracking use of field/method f, assumed
// to be a member of struct/interface type t.
func (pstate *PackageState) tracksym(t *types.Type, f *types.Field) *types.Sym {
	return pstate.trackpkg.Lookup(pstate.types, t.ShortString(pstate.types)+"."+f.Sym.Name)
}

func (pstate *PackageState) typesymprefix(prefix string, t *types.Type) *types.Sym {
	p := prefix + "." + t.ShortString(pstate.types)
	s := pstate.typeLookup(p)

	//print("algsym: %s -> %+S\n", p, s);

	return s
}

func (pstate *PackageState) typenamesym(t *types.Type) *types.Sym {
	if t == nil || (t.IsPtr() && t.Elem(pstate.types) == nil) || t.IsUntyped(pstate.types) {
		pstate.Fatalf("typenamesym %v", t)
	}
	s := pstate.typesym(t)
	pstate.signatsetmu.Lock()
	pstate.addsignat(t)
	pstate.signatsetmu.Unlock()
	return s
}

func (pstate *PackageState) typename(t *types.Type) *Node {
	s := pstate.typenamesym(t)
	if s.Def == nil {
		n := pstate.newnamel(pstate.src.NoXPos, s)
		n.Type = pstate.types.Types[TUINT8]
		n.SetClass(PEXTERN)
		n.SetTypecheck(1)
		s.Def = asTypesNode(n)
	}

	n := pstate.nod(OADDR, asNode(s.Def), nil)
	n.Type = pstate.types.NewPtr(asNode(s.Def).Type)
	n.SetAddable(true)
	n.SetTypecheck(1)
	return n
}

func (pstate *PackageState) itabname(t, itype *types.Type) *Node {
	if t == nil || (t.IsPtr() && t.Elem(pstate.types) == nil) || t.IsUntyped(pstate.types) || !itype.IsInterface() || itype.IsEmptyInterface(pstate.types) {
		pstate.Fatalf("itabname(%v, %v)", t, itype)
	}
	s := pstate.itabpkg.Lookup(pstate.types, t.ShortString(pstate.types)+","+itype.ShortString(pstate.types))
	if s.Def == nil {
		n := pstate.newname(s)
		n.Type = pstate.types.Types[TUINT8]
		n.SetClass(PEXTERN)
		n.SetTypecheck(1)
		s.Def = asTypesNode(n)
		pstate.itabs = append(pstate.itabs, itabEntry{t: t, itype: itype, lsym: s.Linksym(pstate.types)})
	}

	n := pstate.nod(OADDR, asNode(s.Def), nil)
	n.Type = pstate.types.NewPtr(asNode(s.Def).Type)
	n.SetAddable(true)
	n.SetTypecheck(1)
	return n
}

// isreflexive reports whether t has a reflexive equality operator.
// That is, if x==x for all x of type t.
func (pstate *PackageState) isreflexive(t *types.Type) bool {
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
		return pstate.isreflexive(t.Elem(pstate.types))

	case TSTRUCT:
		for _, t1 := range t.Fields(pstate.types).Slice() {
			if !pstate.isreflexive(t1.Type) {
				return false
			}
		}
		return true

	default:
		pstate.Fatalf("bad type for map key: %v", t)
		return false
	}
}

// needkeyupdate reports whether map updates with t as a key
// need the key to be updated.
func (pstate *PackageState) needkeyupdate(t *types.Type) bool {
	switch t.Etype {
	case TBOOL, TINT, TUINT, TINT8, TUINT8, TINT16, TUINT16, TINT32, TUINT32,
		TINT64, TUINT64, TUINTPTR, TPTR32, TPTR64, TUNSAFEPTR, TCHAN:
		return false

	case TFLOAT32, TFLOAT64, TCOMPLEX64, TCOMPLEX128, // floats and complex can be +0/-0
		TINTER,
		TSTRING: // strings might have smaller backing stores
		return true

	case TARRAY:
		return pstate.needkeyupdate(t.Elem(pstate.types))

	case TSTRUCT:
		for _, t1 := range t.Fields(pstate.types).Slice() {
			if pstate.needkeyupdate(t1.Type) {
				return true
			}
		}
		return false

	default:
		pstate.Fatalf("bad type for map key: %v", t)
		return true
	}
}

// formalType replaces byte and rune aliases with real types.
// They've been separate internally to make error messages
// better, but we have to merge them in the reflect tables.
func (pstate *PackageState) formalType(t *types.Type) *types.Type {
	if t == pstate.types.Bytetype || t == pstate.types.Runetype {
		return pstate.types.Types[t.Etype]
	}
	return t
}

func (pstate *PackageState) dtypesym(t *types.Type) *obj.LSym {
	t = pstate.formalType(t)
	if t.IsUntyped(pstate.types) {
		pstate.Fatalf("dtypesym %v", t)
	}

	s := pstate.typesym(t)
	lsym := s.Linksym(pstate.types)
	if s.Siggen() {
		return lsym
	}
	s.SetSiggen(true)

	// special case (look for runtime below):
	// when compiling package runtime,
	// emit the type structures for int, float, etc.
	tbase := t

	if t.IsPtr() && t.Sym == nil && t.Elem(pstate.types).Sym != nil {
		tbase = t.Elem(pstate.types)
	}
	dupok := 0
	if tbase.Sym == nil {
		dupok = obj.DUPOK
	}

	if pstate.myimportpath != "runtime" || (tbase != pstate.types.Types[tbase.Etype] && tbase != pstate.types.Bytetype && tbase != pstate.types.Runetype && tbase != pstate.types.Errortype) { // int, float, etc
		// named types from other files are defined only by those files
		if tbase.Sym != nil && tbase.Sym.Pkg != pstate.localpkg {
			return lsym
		}
		// TODO(mdempsky): Investigate whether this can happen.
		if pstate.isforw[tbase.Etype] {
			return lsym
		}
	}

	ot := 0
	switch t.Etype {
	default:
		ot = pstate.dcommontype(lsym, t)
		ot = pstate.dextratype(lsym, ot, t, 0)

	case TARRAY:
		// ../../../../runtime/type.go:/arrayType
		s1 := pstate.dtypesym(t.Elem(pstate.types))
		t2 := pstate.types.NewSlice(t.Elem(pstate.types))
		s2 := pstate.dtypesym(t2)
		ot = pstate.dcommontype(lsym, t)
		ot = pstate.dsymptr(lsym, ot, s1, 0)
		ot = pstate.dsymptr(lsym, ot, s2, 0)
		ot = pstate.duintptr(lsym, ot, uint64(t.NumElem(pstate.types)))
		ot = pstate.dextratype(lsym, ot, t, 0)

	case TSLICE:
		// ../../../../runtime/type.go:/sliceType
		s1 := pstate.dtypesym(t.Elem(pstate.types))
		ot = pstate.dcommontype(lsym, t)
		ot = pstate.dsymptr(lsym, ot, s1, 0)
		ot = pstate.dextratype(lsym, ot, t, 0)

	case TCHAN:
		// ../../../../runtime/type.go:/chanType
		s1 := pstate.dtypesym(t.Elem(pstate.types))
		ot = pstate.dcommontype(lsym, t)
		ot = pstate.dsymptr(lsym, ot, s1, 0)
		ot = pstate.duintptr(lsym, ot, uint64(t.ChanDir(pstate.types)))
		ot = pstate.dextratype(lsym, ot, t, 0)

	case TFUNC:
		for _, t1 := range t.Recvs(pstate.types).Fields(pstate.types).Slice() {
			pstate.dtypesym(t1.Type)
		}
		isddd := false
		for _, t1 := range t.Params(pstate.types).Fields(pstate.types).Slice() {
			isddd = t1.Isddd()
			pstate.dtypesym(t1.Type)
		}
		for _, t1 := range t.Results(pstate.types).Fields(pstate.types).Slice() {
			pstate.dtypesym(t1.Type)
		}

		ot = pstate.dcommontype(lsym, t)
		inCount := t.NumRecvs(pstate.types) + t.NumParams(pstate.types)
		outCount := t.NumResults(pstate.types)
		if isddd {
			outCount |= 1 << 15
		}
		ot = pstate.duint16(lsym, ot, uint16(inCount))
		ot = pstate.duint16(lsym, ot, uint16(outCount))
		if pstate.Widthptr == 8 {
			ot += 4 // align for *rtype
		}

		dataAdd := (inCount + t.NumResults(pstate.types)) * pstate.Widthptr
		ot = pstate.dextratype(lsym, ot, t, dataAdd)

		// Array of rtype pointers follows funcType.
		for _, t1 := range t.Recvs(pstate.types).Fields(pstate.types).Slice() {
			ot = pstate.dsymptr(lsym, ot, pstate.dtypesym(t1.Type), 0)
		}
		for _, t1 := range t.Params(pstate.types).Fields(pstate.types).Slice() {
			ot = pstate.dsymptr(lsym, ot, pstate.dtypesym(t1.Type), 0)
		}
		for _, t1 := range t.Results(pstate.types).Fields(pstate.types).Slice() {
			ot = pstate.dsymptr(lsym, ot, pstate.dtypesym(t1.Type), 0)
		}

	case TINTER:
		m := pstate.imethods(t)
		n := len(m)
		for _, a := range m {
			pstate.dtypesym(a.type_)
		}

		// ../../../../runtime/type.go:/interfaceType
		ot = pstate.dcommontype(lsym, t)

		var tpkg *types.Pkg
		if t.Sym != nil && t != pstate.types.Types[t.Etype] && t != pstate.types.Errortype {
			tpkg = t.Sym.Pkg
		}
		ot = pstate.dgopkgpath(lsym, ot, tpkg)

		ot = pstate.dsymptr(lsym, ot, lsym, ot+3*pstate.Widthptr+pstate.uncommonSize(t))
		ot = pstate.duintptr(lsym, ot, uint64(n))
		ot = pstate.duintptr(lsym, ot, uint64(n))
		dataAdd := imethodSize() * n
		ot = pstate.dextratype(lsym, ot, t, dataAdd)

		for _, a := range m {
			// ../../../../runtime/type.go:/imethod
			exported := types.IsExported(a.name.Name)
			var pkg *types.Pkg
			if !exported && a.name.Pkg != tpkg {
				pkg = a.name.Pkg
			}
			nsym := pstate.dname(a.name.Name, "", pkg, exported)

			ot = pstate.dsymptrOff(lsym, ot, nsym)
			ot = pstate.dsymptrOff(lsym, ot, pstate.dtypesym(a.type_))
		}

	// ../../../../runtime/type.go:/mapType
	case TMAP:
		s1 := pstate.dtypesym(t.Key(pstate.types))
		s2 := pstate.dtypesym(t.Elem(pstate.types))
		s3 := pstate.dtypesym(pstate.bmap(t))
		ot = pstate.dcommontype(lsym, t)
		ot = pstate.dsymptr(lsym, ot, s1, 0)
		ot = pstate.dsymptr(lsym, ot, s2, 0)
		ot = pstate.dsymptr(lsym, ot, s3, 0)
		if t.Key(pstate.types).Width > MAXKEYSIZE {
			ot = pstate.duint8(lsym, ot, uint8(pstate.Widthptr))
			ot = pstate.duint8(lsym, ot, 1) // indirect
		} else {
			ot = pstate.duint8(lsym, ot, uint8(t.Key(pstate.types).Width))
			ot = pstate.duint8(lsym, ot, 0) // not indirect
		}

		if t.Elem(pstate.types).Width > MAXVALSIZE {
			ot = pstate.duint8(lsym, ot, uint8(pstate.Widthptr))
			ot = pstate.duint8(lsym, ot, 1) // indirect
		} else {
			ot = pstate.duint8(lsym, ot, uint8(t.Elem(pstate.types).Width))
			ot = pstate.duint8(lsym, ot, 0) // not indirect
		}

		ot = pstate.duint16(lsym, ot, uint16(pstate.bmap(t).Width))
		ot = pstate.duint8(lsym, ot, uint8(obj.Bool2int(pstate.isreflexive(t.Key(pstate.types)))))
		ot = pstate.duint8(lsym, ot, uint8(obj.Bool2int(pstate.needkeyupdate(t.Key(pstate.types)))))
		ot = pstate.dextratype(lsym, ot, t, 0)

	case TPTR32, TPTR64:
		if t.Elem(pstate.types).Etype == TANY {
			// ../../../../runtime/type.go:/UnsafePointerType
			ot = pstate.dcommontype(lsym, t)
			ot = pstate.dextratype(lsym, ot, t, 0)

			break
		}

		// ../../../../runtime/type.go:/ptrType
		s1 := pstate.dtypesym(t.Elem(pstate.types))

		ot = pstate.dcommontype(lsym, t)
		ot = pstate.dsymptr(lsym, ot, s1, 0)
		ot = pstate.dextratype(lsym, ot, t, 0)

	// ../../../../runtime/type.go:/structType
	// for security, only the exported fields.
	case TSTRUCT:
		fields := t.Fields(pstate.types).Slice()
		for _, t1 := range fields {
			pstate.dtypesym(t1.Type)
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

		ot = pstate.dcommontype(lsym, t)
		ot = pstate.dgopkgpath(lsym, ot, spkg)
		ot = pstate.dsymptr(lsym, ot, lsym, ot+3*pstate.Widthptr+pstate.uncommonSize(t))
		ot = pstate.duintptr(lsym, ot, uint64(len(fields)))
		ot = pstate.duintptr(lsym, ot, uint64(len(fields)))

		dataAdd := len(fields) * pstate.structfieldSize()
		ot = pstate.dextratype(lsym, ot, t, dataAdd)

		for _, f := range fields {
			// ../../../../runtime/type.go:/structField
			ot = pstate.dnameField(lsym, ot, spkg, f)
			ot = pstate.dsymptr(lsym, ot, pstate.dtypesym(f.Type), 0)
			offsetAnon := uint64(f.Offset) << 1
			if offsetAnon>>1 != uint64(f.Offset) {
				pstate.Fatalf("%v: bad field offset for %s", t, f.Sym.Name)
			}
			if f.Embedded != 0 {
				offsetAnon |= 1
			}
			ot = pstate.duintptr(lsym, ot, offsetAnon)
		}
	}

	ot = pstate.dextratypeData(lsym, ot, t)
	pstate.ggloblsym(lsym, int32(ot), int16(dupok|obj.RODATA))

	// The linker will leave a table of all the typelinks for
	// types in the binary, so the runtime can find them.
	//
	// When buildmode=shared, all types are in typelinks so the
	// runtime can deduplicate type pointers.
	keep := pstate.Ctxt.Flag_dynlink
	if !keep && t.Sym == nil {
		// For an unnamed type, we only need the link if the type can
		// be created at run time by reflect.PtrTo and similar
		// functions. If the type exists in the program, those
		// functions must return the existing type structure rather
		// than creating a new one.
		switch t.Etype {
		case TPTR32, TPTR64, TARRAY, TCHAN, TFUNC, TMAP, TSLICE, TSTRUCT:
			keep = true
		}
	}
	// Do not put Noalg types in typelinks.  See issue #22605.
	if pstate.typeHasNoAlg(t) {
		keep = false
	}
	lsym.Set(obj.AttrMakeTypelink, keep)

	return lsym
}

// for each itabEntry, gather the methods on
// the concrete type that implement the interface
func (pstate *PackageState) peekitabs() {
	for i := range pstate.itabs {
		tab := &pstate.itabs[i]
		methods := pstate.genfun(tab.t, tab.itype)
		if len(methods) == 0 {
			continue
		}
		tab.entries = methods
	}
}

// for the given concrete type and interface
// type, return the (sorted) set of methods
// on the concrete type that implement the interface
func (pstate *PackageState) genfun(t, it *types.Type) []*obj.LSym {
	if t == nil || it == nil {
		return nil
	}
	sigs := pstate.imethods(it)
	methods := pstate.methods(t)
	out := make([]*obj.LSym, 0, len(sigs))
	// TODO(mdempsky): Short circuit before calling methods(t)?
	// See discussion on CL 105039.
	if len(sigs) == 0 {
		return nil
	}

	// both sigs and methods are sorted by name,
	// so we can find the intersect in a single pass
	for _, m := range methods {
		if m.name == sigs[0].name {
			out = append(out, m.isym.Linksym(pstate.types))
			sigs = sigs[1:]
			if len(sigs) == 0 {
				break
			}
		}
	}

	if len(sigs) != 0 {
		pstate.Fatalf("incomplete itab")
	}

	return out
}

// itabsym uses the information gathered in
// peekitabs to de-virtualize interface methods.
// Since this is called by the SSA backend, it shouldn't
// generate additional Nodes, Syms, etc.
func (pstate *PackageState) itabsym(it *obj.LSym, offset int64) *obj.LSym {
	var syms []*obj.LSym
	if it == nil {
		return nil
	}

	for i := range pstate.itabs {
		e := &pstate.itabs[i]
		if e.lsym == it {
			syms = e.entries
			break
		}
	}
	if syms == nil {
		return nil
	}

	// keep this arithmetic in sync with *itab layout
	methodnum := int((offset - 2*int64(pstate.Widthptr) - 8) / int64(pstate.Widthptr))
	if methodnum >= len(syms) {
		return nil
	}
	return syms[methodnum]
}

// addsignat ensures that a runtime type descriptor is emitted for t.
func (pstate *PackageState) addsignat(t *types.Type) {
	pstate.signatset[t] = struct{}{}
}

func (pstate *PackageState) addsignats(dcls []*Node) {
	// copy types from dcl list to signatset
	for _, n := range dcls {
		if n.Op == OTYPE {
			pstate.addsignat(n.Type)
		}
	}
}

func (pstate *PackageState) dumpsignats() {
	// Process signatset. Use a loop, as dtypesym adds
	// entries to signatset while it is being processed.
	signats := make([]typeAndStr, len(pstate.signatset))
	for len(pstate.signatset) > 0 {
		signats = signats[:0]
		// Transfer entries to a slice and sort, for reproducible builds.
		for t := range pstate.signatset {
			signats = append(signats, typeAndStr{t: t, short: pstate.typesymname(t), regular: t.String(pstate.types)})
			delete(pstate.signatset, t)
		}
		sort.Sort(typesByString(signats))
		for _, ts := range signats {
			t := ts.t
			pstate.dtypesym(t)
			if t.Sym != nil {
				pstate.dtypesym(pstate.types.NewPtr(t))
			}
		}
	}
}

func (pstate *PackageState) dumptabs() {
	// process itabs
	for _, i := range pstate.itabs {
		// dump empty itab symbol into i.sym
		// type itab struct {
		//   inter  *interfacetype
		//   _type  *_type
		//   hash   uint32
		//   _      [4]byte
		//   fun    [1]uintptr // variable sized
		// }
		o := pstate.dsymptr(i.lsym, 0, pstate.dtypesym(i.itype), 0)
		o = pstate.dsymptr(i.lsym, o, pstate.dtypesym(i.t), 0)
		o = pstate.duint32(i.lsym, o, pstate.typehash(i.t)) // copy of type hash
		o += 4                                              // skip unused field
		for _, fn := range pstate.genfun(i.t, i.itype) {
			o = pstate.dsymptr(i.lsym, o, fn, 0) // method pointer for each method
		}
		// Nothing writes static itabs, so they are read only.
		pstate.ggloblsym(i.lsym, int32(o), int16(obj.DUPOK|obj.RODATA))
		ilink := pstate.itablinkpkg.Lookup(pstate.types, i.t.ShortString(pstate.types)+","+i.itype.ShortString(pstate.types)).Linksym(pstate.types)
		pstate.dsymptr(ilink, 0, i.lsym, 0)
		pstate.ggloblsym(ilink, int32(pstate.Widthptr), int16(obj.DUPOK|obj.RODATA))
	}

	// process ptabs
	if pstate.localpkg.Name == "main" && len(pstate.ptabs) > 0 {
		ot := 0
		s := pstate.Ctxt.Lookup("go.plugin.tabs")
		for _, p := range pstate.ptabs {
			// Dump ptab symbol into go.pluginsym package.
			//
			// type ptab struct {
			//	name nameOff
			//	typ  typeOff // pointer to symbol
			// }
			nsym := pstate.dname(p.s.Name, "", nil, true)
			ot = pstate.dsymptrOff(s, ot, nsym)
			ot = pstate.dsymptrOff(s, ot, pstate.dtypesym(p.t))
		}
		pstate.ggloblsym(s, int32(ot), int16(obj.RODATA))

		ot = 0
		s = pstate.Ctxt.Lookup("go.plugin.exports")
		for _, p := range pstate.ptabs {
			ot = pstate.dsymptr(s, ot, p.s.Linksym(pstate.types), 0)
		}
		pstate.ggloblsym(s, int32(ot), int16(obj.RODATA))
	}
}

func (pstate *PackageState) dumpimportstrings() {
	// generate import strings for imported packages
	for _, p := range pstate.types.ImportedPkgList() {
		pstate.dimportpath(p)
	}
}

func (pstate *PackageState) dumpbasictypes() {
	// do basic types if compiling package runtime.
	// they have to be in at least one package,
	// and runtime is always loaded implicitly,
	// so this is as good as any.
	// another possible choice would be package main,
	// but using runtime means fewer copies in object files.
	if pstate.myimportpath == "runtime" {
		for i := types.EType(1); i <= TBOOL; i++ {
			pstate.dtypesym(pstate.types.NewPtr(pstate.types.Types[i]))
		}
		pstate.dtypesym(pstate.types.NewPtr(pstate.types.Types[TSTRING]))
		pstate.dtypesym(pstate.types.NewPtr(pstate.types.Types[TUNSAFEPTR]))

		// emit type structs for error and func(error) string.
		// The latter is the type of an auto-generated wrapper.
		pstate.dtypesym(pstate.types.NewPtr(pstate.types.Errortype))

		pstate.dtypesym(pstate.functype(nil, []*Node{pstate.anonfield(pstate.types.Errortype)}, []*Node{pstate.anonfield(pstate.types.Types[TSTRING])}))

		// add paths for runtime and main, which 6l imports implicitly.
		pstate.dimportpath(pstate.Runtimepkg)

		if pstate.flag_race {
			pstate.dimportpath(pstate.racepkg)
		}
		if pstate.flag_msan {
			pstate.dimportpath(pstate.msanpkg)
		}
		pstate.dimportpath(pstate.types.NewPkg("main", ""))
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
	// When the only difference between the types is whether
	// they refer to byte or uint8, such as **byte vs **uint8,
	// the types' ShortStrings can be identical.
	// To preserve deterministic sort ordering, sort these by String().
	return a[i].regular < a[j].regular
}
func (a typesByString) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (pstate *PackageState) dalgsym(t *types.Type) *obj.LSym {
	var lsym *obj.LSym
	var hashfunc *obj.LSym
	var eqfunc *obj.LSym

	// dalgsym is only called for a type that needs an algorithm table,
	// which implies that the type is comparable (or else it would use ANOEQ).

	if pstate.algtype(t) == AMEM {
		// we use one algorithm table for all AMEM types of a given size
		p := fmt.Sprintf(".alg%d", t.Width)

		s := pstate.typeLookup(p)
		lsym = s.Linksym(pstate.types)
		if s.AlgGen() {
			return lsym
		}
		s.SetAlgGen(true)

		if pstate.memhashvarlen == nil {
			pstate.memhashvarlen = pstate.sysfunc("memhash_varlen")
			pstate.memequalvarlen = pstate.sysfunc("memequal_varlen")
		}

		// make hash closure
		p = fmt.Sprintf(".hashfunc%d", t.Width)

		hashfunc = pstate.typeLookup(p).Linksym(pstate.types)

		ot := 0
		ot = pstate.dsymptr(hashfunc, ot, pstate.memhashvarlen, 0)
		ot = pstate.duintptr(hashfunc, ot, uint64(t.Width)) // size encoded in closure
		pstate.ggloblsym(hashfunc, int32(ot), obj.DUPOK|obj.RODATA)

		// make equality closure
		p = fmt.Sprintf(".eqfunc%d", t.Width)

		eqfunc = pstate.typeLookup(p).Linksym(pstate.types)

		ot = 0
		ot = pstate.dsymptr(eqfunc, ot, pstate.memequalvarlen, 0)
		ot = pstate.duintptr(eqfunc, ot, uint64(t.Width))
		pstate.ggloblsym(eqfunc, int32(ot), obj.DUPOK|obj.RODATA)
	} else {
		// generate an alg table specific to this type
		s := pstate.typesymprefix(".alg", t)
		lsym = s.Linksym(pstate.types)

		hash := pstate.typesymprefix(".hash", t)
		eq := pstate.typesymprefix(".eq", t)
		hashfunc = pstate.typesymprefix(".hashfunc", t).Linksym(pstate.types)
		eqfunc = pstate.typesymprefix(".eqfunc", t).Linksym(pstate.types)

		pstate.genhash(hash, t)
		pstate.geneq(eq, t)

		// make Go funcs (closures) for calling hash and equal from Go
		pstate.dsymptr(hashfunc, 0, hash.Linksym(pstate.types), 0)
		pstate.ggloblsym(hashfunc, int32(pstate.Widthptr), obj.DUPOK|obj.RODATA)
		pstate.dsymptr(eqfunc, 0, eq.Linksym(pstate.types), 0)
		pstate.ggloblsym(eqfunc, int32(pstate.Widthptr), obj.DUPOK|obj.RODATA)
	}

	// ../../../../runtime/alg.go:/typeAlg
	ot := 0

	ot = pstate.dsymptr(lsym, ot, hashfunc, 0)
	ot = pstate.dsymptr(lsym, ot, eqfunc, 0)
	pstate.ggloblsym(lsym, int32(ot), obj.DUPOK|obj.RODATA)
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
func (pstate *PackageState) dgcsym(t *types.Type) (lsym *obj.LSym, useGCProg bool, ptrdata int64) {
	ptrdata = pstate.typeptrdata(t)
	if ptrdata/int64(pstate.Widthptr) <= maxPtrmaskBytes*8 {
		lsym = pstate.dgcptrmask(t)
		return
	}

	useGCProg = true
	lsym, ptrdata = pstate.dgcprog(t)
	return
}

// dgcptrmask emits and returns the symbol containing a pointer mask for type t.
func (pstate *PackageState) dgcptrmask(t *types.Type) *obj.LSym {
	ptrmask := make([]byte, (pstate.typeptrdata(t)/int64(pstate.Widthptr)+7)/8)
	pstate.fillptrmask(t, ptrmask)
	p := fmt.Sprintf("gcbits.%x", ptrmask)

	sym := pstate.Runtimepkg.Lookup(pstate.types, p)
	lsym := sym.Linksym(pstate.types)
	if !sym.Uniq() {
		sym.SetUniq(true)
		for i, x := range ptrmask {
			pstate.duint8(lsym, i, x)
		}
		pstate.ggloblsym(lsym, int32(len(ptrmask)), obj.DUPOK|obj.RODATA|obj.LOCAL)
	}
	return lsym
}

// fillptrmask fills in ptrmask with 1s corresponding to the
// word offsets in t that hold pointers.
// ptrmask is assumed to fit at least typeptrdata(t)/Widthptr bits.
func (pstate *PackageState) fillptrmask(t *types.Type, ptrmask []byte) {
	for i := range ptrmask {
		ptrmask[i] = 0
	}
	if !pstate.types.Haspointers(t) {
		return
	}

	vec := bvalloc(8 * int32(len(ptrmask)))
	pstate.onebitwalktype1(t, 0, vec)

	nptr := pstate.typeptrdata(t) / int64(pstate.Widthptr)
	for i := int64(0); i < nptr; i++ {
		if vec.Get(pstate, int32(i)) {
			ptrmask[i/8] |= 1 << (uint(i) % 8)
		}
	}
}

// dgcprog emits and returns the symbol containing a GC program for type t
// along with the size of the data described by the program (in the range [typeptrdata(t), t.Width]).
// In practice, the size is typeptrdata(t) except for non-trivial arrays.
// For non-trivial arrays, the program describes the full t.Width size.
func (pstate *PackageState) dgcprog(t *types.Type) (*obj.LSym, int64) {
	pstate.dowidth(t)
	if t.Width == BADWIDTH {
		pstate.Fatalf("dgcprog: %v badwidth", t)
	}
	lsym := pstate.typesymprefix(".gcprog", t).Linksym(pstate.types)
	var p GCProg
	p.init(pstate, lsym)
	p.emit(pstate, t, 0)
	offset := p.w.BitIndex() * int64(pstate.Widthptr)
	p.end(pstate)
	if ptrdata := pstate.typeptrdata(t); offset < ptrdata || offset > t.Width {
		pstate.Fatalf("dgcprog: %v: offset=%d but ptrdata=%d size=%d", t, offset, ptrdata, t.Width)
	}
	return lsym, offset
}

type GCProg struct {
	lsym   *obj.LSym
	symoff int
	w      gcprog.Writer
}

func (p *GCProg) init(pstate *PackageState, lsym *obj.LSym) {
	p.lsym = lsym
	p.symoff = 4 // first 4 bytes hold program length
	p.w.Init(p.writeByte)
	if pstate.Debug_gcprog > 0 {
		fmt.Fprintf(os.Stderr, "compile: start GCProg for %v\n", lsym)
		p.w.Debug(os.Stderr)
	}
}

func (p *GCProg) writeByte(pstate *PackageState, x byte) {
	p.symoff = pstate.duint8(p.lsym, p.symoff, x)
}

func (p *GCProg) end(pstate *PackageState) {
	p.w.End()
	pstate.duint32(p.lsym, 0, uint32(p.symoff-4))
	pstate.ggloblsym(p.lsym, int32(p.symoff), obj.DUPOK|obj.RODATA|obj.LOCAL)
	if pstate.Debug_gcprog > 0 {
		fmt.Fprintf(os.Stderr, "compile: end GCProg for %v\n", p.lsym)
	}
}

func (p *GCProg) emit(pstate *PackageState, t *types.Type, offset int64) {
	pstate.dowidth(t)
	if !pstate.types.Haspointers(t) {
		return
	}
	if t.Width == int64(pstate.Widthptr) {
		p.w.Ptr(offset / int64(pstate.Widthptr))
		return
	}
	switch t.Etype {
	default:
		pstate.Fatalf("GCProg.emit: unexpected type %v", t)

	case TSTRING:
		p.w.Ptr(offset / int64(pstate.Widthptr))

	case TINTER:
		// Note: the first word isn't a pointer. See comment in plive.go:onebitwalktype1.
		p.w.Ptr(offset/int64(pstate.Widthptr) + 1)

	case TSLICE:
		p.w.Ptr(offset / int64(pstate.Widthptr))

	case TARRAY:
		if t.NumElem(pstate.types) == 0 {
			// should have been handled by haspointers check above
			pstate.Fatalf("GCProg.emit: empty array")
		}

		// Flatten array-of-array-of-array to just a big array by multiplying counts.
		count := t.NumElem(pstate.types)
		elem := t.Elem(pstate.types)
		for elem.IsArray() {
			count *= elem.NumElem(pstate.types)
			elem = elem.Elem(pstate.types)
		}

		if !p.w.ShouldRepeat(elem.Width/int64(pstate.Widthptr), count) {
			// Cheaper to just emit the bits.
			for i := int64(0); i < count; i++ {
				p.emit(pstate, elem, offset+i*elem.Width)
			}
			return
		}
		p.emit(pstate, elem, offset)
		p.w.ZeroUntil((offset + elem.Width) / int64(pstate.Widthptr))
		p.w.Repeat(elem.Width/int64(pstate.Widthptr), count-1)

	case TSTRUCT:
		for _, t1 := range t.Fields(pstate.types).Slice() {
			p.emit(pstate, t1.Type, offset+t1.Offset)
		}
	}
}

// zeroaddr returns the address of a symbol with at least
// size bytes of zeros.
func (pstate *PackageState) zeroaddr(size int64) *Node {
	if size >= 1<<31 {
		pstate.Fatalf("map value too big %d", size)
	}
	if pstate.zerosize < size {
		pstate.zerosize = size
	}
	s := pstate.mappkg.Lookup(pstate.types, "zero")
	if s.Def == nil {
		x := pstate.newname(s)
		x.Type = pstate.types.Types[TUINT8]
		x.SetClass(PEXTERN)
		x.SetTypecheck(1)
		s.Def = asTypesNode(x)
	}
	z := pstate.nod(OADDR, asNode(s.Def), nil)
	z.Type = pstate.types.NewPtr(pstate.types.Types[TUINT8])
	z.SetAddable(true)
	z.SetTypecheck(1)
	return z
}

package types

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/src"
)

// Dummy Node so we can refer to *Node without actually
// having a gc.Node. Necessary to break import cycles.
// TODO(gri) try to eliminate soon
type Node struct{ _ int }

// EType describes a kind of type.
type EType uint8

const (
	Txxx EType = iota

	TINT8
	TUINT8
	TINT16
	TUINT16
	TINT32
	TUINT32
	TINT64
	TUINT64
	TINT
	TUINT
	TUINTPTR

	TCOMPLEX64
	TCOMPLEX128

	TFLOAT32
	TFLOAT64

	TBOOL

	TPTR32
	TPTR64

	TFUNC
	TSLICE
	TARRAY
	TSTRUCT
	TCHAN
	TMAP
	TINTER
	TFORW
	TANY
	TSTRING
	TUNSAFEPTR

	// pseudo-types for literals
	TIDEAL
	TNIL
	TBLANK

	// pseudo-types for frame layout
	TFUNCARGS
	TCHANARGS

	// pseudo-types for import/export
	TDDDFIELD // wrapper: contained type is a ... field

	// SSA backend types
	TSSA   // internal types used by SSA backend (flags, memory, etc.)
	TTUPLE // a pair of types, used by SSA backend

	NTYPE
)

// ChanDir is whether a channel can send, receive, or both.
type ChanDir uint8

func (c ChanDir) CanRecv() bool { return c&Crecv != 0 }
func (c ChanDir) CanSend() bool { return c&Csend != 0 }

const (
	// types of channel
	// must match ../../../../reflect/type.go:/ChanDir
	Crecv ChanDir = 1 << 0
	Csend ChanDir = 1 << 1
	Cboth ChanDir = Crecv | Csend
)

// Types stores pointers to predeclared named types.
//
// It also stores pointers to several special types:
//   - Types[TANY] is the placeholder "any" type recognized by substArgTypes.
//   - Types[TBLANK] represents the blank variable's type.
//   - Types[TIDEAL] represents untyped numeric constants.
//   - Types[TNIL] represents the predeclared "nil" value's type.
//   - Types[TUNSAFEPTR] is package unsafe's Pointer type.

// Predeclared alias types. Kept separate for better error messages.

// Predeclared error interface type.

// Types to represent untyped string and boolean constants.

// Types to represent untyped numeric constants.
// Note: Currently these are only used within the binary export
// data format. The rest of the compiler only uses Types[TIDEAL].

// A Type represents a Go type.
type Type struct {
	// Extra contains extra etype-specific fields.
	// As an optimization, those etype-specific structs which contain exactly
	// one pointer-shaped field are stored as values rather than pointers when possible.
	//
	// TMAP: *Map
	// TFORW: *Forward
	// TFUNC: *Func
	// TSTRUCT: *Struct
	// TINTER: *Interface
	// TDDDFIELD: DDDField
	// TFUNCARGS: FuncArgs
	// TCHANARGS: ChanArgs
	// TCHAN: *Chan
	// TPTR32, TPTR64: Ptr
	// TARRAY: *Array
	// TSLICE: Slice
	Extra interface{}

	// Width is the width of this Type in bytes.
	Width int64

	methods    Fields
	allMethods Fields

	Nod  *Node // canonical OTYPE node
	Orig *Type // original type (type literal or predefined type)

	SliceOf *Type
	PtrBase *Type

	Sym    *Sym  // symbol containing name, for named types
	Vargen int32 // unique name for OTYPE/ONAME

	Etype EType // kind of type
	Align uint8 // the required alignment of this type, in bytes

	flags bitset8
}

const (
	typeNotInHeap = 1 << iota // type cannot be heap allocated
	typeBroke                 // broken type definition
	typeNoalg                 // suppress hash and eq algorithm generation
	typeDeferwidth
	typeRecur
)

func (t *Type) NotInHeap() bool  { return t.flags&typeNotInHeap != 0 }
func (t *Type) Broke() bool      { return t.flags&typeBroke != 0 }
func (t *Type) Noalg() bool      { return t.flags&typeNoalg != 0 }
func (t *Type) Deferwidth() bool { return t.flags&typeDeferwidth != 0 }
func (t *Type) Recur() bool      { return t.flags&typeRecur != 0 }

func (t *Type) SetNotInHeap(b bool)  { t.flags.set(typeNotInHeap, b) }
func (t *Type) SetBroke(b bool)      { t.flags.set(typeBroke, b) }
func (t *Type) SetNoalg(b bool)      { t.flags.set(typeNoalg, b) }
func (t *Type) SetDeferwidth(b bool) { t.flags.set(typeDeferwidth, b) }
func (t *Type) SetRecur(b bool)      { t.flags.set(typeRecur, b) }

// Pkg returns the package that t appeared in.
//
// Pkg is only defined for function, struct, and interface types
// (i.e., types with named elements). This information isn't used by
// cmd/compile itself, but we need to track it because it's exposed by
// the go/types API.
func (t *Type) Pkg(psess *PackageSession) *Pkg {
	switch t.Etype {
	case TFUNC:
		return t.Extra.(*Func).pkg
	case TSTRUCT:
		return t.Extra.(*Struct).pkg
	case TINTER:
		return t.Extra.(*Interface).pkg
	default:
		psess.
			Fatalf("Pkg: unexpected kind: %v", t)
		return nil
	}
}

// SetPkg sets the package that t appeared in.
func (t *Type) SetPkg(psess *PackageSession, pkg *Pkg) {
	switch t.Etype {
	case TFUNC:
		t.Extra.(*Func).pkg = pkg
	case TSTRUCT:
		t.Extra.(*Struct).pkg = pkg
	case TINTER:
		t.Extra.(*Interface).pkg = pkg
	default:
		psess.
			Fatalf("Pkg: unexpected kind: %v", t)
	}
}

// Map contains Type fields specific to maps.
type Map struct {
	Key  *Type // Key type
	Elem *Type // Val (elem) type

	Bucket *Type // internal struct type representing a hash bucket
	Hmap   *Type // internal struct type representing the Hmap (map header object)
	Hiter  *Type // internal struct type representing hash iterator state
}

// MapType returns t's extra map-specific fields.
func (t *Type) MapType(psess *PackageSession) *Map {
	t.wantEtype(psess, TMAP)
	return t.Extra.(*Map)
}

// Forward contains Type fields specific to forward types.
type Forward struct {
	Copyto      []*Node  // where to copy the eventual value to
	Embedlineno src.XPos // first use of this type as an embedded type
}

// ForwardType returns t's extra forward-type-specific fields.
func (t *Type) ForwardType(psess *PackageSession) *Forward {
	t.wantEtype(psess, TFORW)
	return t.Extra.(*Forward)
}

// Func contains Type fields specific to func types.
type Func struct {
	Receiver *Type // function receiver
	Results  *Type // function results
	Params   *Type // function params

	Nname *Node
	pkg   *Pkg

	// Argwid is the total width of the function receiver, params, and results.
	// It gets calculated via a temporary TFUNCARGS type.
	// Note that TFUNC's Width is Widthptr.
	Argwid int64

	Outnamed bool
}

// FuncType returns t's extra func-specific fields.
func (t *Type) FuncType(psess *PackageSession) *Func {
	t.wantEtype(psess, TFUNC)
	return t.Extra.(*Func)
}

// StructType contains Type fields specific to struct types.
type Struct struct {
	fields Fields
	pkg    *Pkg

	// Maps have three associated internal structs (see struct MapType).
	// Map links such structs back to their map type.
	Map *Type

	Funarg Funarg // type of function arguments for arg struct
}

// Fnstruct records the kind of function argument
type Funarg uint8

const (
	FunargNone    Funarg = iota
	FunargRcvr           // receiver
	FunargParams         // input parameters
	FunargResults        // output results
)

// StructType returns t's extra struct-specific fields.
func (t *Type) StructType(psess *PackageSession) *Struct {
	t.wantEtype(psess, TSTRUCT)
	return t.Extra.(*Struct)
}

// Interface contains Type fields specific to interface types.
type Interface struct {
	Fields Fields
	pkg    *Pkg
}

// Ptr contains Type fields specific to pointer types.
type Ptr struct {
	Elem *Type // element type
}

// DDDField contains Type fields specific to TDDDFIELD types.
type DDDField struct {
	T *Type // reference to a slice type for ... args
}

// ChanArgs contains Type fields specific to TCHANARGS types.
type ChanArgs struct {
	T *Type // reference to a chan type whose elements need a width check
}

// // FuncArgs contains Type fields specific to TFUNCARGS types.
type FuncArgs struct {
	T *Type // reference to a func type whose elements need a width check
}

// Chan contains Type fields specific to channel types.
type Chan struct {
	Elem *Type   // element type
	Dir  ChanDir // channel direction
}

// ChanType returns t's extra channel-specific fields.
func (t *Type) ChanType(psess *PackageSession) *Chan {
	t.wantEtype(psess, TCHAN)
	return t.Extra.(*Chan)
}

type Tuple struct {
	first  *Type
	second *Type
}

// Array contains Type fields specific to array types.
type Array struct {
	Elem  *Type // element type
	Bound int64 // number of elements; <0 if unknown yet
}

// Slice contains Type fields specific to slice types.
type Slice struct {
	Elem *Type // element type
}

// A Field represents a field in a struct or a method in an interface or
// associated with a named type.
type Field struct {
	flags bitset8

	Embedded uint8 // embedded field

	Pos  src.XPos
	Sym  *Sym
	Type *Type  // field type
	Note string // literal string annotation

	// For fields that represent function parameters, Nname points
	// to the associated ONAME Node.
	Nname *Node

	// Offset in bytes of this field or method within its enclosing struct
	// or interface Type.
	Offset int64
}

const (
	fieldIsddd = 1 << iota // field is ... argument
	fieldBroke             // broken field definition
	fieldNointerface
)

func (f *Field) Isddd() bool       { return f.flags&fieldIsddd != 0 }
func (f *Field) Broke() bool       { return f.flags&fieldBroke != 0 }
func (f *Field) Nointerface() bool { return f.flags&fieldNointerface != 0 }

func (f *Field) SetIsddd(b bool)       { f.flags.set(fieldIsddd, b) }
func (f *Field) SetBroke(b bool)       { f.flags.set(fieldBroke, b) }
func (f *Field) SetNointerface(b bool) { f.flags.set(fieldNointerface, b) }

// End returns the offset of the first byte immediately after this field.
func (f *Field) End() int64 {
	return f.Offset + f.Type.Width
}

// Fields is a pointer to a slice of *Field.
// This saves space in Types that do not have fields or methods
// compared to a simple slice of *Field.
type Fields struct {
	s *[]*Field
}

// Len returns the number of entries in f.
func (f *Fields) Len() int {
	if f.s == nil {
		return 0
	}
	return len(*f.s)
}

// Slice returns the entries in f as a slice.
// Changes to the slice entries will be reflected in f.
func (f *Fields) Slice() []*Field {
	if f.s == nil {
		return nil
	}
	return *f.s
}

// Index returns the i'th element of Fields.
// It panics if f does not have at least i+1 elements.
func (f *Fields) Index(i int) *Field {
	return (*f.s)[i]
}

// Set sets f to a slice.
// This takes ownership of the slice.
func (f *Fields) Set(s []*Field) {
	if len(s) == 0 {
		f.s = nil
	} else {

		t := s
		f.s = &t
	}
}

// Append appends entries to f.
func (f *Fields) Append(s ...*Field) {
	if f.s == nil {
		f.s = new([]*Field)
	}
	*f.s = append(*f.s, s...)
}

// New returns a new Type of the specified kind.
func New(et EType) *Type {
	t := &Type{
		Etype: et,
		Width: BADWIDTH,
	}
	t.Orig = t

	switch t.Etype {
	case TMAP:
		t.Extra = new(Map)
	case TFORW:
		t.Extra = new(Forward)
	case TFUNC:
		t.Extra = new(Func)
	case TSTRUCT:
		t.Extra = new(Struct)
	case TINTER:
		t.Extra = new(Interface)
	case TPTR32, TPTR64:
		t.Extra = Ptr{}
	case TCHANARGS:
		t.Extra = ChanArgs{}
	case TFUNCARGS:
		t.Extra = FuncArgs{}
	case TDDDFIELD:
		t.Extra = DDDField{}
	case TCHAN:
		t.Extra = new(Chan)
	case TTUPLE:
		t.Extra = new(Tuple)
	}
	return t
}

// NewArray returns a new fixed-length array Type.
func (psess *PackageSession) NewArray(elem *Type, bound int64) *Type {
	if bound < 0 {
		psess.
			Fatalf("NewArray: invalid bound %v", bound)
	}
	t := New(TARRAY)
	t.Extra = &Array{Elem: elem, Bound: bound}
	t.SetNotInHeap(elem.NotInHeap())
	return t
}

// NewSlice returns the slice Type with element type elem.
func (psess *PackageSession) NewSlice(elem *Type) *Type {
	if t := elem.SliceOf; t != nil {
		if t.Elem(psess) != elem {
			psess.
				Fatalf("elem mismatch")
		}
		return t
	}

	t := New(TSLICE)
	t.Extra = Slice{Elem: elem}
	elem.SliceOf = t
	return t
}

// NewDDDArray returns a new [...]T array Type.
func NewDDDArray(elem *Type) *Type {
	t := New(TARRAY)
	t.Extra = &Array{Elem: elem, Bound: -1}
	t.SetNotInHeap(elem.NotInHeap())
	return t
}

// NewChan returns a new chan Type with direction dir.
func (psess *PackageSession) NewChan(elem *Type, dir ChanDir) *Type {
	t := New(TCHAN)
	ct := t.ChanType(psess)
	ct.Elem = elem
	ct.Dir = dir
	return t
}

func NewTuple(t1, t2 *Type) *Type {
	t := New(TTUPLE)
	t.Extra.(*Tuple).first = t1
	t.Extra.(*Tuple).second = t2
	return t
}

func newSSA(name string) *Type {
	t := New(TSSA)
	t.Extra = name
	return t
}

// NewMap returns a new map Type with key type k and element (aka value) type v.
func (psess *PackageSession) NewMap(k, v *Type) *Type {
	t := New(TMAP)
	mt := t.MapType(psess)
	mt.Key = k
	mt.Elem = v
	return t
}

// NewPtrCacheEnabled controls whether *T Types are cached in T.
// Caching is disabled just before starting the backend.
// This allows the backend to run concurrently.

// NewPtr returns the pointer type pointing to t.
func (psess *PackageSession) NewPtr(elem *Type) *Type {
	if elem == nil {
		psess.
			Fatalf("NewPtr: pointer to elem Type is nil")
	}

	if t := elem.PtrBase; t != nil {
		if t.Elem(psess) != elem {
			psess.
				Fatalf("NewPtr: elem mismatch")
		}
		return t
	}

	if psess.Tptr == 0 {
		psess.
			Fatalf("NewPtr: Tptr not initialized")
	}

	t := New(psess.Tptr)
	t.Extra = Ptr{Elem: elem}
	t.Width = int64(psess.Widthptr)
	t.Align = uint8(psess.Widthptr)
	if psess.NewPtrCacheEnabled {
		elem.PtrBase = t
	}
	return t
}

// NewDDDField returns a new TDDDFIELD type for slice type s.
func NewDDDField(s *Type) *Type {
	t := New(TDDDFIELD)
	t.Extra = DDDField{T: s}
	return t
}

// NewChanArgs returns a new TCHANARGS type for channel type c.
func NewChanArgs(c *Type) *Type {
	t := New(TCHANARGS)
	t.Extra = ChanArgs{T: c}
	return t
}

// NewFuncArgs returns a new TFUNCARGS type for func type f.
func NewFuncArgs(f *Type) *Type {
	t := New(TFUNCARGS)
	t.Extra = FuncArgs{T: f}
	return t
}

func NewField() *Field {
	return &Field{
		Offset: BADWIDTH,
	}
}

// SubstAny walks t, replacing instances of "any" with successive
// elements removed from types.  It returns the substituted type.
func (psess *PackageSession) SubstAny(t *Type, types *[]*Type) *Type {
	if t == nil {
		return nil
	}

	switch t.Etype {
	default:

	case TANY:
		if len(*types) == 0 {
			psess.
				Fatalf("substArgTypes: not enough argument types")
		}
		t = (*types)[0]
		*types = (*types)[1:]

	case TPTR32, TPTR64:
		elem := psess.SubstAny(t.Elem(psess), types)
		if elem != t.Elem(psess) {
			t = t.copy(psess)
			t.Extra = Ptr{Elem: elem}
		}

	case TARRAY:
		elem := psess.SubstAny(t.Elem(psess), types)
		if elem != t.Elem(psess) {
			t = t.copy(psess)
			t.Extra.(*Array).Elem = elem
		}

	case TSLICE:
		elem := psess.SubstAny(t.Elem(psess), types)
		if elem != t.Elem(psess) {
			t = t.copy(psess)
			t.Extra = Slice{Elem: elem}
		}

	case TCHAN:
		elem := psess.SubstAny(t.Elem(psess), types)
		if elem != t.Elem(psess) {
			t = t.copy(psess)
			t.Extra.(*Chan).Elem = elem
		}

	case TMAP:
		key := psess.SubstAny(t.Key(psess), types)
		elem := psess.SubstAny(t.Elem(psess), types)
		if key != t.Key(psess) || elem != t.Elem(psess) {
			t = t.copy(psess)
			t.Extra.(*Map).Key = key
			t.Extra.(*Map).Elem = elem
		}

	case TFUNC:
		recvs := psess.SubstAny(t.Recvs(psess), types)
		params := psess.SubstAny(t.Params(psess), types)
		results := psess.SubstAny(t.Results(psess), types)
		if recvs != t.Recvs(psess) || params != t.Params(psess) || results != t.Results(psess) {
			t = t.copy(psess)
			t.FuncType(psess).Receiver = recvs
			t.FuncType(psess).Results = results
			t.FuncType(psess).Params = params
		}

	case TSTRUCT:
		fields := t.FieldSlice(psess)
		var nfs []*Field
		for i, f := range fields {
			nft := psess.SubstAny(f.Type, types)
			if nft == f.Type {
				continue
			}
			if nfs == nil {
				nfs = append([]*Field(nil), fields...)
			}
			nfs[i] = f.Copy()
			nfs[i].Type = nft
		}
		if nfs != nil {
			t = t.copy(psess)
			t.SetFields(psess, nfs)
		}
	}

	return t
}

// copy returns a shallow copy of the Type.
func (t *Type) copy(psess *PackageSession) *Type {
	if t == nil {
		return nil
	}
	nt := *t

	switch t.Etype {
	case TMAP:
		x := *t.Extra.(*Map)
		nt.Extra = &x
	case TFORW:
		x := *t.Extra.(*Forward)
		nt.Extra = &x
	case TFUNC:
		x := *t.Extra.(*Func)
		nt.Extra = &x
	case TSTRUCT:
		x := *t.Extra.(*Struct)
		nt.Extra = &x
	case TINTER:
		x := *t.Extra.(*Interface)
		nt.Extra = &x
	case TCHAN:
		x := *t.Extra.(*Chan)
		nt.Extra = &x
	case TARRAY:
		x := *t.Extra.(*Array)
		nt.Extra = &x
	case TTUPLE, TSSA:
		psess.
			Fatalf("ssa types cannot be copied")
	}

	if t.Orig == t {
		nt.Orig = &nt
	}
	return &nt
}

func (f *Field) Copy() *Field {
	nf := *f
	return &nf
}

func (t *Type) wantEtype(psess *PackageSession, et EType) {
	if t.Etype != et {
		psess.
			Fatalf("want %v, but have %v", et, t)
	}
}

func (t *Type) Recvs(psess *PackageSession) *Type   { return t.FuncType(psess).Receiver }
func (t *Type) Params(psess *PackageSession) *Type  { return t.FuncType(psess).Params }
func (t *Type) Results(psess *PackageSession) *Type { return t.FuncType(psess).Results }

func (t *Type) NumRecvs(psess *PackageSession) int {
	return t.FuncType(psess).Receiver.NumFields(psess)
}
func (t *Type) NumParams(psess *PackageSession) int { return t.FuncType(psess).Params.NumFields(psess) }
func (t *Type) NumResults(psess *PackageSession) int {
	return t.FuncType(psess).Results.NumFields(psess)
}

// IsVariadic reports whether function type t is variadic.
func (t *Type) IsVariadic(psess *PackageSession) bool {
	n := t.NumParams(psess)
	return n > 0 && t.Params(psess).Field(psess, n-1).Isddd()
}

// Recv returns the receiver of function type t, if any.
func (t *Type) Recv(psess *PackageSession) *Field {
	s := t.Recvs(psess)
	if s.NumFields(psess) == 0 {
		return nil
	}
	return s.Field(psess, 0)
}

// RecvsParamsResults stores the accessor functions for a function Type's
// receiver, parameters, and result parameters, in that order.
// It can be used to iterate over all of a function's parameter lists.

// RecvsParams is like RecvsParamsResults, but omits result parameters.

// ParamsResults is like RecvsParamsResults, but omits receiver parameters.

// Key returns the key type of map type t.
func (t *Type) Key(psess *PackageSession) *Type {
	t.wantEtype(psess, TMAP)
	return t.Extra.(*Map).Key
}

// Elem returns the type of elements of t.
// Usable with pointers, channels, arrays, slices, and maps.
func (t *Type) Elem(psess *PackageSession) *Type {
	switch t.Etype {
	case TPTR32, TPTR64:
		return t.Extra.(Ptr).Elem
	case TARRAY:
		return t.Extra.(*Array).Elem
	case TSLICE:
		return t.Extra.(Slice).Elem
	case TCHAN:
		return t.Extra.(*Chan).Elem
	case TMAP:
		return t.Extra.(*Map).Elem
	}
	psess.
		Fatalf("Type.Elem %s", t.Etype)
	return nil
}

// DDDField returns the slice ... type for TDDDFIELD type t.
func (t *Type) DDDField(psess *PackageSession) *Type {
	t.wantEtype(psess, TDDDFIELD)
	return t.Extra.(DDDField).T
}

// ChanArgs returns the channel type for TCHANARGS type t.
func (t *Type) ChanArgs(psess *PackageSession) *Type {
	t.wantEtype(psess, TCHANARGS)
	return t.Extra.(ChanArgs).T
}

// FuncArgs returns the channel type for TFUNCARGS type t.
func (t *Type) FuncArgs(psess *PackageSession) *Type {
	t.wantEtype(psess, TFUNCARGS)
	return t.Extra.(FuncArgs).T
}

// Nname returns the associated function's nname.
func (t *Type) Nname(psess *PackageSession) *Node {
	switch t.Etype {
	case TFUNC:
		return t.Extra.(*Func).Nname
	}
	psess.
		Fatalf("Type.Nname %v %v", t.Etype, t)
	return nil
}

// Nname sets the associated function's nname.
func (t *Type) SetNname(psess *PackageSession, n *Node) {
	switch t.Etype {
	case TFUNC:
		t.Extra.(*Func).Nname = n
	default:
		psess.
			Fatalf("Type.SetNname %v %v", t.Etype, t)
	}
}

// IsFuncArgStruct reports whether t is a struct representing function parameters.
func (t *Type) IsFuncArgStruct() bool {
	return t.Etype == TSTRUCT && t.Extra.(*Struct).Funarg != FunargNone
}

func (t *Type) Methods() *Fields {

	return &t.methods
}

func (t *Type) AllMethods() *Fields {

	return &t.allMethods
}

func (t *Type) Fields(psess *PackageSession) *Fields {
	switch t.Etype {
	case TSTRUCT:
		return &t.Extra.(*Struct).fields
	case TINTER:
		psess.
			Dowidth(t)
		return &t.Extra.(*Interface).Fields
	}
	psess.
		Fatalf("Fields: type %v does not have fields", t)
	return nil
}

// Field returns the i'th field/method of struct/interface type t.
func (t *Type) Field(psess *PackageSession, i int) *Field {
	return t.Fields(psess).Slice()[i]
}

// FieldSlice returns a slice of containing all fields/methods of
// struct/interface type t.
func (t *Type) FieldSlice(psess *PackageSession) []*Field {
	return t.Fields(psess).Slice()
}

// SetFields sets struct/interface type t's fields/methods to fields.
func (t *Type) SetFields(psess *PackageSession, fields []*Field) {

	if t.WidthCalculated() {
		psess.
			Fatalf("SetFields of %v: width previously calculated", t)
	}
	t.wantEtype(psess, TSTRUCT)
	for _, f := range fields {

		if f.Type != nil && f.Type.NotInHeap() {
			t.SetNotInHeap(true)
			break
		}
	}
	t.Fields(psess).Set(fields)
}

func (t *Type) SetInterface(psess *PackageSession, methods []*Field) {
	t.wantEtype(psess, TINTER)
	t.Methods().Set(methods)
}

func (t *Type) IsDDDArray() bool {
	if t.Etype != TARRAY {
		return false
	}
	return t.Extra.(*Array).Bound < 0
}

func (t *Type) WidthCalculated() bool {
	return t.Align > 0
}

// ArgWidth returns the total aligned argument size for a function.
// It includes the receiver, parameters, and results.
func (t *Type) ArgWidth(psess *PackageSession) int64 {
	t.wantEtype(psess, TFUNC)
	return t.Extra.(*Func).Argwid
}

func (t *Type) Size(psess *PackageSession) int64 {
	if t.Etype == TSSA {
		if t == psess.TypeInt128 {
			return 16
		}
		return 0
	}
	psess.
		Dowidth(t)
	return t.Width
}

func (t *Type) Alignment(psess *PackageSession) int64 {
	psess.
		Dowidth(t)
	return int64(t.Align)
}

func (t *Type) SimpleString(psess *PackageSession) string {
	return t.Etype.String(psess)
}

// Cmp is a comparison between values a and b.
// -1 if a < b
//  0 if a == b
//  1 if a > b
type Cmp int8

const (
	CMPlt = Cmp(-1)
	CMPeq = Cmp(0)
	CMPgt = Cmp(1)
)

// Compare compares types for purposes of the SSA back
// end, returning a Cmp (one of CMPlt, CMPeq, CMPgt).
// The answers are correct for an optimizer
// or code generator, but not necessarily typechecking.
// The order chosen is arbitrary, only consistency and division
// into equivalence classes (Types that compare CMPeq) matters.
func (t *Type) Compare(psess *PackageSession, x *Type) Cmp {
	if x == t {
		return CMPeq
	}
	return t.cmp(psess, x)
}

func cmpForNe(x bool) Cmp {
	if x {
		return CMPlt
	}
	return CMPgt
}

func (r *Sym) cmpsym(s *Sym) Cmp {
	if r == s {
		return CMPeq
	}
	if r == nil {
		return CMPlt
	}
	if s == nil {
		return CMPgt
	}

	if len(r.Name) != len(s.Name) {
		return cmpForNe(len(r.Name) < len(s.Name))
	}
	if r.Pkg != s.Pkg {
		if len(r.Pkg.Prefix) != len(s.Pkg.Prefix) {
			return cmpForNe(len(r.Pkg.Prefix) < len(s.Pkg.Prefix))
		}
		if r.Pkg.Prefix != s.Pkg.Prefix {
			return cmpForNe(r.Pkg.Prefix < s.Pkg.Prefix)
		}
	}
	if r.Name != s.Name {
		return cmpForNe(r.Name < s.Name)
	}
	return CMPeq
}

// cmp compares two *Types t and x, returning CMPlt,
// CMPeq, CMPgt as t<x, t==x, t>x, for an arbitrary
// and optimizer-centric notion of comparison.
// TODO(josharian): make this safe for recursive interface types
// and use in signatlist sorting. See issue 19869.
func (t *Type) cmp(psess *PackageSession, x *Type) Cmp {

	if t == x {
		return CMPeq
	}
	if t == nil {
		return CMPlt
	}
	if x == nil {
		return CMPgt
	}

	if t.Etype != x.Etype {
		return cmpForNe(t.Etype < x.Etype)
	}

	if t.Sym != nil || x.Sym != nil {

		switch t.Etype {
		case TUINT8:
			if (t == psess.Types[TUINT8] || t == psess.Bytetype) && (x == psess.Types[TUINT8] || x == psess.Bytetype) {
				return CMPeq
			}

		case TINT32:
			if (t == psess.Types[psess.Runetype.Etype] || t == psess.Runetype) && (x == psess.Types[psess.Runetype.Etype] || x == psess.Runetype) {
				return CMPeq
			}
		}
	}

	if c := t.Sym.cmpsym(x.Sym); c != CMPeq {
		return c
	}

	if x.Sym != nil {

		if t.Vargen != x.Vargen {
			return cmpForNe(t.Vargen < x.Vargen)
		}
		return CMPeq
	}

	switch t.Etype {
	case TBOOL, TFLOAT32, TFLOAT64, TCOMPLEX64, TCOMPLEX128, TUNSAFEPTR, TUINTPTR,
		TINT8, TINT16, TINT32, TINT64, TINT, TUINT8, TUINT16, TUINT32, TUINT64, TUINT:
		return CMPeq

	case TSSA:
		tname := t.Extra.(string)
		xname := t.Extra.(string)

		if len(tname) == len(xname) {
			if tname == xname {
				return CMPeq
			}
			if tname < xname {
				return CMPlt
			}
			return CMPgt
		}
		if len(tname) > len(xname) {
			return CMPgt
		}
		return CMPlt

	case TTUPLE:
		xtup := x.Extra.(*Tuple)
		ttup := t.Extra.(*Tuple)
		if c := ttup.first.Compare(psess, xtup.first); c != CMPeq {
			return c
		}
		return ttup.second.Compare(psess, xtup.second)

	case TMAP:
		if c := t.Key(psess).cmp(psess, x.Key(psess)); c != CMPeq {
			return c
		}
		return t.Elem(psess).cmp(psess, x.Elem(psess))

	case TPTR32, TPTR64, TSLICE:

	case TSTRUCT:
		if t.StructType(psess).Map == nil {
			if x.StructType(psess).Map != nil {
				return CMPlt
			}

		} else if x.StructType(psess).Map == nil {
			return CMPgt
		} else if t.StructType(psess).Map.MapType(psess).Bucket == t {

			if x.StructType(psess).Map.MapType(psess).Bucket != x {
				return CMPlt
			}
			return t.StructType(psess).Map.cmp(psess, x.StructType(psess).Map)
		} else if x.StructType(psess).Map.MapType(psess).Bucket == x {
			return CMPgt
		}

		tfs := t.FieldSlice(psess)
		xfs := x.FieldSlice(psess)
		for i := 0; i < len(tfs) && i < len(xfs); i++ {
			t1, x1 := tfs[i], xfs[i]
			if t1.Embedded != x1.Embedded {
				return cmpForNe(t1.Embedded < x1.Embedded)
			}
			if t1.Note != x1.Note {
				return cmpForNe(t1.Note < x1.Note)
			}
			if c := t1.Sym.cmpsym(x1.Sym); c != CMPeq {
				return c
			}
			if c := t1.Type.cmp(psess, x1.Type); c != CMPeq {
				return c
			}
		}
		if len(tfs) != len(xfs) {
			return cmpForNe(len(tfs) < len(xfs))
		}
		return CMPeq

	case TINTER:
		tfs := t.FieldSlice(psess)
		xfs := x.FieldSlice(psess)
		for i := 0; i < len(tfs) && i < len(xfs); i++ {
			t1, x1 := tfs[i], xfs[i]
			if c := t1.Sym.cmpsym(x1.Sym); c != CMPeq {
				return c
			}
			if c := t1.Type.cmp(psess, x1.Type); c != CMPeq {
				return c
			}
		}
		if len(tfs) != len(xfs) {
			return cmpForNe(len(tfs) < len(xfs))
		}
		return CMPeq

	case TFUNC:
		for _, f := range psess.RecvsParamsResults {

			tfs := f(t).FieldSlice(psess)
			xfs := f(x).FieldSlice(psess)
			for i := 0; i < len(tfs) && i < len(xfs); i++ {
				ta := tfs[i]
				tb := xfs[i]
				if ta.Isddd() != tb.Isddd() {
					return cmpForNe(!ta.Isddd())
				}
				if c := ta.Type.cmp(psess, tb.Type); c != CMPeq {
					return c
				}
			}
			if len(tfs) != len(xfs) {
				return cmpForNe(len(tfs) < len(xfs))
			}
		}
		return CMPeq

	case TARRAY:
		if t.NumElem(psess) != x.NumElem(psess) {
			return cmpForNe(t.NumElem(psess) < x.NumElem(psess))
		}

	case TCHAN:
		if t.ChanDir(psess) != x.ChanDir(psess) {
			return cmpForNe(t.ChanDir(psess) < x.ChanDir(psess))
		}

	default:
		e := fmt.Sprintf("Do not know how to compare %v with %v", t, x)
		panic(e)
	}

	return t.Elem(psess).cmp(psess, x.Elem(psess))
}

// IsKind reports whether t is a Type of the specified kind.
func (t *Type) IsKind(et EType) bool {
	return t != nil && t.Etype == et
}

func (t *Type) IsBoolean() bool {
	return t.Etype == TBOOL
}

// ToUnsigned returns the unsigned equivalent of integer type t.
func (t *Type) ToUnsigned(psess *PackageSession) *Type {
	if !t.IsInteger() {
		psess.
			Fatalf("unsignedType(%v)", t)
	}
	return psess.Types[psess.unsignedEType[t.Etype]]
}

func (t *Type) IsInteger() bool {
	switch t.Etype {
	case TINT8, TUINT8, TINT16, TUINT16, TINT32, TUINT32, TINT64, TUINT64, TINT, TUINT, TUINTPTR:
		return true
	}
	return false
}

func (t *Type) IsSigned() bool {
	switch t.Etype {
	case TINT8, TINT16, TINT32, TINT64, TINT:
		return true
	}
	return false
}

func (t *Type) IsFloat() bool {
	return t.Etype == TFLOAT32 || t.Etype == TFLOAT64
}

func (t *Type) IsComplex() bool {
	return t.Etype == TCOMPLEX64 || t.Etype == TCOMPLEX128
}

// IsPtr reports whether t is a regular Go pointer type.
// This does not include unsafe.Pointer.
func (t *Type) IsPtr() bool {
	return t.Etype == TPTR32 || t.Etype == TPTR64
}

// IsUnsafePtr reports whether t is an unsafe pointer.
func (t *Type) IsUnsafePtr() bool {
	return t.Etype == TUNSAFEPTR
}

// IsPtrShaped reports whether t is represented by a single machine pointer.
// In addition to regular Go pointer types, this includes map, channel, and
// function types and unsafe.Pointer. It does not include array or struct types
// that consist of a single pointer shaped type.
// TODO(mdempsky): Should it? See golang.org/issue/15028.
func (t *Type) IsPtrShaped() bool {
	return t.Etype == TPTR32 || t.Etype == TPTR64 || t.Etype == TUNSAFEPTR ||
		t.Etype == TMAP || t.Etype == TCHAN || t.Etype == TFUNC
}

func (t *Type) IsString() bool {
	return t.Etype == TSTRING
}

func (t *Type) IsMap() bool {
	return t.Etype == TMAP
}

func (t *Type) IsChan() bool {
	return t.Etype == TCHAN
}

func (t *Type) IsSlice() bool {
	return t.Etype == TSLICE
}

func (t *Type) IsArray() bool {
	return t.Etype == TARRAY
}

func (t *Type) IsStruct() bool {
	return t.Etype == TSTRUCT
}

func (t *Type) IsInterface() bool {
	return t.Etype == TINTER
}

// IsEmptyInterface reports whether t is an empty interface type.
func (t *Type) IsEmptyInterface(psess *PackageSession) bool {
	return t.IsInterface() && t.NumFields(psess) == 0
}

func (t *Type) PtrTo(psess *PackageSession) *Type {
	return psess.NewPtr(t)
}

func (t *Type) NumFields(psess *PackageSession) int {
	return t.Fields(psess).Len()
}
func (t *Type) FieldType(psess *PackageSession, i int) *Type {
	if t.Etype == TTUPLE {
		switch i {
		case 0:
			return t.Extra.(*Tuple).first
		case 1:
			return t.Extra.(*Tuple).second
		default:
			panic("bad tuple index")
		}
	}
	return t.Field(psess, i).Type
}
func (t *Type) FieldOff(psess *PackageSession, i int) int64 {
	return t.Field(psess, i).Offset
}
func (t *Type) FieldName(psess *PackageSession, i int) string {
	return t.Field(psess, i).Sym.Name
}

func (t *Type) NumElem(psess *PackageSession) int64 {
	t.wantEtype(psess, TARRAY)
	at := t.Extra.(*Array)
	if at.Bound < 0 {
		psess.
			Fatalf("NumElem array %v does not have bound yet", t)
	}
	return at.Bound
}

// SetNumElem sets the number of elements in an array type.
// The only allowed use is on array types created with NewDDDArray.
// For other uses, create a new array with NewArray instead.
func (t *Type) SetNumElem(psess *PackageSession, n int64) {
	t.wantEtype(psess, TARRAY)
	at := t.Extra.(*Array)
	if at.Bound >= 0 {
		psess.
			Fatalf("SetNumElem array %v already has bound %d", t, at.Bound)
	}
	at.Bound = n
}

type componentsIncludeBlankFields bool

const (
	IgnoreBlankFields componentsIncludeBlankFields = false
	CountBlankFields  componentsIncludeBlankFields = true
)

// NumComponents returns the number of primitive elements that compose t.
// Struct and array types are flattened for the purpose of counting.
// All other types (including string, slice, and interface types) count as one element.
// If countBlank is IgnoreBlankFields, then blank struct fields
// (and their comprised elements) are excluded from the count.
// struct { x, y [3]int } has six components; [10]struct{ x, y string } has twenty.
func (t *Type) NumComponents(psess *PackageSession, countBlank componentsIncludeBlankFields) int64 {
	switch t.Etype {
	case TSTRUCT:
		if t.IsFuncArgStruct() {
			psess.
				Fatalf("NumComponents func arg struct")
		}
		var n int64
		for _, f := range t.FieldSlice(psess) {
			if countBlank == IgnoreBlankFields && f.Sym.IsBlank() {
				continue
			}
			n += f.Type.NumComponents(psess, countBlank)
		}
		return n
	case TARRAY:
		return t.NumElem(psess) * t.Elem(psess).NumComponents(psess, countBlank)
	}
	return 1
}

// ChanDir returns the direction of a channel type t.
// The direction will be one of Crecv, Csend, or Cboth.
func (t *Type) ChanDir(psess *PackageSession) ChanDir {
	t.wantEtype(psess, TCHAN)
	return t.Extra.(*Chan).Dir
}

func (t *Type) IsMemory(psess *PackageSession) bool {
	return t == psess.TypeMem || t.Etype == TTUPLE && t.Extra.(*Tuple).second == psess.TypeMem
}
func (t *Type) IsFlags(psess *PackageSession) bool { return t == psess.TypeFlags }
func (t *Type) IsVoid(psess *PackageSession) bool  { return t == psess.TypeVoid }
func (t *Type) IsTuple() bool                      { return t.Etype == TTUPLE }

// IsUntyped reports whether t is an untyped type.
func (t *Type) IsUntyped(psess *PackageSession) bool {
	if t == nil {
		return false
	}
	if t == psess.Idealstring || t == psess.Idealbool {
		return true
	}
	switch t.Etype {
	case TNIL, TIDEAL:
		return true
	}
	return false
}

func (psess *PackageSession) Haspointers(t *Type) bool {
	return psess.Haspointers1(t, false)
}

func (psess *PackageSession) Haspointers1(t *Type, ignoreNotInHeap bool) bool {
	switch t.Etype {
	case TINT, TUINT, TINT8, TUINT8, TINT16, TUINT16, TINT32, TUINT32, TINT64,
		TUINT64, TUINTPTR, TFLOAT32, TFLOAT64, TCOMPLEX64, TCOMPLEX128, TBOOL, TSSA:
		return false

	case TARRAY:
		if t.NumElem(psess) == 0 {
			return false
		}
		return psess.Haspointers1(t.Elem(psess), ignoreNotInHeap)

	case TSTRUCT:
		for _, t1 := range t.Fields(psess).Slice() {
			if psess.Haspointers1(t1.Type, ignoreNotInHeap) {
				return true
			}
		}
		return false

	case TPTR32, TPTR64, TSLICE:
		return !(ignoreNotInHeap && t.Elem(psess).NotInHeap())

	case TTUPLE:
		ttup := t.Extra.(*Tuple)
		return psess.Haspointers1(ttup.first, ignoreNotInHeap) || psess.Haspointers1(ttup.second, ignoreNotInHeap)
	}

	return true
}

// HasHeapPointer returns whether t contains a heap pointer.
// This is used for write barrier insertion, so it ignores
// pointers to go:notinheap types.
func (t *Type) HasHeapPointer(psess *PackageSession) bool {
	return psess.Haspointers1(t, true)
}

func (t *Type) Symbol(psess *PackageSession) *obj.LSym {
	return psess.TypeLinkSym(t)
}

// Tie returns 'T' if t is a concrete type,
// 'I' if t is an interface type, and 'E' if t is an empty interface type.
// It is used to build calls to the conv* and assert* runtime routines.
func (t *Type) Tie(psess *PackageSession) byte {
	if t.IsEmptyInterface(psess) {
		return 'E'
	}
	if t.IsInterface() {
		return 'I'
	}
	return 'T'
}

// FakeRecvType returns the singleton type used for interface method receivers.
func (psess *PackageSession) FakeRecvType() *Type {
	if psess.recvType == nil {
		psess.
			recvType = psess.NewPtr(New(TSTRUCT))
	}
	return psess.recvType
}

// TSSA types. Haspointers assumes these are pointer-free.

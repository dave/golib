// Package dwarf generates DWARF debugging information.
// DWARF generation is split between the compiler and the linker,
// this package contains the shared code.
package dwarf

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// InfoPrefix is the prefix for all the symbols containing DWARF info entries.
const InfoPrefix = "go.info."

// RangePrefix is the prefix for all the symbols containing DWARF location lists.
const LocPrefix = "go.loc."

// RangePrefix is the prefix for all the symbols containing DWARF range lists.
const RangePrefix = "go.range."

// IsStmtPrefix is the prefix for all the symbols containing DWARF is_stmt info for the line number table.
const IsStmtPrefix = "go.isstmt."

// ConstInfoPrefix is the prefix for all symbols containing DWARF info
// entries that contain constants.
const ConstInfoPrefix = "go.constinfo."

// CUInfoPrefix is the prefix for symbols containing information to
// populate the DWARF compilation unit info entries.
const CUInfoPrefix = "go.cuinfo."

// Used to form the symbol name assigned to the DWARF 'abstract subprogram"
// info entry for a function
const AbstractFuncSuffix = "$abstract"

// Controls logging/debugging for selected aspects of DWARF subprogram
// generation (functions, scopes).

// Sym represents a symbol.
type Sym interface {
	Len() int64
}

// A Var represents a local variable or a function parameter.
type Var struct {
	Name          string
	Abbrev        int // Either DW_ABRV_AUTO[_LOCLIST] or DW_ABRV_PARAM[_LOCLIST]
	IsReturnValue bool
	IsInlFormal   bool
	StackOffset   int32
	// This package can't use the ssa package, so it can't mention ssa.FuncDebug,
	// so indirect through a closure.
	PutLocationList func(listSym, startPC Sym)
	Scope           int32
	Type            Sym
	DeclFile        string
	DeclLine        uint
	DeclCol         uint
	InlIndex        int32 // subtract 1 to form real index into InlTree
	ChildIndex      int32 // child DIE index in abstract function
	IsInAbstract    bool  // variable exists in abstract function
}

// A Scope represents a lexical scope. All variables declared within a
// scope will only be visible to instructions covered by the scope.
// Lexical scopes are contiguous in source files but can end up being
// compiled to discontiguous blocks of instructions in the executable.
// The Ranges field lists all the blocks of instructions that belong
// in this scope.
type Scope struct {
	Parent int32
	Ranges []Range
	Vars   []*Var
}

// A Range represents a half-open interval [Start, End).
type Range struct {
	Start, End int64
}

// This container is used by the PutFunc* variants below when
// creating the DWARF subprogram DIE(s) for a function.
type FnState struct {
	Name       string
	Importpath string
	Info       Sym
	Filesym    Sym
	Loc        Sym
	Ranges     Sym
	Absfn      Sym
	StartPC    Sym
	Size       int64
	External   bool
	Scopes     []Scope
	InlCalls   InlCalls
}

func (psess *PackageSession) EnableLogging(doit bool) {
	psess.
		logDwarf = doit
}

// UnifyRanges merges the list of ranges of c into the list of ranges of s
func (s *Scope) UnifyRanges(c *Scope) {
	out := make([]Range, 0, len(s.Ranges)+len(c.Ranges))

	i, j := 0, 0
	for {
		var cur Range
		if i < len(s.Ranges) && j < len(c.Ranges) {
			if s.Ranges[i].Start < c.Ranges[j].Start {
				cur = s.Ranges[i]
				i++
			} else {
				cur = c.Ranges[j]
				j++
			}
		} else if i < len(s.Ranges) {
			cur = s.Ranges[i]
			i++
		} else if j < len(c.Ranges) {
			cur = c.Ranges[j]
			j++
		} else {
			break
		}

		if n := len(out); n > 0 && cur.Start <= out[n-1].End {
			out[n-1].End = cur.End
		} else {
			out = append(out, cur)
		}
	}

	s.Ranges = out
}

type InlCalls struct {
	Calls []InlCall
}

type InlCall struct {
	// index into ctx.InlTree describing the call inlined here
	InlIndex int

	// Symbol of file containing inlined call site (really *obj.LSym).
	CallFile Sym

	// Line number of inlined call site.
	CallLine uint32

	// Dwarf abstract subroutine symbol (really *obj.LSym).
	AbsFunSym Sym

	// Indices of child inlines within Calls array above.
	Children []int

	// entries in this list are PAUTO's created by the inliner to
	// capture the promoted formals and locals of the inlined callee.
	InlVars []*Var

	// PC ranges for this inlined call.
	Ranges []Range

	// Root call (not a child of some other call).
	Root bool
}

// A Context specifies how to add data to a Sym.
type Context interface {
	PtrSize() int
	AddInt(s Sym, size int, i int64)
	AddBytes(s Sym, b []byte)
	AddAddress(s Sym, t interface{}, ofs int64)
	AddSectionOffset(s Sym, size int, t interface{}, ofs int64)
	AddDWARFSectionOffset(s Sym, size int, t interface{}, ofs int64)
	CurrentOffset(s Sym) int64
	RecordDclReference(from Sym, to Sym, dclIdx int, inlIndex int)
	RecordChildDieOffsets(s Sym, vars []*Var, offsets []int32)
	AddString(s Sym, v string)
	AddFileRef(s Sym, f interface{})
	Logf(format string, args ...interface{})
}

// AppendUleb128 appends v to b using DWARF's unsigned LEB128 encoding.
func AppendUleb128(b []byte, v uint64) []byte {
	for {
		c := uint8(v & 0x7f)
		v >>= 7
		if v != 0 {
			c |= 0x80
		}
		b = append(b, c)
		if c&0x80 == 0 {
			break
		}
	}
	return b
}

// AppendSleb128 appends v to b using DWARF's signed LEB128 encoding.
func AppendSleb128(b []byte, v int64) []byte {
	for {
		c := uint8(v & 0x7f)
		s := uint8(v & 0x40)
		v >>= 7
		if (v != -1 || s == 0) && (v != 0 || s != 0) {
			c |= 0x80
		}
		b = append(b, c)
		if c&0x80 == 0 {
			break
		}
	}
	return b
}

// sevenbits contains all unsigned seven bit numbers, indexed by their value.

// sevenBitU returns the unsigned LEB128 encoding of v if v is seven bits and nil otherwise.
// The contents of the returned slice must not be modified.
func (psess *PackageSession) sevenBitU(v int64) []byte {
	if uint64(v) < uint64(len(psess.sevenbits)) {
		return psess.sevenbits[v : v+1]
	}
	return nil
}

// sevenBitS returns the signed LEB128 encoding of v if v is seven bits and nil otherwise.
// The contents of the returned slice must not be modified.
func (psess *PackageSession) sevenBitS(v int64) []byte {
	if uint64(v) <= 63 {
		return psess.sevenbits[v : v+1]
	}
	if uint64(-v) <= 64 {
		return psess.sevenbits[128+v : 128+v+1]
	}
	return nil
}

// Uleb128put appends v to s using DWARF's unsigned LEB128 encoding.
func (psess *PackageSession) Uleb128put(ctxt Context, s Sym, v int64) {
	b := psess.sevenBitU(v)
	if b == nil {
		var encbuf [20]byte
		b = AppendUleb128(encbuf[:0], uint64(v))
	}
	ctxt.AddBytes(s, b)
}

// Sleb128put appends v to s using DWARF's signed LEB128 encoding.
func (psess *PackageSession) Sleb128put(ctxt Context, s Sym, v int64) {
	b := psess.sevenBitS(v)
	if b == nil {
		var encbuf [20]byte
		b = AppendSleb128(encbuf[:0], v)
	}
	ctxt.AddBytes(s, b)
}

/*
 * Defining Abbrevs.  This is hardcoded, and there will be
 * only a handful of them.  The DWARF spec places no restriction on
 * the ordering of attributes in the Abbrevs and DIEs, and we will
 * always write them out in the order of declaration in the abbrev.
 */
type dwAttrForm struct {
	attr uint16
	form uint8
}

// Go-specific type attributes.
const (
	DW_AT_go_kind = 0x2900
	DW_AT_go_key  = 0x2901
	DW_AT_go_elem = 0x2902
	// Attribute for DW_TAG_member of a struct type.
	// Nonzero value indicates the struct field is an embedded field.
	DW_AT_go_embedded_field = 0x2903
	DW_AT_go_runtime_type   = 0x2904

	DW_AT_internal_location = 253 // params and locals; not emitted
)

// Index into the abbrevs table below.
// Keep in sync with ispubname() and ispubtype() in ld/dwarf.go.
// ispubtype considers >= NULLTYPE public
const (
	DW_ABRV_NULL = iota
	DW_ABRV_COMPUNIT
	DW_ABRV_FUNCTION
	DW_ABRV_FUNCTION_ABSTRACT
	DW_ABRV_FUNCTION_CONCRETE
	DW_ABRV_INLINED_SUBROUTINE
	DW_ABRV_INLINED_SUBROUTINE_RANGES
	DW_ABRV_VARIABLE
	DW_ABRV_INT_CONSTANT
	DW_ABRV_AUTO
	DW_ABRV_AUTO_LOCLIST
	DW_ABRV_AUTO_ABSTRACT
	DW_ABRV_AUTO_CONCRETE
	DW_ABRV_AUTO_CONCRETE_LOCLIST
	DW_ABRV_PARAM
	DW_ABRV_PARAM_LOCLIST
	DW_ABRV_PARAM_ABSTRACT
	DW_ABRV_PARAM_CONCRETE
	DW_ABRV_PARAM_CONCRETE_LOCLIST
	DW_ABRV_LEXICAL_BLOCK_RANGES
	DW_ABRV_LEXICAL_BLOCK_SIMPLE
	DW_ABRV_STRUCTFIELD
	DW_ABRV_FUNCTYPEPARAM
	DW_ABRV_DOTDOTDOT
	DW_ABRV_ARRAYRANGE
	DW_ABRV_NULLTYPE
	DW_ABRV_BASETYPE
	DW_ABRV_ARRAYTYPE
	DW_ABRV_CHANTYPE
	DW_ABRV_FUNCTYPE
	DW_ABRV_IFACETYPE
	DW_ABRV_MAPTYPE
	DW_ABRV_PTRTYPE
	DW_ABRV_BARE_PTRTYPE // only for void*, no DW_AT_type attr to please gdb 6.
	DW_ABRV_SLICETYPE
	DW_ABRV_STRINGTYPE
	DW_ABRV_STRUCTTYPE
	DW_ABRV_TYPEDECL
	DW_NABRV
)

type dwAbbrev struct {
	tag      uint8
	children uint8
	attr     []dwAttrForm
}

// GetAbbrev returns the contents of the .debug_abbrev section.
func (psess *PackageSession) GetAbbrev() []byte {
	var buf []byte
	for i := 1; i < DW_NABRV; i++ {

		buf = AppendUleb128(buf, uint64(i))
		buf = AppendUleb128(buf, uint64(psess.abbrevs[i].tag))
		buf = append(buf, psess.abbrevs[i].children)
		for _, f := range psess.abbrevs[i].attr {
			buf = AppendUleb128(buf, uint64(f.attr))
			buf = AppendUleb128(buf, uint64(f.form))
		}
		buf = append(buf, 0, 0)
	}
	return append(buf, 0)
}

// DWAttr represents an attribute of a DWDie.
//
// For DW_CLS_string and _block, value should contain the length, and
// data the data, for _reference, value is 0 and data is a DWDie* to
// the referenced instance, for all others, value is the whole thing
// and data is null.
type DWAttr struct {
	Link  *DWAttr
	Atr   uint16 // DW_AT_
	Cls   uint8  // DW_CLS_
	Value int64
	Data  interface{}
}

// DWDie represents a DWARF debug info entry.
type DWDie struct {
	Abbrev int
	Link   *DWDie
	Child  *DWDie
	Attr   *DWAttr
	Sym    Sym
}

func (psess *PackageSession) putattr(ctxt Context, s Sym, abbrev int, form int, cls int, value int64, data interface{}) error {
	switch form {
	case DW_FORM_addr:

		if data == nil && value == 0 {
			ctxt.AddInt(s, ctxt.PtrSize(), 0)
			break
		}
		if cls == DW_CLS_GO_TYPEREF {
			ctxt.AddSectionOffset(s, ctxt.PtrSize(), data, value)
			break
		}
		ctxt.AddAddress(s, data, value)

	case DW_FORM_block1:
		if cls == DW_CLS_ADDRESS {
			ctxt.AddInt(s, 1, int64(1+ctxt.PtrSize()))
			ctxt.AddInt(s, 1, DW_OP_addr)
			ctxt.AddAddress(s, data, 0)
			break
		}

		value &= 0xff
		ctxt.AddInt(s, 1, value)
		p := data.([]byte)[:value]
		ctxt.AddBytes(s, p)

	case DW_FORM_block2:
		value &= 0xffff

		ctxt.AddInt(s, 2, value)
		p := data.([]byte)[:value]
		ctxt.AddBytes(s, p)

	case DW_FORM_block4:
		value &= 0xffffffff

		ctxt.AddInt(s, 4, value)
		p := data.([]byte)[:value]
		ctxt.AddBytes(s, p)

	case DW_FORM_block:
		psess.
			Uleb128put(ctxt, s, value)

		p := data.([]byte)[:value]
		ctxt.AddBytes(s, p)

	case DW_FORM_data1:
		ctxt.AddInt(s, 1, value)

	case DW_FORM_data2:
		ctxt.AddInt(s, 2, value)

	case DW_FORM_data4:
		if cls == DW_CLS_PTR {
			ctxt.AddDWARFSectionOffset(s, 4, data, value)
			break
		}
		ctxt.AddInt(s, 4, value)

	case DW_FORM_data8:
		ctxt.AddInt(s, 8, value)

	case DW_FORM_sdata:
		psess.
			Sleb128put(ctxt, s, value)

	case DW_FORM_udata:
		psess.
			Uleb128put(ctxt, s, value)

	case DW_FORM_string:
		str := data.(string)
		ctxt.AddString(s, str)

		for i := int64(len(str)); i < value; i++ {
			ctxt.AddInt(s, 1, 0)
		}

	case DW_FORM_flag:
		if value != 0 {
			ctxt.AddInt(s, 1, 1)
		} else {
			ctxt.AddInt(s, 1, 0)
		}

	case DW_FORM_ref_addr:
		fallthrough
	case DW_FORM_sec_offset:
		if data == nil {
			return fmt.Errorf("dwarf: null reference in %d", abbrev)
		}
		ctxt.AddDWARFSectionOffset(s, 4, data, value)

	case DW_FORM_ref1,
		DW_FORM_ref2,
		DW_FORM_ref4,
		DW_FORM_ref8,
		DW_FORM_ref_udata,

		DW_FORM_strp,
		DW_FORM_indirect:
		fallthrough
	default:
		return fmt.Errorf("dwarf: unsupported attribute form %d / class %d", form, cls)
	}
	return nil
}

// PutAttrs writes the attributes for a DIE to symbol 's'.
//
// Note that we can (and do) add arbitrary attributes to a DIE, but
// only the ones actually listed in the Abbrev will be written out.
func (psess *PackageSession) PutAttrs(ctxt Context, s Sym, abbrev int, attr *DWAttr) {
Outer:
	for _, f := range psess.abbrevs[abbrev].attr {
		for ap := attr; ap != nil; ap = ap.Link {
			if ap.Atr == f.attr {
				psess.
					putattr(ctxt, s, abbrev, int(f.form), int(ap.Cls), ap.Value, ap.Data)
				continue Outer
			}
		}
		psess.
			putattr(ctxt, s, abbrev, int(f.form), 0, 0, nil)
	}
}

// HasChildren returns true if 'die' uses an abbrev that supports children.
func (psess *PackageSession) HasChildren(die *DWDie) bool {
	return psess.abbrevs[die.Abbrev].children != 0
}

// PutIntConst writes a DIE for an integer constant
func (psess *PackageSession) PutIntConst(ctxt Context, info, typ Sym, name string, val int64) {
	psess.
		Uleb128put(ctxt, info, DW_ABRV_INT_CONSTANT)
	psess.
		putattr(ctxt, info, DW_ABRV_INT_CONSTANT, DW_FORM_string, DW_CLS_STRING, int64(len(name)), name)
	psess.
		putattr(ctxt, info, DW_ABRV_INT_CONSTANT, DW_FORM_ref_addr, DW_CLS_REFERENCE, 0, typ)
	psess.
		putattr(ctxt, info, DW_ABRV_INT_CONSTANT, DW_FORM_sdata, DW_CLS_CONSTANT, val, nil)
}

// PutRanges writes a range table to sym. All addresses in ranges are
// relative to some base address. If base is not nil, then they're
// relative to the start of base. If base is nil, then the caller must
// arrange a base address some other way (such as a DW_AT_low_pc
// attribute).
func PutRanges(ctxt Context, sym Sym, base Sym, ranges []Range) {
	ps := ctxt.PtrSize()

	if base != nil {
		ctxt.AddInt(sym, ps, -1)
		ctxt.AddAddress(sym, base, 0)
	}

	for _, r := range ranges {
		ctxt.AddInt(sym, ps, r.Start)
		ctxt.AddInt(sym, ps, r.End)
	}

	ctxt.AddInt(sym, ps, 0)
	ctxt.AddInt(sym, ps, 0)
}

// Return TRUE if the inlined call in the specified slot is empty,
// meaning it has a zero-length range (no instructions), and all
// of its children are empty.
func isEmptyInlinedCall(slot int, calls *InlCalls) bool {
	ic := &calls.Calls[slot]
	if ic.InlIndex == -2 {
		return true
	}
	live := false
	for _, k := range ic.Children {
		if !isEmptyInlinedCall(k, calls) {
			live = true
		}
	}
	if len(ic.Ranges) > 0 {
		live = true
	}
	if !live {
		ic.InlIndex = -2
	}
	return !live
}

// Slot -1:    return top-level inlines
// Slot >= 0:  return children of that slot
func inlChildren(slot int, calls *InlCalls) []int {
	var kids []int
	if slot != -1 {
		for _, k := range calls.Calls[slot].Children {
			if !isEmptyInlinedCall(k, calls) {
				kids = append(kids, k)
			}
		}
	} else {
		for k := 0; k < len(calls.Calls); k += 1 {
			if calls.Calls[k].Root && !isEmptyInlinedCall(k, calls) {
				kids = append(kids, k)
			}
		}
	}
	return kids
}

func inlinedVarTable(inlcalls *InlCalls) map[*Var]bool {
	vars := make(map[*Var]bool)
	for _, ic := range inlcalls.Calls {
		for _, v := range ic.InlVars {
			vars[v] = true
		}
	}
	return vars
}

// The s.Scopes slice contains variables were originally part of the
// function being emitted, as well as variables that were imported
// from various callee functions during the inlining process. This
// function prunes out any variables from the latter category (since
// they will be emitted as part of DWARF inlined_subroutine DIEs) and
// then generates scopes for vars in the former category.
func (psess *PackageSession) putPrunedScopes(ctxt Context, s *FnState, fnabbrev int) error {
	if len(s.Scopes) == 0 {
		return nil
	}
	scopes := make([]Scope, len(s.Scopes), len(s.Scopes))
	pvars := inlinedVarTable(&s.InlCalls)
	for k, s := range s.Scopes {
		var pruned Scope = Scope{Parent: s.Parent, Ranges: s.Ranges}
		for i := 0; i < len(s.Vars); i++ {
			_, found := pvars[s.Vars[i]]
			if !found {
				pruned.Vars = append(pruned.Vars, s.Vars[i])
			}
		}
		sort.Sort(byChildIndex(pruned.Vars))
		scopes[k] = pruned
	}
	var encbuf [20]byte
	if psess.putscope(ctxt, s, scopes, 0, fnabbrev, encbuf[:0]) < int32(len(scopes)) {
		return errors.New("multiple toplevel scopes")
	}
	return nil
}

// Emit DWARF attributes and child DIEs for an 'abstract' subprogram.
// The abstract subprogram DIE for a function contains its
// location-independent attributes (name, type, etc). Other instances
// of the function (any inlined copy of it, or the single out-of-line
// 'concrete' instance) will contain a pointer back to this abstract
// DIE (as a space-saving measure, so that name/type etc doesn't have
// to be repeated for each inlined copy).
func (psess *PackageSession) PutAbstractFunc(ctxt Context, s *FnState) error {

	if psess.logDwarf {
		ctxt.Logf("PutAbstractFunc(%v)\n", s.Absfn)
	}

	abbrev := DW_ABRV_FUNCTION_ABSTRACT
	psess.
		Uleb128put(ctxt, s.Absfn, int64(abbrev))

	fullname := s.Name
	if strings.HasPrefix(s.Name, "\"\".") {

		fullname = s.Importpath + "." + s.Name[3:]
	}
	psess.
		putattr(ctxt, s.Absfn, abbrev, DW_FORM_string, DW_CLS_STRING, int64(len(fullname)), fullname)
	psess.
		putattr(ctxt, s.Absfn, abbrev, DW_FORM_data1, DW_CLS_CONSTANT, int64(DW_INL_inlined), nil)

	var ev int64
	if s.External {
		ev = 1
	}
	psess.
		putattr(ctxt, s.Absfn, abbrev, DW_FORM_flag, DW_CLS_FLAG, ev, 0)

	// Child variables (may be empty)
	var flattened []*Var

	// This slice will hold the offset in bytes for each child var DIE
	// with respect to the start of the parent subprogram DIE.
	var offsets []int32

	if len(s.Scopes) > 0 {

		pvars := inlinedVarTable(&s.InlCalls)
		for _, scope := range s.Scopes {
			for i := 0; i < len(scope.Vars); i++ {
				_, found := pvars[scope.Vars[i]]
				if found || !scope.Vars[i].IsInAbstract {
					continue
				}
				flattened = append(flattened, scope.Vars[i])
			}
		}
		if len(flattened) > 0 {
			sort.Sort(byChildIndex(flattened))

			if psess.logDwarf {
				ctxt.Logf("putAbstractScope(%v): vars:", s.Info)
				for i, v := range flattened {
					ctxt.Logf(" %d:%s", i, v.Name)
				}
				ctxt.Logf("\n")
			}

			for _, v := range flattened {
				offsets = append(offsets, int32(ctxt.CurrentOffset(s.Absfn)))
				psess.
					putAbstractVar(ctxt, s.Absfn, v)
			}
		}
	}
	ctxt.RecordChildDieOffsets(s.Absfn, flattened, offsets)
	psess.
		Uleb128put(ctxt, s.Absfn, 0)
	return nil
}

// Emit DWARF attributes and child DIEs for an inlined subroutine. The
// first attribute of an inlined subroutine DIE is a reference back to
// its corresponding 'abstract' DIE (containing location-independent
// attributes such as name, type, etc). Inlined subroutine DIEs can
// have other inlined subroutine DIEs as children.
func (psess *PackageSession) PutInlinedFunc(ctxt Context, s *FnState, callersym Sym, callIdx int) error {
	ic := s.InlCalls.Calls[callIdx]
	callee := ic.AbsFunSym

	abbrev := DW_ABRV_INLINED_SUBROUTINE_RANGES
	if len(ic.Ranges) == 1 {
		abbrev = DW_ABRV_INLINED_SUBROUTINE
	}
	psess.
		Uleb128put(ctxt, s.Info, int64(abbrev))

	if psess.logDwarf {
		ctxt.Logf("PutInlinedFunc(caller=%v,callee=%v,abbrev=%d)\n", callersym, callee, abbrev)
	}
	psess.
		putattr(ctxt, s.Info, abbrev, DW_FORM_ref_addr, DW_CLS_REFERENCE, 0, callee)

	if abbrev == DW_ABRV_INLINED_SUBROUTINE_RANGES {
		psess.
			putattr(ctxt, s.Info, abbrev, DW_FORM_sec_offset, DW_CLS_PTR, s.Ranges.Len(), s.Ranges)
		PutRanges(ctxt, s.Ranges, s.StartPC, ic.Ranges)
	} else {
		st := ic.Ranges[0].Start
		en := ic.Ranges[0].End
		psess.
			putattr(ctxt, s.Info, abbrev, DW_FORM_addr, DW_CLS_ADDRESS, st, s.StartPC)
		psess.
			putattr(ctxt, s.Info, abbrev, DW_FORM_addr, DW_CLS_ADDRESS, en, s.StartPC)
	}

	ctxt.AddFileRef(s.Info, ic.CallFile)
	psess.
		putattr(ctxt, s.Info, abbrev, DW_FORM_udata, DW_CLS_CONSTANT, int64(ic.CallLine), nil)

	vars := ic.InlVars
	sort.Sort(byChildIndex(vars))
	inlIndex := ic.InlIndex
	var encbuf [20]byte
	for _, v := range vars {
		if !v.IsInAbstract {
			continue
		}
		psess.
			putvar(ctxt, s, v, callee, abbrev, inlIndex, encbuf[:0])
	}

	for _, sib := range inlChildren(callIdx, &s.InlCalls) {
		absfn := s.InlCalls.Calls[sib].AbsFunSym
		err := psess.PutInlinedFunc(ctxt, s, absfn, sib)
		if err != nil {
			return err
		}
	}
	psess.
		Uleb128put(ctxt, s.Info, 0)
	return nil
}

// Emit DWARF attributes and child DIEs for a 'concrete' subprogram,
// meaning the out-of-line copy of a function that was inlined at some
// point during the compilation of its containing package. The first
// attribute for a concrete DIE is a reference to the 'abstract' DIE
// for the function (which holds location-independent attributes such
// as name, type), then the remainder of the attributes are specific
// to this instance (location, frame base, etc).
func (psess *PackageSession) PutConcreteFunc(ctxt Context, s *FnState) error {
	if psess.logDwarf {
		ctxt.Logf("PutConcreteFunc(%v)\n", s.Info)
	}
	abbrev := DW_ABRV_FUNCTION_CONCRETE
	psess.
		Uleb128put(ctxt, s.Info, int64(abbrev))
	psess.
		putattr(ctxt, s.Info, abbrev, DW_FORM_ref_addr, DW_CLS_REFERENCE, 0, s.Absfn)
	psess.
		putattr(ctxt, s.Info, abbrev, DW_FORM_addr, DW_CLS_ADDRESS, 0, s.StartPC)
	psess.
		putattr(ctxt, s.Info, abbrev, DW_FORM_addr, DW_CLS_ADDRESS, s.Size, s.StartPC)
	psess.
		putattr(ctxt, s.Info, abbrev, DW_FORM_block1, DW_CLS_BLOCK, 1, []byte{DW_OP_call_frame_cfa})

	if err := psess.putPrunedScopes(ctxt, s, abbrev); err != nil {
		return err
	}

	for _, sib := range inlChildren(-1, &s.InlCalls) {
		absfn := s.InlCalls.Calls[sib].AbsFunSym
		err := psess.PutInlinedFunc(ctxt, s, absfn, sib)
		if err != nil {
			return err
		}
	}
	psess.
		Uleb128put(ctxt, s.Info, 0)
	return nil
}

// Emit DWARF attributes and child DIEs for a subprogram. Here
// 'default' implies that the function in question was not inlined
// when its containing package was compiled (hence there is no need to
// emit an abstract version for it to use as a base for inlined
// routine records).
func (psess *PackageSession) PutDefaultFunc(ctxt Context, s *FnState) error {
	if psess.logDwarf {
		ctxt.Logf("PutDefaultFunc(%v)\n", s.Info)
	}
	abbrev := DW_ABRV_FUNCTION
	psess.
		Uleb128put(ctxt, s.Info, int64(abbrev))
	psess.
		putattr(ctxt, s.Info, DW_ABRV_FUNCTION, DW_FORM_string, DW_CLS_STRING, int64(len(s.Name)), s.Name)
	psess.
		putattr(ctxt, s.Info, abbrev, DW_FORM_addr, DW_CLS_ADDRESS, 0, s.StartPC)
	psess.
		putattr(ctxt, s.Info, abbrev, DW_FORM_addr, DW_CLS_ADDRESS, s.Size, s.StartPC)
	psess.
		putattr(ctxt, s.Info, abbrev, DW_FORM_block1, DW_CLS_BLOCK, 1, []byte{DW_OP_call_frame_cfa})
	ctxt.AddFileRef(s.Info, s.Filesym)

	var ev int64
	if s.External {
		ev = 1
	}
	psess.
		putattr(ctxt, s.Info, abbrev, DW_FORM_flag, DW_CLS_FLAG, ev, 0)

	if err := psess.putPrunedScopes(ctxt, s, abbrev); err != nil {
		return err
	}

	for _, sib := range inlChildren(-1, &s.InlCalls) {
		absfn := s.InlCalls.Calls[sib].AbsFunSym
		err := psess.PutInlinedFunc(ctxt, s, absfn, sib)
		if err != nil {
			return err
		}
	}
	psess.
		Uleb128put(ctxt, s.Info, 0)
	return nil
}

func (psess *PackageSession) putscope(ctxt Context, s *FnState, scopes []Scope, curscope int32, fnabbrev int, encbuf []byte) int32 {

	if psess.logDwarf {
		ctxt.Logf("putscope(%v,%d): vars:", s.Info, curscope)
		for i, v := range scopes[curscope].Vars {
			ctxt.Logf(" %d:%d:%s", i, v.ChildIndex, v.Name)
		}
		ctxt.Logf("\n")
	}

	for _, v := range scopes[curscope].Vars {
		psess.
			putvar(ctxt, s, v, s.Absfn, fnabbrev, -1, encbuf)
	}
	this := curscope
	curscope++
	for curscope < int32(len(scopes)) {
		scope := scopes[curscope]
		if scope.Parent != this {
			return curscope
		}

		if len(scopes[curscope].Vars) == 0 {
			curscope = psess.putscope(ctxt, s, scopes, curscope, fnabbrev, encbuf)
			continue
		}

		if len(scope.Ranges) == 1 {
			psess.
				Uleb128put(ctxt, s.Info, DW_ABRV_LEXICAL_BLOCK_SIMPLE)
			psess.
				putattr(ctxt, s.Info, DW_ABRV_LEXICAL_BLOCK_SIMPLE, DW_FORM_addr, DW_CLS_ADDRESS, scope.Ranges[0].Start, s.StartPC)
			psess.
				putattr(ctxt, s.Info, DW_ABRV_LEXICAL_BLOCK_SIMPLE, DW_FORM_addr, DW_CLS_ADDRESS, scope.Ranges[0].End, s.StartPC)
		} else {
			psess.
				Uleb128put(ctxt, s.Info, DW_ABRV_LEXICAL_BLOCK_RANGES)
			psess.
				putattr(ctxt, s.Info, DW_ABRV_LEXICAL_BLOCK_RANGES, DW_FORM_sec_offset, DW_CLS_PTR, s.Ranges.Len(), s.Ranges)

			PutRanges(ctxt, s.Ranges, s.StartPC, scope.Ranges)
		}

		curscope = psess.putscope(ctxt, s, scopes, curscope, fnabbrev, encbuf)
		psess.
			Uleb128put(ctxt, s.Info, 0)
	}
	return curscope
}

// Given a default var abbrev code, select corresponding concrete code.
func concreteVarAbbrev(varAbbrev int) int {
	switch varAbbrev {
	case DW_ABRV_AUTO:
		return DW_ABRV_AUTO_CONCRETE
	case DW_ABRV_PARAM:
		return DW_ABRV_PARAM_CONCRETE
	case DW_ABRV_AUTO_LOCLIST:
		return DW_ABRV_AUTO_CONCRETE_LOCLIST
	case DW_ABRV_PARAM_LOCLIST:
		return DW_ABRV_PARAM_CONCRETE_LOCLIST
	default:
		panic("should never happen")
	}
}

// Pick the correct abbrev code for variable or parameter DIE.
func determineVarAbbrev(v *Var, fnabbrev int) (int, bool, bool) {
	abbrev := v.Abbrev

	missing := false
	switch {
	case abbrev == DW_ABRV_AUTO_LOCLIST && v.PutLocationList == nil:
		missing = true
		abbrev = DW_ABRV_AUTO
	case abbrev == DW_ABRV_PARAM_LOCLIST && v.PutLocationList == nil:
		missing = true
		abbrev = DW_ABRV_PARAM
	}

	concrete := true
	switch fnabbrev {
	case DW_ABRV_FUNCTION:
		concrete = false
		break
	case DW_ABRV_FUNCTION_CONCRETE:

		if !v.IsInAbstract {
			concrete = false
		}
	case DW_ABRV_INLINED_SUBROUTINE, DW_ABRV_INLINED_SUBROUTINE_RANGES:
	default:
		panic("should never happen")
	}

	if concrete {
		abbrev = concreteVarAbbrev(abbrev)
	}

	return abbrev, missing, concrete
}

func abbrevUsesLoclist(abbrev int) bool {
	switch abbrev {
	case DW_ABRV_AUTO_LOCLIST, DW_ABRV_AUTO_CONCRETE_LOCLIST,
		DW_ABRV_PARAM_LOCLIST, DW_ABRV_PARAM_CONCRETE_LOCLIST:
		return true
	default:
		return false
	}
}

// Emit DWARF attributes for a variable belonging to an 'abstract' subprogram.
func (psess *PackageSession) putAbstractVar(ctxt Context, info Sym, v *Var) {

	abbrev := v.Abbrev
	switch abbrev {
	case DW_ABRV_AUTO, DW_ABRV_AUTO_LOCLIST:
		abbrev = DW_ABRV_AUTO_ABSTRACT
	case DW_ABRV_PARAM, DW_ABRV_PARAM_LOCLIST:
		abbrev = DW_ABRV_PARAM_ABSTRACT
	}
	psess.
		Uleb128put(ctxt, info, int64(abbrev))
	psess.
		putattr(ctxt, info, abbrev, DW_FORM_string, DW_CLS_STRING, int64(len(v.Name)), v.Name)

	if abbrev == DW_ABRV_PARAM_ABSTRACT {
		var isReturn int64
		if v.IsReturnValue {
			isReturn = 1
		}
		psess.
			putattr(ctxt, info, abbrev, DW_FORM_flag, DW_CLS_FLAG, isReturn, nil)
	}

	if abbrev != DW_ABRV_PARAM_ABSTRACT {
		psess.
			putattr(ctxt, info, abbrev, DW_FORM_udata, DW_CLS_CONSTANT, int64(v.DeclLine), nil)
	}
	psess.
		putattr(ctxt, info, abbrev, DW_FORM_ref_addr, DW_CLS_REFERENCE, 0, v.Type)

}

func (psess *PackageSession) putvar(ctxt Context, s *FnState, v *Var, absfn Sym, fnabbrev, inlIndex int, encbuf []byte) {

	abbrev, missing, concrete := determineVarAbbrev(v, fnabbrev)
	psess.
		Uleb128put(ctxt, s.Info, int64(abbrev))

	if concrete {
		psess.
			putattr(ctxt, s.Info, abbrev, DW_FORM_ref_addr, DW_CLS_REFERENCE, 0, absfn)
		ctxt.RecordDclReference(s.Info, absfn, int(v.ChildIndex), inlIndex)
	} else {

		n := v.Name
		psess.
			putattr(ctxt, s.Info, abbrev, DW_FORM_string, DW_CLS_STRING, int64(len(n)), n)
		if abbrev == DW_ABRV_PARAM || abbrev == DW_ABRV_PARAM_LOCLIST || abbrev == DW_ABRV_PARAM_ABSTRACT {
			var isReturn int64
			if v.IsReturnValue {
				isReturn = 1
			}
			psess.
				putattr(ctxt, s.Info, abbrev, DW_FORM_flag, DW_CLS_FLAG, isReturn, nil)
		}
		psess.
			putattr(ctxt, s.Info, abbrev, DW_FORM_udata, DW_CLS_CONSTANT, int64(v.DeclLine), nil)
		psess.
			putattr(ctxt, s.Info, abbrev, DW_FORM_ref_addr, DW_CLS_REFERENCE, 0, v.Type)
	}

	if abbrevUsesLoclist(abbrev) {
		psess.
			putattr(ctxt, s.Info, abbrev, DW_FORM_sec_offset, DW_CLS_PTR, s.Loc.Len(), s.Loc)
		v.PutLocationList(s.Loc, s.StartPC)
	} else {
		loc := encbuf[:0]
		switch {
		case missing:
			break
		case v.StackOffset == 0:
			loc = append(loc, DW_OP_call_frame_cfa)
		default:
			loc = append(loc, DW_OP_fbreg)
			loc = AppendSleb128(loc, int64(v.StackOffset))
		}
		psess.
			putattr(ctxt, s.Info, abbrev, DW_FORM_block1, DW_CLS_BLOCK, int64(len(loc)), loc)
	}

}

// VarsByOffset attaches the methods of sort.Interface to []*Var,
// sorting in increasing StackOffset.
type VarsByOffset []*Var

func (s VarsByOffset) Len() int           { return len(s) }
func (s VarsByOffset) Less(i, j int) bool { return s[i].StackOffset < s[j].StackOffset }
func (s VarsByOffset) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// byChildIndex implements sort.Interface for []*dwarf.Var by child index.
type byChildIndex []*Var

func (s byChildIndex) Len() int           { return len(s) }
func (s byChildIndex) Less(i, j int) bool { return s[i].ChildIndex < s[j].ChildIndex }
func (s byChildIndex) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

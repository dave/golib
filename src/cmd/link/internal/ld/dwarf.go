package ld

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/dwarf"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"log"
	"strings"
)

type dwctxt struct {
	linkctxt *Link
}

func (c dwctxt) PtrSize() int {
	return c.linkctxt.Arch.PtrSize
}
func (c dwctxt) AddInt(s dwarf.Sym, size int, i int64) {
	ls := s.(*sym.Symbol)
	ls.AddUintXX(c.linkctxt.Arch, uint64(i), size)
}
func (c dwctxt) AddBytes(s dwarf.Sym, b []byte) {
	ls := s.(*sym.Symbol)
	ls.AddBytes(b)
}
func (c dwctxt) AddString(psess *PackageSession, s dwarf.Sym, v string) {
	psess.
		Addstring(s.(*sym.Symbol), v)
}

func (c dwctxt) AddAddress(s dwarf.Sym, data interface{}, value int64) {
	if value != 0 {
		value -= (data.(*sym.Symbol)).Value
	}
	s.(*sym.Symbol).AddAddrPlus(c.linkctxt.Arch, data.(*sym.Symbol), value)
}

func (c dwctxt) AddSectionOffset(psess *PackageSession, s dwarf.Sym, size int, t interface{}, ofs int64) {
	ls := s.(*sym.Symbol)
	switch size {
	default:
		psess.
			Errorf(ls, "invalid size %d in adddwarfref\n", size)
		fallthrough
	case c.linkctxt.Arch.PtrSize:
		ls.AddAddr(c.linkctxt.Arch, t.(*sym.Symbol))
	case 4:
		ls.AddAddrPlus4(t.(*sym.Symbol), 0)
	}
	r := &ls.R[len(ls.R)-1]
	r.Type = objabi.R_ADDROFF
	r.Add = ofs
}

func (c dwctxt) AddDWARFSectionOffset(psess *PackageSession, s dwarf.Sym, size int, t interface{}, ofs int64) {
	c.AddSectionOffset(psess, s, size, t, ofs)
	ls := s.(*sym.Symbol)
	ls.R[len(ls.R)-1].Type = objabi.R_DWARFSECREF
}

func (c dwctxt) Logf(format string, args ...interface{}) {
	c.linkctxt.Logf(format, args...)
}

func (c dwctxt) AddFileRef(s dwarf.Sym, f interface{}) {
	panic("should be used only in the compiler")
}

func (c dwctxt) CurrentOffset(s dwarf.Sym) int64 {
	panic("should be used only in the compiler")
}

func (c dwctxt) RecordDclReference(s dwarf.Sym, t dwarf.Sym, dclIdx int, inlIndex int) {
	panic("should be used only in the compiler")
}

func (c dwctxt) RecordChildDieOffsets(s dwarf.Sym, vars []*dwarf.Var, offsets []int32) {
	panic("should be used only in the compiler")
}

func (psess *PackageSession) writeabbrev(ctxt *Link) *sym.Symbol {
	s := ctxt.Syms.Lookup(".debug_abbrev", 0)
	s.Type = sym.SDWARFSECT
	s.AddBytes(psess.dwarf.GetAbbrev())
	return s
}

/*
 * Root DIEs for compilation units, types and global variables.
 */

func newattr(die *dwarf.DWDie, attr uint16, cls int, value int64, data interface{}) *dwarf.DWAttr {
	a := new(dwarf.DWAttr)
	a.Link = die.Attr
	die.Attr = a
	a.Atr = attr
	a.Cls = uint8(cls)
	a.Value = value
	a.Data = data
	return a
}

// Each DIE (except the root ones) has at least 1 attribute: its
// name. getattr moves the desired one to the front so
// frequently searched ones are found faster.
func getattr(die *dwarf.DWDie, attr uint16) *dwarf.DWAttr {
	if die.Attr.Atr == attr {
		return die.Attr
	}

	a := die.Attr
	b := a.Link
	for b != nil {
		if b.Atr == attr {
			a.Link = b.Link
			b.Link = die.Attr
			die.Attr = b
			return b
		}

		a = b
		b = b.Link
	}

	return nil
}

// Every DIE manufactured by the linker has at least an AT_name
// attribute (but it will only be written out if it is listed in the abbrev).
// The compiler does create nameless DWARF DIEs (ex: concrete subprogram
// instance).
func newdie(ctxt *Link, parent *dwarf.DWDie, abbrev int, name string, version int) *dwarf.DWDie {
	die := new(dwarf.DWDie)
	die.Abbrev = abbrev
	die.Link = parent.Child
	parent.Child = die

	newattr(die, dwarf.DW_AT_name, dwarf.DW_CLS_STRING, int64(len(name)), name)

	if name != "" && (abbrev <= dwarf.DW_ABRV_VARIABLE || abbrev >= dwarf.DW_ABRV_NULLTYPE) {
		if abbrev != dwarf.DW_ABRV_VARIABLE || version == 0 {
			if abbrev == dwarf.DW_ABRV_COMPUNIT {

				name = ".pkg." + name
			}
			s := ctxt.Syms.Lookup(dwarf.InfoPrefix+name, version)
			s.Attr |= sym.AttrNotInSymbolTable
			s.Type = sym.SDWARFINFO
			die.Sym = s
		}
	}

	return die
}

func walktypedef(die *dwarf.DWDie) *dwarf.DWDie {
	if die == nil {
		return nil
	}

	if die.Abbrev == dwarf.DW_ABRV_TYPEDECL {
		for attr := die.Attr; attr != nil; attr = attr.Link {
			if attr.Atr == dwarf.DW_AT_type && attr.Cls == dwarf.DW_CLS_REFERENCE && attr.Data != nil {
				return attr.Data.(*dwarf.DWDie)
			}
		}
	}

	return die
}

func walksymtypedef(ctxt *Link, s *sym.Symbol) *sym.Symbol {
	if t := ctxt.Syms.ROLookup(s.Name+"..def", int(s.Version)); t != nil {
		return t
	}
	return s
}

// Find child by AT_name using hashtable if available or linear scan
// if not.
func findchild(die *dwarf.DWDie, name string) *dwarf.DWDie {
	var prev *dwarf.DWDie
	for ; die != prev; prev, die = die, walktypedef(die) {
		for a := die.Child; a != nil; a = a.Link {
			if name == getattr(a, dwarf.DW_AT_name).Data {
				return a
			}
		}
		continue
	}
	return nil
}

// Used to avoid string allocation when looking up dwarf symbols

func (psess *PackageSession) find(ctxt *Link, name string) *sym.Symbol {
	n := append(psess.prefixBuf, name...)

	s := ctxt.Syms.ROLookup(string(n), 0)
	psess.
		prefixBuf = n[:len(dwarf.InfoPrefix)]
	if s != nil && s.Type == sym.SDWARFINFO {
		return s
	}
	return nil
}

func (psess *PackageSession) mustFind(ctxt *Link, name string) *sym.Symbol {
	r := psess.find(ctxt, name)
	if r == nil {
		psess.
			Exitf("dwarf find: cannot find %s", name)
	}
	return r
}

func (psess *PackageSession) adddwarfref(ctxt *Link, s *sym.Symbol, t *sym.Symbol, size int) int64 {
	var result int64
	switch size {
	default:
		psess.
			Errorf(s, "invalid size %d in adddwarfref\n", size)
		fallthrough
	case ctxt.Arch.PtrSize:
		result = s.AddAddr(ctxt.Arch, t)
	case 4:
		result = s.AddAddrPlus4(t, 0)
	}
	r := &s.R[len(s.R)-1]
	r.Type = objabi.R_DWARFSECREF
	return result
}

func newrefattr(die *dwarf.DWDie, attr uint16, ref *sym.Symbol) *dwarf.DWAttr {
	if ref == nil {
		return nil
	}
	return newattr(die, attr, dwarf.DW_CLS_REFERENCE, 0, ref)
}

func (psess *PackageSession) putdies(linkctxt *Link, ctxt dwarf.Context, syms []*sym.Symbol, die *dwarf.DWDie) []*sym.Symbol {
	for ; die != nil; die = die.Link {
		syms = psess.putdie(linkctxt, ctxt, syms, die)
	}
	syms[len(syms)-1].AddUint8(0)

	return syms
}

func dtolsym(s dwarf.Sym) *sym.Symbol {
	if s == nil {
		return nil
	}
	return s.(*sym.Symbol)
}

func (psess *PackageSession) putdie(linkctxt *Link, ctxt dwarf.Context, syms []*sym.Symbol, die *dwarf.DWDie) []*sym.Symbol {
	s := dtolsym(die.Sym)
	if s == nil {
		s = syms[len(syms)-1]
	} else {
		if s.Attr.OnList() {
			log.Fatalf("symbol %s listed multiple times", s.Name)
		}
		s.Attr |= sym.AttrOnList
		syms = append(syms, s)
	}
	psess.dwarf.
		Uleb128put(ctxt, s, int64(die.Abbrev))
	psess.dwarf.
		PutAttrs(ctxt, s, die.Abbrev, die.Attr)
	if psess.dwarf.HasChildren(die) {
		return psess.putdies(linkctxt, ctxt, syms, die.Child)
	}
	return syms
}

func reverselist(list **dwarf.DWDie) {
	curr := *list
	var prev *dwarf.DWDie
	for curr != nil {
		next := curr.Link
		curr.Link = prev
		prev = curr
		curr = next
	}

	*list = prev
}

func (psess *PackageSession) reversetree(list **dwarf.DWDie) {
	reverselist(list)
	for die := *list; die != nil; die = die.Link {
		if psess.dwarf.HasChildren(die) {
			psess.
				reversetree(&die.Child)
		}
	}
}

func newmemberoffsetattr(die *dwarf.DWDie, offs int32) {
	newattr(die, dwarf.DW_AT_data_member_location, dwarf.DW_CLS_CONSTANT, int64(offs), nil)
}

// GDB doesn't like FORM_addr for AT_location, so emit a
// location expression that evals to a const.
func newabslocexprattr(die *dwarf.DWDie, addr int64, sym *sym.Symbol) {
	newattr(die, dwarf.DW_AT_location, dwarf.DW_CLS_ADDRESS, addr, sym)

}

// Lookup predefined types
func (psess *PackageSession) lookupOrDiag(ctxt *Link, n string) *sym.Symbol {
	s := ctxt.Syms.ROLookup(n, 0)
	if s == nil || s.Size == 0 {
		psess.
			Exitf("dwarf: missing type: %s", n)
	}

	return s
}

func (psess *PackageSession) dotypedef(ctxt *Link, parent *dwarf.DWDie, name string, def *dwarf.DWDie) {

	if strings.HasPrefix(name, "map[") {
		return
	}
	if strings.HasPrefix(name, "struct {") {
		return
	}
	if strings.HasPrefix(name, "chan ") {
		return
	}
	if name[0] == '[' || name[0] == '*' {
		return
	}
	if def == nil {
		psess.
			Errorf(nil, "dwarf: bad def in dotypedef")
	}

	s := ctxt.Syms.Lookup(dtolsym(def.Sym).Name+"..def", 0)
	s.Attr |= sym.AttrNotInSymbolTable
	s.Type = sym.SDWARFINFO
	def.Sym = s

	die := newdie(ctxt, parent, dwarf.DW_ABRV_TYPEDECL, name, 0)

	newrefattr(die, dwarf.DW_AT_type, s)
}

// Define gotype, for composite ones recurse into constituents.
func (psess *PackageSession) defgotype(ctxt *Link, gotype *sym.Symbol) *sym.Symbol {
	if gotype == nil {
		return psess.mustFind(ctxt, "<unspecified>")
	}

	if !strings.HasPrefix(gotype.Name, "type.") {
		psess.
			Errorf(gotype, "dwarf: type name doesn't start with \"type.\"")
		return psess.mustFind(ctxt, "<unspecified>")
	}

	name := gotype.Name[5:]

	sdie := psess.find(ctxt, name)

	if sdie != nil {
		return sdie
	}

	return psess.newtype(ctxt, gotype).Sym.(*sym.Symbol)
}

func (psess *PackageSession) newtype(ctxt *Link, gotype *sym.Symbol) *dwarf.DWDie {
	name := gotype.Name[5:]
	kind := decodetypeKind(ctxt.Arch, gotype)
	bytesize := psess.decodetypeSize(ctxt.Arch, gotype)

	var die *dwarf.DWDie
	switch kind {
	case objabi.KindBool:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_BASETYPE, name, 0)
		newattr(die, dwarf.DW_AT_encoding, dwarf.DW_CLS_CONSTANT, dwarf.DW_ATE_boolean, 0)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)

	case objabi.KindInt,
		objabi.KindInt8,
		objabi.KindInt16,
		objabi.KindInt32,
		objabi.KindInt64:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_BASETYPE, name, 0)
		newattr(die, dwarf.DW_AT_encoding, dwarf.DW_CLS_CONSTANT, dwarf.DW_ATE_signed, 0)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)

	case objabi.KindUint,
		objabi.KindUint8,
		objabi.KindUint16,
		objabi.KindUint32,
		objabi.KindUint64,
		objabi.KindUintptr:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_BASETYPE, name, 0)
		newattr(die, dwarf.DW_AT_encoding, dwarf.DW_CLS_CONSTANT, dwarf.DW_ATE_unsigned, 0)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)

	case objabi.KindFloat32,
		objabi.KindFloat64:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_BASETYPE, name, 0)
		newattr(die, dwarf.DW_AT_encoding, dwarf.DW_CLS_CONSTANT, dwarf.DW_ATE_float, 0)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)

	case objabi.KindComplex64,
		objabi.KindComplex128:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_BASETYPE, name, 0)
		newattr(die, dwarf.DW_AT_encoding, dwarf.DW_CLS_CONSTANT, dwarf.DW_ATE_complex_float, 0)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)

	case objabi.KindArray:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_ARRAYTYPE, name, 0)
		psess.
			dotypedef(ctxt, &psess.dwtypes, name, die)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)
		s := decodetypeArrayElem(ctxt.Arch, gotype)
		newrefattr(die, dwarf.DW_AT_type, psess.defgotype(ctxt, s))
		fld := newdie(ctxt, die, dwarf.DW_ABRV_ARRAYRANGE, "range", 0)

		newattr(fld, dwarf.DW_AT_count, dwarf.DW_CLS_CONSTANT, psess.decodetypeArrayLen(ctxt.Arch, gotype), 0)

		newrefattr(fld, dwarf.DW_AT_type, psess.mustFind(ctxt, "uintptr"))

	case objabi.KindChan:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_CHANTYPE, name, 0)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)
		s := decodetypeChanElem(ctxt.Arch, gotype)
		newrefattr(die, dwarf.DW_AT_go_elem, psess.defgotype(ctxt, s))

		newrefattr(die, dwarf.DW_AT_type, s)

	case objabi.KindFunc:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_FUNCTYPE, name, 0)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)
		psess.
			dotypedef(ctxt, &psess.dwtypes, name, die)
		newrefattr(die, dwarf.DW_AT_type, psess.mustFind(ctxt, "void"))
		nfields := psess.decodetypeFuncInCount(ctxt.Arch, gotype)
		for i := 0; i < nfields; i++ {
			s := decodetypeFuncInType(ctxt.Arch, gotype, i)
			fld := newdie(ctxt, die, dwarf.DW_ABRV_FUNCTYPEPARAM, s.Name[5:], 0)
			newrefattr(fld, dwarf.DW_AT_type, psess.defgotype(ctxt, s))
		}

		if psess.decodetypeFuncDotdotdot(ctxt.Arch, gotype) {
			newdie(ctxt, die, dwarf.DW_ABRV_DOTDOTDOT, "...", 0)
		}
		nfields = psess.decodetypeFuncOutCount(ctxt.Arch, gotype)
		for i := 0; i < nfields; i++ {
			s := psess.decodetypeFuncOutType(ctxt.Arch, gotype, i)
			fld := newdie(ctxt, die, dwarf.DW_ABRV_FUNCTYPEPARAM, s.Name[5:], 0)
			newrefattr(fld, dwarf.DW_AT_type, psess.defptrto(ctxt, psess.defgotype(ctxt, s)))
		}

	case objabi.KindInterface:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_IFACETYPE, name, 0)
		psess.
			dotypedef(ctxt, &psess.dwtypes, name, die)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)
		nfields := int(psess.decodetypeIfaceMethodCount(ctxt.Arch, gotype))
		var s *sym.Symbol
		if nfields == 0 {
			s = psess.lookupOrDiag(ctxt, "type.runtime.eface")
		} else {
			s = psess.lookupOrDiag(ctxt, "type.runtime.iface")
		}
		newrefattr(die, dwarf.DW_AT_type, psess.defgotype(ctxt, s))

	case objabi.KindMap:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_MAPTYPE, name, 0)
		s := decodetypeMapKey(ctxt.Arch, gotype)
		newrefattr(die, dwarf.DW_AT_go_key, psess.defgotype(ctxt, s))
		s = decodetypeMapValue(ctxt.Arch, gotype)
		newrefattr(die, dwarf.DW_AT_go_elem, psess.defgotype(ctxt, s))

		newrefattr(die, dwarf.DW_AT_type, gotype)

	case objabi.KindPtr:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_PTRTYPE, name, 0)
		psess.
			dotypedef(ctxt, &psess.dwtypes, name, die)
		s := decodetypePtrElem(ctxt.Arch, gotype)
		newrefattr(die, dwarf.DW_AT_type, psess.defgotype(ctxt, s))

	case objabi.KindSlice:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_SLICETYPE, name, 0)
		psess.
			dotypedef(ctxt, &psess.dwtypes, name, die)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)
		s := decodetypeArrayElem(ctxt.Arch, gotype)
		elem := psess.defgotype(ctxt, s)
		newrefattr(die, dwarf.DW_AT_go_elem, elem)

	case objabi.KindString:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_STRINGTYPE, name, 0)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)

	case objabi.KindStruct:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_STRUCTTYPE, name, 0)
		psess.
			dotypedef(ctxt, &psess.dwtypes, name, die)
		newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, bytesize, 0)
		nfields := psess.decodetypeStructFieldCount(ctxt.Arch, gotype)
		for i := 0; i < nfields; i++ {
			f := decodetypeStructFieldName(ctxt.Arch, gotype, i)
			s := decodetypeStructFieldType(ctxt.Arch, gotype, i)
			if f == "" {
				f = s.Name[5:]
			}
			fld := newdie(ctxt, die, dwarf.DW_ABRV_STRUCTFIELD, f, 0)
			newrefattr(fld, dwarf.DW_AT_type, psess.defgotype(ctxt, s))
			offsetAnon := psess.decodetypeStructFieldOffsAnon(ctxt.Arch, gotype, i)
			newmemberoffsetattr(fld, int32(offsetAnon>>1))
			if offsetAnon&1 != 0 {
				newattr(fld, dwarf.DW_AT_go_embedded_field, dwarf.DW_CLS_FLAG, 1, 0)
			}
		}

	case objabi.KindUnsafePointer:
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_BARE_PTRTYPE, name, 0)

	default:
		psess.
			Errorf(gotype, "dwarf: definition of unknown kind %d", kind)
		die = newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_TYPEDECL, name, 0)
		newrefattr(die, dwarf.DW_AT_type, psess.mustFind(ctxt, "<unspecified>"))
	}

	newattr(die, dwarf.DW_AT_go_kind, dwarf.DW_CLS_CONSTANT, int64(kind), 0)
	if gotype.Attr.Reachable() {
		newattr(die, dwarf.DW_AT_go_runtime_type, dwarf.DW_CLS_GO_TYPEREF, 0, gotype)
	}

	if _, ok := psess.prototypedies[gotype.Name]; ok {
		psess.
			prototypedies[gotype.Name] = die
	}

	return die
}

func nameFromDIESym(dwtype *sym.Symbol) string {
	return strings.TrimSuffix(dwtype.Name[len(dwarf.InfoPrefix):], "..def")
}

// Find or construct *T given T.
func (psess *PackageSession) defptrto(ctxt *Link, dwtype *sym.Symbol) *sym.Symbol {
	ptrname := "*" + nameFromDIESym(dwtype)
	if die := psess.find(ctxt, ptrname); die != nil {
		return die
	}

	pdie := newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_PTRTYPE, ptrname, 0)
	newrefattr(pdie, dwarf.DW_AT_type, dwtype)

	gotype := ctxt.Syms.ROLookup("type."+ptrname, 0)
	if gotype != nil && gotype.Attr.Reachable() {
		newattr(pdie, dwarf.DW_AT_go_runtime_type, dwarf.DW_CLS_GO_TYPEREF, 0, gotype)
	}
	return dtolsym(pdie.Sym)
}

// Copies src's children into dst. Copies attributes by value.
// DWAttr.data is copied as pointer only. If except is one of
// the top-level children, it will not be copied.
func copychildrenexcept(ctxt *Link, dst *dwarf.DWDie, src *dwarf.DWDie, except *dwarf.DWDie) {
	for src = src.Child; src != nil; src = src.Link {
		if src == except {
			continue
		}
		c := newdie(ctxt, dst, src.Abbrev, getattr(src, dwarf.DW_AT_name).Data.(string), 0)
		for a := src.Attr; a != nil; a = a.Link {
			newattr(c, a.Atr, int(a.Cls), a.Value, a.Data)
		}
		copychildrenexcept(ctxt, c, src, nil)
	}

	reverselist(&dst.Child)
}

func copychildren(ctxt *Link, dst *dwarf.DWDie, src *dwarf.DWDie) {
	copychildrenexcept(ctxt, dst, src, nil)
}

// Search children (assumed to have TAG_member) for the one named
// field and set its AT_type to dwtype
func (psess *PackageSession) substitutetype(structdie *dwarf.DWDie, field string, dwtype *sym.Symbol) {
	child := findchild(structdie, field)
	if child == nil {
		psess.
			Exitf("dwarf substitutetype: %s does not have member %s",
				getattr(structdie, dwarf.DW_AT_name).Data, field)
		return
	}

	a := getattr(child, dwarf.DW_AT_type)
	if a != nil {
		a.Data = dwtype
	} else {
		newrefattr(child, dwarf.DW_AT_type, dwtype)
	}
}

func (psess *PackageSession) findprotodie(ctxt *Link, name string) *dwarf.DWDie {
	die, ok := psess.prototypedies[name]
	if ok && die == nil {
		psess.
			defgotype(ctxt, psess.lookupOrDiag(ctxt, name))
		die = psess.prototypedies[name]
	}
	return die
}

func (psess *PackageSession) synthesizestringtypes(ctxt *Link, die *dwarf.DWDie) {
	prototype := walktypedef(psess.findprotodie(ctxt, "type.runtime.stringStructDWARF"))
	if prototype == nil {
		return
	}

	for ; die != nil; die = die.Link {
		if die.Abbrev != dwarf.DW_ABRV_STRINGTYPE {
			continue
		}
		copychildren(ctxt, die, prototype)
	}
}

func (psess *PackageSession) synthesizeslicetypes(ctxt *Link, die *dwarf.DWDie) {
	prototype := walktypedef(psess.findprotodie(ctxt, "type.runtime.slice"))
	if prototype == nil {
		return
	}

	for ; die != nil; die = die.Link {
		if die.Abbrev != dwarf.DW_ABRV_SLICETYPE {
			continue
		}
		copychildren(ctxt, die, prototype)
		elem := getattr(die, dwarf.DW_AT_go_elem).Data.(*sym.Symbol)
		psess.
			substitutetype(die, "array", psess.defptrto(ctxt, elem))
	}
}

func mkinternaltypename(base string, arg1 string, arg2 string) string {
	if arg2 == "" {
		return fmt.Sprintf("%s<%s>", base, arg1)
	}
	return fmt.Sprintf("%s<%s,%s>", base, arg1, arg2)
}

// synthesizemaptypes is way too closely married to runtime/hashmap.c
const (
	MaxKeySize = 128
	MaxValSize = 128
	BucketSize = 8
)

func (psess *PackageSession) mkinternaltype(ctxt *Link, abbrev int, typename, keyname, valname string, f func(*dwarf.DWDie)) *sym.Symbol {
	name := mkinternaltypename(typename, keyname, valname)
	symname := dwarf.InfoPrefix + name
	s := ctxt.Syms.ROLookup(symname, 0)
	if s != nil && s.Type == sym.SDWARFINFO {
		return s
	}
	die := newdie(ctxt, &psess.dwtypes, abbrev, name, 0)
	f(die)
	return dtolsym(die.Sym)
}

func (psess *PackageSession) synthesizemaptypes(ctxt *Link, die *dwarf.DWDie) {
	hash := walktypedef(psess.findprotodie(ctxt, "type.runtime.hmap"))
	bucket := walktypedef(psess.findprotodie(ctxt, "type.runtime.bmap"))

	if hash == nil {
		return
	}

	for ; die != nil; die = die.Link {
		if die.Abbrev != dwarf.DW_ABRV_MAPTYPE {
			continue
		}
		gotype := getattr(die, dwarf.DW_AT_type).Data.(*sym.Symbol)
		keytype := decodetypeMapKey(ctxt.Arch, gotype)
		valtype := decodetypeMapValue(ctxt.Arch, gotype)
		keysize, valsize := psess.decodetypeSize(ctxt.Arch, keytype), psess.decodetypeSize(ctxt.Arch, valtype)
		keytype, valtype = walksymtypedef(ctxt, psess.defgotype(ctxt, keytype)), walksymtypedef(ctxt, psess.defgotype(ctxt, valtype))

		indirectKey, indirectVal := false, false
		if keysize > MaxKeySize {
			keysize = int64(ctxt.Arch.PtrSize)
			indirectKey = true
		}
		if valsize > MaxValSize {
			valsize = int64(ctxt.Arch.PtrSize)
			indirectVal = true
		}

		keyname := nameFromDIESym(keytype)
		dwhks := psess.mkinternaltype(ctxt, dwarf.DW_ABRV_ARRAYTYPE, "[]key", keyname, "", func(dwhk *dwarf.DWDie) {
			newattr(dwhk, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, BucketSize*keysize, 0)
			t := keytype
			if indirectKey {
				t = psess.defptrto(ctxt, keytype)
			}
			newrefattr(dwhk, dwarf.DW_AT_type, t)
			fld := newdie(ctxt, dwhk, dwarf.DW_ABRV_ARRAYRANGE, "size", 0)
			newattr(fld, dwarf.DW_AT_count, dwarf.DW_CLS_CONSTANT, BucketSize, 0)
			newrefattr(fld, dwarf.DW_AT_type, psess.mustFind(ctxt, "uintptr"))
		})

		valname := nameFromDIESym(valtype)
		dwhvs := psess.mkinternaltype(ctxt, dwarf.DW_ABRV_ARRAYTYPE, "[]val", valname, "", func(dwhv *dwarf.DWDie) {
			newattr(dwhv, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, BucketSize*valsize, 0)
			t := valtype
			if indirectVal {
				t = psess.defptrto(ctxt, valtype)
			}
			newrefattr(dwhv, dwarf.DW_AT_type, t)
			fld := newdie(ctxt, dwhv, dwarf.DW_ABRV_ARRAYRANGE, "size", 0)
			newattr(fld, dwarf.DW_AT_count, dwarf.DW_CLS_CONSTANT, BucketSize, 0)
			newrefattr(fld, dwarf.DW_AT_type, psess.mustFind(ctxt, "uintptr"))
		})

		dwhbs := psess.mkinternaltype(ctxt, dwarf.DW_ABRV_STRUCTTYPE, "bucket", keyname, valname, func(dwhb *dwarf.DWDie) {

			copychildrenexcept(ctxt, dwhb, bucket, findchild(bucket, "data"))

			fld := newdie(ctxt, dwhb, dwarf.DW_ABRV_STRUCTFIELD, "keys", 0)
			newrefattr(fld, dwarf.DW_AT_type, dwhks)
			newmemberoffsetattr(fld, BucketSize)
			fld = newdie(ctxt, dwhb, dwarf.DW_ABRV_STRUCTFIELD, "values", 0)
			newrefattr(fld, dwarf.DW_AT_type, dwhvs)
			newmemberoffsetattr(fld, BucketSize+BucketSize*int32(keysize))
			fld = newdie(ctxt, dwhb, dwarf.DW_ABRV_STRUCTFIELD, "overflow", 0)
			newrefattr(fld, dwarf.DW_AT_type, psess.defptrto(ctxt, dtolsym(dwhb.Sym)))
			newmemberoffsetattr(fld, BucketSize+BucketSize*(int32(keysize)+int32(valsize)))
			if ctxt.Arch.RegSize > ctxt.Arch.PtrSize {
				fld = newdie(ctxt, dwhb, dwarf.DW_ABRV_STRUCTFIELD, "pad", 0)
				newrefattr(fld, dwarf.DW_AT_type, psess.mustFind(ctxt, "uintptr"))
				newmemberoffsetattr(fld, BucketSize+BucketSize*(int32(keysize)+int32(valsize))+int32(ctxt.Arch.PtrSize))
			}

			newattr(dwhb, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, BucketSize+BucketSize*keysize+BucketSize*valsize+int64(ctxt.Arch.RegSize), 0)
		})

		dwhs := psess.mkinternaltype(ctxt, dwarf.DW_ABRV_STRUCTTYPE, "hash", keyname, valname, func(dwh *dwarf.DWDie) {
			copychildren(ctxt, dwh, hash)
			psess.
				substitutetype(dwh, "buckets", psess.defptrto(ctxt, dwhbs))
			psess.
				substitutetype(dwh, "oldbuckets", psess.defptrto(ctxt, dwhbs))
			newattr(dwh, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, getattr(hash, dwarf.DW_AT_byte_size).Value, nil)
		})

		newrefattr(die, dwarf.DW_AT_type, psess.defptrto(ctxt, dwhs))
	}
}

func (psess *PackageSession) synthesizechantypes(ctxt *Link, die *dwarf.DWDie) {
	sudog := walktypedef(psess.findprotodie(ctxt, "type.runtime.sudog"))
	waitq := walktypedef(psess.findprotodie(ctxt, "type.runtime.waitq"))
	hchan := walktypedef(psess.findprotodie(ctxt, "type.runtime.hchan"))
	if sudog == nil || waitq == nil || hchan == nil {
		return
	}

	sudogsize := int(getattr(sudog, dwarf.DW_AT_byte_size).Value)

	for ; die != nil; die = die.Link {
		if die.Abbrev != dwarf.DW_ABRV_CHANTYPE {
			continue
		}
		elemgotype := getattr(die, dwarf.DW_AT_type).Data.(*sym.Symbol)
		elemname := elemgotype.Name[5:]
		elemtype := walksymtypedef(ctxt, psess.defgotype(ctxt, elemgotype))

		dwss := psess.mkinternaltype(ctxt, dwarf.DW_ABRV_STRUCTTYPE, "sudog", elemname, "", func(dws *dwarf.DWDie) {
			copychildren(ctxt, dws, sudog)
			psess.
				substitutetype(dws, "elem", psess.defptrto(ctxt, elemtype))
			newattr(dws, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, int64(sudogsize), nil)
		})

		dwws := psess.mkinternaltype(ctxt, dwarf.DW_ABRV_STRUCTTYPE, "waitq", elemname, "", func(dww *dwarf.DWDie) {

			copychildren(ctxt, dww, waitq)
			psess.
				substitutetype(dww, "first", psess.defptrto(ctxt, dwss))
			psess.
				substitutetype(dww, "last", psess.defptrto(ctxt, dwss))
			newattr(dww, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, getattr(waitq, dwarf.DW_AT_byte_size).Value, nil)
		})

		dwhs := psess.mkinternaltype(ctxt, dwarf.DW_ABRV_STRUCTTYPE, "hchan", elemname, "", func(dwh *dwarf.DWDie) {
			copychildren(ctxt, dwh, hchan)
			psess.
				substitutetype(dwh, "recvq", dwws)
			psess.
				substitutetype(dwh, "sendq", dwws)
			newattr(dwh, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, getattr(hchan, dwarf.DW_AT_byte_size).Value, nil)
		})

		newrefattr(die, dwarf.DW_AT_type, psess.defptrto(ctxt, dwhs))
	}
}

// For use with pass.c::genasmsym
func (psess *PackageSession) defdwsymb(ctxt *Link, s *sym.Symbol, str string, t SymbolType, v int64, gotype *sym.Symbol) {
	if strings.HasPrefix(str, "go.string.") {
		return
	}
	if strings.HasPrefix(str, "runtime.gcbits.") {
		return
	}

	if strings.HasPrefix(str, "type.") && str != "type.*" && !strings.HasPrefix(str, "type..") {
		psess.
			defgotype(ctxt, s)
		return
	}

	var dv *dwarf.DWDie

	var dt *sym.Symbol
	switch t {
	default:
		return

	case DataSym, BSSSym:
		dv = newdie(ctxt, &psess.dwglobals, dwarf.DW_ABRV_VARIABLE, str, int(s.Version))
		newabslocexprattr(dv, v, s)
		if s.Version == 0 {
			newattr(dv, dwarf.DW_AT_external, dwarf.DW_CLS_FLAG, 1, 0)
		}
		fallthrough

	case AutoSym, ParamSym, DeletedAutoSym:
		dt = psess.defgotype(ctxt, gotype)
	}

	if dv != nil {
		newrefattr(dv, dwarf.DW_AT_type, dt)
	}
}

// compilationUnit is per-compilation unit (equivalently, per-package)
// debug-related data.
type compilationUnit struct {
	lib       *sym.Library
	consts    *sym.Symbol   // Package constants DIEs
	pcs       []dwarf.Range // PC ranges, relative to textp[0]
	dwinfo    *dwarf.DWDie  // CU root DIE
	funcDIEs  []*sym.Symbol // Function DIE subtrees
	absFnDIEs []*sym.Symbol // Abstract function DIE subtrees
}

// getCompilationUnits divides the symbols in ctxt.Textp by package.
func (psess *PackageSession) getCompilationUnits(ctxt *Link) []*compilationUnit {
	units := []*compilationUnit{}
	index := make(map[*sym.Library]*compilationUnit)
	var prevUnit *compilationUnit
	for _, s := range ctxt.Textp {
		if s.FuncInfo == nil {
			continue
		}
		unit := index[s.Lib]
		if unit == nil {
			unit = &compilationUnit{lib: s.Lib}
			if s := ctxt.Syms.ROLookup(dwarf.ConstInfoPrefix+s.Lib.Pkg, 0); s != nil {
				psess.
					importInfoSymbol(ctxt, s)
				unit.consts = s
			}
			units = append(units, unit)
			index[s.Lib] = unit
		}

		if prevUnit != unit {
			unit.pcs = append(unit.pcs, dwarf.Range{Start: s.Value - unit.lib.Textp[0].Value})
			prevUnit = unit
		}
		unit.pcs[len(unit.pcs)-1].End = s.Value - unit.lib.Textp[0].Value + s.Size
	}
	return units
}

func (psess *PackageSession) movetomodule(parent *dwarf.DWDie) {
	die := psess.dwroot.Child.Child
	if die == nil {
		psess.
			dwroot.Child.Child = parent.Child
		return
	}
	for die.Link != nil {
		die = die.Link
	}
	die.Link = parent.Child
}

// If the pcln table contains runtime/proc.go, use that to set gdbscript path.
func (psess *PackageSession) finddebugruntimepath(s *sym.Symbol) {
	if psess.gdbscript != "" {
		return
	}

	for i := range s.FuncInfo.File {
		f := s.FuncInfo.File[i]

		if i := strings.Index(f.Name, "runtime/proc.go"); i >= 0 {
			psess.
				gdbscript = f.Name[:i] + "runtime/runtime-gdb.py"
			break
		}
	}
}

/*
 * Generate a sequence of opcodes that is as short as possible.
 * See section 6.2.5
 */
const (
	LINE_BASE   = -4
	LINE_RANGE  = 10
	PC_RANGE    = (255 - OPCODE_BASE) / LINE_RANGE
	OPCODE_BASE = 11
)

func (psess *PackageSession) putpclcdelta(linkctxt *Link, ctxt dwarf.Context, s *sym.Symbol, deltaPC uint64, deltaLC int64) {
	// Choose a special opcode that minimizes the number of bytes needed to
	// encode the remaining PC delta and LC delta.
	var opcode int64
	if deltaLC < LINE_BASE {
		if deltaPC >= PC_RANGE {
			opcode = OPCODE_BASE + (LINE_RANGE * PC_RANGE)
		} else {
			opcode = OPCODE_BASE + (LINE_RANGE * int64(deltaPC))
		}
	} else if deltaLC < LINE_BASE+LINE_RANGE {
		if deltaPC >= PC_RANGE {
			opcode = OPCODE_BASE + (deltaLC - LINE_BASE) + (LINE_RANGE * PC_RANGE)
			if opcode > 255 {
				opcode -= LINE_RANGE
			}
		} else {
			opcode = OPCODE_BASE + (deltaLC - LINE_BASE) + (LINE_RANGE * int64(deltaPC))
		}
	} else {
		if deltaPC <= PC_RANGE {
			opcode = OPCODE_BASE + (LINE_RANGE - 1) + (LINE_RANGE * int64(deltaPC))
			if opcode > 255 {
				opcode = 255
			}
		} else {

			switch deltaPC - PC_RANGE {

			case PC_RANGE, (1 << 7) - 1, (1 << 16) - 1, (1 << 21) - 1, (1 << 28) - 1,
				(1 << 35) - 1, (1 << 42) - 1, (1 << 49) - 1, (1 << 56) - 1, (1 << 63) - 1:
				opcode = 255
			default:
				opcode = OPCODE_BASE + LINE_RANGE*PC_RANGE - 1
			}
		}
	}
	if opcode < OPCODE_BASE || opcode > 255 {
		panic(fmt.Sprintf("produced invalid special opcode %d", opcode))
	}

	deltaPC -= uint64((opcode - OPCODE_BASE) / LINE_RANGE)
	deltaLC -= (opcode-OPCODE_BASE)%LINE_RANGE + LINE_BASE

	if deltaPC != 0 {
		if deltaPC <= PC_RANGE {

			opcode -= LINE_RANGE * int64(PC_RANGE-deltaPC)
			if opcode < OPCODE_BASE {
				panic(fmt.Sprintf("produced invalid special opcode %d", opcode))
			}
			s.AddUint8(dwarf.DW_LNS_const_add_pc)
		} else if (1<<14) <= deltaPC && deltaPC < (1<<16) {
			s.AddUint8(dwarf.DW_LNS_fixed_advance_pc)
			s.AddUint16(linkctxt.Arch, uint16(deltaPC))
		} else {
			s.AddUint8(dwarf.DW_LNS_advance_pc)
			psess.dwarf.
				Uleb128put(ctxt, s, int64(deltaPC))
		}
	}

	if deltaLC != 0 {
		s.AddUint8(dwarf.DW_LNS_advance_line)
		psess.dwarf.
			Sleb128put(ctxt, s, deltaLC)
	}

	s.AddUint8(uint8(opcode))
}

func getCompilationDir() string {

	return "."
}

func (psess *PackageSession) importInfoSymbol(ctxt *Link, dsym *sym.Symbol) {
	dsym.Attr |= sym.AttrNotInSymbolTable | sym.AttrReachable
	dsym.Type = sym.SDWARFINFO
	for _, r := range dsym.R {
		if r.Type == objabi.R_DWARFSECREF && r.Sym.Size == 0 {
			if ctxt.BuildMode == BuildModeShared {

				continue
			}
			n := nameFromDIESym(r.Sym)
			psess.
				defgotype(ctxt, ctxt.Syms.Lookup("type."+n, 0))
		}
	}
}

func (psess *PackageSession) collectAbstractFunctions(ctxt *Link, fn *sym.Symbol, dsym *sym.Symbol, absfuncs []*sym.Symbol) []*sym.Symbol {

	var newabsfns []*sym.Symbol

	for _, reloc := range dsym.R {
		candsym := reloc.Sym
		if reloc.Type != objabi.R_DWARFSECREF {
			continue
		}
		if !strings.HasPrefix(candsym.Name, dwarf.InfoPrefix) {
			continue
		}
		if !strings.HasSuffix(candsym.Name, dwarf.AbstractFuncSuffix) {
			continue
		}
		if candsym.Attr.OnList() {
			continue
		}
		candsym.Attr |= sym.AttrOnList
		newabsfns = append(newabsfns, candsym)
	}

	for _, absdsym := range newabsfns {
		psess.
			importInfoSymbol(ctxt, absdsym)
		absfuncs = append(absfuncs, absdsym)
	}

	return absfuncs
}

func (psess *PackageSession) writelines(ctxt *Link, lib *sym.Library, textp []*sym.Symbol, ls *sym.Symbol) (dwinfo *dwarf.DWDie, funcs []*sym.Symbol, absfuncs []*sym.Symbol) {

	var dwarfctxt dwarf.Context = dwctxt{ctxt}
	is_stmt := uint8(1)

	unitstart := int64(-1)
	headerstart := int64(-1)
	headerend := int64(-1)

	lang := dwarf.DW_LANG_Go

	dwinfo = newdie(ctxt, &psess.dwroot, dwarf.DW_ABRV_COMPUNIT, lib.Pkg, 0)
	newattr(dwinfo, dwarf.DW_AT_language, dwarf.DW_CLS_CONSTANT, int64(lang), 0)
	newattr(dwinfo, dwarf.DW_AT_stmt_list, dwarf.DW_CLS_PTR, ls.Size, ls)

	compDir := getCompilationDir()

	newattr(dwinfo, dwarf.DW_AT_comp_dir, dwarf.DW_CLS_STRING, int64(len(compDir)), compDir)
	producerExtra := ctxt.Syms.Lookup(dwarf.CUInfoPrefix+"producer."+lib.Pkg, 0)
	producer := "Go github.com/dave/golib/src/cmd/compile " + psess.objabi.Version
	if len(producerExtra.P) > 0 {

		producer += "; " + string(producerExtra.P)
	}
	newattr(dwinfo, dwarf.DW_AT_producer, dwarf.DW_CLS_STRING, int64(len(producer)), producer)

	unitLengthOffset := ls.Size
	ls.AddUint32(ctxt.Arch, 0)
	unitstart = ls.Size
	ls.AddUint16(ctxt.Arch, 2)
	headerLengthOffset := ls.Size
	ls.AddUint32(ctxt.Arch, 0)
	headerstart = ls.Size

	ls.AddUint8(1)
	ls.AddUint8(is_stmt)
	ls.AddUint8(LINE_BASE & 0xFF)
	ls.AddUint8(LINE_RANGE)
	ls.AddUint8(OPCODE_BASE)
	ls.AddUint8(0)
	ls.AddUint8(1)
	ls.AddUint8(1)
	ls.AddUint8(1)
	ls.AddUint8(1)
	ls.AddUint8(0)
	ls.AddUint8(0)
	ls.AddUint8(0)
	ls.AddUint8(1)
	ls.AddUint8(0)
	ls.AddUint8(0)

	fileNums := make(map[int]int)
	for _, s := range textp {
		for _, f := range s.FuncInfo.File {
			if _, ok := fileNums[int(f.Value)]; ok {
				continue
			}

			fileNums[int(f.Value)] = len(fileNums) + 1
			psess.
				Addstring(ls, f.Name)
			ls.AddUint8(0)
			ls.AddUint8(0)
			ls.AddUint8(0)
		}

		dsym := ctxt.Syms.Lookup(dwarf.InfoPrefix+s.Name, int(s.Version))
		psess.
			importInfoSymbol(ctxt, dsym)
		for ri := range dsym.R {
			r := &dsym.R[ri]
			if r.Type != objabi.R_DWARFFILEREF {
				continue
			}
			_, ok := fileNums[int(r.Sym.Value)]
			if !ok {
				fileNums[int(r.Sym.Value)] = len(fileNums) + 1
				psess.
					Addstring(ls, r.Sym.Name)
				ls.AddUint8(0)
				ls.AddUint8(0)
				ls.AddUint8(0)
			}
		}
	}

	ls.AddUint8(0)

	headerend = ls.Size

	ls.AddUint8(0)
	psess.dwarf.
		Uleb128put(dwarfctxt, ls, 1+int64(ctxt.Arch.PtrSize))
	ls.AddUint8(dwarf.DW_LNE_set_address)

	s := textp[0]
	pc := s.Value
	line := 1
	file := 1
	ls.AddAddr(ctxt.Arch, s)

	var pcfile Pciter
	var pcline Pciter
	var pcstmt Pciter
	for i, s := range textp {
		dsym := ctxt.Syms.Lookup(dwarf.InfoPrefix+s.Name, int(s.Version))
		funcs = append(funcs, dsym)
		absfuncs = psess.collectAbstractFunctions(ctxt, s, dsym, absfuncs)
		psess.
			finddebugruntimepath(s)

		isStmtsSym := ctxt.Syms.ROLookup(dwarf.IsStmtPrefix+s.Name, int(s.Version))
		pctostmtData := sym.Pcdata{P: isStmtsSym.P}

		pciterinit(ctxt, &pcfile, &s.FuncInfo.Pcfile)
		pciterinit(ctxt, &pcline, &s.FuncInfo.Pcline)
		pciterinit(ctxt, &pcstmt, &pctostmtData)

		if pcstmt.done != 0 {

			pcstmt.value = 1
		}

		var thispc uint32

		for pcfile.done == 0 && pcline.done == 0 {

			if int32(file) != pcfile.value {
				ls.AddUint8(dwarf.DW_LNS_set_file)
				idx, ok := fileNums[int(pcfile.value)]
				if !ok {
					psess.
						Exitf("pcln table file missing from DWARF line table")
				}
				psess.dwarf.
					Uleb128put(dwarfctxt, ls, int64(idx))
				file = int(pcfile.value)
			}

			if is_stmt != uint8(pcstmt.value) {
				new_stmt := uint8(pcstmt.value)
				switch new_stmt &^ 1 {
				case obj.PrologueEnd:
					ls.AddUint8(uint8(dwarf.DW_LNS_set_prologue_end))
				case obj.EpilogueBegin:

				}
				new_stmt &= 1
				if is_stmt != new_stmt {
					is_stmt = new_stmt
					ls.AddUint8(uint8(dwarf.DW_LNS_negate_stmt))
				}
			}
			psess.
				putpclcdelta(ctxt, dwarfctxt, ls, uint64(s.Value+int64(thispc)-pc), int64(pcline.value)-int64(line))

			pc = s.Value + int64(thispc)
			line = int(pcline.value)

			thispc = pcfile.nextpc
			if pcline.nextpc < thispc {
				thispc = pcline.nextpc
			}
			if pcstmt.done == 0 && pcstmt.nextpc < thispc {
				thispc = pcstmt.nextpc
			}

			if pcfile.nextpc == thispc {
				pciternext(&pcfile)
			}
			if pcstmt.done == 0 && pcstmt.nextpc == thispc {
				pciternext(&pcstmt)
			}
			if pcline.nextpc == thispc {
				pciternext(&pcline)
			}
		}
		if is_stmt == 0 && i < len(textp)-1 {

			is_stmt = 1
			ls.AddUint8(uint8(dwarf.DW_LNS_negate_stmt))
		}
	}

	ls.AddUint8(0)
	psess.dwarf.
		Uleb128put(dwarfctxt, ls, 1)
	ls.AddUint8(dwarf.DW_LNE_end_sequence)

	ls.SetUint32(ctxt.Arch, unitLengthOffset, uint32(ls.Size-unitstart))
	ls.SetUint32(ctxt.Arch, headerLengthOffset, uint32(headerend-headerstart))

	missing := make(map[int]interface{})
	for _, f := range funcs {
		for ri := range f.R {
			r := &f.R[ri]
			if r.Type != objabi.R_DWARFFILEREF {
				continue
			}

			r.Done = true
			idx, ok := fileNums[int(r.Sym.Value)]
			if ok {
				if int(int32(idx)) != idx {
					psess.
						Errorf(f, "bad R_DWARFFILEREF relocation: file index overflow")
				}
				if r.Siz != 4 {
					psess.
						Errorf(f, "bad R_DWARFFILEREF relocation: has size %d, expected 4", r.Siz)
				}
				if r.Off < 0 || r.Off+4 > int32(len(f.P)) {
					psess.
						Errorf(f, "bad R_DWARFFILEREF relocation offset %d + 4 would write past length %d", r.Off, len(s.P))
					continue
				}
				ctxt.Arch.ByteOrder.PutUint32(f.P[r.Off:r.Off+4], uint32(idx))
			} else {
				_, found := missing[int(r.Sym.Value)]
				if !found {
					psess.
						Errorf(f, "R_DWARFFILEREF relocation file missing: %v idx %d", r.Sym, r.Sym.Value)
					missing[int(r.Sym.Value)] = nil
				}
			}
		}
	}

	return dwinfo, funcs, absfuncs
}

// writepcranges generates the DW_AT_ranges table for compilation unit cu.
func writepcranges(ctxt *Link, cu *dwarf.DWDie, base *sym.Symbol, pcs []dwarf.Range, ranges *sym.Symbol) {
	var dwarfctxt dwarf.Context = dwctxt{ctxt}

	newattr(cu, dwarf.DW_AT_ranges, dwarf.DW_CLS_PTR, ranges.Size, ranges)
	newattr(cu, dwarf.DW_AT_low_pc, dwarf.DW_CLS_ADDRESS, base.Value, base)
	dwarf.PutRanges(dwarfctxt, ranges, nil, pcs)
}

/*
 *  Emit .debug_frame
 */
const (
	dataAlignmentFactor = -4
)

// appendPCDeltaCFA appends per-PC CFA deltas to b and returns the final slice.
func appendPCDeltaCFA(arch *sys.Arch, b []byte, deltapc, cfa int64) []byte {
	b = append(b, dwarf.DW_CFA_def_cfa_offset_sf)
	b = dwarf.AppendSleb128(b, cfa/dataAlignmentFactor)

	switch {
	case deltapc < 0x40:
		b = append(b, uint8(dwarf.DW_CFA_advance_loc+deltapc))
	case deltapc < 0x100:
		b = append(b, dwarf.DW_CFA_advance_loc1)
		b = append(b, uint8(deltapc))
	case deltapc < 0x10000:
		b = append(b, dwarf.DW_CFA_advance_loc2, 0, 0)
		arch.ByteOrder.PutUint16(b[len(b)-2:], uint16(deltapc))
	default:
		b = append(b, dwarf.DW_CFA_advance_loc4, 0, 0, 0, 0)
		arch.ByteOrder.PutUint32(b[len(b)-4:], uint32(deltapc))
	}
	return b
}

func (psess *PackageSession) writeframes(ctxt *Link, syms []*sym.Symbol) []*sym.Symbol {
	var dwarfctxt dwarf.Context = dwctxt{ctxt}
	fs := ctxt.Syms.Lookup(".debug_frame", 0)
	fs.Type = sym.SDWARFSECT
	syms = append(syms, fs)

	cieReserve := uint32(16)
	if haslinkregister(ctxt) {
		cieReserve = 32
	}
	fs.AddUint32(ctxt.Arch, cieReserve)
	fs.AddUint32(ctxt.Arch, 0xffffffff)
	fs.AddUint8(3)
	fs.AddUint8(0)
	psess.dwarf.
		Uleb128put(dwarfctxt, fs, 1)
	psess.dwarf.
		Sleb128put(dwarfctxt, fs, dataAlignmentFactor)
	psess.dwarf.
		Uleb128put(dwarfctxt, fs, int64(psess.thearch.Dwarfreglr))

	fs.AddUint8(dwarf.DW_CFA_def_cfa)
	psess.dwarf.
		Uleb128put(dwarfctxt, fs, int64(psess.thearch.Dwarfregsp))
	if haslinkregister(ctxt) {
		psess.dwarf.
			Uleb128put(dwarfctxt, fs, int64(0))

		fs.AddUint8(dwarf.DW_CFA_same_value)
		psess.dwarf.
			Uleb128put(dwarfctxt, fs, int64(psess.thearch.Dwarfreglr))

		fs.AddUint8(dwarf.DW_CFA_val_offset)
		psess.dwarf.
			Uleb128put(dwarfctxt, fs, int64(psess.thearch.Dwarfregsp))
		psess.dwarf.
			Uleb128put(dwarfctxt, fs, int64(0))
	} else {
		psess.dwarf.
			Uleb128put(dwarfctxt, fs, int64(ctxt.Arch.PtrSize))

		fs.AddUint8(dwarf.DW_CFA_offset_extended)
		psess.dwarf.
			Uleb128put(dwarfctxt, fs, int64(psess.thearch.Dwarfreglr))
		psess.dwarf.
			Uleb128put(dwarfctxt, fs, int64(-ctxt.Arch.PtrSize)/dataAlignmentFactor)
	}

	pad := int64(cieReserve) + 4 - fs.Size

	if pad < 0 {
		psess.
			Exitf("dwarf: cieReserve too small by %d bytes.", -pad)
	}

	fs.AddBytes(psess.zeros[:pad])

	var deltaBuf []byte
	var pcsp Pciter
	for _, s := range ctxt.Textp {
		if s.FuncInfo == nil {
			continue
		}

		deltaBuf = deltaBuf[:0]
		for pciterinit(ctxt, &pcsp, &s.FuncInfo.Pcsp); pcsp.done == 0; pciternext(&pcsp) {
			nextpc := pcsp.nextpc

			if int64(nextpc) == s.Size {
				nextpc--
				if nextpc < pcsp.pc {
					continue
				}
			}

			if haslinkregister(ctxt) {

				if pcsp.value > 0 {

					deltaBuf = append(deltaBuf, dwarf.DW_CFA_offset_extended_sf)
					deltaBuf = dwarf.AppendUleb128(deltaBuf, uint64(psess.thearch.Dwarfreglr))
					deltaBuf = dwarf.AppendSleb128(deltaBuf, -int64(pcsp.value)/dataAlignmentFactor)
				} else {

					deltaBuf = append(deltaBuf, dwarf.DW_CFA_same_value)
					deltaBuf = dwarf.AppendUleb128(deltaBuf, uint64(psess.thearch.Dwarfreglr))
				}
				deltaBuf = appendPCDeltaCFA(ctxt.Arch, deltaBuf, int64(nextpc)-int64(pcsp.pc), int64(pcsp.value))
			} else {
				deltaBuf = appendPCDeltaCFA(ctxt.Arch, deltaBuf, int64(nextpc)-int64(pcsp.pc), int64(ctxt.Arch.PtrSize)+int64(pcsp.value))
			}
		}
		pad := int(Rnd(int64(len(deltaBuf)), int64(ctxt.Arch.PtrSize))) - len(deltaBuf)
		deltaBuf = append(deltaBuf, psess.zeros[:pad]...)

		fs.AddUint32(ctxt.Arch, uint32(4+2*ctxt.Arch.PtrSize+len(deltaBuf)))
		if ctxt.LinkMode == LinkExternal {
			psess.
				adddwarfref(ctxt, fs, fs, 4)
		} else {
			fs.AddUint32(ctxt.Arch, 0)
		}
		fs.AddAddr(ctxt.Arch, s)
		fs.AddUintXX(ctxt.Arch, uint64(s.Size), ctxt.Arch.PtrSize)
		fs.AddBytes(deltaBuf)
	}
	return syms
}

func writeranges(ctxt *Link, syms []*sym.Symbol) []*sym.Symbol {
	for _, s := range ctxt.Textp {
		rangeSym := ctxt.Syms.ROLookup(dwarf.RangePrefix+s.Name, int(s.Version))
		if rangeSym == nil || rangeSym.Size == 0 {
			continue
		}
		rangeSym.Attr |= sym.AttrReachable | sym.AttrNotInSymbolTable
		rangeSym.Type = sym.SDWARFRANGE

		if ctxt.HeadType == objabi.Hdarwin {
			fn := ctxt.Syms.ROLookup(dwarf.InfoPrefix+s.Name, int(s.Version))
			removeDwarfAddrListBaseAddress(ctxt, fn, rangeSym, false)
		}
		syms = append(syms, rangeSym)
	}
	return syms
}

/*
 *  Walk DWarfDebugInfoEntries, and emit .debug_info
 */
const (
	COMPUNITHEADERSIZE = 4 + 2 + 4 + 1
)

func (psess *PackageSession) writeinfo(ctxt *Link, syms []*sym.Symbol, units []*compilationUnit, abbrevsym *sym.Symbol) []*sym.Symbol {
	infosec := ctxt.Syms.Lookup(".debug_info", 0)
	infosec.Type = sym.SDWARFINFO
	infosec.Attr |= sym.AttrReachable
	syms = append(syms, infosec)

	var dwarfctxt dwarf.Context = dwctxt{ctxt}

	unitByDIE := make(map[*dwarf.DWDie]*compilationUnit)
	for _, u := range units {
		unitByDIE[u.dwinfo] = u
	}

	for compunit := psess.dwroot.Child; compunit != nil; compunit = compunit.Link {
		s := dtolsym(compunit.Sym)
		u := unitByDIE[compunit]

		s.AddUint32(ctxt.Arch, 0)
		s.AddUint16(ctxt.Arch, 4)
		psess.
			adddwarfref(ctxt, s, abbrevsym, 4)

		s.AddUint8(uint8(ctxt.Arch.PtrSize))
		psess.dwarf.
			Uleb128put(dwarfctxt, s, int64(compunit.Abbrev))
		psess.dwarf.
			PutAttrs(dwarfctxt, s, compunit.Abbrev, compunit.Attr)

		cu := []*sym.Symbol{s}
		cu = append(cu, u.absFnDIEs...)
		cu = append(cu, u.funcDIEs...)
		if u.consts != nil {
			cu = append(cu, u.consts)
		}
		cu = psess.putdies(ctxt, dwarfctxt, cu, compunit.Child)
		var cusize int64
		for _, child := range cu {
			cusize += child.Size
		}
		cusize -= 4
		s.SetUint32(ctxt.Arch, 0, uint32(cusize))

		newattr(compunit, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, cusize, 0)
		syms = append(syms, cu...)
	}
	return syms
}

/*
 *  Emit .debug_pubnames/_types.  _info must have been written before,
 *  because we need die->offs and infoo/infosize;
 */
func ispubname(die *dwarf.DWDie) bool {
	switch die.Abbrev {
	case dwarf.DW_ABRV_FUNCTION, dwarf.DW_ABRV_VARIABLE:
		a := getattr(die, dwarf.DW_AT_external)
		return a != nil && a.Value != 0
	}

	return false
}

func ispubtype(die *dwarf.DWDie) bool {
	return die.Abbrev >= dwarf.DW_ABRV_NULLTYPE
}

func (psess *PackageSession) writepub(ctxt *Link, sname string, ispub func(*dwarf.DWDie) bool, syms []*sym.Symbol) []*sym.Symbol {
	s := ctxt.Syms.Lookup(sname, 0)
	s.Type = sym.SDWARFSECT
	syms = append(syms, s)

	for compunit := psess.dwroot.Child; compunit != nil; compunit = compunit.Link {
		sectionstart := s.Size
		culength := uint32(getattr(compunit, dwarf.DW_AT_byte_size).Value) + 4

		s.AddUint32(ctxt.Arch, 0)
		s.AddUint16(ctxt.Arch, 2)
		psess.
			adddwarfref(ctxt, s, dtolsym(compunit.Sym), 4)
		s.AddUint32(ctxt.Arch, culength)

		for die := compunit.Child; die != nil; die = die.Link {
			if !ispub(die) {
				continue
			}
			dwa := getattr(die, dwarf.DW_AT_name)
			name := dwa.Data.(string)
			if die.Sym == nil {
				fmt.Println("Missing sym for ", name)
			}
			psess.
				adddwarfref(ctxt, s, dtolsym(die.Sym), 4)
			psess.
				Addstring(s, name)
		}

		s.AddUint32(ctxt.Arch, 0)

		s.SetUint32(ctxt.Arch, sectionstart, uint32(s.Size-sectionstart)-4)
	}

	return syms
}

func (psess *PackageSession) writegdbscript(ctxt *Link, syms []*sym.Symbol) []*sym.Symbol {
	if ctxt.LinkMode == LinkExternal && ctxt.HeadType == objabi.Hwindows && ctxt.BuildMode == BuildModeCArchive {

		return syms
	}

	if psess.gdbscript != "" {
		s := ctxt.Syms.Lookup(".debug_gdb_scripts", 0)
		s.Type = sym.SDWARFSECT
		syms = append(syms, s)
		s.AddUint8(1)
		psess.
			Addstring(s, psess.gdbscript)
	}

	return syms
}

/*
 * This is the main entry point for generating dwarf.  After emitting
 * the mandatory debug_abbrev section, it calls writelines() to set up
 * the per-compilation unit part of the DIE tree, while simultaneously
 * emitting the debug_line section.  When the final tree contains
 * forward references, it will write the debug_info section in 2
 * passes.
 *
 */
func (psess *PackageSession) dwarfgeneratedebugsyms(ctxt *Link) {
	if *psess.FlagW {
		return
	}
	if *psess.FlagS && ctxt.HeadType != objabi.Hdarwin {
		return
	}
	if ctxt.HeadType == objabi.Hplan9 || ctxt.HeadType == objabi.Hjs {
		return
	}

	if ctxt.LinkMode == LinkExternal {
		switch {
		case ctxt.IsELF:
		case ctxt.HeadType == objabi.Hdarwin:
		case ctxt.HeadType == objabi.Hwindows:
		default:
			return
		}
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f dwarf\n", psess.Cputime())
	}

	newattr(&psess.dwtypes, dwarf.DW_AT_name, dwarf.DW_CLS_STRING, int64(len("dwtypes")), "dwtypes")

	newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_NULLTYPE, "<unspecified>", 0)

	newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_NULLTYPE, "void", 0)
	newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_BARE_PTRTYPE, "unsafe.Pointer", 0)

	die := newdie(ctxt, &psess.dwtypes, dwarf.DW_ABRV_BASETYPE, "uintptr", 0)
	newattr(die, dwarf.DW_AT_encoding, dwarf.DW_CLS_CONSTANT, dwarf.DW_ATE_unsigned, 0)
	newattr(die, dwarf.DW_AT_byte_size, dwarf.DW_CLS_CONSTANT, int64(ctxt.Arch.PtrSize), 0)
	newattr(die, dwarf.DW_AT_go_kind, dwarf.DW_CLS_CONSTANT, objabi.KindUintptr, 0)
	newattr(die, dwarf.DW_AT_go_runtime_type, dwarf.DW_CLS_ADDRESS, 0, psess.lookupOrDiag(ctxt, "type.uintptr"))
	psess.
		prototypedies = map[string]*dwarf.DWDie{
		"type.runtime.stringStructDWARF": nil,
		"type.runtime.slice":             nil,
		"type.runtime.hmap":              nil,
		"type.runtime.bmap":              nil,
		"type.runtime.sudog":             nil,
		"type.runtime.waitq":             nil,
		"type.runtime.hchan":             nil,
	}

	for _, typ := range []string{
		"type.runtime._type",
		"type.runtime.arraytype",
		"type.runtime.chantype",
		"type.runtime.functype",
		"type.runtime.maptype",
		"type.runtime.ptrtype",
		"type.runtime.slicetype",
		"type.runtime.structtype",
		"type.runtime.interfacetype",
		"type.runtime.itab",
		"type.runtime.imethod"} {
		psess.
			defgotype(ctxt, psess.lookupOrDiag(ctxt, typ))
	}
	psess.
		genasmsym(ctxt, psess.defdwsymb)

	abbrev := psess.writeabbrev(ctxt)
	syms := []*sym.Symbol{abbrev}

	units := psess.getCompilationUnits(ctxt)

	debugLine := ctxt.Syms.Lookup(".debug_line", 0)
	debugLine.Type = sym.SDWARFSECT
	debugRanges := ctxt.Syms.Lookup(".debug_ranges", 0)
	debugRanges.Type = sym.SDWARFRANGE
	debugRanges.Attr |= sym.AttrReachable
	syms = append(syms, debugLine)
	for _, u := range units {
		u.dwinfo, u.funcDIEs, u.absFnDIEs = psess.writelines(ctxt, u.lib, u.lib.Textp, debugLine)
		writepcranges(ctxt, u.dwinfo, u.lib.Textp[0], u.pcs, debugRanges)
	}
	psess.
		synthesizestringtypes(ctxt, psess.dwtypes.Child)
	psess.
		synthesizeslicetypes(ctxt, psess.dwtypes.Child)
	psess.
		synthesizemaptypes(ctxt, psess.dwtypes.Child)
	psess.
		synthesizechantypes(ctxt, psess.dwtypes.Child)
	psess.
		reversetree(&psess.dwroot.Child)
	psess.
		reversetree(&psess.dwtypes.Child)
	psess.
		reversetree(&psess.dwglobals.Child)
	psess.
		movetomodule(&psess.dwtypes)
	psess.
		movetomodule(&psess.dwglobals)

	infosyms := psess.writeinfo(ctxt, nil, units, abbrev)

	syms = psess.writeframes(ctxt, syms)
	syms = psess.writepub(ctxt, ".debug_pubnames", ispubname, syms)
	syms = psess.writepub(ctxt, ".debug_pubtypes", ispubtype, syms)
	syms = psess.writegdbscript(ctxt, syms)

	syms = append(syms, infosyms...)
	syms = collectlocs(ctxt, syms, units)
	syms = append(syms, debugRanges)
	syms = writeranges(ctxt, syms)
	psess.
		dwarfp = syms
}

func collectlocs(ctxt *Link, syms []*sym.Symbol, units []*compilationUnit) []*sym.Symbol {
	empty := true
	for _, u := range units {
		for _, fn := range u.funcDIEs {
			for _, reloc := range fn.R {
				if reloc.Type == objabi.R_DWARFSECREF && strings.HasPrefix(reloc.Sym.Name, dwarf.LocPrefix) {
					reloc.Sym.Attr |= sym.AttrReachable | sym.AttrNotInSymbolTable
					syms = append(syms, reloc.Sym)
					empty = false

					if ctxt.HeadType == objabi.Hdarwin {
						removeDwarfAddrListBaseAddress(ctxt, fn, reloc.Sym, true)
					}

					break
				}
			}
		}
	}

	if !empty {
		locsym := ctxt.Syms.Lookup(".debug_loc", 0)
		locsym.Type = sym.SDWARFLOC
		locsym.Attr |= sym.AttrReachable
		syms = append(syms, locsym)
	}
	return syms
}

// removeDwarfAddrListBaseAddress removes base address selector entries from
// DWARF location lists and range lists.
func removeDwarfAddrListBaseAddress(ctxt *Link, info, list *sym.Symbol, isloclist bool) {

	fn := list.R[0].Sym

	list.R = list.R[:0]

	relocate := func(addr uint64, offset int) {
		list.R = append(list.R, sym.Reloc{
			Off:  int32(offset),
			Siz:  uint8(ctxt.Arch.PtrSize),
			Type: objabi.R_ADDRCUOFF,
			Add:  int64(addr),
			Sym:  fn,
		})
	}

	for i := 0; i < len(list.P); {
		first := readPtr(ctxt, list.P[i:])
		second := readPtr(ctxt, list.P[i+ctxt.Arch.PtrSize:])

		if (first == 0 && second == 0) ||
			first == ^uint64(0) ||
			(ctxt.Arch.PtrSize == 4 && first == uint64(^uint32(0))) {

			i += ctxt.Arch.PtrSize * 2
			continue
		}

		relocate(first, i)
		relocate(second, i+ctxt.Arch.PtrSize)

		i += ctxt.Arch.PtrSize * 2
		if isloclist {
			i += 2 + int(ctxt.Arch.ByteOrder.Uint16(list.P[i:]))
		}
	}

	for i := range info.R {
		r := &info.R[i]
		if r.Sym != list {
			continue
		}
		r.Add += int64(2 * ctxt.Arch.PtrSize)
	}
}

// Read a pointer-sized uint from the beginning of buf.
func readPtr(ctxt *Link, buf []byte) uint64 {
	switch ctxt.Arch.PtrSize {
	case 4:
		return uint64(ctxt.Arch.ByteOrder.Uint32(buf))
	case 8:
		return ctxt.Arch.ByteOrder.Uint64(buf)
	default:
		panic("unexpected pointer size")
	}
}

/*
 *  Elf.
 */
func (psess *PackageSession) dwarfaddshstrings(ctxt *Link, shstrtab *sym.Symbol) {
	if *psess.FlagW {
		return
	}

	secs := []string{"abbrev", "frame", "info", "loc", "line", "pubnames", "pubtypes", "gdb_scripts", "ranges"}
	for _, sec := range secs {
		psess.
			Addstring(shstrtab, ".debug_"+sec)
		if ctxt.LinkMode == LinkExternal {
			psess.
				Addstring(shstrtab, psess.elfRelType+".debug_"+sec)
		} else {
			psess.
				Addstring(shstrtab, ".zdebug_"+sec)
		}
	}
}

// Add section symbols for DWARF debug info.  This is called before
// dwarfaddelfheaders.
func (psess *PackageSession) dwarfaddelfsectionsyms(ctxt *Link) {
	if *psess.FlagW {
		return
	}
	if ctxt.LinkMode != LinkExternal {
		return
	}

	s := ctxt.Syms.Lookup(".debug_info", 0)
	psess.
		putelfsectionsym(ctxt.Out, s, s.Sect.Elfsect.(*ElfShdr).shnum)
	s = ctxt.Syms.Lookup(".debug_abbrev", 0)
	psess.
		putelfsectionsym(ctxt.Out, s, s.Sect.Elfsect.(*ElfShdr).shnum)
	s = ctxt.Syms.Lookup(".debug_line", 0)
	psess.
		putelfsectionsym(ctxt.Out, s, s.Sect.Elfsect.(*ElfShdr).shnum)
	s = ctxt.Syms.Lookup(".debug_frame", 0)
	psess.
		putelfsectionsym(ctxt.Out, s, s.Sect.Elfsect.(*ElfShdr).shnum)
	s = ctxt.Syms.Lookup(".debug_loc", 0)
	if s.Sect != nil {
		psess.
			putelfsectionsym(ctxt.Out, s, s.Sect.Elfsect.(*ElfShdr).shnum)
	}
	s = ctxt.Syms.Lookup(".debug_ranges", 0)
	if s.Sect != nil {
		psess.
			putelfsectionsym(ctxt.Out, s, s.Sect.Elfsect.(*ElfShdr).shnum)
	}
}

// dwarfcompress compresses the DWARF sections. This must happen after
// relocations are applied. After this, dwarfp will contain a
// different (new) set of symbols, and sections may have been replaced.
func (psess *PackageSession) dwarfcompress(ctxt *Link) {
	if !(ctxt.IsELF || ctxt.HeadType == objabi.Hwindows) || ctxt.LinkMode == LinkExternal {
		return
	}

	var start int
	var newDwarfp []*sym.Symbol
	psess.
		Segdwarf.Sections = psess.Segdwarf.Sections[:0]
	for i, s := range psess.dwarfp {

		if i+1 >= len(psess.dwarfp) || s.Sect != psess.dwarfp[i+1].Sect {
			s1 := psess.compressSyms(ctxt, psess.dwarfp[start:i+1])
			if s1 == nil {

				newDwarfp = append(newDwarfp, psess.dwarfp[start:i+1]...)
				psess.
					Segdwarf.Sections = append(psess.Segdwarf.Sections, s.Sect)
			} else {
				compressedSegName := ".zdebug_" + s.Sect.Name[len(".debug_"):]
				sect := addsection(ctxt.Arch, &psess.Segdwarf, compressedSegName, 04)
				sect.Length = uint64(len(s1))
				newSym := ctxt.Syms.Lookup(compressedSegName, 0)
				newSym.P = s1
				newSym.Size = int64(len(s1))
				newSym.Sect = sect
				newDwarfp = append(newDwarfp, newSym)
			}
			start = i + 1
		}
	}
	psess.
		dwarfp = newDwarfp

	pos := psess.Segdwarf.Vaddr
	var prevSect *sym.Section
	for _, s := range psess.dwarfp {
		s.Value = int64(pos)
		if s.Sect != prevSect {
			s.Sect.Vaddr = uint64(s.Value)
			prevSect = s.Sect
		}
		if s.Sub != nil {
			log.Fatalf("%s: unexpected sub-symbols", s)
		}
		pos += uint64(s.Size)
		if ctxt.HeadType == objabi.Hwindows {
			pos = uint64(Rnd(int64(pos), psess.PEFILEALIGN))
		}

	}
	psess.
		Segdwarf.Length = pos - psess.Segdwarf.Vaddr
}

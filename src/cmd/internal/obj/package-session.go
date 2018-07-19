package obj

import (
	"github.com/dave/golib/src/cmd/internal/dwarf"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"github.com/dave/golib/src/cmd/internal/sys"
)

type PackageSession struct {
	dwarf  *dwarf.PackageSession
	objabi *objabi.PackageSession
	src    *src.PackageSession
	sys    *sys.PackageSession

	Anames          []string
	_AddrType_index [15]uint8

	aSpace      []opSet
	armCondCode []string

	opSuffixSpace []opSuffixSet

	regListSpace    []regListSet
	regSpace        []regSet
	textAttrStrings [14]struct {
		bit Attribute
		s   string
	}
}

func NewPackageSession(src_psess *src.PackageSession, dwarf_psess *dwarf.PackageSession, objabi_psess *objabi.PackageSession, sys_psess *sys.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.src = src_psess
	psess.dwarf = dwarf_psess
	psess.objabi = objabi_psess
	psess.sys = sys_psess
	psess._AddrType_index = [...]uint8{0, 9, 20, 33, 41, 51, 62, 73, 81, 90, 100, 111, 123, 133, 145}
	psess.armCondCode = []string{
		".EQ",
		".NE",
		".CS",
		".CC",
		".MI",
		".PL",
		".VS",
		".VC",
		".HI",
		".LS",
		".GE",
		".LT",
		".GT",
		".LE",
		"",
		".NV",
	}
	psess.Anames = []string{
		"XXX",
		"CALL",
		"DUFFCOPY",
		"DUFFZERO",
		"END",
		"FUNCDATA",
		"JMP",
		"NOP",
		"PCDATA",
		"RET",
		"GETCALLERPC",
		"TEXT",
		"UNDEF",
	}
	psess.textAttrStrings = [...]struct {
		bit Attribute
		s   string
	}{
		{bit: AttrDuplicateOK, s: "DUPOK"},
		{bit: AttrMakeTypelink, s: ""},
		{bit: AttrCFunc, s: "CFUNC"},
		{bit: AttrNoSplit, s: "NOSPLIT"},
		{bit: AttrLeaf, s: "LEAF"},
		{bit: AttrSeenGlobl, s: ""},
		{bit: AttrOnList, s: ""},
		{bit: AttrReflectMethod, s: "REFLECTMETHOD"},
		{bit: AttrLocal, s: "LOCAL"},
		{bit: AttrWrapper, s: "WRAPPER"},
		{bit: AttrNeedCtxt, s: "NEEDCTXT"},
		{bit: AttrNoFrame, s: "NOFRAME"},
		{bit: AttrStatic, s: "STATIC"},
		{bit: AttrWasInlined, s: ""},
	}
	return psess
}

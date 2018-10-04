package obj

type PackageState struct {
	dwarf           *dwarf.PackageState
	objabi          *objabi.PackageState
	src             *src.PackageState
	sys             *sys.PackageState
	Anames          []string
	_AddrType_index [15]uint8
	aSpace          []opSet
	armCondCode     []string
	opSuffixSpace   []opSuffixSet
	regListSpace    []regListSet
	regSpace        []regSet
	textAttrStrings [14]struct {
		bit Attribute
		s   string
	}
}

func NewPackageState(objabi_pstate *objabi.PackageState, dwarf_pstate *dwarf.PackageState, src_pstate *src.PackageState, sys_pstate *sys.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.objabi = objabi_pstate
	pstate.dwarf = dwarf_pstate
	pstate.src = src_pstate
	pstate.sys = sys_pstate
	pstate.textAttrStrings = [...]struct {
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
	pstate.armCondCode = []string{
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
	pstate.Anames = []string{
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
	pstate._AddrType_index = [...]uint8{0, 9, 20, 33, 41, 51, 62, 73, 81, 90, 100, 111, 123, 133, 145}
	return pstate
}

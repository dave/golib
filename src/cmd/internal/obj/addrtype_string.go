package obj

import "strconv"

const _AddrType_name = "TYPE_NONETYPE_BRANCHTYPE_TEXTSIZETYPE_MEMTYPE_CONSTTYPE_FCONSTTYPE_SCONSTTYPE_REGTYPE_ADDRTYPE_SHIFTTYPE_REGREGTYPE_REGREG2TYPE_INDIRTYPE_REGLIST"

func (i AddrType) String(psess *PackageSession,) string {
	if i >= AddrType(len(psess._AddrType_index)-1) {
		return "AddrType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AddrType_name[psess._AddrType_index[i]:psess._AddrType_index[i+1]]
}

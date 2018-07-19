package types

import "strconv"

const _EType_name = "xxxINT8UINT8INT16UINT16INT32UINT32INT64UINT64INTUINTUINTPTRCOMPLEX64COMPLEX128FLOAT32FLOAT64BOOLPTR32PTR64FUNCSLICEARRAYSTRUCTCHANMAPINTERFORWANYSTRINGUNSAFEPTRIDEALNILBLANKFUNCARGSCHANARGSDDDFIELDSSATUPLENTYPE"

func (i EType) String(psess *PackageSession,) string {
	if i >= EType(len(psess._EType_index)-1) {
		return "EType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EType_name[psess._EType_index[i]:psess._EType_index[i+1]]
}

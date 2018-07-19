package gc

import "strconv"

const _Class_name = "PxxxPEXTERNPAUTOPAUTOHEAPPPARAMPPARAMOUTPFUNCPDISCARD"

func (i Class) String(psess *PackageSession,) string {
	if i >= Class(len(psess._Class_index)-1) {
		return "Class(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Class_name[psess._Class_index[i]:psess._Class_index[i+1]]
}

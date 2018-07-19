package syntax

import "strconv"

const _Operator_name = ":!<-||&&==!=<<=>>=+-|^*/%&&^<<>>"

func (i Operator) String(psess *PackageSession,) string {
	i -= 1
	if i >= Operator(len(psess._Operator_index)-1) {
		return "Operator(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Operator_name[psess._Operator_index[i]:psess._Operator_index[i+1]]
}

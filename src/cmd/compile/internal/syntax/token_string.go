package syntax

import "strconv"

const _token_name = "EOFnameliteralopop=opop=:=<-*([{)]},;:....breakcasechanconstcontinuedefaultdeferelsefallthroughforfuncgogotoifimportinterfacemappackagerangereturnselectstructswitchtypevar"

func (i token) String(psess *PackageSession,) string {
	i -= 1
	if i >= token(len(psess._token_index)-1) {
		return "token(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _token_name[psess._token_index[i]:psess._token_index[i+1]]
}

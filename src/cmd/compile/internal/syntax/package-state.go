package syntax

type PackageState struct {
	ImplicitOne     *BasicLit
	_Operator_index [23]uint8
	_token_index    [48]uint8
	blankByte       []byte
	indentBytes     []byte
	invalid         *LabeledStmt
	keywordMap      [64]token
	newlineByte     []byte
	tabBytes        []byte
}

func NewPackageState() *PackageState {
	pstate := &PackageState{}
	pstate.indentBytes = []byte(".  ")
	pstate._Operator_index = [...]uint8{0, 1, 2, 4, 6, 8, 10, 12, 13, 15, 16, 18, 19, 20, 21, 22, 23, 24, 25, 26, 28, 30, 32}
	pstate.invalid = new(LabeledStmt)
	pstate.ImplicitOne = &BasicLit{Value: "1"}
	pstate.tabBytes = []byte("\t\t\t\t\t\t\t\t")
	pstate.newlineByte = []byte("\n")
	pstate.blankByte = []byte(" ")
	pstate._token_index = [...]uint8{0, 3, 7, 14, 16, 19, 23, 24, 26, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 42, 47, 51, 55, 60, 68, 75, 80, 84, 95, 98, 102, 104, 108, 110, 116, 125, 128, 135, 140, 146, 152, 158, 164, 168, 171, 171}
	return pstate
}

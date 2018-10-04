package src

type PackageState struct {
	NoPos  Pos
	NoXPos XPos
	noPos  Pos
}

func NewPackageState() *PackageState { pstate := &PackageState{}; return pstate }

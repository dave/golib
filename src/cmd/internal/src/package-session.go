package src

type PackageSession struct {
	NoPos  Pos
	NoXPos XPos

	noPos Pos
}

func NewPackageSession() *PackageSession {
	psess := &PackageSession{}
	return psess
}

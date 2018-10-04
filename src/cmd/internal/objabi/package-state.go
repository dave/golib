package objabi

type PackageState struct {
	Clobberdead_enabled      int
	DebugCPU_enabled         int
	Fieldtrack_enabled       int
	GO386                    string
	GOARCH                   string
	GOARM                    int
	GOMIPS                   string
	GOMIPS64                 string
	GOOS                     string
	GOROOT                   string
	Preemptibleloops_enabled int
	Version                  string
	_RelocType_index         [46]uint16
	_SymKind_index           [13]uint8
	buildID                  string
	defaultExpstring         string
	defaultGOROOT            string
	exper                    []struct {
		name string
		val  *int
	}
	framepointer_enabled int
}

func NewPackageState() *PackageState {
	pstate := &PackageState{}
	pstate._RelocType_index = [...]uint16{0, 6, 17, 28, 38, 47, 60, 66, 72, 81, 92, 101, 112, 122, 129, 136, 144, 152, 160, 166, 172, 178, 188, 197, 208, 219, 229, 238, 251, 265, 279, 293, 309, 323, 337, 348, 362, 377, 394, 412, 433, 443, 454, 467, 478, 490}
	pstate.GOROOT = envOr("GOROOT", pstate.defaultGOROOT)
	pstate.framepointer_enabled = 1
	pstate.exper = []struct {
		name string
		val  *int
	}{
		{"fieldtrack", &pstate.Fieldtrack_enabled},
		{"framepointer", &pstate.framepointer_enabled},
		{"preemptibleloops", &pstate.Preemptibleloops_enabled},
		{"clobberdead", &pstate.Clobberdead_enabled},
		{"debugcpu", &pstate.DebugCPU_enabled},
	}
	pstate.defaultExpstring = pstate.Expstring()
	pstate.GO386 = envOr("GO386", defaultGO386)
	pstate.GOARM = goarm()
	pstate.GOMIPS = gomips()
	pstate.GOMIPS64 = gomips64()
	pstate.GOOS = envOr("GOOS", defaultGOOS)
	pstate.GOARCH = envOr("GOARCH", defaultGOARCH)
	pstate.Version = version
	pstate._SymKind_index = [...]uint8{0, 4, 9, 16, 26, 31, 35, 44, 51, 61, 72, 81, 91}
	return pstate
}

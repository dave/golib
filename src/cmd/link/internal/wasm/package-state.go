package wasm

type PackageState struct {
	ld            *ld.PackageState
	objabi        *objabi.PackageState
	sym           *sym.PackageState
	sys           *sys.PackageState
	nameRegexp    *regexp.Regexp
	wasmFuncTypes map[string]*wasmFuncType
}

func NewPackageState(sys_pstate *sys.PackageState, ld_pstate *ld.PackageState, objabi_pstate *objabi.PackageState, sym_pstate *sym.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.sys = sys_pstate
	pstate.ld = ld_pstate
	pstate.objabi = objabi_pstate
	pstate.sym = sym_pstate
	pstate.wasmFuncTypes = map[string]*wasmFuncType{
		"_rt0_wasm_js":           &wasmFuncType{Params: []byte{I32, I32}},                                 // argc, argv
		"runtime.wasmMove":       &wasmFuncType{Params: []byte{I32, I32, I32}},                            // dst, src, len
		"runtime.wasmZero":       &wasmFuncType{Params: []byte{I32, I32}},                                 // ptr, len
		"runtime.wasmDiv":        &wasmFuncType{Params: []byte{I64, I64}, Results: []byte{I64}},           // x, y -> x/y
		"runtime.wasmTruncS":     &wasmFuncType{Params: []byte{F64}, Results: []byte{I64}},                // x -> int(x)
		"runtime.wasmTruncU":     &wasmFuncType{Params: []byte{F64}, Results: []byte{I64}},                // x -> uint(x)
		"runtime.gcWriteBarrier": &wasmFuncType{Params: []byte{I64, I64}},                                 // ptr, val
		"cmpbody":                &wasmFuncType{Params: []byte{I64, I64, I64, I64}, Results: []byte{I64}}, // a, alen, b, blen -> -1/0/1
		"memeqbody":              &wasmFuncType{Params: []byte{I64, I64, I64}, Results: []byte{I64}},      // a, b, len -> 0/1
		"memcmp":                 &wasmFuncType{Params: []byte{I32, I32, I32}, Results: []byte{I32}},      // a, b, len -> <0/0/>0
		"memchr":                 &wasmFuncType{Params: []byte{I32, I32, I32}, Results: []byte{I32}},      // s, c, len -> index
	}
	pstate.nameRegexp = regexp.MustCompile("[^\\w\\.]")
	return pstate
}

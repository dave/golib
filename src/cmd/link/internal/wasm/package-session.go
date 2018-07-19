package wasm

import (
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"regexp"
)

type PackageSession struct {
	ld     *ld.PackageSession
	objabi *objabi.PackageSession
	sym    *sym.PackageSession
	sys    *sys.PackageSession

	nameRegexp    *regexp.Regexp
	wasmFuncTypes map[string]*wasmFuncType
}

func NewPackageSession(objabi_psess *objabi.PackageSession, ld_psess *ld.PackageSession, sym_psess *sym.PackageSession, sys_psess *sys.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.objabi = objabi_psess
	psess.ld = ld_psess
	psess.sym = sym_psess
	psess.sys = sys_psess
	psess.wasmFuncTypes = map[string]*wasmFuncType{
		"_rt0_wasm_js":           &wasmFuncType{Params: []byte{I32, I32}},
		"runtime.wasmMove":       &wasmFuncType{Params: []byte{I32, I32, I32}},
		"runtime.wasmZero":       &wasmFuncType{Params: []byte{I32, I32}},
		"runtime.wasmDiv":        &wasmFuncType{Params: []byte{I64, I64}, Results: []byte{I64}},
		"runtime.wasmTruncS":     &wasmFuncType{Params: []byte{F64}, Results: []byte{I64}},
		"runtime.wasmTruncU":     &wasmFuncType{Params: []byte{F64}, Results: []byte{I64}},
		"runtime.gcWriteBarrier": &wasmFuncType{Params: []byte{I64, I64}},
		"cmpbody":                &wasmFuncType{Params: []byte{I64, I64, I64, I64}, Results: []byte{I64}},
		"memeqbody":              &wasmFuncType{Params: []byte{I64, I64, I64}, Results: []byte{I64}},
		"memcmp":                 &wasmFuncType{Params: []byte{I32, I32, I32}, Results: []byte{I32}},
		"memchr":                 &wasmFuncType{Params: []byte{I32, I32, I32}, Results: []byte{I32}},
	}
	psess.nameRegexp = regexp.MustCompile("[^\\w\\.]")
	return psess
}

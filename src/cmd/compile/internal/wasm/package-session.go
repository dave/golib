package wasm

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/wasm"
)

type PackageSession struct {
	gc    *gc.PackageSession
	obj   *obj.PackageSession
	ssa   *ssa.PackageSession
	types *types.PackageSession
	wasm  *wasm.PackageSession
}

func NewPackageSession(gc_psess *gc.PackageSession, ssa_psess *ssa.PackageSession, types_psess *types.PackageSession, obj_psess *obj.PackageSession, wasm_psess *wasm.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.gc = gc_psess
	psess.ssa = ssa_psess
	psess.types = types_psess
	psess.obj = obj_psess
	psess.wasm = wasm_psess
	return psess
}

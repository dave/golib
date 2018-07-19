package s390x

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/s390x"
)

type PackageSession struct {
	gc    *gc.PackageSession
	obj   *obj.PackageSession
	s390x *s390x.PackageSession
	ssa   *ssa.PackageSession
	types *types.PackageSession

	blockJump [89]struct {
		asm    obj.As
		invasm obj.As
	}
}

func NewPackageSession(gc_psess *gc.PackageSession, s390x_psess *s390x.PackageSession, obj_psess *obj.PackageSession, ssa_psess *ssa.PackageSession, types_psess *types.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.gc = gc_psess
	psess.s390x = s390x_psess
	psess.obj = obj_psess
	psess.ssa = ssa_psess
	psess.types = types_psess
	psess.blockJump = [...]struct {
		asm, invasm obj.As
	}{
		ssa.BlockS390XEQ:  {s390x.ABEQ, s390x.ABNE},
		ssa.BlockS390XNE:  {s390x.ABNE, s390x.ABEQ},
		ssa.BlockS390XLT:  {s390x.ABLT, s390x.ABGE},
		ssa.BlockS390XGE:  {s390x.ABGE, s390x.ABLT},
		ssa.BlockS390XLE:  {s390x.ABLE, s390x.ABGT},
		ssa.BlockS390XGT:  {s390x.ABGT, s390x.ABLE},
		ssa.BlockS390XGTF: {s390x.ABGT, s390x.ABLEU},
		ssa.BlockS390XGEF: {s390x.ABGE, s390x.ABLTU},
	}
	return psess
}

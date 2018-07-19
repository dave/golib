package x86

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/x86"
	"github.com/dave/golib/src/cmd/internal/objabi"
)

type PackageSession struct {
	gc     *gc.PackageSession
	obj    *obj.PackageSession
	objabi *objabi.PackageSession
	ssa    *ssa.PackageSession
	types  *types.PackageSession
	x86    *x86.PackageSession

	blockJump [15]struct {
		asm    obj.As
		invasm obj.As
	}

	eqfJumps [2][2]gc.FloatingEQNEJump

	nefJumps [2][2]gc.FloatingEQNEJump
}

func NewPackageSession(gc_psess *gc.PackageSession, ssa_psess *ssa.PackageSession, types_psess *types.PackageSession, obj_psess *obj.PackageSession, x86_psess *x86.PackageSession, objabi_psess *objabi.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.gc = gc_psess
	psess.ssa = ssa_psess
	psess.types = types_psess
	psess.obj = obj_psess
	psess.x86 = x86_psess
	psess.objabi = objabi_psess
	psess.blockJump = [...]struct {
		asm, invasm obj.As
	}{
		ssa.Block386EQ:  {x86.AJEQ, x86.AJNE},
		ssa.Block386NE:  {x86.AJNE, x86.AJEQ},
		ssa.Block386LT:  {x86.AJLT, x86.AJGE},
		ssa.Block386GE:  {x86.AJGE, x86.AJLT},
		ssa.Block386LE:  {x86.AJLE, x86.AJGT},
		ssa.Block386GT:  {x86.AJGT, x86.AJLE},
		ssa.Block386ULT: {x86.AJCS, x86.AJCC},
		ssa.Block386UGE: {x86.AJCC, x86.AJCS},
		ssa.Block386UGT: {x86.AJHI, x86.AJLS},
		ssa.Block386ULE: {x86.AJLS, x86.AJHI},
		ssa.Block386ORD: {x86.AJPC, x86.AJPS},
		ssa.Block386NAN: {x86.AJPS, x86.AJPC},
	}
	psess.eqfJumps = [2][2]gc.FloatingEQNEJump{
		{{Jump: x86.AJNE, Index: 1}, {Jump: x86.AJPS, Index: 1}},
		{{Jump: x86.AJNE, Index: 1}, {Jump: x86.AJPC, Index: 0}},
	}
	psess.nefJumps = [2][2]gc.FloatingEQNEJump{
		{{Jump: x86.AJNE, Index: 0}, {Jump: x86.AJPC, Index: 1}},
		{{Jump: x86.AJNE, Index: 0}, {Jump: x86.AJPS, Index: 0}},
	}
	return psess
}

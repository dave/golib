package amd64

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

	blockJump [29]struct {
		asm    obj.As
		invasm obj.As
	}

	eqfJumps [2][2]gc.FloatingEQNEJump
	isPlan9  bool
	leaptr   obj.As

	nefJumps [2][2]gc.FloatingEQNEJump
}

func NewPackageSession(gc_psess *gc.PackageSession, x86_psess *x86.PackageSession, objabi_psess *objabi.PackageSession, obj_psess *obj.PackageSession, ssa_psess *ssa.PackageSession, types_psess *types.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.gc = gc_psess
	psess.x86 = x86_psess
	psess.objabi = objabi_psess
	psess.obj = obj_psess
	psess.ssa = ssa_psess
	psess.types = types_psess
	psess.leaptr = x86.ALEAQ
	psess.isPlan9 = psess.objabi.
		GOOS == "plan9"
	psess.blockJump = [...]struct {
		asm, invasm obj.As
	}{
		ssa.BlockAMD64EQ:  {x86.AJEQ, x86.AJNE},
		ssa.BlockAMD64NE:  {x86.AJNE, x86.AJEQ},
		ssa.BlockAMD64LT:  {x86.AJLT, x86.AJGE},
		ssa.BlockAMD64GE:  {x86.AJGE, x86.AJLT},
		ssa.BlockAMD64LE:  {x86.AJLE, x86.AJGT},
		ssa.BlockAMD64GT:  {x86.AJGT, x86.AJLE},
		ssa.BlockAMD64ULT: {x86.AJCS, x86.AJCC},
		ssa.BlockAMD64UGE: {x86.AJCC, x86.AJCS},
		ssa.BlockAMD64UGT: {x86.AJHI, x86.AJLS},
		ssa.BlockAMD64ULE: {x86.AJLS, x86.AJHI},
		ssa.BlockAMD64ORD: {x86.AJPC, x86.AJPS},
		ssa.BlockAMD64NAN: {x86.AJPS, x86.AJPC},
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

package arm

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/arm"
	"github.com/dave/golib/src/cmd/internal/objabi"
)

type PackageSession struct {
	arm    *arm.PackageSession
	gc     *gc.PackageSession
	obj    *obj.PackageSession
	objabi *objabi.PackageSession
	ssa    *ssa.PackageSession
	types  *types.PackageSession

	blockJump map[ssa.BlockKind]struct {
		asm    obj.As
		invasm obj.As
	}
	condBits map[ssa.Op]uint8
}

func NewPackageSession(gc_psess *gc.PackageSession, obj_psess *obj.PackageSession, arm_psess *arm.PackageSession, ssa_psess *ssa.PackageSession, types_psess *types.PackageSession, objabi_psess *objabi.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.gc = gc_psess
	psess.obj = obj_psess
	psess.arm = arm_psess
	psess.ssa = ssa_psess
	psess.types = types_psess
	psess.objabi = objabi_psess
	psess.condBits = map[ssa.Op]uint8{
		ssa.OpARMEqual:         arm.C_SCOND_EQ,
		ssa.OpARMNotEqual:      arm.C_SCOND_NE,
		ssa.OpARMLessThan:      arm.C_SCOND_LT,
		ssa.OpARMLessThanU:     arm.C_SCOND_LO,
		ssa.OpARMLessEqual:     arm.C_SCOND_LE,
		ssa.OpARMLessEqualU:    arm.C_SCOND_LS,
		ssa.OpARMGreaterThan:   arm.C_SCOND_GT,
		ssa.OpARMGreaterThanU:  arm.C_SCOND_HI,
		ssa.OpARMGreaterEqual:  arm.C_SCOND_GE,
		ssa.OpARMGreaterEqualU: arm.C_SCOND_HS,
	}
	psess.blockJump = map[ssa.BlockKind]struct {
		asm, invasm obj.As
	}{
		ssa.BlockARMEQ:  {arm.ABEQ, arm.ABNE},
		ssa.BlockARMNE:  {arm.ABNE, arm.ABEQ},
		ssa.BlockARMLT:  {arm.ABLT, arm.ABGE},
		ssa.BlockARMGE:  {arm.ABGE, arm.ABLT},
		ssa.BlockARMLE:  {arm.ABLE, arm.ABGT},
		ssa.BlockARMGT:  {arm.ABGT, arm.ABLE},
		ssa.BlockARMULT: {arm.ABLO, arm.ABHS},
		ssa.BlockARMUGE: {arm.ABHS, arm.ABLO},
		ssa.BlockARMUGT: {arm.ABHI, arm.ABLS},
		ssa.BlockARMULE: {arm.ABLS, arm.ABHI},
	}
	return psess
}

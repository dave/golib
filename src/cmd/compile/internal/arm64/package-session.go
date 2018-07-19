package arm64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/arm64"
	"github.com/dave/golib/src/cmd/internal/objabi"
)

type PackageSession struct {
	arm64  *arm64.PackageSession
	gc     *gc.PackageSession
	obj    *obj.PackageSession
	objabi *objabi.PackageSession
	ssa    *ssa.PackageSession
	types  *types.PackageSession

	blockJump map[ssa.BlockKind]struct {
		asm    obj.As
		invasm obj.As
	}
	condBits map[ssa.Op]int16
	darwin   bool
}

func NewPackageSession(gc_psess *gc.PackageSession, obj_psess *obj.PackageSession, arm64_psess *arm64.PackageSession, objabi_psess *objabi.PackageSession, ssa_psess *ssa.PackageSession, types_psess *types.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.gc = gc_psess
	psess.obj = obj_psess
	psess.arm64 = arm64_psess
	psess.objabi = objabi_psess
	psess.ssa = ssa_psess
	psess.types = types_psess
	psess.darwin = psess.objabi.GOOS == "darwin"
	psess.condBits = map[ssa.Op]int16{
		ssa.OpARM64Equal:         arm64.COND_EQ,
		ssa.OpARM64NotEqual:      arm64.COND_NE,
		ssa.OpARM64LessThan:      arm64.COND_LT,
		ssa.OpARM64LessThanU:     arm64.COND_LO,
		ssa.OpARM64LessEqual:     arm64.COND_LE,
		ssa.OpARM64LessEqualU:    arm64.COND_LS,
		ssa.OpARM64GreaterThan:   arm64.COND_GT,
		ssa.OpARM64GreaterThanU:  arm64.COND_HI,
		ssa.OpARM64GreaterEqual:  arm64.COND_GE,
		ssa.OpARM64GreaterEqualU: arm64.COND_HS,
	}
	psess.blockJump = map[ssa.BlockKind]struct {
		asm, invasm obj.As
	}{
		ssa.BlockARM64EQ:   {arm64.ABEQ, arm64.ABNE},
		ssa.BlockARM64NE:   {arm64.ABNE, arm64.ABEQ},
		ssa.BlockARM64LT:   {arm64.ABLT, arm64.ABGE},
		ssa.BlockARM64GE:   {arm64.ABGE, arm64.ABLT},
		ssa.BlockARM64LE:   {arm64.ABLE, arm64.ABGT},
		ssa.BlockARM64GT:   {arm64.ABGT, arm64.ABLE},
		ssa.BlockARM64ULT:  {arm64.ABLO, arm64.ABHS},
		ssa.BlockARM64UGE:  {arm64.ABHS, arm64.ABLO},
		ssa.BlockARM64UGT:  {arm64.ABHI, arm64.ABLS},
		ssa.BlockARM64ULE:  {arm64.ABLS, arm64.ABHI},
		ssa.BlockARM64Z:    {arm64.ACBZ, arm64.ACBNZ},
		ssa.BlockARM64NZ:   {arm64.ACBNZ, arm64.ACBZ},
		ssa.BlockARM64ZW:   {arm64.ACBZW, arm64.ACBNZW},
		ssa.BlockARM64NZW:  {arm64.ACBNZW, arm64.ACBZW},
		ssa.BlockARM64TBZ:  {arm64.ATBZ, arm64.ATBNZ},
		ssa.BlockARM64TBNZ: {arm64.ATBNZ, arm64.ATBZ},
	}
	return psess
}

package ppc64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/ppc64"
	"github.com/dave/golib/src/cmd/internal/objabi"
)

type PackageSession struct {
	gc     *gc.PackageSession
	obj    *obj.PackageSession
	objabi *objabi.PackageSession
	ppc64  *ppc64.PackageSession
	ssa    *ssa.PackageSession
	types  *types.PackageSession

	blockJump [81]struct {
		asm      obj.As
		invasm   obj.As
		asmeq    bool
		invasmun bool
	}
	iselOps  map[ssa.Op]iselOp
	iselRegs [2]int16
}

func NewPackageSession(gc_psess *gc.PackageSession, obj_psess *obj.PackageSession, ppc64_psess *ppc64.PackageSession, ssa_psess *ssa.PackageSession, types_psess *types.PackageSession, objabi_psess *objabi.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.gc = gc_psess
	psess.obj = obj_psess
	psess.ppc64 = ppc64_psess
	psess.ssa = ssa_psess
	psess.types = types_psess
	psess.objabi = objabi_psess
	psess.iselRegs = [2]int16{ppc64.REG_R0, ppc64.REGTMP}
	psess.iselOps = map[ssa.Op]iselOp{
		ssa.OpPPC64Equal:         iselOp{cond: ppc64.C_COND_EQ, valueIfCond: 1},
		ssa.OpPPC64NotEqual:      iselOp{cond: ppc64.C_COND_EQ, valueIfCond: 0},
		ssa.OpPPC64LessThan:      iselOp{cond: ppc64.C_COND_LT, valueIfCond: 1},
		ssa.OpPPC64GreaterEqual:  iselOp{cond: ppc64.C_COND_LT, valueIfCond: 0},
		ssa.OpPPC64GreaterThan:   iselOp{cond: ppc64.C_COND_GT, valueIfCond: 1},
		ssa.OpPPC64LessEqual:     iselOp{cond: ppc64.C_COND_GT, valueIfCond: 0},
		ssa.OpPPC64FLessThan:     iselOp{cond: ppc64.C_COND_LT, valueIfCond: 1},
		ssa.OpPPC64FGreaterThan:  iselOp{cond: ppc64.C_COND_GT, valueIfCond: 1},
		ssa.OpPPC64FLessEqual:    iselOp{cond: ppc64.C_COND_LT, valueIfCond: 1},
		ssa.OpPPC64FGreaterEqual: iselOp{cond: ppc64.C_COND_GT, valueIfCond: 1},
	}
	psess.blockJump = [...]struct {
		asm, invasm     obj.As
		asmeq, invasmun bool
	}{
		ssa.BlockPPC64EQ: {ppc64.ABEQ, ppc64.ABNE, false, false},
		ssa.BlockPPC64NE: {ppc64.ABNE, ppc64.ABEQ, false, false},

		ssa.BlockPPC64LT: {ppc64.ABLT, ppc64.ABGE, false, false},
		ssa.BlockPPC64GE: {ppc64.ABGE, ppc64.ABLT, false, false},
		ssa.BlockPPC64LE: {ppc64.ABLE, ppc64.ABGT, false, false},
		ssa.BlockPPC64GT: {ppc64.ABGT, ppc64.ABLE, false, false},

		ssa.BlockPPC64FLT: {ppc64.ABLT, ppc64.ABGE, false, false},
		ssa.BlockPPC64FGE: {ppc64.ABGT, ppc64.ABLT, true, true},
		ssa.BlockPPC64FLE: {ppc64.ABLT, ppc64.ABGT, true, true},
		ssa.BlockPPC64FGT: {ppc64.ABGT, ppc64.ABLE, false, false},
	}
	return psess
}

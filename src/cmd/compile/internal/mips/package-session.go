package mips

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/mips"
	"github.com/dave/golib/src/cmd/internal/objabi"
)

type PackageSession struct {
	gc     *gc.PackageSession
	mips   *mips.PackageSession
	obj    *obj.PackageSession
	objabi *objabi.PackageSession
	ssa    *ssa.PackageSession
	types  *types.PackageSession

	blockJump map[ssa.BlockKind]struct {
		asm    obj.As
		invasm obj.As
	}
}

func NewPackageSession(gc_psess *gc.PackageSession, ssa_psess *ssa.PackageSession, mips_psess *mips.PackageSession, objabi_psess *objabi.PackageSession, obj_psess *obj.PackageSession, types_psess *types.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.gc = gc_psess
	psess.ssa = ssa_psess
	psess.mips = mips_psess
	psess.objabi = objabi_psess
	psess.obj = obj_psess
	psess.types = types_psess
	psess.blockJump = map[ssa.BlockKind]struct {
		asm, invasm obj.As
	}{
		ssa.BlockMIPSEQ:  {mips.ABEQ, mips.ABNE},
		ssa.BlockMIPSNE:  {mips.ABNE, mips.ABEQ},
		ssa.BlockMIPSLTZ: {mips.ABLTZ, mips.ABGEZ},
		ssa.BlockMIPSGEZ: {mips.ABGEZ, mips.ABLTZ},
		ssa.BlockMIPSLEZ: {mips.ABLEZ, mips.ABGTZ},
		ssa.BlockMIPSGTZ: {mips.ABGTZ, mips.ABLEZ},
		ssa.BlockMIPSFPT: {mips.ABFPT, mips.ABFPF},
		ssa.BlockMIPSFPF: {mips.ABFPF, mips.ABFPT},
	}
	return psess
}

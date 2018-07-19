package mips64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/internal/obj/mips"
)

func (psess *PackageSession) Init(arch *gc.Arch) {
	arch.LinkArch = &psess.mips.Linkmips64
	if psess.objabi.GOARCH == "mips64le" {
		arch.LinkArch = &psess.mips.Linkmips64le
	}
	arch.REGSP = mips.REGSP
	arch.MAXWIDTH = 1 << 50
	arch.SoftFloat = psess.objabi.GOMIPS64 == "softfloat"
	arch.ZeroRange = psess.zerorange
	arch.ZeroAuto = psess.zeroAuto
	arch.Ginsnop = psess.ginsnop

	arch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	arch.SSAGenValue = psess.ssaGenValue
	arch.SSAGenBlock = psess.ssaGenBlock
}

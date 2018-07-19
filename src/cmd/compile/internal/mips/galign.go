package mips

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/internal/obj/mips"
)

func (psess *PackageSession) Init(arch *gc.Arch) {
	arch.LinkArch = &psess.mips.Linkmips
	if psess.objabi.GOARCH == "mipsle" {
		arch.LinkArch = &psess.mips.Linkmipsle
	}
	arch.REGSP = mips.REGSP
	arch.MAXWIDTH = (1 << 31) - 1
	arch.SoftFloat = (psess.objabi.GOMIPS == "softfloat")
	arch.ZeroRange = psess.zerorange
	arch.ZeroAuto = psess.zeroAuto
	arch.Ginsnop = psess.ginsnop
	arch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	arch.SSAGenValue = psess.ssaGenValue
	arch.SSAGenBlock = psess.ssaGenBlock
}

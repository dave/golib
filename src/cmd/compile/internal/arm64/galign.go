package arm64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/internal/obj/arm64"
)

func (psess *PackageSession) Init(arch *gc.Arch) {
	arch.LinkArch = &psess.arm64.Linkarm64
	arch.REGSP = arm64.REGSP
	arch.MAXWIDTH = 1 << 50

	arch.PadFrame = padframe
	arch.ZeroRange = psess.zerorange
	arch.ZeroAuto = psess.zeroAuto
	arch.Ginsnop = psess.ginsnop

	arch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	arch.SSAGenValue = psess.ssaGenValue
	arch.SSAGenBlock = psess.ssaGenBlock
}

package s390x

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj/s390x"
)

func (psess *PackageSession) Init(arch *gc.Arch) {
	arch.LinkArch = &psess.s390x.Links390x
	arch.REGSP = s390x.REGSP
	arch.MAXWIDTH = 1 << 50

	arch.ZeroRange = psess.zerorange
	arch.ZeroAuto = psess.zeroAuto
	arch.Ginsnop = psess.ginsnop

	arch.SSAMarkMoves = psess.ssaMarkMoves
	arch.SSAGenValue = psess.ssaGenValue
	arch.SSAGenBlock = psess.ssaGenBlock
}

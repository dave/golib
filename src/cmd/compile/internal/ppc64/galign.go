package ppc64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj/ppc64"
)

func (psess *PackageSession) Init(arch *gc.Arch) {
	arch.LinkArch = &psess.ppc64.Linkppc64
	if psess.objabi.GOARCH == "ppc64le" {
		arch.LinkArch = &psess.ppc64.Linkppc64le
	}
	arch.REGSP = ppc64.REGSP
	arch.MAXWIDTH = 1 << 50

	arch.ZeroRange = psess.zerorange
	arch.ZeroAuto = psess.zeroAuto
	arch.Ginsnop = psess.ginsnop2

	arch.SSAMarkMoves = ssaMarkMoves
	arch.SSAGenValue = psess.ssaGenValue
	arch.SSAGenBlock = psess.ssaGenBlock
}

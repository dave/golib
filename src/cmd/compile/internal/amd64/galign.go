package amd64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj/x86"
)

func (psess *PackageSession) Init(arch *gc.Arch) {
	arch.LinkArch = &psess.x86.Linkamd64
	if psess.objabi.GOARCH == "amd64p32" {
		arch.LinkArch = &psess.x86.Linkamd64p32
		psess.
			leaptr = x86.ALEAL
	}
	arch.REGSP = x86.REGSP
	arch.MAXWIDTH = 1 << 50

	arch.ZeroRange = psess.zerorange
	arch.ZeroAuto = psess.zeroAuto
	arch.Ginsnop = psess.ginsnop

	arch.SSAMarkMoves = psess.ssaMarkMoves
	arch.SSAGenValue = psess.ssaGenValue
	arch.SSAGenBlock = psess.ssaGenBlock
}

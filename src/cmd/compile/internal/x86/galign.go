package x86

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj/x86"

	"os"
)

func (psess *PackageSession) Init(arch *gc.Arch) {
	arch.LinkArch = &psess.x86.Link386
	arch.REGSP = x86.REGSP
	switch v := psess.objabi.GO386; v {
	case "387":
		arch.Use387 = true
		arch.SSAGenValue = psess.ssaGenValue387
		arch.SSAGenBlock = psess.ssaGenBlock387
	case "sse2":
		arch.SSAGenValue = psess.ssaGenValue
		arch.SSAGenBlock = psess.ssaGenBlock
	default:
		fmt.Fprintf(os.Stderr, "unsupported setting GO386=%s\n", v)
		psess.gc.
			Exit(1)
	}
	arch.MAXWIDTH = (1 << 32) - 1

	arch.ZeroRange = psess.zerorange
	arch.ZeroAuto = psess.zeroAuto
	arch.Ginsnop = psess.ginsnop

	arch.SSAMarkMoves = psess.ssaMarkMoves
}

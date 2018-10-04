// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x86

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj/x86"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"os"
)

func (pstate *PackageState) Init(arch *gc.Arch) {
	arch.LinkArch = &pstate.x86.Link386
	arch.REGSP = x86.REGSP
	switch v := pstate.objabi.GO386; v {
	case "387":
		arch.Use387 = true
		arch.SSAGenValue = pstate.ssaGenValue387
		arch.SSAGenBlock = pstate.ssaGenBlock387
	case "sse2":
		arch.SSAGenValue = pstate.ssaGenValue
		arch.SSAGenBlock = pstate.ssaGenBlock
	default:
		fmt.Fprintf(os.Stderr, "unsupported setting GO386=%s\n", v)
		pstate.gc.Exit(1)
	}
	arch.MAXWIDTH = (1 << 32) - 1

	arch.ZeroRange = pstate.zerorange
	arch.ZeroAuto = pstate.zeroAuto
	arch.Ginsnop = pstate.ginsnop

	arch.SSAMarkMoves = pstate.ssaMarkMoves
}

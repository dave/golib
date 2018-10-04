// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppc64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj/ppc64"
	"github.com/dave/golib/src/cmd/internal/objabi"
)

func (pstate *PackageState) Init(arch *gc.Arch) {
	arch.LinkArch = &pstate.ppc64.Linkppc64
	if pstate.objabi.GOARCH == "ppc64le" {
		arch.LinkArch = &pstate.ppc64.Linkppc64le
	}
	arch.REGSP = ppc64.REGSP
	arch.MAXWIDTH = 1 << 50

	arch.ZeroRange = pstate.zerorange
	arch.ZeroAuto = pstate.zeroAuto
	arch.Ginsnop = pstate.ginsnop2

	arch.SSAMarkMoves = ssaMarkMoves
	arch.SSAGenValue = pstate.ssaGenValue
	arch.SSAGenBlock = pstate.ssaGenBlock
}

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package amd64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj/x86"
	"github.com/dave/golib/src/cmd/internal/objabi"
)

func (pstate *PackageState) Init(arch *gc.Arch) {
	arch.LinkArch = &pstate.x86.Linkamd64
	if pstate.objabi.GOARCH == "amd64p32" {
		arch.LinkArch = &pstate.x86.Linkamd64p32
		pstate.leaptr = x86.ALEAL
	}
	arch.REGSP = x86.REGSP
	arch.MAXWIDTH = 1 << 50

	arch.ZeroRange = pstate.zerorange
	arch.ZeroAuto = pstate.zeroAuto
	arch.Ginsnop = pstate.ginsnop

	arch.SSAMarkMoves = pstate.ssaMarkMoves
	arch.SSAGenValue = pstate.ssaGenValue
	arch.SSAGenBlock = pstate.ssaGenBlock
}

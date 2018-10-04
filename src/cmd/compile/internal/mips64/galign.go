// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mips64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/internal/obj/mips"
	"github.com/dave/golib/src/cmd/internal/objabi"
)

func (pstate *PackageState) Init(arch *gc.Arch) {
	arch.LinkArch = &pstate.mips.Linkmips64
	if pstate.objabi.GOARCH == "mips64le" {
		arch.LinkArch = &pstate.mips.Linkmips64le
	}
	arch.REGSP = mips.REGSP
	arch.MAXWIDTH = 1 << 50
	arch.SoftFloat = pstate.objabi.GOMIPS64 == "softfloat"
	arch.ZeroRange = pstate.zerorange
	arch.ZeroAuto = pstate.zeroAuto
	arch.Ginsnop = pstate.ginsnop

	arch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	arch.SSAGenValue = pstate.ssaGenValue
	arch.SSAGenBlock = pstate.ssaGenBlock
}

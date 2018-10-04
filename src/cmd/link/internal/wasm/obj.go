// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wasm

import (
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
)

func (pstate *PackageState) Init() (*sys.Arch, ld.Arch) {
	theArch := ld.Arch{
		Funcalign: 16,
		Maxalign:  32,
		Minalign:  1,

		Archinit:      pstate.archinit,
		AssignAddress: assignAddress,
		Asmb:          pstate.asmb,
		Gentext:       gentext,
	}

	return pstate.sys.ArchWasm, theArch
}

func (pstate *PackageState) archinit(ctxt *ld.Link) {
	if *pstate.ld.FlagRound == -1 {
		*pstate.ld.FlagRound = 4096
	}
	if *pstate.ld.FlagTextAddr == -1 {
		*pstate.ld.FlagTextAddr = 0
	}
}

package wasm

import (
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
)

func (psess *PackageSession) Init() (*sys.Arch, ld.Arch) {
	theArch := ld.Arch{
		Funcalign: 16,
		Maxalign:  32,
		Minalign:  1,

		Archinit:      psess.archinit,
		AssignAddress: assignAddress,
		Asmb:          psess.asmb,
		Gentext:       gentext,
	}

	return psess.sys.ArchWasm, theArch
}

func (psess *PackageSession) archinit(ctxt *ld.Link) {
	if *psess.ld.FlagRound == -1 {
		*psess.ld.FlagRound = 4096
	}
	if *psess.ld.FlagTextAddr == -1 {
		*psess.ld.FlagTextAddr = 0
	}
}

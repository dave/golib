package ppc64

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
)

func (psess *PackageSession) Init() (*sys.Arch, ld.Arch) {
	arch := psess.sys.ArchPPC64
	if psess.objabi.GOARCH == "ppc64le" {
		arch = psess.sys.ArchPPC64LE
	}

	theArch := ld.Arch{
		Funcalign:  funcAlign,
		Maxalign:   maxAlign,
		Minalign:   minAlign,
		Dwarfregsp: dwarfRegSP,
		Dwarfreglr: dwarfRegLR,

		Adddynrel:        psess.adddynrel,
		Archinit:         psess.archinit,
		Archreloc:        psess.archreloc,
		Archrelocvariant: psess.archrelocvariant,
		Asmb:             psess.asmb,
		Elfreloc1:        elfreloc1,
		Elfsetupplt:      elfsetupplt,
		Gentext:          psess.gentext,
		Trampoline:       psess.trampoline,
		Machoreloc1:      machoreloc1,

		Linuxdynld: "/lib64/ld64.so.1",

		Freebsddynld:   "XXX",
		Openbsddynld:   "XXX",
		Netbsddynld:    "XXX",
		Dragonflydynld: "XXX",
		Solarisdynld:   "XXX",
	}

	return arch, theArch
}

func (psess *PackageSession) archinit(ctxt *ld.Link) {
	switch ctxt.HeadType {
	default:
		psess.ld.
			Exitf("unknown -H option: %v", ctxt.HeadType)

	case objabi.Hplan9:
		psess.ld.
			HEADR = 32

		if *psess.ld.FlagTextAddr == -1 {
			*psess.ld.FlagTextAddr = 4128
		}
		if *psess.ld.FlagDataAddr == -1 {
			*psess.ld.FlagDataAddr = 0
		}
		if *psess.ld.FlagRound == -1 {
			*psess.ld.FlagRound = 4096
		}

	case objabi.Hlinux:
		if ctxt.Arch == psess.sys.ArchPPC64 {
			*psess.ld.FlagD = true
		}
		psess.ld.
			Elfinit(ctxt)
		psess.ld.
			HEADR = ld.ELFRESERVE
		if *psess.ld.FlagTextAddr == -1 {
			*psess.ld.FlagTextAddr = 0x10000 + int64(psess.ld.HEADR)
		}
		if *psess.ld.FlagDataAddr == -1 {
			*psess.ld.FlagDataAddr = 0
		}
		if *psess.ld.FlagRound == -1 {
			*psess.ld.FlagRound = 0x10000
		}

	case objabi.Hnacl:
		psess.ld.
			Elfinit(ctxt)
		psess.ld.
			HEADR = 0x10000
		psess.ld.
			Funcalign = 16
		if *psess.ld.FlagTextAddr == -1 {
			*psess.ld.FlagTextAddr = 0x20000
		}
		if *psess.ld.FlagDataAddr == -1 {
			*psess.ld.FlagDataAddr = 0
		}
		if *psess.ld.FlagRound == -1 {
			*psess.ld.FlagRound = 0x10000
		}
	}

	if *psess.ld.FlagDataAddr != 0 && *psess.ld.FlagRound != 0 {
		fmt.Printf("warning: -D0x%x is ignored because of -R0x%x\n", uint64(*psess.ld.FlagDataAddr), uint32(*psess.ld.FlagRound))
	}
}

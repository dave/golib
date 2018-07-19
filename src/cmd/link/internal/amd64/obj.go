package amd64

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
)

func (psess *PackageSession) Init() (*sys.Arch, ld.Arch) {
	arch := psess.sys.ArchAMD64
	if psess.objabi.GOARCH == "amd64p32" {
		arch = psess.sys.ArchAMD64P32
	}

	theArch := ld.Arch{
		Funcalign:  funcAlign,
		Maxalign:   maxAlign,
		Minalign:   minAlign,
		Dwarfregsp: dwarfRegSP,
		Dwarfreglr: dwarfRegLR,

		Adddynrel:        psess.adddynrel,
		Archinit:         psess.archinit,
		Archreloc:        archreloc,
		Archrelocvariant: archrelocvariant,
		Asmb:             psess.asmb,
		Elfreloc1:        elfreloc1,
		Elfsetupplt:      elfsetupplt,
		Gentext:          gentext,
		Machoreloc1:      psess.machoreloc1,
		PEreloc1:         psess.pereloc1,
		TLSIEtoLE:        tlsIEtoLE,

		Linuxdynld:     "/lib64/ld-linux-x86-64.so.2",
		Freebsddynld:   "/libexec/ld-elf.so.1",
		Openbsddynld:   "/usr/libexec/ld.so",
		Netbsddynld:    "/libexec/ld.elf_so",
		Dragonflydynld: "/usr/libexec/ld-elf.so.2",
		Solarisdynld:   "/lib/amd64/ld.so.1",
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
			HEADR = 32 + 8

		if *psess.ld.FlagTextAddr == -1 {
			*psess.ld.FlagTextAddr = 0x200000 + int64(psess.ld.HEADR)
		}
		if *psess.ld.FlagDataAddr == -1 {
			*psess.ld.FlagDataAddr = 0
		}
		if *psess.ld.FlagRound == -1 {
			*psess.ld.FlagRound = 0x200000
		}

	case objabi.Hdarwin:
		psess.ld.
			HEADR = ld.INITIAL_MACHO_HEADR
		if *psess.ld.FlagRound == -1 {
			*psess.ld.FlagRound = 4096
		}
		if *psess.ld.FlagTextAddr == -1 {
			*psess.ld.FlagTextAddr = 0x1000000 + int64(psess.ld.HEADR)
		}
		if *psess.ld.FlagDataAddr == -1 {
			*psess.ld.FlagDataAddr = 0
		}

	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd,
		objabi.Hdragonfly,
		objabi.Hsolaris:
		psess.ld.
			Elfinit(ctxt)
		psess.ld.
			HEADR = ld.ELFRESERVE
		if *psess.ld.FlagTextAddr == -1 {
			*psess.ld.FlagTextAddr = (1 << 22) + int64(psess.ld.HEADR)
		}
		if *psess.ld.FlagDataAddr == -1 {
			*psess.ld.FlagDataAddr = 0
		}
		if *psess.ld.FlagRound == -1 {
			*psess.ld.FlagRound = 4096
		}

	case objabi.Hnacl:
		psess.ld.
			Elfinit(ctxt)
		*psess.ld.FlagW = true
		psess.ld.
			HEADR = 0x10000
		psess.ld.
			Funcalign = 32
		if *psess.ld.FlagTextAddr == -1 {
			*psess.ld.FlagTextAddr = 0x20000
		}
		if *psess.ld.FlagDataAddr == -1 {
			*psess.ld.FlagDataAddr = 0
		}
		if *psess.ld.FlagRound == -1 {
			*psess.ld.FlagRound = 0x10000
		}

	case objabi.Hwindows:

		return
	}

	if *psess.ld.FlagDataAddr != 0 && *psess.ld.FlagRound != 0 {
		fmt.Printf("warning: -D0x%x is ignored because of -R0x%x\n", uint64(*psess.ld.FlagDataAddr), uint32(*psess.ld.FlagRound))
	}
}

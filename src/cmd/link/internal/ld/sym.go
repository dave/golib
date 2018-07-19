package ld

import (
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"log"
)

func (psess *PackageSession) linknew(arch *sys.Arch) *Link {
	ctxt := &Link{
		Syms:         sym.NewSymbols(),
		Out:          &OutBuf{arch: arch},
		Arch:         arch,
		LibraryByPkg: make(map[string]*sym.Library),
	}

	if psess.objabi.GOARCH != arch.Name {
		log.Fatalf("invalid objabi.GOARCH %s (want %s)", psess.objabi.GOARCH, arch.Name)
	}
	psess.
		AtExit(func() {
			if psess.nerrors > 0 && ctxt.Out.f != nil {
				ctxt.Out.f.Close()
				psess.
					mayberemoveoutfile()
			}
		})

	return ctxt
}

// computeTLSOffset records the thread-local storage offset.
func (ctxt *Link) computeTLSOffset(psess *PackageSession) {
	switch ctxt.HeadType {
	default:
		log.Fatalf("unknown thread-local storage offset for %v", ctxt.HeadType)

	case objabi.Hplan9, objabi.Hwindows, objabi.Hjs:
		break

	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd,
		objabi.Hdragonfly,
		objabi.Hsolaris:
		if psess.objabi.GOOS == "android" {
			switch ctxt.Arch.Family {
			case sys.AMD64:

				ctxt.Tlsoffset = 0x1d0
			case sys.I386:

				ctxt.Tlsoffset = 0xf8
			default:
				ctxt.Tlsoffset = -1 * ctxt.Arch.PtrSize
			}
		} else {
			ctxt.Tlsoffset = -1 * ctxt.Arch.PtrSize
		}

	case objabi.Hnacl:
		switch ctxt.Arch.Family {
		default:
			log.Fatalf("unknown thread-local storage offset for nacl/%s", ctxt.Arch.Name)

		case sys.ARM:
			ctxt.Tlsoffset = 0

		case sys.AMD64:
			ctxt.Tlsoffset = 0

		case sys.I386:
			ctxt.Tlsoffset = -8
		}

	case objabi.Hdarwin:
		switch ctxt.Arch.Family {
		default:
			log.Fatalf("unknown thread-local storage offset for darwin/%s", ctxt.Arch.Name)

		case sys.I386:
			ctxt.Tlsoffset = 0x18

		case sys.AMD64:
			ctxt.Tlsoffset = 0x30

		case sys.ARM:
			ctxt.Tlsoffset = 0

		case sys.ARM64:
			ctxt.Tlsoffset = 0
		}
	}

}

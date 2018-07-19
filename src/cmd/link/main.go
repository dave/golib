package main

import (
	"fmt"

	"github.com/dave/golib/src/cmd/internal/sys"

	"github.com/dave/golib/src/cmd/link/internal/ld"

	"os"
)

func (psess *PackageSession) main() {
	var arch *sys.Arch
	var theArch ld.Arch

	switch psess.objabi.GOARCH {
	default:
		fmt.Fprintf(os.Stderr, "link: unknown architecture %q\n", psess.objabi.GOARCH)
		os.Exit(2)
	case "386":
		arch, theArch = psess.x86.Init()
	case "amd64", "amd64p32":
		arch, theArch = psess.amd64.Init()
	case "arm":
		arch, theArch = psess.arm.Init()
	case "arm64":
		arch, theArch = psess.arm64.Init()
	case "mips", "mipsle":
		arch, theArch = psess.mips.Init()
	case "mips64", "mips64le":
		arch, theArch = psess.mips64.Init()
	case "ppc64", "ppc64le":
		arch, theArch = psess.ppc64.Init()
	case "s390x":
		arch, theArch = psess.s390x.Init()
	case "wasm":
		arch, theArch = psess.wasm.Init()
	}
	psess.ld.
		Main(arch, theArch)
}

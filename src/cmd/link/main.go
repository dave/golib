// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/amd64"
	"github.com/dave/golib/src/cmd/link/internal/arm"
	"github.com/dave/golib/src/cmd/link/internal/arm64"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/mips"
	"github.com/dave/golib/src/cmd/link/internal/mips64"
	"github.com/dave/golib/src/cmd/link/internal/ppc64"
	"github.com/dave/golib/src/cmd/link/internal/s390x"
	"github.com/dave/golib/src/cmd/link/internal/wasm"
	"github.com/dave/golib/src/cmd/link/internal/x86"
	"os"
)

// The bulk of the linker implementation lives in cmd/link/internal/ld.
// Architecture-specific code lives in cmd/link/internal/GOARCH.
//
// Program initialization:
//
// Before any argument parsing is done, the Init function of relevant
// architecture package is called. The only job done in Init is
// configuration of the architecture-specific variables.
//
// Then control flow passes to ld.Main, which parses flags, makes
// some configuration decisions, and then gives the architecture
// packages a second chance to modify the linker's configuration
// via the ld.Arch.Archinit function.

func (pstate *PackageState) main() {
	var arch *sys.Arch
	var theArch ld.Arch

	switch pstate.objabi.GOARCH {
	default:
		fmt.Fprintf(os.Stderr, "link: unknown architecture %q\n", pstate.objabi.GOARCH)
		os.Exit(2)
	case "386":
		arch, theArch = pstate.x86.Init()
	case "amd64", "amd64p32":
		arch, theArch = pstate.amd64.Init()
	case "arm":
		arch, theArch = pstate.arm.Init()
	case "arm64":
		arch, theArch = pstate.arm64.Init()
	case "mips", "mipsle":
		arch, theArch = pstate.mips.Init()
	case "mips64", "mips64le":
		arch, theArch = pstate.mips64.Init()
	case "ppc64", "ppc64le":
		arch, theArch = pstate.ppc64.Init()
	case "s390x":
		arch, theArch = pstate.s390x.Init()
	case "wasm":
		arch, theArch = pstate.wasm.Init()
	}
	pstate.ld.Main(arch, theArch)
}

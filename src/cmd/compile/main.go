// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/amd64"
	"github.com/dave/golib/src/cmd/compile/internal/arm"
	"github.com/dave/golib/src/cmd/compile/internal/arm64"
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/mips"
	"github.com/dave/golib/src/cmd/compile/internal/mips64"
	"github.com/dave/golib/src/cmd/compile/internal/ppc64"
	"github.com/dave/golib/src/cmd/compile/internal/s390x"
	"github.com/dave/golib/src/cmd/compile/internal/wasm"
	"github.com/dave/golib/src/cmd/compile/internal/x86"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"log"
	"os"
)

func (pstate *PackageState) main() {
	// disable timestamps for reproducible output
	log.SetFlags(0)
	log.SetPrefix("compile: ")

	archInit, ok := pstate.archInits[pstate.objabi.GOARCH]
	if !ok {
		fmt.Fprintf(os.Stderr, "compile: unknown architecture %q\n", pstate.objabi.GOARCH)
		os.Exit(2)
	}

	pstate.gc.Main(archInit)
	pstate.gc.Exit(0)
}

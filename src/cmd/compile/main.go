package main

import (
	"fmt"

	"log"
	"os"
)

func (psess *PackageSession) main() {

	log.SetFlags(0)
	log.SetPrefix("compile: ")

	archInit, ok := psess.archInits[psess.objabi.GOARCH]
	if !ok {
		fmt.Fprintf(os.Stderr, "compile: unknown architecture %q\n", psess.objabi.GOARCH)
		os.Exit(2)
	}
	psess.gc.
		Main(archInit)
	psess.gc.
		Exit(0)
}

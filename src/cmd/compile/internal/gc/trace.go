// +build go1.7

package gc

import (
	"os"
	tracepkg "runtime/trace"
)

func (psess *PackageSession) init() {
	psess.
		traceHandler = psess.traceHandlerGo17
}

func (psess *PackageSession) traceHandlerGo17(traceprofile string) {
	f, err := os.Create(traceprofile)
	if err != nil {
		psess.
			Fatalf("%v", err)
	}
	if err := tracepkg.Start(f); err != nil {
		psess.
			Fatalf("%v", err)
	}
	psess.
		atExit(tracepkg.Stop)
}

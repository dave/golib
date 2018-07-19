package gc

import (
	"os"
	"runtime"
	"runtime/pprof"
)

// Line returns n's position as a string. If n has been inlined,
// it uses the outermost position where n has been inlined.
func (n *Node) Line(psess *PackageSession) string {
	return psess.linestr(n.Pos)
}

func (psess *PackageSession) atExit(f func()) {
	psess.
		atExitFuncs = append(psess.atExitFuncs, f)
}

func (psess *PackageSession) Exit(code int) {
	for i := len(psess.atExitFuncs) - 1; i >= 0; i-- {
		f := psess.atExitFuncs[i]
		psess.
			atExitFuncs = psess.atExitFuncs[:i]
		f()
	}
	os.Exit(code)
}

func (psess *PackageSession) startProfile() {
	if psess.cpuprofile != "" {
		f, err := os.Create(psess.cpuprofile)
		if err != nil {
			psess.
				Fatalf("%v", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			psess.
				Fatalf("%v", err)
		}
		psess.
			atExit(pprof.StopCPUProfile)
	}
	if psess.memprofile != "" {
		if psess.memprofilerate != 0 {
			runtime.MemProfileRate = int(psess.memprofilerate)
		}
		f, err := os.Create(psess.memprofile)
		if err != nil {
			psess.
				Fatalf("%v", err)
		}
		psess.
			atExit(func() {

				runtime.GC()
				// compilebench parses the memory profile to extract memstats,
				// which are only written in the legacy pprof format.
				// See golang.org/issue/18641 and runtime/pprof/pprof.go:writeHeap.
				const writeLegacyFormat = 1
				if err := pprof.Lookup("heap").WriteTo(f, writeLegacyFormat); err != nil {
					psess.
						Fatalf("%v", err)
				}
			})
	} else {

		runtime.MemProfileRate = 0
	}
	if psess.blockprofile != "" {
		f, err := os.Create(psess.blockprofile)
		if err != nil {
			psess.
				Fatalf("%v", err)
		}
		runtime.SetBlockProfileRate(1)
		psess.
			atExit(func() {
				pprof.Lookup("block").WriteTo(f, 0)
				f.Close()
			})
	}
	if psess.mutexprofile != "" {
		f, err := os.Create(psess.mutexprofile)
		if err != nil {
			psess.
				Fatalf("%v", err)
		}
		startMutexProfiling()
		psess.
			atExit(func() {
				pprof.Lookup("mutex").WriteTo(f, 0)
				f.Close()
			})
	}
	if psess.traceprofile != "" && psess.traceHandler != nil {
		psess.
			traceHandler(psess.traceprofile)
	}
}

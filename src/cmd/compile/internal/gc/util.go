// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"os"
	"runtime"
	"runtime/pprof"
)

// Line returns n's position as a string. If n has been inlined,
// it uses the outermost position where n has been inlined.
func (n *Node) Line(pstate *PackageState) string {
	return pstate.linestr(n.Pos)
}

func (pstate *PackageState) atExit(f func()) {
	pstate.atExitFuncs = append(pstate.atExitFuncs, f)
}

func (pstate *PackageState) Exit(code int) {
	for i := len(pstate.atExitFuncs) - 1; i >= 0; i-- {
		f := pstate.atExitFuncs[i]
		pstate.atExitFuncs = pstate.atExitFuncs[:i]
		f()
	}
	os.Exit(code)
}

func (pstate *PackageState) startProfile() {
	if pstate.cpuprofile != "" {
		f, err := os.Create(pstate.cpuprofile)
		if err != nil {
			pstate.Fatalf("%v", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			pstate.Fatalf("%v", err)
		}
		pstate.atExit(pprof.StopCPUProfile)
	}
	if pstate.memprofile != "" {
		if pstate.memprofilerate != 0 {
			runtime.MemProfileRate = int(pstate.memprofilerate)
		}
		f, err := os.Create(pstate.memprofile)
		if err != nil {
			pstate.Fatalf("%v", err)
		}
		pstate.atExit(func() {
			// Profile all outstanding allocations.
			runtime.GC()
			// compilebench parses the memory profile to extract memstats,
			// which are only written in the legacy pprof format.
			// See golang.org/issue/18641 and runtime/pprof/pprof.go:writeHeap.
			const writeLegacyFormat = 1
			if err := pprof.Lookup("heap").WriteTo(f, writeLegacyFormat); err != nil {
				pstate.Fatalf("%v", err)
			}
		})
	} else {
		// Not doing memory profiling; disable it entirely.
		runtime.MemProfileRate = 0
	}
	if pstate.blockprofile != "" {
		f, err := os.Create(pstate.blockprofile)
		if err != nil {
			pstate.Fatalf("%v", err)
		}
		runtime.SetBlockProfileRate(1)
		pstate.atExit(func() {
			pprof.Lookup("block").WriteTo(f, 0)
			f.Close()
		})
	}
	if pstate.mutexprofile != "" {
		f, err := os.Create(pstate.mutexprofile)
		if err != nil {
			pstate.Fatalf("%v", err)
		}
		startMutexProfiling()
		pstate.atExit(func() {
			pprof.Lookup("mutex").WriteTo(f, 0)
			f.Close()
		})
	}
	if pstate.traceprofile != "" && pstate.traceHandler != nil {
		pstate.traceHandler(pstate.traceprofile)
	}
}

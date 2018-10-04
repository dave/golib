// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// Compile is the main entry point for this package.
// Compile modifies f so that on return:
//   · all Values in f map to 0 or 1 assembly instructions of the target architecture
//   · the order of f.Blocks is the order to emit the Blocks
//   · the order of b.Values is the order to emit the Values in each Block
//   · f has a non-nil regAlloc field
func (pstate *PackageState) Compile(f *Func) {
	// TODO: debugging - set flags to control verbosity of compiler,
	// which phases to dump IR before/after, etc.
	if f.Log() {
		f.Logf("compiling %s\n", f.Name)
	}

	// hook to print function & phase if panic happens
	phaseName := "init"
	defer func() {
		if phaseName != "" {
			err := recover()
			stack := make([]byte, 16384)
			n := runtime.Stack(stack, false)
			stack = stack[:n]
			f.Fatalf("panic during %s while compiling %s:\n\n%v\n\n%s\n", phaseName, f.Name, err, stack)
		}
	}()

	// Run all the passes
	printFunc(f)
	f.HTMLWriter.WriteFunc(pstate, "start", "start", f)
	if pstate.BuildDump != "" && pstate.BuildDump == f.Name {
		f.dumpFile(pstate, "build")
	}
	if pstate.checkEnabled {
		pstate.checkFunc(f)
	}
	const logMemStats = false
	for _, p := range pstate.passes {
		if !f.Config.optimize && !p.required || p.disabled {
			continue
		}
		f.pass = &p
		phaseName = p.name
		if f.Log() {
			f.Logf("  pass %s begin\n", p.name)
		}
		// TODO: capture logging during this pass, add it to the HTML
		var mStart runtime.MemStats
		if logMemStats || p.mem {
			runtime.ReadMemStats(&mStart)
		}

		tStart := time.Now()
		p.fn(f)
		tEnd := time.Now()

		// Need something less crude than "Log the whole intermediate result".
		if f.Log() || f.HTMLWriter != nil {
			time := tEnd.Sub(tStart).Nanoseconds()
			var stats string
			if logMemStats {
				var mEnd runtime.MemStats
				runtime.ReadMemStats(&mEnd)
				nBytes := mEnd.TotalAlloc - mStart.TotalAlloc
				nAllocs := mEnd.Mallocs - mStart.Mallocs
				stats = fmt.Sprintf("[%d ns %d allocs %d bytes]", time, nAllocs, nBytes)
			} else {
				stats = fmt.Sprintf("[%d ns]", time)
			}

			f.Logf("  pass %s end %s\n", p.name, stats)
			printFunc(f)
			f.HTMLWriter.WriteFunc(pstate, phaseName, fmt.Sprintf("%s <span class=\"stats\">%s</span>", phaseName, stats), f)
		}
		if p.time || p.mem {
			// Surround timing information w/ enough context to allow comparisons.
			time := tEnd.Sub(tStart).Nanoseconds()
			if p.time {
				f.LogStat("TIME(ns)", time)
			}
			if p.mem {
				var mEnd runtime.MemStats
				runtime.ReadMemStats(&mEnd)
				nBytes := mEnd.TotalAlloc - mStart.TotalAlloc
				nAllocs := mEnd.Mallocs - mStart.Mallocs
				f.LogStat("TIME(ns):BYTES:ALLOCS", time, nBytes, nAllocs)
			}
		}
		if p.dump != nil && p.dump[f.Name] {
			// Dump function to appropriately named file
			f.dumpFile(pstate, phaseName)
		}
		if pstate.checkEnabled {
			pstate.checkFunc(f)
		}
	}

	// Squash error printing defer
	phaseName = ""
}

// dumpFile creates a file from the phase name and function name
// Dumping is done to files to avoid buffering huge strings before
// output.
func (f *Func) dumpFile(pstate *PackageState, phaseName string) {
	pstate.dumpFileSeq++
	fname := fmt.Sprintf("%s_%02d__%s.dump", f.Name, pstate.dumpFileSeq, phaseName)
	fname = strings.Replace(fname, " ", "_", -1)
	fname = strings.Replace(fname, "/", "_", -1)
	fname = strings.Replace(fname, ":", "_", -1)

	fi, err := os.Create(fname)
	if err != nil {
		f.Warnl(pstate.src.NoXPos, "Unable to create after-phase dump file %s", fname)
		return
	}

	p := stringFuncPrinter{w: fi}
	pstate.fprintFunc(p, f)
	fi.Close()
}

type pass struct {
	name     string
	fn       func(*Func)
	required bool
	disabled bool
	time     bool            // report time to run pass
	mem      bool            // report mem stats to run pass
	stats    int             // pass reports own "stats" (e.g., branches removed)
	debug    int             // pass performs some debugging. =1 should be in error-testing-friendly Warnl format.
	test     int             // pass-specific ad-hoc option, perhaps useful in development
	dump     map[string]bool // dump if function name matches
}

func (p *pass) addDump(s string) {
	if p.dump == nil {
		p.dump = make(map[string]bool)
	}
	p.dump[s] = true
}

// PhaseOption sets the specified flag in the specified ssa phase,
// returning empty string if this was successful or a string explaining
// the error if it was not.
// A version of the phase name with "_" replaced by " " is also checked for a match.
// If the phase name begins a '~' then the rest of the underscores-replaced-with-blanks
// version is used as a regular expression to match the phase name(s).
//
// Special cases that have turned out to be useful:
//  ssa/check/on enables checking after each phase
//  ssa/all/time enables time reporting for all phases
//
// See gc/lex.go for dissection of the option string.
// Example uses:
//
// GO_GCFLAGS=-d=ssa/generic_cse/time,ssa/generic_cse/stats,ssa/generic_cse/debug=3 ./make.bash
//
// BOOT_GO_GCFLAGS=-d='ssa/~^.*scc$/off' GO_GCFLAGS='-d=ssa/~^.*scc$/off' ./make.bash
//
func (pstate *PackageState) PhaseOption(phase, flag string, val int, valString string) string {
	if phase == "help" {
		lastcr := 0
		phasenames := "    check, all, build, intrinsics"
		for _, p := range pstate.passes {
			pn := strings.Replace(p.name, " ", "_", -1)
			if len(pn)+len(phasenames)-lastcr > 70 {
				phasenames += "\n    "
				lastcr = len(phasenames)
				phasenames += pn
			} else {
				phasenames += ", " + pn
			}
		}
		return "PhaseOptions usage:\n\n    go tool compile -d=ssa/<phase>/<flag>[=<value>|<function_name>]\n\nwhere:\n\n- <phase> is one of:\n" + phasenames + "\n\n- <flag> is one of:\n    on, off, debug, mem, time, test, stats, dump\n\n- <value> defaults to 1\n\n- <function_name> is required for the \"dump\" flag, and specifies the\n  name of function to dump after <phase>\n\nPhase \"all\" supports flags \"time\", \"mem\", and \"dump\".\nPhase \"intrinsics\" supports flags \"on\", \"off\", and \"debug\".\n\nIf the \"dump\" flag is specified, the output is written on a file named\n<phase>__<function_name>_<seq>.dump; otherwise it is directed to stdout.\n\nExamples:\n\n    -d=ssa/check/on\nenables checking after each phase\n\n    -d=ssa/all/time\nenables time reporting for all phases\n\n    -d=ssa/prove/debug=2\nsets debugging level to 2 in the prove pass\n\nMultiple flags can be passed at once, by separating them with\ncommas. For example:\n\n    -d=ssa/check/on,ssa/all/time\n"
	}

	if phase == "check" && flag == "on" {
		pstate.checkEnabled = val != 0
		return ""
	}
	if phase == "check" && flag == "off" {
		pstate.checkEnabled = val == 0
		return ""
	}

	alltime := false
	allmem := false
	alldump := false
	if phase == "all" {
		if flag == "time" {
			alltime = val != 0
		} else if flag == "mem" {
			allmem = val != 0
		} else if flag == "dump" {
			alldump = val != 0
			if alldump {
				pstate.BuildDump = valString
			}
		} else {
			return fmt.Sprintf("Did not find a flag matching %s in -d=ssa/%s debug option", flag, phase)
		}
	}

	if phase == "intrinsics" {
		switch flag {
		case "on":
			pstate.IntrinsicsDisable = val == 0
		case "off":
			pstate.IntrinsicsDisable = val != 0
		case "debug":
			pstate.IntrinsicsDebug = val
		default:
			return fmt.Sprintf("Did not find a flag matching %s in -d=ssa/%s debug option", flag, phase)
		}
		return ""
	}
	if phase == "build" {
		switch flag {
		case "debug":
			pstate.BuildDebug = val
		case "test":
			pstate.BuildTest = val
		case "stats":
			pstate.BuildStats = val
		case "dump":
			pstate.BuildDump = valString
		default:
			return fmt.Sprintf("Did not find a flag matching %s in -d=ssa/%s debug option", flag, phase)
		}
		return ""
	}

	underphase := strings.Replace(phase, "_", " ", -1)
	var re *regexp.Regexp
	if phase[0] == '~' {
		r, ok := regexp.Compile(underphase[1:])
		if ok != nil {
			return fmt.Sprintf("Error %s in regexp for phase %s, flag %s", ok.Error(), phase, flag)
		}
		re = r
	}
	matchedOne := false
	for i, p := range pstate.passes {
		if phase == "all" {
			p.time = alltime
			p.mem = allmem
			if alldump {
				p.addDump(valString)
			}
			pstate.passes[i] = p
			matchedOne = true
		} else if p.name == phase || p.name == underphase || re != nil && re.MatchString(p.name) {
			switch flag {
			case "on":
				p.disabled = val == 0
			case "off":
				p.disabled = val != 0
			case "time":
				p.time = val != 0
			case "mem":
				p.mem = val != 0
			case "debug":
				p.debug = val
			case "stats":
				p.stats = val
			case "test":
				p.test = val
			case "dump":
				p.addDump(valString)
			default:
				return fmt.Sprintf("Did not find a flag matching %s in -d=ssa/%s debug option", flag, phase)
			}
			if p.disabled && p.required {
				return fmt.Sprintf("Cannot disable required SSA phase %s using -d=ssa/%s debug option", phase, phase)
			}
			pstate.passes[i] = p
			matchedOne = true
		}
	}
	if matchedOne {
		return ""
	}
	return fmt.Sprintf("Did not find a phase matching %s in -d=ssa/... debug option", phase)
}

// Double-check phase ordering constraints.
// This code is intended to document the ordering requirements
// between different phases. It does not override the passes
// list above.
type constraint struct {
	a, b string // a must come before b
}

func (pstate *PackageState) init() {
	for _, c := range pstate.passOrder {
		a, b := c.a, c.b
		i := -1
		j := -1
		for k, p := range pstate.passes {
			if p.name == a {
				i = k
			}
			if p.name == b {
				j = k
			}
		}
		if i < 0 {
			log.Panicf("pass %s not found", a)
		}
		if j < 0 {
			log.Panicf("pass %s not found", b)
		}
		if i >= j {
			log.Panicf("passes %s and %s out of order", a, b)
		}
	}
}

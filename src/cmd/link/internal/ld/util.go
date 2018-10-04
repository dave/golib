// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"os"
	"time"
)

// TODO(josharian): delete. See issue 19865.
func (pstate *PackageState) Cputime() float64 {
	if pstate.startTime.IsZero() {
		pstate.startTime = time.Now()
	}
	return time.Since(pstate.startTime).Seconds()
}

func (pstate *PackageState) AtExit(f func()) {
	pstate.atExitFuncs = append(pstate.atExitFuncs, f)
}

// Exit exits with code after executing all atExitFuncs.
func (pstate *PackageState) Exit(code int) {
	for i := len(pstate.atExitFuncs) - 1; i >= 0; i-- {
		pstate.atExitFuncs[i]()
	}
	os.Exit(code)
}

// Exitf logs an error message then calls Exit(2).
func (pstate *PackageState) Exitf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, os.Args[0]+": "+format+"\n", a...)
	pstate.nerrors++
	pstate.Exit(2)
}

// Errorf logs an error message.
//
// If more than 20 errors have been printed, exit with an error.
//
// Logging an error means that on exit cmd/link will delete any
// output file and return a non-zero error code.
func (pstate *PackageState) Errorf(s *sym.Symbol, format string, args ...interface{}) {
	if s != nil {
		format = s.Name + ": " + format
	}
	format += "\n"
	fmt.Fprintf(os.Stderr, format, args...)
	pstate.nerrors++
	if *pstate.flagH {
		panic("error")
	}
	if pstate.nerrors > 20 {
		pstate.Exitf("too many errors")
	}
}

func artrim(x []byte) string {
	i := 0
	j := len(x)
	for i < len(x) && x[i] == ' ' {
		i++
	}
	for j > i && x[j-1] == ' ' {
		j--
	}
	return string(x[i:j])
}

func stringtouint32(x []uint32, s string) {
	for i := 0; len(s) > 0; i++ {
		var buf [4]byte
		s = s[copy(buf[:], s):]
		x[i] = binary.LittleEndian.Uint32(buf[:])
	}
}

func (pstate *PackageState) elapsed() float64 {
	return time.Since(pstate.start).Seconds()
}

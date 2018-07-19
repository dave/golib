package ld

import (
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"os"
	"time"
)

// TODO(josharian): delete. See issue 19865.
func (psess *PackageSession) Cputime() float64 {
	if psess.startTime.IsZero() {
		psess.
			startTime = time.Now()
	}
	return time.Since(psess.startTime).Seconds()
}

func (psess *PackageSession) AtExit(f func()) {
	psess.
		atExitFuncs = append(psess.atExitFuncs, f)
}

// Exit exits with code after executing all atExitFuncs.
func (psess *PackageSession) Exit(code int) {
	for i := len(psess.atExitFuncs) - 1; i >= 0; i-- {
		psess.
			atExitFuncs[i]()
	}
	os.Exit(code)
}

// Exitf logs an error message then calls Exit(2).
func (psess *PackageSession) Exitf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, os.Args[0]+": "+format+"\n", a...)
	psess.
		nerrors++
	psess.
		Exit(2)
}

// Errorf logs an error message.
//
// If more than 20 errors have been printed, exit with an error.
//
// Logging an error means that on exit cmd/link will delete any
// output file and return a non-zero error code.
func (psess *PackageSession) Errorf(s *sym.Symbol, format string, args ...interface{}) {
	if s != nil {
		format = s.Name + ": " + format
	}
	format += "\n"
	fmt.Fprintf(os.Stderr, format, args...)
	psess.
		nerrors++
	if *psess.flagH {
		panic("error")
	}
	if psess.nerrors > 20 {
		psess.
			Exitf("too many errors")
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

func (psess *PackageSession) elapsed() float64 {
	return time.Since(psess.start).Seconds()
}

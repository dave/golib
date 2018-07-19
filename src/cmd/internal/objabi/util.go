package objabi

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func envOr(key, value string) string {
	if x := os.Getenv(key); x != "" {
		return x
	}
	return value
}

// set by linker

func goarm() int {
	switch v := envOr("GOARM", defaultGOARM); v {
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	}

	log.Fatalf("Invalid GOARM value. Must be 5, 6, or 7.")
	panic("unreachable")
}

func gomips() string {
	switch v := envOr("GOMIPS", defaultGOMIPS); v {
	case "hardfloat", "softfloat":
		return v
	}
	log.Fatalf("Invalid GOMIPS value. Must be hardfloat or softfloat.")
	panic("unreachable")
}

func gomips64() string {
	switch v := envOr("GOMIPS64", defaultGOMIPS64); v {
	case "hardfloat", "softfloat":
		return v
	}
	log.Fatalf("Invalid GOMIPS64 value. Must be hardfloat or softfloat.")
	panic("unreachable")
}

func Getgoextlinkenabled() string {
	return envOr("GO_EXTLINK_ENABLED", defaultGO_EXTLINK_ENABLED)
}

func (psess *PackageSession) init() {
	for _, f := range strings.Split(goexperiment, ",") {
		if f != "" {
			psess.
				addexp(f)
		}
	}
}

func (psess *PackageSession) Framepointer_enabled(goos, goarch string) bool {
	return psess.framepointer_enabled != 0 && goarch == "amd64" && goos != "nacl"
}

func (psess *PackageSession) addexp(s string) {

	v := 1
	name := s
	if len(name) > 2 && name[:2] == "no" {
		v = 0
		name = name[2:]
	}
	for i := 0; i < len(psess.exper); i++ {
		if psess.exper[i].name == name {
			if psess.exper[i].val != nil {
				*psess.exper[i].val = v
			}
			return
		}
	}

	fmt.Printf("unknown experiment %s\n", s)
	os.Exit(2)
}

// Toolchain experiments.
// These are controlled by the GOEXPERIMENT environment
// variable recorded when the toolchain is built.
// This list is also known to cmd/gc.

func (psess *PackageSession) DefaultExpstring() string {
	return psess.defaultExpstring
}

func (psess *PackageSession) Expstring() string {
	buf := "X"
	for i := range psess.exper {
		if *psess.exper[i].val != 0 {
			buf += "," + psess.exper[i].name
		}
	}
	if buf == "X" {
		buf += ",none"
	}
	return "X:" + buf[2:]
}

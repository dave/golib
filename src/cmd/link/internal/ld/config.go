// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"log"
)

// A BuildMode indicates the sort of object we are building.
//
// Possible build modes are the same as those for the -buildmode flag
// in cmd/go, and are documented in 'go help buildmode'.
type BuildMode uint8

const (
	BuildModeUnset BuildMode = iota
	BuildModeExe
	BuildModePIE
	BuildModeCArchive
	BuildModeCShared
	BuildModeShared
	BuildModePlugin
)

func (mode *BuildMode) Set(pstate *PackageState, s string) error {
	badmode := func() error {
		return fmt.Errorf("buildmode %s not supported on %s/%s", s, pstate.objabi.GOOS, pstate.objabi.GOARCH)
	}
	switch s {
	default:
		return fmt.Errorf("invalid buildmode: %q", s)
	case "exe":
		*mode = BuildModeExe
	case "pie":
		switch pstate.objabi.GOOS {
		case "android", "linux":
		case "darwin", "freebsd":
			switch pstate.objabi.GOARCH {
			case "amd64":
			default:
				return badmode()
			}
		default:
			return badmode()
		}
		*mode = BuildModePIE
	case "c-archive":
		switch pstate.objabi.GOOS {
		case "darwin", "linux":
		case "freebsd":
			switch pstate.objabi.GOARCH {
			case "amd64":
			default:
				return badmode()
			}
		case "windows":
			switch pstate.objabi.GOARCH {
			case "amd64", "386":
			default:
				return badmode()
			}
		default:
			return badmode()
		}
		*mode = BuildModeCArchive
	case "c-shared":
		switch pstate.objabi.GOARCH {
		case "386", "amd64", "arm", "arm64", "ppc64le", "s390x":
		default:
			return badmode()
		}
		*mode = BuildModeCShared
	case "shared":
		switch pstate.objabi.GOOS {
		case "linux":
			switch pstate.objabi.GOARCH {
			case "386", "amd64", "arm", "arm64", "ppc64le", "s390x":
			default:
				return badmode()
			}
		default:
			return badmode()
		}
		*mode = BuildModeShared
	case "plugin":
		switch pstate.objabi.GOOS {
		case "linux":
			switch pstate.objabi.GOARCH {
			case "386", "amd64", "arm", "arm64", "s390x", "ppc64le":
			default:
				return badmode()
			}
		case "darwin":
			switch pstate.objabi.GOARCH {
			case "amd64":
			default:
				return badmode()
			}
		default:
			return badmode()
		}
		*mode = BuildModePlugin
	}
	return nil
}

func (mode *BuildMode) String() string {
	switch *mode {
	case BuildModeUnset:
		return "" // avoid showing a default in usage message
	case BuildModeExe:
		return "exe"
	case BuildModePIE:
		return "pie"
	case BuildModeCArchive:
		return "c-archive"
	case BuildModeCShared:
		return "c-shared"
	case BuildModeShared:
		return "shared"
	case BuildModePlugin:
		return "plugin"
	}
	return fmt.Sprintf("BuildMode(%d)", uint8(*mode))
}

// LinkMode indicates whether an external linker is used for the final link.
type LinkMode uint8

const (
	LinkAuto LinkMode = iota
	LinkInternal
	LinkExternal
)

func (mode *LinkMode) Set(s string) error {
	switch s {
	default:
		return fmt.Errorf("invalid linkmode: %q", s)
	case "auto":
		*mode = LinkAuto
	case "internal":
		*mode = LinkInternal
	case "external":
		*mode = LinkExternal
	}
	return nil
}

func (mode *LinkMode) String() string {
	switch *mode {
	case LinkAuto:
		return "auto"
	case LinkInternal:
		return "internal"
	case LinkExternal:
		return "external"
	}
	return fmt.Sprintf("LinkMode(%d)", uint8(*mode))
}

// mustLinkExternal reports whether the program being linked requires
// the external linker be used to complete the link.
func (pstate *PackageState) mustLinkExternal(ctxt *Link) (res bool, reason string) {
	if ctxt.Debugvlog > 1 {
		defer func() {
			if res {
				log.Printf("external linking is forced by: %s\n", reason)
			}
		}()
	}

	switch pstate.objabi.GOOS {
	case "android":
		return true, "android"
	case "darwin":
		if ctxt.Arch.InFamily(sys.ARM, sys.ARM64) {
			return true, "iOS"
		}
	}

	if *pstate.flagMsan {
		return true, "msan"
	}

	// Internally linking cgo is incomplete on some architectures.
	// https://golang.org/issue/10373
	// https://golang.org/issue/14449
	// https://golang.org/issue/21961
	if pstate.iscgo && ctxt.Arch.InFamily(sys.ARM64, sys.MIPS64, sys.MIPS, sys.PPC64) {
		return true, pstate.objabi.GOARCH + " does not support internal cgo"
	}

	// When the race flag is set, the LLVM tsan relocatable file is linked
	// into the final binary, which means external linking is required because
	// internal linking does not support it.
	if *pstate.flagRace && ctxt.Arch.InFamily(sys.PPC64) {
		return true, "race on ppc64le"
	}

	// Some build modes require work the internal linker cannot do (yet).
	switch ctxt.BuildMode {
	case BuildModeCArchive:
		return true, "buildmode=c-archive"
	case BuildModeCShared:
		return true, "buildmode=c-shared"
	case BuildModePIE:
		switch pstate.objabi.GOOS + "/" + pstate.objabi.GOARCH {
		case "linux/amd64":
		default:
			// Internal linking does not support TLS_IE.
			return true, "buildmode=pie"
		}
	case BuildModePlugin:
		return true, "buildmode=plugin"
	case BuildModeShared:
		return true, "buildmode=shared"
	}
	if ctxt.linkShared {
		return true, "dynamically linking with a shared library"
	}

	return false, ""
}

// determineLinkMode sets ctxt.LinkMode.
//
// It is called after flags are processed and inputs are processed,
// so the ctxt.LinkMode variable has an initial value from the -linkmode
// flag and the iscgo externalobj variables are set.
func (pstate *PackageState) determineLinkMode(ctxt *Link) {
	switch ctxt.LinkMode {
	case LinkAuto:
		// The environment variable GO_EXTLINK_ENABLED controls the
		// default value of -linkmode. If it is not set when the
		// linker is called we take the value it was set to when
		// cmd/link was compiled. (See make.bash.)
		switch objabi.Getgoextlinkenabled() {
		case "0":
			if needed, reason := pstate.mustLinkExternal(ctxt); needed {
				pstate.Exitf("internal linking requested via GO_EXTLINK_ENABLED, but external linking required: %s", reason)
			}
			ctxt.LinkMode = LinkInternal
		case "1":
			if pstate.objabi.GOARCH == "ppc64" {
				pstate.Exitf("external linking requested via GO_EXTLINK_ENABLED but not supported for linux/ppc64")
			}
			ctxt.LinkMode = LinkExternal
		default:
			if needed, _ := pstate.mustLinkExternal(ctxt); needed {
				ctxt.LinkMode = LinkExternal
			} else if pstate.iscgo && pstate.externalobj {
				ctxt.LinkMode = LinkExternal
			} else if ctxt.BuildMode == BuildModePIE {
				ctxt.LinkMode = LinkExternal // https://golang.org/issue/18968
			} else {
				ctxt.LinkMode = LinkInternal
			}
			if pstate.objabi.GOARCH == "ppc64" && ctxt.LinkMode == LinkExternal {
				pstate.Exitf("external linking is not supported for linux/ppc64")
			}
		}
	case LinkInternal:
		if needed, reason := pstate.mustLinkExternal(ctxt); needed {
			pstate.Exitf("internal linking requested but external linking required: %s", reason)
		}
	case LinkExternal:
		if pstate.objabi.GOARCH == "ppc64" {
			pstate.Exitf("external linking not supported for linux/ppc64")
		}
	}
}

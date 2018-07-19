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

func (mode *BuildMode) Set(psess *PackageSession, s string) error {
	badmode := func() error {
		return fmt.Errorf("buildmode %s not supported on %s/%s", s, psess.objabi.GOOS, psess.objabi.GOARCH)
	}
	switch s {
	default:
		return fmt.Errorf("invalid buildmode: %q", s)
	case "exe":
		*mode = BuildModeExe
	case "pie":
		switch psess.objabi.GOOS {
		case "android", "linux":
		case "darwin", "freebsd":
			switch psess.objabi.GOARCH {
			case "amd64":
			default:
				return badmode()
			}
		default:
			return badmode()
		}
		*mode = BuildModePIE
	case "c-archive":
		switch psess.objabi.GOOS {
		case "darwin", "linux":
		case "freebsd":
			switch psess.objabi.GOARCH {
			case "amd64":
			default:
				return badmode()
			}
		case "windows":
			switch psess.objabi.GOARCH {
			case "amd64", "386":
			default:
				return badmode()
			}
		default:
			return badmode()
		}
		*mode = BuildModeCArchive
	case "c-shared":
		switch psess.objabi.GOARCH {
		case "386", "amd64", "arm", "arm64", "ppc64le", "s390x":
		default:
			return badmode()
		}
		*mode = BuildModeCShared
	case "shared":
		switch psess.objabi.GOOS {
		case "linux":
			switch psess.objabi.GOARCH {
			case "386", "amd64", "arm", "arm64", "ppc64le", "s390x":
			default:
				return badmode()
			}
		default:
			return badmode()
		}
		*mode = BuildModeShared
	case "plugin":
		switch psess.objabi.GOOS {
		case "linux":
			switch psess.objabi.GOARCH {
			case "386", "amd64", "arm", "arm64", "s390x", "ppc64le":
			default:
				return badmode()
			}
		case "darwin":
			switch psess.objabi.GOARCH {
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
		return ""
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
func (psess *PackageSession) mustLinkExternal(ctxt *Link) (res bool, reason string) {
	if ctxt.Debugvlog > 1 {
		defer func() {
			if res {
				log.Printf("external linking is forced by: %s\n", reason)
			}
		}()
	}

	switch psess.objabi.GOOS {
	case "android":
		return true, "android"
	case "darwin":
		if ctxt.Arch.InFamily(sys.ARM, sys.ARM64) {
			return true, "iOS"
		}
	}

	if *psess.flagMsan {
		return true, "msan"
	}

	if psess.iscgo && ctxt.Arch.InFamily(sys.ARM64, sys.MIPS64, sys.MIPS, sys.PPC64) {
		return true, psess.objabi.GOARCH + " does not support internal cgo"
	}

	if *psess.flagRace && ctxt.Arch.InFamily(sys.PPC64) {
		return true, "race on ppc64le"
	}

	switch ctxt.BuildMode {
	case BuildModeCArchive:
		return true, "buildmode=c-archive"
	case BuildModeCShared:
		return true, "buildmode=c-shared"
	case BuildModePIE:
		switch psess.objabi.GOOS + "/" + psess.objabi.GOARCH {
		case "linux/amd64":
		default:

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
func (psess *PackageSession) determineLinkMode(ctxt *Link) {
	switch ctxt.LinkMode {
	case LinkAuto:

		switch objabi.Getgoextlinkenabled() {
		case "0":
			if needed, reason := psess.mustLinkExternal(ctxt); needed {
				psess.
					Exitf("internal linking requested via GO_EXTLINK_ENABLED, but external linking required: %s", reason)
			}
			ctxt.LinkMode = LinkInternal
		case "1":
			if psess.objabi.GOARCH == "ppc64" {
				psess.
					Exitf("external linking requested via GO_EXTLINK_ENABLED but not supported for linux/ppc64")
			}
			ctxt.LinkMode = LinkExternal
		default:
			if needed, _ := psess.mustLinkExternal(ctxt); needed {
				ctxt.LinkMode = LinkExternal
			} else if psess.iscgo && psess.externalobj {
				ctxt.LinkMode = LinkExternal
			} else if ctxt.BuildMode == BuildModePIE {
				ctxt.LinkMode = LinkExternal
			} else {
				ctxt.LinkMode = LinkInternal
			}
			if psess.objabi.GOARCH == "ppc64" && ctxt.LinkMode == LinkExternal {
				psess.
					Exitf("external linking is not supported for linux/ppc64")
			}
		}
	case LinkInternal:
		if needed, reason := psess.mustLinkExternal(ctxt); needed {
			psess.
				Exitf("internal linking requested but external linking required: %s", reason)
		}
	case LinkExternal:
		if psess.objabi.GOARCH == "ppc64" {
			psess.
				Exitf("external linking not supported for linux/ppc64")
		}
	}
}

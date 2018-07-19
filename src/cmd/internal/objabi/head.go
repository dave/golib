package objabi

import "fmt"

// HeadType is the executable header type.
type HeadType uint8

const (
	Hunknown HeadType = iota
	Hdarwin
	Hdragonfly
	Hfreebsd
	Hjs
	Hlinux
	Hnacl
	Hnetbsd
	Hopenbsd
	Hplan9
	Hsolaris
	Hwindows
)

func (h *HeadType) Set(s string) error {
	switch s {
	case "darwin":
		*h = Hdarwin
	case "dragonfly":
		*h = Hdragonfly
	case "freebsd":
		*h = Hfreebsd
	case "js":
		*h = Hjs
	case "linux", "android":
		*h = Hlinux
	case "nacl":
		*h = Hnacl
	case "netbsd":
		*h = Hnetbsd
	case "openbsd":
		*h = Hopenbsd
	case "plan9":
		*h = Hplan9
	case "solaris":
		*h = Hsolaris
	case "windows":
		*h = Hwindows
	default:
		return fmt.Errorf("invalid headtype: %q", s)
	}
	return nil
}

func (h *HeadType) String() string {
	switch *h {
	case Hdarwin:
		return "darwin"
	case Hdragonfly:
		return "dragonfly"
	case Hfreebsd:
		return "freebsd"
	case Hjs:
		return "js"
	case Hlinux:
		return "linux"
	case Hnacl:
		return "nacl"
	case Hnetbsd:
		return "netbsd"
	case Hopenbsd:
		return "openbsd"
	case Hplan9:
		return "plan9"
	case Hsolaris:
		return "solaris"
	case Hwindows:
		return "windows"
	}
	return fmt.Sprintf("HeadType(%d)", *h)
}

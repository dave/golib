// +build !windows

package testenv

import (
	"runtime"
)

func hasSymlink() (ok bool, reason string) {
	switch runtime.GOOS {
	case "android", "nacl", "plan9":
		return false, ""
	}

	return true, ""
}

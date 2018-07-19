// +build go1.8

package gc

import "runtime"

func startMutexProfiling() {
	runtime.SetMutexProfileFraction(1)
}

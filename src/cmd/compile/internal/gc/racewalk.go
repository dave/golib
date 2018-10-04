// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/src"
	"github.com/dave/golib/src/cmd/internal/sys"
)

func (pstate *PackageState) ispkgin(pkgs []string) bool {
	if pstate.myimportpath != "" {
		for _, p := range pkgs {
			if pstate.myimportpath == p {
				return true
			}
		}
	}

	return false
}

func (pstate *PackageState) instrument(fn *Node) {
	if fn.Func.Pragma&Norace != 0 {
		return
	}

	if !pstate.flag_race || !pstate.ispkgin(pstate.norace_inst_pkgs) {
		fn.Func.SetInstrumentBody(true)
	}

	if pstate.flag_race {
		lno := pstate.lineno
		pstate.lineno = pstate.src.NoXPos

		if pstate.thearch.LinkArch.Arch == pstate.sys.ArchPPC64LE {
			fn.Func.Enter.Prepend(pstate.mkcall("racefuncenterfp", nil, nil))
			fn.Func.Exit.Append(pstate.mkcall("racefuncexit", nil, nil))
		} else {

			// nodpc is the PC of the caller as extracted by
			// getcallerpc. We use -widthptr(FP) for x86.
			// BUG: This only works for amd64. This will not
			// work on arm or others that might support
			// race in the future.
			nodpc := pstate.nodfp.copy()
			nodpc.Type = pstate.types.Types[TUINTPTR]
			nodpc.Xoffset = int64(-pstate.Widthptr)
			fn.Func.Dcl = append(fn.Func.Dcl, nodpc)
			fn.Func.Enter.Prepend(pstate.mkcall("racefuncenter", nil, nil, nodpc))
			fn.Func.Exit.Append(pstate.mkcall("racefuncexit", nil, nil))
		}
		pstate.lineno = lno
	}
}

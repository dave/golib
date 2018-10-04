// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/src"
	"sync"
)

const (
	BADWIDTH        = types.BADWIDTH
	maxStackVarSize = 10 * 1024 * 1024
)

// isRuntimePkg reports whether p is package runtime.
func (pstate *PackageState) isRuntimePkg(p *types.Pkg) bool {
	if pstate.compiling_runtime && p == pstate.localpkg {
		return true
	}
	return p.Path == "runtime"
}

// The Class of a variable/function describes the "storage class"
// of a variable or function. During parsing, storage classes are
// called declaration contexts.
type Class uint8

//go:generate stringer -type=Class
const (
	Pxxx      Class = iota // no class; used during ssa conversion to indicate pseudo-variables
	PEXTERN                // global variable
	PAUTO                  // local variables
	PAUTOHEAP              // local variable or parameter moved to heap
	PPARAM                 // input arguments
	PPARAMOUT              // output results
	PFUNC                  // global function

	PDISCARD // discard during parse of duplicate import
// Careful: Class is stored in three bits in Node.flags.
// Adding a new Class will overflow that.
)

func init() {
	if PDISCARD != 7 {
		panic("PDISCARD changed; does all Class values still fit in three bits?")
	}
}

// interface to back end

type Arch struct {
	LinkArch *obj.LinkArch

	REGSP     int
	MAXWIDTH  int64
	Use387    bool // should 386 backend use 387 FP instructions instead of sse2.
	SoftFloat bool

	PadFrame  func(int64) int64
	ZeroRange func(*Progs, *obj.Prog, int64, int64, *uint32) *obj.Prog
	Ginsnop   func(*Progs)

	// SSAMarkMoves marks any MOVXconst ops that need to avoid clobbering flags.
	SSAMarkMoves func(*SSAGenState, *ssa.Block)

	// SSAGenValue emits Prog(s) for the Value.
	SSAGenValue func(*SSAGenState, *ssa.Value)

	// SSAGenBlock emits end-of-block Progs. SSAGenValue should be called
	// for all values in the block before SSAGenBlock.
	SSAGenBlock func(s *SSAGenState, b, next *ssa.Block)

	// ZeroAuto emits code to zero the given auto stack variable.
	// ZeroAuto must not use any non-temporary registers.
	// ZeroAuto will only be called for variables which contain a pointer.
	ZeroAuto func(*Progs, *Node)
}

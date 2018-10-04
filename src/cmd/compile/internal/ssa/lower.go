// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// convert to machine-dependent ops
func (pstate *PackageState) lower(f *Func) {
	// repeat rewrites until we find no more rewrites
	pstate.applyRewrite(f, f.Config.lowerBlock, f.Config.lowerValue)
}

// checkLower checks for unlowered opcodes and fails if we find one.
func (pstate *PackageState) checkLower(f *Func) {
	// Needs to be a separate phase because it must run after both
	// lowering and a subsequent dead code elimination (because lowering
	// rules may leave dead generic ops behind).
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if !pstate.opcodeTable[v.Op].generic {
				continue // lowered
			}
			switch v.Op {
			case OpSP, OpSB, OpInitMem, OpArg, OpPhi, OpVarDef, OpVarKill, OpVarLive, OpKeepAlive, OpSelect0, OpSelect1, OpConvert:
				continue // ok not to lower
			case OpGetG:
				if f.Config.hasGReg {
					// has hardware g register, regalloc takes care of it
					continue // ok not to lower
				}
			}
			s := "not lowered: " + v.String() + ", " + v.Op.String(pstate) + " " + v.Type.SimpleString(pstate.types)
			for _, a := range v.Args {
				s += " " + a.Type.SimpleString(pstate.types)
			}
			f.Fatalf("%s", s)
		}
	}
}

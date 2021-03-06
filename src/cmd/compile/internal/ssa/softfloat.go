// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import "math"

func (pstate *PackageState) softfloat(f *Func) {
	if !f.Config.SoftFloat {
		return
	}
	newInt64 := false

	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Type.IsFloat() {
				switch v.Op {
				case OpPhi, OpLoad, OpArg:
					if v.Type.Size(pstate.types) == 4 {
						v.Type = f.Config.Types.UInt32
					} else {
						v.Type = f.Config.Types.UInt64
					}
				case OpConst32F:
					v.Op = OpConst32
					v.Type = f.Config.Types.UInt32
					v.AuxInt = int64(int32(math.Float32bits(i2f32(v.AuxInt))))
				case OpConst64F:
					v.Op = OpConst64
					v.Type = f.Config.Types.UInt64
				case OpNeg32F:
					arg0 := v.Args[0]
					v.reset(OpXor32)
					v.Type = f.Config.Types.UInt32
					v.AddArg(arg0)
					mask := v.Block.NewValue0(v.Pos, OpConst32, v.Type)
					mask.AuxInt = -0x80000000
					v.AddArg(mask)
				case OpNeg64F:
					arg0 := v.Args[0]
					v.reset(OpXor64)
					v.Type = f.Config.Types.UInt64
					v.AddArg(arg0)
					mask := v.Block.NewValue0(v.Pos, OpConst64, v.Type)
					mask.AuxInt = -0x8000000000000000
					v.AddArg(mask)
				case OpRound32F:
					v.Op = OpCopy
					v.Type = f.Config.Types.UInt32
				case OpRound64F:
					v.Op = OpCopy
					v.Type = f.Config.Types.UInt64
				}
				newInt64 = newInt64 || v.Type.Size(pstate.types) == 8
			}
		}
	}

	if newInt64 && f.Config.RegSize == 4 {
		// On 32bit arch, decompose Uint64 introduced in the switch above.
		pstate.decomposeBuiltIn(f)
		pstate.applyRewrite(f, rewriteBlockdec64, pstate.rewriteValuedec64)
	}

}

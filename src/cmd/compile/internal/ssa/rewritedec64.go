package ssa

import "github.com/dave/golib/src/cmd/compile/internal/types"

// in case not otherwise used
// in case not otherwise used
// in case not otherwise used
// in case not otherwise used

func (psess *PackageSession) rewriteValuedec64(v *Value) bool {
	switch v.Op {
	case OpAdd64:
		return psess.rewriteValuedec64_OpAdd64_0(v)
	case OpAnd64:
		return rewriteValuedec64_OpAnd64_0(v)
	case OpArg:
		return psess.rewriteValuedec64_OpArg_0(v)
	case OpBitLen64:
		return rewriteValuedec64_OpBitLen64_0(v)
	case OpBswap64:
		return rewriteValuedec64_OpBswap64_0(v)
	case OpCom64:
		return rewriteValuedec64_OpCom64_0(v)
	case OpConst64:
		return rewriteValuedec64_OpConst64_0(v)
	case OpCtz64:
		return rewriteValuedec64_OpCtz64_0(v)
	case OpCtz64NonZero:
		return rewriteValuedec64_OpCtz64NonZero_0(v)
	case OpEq64:
		return rewriteValuedec64_OpEq64_0(v)
	case OpGeq64:
		return rewriteValuedec64_OpGeq64_0(v)
	case OpGeq64U:
		return rewriteValuedec64_OpGeq64U_0(v)
	case OpGreater64:
		return rewriteValuedec64_OpGreater64_0(v)
	case OpGreater64U:
		return rewriteValuedec64_OpGreater64U_0(v)
	case OpInt64Hi:
		return rewriteValuedec64_OpInt64Hi_0(v)
	case OpInt64Lo:
		return rewriteValuedec64_OpInt64Lo_0(v)
	case OpLeq64:
		return rewriteValuedec64_OpLeq64_0(v)
	case OpLeq64U:
		return rewriteValuedec64_OpLeq64U_0(v)
	case OpLess64:
		return rewriteValuedec64_OpLess64_0(v)
	case OpLess64U:
		return rewriteValuedec64_OpLess64U_0(v)
	case OpLoad:
		return psess.rewriteValuedec64_OpLoad_0(v)
	case OpLsh16x64:
		return rewriteValuedec64_OpLsh16x64_0(v)
	case OpLsh32x64:
		return rewriteValuedec64_OpLsh32x64_0(v)
	case OpLsh64x16:
		return rewriteValuedec64_OpLsh64x16_0(v)
	case OpLsh64x32:
		return rewriteValuedec64_OpLsh64x32_0(v)
	case OpLsh64x64:
		return rewriteValuedec64_OpLsh64x64_0(v)
	case OpLsh64x8:
		return rewriteValuedec64_OpLsh64x8_0(v)
	case OpLsh8x64:
		return rewriteValuedec64_OpLsh8x64_0(v)
	case OpMul64:
		return rewriteValuedec64_OpMul64_0(v)
	case OpNeg64:
		return rewriteValuedec64_OpNeg64_0(v)
	case OpNeq64:
		return rewriteValuedec64_OpNeq64_0(v)
	case OpOr64:
		return rewriteValuedec64_OpOr64_0(v)
	case OpRsh16Ux64:
		return rewriteValuedec64_OpRsh16Ux64_0(v)
	case OpRsh16x64:
		return rewriteValuedec64_OpRsh16x64_0(v)
	case OpRsh32Ux64:
		return rewriteValuedec64_OpRsh32Ux64_0(v)
	case OpRsh32x64:
		return rewriteValuedec64_OpRsh32x64_0(v)
	case OpRsh64Ux16:
		return rewriteValuedec64_OpRsh64Ux16_0(v)
	case OpRsh64Ux32:
		return rewriteValuedec64_OpRsh64Ux32_0(v)
	case OpRsh64Ux64:
		return rewriteValuedec64_OpRsh64Ux64_0(v)
	case OpRsh64Ux8:
		return rewriteValuedec64_OpRsh64Ux8_0(v)
	case OpRsh64x16:
		return rewriteValuedec64_OpRsh64x16_0(v)
	case OpRsh64x32:
		return rewriteValuedec64_OpRsh64x32_0(v)
	case OpRsh64x64:
		return rewriteValuedec64_OpRsh64x64_0(v)
	case OpRsh64x8:
		return rewriteValuedec64_OpRsh64x8_0(v)
	case OpRsh8Ux64:
		return rewriteValuedec64_OpRsh8Ux64_0(v)
	case OpRsh8x64:
		return rewriteValuedec64_OpRsh8x64_0(v)
	case OpSignExt16to64:
		return rewriteValuedec64_OpSignExt16to64_0(v)
	case OpSignExt32to64:
		return rewriteValuedec64_OpSignExt32to64_0(v)
	case OpSignExt8to64:
		return rewriteValuedec64_OpSignExt8to64_0(v)
	case OpStore:
		return psess.rewriteValuedec64_OpStore_0(v)
	case OpSub64:
		return psess.rewriteValuedec64_OpSub64_0(v)
	case OpTrunc64to16:
		return rewriteValuedec64_OpTrunc64to16_0(v)
	case OpTrunc64to32:
		return rewriteValuedec64_OpTrunc64to32_0(v)
	case OpTrunc64to8:
		return rewriteValuedec64_OpTrunc64to8_0(v)
	case OpXor64:
		return rewriteValuedec64_OpXor64_0(v)
	case OpZeroExt16to64:
		return rewriteValuedec64_OpZeroExt16to64_0(v)
	case OpZeroExt32to64:
		return rewriteValuedec64_OpZeroExt32to64_0(v)
	case OpZeroExt8to64:
		return rewriteValuedec64_OpZeroExt8to64_0(v)
	}
	return false
}
func (psess *PackageSession) rewriteValuedec64_OpAdd64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpAdd32withcarry, typ.Int32)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSelect1, psess.types.TypeFlags)
		v4 := b.NewValue0(v.Pos, OpAdd32carry, types.NewTuple(typ.UInt32, psess.types.TypeFlags))
		v5 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v6.AddArg(y)
		v4.AddArg(v6)
		v3.AddArg(v4)
		v0.AddArg(v3)
		v.AddArg(v0)
		v7 := b.NewValue0(v.Pos, OpSelect0, typ.UInt32)
		v8 := b.NewValue0(v.Pos, OpAdd32carry, types.NewTuple(typ.UInt32, psess.types.TypeFlags))
		v9 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v9.AddArg(x)
		v8.AddArg(v9)
		v10 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v10.AddArg(y)
		v8.AddArg(v10)
		v7.AddArg(v8)
		v.AddArg(v7)
		return true
	}
}
func rewriteValuedec64_OpAnd64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpAnd32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpAnd32, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v4.AddArg(x)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v5.AddArg(y)
		v3.AddArg(v5)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValuedec64_OpArg_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		off := v.AuxInt
		n := v.Aux
		if !(psess.is64BitInt(v.Type) && !config.BigEndian && v.Type.IsSigned()) {
			break
		}
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpArg, typ.Int32)
		v0.AuxInt = off + 4
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, typ.UInt32)
		v1.AuxInt = off
		v1.Aux = n
		v.AddArg(v1)
		return true
	}

	for {
		off := v.AuxInt
		n := v.Aux
		if !(psess.is64BitInt(v.Type) && !config.BigEndian && !v.Type.IsSigned()) {
			break
		}
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpArg, typ.UInt32)
		v0.AuxInt = off + 4
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, typ.UInt32)
		v1.AuxInt = off
		v1.Aux = n
		v.AddArg(v1)
		return true
	}

	for {
		off := v.AuxInt
		n := v.Aux
		if !(psess.is64BitInt(v.Type) && config.BigEndian && v.Type.IsSigned()) {
			break
		}
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpArg, typ.Int32)
		v0.AuxInt = off
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, typ.UInt32)
		v1.AuxInt = off + 4
		v1.Aux = n
		v.AddArg(v1)
		return true
	}

	for {
		off := v.AuxInt
		n := v.Aux
		if !(psess.is64BitInt(v.Type) && config.BigEndian && !v.Type.IsSigned()) {
			break
		}
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpArg, typ.UInt32)
		v0.AuxInt = off
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, typ.UInt32)
		v1.AuxInt = off + 4
		v1.Aux = n
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec64_OpBitLen64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpAdd32)
		v.Type = typ.Int
		v0 := b.NewValue0(v.Pos, OpBitLen32, typ.Int)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpBitLen32, typ.Int)
		v3 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v4.AddArg(x)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v6 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v6.AddArg(x)
		v5.AddArg(v6)
		v3.AddArg(v5)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValuedec64_OpBswap64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpBswap32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpBswap32, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v3.AddArg(x)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValuedec64_OpCom64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpCom32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpCom32, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v3.AddArg(x)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValuedec64_OpConst64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		c := v.AuxInt
		if !(t.IsSigned()) {
			break
		}
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpConst32, typ.Int32)
		v0.AuxInt = c >> 32
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v1.AuxInt = int64(int32(c))
		v.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		c := v.AuxInt
		if !(!t.IsSigned()) {
			break
		}
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v0.AuxInt = c >> 32
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v1.AuxInt = int64(int32(c))
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec64_OpCtz64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpAdd32)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpCtz32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpAnd32, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpCom32, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v5 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v6 := b.NewValue0(v.Pos, OpCtz32, typ.UInt32)
		v7 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v7.AddArg(x)
		v6.AddArg(v7)
		v2.AddArg(v6)
		v.AddArg(v2)
		return true
	}
}
func rewriteValuedec64_OpCtz64NonZero_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCtz64)
		v.AddArg(x)
		return true
	}
}
func rewriteValuedec64_OpEq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpAndB)
		v0 := b.NewValue0(v.Pos, OpEq32, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpEq32, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v4.AddArg(x)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v5.AddArg(y)
		v3.AddArg(v5)
		v.AddArg(v3)
		return true
	}
}
func rewriteValuedec64_OpGeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpOrB)
		v0 := b.NewValue0(v.Pos, OpGreater32, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpAndB, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpEq32, typ.Bool)
		v5 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v6.AddArg(y)
		v4.AddArg(v6)
		v3.AddArg(v4)
		v7 := b.NewValue0(v.Pos, OpGeq32U, typ.Bool)
		v8 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v8.AddArg(x)
		v7.AddArg(v8)
		v9 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v9.AddArg(y)
		v7.AddArg(v9)
		v3.AddArg(v7)
		v.AddArg(v3)
		return true
	}
}
func rewriteValuedec64_OpGeq64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpOrB)
		v0 := b.NewValue0(v.Pos, OpGreater32U, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpAndB, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpEq32, typ.Bool)
		v5 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v6.AddArg(y)
		v4.AddArg(v6)
		v3.AddArg(v4)
		v7 := b.NewValue0(v.Pos, OpGeq32U, typ.Bool)
		v8 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v8.AddArg(x)
		v7.AddArg(v8)
		v9 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v9.AddArg(y)
		v7.AddArg(v9)
		v3.AddArg(v7)
		v.AddArg(v3)
		return true
	}
}
func rewriteValuedec64_OpGreater64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpOrB)
		v0 := b.NewValue0(v.Pos, OpGreater32, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpAndB, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpEq32, typ.Bool)
		v5 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v6.AddArg(y)
		v4.AddArg(v6)
		v3.AddArg(v4)
		v7 := b.NewValue0(v.Pos, OpGreater32U, typ.Bool)
		v8 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v8.AddArg(x)
		v7.AddArg(v8)
		v9 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v9.AddArg(y)
		v7.AddArg(v9)
		v3.AddArg(v7)
		v.AddArg(v3)
		return true
	}
}
func rewriteValuedec64_OpGreater64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpOrB)
		v0 := b.NewValue0(v.Pos, OpGreater32U, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpAndB, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpEq32, typ.Bool)
		v5 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v6.AddArg(y)
		v4.AddArg(v6)
		v3.AddArg(v4)
		v7 := b.NewValue0(v.Pos, OpGreater32U, typ.Bool)
		v8 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v8.AddArg(x)
		v7.AddArg(v8)
		v9 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v9.AddArg(y)
		v7.AddArg(v9)
		v3.AddArg(v7)
		v.AddArg(v3)
		return true
	}
}
func rewriteValuedec64_OpInt64Hi_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		hi := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = hi.Type
		v.AddArg(hi)
		return true
	}
	return false
}
func rewriteValuedec64_OpInt64Lo_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		lo := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = lo.Type
		v.AddArg(lo)
		return true
	}
	return false
}
func rewriteValuedec64_OpLeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpOrB)
		v0 := b.NewValue0(v.Pos, OpLess32, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpAndB, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpEq32, typ.Bool)
		v5 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v6.AddArg(y)
		v4.AddArg(v6)
		v3.AddArg(v4)
		v7 := b.NewValue0(v.Pos, OpLeq32U, typ.Bool)
		v8 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v8.AddArg(x)
		v7.AddArg(v8)
		v9 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v9.AddArg(y)
		v7.AddArg(v9)
		v3.AddArg(v7)
		v.AddArg(v3)
		return true
	}
}
func rewriteValuedec64_OpLeq64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpOrB)
		v0 := b.NewValue0(v.Pos, OpLess32U, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpAndB, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpEq32, typ.Bool)
		v5 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v6.AddArg(y)
		v4.AddArg(v6)
		v3.AddArg(v4)
		v7 := b.NewValue0(v.Pos, OpLeq32U, typ.Bool)
		v8 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v8.AddArg(x)
		v7.AddArg(v8)
		v9 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v9.AddArg(y)
		v7.AddArg(v9)
		v3.AddArg(v7)
		v.AddArg(v3)
		return true
	}
}
func rewriteValuedec64_OpLess64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpOrB)
		v0 := b.NewValue0(v.Pos, OpLess32, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpAndB, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpEq32, typ.Bool)
		v5 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v6.AddArg(y)
		v4.AddArg(v6)
		v3.AddArg(v4)
		v7 := b.NewValue0(v.Pos, OpLess32U, typ.Bool)
		v8 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v8.AddArg(x)
		v7.AddArg(v8)
		v9 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v9.AddArg(y)
		v7.AddArg(v9)
		v3.AddArg(v7)
		v.AddArg(v3)
		return true
	}
}
func rewriteValuedec64_OpLess64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpOrB)
		v0 := b.NewValue0(v.Pos, OpLess32U, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpAndB, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpEq32, typ.Bool)
		v5 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v6.AddArg(y)
		v4.AddArg(v6)
		v3.AddArg(v4)
		v7 := b.NewValue0(v.Pos, OpLess32U, typ.Bool)
		v8 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v8.AddArg(x)
		v7.AddArg(v8)
		v9 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v9.AddArg(y)
		v7.AddArg(v9)
		v3.AddArg(v7)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValuedec64_OpLoad_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is64BitInt(t) && !config.BigEndian && t.IsSigned()) {
			break
		}
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpLoad, typ.Int32)
		v1 := b.NewValue0(v.Pos, OpOffPtr, typ.Int32Ptr)
		v1.AuxInt = 4
		v1.AddArg(ptr)
		v0.AddArg(v1)
		v0.AddArg(mem)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpLoad, typ.UInt32)
		v2.AddArg(ptr)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is64BitInt(t) && !config.BigEndian && !t.IsSigned()) {
			break
		}
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpLoad, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpOffPtr, typ.UInt32Ptr)
		v1.AuxInt = 4
		v1.AddArg(ptr)
		v0.AddArg(v1)
		v0.AddArg(mem)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpLoad, typ.UInt32)
		v2.AddArg(ptr)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is64BitInt(t) && config.BigEndian && t.IsSigned()) {
			break
		}
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpLoad, typ.Int32)
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpLoad, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpOffPtr, typ.UInt32Ptr)
		v2.AuxInt = 4
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is64BitInt(t) && config.BigEndian && !t.IsSigned()) {
			break
		}
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpLoad, typ.UInt32)
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpLoad, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpOffPtr, typ.UInt32Ptr)
		v2.AuxInt = 4
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec64_OpLsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpLsh16x32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpLsh16x32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpLsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpLsh32x32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpLsh32x32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpLsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		hi := v_0.Args[0]
		lo := v_0.Args[1]
		s := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpLsh32x16, typ.UInt32)
		v2.AddArg(hi)
		v2.AddArg(s)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpRsh32Ux16, typ.UInt32)
		v3.AddArg(lo)
		v4 := b.NewValue0(v.Pos, OpSub16, typ.UInt16)
		v5 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
		v5.AuxInt = 32
		v4.AddArg(v5)
		v4.AddArg(s)
		v3.AddArg(v4)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v6 := b.NewValue0(v.Pos, OpLsh32x16, typ.UInt32)
		v6.AddArg(lo)
		v7 := b.NewValue0(v.Pos, OpSub16, typ.UInt16)
		v7.AddArg(s)
		v8 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
		v8.AuxInt = 32
		v7.AddArg(v8)
		v6.AddArg(v7)
		v0.AddArg(v6)
		v.AddArg(v0)
		v9 := b.NewValue0(v.Pos, OpLsh32x16, typ.UInt32)
		v9.AddArg(lo)
		v9.AddArg(s)
		v.AddArg(v9)
		return true
	}
	return false
}
func rewriteValuedec64_OpLsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		hi := v_0.Args[0]
		lo := v_0.Args[1]
		s := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpLsh32x32, typ.UInt32)
		v2.AddArg(hi)
		v2.AddArg(s)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpRsh32Ux32, typ.UInt32)
		v3.AddArg(lo)
		v4 := b.NewValue0(v.Pos, OpSub32, typ.UInt32)
		v5 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v5.AuxInt = 32
		v4.AddArg(v5)
		v4.AddArg(s)
		v3.AddArg(v4)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v6 := b.NewValue0(v.Pos, OpLsh32x32, typ.UInt32)
		v6.AddArg(lo)
		v7 := b.NewValue0(v.Pos, OpSub32, typ.UInt32)
		v7.AddArg(s)
		v8 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v8.AuxInt = 32
		v7.AddArg(v8)
		v6.AddArg(v7)
		v0.AddArg(v6)
		v.AddArg(v0)
		v9 := b.NewValue0(v.Pos, OpLsh32x32, typ.UInt32)
		v9.AddArg(lo)
		v9.AddArg(s)
		v.AddArg(v9)
		return true
	}
	return false
}
func rewriteValuedec64_OpLsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpLsh64x32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpLsh64x32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpLsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		hi := v_0.Args[0]
		lo := v_0.Args[1]
		s := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpLsh32x8, typ.UInt32)
		v2.AddArg(hi)
		v2.AddArg(s)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpRsh32Ux8, typ.UInt32)
		v3.AddArg(lo)
		v4 := b.NewValue0(v.Pos, OpSub8, typ.UInt8)
		v5 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
		v5.AuxInt = 32
		v4.AddArg(v5)
		v4.AddArg(s)
		v3.AddArg(v4)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v6 := b.NewValue0(v.Pos, OpLsh32x8, typ.UInt32)
		v6.AddArg(lo)
		v7 := b.NewValue0(v.Pos, OpSub8, typ.UInt8)
		v7.AddArg(s)
		v8 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
		v8.AuxInt = 32
		v7.AddArg(v8)
		v6.AddArg(v7)
		v0.AddArg(v6)
		v.AddArg(v0)
		v9 := b.NewValue0(v.Pos, OpLsh32x8, typ.UInt32)
		v9.AddArg(lo)
		v9.AddArg(s)
		v.AddArg(v9)
		return true
	}
	return false
}
func rewriteValuedec64_OpLsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpLsh8x32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpLsh8x32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpMul64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpAdd32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpAdd32, typ.UInt32)
		v5 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v6 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v6.AddArg(x)
		v5.AddArg(v6)
		v7 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v7.AddArg(y)
		v5.AddArg(v7)
		v4.AddArg(v5)
		v8 := b.NewValue0(v.Pos, OpSelect0, typ.UInt32)
		v9 := b.NewValue0(v.Pos, OpMul32uhilo, types.NewTuple(typ.UInt32, typ.UInt32))
		v10 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v10.AddArg(x)
		v9.AddArg(v10)
		v11 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v11.AddArg(y)
		v9.AddArg(v11)
		v8.AddArg(v9)
		v4.AddArg(v8)
		v0.AddArg(v4)
		v.AddArg(v0)
		v12 := b.NewValue0(v.Pos, OpSelect1, typ.UInt32)
		v13 := b.NewValue0(v.Pos, OpMul32uhilo, types.NewTuple(typ.UInt32, typ.UInt32))
		v14 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v14.AddArg(x)
		v13.AddArg(v14)
		v15 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v15.AddArg(y)
		v13.AddArg(v15)
		v12.AddArg(v13)
		v.AddArg(v12)
		return true
	}
}
func rewriteValuedec64_OpNeg64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValuedec64_OpNeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpOrB)
		v0 := b.NewValue0(v.Pos, OpNeq32, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpNeq32, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v4.AddArg(x)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v5.AddArg(y)
		v3.AddArg(v5)
		v.AddArg(v3)
		return true
	}
}
func rewriteValuedec64_OpOr64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v4.AddArg(x)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v5.AddArg(y)
		v3.AddArg(v5)
		v.AddArg(v3)
		return true
	}
}
func rewriteValuedec64_OpRsh16Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpRsh16Ux32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpRsh16Ux32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpSignmask)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpRsh16x32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpRsh16x32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh32Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpRsh32Ux32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpRsh32Ux32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpSignmask)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpRsh32x32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpRsh32x32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh64Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		hi := v_0.Args[0]
		lo := v_0.Args[1]
		s := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpRsh32Ux16, typ.UInt32)
		v0.AddArg(hi)
		v0.AddArg(s)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpRsh32Ux16, typ.UInt32)
		v3.AddArg(lo)
		v3.AddArg(s)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpLsh32x16, typ.UInt32)
		v4.AddArg(hi)
		v5 := b.NewValue0(v.Pos, OpSub16, typ.UInt16)
		v6 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
		v6.AuxInt = 32
		v5.AddArg(v6)
		v5.AddArg(s)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v1.AddArg(v2)
		v7 := b.NewValue0(v.Pos, OpRsh32Ux16, typ.UInt32)
		v7.AddArg(hi)
		v8 := b.NewValue0(v.Pos, OpSub16, typ.UInt16)
		v8.AddArg(s)
		v9 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
		v9.AuxInt = 32
		v8.AddArg(v9)
		v7.AddArg(v8)
		v1.AddArg(v7)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh64Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		hi := v_0.Args[0]
		lo := v_0.Args[1]
		s := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpRsh32Ux32, typ.UInt32)
		v0.AddArg(hi)
		v0.AddArg(s)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpRsh32Ux32, typ.UInt32)
		v3.AddArg(lo)
		v3.AddArg(s)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpLsh32x32, typ.UInt32)
		v4.AddArg(hi)
		v5 := b.NewValue0(v.Pos, OpSub32, typ.UInt32)
		v6 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v6.AuxInt = 32
		v5.AddArg(v6)
		v5.AddArg(s)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v1.AddArg(v2)
		v7 := b.NewValue0(v.Pos, OpRsh32Ux32, typ.UInt32)
		v7.AddArg(hi)
		v8 := b.NewValue0(v.Pos, OpSub32, typ.UInt32)
		v8.AddArg(s)
		v9 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v9.AuxInt = 32
		v8.AddArg(v9)
		v7.AddArg(v8)
		v1.AddArg(v7)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh64Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpRsh64Ux32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpRsh64Ux32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh64Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		hi := v_0.Args[0]
		lo := v_0.Args[1]
		s := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpRsh32Ux8, typ.UInt32)
		v0.AddArg(hi)
		v0.AddArg(s)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpRsh32Ux8, typ.UInt32)
		v3.AddArg(lo)
		v3.AddArg(s)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpLsh32x8, typ.UInt32)
		v4.AddArg(hi)
		v5 := b.NewValue0(v.Pos, OpSub8, typ.UInt8)
		v6 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
		v6.AuxInt = 32
		v5.AddArg(v6)
		v5.AddArg(s)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v1.AddArg(v2)
		v7 := b.NewValue0(v.Pos, OpRsh32Ux8, typ.UInt32)
		v7.AddArg(hi)
		v8 := b.NewValue0(v.Pos, OpSub8, typ.UInt8)
		v8.AddArg(s)
		v9 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
		v9.AuxInt = 32
		v8.AddArg(v9)
		v7.AddArg(v8)
		v1.AddArg(v7)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		hi := v_0.Args[0]
		lo := v_0.Args[1]
		s := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpRsh32x16, typ.UInt32)
		v0.AddArg(hi)
		v0.AddArg(s)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpRsh32Ux16, typ.UInt32)
		v3.AddArg(lo)
		v3.AddArg(s)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpLsh32x16, typ.UInt32)
		v4.AddArg(hi)
		v5 := b.NewValue0(v.Pos, OpSub16, typ.UInt16)
		v6 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
		v6.AuxInt = 32
		v5.AddArg(v6)
		v5.AddArg(s)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v1.AddArg(v2)
		v7 := b.NewValue0(v.Pos, OpAnd32, typ.UInt32)
		v8 := b.NewValue0(v.Pos, OpRsh32x16, typ.UInt32)
		v8.AddArg(hi)
		v9 := b.NewValue0(v.Pos, OpSub16, typ.UInt16)
		v9.AddArg(s)
		v10 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
		v10.AuxInt = 32
		v9.AddArg(v10)
		v8.AddArg(v9)
		v7.AddArg(v8)
		v11 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v12 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v13 := b.NewValue0(v.Pos, OpRsh16Ux32, typ.UInt16)
		v13.AddArg(s)
		v14 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v14.AuxInt = 5
		v13.AddArg(v14)
		v12.AddArg(v13)
		v11.AddArg(v12)
		v7.AddArg(v11)
		v1.AddArg(v7)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		hi := v_0.Args[0]
		lo := v_0.Args[1]
		s := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpRsh32x32, typ.UInt32)
		v0.AddArg(hi)
		v0.AddArg(s)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpRsh32Ux32, typ.UInt32)
		v3.AddArg(lo)
		v3.AddArg(s)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpLsh32x32, typ.UInt32)
		v4.AddArg(hi)
		v5 := b.NewValue0(v.Pos, OpSub32, typ.UInt32)
		v6 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v6.AuxInt = 32
		v5.AddArg(v6)
		v5.AddArg(s)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v1.AddArg(v2)
		v7 := b.NewValue0(v.Pos, OpAnd32, typ.UInt32)
		v8 := b.NewValue0(v.Pos, OpRsh32x32, typ.UInt32)
		v8.AddArg(hi)
		v9 := b.NewValue0(v.Pos, OpSub32, typ.UInt32)
		v9.AddArg(s)
		v10 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v10.AuxInt = 32
		v9.AddArg(v10)
		v8.AddArg(v9)
		v7.AddArg(v8)
		v11 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v12 := b.NewValue0(v.Pos, OpRsh32Ux32, typ.UInt32)
		v12.AddArg(s)
		v13 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v13.AuxInt = 5
		v12.AddArg(v13)
		v11.AddArg(v12)
		v7.AddArg(v11)
		v1.AddArg(v7)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v3 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v3.AddArg(x)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpRsh64x32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpRsh64x32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		hi := v_0.Args[0]
		lo := v_0.Args[1]
		s := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpRsh32x8, typ.UInt32)
		v0.AddArg(hi)
		v0.AddArg(s)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpRsh32Ux8, typ.UInt32)
		v3.AddArg(lo)
		v3.AddArg(s)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpLsh32x8, typ.UInt32)
		v4.AddArg(hi)
		v5 := b.NewValue0(v.Pos, OpSub8, typ.UInt8)
		v6 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
		v6.AuxInt = 32
		v5.AddArg(v6)
		v5.AddArg(s)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v1.AddArg(v2)
		v7 := b.NewValue0(v.Pos, OpAnd32, typ.UInt32)
		v8 := b.NewValue0(v.Pos, OpRsh32x8, typ.UInt32)
		v8.AddArg(hi)
		v9 := b.NewValue0(v.Pos, OpSub8, typ.UInt8)
		v9.AddArg(s)
		v10 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
		v10.AuxInt = 32
		v9.AddArg(v10)
		v8.AddArg(v9)
		v7.AddArg(v8)
		v11 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v12 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v13 := b.NewValue0(v.Pos, OpRsh8Ux32, typ.UInt8)
		v13.AddArg(s)
		v14 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v14.AuxInt = 5
		v13.AddArg(v14)
		v12.AddArg(v13)
		v11.AddArg(v12)
		v7.AddArg(v11)
		v1.AddArg(v7)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh8Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpRsh8Ux32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpRsh8Ux32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpRsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c := v_1_0.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpSignmask)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.AuxInt != 0 {
			break
		}
		lo := v_1.Args[1]
		v.reset(OpRsh8x32)
		v.AddArg(x)
		v.AddArg(lo)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		if !(hi.Op != OpConst32) {
			break
		}
		v.reset(OpRsh8x32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpOr32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeromask, typ.UInt32)
		v1.AddArg(hi)
		v0.AddArg(v1)
		v0.AddArg(lo)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuedec64_OpSignExt16to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpSignExt32to64)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuedec64_OpSignExt32to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValuedec64_OpSignExt8to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpSignExt32to64)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuedec64_OpStore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 8 && !config.BigEndian) {
			break
		}
		v.reset(OpStore)
		v.Aux = hi.Type
		v0 := b.NewValue0(v.Pos, OpOffPtr, hi.Type.PtrTo(psess.types))
		v0.AuxInt = 4
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(hi)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = lo.Type
		v1.AddArg(dst)
		v1.AddArg(lo)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpInt64Make {
			break
		}
		_ = v_1.Args[1]
		hi := v_1.Args[0]
		lo := v_1.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 8 && config.BigEndian) {
			break
		}
		v.reset(OpStore)
		v.Aux = lo.Type
		v0 := b.NewValue0(v.Pos, OpOffPtr, lo.Type.PtrTo(psess.types))
		v0.AuxInt = 4
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(lo)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = hi.Type
		v1.AddArg(dst)
		v1.AddArg(hi)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuedec64_OpSub64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpSub32withcarry, typ.Int32)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSelect1, psess.types.TypeFlags)
		v4 := b.NewValue0(v.Pos, OpSub32carry, types.NewTuple(typ.UInt32, psess.types.TypeFlags))
		v5 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v6.AddArg(y)
		v4.AddArg(v6)
		v3.AddArg(v4)
		v0.AddArg(v3)
		v.AddArg(v0)
		v7 := b.NewValue0(v.Pos, OpSelect0, typ.UInt32)
		v8 := b.NewValue0(v.Pos, OpSub32carry, types.NewTuple(typ.UInt32, psess.types.TypeFlags))
		v9 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v9.AddArg(x)
		v8.AddArg(v9)
		v10 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v10.AddArg(y)
		v8.AddArg(v10)
		v7.AddArg(v8)
		v.AddArg(v7)
		return true
	}
}
func rewriteValuedec64_OpTrunc64to16_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		lo := v_0.Args[1]
		v.reset(OpTrunc32to16)
		v.AddArg(lo)
		return true
	}
	return false
}
func rewriteValuedec64_OpTrunc64to32_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		lo := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = lo.Type
		v.AddArg(lo)
		return true
	}
	return false
}
func rewriteValuedec64_OpTrunc64to8_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpInt64Make {
			break
		}
		_ = v_0.Args[1]
		lo := v_0.Args[1]
		v.reset(OpTrunc32to8)
		v.AddArg(lo)
		return true
	}
	return false
}
func rewriteValuedec64_OpXor64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpXor32, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpInt64Hi, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpXor32, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v4.AddArg(x)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpInt64Lo, typ.UInt32)
		v5.AddArg(y)
		v3.AddArg(v5)
		v.AddArg(v3)
		return true
	}
}
func rewriteValuedec64_OpZeroExt16to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpZeroExt32to64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuedec64_OpZeroExt32to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpInt64Make)
		v0 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValuedec64_OpZeroExt8to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpZeroExt32to64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteBlockdec64(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	}
	return false
}

package ssa

import "math"

import "github.com/dave/golib/src/cmd/compile/internal/types"

// in case not otherwise used
// in case not otherwise used
// in case not otherwise used
// in case not otherwise used

func (psess *PackageSession) rewriteValuegeneric(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValuegeneric_OpAdd16_0(v) || rewriteValuegeneric_OpAdd16_10(v) || rewriteValuegeneric_OpAdd16_20(v) || rewriteValuegeneric_OpAdd16_30(v)
	case OpAdd32:
		return rewriteValuegeneric_OpAdd32_0(v) || rewriteValuegeneric_OpAdd32_10(v) || rewriteValuegeneric_OpAdd32_20(v) || rewriteValuegeneric_OpAdd32_30(v)
	case OpAdd32F:
		return rewriteValuegeneric_OpAdd32F_0(v)
	case OpAdd64:
		return rewriteValuegeneric_OpAdd64_0(v) || rewriteValuegeneric_OpAdd64_10(v) || rewriteValuegeneric_OpAdd64_20(v) || rewriteValuegeneric_OpAdd64_30(v)
	case OpAdd64F:
		return rewriteValuegeneric_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValuegeneric_OpAdd8_0(v) || rewriteValuegeneric_OpAdd8_10(v) || rewriteValuegeneric_OpAdd8_20(v) || rewriteValuegeneric_OpAdd8_30(v)
	case OpAddPtr:
		return rewriteValuegeneric_OpAddPtr_0(v)
	case OpAnd16:
		return rewriteValuegeneric_OpAnd16_0(v) || rewriteValuegeneric_OpAnd16_10(v) || rewriteValuegeneric_OpAnd16_20(v)
	case OpAnd32:
		return rewriteValuegeneric_OpAnd32_0(v) || rewriteValuegeneric_OpAnd32_10(v) || rewriteValuegeneric_OpAnd32_20(v)
	case OpAnd64:
		return rewriteValuegeneric_OpAnd64_0(v) || rewriteValuegeneric_OpAnd64_10(v) || rewriteValuegeneric_OpAnd64_20(v)
	case OpAnd8:
		return rewriteValuegeneric_OpAnd8_0(v) || rewriteValuegeneric_OpAnd8_10(v) || rewriteValuegeneric_OpAnd8_20(v)
	case OpArg:
		return psess.rewriteValuegeneric_OpArg_0(v) || psess.rewriteValuegeneric_OpArg_10(v)
	case OpArraySelect:
		return rewriteValuegeneric_OpArraySelect_0(v)
	case OpCom16:
		return rewriteValuegeneric_OpCom16_0(v)
	case OpCom32:
		return rewriteValuegeneric_OpCom32_0(v)
	case OpCom64:
		return rewriteValuegeneric_OpCom64_0(v)
	case OpCom8:
		return rewriteValuegeneric_OpCom8_0(v)
	case OpConstInterface:
		return rewriteValuegeneric_OpConstInterface_0(v)
	case OpConstSlice:
		return psess.rewriteValuegeneric_OpConstSlice_0(v)
	case OpConstString:
		return rewriteValuegeneric_OpConstString_0(v)
	case OpConvert:
		return rewriteValuegeneric_OpConvert_0(v)
	case OpCvt32Fto32:
		return rewriteValuegeneric_OpCvt32Fto32_0(v)
	case OpCvt32Fto64:
		return rewriteValuegeneric_OpCvt32Fto64_0(v)
	case OpCvt32Fto64F:
		return rewriteValuegeneric_OpCvt32Fto64F_0(v)
	case OpCvt32to32F:
		return rewriteValuegeneric_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValuegeneric_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValuegeneric_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValuegeneric_OpCvt64Fto32F_0(v)
	case OpCvt64Fto64:
		return rewriteValuegeneric_OpCvt64Fto64_0(v)
	case OpCvt64to32F:
		return rewriteValuegeneric_OpCvt64to32F_0(v)
	case OpCvt64to64F:
		return rewriteValuegeneric_OpCvt64to64F_0(v)
	case OpDiv16:
		return rewriteValuegeneric_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValuegeneric_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValuegeneric_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValuegeneric_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValuegeneric_OpDiv32u_0(v)
	case OpDiv64:
		return rewriteValuegeneric_OpDiv64_0(v)
	case OpDiv64F:
		return rewriteValuegeneric_OpDiv64F_0(v)
	case OpDiv64u:
		return rewriteValuegeneric_OpDiv64u_0(v)
	case OpDiv8:
		return rewriteValuegeneric_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValuegeneric_OpDiv8u_0(v)
	case OpEq16:
		return rewriteValuegeneric_OpEq16_0(v)
	case OpEq32:
		return rewriteValuegeneric_OpEq32_0(v)
	case OpEq32F:
		return rewriteValuegeneric_OpEq32F_0(v)
	case OpEq64:
		return rewriteValuegeneric_OpEq64_0(v)
	case OpEq64F:
		return rewriteValuegeneric_OpEq64F_0(v)
	case OpEq8:
		return rewriteValuegeneric_OpEq8_0(v)
	case OpEqB:
		return rewriteValuegeneric_OpEqB_0(v)
	case OpEqInter:
		return rewriteValuegeneric_OpEqInter_0(v)
	case OpEqPtr:
		return rewriteValuegeneric_OpEqPtr_0(v) || rewriteValuegeneric_OpEqPtr_10(v)
	case OpEqSlice:
		return rewriteValuegeneric_OpEqSlice_0(v)
	case OpGeq16:
		return rewriteValuegeneric_OpGeq16_0(v)
	case OpGeq16U:
		return rewriteValuegeneric_OpGeq16U_0(v)
	case OpGeq32:
		return rewriteValuegeneric_OpGeq32_0(v)
	case OpGeq32F:
		return rewriteValuegeneric_OpGeq32F_0(v)
	case OpGeq32U:
		return rewriteValuegeneric_OpGeq32U_0(v)
	case OpGeq64:
		return rewriteValuegeneric_OpGeq64_0(v)
	case OpGeq64F:
		return rewriteValuegeneric_OpGeq64F_0(v)
	case OpGeq64U:
		return rewriteValuegeneric_OpGeq64U_0(v)
	case OpGeq8:
		return rewriteValuegeneric_OpGeq8_0(v)
	case OpGeq8U:
		return rewriteValuegeneric_OpGeq8U_0(v)
	case OpGreater16:
		return rewriteValuegeneric_OpGreater16_0(v)
	case OpGreater16U:
		return rewriteValuegeneric_OpGreater16U_0(v)
	case OpGreater32:
		return rewriteValuegeneric_OpGreater32_0(v)
	case OpGreater32F:
		return rewriteValuegeneric_OpGreater32F_0(v)
	case OpGreater32U:
		return rewriteValuegeneric_OpGreater32U_0(v)
	case OpGreater64:
		return rewriteValuegeneric_OpGreater64_0(v)
	case OpGreater64F:
		return rewriteValuegeneric_OpGreater64F_0(v)
	case OpGreater64U:
		return rewriteValuegeneric_OpGreater64U_0(v)
	case OpGreater8:
		return rewriteValuegeneric_OpGreater8_0(v)
	case OpGreater8U:
		return rewriteValuegeneric_OpGreater8U_0(v)
	case OpIMake:
		return rewriteValuegeneric_OpIMake_0(v)
	case OpInterCall:
		return rewriteValuegeneric_OpInterCall_0(v)
	case OpIsInBounds:
		return rewriteValuegeneric_OpIsInBounds_0(v) || rewriteValuegeneric_OpIsInBounds_10(v) || rewriteValuegeneric_OpIsInBounds_20(v) || rewriteValuegeneric_OpIsInBounds_30(v)
	case OpIsNonNil:
		return rewriteValuegeneric_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return rewriteValuegeneric_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return rewriteValuegeneric_OpLeq16_0(v)
	case OpLeq16U:
		return rewriteValuegeneric_OpLeq16U_0(v)
	case OpLeq32:
		return rewriteValuegeneric_OpLeq32_0(v)
	case OpLeq32F:
		return rewriteValuegeneric_OpLeq32F_0(v)
	case OpLeq32U:
		return rewriteValuegeneric_OpLeq32U_0(v)
	case OpLeq64:
		return rewriteValuegeneric_OpLeq64_0(v)
	case OpLeq64F:
		return rewriteValuegeneric_OpLeq64F_0(v)
	case OpLeq64U:
		return rewriteValuegeneric_OpLeq64U_0(v)
	case OpLeq8:
		return rewriteValuegeneric_OpLeq8_0(v)
	case OpLeq8U:
		return rewriteValuegeneric_OpLeq8U_0(v)
	case OpLess16:
		return rewriteValuegeneric_OpLess16_0(v)
	case OpLess16U:
		return rewriteValuegeneric_OpLess16U_0(v)
	case OpLess32:
		return rewriteValuegeneric_OpLess32_0(v)
	case OpLess32F:
		return rewriteValuegeneric_OpLess32F_0(v)
	case OpLess32U:
		return rewriteValuegeneric_OpLess32U_0(v)
	case OpLess64:
		return rewriteValuegeneric_OpLess64_0(v)
	case OpLess64F:
		return rewriteValuegeneric_OpLess64F_0(v)
	case OpLess64U:
		return rewriteValuegeneric_OpLess64U_0(v)
	case OpLess8:
		return rewriteValuegeneric_OpLess8_0(v)
	case OpLess8U:
		return rewriteValuegeneric_OpLess8U_0(v)
	case OpLoad:
		return psess.rewriteValuegeneric_OpLoad_0(v) || psess.rewriteValuegeneric_OpLoad_10(v) || psess.rewriteValuegeneric_OpLoad_20(v)
	case OpLsh16x16:
		return rewriteValuegeneric_OpLsh16x16_0(v)
	case OpLsh16x32:
		return rewriteValuegeneric_OpLsh16x32_0(v)
	case OpLsh16x64:
		return rewriteValuegeneric_OpLsh16x64_0(v)
	case OpLsh16x8:
		return rewriteValuegeneric_OpLsh16x8_0(v)
	case OpLsh32x16:
		return rewriteValuegeneric_OpLsh32x16_0(v)
	case OpLsh32x32:
		return rewriteValuegeneric_OpLsh32x32_0(v)
	case OpLsh32x64:
		return rewriteValuegeneric_OpLsh32x64_0(v)
	case OpLsh32x8:
		return rewriteValuegeneric_OpLsh32x8_0(v)
	case OpLsh64x16:
		return rewriteValuegeneric_OpLsh64x16_0(v)
	case OpLsh64x32:
		return rewriteValuegeneric_OpLsh64x32_0(v)
	case OpLsh64x64:
		return rewriteValuegeneric_OpLsh64x64_0(v)
	case OpLsh64x8:
		return rewriteValuegeneric_OpLsh64x8_0(v)
	case OpLsh8x16:
		return rewriteValuegeneric_OpLsh8x16_0(v)
	case OpLsh8x32:
		return rewriteValuegeneric_OpLsh8x32_0(v)
	case OpLsh8x64:
		return rewriteValuegeneric_OpLsh8x64_0(v)
	case OpLsh8x8:
		return rewriteValuegeneric_OpLsh8x8_0(v)
	case OpMod16:
		return rewriteValuegeneric_OpMod16_0(v)
	case OpMod16u:
		return rewriteValuegeneric_OpMod16u_0(v)
	case OpMod32:
		return rewriteValuegeneric_OpMod32_0(v)
	case OpMod32u:
		return rewriteValuegeneric_OpMod32u_0(v)
	case OpMod64:
		return rewriteValuegeneric_OpMod64_0(v)
	case OpMod64u:
		return rewriteValuegeneric_OpMod64u_0(v)
	case OpMod8:
		return rewriteValuegeneric_OpMod8_0(v)
	case OpMod8u:
		return rewriteValuegeneric_OpMod8u_0(v)
	case OpMove:
		return psess.rewriteValuegeneric_OpMove_0(v) || psess.rewriteValuegeneric_OpMove_10(v) || psess.rewriteValuegeneric_OpMove_20(v)
	case OpMul16:
		return rewriteValuegeneric_OpMul16_0(v) || rewriteValuegeneric_OpMul16_10(v)
	case OpMul32:
		return rewriteValuegeneric_OpMul32_0(v) || rewriteValuegeneric_OpMul32_10(v)
	case OpMul32F:
		return rewriteValuegeneric_OpMul32F_0(v)
	case OpMul64:
		return rewriteValuegeneric_OpMul64_0(v) || rewriteValuegeneric_OpMul64_10(v)
	case OpMul64F:
		return rewriteValuegeneric_OpMul64F_0(v)
	case OpMul8:
		return rewriteValuegeneric_OpMul8_0(v) || rewriteValuegeneric_OpMul8_10(v)
	case OpNeg16:
		return rewriteValuegeneric_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValuegeneric_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValuegeneric_OpNeg32F_0(v)
	case OpNeg64:
		return rewriteValuegeneric_OpNeg64_0(v)
	case OpNeg64F:
		return rewriteValuegeneric_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValuegeneric_OpNeg8_0(v)
	case OpNeq16:
		return rewriteValuegeneric_OpNeq16_0(v)
	case OpNeq32:
		return rewriteValuegeneric_OpNeq32_0(v)
	case OpNeq32F:
		return rewriteValuegeneric_OpNeq32F_0(v)
	case OpNeq64:
		return rewriteValuegeneric_OpNeq64_0(v)
	case OpNeq64F:
		return rewriteValuegeneric_OpNeq64F_0(v)
	case OpNeq8:
		return rewriteValuegeneric_OpNeq8_0(v)
	case OpNeqB:
		return rewriteValuegeneric_OpNeqB_0(v)
	case OpNeqInter:
		return rewriteValuegeneric_OpNeqInter_0(v)
	case OpNeqPtr:
		return rewriteValuegeneric_OpNeqPtr_0(v) || rewriteValuegeneric_OpNeqPtr_10(v)
	case OpNeqSlice:
		return rewriteValuegeneric_OpNeqSlice_0(v)
	case OpNilCheck:
		return rewriteValuegeneric_OpNilCheck_0(v)
	case OpNot:
		return rewriteValuegeneric_OpNot_0(v) || rewriteValuegeneric_OpNot_10(v) || rewriteValuegeneric_OpNot_20(v) || rewriteValuegeneric_OpNot_30(v) || rewriteValuegeneric_OpNot_40(v)
	case OpOffPtr:
		return psess.rewriteValuegeneric_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValuegeneric_OpOr16_0(v) || rewriteValuegeneric_OpOr16_10(v) || rewriteValuegeneric_OpOr16_20(v)
	case OpOr32:
		return rewriteValuegeneric_OpOr32_0(v) || rewriteValuegeneric_OpOr32_10(v) || rewriteValuegeneric_OpOr32_20(v)
	case OpOr64:
		return rewriteValuegeneric_OpOr64_0(v) || rewriteValuegeneric_OpOr64_10(v) || rewriteValuegeneric_OpOr64_20(v)
	case OpOr8:
		return rewriteValuegeneric_OpOr8_0(v) || rewriteValuegeneric_OpOr8_10(v) || rewriteValuegeneric_OpOr8_20(v)
	case OpPhi:
		return rewriteValuegeneric_OpPhi_0(v)
	case OpPtrIndex:
		return psess.rewriteValuegeneric_OpPtrIndex_0(v)
	case OpRound32F:
		return rewriteValuegeneric_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValuegeneric_OpRound64F_0(v)
	case OpRsh16Ux16:
		return rewriteValuegeneric_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return rewriteValuegeneric_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return rewriteValuegeneric_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return rewriteValuegeneric_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return rewriteValuegeneric_OpRsh16x16_0(v)
	case OpRsh16x32:
		return rewriteValuegeneric_OpRsh16x32_0(v)
	case OpRsh16x64:
		return rewriteValuegeneric_OpRsh16x64_0(v)
	case OpRsh16x8:
		return rewriteValuegeneric_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return rewriteValuegeneric_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return rewriteValuegeneric_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return rewriteValuegeneric_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return rewriteValuegeneric_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return rewriteValuegeneric_OpRsh32x16_0(v)
	case OpRsh32x32:
		return rewriteValuegeneric_OpRsh32x32_0(v)
	case OpRsh32x64:
		return rewriteValuegeneric_OpRsh32x64_0(v)
	case OpRsh32x8:
		return rewriteValuegeneric_OpRsh32x8_0(v)
	case OpRsh64Ux16:
		return rewriteValuegeneric_OpRsh64Ux16_0(v)
	case OpRsh64Ux32:
		return rewriteValuegeneric_OpRsh64Ux32_0(v)
	case OpRsh64Ux64:
		return rewriteValuegeneric_OpRsh64Ux64_0(v)
	case OpRsh64Ux8:
		return rewriteValuegeneric_OpRsh64Ux8_0(v)
	case OpRsh64x16:
		return rewriteValuegeneric_OpRsh64x16_0(v)
	case OpRsh64x32:
		return rewriteValuegeneric_OpRsh64x32_0(v)
	case OpRsh64x64:
		return rewriteValuegeneric_OpRsh64x64_0(v)
	case OpRsh64x8:
		return rewriteValuegeneric_OpRsh64x8_0(v)
	case OpRsh8Ux16:
		return rewriteValuegeneric_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return rewriteValuegeneric_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return rewriteValuegeneric_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return rewriteValuegeneric_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return rewriteValuegeneric_OpRsh8x16_0(v)
	case OpRsh8x32:
		return rewriteValuegeneric_OpRsh8x32_0(v)
	case OpRsh8x64:
		return rewriteValuegeneric_OpRsh8x64_0(v)
	case OpRsh8x8:
		return rewriteValuegeneric_OpRsh8x8_0(v)
	case OpSignExt16to32:
		return rewriteValuegeneric_OpSignExt16to32_0(v)
	case OpSignExt16to64:
		return rewriteValuegeneric_OpSignExt16to64_0(v)
	case OpSignExt32to64:
		return rewriteValuegeneric_OpSignExt32to64_0(v)
	case OpSignExt8to16:
		return rewriteValuegeneric_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValuegeneric_OpSignExt8to32_0(v)
	case OpSignExt8to64:
		return rewriteValuegeneric_OpSignExt8to64_0(v)
	case OpSliceCap:
		return rewriteValuegeneric_OpSliceCap_0(v)
	case OpSliceLen:
		return rewriteValuegeneric_OpSliceLen_0(v)
	case OpSlicePtr:
		return rewriteValuegeneric_OpSlicePtr_0(v)
	case OpSlicemask:
		return rewriteValuegeneric_OpSlicemask_0(v)
	case OpSqrt:
		return rewriteValuegeneric_OpSqrt_0(v)
	case OpStaticCall:
		return psess.rewriteValuegeneric_OpStaticCall_0(v)
	case OpStore:
		return psess.rewriteValuegeneric_OpStore_0(v) || psess.rewriteValuegeneric_OpStore_10(v) || psess.rewriteValuegeneric_OpStore_20(v)
	case OpStringLen:
		return rewriteValuegeneric_OpStringLen_0(v)
	case OpStringPtr:
		return rewriteValuegeneric_OpStringPtr_0(v)
	case OpStructSelect:
		return rewriteValuegeneric_OpStructSelect_0(v) || psess.rewriteValuegeneric_OpStructSelect_10(v)
	case OpSub16:
		return rewriteValuegeneric_OpSub16_0(v) || rewriteValuegeneric_OpSub16_10(v)
	case OpSub32:
		return rewriteValuegeneric_OpSub32_0(v) || rewriteValuegeneric_OpSub32_10(v)
	case OpSub32F:
		return rewriteValuegeneric_OpSub32F_0(v)
	case OpSub64:
		return rewriteValuegeneric_OpSub64_0(v) || rewriteValuegeneric_OpSub64_10(v)
	case OpSub64F:
		return rewriteValuegeneric_OpSub64F_0(v)
	case OpSub8:
		return rewriteValuegeneric_OpSub8_0(v) || rewriteValuegeneric_OpSub8_10(v)
	case OpTrunc16to8:
		return rewriteValuegeneric_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValuegeneric_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValuegeneric_OpTrunc32to8_0(v)
	case OpTrunc64to16:
		return rewriteValuegeneric_OpTrunc64to16_0(v)
	case OpTrunc64to32:
		return rewriteValuegeneric_OpTrunc64to32_0(v)
	case OpTrunc64to8:
		return rewriteValuegeneric_OpTrunc64to8_0(v)
	case OpXor16:
		return rewriteValuegeneric_OpXor16_0(v) || rewriteValuegeneric_OpXor16_10(v)
	case OpXor32:
		return rewriteValuegeneric_OpXor32_0(v) || rewriteValuegeneric_OpXor32_10(v)
	case OpXor64:
		return rewriteValuegeneric_OpXor64_0(v) || rewriteValuegeneric_OpXor64_10(v)
	case OpXor8:
		return rewriteValuegeneric_OpXor8_0(v) || rewriteValuegeneric_OpXor8_10(v)
	case OpZero:
		return psess.rewriteValuegeneric_OpZero_0(v)
	case OpZeroExt16to32:
		return rewriteValuegeneric_OpZeroExt16to32_0(v)
	case OpZeroExt16to64:
		return rewriteValuegeneric_OpZeroExt16to64_0(v)
	case OpZeroExt32to64:
		return rewriteValuegeneric_OpZeroExt32to64_0(v)
	case OpZeroExt8to16:
		return rewriteValuegeneric_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValuegeneric_OpZeroExt8to32_0(v)
	case OpZeroExt8to64:
		return rewriteValuegeneric_OpZeroExt8to64_0(v)
	}
	return false
}
func rewriteValuegeneric_OpAdd16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c + d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c + d))
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		z := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		z := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd16_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpCom16 {
			break
		}
		x := v_1.Args[0]
		v.reset(OpNeg16)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpCom16 {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpNeg16)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAdd16 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAdd16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub16 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub16, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub16, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd16_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub16, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub16 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub16, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub16 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub16 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd16 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst16 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd16 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst16 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst16 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst16 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd16_30(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst16 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst16 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst16 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst16 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c + d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c + d))
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		z := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		z := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd32_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpCom32 {
			break
		}
		x := v_1.Args[0]
		v.reset(OpNeg32)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpCom32 {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpNeg32)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAdd32 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAdd32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub32 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub32, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub32, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd32_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub32, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub32 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub32, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub32 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub32 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd32 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd32 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd32_30(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = f2i(float64(i2f32(c) + i2f32(d)))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = f2i(float64(i2f32(c) + i2f32(d)))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c + d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c + d
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		z := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		z := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd64_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpCom64 {
			break
		}
		x := v_1.Args[0]
		v.reset(OpNeg64)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpCom64 {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpNeg64)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAdd64 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAdd64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub64 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub64, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub64, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd64_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub64, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub64 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub64, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub64 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub64 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd64 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd64 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd64_30(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = f2i(i2f(c) + i2f(d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = f2i(i2f(c) + i2f(d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c + d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c + d))
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		z := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		z := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd8_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpCom8 {
			break
		}
		x := v_1.Args[0]
		v.reset(OpNeg8)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpCom8 {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpNeg8)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAdd8 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAdd8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub8 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub8, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub8, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd8_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub8, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub8 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub8, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub8 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub8 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd8 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst8 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd8 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst8 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst8 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst8 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAdd8_30(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst8 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst8 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst8 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSub8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst8 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAddPtr_0(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOffPtr)
		v.Type = t
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOffPtr)
		v.Type = t
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c & d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c & d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		m := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpRsh16Ux64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c >= 64-ntz(m)) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		m := v_1.AuxInt
		if !(c >= 64-ntz(m)) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		m := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpLsh16x64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c >= 64-nlz(m)) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh16x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		m := v_1.AuxInt
		if !(c >= 64-nlz(m)) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd16_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd16 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpAnd16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd16 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpAnd16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpAnd16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpAnd16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAnd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAnd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd16 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAnd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAnd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd16 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst16 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAnd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd16_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd16 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst16 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAnd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst16 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAnd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst16 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAnd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c & d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c & d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		m := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpRsh32Ux64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c >= 64-ntz(m)) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh32Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		m := v_1.AuxInt
		if !(c >= 64-ntz(m)) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		m := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpLsh32x64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c >= 64-nlz(m)) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		m := v_1.AuxInt
		if !(c >= 64-nlz(m)) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd32_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != 0 {
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
		if v_1.Op != OpAnd32 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpAnd32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd32 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpAnd32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpAnd32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpAnd32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAnd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAnd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd32 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAnd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAnd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd32 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAnd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd32_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd32 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAnd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAnd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAnd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c & d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c & d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		m := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpRsh64Ux64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c >= 64-ntz(m)) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh64Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		m := v_1.AuxInt
		if !(c >= 64-ntz(m)) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		m := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpLsh64x64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c >= 64-nlz(m)) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		m := v_1.AuxInt
		if !(c >= 64-nlz(m)) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd64_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
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
		if v_1.Op != OpAnd64 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpAnd64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd64 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpAnd64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpAnd64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpAnd64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		y := v_0.AuxInt
		x := v.Args[1]
		if !(nlz(y)+nto(y) == 64 && nto(y) >= 32) {
			break
		}
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpLsh64x64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpConst64, t)
		v1.AuxInt = nlz(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = nlz(y)
		v.AddArg(v2)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		y := v_1.AuxInt
		if !(nlz(y)+nto(y) == 64 && nto(y) >= 32) {
			break
		}
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpLsh64x64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpConst64, t)
		v1.AuxInt = nlz(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = nlz(y)
		v.AddArg(v2)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		y := v_0.AuxInt
		x := v.Args[1]
		if !(nlo(y)+ntz(y) == 64 && ntz(y) >= 32) {
			break
		}
		v.reset(OpLsh64x64)
		v0 := b.NewValue0(v.Pos, OpRsh64Ux64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpConst64, t)
		v1.AuxInt = ntz(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = ntz(y)
		v.AddArg(v2)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		y := v_1.AuxInt
		if !(nlo(y)+ntz(y) == 64 && ntz(y) >= 32) {
			break
		}
		v.reset(OpLsh64x64)
		v0 := b.NewValue0(v.Pos, OpRsh64Ux64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpConst64, t)
		v1.AuxInt = ntz(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = ntz(y)
		v.AddArg(v2)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAnd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd64_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAnd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd64 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAnd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAnd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd64 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAnd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c & d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd64 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAnd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c & d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAnd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c & d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAnd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c & d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c & d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c & d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		m := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpRsh8Ux64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c >= 64-ntz(m)) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		m := v_1.AuxInt
		if !(c >= 64-ntz(m)) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		m := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpLsh8x64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c >= 64-nlz(m)) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh8x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		m := v_1.AuxInt
		if !(c >= 64-nlz(m)) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd8_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd8 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpAnd8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd8 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpAnd8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpAnd8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd8 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpAnd8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd8 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAnd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd8 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAnd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd8 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAnd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAnd8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAnd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpAnd8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd8 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst8 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAnd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd8_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd8 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst8 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAnd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst8 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAnd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst8 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAnd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c & d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpArg_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &b.Func.Config.Types
	_ = typ

	for {
		off := v.AuxInt
		n := v.Aux
		if !(v.Type.IsString()) {
			break
		}
		v.reset(OpStringMake)
		v0 := b.NewValue0(v.Pos, OpArg, typ.BytePtr)
		v0.AuxInt = off
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, typ.Int)
		v1.AuxInt = off + config.PtrSize
		v1.Aux = n
		v.AddArg(v1)
		return true
	}

	for {
		off := v.AuxInt
		n := v.Aux
		if !(v.Type.IsSlice()) {
			break
		}
		v.reset(OpSliceMake)
		v0 := b.NewValue0(v.Pos, OpArg, v.Type.Elem(psess.types).PtrTo(psess.types))
		v0.AuxInt = off
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, typ.Int)
		v1.AuxInt = off + config.PtrSize
		v1.Aux = n
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpArg, typ.Int)
		v2.AuxInt = off + 2*config.PtrSize
		v2.Aux = n
		v.AddArg(v2)
		return true
	}

	for {
		off := v.AuxInt
		n := v.Aux
		if !(v.Type.IsInterface()) {
			break
		}
		v.reset(OpIMake)
		v0 := b.NewValue0(v.Pos, OpArg, typ.Uintptr)
		v0.AuxInt = off
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, typ.BytePtr)
		v1.AuxInt = off + config.PtrSize
		v1.Aux = n
		v.AddArg(v1)
		return true
	}

	for {
		off := v.AuxInt
		n := v.Aux
		if !(v.Type.IsComplex() && v.Type.Size(psess.types) == 16) {
			break
		}
		v.reset(OpComplexMake)
		v0 := b.NewValue0(v.Pos, OpArg, typ.Float64)
		v0.AuxInt = off
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, typ.Float64)
		v1.AuxInt = off + 8
		v1.Aux = n
		v.AddArg(v1)
		return true
	}

	for {
		off := v.AuxInt
		n := v.Aux
		if !(v.Type.IsComplex() && v.Type.Size(psess.types) == 8) {
			break
		}
		v.reset(OpComplexMake)
		v0 := b.NewValue0(v.Pos, OpArg, typ.Float32)
		v0.AuxInt = off
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, typ.Float32)
		v1.AuxInt = off + 4
		v1.Aux = n
		v.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		if !(t.IsStruct() && t.NumFields(psess.types) == 0 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake0)
		return true
	}

	for {
		t := v.Type
		off := v.AuxInt
		n := v.Aux
		if !(t.IsStruct() && t.NumFields(psess.types) == 1 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake1)
		v0 := b.NewValue0(v.Pos, OpArg, t.FieldType(psess.types, 0))
		v0.AuxInt = off + t.FieldOff(psess.types, 0)
		v0.Aux = n
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		off := v.AuxInt
		n := v.Aux
		if !(t.IsStruct() && t.NumFields(psess.types) == 2 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake2)
		v0 := b.NewValue0(v.Pos, OpArg, t.FieldType(psess.types, 0))
		v0.AuxInt = off + t.FieldOff(psess.types, 0)
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, t.FieldType(psess.types, 1))
		v1.AuxInt = off + t.FieldOff(psess.types, 1)
		v1.Aux = n
		v.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		off := v.AuxInt
		n := v.Aux
		if !(t.IsStruct() && t.NumFields(psess.types) == 3 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake3)
		v0 := b.NewValue0(v.Pos, OpArg, t.FieldType(psess.types, 0))
		v0.AuxInt = off + t.FieldOff(psess.types, 0)
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, t.FieldType(psess.types, 1))
		v1.AuxInt = off + t.FieldOff(psess.types, 1)
		v1.Aux = n
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpArg, t.FieldType(psess.types, 2))
		v2.AuxInt = off + t.FieldOff(psess.types, 2)
		v2.Aux = n
		v.AddArg(v2)
		return true
	}

	for {
		t := v.Type
		off := v.AuxInt
		n := v.Aux
		if !(t.IsStruct() && t.NumFields(psess.types) == 4 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake4)
		v0 := b.NewValue0(v.Pos, OpArg, t.FieldType(psess.types, 0))
		v0.AuxInt = off + t.FieldOff(psess.types, 0)
		v0.Aux = n
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpArg, t.FieldType(psess.types, 1))
		v1.AuxInt = off + t.FieldOff(psess.types, 1)
		v1.Aux = n
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpArg, t.FieldType(psess.types, 2))
		v2.AuxInt = off + t.FieldOff(psess.types, 2)
		v2.Aux = n
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpArg, t.FieldType(psess.types, 3))
		v3.AuxInt = off + t.FieldOff(psess.types, 3)
		v3.Aux = n
		v.AddArg(v3)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpArg_10(v *Value) bool {
	b := v.Block
	_ = b
	fe := b.Func.fe
	_ = fe

	for {
		t := v.Type
		if !(t.IsArray() && t.NumElem(psess.types) == 0) {
			break
		}
		v.reset(OpArrayMake0)
		return true
	}

	for {
		t := v.Type
		off := v.AuxInt
		n := v.Aux
		if !(t.IsArray() && t.NumElem(psess.types) == 1 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpArrayMake1)
		v0 := b.NewValue0(v.Pos, OpArg, t.Elem(psess.types))
		v0.AuxInt = off
		v0.Aux = n
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpArraySelect_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpArrayMake1 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		if x.Op != OpIData {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpCom16_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpCom16 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = ^c
		return true
	}
	return false
}
func rewriteValuegeneric_OpCom32_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpCom32 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = ^c
		return true
	}
	return false
}
func rewriteValuegeneric_OpCom64_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpCom64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = ^c
		return true
	}
	return false
}
func rewriteValuegeneric_OpCom8_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpCom8 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst8)
		v.AuxInt = ^c
		return true
	}
	return false
}
func rewriteValuegeneric_OpConstInterface_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		v.reset(OpIMake)
		v0 := b.NewValue0(v.Pos, OpConstNil, typ.Uintptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConstNil, typ.BytePtr)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuegeneric_OpConstSlice_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		if !(config.PtrSize == 4) {
			break
		}
		v.reset(OpSliceMake)
		v0 := b.NewValue0(v.Pos, OpConstNil, v.Type.Elem(psess.types).PtrTo(psess.types))
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConst32, typ.Int)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.Int)
		v2.AuxInt = 0
		v.AddArg(v2)
		return true
	}

	for {
		if !(config.PtrSize == 8) {
			break
		}
		v.reset(OpSliceMake)
		v0 := b.NewValue0(v.Pos, OpConstNil, v.Type.Elem(psess.types).PtrTo(psess.types))
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.Int)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.Int)
		v2.AuxInt = 0
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValuegeneric_OpConstString_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &b.Func.Config.Types
	_ = typ

	for {
		s := v.Aux
		if !(config.PtrSize == 4 && s.(string) == "") {
			break
		}
		v.reset(OpStringMake)
		v0 := b.NewValue0(v.Pos, OpConstNil, typ.BytePtr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConst32, typ.Int)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}

	for {
		s := v.Aux
		if !(config.PtrSize == 8 && s.(string) == "") {
			break
		}
		v.reset(OpStringMake)
		v0 := b.NewValue0(v.Pos, OpConstNil, typ.BytePtr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.Int)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}

	for {
		s := v.Aux
		if !(config.PtrSize == 4 && s.(string) != "") {
			break
		}
		v.reset(OpStringMake)
		v0 := b.NewValue0(v.Pos, OpAddr, typ.BytePtr)
		v0.Aux = fe.StringData(s.(string))
		v1 := b.NewValue0(v.Pos, OpSB, typ.Uintptr)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.Int)
		v2.AuxInt = int64(len(s.(string)))
		v.AddArg(v2)
		return true
	}

	for {
		s := v.Aux
		if !(config.PtrSize == 8 && s.(string) != "") {
			break
		}
		v.reset(OpStringMake)
		v0 := b.NewValue0(v.Pos, OpAddr, typ.BytePtr)
		v0.Aux = fe.StringData(s.(string))
		v1 := b.NewValue0(v.Pos, OpSB, typ.Uintptr)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.Int)
		v2.AuxInt = int64(len(s.(string)))
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValuegeneric_OpConvert_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConvert {
			break
		}
		_ = v_0_0.Args[1]
		ptr := v_0_0.Args[0]
		mem := v_0_0.Args[1]
		off := v_0.Args[1]
		if mem != v.Args[1] {
			break
		}
		v.reset(OpAdd64)
		v.AddArg(ptr)
		v.AddArg(off)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		off := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConvert {
			break
		}
		_ = v_0_1.Args[1]
		ptr := v_0_1.Args[0]
		mem := v_0_1.Args[1]
		if mem != v.Args[1] {
			break
		}
		v.reset(OpAdd64)
		v.AddArg(ptr)
		v.AddArg(off)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConvert {
			break
		}
		_ = v_0_0.Args[1]
		ptr := v_0_0.Args[0]
		mem := v_0_0.Args[1]
		off := v_0.Args[1]
		if mem != v.Args[1] {
			break
		}
		v.reset(OpAdd32)
		v.AddArg(ptr)
		v.AddArg(off)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		off := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConvert {
			break
		}
		_ = v_0_1.Args[1]
		ptr := v_0_1.Args[0]
		mem := v_0_1.Args[1]
		if mem != v.Args[1] {
			break
		}
		v.reset(OpAdd32)
		v.AddArg(ptr)
		v.AddArg(off)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConvert {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_0.Args[1]
		if mem != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = ptr.Type
		v.AddArg(ptr)
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt32Fto32_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(i2f(c)))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt32Fto64_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(i2f(c))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt32Fto64F_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt32to32F_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = f2i(float64(float32(int32(c))))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt32to64F_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = f2i(float64(int32(c)))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt64Fto32_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(i2f(c)))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt64Fto32F_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = f2i(float64(i2f32(c)))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt64Fto64_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(i2f(c))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt64to32F_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = f2i(float64(float32(c)))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt64to64F_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = f2i(float64(c))
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c) / int16(d))
		return true
	}

	for {
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xffff)) {
			break
		}
		v.reset(OpRsh16Ux64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xffff)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<15) {
			break
		}
		v.reset(OpNeg16)
		v0 := b.NewValue0(v.Pos, OpDiv16, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst16, t)
		v1.AuxInt = -c
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != -1<<15 {
			break
		}
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpAnd16, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpNeg16, t)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = 15
		v.AddArg(v2)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh16x64)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpRsh16Ux64, t)
		v2 := b.NewValue0(v.Pos, OpRsh16x64, t)
		v2.AddArg(n)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 15
		v2.AddArg(v3)
		v1.AddArg(v2)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 16 - log2(c)
		v1.AddArg(v4)
		v0.AddArg(v1)
		v.AddArg(v0)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = log2(c)
		v.AddArg(v5)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(16, c)) {
			break
		}
		v.reset(OpSub16)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(smagic(16, c).m)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v3.AddArg(x)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 16 + smagic(16, c).s
		v0.AddArg(v4)
		v.AddArg(v0)
		v5 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v6 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v6.AddArg(x)
		v5.AddArg(v6)
		v7 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v7.AuxInt = 31
		v5.AddArg(v7)
		v.AddArg(v5)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv16u_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = int64(int16(uint16(c) / uint16(d)))
		return true
	}

	for {
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xffff)) {
			break
		}
		v.reset(OpRsh16Ux64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xffff)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(16, c) && config.RegSize == 8) {
			break
		}
		v.reset(OpTrunc64to16)
		v0 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(1<<16 + umagic(16, c).m)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 16 + umagic(16, c).s
		v0.AddArg(v4)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(16, c) && config.RegSize == 4 && umagic(16, c).m&1 == 0) {
			break
		}
		v.reset(OpTrunc32to16)
		v0 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(1<<15 + umagic(16, c).m/2)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(x)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 16 + umagic(16, c).s - 1
		v0.AddArg(v4)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(16, c) && config.RegSize == 4 && c&1 == 0) {
			break
		}
		v.reset(OpTrunc32to16)
		v0 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(1<<15 + (umagic(16, c).m+1)/2)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(x)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = 1
		v3.AddArg(v5)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v6.AuxInt = 16 + umagic(16, c).s - 2
		v0.AddArg(v6)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(16, c) && config.RegSize == 4 && config.useAvg) {
			break
		}
		v.reset(OpTrunc32to16)
		v0 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpAvg32u, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpLsh32x64, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 16
		v2.AddArg(v4)
		v1.AddArg(v2)
		v5 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v6 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v6.AuxInt = int64(umagic(16, c).m)
		v5.AddArg(v6)
		v7 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v7.AddArg(x)
		v5.AddArg(v7)
		v1.AddArg(v5)
		v0.AddArg(v1)
		v8 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v8.AuxInt = 16 + umagic(16, c).s - 1
		v0.AddArg(v8)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv32_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c) / int32(d))
		return true
	}

	for {
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xffffffff)) {
			break
		}
		v.reset(OpRsh32Ux64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xffffffff)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<31) {
			break
		}
		v.reset(OpNeg32)
		v0 := b.NewValue0(v.Pos, OpDiv32, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst32, t)
		v1.AuxInt = -c
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != -1<<31 {
			break
		}
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpAnd32, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpNeg32, t)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = 31
		v.AddArg(v2)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh32x64)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpRsh32Ux64, t)
		v2 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v2.AddArg(n)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 31
		v2.AddArg(v3)
		v1.AddArg(v2)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 32 - log2(c)
		v1.AddArg(v4)
		v0.AddArg(v1)
		v.AddArg(v0)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = log2(c)
		v.AddArg(v5)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(32, c) && config.RegSize == 8) {
			break
		}
		v.reset(OpSub32)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v1 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(smagic(32, c).m)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 32 + smagic(32, c).s
		v0.AddArg(v4)
		v.AddArg(v0)
		v5 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v6 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v6.AddArg(x)
		v5.AddArg(v6)
		v7 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v7.AuxInt = 63
		v5.AddArg(v7)
		v.AddArg(v5)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(32, c) && config.RegSize == 4 && smagic(32, c).m&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpSub32)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v1 := b.NewValue0(v.Pos, OpHmul32, t)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(int32(smagic(32, c).m / 2))
		v1.AddArg(v2)
		v1.AddArg(x)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = smagic(32, c).s - 1
		v0.AddArg(v3)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = 31
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(32, c) && config.RegSize == 4 && smagic(32, c).m&1 != 0 && config.useHmul) {
			break
		}
		v.reset(OpSub32)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v1 := b.NewValue0(v.Pos, OpAdd32, t)
		v2 := b.NewValue0(v.Pos, OpHmul32, t)
		v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v3.AuxInt = int64(int32(smagic(32, c).m))
		v2.AddArg(v3)
		v2.AddArg(x)
		v1.AddArg(v2)
		v1.AddArg(x)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = smagic(32, c).s
		v0.AddArg(v4)
		v.AddArg(v0)
		v5 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v5.AddArg(x)
		v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v6.AuxInt = 31
		v5.AddArg(v6)
		v.AddArg(v5)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = f2i(float64(i2f32(c) / i2f32(d)))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(reciprocalExact32(float32(i2f(c)))) {
			break
		}
		v.reset(OpMul32F)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst32F, t)
		v0.AuxInt = f2i(1 / i2f(c))
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv32u_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int64(int32(uint32(c) / uint32(d)))
		return true
	}

	for {
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xffffffff)) {
			break
		}
		v.reset(OpRsh32Ux64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xffffffff)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 4 && umagic(32, c).m&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpRsh32Ux64)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpHmul32u, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v1.AuxInt = int64(int32(1<<31 + umagic(32, c).m/2))
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = umagic(32, c).s - 1
		v.AddArg(v2)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 4 && c&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpRsh32Ux64)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpHmul32u, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v1.AuxInt = int64(int32(1<<31 + (umagic(32, c).m+1)/2))
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 1
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = umagic(32, c).s - 2
		v.AddArg(v4)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 4 && config.useAvg && config.useHmul) {
			break
		}
		v.reset(OpRsh32Ux64)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpAvg32u, typ.UInt32)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpHmul32u, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(int32(umagic(32, c).m))
		v1.AddArg(v2)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = umagic(32, c).s - 1
		v.AddArg(v3)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 8 && umagic(32, c).m&1 == 0) {
			break
		}
		v.reset(OpTrunc64to32)
		v0 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(1<<31 + umagic(32, c).m/2)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 32 + umagic(32, c).s - 1
		v0.AddArg(v4)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 8 && c&1 == 0) {
			break
		}
		v.reset(OpTrunc64to32)
		v0 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(1<<31 + (umagic(32, c).m+1)/2)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = 1
		v3.AddArg(v5)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v6.AuxInt = 32 + umagic(32, c).s - 2
		v0.AddArg(v6)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 8 && config.useAvg) {
			break
		}
		v.reset(OpTrunc64to32)
		v0 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpAvg64u, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpLsh64x64, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 32
		v2.AddArg(v4)
		v1.AddArg(v2)
		v5 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
		v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt32)
		v6.AuxInt = int64(umagic(32, c).m)
		v5.AddArg(v6)
		v7 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v7.AddArg(x)
		v5.AddArg(v7)
		v1.AddArg(v5)
		v0.AddArg(v1)
		v8 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v8.AuxInt = 32 + umagic(32, c).s - 1
		v0.AddArg(v8)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv64_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = c / d
		return true
	}

	for {
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh64Ux64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != -1<<63 {
			break
		}
		if !(isNonNegative(n)) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<63) {
			break
		}
		v.reset(OpNeg64)
		v0 := b.NewValue0(v.Pos, OpDiv64, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst64, t)
		v1.AuxInt = -c
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != -1<<63 {
			break
		}
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpAnd64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = 63
		v.AddArg(v2)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpRsh64Ux64, t)
		v2 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v2.AddArg(n)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 63
		v2.AddArg(v3)
		v1.AddArg(v2)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 64 - log2(c)
		v1.AddArg(v4)
		v0.AddArg(v1)
		v.AddArg(v0)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = log2(c)
		v.AddArg(v5)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(64, c) && smagic(64, c).m&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpSub64)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v1 := b.NewValue0(v.Pos, OpHmul64, t)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(smagic(64, c).m / 2)
		v1.AddArg(v2)
		v1.AddArg(x)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = smagic(64, c).s - 1
		v0.AddArg(v3)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = 63
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(64, c) && smagic(64, c).m&1 != 0 && config.useHmul) {
			break
		}
		v.reset(OpSub64)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v1 := b.NewValue0(v.Pos, OpAdd64, t)
		v2 := b.NewValue0(v.Pos, OpHmul64, t)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = int64(smagic(64, c).m)
		v2.AddArg(v3)
		v2.AddArg(x)
		v1.AddArg(v2)
		v1.AddArg(x)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = smagic(64, c).s
		v0.AddArg(v4)
		v.AddArg(v0)
		v5 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v5.AddArg(x)
		v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v6.AuxInt = 63
		v5.AddArg(v6)
		v.AddArg(v5)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = f2i(i2f(c) / i2f(d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(reciprocalExact64(i2f(c))) {
			break
		}
		v.reset(OpMul64F)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64F, t)
		v0.AuxInt = f2i(1 / i2f(c))
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv64u_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = int64(uint64(c) / uint64(d))
		return true
	}

	for {
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh64Ux64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != -1<<63 {
			break
		}
		v.reset(OpRsh64Ux64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = 63
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(64, c) && config.RegSize == 8 && umagic(64, c).m&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpRsh64Ux64)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpHmul64u, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = int64(1<<63 + umagic(64, c).m/2)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = umagic(64, c).s - 1
		v.AddArg(v2)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(64, c) && config.RegSize == 8 && c&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpRsh64Ux64)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpHmul64u, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = int64(1<<63 + (umagic(64, c).m+1)/2)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 1
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = umagic(64, c).s - 2
		v.AddArg(v4)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(64, c) && config.RegSize == 8 && config.useAvg && config.useHmul) {
			break
		}
		v.reset(OpRsh64Ux64)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpAvg64u, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpHmul64u, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(umagic(64, c).m)
		v1.AddArg(v2)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = umagic(64, c).s - 1
		v.AddArg(v3)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c) / int8(d))
		return true
	}

	for {
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xff)) {
			break
		}
		v.reset(OpRsh8Ux64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xff)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<7) {
			break
		}
		v.reset(OpNeg8)
		v0 := b.NewValue0(v.Pos, OpDiv8, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst8, t)
		v1.AuxInt = -c
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != -1<<7 {
			break
		}
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpAnd8, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpNeg8, t)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = 7
		v.AddArg(v2)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh8x64)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpRsh8Ux64, t)
		v2 := b.NewValue0(v.Pos, OpRsh8x64, t)
		v2.AddArg(n)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 7
		v2.AddArg(v3)
		v1.AddArg(v2)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 8 - log2(c)
		v1.AddArg(v4)
		v0.AddArg(v1)
		v.AddArg(v0)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = log2(c)
		v.AddArg(v5)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(8, c)) {
			break
		}
		v.reset(OpSub8)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(smagic(8, c).m)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v3.AddArg(x)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 8 + smagic(8, c).s
		v0.AddArg(v4)
		v.AddArg(v0)
		v5 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v6 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v6.AddArg(x)
		v5.AddArg(v6)
		v7 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v7.AuxInt = 31
		v5.AddArg(v7)
		v.AddArg(v5)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = int64(int8(uint8(c) / uint8(d)))
		return true
	}

	for {
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xff)) {
			break
		}
		v.reset(OpRsh8Ux64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xff)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(8, c)) {
			break
		}
		v.reset(OpTrunc32to8)
		v0 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(1<<8 + umagic(8, c).m)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v3.AddArg(x)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 8 + umagic(8, c).s
		v0.AddArg(v4)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpEq16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd16 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst16 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpEq16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd16 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst16 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpEq16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst16 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpEq16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst16 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpEq16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		s := v.Args[0]
		if s.Op != OpSub16 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpEq16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		s := v.Args[1]
		if s.Op != OpSub16 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpEq16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpEq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd32 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpEq32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd32 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpEq32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpEq32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpEq32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		s := v.Args[0]
		if s.Op != OpSub32 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpEq32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		s := v.Args[1]
		if s.Op != OpSub32 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpEq32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpEq32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) == i2f(d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) == i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpEq64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd64 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpEq64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd64 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpEq64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpEq64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpEq64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		s := v.Args[0]
		if s.Op != OpSub64 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpEq64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		s := v.Args[1]
		if s.Op != OpSub64 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpEq64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpEq64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) == i2f(d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) == i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpEq8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd8 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst8 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpEq8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd8 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst8 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpEq8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst8 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpEq8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst8 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpEq8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		s := v.Args[0]
		if s.Op != OpSub8 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpEq8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		s := v.Args[1]
		if s.Op != OpSub8 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpEq8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpEqB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConstBool {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConstBool {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConstBool {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConstBool {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConstBool {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpNot)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConstBool {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpNot)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConstBool {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConstBool {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpEqInter_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpEqPtr)
		v0 := b.NewValue0(v.Pos, OpITab, typ.Uintptr)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpITab, typ.Uintptr)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuegeneric_OpEqPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAddr {
			break
		}
		a := v_0.Aux
		v_1 := v.Args[1]
		if v_1.Op != OpAddr {
			break
		}
		b := v_1.Aux
		v.reset(OpConstBool)
		v.AuxInt = b2i(a == b)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAddr {
			break
		}
		b := v_0.Aux
		v_1 := v.Args[1]
		if v_1.Op != OpAddr {
			break
		}
		a := v_1.Aux
		v.reset(OpConstBool)
		v.AuxInt = b2i(a == b)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o1 := v_0.AuxInt
		p1 := v_0.Args[0]
		p2 := v.Args[1]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = b2i(o1 == 0)
		return true
	}

	for {
		_ = v.Args[1]
		p2 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOffPtr {
			break
		}
		o1 := v_1.AuxInt
		p1 := v_1.Args[0]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = b2i(o1 == 0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o1 := v_0.AuxInt
		p1 := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOffPtr {
			break
		}
		o2 := v_1.AuxInt
		p2 := v_1.Args[0]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = b2i(o1 == o2)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o2 := v_0.AuxInt
		p2 := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOffPtr {
			break
		}
		o1 := v_1.AuxInt
		p1 := v_1.Args[0]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = b2i(o1 == o2)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpEqPtr_10(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c == d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAddPtr {
			break
		}
		_ = v_0.Args[1]
		p1 := v_0.Args[0]
		o1 := v_0.Args[1]
		p2 := v.Args[1]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
		v0.AddArg(o1)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		p2 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAddPtr {
			break
		}
		_ = v_1.Args[1]
		p1 := v_1.Args[0]
		o1 := v_1.Args[1]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
		v0.AddArg(o1)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		p := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
		v0.AddArg(p)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
		v0.AddArg(p)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		p := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
		v0.AddArg(p)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
		v0.AddArg(p)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConstNil {
			break
		}
		p := v.Args[1]
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
		v0.AddArg(p)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConstNil {
			break
		}
		v.reset(OpNot)
		v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
		v0.AddArg(p)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpEqSlice_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpEqPtr)
		v0 := b.NewValue0(v.Pos, OpSlicePtr, typ.BytePtr)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSlicePtr, typ.BytePtr)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuegeneric_OpGeq16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c >= d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpGeq16U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint16(c) >= uint16(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGeq32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c >= d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpGeq32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) >= i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGeq32U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint32(c) >= uint32(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGeq64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c >= d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpGeq64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) >= i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGeq64U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint64(c) >= uint64(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGeq8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c >= d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpGeq8U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint8(c) >= uint8(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c > d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater16U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint16(c) > uint16(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c > d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) > i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater32U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint32(c) > uint32(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c > d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) > i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater64U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint64(c) > uint64(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c > d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater8U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint8(c) > uint8(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpIMake_0(v *Value) bool {

	for {
		_ = v.Args[1]
		typ := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStructMake1 {
			break
		}
		val := v_1.Args[0]
		v.reset(OpIMake)
		v.AddArg(typ)
		v.AddArg(val)
		return true
	}

	for {
		_ = v.Args[1]
		typ := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpArrayMake1 {
			break
		}
		val := v_1.Args[0]
		v.reset(OpIMake)
		v.AddArg(typ)
		v.AddArg(val)
		return true
	}
	return false
}
func rewriteValuegeneric_OpInterCall_0(v *Value) bool {

	for {
		argsize := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLoad {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpOffPtr {
			break
		}
		off := v_0_0.AuxInt
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpITab {
			break
		}
		v_0_0_0_0 := v_0_0_0.Args[0]
		if v_0_0_0_0.Op != OpIMake {
			break
		}
		_ = v_0_0_0_0.Args[1]
		v_0_0_0_0_0 := v_0_0_0_0.Args[0]
		if v_0_0_0_0_0.Op != OpAddr {
			break
		}
		itab := v_0_0_0_0_0.Aux
		v_0_0_0_0_0_0 := v_0_0_0_0_0.Args[0]
		if v_0_0_0_0_0_0.Op != OpSB {
			break
		}
		mem := v.Args[1]
		if !(devirt(v, itab, off) != nil) {
			break
		}
		v.reset(OpStaticCall)
		v.AuxInt = argsize
		v.Aux = devirt(v, itab, off)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuegeneric_OpIsInBounds_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to32 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !((1 << 8) <= c) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !((1 << 8) <= c) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt16to32 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !((1 << 16) <= c) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt16to64 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !((1 << 16) <= c) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst8 {
			break
		}
		c := v_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd8 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst8 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to16 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd8 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpConst8 {
			break
		}
		c := v_0_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to16 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd8 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst8 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to32 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd8 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpConst8 {
			break
		}
		c := v_0_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpIsInBounds_10(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to32 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd8 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst8 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd8 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpConst8 {
			break
		}
		c := v_0_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd8 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst8 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst16 {
			break
		}
		c := v_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst16 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt16to32 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd16 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpConst16 {
			break
		}
		c := v_0_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt16to32 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd16 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst16 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt16to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd16 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpConst16 {
			break
		}
		c := v_0_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt16to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd16 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst16 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		c := v_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpIsInBounds_20(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt32to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd32 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpConst32 {
			break
		}
		c := v_0_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt32to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd32 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst32 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		c := v_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(0 <= c && c < d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(0 <= c && c < d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMod32u {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[1]
		if y != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMod64u {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[1]
		if y != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 8 && 1<<uint(8-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpIsInBounds_30(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to32 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 8 && 1<<uint(8-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to16 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 8 && 1<<uint(8-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 8 && 1<<uint(8-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt16to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 16 && 1<<uint(16-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt16to32 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 16 && 1<<uint(16-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 16 && 1<<uint(16-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt32to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh32Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 32 && 1<<uint(32-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh32Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 32 && 1<<uint(32-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh64Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 64 && 1<<uint(64-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpIsNonNil_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConstNil {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != 0)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != 0)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAddr {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpIsSliceInBounds_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		c := v_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c <= d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c <= d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		c := v_0_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c <= d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 <= c && c <= d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(0 <= c && c <= d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(0 <= c && c <= d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpSliceLen {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSliceCap {
			break
		}
		if x != v_1.Args[0] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c <= d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq16U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint16(c) <= uint16(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c <= d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) <= i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq32U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint32(c) <= uint32(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c <= d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) <= i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq64U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint64(c) <= uint64(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c <= d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq8U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint8(c) <= uint8(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c < d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess16U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint16(c) < uint16(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c < d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) < i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess32U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint32(c) < uint32(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c < d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) < i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess64U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint64(c) < uint64(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c < d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess8U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint8(c) < uint8(d))
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpLoad_0(v *Value) bool {
	b := v.Block
	_ = b
	fe := b.Func.fe
	_ = fe

	for {
		t1 := v.Type
		_ = v.Args[1]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(p1, p2) && t1.Compare(psess.types, x.Type) == types.CMPeq && t1.Size(psess.types) == psess.sizeof(t2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		x := v_1_2.Args[1]
		if !(isSamePtr(p1, p3) && t1.Compare(psess.types, x.Type) == types.CMPeq && t1.Size(psess.types) == psess.sizeof(t2) && disjoint(p3, psess.sizeof(t3), p2, psess.sizeof(t2))) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		v_1_2_2 := v_1_2.Args[2]
		if v_1_2_2.Op != OpStore {
			break
		}
		t4 := v_1_2_2.Aux
		_ = v_1_2_2.Args[2]
		p4 := v_1_2_2.Args[0]
		x := v_1_2_2.Args[1]
		if !(isSamePtr(p1, p4) && t1.Compare(psess.types, x.Type) == types.CMPeq && t1.Size(psess.types) == psess.sizeof(t2) && disjoint(p4, psess.sizeof(t4), p2, psess.sizeof(t2)) && disjoint(p4, psess.sizeof(t4), p3, psess.sizeof(t3))) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		v_1_2_2 := v_1_2.Args[2]
		if v_1_2_2.Op != OpStore {
			break
		}
		t4 := v_1_2_2.Aux
		_ = v_1_2_2.Args[2]
		p4 := v_1_2_2.Args[0]
		v_1_2_2_2 := v_1_2_2.Args[2]
		if v_1_2_2_2.Op != OpStore {
			break
		}
		t5 := v_1_2_2_2.Aux
		_ = v_1_2_2_2.Args[2]
		p5 := v_1_2_2_2.Args[0]
		x := v_1_2_2_2.Args[1]
		if !(isSamePtr(p1, p5) && t1.Compare(psess.types, x.Type) == types.CMPeq && t1.Size(psess.types) == psess.sizeof(t2) && disjoint(p5, psess.sizeof(t5), p2, psess.sizeof(t2)) && disjoint(p5, psess.sizeof(t5), p3, psess.sizeof(t3)) && disjoint(p5, psess.sizeof(t5), p4, psess.sizeof(t4))) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		x := v_1_1.AuxInt
		if !(isSamePtr(p1, p2) && psess.sizeof(t2) == 8 && psess.is64BitFloat(t1)) {
			break
		}
		v.reset(OpConst64F)
		v.AuxInt = x
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		x := v_1_1.AuxInt
		if !(isSamePtr(p1, p2) && psess.sizeof(t2) == 4 && psess.is32BitFloat(t1)) {
			break
		}
		v.reset(OpConst32F)
		v.AuxInt = f2i(float64(math.Float32frombits(uint32(x))))
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64F {
			break
		}
		x := v_1_1.AuxInt
		if !(isSamePtr(p1, p2) && psess.sizeof(t2) == 8 && psess.is64BitInt(t1)) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = x
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32F {
			break
		}
		x := v_1_1.AuxInt
		if !(isSamePtr(p1, p2) && psess.sizeof(t2) == 4 && psess.is32BitInt(t1)) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int64(int32(math.Float32bits(float32(i2f(x)))))
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		op := v.Args[0]
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		mem := v_1.Args[2]
		if mem.Op != OpZero {
			break
		}
		n := mem.AuxInt
		_ = mem.Args[1]
		p3 := mem.Args[0]
		if !(o1 >= 0 && o1+t1.Size(psess.types) <= n && isSamePtr(p1, p3) && fe.CanSSA(t1) && disjoint(op, t1.Size(psess.types), p2, psess.sizeof(t2))) {
			break
		}
		b = mem.Block
		v0 := b.NewValue0(v.Pos, OpLoad, t1)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, op.Type)
		v1.AuxInt = o1
		v1.AddArg(p3)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		op := v.Args[0]
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		mem := v_1_2.Args[2]
		if mem.Op != OpZero {
			break
		}
		n := mem.AuxInt
		_ = mem.Args[1]
		p4 := mem.Args[0]
		if !(o1 >= 0 && o1+t1.Size(psess.types) <= n && isSamePtr(p1, p4) && fe.CanSSA(t1) && disjoint(op, t1.Size(psess.types), p2, psess.sizeof(t2)) && disjoint(op, t1.Size(psess.types), p3, psess.sizeof(t3))) {
			break
		}
		b = mem.Block
		v0 := b.NewValue0(v.Pos, OpLoad, t1)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, op.Type)
		v1.AuxInt = o1
		v1.AddArg(p4)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpLoad_10(v *Value) bool {
	b := v.Block
	_ = b
	fe := b.Func.fe
	_ = fe

	for {
		t1 := v.Type
		_ = v.Args[1]
		op := v.Args[0]
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		v_1_2_2 := v_1_2.Args[2]
		if v_1_2_2.Op != OpStore {
			break
		}
		t4 := v_1_2_2.Aux
		_ = v_1_2_2.Args[2]
		p4 := v_1_2_2.Args[0]
		mem := v_1_2_2.Args[2]
		if mem.Op != OpZero {
			break
		}
		n := mem.AuxInt
		_ = mem.Args[1]
		p5 := mem.Args[0]
		if !(o1 >= 0 && o1+t1.Size(psess.types) <= n && isSamePtr(p1, p5) && fe.CanSSA(t1) && disjoint(op, t1.Size(psess.types), p2, psess.sizeof(t2)) && disjoint(op, t1.Size(psess.types), p3, psess.sizeof(t3)) && disjoint(op, t1.Size(psess.types), p4, psess.sizeof(t4))) {
			break
		}
		b = mem.Block
		v0 := b.NewValue0(v.Pos, OpLoad, t1)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, op.Type)
		v1.AuxInt = o1
		v1.AddArg(p5)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		op := v.Args[0]
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		v_1_2_2 := v_1_2.Args[2]
		if v_1_2_2.Op != OpStore {
			break
		}
		t4 := v_1_2_2.Aux
		_ = v_1_2_2.Args[2]
		p4 := v_1_2_2.Args[0]
		v_1_2_2_2 := v_1_2_2.Args[2]
		if v_1_2_2_2.Op != OpStore {
			break
		}
		t5 := v_1_2_2_2.Aux
		_ = v_1_2_2_2.Args[2]
		p5 := v_1_2_2_2.Args[0]
		mem := v_1_2_2_2.Args[2]
		if mem.Op != OpZero {
			break
		}
		n := mem.AuxInt
		_ = mem.Args[1]
		p6 := mem.Args[0]
		if !(o1 >= 0 && o1+t1.Size(psess.types) <= n && isSamePtr(p1, p6) && fe.CanSSA(t1) && disjoint(op, t1.Size(psess.types), p2, psess.sizeof(t2)) && disjoint(op, t1.Size(psess.types), p3, psess.sizeof(t3)) && disjoint(op, t1.Size(psess.types), p4, psess.sizeof(t4)) && disjoint(op, t1.Size(psess.types), p5, psess.sizeof(t5))) {
			break
		}
		b = mem.Block
		v0 := b.NewValue0(v.Pos, OpLoad, t1)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, op.Type)
		v1.AuxInt = o1
		v1.AddArg(p6)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(t1.IsBoolean() && isSamePtr(p1, p2) && n >= o+1) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(psess.is8BitInt(t1) && isSamePtr(p1, p2) && n >= o+1) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(psess.is16BitInt(t1) && isSamePtr(p1, p2) && n >= o+2) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(psess.is32BitInt(t1) && isSamePtr(p1, p2) && n >= o+4) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(psess.is64BitInt(t1) && isSamePtr(p1, p2) && n >= o+8) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(psess.is32BitFloat(t1) && isSamePtr(p1, p2) && n >= o+4) {
			break
		}
		v.reset(OpConst32F)
		v.AuxInt = 0
		return true
	}

	for {
		t1 := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(psess.is64BitFloat(t1) && isSamePtr(p1, p2) && n >= o+8) {
			break
		}
		v.reset(OpConst64F)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		if !(t.IsStruct() && t.NumFields(psess.types) == 0 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpLoad_20(v *Value) bool {
	b := v.Block
	_ = b
	fe := b.Func.fe
	_ = fe

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsStruct() && t.NumFields(psess.types) == 1 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake1)
		v0 := b.NewValue0(v.Pos, OpLoad, t.FieldType(psess.types, 0))
		v1 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 0).PtrTo(psess.types))
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v0.AddArg(v1)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsStruct() && t.NumFields(psess.types) == 2 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake2)
		v0 := b.NewValue0(v.Pos, OpLoad, t.FieldType(psess.types, 0))
		v1 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 0).PtrTo(psess.types))
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v0.AddArg(v1)
		v0.AddArg(mem)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpLoad, t.FieldType(psess.types, 1))
		v3 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 1).PtrTo(psess.types))
		v3.AuxInt = t.FieldOff(psess.types, 1)
		v3.AddArg(ptr)
		v2.AddArg(v3)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsStruct() && t.NumFields(psess.types) == 3 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake3)
		v0 := b.NewValue0(v.Pos, OpLoad, t.FieldType(psess.types, 0))
		v1 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 0).PtrTo(psess.types))
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v0.AddArg(v1)
		v0.AddArg(mem)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpLoad, t.FieldType(psess.types, 1))
		v3 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 1).PtrTo(psess.types))
		v3.AuxInt = t.FieldOff(psess.types, 1)
		v3.AddArg(ptr)
		v2.AddArg(v3)
		v2.AddArg(mem)
		v.AddArg(v2)
		v4 := b.NewValue0(v.Pos, OpLoad, t.FieldType(psess.types, 2))
		v5 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 2).PtrTo(psess.types))
		v5.AuxInt = t.FieldOff(psess.types, 2)
		v5.AddArg(ptr)
		v4.AddArg(v5)
		v4.AddArg(mem)
		v.AddArg(v4)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsStruct() && t.NumFields(psess.types) == 4 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake4)
		v0 := b.NewValue0(v.Pos, OpLoad, t.FieldType(psess.types, 0))
		v1 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 0).PtrTo(psess.types))
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v0.AddArg(v1)
		v0.AddArg(mem)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpLoad, t.FieldType(psess.types, 1))
		v3 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 1).PtrTo(psess.types))
		v3.AuxInt = t.FieldOff(psess.types, 1)
		v3.AddArg(ptr)
		v2.AddArg(v3)
		v2.AddArg(mem)
		v.AddArg(v2)
		v4 := b.NewValue0(v.Pos, OpLoad, t.FieldType(psess.types, 2))
		v5 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 2).PtrTo(psess.types))
		v5.AuxInt = t.FieldOff(psess.types, 2)
		v5.AddArg(ptr)
		v4.AddArg(v5)
		v4.AddArg(mem)
		v.AddArg(v4)
		v6 := b.NewValue0(v.Pos, OpLoad, t.FieldType(psess.types, 3))
		v7 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 3).PtrTo(psess.types))
		v7.AuxInt = t.FieldOff(psess.types, 3)
		v7.AddArg(ptr)
		v6.AddArg(v7)
		v6.AddArg(mem)
		v.AddArg(v6)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		if !(t.IsArray() && t.NumElem(psess.types) == 0) {
			break
		}
		v.reset(OpArrayMake0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsArray() && t.NumElem(psess.types) == 1 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpArrayMake1)
		v0 := b.NewValue0(v.Pos, OpLoad, t.Elem(psess.types))
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh16x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh16x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c) << uint64(d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh16x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpLsh16x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLsh16x64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpLsh16x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh16x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh32x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh32x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c) << uint64(d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpLsh32x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh32Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpLsh32x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh32x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c << uint64(d)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh64Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh8x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh8x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c) << uint64(d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh8x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpLsh8x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLsh8x64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpLsh8x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh8x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c % d))
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xffff)) {
			break
		}
		v.reset(OpAnd16)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = (c & 0xffff) - 1
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<15) {
			break
		}
		v.reset(OpMod16)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = -c
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst16 && (c > 0 || c == -1<<15)) {
			break
		}
		v.reset(OpSub16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMul16, t)
		v1 := b.NewValue0(v.Pos, OpDiv16, t)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpConst16, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpConst16, t)
		v3.AuxInt = c
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod16u_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = int64(uint16(c) % uint16(d))
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xffff)) {
			break
		}
		v.reset(OpAnd16)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = (c & 0xffff) - 1
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst16 && c > 0 && umagicOK(16, c)) {
			break
		}
		v.reset(OpSub16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMul16, t)
		v1 := b.NewValue0(v.Pos, OpDiv16u, t)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpConst16, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpConst16, t)
		v3.AuxInt = c
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c % d))
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xffffffff)) {
			break
		}
		v.reset(OpAnd32)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = (c & 0xffffffff) - 1
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<31) {
			break
		}
		v.reset(OpMod32)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = -c
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst32 && (c > 0 || c == -1<<31)) {
			break
		}
		v.reset(OpSub32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMul32, t)
		v1 := b.NewValue0(v.Pos, OpDiv32, t)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpConst32, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpConst32, t)
		v3.AuxInt = c
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod32u_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int64(uint32(c) % uint32(d))
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xffffffff)) {
			break
		}
		v.reset(OpAnd32)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = (c & 0xffffffff) - 1
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst32 && c > 0 && umagicOK(32, c)) {
			break
		}
		v.reset(OpSub32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMul32, t)
		v1 := b.NewValue0(v.Pos, OpDiv32u, t)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpConst32, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpConst32, t)
		v3.AuxInt = c
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = c % d
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c)) {
			break
		}
		v.reset(OpAnd64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - 1
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != -1<<63 {
			break
		}
		if !(isNonNegative(n)) {
			break
		}
		v.reset(OpCopy)
		v.Type = n.Type
		v.AddArg(n)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<63) {
			break
		}
		v.reset(OpMod64)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = -c
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst64 && (c > 0 || c == -1<<63)) {
			break
		}
		v.reset(OpSub64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMul64, t)
		v1 := b.NewValue0(v.Pos, OpDiv64, t)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = c
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod64u_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = int64(uint64(c) % uint64(d))
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpAnd64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - 1
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != -1<<63 {
			break
		}
		v.reset(OpAnd64)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = 1<<63 - 1
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst64 && c > 0 && umagicOK(64, c)) {
			break
		}
		v.reset(OpSub64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMul64, t)
		v1 := b.NewValue0(v.Pos, OpDiv64u, t)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = c
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c % d))
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xff)) {
			break
		}
		v.reset(OpAnd8)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = (c & 0xff) - 1
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<7) {
			break
		}
		v.reset(OpMod8)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = -c
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst8 && (c > 0 || c == -1<<7)) {
			break
		}
		v.reset(OpSub8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMul8, t)
		v1 := b.NewValue0(v.Pos, OpDiv8, t)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpConst8, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpConst8, t)
		v3.AuxInt = c
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod8u_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = int64(uint8(c) % uint8(d))
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xff)) {
			break
		}
		v.reset(OpAnd8)
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = (c & 0xff) - 1
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst8 && c > 0 && umagicOK(8, c)) {
			break
		}
		v.reset(OpSub8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMul8, t)
		v1 := b.NewValue0(v.Pos, OpDiv8u, t)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpConst8, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpConst8, t)
		v3.AuxInt = c
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpMove_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		n := v.AuxInt
		t := v.Aux
		_ = v.Args[2]
		dst1 := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpZero {
			break
		}
		if mem.AuxInt != n {
			break
		}
		if mem.Aux != t {
			break
		}
		_ = mem.Args[1]
		dst2 := mem.Args[0]
		if !(isSamePtr(src, dst2)) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = n
		v.Aux = t
		v.AddArg(dst1)
		v.AddArg(mem)
		return true
	}

	for {
		n := v.AuxInt
		t := v.Aux
		_ = v.Args[2]
		dst1 := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpZero {
			break
		}
		if mem_0.AuxInt != n {
			break
		}
		if mem_0.Aux != t {
			break
		}
		_ = mem_0.Args[1]
		dst0 := mem_0.Args[0]
		if !(isSamePtr(src, dst0)) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = n
		v.Aux = t
		v.AddArg(dst1)
		v.AddArg(mem)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst1 := v.Args[0]
		src1 := v.Args[1]
		store := v.Args[2]
		if store.Op != OpStore {
			break
		}
		t2 := store.Aux
		_ = store.Args[2]
		op := store.Args[0]
		if op.Op != OpOffPtr {
			break
		}
		o2 := op.AuxInt
		dst2 := op.Args[0]
		mem := store.Args[2]
		if !(isSamePtr(dst1, dst2) && store.Uses == 1 && n >= o2+psess.sizeof(t2) && disjoint(src1, n, op, psess.sizeof(t2)) && clobber(store)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = n
		v.Aux = t1
		v.AddArg(dst1)
		v.AddArg(src1)
		v.AddArg(mem)
		return true
	}

	for {
		n := v.AuxInt
		t := v.Aux
		_ = v.Args[2]
		dst1 := v.Args[0]
		src1 := v.Args[1]
		move := v.Args[2]
		if move.Op != OpMove {
			break
		}
		if move.AuxInt != n {
			break
		}
		if move.Aux != t {
			break
		}
		_ = move.Args[2]
		dst2 := move.Args[0]
		mem := move.Args[2]
		if !(move.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(move)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = n
		v.Aux = t
		v.AddArg(dst1)
		v.AddArg(src1)
		v.AddArg(mem)
		return true
	}

	for {
		n := v.AuxInt
		t := v.Aux
		_ = v.Args[2]
		dst1 := v.Args[0]
		src1 := v.Args[1]
		vardef := v.Args[2]
		if vardef.Op != OpVarDef {
			break
		}
		x := vardef.Aux
		move := vardef.Args[0]
		if move.Op != OpMove {
			break
		}
		if move.AuxInt != n {
			break
		}
		if move.Aux != t {
			break
		}
		_ = move.Args[2]
		dst2 := move.Args[0]
		mem := move.Args[2]
		if !(move.Uses == 1 && vardef.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(move) && clobber(vardef)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = n
		v.Aux = t
		v.AddArg(dst1)
		v.AddArg(src1)
		v0 := b.NewValue0(v.Pos, OpVarDef, psess.types.TypeMem)
		v0.Aux = x
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		n := v.AuxInt
		t := v.Aux
		_ = v.Args[2]
		dst1 := v.Args[0]
		src1 := v.Args[1]
		zero := v.Args[2]
		if zero.Op != OpZero {
			break
		}
		if zero.AuxInt != n {
			break
		}
		if zero.Aux != t {
			break
		}
		_ = zero.Args[1]
		dst2 := zero.Args[0]
		mem := zero.Args[1]
		if !(zero.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(zero)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = n
		v.Aux = t
		v.AddArg(dst1)
		v.AddArg(src1)
		v.AddArg(mem)
		return true
	}

	for {
		n := v.AuxInt
		t := v.Aux
		_ = v.Args[2]
		dst1 := v.Args[0]
		src1 := v.Args[1]
		vardef := v.Args[2]
		if vardef.Op != OpVarDef {
			break
		}
		x := vardef.Aux
		zero := vardef.Args[0]
		if zero.Op != OpZero {
			break
		}
		if zero.AuxInt != n {
			break
		}
		if zero.Aux != t {
			break
		}
		_ = zero.Args[1]
		dst2 := zero.Args[0]
		mem := zero.Args[1]
		if !(zero.Uses == 1 && vardef.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(zero) && clobber(vardef)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = n
		v.Aux = t
		v.AddArg(dst1)
		v.AddArg(src1)
		v0 := b.NewValue0(v.Pos, OpVarDef, psess.types.TypeMem)
		v0.Aux = x
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		op2 := mem.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		op3 := mem_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		if op3.AuxInt != 0 {
			break
		}
		p3 := op3.Args[0]
		d2 := mem_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && o2 == psess.sizeof(t3) && n == psess.sizeof(t2)+psess.sizeof(t3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, t2.(*types.Type))
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, t3.(*types.Type))
		v2.AuxInt = 0
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		op2 := mem.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		op3 := mem_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d2 := mem_2.Args[1]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		op4 := mem_2_2.Args[0]
		if op4.Op != OpOffPtr {
			break
		}
		if op4.AuxInt != 0 {
			break
		}
		p4 := op4.Args[0]
		d3 := mem_2_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.alignof(t4) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && psess.registerizable(b, t4) && o3 == psess.sizeof(t4) && o2-o3 == psess.sizeof(t3) && n == psess.sizeof(t2)+psess.sizeof(t3)+psess.sizeof(t4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, t2.(*types.Type))
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, t3.(*types.Type))
		v2.AuxInt = o3
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v3 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, t4.(*types.Type))
		v4.AuxInt = 0
		v4.AddArg(dst)
		v3.AddArg(v4)
		v3.AddArg(d3)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		op2 := mem.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		op3 := mem_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d2 := mem_2.Args[1]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		op4 := mem_2_2.Args[0]
		if op4.Op != OpOffPtr {
			break
		}
		o4 := op4.AuxInt
		p4 := op4.Args[0]
		d3 := mem_2_2.Args[1]
		mem_2_2_2 := mem_2_2.Args[2]
		if mem_2_2_2.Op != OpStore {
			break
		}
		t5 := mem_2_2_2.Aux
		_ = mem_2_2_2.Args[2]
		op5 := mem_2_2_2.Args[0]
		if op5.Op != OpOffPtr {
			break
		}
		if op5.AuxInt != 0 {
			break
		}
		p5 := op5.Args[0]
		d4 := mem_2_2_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.alignof(t4) <= psess.alignof(t1) && psess.alignof(t5) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && psess.registerizable(b, t4) && psess.registerizable(b, t5) && o4 == psess.sizeof(t5) && o3-o4 == psess.sizeof(t4) && o2-o3 == psess.sizeof(t3) && n == psess.sizeof(t2)+psess.sizeof(t3)+psess.sizeof(t4)+psess.sizeof(t5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, t2.(*types.Type))
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, t3.(*types.Type))
		v2.AuxInt = o3
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v3 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, t4.(*types.Type))
		v4.AuxInt = o4
		v4.AddArg(dst)
		v3.AddArg(v4)
		v3.AddArg(d3)
		v5 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v5.Aux = t5
		v6 := b.NewValue0(v.Pos, OpOffPtr, t5.(*types.Type))
		v6.AuxInt = 0
		v6.AddArg(dst)
		v5.AddArg(v6)
		v5.AddArg(d4)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpMove_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		op2 := mem_0.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		op3 := mem_0_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		if op3.AuxInt != 0 {
			break
		}
		p3 := op3.Args[0]
		d2 := mem_0_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && o2 == psess.sizeof(t3) && n == psess.sizeof(t2)+psess.sizeof(t3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, t2.(*types.Type))
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, t3.(*types.Type))
		v2.AuxInt = 0
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		op2 := mem_0.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		op3 := mem_0_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d2 := mem_0_2.Args[1]
		mem_0_2_2 := mem_0_2.Args[2]
		if mem_0_2_2.Op != OpStore {
			break
		}
		t4 := mem_0_2_2.Aux
		_ = mem_0_2_2.Args[2]
		op4 := mem_0_2_2.Args[0]
		if op4.Op != OpOffPtr {
			break
		}
		if op4.AuxInt != 0 {
			break
		}
		p4 := op4.Args[0]
		d3 := mem_0_2_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.alignof(t4) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && psess.registerizable(b, t4) && o3 == psess.sizeof(t4) && o2-o3 == psess.sizeof(t3) && n == psess.sizeof(t2)+psess.sizeof(t3)+psess.sizeof(t4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, t2.(*types.Type))
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, t3.(*types.Type))
		v2.AuxInt = o3
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v3 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, t4.(*types.Type))
		v4.AuxInt = 0
		v4.AddArg(dst)
		v3.AddArg(v4)
		v3.AddArg(d3)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		op2 := mem_0.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		op3 := mem_0_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d2 := mem_0_2.Args[1]
		mem_0_2_2 := mem_0_2.Args[2]
		if mem_0_2_2.Op != OpStore {
			break
		}
		t4 := mem_0_2_2.Aux
		_ = mem_0_2_2.Args[2]
		op4 := mem_0_2_2.Args[0]
		if op4.Op != OpOffPtr {
			break
		}
		o4 := op4.AuxInt
		p4 := op4.Args[0]
		d3 := mem_0_2_2.Args[1]
		mem_0_2_2_2 := mem_0_2_2.Args[2]
		if mem_0_2_2_2.Op != OpStore {
			break
		}
		t5 := mem_0_2_2_2.Aux
		_ = mem_0_2_2_2.Args[2]
		op5 := mem_0_2_2_2.Args[0]
		if op5.Op != OpOffPtr {
			break
		}
		if op5.AuxInt != 0 {
			break
		}
		p5 := op5.Args[0]
		d4 := mem_0_2_2_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.alignof(t4) <= psess.alignof(t1) && psess.alignof(t5) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && psess.registerizable(b, t4) && psess.registerizable(b, t5) && o4 == psess.sizeof(t5) && o3-o4 == psess.sizeof(t4) && o2-o3 == psess.sizeof(t3) && n == psess.sizeof(t2)+psess.sizeof(t3)+psess.sizeof(t4)+psess.sizeof(t5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, t2.(*types.Type))
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, t3.(*types.Type))
		v2.AuxInt = o3
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v3 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, t4.(*types.Type))
		v4.AuxInt = o4
		v4.AddArg(dst)
		v3.AddArg(v4)
		v3.AddArg(d3)
		v5 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v5.Aux = t5
		v6 := b.NewValue0(v.Pos, OpOffPtr, t5.(*types.Type))
		v6.AuxInt = 0
		v6.AddArg(dst)
		v5.AddArg(v6)
		v5.AddArg(d4)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		op2 := mem.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		tt2 := op2.Type
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpZero {
			break
		}
		if mem_2.AuxInt != n {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[1]
		p3 := mem_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.registerizable(b, t2) && n >= o2+psess.sizeof(t2)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpZero, psess.types.TypeMem)
		v1.AuxInt = n
		v1.Aux = t1
		v1.AddArg(dst)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		mem_0 := mem.Args[0]
		if mem_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0.Type
		o2 := mem_0.AuxInt
		p2 := mem_0.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		mem_2_0 := mem_2.Args[0]
		if mem_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_2_0.Type
		o3 := mem_2_0.AuxInt
		p3 := mem_2_0.Args[0]
		d2 := mem_2.Args[1]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpZero {
			break
		}
		if mem_2_2.AuxInt != n {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[1]
		p4 := mem_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.alignof(t4) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && n >= o2+psess.sizeof(t2) && n >= o3+psess.sizeof(t3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v3 := b.NewValue0(v.Pos, OpZero, psess.types.TypeMem)
		v3.AuxInt = n
		v3.Aux = t1
		v3.AddArg(dst)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		mem_0 := mem.Args[0]
		if mem_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0.Type
		o2 := mem_0.AuxInt
		p2 := mem_0.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		mem_2_0 := mem_2.Args[0]
		if mem_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_2_0.Type
		o3 := mem_2_0.AuxInt
		p3 := mem_2_0.Args[0]
		d2 := mem_2.Args[1]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		mem_2_2_0 := mem_2_2.Args[0]
		if mem_2_2_0.Op != OpOffPtr {
			break
		}
		tt4 := mem_2_2_0.Type
		o4 := mem_2_2_0.AuxInt
		p4 := mem_2_2_0.Args[0]
		d3 := mem_2_2.Args[1]
		mem_2_2_2 := mem_2_2.Args[2]
		if mem_2_2_2.Op != OpZero {
			break
		}
		if mem_2_2_2.AuxInt != n {
			break
		}
		t5 := mem_2_2_2.Aux
		_ = mem_2_2_2.Args[1]
		p5 := mem_2_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.alignof(t4) <= psess.alignof(t1) && psess.alignof(t5) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && psess.registerizable(b, t4) && n >= o2+psess.sizeof(t2) && n >= o3+psess.sizeof(t3) && n >= o4+psess.sizeof(t4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v3 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = o4
		v4.AddArg(dst)
		v3.AddArg(v4)
		v3.AddArg(d3)
		v5 := b.NewValue0(v.Pos, OpZero, psess.types.TypeMem)
		v5.AuxInt = n
		v5.Aux = t1
		v5.AddArg(dst)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		mem_0 := mem.Args[0]
		if mem_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0.Type
		o2 := mem_0.AuxInt
		p2 := mem_0.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		mem_2_0 := mem_2.Args[0]
		if mem_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_2_0.Type
		o3 := mem_2_0.AuxInt
		p3 := mem_2_0.Args[0]
		d2 := mem_2.Args[1]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		mem_2_2_0 := mem_2_2.Args[0]
		if mem_2_2_0.Op != OpOffPtr {
			break
		}
		tt4 := mem_2_2_0.Type
		o4 := mem_2_2_0.AuxInt
		p4 := mem_2_2_0.Args[0]
		d3 := mem_2_2.Args[1]
		mem_2_2_2 := mem_2_2.Args[2]
		if mem_2_2_2.Op != OpStore {
			break
		}
		t5 := mem_2_2_2.Aux
		_ = mem_2_2_2.Args[2]
		mem_2_2_2_0 := mem_2_2_2.Args[0]
		if mem_2_2_2_0.Op != OpOffPtr {
			break
		}
		tt5 := mem_2_2_2_0.Type
		o5 := mem_2_2_2_0.AuxInt
		p5 := mem_2_2_2_0.Args[0]
		d4 := mem_2_2_2.Args[1]
		mem_2_2_2_2 := mem_2_2_2.Args[2]
		if mem_2_2_2_2.Op != OpZero {
			break
		}
		if mem_2_2_2_2.AuxInt != n {
			break
		}
		t6 := mem_2_2_2_2.Aux
		_ = mem_2_2_2_2.Args[1]
		p6 := mem_2_2_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && isSamePtr(p5, p6) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.alignof(t4) <= psess.alignof(t1) && psess.alignof(t5) <= psess.alignof(t1) && psess.alignof(t6) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && psess.registerizable(b, t4) && psess.registerizable(b, t5) && n >= o2+psess.sizeof(t2) && n >= o3+psess.sizeof(t3) && n >= o4+psess.sizeof(t4) && n >= o5+psess.sizeof(t5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v3 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = o4
		v4.AddArg(dst)
		v3.AddArg(v4)
		v3.AddArg(d3)
		v5 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v5.Aux = t5
		v6 := b.NewValue0(v.Pos, OpOffPtr, tt5)
		v6.AuxInt = o5
		v6.AddArg(dst)
		v5.AddArg(v6)
		v5.AddArg(d4)
		v7 := b.NewValue0(v.Pos, OpZero, psess.types.TypeMem)
		v7.AuxInt = n
		v7.Aux = t1
		v7.AddArg(dst)
		v7.AddArg(mem)
		v5.AddArg(v7)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		op2 := mem_0.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		tt2 := op2.Type
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpZero {
			break
		}
		if mem_0_2.AuxInt != n {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[1]
		p3 := mem_0_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.registerizable(b, t2) && n >= o2+psess.sizeof(t2)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpZero, psess.types.TypeMem)
		v1.AuxInt = n
		v1.Aux = t1
		v1.AddArg(dst)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		mem_0_0 := mem_0.Args[0]
		if mem_0_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0_0.Type
		o2 := mem_0_0.AuxInt
		p2 := mem_0_0.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		mem_0_2_0 := mem_0_2.Args[0]
		if mem_0_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_0_2_0.Type
		o3 := mem_0_2_0.AuxInt
		p3 := mem_0_2_0.Args[0]
		d2 := mem_0_2.Args[1]
		mem_0_2_2 := mem_0_2.Args[2]
		if mem_0_2_2.Op != OpZero {
			break
		}
		if mem_0_2_2.AuxInt != n {
			break
		}
		t4 := mem_0_2_2.Aux
		_ = mem_0_2_2.Args[1]
		p4 := mem_0_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.alignof(t4) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && n >= o2+psess.sizeof(t2) && n >= o3+psess.sizeof(t3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v3 := b.NewValue0(v.Pos, OpZero, psess.types.TypeMem)
		v3.AuxInt = n
		v3.Aux = t1
		v3.AddArg(dst)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		mem_0_0 := mem_0.Args[0]
		if mem_0_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0_0.Type
		o2 := mem_0_0.AuxInt
		p2 := mem_0_0.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		mem_0_2_0 := mem_0_2.Args[0]
		if mem_0_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_0_2_0.Type
		o3 := mem_0_2_0.AuxInt
		p3 := mem_0_2_0.Args[0]
		d2 := mem_0_2.Args[1]
		mem_0_2_2 := mem_0_2.Args[2]
		if mem_0_2_2.Op != OpStore {
			break
		}
		t4 := mem_0_2_2.Aux
		_ = mem_0_2_2.Args[2]
		mem_0_2_2_0 := mem_0_2_2.Args[0]
		if mem_0_2_2_0.Op != OpOffPtr {
			break
		}
		tt4 := mem_0_2_2_0.Type
		o4 := mem_0_2_2_0.AuxInt
		p4 := mem_0_2_2_0.Args[0]
		d3 := mem_0_2_2.Args[1]
		mem_0_2_2_2 := mem_0_2_2.Args[2]
		if mem_0_2_2_2.Op != OpZero {
			break
		}
		if mem_0_2_2_2.AuxInt != n {
			break
		}
		t5 := mem_0_2_2_2.Aux
		_ = mem_0_2_2_2.Args[1]
		p5 := mem_0_2_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.alignof(t4) <= psess.alignof(t1) && psess.alignof(t5) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && psess.registerizable(b, t4) && n >= o2+psess.sizeof(t2) && n >= o3+psess.sizeof(t3) && n >= o4+psess.sizeof(t4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v3 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = o4
		v4.AddArg(dst)
		v3.AddArg(v4)
		v3.AddArg(d3)
		v5 := b.NewValue0(v.Pos, OpZero, psess.types.TypeMem)
		v5.AuxInt = n
		v5.Aux = t1
		v5.AddArg(dst)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpMove_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		p1 := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		mem_0_0 := mem_0.Args[0]
		if mem_0_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0_0.Type
		o2 := mem_0_0.AuxInt
		p2 := mem_0_0.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		mem_0_2_0 := mem_0_2.Args[0]
		if mem_0_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_0_2_0.Type
		o3 := mem_0_2_0.AuxInt
		p3 := mem_0_2_0.Args[0]
		d2 := mem_0_2.Args[1]
		mem_0_2_2 := mem_0_2.Args[2]
		if mem_0_2_2.Op != OpStore {
			break
		}
		t4 := mem_0_2_2.Aux
		_ = mem_0_2_2.Args[2]
		mem_0_2_2_0 := mem_0_2_2.Args[0]
		if mem_0_2_2_0.Op != OpOffPtr {
			break
		}
		tt4 := mem_0_2_2_0.Type
		o4 := mem_0_2_2_0.AuxInt
		p4 := mem_0_2_2_0.Args[0]
		d3 := mem_0_2_2.Args[1]
		mem_0_2_2_2 := mem_0_2_2.Args[2]
		if mem_0_2_2_2.Op != OpStore {
			break
		}
		t5 := mem_0_2_2_2.Aux
		_ = mem_0_2_2_2.Args[2]
		mem_0_2_2_2_0 := mem_0_2_2_2.Args[0]
		if mem_0_2_2_2_0.Op != OpOffPtr {
			break
		}
		tt5 := mem_0_2_2_2_0.Type
		o5 := mem_0_2_2_2_0.AuxInt
		p5 := mem_0_2_2_2_0.Args[0]
		d4 := mem_0_2_2_2.Args[1]
		mem_0_2_2_2_2 := mem_0_2_2_2.Args[2]
		if mem_0_2_2_2_2.Op != OpZero {
			break
		}
		if mem_0_2_2_2_2.AuxInt != n {
			break
		}
		t6 := mem_0_2_2_2_2.Aux
		_ = mem_0_2_2_2_2.Args[1]
		p6 := mem_0_2_2_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && isSamePtr(p5, p6) && psess.alignof(t2) <= psess.alignof(t1) && psess.alignof(t3) <= psess.alignof(t1) && psess.alignof(t4) <= psess.alignof(t1) && psess.alignof(t5) <= psess.alignof(t1) && psess.alignof(t6) <= psess.alignof(t1) && psess.registerizable(b, t2) && psess.registerizable(b, t3) && psess.registerizable(b, t4) && psess.registerizable(b, t5) && n >= o2+psess.sizeof(t2) && n >= o3+psess.sizeof(t3) && n >= o4+psess.sizeof(t4) && n >= o5+psess.sizeof(t5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(d1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(d2)
		v3 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = o4
		v4.AddArg(dst)
		v3.AddArg(v4)
		v3.AddArg(d3)
		v5 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v5.Aux = t5
		v6 := b.NewValue0(v.Pos, OpOffPtr, tt5)
		v6.AuxInt = o5
		v6.AddArg(dst)
		v5.AddArg(v6)
		v5.AddArg(d4)
		v7 := b.NewValue0(v.Pos, OpZero, psess.types.TypeMem)
		v7.AuxInt = n
		v7.Aux = t1
		v7.AddArg(dst)
		v7.AddArg(mem)
		v5.AddArg(v7)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMul16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c * d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c * d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		x := v.Args[1]
		v.reset(OpNeg16)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpNeg16)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpLsh16x64)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		n := v.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpLsh16x64)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(t.IsSigned() && isPowerOfTwo(-c)) {
			break
		}
		v.reset(OpNeg16)
		v0 := b.NewValue0(v.Pos, OpLsh16x64, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = log2(-c)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		n := v.Args[1]
		if !(t.IsSigned() && isPowerOfTwo(-c)) {
			break
		}
		v.reset(OpNeg16)
		v0 := b.NewValue0(v.Pos, OpLsh16x64, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = log2(-c)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMul16_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst16 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpMul16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst16 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpMul16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst16 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMul16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst16 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMul16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMul32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c * d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c * d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		x := v.Args[1]
		v.reset(OpNeg32)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpNeg32)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpLsh32x64)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		n := v.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpLsh32x64)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(t.IsSigned() && isPowerOfTwo(-c)) {
			break
		}
		v.reset(OpNeg32)
		v0 := b.NewValue0(v.Pos, OpLsh32x64, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = log2(-c)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		n := v.Args[1]
		if !(t.IsSigned() && isPowerOfTwo(-c)) {
			break
		}
		v.reset(OpNeg32)
		v0 := b.NewValue0(v.Pos, OpLsh32x64, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = log2(-c)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMul32_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd32 {
			break
		}
		if v_1.Type != t {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c * d))
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMul32, t)
		v2 := b.NewValue0(v.Pos, OpConst32, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd32 {
			break
		}
		if v_1.Type != t {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c * d))
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMul32, t)
		v2 := b.NewValue0(v.Pos, OpConst32, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		t := v_0.Type
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		if v_0_0.Type != t {
			break
		}
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c * d))
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMul32, t)
		v2 := b.NewValue0(v.Pos, OpConst32, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		t := v_0.Type
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		if v_0_1.Type != t {
			break
		}
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c * d))
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMul32, t)
		v2 := b.NewValue0(v.Pos, OpConst32, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpMul32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpMul32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMul32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMul32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMul32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = f2i(float64(i2f32(c) * i2f32(d)))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = f2i(float64(i2f32(c) * i2f32(d)))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		if v_1.AuxInt != f2i(1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		if v_0.AuxInt != f2i(1) {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		if v_1.AuxInt != f2i(-1) {
			break
		}
		v.reset(OpNeg32F)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		if v_0.AuxInt != f2i(-1) {
			break
		}
		x := v.Args[1]
		v.reset(OpNeg32F)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		if v_1.AuxInt != f2i(2) {
			break
		}
		v.reset(OpAdd32F)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		if v_0.AuxInt != f2i(2) {
			break
		}
		x := v.Args[1]
		v.reset(OpAdd32F)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMul64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c * d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c * d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		x := v.Args[1]
		v.reset(OpNeg64)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpNeg64)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpLsh64x64)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		n := v.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpLsh64x64)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(t.IsSigned() && isPowerOfTwo(-c)) {
			break
		}
		v.reset(OpNeg64)
		v0 := b.NewValue0(v.Pos, OpLsh64x64, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = log2(-c)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		n := v.Args[1]
		if !(t.IsSigned() && isPowerOfTwo(-c)) {
			break
		}
		v.reset(OpNeg64)
		v0 := b.NewValue0(v.Pos, OpLsh64x64, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = log2(-c)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMul64_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd64 {
			break
		}
		if v_1.Type != t {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c * d
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMul64, t)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd64 {
			break
		}
		if v_1.Type != t {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c * d
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMul64, t)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		t := v_0.Type
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		if v_0_0.Type != t {
			break
		}
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c * d
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMul64, t)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		t := v_0.Type
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.Type != t {
			break
		}
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c * d
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMul64, t)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = c
		v1.AddArg(v2)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpMul64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c * d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpMul64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c * d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMul64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c * d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMul64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c * d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMul64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = f2i(i2f(c) * i2f(d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = f2i(i2f(c) * i2f(d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		if v_1.AuxInt != f2i(1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		if v_0.AuxInt != f2i(1) {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		if v_1.AuxInt != f2i(-1) {
			break
		}
		v.reset(OpNeg64F)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		if v_0.AuxInt != f2i(-1) {
			break
		}
		x := v.Args[1]
		v.reset(OpNeg64F)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		if v_1.AuxInt != f2i(2) {
			break
		}
		v.reset(OpAdd64F)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		if v_0.AuxInt != f2i(2) {
			break
		}
		x := v.Args[1]
		v.reset(OpAdd64F)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMul8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c * d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c * d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		x := v.Args[1]
		v.reset(OpNeg8)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpNeg8)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpLsh8x64)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		n := v.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpLsh8x64)
		v.Type = t
		v.AddArg(n)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		n := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(t.IsSigned() && isPowerOfTwo(-c)) {
			break
		}
		v.reset(OpNeg8)
		v0 := b.NewValue0(v.Pos, OpLsh8x64, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = log2(-c)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		n := v.Args[1]
		if !(t.IsSigned() && isPowerOfTwo(-c)) {
			break
		}
		v.reset(OpNeg8)
		v0 := b.NewValue0(v.Pos, OpLsh8x64, t)
		v0.AddArg(n)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = log2(-c)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMul8_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst8 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpMul8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst8 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpMul8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst8 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMul8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst8 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMul8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c * d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeg16_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(-int16(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSub16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpSub16)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeg32_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(-int32(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSub32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpSub32)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeg32F_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		if !(i2f(c) != 0) {
			break
		}
		v.reset(OpConst32F)
		v.AuxInt = f2i(-i2f(c))
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeg64_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = -c
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSub64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpSub64)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeg64F_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		if !(i2f(c) != 0) {
			break
		}
		v.reset(OpConst64F)
		v.AuxInt = f2i(-i2f(c))
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeg8_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(-int8(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSub8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpSub8)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeq16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd16 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst16 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpNeq16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd16 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst16 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpNeq16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst16 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpNeq16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst16 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpNeq16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		s := v.Args[0]
		if s.Op != OpSub16 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpNeq16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		s := v.Args[1]
		if s.Op != OpSub16 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpNeq16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd32 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpNeq32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd32 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpNeq32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpNeq32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpNeq32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		s := v.Args[0]
		if s.Op != OpSub32 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpNeq32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		s := v.Args[1]
		if s.Op != OpSub32 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpNeq32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeq32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) != i2f(d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) != i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeq64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd64 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpNeq64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd64 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpNeq64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpNeq64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpNeq64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		s := v.Args[0]
		if s.Op != OpSub64 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpNeq64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		s := v.Args[1]
		if s.Op != OpSub64 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpNeq64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeq64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) != i2f(d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(i2f(c) != i2f(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeq8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd8 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst8 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpNeq8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAdd8 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst8 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpNeq8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst8 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpNeq8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst8 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpNeq8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		s := v.Args[0]
		if s.Op != OpSub8 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpNeq8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		s := v.Args[1]
		if s.Op != OpSub8 {
			break
		}
		_ = s.Args[1]
		x := s.Args[0]
		y := s.Args[1]
		if !(s.Uses == 1) {
			break
		}
		v.reset(OpNeq8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeqB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConstBool {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConstBool {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConstBool {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConstBool {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConstBool {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConstBool {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConstBool {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		v.reset(OpNot)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConstBool {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpNot)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpNot {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpNot {
			break
		}
		y := v_1.Args[0]
		v.reset(OpNeqB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpNot {
			break
		}
		y := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpNot {
			break
		}
		x := v_1.Args[0]
		v.reset(OpNeqB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeqInter_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNeqPtr)
		v0 := b.NewValue0(v.Pos, OpITab, typ.Uintptr)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpITab, typ.Uintptr)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuegeneric_OpNeqPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAddr {
			break
		}
		a := v_0.Aux
		v_1 := v.Args[1]
		if v_1.Op != OpAddr {
			break
		}
		b := v_1.Aux
		v.reset(OpConstBool)
		v.AuxInt = b2i(a != b)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAddr {
			break
		}
		b := v_0.Aux
		v_1 := v.Args[1]
		if v_1.Op != OpAddr {
			break
		}
		a := v_1.Aux
		v.reset(OpConstBool)
		v.AuxInt = b2i(a != b)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o1 := v_0.AuxInt
		p1 := v_0.Args[0]
		p2 := v.Args[1]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = b2i(o1 != 0)
		return true
	}

	for {
		_ = v.Args[1]
		p2 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOffPtr {
			break
		}
		o1 := v_1.AuxInt
		p1 := v_1.Args[0]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = b2i(o1 != 0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o1 := v_0.AuxInt
		p1 := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOffPtr {
			break
		}
		o2 := v_1.AuxInt
		p2 := v_1.Args[0]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = b2i(o1 != o2)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o2 := v_0.AuxInt
		p2 := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOffPtr {
			break
		}
		o1 := v_1.AuxInt
		p1 := v_1.Args[0]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = b2i(o1 != o2)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeqPtr_10(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != d)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAddPtr {
			break
		}
		_ = v_0.Args[1]
		p1 := v_0.Args[0]
		o1 := v_0.Args[1]
		p2 := v.Args[1]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpIsNonNil)
		v.AddArg(o1)
		return true
	}

	for {
		_ = v.Args[1]
		p2 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpAddPtr {
			break
		}
		_ = v_1.Args[1]
		p1 := v_1.Args[0]
		o1 := v_1.Args[1]
		if !(isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpIsNonNil)
		v.AddArg(o1)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		p := v.Args[1]
		v.reset(OpIsNonNil)
		v.AddArg(p)
		return true
	}

	for {
		_ = v.Args[1]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpIsNonNil)
		v.AddArg(p)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		p := v.Args[1]
		v.reset(OpIsNonNil)
		v.AddArg(p)
		return true
	}

	for {
		_ = v.Args[1]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpIsNonNil)
		v.AddArg(p)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConstNil {
			break
		}
		p := v.Args[1]
		v.reset(OpIsNonNil)
		v.AddArg(p)
		return true
	}

	for {
		_ = v.Args[1]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConstNil {
			break
		}
		v.reset(OpIsNonNil)
		v.AddArg(p)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeqSlice_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpNeqPtr)
		v0 := b.NewValue0(v.Pos, OpSlicePtr, typ.BytePtr)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSlicePtr, typ.BytePtr)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuegeneric_OpNilCheck_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpGetG {
			break
		}
		mem := v_0.Args[0]
		if mem != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLoad {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpOffPtr {
			break
		}
		c := v_0_0.AuxInt
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpSP {
			break
		}
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpStaticCall {
			break
		}
		sym := v_0_1.Aux
		if !(isSameSym(sym, "runtime.newobject") && c == config.ctxt.FixedFrameSize()+config.RegSize && warnRule(fe.Debug_checknil() && v.Pos.Line() > 1, v, "removed nil check")) {
			break
		}
		v.reset(OpInvalid)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLoad {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpOffPtr {
			break
		}
		c := v_0_0_0.AuxInt
		v_0_0_0_0 := v_0_0_0.Args[0]
		if v_0_0_0_0.Op != OpSP {
			break
		}
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpStaticCall {
			break
		}
		sym := v_0_0_1.Aux
		if !(isSameSym(sym, "runtime.newobject") && c == config.ctxt.FixedFrameSize()+config.RegSize && warnRule(fe.Debug_checknil() && v.Pos.Line() > 1, v, "removed nil check")) {
			break
		}
		v.reset(OpInvalid)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNot_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConstBool {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = 1 - c
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpEq64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpNeq64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpEq32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpNeq32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpEq16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpNeq16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpEq8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpNeq8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpEqB {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpNeqB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpNeq64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpEq64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpNeq32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpEq32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpNeq16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpEq16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpNeq8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpEq8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNot_10(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpNeqB {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpEqB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGreater64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLeq64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGreater32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLeq32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGreater16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLeq16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGreater8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLeq8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGreater64U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLeq64U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGreater32U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLeq32U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGreater16U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLeq16U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGreater8U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLeq8U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGeq64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLess64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNot_20(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGeq32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLess32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGeq16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLess16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGeq8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLess8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGeq64U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLess64U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGeq32U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLess32U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGeq16U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLess16U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpGeq8U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpLess8U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLess64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGeq64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLess32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGeq32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLess16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGeq16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNot_30(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLess8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGeq8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLess64U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGeq64U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLess32U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGeq32U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLess16U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGeq16U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLess8U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGeq8U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLeq64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGreater64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLeq32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGreater32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLeq16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGreater16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLeq8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGreater8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLeq64U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGreater64U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNot_40(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLeq32U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGreater32U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLeq16U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGreater16U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLeq8U {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpGreater8U)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpOffPtr_0(v *Value) bool {

	for {
		a := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		b := v_0.AuxInt
		p := v_0.Args[0]
		v.reset(OpOffPtr)
		v.AuxInt = a + b
		v.AddArg(p)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		p := v.Args[0]
		if !(v.Type.Compare(psess.types, p.Type) == types.CMPeq) {
			break
		}
		v.reset(OpCopy)
		v.Type = p.Type
		v.AddArg(p)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c | d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c | d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = -1
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = -1
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr16 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpOr16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr16 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpOr16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpOr16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr16_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr16 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpOr16)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst16 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		t := v_1.Type
		c1 := v_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst16 {
			break
		}
		c2 := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		t := v_1.Type
		c1 := v_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c1 := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd16 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst16 {
			break
		}
		c2 := v_1_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c1 := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd16 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst16 {
			break
		}
		c2 := v_1_0.AuxInt
		x := v_1.Args[1]
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr16 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpOr16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr16 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpOr16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr16 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpOr16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpOr16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpOr16 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst16 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr16_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpOr16 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst16 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst16 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst16 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c | d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c | d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = -1
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = -1
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr32 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpOr32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr32 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpOr32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpOr32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr32_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr32 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpOr32)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		t := v_1.Type
		c1 := v_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		c2 := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		t := v_1.Type
		c1 := v_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c1 := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd32 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		c2 := v_1_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c1 := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd32 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		c2 := v_1_0.AuxInt
		x := v_1.Args[1]
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr32 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpOr32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr32 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpOr32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr32 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpOr32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpOr32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpOr32 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr32_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpOr32 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c | d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c | d
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = -1
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = -1
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr64 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpOr64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr64 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpOr64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpOr64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr64_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr64 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpOr64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		t := v_1.Type
		c1 := v_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		c2 := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		t := v_1.Type
		c1 := v_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c1 := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd64 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c2 := v_1_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c1 := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd64 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 {
			break
		}
		c2 := v_1_0.AuxInt
		x := v_1.Args[1]
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr64 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpOr64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr64 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpOr64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr64 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpOr64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpOr64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpOr64 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c | d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr64_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpOr64 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c | d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c | d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c | d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c | d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c | d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = -1
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = -1
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr8 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpOr8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr8 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpOr8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpOr8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr8_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr8 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpOr8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst8 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		t := v_1.Type
		c1 := v_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAnd8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst8 {
			break
		}
		c2 := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		t := v_1.Type
		c1 := v_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c1 := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd8 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst8 {
			break
		}
		c2 := v_1_1.AuxInt
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c1 := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpAnd8 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst8 {
			break
		}
		c2 := v_1_0.AuxInt
		x := v_1.Args[1]
		if !(^(c1 | c2) == 0) {
			break
		}
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = c1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr8 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpOr8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr8 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpOr8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr8 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpOr8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpOr8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpOr8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpOr8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpOr8 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst8 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr8_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpOr8 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst8 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst8 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpOr8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst8 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c | d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpPhi_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != c {
			break
		}
		if len(v.Args) != 2 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = c
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != c {
			break
		}
		if len(v.Args) != 2 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = c
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != c {
			break
		}
		if len(v.Args) != 2 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = c
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != c {
			break
		}
		if len(v.Args) != 2 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = c
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpPtrIndex_0(v *Value) bool {
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
		idx := v.Args[1]
		if !(config.PtrSize == 4) {
			break
		}
		v.reset(OpAddPtr)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMul32, typ.Int)
		v0.AddArg(idx)
		v1 := b.NewValue0(v.Pos, OpConst32, typ.Int)
		v1.AuxInt = t.Elem(psess.types).Size(psess.types)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		idx := v.Args[1]
		if !(config.PtrSize == 8) {
			break
		}
		v.reset(OpAddPtr)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMul64, typ.Int)
		v0.AddArg(idx)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.Int)
		v1.AuxInt = t.Elem(psess.types).Size(psess.types)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRound32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpConst32F {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRound64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpConst64F {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(uint16(c) >> uint64(d)))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh16Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh16x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpRsh16Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh16x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 8 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 8 {
			break
		}
		v.reset(OpZeroExt8to16)
		v0 := b.NewValue0(v.Pos, OpTrunc16to8, typ.UInt8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c) >> uint64(d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh16x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh16x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh16x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 8 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 8 {
			break
		}
		v.reset(OpSignExt8to16)
		v0 := b.NewValue0(v.Pos, OpTrunc16to8, typ.Int8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh32Ux64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh32Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh32Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpRsh32Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 24 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 24 {
			break
		}
		v.reset(OpZeroExt8to32)
		v0 := b.NewValue0(v.Pos, OpTrunc32to8, typ.UInt8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 16 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		v.reset(OpZeroExt16to32)
		v0 := b.NewValue0(v.Pos, OpTrunc32to16, typ.UInt16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c) >> uint64(d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh32x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 24 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 24 {
			break
		}
		v.reset(OpSignExt8to32)
		v0 := b.NewValue0(v.Pos, OpTrunc32to8, typ.Int8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 16 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		v.reset(OpSignExt16to32)
		v0 := b.NewValue0(v.Pos, OpTrunc32to16, typ.Int16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64Ux16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64Ux32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(uint64(c) >> uint64(d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh64Ux64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh64Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh64Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpRsh64Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 56 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 56 {
			break
		}
		v.reset(OpZeroExt8to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to8, typ.UInt8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 48 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 48 {
			break
		}
		v.reset(OpZeroExt16to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to16, typ.UInt16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 32 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 32 {
			break
		}
		v.reset(OpZeroExt32to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64Ux8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c >> uint64(d)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 56 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 56 {
			break
		}
		v.reset(OpSignExt8to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to8, typ.Int8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 48 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 48 {
			break
		}
		v.reset(OpSignExt16to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to16, typ.Int16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		if v_0_1.AuxInt != 32 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 32 {
			break
		}
		v.reset(OpSignExt32to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(uint8(c) >> uint64(d)))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh8Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLsh8x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpRsh8Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c) >> uint64(d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpRsh8x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh8x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt16to32_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int16(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc32to16 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh32x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 16) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt16to64_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(int16(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc64to16 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 48) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt32to64_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(int32(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc64to32 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 32) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt8to16_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int8(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc16to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh16x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt8to32_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int8(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc32to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh32x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 24) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt8to64_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(int8(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc64to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 56) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSliceCap_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_2 := v_0.Args[2]
		if v_0_2.Op != OpConst64 {
			break
		}
		t := v_0_2.Type
		c := v_0_2.AuxInt
		v.reset(OpConst64)
		v.Type = t
		v.AuxInt = c
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_2 := v_0.Args[2]
		if v_0_2.Op != OpConst32 {
			break
		}
		t := v_0_2.Type
		c := v_0_2.AuxInt
		v.reset(OpConst32)
		v.Type = t
		v.AuxInt = c
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_2 := v_0.Args[2]
		if v_0_2.Op != OpSliceCap {
			break
		}
		x := v_0_2.Args[0]
		v.reset(OpSliceCap)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_2 := v_0.Args[2]
		if v_0_2.Op != OpSliceLen {
			break
		}
		x := v_0_2.Args[0]
		v.reset(OpSliceLen)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSliceLen_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		c := v_0_1.AuxInt
		v.reset(OpConst64)
		v.Type = t
		v.AuxInt = c
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		t := v_0_1.Type
		c := v_0_1.AuxInt
		v.reset(OpConst32)
		v.Type = t
		v.AuxInt = c
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpSliceLen {
			break
		}
		x := v_0_1.Args[0]
		v.reset(OpSliceLen)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSlicePtr_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpSlicePtr {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpSlicePtr)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSlicemask_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		x := v_0.AuxInt
		if !(x > 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = -1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		x := v_0.AuxInt
		if !(x > 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = -1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpSqrt_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = f2i(math.Sqrt(i2f(c)))
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpStaticCall_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		sym := v.Aux
		s1 := v.Args[0]
		if s1.Op != OpStore {
			break
		}
		_ = s1.Args[2]
		s1_1 := s1.Args[1]
		if s1_1.Op != OpConst64 {
			break
		}
		sz := s1_1.AuxInt
		s2 := s1.Args[2]
		if s2.Op != OpStore {
			break
		}
		_ = s2.Args[2]
		src := s2.Args[1]
		s3 := s2.Args[2]
		if s3.Op != OpStore {
			break
		}
		t := s3.Aux
		_ = s3.Args[2]
		dst := s3.Args[1]
		mem := s3.Args[2]
		if !(isSameSym(sym, "runtime.memmove") && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && isInlinableMemmove(dst, src, sz, config) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = sz
		v.Aux = t.(*types.Type).Elem(psess.types)
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}

	for {
		sym := v.Aux
		s1 := v.Args[0]
		if s1.Op != OpStore {
			break
		}
		_ = s1.Args[2]
		s1_1 := s1.Args[1]
		if s1_1.Op != OpConst32 {
			break
		}
		sz := s1_1.AuxInt
		s2 := s1.Args[2]
		if s2.Op != OpStore {
			break
		}
		_ = s2.Args[2]
		src := s2.Args[1]
		s3 := s2.Args[2]
		if s3.Op != OpStore {
			break
		}
		t := s3.Aux
		_ = s3.Args[2]
		dst := s3.Args[1]
		mem := s3.Args[2]
		if !(isSameSym(sym, "runtime.memmove") && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && isInlinableMemmove(dst, src, sz, config) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = sz
		v.Aux = t.(*types.Type).Elem(psess.types)
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpStore_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t1 := v.Aux
		_ = v.Args[2]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpLoad {
			break
		}
		t2 := v_1.Type
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		mem := v_1.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(isSamePtr(p1, p2) && t2.Size(psess.types) == psess.sizeof(t1)) {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		t1 := v.Aux
		_ = v.Args[2]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpLoad {
			break
		}
		t2 := v_1.Type
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		oldmem := v_1.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t3 := mem.Aux
		_ = mem.Args[2]
		p3 := mem.Args[0]
		if oldmem != mem.Args[2] {
			break
		}
		if !(isSamePtr(p1, p2) && t2.Size(psess.types) == psess.sizeof(t1) && disjoint(p1, psess.sizeof(t1), p3, psess.sizeof(t3))) {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		t1 := v.Aux
		_ = v.Args[2]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpLoad {
			break
		}
		t2 := v_1.Type
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		oldmem := v_1.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t3 := mem.Aux
		_ = mem.Args[2]
		p3 := mem.Args[0]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t4 := mem_2.Aux
		_ = mem_2.Args[2]
		p4 := mem_2.Args[0]
		if oldmem != mem_2.Args[2] {
			break
		}
		if !(isSamePtr(p1, p2) && t2.Size(psess.types) == psess.sizeof(t1) && disjoint(p1, psess.sizeof(t1), p3, psess.sizeof(t3)) && disjoint(p1, psess.sizeof(t1), p4, psess.sizeof(t4))) {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		t1 := v.Aux
		_ = v.Args[2]
		p1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpLoad {
			break
		}
		t2 := v_1.Type
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		oldmem := v_1.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t3 := mem.Aux
		_ = mem.Args[2]
		p3 := mem.Args[0]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t4 := mem_2.Aux
		_ = mem_2.Args[2]
		p4 := mem_2.Args[0]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t5 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		p5 := mem_2_2.Args[0]
		if oldmem != mem_2_2.Args[2] {
			break
		}
		if !(isSamePtr(p1, p2) && t2.Size(psess.types) == psess.sizeof(t1) && disjoint(p1, psess.sizeof(t1), p3, psess.sizeof(t3)) && disjoint(p1, psess.sizeof(t1), p4, psess.sizeof(t4)) && disjoint(p1, psess.sizeof(t1), p5, psess.sizeof(t5))) {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		x := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpZero {
			break
		}
		n := mem.AuxInt
		_ = mem.Args[1]
		p2 := mem.Args[0]
		if !(isConstZero(x) && o >= 0 && psess.sizeof(t)+o <= n && isSamePtr(p1, p2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		t1 := v.Aux
		_ = v.Args[2]
		op := v.Args[0]
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		x := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		p2 := mem.Args[0]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpZero {
			break
		}
		n := mem_2.AuxInt
		_ = mem_2.Args[1]
		p3 := mem_2.Args[0]
		if !(isConstZero(x) && o1 >= 0 && psess.sizeof(t1)+o1 <= n && isSamePtr(p1, p3) && disjoint(op, psess.sizeof(t1), p2, psess.sizeof(t2))) {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		t1 := v.Aux
		_ = v.Args[2]
		op := v.Args[0]
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		x := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		p2 := mem.Args[0]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		p3 := mem_2.Args[0]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpZero {
			break
		}
		n := mem_2_2.AuxInt
		_ = mem_2_2.Args[1]
		p4 := mem_2_2.Args[0]
		if !(isConstZero(x) && o1 >= 0 && psess.sizeof(t1)+o1 <= n && isSamePtr(p1, p4) && disjoint(op, psess.sizeof(t1), p2, psess.sizeof(t2)) && disjoint(op, psess.sizeof(t1), p3, psess.sizeof(t3))) {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		t1 := v.Aux
		_ = v.Args[2]
		op := v.Args[0]
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		x := v.Args[1]
		mem := v.Args[2]
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		p2 := mem.Args[0]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		p3 := mem_2.Args[0]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		p4 := mem_2_2.Args[0]
		mem_2_2_2 := mem_2_2.Args[2]
		if mem_2_2_2.Op != OpZero {
			break
		}
		n := mem_2_2_2.AuxInt
		_ = mem_2_2_2.Args[1]
		p5 := mem_2_2_2.Args[0]
		if !(isConstZero(x) && o1 >= 0 && psess.sizeof(t1)+o1 <= n && isSamePtr(p1, p5) && disjoint(op, psess.sizeof(t1), p2, psess.sizeof(t2)) && disjoint(op, psess.sizeof(t1), p3, psess.sizeof(t3)) && disjoint(op, psess.sizeof(t1), p4, psess.sizeof(t4))) {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_1 := v.Args[1]
		if v_1.Op != OpStructMake0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStructMake1 {
			break
		}
		t := v_1.Type
		f0 := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpStore)
		v.Aux = t.FieldType(psess.types, 0)
		v0 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 0).PtrTo(psess.types))
		v0.AuxInt = 0
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(f0)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpStore_10(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe

	for {
		_ = v.Args[2]
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStructMake2 {
			break
		}
		t := v_1.Type
		_ = v_1.Args[1]
		f0 := v_1.Args[0]
		f1 := v_1.Args[1]
		mem := v.Args[2]
		v.reset(OpStore)
		v.Aux = t.FieldType(psess.types, 1)
		v0 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 1).PtrTo(psess.types))
		v0.AuxInt = t.FieldOff(psess.types, 1)
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(f1)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t.FieldType(psess.types, 0)
		v2 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 0).PtrTo(psess.types))
		v2.AuxInt = 0
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(f0)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[2]
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStructMake3 {
			break
		}
		t := v_1.Type
		_ = v_1.Args[2]
		f0 := v_1.Args[0]
		f1 := v_1.Args[1]
		f2 := v_1.Args[2]
		mem := v.Args[2]
		v.reset(OpStore)
		v.Aux = t.FieldType(psess.types, 2)
		v0 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 2).PtrTo(psess.types))
		v0.AuxInt = t.FieldOff(psess.types, 2)
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(f2)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t.FieldType(psess.types, 1)
		v2 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 1).PtrTo(psess.types))
		v2.AuxInt = t.FieldOff(psess.types, 1)
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(f1)
		v3 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v3.Aux = t.FieldType(psess.types, 0)
		v4 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 0).PtrTo(psess.types))
		v4.AuxInt = 0
		v4.AddArg(dst)
		v3.AddArg(v4)
		v3.AddArg(f0)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[2]
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStructMake4 {
			break
		}
		t := v_1.Type
		_ = v_1.Args[3]
		f0 := v_1.Args[0]
		f1 := v_1.Args[1]
		f2 := v_1.Args[2]
		f3 := v_1.Args[3]
		mem := v.Args[2]
		v.reset(OpStore)
		v.Aux = t.FieldType(psess.types, 3)
		v0 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 3).PtrTo(psess.types))
		v0.AuxInt = t.FieldOff(psess.types, 3)
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(f3)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t.FieldType(psess.types, 2)
		v2 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 2).PtrTo(psess.types))
		v2.AuxInt = t.FieldOff(psess.types, 2)
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(f2)
		v3 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v3.Aux = t.FieldType(psess.types, 1)
		v4 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 1).PtrTo(psess.types))
		v4.AuxInt = t.FieldOff(psess.types, 1)
		v4.AddArg(dst)
		v3.AddArg(v4)
		v3.AddArg(f1)
		v5 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v5.Aux = t.FieldType(psess.types, 0)
		v6 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(psess.types, 0).PtrTo(psess.types))
		v6.AuxInt = 0
		v6.AddArg(dst)
		v5.AddArg(v6)
		v5.AddArg(f0)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpLoad {
			break
		}
		_ = v_1.Args[1]
		src := v_1.Args[0]
		mem := v_1.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(!fe.CanSSA(t.(*types.Type))) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = psess.sizeof(t)
		v.Aux = t
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpLoad {
			break
		}
		_ = v_1.Args[1]
		src := v_1.Args[0]
		mem := v_1.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpVarDef {
			break
		}
		x := v_2.Aux
		if mem != v_2.Args[0] {
			break
		}
		if !(!fe.CanSSA(t.(*types.Type))) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = psess.sizeof(t)
		v.Aux = t
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpVarDef, psess.types.TypeMem)
		v0.Aux = x
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		v_1 := v.Args[1]
		if v_1.Op != OpArrayMake0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpArrayMake1 {
			break
		}
		e := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpStore)
		v.Aux = e.Type
		v.AddArg(dst)
		v.AddArg(e)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpLoad {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpOffPtr {
			break
		}
		c := v_0_0.AuxInt
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpSP {
			break
		}
		mem := v_0.Args[1]
		x := v.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(isConstZero(x) && mem.Op == OpStaticCall && isSameSym(mem.Aux, "runtime.newobject") && c == config.ctxt.FixedFrameSize()+config.RegSize) {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpOffPtr {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLoad {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpOffPtr {
			break
		}
		c := v_0_0_0.AuxInt
		v_0_0_0_0 := v_0_0_0.Args[0]
		if v_0_0_0_0.Op != OpSP {
			break
		}
		mem := v_0_0.Args[1]
		x := v.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(isConstZero(x) && mem.Op == OpStaticCall && isSameSym(mem.Aux, "runtime.newobject") && c == config.ctxt.FixedFrameSize()+config.RegSize) {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		t1 := v.Aux
		_ = v.Args[2]
		op1 := v.Args[0]
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v.Args[1]
		m2 := v.Args[2]
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		if op2.AuxInt != 0 {
			break
		}
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpMove {
			break
		}
		n := m3.AuxInt
		_ = m3.Args[2]
		p3 := m3.Args[0]
		mem := m3.Args[2]
		if !(m2.Uses == 1 && m3.Uses == 1 && o1 == psess.sizeof(t2) && n == psess.sizeof(t2)+psess.sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && clobber(m2) && clobber(m3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v.AddArg(op1)
		v.AddArg(d1)
		v0 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v0.Aux = t2
		v0.AddArg(op2)
		v0.AddArg(d2)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpStore_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t1 := v.Aux
		_ = v.Args[2]
		op1 := v.Args[0]
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v.Args[1]
		m2 := v.Args[2]
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpStore {
			break
		}
		t3 := m3.Aux
		_ = m3.Args[2]
		op3 := m3.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		if op3.AuxInt != 0 {
			break
		}
		p3 := op3.Args[0]
		d3 := m3.Args[1]
		m4 := m3.Args[2]
		if m4.Op != OpMove {
			break
		}
		n := m4.AuxInt
		_ = m4.Args[2]
		p4 := m4.Args[0]
		mem := m4.Args[2]
		if !(m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && o2 == psess.sizeof(t3) && o1-o2 == psess.sizeof(t2) && n == psess.sizeof(t3)+psess.sizeof(t2)+psess.sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && clobber(m2) && clobber(m3) && clobber(m4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v.AddArg(op1)
		v.AddArg(d1)
		v0 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v0.Aux = t2
		v0.AddArg(op2)
		v0.AddArg(d2)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v1.AddArg(op3)
		v1.AddArg(d3)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t1 := v.Aux
		_ = v.Args[2]
		op1 := v.Args[0]
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v.Args[1]
		m2 := v.Args[2]
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpStore {
			break
		}
		t3 := m3.Aux
		_ = m3.Args[2]
		op3 := m3.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d3 := m3.Args[1]
		m4 := m3.Args[2]
		if m4.Op != OpStore {
			break
		}
		t4 := m4.Aux
		_ = m4.Args[2]
		op4 := m4.Args[0]
		if op4.Op != OpOffPtr {
			break
		}
		if op4.AuxInt != 0 {
			break
		}
		p4 := op4.Args[0]
		d4 := m4.Args[1]
		m5 := m4.Args[2]
		if m5.Op != OpMove {
			break
		}
		n := m5.AuxInt
		_ = m5.Args[2]
		p5 := m5.Args[0]
		mem := m5.Args[2]
		if !(m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && m5.Uses == 1 && o3 == psess.sizeof(t4) && o2-o3 == psess.sizeof(t3) && o1-o2 == psess.sizeof(t2) && n == psess.sizeof(t4)+psess.sizeof(t3)+psess.sizeof(t2)+psess.sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && clobber(m2) && clobber(m3) && clobber(m4) && clobber(m5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v.AddArg(op1)
		v.AddArg(d1)
		v0 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v0.Aux = t2
		v0.AddArg(op2)
		v0.AddArg(d2)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v1.AddArg(op3)
		v1.AddArg(d3)
		v2 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v2.Aux = t4
		v2.AddArg(op4)
		v2.AddArg(d4)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t1 := v.Aux
		_ = v.Args[2]
		op1 := v.Args[0]
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v.Args[1]
		m2 := v.Args[2]
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		if op2.AuxInt != 0 {
			break
		}
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpZero {
			break
		}
		n := m3.AuxInt
		_ = m3.Args[1]
		p3 := m3.Args[0]
		mem := m3.Args[1]
		if !(m2.Uses == 1 && m3.Uses == 1 && o1 == psess.sizeof(t2) && n == psess.sizeof(t2)+psess.sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && clobber(m2) && clobber(m3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v.AddArg(op1)
		v.AddArg(d1)
		v0 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v0.Aux = t2
		v0.AddArg(op2)
		v0.AddArg(d2)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		t1 := v.Aux
		_ = v.Args[2]
		op1 := v.Args[0]
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v.Args[1]
		m2 := v.Args[2]
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpStore {
			break
		}
		t3 := m3.Aux
		_ = m3.Args[2]
		op3 := m3.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		if op3.AuxInt != 0 {
			break
		}
		p3 := op3.Args[0]
		d3 := m3.Args[1]
		m4 := m3.Args[2]
		if m4.Op != OpZero {
			break
		}
		n := m4.AuxInt
		_ = m4.Args[1]
		p4 := m4.Args[0]
		mem := m4.Args[1]
		if !(m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && o2 == psess.sizeof(t3) && o1-o2 == psess.sizeof(t2) && n == psess.sizeof(t3)+psess.sizeof(t2)+psess.sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && clobber(m2) && clobber(m3) && clobber(m4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v.AddArg(op1)
		v.AddArg(d1)
		v0 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v0.Aux = t2
		v0.AddArg(op2)
		v0.AddArg(d2)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v1.AddArg(op3)
		v1.AddArg(d3)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t1 := v.Aux
		_ = v.Args[2]
		op1 := v.Args[0]
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v.Args[1]
		m2 := v.Args[2]
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpStore {
			break
		}
		t3 := m3.Aux
		_ = m3.Args[2]
		op3 := m3.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d3 := m3.Args[1]
		m4 := m3.Args[2]
		if m4.Op != OpStore {
			break
		}
		t4 := m4.Aux
		_ = m4.Args[2]
		op4 := m4.Args[0]
		if op4.Op != OpOffPtr {
			break
		}
		if op4.AuxInt != 0 {
			break
		}
		p4 := op4.Args[0]
		d4 := m4.Args[1]
		m5 := m4.Args[2]
		if m5.Op != OpZero {
			break
		}
		n := m5.AuxInt
		_ = m5.Args[1]
		p5 := m5.Args[0]
		mem := m5.Args[1]
		if !(m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && m5.Uses == 1 && o3 == psess.sizeof(t4) && o2-o3 == psess.sizeof(t3) && o1-o2 == psess.sizeof(t2) && n == psess.sizeof(t4)+psess.sizeof(t3)+psess.sizeof(t2)+psess.sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && clobber(m2) && clobber(m3) && clobber(m4) && clobber(m5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v.AddArg(op1)
		v.AddArg(d1)
		v0 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v0.Aux = t2
		v0.AddArg(op2)
		v0.AddArg(d2)
		v1 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v1.Aux = t3
		v1.AddArg(op3)
		v1.AddArg(d3)
		v2 := b.NewValue0(v.Pos, OpStore, psess.types.TypeMem)
		v2.Aux = t4
		v2.AddArg(op4)
		v2.AddArg(d4)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpStringLen_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpStringMake {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		c := v_0_1.AuxInt
		v.reset(OpConst64)
		v.Type = t
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValuegeneric_OpStringPtr_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpStringMake {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		t := v_0_0.Type
		c := v_0_0.AuxInt
		v.reset(OpConst64)
		v.Type = t
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValuegeneric_OpStructSelect_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpStructMake1 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpStructMake2 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpStructMake2 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpStructMake3 {
			break
		}
		_ = v_0.Args[2]
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpStructMake3 {
			break
		}
		_ = v_0.Args[2]
		x := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpStructMake3 {
			break
		}
		_ = v_0.Args[2]
		x := v_0.Args[2]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpStructMake4 {
			break
		}
		_ = v_0.Args[3]
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpStructMake4 {
			break
		}
		_ = v_0.Args[3]
		x := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpStructMake4 {
			break
		}
		_ = v_0.Args[3]
		x := v_0.Args[2]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 3 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpStructMake4 {
			break
		}
		_ = v_0.Args[3]
		x := v_0.Args[3]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpStructSelect_10(v *Value) bool {
	b := v.Block
	_ = b
	fe := b.Func.fe
	_ = fe

	for {
		i := v.AuxInt
		x := v.Args[0]
		if x.Op != OpLoad {
			break
		}
		t := x.Type
		_ = x.Args[1]
		ptr := x.Args[0]
		mem := x.Args[1]
		if !(!fe.CanSSA(t)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpLoad, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, v.Type.PtrTo(psess.types))
		v1.AuxInt = t.FieldOff(psess.types, int(i))
		v1.AddArg(ptr)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		if x.Op != OpIData {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c - d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(-c))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul16)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub16, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if y != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub16_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if y != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub16, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst16 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst16 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c - d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(-c))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul32)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub32, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if y != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub32_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if y != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub32, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = f2i(float64(i2f32(c) - i2f32(d)))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32F {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c - d
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = -c
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub64, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if y != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub64_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if y != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub64, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = f2i(i2f(c) - i2f(d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64F {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c - d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(-c))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		z := v_1.Args[1]
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMul8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpMul8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpSub8, t)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if y != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub8_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if y != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(i)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpSub8, t)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst8 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c + d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst8 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpTrunc16to8_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to16 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSignExt8to16 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst16 {
			break
		}
		y := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(y&0xFF == 0xFF) {
			break
		}
		v.reset(OpTrunc16to8)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst16 {
			break
		}
		y := v_0_1.AuxInt
		if !(y&0xFF == 0xFF) {
			break
		}
		v.reset(OpTrunc16to8)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpTrunc32to16_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to32 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpZeroExt8to16)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt16to32 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSignExt8to32 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpSignExt8to16)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSignExt16to32 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		y := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(y&0xFFFF == 0xFFFF) {
			break
		}
		v.reset(OpTrunc32to16)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		y := v_0_1.AuxInt
		if !(y&0xFFFF == 0xFFFF) {
			break
		}
		v.reset(OpTrunc32to16)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpTrunc32to8_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to32 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSignExt8to32 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		y := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(y&0xFF == 0xFF) {
			break
		}
		v.reset(OpTrunc32to8)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		y := v_0_1.AuxInt
		if !(y&0xFF == 0xFF) {
			break
		}
		v.reset(OpTrunc32to8)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpTrunc64to16_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpZeroExt8to16)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt16to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSignExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpSignExt8to16)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSignExt16to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		y := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(y&0xFFFF == 0xFFFF) {
			break
		}
		v.reset(OpTrunc64to16)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		y := v_0_1.AuxInt
		if !(y&0xFFFF == 0xFFFF) {
			break
		}
		v.reset(OpTrunc64to16)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpTrunc64to32_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpZeroExt8to32)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt16to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpZeroExt16to32)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt32to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSignExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpSignExt8to32)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSignExt16to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpSignExt16to32)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSignExt32to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		y := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(y&0xFFFFFFFF == 0xFFFFFFFF) {
			break
		}
		v.reset(OpTrunc64to32)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		y := v_0_1.AuxInt
		if !(y&0xFFFFFFFF == 0xFFFFFFFF) {
			break
		}
		v.reset(OpTrunc64to32)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpTrunc64to8_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSignExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		y := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(y&0xFF == 0xFF) {
			break
		}
		v.reset(OpTrunc64to8)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		y := v_0_1.AuxInt
		if !(y&0xFF == 0xFF) {
			break
		}
		v.reset(OpTrunc64to8)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpXor16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c ^ d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c ^ d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor16 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor16 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor16 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor16 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpXor16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpXor16_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor16 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpXor16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor16 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpXor16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpXor16)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor16, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpXor16 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst16 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpXor16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpXor16 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst16 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpXor16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst16 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpXor16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor16 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst16 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpXor16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpXor32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c ^ d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c ^ d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor32 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor32 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor32 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor32 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpXor32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpXor32_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor32 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpXor32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor32 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpXor32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpXor32)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor32, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpXor32 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpXor32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpXor32 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpXor32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst32 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpXor32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor32 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpXor32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpXor64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c ^ d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c ^ d
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor64 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor64 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor64 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor64 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpXor64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpXor64_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor64 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpXor64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor64 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpXor64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpXor64)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor64, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpXor64 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpXor64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c ^ d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpXor64 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpXor64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c ^ d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpXor64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c ^ d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpXor64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c ^ d
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpXor8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c ^ d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c ^ d))
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor8 {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor8 {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor8 {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor8 {
			break
		}
		_ = v_0.Args[1]
		i := v_0.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_0.Args[1]
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpXor8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpXor8_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor8 {
			break
		}
		_ = v_0.Args[1]
		z := v_0.Args[0]
		i := v_0.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		x := v.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpXor8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor8 {
			break
		}
		_ = v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		z := v_1.Args[1]
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpXor8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpXor8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpXor8)
		v.AddArg(i)
		v0 := b.NewValue0(v.Pos, OpXor8, t)
		v0.AddArg(z)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpXor8 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst8 {
			break
		}
		if v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		x := v_1.Args[1]
		v.reset(OpXor8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpXor8 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst8 {
			break
		}
		if v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpXor8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst8 {
			break
		}
		t := v_0_0.Type
		d := v_0_0.AuxInt
		x := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpXor8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpXor8 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst8 {
			break
		}
		t := v_0_1.Type
		d := v_0_1.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 {
			break
		}
		if v_1.Type != t {
			break
		}
		c := v_1.AuxInt
		v.reset(OpXor8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c ^ d))
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuegeneric_OpZero_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpLoad {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpOffPtr {
			break
		}
		c := v_0_0.AuxInt
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpSP {
			break
		}
		mem := v_0.Args[1]
		if mem != v.Args[1] {
			break
		}
		if !(mem.Op == OpStaticCall && isSameSym(mem.Aux, "runtime.newobject") && c == config.ctxt.FixedFrameSize()+config.RegSize) {
			break
		}
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		n := v.AuxInt
		t1 := v.Aux
		_ = v.Args[1]
		p1 := v.Args[0]
		store := v.Args[1]
		if store.Op != OpStore {
			break
		}
		t2 := store.Aux
		_ = store.Args[2]
		store_0 := store.Args[0]
		if store_0.Op != OpOffPtr {
			break
		}
		o2 := store_0.AuxInt
		p2 := store_0.Args[0]
		mem := store.Args[2]
		if !(isSamePtr(p1, p2) && store.Uses == 1 && n >= o2+psess.sizeof(t2) && clobber(store)) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = n
		v.Aux = t1
		v.AddArg(p1)
		v.AddArg(mem)
		return true
	}

	for {
		n := v.AuxInt
		t := v.Aux
		_ = v.Args[1]
		dst1 := v.Args[0]
		move := v.Args[1]
		if move.Op != OpMove {
			break
		}
		if move.AuxInt != n {
			break
		}
		if move.Aux != t {
			break
		}
		_ = move.Args[2]
		dst2 := move.Args[0]
		mem := move.Args[2]
		if !(move.Uses == 1 && isSamePtr(dst1, dst2) && clobber(move)) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = n
		v.Aux = t
		v.AddArg(dst1)
		v.AddArg(mem)
		return true
	}

	for {
		n := v.AuxInt
		t := v.Aux
		_ = v.Args[1]
		dst1 := v.Args[0]
		vardef := v.Args[1]
		if vardef.Op != OpVarDef {
			break
		}
		x := vardef.Aux
		move := vardef.Args[0]
		if move.Op != OpMove {
			break
		}
		if move.AuxInt != n {
			break
		}
		if move.Aux != t {
			break
		}
		_ = move.Args[2]
		dst2 := move.Args[0]
		mem := move.Args[2]
		if !(move.Uses == 1 && vardef.Uses == 1 && isSamePtr(dst1, dst2) && clobber(move) && clobber(vardef)) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = n
		v.Aux = t
		v.AddArg(dst1)
		v0 := b.NewValue0(v.Pos, OpVarDef, psess.types.TypeMem)
		v0.Aux = x
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt16to32_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(uint16(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc32to16 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh32Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 16) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt16to64_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(uint16(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc64to16 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 48) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt32to64_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(uint32(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc64to32 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 32) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt8to16_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(uint8(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc16to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh16Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 8) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt8to32_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(uint8(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc32to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh32Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 24) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt8to64_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(uint8(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpTrunc64to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 56) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteBlockgeneric(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	case BlockIf:

		for {
			v := b.Control
			if v.Op != OpNot {
				break
			}
			cond := v.Args[0]
			b.Kind = BlockIf
			b.SetControl(cond)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}

		for {
			v := b.Control
			if v.Op != OpConstBool {
				break
			}
			c := v.AuxInt
			if !(c == 1) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpConstBool {
				break
			}
			c := v.AuxInt
			if !(c == 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	}
	return false
}

package ssa

import "math"

import "github.com/dave/golib/src/cmd/compile/internal/types"

// in case not otherwise used
// in case not otherwise used
// in case not otherwise used
// in case not otherwise used

func (psess *PackageSession) rewriteValuePPC64(v *Value) bool {
	switch v.Op {
	case OpAbs:
		return rewriteValuePPC64_OpAbs_0(v)
	case OpAdd16:
		return rewriteValuePPC64_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValuePPC64_OpAdd32_0(v)
	case OpAdd32F:
		return rewriteValuePPC64_OpAdd32F_0(v)
	case OpAdd64:
		return rewriteValuePPC64_OpAdd64_0(v)
	case OpAdd64F:
		return rewriteValuePPC64_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValuePPC64_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValuePPC64_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValuePPC64_OpAddr_0(v)
	case OpAnd16:
		return rewriteValuePPC64_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValuePPC64_OpAnd32_0(v)
	case OpAnd64:
		return rewriteValuePPC64_OpAnd64_0(v)
	case OpAnd8:
		return rewriteValuePPC64_OpAnd8_0(v)
	case OpAndB:
		return rewriteValuePPC64_OpAndB_0(v)
	case OpAtomicAdd32:
		return rewriteValuePPC64_OpAtomicAdd32_0(v)
	case OpAtomicAdd64:
		return rewriteValuePPC64_OpAtomicAdd64_0(v)
	case OpAtomicAnd8:
		return rewriteValuePPC64_OpAtomicAnd8_0(v)
	case OpAtomicCompareAndSwap32:
		return rewriteValuePPC64_OpAtomicCompareAndSwap32_0(v)
	case OpAtomicCompareAndSwap64:
		return rewriteValuePPC64_OpAtomicCompareAndSwap64_0(v)
	case OpAtomicExchange32:
		return rewriteValuePPC64_OpAtomicExchange32_0(v)
	case OpAtomicExchange64:
		return rewriteValuePPC64_OpAtomicExchange64_0(v)
	case OpAtomicLoad32:
		return rewriteValuePPC64_OpAtomicLoad32_0(v)
	case OpAtomicLoad64:
		return rewriteValuePPC64_OpAtomicLoad64_0(v)
	case OpAtomicLoadPtr:
		return rewriteValuePPC64_OpAtomicLoadPtr_0(v)
	case OpAtomicOr8:
		return rewriteValuePPC64_OpAtomicOr8_0(v)
	case OpAtomicStore32:
		return rewriteValuePPC64_OpAtomicStore32_0(v)
	case OpAtomicStore64:
		return rewriteValuePPC64_OpAtomicStore64_0(v)
	case OpAvg64u:
		return rewriteValuePPC64_OpAvg64u_0(v)
	case OpBitLen32:
		return rewriteValuePPC64_OpBitLen32_0(v)
	case OpBitLen64:
		return rewriteValuePPC64_OpBitLen64_0(v)
	case OpCeil:
		return rewriteValuePPC64_OpCeil_0(v)
	case OpClosureCall:
		return rewriteValuePPC64_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValuePPC64_OpCom16_0(v)
	case OpCom32:
		return rewriteValuePPC64_OpCom32_0(v)
	case OpCom64:
		return rewriteValuePPC64_OpCom64_0(v)
	case OpCom8:
		return rewriteValuePPC64_OpCom8_0(v)
	case OpConst16:
		return rewriteValuePPC64_OpConst16_0(v)
	case OpConst32:
		return rewriteValuePPC64_OpConst32_0(v)
	case OpConst32F:
		return rewriteValuePPC64_OpConst32F_0(v)
	case OpConst64:
		return rewriteValuePPC64_OpConst64_0(v)
	case OpConst64F:
		return rewriteValuePPC64_OpConst64F_0(v)
	case OpConst8:
		return rewriteValuePPC64_OpConst8_0(v)
	case OpConstBool:
		return rewriteValuePPC64_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValuePPC64_OpConstNil_0(v)
	case OpCopysign:
		return rewriteValuePPC64_OpCopysign_0(v)
	case OpCtz32:
		return rewriteValuePPC64_OpCtz32_0(v)
	case OpCtz32NonZero:
		return rewriteValuePPC64_OpCtz32NonZero_0(v)
	case OpCtz64:
		return rewriteValuePPC64_OpCtz64_0(v)
	case OpCtz64NonZero:
		return rewriteValuePPC64_OpCtz64NonZero_0(v)
	case OpCvt32Fto32:
		return rewriteValuePPC64_OpCvt32Fto32_0(v)
	case OpCvt32Fto64:
		return rewriteValuePPC64_OpCvt32Fto64_0(v)
	case OpCvt32Fto64F:
		return rewriteValuePPC64_OpCvt32Fto64F_0(v)
	case OpCvt32to32F:
		return rewriteValuePPC64_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValuePPC64_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValuePPC64_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValuePPC64_OpCvt64Fto32F_0(v)
	case OpCvt64Fto64:
		return rewriteValuePPC64_OpCvt64Fto64_0(v)
	case OpCvt64to32F:
		return rewriteValuePPC64_OpCvt64to32F_0(v)
	case OpCvt64to64F:
		return rewriteValuePPC64_OpCvt64to64F_0(v)
	case OpDiv16:
		return rewriteValuePPC64_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValuePPC64_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValuePPC64_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValuePPC64_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValuePPC64_OpDiv32u_0(v)
	case OpDiv64:
		return rewriteValuePPC64_OpDiv64_0(v)
	case OpDiv64F:
		return rewriteValuePPC64_OpDiv64F_0(v)
	case OpDiv64u:
		return rewriteValuePPC64_OpDiv64u_0(v)
	case OpDiv8:
		return rewriteValuePPC64_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValuePPC64_OpDiv8u_0(v)
	case OpEq16:
		return psess.rewriteValuePPC64_OpEq16_0(v)
	case OpEq32:
		return psess.rewriteValuePPC64_OpEq32_0(v)
	case OpEq32F:
		return psess.rewriteValuePPC64_OpEq32F_0(v)
	case OpEq64:
		return psess.rewriteValuePPC64_OpEq64_0(v)
	case OpEq64F:
		return psess.rewriteValuePPC64_OpEq64F_0(v)
	case OpEq8:
		return psess.rewriteValuePPC64_OpEq8_0(v)
	case OpEqB:
		return rewriteValuePPC64_OpEqB_0(v)
	case OpEqPtr:
		return psess.rewriteValuePPC64_OpEqPtr_0(v)
	case OpFloor:
		return rewriteValuePPC64_OpFloor_0(v)
	case OpGeq16:
		return psess.rewriteValuePPC64_OpGeq16_0(v)
	case OpGeq16U:
		return psess.rewriteValuePPC64_OpGeq16U_0(v)
	case OpGeq32:
		return psess.rewriteValuePPC64_OpGeq32_0(v)
	case OpGeq32F:
		return psess.rewriteValuePPC64_OpGeq32F_0(v)
	case OpGeq32U:
		return psess.rewriteValuePPC64_OpGeq32U_0(v)
	case OpGeq64:
		return psess.rewriteValuePPC64_OpGeq64_0(v)
	case OpGeq64F:
		return psess.rewriteValuePPC64_OpGeq64F_0(v)
	case OpGeq64U:
		return psess.rewriteValuePPC64_OpGeq64U_0(v)
	case OpGeq8:
		return psess.rewriteValuePPC64_OpGeq8_0(v)
	case OpGeq8U:
		return psess.rewriteValuePPC64_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValuePPC64_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValuePPC64_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValuePPC64_OpGetClosurePtr_0(v)
	case OpGreater16:
		return psess.rewriteValuePPC64_OpGreater16_0(v)
	case OpGreater16U:
		return psess.rewriteValuePPC64_OpGreater16U_0(v)
	case OpGreater32:
		return psess.rewriteValuePPC64_OpGreater32_0(v)
	case OpGreater32F:
		return psess.rewriteValuePPC64_OpGreater32F_0(v)
	case OpGreater32U:
		return psess.rewriteValuePPC64_OpGreater32U_0(v)
	case OpGreater64:
		return psess.rewriteValuePPC64_OpGreater64_0(v)
	case OpGreater64F:
		return psess.rewriteValuePPC64_OpGreater64F_0(v)
	case OpGreater64U:
		return psess.rewriteValuePPC64_OpGreater64U_0(v)
	case OpGreater8:
		return psess.rewriteValuePPC64_OpGreater8_0(v)
	case OpGreater8U:
		return psess.rewriteValuePPC64_OpGreater8U_0(v)
	case OpHmul32:
		return rewriteValuePPC64_OpHmul32_0(v)
	case OpHmul32u:
		return rewriteValuePPC64_OpHmul32u_0(v)
	case OpHmul64:
		return rewriteValuePPC64_OpHmul64_0(v)
	case OpHmul64u:
		return rewriteValuePPC64_OpHmul64u_0(v)
	case OpInterCall:
		return rewriteValuePPC64_OpInterCall_0(v)
	case OpIsInBounds:
		return psess.rewriteValuePPC64_OpIsInBounds_0(v)
	case OpIsNonNil:
		return psess.rewriteValuePPC64_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return psess.rewriteValuePPC64_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return psess.rewriteValuePPC64_OpLeq16_0(v)
	case OpLeq16U:
		return psess.rewriteValuePPC64_OpLeq16U_0(v)
	case OpLeq32:
		return psess.rewriteValuePPC64_OpLeq32_0(v)
	case OpLeq32F:
		return psess.rewriteValuePPC64_OpLeq32F_0(v)
	case OpLeq32U:
		return psess.rewriteValuePPC64_OpLeq32U_0(v)
	case OpLeq64:
		return psess.rewriteValuePPC64_OpLeq64_0(v)
	case OpLeq64F:
		return psess.rewriteValuePPC64_OpLeq64F_0(v)
	case OpLeq64U:
		return psess.rewriteValuePPC64_OpLeq64U_0(v)
	case OpLeq8:
		return psess.rewriteValuePPC64_OpLeq8_0(v)
	case OpLeq8U:
		return psess.rewriteValuePPC64_OpLeq8U_0(v)
	case OpLess16:
		return psess.rewriteValuePPC64_OpLess16_0(v)
	case OpLess16U:
		return psess.rewriteValuePPC64_OpLess16U_0(v)
	case OpLess32:
		return psess.rewriteValuePPC64_OpLess32_0(v)
	case OpLess32F:
		return psess.rewriteValuePPC64_OpLess32F_0(v)
	case OpLess32U:
		return psess.rewriteValuePPC64_OpLess32U_0(v)
	case OpLess64:
		return psess.rewriteValuePPC64_OpLess64_0(v)
	case OpLess64F:
		return psess.rewriteValuePPC64_OpLess64F_0(v)
	case OpLess64U:
		return psess.rewriteValuePPC64_OpLess64U_0(v)
	case OpLess8:
		return psess.rewriteValuePPC64_OpLess8_0(v)
	case OpLess8U:
		return psess.rewriteValuePPC64_OpLess8U_0(v)
	case OpLoad:
		return psess.rewriteValuePPC64_OpLoad_0(v)
	case OpLsh16x16:
		return psess.rewriteValuePPC64_OpLsh16x16_0(v)
	case OpLsh16x32:
		return psess.rewriteValuePPC64_OpLsh16x32_0(v)
	case OpLsh16x64:
		return psess.rewriteValuePPC64_OpLsh16x64_0(v)
	case OpLsh16x8:
		return psess.rewriteValuePPC64_OpLsh16x8_0(v)
	case OpLsh32x16:
		return psess.rewriteValuePPC64_OpLsh32x16_0(v)
	case OpLsh32x32:
		return psess.rewriteValuePPC64_OpLsh32x32_0(v)
	case OpLsh32x64:
		return psess.rewriteValuePPC64_OpLsh32x64_0(v)
	case OpLsh32x8:
		return psess.rewriteValuePPC64_OpLsh32x8_0(v)
	case OpLsh64x16:
		return psess.rewriteValuePPC64_OpLsh64x16_0(v)
	case OpLsh64x32:
		return psess.rewriteValuePPC64_OpLsh64x32_0(v)
	case OpLsh64x64:
		return psess.rewriteValuePPC64_OpLsh64x64_0(v)
	case OpLsh64x8:
		return psess.rewriteValuePPC64_OpLsh64x8_0(v)
	case OpLsh8x16:
		return psess.rewriteValuePPC64_OpLsh8x16_0(v)
	case OpLsh8x32:
		return psess.rewriteValuePPC64_OpLsh8x32_0(v)
	case OpLsh8x64:
		return psess.rewriteValuePPC64_OpLsh8x64_0(v)
	case OpLsh8x8:
		return psess.rewriteValuePPC64_OpLsh8x8_0(v)
	case OpMod16:
		return rewriteValuePPC64_OpMod16_0(v)
	case OpMod16u:
		return rewriteValuePPC64_OpMod16u_0(v)
	case OpMod32:
		return rewriteValuePPC64_OpMod32_0(v)
	case OpMod32u:
		return rewriteValuePPC64_OpMod32u_0(v)
	case OpMod64:
		return rewriteValuePPC64_OpMod64_0(v)
	case OpMod64u:
		return rewriteValuePPC64_OpMod64u_0(v)
	case OpMod8:
		return rewriteValuePPC64_OpMod8_0(v)
	case OpMod8u:
		return rewriteValuePPC64_OpMod8u_0(v)
	case OpMove:
		return psess.rewriteValuePPC64_OpMove_0(v) || rewriteValuePPC64_OpMove_10(v)
	case OpMul16:
		return rewriteValuePPC64_OpMul16_0(v)
	case OpMul32:
		return rewriteValuePPC64_OpMul32_0(v)
	case OpMul32F:
		return rewriteValuePPC64_OpMul32F_0(v)
	case OpMul64:
		return rewriteValuePPC64_OpMul64_0(v)
	case OpMul64F:
		return rewriteValuePPC64_OpMul64F_0(v)
	case OpMul8:
		return rewriteValuePPC64_OpMul8_0(v)
	case OpNeg16:
		return rewriteValuePPC64_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValuePPC64_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValuePPC64_OpNeg32F_0(v)
	case OpNeg64:
		return rewriteValuePPC64_OpNeg64_0(v)
	case OpNeg64F:
		return rewriteValuePPC64_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValuePPC64_OpNeg8_0(v)
	case OpNeq16:
		return psess.rewriteValuePPC64_OpNeq16_0(v)
	case OpNeq32:
		return psess.rewriteValuePPC64_OpNeq32_0(v)
	case OpNeq32F:
		return psess.rewriteValuePPC64_OpNeq32F_0(v)
	case OpNeq64:
		return psess.rewriteValuePPC64_OpNeq64_0(v)
	case OpNeq64F:
		return psess.rewriteValuePPC64_OpNeq64F_0(v)
	case OpNeq8:
		return psess.rewriteValuePPC64_OpNeq8_0(v)
	case OpNeqB:
		return rewriteValuePPC64_OpNeqB_0(v)
	case OpNeqPtr:
		return psess.rewriteValuePPC64_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValuePPC64_OpNilCheck_0(v)
	case OpNot:
		return rewriteValuePPC64_OpNot_0(v)
	case OpOffPtr:
		return rewriteValuePPC64_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValuePPC64_OpOr16_0(v)
	case OpOr32:
		return rewriteValuePPC64_OpOr32_0(v)
	case OpOr64:
		return rewriteValuePPC64_OpOr64_0(v)
	case OpOr8:
		return rewriteValuePPC64_OpOr8_0(v)
	case OpOrB:
		return rewriteValuePPC64_OpOrB_0(v)
	case OpPPC64ADD:
		return rewriteValuePPC64_OpPPC64ADD_0(v)
	case OpPPC64ADDconst:
		return rewriteValuePPC64_OpPPC64ADDconst_0(v)
	case OpPPC64AND:
		return rewriteValuePPC64_OpPPC64AND_0(v)
	case OpPPC64ANDconst:
		return rewriteValuePPC64_OpPPC64ANDconst_0(v)
	case OpPPC64CMP:
		return psess.rewriteValuePPC64_OpPPC64CMP_0(v)
	case OpPPC64CMPU:
		return psess.rewriteValuePPC64_OpPPC64CMPU_0(v)
	case OpPPC64CMPUconst:
		return rewriteValuePPC64_OpPPC64CMPUconst_0(v)
	case OpPPC64CMPW:
		return psess.rewriteValuePPC64_OpPPC64CMPW_0(v)
	case OpPPC64CMPWU:
		return psess.rewriteValuePPC64_OpPPC64CMPWU_0(v)
	case OpPPC64CMPWUconst:
		return rewriteValuePPC64_OpPPC64CMPWUconst_0(v)
	case OpPPC64CMPWconst:
		return rewriteValuePPC64_OpPPC64CMPWconst_0(v)
	case OpPPC64CMPconst:
		return rewriteValuePPC64_OpPPC64CMPconst_0(v)
	case OpPPC64Equal:
		return rewriteValuePPC64_OpPPC64Equal_0(v)
	case OpPPC64FABS:
		return rewriteValuePPC64_OpPPC64FABS_0(v)
	case OpPPC64FADD:
		return rewriteValuePPC64_OpPPC64FADD_0(v)
	case OpPPC64FADDS:
		return rewriteValuePPC64_OpPPC64FADDS_0(v)
	case OpPPC64FCEIL:
		return rewriteValuePPC64_OpPPC64FCEIL_0(v)
	case OpPPC64FFLOOR:
		return rewriteValuePPC64_OpPPC64FFLOOR_0(v)
	case OpPPC64FMOVDload:
		return rewriteValuePPC64_OpPPC64FMOVDload_0(v)
	case OpPPC64FMOVDstore:
		return rewriteValuePPC64_OpPPC64FMOVDstore_0(v)
	case OpPPC64FMOVSload:
		return rewriteValuePPC64_OpPPC64FMOVSload_0(v)
	case OpPPC64FMOVSstore:
		return rewriteValuePPC64_OpPPC64FMOVSstore_0(v)
	case OpPPC64FNEG:
		return rewriteValuePPC64_OpPPC64FNEG_0(v)
	case OpPPC64FSQRT:
		return rewriteValuePPC64_OpPPC64FSQRT_0(v)
	case OpPPC64FSUB:
		return rewriteValuePPC64_OpPPC64FSUB_0(v)
	case OpPPC64FSUBS:
		return rewriteValuePPC64_OpPPC64FSUBS_0(v)
	case OpPPC64FTRUNC:
		return rewriteValuePPC64_OpPPC64FTRUNC_0(v)
	case OpPPC64GreaterEqual:
		return rewriteValuePPC64_OpPPC64GreaterEqual_0(v)
	case OpPPC64GreaterThan:
		return rewriteValuePPC64_OpPPC64GreaterThan_0(v)
	case OpPPC64LessEqual:
		return rewriteValuePPC64_OpPPC64LessEqual_0(v)
	case OpPPC64LessThan:
		return rewriteValuePPC64_OpPPC64LessThan_0(v)
	case OpPPC64MFVSRD:
		return rewriteValuePPC64_OpPPC64MFVSRD_0(v)
	case OpPPC64MOVBZload:
		return rewriteValuePPC64_OpPPC64MOVBZload_0(v)
	case OpPPC64MOVBZreg:
		return rewriteValuePPC64_OpPPC64MOVBZreg_0(v)
	case OpPPC64MOVBreg:
		return rewriteValuePPC64_OpPPC64MOVBreg_0(v)
	case OpPPC64MOVBstore:
		return rewriteValuePPC64_OpPPC64MOVBstore_0(v) || rewriteValuePPC64_OpPPC64MOVBstore_10(v) || rewriteValuePPC64_OpPPC64MOVBstore_20(v)
	case OpPPC64MOVBstorezero:
		return rewriteValuePPC64_OpPPC64MOVBstorezero_0(v)
	case OpPPC64MOVDload:
		return rewriteValuePPC64_OpPPC64MOVDload_0(v)
	case OpPPC64MOVDstore:
		return rewriteValuePPC64_OpPPC64MOVDstore_0(v)
	case OpPPC64MOVDstorezero:
		return rewriteValuePPC64_OpPPC64MOVDstorezero_0(v)
	case OpPPC64MOVHBRstore:
		return rewriteValuePPC64_OpPPC64MOVHBRstore_0(v)
	case OpPPC64MOVHZload:
		return rewriteValuePPC64_OpPPC64MOVHZload_0(v)
	case OpPPC64MOVHZreg:
		return rewriteValuePPC64_OpPPC64MOVHZreg_0(v)
	case OpPPC64MOVHload:
		return rewriteValuePPC64_OpPPC64MOVHload_0(v)
	case OpPPC64MOVHreg:
		return rewriteValuePPC64_OpPPC64MOVHreg_0(v)
	case OpPPC64MOVHstore:
		return rewriteValuePPC64_OpPPC64MOVHstore_0(v)
	case OpPPC64MOVHstorezero:
		return rewriteValuePPC64_OpPPC64MOVHstorezero_0(v)
	case OpPPC64MOVWBRstore:
		return rewriteValuePPC64_OpPPC64MOVWBRstore_0(v)
	case OpPPC64MOVWZload:
		return rewriteValuePPC64_OpPPC64MOVWZload_0(v)
	case OpPPC64MOVWZreg:
		return rewriteValuePPC64_OpPPC64MOVWZreg_0(v)
	case OpPPC64MOVWload:
		return rewriteValuePPC64_OpPPC64MOVWload_0(v)
	case OpPPC64MOVWreg:
		return rewriteValuePPC64_OpPPC64MOVWreg_0(v)
	case OpPPC64MOVWstore:
		return rewriteValuePPC64_OpPPC64MOVWstore_0(v)
	case OpPPC64MOVWstorezero:
		return rewriteValuePPC64_OpPPC64MOVWstorezero_0(v)
	case OpPPC64MTVSRD:
		return rewriteValuePPC64_OpPPC64MTVSRD_0(v)
	case OpPPC64MaskIfNotCarry:
		return rewriteValuePPC64_OpPPC64MaskIfNotCarry_0(v)
	case OpPPC64NotEqual:
		return rewriteValuePPC64_OpPPC64NotEqual_0(v)
	case OpPPC64OR:
		return rewriteValuePPC64_OpPPC64OR_0(v) || rewriteValuePPC64_OpPPC64OR_10(v) || rewriteValuePPC64_OpPPC64OR_20(v) || rewriteValuePPC64_OpPPC64OR_30(v) || rewriteValuePPC64_OpPPC64OR_40(v) || rewriteValuePPC64_OpPPC64OR_50(v) || rewriteValuePPC64_OpPPC64OR_60(v) || rewriteValuePPC64_OpPPC64OR_70(v) || rewriteValuePPC64_OpPPC64OR_80(v) || rewriteValuePPC64_OpPPC64OR_90(v) || rewriteValuePPC64_OpPPC64OR_100(v) || rewriteValuePPC64_OpPPC64OR_110(v)
	case OpPPC64ORN:
		return rewriteValuePPC64_OpPPC64ORN_0(v)
	case OpPPC64ORconst:
		return rewriteValuePPC64_OpPPC64ORconst_0(v)
	case OpPPC64SUB:
		return rewriteValuePPC64_OpPPC64SUB_0(v)
	case OpPPC64XOR:
		return rewriteValuePPC64_OpPPC64XOR_0(v) || rewriteValuePPC64_OpPPC64XOR_10(v)
	case OpPPC64XORconst:
		return rewriteValuePPC64_OpPPC64XORconst_0(v)
	case OpPopCount16:
		return rewriteValuePPC64_OpPopCount16_0(v)
	case OpPopCount32:
		return rewriteValuePPC64_OpPopCount32_0(v)
	case OpPopCount64:
		return rewriteValuePPC64_OpPopCount64_0(v)
	case OpPopCount8:
		return rewriteValuePPC64_OpPopCount8_0(v)
	case OpRound:
		return rewriteValuePPC64_OpRound_0(v)
	case OpRound32F:
		return rewriteValuePPC64_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValuePPC64_OpRound64F_0(v)
	case OpRsh16Ux16:
		return psess.rewriteValuePPC64_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return psess.rewriteValuePPC64_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return psess.rewriteValuePPC64_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return psess.rewriteValuePPC64_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return psess.rewriteValuePPC64_OpRsh16x16_0(v)
	case OpRsh16x32:
		return psess.rewriteValuePPC64_OpRsh16x32_0(v)
	case OpRsh16x64:
		return psess.rewriteValuePPC64_OpRsh16x64_0(v)
	case OpRsh16x8:
		return psess.rewriteValuePPC64_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return psess.rewriteValuePPC64_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return psess.rewriteValuePPC64_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return psess.rewriteValuePPC64_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return psess.rewriteValuePPC64_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return psess.rewriteValuePPC64_OpRsh32x16_0(v)
	case OpRsh32x32:
		return psess.rewriteValuePPC64_OpRsh32x32_0(v)
	case OpRsh32x64:
		return psess.rewriteValuePPC64_OpRsh32x64_0(v)
	case OpRsh32x8:
		return psess.rewriteValuePPC64_OpRsh32x8_0(v)
	case OpRsh64Ux16:
		return psess.rewriteValuePPC64_OpRsh64Ux16_0(v)
	case OpRsh64Ux32:
		return psess.rewriteValuePPC64_OpRsh64Ux32_0(v)
	case OpRsh64Ux64:
		return psess.rewriteValuePPC64_OpRsh64Ux64_0(v)
	case OpRsh64Ux8:
		return psess.rewriteValuePPC64_OpRsh64Ux8_0(v)
	case OpRsh64x16:
		return psess.rewriteValuePPC64_OpRsh64x16_0(v)
	case OpRsh64x32:
		return psess.rewriteValuePPC64_OpRsh64x32_0(v)
	case OpRsh64x64:
		return psess.rewriteValuePPC64_OpRsh64x64_0(v)
	case OpRsh64x8:
		return psess.rewriteValuePPC64_OpRsh64x8_0(v)
	case OpRsh8Ux16:
		return psess.rewriteValuePPC64_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return psess.rewriteValuePPC64_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return psess.rewriteValuePPC64_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return psess.rewriteValuePPC64_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return psess.rewriteValuePPC64_OpRsh8x16_0(v)
	case OpRsh8x32:
		return psess.rewriteValuePPC64_OpRsh8x32_0(v)
	case OpRsh8x64:
		return psess.rewriteValuePPC64_OpRsh8x64_0(v)
	case OpRsh8x8:
		return psess.rewriteValuePPC64_OpRsh8x8_0(v)
	case OpSignExt16to32:
		return rewriteValuePPC64_OpSignExt16to32_0(v)
	case OpSignExt16to64:
		return rewriteValuePPC64_OpSignExt16to64_0(v)
	case OpSignExt32to64:
		return rewriteValuePPC64_OpSignExt32to64_0(v)
	case OpSignExt8to16:
		return rewriteValuePPC64_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValuePPC64_OpSignExt8to32_0(v)
	case OpSignExt8to64:
		return rewriteValuePPC64_OpSignExt8to64_0(v)
	case OpSlicemask:
		return rewriteValuePPC64_OpSlicemask_0(v)
	case OpSqrt:
		return rewriteValuePPC64_OpSqrt_0(v)
	case OpStaticCall:
		return rewriteValuePPC64_OpStaticCall_0(v)
	case OpStore:
		return psess.rewriteValuePPC64_OpStore_0(v)
	case OpSub16:
		return rewriteValuePPC64_OpSub16_0(v)
	case OpSub32:
		return rewriteValuePPC64_OpSub32_0(v)
	case OpSub32F:
		return rewriteValuePPC64_OpSub32F_0(v)
	case OpSub64:
		return rewriteValuePPC64_OpSub64_0(v)
	case OpSub64F:
		return rewriteValuePPC64_OpSub64F_0(v)
	case OpSub8:
		return rewriteValuePPC64_OpSub8_0(v)
	case OpSubPtr:
		return rewriteValuePPC64_OpSubPtr_0(v)
	case OpTrunc:
		return rewriteValuePPC64_OpTrunc_0(v)
	case OpTrunc16to8:
		return rewriteValuePPC64_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValuePPC64_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValuePPC64_OpTrunc32to8_0(v)
	case OpTrunc64to16:
		return rewriteValuePPC64_OpTrunc64to16_0(v)
	case OpTrunc64to32:
		return rewriteValuePPC64_OpTrunc64to32_0(v)
	case OpTrunc64to8:
		return rewriteValuePPC64_OpTrunc64to8_0(v)
	case OpWB:
		return rewriteValuePPC64_OpWB_0(v)
	case OpXor16:
		return rewriteValuePPC64_OpXor16_0(v)
	case OpXor32:
		return rewriteValuePPC64_OpXor32_0(v)
	case OpXor64:
		return rewriteValuePPC64_OpXor64_0(v)
	case OpXor8:
		return rewriteValuePPC64_OpXor8_0(v)
	case OpZero:
		return psess.rewriteValuePPC64_OpZero_0(v) || psess.rewriteValuePPC64_OpZero_10(v)
	case OpZeroExt16to32:
		return rewriteValuePPC64_OpZeroExt16to32_0(v)
	case OpZeroExt16to64:
		return rewriteValuePPC64_OpZeroExt16to64_0(v)
	case OpZeroExt32to64:
		return rewriteValuePPC64_OpZeroExt32to64_0(v)
	case OpZeroExt8to16:
		return rewriteValuePPC64_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValuePPC64_OpZeroExt8to32_0(v)
	case OpZeroExt8to64:
		return rewriteValuePPC64_OpZeroExt8to64_0(v)
	}
	return false
}
func rewriteValuePPC64_OpAbs_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64FABS)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpAdd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAdd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAdd32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FADDS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAdd64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAdd64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAdd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAddPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAddr_0(v *Value) bool {

	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpPPC64MOVDaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValuePPC64_OpAnd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAnd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAnd64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAnd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAndB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpAtomicAdd32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicAdd32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicAdd64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicAdd64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicAnd8_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicAnd8)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicCompareAndSwap32_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpPPC64LoweredAtomicCas32)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicCompareAndSwap64_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpPPC64LoweredAtomicCas64)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicExchange32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicExchange32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicExchange64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicExchange64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicLoad32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64LoweredAtomicLoad32)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicLoad64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64LoweredAtomicLoad64)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicLoadPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64LoweredAtomicLoadPtr)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicOr8_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicOr8)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicStore32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicStore32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAtomicStore64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredAtomicStore64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpAvg64u_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ADD)
		v0 := b.NewValue0(v.Pos, OpPPC64SRDconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpPPC64SUB, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpBitLen32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64SUB)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 32
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64CNTLZW, typ.Int)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpBitLen64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64SUB)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = 64
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64CNTLZD, typ.Int)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpCeil_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64FCEIL)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpClosureCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64CALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpCom16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCom32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCom64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCom8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64NOR)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpConst16_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConst32_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConst32F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpPPC64FMOVSconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConst64_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConst64F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConst8_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValuePPC64_OpConstBool_0(v *Value) bool {

	for {
		b := v.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValuePPC64_OpConstNil_0(v *Value) bool {

	for {
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValuePPC64_OpCopysign_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FCPSGN)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCtz32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64POPCNTW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZreg, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpPPC64ANDN, typ.Int)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconst, typ.Int)
		v2.AuxInt = -1
		v2.AddArg(x)
		v1.AddArg(v2)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCtz32NonZero_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCtz32)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCtz64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64POPCNTD)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDN, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpPPC64ADDconst, typ.Int64)
		v1.AuxInt = -1
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCtz64NonZero_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCtz64)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCvt32Fto32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64MFVSRD)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIWZ, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt32Fto64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64MFVSRD)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIDZ, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt32Fto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCvt32to32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64FCFIDS)
		v0 := b.NewValue0(v.Pos, OpPPC64MTVSRD, typ.Float64)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt32to64F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64FCFID)
		v0 := b.NewValue0(v.Pos, OpPPC64MTVSRD, typ.Float64)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64Fto32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64MFVSRD)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIWZ, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64Fto32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64FRSP)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpCvt64Fto64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64MFVSRD)
		v0 := b.NewValue0(v.Pos, OpPPC64FCTIDZ, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64to32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64FCFIDS)
		v0 := b.NewValue0(v.Pos, OpPPC64MTVSRD, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpCvt64to64F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64FCFID)
		v0 := b.NewValue0(v.Pos, OpPPC64MTVSRD, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpDiv16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpDiv16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVWU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpDiv32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FDIVS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv32u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FDIV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv64u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVDU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpDiv8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpDiv8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64DIVWU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpEq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(isSigned(x.Type) && isSigned(y.Type)) {
			break
		}
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpEq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpEq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpEq64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpEq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpEq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(isSigned(x.Type) && isSigned(y.Type)) {
			break
		}
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpEqB_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpPPC64EQV, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpEqPtr_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64Equal)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpFloor_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64FFLOOR)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGeq32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGeq64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGeq64U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpGetCallerPC_0(v *Value) bool {

	for {
		v.reset(OpPPC64LoweredGetCallerPC)
		return true
	}
}
func rewriteValuePPC64_OpGetCallerSP_0(v *Value) bool {

	for {
		v.reset(OpPPC64LoweredGetCallerSP)
		return true
	}
}
func rewriteValuePPC64_OpGetClosurePtr_0(v *Value) bool {

	for {
		v.reset(OpPPC64LoweredGetClosurePtr)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGreater16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGreater16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGreater32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGreater32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FGreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGreater32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGreater64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGreater64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FGreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGreater64U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGreater8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpGreater8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpHmul32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULHW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpHmul32u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULHWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpHmul64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULHD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpHmul64u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULHDU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpInterCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[1]
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64CALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpIsInBounds_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, psess.types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		ptr := v.Args[0]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPconst, psess.types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(ptr)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpIsSliceInBounds_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, psess.types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FLessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLeq32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLeq64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FLessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLeq64U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLess16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLess16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLess32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLess32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FLessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLess32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLess64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLess64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FLessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLess64U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLess8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLess8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64LessThan)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWU, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLoad_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpPPC64MOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is32BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is32BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVWZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVHload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVHZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVBreg)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, typ.UInt8)
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is32BitFloat(t)) {
			break
		}
		v.reset(OpPPC64FMOVSload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is64BitFloat(t)) {
			break
		}
		v.reset(OpPPC64FMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -16
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -16
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
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
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -16
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -16
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
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
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1.AuxInt != 31 {
			break
		}
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int32)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 31 {
			break
		}
		y := v_1.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int32)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1.Type != typ.Int32 {
			break
		}
		if v_1.AuxInt != 31 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int32)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
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
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1.AuxInt != 63 {
			break
		}
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int64)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 63 {
			break
		}
		y := v_1.Args[1]
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int64)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1.Type != typ.Int64 {
			break
		}
		if v_1.AuxInt != 63 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int64)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -8
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -8
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
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
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -8
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -8
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpMod16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpMod32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLW, typ.Int32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVW, typ.Int32)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLW, typ.Int32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVWU, typ.Int32)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLD, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVD, typ.Int64)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod64u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64MULLD, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64DIVDU, typ.Int64)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpMod8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpMod8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpMove_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		if v.AuxInt != 0 {
			break
		}
		_ = v.Args[2]
		mem := v.Args[2]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, typ.UInt8)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, typ.UInt16)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, typ.Int64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 3 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVHload, typ.Int16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 5 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, typ.UInt8)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 6 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, typ.UInt16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 7 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBZload, typ.UInt8)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHstore, psess.types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVHZload, typ.UInt16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpPPC64MOVWstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpPPC64MOVWZload, typ.UInt32)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuePPC64_OpMove_10(v *Value) bool {

	for {
		s := v.AuxInt
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 8) {
			break
		}
		v.reset(OpPPC64LoweredMove)
		v.AuxInt = s
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpMul16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpMul32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpMul32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FMULS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpMul64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpMul64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpMul8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64MULLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpNeg16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpNeg32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpNeg32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64FNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpNeg64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpNeg64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64FNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpNeg8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64NEG)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpNeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(isSigned(x.Type) && isSigned(y.Type)) {
			break
		}
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpNeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpNeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpNeq64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpNeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64FCMPU, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpNeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(isSigned(x.Type) && isSigned(y.Type)) {
			break
		}
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPW, psess.types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNeqB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpNeqPtr_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpPPC64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpNilCheck_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64LoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpNot_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64XORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpOffPtr_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpPPC64ADD)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v0.AuxInt = off
		v.AddArg(v0)
		v.AddArg(ptr)
		return true
	}
}
func rewriteValuePPC64_OpOr16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpOr32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpOr64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpOr8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpOrB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpPPC64ADD_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRDconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpPPC64ROTLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLDconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpPPC64ROTLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpPPC64ROTLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLWconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpPPC64ROTLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1.Type != typ.Int64 {
			break
		}
		if v_0_1.AuxInt != 63 {
			break
		}
		y := v_0_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRD {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64SUB {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		if v_1_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_0.AuxInt != 64 {
			break
		}
		v_1_1_1 := v_1_1.Args[1]
		if v_1_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1_1.Type != typ.UInt {
			break
		}
		if v_1_1_1.AuxInt != 63 {
			break
		}
		if y != v_1_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64SUB {
			break
		}
		if v_0_1.Type != typ.UInt {
			break
		}
		_ = v_0_1.Args[1]
		v_0_1_0 := v_0_1.Args[0]
		if v_0_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_0_1_0.AuxInt != 64 {
			break
		}
		v_0_1_1 := v_0_1.Args[1]
		if v_0_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1_1.Type != typ.UInt {
			break
		}
		if v_0_1_1.AuxInt != 63 {
			break
		}
		y := v_0_1_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLD {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1.Type != typ.Int64 {
			break
		}
		if v_1_1.AuxInt != 63 {
			break
		}
		if y != v_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLW {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1.Type != typ.Int32 {
			break
		}
		if v_0_1.AuxInt != 31 {
			break
		}
		y := v_0_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRW {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64SUB {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		if v_1_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_0.AuxInt != 32 {
			break
		}
		v_1_1_1 := v_1_1.Args[1]
		if v_1_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1_1.Type != typ.UInt {
			break
		}
		if v_1_1_1.AuxInt != 31 {
			break
		}
		if y != v_1_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRW {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64SUB {
			break
		}
		if v_0_1.Type != typ.UInt {
			break
		}
		_ = v_0_1.Args[1]
		v_0_1_0 := v_0_1.Args[0]
		if v_0_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_0_1_0.AuxInt != 32 {
			break
		}
		v_0_1_1 := v_0_1.Args[1]
		if v_0_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1_1.Type != typ.UInt {
			break
		}
		if v_0_1_1.AuxInt != 31 {
			break
		}
		y := v_0_1_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLW {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1.Type != typ.Int32 {
			break
		}
		if v_1_1.AuxInt != 31 {
			break
		}
		if y != v_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpPPC64ADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpPPC64ADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64ADDconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpPPC64ADDconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDaddr {
			break
		}
		d := v_0.AuxInt
		sym := v_0.Aux
		x := v_0.Args[0]
		v.reset(OpPPC64MOVDaddr)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64AND_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64NOR {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if y != v_1.Args[1] {
			break
		}
		v.reset(OpPPC64ANDN)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64NOR {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		if y != v_0.Args[1] {
			break
		}
		x := v.Args[1]
		v.reset(OpPPC64ANDN)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c & d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c & d
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if x.Op != OpPPC64MOVBZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpPPC64MOVBZload {
			break
		}
		_ = x.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpPPC64MOVBZload {
			break
		}
		_ = x.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if x.Op != OpPPC64MOVBZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64ANDconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		y := v.Args[0]
		if y.Op != OpPPC64MOVBZreg {
			break
		}
		if !(c&0xFF == 0xFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		y := v.Args[0]
		if y.Op != OpPPC64MOVHZreg {
			break
		}
		if !(c&0xFFFF == 0xFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		y := v.Args[0]
		if y.Op != OpPPC64MOVWZreg {
			break
		}
		if !(c&0xFFFFFFFF == 0xFFFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFF
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFFFF
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64ANDconst)
		v.AuxInt = c & 0xFFFFFFFF
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuePPC64_OpPPC64CMP_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v.Args[1]
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuePPC64_OpPPC64CMPU_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v.Args[1]
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPUconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPUconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x == y) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) < uint64(y)) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) > uint64(y)) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuePPC64_OpPPC64CMPW_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64CMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVWreg {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64CMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v.Args[1]
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuePPC64_OpPPC64CMPWU_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64CMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64CMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64CMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		y := v.Args[1]
		if !(isU16Bit(c)) {
			break
		}
		v.reset(OpPPC64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpPPC64CMPWUconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPWUconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPWconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y)) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y)) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64CMPconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x == y) {
			break
		}
		v.reset(OpPPC64FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x < y) {
			break
		}
		v.reset(OpPPC64FlagLT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x > y) {
			break
		}
		v.reset(OpPPC64FlagGT)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64Equal_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64Equal)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FABS_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		x := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = f2i(math.Abs(i2f(x)))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FADD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMUL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		z := v.Args[1]
		v.reset(OpPPC64FMADD)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		z := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64FMUL {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpPPC64FMADD)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FADDS_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMULS {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		z := v.Args[1]
		v.reset(OpPPC64FMADDS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		z := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64FMULS {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpPPC64FMADDS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FCEIL_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		x := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = f2i(math.Ceil(i2f(x)))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FFLOOR_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		x := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = f2i(math.Floor(i2f(x)))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FMOVDload_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		if ptr != v_1.Args[0] {
			break
		}
		x := v_1.Args[1]
		v.reset(OpPPC64MTVSRD)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64FMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FMOVDstore_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MTVSRD {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64FMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FMOVSload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64FMOVSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FMOVSstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64FMOVSstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64FMOVSstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FNEG_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FABS {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64FNABS)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FNABS {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64FABS)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FSQRT_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		x := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = f2i(math.Sqrt(i2f(x)))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FSUB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMUL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		z := v.Args[1]
		v.reset(OpPPC64FMSUB)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FSUBS_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMULS {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		z := v.Args[1]
		v.reset(OpPPC64FMSUBS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64FTRUNC_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		x := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = f2i(math.Trunc(i2f(x)))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64GreaterEqual_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64LessEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64GreaterThan_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64LessThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64LessEqual_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64GreaterEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64LessThan_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64GreaterThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MFVSRD_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpPPC64FMOVDload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[1]
		ptr := x.Args[0]
		mem := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, typ.Int64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBZload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVBZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBZreg_0(v *Value) bool {

	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVBreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpPPC64MOVBZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBreg_0(v *Value) bool {

	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0x7F) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBstore_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(c == 0) {
			break
		}
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVBZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		c := v_1.AuxInt
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVHreg {
			break
		}
		x := v_1_0.Args[0]
		mem := v.Args[2]
		if !(c <= 8) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBstore_10(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		c := v_1.AuxInt
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_1_0.Args[0]
		mem := v.Args[2]
		if !(c <= 8) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		c := v_1.AuxInt
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1_0.Args[0]
		mem := v.Args[2]
		if !(c <= 24) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		c := v_1.AuxInt
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1_0.Args[0]
		mem := v.Args[2]
		if !(c <= 24) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt32)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		i1 := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		if v_1.AuxInt != 24 {
			break
		}
		w := v_1.Args[0]
		x0 := v.Args[2]
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRWconst {
			break
		}
		if x0_1.AuxInt != 16 {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		mem := x0.Args[2]
		if !(!config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt16)
		v0.AuxInt = 16
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		i1 := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRDconst {
			break
		}
		if v_1.AuxInt != 24 {
			break
		}
		w := v_1.Args[0]
		x0 := v.Args[2]
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRDconst {
			break
		}
		if x0_1.AuxInt != 16 {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		mem := x0.Args[2]
		if !(!config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpPPC64SRWconst, typ.UInt16)
		v0.AuxInt = 16
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		i1 := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		if v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x0 := v.Args[2]
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		if w != x0.Args[1] {
			break
		}
		mem := x0.Args[2]
		if !(!config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i1 := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRDconst {
			break
		}
		if v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x0 := v.Args[2]
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		if w != x0.Args[1] {
			break
		}
		mem := x0.Args[2]
		if !(!config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i3 := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i2 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRWconst {
			break
		}
		if x0_1.AuxInt != 8 {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpPPC64MOVBstore {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if p != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpPPC64SRWconst {
			break
		}
		if x1_1.AuxInt != 16 {
			break
		}
		if w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpPPC64SRWconst {
			break
		}
		if x2_1.AuxInt != 24 {
			break
		}
		if w != x2_1.Args[0] {
			break
		}
		mem := x2.Args[2]
		if !(!config.BigEndian && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpPPC64MOVWBRstore)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i1 := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRWconst {
			break
		}
		if x0_1.AuxInt != 8 {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		mem := x0.Args[2]
		if !(!config.BigEndian && x0.Uses == 1 && i1 == i0+1 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVHBRstore)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i7 := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRDconst {
			break
		}
		if v_1.AuxInt != 56 {
			break
		}
		w := v_1.Args[0]
		x0 := v.Args[2]
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i6 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRDconst {
			break
		}
		if x0_1.AuxInt != 48 {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpPPC64MOVBstore {
			break
		}
		i5 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if p != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpPPC64SRDconst {
			break
		}
		if x1_1.AuxInt != 40 {
			break
		}
		if w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpPPC64MOVBstore {
			break
		}
		i4 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpPPC64SRDconst {
			break
		}
		if x2_1.AuxInt != 32 {
			break
		}
		if w != x2_1.Args[0] {
			break
		}
		x3 := x2.Args[2]
		if x3.Op != OpPPC64MOVWstore {
			break
		}
		i0 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[2]
		if p != x3.Args[0] {
			break
		}
		if w != x3.Args[1] {
			break
		}
		mem := x3.Args[2]
		if !(!config.BigEndian && i0%4 == 0 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3)) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBstore_20(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		i7 := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpPPC64MOVBstore {
			break
		}
		i6 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpPPC64SRDconst {
			break
		}
		if x0_1.AuxInt != 8 {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpPPC64MOVBstore {
			break
		}
		i5 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if p != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpPPC64SRDconst {
			break
		}
		if x1_1.AuxInt != 16 {
			break
		}
		if w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpPPC64MOVBstore {
			break
		}
		i4 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpPPC64SRDconst {
			break
		}
		if x2_1.AuxInt != 24 {
			break
		}
		if w != x2_1.Args[0] {
			break
		}
		x3 := x2.Args[2]
		if x3.Op != OpPPC64MOVBstore {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[2]
		if p != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpPPC64SRDconst {
			break
		}
		if x3_1.AuxInt != 32 {
			break
		}
		if w != x3_1.Args[0] {
			break
		}
		x4 := x3.Args[2]
		if x4.Op != OpPPC64MOVBstore {
			break
		}
		i2 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[2]
		if p != x4.Args[0] {
			break
		}
		x4_1 := x4.Args[1]
		if x4_1.Op != OpPPC64SRDconst {
			break
		}
		if x4_1.AuxInt != 40 {
			break
		}
		if w != x4_1.Args[0] {
			break
		}
		x5 := x4.Args[2]
		if x5.Op != OpPPC64MOVBstore {
			break
		}
		i1 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[2]
		if p != x5.Args[0] {
			break
		}
		x5_1 := x5.Args[1]
		if x5_1.Op != OpPPC64SRDconst {
			break
		}
		if x5_1.AuxInt != 48 {
			break
		}
		if w != x5_1.Args[0] {
			break
		}
		x6 := x5.Args[2]
		if x6.Op != OpPPC64MOVBstore {
			break
		}
		i0 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[2]
		if p != x6.Args[0] {
			break
		}
		x6_1 := x6.Args[1]
		if x6_1.Op != OpPPC64SRDconst {
			break
		}
		if x6_1.AuxInt != 56 {
			break
		}
		if w != x6_1.Args[0] {
			break
		}
		mem := x6.Args[2]
		if !(!config.BigEndian && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6)) {
			break
		}
		v.reset(OpPPC64MOVDBRstore)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVBstorezero_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		x := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (x.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVDload_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64FMOVDstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		if ptr != v_1.Args[0] {
			break
		}
		x := v_1.Args[1]
		v.reset(OpPPC64MFVSRD)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVDstore_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MFVSRD {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64FMOVDstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(c == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVDstorezero_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		x := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (x.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHBRstore_0(v *Value) bool {

	for {
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHZload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVHZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHZreg_0(v *Value) bool {

	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHBRload {
			break
		}
		_ = y.Args[1]
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVHZreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpPPC64MOVHZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHreg_0(v *Value) bool {

	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0x7FFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHZreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpPPC64MOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(c == 0) {
			break
		}
		v.reset(OpPPC64MOVHstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVHZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		i1 := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x0 := v.Args[2]
		if x0.Op != OpPPC64MOVHstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		if w != x0.Args[1] {
			break
		}
		mem := x0.Args[2]
		if !(!config.BigEndian && x0.Uses == 1 && i1 == i0+2 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i1 := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRDconst {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x0 := v.Args[2]
		if x0.Op != OpPPC64MOVHstore {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		if w != x0.Args[1] {
			break
		}
		mem := x0.Args[2]
		if !(!config.BigEndian && x0.Uses == 1 && i1 == i0+2 && clobber(x0)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = i0
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVHstorezero_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		x := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (x.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWBRstore_0(v *Value) bool {

	for {
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVWBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVWBRstore)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWZload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVWZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVWZload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWZreg_0(v *Value) bool {

	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFFFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64AND {
			break
		}
		_ = y.Args[1]
		y_0 := y.Args[0]
		if y_0.Op != OpPPC64MOVDconst {
			break
		}
		c := y_0.AuxInt
		if !(uint64(c) <= 0xFFFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64AND {
			break
		}
		_ = y.Args[1]
		y_1 := y.Args[1]
		if y_1.Op != OpPPC64MOVDconst {
			break
		}
		c := y_1.AuxInt
		if !(uint64(c) <= 0xFFFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVWZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBZreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHBRload {
			break
		}
		_ = y.Args[1]
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVWBRload {
			break
		}
		_ = y.Args[1]
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVWreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVWZreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWreg_0(v *Value) bool {

	for {
		y := v.Args[0]
		if y.Op != OpPPC64ANDconst {
			break
		}
		c := y.AuxInt
		if !(uint64(c) <= 0xFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64AND {
			break
		}
		_ = y.Args[1]
		y_0 := y.Args[0]
		if y_0.Op != OpPPC64MOVDconst {
			break
		}
		c := y_0.AuxInt
		if !(uint64(c) <= 0x7FFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64AND {
			break
		}
		_ = y.Args[1]
		y_1 := y.Args[1]
		if y_1.Op != OpPPC64MOVDconst {
			break
		}
		c := y_1.AuxInt
		if !(uint64(c) <= 0x7FFFFFFF) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVWreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVHreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVBreg {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		y := v.Args[0]
		if y.Op != OpPPC64MOVWZreg {
			break
		}
		x := y.Args[0]
		v.reset(OpPPC64MOVWreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		ptr := p.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && (ptr.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(c == 0) {
			break
		}
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpPPC64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MOVWstorezero_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		x := v_0.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1 + off2)) {
			break
		}
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		if p.Op != OpPPC64MOVDaddr {
			break
		}
		off2 := p.AuxInt
		sym2 := p.Aux
		x := p.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && (x.Op != OpSB || p.Uses == 1)) {
			break
		}
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MTVSRD_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpPPC64FMOVDconst)
		v.AuxInt = c
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpPPC64MOVDload {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[1]
		ptr := x.Args[0]
		mem := x.Args[1]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpPPC64FMOVDload, typ.Float64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64MaskIfNotCarry_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ADDconstForCarry {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpPPC64ANDconst {
			break
		}
		d := v_0_0.AuxInt
		if !(c < 0 && d > 0 && c+d < 0) {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = -1
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64NotEqual_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagEQ {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagLT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64FlagGT {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpPPC64NotEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRDconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpPPC64ROTLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLDconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpPPC64ROTLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpPPC64ROTLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLWconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpPPC64ROTLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1.Type != typ.Int64 {
			break
		}
		if v_0_1.AuxInt != 63 {
			break
		}
		y := v_0_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRD {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64SUB {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		if v_1_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_0.AuxInt != 64 {
			break
		}
		v_1_1_1 := v_1_1.Args[1]
		if v_1_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1_1.Type != typ.UInt {
			break
		}
		if v_1_1_1.AuxInt != 63 {
			break
		}
		if y != v_1_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64SUB {
			break
		}
		if v_0_1.Type != typ.UInt {
			break
		}
		_ = v_0_1.Args[1]
		v_0_1_0 := v_0_1.Args[0]
		if v_0_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_0_1_0.AuxInt != 64 {
			break
		}
		v_0_1_1 := v_0_1.Args[1]
		if v_0_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1_1.Type != typ.UInt {
			break
		}
		if v_0_1_1.AuxInt != 63 {
			break
		}
		y := v_0_1_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLD {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1.Type != typ.Int64 {
			break
		}
		if v_1_1.AuxInt != 63 {
			break
		}
		if y != v_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLW {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1.Type != typ.Int32 {
			break
		}
		if v_0_1.AuxInt != 31 {
			break
		}
		y := v_0_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRW {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64SUB {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		if v_1_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_0.AuxInt != 32 {
			break
		}
		v_1_1_1 := v_1_1.Args[1]
		if v_1_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1_1.Type != typ.UInt {
			break
		}
		if v_1_1_1.AuxInt != 31 {
			break
		}
		if y != v_1_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRW {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64SUB {
			break
		}
		if v_0_1.Type != typ.UInt {
			break
		}
		_ = v_0_1.Args[1]
		v_0_1_0 := v_0_1.Args[0]
		if v_0_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_0_1_0.AuxInt != 32 {
			break
		}
		v_0_1_1 := v_0_1.Args[1]
		if v_0_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1_1.Type != typ.UInt {
			break
		}
		if v_0_1_1.AuxInt != 31 {
			break
		}
		y := v_0_1_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLW {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1.Type != typ.Int32 {
			break
		}
		if v_1_1.AuxInt != 31 {
			break
		}
		if y != v_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c | d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c | d
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_10(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpPPC64ORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpPPC64ORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o1 := v.Args[1]
		if o1.Op != OpPPC64SLWconst {
			break
		}
		if o1.AuxInt != 8 {
			break
		}
		x1 := o1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o1 := v.Args[0]
		if o1.Op != OpPPC64SLWconst {
			break
		}
		if o1.AuxInt != 8 {
			break
		}
		x1 := o1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := v.Args[1]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o1 := v.Args[1]
		if o1.Op != OpPPC64SLDconst {
			break
		}
		if o1.AuxInt != 8 {
			break
		}
		x1 := o1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o1 := v.Args[0]
		if o1.Op != OpPPC64SLDconst {
			break
		}
		if o1.AuxInt != 8 {
			break
		}
		x1 := o1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := v.Args[1]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o1 := v.Args[1]
		if o1.Op != OpPPC64SLWconst {
			break
		}
		if o1.AuxInt != 8 {
			break
		}
		x1 := o1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o1 := v.Args[0]
		if o1.Op != OpPPC64SLWconst {
			break
		}
		if o1.AuxInt != 8 {
			break
		}
		x1 := o1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := v.Args[1]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o1 := v.Args[1]
		if o1.Op != OpPPC64SLDconst {
			break
		}
		if o1.AuxInt != 8 {
			break
		}
		x1 := o1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o1 := v.Args[0]
		if o1.Op != OpPPC64SLDconst {
			break
		}
		if o1.AuxInt != 8 {
			break
		}
		x1 := o1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := v.Args[1]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && o1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(o1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_20(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		n1 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		n2 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && n1%8 == 0 && n2 == n1+8 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = n1
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		n2 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		n1 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && n1%8 == 0 && n2 == n1+8 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = n1
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		n1 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		n2 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && n1%8 == 0 && n2 == n1+8 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = n1
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		n2 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		n1 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && n1%8 == 0 && n2 == n1+8 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = n1
		v1 := b.NewValue0(v.Pos, OpPPC64MOVHBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		x0 := o0.Args[1]
		if x0.Op != OpPPC64MOVHZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpPPC64MOVHZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := o0.Args[1]
		if x0.Op != OpPPC64MOVHZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s1 := v.Args[1]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpPPC64MOVHZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s1 := v.Args[1]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		x0 := o0.Args[1]
		if x0.Op != OpPPC64MOVHZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpPPC64MOVHZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_30(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := o0.Args[1]
		if x0.Op != OpPPC64MOVHZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s1 := v.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpPPC64MOVHZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s1 := v.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		x0 := o0.Args[1]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := o0.Args[1]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s1 := v.Args[1]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		s := x0_0.Aux
		p := x0_0.Args[0]
		mem := x0.Args[1]
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s1 := v.Args[1]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		x0 := o0.Args[1]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := o0.Args[1]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s1 := v.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		s := x0_0.Aux
		p := x0_0.Args[0]
		mem := x0.Args[1]
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s1 := v.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_40(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 16 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVHBRload {
			break
		}
		if x2.Type != t {
			break
		}
		_ = x2.Args[1]
		x2_0 := x2.Args[0]
		if x2_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x2_0.Type != typ.Uintptr {
			break
		}
		i0 := x2_0.AuxInt
		if x2_0.Aux != s {
			break
		}
		if p != x2_0.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 16 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVHBRload {
			break
		}
		if x2.Type != t {
			break
		}
		_ = x2.Args[1]
		x2_0 := x2.Args[0]
		if x2_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x2_0.Type != typ.Uintptr {
			break
		}
		i0 := x2_0.AuxInt
		if x2_0.Aux != s {
			break
		}
		if p != x2_0.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 16 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVHBRload {
			break
		}
		if x2.Type != t {
			break
		}
		_ = x2.Args[1]
		x2_0 := x2.Args[0]
		if x2_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x2_0.Type != typ.Uintptr {
			break
		}
		i0 := x2_0.AuxInt
		if x2_0.Aux != s {
			break
		}
		if p != x2_0.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		x0 := v.Args[1]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLWconst {
			break
		}
		if s1.AuxInt != 16 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVHBRload {
			break
		}
		if x2.Type != t {
			break
		}
		_ = x2.Args[1]
		x2_0 := x2.Args[0]
		if x2_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x2_0.Type != typ.Uintptr {
			break
		}
		i0 := x2_0.AuxInt
		s := x2_0.Aux
		p := x2_0.Args[0]
		mem := x2.Args[1]
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		x0 := v.Args[1]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 16 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVHBRload {
			break
		}
		if x2.Type != t {
			break
		}
		_ = x2.Args[1]
		x2_0 := x2.Args[0]
		if x2_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x2_0.Type != typ.Uintptr {
			break
		}
		i0 := x2_0.AuxInt
		if x2_0.Aux != s {
			break
		}
		if p != x2_0.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 16 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVHBRload {
			break
		}
		if x2.Type != t {
			break
		}
		_ = x2.Args[1]
		x2_0 := x2.Args[0]
		if x2_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x2_0.Type != typ.Uintptr {
			break
		}
		i0 := x2_0.AuxInt
		if x2_0.Aux != s {
			break
		}
		if p != x2_0.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 16 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVHBRload {
			break
		}
		if x2.Type != t {
			break
		}
		_ = x2.Args[1]
		x2_0 := x2.Args[0]
		if x2_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x2_0.Type != typ.Uintptr {
			break
		}
		i0 := x2_0.AuxInt
		if x2_0.Aux != s {
			break
		}
		if p != x2_0.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		x0 := v.Args[1]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 16 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != OpPPC64MOVHBRload {
			break
		}
		if x2.Type != t {
			break
		}
		_ = x2.Args[1]
		x2_0 := x2.Args[0]
		if x2_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x2_0.Type != typ.Uintptr {
			break
		}
		i0 := x2_0.AuxInt
		s := x2_0.Aux
		p := x2_0.Args[0]
		mem := x2.Args[1]
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		x0 := v.Args[1]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s2 := v.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 32 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 40 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 48 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i0 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s2 := v.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 32 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 48 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i0 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 40 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_50(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 40 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 48 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i0 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s2 := v.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 32 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 48 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i0 := x0_0.AuxInt
		s := x0_0.Aux
		p := x0_0.Args[0]
		mem := x0.Args[1]
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 40 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s2 := v.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 32 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s2 := v.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 56 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s2 := v.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 56 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		s0 := o0.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		if x0_0.Aux != s {
			break
		}
		if p != x0_0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s2 := v.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 56 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVHBRload {
			break
		}
		if x0.Type != t {
			break
		}
		_ = x0.Args[1]
		x0_0 := x0.Args[0]
		if x0_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x0_0.Type != typ.Uintptr {
			break
		}
		i2 := x0_0.AuxInt
		s := x0_0.Aux
		p := x0_0.Args[0]
		mem := x0.Args[1]
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s2 := v.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 56 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && o0.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpPPC64SLDconst, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = 32
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWBRload, t)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s6 := v.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s5 := o5.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s4 := o4.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s3 := o3.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		x0 := o3.Args[1]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s6 := v.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s5 := o5.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s4 := o4.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		x0 := o3.Args[0]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s3 := o3.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s6 := v.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s5 := o5.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s3 := o3.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		x0 := o3.Args[1]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s4 := o4.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s6 := v.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s5 := o5.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		x0 := o3.Args[0]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s3 := o3.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s4 := o4.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_60(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		t := v.Type
		_ = v.Args[1]
		s6 := v.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s4 := o4.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s3 := o3.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		x0 := o3.Args[1]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s5 := o5.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s6 := v.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s4 := o4.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		x0 := o3.Args[0]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s3 := o3.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s5 := o5.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s6 := v.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s3 := o3.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		x0 := o3.Args[1]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s4 := o4.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s5 := o5.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s6 := v.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		x0 := o3.Args[0]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s3 := o3.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s4 := o4.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s5 := o5.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s5 := o5.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s4 := o4.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s3 := o3.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		x0 := o3.Args[1]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s6 := v.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s5 := o5.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s4 := o4.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		x0 := o3.Args[0]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s3 := o3.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s6 := v.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s5 := o5.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s3 := o3.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		x0 := o3.Args[1]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s4 := o4.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s6 := v.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s5 := o5.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		x0 := o3.Args[0]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s3 := o3.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s4 := o4.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s6 := v.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s4 := o4.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		s := x5.Aux
		_ = x5.Args[1]
		p := x5.Args[0]
		mem := x5.Args[1]
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s3 := o3.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		x0 := o3.Args[1]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s5 := o5.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		s6 := v.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s4 := o4.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		s := x5.Aux
		_ = x5.Args[1]
		p := x5.Args[0]
		mem := x5.Args[1]
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		x0 := o3.Args[0]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s3 := o3.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s5 := o5.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		s6 := v.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_70(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s3 := o3.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		s := x4.Aux
		_ = x4.Args[1]
		p := x4.Args[0]
		mem := x4.Args[1]
		x0 := o3.Args[1]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		s4 := o4.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s5 := o5.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		s6 := v.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		x0 := o3.Args[0]
		if x0.Op != OpPPC64MOVWZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		s3 := o3.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x4 := s3.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s4 := o4.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 40 {
			break
		}
		x5 := s4.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s5 := o5.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 48 {
			break
		}
		x6 := s5.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		s6 := v.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 56 {
			break
		}
		x7 := s6.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i0%4 == 0 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s3.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x0, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(s3) && clobber(s4) && clobber(s5) && clobber(s6) && clobber(o3) && clobber(o4) && clobber(o5)) {
			break
		}
		b = mergePoint(b, x0, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		o1 := o0.Args[1]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		s2 := o1.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		o2 := o1.Args[1]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		s3 := o2.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		x4 := o2.Args[1]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		o1 := o0.Args[1]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		s2 := o1.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		o2 := o1.Args[1]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		x4 := o2.Args[0]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s3 := o2.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		o1 := o0.Args[1]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		s3 := o2.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		x4 := o2.Args[1]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s2 := o1.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		o1 := o0.Args[1]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		x4 := o2.Args[0]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s3 := o2.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s2 := o1.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		s2 := o1.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		o2 := o1.Args[1]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		s3 := o2.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		x4 := o2.Args[1]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		s2 := o1.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		o2 := o1.Args[1]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		x4 := o2.Args[0]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s3 := o2.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		s3 := o2.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		x4 := o2.Args[1]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s2 := o1.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		x4 := o2.Args[0]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s3 := o2.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s2 := o1.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_80(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		o1 := o0.Args[1]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		s2 := o1.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		o2 := o1.Args[1]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		s3 := o2.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		x4 := o2.Args[1]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := v.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		o1 := o0.Args[1]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		s2 := o1.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		o2 := o1.Args[1]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		x4 := o2.Args[0]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s3 := o2.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s0 := v.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		o1 := o0.Args[1]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		s3 := o2.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		x4 := o2.Args[1]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s2 := o1.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		s0 := v.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		s1 := o0.Args[0]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		o1 := o0.Args[1]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		x4 := o2.Args[0]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s3 := o2.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s2 := o1.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		s0 := v.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		s2 := o1.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o2 := o1.Args[1]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		s3 := o2.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		x4 := o2.Args[1]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s0 := v.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		s2 := o1.Args[0]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o2 := o1.Args[1]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		x4 := o2.Args[0]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s3 := o2.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s0 := v.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		s3 := o2.Args[0]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		s := x3.Aux
		_ = x3.Args[1]
		p := x3.Args[0]
		mem := x3.Args[1]
		x4 := o2.Args[1]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		if p != x4_0.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s2 := o1.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s0 := v.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpPPC64OR {
			break
		}
		if o0.Type != t {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpPPC64OR {
			break
		}
		if o1.Type != t {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpPPC64OR {
			break
		}
		if o2.Type != t {
			break
		}
		_ = o2.Args[1]
		x4 := o2.Args[0]
		if x4.Op != OpPPC64MOVWBRload {
			break
		}
		if x4.Type != t {
			break
		}
		_ = x4.Args[1]
		x4_0 := x4.Args[0]
		if x4_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x4_0.Type != typ.Uintptr {
			break
		}
		i4 := x4_0.AuxInt
		p := x4_0.Args[0]
		mem := x4.Args[1]
		s3 := o2.Args[1]
		if s3.Op != OpPPC64SLDconst {
			break
		}
		if s3.AuxInt != 32 {
			break
		}
		x3 := s3.Args[0]
		if x3.Op != OpPPC64MOVBZload {
			break
		}
		i3 := x3.AuxInt
		s := x3.Aux
		_ = x3.Args[1]
		if p != x3.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s2 := o1.Args[1]
		if s2.Op != OpPPC64SLDconst {
			break
		}
		if s2.AuxInt != 40 {
			break
		}
		x2 := s2.Args[0]
		if x2.Op != OpPPC64MOVBZload {
			break
		}
		i2 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		if p != x2.Args[0] {
			break
		}
		if mem != x2.Args[1] {
			break
		}
		s1 := o0.Args[1]
		if s1.Op != OpPPC64SLDconst {
			break
		}
		if s1.AuxInt != 48 {
			break
		}
		x1 := s1.Args[0]
		if x1.Op != OpPPC64MOVBZload {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		if p != x1.Args[0] {
			break
		}
		if mem != x1.Args[1] {
			break
		}
		s0 := v.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		x0 := s0.Args[0]
		if x0.Op != OpPPC64MOVBZload {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[1]
		if p != x0.Args[0] {
			break
		}
		if mem != x0.Args[1] {
			break
		}
		if !(!config.BigEndian && i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(s0) && clobber(s1) && clobber(s2) && clobber(s3)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_90(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_100(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		s := x5.Aux
		_ = x5.Args[1]
		p := x5.Args[0]
		mem := x5.Args[1]
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		s := x5.Aux
		_ = x5.Args[1]
		p := x5.Args[0]
		mem := x5.Args[1]
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		s := x4.Aux
		_ = x4.Args[1]
		p := x4.Args[0]
		mem := x4.Args[1]
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLWconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		s := x3_0.Aux
		p := x3_0.Args[0]
		mem := x3.Args[1]
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64OR_110(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x7 := v.Args[0]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o5 := v.Args[1]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		s6 := o5.Args[0]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		s := x6.Aux
		_ = x6.Args[1]
		p := x6.Args[0]
		mem := x6.Args[1]
		o4 := o5.Args[1]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		s := x5.Aux
		_ = x5.Args[1]
		p := x5.Args[0]
		mem := x5.Args[1]
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		s5 := o4.Args[0]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		s := x5.Aux
		_ = x5.Args[1]
		p := x5.Args[0]
		mem := x5.Args[1]
		o3 := o4.Args[1]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s4 := o3.Args[0]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		s := x4.Aux
		_ = x4.Args[1]
		p := x4.Args[0]
		mem := x4.Args[1]
		s0 := o3.Args[1]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		if x3_0.Aux != s {
			break
		}
		if p != x3_0.Args[0] {
			break
		}
		if mem != x3.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o5 := v.Args[0]
		if o5.Op != OpPPC64OR {
			break
		}
		if o5.Type != t {
			break
		}
		_ = o5.Args[1]
		o4 := o5.Args[0]
		if o4.Op != OpPPC64OR {
			break
		}
		if o4.Type != t {
			break
		}
		_ = o4.Args[1]
		o3 := o4.Args[0]
		if o3.Op != OpPPC64OR {
			break
		}
		if o3.Type != t {
			break
		}
		_ = o3.Args[1]
		s0 := o3.Args[0]
		if s0.Op != OpPPC64SLDconst {
			break
		}
		if s0.AuxInt != 32 {
			break
		}
		x3 := s0.Args[0]
		if x3.Op != OpPPC64MOVWBRload {
			break
		}
		if x3.Type != t {
			break
		}
		_ = x3.Args[1]
		x3_0 := x3.Args[0]
		if x3_0.Op != OpPPC64MOVDaddr {
			break
		}
		if x3_0.Type != typ.Uintptr {
			break
		}
		i0 := x3_0.AuxInt
		s := x3_0.Aux
		p := x3_0.Args[0]
		mem := x3.Args[1]
		s4 := o3.Args[1]
		if s4.Op != OpPPC64SLDconst {
			break
		}
		if s4.AuxInt != 24 {
			break
		}
		x4 := s4.Args[0]
		if x4.Op != OpPPC64MOVBZload {
			break
		}
		i4 := x4.AuxInt
		if x4.Aux != s {
			break
		}
		_ = x4.Args[1]
		if p != x4.Args[0] {
			break
		}
		if mem != x4.Args[1] {
			break
		}
		s5 := o4.Args[1]
		if s5.Op != OpPPC64SLDconst {
			break
		}
		if s5.AuxInt != 16 {
			break
		}
		x5 := s5.Args[0]
		if x5.Op != OpPPC64MOVBZload {
			break
		}
		i5 := x5.AuxInt
		if x5.Aux != s {
			break
		}
		_ = x5.Args[1]
		if p != x5.Args[0] {
			break
		}
		if mem != x5.Args[1] {
			break
		}
		s6 := o5.Args[1]
		if s6.Op != OpPPC64SLDconst {
			break
		}
		if s6.AuxInt != 8 {
			break
		}
		x6 := s6.Args[0]
		if x6.Op != OpPPC64MOVBZload {
			break
		}
		i6 := x6.AuxInt
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		if p != x6.Args[0] {
			break
		}
		if mem != x6.Args[1] {
			break
		}
		x7 := v.Args[1]
		if x7.Op != OpPPC64MOVBZload {
			break
		}
		i7 := x7.AuxInt
		if x7.Aux != s {
			break
		}
		_ = x7.Args[1]
		if p != x7.Args[0] {
			break
		}
		if mem != x7.Args[1] {
			break
		}
		if !(!config.BigEndian && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && s4.Uses == 1 && s5.Uses == 1 && s6.Uses == 1 && mergePoint(b, x3, x4, x5, x6, x7) != nil && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0) && clobber(s4) && clobber(s5) && clobber(s6)) {
			break
		}
		b = mergePoint(b, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDBRload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDaddr, typ.Uintptr)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64ORN_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
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
	return false
}
func rewriteValuePPC64_OpPPC64ORconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64ORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpPPC64ORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != -1 {
			break
		}
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = -1
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64SUB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(-c)) {
			break
		}
		v.reset(OpPPC64ADDconst)
		v.AuxInt = -c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64XOR_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRDconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpPPC64ROTLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLDconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpPPC64ROTLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRWconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpPPC64ROTLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLWconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpPPC64ROTLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1.Type != typ.Int64 {
			break
		}
		if v_0_1.AuxInt != 63 {
			break
		}
		y := v_0_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRD {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64SUB {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		if v_1_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_0.AuxInt != 64 {
			break
		}
		v_1_1_1 := v_1_1.Args[1]
		if v_1_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1_1.Type != typ.UInt {
			break
		}
		if v_1_1_1.AuxInt != 63 {
			break
		}
		if y != v_1_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64SUB {
			break
		}
		if v_0_1.Type != typ.UInt {
			break
		}
		_ = v_0_1.Args[1]
		v_0_1_0 := v_0_1.Args[0]
		if v_0_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_0_1_0.AuxInt != 64 {
			break
		}
		v_0_1_1 := v_0_1.Args[1]
		if v_0_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1_1.Type != typ.UInt {
			break
		}
		if v_0_1_1.AuxInt != 63 {
			break
		}
		y := v_0_1_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLD {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1.Type != typ.Int64 {
			break
		}
		if v_1_1.AuxInt != 63 {
			break
		}
		if y != v_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SLW {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1.Type != typ.Int32 {
			break
		}
		if v_0_1.AuxInt != 31 {
			break
		}
		y := v_0_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SRW {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64SUB {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		if v_1_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_0.AuxInt != 32 {
			break
		}
		v_1_1_1 := v_1_1.Args[1]
		if v_1_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1_1.Type != typ.UInt {
			break
		}
		if v_1_1_1.AuxInt != 31 {
			break
		}
		if y != v_1_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64SRW {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpPPC64SUB {
			break
		}
		if v_0_1.Type != typ.UInt {
			break
		}
		_ = v_0_1.Args[1]
		v_0_1_0 := v_0_1.Args[0]
		if v_0_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_0_1_0.AuxInt != 32 {
			break
		}
		v_0_1_1 := v_0_1.Args[1]
		if v_0_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_0_1_1.Type != typ.UInt {
			break
		}
		if v_0_1_1.AuxInt != 31 {
			break
		}
		y := v_0_1_1.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SLW {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1.Type != typ.Int32 {
			break
		}
		if v_1_1.AuxInt != 31 {
			break
		}
		if y != v_1_1.Args[0] {
			break
		}
		v.reset(OpPPC64ROTLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c ^ d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = c ^ d
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64XOR_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpPPC64XORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpPPC64XORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPPC64XORconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpPPC64XORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpPPC64XORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuePPC64_OpPopCount16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64POPCNTW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpPopCount32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64POPCNTW)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWZreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpPopCount64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64POPCNTD)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpPopCount8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpPPC64POPCNTB)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpRound_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64FROUND)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpRound32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64LoweredRound32F)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpRound64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64LoweredRound64F)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh16Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
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
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
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
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -16
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh32Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
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
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1.AuxInt != 31 {
			break
		}
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int32)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 31 {
			break
		}
		y := v_1.Args[1]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int32)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		if v_1.AuxInt != 31 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 32 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		if v_1_1.AuxInt != 31 {
			break
		}
		y := v_1_1.Args[0]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 32
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 31
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 32 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		y := v_1_1.Args[0]
		v_1_1_1 := v_1_1.Args[1]
		if v_1_1_1.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_1.AuxInt != 31 {
			break
		}
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 32
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 31
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 32 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		if v_1_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_0.AuxInt != 31 {
			break
		}
		y := v_1_1.Args[1]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 32
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 31
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
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
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = 63
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1.AuxInt != 31 {
			break
		}
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int32)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 31 {
			break
		}
		y := v_1.Args[1]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int32)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		if v_1.AuxInt != 31 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v0.AuxInt = 31
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 32 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		if v_1_1.AuxInt != 31 {
			break
		}
		y := v_1_1.Args[0]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 32
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 31
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 32 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		y := v_1_1.Args[0]
		v_1_1_1 := v_1_1.Args[1]
		if v_1_1_1.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_1.AuxInt != 31 {
			break
		}
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 32
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 31
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 32 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		if v_1_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_0.AuxInt != 31 {
			break
		}
		y := v_1_1.Args[1]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 32
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 31
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -32
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh64Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh64Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh64Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
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
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1.AuxInt != 63 {
			break
		}
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int64)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 63 {
			break
		}
		y := v_1.Args[1]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int64)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		if v_1.AuxInt != 63 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 64 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		if v_1_1.AuxInt != 63 {
			break
		}
		y := v_1_1.Args[0]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 64
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 63
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 64 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		y := v_1_1.Args[0]
		v_1_1_1 := v_1_1.Args[1]
		if v_1_1_1.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_1.AuxInt != 63 {
			break
		}
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 64
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 63
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 64 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		if v_1_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_0.AuxInt != 63 {
			break
		}
		y := v_1_1.Args[1]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 64
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 63
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh64Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
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
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = 63
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpPPC64SRADconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1.AuxInt != 63 {
			break
		}
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int64)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64AND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 63 {
			break
		}
		y := v_1.Args[1]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.Int64)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		if v_1.AuxInt != 63 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v0.AuxInt = 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 64 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64ANDconst {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		if v_1_1.AuxInt != 63 {
			break
		}
		y := v_1_1.Args[0]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 64
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 63
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 64 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		y := v_1_1.Args[0]
		v_1_1_1 := v_1_1.Args[1]
		if v_1_1_1.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_1.AuxInt != 63 {
			break
		}
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 64
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 63
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64SUB {
			break
		}
		if v_1.Type != typ.UInt {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_0.AuxInt != 64 {
			break
		}
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpPPC64AND {
			break
		}
		if v_1_1.Type != typ.UInt {
			break
		}
		_ = v_1_1.Args[1]
		v_1_1_0 := v_1_1.Args[0]
		if v_1_1_0.Op != OpPPC64MOVDconst {
			break
		}
		if v_1_1_0.AuxInt != 63 {
			break
		}
		y := v_1_1.Args[1]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64SUB, typ.UInt)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDconst, typ.Int64)
		v1.AuxInt = 64
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpPPC64ANDconst, typ.UInt)
		v2.AuxInt = 63
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v2 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v2.AuxInt = -64
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh8Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
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
		v.reset(OpPPC64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) < 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
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
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpPPC64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpPPC64SRAWconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SRAW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpPPC64ORN, typ.Int64)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpPPC64MaskIfNotCarry, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpPPC64ADDconstForCarry, psess.types.TypeFlags)
		v3.AuxInt = -8
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValuePPC64_OpSignExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSignExt16to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSignExt32to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVWreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSignExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSignExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSignExt8to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpSlicemask_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpPPC64SRADconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpPPC64NEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValuePPC64_OpSqrt_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64FSQRT)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpStaticCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpPPC64CALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpStore_0(v *Value) bool {

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 8 && psess.is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpPPC64FMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 8 && psess.is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpPPC64FMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 4 && psess.is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpPPC64FMOVSstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 8 && (psess.is64BitInt(val.Type) || isPtr(val.Type))) {
			break
		}
		v.reset(OpPPC64MOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 4 && psess.is32BitInt(val.Type)) {
			break
		}
		v.reset(OpPPC64MOVWstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 2) {
			break
		}
		v.reset(OpPPC64MOVHstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 1) {
			break
		}
		v.reset(OpPPC64MOVBstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuePPC64_OpSub16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSub32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSub32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FSUBS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSub64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSub64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64FSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSub8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpSubPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpTrunc_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64FTRUNC)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc16to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc32to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc32to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc64to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc64to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVWreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpTrunc64to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpWB_0(v *Value) bool {

	for {
		fn := v.Aux
		_ = v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpPPC64LoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpXor16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpXor32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpXor64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValuePPC64_OpXor8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpPPC64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValuePPC64_OpZero_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		if v.AuxInt != 0 {
			break
		}
		_ = v.Args[1]
		mem := v.Args[1]
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVBstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVHstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 3 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = 2
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHstorezero, psess.types.TypeMem)
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVWstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 5 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = 4
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWstorezero, psess.types.TypeMem)
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 6 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVHstorezero)
		v.AuxInt = 4
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWstorezero, psess.types.TypeMem)
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 7 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVBstorezero)
		v.AuxInt = 6
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVHstorezero, psess.types.TypeMem)
		v0.AuxInt = 4
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVWstorezero, psess.types.TypeMem)
		v1.AddArg(destptr)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = 4
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVWstorezero, psess.types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValuePPC64_OpZero_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		if v.AuxInt != 12 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVWstorezero)
		v.AuxInt = 8
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, psess.types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 8
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, psess.types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 24 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 16
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, psess.types.TypeMem)
		v0.AuxInt = 8
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(destptr)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 32 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpPPC64MOVDstorezero)
		v.AuxInt = 24
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, psess.types.TypeMem)
		v0.AuxInt = 16
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, psess.types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpPPC64MOVDstorezero, psess.types.TypeMem)
		v2.AuxInt = 0
		v2.AddArg(destptr)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpPPC64LoweredZero)
		v.AuxInt = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt16to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVHZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt32to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVWZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValuePPC64_OpZeroExt8to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpPPC64MOVBZreg)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteBlockPPC64(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	case BlockPPC64EQ:

		for {
			v := b.Control
			if v.Op != OpPPC64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockPPC64EQ
			v0 := b.NewValue0(v.Pos, OpPPC64ANDCCconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockPPC64EQ
			v0 := b.NewValue0(v.Pos, OpPPC64ANDCCconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockPPC64EQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockPPC64GE:

		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockPPC64LE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockPPC64GT:

		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockPPC64LT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockIf:

		for {
			v := b.Control
			if v.Op != OpPPC64Equal {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockPPC64EQ
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64NotEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockPPC64NE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64LessThan {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockPPC64LT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64LessEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockPPC64LE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64GreaterThan {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockPPC64GT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64GreaterEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockPPC64GE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FLessThan {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockPPC64FLT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FLessEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockPPC64FLE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FGreaterThan {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockPPC64FGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FGreaterEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockPPC64FGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			_ = v
			cond := b.Control
			b.Kind = BlockPPC64NE
			v0 := b.NewValue0(v.Pos, OpPPC64CMPWconst, psess.types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(cond)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockPPC64LE:

		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockPPC64GE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockPPC64LT:

		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockPPC64GT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockPPC64NE:

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64Equal {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockPPC64EQ
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64NotEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockPPC64NE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64LessThan {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockPPC64LT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64LessEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockPPC64LE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64GreaterThan {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockPPC64GT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64GreaterEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockPPC64GE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64FLessThan {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockPPC64FLT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64FLessEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockPPC64FLE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64FGreaterThan {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockPPC64FGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64FGreaterEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockPPC64FGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockPPC64NE
			v0 := b.NewValue0(v.Pos, OpPPC64ANDCCconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpPPC64ANDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockPPC64NE
			v0 := b.NewValue0(v.Pos, OpPPC64ANDCCconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagLT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64FlagGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpPPC64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockPPC64NE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	}
	return false
}

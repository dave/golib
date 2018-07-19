package ssa

import "github.com/dave/golib/src/cmd/compile/internal/types"

// in case not otherwise used
// in case not otherwise used
// in case not otherwise used
// in case not otherwise used

func (psess *PackageSession) rewriteValueMIPS64(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValueMIPS64_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValueMIPS64_OpAdd32_0(v)
	case OpAdd32F:
		return rewriteValueMIPS64_OpAdd32F_0(v)
	case OpAdd64:
		return rewriteValueMIPS64_OpAdd64_0(v)
	case OpAdd64F:
		return rewriteValueMIPS64_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValueMIPS64_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValueMIPS64_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValueMIPS64_OpAddr_0(v)
	case OpAnd16:
		return rewriteValueMIPS64_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValueMIPS64_OpAnd32_0(v)
	case OpAnd64:
		return rewriteValueMIPS64_OpAnd64_0(v)
	case OpAnd8:
		return rewriteValueMIPS64_OpAnd8_0(v)
	case OpAndB:
		return rewriteValueMIPS64_OpAndB_0(v)
	case OpAtomicAdd32:
		return rewriteValueMIPS64_OpAtomicAdd32_0(v)
	case OpAtomicAdd64:
		return rewriteValueMIPS64_OpAtomicAdd64_0(v)
	case OpAtomicCompareAndSwap32:
		return rewriteValueMIPS64_OpAtomicCompareAndSwap32_0(v)
	case OpAtomicCompareAndSwap64:
		return rewriteValueMIPS64_OpAtomicCompareAndSwap64_0(v)
	case OpAtomicExchange32:
		return rewriteValueMIPS64_OpAtomicExchange32_0(v)
	case OpAtomicExchange64:
		return rewriteValueMIPS64_OpAtomicExchange64_0(v)
	case OpAtomicLoad32:
		return rewriteValueMIPS64_OpAtomicLoad32_0(v)
	case OpAtomicLoad64:
		return rewriteValueMIPS64_OpAtomicLoad64_0(v)
	case OpAtomicLoadPtr:
		return rewriteValueMIPS64_OpAtomicLoadPtr_0(v)
	case OpAtomicStore32:
		return rewriteValueMIPS64_OpAtomicStore32_0(v)
	case OpAtomicStore64:
		return rewriteValueMIPS64_OpAtomicStore64_0(v)
	case OpAtomicStorePtrNoWB:
		return rewriteValueMIPS64_OpAtomicStorePtrNoWB_0(v)
	case OpAvg64u:
		return rewriteValueMIPS64_OpAvg64u_0(v)
	case OpClosureCall:
		return rewriteValueMIPS64_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValueMIPS64_OpCom16_0(v)
	case OpCom32:
		return rewriteValueMIPS64_OpCom32_0(v)
	case OpCom64:
		return rewriteValueMIPS64_OpCom64_0(v)
	case OpCom8:
		return rewriteValueMIPS64_OpCom8_0(v)
	case OpConst16:
		return rewriteValueMIPS64_OpConst16_0(v)
	case OpConst32:
		return rewriteValueMIPS64_OpConst32_0(v)
	case OpConst32F:
		return rewriteValueMIPS64_OpConst32F_0(v)
	case OpConst64:
		return rewriteValueMIPS64_OpConst64_0(v)
	case OpConst64F:
		return rewriteValueMIPS64_OpConst64F_0(v)
	case OpConst8:
		return rewriteValueMIPS64_OpConst8_0(v)
	case OpConstBool:
		return rewriteValueMIPS64_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValueMIPS64_OpConstNil_0(v)
	case OpCvt32Fto32:
		return rewriteValueMIPS64_OpCvt32Fto32_0(v)
	case OpCvt32Fto64:
		return rewriteValueMIPS64_OpCvt32Fto64_0(v)
	case OpCvt32Fto64F:
		return rewriteValueMIPS64_OpCvt32Fto64F_0(v)
	case OpCvt32to32F:
		return rewriteValueMIPS64_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValueMIPS64_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValueMIPS64_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValueMIPS64_OpCvt64Fto32F_0(v)
	case OpCvt64Fto64:
		return rewriteValueMIPS64_OpCvt64Fto64_0(v)
	case OpCvt64to32F:
		return rewriteValueMIPS64_OpCvt64to32F_0(v)
	case OpCvt64to64F:
		return rewriteValueMIPS64_OpCvt64to64F_0(v)
	case OpDiv16:
		return rewriteValueMIPS64_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValueMIPS64_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValueMIPS64_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValueMIPS64_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValueMIPS64_OpDiv32u_0(v)
	case OpDiv64:
		return rewriteValueMIPS64_OpDiv64_0(v)
	case OpDiv64F:
		return rewriteValueMIPS64_OpDiv64F_0(v)
	case OpDiv64u:
		return rewriteValueMIPS64_OpDiv64u_0(v)
	case OpDiv8:
		return rewriteValueMIPS64_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValueMIPS64_OpDiv8u_0(v)
	case OpEq16:
		return rewriteValueMIPS64_OpEq16_0(v)
	case OpEq32:
		return rewriteValueMIPS64_OpEq32_0(v)
	case OpEq32F:
		return psess.rewriteValueMIPS64_OpEq32F_0(v)
	case OpEq64:
		return rewriteValueMIPS64_OpEq64_0(v)
	case OpEq64F:
		return psess.rewriteValueMIPS64_OpEq64F_0(v)
	case OpEq8:
		return rewriteValueMIPS64_OpEq8_0(v)
	case OpEqB:
		return rewriteValueMIPS64_OpEqB_0(v)
	case OpEqPtr:
		return rewriteValueMIPS64_OpEqPtr_0(v)
	case OpGeq16:
		return rewriteValueMIPS64_OpGeq16_0(v)
	case OpGeq16U:
		return rewriteValueMIPS64_OpGeq16U_0(v)
	case OpGeq32:
		return rewriteValueMIPS64_OpGeq32_0(v)
	case OpGeq32F:
		return psess.rewriteValueMIPS64_OpGeq32F_0(v)
	case OpGeq32U:
		return rewriteValueMIPS64_OpGeq32U_0(v)
	case OpGeq64:
		return rewriteValueMIPS64_OpGeq64_0(v)
	case OpGeq64F:
		return psess.rewriteValueMIPS64_OpGeq64F_0(v)
	case OpGeq64U:
		return rewriteValueMIPS64_OpGeq64U_0(v)
	case OpGeq8:
		return rewriteValueMIPS64_OpGeq8_0(v)
	case OpGeq8U:
		return rewriteValueMIPS64_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValueMIPS64_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValueMIPS64_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValueMIPS64_OpGetClosurePtr_0(v)
	case OpGreater16:
		return rewriteValueMIPS64_OpGreater16_0(v)
	case OpGreater16U:
		return rewriteValueMIPS64_OpGreater16U_0(v)
	case OpGreater32:
		return rewriteValueMIPS64_OpGreater32_0(v)
	case OpGreater32F:
		return psess.rewriteValueMIPS64_OpGreater32F_0(v)
	case OpGreater32U:
		return rewriteValueMIPS64_OpGreater32U_0(v)
	case OpGreater64:
		return rewriteValueMIPS64_OpGreater64_0(v)
	case OpGreater64F:
		return psess.rewriteValueMIPS64_OpGreater64F_0(v)
	case OpGreater64U:
		return rewriteValueMIPS64_OpGreater64U_0(v)
	case OpGreater8:
		return rewriteValueMIPS64_OpGreater8_0(v)
	case OpGreater8U:
		return rewriteValueMIPS64_OpGreater8U_0(v)
	case OpHmul32:
		return rewriteValueMIPS64_OpHmul32_0(v)
	case OpHmul32u:
		return rewriteValueMIPS64_OpHmul32u_0(v)
	case OpHmul64:
		return rewriteValueMIPS64_OpHmul64_0(v)
	case OpHmul64u:
		return rewriteValueMIPS64_OpHmul64u_0(v)
	case OpInterCall:
		return rewriteValueMIPS64_OpInterCall_0(v)
	case OpIsInBounds:
		return rewriteValueMIPS64_OpIsInBounds_0(v)
	case OpIsNonNil:
		return rewriteValueMIPS64_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return rewriteValueMIPS64_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return rewriteValueMIPS64_OpLeq16_0(v)
	case OpLeq16U:
		return rewriteValueMIPS64_OpLeq16U_0(v)
	case OpLeq32:
		return rewriteValueMIPS64_OpLeq32_0(v)
	case OpLeq32F:
		return psess.rewriteValueMIPS64_OpLeq32F_0(v)
	case OpLeq32U:
		return rewriteValueMIPS64_OpLeq32U_0(v)
	case OpLeq64:
		return rewriteValueMIPS64_OpLeq64_0(v)
	case OpLeq64F:
		return psess.rewriteValueMIPS64_OpLeq64F_0(v)
	case OpLeq64U:
		return rewriteValueMIPS64_OpLeq64U_0(v)
	case OpLeq8:
		return rewriteValueMIPS64_OpLeq8_0(v)
	case OpLeq8U:
		return rewriteValueMIPS64_OpLeq8U_0(v)
	case OpLess16:
		return rewriteValueMIPS64_OpLess16_0(v)
	case OpLess16U:
		return rewriteValueMIPS64_OpLess16U_0(v)
	case OpLess32:
		return rewriteValueMIPS64_OpLess32_0(v)
	case OpLess32F:
		return psess.rewriteValueMIPS64_OpLess32F_0(v)
	case OpLess32U:
		return rewriteValueMIPS64_OpLess32U_0(v)
	case OpLess64:
		return rewriteValueMIPS64_OpLess64_0(v)
	case OpLess64F:
		return psess.rewriteValueMIPS64_OpLess64F_0(v)
	case OpLess64U:
		return rewriteValueMIPS64_OpLess64U_0(v)
	case OpLess8:
		return rewriteValueMIPS64_OpLess8_0(v)
	case OpLess8U:
		return rewriteValueMIPS64_OpLess8U_0(v)
	case OpLoad:
		return psess.rewriteValueMIPS64_OpLoad_0(v)
	case OpLsh16x16:
		return rewriteValueMIPS64_OpLsh16x16_0(v)
	case OpLsh16x32:
		return rewriteValueMIPS64_OpLsh16x32_0(v)
	case OpLsh16x64:
		return rewriteValueMIPS64_OpLsh16x64_0(v)
	case OpLsh16x8:
		return rewriteValueMIPS64_OpLsh16x8_0(v)
	case OpLsh32x16:
		return rewriteValueMIPS64_OpLsh32x16_0(v)
	case OpLsh32x32:
		return rewriteValueMIPS64_OpLsh32x32_0(v)
	case OpLsh32x64:
		return rewriteValueMIPS64_OpLsh32x64_0(v)
	case OpLsh32x8:
		return rewriteValueMIPS64_OpLsh32x8_0(v)
	case OpLsh64x16:
		return rewriteValueMIPS64_OpLsh64x16_0(v)
	case OpLsh64x32:
		return rewriteValueMIPS64_OpLsh64x32_0(v)
	case OpLsh64x64:
		return rewriteValueMIPS64_OpLsh64x64_0(v)
	case OpLsh64x8:
		return rewriteValueMIPS64_OpLsh64x8_0(v)
	case OpLsh8x16:
		return rewriteValueMIPS64_OpLsh8x16_0(v)
	case OpLsh8x32:
		return rewriteValueMIPS64_OpLsh8x32_0(v)
	case OpLsh8x64:
		return rewriteValueMIPS64_OpLsh8x64_0(v)
	case OpLsh8x8:
		return rewriteValueMIPS64_OpLsh8x8_0(v)
	case OpMIPS64ADDV:
		return rewriteValueMIPS64_OpMIPS64ADDV_0(v)
	case OpMIPS64ADDVconst:
		return rewriteValueMIPS64_OpMIPS64ADDVconst_0(v)
	case OpMIPS64AND:
		return rewriteValueMIPS64_OpMIPS64AND_0(v)
	case OpMIPS64ANDconst:
		return rewriteValueMIPS64_OpMIPS64ANDconst_0(v)
	case OpMIPS64LoweredAtomicAdd32:
		return rewriteValueMIPS64_OpMIPS64LoweredAtomicAdd32_0(v)
	case OpMIPS64LoweredAtomicAdd64:
		return rewriteValueMIPS64_OpMIPS64LoweredAtomicAdd64_0(v)
	case OpMIPS64LoweredAtomicStore32:
		return rewriteValueMIPS64_OpMIPS64LoweredAtomicStore32_0(v)
	case OpMIPS64LoweredAtomicStore64:
		return rewriteValueMIPS64_OpMIPS64LoweredAtomicStore64_0(v)
	case OpMIPS64MOVBUload:
		return rewriteValueMIPS64_OpMIPS64MOVBUload_0(v)
	case OpMIPS64MOVBUreg:
		return rewriteValueMIPS64_OpMIPS64MOVBUreg_0(v)
	case OpMIPS64MOVBload:
		return rewriteValueMIPS64_OpMIPS64MOVBload_0(v)
	case OpMIPS64MOVBreg:
		return rewriteValueMIPS64_OpMIPS64MOVBreg_0(v)
	case OpMIPS64MOVBstore:
		return rewriteValueMIPS64_OpMIPS64MOVBstore_0(v)
	case OpMIPS64MOVBstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVBstorezero_0(v)
	case OpMIPS64MOVDload:
		return rewriteValueMIPS64_OpMIPS64MOVDload_0(v)
	case OpMIPS64MOVDstore:
		return rewriteValueMIPS64_OpMIPS64MOVDstore_0(v)
	case OpMIPS64MOVFload:
		return rewriteValueMIPS64_OpMIPS64MOVFload_0(v)
	case OpMIPS64MOVFstore:
		return rewriteValueMIPS64_OpMIPS64MOVFstore_0(v)
	case OpMIPS64MOVHUload:
		return rewriteValueMIPS64_OpMIPS64MOVHUload_0(v)
	case OpMIPS64MOVHUreg:
		return rewriteValueMIPS64_OpMIPS64MOVHUreg_0(v)
	case OpMIPS64MOVHload:
		return rewriteValueMIPS64_OpMIPS64MOVHload_0(v)
	case OpMIPS64MOVHreg:
		return rewriteValueMIPS64_OpMIPS64MOVHreg_0(v)
	case OpMIPS64MOVHstore:
		return rewriteValueMIPS64_OpMIPS64MOVHstore_0(v)
	case OpMIPS64MOVHstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVHstorezero_0(v)
	case OpMIPS64MOVVload:
		return rewriteValueMIPS64_OpMIPS64MOVVload_0(v)
	case OpMIPS64MOVVreg:
		return rewriteValueMIPS64_OpMIPS64MOVVreg_0(v)
	case OpMIPS64MOVVstore:
		return rewriteValueMIPS64_OpMIPS64MOVVstore_0(v)
	case OpMIPS64MOVVstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVVstorezero_0(v)
	case OpMIPS64MOVWUload:
		return rewriteValueMIPS64_OpMIPS64MOVWUload_0(v)
	case OpMIPS64MOVWUreg:
		return rewriteValueMIPS64_OpMIPS64MOVWUreg_0(v)
	case OpMIPS64MOVWload:
		return rewriteValueMIPS64_OpMIPS64MOVWload_0(v)
	case OpMIPS64MOVWreg:
		return rewriteValueMIPS64_OpMIPS64MOVWreg_0(v) || rewriteValueMIPS64_OpMIPS64MOVWreg_10(v)
	case OpMIPS64MOVWstore:
		return rewriteValueMIPS64_OpMIPS64MOVWstore_0(v)
	case OpMIPS64MOVWstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVWstorezero_0(v)
	case OpMIPS64NEGV:
		return rewriteValueMIPS64_OpMIPS64NEGV_0(v)
	case OpMIPS64NOR:
		return rewriteValueMIPS64_OpMIPS64NOR_0(v)
	case OpMIPS64NORconst:
		return rewriteValueMIPS64_OpMIPS64NORconst_0(v)
	case OpMIPS64OR:
		return rewriteValueMIPS64_OpMIPS64OR_0(v)
	case OpMIPS64ORconst:
		return rewriteValueMIPS64_OpMIPS64ORconst_0(v)
	case OpMIPS64SGT:
		return rewriteValueMIPS64_OpMIPS64SGT_0(v)
	case OpMIPS64SGTU:
		return rewriteValueMIPS64_OpMIPS64SGTU_0(v)
	case OpMIPS64SGTUconst:
		return rewriteValueMIPS64_OpMIPS64SGTUconst_0(v)
	case OpMIPS64SGTconst:
		return rewriteValueMIPS64_OpMIPS64SGTconst_0(v) || rewriteValueMIPS64_OpMIPS64SGTconst_10(v)
	case OpMIPS64SLLV:
		return rewriteValueMIPS64_OpMIPS64SLLV_0(v)
	case OpMIPS64SLLVconst:
		return rewriteValueMIPS64_OpMIPS64SLLVconst_0(v)
	case OpMIPS64SRAV:
		return rewriteValueMIPS64_OpMIPS64SRAV_0(v)
	case OpMIPS64SRAVconst:
		return rewriteValueMIPS64_OpMIPS64SRAVconst_0(v)
	case OpMIPS64SRLV:
		return rewriteValueMIPS64_OpMIPS64SRLV_0(v)
	case OpMIPS64SRLVconst:
		return rewriteValueMIPS64_OpMIPS64SRLVconst_0(v)
	case OpMIPS64SUBV:
		return rewriteValueMIPS64_OpMIPS64SUBV_0(v)
	case OpMIPS64SUBVconst:
		return rewriteValueMIPS64_OpMIPS64SUBVconst_0(v)
	case OpMIPS64XOR:
		return rewriteValueMIPS64_OpMIPS64XOR_0(v)
	case OpMIPS64XORconst:
		return rewriteValueMIPS64_OpMIPS64XORconst_0(v)
	case OpMod16:
		return rewriteValueMIPS64_OpMod16_0(v)
	case OpMod16u:
		return rewriteValueMIPS64_OpMod16u_0(v)
	case OpMod32:
		return rewriteValueMIPS64_OpMod32_0(v)
	case OpMod32u:
		return rewriteValueMIPS64_OpMod32u_0(v)
	case OpMod64:
		return rewriteValueMIPS64_OpMod64_0(v)
	case OpMod64u:
		return rewriteValueMIPS64_OpMod64u_0(v)
	case OpMod8:
		return rewriteValueMIPS64_OpMod8_0(v)
	case OpMod8u:
		return rewriteValueMIPS64_OpMod8u_0(v)
	case OpMove:
		return psess.rewriteValueMIPS64_OpMove_0(v) || psess.rewriteValueMIPS64_OpMove_10(v)
	case OpMul16:
		return rewriteValueMIPS64_OpMul16_0(v)
	case OpMul32:
		return rewriteValueMIPS64_OpMul32_0(v)
	case OpMul32F:
		return rewriteValueMIPS64_OpMul32F_0(v)
	case OpMul64:
		return rewriteValueMIPS64_OpMul64_0(v)
	case OpMul64F:
		return rewriteValueMIPS64_OpMul64F_0(v)
	case OpMul8:
		return rewriteValueMIPS64_OpMul8_0(v)
	case OpNeg16:
		return rewriteValueMIPS64_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValueMIPS64_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValueMIPS64_OpNeg32F_0(v)
	case OpNeg64:
		return rewriteValueMIPS64_OpNeg64_0(v)
	case OpNeg64F:
		return rewriteValueMIPS64_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValueMIPS64_OpNeg8_0(v)
	case OpNeq16:
		return rewriteValueMIPS64_OpNeq16_0(v)
	case OpNeq32:
		return rewriteValueMIPS64_OpNeq32_0(v)
	case OpNeq32F:
		return psess.rewriteValueMIPS64_OpNeq32F_0(v)
	case OpNeq64:
		return rewriteValueMIPS64_OpNeq64_0(v)
	case OpNeq64F:
		return psess.rewriteValueMIPS64_OpNeq64F_0(v)
	case OpNeq8:
		return rewriteValueMIPS64_OpNeq8_0(v)
	case OpNeqB:
		return rewriteValueMIPS64_OpNeqB_0(v)
	case OpNeqPtr:
		return rewriteValueMIPS64_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValueMIPS64_OpNilCheck_0(v)
	case OpNot:
		return rewriteValueMIPS64_OpNot_0(v)
	case OpOffPtr:
		return rewriteValueMIPS64_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValueMIPS64_OpOr16_0(v)
	case OpOr32:
		return rewriteValueMIPS64_OpOr32_0(v)
	case OpOr64:
		return rewriteValueMIPS64_OpOr64_0(v)
	case OpOr8:
		return rewriteValueMIPS64_OpOr8_0(v)
	case OpOrB:
		return rewriteValueMIPS64_OpOrB_0(v)
	case OpRound32F:
		return rewriteValueMIPS64_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValueMIPS64_OpRound64F_0(v)
	case OpRsh16Ux16:
		return rewriteValueMIPS64_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return rewriteValueMIPS64_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return rewriteValueMIPS64_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return rewriteValueMIPS64_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return rewriteValueMIPS64_OpRsh16x16_0(v)
	case OpRsh16x32:
		return rewriteValueMIPS64_OpRsh16x32_0(v)
	case OpRsh16x64:
		return rewriteValueMIPS64_OpRsh16x64_0(v)
	case OpRsh16x8:
		return rewriteValueMIPS64_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return rewriteValueMIPS64_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return rewriteValueMIPS64_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return rewriteValueMIPS64_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return rewriteValueMIPS64_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return rewriteValueMIPS64_OpRsh32x16_0(v)
	case OpRsh32x32:
		return rewriteValueMIPS64_OpRsh32x32_0(v)
	case OpRsh32x64:
		return rewriteValueMIPS64_OpRsh32x64_0(v)
	case OpRsh32x8:
		return rewriteValueMIPS64_OpRsh32x8_0(v)
	case OpRsh64Ux16:
		return rewriteValueMIPS64_OpRsh64Ux16_0(v)
	case OpRsh64Ux32:
		return rewriteValueMIPS64_OpRsh64Ux32_0(v)
	case OpRsh64Ux64:
		return rewriteValueMIPS64_OpRsh64Ux64_0(v)
	case OpRsh64Ux8:
		return rewriteValueMIPS64_OpRsh64Ux8_0(v)
	case OpRsh64x16:
		return rewriteValueMIPS64_OpRsh64x16_0(v)
	case OpRsh64x32:
		return rewriteValueMIPS64_OpRsh64x32_0(v)
	case OpRsh64x64:
		return rewriteValueMIPS64_OpRsh64x64_0(v)
	case OpRsh64x8:
		return rewriteValueMIPS64_OpRsh64x8_0(v)
	case OpRsh8Ux16:
		return rewriteValueMIPS64_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return rewriteValueMIPS64_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return rewriteValueMIPS64_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return rewriteValueMIPS64_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return rewriteValueMIPS64_OpRsh8x16_0(v)
	case OpRsh8x32:
		return rewriteValueMIPS64_OpRsh8x32_0(v)
	case OpRsh8x64:
		return rewriteValueMIPS64_OpRsh8x64_0(v)
	case OpRsh8x8:
		return rewriteValueMIPS64_OpRsh8x8_0(v)
	case OpSelect0:
		return rewriteValueMIPS64_OpSelect0_0(v)
	case OpSelect1:
		return rewriteValueMIPS64_OpSelect1_0(v) || rewriteValueMIPS64_OpSelect1_10(v) || rewriteValueMIPS64_OpSelect1_20(v)
	case OpSignExt16to32:
		return rewriteValueMIPS64_OpSignExt16to32_0(v)
	case OpSignExt16to64:
		return rewriteValueMIPS64_OpSignExt16to64_0(v)
	case OpSignExt32to64:
		return rewriteValueMIPS64_OpSignExt32to64_0(v)
	case OpSignExt8to16:
		return rewriteValueMIPS64_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValueMIPS64_OpSignExt8to32_0(v)
	case OpSignExt8to64:
		return rewriteValueMIPS64_OpSignExt8to64_0(v)
	case OpSlicemask:
		return rewriteValueMIPS64_OpSlicemask_0(v)
	case OpSqrt:
		return rewriteValueMIPS64_OpSqrt_0(v)
	case OpStaticCall:
		return rewriteValueMIPS64_OpStaticCall_0(v)
	case OpStore:
		return psess.rewriteValueMIPS64_OpStore_0(v)
	case OpSub16:
		return rewriteValueMIPS64_OpSub16_0(v)
	case OpSub32:
		return rewriteValueMIPS64_OpSub32_0(v)
	case OpSub32F:
		return rewriteValueMIPS64_OpSub32F_0(v)
	case OpSub64:
		return rewriteValueMIPS64_OpSub64_0(v)
	case OpSub64F:
		return rewriteValueMIPS64_OpSub64F_0(v)
	case OpSub8:
		return rewriteValueMIPS64_OpSub8_0(v)
	case OpSubPtr:
		return rewriteValueMIPS64_OpSubPtr_0(v)
	case OpTrunc16to8:
		return rewriteValueMIPS64_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValueMIPS64_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValueMIPS64_OpTrunc32to8_0(v)
	case OpTrunc64to16:
		return rewriteValueMIPS64_OpTrunc64to16_0(v)
	case OpTrunc64to32:
		return rewriteValueMIPS64_OpTrunc64to32_0(v)
	case OpTrunc64to8:
		return rewriteValueMIPS64_OpTrunc64to8_0(v)
	case OpWB:
		return rewriteValueMIPS64_OpWB_0(v)
	case OpXor16:
		return rewriteValueMIPS64_OpXor16_0(v)
	case OpXor32:
		return rewriteValueMIPS64_OpXor32_0(v)
	case OpXor64:
		return rewriteValueMIPS64_OpXor64_0(v)
	case OpXor8:
		return rewriteValueMIPS64_OpXor8_0(v)
	case OpZero:
		return psess.rewriteValueMIPS64_OpZero_0(v) || psess.rewriteValueMIPS64_OpZero_10(v)
	case OpZeroExt16to32:
		return rewriteValueMIPS64_OpZeroExt16to32_0(v)
	case OpZeroExt16to64:
		return rewriteValueMIPS64_OpZeroExt16to64_0(v)
	case OpZeroExt32to64:
		return rewriteValueMIPS64_OpZeroExt32to64_0(v)
	case OpZeroExt8to16:
		return rewriteValueMIPS64_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValueMIPS64_OpZeroExt8to32_0(v)
	case OpZeroExt8to64:
		return rewriteValueMIPS64_OpZeroExt8to64_0(v)
	}
	return false
}
func rewriteValueMIPS64_OpAdd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAdd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAdd32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAdd64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAdd64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAdd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAddPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAddr_0(v *Value) bool {

	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpMIPS64MOVVaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueMIPS64_OpAnd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAnd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAnd64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAnd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAndB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicAdd32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64LoweredAtomicAdd32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicAdd64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64LoweredAtomicAdd64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicCompareAndSwap32_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpMIPS64LoweredAtomicCas32)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicCompareAndSwap64_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpMIPS64LoweredAtomicCas64)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicExchange32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64LoweredAtomicExchange32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicExchange64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64LoweredAtomicExchange64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicLoad32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64LoweredAtomicLoad32)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicLoad64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64LoweredAtomicLoad64)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicLoadPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64LoweredAtomicLoad64)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicStore32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64LoweredAtomicStore32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicStore64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64LoweredAtomicStore64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAtomicStorePtrNoWB_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64LoweredAtomicStore64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpAvg64u_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64ADDV)
		v0 := b.NewValue0(v.Pos, OpMIPS64SRLVconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpMIPS64SUBV, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpClosureCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64CALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpCom16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCom32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCom64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCom8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpConst16_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConst32_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConst32F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVFconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConst64_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConst64F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConst8_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS64_OpConstBool_0(v *Value) bool {

	for {
		b := v.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValueMIPS64_OpConstNil_0(v *Value) bool {

	for {
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueMIPS64_OpCvt32Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64TRUNCFW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt32Fto64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64TRUNCFV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt32Fto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVFD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt32to32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVWF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt32to64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVWD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt64Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64TRUNCDW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt64Fto32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVDF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt64Fto64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64TRUNCDV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt64to32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVVF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCvt64to64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVVD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpDiv16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64DIVF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpDiv32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64DIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpDiv64u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpEq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpEq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpEq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpEq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpEq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEqB_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.Bool)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEqPtr_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpGeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGEF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpGeq32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v1.AddArg(y)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpGeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGED, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpGeq64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v1.AddArg(y)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(x)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGetCallerPC_0(v *Value) bool {

	for {
		v.reset(OpMIPS64LoweredGetCallerPC)
		return true
	}
}
func rewriteValueMIPS64_OpGetCallerSP_0(v *Value) bool {

	for {
		v.reset(OpMIPS64LoweredGetCallerSP)
		return true
	}
}
func rewriteValueMIPS64_OpGetClosurePtr_0(v *Value) bool {

	for {
		v.reset(OpMIPS64LoweredGetClosurePtr)
		return true
	}
}
func rewriteValueMIPS64_OpGreater16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGreater16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGreater32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpGreater32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpGreater32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGreater64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpGreater64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpGreater64U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpGreater8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGreater8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpHmul32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpSelect1, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpMIPS64MULV, types.NewTuple(typ.Int64, typ.Int64))
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpHmul32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRLVconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpSelect1, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpHmul64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULV, types.NewTuple(typ.Int64, typ.Int64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpHmul64u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpInterCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[1]
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64CALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpIsInBounds_0(v *Value) bool {

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v.AddArg(len)
		v.AddArg(idx)
		return true
	}
}
func rewriteValueMIPS64_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		ptr := v.Args[0]
		v.reset(OpMIPS64SGTU)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpIsSliceInBounds_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v1.AddArg(idx)
		v1.AddArg(len)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpLeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGEF, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLeq32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpLeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGED, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLeq64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpLess32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTF, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLess32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpLess64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTD, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLess64U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpLess8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpLoad_0(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpMIPS64MOVBUload)
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
		v.reset(OpMIPS64MOVBload)
		v.AddArg(ptr)
		v.AddArg(mem)
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
		v.reset(OpMIPS64MOVBUload)
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
		v.reset(OpMIPS64MOVHload)
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
		v.reset(OpMIPS64MOVHUload)
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
		v.reset(OpMIPS64MOVWload)
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
		v.reset(OpMIPS64MOVWUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpMIPS64MOVVload)
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
		v.reset(OpMIPS64MOVFload)
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
		v.reset(OpMIPS64MOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpMIPS64ADDV_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64NEGV {
			break
		}
		y := v_1.Args[0]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64NEGV {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64ADDVconst_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpMIPS64MOVVaddr)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
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
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c + d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64SUBVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c - d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = c - d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64AND_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ANDconst)
		v.AuxInt = c
		v.AddArg(x)
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
	return false
}
func rewriteValueMIPS64_OpMIPS64ANDconst_0(v *Value) bool {

	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
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
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c & d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPS64ANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64LoweredAtomicAdd32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64LoweredAtomicAddconst32)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64LoweredAtomicAdd64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64LoweredAtomicAddconst64)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64LoweredAtomicStore32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPS64LoweredAtomicStorezero32)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64LoweredAtomicStore64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPS64LoweredAtomicStorezero64)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBUload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBUreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstore)
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
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstore)
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
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstorezero)
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
		if v_1.Op != OpMIPS64MOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
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
		if v_1.Op != OpMIPS64MOVBUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
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
		if v_1.Op != OpMIPS64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
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
		if v_1.Op != OpMIPS64MOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
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
		if v_1.Op != OpMIPS64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
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
		if v_1.Op != OpMIPS64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBstorezero_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVDload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVDstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVDstore)
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
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVFload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVFload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVFload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVFstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVFstore)
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
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHUload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHUreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
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
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
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
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPS64MOVHstorezero)
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
		if v_1.Op != OpMIPS64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVHstore)
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
		if v_1.Op != OpMIPS64MOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVHstore)
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
		if v_1.Op != OpMIPS64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVHstore)
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
		if v_1.Op != OpMIPS64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHstorezero_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVVload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVVload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpMIPS64MOVVnop)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
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
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
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
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPS64MOVVstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVstorezero_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWUload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWUreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVWUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVWUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint32(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVWload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVHreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPS64MOVWreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWreg_10(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(int32(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
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
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
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
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPS64MOVWstorezero)
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
		if v_1.Op != OpMIPS64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVWstore)
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
		if v_1.Op != OpMIPS64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWstorezero_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64NEGV_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = -c
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64NOR_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64NORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64NORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64NORconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = ^(c | d)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64OR_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64ORconst)
		v.AuxInt = c
		v.AddArg(x)
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
	return false
}
func rewriteValueMIPS64_OpMIPS64ORconst_0(v *Value) bool {

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
		if v.AuxInt != -1 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = -1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c | d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c | d)) {
			break
		}
		v.reset(OpMIPS64ORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGT_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64SGTconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGTU_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64SGTUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGTUconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(uint64(c) > uint64(d)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(uint64(c) <= uint64(d)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVBUreg {
			break
		}
		if !(0xff < uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVHUreg {
			break
		}
		if !(0xffff < uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ANDconst {
			break
		}
		m := v_0.AuxInt
		if !(uint64(m) < uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64SRLVconst {
			break
		}
		d := v_0.AuxInt
		if !(0 < d && d <= 63 && 1<<uint64(64-d) <= uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGTconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(c > d) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(c <= d) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVBreg {
			break
		}
		if !(0x7f < c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVBreg {
			break
		}
		if !(c <= -0x80) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVBUreg {
			break
		}
		if !(0xff < c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVBUreg {
			break
		}
		if !(c < 0) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVHreg {
			break
		}
		if !(0x7fff < c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVHreg {
			break
		}
		if !(c <= -0x8000) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVHUreg {
			break
		}
		if !(0xffff < c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVHUreg {
			break
		}
		if !(c < 0) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGTconst_10(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVWUreg {
			break
		}
		if !(c < 0) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ANDconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= m && m < c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64SRLVconst {
			break
		}
		d := v_0.AuxInt
		if !(0 <= c && 0 < d && d <= 63 && 1<<uint64(64-d) <= c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SLLV_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPS64SLLVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SLLVconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = d << uint64(c)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRAV_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = 63
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRAVconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = d >> uint64(c)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRLV_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPS64SRLVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRLVconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint64(d) >> uint64(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SUBV_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64SUBVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SUBVconst_0(v *Value) bool {

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
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = d - c
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64SUBVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(-c - d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = -c - d
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(-c + d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = -c + d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64XOR_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64XORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64XORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64XORconst_0(v *Value) bool {

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
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		v.reset(OpMIPS64NORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c ^ d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64XORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c ^ d)) {
			break
		}
		v.reset(OpMIPS64XORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMod16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod64u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpMove_0(v *Value) bool {
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
		v.reset(OpMIPS64MOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
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
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Alignment(psess.types)%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
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
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 1
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v0.AuxInt = 1
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 4 {
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
		v.reset(OpMIPS64MOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
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
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Alignment(psess.types)%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
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
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v4.AuxInt = 1
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v6.AddArg(src)
		v6.AddArg(mem)
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
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
		if !(t.(*types.Type).Alignment(psess.types)%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
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
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
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
		if !(t.(*types.Type).Alignment(psess.types)%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v3.AuxInt = 2
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v4.AuxInt = 2
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v6.AddArg(src)
		v6.AddArg(mem)
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueMIPS64_OpMove_10(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		if v.AuxInt != 3 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v2.AuxInt = 1
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 6 {
			break
		}
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Alignment(psess.types)%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 12 {
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
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, psess.types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Alignment(psess.types)%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 24 {
			break
		}
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Alignment(psess.types)%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 16
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
		v0.AuxInt = 16
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, psess.types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
		v2.AuxInt = 8
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		s := v.AuxInt
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 24 || t.(*types.Type).Alignment(psess.types)%8 != 0) {
			break
		}
		v.reset(OpMIPS64LoweredMove)
		v.AuxInt = t.(*types.Type).Alignment(psess.types)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpMIPS64ADDVconst, src.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(psess.types), config)
		v0.AddArg(src)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMul16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMul32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMul32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64MULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpMul64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMul64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64MULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpMul8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpNeg16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeg32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeg32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeg64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeg64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeg8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpNeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpNeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpNeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpNeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpNeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64FPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpNeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpNeqB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpNeqPtr_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpNilCheck_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64LoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpNot_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64XORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpOffPtr_0(v *Value) bool {

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpMIPS64MOVVaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueMIPS64_OpOr16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpOr32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpOr64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpOr8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpOrB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpRound32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpRound64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg(v4)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3.AddArg(y)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg(v4)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3.AddArg(y)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v2.AddArg(v4)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v0.AddArg(v5)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v2.AddArg(v4)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v0.AddArg(v5)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = 63
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v2.AddArg(v4)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v0.AddArg(v5)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg(v4)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3.AddArg(y)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpSelect0_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64ANDconst)
		v.AuxInt = c - 1
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVV {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c % d
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint64(c) % uint64(d))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpSelect1_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != -1 {
			break
		}
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_0.AuxInt != -1 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_0.AuxInt != 0 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_0.AuxInt != 1 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64SLLVconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64SLLVconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_0.AuxInt != -1 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != -1 {
			break
		}
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpSelect1_10(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_0.AuxInt != 0 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_0.AuxInt != 1 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64SLLVconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64SLLVconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		if v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64SRLVconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c * d
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c * d
		return true
	}
	return false
}
func rewriteValueMIPS64_OpSelect1_20(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVV {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c / d
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint64(c) / uint64(d))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpSignExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSignExt16to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSignExt32to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVWreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSignExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSignExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSignExt8to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpSlicemask_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpSqrt_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64SQRTD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpStaticCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpMIPS64CALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpStore_0(v *Value) bool {

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 1) {
			break
		}
		v.reset(OpMIPS64MOVBstore)
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
		v.reset(OpMIPS64MOVHstore)
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
		if !(t.(*types.Type).Size(psess.types) == 4 && !psess.is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
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
		if !(t.(*types.Type).Size(psess.types) == 8 && !psess.is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
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
		v.reset(OpMIPS64MOVFstore)
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
		if !(t.(*types.Type).Size(psess.types) == 8 && psess.is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPS64MOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpSub16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSub32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSub32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSub64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSub64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSub8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpSubPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64SUBV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpTrunc16to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpTrunc32to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpTrunc32to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpTrunc64to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpTrunc64to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpTrunc64to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpWB_0(v *Value) bool {

	for {
		fn := v.Aux
		_ = v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPS64LoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS64_OpXor16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpXor32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpXor64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpXor8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPS64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS64_OpZero_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

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
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64MOVBstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 1
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 3
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 6
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v3.AuxInt = 2
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueMIPS64_OpZero_10(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		if v.AuxInt != 3 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, psess.types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 6 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, psess.types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 12 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 8
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, psess.types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, psess.types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 8
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 24 {
			break
		}
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.(*types.Type).Alignment(psess.types)%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 16
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, psess.types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, psess.types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		s := v.AuxInt
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(s%8 == 0 && s > 24 && s <= 8*128 && t.(*types.Type).Alignment(psess.types)%8 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpMIPS64DUFFZERO)
		v.AuxInt = 8 * (128 - s/8)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		s := v.AuxInt
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !((s > 8*128 || config.noDuffDevice) || t.(*types.Type).Alignment(psess.types)%8 != 0) {
			break
		}
		v.reset(OpMIPS64LoweredZero)
		v.AuxInt = t.(*types.Type).Alignment(psess.types)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64ADDVconst, ptr.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(psess.types), config)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpZeroExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVHUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpZeroExt16to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVHUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpZeroExt32to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVWUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpZeroExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpZeroExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpZeroExt8to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPS64MOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteBlockMIPS64(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	case BlockMIPS64EQ:

		for {
			v := b.Control
			if v.Op != OpMIPS64FPFlagTrue {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockMIPS64FPF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64FPFlagFalse {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockMIPS64FPT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGT {
				break
			}
			_ = cmp.Args[1]
			b.Kind = BlockMIPS64NE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTU {
				break
			}
			_ = cmp.Args[1]
			b.Kind = BlockMIPS64NE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTconst {
				break
			}
			b.Kind = BlockMIPS64NE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTUconst {
				break
			}
			b.Kind = BlockMIPS64NE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64SGTUconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPS64NE
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64SGTU {
				break
			}
			_ = v.Args[1]
			x := v.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != OpMIPS64MOVVconst {
				break
			}
			if v_1.AuxInt != 0 {
				break
			}
			b.Kind = BlockMIPS64EQ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64SGTconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPS64GEZ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64SGT {
				break
			}
			_ = v.Args[1]
			x := v.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != OpMIPS64MOVVconst {
				break
			}
			if v_1.AuxInt != 0 {
				break
			}
			b.Kind = BlockMIPS64LEZ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			if !(c != 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case BlockMIPS64GEZ:

		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			if !(c >= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			if !(c < 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case BlockMIPS64GTZ:

		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			if !(c > 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			if !(c <= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case BlockIf:

		for {
			v := b.Control
			_ = v
			cond := b.Control
			b.Kind = BlockMIPS64NE
			b.SetControl(cond)
			b.Aux = nil
			return true
		}
	case BlockMIPS64LEZ:

		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			if !(c <= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			if !(c > 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case BlockMIPS64LTZ:

		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			if !(c < 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			if !(c >= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case BlockMIPS64NE:

		for {
			v := b.Control
			if v.Op != OpMIPS64FPFlagTrue {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockMIPS64FPT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64FPFlagFalse {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockMIPS64FPF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGT {
				break
			}
			_ = cmp.Args[1]
			b.Kind = BlockMIPS64EQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTU {
				break
			}
			_ = cmp.Args[1]
			b.Kind = BlockMIPS64EQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTconst {
				break
			}
			b.Kind = BlockMIPS64EQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64XORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPS64SGTUconst {
				break
			}
			b.Kind = BlockMIPS64EQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64SGTUconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPS64EQ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64SGTU {
				break
			}
			_ = v.Args[1]
			x := v.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != OpMIPS64MOVVconst {
				break
			}
			if v_1.AuxInt != 0 {
				break
			}
			b.Kind = BlockMIPS64NE
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64SGTconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPS64LTZ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64SGT {
				break
			}
			_ = v.Args[1]
			x := v.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != OpMIPS64MOVVconst {
				break
			}
			if v_1.AuxInt != 0 {
				break
			}
			b.Kind = BlockMIPS64GTZ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			if v.AuxInt != 0 {
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
			if v.Op != OpMIPS64MOVVconst {
				break
			}
			c := v.AuxInt
			if !(c != 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
	}
	return false
}

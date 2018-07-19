package ssa

import "github.com/dave/golib/src/cmd/compile/internal/types"

// in case not otherwise used
// in case not otherwise used
// in case not otherwise used
// in case not otherwise used

func (psess *PackageSession) rewriteValueMIPS(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValueMIPS_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValueMIPS_OpAdd32_0(v)
	case OpAdd32F:
		return rewriteValueMIPS_OpAdd32F_0(v)
	case OpAdd32withcarry:
		return rewriteValueMIPS_OpAdd32withcarry_0(v)
	case OpAdd64F:
		return rewriteValueMIPS_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValueMIPS_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValueMIPS_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValueMIPS_OpAddr_0(v)
	case OpAnd16:
		return rewriteValueMIPS_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValueMIPS_OpAnd32_0(v)
	case OpAnd8:
		return rewriteValueMIPS_OpAnd8_0(v)
	case OpAndB:
		return rewriteValueMIPS_OpAndB_0(v)
	case OpAtomicAdd32:
		return rewriteValueMIPS_OpAtomicAdd32_0(v)
	case OpAtomicAnd8:
		return rewriteValueMIPS_OpAtomicAnd8_0(v)
	case OpAtomicCompareAndSwap32:
		return rewriteValueMIPS_OpAtomicCompareAndSwap32_0(v)
	case OpAtomicExchange32:
		return rewriteValueMIPS_OpAtomicExchange32_0(v)
	case OpAtomicLoad32:
		return rewriteValueMIPS_OpAtomicLoad32_0(v)
	case OpAtomicLoadPtr:
		return rewriteValueMIPS_OpAtomicLoadPtr_0(v)
	case OpAtomicOr8:
		return rewriteValueMIPS_OpAtomicOr8_0(v)
	case OpAtomicStore32:
		return rewriteValueMIPS_OpAtomicStore32_0(v)
	case OpAtomicStorePtrNoWB:
		return rewriteValueMIPS_OpAtomicStorePtrNoWB_0(v)
	case OpAvg32u:
		return rewriteValueMIPS_OpAvg32u_0(v)
	case OpBitLen32:
		return rewriteValueMIPS_OpBitLen32_0(v)
	case OpClosureCall:
		return rewriteValueMIPS_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValueMIPS_OpCom16_0(v)
	case OpCom32:
		return rewriteValueMIPS_OpCom32_0(v)
	case OpCom8:
		return rewriteValueMIPS_OpCom8_0(v)
	case OpConst16:
		return rewriteValueMIPS_OpConst16_0(v)
	case OpConst32:
		return rewriteValueMIPS_OpConst32_0(v)
	case OpConst32F:
		return rewriteValueMIPS_OpConst32F_0(v)
	case OpConst64F:
		return rewriteValueMIPS_OpConst64F_0(v)
	case OpConst8:
		return rewriteValueMIPS_OpConst8_0(v)
	case OpConstBool:
		return rewriteValueMIPS_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValueMIPS_OpConstNil_0(v)
	case OpCtz32:
		return rewriteValueMIPS_OpCtz32_0(v)
	case OpCtz32NonZero:
		return rewriteValueMIPS_OpCtz32NonZero_0(v)
	case OpCvt32Fto32:
		return rewriteValueMIPS_OpCvt32Fto32_0(v)
	case OpCvt32Fto64F:
		return rewriteValueMIPS_OpCvt32Fto64F_0(v)
	case OpCvt32to32F:
		return rewriteValueMIPS_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValueMIPS_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValueMIPS_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValueMIPS_OpCvt64Fto32F_0(v)
	case OpDiv16:
		return rewriteValueMIPS_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValueMIPS_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValueMIPS_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValueMIPS_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValueMIPS_OpDiv32u_0(v)
	case OpDiv64F:
		return rewriteValueMIPS_OpDiv64F_0(v)
	case OpDiv8:
		return rewriteValueMIPS_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValueMIPS_OpDiv8u_0(v)
	case OpEq16:
		return rewriteValueMIPS_OpEq16_0(v)
	case OpEq32:
		return rewriteValueMIPS_OpEq32_0(v)
	case OpEq32F:
		return psess.rewriteValueMIPS_OpEq32F_0(v)
	case OpEq64F:
		return psess.rewriteValueMIPS_OpEq64F_0(v)
	case OpEq8:
		return rewriteValueMIPS_OpEq8_0(v)
	case OpEqB:
		return rewriteValueMIPS_OpEqB_0(v)
	case OpEqPtr:
		return rewriteValueMIPS_OpEqPtr_0(v)
	case OpGeq16:
		return rewriteValueMIPS_OpGeq16_0(v)
	case OpGeq16U:
		return rewriteValueMIPS_OpGeq16U_0(v)
	case OpGeq32:
		return rewriteValueMIPS_OpGeq32_0(v)
	case OpGeq32F:
		return psess.rewriteValueMIPS_OpGeq32F_0(v)
	case OpGeq32U:
		return rewriteValueMIPS_OpGeq32U_0(v)
	case OpGeq64F:
		return psess.rewriteValueMIPS_OpGeq64F_0(v)
	case OpGeq8:
		return rewriteValueMIPS_OpGeq8_0(v)
	case OpGeq8U:
		return rewriteValueMIPS_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValueMIPS_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValueMIPS_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValueMIPS_OpGetClosurePtr_0(v)
	case OpGreater16:
		return rewriteValueMIPS_OpGreater16_0(v)
	case OpGreater16U:
		return rewriteValueMIPS_OpGreater16U_0(v)
	case OpGreater32:
		return rewriteValueMIPS_OpGreater32_0(v)
	case OpGreater32F:
		return psess.rewriteValueMIPS_OpGreater32F_0(v)
	case OpGreater32U:
		return rewriteValueMIPS_OpGreater32U_0(v)
	case OpGreater64F:
		return psess.rewriteValueMIPS_OpGreater64F_0(v)
	case OpGreater8:
		return rewriteValueMIPS_OpGreater8_0(v)
	case OpGreater8U:
		return rewriteValueMIPS_OpGreater8U_0(v)
	case OpHmul32:
		return rewriteValueMIPS_OpHmul32_0(v)
	case OpHmul32u:
		return rewriteValueMIPS_OpHmul32u_0(v)
	case OpInterCall:
		return rewriteValueMIPS_OpInterCall_0(v)
	case OpIsInBounds:
		return rewriteValueMIPS_OpIsInBounds_0(v)
	case OpIsNonNil:
		return rewriteValueMIPS_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return rewriteValueMIPS_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return rewriteValueMIPS_OpLeq16_0(v)
	case OpLeq16U:
		return rewriteValueMIPS_OpLeq16U_0(v)
	case OpLeq32:
		return rewriteValueMIPS_OpLeq32_0(v)
	case OpLeq32F:
		return psess.rewriteValueMIPS_OpLeq32F_0(v)
	case OpLeq32U:
		return rewriteValueMIPS_OpLeq32U_0(v)
	case OpLeq64F:
		return psess.rewriteValueMIPS_OpLeq64F_0(v)
	case OpLeq8:
		return rewriteValueMIPS_OpLeq8_0(v)
	case OpLeq8U:
		return rewriteValueMIPS_OpLeq8U_0(v)
	case OpLess16:
		return rewriteValueMIPS_OpLess16_0(v)
	case OpLess16U:
		return rewriteValueMIPS_OpLess16U_0(v)
	case OpLess32:
		return rewriteValueMIPS_OpLess32_0(v)
	case OpLess32F:
		return psess.rewriteValueMIPS_OpLess32F_0(v)
	case OpLess32U:
		return rewriteValueMIPS_OpLess32U_0(v)
	case OpLess64F:
		return psess.rewriteValueMIPS_OpLess64F_0(v)
	case OpLess8:
		return rewriteValueMIPS_OpLess8_0(v)
	case OpLess8U:
		return rewriteValueMIPS_OpLess8U_0(v)
	case OpLoad:
		return psess.rewriteValueMIPS_OpLoad_0(v)
	case OpLsh16x16:
		return rewriteValueMIPS_OpLsh16x16_0(v)
	case OpLsh16x32:
		return rewriteValueMIPS_OpLsh16x32_0(v)
	case OpLsh16x64:
		return rewriteValueMIPS_OpLsh16x64_0(v)
	case OpLsh16x8:
		return rewriteValueMIPS_OpLsh16x8_0(v)
	case OpLsh32x16:
		return rewriteValueMIPS_OpLsh32x16_0(v)
	case OpLsh32x32:
		return rewriteValueMIPS_OpLsh32x32_0(v)
	case OpLsh32x64:
		return rewriteValueMIPS_OpLsh32x64_0(v)
	case OpLsh32x8:
		return rewriteValueMIPS_OpLsh32x8_0(v)
	case OpLsh8x16:
		return rewriteValueMIPS_OpLsh8x16_0(v)
	case OpLsh8x32:
		return rewriteValueMIPS_OpLsh8x32_0(v)
	case OpLsh8x64:
		return rewriteValueMIPS_OpLsh8x64_0(v)
	case OpLsh8x8:
		return rewriteValueMIPS_OpLsh8x8_0(v)
	case OpMIPSADD:
		return rewriteValueMIPS_OpMIPSADD_0(v)
	case OpMIPSADDconst:
		return rewriteValueMIPS_OpMIPSADDconst_0(v)
	case OpMIPSAND:
		return rewriteValueMIPS_OpMIPSAND_0(v)
	case OpMIPSANDconst:
		return rewriteValueMIPS_OpMIPSANDconst_0(v)
	case OpMIPSCMOVZ:
		return rewriteValueMIPS_OpMIPSCMOVZ_0(v)
	case OpMIPSCMOVZzero:
		return rewriteValueMIPS_OpMIPSCMOVZzero_0(v)
	case OpMIPSLoweredAtomicAdd:
		return rewriteValueMIPS_OpMIPSLoweredAtomicAdd_0(v)
	case OpMIPSLoweredAtomicStore:
		return rewriteValueMIPS_OpMIPSLoweredAtomicStore_0(v)
	case OpMIPSMOVBUload:
		return rewriteValueMIPS_OpMIPSMOVBUload_0(v)
	case OpMIPSMOVBUreg:
		return rewriteValueMIPS_OpMIPSMOVBUreg_0(v)
	case OpMIPSMOVBload:
		return rewriteValueMIPS_OpMIPSMOVBload_0(v)
	case OpMIPSMOVBreg:
		return rewriteValueMIPS_OpMIPSMOVBreg_0(v)
	case OpMIPSMOVBstore:
		return rewriteValueMIPS_OpMIPSMOVBstore_0(v)
	case OpMIPSMOVBstorezero:
		return rewriteValueMIPS_OpMIPSMOVBstorezero_0(v)
	case OpMIPSMOVDload:
		return rewriteValueMIPS_OpMIPSMOVDload_0(v)
	case OpMIPSMOVDstore:
		return rewriteValueMIPS_OpMIPSMOVDstore_0(v)
	case OpMIPSMOVFload:
		return rewriteValueMIPS_OpMIPSMOVFload_0(v)
	case OpMIPSMOVFstore:
		return rewriteValueMIPS_OpMIPSMOVFstore_0(v)
	case OpMIPSMOVHUload:
		return rewriteValueMIPS_OpMIPSMOVHUload_0(v)
	case OpMIPSMOVHUreg:
		return rewriteValueMIPS_OpMIPSMOVHUreg_0(v)
	case OpMIPSMOVHload:
		return rewriteValueMIPS_OpMIPSMOVHload_0(v)
	case OpMIPSMOVHreg:
		return rewriteValueMIPS_OpMIPSMOVHreg_0(v)
	case OpMIPSMOVHstore:
		return rewriteValueMIPS_OpMIPSMOVHstore_0(v)
	case OpMIPSMOVHstorezero:
		return rewriteValueMIPS_OpMIPSMOVHstorezero_0(v)
	case OpMIPSMOVWload:
		return rewriteValueMIPS_OpMIPSMOVWload_0(v)
	case OpMIPSMOVWreg:
		return rewriteValueMIPS_OpMIPSMOVWreg_0(v)
	case OpMIPSMOVWstore:
		return rewriteValueMIPS_OpMIPSMOVWstore_0(v)
	case OpMIPSMOVWstorezero:
		return rewriteValueMIPS_OpMIPSMOVWstorezero_0(v)
	case OpMIPSMUL:
		return rewriteValueMIPS_OpMIPSMUL_0(v)
	case OpMIPSNEG:
		return rewriteValueMIPS_OpMIPSNEG_0(v)
	case OpMIPSNOR:
		return rewriteValueMIPS_OpMIPSNOR_0(v)
	case OpMIPSNORconst:
		return rewriteValueMIPS_OpMIPSNORconst_0(v)
	case OpMIPSOR:
		return rewriteValueMIPS_OpMIPSOR_0(v)
	case OpMIPSORconst:
		return rewriteValueMIPS_OpMIPSORconst_0(v)
	case OpMIPSSGT:
		return rewriteValueMIPS_OpMIPSSGT_0(v)
	case OpMIPSSGTU:
		return rewriteValueMIPS_OpMIPSSGTU_0(v)
	case OpMIPSSGTUconst:
		return rewriteValueMIPS_OpMIPSSGTUconst_0(v)
	case OpMIPSSGTUzero:
		return rewriteValueMIPS_OpMIPSSGTUzero_0(v)
	case OpMIPSSGTconst:
		return rewriteValueMIPS_OpMIPSSGTconst_0(v) || rewriteValueMIPS_OpMIPSSGTconst_10(v)
	case OpMIPSSGTzero:
		return rewriteValueMIPS_OpMIPSSGTzero_0(v)
	case OpMIPSSLL:
		return rewriteValueMIPS_OpMIPSSLL_0(v)
	case OpMIPSSLLconst:
		return rewriteValueMIPS_OpMIPSSLLconst_0(v)
	case OpMIPSSRA:
		return rewriteValueMIPS_OpMIPSSRA_0(v)
	case OpMIPSSRAconst:
		return rewriteValueMIPS_OpMIPSSRAconst_0(v)
	case OpMIPSSRL:
		return rewriteValueMIPS_OpMIPSSRL_0(v)
	case OpMIPSSRLconst:
		return rewriteValueMIPS_OpMIPSSRLconst_0(v)
	case OpMIPSSUB:
		return rewriteValueMIPS_OpMIPSSUB_0(v)
	case OpMIPSSUBconst:
		return rewriteValueMIPS_OpMIPSSUBconst_0(v)
	case OpMIPSXOR:
		return rewriteValueMIPS_OpMIPSXOR_0(v)
	case OpMIPSXORconst:
		return rewriteValueMIPS_OpMIPSXORconst_0(v)
	case OpMod16:
		return rewriteValueMIPS_OpMod16_0(v)
	case OpMod16u:
		return rewriteValueMIPS_OpMod16u_0(v)
	case OpMod32:
		return rewriteValueMIPS_OpMod32_0(v)
	case OpMod32u:
		return rewriteValueMIPS_OpMod32u_0(v)
	case OpMod8:
		return rewriteValueMIPS_OpMod8_0(v)
	case OpMod8u:
		return rewriteValueMIPS_OpMod8u_0(v)
	case OpMove:
		return psess.rewriteValueMIPS_OpMove_0(v) || psess.rewriteValueMIPS_OpMove_10(v)
	case OpMul16:
		return rewriteValueMIPS_OpMul16_0(v)
	case OpMul32:
		return rewriteValueMIPS_OpMul32_0(v)
	case OpMul32F:
		return rewriteValueMIPS_OpMul32F_0(v)
	case OpMul32uhilo:
		return rewriteValueMIPS_OpMul32uhilo_0(v)
	case OpMul64F:
		return rewriteValueMIPS_OpMul64F_0(v)
	case OpMul8:
		return rewriteValueMIPS_OpMul8_0(v)
	case OpNeg16:
		return rewriteValueMIPS_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValueMIPS_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValueMIPS_OpNeg32F_0(v)
	case OpNeg64F:
		return rewriteValueMIPS_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValueMIPS_OpNeg8_0(v)
	case OpNeq16:
		return rewriteValueMIPS_OpNeq16_0(v)
	case OpNeq32:
		return rewriteValueMIPS_OpNeq32_0(v)
	case OpNeq32F:
		return psess.rewriteValueMIPS_OpNeq32F_0(v)
	case OpNeq64F:
		return psess.rewriteValueMIPS_OpNeq64F_0(v)
	case OpNeq8:
		return rewriteValueMIPS_OpNeq8_0(v)
	case OpNeqB:
		return rewriteValueMIPS_OpNeqB_0(v)
	case OpNeqPtr:
		return rewriteValueMIPS_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValueMIPS_OpNilCheck_0(v)
	case OpNot:
		return rewriteValueMIPS_OpNot_0(v)
	case OpOffPtr:
		return rewriteValueMIPS_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValueMIPS_OpOr16_0(v)
	case OpOr32:
		return rewriteValueMIPS_OpOr32_0(v)
	case OpOr8:
		return rewriteValueMIPS_OpOr8_0(v)
	case OpOrB:
		return rewriteValueMIPS_OpOrB_0(v)
	case OpRound32F:
		return rewriteValueMIPS_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValueMIPS_OpRound64F_0(v)
	case OpRsh16Ux16:
		return rewriteValueMIPS_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return rewriteValueMIPS_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return rewriteValueMIPS_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return rewriteValueMIPS_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return rewriteValueMIPS_OpRsh16x16_0(v)
	case OpRsh16x32:
		return rewriteValueMIPS_OpRsh16x32_0(v)
	case OpRsh16x64:
		return rewriteValueMIPS_OpRsh16x64_0(v)
	case OpRsh16x8:
		return rewriteValueMIPS_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return rewriteValueMIPS_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return rewriteValueMIPS_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return rewriteValueMIPS_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return rewriteValueMIPS_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return rewriteValueMIPS_OpRsh32x16_0(v)
	case OpRsh32x32:
		return rewriteValueMIPS_OpRsh32x32_0(v)
	case OpRsh32x64:
		return rewriteValueMIPS_OpRsh32x64_0(v)
	case OpRsh32x8:
		return rewriteValueMIPS_OpRsh32x8_0(v)
	case OpRsh8Ux16:
		return rewriteValueMIPS_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return rewriteValueMIPS_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return rewriteValueMIPS_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return rewriteValueMIPS_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return rewriteValueMIPS_OpRsh8x16_0(v)
	case OpRsh8x32:
		return rewriteValueMIPS_OpRsh8x32_0(v)
	case OpRsh8x64:
		return rewriteValueMIPS_OpRsh8x64_0(v)
	case OpRsh8x8:
		return rewriteValueMIPS_OpRsh8x8_0(v)
	case OpSelect0:
		return psess.rewriteValueMIPS_OpSelect0_0(v) || rewriteValueMIPS_OpSelect0_10(v)
	case OpSelect1:
		return psess.rewriteValueMIPS_OpSelect1_0(v) || rewriteValueMIPS_OpSelect1_10(v)
	case OpSignExt16to32:
		return rewriteValueMIPS_OpSignExt16to32_0(v)
	case OpSignExt8to16:
		return rewriteValueMIPS_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValueMIPS_OpSignExt8to32_0(v)
	case OpSignmask:
		return rewriteValueMIPS_OpSignmask_0(v)
	case OpSlicemask:
		return rewriteValueMIPS_OpSlicemask_0(v)
	case OpSqrt:
		return rewriteValueMIPS_OpSqrt_0(v)
	case OpStaticCall:
		return rewriteValueMIPS_OpStaticCall_0(v)
	case OpStore:
		return psess.rewriteValueMIPS_OpStore_0(v)
	case OpSub16:
		return rewriteValueMIPS_OpSub16_0(v)
	case OpSub32:
		return rewriteValueMIPS_OpSub32_0(v)
	case OpSub32F:
		return rewriteValueMIPS_OpSub32F_0(v)
	case OpSub32withcarry:
		return rewriteValueMIPS_OpSub32withcarry_0(v)
	case OpSub64F:
		return rewriteValueMIPS_OpSub64F_0(v)
	case OpSub8:
		return rewriteValueMIPS_OpSub8_0(v)
	case OpSubPtr:
		return rewriteValueMIPS_OpSubPtr_0(v)
	case OpTrunc16to8:
		return rewriteValueMIPS_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValueMIPS_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValueMIPS_OpTrunc32to8_0(v)
	case OpWB:
		return rewriteValueMIPS_OpWB_0(v)
	case OpXor16:
		return rewriteValueMIPS_OpXor16_0(v)
	case OpXor32:
		return rewriteValueMIPS_OpXor32_0(v)
	case OpXor8:
		return rewriteValueMIPS_OpXor8_0(v)
	case OpZero:
		return psess.rewriteValueMIPS_OpZero_0(v) || psess.rewriteValueMIPS_OpZero_10(v)
	case OpZeroExt16to32:
		return rewriteValueMIPS_OpZeroExt16to32_0(v)
	case OpZeroExt8to16:
		return rewriteValueMIPS_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValueMIPS_OpZeroExt8to32_0(v)
	case OpZeromask:
		return rewriteValueMIPS_OpZeromask_0(v)
	}
	return false
}
func rewriteValueMIPS_OpAdd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAdd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAdd32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADDF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAdd32withcarry_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		c := v.Args[2]
		v.reset(OpMIPSADD)
		v.AddArg(c)
		v0 := b.NewValue0(v.Pos, OpMIPSADD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpAdd64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADDD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAdd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAddPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAddr_0(v *Value) bool {

	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpMIPSMOVWaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueMIPS_OpAnd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAnd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAnd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAndB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpAtomicAdd32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSLoweredAtomicAdd)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicAnd8_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(!config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicAnd)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = ^3
		v0.AddArg(v1)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSOR, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(val)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v5.AuxInt = 3
		v6 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v6.AuxInt = 3
		v6.AddArg(ptr)
		v5.AddArg(v6)
		v3.AddArg(v5)
		v2.AddArg(v3)
		v7 := b.NewValue0(v.Pos, OpMIPSNORconst, typ.UInt32)
		v7.AuxInt = 0
		v8 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v9 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v9.AuxInt = 0xff
		v8.AddArg(v9)
		v10 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v10.AuxInt = 3
		v11 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v11.AuxInt = 3
		v11.AddArg(ptr)
		v10.AddArg(v11)
		v8.AddArg(v10)
		v7.AddArg(v8)
		v2.AddArg(v7)
		v.AddArg(v2)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicAnd)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = ^3
		v0.AddArg(v1)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSOR, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(val)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v5.AuxInt = 3
		v6 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v6.AuxInt = 3
		v7 := b.NewValue0(v.Pos, OpMIPSXORconst, typ.UInt32)
		v7.AuxInt = 3
		v7.AddArg(ptr)
		v6.AddArg(v7)
		v5.AddArg(v6)
		v3.AddArg(v5)
		v2.AddArg(v3)
		v8 := b.NewValue0(v.Pos, OpMIPSNORconst, typ.UInt32)
		v8.AuxInt = 0
		v9 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v10 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v10.AuxInt = 0xff
		v9.AddArg(v10)
		v11 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v11.AuxInt = 3
		v12 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v12.AuxInt = 3
		v13 := b.NewValue0(v.Pos, OpMIPSXORconst, typ.UInt32)
		v13.AuxInt = 3
		v13.AddArg(ptr)
		v12.AddArg(v13)
		v11.AddArg(v12)
		v9.AddArg(v11)
		v8.AddArg(v9)
		v2.AddArg(v8)
		v.AddArg(v2)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpAtomicCompareAndSwap32_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpMIPSLoweredAtomicCas)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicExchange32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSLoweredAtomicExchange)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicLoad32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSLoweredAtomicLoad)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicLoadPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSLoweredAtomicLoad)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicOr8_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(!config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicOr)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = ^3
		v0.AddArg(v1)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v3.AddArg(val)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v4.AuxInt = 3
		v5 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v5.AuxInt = 3
		v5.AddArg(ptr)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v.AddArg(v2)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(config.BigEndian) {
			break
		}
		v.reset(OpMIPSLoweredAtomicOr)
		v0 := b.NewValue0(v.Pos, OpMIPSAND, typ.UInt32Ptr)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = ^3
		v0.AddArg(v1)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSSLL, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v3.AddArg(val)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v4.AuxInt = 3
		v5 := b.NewValue0(v.Pos, OpMIPSANDconst, typ.UInt32)
		v5.AuxInt = 3
		v6 := b.NewValue0(v.Pos, OpMIPSXORconst, typ.UInt32)
		v6.AuxInt = 3
		v6.AddArg(ptr)
		v5.AddArg(v6)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v.AddArg(v2)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpAtomicStore32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSLoweredAtomicStore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAtomicStorePtrNoWB_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSLoweredAtomicStore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpAvg32u_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSADD)
		v0 := b.NewValue0(v.Pos, OpMIPSSRLconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpMIPSSUB, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpBitLen32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpMIPSSUB)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 32
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCLZ, t)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpClosureCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSCALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpCom16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCom32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCom8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpConst16_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS_OpConst32_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS_OpConst32F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpMIPSMOVFconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS_OpConst64F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpMIPSMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS_OpConst8_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueMIPS_OpConstBool_0(v *Value) bool {

	for {
		b := v.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValueMIPS_OpConstNil_0(v *Value) bool {

	for {
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueMIPS_OpCtz32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpMIPSSUB)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 32
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCLZ, t)
		v2 := b.NewValue0(v.Pos, OpMIPSSUBconst, t)
		v2.AuxInt = 1
		v3 := b.NewValue0(v.Pos, OpMIPSAND, t)
		v3.AddArg(x)
		v4 := b.NewValue0(v.Pos, OpMIPSNEG, t)
		v4.AddArg(x)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpCtz32NonZero_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCtz32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCvt32Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSTRUNCFW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCvt32Fto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVFD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCvt32to32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVWF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCvt32to64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVWD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCvt64Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSTRUNCDW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpCvt64Fto32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVDF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpDiv16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
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
func rewriteValueMIPS_OpDiv16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
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
func rewriteValueMIPS_OpDiv32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSDIVF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpDiv32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpDiv64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSDIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpDiv8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
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
func rewriteValueMIPS_OpDiv8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
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
func rewriteValueMIPS_OpEq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
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
func rewriteValueMIPS_OpEq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpEq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpEq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
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
func rewriteValueMIPS_OpEqB_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpEqPtr_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(x)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(x)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, typ.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpGeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGEF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpGeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGED, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(x)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(x)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGetCallerPC_0(v *Value) bool {

	for {
		v.reset(OpMIPSLoweredGetCallerPC)
		return true
	}
}
func rewriteValueMIPS_OpGetCallerSP_0(v *Value) bool {

	for {
		v.reset(OpMIPSLoweredGetCallerSP)
		return true
	}
}
func rewriteValueMIPS_OpGetClosurePtr_0(v *Value) bool {

	for {
		v.reset(OpMIPSLoweredGetClosurePtr)
		return true
	}
}
func rewriteValueMIPS_OpGreater16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpGreater16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpGreater32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpGreater32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGreater32U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpGreater64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpGreater8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpGreater8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpHmul32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSMULT, types.NewTuple(typ.Int32, typ.Int32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpHmul32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSMULTU, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpInterCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[1]
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSCALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpIsInBounds_0(v *Value) bool {

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpMIPSSGTU)
		v.AddArg(len)
		v.AddArg(idx)
		return true
	}
}
func rewriteValueMIPS_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		ptr := v.Args[0]
		v.reset(OpMIPSSGTU)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpIsSliceInBounds_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, typ.Bool)
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
func rewriteValueMIPS_OpLeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
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
func rewriteValueMIPS_OpLeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpLeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGEF, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpLeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGED, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGT, typ.Bool)
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
func rewriteValueMIPS_OpLeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
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
func rewriteValueMIPS_OpLess16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpLess16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpLess32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpLess32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTF, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLess32U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpLess64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPGTD, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpLess8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpLess8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpLoad_0(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpMIPSMOVBUload)
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
		v.reset(OpMIPSMOVBload)
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
		v.reset(OpMIPSMOVBUload)
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
		v.reset(OpMIPSMOVHload)
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
		v.reset(OpMIPSMOVHUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is32BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpMIPSMOVWload)
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
		v.reset(OpMIPSMOVFload)
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
		v.reset(OpMIPSMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueMIPS_OpLsh16x64_0(v *Value) bool {

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
		v.reset(OpMIPSSLLconst)
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
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueMIPS_OpLsh32x64_0(v *Value) bool {

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
		v.reset(OpMIPSSLLconst)
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
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueMIPS_OpLsh8x64_0(v *Value) bool {

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
		v.reset(OpMIPSSLLconst)
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
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpMIPSADD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSNEG {
			break
		}
		y := v_1.Args[0]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSNEG {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSADDconst_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpMIPSMOVWaddr)
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
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c + d))
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSAND_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSANDconst)
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

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSGTUconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSSGTUconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSOR, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSGTUconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		y := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSSGTUconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		x := v_1.Args[0]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpMIPSOR, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSANDconst_0(v *Value) bool {

	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
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
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c & d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSCMOVZ_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		b := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpMIPSMOVWconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		v.reset(OpCopy)
		v.Type = b.Type
		v.AddArg(b)
		return true
	}

	for {
		_ = v.Args[2]
		a := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpMIPSMOVWconst {
			break
		}
		c := v_2.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		c := v.Args[2]
		v.reset(OpMIPSCMOVZzero)
		v.AddArg(a)
		v.AddArg(c)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSCMOVZzero_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c != 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSLoweredAtomicAdd_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(is16Bit(c)) {
			break
		}
		v.reset(OpMIPSLoweredAtomicAddconst)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSLoweredAtomicStore_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPSLoweredAtomicStorezero)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBUload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBUload)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpMIPSMOVBUreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBUreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpMIPSMOVBload {
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
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0xff
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBload)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpMIPSMOVBreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUload {
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
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x80 == 0) {
			break
		}
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0x7f
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBstore)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBstore)
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
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstorezero)
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
		if v_1.Op != OpMIPSMOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
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
		if v_1.Op != OpMIPSMOVBUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
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
		if v_1.Op != OpMIPSMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
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
		if v_1.Op != OpMIPSMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
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
		if v_1.Op != OpMIPSMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVBstorezero_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVBstorezero)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVDload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVDload)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVDstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVDstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVDstore)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVFload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVFload)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVFload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVFstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVFstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVFstore)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHUload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHUload)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVHstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpMIPSMOVHUreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHUreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVHUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpMIPSMOVHload {
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
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHUload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0xffff
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHload)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVHstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpMIPSMOVHreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVBUreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpMIPSMOVHreg {
			break
		}
		v.reset(OpMIPSMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpMIPSMOVHUload {
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
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x8000 == 0) {
			break
		}
		v.reset(OpMIPSANDconst)
		v.AuxInt = c & 0x7fff
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHstore)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHstore)
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
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPSMOVHstorezero)
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
		if v_1.Op != OpMIPSMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVHstore)
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
		if v_1.Op != OpMIPSMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVHstore)
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
		if v_1.Op != OpMIPSMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVHstorezero_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVHstorezero)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWload)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWnop)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWstore)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVWstore)
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
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpMIPSMOVWstorezero)
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
		if v_1.Op != OpMIPSMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMOVWstorezero_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		if x.Op != OpMIPSADDconst {
			break
		}
		off2 := x.AuxInt
		ptr := x.Args[0]
		mem := v.Args[1]
		if !(is16Bit(off1+off2) || x.Uses == 1) {
			break
		}
		v.reset(OpMIPSMOVWstorezero)
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
		if v_0.Op != OpMIPSMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpMIPSMOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSMUL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
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
		if v_1.Op != OpMIPSMOVWconst {
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
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		x := v.Args[1]
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c) * int32(d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c) * int32(d))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSNEG_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(-c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSNOR_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSNORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSNORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSNORconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = ^(c | d)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSOR_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSORconst)
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

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSGTUzero {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSSGTUzero {
			break
		}
		y := v_1.Args[0]
		v.reset(OpMIPSSGTUzero)
		v0 := b.NewValue0(v.Pos, OpMIPSOR, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSGTUzero {
			break
		}
		y := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSSGTUzero {
			break
		}
		x := v_1.Args[0]
		v.reset(OpMIPSSGTUzero)
		v0 := b.NewValue0(v.Pos, OpMIPSOR, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSORconst_0(v *Value) bool {

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
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = -1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c | d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGT_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSSGTconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSSGTzero)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTU_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSSGTUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSSGTUzero)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTUconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(c) > uint32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(c) <= uint32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVBUreg {
			break
		}
		if !(0xff < uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVHUreg {
			break
		}
		if !(0xffff < uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		m := v_0.AuxInt
		if !(uint32(m) < uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSRLconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(d) <= 31 && 1<<(32-uint32(d)) <= uint32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTUzero_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(d) != 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(uint32(d) == 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(c) > int32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(c) <= int32(d)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVBreg {
			break
		}
		if !(0x7f < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVBreg {
			break
		}
		if !(int32(c) <= -0x80) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVBUreg {
			break
		}
		if !(0xff < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVBUreg {
			break
		}
		if !(int32(c) < 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVHreg {
			break
		}
		if !(0x7fff < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVHreg {
			break
		}
		if !(int32(c) <= -0x8000) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVHUreg {
			break
		}
		if !(0xffff < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVHUreg {
			break
		}
		if !(int32(c) < 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTconst_10(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSANDconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int32(m) && int32(m) < int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSRLconst {
			break
		}
		d := v_0.AuxInt
		if !(0 <= int32(c) && uint32(d) <= 31 && 1<<(32-uint32(d)) <= int32(c)) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSGTzero_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(d) > 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		if !(int32(d) <= 0) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSLL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSLLconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(d) << uint32(c)))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRA_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSRAconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRAconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(d) >> uint32(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSRLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSRLconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(uint32(d) >> uint32(c))
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSUB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSSUBconst)
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
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		x := v.Args[1]
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSSUBconst_0(v *Value) bool {

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
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(d - c))
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(-c - d))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = int64(int32(-c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSXOR_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPSXORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpMIPSXORconst)
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
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpMIPSXORconst_0(v *Value) bool {

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
		v.reset(OpMIPSNORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = c ^ d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSXORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPSXORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMod16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
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
func rewriteValueMIPS_OpMod16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
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
func rewriteValueMIPS_OpMod32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpMod8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIV, types.NewTuple(typ.Int32, typ.Int32))
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
func rewriteValueMIPS_OpMod8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPSDIVU, types.NewTuple(typ.UInt32, typ.UInt32))
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
func (psess *PackageSession) rewriteValueMIPS_OpMove_0(v *Value) bool {
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
		v.reset(OpMIPSMOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
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
		v.reset(OpMIPSMOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHUload, typ.UInt16)
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
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 1
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v0.AuxInt = 1
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
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
		v.reset(OpMIPSMOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
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
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHUload, typ.UInt16)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVHUload, typ.UInt16)
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
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v4.AuxInt = 1
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
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
		if v.AuxInt != 3 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v2.AuxInt = 1
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVBUload, typ.UInt8)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
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
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
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
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, psess.types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVHstore, psess.types.TypeMem)
		v3.AuxInt = 2
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v4.AuxInt = 2
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVHstore, psess.types.TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
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
func (psess *PackageSession) rewriteValueMIPS_OpMove_10(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

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
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, psess.types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVHstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVHload, typ.Int16)
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
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
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
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 12
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v0.AuxInt = 12
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v2.AuxInt = 8
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v3.AuxInt = 4
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
		v4.AuxInt = 4
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVWload, typ.UInt32)
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
		s := v.AuxInt
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 16 || t.(*types.Type).Alignment(psess.types)%4 != 0) {
			break
		}
		v.reset(OpMIPSLoweredMove)
		v.AuxInt = t.(*types.Type).Alignment(psess.types)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpMIPSADDconst, src.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(psess.types), config)
		v0.AddArg(src)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpMul16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpMul32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpMul32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpMul32uhilo_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMULTU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpMul64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpMul8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpNeg16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpNeg32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpNeg32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSNEGF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpNeg64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSNEGD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpNeg8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpNeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpNeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpNeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpNeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSFPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPSCMPEQD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpNeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpNeqB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpNeqPtr_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSGTU)
		v0 := b.NewValue0(v.Pos, OpMIPSXOR, typ.UInt32)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpNilCheck_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSLoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpNot_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSXORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpOffPtr_0(v *Value) bool {

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpMIPSMOVWaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpMIPSADDconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueMIPS_OpOr16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpOr32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpOr8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpOrB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpRound32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpRound64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh16Ux64_0(v *Value) bool {
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
		v.reset(OpMIPSSRLconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 16
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
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -1
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = -1
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh16x64_0(v *Value) bool {
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
		v.reset(OpMIPSSRAconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 16
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
		if !(uint32(c) >= 16) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -1
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueMIPS_OpRsh32Ux64_0(v *Value) bool {

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
		v.reset(OpMIPSSRLconst)
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
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = -1
		v0.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = -1
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v2.AuxInt = 32
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpRsh32x64_0(v *Value) bool {

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
		v.reset(OpMIPSSRAconst)
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
		if !(uint32(c) >= 32) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = -1
		v0.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v4 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS_OpRsh8Ux64_0(v *Value) bool {
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
		v.reset(OpMIPSSRLconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 24
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
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSSRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -1
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = -1
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v3.AuxInt = 32
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS_OpRsh8x64_0(v *Value) bool {
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
		v.reset(OpMIPSSRAconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 24
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
		if !(uint32(c) >= 8) {
			break
		}
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpMIPSSLLconst, typ.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueMIPS_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSCMOVZ, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v3.AuxInt = -1
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPSSGTUconst, typ.Bool)
		v4.AuxInt = 32
		v5 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpSelect0_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32carry {
			break
		}
		t := v_0.Type
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpMIPSADD)
		v.Type = t.FieldType(psess.types, 0)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSub32carry {
			break
		}
		t := v_0.Type
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpMIPSSUB)
		v.Type = t.FieldType(psess.types, 0)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_0.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_0.AuxInt != 1 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_0.AuxInt != -1 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSADDconst, x.Type)
		v0.AuxInt = -1
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_1.AuxInt != -1 {
			break
		}
		v.reset(OpMIPSCMOVZ)
		v0 := b.NewValue0(v.Pos, OpMIPSADDconst, x.Type)
		v0.AuxInt = -1
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSRLconst)
		v.AuxInt = 32 - log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSRLconst)
		v.AuxInt = 32 - log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpSelect0_10(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = (c * d) >> 32
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = (c * d) >> 32
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSDIV {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c) % int32(d))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSDIVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(c) % uint32(d)))
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueMIPS_OpSelect1_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpAdd32carry {
			break
		}
		t := v_0.Type
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpMIPSSGTU)
		v.Type = typ.Bool
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPSADD, t.FieldType(psess.types, 0))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSub32carry {
			break
		}
		t := v_0.Type
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpMIPSSGTU)
		v.Type = typ.Bool
		v0 := b.NewValue0(v.Pos, OpMIPSSUB, t.FieldType(psess.types, 0))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_0.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_1.AuxInt != 0 {
			break
		}
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
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
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
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
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_0.AuxInt != -1 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpMIPSNEG)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		if v_0_1.AuxInt != -1 {
			break
		}
		v.reset(OpMIPSNEG)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		x := v_0.Args[1]
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(int64(uint32(c)))) {
			break
		}
		v.reset(OpMIPSSLLconst)
		v.AuxInt = log2(int64(uint32(c)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS_OpSelect1_10(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(c) * uint32(d)))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSMULTU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(c) * uint32(d)))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSDIV {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(c) / int32(d))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpMIPSDIVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPSMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPSMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPSMOVWconst)
		v.AuxInt = int64(int32(uint32(c) / uint32(d)))
		return true
	}
	return false
}
func rewriteValueMIPS_OpSignExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpSignExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpSignExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpSignmask_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpSlicemask_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpMIPSSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpMIPSNEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS_OpSqrt_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSSQRTD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpStaticCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpMIPSCALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpStore_0(v *Value) bool {

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 1) {
			break
		}
		v.reset(OpMIPSMOVBstore)
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
		v.reset(OpMIPSMOVHstore)
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
		v.reset(OpMIPSMOVWstore)
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
		v.reset(OpMIPSMOVFstore)
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
		v.reset(OpMIPSMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpSub16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpSub32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpSub32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUBF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpSub32withcarry_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		c := v.Args[2]
		v.reset(OpMIPSSUB)
		v0 := b.NewValue0(v.Pos, OpMIPSSUB, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(c)
		return true
	}
}
func rewriteValueMIPS_OpSub64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUBD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpSub8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpSubPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpTrunc16to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpTrunc32to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpTrunc32to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpWB_0(v *Value) bool {

	for {
		fn := v.Aux
		_ = v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpMIPSLoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueMIPS_OpXor16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpXor32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS_OpXor8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpMIPSXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueMIPS_OpZero_0(v *Value) bool {
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
		v.reset(OpMIPSMOVBstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
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
		v.reset(OpMIPSMOVHstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
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
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 1
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
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
		v.reset(OpMIPSMOVWstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
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
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
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
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 3
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 3 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpMIPSMOVBstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVBstore, psess.types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
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
		v.reset(OpMIPSMOVHstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVHstore, psess.types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVHstore, psess.types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
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
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueMIPS_OpZero_10(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

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
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 8
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
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
		if !(t.(*types.Type).Alignment(psess.types)%4 == 0) {
			break
		}
		v.reset(OpMIPSMOVWstore)
		v.AuxInt = 12
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v3.AuxInt = 4
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPSMOVWstore, psess.types.TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
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
		if !(s > 16 || t.(*types.Type).Alignment(psess.types)%4 != 0) {
			break
		}
		v.reset(OpMIPSLoweredZero)
		v.AuxInt = t.(*types.Type).Alignment(psess.types)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPSADDconst, ptr.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(psess.types), config)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS_OpZeroExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVHUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpZeroExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpZeroExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpMIPSMOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS_OpZeromask_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpMIPSNEG)
		v0 := b.NewValue0(v.Pos, OpMIPSSGTU, typ.Bool)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPSMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteBlockMIPS(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	case BlockMIPSEQ:

		for {
			v := b.Control
			if v.Op != OpMIPSFPFlagTrue {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockMIPSFPF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSFPFlagFalse {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockMIPSFPT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGT {
				break
			}
			_ = cmp.Args[1]
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTU {
				break
			}
			_ = cmp.Args[1]
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTconst {
				break
			}
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTUconst {
				break
			}
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTzero {
				break
			}
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTUzero {
				break
			}
			b.Kind = BlockMIPSNE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSSGTUconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPSNE
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSSGTUzero {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPSEQ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSSGTconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPSGEZ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSSGTzero {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPSLEZ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
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
			if v.Op != OpMIPSMOVWconst {
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
	case BlockMIPSGEZ:

		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) >= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) < 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case BlockMIPSGTZ:

		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) > 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) <= 0) {
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
			b.Kind = BlockMIPSNE
			b.SetControl(cond)
			b.Aux = nil
			return true
		}
	case BlockMIPSLEZ:

		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) <= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) > 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case BlockMIPSLTZ:

		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) < 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) >= 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case BlockMIPSNE:

		for {
			v := b.Control
			if v.Op != OpMIPSFPFlagTrue {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockMIPSFPT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSFPFlagFalse {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockMIPSFPF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGT {
				break
			}
			_ = cmp.Args[1]
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTU {
				break
			}
			_ = cmp.Args[1]
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTconst {
				break
			}
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTUconst {
				break
			}
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTzero {
				break
			}
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSXORconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			cmp := v.Args[0]
			if cmp.Op != OpMIPSSGTUzero {
				break
			}
			b.Kind = BlockMIPSEQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSSGTUconst {
				break
			}
			if v.AuxInt != 1 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPSEQ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSSGTUzero {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPSNE
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSSGTconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPSLTZ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSSGTzero {
				break
			}
			x := v.Args[0]
			b.Kind = BlockMIPSGTZ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpMIPSMOVWconst {
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
			if v.Op != OpMIPSMOVWconst {
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

package ssa

import "github.com/dave/golib/src/cmd/compile/internal/types"

// in case not otherwise used
// in case not otherwise used
// in case not otherwise used
// in case not otherwise used

func (psess *PackageSession) rewriteValueS390X(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValueS390X_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValueS390X_OpAdd32_0(v)
	case OpAdd32F:
		return rewriteValueS390X_OpAdd32F_0(v)
	case OpAdd64:
		return rewriteValueS390X_OpAdd64_0(v)
	case OpAdd64F:
		return rewriteValueS390X_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValueS390X_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValueS390X_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValueS390X_OpAddr_0(v)
	case OpAnd16:
		return rewriteValueS390X_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValueS390X_OpAnd32_0(v)
	case OpAnd64:
		return rewriteValueS390X_OpAnd64_0(v)
	case OpAnd8:
		return rewriteValueS390X_OpAnd8_0(v)
	case OpAndB:
		return rewriteValueS390X_OpAndB_0(v)
	case OpAtomicAdd32:
		return psess.rewriteValueS390X_OpAtomicAdd32_0(v)
	case OpAtomicAdd64:
		return psess.rewriteValueS390X_OpAtomicAdd64_0(v)
	case OpAtomicCompareAndSwap32:
		return rewriteValueS390X_OpAtomicCompareAndSwap32_0(v)
	case OpAtomicCompareAndSwap64:
		return rewriteValueS390X_OpAtomicCompareAndSwap64_0(v)
	case OpAtomicExchange32:
		return rewriteValueS390X_OpAtomicExchange32_0(v)
	case OpAtomicExchange64:
		return rewriteValueS390X_OpAtomicExchange64_0(v)
	case OpAtomicLoad32:
		return rewriteValueS390X_OpAtomicLoad32_0(v)
	case OpAtomicLoad64:
		return rewriteValueS390X_OpAtomicLoad64_0(v)
	case OpAtomicLoadPtr:
		return rewriteValueS390X_OpAtomicLoadPtr_0(v)
	case OpAtomicStore32:
		return rewriteValueS390X_OpAtomicStore32_0(v)
	case OpAtomicStore64:
		return rewriteValueS390X_OpAtomicStore64_0(v)
	case OpAtomicStorePtrNoWB:
		return rewriteValueS390X_OpAtomicStorePtrNoWB_0(v)
	case OpAvg64u:
		return rewriteValueS390X_OpAvg64u_0(v)
	case OpBitLen64:
		return rewriteValueS390X_OpBitLen64_0(v)
	case OpBswap32:
		return rewriteValueS390X_OpBswap32_0(v)
	case OpBswap64:
		return rewriteValueS390X_OpBswap64_0(v)
	case OpCeil:
		return rewriteValueS390X_OpCeil_0(v)
	case OpClosureCall:
		return rewriteValueS390X_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValueS390X_OpCom16_0(v)
	case OpCom32:
		return rewriteValueS390X_OpCom32_0(v)
	case OpCom64:
		return rewriteValueS390X_OpCom64_0(v)
	case OpCom8:
		return rewriteValueS390X_OpCom8_0(v)
	case OpConst16:
		return rewriteValueS390X_OpConst16_0(v)
	case OpConst32:
		return rewriteValueS390X_OpConst32_0(v)
	case OpConst32F:
		return rewriteValueS390X_OpConst32F_0(v)
	case OpConst64:
		return rewriteValueS390X_OpConst64_0(v)
	case OpConst64F:
		return rewriteValueS390X_OpConst64F_0(v)
	case OpConst8:
		return rewriteValueS390X_OpConst8_0(v)
	case OpConstBool:
		return rewriteValueS390X_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValueS390X_OpConstNil_0(v)
	case OpCtz32:
		return rewriteValueS390X_OpCtz32_0(v)
	case OpCtz32NonZero:
		return rewriteValueS390X_OpCtz32NonZero_0(v)
	case OpCtz64:
		return rewriteValueS390X_OpCtz64_0(v)
	case OpCtz64NonZero:
		return rewriteValueS390X_OpCtz64NonZero_0(v)
	case OpCvt32Fto32:
		return rewriteValueS390X_OpCvt32Fto32_0(v)
	case OpCvt32Fto64:
		return rewriteValueS390X_OpCvt32Fto64_0(v)
	case OpCvt32Fto64F:
		return rewriteValueS390X_OpCvt32Fto64F_0(v)
	case OpCvt32to32F:
		return rewriteValueS390X_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValueS390X_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValueS390X_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValueS390X_OpCvt64Fto32F_0(v)
	case OpCvt64Fto64:
		return rewriteValueS390X_OpCvt64Fto64_0(v)
	case OpCvt64to32F:
		return rewriteValueS390X_OpCvt64to32F_0(v)
	case OpCvt64to64F:
		return rewriteValueS390X_OpCvt64to64F_0(v)
	case OpDiv16:
		return rewriteValueS390X_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValueS390X_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValueS390X_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValueS390X_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValueS390X_OpDiv32u_0(v)
	case OpDiv64:
		return rewriteValueS390X_OpDiv64_0(v)
	case OpDiv64F:
		return rewriteValueS390X_OpDiv64F_0(v)
	case OpDiv64u:
		return rewriteValueS390X_OpDiv64u_0(v)
	case OpDiv8:
		return rewriteValueS390X_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValueS390X_OpDiv8u_0(v)
	case OpEq16:
		return psess.rewriteValueS390X_OpEq16_0(v)
	case OpEq32:
		return psess.rewriteValueS390X_OpEq32_0(v)
	case OpEq32F:
		return psess.rewriteValueS390X_OpEq32F_0(v)
	case OpEq64:
		return psess.rewriteValueS390X_OpEq64_0(v)
	case OpEq64F:
		return psess.rewriteValueS390X_OpEq64F_0(v)
	case OpEq8:
		return psess.rewriteValueS390X_OpEq8_0(v)
	case OpEqB:
		return psess.rewriteValueS390X_OpEqB_0(v)
	case OpEqPtr:
		return psess.rewriteValueS390X_OpEqPtr_0(v)
	case OpFloor:
		return rewriteValueS390X_OpFloor_0(v)
	case OpGeq16:
		return psess.rewriteValueS390X_OpGeq16_0(v)
	case OpGeq16U:
		return psess.rewriteValueS390X_OpGeq16U_0(v)
	case OpGeq32:
		return psess.rewriteValueS390X_OpGeq32_0(v)
	case OpGeq32F:
		return psess.rewriteValueS390X_OpGeq32F_0(v)
	case OpGeq32U:
		return psess.rewriteValueS390X_OpGeq32U_0(v)
	case OpGeq64:
		return psess.rewriteValueS390X_OpGeq64_0(v)
	case OpGeq64F:
		return psess.rewriteValueS390X_OpGeq64F_0(v)
	case OpGeq64U:
		return psess.rewriteValueS390X_OpGeq64U_0(v)
	case OpGeq8:
		return psess.rewriteValueS390X_OpGeq8_0(v)
	case OpGeq8U:
		return psess.rewriteValueS390X_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValueS390X_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValueS390X_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValueS390X_OpGetClosurePtr_0(v)
	case OpGetG:
		return rewriteValueS390X_OpGetG_0(v)
	case OpGreater16:
		return psess.rewriteValueS390X_OpGreater16_0(v)
	case OpGreater16U:
		return psess.rewriteValueS390X_OpGreater16U_0(v)
	case OpGreater32:
		return psess.rewriteValueS390X_OpGreater32_0(v)
	case OpGreater32F:
		return psess.rewriteValueS390X_OpGreater32F_0(v)
	case OpGreater32U:
		return psess.rewriteValueS390X_OpGreater32U_0(v)
	case OpGreater64:
		return psess.rewriteValueS390X_OpGreater64_0(v)
	case OpGreater64F:
		return psess.rewriteValueS390X_OpGreater64F_0(v)
	case OpGreater64U:
		return psess.rewriteValueS390X_OpGreater64U_0(v)
	case OpGreater8:
		return psess.rewriteValueS390X_OpGreater8_0(v)
	case OpGreater8U:
		return psess.rewriteValueS390X_OpGreater8U_0(v)
	case OpHmul32:
		return rewriteValueS390X_OpHmul32_0(v)
	case OpHmul32u:
		return rewriteValueS390X_OpHmul32u_0(v)
	case OpHmul64:
		return rewriteValueS390X_OpHmul64_0(v)
	case OpHmul64u:
		return rewriteValueS390X_OpHmul64u_0(v)
	case OpITab:
		return rewriteValueS390X_OpITab_0(v)
	case OpInterCall:
		return rewriteValueS390X_OpInterCall_0(v)
	case OpIsInBounds:
		return psess.rewriteValueS390X_OpIsInBounds_0(v)
	case OpIsNonNil:
		return psess.rewriteValueS390X_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return psess.rewriteValueS390X_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return psess.rewriteValueS390X_OpLeq16_0(v)
	case OpLeq16U:
		return psess.rewriteValueS390X_OpLeq16U_0(v)
	case OpLeq32:
		return psess.rewriteValueS390X_OpLeq32_0(v)
	case OpLeq32F:
		return psess.rewriteValueS390X_OpLeq32F_0(v)
	case OpLeq32U:
		return psess.rewriteValueS390X_OpLeq32U_0(v)
	case OpLeq64:
		return psess.rewriteValueS390X_OpLeq64_0(v)
	case OpLeq64F:
		return psess.rewriteValueS390X_OpLeq64F_0(v)
	case OpLeq64U:
		return psess.rewriteValueS390X_OpLeq64U_0(v)
	case OpLeq8:
		return psess.rewriteValueS390X_OpLeq8_0(v)
	case OpLeq8U:
		return psess.rewriteValueS390X_OpLeq8U_0(v)
	case OpLess16:
		return psess.rewriteValueS390X_OpLess16_0(v)
	case OpLess16U:
		return psess.rewriteValueS390X_OpLess16U_0(v)
	case OpLess32:
		return psess.rewriteValueS390X_OpLess32_0(v)
	case OpLess32F:
		return psess.rewriteValueS390X_OpLess32F_0(v)
	case OpLess32U:
		return psess.rewriteValueS390X_OpLess32U_0(v)
	case OpLess64:
		return psess.rewriteValueS390X_OpLess64_0(v)
	case OpLess64F:
		return psess.rewriteValueS390X_OpLess64F_0(v)
	case OpLess64U:
		return psess.rewriteValueS390X_OpLess64U_0(v)
	case OpLess8:
		return psess.rewriteValueS390X_OpLess8_0(v)
	case OpLess8U:
		return psess.rewriteValueS390X_OpLess8U_0(v)
	case OpLoad:
		return psess.rewriteValueS390X_OpLoad_0(v)
	case OpLsh16x16:
		return psess.rewriteValueS390X_OpLsh16x16_0(v)
	case OpLsh16x32:
		return psess.rewriteValueS390X_OpLsh16x32_0(v)
	case OpLsh16x64:
		return psess.rewriteValueS390X_OpLsh16x64_0(v)
	case OpLsh16x8:
		return psess.rewriteValueS390X_OpLsh16x8_0(v)
	case OpLsh32x16:
		return psess.rewriteValueS390X_OpLsh32x16_0(v)
	case OpLsh32x32:
		return psess.rewriteValueS390X_OpLsh32x32_0(v)
	case OpLsh32x64:
		return psess.rewriteValueS390X_OpLsh32x64_0(v)
	case OpLsh32x8:
		return psess.rewriteValueS390X_OpLsh32x8_0(v)
	case OpLsh64x16:
		return psess.rewriteValueS390X_OpLsh64x16_0(v)
	case OpLsh64x32:
		return psess.rewriteValueS390X_OpLsh64x32_0(v)
	case OpLsh64x64:
		return psess.rewriteValueS390X_OpLsh64x64_0(v)
	case OpLsh64x8:
		return psess.rewriteValueS390X_OpLsh64x8_0(v)
	case OpLsh8x16:
		return psess.rewriteValueS390X_OpLsh8x16_0(v)
	case OpLsh8x32:
		return psess.rewriteValueS390X_OpLsh8x32_0(v)
	case OpLsh8x64:
		return psess.rewriteValueS390X_OpLsh8x64_0(v)
	case OpLsh8x8:
		return psess.rewriteValueS390X_OpLsh8x8_0(v)
	case OpMod16:
		return rewriteValueS390X_OpMod16_0(v)
	case OpMod16u:
		return rewriteValueS390X_OpMod16u_0(v)
	case OpMod32:
		return rewriteValueS390X_OpMod32_0(v)
	case OpMod32u:
		return rewriteValueS390X_OpMod32u_0(v)
	case OpMod64:
		return rewriteValueS390X_OpMod64_0(v)
	case OpMod64u:
		return rewriteValueS390X_OpMod64u_0(v)
	case OpMod8:
		return rewriteValueS390X_OpMod8_0(v)
	case OpMod8u:
		return rewriteValueS390X_OpMod8u_0(v)
	case OpMove:
		return psess.rewriteValueS390X_OpMove_0(v) || psess.rewriteValueS390X_OpMove_10(v)
	case OpMul16:
		return rewriteValueS390X_OpMul16_0(v)
	case OpMul32:
		return rewriteValueS390X_OpMul32_0(v)
	case OpMul32F:
		return rewriteValueS390X_OpMul32F_0(v)
	case OpMul64:
		return rewriteValueS390X_OpMul64_0(v)
	case OpMul64F:
		return rewriteValueS390X_OpMul64F_0(v)
	case OpMul8:
		return rewriteValueS390X_OpMul8_0(v)
	case OpNeg16:
		return rewriteValueS390X_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValueS390X_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValueS390X_OpNeg32F_0(v)
	case OpNeg64:
		return rewriteValueS390X_OpNeg64_0(v)
	case OpNeg64F:
		return rewriteValueS390X_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValueS390X_OpNeg8_0(v)
	case OpNeq16:
		return psess.rewriteValueS390X_OpNeq16_0(v)
	case OpNeq32:
		return psess.rewriteValueS390X_OpNeq32_0(v)
	case OpNeq32F:
		return psess.rewriteValueS390X_OpNeq32F_0(v)
	case OpNeq64:
		return psess.rewriteValueS390X_OpNeq64_0(v)
	case OpNeq64F:
		return psess.rewriteValueS390X_OpNeq64F_0(v)
	case OpNeq8:
		return psess.rewriteValueS390X_OpNeq8_0(v)
	case OpNeqB:
		return psess.rewriteValueS390X_OpNeqB_0(v)
	case OpNeqPtr:
		return psess.rewriteValueS390X_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValueS390X_OpNilCheck_0(v)
	case OpNot:
		return rewriteValueS390X_OpNot_0(v)
	case OpOffPtr:
		return rewriteValueS390X_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValueS390X_OpOr16_0(v)
	case OpOr32:
		return rewriteValueS390X_OpOr32_0(v)
	case OpOr64:
		return rewriteValueS390X_OpOr64_0(v)
	case OpOr8:
		return rewriteValueS390X_OpOr8_0(v)
	case OpOrB:
		return rewriteValueS390X_OpOrB_0(v)
	case OpRound:
		return rewriteValueS390X_OpRound_0(v)
	case OpRound32F:
		return rewriteValueS390X_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValueS390X_OpRound64F_0(v)
	case OpRoundToEven:
		return rewriteValueS390X_OpRoundToEven_0(v)
	case OpRsh16Ux16:
		return psess.rewriteValueS390X_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return psess.rewriteValueS390X_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return psess.rewriteValueS390X_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return psess.rewriteValueS390X_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return psess.rewriteValueS390X_OpRsh16x16_0(v)
	case OpRsh16x32:
		return psess.rewriteValueS390X_OpRsh16x32_0(v)
	case OpRsh16x64:
		return psess.rewriteValueS390X_OpRsh16x64_0(v)
	case OpRsh16x8:
		return psess.rewriteValueS390X_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return psess.rewriteValueS390X_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return psess.rewriteValueS390X_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return psess.rewriteValueS390X_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return psess.rewriteValueS390X_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return psess.rewriteValueS390X_OpRsh32x16_0(v)
	case OpRsh32x32:
		return psess.rewriteValueS390X_OpRsh32x32_0(v)
	case OpRsh32x64:
		return psess.rewriteValueS390X_OpRsh32x64_0(v)
	case OpRsh32x8:
		return psess.rewriteValueS390X_OpRsh32x8_0(v)
	case OpRsh64Ux16:
		return psess.rewriteValueS390X_OpRsh64Ux16_0(v)
	case OpRsh64Ux32:
		return psess.rewriteValueS390X_OpRsh64Ux32_0(v)
	case OpRsh64Ux64:
		return psess.rewriteValueS390X_OpRsh64Ux64_0(v)
	case OpRsh64Ux8:
		return psess.rewriteValueS390X_OpRsh64Ux8_0(v)
	case OpRsh64x16:
		return psess.rewriteValueS390X_OpRsh64x16_0(v)
	case OpRsh64x32:
		return psess.rewriteValueS390X_OpRsh64x32_0(v)
	case OpRsh64x64:
		return psess.rewriteValueS390X_OpRsh64x64_0(v)
	case OpRsh64x8:
		return psess.rewriteValueS390X_OpRsh64x8_0(v)
	case OpRsh8Ux16:
		return psess.rewriteValueS390X_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return psess.rewriteValueS390X_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return psess.rewriteValueS390X_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return psess.rewriteValueS390X_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return psess.rewriteValueS390X_OpRsh8x16_0(v)
	case OpRsh8x32:
		return psess.rewriteValueS390X_OpRsh8x32_0(v)
	case OpRsh8x64:
		return psess.rewriteValueS390X_OpRsh8x64_0(v)
	case OpRsh8x8:
		return psess.rewriteValueS390X_OpRsh8x8_0(v)
	case OpS390XADD:
		return psess.rewriteValueS390X_OpS390XADD_0(v) || psess.rewriteValueS390X_OpS390XADD_10(v)
	case OpS390XADDW:
		return psess.rewriteValueS390X_OpS390XADDW_0(v) || psess.rewriteValueS390X_OpS390XADDW_10(v)
	case OpS390XADDWconst:
		return rewriteValueS390X_OpS390XADDWconst_0(v)
	case OpS390XADDWload:
		return rewriteValueS390X_OpS390XADDWload_0(v)
	case OpS390XADDconst:
		return rewriteValueS390X_OpS390XADDconst_0(v)
	case OpS390XADDload:
		return rewriteValueS390X_OpS390XADDload_0(v)
	case OpS390XAND:
		return rewriteValueS390X_OpS390XAND_0(v) || psess.rewriteValueS390X_OpS390XAND_10(v)
	case OpS390XANDW:
		return psess.rewriteValueS390X_OpS390XANDW_0(v) || psess.rewriteValueS390X_OpS390XANDW_10(v)
	case OpS390XANDWconst:
		return rewriteValueS390X_OpS390XANDWconst_0(v)
	case OpS390XANDWload:
		return rewriteValueS390X_OpS390XANDWload_0(v)
	case OpS390XANDconst:
		return rewriteValueS390X_OpS390XANDconst_0(v)
	case OpS390XANDload:
		return rewriteValueS390X_OpS390XANDload_0(v)
	case OpS390XCMP:
		return psess.rewriteValueS390X_OpS390XCMP_0(v)
	case OpS390XCMPU:
		return psess.rewriteValueS390X_OpS390XCMPU_0(v)
	case OpS390XCMPUconst:
		return rewriteValueS390X_OpS390XCMPUconst_0(v) || rewriteValueS390X_OpS390XCMPUconst_10(v)
	case OpS390XCMPW:
		return psess.rewriteValueS390X_OpS390XCMPW_0(v)
	case OpS390XCMPWU:
		return psess.rewriteValueS390X_OpS390XCMPWU_0(v)
	case OpS390XCMPWUconst:
		return rewriteValueS390X_OpS390XCMPWUconst_0(v)
	case OpS390XCMPWconst:
		return rewriteValueS390X_OpS390XCMPWconst_0(v)
	case OpS390XCMPconst:
		return rewriteValueS390X_OpS390XCMPconst_0(v) || rewriteValueS390X_OpS390XCMPconst_10(v)
	case OpS390XCPSDR:
		return rewriteValueS390X_OpS390XCPSDR_0(v)
	case OpS390XFADD:
		return rewriteValueS390X_OpS390XFADD_0(v)
	case OpS390XFADDS:
		return rewriteValueS390X_OpS390XFADDS_0(v)
	case OpS390XFMOVDload:
		return rewriteValueS390X_OpS390XFMOVDload_0(v)
	case OpS390XFMOVDloadidx:
		return rewriteValueS390X_OpS390XFMOVDloadidx_0(v)
	case OpS390XFMOVDstore:
		return rewriteValueS390X_OpS390XFMOVDstore_0(v)
	case OpS390XFMOVDstoreidx:
		return rewriteValueS390X_OpS390XFMOVDstoreidx_0(v)
	case OpS390XFMOVSload:
		return rewriteValueS390X_OpS390XFMOVSload_0(v)
	case OpS390XFMOVSloadidx:
		return rewriteValueS390X_OpS390XFMOVSloadidx_0(v)
	case OpS390XFMOVSstore:
		return rewriteValueS390X_OpS390XFMOVSstore_0(v)
	case OpS390XFMOVSstoreidx:
		return rewriteValueS390X_OpS390XFMOVSstoreidx_0(v)
	case OpS390XFNEG:
		return rewriteValueS390X_OpS390XFNEG_0(v)
	case OpS390XFNEGS:
		return rewriteValueS390X_OpS390XFNEGS_0(v)
	case OpS390XFSUB:
		return rewriteValueS390X_OpS390XFSUB_0(v)
	case OpS390XFSUBS:
		return rewriteValueS390X_OpS390XFSUBS_0(v)
	case OpS390XLDGR:
		return rewriteValueS390X_OpS390XLDGR_0(v)
	case OpS390XLEDBR:
		return rewriteValueS390X_OpS390XLEDBR_0(v)
	case OpS390XLGDR:
		return rewriteValueS390X_OpS390XLGDR_0(v)
	case OpS390XLoweredRound32F:
		return rewriteValueS390X_OpS390XLoweredRound32F_0(v)
	case OpS390XLoweredRound64F:
		return rewriteValueS390X_OpS390XLoweredRound64F_0(v)
	case OpS390XMOVBZload:
		return rewriteValueS390X_OpS390XMOVBZload_0(v)
	case OpS390XMOVBZloadidx:
		return rewriteValueS390X_OpS390XMOVBZloadidx_0(v)
	case OpS390XMOVBZreg:
		return psess.rewriteValueS390X_OpS390XMOVBZreg_0(v) || rewriteValueS390X_OpS390XMOVBZreg_10(v)
	case OpS390XMOVBload:
		return rewriteValueS390X_OpS390XMOVBload_0(v)
	case OpS390XMOVBloadidx:
		return rewriteValueS390X_OpS390XMOVBloadidx_0(v)
	case OpS390XMOVBreg:
		return psess.rewriteValueS390X_OpS390XMOVBreg_0(v)
	case OpS390XMOVBstore:
		return rewriteValueS390X_OpS390XMOVBstore_0(v) || rewriteValueS390X_OpS390XMOVBstore_10(v)
	case OpS390XMOVBstoreconst:
		return rewriteValueS390X_OpS390XMOVBstoreconst_0(v)
	case OpS390XMOVBstoreidx:
		return rewriteValueS390X_OpS390XMOVBstoreidx_0(v) || rewriteValueS390X_OpS390XMOVBstoreidx_10(v) || rewriteValueS390X_OpS390XMOVBstoreidx_20(v) || rewriteValueS390X_OpS390XMOVBstoreidx_30(v)
	case OpS390XMOVDEQ:
		return rewriteValueS390X_OpS390XMOVDEQ_0(v)
	case OpS390XMOVDGE:
		return rewriteValueS390X_OpS390XMOVDGE_0(v)
	case OpS390XMOVDGT:
		return rewriteValueS390X_OpS390XMOVDGT_0(v)
	case OpS390XMOVDLE:
		return rewriteValueS390X_OpS390XMOVDLE_0(v)
	case OpS390XMOVDLT:
		return rewriteValueS390X_OpS390XMOVDLT_0(v)
	case OpS390XMOVDNE:
		return rewriteValueS390X_OpS390XMOVDNE_0(v)
	case OpS390XMOVDaddridx:
		return rewriteValueS390X_OpS390XMOVDaddridx_0(v)
	case OpS390XMOVDload:
		return psess.rewriteValueS390X_OpS390XMOVDload_0(v)
	case OpS390XMOVDloadidx:
		return rewriteValueS390X_OpS390XMOVDloadidx_0(v)
	case OpS390XMOVDnop:
		return psess.rewriteValueS390X_OpS390XMOVDnop_0(v) || rewriteValueS390X_OpS390XMOVDnop_10(v)
	case OpS390XMOVDreg:
		return psess.rewriteValueS390X_OpS390XMOVDreg_0(v) || rewriteValueS390X_OpS390XMOVDreg_10(v)
	case OpS390XMOVDstore:
		return psess.rewriteValueS390X_OpS390XMOVDstore_0(v)
	case OpS390XMOVDstoreconst:
		return rewriteValueS390X_OpS390XMOVDstoreconst_0(v)
	case OpS390XMOVDstoreidx:
		return rewriteValueS390X_OpS390XMOVDstoreidx_0(v)
	case OpS390XMOVHBRstore:
		return rewriteValueS390X_OpS390XMOVHBRstore_0(v)
	case OpS390XMOVHBRstoreidx:
		return rewriteValueS390X_OpS390XMOVHBRstoreidx_0(v) || rewriteValueS390X_OpS390XMOVHBRstoreidx_10(v)
	case OpS390XMOVHZload:
		return psess.rewriteValueS390X_OpS390XMOVHZload_0(v)
	case OpS390XMOVHZloadidx:
		return rewriteValueS390X_OpS390XMOVHZloadidx_0(v)
	case OpS390XMOVHZreg:
		return psess.rewriteValueS390X_OpS390XMOVHZreg_0(v) || rewriteValueS390X_OpS390XMOVHZreg_10(v)
	case OpS390XMOVHload:
		return psess.rewriteValueS390X_OpS390XMOVHload_0(v)
	case OpS390XMOVHloadidx:
		return rewriteValueS390X_OpS390XMOVHloadidx_0(v)
	case OpS390XMOVHreg:
		return psess.rewriteValueS390X_OpS390XMOVHreg_0(v) || rewriteValueS390X_OpS390XMOVHreg_10(v)
	case OpS390XMOVHstore:
		return psess.rewriteValueS390X_OpS390XMOVHstore_0(v) || rewriteValueS390X_OpS390XMOVHstore_10(v)
	case OpS390XMOVHstoreconst:
		return rewriteValueS390X_OpS390XMOVHstoreconst_0(v)
	case OpS390XMOVHstoreidx:
		return rewriteValueS390X_OpS390XMOVHstoreidx_0(v) || rewriteValueS390X_OpS390XMOVHstoreidx_10(v)
	case OpS390XMOVWBRstore:
		return rewriteValueS390X_OpS390XMOVWBRstore_0(v)
	case OpS390XMOVWBRstoreidx:
		return rewriteValueS390X_OpS390XMOVWBRstoreidx_0(v)
	case OpS390XMOVWZload:
		return psess.rewriteValueS390X_OpS390XMOVWZload_0(v)
	case OpS390XMOVWZloadidx:
		return rewriteValueS390X_OpS390XMOVWZloadidx_0(v)
	case OpS390XMOVWZreg:
		return psess.rewriteValueS390X_OpS390XMOVWZreg_0(v) || rewriteValueS390X_OpS390XMOVWZreg_10(v)
	case OpS390XMOVWload:
		return psess.rewriteValueS390X_OpS390XMOVWload_0(v)
	case OpS390XMOVWloadidx:
		return rewriteValueS390X_OpS390XMOVWloadidx_0(v)
	case OpS390XMOVWreg:
		return psess.rewriteValueS390X_OpS390XMOVWreg_0(v) || rewriteValueS390X_OpS390XMOVWreg_10(v)
	case OpS390XMOVWstore:
		return psess.rewriteValueS390X_OpS390XMOVWstore_0(v) || rewriteValueS390X_OpS390XMOVWstore_10(v)
	case OpS390XMOVWstoreconst:
		return rewriteValueS390X_OpS390XMOVWstoreconst_0(v)
	case OpS390XMOVWstoreidx:
		return rewriteValueS390X_OpS390XMOVWstoreidx_0(v) || rewriteValueS390X_OpS390XMOVWstoreidx_10(v)
	case OpS390XMULLD:
		return psess.rewriteValueS390X_OpS390XMULLD_0(v)
	case OpS390XMULLDconst:
		return rewriteValueS390X_OpS390XMULLDconst_0(v)
	case OpS390XMULLDload:
		return rewriteValueS390X_OpS390XMULLDload_0(v)
	case OpS390XMULLW:
		return psess.rewriteValueS390X_OpS390XMULLW_0(v)
	case OpS390XMULLWconst:
		return rewriteValueS390X_OpS390XMULLWconst_0(v)
	case OpS390XMULLWload:
		return rewriteValueS390X_OpS390XMULLWload_0(v)
	case OpS390XNEG:
		return rewriteValueS390X_OpS390XNEG_0(v)
	case OpS390XNEGW:
		return rewriteValueS390X_OpS390XNEGW_0(v)
	case OpS390XNOT:
		return rewriteValueS390X_OpS390XNOT_0(v)
	case OpS390XNOTW:
		return rewriteValueS390X_OpS390XNOTW_0(v)
	case OpS390XOR:
		return rewriteValueS390X_OpS390XOR_0(v) || psess.rewriteValueS390X_OpS390XOR_10(v) || rewriteValueS390X_OpS390XOR_20(v) || rewriteValueS390X_OpS390XOR_30(v) || rewriteValueS390X_OpS390XOR_40(v) || rewriteValueS390X_OpS390XOR_50(v) || rewriteValueS390X_OpS390XOR_60(v) || rewriteValueS390X_OpS390XOR_70(v) || rewriteValueS390X_OpS390XOR_80(v) || rewriteValueS390X_OpS390XOR_90(v) || rewriteValueS390X_OpS390XOR_100(v) || rewriteValueS390X_OpS390XOR_110(v) || rewriteValueS390X_OpS390XOR_120(v) || rewriteValueS390X_OpS390XOR_130(v) || rewriteValueS390X_OpS390XOR_140(v) || rewriteValueS390X_OpS390XOR_150(v)
	case OpS390XORW:
		return psess.rewriteValueS390X_OpS390XORW_0(v) || psess.rewriteValueS390X_OpS390XORW_10(v) || rewriteValueS390X_OpS390XORW_20(v) || rewriteValueS390X_OpS390XORW_30(v) || rewriteValueS390X_OpS390XORW_40(v) || rewriteValueS390X_OpS390XORW_50(v) || rewriteValueS390X_OpS390XORW_60(v) || rewriteValueS390X_OpS390XORW_70(v) || rewriteValueS390X_OpS390XORW_80(v) || rewriteValueS390X_OpS390XORW_90(v)
	case OpS390XORWconst:
		return rewriteValueS390X_OpS390XORWconst_0(v)
	case OpS390XORWload:
		return rewriteValueS390X_OpS390XORWload_0(v)
	case OpS390XORconst:
		return rewriteValueS390X_OpS390XORconst_0(v)
	case OpS390XORload:
		return rewriteValueS390X_OpS390XORload_0(v)
	case OpS390XSLD:
		return rewriteValueS390X_OpS390XSLD_0(v) || rewriteValueS390X_OpS390XSLD_10(v)
	case OpS390XSLW:
		return rewriteValueS390X_OpS390XSLW_0(v) || rewriteValueS390X_OpS390XSLW_10(v)
	case OpS390XSRAD:
		return rewriteValueS390X_OpS390XSRAD_0(v) || rewriteValueS390X_OpS390XSRAD_10(v)
	case OpS390XSRADconst:
		return rewriteValueS390X_OpS390XSRADconst_0(v)
	case OpS390XSRAW:
		return rewriteValueS390X_OpS390XSRAW_0(v) || rewriteValueS390X_OpS390XSRAW_10(v)
	case OpS390XSRAWconst:
		return rewriteValueS390X_OpS390XSRAWconst_0(v)
	case OpS390XSRD:
		return rewriteValueS390X_OpS390XSRD_0(v) || rewriteValueS390X_OpS390XSRD_10(v)
	case OpS390XSRDconst:
		return rewriteValueS390X_OpS390XSRDconst_0(v)
	case OpS390XSRW:
		return rewriteValueS390X_OpS390XSRW_0(v) || rewriteValueS390X_OpS390XSRW_10(v)
	case OpS390XSTM2:
		return rewriteValueS390X_OpS390XSTM2_0(v)
	case OpS390XSTMG2:
		return rewriteValueS390X_OpS390XSTMG2_0(v)
	case OpS390XSUB:
		return psess.rewriteValueS390X_OpS390XSUB_0(v)
	case OpS390XSUBW:
		return psess.rewriteValueS390X_OpS390XSUBW_0(v)
	case OpS390XSUBWconst:
		return rewriteValueS390X_OpS390XSUBWconst_0(v)
	case OpS390XSUBWload:
		return rewriteValueS390X_OpS390XSUBWload_0(v)
	case OpS390XSUBconst:
		return rewriteValueS390X_OpS390XSUBconst_0(v)
	case OpS390XSUBload:
		return rewriteValueS390X_OpS390XSUBload_0(v)
	case OpS390XXOR:
		return psess.rewriteValueS390X_OpS390XXOR_0(v) || psess.rewriteValueS390X_OpS390XXOR_10(v)
	case OpS390XXORW:
		return psess.rewriteValueS390X_OpS390XXORW_0(v) || psess.rewriteValueS390X_OpS390XXORW_10(v)
	case OpS390XXORWconst:
		return rewriteValueS390X_OpS390XXORWconst_0(v)
	case OpS390XXORWload:
		return rewriteValueS390X_OpS390XXORWload_0(v)
	case OpS390XXORconst:
		return rewriteValueS390X_OpS390XXORconst_0(v)
	case OpS390XXORload:
		return rewriteValueS390X_OpS390XXORload_0(v)
	case OpSelect0:
		return rewriteValueS390X_OpSelect0_0(v)
	case OpSelect1:
		return rewriteValueS390X_OpSelect1_0(v)
	case OpSignExt16to32:
		return rewriteValueS390X_OpSignExt16to32_0(v)
	case OpSignExt16to64:
		return rewriteValueS390X_OpSignExt16to64_0(v)
	case OpSignExt32to64:
		return rewriteValueS390X_OpSignExt32to64_0(v)
	case OpSignExt8to16:
		return rewriteValueS390X_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValueS390X_OpSignExt8to32_0(v)
	case OpSignExt8to64:
		return rewriteValueS390X_OpSignExt8to64_0(v)
	case OpSlicemask:
		return rewriteValueS390X_OpSlicemask_0(v)
	case OpSqrt:
		return rewriteValueS390X_OpSqrt_0(v)
	case OpStaticCall:
		return rewriteValueS390X_OpStaticCall_0(v)
	case OpStore:
		return psess.rewriteValueS390X_OpStore_0(v)
	case OpSub16:
		return rewriteValueS390X_OpSub16_0(v)
	case OpSub32:
		return rewriteValueS390X_OpSub32_0(v)
	case OpSub32F:
		return rewriteValueS390X_OpSub32F_0(v)
	case OpSub64:
		return rewriteValueS390X_OpSub64_0(v)
	case OpSub64F:
		return rewriteValueS390X_OpSub64F_0(v)
	case OpSub8:
		return rewriteValueS390X_OpSub8_0(v)
	case OpSubPtr:
		return rewriteValueS390X_OpSubPtr_0(v)
	case OpTrunc:
		return rewriteValueS390X_OpTrunc_0(v)
	case OpTrunc16to8:
		return rewriteValueS390X_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValueS390X_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValueS390X_OpTrunc32to8_0(v)
	case OpTrunc64to16:
		return rewriteValueS390X_OpTrunc64to16_0(v)
	case OpTrunc64to32:
		return rewriteValueS390X_OpTrunc64to32_0(v)
	case OpTrunc64to8:
		return rewriteValueS390X_OpTrunc64to8_0(v)
	case OpWB:
		return rewriteValueS390X_OpWB_0(v)
	case OpXor16:
		return rewriteValueS390X_OpXor16_0(v)
	case OpXor32:
		return rewriteValueS390X_OpXor32_0(v)
	case OpXor64:
		return rewriteValueS390X_OpXor64_0(v)
	case OpXor8:
		return rewriteValueS390X_OpXor8_0(v)
	case OpZero:
		return psess.rewriteValueS390X_OpZero_0(v) || rewriteValueS390X_OpZero_10(v)
	case OpZeroExt16to32:
		return rewriteValueS390X_OpZeroExt16to32_0(v)
	case OpZeroExt16to64:
		return rewriteValueS390X_OpZeroExt16to64_0(v)
	case OpZeroExt32to64:
		return rewriteValueS390X_OpZeroExt32to64_0(v)
	case OpZeroExt8to16:
		return rewriteValueS390X_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValueS390X_OpZeroExt8to32_0(v)
	case OpZeroExt8to64:
		return rewriteValueS390X_OpZeroExt8to64_0(v)
	}
	return false
}
func rewriteValueS390X_OpAdd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XADDW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpAdd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XADDW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpAdd32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XFADDS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpAdd64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpAdd64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XFADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpAdd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XADDW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpAddPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpAddr_0(v *Value) bool {

	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpS390XMOVDaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueS390X_OpAnd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XANDW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpAnd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XANDW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpAnd64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpAnd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XANDW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpAndB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XANDW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpAtomicAdd32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XAddTupleFirst32)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpS390XLAA, types.NewTuple(typ.UInt32, psess.types.TypeMem))
		v0.AddArg(ptr)
		v0.AddArg(val)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpAtomicAdd64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XAddTupleFirst64)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpS390XLAAG, types.NewTuple(typ.UInt64, psess.types.TypeMem))
		v0.AddArg(ptr)
		v0.AddArg(val)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpAtomicCompareAndSwap32_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpS390XLoweredAtomicCas32)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAtomicCompareAndSwap64_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpS390XLoweredAtomicCas64)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAtomicExchange32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XLoweredAtomicExchange32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAtomicExchange64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XLoweredAtomicExchange64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAtomicLoad32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpS390XMOVWZatomicload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAtomicLoad64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpS390XMOVDatomicload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAtomicLoadPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpS390XMOVDatomicload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAtomicStore32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XMOVWatomicstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAtomicStore64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XMOVDatomicstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAtomicStorePtrNoWB_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XMOVDatomicstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpAvg64u_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XADD)
		v0 := b.NewValue0(v.Pos, OpS390XSRDconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpS390XSUB, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpBitLen64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpS390XSUB)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 64
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XFLOGR, typ.UInt64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpBswap32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVWBR)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpBswap64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVDBR)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCeil_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XFIDBR)
		v.AuxInt = 6
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpClosureCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XCALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpCom16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XNOTW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCom32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XNOTW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCom64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XNOT)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCom8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XNOTW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpConst16_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueS390X_OpConst32_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueS390X_OpConst32F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpS390XFMOVSconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueS390X_OpConst64_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueS390X_OpConst64F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpS390XFMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueS390X_OpConst8_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueS390X_OpConstBool_0(v *Value) bool {

	for {
		b := v.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValueS390X_OpConstNil_0(v *Value) bool {

	for {
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueS390X_OpCtz32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpS390XSUB)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 64
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XFLOGR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XANDW, t)
		v4 := b.NewValue0(v.Pos, OpS390XSUBWconst, t)
		v4.AuxInt = 1
		v4.AddArg(x)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpS390XNOTW, t)
		v5.AddArg(x)
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpCtz32NonZero_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCtz32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCtz64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpS390XSUB)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 64
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XFLOGR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpS390XAND, t)
		v3 := b.NewValue0(v.Pos, OpS390XSUBconst, t)
		v3.AuxInt = 1
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XNOT, t)
		v4.AddArg(x)
		v2.AddArg(v4)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpCtz64NonZero_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCtz64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCvt32Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XCFEBRA)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCvt32Fto64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XCGEBRA)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCvt32Fto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XLDEBR)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCvt32to32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XCEFBRA)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCvt32to64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XCDFBRA)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCvt64Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XCFDBRA)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCvt64Fto32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XLEDBR)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCvt64Fto64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XCGDBRA)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCvt64to32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XCEGBRA)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpCvt64to64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XCDGBRA)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpDiv16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XDIVW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpDiv16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XDIVWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpDiv32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XDIVW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpDiv32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XFDIVS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpDiv32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XDIVWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpDiv64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XDIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpDiv64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XFDIV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpDiv64u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XDIVDU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpDiv8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XDIVW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpDiv8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XDIVWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpEq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDEQ)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpEq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDEQ)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpEq32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDEQ)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpEq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDEQ)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpEq64F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDEQ)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpEq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDEQ)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpEqB_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDEQ)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpEqPtr_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDEQ)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpFloor_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XFIDBR)
		v.AuxInt = 7
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGeq32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGEnoinv)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGeq32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGeq64F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGEnoinv)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGeq64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPU, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpGetCallerPC_0(v *Value) bool {

	for {
		v.reset(OpS390XLoweredGetCallerPC)
		return true
	}
}
func rewriteValueS390X_OpGetCallerSP_0(v *Value) bool {

	for {
		v.reset(OpS390XLoweredGetCallerSP)
		return true
	}
}
func rewriteValueS390X_OpGetClosurePtr_0(v *Value) bool {

	for {
		v.reset(OpS390XLoweredGetClosurePtr)
		return true
	}
}
func rewriteValueS390X_OpGetG_0(v *Value) bool {

	for {
		mem := v.Args[0]
		v.reset(OpS390XLoweredGetG)
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGreater16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGreater16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGreater32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGreater32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGTnoinv)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGreater32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGreater64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGreater64F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGTnoinv)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGreater64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPU, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGreater8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpGreater8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpHmul32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRDconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpS390XMULLD, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWreg, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XMOVWreg, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpHmul32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRDconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpS390XMULLD, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpHmul64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMULHD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpHmul64u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMULHDU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpITab_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpLoad {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_0.Args[1]
		v.reset(OpS390XMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpInterCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[1]
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpS390XCALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpIsInBounds_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpS390XMOVDLT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPU, psess.types.TypeFlags)
		v2.AddArg(idx)
		v2.AddArg(len)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		p := v.Args[0]
		v.reset(OpS390XMOVDNE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPconst, psess.types.TypeFlags)
		v2.AuxInt = 0
		v2.AddArg(p)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpIsSliceInBounds_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpS390XMOVDLE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPU, psess.types.TypeFlags)
		v2.AddArg(idx)
		v2.AddArg(len)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLeq32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGEnoinv)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, psess.types.TypeFlags)
		v2.AddArg(y)
		v2.AddArg(x)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLeq32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLeq64F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGEnoinv)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, psess.types.TypeFlags)
		v2.AddArg(y)
		v2.AddArg(x)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLeq64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPU, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLess16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLess16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLess32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLess32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGTnoinv)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, psess.types.TypeFlags)
		v2.AddArg(y)
		v2.AddArg(x)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLess32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLess64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLess64F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGTnoinv)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, psess.types.TypeFlags)
		v2.AddArg(y)
		v2.AddArg(x)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLess64U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPU, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLess8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLess8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDLT)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWU, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLoad_0(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpS390XMOVDload)
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
		v.reset(OpS390XMOVWload)
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
		v.reset(OpS390XMOVWZload)
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
		v.reset(OpS390XMOVHload)
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
		v.reset(OpS390XMOVHZload)
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
		v.reset(OpS390XMOVBload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsBoolean() || (psess.is8BitInt(t) && !isSigned(t))) {
			break
		}
		v.reset(OpS390XMOVBZload)
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
		v.reset(OpS390XFMOVSload)
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
		v.reset(OpS390XFMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSLW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpMod16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMODW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpMod16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMODWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpMod32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMODW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMod32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMODWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMod64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMODD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMod64u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMODDU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMod8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMODW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueS390X_OpMod8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMODWU)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpMove_0(v *Value) bool {
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
		v.reset(OpS390XMOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, typ.UInt8)
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
		v.reset(OpS390XMOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
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
		v.reset(OpS390XMOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
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
		v.reset(OpS390XMOVDstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 16 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XMOVDstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
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
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XMOVDstore)
		v.AuxInt = 16
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v0.AuxInt = 16
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDstore, psess.types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v2.AuxInt = 8
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XMOVDstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
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
		v.reset(OpS390XMOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
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
		v.reset(OpS390XMOVBstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, typ.UInt8)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
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
		v.reset(OpS390XMOVHstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpMove_10(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		if v.AuxInt != 7 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XMOVBstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, typ.UInt8)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHstore, psess.types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
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
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 0 && s <= 256) {
			break
		}
		v.reset(OpS390XMVC)
		v.AuxInt = makeValAndOff(s, 0)
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 256 && s <= 512) {
			break
		}
		v.reset(OpS390XMVC)
		v.AuxInt = makeValAndOff(s-256, 256)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpS390XMVC, psess.types.TypeMem)
		v0.AuxInt = makeValAndOff(256, 0)
		v0.AddArg(dst)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 512 && s <= 768) {
			break
		}
		v.reset(OpS390XMVC)
		v.AuxInt = makeValAndOff(s-512, 512)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpS390XMVC, psess.types.TypeMem)
		v0.AuxInt = makeValAndOff(256, 256)
		v0.AddArg(dst)
		v0.AddArg(src)
		v1 := b.NewValue0(v.Pos, OpS390XMVC, psess.types.TypeMem)
		v1.AuxInt = makeValAndOff(256, 0)
		v1.AddArg(dst)
		v1.AddArg(src)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 768 && s <= 1024) {
			break
		}
		v.reset(OpS390XMVC)
		v.AuxInt = makeValAndOff(s-768, 768)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpS390XMVC, psess.types.TypeMem)
		v0.AuxInt = makeValAndOff(256, 512)
		v0.AddArg(dst)
		v0.AddArg(src)
		v1 := b.NewValue0(v.Pos, OpS390XMVC, psess.types.TypeMem)
		v1.AuxInt = makeValAndOff(256, 256)
		v1.AddArg(dst)
		v1.AddArg(src)
		v2 := b.NewValue0(v.Pos, OpS390XMVC, psess.types.TypeMem)
		v2.AuxInt = makeValAndOff(256, 0)
		v2.AddArg(dst)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 1024) {
			break
		}
		v.reset(OpS390XLoweredMove)
		v.AuxInt = s % 256
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpS390XADDconst, src.Type)
		v0.AuxInt = (s / 256) * 256
		v0.AddArg(src)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpMul16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMULLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMul32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMULLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMul32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XFMULS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMul64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMULLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMul64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XFMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpMul8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMULLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpNeg16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XNEGW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpNeg32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XNEGW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpNeg32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XFNEGS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpNeg64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpNeg64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XFNEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpNeg8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XNEGW)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpNeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDNE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpNeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDNE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpNeq32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDNE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMPS, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpNeq64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDNE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpNeq64F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDNE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XFCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpNeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDNE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpNeqB_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDNE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPW, psess.types.TypeFlags)
		v3 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v4.AddArg(y)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpNeqPtr_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDNE)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 1
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMP, psess.types.TypeFlags)
		v2.AddArg(x)
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueS390X_OpNilCheck_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpS390XLoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpNot_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XXORWconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpOffPtr_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpS390XMOVDaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		if !(is32Bit(off)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpS390XADD)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = off
		v.AddArg(v0)
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueS390X_OpOr16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XORW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpOr32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XORW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpOr64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpOr8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XORW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpOrB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XORW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpRound_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XFIDBR)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpRound32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XLoweredRound32F)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpRound64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XLoweredRound64F)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpRoundToEven_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XFIDBR)
		v.AuxInt = 4
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh16Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh32Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh64Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh64Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh64Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh64Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v3 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh8Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XMOVDGE)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XSRW, t)
		v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSRAW)
		v0 := b.NewValue0(v.Pos, OpS390XMOVBreg, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVDGE, y.Type)
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDconst, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpS390XADD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSLDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRDconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpS390XRLLGconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSRDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSLDconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpS390XRLLGconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		c := v_1.AuxInt
		s := v_1.Aux
		ptr := v_1.Args[0]
		if !(ptr.Op != OpSB && idx.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(idx)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		c := v_0.AuxInt
		s := v_0.Aux
		ptr := v_0.Args[0]
		idx := v.Args[1]
		if !(ptr.Op != OpSB && idx.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(idx)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XNEG {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XNEG {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpS390XSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XADD_10(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XADDW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XADDWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpS390XADDWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSLWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRWconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpS390XRLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSRWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSLWconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpS390XRLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XNEGW {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSUBW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XNEGW {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpS390XSUBW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XADDW_10(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XADDWconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int32(c + d))
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpS390XADDWconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XADDWload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XADDWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XADDconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		x := v_0.Args[0]
		if x.Op != OpSB {
			break
		}
		if !(((c+d)&1 == 0) && is32Bit(c+d)) {
			break
		}
		v.reset(OpS390XMOVDaddr)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		x := v_0.Args[0]
		if !(x.Op != OpSB && is20Bit(c+d)) {
			break
		}
		v.reset(OpS390XMOVDaddr)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
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
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c + d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XADDload_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr1 := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFMOVDstore {
			break
		}
		if v_2.AuxInt != off {
			break
		}
		if v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XADD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XADDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XADDload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XAND_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c) && c < 0) {
			break
		}
		v.reset(OpS390XANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c) && c < 0) {
			break
		}
		v.reset(OpS390XANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c) && c >= 0) {
			break
		}
		v.reset(OpS390XMOVWZreg)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = int64(int32(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c) && c >= 0) {
			break
		}
		v.reset(OpS390XMOVWZreg)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = int64(int32(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		if v_1.AuxInt != 0xFF {
			break
		}
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		if v_0.AuxInt != 0xFF {
			break
		}
		x := v.Args[1]
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		if v_1.AuxInt != 0xFFFF {
			break
		}
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		if v_0.AuxInt != 0xFFFF {
			break
		}
		x := v.Args[1]
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		if v_1.AuxInt != 0xFFFFFFFF {
			break
		}
		v.reset(OpS390XMOVWZreg)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		if v_0.AuxInt != 0xFFFFFFFF {
			break
		}
		x := v.Args[1]
		v.reset(OpS390XMOVWZreg)
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XAND_10(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c & d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c & d
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
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XANDW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XANDWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpS390XANDWconst)
		v.AuxInt = int64(int32(c))
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
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XANDW_10(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XANDWconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XANDWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpS390XANDWconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0xFF {
			break
		}
		x := v.Args[0]
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0xFFFF {
			break
		}
		x := v.Args[0]
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c & d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XANDWload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XANDWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XANDconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpS390XANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpS390XMOVDconst)
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
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c & d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XANDload_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr1 := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFMOVDstore {
			break
		}
		if v_2.AuxInt != off {
			break
		}
		if v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XAND)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XANDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XANDload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XCMP_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XCMPconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XCMPU_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpS390XCMPUconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMPUconst, psess.types.TypeFlags)
		v0.AuxInt = int64(int32(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPUconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) == uint64(y)) {
			break
		}
		v.reset(OpS390XFlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) < uint64(y)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint64(x) > uint64(y)) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSRDconst {
			break
		}
		c := v_0.AuxInt
		if !(c > 0 && c < 64 && (1<<uint(64-c)) <= uint64(n)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if x.Op != OpS390XMOVHreg {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if x.Op != OpS390XMOVHZreg {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if x.Op != OpS390XMOVBreg {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if x.Op != OpS390XMOVBZreg {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpS390XANDWconst {
			break
		}
		m := x.AuxInt
		if !(int32(m) >= 0) {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPUconst_10(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpS390XANDWconst {
			break
		}
		m := x.AuxInt
		if !(int32(m) >= 0) {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XCMPW_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XCMPWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMPWconst, psess.types.TypeFlags)
		v0.AuxInt = int64(int32(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XCMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XCMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpS390XCMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpS390XCMPW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XCMPWU_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpS390XInvertFlags)
		v0 := b.NewValue0(v.Pos, OpS390XCMPWUconst, psess.types.TypeFlags)
		v0.AuxInt = int64(int32(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XCMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XCMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpS390XCMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpS390XCMPWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPWUconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) == uint32(y)) {
			break
		}
		v.reset(OpS390XFlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVBZreg {
			break
		}
		if !(0xff < c) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVHZreg {
			break
		}
		if !(0xffff < c) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSRWconst {
			break
		}
		c := v_0.AuxInt
		if !(c > 0 && c < 32 && (1<<uint(32-c)) <= uint32(n)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		if !(uint32(m) < uint32(n)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPWconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpS390XFlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y)) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVBZreg {
			break
		}
		if !(0xff < c) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVHZreg {
			break
		}
		if !(0xffff < c) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSRWconst {
			break
		}
		c := v_0.AuxInt
		if !(c > 0 && n < 0) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		if !(int32(m) >= 0 && int32(m) < int32(n)) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		n := v.AuxInt
		x := v.Args[0]
		if x.Op != OpS390XSRWconst {
			break
		}
		c := x.AuxInt
		if !(c > 0 && n >= 0) {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = n
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x == y) {
			break
		}
		v.reset(OpS390XFlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x < y) {
			break
		}
		v.reset(OpS390XFlagLT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x > y) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSRDconst {
			break
		}
		c := v_0.AuxInt
		if !(c > 0 && n < 0) {
			break
		}
		v.reset(OpS390XFlagGT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if x.Op != OpS390XMOVHreg {
			break
		}
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if x.Op != OpS390XMOVHZreg {
			break
		}
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if x.Op != OpS390XMOVBreg {
			break
		}
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if x.Op != OpS390XMOVBZreg {
			break
		}
		v.reset(OpS390XCMPWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpS390XANDWconst {
			break
		}
		m := x.AuxInt
		if !(int32(m) >= 0 && c >= 0) {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCMPconst_10(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpS390XANDWconst {
			break
		}
		m := x.AuxInt
		if !(int32(m) >= 0 && c >= 0) {
			break
		}
		v.reset(OpS390XCMPWUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		n := v.AuxInt
		x := v.Args[0]
		if x.Op != OpS390XSRDconst {
			break
		}
		c := x.AuxInt
		if !(c > 0 && n >= 0) {
			break
		}
		v.reset(OpS390XCMPUconst)
		v.AuxInt = n
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XCPSDR_0(v *Value) bool {

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XFMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c&-1<<63 == 0) {
			break
		}
		v.reset(OpS390XLPDFR)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XFMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c&-1<<63 != 0) {
			break
		}
		v.reset(OpS390XLNDFR)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFADD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XFMUL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpS390XFMADD)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XFMUL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpS390XFMADD)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFADDS_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XFMULS {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpS390XFMADDS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XFMULS {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpS390XFMADDS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVDload_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XLDGR)
		v.AddArg(x)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XFMOVDstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XFMOVDload)
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
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVDloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XFMOVDloadidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVDloadidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVDloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVDloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVDstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XFMOVDstore)
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
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVDstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XFMOVDstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVDstoreidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVDstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVDstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVSload_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XFMOVSstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XFMOVSload)
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
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVSloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XFMOVSloadidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVSloadidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVSloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVSloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVSstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XFMOVSstore)
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
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVSstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XFMOVSstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XFMOVSstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFMOVSstoreidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVSstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XFMOVSstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFNEG_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XLPDFR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XLNDFR)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XLNDFR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XLPDFR)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFNEGS_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XLPDFR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XLNDFR)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XLNDFR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XLPDFR)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFSUB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XFMUL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpS390XFMSUB)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XFSUBS_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XFMULS {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpS390XFMSUBS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XLDGR_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSRDconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XSLDconst {
			break
		}
		if v_0_0.AuxInt != 1 {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpS390XLPDFR)
		v0 := b.NewValue0(v.Pos, OpS390XLDGR, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpS390XOR {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XMOVDconst {
			break
		}
		if v_0_0.AuxInt != -1<<63 {
			break
		}
		x := v_0.Args[1]
		v.reset(OpS390XLNDFR)
		v0 := b.NewValue0(v.Pos, OpS390XLDGR, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpS390XOR {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpS390XMOVDconst {
			break
		}
		if v_0_1.AuxInt != -1<<63 {
			break
		}
		v.reset(OpS390XLNDFR)
		v0 := b.NewValue0(v.Pos, OpS390XLDGR, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XORload {
			break
		}
		t1 := x.Type
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		x_0 := x.Args[0]
		if x_0.Op != OpS390XMOVDconst {
			break
		}
		if x_0.AuxInt != -1<<63 {
			break
		}
		ptr := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XLNDFR, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XLDGR, t)
		v2 := b.NewValue0(v.Pos, OpS390XMOVDload, t1)
		v2.AuxInt = off
		v2.Aux = sym
		v2.AddArg(ptr)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XLGDR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XLEDBR_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XLPDFR {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XLDEBR {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpS390XLPDFR)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XLNDFR {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XLDEBR {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpS390XLNDFR)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XLGDR_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XLDGR {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XLoweredRound32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpS390XFMOVSconst {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XLoweredRound64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpS390XFMOVDconst {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBZload_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVBZload)
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
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBZloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVBZloadidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBZloadidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVBZreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVDLT {
			break
		}
		_ = x.Args[2]
		x_0 := x.Args[0]
		if x_0.Op != OpS390XMOVDconst {
			break
		}
		c := x_0.AuxInt
		x_1 := x.Args[1]
		if x_1.Op != OpS390XMOVDconst {
			break
		}
		d := x_1.AuxInt
		if !(int64(uint8(c)) == c && int64(uint8(d)) == d) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVDLE {
			break
		}
		_ = x.Args[2]
		x_0 := x.Args[0]
		if x_0.Op != OpS390XMOVDconst {
			break
		}
		c := x_0.AuxInt
		x_1 := x.Args[1]
		if x_1.Op != OpS390XMOVDconst {
			break
		}
		d := x_1.AuxInt
		if !(int64(uint8(c)) == c && int64(uint8(d)) == d) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVDGT {
			break
		}
		_ = x.Args[2]
		x_0 := x.Args[0]
		if x_0.Op != OpS390XMOVDconst {
			break
		}
		c := x_0.AuxInt
		x_1 := x.Args[1]
		if x_1.Op != OpS390XMOVDconst {
			break
		}
		d := x_1.AuxInt
		if !(int64(uint8(c)) == c && int64(uint8(d)) == d) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVDGE {
			break
		}
		_ = x.Args[2]
		x_0 := x.Args[0]
		if x_0.Op != OpS390XMOVDconst {
			break
		}
		c := x_0.AuxInt
		x_1 := x.Args[1]
		if x_1.Op != OpS390XMOVDconst {
			break
		}
		d := x_1.AuxInt
		if !(int64(uint8(c)) == c && int64(uint8(d)) == d) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVDEQ {
			break
		}
		_ = x.Args[2]
		x_0 := x.Args[0]
		if x_0.Op != OpS390XMOVDconst {
			break
		}
		c := x_0.AuxInt
		x_1 := x.Args[1]
		if x_1.Op != OpS390XMOVDconst {
			break
		}
		d := x_1.AuxInt
		if !(int64(uint8(c)) == c && int64(uint8(d)) == d) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVDNE {
			break
		}
		_ = x.Args[2]
		x_0 := x.Args[0]
		if x_0.Op != OpS390XMOVDconst {
			break
		}
		c := x_0.AuxInt
		x_1 := x.Args[1]
		if x_1.Op != OpS390XMOVDconst {
			break
		}
		d := x_1.AuxInt
		if !(int64(uint8(c)) == c && int64(uint8(d)) == d) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVDGTnoinv {
			break
		}
		_ = x.Args[2]
		x_0 := x.Args[0]
		if x_0.Op != OpS390XMOVDconst {
			break
		}
		c := x_0.AuxInt
		x_1 := x.Args[1]
		if x_1.Op != OpS390XMOVDconst {
			break
		}
		d := x_1.AuxInt
		if !(int64(uint8(c)) == c && int64(uint8(d)) == d) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVDGEnoinv {
			break
		}
		_ = x.Args[2]
		x_0 := x.Args[0]
		if x_0.Op != OpS390XMOVDconst {
			break
		}
		c := x_0.AuxInt
		x_1 := x.Args[1]
		if x_1.Op != OpS390XMOVDconst {
			break
		}
		d := x_1.AuxInt
		if !(int64(uint8(c)) == c && int64(uint8(d)) == d) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !(psess.is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBZreg_10(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVBreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(uint8(c))
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpS390XMOVWZreg)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = int64(uint8(m))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBload_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVBload)
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
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVBloadidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBloadidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVBreg_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !(psess.is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVBZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int8(c))
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVBload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVBload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVBloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVBloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		x := v_0.Args[0]
		if !(int8(m) >= 0) {
			break
		}
		v.reset(OpS390XMOVWZreg)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = int64(uint8(m))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBstore_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpS390XMOVBstore)
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
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpS390XMOVBstore)
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
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(is20Bit(off) && ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = makeValAndOff(int64(int8(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVBstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVBstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpS390XMOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRDconst {
			break
		}
		if x_1.AuxInt != 8 {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w0 := v.Args[1]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRDconst {
			break
		}
		if x_1.AuxInt != j+8 {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpS390XMOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRWconst {
			break
		}
		if x_1.AuxInt != 8 {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBstore_10(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w0 := v.Args[1]
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRWconst {
			break
		}
		if x_1.AuxInt != j+8 {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRDconst {
			break
		}
		if v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRDconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRWconst {
			break
		}
		if v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRWconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpS390XSRWconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBstoreconst_0(v *Value) bool {

	for {
		sc := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is20Bit(ValAndOff(sc).Off() + off)) {
			break
		}
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		sc := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != OpS390XMOVBstoreconst {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		_ = x.Args[1]
		if p != x.Args[0] {
			break
		}
		mem := x.Args[1]
		if !(p.Op != OpSB && x.Uses == 1 && ValAndOff(a).Off()+1 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = makeValAndOff(ValAndOff(c).Val()&0xff|ValAndOff(a).Val()<<8, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBstoreidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVBstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBstoreidx_10(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != 8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != 8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != 8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != 8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != j+8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != j+8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != j+8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != j+8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBstoreidx_20(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVBstoreidx_30(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVBstoreidx {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVHBRstoreidx)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDEQ_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XInvertFlags {
			break
		}
		cmp := v_2.Args[0]
		v.reset(OpS390XMOVDEQ)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cmp)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDGE_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XInvertFlags {
			break
		}
		cmp := v_2.Args[0]
		v.reset(OpS390XMOVDLE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cmp)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDGT_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XInvertFlags {
			break
		}
		cmp := v_2.Args[0]
		v.reset(OpS390XMOVDLT)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cmp)
		return true
	}

	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDLE_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XInvertFlags {
			break
		}
		cmp := v_2.Args[0]
		v.reset(OpS390XMOVDGE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cmp)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDLT_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XInvertFlags {
			break
		}
		cmp := v_2.Args[0]
		v.reset(OpS390XMOVDGT)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cmp)
		return true
	}

	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDNE_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XInvertFlags {
			break
		}
		cmp := v_2.Args[0]
		v.reset(OpS390XMOVDNE)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cmp)
		return true
	}

	for {
		_ = v.Args[2]
		y := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagEQ {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagLT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFlagGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDaddridx_0(v *Value) bool {

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		y := v.Args[1]
		if !(is20Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		if !(is20Bit(c+d) && y.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		y := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		y := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && y.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDaddridx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVDload_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XFMOVDstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XLGDR)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVDload)
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
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem(psess.types).Alignment(psess.types)%8 == 0 && (off1+off2)%8 == 0))) {
			break
		}
		v.reset(OpS390XMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVDloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDloadidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDloadidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVDloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVDloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVDloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVDloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVDnop_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		if !(t.Compare(psess.types, x.Type) == types.CMPeq) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVBZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVBload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVBload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVHZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVHload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVHload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVWZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVWload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVWload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVDload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVBZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDnop_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVBloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVBloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVHZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVHloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVHloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVWZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVWloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVWloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVDloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVDloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVDreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		if !(t.Compare(psess.types, x.Type) == types.CMPeq) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c
		return true
	}

	for {
		x := v.Args[0]
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpS390XMOVDnop)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVBZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVBload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVBload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVHZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVHload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVHload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVWZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVWload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVWload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVDload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVDload, t)
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
func rewriteValueS390X_OpS390XMOVDreg_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVBZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVBZloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVBloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVBloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVHZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVHloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVHloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVWZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVWloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVWloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if x.Op != OpS390XMOVDloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVDloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVDstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(is16Bit(c) && isU12Bit(off) && ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDstoreconst)
		v.AuxInt = makeValAndOff(c, off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem(psess.types).Alignment(psess.types)%8 == 0 && (off1+off2)%8 == 0))) {
			break
		}
		v.reset(OpS390XMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w1 := v.Args[1]
		x := v.Args[2]
		if x.Op != OpS390XMOVDstore {
			break
		}
		if x.AuxInt != i-8 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && is20Bit(i-8) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTMG2)
		v.AuxInt = i - 8
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w2 := v.Args[1]
		x := v.Args[2]
		if x.Op != OpS390XSTMG2 {
			break
		}
		if x.AuxInt != i-16 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		mem := x.Args[3]
		if !(x.Uses == 1 && is20Bit(i-16) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTMG3)
		v.AuxInt = i - 16
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w3 := v.Args[1]
		x := v.Args[2]
		if x.Op != OpS390XSTMG3 {
			break
		}
		if x.AuxInt != i-24 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[4]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		w2 := x.Args[3]
		mem := x.Args[4]
		if !(x.Uses == 1 && is20Bit(i-24) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTMG4)
		v.AuxInt = i - 24
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(w3)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDstoreconst_0(v *Value) bool {

	for {
		sc := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(isU12Bit(ValAndOff(sc).Off() + off)) {
			break
		}
		v.reset(OpS390XMOVDstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		sc := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpS390XMOVDstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVDstoreidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHBRstore_0(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRDconst {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVHBRstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRDconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVHBRstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRWconst {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVHBRstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRWconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVHBRstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpS390XSRWconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHBRstoreidx_0(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHBRstoreidx_10(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRWconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHBRstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWBRstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVHZload_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVHZload)
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
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem(psess.types).Alignment(psess.types)%2 == 0 && (off1+off2)%2 == 0))) {
			break
		}
		v.reset(OpS390XMOVHZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVHZloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVHZloadidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHZloadidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVHZreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !((psess.is8BitInt(t) || psess.is16BitInt(t)) && !isSigned(t)) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHZreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVHreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(uint16(c))
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHZreg_10(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpS390XMOVWZreg)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = int64(uint16(m))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVHload_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVHreg)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVHload)
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
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem(psess.types).Alignment(psess.types)%2 == 0 && (off1+off2)%2 == 0))) {
			break
		}
		v.reset(OpS390XMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVHloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVHloadidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHloadidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVHreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !((psess.is8BitInt(t) || psess.is16BitInt(t)) && isSigned(t)) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVHZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XMOVHreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int16(c))
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVHload, v.Type)
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
func rewriteValueS390X_OpS390XMOVHreg_10(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVHload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVHloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVHloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XANDWconst {
			break
		}
		m := v_0.AuxInt
		x := v_0.Args[0]
		if !(int16(m) >= 0) {
			break
		}
		v.reset(OpS390XMOVWZreg)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = int64(uint16(m))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVHstore_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpS390XMOVHstore)
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
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpS390XMOVHstore)
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
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(isU12Bit(off) && ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = makeValAndOff(int64(int16(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem(psess.types).Alignment(psess.types)%2 == 0 && (off1+off2)%2 == 0))) {
			break
		}
		v.reset(OpS390XMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpS390XMOVHstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRDconst {
			break
		}
		if x_1.AuxInt != 16 {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w0 := v.Args[1]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVHstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRDconst {
			break
		}
		if x_1.AuxInt != j+16 {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpS390XMOVHstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRWconst {
			break
		}
		if x_1.AuxInt != 16 {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHstore_10(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w0 := v.Args[1]
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVHstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRWconst {
			break
		}
		if x_1.AuxInt != j+16 {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHstoreconst_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		sc := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(isU12Bit(ValAndOff(sc).Off() + off)) {
			break
		}
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		sc := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != OpS390XMOVHstoreconst {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		_ = x.Args[1]
		if p != x.Args[0] {
			break
		}
		mem := x.Args[1]
		if !(p.Op != OpSB && x.Uses == 1 && ValAndOff(a).Off()+2 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = ValAndOff(a).Off()
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = int64(int32(ValAndOff(c).Val()&0xffff | ValAndOff(a).Val()<<16))
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHstoreidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVHstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVHstoreidx_10(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != 16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != 16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != 16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != 16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != j+16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != j+16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != j+16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRWconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVHstoreidx {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRWconst {
			break
		}
		if x_2.AuxInt != j+16 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWBRstore_0(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRDconst {
			break
		}
		if v_1.AuxInt != 32 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVWBRstore {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRDconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVWBRstore {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-32 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWBRstoreidx_0(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 32 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWBRstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 32 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWBRstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 32 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWBRstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		if v_2.AuxInt != 32 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWBRstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWBRstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-32 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWBRstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-32 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWBRstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-32 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XSRDconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWBRstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		w0 := x.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		if w0.AuxInt != j-32 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDBRstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVWZload_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVWZreg)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVWZload)
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
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem(psess.types).Alignment(psess.types)%4 == 0 && (off1+off2)%4 == 0))) {
			break
		}
		v.reset(OpS390XMOVWZload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVWZloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVWZloadidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWZloadidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWZloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVWZreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !((psess.is8BitInt(t) || psess.is16BitInt(t) || psess.is32BitInt(t)) && !isSigned(t)) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHZreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWZreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XMOVWZreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(uint32(c))
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZload, v.Type)
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
func rewriteValueS390X_OpS390XMOVWZreg_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVWload_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr1 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWstore {
			break
		}
		if v_1.AuxInt != off {
			break
		}
		if v_1.Aux != sym {
			break
		}
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVWload)
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
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem(psess.types).Alignment(psess.types)%4 == 0 && (off1+off2)%4 == 0))) {
			break
		}
		v.reset(OpS390XMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVWloadidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVWloadidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWloadidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWloadidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVWreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHZload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWload {
			break
		}
		_ = x.Args[1]
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpArg {
			break
		}
		t := x.Type
		if !((psess.is8BitInt(t) || psess.is16BitInt(t) || psess.is32BitInt(t)) && isSigned(t)) {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVBZreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVHZreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWreg_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWreg {
			break
		}
		v.reset(OpS390XMOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVWZreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpS390XMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int32(c))
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWZload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVWload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWload {
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
		v0 := b.NewValue0(v.Pos, OpS390XMOVWload, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWZloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVWloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpS390XMOVWloadidx {
			break
		}
		off := x.AuxInt
		sym := x.Aux
		_ = x.Args[2]
		ptr := x.Args[0]
		idx := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpS390XMOVWloadidx, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMOVWstore_0(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpS390XMOVWstore)
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
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpS390XMOVWstore)
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
		if v_0.Op != OpS390XADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is20Bit(off1 + off2)) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(is16Bit(c) && isU12Bit(off) && ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVWstoreconst)
		v.AuxInt = makeValAndOff(int64(int32(c)), off)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		t := v_0.Type
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || (t.IsPtr() && t.Elem(psess.types).Alignment(psess.types)%4 == 0 && (off1+off2)%4 == 0))) {
			break
		}
		v.reset(OpS390XMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddridx {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRDconst {
			break
		}
		if v_1.AuxInt != 32 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVWstore {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w0 := v.Args[1]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[2]
		if x.Op != OpS390XMOVWstore {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpS390XSRDconst {
			break
		}
		if x_1.AuxInt != j+32 {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w1 := v.Args[1]
		x := v.Args[2]
		if x.Op != OpS390XMOVWstore {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		mem := x.Args[2]
		if !(p.Op != OpSB && x.Uses == 1 && is20Bit(i-4) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTM2)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWstore_10(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w2 := v.Args[1]
		x := v.Args[2]
		if x.Op != OpS390XSTM2 {
			break
		}
		if x.AuxInt != i-8 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		mem := x.Args[3]
		if !(x.Uses == 1 && is20Bit(i-8) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTM3)
		v.AuxInt = i - 8
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w3 := v.Args[1]
		x := v.Args[2]
		if x.Op != OpS390XSTM3 {
			break
		}
		if x.AuxInt != i-12 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[4]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		w2 := x.Args[3]
		mem := x.Args[4]
		if !(x.Uses == 1 && is20Bit(i-12) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTM4)
		v.AuxInt = i - 12
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(w3)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWstoreconst_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		sc := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(isU12Bit(ValAndOff(sc).Off() + off)) {
			break
		}
		v.reset(OpS390XMOVWstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = s
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		sc := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDaddr {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(ptr.Op != OpSB && canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(OpS390XMOVWstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != OpS390XMOVWstoreconst {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		_ = x.Args[1]
		if p != x.Args[0] {
			break
		}
		mem := x.Args[1]
		if !(p.Op != OpSB && x.Uses == 1 && ValAndOff(a).Off()+4 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstore)
		v.AuxInt = ValAndOff(a).Off()
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = ValAndOff(c).Val()&0xffffffff | ValAndOff(a).Val()<<32
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWstoreidx_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		if !(is20Bit(c + d)) {
			break
		}
		v.reset(OpS390XMOVWstoreidx)
		v.AuxInt = c + d
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVWstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 32 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVWstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 32 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVWstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 32 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XMOVWstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != 32 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+32 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+32 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMOVWstoreidx_10(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+32 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		w0 := v.Args[2]
		if w0.Op != OpS390XSRDconst {
			break
		}
		j := w0.AuxInt
		w := w0.Args[0]
		x := v.Args[3]
		if x.Op != OpS390XMOVWstoreidx {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if idx != x.Args[0] {
			break
		}
		if p != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpS390XSRDconst {
			break
		}
		if x_2.AuxInt != j+32 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpS390XMOVDstoreidx)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(p)
		v.AddArg(idx)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMULLD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XMULLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XMULLDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLDload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMULLDconst_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		v.reset(OpS390XNEG)
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		if v.AuxInt != 1 {
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
		x := v.Args[0]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpS390XSLDconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c+1) && c >= 15) {
			break
		}
		v.reset(OpS390XSUB)
		v0 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-1) && c >= 17) {
			break
		}
		v.reset(OpS390XADD)
		v0 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c * d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMULLDload_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr1 := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFMOVDstore {
			break
		}
		if v_2.AuxInt != off {
			break
		}
		if v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XMULLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XMULLDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XMULLDload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XMULLW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XMULLWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpS390XMULLWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMULLWconst_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		v.reset(OpS390XNEGW)
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		if v.AuxInt != 1 {
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
		x := v.Args[0]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpS390XSLWconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c+1) && c >= 15) {
			break
		}
		v.reset(OpS390XSUBW)
		v0 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-1) && c >= 17) {
			break
		}
		v.reset(OpS390XADDW)
		v0 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int32(c * d))
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XMULLWload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XMULLWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XNEG_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = -c
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XADDconst {
			break
		}
		c := v_0.AuxInt
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XNEG {
			break
		}
		x := v_0_0.Args[0]
		if !(c != -(1 << 31)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = -c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XNEGW_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int32(-c))
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XNOT_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if !(true) {
			break
		}
		v.reset(OpS390XXOR)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDconst, typ.UInt64)
		v0.AuxInt = -1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XNOTW_0(v *Value) bool {

	for {
		x := v.Args[0]
		if !(true) {
			break
		}
		v.reset(OpS390XXORWconst)
		v.AuxInt = -1
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpS390XORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpS390XORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSLDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRDconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpS390XRLLGconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSRDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSLDconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpS390XRLLGconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		if v_0.AuxInt != -1<<63 {
			break
		}
		v_1 := v.Args[1]
		if v_1.Op != OpS390XLGDR {
			break
		}
		t := v_1.Type
		x := v_1.Args[0]
		v.reset(OpS390XLGDR)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XLNDFR, x.Type)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XLGDR {
			break
		}
		t := v_0.Type
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		if v_1.AuxInt != -1<<63 {
			break
		}
		v.reset(OpS390XLGDR)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XLNDFR, x.Type)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSLDconst {
			break
		}
		if v_0.AuxInt != 63 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XSRDconst {
			break
		}
		if v_0_0.AuxInt != 63 {
			break
		}
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpS390XLGDR {
			break
		}
		x := v_0_0_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XLGDR {
			break
		}
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpS390XLPDFR {
			break
		}
		t := v_1_0.Type
		y := v_1_0.Args[0]
		v.reset(OpS390XLGDR)
		v0 := b.NewValue0(v.Pos, OpS390XCPSDR, t)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XLGDR {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XLPDFR {
			break
		}
		t := v_0_0.Type
		y := v_0_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSLDconst {
			break
		}
		if v_1.AuxInt != 63 {
			break
		}
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpS390XSRDconst {
			break
		}
		if v_1_0.AuxInt != 63 {
			break
		}
		v_1_0_0 := v_1_0.Args[0]
		if v_1_0_0.Op != OpS390XLGDR {
			break
		}
		x := v_1_0_0.Args[0]
		v.reset(OpS390XLGDR)
		v0 := b.NewValue0(v.Pos, OpS390XCPSDR, t)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSLDconst {
			break
		}
		if v_0.AuxInt != 63 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XSRDconst {
			break
		}
		if v_0_0.AuxInt != 63 {
			break
		}
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpS390XLGDR {
			break
		}
		x := v_0_0_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c&-1<<63 == 0) {
			break
		}
		v.reset(OpS390XLGDR)
		v0 := b.NewValue0(v.Pos, OpS390XCPSDR, x.Type)
		v1 := b.NewValue0(v.Pos, OpS390XFMOVDconst, x.Type)
		v1.AuxInt = c
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSLDconst {
			break
		}
		if v_1.AuxInt != 63 {
			break
		}
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpS390XSRDconst {
			break
		}
		if v_1_0.AuxInt != 63 {
			break
		}
		v_1_0_0 := v_1_0.Args[0]
		if v_1_0_0.Op != OpS390XLGDR {
			break
		}
		x := v_1_0_0.Args[0]
		if !(c&-1<<63 == 0) {
			break
		}
		v.reset(OpS390XLGDR)
		v0 := b.NewValue0(v.Pos, OpS390XCPSDR, x.Type)
		v1 := b.NewValue0(v.Pos, OpS390XFMOVDconst, x.Type)
		v1.AuxInt = c
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XOR_10(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c | d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XMOVDconst)
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
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZload {
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
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVBZload {
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
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVHZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZload {
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
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
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
func rewriteValueS390X_OpS390XOR_20(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVHZload {
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
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVWZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVWZload {
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
		if !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVWZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVWZload {
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
		if !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDload, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
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
		y := or.Args[1]
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
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
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
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
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
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
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZload {
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
		y := or.Args[1]
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZload {
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
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZload {
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
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_30(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZload {
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
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_40(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVHZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVHZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVWZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVWZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDloadidx, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVWZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVWZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDloadidx, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVWZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVWZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDloadidx, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_50(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVWZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVWZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDloadidx, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVWZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVWZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDloadidx, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVWZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVWZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDloadidx, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVWZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVWZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDloadidx, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVWZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVWZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+4 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDloadidx, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_60(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_70(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_80(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0-16 && j1%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZload {
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
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := v.Args[1]
		if x0.Op != OpS390XMOVBZload {
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
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRload {
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
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWBRload, typ.UInt32)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_90(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRload {
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
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWBRload, typ.UInt32)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVWZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVWBRload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVWZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVWBRload {
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
		if !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDBRload, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVWZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVWBRload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVWZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVWBRload {
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
		if !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDBRload, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
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
		y := or.Args[1]
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
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
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
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
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
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
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRload {
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
		y := or.Args[1]
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRload, typ.UInt32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRload {
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
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRload, typ.UInt32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRload {
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
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRload, typ.UInt32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_100(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRload {
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
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRload, typ.UInt32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_110(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVWZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVWBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVWZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVWBRloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDBRloadidx, typ.Int64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVWZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVWBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVWZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVWBRloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDBRloadidx, typ.Int64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVWZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVWBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVWZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVWBRloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDBRloadidx, typ.Int64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_120(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVWZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVWBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVWZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVWBRloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDBRloadidx, typ.Int64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVWZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVWBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVWZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVWBRloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDBRloadidx, typ.Int64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVWZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVWBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVWZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVWBRloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDBRloadidx, typ.Int64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVWZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVWBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVWZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVWBRloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDBRloadidx, typ.Int64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLDconst {
			break
		}
		if sh.AuxInt != 32 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVWZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVWBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVWZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVWBRloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVDBRloadidx, typ.Int64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_130(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_140(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XOR_150(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XOR {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLDconst {
			break
		}
		j0 := s0.AuxInt
		r0 := s0.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLDconst {
			break
		}
		j1 := s1.AuxInt
		r1 := s1.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && j1 == j0+16 && j0%32 == 0 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XOR, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLDconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVWZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XORW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XORWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpS390XORWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSLWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRWconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpS390XRLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSRWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSLWconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpS390XRLLconst)
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
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XORW_10(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZload {
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
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVBZload {
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
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVHZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZload {
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
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVHZload {
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
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZload, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
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
		y := or.Args[1]
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
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
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
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
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORW_20(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
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
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZload, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORW_30(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVHZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVHZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		x0 := sh.Args[0]
		if x0.Op != OpS390XMOVHZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		x1 := v.Args[1]
		if x1.Op != OpS390XMOVHZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && p.Op != OpSB && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWZloadidx, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORW_40(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y := or.Args[1]
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s0 := v.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORW_50(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s1 := or.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		y := or.Args[1]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s1 := or.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		s0 := v.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+1 && j1 == j0-8 && j1%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j1
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZloadidx, typ.UInt16)
		v2.AuxInt = i0
		v2.Aux = s
		v2.AddArg(p)
		v2.AddArg(idx)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZload {
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
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := v.Args[1]
		if x0.Op != OpS390XMOVBZload {
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
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRload {
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
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWBRload, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRload {
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
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWBRload, typ.UInt32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
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
		y := or.Args[1]
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
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
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
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
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORW_60(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZload {
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
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRload, typ.UInt16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 8 {
			break
		}
		x1 := sh.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORW_70(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		r0 := v.Args[0]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		sh := v.Args[1]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		sh := v.Args[0]
		if sh.Op != OpS390XSLWconst {
			break
		}
		if sh.AuxInt != 16 {
			break
		}
		r1 := sh.Args[0]
		if r1.Op != OpS390XMOVHZreg {
			break
		}
		x1 := r1.Args[0]
		if x1.Op != OpS390XMOVHBRloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		r0 := v.Args[1]
		if r0.Op != OpS390XMOVHZreg {
			break
		}
		x0 := r0.Args[0]
		if x0.Op != OpS390XMOVHBRloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(i1 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && r0.Uses == 1 && r1.Uses == 1 && sh.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(r0) && clobber(r1) && clobber(sh)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWBRloadidx, typ.Int32)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = i0
		v0.Aux = s
		v0.AddArg(p)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORW_80(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y := or.Args[1]
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
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
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		or := v.Args[1]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if idx != x0.Args[0] {
			break
		}
		if p != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
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
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORW_90(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		s0 := or.Args[0]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		y := or.Args[1]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		or := v.Args[0]
		if or.Op != OpS390XORW {
			break
		}
		_ = or.Args[1]
		y := or.Args[0]
		s0 := or.Args[1]
		if s0.Op != OpS390XSLWconst {
			break
		}
		j0 := s0.AuxInt
		x0 := s0.Args[0]
		if x0.Op != OpS390XMOVBZloadidx {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s1 := v.Args[1]
		if s1.Op != OpS390XSLWconst {
			break
		}
		j1 := s1.AuxInt
		x1 := s1.Args[0]
		if x1.Op != OpS390XMOVBZloadidx {
			break
		}
		i1 := x1.AuxInt
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if idx != x1.Args[0] {
			break
		}
		if p != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(p.Op != OpSB && i1 == i0+1 && j1 == j0+8 && j0%16 == 0 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && or.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0) && clobber(s1) && clobber(or)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpS390XORW, v.Type)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpS390XSLWconst, v.Type)
		v1.AuxInt = j0
		v2 := b.NewValue0(v.Pos, OpS390XMOVHZreg, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpS390XMOVHBRloadidx, typ.Int16)
		v3.AuxInt = i0
		v3.Aux = s
		v3.AddArg(p)
		v3.AddArg(idx)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORWconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = -1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c | d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORWload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XORWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XORWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORconst_0(v *Value) bool {

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
		v.reset(OpS390XMOVDconst)
		v.AuxInt = -1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c | d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XORload_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr1 := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFMOVDstore {
			break
		}
		if v_2.AuxInt != off {
			break
		}
		if v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XOR)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XORload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XORload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSLD_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSLDconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1.Args[1]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_1.AuxInt
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSLD_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSLW_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSLWconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1.Args[1]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_1.AuxInt
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSLW_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSLW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRAD_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSRADconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1.Args[1]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_1.AuxInt
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRAD_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRADconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = d >> uint64(c)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRAW_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSRAWconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1.Args[1]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_1.AuxInt
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRAW_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRAW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRAWconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = int64(int32(d)) >> uint64(c)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRD_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSRDconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1.Args[1]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_1.AuxInt
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRD_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRDconst_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		if v.AuxInt != 1 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSLDconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpS390XLGDR {
			break
		}
		t := v_0_0.Type
		x := v_0_0.Args[0]
		v.reset(OpS390XLGDR)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpS390XLPDFR, x.Type)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRW_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSRWconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_0.AuxInt
		y := v_1.Args[1]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XAND {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1_1.AuxInt
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XANDWconst, typ.UInt32)
		v0.AuxInt = c & 63
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XANDWconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		if !(c&63 == 63) {
			break
		}
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVWZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVHZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSRW_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVBZreg {
			break
		}
		y := v_1.Args[0]
		v.reset(OpS390XSRW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSTM2_0(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		w2 := v.Args[1]
		w3 := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XSTM2 {
			break
		}
		if x.AuxInt != i-8 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		mem := x.Args[3]
		if !(x.Uses == 1 && is20Bit(i-8) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTM4)
		v.AuxInt = i - 8
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(w3)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRDconst {
			break
		}
		if v_1.AuxInt != 32 {
			break
		}
		x := v_1.Args[0]
		if x != v.Args[2] {
			break
		}
		mem := v.Args[3]
		v.reset(OpS390XMOVDstore)
		v.AuxInt = i
		v.Aux = s
		v.AddArg(p)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSTMG2_0(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		w2 := v.Args[1]
		w3 := v.Args[2]
		x := v.Args[3]
		if x.Op != OpS390XSTMG2 {
			break
		}
		if x.AuxInt != i-16 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[3]
		if p != x.Args[0] {
			break
		}
		w0 := x.Args[1]
		w1 := x.Args[2]
		mem := x.Args[3]
		if !(x.Uses == 1 && is20Bit(i-16) && clobber(x)) {
			break
		}
		v.reset(OpS390XSTMG4)
		v.AuxInt = i - 16
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(w1)
		v.AddArg(w2)
		v.AddArg(w3)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XSUB_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XSUBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpS390XNEG)
		v0 := b.NewValue0(v.Pos, OpS390XSUBconst, v.Type)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XSUBload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XSUBW_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XSUBWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpS390XNEGW)
		v0 := b.NewValue0(v.Pos, OpS390XSUBWconst, v.Type)
		v0.AuxInt = int64(int32(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XSUBWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XSUBWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSUBWconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		v.reset(OpS390XADDWconst)
		v.AuxInt = int64(int32(-c))
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpS390XSUBWload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XSUBWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XSUBWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSUBconst_0(v *Value) bool {

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
		x := v.Args[0]
		if !(c != -(1 << 31)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = -c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = d - c
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(-c - d)) {
			break
		}
		v.reset(OpS390XADDconst)
		v.AuxInt = -c - d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XSUBload_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr1 := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFMOVDstore {
			break
		}
		if v_2.AuxInt != off {
			break
		}
		if v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XSUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XSUBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XSUBload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XXOR_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpS390XXORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isU32Bit(c)) {
			break
		}
		v.reset(OpS390XXORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSLDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRDconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpS390XRLLGconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSRDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSLDconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 64-c) {
			break
		}
		v.reset(OpS390XRLLGconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c ^ d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c ^ d
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XXOR_10(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVDload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XXORW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpS390XXORWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpS390XXORWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSLWconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSRWconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpS390XRLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpS390XSRWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XSLWconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(OpS390XRLLconst)
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
		v.reset(OpS390XMOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueS390X_OpS390XXORW_10(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		g := v.Args[0]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		x := v.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		g := v.Args[1]
		if g.Op != OpS390XMOVWZload {
			break
		}
		off := g.AuxInt
		sym := g.Aux
		_ = g.Args[1]
		ptr := g.Args[0]
		mem := g.Args[1]
		if !(ptr.Op != OpSB && is20Bit(off) && psess.canMergeLoad(v, g, x) && clobber(g)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.Type = t
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XXORWconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c ^ d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XXORWload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XXORWload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XXORconst_0(v *Value) bool {

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
		if v_0.Op != OpS390XMOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpS390XMOVDconst)
		v.AuxInt = c ^ d
		return true
	}
	return false
}
func rewriteValueS390X_OpS390XXORload_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		ptr1 := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpS390XFMOVDstore {
			break
		}
		if v_2.AuxInt != off {
			break
		}
		if v_2.Aux != sym {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		y := v_2.Args[1]
		if !(isSamePtr(ptr1, ptr2)) {
			break
		}
		v.reset(OpS390XXOR)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpS390XLGDR, t)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XADDconst {
			break
		}
		off2 := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(off1+off2)) {
			break
		}
		v.reset(OpS390XXORload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		o1 := v.AuxInt
		s1 := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpS390XMOVDaddr {
			break
		}
		o2 := v_1.AuxInt
		s2 := v_1.Aux
		ptr := v_1.Args[0]
		mem := v.Args[2]
		if !(ptr.Op != OpSB && is20Bit(o1+o2) && canMergeSym(s1, s2)) {
			break
		}
		v.reset(OpS390XXORload)
		v.AuxInt = o1 + o2
		v.Aux = mergeSym(s1, s2)
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpSelect0_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpS390XAddTupleFirst32 {
			break
		}
		_ = v_0.Args[1]
		val := v_0.Args[0]
		tuple := v_0.Args[1]
		v.reset(OpS390XADDW)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpSelect0, t)
		v0.AddArg(tuple)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpS390XAddTupleFirst64 {
			break
		}
		_ = v_0.Args[1]
		val := v_0.Args[0]
		tuple := v_0.Args[1]
		v.reset(OpS390XADD)
		v.AddArg(val)
		v0 := b.NewValue0(v.Pos, OpSelect0, t)
		v0.AddArg(tuple)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueS390X_OpSelect1_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XAddTupleFirst32 {
			break
		}
		_ = v_0.Args[1]
		tuple := v_0.Args[1]
		v.reset(OpSelect1)
		v.AddArg(tuple)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpS390XAddTupleFirst64 {
			break
		}
		_ = v_0.Args[1]
		tuple := v_0.Args[1]
		v.reset(OpSelect1)
		v.AddArg(tuple)
		return true
	}
	return false
}
func rewriteValueS390X_OpSignExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpSignExt16to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpSignExt32to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVWreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpSignExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpSignExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpSignExt8to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpSlicemask_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpS390XSRADconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpS390XNEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueS390X_OpSqrt_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XFSQRT)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpStaticCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpS390XCALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpStore_0(v *Value) bool {

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 8 && psess.is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpS390XFMOVDstore)
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
		v.reset(OpS390XFMOVSstore)
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
		if !(t.(*types.Type).Size(psess.types) == 8) {
			break
		}
		v.reset(OpS390XMOVDstore)
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
		if !(t.(*types.Type).Size(psess.types) == 4) {
			break
		}
		v.reset(OpS390XMOVWstore)
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
		v.reset(OpS390XMOVHstore)
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
		v.reset(OpS390XMOVBstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpSub16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSUBW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpSub32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSUBW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpSub32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XFSUBS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpSub64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpSub64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XFSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpSub8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSUBW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpSubPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpTrunc_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XFIDBR)
		v.AuxInt = 5
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpTrunc16to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpTrunc32to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpTrunc32to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpTrunc64to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpTrunc64to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpTrunc64to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpWB_0(v *Value) bool {

	for {
		fn := v.Aux
		_ = v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpS390XLoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueS390X_OpXor16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XXORW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpXor32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XXORW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpXor64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueS390X_OpXor8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpS390XXORW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueS390X_OpZero_0(v *Value) bool {
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
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = 0
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
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = 0
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpS390XMOVWstoreconst)
		v.AuxInt = 0
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
		v.reset(OpS390XMOVDstoreconst)
		v.AuxInt = 0
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
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = makeValAndOff(0, 2)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpS390XMOVHstoreconst, psess.types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 5 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpS390XMOVBstoreconst)
		v.AuxInt = makeValAndOff(0, 4)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWstoreconst, psess.types.TypeMem)
		v0.AuxInt = 0
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
		v.reset(OpS390XMOVHstoreconst)
		v.AuxInt = makeValAndOff(0, 4)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWstoreconst, psess.types.TypeMem)
		v0.AuxInt = 0
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
		v.reset(OpS390XMOVWstoreconst)
		v.AuxInt = makeValAndOff(0, 3)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpS390XMOVWstoreconst, psess.types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		if !(s > 0 && s <= 1024) {
			break
		}
		v.reset(OpS390XCLEAR)
		v.AuxInt = makeValAndOff(s, 0)
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpZero_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		s := v.AuxInt
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		if !(s > 1024) {
			break
		}
		v.reset(OpS390XLoweredZero)
		v.AuxInt = s % 256
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpS390XADDconst, destptr.Type)
		v0.AuxInt = (s / 256) * 256
		v0.AddArg(destptr)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueS390X_OpZeroExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpZeroExt16to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVHZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpZeroExt32to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVWZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpZeroExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpZeroExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueS390X_OpZeroExt8to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpS390XMOVBZreg)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteBlockS390X(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	case BlockS390XEQ:

		for {
			v := b.Control
			if v.Op != OpS390XInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockS390XEQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagLT {
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
			if v.Op != OpS390XFlagGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case BlockS390XGE:

		for {
			v := b.Control
			if v.Op != OpS390XInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockS390XLE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagLT {
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
			if v.Op != OpS390XFlagGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
	case BlockS390XGT:

		for {
			v := b.Control
			if v.Op != OpS390XInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockS390XLT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagEQ {
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
			if v.Op != OpS390XFlagLT {
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
			if v.Op != OpS390XFlagGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
	case BlockIf:

		for {
			v := b.Control
			if v.Op != OpS390XMOVDLT {
				break
			}
			_ = v.Args[2]
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0.AuxInt != 0 {
				break
			}
			v_1 := v.Args[1]
			if v_1.Op != OpS390XMOVDconst {
				break
			}
			if v_1.AuxInt != 1 {
				break
			}
			cmp := v.Args[2]
			b.Kind = BlockS390XLT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XMOVDLE {
				break
			}
			_ = v.Args[2]
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0.AuxInt != 0 {
				break
			}
			v_1 := v.Args[1]
			if v_1.Op != OpS390XMOVDconst {
				break
			}
			if v_1.AuxInt != 1 {
				break
			}
			cmp := v.Args[2]
			b.Kind = BlockS390XLE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XMOVDGT {
				break
			}
			_ = v.Args[2]
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0.AuxInt != 0 {
				break
			}
			v_1 := v.Args[1]
			if v_1.Op != OpS390XMOVDconst {
				break
			}
			if v_1.AuxInt != 1 {
				break
			}
			cmp := v.Args[2]
			b.Kind = BlockS390XGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XMOVDGE {
				break
			}
			_ = v.Args[2]
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0.AuxInt != 0 {
				break
			}
			v_1 := v.Args[1]
			if v_1.Op != OpS390XMOVDconst {
				break
			}
			if v_1.AuxInt != 1 {
				break
			}
			cmp := v.Args[2]
			b.Kind = BlockS390XGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XMOVDEQ {
				break
			}
			_ = v.Args[2]
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0.AuxInt != 0 {
				break
			}
			v_1 := v.Args[1]
			if v_1.Op != OpS390XMOVDconst {
				break
			}
			if v_1.AuxInt != 1 {
				break
			}
			cmp := v.Args[2]
			b.Kind = BlockS390XEQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XMOVDNE {
				break
			}
			_ = v.Args[2]
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0.AuxInt != 0 {
				break
			}
			v_1 := v.Args[1]
			if v_1.Op != OpS390XMOVDconst {
				break
			}
			if v_1.AuxInt != 1 {
				break
			}
			cmp := v.Args[2]
			b.Kind = BlockS390XNE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XMOVDGTnoinv {
				break
			}
			_ = v.Args[2]
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0.AuxInt != 0 {
				break
			}
			v_1 := v.Args[1]
			if v_1.Op != OpS390XMOVDconst {
				break
			}
			if v_1.AuxInt != 1 {
				break
			}
			cmp := v.Args[2]
			b.Kind = BlockS390XGTF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XMOVDGEnoinv {
				break
			}
			_ = v.Args[2]
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0.AuxInt != 0 {
				break
			}
			v_1 := v.Args[1]
			if v_1.Op != OpS390XMOVDconst {
				break
			}
			if v_1.AuxInt != 1 {
				break
			}
			cmp := v.Args[2]
			b.Kind = BlockS390XGEF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			_ = v
			cond := b.Control
			b.Kind = BlockS390XNE
			v0 := b.NewValue0(v.Pos, OpS390XCMPWconst, psess.types.TypeFlags)
			v0.AuxInt = 0
			v1 := b.NewValue0(v.Pos, OpS390XMOVBZreg, typ.Bool)
			v1.AddArg(cond)
			v0.AddArg(v1)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockS390XLE:

		for {
			v := b.Control
			if v.Op != OpS390XInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockS390XGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagLT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case BlockS390XLT:

		for {
			v := b.Control
			if v.Op != OpS390XInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockS390XGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagEQ {
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
			if v.Op != OpS390XFlagLT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case BlockS390XNE:

		for {
			v := b.Control
			if v.Op != OpS390XCMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDLT {
				break
			}
			_ = v_0.Args[2]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0_0.AuxInt != 0 {
				break
			}
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpS390XMOVDconst {
				break
			}
			if v_0_1.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[2]
			b.Kind = BlockS390XLT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XCMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDLE {
				break
			}
			_ = v_0.Args[2]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0_0.AuxInt != 0 {
				break
			}
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpS390XMOVDconst {
				break
			}
			if v_0_1.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[2]
			b.Kind = BlockS390XLE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XCMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDGT {
				break
			}
			_ = v_0.Args[2]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0_0.AuxInt != 0 {
				break
			}
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpS390XMOVDconst {
				break
			}
			if v_0_1.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[2]
			b.Kind = BlockS390XGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XCMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDGE {
				break
			}
			_ = v_0.Args[2]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0_0.AuxInt != 0 {
				break
			}
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpS390XMOVDconst {
				break
			}
			if v_0_1.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[2]
			b.Kind = BlockS390XGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XCMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDEQ {
				break
			}
			_ = v_0.Args[2]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0_0.AuxInt != 0 {
				break
			}
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpS390XMOVDconst {
				break
			}
			if v_0_1.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[2]
			b.Kind = BlockS390XEQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XCMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDNE {
				break
			}
			_ = v_0.Args[2]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0_0.AuxInt != 0 {
				break
			}
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpS390XMOVDconst {
				break
			}
			if v_0_1.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[2]
			b.Kind = BlockS390XNE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XCMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDGTnoinv {
				break
			}
			_ = v_0.Args[2]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0_0.AuxInt != 0 {
				break
			}
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpS390XMOVDconst {
				break
			}
			if v_0_1.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[2]
			b.Kind = BlockS390XGTF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XCMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpS390XMOVDGEnoinv {
				break
			}
			_ = v_0.Args[2]
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpS390XMOVDconst {
				break
			}
			if v_0_0.AuxInt != 0 {
				break
			}
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpS390XMOVDconst {
				break
			}
			if v_0_1.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[2]
			b.Kind = BlockS390XGEF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockS390XNE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagEQ {
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
			if v.Op != OpS390XFlagLT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpS390XFlagGT {
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

package ssa

import "github.com/dave/golib/src/cmd/compile/internal/types"

// in case not otherwise used
// in case not otherwise used
// in case not otherwise used
// in case not otherwise used

func (psess *PackageSession) rewriteValueARM64(v *Value) bool {
	switch v.Op {
	case OpARM64ADD:
		return rewriteValueARM64_OpARM64ADD_0(v)
	case OpARM64ADDconst:
		return rewriteValueARM64_OpARM64ADDconst_0(v)
	case OpARM64ADDshiftLL:
		return psess.rewriteValueARM64_OpARM64ADDshiftLL_0(v)
	case OpARM64ADDshiftRA:
		return rewriteValueARM64_OpARM64ADDshiftRA_0(v)
	case OpARM64ADDshiftRL:
		return psess.rewriteValueARM64_OpARM64ADDshiftRL_0(v)
	case OpARM64AND:
		return rewriteValueARM64_OpARM64AND_0(v) || rewriteValueARM64_OpARM64AND_10(v)
	case OpARM64ANDconst:
		return rewriteValueARM64_OpARM64ANDconst_0(v)
	case OpARM64ANDshiftLL:
		return rewriteValueARM64_OpARM64ANDshiftLL_0(v)
	case OpARM64ANDshiftRA:
		return rewriteValueARM64_OpARM64ANDshiftRA_0(v)
	case OpARM64ANDshiftRL:
		return rewriteValueARM64_OpARM64ANDshiftRL_0(v)
	case OpARM64BIC:
		return rewriteValueARM64_OpARM64BIC_0(v)
	case OpARM64BICshiftLL:
		return rewriteValueARM64_OpARM64BICshiftLL_0(v)
	case OpARM64BICshiftRA:
		return rewriteValueARM64_OpARM64BICshiftRA_0(v)
	case OpARM64BICshiftRL:
		return rewriteValueARM64_OpARM64BICshiftRL_0(v)
	case OpARM64CMN:
		return rewriteValueARM64_OpARM64CMN_0(v)
	case OpARM64CMNWconst:
		return rewriteValueARM64_OpARM64CMNWconst_0(v)
	case OpARM64CMNconst:
		return rewriteValueARM64_OpARM64CMNconst_0(v)
	case OpARM64CMP:
		return psess.rewriteValueARM64_OpARM64CMP_0(v)
	case OpARM64CMPW:
		return psess.rewriteValueARM64_OpARM64CMPW_0(v)
	case OpARM64CMPWconst:
		return rewriteValueARM64_OpARM64CMPWconst_0(v)
	case OpARM64CMPconst:
		return rewriteValueARM64_OpARM64CMPconst_0(v)
	case OpARM64CMPshiftLL:
		return psess.rewriteValueARM64_OpARM64CMPshiftLL_0(v)
	case OpARM64CMPshiftRA:
		return psess.rewriteValueARM64_OpARM64CMPshiftRA_0(v)
	case OpARM64CMPshiftRL:
		return psess.rewriteValueARM64_OpARM64CMPshiftRL_0(v)
	case OpARM64CSEL:
		return psess.rewriteValueARM64_OpARM64CSEL_0(v)
	case OpARM64CSEL0:
		return psess.rewriteValueARM64_OpARM64CSEL0_0(v)
	case OpARM64DIV:
		return rewriteValueARM64_OpARM64DIV_0(v)
	case OpARM64DIVW:
		return rewriteValueARM64_OpARM64DIVW_0(v)
	case OpARM64EON:
		return rewriteValueARM64_OpARM64EON_0(v)
	case OpARM64EONshiftLL:
		return rewriteValueARM64_OpARM64EONshiftLL_0(v)
	case OpARM64EONshiftRA:
		return rewriteValueARM64_OpARM64EONshiftRA_0(v)
	case OpARM64EONshiftRL:
		return rewriteValueARM64_OpARM64EONshiftRL_0(v)
	case OpARM64Equal:
		return rewriteValueARM64_OpARM64Equal_0(v)
	case OpARM64FADDD:
		return rewriteValueARM64_OpARM64FADDD_0(v)
	case OpARM64FADDS:
		return rewriteValueARM64_OpARM64FADDS_0(v)
	case OpARM64FMOVDgpfp:
		return rewriteValueARM64_OpARM64FMOVDgpfp_0(v)
	case OpARM64FMOVDload:
		return rewriteValueARM64_OpARM64FMOVDload_0(v)
	case OpARM64FMOVDstore:
		return rewriteValueARM64_OpARM64FMOVDstore_0(v)
	case OpARM64FMOVSload:
		return rewriteValueARM64_OpARM64FMOVSload_0(v)
	case OpARM64FMOVSstore:
		return rewriteValueARM64_OpARM64FMOVSstore_0(v)
	case OpARM64FMULD:
		return rewriteValueARM64_OpARM64FMULD_0(v)
	case OpARM64FMULS:
		return rewriteValueARM64_OpARM64FMULS_0(v)
	case OpARM64FNEGD:
		return rewriteValueARM64_OpARM64FNEGD_0(v)
	case OpARM64FNEGS:
		return rewriteValueARM64_OpARM64FNEGS_0(v)
	case OpARM64FNMULD:
		return rewriteValueARM64_OpARM64FNMULD_0(v)
	case OpARM64FNMULS:
		return rewriteValueARM64_OpARM64FNMULS_0(v)
	case OpARM64FSUBD:
		return rewriteValueARM64_OpARM64FSUBD_0(v)
	case OpARM64FSUBS:
		return rewriteValueARM64_OpARM64FSUBS_0(v)
	case OpARM64GreaterEqual:
		return rewriteValueARM64_OpARM64GreaterEqual_0(v)
	case OpARM64GreaterEqualU:
		return rewriteValueARM64_OpARM64GreaterEqualU_0(v)
	case OpARM64GreaterThan:
		return rewriteValueARM64_OpARM64GreaterThan_0(v)
	case OpARM64GreaterThanU:
		return rewriteValueARM64_OpARM64GreaterThanU_0(v)
	case OpARM64LessEqual:
		return rewriteValueARM64_OpARM64LessEqual_0(v)
	case OpARM64LessEqualU:
		return rewriteValueARM64_OpARM64LessEqualU_0(v)
	case OpARM64LessThan:
		return rewriteValueARM64_OpARM64LessThan_0(v)
	case OpARM64LessThanU:
		return rewriteValueARM64_OpARM64LessThanU_0(v)
	case OpARM64MNEG:
		return rewriteValueARM64_OpARM64MNEG_0(v) || rewriteValueARM64_OpARM64MNEG_10(v) || rewriteValueARM64_OpARM64MNEG_20(v)
	case OpARM64MNEGW:
		return rewriteValueARM64_OpARM64MNEGW_0(v) || rewriteValueARM64_OpARM64MNEGW_10(v) || rewriteValueARM64_OpARM64MNEGW_20(v)
	case OpARM64MOD:
		return rewriteValueARM64_OpARM64MOD_0(v)
	case OpARM64MODW:
		return rewriteValueARM64_OpARM64MODW_0(v)
	case OpARM64MOVBUload:
		return rewriteValueARM64_OpARM64MOVBUload_0(v)
	case OpARM64MOVBUloadidx:
		return rewriteValueARM64_OpARM64MOVBUloadidx_0(v)
	case OpARM64MOVBUreg:
		return rewriteValueARM64_OpARM64MOVBUreg_0(v)
	case OpARM64MOVBload:
		return rewriteValueARM64_OpARM64MOVBload_0(v)
	case OpARM64MOVBloadidx:
		return rewriteValueARM64_OpARM64MOVBloadidx_0(v)
	case OpARM64MOVBreg:
		return rewriteValueARM64_OpARM64MOVBreg_0(v)
	case OpARM64MOVBstore:
		return rewriteValueARM64_OpARM64MOVBstore_0(v) || rewriteValueARM64_OpARM64MOVBstore_10(v) || rewriteValueARM64_OpARM64MOVBstore_20(v) || rewriteValueARM64_OpARM64MOVBstore_30(v) || rewriteValueARM64_OpARM64MOVBstore_40(v)
	case OpARM64MOVBstoreidx:
		return rewriteValueARM64_OpARM64MOVBstoreidx_0(v) || rewriteValueARM64_OpARM64MOVBstoreidx_10(v)
	case OpARM64MOVBstorezero:
		return rewriteValueARM64_OpARM64MOVBstorezero_0(v)
	case OpARM64MOVBstorezeroidx:
		return rewriteValueARM64_OpARM64MOVBstorezeroidx_0(v)
	case OpARM64MOVDload:
		return rewriteValueARM64_OpARM64MOVDload_0(v)
	case OpARM64MOVDloadidx:
		return rewriteValueARM64_OpARM64MOVDloadidx_0(v)
	case OpARM64MOVDloadidx8:
		return rewriteValueARM64_OpARM64MOVDloadidx8_0(v)
	case OpARM64MOVDreg:
		return rewriteValueARM64_OpARM64MOVDreg_0(v)
	case OpARM64MOVDstore:
		return rewriteValueARM64_OpARM64MOVDstore_0(v)
	case OpARM64MOVDstoreidx:
		return rewriteValueARM64_OpARM64MOVDstoreidx_0(v)
	case OpARM64MOVDstoreidx8:
		return rewriteValueARM64_OpARM64MOVDstoreidx8_0(v)
	case OpARM64MOVDstorezero:
		return rewriteValueARM64_OpARM64MOVDstorezero_0(v)
	case OpARM64MOVDstorezeroidx:
		return rewriteValueARM64_OpARM64MOVDstorezeroidx_0(v)
	case OpARM64MOVDstorezeroidx8:
		return rewriteValueARM64_OpARM64MOVDstorezeroidx8_0(v)
	case OpARM64MOVHUload:
		return rewriteValueARM64_OpARM64MOVHUload_0(v)
	case OpARM64MOVHUloadidx:
		return rewriteValueARM64_OpARM64MOVHUloadidx_0(v)
	case OpARM64MOVHUloadidx2:
		return rewriteValueARM64_OpARM64MOVHUloadidx2_0(v)
	case OpARM64MOVHUreg:
		return rewriteValueARM64_OpARM64MOVHUreg_0(v) || rewriteValueARM64_OpARM64MOVHUreg_10(v)
	case OpARM64MOVHload:
		return rewriteValueARM64_OpARM64MOVHload_0(v)
	case OpARM64MOVHloadidx:
		return rewriteValueARM64_OpARM64MOVHloadidx_0(v)
	case OpARM64MOVHloadidx2:
		return rewriteValueARM64_OpARM64MOVHloadidx2_0(v)
	case OpARM64MOVHreg:
		return rewriteValueARM64_OpARM64MOVHreg_0(v) || rewriteValueARM64_OpARM64MOVHreg_10(v)
	case OpARM64MOVHstore:
		return rewriteValueARM64_OpARM64MOVHstore_0(v) || rewriteValueARM64_OpARM64MOVHstore_10(v) || rewriteValueARM64_OpARM64MOVHstore_20(v)
	case OpARM64MOVHstoreidx:
		return rewriteValueARM64_OpARM64MOVHstoreidx_0(v) || rewriteValueARM64_OpARM64MOVHstoreidx_10(v)
	case OpARM64MOVHstoreidx2:
		return rewriteValueARM64_OpARM64MOVHstoreidx2_0(v)
	case OpARM64MOVHstorezero:
		return rewriteValueARM64_OpARM64MOVHstorezero_0(v)
	case OpARM64MOVHstorezeroidx:
		return rewriteValueARM64_OpARM64MOVHstorezeroidx_0(v)
	case OpARM64MOVHstorezeroidx2:
		return rewriteValueARM64_OpARM64MOVHstorezeroidx2_0(v)
	case OpARM64MOVQstorezero:
		return rewriteValueARM64_OpARM64MOVQstorezero_0(v)
	case OpARM64MOVWUload:
		return rewriteValueARM64_OpARM64MOVWUload_0(v)
	case OpARM64MOVWUloadidx:
		return rewriteValueARM64_OpARM64MOVWUloadidx_0(v)
	case OpARM64MOVWUloadidx4:
		return rewriteValueARM64_OpARM64MOVWUloadidx4_0(v)
	case OpARM64MOVWUreg:
		return rewriteValueARM64_OpARM64MOVWUreg_0(v) || rewriteValueARM64_OpARM64MOVWUreg_10(v)
	case OpARM64MOVWload:
		return rewriteValueARM64_OpARM64MOVWload_0(v)
	case OpARM64MOVWloadidx:
		return rewriteValueARM64_OpARM64MOVWloadidx_0(v)
	case OpARM64MOVWloadidx4:
		return rewriteValueARM64_OpARM64MOVWloadidx4_0(v)
	case OpARM64MOVWreg:
		return rewriteValueARM64_OpARM64MOVWreg_0(v) || rewriteValueARM64_OpARM64MOVWreg_10(v)
	case OpARM64MOVWstore:
		return rewriteValueARM64_OpARM64MOVWstore_0(v) || rewriteValueARM64_OpARM64MOVWstore_10(v)
	case OpARM64MOVWstoreidx:
		return rewriteValueARM64_OpARM64MOVWstoreidx_0(v)
	case OpARM64MOVWstoreidx4:
		return rewriteValueARM64_OpARM64MOVWstoreidx4_0(v)
	case OpARM64MOVWstorezero:
		return rewriteValueARM64_OpARM64MOVWstorezero_0(v)
	case OpARM64MOVWstorezeroidx:
		return rewriteValueARM64_OpARM64MOVWstorezeroidx_0(v)
	case OpARM64MOVWstorezeroidx4:
		return rewriteValueARM64_OpARM64MOVWstorezeroidx4_0(v)
	case OpARM64MUL:
		return rewriteValueARM64_OpARM64MUL_0(v) || rewriteValueARM64_OpARM64MUL_10(v) || rewriteValueARM64_OpARM64MUL_20(v)
	case OpARM64MULW:
		return rewriteValueARM64_OpARM64MULW_0(v) || rewriteValueARM64_OpARM64MULW_10(v) || rewriteValueARM64_OpARM64MULW_20(v)
	case OpARM64MVN:
		return rewriteValueARM64_OpARM64MVN_0(v)
	case OpARM64NEG:
		return rewriteValueARM64_OpARM64NEG_0(v)
	case OpARM64NotEqual:
		return rewriteValueARM64_OpARM64NotEqual_0(v)
	case OpARM64OR:
		return rewriteValueARM64_OpARM64OR_0(v) || rewriteValueARM64_OpARM64OR_10(v) || rewriteValueARM64_OpARM64OR_20(v) || rewriteValueARM64_OpARM64OR_30(v)
	case OpARM64ORN:
		return rewriteValueARM64_OpARM64ORN_0(v)
	case OpARM64ORNshiftLL:
		return rewriteValueARM64_OpARM64ORNshiftLL_0(v)
	case OpARM64ORNshiftRA:
		return rewriteValueARM64_OpARM64ORNshiftRA_0(v)
	case OpARM64ORNshiftRL:
		return rewriteValueARM64_OpARM64ORNshiftRL_0(v)
	case OpARM64ORconst:
		return rewriteValueARM64_OpARM64ORconst_0(v)
	case OpARM64ORshiftLL:
		return psess.rewriteValueARM64_OpARM64ORshiftLL_0(v) || rewriteValueARM64_OpARM64ORshiftLL_10(v) || rewriteValueARM64_OpARM64ORshiftLL_20(v)
	case OpARM64ORshiftRA:
		return rewriteValueARM64_OpARM64ORshiftRA_0(v)
	case OpARM64ORshiftRL:
		return psess.rewriteValueARM64_OpARM64ORshiftRL_0(v)
	case OpARM64SLL:
		return rewriteValueARM64_OpARM64SLL_0(v)
	case OpARM64SLLconst:
		return rewriteValueARM64_OpARM64SLLconst_0(v)
	case OpARM64SRA:
		return rewriteValueARM64_OpARM64SRA_0(v)
	case OpARM64SRAconst:
		return rewriteValueARM64_OpARM64SRAconst_0(v)
	case OpARM64SRL:
		return rewriteValueARM64_OpARM64SRL_0(v)
	case OpARM64SRLconst:
		return rewriteValueARM64_OpARM64SRLconst_0(v) || rewriteValueARM64_OpARM64SRLconst_10(v)
	case OpARM64STP:
		return rewriteValueARM64_OpARM64STP_0(v)
	case OpARM64SUB:
		return rewriteValueARM64_OpARM64SUB_0(v)
	case OpARM64SUBconst:
		return rewriteValueARM64_OpARM64SUBconst_0(v)
	case OpARM64SUBshiftLL:
		return rewriteValueARM64_OpARM64SUBshiftLL_0(v)
	case OpARM64SUBshiftRA:
		return rewriteValueARM64_OpARM64SUBshiftRA_0(v)
	case OpARM64SUBshiftRL:
		return rewriteValueARM64_OpARM64SUBshiftRL_0(v)
	case OpARM64TST:
		return rewriteValueARM64_OpARM64TST_0(v)
	case OpARM64TSTWconst:
		return rewriteValueARM64_OpARM64TSTWconst_0(v)
	case OpARM64TSTconst:
		return rewriteValueARM64_OpARM64TSTconst_0(v)
	case OpARM64UBFIZ:
		return rewriteValueARM64_OpARM64UBFIZ_0(v)
	case OpARM64UBFX:
		return rewriteValueARM64_OpARM64UBFX_0(v)
	case OpARM64UDIV:
		return rewriteValueARM64_OpARM64UDIV_0(v)
	case OpARM64UDIVW:
		return rewriteValueARM64_OpARM64UDIVW_0(v)
	case OpARM64UMOD:
		return rewriteValueARM64_OpARM64UMOD_0(v)
	case OpARM64UMODW:
		return rewriteValueARM64_OpARM64UMODW_0(v)
	case OpARM64XOR:
		return rewriteValueARM64_OpARM64XOR_0(v) || rewriteValueARM64_OpARM64XOR_10(v)
	case OpARM64XORconst:
		return rewriteValueARM64_OpARM64XORconst_0(v)
	case OpARM64XORshiftLL:
		return psess.rewriteValueARM64_OpARM64XORshiftLL_0(v)
	case OpARM64XORshiftRA:
		return rewriteValueARM64_OpARM64XORshiftRA_0(v)
	case OpARM64XORshiftRL:
		return psess.rewriteValueARM64_OpARM64XORshiftRL_0(v)
	case OpAdd16:
		return rewriteValueARM64_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValueARM64_OpAdd32_0(v)
	case OpAdd32F:
		return rewriteValueARM64_OpAdd32F_0(v)
	case OpAdd64:
		return rewriteValueARM64_OpAdd64_0(v)
	case OpAdd64F:
		return rewriteValueARM64_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValueARM64_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValueARM64_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValueARM64_OpAddr_0(v)
	case OpAnd16:
		return rewriteValueARM64_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValueARM64_OpAnd32_0(v)
	case OpAnd64:
		return rewriteValueARM64_OpAnd64_0(v)
	case OpAnd8:
		return rewriteValueARM64_OpAnd8_0(v)
	case OpAndB:
		return rewriteValueARM64_OpAndB_0(v)
	case OpAtomicAdd32:
		return rewriteValueARM64_OpAtomicAdd32_0(v)
	case OpAtomicAdd32Variant:
		return rewriteValueARM64_OpAtomicAdd32Variant_0(v)
	case OpAtomicAdd64:
		return rewriteValueARM64_OpAtomicAdd64_0(v)
	case OpAtomicAdd64Variant:
		return rewriteValueARM64_OpAtomicAdd64Variant_0(v)
	case OpAtomicAnd8:
		return psess.rewriteValueARM64_OpAtomicAnd8_0(v)
	case OpAtomicCompareAndSwap32:
		return rewriteValueARM64_OpAtomicCompareAndSwap32_0(v)
	case OpAtomicCompareAndSwap64:
		return rewriteValueARM64_OpAtomicCompareAndSwap64_0(v)
	case OpAtomicExchange32:
		return rewriteValueARM64_OpAtomicExchange32_0(v)
	case OpAtomicExchange64:
		return rewriteValueARM64_OpAtomicExchange64_0(v)
	case OpAtomicLoad32:
		return rewriteValueARM64_OpAtomicLoad32_0(v)
	case OpAtomicLoad64:
		return rewriteValueARM64_OpAtomicLoad64_0(v)
	case OpAtomicLoadPtr:
		return rewriteValueARM64_OpAtomicLoadPtr_0(v)
	case OpAtomicOr8:
		return psess.rewriteValueARM64_OpAtomicOr8_0(v)
	case OpAtomicStore32:
		return rewriteValueARM64_OpAtomicStore32_0(v)
	case OpAtomicStore64:
		return rewriteValueARM64_OpAtomicStore64_0(v)
	case OpAtomicStorePtrNoWB:
		return rewriteValueARM64_OpAtomicStorePtrNoWB_0(v)
	case OpAvg64u:
		return rewriteValueARM64_OpAvg64u_0(v)
	case OpBitLen64:
		return rewriteValueARM64_OpBitLen64_0(v)
	case OpBitRev16:
		return rewriteValueARM64_OpBitRev16_0(v)
	case OpBitRev32:
		return rewriteValueARM64_OpBitRev32_0(v)
	case OpBitRev64:
		return rewriteValueARM64_OpBitRev64_0(v)
	case OpBitRev8:
		return rewriteValueARM64_OpBitRev8_0(v)
	case OpBswap32:
		return rewriteValueARM64_OpBswap32_0(v)
	case OpBswap64:
		return rewriteValueARM64_OpBswap64_0(v)
	case OpCeil:
		return rewriteValueARM64_OpCeil_0(v)
	case OpClosureCall:
		return rewriteValueARM64_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValueARM64_OpCom16_0(v)
	case OpCom32:
		return rewriteValueARM64_OpCom32_0(v)
	case OpCom64:
		return rewriteValueARM64_OpCom64_0(v)
	case OpCom8:
		return rewriteValueARM64_OpCom8_0(v)
	case OpCondSelect:
		return psess.rewriteValueARM64_OpCondSelect_0(v)
	case OpConst16:
		return rewriteValueARM64_OpConst16_0(v)
	case OpConst32:
		return rewriteValueARM64_OpConst32_0(v)
	case OpConst32F:
		return rewriteValueARM64_OpConst32F_0(v)
	case OpConst64:
		return rewriteValueARM64_OpConst64_0(v)
	case OpConst64F:
		return rewriteValueARM64_OpConst64F_0(v)
	case OpConst8:
		return rewriteValueARM64_OpConst8_0(v)
	case OpConstBool:
		return rewriteValueARM64_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValueARM64_OpConstNil_0(v)
	case OpCtz32:
		return rewriteValueARM64_OpCtz32_0(v)
	case OpCtz32NonZero:
		return rewriteValueARM64_OpCtz32NonZero_0(v)
	case OpCtz64:
		return rewriteValueARM64_OpCtz64_0(v)
	case OpCtz64NonZero:
		return rewriteValueARM64_OpCtz64NonZero_0(v)
	case OpCvt32Fto32:
		return rewriteValueARM64_OpCvt32Fto32_0(v)
	case OpCvt32Fto32U:
		return rewriteValueARM64_OpCvt32Fto32U_0(v)
	case OpCvt32Fto64:
		return rewriteValueARM64_OpCvt32Fto64_0(v)
	case OpCvt32Fto64F:
		return rewriteValueARM64_OpCvt32Fto64F_0(v)
	case OpCvt32Fto64U:
		return rewriteValueARM64_OpCvt32Fto64U_0(v)
	case OpCvt32Uto32F:
		return rewriteValueARM64_OpCvt32Uto32F_0(v)
	case OpCvt32Uto64F:
		return rewriteValueARM64_OpCvt32Uto64F_0(v)
	case OpCvt32to32F:
		return rewriteValueARM64_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValueARM64_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValueARM64_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValueARM64_OpCvt64Fto32F_0(v)
	case OpCvt64Fto32U:
		return rewriteValueARM64_OpCvt64Fto32U_0(v)
	case OpCvt64Fto64:
		return rewriteValueARM64_OpCvt64Fto64_0(v)
	case OpCvt64Fto64U:
		return rewriteValueARM64_OpCvt64Fto64U_0(v)
	case OpCvt64Uto32F:
		return rewriteValueARM64_OpCvt64Uto32F_0(v)
	case OpCvt64Uto64F:
		return rewriteValueARM64_OpCvt64Uto64F_0(v)
	case OpCvt64to32F:
		return rewriteValueARM64_OpCvt64to32F_0(v)
	case OpCvt64to64F:
		return rewriteValueARM64_OpCvt64to64F_0(v)
	case OpDiv16:
		return rewriteValueARM64_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValueARM64_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValueARM64_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValueARM64_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValueARM64_OpDiv32u_0(v)
	case OpDiv64:
		return rewriteValueARM64_OpDiv64_0(v)
	case OpDiv64F:
		return rewriteValueARM64_OpDiv64F_0(v)
	case OpDiv64u:
		return rewriteValueARM64_OpDiv64u_0(v)
	case OpDiv8:
		return rewriteValueARM64_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValueARM64_OpDiv8u_0(v)
	case OpEq16:
		return psess.rewriteValueARM64_OpEq16_0(v)
	case OpEq32:
		return psess.rewriteValueARM64_OpEq32_0(v)
	case OpEq32F:
		return psess.rewriteValueARM64_OpEq32F_0(v)
	case OpEq64:
		return psess.rewriteValueARM64_OpEq64_0(v)
	case OpEq64F:
		return psess.rewriteValueARM64_OpEq64F_0(v)
	case OpEq8:
		return psess.rewriteValueARM64_OpEq8_0(v)
	case OpEqB:
		return rewriteValueARM64_OpEqB_0(v)
	case OpEqPtr:
		return psess.rewriteValueARM64_OpEqPtr_0(v)
	case OpFloor:
		return rewriteValueARM64_OpFloor_0(v)
	case OpGeq16:
		return psess.rewriteValueARM64_OpGeq16_0(v)
	case OpGeq16U:
		return psess.rewriteValueARM64_OpGeq16U_0(v)
	case OpGeq32:
		return psess.rewriteValueARM64_OpGeq32_0(v)
	case OpGeq32F:
		return psess.rewriteValueARM64_OpGeq32F_0(v)
	case OpGeq32U:
		return psess.rewriteValueARM64_OpGeq32U_0(v)
	case OpGeq64:
		return psess.rewriteValueARM64_OpGeq64_0(v)
	case OpGeq64F:
		return psess.rewriteValueARM64_OpGeq64F_0(v)
	case OpGeq64U:
		return psess.rewriteValueARM64_OpGeq64U_0(v)
	case OpGeq8:
		return psess.rewriteValueARM64_OpGeq8_0(v)
	case OpGeq8U:
		return psess.rewriteValueARM64_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValueARM64_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValueARM64_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValueARM64_OpGetClosurePtr_0(v)
	case OpGreater16:
		return psess.rewriteValueARM64_OpGreater16_0(v)
	case OpGreater16U:
		return psess.rewriteValueARM64_OpGreater16U_0(v)
	case OpGreater32:
		return psess.rewriteValueARM64_OpGreater32_0(v)
	case OpGreater32F:
		return psess.rewriteValueARM64_OpGreater32F_0(v)
	case OpGreater32U:
		return psess.rewriteValueARM64_OpGreater32U_0(v)
	case OpGreater64:
		return psess.rewriteValueARM64_OpGreater64_0(v)
	case OpGreater64F:
		return psess.rewriteValueARM64_OpGreater64F_0(v)
	case OpGreater64U:
		return psess.rewriteValueARM64_OpGreater64U_0(v)
	case OpGreater8:
		return psess.rewriteValueARM64_OpGreater8_0(v)
	case OpGreater8U:
		return psess.rewriteValueARM64_OpGreater8U_0(v)
	case OpHmul32:
		return rewriteValueARM64_OpHmul32_0(v)
	case OpHmul32u:
		return rewriteValueARM64_OpHmul32u_0(v)
	case OpHmul64:
		return rewriteValueARM64_OpHmul64_0(v)
	case OpHmul64u:
		return rewriteValueARM64_OpHmul64u_0(v)
	case OpInterCall:
		return rewriteValueARM64_OpInterCall_0(v)
	case OpIsInBounds:
		return psess.rewriteValueARM64_OpIsInBounds_0(v)
	case OpIsNonNil:
		return psess.rewriteValueARM64_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return psess.rewriteValueARM64_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return psess.rewriteValueARM64_OpLeq16_0(v)
	case OpLeq16U:
		return psess.rewriteValueARM64_OpLeq16U_0(v)
	case OpLeq32:
		return psess.rewriteValueARM64_OpLeq32_0(v)
	case OpLeq32F:
		return psess.rewriteValueARM64_OpLeq32F_0(v)
	case OpLeq32U:
		return psess.rewriteValueARM64_OpLeq32U_0(v)
	case OpLeq64:
		return psess.rewriteValueARM64_OpLeq64_0(v)
	case OpLeq64F:
		return psess.rewriteValueARM64_OpLeq64F_0(v)
	case OpLeq64U:
		return psess.rewriteValueARM64_OpLeq64U_0(v)
	case OpLeq8:
		return psess.rewriteValueARM64_OpLeq8_0(v)
	case OpLeq8U:
		return psess.rewriteValueARM64_OpLeq8U_0(v)
	case OpLess16:
		return psess.rewriteValueARM64_OpLess16_0(v)
	case OpLess16U:
		return psess.rewriteValueARM64_OpLess16U_0(v)
	case OpLess32:
		return psess.rewriteValueARM64_OpLess32_0(v)
	case OpLess32F:
		return psess.rewriteValueARM64_OpLess32F_0(v)
	case OpLess32U:
		return psess.rewriteValueARM64_OpLess32U_0(v)
	case OpLess64:
		return psess.rewriteValueARM64_OpLess64_0(v)
	case OpLess64F:
		return psess.rewriteValueARM64_OpLess64F_0(v)
	case OpLess64U:
		return psess.rewriteValueARM64_OpLess64U_0(v)
	case OpLess8:
		return psess.rewriteValueARM64_OpLess8_0(v)
	case OpLess8U:
		return psess.rewriteValueARM64_OpLess8U_0(v)
	case OpLoad:
		return psess.rewriteValueARM64_OpLoad_0(v)
	case OpLsh16x16:
		return psess.rewriteValueARM64_OpLsh16x16_0(v)
	case OpLsh16x32:
		return psess.rewriteValueARM64_OpLsh16x32_0(v)
	case OpLsh16x64:
		return psess.rewriteValueARM64_OpLsh16x64_0(v)
	case OpLsh16x8:
		return psess.rewriteValueARM64_OpLsh16x8_0(v)
	case OpLsh32x16:
		return psess.rewriteValueARM64_OpLsh32x16_0(v)
	case OpLsh32x32:
		return psess.rewriteValueARM64_OpLsh32x32_0(v)
	case OpLsh32x64:
		return psess.rewriteValueARM64_OpLsh32x64_0(v)
	case OpLsh32x8:
		return psess.rewriteValueARM64_OpLsh32x8_0(v)
	case OpLsh64x16:
		return psess.rewriteValueARM64_OpLsh64x16_0(v)
	case OpLsh64x32:
		return psess.rewriteValueARM64_OpLsh64x32_0(v)
	case OpLsh64x64:
		return psess.rewriteValueARM64_OpLsh64x64_0(v)
	case OpLsh64x8:
		return psess.rewriteValueARM64_OpLsh64x8_0(v)
	case OpLsh8x16:
		return psess.rewriteValueARM64_OpLsh8x16_0(v)
	case OpLsh8x32:
		return psess.rewriteValueARM64_OpLsh8x32_0(v)
	case OpLsh8x64:
		return psess.rewriteValueARM64_OpLsh8x64_0(v)
	case OpLsh8x8:
		return psess.rewriteValueARM64_OpLsh8x8_0(v)
	case OpMod16:
		return rewriteValueARM64_OpMod16_0(v)
	case OpMod16u:
		return rewriteValueARM64_OpMod16u_0(v)
	case OpMod32:
		return rewriteValueARM64_OpMod32_0(v)
	case OpMod32u:
		return rewriteValueARM64_OpMod32u_0(v)
	case OpMod64:
		return rewriteValueARM64_OpMod64_0(v)
	case OpMod64u:
		return rewriteValueARM64_OpMod64u_0(v)
	case OpMod8:
		return rewriteValueARM64_OpMod8_0(v)
	case OpMod8u:
		return rewriteValueARM64_OpMod8u_0(v)
	case OpMove:
		return psess.rewriteValueARM64_OpMove_0(v) || psess.rewriteValueARM64_OpMove_10(v)
	case OpMul16:
		return rewriteValueARM64_OpMul16_0(v)
	case OpMul32:
		return rewriteValueARM64_OpMul32_0(v)
	case OpMul32F:
		return rewriteValueARM64_OpMul32F_0(v)
	case OpMul64:
		return rewriteValueARM64_OpMul64_0(v)
	case OpMul64F:
		return rewriteValueARM64_OpMul64F_0(v)
	case OpMul64uhilo:
		return rewriteValueARM64_OpMul64uhilo_0(v)
	case OpMul8:
		return rewriteValueARM64_OpMul8_0(v)
	case OpNeg16:
		return rewriteValueARM64_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValueARM64_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValueARM64_OpNeg32F_0(v)
	case OpNeg64:
		return rewriteValueARM64_OpNeg64_0(v)
	case OpNeg64F:
		return rewriteValueARM64_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValueARM64_OpNeg8_0(v)
	case OpNeq16:
		return psess.rewriteValueARM64_OpNeq16_0(v)
	case OpNeq32:
		return psess.rewriteValueARM64_OpNeq32_0(v)
	case OpNeq32F:
		return psess.rewriteValueARM64_OpNeq32F_0(v)
	case OpNeq64:
		return psess.rewriteValueARM64_OpNeq64_0(v)
	case OpNeq64F:
		return psess.rewriteValueARM64_OpNeq64F_0(v)
	case OpNeq8:
		return psess.rewriteValueARM64_OpNeq8_0(v)
	case OpNeqB:
		return rewriteValueARM64_OpNeqB_0(v)
	case OpNeqPtr:
		return psess.rewriteValueARM64_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValueARM64_OpNilCheck_0(v)
	case OpNot:
		return rewriteValueARM64_OpNot_0(v)
	case OpOffPtr:
		return rewriteValueARM64_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValueARM64_OpOr16_0(v)
	case OpOr32:
		return rewriteValueARM64_OpOr32_0(v)
	case OpOr64:
		return rewriteValueARM64_OpOr64_0(v)
	case OpOr8:
		return rewriteValueARM64_OpOr8_0(v)
	case OpOrB:
		return rewriteValueARM64_OpOrB_0(v)
	case OpPopCount16:
		return rewriteValueARM64_OpPopCount16_0(v)
	case OpPopCount32:
		return rewriteValueARM64_OpPopCount32_0(v)
	case OpPopCount64:
		return rewriteValueARM64_OpPopCount64_0(v)
	case OpRound:
		return rewriteValueARM64_OpRound_0(v)
	case OpRound32F:
		return rewriteValueARM64_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValueARM64_OpRound64F_0(v)
	case OpRsh16Ux16:
		return psess.rewriteValueARM64_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return psess.rewriteValueARM64_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return psess.rewriteValueARM64_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return psess.rewriteValueARM64_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return psess.rewriteValueARM64_OpRsh16x16_0(v)
	case OpRsh16x32:
		return psess.rewriteValueARM64_OpRsh16x32_0(v)
	case OpRsh16x64:
		return psess.rewriteValueARM64_OpRsh16x64_0(v)
	case OpRsh16x8:
		return psess.rewriteValueARM64_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return psess.rewriteValueARM64_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return psess.rewriteValueARM64_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return psess.rewriteValueARM64_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return psess.rewriteValueARM64_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return psess.rewriteValueARM64_OpRsh32x16_0(v)
	case OpRsh32x32:
		return psess.rewriteValueARM64_OpRsh32x32_0(v)
	case OpRsh32x64:
		return psess.rewriteValueARM64_OpRsh32x64_0(v)
	case OpRsh32x8:
		return psess.rewriteValueARM64_OpRsh32x8_0(v)
	case OpRsh64Ux16:
		return psess.rewriteValueARM64_OpRsh64Ux16_0(v)
	case OpRsh64Ux32:
		return psess.rewriteValueARM64_OpRsh64Ux32_0(v)
	case OpRsh64Ux64:
		return psess.rewriteValueARM64_OpRsh64Ux64_0(v)
	case OpRsh64Ux8:
		return psess.rewriteValueARM64_OpRsh64Ux8_0(v)
	case OpRsh64x16:
		return psess.rewriteValueARM64_OpRsh64x16_0(v)
	case OpRsh64x32:
		return psess.rewriteValueARM64_OpRsh64x32_0(v)
	case OpRsh64x64:
		return psess.rewriteValueARM64_OpRsh64x64_0(v)
	case OpRsh64x8:
		return psess.rewriteValueARM64_OpRsh64x8_0(v)
	case OpRsh8Ux16:
		return psess.rewriteValueARM64_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return psess.rewriteValueARM64_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return psess.rewriteValueARM64_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return psess.rewriteValueARM64_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return psess.rewriteValueARM64_OpRsh8x16_0(v)
	case OpRsh8x32:
		return psess.rewriteValueARM64_OpRsh8x32_0(v)
	case OpRsh8x64:
		return psess.rewriteValueARM64_OpRsh8x64_0(v)
	case OpRsh8x8:
		return psess.rewriteValueARM64_OpRsh8x8_0(v)
	case OpSignExt16to32:
		return rewriteValueARM64_OpSignExt16to32_0(v)
	case OpSignExt16to64:
		return rewriteValueARM64_OpSignExt16to64_0(v)
	case OpSignExt32to64:
		return rewriteValueARM64_OpSignExt32to64_0(v)
	case OpSignExt8to16:
		return rewriteValueARM64_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValueARM64_OpSignExt8to32_0(v)
	case OpSignExt8to64:
		return rewriteValueARM64_OpSignExt8to64_0(v)
	case OpSlicemask:
		return rewriteValueARM64_OpSlicemask_0(v)
	case OpSqrt:
		return rewriteValueARM64_OpSqrt_0(v)
	case OpStaticCall:
		return rewriteValueARM64_OpStaticCall_0(v)
	case OpStore:
		return psess.rewriteValueARM64_OpStore_0(v)
	case OpSub16:
		return rewriteValueARM64_OpSub16_0(v)
	case OpSub32:
		return rewriteValueARM64_OpSub32_0(v)
	case OpSub32F:
		return rewriteValueARM64_OpSub32F_0(v)
	case OpSub64:
		return rewriteValueARM64_OpSub64_0(v)
	case OpSub64F:
		return rewriteValueARM64_OpSub64F_0(v)
	case OpSub8:
		return rewriteValueARM64_OpSub8_0(v)
	case OpSubPtr:
		return rewriteValueARM64_OpSubPtr_0(v)
	case OpTrunc:
		return rewriteValueARM64_OpTrunc_0(v)
	case OpTrunc16to8:
		return rewriteValueARM64_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValueARM64_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValueARM64_OpTrunc32to8_0(v)
	case OpTrunc64to16:
		return rewriteValueARM64_OpTrunc64to16_0(v)
	case OpTrunc64to32:
		return rewriteValueARM64_OpTrunc64to32_0(v)
	case OpTrunc64to8:
		return rewriteValueARM64_OpTrunc64to8_0(v)
	case OpWB:
		return rewriteValueARM64_OpWB_0(v)
	case OpXor16:
		return rewriteValueARM64_OpXor16_0(v)
	case OpXor32:
		return rewriteValueARM64_OpXor32_0(v)
	case OpXor64:
		return rewriteValueARM64_OpXor64_0(v)
	case OpXor8:
		return rewriteValueARM64_OpXor8_0(v)
	case OpZero:
		return psess.rewriteValueARM64_OpZero_0(v) || psess.rewriteValueARM64_OpZero_10(v) || psess.rewriteValueARM64_OpZero_20(v)
	case OpZeroExt16to32:
		return rewriteValueARM64_OpZeroExt16to32_0(v)
	case OpZeroExt16to64:
		return rewriteValueARM64_OpZeroExt16to64_0(v)
	case OpZeroExt32to64:
		return rewriteValueARM64_OpZeroExt32to64_0(v)
	case OpZeroExt8to16:
		return rewriteValueARM64_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValueARM64_OpZeroExt8to32_0(v)
	case OpZeroExt8to64:
		return rewriteValueARM64_OpZeroExt8to64_0(v)
	}
	return false
}
func rewriteValueARM64_OpARM64ADD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64NEG {
			break
		}
		y := v_1.Args[0]
		v.reset(OpARM64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64NEG {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARM64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ADDshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ADDshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ADDshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ADDshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ADDshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ADDshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ADDconst_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpARM64MOVDaddr)
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
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = c + d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARM64ADDconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARM64ADDconst)
		v.AuxInt = c - d
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64ADDshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ADDconst)
		v.AuxInt = int64(uint64(c) << uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		if v_0.AuxInt != 64-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARM64RORconst)
		v.AuxInt = 64 - c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFX {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		if !(c < 32 && t.Size(psess.types) == 4 && bfc == arm64BFAuxInt(32-c, c)) {
			break
		}
		v.reset(OpARM64RORWconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		if v_0.AuxInt != 64-c {
			break
		}
		x := v_0.Args[0]
		x2 := v.Args[1]
		v.reset(OpARM64EXTRconst)
		v.AuxInt = 64 - c
		v.AddArg(x2)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFX {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		x2 := v.Args[1]
		if !(c < 32 && t.Size(psess.types) == 4 && bfc == arm64BFAuxInt(32-c, c)) {
			break
		}
		v.reset(OpARM64EXTRWconst)
		v.AuxInt = 32 - c
		v.AddArg(x2)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ADDshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ADDconst)
		v.AuxInt = c >> uint64(d)
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64ADDshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ADDconst)
		v.AuxInt = int64(uint64(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 64-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARM64RORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVWUreg {
			break
		}
		if x != v_1.Args[0] {
			break
		}
		if !(c < 32 && t.Size(psess.types) == 4) {
			break
		}
		v.reset(OpARM64RORWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64AND_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ANDconst)
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
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MVN {
			break
		}
		y := v_1.Args[0]
		v.reset(OpARM64BIC)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MVN {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARM64BIC)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ANDshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ANDshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ANDshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ANDshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ANDshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64AND_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ANDshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ANDconst_0(v *Value) bool {

	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpARM64MOVDconst)
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
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = c & d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARM64ANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVWUreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64ANDconst)
		v.AuxInt = c & (1<<32 - 1)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVHUreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64ANDconst)
		v.AuxInt = c & (1<<16 - 1)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVBUreg {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64ANDconst)
		v.AuxInt = c & (1<<8 - 1)
		v.AddArg(x)
		return true
	}

	for {
		ac := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, ac, sc)) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(sc, arm64BFWidth(ac, sc))
		v.AddArg(x)
		return true
	}

	for {
		ac := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, ac, 0)) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(sc, arm64BFWidth(ac, 0))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ANDshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ANDconst)
		v.AuxInt = int64(uint64(c) << uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARM64SLLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ANDshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ANDconst)
		v.AuxInt = c >> uint64(d)
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARM64SRAconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ANDshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ANDconst)
		v.AuxInt = int64(uint64(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARM64SRLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64BIC_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ANDconst)
		v.AuxInt = ^c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64BICshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64BICshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64BICshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64BICshiftLL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ANDconst)
		v.AuxInt = ^int64(uint64(c) << uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64BICshiftRA_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ANDconst)
		v.AuxInt = ^(c >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64BICshiftRL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ANDconst)
		v.AuxInt = ^int64(uint64(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64CMN_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64CMNconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64CMNWconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(-y)) {
			break
		}
		v.reset(OpARM64FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(-y) && uint32(x) < uint32(-y)) {
			break
		}
		v.reset(OpARM64FlagLT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(-y) && uint32(x) > uint32(-y)) {
			break
		}
		v.reset(OpARM64FlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(-y) && uint32(x) < uint32(-y)) {
			break
		}
		v.reset(OpARM64FlagGT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(-y) && uint32(x) > uint32(-y)) {
			break
		}
		v.reset(OpARM64FlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64CMNconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x) == int64(-y)) {
			break
		}
		v.reset(OpARM64FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x) < int64(-y) && uint64(x) < uint64(-y)) {
			break
		}
		v.reset(OpARM64FlagLT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x) < int64(-y) && uint64(x) > uint64(-y)) {
			break
		}
		v.reset(OpARM64FlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x) > int64(-y) && uint64(x) < uint64(-y)) {
			break
		}
		v.reset(OpARM64FlagGT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x) > int64(-y) && uint64(x) > uint64(-y)) {
			break
		}
		v.reset(OpARM64FlagGT_UGT)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64CMP_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64CMPconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64CMPshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpARM64SLLconst {
			break
		}
		c := x0.AuxInt
		y := x0.Args[0]
		x1 := v.Args[1]
		if !(clobberIfDead(x0)) {
			break
		}
		v.reset(OpARM64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpARM64CMPshiftLL, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x1)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64CMPshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpARM64SRLconst {
			break
		}
		c := x0.AuxInt
		y := x0.Args[0]
		x1 := v.Args[1]
		if !(clobberIfDead(x0)) {
			break
		}
		v.reset(OpARM64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpARM64CMPshiftRL, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x1)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64CMPshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != OpARM64SRAconst {
			break
		}
		c := x0.AuxInt
		y := x0.Args[0]
		x1 := v.Args[1]
		if !(clobberIfDead(x0)) {
			break
		}
		v.reset(OpARM64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpARM64CMPshiftRA, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x1)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64CMPW_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64CMPWconst)
		v.AuxInt = int64(int32(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpARM64CMPWconst, psess.types.TypeFlags)
		v0.AuxInt = int64(int32(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64CMPWconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpARM64FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpARM64FlagLT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpARM64FlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpARM64FlagGT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpARM64FlagGT_UGT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVBUreg {
			break
		}
		if !(0xff < int32(c)) {
			break
		}
		v.reset(OpARM64FlagLT_ULT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVHUreg {
			break
		}
		if !(0xffff < int32(c)) {
			break
		}
		v.reset(OpARM64FlagLT_ULT)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64CMPconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x == y) {
			break
		}
		v.reset(OpARM64FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x < y && uint64(x) < uint64(y)) {
			break
		}
		v.reset(OpARM64FlagLT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x < y && uint64(x) > uint64(y)) {
			break
		}
		v.reset(OpARM64FlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x > y && uint64(x) < uint64(y)) {
			break
		}
		v.reset(OpARM64FlagGT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(x > y && uint64(x) > uint64(y)) {
			break
		}
		v.reset(OpARM64FlagGT_UGT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVBUreg {
			break
		}
		if !(0xff < c) {
			break
		}
		v.reset(OpARM64FlagLT_ULT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVHUreg {
			break
		}
		if !(0xffff < c) {
			break
		}
		v.reset(OpARM64FlagLT_ULT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVWUreg {
			break
		}
		if !(0xffffffff < c) {
			break
		}
		v.reset(OpARM64FlagLT_ULT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ANDconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= m && m < n) {
			break
		}
		v.reset(OpARM64FlagLT_ULT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		c := v_0.AuxInt
		if !(0 <= n && 0 < c && c <= 63 && (1<<uint64(64-c)) <= uint64(n)) {
			break
		}
		v.reset(OpARM64FlagLT_ULT)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64CMPshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v1.AuxInt = d
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64CMPconst)
		v.AuxInt = int64(uint64(c) << uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64CMPshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARM64SRAconst, x.Type)
		v1.AuxInt = d
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64CMPconst)
		v.AuxInt = c >> uint64(d)
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64CMPshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64InvertFlags)
		v0 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARM64SRLconst, x.Type)
		v1.AuxInt = d
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64CMPconst)
		v.AuxInt = int64(uint64(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64CSEL_0(v *Value) bool {

	for {
		cc := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		flag := v.Args[2]
		v.reset(OpARM64CSEL0)
		v.Aux = cc
		v.AddArg(x)
		v.AddArg(flag)
		return true
	}

	for {
		cc := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		y := v.Args[1]
		flag := v.Args[2]
		v.reset(OpARM64CSEL0)
		v.Aux = arm64Negate(cc.(Op))
		v.AddArg(y)
		v.AddArg(flag)
		return true
	}

	for {
		cc := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64InvertFlags {
			break
		}
		cmp := v_2.Args[0]
		v.reset(OpARM64CSEL)
		v.Aux = arm64Invert(cc.(Op))
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(cmp)
		return true
	}

	for {
		cc := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		flag := v.Args[2]
		if !(ccARM64Eval(cc, flag) > 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		cc := v.Aux
		_ = v.Args[2]
		y := v.Args[1]
		flag := v.Args[2]
		if !(ccARM64Eval(cc, flag) < 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		cc := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64CMPWconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		bool := v_2.Args[0]
		if !(cc.(Op) == OpARM64NotEqual && psess.flagArg(bool) != nil) {
			break
		}
		v.reset(OpARM64CSEL)
		v.Aux = bool.Op
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(psess.flagArg(bool))
		return true
	}

	for {
		cc := v.Aux
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64CMPWconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		bool := v_2.Args[0]
		if !(cc.(Op) == OpARM64Equal && psess.flagArg(bool) != nil) {
			break
		}
		v.reset(OpARM64CSEL)
		v.Aux = arm64Negate(bool.Op)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(psess.flagArg(bool))
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64CSEL0_0(v *Value) bool {

	for {
		cc := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64InvertFlags {
			break
		}
		cmp := v_1.Args[0]
		v.reset(OpARM64CSEL0)
		v.Aux = arm64Invert(cc.(Op))
		v.AddArg(x)
		v.AddArg(cmp)
		return true
	}

	for {
		cc := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		flag := v.Args[1]
		if !(ccARM64Eval(cc, flag) > 0) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		cc := v.Aux
		_ = v.Args[1]
		flag := v.Args[1]
		if !(ccARM64Eval(cc, flag) < 0) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		cc := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64CMPWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		bool := v_1.Args[0]
		if !(cc.(Op) == OpARM64NotEqual && psess.flagArg(bool) != nil) {
			break
		}
		v.reset(OpARM64CSEL0)
		v.Aux = bool.Op
		v.AddArg(x)
		v.AddArg(psess.flagArg(bool))
		return true
	}

	for {
		cc := v.Aux
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64CMPWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		bool := v_1.Args[0]
		if !(cc.(Op) == OpARM64Equal && psess.flagArg(bool) != nil) {
			break
		}
		v.reset(OpARM64CSEL0)
		v.Aux = arm64Negate(bool.Op)
		v.AddArg(x)
		v.AddArg(psess.flagArg(bool))
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64DIV_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = c / d
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64DIVW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(int32(c) / int32(d))
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64EON_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64XORconst)
		v.AuxInt = ^c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -1
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64EONshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64EONshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64EONshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64EONshiftLL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64XORconst)
		v.AuxInt = ^int64(uint64(c) << uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -1
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64EONshiftRA_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64XORconst)
		v.AuxInt = ^(c >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -1
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64EONshiftRL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64XORconst)
		v.AuxInt = ^int64(uint64(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -1
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64Equal_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagEQ {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64Equal)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FADDD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FMULD {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpARM64FMADDD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FMULD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		v.reset(OpARM64FMADDD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FNMULD {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpARM64FMSUBD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FNMULD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		v.reset(OpARM64FMSUBD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FADDS_0(v *Value) bool {

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FMULS {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpARM64FMADDS)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FMULS {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		v.reset(OpARM64FMADDS)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FNMULS {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpARM64FMSUBS)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FNMULS {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		v.reset(OpARM64FMSUBS)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FMOVDgpfp_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		v_0 := v.Args[0]
		if v_0.Op != OpArg {
			break
		}
		off := v_0.AuxInt
		sym := v_0.Aux
		b = b.Func.Entry
		v0 := b.NewValue0(v.Pos, OpArg, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AuxInt = off
		v0.Aux = sym
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FMOVDload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64FMOVDload)
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
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64FMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FMOVDstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FMOVDgpfp {
			break
		}
		val := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64FMOVDstore)
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
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64FMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FMOVSload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64FMOVSload)
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
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64FMOVSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FMOVSstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64FMOVSstore)
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
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64FMOVSstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FMULD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FNEGD {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FNMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FNEGD {
			break
		}
		x := v_1.Args[0]
		v.reset(OpARM64FNMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FMULS_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FNEGS {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FNMULS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FNEGS {
			break
		}
		x := v_1.Args[0]
		v.reset(OpARM64FNMULS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FNEGD_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FMULD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpARM64FNMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FNMULD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpARM64FMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FNEGS_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FMULS {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpARM64FNMULS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FNMULS {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpARM64FMULS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FNMULD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FNEGD {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FNEGD {
			break
		}
		x := v_1.Args[0]
		v.reset(OpARM64FMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FNMULS_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FNEGS {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FMULS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FNEGS {
			break
		}
		x := v_1.Args[0]
		v.reset(OpARM64FMULS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FSUBD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FMULD {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpARM64FMSUBD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FMULD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		v.reset(OpARM64FNMSUBD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FNMULD {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpARM64FMADDD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FNMULD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		v.reset(OpARM64FNMADDD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64FSUBS_0(v *Value) bool {

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FMULS {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpARM64FMSUBS)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FMULS {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		v.reset(OpARM64FNMSUBS)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FNMULS {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpARM64FMADDS)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FNMULS {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		v.reset(OpARM64FNMADDS)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64GreaterEqual_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagEQ {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64LessEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64GreaterEqualU_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagEQ {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64LessEqualU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64GreaterThan_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagEQ {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64LessThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64GreaterThanU_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagEQ {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64LessThanU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64LessEqual_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagEQ {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64GreaterEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64LessEqualU_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagEQ {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64GreaterEqualU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64LessThan_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagEQ {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64GreaterThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64LessThanU_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagEQ {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64GreaterThanU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MNEG_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
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
		if v_0.Op != OpARM64MOVDconst {
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
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c-1) && c >= 3) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c-1) && c >= 3) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MNEG_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c+1) && c >= 7) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v1 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c+1) && c >= 7) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v1 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%3 == 0 && isPowerOfTwo(c/3)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.Type = x.Type
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpARM64SUBshiftLL, x.Type)
		v0.AuxInt = 2
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%3 == 0 && isPowerOfTwo(c/3)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.Type = x.Type
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpARM64SUBshiftLL, x.Type)
		v0.AuxInt = 2
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%5 == 0 && isPowerOfTwo(c/5)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%5 == 0 && isPowerOfTwo(c/5)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%7 == 0 && isPowerOfTwo(c/7)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.Type = x.Type
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpARM64SUBshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%7 == 0 && isPowerOfTwo(c/7)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.Type = x.Type
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpARM64SUBshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%9 == 0 && isPowerOfTwo(c/9)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%9 == 0 && isPowerOfTwo(c/9)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MNEG_20(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -c * d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -c * d
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MNEGW_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(int32(c) == -1) {
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
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(int32(c) == 1) {
			break
		}
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(int32(c) == 1) {
			break
		}
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MNEGW_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v1 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v1 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.Type = x.Type
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpARM64SUBshiftLL, x.Type)
		v0.AuxInt = 2
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.Type = x.Type
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpARM64SUBshiftLL, x.Type)
		v0.AuxInt = 2
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.Type = x.Type
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpARM64SUBshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.Type = x.Type
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpARM64SUBshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64NEG)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MNEGW_20(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -int64(int32(c) * int32(d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -int64(int32(c) * int32(d))
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = c % d
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MODW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(int32(c) % int32(d))
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBUload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVBUloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVBUload)
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
		if v_1.Op != OpARM64MOVBstorezero {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBUloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVBUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVBUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVBstorezeroidx {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2) || isSamePtr(ptr, idx2) && isSamePtr(idx, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBUreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARM64ANDconst)
		v.AuxInt = c & (1<<8 - 1)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(uint8(c))
		return true
	}

	for {
		x := v.Args[0]
		if !(x.Type.IsBoolean()) {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<8-1, sc)) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(sc, arm64BFWidth(1<<8-1, sc))
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<8-1, 0)) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(sc, 8)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVBloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVBload)
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
		if v_1.Op != OpARM64MOVBstorezero {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVBload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVBload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVBstorezeroidx {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2) || isSamePtr(ptr, idx2) && isSamePtr(idx, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(int8(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		lc := v_0.AuxInt
		x := v_0.Args[0]
		if !(lc < 8) {
			break
		}
		v.reset(OpARM64SBFIZ)
		v.AuxInt = arm64BFAuxInt(lc, 8-lc)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVBstore)
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
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVBstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVBstore)
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
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpARM64MOVBstorezero)
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
		if v_1.Op != OpARM64MOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVBstore)
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
		if v_1.Op != OpARM64MOVBUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVBstore)
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
		if v_1.Op != OpARM64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVBstore)
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
		if v_1.Op != OpARM64MOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVBstore)
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
		if v_1.Op != OpARM64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVBstore)
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
		if v_1.Op != OpARM64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBstore_10(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64UBFX {
			break
		}
		if v_1.AuxInt != arm64BFAuxInt(8, 8) {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64UBFX {
			break
		}
		if v_1.AuxInt != arm64BFAuxInt(8, 8) {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64UBFX {
			break
		}
		if v_1.AuxInt != arm64BFAuxInt(8, 24) {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64UBFX {
			break
		}
		if v_1.AuxInt != arm64BFAuxInt(8, 24) {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 8 {
			break
		}
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpARM64MOVDreg {
			break
		}
		w := v_1_0.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 8 {
			break
		}
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpARM64MOVDreg {
			break
		}
		w := v_1_0.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		w0 := x.Args[1]
		if w0.Op != OpARM64SRLconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		w0 := x.Args[2]
		if w0.Op != OpARM64SRLconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBstore_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64UBFX {
			break
		}
		bfc := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		w0 := x.Args[1]
		if w0.Op != OpARM64UBFX {
			break
		}
		bfc2 := w0.AuxInt
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && getARM64BFwidth(bfc) == 32-getARM64BFlsb(bfc) && getARM64BFwidth(bfc2) == 32-getARM64BFlsb(bfc2) && getARM64BFlsb(bfc2) == getARM64BFlsb(bfc)-8 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64UBFX {
			break
		}
		bfc := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		w0 := x.Args[2]
		if w0.Op != OpARM64UBFX {
			break
		}
		bfc2 := w0.AuxInt
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && getARM64BFwidth(bfc) == 32-getARM64BFlsb(bfc) && getARM64BFwidth(bfc2) == 32-getARM64BFlsb(bfc2) && getARM64BFlsb(bfc2) == getARM64BFlsb(bfc)-8 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		j := v_1.AuxInt
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpARM64MOVDreg {
			break
		}
		w := v_1_0.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		w0 := x.Args[1]
		if w0.Op != OpARM64SRLconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		w0_0 := w0.Args[0]
		if w0_0.Op != OpARM64MOVDreg {
			break
		}
		if w != w0_0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		j := v_1.AuxInt
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpARM64MOVDreg {
			break
		}
		w := v_1_0.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		w0 := x.Args[2]
		if w0.Op != OpARM64SRLconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		w0_0 := w0.Args[0]
		if w0_0.Op != OpARM64MOVDreg {
			break
		}
		if w != w0_0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpARM64MOVBstore {
			break
		}
		if x0.AuxInt != i-1 {
			break
		}
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if ptr != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64SRLconst {
			break
		}
		if x0_1.AuxInt != 8 {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpARM64MOVBstore {
			break
		}
		if x1.AuxInt != i-2 {
			break
		}
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64SRLconst {
			break
		}
		if x1_1.AuxInt != 16 {
			break
		}
		if w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpARM64MOVBstore {
			break
		}
		if x2.AuxInt != i-3 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64SRLconst {
			break
		}
		if x2_1.AuxInt != 24 {
			break
		}
		if w != x2_1.Args[0] {
			break
		}
		x3 := x2.Args[2]
		if x3.Op != OpARM64MOVBstore {
			break
		}
		if x3.AuxInt != i-4 {
			break
		}
		if x3.Aux != s {
			break
		}
		_ = x3.Args[2]
		if ptr != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpARM64SRLconst {
			break
		}
		if x3_1.AuxInt != 32 {
			break
		}
		if w != x3_1.Args[0] {
			break
		}
		x4 := x3.Args[2]
		if x4.Op != OpARM64MOVBstore {
			break
		}
		if x4.AuxInt != i-5 {
			break
		}
		if x4.Aux != s {
			break
		}
		_ = x4.Args[2]
		if ptr != x4.Args[0] {
			break
		}
		x4_1 := x4.Args[1]
		if x4_1.Op != OpARM64SRLconst {
			break
		}
		if x4_1.AuxInt != 40 {
			break
		}
		if w != x4_1.Args[0] {
			break
		}
		x5 := x4.Args[2]
		if x5.Op != OpARM64MOVBstore {
			break
		}
		if x5.AuxInt != i-6 {
			break
		}
		if x5.Aux != s {
			break
		}
		_ = x5.Args[2]
		if ptr != x5.Args[0] {
			break
		}
		x5_1 := x5.Args[1]
		if x5_1.Op != OpARM64SRLconst {
			break
		}
		if x5_1.AuxInt != 48 {
			break
		}
		if w != x5_1.Args[0] {
			break
		}
		x6 := x5.Args[2]
		if x6.Op != OpARM64MOVBstore {
			break
		}
		if x6.AuxInt != i-7 {
			break
		}
		if x6.Aux != s {
			break
		}
		_ = x6.Args[2]
		if ptr != x6.Args[0] {
			break
		}
		x6_1 := x6.Args[1]
		if x6_1.Op != OpARM64SRLconst {
			break
		}
		if x6_1.AuxInt != 56 {
			break
		}
		if w != x6_1.Args[0] {
			break
		}
		mem := x6.Args[2]
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6)) {
			break
		}
		v.reset(OpARM64MOVDstore)
		v.AuxInt = i - 7
		v.Aux = s
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64REV, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 7 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpARM64MOVBstore {
			break
		}
		if x0.AuxInt != 6 {
			break
		}
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64SRLconst {
			break
		}
		if x0_1.AuxInt != 8 {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpARM64MOVBstore {
			break
		}
		if x1.AuxInt != 5 {
			break
		}
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if p != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64SRLconst {
			break
		}
		if x1_1.AuxInt != 16 {
			break
		}
		if w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpARM64MOVBstore {
			break
		}
		if x2.AuxInt != 4 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64SRLconst {
			break
		}
		if x2_1.AuxInt != 24 {
			break
		}
		if w != x2_1.Args[0] {
			break
		}
		x3 := x2.Args[2]
		if x3.Op != OpARM64MOVBstore {
			break
		}
		if x3.AuxInt != 3 {
			break
		}
		if x3.Aux != s {
			break
		}
		_ = x3.Args[2]
		if p != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpARM64SRLconst {
			break
		}
		if x3_1.AuxInt != 32 {
			break
		}
		if w != x3_1.Args[0] {
			break
		}
		x4 := x3.Args[2]
		if x4.Op != OpARM64MOVBstore {
			break
		}
		if x4.AuxInt != 2 {
			break
		}
		if x4.Aux != s {
			break
		}
		_ = x4.Args[2]
		if p != x4.Args[0] {
			break
		}
		x4_1 := x4.Args[1]
		if x4_1.Op != OpARM64SRLconst {
			break
		}
		if x4_1.AuxInt != 40 {
			break
		}
		if w != x4_1.Args[0] {
			break
		}
		x5 := x4.Args[2]
		if x5.Op != OpARM64MOVBstore {
			break
		}
		if x5.AuxInt != 1 {
			break
		}
		if x5.Aux != s {
			break
		}
		_ = x5.Args[2]
		p1 := x5.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		x5_1 := x5.Args[1]
		if x5_1.Op != OpARM64SRLconst {
			break
		}
		if x5_1.AuxInt != 48 {
			break
		}
		if w != x5_1.Args[0] {
			break
		}
		x6 := x5.Args[2]
		if x6.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x6.Args[3]
		ptr0 := x6.Args[0]
		idx0 := x6.Args[1]
		x6_2 := x6.Args[2]
		if x6_2.Op != OpARM64SRLconst {
			break
		}
		if x6_2.AuxInt != 56 {
			break
		}
		if w != x6_2.Args[0] {
			break
		}
		mem := x6.Args[3]
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6)) {
			break
		}
		v.reset(OpARM64MOVDstoreidx)
		v.AddArg(ptr0)
		v.AddArg(idx0)
		v0 := b.NewValue0(v.Pos, OpARM64REV, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpARM64MOVBstore {
			break
		}
		if x0.AuxInt != i-1 {
			break
		}
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if ptr != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64UBFX {
			break
		}
		if x0_1.AuxInt != arm64BFAuxInt(8, 24) {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpARM64MOVBstore {
			break
		}
		if x1.AuxInt != i-2 {
			break
		}
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64UBFX {
			break
		}
		if x1_1.AuxInt != arm64BFAuxInt(16, 16) {
			break
		}
		if w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpARM64MOVBstore {
			break
		}
		if x2.AuxInt != i-3 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64UBFX {
			break
		}
		if x2_1.AuxInt != arm64BFAuxInt(24, 8) {
			break
		}
		if w != x2_1.Args[0] {
			break
		}
		mem := x2.Args[2]
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpARM64MOVWstore)
		v.AuxInt = i - 3
		v.Aux = s
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 3 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpARM64MOVBstore {
			break
		}
		if x0.AuxInt != 2 {
			break
		}
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64UBFX {
			break
		}
		if x0_1.AuxInt != arm64BFAuxInt(8, 24) {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpARM64MOVBstore {
			break
		}
		if x1.AuxInt != 1 {
			break
		}
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64UBFX {
			break
		}
		if x1_1.AuxInt != arm64BFAuxInt(16, 16) {
			break
		}
		if w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x2.Args[3]
		ptr0 := x2.Args[0]
		idx0 := x2.Args[1]
		x2_2 := x2.Args[2]
		if x2_2.Op != OpARM64UBFX {
			break
		}
		if x2_2.AuxInt != arm64BFAuxInt(24, 8) {
			break
		}
		if w != x2_2.Args[0] {
			break
		}
		mem := x2.Args[3]
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr0)
		v.AddArg(idx0)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpARM64MOVBstore {
			break
		}
		if x0.AuxInt != i-1 {
			break
		}
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if ptr != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64SRLconst {
			break
		}
		if x0_1.AuxInt != 8 {
			break
		}
		x0_1_0 := x0_1.Args[0]
		if x0_1_0.Op != OpARM64MOVDreg {
			break
		}
		if w != x0_1_0.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpARM64MOVBstore {
			break
		}
		if x1.AuxInt != i-2 {
			break
		}
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64SRLconst {
			break
		}
		if x1_1.AuxInt != 16 {
			break
		}
		x1_1_0 := x1_1.Args[0]
		if x1_1_0.Op != OpARM64MOVDreg {
			break
		}
		if w != x1_1_0.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpARM64MOVBstore {
			break
		}
		if x2.AuxInt != i-3 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64SRLconst {
			break
		}
		if x2_1.AuxInt != 24 {
			break
		}
		x2_1_0 := x2_1.Args[0]
		if x2_1_0.Op != OpARM64MOVDreg {
			break
		}
		if w != x2_1_0.Args[0] {
			break
		}
		mem := x2.Args[2]
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpARM64MOVWstore)
		v.AuxInt = i - 3
		v.Aux = s
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 3 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpARM64MOVBstore {
			break
		}
		if x0.AuxInt != 2 {
			break
		}
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64SRLconst {
			break
		}
		if x0_1.AuxInt != 8 {
			break
		}
		x0_1_0 := x0_1.Args[0]
		if x0_1_0.Op != OpARM64MOVDreg {
			break
		}
		if w != x0_1_0.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpARM64MOVBstore {
			break
		}
		if x1.AuxInt != 1 {
			break
		}
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64SRLconst {
			break
		}
		if x1_1.AuxInt != 16 {
			break
		}
		x1_1_0 := x1_1.Args[0]
		if x1_1_0.Op != OpARM64MOVDreg {
			break
		}
		if w != x1_1_0.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x2.Args[3]
		ptr0 := x2.Args[0]
		idx0 := x2.Args[1]
		x2_2 := x2.Args[2]
		if x2_2.Op != OpARM64SRLconst {
			break
		}
		if x2_2.AuxInt != 24 {
			break
		}
		x2_2_0 := x2_2.Args[0]
		if x2_2_0.Op != OpARM64MOVDreg {
			break
		}
		if w != x2_2_0.Args[0] {
			break
		}
		mem := x2.Args[3]
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr0)
		v.AddArg(idx0)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBstore_30(v *Value) bool {
	b := v.Block
	_ = b

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpARM64MOVBstore {
			break
		}
		if x0.AuxInt != i-1 {
			break
		}
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if ptr != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64SRLconst {
			break
		}
		if x0_1.AuxInt != 8 {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpARM64MOVBstore {
			break
		}
		if x1.AuxInt != i-2 {
			break
		}
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64SRLconst {
			break
		}
		if x1_1.AuxInt != 16 {
			break
		}
		if w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpARM64MOVBstore {
			break
		}
		if x2.AuxInt != i-3 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64SRLconst {
			break
		}
		if x2_1.AuxInt != 24 {
			break
		}
		if w != x2_1.Args[0] {
			break
		}
		mem := x2.Args[2]
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpARM64MOVWstore)
		v.AuxInt = i - 3
		v.Aux = s
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 3 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		w := v.Args[1]
		x0 := v.Args[2]
		if x0.Op != OpARM64MOVBstore {
			break
		}
		if x0.AuxInt != 2 {
			break
		}
		if x0.Aux != s {
			break
		}
		_ = x0.Args[2]
		if p != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64SRLconst {
			break
		}
		if x0_1.AuxInt != 8 {
			break
		}
		if w != x0_1.Args[0] {
			break
		}
		x1 := x0.Args[2]
		if x1.Op != OpARM64MOVBstore {
			break
		}
		if x1.AuxInt != 1 {
			break
		}
		if x1.Aux != s {
			break
		}
		_ = x1.Args[2]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64SRLconst {
			break
		}
		if x1_1.AuxInt != 16 {
			break
		}
		if w != x1_1.Args[0] {
			break
		}
		x2 := x1.Args[2]
		if x2.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x2.Args[3]
		ptr0 := x2.Args[0]
		idx0 := x2.Args[1]
		x2_2 := x2.Args[2]
		if x2_2.Op != OpARM64SRLconst {
			break
		}
		if x2_2.AuxInt != 24 {
			break
		}
		if w != x2_2.Args[0] {
			break
		}
		mem := x2.Args[3]
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr0)
		v.AddArg(idx0)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if ptr != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpARM64SRLconst {
			break
		}
		if x_1.AuxInt != 8 {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr1 := v_0.Args[0]
		idx1 := v_0.Args[1]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr0 := x.Args[0]
		idx0 := x.Args[1]
		x_2 := x.Args[2]
		if x_2.Op != OpARM64SRLconst {
			break
		}
		if x_2.AuxInt != 8 {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr0)
		v.AddArg(idx0)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if ptr != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpARM64UBFX {
			break
		}
		if x_1.AuxInt != arm64BFAuxInt(8, 8) {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr1 := v_0.Args[0]
		idx1 := v_0.Args[1]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr0 := x.Args[0]
		idx0 := x.Args[1]
		x_2 := x.Args[2]
		if x_2.Op != OpARM64UBFX {
			break
		}
		if x_2.AuxInt != arm64BFAuxInt(8, 8) {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr0)
		v.AddArg(idx0)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if ptr != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpARM64SRLconst {
			break
		}
		if x_1.AuxInt != 8 {
			break
		}
		x_1_0 := x_1.Args[0]
		if x_1_0.Op != OpARM64MOVDreg {
			break
		}
		if w != x_1_0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr1 := v_0.Args[0]
		idx1 := v_0.Args[1]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr0 := x.Args[0]
		idx0 := x.Args[1]
		x_2 := x.Args[2]
		if x_2.Op != OpARM64SRLconst {
			break
		}
		if x_2.AuxInt != 8 {
			break
		}
		x_2_0 := x_2.Args[0]
		if x_2_0.Op != OpARM64MOVDreg {
			break
		}
		if w != x_2_0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr0)
		v.AddArg(idx0)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if ptr != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpARM64UBFX {
			break
		}
		if x_1.AuxInt != arm64BFAuxInt(8, 24) {
			break
		}
		if w != x_1.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr1 := v_0.Args[0]
		idx1 := v_0.Args[1]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr0 := x.Args[0]
		idx0 := x.Args[1]
		x_2 := x.Args[2]
		if x_2.Op != OpARM64UBFX {
			break
		}
		if x_2.AuxInt != arm64BFAuxInt(8, 24) {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr0)
		v.AddArg(idx0)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBstore_40(v *Value) bool {
	b := v.Block
	_ = b

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstore {
			break
		}
		if x.AuxInt != i-1 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if ptr != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpARM64SRLconst {
			break
		}
		if x_1.AuxInt != 8 {
			break
		}
		x_1_0 := x_1.Args[0]
		if x_1_0.Op != OpARM64MOVDreg {
			break
		}
		if w != x_1_0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr1 := v_0.Args[0]
		idx1 := v_0.Args[1]
		w := v.Args[1]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		ptr0 := x.Args[0]
		idx0 := x.Args[1]
		x_2 := x.Args[2]
		if x_2.Op != OpARM64SRLconst {
			break
		}
		if x_2.AuxInt != 8 {
			break
		}
		x_2_0 := x_2.Args[0]
		if x_2_0.Op != OpARM64MOVDreg {
			break
		}
		if w != x_2_0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr0)
		v.AddArg(idx0)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBstoreidx_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVBstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVBstore)
		v.AuxInt = c
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVDconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		mem := v.Args[3]
		v.reset(OpARM64MOVBstorezeroidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVBreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVBstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVBUreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVBstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVHreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVBstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVHUreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVBstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVBstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWUreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVBstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADDconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64SRLconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		if ptr != x.Args[0] {
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
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBstoreidx_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADDconst {
			break
		}
		if v_1.AuxInt != 3 {
			break
		}
		idx := v_1.Args[0]
		w := v.Args[2]
		x0 := v.Args[3]
		if x0.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x0.Args[3]
		if ptr != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64ADDconst {
			break
		}
		if x0_1.AuxInt != 2 {
			break
		}
		if idx != x0_1.Args[0] {
			break
		}
		x0_2 := x0.Args[2]
		if x0_2.Op != OpARM64UBFX {
			break
		}
		if x0_2.AuxInt != arm64BFAuxInt(8, 24) {
			break
		}
		if w != x0_2.Args[0] {
			break
		}
		x1 := x0.Args[3]
		if x1.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x1.Args[3]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 1 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		x1_2 := x1.Args[2]
		if x1_2.Op != OpARM64UBFX {
			break
		}
		if x1_2.AuxInt != arm64BFAuxInt(16, 16) {
			break
		}
		if w != x1_2.Args[0] {
			break
		}
		x2 := x1.Args[3]
		if x2.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x2.Args[3]
		if ptr != x2.Args[0] {
			break
		}
		if idx != x2.Args[1] {
			break
		}
		x2_2 := x2.Args[2]
		if x2_2.Op != OpARM64UBFX {
			break
		}
		if x2_2.AuxInt != arm64BFAuxInt(24, 8) {
			break
		}
		if w != x2_2.Args[0] {
			break
		}
		mem := x2.Args[3]
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x0 := v.Args[3]
		if x0.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x0.Args[3]
		if ptr != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64ADDconst {
			break
		}
		if x0_1.AuxInt != 1 {
			break
		}
		if idx != x0_1.Args[0] {
			break
		}
		x0_2 := x0.Args[2]
		if x0_2.Op != OpARM64UBFX {
			break
		}
		if x0_2.AuxInt != arm64BFAuxInt(8, 24) {
			break
		}
		if w != x0_2.Args[0] {
			break
		}
		x1 := x0.Args[3]
		if x1.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x1.Args[3]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 2 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		x1_2 := x1.Args[2]
		if x1_2.Op != OpARM64UBFX {
			break
		}
		if x1_2.AuxInt != arm64BFAuxInt(16, 16) {
			break
		}
		if w != x1_2.Args[0] {
			break
		}
		x2 := x1.Args[3]
		if x2.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x2.Args[3]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 3 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		x2_2 := x2.Args[2]
		if x2_2.Op != OpARM64UBFX {
			break
		}
		if x2_2.AuxInt != arm64BFAuxInt(24, 8) {
			break
		}
		if w != x2_2.Args[0] {
			break
		}
		mem := x2.Args[3]
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && clobber(x0) && clobber(x1) && clobber(x2)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADDconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		if ptr != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpARM64UBFX {
			break
		}
		if x_2.AuxInt != arm64BFAuxInt(8, 8) {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, w.Type)
		v0.AddArg(w)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		w := v.Args[2]
		x := v.Args[3]
		if x.Op != OpARM64MOVBstoreidx {
			break
		}
		_ = x.Args[3]
		if ptr != x.Args[0] {
			break
		}
		x_1 := x.Args[1]
		if x_1.Op != OpARM64ADDconst {
			break
		}
		if x_1.AuxInt != 1 {
			break
		}
		if idx != x_1.Args[0] {
			break
		}
		x_2 := x.Args[2]
		if x_2.Op != OpARM64UBFX {
			break
		}
		if x_2.AuxInt != arm64BFAuxInt(8, 8) {
			break
		}
		if w != x_2.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBstorezero_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVBstorezero)
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
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVBstorezero)
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
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVBstorezeroidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		ptr0 := v.Args[0]
		x := v.Args[1]
		if x.Op != OpARM64MOVBstorezero {
			break
		}
		j := x.AuxInt
		if x.Aux != s {
			break
		}
		_ = x.Args[1]
		ptr1 := x.Args[0]
		mem := x.Args[1]
		if !(x.Uses == 1 && areAdjacentOffsets(i, j, 1) && is32Bit(min(i, j)) && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstorezero)
		v.AuxInt = min(i, j)
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 1 {
			break
		}
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		x := v.Args[1]
		if x.Op != OpARM64MOVBstorezeroidx {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstorezeroidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVBstorezeroidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVBstorezero)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVBstorezero)
		v.AuxInt = c
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADDconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVBstorezeroidx {
			break
		}
		_ = x.Args[2]
		if ptr != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVHstorezeroidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVDload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVDloadidx)
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
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 3 {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVDloadidx8)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVDload)
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
		if v_1.Op != OpARM64MOVDstorezero {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVDloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVDload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVDload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		if v_1.AuxInt != 3 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVDloadidx8)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 3 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVDloadidx8)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVDstorezeroidx {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2) || isSamePtr(ptr, idx2) && isSamePtr(idx, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVDloadidx8_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVDload)
		v.AuxInt = c << 3
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVDstorezeroidx8 {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVDreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpARM64MOVDnop)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVDstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64FMOVDfpgp {
			break
		}
		val := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64FMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVDstore)
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
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVDstoreidx)
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
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 3 {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVDstoreidx8)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVDstore)
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
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpARM64MOVDstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVDstoreidx_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVDstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVDstore)
		v.AuxInt = c
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		if v_1.AuxInt != 3 {
			break
		}
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVDstoreidx8)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 3 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVDstoreidx8)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVDconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		mem := v.Args[3]
		v.reset(OpARM64MOVDstorezeroidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVDstoreidx8_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVDstore)
		v.AuxInt = c << 3
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVDconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		mem := v.Args[3]
		v.reset(OpARM64MOVDstorezeroidx8)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVDstorezero_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVDstorezero)
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
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVDstorezero)
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
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVDstorezeroidx)
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
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 3 {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVDstorezeroidx8)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		ptr0 := v.Args[0]
		x := v.Args[1]
		if x.Op != OpARM64MOVDstorezero {
			break
		}
		j := x.AuxInt
		if x.Aux != s {
			break
		}
		_ = x.Args[1]
		ptr1 := x.Args[0]
		mem := x.Args[1]
		if !(x.Uses == 1 && areAdjacentOffsets(i, j, 8) && is32Bit(min(i, j)) && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVQstorezero)
		v.AuxInt = min(i, j)
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 8 {
			break
		}
		s := v.Aux
		_ = v.Args[1]
		p0 := v.Args[0]
		if p0.Op != OpARM64ADD {
			break
		}
		_ = p0.Args[1]
		ptr0 := p0.Args[0]
		idx0 := p0.Args[1]
		x := v.Args[1]
		if x.Op != OpARM64MOVDstorezeroidx {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVQstorezero)
		v.AuxInt = 0
		v.Aux = s
		v.AddArg(p0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 8 {
			break
		}
		s := v.Aux
		_ = v.Args[1]
		p0 := v.Args[0]
		if p0.Op != OpARM64ADDshiftLL {
			break
		}
		if p0.AuxInt != 3 {
			break
		}
		_ = p0.Args[1]
		ptr0 := p0.Args[0]
		idx0 := p0.Args[1]
		x := v.Args[1]
		if x.Op != OpARM64MOVDstorezeroidx8 {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && s == nil && isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVQstorezero)
		v.AuxInt = 0
		v.Aux = s
		v.AddArg(p0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVDstorezeroidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVDstorezero)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVDstorezero)
		v.AuxInt = c
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		if v_1.AuxInt != 3 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVDstorezeroidx8)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 3 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVDstorezeroidx8)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVDstorezeroidx8_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVDstorezero)
		v.AuxInt = c << 3
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHUload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVHUloadidx)
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
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVHUloadidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVHUload)
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
		if v_1.Op != OpARM64MOVHstorezero {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHUloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVHUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVHUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVHUloadidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADD {
			break
		}
		_ = v_1.Args[1]
		idx := v_1.Args[0]
		if idx != v_1.Args[1] {
			break
		}
		mem := v.Args[2]
		v.reset(OpARM64MOVHUloadidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		idx := v_0.Args[0]
		if idx != v_0.Args[1] {
			break
		}
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVHUloadidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVHstorezeroidx {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2) || isSamePtr(ptr, idx2) && isSamePtr(idx, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHUloadidx2_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVHUload)
		v.AuxInt = c << 1
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVHstorezeroidx2 {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHUreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHUloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHUloadidx2 {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHUreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARM64ANDconst)
		v.AuxInt = c & (1<<16 - 1)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(uint16(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<16-1, sc)) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(sc, arm64BFWidth(1<<16-1, sc))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHUreg_10(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<16-1, 0)) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(sc, 16)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVHloadidx)
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
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVHloadidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVHload)
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
		if v_1.Op != OpARM64MOVHstorezero {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVHload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVHload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVHloadidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADD {
			break
		}
		_ = v_1.Args[1]
		idx := v_1.Args[0]
		if idx != v_1.Args[1] {
			break
		}
		mem := v.Args[2]
		v.reset(OpARM64MOVHloadidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		idx := v_0.Args[0]
		if idx != v_0.Args[1] {
			break
		}
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVHloadidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVHstorezeroidx {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2) || isSamePtr(ptr, idx2) && isSamePtr(idx, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHloadidx2_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVHload)
		v.AuxInt = c << 1
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVHstorezeroidx2 {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHloadidx2 {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHreg_10(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(int16(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		lc := v_0.AuxInt
		x := v_0.Args[0]
		if !(lc < 16) {
			break
		}
		v.reset(OpARM64SBFIZ)
		v.AuxInt = arm64BFAuxInt(lc, 16-lc)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVHstore)
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
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVHstoreidx)
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
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVHstoreidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVHstore)
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
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpARM64MOVHstorezero)
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
		if v_1.Op != OpARM64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVHstore)
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
		if v_1.Op != OpARM64MOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVHstore)
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
		if v_1.Op != OpARM64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVHstore)
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
		if v_1.Op != OpARM64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHstore_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		if v.AuxInt != 2 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstoreidx2 {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr1)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, idx1.Type)
		v0.AuxInt = 1
		v0.AddArg(idx1)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64UBFX {
			break
		}
		if v_1.AuxInt != arm64BFAuxInt(16, 16) {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64UBFX {
			break
		}
		if v_1.AuxInt != arm64BFAuxInt(16, 16) {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64UBFX {
			break
		}
		if v_1.AuxInt != arm64BFAuxInt(16, 16) {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstoreidx2 {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr1)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, idx1.Type)
		v0.AuxInt = 1
		v0.AddArg(idx1)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpARM64MOVDreg {
			break
		}
		w := v_1_0.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpARM64MOVDreg {
			break
		}
		w := v_1_0.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpARM64MOVDreg {
			break
		}
		w := v_1_0.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstoreidx2 {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr1)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, idx1.Type)
		v0.AuxInt = 1
		v0.AddArg(idx1)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstore {
			break
		}
		if x.AuxInt != i-2 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		w0 := x.Args[1]
		if w0.Op != OpARM64SRLconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		w0 := x.Args[2]
		if w0.Op != OpARM64SRLconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHstore_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		if v.AuxInt != 2 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstoreidx2 {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		w0 := x.Args[2]
		if w0.Op != OpARM64SRLconst {
			break
		}
		if w0.AuxInt != j-16 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr1)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, idx1.Type)
		v0.AuxInt = 1
		v0.AddArg(idx1)
		v.AddArg(v0)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHstoreidx_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstore)
		v.AuxInt = c
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADD {
			break
		}
		_ = v_1.Args[1]
		idx := v_1.Args[0]
		if idx != v_1.Args[1] {
			break
		}
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		idx := v_0.Args[0]
		if idx != v_0.Args[1] {
			break
		}
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVDconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		mem := v.Args[3]
		v.reset(OpARM64MOVHstorezeroidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVHreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVHUreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHstoreidx_10(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWUreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADDconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64SRLconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpARM64MOVHstoreidx {
			break
		}
		_ = x.Args[3]
		if ptr != x.Args[0] {
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
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHstoreidx2_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstore)
		v.AuxInt = c << 1
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVDconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		mem := v.Args[3]
		v.reset(OpARM64MOVHstorezeroidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVHreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVHUreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWUreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVHstoreidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHstorezero_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVHstorezero)
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
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVHstorezero)
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
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVHstorezeroidx)
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
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVHstorezeroidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		ptr0 := v.Args[0]
		x := v.Args[1]
		if x.Op != OpARM64MOVHstorezero {
			break
		}
		j := x.AuxInt
		if x.Aux != s {
			break
		}
		_ = x.Args[1]
		ptr1 := x.Args[0]
		mem := x.Args[1]
		if !(x.Uses == 1 && areAdjacentOffsets(i, j, 2) && is32Bit(min(i, j)) && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstorezero)
		v.AuxInt = min(i, j)
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		x := v.Args[1]
		if x.Op != OpARM64MOVHstorezeroidx {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstorezeroidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 2 {
			break
		}
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		x := v.Args[1]
		if x.Op != OpARM64MOVHstorezeroidx2 {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && s == nil && isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstorezeroidx)
		v.AddArg(ptr1)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, idx1.Type)
		v0.AuxInt = 1
		v0.AddArg(idx1)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHstorezeroidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVHstorezero)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVHstorezero)
		v.AuxInt = c
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVHstorezeroidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADD {
			break
		}
		_ = v_1.Args[1]
		idx := v_1.Args[0]
		if idx != v_1.Args[1] {
			break
		}
		mem := v.Args[2]
		v.reset(OpARM64MOVHstorezeroidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVHstorezeroidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		idx := v_0.Args[0]
		if idx != v_0.Args[1] {
			break
		}
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVHstorezeroidx2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADDconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVHstorezeroidx {
			break
		}
		_ = x.Args[2]
		if ptr != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVWstorezeroidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVHstorezeroidx2_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVHstorezero)
		v.AuxInt = c << 1
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVQstorezero_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVQstorezero)
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
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVQstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWUload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVWUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVWUloadidx)
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
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVWUloadidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVWUload)
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
		if v_1.Op != OpARM64MOVWstorezero {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWUloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVWUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVWUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVWUloadidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVWUloadidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWstorezeroidx {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2) || isSamePtr(ptr, idx2) && isSamePtr(idx, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWUloadidx4_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVWUload)
		v.AuxInt = c << 2
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWstorezeroidx4 {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWUreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVWUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHUloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVWUloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHUloadidx2 {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVWUloadidx4 {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHUreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWUreg_10(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVWUreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARM64ANDconst)
		v.AuxInt = c & (1<<32 - 1)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(uint32(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<32-1, sc)) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(sc, arm64BFWidth(1<<32-1, sc))
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<32-1, 0)) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(sc, 32)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVWloadidx)
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
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVWloadidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVWload)
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
		if v_1.Op != OpARM64MOVWstorezero {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[1]
		ptr2 := v_1.Args[0]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVWload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVWload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVWloadidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVWloadidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWstorezeroidx {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2) || isSamePtr(ptr, idx2) && isSamePtr(idx, ptr2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWloadidx4_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVWload)
		v.AuxInt = c << 2
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWstorezeroidx4 {
			break
		}
		_ = v_2.Args[2]
		ptr2 := v_2.Args[0]
		idx2 := v_2.Args[1]
		if !(isSamePtr(ptr, ptr2) && isSamePtr(idx, idx2)) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVWload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHUloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVWloadidx {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWreg_10(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHloadidx2 {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHUloadidx2 {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVWloadidx4 {
			break
		}
		_ = x.Args[2]
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVBUreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVHreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARM64MOVWreg {
			break
		}
		v.reset(OpARM64MOVDreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(int32(c))
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		lc := v_0.AuxInt
		x := v_0.Args[0]
		if !(lc < 32) {
			break
		}
		v.reset(OpARM64SBFIZ)
		v.AuxInt = arm64BFAuxInt(lc, 32-lc)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVWstore)
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
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVWstoreidx)
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
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVWstoreidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVWstore)
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
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		mem := v.Args[2]
		v.reset(OpARM64MOVWstorezero)
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
		if v_1.Op != OpARM64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVWstore)
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
		if v_1.Op != OpARM64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 32 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVWstore {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		if w != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVDstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 32 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVWstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVDstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != 32 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVWstoreidx4 {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		if w != x.Args[2] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVDstoreidx)
		v.AddArg(ptr1)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, idx1.Type)
		v0.AuxInt = 2
		v0.AddArg(idx1)
		v.AddArg(v0)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWstore_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		ptr0 := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVWstore {
			break
		}
		if x.AuxInt != i-4 {
			break
		}
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		w0 := x.Args[1]
		if w0.Op != OpARM64SRLconst {
			break
		}
		if w0.AuxInt != j-32 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVDstore)
		v.AuxInt = i - 4
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVWstoreidx {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		w0 := x.Args[2]
		if w0.Op != OpARM64SRLconst {
			break
		}
		if w0.AuxInt != j-32 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVDstoreidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		s := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVWstoreidx4 {
			break
		}
		_ = x.Args[3]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		w0 := x.Args[2]
		if w0.Op != OpARM64SRLconst {
			break
		}
		if w0.AuxInt != j-32 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[3]
		if !(x.Uses == 1 && s == nil && isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVDstoreidx)
		v.AddArg(ptr1)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, idx1.Type)
		v0.AuxInt = 2
		v0.AddArg(idx1)
		v.AddArg(v0)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWstoreidx_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVWstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVWstore)
		v.AuxInt = c
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVWstoreidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVWstoreidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVDconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		mem := v.Args[3]
		v.reset(OpARM64MOVWstorezeroidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWUreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVWstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADDconst {
			break
		}
		if v_1.AuxInt != 4 {
			break
		}
		idx := v_1.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64SRLconst {
			break
		}
		if v_2.AuxInt != 32 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != OpARM64MOVWstoreidx {
			break
		}
		_ = x.Args[3]
		if ptr != x.Args[0] {
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
		v.reset(OpARM64MOVDstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(w)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWstoreidx4_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64MOVWstore)
		v.AuxInt = c << 2
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVDconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		mem := v.Args[3]
		v.reset(OpARM64MOVWstorezeroidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVWstoreidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVWUreg {
			break
		}
		x := v_2.Args[0]
		mem := v.Args[3]
		v.reset(OpARM64MOVWstoreidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWstorezero_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVWstorezero)
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
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64MOVWstorezero)
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
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVWstorezeroidx)
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
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(off == 0 && sym == nil) {
			break
		}
		v.reset(OpARM64MOVWstorezeroidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		ptr0 := v.Args[0]
		x := v.Args[1]
		if x.Op != OpARM64MOVWstorezero {
			break
		}
		j := x.AuxInt
		if x.Aux != s {
			break
		}
		_ = x.Args[1]
		ptr1 := x.Args[0]
		mem := x.Args[1]
		if !(x.Uses == 1 && areAdjacentOffsets(i, j, 4) && is32Bit(min(i, j)) && isSamePtr(ptr0, ptr1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVDstorezero)
		v.AuxInt = min(i, j)
		v.Aux = s
		v.AddArg(ptr0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADD {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		x := v.Args[1]
		if x.Op != OpARM64MOVWstorezeroidx {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && s == nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVDstorezeroidx)
		v.AddArg(ptr1)
		v.AddArg(idx1)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDshiftLL {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		_ = v_0.Args[1]
		ptr0 := v_0.Args[0]
		idx0 := v_0.Args[1]
		x := v.Args[1]
		if x.Op != OpARM64MOVWstorezeroidx4 {
			break
		}
		_ = x.Args[2]
		ptr1 := x.Args[0]
		idx1 := x.Args[1]
		mem := x.Args[2]
		if !(x.Uses == 1 && s == nil && isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVDstorezeroidx)
		v.AddArg(ptr1)
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, idx1.Type)
		v0.AuxInt = 2
		v0.AddArg(idx1)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWstorezeroidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVWstorezero)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVWstorezero)
		v.AuxInt = c
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARM64MOVWstorezeroidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVWstorezeroidx4)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ADDconst {
			break
		}
		if v_1.AuxInt != 4 {
			break
		}
		idx := v_1.Args[0]
		x := v.Args[2]
		if x.Op != OpARM64MOVWstorezeroidx {
			break
		}
		_ = x.Args[2]
		if ptr != x.Args[0] {
			break
		}
		if idx != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(OpARM64MOVDstorezeroidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MOVWstorezeroidx4_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARM64MOVWstorezero)
		v.AuxInt = c << 2
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MUL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64NEG {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpARM64MNEG)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64NEG {
			break
		}
		x := v_1.Args[0]
		v.reset(OpARM64MNEG)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		if v_1.AuxInt != -1 {
			break
		}
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		if v_0.AuxInt != -1 {
			break
		}
		x := v.Args[1]
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
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
		if v_0.Op != OpARM64MOVDconst {
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
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MUL_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c-1) && c >= 3) {
			break
		}
		v.reset(OpARM64ADDshiftLL)
		v.AuxInt = log2(c - 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c-1) && c >= 3) {
			break
		}
		v.reset(OpARM64ADDshiftLL)
		v.AuxInt = log2(c - 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c+1) && c >= 7) {
			break
		}
		v.reset(OpARM64ADDshiftLL)
		v.AuxInt = log2(c + 1)
		v0 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c+1) && c >= 7) {
			break
		}
		v.reset(OpARM64ADDshiftLL)
		v.AuxInt = log2(c + 1)
		v0 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%3 == 0 && isPowerOfTwo(c/3)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 1
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%3 == 0 && isPowerOfTwo(c/3)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 1
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%5 == 0 && isPowerOfTwo(c/5)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 5)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 2
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%5 == 0 && isPowerOfTwo(c/5)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 5)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 2
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%7 == 0 && isPowerOfTwo(c/7)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 3
		v1 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%7 == 0 && isPowerOfTwo(c/7)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 3
		v1 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MUL_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%9 == 0 && isPowerOfTwo(c/9)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 9)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%9 == 0 && isPowerOfTwo(c/9)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 9)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = c * d
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = c * d
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MULW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64NEG {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpARM64MNEGW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64NEG {
			break
		}
		x := v_1.Args[0]
		v.reset(OpARM64MNEGW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		if !(int32(c) == 0) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(int32(c) == 1) {
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
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(int32(c) == 1) {
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
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MULW_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARM64ADDshiftLL)
		v.AuxInt = log2(c - 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARM64ADDshiftLL)
		v.AuxInt = log2(c - 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARM64ADDshiftLL)
		v.AuxInt = log2(c + 1)
		v0 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARM64ADDshiftLL)
		v.AuxInt = log2(c + 1)
		v0 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 1
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 1
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 5)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 2
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 5)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 2
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 3
		v1 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 3
		v1 := b.NewValue0(v.Pos, OpARM64NEG, x.Type)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MULW_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 9)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SLLconst)
		v.AuxInt = log2(c / 9)
		v0 := b.NewValue0(v.Pos, OpARM64ADDshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(int32(c) * int32(d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(int32(c) * int32(d))
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64MVN_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = ^c
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64NEG_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MUL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpARM64MNEG)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MULW {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpARM64MNEGW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -c
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64NotEqual_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagEQ {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagLT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_ULT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64FlagGT_UGT {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARM64InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARM64NotEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64OR_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ORconst)
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
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MVN {
			break
		}
		y := v_1.Args[0]
		v.reset(OpARM64ORN)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MVN {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARM64ORN)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ORshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ORshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ORshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ORshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ORshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64OR_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ORshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFIZ {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ANDconst {
			break
		}
		ac := v_1.AuxInt
		y := v_1.Args[0]
		if !(ac == ^((1<<uint(getARM64BFwidth(bfc)) - 1) << uint(getARM64BFlsb(bfc)))) {
			break
		}
		v.reset(OpARM64BFI)
		v.AuxInt = bfc
		v.AddArg(y)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ANDconst {
			break
		}
		ac := v_0.AuxInt
		y := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64UBFIZ {
			break
		}
		bfc := v_1.AuxInt
		x := v_1.Args[0]
		if !(ac == ^((1<<uint(getARM64BFwidth(bfc)) - 1) << uint(getARM64BFlsb(bfc)))) {
			break
		}
		v.reset(OpARM64BFI)
		v.AuxInt = bfc
		v.AddArg(y)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFX {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64ANDconst {
			break
		}
		ac := v_1.AuxInt
		y := v_1.Args[0]
		if !(ac == ^(1<<uint(getARM64BFwidth(bfc)) - 1)) {
			break
		}
		v.reset(OpARM64BFXIL)
		v.AuxInt = bfc
		v.AddArg(y)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ANDconst {
			break
		}
		ac := v_0.AuxInt
		y := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64UBFX {
			break
		}
		bfc := v_1.AuxInt
		x := v_1.Args[0]
		if !(ac == ^(1<<uint(getARM64BFwidth(bfc)) - 1)) {
			break
		}
		v.reset(OpARM64BFXIL)
		v.AuxInt = bfc
		v.AddArg(y)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		i3 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
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
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		i1 := x2.AuxInt
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
		y3 := v.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		i0 := x3.AuxInt
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
		if !(i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.Aux = s
		v1 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v1.AuxInt = i0
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		y3 := v.Args[0]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		i0 := x3.AuxInt
		s := x3.Aux
		_ = x3.Args[1]
		p := x3.Args[0]
		mem := x3.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
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
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
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
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		i1 := x2.AuxInt
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
		if !(i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.Aux = s
		v1 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v1.AuxInt = i0
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		if x0.AuxInt != 3 {
			break
		}
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 2 {
			break
		}
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
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 1 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		p1 := x2.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x2.Args[1] {
			break
		}
		y3 := v.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		ptr0 := x3.Args[0]
		idx0 := x3.Args[1]
		if mem != x3.Args[2] {
			break
		}
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr0)
		v0.AddArg(idx0)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		y3 := v.Args[0]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		ptr0 := x3.Args[0]
		idx0 := x3.Args[1]
		mem := x3.Args[2]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		if x0.AuxInt != 3 {
			break
		}
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		if mem != x0.Args[1] {
			break
		}
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 2 {
			break
		}
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
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 1 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		p1 := x2.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x2.Args[1] {
			break
		}
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr0)
		v0.AddArg(idx0)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr := x0.Args[0]
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64ADDconst {
			break
		}
		if x0_1.AuxInt != 3 {
			break
		}
		idx := x0_1.Args[0]
		mem := x0.Args[2]
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 2 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 1 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		y3 := v.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		if ptr != x3.Args[0] {
			break
		}
		if idx != x3.Args[1] {
			break
		}
		if mem != x3.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64OR_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		y3 := v.Args[0]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		ptr := x3.Args[0]
		idx := x3.Args[1]
		mem := x3.Args[2]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		if ptr != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64ADDconst {
			break
		}
		if x0_1.AuxInt != 3 {
			break
		}
		if idx != x0_1.Args[0] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 2 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 1 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		i7 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		i6 := x1.AuxInt
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
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		i5 := x2.AuxInt
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
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		i4 := x3.AuxInt
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
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
			break
		}
		i3 := x4.AuxInt
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
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUload {
			break
		}
		i2 := x5.AuxInt
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
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUload {
			break
		}
		i1 := x6.AuxInt
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
		y7 := v.Args[1]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUload {
			break
		}
		i0 := x7.AuxInt
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
		if !(i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.Aux = s
		v1 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v1.AuxInt = i0
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		y7 := v.Args[0]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUload {
			break
		}
		i0 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		i7 := x0.AuxInt
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
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		i6 := x1.AuxInt
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
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		i5 := x2.AuxInt
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
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		i4 := x3.AuxInt
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
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
			break
		}
		i3 := x4.AuxInt
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
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUload {
			break
		}
		i2 := x5.AuxInt
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
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUload {
			break
		}
		i1 := x6.AuxInt
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
		if !(i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.Aux = s
		v1 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v1.AuxInt = i0
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		if x0.AuxInt != 7 {
			break
		}
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 6 {
			break
		}
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
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 5 {
			break
		}
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
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		if x3.AuxInt != 4 {
			break
		}
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
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
			break
		}
		if x4.AuxInt != 3 {
			break
		}
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
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUload {
			break
		}
		if x5.AuxInt != 2 {
			break
		}
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
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUload {
			break
		}
		if x6.AuxInt != 1 {
			break
		}
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		p1 := x6.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x6.Args[1] {
			break
		}
		y7 := v.Args[1]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x7.Args[2]
		ptr0 := x7.Args[0]
		idx0 := x7.Args[1]
		if mem != x7.Args[2] {
			break
		}
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr0)
		v0.AddArg(idx0)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		y7 := v.Args[0]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x7.Args[2]
		ptr0 := x7.Args[0]
		idx0 := x7.Args[1]
		mem := x7.Args[2]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		if x0.AuxInt != 7 {
			break
		}
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		if mem != x0.Args[1] {
			break
		}
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 6 {
			break
		}
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
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 5 {
			break
		}
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
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		if x3.AuxInt != 4 {
			break
		}
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
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
			break
		}
		if x4.AuxInt != 3 {
			break
		}
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
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUload {
			break
		}
		if x5.AuxInt != 2 {
			break
		}
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
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUload {
			break
		}
		if x6.AuxInt != 1 {
			break
		}
		if x6.Aux != s {
			break
		}
		_ = x6.Args[1]
		p1 := x6.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x6.Args[1] {
			break
		}
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr0)
		v0.AddArg(idx0)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr := x0.Args[0]
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64ADDconst {
			break
		}
		if x0_1.AuxInt != 7 {
			break
		}
		idx := x0_1.Args[0]
		mem := x0.Args[2]
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 6 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 5 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		if ptr != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpARM64ADDconst {
			break
		}
		if x3_1.AuxInt != 4 {
			break
		}
		if idx != x3_1.Args[0] {
			break
		}
		if mem != x3.Args[2] {
			break
		}
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x4.Args[2]
		if ptr != x4.Args[0] {
			break
		}
		x4_1 := x4.Args[1]
		if x4_1.Op != OpARM64ADDconst {
			break
		}
		if x4_1.AuxInt != 3 {
			break
		}
		if idx != x4_1.Args[0] {
			break
		}
		if mem != x4.Args[2] {
			break
		}
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x5.Args[2]
		if ptr != x5.Args[0] {
			break
		}
		x5_1 := x5.Args[1]
		if x5_1.Op != OpARM64ADDconst {
			break
		}
		if x5_1.AuxInt != 2 {
			break
		}
		if idx != x5_1.Args[0] {
			break
		}
		if mem != x5.Args[2] {
			break
		}
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x6.Args[2]
		if ptr != x6.Args[0] {
			break
		}
		x6_1 := x6.Args[1]
		if x6_1.Op != OpARM64ADDconst {
			break
		}
		if x6_1.AuxInt != 1 {
			break
		}
		if idx != x6_1.Args[0] {
			break
		}
		if mem != x6.Args[2] {
			break
		}
		y7 := v.Args[1]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x7.Args[2]
		if ptr != x7.Args[0] {
			break
		}
		if idx != x7.Args[1] {
			break
		}
		if mem != x7.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		y7 := v.Args[0]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x7.Args[2]
		ptr := x7.Args[0]
		idx := x7.Args[1]
		mem := x7.Args[2]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		if ptr != x0.Args[0] {
			break
		}
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64ADDconst {
			break
		}
		if x0_1.AuxInt != 7 {
			break
		}
		if idx != x0_1.Args[0] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 6 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 5 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		if ptr != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpARM64ADDconst {
			break
		}
		if x3_1.AuxInt != 4 {
			break
		}
		if idx != x3_1.Args[0] {
			break
		}
		if mem != x3.Args[2] {
			break
		}
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x4.Args[2]
		if ptr != x4.Args[0] {
			break
		}
		x4_1 := x4.Args[1]
		if x4_1.Op != OpARM64ADDconst {
			break
		}
		if x4_1.AuxInt != 3 {
			break
		}
		if idx != x4_1.Args[0] {
			break
		}
		if mem != x4.Args[2] {
			break
		}
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x5.Args[2]
		if ptr != x5.Args[0] {
			break
		}
		x5_1 := x5.Args[1]
		if x5_1.Op != OpARM64ADDconst {
			break
		}
		if x5_1.AuxInt != 2 {
			break
		}
		if idx != x5_1.Args[0] {
			break
		}
		if mem != x5.Args[2] {
			break
		}
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x6.Args[2]
		if ptr != x6.Args[0] {
			break
		}
		x6_1 := x6.Args[1]
		if x6_1.Op != OpARM64ADDconst {
			break
		}
		if x6_1.AuxInt != 1 {
			break
		}
		if idx != x6_1.Args[0] {
			break
		}
		if mem != x6.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
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
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
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
		y3 := v.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
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
		if !(i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWUload, t)
		v1.Aux = s
		v2 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v2.AuxInt = i0
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		y3 := v.Args[0]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		i3 := x3.AuxInt
		s := x3.Aux
		_ = x3.Args[1]
		p := x3.Args[0]
		mem := x3.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
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
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
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
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
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
		if !(i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWUload, t)
		v1.Aux = s
		v2 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v2.AuxInt = i0
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
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr0 := x0.Args[0]
		idx0 := x0.Args[1]
		mem := x0.Args[2]
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 1 {
			break
		}
		s := x1.Aux
		_ = x1.Args[1]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x1.Args[1] {
			break
		}
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 2 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		p := x2.Args[0]
		if mem != x2.Args[1] {
			break
		}
		y3 := v.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		if x3.AuxInt != 3 {
			break
		}
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
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v1.AddArg(ptr0)
		v1.AddArg(idx0)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64OR_30(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		y3 := v.Args[0]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		if x3.AuxInt != 3 {
			break
		}
		s := x3.Aux
		_ = x3.Args[1]
		p := x3.Args[0]
		mem := x3.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr0 := x0.Args[0]
		idx0 := x0.Args[1]
		if mem != x0.Args[2] {
			break
		}
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 1 {
			break
		}
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x1.Args[1] {
			break
		}
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 2 {
			break
		}
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
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v1.AddArg(ptr0)
		v1.AddArg(idx0)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 1 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 2 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		y3 := v.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		if ptr != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpARM64ADDconst {
			break
		}
		if x3_1.AuxInt != 3 {
			break
		}
		if idx != x3_1.Args[0] {
			break
		}
		if mem != x3.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v1.AddArg(ptr)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		y3 := v.Args[0]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		ptr := x3.Args[0]
		x3_1 := x3.Args[1]
		if x3_1.Op != OpARM64ADDconst {
			break
		}
		if x3_1.AuxInt != 3 {
			break
		}
		idx := x3_1.Args[0]
		mem := x3.Args[2]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		s0 := o1.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 24 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		if ptr != x0.Args[0] {
			break
		}
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y1 := o1.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 1 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := o0.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 2 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(o0) && clobber(o1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v1.AddArg(ptr)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
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
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
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
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
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
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
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
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUload {
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
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUload {
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
		y7 := v.Args[1]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUload {
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
		if !(i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64REV, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDload, t)
		v1.Aux = s
		v2 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v2.AuxInt = i0
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		y7 := v.Args[0]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUload {
			break
		}
		i7 := x7.AuxInt
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
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
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
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
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
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
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
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
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
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
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUload {
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
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUload {
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
		if !(i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64REV, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDload, t)
		v1.Aux = s
		v2 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v2.AuxInt = i0
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
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr0 := x0.Args[0]
		idx0 := x0.Args[1]
		mem := x0.Args[2]
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 1 {
			break
		}
		s := x1.Aux
		_ = x1.Args[1]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x1.Args[1] {
			break
		}
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 2 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		p := x2.Args[0]
		if mem != x2.Args[1] {
			break
		}
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		if x3.AuxInt != 3 {
			break
		}
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
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
			break
		}
		if x4.AuxInt != 4 {
			break
		}
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
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUload {
			break
		}
		if x5.AuxInt != 5 {
			break
		}
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
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUload {
			break
		}
		if x6.AuxInt != 6 {
			break
		}
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
		y7 := v.Args[1]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUload {
			break
		}
		if x7.AuxInt != 7 {
			break
		}
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
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64REV, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v1.AddArg(ptr0)
		v1.AddArg(idx0)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		y7 := v.Args[0]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUload {
			break
		}
		if x7.AuxInt != 7 {
			break
		}
		s := x7.Aux
		_ = x7.Args[1]
		p := x7.Args[0]
		mem := x7.Args[1]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr0 := x0.Args[0]
		idx0 := x0.Args[1]
		if mem != x0.Args[2] {
			break
		}
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 1 {
			break
		}
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x1.Args[1] {
			break
		}
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 2 {
			break
		}
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
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		if x3.AuxInt != 3 {
			break
		}
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
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
			break
		}
		if x4.AuxInt != 4 {
			break
		}
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
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUload {
			break
		}
		if x5.AuxInt != 5 {
			break
		}
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
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUload {
			break
		}
		if x6.AuxInt != 6 {
			break
		}
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
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64REV, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v1.AddArg(ptr0)
		v1.AddArg(idx0)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 1 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 2 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		if ptr != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpARM64ADDconst {
			break
		}
		if x3_1.AuxInt != 3 {
			break
		}
		if idx != x3_1.Args[0] {
			break
		}
		if mem != x3.Args[2] {
			break
		}
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x4.Args[2]
		if ptr != x4.Args[0] {
			break
		}
		x4_1 := x4.Args[1]
		if x4_1.Op != OpARM64ADDconst {
			break
		}
		if x4_1.AuxInt != 4 {
			break
		}
		if idx != x4_1.Args[0] {
			break
		}
		if mem != x4.Args[2] {
			break
		}
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x5.Args[2]
		if ptr != x5.Args[0] {
			break
		}
		x5_1 := x5.Args[1]
		if x5_1.Op != OpARM64ADDconst {
			break
		}
		if x5_1.AuxInt != 5 {
			break
		}
		if idx != x5_1.Args[0] {
			break
		}
		if mem != x5.Args[2] {
			break
		}
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x6.Args[2]
		if ptr != x6.Args[0] {
			break
		}
		x6_1 := x6.Args[1]
		if x6_1.Op != OpARM64ADDconst {
			break
		}
		if x6_1.AuxInt != 6 {
			break
		}
		if idx != x6_1.Args[0] {
			break
		}
		if mem != x6.Args[2] {
			break
		}
		y7 := v.Args[1]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x7.Args[2]
		if ptr != x7.Args[0] {
			break
		}
		x7_1 := x7.Args[1]
		if x7_1.Op != OpARM64ADDconst {
			break
		}
		if x7_1.AuxInt != 7 {
			break
		}
		if idx != x7_1.Args[0] {
			break
		}
		if mem != x7.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64REV, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v1.AddArg(ptr)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		y7 := v.Args[0]
		if y7.Op != OpARM64MOVDnop {
			break
		}
		x7 := y7.Args[0]
		if x7.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x7.Args[2]
		ptr := x7.Args[0]
		x7_1 := x7.Args[1]
		if x7_1.Op != OpARM64ADDconst {
			break
		}
		if x7_1.AuxInt != 7 {
			break
		}
		idx := x7_1.Args[0]
		mem := x7.Args[2]
		o0 := v.Args[1]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 8 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 16 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 24 {
			break
		}
		_ = o2.Args[1]
		o3 := o2.Args[0]
		if o3.Op != OpARM64ORshiftLL {
			break
		}
		if o3.AuxInt != 32 {
			break
		}
		_ = o3.Args[1]
		o4 := o3.Args[0]
		if o4.Op != OpARM64ORshiftLL {
			break
		}
		if o4.AuxInt != 40 {
			break
		}
		_ = o4.Args[1]
		o5 := o4.Args[0]
		if o5.Op != OpARM64ORshiftLL {
			break
		}
		if o5.AuxInt != 48 {
			break
		}
		_ = o5.Args[1]
		s0 := o5.Args[0]
		if s0.Op != OpARM64SLLconst {
			break
		}
		if s0.AuxInt != 56 {
			break
		}
		y0 := s0.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		if ptr != x0.Args[0] {
			break
		}
		if idx != x0.Args[1] {
			break
		}
		if mem != x0.Args[2] {
			break
		}
		y1 := o5.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 1 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := o4.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 2 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		y3 := o3.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		if ptr != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpARM64ADDconst {
			break
		}
		if x3_1.AuxInt != 3 {
			break
		}
		if idx != x3_1.Args[0] {
			break
		}
		if mem != x3.Args[2] {
			break
		}
		y4 := o2.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x4.Args[2]
		if ptr != x4.Args[0] {
			break
		}
		x4_1 := x4.Args[1]
		if x4_1.Op != OpARM64ADDconst {
			break
		}
		if x4_1.AuxInt != 4 {
			break
		}
		if idx != x4_1.Args[0] {
			break
		}
		if mem != x4.Args[2] {
			break
		}
		y5 := o1.Args[1]
		if y5.Op != OpARM64MOVDnop {
			break
		}
		x5 := y5.Args[0]
		if x5.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x5.Args[2]
		if ptr != x5.Args[0] {
			break
		}
		x5_1 := x5.Args[1]
		if x5_1.Op != OpARM64ADDconst {
			break
		}
		if x5_1.AuxInt != 5 {
			break
		}
		if idx != x5_1.Args[0] {
			break
		}
		if mem != x5.Args[2] {
			break
		}
		y6 := o0.Args[1]
		if y6.Op != OpARM64MOVDnop {
			break
		}
		x6 := y6.Args[0]
		if x6.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x6.Args[2]
		if ptr != x6.Args[0] {
			break
		}
		x6_1 := x6.Args[1]
		if x6_1.Op != OpARM64ADDconst {
			break
		}
		if x6_1.AuxInt != 6 {
			break
		}
		if idx != x6_1.Args[0] {
			break
		}
		if mem != x6.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && x5.Uses == 1 && x6.Uses == 1 && x7.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && y5.Uses == 1 && y6.Uses == 1 && y7.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && o3.Uses == 1 && o4.Uses == 1 && o5.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(x5) && clobber(x6) && clobber(x7) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(y5) && clobber(y6) && clobber(y7) && clobber(o0) && clobber(o1) && clobber(o2) && clobber(o3) && clobber(o4) && clobber(o5) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4, x5, x6, x7)
		v0 := b.NewValue0(v.Pos, OpARM64REV, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v1.AddArg(ptr)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ORN_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ORconst)
		v.AuxInt = ^c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -1
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ORNshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ORNshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64ORNshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ORNshiftLL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ORconst)
		v.AuxInt = ^int64(uint64(c) << uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -1
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ORNshiftRA_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ORconst)
		v.AuxInt = ^(c >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -1
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ORNshiftRL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ORconst)
		v.AuxInt = ^int64(uint64(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -1
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ORconst_0(v *Value) bool {

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
		v.reset(OpARM64MOVDconst)
		v.AuxInt = -1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = c | d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARM64ORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64ORshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ORconst)
		v.AuxInt = int64(uint64(c) << uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARM64SLLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		if v_0.AuxInt != 64-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARM64RORconst)
		v.AuxInt = 64 - c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFX {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		if !(c < 32 && t.Size(psess.types) == 4 && bfc == arm64BFAuxInt(32-c, c)) {
			break
		}
		v.reset(OpARM64RORWconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		if v_0.AuxInt != 64-c {
			break
		}
		x := v_0.Args[0]
		x2 := v.Args[1]
		v.reset(OpARM64EXTRconst)
		v.AuxInt = 64 - c
		v.AddArg(x2)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFX {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		x2 := v.Args[1]
		if !(c < 32 && t.Size(psess.types) == 4 && bfc == arm64BFAuxInt(32-c, c)) {
			break
		}
		v.reset(OpARM64EXTRWconst)
		v.AuxInt = 32 - c
		v.AddArg(x2)
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFX {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		if v_1.AuxInt != sc {
			break
		}
		y := v_1.Args[0]
		if !(sc == getARM64BFwidth(bfc)) {
			break
		}
		v.reset(OpARM64BFXIL)
		v.AuxInt = bfc
		v.AddArg(y)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[1]
		y0 := v.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := v.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(y0) && clobber(y1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpARM64MOVHUload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.Aux = s
		v1 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v1.AuxInt = i0
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[1]
		y0 := v.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr0 := x0.Args[0]
		idx0 := x0.Args[1]
		mem := x0.Args[2]
		y1 := v.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 1 {
			break
		}
		s := x1.Aux
		_ = x1.Args[1]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x1.Args[1] {
			break
		}
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && mergePoint(b, x0, x1) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x0) && clobber(x1) && clobber(y0) && clobber(y1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpARM64MOVHUloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr0)
		v0.AddArg(idx0)
		v0.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ORshiftLL_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[1]
		y0 := v.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		y1 := v.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 1 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(y0) && clobber(y1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpARM64MOVHUloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 24 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 16 {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpARM64MOVHUload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o0.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
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
		y2 := v.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(y1) && clobber(y2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.Aux = s
		v1 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v1.AuxInt = i0
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 24 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 16 {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpARM64MOVHUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr0 := x0.Args[0]
		idx0 := x0.Args[1]
		mem := x0.Args[2]
		y1 := o0.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 2 {
			break
		}
		s := x1.Aux
		_ = x1.Args[1]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x1.Args[1] {
			break
		}
		y2 := v.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 3 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		p := x2.Args[0]
		if mem != x2.Args[1] {
			break
		}
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(y1) && clobber(y2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr0)
		v0.AddArg(idx0)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 24 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 16 {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpARM64MOVHUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		y1 := o0.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 2 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := v.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 3 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(y1) && clobber(y2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 24 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 16 {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != OpARM64MOVHUloadidx2 {
			break
		}
		_ = x0.Args[2]
		ptr0 := x0.Args[0]
		idx0 := x0.Args[1]
		mem := x0.Args[2]
		y1 := o0.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 2 {
			break
		}
		s := x1.Aux
		_ = x1.Args[1]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADDshiftLL {
			break
		}
		if p1.AuxInt != 1 {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x1.Args[1] {
			break
		}
		y2 := v.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 3 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		p := x2.Args[0]
		if mem != x2.Args[1] {
			break
		}
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(y1) && clobber(y2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr0)
		v1 := b.NewValue0(v.Pos, OpARM64SLLconst, idx0.Type)
		v1.AuxInt = 1
		v1.AddArg(idx0)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 56 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 48 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 40 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 32 {
			break
		}
		_ = o2.Args[1]
		x0 := o2.Args[0]
		if x0.Op != OpARM64MOVWUload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o2.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		i4 := x1.AuxInt
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
		y2 := o1.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		i5 := x2.AuxInt
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
		y3 := o0.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		i6 := x3.AuxInt
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
		y4 := v.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
			break
		}
		i7 := x4.AuxInt
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
		if !(i4 == i0+4 && i5 == i0+5 && i6 == i0+6 && i7 == i0+7 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(o0) && clobber(o1) && clobber(o2)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDload, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.Aux = s
		v1 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v1.AuxInt = i0
		v1.AddArg(p)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 56 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 48 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 40 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 32 {
			break
		}
		_ = o2.Args[1]
		x0 := o2.Args[0]
		if x0.Op != OpARM64MOVWUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr0 := x0.Args[0]
		idx0 := x0.Args[1]
		mem := x0.Args[2]
		y1 := o2.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 4 {
			break
		}
		s := x1.Aux
		_ = x1.Args[1]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x1.Args[1] {
			break
		}
		y2 := o1.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 5 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		p := x2.Args[0]
		if mem != x2.Args[1] {
			break
		}
		y3 := o0.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		if x3.AuxInt != 6 {
			break
		}
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
		y4 := v.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
			break
		}
		if x4.AuxInt != 7 {
			break
		}
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
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(o0) && clobber(o1) && clobber(o2)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr0)
		v0.AddArg(idx0)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 56 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 48 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 40 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 32 {
			break
		}
		_ = o2.Args[1]
		x0 := o2.Args[0]
		if x0.Op != OpARM64MOVWUloadidx4 {
			break
		}
		_ = x0.Args[2]
		ptr0 := x0.Args[0]
		idx0 := x0.Args[1]
		mem := x0.Args[2]
		y1 := o2.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 4 {
			break
		}
		s := x1.Aux
		_ = x1.Args[1]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADDshiftLL {
			break
		}
		if p1.AuxInt != 2 {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x1.Args[1] {
			break
		}
		y2 := o1.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 5 {
			break
		}
		if x2.Aux != s {
			break
		}
		_ = x2.Args[1]
		p := x2.Args[0]
		if mem != x2.Args[1] {
			break
		}
		y3 := o0.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		if x3.AuxInt != 6 {
			break
		}
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
		y4 := v.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
			break
		}
		if x4.AuxInt != 7 {
			break
		}
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
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(o0) && clobber(o1) && clobber(o2)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr0)
		v1 := b.NewValue0(v.Pos, OpARM64SLLconst, idx0.Type)
		v1.AuxInt = 2
		v1.AddArg(idx0)
		v0.AddArg(v1)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 56 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 48 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 40 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 32 {
			break
		}
		_ = o2.Args[1]
		x0 := o2.Args[0]
		if x0.Op != OpARM64MOVWUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		y1 := o2.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 4 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := o1.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 5 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		y3 := o0.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		if ptr != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpARM64ADDconst {
			break
		}
		if x3_1.AuxInt != 6 {
			break
		}
		if idx != x3_1.Args[0] {
			break
		}
		if mem != x3.Args[2] {
			break
		}
		y4 := v.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x4.Args[2]
		if ptr != x4.Args[0] {
			break
		}
		x4_1 := x4.Args[1]
		if x4_1.Op != OpARM64ADDconst {
			break
		}
		if x4_1.AuxInt != 7 {
			break
		}
		if idx != x4_1.Args[0] {
			break
		}
		if mem != x4.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(o0) && clobber(o1) && clobber(o2)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v0.AddArg(ptr)
		v0.AddArg(idx)
		v0.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[1]
		y0 := v.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		i1 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := v.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(y0) && clobber(y1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVHUload, t)
		v1.AuxInt = i0
		v1.Aux = s
		v1.AddArg(p)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ORshiftLL_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[1]
		y0 := v.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUload {
			break
		}
		if x0.AuxInt != 1 {
			break
		}
		s := x0.Aux
		_ = x0.Args[1]
		p1 := x0.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		mem := x0.Args[1]
		y1 := v.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		ptr0 := x1.Args[0]
		idx0 := x1.Args[1]
		if mem != x1.Args[2] {
			break
		}
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && mergePoint(b, x0, x1) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && clobber(x0) && clobber(x1) && clobber(y0) && clobber(y1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVHUloadidx, t)
		v1.AddArg(ptr0)
		v1.AddArg(idx0)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[1]
		y0 := v.Args[0]
		if y0.Op != OpARM64MOVDnop {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr := x0.Args[0]
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64ADDconst {
			break
		}
		if x0_1.AuxInt != 1 {
			break
		}
		idx := x0_1.Args[0]
		mem := x0.Args[2]
		y1 := v.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		if idx != x1.Args[1] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(y0) && clobber(y1)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, OpARM64REV16W, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVHUloadidx, t)
		v1.AddArg(ptr)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 24 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 16 {
			break
		}
		_ = o0.Args[1]
		y0 := o0.Args[0]
		if y0.Op != OpARM64REV16W {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVHUload {
			break
		}
		i2 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o0.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
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
		y2 := v.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
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
		if !(i1 == i0+1 && i2 == i0+2 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWUload, t)
		v1.Aux = s
		v2 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v2.AuxInt = i0
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 24 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 16 {
			break
		}
		_ = o0.Args[1]
		y0 := o0.Args[0]
		if y0.Op != OpARM64REV16W {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVHUload {
			break
		}
		if x0.AuxInt != 2 {
			break
		}
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o0.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 1 {
			break
		}
		if x1.Aux != s {
			break
		}
		_ = x1.Args[1]
		p1 := x1.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x1.Args[1] {
			break
		}
		y2 := v.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		ptr0 := x2.Args[0]
		idx0 := x2.Args[1]
		if mem != x2.Args[2] {
			break
		}
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v1.AddArg(ptr0)
		v1.AddArg(idx0)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 24 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 16 {
			break
		}
		_ = o0.Args[1]
		y0 := o0.Args[0]
		if y0.Op != OpARM64REV16W {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVHUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr := x0.Args[0]
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64ADDconst {
			break
		}
		if x0_1.AuxInt != 2 {
			break
		}
		idx := x0_1.Args[0]
		mem := x0.Args[2]
		y1 := o0.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 1 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := v.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		if idx != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, OpARM64REVW, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWUloadidx, t)
		v1.AddArg(ptr)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 56 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 48 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 40 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 32 {
			break
		}
		_ = o2.Args[1]
		y0 := o2.Args[0]
		if y0.Op != OpARM64REVW {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVWUload {
			break
		}
		i4 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o2.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		i3 := x1.AuxInt
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
		y2 := o1.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
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
		y3 := o0.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		i1 := x3.AuxInt
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
		y4 := v.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUload {
			break
		}
		i0 := x4.AuxInt
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
		if !(i1 == i0+1 && i2 == i0+2 && i3 == i0+3 && i4 == i0+4 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(o0) && clobber(o1) && clobber(o2)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpARM64REV, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDload, t)
		v1.Aux = s
		v2 := b.NewValue0(v.Pos, OpOffPtr, p.Type)
		v2.AuxInt = i0
		v2.AddArg(p)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 56 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 48 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 40 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 32 {
			break
		}
		_ = o2.Args[1]
		y0 := o2.Args[0]
		if y0.Op != OpARM64REVW {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVWUload {
			break
		}
		if x0.AuxInt != 4 {
			break
		}
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		y1 := o2.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUload {
			break
		}
		if x1.AuxInt != 3 {
			break
		}
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
		y2 := o1.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUload {
			break
		}
		if x2.AuxInt != 2 {
			break
		}
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
		y3 := o0.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUload {
			break
		}
		if x3.AuxInt != 1 {
			break
		}
		if x3.Aux != s {
			break
		}
		_ = x3.Args[1]
		p1 := x3.Args[0]
		if p1.Op != OpARM64ADD {
			break
		}
		_ = p1.Args[1]
		ptr1 := p1.Args[0]
		idx1 := p1.Args[1]
		if mem != x3.Args[1] {
			break
		}
		y4 := v.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x4.Args[2]
		ptr0 := x4.Args[0]
		idx0 := x4.Args[1]
		if mem != x4.Args[2] {
			break
		}
		if !(s == nil && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && (isSamePtr(ptr0, ptr1) && isSamePtr(idx0, idx1) || isSamePtr(ptr0, idx1) && isSamePtr(idx0, ptr1)) && isSamePtr(p1, p) && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(o0) && clobber(o1) && clobber(o2)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpARM64REV, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v1.AddArg(ptr0)
		v1.AddArg(idx0)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}

	for {
		t := v.Type
		if v.AuxInt != 56 {
			break
		}
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != OpARM64ORshiftLL {
			break
		}
		if o0.AuxInt != 48 {
			break
		}
		_ = o0.Args[1]
		o1 := o0.Args[0]
		if o1.Op != OpARM64ORshiftLL {
			break
		}
		if o1.AuxInt != 40 {
			break
		}
		_ = o1.Args[1]
		o2 := o1.Args[0]
		if o2.Op != OpARM64ORshiftLL {
			break
		}
		if o2.AuxInt != 32 {
			break
		}
		_ = o2.Args[1]
		y0 := o2.Args[0]
		if y0.Op != OpARM64REVW {
			break
		}
		x0 := y0.Args[0]
		if x0.Op != OpARM64MOVWUloadidx {
			break
		}
		_ = x0.Args[2]
		ptr := x0.Args[0]
		x0_1 := x0.Args[1]
		if x0_1.Op != OpARM64ADDconst {
			break
		}
		if x0_1.AuxInt != 4 {
			break
		}
		idx := x0_1.Args[0]
		mem := x0.Args[2]
		y1 := o2.Args[1]
		if y1.Op != OpARM64MOVDnop {
			break
		}
		x1 := y1.Args[0]
		if x1.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x1.Args[2]
		if ptr != x1.Args[0] {
			break
		}
		x1_1 := x1.Args[1]
		if x1_1.Op != OpARM64ADDconst {
			break
		}
		if x1_1.AuxInt != 3 {
			break
		}
		if idx != x1_1.Args[0] {
			break
		}
		if mem != x1.Args[2] {
			break
		}
		y2 := o1.Args[1]
		if y2.Op != OpARM64MOVDnop {
			break
		}
		x2 := y2.Args[0]
		if x2.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x2.Args[2]
		if ptr != x2.Args[0] {
			break
		}
		x2_1 := x2.Args[1]
		if x2_1.Op != OpARM64ADDconst {
			break
		}
		if x2_1.AuxInt != 2 {
			break
		}
		if idx != x2_1.Args[0] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		y3 := o0.Args[1]
		if y3.Op != OpARM64MOVDnop {
			break
		}
		x3 := y3.Args[0]
		if x3.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x3.Args[2]
		if ptr != x3.Args[0] {
			break
		}
		x3_1 := x3.Args[1]
		if x3_1.Op != OpARM64ADDconst {
			break
		}
		if x3_1.AuxInt != 1 {
			break
		}
		if idx != x3_1.Args[0] {
			break
		}
		if mem != x3.Args[2] {
			break
		}
		y4 := v.Args[1]
		if y4.Op != OpARM64MOVDnop {
			break
		}
		x4 := y4.Args[0]
		if x4.Op != OpARM64MOVBUloadidx {
			break
		}
		_ = x4.Args[2]
		if ptr != x4.Args[0] {
			break
		}
		if idx != x4.Args[1] {
			break
		}
		if mem != x4.Args[2] {
			break
		}
		if !(x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && x3.Uses == 1 && x4.Uses == 1 && y0.Uses == 1 && y1.Uses == 1 && y2.Uses == 1 && y3.Uses == 1 && y4.Uses == 1 && o0.Uses == 1 && o1.Uses == 1 && o2.Uses == 1 && mergePoint(b, x0, x1, x2, x3, x4) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(x3) && clobber(x4) && clobber(y0) && clobber(y1) && clobber(y2) && clobber(y3) && clobber(y4) && clobber(o0) && clobber(o1) && clobber(o2)) {
			break
		}
		b = mergePoint(b, x0, x1, x2, x3, x4)
		v0 := b.NewValue0(v.Pos, OpARM64REV, t)
		v.reset(OpCopy)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDloadidx, t)
		v1.AddArg(ptr)
		v1.AddArg(idx)
		v1.AddArg(mem)
		v0.AddArg(v1)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64ORshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ORconst)
		v.AuxInt = c >> uint64(d)
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARM64SRAconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64ORshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64ORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64ORconst)
		v.AuxInt = int64(uint64(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARM64SRLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 64-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARM64RORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVWUreg {
			break
		}
		if x != v_1.Args[0] {
			break
		}
		if !(c < 32 && t.Size(psess.types) == 4) {
			break
		}
		v.reset(OpARM64RORWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		rc := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ANDconst {
			break
		}
		ac := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		lc := v_1.AuxInt
		y := v_1.Args[0]
		if !(lc > rc && ac == ^((1<<uint(64-lc)-1)<<uint64(lc-rc))) {
			break
		}
		v.reset(OpARM64BFI)
		v.AuxInt = arm64BFAuxInt(lc-rc, 64-lc)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SLL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64SLLconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SLLconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = d << uint64(c)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		if v_0.AuxInt != c {
			break
		}
		x := v_0.Args[0]
		if !(0 < c && c < 64) {
			break
		}
		v.reset(OpARM64ANDconst)
		v.AuxInt = ^(1<<uint(c) - 1)
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ANDconst {
			break
		}
		ac := v_0.AuxInt
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, ac, 0)) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(sc, arm64BFWidth(ac, 0))
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVWUreg {
			break
		}
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<32-1, 0)) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(sc, 32)
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVHUreg {
			break
		}
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<16-1, 0)) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(sc, 16)
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVBUreg {
			break
		}
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<8-1, 0)) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(sc, 8)
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFIZ {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc+getARM64BFwidth(bfc)+getARM64BFlsb(bfc) < 64) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(getARM64BFlsb(bfc)+sc, getARM64BFwidth(bfc))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SRA_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64SRAconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SRAconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = d >> uint64(c)
		return true
	}

	for {
		rc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		lc := v_0.AuxInt
		x := v_0.Args[0]
		if !(lc > rc) {
			break
		}
		v.reset(OpARM64SBFIZ)
		v.AuxInt = arm64BFAuxInt(lc-rc, 64-lc)
		v.AddArg(x)
		return true
	}

	for {
		rc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		lc := v_0.AuxInt
		x := v_0.Args[0]
		if !(lc <= rc) {
			break
		}
		v.reset(OpARM64SBFX)
		v.AuxInt = arm64BFAuxInt(rc-lc, 64-rc)
		v.AddArg(x)
		return true
	}

	for {
		rc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVWreg {
			break
		}
		x := v_0.Args[0]
		if !(rc < 32) {
			break
		}
		v.reset(OpARM64SBFX)
		v.AuxInt = arm64BFAuxInt(rc, 32-rc)
		v.AddArg(x)
		return true
	}

	for {
		rc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVHreg {
			break
		}
		x := v_0.Args[0]
		if !(rc < 16) {
			break
		}
		v.reset(OpARM64SBFX)
		v.AuxInt = arm64BFAuxInt(rc, 16-rc)
		v.AddArg(x)
		return true
	}

	for {
		rc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVBreg {
			break
		}
		x := v_0.Args[0]
		if !(rc < 8) {
			break
		}
		v.reset(OpARM64SBFX)
		v.AuxInt = arm64BFAuxInt(rc, 8-rc)
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SBFIZ {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc < getARM64BFlsb(bfc)) {
			break
		}
		v.reset(OpARM64SBFIZ)
		v.AuxInt = arm64BFAuxInt(getARM64BFlsb(bfc)-sc, getARM64BFwidth(bfc))
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SBFIZ {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc >= getARM64BFlsb(bfc) && sc < getARM64BFlsb(bfc)+getARM64BFwidth(bfc)) {
			break
		}
		v.reset(OpARM64SBFX)
		v.AuxInt = arm64BFAuxInt(sc-getARM64BFlsb(bfc), getARM64BFlsb(bfc)+getARM64BFwidth(bfc)-sc)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SRL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64SRLconst)
		v.AuxInt = c & 63
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SRLconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(uint64(d) >> uint64(c))
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != c {
			break
		}
		x := v_0.Args[0]
		if !(0 < c && c < 64) {
			break
		}
		v.reset(OpARM64ANDconst)
		v.AuxInt = 1<<uint(64-c) - 1
		v.AddArg(x)
		return true
	}

	for {
		rc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		lc := v_0.AuxInt
		x := v_0.Args[0]
		if !(lc > rc) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(lc-rc, 64-lc)
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ANDconst {
			break
		}
		ac := v_0.AuxInt
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, ac, sc)) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(sc, arm64BFWidth(ac, sc))
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVWUreg {
			break
		}
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<32-1, sc)) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(sc, arm64BFWidth(1<<32-1, sc))
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVHUreg {
			break
		}
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<16-1, sc)) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(sc, arm64BFWidth(1<<16-1, sc))
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVBUreg {
			break
		}
		x := v_0.Args[0]
		if !(isARM64BFMask(sc, 1<<8-1, sc)) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(sc, arm64BFWidth(1<<8-1, sc))
		v.AddArg(x)
		return true
	}

	for {
		rc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		lc := v_0.AuxInt
		x := v_0.Args[0]
		if !(lc < rc) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(rc-lc, 64-rc)
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFX {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc < getARM64BFwidth(bfc)) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(getARM64BFlsb(bfc)+sc, getARM64BFwidth(bfc)-sc)
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFIZ {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc == getARM64BFlsb(bfc)) {
			break
		}
		v.reset(OpARM64ANDconst)
		v.AuxInt = 1<<uint(getARM64BFwidth(bfc)) - 1
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SRLconst_10(v *Value) bool {

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFIZ {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc < getARM64BFlsb(bfc)) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(getARM64BFlsb(bfc)-sc, getARM64BFwidth(bfc))
		v.AddArg(x)
		return true
	}

	for {
		sc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFIZ {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc > getARM64BFlsb(bfc) && sc < getARM64BFlsb(bfc)+getARM64BFwidth(bfc)) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(sc-getARM64BFlsb(bfc), getARM64BFlsb(bfc)+getARM64BFwidth(bfc)-sc)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64STP_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val1 := v.Args[1]
		val2 := v.Args[2]
		mem := v.Args[3]
		if !(is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64STP)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val1)
		v.AddArg(val2)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val1 := v.Args[1]
		val2 := v.Args[2]
		mem := v.Args[3]
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(OpARM64STP)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val1)
		v.AddArg(val2)
		v.AddArg(mem)
		return true
	}

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v_2 := v.Args[2]
		if v_2.Op != OpARM64MOVDconst {
			break
		}
		if v_2.AuxInt != 0 {
			break
		}
		mem := v.Args[3]
		v.reset(OpARM64MOVQstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SUB_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64SUBconst)
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
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SUB {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARM64SUB)
		v0 := b.NewValue0(v.Pos, OpARM64ADD, v.Type)
		v0.AddArg(x)
		v0.AddArg(z)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SUB {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		z := v.Args[1]
		v.reset(OpARM64SUB)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpARM64ADD, y.Type)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64SUBshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64SUBshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64SUBshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SUBconst_0(v *Value) bool {

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
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = d - c
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARM64ADDconst)
		v.AuxInt = -c - d
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64ADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARM64ADDconst)
		v.AuxInt = -c + d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SUBshiftLL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64SUBconst)
		v.AuxInt = int64(uint64(c) << uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SUBshiftRA_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64SUBconst)
		v.AuxInt = c >> uint64(d)
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64SUBshiftRL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64SUBconst)
		v.AuxInt = int64(uint64(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64TST_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64TSTconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64TSTWconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) == 0) {
			break
		}
		v.reset(OpARM64FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) < 0) {
			break
		}
		v.reset(OpARM64FlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) > 0) {
			break
		}
		v.reset(OpARM64FlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64TSTconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x&y) == 0) {
			break
		}
		v.reset(OpARM64FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x&y) < 0) {
			break
		}
		v.reset(OpARM64FlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		x := v_0.AuxInt
		if !(int64(x&y) > 0) {
			break
		}
		v.reset(OpARM64FlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64UBFIZ_0(v *Value) bool {

	for {
		bfc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc < getARM64BFwidth(bfc)) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(getARM64BFlsb(bfc)+sc, getARM64BFwidth(bfc)-sc)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64UBFX_0(v *Value) bool {

	for {
		bfc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc+getARM64BFwidth(bfc)+getARM64BFlsb(bfc) < 64) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(getARM64BFlsb(bfc)+sc, getARM64BFwidth(bfc))
		v.AddArg(x)
		return true
	}

	for {
		bfc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc == getARM64BFlsb(bfc)) {
			break
		}
		v.reset(OpARM64ANDconst)
		v.AuxInt = 1<<uint(getARM64BFwidth(bfc)) - 1
		v.AddArg(x)
		return true
	}

	for {
		bfc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc < getARM64BFlsb(bfc)) {
			break
		}
		v.reset(OpARM64UBFX)
		v.AuxInt = arm64BFAuxInt(getARM64BFlsb(bfc)-sc, getARM64BFwidth(bfc))
		v.AddArg(x)
		return true
	}

	for {
		bfc := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		sc := v_0.AuxInt
		x := v_0.Args[0]
		if !(sc > getARM64BFlsb(bfc) && sc < getARM64BFlsb(bfc)+getARM64BFwidth(bfc)) {
			break
		}
		v.reset(OpARM64UBFIZ)
		v.AuxInt = arm64BFAuxInt(sc-getARM64BFlsb(bfc), getARM64BFlsb(bfc)+getARM64BFwidth(bfc)-sc)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64UDIV_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
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
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARM64SRLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(uint64(c) / uint64(d))
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64UDIVW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) == 1) {
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
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64SRLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(uint32(c) / uint32(d))
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64UMOD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARM64ANDconst)
		v.AuxInt = c - 1
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(uint64(c) % uint64(d))
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64UMODW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(uint32(c) == 1) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c) && is32Bit(c)) {
			break
		}
		v.reset(OpARM64ANDconst)
		v.AuxInt = c - 1
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = int64(uint32(c) % uint32(d))
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64XOR_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64XORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64XORconst)
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
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MVN {
			break
		}
		y := v_1.Args[0]
		v.reset(OpARM64EON)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MVN {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARM64EON)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64XORshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SLLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64XORshiftLL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64XORshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SRLconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64XORshiftRL)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		x1 := v.Args[1]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64XORshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64XOR_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x1 := v.Args[0]
		if x1.Op != OpARM64SRAconst {
			break
		}
		c := x1.AuxInt
		y := x1.Args[0]
		x0 := v.Args[1]
		if !(clobberIfDead(x1)) {
			break
		}
		v.reset(OpARM64XORshiftRA)
		v.AuxInt = c
		v.AddArg(x0)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64XORconst_0(v *Value) bool {

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
		v.reset(OpARM64MVN)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = c ^ d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARM64XORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARM64XORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64XORshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64XORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64XORconst)
		v.AuxInt = int64(uint64(c) << uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		if v_0.AuxInt != 64-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARM64RORconst)
		v.AuxInt = 64 - c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFX {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		if !(c < 32 && t.Size(psess.types) == 4 && bfc == arm64BFAuxInt(32-c, c)) {
			break
		}
		v.reset(OpARM64RORWconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SRLconst {
			break
		}
		if v_0.AuxInt != 64-c {
			break
		}
		x := v_0.Args[0]
		x2 := v.Args[1]
		v.reset(OpARM64EXTRconst)
		v.AuxInt = 64 - c
		v.AddArg(x2)
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64UBFX {
			break
		}
		bfc := v_0.AuxInt
		x := v_0.Args[0]
		x2 := v.Args[1]
		if !(c < 32 && t.Size(psess.types) == 4 && bfc == arm64BFAuxInt(32-c, c)) {
			break
		}
		v.reset(OpARM64EXTRWconst)
		v.AuxInt = 32 - c
		v.AddArg(x2)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpARM64XORshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64XORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64XORconst)
		v.AuxInt = c >> uint64(d)
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpARM64XORshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64MOVDconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARM64XORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARM64SRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVDconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARM64XORconst)
		v.AuxInt = int64(uint64(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64SRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 64-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARM64RORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARM64SLLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARM64MOVWUreg {
			break
		}
		if x != v_1.Args[0] {
			break
		}
		if !(c < 32 && t.Size(psess.types) == 4) {
			break
		}
		v.reset(OpARM64RORWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM64_OpAdd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAdd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAdd32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FADDS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAdd64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAdd64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FADDD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAdd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAddPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64ADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAddr_0(v *Value) bool {

	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpARM64MOVDaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueARM64_OpAnd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAnd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAnd64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAnd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAndB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64AND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpAtomicAdd32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64LoweredAtomicAdd32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAtomicAdd32Variant_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64LoweredAtomicAdd32Variant)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAtomicAdd64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64LoweredAtomicAdd64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAtomicAdd64Variant_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64LoweredAtomicAdd64Variant)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpAtomicAnd8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpARM64LoweredAtomicAnd8, types.NewTuple(typ.UInt8, psess.types.TypeMem))
		v0.AddArg(ptr)
		v0.AddArg(val)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpAtomicCompareAndSwap32_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64LoweredAtomicCas32)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAtomicCompareAndSwap64_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		old := v.Args[1]
		new_ := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARM64LoweredAtomicCas64)
		v.AddArg(ptr)
		v.AddArg(old)
		v.AddArg(new_)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAtomicExchange32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64LoweredAtomicExchange32)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAtomicExchange64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64LoweredAtomicExchange64)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAtomicLoad32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64LDARW)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAtomicLoad64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64LDAR)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAtomicLoadPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64LDAR)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpAtomicOr8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpARM64LoweredAtomicOr8, types.NewTuple(typ.UInt8, psess.types.TypeMem))
		v0.AddArg(ptr)
		v0.AddArg(val)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpAtomicStore32_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64STLRW)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAtomicStore64_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64STLR)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAtomicStorePtrNoWB_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64STLR)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpAvg64u_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64ADD)
		v0 := b.NewValue0(v.Pos, OpARM64SRLconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpARM64SUB, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpBitLen64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpARM64SUB)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 64
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CLZ, typ.Int)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM64_OpBitRev16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpARM64SRLconst)
		v.AuxInt = 48
		v0 := b.NewValue0(v.Pos, OpARM64RBIT, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpBitRev32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64RBITW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpBitRev64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64RBIT)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpBitRev8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpARM64SRLconst)
		v.AuxInt = 56
		v0 := b.NewValue0(v.Pos, OpARM64RBIT, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpBswap32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64REVW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpBswap64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64REV)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCeil_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FRINTPD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpClosureCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64CALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpCom16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCom32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCom64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCom8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MVN)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpCondSelect_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		bool := v.Args[2]
		if !(psess.flagArg(bool) != nil) {
			break
		}
		v.reset(OpARM64CSEL)
		v.Aux = bool.Op
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(psess.flagArg(bool))
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		bool := v.Args[2]
		if !(psess.flagArg(bool) == nil) {
			break
		}
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64NotEqual
		v.AddArg(x)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpARM64CMPWconst, psess.types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(bool)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM64_OpConst16_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM64_OpConst32_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM64_OpConst32F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpARM64FMOVSconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM64_OpConst64_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM64_OpConst64F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpARM64FMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM64_OpConst8_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM64_OpConstBool_0(v *Value) bool {

	for {
		b := v.AuxInt
		v.reset(OpARM64MOVDconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValueARM64_OpConstNil_0(v *Value) bool {

	for {
		v.reset(OpARM64MOVDconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueARM64_OpCtz32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpARM64CLZW)
		v0 := b.NewValue0(v.Pos, OpARM64RBITW, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpCtz32NonZero_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCtz32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCtz64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpARM64CLZ)
		v0 := b.NewValue0(v.Pos, OpARM64RBIT, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpCtz64NonZero_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCtz64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt32Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FCVTZSSW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt32Fto32U_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FCVTZUSW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt32Fto64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FCVTZSS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt32Fto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FCVTSD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt32Fto64U_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FCVTZUS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt32Uto32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64UCVTFWS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt32Uto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64UCVTFWD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt32to32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64SCVTFWS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt32to64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64SCVTFWD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt64Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FCVTZSDW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt64Fto32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FCVTDS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt64Fto32U_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FCVTZUDW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt64Fto64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FCVTZSD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt64Fto64U_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FCVTZUD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt64Uto32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64UCVTFS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt64Uto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64UCVTFD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt64to32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64SCVTFS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpCvt64to64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64SCVTFD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpDiv16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64DIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM64_OpDiv16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64UDIVW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM64_OpDiv32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64DIVW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpDiv32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FDIVS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpDiv32u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64UDIVW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpDiv64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64DIV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpDiv64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FDIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpDiv64u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64UDIV)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpDiv8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64DIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM64_OpDiv8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64UDIVW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpEq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64Equal)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpEq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64Equal)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpEq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64Equal)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPS, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpEq64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64Equal)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpEq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64Equal)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpEq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64Equal)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func rewriteValueARM64_OpEqB_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64XOR)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64XOR, typ.Bool)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpEqPtr_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64Equal)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpFloor_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FRINTMD)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpGeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpGeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPS, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGeq32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGeq64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGeq64U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpGeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func rewriteValueARM64_OpGetCallerPC_0(v *Value) bool {

	for {
		v.reset(OpARM64LoweredGetCallerPC)
		return true
	}
}
func rewriteValueARM64_OpGetCallerSP_0(v *Value) bool {

	for {
		v.reset(OpARM64LoweredGetCallerSP)
		return true
	}
}
func rewriteValueARM64_OpGetClosurePtr_0(v *Value) bool {

	for {
		v.reset(OpARM64LoweredGetClosurePtr)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGreater16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpGreater16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThanU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpGreater32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGreater32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPS, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGreater32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThanU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGreater64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGreater64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGreater64U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThanU)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpGreater8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpGreater8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThanU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func rewriteValueARM64_OpHmul32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRAconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpARM64MULL, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpHmul32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRAconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpARM64UMULL, typ.UInt64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpHmul64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64MULH)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpHmul64u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64UMULH)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpInterCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[1]
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64CALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpIsInBounds_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpARM64LessThanU)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		ptr := v.Args[0]
		v.reset(OpARM64NotEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(ptr)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpIsSliceInBounds_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpARM64LessEqualU)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpLeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessEqualU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpLeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPS, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLeq32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessEqualU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLeq64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPD, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLeq64U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessEqualU)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpLeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessEqualU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpLess16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessThan)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpLess16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessThanU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpLess32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessThan)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLess32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPS, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLess32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessThanU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLess64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessThan)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLess64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64GreaterThan)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPD, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLess64U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessThanU)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLess8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessThan)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpLess8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LessThanU)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpLoad_0(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpARM64MOVBUload)
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
		v.reset(OpARM64MOVBload)
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
		v.reset(OpARM64MOVBUload)
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
		v.reset(OpARM64MOVHload)
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
		v.reset(OpARM64MOVHUload)
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
		v.reset(OpARM64MOVWload)
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
		v.reset(OpARM64MOVWUload)
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
		v.reset(OpARM64MOVDload)
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
		v.reset(OpARM64FMOVSload)
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
		v.reset(OpARM64FMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConst64, t)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConst64, t)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConst64, t)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConst64, t)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SLL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueARM64_OpMod16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64MODW)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM64_OpMod16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64UMODW)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM64_OpMod32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64MODW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpMod32u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64UMODW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpMod64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64MOD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpMod64u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64UMOD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpMod8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64MODW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM64_OpMod8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64UMODW)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpMove_0(v *Value) bool {
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
		v.reset(OpARM64MOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVBUload, typ.UInt8)
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
		v.reset(OpARM64MOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVHUload, typ.UInt16)
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
		v.reset(OpARM64MOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUload, typ.UInt32)
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
		v.reset(OpARM64MOVDstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDload, typ.UInt64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
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
		v.reset(OpARM64MOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVBUload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVHstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARM64MOVHUload, typ.UInt16)
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
		v.reset(OpARM64MOVBstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVBUload, typ.UInt8)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARM64MOVWUload, typ.UInt32)
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
		v.reset(OpARM64MOVHstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVHUload, typ.UInt16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARM64MOVWUload, typ.UInt32)
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
		v.reset(OpARM64MOVBstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVBUload, typ.UInt8)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVHstore, psess.types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARM64MOVHUload, typ.UInt16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64MOVWstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpARM64MOVWUload, typ.UInt32)
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
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVWstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVWUload, typ.UInt32)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDload, typ.UInt64)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpMove_10(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		if v.AuxInt != 16 {
			break
		}
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64MOVDstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDload, typ.UInt64)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDload, typ.UInt64)
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
		v.reset(OpARM64MOVDstore)
		v.AuxInt = 16
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDload, typ.UInt64)
		v0.AuxInt = 16
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDstore, psess.types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDload, typ.UInt64)
		v2.AuxInt = 8
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64MOVDstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpARM64MOVDload, typ.UInt64)
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
		if !(s%8 != 0 && s > 8) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = s % 8
		v0 := b.NewValue0(v.Pos, OpOffPtr, dst.Type)
		v0.AuxInt = s - s%8
		v0.AddArg(dst)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, src.Type)
		v1.AuxInt = s - s%8
		v1.AddArg(src)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpMove, psess.types.TypeMem)
		v2.AuxInt = s - s%8
		v2.AddArg(dst)
		v2.AddArg(src)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 32 && s <= 16*64 && s%16 == 8 && !config.noDuffDevice) {
			break
		}
		v.reset(OpARM64MOVDstore)
		v.AuxInt = s - 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDload, typ.UInt64)
		v0.AuxInt = s - 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64DUFFCOPY, psess.types.TypeMem)
		v1.AuxInt = 8 * (64 - (s-8)/16)
		v1.AddArg(dst)
		v1.AddArg(src)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 32 && s <= 16*64 && s%16 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpARM64DUFFCOPY)
		v.AuxInt = 8 * (64 - s/16)
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
		if !(s > 24 && s%8 == 0) {
			break
		}
		v.reset(OpARM64LoweredMove)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpARM64ADDconst, src.Type)
		v0.AuxInt = s - 8
		v0.AddArg(src)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpMul16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64MULW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpMul32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64MULW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpMul32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FMULS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpMul64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64MUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpMul64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpMul64uhilo_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64LoweredMuluhilo)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpMul8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64MULW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpNeg16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpNeg32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpNeg32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FNEGS)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpNeg64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpNeg64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FNEGD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpNeg8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64NEG)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpNeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64NotEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM64_OpNeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64NotEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpNeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64NotEqual)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPS, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpNeq64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64NotEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpNeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64NotEqual)
		v0 := b.NewValue0(v.Pos, OpARM64FCMPD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpNeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64NotEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMPW, psess.types.TypeFlags)
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
func rewriteValueARM64_OpNeqB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpNeqPtr_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64NotEqual)
		v0 := b.NewValue0(v.Pos, OpARM64CMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpNilCheck_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64LoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpNot_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpARM64XOR)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpOffPtr_0(v *Value) bool {

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpARM64MOVDaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpARM64ADDconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueARM64_OpOr16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpOr32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpOr64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpOr8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpOrB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64OR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpPopCount16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpARM64FMOVDfpgp)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpARM64VUADDLV, typ.Float64)
		v1 := b.NewValue0(v.Pos, OpARM64VCNT, typ.Float64)
		v2 := b.NewValue0(v.Pos, OpARM64FMOVDgpfp, typ.Float64)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpPopCount32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpARM64FMOVDfpgp)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpARM64VUADDLV, typ.Float64)
		v1 := b.NewValue0(v.Pos, OpARM64VCNT, typ.Float64)
		v2 := b.NewValue0(v.Pos, OpARM64FMOVDgpfp, typ.Float64)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(x)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpPopCount64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpARM64FMOVDfpgp)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpARM64VUADDLV, typ.Float64)
		v1 := b.NewValue0(v.Pos, OpARM64VCNT, typ.Float64)
		v2 := b.NewValue0(v.Pos, OpARM64FMOVDgpfp, typ.Float64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpRound_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FRINTAD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpRound32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64LoweredRound32F)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpRound64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64LoweredRound64F)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh16Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v3.AuxInt = 63
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v3.AuxInt = 63
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v3.AuxInt = 63
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh32Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v3.AuxInt = 63
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v3.AuxInt = 63
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v3.AuxInt = 63
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh64Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh64Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh64Ux64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpConst64, t)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh64Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v0.Aux = OpARM64LessThanU
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v2.AuxInt = 63
		v0.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v0.Aux = OpARM64LessThanU
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v2.AuxInt = 63
		v0.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v0.Aux = OpARM64LessThanU
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v2.AuxInt = 64
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v0.Aux = OpARM64LessThanU
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v2.AuxInt = 63
		v0.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v0.AddArg(v3)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh8Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = 0
		v.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64CSEL)
		v.Aux = OpARM64LessThanU
		v0 := b.NewValue0(v.Pos, OpARM64SRL, t)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = 0
		v.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v3.AuxInt = 63
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v3.AuxInt = 63
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v1.AddArg(y)
		v2 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v2.AuxInt = 63
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v3.AuxInt = 64
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64CSEL, y.Type)
		v1.Aux = OpARM64LessThanU
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpConst64, y.Type)
		v3.AuxInt = 63
		v1.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64CMPconst, psess.types.TypeFlags)
		v4.AuxInt = 64
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v1.AddArg(v4)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM64_OpSignExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpSignExt16to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpSignExt32to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVWreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpSignExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpSignExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpSignExt8to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpSlicemask_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpARM64SRAconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpARM64NEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM64_OpSqrt_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FSQRTD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpStaticCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpARM64CALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpStore_0(v *Value) bool {

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 1) {
			break
		}
		v.reset(OpARM64MOVBstore)
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
		v.reset(OpARM64MOVHstore)
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
		v.reset(OpARM64MOVWstore)
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
		v.reset(OpARM64MOVDstore)
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
		v.reset(OpARM64FMOVSstore)
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
		v.reset(OpARM64FMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpSub16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpSub32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpSub32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FSUBS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpSub64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpSub64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64FSUBD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpSub8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpSubPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64SUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpTrunc_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64FRINTZD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpTrunc16to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpTrunc32to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpTrunc32to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpTrunc64to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpTrunc64to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpTrunc64to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpWB_0(v *Value) bool {

	for {
		fn := v.Aux
		_ = v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARM64LoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM64_OpXor16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpXor32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpXor64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM64_OpXor8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARM64XOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM64_OpZero_0(v *Value) bool {
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
		v.reset(OpARM64MOVBstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
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
		v.reset(OpARM64MOVHstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 4 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVWstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVDstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 3 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVBstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVHstore, psess.types.TypeMem)
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 5 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVBstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWstore, psess.types.TypeMem)
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 6 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVHstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWstore, psess.types.TypeMem)
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 7 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVBstore)
		v.AuxInt = 6
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVHstore, psess.types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64MOVWstore, psess.types.TypeMem)
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 9 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVBstore)
		v.AuxInt = 8
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDstore, psess.types.TypeMem)
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpZero_10(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		if v.AuxInt != 10 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVHstore)
		v.AuxInt = 8
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDstore, psess.types.TypeMem)
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 11 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVBstore)
		v.AuxInt = 10
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVHstore, psess.types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64MOVDstore, psess.types.TypeMem)
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
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
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVWstore)
		v.AuxInt = 8
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDstore, psess.types.TypeMem)
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 13 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVBstore)
		v.AuxInt = 12
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWstore, psess.types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64MOVDstore, psess.types.TypeMem)
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 14 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVHstore)
		v.AuxInt = 12
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVWstore, psess.types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64MOVDstore, psess.types.TypeMem)
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 15 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64MOVBstore)
		v.AuxInt = 14
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVHstore, psess.types.TypeMem)
		v1.AuxInt = 12
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARM64MOVWstore, psess.types.TypeMem)
		v3.AuxInt = 8
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpARM64MOVDstore, psess.types.TypeMem)
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 16 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64STP)
		v.AuxInt = 0
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 32 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64STP)
		v.AuxInt = 16
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARM64STP, psess.types.TypeMem)
		v2.AuxInt = 0
		v2.AddArg(ptr)
		v3 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v3.AuxInt = 0
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v4.AuxInt = 0
		v2.AddArg(v4)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}

	for {
		if v.AuxInt != 48 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64STP)
		v.AuxInt = 32
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARM64STP, psess.types.TypeMem)
		v2.AuxInt = 16
		v2.AddArg(ptr)
		v3 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v3.AuxInt = 0
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v4.AuxInt = 0
		v2.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpARM64STP, psess.types.TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v7 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v7.AuxInt = 0
		v5.AddArg(v7)
		v5.AddArg(mem)
		v2.AddArg(v5)
		v.AddArg(v2)
		return true
	}

	for {
		if v.AuxInt != 64 {
			break
		}
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARM64STP)
		v.AuxInt = 48
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARM64STP, psess.types.TypeMem)
		v2.AuxInt = 32
		v2.AddArg(ptr)
		v3 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v3.AuxInt = 0
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v4.AuxInt = 0
		v2.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpARM64STP, psess.types.TypeMem)
		v5.AuxInt = 16
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v7 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v7.AuxInt = 0
		v5.AddArg(v7)
		v8 := b.NewValue0(v.Pos, OpARM64STP, psess.types.TypeMem)
		v8.AuxInt = 0
		v8.AddArg(ptr)
		v9 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v9.AuxInt = 0
		v8.AddArg(v9)
		v10 := b.NewValue0(v.Pos, OpARM64MOVDconst, typ.UInt64)
		v10.AuxInt = 0
		v8.AddArg(v10)
		v8.AddArg(mem)
		v5.AddArg(v8)
		v2.AddArg(v5)
		v.AddArg(v2)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM64_OpZero_20(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		s := v.AuxInt
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(s%16 != 0 && s%16 <= 8 && s > 16) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = 8
		v0 := b.NewValue0(v.Pos, OpOffPtr, ptr.Type)
		v0.AuxInt = s - 8
		v0.AddArg(ptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZero, psess.types.TypeMem)
		v1.AuxInt = s - s%16
		v1.AddArg(ptr)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(s%16 != 0 && s%16 > 8 && s > 16) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = 16
		v0 := b.NewValue0(v.Pos, OpOffPtr, ptr.Type)
		v0.AuxInt = s - 16
		v0.AddArg(ptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZero, psess.types.TypeMem)
		v1.AuxInt = s - s%16
		v1.AddArg(ptr)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(s%16 == 0 && s > 64 && s <= 16*64 && !config.noDuffDevice) {
			break
		}
		v.reset(OpARM64DUFFZERO)
		v.AuxInt = 4 * (64 - s/16)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(s%16 == 0 && (s > 16*64 || config.noDuffDevice)) {
			break
		}
		v.reset(OpARM64LoweredZero)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARM64ADDconst, ptr.Type)
		v0.AuxInt = s - 16
		v0.AddArg(ptr)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM64_OpZeroExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVHUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpZeroExt16to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVHUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpZeroExt32to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVWUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpZeroExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpZeroExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM64_OpZeroExt8to64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARM64MOVBUreg)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteBlockARM64(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	case BlockARM64EQ:

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64EQ
			v0 := b.NewValue0(v.Pos, OpARM64TSTWconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			z := v.Args[0]
			if z.Op != OpARM64AND {
				break
			}
			_ = z.Args[1]
			x := z.Args[0]
			y := z.Args[1]
			if !(z.Uses == 1) {
				break
			}
			b.Kind = BlockARM64EQ
			v0 := b.NewValue0(v.Pos, OpARM64TST, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			z := v.Args[0]
			if z.Op != OpARM64AND {
				break
			}
			_ = z.Args[1]
			x := z.Args[0]
			y := z.Args[1]
			if !(z.Uses == 1) {
				break
			}
			b.Kind = BlockARM64EQ
			v0 := b.NewValue0(v.Pos, OpARM64TST, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64EQ
			v0 := b.NewValue0(v.Pos, OpARM64TSTconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			z := v.Args[0]
			if z.Op != OpARM64ADD {
				break
			}
			_ = z.Args[1]
			x := z.Args[0]
			y := z.Args[1]
			if !(z.Uses == 1) {
				break
			}
			b.Kind = BlockARM64EQ
			v0 := b.NewValue0(v.Pos, OpARM64CMN, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMP {
				break
			}
			_ = v.Args[1]
			x := v.Args[0]
			z := v.Args[1]
			if z.Op != OpARM64NEG {
				break
			}
			y := z.Args[0]
			if !(z.Uses == 1) {
				break
			}
			b.Kind = BlockARM64EQ
			v0 := b.NewValue0(v.Pos, OpARM64CMN, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockARM64Z
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockARM64ZW
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64TSTconst {
				break
			}
			c := v.AuxInt
			x := v.Args[0]
			if !(oneBit(c)) {
				break
			}
			b.Kind = BlockARM64TBZ
			b.SetControl(x)
			b.Aux = ntz(c)
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64TSTWconst {
				break
			}
			c := v.AuxInt
			x := v.Args[0]
			if !(oneBit(int64(uint32(c)))) {
				break
			}
			b.Kind = BlockARM64TBZ
			b.SetControl(x)
			b.Aux = ntz(int64(uint32(c)))
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagLT_ULT {
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
			if v.Op != OpARM64FlagLT_UGT {
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
			if v.Op != OpARM64FlagGT_ULT {
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
			if v.Op != OpARM64FlagGT_UGT {
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
			if v.Op != OpARM64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARM64EQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARM64GE:

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64GE
			v0 := b.NewValue0(v.Pos, OpARM64TSTWconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64GE
			v0 := b.NewValue0(v.Pos, OpARM64TSTconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockARM64TBZ
			b.SetControl(x)
			b.Aux = int64(31)
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockARM64TBZ
			b.SetControl(x)
			b.Aux = int64(63)
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagLT_ULT {
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
			if v.Op != OpARM64FlagLT_UGT {
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
			if v.Op != OpARM64FlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARM64LE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARM64GT:

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64GT
			v0 := b.NewValue0(v.Pos, OpARM64TSTWconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64GT
			v0 := b.NewValue0(v.Pos, OpARM64TSTconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagEQ {
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
			if v.Op != OpARM64FlagLT_ULT {
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
			if v.Op != OpARM64FlagLT_UGT {
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
			if v.Op != OpARM64FlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARM64LT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockIf:

		for {
			v := b.Control
			if v.Op != OpARM64Equal {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64EQ
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64NotEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64NE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64LessThan {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64LT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64LessThanU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64ULT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64LessEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64LE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64LessEqualU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64ULE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64GreaterThan {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64GT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64GreaterThanU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64UGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64GreaterEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64GE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64GreaterEqualU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64UGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			_ = v
			cond := b.Control
			b.Kind = BlockARM64NZ
			b.SetControl(cond)
			b.Aux = nil
			return true
		}
	case BlockARM64LE:

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64LE
			v0 := b.NewValue0(v.Pos, OpARM64TSTWconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64LE
			v0 := b.NewValue0(v.Pos, OpARM64TSTconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagGT_ULT {
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
			if v.Op != OpARM64FlagGT_UGT {
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
			if v.Op != OpARM64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARM64GE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARM64LT:

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64LT
			v0 := b.NewValue0(v.Pos, OpARM64TSTWconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64LT
			v0 := b.NewValue0(v.Pos, OpARM64TSTconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockARM64TBNZ
			b.SetControl(x)
			b.Aux = int64(31)
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockARM64TBNZ
			b.SetControl(x)
			b.Aux = int64(63)
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagEQ {
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
			if v.Op != OpARM64FlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagGT_ULT {
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
			if v.Op != OpARM64FlagGT_UGT {
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
			if v.Op != OpARM64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARM64GT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARM64NE:

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64NE
			v0 := b.NewValue0(v.Pos, OpARM64TSTWconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			z := v.Args[0]
			if z.Op != OpARM64AND {
				break
			}
			_ = z.Args[1]
			x := z.Args[0]
			y := z.Args[1]
			if !(z.Uses == 1) {
				break
			}
			b.Kind = BlockARM64NE
			v0 := b.NewValue0(v.Pos, OpARM64TST, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			z := v.Args[0]
			if z.Op != OpARM64AND {
				break
			}
			_ = z.Args[1]
			x := z.Args[0]
			y := z.Args[1]
			if !(z.Uses == 1) {
				break
			}
			b.Kind = BlockARM64NE
			v0 := b.NewValue0(v.Pos, OpARM64TST, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			if x.Op != OpARM64ANDconst {
				break
			}
			c := x.AuxInt
			y := x.Args[0]
			if !(x.Uses == 1) {
				break
			}
			b.Kind = BlockARM64NE
			v0 := b.NewValue0(v.Pos, OpARM64TSTconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			z := v.Args[0]
			if z.Op != OpARM64ADD {
				break
			}
			_ = z.Args[1]
			x := z.Args[0]
			y := z.Args[1]
			if !(z.Uses == 1) {
				break
			}
			b.Kind = BlockARM64NE
			v0 := b.NewValue0(v.Pos, OpARM64CMN, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMP {
				break
			}
			_ = v.Args[1]
			x := v.Args[0]
			z := v.Args[1]
			if z.Op != OpARM64NEG {
				break
			}
			y := z.Args[0]
			if !(z.Uses == 1) {
				break
			}
			b.Kind = BlockARM64NE
			v0 := b.NewValue0(v.Pos, OpARM64CMN, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockARM64NZ
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64CMPWconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			x := v.Args[0]
			b.Kind = BlockARM64NZW
			b.SetControl(x)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64TSTconst {
				break
			}
			c := v.AuxInt
			x := v.Args[0]
			if !(oneBit(c)) {
				break
			}
			b.Kind = BlockARM64TBNZ
			b.SetControl(x)
			b.Aux = ntz(c)
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64TSTWconst {
				break
			}
			c := v.AuxInt
			x := v.Args[0]
			if !(oneBit(int64(uint32(c)))) {
				break
			}
			b.Kind = BlockARM64TBNZ
			b.SetControl(x)
			b.Aux = ntz(int64(uint32(c)))
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagEQ {
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
			if v.Op != OpARM64FlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARM64NE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARM64NZ:

		for {
			v := b.Control
			if v.Op != OpARM64Equal {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64EQ
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64NotEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64NE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64LessThan {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64LT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64LessThanU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64ULT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64LessEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64LE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64LessEqualU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64ULE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64GreaterThan {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64GT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64GreaterThanU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64UGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64GreaterEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64GE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64GreaterEqualU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARM64UGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64ANDconst {
				break
			}
			c := v.AuxInt
			x := v.Args[0]
			if !(oneBit(c)) {
				break
			}
			b.Kind = BlockARM64TBNZ
			b.SetControl(x)
			b.Aux = ntz(c)
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64MOVDconst {
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
			if v.Op != OpARM64MOVDconst {
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
	case BlockARM64NZW:

		for {
			v := b.Control
			if v.Op != OpARM64ANDconst {
				break
			}
			c := v.AuxInt
			x := v.Args[0]
			if !(oneBit(int64(uint32(c)))) {
				break
			}
			b.Kind = BlockARM64TBNZ
			b.SetControl(x)
			b.Aux = ntz(int64(uint32(c)))
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64MOVDconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) == 0) {
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
			if v.Op != OpARM64MOVDconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) != 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
	case BlockARM64UGE:

		for {
			v := b.Control
			if v.Op != OpARM64FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagLT_ULT {
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
			if v.Op != OpARM64FlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagGT_ULT {
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
			if v.Op != OpARM64FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARM64ULE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARM64UGT:

		for {
			v := b.Control
			if v.Op != OpARM64FlagEQ {
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
			if v.Op != OpARM64FlagLT_ULT {
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
			if v.Op != OpARM64FlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagGT_ULT {
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
			if v.Op != OpARM64FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARM64ULT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARM64ULE:

		for {
			v := b.Control
			if v.Op != OpARM64FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagLT_UGT {
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
			if v.Op != OpARM64FlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagGT_UGT {
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
			if v.Op != OpARM64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARM64UGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARM64ULT:

		for {
			v := b.Control
			if v.Op != OpARM64FlagEQ {
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
			if v.Op != OpARM64FlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagLT_UGT {
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
			if v.Op != OpARM64FlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64FlagGT_UGT {
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
			if v.Op != OpARM64InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARM64UGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARM64Z:

		for {
			v := b.Control
			if v.Op != OpARM64ANDconst {
				break
			}
			c := v.AuxInt
			x := v.Args[0]
			if !(oneBit(c)) {
				break
			}
			b.Kind = BlockARM64TBZ
			b.SetControl(x)
			b.Aux = ntz(c)
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64MOVDconst {
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
			if v.Op != OpARM64MOVDconst {
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
	case BlockARM64ZW:

		for {
			v := b.Control
			if v.Op != OpARM64ANDconst {
				break
			}
			c := v.AuxInt
			x := v.Args[0]
			if !(oneBit(int64(uint32(c)))) {
				break
			}
			b.Kind = BlockARM64TBZ
			b.SetControl(x)
			b.Aux = ntz(int64(uint32(c)))
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64MOVDconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) == 0) {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARM64MOVDconst {
				break
			}
			c := v.AuxInt
			if !(int32(c) != 0) {
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

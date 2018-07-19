package ssa

import "math"

import "github.com/dave/golib/src/cmd/compile/internal/types"

// in case not otherwise used
// in case not otherwise used
// in case not otherwise used
// in case not otherwise used

func (psess *PackageSession) rewriteValue386(v *Value) bool {
	switch v.Op {
	case Op386ADCL:
		return rewriteValue386_Op386ADCL_0(v)
	case Op386ADDL:
		return psess.rewriteValue386_Op386ADDL_0(v) || rewriteValue386_Op386ADDL_10(v) || psess.rewriteValue386_Op386ADDL_20(v)
	case Op386ADDLcarry:
		return rewriteValue386_Op386ADDLcarry_0(v)
	case Op386ADDLconst:
		return rewriteValue386_Op386ADDLconst_0(v)
	case Op386ADDLload:
		return rewriteValue386_Op386ADDLload_0(v)
	case Op386ADDLmodify:
		return rewriteValue386_Op386ADDLmodify_0(v)
	case Op386ADDSD:
		return psess.rewriteValue386_Op386ADDSD_0(v)
	case Op386ADDSDload:
		return rewriteValue386_Op386ADDSDload_0(v)
	case Op386ADDSS:
		return psess.rewriteValue386_Op386ADDSS_0(v)
	case Op386ADDSSload:
		return rewriteValue386_Op386ADDSSload_0(v)
	case Op386ANDL:
		return psess.rewriteValue386_Op386ANDL_0(v)
	case Op386ANDLconst:
		return rewriteValue386_Op386ANDLconst_0(v)
	case Op386ANDLload:
		return rewriteValue386_Op386ANDLload_0(v)
	case Op386ANDLmodify:
		return rewriteValue386_Op386ANDLmodify_0(v)
	case Op386CMPB:
		return psess.rewriteValue386_Op386CMPB_0(v)
	case Op386CMPBconst:
		return rewriteValue386_Op386CMPBconst_0(v)
	case Op386CMPL:
		return psess.rewriteValue386_Op386CMPL_0(v)
	case Op386CMPLconst:
		return rewriteValue386_Op386CMPLconst_0(v)
	case Op386CMPW:
		return psess.rewriteValue386_Op386CMPW_0(v)
	case Op386CMPWconst:
		return rewriteValue386_Op386CMPWconst_0(v)
	case Op386LEAL:
		return rewriteValue386_Op386LEAL_0(v)
	case Op386LEAL1:
		return rewriteValue386_Op386LEAL1_0(v)
	case Op386LEAL2:
		return rewriteValue386_Op386LEAL2_0(v)
	case Op386LEAL4:
		return rewriteValue386_Op386LEAL4_0(v)
	case Op386LEAL8:
		return rewriteValue386_Op386LEAL8_0(v)
	case Op386MOVBLSX:
		return rewriteValue386_Op386MOVBLSX_0(v)
	case Op386MOVBLSXload:
		return rewriteValue386_Op386MOVBLSXload_0(v)
	case Op386MOVBLZX:
		return rewriteValue386_Op386MOVBLZX_0(v)
	case Op386MOVBload:
		return rewriteValue386_Op386MOVBload_0(v)
	case Op386MOVBloadidx1:
		return rewriteValue386_Op386MOVBloadidx1_0(v)
	case Op386MOVBstore:
		return rewriteValue386_Op386MOVBstore_0(v)
	case Op386MOVBstoreconst:
		return rewriteValue386_Op386MOVBstoreconst_0(v)
	case Op386MOVBstoreconstidx1:
		return rewriteValue386_Op386MOVBstoreconstidx1_0(v)
	case Op386MOVBstoreidx1:
		return rewriteValue386_Op386MOVBstoreidx1_0(v) || rewriteValue386_Op386MOVBstoreidx1_10(v)
	case Op386MOVLload:
		return rewriteValue386_Op386MOVLload_0(v)
	case Op386MOVLloadidx1:
		return rewriteValue386_Op386MOVLloadidx1_0(v)
	case Op386MOVLloadidx4:
		return rewriteValue386_Op386MOVLloadidx4_0(v)
	case Op386MOVLstore:
		return rewriteValue386_Op386MOVLstore_0(v) || rewriteValue386_Op386MOVLstore_10(v)
	case Op386MOVLstoreconst:
		return rewriteValue386_Op386MOVLstoreconst_0(v)
	case Op386MOVLstoreconstidx1:
		return rewriteValue386_Op386MOVLstoreconstidx1_0(v)
	case Op386MOVLstoreconstidx4:
		return rewriteValue386_Op386MOVLstoreconstidx4_0(v)
	case Op386MOVLstoreidx1:
		return rewriteValue386_Op386MOVLstoreidx1_0(v)
	case Op386MOVLstoreidx4:
		return rewriteValue386_Op386MOVLstoreidx4_0(v)
	case Op386MOVSDconst:
		return rewriteValue386_Op386MOVSDconst_0(v)
	case Op386MOVSDload:
		return rewriteValue386_Op386MOVSDload_0(v)
	case Op386MOVSDloadidx1:
		return rewriteValue386_Op386MOVSDloadidx1_0(v)
	case Op386MOVSDloadidx8:
		return rewriteValue386_Op386MOVSDloadidx8_0(v)
	case Op386MOVSDstore:
		return rewriteValue386_Op386MOVSDstore_0(v)
	case Op386MOVSDstoreidx1:
		return rewriteValue386_Op386MOVSDstoreidx1_0(v)
	case Op386MOVSDstoreidx8:
		return rewriteValue386_Op386MOVSDstoreidx8_0(v)
	case Op386MOVSSconst:
		return rewriteValue386_Op386MOVSSconst_0(v)
	case Op386MOVSSload:
		return rewriteValue386_Op386MOVSSload_0(v)
	case Op386MOVSSloadidx1:
		return rewriteValue386_Op386MOVSSloadidx1_0(v)
	case Op386MOVSSloadidx4:
		return rewriteValue386_Op386MOVSSloadidx4_0(v)
	case Op386MOVSSstore:
		return rewriteValue386_Op386MOVSSstore_0(v)
	case Op386MOVSSstoreidx1:
		return rewriteValue386_Op386MOVSSstoreidx1_0(v)
	case Op386MOVSSstoreidx4:
		return rewriteValue386_Op386MOVSSstoreidx4_0(v)
	case Op386MOVWLSX:
		return rewriteValue386_Op386MOVWLSX_0(v)
	case Op386MOVWLSXload:
		return rewriteValue386_Op386MOVWLSXload_0(v)
	case Op386MOVWLZX:
		return rewriteValue386_Op386MOVWLZX_0(v)
	case Op386MOVWload:
		return rewriteValue386_Op386MOVWload_0(v)
	case Op386MOVWloadidx1:
		return rewriteValue386_Op386MOVWloadidx1_0(v)
	case Op386MOVWloadidx2:
		return rewriteValue386_Op386MOVWloadidx2_0(v)
	case Op386MOVWstore:
		return rewriteValue386_Op386MOVWstore_0(v)
	case Op386MOVWstoreconst:
		return rewriteValue386_Op386MOVWstoreconst_0(v)
	case Op386MOVWstoreconstidx1:
		return rewriteValue386_Op386MOVWstoreconstidx1_0(v)
	case Op386MOVWstoreconstidx2:
		return rewriteValue386_Op386MOVWstoreconstidx2_0(v)
	case Op386MOVWstoreidx1:
		return rewriteValue386_Op386MOVWstoreidx1_0(v) || rewriteValue386_Op386MOVWstoreidx1_10(v)
	case Op386MOVWstoreidx2:
		return rewriteValue386_Op386MOVWstoreidx2_0(v)
	case Op386MULL:
		return rewriteValue386_Op386MULL_0(v)
	case Op386MULLconst:
		return rewriteValue386_Op386MULLconst_0(v) || rewriteValue386_Op386MULLconst_10(v) || rewriteValue386_Op386MULLconst_20(v) || rewriteValue386_Op386MULLconst_30(v)
	case Op386MULSD:
		return psess.rewriteValue386_Op386MULSD_0(v)
	case Op386MULSDload:
		return rewriteValue386_Op386MULSDload_0(v)
	case Op386MULSS:
		return psess.rewriteValue386_Op386MULSS_0(v)
	case Op386MULSSload:
		return rewriteValue386_Op386MULSSload_0(v)
	case Op386NEGL:
		return rewriteValue386_Op386NEGL_0(v)
	case Op386NOTL:
		return rewriteValue386_Op386NOTL_0(v)
	case Op386ORL:
		return psess.rewriteValue386_Op386ORL_0(v) || rewriteValue386_Op386ORL_10(v) || rewriteValue386_Op386ORL_20(v) || rewriteValue386_Op386ORL_30(v) || rewriteValue386_Op386ORL_40(v) || rewriteValue386_Op386ORL_50(v)
	case Op386ORLconst:
		return rewriteValue386_Op386ORLconst_0(v)
	case Op386ORLload:
		return rewriteValue386_Op386ORLload_0(v)
	case Op386ORLmodify:
		return rewriteValue386_Op386ORLmodify_0(v)
	case Op386ROLBconst:
		return rewriteValue386_Op386ROLBconst_0(v)
	case Op386ROLLconst:
		return rewriteValue386_Op386ROLLconst_0(v)
	case Op386ROLWconst:
		return rewriteValue386_Op386ROLWconst_0(v)
	case Op386SARB:
		return rewriteValue386_Op386SARB_0(v)
	case Op386SARBconst:
		return rewriteValue386_Op386SARBconst_0(v)
	case Op386SARL:
		return rewriteValue386_Op386SARL_0(v)
	case Op386SARLconst:
		return rewriteValue386_Op386SARLconst_0(v)
	case Op386SARW:
		return rewriteValue386_Op386SARW_0(v)
	case Op386SARWconst:
		return rewriteValue386_Op386SARWconst_0(v)
	case Op386SBBL:
		return rewriteValue386_Op386SBBL_0(v)
	case Op386SBBLcarrymask:
		return rewriteValue386_Op386SBBLcarrymask_0(v)
	case Op386SETA:
		return rewriteValue386_Op386SETA_0(v)
	case Op386SETAE:
		return rewriteValue386_Op386SETAE_0(v)
	case Op386SETB:
		return rewriteValue386_Op386SETB_0(v)
	case Op386SETBE:
		return rewriteValue386_Op386SETBE_0(v)
	case Op386SETEQ:
		return rewriteValue386_Op386SETEQ_0(v)
	case Op386SETG:
		return rewriteValue386_Op386SETG_0(v)
	case Op386SETGE:
		return rewriteValue386_Op386SETGE_0(v)
	case Op386SETL:
		return rewriteValue386_Op386SETL_0(v)
	case Op386SETLE:
		return rewriteValue386_Op386SETLE_0(v)
	case Op386SETNE:
		return rewriteValue386_Op386SETNE_0(v)
	case Op386SHLL:
		return rewriteValue386_Op386SHLL_0(v)
	case Op386SHLLconst:
		return rewriteValue386_Op386SHLLconst_0(v)
	case Op386SHRB:
		return rewriteValue386_Op386SHRB_0(v)
	case Op386SHRBconst:
		return rewriteValue386_Op386SHRBconst_0(v)
	case Op386SHRL:
		return rewriteValue386_Op386SHRL_0(v)
	case Op386SHRLconst:
		return rewriteValue386_Op386SHRLconst_0(v)
	case Op386SHRW:
		return rewriteValue386_Op386SHRW_0(v)
	case Op386SHRWconst:
		return rewriteValue386_Op386SHRWconst_0(v)
	case Op386SUBL:
		return psess.rewriteValue386_Op386SUBL_0(v)
	case Op386SUBLcarry:
		return rewriteValue386_Op386SUBLcarry_0(v)
	case Op386SUBLconst:
		return rewriteValue386_Op386SUBLconst_0(v)
	case Op386SUBLload:
		return rewriteValue386_Op386SUBLload_0(v)
	case Op386SUBLmodify:
		return rewriteValue386_Op386SUBLmodify_0(v)
	case Op386SUBSD:
		return psess.rewriteValue386_Op386SUBSD_0(v)
	case Op386SUBSDload:
		return rewriteValue386_Op386SUBSDload_0(v)
	case Op386SUBSS:
		return psess.rewriteValue386_Op386SUBSS_0(v)
	case Op386SUBSSload:
		return rewriteValue386_Op386SUBSSload_0(v)
	case Op386XORL:
		return psess.rewriteValue386_Op386XORL_0(v) || rewriteValue386_Op386XORL_10(v)
	case Op386XORLconst:
		return rewriteValue386_Op386XORLconst_0(v)
	case Op386XORLload:
		return rewriteValue386_Op386XORLload_0(v)
	case Op386XORLmodify:
		return rewriteValue386_Op386XORLmodify_0(v)
	case OpAdd16:
		return rewriteValue386_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValue386_OpAdd32_0(v)
	case OpAdd32F:
		return rewriteValue386_OpAdd32F_0(v)
	case OpAdd32carry:
		return rewriteValue386_OpAdd32carry_0(v)
	case OpAdd32withcarry:
		return rewriteValue386_OpAdd32withcarry_0(v)
	case OpAdd64F:
		return rewriteValue386_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValue386_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValue386_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValue386_OpAddr_0(v)
	case OpAnd16:
		return rewriteValue386_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValue386_OpAnd32_0(v)
	case OpAnd8:
		return rewriteValue386_OpAnd8_0(v)
	case OpAndB:
		return rewriteValue386_OpAndB_0(v)
	case OpAvg32u:
		return rewriteValue386_OpAvg32u_0(v)
	case OpBswap32:
		return rewriteValue386_OpBswap32_0(v)
	case OpClosureCall:
		return rewriteValue386_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValue386_OpCom16_0(v)
	case OpCom32:
		return rewriteValue386_OpCom32_0(v)
	case OpCom8:
		return rewriteValue386_OpCom8_0(v)
	case OpConst16:
		return rewriteValue386_OpConst16_0(v)
	case OpConst32:
		return rewriteValue386_OpConst32_0(v)
	case OpConst32F:
		return rewriteValue386_OpConst32F_0(v)
	case OpConst64F:
		return rewriteValue386_OpConst64F_0(v)
	case OpConst8:
		return rewriteValue386_OpConst8_0(v)
	case OpConstBool:
		return rewriteValue386_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValue386_OpConstNil_0(v)
	case OpCvt32Fto32:
		return rewriteValue386_OpCvt32Fto32_0(v)
	case OpCvt32Fto64F:
		return rewriteValue386_OpCvt32Fto64F_0(v)
	case OpCvt32to32F:
		return rewriteValue386_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValue386_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValue386_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValue386_OpCvt64Fto32F_0(v)
	case OpDiv16:
		return rewriteValue386_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValue386_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValue386_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValue386_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValue386_OpDiv32u_0(v)
	case OpDiv64F:
		return rewriteValue386_OpDiv64F_0(v)
	case OpDiv8:
		return rewriteValue386_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValue386_OpDiv8u_0(v)
	case OpEq16:
		return psess.rewriteValue386_OpEq16_0(v)
	case OpEq32:
		return psess.rewriteValue386_OpEq32_0(v)
	case OpEq32F:
		return psess.rewriteValue386_OpEq32F_0(v)
	case OpEq64F:
		return psess.rewriteValue386_OpEq64F_0(v)
	case OpEq8:
		return psess.rewriteValue386_OpEq8_0(v)
	case OpEqB:
		return psess.rewriteValue386_OpEqB_0(v)
	case OpEqPtr:
		return psess.rewriteValue386_OpEqPtr_0(v)
	case OpGeq16:
		return psess.rewriteValue386_OpGeq16_0(v)
	case OpGeq16U:
		return psess.rewriteValue386_OpGeq16U_0(v)
	case OpGeq32:
		return psess.rewriteValue386_OpGeq32_0(v)
	case OpGeq32F:
		return psess.rewriteValue386_OpGeq32F_0(v)
	case OpGeq32U:
		return psess.rewriteValue386_OpGeq32U_0(v)
	case OpGeq64F:
		return psess.rewriteValue386_OpGeq64F_0(v)
	case OpGeq8:
		return psess.rewriteValue386_OpGeq8_0(v)
	case OpGeq8U:
		return psess.rewriteValue386_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValue386_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValue386_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValue386_OpGetClosurePtr_0(v)
	case OpGetG:
		return rewriteValue386_OpGetG_0(v)
	case OpGreater16:
		return psess.rewriteValue386_OpGreater16_0(v)
	case OpGreater16U:
		return psess.rewriteValue386_OpGreater16U_0(v)
	case OpGreater32:
		return psess.rewriteValue386_OpGreater32_0(v)
	case OpGreater32F:
		return psess.rewriteValue386_OpGreater32F_0(v)
	case OpGreater32U:
		return psess.rewriteValue386_OpGreater32U_0(v)
	case OpGreater64F:
		return psess.rewriteValue386_OpGreater64F_0(v)
	case OpGreater8:
		return psess.rewriteValue386_OpGreater8_0(v)
	case OpGreater8U:
		return psess.rewriteValue386_OpGreater8U_0(v)
	case OpHmul32:
		return rewriteValue386_OpHmul32_0(v)
	case OpHmul32u:
		return rewriteValue386_OpHmul32u_0(v)
	case OpInterCall:
		return rewriteValue386_OpInterCall_0(v)
	case OpIsInBounds:
		return psess.rewriteValue386_OpIsInBounds_0(v)
	case OpIsNonNil:
		return psess.rewriteValue386_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return psess.rewriteValue386_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return psess.rewriteValue386_OpLeq16_0(v)
	case OpLeq16U:
		return psess.rewriteValue386_OpLeq16U_0(v)
	case OpLeq32:
		return psess.rewriteValue386_OpLeq32_0(v)
	case OpLeq32F:
		return psess.rewriteValue386_OpLeq32F_0(v)
	case OpLeq32U:
		return psess.rewriteValue386_OpLeq32U_0(v)
	case OpLeq64F:
		return psess.rewriteValue386_OpLeq64F_0(v)
	case OpLeq8:
		return psess.rewriteValue386_OpLeq8_0(v)
	case OpLeq8U:
		return psess.rewriteValue386_OpLeq8U_0(v)
	case OpLess16:
		return psess.rewriteValue386_OpLess16_0(v)
	case OpLess16U:
		return psess.rewriteValue386_OpLess16U_0(v)
	case OpLess32:
		return psess.rewriteValue386_OpLess32_0(v)
	case OpLess32F:
		return psess.rewriteValue386_OpLess32F_0(v)
	case OpLess32U:
		return psess.rewriteValue386_OpLess32U_0(v)
	case OpLess64F:
		return psess.rewriteValue386_OpLess64F_0(v)
	case OpLess8:
		return psess.rewriteValue386_OpLess8_0(v)
	case OpLess8U:
		return psess.rewriteValue386_OpLess8U_0(v)
	case OpLoad:
		return psess.rewriteValue386_OpLoad_0(v)
	case OpLsh16x16:
		return psess.rewriteValue386_OpLsh16x16_0(v)
	case OpLsh16x32:
		return psess.rewriteValue386_OpLsh16x32_0(v)
	case OpLsh16x64:
		return rewriteValue386_OpLsh16x64_0(v)
	case OpLsh16x8:
		return psess.rewriteValue386_OpLsh16x8_0(v)
	case OpLsh32x16:
		return psess.rewriteValue386_OpLsh32x16_0(v)
	case OpLsh32x32:
		return psess.rewriteValue386_OpLsh32x32_0(v)
	case OpLsh32x64:
		return rewriteValue386_OpLsh32x64_0(v)
	case OpLsh32x8:
		return psess.rewriteValue386_OpLsh32x8_0(v)
	case OpLsh8x16:
		return psess.rewriteValue386_OpLsh8x16_0(v)
	case OpLsh8x32:
		return psess.rewriteValue386_OpLsh8x32_0(v)
	case OpLsh8x64:
		return rewriteValue386_OpLsh8x64_0(v)
	case OpLsh8x8:
		return psess.rewriteValue386_OpLsh8x8_0(v)
	case OpMod16:
		return rewriteValue386_OpMod16_0(v)
	case OpMod16u:
		return rewriteValue386_OpMod16u_0(v)
	case OpMod32:
		return rewriteValue386_OpMod32_0(v)
	case OpMod32u:
		return rewriteValue386_OpMod32u_0(v)
	case OpMod8:
		return rewriteValue386_OpMod8_0(v)
	case OpMod8u:
		return rewriteValue386_OpMod8u_0(v)
	case OpMove:
		return psess.rewriteValue386_OpMove_0(v) || rewriteValue386_OpMove_10(v)
	case OpMul16:
		return rewriteValue386_OpMul16_0(v)
	case OpMul32:
		return rewriteValue386_OpMul32_0(v)
	case OpMul32F:
		return rewriteValue386_OpMul32F_0(v)
	case OpMul32uhilo:
		return rewriteValue386_OpMul32uhilo_0(v)
	case OpMul64F:
		return rewriteValue386_OpMul64F_0(v)
	case OpMul8:
		return rewriteValue386_OpMul8_0(v)
	case OpNeg16:
		return rewriteValue386_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValue386_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValue386_OpNeg32F_0(v)
	case OpNeg64F:
		return rewriteValue386_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValue386_OpNeg8_0(v)
	case OpNeq16:
		return psess.rewriteValue386_OpNeq16_0(v)
	case OpNeq32:
		return psess.rewriteValue386_OpNeq32_0(v)
	case OpNeq32F:
		return psess.rewriteValue386_OpNeq32F_0(v)
	case OpNeq64F:
		return psess.rewriteValue386_OpNeq64F_0(v)
	case OpNeq8:
		return psess.rewriteValue386_OpNeq8_0(v)
	case OpNeqB:
		return psess.rewriteValue386_OpNeqB_0(v)
	case OpNeqPtr:
		return psess.rewriteValue386_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValue386_OpNilCheck_0(v)
	case OpNot:
		return rewriteValue386_OpNot_0(v)
	case OpOffPtr:
		return rewriteValue386_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValue386_OpOr16_0(v)
	case OpOr32:
		return rewriteValue386_OpOr32_0(v)
	case OpOr8:
		return rewriteValue386_OpOr8_0(v)
	case OpOrB:
		return rewriteValue386_OpOrB_0(v)
	case OpRound32F:
		return rewriteValue386_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValue386_OpRound64F_0(v)
	case OpRsh16Ux16:
		return psess.rewriteValue386_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return psess.rewriteValue386_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return rewriteValue386_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return psess.rewriteValue386_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return psess.rewriteValue386_OpRsh16x16_0(v)
	case OpRsh16x32:
		return psess.rewriteValue386_OpRsh16x32_0(v)
	case OpRsh16x64:
		return rewriteValue386_OpRsh16x64_0(v)
	case OpRsh16x8:
		return psess.rewriteValue386_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return psess.rewriteValue386_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return psess.rewriteValue386_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return rewriteValue386_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return psess.rewriteValue386_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return psess.rewriteValue386_OpRsh32x16_0(v)
	case OpRsh32x32:
		return psess.rewriteValue386_OpRsh32x32_0(v)
	case OpRsh32x64:
		return rewriteValue386_OpRsh32x64_0(v)
	case OpRsh32x8:
		return psess.rewriteValue386_OpRsh32x8_0(v)
	case OpRsh8Ux16:
		return psess.rewriteValue386_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return psess.rewriteValue386_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return rewriteValue386_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return psess.rewriteValue386_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return psess.rewriteValue386_OpRsh8x16_0(v)
	case OpRsh8x32:
		return psess.rewriteValue386_OpRsh8x32_0(v)
	case OpRsh8x64:
		return rewriteValue386_OpRsh8x64_0(v)
	case OpRsh8x8:
		return psess.rewriteValue386_OpRsh8x8_0(v)
	case OpSignExt16to32:
		return rewriteValue386_OpSignExt16to32_0(v)
	case OpSignExt8to16:
		return rewriteValue386_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValue386_OpSignExt8to32_0(v)
	case OpSignmask:
		return rewriteValue386_OpSignmask_0(v)
	case OpSlicemask:
		return rewriteValue386_OpSlicemask_0(v)
	case OpSqrt:
		return rewriteValue386_OpSqrt_0(v)
	case OpStaticCall:
		return rewriteValue386_OpStaticCall_0(v)
	case OpStore:
		return psess.rewriteValue386_OpStore_0(v)
	case OpSub16:
		return rewriteValue386_OpSub16_0(v)
	case OpSub32:
		return rewriteValue386_OpSub32_0(v)
	case OpSub32F:
		return rewriteValue386_OpSub32F_0(v)
	case OpSub32carry:
		return rewriteValue386_OpSub32carry_0(v)
	case OpSub32withcarry:
		return rewriteValue386_OpSub32withcarry_0(v)
	case OpSub64F:
		return rewriteValue386_OpSub64F_0(v)
	case OpSub8:
		return rewriteValue386_OpSub8_0(v)
	case OpSubPtr:
		return rewriteValue386_OpSubPtr_0(v)
	case OpTrunc16to8:
		return rewriteValue386_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValue386_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValue386_OpTrunc32to8_0(v)
	case OpWB:
		return rewriteValue386_OpWB_0(v)
	case OpXor16:
		return rewriteValue386_OpXor16_0(v)
	case OpXor32:
		return rewriteValue386_OpXor32_0(v)
	case OpXor8:
		return rewriteValue386_OpXor8_0(v)
	case OpZero:
		return psess.rewriteValue386_OpZero_0(v) || psess.rewriteValue386_OpZero_10(v)
	case OpZeroExt16to32:
		return rewriteValue386_OpZeroExt16to32_0(v)
	case OpZeroExt8to16:
		return rewriteValue386_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValue386_OpZeroExt8to32_0(v)
	case OpZeromask:
		return psess.rewriteValue386_OpZeromask_0(v)
	}
	return false
}
func rewriteValue386_Op386ADCL_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		f := v.Args[2]
		v.reset(Op386ADCLconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(f)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		f := v.Args[2]
		v.reset(Op386ADCLconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(f)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		f := v.Args[2]
		v.reset(Op386ADCLconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(f)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		f := v.Args[2]
		v.reset(Op386ADCLconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(f)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386ADDL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386ADDLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(Op386ADDLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHRLconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(Op386ROLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHRLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(Op386ROLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHRWconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 16 && d == 16-c && t.Size(psess.types) == 2) {
			break
		}
		v.reset(Op386ROLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHRWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 16 && d == 16-c && t.Size(psess.types) == 2) {
			break
		}
		v.reset(Op386ROLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHRBconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 8 && d == 8-c && t.Size(psess.types) == 1) {
			break
		}
		v.reset(Op386ROLBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHRBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 8 && d == 8-c && t.Size(psess.types) == 1) {
			break
		}
		v.reset(Op386ROLBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 3 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386LEAL8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		if v_0.AuxInt != 3 {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(Op386LEAL8)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValue386_Op386ADDL_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386LEAL4)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(Op386LEAL4)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386LEAL2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(Op386LEAL2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if y != v_1.Args[1] {
			break
		}
		v.reset(Op386LEAL2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		if y != v_0.Args[1] {
			break
		}
		x := v.Args[1]
		v.reset(Op386LEAL2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDL {
			break
		}
		_ = v_1.Args[1]
		if x != v_1.Args[0] {
			break
		}
		y := v_1.Args[1]
		v.reset(Op386LEAL2)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		if x != v_1.Args[1] {
			break
		}
		v.reset(Op386LEAL2)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(Op386LEAL2)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		x := v_0.Args[1]
		if x != v.Args[1] {
			break
		}
		v.reset(Op386LEAL2)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386ADDL_20(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(Op386LEAL1)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		c := v_1.AuxInt
		x := v_1.Args[0]
		v.reset(Op386LEAL1)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		c := v_1.AuxInt
		s := v_1.Aux
		y := v_1.Args[0]
		if !(x.Op != OpSB && y.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL1)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL {
			break
		}
		c := v_0.AuxInt
		s := v_0.Aux
		y := v_0.Args[0]
		x := v.Args[1]
		if !(x.Op != OpSB && y.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL1)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != Op386MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		if !(psess.canMergeLoad(v, l, x) && clobber(l)) {
			break
		}
		v.reset(Op386ADDLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		l := v.Args[0]
		if l.Op != Op386MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		x := v.Args[1]
		if !(psess.canMergeLoad(v, l, x) && clobber(l)) {
			break
		}
		v.reset(Op386ADDLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386NEGL {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386SUBL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386NEGL {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(Op386SUBL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValue386_Op386ADDLcarry_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386ADDLconstcarry)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(Op386ADDLconstcarry)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_Op386ADDLconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(Op386LEAL1)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(Op386LEAL)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL1 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(Op386LEAL1)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL2 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(Op386LEAL2)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL4 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(Op386LEAL4)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL8 {
			break
		}
		d := v_0.AuxInt
		s := v_0.Aux
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(Op386LEAL8)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

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
		if v_0.Op != Op386MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = int64(int32(c + d))
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386ADDLconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_Op386ADDLload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386ADDLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386ADDLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386ADDLmodify_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386ADDLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386ADDLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386ADDSD_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != Op386MOVSDload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		if !(psess.canMergeLoad(v, l, x) && !config.use387 && clobber(l)) {
			break
		}
		v.reset(Op386ADDSDload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		l := v.Args[0]
		if l.Op != Op386MOVSDload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		x := v.Args[1]
		if !(psess.canMergeLoad(v, l, x) && !config.use387 && clobber(l)) {
			break
		}
		v.reset(Op386ADDSDload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386ADDSDload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386ADDSDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386ADDSDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386ADDSS_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != Op386MOVSSload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		if !(psess.canMergeLoad(v, l, x) && !config.use387 && clobber(l)) {
			break
		}
		v.reset(Op386ADDSSload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		l := v.Args[0]
		if l.Op != Op386MOVSSload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		x := v.Args[1]
		if !(psess.canMergeLoad(v, l, x) && !config.use387 && clobber(l)) {
			break
		}
		v.reset(Op386ADDSSload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386ADDSSload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386ADDSSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386ADDSSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386ANDL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386ANDLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(Op386ANDLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != Op386MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		if !(psess.canMergeLoad(v, l, x) && clobber(l)) {
			break
		}
		v.reset(Op386ANDLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		l := v.Args[0]
		if l.Op != Op386MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		x := v.Args[1]
		if !(psess.canMergeLoad(v, l, x) && clobber(l)) {
			break
		}
		v.reset(Op386ANDLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
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
func rewriteValue386_Op386ANDLconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386ANDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386ANDLconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		if !(int32(c) == 0) {
			break
		}
		v.reset(Op386MOVLconst)
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
		if v_0.Op != Op386MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = c & d
		return true
	}
	return false
}
func rewriteValue386_Op386ANDLload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386ANDLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386ANDLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386ANDLmodify_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386ANDLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386ANDLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386CMPB_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386CMPBconst)
		v.AuxInt = int64(int8(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(Op386InvertFlags)
		v0 := b.NewValue0(v.Pos, Op386CMPBconst, psess.types.TypeFlags)
		v0.AuxInt = int64(int8(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValue386_Op386CMPBconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int8(x) == int8(y)) {
			break
		}
		v.reset(Op386FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int8(x) < int8(y) && uint8(x) < uint8(y)) {
			break
		}
		v.reset(Op386FlagLT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int8(x) < int8(y) && uint8(x) > uint8(y)) {
			break
		}
		v.reset(Op386FlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int8(x) > int8(y) && uint8(x) < uint8(y)) {
			break
		}
		v.reset(Op386FlagGT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int8(x) > int8(y) && uint8(x) > uint8(y)) {
			break
		}
		v.reset(Op386FlagGT_UGT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386ANDLconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int8(m) && int8(m) < int8(n)) {
			break
		}
		v.reset(Op386FlagLT_ULT)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != Op386ANDL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(Op386TESTB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != Op386ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386TESTBconst)
		v.AuxInt = int64(int8(c))
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(Op386TESTB)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386CMPL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386CMPLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(Op386InvertFlags)
		v0 := b.NewValue0(v.Pos, Op386CMPLconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValue386_Op386CMPLconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(Op386FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(Op386FlagLT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(Op386FlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(Op386FlagGT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(Op386FlagGT_UGT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386SHRLconst {
			break
		}
		c := v_0.AuxInt
		if !(0 <= n && 0 < c && c <= 32 && (1<<uint64(32-c)) <= uint64(n)) {
			break
		}
		v.reset(Op386FlagLT_ULT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386ANDLconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int32(m) && int32(m) < int32(n)) {
			break
		}
		v.reset(Op386FlagLT_ULT)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != Op386ANDL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(Op386TESTL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != Op386ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386TESTLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(Op386TESTL)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386CMPW_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386CMPWconst)
		v.AuxInt = int64(int16(c))
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(Op386InvertFlags)
		v0 := b.NewValue0(v.Pos, Op386CMPWconst, psess.types.TypeFlags)
		v0.AuxInt = int64(int16(c))
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValue386_Op386CMPWconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int16(x) == int16(y)) {
			break
		}
		v.reset(Op386FlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int16(x) < int16(y) && uint16(x) < uint16(y)) {
			break
		}
		v.reset(Op386FlagLT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int16(x) < int16(y) && uint16(x) > uint16(y)) {
			break
		}
		v.reset(Op386FlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int16(x) > int16(y) && uint16(x) < uint16(y)) {
			break
		}
		v.reset(Op386FlagGT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		x := v_0.AuxInt
		if !(int16(x) > int16(y) && uint16(x) > uint16(y)) {
			break
		}
		v.reset(Op386FlagGT_UGT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386ANDLconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int16(m) && int16(m) < int16(n)) {
			break
		}
		v.reset(Op386FlagLT_ULT)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != Op386ANDL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(Op386TESTW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v_0 := v.Args[0]
		if v_0.Op != Op386ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386TESTWconst)
		v.AuxInt = int64(int16(c))
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		x := v.Args[0]
		v.reset(Op386TESTW)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_Op386LEAL_0(v *Value) bool {

	for {
		c := v.AuxInt
		s := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(Op386LEAL)
		v.AuxInt = c + d
		v.Aux = s
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(x.Op != OpSB && y.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL1)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(Op386LEAL)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL1 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(Op386LEAL1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL2 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(Op386LEAL2)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL4 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(Op386LEAL4)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL8 {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(Op386LEAL8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValue386_Op386LEAL1_0(v *Value) bool {

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		y := v.Args[1]
		if !(is32Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL1)
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
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		x := v_1.Args[0]
		if !(is32Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL1)
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
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386LEAL2)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(Op386LEAL2)
		v.AuxInt = c
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
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386LEAL4)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(Op386LEAL4)
		v.AuxInt = c
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
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 3 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386LEAL8)
		v.AuxInt = c
		v.Aux = s
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		if v_0.AuxInt != 3 {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(Op386LEAL8)
		v.AuxInt = c
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		y := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL1)
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
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		x := v_1.Args[0]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValue386_Op386LEAL2_0(v *Value) bool {

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		y := v.Args[1]
		if !(is32Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL2)
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		if !(is32Bit(c+2*d) && y.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL2)
		v.AuxInt = c + 2*d
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
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386LEAL4)
		v.AuxInt = c
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
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386LEAL8)
		v.AuxInt = c
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		y := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL2)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValue386_Op386LEAL4_0(v *Value) bool {

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		y := v.Args[1]
		if !(is32Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL4)
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		if !(is32Bit(c+4*d) && y.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL4)
		v.AuxInt = c + 4*d
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
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386LEAL8)
		v.AuxInt = c
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		y := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL4)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValue386_Op386LEAL8_0(v *Value) bool {

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		y := v.Args[1]
		if !(is32Bit(c+d) && x.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL8)
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		if !(is32Bit(c+8*d) && y.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL8)
		v.AuxInt = c + 8*d
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		x := v_0.Args[0]
		y := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && x.Op != OpSB) {
			break
		}
		v.reset(Op386LEAL8)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVBLSX_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != Op386MOVBload {
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
		v0 := b.NewValue0(v.Pos, Op386MOVBLSXload, v.Type)
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
		if v_0.Op != Op386ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x80 == 0) {
			break
		}
		v.reset(Op386ANDLconst)
		v.AuxInt = c & 0x7f
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVBLSXload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVBstore {
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
		v.reset(Op386MOVBLSX)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVBLSXload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVBLZX_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != Op386MOVBload {
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
		v0 := b.NewValue0(v.Pos, Op386MOVBload, v.Type)
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
		if x.Op != Op386MOVBloadidx1 {
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
		v0 := b.NewValue0(v.Pos, Op386MOVBloadidx1, v.Type)
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
		if v_0.Op != Op386ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386ANDLconst)
		v.AuxInt = c & 0xff
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVBload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVBstore {
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
		v.reset(Op386MOVBLZX)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MOVBload)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVBload)
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
		if v_0.Op != Op386LEAL1 {
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
		v.reset(Op386MOVBloadidx1)
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
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(Op386MOVBloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVBloadidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVBloadidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVBloadidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVBloadidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVBloadidx1)
		v.AuxInt = int64(int32(c + d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVBstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVBLSX {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVBstore)
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
		if v_1.Op != Op386MOVBLZX {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVBstore)
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MOVBstore)
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
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(validOff(off)) {
			break
		}
		v.reset(Op386MOVBstoreconst)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVBstore)
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
		if v_0.Op != Op386LEAL1 {
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
		v.reset(Op386MOVBstoreidx1)
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
		if v_0.Op != Op386ADDL {
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
		v.reset(Op386MOVBstoreidx1)
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
		if v_1.Op != Op386SHRLconst {
			break
		}
		if v_1.AuxInt != 8 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != Op386MOVBstore {
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
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(Op386MOVWstore)
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
		if v_1.Op != Op386SHRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != Op386MOVBstore {
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
		if w0.Op != Op386SHRLconst {
			break
		}
		if w0.AuxInt != j-8 {
			break
		}
		if w != w0.Args[0] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && clobber(x)) {
			break
		}
		v.reset(Op386MOVWstore)
		v.AuxInt = i - 1
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVBstoreconst_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		sc := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(Op386MOVBstoreconst)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVBstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL1 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(Op386MOVBstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		v.reset(Op386MOVBstoreconstidx1)
		v.AuxInt = x
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != Op386MOVBstoreconst {
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
		if !(x.Uses == 1 && ValAndOff(a).Off()+1 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(Op386MOVWstoreconst)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xff|ValAndOff(c).Val()<<8, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVBstoreconstidx1_0(v *Value) bool {

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVBstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVBstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		i := v.Args[1]
		x := v.Args[2]
		if x.Op != Op386MOVBstoreconstidx1 {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		if i != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && ValAndOff(a).Off()+1 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(Op386MOVWstoreconstidx1)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xff|ValAndOff(c).Val()<<8, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(i)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVBstoreidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVBstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVBstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVBstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVBstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		v_2 := v.Args[2]
		if v_2.Op != Op386SHRLconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVBstoreidx1 {
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
		v.reset(Op386MOVWstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVBstoreidx1 {
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
		v.reset(Op386MOVWstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVBstoreidx1 {
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
		v.reset(Op386MOVWstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		if v_2.AuxInt != 8 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVBstoreidx1 {
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
		v.reset(Op386MOVWstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVBstoreidx1 {
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
		if w0.Op != Op386SHRLconst {
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
		v.reset(Op386MOVWstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVBstoreidx1 {
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
		if w0.Op != Op386SHRLconst {
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
		v.reset(Op386MOVWstoreidx1)
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
func rewriteValue386_Op386MOVBstoreidx1_10(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		idx := v.Args[0]
		p := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != Op386SHRLconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVBstoreidx1 {
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
		if w0.Op != Op386SHRLconst {
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
		v.reset(Op386MOVWstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVBstoreidx1 {
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
		if w0.Op != Op386SHRLconst {
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
		v.reset(Op386MOVWstoreidx1)
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
func rewriteValue386_Op386MOVLload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLstore {
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

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MOVLload)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVLload)
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
		if v_0.Op != Op386LEAL1 {
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
		v.reset(Op386MOVLloadidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
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
		if v_0.Op != Op386LEAL4 {
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
		v.reset(Op386MOVLloadidx4)
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
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(Op386MOVLloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVLloadidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVLloadidx4)
		v.AuxInt = c
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
		if v_0.Op != Op386SHLLconst {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVLloadidx4)
		v.AuxInt = c
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVLloadidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVLloadidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVLloadidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVLloadidx1)
		v.AuxInt = int64(int32(c + d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVLloadidx4_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVLloadidx4)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVLloadidx4)
		v.AuxInt = int64(int32(c + 4*d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVLstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MOVLstore)
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
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(validOff(off)) {
			break
		}
		v.reset(Op386MOVLstoreconst)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVLstore)
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
		if v_0.Op != Op386LEAL1 {
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
		v.reset(Op386MOVLstoreidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
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
		if v_0.Op != Op386LEAL4 {
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
		v.reset(Op386MOVLstoreidx4)
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
		if v_0.Op != Op386ADDL {
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
		v.reset(Op386MOVLstoreidx1)
		v.AuxInt = off
		v.Aux = sym
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
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != Op386ADDLload {
			break
		}
		if y.AuxInt != off {
			break
		}
		if y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] {
			break
		}
		mem := y.Args[2]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(Op386ADDLmodify)
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
		y := v.Args[1]
		if y.Op != Op386ANDLload {
			break
		}
		if y.AuxInt != off {
			break
		}
		if y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] {
			break
		}
		mem := y.Args[2]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(Op386ANDLmodify)
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
		y := v.Args[1]
		if y.Op != Op386ORLload {
			break
		}
		if y.AuxInt != off {
			break
		}
		if y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] {
			break
		}
		mem := y.Args[2]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(Op386ORLmodify)
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
		y := v.Args[1]
		if y.Op != Op386XORLload {
			break
		}
		if y.AuxInt != off {
			break
		}
		if y.Aux != sym {
			break
		}
		_ = y.Args[2]
		x := y.Args[0]
		if ptr != y.Args[1] {
			break
		}
		mem := y.Args[2]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && clobber(y)) {
			break
		}
		v.reset(Op386XORLmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVLstore_10(v *Value) bool {

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		y := v.Args[1]
		if y.Op != Op386ADDL {
			break
		}
		_ = y.Args[1]
		l := y.Args[0]
		if l.Op != Op386MOVLload {
			break
		}
		if l.AuxInt != off {
			break
		}
		if l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] {
			break
		}
		mem := l.Args[1]
		x := y.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(Op386ADDLmodify)
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
		y := v.Args[1]
		if y.Op != Op386ADDL {
			break
		}
		_ = y.Args[1]
		x := y.Args[0]
		l := y.Args[1]
		if l.Op != Op386MOVLload {
			break
		}
		if l.AuxInt != off {
			break
		}
		if l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] {
			break
		}
		mem := l.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(Op386ADDLmodify)
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
		y := v.Args[1]
		if y.Op != Op386SUBL {
			break
		}
		_ = y.Args[1]
		l := y.Args[0]
		if l.Op != Op386MOVLload {
			break
		}
		if l.AuxInt != off {
			break
		}
		if l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] {
			break
		}
		mem := l.Args[1]
		x := y.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(Op386SUBLmodify)
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
		y := v.Args[1]
		if y.Op != Op386ANDL {
			break
		}
		_ = y.Args[1]
		l := y.Args[0]
		if l.Op != Op386MOVLload {
			break
		}
		if l.AuxInt != off {
			break
		}
		if l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] {
			break
		}
		mem := l.Args[1]
		x := y.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(Op386ANDLmodify)
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
		y := v.Args[1]
		if y.Op != Op386ANDL {
			break
		}
		_ = y.Args[1]
		x := y.Args[0]
		l := y.Args[1]
		if l.Op != Op386MOVLload {
			break
		}
		if l.AuxInt != off {
			break
		}
		if l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] {
			break
		}
		mem := l.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(Op386ANDLmodify)
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
		y := v.Args[1]
		if y.Op != Op386ORL {
			break
		}
		_ = y.Args[1]
		l := y.Args[0]
		if l.Op != Op386MOVLload {
			break
		}
		if l.AuxInt != off {
			break
		}
		if l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] {
			break
		}
		mem := l.Args[1]
		x := y.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(Op386ORLmodify)
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
		y := v.Args[1]
		if y.Op != Op386ORL {
			break
		}
		_ = y.Args[1]
		x := y.Args[0]
		l := y.Args[1]
		if l.Op != Op386MOVLload {
			break
		}
		if l.AuxInt != off {
			break
		}
		if l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] {
			break
		}
		mem := l.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(Op386ORLmodify)
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
		y := v.Args[1]
		if y.Op != Op386XORL {
			break
		}
		_ = y.Args[1]
		l := y.Args[0]
		if l.Op != Op386MOVLload {
			break
		}
		if l.AuxInt != off {
			break
		}
		if l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] {
			break
		}
		mem := l.Args[1]
		x := y.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(Op386XORLmodify)
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
		y := v.Args[1]
		if y.Op != Op386XORL {
			break
		}
		_ = y.Args[1]
		x := y.Args[0]
		l := y.Args[1]
		if l.Op != Op386MOVLload {
			break
		}
		if l.AuxInt != off {
			break
		}
		if l.Aux != sym {
			break
		}
		_ = l.Args[1]
		if ptr != l.Args[0] {
			break
		}
		mem := l.Args[1]
		if mem != v.Args[2] {
			break
		}
		if !(y.Uses == 1 && l.Uses == 1 && clobber(y) && clobber(l)) {
			break
		}
		v.reset(Op386XORLmodify)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVLstoreconst_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		sc := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(Op386MOVLstoreconst)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVLstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL1 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(Op386MOVLstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL4 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(Op386MOVLstoreconstidx4)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		v.reset(Op386MOVLstoreconstidx1)
		v.AuxInt = x
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVLstoreconstidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVLstoreconstidx4)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVLstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVLstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVLstoreconstidx4_0(v *Value) bool {

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVLstoreconstidx4)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVLstoreconstidx4)
		v.AuxInt = ValAndOff(x).add(4 * c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVLstoreidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 2 {
			break
		}
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVLstoreidx4)
		v.AuxInt = c
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
		if v_0.Op != Op386SHLLconst {
			break
		}
		if v_0.AuxInt != 2 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVLstoreidx4)
		v.AuxInt = c
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVLstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVLstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVLstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVLstoreidx1)
		v.AuxInt = int64(int32(c + d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVLstoreidx4_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVLstoreidx4)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVLstoreidx4)
		v.AuxInt = int64(int32(c + 4*d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSDconst_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		c := v.AuxInt
		if !(config.ctxt.Flag_shared) {
			break
		}
		v.reset(Op386MOVSDconst2)
		v0 := b.NewValue0(v.Pos, Op386MOVSDconst1, typ.UInt32)
		v0.AuxInt = c
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSDload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MOVSDload)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVSDload)
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
		if v_0.Op != Op386LEAL1 {
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
		v.reset(Op386MOVSDloadidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
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
		if v_0.Op != Op386LEAL8 {
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
		v.reset(Op386MOVSDloadidx8)
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
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(Op386MOVSDloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSDloadidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVSDloadidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVSDloadidx1)
		v.AuxInt = int64(int32(c + d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSDloadidx8_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVSDloadidx8)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVSDloadidx8)
		v.AuxInt = int64(int32(c + 8*d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSDstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MOVSDstore)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVSDstore)
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
		if v_0.Op != Op386LEAL1 {
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
		v.reset(Op386MOVSDstoreidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
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
		if v_0.Op != Op386LEAL8 {
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
		v.reset(Op386MOVSDstoreidx8)
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
		if v_0.Op != Op386ADDL {
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
		v.reset(Op386MOVSDstoreidx1)
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
func rewriteValue386_Op386MOVSDstoreidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVSDstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVSDstoreidx1)
		v.AuxInt = int64(int32(c + d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSDstoreidx8_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVSDstoreidx8)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVSDstoreidx8)
		v.AuxInt = int64(int32(c + 8*d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSSconst_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		c := v.AuxInt
		if !(config.ctxt.Flag_shared) {
			break
		}
		v.reset(Op386MOVSSconst2)
		v0 := b.NewValue0(v.Pos, Op386MOVSSconst1, typ.UInt32)
		v0.AuxInt = c
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSSload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MOVSSload)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVSSload)
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
		if v_0.Op != Op386LEAL1 {
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
		v.reset(Op386MOVSSloadidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
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
		if v_0.Op != Op386LEAL4 {
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
		v.reset(Op386MOVSSloadidx4)
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
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(Op386MOVSSloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSSloadidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVSSloadidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVSSloadidx1)
		v.AuxInt = int64(int32(c + d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSSloadidx4_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVSSloadidx4)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVSSloadidx4)
		v.AuxInt = int64(int32(c + 4*d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSSstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MOVSSstore)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVSSstore)
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
		if v_0.Op != Op386LEAL1 {
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
		v.reset(Op386MOVSSstoreidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
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
		if v_0.Op != Op386LEAL4 {
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
		v.reset(Op386MOVSSstoreidx4)
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
		if v_0.Op != Op386ADDL {
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
		v.reset(Op386MOVSSstoreidx1)
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
func rewriteValue386_Op386MOVSSstoreidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVSSstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVSSstoreidx1)
		v.AuxInt = int64(int32(c + d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVSSstoreidx4_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVSSstoreidx4)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVSSstoreidx4)
		v.AuxInt = int64(int32(c + 4*d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVWLSX_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != Op386MOVWload {
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
		v0 := b.NewValue0(v.Pos, Op386MOVWLSXload, v.Type)
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
		if v_0.Op != Op386ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x8000 == 0) {
			break
		}
		v.reset(Op386ANDLconst)
		v.AuxInt = c & 0x7fff
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVWLSXload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVWstore {
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
		v.reset(Op386MOVWLSX)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVWLSXload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVWLZX_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.Args[0]
		if x.Op != Op386MOVWload {
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
		v0 := b.NewValue0(v.Pos, Op386MOVWload, v.Type)
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
		if x.Op != Op386MOVWloadidx1 {
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
		v0 := b.NewValue0(v.Pos, Op386MOVWloadidx1, v.Type)
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
		if x.Op != Op386MOVWloadidx2 {
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
		v0 := b.NewValue0(v.Pos, Op386MOVWloadidx2, v.Type)
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
		if v_0.Op != Op386ANDLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386ANDLconst)
		v.AuxInt = c & 0xffff
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVWload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVWstore {
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
		v.reset(Op386MOVWLZX)
		v.AddArg(x)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MOVWload)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		mem := v.Args[1]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVWload)
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
		if v_0.Op != Op386LEAL1 {
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
		v.reset(Op386MOVWloadidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
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
		if v_0.Op != Op386LEAL2 {
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
		v.reset(Op386MOVWloadidx2)
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
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(ptr.Op != OpSB) {
			break
		}
		v.reset(Op386MOVWloadidx1)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVWloadidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVWloadidx2)
		v.AuxInt = c
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
		if v_0.Op != Op386SHLLconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVWloadidx2)
		v.AuxInt = c
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVWloadidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVWloadidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVWloadidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVWloadidx1)
		v.AuxInt = int64(int32(c + d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVWloadidx2_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVWloadidx2)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVWloadidx2)
		v.AuxInt = int64(int32(c + 2*d))
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVWstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVWLSX {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVWstore)
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
		if v_1.Op != Op386MOVWLZX {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVWstore)
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MOVWstore)
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
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		if !(validOff(off)) {
			break
		}
		v.reset(Op386MOVWstoreconst)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVWstore)
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
		if v_0.Op != Op386LEAL1 {
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
		v.reset(Op386MOVWstoreidx1)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
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
		if v_0.Op != Op386LEAL2 {
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
		v.reset(Op386MOVWstoreidx2)
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
		if v_0.Op != Op386ADDL {
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
		v.reset(Op386MOVWstoreidx1)
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
		if v_1.Op != Op386SHRLconst {
			break
		}
		if v_1.AuxInt != 16 {
			break
		}
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != Op386MOVWstore {
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
		v.reset(Op386MOVLstore)
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
		if v_1.Op != Op386SHRLconst {
			break
		}
		j := v_1.AuxInt
		w := v_1.Args[0]
		x := v.Args[2]
		if x.Op != Op386MOVWstore {
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
		if w0.Op != Op386SHRLconst {
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
		v.reset(Op386MOVLstore)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVWstoreconst_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		sc := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(ValAndOff(sc).canAdd(off)) {
			break
		}
		v.reset(Op386MOVWstoreconst)
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
		if v_0.Op != Op386LEAL {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2) && ValAndOff(sc).canAdd(off) && (ptr.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MOVWstoreconst)
		v.AuxInt = ValAndOff(sc).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL1 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(Op386MOVWstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386LEAL2 {
			break
		}
		off := v_0.AuxInt
		sym2 := v_0.Aux
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(Op386MOVWstoreconstidx2)
		v.AuxInt = ValAndOff(x).add(off)
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDL {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		v.reset(Op386MOVWstoreconstidx1)
		v.AuxInt = x
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[1]
		p := v.Args[0]
		x := v.Args[1]
		if x.Op != Op386MOVWstoreconst {
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
		if !(x.Uses == 1 && ValAndOff(a).Off()+2 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(Op386MOVLstoreconst)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xffff|ValAndOff(c).Val()<<16, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVWstoreconstidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVWstoreconstidx2)
		v.AuxInt = c
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVWstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVWstoreconstidx1)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		i := v.Args[1]
		x := v.Args[2]
		if x.Op != Op386MOVWstoreconstidx1 {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		if i != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && ValAndOff(a).Off()+2 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(Op386MOVLstoreconstidx1)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xffff|ValAndOff(c).Val()<<16, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v.AddArg(i)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVWstoreconstidx2_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVWstoreconstidx2)
		v.AuxInt = ValAndOff(x).add(c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		x := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(Op386MOVWstoreconstidx2)
		v.AuxInt = ValAndOff(x).add(2 * c)
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		c := v.AuxInt
		s := v.Aux
		_ = v.Args[2]
		p := v.Args[0]
		i := v.Args[1]
		x := v.Args[2]
		if x.Op != Op386MOVWstoreconstidx2 {
			break
		}
		a := x.AuxInt
		if x.Aux != s {
			break
		}
		_ = x.Args[2]
		if p != x.Args[0] {
			break
		}
		if i != x.Args[1] {
			break
		}
		mem := x.Args[2]
		if !(x.Uses == 1 && ValAndOff(a).Off()+2 == ValAndOff(c).Off() && clobber(x)) {
			break
		}
		v.reset(Op386MOVLstoreconstidx1)
		v.AuxInt = makeValAndOff(ValAndOff(a).Val()&0xffff|ValAndOff(c).Val()<<16, ValAndOff(a).Off())
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, Op386SHLLconst, i.Type)
		v0.AuxInt = 1
		v0.AddArg(i)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MOVWstoreidx1_0(v *Value) bool {

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVWstoreidx2)
		v.AuxInt = c
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
		if v_0.Op != Op386SHLLconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVWstoreidx2)
		v.AuxInt = c
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVWstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		ptr := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVWstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVWstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVWstoreidx1)
		v.AuxInt = int64(int32(c + d))
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
		v_2 := v.Args[2]
		if v_2.Op != Op386SHRLconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVWstoreidx1 {
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
		v.reset(Op386MOVLstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVWstoreidx1 {
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
		v.reset(Op386MOVLstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVWstoreidx1 {
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
		v.reset(Op386MOVLstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVWstoreidx1 {
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
		v.reset(Op386MOVLstoreidx1)
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
func rewriteValue386_Op386MOVWstoreidx1_10(v *Value) bool {

	for {
		i := v.AuxInt
		s := v.Aux
		_ = v.Args[3]
		p := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != Op386SHRLconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVWstoreidx1 {
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
		if w0.Op != Op386SHRLconst {
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
		v.reset(Op386MOVLstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVWstoreidx1 {
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
		if w0.Op != Op386SHRLconst {
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
		v.reset(Op386MOVLstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVWstoreidx1 {
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
		if w0.Op != Op386SHRLconst {
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
		v.reset(Op386MOVLstoreidx1)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVWstoreidx1 {
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
		if w0.Op != Op386SHRLconst {
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
		v.reset(Op386MOVLstoreidx1)
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
func rewriteValue386_Op386MOVWstoreidx2_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		c := v.AuxInt
		sym := v.Aux
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		d := v_0.AuxInt
		ptr := v_0.Args[0]
		idx := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVWstoreidx2)
		v.AuxInt = int64(int32(c + d))
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
		if v_1.Op != Op386ADDLconst {
			break
		}
		d := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(Op386MOVWstoreidx2)
		v.AuxInt = int64(int32(c + 2*d))
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
		v_2 := v.Args[2]
		if v_2.Op != Op386SHRLconst {
			break
		}
		if v_2.AuxInt != 16 {
			break
		}
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVWstoreidx2 {
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
		v.reset(Op386MOVLstoreidx1)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, Op386SHLLconst, idx.Type)
		v0.AuxInt = 1
		v0.AddArg(idx)
		v.AddArg(v0)
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
		if v_2.Op != Op386SHRLconst {
			break
		}
		j := v_2.AuxInt
		w := v_2.Args[0]
		x := v.Args[3]
		if x.Op != Op386MOVWstoreidx2 {
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
		if w0.Op != Op386SHRLconst {
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
		v.reset(Op386MOVLstoreidx1)
		v.AuxInt = i - 2
		v.Aux = s
		v.AddArg(p)
		v0 := b.NewValue0(v.Pos, Op386SHLLconst, idx.Type)
		v0.AuxInt = 1
		v0.AddArg(idx)
		v.AddArg(v0)
		v.AddArg(w0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MULL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386MULLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(Op386MULLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_Op386MULLconst_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MULLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386MULLconst)
		v.AuxInt = int64(int32(c * d))
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != -9 {
			break
		}
		x := v.Args[0]
		v.reset(Op386NEGL)
		v0 := b.NewValue0(v.Pos, Op386LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != -5 {
			break
		}
		x := v.Args[0]
		v.reset(Op386NEGL)
		v0 := b.NewValue0(v.Pos, Op386LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != -3 {
			break
		}
		x := v.Args[0]
		v.reset(Op386NEGL)
		v0 := b.NewValue0(v.Pos, Op386LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != -1 {
			break
		}
		x := v.Args[0]
		v.reset(Op386NEGL)
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(Op386MOVLconst)
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
		if v.AuxInt != 3 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL2)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 5 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL4)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 7 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL2)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValue386_Op386MULLconst_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		if v.AuxInt != 9 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL8)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 11 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL2)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 13 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL4)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 19 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL2)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 21 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL4)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 25 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 27 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL8)
		v0 := b.NewValue0(v.Pos, Op386LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386LEAL2, v.Type)
		v1.AddArg(x)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 37 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL4)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 41 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 45 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL8)
		v0 := b.NewValue0(v.Pos, Op386LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386LEAL4, v.Type)
		v1.AddArg(x)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValue386_Op386MULLconst_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		if v.AuxInt != 73 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL8)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 81 {
			break
		}
		x := v.Args[0]
		v.reset(Op386LEAL8)
		v0 := b.NewValue0(v.Pos, Op386LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386LEAL8, v.Type)
		v1.AddArg(x)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c+1) && c >= 15) {
			break
		}
		v.reset(Op386SUBL)
		v0 := b.NewValue0(v.Pos, Op386SHLLconst, v.Type)
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
		v.reset(Op386LEAL1)
		v0 := b.NewValue0(v.Pos, Op386SHLLconst, v.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-2) && c >= 34) {
			break
		}
		v.reset(Op386LEAL2)
		v0 := b.NewValue0(v.Pos, Op386SHLLconst, v.Type)
		v0.AuxInt = log2(c - 2)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-4) && c >= 68) {
			break
		}
		v.reset(Op386LEAL4)
		v0 := b.NewValue0(v.Pos, Op386SHLLconst, v.Type)
		v0.AuxInt = log2(c - 4)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(isPowerOfTwo(c-8) && c >= 136) {
			break
		}
		v.reset(Op386LEAL8)
		v0 := b.NewValue0(v.Pos, Op386SHLLconst, v.Type)
		v0.AuxInt = log2(c - 8)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(c%3 == 0 && isPowerOfTwo(c/3)) {
			break
		}
		v.reset(Op386SHLLconst)
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, Op386LEAL2, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(c%5 == 0 && isPowerOfTwo(c/5)) {
			break
		}
		v.reset(Op386SHLLconst)
		v.AuxInt = log2(c / 5)
		v0 := b.NewValue0(v.Pos, Op386LEAL4, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(c%9 == 0 && isPowerOfTwo(c/9)) {
			break
		}
		v.reset(Op386SHLLconst)
		v.AuxInt = log2(c / 9)
		v0 := b.NewValue0(v.Pos, Op386LEAL8, v.Type)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValue386_Op386MULLconst_30(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = int64(int32(c * d))
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386MULSD_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != Op386MOVSDload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		if !(psess.canMergeLoad(v, l, x) && !config.use387 && clobber(l)) {
			break
		}
		v.reset(Op386MULSDload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		l := v.Args[0]
		if l.Op != Op386MOVSDload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		x := v.Args[1]
		if !(psess.canMergeLoad(v, l, x) && !config.use387 && clobber(l)) {
			break
		}
		v.reset(Op386MULSDload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MULSDload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MULSDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MULSDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386MULSS_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != Op386MOVSSload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		if !(psess.canMergeLoad(v, l, x) && !config.use387 && clobber(l)) {
			break
		}
		v.reset(Op386MULSSload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		l := v.Args[0]
		if l.Op != Op386MOVSSload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		x := v.Args[1]
		if !(psess.canMergeLoad(v, l, x) && !config.use387 && clobber(l)) {
			break
		}
		v.reset(Op386MULSSload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386MULSSload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386MULSSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386MULSSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386NEGL_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = int64(int32(-c))
		return true
	}
	return false
}
func rewriteValue386_Op386NOTL_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = ^c
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386ORL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386ORLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(Op386ORLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHRLconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(Op386ROLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHRLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(Op386ROLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHRWconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 16 && d == 16-c && t.Size(psess.types) == 2) {
			break
		}
		v.reset(Op386ROLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHRWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 16 && d == 16-c && t.Size(psess.types) == 2) {
			break
		}
		v.reset(Op386ROLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHRBconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 8 && d == 8-c && t.Size(psess.types) == 1) {
			break
		}
		v.reset(Op386ROLBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHRBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 8 && d == 8-c && t.Size(psess.types) == 1) {
			break
		}
		v.reset(Op386ROLBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != Op386MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		if !(psess.canMergeLoad(v, l, x) && clobber(l)) {
			break
		}
		v.reset(Op386ORLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		l := v.Args[0]
		if l.Op != Op386MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		x := v.Args[1]
		if !(psess.canMergeLoad(v, l, x) && clobber(l)) {
			break
		}
		v.reset(Op386ORLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386ORL_10(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

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
		x0 := v.Args[0]
		if x0.Op != Op386MOVBload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		s0 := v.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBload {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, Op386MOVWload, typ.UInt16)
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
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBload {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := v.Args[1]
		if x0.Op != Op386MOVBload {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, Op386MOVWload, typ.UInt16)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWload {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[1]
		p := x0.Args[0]
		mem := x0.Args[1]
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBload {
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBload {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBload {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[1]
		p := x1.Args[0]
		mem := x1.Args[1]
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWload {
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBload {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBload {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWload {
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
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBload {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBload {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[1]
		p := x2.Args[0]
		mem := x2.Args[1]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBload {
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
		if x0.Op != Op386MOVWload {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
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
		x0 := v.Args[0]
		if x0.Op != Op386MOVBloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s0 := v.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, Op386MOVWloadidx1, v.Type)
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
		x0 := v.Args[0]
		if x0.Op != Op386MOVBloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s0 := v.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, Op386MOVWloadidx1, v.Type)
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
		x0 := v.Args[0]
		if x0.Op != Op386MOVBloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s0 := v.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, Op386MOVWloadidx1, v.Type)
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
func rewriteValue386_Op386ORL_20(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x0 := v.Args[0]
		if x0.Op != Op386MOVBloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s0 := v.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, Op386MOVWloadidx1, v.Type)
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
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != Op386MOVBloadidx1 {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, Op386MOVWloadidx1, v.Type)
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
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != Op386MOVBloadidx1 {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, Op386MOVWloadidx1, v.Type)
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
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != Op386MOVBloadidx1 {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, Op386MOVWloadidx1, v.Type)
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
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 8 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i1 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		x0 := v.Args[1]
		if x0.Op != Op386MOVBloadidx1 {
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
		if !(i1 == i0+1 && x0.Uses == 1 && x1.Uses == 1 && s0.Uses == 1 && mergePoint(b, x0, x1) != nil && clobber(x0) && clobber(x1) && clobber(s0)) {
			break
		}
		b = mergePoint(b, x0, x1)
		v0 := b.NewValue0(v.Pos, Op386MOVWloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		if idx != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		if idx != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		if idx != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		if idx != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		if idx != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
func rewriteValue386_Op386ORL_30(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		if idx != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		if idx != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if p != x2.Args[0] {
			break
		}
		if idx != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if idx != x2.Args[0] {
			break
		}
		if p != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if idx != x2.Args[0] {
			break
		}
		if p != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		p := x0.Args[0]
		idx := x0.Args[1]
		mem := x0.Args[2]
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if idx != x2.Args[0] {
			break
		}
		if p != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
			break
		}
		i0 := x0.AuxInt
		s := x0.Aux
		_ = x0.Args[2]
		idx := x0.Args[0]
		p := x0.Args[1]
		mem := x0.Args[2]
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if idx != x2.Args[0] {
			break
		}
		if p != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if idx != x2.Args[0] {
			break
		}
		if p != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if idx != x2.Args[0] {
			break
		}
		if p != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		p := x1.Args[0]
		idx := x1.Args[1]
		mem := x1.Args[2]
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if idx != x2.Args[0] {
			break
		}
		if p != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
func rewriteValue386_Op386ORL_40(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		o0 := v.Args[0]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
		s := x1.Aux
		_ = x1.Args[2]
		idx := x1.Args[0]
		p := x1.Args[1]
		mem := x1.Args[2]
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		s1 := v.Args[1]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		if x2.Aux != s {
			break
		}
		_ = x2.Args[2]
		if idx != x2.Args[0] {
			break
		}
		if p != x2.Args[1] {
			break
		}
		if mem != x2.Args[2] {
			break
		}
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		p := x2.Args[0]
		idx := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
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
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		idx := x2.Args[0]
		p := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
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
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		p := x2.Args[0]
		idx := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
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
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		idx := x2.Args[0]
		p := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
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
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		p := x2.Args[0]
		idx := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
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
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		idx := x2.Args[0]
		p := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
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
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		p := x2.Args[0]
		idx := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
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
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		idx := x2.Args[0]
		p := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		x0 := o0.Args[0]
		if x0.Op != Op386MOVWloadidx1 {
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
		s0 := o0.Args[1]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		p := x2.Args[0]
		idx := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
func rewriteValue386_Op386ORL_50(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		s1 := v.Args[0]
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		idx := x2.Args[0]
		p := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		p := x2.Args[0]
		idx := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		idx := x2.Args[0]
		p := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		p := x2.Args[0]
		idx := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		idx := x2.Args[0]
		p := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		p := x2.Args[0]
		idx := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
		if s1.Op != Op386SHLLconst {
			break
		}
		if s1.AuxInt != 24 {
			break
		}
		x2 := s1.Args[0]
		if x2.Op != Op386MOVBloadidx1 {
			break
		}
		i3 := x2.AuxInt
		s := x2.Aux
		_ = x2.Args[2]
		idx := x2.Args[0]
		p := x2.Args[1]
		mem := x2.Args[2]
		o0 := v.Args[1]
		if o0.Op != Op386ORL {
			break
		}
		_ = o0.Args[1]
		s0 := o0.Args[0]
		if s0.Op != Op386SHLLconst {
			break
		}
		if s0.AuxInt != 16 {
			break
		}
		x1 := s0.Args[0]
		if x1.Op != Op386MOVBloadidx1 {
			break
		}
		i2 := x1.AuxInt
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
		x0 := o0.Args[1]
		if x0.Op != Op386MOVWloadidx1 {
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
		if !(i2 == i0+2 && i3 == i0+3 && x0.Uses == 1 && x1.Uses == 1 && x2.Uses == 1 && s0.Uses == 1 && s1.Uses == 1 && o0.Uses == 1 && mergePoint(b, x0, x1, x2) != nil && clobber(x0) && clobber(x1) && clobber(x2) && clobber(s0) && clobber(s1) && clobber(o0)) {
			break
		}
		b = mergePoint(b, x0, x1, x2)
		v0 := b.NewValue0(v.Pos, Op386MOVLloadidx1, v.Type)
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
func rewriteValue386_Op386ORLconst_0(v *Value) bool {

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
		v.reset(Op386MOVLconst)
		v.AuxInt = -1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = c | d
		return true
	}
	return false
}
func rewriteValue386_Op386ORLload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386ORLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386ORLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386ORLmodify_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386ORLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386ORLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386ROLBconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386ROLBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386ROLBconst)
		v.AuxInt = (c + d) & 7
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
func rewriteValue386_Op386ROLLconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386ROLLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386ROLLconst)
		v.AuxInt = (c + d) & 31
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
func rewriteValue386_Op386ROLWconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386ROLWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386ROLWconst)
		v.AuxInt = (c + d) & 15
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
func rewriteValue386_Op386SARB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386SARBconst)
		v.AuxInt = min(c&31, 7)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_Op386SARBconst_0(v *Value) bool {

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
		if v_0.Op != Op386MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = d >> uint64(c)
		return true
	}
	return false
}
func rewriteValue386_Op386SARL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386SARLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ANDLconst {
			break
		}
		if v_1.AuxInt != 31 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386SARL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValue386_Op386SARLconst_0(v *Value) bool {

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
		if v_0.Op != Op386MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = d >> uint64(c)
		return true
	}
	return false
}
func rewriteValue386_Op386SARW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386SARWconst)
		v.AuxInt = min(c&31, 15)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_Op386SARWconst_0(v *Value) bool {

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
		if v_0.Op != Op386MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = d >> uint64(c)
		return true
	}
	return false
}
func rewriteValue386_Op386SBBL_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		f := v.Args[2]
		v.reset(Op386SBBLconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(f)
		return true
	}
	return false
}
func rewriteValue386_Op386SBBLcarrymask_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagEQ {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = -1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = -1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValue386_Op386SETA_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(Op386SETB)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagEQ {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValue386_Op386SETAE_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(Op386SETBE)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagEQ {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValue386_Op386SETB_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(Op386SETA)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagEQ {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValue386_Op386SETBE_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(Op386SETAE)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagEQ {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValue386_Op386SETEQ_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(Op386SETEQ)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagEQ {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValue386_Op386SETG_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(Op386SETL)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagEQ {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValue386_Op386SETGE_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(Op386SETLE)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagEQ {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValue386_Op386SETL_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(Op386SETG)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagEQ {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValue386_Op386SETLE_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(Op386SETGE)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagEQ {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValue386_Op386SETNE_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386InvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(Op386SETNE)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagEQ {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagLT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_ULT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != Op386FlagGT_UGT {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValue386_Op386SHLL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386SHLLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ANDLconst {
			break
		}
		if v_1.AuxInt != 31 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386SHLL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValue386_Op386SHLLconst_0(v *Value) bool {

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
func rewriteValue386_Op386SHRB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 < 8) {
			break
		}
		v.reset(Op386SHRBconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 >= 8) {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValue386_Op386SHRBconst_0(v *Value) bool {

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
func rewriteValue386_Op386SHRL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386SHRLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ANDLconst {
			break
		}
		if v_1.AuxInt != 31 {
			break
		}
		y := v_1.Args[0]
		v.reset(Op386SHRL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValue386_Op386SHRLconst_0(v *Value) bool {

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
func rewriteValue386_Op386SHRW_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 < 16) {
			break
		}
		v.reset(Op386SHRWconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		if !(c&31 >= 16) {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValue386_Op386SHRWconst_0(v *Value) bool {

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
func (psess *PackageSession) rewriteValue386_Op386SUBL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386SUBLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(Op386NEGL)
		v0 := b.NewValue0(v.Pos, Op386SUBLconst, v.Type)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != Op386MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		if !(psess.canMergeLoad(v, l, x) && clobber(l)) {
			break
		}
		v.reset(Op386SUBLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValue386_Op386SUBLcarry_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386SUBLconstcarry)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_Op386SUBLconst_0(v *Value) bool {

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
		v.reset(Op386ADDLconst)
		v.AuxInt = int64(int32(-c))
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_Op386SUBLload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386SUBLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386SUBLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386SUBLmodify_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386SUBLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386SUBLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386SUBSD_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != Op386MOVSDload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		if !(psess.canMergeLoad(v, l, x) && !config.use387 && clobber(l)) {
			break
		}
		v.reset(Op386SUBSDload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386SUBSDload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386SUBSDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386SUBSDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386SUBSS_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != Op386MOVSSload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		if !(psess.canMergeLoad(v, l, x) && !config.use387 && clobber(l)) {
			break
		}
		v.reset(Op386SUBSSload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386SUBSSload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386SUBSSload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386SUBSSload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_Op386XORL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386MOVLconst {
			break
		}
		c := v_1.AuxInt
		v.reset(Op386XORLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386MOVLconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(Op386XORLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHRLconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(Op386ROLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHRLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(d == 32-c) {
			break
		}
		v.reset(Op386ROLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHRWconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 16 && d == 16-c && t.Size(psess.types) == 2) {
			break
		}
		v.reset(Op386ROLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHRWconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 16 && d == 16-c && t.Size(psess.types) == 2) {
			break
		}
		v.reset(Op386ROLWconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHRBconst {
			break
		}
		d := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 8 && d == 8-c && t.Size(psess.types) == 1) {
			break
		}
		v.reset(Op386ROLBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != Op386SHRBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386SHLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c < 8 && d == 8-c && t.Size(psess.types) == 1) {
			break
		}
		v.reset(Op386ROLBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		l := v.Args[1]
		if l.Op != Op386MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		if !(psess.canMergeLoad(v, l, x) && clobber(l)) {
			break
		}
		v.reset(Op386XORLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[1]
		l := v.Args[0]
		if l.Op != Op386MOVLload {
			break
		}
		off := l.AuxInt
		sym := l.Aux
		_ = l.Args[1]
		ptr := l.Args[0]
		mem := l.Args[1]
		x := v.Args[1]
		if !(psess.canMergeLoad(v, l, x) && clobber(l)) {
			break
		}
		v.reset(Op386XORLload)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(x)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386XORL_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValue386_Op386XORLconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != Op386XORLconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(Op386XORLconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}

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
		if v_0.Op != Op386MOVLconst {
			break
		}
		d := v_0.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = c ^ d
		return true
	}
	return false
}
func rewriteValue386_Op386XORLload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386ADDLconst {
			break
		}
		off2 := v_1.AuxInt
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386XORLload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		_ = v.Args[2]
		val := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != Op386LEAL {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		base := v_1.Args[0]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386XORLload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(val)
		v.AddArg(base)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_Op386XORLmodify_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != Op386ADDLconst {
			break
		}
		off2 := v_0.AuxInt
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(Op386XORLmodify)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_0.Op != Op386LEAL {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		base := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(is32Bit(off1+off2) && canMergeSym(sym1, sym2) && (base.Op != OpSB || !config.ctxt.Flag_shared)) {
			break
		}
		v.reset(Op386XORLmodify)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(base)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_OpAdd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ADDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpAdd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ADDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpAdd32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ADDSS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpAdd32carry_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ADDLcarry)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpAdd32withcarry_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		c := v.Args[2]
		v.reset(Op386ADCL)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(c)
		return true
	}
}
func rewriteValue386_OpAdd64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ADDSD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpAdd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ADDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpAddPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ADDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpAddr_0(v *Value) bool {

	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(Op386LEAL)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValue386_OpAnd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpAnd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpAnd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpAndB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpAvg32u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386AVGLU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpBswap32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386BSWAPL)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpClosureCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386CALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValue386_OpCom16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386NOTL)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpCom32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386NOTL)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpCom8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386NOTL)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpConst16_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValue386_OpConst32_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValue386_OpConst32F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(Op386MOVSSconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValue386_OpConst64F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(Op386MOVSDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValue386_OpConst8_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValue386_OpConstBool_0(v *Value) bool {

	for {
		b := v.AuxInt
		v.reset(Op386MOVLconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValue386_OpConstNil_0(v *Value) bool {

	for {
		v.reset(Op386MOVLconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValue386_OpCvt32Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386CVTTSS2SL)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpCvt32Fto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386CVTSS2SD)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpCvt32to32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386CVTSL2SS)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpCvt32to64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386CVTSL2SD)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpCvt64Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386CVTTSD2SL)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpCvt64Fto32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386CVTSD2SS)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpDiv16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386DIVW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpDiv16u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386DIVWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpDiv32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386DIVL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpDiv32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386DIVSS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpDiv32u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386DIVLU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpDiv64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386DIVSD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpDiv8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386DIVW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to16, typ.Int16)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to16, typ.Int16)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValue386_OpDiv8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386DIVWU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to16, typ.UInt16)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to16, typ.UInt16)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpEq16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETEQ)
		v0 := b.NewValue0(v.Pos, Op386CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpEq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETEQ)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpEq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETEQF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISS, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpEq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETEQF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpEq8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETEQ)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpEqB_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETEQ)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpEqPtr_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETEQ)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGeq16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETGE)
		v0 := b.NewValue0(v.Pos, Op386CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGeq16U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETAE)
		v0 := b.NewValue0(v.Pos, Op386CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETGE)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETGEF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISS, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGeq32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETAE)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETGEF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGeq8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETGE)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGeq8U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETAE)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue386_OpGetCallerPC_0(v *Value) bool {

	for {
		v.reset(Op386LoweredGetCallerPC)
		return true
	}
}
func rewriteValue386_OpGetCallerSP_0(v *Value) bool {

	for {
		v.reset(Op386LoweredGetCallerSP)
		return true
	}
}
func rewriteValue386_OpGetClosurePtr_0(v *Value) bool {

	for {
		v.reset(Op386LoweredGetClosurePtr)
		return true
	}
}
func rewriteValue386_OpGetG_0(v *Value) bool {

	for {
		mem := v.Args[0]
		v.reset(Op386LoweredGetG)
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGreater16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETG)
		v0 := b.NewValue0(v.Pos, Op386CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGreater16U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETA)
		v0 := b.NewValue0(v.Pos, Op386CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGreater32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETG)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGreater32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETGF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISS, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGreater32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETA)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGreater64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETGF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGreater8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETG)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpGreater8U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETA)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue386_OpHmul32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386HMULL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpHmul32u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386HMULLU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpInterCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[1]
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(Op386CALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpIsInBounds_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(Op386SETB)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		p := v.Args[0]
		v.reset(Op386SETNE)
		v0 := b.NewValue0(v.Pos, Op386TESTL, psess.types.TypeFlags)
		v0.AddArg(p)
		v0.AddArg(p)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpIsSliceInBounds_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(Op386SETBE)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLeq16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETLE)
		v0 := b.NewValue0(v.Pos, Op386CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLeq16U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETBE)
		v0 := b.NewValue0(v.Pos, Op386CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETLE)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETGEF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISS, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLeq32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETBE)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETGEF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISD, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLeq8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETLE)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLeq8U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETBE)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLess16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETL)
		v0 := b.NewValue0(v.Pos, Op386CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLess16U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETB)
		v0 := b.NewValue0(v.Pos, Op386CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLess32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETL)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLess32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETGF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISS, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLess32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETB)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLess64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETGF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISD, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLess8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETL)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLess8U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETB)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLoad_0(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is32BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(Op386MOVLload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is16BitInt(t)) {
			break
		}
		v.reset(Op386MOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsBoolean() || psess.is8BitInt(t)) {
			break
		}
		v.reset(Op386MOVBload)
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
		v.reset(Op386MOVSSload)
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
		v.reset(Op386MOVSDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPWconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPLconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValue386_OpLsh16x64_0(v *Value) bool {

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
		v.reset(Op386SHLLconst)
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
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPBconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPWconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPLconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValue386_OpLsh32x64_0(v *Value) bool {

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
		v.reset(Op386SHLLconst)
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
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPBconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPWconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPLconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValue386_OpLsh8x64_0(v *Value) bool {

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
		v.reset(Op386SHLLconst)
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
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHLL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPBconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValue386_OpMod16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MODW)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpMod16u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MODWU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpMod32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MODL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpMod32u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MODLU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpMod8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MODW)
		v0 := b.NewValue0(v.Pos, OpSignExt8to16, typ.Int16)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to16, typ.Int16)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValue386_OpMod8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MODWU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to16, typ.UInt16)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to16, typ.UInt16)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpMove_0(v *Value) bool {
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
		v.reset(Op386MOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, Op386MOVBload, typ.UInt8)
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
		v.reset(Op386MOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, Op386MOVWload, typ.UInt16)
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
		v.reset(Op386MOVLstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
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
		v.reset(Op386MOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, Op386MOVBload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386MOVWstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, Op386MOVWload, typ.UInt16)
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
		v.reset(Op386MOVBstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, Op386MOVBload, typ.UInt8)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386MOVLstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
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
		v.reset(Op386MOVWstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, Op386MOVWload, typ.UInt16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386MOVLstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
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
		v.reset(Op386MOVLstore)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386MOVLstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
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
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386MOVLstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386MOVLstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
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
		if !(s > 8 && s%4 != 0) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = s - s%4
		v0 := b.NewValue0(v.Pos, Op386ADDLconst, dst.Type)
		v0.AuxInt = s % 4
		v0.AddArg(dst)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386ADDLconst, src.Type)
		v1.AuxInt = s % 4
		v1.AddArg(src)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, Op386MOVLstore, psess.types.TypeMem)
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, Op386MOVLload, typ.UInt32)
		v3.AddArg(src)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v2.AddArg(mem)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValue386_OpMove_10(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		s := v.AuxInt
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 8 && s <= 4*128 && s%4 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(Op386DUFFCOPY)
		v.AuxInt = 10 * (128 - s/4)
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
		if !((s > 4*128 || config.noDuffDevice) && s%4 == 0) {
			break
		}
		v.reset(Op386REPMOVSL)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, Op386MOVLconst, typ.UInt32)
		v0.AuxInt = s / 4
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_OpMul16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MULL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpMul32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MULL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpMul32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MULSS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpMul32uhilo_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MULLQU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpMul64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MULSD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpMul8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386MULL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpNeg16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386NEGL)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpNeg32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386NEGL)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpNeg32F_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if !(!config.use387) {
			break
		}
		v.reset(Op386PXOR)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386MOVSSconst, typ.Float32)
		v0.AuxInt = f2i(math.Copysign(0, -1))
		v.AddArg(v0)
		return true
	}

	for {
		x := v.Args[0]
		if !(config.use387) {
			break
		}
		v.reset(Op386FCHS)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_OpNeg64F_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if !(!config.use387) {
			break
		}
		v.reset(Op386PXOR)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386MOVSDconst, typ.Float64)
		v0.AuxInt = f2i(math.Copysign(0, -1))
		v.AddArg(v0)
		return true
	}

	for {
		x := v.Args[0]
		if !(config.use387) {
			break
		}
		v.reset(Op386FCHS)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue386_OpNeg8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386NEGL)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpNeq16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETNE)
		v0 := b.NewValue0(v.Pos, Op386CMPW, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpNeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETNE)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpNeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETNEF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISS, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpNeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETNEF)
		v0 := b.NewValue0(v.Pos, Op386UCOMISD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpNeq8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETNE)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpNeqB_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETNE)
		v0 := b.NewValue0(v.Pos, Op386CMPB, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpNeqPtr_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SETNE)
		v0 := b.NewValue0(v.Pos, Op386CMPL, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue386_OpNilCheck_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(Op386LoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValue386_OpNot_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386XORLconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpOffPtr_0(v *Value) bool {

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(Op386ADDLconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValue386_OpOr16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpOr32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpOr8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpOrB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpRound32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpRound64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPWconst, psess.types.TypeFlags)
		v2.AuxInt = 16
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPLconst, psess.types.TypeFlags)
		v2.AuxInt = 16
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValue386_OpRsh16Ux64_0(v *Value) bool {

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
		v.reset(Op386SHRWconst)
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
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHRW, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPBconst, psess.types.TypeFlags)
		v2.AuxInt = 16
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SARW)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, Op386NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, Op386SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, Op386CMPWconst, psess.types.TypeFlags)
		v3.AuxInt = 16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SARW)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, Op386NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, Op386SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, Op386CMPLconst, psess.types.TypeFlags)
		v3.AuxInt = 16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue386_OpRsh16x64_0(v *Value) bool {

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
		v.reset(Op386SARWconst)
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
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(Op386SARWconst)
		v.AuxInt = 15
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SARW)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, Op386NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, Op386SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, Op386CMPBconst, psess.types.TypeFlags)
		v3.AuxInt = 16
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPWconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPLconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValue386_OpRsh32Ux64_0(v *Value) bool {

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
		v.reset(Op386SHRLconst)
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
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHRL, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPBconst, psess.types.TypeFlags)
		v2.AuxInt = 32
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SARL)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, Op386NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, Op386SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, Op386CMPWconst, psess.types.TypeFlags)
		v3.AuxInt = 32
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SARL)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, Op386NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, Op386SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, Op386CMPLconst, psess.types.TypeFlags)
		v3.AuxInt = 32
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue386_OpRsh32x64_0(v *Value) bool {

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
		v.reset(Op386SARLconst)
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
		v.reset(Op386SARLconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SARL)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, Op386NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, Op386SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, Op386CMPBconst, psess.types.TypeFlags)
		v3.AuxInt = 32
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHRB, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPWconst, psess.types.TypeFlags)
		v2.AuxInt = 8
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHRB, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPLconst, psess.types.TypeFlags)
		v2.AuxInt = 8
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValue386_OpRsh8Ux64_0(v *Value) bool {

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
		v.reset(Op386SHRBconst)
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
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386ANDL)
		v0 := b.NewValue0(v.Pos, Op386SHRB, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v2 := b.NewValue0(v.Pos, Op386CMPBconst, psess.types.TypeFlags)
		v2.AuxInt = 8
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SARB)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, Op386NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, Op386SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, Op386CMPWconst, psess.types.TypeFlags)
		v3.AuxInt = 8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SARB)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, Op386NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, Op386SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, Op386CMPLconst, psess.types.TypeFlags)
		v3.AuxInt = 8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue386_OpRsh8x64_0(v *Value) bool {

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
		v.reset(Op386SARBconst)
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
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(Op386SARBconst)
		v.AuxInt = 7
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SARB)
		v.Type = t
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, Op386ORL, y.Type)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, Op386NOTL, y.Type)
		v2 := b.NewValue0(v.Pos, Op386SBBLcarrymask, y.Type)
		v3 := b.NewValue0(v.Pos, Op386CMPBconst, psess.types.TypeFlags)
		v3.AuxInt = 8
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue386_OpSignExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386MOVWLSX)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpSignExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386MOVBLSX)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpSignExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386MOVBLSX)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpSignmask_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386SARLconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpSlicemask_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(Op386SARLconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, Op386NEGL, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue386_OpSqrt_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386SQRTSD)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpStaticCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(Op386CALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpStore_0(v *Value) bool {

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 8 && psess.is64BitFloat(val.Type)) {
			break
		}
		v.reset(Op386MOVSDstore)
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
		v.reset(Op386MOVSSstore)
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
		v.reset(Op386MOVLstore)
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
		v.reset(Op386MOVWstore)
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
		v.reset(Op386MOVBstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_OpSub16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SUBL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpSub32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SUBL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpSub32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SUBSS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpSub32carry_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SUBLcarry)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpSub32withcarry_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		c := v.Args[2]
		v.reset(Op386SBBL)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(c)
		return true
	}
}
func rewriteValue386_OpSub64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SUBSD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpSub8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SUBL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpSubPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386SUBL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpTrunc16to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpTrunc32to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpTrunc32to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpWB_0(v *Value) bool {

	for {
		fn := v.Aux
		_ = v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		mem := v.Args[2]
		v.reset(Op386LoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValue386_OpXor16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386XORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpXor32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386XORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValue386_OpXor8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(Op386XORL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpZero_0(v *Value) bool {
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
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(Op386MOVBstoreconst)
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
		v.reset(Op386MOVWstoreconst)
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
		v.reset(Op386MOVLstoreconst)
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
		v.reset(Op386MOVBstoreconst)
		v.AuxInt = makeValAndOff(0, 2)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, Op386MOVWstoreconst, psess.types.TypeMem)
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
		v.reset(Op386MOVBstoreconst)
		v.AuxInt = makeValAndOff(0, 4)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, Op386MOVLstoreconst, psess.types.TypeMem)
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
		v.reset(Op386MOVWstoreconst)
		v.AuxInt = makeValAndOff(0, 4)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, Op386MOVLstoreconst, psess.types.TypeMem)
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
		v.reset(Op386MOVLstoreconst)
		v.AuxInt = makeValAndOff(0, 3)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, Op386MOVLstoreconst, psess.types.TypeMem)
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
		if !(s%4 != 0 && s > 4) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = s - s%4
		v0 := b.NewValue0(v.Pos, Op386ADDLconst, typ.UInt32)
		v0.AuxInt = s % 4
		v0.AddArg(destptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386MOVLstoreconst, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(destptr)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 8 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(Op386MOVLstoreconst)
		v.AuxInt = makeValAndOff(0, 4)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, Op386MOVLstoreconst, psess.types.TypeMem)
		v0.AuxInt = 0
		v0.AddArg(destptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValue386_OpZero_10(v *Value) bool {
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
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(Op386MOVLstoreconst)
		v.AuxInt = makeValAndOff(0, 8)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, Op386MOVLstoreconst, psess.types.TypeMem)
		v0.AuxInt = makeValAndOff(0, 4)
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, Op386MOVLstoreconst, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(destptr)
		v1.AddArg(mem)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		if v.AuxInt != 16 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(Op386MOVLstoreconst)
		v.AuxInt = makeValAndOff(0, 12)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, Op386MOVLstoreconst, psess.types.TypeMem)
		v0.AuxInt = makeValAndOff(0, 8)
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, Op386MOVLstoreconst, psess.types.TypeMem)
		v1.AuxInt = makeValAndOff(0, 4)
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, Op386MOVLstoreconst, psess.types.TypeMem)
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
		destptr := v.Args[0]
		mem := v.Args[1]
		if !(s > 16 && s <= 4*128 && s%4 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(Op386DUFFZERO)
		v.AuxInt = 1 * (128 - s/4)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, Op386MOVLconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		if !((s > 4*128 || (config.noDuffDevice && s > 16)) && s%4 == 0) {
			break
		}
		v.reset(Op386REPSTOSL)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, Op386MOVLconst, typ.UInt32)
		v0.AuxInt = s / 4
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, Op386MOVLconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue386_OpZeroExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386MOVWLZX)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpZeroExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386MOVBLZX)
		v.AddArg(x)
		return true
	}
}
func rewriteValue386_OpZeroExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(Op386MOVBLZX)
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValue386_OpZeromask_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(Op386XORLconst)
		v.AuxInt = -1
		v0 := b.NewValue0(v.Pos, Op386SBBLcarrymask, t)
		v1 := b.NewValue0(v.Pos, Op386CMPLconst, psess.types.TypeFlags)
		v1.AuxInt = 1
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteBlock386(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	case Block386EQ:

		for {
			v := b.Control
			if v.Op != Op386InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386EQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagLT_ULT {
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
			if v.Op != Op386FlagLT_UGT {
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
			if v.Op != Op386FlagGT_ULT {
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
			if v.Op != Op386FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case Block386GE:

		for {
			v := b.Control
			if v.Op != Op386InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386LE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagLT_ULT {
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
			if v.Op != Op386FlagLT_UGT {
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
			if v.Op != Op386FlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
	case Block386GT:

		for {
			v := b.Control
			if v.Op != Op386InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386LT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagEQ {
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
			if v.Op != Op386FlagLT_ULT {
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
			if v.Op != Op386FlagLT_UGT {
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
			if v.Op != Op386FlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagGT_UGT {
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
			if v.Op != Op386SETL {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386LT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETLE {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386LE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETG {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386GT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETGE {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386GE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETEQ {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386EQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETNE {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386NE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETB {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386ULT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETBE {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386ULE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETA {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386UGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETAE {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386UGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETGF {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386UGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETGEF {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386UGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETEQF {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386EQF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386SETNEF {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386NEF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			_ = v
			cond := b.Control
			b.Kind = Block386NE
			v0 := b.NewValue0(v.Pos, Op386TESTB, psess.types.TypeFlags)
			v0.AddArg(cond)
			v0.AddArg(cond)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case Block386LE:

		for {
			v := b.Control
			if v.Op != Op386InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386GE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagGT_ULT {
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
			if v.Op != Op386FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case Block386LT:

		for {
			v := b.Control
			if v.Op != Op386InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386GT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagEQ {
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
			if v.Op != Op386FlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagGT_ULT {
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
			if v.Op != Op386FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case Block386NE:

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETL {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETL {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386LT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETL {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETL {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386LT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETLE {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETLE {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386LE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETLE {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETLE {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386LE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETG {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETG {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386GT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETG {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETG {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386GT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETGE {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETGE {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386GE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETGE {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETGE {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386GE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETEQ {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETEQ {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386EQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETEQ {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETEQ {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386EQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETNE {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETNE {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386NE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETNE {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETNE {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386NE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETB {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETB {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386ULT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETB {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETB {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386ULT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETBE {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETBE {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386ULE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETBE {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETBE {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386ULE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETA {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETA {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386UGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETA {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETA {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386UGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETAE {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETAE {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386UGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETAE {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETAE {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386UGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETGF {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETGF {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386UGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETGF {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETGF {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386UGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETGEF {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETGEF {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386UGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETGEF {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETGEF {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386UGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETEQF {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETEQF {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386EQF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETEQF {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETEQF {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386EQF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETNEF {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETNEF {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386NEF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386TESTB {
				break
			}
			_ = v.Args[1]
			v_0 := v.Args[0]
			if v_0.Op != Op386SETNEF {
				break
			}
			cmp := v_0.Args[0]
			v_1 := v.Args[1]
			if v_1.Op != Op386SETNEF {
				break
			}
			if cmp != v_1.Args[0] {
				break
			}
			b.Kind = Block386NEF
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386NE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagEQ {
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
			if v.Op != Op386FlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
	case Block386UGE:

		for {
			v := b.Control
			if v.Op != Op386InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386ULE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagLT_ULT {
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
			if v.Op != Op386FlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagGT_ULT {
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
			if v.Op != Op386FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
	case Block386UGT:

		for {
			v := b.Control
			if v.Op != Op386InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386ULT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagEQ {
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
			if v.Op != Op386FlagLT_ULT {
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
			if v.Op != Op386FlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagGT_ULT {
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
			if v.Op != Op386FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}
	case Block386ULE:

		for {
			v := b.Control
			if v.Op != Op386InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386UGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagLT_UGT {
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
			if v.Op != Op386FlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			b.swapSuccessors()
			return true
		}
	case Block386ULT:

		for {
			v := b.Control
			if v.Op != Op386InvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = Block386UGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagEQ {
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
			if v.Op != Op386FlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagLT_UGT {
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
			if v.Op != Op386FlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != Op386FlagGT_UGT {
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

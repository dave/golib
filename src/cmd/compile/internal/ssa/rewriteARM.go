package ssa

import "github.com/dave/golib/src/cmd/compile/internal/types"

// in case not otherwise used
// in case not otherwise used
// in case not otherwise used
// in case not otherwise used

func (psess *PackageSession) rewriteValueARM(v *Value) bool {
	switch v.Op {
	case OpARMADC:
		return rewriteValueARM_OpARMADC_0(v) || rewriteValueARM_OpARMADC_10(v) || rewriteValueARM_OpARMADC_20(v)
	case OpARMADCconst:
		return rewriteValueARM_OpARMADCconst_0(v)
	case OpARMADCshiftLL:
		return rewriteValueARM_OpARMADCshiftLL_0(v)
	case OpARMADCshiftLLreg:
		return rewriteValueARM_OpARMADCshiftLLreg_0(v)
	case OpARMADCshiftRA:
		return rewriteValueARM_OpARMADCshiftRA_0(v)
	case OpARMADCshiftRAreg:
		return rewriteValueARM_OpARMADCshiftRAreg_0(v)
	case OpARMADCshiftRL:
		return rewriteValueARM_OpARMADCshiftRL_0(v)
	case OpARMADCshiftRLreg:
		return rewriteValueARM_OpARMADCshiftRLreg_0(v)
	case OpARMADD:
		return rewriteValueARM_OpARMADD_0(v) || rewriteValueARM_OpARMADD_10(v)
	case OpARMADDD:
		return psess.rewriteValueARM_OpARMADDD_0(v)
	case OpARMADDF:
		return psess.rewriteValueARM_OpARMADDF_0(v)
	case OpARMADDS:
		return rewriteValueARM_OpARMADDS_0(v) || rewriteValueARM_OpARMADDS_10(v)
	case OpARMADDSshiftLL:
		return rewriteValueARM_OpARMADDSshiftLL_0(v)
	case OpARMADDSshiftLLreg:
		return rewriteValueARM_OpARMADDSshiftLLreg_0(v)
	case OpARMADDSshiftRA:
		return rewriteValueARM_OpARMADDSshiftRA_0(v)
	case OpARMADDSshiftRAreg:
		return rewriteValueARM_OpARMADDSshiftRAreg_0(v)
	case OpARMADDSshiftRL:
		return rewriteValueARM_OpARMADDSshiftRL_0(v)
	case OpARMADDSshiftRLreg:
		return rewriteValueARM_OpARMADDSshiftRLreg_0(v)
	case OpARMADDconst:
		return rewriteValueARM_OpARMADDconst_0(v)
	case OpARMADDshiftLL:
		return rewriteValueARM_OpARMADDshiftLL_0(v)
	case OpARMADDshiftLLreg:
		return rewriteValueARM_OpARMADDshiftLLreg_0(v)
	case OpARMADDshiftRA:
		return rewriteValueARM_OpARMADDshiftRA_0(v)
	case OpARMADDshiftRAreg:
		return rewriteValueARM_OpARMADDshiftRAreg_0(v)
	case OpARMADDshiftRL:
		return rewriteValueARM_OpARMADDshiftRL_0(v)
	case OpARMADDshiftRLreg:
		return rewriteValueARM_OpARMADDshiftRLreg_0(v)
	case OpARMAND:
		return rewriteValueARM_OpARMAND_0(v) || rewriteValueARM_OpARMAND_10(v) || rewriteValueARM_OpARMAND_20(v)
	case OpARMANDconst:
		return rewriteValueARM_OpARMANDconst_0(v)
	case OpARMANDshiftLL:
		return rewriteValueARM_OpARMANDshiftLL_0(v)
	case OpARMANDshiftLLreg:
		return rewriteValueARM_OpARMANDshiftLLreg_0(v)
	case OpARMANDshiftRA:
		return rewriteValueARM_OpARMANDshiftRA_0(v)
	case OpARMANDshiftRAreg:
		return rewriteValueARM_OpARMANDshiftRAreg_0(v)
	case OpARMANDshiftRL:
		return rewriteValueARM_OpARMANDshiftRL_0(v)
	case OpARMANDshiftRLreg:
		return rewriteValueARM_OpARMANDshiftRLreg_0(v)
	case OpARMBFX:
		return rewriteValueARM_OpARMBFX_0(v)
	case OpARMBFXU:
		return rewriteValueARM_OpARMBFXU_0(v)
	case OpARMBIC:
		return rewriteValueARM_OpARMBIC_0(v)
	case OpARMBICconst:
		return rewriteValueARM_OpARMBICconst_0(v)
	case OpARMBICshiftLL:
		return rewriteValueARM_OpARMBICshiftLL_0(v)
	case OpARMBICshiftLLreg:
		return rewriteValueARM_OpARMBICshiftLLreg_0(v)
	case OpARMBICshiftRA:
		return rewriteValueARM_OpARMBICshiftRA_0(v)
	case OpARMBICshiftRAreg:
		return rewriteValueARM_OpARMBICshiftRAreg_0(v)
	case OpARMBICshiftRL:
		return rewriteValueARM_OpARMBICshiftRL_0(v)
	case OpARMBICshiftRLreg:
		return rewriteValueARM_OpARMBICshiftRLreg_0(v)
	case OpARMCMN:
		return rewriteValueARM_OpARMCMN_0(v) || rewriteValueARM_OpARMCMN_10(v)
	case OpARMCMNconst:
		return rewriteValueARM_OpARMCMNconst_0(v)
	case OpARMCMNshiftLL:
		return rewriteValueARM_OpARMCMNshiftLL_0(v)
	case OpARMCMNshiftLLreg:
		return rewriteValueARM_OpARMCMNshiftLLreg_0(v)
	case OpARMCMNshiftRA:
		return rewriteValueARM_OpARMCMNshiftRA_0(v)
	case OpARMCMNshiftRAreg:
		return rewriteValueARM_OpARMCMNshiftRAreg_0(v)
	case OpARMCMNshiftRL:
		return rewriteValueARM_OpARMCMNshiftRL_0(v)
	case OpARMCMNshiftRLreg:
		return rewriteValueARM_OpARMCMNshiftRLreg_0(v)
	case OpARMCMOVWHSconst:
		return rewriteValueARM_OpARMCMOVWHSconst_0(v)
	case OpARMCMOVWLSconst:
		return rewriteValueARM_OpARMCMOVWLSconst_0(v)
	case OpARMCMP:
		return psess.rewriteValueARM_OpARMCMP_0(v) || psess.rewriteValueARM_OpARMCMP_10(v)
	case OpARMCMPD:
		return rewriteValueARM_OpARMCMPD_0(v)
	case OpARMCMPF:
		return rewriteValueARM_OpARMCMPF_0(v)
	case OpARMCMPconst:
		return rewriteValueARM_OpARMCMPconst_0(v)
	case OpARMCMPshiftLL:
		return psess.rewriteValueARM_OpARMCMPshiftLL_0(v)
	case OpARMCMPshiftLLreg:
		return psess.rewriteValueARM_OpARMCMPshiftLLreg_0(v)
	case OpARMCMPshiftRA:
		return psess.rewriteValueARM_OpARMCMPshiftRA_0(v)
	case OpARMCMPshiftRAreg:
		return psess.rewriteValueARM_OpARMCMPshiftRAreg_0(v)
	case OpARMCMPshiftRL:
		return psess.rewriteValueARM_OpARMCMPshiftRL_0(v)
	case OpARMCMPshiftRLreg:
		return psess.rewriteValueARM_OpARMCMPshiftRLreg_0(v)
	case OpARMEqual:
		return rewriteValueARM_OpARMEqual_0(v)
	case OpARMGreaterEqual:
		return rewriteValueARM_OpARMGreaterEqual_0(v)
	case OpARMGreaterEqualU:
		return rewriteValueARM_OpARMGreaterEqualU_0(v)
	case OpARMGreaterThan:
		return rewriteValueARM_OpARMGreaterThan_0(v)
	case OpARMGreaterThanU:
		return rewriteValueARM_OpARMGreaterThanU_0(v)
	case OpARMLessEqual:
		return rewriteValueARM_OpARMLessEqual_0(v)
	case OpARMLessEqualU:
		return rewriteValueARM_OpARMLessEqualU_0(v)
	case OpARMLessThan:
		return rewriteValueARM_OpARMLessThan_0(v)
	case OpARMLessThanU:
		return rewriteValueARM_OpARMLessThanU_0(v)
	case OpARMMOVBUload:
		return rewriteValueARM_OpARMMOVBUload_0(v)
	case OpARMMOVBUloadidx:
		return rewriteValueARM_OpARMMOVBUloadidx_0(v)
	case OpARMMOVBUreg:
		return rewriteValueARM_OpARMMOVBUreg_0(v)
	case OpARMMOVBload:
		return rewriteValueARM_OpARMMOVBload_0(v)
	case OpARMMOVBloadidx:
		return rewriteValueARM_OpARMMOVBloadidx_0(v)
	case OpARMMOVBreg:
		return rewriteValueARM_OpARMMOVBreg_0(v)
	case OpARMMOVBstore:
		return rewriteValueARM_OpARMMOVBstore_0(v)
	case OpARMMOVBstoreidx:
		return rewriteValueARM_OpARMMOVBstoreidx_0(v)
	case OpARMMOVDload:
		return rewriteValueARM_OpARMMOVDload_0(v)
	case OpARMMOVDstore:
		return rewriteValueARM_OpARMMOVDstore_0(v)
	case OpARMMOVFload:
		return rewriteValueARM_OpARMMOVFload_0(v)
	case OpARMMOVFstore:
		return rewriteValueARM_OpARMMOVFstore_0(v)
	case OpARMMOVHUload:
		return rewriteValueARM_OpARMMOVHUload_0(v)
	case OpARMMOVHUloadidx:
		return rewriteValueARM_OpARMMOVHUloadidx_0(v)
	case OpARMMOVHUreg:
		return rewriteValueARM_OpARMMOVHUreg_0(v)
	case OpARMMOVHload:
		return rewriteValueARM_OpARMMOVHload_0(v)
	case OpARMMOVHloadidx:
		return rewriteValueARM_OpARMMOVHloadidx_0(v)
	case OpARMMOVHreg:
		return rewriteValueARM_OpARMMOVHreg_0(v)
	case OpARMMOVHstore:
		return rewriteValueARM_OpARMMOVHstore_0(v)
	case OpARMMOVHstoreidx:
		return rewriteValueARM_OpARMMOVHstoreidx_0(v)
	case OpARMMOVWload:
		return rewriteValueARM_OpARMMOVWload_0(v)
	case OpARMMOVWloadidx:
		return rewriteValueARM_OpARMMOVWloadidx_0(v)
	case OpARMMOVWloadshiftLL:
		return rewriteValueARM_OpARMMOVWloadshiftLL_0(v)
	case OpARMMOVWloadshiftRA:
		return rewriteValueARM_OpARMMOVWloadshiftRA_0(v)
	case OpARMMOVWloadshiftRL:
		return rewriteValueARM_OpARMMOVWloadshiftRL_0(v)
	case OpARMMOVWreg:
		return rewriteValueARM_OpARMMOVWreg_0(v)
	case OpARMMOVWstore:
		return rewriteValueARM_OpARMMOVWstore_0(v)
	case OpARMMOVWstoreidx:
		return rewriteValueARM_OpARMMOVWstoreidx_0(v)
	case OpARMMOVWstoreshiftLL:
		return rewriteValueARM_OpARMMOVWstoreshiftLL_0(v)
	case OpARMMOVWstoreshiftRA:
		return rewriteValueARM_OpARMMOVWstoreshiftRA_0(v)
	case OpARMMOVWstoreshiftRL:
		return rewriteValueARM_OpARMMOVWstoreshiftRL_0(v)
	case OpARMMUL:
		return rewriteValueARM_OpARMMUL_0(v) || rewriteValueARM_OpARMMUL_10(v) || rewriteValueARM_OpARMMUL_20(v)
	case OpARMMULA:
		return rewriteValueARM_OpARMMULA_0(v) || rewriteValueARM_OpARMMULA_10(v) || rewriteValueARM_OpARMMULA_20(v)
	case OpARMMULD:
		return psess.rewriteValueARM_OpARMMULD_0(v)
	case OpARMMULF:
		return psess.rewriteValueARM_OpARMMULF_0(v)
	case OpARMMULS:
		return rewriteValueARM_OpARMMULS_0(v) || rewriteValueARM_OpARMMULS_10(v) || rewriteValueARM_OpARMMULS_20(v)
	case OpARMMVN:
		return rewriteValueARM_OpARMMVN_0(v)
	case OpARMMVNshiftLL:
		return rewriteValueARM_OpARMMVNshiftLL_0(v)
	case OpARMMVNshiftLLreg:
		return rewriteValueARM_OpARMMVNshiftLLreg_0(v)
	case OpARMMVNshiftRA:
		return rewriteValueARM_OpARMMVNshiftRA_0(v)
	case OpARMMVNshiftRAreg:
		return rewriteValueARM_OpARMMVNshiftRAreg_0(v)
	case OpARMMVNshiftRL:
		return rewriteValueARM_OpARMMVNshiftRL_0(v)
	case OpARMMVNshiftRLreg:
		return rewriteValueARM_OpARMMVNshiftRLreg_0(v)
	case OpARMNEGD:
		return psess.rewriteValueARM_OpARMNEGD_0(v)
	case OpARMNEGF:
		return psess.rewriteValueARM_OpARMNEGF_0(v)
	case OpARMNMULD:
		return rewriteValueARM_OpARMNMULD_0(v)
	case OpARMNMULF:
		return rewriteValueARM_OpARMNMULF_0(v)
	case OpARMNotEqual:
		return rewriteValueARM_OpARMNotEqual_0(v)
	case OpARMOR:
		return rewriteValueARM_OpARMOR_0(v) || rewriteValueARM_OpARMOR_10(v)
	case OpARMORconst:
		return rewriteValueARM_OpARMORconst_0(v)
	case OpARMORshiftLL:
		return rewriteValueARM_OpARMORshiftLL_0(v)
	case OpARMORshiftLLreg:
		return rewriteValueARM_OpARMORshiftLLreg_0(v)
	case OpARMORshiftRA:
		return rewriteValueARM_OpARMORshiftRA_0(v)
	case OpARMORshiftRAreg:
		return rewriteValueARM_OpARMORshiftRAreg_0(v)
	case OpARMORshiftRL:
		return rewriteValueARM_OpARMORshiftRL_0(v)
	case OpARMORshiftRLreg:
		return rewriteValueARM_OpARMORshiftRLreg_0(v)
	case OpARMRSB:
		return rewriteValueARM_OpARMRSB_0(v) || psess.rewriteValueARM_OpARMRSB_10(v)
	case OpARMRSBSshiftLL:
		return rewriteValueARM_OpARMRSBSshiftLL_0(v)
	case OpARMRSBSshiftLLreg:
		return rewriteValueARM_OpARMRSBSshiftLLreg_0(v)
	case OpARMRSBSshiftRA:
		return rewriteValueARM_OpARMRSBSshiftRA_0(v)
	case OpARMRSBSshiftRAreg:
		return rewriteValueARM_OpARMRSBSshiftRAreg_0(v)
	case OpARMRSBSshiftRL:
		return rewriteValueARM_OpARMRSBSshiftRL_0(v)
	case OpARMRSBSshiftRLreg:
		return rewriteValueARM_OpARMRSBSshiftRLreg_0(v)
	case OpARMRSBconst:
		return rewriteValueARM_OpARMRSBconst_0(v)
	case OpARMRSBshiftLL:
		return rewriteValueARM_OpARMRSBshiftLL_0(v)
	case OpARMRSBshiftLLreg:
		return rewriteValueARM_OpARMRSBshiftLLreg_0(v)
	case OpARMRSBshiftRA:
		return rewriteValueARM_OpARMRSBshiftRA_0(v)
	case OpARMRSBshiftRAreg:
		return rewriteValueARM_OpARMRSBshiftRAreg_0(v)
	case OpARMRSBshiftRL:
		return rewriteValueARM_OpARMRSBshiftRL_0(v)
	case OpARMRSBshiftRLreg:
		return rewriteValueARM_OpARMRSBshiftRLreg_0(v)
	case OpARMRSCconst:
		return rewriteValueARM_OpARMRSCconst_0(v)
	case OpARMRSCshiftLL:
		return rewriteValueARM_OpARMRSCshiftLL_0(v)
	case OpARMRSCshiftLLreg:
		return rewriteValueARM_OpARMRSCshiftLLreg_0(v)
	case OpARMRSCshiftRA:
		return rewriteValueARM_OpARMRSCshiftRA_0(v)
	case OpARMRSCshiftRAreg:
		return rewriteValueARM_OpARMRSCshiftRAreg_0(v)
	case OpARMRSCshiftRL:
		return rewriteValueARM_OpARMRSCshiftRL_0(v)
	case OpARMRSCshiftRLreg:
		return rewriteValueARM_OpARMRSCshiftRLreg_0(v)
	case OpARMSBC:
		return rewriteValueARM_OpARMSBC_0(v) || rewriteValueARM_OpARMSBC_10(v)
	case OpARMSBCconst:
		return rewriteValueARM_OpARMSBCconst_0(v)
	case OpARMSBCshiftLL:
		return rewriteValueARM_OpARMSBCshiftLL_0(v)
	case OpARMSBCshiftLLreg:
		return rewriteValueARM_OpARMSBCshiftLLreg_0(v)
	case OpARMSBCshiftRA:
		return rewriteValueARM_OpARMSBCshiftRA_0(v)
	case OpARMSBCshiftRAreg:
		return rewriteValueARM_OpARMSBCshiftRAreg_0(v)
	case OpARMSBCshiftRL:
		return rewriteValueARM_OpARMSBCshiftRL_0(v)
	case OpARMSBCshiftRLreg:
		return rewriteValueARM_OpARMSBCshiftRLreg_0(v)
	case OpARMSLL:
		return rewriteValueARM_OpARMSLL_0(v)
	case OpARMSLLconst:
		return rewriteValueARM_OpARMSLLconst_0(v)
	case OpARMSRA:
		return rewriteValueARM_OpARMSRA_0(v)
	case OpARMSRAcond:
		return rewriteValueARM_OpARMSRAcond_0(v)
	case OpARMSRAconst:
		return psess.rewriteValueARM_OpARMSRAconst_0(v)
	case OpARMSRL:
		return rewriteValueARM_OpARMSRL_0(v)
	case OpARMSRLconst:
		return psess.rewriteValueARM_OpARMSRLconst_0(v)
	case OpARMSUB:
		return rewriteValueARM_OpARMSUB_0(v) || psess.rewriteValueARM_OpARMSUB_10(v)
	case OpARMSUBD:
		return psess.rewriteValueARM_OpARMSUBD_0(v)
	case OpARMSUBF:
		return psess.rewriteValueARM_OpARMSUBF_0(v)
	case OpARMSUBS:
		return rewriteValueARM_OpARMSUBS_0(v) || rewriteValueARM_OpARMSUBS_10(v)
	case OpARMSUBSshiftLL:
		return rewriteValueARM_OpARMSUBSshiftLL_0(v)
	case OpARMSUBSshiftLLreg:
		return rewriteValueARM_OpARMSUBSshiftLLreg_0(v)
	case OpARMSUBSshiftRA:
		return rewriteValueARM_OpARMSUBSshiftRA_0(v)
	case OpARMSUBSshiftRAreg:
		return rewriteValueARM_OpARMSUBSshiftRAreg_0(v)
	case OpARMSUBSshiftRL:
		return rewriteValueARM_OpARMSUBSshiftRL_0(v)
	case OpARMSUBSshiftRLreg:
		return rewriteValueARM_OpARMSUBSshiftRLreg_0(v)
	case OpARMSUBconst:
		return rewriteValueARM_OpARMSUBconst_0(v)
	case OpARMSUBshiftLL:
		return rewriteValueARM_OpARMSUBshiftLL_0(v)
	case OpARMSUBshiftLLreg:
		return rewriteValueARM_OpARMSUBshiftLLreg_0(v)
	case OpARMSUBshiftRA:
		return rewriteValueARM_OpARMSUBshiftRA_0(v)
	case OpARMSUBshiftRAreg:
		return rewriteValueARM_OpARMSUBshiftRAreg_0(v)
	case OpARMSUBshiftRL:
		return rewriteValueARM_OpARMSUBshiftRL_0(v)
	case OpARMSUBshiftRLreg:
		return rewriteValueARM_OpARMSUBshiftRLreg_0(v)
	case OpARMTEQ:
		return rewriteValueARM_OpARMTEQ_0(v) || rewriteValueARM_OpARMTEQ_10(v)
	case OpARMTEQconst:
		return rewriteValueARM_OpARMTEQconst_0(v)
	case OpARMTEQshiftLL:
		return rewriteValueARM_OpARMTEQshiftLL_0(v)
	case OpARMTEQshiftLLreg:
		return rewriteValueARM_OpARMTEQshiftLLreg_0(v)
	case OpARMTEQshiftRA:
		return rewriteValueARM_OpARMTEQshiftRA_0(v)
	case OpARMTEQshiftRAreg:
		return rewriteValueARM_OpARMTEQshiftRAreg_0(v)
	case OpARMTEQshiftRL:
		return rewriteValueARM_OpARMTEQshiftRL_0(v)
	case OpARMTEQshiftRLreg:
		return rewriteValueARM_OpARMTEQshiftRLreg_0(v)
	case OpARMTST:
		return rewriteValueARM_OpARMTST_0(v) || rewriteValueARM_OpARMTST_10(v)
	case OpARMTSTconst:
		return rewriteValueARM_OpARMTSTconst_0(v)
	case OpARMTSTshiftLL:
		return rewriteValueARM_OpARMTSTshiftLL_0(v)
	case OpARMTSTshiftLLreg:
		return rewriteValueARM_OpARMTSTshiftLLreg_0(v)
	case OpARMTSTshiftRA:
		return rewriteValueARM_OpARMTSTshiftRA_0(v)
	case OpARMTSTshiftRAreg:
		return rewriteValueARM_OpARMTSTshiftRAreg_0(v)
	case OpARMTSTshiftRL:
		return rewriteValueARM_OpARMTSTshiftRL_0(v)
	case OpARMTSTshiftRLreg:
		return rewriteValueARM_OpARMTSTshiftRLreg_0(v)
	case OpARMXOR:
		return rewriteValueARM_OpARMXOR_0(v) || rewriteValueARM_OpARMXOR_10(v)
	case OpARMXORconst:
		return rewriteValueARM_OpARMXORconst_0(v)
	case OpARMXORshiftLL:
		return rewriteValueARM_OpARMXORshiftLL_0(v)
	case OpARMXORshiftLLreg:
		return rewriteValueARM_OpARMXORshiftLLreg_0(v)
	case OpARMXORshiftRA:
		return rewriteValueARM_OpARMXORshiftRA_0(v)
	case OpARMXORshiftRAreg:
		return rewriteValueARM_OpARMXORshiftRAreg_0(v)
	case OpARMXORshiftRL:
		return rewriteValueARM_OpARMXORshiftRL_0(v)
	case OpARMXORshiftRLreg:
		return rewriteValueARM_OpARMXORshiftRLreg_0(v)
	case OpARMXORshiftRR:
		return rewriteValueARM_OpARMXORshiftRR_0(v)
	case OpAdd16:
		return rewriteValueARM_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValueARM_OpAdd32_0(v)
	case OpAdd32F:
		return rewriteValueARM_OpAdd32F_0(v)
	case OpAdd32carry:
		return rewriteValueARM_OpAdd32carry_0(v)
	case OpAdd32withcarry:
		return rewriteValueARM_OpAdd32withcarry_0(v)
	case OpAdd64F:
		return rewriteValueARM_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValueARM_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValueARM_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValueARM_OpAddr_0(v)
	case OpAnd16:
		return rewriteValueARM_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValueARM_OpAnd32_0(v)
	case OpAnd8:
		return rewriteValueARM_OpAnd8_0(v)
	case OpAndB:
		return rewriteValueARM_OpAndB_0(v)
	case OpAvg32u:
		return rewriteValueARM_OpAvg32u_0(v)
	case OpBitLen32:
		return rewriteValueARM_OpBitLen32_0(v)
	case OpBswap32:
		return psess.rewriteValueARM_OpBswap32_0(v)
	case OpClosureCall:
		return rewriteValueARM_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValueARM_OpCom16_0(v)
	case OpCom32:
		return rewriteValueARM_OpCom32_0(v)
	case OpCom8:
		return rewriteValueARM_OpCom8_0(v)
	case OpConst16:
		return rewriteValueARM_OpConst16_0(v)
	case OpConst32:
		return rewriteValueARM_OpConst32_0(v)
	case OpConst32F:
		return rewriteValueARM_OpConst32F_0(v)
	case OpConst64F:
		return rewriteValueARM_OpConst64F_0(v)
	case OpConst8:
		return rewriteValueARM_OpConst8_0(v)
	case OpConstBool:
		return rewriteValueARM_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValueARM_OpConstNil_0(v)
	case OpCtz32:
		return psess.rewriteValueARM_OpCtz32_0(v)
	case OpCtz32NonZero:
		return rewriteValueARM_OpCtz32NonZero_0(v)
	case OpCvt32Fto32:
		return rewriteValueARM_OpCvt32Fto32_0(v)
	case OpCvt32Fto32U:
		return rewriteValueARM_OpCvt32Fto32U_0(v)
	case OpCvt32Fto64F:
		return rewriteValueARM_OpCvt32Fto64F_0(v)
	case OpCvt32Uto32F:
		return rewriteValueARM_OpCvt32Uto32F_0(v)
	case OpCvt32Uto64F:
		return rewriteValueARM_OpCvt32Uto64F_0(v)
	case OpCvt32to32F:
		return rewriteValueARM_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValueARM_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValueARM_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValueARM_OpCvt64Fto32F_0(v)
	case OpCvt64Fto32U:
		return rewriteValueARM_OpCvt64Fto32U_0(v)
	case OpDiv16:
		return rewriteValueARM_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValueARM_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValueARM_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValueARM_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValueARM_OpDiv32u_0(v)
	case OpDiv64F:
		return rewriteValueARM_OpDiv64F_0(v)
	case OpDiv8:
		return rewriteValueARM_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValueARM_OpDiv8u_0(v)
	case OpEq16:
		return psess.rewriteValueARM_OpEq16_0(v)
	case OpEq32:
		return psess.rewriteValueARM_OpEq32_0(v)
	case OpEq32F:
		return psess.rewriteValueARM_OpEq32F_0(v)
	case OpEq64F:
		return psess.rewriteValueARM_OpEq64F_0(v)
	case OpEq8:
		return psess.rewriteValueARM_OpEq8_0(v)
	case OpEqB:
		return rewriteValueARM_OpEqB_0(v)
	case OpEqPtr:
		return psess.rewriteValueARM_OpEqPtr_0(v)
	case OpGeq16:
		return psess.rewriteValueARM_OpGeq16_0(v)
	case OpGeq16U:
		return psess.rewriteValueARM_OpGeq16U_0(v)
	case OpGeq32:
		return psess.rewriteValueARM_OpGeq32_0(v)
	case OpGeq32F:
		return psess.rewriteValueARM_OpGeq32F_0(v)
	case OpGeq32U:
		return psess.rewriteValueARM_OpGeq32U_0(v)
	case OpGeq64F:
		return psess.rewriteValueARM_OpGeq64F_0(v)
	case OpGeq8:
		return psess.rewriteValueARM_OpGeq8_0(v)
	case OpGeq8U:
		return psess.rewriteValueARM_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValueARM_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValueARM_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValueARM_OpGetClosurePtr_0(v)
	case OpGreater16:
		return psess.rewriteValueARM_OpGreater16_0(v)
	case OpGreater16U:
		return psess.rewriteValueARM_OpGreater16U_0(v)
	case OpGreater32:
		return psess.rewriteValueARM_OpGreater32_0(v)
	case OpGreater32F:
		return psess.rewriteValueARM_OpGreater32F_0(v)
	case OpGreater32U:
		return psess.rewriteValueARM_OpGreater32U_0(v)
	case OpGreater64F:
		return psess.rewriteValueARM_OpGreater64F_0(v)
	case OpGreater8:
		return psess.rewriteValueARM_OpGreater8_0(v)
	case OpGreater8U:
		return psess.rewriteValueARM_OpGreater8U_0(v)
	case OpHmul32:
		return rewriteValueARM_OpHmul32_0(v)
	case OpHmul32u:
		return rewriteValueARM_OpHmul32u_0(v)
	case OpInterCall:
		return rewriteValueARM_OpInterCall_0(v)
	case OpIsInBounds:
		return psess.rewriteValueARM_OpIsInBounds_0(v)
	case OpIsNonNil:
		return psess.rewriteValueARM_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return psess.rewriteValueARM_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return psess.rewriteValueARM_OpLeq16_0(v)
	case OpLeq16U:
		return psess.rewriteValueARM_OpLeq16U_0(v)
	case OpLeq32:
		return psess.rewriteValueARM_OpLeq32_0(v)
	case OpLeq32F:
		return psess.rewriteValueARM_OpLeq32F_0(v)
	case OpLeq32U:
		return psess.rewriteValueARM_OpLeq32U_0(v)
	case OpLeq64F:
		return psess.rewriteValueARM_OpLeq64F_0(v)
	case OpLeq8:
		return psess.rewriteValueARM_OpLeq8_0(v)
	case OpLeq8U:
		return psess.rewriteValueARM_OpLeq8U_0(v)
	case OpLess16:
		return psess.rewriteValueARM_OpLess16_0(v)
	case OpLess16U:
		return psess.rewriteValueARM_OpLess16U_0(v)
	case OpLess32:
		return psess.rewriteValueARM_OpLess32_0(v)
	case OpLess32F:
		return psess.rewriteValueARM_OpLess32F_0(v)
	case OpLess32U:
		return psess.rewriteValueARM_OpLess32U_0(v)
	case OpLess64F:
		return psess.rewriteValueARM_OpLess64F_0(v)
	case OpLess8:
		return psess.rewriteValueARM_OpLess8_0(v)
	case OpLess8U:
		return psess.rewriteValueARM_OpLess8U_0(v)
	case OpLoad:
		return psess.rewriteValueARM_OpLoad_0(v)
	case OpLsh16x16:
		return psess.rewriteValueARM_OpLsh16x16_0(v)
	case OpLsh16x32:
		return psess.rewriteValueARM_OpLsh16x32_0(v)
	case OpLsh16x64:
		return rewriteValueARM_OpLsh16x64_0(v)
	case OpLsh16x8:
		return rewriteValueARM_OpLsh16x8_0(v)
	case OpLsh32x16:
		return psess.rewriteValueARM_OpLsh32x16_0(v)
	case OpLsh32x32:
		return psess.rewriteValueARM_OpLsh32x32_0(v)
	case OpLsh32x64:
		return rewriteValueARM_OpLsh32x64_0(v)
	case OpLsh32x8:
		return rewriteValueARM_OpLsh32x8_0(v)
	case OpLsh8x16:
		return psess.rewriteValueARM_OpLsh8x16_0(v)
	case OpLsh8x32:
		return psess.rewriteValueARM_OpLsh8x32_0(v)
	case OpLsh8x64:
		return rewriteValueARM_OpLsh8x64_0(v)
	case OpLsh8x8:
		return rewriteValueARM_OpLsh8x8_0(v)
	case OpMod16:
		return rewriteValueARM_OpMod16_0(v)
	case OpMod16u:
		return rewriteValueARM_OpMod16u_0(v)
	case OpMod32:
		return rewriteValueARM_OpMod32_0(v)
	case OpMod32u:
		return rewriteValueARM_OpMod32u_0(v)
	case OpMod8:
		return rewriteValueARM_OpMod8_0(v)
	case OpMod8u:
		return rewriteValueARM_OpMod8u_0(v)
	case OpMove:
		return psess.rewriteValueARM_OpMove_0(v)
	case OpMul16:
		return rewriteValueARM_OpMul16_0(v)
	case OpMul32:
		return rewriteValueARM_OpMul32_0(v)
	case OpMul32F:
		return rewriteValueARM_OpMul32F_0(v)
	case OpMul32uhilo:
		return rewriteValueARM_OpMul32uhilo_0(v)
	case OpMul64F:
		return rewriteValueARM_OpMul64F_0(v)
	case OpMul8:
		return rewriteValueARM_OpMul8_0(v)
	case OpNeg16:
		return rewriteValueARM_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValueARM_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValueARM_OpNeg32F_0(v)
	case OpNeg64F:
		return rewriteValueARM_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValueARM_OpNeg8_0(v)
	case OpNeq16:
		return psess.rewriteValueARM_OpNeq16_0(v)
	case OpNeq32:
		return psess.rewriteValueARM_OpNeq32_0(v)
	case OpNeq32F:
		return psess.rewriteValueARM_OpNeq32F_0(v)
	case OpNeq64F:
		return psess.rewriteValueARM_OpNeq64F_0(v)
	case OpNeq8:
		return psess.rewriteValueARM_OpNeq8_0(v)
	case OpNeqB:
		return rewriteValueARM_OpNeqB_0(v)
	case OpNeqPtr:
		return psess.rewriteValueARM_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValueARM_OpNilCheck_0(v)
	case OpNot:
		return rewriteValueARM_OpNot_0(v)
	case OpOffPtr:
		return rewriteValueARM_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValueARM_OpOr16_0(v)
	case OpOr32:
		return rewriteValueARM_OpOr32_0(v)
	case OpOr8:
		return rewriteValueARM_OpOr8_0(v)
	case OpOrB:
		return rewriteValueARM_OpOrB_0(v)
	case OpRound32F:
		return rewriteValueARM_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValueARM_OpRound64F_0(v)
	case OpRsh16Ux16:
		return psess.rewriteValueARM_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return psess.rewriteValueARM_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return rewriteValueARM_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return rewriteValueARM_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return psess.rewriteValueARM_OpRsh16x16_0(v)
	case OpRsh16x32:
		return psess.rewriteValueARM_OpRsh16x32_0(v)
	case OpRsh16x64:
		return rewriteValueARM_OpRsh16x64_0(v)
	case OpRsh16x8:
		return rewriteValueARM_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return psess.rewriteValueARM_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return psess.rewriteValueARM_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return rewriteValueARM_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return rewriteValueARM_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return psess.rewriteValueARM_OpRsh32x16_0(v)
	case OpRsh32x32:
		return psess.rewriteValueARM_OpRsh32x32_0(v)
	case OpRsh32x64:
		return rewriteValueARM_OpRsh32x64_0(v)
	case OpRsh32x8:
		return rewriteValueARM_OpRsh32x8_0(v)
	case OpRsh8Ux16:
		return psess.rewriteValueARM_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return psess.rewriteValueARM_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return rewriteValueARM_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return rewriteValueARM_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return psess.rewriteValueARM_OpRsh8x16_0(v)
	case OpRsh8x32:
		return psess.rewriteValueARM_OpRsh8x32_0(v)
	case OpRsh8x64:
		return rewriteValueARM_OpRsh8x64_0(v)
	case OpRsh8x8:
		return rewriteValueARM_OpRsh8x8_0(v)
	case OpSelect0:
		return rewriteValueARM_OpSelect0_0(v)
	case OpSelect1:
		return rewriteValueARM_OpSelect1_0(v)
	case OpSignExt16to32:
		return rewriteValueARM_OpSignExt16to32_0(v)
	case OpSignExt8to16:
		return rewriteValueARM_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValueARM_OpSignExt8to32_0(v)
	case OpSignmask:
		return rewriteValueARM_OpSignmask_0(v)
	case OpSlicemask:
		return rewriteValueARM_OpSlicemask_0(v)
	case OpSqrt:
		return rewriteValueARM_OpSqrt_0(v)
	case OpStaticCall:
		return rewriteValueARM_OpStaticCall_0(v)
	case OpStore:
		return psess.rewriteValueARM_OpStore_0(v)
	case OpSub16:
		return rewriteValueARM_OpSub16_0(v)
	case OpSub32:
		return rewriteValueARM_OpSub32_0(v)
	case OpSub32F:
		return rewriteValueARM_OpSub32F_0(v)
	case OpSub32carry:
		return rewriteValueARM_OpSub32carry_0(v)
	case OpSub32withcarry:
		return rewriteValueARM_OpSub32withcarry_0(v)
	case OpSub64F:
		return rewriteValueARM_OpSub64F_0(v)
	case OpSub8:
		return rewriteValueARM_OpSub8_0(v)
	case OpSubPtr:
		return rewriteValueARM_OpSubPtr_0(v)
	case OpTrunc16to8:
		return rewriteValueARM_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValueARM_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValueARM_OpTrunc32to8_0(v)
	case OpWB:
		return rewriteValueARM_OpWB_0(v)
	case OpXor16:
		return rewriteValueARM_OpXor16_0(v)
	case OpXor32:
		return rewriteValueARM_OpXor32_0(v)
	case OpXor8:
		return rewriteValueARM_OpXor8_0(v)
	case OpZero:
		return psess.rewriteValueARM_OpZero_0(v)
	case OpZeroExt16to32:
		return rewriteValueARM_OpZeroExt16to32_0(v)
	case OpZeroExt8to16:
		return rewriteValueARM_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValueARM_OpZeroExt8to32_0(v)
	case OpZeromask:
		return rewriteValueARM_OpZeromask_0(v)
	}
	return false
}
func rewriteValueARM_OpARMADC_0(v *Value) bool {

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v.Args[2]
		v.reset(OpARMADCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v.Args[2]
		v.reset(OpARMADCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADC_10(v *Value) bool {

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADC_20(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v.Args[1]
		v.reset(OpARMADCconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v.Args[1]
		v.reset(OpARMADCconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMADCconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		flags := v.Args[3]
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[3]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v.Args[3]
		v.reset(OpARMADCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMADCconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		flags := v.Args[3]
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[3]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v.Args[3]
		v.reset(OpARMADCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMADCconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		flags := v.Args[3]
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[3]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v.Args[3]
		v.reset(OpARMADCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMADDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMADDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMADDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMADDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMADDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMADDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMADDshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMADDshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADD_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMADDshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMADDshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMADDshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMADDshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMRSBconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpARMSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMRSBconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMRSBconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMRSBconst {
			break
		}
		d := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = c + d
		v0 := b.NewValue0(v.Pos, OpARMADD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMRSBconst {
			break
		}
		d := v_0.AuxInt
		y := v_0.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMRSBconst {
			break
		}
		c := v_1.AuxInt
		x := v_1.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = c + d
		v0 := b.NewValue0(v.Pos, OpARMADD, t)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMUL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		v.reset(OpARMMULA)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMUL {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		v.reset(OpARMMULA)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(a)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMADDD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMULD {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULAD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMULD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULAD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMNMULD {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULSD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMNMULD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULSD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMADDF_0(v *Value) bool {

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMULF {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULAF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMULF {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULAF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMNMULF {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULSF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMNMULF {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULSF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDS_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMADDSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMADDSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMADDSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMADDSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMADDSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMADDSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMADDSshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMADDSshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDS_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMADDSshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMADDSshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMADDSshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMADDSshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDSconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDSconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDSconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDconst_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpARMMOVWaddr)
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
		x := v.Args[0]
		if !(!isARMImmRot(uint32(c)) && isARMImmRot(uint32(-c))) {
			break
		}
		v.reset(OpARMSUBconst)
		v.AuxInt = int64(int32(-c))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(c + d))
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMRSBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMAND_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMANDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMANDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMANDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMANDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMANDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMANDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMANDshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMANDshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMAND_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMANDshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMANDshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMANDshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMANDshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
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
		if v_1.Op != OpARMMVN {
			break
		}
		y := v_1.Args[0]
		v.reset(OpARMBIC)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMVN {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMBIC)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMVNshiftLL {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMBICshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMVNshiftLL {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMBICshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMVNshiftRL {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMBICshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMAND_20(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMVNshiftRL {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMBICshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMVNshiftRA {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMBICshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMVNshiftRA {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMBICshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMANDconst_0(v *Value) bool {

	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpARMMOVWconst)
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
		x := v.Args[0]
		if !(!isARMImmRot(uint32(c)) && isARMImmRot(^uint32(c))) {
			break
		}
		v.reset(OpARMBICconst)
		v.AuxInt = int64(int32(^uint32(c)))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = c & d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMANDshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMANDconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARMSLLconst {
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
func rewriteValueARM_OpARMANDshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMANDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMANDshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMANDconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARMSRAconst {
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
func rewriteValueARM_OpARMANDshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMANDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMANDshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMANDconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARMSRLconst {
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
func rewriteValueARM_OpARMANDshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMANDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMBFX_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(d) << (32 - uint32(c&0xff) - uint32(c>>8)) >> (32 - uint32(c>>8)))
		return true
	}
	return false
}
func rewriteValueARM_OpARMBFXU_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(uint32(d) << (32 - uint32(c&0xff) - uint32(c>>8)) >> (32 - uint32(c>>8)))
		return true
	}
	return false
}
func rewriteValueARM_OpARMBIC_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMBICconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMBICshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMBICshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMBICshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMBICshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMBICshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMBICshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICconst_0(v *Value) bool {

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
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		c := v.AuxInt
		x := v.Args[0]
		if !(!isARMImmRot(uint32(c)) && isARMImmRot(^uint32(c))) {
			break
		}
		v.reset(OpARMANDconst)
		v.AuxInt = int64(int32(^uint32(c)))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = d &^ c
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMBICconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMBICconst)
		v.AuxInt = int64(int32(c | d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftLL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMBICconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftLLreg_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMBICshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftRA_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMBICconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftRAreg_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMBICshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftRL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMBICconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftRLreg_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMBICshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMN_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMCMNshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMCMNshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMCMNshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMCMNshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMCMNshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMCMNshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMCMNshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMCMNshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMN_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMCMNshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMCMNshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMCMNshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMCMNshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMRSBconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpARMCMP)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMRSBconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMCMP)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(-y)) {
			break
		}
		v.reset(OpARMFlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(-y) && uint32(x) < uint32(-y)) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(-y) && uint32(x) > uint32(-y)) {
			break
		}
		v.reset(OpARMFlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(-y) && uint32(x) < uint32(-y)) {
			break
		}
		v.reset(OpARMFlagGT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(-y) && uint32(x) > uint32(-y)) {
			break
		}
		v.reset(OpARMFlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMNconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMNshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMNconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMNshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMNconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMNshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMOVWHSconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMInvertFlags {
			break
		}
		flags := v_1.Args[0]
		v.reset(OpARMCMOVWLSconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMOVWLSconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMInvertFlags {
			break
		}
		flags := v_1.Args[0]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMCMP_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMPconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMCMPshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftLL, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMCMPshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftRL, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMCMPshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftRA, psess.types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMCMPshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftLLreg, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMCMP_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMCMPshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftRLreg, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMCMPshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftRAreg, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMRSBconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpARMCMN)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVDconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpARMCMPD0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPF_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVFconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpARMCMPF0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpARMFlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpARMFlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpARMFlagGT_ULT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpARMFlagGT_UGT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVBUreg {
			break
		}
		if !(0xff < c) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVHUreg {
			break
		}
		if !(0xffff < c) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMANDconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int32(m) && int32(m) < int32(n)) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}

	for {
		n := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		if !(0 <= n && 0 < c && c <= 32 && (1<<uint32(32-c)) <= uint32(n)) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMCMPshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMPconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMCMPshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMPshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMCMPshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMPconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMCMPshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMPshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMCMPshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMPconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMCMPshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMPshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMEqual_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMGreaterEqual_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMLessEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMGreaterEqualU_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMLessEqualU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMGreaterThan_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMLessThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMGreaterThanU_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMLessThanU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMLessEqual_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMGreaterEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMLessEqualU_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMGreaterEqualU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMLessThan_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMGreaterThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMLessThanU_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMGreaterThanU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBUload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVBUload)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVBUload)
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
		if v_1.Op != OpARMMOVBstore {
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
		v.reset(OpARMMOVBUreg)
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVBUloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBUloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVBstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVBUreg)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARMMOVBUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVBUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBUreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMANDconst)
		v.AuxInt = c & 0xff
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVBUreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVBload)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVBload)
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
		if v_1.Op != OpARMMOVBstore {
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
		v.reset(OpARMMOVBreg)
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVBloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVBstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVBreg)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARMMOVBload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVBload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x80 == 0) {
			break
		}
		v.reset(OpARMANDconst)
		v.AuxInt = c & 0x7f
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVBreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVBstore)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
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
		v.reset(OpARMMOVBstore)
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
		if v_1.Op != OpARMMOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARMMOVBstore)
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
		if v_1.Op != OpARMMOVBUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARMMOVBstore)
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
		if v_1.Op != OpARMMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARMMOVBstore)
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
		if v_1.Op != OpARMMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARMMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVBstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBstoreidx_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVBstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVBstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVDload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVDload)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVDload)
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
		if v_1.Op != OpARMMOVDstore {
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
func rewriteValueARM_OpARMMOVDstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVDstore)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
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
		v.reset(OpARMMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVFload_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVFload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVFload)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVFload)
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
		if v_1.Op != OpARMMOVFstore {
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
func rewriteValueARM_OpARMMOVFstore_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVFstore)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
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
		v.reset(OpARMMOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHUload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVHUload)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVHUload)
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
		if v_1.Op != OpARMMOVHstore {
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
		v.reset(OpARMMOVHUreg)
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVHUloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHUloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVHstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVHUreg)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARMMOVHUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVHUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHUreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMANDconst)
		v.AuxInt = c & 0xffff
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVBUreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVHUreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVHload)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVHload)
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
		if v_1.Op != OpARMMOVHstore {
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
		v.reset(OpARMMOVHreg)
		v.AddArg(x)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVHloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVHstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVHreg)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARMMOVHload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVHload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x8000 == 0) {
			break
		}
		v.reset(OpARMANDconst)
		v.AuxInt = c & 0x7fff
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVBreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVBUreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		if x.Op != OpARMMOVHreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVHstore)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
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
		v.reset(OpARMMOVHstore)
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
		if v_1.Op != OpARMMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARMMOVHstore)
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
		if v_1.Op != OpARMMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARMMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVHstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHstoreidx_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVHstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVHstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWload_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		v.reset(OpARMMOVWload)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVWload)
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
		if v_1.Op != OpARMMOVWstore {
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
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVWloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDshiftLL {
			break
		}
		c := v_0.AuxInt
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVWloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDshiftRL {
			break
		}
		c := v_0.AuxInt
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVWloadshiftRL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDshiftRA {
			break
		}
		c := v_0.AuxInt
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		mem := v.Args[1]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVWloadshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWloadidx_0(v *Value) bool {

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARMMOVWload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVWload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARMMOVWloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVWloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARMMOVWloadshiftRL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVWloadshiftRL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v.Args[2]
		v.reset(OpARMMOVWloadshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVWloadshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWloadshiftLL_0(v *Value) bool {

	for {
		c := v.AuxInt
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWstoreshiftLL {
			break
		}
		d := v_2.AuxInt
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARMMOVWload)
		v.AuxInt = int64(uint32(c) << uint64(d))
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWloadshiftRA_0(v *Value) bool {

	for {
		c := v.AuxInt
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWstoreshiftRA {
			break
		}
		d := v_2.AuxInt
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARMMOVWload)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWloadshiftRL_0(v *Value) bool {

	for {
		c := v.AuxInt
		_ = v.Args[2]
		ptr := v.Args[0]
		idx := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWstoreshiftRL {
			break
		}
		d := v_2.AuxInt
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v.Args[2]
		v.reset(OpARMMOVWload)
		v.AuxInt = int64(uint32(c) >> uint64(d))
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWreg_0(v *Value) bool {

	for {
		x := v.Args[0]
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpARMMOVWnop)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWstore_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config

	for {
		off1 := v.AuxInt
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
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
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMMOVWstore)
		v.AuxInt = off1 - off2
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
		if v_0.Op != OpARMMOVWaddr {
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
		v.reset(OpARMMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADD {
			break
		}
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVWstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDshiftLL {
			break
		}
		c := v_0.AuxInt
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVWstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDshiftRL {
			break
		}
		c := v_0.AuxInt
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVWstoreshiftRL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDshiftRA {
			break
		}
		c := v_0.AuxInt
		_ = v_0.Args[1]
		ptr := v_0.Args[0]
		idx := v_0.Args[1]
		val := v.Args[1]
		mem := v.Args[2]
		if !(sym == nil && !config.nacl) {
			break
		}
		v.reset(OpARMMOVWstoreshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWstoreidx_0(v *Value) bool {

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVWstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVWstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVWstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVWstoreshiftLL)
		v.AuxInt = c
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
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVWstoreshiftRL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVWstoreshiftRL)
		v.AuxInt = c
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
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVWstoreshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v.Args[1]
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVWstoreshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWstoreshiftLL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVWstore)
		v.AuxInt = int64(uint32(c) << uint64(d))
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWstoreshiftRA_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVWstore)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWstoreshiftRL_0(v *Value) bool {

	for {
		d := v.AuxInt
		_ = v.Args[3]
		ptr := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v.Args[2]
		mem := v.Args[3]
		v.reset(OpARMMOVWstore)
		v.AuxInt = int64(uint32(c) >> uint64(d))
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMUL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
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
		if v_0.Op != OpARMMOVWconst {
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARMADDshiftLL)
		v.AuxInt = log2(c - 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARMADDshiftLL)
		v.AuxInt = log2(c - 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMUL_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARMRSBshiftLL)
		v.AuxInt = log2(c + 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARMRSBshiftLL)
		v.AuxInt = log2(c + 1)
		v.AddArg(x)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = 1
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = log2(c / 3)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = log2(c / 5)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = 2
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = log2(c / 5)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = log2(c / 7)
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = log2(c / 9)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = log2(c / 9)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = 3
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMUL_20(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		d := v_1.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(c * d))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(c * d))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMULA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMSUB)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		a := v.Args[2]
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		a := v.Args[2]
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMULA_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMSUB)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		a := v.Args[2]
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		a := v.Args[2]
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMULA_20(v *Value) bool {

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		d := v_1.AuxInt
		a := v.Args[2]
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(c * d))
		v.AddArg(a)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMMULD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMNEGD {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		if !(psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMNMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMNEGD {
			break
		}
		x := v_1.Args[0]
		if !(psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMNMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMMULF_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMNEGF {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		if !(psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMNMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMNEGF {
			break
		}
		x := v_1.Args[0]
		if !(psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMNMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMULS_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMADD)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		a := v.Args[2]
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		if v_1.AuxInt != 1 {
			break
		}
		a := v.Args[2]
		v.reset(OpARMRSB)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v.Args[2]
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMULS_10(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMADD)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		if v_0.AuxInt != 0 {
			break
		}
		a := v.Args[2]
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		if v_0.AuxInt != 1 {
			break
		}
		x := v.Args[1]
		a := v.Args[2]
		v.reset(OpARMRSB)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		a := v.Args[2]
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMULS_20(v *Value) bool {

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		d := v_1.AuxInt
		a := v.Args[2]
		v.reset(OpARMSUBconst)
		v.AuxInt = int64(int32(c * d))
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVN_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = ^c
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMMVNshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMMVNshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMMVNshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpARMMVNshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpARMMVNshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		v.reset(OpARMMVNshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftLL_0(v *Value) bool {

	for {
		d := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = ^int64(uint32(c) << uint64(d))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftLLreg_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMMVNshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftRA_0(v *Value) bool {

	for {
		d := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = ^int64(int32(c) >> uint64(d))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftRAreg_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMMVNshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftRL_0(v *Value) bool {

	for {
		d := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = ^int64(uint32(c) >> uint64(d))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftRLreg_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMMVNshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMNEGD_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMMULD {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMNMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMNEGF_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMMULF {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		if !(psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMNMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMNMULD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMNEGD {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpARMMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMNEGD {
			break
		}
		x := v_1.Args[0]
		v.reset(OpARMMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMNMULF_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMNEGF {
			break
		}
		x := v_0.Args[0]
		y := v.Args[1]
		v.reset(OpARMMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		y := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMNEGF {
			break
		}
		x := v_1.Args[0]
		v.reset(OpARMMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMNotEqual_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMNotEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMOR_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMORshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMORshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMOR_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMORshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMORshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMORshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMORshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
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
func rewriteValueARM_OpARMORconst_0(v *Value) bool {

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
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = -1
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = c | d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMORshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMORconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARMSLLconst {
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
func rewriteValueARM_OpARMORshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMORshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMORconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARMSRAconst {
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
func rewriteValueARM_OpARMORshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMORshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMORconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		if y.Op != OpARMSRLconst {
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
func rewriteValueARM_OpARMORshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMRSBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMSUBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMRSBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMSUBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMRSBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMSUBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMRSBshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMSUBshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMRSB_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMRSBshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMSUBshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMRSBshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMSUBshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMUL {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		y := v_0.Args[1]
		a := v.Args[1]
		if !(psess.objabi.GOARM == 7) {
			break
		}
		v.reset(OpARMMULS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBSconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBSconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBSconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(c - d))
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMRSBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v.Args[1]
		v.reset(OpARMRSCconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v.Args[1]
		v.reset(OpARMRSCconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMRSCconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		flags := v.Args[3]
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[3]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v.Args[3]
		v.reset(OpARMRSCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMRSCconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		flags := v.Args[3]
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[3]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v.Args[3]
		v.reset(OpARMRSCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMRSCconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		flags := v.Args[3]
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[3]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v.Args[3]
		v.reset(OpARMRSCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBC_0(v *Value) bool {

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v.Args[2]
		v.reset(OpARMSBCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMRSCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v.Args[2]
		v.reset(OpARMSBCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMRSCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v.Args[2]
		v.reset(OpARMSBCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMRSCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		flags := v.Args[2]
		v.reset(OpARMSBCshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMRSCshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBC_10(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		flags := v.Args[2]
		v.reset(OpARMSBCshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMRSCshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		flags := v.Args[2]
		v.reset(OpARMSBCshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMRSCshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v.Args[1]
		v.reset(OpARMSBCconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v.Args[1]
		v.reset(OpARMSBCconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMSBCconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		flags := v.Args[3]
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[3]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v.Args[3]
		v.reset(OpARMSBCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMSBCconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		flags := v.Args[3]
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[3]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v.Args[3]
		v.reset(OpARMSBCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		flags := v.Args[2]
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[2]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v.Args[2]
		v.reset(OpARMSBCconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[3]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		flags := v.Args[3]
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}

	for {
		_ = v.Args[3]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v.Args[3]
		v.reset(OpARMSBCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSLL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSLLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSLLconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(uint32(d) << uint64(c))
		return true
	}
	return false
}
func rewriteValueARM_OpARMSRA_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSRAconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSRAcond_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMSRA)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMSRA)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		v_2 := v.Args[2]
		if v_2.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMSRAconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(d) >> uint64(c))
		return true
	}

	for {
		d := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(psess.objabi.GOARM == 7 && uint64(d) >= uint64(c) && uint64(d) <= 31) {
			break
		}
		v.reset(OpARMBFX)
		v.AuxInt = (d - c) | (32-d)<<8
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSRL_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSRLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMSRLconst_0(v *Value) bool {

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(uint32(d) >> uint64(c))
		return true
	}

	for {
		d := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(psess.objabi.GOARM == 7 && uint64(d) >= uint64(c) && uint64(d) <= 31) {
			break
		}
		v.reset(OpARMBFXU)
		v.AuxInt = (d - c) | (32-d)<<8
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMRSBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMRSBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMRSBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMSUBshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMRSBshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMSUB_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMSUBshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMRSBshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMSUBshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMRSBshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMUL {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		if !(psess.objabi.GOARM == 7) {
			break
		}
		v.reset(OpARMMULS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(a)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMSUBD_0(v *Value) bool {

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMULD {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULSD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMNMULD {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULAD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpARMSUBF_0(v *Value) bool {

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMULF {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULSF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		a := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMNMULF {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		y := v_1.Args[1]
		if !(a.Uses == 1 && psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULAF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBS_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMRSBSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMRSBSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMRSBSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMSUBSshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMRSBSshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMSUBSshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBS_10(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMRSBSshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMSUBSshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMRSBSshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBSconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBSconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBSconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBconst_0(v *Value) bool {

	for {
		off1 := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpARMMOVWaddr)
		v.AuxInt = off2 - off1
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
		x := v.Args[0]
		if !(!isARMImmRot(uint32(c)) && isARMImmRot(uint32(-c))) {
			break
		}
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(-c))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(d - c))
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(-c - d))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(-c + d))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMRSBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(-c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQ_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMTEQshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMTEQshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMTEQshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMTEQshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMTEQshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMTEQshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMTEQshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMTEQshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQ_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMTEQshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMTEQshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMTEQshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMTEQshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x^y) == 0) {
			break
		}
		v.reset(OpARMFlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x^y) < 0) {
			break
		}
		v.reset(OpARMFlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x^y) > 0) {
			break
		}
		v.reset(OpARMFlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTEQconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTEQshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTEQconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTEQshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTEQconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTEQshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTST_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMTSTshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMTSTshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMTSTshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMTSTshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMTSTshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMTSTshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMTSTshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMTSTshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTST_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMTSTshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMTSTshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMTSTshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMTSTshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTconst_0(v *Value) bool {

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) == 0) {
			break
		}
		v.reset(OpARMFlagEQ)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) < 0) {
			break
		}
		v.reset(OpARMFlagLT_UGT)
		return true
	}

	for {
		y := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) > 0) {
			break
		}
		v.reset(OpARMFlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTSTconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTSTshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTSTconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTSTshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTSTconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTSTshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMXOR_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMXORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMXORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMXORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMXORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMXORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMXORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRRconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMXORshiftRR)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRRconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v.Args[1]
		v.reset(OpARMXORshiftRR)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMXOR_10(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMXORshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMXORshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRL {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMXORshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRL {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMXORshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRA {
			break
		}
		_ = v_1.Args[1]
		y := v_1.Args[0]
		z := v_1.Args[1]
		v.reset(OpARMXORshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRA {
			break
		}
		_ = v_0.Args[1]
		y := v_0.Args[0]
		z := v_0.Args[1]
		x := v.Args[1]
		v.reset(OpARMXORshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORconst_0(v *Value) bool {

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
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = c ^ d
		return true
	}

	for {
		c := v.AuxInt
		v_0 := v.Args[0]
		if v_0.Op != OpARMXORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMXORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftLL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMXORconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSRLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftLLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMXORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftRA_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMXORconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftRAreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMXORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftRL_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMXORconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}

	for {
		c := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMSLLconst {
			break
		}
		if v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v.Args[1] {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}

	for {
		d := v.AuxInt
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] {
			break
		}
		if !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftRLreg_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		y := v.Args[2]
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		v_2 := v.Args[2]
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMXORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftRR_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		d := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v.Args[1]
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRRconst, x.Type)
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
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMXORconst)
		v.AuxInt = int64(int32(uint32(c)>>uint64(d) | uint32(c)<<uint64(32-d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpAdd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAdd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAdd32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMADDF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAdd32carry_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMADDS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAdd32withcarry_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		c := v.Args[2]
		v.reset(OpARMADC)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(c)
		return true
	}
}
func rewriteValueARM_OpAdd64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMADDD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAdd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAddPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAddr_0(v *Value) bool {

	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpARMMOVWaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueARM_OpAnd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAnd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAnd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAndB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAvg32u_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpARMSUB, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpBitLen32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpARMCLZ, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpBswap32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		if !(psess.objabi.GOARM == 5) {
			break
		}
		v.reset(OpARMXOR)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, t)
		v0.AuxInt = 8
		v1 := b.NewValue0(v.Pos, OpARMBICconst, t)
		v1.AuxInt = 0xff0000
		v2 := b.NewValue0(v.Pos, OpARMXOR, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpARMSRRconst, t)
		v3.AuxInt = 16
		v3.AddArg(x)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpARMSRRconst, t)
		v4.AuxInt = 8
		v4.AddArg(x)
		v.AddArg(v4)
		return true
	}

	for {
		x := v.Args[0]
		if !(psess.objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMREV)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpClosureCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMCALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM_OpCom16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCom32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCom8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpConst16_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM_OpConst32_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM_OpConst32F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpARMMOVFconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM_OpConst64F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpARMMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM_OpConst8_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM_OpConstBool_0(v *Value) bool {

	for {
		b := v.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValueARM_OpConstNil_0(v *Value) bool {

	for {
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpCtz32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		if !(psess.objabi.GOARM <= 6) {
			break
		}
		v.reset(OpARMRSBconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpARMCLZ, t)
		v1 := b.NewValue0(v.Pos, OpARMSUBconst, t)
		v1.AuxInt = 1
		v2 := b.NewValue0(v.Pos, OpARMAND, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpARMRSBconst, t)
		v3.AuxInt = 0
		v3.AddArg(x)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}

	for {
		t := v.Type
		x := v.Args[0]
		if !(psess.objabi.GOARM == 7) {
			break
		}
		v.reset(OpARMCLZ)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpARMRBIT, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM_OpCtz32NonZero_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCtz32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVFW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32Fto32U_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVFWU)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32Fto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVFD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32Uto32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVWUF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32Uto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVWUD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32to32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVWF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32to64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVWD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt64Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVDW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt64Fto32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVDF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt64Fto32U_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVDWU)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpDiv16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpDiv32)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpDiv16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpDiv32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpDiv32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSUB)
		v0 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpSelect0, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpARMCALLudiv, types.NewTuple(typ.UInt32, typ.UInt32))
		v3 := b.NewValue0(v.Pos, OpARMSUB, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v3.AddArg(v4)
		v6 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v6.AddArg(x)
		v3.AddArg(v6)
		v2.AddArg(v3)
		v7 := b.NewValue0(v.Pos, OpARMSUB, typ.UInt32)
		v8 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v8.AddArg(y)
		v9 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v9.AddArg(y)
		v8.AddArg(v9)
		v7.AddArg(v8)
		v10 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v10.AddArg(y)
		v7.AddArg(v10)
		v2.AddArg(v7)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v11 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v12 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v12.AddArg(x)
		v12.AddArg(y)
		v11.AddArg(v12)
		v0.AddArg(v11)
		v.AddArg(v0)
		v13 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v14 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v14.AddArg(x)
		v14.AddArg(y)
		v13.AddArg(v14)
		v.AddArg(v13)
		return true
	}
}
func rewriteValueARM_OpDiv32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMDIVF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpDiv32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect0)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpARMCALLudiv, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpDiv64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMDIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpDiv8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpDiv32)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpDiv8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpDiv32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpEq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpEq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpEq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpEq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpEq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func rewriteValueARM_OpEqB_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpARMXOR, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpEqPtr_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpGeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpGeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpGeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpGeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpGeq32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpGeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpGeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpGeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func rewriteValueARM_OpGetCallerPC_0(v *Value) bool {

	for {
		v.reset(OpARMLoweredGetCallerPC)
		return true
	}
}
func rewriteValueARM_OpGetCallerSP_0(v *Value) bool {

	for {
		v.reset(OpARMLoweredGetCallerSP)
		return true
	}
}
func rewriteValueARM_OpGetClosurePtr_0(v *Value) bool {

	for {
		v.reset(OpARMLoweredGetClosurePtr)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpGreater16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpGreater16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpGreater32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpGreater32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpGreater32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpGreater64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpGreater8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpGreater8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func rewriteValueARM_OpHmul32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMHMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpHmul32u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMHMULU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpInterCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[1]
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARMCALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpIsInBounds_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpARMLessThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		ptr := v.Args[0]
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(ptr)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpIsSliceInBounds_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpARMLessEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpLeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpLeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLeq32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpLeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpLess16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpLess16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpLess32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLess32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLess32U_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLess64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, psess.types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLess8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpLess8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMLessThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpLoad_0(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpARMMOVBUload)
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
		v.reset(OpARMMOVBload)
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
		v.reset(OpARMMOVBUload)
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
		v.reset(OpARMMOVHload)
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
		v.reset(OpARMMOVHUload)
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
		v.reset(OpARMMOVWload)
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
		v.reset(OpARMMOVFload)
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
		v.reset(OpARMMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueARM_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpLsh16x64_0(v *Value) bool {

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
		v.reset(OpARMSLLconst)
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
func rewriteValueARM_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpLsh32x64_0(v *Value) bool {

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
		v.reset(OpARMSLLconst)
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
func rewriteValueARM_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpLsh8x64_0(v *Value) bool {

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
		v.reset(OpARMSLLconst)
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
func rewriteValueARM_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpMod16_0(v *Value) bool {
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
func rewriteValueARM_OpMod16u_0(v *Value) bool {
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
func rewriteValueARM_OpMod32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSUB)
		v0 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpSelect1, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpARMCALLudiv, types.NewTuple(typ.UInt32, typ.UInt32))
		v3 := b.NewValue0(v.Pos, OpARMSUB, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v3.AddArg(v4)
		v6 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v6.AddArg(x)
		v3.AddArg(v6)
		v2.AddArg(v3)
		v7 := b.NewValue0(v.Pos, OpARMSUB, typ.UInt32)
		v8 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v8.AddArg(y)
		v9 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v9.AddArg(y)
		v8.AddArg(v9)
		v7.AddArg(v8)
		v10 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v10.AddArg(y)
		v7.AddArg(v10)
		v2.AddArg(v7)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v11 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v11.AddArg(x)
		v0.AddArg(v11)
		v.AddArg(v0)
		v12 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v12.AddArg(x)
		v.AddArg(v12)
		return true
	}
}
func rewriteValueARM_OpMod32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpSelect1)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpARMCALLudiv, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpMod8_0(v *Value) bool {
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
func rewriteValueARM_OpMod8u_0(v *Value) bool {
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
func (psess *PackageSession) rewriteValueARM_OpMove_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
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
		v.reset(OpARMMOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
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
		v.reset(OpARMMOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVHUload, typ.UInt16)
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
		v.reset(OpARMMOVBstore)
		v.AuxInt = 1
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v0.AuxInt = 1
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
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
		v.reset(OpARMMOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVWload, typ.UInt32)
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
		v.reset(OpARMMOVHstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVHUload, typ.UInt16)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVHstore, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARMMOVHUload, typ.UInt16)
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
		v.reset(OpARMMOVBstore)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v4.AuxInt = 1
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
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
		v.reset(OpARMMOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v2.AuxInt = 1
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
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
		if !(s%4 == 0 && s > 4 && s <= 512 && t.(*types.Type).Alignment(psess.types)%4 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpARMDUFFCOPY)
		v.AuxInt = 8 * (128 - s/4)
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}

	for {
		s := v.AuxInt
		t := v.Aux
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !((s > 512 || config.noDuffDevice) || t.(*types.Type).Alignment(psess.types)%4 != 0) {
			break
		}
		v.reset(OpARMLoweredMove)
		v.AuxInt = t.(*types.Type).Alignment(psess.types)
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpARMADDconst, src.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(psess.types), config)
		v0.AddArg(src)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpMul16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpMul32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpMul32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpMul32uhilo_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMMULLU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpMul64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpMul8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpNeg16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpNeg32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpNeg32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMNEGF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpNeg64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMNEGD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpNeg8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpNeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func (psess *PackageSession) rewriteValueARM_OpNeq32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpNeq32F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpNeq64F_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpNeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
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
func rewriteValueARM_OpNeqB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpNeqPtr_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpNilCheck_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpARMLoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM_OpNot_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMXORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpOffPtr_0(v *Value) bool {

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpARMMOVWaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpARMADDconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueARM_OpOr16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpOr32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpOr8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpOrB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpRound32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpRound64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v3.AuxInt = 256
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v2.AuxInt = 256
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueARM_OpRsh16Ux64_0(v *Value) bool {
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
		v.reset(OpARMSRLconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
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
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRL)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpRsh16x64_0(v *Value) bool {
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
		v.reset(OpARMSRAconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
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
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpRsh32Ux64_0(v *Value) bool {

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
		v.reset(OpARMSRLconst)
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
func rewriteValueARM_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRAcond)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v1.AuxInt = 256
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRAcond)
		v.AddArg(x)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v0.AuxInt = 256
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpRsh32x64_0(v *Value) bool {

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
		v.reset(OpARMSRAconst)
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
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v3.AuxInt = 256
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v2.AuxInt = 256
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueARM_OpRsh8Ux64_0(v *Value) bool {
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
		v.reset(OpARMSRLconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
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
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRL)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpRsh8x64_0(v *Value) bool {
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
		v.reset(OpARMSRAconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
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
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpSelect0_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst {
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
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMSRLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(uint32(c) / uint32(d))
		return true
	}
	return false
}
func rewriteValueARM_OpSelect1_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst {
			break
		}
		if v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMANDconst)
		v.AuxInt = c - 1
		v.AddArg(x)
		return true
	}

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(uint32(c) % uint32(d))
		return true
	}
	return false
}
func rewriteValueARM_OpSignExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpSignExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpSignExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpSignmask_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpSlicemask_0(v *Value) bool {
	b := v.Block
	_ = b

	for {
		t := v.Type
		x := v.Args[0]
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpARMRSBconst, t)
		v0.AuxInt = 0
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpSqrt_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMSQRTD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpStaticCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpARMCALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpStore_0(v *Value) bool {

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(t.(*types.Type).Size(psess.types) == 1) {
			break
		}
		v.reset(OpARMMOVBstore)
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
		v.reset(OpARMMOVHstore)
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
		v.reset(OpARMMOVWstore)
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
		v.reset(OpARMMOVFstore)
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
		v.reset(OpARMMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpSub16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSub32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSub32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSUBF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSub32carry_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSUBS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSub32withcarry_0(v *Value) bool {

	for {
		_ = v.Args[2]
		x := v.Args[0]
		y := v.Args[1]
		c := v.Args[2]
		v.reset(OpARMSBC)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(c)
		return true
	}
}
func rewriteValueARM_OpSub64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSUBD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSub8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSubPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpTrunc16to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpTrunc32to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpTrunc32to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpWB_0(v *Value) bool {

	for {
		fn := v.Aux
		_ = v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpARMLoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM_OpXor16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpXor32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpXor8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpARMXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueARM_OpZero_0(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
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
		v.reset(OpARMMOVBstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
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
		v.reset(OpARMMOVHstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
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
		v.reset(OpARMMOVBstore)
		v.AuxInt = 1
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
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
		v.reset(OpARMMOVWstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
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
		v.reset(OpARMMOVHstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVHstore, psess.types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
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
		v.reset(OpARMMOVBstore)
		v.AuxInt = 3
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
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
		v.reset(OpARMMOVBstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARMMOVBstore, psess.types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
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
		if !(s%4 == 0 && s > 4 && s <= 512 && t.(*types.Type).Alignment(psess.types)%4 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpARMDUFFZERO)
		v.AuxInt = 4 * (128 - s/4)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}

	for {
		s := v.AuxInt
		t := v.Aux
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !((s > 512 || config.noDuffDevice) || t.(*types.Type).Alignment(psess.types)%4 != 0) {
			break
		}
		v.reset(OpARMLoweredZero)
		v.AuxInt = t.(*types.Type).Alignment(psess.types)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMADDconst, ptr.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(psess.types), config)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpZeroExt16to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVHUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpZeroExt8to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpZeroExt8to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpARMMOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpZeromask_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftRL, typ.Int32)
		v0.AuxInt = 1
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func (psess *PackageSession) rewriteBlockARM(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	typ := &config.Types
	_ = typ
	switch b.Kind {
	case BlockARMEQ:

		for {
			v := b.Control
			if v.Op != OpARMFlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagLT_ULT {
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
			if v.Op != OpARMFlagLT_UGT {
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
			if v.Op != OpARMFlagGT_ULT {
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
			if v.Op != OpARMFlagGT_UGT {
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
			if v.Op != OpARMInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARMEQ
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUB {
				break
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftLL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftLL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftRL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftRL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftRA {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftRA, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftLLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftLLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftRLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftRLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftRAreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftRAreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADD {
				break
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMN, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMNconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftLL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftLL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftRL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftRL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftRA {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftRA, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftLLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftLLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftRLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftRLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftRAreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftRAreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMAND {
				break
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTST, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTSTconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftLL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftLL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftRL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftRL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftRA {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftRA, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftLLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftLLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftRLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftRLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftRAreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftRAreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXOR {
				break
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTEQ, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTEQconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftLL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftLL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftRL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftRL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftRA {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftRA, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftLLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftLLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftRLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftRLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftRAreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMEQ
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftRAreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockARMGE:

		for {
			v := b.Control
			if v.Op != OpARMFlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagLT_ULT {
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
			if v.Op != OpARMFlagLT_UGT {
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
			if v.Op != OpARMFlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARMLE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARMGT:

		for {
			v := b.Control
			if v.Op != OpARMFlagEQ {
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
			if v.Op != OpARMFlagLT_ULT {
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
			if v.Op != OpARMFlagLT_UGT {
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
			if v.Op != OpARMFlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARMLT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockIf:

		for {
			v := b.Control
			if v.Op != OpARMEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARMEQ
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMNotEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARMNE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMLessThan {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARMLT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMLessThanU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARMULT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMLessEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARMLE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMLessEqualU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARMULE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMGreaterThan {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARMGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMGreaterThanU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARMUGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMGreaterEqual {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARMGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMGreaterEqualU {
				break
			}
			cc := v.Args[0]
			b.Kind = BlockARMUGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			_ = v
			cond := b.Control
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(cond)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockARMLE:

		for {
			v := b.Control
			if v.Op != OpARMFlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagGT_ULT {
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
			if v.Op != OpARMFlagGT_UGT {
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
			if v.Op != OpARMInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARMGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARMLT:

		for {
			v := b.Control
			if v.Op != OpARMFlagEQ {
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
			if v.Op != OpARMFlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagGT_ULT {
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
			if v.Op != OpARMFlagGT_UGT {
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
			if v.Op != OpARMInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARMGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARMNE:

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockARMEQ
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMNotEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockARMNE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMLessThan {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockARMLT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMLessThanU {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockARMULT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMLessEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockARMLE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMLessEqualU {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockARMULE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMGreaterThan {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockARMGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMGreaterThanU {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockARMUGT
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMGreaterEqual {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockARMGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMGreaterEqualU {
				break
			}
			cc := v_0.Args[0]
			b.Kind = BlockARMUGE
			b.SetControl(cc)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagEQ {
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
			if v.Op != OpARMFlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARMNE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUB {
				break
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMP, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMPconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftLL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftLL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftRL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftRL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftRA {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftRA, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftLLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftLLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftRLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftRLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMSUBshiftRAreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMPshiftRAreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADD {
				break
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMN, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMNconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftLL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftLL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftRL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftRL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftRA {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftRA, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftLLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftLLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftRLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftRLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMADDshiftRAreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMCMNshiftRAreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMAND {
				break
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTST, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTSTconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftLL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftLL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftRL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftRL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftRA {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftRA, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftLLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftLLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftRLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftRLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMANDshiftRAreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTSTshiftRAreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXOR {
				break
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTEQ, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORconst {
				break
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTEQconst, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftLL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftLL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftRL {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftRL, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftRA {
				break
			}
			c := v_0.AuxInt
			_ = v_0.Args[1]
			x := v_0.Args[0]
			y := v_0.Args[1]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftRA, psess.types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftLLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftLLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftRLreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftRLreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMCMPconst {
				break
			}
			if v.AuxInt != 0 {
				break
			}
			v_0 := v.Args[0]
			if v_0.Op != OpARMXORshiftRAreg {
				break
			}
			_ = v_0.Args[2]
			x := v_0.Args[0]
			y := v_0.Args[1]
			z := v_0.Args[2]
			b.Kind = BlockARMNE
			v0 := b.NewValue0(v.Pos, OpARMTEQshiftRAreg, psess.types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.SetControl(v0)
			b.Aux = nil
			return true
		}
	case BlockARMUGE:

		for {
			v := b.Control
			if v.Op != OpARMFlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagLT_ULT {
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
			if v.Op != OpARMFlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagGT_ULT {
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
			if v.Op != OpARMFlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARMULE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARMUGT:

		for {
			v := b.Control
			if v.Op != OpARMFlagEQ {
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
			if v.Op != OpARMFlagLT_ULT {
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
			if v.Op != OpARMFlagLT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagGT_ULT {
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
			if v.Op != OpARMFlagGT_UGT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARMULT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARMULE:

		for {
			v := b.Control
			if v.Op != OpARMFlagEQ {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagLT_UGT {
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
			if v.Op != OpARMFlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagGT_UGT {
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
			if v.Op != OpARMInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARMUGE
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	case BlockARMULT:

		for {
			v := b.Control
			if v.Op != OpARMFlagEQ {
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
			if v.Op != OpARMFlagLT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagLT_UGT {
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
			if v.Op != OpARMFlagGT_ULT {
				break
			}
			b.Kind = BlockFirst
			b.SetControl(nil)
			b.Aux = nil
			return true
		}

		for {
			v := b.Control
			if v.Op != OpARMFlagGT_UGT {
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
			if v.Op != OpARMInvertFlags {
				break
			}
			cmp := v.Args[0]
			b.Kind = BlockARMUGT
			b.SetControl(cmp)
			b.Aux = nil
			return true
		}
	}
	return false
}

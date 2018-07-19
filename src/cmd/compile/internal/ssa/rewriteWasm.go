package ssa

import "github.com/dave/golib/src/cmd/compile/internal/types"

// in case not otherwise used
// in case not otherwise used
// in case not otherwise used
// in case not otherwise used

func (psess *PackageSession) rewriteValueWasm(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValueWasm_OpAdd16_0(v)
	case OpAdd32:
		return rewriteValueWasm_OpAdd32_0(v)
	case OpAdd32F:
		return rewriteValueWasm_OpAdd32F_0(v)
	case OpAdd64:
		return rewriteValueWasm_OpAdd64_0(v)
	case OpAdd64F:
		return rewriteValueWasm_OpAdd64F_0(v)
	case OpAdd8:
		return rewriteValueWasm_OpAdd8_0(v)
	case OpAddPtr:
		return rewriteValueWasm_OpAddPtr_0(v)
	case OpAddr:
		return rewriteValueWasm_OpAddr_0(v)
	case OpAnd16:
		return rewriteValueWasm_OpAnd16_0(v)
	case OpAnd32:
		return rewriteValueWasm_OpAnd32_0(v)
	case OpAnd64:
		return rewriteValueWasm_OpAnd64_0(v)
	case OpAnd8:
		return rewriteValueWasm_OpAnd8_0(v)
	case OpAndB:
		return rewriteValueWasm_OpAndB_0(v)
	case OpClosureCall:
		return rewriteValueWasm_OpClosureCall_0(v)
	case OpCom16:
		return rewriteValueWasm_OpCom16_0(v)
	case OpCom32:
		return rewriteValueWasm_OpCom32_0(v)
	case OpCom64:
		return rewriteValueWasm_OpCom64_0(v)
	case OpCom8:
		return rewriteValueWasm_OpCom8_0(v)
	case OpConst16:
		return rewriteValueWasm_OpConst16_0(v)
	case OpConst32:
		return rewriteValueWasm_OpConst32_0(v)
	case OpConst32F:
		return rewriteValueWasm_OpConst32F_0(v)
	case OpConst64:
		return rewriteValueWasm_OpConst64_0(v)
	case OpConst64F:
		return rewriteValueWasm_OpConst64F_0(v)
	case OpConst8:
		return rewriteValueWasm_OpConst8_0(v)
	case OpConstBool:
		return rewriteValueWasm_OpConstBool_0(v)
	case OpConstNil:
		return rewriteValueWasm_OpConstNil_0(v)
	case OpConvert:
		return rewriteValueWasm_OpConvert_0(v)
	case OpCvt32Fto32:
		return rewriteValueWasm_OpCvt32Fto32_0(v)
	case OpCvt32Fto32U:
		return rewriteValueWasm_OpCvt32Fto32U_0(v)
	case OpCvt32Fto64:
		return rewriteValueWasm_OpCvt32Fto64_0(v)
	case OpCvt32Fto64F:
		return rewriteValueWasm_OpCvt32Fto64F_0(v)
	case OpCvt32Fto64U:
		return rewriteValueWasm_OpCvt32Fto64U_0(v)
	case OpCvt32Uto32F:
		return rewriteValueWasm_OpCvt32Uto32F_0(v)
	case OpCvt32Uto64F:
		return rewriteValueWasm_OpCvt32Uto64F_0(v)
	case OpCvt32to32F:
		return rewriteValueWasm_OpCvt32to32F_0(v)
	case OpCvt32to64F:
		return rewriteValueWasm_OpCvt32to64F_0(v)
	case OpCvt64Fto32:
		return rewriteValueWasm_OpCvt64Fto32_0(v)
	case OpCvt64Fto32F:
		return rewriteValueWasm_OpCvt64Fto32F_0(v)
	case OpCvt64Fto32U:
		return rewriteValueWasm_OpCvt64Fto32U_0(v)
	case OpCvt64Fto64:
		return rewriteValueWasm_OpCvt64Fto64_0(v)
	case OpCvt64Fto64U:
		return rewriteValueWasm_OpCvt64Fto64U_0(v)
	case OpCvt64Uto32F:
		return rewriteValueWasm_OpCvt64Uto32F_0(v)
	case OpCvt64Uto64F:
		return rewriteValueWasm_OpCvt64Uto64F_0(v)
	case OpCvt64to32F:
		return rewriteValueWasm_OpCvt64to32F_0(v)
	case OpCvt64to64F:
		return rewriteValueWasm_OpCvt64to64F_0(v)
	case OpDiv16:
		return rewriteValueWasm_OpDiv16_0(v)
	case OpDiv16u:
		return rewriteValueWasm_OpDiv16u_0(v)
	case OpDiv32:
		return rewriteValueWasm_OpDiv32_0(v)
	case OpDiv32F:
		return rewriteValueWasm_OpDiv32F_0(v)
	case OpDiv32u:
		return rewriteValueWasm_OpDiv32u_0(v)
	case OpDiv64:
		return rewriteValueWasm_OpDiv64_0(v)
	case OpDiv64F:
		return rewriteValueWasm_OpDiv64F_0(v)
	case OpDiv64u:
		return rewriteValueWasm_OpDiv64u_0(v)
	case OpDiv8:
		return rewriteValueWasm_OpDiv8_0(v)
	case OpDiv8u:
		return rewriteValueWasm_OpDiv8u_0(v)
	case OpEq16:
		return rewriteValueWasm_OpEq16_0(v)
	case OpEq32:
		return rewriteValueWasm_OpEq32_0(v)
	case OpEq32F:
		return rewriteValueWasm_OpEq32F_0(v)
	case OpEq64:
		return rewriteValueWasm_OpEq64_0(v)
	case OpEq64F:
		return rewriteValueWasm_OpEq64F_0(v)
	case OpEq8:
		return rewriteValueWasm_OpEq8_0(v)
	case OpEqB:
		return rewriteValueWasm_OpEqB_0(v)
	case OpEqPtr:
		return rewriteValueWasm_OpEqPtr_0(v)
	case OpGeq16:
		return rewriteValueWasm_OpGeq16_0(v)
	case OpGeq16U:
		return rewriteValueWasm_OpGeq16U_0(v)
	case OpGeq32:
		return rewriteValueWasm_OpGeq32_0(v)
	case OpGeq32F:
		return rewriteValueWasm_OpGeq32F_0(v)
	case OpGeq32U:
		return rewriteValueWasm_OpGeq32U_0(v)
	case OpGeq64:
		return rewriteValueWasm_OpGeq64_0(v)
	case OpGeq64F:
		return rewriteValueWasm_OpGeq64F_0(v)
	case OpGeq64U:
		return rewriteValueWasm_OpGeq64U_0(v)
	case OpGeq8:
		return rewriteValueWasm_OpGeq8_0(v)
	case OpGeq8U:
		return rewriteValueWasm_OpGeq8U_0(v)
	case OpGetCallerPC:
		return rewriteValueWasm_OpGetCallerPC_0(v)
	case OpGetCallerSP:
		return rewriteValueWasm_OpGetCallerSP_0(v)
	case OpGetClosurePtr:
		return rewriteValueWasm_OpGetClosurePtr_0(v)
	case OpGreater16:
		return rewriteValueWasm_OpGreater16_0(v)
	case OpGreater16U:
		return rewriteValueWasm_OpGreater16U_0(v)
	case OpGreater32:
		return rewriteValueWasm_OpGreater32_0(v)
	case OpGreater32F:
		return rewriteValueWasm_OpGreater32F_0(v)
	case OpGreater32U:
		return rewriteValueWasm_OpGreater32U_0(v)
	case OpGreater64:
		return rewriteValueWasm_OpGreater64_0(v)
	case OpGreater64F:
		return rewriteValueWasm_OpGreater64F_0(v)
	case OpGreater64U:
		return rewriteValueWasm_OpGreater64U_0(v)
	case OpGreater8:
		return rewriteValueWasm_OpGreater8_0(v)
	case OpGreater8U:
		return rewriteValueWasm_OpGreater8U_0(v)
	case OpInterCall:
		return rewriteValueWasm_OpInterCall_0(v)
	case OpIsInBounds:
		return rewriteValueWasm_OpIsInBounds_0(v)
	case OpIsNonNil:
		return rewriteValueWasm_OpIsNonNil_0(v)
	case OpIsSliceInBounds:
		return rewriteValueWasm_OpIsSliceInBounds_0(v)
	case OpLeq16:
		return rewriteValueWasm_OpLeq16_0(v)
	case OpLeq16U:
		return rewriteValueWasm_OpLeq16U_0(v)
	case OpLeq32:
		return rewriteValueWasm_OpLeq32_0(v)
	case OpLeq32F:
		return rewriteValueWasm_OpLeq32F_0(v)
	case OpLeq32U:
		return rewriteValueWasm_OpLeq32U_0(v)
	case OpLeq64:
		return rewriteValueWasm_OpLeq64_0(v)
	case OpLeq64F:
		return rewriteValueWasm_OpLeq64F_0(v)
	case OpLeq64U:
		return rewriteValueWasm_OpLeq64U_0(v)
	case OpLeq8:
		return rewriteValueWasm_OpLeq8_0(v)
	case OpLeq8U:
		return rewriteValueWasm_OpLeq8U_0(v)
	case OpLess16:
		return rewriteValueWasm_OpLess16_0(v)
	case OpLess16U:
		return rewriteValueWasm_OpLess16U_0(v)
	case OpLess32:
		return rewriteValueWasm_OpLess32_0(v)
	case OpLess32F:
		return rewriteValueWasm_OpLess32F_0(v)
	case OpLess32U:
		return rewriteValueWasm_OpLess32U_0(v)
	case OpLess64:
		return rewriteValueWasm_OpLess64_0(v)
	case OpLess64F:
		return rewriteValueWasm_OpLess64F_0(v)
	case OpLess64U:
		return rewriteValueWasm_OpLess64U_0(v)
	case OpLess8:
		return rewriteValueWasm_OpLess8_0(v)
	case OpLess8U:
		return rewriteValueWasm_OpLess8U_0(v)
	case OpLoad:
		return psess.rewriteValueWasm_OpLoad_0(v)
	case OpLsh16x16:
		return rewriteValueWasm_OpLsh16x16_0(v)
	case OpLsh16x32:
		return rewriteValueWasm_OpLsh16x32_0(v)
	case OpLsh16x64:
		return rewriteValueWasm_OpLsh16x64_0(v)
	case OpLsh16x8:
		return rewriteValueWasm_OpLsh16x8_0(v)
	case OpLsh32x16:
		return rewriteValueWasm_OpLsh32x16_0(v)
	case OpLsh32x32:
		return rewriteValueWasm_OpLsh32x32_0(v)
	case OpLsh32x64:
		return rewriteValueWasm_OpLsh32x64_0(v)
	case OpLsh32x8:
		return rewriteValueWasm_OpLsh32x8_0(v)
	case OpLsh64x16:
		return rewriteValueWasm_OpLsh64x16_0(v)
	case OpLsh64x32:
		return rewriteValueWasm_OpLsh64x32_0(v)
	case OpLsh64x64:
		return rewriteValueWasm_OpLsh64x64_0(v)
	case OpLsh64x8:
		return rewriteValueWasm_OpLsh64x8_0(v)
	case OpLsh8x16:
		return rewriteValueWasm_OpLsh8x16_0(v)
	case OpLsh8x32:
		return rewriteValueWasm_OpLsh8x32_0(v)
	case OpLsh8x64:
		return rewriteValueWasm_OpLsh8x64_0(v)
	case OpLsh8x8:
		return rewriteValueWasm_OpLsh8x8_0(v)
	case OpMod16:
		return rewriteValueWasm_OpMod16_0(v)
	case OpMod16u:
		return rewriteValueWasm_OpMod16u_0(v)
	case OpMod32:
		return rewriteValueWasm_OpMod32_0(v)
	case OpMod32u:
		return rewriteValueWasm_OpMod32u_0(v)
	case OpMod64:
		return rewriteValueWasm_OpMod64_0(v)
	case OpMod64u:
		return rewriteValueWasm_OpMod64u_0(v)
	case OpMod8:
		return rewriteValueWasm_OpMod8_0(v)
	case OpMod8u:
		return rewriteValueWasm_OpMod8u_0(v)
	case OpMove:
		return psess.rewriteValueWasm_OpMove_0(v) || psess.rewriteValueWasm_OpMove_10(v)
	case OpMul16:
		return rewriteValueWasm_OpMul16_0(v)
	case OpMul32:
		return rewriteValueWasm_OpMul32_0(v)
	case OpMul32F:
		return rewriteValueWasm_OpMul32F_0(v)
	case OpMul64:
		return rewriteValueWasm_OpMul64_0(v)
	case OpMul64F:
		return rewriteValueWasm_OpMul64F_0(v)
	case OpMul8:
		return rewriteValueWasm_OpMul8_0(v)
	case OpNeg16:
		return rewriteValueWasm_OpNeg16_0(v)
	case OpNeg32:
		return rewriteValueWasm_OpNeg32_0(v)
	case OpNeg32F:
		return rewriteValueWasm_OpNeg32F_0(v)
	case OpNeg64:
		return rewriteValueWasm_OpNeg64_0(v)
	case OpNeg64F:
		return rewriteValueWasm_OpNeg64F_0(v)
	case OpNeg8:
		return rewriteValueWasm_OpNeg8_0(v)
	case OpNeq16:
		return rewriteValueWasm_OpNeq16_0(v)
	case OpNeq32:
		return rewriteValueWasm_OpNeq32_0(v)
	case OpNeq32F:
		return rewriteValueWasm_OpNeq32F_0(v)
	case OpNeq64:
		return rewriteValueWasm_OpNeq64_0(v)
	case OpNeq64F:
		return rewriteValueWasm_OpNeq64F_0(v)
	case OpNeq8:
		return rewriteValueWasm_OpNeq8_0(v)
	case OpNeqB:
		return rewriteValueWasm_OpNeqB_0(v)
	case OpNeqPtr:
		return rewriteValueWasm_OpNeqPtr_0(v)
	case OpNilCheck:
		return rewriteValueWasm_OpNilCheck_0(v)
	case OpNot:
		return rewriteValueWasm_OpNot_0(v)
	case OpOffPtr:
		return rewriteValueWasm_OpOffPtr_0(v)
	case OpOr16:
		return rewriteValueWasm_OpOr16_0(v)
	case OpOr32:
		return rewriteValueWasm_OpOr32_0(v)
	case OpOr64:
		return rewriteValueWasm_OpOr64_0(v)
	case OpOr8:
		return rewriteValueWasm_OpOr8_0(v)
	case OpOrB:
		return rewriteValueWasm_OpOrB_0(v)
	case OpRound32F:
		return rewriteValueWasm_OpRound32F_0(v)
	case OpRound64F:
		return rewriteValueWasm_OpRound64F_0(v)
	case OpRsh16Ux16:
		return rewriteValueWasm_OpRsh16Ux16_0(v)
	case OpRsh16Ux32:
		return rewriteValueWasm_OpRsh16Ux32_0(v)
	case OpRsh16Ux64:
		return rewriteValueWasm_OpRsh16Ux64_0(v)
	case OpRsh16Ux8:
		return rewriteValueWasm_OpRsh16Ux8_0(v)
	case OpRsh16x16:
		return rewriteValueWasm_OpRsh16x16_0(v)
	case OpRsh16x32:
		return rewriteValueWasm_OpRsh16x32_0(v)
	case OpRsh16x64:
		return rewriteValueWasm_OpRsh16x64_0(v)
	case OpRsh16x8:
		return rewriteValueWasm_OpRsh16x8_0(v)
	case OpRsh32Ux16:
		return rewriteValueWasm_OpRsh32Ux16_0(v)
	case OpRsh32Ux32:
		return rewriteValueWasm_OpRsh32Ux32_0(v)
	case OpRsh32Ux64:
		return rewriteValueWasm_OpRsh32Ux64_0(v)
	case OpRsh32Ux8:
		return rewriteValueWasm_OpRsh32Ux8_0(v)
	case OpRsh32x16:
		return rewriteValueWasm_OpRsh32x16_0(v)
	case OpRsh32x32:
		return rewriteValueWasm_OpRsh32x32_0(v)
	case OpRsh32x64:
		return rewriteValueWasm_OpRsh32x64_0(v)
	case OpRsh32x8:
		return rewriteValueWasm_OpRsh32x8_0(v)
	case OpRsh64Ux16:
		return rewriteValueWasm_OpRsh64Ux16_0(v)
	case OpRsh64Ux32:
		return rewriteValueWasm_OpRsh64Ux32_0(v)
	case OpRsh64Ux64:
		return rewriteValueWasm_OpRsh64Ux64_0(v)
	case OpRsh64Ux8:
		return rewriteValueWasm_OpRsh64Ux8_0(v)
	case OpRsh64x16:
		return rewriteValueWasm_OpRsh64x16_0(v)
	case OpRsh64x32:
		return rewriteValueWasm_OpRsh64x32_0(v)
	case OpRsh64x64:
		return rewriteValueWasm_OpRsh64x64_0(v)
	case OpRsh64x8:
		return rewriteValueWasm_OpRsh64x8_0(v)
	case OpRsh8Ux16:
		return rewriteValueWasm_OpRsh8Ux16_0(v)
	case OpRsh8Ux32:
		return rewriteValueWasm_OpRsh8Ux32_0(v)
	case OpRsh8Ux64:
		return rewriteValueWasm_OpRsh8Ux64_0(v)
	case OpRsh8Ux8:
		return rewriteValueWasm_OpRsh8Ux8_0(v)
	case OpRsh8x16:
		return rewriteValueWasm_OpRsh8x16_0(v)
	case OpRsh8x32:
		return rewriteValueWasm_OpRsh8x32_0(v)
	case OpRsh8x64:
		return rewriteValueWasm_OpRsh8x64_0(v)
	case OpRsh8x8:
		return rewriteValueWasm_OpRsh8x8_0(v)
	case OpSignExt16to32:
		return rewriteValueWasm_OpSignExt16to32_0(v)
	case OpSignExt16to64:
		return rewriteValueWasm_OpSignExt16to64_0(v)
	case OpSignExt32to64:
		return rewriteValueWasm_OpSignExt32to64_0(v)
	case OpSignExt8to16:
		return rewriteValueWasm_OpSignExt8to16_0(v)
	case OpSignExt8to32:
		return rewriteValueWasm_OpSignExt8to32_0(v)
	case OpSignExt8to64:
		return rewriteValueWasm_OpSignExt8to64_0(v)
	case OpSlicemask:
		return rewriteValueWasm_OpSlicemask_0(v)
	case OpStaticCall:
		return rewriteValueWasm_OpStaticCall_0(v)
	case OpStore:
		return psess.rewriteValueWasm_OpStore_0(v)
	case OpSub16:
		return rewriteValueWasm_OpSub16_0(v)
	case OpSub32:
		return rewriteValueWasm_OpSub32_0(v)
	case OpSub32F:
		return rewriteValueWasm_OpSub32F_0(v)
	case OpSub64:
		return rewriteValueWasm_OpSub64_0(v)
	case OpSub64F:
		return rewriteValueWasm_OpSub64F_0(v)
	case OpSub8:
		return rewriteValueWasm_OpSub8_0(v)
	case OpSubPtr:
		return rewriteValueWasm_OpSubPtr_0(v)
	case OpTrunc16to8:
		return rewriteValueWasm_OpTrunc16to8_0(v)
	case OpTrunc32to16:
		return rewriteValueWasm_OpTrunc32to16_0(v)
	case OpTrunc32to8:
		return rewriteValueWasm_OpTrunc32to8_0(v)
	case OpTrunc64to16:
		return rewriteValueWasm_OpTrunc64to16_0(v)
	case OpTrunc64to32:
		return rewriteValueWasm_OpTrunc64to32_0(v)
	case OpTrunc64to8:
		return rewriteValueWasm_OpTrunc64to8_0(v)
	case OpWB:
		return rewriteValueWasm_OpWB_0(v)
	case OpWasmF64Add:
		return rewriteValueWasm_OpWasmF64Add_0(v)
	case OpWasmF64Mul:
		return rewriteValueWasm_OpWasmF64Mul_0(v)
	case OpWasmI64Add:
		return rewriteValueWasm_OpWasmI64Add_0(v)
	case OpWasmI64AddConst:
		return rewriteValueWasm_OpWasmI64AddConst_0(v)
	case OpWasmI64And:
		return rewriteValueWasm_OpWasmI64And_0(v)
	case OpWasmI64Eq:
		return rewriteValueWasm_OpWasmI64Eq_0(v)
	case OpWasmI64Eqz:
		return rewriteValueWasm_OpWasmI64Eqz_0(v)
	case OpWasmI64Load:
		return rewriteValueWasm_OpWasmI64Load_0(v)
	case OpWasmI64Load16S:
		return rewriteValueWasm_OpWasmI64Load16S_0(v)
	case OpWasmI64Load16U:
		return rewriteValueWasm_OpWasmI64Load16U_0(v)
	case OpWasmI64Load32S:
		return rewriteValueWasm_OpWasmI64Load32S_0(v)
	case OpWasmI64Load32U:
		return rewriteValueWasm_OpWasmI64Load32U_0(v)
	case OpWasmI64Load8S:
		return rewriteValueWasm_OpWasmI64Load8S_0(v)
	case OpWasmI64Load8U:
		return rewriteValueWasm_OpWasmI64Load8U_0(v)
	case OpWasmI64Mul:
		return rewriteValueWasm_OpWasmI64Mul_0(v)
	case OpWasmI64Ne:
		return rewriteValueWasm_OpWasmI64Ne_0(v)
	case OpWasmI64Or:
		return rewriteValueWasm_OpWasmI64Or_0(v)
	case OpWasmI64Shl:
		return rewriteValueWasm_OpWasmI64Shl_0(v)
	case OpWasmI64ShrS:
		return rewriteValueWasm_OpWasmI64ShrS_0(v)
	case OpWasmI64ShrU:
		return rewriteValueWasm_OpWasmI64ShrU_0(v)
	case OpWasmI64Store:
		return rewriteValueWasm_OpWasmI64Store_0(v)
	case OpWasmI64Store16:
		return rewriteValueWasm_OpWasmI64Store16_0(v)
	case OpWasmI64Store32:
		return rewriteValueWasm_OpWasmI64Store32_0(v)
	case OpWasmI64Store8:
		return rewriteValueWasm_OpWasmI64Store8_0(v)
	case OpWasmI64Xor:
		return rewriteValueWasm_OpWasmI64Xor_0(v)
	case OpXor16:
		return rewriteValueWasm_OpXor16_0(v)
	case OpXor32:
		return rewriteValueWasm_OpXor32_0(v)
	case OpXor64:
		return rewriteValueWasm_OpXor64_0(v)
	case OpXor8:
		return rewriteValueWasm_OpXor8_0(v)
	case OpZero:
		return psess.rewriteValueWasm_OpZero_0(v) || psess.rewriteValueWasm_OpZero_10(v)
	case OpZeroExt16to32:
		return rewriteValueWasm_OpZeroExt16to32_0(v)
	case OpZeroExt16to64:
		return rewriteValueWasm_OpZeroExt16to64_0(v)
	case OpZeroExt32to64:
		return rewriteValueWasm_OpZeroExt32to64_0(v)
	case OpZeroExt8to16:
		return rewriteValueWasm_OpZeroExt8to16_0(v)
	case OpZeroExt8to32:
		return rewriteValueWasm_OpZeroExt8to32_0(v)
	case OpZeroExt8to64:
		return rewriteValueWasm_OpZeroExt8to64_0(v)
	}
	return false
}
func rewriteValueWasm_OpAdd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Add)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpAdd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Add)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpAdd32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Add)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpAdd64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Add)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpAdd64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Add)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpAdd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Add)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpAddPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Add)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpAddr_0(v *Value) bool {

	for {
		sym := v.Aux
		base := v.Args[0]
		v.reset(OpWasmLoweredAddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueWasm_OpAnd16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64And)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpAnd32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64And)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpAnd64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64And)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpAnd8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64And)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpAndB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64And)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpClosureCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[2]
		entry := v.Args[0]
		closure := v.Args[1]
		mem := v.Args[2]
		v.reset(OpWasmLoweredClosureCall)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueWasm_OpCom16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmI64Xor)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = -1
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCom32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmI64Xor)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = -1
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCom64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmI64Xor)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = -1
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCom8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmI64Xor)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = -1
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpConst16_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = val
		return true
	}
}
func rewriteValueWasm_OpConst32_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = val
		return true
	}
}
func rewriteValueWasm_OpConst32F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpWasmF64Const)
		v.AuxInt = val
		return true
	}
}
func rewriteValueWasm_OpConst64_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = val
		return true
	}
}
func rewriteValueWasm_OpConst64F_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpWasmF64Const)
		v.AuxInt = val
		return true
	}
}
func rewriteValueWasm_OpConst8_0(v *Value) bool {

	for {
		val := v.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = val
		return true
	}
}
func rewriteValueWasm_OpConstBool_0(v *Value) bool {

	for {
		b := v.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = b
		return true
	}
}
func rewriteValueWasm_OpConstNil_0(v *Value) bool {

	for {
		v.reset(OpWasmI64Const)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueWasm_OpConvert_0(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		x := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmLoweredConvert)
		v.Type = t
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueWasm_OpCvt32Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmI64TruncSF64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt32Fto32U_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmI64TruncUF64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt32Fto64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmI64TruncSF64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt32Fto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt32Fto64U_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmI64TruncUF64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt32Uto32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmLoweredRound32F)
		v0 := b.NewValue0(v.Pos, OpWasmF64ConvertUI64, typ.Float64)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCvt32Uto64F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmF64ConvertUI64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCvt32to32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmLoweredRound32F)
		v0 := b.NewValue0(v.Pos, OpWasmF64ConvertSI64, typ.Float64)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCvt32to64F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmF64ConvertSI64)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCvt64Fto32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmI64TruncSF64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt64Fto32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmLoweredRound32F)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt64Fto32U_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmI64TruncUF64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt64Fto64_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmI64TruncSF64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt64Fto64U_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmI64TruncUF64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt64Uto32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmLoweredRound32F)
		v0 := b.NewValue0(v.Pos, OpWasmF64ConvertUI64, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCvt64Uto64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmF64ConvertUI64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpCvt64to32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmLoweredRound32F)
		v0 := b.NewValue0(v.Pos, OpWasmF64ConvertSI64, typ.Float64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpCvt64to64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmF64ConvertSI64)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpDiv16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64DivS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpDiv16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64DivU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpDiv32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64DivS)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpDiv32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Div)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpDiv32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64DivU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpDiv64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64DivS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpDiv64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Div)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpDiv64u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64DivU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpDiv8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64DivS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpDiv8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64DivU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpEq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Eq)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpEq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Eq)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpEq32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Eq)
		v0 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpEq64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Eq)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpEq64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Eq)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpEq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Eq)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpEqB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Eq)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpEqPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Eq)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpGeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GeS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GeU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GeS)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGeq32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Ge)
		v0 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGeq32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GeU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGeq64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GeS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpGeq64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Ge)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpGeq64U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GeU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpGeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GeS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GeU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGetCallerPC_0(v *Value) bool {

	for {
		v.reset(OpWasmLoweredGetCallerPC)
		return true
	}
}
func rewriteValueWasm_OpGetCallerSP_0(v *Value) bool {

	for {
		v.reset(OpWasmLoweredGetCallerSP)
		return true
	}
}
func rewriteValueWasm_OpGetClosurePtr_0(v *Value) bool {

	for {
		v.reset(OpWasmLoweredGetClosurePtr)
		return true
	}
}
func rewriteValueWasm_OpGreater16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GtS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGreater16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GtU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGreater32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GtS)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGreater32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Gt)
		v0 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGreater32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GtU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGreater64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GtS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpGreater64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Gt)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpGreater64U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GtU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpGreater8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GtS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpGreater8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64GtU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpInterCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		_ = v.Args[1]
		entry := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmLoweredInterCall)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueWasm_OpIsInBounds_0(v *Value) bool {

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpWasmI64LtU)
		v.AddArg(idx)
		v.AddArg(len)
		return true
	}
}
func rewriteValueWasm_OpIsNonNil_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		p := v.Args[0]
		v.reset(OpWasmI64Eqz)
		v0 := b.NewValue0(v.Pos, OpWasmI64Eqz, typ.Bool)
		v0.AddArg(p)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpIsSliceInBounds_0(v *Value) bool {

	for {
		_ = v.Args[1]
		idx := v.Args[0]
		len := v.Args[1]
		v.reset(OpWasmI64LeU)
		v.AddArg(idx)
		v.AddArg(len)
		return true
	}
}
func rewriteValueWasm_OpLeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LeS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLeq16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LeU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LeS)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLeq32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Le)
		v0 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLeq32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LeU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLeq64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LeS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpLeq64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Le)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpLeq64U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LeU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpLeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LeS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLeq8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LeU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLess16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LtS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLess16U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LtU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLess32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LtS)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLess32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Lt)
		v0 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLess32U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LtU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLess64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LtS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpLess64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Lt)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpLess64U_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LtU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpLess8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LtS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpLess8U_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64LtU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueWasm_OpLoad_0(v *Value) bool {

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(psess.is32BitFloat(t)) {
			break
		}
		v.reset(OpWasmF32Load)
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
		v.reset(OpWasmF64Load)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size(psess.types) == 8) {
			break
		}
		v.reset(OpWasmI64Load)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size(psess.types) == 4 && !t.IsSigned()) {
			break
		}
		v.reset(OpWasmI64Load32U)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size(psess.types) == 4 && t.IsSigned()) {
			break
		}
		v.reset(OpWasmI64Load32S)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size(psess.types) == 2 && !t.IsSigned()) {
			break
		}
		v.reset(OpWasmI64Load16U)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size(psess.types) == 2 && t.IsSigned()) {
			break
		}
		v.reset(OpWasmI64Load16S)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size(psess.types) == 1 && !t.IsSigned()) {
			break
		}
		v.reset(OpWasmI64Load8U)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}

	for {
		t := v.Type
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.Size(psess.types) == 1 && t.IsSigned()) {
			break
		}
		v.reset(OpWasmI64Load8S)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpLsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLsh16x64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpLsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLsh32x64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpLsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmSelect)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpWasmI64LtU, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v3.AuxInt = 64
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpLsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpLsh8x64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpLsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpLsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpMod16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64RemS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpMod16u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64RemU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpMod32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64RemS)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpMod32u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64RemU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpMod64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64RemS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpMod64u_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64RemU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpMod8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64RemS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpMod8u_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64RemU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func (psess *PackageSession) rewriteValueWasm_OpMove_0(v *Value) bool {
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
		v.reset(OpWasmI64Store8)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpWasmI64Load8U, typ.UInt8)
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
		v.reset(OpWasmI64Store16)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpWasmI64Load16U, typ.UInt16)
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
		v.reset(OpWasmI64Store32)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpWasmI64Load32U, typ.UInt32)
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
		v.reset(OpWasmI64Store)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpWasmI64Load, typ.UInt64)
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
		v.reset(OpWasmI64Store)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpWasmI64Load, typ.UInt64)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpWasmI64Load, typ.UInt64)
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
		v.reset(OpWasmI64Store8)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpWasmI64Load8U, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store16, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpWasmI64Load16U, typ.UInt16)
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
		v.reset(OpWasmI64Store8)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpWasmI64Load8U, typ.UInt8)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store32, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpWasmI64Load32U, typ.UInt32)
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
		v.reset(OpWasmI64Store16)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpWasmI64Load16U, typ.UInt16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store32, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpWasmI64Load32U, typ.UInt32)
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
		v.reset(OpWasmI64Store32)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpWasmI64Load32U, typ.UInt32)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store32, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpWasmI64Load32U, typ.UInt32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueWasm_OpMove_10(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		s := v.AuxInt
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s > 8 && s < 16) {
			break
		}
		v.reset(OpWasmI64Store)
		v.AuxInt = s - 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpWasmI64Load, typ.UInt64)
		v0.AuxInt = s - 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpWasmI64Load, typ.UInt64)
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
		if !(s > 16 && s%16 != 0 && s%16 <= 8) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = s - s%16
		v0 := b.NewValue0(v.Pos, OpOffPtr, dst.Type)
		v0.AuxInt = s % 16
		v0.AddArg(dst)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, src.Type)
		v1.AuxInt = s % 16
		v1.AddArg(src)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpWasmI64Load, typ.UInt64)
		v3.AddArg(src)
		v3.AddArg(mem)
		v2.AddArg(v3)
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
		if !(s > 16 && s%16 != 0 && s%16 > 8) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = s - s%16
		v0 := b.NewValue0(v.Pos, OpOffPtr, dst.Type)
		v0.AuxInt = s % 16
		v0.AddArg(dst)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, src.Type)
		v1.AuxInt = s % 16
		v1.AddArg(src)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v2.AuxInt = 8
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpWasmI64Load, typ.UInt64)
		v3.AuxInt = 8
		v3.AddArg(src)
		v3.AddArg(mem)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v4.AddArg(dst)
		v5 := b.NewValue0(v.Pos, OpWasmI64Load, typ.UInt64)
		v5.AddArg(src)
		v5.AddArg(mem)
		v4.AddArg(v5)
		v4.AddArg(mem)
		v2.AddArg(v4)
		v.AddArg(v2)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[2]
		dst := v.Args[0]
		src := v.Args[1]
		mem := v.Args[2]
		if !(s%8 == 0) {
			break
		}
		v.reset(OpWasmLoweredMove)
		v.AuxInt = s / 8
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpMul16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Mul)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpMul32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Mul)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpMul32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Mul)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpMul64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Mul)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpMul64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Mul)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpMul8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Mul)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpNeg16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmI64Sub)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpNeg32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmI64Sub)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpNeg32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmF64Neg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpNeg64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmI64Sub)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpNeg64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmF64Neg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpNeg8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmI64Sub)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpNeq16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Ne)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpNeq32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Ne)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpNeq32F_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Ne)
		v0 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmLoweredRound32F, typ.Float32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpNeq64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Ne)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpNeq64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Ne)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpNeq8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Ne)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpNeqB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Ne)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpNeqPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Ne)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpNilCheck_0(v *Value) bool {

	for {
		_ = v.Args[1]
		ptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmLoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueWasm_OpNot_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmI64Eqz)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpOffPtr_0(v *Value) bool {

	for {
		off := v.AuxInt
		ptr := v.Args[0]
		v.reset(OpWasmI64AddConst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueWasm_OpOr16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Or)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpOr32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Or)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpOr64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Or)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpOr8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Or)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpOrB_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Or)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpRound32F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpWasmLoweredRound32F)
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpRound64F_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpRsh16Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh16Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh16Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpRsh16Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh16x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh16x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh16x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpRsh16x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh32Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh32Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh32Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpRsh32Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh32x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh32x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh32x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpRsh32x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh64Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmSelect)
		v0 := b.NewValue0(v.Pos, OpWasmI64ShrU, typ.Int64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 0
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpWasmI64LtU, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v3.AuxInt = 64
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpRsh64Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64ShrS)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpWasmSelect, typ.Int64)
		v0.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 63
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpWasmI64LtU, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v3.AuxInt = 64
		v2.AddArg(v3)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpRsh64x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm_OpRsh8Ux16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh8Ux32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh8Ux64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpRsh8Ux8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh8x16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh8x32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpRsh8x64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpRsh8x8_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueWasm_OpSignExt16to32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load16S {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 48
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 48
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpSignExt16to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load16S {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 48
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 48
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpSignExt32to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load32S {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 32
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 32
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpSignExt8to16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load8S {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 56
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 56
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpSignExt8to32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load8S {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 56
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 56
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpSignExt8to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load8S {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 56
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 56
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpSlicemask_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasmI64Sub, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 0
		v0.AddArg(v1)
		v0.AddArg(x)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 63
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpStaticCall_0(v *Value) bool {

	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v.Args[0]
		v.reset(OpWasmLoweredStaticCall)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func (psess *PackageSession) rewriteValueWasm_OpStore_0(v *Value) bool {

	for {
		t := v.Aux
		_ = v.Args[2]
		ptr := v.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(psess.is64BitFloat(t.(*types.Type))) {
			break
		}
		v.reset(OpWasmF64Store)
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
		if !(psess.is32BitFloat(t.(*types.Type))) {
			break
		}
		v.reset(OpWasmF32Store)
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
		v.reset(OpWasmI64Store)
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
		v.reset(OpWasmI64Store32)
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
		v.reset(OpWasmI64Store16)
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
		v.reset(OpWasmI64Store8)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpSub16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Sub)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpSub32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Sub)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpSub32F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Sub)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpSub64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Sub)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpSub64F_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmF64Sub)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpSub8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Sub)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpSubPtr_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Sub)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpTrunc16to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpTrunc32to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpTrunc32to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpTrunc64to16_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpTrunc64to32_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpTrunc64to8_0(v *Value) bool {

	for {
		x := v.Args[0]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueWasm_OpWB_0(v *Value) bool {

	for {
		fn := v.Aux
		_ = v.Args[2]
		destptr := v.Args[0]
		srcptr := v.Args[1]
		mem := v.Args[2]
		v.reset(OpWasmLoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueWasm_OpWasmF64Add_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmF64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmF64Const {
			break
		}
		y := v_1.AuxInt
		v.reset(OpWasmF64Const)
		v.AuxInt = f2i(i2f(x) + i2f(y))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmF64Const {
			break
		}
		x := v_0.AuxInt
		y := v.Args[1]
		v.reset(OpWasmF64Add)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpWasmF64Const, typ.Float64)
		v0.AuxInt = x
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmF64Mul_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmF64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmF64Const {
			break
		}
		y := v_1.AuxInt
		v.reset(OpWasmF64Const)
		v.AuxInt = f2i(i2f(x) * i2f(y))
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmF64Const {
			break
		}
		x := v_0.AuxInt
		y := v.Args[1]
		v.reset(OpWasmF64Mul)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpWasmF64Const, typ.Float64)
		v0.AuxInt = x
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Add_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = x + y
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		y := v.Args[1]
		v.reset(OpWasmI64Add)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = x
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		v.reset(OpWasmI64AddConst)
		v.AuxInt = y
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64AddConst_0(v *Value) bool {

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
func rewriteValueWasm_OpWasmI64And_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = x & y
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		y := v.Args[1]
		v.reset(OpWasmI64And)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = x
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Eq_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		if !(x == y) {
			break
		}
		v.reset(OpWasmI64Const)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		if !(x != y) {
			break
		}
		v.reset(OpWasmI64Const)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		y := v.Args[1]
		v.reset(OpWasmI64Eq)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = x
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpWasmI64Eqz)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Eqz_0(v *Value) bool {

	for {
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Eqz {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpWasmI64Eqz {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpWasmI64Eqz)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load_0(v *Value) bool {

	for {
		off := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasmI64Load)
		v.AuxInt = off + off2
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load16S_0(v *Value) bool {

	for {
		off := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasmI64Load16S)
		v.AuxInt = off + off2
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load16U_0(v *Value) bool {

	for {
		off := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasmI64Load16U)
		v.AuxInt = off + off2
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load32S_0(v *Value) bool {

	for {
		off := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasmI64Load32S)
		v.AuxInt = off + off2
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load32U_0(v *Value) bool {

	for {
		off := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasmI64Load32U)
		v.AuxInt = off + off2
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load8S_0(v *Value) bool {

	for {
		off := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasmI64Load8S)
		v.AuxInt = off + off2
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Load8U_0(v *Value) bool {

	for {
		off := v.AuxInt
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v.Args[1]
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasmI64Load8U)
		v.AuxInt = off + off2
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Mul_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = x * y
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		y := v.Args[1]
		v.reset(OpWasmI64Mul)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = x
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Ne_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		if !(x == y) {
			break
		}
		v.reset(OpWasmI64Const)
		v.AuxInt = 0
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		if !(x != y) {
			break
		}
		v.reset(OpWasmI64Const)
		v.AuxInt = 1
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		y := v.Args[1]
		v.reset(OpWasmI64Ne)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = x
		v.AddArg(v0)
		return true
	}

	for {
		_ = v.Args[1]
		x := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		if v_1.AuxInt != 0 {
			break
		}
		v.reset(OpWasmI64Eqz)
		v0 := b.NewValue0(v.Pos, OpWasmI64Eqz, typ.Bool)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Or_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = x | y
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		y := v.Args[1]
		v.reset(OpWasmI64Or)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = x
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Shl_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = x << uint64(y)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64ShrS_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = x >> uint64(y)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64ShrU_0(v *Value) bool {

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = int64(uint64(x) >> uint64(y))
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Store_0(v *Value) bool {

	for {
		off := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasmI64Store)
		v.AuxInt = off + off2
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Store16_0(v *Value) bool {

	for {
		off := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasmI64Store16)
		v.AuxInt = off + off2
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Store32_0(v *Value) bool {

	for {
		off := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasmI64Store32)
		v.AuxInt = off + off2
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Store8_0(v *Value) bool {

	for {
		off := v.AuxInt
		_ = v.Args[2]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64AddConst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v.Args[1]
		mem := v.Args[2]
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasmI64Store8)
		v.AuxInt = off + off2
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpWasmI64Xor_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpWasmI64Const {
			break
		}
		y := v_1.AuxInt
		v.reset(OpWasmI64Const)
		v.AuxInt = x ^ y
		return true
	}

	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpWasmI64Const {
			break
		}
		x := v_0.AuxInt
		y := v.Args[1]
		v.reset(OpWasmI64Xor)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = x
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm_OpXor16_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Xor)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpXor32_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Xor)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpXor64_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Xor)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueWasm_OpXor8_0(v *Value) bool {

	for {
		_ = v.Args[1]
		x := v.Args[0]
		y := v.Args[1]
		v.reset(OpWasmI64Xor)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func (psess *PackageSession) rewriteValueWasm_OpZero_0(v *Value) bool {
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
		v.reset(OpWasmI64Store8)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
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
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmI64Store16)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
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
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmI64Store32)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
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
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmI64Store)
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
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
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmI64Store8)
		v.AuxInt = 2
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store16, psess.types.TypeMem)
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
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
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmI64Store8)
		v.AuxInt = 4
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store32, psess.types.TypeMem)
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
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
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmI64Store16)
		v.AuxInt = 4
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store32, psess.types.TypeMem)
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
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
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmI64Store32)
		v.AuxInt = 3
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store32, psess.types.TypeMem)
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}

	for {
		s := v.AuxInt
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		if !(s%8 != 0 && s > 8) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = s - s%8
		v0 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v0.AuxInt = s % 8
		v0.AddArg(destptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func (psess *PackageSession) rewriteValueWasm_OpZero_10(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		if v.AuxInt != 16 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmI64Store)
		v.AuxInt = 8
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
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
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmI64Store)
		v.AuxInt = 16
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v3.AddArg(destptr)
		v4 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}

	for {
		if v.AuxInt != 32 {
			break
		}
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		v.reset(OpWasmI64Store)
		v.AuxInt = 24
		v.AddArg(destptr)
		v0 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v1.AuxInt = 16
		v1.AddArg(destptr)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v3.AuxInt = 8
		v3.AddArg(destptr)
		v4 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpWasmI64Store, psess.types.TypeMem)
		v5.AddArg(destptr)
		v6 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
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
		_ = v.Args[1]
		destptr := v.Args[0]
		mem := v.Args[1]
		if !(s%8 == 0 && s > 32) {
			break
		}
		v.reset(OpWasmLoweredZero)
		v.AuxInt = s / 8
		v.AddArg(destptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueWasm_OpZeroExt16to32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load16U {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrU)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 48
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 48
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpZeroExt16to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load16U {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrU)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 48
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 48
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpZeroExt32to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load32U {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrU)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 32
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 32
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpZeroExt8to16_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load8U {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrU)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 56
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 56
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpZeroExt8to32_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load8U {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrU)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 56
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 56
		v.AddArg(v2)
		return true
	}
}
func rewriteValueWasm_OpZeroExt8to64_0(v *Value) bool {
	b := v.Block
	_ = b
	typ := &b.Func.Config.Types
	_ = typ

	for {
		x := v.Args[0]
		if x.Op != OpWasmI64Load8U {
			break
		}
		_ = x.Args[1]
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}

	for {
		x := v.Args[0]
		v.reset(OpWasmI64ShrU)
		v0 := b.NewValue0(v.Pos, OpWasmI64Shl, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v1.AuxInt = 56
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpWasmI64Const, typ.Int64)
		v2.AuxInt = 56
		v.AddArg(v2)
		return true
	}
}
func rewriteBlockWasm(b *Block) bool {
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

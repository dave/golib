package ssa

type PackageState struct {
	arm                 *arm.PackageState
	arm64               *arm64.PackageState
	dwarf               *dwarf.PackageState
	mips                *mips.PackageState
	obj                 *obj.PackageState
	objabi              *objabi.PackageState
	ppc64               *ppc64.PackageState
	s390x               *s390x.PackageState
	src                 *src.PackageState
	sys                 *sys.PackageState
	types               *types.PackageState
	wasm                *wasm.PackageState
	x86                 *x86.PackageState
	BlockEnd            *Value
	BlockStart          *Value
	BuildDebug          int
	BuildDump           string
	BuildStats          int
	BuildTest           int
	IntrinsicsDebug     int
	IntrinsicsDisable   bool
	bllikelies          [4]string
	blockString         [96]string
	checkEnabled        bool
	checkpointBound     limitFact
	checkpointFact      fact
	ctzNonZeroOp        map[Op]Op
	domainRelationTable map[Op]struct {
		d domain
		r relation
	}
	domainStrings         [4]string
	dumpFileSeq           int
	fpRegMask386          regMask
	fpRegMaskAMD64        regMask
	fpRegMaskARM          regMask
	fpRegMaskARM64        regMask
	fpRegMaskMIPS         regMask
	fpRegMaskMIPS64       regMask
	fpRegMaskPPC64        regMask
	fpRegMaskS390X        regMask
	fpRegMaskWasm         regMask
	framepointerReg386    int8
	framepointerRegAMD64  int8
	framepointerRegARM    int8
	framepointerRegARM64  int8
	framepointerRegMIPS   int8
	framepointerRegMIPS64 int8
	framepointerRegPPC64  int8
	framepointerRegS390X  int8
	framepointerRegWasm   int8
	gpRegMask386          regMask
	gpRegMaskAMD64        regMask
	gpRegMaskARM          regMask
	gpRegMaskARM64        regMask
	gpRegMaskMIPS         regMask
	gpRegMaskMIPS64       regMask
	gpRegMaskPPC64        regMask
	gpRegMaskS390X        regMask
	gpRegMaskWasm         regMask
	linkReg386            int8
	linkRegAMD64          int8
	linkRegARM            int8
	linkRegARM64          int8
	linkRegMIPS           int8
	linkRegMIPS64         int8
	linkRegPPC64          int8
	linkRegS390X          int8
	linkRegWasm           int8
	noLimit               limit
	opMax                 map[Op]int64
	opMin                 map[Op]int64
	opcodeTable           [2050]opInfo
	passOrder             [26]constraint
	passes                [44]pass
	registers386          [17]Register
	registersAMD64        [33]Register
	registersARM          [33]Register
	registersARM64        [64]Register
	registersMIPS         [48]Register
	registersMIPS64       [63]Register
	registersPPC64        [64]Register
	registersS390X        [33]Register
	registersWasm         [35]Register
	relationStrings       [8]string
	reverseBits           [8]relation
	ruleFile              io.Writer
	specialRegMask386     regMask
	specialRegMaskAMD64   regMask
	specialRegMaskARM     regMask
	specialRegMaskARM64   regMask
	specialRegMaskMIPS    regMask
	specialRegMaskMIPS64  regMask
	specialRegMaskPPC64   regMask
	specialRegMaskS390X   regMask
	specialRegMaskWasm    regMask
}

func NewPackageState(obj_pstate *obj.PackageState, objabi_pstate *objabi.PackageState, types_pstate *types.PackageState, src_pstate *src.PackageState, sys_pstate *sys.PackageState, arm_pstate *arm.PackageState, arm64_pstate *arm64.PackageState, mips_pstate *mips.PackageState, ppc64_pstate *ppc64.PackageState, s390x_pstate *s390x.PackageState, wasm_pstate *wasm.PackageState, x86_pstate *x86.PackageState, dwarf_pstate *dwarf.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.obj = obj_pstate
	pstate.objabi = objabi_pstate
	pstate.types = types_pstate
	pstate.src = src_pstate
	pstate.sys = sys_pstate
	pstate.arm = arm_pstate
	pstate.arm64 = arm64_pstate
	pstate.mips = mips_pstate
	pstate.ppc64 = ppc64_pstate
	pstate.s390x = s390x_pstate
	pstate.wasm = wasm_pstate
	pstate.x86 = x86_pstate
	pstate.dwarf = dwarf_pstate
	pstate.relationStrings = [...]string{
		0: "none", lt: "<", eq: "==", lt | eq: "<=",
		gt: ">", gt | lt: "!=", gt | eq: ">=", gt | eq | lt: "any",
	}
	pstate.domainStrings = [...]string{
		"signed", "unsigned", "pointer", "boolean",
	}
	pstate.noLimit = limit{math.MinInt64, math.MaxInt64, 0, math.MaxUint64}
	pstate.checkpointFact = fact{}
	pstate.checkpointBound = limitFact{}
	pstate.reverseBits = [...]relation{0, 4, 2, 6, 1, 5, 3, 7}
	pstate.blockString = [...]string{
		BlockInvalid: "BlockInvalid",

		Block386EQ:  "EQ",
		Block386NE:  "NE",
		Block386LT:  "LT",
		Block386LE:  "LE",
		Block386GT:  "GT",
		Block386GE:  "GE",
		Block386ULT: "ULT",
		Block386ULE: "ULE",
		Block386UGT: "UGT",
		Block386UGE: "UGE",
		Block386EQF: "EQF",
		Block386NEF: "NEF",
		Block386ORD: "ORD",
		Block386NAN: "NAN",

		BlockAMD64EQ:  "EQ",
		BlockAMD64NE:  "NE",
		BlockAMD64LT:  "LT",
		BlockAMD64LE:  "LE",
		BlockAMD64GT:  "GT",
		BlockAMD64GE:  "GE",
		BlockAMD64ULT: "ULT",
		BlockAMD64ULE: "ULE",
		BlockAMD64UGT: "UGT",
		BlockAMD64UGE: "UGE",
		BlockAMD64EQF: "EQF",
		BlockAMD64NEF: "NEF",
		BlockAMD64ORD: "ORD",
		BlockAMD64NAN: "NAN",

		BlockARMEQ:  "EQ",
		BlockARMNE:  "NE",
		BlockARMLT:  "LT",
		BlockARMLE:  "LE",
		BlockARMGT:  "GT",
		BlockARMGE:  "GE",
		BlockARMULT: "ULT",
		BlockARMULE: "ULE",
		BlockARMUGT: "UGT",
		BlockARMUGE: "UGE",

		BlockARM64EQ:   "EQ",
		BlockARM64NE:   "NE",
		BlockARM64LT:   "LT",
		BlockARM64LE:   "LE",
		BlockARM64GT:   "GT",
		BlockARM64GE:   "GE",
		BlockARM64ULT:  "ULT",
		BlockARM64ULE:  "ULE",
		BlockARM64UGT:  "UGT",
		BlockARM64UGE:  "UGE",
		BlockARM64Z:    "Z",
		BlockARM64NZ:   "NZ",
		BlockARM64ZW:   "ZW",
		BlockARM64NZW:  "NZW",
		BlockARM64TBZ:  "TBZ",
		BlockARM64TBNZ: "TBNZ",

		BlockMIPSEQ:  "EQ",
		BlockMIPSNE:  "NE",
		BlockMIPSLTZ: "LTZ",
		BlockMIPSLEZ: "LEZ",
		BlockMIPSGTZ: "GTZ",
		BlockMIPSGEZ: "GEZ",
		BlockMIPSFPT: "FPT",
		BlockMIPSFPF: "FPF",

		BlockMIPS64EQ:  "EQ",
		BlockMIPS64NE:  "NE",
		BlockMIPS64LTZ: "LTZ",
		BlockMIPS64LEZ: "LEZ",
		BlockMIPS64GTZ: "GTZ",
		BlockMIPS64GEZ: "GEZ",
		BlockMIPS64FPT: "FPT",
		BlockMIPS64FPF: "FPF",

		BlockPPC64EQ:  "EQ",
		BlockPPC64NE:  "NE",
		BlockPPC64LT:  "LT",
		BlockPPC64LE:  "LE",
		BlockPPC64GT:  "GT",
		BlockPPC64GE:  "GE",
		BlockPPC64FLT: "FLT",
		BlockPPC64FLE: "FLE",
		BlockPPC64FGT: "FGT",
		BlockPPC64FGE: "FGE",

		BlockS390XEQ:  "EQ",
		BlockS390XNE:  "NE",
		BlockS390XLT:  "LT",
		BlockS390XLE:  "LE",
		BlockS390XGT:  "GT",
		BlockS390XGE:  "GE",
		BlockS390XGTF: "GTF",
		BlockS390XGEF: "GEF",

		BlockPlain:  "Plain",
		BlockIf:     "If",
		BlockDefer:  "Defer",
		BlockRet:    "Ret",
		BlockRetJmp: "RetJmp",
		BlockExit:   "Exit",
		BlockFirst:  "First",
	}
	pstate.opMin = map[Op]int64{
		OpAdd64: math.MinInt64, OpSub64: math.MinInt64,
		OpAdd32: math.MinInt32, OpSub32: math.MinInt32,
	}
	pstate.opMax = map[Op]int64{
		OpAdd64: math.MaxInt64, OpSub64: math.MaxInt64,
		OpAdd32: math.MaxInt32, OpSub32: math.MaxInt32,
	}
	pstate.ctzNonZeroOp = map[Op]Op{OpCtz8: OpCtz8NonZero, OpCtz16: OpCtz16NonZero, OpCtz32: OpCtz32NonZero, OpCtz64: OpCtz64NonZero}
	pstate.domainRelationTable = map[Op]struct {
		d domain
		r relation
	}{
		OpEq8:   {signed | unsigned, eq},
		OpEq16:  {signed | unsigned, eq},
		OpEq32:  {signed | unsigned, eq},
		OpEq64:  {signed | unsigned, eq},
		OpEqPtr: {pointer, eq},

		OpNeq8:   {signed | unsigned, lt | gt},
		OpNeq16:  {signed | unsigned, lt | gt},
		OpNeq32:  {signed | unsigned, lt | gt},
		OpNeq64:  {signed | unsigned, lt | gt},
		OpNeqPtr: {pointer, lt | gt},

		OpLess8:   {signed, lt},
		OpLess8U:  {unsigned, lt},
		OpLess16:  {signed, lt},
		OpLess16U: {unsigned, lt},
		OpLess32:  {signed, lt},
		OpLess32U: {unsigned, lt},
		OpLess64:  {signed, lt},
		OpLess64U: {unsigned, lt},

		OpLeq8:   {signed, lt | eq},
		OpLeq8U:  {unsigned, lt | eq},
		OpLeq16:  {signed, lt | eq},
		OpLeq16U: {unsigned, lt | eq},
		OpLeq32:  {signed, lt | eq},
		OpLeq32U: {unsigned, lt | eq},
		OpLeq64:  {signed, lt | eq},
		OpLeq64U: {unsigned, lt | eq},

		OpGeq8:   {signed, eq | gt},
		OpGeq8U:  {unsigned, eq | gt},
		OpGeq16:  {signed, eq | gt},
		OpGeq16U: {unsigned, eq | gt},
		OpGeq32:  {signed, eq | gt},
		OpGeq32U: {unsigned, eq | gt},
		OpGeq64:  {signed, eq | gt},
		OpGeq64U: {unsigned, eq | gt},

		OpGreater8:   {signed, gt},
		OpGreater8U:  {unsigned, gt},
		OpGreater16:  {signed, gt},
		OpGreater16U: {unsigned, gt},
		OpGreater32:  {signed, gt},
		OpGreater32U: {unsigned, gt},
		OpGreater64:  {signed, gt},
		OpGreater64U: {unsigned, gt},

		// For these ops, the negative branch is different: we can only
		// prove signed/GE (signed/GT) if we can prove that arg0 is non-negative.
		// See the special case in addBranchRestrictions.
		OpIsInBounds:      {signed | unsigned, lt},      // 0 <= arg0 < arg1
		OpIsSliceInBounds: {signed | unsigned, lt | eq}, // 0 <= arg0 <= arg1
	}
	pstate.registers386 = [...]Register{
		{0, x86.REG_AX, 0, "AX"},
		{1, x86.REG_CX, 1, "CX"},
		{2, x86.REG_DX, 2, "DX"},
		{3, x86.REG_BX, 3, "BX"},
		{4, x86.REGSP, -1, "SP"},
		{5, x86.REG_BP, 4, "BP"},
		{6, x86.REG_SI, 5, "SI"},
		{7, x86.REG_DI, 6, "DI"},
		{8, x86.REG_X0, -1, "X0"},
		{9, x86.REG_X1, -1, "X1"},
		{10, x86.REG_X2, -1, "X2"},
		{11, x86.REG_X3, -1, "X3"},
		{12, x86.REG_X4, -1, "X4"},
		{13, x86.REG_X5, -1, "X5"},
		{14, x86.REG_X6, -1, "X6"},
		{15, x86.REG_X7, -1, "X7"},
		{16, 0, -1, "SB"},
	}
	pstate.gpRegMask386 = regMask(239)
	pstate.fpRegMask386 = regMask(65280)
	pstate.specialRegMask386 = regMask(0)
	pstate.framepointerReg386 = int8(5)
	pstate.linkReg386 = int8(-1)
	pstate.registersAMD64 = [...]Register{
		{0, x86.REG_AX, 0, "AX"},
		{1, x86.REG_CX, 1, "CX"},
		{2, x86.REG_DX, 2, "DX"},
		{3, x86.REG_BX, 3, "BX"},
		{4, x86.REGSP, -1, "SP"},
		{5, x86.REG_BP, 4, "BP"},
		{6, x86.REG_SI, 5, "SI"},
		{7, x86.REG_DI, 6, "DI"},
		{8, x86.REG_R8, 7, "R8"},
		{9, x86.REG_R9, 8, "R9"},
		{10, x86.REG_R10, 9, "R10"},
		{11, x86.REG_R11, 10, "R11"},
		{12, x86.REG_R12, 11, "R12"},
		{13, x86.REG_R13, 12, "R13"},
		{14, x86.REG_R14, 13, "R14"},
		{15, x86.REG_R15, 14, "R15"},
		{16, x86.REG_X0, -1, "X0"},
		{17, x86.REG_X1, -1, "X1"},
		{18, x86.REG_X2, -1, "X2"},
		{19, x86.REG_X3, -1, "X3"},
		{20, x86.REG_X4, -1, "X4"},
		{21, x86.REG_X5, -1, "X5"},
		{22, x86.REG_X6, -1, "X6"},
		{23, x86.REG_X7, -1, "X7"},
		{24, x86.REG_X8, -1, "X8"},
		{25, x86.REG_X9, -1, "X9"},
		{26, x86.REG_X10, -1, "X10"},
		{27, x86.REG_X11, -1, "X11"},
		{28, x86.REG_X12, -1, "X12"},
		{29, x86.REG_X13, -1, "X13"},
		{30, x86.REG_X14, -1, "X14"},
		{31, x86.REG_X15, -1, "X15"},
		{32, 0, -1, "SB"},
	}
	pstate.gpRegMaskAMD64 = regMask(65519)
	pstate.fpRegMaskAMD64 = regMask(4294901760)
	pstate.specialRegMaskAMD64 = regMask(0)
	pstate.framepointerRegAMD64 = int8(5)
	pstate.linkRegAMD64 = int8(-1)
	pstate.registersARM = [...]Register{
		{0, arm.REG_R0, 0, "R0"},
		{1, arm.REG_R1, 1, "R1"},
		{2, arm.REG_R2, 2, "R2"},
		{3, arm.REG_R3, 3, "R3"},
		{4, arm.REG_R4, 4, "R4"},
		{5, arm.REG_R5, 5, "R5"},
		{6, arm.REG_R6, 6, "R6"},
		{7, arm.REG_R7, 7, "R7"},
		{8, arm.REG_R8, 8, "R8"},
		{9, arm.REG_R9, 9, "R9"},
		{10, arm.REGG, -1, "g"},
		{11, arm.REG_R11, -1, "R11"},
		{12, arm.REG_R12, 10, "R12"},
		{13, arm.REGSP, -1, "SP"},
		{14, arm.REG_R14, 11, "R14"},
		{15, arm.REG_R15, -1, "R15"},
		{16, arm.REG_F0, -1, "F0"},
		{17, arm.REG_F1, -1, "F1"},
		{18, arm.REG_F2, -1, "F2"},
		{19, arm.REG_F3, -1, "F3"},
		{20, arm.REG_F4, -1, "F4"},
		{21, arm.REG_F5, -1, "F5"},
		{22, arm.REG_F6, -1, "F6"},
		{23, arm.REG_F7, -1, "F7"},
		{24, arm.REG_F8, -1, "F8"},
		{25, arm.REG_F9, -1, "F9"},
		{26, arm.REG_F10, -1, "F10"},
		{27, arm.REG_F11, -1, "F11"},
		{28, arm.REG_F12, -1, "F12"},
		{29, arm.REG_F13, -1, "F13"},
		{30, arm.REG_F14, -1, "F14"},
		{31, arm.REG_F15, -1, "F15"},
		{32, 0, -1, "SB"},
	}
	pstate.gpRegMaskARM = regMask(21503)
	pstate.fpRegMaskARM = regMask(4294901760)
	pstate.specialRegMaskARM = regMask(0)
	pstate.framepointerRegARM = int8(-1)
	pstate.linkRegARM = int8(14)
	pstate.registersARM64 = [...]Register{
		{0, arm64.REG_R0, 0, "R0"},
		{1, arm64.REG_R1, 1, "R1"},
		{2, arm64.REG_R2, 2, "R2"},
		{3, arm64.REG_R3, 3, "R3"},
		{4, arm64.REG_R4, 4, "R4"},
		{5, arm64.REG_R5, 5, "R5"},
		{6, arm64.REG_R6, 6, "R6"},
		{7, arm64.REG_R7, 7, "R7"},
		{8, arm64.REG_R8, 8, "R8"},
		{9, arm64.REG_R9, 9, "R9"},
		{10, arm64.REG_R10, 10, "R10"},
		{11, arm64.REG_R11, 11, "R11"},
		{12, arm64.REG_R12, 12, "R12"},
		{13, arm64.REG_R13, 13, "R13"},
		{14, arm64.REG_R14, 14, "R14"},
		{15, arm64.REG_R15, 15, "R15"},
		{16, arm64.REG_R16, 16, "R16"},
		{17, arm64.REG_R17, 17, "R17"},
		{18, arm64.REG_R18, -1, "R18"},
		{19, arm64.REG_R19, 18, "R19"},
		{20, arm64.REG_R20, 19, "R20"},
		{21, arm64.REG_R21, 20, "R21"},
		{22, arm64.REG_R22, 21, "R22"},
		{23, arm64.REG_R23, 22, "R23"},
		{24, arm64.REG_R24, 23, "R24"},
		{25, arm64.REG_R25, 24, "R25"},
		{26, arm64.REG_R26, 25, "R26"},
		{27, arm64.REGG, -1, "g"},
		{28, arm64.REG_R29, -1, "R29"},
		{29, arm64.REG_R30, 26, "R30"},
		{30, arm64.REGSP, -1, "SP"},
		{31, arm64.REG_F0, -1, "F0"},
		{32, arm64.REG_F1, -1, "F1"},
		{33, arm64.REG_F2, -1, "F2"},
		{34, arm64.REG_F3, -1, "F3"},
		{35, arm64.REG_F4, -1, "F4"},
		{36, arm64.REG_F5, -1, "F5"},
		{37, arm64.REG_F6, -1, "F6"},
		{38, arm64.REG_F7, -1, "F7"},
		{39, arm64.REG_F8, -1, "F8"},
		{40, arm64.REG_F9, -1, "F9"},
		{41, arm64.REG_F10, -1, "F10"},
		{42, arm64.REG_F11, -1, "F11"},
		{43, arm64.REG_F12, -1, "F12"},
		{44, arm64.REG_F13, -1, "F13"},
		{45, arm64.REG_F14, -1, "F14"},
		{46, arm64.REG_F15, -1, "F15"},
		{47, arm64.REG_F16, -1, "F16"},
		{48, arm64.REG_F17, -1, "F17"},
		{49, arm64.REG_F18, -1, "F18"},
		{50, arm64.REG_F19, -1, "F19"},
		{51, arm64.REG_F20, -1, "F20"},
		{52, arm64.REG_F21, -1, "F21"},
		{53, arm64.REG_F22, -1, "F22"},
		{54, arm64.REG_F23, -1, "F23"},
		{55, arm64.REG_F24, -1, "F24"},
		{56, arm64.REG_F25, -1, "F25"},
		{57, arm64.REG_F26, -1, "F26"},
		{58, arm64.REG_F27, -1, "F27"},
		{59, arm64.REG_F28, -1, "F28"},
		{60, arm64.REG_F29, -1, "F29"},
		{61, arm64.REG_F30, -1, "F30"},
		{62, arm64.REG_F31, -1, "F31"},
		{63, 0, -1, "SB"},
	}
	pstate.gpRegMaskARM64 = regMask(670826495)
	pstate.fpRegMaskARM64 = regMask(9223372034707292160)
	pstate.specialRegMaskARM64 = regMask(0)
	pstate.framepointerRegARM64 = int8(-1)
	pstate.linkRegARM64 = int8(29)
	pstate.registersMIPS = [...]Register{
		{0, mips.REG_R0, -1, "R0"},
		{1, mips.REG_R1, 0, "R1"},
		{2, mips.REG_R2, 1, "R2"},
		{3, mips.REG_R3, 2, "R3"},
		{4, mips.REG_R4, 3, "R4"},
		{5, mips.REG_R5, 4, "R5"},
		{6, mips.REG_R6, 5, "R6"},
		{7, mips.REG_R7, 6, "R7"},
		{8, mips.REG_R8, 7, "R8"},
		{9, mips.REG_R9, 8, "R9"},
		{10, mips.REG_R10, 9, "R10"},
		{11, mips.REG_R11, 10, "R11"},
		{12, mips.REG_R12, 11, "R12"},
		{13, mips.REG_R13, 12, "R13"},
		{14, mips.REG_R14, 13, "R14"},
		{15, mips.REG_R15, 14, "R15"},
		{16, mips.REG_R16, 15, "R16"},
		{17, mips.REG_R17, 16, "R17"},
		{18, mips.REG_R18, 17, "R18"},
		{19, mips.REG_R19, 18, "R19"},
		{20, mips.REG_R20, 19, "R20"},
		{21, mips.REG_R21, 20, "R21"},
		{22, mips.REG_R22, 21, "R22"},
		{23, mips.REG_R24, 22, "R24"},
		{24, mips.REG_R25, 23, "R25"},
		{25, mips.REG_R28, 24, "R28"},
		{26, mips.REGSP, -1, "SP"},
		{27, mips.REGG, -1, "g"},
		{28, mips.REG_R31, 25, "R31"},
		{29, mips.REG_F0, -1, "F0"},
		{30, mips.REG_F2, -1, "F2"},
		{31, mips.REG_F4, -1, "F4"},
		{32, mips.REG_F6, -1, "F6"},
		{33, mips.REG_F8, -1, "F8"},
		{34, mips.REG_F10, -1, "F10"},
		{35, mips.REG_F12, -1, "F12"},
		{36, mips.REG_F14, -1, "F14"},
		{37, mips.REG_F16, -1, "F16"},
		{38, mips.REG_F18, -1, "F18"},
		{39, mips.REG_F20, -1, "F20"},
		{40, mips.REG_F22, -1, "F22"},
		{41, mips.REG_F24, -1, "F24"},
		{42, mips.REG_F26, -1, "F26"},
		{43, mips.REG_F28, -1, "F28"},
		{44, mips.REG_F30, -1, "F30"},
		{45, mips.REG_HI, -1, "HI"},
		{46, mips.REG_LO, -1, "LO"},
		{47, 0, -1, "SB"},
	}
	pstate.gpRegMaskMIPS = regMask(335544318)
	pstate.fpRegMaskMIPS = regMask(35183835217920)
	pstate.specialRegMaskMIPS = regMask(105553116266496)
	pstate.framepointerRegMIPS = int8(-1)
	pstate.linkRegMIPS = int8(28)
	pstate.registersMIPS64 = [...]Register{
		{0, mips.REG_R0, -1, "R0"},
		{1, mips.REG_R1, 0, "R1"},
		{2, mips.REG_R2, 1, "R2"},
		{3, mips.REG_R3, 2, "R3"},
		{4, mips.REG_R4, 3, "R4"},
		{5, mips.REG_R5, 4, "R5"},
		{6, mips.REG_R6, 5, "R6"},
		{7, mips.REG_R7, 6, "R7"},
		{8, mips.REG_R8, 7, "R8"},
		{9, mips.REG_R9, 8, "R9"},
		{10, mips.REG_R10, 9, "R10"},
		{11, mips.REG_R11, 10, "R11"},
		{12, mips.REG_R12, 11, "R12"},
		{13, mips.REG_R13, 12, "R13"},
		{14, mips.REG_R14, 13, "R14"},
		{15, mips.REG_R15, 14, "R15"},
		{16, mips.REG_R16, 15, "R16"},
		{17, mips.REG_R17, 16, "R17"},
		{18, mips.REG_R18, 17, "R18"},
		{19, mips.REG_R19, 18, "R19"},
		{20, mips.REG_R20, 19, "R20"},
		{21, mips.REG_R21, 20, "R21"},
		{22, mips.REG_R22, 21, "R22"},
		{23, mips.REG_R24, 22, "R24"},
		{24, mips.REG_R25, 23, "R25"},
		{25, mips.REGSP, -1, "SP"},
		{26, mips.REGG, -1, "g"},
		{27, mips.REG_R31, 24, "R31"},
		{28, mips.REG_F0, -1, "F0"},
		{29, mips.REG_F1, -1, "F1"},
		{30, mips.REG_F2, -1, "F2"},
		{31, mips.REG_F3, -1, "F3"},
		{32, mips.REG_F4, -1, "F4"},
		{33, mips.REG_F5, -1, "F5"},
		{34, mips.REG_F6, -1, "F6"},
		{35, mips.REG_F7, -1, "F7"},
		{36, mips.REG_F8, -1, "F8"},
		{37, mips.REG_F9, -1, "F9"},
		{38, mips.REG_F10, -1, "F10"},
		{39, mips.REG_F11, -1, "F11"},
		{40, mips.REG_F12, -1, "F12"},
		{41, mips.REG_F13, -1, "F13"},
		{42, mips.REG_F14, -1, "F14"},
		{43, mips.REG_F15, -1, "F15"},
		{44, mips.REG_F16, -1, "F16"},
		{45, mips.REG_F17, -1, "F17"},
		{46, mips.REG_F18, -1, "F18"},
		{47, mips.REG_F19, -1, "F19"},
		{48, mips.REG_F20, -1, "F20"},
		{49, mips.REG_F21, -1, "F21"},
		{50, mips.REG_F22, -1, "F22"},
		{51, mips.REG_F23, -1, "F23"},
		{52, mips.REG_F24, -1, "F24"},
		{53, mips.REG_F25, -1, "F25"},
		{54, mips.REG_F26, -1, "F26"},
		{55, mips.REG_F27, -1, "F27"},
		{56, mips.REG_F28, -1, "F28"},
		{57, mips.REG_F29, -1, "F29"},
		{58, mips.REG_F30, -1, "F30"},
		{59, mips.REG_F31, -1, "F31"},
		{60, mips.REG_HI, -1, "HI"},
		{61, mips.REG_LO, -1, "LO"},
		{62, 0, -1, "SB"},
	}
	pstate.gpRegMaskMIPS64 = regMask(167772158)
	pstate.fpRegMaskMIPS64 = regMask(1152921504338411520)
	pstate.specialRegMaskMIPS64 = regMask(3458764513820540928)
	pstate.framepointerRegMIPS64 = int8(-1)
	pstate.linkRegMIPS64 = int8(27)
	pstate.registersPPC64 = [...]Register{
		{0, ppc64.REG_R0, -1, "R0"},
		{1, ppc64.REGSP, -1, "SP"},
		{2, 0, -1, "SB"},
		{3, ppc64.REG_R3, 0, "R3"},
		{4, ppc64.REG_R4, 1, "R4"},
		{5, ppc64.REG_R5, 2, "R5"},
		{6, ppc64.REG_R6, 3, "R6"},
		{7, ppc64.REG_R7, 4, "R7"},
		{8, ppc64.REG_R8, 5, "R8"},
		{9, ppc64.REG_R9, 6, "R9"},
		{10, ppc64.REG_R10, 7, "R10"},
		{11, ppc64.REG_R11, 8, "R11"},
		{12, ppc64.REG_R12, 9, "R12"},
		{13, ppc64.REG_R13, -1, "R13"},
		{14, ppc64.REG_R14, 10, "R14"},
		{15, ppc64.REG_R15, 11, "R15"},
		{16, ppc64.REG_R16, 12, "R16"},
		{17, ppc64.REG_R17, 13, "R17"},
		{18, ppc64.REG_R18, 14, "R18"},
		{19, ppc64.REG_R19, 15, "R19"},
		{20, ppc64.REG_R20, 16, "R20"},
		{21, ppc64.REG_R21, 17, "R21"},
		{22, ppc64.REG_R22, 18, "R22"},
		{23, ppc64.REG_R23, 19, "R23"},
		{24, ppc64.REG_R24, 20, "R24"},
		{25, ppc64.REG_R25, 21, "R25"},
		{26, ppc64.REG_R26, 22, "R26"},
		{27, ppc64.REG_R27, 23, "R27"},
		{28, ppc64.REG_R28, 24, "R28"},
		{29, ppc64.REG_R29, 25, "R29"},
		{30, ppc64.REGG, -1, "g"},
		{31, ppc64.REG_R31, -1, "R31"},
		{32, ppc64.REG_F0, -1, "F0"},
		{33, ppc64.REG_F1, -1, "F1"},
		{34, ppc64.REG_F2, -1, "F2"},
		{35, ppc64.REG_F3, -1, "F3"},
		{36, ppc64.REG_F4, -1, "F4"},
		{37, ppc64.REG_F5, -1, "F5"},
		{38, ppc64.REG_F6, -1, "F6"},
		{39, ppc64.REG_F7, -1, "F7"},
		{40, ppc64.REG_F8, -1, "F8"},
		{41, ppc64.REG_F9, -1, "F9"},
		{42, ppc64.REG_F10, -1, "F10"},
		{43, ppc64.REG_F11, -1, "F11"},
		{44, ppc64.REG_F12, -1, "F12"},
		{45, ppc64.REG_F13, -1, "F13"},
		{46, ppc64.REG_F14, -1, "F14"},
		{47, ppc64.REG_F15, -1, "F15"},
		{48, ppc64.REG_F16, -1, "F16"},
		{49, ppc64.REG_F17, -1, "F17"},
		{50, ppc64.REG_F18, -1, "F18"},
		{51, ppc64.REG_F19, -1, "F19"},
		{52, ppc64.REG_F20, -1, "F20"},
		{53, ppc64.REG_F21, -1, "F21"},
		{54, ppc64.REG_F22, -1, "F22"},
		{55, ppc64.REG_F23, -1, "F23"},
		{56, ppc64.REG_F24, -1, "F24"},
		{57, ppc64.REG_F25, -1, "F25"},
		{58, ppc64.REG_F26, -1, "F26"},
		{59, ppc64.REG_F27, -1, "F27"},
		{60, ppc64.REG_F28, -1, "F28"},
		{61, ppc64.REG_F29, -1, "F29"},
		{62, ppc64.REG_F30, -1, "F30"},
		{63, ppc64.REG_F31, -1, "F31"},
	}
	pstate.gpRegMaskPPC64 = regMask(1073733624)
	pstate.fpRegMaskPPC64 = regMask(576460743713488896)
	pstate.specialRegMaskPPC64 = regMask(0)
	pstate.framepointerRegPPC64 = int8(1)
	pstate.linkRegPPC64 = int8(-1)
	pstate.registersS390X = [...]Register{
		{0, s390x.REG_R0, 0, "R0"},
		{1, s390x.REG_R1, 1, "R1"},
		{2, s390x.REG_R2, 2, "R2"},
		{3, s390x.REG_R3, 3, "R3"},
		{4, s390x.REG_R4, 4, "R4"},
		{5, s390x.REG_R5, 5, "R5"},
		{6, s390x.REG_R6, 6, "R6"},
		{7, s390x.REG_R7, 7, "R7"},
		{8, s390x.REG_R8, 8, "R8"},
		{9, s390x.REG_R9, 9, "R9"},
		{10, s390x.REG_R10, -1, "R10"},
		{11, s390x.REG_R11, 10, "R11"},
		{12, s390x.REG_R12, 11, "R12"},
		{13, s390x.REGG, -1, "g"},
		{14, s390x.REG_R14, 12, "R14"},
		{15, s390x.REGSP, -1, "SP"},
		{16, s390x.REG_F0, -1, "F0"},
		{17, s390x.REG_F1, -1, "F1"},
		{18, s390x.REG_F2, -1, "F2"},
		{19, s390x.REG_F3, -1, "F3"},
		{20, s390x.REG_F4, -1, "F4"},
		{21, s390x.REG_F5, -1, "F5"},
		{22, s390x.REG_F6, -1, "F6"},
		{23, s390x.REG_F7, -1, "F7"},
		{24, s390x.REG_F8, -1, "F8"},
		{25, s390x.REG_F9, -1, "F9"},
		{26, s390x.REG_F10, -1, "F10"},
		{27, s390x.REG_F11, -1, "F11"},
		{28, s390x.REG_F12, -1, "F12"},
		{29, s390x.REG_F13, -1, "F13"},
		{30, s390x.REG_F14, -1, "F14"},
		{31, s390x.REG_F15, -1, "F15"},
		{32, 0, -1, "SB"},
	}
	pstate.gpRegMaskS390X = regMask(23551)
	pstate.fpRegMaskS390X = regMask(4294901760)
	pstate.specialRegMaskS390X = regMask(0)
	pstate.framepointerRegS390X = int8(-1)
	pstate.linkRegS390X = int8(14)
	pstate.registersWasm = [...]Register{
		{0, wasm.REG_R0, 0, "R0"},
		{1, wasm.REG_R1, 1, "R1"},
		{2, wasm.REG_R2, 2, "R2"},
		{3, wasm.REG_R3, 3, "R3"},
		{4, wasm.REG_R4, 4, "R4"},
		{5, wasm.REG_R5, 5, "R5"},
		{6, wasm.REG_R6, 6, "R6"},
		{7, wasm.REG_R7, 7, "R7"},
		{8, wasm.REG_R8, 8, "R8"},
		{9, wasm.REG_R9, 9, "R9"},
		{10, wasm.REG_R10, 10, "R10"},
		{11, wasm.REG_R11, 11, "R11"},
		{12, wasm.REG_R12, 12, "R12"},
		{13, wasm.REG_R13, 13, "R13"},
		{14, wasm.REG_R14, 14, "R14"},
		{15, wasm.REG_R15, 15, "R15"},
		{16, wasm.REG_F0, -1, "F0"},
		{17, wasm.REG_F1, -1, "F1"},
		{18, wasm.REG_F2, -1, "F2"},
		{19, wasm.REG_F3, -1, "F3"},
		{20, wasm.REG_F4, -1, "F4"},
		{21, wasm.REG_F5, -1, "F5"},
		{22, wasm.REG_F6, -1, "F6"},
		{23, wasm.REG_F7, -1, "F7"},
		{24, wasm.REG_F8, -1, "F8"},
		{25, wasm.REG_F9, -1, "F9"},
		{26, wasm.REG_F10, -1, "F10"},
		{27, wasm.REG_F11, -1, "F11"},
		{28, wasm.REG_F12, -1, "F12"},
		{29, wasm.REG_F13, -1, "F13"},
		{30, wasm.REG_F14, -1, "F14"},
		{31, wasm.REG_F15, -1, "F15"},
		{32, wasm.REGSP, -1, "SP"},
		{33, wasm.REGG, -1, "g"},
		{34, 0, -1, "SB"},
	}
	pstate.gpRegMaskWasm = regMask(65535)
	pstate.fpRegMaskWasm = regMask(4294901760)
	pstate.specialRegMaskWasm = regMask(0)
	pstate.framepointerRegWasm = int8(-1)
	pstate.linkRegWasm = int8(-1)
	pstate.opcodeTable = [...]opInfo{
		{name: "OpInvalid"},

		{
			name:         "ADDSS",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			usesScratch:  true,
			asm:          x86.AADDSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:         "ADDSD",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			asm:          x86.AADDSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:         "SUBSS",
			argLen:       2,
			resultInArg0: true,
			usesScratch:  true,
			asm:          x86.ASUBSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:         "SUBSD",
			argLen:       2,
			resultInArg0: true,
			asm:          x86.ASUBSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:         "MULSS",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			usesScratch:  true,
			asm:          x86.AMULSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:         "MULSD",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			asm:          x86.AMULSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:         "DIVSS",
			argLen:       2,
			resultInArg0: true,
			usesScratch:  true,
			asm:          x86.ADIVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:         "DIVSD",
			argLen:       2,
			resultInArg0: true,
			asm:          x86.ADIVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:           "MOVSSload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:           "MOVSDload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:              "MOVSSconst",
			auxType:           auxFloat32,
			argLen:            0,
			rematerializeable: true,
			asm:               x86.AMOVSS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:              "MOVSDconst",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			asm:               x86.AMOVSD,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:      "MOVSSloadidx1",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:      "MOVSSloadidx4",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:      "MOVSDloadidx1",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:      "MOVSDloadidx8",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:           "MOVSSstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "MOVSDstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:      "MOVSSstoreidx1",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{2, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:      "MOVSSstoreidx4",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{2, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:      "MOVSDstoreidx1",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{2, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:      "MOVSDstoreidx8",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{2, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "ADDSSload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AADDSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:           "ADDSDload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AADDSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:           "SUBSSload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.ASUBSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:           "SUBSDload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.ASUBSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:           "MULSSload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AMULSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:           "MULSDload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AMULSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:         "ADDL",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          x86.AADDL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 239}, // AX CX DX BX BP SI DI
					{0, 255}, // AX CX DX BX SP BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "ADDLconst",
			auxType:      auxInt32,
			argLen:       1,
			clobberFlags: true,
			asm:          x86.AADDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "ADDLcarry",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			asm:          x86.AADDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
					{1, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "ADDLconstcarry",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			asm:          x86.AADDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "ADCL",
			argLen:       3,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AADCL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
					{1, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "ADCLconst",
			auxType:      auxInt32,
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AADCL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SUBL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASUBL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
					{1, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SUBLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASUBL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SUBLcarry",
			argLen:       2,
			resultInArg0: true,
			asm:          x86.ASUBL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
					{1, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SUBLconstcarry",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			asm:          x86.ASUBL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SBBL",
			argLen:       3,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASBBL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
					{1, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SBBLconst",
			auxType:      auxInt32,
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASBBL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "MULL",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AIMULL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
					{1, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "MULLconst",
			auxType:      auxInt32,
			argLen:       1,
			clobberFlags: true,
			asm:          x86.AIMUL3L,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "HMULL",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          x86.AIMULL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},   // AX
					{1, 255}, // AX CX DX BX SP BP SI DI
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:         "HMULLU",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          x86.AMULL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},   // AX
					{1, 255}, // AX CX DX BX SP BP SI DI
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:         "MULLQU",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          x86.AMULL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},   // AX
					{1, 255}, // AX CX DX BX SP BP SI DI
				},
				outputs: []outputInfo{
					{0, 4}, // DX
					{1, 1}, // AX
				},
			},
		},
		{
			name:         "AVGLU",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
					{1, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "DIVL",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.AIDIVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},   // AX
					{1, 251}, // AX CX BX SP BP SI DI
				},
				clobbers: 4, // DX
				outputs: []outputInfo{
					{0, 1}, // AX
				},
			},
		},
		{
			name:         "DIVW",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.AIDIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},   // AX
					{1, 251}, // AX CX BX SP BP SI DI
				},
				clobbers: 4, // DX
				outputs: []outputInfo{
					{0, 1}, // AX
				},
			},
		},
		{
			name:         "DIVLU",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.ADIVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},   // AX
					{1, 251}, // AX CX BX SP BP SI DI
				},
				clobbers: 4, // DX
				outputs: []outputInfo{
					{0, 1}, // AX
				},
			},
		},
		{
			name:         "DIVWU",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.ADIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},   // AX
					{1, 251}, // AX CX BX SP BP SI DI
				},
				clobbers: 4, // DX
				outputs: []outputInfo{
					{0, 1}, // AX
				},
			},
		},
		{
			name:         "MODL",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.AIDIVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},   // AX
					{1, 251}, // AX CX BX SP BP SI DI
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:         "MODW",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.AIDIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},   // AX
					{1, 251}, // AX CX BX SP BP SI DI
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:         "MODLU",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.ADIVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},   // AX
					{1, 251}, // AX CX BX SP BP SI DI
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:         "MODWU",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.ADIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},   // AX
					{1, 251}, // AX CX BX SP BP SI DI
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:         "ANDL",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AANDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
					{1, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "ANDLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AANDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "ORL",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
					{1, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "ORLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "XORL",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AXORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
					{1, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "XORLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AXORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "CMPL",
			argLen: 2,
			asm:    x86.ACMPL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
					{1, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:   "CMPW",
			argLen: 2,
			asm:    x86.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
					{1, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:   "CMPB",
			argLen: 2,
			asm:    x86.ACMPB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
					{1, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:    "CMPLconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     x86.ACMPL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:    "CMPWconst",
			auxType: auxInt16,
			argLen:  1,
			asm:     x86.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:    "CMPBconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     x86.ACMPB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:        "UCOMISS",
			argLen:      2,
			usesScratch: true,
			asm:         x86.AUCOMISS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:        "UCOMISD",
			argLen:      2,
			usesScratch: true,
			asm:         x86.AUCOMISD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:        "TESTL",
			argLen:      2,
			commutative: true,
			asm:         x86.ATESTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
					{1, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:        "TESTW",
			argLen:      2,
			commutative: true,
			asm:         x86.ATESTW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
					{1, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:        "TESTB",
			argLen:      2,
			commutative: true,
			asm:         x86.ATESTB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
					{1, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:    "TESTLconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     x86.ATESTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:    "TESTWconst",
			auxType: auxInt16,
			argLen:  1,
			asm:     x86.ATESTW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:    "TESTBconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     x86.ATESTB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:         "SHLL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHLL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},   // CX
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SHLLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHLL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SHRL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},   // CX
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SHRW",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},   // CX
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SHRB",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},   // CX
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SHRLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SHRWconst",
			auxType:      auxInt16,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SHRBconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SARL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},   // CX
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SARW",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},   // CX
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SARB",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},   // CX
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SARLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SARWconst",
			auxType:      auxInt16,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SARBconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "ROLLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AROLL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "ROLWconst",
			auxType:      auxInt16,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AROLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "ROLBconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AROLB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "ADDLload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AADDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239},   // AX CX DX BX BP SI DI
					{1, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "SUBLload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.ASUBL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239},   // AX CX DX BX BP SI DI
					{1, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "ANDLload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AANDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239},   // AX CX DX BX BP SI DI
					{1, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "ORLload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239},   // AX CX DX BX BP SI DI
					{1, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "XORLload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AXORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239},   // AX CX DX BX BP SI DI
					{1, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "NEGL",
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ANEGL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "NOTL",
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ANOTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "BSFL",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ABSFL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "BSFW",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ABSFW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "BSRL",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ABSRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "BSRW",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ABSRW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "BSWAPL",
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABSWAPL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SQRTSD",
			argLen: 1,
			asm:    x86.ASQRTSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:   "SBBLcarrymask",
			argLen: 1,
			asm:    x86.ASBBL,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETEQ",
			argLen: 1,
			asm:    x86.ASETEQ,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETNE",
			argLen: 1,
			asm:    x86.ASETNE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETL",
			argLen: 1,
			asm:    x86.ASETLT,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETLE",
			argLen: 1,
			asm:    x86.ASETLE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETG",
			argLen: 1,
			asm:    x86.ASETGT,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETGE",
			argLen: 1,
			asm:    x86.ASETGE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETB",
			argLen: 1,
			asm:    x86.ASETCS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETBE",
			argLen: 1,
			asm:    x86.ASETLS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETA",
			argLen: 1,
			asm:    x86.ASETHI,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETAE",
			argLen: 1,
			asm:    x86.ASETCC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SETEQF",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ASETEQ,
			reg: regInfo{
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 238}, // CX DX BX BP SI DI
				},
			},
		},
		{
			name:         "SETNEF",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ASETNE,
			reg: regInfo{
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 238}, // CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETORD",
			argLen: 1,
			asm:    x86.ASETPC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETNAN",
			argLen: 1,
			asm:    x86.ASETPS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETGF",
			argLen: 1,
			asm:    x86.ASETHI,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "SETGEF",
			argLen: 1,
			asm:    x86.ASETCC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "MOVBLSX",
			argLen: 1,
			asm:    x86.AMOVBLSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "MOVBLZX",
			argLen: 1,
			asm:    x86.AMOVBLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "MOVWLSX",
			argLen: 1,
			asm:    x86.AMOVWLSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "MOVWLZX",
			argLen: 1,
			asm:    x86.AMOVWLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:              "MOVLconst",
			auxType:           auxInt32,
			argLen:            0,
			rematerializeable: true,
			asm:               x86.AMOVL,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:        "CVTTSD2SL",
			argLen:      1,
			usesScratch: true,
			asm:         x86.ACVTTSD2SL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:        "CVTTSS2SL",
			argLen:      1,
			usesScratch: true,
			asm:         x86.ACVTTSS2SL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:        "CVTSL2SS",
			argLen:      1,
			usesScratch: true,
			asm:         x86.ACVTSL2SS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:        "CVTSL2SD",
			argLen:      1,
			usesScratch: true,
			asm:         x86.ACVTSL2SD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:        "CVTSD2SS",
			argLen:      1,
			usesScratch: true,
			asm:         x86.ACVTSD2SS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:   "CVTSS2SD",
			argLen: 1,
			asm:    x86.ACVTSS2SD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:         "PXOR",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			asm:          x86.APXOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
					{1, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:              "LEAL",
			auxType:           auxSymOff,
			argLen:            1,
			rematerializeable: true,
			symEffect:         SymAddr,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:        "LEAL1",
			auxType:     auxSymOff,
			argLen:      2,
			commutative: true,
			symEffect:   SymAddr,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:      "LEAL2",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:      "LEAL4",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:      "LEAL8",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "MOVBload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVBLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "MOVBLSXload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVBLSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "MOVWload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVWLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "MOVWLSXload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVWLSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "MOVLload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "MOVBstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "MOVWstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "MOVLstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "ADDLmodify",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymRead | SymWrite,
			asm:            x86.AADDL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "SUBLmodify",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymRead | SymWrite,
			asm:            x86.ASUBL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "ANDLmodify",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymRead | SymWrite,
			asm:            x86.AANDL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "ORLmodify",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymRead | SymWrite,
			asm:            x86.AORL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "XORLmodify",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymRead | SymWrite,
			asm:            x86.AXORL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:        "MOVBloadidx1",
			auxType:     auxSymOff,
			argLen:      3,
			commutative: true,
			symEffect:   SymRead,
			asm:         x86.AMOVBLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:        "MOVWloadidx1",
			auxType:     auxSymOff,
			argLen:      3,
			commutative: true,
			symEffect:   SymRead,
			asm:         x86.AMOVWLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:      "MOVWloadidx2",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVWLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:        "MOVLloadidx1",
			auxType:     auxSymOff,
			argLen:      3,
			commutative: true,
			symEffect:   SymRead,
			asm:         x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:      "MOVLloadidx4",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:        "MOVBstoreidx1",
			auxType:     auxSymOff,
			argLen:      4,
			commutative: true,
			symEffect:   SymWrite,
			asm:         x86.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{2, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:        "MOVWstoreidx1",
			auxType:     auxSymOff,
			argLen:      4,
			commutative: true,
			symEffect:   SymWrite,
			asm:         x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{2, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:      "MOVWstoreidx2",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{2, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:        "MOVLstoreidx1",
			auxType:     auxSymOff,
			argLen:      4,
			commutative: true,
			symEffect:   SymWrite,
			asm:         x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{2, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:      "MOVLstoreidx4",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{2, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "MOVBstoreconst",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "MOVWstoreconst",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "MOVLstoreconst",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:      "MOVBstoreconstidx1",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:      "MOVWstoreconstidx1",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:      "MOVWstoreconstidx2",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:      "MOVLstoreconstidx1",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:      "MOVLstoreconstidx4",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 255},   // AX CX DX BX SP BP SI DI
					{0, 65791}, // AX CX DX BX SP BP SI DI SB
				},
			},
		},
		{
			name:           "DUFFZERO",
			auxType:        auxInt64,
			argLen:         3,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128}, // DI
					{1, 1},   // AX
				},
				clobbers: 130, // CX DI
			},
		},
		{
			name:           "REPSTOSL",
			argLen:         4,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128}, // DI
					{1, 2},   // CX
					{2, 1},   // AX
				},
				clobbers: 130, // CX DI
			},
		},
		{
			name:         "CALLstatic",
			auxType:      auxSymOff,
			argLen:       1,
			clobberFlags: true,
			call:         true,
			symEffect:    SymNone,
			reg: regInfo{
				clobbers: 65519, // AX CX DX BX BP SI DI X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
		{
			name:         "CALLclosure",
			auxType:      auxInt64,
			argLen:       3,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4},   // DX
					{0, 255}, // AX CX DX BX SP BP SI DI
				},
				clobbers: 65519, // AX CX DX BX BP SI DI X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
		{
			name:         "CALLinter",
			auxType:      auxInt64,
			argLen:       2,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				clobbers: 65519, // AX CX DX BX BP SI DI X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
		{
			name:           "DUFFCOPY",
			auxType:        auxInt64,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128}, // DI
					{1, 64},  // SI
				},
				clobbers: 194, // CX SI DI
			},
		},
		{
			name:           "REPMOVSL",
			argLen:         4,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128}, // DI
					{1, 64},  // SI
					{2, 2},   // CX
				},
				clobbers: 194, // CX SI DI
			},
		},
		{
			name:   "InvertFlags",
			argLen: 1,
			reg:    regInfo{},
		},
		{
			name:   "LoweredGetG",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			clobberFlags:   true,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255}, // AX CX DX BX SP BP SI DI
				},
			},
		},
		{
			name:         "LoweredWB",
			auxType:      auxSym,
			argLen:       3,
			clobberFlags: true,
			symEffect:    SymNone,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128}, // DI
					{1, 1},   // AX
				},
				clobbers: 65280, // X0 X1 X2 X3 X4 X5 X6 X7
			},
		},
		{
			name:   "FlagEQ",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagLT_ULT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagLT_UGT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagGT_UGT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagGT_ULT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FCHS",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:    "MOVSSconst1",
			auxType: auxFloat32,
			argLen:  0,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:    "MOVSDconst1",
			auxType: auxFloat64,
			argLen:  0,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
			},
		},
		{
			name:   "MOVSSconst2",
			argLen: 1,
			asm:    x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},
		{
			name:   "MOVSDconst2",
			argLen: 1,
			asm:    x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239}, // AX CX DX BX BP SI DI
				},
				outputs: []outputInfo{
					{0, 65280}, // X0 X1 X2 X3 X4 X5 X6 X7
				},
			},
		},

		{
			name:         "ADDSS",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			asm:          x86.AADDSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:         "ADDSD",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			asm:          x86.AADDSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:         "SUBSS",
			argLen:       2,
			resultInArg0: true,
			asm:          x86.ASUBSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:         "SUBSD",
			argLen:       2,
			resultInArg0: true,
			asm:          x86.ASUBSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:         "MULSS",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			asm:          x86.AMULSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:         "MULSD",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			asm:          x86.AMULSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:         "DIVSS",
			argLen:       2,
			resultInArg0: true,
			asm:          x86.ADIVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:         "DIVSD",
			argLen:       2,
			resultInArg0: true,
			asm:          x86.ADIVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:           "MOVSSload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:           "MOVSDload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:              "MOVSSconst",
			auxType:           auxFloat32,
			argLen:            0,
			rematerializeable: true,
			asm:               x86.AMOVSS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:              "MOVSDconst",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			asm:               x86.AMOVSD,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:      "MOVSSloadidx1",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:      "MOVSSloadidx4",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:      "MOVSDloadidx1",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:      "MOVSDloadidx8",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:           "MOVSSstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "MOVSDstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVSSstoreidx1",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVSSstoreidx4",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVSDstoreidx1",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVSDstoreidx8",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "ADDSSload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AADDSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:           "ADDSDload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AADDSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:           "SUBSSload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.ASUBSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:           "SUBSDload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.ASUBSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:           "MULSSload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AMULSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:           "MULSDload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AMULSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:         "ADDQ",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          x86.AADDQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ADDL",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          x86.AADDL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ADDQconst",
			auxType:      auxInt32,
			argLen:       1,
			clobberFlags: true,
			asm:          x86.AADDQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ADDLconst",
			auxType:      auxInt32,
			argLen:       1,
			clobberFlags: true,
			asm:          x86.AADDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "ADDQconstmodify",
			auxType:        auxSymValAndOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead | SymWrite,
			asm:            x86.AADDQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "ADDLconstmodify",
			auxType:        auxSymValAndOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead | SymWrite,
			asm:            x86.AADDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:         "SUBQ",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASUBQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SUBL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASUBL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SUBQconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASUBQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SUBLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASUBL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "MULQ",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AIMULQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "MULL",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AIMULL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "MULQconst",
			auxType:      auxInt32,
			argLen:       1,
			clobberFlags: true,
			asm:          x86.AIMUL3Q,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "MULLconst",
			auxType:      auxInt32,
			argLen:       1,
			clobberFlags: true,
			asm:          x86.AIMUL3L,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "HMULQ",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          x86.AIMULQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},     // AX
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:         "HMULL",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          x86.AIMULL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},     // AX
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:         "HMULQU",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          x86.AMULQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},     // AX
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:         "HMULLU",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          x86.AMULL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},     // AX
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:         "AVGQU",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "DIVQ",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.AIDIVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},     // AX
					{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 1}, // AX
					{1, 4}, // DX
				},
			},
		},
		{
			name:         "DIVL",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.AIDIVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},     // AX
					{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 1}, // AX
					{1, 4}, // DX
				},
			},
		},
		{
			name:         "DIVW",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.AIDIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},     // AX
					{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 1}, // AX
					{1, 4}, // DX
				},
			},
		},
		{
			name:         "DIVQU",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.ADIVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},     // AX
					{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 1}, // AX
					{1, 4}, // DX
				},
			},
		},
		{
			name:         "DIVLU",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.ADIVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},     // AX
					{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 1}, // AX
					{1, 4}, // DX
				},
			},
		},
		{
			name:         "DIVWU",
			argLen:       2,
			clobberFlags: true,
			asm:          x86.ADIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},     // AX
					{1, 65531}, // AX CX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 1}, // AX
					{1, 4}, // DX
				},
			},
		},
		{
			name:         "MULQU2",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          x86.AMULQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1},     // AX
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 4}, // DX
					{1, 1}, // AX
				},
			},
		},
		{
			name:         "DIVQU2",
			argLen:       3,
			clobberFlags: true,
			asm:          x86.ADIVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4},     // DX
					{1, 1},     // AX
					{2, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 1}, // AX
					{1, 4}, // DX
				},
			},
		},
		{
			name:         "ANDQ",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AANDQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ANDL",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AANDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ANDQconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AANDQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ANDLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AANDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ORQ",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AORQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ORL",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ORQconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AORQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ORLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "XORQ",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AXORQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "XORL",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AXORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "XORQconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AXORQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "XORLconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AXORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "CMPQ",
			argLen: 2,
			asm:    x86.ACMPQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "CMPL",
			argLen: 2,
			asm:    x86.ACMPL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "CMPW",
			argLen: 2,
			asm:    x86.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "CMPB",
			argLen: 2,
			asm:    x86.ACMPB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "CMPQconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     x86.ACMPQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "CMPLconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     x86.ACMPL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "CMPWconst",
			auxType: auxInt16,
			argLen:  1,
			asm:     x86.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "CMPBconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     x86.ACMPB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "CMPQload",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.ACMPQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "CMPLload",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.ACMPL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "CMPWload",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "CMPBload",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.ACMPB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "CMPQconstload",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.ACMPQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "CMPLconstload",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.ACMPL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "CMPWconstload",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "CMPBconstload",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.ACMPB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:   "UCOMISS",
			argLen: 2,
			asm:    x86.AUCOMISS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:   "UCOMISD",
			argLen: 2,
			asm:    x86.AUCOMISD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:   "BTL",
			argLen: 2,
			asm:    x86.ABTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "BTQ",
			argLen: 2,
			asm:    x86.ABTQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTCL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTCL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTCQ",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTCQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTRL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTRQ",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTRQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTSL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTSL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTSQ",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTSQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "BTLconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     x86.ABTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "BTQconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     x86.ABTQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTCLconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTCL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTCQconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTCQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTRLconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTRQconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTRQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTSLconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTSL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BTSQconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABTSQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:        "TESTQ",
			argLen:      2,
			commutative: true,
			asm:         x86.ATESTQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:        "TESTL",
			argLen:      2,
			commutative: true,
			asm:         x86.ATESTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:        "TESTW",
			argLen:      2,
			commutative: true,
			asm:         x86.ATESTW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:        "TESTB",
			argLen:      2,
			commutative: true,
			asm:         x86.ATESTB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "TESTQconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     x86.ATESTQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "TESTLconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     x86.ATESTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "TESTWconst",
			auxType: auxInt16,
			argLen:  1,
			asm:     x86.ATESTW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "TESTBconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     x86.ATESTB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHLQ",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHLQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHLL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHLL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHLQconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHLQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHLLconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHLL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHRQ",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHRL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHRW",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHRB",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHRQconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHRLconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHRWconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SHRBconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASHRB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SARQ",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SARL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SARW",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SARB",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SARQconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SARLconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SARWconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SARBconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ASARB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ROLQ",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AROLQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ROLL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AROLL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ROLW",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AROLW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ROLB",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AROLB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "RORQ",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ARORQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "RORL",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ARORL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "RORW",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ARORW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "RORB",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ARORB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // CX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ROLQconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AROLQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ROLLconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AROLL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ROLWconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AROLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "ROLBconst",
			auxType:      auxInt8,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.AROLB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "ADDLload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AADDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "ADDQload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AADDQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "SUBQload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.ASUBQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "SUBLload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.ASUBL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "ANDLload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AANDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "ANDQload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AANDQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "ORQload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AORQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "ORLload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "XORQload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AXORQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "XORLload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            x86.AXORL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "NEGQ",
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ANEGQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "NEGL",
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ANEGL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "NOTQ",
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ANOTQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "NOTL",
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ANOTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "BSFQ",
			argLen: 1,
			asm:    x86.ABSFQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BSFL",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ABSFL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "BSRQ",
			argLen: 1,
			asm:    x86.ABSRQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BSRL",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ABSRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQEQ",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQNE",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQNE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQLT",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQLT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQGT",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQLE",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQLE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQGE",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQGE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQLS",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQLS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQHI",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQHI,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQCC",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQCC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQCS",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQCS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLEQ",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLNE",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLNE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLLT",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLLT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLGT",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLLE",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLLE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLGE",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLGE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLLS",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLLS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLHI",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLHI,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLCC",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLCC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLCS",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLCS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWEQ",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWNE",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWNE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWLT",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWLT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWGT",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWLE",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWLE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWGE",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWGE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWLS",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWLS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWHI",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWHI,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWCC",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWCC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWCS",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWCS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQEQF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQNE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQNEF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQNE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQGTF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQHI,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVQGEF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVQCC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLEQF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLNE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLNEF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLNE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLGTF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLHI,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVLGEF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVLCC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWEQF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWNE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWNEF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWNE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWGTF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWHI,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "CMOVWGEF",
			argLen:       3,
			resultInArg0: true,
			asm:          x86.ACMOVWCC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BSWAPQ",
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABSWAPQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "BSWAPL",
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          x86.ABSWAPL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "POPCNTQ",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.APOPCNTQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "POPCNTL",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.APOPCNTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SQRTSD",
			argLen: 1,
			asm:    x86.ASQRTSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:    "ROUNDSD",
			auxType: auxInt8,
			argLen:  1,
			asm:     x86.AROUNDSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:   "SBBQcarrymask",
			argLen: 1,
			asm:    x86.ASBBQ,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SBBLcarrymask",
			argLen: 1,
			asm:    x86.ASBBL,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETEQ",
			argLen: 1,
			asm:    x86.ASETEQ,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETNE",
			argLen: 1,
			asm:    x86.ASETNE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETL",
			argLen: 1,
			asm:    x86.ASETLT,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETLE",
			argLen: 1,
			asm:    x86.ASETLE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETG",
			argLen: 1,
			asm:    x86.ASETGT,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETGE",
			argLen: 1,
			asm:    x86.ASETGE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETB",
			argLen: 1,
			asm:    x86.ASETCS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETBE",
			argLen: 1,
			asm:    x86.ASETLS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETA",
			argLen: 1,
			asm:    x86.ASETHI,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETAE",
			argLen: 1,
			asm:    x86.ASETCC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "SETEQstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.ASETEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "SETNEstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.ASETNE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "SETLstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.ASETLT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "SETLEstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.ASETLE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "SETGstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.ASETGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "SETGEstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.ASETGE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "SETBstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.ASETCS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "SETBEstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.ASETLS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "SETAstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.ASETHI,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "SETAEstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.ASETCC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:         "SETEQF",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ASETEQ,
			reg: regInfo{
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "SETNEF",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ASETNE,
			reg: regInfo{
				clobbers: 1, // AX
				outputs: []outputInfo{
					{0, 65518}, // CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETORD",
			argLen: 1,
			asm:    x86.ASETPC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETNAN",
			argLen: 1,
			asm:    x86.ASETPS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETGF",
			argLen: 1,
			asm:    x86.ASETHI,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "SETGEF",
			argLen: 1,
			asm:    x86.ASETCC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "MOVBQSX",
			argLen: 1,
			asm:    x86.AMOVBQSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "MOVBQZX",
			argLen: 1,
			asm:    x86.AMOVBLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "MOVWQSX",
			argLen: 1,
			asm:    x86.AMOVWQSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "MOVWQZX",
			argLen: 1,
			asm:    x86.AMOVWLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "MOVLQSX",
			argLen: 1,
			asm:    x86.AMOVLQSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "MOVLQZX",
			argLen: 1,
			asm:    x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:              "MOVLconst",
			auxType:           auxInt32,
			argLen:            0,
			rematerializeable: true,
			asm:               x86.AMOVL,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:              "MOVQconst",
			auxType:           auxInt64,
			argLen:            0,
			rematerializeable: true,
			asm:               x86.AMOVQ,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "CVTTSD2SL",
			argLen: 1,
			asm:    x86.ACVTTSD2SL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "CVTTSD2SQ",
			argLen: 1,
			asm:    x86.ACVTTSD2SQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "CVTTSS2SL",
			argLen: 1,
			asm:    x86.ACVTTSS2SL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "CVTTSS2SQ",
			argLen: 1,
			asm:    x86.ACVTTSS2SQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "CVTSL2SS",
			argLen: 1,
			asm:    x86.ACVTSL2SS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:   "CVTSL2SD",
			argLen: 1,
			asm:    x86.ACVTSL2SD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:   "CVTSQ2SS",
			argLen: 1,
			asm:    x86.ACVTSQ2SS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:   "CVTSQ2SD",
			argLen: 1,
			asm:    x86.ACVTSQ2SD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:   "CVTSD2SS",
			argLen: 1,
			asm:    x86.ACVTSD2SS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:   "CVTSS2SD",
			argLen: 1,
			asm:    x86.ACVTSS2SD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:   "MOVQi2f",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:   "MOVQf2i",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "MOVLi2f",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:   "MOVLf2i",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "PXOR",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			asm:          x86.APXOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:              "LEAQ",
			auxType:           auxSymOff,
			argLen:            1,
			rematerializeable: true,
			symEffect:         SymAddr,
			asm:               x86.ALEAQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:              "LEAL",
			auxType:           auxSymOff,
			argLen:            1,
			rematerializeable: true,
			symEffect:         SymAddr,
			asm:               x86.ALEAL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:              "LEAW",
			auxType:           auxSymOff,
			argLen:            1,
			rematerializeable: true,
			symEffect:         SymAddr,
			asm:               x86.ALEAW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:        "LEAQ1",
			auxType:     auxSymOff,
			argLen:      2,
			commutative: true,
			symEffect:   SymAddr,
			asm:         x86.ALEAQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:        "LEAL1",
			auxType:     auxSymOff,
			argLen:      2,
			commutative: true,
			symEffect:   SymAddr,
			asm:         x86.ALEAL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:        "LEAW1",
			auxType:     auxSymOff,
			argLen:      2,
			commutative: true,
			symEffect:   SymAddr,
			asm:         x86.ALEAW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "LEAQ2",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			asm:       x86.ALEAQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "LEAL2",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			asm:       x86.ALEAL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "LEAW2",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			asm:       x86.ALEAW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "LEAQ4",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			asm:       x86.ALEAQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "LEAL4",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			asm:       x86.ALEAL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "LEAW4",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			asm:       x86.ALEAW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "LEAQ8",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			asm:       x86.ALEAQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "LEAL8",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			asm:       x86.ALEAL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "LEAW8",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymAddr,
			asm:       x86.ALEAW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "MOVBload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVBLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "MOVBQSXload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVBQSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "MOVWload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVWLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "MOVWQSXload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVWQSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "MOVLload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "MOVLQSXload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVLQSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "MOVQload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "MOVBstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "MOVWstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "MOVLstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "MOVQstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "MOVOload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVUPS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:           "MOVOstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVUPS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:        "MOVBloadidx1",
			auxType:     auxSymOff,
			argLen:      3,
			commutative: true,
			symEffect:   SymRead,
			asm:         x86.AMOVBLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:        "MOVWloadidx1",
			auxType:     auxSymOff,
			argLen:      3,
			commutative: true,
			symEffect:   SymRead,
			asm:         x86.AMOVWLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "MOVWloadidx2",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVWLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:        "MOVLloadidx1",
			auxType:     auxSymOff,
			argLen:      3,
			commutative: true,
			symEffect:   SymRead,
			asm:         x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "MOVLloadidx4",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "MOVLloadidx8",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:        "MOVQloadidx1",
			auxType:     auxSymOff,
			argLen:      3,
			commutative: true,
			symEffect:   SymRead,
			asm:         x86.AMOVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "MOVQloadidx8",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       x86.AMOVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "MOVBstoreidx1",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVWstoreidx1",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVWstoreidx2",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVLstoreidx1",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVLstoreidx4",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVLstoreidx8",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVQstoreidx1",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVQstoreidx8",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       x86.AMOVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "MOVBstoreconst",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "MOVWstoreconst",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "MOVLstoreconst",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "MOVQstoreconst",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            x86.AMOVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVBstoreconstidx1",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVWstoreconstidx1",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVWstoreconstidx2",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVLstoreconstidx1",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVLstoreconstidx4",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVQstoreconstidx1",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:      "MOVQstoreconstidx8",
			auxType:   auxSymValAndOff,
			argLen:    3,
			symEffect: SymWrite,
			asm:       x86.AMOVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "DUFFZERO",
			auxType:        auxInt64,
			argLen:         3,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128},   // DI
					{1, 65536}, // X0
				},
				clobbers: 128, // DI
			},
		},
		{
			name:              "MOVOconst",
			auxType:           auxInt128,
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4294901760}, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
				},
			},
		},
		{
			name:           "REPSTOSQ",
			argLen:         4,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128}, // DI
					{1, 2},   // CX
					{2, 1},   // AX
				},
				clobbers: 130, // CX DI
			},
		},
		{
			name:         "CALLstatic",
			auxType:      auxSymOff,
			argLen:       1,
			clobberFlags: true,
			call:         true,
			symEffect:    SymNone,
			reg: regInfo{
				clobbers: 4294967279, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
		{
			name:         "CALLclosure",
			auxType:      auxInt64,
			argLen:       3,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4},     // DX
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				clobbers: 4294967279, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
		{
			name:         "CALLinter",
			auxType:      auxInt64,
			argLen:       2,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				clobbers: 4294967279, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
		{
			name:           "DUFFCOPY",
			auxType:        auxInt64,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128}, // DI
					{1, 64},  // SI
				},
				clobbers: 65728, // SI DI X0
			},
		},
		{
			name:           "REPMOVSQ",
			argLen:         4,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128}, // DI
					{1, 64},  // SI
					{2, 2},   // CX
				},
				clobbers: 194, // CX SI DI
			},
		},
		{
			name:   "InvertFlags",
			argLen: 1,
			reg:    regInfo{},
		},
		{
			name:   "LoweredGetG",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4}, // DX
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			clobberFlags:   true,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:         "LoweredWB",
			auxType:      auxSym,
			argLen:       3,
			clobberFlags: true,
			symEffect:    SymNone,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128}, // DI
					{1, 1},   // AX
				},
				clobbers: 4294901760, // X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15
			},
		},
		{
			name:   "FlagEQ",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagLT_ULT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagLT_UGT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagGT_UGT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagGT_ULT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:           "MOVLatomicload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "MOVQatomicload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            x86.AMOVQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "XCHGL",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            x86.AXCHGL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "XCHGQ",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			faultOnNilArg1: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            x86.AXCHGQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "XADDLlock",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            x86.AXADDL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "XADDQlock",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            x86.AXADDQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},      // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{1, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
				outputs: []outputInfo{
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "AddTupleFirst32",
			argLen: 2,
			reg:    regInfo{},
		},
		{
			name:   "AddTupleFirst64",
			argLen: 2,
			reg:    regInfo{},
		},
		{
			name:           "CMPXCHGLlock",
			auxType:        auxSymOff,
			argLen:         4,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            x86.ACMPXCHGL,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 1},     // AX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{1, 0},
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "CMPXCHGQlock",
			auxType:        auxSymOff,
			argLen:         4,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            x86.ACMPXCHGQ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 1},     // AX
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{2, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
				clobbers: 1, // AX
				outputs: []outputInfo{
					{1, 0},
					{0, 65519}, // AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "ANDBlock",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            x86.AANDB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},
		{
			name:           "ORBlock",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            x86.AORB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 65535},      // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15
					{0, 4295032831}, // AX CX DX BX SP BP SI DI R8 R9 R10 R11 R12 R13 R14 R15 SB
				},
			},
		},

		{
			name:        "ADD",
			argLen:      2,
			commutative: true,
			asm:         arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADDconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 30719}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SUB",
			argLen: 2,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SUBconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "RSB",
			argLen: 2,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSBconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:        "MUL",
			argLen:      2,
			commutative: true,
			asm:         arm.AMUL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:        "HMUL",
			argLen:      2,
			commutative: true,
			asm:         arm.AMULL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:        "HMULU",
			argLen:      2,
			commutative: true,
			asm:         arm.AMULLU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "CALLudiv",
			argLen:       2,
			clobberFlags: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 2}, // R1
					{1, 1}, // R0
				},
				clobbers: 16396, // R2 R3 R14
				outputs: []outputInfo{
					{0, 1}, // R0
					{1, 2}, // R1
				},
			},
		},
		{
			name:        "ADDS",
			argLen:      2,
			commutative: true,
			asm:         arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADDSconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:        "ADC",
			argLen:      3,
			commutative: true,
			asm:         arm.AADC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADCconst",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AADC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SUBS",
			argLen: 2,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SUBSconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSBSconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SBC",
			argLen: 3,
			asm:    arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SBCconst",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSCconst",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ARSC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:        "MULLU",
			argLen:      2,
			commutative: true,
			asm:         arm.AMULLU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MULA",
			argLen: 3,
			asm:    arm.AMULA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MULS",
			argLen: 3,
			asm:    arm.AMULS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:        "ADDF",
			argLen:      2,
			commutative: true,
			asm:         arm.AADDF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:        "ADDD",
			argLen:      2,
			commutative: true,
			asm:         arm.AADDD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "SUBF",
			argLen: 2,
			asm:    arm.ASUBF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "SUBD",
			argLen: 2,
			asm:    arm.ASUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:        "MULF",
			argLen:      2,
			commutative: true,
			asm:         arm.AMULF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:        "MULD",
			argLen:      2,
			commutative: true,
			asm:         arm.AMULD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:        "NMULF",
			argLen:      2,
			commutative: true,
			asm:         arm.ANMULF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:        "NMULD",
			argLen:      2,
			commutative: true,
			asm:         arm.ANMULD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "DIVF",
			argLen: 2,
			asm:    arm.ADIVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "DIVD",
			argLen: 2,
			asm:    arm.ADIVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "MULAF",
			argLen:       3,
			resultInArg0: true,
			asm:          arm.AMULAF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "MULAD",
			argLen:       3,
			resultInArg0: true,
			asm:          arm.AMULAD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "MULSF",
			argLen:       3,
			resultInArg0: true,
			asm:          arm.AMULSF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "MULSD",
			argLen:       3,
			resultInArg0: true,
			asm:          arm.AMULSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:        "AND",
			argLen:      2,
			commutative: true,
			asm:         arm.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ANDconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:        "OR",
			argLen:      2,
			commutative: true,
			asm:         arm.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ORconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:        "XOR",
			argLen:      2,
			commutative: true,
			asm:         arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "XORconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "BIC",
			argLen: 2,
			asm:    arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "BICconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "BFX",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ABFX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "BFXU",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ABFXU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MVN",
			argLen: 1,
			asm:    arm.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "NEGF",
			argLen: 1,
			asm:    arm.ANEGF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "NEGD",
			argLen: 1,
			asm:    arm.ANEGD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "SQRTD",
			argLen: 1,
			asm:    arm.ASQRTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "CLZ",
			argLen: 1,
			asm:    arm.ACLZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "REV",
			argLen: 1,
			asm:    arm.AREV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "RBIT",
			argLen: 1,
			asm:    arm.ARBIT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SLL",
			argLen: 2,
			asm:    arm.ASLL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SLLconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ASLL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SRL",
			argLen: 2,
			asm:    arm.ASRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SRLconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ASRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SRA",
			argLen: 2,
			asm:    arm.ASRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SRAconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ASRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SRRconst",
			auxType: auxInt32,
			argLen:  1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADDshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADDshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADDshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SUBshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SUBshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SUBshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSBshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSBshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSBshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ANDshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ANDshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ANDshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ORshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ORshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ORshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "XORshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "XORshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "XORshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "XORshiftRR",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "BICshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "BICshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "BICshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "MVNshiftLL",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "MVNshiftRL",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "MVNshiftRA",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADCshiftLL",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.AADC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADCshiftRL",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.AADC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADCshiftRA",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.AADC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SBCshiftLL",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SBCshiftRL",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SBCshiftRA",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSCshiftLL",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.ARSC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSCshiftRL",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.ARSC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSCshiftRA",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.ARSC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADDSshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADDSshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "ADDSshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SUBSshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SUBSshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "SUBSshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSBSshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSBSshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "RSBSshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ADDshiftLLreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ADDshiftRLreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ADDshiftRAreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SUBshiftLLreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SUBshiftRLreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SUBshiftRAreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "RSBshiftLLreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "RSBshiftRLreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "RSBshiftRAreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ANDshiftLLreg",
			argLen: 3,
			asm:    arm.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ANDshiftRLreg",
			argLen: 3,
			asm:    arm.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ANDshiftRAreg",
			argLen: 3,
			asm:    arm.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ORshiftLLreg",
			argLen: 3,
			asm:    arm.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ORshiftRLreg",
			argLen: 3,
			asm:    arm.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ORshiftRAreg",
			argLen: 3,
			asm:    arm.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "XORshiftLLreg",
			argLen: 3,
			asm:    arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "XORshiftRLreg",
			argLen: 3,
			asm:    arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "XORshiftRAreg",
			argLen: 3,
			asm:    arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "BICshiftLLreg",
			argLen: 3,
			asm:    arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "BICshiftRLreg",
			argLen: 3,
			asm:    arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "BICshiftRAreg",
			argLen: 3,
			asm:    arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MVNshiftLLreg",
			argLen: 2,
			asm:    arm.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MVNshiftRLreg",
			argLen: 2,
			asm:    arm.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MVNshiftRAreg",
			argLen: 2,
			asm:    arm.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ADCshiftLLreg",
			argLen: 4,
			asm:    arm.AADC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ADCshiftRLreg",
			argLen: 4,
			asm:    arm.AADC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ADCshiftRAreg",
			argLen: 4,
			asm:    arm.AADC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SBCshiftLLreg",
			argLen: 4,
			asm:    arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SBCshiftRLreg",
			argLen: 4,
			asm:    arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SBCshiftRAreg",
			argLen: 4,
			asm:    arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "RSCshiftLLreg",
			argLen: 4,
			asm:    arm.ARSC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "RSCshiftRLreg",
			argLen: 4,
			asm:    arm.ARSC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "RSCshiftRAreg",
			argLen: 4,
			asm:    arm.ARSC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ADDSshiftLLreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ADDSshiftRLreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "ADDSshiftRAreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SUBSshiftLLreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SUBSshiftRLreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SUBSshiftRAreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "RSBSshiftLLreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "RSBSshiftRLreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "RSBSshiftRAreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "CMP",
			argLen: 2,
			asm:    arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "CMPconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:        "CMN",
			argLen:      2,
			commutative: true,
			asm:         arm.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "CMNconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:        "TST",
			argLen:      2,
			commutative: true,
			asm:         arm.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "TSTconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:        "TEQ",
			argLen:      2,
			commutative: true,
			asm:         arm.ATEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "TEQconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm.ATEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:   "CMPF",
			argLen: 2,
			asm:    arm.ACMPF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "CMPD",
			argLen: 2,
			asm:    arm.ACMPD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:    "CMPshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "CMPshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "CMPshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "CMNshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "CMNshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "CMNshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "TSTshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "TSTshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "TSTshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "TEQshiftLL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ATEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "TEQshiftRL",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ATEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:    "TEQshiftRA",
			auxType: auxInt32,
			argLen:  2,
			asm:     arm.ATEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{1, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:   "CMPshiftLLreg",
			argLen: 3,
			asm:    arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "CMPshiftRLreg",
			argLen: 3,
			asm:    arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "CMPshiftRAreg",
			argLen: 3,
			asm:    arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "CMNshiftLLreg",
			argLen: 3,
			asm:    arm.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "CMNshiftRLreg",
			argLen: 3,
			asm:    arm.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "CMNshiftRAreg",
			argLen: 3,
			asm:    arm.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "TSTshiftLLreg",
			argLen: 3,
			asm:    arm.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "TSTshiftRLreg",
			argLen: 3,
			asm:    arm.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "TSTshiftRAreg",
			argLen: 3,
			asm:    arm.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "TEQshiftLLreg",
			argLen: 3,
			asm:    arm.ATEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "TEQshiftRLreg",
			argLen: 3,
			asm:    arm.ATEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "TEQshiftRAreg",
			argLen: 3,
			asm:    arm.ATEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "CMPF0",
			argLen: 1,
			asm:    arm.ACMPF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "CMPD0",
			argLen: 1,
			asm:    arm.ACMPD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:              "MOVWconst",
			auxType:           auxInt32,
			argLen:            0,
			rematerializeable: true,
			asm:               arm.AMOVW,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:              "MOVFconst",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			asm:               arm.AMOVF,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:              "MOVDconst",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			asm:               arm.AMOVD,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:              "MOVWaddr",
			auxType:           auxSymOff,
			argLen:            1,
			rematerializeable: true,
			symEffect:         SymAddr,
			asm:               arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294975488}, // SP SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:           "MOVBload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:           "MOVBUload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:           "MOVHload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:           "MOVHUload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:           "MOVWload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:           "MOVFload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm.AMOVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:           "MOVDload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:           "MOVBstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
			},
		},
		{
			name:           "MOVHstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
			},
		},
		{
			name:           "MOVWstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
			},
		},
		{
			name:           "MOVFstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm.AMOVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:           "MOVDstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "MOVWloadidx",
			argLen: 3,
			asm:    arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "MOVWloadshiftLL",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "MOVWloadshiftRL",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:    "MOVWloadshiftRA",
			auxType: auxInt32,
			argLen:  3,
			asm:     arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVBUloadidx",
			argLen: 3,
			asm:    arm.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVBloadidx",
			argLen: 3,
			asm:    arm.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVHUloadidx",
			argLen: 3,
			asm:    arm.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVHloadidx",
			argLen: 3,
			asm:    arm.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVWstoreidx",
			argLen: 4,
			asm:    arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
			},
		},
		{
			name:    "MOVWstoreshiftLL",
			auxType: auxInt32,
			argLen:  4,
			asm:     arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
			},
		},
		{
			name:    "MOVWstoreshiftRL",
			auxType: auxInt32,
			argLen:  4,
			asm:     arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
			},
		},
		{
			name:    "MOVWstoreshiftRA",
			auxType: auxInt32,
			argLen:  4,
			asm:     arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
			},
		},
		{
			name:   "MOVBstoreidx",
			argLen: 4,
			asm:    arm.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
			},
		},
		{
			name:   "MOVHstoreidx",
			argLen: 4,
			asm:    arm.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{2, 22527},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
					{0, 4294998015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 SP R14 SB
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    arm.AMOVBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVBUreg",
			argLen: 1,
			asm:    arm.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    arm.AMOVHS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVHUreg",
			argLen: 1,
			asm:    arm.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "MOVWnop",
			argLen:       1,
			resultInArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVWF",
			argLen: 1,
			asm:    arm.AMOVWF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2147483648, // F15
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "MOVWD",
			argLen: 1,
			asm:    arm.AMOVWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2147483648, // F15
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "MOVWUF",
			argLen: 1,
			asm:    arm.AMOVWF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2147483648, // F15
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "MOVWUD",
			argLen: 1,
			asm:    arm.AMOVWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2147483648, // F15
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "MOVFW",
			argLen: 1,
			asm:    arm.AMOVFW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				clobbers: 2147483648, // F15
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVDW",
			argLen: 1,
			asm:    arm.AMOVDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				clobbers: 2147483648, // F15
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVFWU",
			argLen: 1,
			asm:    arm.AMOVFW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				clobbers: 2147483648, // F15
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVDWU",
			argLen: 1,
			asm:    arm.AMOVDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				clobbers: 2147483648, // F15
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "MOVFD",
			argLen: 1,
			asm:    arm.AMOVFD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "MOVDF",
			argLen: 1,
			asm:    arm.AMOVDF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "CMOVWHSconst",
			auxType:      auxInt32,
			argLen:       2,
			resultInArg0: true,
			asm:          arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "CMOVWLSconst",
			auxType:      auxInt32,
			argLen:       2,
			resultInArg0: true,
			asm:          arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "SRAcond",
			argLen: 3,
			asm:    arm.ASRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "CALLstatic",
			auxType:      auxSymOff,
			argLen:       1,
			clobberFlags: true,
			call:         true,
			symEffect:    SymNone,
			reg: regInfo{
				clobbers: 4294924287, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
		{
			name:         "CALLclosure",
			auxType:      auxInt64,
			argLen:       3,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 128},   // R7
					{0, 29695}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 SP R14
				},
				clobbers: 4294924287, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
		{
			name:         "CALLinter",
			auxType:      auxInt64,
			argLen:       2,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 4294924287, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 g R12 R14
				},
			},
		},
		{
			name:   "Equal",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "NotEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "LessThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "LessEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "GreaterThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "GreaterEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "LessThanU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "LessEqualU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "GreaterThanU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "GreaterEqualU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:           "DUFFZERO",
			auxType:        auxInt64,
			argLen:         3,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 2}, // R1
					{1, 1}, // R0
				},
				clobbers: 16386, // R1 R14
			},
		},
		{
			name:           "DUFFCOPY",
			auxType:        auxInt64,
			argLen:         3,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4}, // R2
					{1, 2}, // R1
				},
				clobbers: 16391, // R0 R1 R2 R14
			},
		},
		{
			name:           "LoweredZero",
			auxType:        auxInt64,
			argLen:         4,
			clobberFlags:   true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 2},     // R1
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2, // R1
			},
		},
		{
			name:           "LoweredMove",
			auxType:        auxInt64,
			argLen:         4,
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4},     // R2
					{1, 2},     // R1
					{2, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 6, // R1 R2
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 128}, // R7
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:   "FlagEQ",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagLT_ULT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagLT_UGT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagGT_UGT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagGT_ULT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "InvertFlags",
			argLen: 1,
			reg:    regInfo{},
		},
		{
			name:         "LoweredWB",
			auxType:      auxSym,
			argLen:       3,
			clobberFlags: true,
			symEffect:    SymNone,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4}, // R2
					{1, 8}, // R3
				},
				clobbers: 4294918144, // R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},

		{
			name:        "ADD",
			argLen:      2,
			commutative: true,
			asm:         arm64.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ADDconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1878786047}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "SUB",
			argLen: 2,
			asm:    arm64.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "SUBconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:        "MUL",
			argLen:      2,
			commutative: true,
			asm:         arm64.AMUL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:        "MULW",
			argLen:      2,
			commutative: true,
			asm:         arm64.AMULW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:        "MNEG",
			argLen:      2,
			commutative: true,
			asm:         arm64.AMNEG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:        "MNEGW",
			argLen:      2,
			commutative: true,
			asm:         arm64.AMNEGW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:        "MULH",
			argLen:      2,
			commutative: true,
			asm:         arm64.ASMULH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:        "UMULH",
			argLen:      2,
			commutative: true,
			asm:         arm64.AUMULH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:        "MULL",
			argLen:      2,
			commutative: true,
			asm:         arm64.ASMULL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:        "UMULL",
			argLen:      2,
			commutative: true,
			asm:         arm64.AUMULL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "DIV",
			argLen: 2,
			asm:    arm64.ASDIV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "UDIV",
			argLen: 2,
			asm:    arm64.AUDIV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "DIVW",
			argLen: 2,
			asm:    arm64.ASDIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "UDIVW",
			argLen: 2,
			asm:    arm64.AUDIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOD",
			argLen: 2,
			asm:    arm64.AREM,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "UMOD",
			argLen: 2,
			asm:    arm64.AUREM,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MODW",
			argLen: 2,
			asm:    arm64.AREMW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "UMODW",
			argLen: 2,
			asm:    arm64.AUREMW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:        "FADDS",
			argLen:      2,
			commutative: true,
			asm:         arm64.AFADDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:        "FADDD",
			argLen:      2,
			commutative: true,
			asm:         arm64.AFADDD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FSUBS",
			argLen: 2,
			asm:    arm64.AFSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FSUBD",
			argLen: 2,
			asm:    arm64.AFSUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:        "FMULS",
			argLen:      2,
			commutative: true,
			asm:         arm64.AFMULS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:        "FMULD",
			argLen:      2,
			commutative: true,
			asm:         arm64.AFMULD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:        "FNMULS",
			argLen:      2,
			commutative: true,
			asm:         arm64.AFNMULS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:        "FNMULD",
			argLen:      2,
			commutative: true,
			asm:         arm64.AFNMULD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FDIVS",
			argLen: 2,
			asm:    arm64.AFDIVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FDIVD",
			argLen: 2,
			asm:    arm64.AFDIVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:        "AND",
			argLen:      2,
			commutative: true,
			asm:         arm64.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ANDconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:        "OR",
			argLen:      2,
			commutative: true,
			asm:         arm64.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ORconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:        "XOR",
			argLen:      2,
			commutative: true,
			asm:         arm64.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "XORconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "BIC",
			argLen: 2,
			asm:    arm64.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "EON",
			argLen: 2,
			asm:    arm64.AEON,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "ORN",
			argLen: 2,
			asm:    arm64.AORN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:            "LoweredMuluhilo",
			argLen:          2,
			resultNotInArgs: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
					{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MVN",
			argLen: 1,
			asm:    arm64.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "NEG",
			argLen: 1,
			asm:    arm64.ANEG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "FNEGS",
			argLen: 1,
			asm:    arm64.AFNEGS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FNEGD",
			argLen: 1,
			asm:    arm64.AFNEGD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FSQRTD",
			argLen: 1,
			asm:    arm64.AFSQRTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "REV",
			argLen: 1,
			asm:    arm64.AREV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "REVW",
			argLen: 1,
			asm:    arm64.AREVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "REV16W",
			argLen: 1,
			asm:    arm64.AREV16W,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "RBIT",
			argLen: 1,
			asm:    arm64.ARBIT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "RBITW",
			argLen: 1,
			asm:    arm64.ARBITW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "CLZ",
			argLen: 1,
			asm:    arm64.ACLZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "CLZW",
			argLen: 1,
			asm:    arm64.ACLZW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "VCNT",
			argLen: 1,
			asm:    arm64.AVCNT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "VUADDLV",
			argLen: 1,
			asm:    arm64.AVUADDLV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:         "LoweredRound32F",
			argLen:       1,
			resultInArg0: true,
			zeroWidth:    true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:         "LoweredRound64F",
			argLen:       1,
			resultInArg0: true,
			zeroWidth:    true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FMADDS",
			argLen: 3,
			asm:    arm64.AFMADDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FMADDD",
			argLen: 3,
			asm:    arm64.AFMADDD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FNMADDS",
			argLen: 3,
			asm:    arm64.AFNMADDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FNMADDD",
			argLen: 3,
			asm:    arm64.AFNMADDD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FMSUBS",
			argLen: 3,
			asm:    arm64.AFMSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FMSUBD",
			argLen: 3,
			asm:    arm64.AFMSUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FNMSUBS",
			argLen: 3,
			asm:    arm64.AFNMSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FNMSUBD",
			argLen: 3,
			asm:    arm64.AFNMSUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{2, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "SLL",
			argLen: 2,
			asm:    arm64.ALSL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "SLLconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.ALSL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "SRL",
			argLen: 2,
			asm:    arm64.ALSR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "SRLconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.ALSR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "SRA",
			argLen: 2,
			asm:    arm64.AASR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "SRAconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.AASR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "RORconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.AROR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "RORWconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.ARORW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "EXTRconst",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AEXTR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "EXTRWconst",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AEXTRW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "CMP",
			argLen: 2,
			asm:    arm64.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:    "CMPconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:   "CMPW",
			argLen: 2,
			asm:    arm64.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:    "CMPWconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm64.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:   "CMN",
			argLen: 2,
			asm:    arm64.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:    "CMNconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:   "CMNW",
			argLen: 2,
			asm:    arm64.ACMNW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:    "CMNWconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm64.ACMNW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:   "TST",
			argLen: 2,
			asm:    arm64.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:    "TSTconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:   "TSTW",
			argLen: 2,
			asm:    arm64.ATSTW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:    "TSTWconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     arm64.ATSTW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:   "FCMPS",
			argLen: 2,
			asm:    arm64.AFCMPS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FCMPD",
			argLen: 2,
			asm:    arm64.AFCMPD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:    "ADDshiftLL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ADDshiftRL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ADDshiftRA",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "SUBshiftLL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "SUBshiftRL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "SUBshiftRA",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ANDshiftLL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ANDshiftRL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ANDshiftRA",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ORshiftLL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ORshiftRL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ORshiftRA",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "XORshiftLL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "XORshiftRL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "XORshiftRA",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "BICshiftLL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "BICshiftRL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "BICshiftRA",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "EONshiftLL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AEON,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "EONshiftRL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AEON,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "EONshiftRA",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AEON,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ORNshiftLL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AORN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ORNshiftRL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AORN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "ORNshiftRA",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.AORN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "CMPshiftLL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:    "CMPshiftRL",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:    "CMPshiftRA",
			auxType: auxInt64,
			argLen:  2,
			asm:     arm64.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{1, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:         "BFI",
			auxType:      auxInt64,
			argLen:       2,
			resultInArg0: true,
			asm:          arm64.ABFI,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
					{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:         "BFXIL",
			auxType:      auxInt64,
			argLen:       2,
			resultInArg0: true,
			asm:          arm64.ABFXIL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
					{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "SBFIZ",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.ASBFIZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "SBFX",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.ASBFX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "UBFIZ",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.AUBFIZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "UBFX",
			auxType: auxInt64,
			argLen:  1,
			asm:     arm64.AUBFX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:              "MOVDconst",
			auxType:           auxInt64,
			argLen:            0,
			rematerializeable: true,
			asm:               arm64.AMOVD,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:              "FMOVSconst",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			asm:               arm64.AFMOVS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:              "FMOVDconst",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			asm:               arm64.AFMOVD,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:              "MOVDaddr",
			auxType:           auxSymOff,
			argLen:            1,
			rematerializeable: true,
			symEffect:         SymAddr,
			asm:               arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372037928517632}, // SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "MOVBload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "MOVBUload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm64.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "MOVHload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "MOVHUload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm64.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "MOVWload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "MOVWUload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm64.AMOVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "MOVDload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "FMOVSload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm64.AFMOVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:           "FMOVDload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            arm64.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "MOVDloadidx",
			argLen: 3,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVWloadidx",
			argLen: 3,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVWUloadidx",
			argLen: 3,
			asm:    arm64.AMOVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVHloadidx",
			argLen: 3,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVHUloadidx",
			argLen: 3,
			asm:    arm64.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVBloadidx",
			argLen: 3,
			asm:    arm64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVBUloadidx",
			argLen: 3,
			asm:    arm64.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVHloadidx2",
			argLen: 3,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVHUloadidx2",
			argLen: 3,
			asm:    arm64.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVWloadidx4",
			argLen: 3,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVWUloadidx4",
			argLen: 3,
			asm:    arm64.AMOVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVDloadidx8",
			argLen: 3,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "MOVBstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:           "MOVHstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:           "MOVWstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:           "MOVDstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:           "STP",
			auxType:        auxSymOff,
			argLen:         4,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.ASTP,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:           "FMOVSstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.AFMOVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:           "FMOVDstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
					{1, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "MOVBstoreidx",
			argLen: 4,
			asm:    arm64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVHstoreidx",
			argLen: 4,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVWstoreidx",
			argLen: 4,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVDstoreidx",
			argLen: 4,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVHstoreidx2",
			argLen: 4,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVWstoreidx4",
			argLen: 4,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVDstoreidx8",
			argLen: 4,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:           "MOVBstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:           "MOVHstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:           "MOVWstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:           "MOVDstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:           "MOVQstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            arm64.ASTP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVBstorezeroidx",
			argLen: 3,
			asm:    arm64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVHstorezeroidx",
			argLen: 3,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVWstorezeroidx",
			argLen: 3,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVDstorezeroidx",
			argLen: 3,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVHstorezeroidx2",
			argLen: 3,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVWstorezeroidx4",
			argLen: 3,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "MOVDstorezeroidx8",
			argLen: 3,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:   "FMOVDgpfp",
			argLen: 1,
			asm:    arm64.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FMOVDfpgp",
			argLen: 1,
			asm:    arm64.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    arm64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVBUreg",
			argLen: 1,
			asm:    arm64.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVHUreg",
			argLen: 1,
			asm:    arm64.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVWUreg",
			argLen: 1,
			asm:    arm64.AMOVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "MOVDreg",
			argLen: 1,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:         "MOVDnop",
			argLen:       1,
			resultInArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "SCVTFWS",
			argLen: 1,
			asm:    arm64.ASCVTFWS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "SCVTFWD",
			argLen: 1,
			asm:    arm64.ASCVTFWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "UCVTFWS",
			argLen: 1,
			asm:    arm64.AUCVTFWS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "UCVTFWD",
			argLen: 1,
			asm:    arm64.AUCVTFWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "SCVTFS",
			argLen: 1,
			asm:    arm64.ASCVTFS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "SCVTFD",
			argLen: 1,
			asm:    arm64.ASCVTFD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "UCVTFS",
			argLen: 1,
			asm:    arm64.AUCVTFS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "UCVTFD",
			argLen: 1,
			asm:    arm64.AUCVTFD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FCVTZSSW",
			argLen: 1,
			asm:    arm64.AFCVTZSSW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "FCVTZSDW",
			argLen: 1,
			asm:    arm64.AFCVTZSDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "FCVTZUSW",
			argLen: 1,
			asm:    arm64.AFCVTZUSW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "FCVTZUDW",
			argLen: 1,
			asm:    arm64.AFCVTZUDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "FCVTZSS",
			argLen: 1,
			asm:    arm64.AFCVTZSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "FCVTZSD",
			argLen: 1,
			asm:    arm64.AFCVTZSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "FCVTZUS",
			argLen: 1,
			asm:    arm64.AFCVTZUS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "FCVTZUD",
			argLen: 1,
			asm:    arm64.AFCVTZUD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "FCVTSD",
			argLen: 1,
			asm:    arm64.AFCVTSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FCVTDS",
			argLen: 1,
			asm:    arm64.AFCVTDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FRINTAD",
			argLen: 1,
			asm:    arm64.AFRINTAD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FRINTMD",
			argLen: 1,
			asm:    arm64.AFRINTMD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FRINTPD",
			argLen: 1,
			asm:    arm64.AFRINTPD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "FRINTZD",
			argLen: 1,
			asm:    arm64.AFRINTZD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 9223372034707292160}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:    "CSEL",
			auxType: auxCCop,
			argLen:  3,
			asm:     arm64.ACSEL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
					{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:    "CSEL0",
			auxType: auxCCop,
			argLen:  2,
			asm:     arm64.ACSEL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:         "CALLstatic",
			auxType:      auxSymOff,
			argLen:       1,
			clobberFlags: true,
			call:         true,
			symEffect:    SymNone,
			reg: regInfo{
				clobbers: 9223372035512336383, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
		{
			name:         "CALLclosure",
			auxType:      auxInt64,
			argLen:       3,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 67108864},   // R26
					{0, 1744568319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30 SP
				},
				clobbers: 9223372035512336383, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
		{
			name:         "CALLinter",
			auxType:      auxInt64,
			argLen:       2,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				clobbers: 9223372035512336383, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
				},
			},
		},
		{
			name:   "Equal",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "NotEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "LessThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "LessEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "GreaterThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "GreaterEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "LessThanU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "LessEqualU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "GreaterThanU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "GreaterEqualU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "DUFFZERO",
			auxType:        auxInt64,
			argLen:         2,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65536}, // R16
				},
				clobbers: 536936448, // R16 R30
			},
		},
		{
			name:           "LoweredZero",
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65536},     // R16
					{1, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				clobbers: 65536, // R16
			},
		},
		{
			name:           "DUFFCOPY",
			auxType:        auxInt64,
			argLen:         3,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 131072}, // R17
					{1, 65536},  // R16
				},
				clobbers: 604176384, // R16 R17 R26 R30
			},
		},
		{
			name:           "LoweredMove",
			argLen:         4,
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 131072},    // R17
					{1, 65536},     // R16
					{2, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
				clobbers: 196608, // R16 R17
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 67108864}, // R26
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:   "FlagEQ",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagLT_ULT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagLT_UGT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagGT_UGT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagGT_ULT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "InvertFlags",
			argLen: 1,
			reg:    regInfo{},
		},
		{
			name:           "LDAR",
			argLen:         2,
			faultOnNilArg0: true,
			asm:            arm64.ALDAR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "LDARW",
			argLen:         2,
			faultOnNilArg0: true,
			asm:            arm64.ALDARW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:           "STLR",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			asm:            arm64.ASTLR,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:           "STLRW",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			asm:            arm64.ASTLRW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
			},
		},
		{
			name:            "LoweredAtomicExchange64",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:            "LoweredAtomicExchange32",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:            "LoweredAtomicAdd64",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:            "LoweredAtomicAdd32",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:            "LoweredAtomicAdd64Variant",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:            "LoweredAtomicAdd32Variant",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:            "LoweredAtomicCas64",
			argLen:          4,
			resultNotInArgs: true,
			clobberFlags:    true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:            "LoweredAtomicCas32",
			argLen:          4,
			resultNotInArgs: true,
			clobberFlags:    true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{2, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:            "LoweredAtomicAnd8",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			asm:             arm64.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:            "LoweredAtomicOr8",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			asm:             arm64.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},           // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30
					{0, 9223372038733561855}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 g R30 SP SB
				},
				outputs: []outputInfo{
					{0, 670826495}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30
				},
			},
		},
		{
			name:         "LoweredWB",
			auxType:      auxSym,
			argLen:       3,
			clobberFlags: true,
			symEffect:    SymNone,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4}, // R2
					{1, 8}, // R3
				},
				clobbers: 9223372035244163072, // R30 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
			},
		},

		{
			name:        "ADD",
			argLen:      2,
			commutative: true,
			asm:         mips.AADDU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:    "ADDconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     mips.AADDU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 536870910}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "SUB",
			argLen: 2,
			asm:    mips.ASUBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:    "SUBconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     mips.ASUBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:        "MUL",
			argLen:      2,
			commutative: true,
			asm:         mips.AMUL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				clobbers: 105553116266496, // HI LO
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:        "MULT",
			argLen:      2,
			commutative: true,
			asm:         mips.AMUL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 35184372088832}, // HI
					{1, 70368744177664}, // LO
				},
			},
		},
		{
			name:        "MULTU",
			argLen:      2,
			commutative: true,
			asm:         mips.AMULU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 35184372088832}, // HI
					{1, 70368744177664}, // LO
				},
			},
		},
		{
			name:   "DIV",
			argLen: 2,
			asm:    mips.ADIV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 35184372088832}, // HI
					{1, 70368744177664}, // LO
				},
			},
		},
		{
			name:   "DIVU",
			argLen: 2,
			asm:    mips.ADIVU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 35184372088832}, // HI
					{1, 70368744177664}, // LO
				},
			},
		},
		{
			name:        "ADDF",
			argLen:      2,
			commutative: true,
			asm:         mips.AADDF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:        "ADDD",
			argLen:      2,
			commutative: true,
			asm:         mips.AADDD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "SUBF",
			argLen: 2,
			asm:    mips.ASUBF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "SUBD",
			argLen: 2,
			asm:    mips.ASUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:        "MULF",
			argLen:      2,
			commutative: true,
			asm:         mips.AMULF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:        "MULD",
			argLen:      2,
			commutative: true,
			asm:         mips.AMULD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "DIVF",
			argLen: 2,
			asm:    mips.ADIVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "DIVD",
			argLen: 2,
			asm:    mips.ADIVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:        "AND",
			argLen:      2,
			commutative: true,
			asm:         mips.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:    "ANDconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     mips.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:        "OR",
			argLen:      2,
			commutative: true,
			asm:         mips.AOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:    "ORconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     mips.AOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:        "XOR",
			argLen:      2,
			commutative: true,
			asm:         mips.AXOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:    "XORconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     mips.AXOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:        "NOR",
			argLen:      2,
			commutative: true,
			asm:         mips.ANOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:    "NORconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     mips.ANOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "NEG",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "NEGF",
			argLen: 1,
			asm:    mips.ANEGF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "NEGD",
			argLen: 1,
			asm:    mips.ANEGD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "SQRTD",
			argLen: 1,
			asm:    mips.ASQRTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "SLL",
			argLen: 2,
			asm:    mips.ASLL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:    "SLLconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     mips.ASLL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "SRL",
			argLen: 2,
			asm:    mips.ASRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:    "SRLconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     mips.ASRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "SRA",
			argLen: 2,
			asm:    mips.ASRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:    "SRAconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     mips.ASRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "CLZ",
			argLen: 1,
			asm:    mips.ACLZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "SGT",
			argLen: 2,
			asm:    mips.ASGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:    "SGTconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     mips.ASGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "SGTzero",
			argLen: 1,
			asm:    mips.ASGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "SGTU",
			argLen: 2,
			asm:    mips.ASGTU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:    "SGTUconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     mips.ASGTU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "SGTUzero",
			argLen: 1,
			asm:    mips.ASGTU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "CMPEQF",
			argLen: 2,
			asm:    mips.ACMPEQF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "CMPEQD",
			argLen: 2,
			asm:    mips.ACMPEQD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "CMPGEF",
			argLen: 2,
			asm:    mips.ACMPGEF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "CMPGED",
			argLen: 2,
			asm:    mips.ACMPGED,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "CMPGTF",
			argLen: 2,
			asm:    mips.ACMPGTF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "CMPGTD",
			argLen: 2,
			asm:    mips.ACMPGTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{1, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:              "MOVWconst",
			auxType:           auxInt32,
			argLen:            0,
			rematerializeable: true,
			asm:               mips.AMOVW,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:              "MOVFconst",
			auxType:           auxFloat32,
			argLen:            0,
			rematerializeable: true,
			asm:               mips.AMOVF,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:              "MOVDconst",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			asm:               mips.AMOVD,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:              "MOVWaddr",
			auxType:           auxSymOff,
			argLen:            1,
			rematerializeable: true,
			symEffect:         SymAddr,
			asm:               mips.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140737555464192}, // SP SB
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:           "MOVBload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:           "MOVBUload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:           "MOVHload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:           "MOVHUload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:           "MOVWload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:           "MOVFload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:           "MOVDload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:           "MOVBstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVHstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVWstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVFstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVF,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 35183835217920},  // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVDstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 35183835217920},  // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVBstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVHstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVWstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    mips.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "MOVBUreg",
			argLen: 1,
			asm:    mips.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    mips.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "MOVHUreg",
			argLen: 1,
			asm:    mips.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    mips.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:         "MOVWnop",
			argLen:       1,
			resultInArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:         "CMOVZ",
			argLen:       3,
			resultInArg0: true,
			asm:          mips.ACMOVZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
					{1, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
					{2, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:         "CMOVZzero",
			argLen:       2,
			resultInArg0: true,
			asm:          mips.ACMOVZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
					{1, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "MOVWF",
			argLen: 1,
			asm:    mips.AMOVWF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "MOVWD",
			argLen: 1,
			asm:    mips.AMOVWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "TRUNCFW",
			argLen: 1,
			asm:    mips.ATRUNCFW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "TRUNCDW",
			argLen: 1,
			asm:    mips.ATRUNCDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "MOVFD",
			argLen: 1,
			asm:    mips.AMOVFD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:   "MOVDF",
			argLen: 1,
			asm:    mips.AMOVDF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
				outputs: []outputInfo{
					{0, 35183835217920}, // F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30
				},
			},
		},
		{
			name:         "CALLstatic",
			auxType:      auxSymOff,
			argLen:       1,
			clobberFlags: true,
			call:         true,
			symEffect:    SymNone,
			reg: regInfo{
				clobbers: 140737421246462, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31 F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30 HI LO
			},
		},
		{
			name:         "CALLclosure",
			auxType:      auxInt64,
			argLen:       3,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4194304},   // R22
					{0, 402653182}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP R31
				},
				clobbers: 140737421246462, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31 F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30 HI LO
			},
		},
		{
			name:         "CALLinter",
			auxType:      auxInt64,
			argLen:       2,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
				clobbers: 140737421246462, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31 F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30 HI LO
			},
		},
		{
			name:           "LoweredAtomicLoad",
			argLen:         2,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:           "LoweredAtomicStore",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:           "LoweredAtomicStorezero",
			argLen:         2,
			faultOnNilArg0: true,
			hasSideEffects: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:            "LoweredAtomicExchange",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:            "LoweredAtomicAdd",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:            "LoweredAtomicAddconst",
			auxType:         auxInt32,
			argLen:          2,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:            "LoweredAtomicCas",
			argLen:          4,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{2, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:           "LoweredAtomicAnd",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			asm:            mips.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:           "LoweredAtomicOr",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			asm:            mips.AOR,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 469762046},       // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
					{0, 140738025226238}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 SP g R31 SB
				},
			},
		},
		{
			name:           "LoweredZero",
			auxType:        auxInt32,
			argLen:         3,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 2},         // R1
					{1, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
				clobbers: 2, // R1
			},
		},
		{
			name:           "LoweredMove",
			auxType:        auxInt32,
			argLen:         4,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4},         // R2
					{1, 2},         // R1
					{2, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
				clobbers: 6, // R1 R2
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 g R31
				},
			},
		},
		{
			name:   "FPFlagTrue",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:   "FPFlagFalse",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4194304}, // R22
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 335544318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R28 R31
				},
			},
		},
		{
			name:         "LoweredWB",
			auxType:      auxSym,
			argLen:       3,
			clobberFlags: true,
			symEffect:    SymNone,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1048576}, // R20
					{1, 2097152}, // R21
				},
				clobbers: 140737219919872, // R31 F0 F2 F4 F6 F8 F10 F12 F14 F16 F18 F20 F22 F24 F26 F28 F30 HI LO
			},
		},

		{
			name:        "ADDV",
			argLen:      2,
			commutative: true,
			asm:         mips.AADDVU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:    "ADDVconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     mips.AADDVU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 268435454}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "SUBV",
			argLen: 2,
			asm:    mips.ASUBVU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:    "SUBVconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     mips.ASUBVU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:        "MULV",
			argLen:      2,
			commutative: true,
			asm:         mips.AMULV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 1152921504606846976}, // HI
					{1, 2305843009213693952}, // LO
				},
			},
		},
		{
			name:        "MULVU",
			argLen:      2,
			commutative: true,
			asm:         mips.AMULVU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 1152921504606846976}, // HI
					{1, 2305843009213693952}, // LO
				},
			},
		},
		{
			name:   "DIVV",
			argLen: 2,
			asm:    mips.ADIVV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 1152921504606846976}, // HI
					{1, 2305843009213693952}, // LO
				},
			},
		},
		{
			name:   "DIVVU",
			argLen: 2,
			asm:    mips.ADIVVU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 1152921504606846976}, // HI
					{1, 2305843009213693952}, // LO
				},
			},
		},
		{
			name:        "ADDF",
			argLen:      2,
			commutative: true,
			asm:         mips.AADDF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:        "ADDD",
			argLen:      2,
			commutative: true,
			asm:         mips.AADDD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "SUBF",
			argLen: 2,
			asm:    mips.ASUBF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "SUBD",
			argLen: 2,
			asm:    mips.ASUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:        "MULF",
			argLen:      2,
			commutative: true,
			asm:         mips.AMULF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:        "MULD",
			argLen:      2,
			commutative: true,
			asm:         mips.AMULD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "DIVF",
			argLen: 2,
			asm:    mips.ADIVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "DIVD",
			argLen: 2,
			asm:    mips.ADIVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:        "AND",
			argLen:      2,
			commutative: true,
			asm:         mips.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:    "ANDconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     mips.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:        "OR",
			argLen:      2,
			commutative: true,
			asm:         mips.AOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:    "ORconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     mips.AOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:        "XOR",
			argLen:      2,
			commutative: true,
			asm:         mips.AXOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:    "XORconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     mips.AXOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:        "NOR",
			argLen:      2,
			commutative: true,
			asm:         mips.ANOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:    "NORconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     mips.ANOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "NEGV",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "NEGF",
			argLen: 1,
			asm:    mips.ANEGF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "NEGD",
			argLen: 1,
			asm:    mips.ANEGD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "SQRTD",
			argLen: 1,
			asm:    mips.ASQRTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "SLLV",
			argLen: 2,
			asm:    mips.ASLLV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:    "SLLVconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     mips.ASLLV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "SRLV",
			argLen: 2,
			asm:    mips.ASRLV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:    "SRLVconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     mips.ASRLV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "SRAV",
			argLen: 2,
			asm:    mips.ASRAV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:    "SRAVconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     mips.ASRAV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "SGT",
			argLen: 2,
			asm:    mips.ASGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:    "SGTconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     mips.ASGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "SGTU",
			argLen: 2,
			asm:    mips.ASGTU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{1, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:    "SGTUconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     mips.ASGTU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "CMPEQF",
			argLen: 2,
			asm:    mips.ACMPEQF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "CMPEQD",
			argLen: 2,
			asm:    mips.ACMPEQD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "CMPGEF",
			argLen: 2,
			asm:    mips.ACMPGEF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "CMPGED",
			argLen: 2,
			asm:    mips.ACMPGED,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "CMPGTF",
			argLen: 2,
			asm:    mips.ACMPGTF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "CMPGTD",
			argLen: 2,
			asm:    mips.ACMPGTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:              "MOVVconst",
			auxType:           auxInt64,
			argLen:            0,
			rematerializeable: true,
			asm:               mips.AMOVV,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:              "MOVFconst",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			asm:               mips.AMOVF,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:              "MOVDconst",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			asm:               mips.AMOVD,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:              "MOVVaddr",
			auxType:           auxSymOff,
			argLen:            1,
			rematerializeable: true,
			symEffect:         SymAddr,
			asm:               mips.AMOVV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018460942336}, // SP SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:           "MOVBload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:           "MOVBUload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:           "MOVHload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:           "MOVHUload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:           "MOVWload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:           "MOVWUload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:           "MOVVload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:           "MOVFload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:           "MOVDload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            mips.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:           "MOVBstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVHstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVWstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVVstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVV,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVFstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:           "MOVDstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
					{1, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:           "MOVBstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVHstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVWstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:           "MOVVstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            mips.AMOVV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    mips.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "MOVBUreg",
			argLen: 1,
			asm:    mips.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    mips.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "MOVHUreg",
			argLen: 1,
			asm:    mips.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    mips.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "MOVWUreg",
			argLen: 1,
			asm:    mips.AMOVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "MOVVreg",
			argLen: 1,
			asm:    mips.AMOVV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:         "MOVVnop",
			argLen:       1,
			resultInArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "MOVWF",
			argLen: 1,
			asm:    mips.AMOVWF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "MOVWD",
			argLen: 1,
			asm:    mips.AMOVWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "MOVVF",
			argLen: 1,
			asm:    mips.AMOVVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "MOVVD",
			argLen: 1,
			asm:    mips.AMOVVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "TRUNCFW",
			argLen: 1,
			asm:    mips.ATRUNCFW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "TRUNCDW",
			argLen: 1,
			asm:    mips.ATRUNCDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "TRUNCFV",
			argLen: 1,
			asm:    mips.ATRUNCFV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "TRUNCDV",
			argLen: 1,
			asm:    mips.ATRUNCDV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "MOVFD",
			argLen: 1,
			asm:    mips.AMOVFD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:   "MOVDF",
			argLen: 1,
			asm:    mips.AMOVDF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
				outputs: []outputInfo{
					{0, 1152921504338411520}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31
				},
			},
		},
		{
			name:         "CALLstatic",
			auxType:      auxSymOff,
			argLen:       1,
			clobberFlags: true,
			call:         true,
			symEffect:    SymNone,
			reg: regInfo{
				clobbers: 4611686018393833470, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31 HI LO
			},
		},
		{
			name:         "CALLclosure",
			auxType:      auxInt64,
			argLen:       3,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4194304},   // R22
					{0, 201326590}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP R31
				},
				clobbers: 4611686018393833470, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31 HI LO
			},
		},
		{
			name:         "CALLinter",
			auxType:      auxInt64,
			argLen:       2,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
				clobbers: 4611686018393833470, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31 HI LO
			},
		},
		{
			name:           "DUFFZERO",
			auxType:        auxInt64,
			argLen:         2,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
				clobbers: 134217730, // R1 R31
			},
		},
		{
			name:           "LoweredZero",
			auxType:        auxInt64,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 2},         // R1
					{1, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
				clobbers: 2, // R1
			},
		},
		{
			name:           "LoweredMove",
			auxType:        auxInt64,
			argLen:         4,
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4},         // R2
					{1, 2},         // R1
					{2, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
				clobbers: 6, // R1 R2
			},
		},
		{
			name:           "LoweredAtomicLoad32",
			argLen:         2,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:           "LoweredAtomicLoad64",
			argLen:         2,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:           "LoweredAtomicStore32",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:           "LoweredAtomicStore64",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:           "LoweredAtomicStorezero32",
			argLen:         2,
			faultOnNilArg0: true,
			hasSideEffects: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:           "LoweredAtomicStorezero64",
			argLen:         2,
			faultOnNilArg0: true,
			hasSideEffects: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
			},
		},
		{
			name:            "LoweredAtomicExchange32",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:            "LoweredAtomicExchange64",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:            "LoweredAtomicAdd32",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:            "LoweredAtomicAdd64",
			argLen:          3,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:            "LoweredAtomicAddconst32",
			auxType:         auxInt32,
			argLen:          2,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:            "LoweredAtomicAddconst64",
			auxType:         auxInt64,
			argLen:          2,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:            "LoweredAtomicCas32",
			argLen:          4,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{2, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:            "LoweredAtomicCas64",
			argLen:          4,
			resultNotInArgs: true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{2, 234881022},           // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
					{0, 4611686018695823358}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 SP g R31 SB
				},
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 g R31
				},
			},
		},
		{
			name:   "FPFlagTrue",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:   "FPFlagFalse",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4194304}, // R22
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 167772158}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31
				},
			},
		},
		{
			name:         "LoweredWB",
			auxType:      auxSym,
			argLen:       3,
			clobberFlags: true,
			symEffect:    SymNone,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1048576}, // R20
					{1, 2097152}, // R21
				},
				clobbers: 4611686018293170176, // R31 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31 HI LO
			},
		},

		{
			name:        "ADD",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "ADDconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:        "FADD",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AFADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:        "FADDS",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AFADDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "SUB",
			argLen: 2,
			asm:    ppc64.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "FSUB",
			argLen: 2,
			asm:    ppc64.AFSUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FSUBS",
			argLen: 2,
			asm:    ppc64.AFSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:        "MULLD",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AMULLD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:        "MULLW",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AMULLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:        "MULHD",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AMULHD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:        "MULHW",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AMULHW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:        "MULHDU",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AMULHDU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:        "MULHWU",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AMULHWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:        "FMUL",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AFMUL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:        "FMULS",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AFMULS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FMADD",
			argLen: 3,
			asm:    ppc64.AFMADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{2, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FMADDS",
			argLen: 3,
			asm:    ppc64.AFMADDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{2, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FMSUB",
			argLen: 3,
			asm:    ppc64.AFMSUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{2, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FMSUBS",
			argLen: 3,
			asm:    ppc64.AFMSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{2, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "SRAD",
			argLen: 2,
			asm:    ppc64.ASRAD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "SRAW",
			argLen: 2,
			asm:    ppc64.ASRAW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "SRD",
			argLen: 2,
			asm:    ppc64.ASRD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "SRW",
			argLen: 2,
			asm:    ppc64.ASRW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "SLD",
			argLen: 2,
			asm:    ppc64.ASLD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "SLW",
			argLen: 2,
			asm:    ppc64.ASLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "ROTL",
			argLen: 2,
			asm:    ppc64.AROTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "ROTLW",
			argLen: 2,
			asm:    ppc64.AROTLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "ADDconstForCarry",
			auxType: auxInt16,
			argLen:  1,
			asm:     ppc64.AADDC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				clobbers: 2147483648, // R31
			},
		},
		{
			name:   "MaskIfNotCarry",
			argLen: 1,
			asm:    ppc64.AADDME,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "SRADconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.ASRAD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "SRAWconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.ASRAW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "SRDconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.ASRD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "SRWconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.ASRW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "SLDconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.ASLD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "SLWconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.ASLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "ROTLconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.AROTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "ROTLWconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.AROTLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:         "CNTLZD",
			argLen:       1,
			clobberFlags: true,
			asm:          ppc64.ACNTLZD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:         "CNTLZW",
			argLen:       1,
			clobberFlags: true,
			asm:          ppc64.ACNTLZW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "POPCNTD",
			argLen: 1,
			asm:    ppc64.APOPCNTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "POPCNTW",
			argLen: 1,
			asm:    ppc64.APOPCNTW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "POPCNTB",
			argLen: 1,
			asm:    ppc64.APOPCNTB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "FDIV",
			argLen: 2,
			asm:    ppc64.AFDIV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FDIVS",
			argLen: 2,
			asm:    ppc64.AFDIVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "DIVD",
			argLen: 2,
			asm:    ppc64.ADIVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "DIVW",
			argLen: 2,
			asm:    ppc64.ADIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "DIVDU",
			argLen: 2,
			asm:    ppc64.ADIVDU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "DIVWU",
			argLen: 2,
			asm:    ppc64.ADIVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "FCTIDZ",
			argLen: 1,
			asm:    ppc64.AFCTIDZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FCTIWZ",
			argLen: 1,
			asm:    ppc64.AFCTIWZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FCFID",
			argLen: 1,
			asm:    ppc64.AFCFID,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FCFIDS",
			argLen: 1,
			asm:    ppc64.AFCFIDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FRSP",
			argLen: 1,
			asm:    ppc64.AFRSP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "MFVSRD",
			argLen: 1,
			asm:    ppc64.AMFVSRD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "MTVSRD",
			argLen: 1,
			asm:    ppc64.AMTVSRD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:        "AND",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "ANDN",
			argLen: 2,
			asm:    ppc64.AANDN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:        "OR",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "ORN",
			argLen: 2,
			asm:    ppc64.AORN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:        "NOR",
			argLen:      2,
			commutative: true,
			asm:         ppc64.ANOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:        "XOR",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AXOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:        "EQV",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AEQV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "NEG",
			argLen: 1,
			asm:    ppc64.ANEG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "FNEG",
			argLen: 1,
			asm:    ppc64.AFNEG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FSQRT",
			argLen: 1,
			asm:    ppc64.AFSQRT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FSQRTS",
			argLen: 1,
			asm:    ppc64.AFSQRTS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FFLOOR",
			argLen: 1,
			asm:    ppc64.AFRIM,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FCEIL",
			argLen: 1,
			asm:    ppc64.AFRIP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FTRUNC",
			argLen: 1,
			asm:    ppc64.AFRIZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FROUND",
			argLen: 1,
			asm:    ppc64.AFRIN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FABS",
			argLen: 1,
			asm:    ppc64.AFABS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FNABS",
			argLen: 1,
			asm:    ppc64.AFNABS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FCPSGN",
			argLen: 2,
			asm:    ppc64.AFCPSGN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:    "ORconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.AOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "XORconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.AXOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:         "ANDconst",
			auxType:      auxInt64,
			argLen:       1,
			clobberFlags: true,
			asm:          ppc64.AANDCC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "ANDCCconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.AANDCC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    ppc64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "MOVBZreg",
			argLen: 1,
			asm:    ppc64.AMOVBZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    ppc64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "MOVHZreg",
			argLen: 1,
			asm:    ppc64.AMOVHZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    ppc64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "MOVWZreg",
			argLen: 1,
			asm:    ppc64.AMOVWZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVBZload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            ppc64.AMOVBZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVHload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            ppc64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVHZload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            ppc64.AMOVHZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVWload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            ppc64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVWZload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            ppc64.AMOVWZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVDload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            ppc64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVDBRload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            ppc64.AMOVDBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVWBRload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            ppc64.AMOVWBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVHBRload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            ppc64.AMOVHBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVDBRstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AMOVDBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVWBRstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AMOVWBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVHBRstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AMOVHBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "FMOVDload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            ppc64.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:           "FMOVSload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            ppc64.AFMOVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:           "MOVBstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVHstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVWstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVDstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "FMOVDstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{0, 1073733630},         // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "FMOVSstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AFMOVS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{0, 1073733630},         // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVBstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVHstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVWstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "MOVDstorezero",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            ppc64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:              "MOVDaddr",
			auxType:           auxSymOff,
			argLen:            1,
			rematerializeable: true,
			symEffect:         SymAddr,
			asm:               ppc64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:              "MOVDconst",
			auxType:           auxInt64,
			argLen:            0,
			rematerializeable: true,
			asm:               ppc64.AMOVD,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:              "FMOVDconst",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			asm:               ppc64.AFMOVD,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:              "FMOVSconst",
			auxType:           auxFloat32,
			argLen:            0,
			rematerializeable: true,
			asm:               ppc64.AFMOVS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "FCMPU",
			argLen: 2,
			asm:    ppc64.AFCMPU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
					{1, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:   "CMP",
			argLen: 2,
			asm:    ppc64.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "CMPU",
			argLen: 2,
			asm:    ppc64.ACMPU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "CMPW",
			argLen: 2,
			asm:    ppc64.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "CMPWU",
			argLen: 2,
			asm:    ppc64.ACMPWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "CMPconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "CMPUconst",
			auxType: auxInt64,
			argLen:  1,
			asm:     ppc64.ACMPU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "CMPWconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     ppc64.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:    "CMPWUconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     ppc64.ACMPWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "Equal",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "NotEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "LessThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "FLessThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "LessEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "FLessEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "GreaterThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "FGreaterThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "GreaterEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:   "FGreaterEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 2048}, // R11
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			clobberFlags:   true,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				clobbers: 2147483648, // R31
			},
		},
		{
			name:         "LoweredRound32F",
			argLen:       1,
			resultInArg0: true,
			zeroWidth:    true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:         "LoweredRound64F",
			argLen:       1,
			resultInArg0: true,
			zeroWidth:    true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
				outputs: []outputInfo{
					{0, 576460743713488896}, // F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
				},
			},
		},
		{
			name:         "CALLstatic",
			auxType:      auxSymOff,
			argLen:       1,
			clobberFlags: true,
			call:         true,
			symEffect:    SymNone,
			reg: regInfo{
				clobbers: 576460745860964344, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29 g F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
		{
			name:         "CALLclosure",
			auxType:      auxInt64,
			argLen:       3,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4096}, // R12
					{1, 2048}, // R11
				},
				clobbers: 576460745860964344, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29 g F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
		{
			name:         "CALLinter",
			auxType:      auxInt64,
			argLen:       2,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4096}, // R12
				},
				clobbers: 576460745860964344, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29 g F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
		{
			name:           "LoweredZero",
			auxType:        auxInt64,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 8}, // R3
				},
				clobbers: 8, // R3
			},
		},
		{
			name:           "LoweredMove",
			auxType:        auxInt64,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 8},  // R3
					{1, 16}, // R4
				},
				clobbers: 1944, // R3 R4 R7 R8 R9 R10
			},
		},
		{
			name:           "LoweredAtomicStore32",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "LoweredAtomicStore64",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "LoweredAtomicLoad32",
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "LoweredAtomicLoad64",
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "LoweredAtomicLoadPtr",
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:            "LoweredAtomicAdd32",
			argLen:          3,
			resultNotInArgs: true,
			clobberFlags:    true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:            "LoweredAtomicAdd64",
			argLen:          3,
			resultNotInArgs: true,
			clobberFlags:    true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:            "LoweredAtomicExchange32",
			argLen:          3,
			resultNotInArgs: true,
			clobberFlags:    true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:            "LoweredAtomicExchange64",
			argLen:          3,
			resultNotInArgs: true,
			clobberFlags:    true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:            "LoweredAtomicCas64",
			argLen:          4,
			resultNotInArgs: true,
			clobberFlags:    true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{2, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:            "LoweredAtomicCas32",
			argLen:          4,
			resultNotInArgs: true,
			clobberFlags:    true,
			faultOnNilArg0:  true,
			hasSideEffects:  true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{2, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
				outputs: []outputInfo{
					{0, 1073733624}, // R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "LoweredAtomicAnd8",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			asm:            ppc64.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:           "LoweredAtomicOr8",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			asm:            ppc64.AOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
					{1, 1073733630}, // SP SB R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29
				},
			},
		},
		{
			name:         "LoweredWB",
			auxType:      auxSym,
			argLen:       3,
			clobberFlags: true,
			symEffect:    SymNone,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1048576}, // R20
					{1, 2097152}, // R21
				},
				clobbers: 576460746931503104, // R16 R17 R18 R19 R22 R23 R24 R25 R26 R27 R28 R29 R31 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26
			},
		},
		{
			name:   "InvertFlags",
			argLen: 1,
			reg:    regInfo{},
		},
		{
			name:   "FlagEQ",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagLT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagGT",
			argLen: 0,
			reg:    regInfo{},
		},

		{
			name:         "FADDS",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AFADDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FADD",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AFADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FSUBS",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AFSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FSUB",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AFSUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FMULS",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			asm:          s390x.AFMULS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FMUL",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			asm:          s390x.AFMUL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FDIVS",
			argLen:       2,
			resultInArg0: true,
			asm:          s390x.AFDIVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FDIV",
			argLen:       2,
			resultInArg0: true,
			asm:          s390x.AFDIV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FNEGS",
			argLen:       1,
			clobberFlags: true,
			asm:          s390x.AFNEGS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FNEG",
			argLen:       1,
			clobberFlags: true,
			asm:          s390x.AFNEG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FMADDS",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AFMADDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FMADD",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AFMADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FMSUBS",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AFMSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "FMSUB",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AFMSUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "LPDFR",
			argLen: 1,
			asm:    s390x.ALPDFR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "LNDFR",
			argLen: 1,
			asm:    s390x.ALNDFR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "CPSDR",
			argLen: 2,
			asm:    s390x.ACPSDR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:    "FIDBR",
			auxType: auxInt8,
			argLen:  1,
			asm:     s390x.AFIDBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:           "FMOVSload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AFMOVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:           "FMOVDload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:              "FMOVSconst",
			auxType:           auxFloat32,
			argLen:            0,
			rematerializeable: true,
			asm:               s390x.AFMOVS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:              "FMOVDconst",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			asm:               s390x.AFMOVD,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:      "FMOVSloadidx",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       s390x.AFMOVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:      "FMOVDloadidx",
			auxType:   auxSymOff,
			argLen:    3,
			symEffect: SymRead,
			asm:       s390x.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:           "FMOVSstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AFMOVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:           "FMOVDstore",
			auxType:        auxSymOff,
			argLen:         3,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:      "FMOVSstoreidx",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       s390x.AFMOVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:      "FMOVDstoreidx",
			auxType:   auxSymOff,
			argLen:    4,
			symEffect: SymWrite,
			asm:       s390x.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{2, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "ADD",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          s390x.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "ADDW",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          s390x.AADDW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "ADDconst",
			auxType:      auxInt32,
			argLen:       1,
			clobberFlags: true,
			asm:          s390x.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "ADDWconst",
			auxType:      auxInt32,
			argLen:       1,
			clobberFlags: true,
			asm:          s390x.AADDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "ADDload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "ADDWload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.AADDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "SUB",
			argLen:       2,
			clobberFlags: true,
			asm:          s390x.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "SUBW",
			argLen:       2,
			clobberFlags: true,
			asm:          s390x.ASUBW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "SUBconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "SUBWconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.ASUBW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "SUBload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "SUBWload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.ASUBW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MULLD",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AMULLD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MULLW",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AMULLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MULLDconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AMULLD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MULLWconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AMULLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MULLDload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.AMULLD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MULLWload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.AMULLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MULHD",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AMULHD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2048, // R11
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "MULHDU",
			argLen:       2,
			commutative:  true,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AMULHDU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2048, // R11
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "DIVD",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.ADIVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2048, // R11
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "DIVW",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.ADIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2048, // R11
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "DIVDU",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.ADIVDU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2048, // R11
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "DIVWU",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.ADIVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2048, // R11
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "MODD",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AMODD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2048, // R11
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "MODW",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AMODW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2048, // R11
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "MODDU",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AMODDU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2048, // R11
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "MODWU",
			argLen:       2,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AMODWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
					{1, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
				clobbers: 2048, // R11
				outputs: []outputInfo{
					{0, 21503}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14
				},
			},
		},
		{
			name:         "AND",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          s390x.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "ANDW",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          s390x.AANDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "ANDconst",
			auxType:      auxInt64,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "ANDWconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AANDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "ANDload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "ANDWload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.AANDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "OR",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          s390x.AOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "ORW",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          s390x.AORW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "ORconst",
			auxType:      auxInt64,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "ORWconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AORW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "ORload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.AOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "ORWload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.AORW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "XOR",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          s390x.AXOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "XORW",
			argLen:       2,
			commutative:  true,
			clobberFlags: true,
			asm:          s390x.AXORW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "XORconst",
			auxType:      auxInt64,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AXOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "XORWconst",
			auxType:      auxInt32,
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			asm:          s390x.AXORW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "XORload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.AXOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "XORWload",
			auxType:        auxSymOff,
			argLen:         3,
			resultInArg0:   true,
			clobberFlags:   true,
			faultOnNilArg1: true,
			symEffect:      SymRead,
			asm:            s390x.AXORW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "CMP",
			argLen: 2,
			asm:    s390x.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:   "CMPW",
			argLen: 2,
			asm:    s390x.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:   "CMPU",
			argLen: 2,
			asm:    s390x.ACMPU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:   "CMPWU",
			argLen: 2,
			asm:    s390x.ACMPWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:    "CMPconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     s390x.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:    "CMPWconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     s390x.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:    "CMPUconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     s390x.ACMPU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:    "CMPWUconst",
			auxType: auxInt32,
			argLen:  1,
			asm:     s390x.ACMPWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:   "FCMPS",
			argLen: 2,
			asm:    s390x.ACEBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "FCMP",
			argLen: 2,
			asm:    s390x.AFCMPU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "SLD",
			argLen: 2,
			asm:    s390x.ASLD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "SLW",
			argLen: 2,
			asm:    s390x.ASLW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:    "SLDconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     s390x.ASLD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:    "SLWconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     s390x.ASLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "SRD",
			argLen: 2,
			asm:    s390x.ASRD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "SRW",
			argLen: 2,
			asm:    s390x.ASRW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:    "SRDconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     s390x.ASRD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:    "SRWconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     s390x.ASRW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "SRAD",
			argLen:       2,
			clobberFlags: true,
			asm:          s390x.ASRAD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "SRAW",
			argLen:       2,
			clobberFlags: true,
			asm:          s390x.ASRAW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "SRADconst",
			auxType:      auxInt8,
			argLen:       1,
			clobberFlags: true,
			asm:          s390x.ASRAD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "SRAWconst",
			auxType:      auxInt8,
			argLen:       1,
			clobberFlags: true,
			asm:          s390x.ASRAW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:    "RLLGconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     s390x.ARLLG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:    "RLLconst",
			auxType: auxInt8,
			argLen:  1,
			asm:     s390x.ARLL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "NEG",
			argLen:       1,
			clobberFlags: true,
			asm:          s390x.ANEG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "NEGW",
			argLen:       1,
			clobberFlags: true,
			asm:          s390x.ANEGW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "NOT",
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "NOTW",
			argLen:       1,
			resultInArg0: true,
			clobberFlags: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "FSQRT",
			argLen: 1,
			asm:    s390x.AFSQRT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "MOVDEQ",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AMOVDEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVDNE",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AMOVDNE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVDLT",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AMOVDLT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVDLE",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AMOVDLE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVDGT",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AMOVDGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVDGE",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AMOVDGE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVDGTnoinv",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AMOVDGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVDGEnoinv",
			argLen:       3,
			resultInArg0: true,
			asm:          s390x.AMOVDGE,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
					{1, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    s390x.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "MOVBZreg",
			argLen: 1,
			asm:    s390x.AMOVBZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    s390x.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "MOVHZreg",
			argLen: 1,
			asm:    s390x.AMOVHZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    s390x.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "MOVWZreg",
			argLen: 1,
			asm:    s390x.AMOVWZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "MOVDreg",
			argLen: 1,
			asm:    s390x.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVDnop",
			argLen:       1,
			resultInArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:              "MOVDconst",
			auxType:           auxInt64,
			argLen:            0,
			rematerializeable: true,
			asm:               s390x.AMOVD,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "LDGR",
			argLen: 1,
			asm:    s390x.ALDGR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "LGDR",
			argLen: 1,
			asm:    s390x.ALGDR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "CFDBRA",
			argLen: 1,
			asm:    s390x.ACFDBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "CGDBRA",
			argLen: 1,
			asm:    s390x.ACGDBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "CFEBRA",
			argLen: 1,
			asm:    s390x.ACFEBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "CGEBRA",
			argLen: 1,
			asm:    s390x.ACGEBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "CEFBRA",
			argLen: 1,
			asm:    s390x.ACEFBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "CDFBRA",
			argLen: 1,
			asm:    s390x.ACDFBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "CEGBRA",
			argLen: 1,
			asm:    s390x.ACEGBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "CDGBRA",
			argLen: 1,
			asm:    s390x.ACDGBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "LEDBR",
			argLen: 1,
			asm:    s390x.ALEDBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "LDEBR",
			argLen: 1,
			asm:    s390x.ALDEBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:              "MOVDaddr",
			auxType:           auxSymOff,
			argLen:            1,
			rematerializeable: true,
			symEffect:         SymRead,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295000064}, // SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:      "MOVDaddridx",
			auxType:   auxSymOff,
			argLen:    2,
			symEffect: SymRead,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295000064}, // SP SB
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVBZload",
			auxType:        auxSymOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVBZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVBload",
			auxType:        auxSymOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVHZload",
			auxType:        auxSymOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVHZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVHload",
			auxType:        auxSymOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVWZload",
			auxType:        auxSymOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVWZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVWload",
			auxType:        auxSymOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVDload",
			auxType:        auxSymOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "MOVWBR",
			argLen: 1,
			asm:    s390x.AMOVWBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "MOVDBR",
			argLen: 1,
			asm:    s390x.AMOVDBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVHBRload",
			auxType:        auxSymOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVHBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVWBRload",
			auxType:        auxSymOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVWBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVDBRload",
			auxType:        auxSymOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVDBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVBstore",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
					{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "MOVHstore",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
					{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "MOVWstore",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
					{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "MOVDstore",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
					{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "MOVHBRstore",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVHBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "MOVWBRstore",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVWBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "MOVDBRstore",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVDBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "MVC",
			auxType:        auxSymValAndOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			symEffect:      SymNone,
			asm:            s390x.AMVC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:         "MOVBZloadidx",
			auxType:      auxSymOff,
			argLen:       3,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymRead,
			asm:          s390x.AMOVBZ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVBloadidx",
			auxType:      auxSymOff,
			argLen:       3,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymRead,
			asm:          s390x.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVHZloadidx",
			auxType:      auxSymOff,
			argLen:       3,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymRead,
			asm:          s390x.AMOVHZ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVHloadidx",
			auxType:      auxSymOff,
			argLen:       3,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymRead,
			asm:          s390x.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVWZloadidx",
			auxType:      auxSymOff,
			argLen:       3,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymRead,
			asm:          s390x.AMOVWZ,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVWloadidx",
			auxType:      auxSymOff,
			argLen:       3,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymRead,
			asm:          s390x.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVDloadidx",
			auxType:      auxSymOff,
			argLen:       3,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymRead,
			asm:          s390x.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVHBRloadidx",
			auxType:      auxSymOff,
			argLen:       3,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymRead,
			asm:          s390x.AMOVHBR,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVWBRloadidx",
			auxType:      auxSymOff,
			argLen:       3,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymRead,
			asm:          s390x.AMOVWBR,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVDBRloadidx",
			auxType:      auxSymOff,
			argLen:       3,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymRead,
			asm:          s390x.AMOVDBR,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 56318},      // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "MOVBstoreidx",
			auxType:      auxSymOff,
			argLen:       4,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymWrite,
			asm:          s390x.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:         "MOVHstoreidx",
			auxType:      auxSymOff,
			argLen:       4,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymWrite,
			asm:          s390x.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:         "MOVWstoreidx",
			auxType:      auxSymOff,
			argLen:       4,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymWrite,
			asm:          s390x.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:         "MOVDstoreidx",
			auxType:      auxSymOff,
			argLen:       4,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymWrite,
			asm:          s390x.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:         "MOVHBRstoreidx",
			auxType:      auxSymOff,
			argLen:       4,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymWrite,
			asm:          s390x.AMOVHBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:         "MOVWBRstoreidx",
			auxType:      auxSymOff,
			argLen:       4,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymWrite,
			asm:          s390x.AMOVWBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:         "MOVDBRstoreidx",
			auxType:      auxSymOff,
			argLen:       4,
			commutative:  true,
			clobberFlags: true,
			symEffect:    SymWrite,
			asm:          s390x.AMOVDBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "MOVBstoreconst",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
			},
		},
		{
			name:           "MOVHstoreconst",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
			},
		},
		{
			name:           "MOVWstoreconst",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
			},
		},
		{
			name:           "MOVDstoreconst",
			auxType:        auxSymValAndOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
			},
		},
		{
			name:           "CLEAR",
			auxType:        auxSymValAndOff,
			argLen:         2,
			clobberFlags:   true,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.ACLEAR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:         "CALLstatic",
			auxType:      auxSymOff,
			argLen:       1,
			clobberFlags: true,
			call:         true,
			symEffect:    SymNone,
			reg: regInfo{
				clobbers: 4294933503, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 g R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
		{
			name:         "CALLclosure",
			auxType:      auxInt64,
			argLen:       3,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4096},  // R12
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				clobbers: 4294933503, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 g R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
		{
			name:         "CALLinter",
			auxType:      auxInt64,
			argLen:       2,
			clobberFlags: true,
			call:         true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23550}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				clobbers: 4294933503, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 g R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
		{
			name:   "InvertFlags",
			argLen: 1,
			reg:    regInfo{},
		},
		{
			name:   "LoweredGetG",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4096}, // R12
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			clobberFlags:   true,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:         "LoweredRound32F",
			argLen:       1,
			resultInArg0: true,
			zeroWidth:    true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "LoweredRound64F",
			argLen:       1,
			resultInArg0: true,
			zeroWidth:    true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:         "LoweredWB",
			auxType:      auxSym,
			argLen:       3,
			clobberFlags: true,
			symEffect:    SymNone,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4}, // R2
					{1, 8}, // R3
				},
				clobbers: 4294918144, // R14 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
			},
		},
		{
			name:   "FlagEQ",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagLT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:   "FlagGT",
			argLen: 0,
			reg:    regInfo{},
		},
		{
			name:           "MOVWZatomicload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVWZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVDatomicload",
			auxType:        auxSymOff,
			argLen:         2,
			faultOnNilArg0: true,
			symEffect:      SymRead,
			asm:            s390x.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "MOVWatomicstore",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
					{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "MOVDatomicstore",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymWrite,
			asm:            s390x.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
					{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "LAA",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            s390x.ALAA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
					{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "LAAG",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            s390x.ALAAG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295023614}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP SB
					{1, 56319},      // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:   "AddTupleFirst32",
			argLen: 2,
			reg:    regInfo{},
		},
		{
			name:   "AddTupleFirst64",
			argLen: 2,
			reg:    regInfo{},
		},
		{
			name:           "LoweredAtomicCas32",
			auxType:        auxSymOff,
			argLen:         4,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            s390x.ACS,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 1},     // R0
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				clobbers: 1, // R0
				outputs: []outputInfo{
					{1, 0},
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "LoweredAtomicCas64",
			auxType:        auxSymOff,
			argLen:         4,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            s390x.ACSG,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 1},     // R0
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				clobbers: 1, // R0
				outputs: []outputInfo{
					{1, 0},
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
			},
		},
		{
			name:           "LoweredAtomicExchange32",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            s390x.ACS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 1}, // R0
				},
			},
		},
		{
			name:           "LoweredAtomicExchange64",
			auxType:        auxSymOff,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			hasSideEffects: true,
			symEffect:      SymRdWr,
			asm:            s390x.ACSG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
					{1, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 1}, // R0
				},
			},
		},
		{
			name:         "FLOGR",
			argLen:       1,
			clobberFlags: true,
			asm:          s390x.AFLOGR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14
				},
				clobbers: 2, // R1
				outputs: []outputInfo{
					{0, 1}, // R0
				},
			},
		},
		{
			name:           "STMG2",
			auxType:        auxSymOff,
			argLen:         4,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.ASTMG,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // R1
					{2, 4},     // R2
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "STMG3",
			auxType:        auxSymOff,
			argLen:         5,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.ASTMG,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // R1
					{2, 4},     // R2
					{3, 8},     // R3
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "STMG4",
			auxType:        auxSymOff,
			argLen:         6,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.ASTMG,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // R1
					{2, 4},     // R2
					{3, 8},     // R3
					{4, 16},    // R4
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "STM2",
			auxType:        auxSymOff,
			argLen:         4,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.ASTMY,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // R1
					{2, 4},     // R2
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "STM3",
			auxType:        auxSymOff,
			argLen:         5,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.ASTMY,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // R1
					{2, 4},     // R2
					{3, 8},     // R3
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "STM4",
			auxType:        auxSymOff,
			argLen:         6,
			faultOnNilArg0: true,
			symEffect:      SymWrite,
			asm:            s390x.ASTMY,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 2},     // R1
					{2, 4},     // R2
					{3, 8},     // R3
					{4, 16},    // R4
					{0, 56318}, // R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
			},
		},
		{
			name:           "LoweredMove",
			auxType:        auxInt64,
			argLen:         4,
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 2},     // R1
					{1, 4},     // R2
					{2, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				clobbers: 6, // R1 R2
			},
		},
		{
			name:           "LoweredZero",
			auxType:        auxInt64,
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 2},     // R1
					{1, 56319}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14 SP
				},
				clobbers: 2, // R1
			},
		},

		{
			name:      "LoweredStaticCall",
			auxType:   auxSymOff,
			argLen:    1,
			call:      true,
			symEffect: SymNone,
			reg: regInfo{
				clobbers: 12884901887, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 g
			},
		},
		{
			name:    "LoweredClosureCall",
			auxType: auxInt64,
			argLen:  3,
			call:    true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
				clobbers: 12884901887, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 g
			},
		},
		{
			name:    "LoweredInterCall",
			auxType: auxInt64,
			argLen:  2,
			call:    true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
				clobbers: 12884901887, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 g
			},
		},
		{
			name:              "LoweredAddr",
			auxType:           auxSymOff,
			argLen:            1,
			rematerializeable: true,
			symEffect:         SymAddr,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "LoweredMove",
			auxType: auxInt64,
			argLen:  3,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "LoweredZero",
			auxType: auxInt64,
			argLen:  2,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "LoweredGetClosurePtr",
			argLen: 0,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:      "LoweredWB",
			auxType:   auxSym,
			argLen:    3,
			symEffect: SymNone,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
					{1, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "LoweredRound32F",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "LoweredConvert",
			argLen: 2,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "Select",
			argLen: 3,
			asm:    wasm.ASelect,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{2, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "I64Load8U",
			auxType: auxInt64,
			argLen:  2,
			asm:     wasm.AI64Load8U,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "I64Load8S",
			auxType: auxInt64,
			argLen:  2,
			asm:     wasm.AI64Load8S,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "I64Load16U",
			auxType: auxInt64,
			argLen:  2,
			asm:     wasm.AI64Load16U,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "I64Load16S",
			auxType: auxInt64,
			argLen:  2,
			asm:     wasm.AI64Load16S,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "I64Load32U",
			auxType: auxInt64,
			argLen:  2,
			asm:     wasm.AI64Load32U,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "I64Load32S",
			auxType: auxInt64,
			argLen:  2,
			asm:     wasm.AI64Load32S,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "I64Load",
			auxType: auxInt64,
			argLen:  2,
			asm:     wasm.AI64Load,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "I64Store8",
			auxType: auxInt64,
			argLen:  3,
			asm:     wasm.AI64Store8,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4295032831},  // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
			},
		},
		{
			name:    "I64Store16",
			auxType: auxInt64,
			argLen:  3,
			asm:     wasm.AI64Store16,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4295032831},  // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
			},
		},
		{
			name:    "I64Store32",
			auxType: auxInt64,
			argLen:  3,
			asm:     wasm.AI64Store32,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4295032831},  // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
			},
		},
		{
			name:    "I64Store",
			auxType: auxInt64,
			argLen:  3,
			asm:     wasm.AI64Store,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4295032831},  // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
			},
		},
		{
			name:    "F32Load",
			auxType: auxInt64,
			argLen:  2,
			asm:     wasm.AF32Load,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:    "F64Load",
			auxType: auxInt64,
			argLen:  2,
			asm:     wasm.AF64Load,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:    "F32Store",
			auxType: auxInt64,
			argLen:  3,
			asm:     wasm.AF32Store,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4294901760},  // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
			},
		},
		{
			name:    "F64Store",
			auxType: auxInt64,
			argLen:  3,
			asm:     wasm.AF64Store,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 4294901760},  // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{0, 21474902015}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP SB
				},
			},
		},
		{
			name:              "I64Const",
			auxType:           auxInt64,
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:              "F64Const",
			auxType:           auxFloat64,
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "I64Eqz",
			argLen: 1,
			asm:    wasm.AI64Eqz,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64Eq",
			argLen: 2,
			asm:    wasm.AI64Eq,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64Ne",
			argLen: 2,
			asm:    wasm.AI64Ne,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64LtS",
			argLen: 2,
			asm:    wasm.AI64LtS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64LtU",
			argLen: 2,
			asm:    wasm.AI64LtU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64GtS",
			argLen: 2,
			asm:    wasm.AI64GtS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64GtU",
			argLen: 2,
			asm:    wasm.AI64GtU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64LeS",
			argLen: 2,
			asm:    wasm.AI64LeS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64LeU",
			argLen: 2,
			asm:    wasm.AI64LeU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64GeS",
			argLen: 2,
			asm:    wasm.AI64GeS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64GeU",
			argLen: 2,
			asm:    wasm.AI64GeU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "F64Eq",
			argLen: 2,
			asm:    wasm.AF64Eq,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "F64Ne",
			argLen: 2,
			asm:    wasm.AF64Ne,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "F64Lt",
			argLen: 2,
			asm:    wasm.AF64Lt,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "F64Gt",
			argLen: 2,
			asm:    wasm.AF64Gt,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "F64Le",
			argLen: 2,
			asm:    wasm.AF64Le,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "F64Ge",
			argLen: 2,
			asm:    wasm.AF64Ge,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64Add",
			argLen: 2,
			asm:    wasm.AI64Add,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:    "I64AddConst",
			auxType: auxInt64,
			argLen:  1,
			asm:     wasm.AI64Add,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64Sub",
			argLen: 2,
			asm:    wasm.AI64Sub,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64Mul",
			argLen: 2,
			asm:    wasm.AI64Mul,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64DivS",
			argLen: 2,
			asm:    wasm.AI64DivS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64DivU",
			argLen: 2,
			asm:    wasm.AI64DivU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64RemS",
			argLen: 2,
			asm:    wasm.AI64RemS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64RemU",
			argLen: 2,
			asm:    wasm.AI64RemU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64And",
			argLen: 2,
			asm:    wasm.AI64And,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64Or",
			argLen: 2,
			asm:    wasm.AI64Or,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64Xor",
			argLen: 2,
			asm:    wasm.AI64Xor,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64Shl",
			argLen: 2,
			asm:    wasm.AI64Shl,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64ShrS",
			argLen: 2,
			asm:    wasm.AI64ShrS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64ShrU",
			argLen: 2,
			asm:    wasm.AI64ShrU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
					{1, 4295032831}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 SP
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "F64Neg",
			argLen: 1,
			asm:    wasm.AF64Neg,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "F64Add",
			argLen: 2,
			asm:    wasm.AF64Add,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "F64Sub",
			argLen: 2,
			asm:    wasm.AF64Sub,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "F64Mul",
			argLen: 2,
			asm:    wasm.AF64Mul,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "F64Div",
			argLen: 2,
			asm:    wasm.AF64Div,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
					{1, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "I64TruncSF64",
			argLen: 1,
			asm:    wasm.AI64TruncSF64,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "I64TruncUF64",
			argLen: 1,
			asm:    wasm.AI64TruncUF64,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
				outputs: []outputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
			},
		},
		{
			name:   "F64ConvertSI64",
			argLen: 1,
			asm:    wasm.AF64ConvertSI64,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},
		{
			name:   "F64ConvertUI64",
			argLen: 1,
			asm:    wasm.AF64ConvertUI64,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535}, // R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15
				},
				outputs: []outputInfo{
					{0, 4294901760}, // F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15
				},
			},
		},

		{
			name:        "Add8",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Add16",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Add32",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Add64",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:    "AddPtr",
			argLen:  2,
			generic: true,
		},
		{
			name:        "Add32F",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Add64F",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:    "Sub8",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Sub16",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Sub32",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Sub64",
			argLen:  2,
			generic: true,
		},
		{
			name:    "SubPtr",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Sub32F",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Sub64F",
			argLen:  2,
			generic: true,
		},
		{
			name:        "Mul8",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Mul16",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Mul32",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Mul64",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Mul32F",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Mul64F",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:    "Div32F",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Div64F",
			argLen:  2,
			generic: true,
		},
		{
			name:        "Hmul32",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Hmul32u",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Hmul64",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Hmul64u",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Mul32uhilo",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Mul64uhilo",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:    "Avg32u",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Avg64u",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Div8",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Div8u",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Div16",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Div16u",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Div32",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Div32u",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Div64",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Div64u",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Div128u",
			argLen:  3,
			generic: true,
		},
		{
			name:    "Mod8",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Mod8u",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Mod16",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Mod16u",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Mod32",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Mod32u",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Mod64",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Mod64u",
			argLen:  2,
			generic: true,
		},
		{
			name:        "And8",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "And16",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "And32",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "And64",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Or8",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Or16",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Or32",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Or64",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Xor8",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Xor16",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Xor32",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Xor64",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:    "Lsh8x8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh8x16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh8x32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh8x64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh16x8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh16x16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh16x32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh16x64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh32x8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh32x16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh32x32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh32x64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh64x8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh64x16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh64x32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Lsh64x64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh8x8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh8x16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh8x32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh8x64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh16x8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh16x16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh16x32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh16x64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh32x8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh32x16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh32x32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh32x64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh64x8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh64x16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh64x32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh64x64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh8Ux8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh8Ux16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh8Ux32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh8Ux64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh16Ux8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh16Ux16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh16Ux32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh16Ux64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh32Ux8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh32Ux16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh32Ux32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh32Ux64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh64Ux8",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh64Ux16",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh64Ux32",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:    "Rsh64Ux64",
			auxType: auxBool,
			argLen:  2,
			generic: true,
		},
		{
			name:        "Eq8",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Eq16",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Eq32",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Eq64",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "EqPtr",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:    "EqInter",
			argLen:  2,
			generic: true,
		},
		{
			name:    "EqSlice",
			argLen:  2,
			generic: true,
		},
		{
			name:        "Eq32F",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Eq64F",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Neq8",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Neq16",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Neq32",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Neq64",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "NeqPtr",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:    "NeqInter",
			argLen:  2,
			generic: true,
		},
		{
			name:    "NeqSlice",
			argLen:  2,
			generic: true,
		},
		{
			name:        "Neq32F",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Neq64F",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:    "Less8",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Less8U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Less16",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Less16U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Less32",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Less32U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Less64",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Less64U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Less32F",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Less64F",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Leq8",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Leq8U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Leq16",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Leq16U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Leq32",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Leq32U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Leq64",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Leq64U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Leq32F",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Leq64F",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Greater8",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Greater8U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Greater16",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Greater16U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Greater32",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Greater32U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Greater64",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Greater64U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Greater32F",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Greater64F",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Geq8",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Geq8U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Geq16",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Geq16U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Geq32",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Geq32U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Geq64",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Geq64U",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Geq32F",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Geq64F",
			argLen:  2,
			generic: true,
		},
		{
			name:    "CondSelect",
			argLen:  3,
			generic: true,
		},
		{
			name:        "AndB",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "OrB",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "EqB",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "NeqB",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:    "Not",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Neg8",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Neg16",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Neg32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Neg64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Neg32F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Neg64F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Com8",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Com16",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Com32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Com64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Ctz8",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Ctz16",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Ctz32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Ctz64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Ctz8NonZero",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Ctz16NonZero",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Ctz32NonZero",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Ctz64NonZero",
			argLen:  1,
			generic: true,
		},
		{
			name:    "BitLen8",
			argLen:  1,
			generic: true,
		},
		{
			name:    "BitLen16",
			argLen:  1,
			generic: true,
		},
		{
			name:    "BitLen32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "BitLen64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Bswap32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Bswap64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "BitRev8",
			argLen:  1,
			generic: true,
		},
		{
			name:    "BitRev16",
			argLen:  1,
			generic: true,
		},
		{
			name:    "BitRev32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "BitRev64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "PopCount8",
			argLen:  1,
			generic: true,
		},
		{
			name:    "PopCount16",
			argLen:  1,
			generic: true,
		},
		{
			name:    "PopCount32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "PopCount64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Sqrt",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Floor",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Ceil",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Trunc",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Round",
			argLen:  1,
			generic: true,
		},
		{
			name:    "RoundToEven",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Abs",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Copysign",
			argLen:  2,
			generic: true,
		},
		{
			name:      "Phi",
			argLen:    -1,
			zeroWidth: true,
			generic:   true,
		},
		{
			name:    "Copy",
			argLen:  1,
			generic: true,
		},
		{
			name:         "Convert",
			argLen:       2,
			resultInArg0: true,
			zeroWidth:    true,
			generic:      true,
		},
		{
			name:    "ConstBool",
			auxType: auxBool,
			argLen:  0,
			generic: true,
		},
		{
			name:    "ConstString",
			auxType: auxString,
			argLen:  0,
			generic: true,
		},
		{
			name:    "ConstNil",
			argLen:  0,
			generic: true,
		},
		{
			name:    "Const8",
			auxType: auxInt8,
			argLen:  0,
			generic: true,
		},
		{
			name:    "Const16",
			auxType: auxInt16,
			argLen:  0,
			generic: true,
		},
		{
			name:    "Const32",
			auxType: auxInt32,
			argLen:  0,
			generic: true,
		},
		{
			name:    "Const64",
			auxType: auxInt64,
			argLen:  0,
			generic: true,
		},
		{
			name:    "Const32F",
			auxType: auxFloat32,
			argLen:  0,
			generic: true,
		},
		{
			name:    "Const64F",
			auxType: auxFloat64,
			argLen:  0,
			generic: true,
		},
		{
			name:    "ConstInterface",
			argLen:  0,
			generic: true,
		},
		{
			name:    "ConstSlice",
			argLen:  0,
			generic: true,
		},
		{
			name:      "InitMem",
			argLen:    0,
			zeroWidth: true,
			generic:   true,
		},
		{
			name:      "Arg",
			auxType:   auxSymOff,
			argLen:    0,
			zeroWidth: true,
			symEffect: SymRead,
			generic:   true,
		},
		{
			name:      "Addr",
			auxType:   auxSym,
			argLen:    1,
			symEffect: SymAddr,
			generic:   true,
		},
		{
			name:      "SP",
			argLen:    0,
			zeroWidth: true,
			generic:   true,
		},
		{
			name:      "SB",
			argLen:    0,
			zeroWidth: true,
			generic:   true,
		},
		{
			name:    "Load",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Store",
			auxType: auxTyp,
			argLen:  3,
			generic: true,
		},
		{
			name:    "Move",
			auxType: auxTypSize,
			argLen:  3,
			generic: true,
		},
		{
			name:    "Zero",
			auxType: auxTypSize,
			argLen:  2,
			generic: true,
		},
		{
			name:    "StoreWB",
			auxType: auxTyp,
			argLen:  3,
			generic: true,
		},
		{
			name:    "MoveWB",
			auxType: auxTypSize,
			argLen:  3,
			generic: true,
		},
		{
			name:    "ZeroWB",
			auxType: auxTypSize,
			argLen:  2,
			generic: true,
		},
		{
			name:      "WB",
			auxType:   auxSym,
			argLen:    3,
			symEffect: SymNone,
			generic:   true,
		},
		{
			name:    "ClosureCall",
			auxType: auxInt64,
			argLen:  3,
			call:    true,
			generic: true,
		},
		{
			name:      "StaticCall",
			auxType:   auxSymOff,
			argLen:    1,
			call:      true,
			symEffect: SymNone,
			generic:   true,
		},
		{
			name:    "InterCall",
			auxType: auxInt64,
			argLen:  2,
			call:    true,
			generic: true,
		},
		{
			name:    "SignExt8to16",
			argLen:  1,
			generic: true,
		},
		{
			name:    "SignExt8to32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "SignExt8to64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "SignExt16to32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "SignExt16to64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "SignExt32to64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "ZeroExt8to16",
			argLen:  1,
			generic: true,
		},
		{
			name:    "ZeroExt8to32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "ZeroExt8to64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "ZeroExt16to32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "ZeroExt16to64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "ZeroExt32to64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Trunc16to8",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Trunc32to8",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Trunc32to16",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Trunc64to8",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Trunc64to16",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Trunc64to32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt32to32F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt32to64F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt64to32F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt64to64F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt32Fto32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt32Fto64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt64Fto32",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt64Fto64",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt32Fto64F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt64Fto32F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Round32F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Round64F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "IsNonNil",
			argLen:  1,
			generic: true,
		},
		{
			name:    "IsInBounds",
			argLen:  2,
			generic: true,
		},
		{
			name:    "IsSliceInBounds",
			argLen:  2,
			generic: true,
		},
		{
			name:    "NilCheck",
			argLen:  2,
			generic: true,
		},
		{
			name:      "GetG",
			argLen:    1,
			zeroWidth: true,
			generic:   true,
		},
		{
			name:    "GetClosurePtr",
			argLen:  0,
			generic: true,
		},
		{
			name:    "GetCallerPC",
			argLen:  0,
			generic: true,
		},
		{
			name:    "GetCallerSP",
			argLen:  0,
			generic: true,
		},
		{
			name:    "PtrIndex",
			argLen:  2,
			generic: true,
		},
		{
			name:    "OffPtr",
			auxType: auxInt64,
			argLen:  1,
			generic: true,
		},
		{
			name:    "SliceMake",
			argLen:  3,
			generic: true,
		},
		{
			name:    "SlicePtr",
			argLen:  1,
			generic: true,
		},
		{
			name:    "SliceLen",
			argLen:  1,
			generic: true,
		},
		{
			name:    "SliceCap",
			argLen:  1,
			generic: true,
		},
		{
			name:    "ComplexMake",
			argLen:  2,
			generic: true,
		},
		{
			name:    "ComplexReal",
			argLen:  1,
			generic: true,
		},
		{
			name:    "ComplexImag",
			argLen:  1,
			generic: true,
		},
		{
			name:    "StringMake",
			argLen:  2,
			generic: true,
		},
		{
			name:    "StringPtr",
			argLen:  1,
			generic: true,
		},
		{
			name:    "StringLen",
			argLen:  1,
			generic: true,
		},
		{
			name:    "IMake",
			argLen:  2,
			generic: true,
		},
		{
			name:    "ITab",
			argLen:  1,
			generic: true,
		},
		{
			name:    "IData",
			argLen:  1,
			generic: true,
		},
		{
			name:    "StructMake0",
			argLen:  0,
			generic: true,
		},
		{
			name:    "StructMake1",
			argLen:  1,
			generic: true,
		},
		{
			name:    "StructMake2",
			argLen:  2,
			generic: true,
		},
		{
			name:    "StructMake3",
			argLen:  3,
			generic: true,
		},
		{
			name:    "StructMake4",
			argLen:  4,
			generic: true,
		},
		{
			name:    "StructSelect",
			auxType: auxInt64,
			argLen:  1,
			generic: true,
		},
		{
			name:    "ArrayMake0",
			argLen:  0,
			generic: true,
		},
		{
			name:    "ArrayMake1",
			argLen:  1,
			generic: true,
		},
		{
			name:    "ArraySelect",
			auxType: auxInt64,
			argLen:  1,
			generic: true,
		},
		{
			name:    "StoreReg",
			argLen:  1,
			generic: true,
		},
		{
			name:    "LoadReg",
			argLen:  1,
			generic: true,
		},
		{
			name:      "FwdRef",
			auxType:   auxSym,
			argLen:    0,
			symEffect: SymNone,
			generic:   true,
		},
		{
			name:    "Unknown",
			argLen:  0,
			generic: true,
		},
		{
			name:      "VarDef",
			auxType:   auxSym,
			argLen:    1,
			zeroWidth: true,
			symEffect: SymNone,
			generic:   true,
		},
		{
			name:      "VarKill",
			auxType:   auxSym,
			argLen:    1,
			symEffect: SymNone,
			generic:   true,
		},
		{
			name:      "VarLive",
			auxType:   auxSym,
			argLen:    1,
			zeroWidth: true,
			symEffect: SymRead,
			generic:   true,
		},
		{
			name:      "KeepAlive",
			argLen:    2,
			zeroWidth: true,
			generic:   true,
		},
		{
			name:    "Int64Make",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Int64Hi",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Int64Lo",
			argLen:  1,
			generic: true,
		},
		{
			name:        "Add32carry",
			argLen:      2,
			commutative: true,
			generic:     true,
		},
		{
			name:        "Add32withcarry",
			argLen:      3,
			commutative: true,
			generic:     true,
		},
		{
			name:    "Sub32carry",
			argLen:  2,
			generic: true,
		},
		{
			name:    "Sub32withcarry",
			argLen:  3,
			generic: true,
		},
		{
			name:    "Signmask",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Zeromask",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Slicemask",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt32Uto32F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt32Uto64F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt32Fto32U",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt64Fto32U",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt64Uto32F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt64Uto64F",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt32Fto64U",
			argLen:  1,
			generic: true,
		},
		{
			name:    "Cvt64Fto64U",
			argLen:  1,
			generic: true,
		},
		{
			name:      "Select0",
			argLen:    1,
			zeroWidth: true,
			generic:   true,
		},
		{
			name:      "Select1",
			argLen:    1,
			zeroWidth: true,
			generic:   true,
		},
		{
			name:    "AtomicLoad32",
			argLen:  2,
			generic: true,
		},
		{
			name:    "AtomicLoad64",
			argLen:  2,
			generic: true,
		},
		{
			name:    "AtomicLoadPtr",
			argLen:  2,
			generic: true,
		},
		{
			name:           "AtomicStore32",
			argLen:         3,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicStore64",
			argLen:         3,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicStorePtrNoWB",
			argLen:         3,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicExchange32",
			argLen:         3,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicExchange64",
			argLen:         3,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicAdd32",
			argLen:         3,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicAdd64",
			argLen:         3,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicCompareAndSwap32",
			argLen:         4,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicCompareAndSwap64",
			argLen:         4,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicAnd8",
			argLen:         3,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicOr8",
			argLen:         3,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicAdd32Variant",
			argLen:         3,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:           "AtomicAdd64Variant",
			argLen:         3,
			hasSideEffects: true,
			generic:        true,
		},
		{
			name:      "Clobber",
			auxType:   auxSymOff,
			argLen:    0,
			symEffect: SymNone,
			generic:   true,
		},
	}
	pstate.checkEnabled = false
	pstate.passOrder = [...]constraint{
		// "insert resched checks" uses mem, better to clean out stores first.
		{"dse", "insert resched checks"},
		// insert resched checks adds new blocks containing generic instructions
		{"insert resched checks", "lower"},
		{"insert resched checks", "tighten"},

		// prove relies on common-subexpression elimination for maximum benefits.
		{"generic cse", "prove"},
		// deadcode after prove to eliminate all new dead blocks.
		{"prove", "generic deadcode"},
		// common-subexpression before dead-store elim, so that we recognize
		// when two address expressions are the same.
		{"generic cse", "dse"},
		// cse substantially improves nilcheckelim efficacy
		{"generic cse", "nilcheckelim"},
		// allow deadcode to clean up after nilcheckelim
		{"nilcheckelim", "generic deadcode"},
		// nilcheckelim generates sequences of plain basic blocks
		{"nilcheckelim", "fuse"},
		// nilcheckelim relies on opt to rewrite user nil checks
		{"opt", "nilcheckelim"},
		// tighten will be most effective when as many values have been removed as possible
		{"generic deadcode", "tighten"},
		{"generic cse", "tighten"},
		// checkbce needs the values removed
		{"generic deadcode", "check bce"},
		// don't run optimization pass until we've decomposed builtin objects
		{"decompose builtin", "late opt"},
		// decompose builtin is the last pass that may introduce new float ops, so run softfloat after it
		{"decompose builtin", "softfloat"},
		// don't layout blocks until critical edges have been removed
		{"critical", "layout"},
		// regalloc requires the removal of all critical edges
		{"critical", "regalloc"},
		// regalloc requires all the values in a block to be scheduled
		{"schedule", "regalloc"},
		// checkLower must run after lowering & subsequent dead code elim
		{"lower", "checkLower"},
		{"lowered deadcode", "checkLower"},
		// late nilcheck needs instructions to be scheduled.
		{"schedule", "late nilcheck"},
		// flagalloc needs instructions to be scheduled.
		{"schedule", "flagalloc"},
		// regalloc needs flags to be allocated first.
		{"flagalloc", "regalloc"},
		// loopRotate will confuse regalloc.
		{"regalloc", "loop rotate"},
		// stackframe needs to know about spilled registers.
		{"regalloc", "stackframe"},
		// trim needs regalloc to be done first.
		{"regalloc", "trim"},
	}
	pstate.bllikelies = [4]string{"default", "call", "ret", "exit"}
	pstate.passes = [...]pass{
		// TODO: combine phielim and copyelim into a single pass?
		{name: "number lines", fn: pstate.numberLines, required: true},
		{name: "early phielim", fn: phielim},
		{name: "early copyelim", fn: copyelim},
		{name: "early deadcode", fn: pstate.deadcode}, // remove generated dead code to avoid doing pointless work during opt
		{name: "short circuit", fn: pstate.shortcircuit},
		{name: "decompose user", fn: pstate.decomposeUser, required: true},
		{name: "opt", fn: pstate.opt, required: true},               // TODO: split required rules and optimizing rules
		{name: "zero arg cse", fn: pstate.zcse, required: true},     // required to merge OpSB values
		{name: "opt deadcode", fn: pstate.deadcode, required: true}, // remove any blocks orphaned during opt
		{name: "generic cse", fn: pstate.cse},
		{name: "phiopt", fn: pstate.phiopt},
		{name: "nilcheckelim", fn: pstate.nilcheckelim},
		{name: "prove", fn: pstate.prove},
		{name: "decompose builtin", fn: pstate.decomposeBuiltIn, required: true},
		{name: "softfloat", fn: pstate.softfloat, required: true},
		{name: "late opt", fn: pstate.opt, required: true}, // TODO: split required rules and optimizing rules
		{name: "dead auto elim", fn: pstate.elimDeadAutosGeneric},
		{name: "generic deadcode", fn: pstate.deadcode},
		{name: "check bce", fn: checkbce},
		{name: "branchelim", fn: pstate.branchelim},
		{name: "fuse", fn: fuse},
		{name: "dse", fn: pstate.dse},
		{name: "writebarrier", fn: pstate.writebarrier, required: true}, // expand write barrier ops
		{name: "insert resched checks", fn: pstate.insertLoopReschedChecks,
			disabled: pstate.objabi.Preemptibleloops_enabled == 0}, // insert resched checks in loops.
		{name: "lower", fn: pstate.lower, required: true},
		{name: "lowered cse", fn: pstate.cse},
		{name: "elim unread autos", fn: pstate.elimUnreadAutos},
		{name: "lowered deadcode", fn: pstate.deadcode, required: true},
		{name: "checkLower", fn: pstate.checkLower, required: true},
		{name: "late phielim", fn: phielim},
		{name: "late copyelim", fn: copyelim},
		{name: "tighten", fn: pstate.tighten}, // move values closer to their uses
		{name: "phi tighten", fn: pstate.phiTighten},
		{name: "late deadcode", fn: pstate.deadcode},
		{name: "critical", fn: critical, required: true}, // remove critical edges
		{name: "likelyadjust", fn: pstate.likelyadjust},
		{name: "layout", fn: layout, required: true},            // schedule blocks
		{name: "schedule", fn: pstate.schedule, required: true}, // schedule values
		{name: "late nilcheck", fn: pstate.nilcheckelim2},
		{name: "flagalloc", fn: pstate.flagalloc, required: true}, // allocate flags register
		{name: "regalloc", fn: pstate.regalloc, required: true},   // allocate int & float registers + stack slots
		{name: "loop rotate", fn: pstate.loopRotate},
		{name: "stackframe", fn: stackframe, required: true},
		{name: "trim", fn: pstate.trim}, // remove empty blocks
	}
	pstate.BlockStart = &Value{
		ID:  -10000,
		Op:  OpInvalid,
		Aux: "BlockStart",
	}
	pstate.BlockEnd = &Value{
		ID:  -20000,
		Op:  OpInvalid,
		Aux: "BlockEnd",
	}
	return pstate
}

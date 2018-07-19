package ssa

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/dwarf"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/arm"
	"github.com/dave/golib/src/cmd/internal/obj/arm64"
	"github.com/dave/golib/src/cmd/internal/obj/mips"
	"github.com/dave/golib/src/cmd/internal/obj/ppc64"
	"github.com/dave/golib/src/cmd/internal/obj/s390x"
	"github.com/dave/golib/src/cmd/internal/obj/wasm"
	"github.com/dave/golib/src/cmd/internal/obj/x86"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"github.com/dave/golib/src/cmd/internal/sys"
	"math"
)

type PackageSession struct {
	arm    *arm.PackageSession
	arm64  *arm64.PackageSession
	dwarf  *dwarf.PackageSession
	mips   *mips.PackageSession
	obj    *obj.PackageSession
	objabi *objabi.PackageSession
	ppc64  *ppc64.PackageSession
	s390x  *s390x.PackageSession
	src    *src.PackageSession
	sys    *sys.PackageSession
	types  *types.PackageSession
	wasm   *wasm.PackageSession
	x86    *x86.PackageSession

	BlockEnd   *Value
	BlockStart *Value

	BuildDebug int

	BuildDump         string
	BuildStats        int
	BuildTest         int
	IntrinsicsDebug   int
	IntrinsicsDisable bool
	bllikelies        [4]string

	blockString [96]string

	checkEnabled bool

	checkpointBound limitFact
	checkpointFact  fact

	ctzNonZeroOp        map[Op]Op
	domainRelationTable map[Op]struct {
		d domain
		r relation
	}
	domainStrings [4]string

	dumpFileSeq int

	fpRegMask386 regMask

	fpRegMaskAMD64 regMask

	fpRegMaskARM regMask

	fpRegMaskARM64 regMask

	fpRegMaskMIPS regMask

	fpRegMaskMIPS64 regMask

	fpRegMaskPPC64 regMask

	fpRegMaskS390X regMask

	fpRegMaskWasm      regMask
	framepointerReg386 int8

	framepointerRegAMD64 int8

	framepointerRegARM int8

	framepointerRegARM64 int8

	framepointerRegMIPS int8

	framepointerRegMIPS64 int8

	framepointerRegPPC64 int8

	framepointerRegS390X int8

	framepointerRegWasm int8
	gpRegMask386        regMask

	gpRegMaskAMD64 regMask

	gpRegMaskARM regMask

	gpRegMaskARM64 regMask

	gpRegMaskMIPS regMask

	gpRegMaskMIPS64 regMask

	gpRegMaskPPC64 regMask

	gpRegMaskS390X regMask

	gpRegMaskWasm regMask
	linkReg386    int8

	linkRegAMD64 int8

	linkRegARM int8

	linkRegARM64 int8

	linkRegMIPS int8

	linkRegMIPS64 int8

	linkRegPPC64 int8

	linkRegS390X int8

	linkRegWasm int8
	noLimit     limit

	opMax map[Op]int64
	opMin map[Op]int64

	opcodeTable [2050]opInfo
	passOrder   [26]constraint
	passes      [44]pass

	registers386 [17]Register

	registersAMD64 [33]Register

	registersARM [33]Register

	registersARM64 [64]Register

	registersMIPS [48]Register

	registersMIPS64 [63]Register

	registersPPC64 [64]Register

	registersS390X [33]Register

	registersWasm   [35]Register
	relationStrings [8]string

	reverseBits [8]relation

	ruleFile io.Writer

	specialRegMask386 regMask

	specialRegMaskAMD64 regMask

	specialRegMaskARM regMask

	specialRegMaskARM64 regMask

	specialRegMaskMIPS regMask

	specialRegMaskMIPS64 regMask

	specialRegMaskPPC64 regMask

	specialRegMaskS390X regMask

	specialRegMaskWasm regMask
}

func NewPackageSession(obj_psess *obj.PackageSession, objabi_psess *objabi.PackageSession, types_psess *types.PackageSession, src_psess *src.PackageSession, arm_psess *arm.PackageSession, arm64_psess *arm64.PackageSession, mips_psess *mips.PackageSession, ppc64_psess *ppc64.PackageSession, s390x_psess *s390x.PackageSession, wasm_psess *wasm.PackageSession, x86_psess *x86.PackageSession, dwarf_psess *dwarf.PackageSession, sys_psess *sys.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.obj = obj_psess
	psess.objabi = objabi_psess
	psess.types = types_psess
	psess.src = src_psess
	psess.arm = arm_psess
	psess.arm64 = arm64_psess
	psess.mips = mips_psess
	psess.ppc64 = ppc64_psess
	psess.s390x = s390x_psess
	psess.wasm = wasm_psess
	psess.x86 = x86_psess
	psess.dwarf = dwarf_psess
	psess.sys = sys_psess
	psess.bllikelies = [4]string{"default", "call", "ret", "exit"}
	psess.blockString = [...]string{
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
	psess.registers386 = [...]Register{
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
	psess.gpRegMask386 = regMask(239)
	psess.fpRegMask386 = regMask(65280)
	psess.specialRegMask386 = regMask(0)
	psess.framepointerReg386 = int8(5)
	psess.linkReg386 = int8(-1)
	psess.registersAMD64 = [...]Register{
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
	psess.gpRegMaskAMD64 = regMask(65519)
	psess.fpRegMaskAMD64 = regMask(4294901760)
	psess.specialRegMaskAMD64 = regMask(0)
	psess.framepointerRegAMD64 = int8(5)
	psess.linkRegAMD64 = int8(-1)
	psess.registersARM = [...]Register{
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
	psess.gpRegMaskARM = regMask(21503)
	psess.fpRegMaskARM = regMask(4294901760)
	psess.specialRegMaskARM = regMask(0)
	psess.framepointerRegARM = int8(-1)
	psess.linkRegARM = int8(14)
	psess.registersARM64 = [...]Register{
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
	psess.gpRegMaskARM64 = regMask(670826495)
	psess.fpRegMaskARM64 = regMask(9223372034707292160)
	psess.specialRegMaskARM64 = regMask(0)
	psess.framepointerRegARM64 = int8(-1)
	psess.linkRegARM64 = int8(29)
	psess.registersMIPS = [...]Register{
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
	psess.gpRegMaskMIPS = regMask(335544318)
	psess.fpRegMaskMIPS = regMask(35183835217920)
	psess.specialRegMaskMIPS = regMask(105553116266496)
	psess.framepointerRegMIPS = int8(-1)
	psess.linkRegMIPS = int8(28)
	psess.registersMIPS64 = [...]Register{
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
	psess.gpRegMaskMIPS64 = regMask(167772158)
	psess.fpRegMaskMIPS64 = regMask(1152921504338411520)
	psess.specialRegMaskMIPS64 = regMask(3458764513820540928)
	psess.framepointerRegMIPS64 = int8(-1)
	psess.linkRegMIPS64 = int8(27)
	psess.registersPPC64 = [...]Register{
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
	psess.gpRegMaskPPC64 = regMask(1073733624)
	psess.fpRegMaskPPC64 = regMask(576460743713488896)
	psess.specialRegMaskPPC64 = regMask(0)
	psess.framepointerRegPPC64 = int8(1)
	psess.linkRegPPC64 = int8(-1)
	psess.registersS390X = [...]Register{
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
	psess.gpRegMaskS390X = regMask(23551)
	psess.fpRegMaskS390X = regMask(4294901760)
	psess.specialRegMaskS390X = regMask(0)
	psess.framepointerRegS390X = int8(-1)
	psess.linkRegS390X = int8(14)
	psess.registersWasm = [...]Register{
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
	psess.gpRegMaskWasm = regMask(65535)
	psess.fpRegMaskWasm = regMask(4294901760)
	psess.specialRegMaskWasm = regMask(0)
	psess.framepointerRegWasm = int8(-1)
	psess.linkRegWasm = int8(-1)
	psess.checkEnabled = false
	psess.passOrder = [...]constraint{

		{"dse", "insert resched checks"},

		{"insert resched checks", "lower"},
		{"insert resched checks", "tighten"},

		{"generic cse", "prove"},

		{"prove", "generic deadcode"},

		{"generic cse", "dse"},

		{"generic cse", "nilcheckelim"},

		{"nilcheckelim", "generic deadcode"},

		{"nilcheckelim", "fuse"},

		{"opt", "nilcheckelim"},

		{"generic deadcode", "tighten"},
		{"generic cse", "tighten"},

		{"generic deadcode", "check bce"},

		{"decompose builtin", "late opt"},

		{"decompose builtin", "softfloat"},

		{"critical", "layout"},

		{"critical", "regalloc"},

		{"schedule", "regalloc"},

		{"lower", "checkLower"},
		{"lowered deadcode", "checkLower"},

		{"schedule", "late nilcheck"},

		{"schedule", "flagalloc"},

		{"flagalloc", "regalloc"},

		{"regalloc", "loop rotate"},

		{"regalloc", "stackframe"},

		{"regalloc", "trim"},
	}
	psess.BlockStart = &Value{
		ID:  -10000,
		Op:  OpInvalid,
		Aux: "BlockStart",
	}
	psess.BlockEnd = &Value{
		ID:  -20000,
		Op:  OpInvalid,
		Aux: "BlockEnd",
	}
	psess.opcodeTable = [...]opInfo{
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
					{0, 65280},
					{1, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
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
					{0, 65280},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{1, 65280},
					{0, 65791},
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
					{1, 65280},
					{0, 65791},
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
					{1, 255},
					{2, 65280},
					{0, 65791},
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
					{1, 255},
					{2, 65280},
					{0, 65791},
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
					{1, 255},
					{2, 65280},
					{0, 65791},
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
					{1, 255},
					{2, 65280},
					{0, 65791},
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
					{0, 65280},
					{1, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65791},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{1, 239},
					{0, 255},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 255},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
					{1, 239},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 239},
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
					{0, 239},
					{1, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
					{1, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
					{1, 239},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 239},
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
					{0, 239},
					{1, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
					{1, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 1},
					{1, 255},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{0, 4},
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
					{0, 1},
					{1, 255},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{0, 4},
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
					{0, 1},
					{1, 255},
				},
				outputs: []outputInfo{
					{0, 4},
					{1, 1},
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
					{0, 239},
					{1, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 1},
					{1, 251},
				},
				clobbers: 4,
				outputs: []outputInfo{
					{0, 1},
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
					{0, 1},
					{1, 251},
				},
				clobbers: 4,
				outputs: []outputInfo{
					{0, 1},
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
					{0, 1},
					{1, 251},
				},
				clobbers: 4,
				outputs: []outputInfo{
					{0, 1},
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
					{0, 1},
					{1, 251},
				},
				clobbers: 4,
				outputs: []outputInfo{
					{0, 1},
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
					{0, 1},
					{1, 251},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{0, 4},
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
					{0, 1},
					{1, 251},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{0, 4},
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
					{0, 1},
					{1, 251},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{0, 4},
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
					{0, 1},
					{1, 251},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{0, 4},
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
					{0, 239},
					{1, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
					{1, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
					{1, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "CMPL",
			argLen: 2,
			asm:    x86.ACMPL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255},
					{1, 255},
				},
			},
		},
		{
			name:   "CMPW",
			argLen: 2,
			asm:    x86.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255},
					{1, 255},
				},
			},
		},
		{
			name:   "CMPB",
			argLen: 2,
			asm:    x86.ACMPB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 255},
					{1, 255},
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
					{0, 255},
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
					{0, 255},
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
					{0, 255},
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
					{0, 65280},
					{1, 65280},
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
					{0, 65280},
					{1, 65280},
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
					{0, 255},
					{1, 255},
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
					{0, 255},
					{1, 255},
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
					{0, 255},
					{1, 255},
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
					{0, 255},
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
					{0, 255},
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
					{0, 255},
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
					{1, 2},
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 2},
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 2},
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 2},
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 2},
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 2},
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 2},
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
					{1, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
					{1, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
					{1, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
					{1, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
					{1, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SQRTSD",
			argLen: 1,
			asm:    x86.ASQRTSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
				},
			},
		},
		{
			name:   "SBBLcarrymask",
			argLen: 1,
			asm:    x86.ASBBL,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETEQ",
			argLen: 1,
			asm:    x86.ASETEQ,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETNE",
			argLen: 1,
			asm:    x86.ASETNE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETL",
			argLen: 1,
			asm:    x86.ASETLT,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETLE",
			argLen: 1,
			asm:    x86.ASETLE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETG",
			argLen: 1,
			asm:    x86.ASETGT,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETGE",
			argLen: 1,
			asm:    x86.ASETGE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETB",
			argLen: 1,
			asm:    x86.ASETCS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETBE",
			argLen: 1,
			asm:    x86.ASETLS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETA",
			argLen: 1,
			asm:    x86.ASETHI,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETAE",
			argLen: 1,
			asm:    x86.ASETCC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:         "SETEQF",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ASETEQ,
			reg: regInfo{
				clobbers: 1,
				outputs: []outputInfo{
					{0, 238},
				},
			},
		},
		{
			name:         "SETNEF",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ASETNE,
			reg: regInfo{
				clobbers: 1,
				outputs: []outputInfo{
					{0, 238},
				},
			},
		},
		{
			name:   "SETORD",
			argLen: 1,
			asm:    x86.ASETPC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETNAN",
			argLen: 1,
			asm:    x86.ASETPS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETGF",
			argLen: 1,
			asm:    x86.ASETHI,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "SETGEF",
			argLen: 1,
			asm:    x86.ASETCC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "MOVBLSX",
			argLen: 1,
			asm:    x86.AMOVBLSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "MOVBLZX",
			argLen: 1,
			asm:    x86.AMOVBLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "MOVWLSX",
			argLen: 1,
			asm:    x86.AMOVWLSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "MOVWLZX",
			argLen: 1,
			asm:    x86.AMOVWLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
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
					{0, 65280},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 65280},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
				},
			},
		},
		{
			name:   "CVTSS2SD",
			argLen: 1,
			asm:    x86.ACVTSS2SD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65280},
					{1, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 255},
					{0, 65791},
				},
				outputs: []outputInfo{
					{0, 239},
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
					{1, 255},
					{2, 255},
					{0, 65791},
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
					{1, 255},
					{2, 255},
					{0, 65791},
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
					{1, 255},
					{2, 255},
					{0, 65791},
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
					{1, 255},
					{2, 255},
					{0, 65791},
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
					{1, 255},
					{2, 255},
					{0, 65791},
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
					{0, 65791},
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
					{0, 65791},
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
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{1, 255},
					{0, 65791},
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
					{0, 128},
					{1, 1},
				},
				clobbers: 130,
			},
		},
		{
			name:           "REPSTOSL",
			argLen:         4,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128},
					{1, 2},
					{2, 1},
				},
				clobbers: 130,
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
				clobbers: 65519,
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
					{1, 4},
					{0, 255},
				},
				clobbers: 65519,
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
					{0, 239},
				},
				clobbers: 65519,
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
					{0, 128},
					{1, 64},
				},
				clobbers: 194,
			},
		},
		{
			name:           "REPMOVSL",
			argLen:         4,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128},
					{1, 64},
					{2, 2},
				},
				clobbers: 194,
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
					{0, 239},
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4},
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
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
					{0, 255},
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
					{0, 128},
					{1, 1},
				},
				clobbers: 65280,
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
					{0, 65280},
				},
				outputs: []outputInfo{
					{0, 65280},
				},
			},
		},
		{
			name:    "MOVSSconst1",
			auxType: auxFloat32,
			argLen:  0,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:    "MOVSDconst1",
			auxType: auxFloat64,
			argLen:  0,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 239},
				},
			},
		},
		{
			name:   "MOVSSconst2",
			argLen: 1,
			asm:    x86.AMOVSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 65280},
				},
			},
		},
		{
			name:   "MOVSDconst2",
			argLen: 1,
			asm:    x86.AMOVSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 239},
				},
				outputs: []outputInfo{
					{0, 65280},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
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
					{0, 4294901760},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{1, 4294901760},
					{0, 4295032831},
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
					{1, 4294901760},
					{0, 4295032831},
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
					{1, 65535},
					{2, 4294901760},
					{0, 4295032831},
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
					{1, 65535},
					{2, 4294901760},
					{0, 4295032831},
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
					{1, 65535},
					{2, 4294901760},
					{0, 4295032831},
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
					{1, 65535},
					{2, 4294901760},
					{0, 4295032831},
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
					{0, 4294901760},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{1, 65519},
					{0, 65535},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65519},
					{0, 65535},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65535},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65535},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 1},
					{1, 65535},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{0, 4},
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
					{0, 1},
					{1, 65535},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{0, 4},
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
					{0, 1},
					{1, 65535},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{0, 4},
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
					{0, 1},
					{1, 65535},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{0, 4},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 1},
					{1, 65531},
				},
				outputs: []outputInfo{
					{0, 1},
					{1, 4},
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
					{0, 1},
					{1, 65531},
				},
				outputs: []outputInfo{
					{0, 1},
					{1, 4},
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
					{0, 1},
					{1, 65531},
				},
				outputs: []outputInfo{
					{0, 1},
					{1, 4},
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
					{0, 1},
					{1, 65531},
				},
				outputs: []outputInfo{
					{0, 1},
					{1, 4},
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
					{0, 1},
					{1, 65531},
				},
				outputs: []outputInfo{
					{0, 1},
					{1, 4},
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
					{0, 1},
					{1, 65531},
				},
				outputs: []outputInfo{
					{0, 1},
					{1, 4},
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
					{0, 1},
					{1, 65535},
				},
				outputs: []outputInfo{
					{0, 4},
					{1, 1},
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
					{0, 4},
					{1, 1},
					{2, 65535},
				},
				outputs: []outputInfo{
					{0, 1},
					{1, 4},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "CMPQ",
			argLen: 2,
			asm:    x86.ACMPQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
					{1, 65535},
				},
			},
		},
		{
			name:   "CMPL",
			argLen: 2,
			asm:    x86.ACMPL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
					{1, 65535},
				},
			},
		},
		{
			name:   "CMPW",
			argLen: 2,
			asm:    x86.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
					{1, 65535},
				},
			},
		},
		{
			name:   "CMPB",
			argLen: 2,
			asm:    x86.ACMPB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
					{1, 65535},
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
					{0, 65535},
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
					{0, 65535},
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
					{0, 65535},
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
					{0, 65535},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
				},
			},
		},
		{
			name:   "UCOMISS",
			argLen: 2,
			asm:    x86.AUCOMISS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
			},
		},
		{
			name:   "UCOMISD",
			argLen: 2,
			asm:    x86.AUCOMISD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
			},
		},
		{
			name:   "BTL",
			argLen: 2,
			asm:    x86.ABTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
					{1, 65535},
				},
			},
		},
		{
			name:   "BTQ",
			argLen: 2,
			asm:    x86.ABTQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
					{1, 65535},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65535},
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
					{0, 65535},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65535},
					{1, 65535},
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
					{0, 65535},
					{1, 65535},
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
					{0, 65535},
					{1, 65535},
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
					{0, 65535},
					{1, 65535},
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
					{0, 65535},
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
					{0, 65535},
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
					{0, 65535},
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
					{0, 65535},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 2},
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "BSFQ",
			argLen: 1,
			asm:    x86.ABSFQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "BSRQ",
			argLen: 1,
			asm:    x86.ABSRQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65518},
					{1, 65519},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{0, 65518},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SQRTSD",
			argLen: 1,
			asm:    x86.ASQRTSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "SBBQcarrymask",
			argLen: 1,
			asm:    x86.ASBBQ,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SBBLcarrymask",
			argLen: 1,
			asm:    x86.ASBBL,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETEQ",
			argLen: 1,
			asm:    x86.ASETEQ,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETNE",
			argLen: 1,
			asm:    x86.ASETNE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETL",
			argLen: 1,
			asm:    x86.ASETLT,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETLE",
			argLen: 1,
			asm:    x86.ASETLE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETG",
			argLen: 1,
			asm:    x86.ASETGT,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETGE",
			argLen: 1,
			asm:    x86.ASETGE,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETB",
			argLen: 1,
			asm:    x86.ASETCS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETBE",
			argLen: 1,
			asm:    x86.ASETLS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETA",
			argLen: 1,
			asm:    x86.ASETHI,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETAE",
			argLen: 1,
			asm:    x86.ASETCC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
				},
			},
		},
		{
			name:         "SETEQF",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ASETEQ,
			reg: regInfo{
				clobbers: 1,
				outputs: []outputInfo{
					{0, 65518},
				},
			},
		},
		{
			name:         "SETNEF",
			argLen:       1,
			clobberFlags: true,
			asm:          x86.ASETNE,
			reg: regInfo{
				clobbers: 1,
				outputs: []outputInfo{
					{0, 65518},
				},
			},
		},
		{
			name:   "SETORD",
			argLen: 1,
			asm:    x86.ASETPC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETNAN",
			argLen: 1,
			asm:    x86.ASETPS,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETGF",
			argLen: 1,
			asm:    x86.ASETHI,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "SETGEF",
			argLen: 1,
			asm:    x86.ASETCC,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "MOVBQSX",
			argLen: 1,
			asm:    x86.AMOVBQSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "MOVBQZX",
			argLen: 1,
			asm:    x86.AMOVBLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "MOVWQSX",
			argLen: 1,
			asm:    x86.AMOVWQSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "MOVWQZX",
			argLen: 1,
			asm:    x86.AMOVWLZX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "MOVLQSX",
			argLen: 1,
			asm:    x86.AMOVLQSX,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "MOVLQZX",
			argLen: 1,
			asm:    x86.AMOVL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
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
					{0, 65519},
				},
			},
		},
		{
			name:   "CVTTSD2SL",
			argLen: 1,
			asm:    x86.ACVTTSD2SL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "CVTTSD2SQ",
			argLen: 1,
			asm:    x86.ACVTTSD2SQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "CVTTSS2SL",
			argLen: 1,
			asm:    x86.ACVTTSS2SL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "CVTTSS2SQ",
			argLen: 1,
			asm:    x86.ACVTTSS2SQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "CVTSL2SS",
			argLen: 1,
			asm:    x86.ACVTSL2SS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "CVTSL2SD",
			argLen: 1,
			asm:    x86.ACVTSL2SD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "CVTSQ2SS",
			argLen: 1,
			asm:    x86.ACVTSQ2SS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "CVTSQ2SD",
			argLen: 1,
			asm:    x86.ACVTSQ2SD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "CVTSD2SS",
			argLen: 1,
			asm:    x86.ACVTSD2SS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "CVTSS2SD",
			argLen: 1,
			asm:    x86.ACVTSS2SD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "MOVQi2f",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "MOVQf2i",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:   "MOVLi2f",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65519},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "MOVLf2i",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{1, 4294901760},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 65535},
					{2, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{2, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{2, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{2, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{2, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{2, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{2, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{2, 65535},
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{0, 128},
					{1, 65536},
				},
				clobbers: 128,
			},
		},
		{
			name:              "MOVOconst",
			auxType:           auxInt128,
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:           "REPSTOSQ",
			argLen:         4,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128},
					{1, 2},
					{2, 1},
				},
				clobbers: 130,
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
				clobbers: 4294967279,
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
					{1, 4},
					{0, 65535},
				},
				clobbers: 4294967279,
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
					{0, 65519},
				},
				clobbers: 4294967279,
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
					{0, 128},
					{1, 64},
				},
				clobbers: 65728,
			},
		},
		{
			name:           "REPMOVSQ",
			argLen:         4,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 128},
					{1, 64},
					{2, 2},
				},
				clobbers: 194,
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
					{0, 65519},
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4},
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65535},
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
					{0, 128},
					{1, 1},
				},
				clobbers: 4294901760,
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{0, 65519},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65519},
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
					{1, 1},
					{0, 65519},
					{2, 65519},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{1, 0},
					{0, 65519},
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
					{1, 1},
					{0, 65519},
					{2, 65519},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{1, 0},
					{0, 65519},
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
					{1, 65535},
					{0, 4295032831},
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
					{1, 65535},
					{0, 4295032831},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 30719},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SUB",
			argLen: 2,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "RSB",
			argLen: 2,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:         "CALLudiv",
			argLen:       2,
			clobberFlags: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 2},
					{1, 1},
				},
				clobbers: 16396,
				outputs: []outputInfo{
					{0, 1},
					{1, 2},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SUBS",
			argLen: 2,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
				},
			},
		},
		{
			name:   "SBC",
			argLen: 3,
			asm:    arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
					{1, 21503},
				},
			},
		},
		{
			name:   "MULA",
			argLen: 3,
			asm:    arm.AMULA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MULS",
			argLen: 3,
			asm:    arm.AMULS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "SUBF",
			argLen: 2,
			asm:    arm.ASUBF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "SUBD",
			argLen: 2,
			asm:    arm.ASUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "DIVF",
			argLen: 2,
			asm:    arm.ADIVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "DIVD",
			argLen: 2,
			asm:    arm.ADIVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
					{2, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
					{2, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
					{2, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
					{2, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "BIC",
			argLen: 2,
			asm:    arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MVN",
			argLen: 1,
			asm:    arm.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "NEGF",
			argLen: 1,
			asm:    arm.ANEGF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "NEGD",
			argLen: 1,
			asm:    arm.ANEGD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "SQRTD",
			argLen: 1,
			asm:    arm.ASQRTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "CLZ",
			argLen: 1,
			asm:    arm.ACLZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "REV",
			argLen: 1,
			asm:    arm.AREV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "RBIT",
			argLen: 1,
			asm:    arm.ARBIT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SLL",
			argLen: 2,
			asm:    arm.ASLL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SRL",
			argLen: 2,
			asm:    arm.ASRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SRA",
			argLen: 2,
			asm:    arm.ASRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:    "SRRconst",
			auxType: auxInt32,
			argLen:  1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
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
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
				},
			},
		},
		{
			name:   "ADDshiftLLreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ADDshiftRLreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ADDshiftRAreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SUBshiftLLreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SUBshiftRLreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SUBshiftRAreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "RSBshiftLLreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "RSBshiftRLreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "RSBshiftRAreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ANDshiftLLreg",
			argLen: 3,
			asm:    arm.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ANDshiftRLreg",
			argLen: 3,
			asm:    arm.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ANDshiftRAreg",
			argLen: 3,
			asm:    arm.AAND,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ORshiftLLreg",
			argLen: 3,
			asm:    arm.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ORshiftRLreg",
			argLen: 3,
			asm:    arm.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ORshiftRAreg",
			argLen: 3,
			asm:    arm.AORR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "XORshiftLLreg",
			argLen: 3,
			asm:    arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "XORshiftRLreg",
			argLen: 3,
			asm:    arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "XORshiftRAreg",
			argLen: 3,
			asm:    arm.AEOR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "BICshiftLLreg",
			argLen: 3,
			asm:    arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "BICshiftRLreg",
			argLen: 3,
			asm:    arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "BICshiftRAreg",
			argLen: 3,
			asm:    arm.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MVNshiftLLreg",
			argLen: 2,
			asm:    arm.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MVNshiftRLreg",
			argLen: 2,
			asm:    arm.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MVNshiftRAreg",
			argLen: 2,
			asm:    arm.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
					{1, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ADCshiftLLreg",
			argLen: 4,
			asm:    arm.AADC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ADCshiftRLreg",
			argLen: 4,
			asm:    arm.AADC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ADCshiftRAreg",
			argLen: 4,
			asm:    arm.AADC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SBCshiftLLreg",
			argLen: 4,
			asm:    arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SBCshiftRLreg",
			argLen: 4,
			asm:    arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SBCshiftRAreg",
			argLen: 4,
			asm:    arm.ASBC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "RSCshiftLLreg",
			argLen: 4,
			asm:    arm.ARSC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "RSCshiftRLreg",
			argLen: 4,
			asm:    arm.ARSC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "RSCshiftRAreg",
			argLen: 4,
			asm:    arm.ARSC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "ADDSshiftLLreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
				},
			},
		},
		{
			name:   "ADDSshiftRLreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
				},
			},
		},
		{
			name:   "ADDSshiftRAreg",
			argLen: 3,
			asm:    arm.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
				},
			},
		},
		{
			name:   "SUBSshiftLLreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
				},
			},
		},
		{
			name:   "SUBSshiftRLreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
				},
			},
		},
		{
			name:   "SUBSshiftRAreg",
			argLen: 3,
			asm:    arm.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
				},
			},
		},
		{
			name:   "RSBSshiftLLreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
				},
			},
		},
		{
			name:   "RSBSshiftRLreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
				},
			},
		},
		{
			name:   "RSBSshiftRAreg",
			argLen: 3,
			asm:    arm.ARSB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 21503},
				},
			},
		},
		{
			name:   "CMP",
			argLen: 2,
			asm:    arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
				},
			},
		},
		{
			name:   "CMPF",
			argLen: 2,
			asm:    arm.ACMPF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
			},
		},
		{
			name:   "CMPD",
			argLen: 2,
			asm:    arm.ACMPD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
					{1, 22527},
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
					{0, 22527},
					{1, 22527},
				},
			},
		},
		{
			name:   "CMPshiftLLreg",
			argLen: 3,
			asm:    arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "CMPshiftRLreg",
			argLen: 3,
			asm:    arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "CMPshiftRAreg",
			argLen: 3,
			asm:    arm.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "CMNshiftLLreg",
			argLen: 3,
			asm:    arm.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "CMNshiftRLreg",
			argLen: 3,
			asm:    arm.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "CMNshiftRAreg",
			argLen: 3,
			asm:    arm.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "TSTshiftLLreg",
			argLen: 3,
			asm:    arm.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "TSTshiftRLreg",
			argLen: 3,
			asm:    arm.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "TSTshiftRAreg",
			argLen: 3,
			asm:    arm.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "TEQshiftLLreg",
			argLen: 3,
			asm:    arm.ATEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "TEQshiftRLreg",
			argLen: 3,
			asm:    arm.ATEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "TEQshiftRAreg",
			argLen: 3,
			asm:    arm.ATEQ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
					{2, 21503},
				},
			},
		},
		{
			name:   "CMPF0",
			argLen: 1,
			asm:    arm.ACMPF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "CMPD0",
			argLen: 1,
			asm:    arm.ACMPD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
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
					{0, 21503},
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
					{0, 4294901760},
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
					{0, 4294901760},
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
					{0, 4294975488},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{1, 22527},
					{0, 4294998015},
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
					{1, 22527},
					{0, 4294998015},
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
					{1, 22527},
					{0, 4294998015},
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
					{0, 4294998015},
					{1, 4294901760},
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
					{0, 4294998015},
					{1, 4294901760},
				},
			},
		},
		{
			name:   "MOVWloadidx",
			argLen: 3,
			asm:    arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{1, 22527},
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{1, 22527},
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{1, 22527},
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVBUloadidx",
			argLen: 3,
			asm:    arm.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVBloadidx",
			argLen: 3,
			asm:    arm.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVHUloadidx",
			argLen: 3,
			asm:    arm.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVHloadidx",
			argLen: 3,
			asm:    arm.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},
					{0, 4294998015},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVWstoreidx",
			argLen: 4,
			asm:    arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},
					{2, 22527},
					{0, 4294998015},
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
					{1, 22527},
					{2, 22527},
					{0, 4294998015},
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
					{1, 22527},
					{2, 22527},
					{0, 4294998015},
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
					{1, 22527},
					{2, 22527},
					{0, 4294998015},
				},
			},
		},
		{
			name:   "MOVBstoreidx",
			argLen: 4,
			asm:    arm.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},
					{2, 22527},
					{0, 4294998015},
				},
			},
		},
		{
			name:   "MOVHstoreidx",
			argLen: 4,
			asm:    arm.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 22527},
					{2, 22527},
					{0, 4294998015},
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    arm.AMOVBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVBUreg",
			argLen: 1,
			asm:    arm.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    arm.AMOVHS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVHUreg",
			argLen: 1,
			asm:    arm.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    arm.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:         "MOVWnop",
			argLen:       1,
			resultInArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVWF",
			argLen: 1,
			asm:    arm.AMOVWF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
				},
				clobbers: 2147483648,
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "MOVWD",
			argLen: 1,
			asm:    arm.AMOVWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
				},
				clobbers: 2147483648,
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "MOVWUF",
			argLen: 1,
			asm:    arm.AMOVWF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
				},
				clobbers: 2147483648,
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "MOVWUD",
			argLen: 1,
			asm:    arm.AMOVWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
				},
				clobbers: 2147483648,
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "MOVFW",
			argLen: 1,
			asm:    arm.AMOVFW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				clobbers: 2147483648,
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVDW",
			argLen: 1,
			asm:    arm.AMOVDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				clobbers: 2147483648,
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVFWU",
			argLen: 1,
			asm:    arm.AMOVFW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				clobbers: 2147483648,
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVDWU",
			argLen: 1,
			asm:    arm.AMOVDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				clobbers: 2147483648,
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "MOVFD",
			argLen: 1,
			asm:    arm.AMOVFD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "MOVDF",
			argLen: 1,
			asm:    arm.AMOVDF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "SRAcond",
			argLen: 3,
			asm:    arm.ASRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 21503},
					{1, 21503},
				},
				outputs: []outputInfo{
					{0, 21503},
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
				clobbers: 4294924287,
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
					{1, 128},
					{0, 29695},
				},
				clobbers: 4294924287,
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
					{0, 21503},
				},
				clobbers: 4294924287,
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 22527},
				},
			},
		},
		{
			name:   "Equal",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "NotEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "LessThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "LessEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "GreaterThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "GreaterEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "LessThanU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "LessEqualU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "GreaterThanU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:   "GreaterEqualU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 2},
					{1, 1},
				},
				clobbers: 16386,
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
					{0, 4},
					{1, 2},
				},
				clobbers: 16391,
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
					{0, 2},
					{1, 21503},
					{2, 21503},
				},
				clobbers: 2,
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
					{0, 4},
					{1, 2},
					{2, 21503},
				},
				clobbers: 6,
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 128},
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 4},
					{1, 8},
				},
				clobbers: 4294918144,
			},
		},

		{
			name:        "ADD",
			argLen:      2,
			commutative: true,
			asm:         arm64.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 1878786047},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "SUB",
			argLen: 2,
			asm:    arm64.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "DIV",
			argLen: 2,
			asm:    arm64.ASDIV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "UDIV",
			argLen: 2,
			asm:    arm64.AUDIV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "DIVW",
			argLen: 2,
			asm:    arm64.ASDIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "UDIVW",
			argLen: 2,
			asm:    arm64.AUDIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOD",
			argLen: 2,
			asm:    arm64.AREM,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "UMOD",
			argLen: 2,
			asm:    arm64.AUREM,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MODW",
			argLen: 2,
			asm:    arm64.AREMW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "UMODW",
			argLen: 2,
			asm:    arm64.AUREMW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 9223372034707292160},
					{1, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
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
					{0, 9223372034707292160},
					{1, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FSUBS",
			argLen: 2,
			asm:    arm64.AFSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FSUBD",
			argLen: 2,
			asm:    arm64.AFSUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
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
					{0, 9223372034707292160},
					{1, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
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
					{0, 9223372034707292160},
					{1, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
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
					{0, 9223372034707292160},
					{1, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
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
					{0, 9223372034707292160},
					{1, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FDIVS",
			argLen: 2,
			asm:    arm64.AFDIVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FDIVD",
			argLen: 2,
			asm:    arm64.AFDIVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "BIC",
			argLen: 2,
			asm:    arm64.ABIC,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "EON",
			argLen: 2,
			asm:    arm64.AEON,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "ORN",
			argLen: 2,
			asm:    arm64.AORN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:            "LoweredMuluhilo",
			argLen:          2,
			resultNotInArgs: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
					{1, 670826495},
				},
			},
		},
		{
			name:   "MVN",
			argLen: 1,
			asm:    arm64.AMVN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "NEG",
			argLen: 1,
			asm:    arm64.ANEG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "FNEGS",
			argLen: 1,
			asm:    arm64.AFNEGS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FNEGD",
			argLen: 1,
			asm:    arm64.AFNEGD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FSQRTD",
			argLen: 1,
			asm:    arm64.AFSQRTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "REV",
			argLen: 1,
			asm:    arm64.AREV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "REVW",
			argLen: 1,
			asm:    arm64.AREVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "REV16W",
			argLen: 1,
			asm:    arm64.AREV16W,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "RBIT",
			argLen: 1,
			asm:    arm64.ARBIT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "RBITW",
			argLen: 1,
			asm:    arm64.ARBITW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "CLZ",
			argLen: 1,
			asm:    arm64.ACLZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "CLZW",
			argLen: 1,
			asm:    arm64.ACLZW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "VCNT",
			argLen: 1,
			asm:    arm64.AVCNT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "VUADDLV",
			argLen: 1,
			asm:    arm64.AVUADDLV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
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
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
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
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FMADDS",
			argLen: 3,
			asm:    arm64.AFMADDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
					{2, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FMADDD",
			argLen: 3,
			asm:    arm64.AFMADDD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
					{2, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FNMADDS",
			argLen: 3,
			asm:    arm64.AFNMADDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
					{2, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FNMADDD",
			argLen: 3,
			asm:    arm64.AFNMADDD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
					{2, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FMSUBS",
			argLen: 3,
			asm:    arm64.AFMSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
					{2, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FMSUBD",
			argLen: 3,
			asm:    arm64.AFMSUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
					{2, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FNMSUBS",
			argLen: 3,
			asm:    arm64.AFNMSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
					{2, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FNMSUBD",
			argLen: 3,
			asm:    arm64.AFNMSUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
					{2, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "SLL",
			argLen: 2,
			asm:    arm64.ALSL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "SRL",
			argLen: 2,
			asm:    arm64.ALSR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "SRA",
			argLen: 2,
			asm:    arm64.AASR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "CMP",
			argLen: 2,
			asm:    arm64.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
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
					{0, 805044223},
				},
			},
		},
		{
			name:   "CMPW",
			argLen: 2,
			asm:    arm64.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
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
					{0, 805044223},
				},
			},
		},
		{
			name:   "CMN",
			argLen: 2,
			asm:    arm64.ACMN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
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
					{0, 805044223},
				},
			},
		},
		{
			name:   "CMNW",
			argLen: 2,
			asm:    arm64.ACMNW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
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
					{0, 805044223},
				},
			},
		},
		{
			name:   "TST",
			argLen: 2,
			asm:    arm64.ATST,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
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
					{0, 805044223},
				},
			},
		},
		{
			name:   "TSTW",
			argLen: 2,
			asm:    arm64.ATSTW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
					{1, 805044223},
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
					{0, 805044223},
				},
			},
		},
		{
			name:   "FCMPS",
			argLen: 2,
			asm:    arm64.AFCMPS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
				},
			},
		},
		{
			name:   "FCMPD",
			argLen: 2,
			asm:    arm64.AFCMPD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
					{1, 9223372034707292160},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
					{1, 805044223},
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
					{0, 805044223},
					{1, 805044223},
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
					{0, 805044223},
					{1, 805044223},
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
					{0, 670826495},
					{1, 670826495},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 670826495},
					{1, 670826495},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 670826495},
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
					{0, 9223372034707292160},
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
					{0, 9223372034707292160},
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
					{0, 9223372037928517632},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
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
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "MOVDloadidx",
			argLen: 3,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVWloadidx",
			argLen: 3,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVWUloadidx",
			argLen: 3,
			asm:    arm64.AMOVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVHloadidx",
			argLen: 3,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVHUloadidx",
			argLen: 3,
			asm:    arm64.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVBloadidx",
			argLen: 3,
			asm:    arm64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVBUloadidx",
			argLen: 3,
			asm:    arm64.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVHloadidx2",
			argLen: 3,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVHUloadidx2",
			argLen: 3,
			asm:    arm64.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVWloadidx4",
			argLen: 3,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVWUloadidx4",
			argLen: 3,
			asm:    arm64.AMOVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVDloadidx8",
			argLen: 3,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{1, 805044223},
					{0, 9223372038733561855},
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
					{1, 805044223},
					{0, 9223372038733561855},
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
					{1, 805044223},
					{0, 9223372038733561855},
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
					{1, 805044223},
					{0, 9223372038733561855},
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
					{1, 805044223},
					{2, 805044223},
					{0, 9223372038733561855},
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
					{0, 9223372038733561855},
					{1, 9223372034707292160},
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
					{0, 9223372038733561855},
					{1, 9223372034707292160},
				},
			},
		},
		{
			name:   "MOVBstoreidx",
			argLen: 4,
			asm:    arm64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{2, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVHstoreidx",
			argLen: 4,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{2, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVWstoreidx",
			argLen: 4,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{2, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVDstoreidx",
			argLen: 4,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{2, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVHstoreidx2",
			argLen: 4,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{2, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVWstoreidx4",
			argLen: 4,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{2, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVDstoreidx8",
			argLen: 4,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{2, 805044223},
					{0, 9223372038733561855},
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
					{0, 9223372038733561855},
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
					{0, 9223372038733561855},
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
					{0, 9223372038733561855},
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
					{0, 9223372038733561855},
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
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVBstorezeroidx",
			argLen: 3,
			asm:    arm64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVHstorezeroidx",
			argLen: 3,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVWstorezeroidx",
			argLen: 3,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVDstorezeroidx",
			argLen: 3,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVHstorezeroidx2",
			argLen: 3,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVWstorezeroidx4",
			argLen: 3,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "MOVDstorezeroidx8",
			argLen: 3,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 805044223},
					{0, 9223372038733561855},
				},
			},
		},
		{
			name:   "FMOVDgpfp",
			argLen: 1,
			asm:    arm64.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FMOVDfpgp",
			argLen: 1,
			asm:    arm64.AFMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    arm64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVBUreg",
			argLen: 1,
			asm:    arm64.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    arm64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVHUreg",
			argLen: 1,
			asm:    arm64.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    arm64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVWUreg",
			argLen: 1,
			asm:    arm64.AMOVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "MOVDreg",
			argLen: 1,
			asm:    arm64.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:         "MOVDnop",
			argLen:       1,
			resultInArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "SCVTFWS",
			argLen: 1,
			asm:    arm64.ASCVTFWS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "SCVTFWD",
			argLen: 1,
			asm:    arm64.ASCVTFWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "UCVTFWS",
			argLen: 1,
			asm:    arm64.AUCVTFWS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "UCVTFWD",
			argLen: 1,
			asm:    arm64.AUCVTFWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "SCVTFS",
			argLen: 1,
			asm:    arm64.ASCVTFS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "SCVTFD",
			argLen: 1,
			asm:    arm64.ASCVTFD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "UCVTFS",
			argLen: 1,
			asm:    arm64.AUCVTFS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "UCVTFD",
			argLen: 1,
			asm:    arm64.AUCVTFD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 670826495},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FCVTZSSW",
			argLen: 1,
			asm:    arm64.AFCVTZSSW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "FCVTZSDW",
			argLen: 1,
			asm:    arm64.AFCVTZSDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "FCVTZUSW",
			argLen: 1,
			asm:    arm64.AFCVTZUSW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "FCVTZUDW",
			argLen: 1,
			asm:    arm64.AFCVTZUDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "FCVTZSS",
			argLen: 1,
			asm:    arm64.AFCVTZSS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "FCVTZSD",
			argLen: 1,
			asm:    arm64.AFCVTZSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "FCVTZUS",
			argLen: 1,
			asm:    arm64.AFCVTZUS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "FCVTZUD",
			argLen: 1,
			asm:    arm64.AFCVTZUD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "FCVTSD",
			argLen: 1,
			asm:    arm64.AFCVTSD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FCVTDS",
			argLen: 1,
			asm:    arm64.AFCVTDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FRINTAD",
			argLen: 1,
			asm:    arm64.AFRINTAD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FRINTMD",
			argLen: 1,
			asm:    arm64.AFRINTMD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FRINTPD",
			argLen: 1,
			asm:    arm64.AFRINTPD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
				},
			},
		},
		{
			name:   "FRINTZD",
			argLen: 1,
			asm:    arm64.AFRINTZD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 9223372034707292160},
				},
				outputs: []outputInfo{
					{0, 9223372034707292160},
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
					{0, 670826495},
					{1, 670826495},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 805044223},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
				clobbers: 9223372035512336383,
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
					{1, 67108864},
					{0, 1744568319},
				},
				clobbers: 9223372035512336383,
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
					{0, 670826495},
				},
				clobbers: 9223372035512336383,
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 805044223},
				},
			},
		},
		{
			name:   "Equal",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "NotEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "LessThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "LessEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "GreaterThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "GreaterEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "LessThanU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "LessEqualU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "GreaterThanU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:   "GreaterEqualU",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 65536},
				},
				clobbers: 536936448,
			},
		},
		{
			name:           "LoweredZero",
			argLen:         3,
			clobberFlags:   true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65536},
					{1, 670826495},
				},
				clobbers: 65536,
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
					{0, 131072},
					{1, 65536},
				},
				clobbers: 604176384,
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
					{0, 131072},
					{1, 65536},
					{2, 670826495},
				},
				clobbers: 196608,
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 67108864},
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{1, 805044223},
					{0, 9223372038733561855},
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
					{1, 805044223},
					{0, 9223372038733561855},
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
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{1, 805044223},
					{2, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{1, 805044223},
					{2, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{1, 805044223},
					{0, 9223372038733561855},
				},
				outputs: []outputInfo{
					{0, 670826495},
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
					{0, 4},
					{1, 8},
				},
				clobbers: 9223372035244163072,
			},
		},

		{
			name:        "ADD",
			argLen:      2,
			commutative: true,
			asm:         mips.AADDU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 536870910},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "SUB",
			argLen: 2,
			asm:    mips.ASUBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
					{1, 469762046},
				},
				clobbers: 105553116266496,
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 35184372088832},
					{1, 70368744177664},
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
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 35184372088832},
					{1, 70368744177664},
				},
			},
		},
		{
			name:   "DIV",
			argLen: 2,
			asm:    mips.ADIV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 35184372088832},
					{1, 70368744177664},
				},
			},
		},
		{
			name:   "DIVU",
			argLen: 2,
			asm:    mips.ADIVU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 35184372088832},
					{1, 70368744177664},
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
					{0, 35183835217920},
					{1, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
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
					{0, 35183835217920},
					{1, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "SUBF",
			argLen: 2,
			asm:    mips.ASUBF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
					{1, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "SUBD",
			argLen: 2,
			asm:    mips.ASUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
					{1, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
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
					{0, 35183835217920},
					{1, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
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
					{0, 35183835217920},
					{1, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "DIVF",
			argLen: 2,
			asm:    mips.ADIVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
					{1, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "DIVD",
			argLen: 2,
			asm:    mips.ADIVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
					{1, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
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
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "NEG",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "NEGF",
			argLen: 1,
			asm:    mips.ANEGF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "NEGD",
			argLen: 1,
			asm:    mips.ANEGD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "SQRTD",
			argLen: 1,
			asm:    mips.ASQRTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "SLL",
			argLen: 2,
			asm:    mips.ASLL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "SRL",
			argLen: 2,
			asm:    mips.ASRL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "SRA",
			argLen: 2,
			asm:    mips.ASRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "CLZ",
			argLen: 1,
			asm:    mips.ACLZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "SGT",
			argLen: 2,
			asm:    mips.ASGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "SGTzero",
			argLen: 1,
			asm:    mips.ASGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "SGTU",
			argLen: 2,
			asm:    mips.ASGTU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "SGTUzero",
			argLen: 1,
			asm:    mips.ASGTU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "CMPEQF",
			argLen: 2,
			asm:    mips.ACMPEQF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
					{1, 35183835217920},
				},
			},
		},
		{
			name:   "CMPEQD",
			argLen: 2,
			asm:    mips.ACMPEQD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
					{1, 35183835217920},
				},
			},
		},
		{
			name:   "CMPGEF",
			argLen: 2,
			asm:    mips.ACMPGEF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
					{1, 35183835217920},
				},
			},
		},
		{
			name:   "CMPGED",
			argLen: 2,
			asm:    mips.ACMPGED,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
					{1, 35183835217920},
				},
			},
		},
		{
			name:   "CMPGTF",
			argLen: 2,
			asm:    mips.ACMPGTF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
					{1, 35183835217920},
				},
			},
		},
		{
			name:   "CMPGTD",
			argLen: 2,
			asm:    mips.ACMPGTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
					{1, 35183835217920},
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
					{0, 335544318},
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
					{0, 35183835217920},
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
					{0, 35183835217920},
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
					{0, 140737555464192},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
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
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
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
					{1, 469762046},
					{0, 140738025226238},
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
					{1, 469762046},
					{0, 140738025226238},
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
					{1, 469762046},
					{0, 140738025226238},
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
					{1, 35183835217920},
					{0, 140738025226238},
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
					{1, 35183835217920},
					{0, 140738025226238},
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
					{0, 140738025226238},
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
					{0, 140738025226238},
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
					{0, 140738025226238},
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    mips.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "MOVBUreg",
			argLen: 1,
			asm:    mips.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    mips.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "MOVHUreg",
			argLen: 1,
			asm:    mips.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    mips.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:         "MOVWnop",
			argLen:       1,
			resultInArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 335544318},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 335544318},
					{1, 335544318},
					{2, 335544318},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 335544318},
					{1, 469762046},
				},
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "MOVWF",
			argLen: 1,
			asm:    mips.AMOVWF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "MOVWD",
			argLen: 1,
			asm:    mips.AMOVWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "TRUNCFW",
			argLen: 1,
			asm:    mips.ATRUNCFW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "TRUNCDW",
			argLen: 1,
			asm:    mips.ATRUNCDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "MOVFD",
			argLen: 1,
			asm:    mips.AMOVFD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
				},
			},
		},
		{
			name:   "MOVDF",
			argLen: 1,
			asm:    mips.AMOVDF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 35183835217920},
				},
				outputs: []outputInfo{
					{0, 35183835217920},
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
				clobbers: 140737421246462,
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
					{1, 4194304},
					{0, 402653182},
				},
				clobbers: 140737421246462,
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
					{0, 335544318},
				},
				clobbers: 140737421246462,
			},
		},
		{
			name:           "LoweredAtomicLoad",
			argLen:         2,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{1, 469762046},
					{0, 140738025226238},
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
					{0, 140738025226238},
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
					{1, 469762046},
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{1, 469762046},
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{1, 469762046},
					{2, 469762046},
					{0, 140738025226238},
				},
				outputs: []outputInfo{
					{0, 335544318},
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
					{1, 469762046},
					{0, 140738025226238},
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
					{1, 469762046},
					{0, 140738025226238},
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
					{0, 2},
					{1, 335544318},
				},
				clobbers: 2,
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
					{0, 4},
					{1, 2},
					{2, 335544318},
				},
				clobbers: 6,
			},
		},
		{
			name:           "LoweredNilCheck",
			argLen:         2,
			nilCheck:       true,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 469762046},
				},
			},
		},
		{
			name:   "FPFlagTrue",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:   "FPFlagFalse",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4194304},
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 335544318},
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 335544318},
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
					{0, 1048576},
					{1, 2097152},
				},
				clobbers: 140737219919872,
			},
		},

		{
			name:        "ADDV",
			argLen:      2,
			commutative: true,
			asm:         mips.AADDVU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 268435454},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "SUBV",
			argLen: 2,
			asm:    mips.ASUBVU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 1152921504606846976},
					{1, 2305843009213693952},
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
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 1152921504606846976},
					{1, 2305843009213693952},
				},
			},
		},
		{
			name:   "DIVV",
			argLen: 2,
			asm:    mips.ADIVV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 1152921504606846976},
					{1, 2305843009213693952},
				},
			},
		},
		{
			name:   "DIVVU",
			argLen: 2,
			asm:    mips.ADIVVU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 1152921504606846976},
					{1, 2305843009213693952},
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
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
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
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "SUBF",
			argLen: 2,
			asm:    mips.ASUBF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "SUBD",
			argLen: 2,
			asm:    mips.ASUBD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
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
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
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
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "DIVF",
			argLen: 2,
			asm:    mips.ADIVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "DIVD",
			argLen: 2,
			asm:    mips.ADIVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
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
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "NEGV",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "NEGF",
			argLen: 1,
			asm:    mips.ANEGF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "NEGD",
			argLen: 1,
			asm:    mips.ANEGD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "SQRTD",
			argLen: 1,
			asm:    mips.ASQRTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "SLLV",
			argLen: 2,
			asm:    mips.ASLLV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "SRLV",
			argLen: 2,
			asm:    mips.ASRLV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "SRAV",
			argLen: 2,
			asm:    mips.ASRAV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "SGT",
			argLen: 2,
			asm:    mips.ASGT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "SGTU",
			argLen: 2,
			asm:    mips.ASGTU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
					{1, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "CMPEQF",
			argLen: 2,
			asm:    mips.ACMPEQF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
			},
		},
		{
			name:   "CMPEQD",
			argLen: 2,
			asm:    mips.ACMPEQD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
			},
		},
		{
			name:   "CMPGEF",
			argLen: 2,
			asm:    mips.ACMPGEF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
			},
		},
		{
			name:   "CMPGED",
			argLen: 2,
			asm:    mips.ACMPGED,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
			},
		},
		{
			name:   "CMPGTF",
			argLen: 2,
			asm:    mips.ACMPGTF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
					{1, 1152921504338411520},
				},
			},
		},
		{
			name:   "CMPGTD",
			argLen: 2,
			asm:    mips.ACMPGTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
					{1, 1152921504338411520},
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
					{0, 167772158},
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
					{0, 1152921504338411520},
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
					{0, 1152921504338411520},
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
					{0, 4611686018460942336},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
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
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
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
					{1, 234881022},
					{0, 4611686018695823358},
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
					{1, 234881022},
					{0, 4611686018695823358},
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
					{1, 234881022},
					{0, 4611686018695823358},
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
					{1, 234881022},
					{0, 4611686018695823358},
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
					{0, 4611686018695823358},
					{1, 1152921504338411520},
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
					{0, 4611686018695823358},
					{1, 1152921504338411520},
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
					{0, 4611686018695823358},
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
					{0, 4611686018695823358},
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
					{0, 4611686018695823358},
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
					{0, 4611686018695823358},
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    mips.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "MOVBUreg",
			argLen: 1,
			asm:    mips.AMOVBU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    mips.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "MOVHUreg",
			argLen: 1,
			asm:    mips.AMOVHU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    mips.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "MOVWUreg",
			argLen: 1,
			asm:    mips.AMOVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "MOVVreg",
			argLen: 1,
			asm:    mips.AMOVV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 234881022},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:         "MOVVnop",
			argLen:       1,
			resultInArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 167772158},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "MOVWF",
			argLen: 1,
			asm:    mips.AMOVWF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "MOVWD",
			argLen: 1,
			asm:    mips.AMOVWD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "MOVVF",
			argLen: 1,
			asm:    mips.AMOVVF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "MOVVD",
			argLen: 1,
			asm:    mips.AMOVVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "TRUNCFW",
			argLen: 1,
			asm:    mips.ATRUNCFW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "TRUNCDW",
			argLen: 1,
			asm:    mips.ATRUNCDW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "TRUNCFV",
			argLen: 1,
			asm:    mips.ATRUNCFV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "TRUNCDV",
			argLen: 1,
			asm:    mips.ATRUNCDV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "MOVFD",
			argLen: 1,
			asm:    mips.AMOVFD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
				},
			},
		},
		{
			name:   "MOVDF",
			argLen: 1,
			asm:    mips.AMOVDF,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1152921504338411520},
				},
				outputs: []outputInfo{
					{0, 1152921504338411520},
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
				clobbers: 4611686018393833470,
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
					{1, 4194304},
					{0, 201326590},
				},
				clobbers: 4611686018393833470,
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
					{0, 167772158},
				},
				clobbers: 4611686018393833470,
			},
		},
		{
			name:           "DUFFZERO",
			auxType:        auxInt64,
			argLen:         2,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 167772158},
				},
				clobbers: 134217730,
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
					{0, 2},
					{1, 167772158},
				},
				clobbers: 2,
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
					{0, 4},
					{1, 2},
					{2, 167772158},
				},
				clobbers: 6,
			},
		},
		{
			name:           "LoweredAtomicLoad32",
			argLen:         2,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:           "LoweredAtomicLoad64",
			argLen:         2,
			faultOnNilArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{1, 234881022},
					{0, 4611686018695823358},
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
					{1, 234881022},
					{0, 4611686018695823358},
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
					{0, 4611686018695823358},
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
					{0, 4611686018695823358},
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
					{1, 234881022},
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{1, 234881022},
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{1, 234881022},
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{1, 234881022},
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{1, 234881022},
					{2, 234881022},
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{1, 234881022},
					{2, 234881022},
					{0, 4611686018695823358},
				},
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 234881022},
				},
			},
		},
		{
			name:   "FPFlagTrue",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:   "FPFlagFalse",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4194304},
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 167772158},
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 167772158},
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
					{0, 1048576},
					{1, 2097152},
				},
				clobbers: 4611686018293170176,
			},
		},

		{
			name:        "ADD",
			argLen:      2,
			commutative: true,
			asm:         ppc64.AADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 576460743713488896},
					{1, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
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
					{0, 576460743713488896},
					{1, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "SUB",
			argLen: 2,
			asm:    ppc64.ASUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "FSUB",
			argLen: 2,
			asm:    ppc64.AFSUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
					{1, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FSUBS",
			argLen: 2,
			asm:    ppc64.AFSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
					{1, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
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
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 576460743713488896},
					{1, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
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
					{0, 576460743713488896},
					{1, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FMADD",
			argLen: 3,
			asm:    ppc64.AFMADD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
					{1, 576460743713488896},
					{2, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FMADDS",
			argLen: 3,
			asm:    ppc64.AFMADDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
					{1, 576460743713488896},
					{2, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FMSUB",
			argLen: 3,
			asm:    ppc64.AFMSUB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
					{1, 576460743713488896},
					{2, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FMSUBS",
			argLen: 3,
			asm:    ppc64.AFMSUBS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
					{1, 576460743713488896},
					{2, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "SRAD",
			argLen: 2,
			asm:    ppc64.ASRAD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "SRAW",
			argLen: 2,
			asm:    ppc64.ASRAW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "SRD",
			argLen: 2,
			asm:    ppc64.ASRD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "SRW",
			argLen: 2,
			asm:    ppc64.ASRW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "SLD",
			argLen: 2,
			asm:    ppc64.ASLD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "SLW",
			argLen: 2,
			asm:    ppc64.ASLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "ROTL",
			argLen: 2,
			asm:    ppc64.AROTL,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "ROTLW",
			argLen: 2,
			asm:    ppc64.AROTLW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				clobbers: 2147483648,
			},
		},
		{
			name:   "MaskIfNotCarry",
			argLen: 1,
			asm:    ppc64.AADDME,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "POPCNTD",
			argLen: 1,
			asm:    ppc64.APOPCNTD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "POPCNTW",
			argLen: 1,
			asm:    ppc64.APOPCNTW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "POPCNTB",
			argLen: 1,
			asm:    ppc64.APOPCNTB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "FDIV",
			argLen: 2,
			asm:    ppc64.AFDIV,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
					{1, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FDIVS",
			argLen: 2,
			asm:    ppc64.AFDIVS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
					{1, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "DIVD",
			argLen: 2,
			asm:    ppc64.ADIVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "DIVW",
			argLen: 2,
			asm:    ppc64.ADIVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "DIVDU",
			argLen: 2,
			asm:    ppc64.ADIVDU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "DIVWU",
			argLen: 2,
			asm:    ppc64.ADIVWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "FCTIDZ",
			argLen: 1,
			asm:    ppc64.AFCTIDZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FCTIWZ",
			argLen: 1,
			asm:    ppc64.AFCTIWZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FCFID",
			argLen: 1,
			asm:    ppc64.AFCFID,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FCFIDS",
			argLen: 1,
			asm:    ppc64.AFCFIDS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FRSP",
			argLen: 1,
			asm:    ppc64.AFRSP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "MFVSRD",
			argLen: 1,
			asm:    ppc64.AMFVSRD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "MTVSRD",
			argLen: 1,
			asm:    ppc64.AMTVSRD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733624},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
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
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "ANDN",
			argLen: 2,
			asm:    ppc64.AANDN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "ORN",
			argLen: 2,
			asm:    ppc64.AORN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
					{1, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "NEG",
			argLen: 1,
			asm:    ppc64.ANEG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "FNEG",
			argLen: 1,
			asm:    ppc64.AFNEG,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FSQRT",
			argLen: 1,
			asm:    ppc64.AFSQRT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FSQRTS",
			argLen: 1,
			asm:    ppc64.AFSQRTS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FFLOOR",
			argLen: 1,
			asm:    ppc64.AFRIM,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FCEIL",
			argLen: 1,
			asm:    ppc64.AFRIP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FTRUNC",
			argLen: 1,
			asm:    ppc64.AFRIZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FROUND",
			argLen: 1,
			asm:    ppc64.AFRIN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FABS",
			argLen: 1,
			asm:    ppc64.AFABS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FNABS",
			argLen: 1,
			asm:    ppc64.AFNABS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FCPSGN",
			argLen: 2,
			asm:    ppc64.AFCPSGN,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
					{1, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    ppc64.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "MOVBZreg",
			argLen: 1,
			asm:    ppc64.AMOVBZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    ppc64.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "MOVHZreg",
			argLen: 1,
			asm:    ppc64.AMOVHZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    ppc64.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "MOVWZreg",
			argLen: 1,
			asm:    ppc64.AMOVWZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
					{1, 1073733630},
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
					{0, 1073733630},
					{1, 1073733630},
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
					{0, 1073733630},
					{1, 1073733630},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
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
					{0, 1073733630},
					{1, 1073733630},
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
					{0, 1073733630},
					{1, 1073733630},
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
					{0, 1073733630},
					{1, 1073733630},
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
					{0, 1073733630},
					{1, 1073733630},
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
					{1, 576460743713488896},
					{0, 1073733630},
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
					{1, 576460743713488896},
					{0, 1073733630},
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
					{0, 1073733630},
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
					{0, 1073733630},
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
					{0, 1073733630},
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
					{0, 1073733630},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733624},
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
					{0, 576460743713488896},
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
					{0, 576460743713488896},
				},
			},
		},
		{
			name:   "FCMPU",
			argLen: 2,
			asm:    ppc64.AFCMPU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
					{1, 576460743713488896},
				},
			},
		},
		{
			name:   "CMP",
			argLen: 2,
			asm:    ppc64.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
			},
		},
		{
			name:   "CMPU",
			argLen: 2,
			asm:    ppc64.ACMPU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
			},
		},
		{
			name:   "CMPW",
			argLen: 2,
			asm:    ppc64.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
				},
			},
		},
		{
			name:   "CMPWU",
			argLen: 2,
			asm:    ppc64.ACMPWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
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
					{0, 1073733630},
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
					{0, 1073733630},
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
					{0, 1073733630},
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
					{0, 1073733630},
				},
			},
		},
		{
			name:   "Equal",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "NotEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "LessThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "FLessThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "LessEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "FLessEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "GreaterThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "FGreaterThan",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "GreaterEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:   "FGreaterEqual",
			argLen: 1,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 2048},
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				clobbers: 2147483648,
			},
		},
		{
			name:         "LoweredRound32F",
			argLen:       1,
			resultInArg0: true,
			zeroWidth:    true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
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
					{0, 576460743713488896},
				},
				outputs: []outputInfo{
					{0, 576460743713488896},
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
				clobbers: 576460745860964344,
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
					{0, 4096},
					{1, 2048},
				},
				clobbers: 576460745860964344,
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
					{0, 4096},
				},
				clobbers: 576460745860964344,
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
					{0, 8},
				},
				clobbers: 8,
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
					{0, 8},
					{1, 16},
				},
				clobbers: 1944,
			},
		},
		{
			name:           "LoweredAtomicStore32",
			argLen:         3,
			faultOnNilArg0: true,
			hasSideEffects: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 1073733630},
					{1, 1073733630},
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
					{0, 1073733630},
					{1, 1073733630},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{1, 1073733624},
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{1, 1073733624},
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{1, 1073733624},
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{1, 1073733624},
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{1, 1073733624},
					{2, 1073733624},
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{1, 1073733624},
					{2, 1073733624},
					{0, 1073733630},
				},
				outputs: []outputInfo{
					{0, 1073733624},
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
					{0, 1073733630},
					{1, 1073733630},
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
					{0, 1073733630},
					{1, 1073733630},
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
					{0, 1048576},
					{1, 2097152},
				},
				clobbers: 576460746931503104,
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
					{2, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
					{2, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
					{2, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
					{1, 4294901760},
					{2, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "LPDFR",
			argLen: 1,
			asm:    s390x.ALPDFR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "LNDFR",
			argLen: 1,
			asm:    s390x.ALNDFR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "CPSDR",
			argLen: 2,
			asm:    s390x.ACPSDR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
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
					{0, 4294901760},
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
					{0, 56318},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 56318},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4295023614},
					{1, 4294901760},
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
					{0, 4295023614},
					{1, 4294901760},
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
					{0, 56318},
					{1, 56318},
					{2, 4294901760},
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
					{0, 56318},
					{1, 56318},
					{2, 4294901760},
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
					{1, 23551},
					{0, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 23551},
					{0, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 21503},
					{1, 21503},
				},
				clobbers: 2048,
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				clobbers: 2048,
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				clobbers: 2048,
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				clobbers: 2048,
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				clobbers: 2048,
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				clobbers: 2048,
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				clobbers: 2048,
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				clobbers: 2048,
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				clobbers: 2048,
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 21503},
					{1, 21503},
				},
				clobbers: 2048,
				outputs: []outputInfo{
					{0, 21503},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "CMP",
			argLen: 2,
			asm:    s390x.ACMP,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319},
					{1, 56319},
				},
			},
		},
		{
			name:   "CMPW",
			argLen: 2,
			asm:    s390x.ACMPW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319},
					{1, 56319},
				},
			},
		},
		{
			name:   "CMPU",
			argLen: 2,
			asm:    s390x.ACMPU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319},
					{1, 56319},
				},
			},
		},
		{
			name:   "CMPWU",
			argLen: 2,
			asm:    s390x.ACMPWU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319},
					{1, 56319},
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
					{0, 56319},
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
					{0, 56319},
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
					{0, 56319},
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
					{0, 56319},
				},
			},
		},
		{
			name:   "FCMPS",
			argLen: 2,
			asm:    s390x.ACEBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
			},
		},
		{
			name:   "FCMP",
			argLen: 2,
			asm:    s390x.AFCMPU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
			},
		},
		{
			name:   "SLD",
			argLen: 2,
			asm:    s390x.ASLD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23550},
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "SLW",
			argLen: 2,
			asm:    s390x.ASLW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23550},
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "SRD",
			argLen: 2,
			asm:    s390x.ASRD,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23550},
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "SRW",
			argLen: 2,
			asm:    s390x.ASRW,
			reg: regInfo{
				inputs: []inputInfo{
					{1, 23550},
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 23550},
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 23550},
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "FSQRT",
			argLen: 1,
			asm:    s390x.AFSQRT,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
					{1, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "MOVBreg",
			argLen: 1,
			asm:    s390x.AMOVB,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "MOVBZreg",
			argLen: 1,
			asm:    s390x.AMOVBZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "MOVHreg",
			argLen: 1,
			asm:    s390x.AMOVH,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "MOVHZreg",
			argLen: 1,
			asm:    s390x.AMOVHZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "MOVWreg",
			argLen: 1,
			asm:    s390x.AMOVW,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "MOVWZreg",
			argLen: 1,
			asm:    s390x.AMOVWZ,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "MOVDreg",
			argLen: 1,
			asm:    s390x.AMOVD,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:         "MOVDnop",
			argLen:       1,
			resultInArg0: true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 23551},
				},
			},
		},
		{
			name:   "LDGR",
			argLen: 1,
			asm:    s390x.ALDGR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "LGDR",
			argLen: 1,
			asm:    s390x.ALGDR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "CFDBRA",
			argLen: 1,
			asm:    s390x.ACFDBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "CGDBRA",
			argLen: 1,
			asm:    s390x.ACGDBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "CFEBRA",
			argLen: 1,
			asm:    s390x.ACFEBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "CGEBRA",
			argLen: 1,
			asm:    s390x.ACGEBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "CEFBRA",
			argLen: 1,
			asm:    s390x.ACEFBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "CDFBRA",
			argLen: 1,
			asm:    s390x.ACDFBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "CEGBRA",
			argLen: 1,
			asm:    s390x.ACEGBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "CDGBRA",
			argLen: 1,
			asm:    s390x.ACDGBRA,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "LEDBR",
			argLen: 1,
			asm:    s390x.ALEDBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "LDEBR",
			argLen: 1,
			asm:    s390x.ALDEBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4295000064},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295000064},
					{1, 56318},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "MOVWBR",
			argLen: 1,
			asm:    s390x.AMOVWBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:   "MOVDBR",
			argLen: 1,
			asm:    s390x.AMOVDBR,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 23551},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
					{1, 56319},
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
					{0, 4295023614},
					{1, 56319},
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
					{0, 4295023614},
					{1, 56319},
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
					{0, 4295023614},
					{1, 56319},
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
					{0, 56318},
					{1, 56319},
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
					{0, 56318},
					{1, 56319},
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
					{0, 56318},
					{1, 56319},
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
					{0, 56318},
					{1, 56318},
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
					{1, 56318},
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 56318},
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 56318},
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 56318},
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 56318},
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 56318},
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 56318},
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 56318},
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 56318},
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 56318},
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 56318},
					{1, 56318},
					{2, 56319},
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
					{0, 56318},
					{1, 56318},
					{2, 56319},
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
					{0, 56318},
					{1, 56318},
					{2, 56319},
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
					{0, 56318},
					{1, 56318},
					{2, 56319},
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
					{0, 56318},
					{1, 56318},
					{2, 56319},
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
					{0, 56318},
					{1, 56318},
					{2, 56319},
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
					{0, 56318},
					{1, 56318},
					{2, 56319},
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
					{0, 4295023614},
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
					{0, 4295023614},
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
					{0, 4295023614},
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
					{0, 4295023614},
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
					{0, 23550},
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
				clobbers: 4294933503,
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
					{1, 4096},
					{0, 56318},
				},
				clobbers: 4294933503,
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
					{0, 23550},
				},
				clobbers: 4294933503,
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
					{0, 23551},
				},
			},
		},
		{
			name:      "LoweredGetClosurePtr",
			argLen:    0,
			zeroWidth: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 4096},
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 23551},
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 56318},
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
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 4},
					{1, 8},
				},
				clobbers: 4294918144,
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
					{1, 56319},
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
					{0, 4295023614},
					{1, 56319},
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
					{0, 4295023614},
					{1, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{0, 4295023614},
					{1, 56319},
				},
				outputs: []outputInfo{
					{0, 23551},
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
					{1, 1},
					{0, 56318},
					{2, 56319},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{1, 0},
					{0, 23551},
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
					{1, 1},
					{0, 56318},
					{2, 56319},
				},
				clobbers: 1,
				outputs: []outputInfo{
					{1, 0},
					{0, 23551},
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
					{0, 56318},
					{1, 56318},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 1},
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
					{0, 56318},
					{1, 56318},
				},
				outputs: []outputInfo{
					{1, 0},
					{0, 1},
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
					{0, 23551},
				},
				clobbers: 2,
				outputs: []outputInfo{
					{0, 1},
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
					{1, 2},
					{2, 4},
					{0, 56318},
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
					{1, 2},
					{2, 4},
					{3, 8},
					{0, 56318},
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
					{1, 2},
					{2, 4},
					{3, 8},
					{4, 16},
					{0, 56318},
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
					{1, 2},
					{2, 4},
					{0, 56318},
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
					{1, 2},
					{2, 4},
					{3, 8},
					{0, 56318},
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
					{1, 2},
					{2, 4},
					{3, 8},
					{4, 16},
					{0, 56318},
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
					{0, 2},
					{1, 4},
					{2, 56319},
				},
				clobbers: 6,
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
					{0, 2},
					{1, 56319},
				},
				clobbers: 2,
			},
		},

		{
			name:      "LoweredStaticCall",
			auxType:   auxSymOff,
			argLen:    1,
			call:      true,
			symEffect: SymNone,
			reg: regInfo{
				clobbers: 12884901887,
			},
		},
		{
			name:    "LoweredClosureCall",
			auxType: auxInt64,
			argLen:  3,
			call:    true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
					{1, 65535},
				},
				clobbers: 12884901887,
			},
		},
		{
			name:    "LoweredInterCall",
			auxType: auxInt64,
			argLen:  2,
			call:    true,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
				},
				clobbers: 12884901887,
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:    "LoweredMove",
			auxType: auxInt64,
			argLen:  3,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
					{1, 65535},
				},
			},
		},
		{
			name:    "LoweredZero",
			auxType: auxInt64,
			argLen:  2,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "LoweredGetClosurePtr",
			argLen: 0,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:              "LoweredGetCallerPC",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:              "LoweredGetCallerSP",
			argLen:            0,
			rematerializeable: true,
			reg: regInfo{
				outputs: []outputInfo{
					{0, 65535},
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
					{0, 65535},
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
					{0, 65535},
					{1, 65535},
				},
			},
		},
		{
			name:   "LoweredRound32F",
			argLen: 1,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "LoweredConvert",
			argLen: 2,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "Select",
			argLen: 3,
			asm:    wasm.ASelect,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
					{2, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
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
					{0, 21474902015},
				},
				outputs: []outputInfo{
					{0, 65535},
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
					{0, 21474902015},
				},
				outputs: []outputInfo{
					{0, 65535},
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
					{0, 21474902015},
				},
				outputs: []outputInfo{
					{0, 65535},
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
					{0, 21474902015},
				},
				outputs: []outputInfo{
					{0, 65535},
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
					{0, 21474902015},
				},
				outputs: []outputInfo{
					{0, 65535},
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
					{0, 21474902015},
				},
				outputs: []outputInfo{
					{0, 65535},
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
					{0, 21474902015},
				},
				outputs: []outputInfo{
					{0, 65535},
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
					{1, 4295032831},
					{0, 21474902015},
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
					{1, 4295032831},
					{0, 21474902015},
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
					{1, 4295032831},
					{0, 21474902015},
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
					{1, 4295032831},
					{0, 21474902015},
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
					{0, 21474902015},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{0, 21474902015},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
					{1, 4294901760},
					{0, 21474902015},
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
					{1, 4294901760},
					{0, 21474902015},
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
					{0, 65535},
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
					{0, 4294901760},
				},
			},
		},
		{
			name:   "I64Eqz",
			argLen: 1,
			asm:    wasm.AI64Eqz,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64Eq",
			argLen: 2,
			asm:    wasm.AI64Eq,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64Ne",
			argLen: 2,
			asm:    wasm.AI64Ne,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64LtS",
			argLen: 2,
			asm:    wasm.AI64LtS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64LtU",
			argLen: 2,
			asm:    wasm.AI64LtU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64GtS",
			argLen: 2,
			asm:    wasm.AI64GtS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64GtU",
			argLen: 2,
			asm:    wasm.AI64GtU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64LeS",
			argLen: 2,
			asm:    wasm.AI64LeS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64LeU",
			argLen: 2,
			asm:    wasm.AI64LeU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64GeS",
			argLen: 2,
			asm:    wasm.AI64GeS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64GeU",
			argLen: 2,
			asm:    wasm.AI64GeU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "F64Eq",
			argLen: 2,
			asm:    wasm.AF64Eq,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "F64Ne",
			argLen: 2,
			asm:    wasm.AF64Ne,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "F64Lt",
			argLen: 2,
			asm:    wasm.AF64Lt,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "F64Gt",
			argLen: 2,
			asm:    wasm.AF64Gt,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "F64Le",
			argLen: 2,
			asm:    wasm.AF64Le,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "F64Ge",
			argLen: 2,
			asm:    wasm.AF64Ge,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64Add",
			argLen: 2,
			asm:    wasm.AI64Add,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
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
					{0, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64Sub",
			argLen: 2,
			asm:    wasm.AI64Sub,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64Mul",
			argLen: 2,
			asm:    wasm.AI64Mul,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64DivS",
			argLen: 2,
			asm:    wasm.AI64DivS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64DivU",
			argLen: 2,
			asm:    wasm.AI64DivU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64RemS",
			argLen: 2,
			asm:    wasm.AI64RemS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64RemU",
			argLen: 2,
			asm:    wasm.AI64RemU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64And",
			argLen: 2,
			asm:    wasm.AI64And,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64Or",
			argLen: 2,
			asm:    wasm.AI64Or,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64Xor",
			argLen: 2,
			asm:    wasm.AI64Xor,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64Shl",
			argLen: 2,
			asm:    wasm.AI64Shl,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64ShrS",
			argLen: 2,
			asm:    wasm.AI64ShrS,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64ShrU",
			argLen: 2,
			asm:    wasm.AI64ShrU,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4295032831},
					{1, 4295032831},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "F64Neg",
			argLen: 1,
			asm:    wasm.AF64Neg,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "F64Add",
			argLen: 2,
			asm:    wasm.AF64Add,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "F64Sub",
			argLen: 2,
			asm:    wasm.AF64Sub,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "F64Mul",
			argLen: 2,
			asm:    wasm.AF64Mul,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "F64Div",
			argLen: 2,
			asm:    wasm.AF64Div,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
					{1, 4294901760},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "I64TruncSF64",
			argLen: 1,
			asm:    wasm.AI64TruncSF64,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "I64TruncUF64",
			argLen: 1,
			asm:    wasm.AI64TruncUF64,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 4294901760},
				},
				outputs: []outputInfo{
					{0, 65535},
				},
			},
		},
		{
			name:   "F64ConvertSI64",
			argLen: 1,
			asm:    wasm.AF64ConvertSI64,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
				},
				outputs: []outputInfo{
					{0, 4294901760},
				},
			},
		},
		{
			name:   "F64ConvertUI64",
			argLen: 1,
			asm:    wasm.AF64ConvertUI64,
			reg: regInfo{
				inputs: []inputInfo{
					{0, 65535},
				},
				outputs: []outputInfo{
					{0, 4294901760},
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
	psess.relationStrings = [...]string{
		0: "none", lt: "<", eq: "==", lt | eq: "<=",
		gt: ">", gt | lt: "!=", gt | eq: ">=", gt | eq | lt: "any",
	}
	psess.domainStrings = [...]string{
		"signed", "unsigned", "pointer", "boolean",
	}
	psess.noLimit = limit{math.MinInt64, math.MaxInt64, 0, math.MaxUint64}
	psess.checkpointFact = fact{}
	psess.checkpointBound = limitFact{}
	psess.opMin = map[Op]int64{
		OpAdd64: math.MinInt64, OpSub64: math.MinInt64,
		OpAdd32: math.MinInt32, OpSub32: math.MinInt32,
	}
	psess.opMax = map[Op]int64{
		OpAdd64: math.MaxInt64, OpSub64: math.MaxInt64,
		OpAdd32: math.MaxInt32, OpSub32: math.MaxInt32,
	}
	psess.reverseBits = [...]relation{0, 4, 2, 6, 1, 5, 3, 7}
	psess.domainRelationTable = map[Op]struct {
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

		OpIsInBounds:      {signed | unsigned, lt},
		OpIsSliceInBounds: {signed | unsigned, lt | eq},
	}
	psess.ctzNonZeroOp = map[Op]Op{OpCtz8: OpCtz8NonZero, OpCtz16: OpCtz16NonZero, OpCtz32: OpCtz32NonZero, OpCtz64: OpCtz64NonZero}
	psess.passes = [...]pass{

		{name: "number lines", fn: psess.numberLines, required: true},
		{name: "early phielim", fn: phielim},
		{name: "early copyelim", fn: copyelim},
		{name: "early deadcode", fn: psess.deadcode},
		{name: "short circuit", fn: psess.shortcircuit},
		{name: "decompose user", fn: psess.decomposeUser, required: true},
		{name: "opt", fn: psess.opt, required: true},
		{name: "zero arg cse", fn: psess.zcse, required: true},
		{name: "opt deadcode", fn: psess.deadcode, required: true},
		{name: "generic cse", fn: psess.cse},
		{name: "phiopt", fn: psess.phiopt},
		{name: "nilcheckelim", fn: psess.nilcheckelim},
		{name: "prove", fn: psess.prove},
		{name: "decompose builtin", fn: psess.decomposeBuiltIn, required: true},
		{name: "softfloat", fn: psess.softfloat, required: true},
		{name: "late opt", fn: psess.opt, required: true},
		{name: "dead auto elim", fn: psess.elimDeadAutosGeneric},
		{name: "generic deadcode", fn: psess.deadcode},
		{name: "check bce", fn: checkbce},
		{name: "branchelim", fn: psess.branchelim},
		{name: "fuse", fn: fuse},
		{name: "dse", fn: psess.dse},
		{name: "writebarrier", fn: psess.writebarrier, required: true},
		{name: "insert resched checks", fn: psess.insertLoopReschedChecks,
			disabled: psess.objabi.Preemptibleloops_enabled == 0},
		{name: "lower", fn: psess.lower, required: true},
		{name: "lowered cse", fn: psess.cse},
		{name: "elim unread autos", fn: psess.elimUnreadAutos},
		{name: "lowered deadcode", fn: psess.deadcode, required: true},
		{name: "checkLower", fn: psess.checkLower, required: true},
		{name: "late phielim", fn: phielim},
		{name: "late copyelim", fn: copyelim},
		{name: "tighten", fn: psess.tighten},
		{name: "phi tighten", fn: psess.phiTighten},
		{name: "late deadcode", fn: psess.deadcode},
		{name: "critical", fn: critical, required: true},
		{name: "likelyadjust", fn: psess.likelyadjust},
		{name: "layout", fn: layout, required: true},
		{name: "schedule", fn: psess.schedule, required: true},
		{name: "late nilcheck", fn: psess.nilcheckelim2},
		{name: "flagalloc", fn: psess.flagalloc, required: true},
		{name: "regalloc", fn: psess.regalloc, required: true},
		{name: "loop rotate", fn: psess.loopRotate},
		{name: "stackframe", fn: stackframe, required: true},
		{name: "trim", fn: psess.trim},
	}
	return psess
}

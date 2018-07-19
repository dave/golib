// +build ignore

package main

import "strings"

// Note: registers not used in regalloc are not included in this list,
// so that regmask stays within int64
// Be careful when hand coding regmasks.
var regNamesARM64 = []string{
	"R0",
	"R1",
	"R2",
	"R3",
	"R4",
	"R5",
	"R6",
	"R7",
	"R8",
	"R9",
	"R10",
	"R11",
	"R12",
	"R13",
	"R14",
	"R15",
	"R16",
	"R17",
	"R18",
	"R19",
	"R20",
	"R21",
	"R22",
	"R23",
	"R24",
	"R25",
	"R26",

	"g",
	"R29",
	"R30",
	"SP",

	"F0",
	"F1",
	"F2",
	"F3",
	"F4",
	"F5",
	"F6",
	"F7",
	"F8",
	"F9",
	"F10",
	"F11",
	"F12",
	"F13",
	"F14",
	"F15",
	"F16",
	"F17",
	"F18",
	"F19",
	"F20",
	"F21",
	"F22",
	"F23",
	"F24",
	"F25",
	"F26",
	"F27",
	"F28",
	"F29",
	"F30",
	"F31",

	"SB",
}

func init() {

	if len(regNamesARM64) > 64 {
		panic("too many registers")
	}
	num := map[string]int{}
	for i, name := range regNamesARM64 {
		num[name] = i
	}
	buildReg := func(s string) regMask {
		m := regMask(0)
		for _, r := range strings.Split(s, " ") {
			if n, ok := num[r]; ok {
				m |= regMask(1) << uint(n)
				continue
			}
			panic("register " + r + " not found")
		}
		return m
	}

	// Common individual register masks
	var (
		gp         = buildReg("R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R19 R20 R21 R22 R23 R24 R25 R26 R30")
		gpg        = gp | buildReg("g")
		gpsp       = gp | buildReg("SP")
		gpspg      = gpg | buildReg("SP")
		gpspsbg    = gpspg | buildReg("SB")
		fp         = buildReg("F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31")
		callerSave = gp | fp | buildReg("g") // runtime.setg (and anything calling it) may clobber g
	)
	// Common regInfo
	var (
		gp01      = regInfo{inputs: nil, outputs: []regMask{gp}}
		gp11      = regInfo{inputs: []regMask{gpg}, outputs: []regMask{gp}}
		gp11sp    = regInfo{inputs: []regMask{gpspg}, outputs: []regMask{gp}}
		gp1flags  = regInfo{inputs: []regMask{gpg}}
		gp1flags1 = regInfo{inputs: []regMask{gpg}, outputs: []regMask{gp}}
		gp21      = regInfo{inputs: []regMask{gpg, gpg}, outputs: []regMask{gp}}
		gp21nog   = regInfo{inputs: []regMask{gp, gp}, outputs: []regMask{gp}}
		gp2flags  = regInfo{inputs: []regMask{gpg, gpg}}
		gp2flags1 = regInfo{inputs: []regMask{gp, gp}, outputs: []regMask{gp}}
		gp22      = regInfo{inputs: []regMask{gpg, gpg}, outputs: []regMask{gp, gp}}
		gpload    = regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{gp}}
		gp2load   = regInfo{inputs: []regMask{gpspsbg, gpg}, outputs: []regMask{gp}}
		gpstore   = regInfo{inputs: []regMask{gpspsbg, gpg}}
		gpstore0  = regInfo{inputs: []regMask{gpspsbg}}
		gpstore2  = regInfo{inputs: []regMask{gpspsbg, gpg, gpg}}
		gpxchg    = regInfo{inputs: []regMask{gpspsbg, gpg}, outputs: []regMask{gp}}
		gpcas     = regInfo{inputs: []regMask{gpspsbg, gpg, gpg}, outputs: []regMask{gp}}
		fp01      = regInfo{inputs: nil, outputs: []regMask{fp}}
		fp11      = regInfo{inputs: []regMask{fp}, outputs: []regMask{fp}}
		fpgp      = regInfo{inputs: []regMask{fp}, outputs: []regMask{gp}}
		gpfp      = regInfo{inputs: []regMask{gp}, outputs: []regMask{fp}}
		fp21      = regInfo{inputs: []regMask{fp, fp}, outputs: []regMask{fp}}
		fp31      = regInfo{inputs: []regMask{fp, fp, fp}, outputs: []regMask{fp}}
		fp2flags  = regInfo{inputs: []regMask{fp, fp}}
		fpload    = regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{fp}}
		fpstore   = regInfo{inputs: []regMask{gpspsbg, fp}}
		readflags = regInfo{inputs: nil, outputs: []regMask{gp}}
	)
	ops := []opData{

		{name: "ADD", argLength: 2, reg: gp21, asm: "ADD", commutative: true},
		{name: "ADDconst", argLength: 1, reg: gp11sp, asm: "ADD", aux: "Int64"},
		{name: "SUB", argLength: 2, reg: gp21, asm: "SUB"},
		{name: "SUBconst", argLength: 1, reg: gp11, asm: "SUB", aux: "Int64"},
		{name: "MUL", argLength: 2, reg: gp21, asm: "MUL", commutative: true},
		{name: "MULW", argLength: 2, reg: gp21, asm: "MULW", commutative: true},
		{name: "MNEG", argLength: 2, reg: gp21, asm: "MNEG", commutative: true},
		{name: "MNEGW", argLength: 2, reg: gp21, asm: "MNEGW", commutative: true},
		{name: "MULH", argLength: 2, reg: gp21, asm: "SMULH", commutative: true},
		{name: "UMULH", argLength: 2, reg: gp21, asm: "UMULH", commutative: true},
		{name: "MULL", argLength: 2, reg: gp21, asm: "SMULL", commutative: true},
		{name: "UMULL", argLength: 2, reg: gp21, asm: "UMULL", commutative: true},
		{name: "DIV", argLength: 2, reg: gp21, asm: "SDIV"},
		{name: "UDIV", argLength: 2, reg: gp21, asm: "UDIV"},
		{name: "DIVW", argLength: 2, reg: gp21, asm: "SDIVW"},
		{name: "UDIVW", argLength: 2, reg: gp21, asm: "UDIVW"},
		{name: "MOD", argLength: 2, reg: gp21, asm: "REM"},
		{name: "UMOD", argLength: 2, reg: gp21, asm: "UREM"},
		{name: "MODW", argLength: 2, reg: gp21, asm: "REMW"},
		{name: "UMODW", argLength: 2, reg: gp21, asm: "UREMW"},

		{name: "FADDS", argLength: 2, reg: fp21, asm: "FADDS", commutative: true},
		{name: "FADDD", argLength: 2, reg: fp21, asm: "FADDD", commutative: true},
		{name: "FSUBS", argLength: 2, reg: fp21, asm: "FSUBS"},
		{name: "FSUBD", argLength: 2, reg: fp21, asm: "FSUBD"},
		{name: "FMULS", argLength: 2, reg: fp21, asm: "FMULS", commutative: true},
		{name: "FMULD", argLength: 2, reg: fp21, asm: "FMULD", commutative: true},
		{name: "FNMULS", argLength: 2, reg: fp21, asm: "FNMULS", commutative: true},
		{name: "FNMULD", argLength: 2, reg: fp21, asm: "FNMULD", commutative: true},
		{name: "FDIVS", argLength: 2, reg: fp21, asm: "FDIVS"},
		{name: "FDIVD", argLength: 2, reg: fp21, asm: "FDIVD"},

		{name: "AND", argLength: 2, reg: gp21, asm: "AND", commutative: true},
		{name: "ANDconst", argLength: 1, reg: gp11, asm: "AND", aux: "Int64"},
		{name: "OR", argLength: 2, reg: gp21, asm: "ORR", commutative: true},
		{name: "ORconst", argLength: 1, reg: gp11, asm: "ORR", aux: "Int64"},
		{name: "XOR", argLength: 2, reg: gp21, asm: "EOR", commutative: true},
		{name: "XORconst", argLength: 1, reg: gp11, asm: "EOR", aux: "Int64"},
		{name: "BIC", argLength: 2, reg: gp21, asm: "BIC"},
		{name: "EON", argLength: 2, reg: gp21, asm: "EON"},
		{name: "ORN", argLength: 2, reg: gp21, asm: "ORN"},

		{name: "LoweredMuluhilo", argLength: 2, reg: gp22, resultNotInArgs: true},

		{name: "MVN", argLength: 1, reg: gp11, asm: "MVN"},
		{name: "NEG", argLength: 1, reg: gp11, asm: "NEG"},
		{name: "FNEGS", argLength: 1, reg: fp11, asm: "FNEGS"},
		{name: "FNEGD", argLength: 1, reg: fp11, asm: "FNEGD"},
		{name: "FSQRTD", argLength: 1, reg: fp11, asm: "FSQRTD"},
		{name: "REV", argLength: 1, reg: gp11, asm: "REV"},
		{name: "REVW", argLength: 1, reg: gp11, asm: "REVW"},
		{name: "REV16W", argLength: 1, reg: gp11, asm: "REV16W"},
		{name: "RBIT", argLength: 1, reg: gp11, asm: "RBIT"},
		{name: "RBITW", argLength: 1, reg: gp11, asm: "RBITW"},
		{name: "CLZ", argLength: 1, reg: gp11, asm: "CLZ"},
		{name: "CLZW", argLength: 1, reg: gp11, asm: "CLZW"},
		{name: "VCNT", argLength: 1, reg: fp11, asm: "VCNT"},
		{name: "VUADDLV", argLength: 1, reg: fp11, asm: "VUADDLV"},
		{name: "LoweredRound32F", argLength: 1, reg: fp11, resultInArg0: true, zeroWidth: true},
		{name: "LoweredRound64F", argLength: 1, reg: fp11, resultInArg0: true, zeroWidth: true},

		{name: "FMADDS", argLength: 3, reg: fp31, asm: "FMADDS"},
		{name: "FMADDD", argLength: 3, reg: fp31, asm: "FMADDD"},
		{name: "FNMADDS", argLength: 3, reg: fp31, asm: "FNMADDS"},
		{name: "FNMADDD", argLength: 3, reg: fp31, asm: "FNMADDD"},
		{name: "FMSUBS", argLength: 3, reg: fp31, asm: "FMSUBS"},
		{name: "FMSUBD", argLength: 3, reg: fp31, asm: "FMSUBD"},
		{name: "FNMSUBS", argLength: 3, reg: fp31, asm: "FNMSUBS"},
		{name: "FNMSUBD", argLength: 3, reg: fp31, asm: "FNMSUBD"},

		{name: "SLL", argLength: 2, reg: gp21, asm: "LSL"},
		{name: "SLLconst", argLength: 1, reg: gp11, asm: "LSL", aux: "Int64"},
		{name: "SRL", argLength: 2, reg: gp21, asm: "LSR"},
		{name: "SRLconst", argLength: 1, reg: gp11, asm: "LSR", aux: "Int64"},
		{name: "SRA", argLength: 2, reg: gp21, asm: "ASR"},
		{name: "SRAconst", argLength: 1, reg: gp11, asm: "ASR", aux: "Int64"},
		{name: "RORconst", argLength: 1, reg: gp11, asm: "ROR", aux: "Int64"},
		{name: "RORWconst", argLength: 1, reg: gp11, asm: "RORW", aux: "Int64"},
		{name: "EXTRconst", argLength: 2, reg: gp21, asm: "EXTR", aux: "Int64"},
		{name: "EXTRWconst", argLength: 2, reg: gp21, asm: "EXTRW", aux: "Int64"},

		{name: "CMP", argLength: 2, reg: gp2flags, asm: "CMP", typ: "Flags"},
		{name: "CMPconst", argLength: 1, reg: gp1flags, asm: "CMP", aux: "Int64", typ: "Flags"},
		{name: "CMPW", argLength: 2, reg: gp2flags, asm: "CMPW", typ: "Flags"},
		{name: "CMPWconst", argLength: 1, reg: gp1flags, asm: "CMPW", aux: "Int32", typ: "Flags"},
		{name: "CMN", argLength: 2, reg: gp2flags, asm: "CMN", typ: "Flags"},
		{name: "CMNconst", argLength: 1, reg: gp1flags, asm: "CMN", aux: "Int64", typ: "Flags"},
		{name: "CMNW", argLength: 2, reg: gp2flags, asm: "CMNW", typ: "Flags"},
		{name: "CMNWconst", argLength: 1, reg: gp1flags, asm: "CMNW", aux: "Int32", typ: "Flags"},
		{name: "TST", argLength: 2, reg: gp2flags, asm: "TST", typ: "Flags"},
		{name: "TSTconst", argLength: 1, reg: gp1flags, asm: "TST", aux: "Int64", typ: "Flags"},
		{name: "TSTW", argLength: 2, reg: gp2flags, asm: "TSTW", typ: "Flags"},
		{name: "TSTWconst", argLength: 1, reg: gp1flags, asm: "TSTW", aux: "Int32", typ: "Flags"},
		{name: "FCMPS", argLength: 2, reg: fp2flags, asm: "FCMPS", typ: "Flags"},
		{name: "FCMPD", argLength: 2, reg: fp2flags, asm: "FCMPD", typ: "Flags"},

		{name: "ADDshiftLL", argLength: 2, reg: gp21, asm: "ADD", aux: "Int64"},
		{name: "ADDshiftRL", argLength: 2, reg: gp21, asm: "ADD", aux: "Int64"},
		{name: "ADDshiftRA", argLength: 2, reg: gp21, asm: "ADD", aux: "Int64"},
		{name: "SUBshiftLL", argLength: 2, reg: gp21, asm: "SUB", aux: "Int64"},
		{name: "SUBshiftRL", argLength: 2, reg: gp21, asm: "SUB", aux: "Int64"},
		{name: "SUBshiftRA", argLength: 2, reg: gp21, asm: "SUB", aux: "Int64"},
		{name: "ANDshiftLL", argLength: 2, reg: gp21, asm: "AND", aux: "Int64"},
		{name: "ANDshiftRL", argLength: 2, reg: gp21, asm: "AND", aux: "Int64"},
		{name: "ANDshiftRA", argLength: 2, reg: gp21, asm: "AND", aux: "Int64"},
		{name: "ORshiftLL", argLength: 2, reg: gp21, asm: "ORR", aux: "Int64"},
		{name: "ORshiftRL", argLength: 2, reg: gp21, asm: "ORR", aux: "Int64"},
		{name: "ORshiftRA", argLength: 2, reg: gp21, asm: "ORR", aux: "Int64"},
		{name: "XORshiftLL", argLength: 2, reg: gp21, asm: "EOR", aux: "Int64"},
		{name: "XORshiftRL", argLength: 2, reg: gp21, asm: "EOR", aux: "Int64"},
		{name: "XORshiftRA", argLength: 2, reg: gp21, asm: "EOR", aux: "Int64"},
		{name: "BICshiftLL", argLength: 2, reg: gp21, asm: "BIC", aux: "Int64"},
		{name: "BICshiftRL", argLength: 2, reg: gp21, asm: "BIC", aux: "Int64"},
		{name: "BICshiftRA", argLength: 2, reg: gp21, asm: "BIC", aux: "Int64"},
		{name: "EONshiftLL", argLength: 2, reg: gp21, asm: "EON", aux: "Int64"},
		{name: "EONshiftRL", argLength: 2, reg: gp21, asm: "EON", aux: "Int64"},
		{name: "EONshiftRA", argLength: 2, reg: gp21, asm: "EON", aux: "Int64"},
		{name: "ORNshiftLL", argLength: 2, reg: gp21, asm: "ORN", aux: "Int64"},
		{name: "ORNshiftRL", argLength: 2, reg: gp21, asm: "ORN", aux: "Int64"},
		{name: "ORNshiftRA", argLength: 2, reg: gp21, asm: "ORN", aux: "Int64"},
		{name: "CMPshiftLL", argLength: 2, reg: gp2flags, asm: "CMP", aux: "Int64", typ: "Flags"},
		{name: "CMPshiftRL", argLength: 2, reg: gp2flags, asm: "CMP", aux: "Int64", typ: "Flags"},
		{name: "CMPshiftRA", argLength: 2, reg: gp2flags, asm: "CMP", aux: "Int64", typ: "Flags"},

		{name: "BFI", argLength: 2, reg: gp21nog, asm: "BFI", aux: "Int64", resultInArg0: true},

		{name: "BFXIL", argLength: 2, reg: gp21nog, asm: "BFXIL", aux: "Int64", resultInArg0: true},

		{name: "SBFIZ", argLength: 1, reg: gp11, asm: "SBFIZ", aux: "Int64"},

		{name: "SBFX", argLength: 1, reg: gp11, asm: "SBFX", aux: "Int64"},

		{name: "UBFIZ", argLength: 1, reg: gp11, asm: "UBFIZ", aux: "Int64"},

		{name: "UBFX", argLength: 1, reg: gp11, asm: "UBFX", aux: "Int64"},

		{name: "MOVDconst", argLength: 0, reg: gp01, aux: "Int64", asm: "MOVD", typ: "UInt64", rematerializeable: true},
		{name: "FMOVSconst", argLength: 0, reg: fp01, aux: "Float64", asm: "FMOVS", typ: "Float32", rematerializeable: true},
		{name: "FMOVDconst", argLength: 0, reg: fp01, aux: "Float64", asm: "FMOVD", typ: "Float64", rematerializeable: true},

		{name: "MOVDaddr", argLength: 1, reg: regInfo{inputs: []regMask{buildReg("SP") | buildReg("SB")}, outputs: []regMask{gp}}, aux: "SymOff", asm: "MOVD", rematerializeable: true, symEffect: "Addr"},

		{name: "MOVBload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVB", typ: "Int8", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVBUload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVBU", typ: "UInt8", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVH", typ: "Int16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHUload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVHU", typ: "UInt16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVW", typ: "Int32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWUload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVWU", typ: "UInt32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVDload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVD", typ: "UInt64", faultOnNilArg0: true, symEffect: "Read"},
		{name: "FMOVSload", argLength: 2, reg: fpload, aux: "SymOff", asm: "FMOVS", typ: "Float32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "FMOVDload", argLength: 2, reg: fpload, aux: "SymOff", asm: "FMOVD", typ: "Float64", faultOnNilArg0: true, symEffect: "Read"},

		{name: "MOVDloadidx", argLength: 3, reg: gp2load, asm: "MOVD"},
		{name: "MOVWloadidx", argLength: 3, reg: gp2load, asm: "MOVW"},
		{name: "MOVWUloadidx", argLength: 3, reg: gp2load, asm: "MOVWU"},
		{name: "MOVHloadidx", argLength: 3, reg: gp2load, asm: "MOVH"},
		{name: "MOVHUloadidx", argLength: 3, reg: gp2load, asm: "MOVHU"},
		{name: "MOVBloadidx", argLength: 3, reg: gp2load, asm: "MOVB"},
		{name: "MOVBUloadidx", argLength: 3, reg: gp2load, asm: "MOVBU"},

		{name: "MOVHloadidx2", argLength: 3, reg: gp2load, asm: "MOVH"},
		{name: "MOVHUloadidx2", argLength: 3, reg: gp2load, asm: "MOVHU"},
		{name: "MOVWloadidx4", argLength: 3, reg: gp2load, asm: "MOVW"},
		{name: "MOVWUloadidx4", argLength: 3, reg: gp2load, asm: "MOVWU"},
		{name: "MOVDloadidx8", argLength: 3, reg: gp2load, asm: "MOVD"},

		{name: "MOVBstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "MOVB", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "MOVH", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "MOVW", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVDstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "MOVD", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "STP", argLength: 4, reg: gpstore2, aux: "SymOff", asm: "STP", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "FMOVSstore", argLength: 3, reg: fpstore, aux: "SymOff", asm: "FMOVS", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "FMOVDstore", argLength: 3, reg: fpstore, aux: "SymOff", asm: "FMOVD", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "MOVBstoreidx", argLength: 4, reg: gpstore2, asm: "MOVB", typ: "Mem"},
		{name: "MOVHstoreidx", argLength: 4, reg: gpstore2, asm: "MOVH", typ: "Mem"},
		{name: "MOVWstoreidx", argLength: 4, reg: gpstore2, asm: "MOVW", typ: "Mem"},
		{name: "MOVDstoreidx", argLength: 4, reg: gpstore2, asm: "MOVD", typ: "Mem"},

		{name: "MOVHstoreidx2", argLength: 4, reg: gpstore2, asm: "MOVH", typ: "Mem"},
		{name: "MOVWstoreidx4", argLength: 4, reg: gpstore2, asm: "MOVW", typ: "Mem"},
		{name: "MOVDstoreidx8", argLength: 4, reg: gpstore2, asm: "MOVD", typ: "Mem"},

		{name: "MOVBstorezero", argLength: 2, reg: gpstore0, aux: "SymOff", asm: "MOVB", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHstorezero", argLength: 2, reg: gpstore0, aux: "SymOff", asm: "MOVH", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstorezero", argLength: 2, reg: gpstore0, aux: "SymOff", asm: "MOVW", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVDstorezero", argLength: 2, reg: gpstore0, aux: "SymOff", asm: "MOVD", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVQstorezero", argLength: 2, reg: gpstore0, aux: "SymOff", asm: "STP", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "MOVBstorezeroidx", argLength: 3, reg: gpstore, asm: "MOVB", typ: "Mem"},
		{name: "MOVHstorezeroidx", argLength: 3, reg: gpstore, asm: "MOVH", typ: "Mem"},
		{name: "MOVWstorezeroidx", argLength: 3, reg: gpstore, asm: "MOVW", typ: "Mem"},
		{name: "MOVDstorezeroidx", argLength: 3, reg: gpstore, asm: "MOVD", typ: "Mem"},

		{name: "MOVHstorezeroidx2", argLength: 3, reg: gpstore, asm: "MOVH", typ: "Mem"},
		{name: "MOVWstorezeroidx4", argLength: 3, reg: gpstore, asm: "MOVW", typ: "Mem"},
		{name: "MOVDstorezeroidx8", argLength: 3, reg: gpstore, asm: "MOVD", typ: "Mem"},

		{name: "FMOVDgpfp", argLength: 1, reg: gpfp, asm: "FMOVD"},
		{name: "FMOVDfpgp", argLength: 1, reg: fpgp, asm: "FMOVD"},

		{name: "MOVBreg", argLength: 1, reg: gp11, asm: "MOVB"},
		{name: "MOVBUreg", argLength: 1, reg: gp11, asm: "MOVBU"},
		{name: "MOVHreg", argLength: 1, reg: gp11, asm: "MOVH"},
		{name: "MOVHUreg", argLength: 1, reg: gp11, asm: "MOVHU"},
		{name: "MOVWreg", argLength: 1, reg: gp11, asm: "MOVW"},
		{name: "MOVWUreg", argLength: 1, reg: gp11, asm: "MOVWU"},
		{name: "MOVDreg", argLength: 1, reg: gp11, asm: "MOVD"},

		{name: "MOVDnop", argLength: 1, reg: regInfo{inputs: []regMask{gp}, outputs: []regMask{gp}}, resultInArg0: true},

		{name: "SCVTFWS", argLength: 1, reg: gpfp, asm: "SCVTFWS"},
		{name: "SCVTFWD", argLength: 1, reg: gpfp, asm: "SCVTFWD"},
		{name: "UCVTFWS", argLength: 1, reg: gpfp, asm: "UCVTFWS"},
		{name: "UCVTFWD", argLength: 1, reg: gpfp, asm: "UCVTFWD"},
		{name: "SCVTFS", argLength: 1, reg: gpfp, asm: "SCVTFS"},
		{name: "SCVTFD", argLength: 1, reg: gpfp, asm: "SCVTFD"},
		{name: "UCVTFS", argLength: 1, reg: gpfp, asm: "UCVTFS"},
		{name: "UCVTFD", argLength: 1, reg: gpfp, asm: "UCVTFD"},
		{name: "FCVTZSSW", argLength: 1, reg: fpgp, asm: "FCVTZSSW"},
		{name: "FCVTZSDW", argLength: 1, reg: fpgp, asm: "FCVTZSDW"},
		{name: "FCVTZUSW", argLength: 1, reg: fpgp, asm: "FCVTZUSW"},
		{name: "FCVTZUDW", argLength: 1, reg: fpgp, asm: "FCVTZUDW"},
		{name: "FCVTZSS", argLength: 1, reg: fpgp, asm: "FCVTZSS"},
		{name: "FCVTZSD", argLength: 1, reg: fpgp, asm: "FCVTZSD"},
		{name: "FCVTZUS", argLength: 1, reg: fpgp, asm: "FCVTZUS"},
		{name: "FCVTZUD", argLength: 1, reg: fpgp, asm: "FCVTZUD"},
		{name: "FCVTSD", argLength: 1, reg: fp11, asm: "FCVTSD"},
		{name: "FCVTDS", argLength: 1, reg: fp11, asm: "FCVTDS"},

		{name: "FRINTAD", argLength: 1, reg: fp11, asm: "FRINTAD"},
		{name: "FRINTMD", argLength: 1, reg: fp11, asm: "FRINTMD"},
		{name: "FRINTPD", argLength: 1, reg: fp11, asm: "FRINTPD"},
		{name: "FRINTZD", argLength: 1, reg: fp11, asm: "FRINTZD"},

		{name: "CSEL", argLength: 3, reg: gp2flags1, asm: "CSEL", aux: "CCop"},
		{name: "CSEL0", argLength: 2, reg: gp1flags1, asm: "CSEL", aux: "CCop"},

		{name: "CALLstatic", argLength: 1, reg: regInfo{clobbers: callerSave}, aux: "SymOff", clobberFlags: true, call: true, symEffect: "None"},
		{name: "CALLclosure", argLength: 3, reg: regInfo{inputs: []regMask{gpsp, buildReg("R26"), 0}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},
		{name: "CALLinter", argLength: 2, reg: regInfo{inputs: []regMask{gp}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},

		{name: "LoweredNilCheck", argLength: 2, reg: regInfo{inputs: []regMask{gpg}}, nilCheck: true, faultOnNilArg0: true},

		{name: "Equal", argLength: 1, reg: readflags},
		{name: "NotEqual", argLength: 1, reg: readflags},
		{name: "LessThan", argLength: 1, reg: readflags},
		{name: "LessEqual", argLength: 1, reg: readflags},
		{name: "GreaterThan", argLength: 1, reg: readflags},
		{name: "GreaterEqual", argLength: 1, reg: readflags},
		{name: "LessThanU", argLength: 1, reg: readflags},
		{name: "LessEqualU", argLength: 1, reg: readflags},
		{name: "GreaterThanU", argLength: 1, reg: readflags},
		{name: "GreaterEqualU", argLength: 1, reg: readflags},

		{
			name:      "DUFFZERO",
			aux:       "Int64",
			argLength: 2,
			reg: regInfo{
				inputs:   []regMask{buildReg("R16")},
				clobbers: buildReg("R16 R30"),
			},
			faultOnNilArg0: true,
		},

		{
			name:      "LoweredZero",
			argLength: 3,
			reg: regInfo{
				inputs:   []regMask{buildReg("R16"), gp},
				clobbers: buildReg("R16"),
			},
			clobberFlags:   true,
			faultOnNilArg0: true,
		},

		{
			name:      "DUFFCOPY",
			aux:       "Int64",
			argLength: 3,
			reg: regInfo{
				inputs:   []regMask{buildReg("R17"), buildReg("R16")},
				clobbers: buildReg("R16 R17 R26 R30"),
			},
			faultOnNilArg0: true,
			faultOnNilArg1: true,
		},

		{
			name:      "LoweredMove",
			argLength: 4,
			reg: regInfo{
				inputs:   []regMask{buildReg("R17"), buildReg("R16"), gp},
				clobbers: buildReg("R16 R17"),
			},
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
		},

		{name: "LoweredGetClosurePtr", reg: regInfo{outputs: []regMask{buildReg("R26")}}, zeroWidth: true},

		{name: "LoweredGetCallerSP", reg: gp01, rematerializeable: true},

		{name: "LoweredGetCallerPC", reg: gp01, rematerializeable: true},

		{name: "FlagEQ"},
		{name: "FlagLT_ULT"},
		{name: "FlagLT_UGT"},
		{name: "FlagGT_UGT"},
		{name: "FlagGT_ULT"},

		{name: "InvertFlags", argLength: 1},

		{name: "LDAR", argLength: 2, reg: gpload, asm: "LDAR", faultOnNilArg0: true},
		{name: "LDARW", argLength: 2, reg: gpload, asm: "LDARW", faultOnNilArg0: true},

		{name: "STLR", argLength: 3, reg: gpstore, asm: "STLR", faultOnNilArg0: true, hasSideEffects: true},
		{name: "STLRW", argLength: 3, reg: gpstore, asm: "STLRW", faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicExchange64", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicExchange32", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicAdd64", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicAdd32", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicAdd64Variant", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicAdd32Variant", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicCas64", argLength: 4, reg: gpcas, resultNotInArgs: true, clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicCas32", argLength: 4, reg: gpcas, resultNotInArgs: true, clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicAnd8", argLength: 3, reg: gpxchg, resultNotInArgs: true, asm: "AND", typ: "(UInt8,Mem)", faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicOr8", argLength: 3, reg: gpxchg, resultNotInArgs: true, asm: "ORR", typ: "(UInt8,Mem)", faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredWB", argLength: 3, reg: regInfo{inputs: []regMask{buildReg("R2"), buildReg("R3")}, clobbers: (callerSave &^ gpg) | buildReg("R30")}, clobberFlags: true, aux: "Sym", symEffect: "None"},
	}

	blocks := []blockData{
		{name: "EQ"},
		{name: "NE"},
		{name: "LT"},
		{name: "LE"},
		{name: "GT"},
		{name: "GE"},
		{name: "ULT"},
		{name: "ULE"},
		{name: "UGT"},
		{name: "UGE"},
		{name: "Z"},
		{name: "NZ"},
		{name: "ZW"},
		{name: "NZW"},
		{name: "TBZ"},
		{name: "TBNZ"},
	}

	archs = append(archs, arch{
		name:            "ARM64",
		pkg:             "github.com/dave/golib/src/cmd/internal/obj/arm64",
		genfile:         "../../arm64/ssa.go",
		ops:             ops,
		blocks:          blocks,
		regnames:        regNamesARM64,
		gpregmask:       gp,
		fpregmask:       fp,
		framepointerreg: -1,
		linkreg:         int8(num["R30"]),
	})
}

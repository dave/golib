// +build ignore

package main

import "strings"

var regNamesPPC64 = []string{
	"R0",
	"SP",
	"SB",
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
	"R27",
	"R28",
	"R29",
	"g",
	"R31",

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
}

func init() {

	if len(regNamesPPC64) > 64 {
		panic("too many registers")
	}
	num := map[string]int{}
	for i, name := range regNamesPPC64 {
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

	var (
		gp = buildReg("R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R16 R17 R18 R19 R20 R21 R22 R23 R24 R25 R26 R27 R28 R29")
		fp = buildReg("F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26")
		sp = buildReg("SP")
		sb = buildReg("SB")
		gr = buildReg("g")
		// cr  = buildReg("CR")
		// ctr = buildReg("CTR")
		// lr  = buildReg("LR")
		tmp     = buildReg("R31")
		ctxt    = buildReg("R11")
		callptr = buildReg("R12")
		// tls = buildReg("R13")
		gp01        = regInfo{inputs: nil, outputs: []regMask{gp}}
		gp11        = regInfo{inputs: []regMask{gp | sp | sb}, outputs: []regMask{gp}}
		gp21        = regInfo{inputs: []regMask{gp | sp | sb, gp | sp | sb}, outputs: []regMask{gp}}
		gp1cr       = regInfo{inputs: []regMask{gp | sp | sb}}
		gp2cr       = regInfo{inputs: []regMask{gp | sp | sb, gp | sp | sb}}
		crgp        = regInfo{inputs: nil, outputs: []regMask{gp}}
		gpload      = regInfo{inputs: []regMask{gp | sp | sb}, outputs: []regMask{gp}}
		gpstore     = regInfo{inputs: []regMask{gp | sp | sb, gp | sp | sb}}
		gpstorezero = regInfo{inputs: []regMask{gp | sp | sb}} // ppc64.REGZERO is reserved zero value
		gpxchg      = regInfo{inputs: []regMask{gp | sp | sb, gp}, outputs: []regMask{gp}}
		gpcas       = regInfo{inputs: []regMask{gp | sp | sb, gp, gp}, outputs: []regMask{gp}}
		fp01        = regInfo{inputs: nil, outputs: []regMask{fp}}
		fp11        = regInfo{inputs: []regMask{fp}, outputs: []regMask{fp}}
		fpgp        = regInfo{inputs: []regMask{fp}, outputs: []regMask{gp}}
		gpfp        = regInfo{inputs: []regMask{gp}, outputs: []regMask{fp}}
		fp21        = regInfo{inputs: []regMask{fp, fp}, outputs: []regMask{fp}}
		fp31        = regInfo{inputs: []regMask{fp, fp, fp}, outputs: []regMask{fp}}
		fp2cr       = regInfo{inputs: []regMask{fp, fp}}
		fpload      = regInfo{inputs: []regMask{gp | sp | sb}, outputs: []regMask{fp}}
		fpstore     = regInfo{inputs: []regMask{gp | sp | sb, fp}}
		callerSave  = regMask(gp | fp | gr)
	)
	ops := []opData{
		{name: "ADD", argLength: 2, reg: gp21, asm: "ADD", commutative: true},
		{name: "ADDconst", argLength: 1, reg: gp11, asm: "ADD", aux: "Int64"},
		{name: "FADD", argLength: 2, reg: fp21, asm: "FADD", commutative: true},
		{name: "FADDS", argLength: 2, reg: fp21, asm: "FADDS", commutative: true},
		{name: "SUB", argLength: 2, reg: gp21, asm: "SUB"},
		{name: "FSUB", argLength: 2, reg: fp21, asm: "FSUB"},
		{name: "FSUBS", argLength: 2, reg: fp21, asm: "FSUBS"},

		{name: "MULLD", argLength: 2, reg: gp21, asm: "MULLD", typ: "Int64", commutative: true},
		{name: "MULLW", argLength: 2, reg: gp21, asm: "MULLW", typ: "Int32", commutative: true},

		{name: "MULHD", argLength: 2, reg: gp21, asm: "MULHD", commutative: true},
		{name: "MULHW", argLength: 2, reg: gp21, asm: "MULHW", commutative: true},
		{name: "MULHDU", argLength: 2, reg: gp21, asm: "MULHDU", commutative: true},
		{name: "MULHWU", argLength: 2, reg: gp21, asm: "MULHWU", commutative: true},

		{name: "FMUL", argLength: 2, reg: fp21, asm: "FMUL", commutative: true},
		{name: "FMULS", argLength: 2, reg: fp21, asm: "FMULS", commutative: true},

		{name: "FMADD", argLength: 3, reg: fp31, asm: "FMADD"},
		{name: "FMADDS", argLength: 3, reg: fp31, asm: "FMADDS"},
		{name: "FMSUB", argLength: 3, reg: fp31, asm: "FMSUB"},
		{name: "FMSUBS", argLength: 3, reg: fp31, asm: "FMSUBS"},

		{name: "SRAD", argLength: 2, reg: gp21, asm: "SRAD"},
		{name: "SRAW", argLength: 2, reg: gp21, asm: "SRAW"},
		{name: "SRD", argLength: 2, reg: gp21, asm: "SRD"},
		{name: "SRW", argLength: 2, reg: gp21, asm: "SRW"},
		{name: "SLD", argLength: 2, reg: gp21, asm: "SLD"},
		{name: "SLW", argLength: 2, reg: gp21, asm: "SLW"},

		{name: "ROTL", argLength: 2, reg: gp21, asm: "ROTL"},
		{name: "ROTLW", argLength: 2, reg: gp21, asm: "ROTLW"},

		{name: "ADDconstForCarry", argLength: 1, reg: regInfo{inputs: []regMask{gp | sp | sb}, clobbers: tmp}, aux: "Int16", asm: "ADDC", typ: "Flags"},
		{name: "MaskIfNotCarry", argLength: 1, reg: crgp, asm: "ADDME", typ: "Int64"},

		{name: "SRADconst", argLength: 1, reg: gp11, asm: "SRAD", aux: "Int64"},
		{name: "SRAWconst", argLength: 1, reg: gp11, asm: "SRAW", aux: "Int64"},
		{name: "SRDconst", argLength: 1, reg: gp11, asm: "SRD", aux: "Int64"},
		{name: "SRWconst", argLength: 1, reg: gp11, asm: "SRW", aux: "Int64"},
		{name: "SLDconst", argLength: 1, reg: gp11, asm: "SLD", aux: "Int64"},
		{name: "SLWconst", argLength: 1, reg: gp11, asm: "SLW", aux: "Int64"},

		{name: "ROTLconst", argLength: 1, reg: gp11, asm: "ROTL", aux: "Int64"},
		{name: "ROTLWconst", argLength: 1, reg: gp11, asm: "ROTLW", aux: "Int64"},

		{name: "CNTLZD", argLength: 1, reg: gp11, asm: "CNTLZD", clobberFlags: true},
		{name: "CNTLZW", argLength: 1, reg: gp11, asm: "CNTLZW", clobberFlags: true},

		{name: "POPCNTD", argLength: 1, reg: gp11, asm: "POPCNTD"},
		{name: "POPCNTW", argLength: 1, reg: gp11, asm: "POPCNTW"},
		{name: "POPCNTB", argLength: 1, reg: gp11, asm: "POPCNTB"},

		{name: "FDIV", argLength: 2, reg: fp21, asm: "FDIV"},
		{name: "FDIVS", argLength: 2, reg: fp21, asm: "FDIVS"},

		{name: "DIVD", argLength: 2, reg: gp21, asm: "DIVD", typ: "Int64"},
		{name: "DIVW", argLength: 2, reg: gp21, asm: "DIVW", typ: "Int32"},
		{name: "DIVDU", argLength: 2, reg: gp21, asm: "DIVDU", typ: "Int64"},
		{name: "DIVWU", argLength: 2, reg: gp21, asm: "DIVWU", typ: "Int32"},

		{name: "FCTIDZ", argLength: 1, reg: fp11, asm: "FCTIDZ", typ: "Float64"},
		{name: "FCTIWZ", argLength: 1, reg: fp11, asm: "FCTIWZ", typ: "Float64"},
		{name: "FCFID", argLength: 1, reg: fp11, asm: "FCFID", typ: "Float64"},
		{name: "FCFIDS", argLength: 1, reg: fp11, asm: "FCFIDS", typ: "Float32"},
		{name: "FRSP", argLength: 1, reg: fp11, asm: "FRSP", typ: "Float64"},

		{name: "MFVSRD", argLength: 1, reg: fpgp, asm: "MFVSRD", typ: "Int64"},
		{name: "MTVSRD", argLength: 1, reg: gpfp, asm: "MTVSRD", typ: "Float64"},

		{name: "AND", argLength: 2, reg: gp21, asm: "AND", commutative: true},
		{name: "ANDN", argLength: 2, reg: gp21, asm: "ANDN"},
		{name: "OR", argLength: 2, reg: gp21, asm: "OR", commutative: true},
		{name: "ORN", argLength: 2, reg: gp21, asm: "ORN"},
		{name: "NOR", argLength: 2, reg: gp21, asm: "NOR", commutative: true},
		{name: "XOR", argLength: 2, reg: gp21, asm: "XOR", typ: "Int64", commutative: true},
		{name: "EQV", argLength: 2, reg: gp21, asm: "EQV", typ: "Int64", commutative: true},
		{name: "NEG", argLength: 1, reg: gp11, asm: "NEG"},
		{name: "FNEG", argLength: 1, reg: fp11, asm: "FNEG"},
		{name: "FSQRT", argLength: 1, reg: fp11, asm: "FSQRT"},
		{name: "FSQRTS", argLength: 1, reg: fp11, asm: "FSQRTS"},
		{name: "FFLOOR", argLength: 1, reg: fp11, asm: "FRIM"},
		{name: "FCEIL", argLength: 1, reg: fp11, asm: "FRIP"},
		{name: "FTRUNC", argLength: 1, reg: fp11, asm: "FRIZ"},
		{name: "FROUND", argLength: 1, reg: fp11, asm: "FRIN"},
		{name: "FABS", argLength: 1, reg: fp11, asm: "FABS"},
		{name: "FNABS", argLength: 1, reg: fp11, asm: "FNABS"},
		{name: "FCPSGN", argLength: 2, reg: fp21, asm: "FCPSGN"},

		{name: "ORconst", argLength: 1, reg: gp11, asm: "OR", aux: "Int64"},
		{name: "XORconst", argLength: 1, reg: gp11, asm: "XOR", aux: "Int64"},
		{name: "ANDconst", argLength: 1, reg: regInfo{inputs: []regMask{gp | sp | sb}, outputs: []regMask{gp}}, asm: "ANDCC", aux: "Int64", clobberFlags: true},
		{name: "ANDCCconst", argLength: 1, reg: regInfo{inputs: []regMask{gp | sp | sb}}, asm: "ANDCC", aux: "Int64", typ: "Flags"},

		{name: "MOVBreg", argLength: 1, reg: gp11, asm: "MOVB", typ: "Int64"},
		{name: "MOVBZreg", argLength: 1, reg: gp11, asm: "MOVBZ", typ: "Int64"},
		{name: "MOVHreg", argLength: 1, reg: gp11, asm: "MOVH", typ: "Int64"},
		{name: "MOVHZreg", argLength: 1, reg: gp11, asm: "MOVHZ", typ: "Int64"},
		{name: "MOVWreg", argLength: 1, reg: gp11, asm: "MOVW", typ: "Int64"},
		{name: "MOVWZreg", argLength: 1, reg: gp11, asm: "MOVWZ", typ: "Int64"},

		{name: "MOVBZload", argLength: 2, reg: gpload, asm: "MOVBZ", aux: "SymOff", typ: "UInt8", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHload", argLength: 2, reg: gpload, asm: "MOVH", aux: "SymOff", typ: "Int16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHZload", argLength: 2, reg: gpload, asm: "MOVHZ", aux: "SymOff", typ: "UInt16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWload", argLength: 2, reg: gpload, asm: "MOVW", aux: "SymOff", typ: "Int32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWZload", argLength: 2, reg: gpload, asm: "MOVWZ", aux: "SymOff", typ: "UInt32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVDload", argLength: 2, reg: gpload, asm: "MOVD", aux: "SymOff", typ: "Int64", faultOnNilArg0: true, symEffect: "Read"},

		{name: "MOVDBRload", argLength: 2, reg: gpload, asm: "MOVDBR", aux: "SymOff", typ: "Int64", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWBRload", argLength: 2, reg: gpload, asm: "MOVWBR", aux: "SymOff", typ: "Int32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHBRload", argLength: 2, reg: gpload, asm: "MOVHBR", aux: "SymOff", typ: "Int16", faultOnNilArg0: true, symEffect: "Read"},

		{name: "MOVDBRstore", argLength: 3, reg: gpstore, asm: "MOVDBR", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWBRstore", argLength: 3, reg: gpstore, asm: "MOVWBR", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHBRstore", argLength: 3, reg: gpstore, asm: "MOVHBR", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "FMOVDload", argLength: 2, reg: fpload, asm: "FMOVD", aux: "SymOff", typ: "Float64", faultOnNilArg0: true, symEffect: "Read"},
		{name: "FMOVSload", argLength: 2, reg: fpload, asm: "FMOVS", aux: "SymOff", typ: "Float32", faultOnNilArg0: true, symEffect: "Read"},

		{name: "MOVBstore", argLength: 3, reg: gpstore, asm: "MOVB", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHstore", argLength: 3, reg: gpstore, asm: "MOVH", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstore", argLength: 3, reg: gpstore, asm: "MOVW", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVDstore", argLength: 3, reg: gpstore, asm: "MOVD", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "FMOVDstore", argLength: 3, reg: fpstore, asm: "FMOVD", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "FMOVSstore", argLength: 3, reg: fpstore, asm: "FMOVS", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "MOVBstorezero", argLength: 2, reg: gpstorezero, asm: "MOVB", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHstorezero", argLength: 2, reg: gpstorezero, asm: "MOVH", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstorezero", argLength: 2, reg: gpstorezero, asm: "MOVW", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVDstorezero", argLength: 2, reg: gpstorezero, asm: "MOVD", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "MOVDaddr", argLength: 1, reg: regInfo{inputs: []regMask{sp | sb | gp}, outputs: []regMask{gp}}, aux: "SymOff", asm: "MOVD", rematerializeable: true, symEffect: "Addr"},

		{name: "MOVDconst", argLength: 0, reg: gp01, aux: "Int64", asm: "MOVD", typ: "Int64", rematerializeable: true},
		{name: "FMOVDconst", argLength: 0, reg: fp01, aux: "Float64", asm: "FMOVD", rematerializeable: true},
		{name: "FMOVSconst", argLength: 0, reg: fp01, aux: "Float32", asm: "FMOVS", rematerializeable: true},
		{name: "FCMPU", argLength: 2, reg: fp2cr, asm: "FCMPU", typ: "Flags"},

		{name: "CMP", argLength: 2, reg: gp2cr, asm: "CMP", typ: "Flags"},
		{name: "CMPU", argLength: 2, reg: gp2cr, asm: "CMPU", typ: "Flags"},
		{name: "CMPW", argLength: 2, reg: gp2cr, asm: "CMPW", typ: "Flags"},
		{name: "CMPWU", argLength: 2, reg: gp2cr, asm: "CMPWU", typ: "Flags"},
		{name: "CMPconst", argLength: 1, reg: gp1cr, asm: "CMP", aux: "Int64", typ: "Flags"},
		{name: "CMPUconst", argLength: 1, reg: gp1cr, asm: "CMPU", aux: "Int64", typ: "Flags"},
		{name: "CMPWconst", argLength: 1, reg: gp1cr, asm: "CMPW", aux: "Int32", typ: "Flags"},
		{name: "CMPWUconst", argLength: 1, reg: gp1cr, asm: "CMPWU", aux: "Int32", typ: "Flags"},

		{name: "Equal", argLength: 1, reg: crgp},
		{name: "NotEqual", argLength: 1, reg: crgp},
		{name: "LessThan", argLength: 1, reg: crgp},
		{name: "FLessThan", argLength: 1, reg: crgp},
		{name: "LessEqual", argLength: 1, reg: crgp},
		{name: "FLessEqual", argLength: 1, reg: crgp},
		{name: "GreaterThan", argLength: 1, reg: crgp},
		{name: "FGreaterThan", argLength: 1, reg: crgp},
		{name: "GreaterEqual", argLength: 1, reg: crgp},
		{name: "FGreaterEqual", argLength: 1, reg: crgp},

		{name: "LoweredGetClosurePtr", reg: regInfo{outputs: []regMask{ctxt}}, zeroWidth: true},

		{name: "LoweredGetCallerSP", reg: gp01, rematerializeable: true},

		{name: "LoweredGetCallerPC", reg: gp01, rematerializeable: true},

		{name: "LoweredNilCheck", argLength: 2, reg: regInfo{inputs: []regMask{gp | sp | sb}, clobbers: tmp}, clobberFlags: true, nilCheck: true, faultOnNilArg0: true},

		{name: "LoweredRound32F", argLength: 1, reg: fp11, resultInArg0: true, zeroWidth: true},
		{name: "LoweredRound64F", argLength: 1, reg: fp11, resultInArg0: true, zeroWidth: true},

		{name: "CALLstatic", argLength: 1, reg: regInfo{clobbers: callerSave}, aux: "SymOff", clobberFlags: true, call: true, symEffect: "None"},
		{name: "CALLclosure", argLength: 3, reg: regInfo{inputs: []regMask{callptr, ctxt, 0}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},
		{name: "CALLinter", argLength: 2, reg: regInfo{inputs: []regMask{callptr}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},

		{
			name:      "LoweredZero",
			aux:       "Int64",
			argLength: 2,
			reg: regInfo{
				inputs:   []regMask{buildReg("R3")},
				clobbers: buildReg("R3"),
			},
			clobberFlags:   true,
			typ:            "Mem",
			faultOnNilArg0: true,
		},

		{
			name:      "LoweredMove",
			aux:       "Int64",
			argLength: 3,
			reg: regInfo{
				inputs:   []regMask{buildReg("R3"), buildReg("R4")},
				clobbers: buildReg("R3 R4 R7 R8 R9 R10"),
			},
			clobberFlags:   true,
			typ:            "Mem",
			faultOnNilArg0: true,
			faultOnNilArg1: true,
		},

		{name: "LoweredAtomicStore32", argLength: 3, reg: gpstore, typ: "Mem", faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicStore64", argLength: 3, reg: gpstore, typ: "Mem", faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicLoad32", argLength: 2, reg: gpload, typ: "UInt32", clobberFlags: true, faultOnNilArg0: true},
		{name: "LoweredAtomicLoad64", argLength: 2, reg: gpload, typ: "Int64", clobberFlags: true, faultOnNilArg0: true},
		{name: "LoweredAtomicLoadPtr", argLength: 2, reg: gpload, typ: "Int64", clobberFlags: true, faultOnNilArg0: true},

		{name: "LoweredAtomicAdd32", argLength: 3, reg: gpxchg, resultNotInArgs: true, clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicAdd64", argLength: 3, reg: gpxchg, resultNotInArgs: true, clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicExchange32", argLength: 3, reg: gpxchg, resultNotInArgs: true, clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicExchange64", argLength: 3, reg: gpxchg, resultNotInArgs: true, clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicCas64", argLength: 4, reg: gpcas, resultNotInArgs: true, clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicCas32", argLength: 4, reg: gpcas, resultNotInArgs: true, clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicAnd8", argLength: 3, reg: gpstore, asm: "AND", faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicOr8", argLength: 3, reg: gpstore, asm: "OR", faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredWB", argLength: 3, reg: regInfo{inputs: []regMask{buildReg("R20"), buildReg("R21")}, clobbers: (callerSave &^ buildReg("R0 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R14 R15 R20 R21 g")) | buildReg("R31")}, clobberFlags: true, aux: "Sym", symEffect: "None"},

		{name: "InvertFlags", argLength: 1},

		{name: "FlagEQ"},
		{name: "FlagLT"},
		{name: "FlagGT"},
	}

	blocks := []blockData{
		{name: "EQ"},
		{name: "NE"},
		{name: "LT"},
		{name: "LE"},
		{name: "GT"},
		{name: "GE"},
		{name: "FLT"},
		{name: "FLE"},
		{name: "FGT"},
		{name: "FGE"},
	}

	archs = append(archs, arch{
		name:            "PPC64",
		pkg:             "github.com/dave/golib/src/cmd/internal/obj/ppc64",
		genfile:         "../../ppc64/ssa.go",
		ops:             ops,
		blocks:          blocks,
		regnames:        regNamesPPC64,
		gpregmask:       gp,
		fpregmask:       fp,
		framepointerreg: int8(num["SP"]),
		linkreg:         -1,
	})
}

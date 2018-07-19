// +build ignore

package main

import "strings"

// Note: registers not used in regalloc are not included in this list,
// so that regmask stays within int64
// Be careful when hand coding regmasks.
var regNamesMIPS64 = []string{
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

	"R24",
	"R25",

	"SP",
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

	"HI",
	"LO",

	"SB",
}

func init() {

	if len(regNamesMIPS64) > 64 {
		panic("too many registers")
	}
	num := map[string]int{}
	for i, name := range regNamesMIPS64 {
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
		gp         = buildReg("R1 R2 R3 R4 R5 R6 R7 R8 R9 R10 R11 R12 R13 R14 R15 R16 R17 R18 R19 R20 R21 R22 R24 R25 R31")
		gpg        = gp | buildReg("g")
		gpsp       = gp | buildReg("SP")
		gpspg      = gpg | buildReg("SP")
		gpspsbg    = gpspg | buildReg("SB")
		fp         = buildReg("F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15 F16 F17 F18 F19 F20 F21 F22 F23 F24 F25 F26 F27 F28 F29 F30 F31")
		lo         = buildReg("LO")
		hi         = buildReg("HI")
		callerSave = gp | fp | lo | hi | buildReg("g") // runtime.setg (and anything calling it) may clobber g
	)
	// Common regInfo
	var (
		gp01     = regInfo{inputs: nil, outputs: []regMask{gp}}
		gp11     = regInfo{inputs: []regMask{gpg}, outputs: []regMask{gp}}
		gp11sp   = regInfo{inputs: []regMask{gpspg}, outputs: []regMask{gp}}
		gp21     = regInfo{inputs: []regMask{gpg, gpg}, outputs: []regMask{gp}}
		gp2hilo  = regInfo{inputs: []regMask{gpg, gpg}, outputs: []regMask{hi, lo}}
		gpload   = regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{gp}}
		gpstore  = regInfo{inputs: []regMask{gpspsbg, gpg}}
		gpstore0 = regInfo{inputs: []regMask{gpspsbg}}
		gpxchg   = regInfo{inputs: []regMask{gpspsbg, gpg}, outputs: []regMask{gp}}
		gpcas    = regInfo{inputs: []regMask{gpspsbg, gpg, gpg}, outputs: []regMask{gp}}
		fp01     = regInfo{inputs: nil, outputs: []regMask{fp}}
		fp11     = regInfo{inputs: []regMask{fp}, outputs: []regMask{fp}}
		//fp1flags  = regInfo{inputs: []regMask{fp}}
		//fpgp      = regInfo{inputs: []regMask{fp}, outputs: []regMask{gp}}
		//gpfp      = regInfo{inputs: []regMask{gp}, outputs: []regMask{fp}}
		fp21      = regInfo{inputs: []regMask{fp, fp}, outputs: []regMask{fp}}
		fp2flags  = regInfo{inputs: []regMask{fp, fp}}
		fpload    = regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{fp}}
		fpstore   = regInfo{inputs: []regMask{gpspsbg, fp}}
		readflags = regInfo{inputs: nil, outputs: []regMask{gp}}
	)
	ops := []opData{

		{name: "ADDV", argLength: 2, reg: gp21, asm: "ADDVU", commutative: true},
		{name: "ADDVconst", argLength: 1, reg: gp11sp, asm: "ADDVU", aux: "Int64"},
		{name: "SUBV", argLength: 2, reg: gp21, asm: "SUBVU"},
		{name: "SUBVconst", argLength: 1, reg: gp11, asm: "SUBVU", aux: "Int64"},
		{name: "MULV", argLength: 2, reg: gp2hilo, asm: "MULV", commutative: true, typ: "(Int64,Int64)"},
		{name: "MULVU", argLength: 2, reg: gp2hilo, asm: "MULVU", commutative: true, typ: "(UInt64,UInt64)"},
		{name: "DIVV", argLength: 2, reg: gp2hilo, asm: "DIVV", typ: "(Int64,Int64)"},
		{name: "DIVVU", argLength: 2, reg: gp2hilo, asm: "DIVVU", typ: "(UInt64,UInt64)"},

		{name: "ADDF", argLength: 2, reg: fp21, asm: "ADDF", commutative: true},
		{name: "ADDD", argLength: 2, reg: fp21, asm: "ADDD", commutative: true},
		{name: "SUBF", argLength: 2, reg: fp21, asm: "SUBF"},
		{name: "SUBD", argLength: 2, reg: fp21, asm: "SUBD"},
		{name: "MULF", argLength: 2, reg: fp21, asm: "MULF", commutative: true},
		{name: "MULD", argLength: 2, reg: fp21, asm: "MULD", commutative: true},
		{name: "DIVF", argLength: 2, reg: fp21, asm: "DIVF"},
		{name: "DIVD", argLength: 2, reg: fp21, asm: "DIVD"},

		{name: "AND", argLength: 2, reg: gp21, asm: "AND", commutative: true},
		{name: "ANDconst", argLength: 1, reg: gp11, asm: "AND", aux: "Int64"},
		{name: "OR", argLength: 2, reg: gp21, asm: "OR", commutative: true},
		{name: "ORconst", argLength: 1, reg: gp11, asm: "OR", aux: "Int64"},
		{name: "XOR", argLength: 2, reg: gp21, asm: "XOR", commutative: true, typ: "UInt64"},
		{name: "XORconst", argLength: 1, reg: gp11, asm: "XOR", aux: "Int64", typ: "UInt64"},
		{name: "NOR", argLength: 2, reg: gp21, asm: "NOR", commutative: true},
		{name: "NORconst", argLength: 1, reg: gp11, asm: "NOR", aux: "Int64"},

		{name: "NEGV", argLength: 1, reg: gp11},
		{name: "NEGF", argLength: 1, reg: fp11, asm: "NEGF"},
		{name: "NEGD", argLength: 1, reg: fp11, asm: "NEGD"},
		{name: "SQRTD", argLength: 1, reg: fp11, asm: "SQRTD"},

		{name: "SLLV", argLength: 2, reg: gp21, asm: "SLLV"},
		{name: "SLLVconst", argLength: 1, reg: gp11, asm: "SLLV", aux: "Int64"},
		{name: "SRLV", argLength: 2, reg: gp21, asm: "SRLV"},
		{name: "SRLVconst", argLength: 1, reg: gp11, asm: "SRLV", aux: "Int64"},
		{name: "SRAV", argLength: 2, reg: gp21, asm: "SRAV"},
		{name: "SRAVconst", argLength: 1, reg: gp11, asm: "SRAV", aux: "Int64"},

		{name: "SGT", argLength: 2, reg: gp21, asm: "SGT", typ: "Bool"},
		{name: "SGTconst", argLength: 1, reg: gp11, asm: "SGT", aux: "Int64", typ: "Bool"},
		{name: "SGTU", argLength: 2, reg: gp21, asm: "SGTU", typ: "Bool"},
		{name: "SGTUconst", argLength: 1, reg: gp11, asm: "SGTU", aux: "Int64", typ: "Bool"},

		{name: "CMPEQF", argLength: 2, reg: fp2flags, asm: "CMPEQF", typ: "Flags"},
		{name: "CMPEQD", argLength: 2, reg: fp2flags, asm: "CMPEQD", typ: "Flags"},
		{name: "CMPGEF", argLength: 2, reg: fp2flags, asm: "CMPGEF", typ: "Flags"},
		{name: "CMPGED", argLength: 2, reg: fp2flags, asm: "CMPGED", typ: "Flags"},
		{name: "CMPGTF", argLength: 2, reg: fp2flags, asm: "CMPGTF", typ: "Flags"},
		{name: "CMPGTD", argLength: 2, reg: fp2flags, asm: "CMPGTD", typ: "Flags"},

		{name: "MOVVconst", argLength: 0, reg: gp01, aux: "Int64", asm: "MOVV", typ: "UInt64", rematerializeable: true},
		{name: "MOVFconst", argLength: 0, reg: fp01, aux: "Float64", asm: "MOVF", typ: "Float32", rematerializeable: true},
		{name: "MOVDconst", argLength: 0, reg: fp01, aux: "Float64", asm: "MOVD", typ: "Float64", rematerializeable: true},

		{name: "MOVVaddr", argLength: 1, reg: regInfo{inputs: []regMask{buildReg("SP") | buildReg("SB")}, outputs: []regMask{gp}}, aux: "SymOff", asm: "MOVV", rematerializeable: true, symEffect: "Addr"},

		{name: "MOVBload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVB", typ: "Int8", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVBUload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVBU", typ: "UInt8", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVH", typ: "Int16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHUload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVHU", typ: "UInt16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVW", typ: "Int32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWUload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVWU", typ: "UInt32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVVload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVV", typ: "UInt64", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVFload", argLength: 2, reg: fpload, aux: "SymOff", asm: "MOVF", typ: "Float32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVDload", argLength: 2, reg: fpload, aux: "SymOff", asm: "MOVD", typ: "Float64", faultOnNilArg0: true, symEffect: "Read"},

		{name: "MOVBstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "MOVB", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "MOVH", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "MOVW", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVVstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "MOVV", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVFstore", argLength: 3, reg: fpstore, aux: "SymOff", asm: "MOVF", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVDstore", argLength: 3, reg: fpstore, aux: "SymOff", asm: "MOVD", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "MOVBstorezero", argLength: 2, reg: gpstore0, aux: "SymOff", asm: "MOVB", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHstorezero", argLength: 2, reg: gpstore0, aux: "SymOff", asm: "MOVH", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstorezero", argLength: 2, reg: gpstore0, aux: "SymOff", asm: "MOVW", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVVstorezero", argLength: 2, reg: gpstore0, aux: "SymOff", asm: "MOVV", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "MOVBreg", argLength: 1, reg: gp11, asm: "MOVB"},
		{name: "MOVBUreg", argLength: 1, reg: gp11, asm: "MOVBU"},
		{name: "MOVHreg", argLength: 1, reg: gp11, asm: "MOVH"},
		{name: "MOVHUreg", argLength: 1, reg: gp11, asm: "MOVHU"},
		{name: "MOVWreg", argLength: 1, reg: gp11, asm: "MOVW"},
		{name: "MOVWUreg", argLength: 1, reg: gp11, asm: "MOVWU"},
		{name: "MOVVreg", argLength: 1, reg: gp11, asm: "MOVV"},

		{name: "MOVVnop", argLength: 1, reg: regInfo{inputs: []regMask{gp}, outputs: []regMask{gp}}, resultInArg0: true},

		{name: "MOVWF", argLength: 1, reg: fp11, asm: "MOVWF"},
		{name: "MOVWD", argLength: 1, reg: fp11, asm: "MOVWD"},
		{name: "MOVVF", argLength: 1, reg: fp11, asm: "MOVVF"},
		{name: "MOVVD", argLength: 1, reg: fp11, asm: "MOVVD"},
		{name: "TRUNCFW", argLength: 1, reg: fp11, asm: "TRUNCFW"},
		{name: "TRUNCDW", argLength: 1, reg: fp11, asm: "TRUNCDW"},
		{name: "TRUNCFV", argLength: 1, reg: fp11, asm: "TRUNCFV"},
		{name: "TRUNCDV", argLength: 1, reg: fp11, asm: "TRUNCDV"},
		{name: "MOVFD", argLength: 1, reg: fp11, asm: "MOVFD"},
		{name: "MOVDF", argLength: 1, reg: fp11, asm: "MOVDF"},

		{name: "CALLstatic", argLength: 1, reg: regInfo{clobbers: callerSave}, aux: "SymOff", clobberFlags: true, call: true, symEffect: "None"},
		{name: "CALLclosure", argLength: 3, reg: regInfo{inputs: []regMask{gpsp, buildReg("R22"), 0}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},
		{name: "CALLinter", argLength: 2, reg: regInfo{inputs: []regMask{gp}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},

		{
			name:      "DUFFZERO",
			aux:       "Int64",
			argLength: 2,
			reg: regInfo{
				inputs:   []regMask{gp},
				clobbers: buildReg("R1 R31"),
			},
			faultOnNilArg0: true,
		},

		{
			name:      "LoweredZero",
			aux:       "Int64",
			argLength: 3,
			reg: regInfo{
				inputs:   []regMask{buildReg("R1"), gp},
				clobbers: buildReg("R1"),
			},
			clobberFlags:   true,
			faultOnNilArg0: true,
		},

		{
			name:      "LoweredMove",
			aux:       "Int64",
			argLength: 4,
			reg: regInfo{
				inputs:   []regMask{buildReg("R2"), buildReg("R1"), gp},
				clobbers: buildReg("R1 R2"),
			},
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
		},

		{name: "LoweredAtomicLoad32", argLength: 2, reg: gpload, faultOnNilArg0: true},
		{name: "LoweredAtomicLoad64", argLength: 2, reg: gpload, faultOnNilArg0: true},

		{name: "LoweredAtomicStore32", argLength: 3, reg: gpstore, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicStore64", argLength: 3, reg: gpstore, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicStorezero32", argLength: 2, reg: gpstore0, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicStorezero64", argLength: 2, reg: gpstore0, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicExchange32", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicExchange64", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicAdd32", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicAdd64", argLength: 3, reg: gpxchg, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicAddconst32", argLength: 2, reg: regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{gp}}, aux: "Int32", resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicAddconst64", argLength: 2, reg: regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{gp}}, aux: "Int64", resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredAtomicCas32", argLength: 4, reg: gpcas, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},
		{name: "LoweredAtomicCas64", argLength: 4, reg: gpcas, resultNotInArgs: true, faultOnNilArg0: true, hasSideEffects: true},

		{name: "LoweredNilCheck", argLength: 2, reg: regInfo{inputs: []regMask{gpg}}, nilCheck: true, faultOnNilArg0: true},

		{name: "FPFlagTrue", argLength: 1, reg: readflags},
		{name: "FPFlagFalse", argLength: 1, reg: readflags},

		{name: "LoweredGetClosurePtr", reg: regInfo{outputs: []regMask{buildReg("R22")}}, zeroWidth: true},

		{name: "LoweredGetCallerSP", reg: gp01, rematerializeable: true},

		{name: "LoweredGetCallerPC", reg: gp01, rematerializeable: true},

		{name: "LoweredWB", argLength: 3, reg: regInfo{inputs: []regMask{buildReg("R20"), buildReg("R21")}, clobbers: (callerSave &^ gpg) | buildReg("R31")}, clobberFlags: true, aux: "Sym", symEffect: "None"},
	}

	blocks := []blockData{
		{name: "EQ"},
		{name: "NE"},
		{name: "LTZ"},
		{name: "LEZ"},
		{name: "GTZ"},
		{name: "GEZ"},
		{name: "FPT"},
		{name: "FPF"},
	}

	archs = append(archs, arch{
		name:            "MIPS64",
		pkg:             "github.com/dave/golib/src/cmd/internal/obj/mips",
		genfile:         "../../mips64/ssa.go",
		ops:             ops,
		blocks:          blocks,
		regnames:        regNamesMIPS64,
		gpregmask:       gp,
		fpregmask:       fp,
		specialregmask:  hi | lo,
		framepointerreg: -1,
		linkreg:         int8(num["R31"]),
	})
}

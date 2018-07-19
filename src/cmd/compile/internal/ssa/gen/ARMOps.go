// +build ignore

package main

import "strings"

var regNamesARM = []string{
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
	"g",
	"R11",
	"R12",
	"SP",
	"R14",
	"R15",

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

	"SB",
}

func init() {

	if len(regNamesARM) > 64 {
		panic("too many registers")
	}
	num := map[string]int{}
	for i, name := range regNamesARM {
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
		gp         = buildReg("R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R12 R14")
		gpg        = gp | buildReg("g")
		gpsp       = gp | buildReg("SP")
		gpspg      = gpg | buildReg("SP")
		gpspsbg    = gpspg | buildReg("SB")
		fp         = buildReg("F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15")
		callerSave = gp | fp | buildReg("g") // runtime.setg (and anything calling it) may clobber g
	)
	// Common regInfo
	var (
		gp01      = regInfo{inputs: nil, outputs: []regMask{gp}}
		gp11      = regInfo{inputs: []regMask{gpg}, outputs: []regMask{gp}}
		gp11carry = regInfo{inputs: []regMask{gpg}, outputs: []regMask{gp, 0}}
		gp11sp    = regInfo{inputs: []regMask{gpspg}, outputs: []regMask{gp}}
		gp1flags  = regInfo{inputs: []regMask{gpg}}
		gp1flags1 = regInfo{inputs: []regMask{gp}, outputs: []regMask{gp}}
		gp21      = regInfo{inputs: []regMask{gpg, gpg}, outputs: []regMask{gp}}
		gp21carry = regInfo{inputs: []regMask{gpg, gpg}, outputs: []regMask{gp, 0}}
		gp2flags  = regInfo{inputs: []regMask{gpg, gpg}}
		gp2flags1 = regInfo{inputs: []regMask{gp, gp}, outputs: []regMask{gp}}
		gp22      = regInfo{inputs: []regMask{gpg, gpg}, outputs: []regMask{gp, gp}}
		gp31      = regInfo{inputs: []regMask{gp, gp, gp}, outputs: []regMask{gp}}
		gp31carry = regInfo{inputs: []regMask{gp, gp, gp}, outputs: []regMask{gp, 0}}
		gp3flags  = regInfo{inputs: []regMask{gp, gp, gp}}
		gp3flags1 = regInfo{inputs: []regMask{gp, gp, gp}, outputs: []regMask{gp}}
		gpload    = regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{gp}}
		gpstore   = regInfo{inputs: []regMask{gpspsbg, gpg}}
		gp2load   = regInfo{inputs: []regMask{gpspsbg, gpg}, outputs: []regMask{gp}}
		gp2store  = regInfo{inputs: []regMask{gpspsbg, gpg, gpg}}
		fp01      = regInfo{inputs: nil, outputs: []regMask{fp}}
		fp11      = regInfo{inputs: []regMask{fp}, outputs: []regMask{fp}}
		fp1flags  = regInfo{inputs: []regMask{fp}}
		fpgp      = regInfo{inputs: []regMask{fp}, outputs: []regMask{gp}, clobbers: buildReg("F15")} // int-float conversion uses F15 as tmp
		gpfp      = regInfo{inputs: []regMask{gp}, outputs: []regMask{fp}, clobbers: buildReg("F15")}
		fp21      = regInfo{inputs: []regMask{fp, fp}, outputs: []regMask{fp}}
		fp31      = regInfo{inputs: []regMask{fp, fp, fp}, outputs: []regMask{fp}}
		fp2flags  = regInfo{inputs: []regMask{fp, fp}}
		fpload    = regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{fp}}
		fpstore   = regInfo{inputs: []regMask{gpspsbg, fp}}
		readflags = regInfo{inputs: nil, outputs: []regMask{gp}}
	)
	ops := []opData{

		{name: "ADD", argLength: 2, reg: gp21, asm: "ADD", commutative: true},
		{name: "ADDconst", argLength: 1, reg: gp11sp, asm: "ADD", aux: "Int32"},
		{name: "SUB", argLength: 2, reg: gp21, asm: "SUB"},
		{name: "SUBconst", argLength: 1, reg: gp11, asm: "SUB", aux: "Int32"},
		{name: "RSB", argLength: 2, reg: gp21, asm: "RSB"},
		{name: "RSBconst", argLength: 1, reg: gp11, asm: "RSB", aux: "Int32"},
		{name: "MUL", argLength: 2, reg: gp21, asm: "MUL", commutative: true},
		{name: "HMUL", argLength: 2, reg: gp21, asm: "MULL", commutative: true},
		{name: "HMULU", argLength: 2, reg: gp21, asm: "MULLU", commutative: true},

		{
			name:      "CALLudiv",
			argLength: 2,
			reg: regInfo{
				inputs:   []regMask{buildReg("R1"), buildReg("R0")},
				outputs:  []regMask{buildReg("R0"), buildReg("R1")},
				clobbers: buildReg("R2 R3 R14"),
			},
			clobberFlags: true,
			typ:          "(UInt32,UInt32)",
			call:         false,
		},

		{name: "ADDS", argLength: 2, reg: gp21carry, asm: "ADD", commutative: true},
		{name: "ADDSconst", argLength: 1, reg: gp11carry, asm: "ADD", aux: "Int32"},
		{name: "ADC", argLength: 3, reg: gp2flags1, asm: "ADC", commutative: true},
		{name: "ADCconst", argLength: 2, reg: gp1flags1, asm: "ADC", aux: "Int32"},
		{name: "SUBS", argLength: 2, reg: gp21carry, asm: "SUB"},
		{name: "SUBSconst", argLength: 1, reg: gp11carry, asm: "SUB", aux: "Int32"},
		{name: "RSBSconst", argLength: 1, reg: gp11carry, asm: "RSB", aux: "Int32"},
		{name: "SBC", argLength: 3, reg: gp2flags1, asm: "SBC"},
		{name: "SBCconst", argLength: 2, reg: gp1flags1, asm: "SBC", aux: "Int32"},
		{name: "RSCconst", argLength: 2, reg: gp1flags1, asm: "RSC", aux: "Int32"},

		{name: "MULLU", argLength: 2, reg: gp22, asm: "MULLU", commutative: true},
		{name: "MULA", argLength: 3, reg: gp31, asm: "MULA"},
		{name: "MULS", argLength: 3, reg: gp31, asm: "MULS"},

		{name: "ADDF", argLength: 2, reg: fp21, asm: "ADDF", commutative: true},
		{name: "ADDD", argLength: 2, reg: fp21, asm: "ADDD", commutative: true},
		{name: "SUBF", argLength: 2, reg: fp21, asm: "SUBF"},
		{name: "SUBD", argLength: 2, reg: fp21, asm: "SUBD"},
		{name: "MULF", argLength: 2, reg: fp21, asm: "MULF", commutative: true},
		{name: "MULD", argLength: 2, reg: fp21, asm: "MULD", commutative: true},
		{name: "NMULF", argLength: 2, reg: fp21, asm: "NMULF", commutative: true},
		{name: "NMULD", argLength: 2, reg: fp21, asm: "NMULD", commutative: true},
		{name: "DIVF", argLength: 2, reg: fp21, asm: "DIVF"},
		{name: "DIVD", argLength: 2, reg: fp21, asm: "DIVD"},

		{name: "MULAF", argLength: 3, reg: fp31, asm: "MULAF", resultInArg0: true},
		{name: "MULAD", argLength: 3, reg: fp31, asm: "MULAD", resultInArg0: true},
		{name: "MULSF", argLength: 3, reg: fp31, asm: "MULSF", resultInArg0: true},
		{name: "MULSD", argLength: 3, reg: fp31, asm: "MULSD", resultInArg0: true},

		{name: "AND", argLength: 2, reg: gp21, asm: "AND", commutative: true},
		{name: "ANDconst", argLength: 1, reg: gp11, asm: "AND", aux: "Int32"},
		{name: "OR", argLength: 2, reg: gp21, asm: "ORR", commutative: true},
		{name: "ORconst", argLength: 1, reg: gp11, asm: "ORR", aux: "Int32"},
		{name: "XOR", argLength: 2, reg: gp21, asm: "EOR", commutative: true},
		{name: "XORconst", argLength: 1, reg: gp11, asm: "EOR", aux: "Int32"},
		{name: "BIC", argLength: 2, reg: gp21, asm: "BIC"},
		{name: "BICconst", argLength: 1, reg: gp11, asm: "BIC", aux: "Int32"},

		{name: "BFX", argLength: 1, reg: gp11, asm: "BFX", aux: "Int32"},
		{name: "BFXU", argLength: 1, reg: gp11, asm: "BFXU", aux: "Int32"},

		{name: "MVN", argLength: 1, reg: gp11, asm: "MVN"},

		{name: "NEGF", argLength: 1, reg: fp11, asm: "NEGF"},
		{name: "NEGD", argLength: 1, reg: fp11, asm: "NEGD"},
		{name: "SQRTD", argLength: 1, reg: fp11, asm: "SQRTD"},

		{name: "CLZ", argLength: 1, reg: gp11, asm: "CLZ"},
		{name: "REV", argLength: 1, reg: gp11, asm: "REV"},
		{name: "RBIT", argLength: 1, reg: gp11, asm: "RBIT"},

		{name: "SLL", argLength: 2, reg: gp21, asm: "SLL"},
		{name: "SLLconst", argLength: 1, reg: gp11, asm: "SLL", aux: "Int32"},
		{name: "SRL", argLength: 2, reg: gp21, asm: "SRL"},
		{name: "SRLconst", argLength: 1, reg: gp11, asm: "SRL", aux: "Int32"},
		{name: "SRA", argLength: 2, reg: gp21, asm: "SRA"},
		{name: "SRAconst", argLength: 1, reg: gp11, asm: "SRA", aux: "Int32"},
		{name: "SRRconst", argLength: 1, reg: gp11, aux: "Int32"},

		{name: "ADDshiftLL", argLength: 2, reg: gp21, asm: "ADD", aux: "Int32"},
		{name: "ADDshiftRL", argLength: 2, reg: gp21, asm: "ADD", aux: "Int32"},
		{name: "ADDshiftRA", argLength: 2, reg: gp21, asm: "ADD", aux: "Int32"},
		{name: "SUBshiftLL", argLength: 2, reg: gp21, asm: "SUB", aux: "Int32"},
		{name: "SUBshiftRL", argLength: 2, reg: gp21, asm: "SUB", aux: "Int32"},
		{name: "SUBshiftRA", argLength: 2, reg: gp21, asm: "SUB", aux: "Int32"},
		{name: "RSBshiftLL", argLength: 2, reg: gp21, asm: "RSB", aux: "Int32"},
		{name: "RSBshiftRL", argLength: 2, reg: gp21, asm: "RSB", aux: "Int32"},
		{name: "RSBshiftRA", argLength: 2, reg: gp21, asm: "RSB", aux: "Int32"},
		{name: "ANDshiftLL", argLength: 2, reg: gp21, asm: "AND", aux: "Int32"},
		{name: "ANDshiftRL", argLength: 2, reg: gp21, asm: "AND", aux: "Int32"},
		{name: "ANDshiftRA", argLength: 2, reg: gp21, asm: "AND", aux: "Int32"},
		{name: "ORshiftLL", argLength: 2, reg: gp21, asm: "ORR", aux: "Int32"},
		{name: "ORshiftRL", argLength: 2, reg: gp21, asm: "ORR", aux: "Int32"},
		{name: "ORshiftRA", argLength: 2, reg: gp21, asm: "ORR", aux: "Int32"},
		{name: "XORshiftLL", argLength: 2, reg: gp21, asm: "EOR", aux: "Int32"},
		{name: "XORshiftRL", argLength: 2, reg: gp21, asm: "EOR", aux: "Int32"},
		{name: "XORshiftRA", argLength: 2, reg: gp21, asm: "EOR", aux: "Int32"},
		{name: "XORshiftRR", argLength: 2, reg: gp21, asm: "EOR", aux: "Int32"},
		{name: "BICshiftLL", argLength: 2, reg: gp21, asm: "BIC", aux: "Int32"},
		{name: "BICshiftRL", argLength: 2, reg: gp21, asm: "BIC", aux: "Int32"},
		{name: "BICshiftRA", argLength: 2, reg: gp21, asm: "BIC", aux: "Int32"},
		{name: "MVNshiftLL", argLength: 1, reg: gp11, asm: "MVN", aux: "Int32"},
		{name: "MVNshiftRL", argLength: 1, reg: gp11, asm: "MVN", aux: "Int32"},
		{name: "MVNshiftRA", argLength: 1, reg: gp11, asm: "MVN", aux: "Int32"},

		{name: "ADCshiftLL", argLength: 3, reg: gp2flags1, asm: "ADC", aux: "Int32"},
		{name: "ADCshiftRL", argLength: 3, reg: gp2flags1, asm: "ADC", aux: "Int32"},
		{name: "ADCshiftRA", argLength: 3, reg: gp2flags1, asm: "ADC", aux: "Int32"},
		{name: "SBCshiftLL", argLength: 3, reg: gp2flags1, asm: "SBC", aux: "Int32"},
		{name: "SBCshiftRL", argLength: 3, reg: gp2flags1, asm: "SBC", aux: "Int32"},
		{name: "SBCshiftRA", argLength: 3, reg: gp2flags1, asm: "SBC", aux: "Int32"},
		{name: "RSCshiftLL", argLength: 3, reg: gp2flags1, asm: "RSC", aux: "Int32"},
		{name: "RSCshiftRL", argLength: 3, reg: gp2flags1, asm: "RSC", aux: "Int32"},
		{name: "RSCshiftRA", argLength: 3, reg: gp2flags1, asm: "RSC", aux: "Int32"},

		{name: "ADDSshiftLL", argLength: 2, reg: gp21carry, asm: "ADD", aux: "Int32"},
		{name: "ADDSshiftRL", argLength: 2, reg: gp21carry, asm: "ADD", aux: "Int32"},
		{name: "ADDSshiftRA", argLength: 2, reg: gp21carry, asm: "ADD", aux: "Int32"},
		{name: "SUBSshiftLL", argLength: 2, reg: gp21carry, asm: "SUB", aux: "Int32"},
		{name: "SUBSshiftRL", argLength: 2, reg: gp21carry, asm: "SUB", aux: "Int32"},
		{name: "SUBSshiftRA", argLength: 2, reg: gp21carry, asm: "SUB", aux: "Int32"},
		{name: "RSBSshiftLL", argLength: 2, reg: gp21carry, asm: "RSB", aux: "Int32"},
		{name: "RSBSshiftRL", argLength: 2, reg: gp21carry, asm: "RSB", aux: "Int32"},
		{name: "RSBSshiftRA", argLength: 2, reg: gp21carry, asm: "RSB", aux: "Int32"},

		{name: "ADDshiftLLreg", argLength: 3, reg: gp31, asm: "ADD"},
		{name: "ADDshiftRLreg", argLength: 3, reg: gp31, asm: "ADD"},
		{name: "ADDshiftRAreg", argLength: 3, reg: gp31, asm: "ADD"},
		{name: "SUBshiftLLreg", argLength: 3, reg: gp31, asm: "SUB"},
		{name: "SUBshiftRLreg", argLength: 3, reg: gp31, asm: "SUB"},
		{name: "SUBshiftRAreg", argLength: 3, reg: gp31, asm: "SUB"},
		{name: "RSBshiftLLreg", argLength: 3, reg: gp31, asm: "RSB"},
		{name: "RSBshiftRLreg", argLength: 3, reg: gp31, asm: "RSB"},
		{name: "RSBshiftRAreg", argLength: 3, reg: gp31, asm: "RSB"},
		{name: "ANDshiftLLreg", argLength: 3, reg: gp31, asm: "AND"},
		{name: "ANDshiftRLreg", argLength: 3, reg: gp31, asm: "AND"},
		{name: "ANDshiftRAreg", argLength: 3, reg: gp31, asm: "AND"},
		{name: "ORshiftLLreg", argLength: 3, reg: gp31, asm: "ORR"},
		{name: "ORshiftRLreg", argLength: 3, reg: gp31, asm: "ORR"},
		{name: "ORshiftRAreg", argLength: 3, reg: gp31, asm: "ORR"},
		{name: "XORshiftLLreg", argLength: 3, reg: gp31, asm: "EOR"},
		{name: "XORshiftRLreg", argLength: 3, reg: gp31, asm: "EOR"},
		{name: "XORshiftRAreg", argLength: 3, reg: gp31, asm: "EOR"},
		{name: "BICshiftLLreg", argLength: 3, reg: gp31, asm: "BIC"},
		{name: "BICshiftRLreg", argLength: 3, reg: gp31, asm: "BIC"},
		{name: "BICshiftRAreg", argLength: 3, reg: gp31, asm: "BIC"},
		{name: "MVNshiftLLreg", argLength: 2, reg: gp21, asm: "MVN"},
		{name: "MVNshiftRLreg", argLength: 2, reg: gp21, asm: "MVN"},
		{name: "MVNshiftRAreg", argLength: 2, reg: gp21, asm: "MVN"},

		{name: "ADCshiftLLreg", argLength: 4, reg: gp3flags1, asm: "ADC"},
		{name: "ADCshiftRLreg", argLength: 4, reg: gp3flags1, asm: "ADC"},
		{name: "ADCshiftRAreg", argLength: 4, reg: gp3flags1, asm: "ADC"},
		{name: "SBCshiftLLreg", argLength: 4, reg: gp3flags1, asm: "SBC"},
		{name: "SBCshiftRLreg", argLength: 4, reg: gp3flags1, asm: "SBC"},
		{name: "SBCshiftRAreg", argLength: 4, reg: gp3flags1, asm: "SBC"},
		{name: "RSCshiftLLreg", argLength: 4, reg: gp3flags1, asm: "RSC"},
		{name: "RSCshiftRLreg", argLength: 4, reg: gp3flags1, asm: "RSC"},
		{name: "RSCshiftRAreg", argLength: 4, reg: gp3flags1, asm: "RSC"},

		{name: "ADDSshiftLLreg", argLength: 3, reg: gp31carry, asm: "ADD"},
		{name: "ADDSshiftRLreg", argLength: 3, reg: gp31carry, asm: "ADD"},
		{name: "ADDSshiftRAreg", argLength: 3, reg: gp31carry, asm: "ADD"},
		{name: "SUBSshiftLLreg", argLength: 3, reg: gp31carry, asm: "SUB"},
		{name: "SUBSshiftRLreg", argLength: 3, reg: gp31carry, asm: "SUB"},
		{name: "SUBSshiftRAreg", argLength: 3, reg: gp31carry, asm: "SUB"},
		{name: "RSBSshiftLLreg", argLength: 3, reg: gp31carry, asm: "RSB"},
		{name: "RSBSshiftRLreg", argLength: 3, reg: gp31carry, asm: "RSB"},
		{name: "RSBSshiftRAreg", argLength: 3, reg: gp31carry, asm: "RSB"},

		{name: "CMP", argLength: 2, reg: gp2flags, asm: "CMP", typ: "Flags"},
		{name: "CMPconst", argLength: 1, reg: gp1flags, asm: "CMP", aux: "Int32", typ: "Flags"},
		{name: "CMN", argLength: 2, reg: gp2flags, asm: "CMN", typ: "Flags", commutative: true},
		{name: "CMNconst", argLength: 1, reg: gp1flags, asm: "CMN", aux: "Int32", typ: "Flags"},
		{name: "TST", argLength: 2, reg: gp2flags, asm: "TST", typ: "Flags", commutative: true},
		{name: "TSTconst", argLength: 1, reg: gp1flags, asm: "TST", aux: "Int32", typ: "Flags"},
		{name: "TEQ", argLength: 2, reg: gp2flags, asm: "TEQ", typ: "Flags", commutative: true},
		{name: "TEQconst", argLength: 1, reg: gp1flags, asm: "TEQ", aux: "Int32", typ: "Flags"},
		{name: "CMPF", argLength: 2, reg: fp2flags, asm: "CMPF", typ: "Flags"},
		{name: "CMPD", argLength: 2, reg: fp2flags, asm: "CMPD", typ: "Flags"},

		{name: "CMPshiftLL", argLength: 2, reg: gp2flags, asm: "CMP", aux: "Int32", typ: "Flags"},
		{name: "CMPshiftRL", argLength: 2, reg: gp2flags, asm: "CMP", aux: "Int32", typ: "Flags"},
		{name: "CMPshiftRA", argLength: 2, reg: gp2flags, asm: "CMP", aux: "Int32", typ: "Flags"},
		{name: "CMNshiftLL", argLength: 2, reg: gp2flags, asm: "CMN", aux: "Int32", typ: "Flags"},
		{name: "CMNshiftRL", argLength: 2, reg: gp2flags, asm: "CMN", aux: "Int32", typ: "Flags"},
		{name: "CMNshiftRA", argLength: 2, reg: gp2flags, asm: "CMN", aux: "Int32", typ: "Flags"},
		{name: "TSTshiftLL", argLength: 2, reg: gp2flags, asm: "TST", aux: "Int32", typ: "Flags"},
		{name: "TSTshiftRL", argLength: 2, reg: gp2flags, asm: "TST", aux: "Int32", typ: "Flags"},
		{name: "TSTshiftRA", argLength: 2, reg: gp2flags, asm: "TST", aux: "Int32", typ: "Flags"},
		{name: "TEQshiftLL", argLength: 2, reg: gp2flags, asm: "TEQ", aux: "Int32", typ: "Flags"},
		{name: "TEQshiftRL", argLength: 2, reg: gp2flags, asm: "TEQ", aux: "Int32", typ: "Flags"},
		{name: "TEQshiftRA", argLength: 2, reg: gp2flags, asm: "TEQ", aux: "Int32", typ: "Flags"},

		{name: "CMPshiftLLreg", argLength: 3, reg: gp3flags, asm: "CMP", typ: "Flags"},
		{name: "CMPshiftRLreg", argLength: 3, reg: gp3flags, asm: "CMP", typ: "Flags"},
		{name: "CMPshiftRAreg", argLength: 3, reg: gp3flags, asm: "CMP", typ: "Flags"},
		{name: "CMNshiftLLreg", argLength: 3, reg: gp3flags, asm: "CMN", typ: "Flags"},
		{name: "CMNshiftRLreg", argLength: 3, reg: gp3flags, asm: "CMN", typ: "Flags"},
		{name: "CMNshiftRAreg", argLength: 3, reg: gp3flags, asm: "CMN", typ: "Flags"},
		{name: "TSTshiftLLreg", argLength: 3, reg: gp3flags, asm: "TST", typ: "Flags"},
		{name: "TSTshiftRLreg", argLength: 3, reg: gp3flags, asm: "TST", typ: "Flags"},
		{name: "TSTshiftRAreg", argLength: 3, reg: gp3flags, asm: "TST", typ: "Flags"},
		{name: "TEQshiftLLreg", argLength: 3, reg: gp3flags, asm: "TEQ", typ: "Flags"},
		{name: "TEQshiftRLreg", argLength: 3, reg: gp3flags, asm: "TEQ", typ: "Flags"},
		{name: "TEQshiftRAreg", argLength: 3, reg: gp3flags, asm: "TEQ", typ: "Flags"},

		{name: "CMPF0", argLength: 1, reg: fp1flags, asm: "CMPF", typ: "Flags"},
		{name: "CMPD0", argLength: 1, reg: fp1flags, asm: "CMPD", typ: "Flags"},

		{name: "MOVWconst", argLength: 0, reg: gp01, aux: "Int32", asm: "MOVW", typ: "UInt32", rematerializeable: true},
		{name: "MOVFconst", argLength: 0, reg: fp01, aux: "Float64", asm: "MOVF", typ: "Float32", rematerializeable: true},
		{name: "MOVDconst", argLength: 0, reg: fp01, aux: "Float64", asm: "MOVD", typ: "Float64", rematerializeable: true},

		{name: "MOVWaddr", argLength: 1, reg: regInfo{inputs: []regMask{buildReg("SP") | buildReg("SB")}, outputs: []regMask{gp}}, aux: "SymOff", asm: "MOVW", rematerializeable: true, symEffect: "Addr"},

		{name: "MOVBload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVB", typ: "Int8", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVBUload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVBU", typ: "UInt8", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVH", typ: "Int16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHUload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVHU", typ: "UInt16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWload", argLength: 2, reg: gpload, aux: "SymOff", asm: "MOVW", typ: "UInt32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVFload", argLength: 2, reg: fpload, aux: "SymOff", asm: "MOVF", typ: "Float32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVDload", argLength: 2, reg: fpload, aux: "SymOff", asm: "MOVD", typ: "Float64", faultOnNilArg0: true, symEffect: "Read"},

		{name: "MOVBstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "MOVB", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "MOVH", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstore", argLength: 3, reg: gpstore, aux: "SymOff", asm: "MOVW", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVFstore", argLength: 3, reg: fpstore, aux: "SymOff", asm: "MOVF", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVDstore", argLength: 3, reg: fpstore, aux: "SymOff", asm: "MOVD", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "MOVWloadidx", argLength: 3, reg: gp2load, asm: "MOVW"},
		{name: "MOVWloadshiftLL", argLength: 3, reg: gp2load, asm: "MOVW", aux: "Int32"},
		{name: "MOVWloadshiftRL", argLength: 3, reg: gp2load, asm: "MOVW", aux: "Int32"},
		{name: "MOVWloadshiftRA", argLength: 3, reg: gp2load, asm: "MOVW", aux: "Int32"},
		{name: "MOVBUloadidx", argLength: 3, reg: gp2load, asm: "MOVBU"},
		{name: "MOVBloadidx", argLength: 3, reg: gp2load, asm: "MOVB"},
		{name: "MOVHUloadidx", argLength: 3, reg: gp2load, asm: "MOVHU"},
		{name: "MOVHloadidx", argLength: 3, reg: gp2load, asm: "MOVH"},

		{name: "MOVWstoreidx", argLength: 4, reg: gp2store, asm: "MOVW"},
		{name: "MOVWstoreshiftLL", argLength: 4, reg: gp2store, asm: "MOVW", aux: "Int32"},
		{name: "MOVWstoreshiftRL", argLength: 4, reg: gp2store, asm: "MOVW", aux: "Int32"},
		{name: "MOVWstoreshiftRA", argLength: 4, reg: gp2store, asm: "MOVW", aux: "Int32"},
		{name: "MOVBstoreidx", argLength: 4, reg: gp2store, asm: "MOVB"},
		{name: "MOVHstoreidx", argLength: 4, reg: gp2store, asm: "MOVH"},

		{name: "MOVBreg", argLength: 1, reg: gp11, asm: "MOVBS"},
		{name: "MOVBUreg", argLength: 1, reg: gp11, asm: "MOVBU"},
		{name: "MOVHreg", argLength: 1, reg: gp11, asm: "MOVHS"},
		{name: "MOVHUreg", argLength: 1, reg: gp11, asm: "MOVHU"},
		{name: "MOVWreg", argLength: 1, reg: gp11, asm: "MOVW"},

		{name: "MOVWnop", argLength: 1, reg: regInfo{inputs: []regMask{gp}, outputs: []regMask{gp}}, resultInArg0: true},

		{name: "MOVWF", argLength: 1, reg: gpfp, asm: "MOVWF"},
		{name: "MOVWD", argLength: 1, reg: gpfp, asm: "MOVWD"},
		{name: "MOVWUF", argLength: 1, reg: gpfp, asm: "MOVWF"},
		{name: "MOVWUD", argLength: 1, reg: gpfp, asm: "MOVWD"},
		{name: "MOVFW", argLength: 1, reg: fpgp, asm: "MOVFW"},
		{name: "MOVDW", argLength: 1, reg: fpgp, asm: "MOVDW"},
		{name: "MOVFWU", argLength: 1, reg: fpgp, asm: "MOVFW"},
		{name: "MOVDWU", argLength: 1, reg: fpgp, asm: "MOVDW"},
		{name: "MOVFD", argLength: 1, reg: fp11, asm: "MOVFD"},
		{name: "MOVDF", argLength: 1, reg: fp11, asm: "MOVDF"},

		{name: "CMOVWHSconst", argLength: 2, reg: gp1flags1, asm: "MOVW", aux: "Int32", resultInArg0: true},
		{name: "CMOVWLSconst", argLength: 2, reg: gp1flags1, asm: "MOVW", aux: "Int32", resultInArg0: true},
		{name: "SRAcond", argLength: 3, reg: gp2flags1, asm: "SRA"},

		{name: "CALLstatic", argLength: 1, reg: regInfo{clobbers: callerSave}, aux: "SymOff", clobberFlags: true, call: true, symEffect: "None"},
		{name: "CALLclosure", argLength: 3, reg: regInfo{inputs: []regMask{gpsp, buildReg("R7"), 0}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},
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
			argLength: 3,
			reg: regInfo{
				inputs:   []regMask{buildReg("R1"), buildReg("R0")},
				clobbers: buildReg("R1 R14"),
			},
			faultOnNilArg0: true,
		},

		{
			name:      "DUFFCOPY",
			aux:       "Int64",
			argLength: 3,
			reg: regInfo{
				inputs:   []regMask{buildReg("R2"), buildReg("R1")},
				clobbers: buildReg("R0 R1 R2 R14"),
			},
			faultOnNilArg0: true,
			faultOnNilArg1: true,
		},

		{
			name:      "LoweredZero",
			aux:       "Int64",
			argLength: 4,
			reg: regInfo{
				inputs:   []regMask{buildReg("R1"), gp, gp},
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

		{name: "LoweredGetClosurePtr", reg: regInfo{outputs: []regMask{buildReg("R7")}}, zeroWidth: true},

		{name: "LoweredGetCallerSP", reg: gp01, rematerializeable: true},

		{name: "LoweredGetCallerPC", reg: gp01, rematerializeable: true},

		{name: "FlagEQ"},
		{name: "FlagLT_ULT"},
		{name: "FlagLT_UGT"},
		{name: "FlagGT_UGT"},
		{name: "FlagGT_ULT"},

		{name: "InvertFlags", argLength: 1},

		{name: "LoweredWB", argLength: 3, reg: regInfo{inputs: []regMask{buildReg("R2"), buildReg("R3")}, clobbers: (callerSave &^ gpg) | buildReg("R14")}, clobberFlags: true, aux: "Sym", symEffect: "None"},
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
	}

	archs = append(archs, arch{
		name:            "ARM",
		pkg:             "github.com/dave/golib/src/cmd/internal/obj/arm",
		genfile:         "../../arm/ssa.go",
		ops:             ops,
		blocks:          blocks,
		regnames:        regNamesARM,
		gpregmask:       gp,
		fpregmask:       fp,
		framepointerreg: -1,
		linkreg:         int8(num["R14"]),
	})
}

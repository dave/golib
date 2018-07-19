// +build ignore

package main

import "strings"

// copied from ../../s390x/reg.go
var regNamesS390X = []string{
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
	"g",
	"R14",
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

	"SB",
}

func init() {

	if len(regNamesS390X) > 64 {
		panic("too many registers")
	}
	num := map[string]int{}
	for i, name := range regNamesS390X {
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
		sp  = buildReg("SP")
		sb  = buildReg("SB")
		r0  = buildReg("R0")
		tmp = buildReg("R11") // R11 is used as a temporary in a small number of instructions.

		// R10 is reserved by the assembler.
		gp   = buildReg("R0 R1 R2 R3 R4 R5 R6 R7 R8 R9 R11 R12 R14")
		gpg  = gp | buildReg("g")
		gpsp = gp | sp

		// R0 is considered to contain the value 0 in address calculations.
		ptr     = gp &^ r0
		ptrsp   = ptr | sp
		ptrspsb = ptrsp | sb

		fp         = buildReg("F0 F1 F2 F3 F4 F5 F6 F7 F8 F9 F10 F11 F12 F13 F14 F15")
		callerSave = gp | fp | buildReg("g") // runtime.setg (and anything calling it) may clobber g
	)
	// Common slices of register masks
	var (
		gponly = []regMask{gp}
		fponly = []regMask{fp}
	)

	// Common regInfo
	var (
		gp01    = regInfo{inputs: []regMask{}, outputs: gponly}
		gp11    = regInfo{inputs: []regMask{gp}, outputs: gponly}
		gp11sp  = regInfo{inputs: []regMask{gpsp}, outputs: gponly}
		gp21    = regInfo{inputs: []regMask{gp, gp}, outputs: gponly}
		gp21sp  = regInfo{inputs: []regMask{gpsp, gp}, outputs: gponly}
		gp21tmp = regInfo{inputs: []regMask{gp &^ tmp, gp &^ tmp}, outputs: []regMask{gp &^ tmp}, clobbers: tmp}

		// R0 evaluates to 0 when used as the number of bits to shift
		// so we need to exclude it from that operand.
		sh21 = regInfo{inputs: []regMask{gp, ptr}, outputs: gponly}

		addr    = regInfo{inputs: []regMask{sp | sb}, outputs: gponly}
		addridx = regInfo{inputs: []regMask{sp | sb, ptrsp}, outputs: gponly}

		gp2flags  = regInfo{inputs: []regMask{gpsp, gpsp}}
		gp1flags  = regInfo{inputs: []regMask{gpsp}}
		gp2flags1 = regInfo{inputs: []regMask{gp, gp}, outputs: gponly}

		gpload       = regInfo{inputs: []regMask{ptrspsb, 0}, outputs: gponly}
		gploadidx    = regInfo{inputs: []regMask{ptrspsb, ptrsp, 0}, outputs: gponly}
		gpopload     = regInfo{inputs: []regMask{gp, ptrsp, 0}, outputs: gponly}
		gpstore      = regInfo{inputs: []regMask{ptrspsb, gpsp, 0}}
		gpstoreconst = regInfo{inputs: []regMask{ptrspsb, 0}}
		gpstoreidx   = regInfo{inputs: []regMask{ptrsp, ptrsp, gpsp, 0}}
		gpstorebr    = regInfo{inputs: []regMask{ptrsp, gpsp, 0}}
		gpstorelaa   = regInfo{inputs: []regMask{ptrspsb, gpsp, 0}, outputs: gponly}

		gpmvc = regInfo{inputs: []regMask{ptrsp, ptrsp, 0}}

		fp01        = regInfo{inputs: []regMask{}, outputs: fponly}
		fp21        = regInfo{inputs: []regMask{fp, fp}, outputs: fponly}
		fp31        = regInfo{inputs: []regMask{fp, fp, fp}, outputs: fponly}
		fp21clobber = regInfo{inputs: []regMask{fp, fp}, outputs: fponly}
		fpgp        = regInfo{inputs: fponly, outputs: gponly}
		gpfp        = regInfo{inputs: gponly, outputs: fponly}
		fp11        = regInfo{inputs: fponly, outputs: fponly}
		fp11clobber = regInfo{inputs: fponly, outputs: fponly}
		fp2flags    = regInfo{inputs: []regMask{fp, fp}}

		fpload    = regInfo{inputs: []regMask{ptrspsb, 0}, outputs: fponly}
		fploadidx = regInfo{inputs: []regMask{ptrsp, ptrsp, 0}, outputs: fponly}

		fpstore    = regInfo{inputs: []regMask{ptrspsb, fp, 0}}
		fpstoreidx = regInfo{inputs: []regMask{ptrsp, ptrsp, fp, 0}}

		// LoweredAtomicCas may overwrite arg1, so force it to R0 for now.
		cas = regInfo{inputs: []regMask{ptrsp, r0, gpsp, 0}, outputs: []regMask{gp, 0}, clobbers: r0}

		// LoweredAtomicExchange overwrites the output before executing
		// CS{,G}, so the output register must not be the same as the
		// input register. For now we just force the output register to
		// R0.
		exchange = regInfo{inputs: []regMask{ptrsp, gpsp &^ r0, 0}, outputs: []regMask{r0, 0}}
	)

	var S390Xops = []opData{

		{name: "FADDS", argLength: 2, reg: fp21clobber, asm: "FADDS", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "FADD", argLength: 2, reg: fp21clobber, asm: "FADD", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "FSUBS", argLength: 2, reg: fp21clobber, asm: "FSUBS", resultInArg0: true, clobberFlags: true},
		{name: "FSUB", argLength: 2, reg: fp21clobber, asm: "FSUB", resultInArg0: true, clobberFlags: true},
		{name: "FMULS", argLength: 2, reg: fp21, asm: "FMULS", commutative: true, resultInArg0: true},
		{name: "FMUL", argLength: 2, reg: fp21, asm: "FMUL", commutative: true, resultInArg0: true},
		{name: "FDIVS", argLength: 2, reg: fp21, asm: "FDIVS", resultInArg0: true},
		{name: "FDIV", argLength: 2, reg: fp21, asm: "FDIV", resultInArg0: true},
		{name: "FNEGS", argLength: 1, reg: fp11clobber, asm: "FNEGS", clobberFlags: true},
		{name: "FNEG", argLength: 1, reg: fp11clobber, asm: "FNEG", clobberFlags: true},
		{name: "FMADDS", argLength: 3, reg: fp31, asm: "FMADDS", resultInArg0: true},
		{name: "FMADD", argLength: 3, reg: fp31, asm: "FMADD", resultInArg0: true},
		{name: "FMSUBS", argLength: 3, reg: fp31, asm: "FMSUBS", resultInArg0: true},
		{name: "FMSUB", argLength: 3, reg: fp31, asm: "FMSUB", resultInArg0: true},
		{name: "LPDFR", argLength: 1, reg: fp11, asm: "LPDFR"},
		{name: "LNDFR", argLength: 1, reg: fp11, asm: "LNDFR"},
		{name: "CPSDR", argLength: 2, reg: fp21, asm: "CPSDR"},

		{name: "FIDBR", argLength: 1, reg: fp11, asm: "FIDBR", aux: "Int8"},

		{name: "FMOVSload", argLength: 2, reg: fpload, asm: "FMOVS", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},
		{name: "FMOVDload", argLength: 2, reg: fpload, asm: "FMOVD", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},
		{name: "FMOVSconst", reg: fp01, asm: "FMOVS", aux: "Float32", rematerializeable: true},
		{name: "FMOVDconst", reg: fp01, asm: "FMOVD", aux: "Float64", rematerializeable: true},
		{name: "FMOVSloadidx", argLength: 3, reg: fploadidx, asm: "FMOVS", aux: "SymOff", symEffect: "Read"},
		{name: "FMOVDloadidx", argLength: 3, reg: fploadidx, asm: "FMOVD", aux: "SymOff", symEffect: "Read"},

		{name: "FMOVSstore", argLength: 3, reg: fpstore, asm: "FMOVS", aux: "SymOff", faultOnNilArg0: true, symEffect: "Write"},
		{name: "FMOVDstore", argLength: 3, reg: fpstore, asm: "FMOVD", aux: "SymOff", faultOnNilArg0: true, symEffect: "Write"},
		{name: "FMOVSstoreidx", argLength: 4, reg: fpstoreidx, asm: "FMOVS", aux: "SymOff", symEffect: "Write"},
		{name: "FMOVDstoreidx", argLength: 4, reg: fpstoreidx, asm: "FMOVD", aux: "SymOff", symEffect: "Write"},

		{name: "ADD", argLength: 2, reg: gp21sp, asm: "ADD", commutative: true, clobberFlags: true},
		{name: "ADDW", argLength: 2, reg: gp21sp, asm: "ADDW", commutative: true, clobberFlags: true},
		{name: "ADDconst", argLength: 1, reg: gp11sp, asm: "ADD", aux: "Int32", typ: "UInt64", clobberFlags: true},
		{name: "ADDWconst", argLength: 1, reg: gp11sp, asm: "ADDW", aux: "Int32", clobberFlags: true},
		{name: "ADDload", argLength: 3, reg: gpopload, asm: "ADD", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "ADDWload", argLength: 3, reg: gpopload, asm: "ADDW", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},

		{name: "SUB", argLength: 2, reg: gp21, asm: "SUB", clobberFlags: true},
		{name: "SUBW", argLength: 2, reg: gp21, asm: "SUBW", clobberFlags: true},
		{name: "SUBconst", argLength: 1, reg: gp11, asm: "SUB", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "SUBWconst", argLength: 1, reg: gp11, asm: "SUBW", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "SUBload", argLength: 3, reg: gpopload, asm: "SUB", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "SUBWload", argLength: 3, reg: gpopload, asm: "SUBW", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},

		{name: "MULLD", argLength: 2, reg: gp21, asm: "MULLD", typ: "Int64", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "MULLW", argLength: 2, reg: gp21, asm: "MULLW", typ: "Int32", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "MULLDconst", argLength: 1, reg: gp11, asm: "MULLD", aux: "Int32", typ: "Int64", resultInArg0: true, clobberFlags: true},
		{name: "MULLWconst", argLength: 1, reg: gp11, asm: "MULLW", aux: "Int32", typ: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "MULLDload", argLength: 3, reg: gpopload, asm: "MULLD", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "MULLWload", argLength: 3, reg: gpopload, asm: "MULLW", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},

		{name: "MULHD", argLength: 2, reg: gp21tmp, asm: "MULHD", typ: "Int64", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "MULHDU", argLength: 2, reg: gp21tmp, asm: "MULHDU", typ: "Int64", commutative: true, resultInArg0: true, clobberFlags: true},

		{name: "DIVD", argLength: 2, reg: gp21tmp, asm: "DIVD", resultInArg0: true, clobberFlags: true},
		{name: "DIVW", argLength: 2, reg: gp21tmp, asm: "DIVW", resultInArg0: true, clobberFlags: true},
		{name: "DIVDU", argLength: 2, reg: gp21tmp, asm: "DIVDU", resultInArg0: true, clobberFlags: true},
		{name: "DIVWU", argLength: 2, reg: gp21tmp, asm: "DIVWU", resultInArg0: true, clobberFlags: true},

		{name: "MODD", argLength: 2, reg: gp21tmp, asm: "MODD", resultInArg0: true, clobberFlags: true},
		{name: "MODW", argLength: 2, reg: gp21tmp, asm: "MODW", resultInArg0: true, clobberFlags: true},

		{name: "MODDU", argLength: 2, reg: gp21tmp, asm: "MODDU", resultInArg0: true, clobberFlags: true},
		{name: "MODWU", argLength: 2, reg: gp21tmp, asm: "MODWU", resultInArg0: true, clobberFlags: true},

		{name: "AND", argLength: 2, reg: gp21, asm: "AND", commutative: true, clobberFlags: true},
		{name: "ANDW", argLength: 2, reg: gp21, asm: "ANDW", commutative: true, clobberFlags: true},
		{name: "ANDconst", argLength: 1, reg: gp11, asm: "AND", aux: "Int64", resultInArg0: true, clobberFlags: true},
		{name: "ANDWconst", argLength: 1, reg: gp11, asm: "ANDW", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "ANDload", argLength: 3, reg: gpopload, asm: "AND", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "ANDWload", argLength: 3, reg: gpopload, asm: "ANDW", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},

		{name: "OR", argLength: 2, reg: gp21, asm: "OR", commutative: true, clobberFlags: true},
		{name: "ORW", argLength: 2, reg: gp21, asm: "ORW", commutative: true, clobberFlags: true},
		{name: "ORconst", argLength: 1, reg: gp11, asm: "OR", aux: "Int64", resultInArg0: true, clobberFlags: true},
		{name: "ORWconst", argLength: 1, reg: gp11, asm: "ORW", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "ORload", argLength: 3, reg: gpopload, asm: "OR", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "ORWload", argLength: 3, reg: gpopload, asm: "ORW", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},

		{name: "XOR", argLength: 2, reg: gp21, asm: "XOR", commutative: true, clobberFlags: true},
		{name: "XORW", argLength: 2, reg: gp21, asm: "XORW", commutative: true, clobberFlags: true},
		{name: "XORconst", argLength: 1, reg: gp11, asm: "XOR", aux: "Int64", resultInArg0: true, clobberFlags: true},
		{name: "XORWconst", argLength: 1, reg: gp11, asm: "XORW", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "XORload", argLength: 3, reg: gpopload, asm: "XOR", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "XORWload", argLength: 3, reg: gpopload, asm: "XORW", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},

		{name: "CMP", argLength: 2, reg: gp2flags, asm: "CMP", typ: "Flags"},
		{name: "CMPW", argLength: 2, reg: gp2flags, asm: "CMPW", typ: "Flags"},

		{name: "CMPU", argLength: 2, reg: gp2flags, asm: "CMPU", typ: "Flags"},
		{name: "CMPWU", argLength: 2, reg: gp2flags, asm: "CMPWU", typ: "Flags"},

		{name: "CMPconst", argLength: 1, reg: gp1flags, asm: "CMP", typ: "Flags", aux: "Int32"},
		{name: "CMPWconst", argLength: 1, reg: gp1flags, asm: "CMPW", typ: "Flags", aux: "Int32"},
		{name: "CMPUconst", argLength: 1, reg: gp1flags, asm: "CMPU", typ: "Flags", aux: "Int32"},
		{name: "CMPWUconst", argLength: 1, reg: gp1flags, asm: "CMPWU", typ: "Flags", aux: "Int32"},

		{name: "FCMPS", argLength: 2, reg: fp2flags, asm: "CEBR", typ: "Flags"},
		{name: "FCMP", argLength: 2, reg: fp2flags, asm: "FCMPU", typ: "Flags"},

		{name: "SLD", argLength: 2, reg: sh21, asm: "SLD"},
		{name: "SLW", argLength: 2, reg: sh21, asm: "SLW"},
		{name: "SLDconst", argLength: 1, reg: gp11, asm: "SLD", aux: "Int8"},
		{name: "SLWconst", argLength: 1, reg: gp11, asm: "SLW", aux: "Int8"},

		{name: "SRD", argLength: 2, reg: sh21, asm: "SRD"},
		{name: "SRW", argLength: 2, reg: sh21, asm: "SRW"},
		{name: "SRDconst", argLength: 1, reg: gp11, asm: "SRD", aux: "Int8"},
		{name: "SRWconst", argLength: 1, reg: gp11, asm: "SRW", aux: "Int8"},

		{name: "SRAD", argLength: 2, reg: sh21, asm: "SRAD", clobberFlags: true},
		{name: "SRAW", argLength: 2, reg: sh21, asm: "SRAW", clobberFlags: true},
		{name: "SRADconst", argLength: 1, reg: gp11, asm: "SRAD", aux: "Int8", clobberFlags: true},
		{name: "SRAWconst", argLength: 1, reg: gp11, asm: "SRAW", aux: "Int8", clobberFlags: true},

		{name: "RLLGconst", argLength: 1, reg: gp11, asm: "RLLG", aux: "Int8"},
		{name: "RLLconst", argLength: 1, reg: gp11, asm: "RLL", aux: "Int8"},

		{name: "NEG", argLength: 1, reg: gp11, asm: "NEG", clobberFlags: true},
		{name: "NEGW", argLength: 1, reg: gp11, asm: "NEGW", clobberFlags: true},

		{name: "NOT", argLength: 1, reg: gp11, resultInArg0: true, clobberFlags: true},
		{name: "NOTW", argLength: 1, reg: gp11, resultInArg0: true, clobberFlags: true},

		{name: "FSQRT", argLength: 1, reg: fp11, asm: "FSQRT"},

		{name: "MOVDEQ", argLength: 3, reg: gp2flags1, resultInArg0: true, asm: "MOVDEQ"},
		{name: "MOVDNE", argLength: 3, reg: gp2flags1, resultInArg0: true, asm: "MOVDNE"},
		{name: "MOVDLT", argLength: 3, reg: gp2flags1, resultInArg0: true, asm: "MOVDLT"},
		{name: "MOVDLE", argLength: 3, reg: gp2flags1, resultInArg0: true, asm: "MOVDLE"},
		{name: "MOVDGT", argLength: 3, reg: gp2flags1, resultInArg0: true, asm: "MOVDGT"},
		{name: "MOVDGE", argLength: 3, reg: gp2flags1, resultInArg0: true, asm: "MOVDGE"},

		{name: "MOVDGTnoinv", argLength: 3, reg: gp2flags1, resultInArg0: true, asm: "MOVDGT"},
		{name: "MOVDGEnoinv", argLength: 3, reg: gp2flags1, resultInArg0: true, asm: "MOVDGE"},

		{name: "MOVBreg", argLength: 1, reg: gp11sp, asm: "MOVB", typ: "Int64"},
		{name: "MOVBZreg", argLength: 1, reg: gp11sp, asm: "MOVBZ", typ: "UInt64"},
		{name: "MOVHreg", argLength: 1, reg: gp11sp, asm: "MOVH", typ: "Int64"},
		{name: "MOVHZreg", argLength: 1, reg: gp11sp, asm: "MOVHZ", typ: "UInt64"},
		{name: "MOVWreg", argLength: 1, reg: gp11sp, asm: "MOVW", typ: "Int64"},
		{name: "MOVWZreg", argLength: 1, reg: gp11sp, asm: "MOVWZ", typ: "UInt64"},
		{name: "MOVDreg", argLength: 1, reg: gp11sp, asm: "MOVD"},

		{name: "MOVDnop", argLength: 1, reg: gp11, resultInArg0: true},

		{name: "MOVDconst", reg: gp01, asm: "MOVD", typ: "UInt64", aux: "Int64", rematerializeable: true},

		{name: "LDGR", argLength: 1, reg: gpfp, asm: "LDGR"},
		{name: "LGDR", argLength: 1, reg: fpgp, asm: "LGDR"},
		{name: "CFDBRA", argLength: 1, reg: fpgp, asm: "CFDBRA"},
		{name: "CGDBRA", argLength: 1, reg: fpgp, asm: "CGDBRA"},
		{name: "CFEBRA", argLength: 1, reg: fpgp, asm: "CFEBRA"},
		{name: "CGEBRA", argLength: 1, reg: fpgp, asm: "CGEBRA"},
		{name: "CEFBRA", argLength: 1, reg: gpfp, asm: "CEFBRA"},
		{name: "CDFBRA", argLength: 1, reg: gpfp, asm: "CDFBRA"},
		{name: "CEGBRA", argLength: 1, reg: gpfp, asm: "CEGBRA"},
		{name: "CDGBRA", argLength: 1, reg: gpfp, asm: "CDGBRA"},
		{name: "LEDBR", argLength: 1, reg: fp11, asm: "LEDBR"},
		{name: "LDEBR", argLength: 1, reg: fp11, asm: "LDEBR"},

		{name: "MOVDaddr", argLength: 1, reg: addr, aux: "SymOff", rematerializeable: true, symEffect: "Read"},
		{name: "MOVDaddridx", argLength: 2, reg: addridx, aux: "SymOff", symEffect: "Read"},

		{name: "MOVBZload", argLength: 2, reg: gpload, asm: "MOVBZ", aux: "SymOff", typ: "UInt8", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVBload", argLength: 2, reg: gpload, asm: "MOVB", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHZload", argLength: 2, reg: gpload, asm: "MOVHZ", aux: "SymOff", typ: "UInt16", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHload", argLength: 2, reg: gpload, asm: "MOVH", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWZload", argLength: 2, reg: gpload, asm: "MOVWZ", aux: "SymOff", typ: "UInt32", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWload", argLength: 2, reg: gpload, asm: "MOVW", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVDload", argLength: 2, reg: gpload, asm: "MOVD", aux: "SymOff", typ: "UInt64", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read"},

		{name: "MOVWBR", argLength: 1, reg: gp11, asm: "MOVWBR"},
		{name: "MOVDBR", argLength: 1, reg: gp11, asm: "MOVDBR"},

		{name: "MOVHBRload", argLength: 2, reg: gpload, asm: "MOVHBR", aux: "SymOff", typ: "UInt16", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWBRload", argLength: 2, reg: gpload, asm: "MOVWBR", aux: "SymOff", typ: "UInt32", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVDBRload", argLength: 2, reg: gpload, asm: "MOVDBR", aux: "SymOff", typ: "UInt64", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read"},

		{name: "MOVBstore", argLength: 3, reg: gpstore, asm: "MOVB", aux: "SymOff", typ: "Mem", clobberFlags: true, faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHstore", argLength: 3, reg: gpstore, asm: "MOVH", aux: "SymOff", typ: "Mem", clobberFlags: true, faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstore", argLength: 3, reg: gpstore, asm: "MOVW", aux: "SymOff", typ: "Mem", clobberFlags: true, faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVDstore", argLength: 3, reg: gpstore, asm: "MOVD", aux: "SymOff", typ: "Mem", clobberFlags: true, faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHBRstore", argLength: 3, reg: gpstorebr, asm: "MOVHBR", aux: "SymOff", typ: "Mem", clobberFlags: true, faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWBRstore", argLength: 3, reg: gpstorebr, asm: "MOVWBR", aux: "SymOff", typ: "Mem", clobberFlags: true, faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVDBRstore", argLength: 3, reg: gpstorebr, asm: "MOVDBR", aux: "SymOff", typ: "Mem", clobberFlags: true, faultOnNilArg0: true, symEffect: "Write"},

		{name: "MVC", argLength: 3, reg: gpmvc, asm: "MVC", aux: "SymValAndOff", typ: "Mem", clobberFlags: true, faultOnNilArg0: true, faultOnNilArg1: true, symEffect: "None"},

		{name: "MOVBZloadidx", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVBZ", aux: "SymOff", typ: "UInt8", clobberFlags: true, symEffect: "Read"},
		{name: "MOVBloadidx", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVB", aux: "SymOff", typ: "Int8", clobberFlags: true, symEffect: "Read"},
		{name: "MOVHZloadidx", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVHZ", aux: "SymOff", typ: "UInt16", clobberFlags: true, symEffect: "Read"},
		{name: "MOVHloadidx", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVH", aux: "SymOff", typ: "Int16", clobberFlags: true, symEffect: "Read"},
		{name: "MOVWZloadidx", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVWZ", aux: "SymOff", typ: "UInt32", clobberFlags: true, symEffect: "Read"},
		{name: "MOVWloadidx", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVW", aux: "SymOff", typ: "Int32", clobberFlags: true, symEffect: "Read"},
		{name: "MOVDloadidx", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVD", aux: "SymOff", typ: "UInt64", clobberFlags: true, symEffect: "Read"},
		{name: "MOVHBRloadidx", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVHBR", aux: "SymOff", typ: "Int16", clobberFlags: true, symEffect: "Read"},
		{name: "MOVWBRloadidx", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVWBR", aux: "SymOff", typ: "Int32", clobberFlags: true, symEffect: "Read"},
		{name: "MOVDBRloadidx", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVDBR", aux: "SymOff", typ: "Int64", clobberFlags: true, symEffect: "Read"},
		{name: "MOVBstoreidx", argLength: 4, reg: gpstoreidx, commutative: true, asm: "MOVB", aux: "SymOff", clobberFlags: true, symEffect: "Write"},
		{name: "MOVHstoreidx", argLength: 4, reg: gpstoreidx, commutative: true, asm: "MOVH", aux: "SymOff", clobberFlags: true, symEffect: "Write"},
		{name: "MOVWstoreidx", argLength: 4, reg: gpstoreidx, commutative: true, asm: "MOVW", aux: "SymOff", clobberFlags: true, symEffect: "Write"},
		{name: "MOVDstoreidx", argLength: 4, reg: gpstoreidx, commutative: true, asm: "MOVD", aux: "SymOff", clobberFlags: true, symEffect: "Write"},
		{name: "MOVHBRstoreidx", argLength: 4, reg: gpstoreidx, commutative: true, asm: "MOVHBR", aux: "SymOff", clobberFlags: true, symEffect: "Write"},
		{name: "MOVWBRstoreidx", argLength: 4, reg: gpstoreidx, commutative: true, asm: "MOVWBR", aux: "SymOff", clobberFlags: true, symEffect: "Write"},
		{name: "MOVDBRstoreidx", argLength: 4, reg: gpstoreidx, commutative: true, asm: "MOVDBR", aux: "SymOff", clobberFlags: true, symEffect: "Write"},

		{name: "MOVBstoreconst", argLength: 2, reg: gpstoreconst, asm: "MOVB", aux: "SymValAndOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHstoreconst", argLength: 2, reg: gpstoreconst, asm: "MOVH", aux: "SymValAndOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstoreconst", argLength: 2, reg: gpstoreconst, asm: "MOVW", aux: "SymValAndOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVDstoreconst", argLength: 2, reg: gpstoreconst, asm: "MOVD", aux: "SymValAndOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "CLEAR", argLength: 2, reg: regInfo{inputs: []regMask{ptr, 0}}, asm: "CLEAR", aux: "SymValAndOff", typ: "Mem", clobberFlags: true, faultOnNilArg0: true, symEffect: "Write"},

		{name: "CALLstatic", argLength: 1, reg: regInfo{clobbers: callerSave}, aux: "SymOff", clobberFlags: true, call: true, symEffect: "None"},
		{name: "CALLclosure", argLength: 3, reg: regInfo{inputs: []regMask{ptrsp, buildReg("R12"), 0}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},
		{name: "CALLinter", argLength: 2, reg: regInfo{inputs: []regMask{ptr}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},

		{name: "InvertFlags", argLength: 1},

		{name: "LoweredGetG", argLength: 1, reg: gp01},

		{name: "LoweredGetClosurePtr", reg: regInfo{outputs: []regMask{buildReg("R12")}}, zeroWidth: true},

		{name: "LoweredGetCallerSP", reg: gp01, rematerializeable: true},

		{name: "LoweredGetCallerPC", reg: gp01, rematerializeable: true},
		{name: "LoweredNilCheck", argLength: 2, reg: regInfo{inputs: []regMask{ptrsp}}, clobberFlags: true, nilCheck: true, faultOnNilArg0: true},

		{name: "LoweredRound32F", argLength: 1, reg: fp11, resultInArg0: true, zeroWidth: true},
		{name: "LoweredRound64F", argLength: 1, reg: fp11, resultInArg0: true, zeroWidth: true},

		{name: "LoweredWB", argLength: 3, reg: regInfo{inputs: []regMask{buildReg("R2"), buildReg("R3")}, clobbers: (callerSave &^ gpg) | buildReg("R14")}, clobberFlags: true, aux: "Sym", symEffect: "None"},

		{name: "FlagEQ"},
		{name: "FlagLT"},
		{name: "FlagGT"},

		{name: "MOVWZatomicload", argLength: 2, reg: gpload, asm: "MOVWZ", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVDatomicload", argLength: 2, reg: gpload, asm: "MOVD", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},

		{name: "MOVWatomicstore", argLength: 3, reg: gpstore, asm: "MOVW", aux: "SymOff", typ: "Mem", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "Write"},
		{name: "MOVDatomicstore", argLength: 3, reg: gpstore, asm: "MOVD", aux: "SymOff", typ: "Mem", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "Write"},

		{name: "LAA", argLength: 3, reg: gpstorelaa, asm: "LAA", typ: "(UInt32,Mem)", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "RdWr"},
		{name: "LAAG", argLength: 3, reg: gpstorelaa, asm: "LAAG", typ: "(UInt64,Mem)", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "RdWr"},
		{name: "AddTupleFirst32", argLength: 2},
		{name: "AddTupleFirst64", argLength: 2},

		{name: "LoweredAtomicCas32", argLength: 4, reg: cas, asm: "CS", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "RdWr"},
		{name: "LoweredAtomicCas64", argLength: 4, reg: cas, asm: "CSG", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "RdWr"},

		{name: "LoweredAtomicExchange32", argLength: 3, reg: exchange, asm: "CS", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "RdWr"},
		{name: "LoweredAtomicExchange64", argLength: 3, reg: exchange, asm: "CSG", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "RdWr"},

		{
			name:         "FLOGR",
			argLength:    1,
			reg:          regInfo{inputs: gponly, outputs: []regMask{buildReg("R0")}, clobbers: buildReg("R1")},
			asm:          "FLOGR",
			typ:          "UInt64",
			clobberFlags: true,
		},

		{
			name:           "STMG2",
			argLength:      4,
			reg:            regInfo{inputs: []regMask{ptrsp, buildReg("R1"), buildReg("R2"), 0}},
			aux:            "SymOff",
			typ:            "Mem",
			asm:            "STMG",
			faultOnNilArg0: true,
			symEffect:      "Write",
		},
		{
			name:           "STMG3",
			argLength:      5,
			reg:            regInfo{inputs: []regMask{ptrsp, buildReg("R1"), buildReg("R2"), buildReg("R3"), 0}},
			aux:            "SymOff",
			typ:            "Mem",
			asm:            "STMG",
			faultOnNilArg0: true,
			symEffect:      "Write",
		},
		{
			name:      "STMG4",
			argLength: 6,
			reg: regInfo{inputs: []regMask{
				ptrsp,
				buildReg("R1"),
				buildReg("R2"),
				buildReg("R3"),
				buildReg("R4"),
				0,
			}},
			aux:            "SymOff",
			typ:            "Mem",
			asm:            "STMG",
			faultOnNilArg0: true,
			symEffect:      "Write",
		},
		{
			name:           "STM2",
			argLength:      4,
			reg:            regInfo{inputs: []regMask{ptrsp, buildReg("R1"), buildReg("R2"), 0}},
			aux:            "SymOff",
			typ:            "Mem",
			asm:            "STMY",
			faultOnNilArg0: true,
			symEffect:      "Write",
		},
		{
			name:           "STM3",
			argLength:      5,
			reg:            regInfo{inputs: []regMask{ptrsp, buildReg("R1"), buildReg("R2"), buildReg("R3"), 0}},
			aux:            "SymOff",
			typ:            "Mem",
			asm:            "STMY",
			faultOnNilArg0: true,
			symEffect:      "Write",
		},
		{
			name:      "STM4",
			argLength: 6,
			reg: regInfo{inputs: []regMask{
				ptrsp,
				buildReg("R1"),
				buildReg("R2"),
				buildReg("R3"),
				buildReg("R4"),
				0,
			}},
			aux:            "SymOff",
			typ:            "Mem",
			asm:            "STMY",
			faultOnNilArg0: true,
			symEffect:      "Write",
		},

		{
			name:      "LoweredMove",
			aux:       "Int64",
			argLength: 4,
			reg: regInfo{
				inputs:   []regMask{buildReg("R1"), buildReg("R2"), gpsp},
				clobbers: buildReg("R1 R2"),
			},
			clobberFlags:   true,
			typ:            "Mem",
			faultOnNilArg0: true,
			faultOnNilArg1: true,
		},

		{
			name:      "LoweredZero",
			aux:       "Int64",
			argLength: 3,
			reg: regInfo{
				inputs:   []regMask{buildReg("R1"), gpsp},
				clobbers: buildReg("R1"),
			},
			clobberFlags:   true,
			typ:            "Mem",
			faultOnNilArg0: true,
		},
	}

	var S390Xblocks = []blockData{
		{name: "EQ"},
		{name: "NE"},
		{name: "LT"},
		{name: "LE"},
		{name: "GT"},
		{name: "GE"},
		{name: "GTF"},
		{name: "GEF"},
	}

	archs = append(archs, arch{
		name:            "S390X",
		pkg:             "github.com/dave/golib/src/cmd/internal/obj/s390x",
		genfile:         "../../s390x/ssa.go",
		ops:             S390Xops,
		blocks:          S390Xblocks,
		regnames:        regNamesS390X,
		gpregmask:       gp,
		fpregmask:       fp,
		framepointerreg: -1,
		linkreg:         int8(num["R14"]),
	})
}

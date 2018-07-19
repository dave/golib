// +build ignore

package main

import "strings"

// copied from ../../amd64/reg.go
var regNamesAMD64 = []string{
	"AX",
	"CX",
	"DX",
	"BX",
	"SP",
	"BP",
	"SI",
	"DI",
	"R8",
	"R9",
	"R10",
	"R11",
	"R12",
	"R13",
	"R14",
	"R15",
	"X0",
	"X1",
	"X2",
	"X3",
	"X4",
	"X5",
	"X6",
	"X7",
	"X8",
	"X9",
	"X10",
	"X11",
	"X12",
	"X13",
	"X14",
	"X15",

	"SB",
}

func init() {

	if len(regNamesAMD64) > 64 {
		panic("too many registers")
	}
	num := map[string]int{}
	for i, name := range regNamesAMD64 {
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
		ax         = buildReg("AX")
		cx         = buildReg("CX")
		dx         = buildReg("DX")
		gp         = buildReg("AX CX DX BX BP SI DI R8 R9 R10 R11 R12 R13 R14 R15")
		fp         = buildReg("X0 X1 X2 X3 X4 X5 X6 X7 X8 X9 X10 X11 X12 X13 X14 X15")
		gpsp       = gp | buildReg("SP")
		gpspsb     = gpsp | buildReg("SB")
		callerSave = gp | fp
	)
	// Common slices of register masks
	var (
		gponly = []regMask{gp}
		fponly = []regMask{fp}
	)

	// Common regInfo
	var (
		gp01      = regInfo{inputs: nil, outputs: gponly}
		gp11      = regInfo{inputs: []regMask{gp}, outputs: gponly}
		gp11sp    = regInfo{inputs: []regMask{gpsp}, outputs: gponly}
		gp11sb    = regInfo{inputs: []regMask{gpspsb}, outputs: gponly}
		gp21      = regInfo{inputs: []regMask{gp, gp}, outputs: gponly}
		gp21sp    = regInfo{inputs: []regMask{gpsp, gp}, outputs: gponly}
		gp21sb    = regInfo{inputs: []regMask{gpspsb, gpsp}, outputs: gponly}
		gp21shift = regInfo{inputs: []regMask{gp, cx}, outputs: []regMask{gp}}
		gp11div   = regInfo{inputs: []regMask{ax, gpsp &^ dx}, outputs: []regMask{ax, dx}}
		gp21hmul  = regInfo{inputs: []regMask{ax, gpsp}, outputs: []regMask{dx}, clobbers: ax}

		gp2flags     = regInfo{inputs: []regMask{gpsp, gpsp}}
		gp1flags     = regInfo{inputs: []regMask{gpsp}}
		gp0flagsLoad = regInfo{inputs: []regMask{gpspsb, 0}}
		gp1flagsLoad = regInfo{inputs: []regMask{gpspsb, gpsp, 0}}
		flagsgp      = regInfo{inputs: nil, outputs: gponly}

		gp11flags = regInfo{inputs: []regMask{gp}, outputs: []regMask{gp, 0}}

		readflags = regInfo{inputs: nil, outputs: gponly}
		flagsgpax = regInfo{inputs: nil, clobbers: ax, outputs: []regMask{gp &^ ax}}

		gpload    = regInfo{inputs: []regMask{gpspsb, 0}, outputs: gponly}
		gp21load  = regInfo{inputs: []regMask{gp, gpspsb, 0}, outputs: gponly}
		gploadidx = regInfo{inputs: []regMask{gpspsb, gpsp, 0}, outputs: gponly}
		gp21pax   = regInfo{inputs: []regMask{gp &^ ax, gp}, outputs: []regMask{gp &^ ax}, clobbers: ax}

		gpstore         = regInfo{inputs: []regMask{gpspsb, gpsp, 0}}
		gpstoreconst    = regInfo{inputs: []regMask{gpspsb, 0}}
		gpstoreidx      = regInfo{inputs: []regMask{gpspsb, gpsp, gpsp, 0}}
		gpstoreconstidx = regInfo{inputs: []regMask{gpspsb, gpsp, 0}}
		gpstorexchg     = regInfo{inputs: []regMask{gp, gpspsb, 0}, outputs: []regMask{gp}}
		cmpxchg         = regInfo{inputs: []regMask{gp, ax, gp, 0}, outputs: []regMask{gp, 0}, clobbers: ax}

		fp01     = regInfo{inputs: nil, outputs: fponly}
		fp21     = regInfo{inputs: []regMask{fp, fp}, outputs: fponly}
		fp21load = regInfo{inputs: []regMask{fp, gpspsb, 0}, outputs: fponly}
		fpgp     = regInfo{inputs: fponly, outputs: gponly}
		gpfp     = regInfo{inputs: gponly, outputs: fponly}
		fp11     = regInfo{inputs: fponly, outputs: fponly}
		fp2flags = regInfo{inputs: []regMask{fp, fp}}

		fpload    = regInfo{inputs: []regMask{gpspsb, 0}, outputs: fponly}
		fploadidx = regInfo{inputs: []regMask{gpspsb, gpsp, 0}, outputs: fponly}

		fpstore    = regInfo{inputs: []regMask{gpspsb, fp, 0}}
		fpstoreidx = regInfo{inputs: []regMask{gpspsb, gpsp, fp, 0}}
	)

	var AMD64ops = []opData{

		{name: "ADDSS", argLength: 2, reg: fp21, asm: "ADDSS", commutative: true, resultInArg0: true},
		{name: "ADDSD", argLength: 2, reg: fp21, asm: "ADDSD", commutative: true, resultInArg0: true},
		{name: "SUBSS", argLength: 2, reg: fp21, asm: "SUBSS", resultInArg0: true},
		{name: "SUBSD", argLength: 2, reg: fp21, asm: "SUBSD", resultInArg0: true},
		{name: "MULSS", argLength: 2, reg: fp21, asm: "MULSS", commutative: true, resultInArg0: true},
		{name: "MULSD", argLength: 2, reg: fp21, asm: "MULSD", commutative: true, resultInArg0: true},
		{name: "DIVSS", argLength: 2, reg: fp21, asm: "DIVSS", resultInArg0: true},
		{name: "DIVSD", argLength: 2, reg: fp21, asm: "DIVSD", resultInArg0: true},

		{name: "MOVSSload", argLength: 2, reg: fpload, asm: "MOVSS", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVSDload", argLength: 2, reg: fpload, asm: "MOVSD", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVSSconst", reg: fp01, asm: "MOVSS", aux: "Float32", rematerializeable: true},
		{name: "MOVSDconst", reg: fp01, asm: "MOVSD", aux: "Float64", rematerializeable: true},
		{name: "MOVSSloadidx1", argLength: 3, reg: fploadidx, asm: "MOVSS", aux: "SymOff", symEffect: "Read"},
		{name: "MOVSSloadidx4", argLength: 3, reg: fploadidx, asm: "MOVSS", aux: "SymOff", symEffect: "Read"},
		{name: "MOVSDloadidx1", argLength: 3, reg: fploadidx, asm: "MOVSD", aux: "SymOff", symEffect: "Read"},
		{name: "MOVSDloadidx8", argLength: 3, reg: fploadidx, asm: "MOVSD", aux: "SymOff", symEffect: "Read"},

		{name: "MOVSSstore", argLength: 3, reg: fpstore, asm: "MOVSS", aux: "SymOff", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVSDstore", argLength: 3, reg: fpstore, asm: "MOVSD", aux: "SymOff", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVSSstoreidx1", argLength: 4, reg: fpstoreidx, asm: "MOVSS", aux: "SymOff", symEffect: "Write"},
		{name: "MOVSSstoreidx4", argLength: 4, reg: fpstoreidx, asm: "MOVSS", aux: "SymOff", symEffect: "Write"},
		{name: "MOVSDstoreidx1", argLength: 4, reg: fpstoreidx, asm: "MOVSD", aux: "SymOff", symEffect: "Write"},
		{name: "MOVSDstoreidx8", argLength: 4, reg: fpstoreidx, asm: "MOVSD", aux: "SymOff", symEffect: "Write"},

		{name: "ADDSSload", argLength: 3, reg: fp21load, asm: "ADDSS", aux: "SymOff", resultInArg0: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "ADDSDload", argLength: 3, reg: fp21load, asm: "ADDSD", aux: "SymOff", resultInArg0: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "SUBSSload", argLength: 3, reg: fp21load, asm: "SUBSS", aux: "SymOff", resultInArg0: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "SUBSDload", argLength: 3, reg: fp21load, asm: "SUBSD", aux: "SymOff", resultInArg0: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "MULSSload", argLength: 3, reg: fp21load, asm: "MULSS", aux: "SymOff", resultInArg0: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "MULSDload", argLength: 3, reg: fp21load, asm: "MULSD", aux: "SymOff", resultInArg0: true, faultOnNilArg1: true, symEffect: "Read"},

		{name: "ADDQ", argLength: 2, reg: gp21sp, asm: "ADDQ", commutative: true, clobberFlags: true},
		{name: "ADDL", argLength: 2, reg: gp21sp, asm: "ADDL", commutative: true, clobberFlags: true},
		{name: "ADDQconst", argLength: 1, reg: gp11sp, asm: "ADDQ", aux: "Int32", typ: "UInt64", clobberFlags: true},
		{name: "ADDLconst", argLength: 1, reg: gp11sp, asm: "ADDL", aux: "Int32", clobberFlags: true},
		{name: "ADDQconstmodify", argLength: 2, reg: gpstoreconst, asm: "ADDQ", aux: "SymValAndOff", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read,Write"},
		{name: "ADDLconstmodify", argLength: 2, reg: gpstoreconst, asm: "ADDL", aux: "SymValAndOff", clobberFlags: true, faultOnNilArg0: true, symEffect: "Read,Write"},

		{name: "SUBQ", argLength: 2, reg: gp21, asm: "SUBQ", resultInArg0: true, clobberFlags: true},
		{name: "SUBL", argLength: 2, reg: gp21, asm: "SUBL", resultInArg0: true, clobberFlags: true},
		{name: "SUBQconst", argLength: 1, reg: gp11, asm: "SUBQ", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "SUBLconst", argLength: 1, reg: gp11, asm: "SUBL", aux: "Int32", resultInArg0: true, clobberFlags: true},

		{name: "MULQ", argLength: 2, reg: gp21, asm: "IMULQ", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "MULL", argLength: 2, reg: gp21, asm: "IMULL", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "MULQconst", argLength: 1, reg: gp11, asm: "IMUL3Q", aux: "Int32", clobberFlags: true},
		{name: "MULLconst", argLength: 1, reg: gp11, asm: "IMUL3L", aux: "Int32", clobberFlags: true},

		{name: "HMULQ", argLength: 2, reg: gp21hmul, commutative: true, asm: "IMULQ", clobberFlags: true},
		{name: "HMULL", argLength: 2, reg: gp21hmul, commutative: true, asm: "IMULL", clobberFlags: true},
		{name: "HMULQU", argLength: 2, reg: gp21hmul, commutative: true, asm: "MULQ", clobberFlags: true},
		{name: "HMULLU", argLength: 2, reg: gp21hmul, commutative: true, asm: "MULL", clobberFlags: true},

		{name: "AVGQU", argLength: 2, reg: gp21, commutative: true, resultInArg0: true, clobberFlags: true},

		{name: "DIVQ", argLength: 2, reg: gp11div, typ: "(Int64,Int64)", asm: "IDIVQ", clobberFlags: true},
		{name: "DIVL", argLength: 2, reg: gp11div, typ: "(Int32,Int32)", asm: "IDIVL", clobberFlags: true},
		{name: "DIVW", argLength: 2, reg: gp11div, typ: "(Int16,Int16)", asm: "IDIVW", clobberFlags: true},
		{name: "DIVQU", argLength: 2, reg: gp11div, typ: "(UInt64,UInt64)", asm: "DIVQ", clobberFlags: true},
		{name: "DIVLU", argLength: 2, reg: gp11div, typ: "(UInt32,UInt32)", asm: "DIVL", clobberFlags: true},
		{name: "DIVWU", argLength: 2, reg: gp11div, typ: "(UInt16,UInt16)", asm: "DIVW", clobberFlags: true},

		{name: "MULQU2", argLength: 2, reg: regInfo{inputs: []regMask{ax, gpsp}, outputs: []regMask{dx, ax}}, commutative: true, asm: "MULQ", clobberFlags: true},
		{name: "DIVQU2", argLength: 3, reg: regInfo{inputs: []regMask{dx, ax, gpsp}, outputs: []regMask{ax, dx}}, asm: "DIVQ", clobberFlags: true},

		{name: "ANDQ", argLength: 2, reg: gp21, asm: "ANDQ", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "ANDL", argLength: 2, reg: gp21, asm: "ANDL", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "ANDQconst", argLength: 1, reg: gp11, asm: "ANDQ", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "ANDLconst", argLength: 1, reg: gp11, asm: "ANDL", aux: "Int32", resultInArg0: true, clobberFlags: true},

		{name: "ORQ", argLength: 2, reg: gp21, asm: "ORQ", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "ORL", argLength: 2, reg: gp21, asm: "ORL", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "ORQconst", argLength: 1, reg: gp11, asm: "ORQ", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "ORLconst", argLength: 1, reg: gp11, asm: "ORL", aux: "Int32", resultInArg0: true, clobberFlags: true},

		{name: "XORQ", argLength: 2, reg: gp21, asm: "XORQ", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "XORL", argLength: 2, reg: gp21, asm: "XORL", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "XORQconst", argLength: 1, reg: gp11, asm: "XORQ", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "XORLconst", argLength: 1, reg: gp11, asm: "XORL", aux: "Int32", resultInArg0: true, clobberFlags: true},

		{name: "CMPQ", argLength: 2, reg: gp2flags, asm: "CMPQ", typ: "Flags"},
		{name: "CMPL", argLength: 2, reg: gp2flags, asm: "CMPL", typ: "Flags"},
		{name: "CMPW", argLength: 2, reg: gp2flags, asm: "CMPW", typ: "Flags"},
		{name: "CMPB", argLength: 2, reg: gp2flags, asm: "CMPB", typ: "Flags"},
		{name: "CMPQconst", argLength: 1, reg: gp1flags, asm: "CMPQ", typ: "Flags", aux: "Int32"},
		{name: "CMPLconst", argLength: 1, reg: gp1flags, asm: "CMPL", typ: "Flags", aux: "Int32"},
		{name: "CMPWconst", argLength: 1, reg: gp1flags, asm: "CMPW", typ: "Flags", aux: "Int16"},
		{name: "CMPBconst", argLength: 1, reg: gp1flags, asm: "CMPB", typ: "Flags", aux: "Int8"},

		{name: "CMPQload", argLength: 3, reg: gp1flagsLoad, asm: "CMPQ", aux: "SymOff", typ: "Flags", symEffect: "Read", faultOnNilArg0: true},
		{name: "CMPLload", argLength: 3, reg: gp1flagsLoad, asm: "CMPL", aux: "SymOff", typ: "Flags", symEffect: "Read", faultOnNilArg0: true},
		{name: "CMPWload", argLength: 3, reg: gp1flagsLoad, asm: "CMPW", aux: "SymOff", typ: "Flags", symEffect: "Read", faultOnNilArg0: true},
		{name: "CMPBload", argLength: 3, reg: gp1flagsLoad, asm: "CMPB", aux: "SymOff", typ: "Flags", symEffect: "Read", faultOnNilArg0: true},

		{name: "CMPQconstload", argLength: 2, reg: gp0flagsLoad, asm: "CMPQ", aux: "SymValAndOff", typ: "Flags", symEffect: "Read", faultOnNilArg0: true},
		{name: "CMPLconstload", argLength: 2, reg: gp0flagsLoad, asm: "CMPL", aux: "SymValAndOff", typ: "Flags", symEffect: "Read", faultOnNilArg0: true},
		{name: "CMPWconstload", argLength: 2, reg: gp0flagsLoad, asm: "CMPW", aux: "SymValAndOff", typ: "Flags", symEffect: "Read", faultOnNilArg0: true},
		{name: "CMPBconstload", argLength: 2, reg: gp0flagsLoad, asm: "CMPB", aux: "SymValAndOff", typ: "Flags", symEffect: "Read", faultOnNilArg0: true},

		{name: "UCOMISS", argLength: 2, reg: fp2flags, asm: "UCOMISS", typ: "Flags"},
		{name: "UCOMISD", argLength: 2, reg: fp2flags, asm: "UCOMISD", typ: "Flags"},

		{name: "BTL", argLength: 2, reg: gp2flags, asm: "BTL", typ: "Flags"},
		{name: "BTQ", argLength: 2, reg: gp2flags, asm: "BTQ", typ: "Flags"},
		{name: "BTCL", argLength: 2, reg: gp21, asm: "BTCL", resultInArg0: true, clobberFlags: true},
		{name: "BTCQ", argLength: 2, reg: gp21, asm: "BTCQ", resultInArg0: true, clobberFlags: true},
		{name: "BTRL", argLength: 2, reg: gp21, asm: "BTRL", resultInArg0: true, clobberFlags: true},
		{name: "BTRQ", argLength: 2, reg: gp21, asm: "BTRQ", resultInArg0: true, clobberFlags: true},
		{name: "BTSL", argLength: 2, reg: gp21, asm: "BTSL", resultInArg0: true, clobberFlags: true},
		{name: "BTSQ", argLength: 2, reg: gp21, asm: "BTSQ", resultInArg0: true, clobberFlags: true},
		{name: "BTLconst", argLength: 1, reg: gp1flags, asm: "BTL", typ: "Flags", aux: "Int8"},
		{name: "BTQconst", argLength: 1, reg: gp1flags, asm: "BTQ", typ: "Flags", aux: "Int8"},
		{name: "BTCLconst", argLength: 1, reg: gp11, asm: "BTCL", resultInArg0: true, clobberFlags: true, aux: "Int8"},
		{name: "BTCQconst", argLength: 1, reg: gp11, asm: "BTCQ", resultInArg0: true, clobberFlags: true, aux: "Int8"},
		{name: "BTRLconst", argLength: 1, reg: gp11, asm: "BTRL", resultInArg0: true, clobberFlags: true, aux: "Int8"},
		{name: "BTRQconst", argLength: 1, reg: gp11, asm: "BTRQ", resultInArg0: true, clobberFlags: true, aux: "Int8"},
		{name: "BTSLconst", argLength: 1, reg: gp11, asm: "BTSL", resultInArg0: true, clobberFlags: true, aux: "Int8"},
		{name: "BTSQconst", argLength: 1, reg: gp11, asm: "BTSQ", resultInArg0: true, clobberFlags: true, aux: "Int8"},

		{name: "TESTQ", argLength: 2, reg: gp2flags, commutative: true, asm: "TESTQ", typ: "Flags"},
		{name: "TESTL", argLength: 2, reg: gp2flags, commutative: true, asm: "TESTL", typ: "Flags"},
		{name: "TESTW", argLength: 2, reg: gp2flags, commutative: true, asm: "TESTW", typ: "Flags"},
		{name: "TESTB", argLength: 2, reg: gp2flags, commutative: true, asm: "TESTB", typ: "Flags"},
		{name: "TESTQconst", argLength: 1, reg: gp1flags, asm: "TESTQ", typ: "Flags", aux: "Int32"},
		{name: "TESTLconst", argLength: 1, reg: gp1flags, asm: "TESTL", typ: "Flags", aux: "Int32"},
		{name: "TESTWconst", argLength: 1, reg: gp1flags, asm: "TESTW", typ: "Flags", aux: "Int16"},
		{name: "TESTBconst", argLength: 1, reg: gp1flags, asm: "TESTB", typ: "Flags", aux: "Int8"},

		{name: "SHLQ", argLength: 2, reg: gp21shift, asm: "SHLQ", resultInArg0: true, clobberFlags: true},
		{name: "SHLL", argLength: 2, reg: gp21shift, asm: "SHLL", resultInArg0: true, clobberFlags: true},
		{name: "SHLQconst", argLength: 1, reg: gp11, asm: "SHLQ", aux: "Int8", resultInArg0: true, clobberFlags: true},
		{name: "SHLLconst", argLength: 1, reg: gp11, asm: "SHLL", aux: "Int8", resultInArg0: true, clobberFlags: true},

		{name: "SHRQ", argLength: 2, reg: gp21shift, asm: "SHRQ", resultInArg0: true, clobberFlags: true},
		{name: "SHRL", argLength: 2, reg: gp21shift, asm: "SHRL", resultInArg0: true, clobberFlags: true},
		{name: "SHRW", argLength: 2, reg: gp21shift, asm: "SHRW", resultInArg0: true, clobberFlags: true},
		{name: "SHRB", argLength: 2, reg: gp21shift, asm: "SHRB", resultInArg0: true, clobberFlags: true},
		{name: "SHRQconst", argLength: 1, reg: gp11, asm: "SHRQ", aux: "Int8", resultInArg0: true, clobberFlags: true},
		{name: "SHRLconst", argLength: 1, reg: gp11, asm: "SHRL", aux: "Int8", resultInArg0: true, clobberFlags: true},
		{name: "SHRWconst", argLength: 1, reg: gp11, asm: "SHRW", aux: "Int8", resultInArg0: true, clobberFlags: true},
		{name: "SHRBconst", argLength: 1, reg: gp11, asm: "SHRB", aux: "Int8", resultInArg0: true, clobberFlags: true},

		{name: "SARQ", argLength: 2, reg: gp21shift, asm: "SARQ", resultInArg0: true, clobberFlags: true},
		{name: "SARL", argLength: 2, reg: gp21shift, asm: "SARL", resultInArg0: true, clobberFlags: true},
		{name: "SARW", argLength: 2, reg: gp21shift, asm: "SARW", resultInArg0: true, clobberFlags: true},
		{name: "SARB", argLength: 2, reg: gp21shift, asm: "SARB", resultInArg0: true, clobberFlags: true},
		{name: "SARQconst", argLength: 1, reg: gp11, asm: "SARQ", aux: "Int8", resultInArg0: true, clobberFlags: true},
		{name: "SARLconst", argLength: 1, reg: gp11, asm: "SARL", aux: "Int8", resultInArg0: true, clobberFlags: true},
		{name: "SARWconst", argLength: 1, reg: gp11, asm: "SARW", aux: "Int8", resultInArg0: true, clobberFlags: true},
		{name: "SARBconst", argLength: 1, reg: gp11, asm: "SARB", aux: "Int8", resultInArg0: true, clobberFlags: true},

		{name: "ROLQ", argLength: 2, reg: gp21shift, asm: "ROLQ", resultInArg0: true, clobberFlags: true},
		{name: "ROLL", argLength: 2, reg: gp21shift, asm: "ROLL", resultInArg0: true, clobberFlags: true},
		{name: "ROLW", argLength: 2, reg: gp21shift, asm: "ROLW", resultInArg0: true, clobberFlags: true},
		{name: "ROLB", argLength: 2, reg: gp21shift, asm: "ROLB", resultInArg0: true, clobberFlags: true},
		{name: "RORQ", argLength: 2, reg: gp21shift, asm: "RORQ", resultInArg0: true, clobberFlags: true},
		{name: "RORL", argLength: 2, reg: gp21shift, asm: "RORL", resultInArg0: true, clobberFlags: true},
		{name: "RORW", argLength: 2, reg: gp21shift, asm: "RORW", resultInArg0: true, clobberFlags: true},
		{name: "RORB", argLength: 2, reg: gp21shift, asm: "RORB", resultInArg0: true, clobberFlags: true},
		{name: "ROLQconst", argLength: 1, reg: gp11, asm: "ROLQ", aux: "Int8", resultInArg0: true, clobberFlags: true},
		{name: "ROLLconst", argLength: 1, reg: gp11, asm: "ROLL", aux: "Int8", resultInArg0: true, clobberFlags: true},
		{name: "ROLWconst", argLength: 1, reg: gp11, asm: "ROLW", aux: "Int8", resultInArg0: true, clobberFlags: true},
		{name: "ROLBconst", argLength: 1, reg: gp11, asm: "ROLB", aux: "Int8", resultInArg0: true, clobberFlags: true},

		{name: "ADDLload", argLength: 3, reg: gp21load, asm: "ADDL", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "ADDQload", argLength: 3, reg: gp21load, asm: "ADDQ", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "SUBQload", argLength: 3, reg: gp21load, asm: "SUBQ", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "SUBLload", argLength: 3, reg: gp21load, asm: "SUBL", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "ANDLload", argLength: 3, reg: gp21load, asm: "ANDL", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "ANDQload", argLength: 3, reg: gp21load, asm: "ANDQ", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "ORQload", argLength: 3, reg: gp21load, asm: "ORQ", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "ORLload", argLength: 3, reg: gp21load, asm: "ORL", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "XORQload", argLength: 3, reg: gp21load, asm: "XORQ", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "XORLload", argLength: 3, reg: gp21load, asm: "XORL", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},

		{name: "NEGQ", argLength: 1, reg: gp11, asm: "NEGQ", resultInArg0: true, clobberFlags: true},
		{name: "NEGL", argLength: 1, reg: gp11, asm: "NEGL", resultInArg0: true, clobberFlags: true},

		{name: "NOTQ", argLength: 1, reg: gp11, asm: "NOTQ", resultInArg0: true, clobberFlags: true},
		{name: "NOTL", argLength: 1, reg: gp11, asm: "NOTL", resultInArg0: true, clobberFlags: true},

		{name: "BSFQ", argLength: 1, reg: gp11flags, asm: "BSFQ", typ: "(UInt64,Flags)"},
		{name: "BSFL", argLength: 1, reg: gp11, asm: "BSFL", typ: "UInt32", clobberFlags: true},
		{name: "BSRQ", argLength: 1, reg: gp11flags, asm: "BSRQ", typ: "(UInt64,Flags)"},
		{name: "BSRL", argLength: 1, reg: gp11, asm: "BSRL", typ: "UInt32", clobberFlags: true},

		{name: "CMOVQEQ", argLength: 3, reg: gp21, asm: "CMOVQEQ", resultInArg0: true},
		{name: "CMOVQNE", argLength: 3, reg: gp21, asm: "CMOVQNE", resultInArg0: true},
		{name: "CMOVQLT", argLength: 3, reg: gp21, asm: "CMOVQLT", resultInArg0: true},
		{name: "CMOVQGT", argLength: 3, reg: gp21, asm: "CMOVQGT", resultInArg0: true},
		{name: "CMOVQLE", argLength: 3, reg: gp21, asm: "CMOVQLE", resultInArg0: true},
		{name: "CMOVQGE", argLength: 3, reg: gp21, asm: "CMOVQGE", resultInArg0: true},
		{name: "CMOVQLS", argLength: 3, reg: gp21, asm: "CMOVQLS", resultInArg0: true},
		{name: "CMOVQHI", argLength: 3, reg: gp21, asm: "CMOVQHI", resultInArg0: true},
		{name: "CMOVQCC", argLength: 3, reg: gp21, asm: "CMOVQCC", resultInArg0: true},
		{name: "CMOVQCS", argLength: 3, reg: gp21, asm: "CMOVQCS", resultInArg0: true},

		{name: "CMOVLEQ", argLength: 3, reg: gp21, asm: "CMOVLEQ", resultInArg0: true},
		{name: "CMOVLNE", argLength: 3, reg: gp21, asm: "CMOVLNE", resultInArg0: true},
		{name: "CMOVLLT", argLength: 3, reg: gp21, asm: "CMOVLLT", resultInArg0: true},
		{name: "CMOVLGT", argLength: 3, reg: gp21, asm: "CMOVLGT", resultInArg0: true},
		{name: "CMOVLLE", argLength: 3, reg: gp21, asm: "CMOVLLE", resultInArg0: true},
		{name: "CMOVLGE", argLength: 3, reg: gp21, asm: "CMOVLGE", resultInArg0: true},
		{name: "CMOVLLS", argLength: 3, reg: gp21, asm: "CMOVLLS", resultInArg0: true},
		{name: "CMOVLHI", argLength: 3, reg: gp21, asm: "CMOVLHI", resultInArg0: true},
		{name: "CMOVLCC", argLength: 3, reg: gp21, asm: "CMOVLCC", resultInArg0: true},
		{name: "CMOVLCS", argLength: 3, reg: gp21, asm: "CMOVLCS", resultInArg0: true},

		{name: "CMOVWEQ", argLength: 3, reg: gp21, asm: "CMOVWEQ", resultInArg0: true},
		{name: "CMOVWNE", argLength: 3, reg: gp21, asm: "CMOVWNE", resultInArg0: true},
		{name: "CMOVWLT", argLength: 3, reg: gp21, asm: "CMOVWLT", resultInArg0: true},
		{name: "CMOVWGT", argLength: 3, reg: gp21, asm: "CMOVWGT", resultInArg0: true},
		{name: "CMOVWLE", argLength: 3, reg: gp21, asm: "CMOVWLE", resultInArg0: true},
		{name: "CMOVWGE", argLength: 3, reg: gp21, asm: "CMOVWGE", resultInArg0: true},
		{name: "CMOVWLS", argLength: 3, reg: gp21, asm: "CMOVWLS", resultInArg0: true},
		{name: "CMOVWHI", argLength: 3, reg: gp21, asm: "CMOVWHI", resultInArg0: true},
		{name: "CMOVWCC", argLength: 3, reg: gp21, asm: "CMOVWCC", resultInArg0: true},
		{name: "CMOVWCS", argLength: 3, reg: gp21, asm: "CMOVWCS", resultInArg0: true},

		{name: "CMOVQEQF", argLength: 3, reg: gp21pax, asm: "CMOVQNE", resultInArg0: true},
		{name: "CMOVQNEF", argLength: 3, reg: gp21, asm: "CMOVQNE", resultInArg0: true},
		{name: "CMOVQGTF", argLength: 3, reg: gp21, asm: "CMOVQHI", resultInArg0: true},
		{name: "CMOVQGEF", argLength: 3, reg: gp21, asm: "CMOVQCC", resultInArg0: true},
		{name: "CMOVLEQF", argLength: 3, reg: gp21, asm: "CMOVLNE", resultInArg0: true},
		{name: "CMOVLNEF", argLength: 3, reg: gp21, asm: "CMOVLNE", resultInArg0: true},
		{name: "CMOVLGTF", argLength: 3, reg: gp21, asm: "CMOVLHI", resultInArg0: true},
		{name: "CMOVLGEF", argLength: 3, reg: gp21, asm: "CMOVLCC", resultInArg0: true},
		{name: "CMOVWEQF", argLength: 3, reg: gp21, asm: "CMOVWNE", resultInArg0: true},
		{name: "CMOVWNEF", argLength: 3, reg: gp21, asm: "CMOVWNE", resultInArg0: true},
		{name: "CMOVWGTF", argLength: 3, reg: gp21, asm: "CMOVWHI", resultInArg0: true},
		{name: "CMOVWGEF", argLength: 3, reg: gp21, asm: "CMOVWCC", resultInArg0: true},

		{name: "BSWAPQ", argLength: 1, reg: gp11, asm: "BSWAPQ", resultInArg0: true, clobberFlags: true},
		{name: "BSWAPL", argLength: 1, reg: gp11, asm: "BSWAPL", resultInArg0: true, clobberFlags: true},

		{name: "POPCNTQ", argLength: 1, reg: gp11, asm: "POPCNTQ", clobberFlags: true},
		{name: "POPCNTL", argLength: 1, reg: gp11, asm: "POPCNTL", clobberFlags: true},

		{name: "SQRTSD", argLength: 1, reg: fp11, asm: "SQRTSD"},

		{name: "ROUNDSD", argLength: 1, reg: fp11, aux: "Int8", asm: "ROUNDSD"},

		{name: "SBBQcarrymask", argLength: 1, reg: flagsgp, asm: "SBBQ"},
		{name: "SBBLcarrymask", argLength: 1, reg: flagsgp, asm: "SBBL"},

		{name: "SETEQ", argLength: 1, reg: readflags, asm: "SETEQ"},
		{name: "SETNE", argLength: 1, reg: readflags, asm: "SETNE"},
		{name: "SETL", argLength: 1, reg: readflags, asm: "SETLT"},
		{name: "SETLE", argLength: 1, reg: readflags, asm: "SETLE"},
		{name: "SETG", argLength: 1, reg: readflags, asm: "SETGT"},
		{name: "SETGE", argLength: 1, reg: readflags, asm: "SETGE"},
		{name: "SETB", argLength: 1, reg: readflags, asm: "SETCS"},
		{name: "SETBE", argLength: 1, reg: readflags, asm: "SETLS"},
		{name: "SETA", argLength: 1, reg: readflags, asm: "SETHI"},
		{name: "SETAE", argLength: 1, reg: readflags, asm: "SETCC"},

		{name: "SETEQstore", argLength: 3, reg: gpstoreconst, asm: "SETEQ", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "SETNEstore", argLength: 3, reg: gpstoreconst, asm: "SETNE", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "SETLstore", argLength: 3, reg: gpstoreconst, asm: "SETLT", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "SETLEstore", argLength: 3, reg: gpstoreconst, asm: "SETLE", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "SETGstore", argLength: 3, reg: gpstoreconst, asm: "SETGT", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "SETGEstore", argLength: 3, reg: gpstoreconst, asm: "SETGE", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "SETBstore", argLength: 3, reg: gpstoreconst, asm: "SETCS", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "SETBEstore", argLength: 3, reg: gpstoreconst, asm: "SETLS", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "SETAstore", argLength: 3, reg: gpstoreconst, asm: "SETHI", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "SETAEstore", argLength: 3, reg: gpstoreconst, asm: "SETCC", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "SETEQF", argLength: 1, reg: flagsgpax, asm: "SETEQ", clobberFlags: true},
		{name: "SETNEF", argLength: 1, reg: flagsgpax, asm: "SETNE", clobberFlags: true},
		{name: "SETORD", argLength: 1, reg: flagsgp, asm: "SETPC"},
		{name: "SETNAN", argLength: 1, reg: flagsgp, asm: "SETPS"},

		{name: "SETGF", argLength: 1, reg: flagsgp, asm: "SETHI"},
		{name: "SETGEF", argLength: 1, reg: flagsgp, asm: "SETCC"},

		{name: "MOVBQSX", argLength: 1, reg: gp11, asm: "MOVBQSX"},
		{name: "MOVBQZX", argLength: 1, reg: gp11, asm: "MOVBLZX"},
		{name: "MOVWQSX", argLength: 1, reg: gp11, asm: "MOVWQSX"},
		{name: "MOVWQZX", argLength: 1, reg: gp11, asm: "MOVWLZX"},
		{name: "MOVLQSX", argLength: 1, reg: gp11, asm: "MOVLQSX"},
		{name: "MOVLQZX", argLength: 1, reg: gp11, asm: "MOVL"},

		{name: "MOVLconst", reg: gp01, asm: "MOVL", typ: "UInt32", aux: "Int32", rematerializeable: true},
		{name: "MOVQconst", reg: gp01, asm: "MOVQ", typ: "UInt64", aux: "Int64", rematerializeable: true},

		{name: "CVTTSD2SL", argLength: 1, reg: fpgp, asm: "CVTTSD2SL"},
		{name: "CVTTSD2SQ", argLength: 1, reg: fpgp, asm: "CVTTSD2SQ"},
		{name: "CVTTSS2SL", argLength: 1, reg: fpgp, asm: "CVTTSS2SL"},
		{name: "CVTTSS2SQ", argLength: 1, reg: fpgp, asm: "CVTTSS2SQ"},
		{name: "CVTSL2SS", argLength: 1, reg: gpfp, asm: "CVTSL2SS"},
		{name: "CVTSL2SD", argLength: 1, reg: gpfp, asm: "CVTSL2SD"},
		{name: "CVTSQ2SS", argLength: 1, reg: gpfp, asm: "CVTSQ2SS"},
		{name: "CVTSQ2SD", argLength: 1, reg: gpfp, asm: "CVTSQ2SD"},
		{name: "CVTSD2SS", argLength: 1, reg: fp11, asm: "CVTSD2SS"},
		{name: "CVTSS2SD", argLength: 1, reg: fp11, asm: "CVTSS2SD"},

		{name: "MOVQi2f", argLength: 1, reg: gpfp, typ: "Float64"},
		{name: "MOVQf2i", argLength: 1, reg: fpgp, typ: "UInt64"},
		{name: "MOVLi2f", argLength: 1, reg: gpfp, typ: "Float32"},
		{name: "MOVLf2i", argLength: 1, reg: fpgp, typ: "UInt32"},

		{name: "PXOR", argLength: 2, reg: fp21, asm: "PXOR", commutative: true, resultInArg0: true},

		{name: "LEAQ", argLength: 1, reg: gp11sb, asm: "LEAQ", aux: "SymOff", rematerializeable: true, symEffect: "Addr"},
		{name: "LEAL", argLength: 1, reg: gp11sb, asm: "LEAL", aux: "SymOff", rematerializeable: true, symEffect: "Addr"},
		{name: "LEAW", argLength: 1, reg: gp11sb, asm: "LEAW", aux: "SymOff", rematerializeable: true, symEffect: "Addr"},
		{name: "LEAQ1", argLength: 2, reg: gp21sb, asm: "LEAQ", commutative: true, aux: "SymOff", symEffect: "Addr"},
		{name: "LEAL1", argLength: 2, reg: gp21sb, asm: "LEAL", commutative: true, aux: "SymOff", symEffect: "Addr"},
		{name: "LEAW1", argLength: 2, reg: gp21sb, asm: "LEAW", commutative: true, aux: "SymOff", symEffect: "Addr"},
		{name: "LEAQ2", argLength: 2, reg: gp21sb, asm: "LEAQ", aux: "SymOff", symEffect: "Addr"},
		{name: "LEAL2", argLength: 2, reg: gp21sb, asm: "LEAL", aux: "SymOff", symEffect: "Addr"},
		{name: "LEAW2", argLength: 2, reg: gp21sb, asm: "LEAW", aux: "SymOff", symEffect: "Addr"},
		{name: "LEAQ4", argLength: 2, reg: gp21sb, asm: "LEAQ", aux: "SymOff", symEffect: "Addr"},
		{name: "LEAL4", argLength: 2, reg: gp21sb, asm: "LEAL", aux: "SymOff", symEffect: "Addr"},
		{name: "LEAW4", argLength: 2, reg: gp21sb, asm: "LEAW", aux: "SymOff", symEffect: "Addr"},
		{name: "LEAQ8", argLength: 2, reg: gp21sb, asm: "LEAQ", aux: "SymOff", symEffect: "Addr"},
		{name: "LEAL8", argLength: 2, reg: gp21sb, asm: "LEAL", aux: "SymOff", symEffect: "Addr"},
		{name: "LEAW8", argLength: 2, reg: gp21sb, asm: "LEAW", aux: "SymOff", symEffect: "Addr"},

		{name: "MOVBload", argLength: 2, reg: gpload, asm: "MOVBLZX", aux: "SymOff", typ: "UInt8", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVBQSXload", argLength: 2, reg: gpload, asm: "MOVBQSX", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWload", argLength: 2, reg: gpload, asm: "MOVWLZX", aux: "SymOff", typ: "UInt16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWQSXload", argLength: 2, reg: gpload, asm: "MOVWQSX", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVLload", argLength: 2, reg: gpload, asm: "MOVL", aux: "SymOff", typ: "UInt32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVLQSXload", argLength: 2, reg: gpload, asm: "MOVLQSX", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVQload", argLength: 2, reg: gpload, asm: "MOVQ", aux: "SymOff", typ: "UInt64", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVBstore", argLength: 3, reg: gpstore, asm: "MOVB", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstore", argLength: 3, reg: gpstore, asm: "MOVW", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVLstore", argLength: 3, reg: gpstore, asm: "MOVL", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVQstore", argLength: 3, reg: gpstore, asm: "MOVQ", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVOload", argLength: 2, reg: fpload, asm: "MOVUPS", aux: "SymOff", typ: "Int128", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVOstore", argLength: 3, reg: fpstore, asm: "MOVUPS", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "MOVBloadidx1", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVBLZX", aux: "SymOff", typ: "UInt8", symEffect: "Read"},
		{name: "MOVWloadidx1", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVWLZX", aux: "SymOff", typ: "UInt16", symEffect: "Read"},
		{name: "MOVWloadidx2", argLength: 3, reg: gploadidx, asm: "MOVWLZX", aux: "SymOff", typ: "UInt16", symEffect: "Read"},
		{name: "MOVLloadidx1", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVL", aux: "SymOff", typ: "UInt32", symEffect: "Read"},
		{name: "MOVLloadidx4", argLength: 3, reg: gploadidx, asm: "MOVL", aux: "SymOff", typ: "UInt32", symEffect: "Read"},
		{name: "MOVLloadidx8", argLength: 3, reg: gploadidx, asm: "MOVL", aux: "SymOff", typ: "UInt32", symEffect: "Read"},
		{name: "MOVQloadidx1", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVQ", aux: "SymOff", typ: "UInt64", symEffect: "Read"},
		{name: "MOVQloadidx8", argLength: 3, reg: gploadidx, asm: "MOVQ", aux: "SymOff", typ: "UInt64", symEffect: "Read"},

		{name: "MOVBstoreidx1", argLength: 4, reg: gpstoreidx, asm: "MOVB", aux: "SymOff", symEffect: "Write"},
		{name: "MOVWstoreidx1", argLength: 4, reg: gpstoreidx, asm: "MOVW", aux: "SymOff", symEffect: "Write"},
		{name: "MOVWstoreidx2", argLength: 4, reg: gpstoreidx, asm: "MOVW", aux: "SymOff", symEffect: "Write"},
		{name: "MOVLstoreidx1", argLength: 4, reg: gpstoreidx, asm: "MOVL", aux: "SymOff", symEffect: "Write"},
		{name: "MOVLstoreidx4", argLength: 4, reg: gpstoreidx, asm: "MOVL", aux: "SymOff", symEffect: "Write"},
		{name: "MOVLstoreidx8", argLength: 4, reg: gpstoreidx, asm: "MOVL", aux: "SymOff", symEffect: "Write"},
		{name: "MOVQstoreidx1", argLength: 4, reg: gpstoreidx, asm: "MOVQ", aux: "SymOff", symEffect: "Write"},
		{name: "MOVQstoreidx8", argLength: 4, reg: gpstoreidx, asm: "MOVQ", aux: "SymOff", symEffect: "Write"},

		{name: "MOVBstoreconst", argLength: 2, reg: gpstoreconst, asm: "MOVB", aux: "SymValAndOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstoreconst", argLength: 2, reg: gpstoreconst, asm: "MOVW", aux: "SymValAndOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVLstoreconst", argLength: 2, reg: gpstoreconst, asm: "MOVL", aux: "SymValAndOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVQstoreconst", argLength: 2, reg: gpstoreconst, asm: "MOVQ", aux: "SymValAndOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "MOVBstoreconstidx1", argLength: 3, reg: gpstoreconstidx, asm: "MOVB", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},
		{name: "MOVWstoreconstidx1", argLength: 3, reg: gpstoreconstidx, asm: "MOVW", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},
		{name: "MOVWstoreconstidx2", argLength: 3, reg: gpstoreconstidx, asm: "MOVW", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},
		{name: "MOVLstoreconstidx1", argLength: 3, reg: gpstoreconstidx, asm: "MOVL", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},
		{name: "MOVLstoreconstidx4", argLength: 3, reg: gpstoreconstidx, asm: "MOVL", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},
		{name: "MOVQstoreconstidx1", argLength: 3, reg: gpstoreconstidx, asm: "MOVQ", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},
		{name: "MOVQstoreconstidx8", argLength: 3, reg: gpstoreconstidx, asm: "MOVQ", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},

		{
			name:      "DUFFZERO",
			aux:       "Int64",
			argLength: 3,
			reg: regInfo{
				inputs:   []regMask{buildReg("DI"), buildReg("X0")},
				clobbers: buildReg("DI"),
			},
			faultOnNilArg0: true,
		},
		{name: "MOVOconst", reg: regInfo{nil, 0, []regMask{fp}}, typ: "Int128", aux: "Int128", rematerializeable: true},

		{
			name:      "REPSTOSQ",
			argLength: 4,
			reg: regInfo{
				inputs:   []regMask{buildReg("DI"), buildReg("CX"), buildReg("AX")},
				clobbers: buildReg("DI CX"),
			},
			faultOnNilArg0: true,
		},

		{name: "CALLstatic", argLength: 1, reg: regInfo{clobbers: callerSave}, aux: "SymOff", clobberFlags: true, call: true, symEffect: "None"},
		{name: "CALLclosure", argLength: 3, reg: regInfo{inputs: []regMask{gpsp, buildReg("DX"), 0}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},
		{name: "CALLinter", argLength: 2, reg: regInfo{inputs: []regMask{gp}, clobbers: callerSave}, aux: "Int64", clobberFlags: true, call: true},

		{
			name:      "DUFFCOPY",
			aux:       "Int64",
			argLength: 3,
			reg: regInfo{
				inputs:   []regMask{buildReg("DI"), buildReg("SI")},
				clobbers: buildReg("DI SI X0"),
			},
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
		},

		{
			name:      "REPMOVSQ",
			argLength: 4,
			reg: regInfo{
				inputs:   []regMask{buildReg("DI"), buildReg("SI"), buildReg("CX")},
				clobbers: buildReg("DI SI CX"),
			},
			faultOnNilArg0: true,
			faultOnNilArg1: true,
		},

		{name: "InvertFlags", argLength: 1},

		{name: "LoweredGetG", argLength: 1, reg: gp01},

		{name: "LoweredGetClosurePtr", reg: regInfo{outputs: []regMask{buildReg("DX")}}, zeroWidth: true},

		{name: "LoweredGetCallerPC", reg: gp01, rematerializeable: true},

		{name: "LoweredGetCallerSP", reg: gp01, rematerializeable: true},

		{name: "LoweredNilCheck", argLength: 2, reg: regInfo{inputs: []regMask{gpsp}}, clobberFlags: true, nilCheck: true, faultOnNilArg0: true},

		{name: "LoweredWB", argLength: 3, reg: regInfo{inputs: []regMask{buildReg("DI"), ax}, clobbers: callerSave &^ gp}, clobberFlags: true, aux: "Sym", symEffect: "None"},

		{name: "FlagEQ"},
		{name: "FlagLT_ULT"},
		{name: "FlagLT_UGT"},
		{name: "FlagGT_UGT"},
		{name: "FlagGT_ULT"},

		{name: "MOVLatomicload", argLength: 2, reg: gpload, asm: "MOVL", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVQatomicload", argLength: 2, reg: gpload, asm: "MOVQ", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},

		{name: "XCHGL", argLength: 3, reg: gpstorexchg, asm: "XCHGL", aux: "SymOff", resultInArg0: true, faultOnNilArg1: true, hasSideEffects: true, symEffect: "RdWr"},
		{name: "XCHGQ", argLength: 3, reg: gpstorexchg, asm: "XCHGQ", aux: "SymOff", resultInArg0: true, faultOnNilArg1: true, hasSideEffects: true, symEffect: "RdWr"},

		{name: "XADDLlock", argLength: 3, reg: gpstorexchg, asm: "XADDL", typ: "(UInt32,Mem)", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, hasSideEffects: true, symEffect: "RdWr"},
		{name: "XADDQlock", argLength: 3, reg: gpstorexchg, asm: "XADDQ", typ: "(UInt64,Mem)", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, hasSideEffects: true, symEffect: "RdWr"},
		{name: "AddTupleFirst32", argLength: 2},
		{name: "AddTupleFirst64", argLength: 2},

		{name: "CMPXCHGLlock", argLength: 4, reg: cmpxchg, asm: "CMPXCHGL", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "RdWr"},
		{name: "CMPXCHGQlock", argLength: 4, reg: cmpxchg, asm: "CMPXCHGQ", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "RdWr"},

		{name: "ANDBlock", argLength: 3, reg: gpstore, asm: "ANDB", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "RdWr"},
		{name: "ORBlock", argLength: 3, reg: gpstore, asm: "ORB", aux: "SymOff", clobberFlags: true, faultOnNilArg0: true, hasSideEffects: true, symEffect: "RdWr"},
	}

	var AMD64blocks = []blockData{
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
		{name: "EQF"},
		{name: "NEF"},
		{name: "ORD"},
		{name: "NAN"},
	}

	archs = append(archs, arch{
		name:            "AMD64",
		pkg:             "github.com/dave/golib/src/cmd/internal/obj/x86",
		genfile:         "../../amd64/ssa.go",
		ops:             AMD64ops,
		blocks:          AMD64blocks,
		regnames:        regNamesAMD64,
		gpregmask:       gp,
		fpregmask:       fp,
		framepointerreg: int8(num["BP"]),
		linkreg:         -1,
	})
}

// +build ignore

package main

import "strings"

// copied from ../../x86/reg.go
var regNames386 = []string{
	"AX",
	"CX",
	"DX",
	"BX",
	"SP",
	"BP",
	"SI",
	"DI",
	"X0",
	"X1",
	"X2",
	"X3",
	"X4",
	"X5",
	"X6",
	"X7",

	"SB",
}

func init() {

	if len(regNames386) > 64 {
		panic("too many registers")
	}
	num := map[string]int{}
	for i, name := range regNames386 {
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
		gp         = buildReg("AX CX DX BX BP SI DI")
		fp         = buildReg("X0 X1 X2 X3 X4 X5 X6 X7")
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
		gp11carry = regInfo{inputs: []regMask{gp}, outputs: []regMask{gp, 0}}
		gp21carry = regInfo{inputs: []regMask{gp, gp}, outputs: []regMask{gp, 0}}
		gp1carry1 = regInfo{inputs: []regMask{gp}, outputs: gponly}
		gp2carry1 = regInfo{inputs: []regMask{gp, gp}, outputs: gponly}
		gp21sp    = regInfo{inputs: []regMask{gpsp, gp}, outputs: gponly}
		gp21sb    = regInfo{inputs: []regMask{gpspsb, gpsp}, outputs: gponly}
		gp21shift = regInfo{inputs: []regMask{gp, cx}, outputs: []regMask{gp}}
		gp11div   = regInfo{inputs: []regMask{ax, gpsp &^ dx}, outputs: []regMask{ax}, clobbers: dx}
		gp21hmul  = regInfo{inputs: []regMask{ax, gpsp}, outputs: []regMask{dx}, clobbers: ax}
		gp11mod   = regInfo{inputs: []regMask{ax, gpsp &^ dx}, outputs: []regMask{dx}, clobbers: ax}
		gp21mul   = regInfo{inputs: []regMask{ax, gpsp}, outputs: []regMask{dx, ax}}

		gp2flags = regInfo{inputs: []regMask{gpsp, gpsp}}
		gp1flags = regInfo{inputs: []regMask{gpsp}}
		flagsgp  = regInfo{inputs: nil, outputs: gponly}

		readflags = regInfo{inputs: nil, outputs: gponly}
		flagsgpax = regInfo{inputs: nil, clobbers: ax, outputs: []regMask{gp &^ ax}}

		gpload    = regInfo{inputs: []regMask{gpspsb, 0}, outputs: gponly}
		gp21load  = regInfo{inputs: []regMask{gp, gpspsb, 0}, outputs: gponly}
		gploadidx = regInfo{inputs: []regMask{gpspsb, gpsp, 0}, outputs: gponly}

		gpstore         = regInfo{inputs: []regMask{gpspsb, gpsp, 0}}
		gpstoreconst    = regInfo{inputs: []regMask{gpspsb, 0}}
		gpstoreidx      = regInfo{inputs: []regMask{gpspsb, gpsp, gpsp, 0}}
		gpstoreconstidx = regInfo{inputs: []regMask{gpspsb, gpsp, 0}}

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

	var _386ops = []opData{

		{name: "ADDSS", argLength: 2, reg: fp21, asm: "ADDSS", commutative: true, resultInArg0: true, usesScratch: true},
		{name: "ADDSD", argLength: 2, reg: fp21, asm: "ADDSD", commutative: true, resultInArg0: true},
		{name: "SUBSS", argLength: 2, reg: fp21, asm: "SUBSS", resultInArg0: true, usesScratch: true},
		{name: "SUBSD", argLength: 2, reg: fp21, asm: "SUBSD", resultInArg0: true},
		{name: "MULSS", argLength: 2, reg: fp21, asm: "MULSS", commutative: true, resultInArg0: true, usesScratch: true},
		{name: "MULSD", argLength: 2, reg: fp21, asm: "MULSD", commutative: true, resultInArg0: true},
		{name: "DIVSS", argLength: 2, reg: fp21, asm: "DIVSS", resultInArg0: true, usesScratch: true},
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

		{name: "ADDL", argLength: 2, reg: gp21sp, asm: "ADDL", commutative: true, clobberFlags: true},
		{name: "ADDLconst", argLength: 1, reg: gp11sp, asm: "ADDL", aux: "Int32", typ: "UInt32", clobberFlags: true},

		{name: "ADDLcarry", argLength: 2, reg: gp21carry, asm: "ADDL", commutative: true, resultInArg0: true},
		{name: "ADDLconstcarry", argLength: 1, reg: gp11carry, asm: "ADDL", aux: "Int32", resultInArg0: true},
		{name: "ADCL", argLength: 3, reg: gp2carry1, asm: "ADCL", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "ADCLconst", argLength: 2, reg: gp1carry1, asm: "ADCL", aux: "Int32", resultInArg0: true, clobberFlags: true},

		{name: "SUBL", argLength: 2, reg: gp21, asm: "SUBL", resultInArg0: true, clobberFlags: true},
		{name: "SUBLconst", argLength: 1, reg: gp11, asm: "SUBL", aux: "Int32", resultInArg0: true, clobberFlags: true},

		{name: "SUBLcarry", argLength: 2, reg: gp21carry, asm: "SUBL", resultInArg0: true},
		{name: "SUBLconstcarry", argLength: 1, reg: gp11carry, asm: "SUBL", aux: "Int32", resultInArg0: true},
		{name: "SBBL", argLength: 3, reg: gp2carry1, asm: "SBBL", resultInArg0: true, clobberFlags: true},
		{name: "SBBLconst", argLength: 2, reg: gp1carry1, asm: "SBBL", aux: "Int32", resultInArg0: true, clobberFlags: true},

		{name: "MULL", argLength: 2, reg: gp21, asm: "IMULL", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "MULLconst", argLength: 1, reg: gp11, asm: "IMUL3L", aux: "Int32", clobberFlags: true},

		{name: "HMULL", argLength: 2, reg: gp21hmul, commutative: true, asm: "IMULL", clobberFlags: true},
		{name: "HMULLU", argLength: 2, reg: gp21hmul, commutative: true, asm: "MULL", clobberFlags: true},

		{name: "MULLQU", argLength: 2, reg: gp21mul, commutative: true, asm: "MULL", clobberFlags: true},

		{name: "AVGLU", argLength: 2, reg: gp21, commutative: true, resultInArg0: true, clobberFlags: true},

		{name: "DIVL", argLength: 2, reg: gp11div, asm: "IDIVL", clobberFlags: true},
		{name: "DIVW", argLength: 2, reg: gp11div, asm: "IDIVW", clobberFlags: true},
		{name: "DIVLU", argLength: 2, reg: gp11div, asm: "DIVL", clobberFlags: true},
		{name: "DIVWU", argLength: 2, reg: gp11div, asm: "DIVW", clobberFlags: true},

		{name: "MODL", argLength: 2, reg: gp11mod, asm: "IDIVL", clobberFlags: true},
		{name: "MODW", argLength: 2, reg: gp11mod, asm: "IDIVW", clobberFlags: true},
		{name: "MODLU", argLength: 2, reg: gp11mod, asm: "DIVL", clobberFlags: true},
		{name: "MODWU", argLength: 2, reg: gp11mod, asm: "DIVW", clobberFlags: true},

		{name: "ANDL", argLength: 2, reg: gp21, asm: "ANDL", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "ANDLconst", argLength: 1, reg: gp11, asm: "ANDL", aux: "Int32", resultInArg0: true, clobberFlags: true},

		{name: "ORL", argLength: 2, reg: gp21, asm: "ORL", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "ORLconst", argLength: 1, reg: gp11, asm: "ORL", aux: "Int32", resultInArg0: true, clobberFlags: true},

		{name: "XORL", argLength: 2, reg: gp21, asm: "XORL", commutative: true, resultInArg0: true, clobberFlags: true},
		{name: "XORLconst", argLength: 1, reg: gp11, asm: "XORL", aux: "Int32", resultInArg0: true, clobberFlags: true},

		{name: "CMPL", argLength: 2, reg: gp2flags, asm: "CMPL", typ: "Flags"},
		{name: "CMPW", argLength: 2, reg: gp2flags, asm: "CMPW", typ: "Flags"},
		{name: "CMPB", argLength: 2, reg: gp2flags, asm: "CMPB", typ: "Flags"},
		{name: "CMPLconst", argLength: 1, reg: gp1flags, asm: "CMPL", typ: "Flags", aux: "Int32"},
		{name: "CMPWconst", argLength: 1, reg: gp1flags, asm: "CMPW", typ: "Flags", aux: "Int16"},
		{name: "CMPBconst", argLength: 1, reg: gp1flags, asm: "CMPB", typ: "Flags", aux: "Int8"},

		{name: "UCOMISS", argLength: 2, reg: fp2flags, asm: "UCOMISS", typ: "Flags", usesScratch: true},
		{name: "UCOMISD", argLength: 2, reg: fp2flags, asm: "UCOMISD", typ: "Flags", usesScratch: true},

		{name: "TESTL", argLength: 2, reg: gp2flags, commutative: true, asm: "TESTL", typ: "Flags"},
		{name: "TESTW", argLength: 2, reg: gp2flags, commutative: true, asm: "TESTW", typ: "Flags"},
		{name: "TESTB", argLength: 2, reg: gp2flags, commutative: true, asm: "TESTB", typ: "Flags"},
		{name: "TESTLconst", argLength: 1, reg: gp1flags, asm: "TESTL", typ: "Flags", aux: "Int32"},
		{name: "TESTWconst", argLength: 1, reg: gp1flags, asm: "TESTW", typ: "Flags", aux: "Int16"},
		{name: "TESTBconst", argLength: 1, reg: gp1flags, asm: "TESTB", typ: "Flags", aux: "Int8"},

		{name: "SHLL", argLength: 2, reg: gp21shift, asm: "SHLL", resultInArg0: true, clobberFlags: true},
		{name: "SHLLconst", argLength: 1, reg: gp11, asm: "SHLL", aux: "Int32", resultInArg0: true, clobberFlags: true},

		{name: "SHRL", argLength: 2, reg: gp21shift, asm: "SHRL", resultInArg0: true, clobberFlags: true},
		{name: "SHRW", argLength: 2, reg: gp21shift, asm: "SHRW", resultInArg0: true, clobberFlags: true},
		{name: "SHRB", argLength: 2, reg: gp21shift, asm: "SHRB", resultInArg0: true, clobberFlags: true},
		{name: "SHRLconst", argLength: 1, reg: gp11, asm: "SHRL", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "SHRWconst", argLength: 1, reg: gp11, asm: "SHRW", aux: "Int16", resultInArg0: true, clobberFlags: true},
		{name: "SHRBconst", argLength: 1, reg: gp11, asm: "SHRB", aux: "Int8", resultInArg0: true, clobberFlags: true},

		{name: "SARL", argLength: 2, reg: gp21shift, asm: "SARL", resultInArg0: true, clobberFlags: true},
		{name: "SARW", argLength: 2, reg: gp21shift, asm: "SARW", resultInArg0: true, clobberFlags: true},
		{name: "SARB", argLength: 2, reg: gp21shift, asm: "SARB", resultInArg0: true, clobberFlags: true},
		{name: "SARLconst", argLength: 1, reg: gp11, asm: "SARL", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "SARWconst", argLength: 1, reg: gp11, asm: "SARW", aux: "Int16", resultInArg0: true, clobberFlags: true},
		{name: "SARBconst", argLength: 1, reg: gp11, asm: "SARB", aux: "Int8", resultInArg0: true, clobberFlags: true},

		{name: "ROLLconst", argLength: 1, reg: gp11, asm: "ROLL", aux: "Int32", resultInArg0: true, clobberFlags: true},
		{name: "ROLWconst", argLength: 1, reg: gp11, asm: "ROLW", aux: "Int16", resultInArg0: true, clobberFlags: true},
		{name: "ROLBconst", argLength: 1, reg: gp11, asm: "ROLB", aux: "Int8", resultInArg0: true, clobberFlags: true},

		{name: "ADDLload", argLength: 3, reg: gp21load, asm: "ADDL", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "SUBLload", argLength: 3, reg: gp21load, asm: "SUBL", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "ANDLload", argLength: 3, reg: gp21load, asm: "ANDL", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "ORLload", argLength: 3, reg: gp21load, asm: "ORL", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},
		{name: "XORLload", argLength: 3, reg: gp21load, asm: "XORL", aux: "SymOff", resultInArg0: true, clobberFlags: true, faultOnNilArg1: true, symEffect: "Read"},

		{name: "NEGL", argLength: 1, reg: gp11, asm: "NEGL", resultInArg0: true, clobberFlags: true},

		{name: "NOTL", argLength: 1, reg: gp11, asm: "NOTL", resultInArg0: true, clobberFlags: true},

		{name: "BSFL", argLength: 1, reg: gp11, asm: "BSFL", clobberFlags: true},
		{name: "BSFW", argLength: 1, reg: gp11, asm: "BSFW", clobberFlags: true},

		{name: "BSRL", argLength: 1, reg: gp11, asm: "BSRL", clobberFlags: true},
		{name: "BSRW", argLength: 1, reg: gp11, asm: "BSRW", clobberFlags: true},

		{name: "BSWAPL", argLength: 1, reg: gp11, asm: "BSWAPL", resultInArg0: true, clobberFlags: true},

		{name: "SQRTSD", argLength: 1, reg: fp11, asm: "SQRTSD"},

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

		{name: "SETEQF", argLength: 1, reg: flagsgpax, asm: "SETEQ", clobberFlags: true},
		{name: "SETNEF", argLength: 1, reg: flagsgpax, asm: "SETNE", clobberFlags: true},
		{name: "SETORD", argLength: 1, reg: flagsgp, asm: "SETPC"},
		{name: "SETNAN", argLength: 1, reg: flagsgp, asm: "SETPS"},

		{name: "SETGF", argLength: 1, reg: flagsgp, asm: "SETHI"},
		{name: "SETGEF", argLength: 1, reg: flagsgp, asm: "SETCC"},

		{name: "MOVBLSX", argLength: 1, reg: gp11, asm: "MOVBLSX"},
		{name: "MOVBLZX", argLength: 1, reg: gp11, asm: "MOVBLZX"},
		{name: "MOVWLSX", argLength: 1, reg: gp11, asm: "MOVWLSX"},
		{name: "MOVWLZX", argLength: 1, reg: gp11, asm: "MOVWLZX"},

		{name: "MOVLconst", reg: gp01, asm: "MOVL", typ: "UInt32", aux: "Int32", rematerializeable: true},

		{name: "CVTTSD2SL", argLength: 1, reg: fpgp, asm: "CVTTSD2SL", usesScratch: true},
		{name: "CVTTSS2SL", argLength: 1, reg: fpgp, asm: "CVTTSS2SL", usesScratch: true},
		{name: "CVTSL2SS", argLength: 1, reg: gpfp, asm: "CVTSL2SS", usesScratch: true},
		{name: "CVTSL2SD", argLength: 1, reg: gpfp, asm: "CVTSL2SD", usesScratch: true},
		{name: "CVTSD2SS", argLength: 1, reg: fp11, asm: "CVTSD2SS", usesScratch: true},
		{name: "CVTSS2SD", argLength: 1, reg: fp11, asm: "CVTSS2SD"},

		{name: "PXOR", argLength: 2, reg: fp21, asm: "PXOR", commutative: true, resultInArg0: true},

		{name: "LEAL", argLength: 1, reg: gp11sb, aux: "SymOff", rematerializeable: true, symEffect: "Addr"},
		{name: "LEAL1", argLength: 2, reg: gp21sb, commutative: true, aux: "SymOff", symEffect: "Addr"},
		{name: "LEAL2", argLength: 2, reg: gp21sb, aux: "SymOff", symEffect: "Addr"},
		{name: "LEAL4", argLength: 2, reg: gp21sb, aux: "SymOff", symEffect: "Addr"},
		{name: "LEAL8", argLength: 2, reg: gp21sb, aux: "SymOff", symEffect: "Addr"},

		{name: "MOVBload", argLength: 2, reg: gpload, asm: "MOVBLZX", aux: "SymOff", typ: "UInt8", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVBLSXload", argLength: 2, reg: gpload, asm: "MOVBLSX", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWload", argLength: 2, reg: gpload, asm: "MOVWLZX", aux: "SymOff", typ: "UInt16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWLSXload", argLength: 2, reg: gpload, asm: "MOVWLSX", aux: "SymOff", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVLload", argLength: 2, reg: gpload, asm: "MOVL", aux: "SymOff", typ: "UInt32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVBstore", argLength: 3, reg: gpstore, asm: "MOVB", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstore", argLength: 3, reg: gpstore, asm: "MOVW", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVLstore", argLength: 3, reg: gpstore, asm: "MOVL", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "ADDLmodify", argLength: 3, reg: gpstore, asm: "ADDL", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Read,Write"},
		{name: "SUBLmodify", argLength: 3, reg: gpstore, asm: "SUBL", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Read,Write"},
		{name: "ANDLmodify", argLength: 3, reg: gpstore, asm: "ANDL", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Read,Write"},
		{name: "ORLmodify", argLength: 3, reg: gpstore, asm: "ORL", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Read,Write"},
		{name: "XORLmodify", argLength: 3, reg: gpstore, asm: "XORL", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Read,Write"},

		{name: "MOVBloadidx1", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVBLZX", aux: "SymOff", symEffect: "Read"},
		{name: "MOVWloadidx1", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVWLZX", aux: "SymOff", symEffect: "Read"},
		{name: "MOVWloadidx2", argLength: 3, reg: gploadidx, asm: "MOVWLZX", aux: "SymOff", symEffect: "Read"},
		{name: "MOVLloadidx1", argLength: 3, reg: gploadidx, commutative: true, asm: "MOVL", aux: "SymOff", symEffect: "Read"},
		{name: "MOVLloadidx4", argLength: 3, reg: gploadidx, asm: "MOVL", aux: "SymOff", symEffect: "Read"},

		{name: "MOVBstoreidx1", argLength: 4, reg: gpstoreidx, commutative: true, asm: "MOVB", aux: "SymOff", symEffect: "Write"},
		{name: "MOVWstoreidx1", argLength: 4, reg: gpstoreidx, commutative: true, asm: "MOVW", aux: "SymOff", symEffect: "Write"},
		{name: "MOVWstoreidx2", argLength: 4, reg: gpstoreidx, asm: "MOVW", aux: "SymOff", symEffect: "Write"},
		{name: "MOVLstoreidx1", argLength: 4, reg: gpstoreidx, commutative: true, asm: "MOVL", aux: "SymOff", symEffect: "Write"},
		{name: "MOVLstoreidx4", argLength: 4, reg: gpstoreidx, asm: "MOVL", aux: "SymOff", symEffect: "Write"},

		{name: "MOVBstoreconst", argLength: 2, reg: gpstoreconst, asm: "MOVB", aux: "SymValAndOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstoreconst", argLength: 2, reg: gpstoreconst, asm: "MOVW", aux: "SymValAndOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVLstoreconst", argLength: 2, reg: gpstoreconst, asm: "MOVL", aux: "SymValAndOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		{name: "MOVBstoreconstidx1", argLength: 3, reg: gpstoreconstidx, asm: "MOVB", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},
		{name: "MOVWstoreconstidx1", argLength: 3, reg: gpstoreconstidx, asm: "MOVW", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},
		{name: "MOVWstoreconstidx2", argLength: 3, reg: gpstoreconstidx, asm: "MOVW", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},
		{name: "MOVLstoreconstidx1", argLength: 3, reg: gpstoreconstidx, asm: "MOVL", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},
		{name: "MOVLstoreconstidx4", argLength: 3, reg: gpstoreconstidx, asm: "MOVL", aux: "SymValAndOff", typ: "Mem", symEffect: "Write"},

		{
			name:      "DUFFZERO",
			aux:       "Int64",
			argLength: 3,
			reg: regInfo{
				inputs:   []regMask{buildReg("DI"), buildReg("AX")},
				clobbers: buildReg("DI CX"),
			},
			faultOnNilArg0: true,
		},

		{
			name:      "REPSTOSL",
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
				clobbers: buildReg("DI SI CX"),
			},
			clobberFlags:   true,
			faultOnNilArg0: true,
			faultOnNilArg1: true,
		},

		{
			name:      "REPMOVSL",
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

		{name: "FCHS", argLength: 1, reg: fp11},

		{name: "MOVSSconst1", reg: gp01, typ: "UInt32", aux: "Float32"},
		{name: "MOVSDconst1", reg: gp01, typ: "UInt32", aux: "Float64"},
		{name: "MOVSSconst2", argLength: 1, reg: gpfp, asm: "MOVSS"},
		{name: "MOVSDconst2", argLength: 1, reg: gpfp, asm: "MOVSD"},
	}

	var _386blocks = []blockData{
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
		name:            "386",
		pkg:             "github.com/dave/golib/src/cmd/internal/obj/x86",
		genfile:         "../../x86/ssa.go",
		ops:             _386ops,
		blocks:          _386blocks,
		regnames:        regNames386,
		gpregmask:       gp,
		fpregmask:       fp,
		framepointerreg: int8(num["BP"]),
		linkreg:         -1,
	})
}

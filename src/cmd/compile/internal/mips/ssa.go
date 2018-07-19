package mips

import (
	"math"

	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/mips"
)

// isFPreg returns whether r is an FP register
func isFPreg(r int16) bool {
	return mips.REG_F0 <= r && r <= mips.REG_F31
}

// isHILO returns whether r is HI or LO register
func isHILO(r int16) bool {
	return r == mips.REG_HI || r == mips.REG_LO
}

// loadByType returns the load instruction of the given type.
func (psess *PackageSession) loadByType(t *types.Type, r int16) obj.As {
	if isFPreg(r) {
		if t.Size(psess.types) == 4 {
			return mips.AMOVF
		} else {
			return mips.AMOVD
		}
	} else {
		switch t.Size(psess.types) {
		case 1:
			if t.IsSigned() {
				return mips.AMOVB
			} else {
				return mips.AMOVBU
			}
		case 2:
			if t.IsSigned() {
				return mips.AMOVH
			} else {
				return mips.AMOVHU
			}
		case 4:
			return mips.AMOVW
		}
	}
	panic("bad load type")
}

// storeByType returns the store instruction of the given type.
func (psess *PackageSession) storeByType(t *types.Type, r int16) obj.As {
	if isFPreg(r) {
		if t.Size(psess.types) == 4 {
			return mips.AMOVF
		} else {
			return mips.AMOVD
		}
	} else {
		switch t.Size(psess.types) {
		case 1:
			return mips.AMOVB
		case 2:
			return mips.AMOVH
		case 4:
			return mips.AMOVW
		}
	}
	panic("bad store type")
}

func (psess *PackageSession) ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpCopy, ssa.OpMIPSMOVWreg:
		t := v.Type
		if t.IsMemory(psess.types) {
			return
		}
		x := v.Args[0].Reg(psess.ssa)
		y := v.Reg(psess.ssa)
		if x == y {
			return
		}
		as := mips.AMOVW
		if isFPreg(x) && isFPreg(y) {
			as = mips.AMOVF
			if t.Size(psess.types) == 8 {
				as = mips.AMOVD
			}
		}

		p := s.Prog(psess.gc, as)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x
		p.To.Type = obj.TYPE_REG
		p.To.Reg = y
		if isHILO(x) && isHILO(y) || isHILO(x) && isFPreg(y) || isFPreg(x) && isHILO(y) {

			p.To.Reg = mips.REGTMP
			p = s.Prog(psess.gc, mips.AMOVW)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = mips.REGTMP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = y
		}
	case ssa.OpMIPSMOVWnop:
		if v.Reg(psess.ssa) != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}

	case ssa.OpLoadReg:
		if v.Type.IsFlags(psess.types) {
			v.Fatalf("load flags not implemented: %v", v.LongString(psess.ssa))
			return
		}
		r := v.Reg(psess.ssa)
		p := s.Prog(psess.gc, psess.loadByType(v.Type, r))
		psess.gc.
			AddrAuto(&p.From, v.Args[0])
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		if isHILO(r) {

			p.To.Reg = mips.REGTMP
			p = s.Prog(psess.gc, mips.AMOVW)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = mips.REGTMP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = r
		}
	case ssa.OpStoreReg:
		if v.Type.IsFlags(psess.types) {
			v.Fatalf("store flags not implemented: %v", v.LongString(psess.ssa))
			return
		}
		r := v.Args[0].Reg(psess.ssa)
		if isHILO(r) {

			p := s.Prog(psess.gc, mips.AMOVW)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = r
			p.To.Type = obj.TYPE_REG
			p.To.Reg = mips.REGTMP
			r = mips.REGTMP
		}
		p := s.Prog(psess.gc, psess.storeByType(v.Type, r))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r
		psess.gc.
			AddrAuto(&p.To, v)
	case ssa.OpMIPSADD,
		ssa.OpMIPSSUB,
		ssa.OpMIPSAND,
		ssa.OpMIPSOR,
		ssa.OpMIPSXOR,
		ssa.OpMIPSNOR,
		ssa.OpMIPSSLL,
		ssa.OpMIPSSRL,
		ssa.OpMIPSSRA,
		ssa.OpMIPSADDF,
		ssa.OpMIPSADDD,
		ssa.OpMIPSSUBF,
		ssa.OpMIPSSUBD,
		ssa.OpMIPSMULF,
		ssa.OpMIPSMULD,
		ssa.OpMIPSDIVF,
		ssa.OpMIPSDIVD,
		ssa.OpMIPSMUL:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSSGT,
		ssa.OpMIPSSGTU:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSSGTzero,
		ssa.OpMIPSSGTUzero:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.Reg = mips.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSADDconst,
		ssa.OpMIPSSUBconst,
		ssa.OpMIPSANDconst,
		ssa.OpMIPSORconst,
		ssa.OpMIPSXORconst,
		ssa.OpMIPSNORconst,
		ssa.OpMIPSSLLconst,
		ssa.OpMIPSSRLconst,
		ssa.OpMIPSSRAconst,
		ssa.OpMIPSSGTconst,
		ssa.OpMIPSSGTUconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSMULT,
		ssa.OpMIPSMULTU,
		ssa.OpMIPSDIV,
		ssa.OpMIPSDIVU:

		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.Reg = v.Args[0].Reg(psess.ssa)
	case ssa.OpMIPSMOVWconst:
		r := v.Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		if isFPreg(r) || isHILO(r) {

			p.To.Reg = mips.REGTMP
			p = s.Prog(psess.gc, mips.AMOVW)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = mips.REGTMP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = r
		}
	case ssa.OpMIPSMOVFconst,
		ssa.OpMIPSMOVDconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_FCONST
		p.From.Val = math.Float64frombits(uint64(v.AuxInt))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSCMOVZ:
		if v.Reg(psess.ssa) != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[2].Reg(psess.ssa)
		p.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSCMOVZzero:
		if v.Reg(psess.ssa) != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.Reg = mips.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSCMPEQF,
		ssa.OpMIPSCMPEQD,
		ssa.OpMIPSCMPGEF,
		ssa.OpMIPSCMPGED,
		ssa.OpMIPSCMPGTF,
		ssa.OpMIPSCMPGTD:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.Reg = v.Args[1].Reg(psess.ssa)
	case ssa.OpMIPSMOVWaddr:
		p := s.Prog(psess.gc, mips.AMOVW)
		p.From.Type = obj.TYPE_ADDR
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		var wantreg string

		switch v.Aux.(type) {
		default:
			v.Fatalf("aux is of unknown type %T", v.Aux)
		case *obj.LSym:
			wantreg = "SB"
			psess.gc.
				AddAux(&p.From, v)
		case *gc.Node:
			wantreg = "SP"
			psess.gc.
				AddAux(&p.From, v)
		case nil:

			wantreg = "SP"
			p.From.Offset = v.AuxInt
		}
		if reg := v.Args[0].RegName(psess.ssa); reg != wantreg {
			v.Fatalf("bad reg %s for symbol type %T, want %s", reg, v.Aux, wantreg)
		}
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSMOVBload,
		ssa.OpMIPSMOVBUload,
		ssa.OpMIPSMOVHload,
		ssa.OpMIPSMOVHUload,
		ssa.OpMIPSMOVWload,
		ssa.OpMIPSMOVFload,
		ssa.OpMIPSMOVDload:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSMOVBstore,
		ssa.OpMIPSMOVHstore,
		ssa.OpMIPSMOVWstore,
		ssa.OpMIPSMOVFstore,
		ssa.OpMIPSMOVDstore:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpMIPSMOVBstorezero,
		ssa.OpMIPSMOVHstorezero,
		ssa.OpMIPSMOVWstorezero:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = mips.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.OpMIPSMOVBreg,
		ssa.OpMIPSMOVBUreg,
		ssa.OpMIPSMOVHreg,
		ssa.OpMIPSMOVHUreg:
		a := v.Args[0]
		for a.Op == ssa.OpCopy || a.Op == ssa.OpMIPSMOVWreg || a.Op == ssa.OpMIPSMOVWnop {
			a = a.Args[0]
		}
		if a.Op == ssa.OpLoadReg {
			t := a.Type
			switch {
			case v.Op == ssa.OpMIPSMOVBreg && t.Size(psess.types) == 1 && t.IsSigned(),
				v.Op == ssa.OpMIPSMOVBUreg && t.Size(psess.types) == 1 && !t.IsSigned(),
				v.Op == ssa.OpMIPSMOVHreg && t.Size(psess.types) == 2 && t.IsSigned(),
				v.Op == ssa.OpMIPSMOVHUreg && t.Size(psess.types) == 2 && !t.IsSigned():

				if v.Reg(psess.ssa) == v.Args[0].Reg(psess.ssa) {
					return
				}
				p := s.Prog(psess.gc, mips.AMOVW)
				p.From.Type = obj.TYPE_REG
				p.From.Reg = v.Args[0].Reg(psess.ssa)
				p.To.Type = obj.TYPE_REG
				p.To.Reg = v.Reg(psess.ssa)
				return
			default:
			}
		}
		fallthrough
	case ssa.OpMIPSMOVWF,
		ssa.OpMIPSMOVWD,
		ssa.OpMIPSTRUNCFW,
		ssa.OpMIPSTRUNCDW,
		ssa.OpMIPSMOVFD,
		ssa.OpMIPSMOVDF,
		ssa.OpMIPSNEGF,
		ssa.OpMIPSNEGD,
		ssa.OpMIPSSQRTD,
		ssa.OpMIPSCLZ:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSNEG:

		p := s.Prog(psess.gc, mips.ASUBU)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.Reg = mips.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSLoweredZero:
		// SUBU	$4, R1
		// MOVW	R0, 4(R1)
		// ADDU	$4, R1
		// BNE	Rarg1, R1, -2(PC)
		// arg1 is the address of the last element to zero
		var sz int64
		var mov obj.As
		switch {
		case v.AuxInt%4 == 0:
			sz = 4
			mov = mips.AMOVW
		case v.AuxInt%2 == 0:
			sz = 2
			mov = mips.AMOVH
		default:
			sz = 1
			mov = mips.AMOVB
		}
		p := s.Prog(psess.gc, mips.ASUBU)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = sz
		p.To.Type = obj.TYPE_REG
		p.To.Reg = mips.REG_R1
		p2 := s.Prog(psess.gc, mov)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = mips.REGZERO
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = mips.REG_R1
		p2.To.Offset = sz
		p3 := s.Prog(psess.gc, mips.AADDU)
		p3.From.Type = obj.TYPE_CONST
		p3.From.Offset = sz
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = mips.REG_R1
		p4 := s.Prog(psess.gc, mips.ABNE)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = v.Args[1].Reg(psess.ssa)
		p4.Reg = mips.REG_R1
		p4.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p4, p2)
	case ssa.OpMIPSLoweredMove:
		// SUBU	$4, R1
		// MOVW	4(R1), Rtmp
		// MOVW	Rtmp, (R2)
		// ADDU	$4, R1
		// ADDU	$4, R2
		// BNE	Rarg2, R1, -4(PC)
		// arg2 is the address of the last element of src
		var sz int64
		var mov obj.As
		switch {
		case v.AuxInt%4 == 0:
			sz = 4
			mov = mips.AMOVW
		case v.AuxInt%2 == 0:
			sz = 2
			mov = mips.AMOVH
		default:
			sz = 1
			mov = mips.AMOVB
		}
		p := s.Prog(psess.gc, mips.ASUBU)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = sz
		p.To.Type = obj.TYPE_REG
		p.To.Reg = mips.REG_R1
		p2 := s.Prog(psess.gc, mov)
		p2.From.Type = obj.TYPE_MEM
		p2.From.Reg = mips.REG_R1
		p2.From.Offset = sz
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = mips.REGTMP
		p3 := s.Prog(psess.gc, mov)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = mips.REGTMP
		p3.To.Type = obj.TYPE_MEM
		p3.To.Reg = mips.REG_R2
		p4 := s.Prog(psess.gc, mips.AADDU)
		p4.From.Type = obj.TYPE_CONST
		p4.From.Offset = sz
		p4.To.Type = obj.TYPE_REG
		p4.To.Reg = mips.REG_R1
		p5 := s.Prog(psess.gc, mips.AADDU)
		p5.From.Type = obj.TYPE_CONST
		p5.From.Offset = sz
		p5.To.Type = obj.TYPE_REG
		p5.To.Reg = mips.REG_R2
		p6 := s.Prog(psess.gc, mips.ABNE)
		p6.From.Type = obj.TYPE_REG
		p6.From.Reg = v.Args[2].Reg(psess.ssa)
		p6.Reg = mips.REG_R1
		p6.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p6, p2)
	case ssa.OpMIPSCALLstatic, ssa.OpMIPSCALLclosure, ssa.OpMIPSCALLinter:
		s.Call(psess.gc, v)
	case ssa.OpMIPSLoweredWB:
		p := s.Prog(psess.gc, obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = v.Aux.(*obj.LSym)
	case ssa.OpMIPSLoweredAtomicLoad:
		s.Prog(psess.gc, mips.ASYNC)

		p := s.Prog(psess.gc, mips.AMOVW)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0(psess.ssa)

		s.Prog(psess.gc, mips.ASYNC)
	case ssa.OpMIPSLoweredAtomicStore:
		s.Prog(psess.gc, mips.ASYNC)

		p := s.Prog(psess.gc, mips.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)

		s.Prog(psess.gc, mips.ASYNC)
	case ssa.OpMIPSLoweredAtomicStorezero:
		s.Prog(psess.gc, mips.ASYNC)

		p := s.Prog(psess.gc, mips.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = mips.REGZERO
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)

		s.Prog(psess.gc, mips.ASYNC)
	case ssa.OpMIPSLoweredAtomicExchange:

		s.Prog(psess.gc, mips.ASYNC)

		p := s.Prog(psess.gc, mips.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = mips.REGTMP

		p1 := s.Prog(psess.gc, mips.ALL)
		p1.From.Type = obj.TYPE_MEM
		p1.From.Reg = v.Args[0].Reg(psess.ssa)
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = v.Reg0(psess.ssa)

		p2 := s.Prog(psess.gc, mips.ASC)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = mips.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = v.Args[0].Reg(psess.ssa)

		p3 := s.Prog(psess.gc, mips.ABEQ)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = mips.REGTMP
		p3.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p3, p)

		s.Prog(psess.gc, mips.ASYNC)
	case ssa.OpMIPSLoweredAtomicAdd:

		s.Prog(psess.gc, mips.ASYNC)

		p := s.Prog(psess.gc, mips.ALL)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0(psess.ssa)

		p1 := s.Prog(psess.gc, mips.AADDU)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = v.Args[1].Reg(psess.ssa)
		p1.Reg = v.Reg0(psess.ssa)
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = mips.REGTMP

		p2 := s.Prog(psess.gc, mips.ASC)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = mips.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = v.Args[0].Reg(psess.ssa)

		p3 := s.Prog(psess.gc, mips.ABEQ)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = mips.REGTMP
		p3.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p3, p)

		s.Prog(psess.gc, mips.ASYNC)

		p4 := s.Prog(psess.gc, mips.AADDU)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = v.Args[1].Reg(psess.ssa)
		p4.Reg = v.Reg0(psess.ssa)
		p4.To.Type = obj.TYPE_REG
		p4.To.Reg = v.Reg0(psess.ssa)

	case ssa.OpMIPSLoweredAtomicAddconst:

		s.Prog(psess.gc, mips.ASYNC)

		p := s.Prog(psess.gc, mips.ALL)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0(psess.ssa)

		p1 := s.Prog(psess.gc, mips.AADDU)
		p1.From.Type = obj.TYPE_CONST
		p1.From.Offset = v.AuxInt
		p1.Reg = v.Reg0(psess.ssa)
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = mips.REGTMP

		p2 := s.Prog(psess.gc, mips.ASC)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = mips.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = v.Args[0].Reg(psess.ssa)

		p3 := s.Prog(psess.gc, mips.ABEQ)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = mips.REGTMP
		p3.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p3, p)

		s.Prog(psess.gc, mips.ASYNC)

		p4 := s.Prog(psess.gc, mips.AADDU)
		p4.From.Type = obj.TYPE_CONST
		p4.From.Offset = v.AuxInt
		p4.Reg = v.Reg0(psess.ssa)
		p4.To.Type = obj.TYPE_REG
		p4.To.Reg = v.Reg0(psess.ssa)

	case ssa.OpMIPSLoweredAtomicAnd,
		ssa.OpMIPSLoweredAtomicOr:

		s.Prog(psess.gc, mips.ASYNC)

		p := s.Prog(psess.gc, mips.ALL)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = mips.REGTMP

		p1 := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = v.Args[1].Reg(psess.ssa)
		p1.Reg = mips.REGTMP
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = mips.REGTMP

		p2 := s.Prog(psess.gc, mips.ASC)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = mips.REGTMP
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = v.Args[0].Reg(psess.ssa)

		p3 := s.Prog(psess.gc, mips.ABEQ)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = mips.REGTMP
		p3.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p3, p)

		s.Prog(psess.gc, mips.ASYNC)

	case ssa.OpMIPSLoweredAtomicCas:

		p := s.Prog(psess.gc, mips.AMOVW)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = mips.REGZERO
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0(psess.ssa)

		s.Prog(psess.gc, mips.ASYNC)

		p1 := s.Prog(psess.gc, mips.ALL)
		p1.From.Type = obj.TYPE_MEM
		p1.From.Reg = v.Args[0].Reg(psess.ssa)
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = mips.REGTMP

		p2 := s.Prog(psess.gc, mips.ABNE)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = v.Args[1].Reg(psess.ssa)
		p2.Reg = mips.REGTMP
		p2.To.Type = obj.TYPE_BRANCH

		p3 := s.Prog(psess.gc, mips.AMOVW)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = v.Args[2].Reg(psess.ssa)
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = v.Reg0(psess.ssa)

		p4 := s.Prog(psess.gc, mips.ASC)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = v.Reg0(psess.ssa)
		p4.To.Type = obj.TYPE_MEM
		p4.To.Reg = v.Args[0].Reg(psess.ssa)

		p5 := s.Prog(psess.gc, mips.ABEQ)
		p5.From.Type = obj.TYPE_REG
		p5.From.Reg = v.Reg0(psess.ssa)
		p5.To.Type = obj.TYPE_BRANCH
		psess.gc.
			Patch(p5, p1)

		s.Prog(psess.gc, mips.ASYNC)

		p6 := s.Prog(psess.gc, obj.ANOP)
		psess.gc.
			Patch(p2, p6)

	case ssa.OpMIPSLoweredNilCheck:

		p := s.Prog(psess.gc, mips.AMOVB)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = mips.REGTMP
		if psess.gc.Debug_checknil != 0 && v.Pos.Line() > 1 {
			psess.gc.
				Warnl(v.Pos, "generated nil check")
		}
	case ssa.OpMIPSFPFlagTrue,
		ssa.OpMIPSFPFlagFalse:

		cmov := mips.ACMOVF
		if v.Op == ssa.OpMIPSFPFlagFalse {
			cmov = mips.ACMOVT
		}
		p := s.Prog(psess.gc, mips.AMOVW)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
		p1 := s.Prog(psess.gc, cmov)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = mips.REGZERO
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = v.Reg(psess.ssa)

	case ssa.OpMIPSLoweredGetClosurePtr:
		psess.gc.
			CheckLoweredGetClosurePtr(v)
	case ssa.OpMIPSLoweredGetCallerSP:

		p := s.Prog(psess.gc, mips.AMOVW)
		p.From.Type = obj.TYPE_ADDR
		p.From.Offset = -psess.gc.Ctxt.FixedFrameSize()
		p.From.Name = obj.NAME_PARAM
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpMIPSLoweredGetCallerPC:
		p := s.Prog(psess.gc, obj.AGETCALLERPC)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.OpClobber:

	default:
		v.Fatalf("genValue not implemented: %s", v.LongString(psess.ssa))
	}
}

func (psess *PackageSession) ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	switch b.Kind {
	case ssa.BlockPlain:
		if b.Succs[0].Block() != next {
			p := s.Prog(psess.gc, obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}
	case ssa.BlockDefer:

		p := s.Prog(psess.gc, mips.ABNE)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = mips.REGZERO
		p.Reg = mips.REG_R1
		p.To.Type = obj.TYPE_BRANCH
		s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[1].Block()})
		if b.Succs[0].Block() != next {
			p := s.Prog(psess.gc, obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, gc.Branch{P: p, B: b.Succs[0].Block()})
		}
	case ssa.BlockExit:
		s.Prog(psess.gc, obj.AUNDEF)
	case ssa.BlockRet:
		s.Prog(psess.gc, obj.ARET)
	case ssa.BlockRetJmp:
		p := s.Prog(psess.gc, obj.ARET)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = b.Aux.(*obj.LSym)
	case ssa.BlockMIPSEQ, ssa.BlockMIPSNE,
		ssa.BlockMIPSLTZ, ssa.BlockMIPSGEZ,
		ssa.BlockMIPSLEZ, ssa.BlockMIPSGTZ,
		ssa.BlockMIPSFPT, ssa.BlockMIPSFPF:
		jmp := psess.blockJump[b.Kind]
		var p *obj.Prog
		switch next {
		case b.Succs[0].Block():
			p = s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
		case b.Succs[1].Block():
			p = s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
		default:
			if b.Likely != ssa.BranchUnlikely {
				p = s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
				s.Br(psess.gc, obj.AJMP, b.Succs[1].Block())
			} else {
				p = s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
				s.Br(psess.gc, obj.AJMP, b.Succs[0].Block())
			}
		}
		if !b.Control.Type.IsFlags(psess.types) {
			p.From.Type = obj.TYPE_REG
			p.From.Reg = b.Control.Reg(psess.ssa)
		}
	default:
		b.Fatalf("branch not implemented: %s. Control: %s", b.LongString(psess.ssa), b.Control.LongString(psess.ssa))
	}
}

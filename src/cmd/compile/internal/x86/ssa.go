package x86

import (
	"fmt"
	"math"

	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/x86"
)

// markMoves marks any MOVXconst ops that need to avoid clobbering flags.
func (psess *PackageSession) ssaMarkMoves(s *gc.SSAGenState, b *ssa.Block) {
	flive := b.FlagsLiveAtEnd
	if b.Control != nil && b.Control.Type.IsFlags(psess.types) {
		flive = true
	}
	for i := len(b.Values) - 1; i >= 0; i-- {
		v := b.Values[i]
		if flive && v.Op == ssa.Op386MOVLconst {

			v.Aux = v
		}
		if v.Type.IsFlags(psess.types) {
			flive = false
		}
		for _, a := range v.Args {
			if a.Type.IsFlags(psess.types) {
				flive = true
			}
		}
	}
}

// loadByType returns the load instruction of the given type.
func (psess *PackageSession) loadByType(t *types.Type) obj.As {

	if !t.IsFloat() && t.Size(psess.types) <= 2 {
		if t.Size(psess.types) == 1 {
			return x86.AMOVBLZX
		} else {
			return x86.AMOVWLZX
		}
	}

	return psess.storeByType(t)
}

// storeByType returns the store instruction of the given type.
func (psess *PackageSession) storeByType(t *types.Type) obj.As {
	width := t.Size(psess.types)
	if t.IsFloat() {
		switch width {
		case 4:
			return x86.AMOVSS
		case 8:
			return x86.AMOVSD
		}
	} else {
		switch width {
		case 1:
			return x86.AMOVB
		case 2:
			return x86.AMOVW
		case 4:
			return x86.AMOVL
		}
	}
	panic("bad store type")
}

// moveByType returns the reg->reg move instruction of the given type.
func (psess *PackageSession) moveByType(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(psess.types) {
		case 4:
			return x86.AMOVSS
		case 8:
			return x86.AMOVSD
		default:
			panic(fmt.Sprintf("bad float register width %d:%s", t.Size(psess.types), t))
		}
	} else {
		switch t.Size(psess.types) {
		case 1:

			return x86.AMOVL
		case 2:
			return x86.AMOVL
		case 4:
			return x86.AMOVL
		default:
			panic(fmt.Sprintf("bad int register width %d:%s", t.Size(psess.types), t))
		}
	}
}

// opregreg emits instructions for
//     dest := dest(To) op src(From)
// and also returns the created obj.Prog so it
// may be further adjusted (offset, scale, etc).
func (psess *PackageSession) opregreg(s *gc.SSAGenState, op obj.As, dest, src int16) *obj.Prog {
	p := s.Prog(psess.gc, op)
	p.From.Type = obj.TYPE_REG
	p.To.Type = obj.TYPE_REG
	p.To.Reg = dest
	p.From.Reg = src
	return p
}

func (psess *PackageSession) ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.Op386ADDL:
		r := v.Reg(psess.ssa)
		r1 := v.Args[0].Reg(psess.ssa)
		r2 := v.Args[1].Reg(psess.ssa)
		switch {
		case r == r1:
			p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
			p.From.Type = obj.TYPE_REG
			p.From.Reg = r2
			p.To.Type = obj.TYPE_REG
			p.To.Reg = r
		case r == r2:
			p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
			p.From.Type = obj.TYPE_REG
			p.From.Reg = r1
			p.To.Type = obj.TYPE_REG
			p.To.Reg = r
		default:
			p := s.Prog(psess.gc, x86.ALEAL)
			p.From.Type = obj.TYPE_MEM
			p.From.Reg = r1
			p.From.Scale = 1
			p.From.Index = r2
			p.To.Type = obj.TYPE_REG
			p.To.Reg = r
		}

	case ssa.Op386SUBL,
		ssa.Op386MULL,
		ssa.Op386ANDL,
		ssa.Op386ORL,
		ssa.Op386XORL,
		ssa.Op386SHLL,
		ssa.Op386SHRL, ssa.Op386SHRW, ssa.Op386SHRB,
		ssa.Op386SARL, ssa.Op386SARW, ssa.Op386SARB,
		ssa.Op386ADDSS, ssa.Op386ADDSD, ssa.Op386SUBSS, ssa.Op386SUBSD,
		ssa.Op386MULSS, ssa.Op386MULSD, ssa.Op386DIVSS, ssa.Op386DIVSD,
		ssa.Op386PXOR,
		ssa.Op386ADCL,
		ssa.Op386SBBL:
		r := v.Reg(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		psess.
			opregreg(s, v.Op.Asm(psess.ssa), r, v.Args[1].Reg(psess.ssa))

	case ssa.Op386ADDLcarry, ssa.Op386SUBLcarry:

		r := v.Reg0(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output[0] not in same register %s", v.LongString(psess.ssa))
		}
		psess.
			opregreg(s, v.Op.Asm(psess.ssa), r, v.Args[1].Reg(psess.ssa))

	case ssa.Op386ADDLconstcarry, ssa.Op386SUBLconstcarry:

		r := v.Reg0(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output[0] not in same register %s", v.LongString(psess.ssa))
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

	case ssa.Op386DIVL, ssa.Op386DIVW,
		ssa.Op386DIVLU, ssa.Op386DIVWU,
		ssa.Op386MODL, ssa.Op386MODW,
		ssa.Op386MODLU, ssa.Op386MODWU:

		x := v.Args[1].Reg(psess.

			// CPU faults upon signed overflow, which occurs when most
			// negative int is divided by -1.
			ssa)

		var j *obj.Prog
		if v.Op == ssa.Op386DIVL || v.Op == ssa.Op386DIVW ||
			v.Op == ssa.Op386MODL || v.Op == ssa.Op386MODW {

			var c *obj.Prog
			switch v.Op {
			case ssa.Op386DIVL, ssa.Op386MODL:
				c = s.Prog(psess.gc, x86.ACMPL)
				j = s.Prog(psess.gc, x86.AJEQ)
				s.Prog(psess.gc, x86.ACDQ)

			case ssa.Op386DIVW, ssa.Op386MODW:
				c = s.Prog(psess.gc, x86.ACMPW)
				j = s.Prog(psess.gc, x86.AJEQ)
				s.Prog(psess.gc, x86.ACWD)
			}
			c.From.Type = obj.TYPE_REG
			c.From.Reg = x
			c.To.Type = obj.TYPE_CONST
			c.To.Offset = -1

			j.To.Type = obj.TYPE_BRANCH
		}

		if v.Op == ssa.Op386DIVLU || v.Op == ssa.Op386MODLU ||
			v.Op == ssa.Op386DIVWU || v.Op == ssa.Op386MODWU {
			c := s.Prog(psess.gc, x86.AXORL)
			c.From.Type = obj.TYPE_REG
			c.From.Reg = x86.REG_DX
			c.To.Type = obj.TYPE_REG
			c.To.Reg = x86.REG_DX
		}

		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x

		if j != nil {
			j2 := s.Prog(psess.gc, obj.AJMP)
			j2.To.Type = obj.TYPE_BRANCH

			var n *obj.Prog
			if v.Op == ssa.Op386DIVL || v.Op == ssa.Op386DIVW {

				n = s.Prog(psess.gc, x86.ANEGL)
				n.To.Type = obj.TYPE_REG
				n.To.Reg = x86.REG_AX
			} else {

				n = s.Prog(psess.gc, x86.AXORL)
				n.From.Type = obj.TYPE_REG
				n.From.Reg = x86.REG_DX
				n.To.Type = obj.TYPE_REG
				n.To.Reg = x86.REG_DX
			}

			j.To.Val = n
			j2.To.Val = s.Pc()
		}

	case ssa.Op386HMULL, ssa.Op386HMULLU:

		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)

		if v.Type.Size(psess.types) == 1 {
			m := s.Prog(psess.gc, x86.AMOVB)
			m.From.Type = obj.TYPE_REG
			m.From.Reg = x86.REG_AH
			m.To.Type = obj.TYPE_REG
			m.To.Reg = x86.REG_DX
		}

	case ssa.Op386MULLQU:

		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)

	case ssa.Op386AVGLU:

		r := v.Reg(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		p := s.Prog(psess.gc, x86.AADDL)
		p.From.Type = obj.TYPE_REG
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p = s.Prog(psess.gc, x86.ARCRL)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

	case ssa.Op386ADDLconst:
		r := v.Reg(psess.ssa)
		a := v.Args[0].Reg(psess.ssa)
		if r == a {
			if v.AuxInt == 1 {
				p := s.Prog(psess.gc, x86.AINCL)
				p.To.Type = obj.TYPE_REG
				p.To.Reg = r
				return
			}
			if v.AuxInt == -1 {
				p := s.Prog(psess.gc, x86.ADECL)
				p.To.Type = obj.TYPE_REG
				p.To.Reg = r
				return
			}
			p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = v.AuxInt
			p.To.Type = obj.TYPE_REG
			p.To.Reg = r
			return
		}
		p := s.Prog(psess.gc, x86.ALEAL)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = a
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

	case ssa.Op386MULLconst:
		r := v.Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		p.SetFrom3(obj.Addr{Type: obj.TYPE_REG, Reg: v.Args[0].Reg(psess.ssa)})

	case ssa.Op386SUBLconst,
		ssa.Op386ADCLconst,
		ssa.Op386SBBLconst,
		ssa.Op386ANDLconst,
		ssa.Op386ORLconst,
		ssa.Op386XORLconst,
		ssa.Op386SHLLconst,
		ssa.Op386SHRLconst, ssa.Op386SHRWconst, ssa.Op386SHRBconst,
		ssa.Op386SARLconst, ssa.Op386SARWconst, ssa.Op386SARBconst,
		ssa.Op386ROLLconst, ssa.Op386ROLWconst, ssa.Op386ROLBconst:
		r := v.Reg(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.Op386SBBLcarrymask:
		r := v.Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = r
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.Op386LEAL1, ssa.Op386LEAL2, ssa.Op386LEAL4, ssa.Op386LEAL8:
		r := v.Args[0].Reg(psess.ssa)
		i := v.Args[1].Reg(psess.ssa)
		p := s.Prog(psess.gc, x86.ALEAL)
		switch v.Op {
		case ssa.Op386LEAL1:
			p.From.Scale = 1
			if i == x86.REG_SP {
				r, i = i, r
			}
		case ssa.Op386LEAL2:
			p.From.Scale = 2
		case ssa.Op386LEAL4:
			p.From.Scale = 4
		case ssa.Op386LEAL8:
			p.From.Scale = 8
		}
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r
		p.From.Index = i
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.Op386LEAL:
		p := s.Prog(psess.gc, x86.ALEAL)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.Op386CMPL, ssa.Op386CMPW, ssa.Op386CMPB,
		ssa.Op386TESTL, ssa.Op386TESTW, ssa.Op386TESTB:
		psess.
			opregreg(s, v.Op.Asm(psess.ssa), v.Args[1].Reg(psess.ssa), v.Args[0].Reg(psess.ssa))
	case ssa.Op386UCOMISS, ssa.Op386UCOMISD:
		psess.
			opregreg(s, v.Op.Asm(psess.ssa), v.Args[0].Reg(psess.ssa), v.Args[1].Reg(psess.ssa))
	case ssa.Op386CMPLconst, ssa.Op386CMPWconst, ssa.Op386CMPBconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = v.AuxInt
	case ssa.Op386TESTLconst, ssa.Op386TESTWconst, ssa.Op386TESTBconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Args[0].Reg(psess.ssa)
	case ssa.Op386MOVLconst:
		x := v.Reg(psess.ssa)

		if v.AuxInt == 0 && v.Aux == nil {
			p := s.Prog(psess.gc, x86.AXORL)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = x
			p.To.Type = obj.TYPE_REG
			p.To.Reg = x
			break
		}

		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x
	case ssa.Op386MOVSSconst, ssa.Op386MOVSDconst:
		x := v.Reg(psess.ssa)
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_FCONST
		p.From.Val = math.Float64frombits(uint64(v.AuxInt))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x
	case ssa.Op386MOVSSconst1, ssa.Op386MOVSDconst1:
		p := s.Prog(psess.gc, x86.ALEAL)
		p.From.Type = obj.TYPE_MEM
		p.From.Name = obj.NAME_EXTERN
		f := math.Float64frombits(uint64(v.AuxInt))
		if v.Op == ssa.Op386MOVSDconst1 {
			p.From.Sym = psess.gc.Ctxt.Float64Sym(f)
		} else {
			p.From.Sym = psess.gc.Ctxt.Float32Sym(float32(f))
		}
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.Op386MOVSSconst2, ssa.Op386MOVSDconst2:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.Op386MOVSSload, ssa.Op386MOVSDload, ssa.Op386MOVLload, ssa.Op386MOVWload, ssa.Op386MOVBload, ssa.Op386MOVBLSXload, ssa.Op386MOVWLSXload:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.Op386MOVSDloadidx8:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.From.Scale = 8
		p.From.Index = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.Op386MOVLloadidx4, ssa.Op386MOVSSloadidx4:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.From.Scale = 4
		p.From.Index = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.Op386MOVWloadidx2:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.From.Scale = 2
		p.From.Index = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.Op386MOVBloadidx1, ssa.Op386MOVWloadidx1, ssa.Op386MOVLloadidx1, ssa.Op386MOVSSloadidx1, ssa.Op386MOVSDloadidx1:
		r := v.Args[0].Reg(psess.ssa)
		i := v.Args[1].Reg(psess.ssa)
		if i == x86.REG_SP {
			r, i = i, r
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = r
		p.From.Scale = 1
		p.From.Index = i
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.Op386ADDLload, ssa.Op386SUBLload, ssa.Op386ANDLload, ssa.Op386ORLload, ssa.Op386XORLload,
		ssa.Op386ADDSDload, ssa.Op386ADDSSload, ssa.Op386SUBSDload, ssa.Op386SUBSSload, ssa.Op386MULSDload, ssa.Op386MULSSload:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
		if v.Reg(psess.ssa) != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
	case ssa.Op386MOVSSstore, ssa.Op386MOVSDstore, ssa.Op386MOVLstore, ssa.Op386MOVWstore, ssa.Op386MOVBstore,
		ssa.Op386ADDLmodify, ssa.Op386SUBLmodify, ssa.Op386ANDLmodify, ssa.Op386ORLmodify, ssa.Op386XORLmodify:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.Op386MOVSDstoreidx8:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[2].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Scale = 8
		p.To.Index = v.Args[1].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.Op386MOVSSstoreidx4, ssa.Op386MOVLstoreidx4:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[2].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Scale = 4
		p.To.Index = v.Args[1].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.Op386MOVWstoreidx2:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[2].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Scale = 2
		p.To.Index = v.Args[1].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
	case ssa.Op386MOVBstoreidx1, ssa.Op386MOVWstoreidx1, ssa.Op386MOVLstoreidx1, ssa.Op386MOVSSstoreidx1, ssa.Op386MOVSDstoreidx1:
		r := v.Args[0].Reg(psess.ssa)
		i := v.Args[1].Reg(psess.ssa)
		if i == x86.REG_SP {
			r, i = i, r
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[2].Reg(psess.ssa)
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = r
		p.To.Scale = 1
		p.To.Index = i
		psess.gc.
			AddAux(&p.To, v)
	case ssa.Op386MOVLstoreconst, ssa.Op386MOVWstoreconst, ssa.Op386MOVBstoreconst:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		sc := v.AuxValAndOff(psess.ssa)
		p.From.Offset = sc.Val()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux2(&p.To, v, sc.Off())
	case ssa.Op386MOVLstoreconstidx1, ssa.Op386MOVLstoreconstidx4, ssa.Op386MOVWstoreconstidx1, ssa.Op386MOVWstoreconstidx2, ssa.Op386MOVBstoreconstidx1:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_CONST
		sc := v.AuxValAndOff(psess.ssa)
		p.From.Offset = sc.Val()
		r := v.Args[0].Reg(psess.ssa)
		i := v.Args[1].Reg(psess.ssa)
		switch v.Op {
		case ssa.Op386MOVBstoreconstidx1, ssa.Op386MOVWstoreconstidx1, ssa.Op386MOVLstoreconstidx1:
			p.To.Scale = 1
			if i == x86.REG_SP {
				r, i = i, r
			}
		case ssa.Op386MOVWstoreconstidx2:
			p.To.Scale = 2
		case ssa.Op386MOVLstoreconstidx4:
			p.To.Scale = 4
		}
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = r
		p.To.Index = i
		psess.gc.
			AddAux2(&p.To, v, sc.Off())
	case ssa.Op386MOVWLSX, ssa.Op386MOVBLSX, ssa.Op386MOVWLZX, ssa.Op386MOVBLZX,
		ssa.Op386CVTSL2SS, ssa.Op386CVTSL2SD,
		ssa.Op386CVTTSS2SL, ssa.Op386CVTTSD2SL,
		ssa.Op386CVTSS2SD, ssa.Op386CVTSD2SS:
		psess.
			opregreg(s, v.Op.Asm(psess.ssa), v.Reg(psess.ssa), v.Args[0].Reg(psess.ssa))
	case ssa.Op386DUFFZERO:
		p := s.Prog(psess.gc, obj.ADUFFZERO)
		p.To.Type = obj.TYPE_ADDR
		p.To.Sym = psess.gc.Duffzero
		p.To.Offset = v.AuxInt
	case ssa.Op386DUFFCOPY:
		p := s.Prog(psess.gc, obj.ADUFFCOPY)
		p.To.Type = obj.TYPE_ADDR
		p.To.Sym = psess.gc.Duffcopy
		p.To.Offset = v.AuxInt

	case ssa.OpCopy:
		if v.Type.IsMemory(psess.types) {
			return
		}
		x := v.Args[0].Reg(psess.ssa)
		y := v.Reg(psess.ssa)
		if x != y {
			psess.
				opregreg(s, psess.moveByType(v.Type), y, x)
		}
	case ssa.OpLoadReg:
		if v.Type.IsFlags(psess.types) {
			v.Fatalf("load flags not implemented: %v", v.LongString(psess.ssa))
			return
		}
		p := s.Prog(psess.gc, psess.loadByType(v.Type))
		psess.gc.
			AddrAuto(&p.From, v.Args[0])
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.OpStoreReg:
		if v.Type.IsFlags(psess.types) {
			v.Fatalf("store flags not implemented: %v", v.LongString(psess.ssa))
			return
		}
		p := s.Prog(psess.gc, psess.storeByType(v.Type))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddrAuto(&p.To, v)
	case ssa.Op386LoweredGetClosurePtr:
		psess.gc.
			CheckLoweredGetClosurePtr(v)
	case ssa.Op386LoweredGetG:
		r := v.Reg(psess.ssa)

		if psess.x86.CanUse1InsnTLS(psess.gc.Ctxt) {

			p := s.Prog(psess.gc, x86.AMOVL)
			p.From.Type = obj.TYPE_MEM
			p.From.Reg = x86.REG_TLS
			p.To.Type = obj.TYPE_REG
			p.To.Reg = r
		} else {

			p := s.Prog(psess.gc, x86.AMOVL)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = x86.REG_TLS
			p.To.Type = obj.TYPE_REG
			p.To.Reg = r
			q := s.Prog(psess.gc, x86.AMOVL)
			q.From.Type = obj.TYPE_MEM
			q.From.Reg = r
			q.From.Index = x86.REG_TLS
			q.From.Scale = 1
			q.To.Type = obj.TYPE_REG
			q.To.Reg = r
		}

	case ssa.Op386LoweredGetCallerPC:
		p := s.Prog(psess.gc, x86.AMOVL)
		p.From.Type = obj.TYPE_MEM
		p.From.Offset = -4
		p.From.Name = obj.NAME_PARAM
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.Op386LoweredGetCallerSP:

		p := s.Prog(psess.gc, x86.AMOVL)
		p.From.Type = obj.TYPE_ADDR
		p.From.Offset = -psess.gc.Ctxt.FixedFrameSize()
		p.From.Name = obj.NAME_PARAM
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.Op386LoweredWB:
		p := s.Prog(psess.gc, obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = v.Aux.(*obj.LSym)

	case ssa.Op386CALLstatic, ssa.Op386CALLclosure, ssa.Op386CALLinter:
		s.Call(psess.gc, v)
	case ssa.Op386NEGL,
		ssa.Op386BSWAPL,
		ssa.Op386NOTL:
		r := v.Reg(psess.ssa)
		if r != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
	case ssa.Op386BSFL, ssa.Op386BSFW,
		ssa.Op386BSRL, ssa.Op386BSRW,
		ssa.Op386SQRTSD:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
	case ssa.Op386SETEQ, ssa.Op386SETNE,
		ssa.Op386SETL, ssa.Op386SETLE,
		ssa.Op386SETG, ssa.Op386SETGE,
		ssa.Op386SETGF, ssa.Op386SETGEF,
		ssa.Op386SETB, ssa.Op386SETBE,
		ssa.Op386SETORD, ssa.Op386SETNAN,
		ssa.Op386SETA, ssa.Op386SETAE:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

	case ssa.Op386SETNEF:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
		q := s.Prog(psess.gc, x86.ASETPS)
		q.To.Type = obj.TYPE_REG
		q.To.Reg = x86.REG_AX
		psess.
			opregreg(s, x86.AORL, v.Reg(psess.ssa), x86.REG_AX)

	case ssa.Op386SETEQF:
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)
		q := s.Prog(psess.gc, x86.ASETPC)
		q.To.Type = obj.TYPE_REG
		q.To.Reg = x86.REG_AX
		psess.
			opregreg(s, x86.AANDL, v.Reg(psess.ssa), x86.REG_AX)

	case ssa.Op386InvertFlags:
		v.Fatalf("InvertFlags should never make it to codegen %v", v.LongString(psess.ssa))
	case ssa.Op386FlagEQ, ssa.Op386FlagLT_ULT, ssa.Op386FlagLT_UGT, ssa.Op386FlagGT_ULT, ssa.Op386FlagGT_UGT:
		v.Fatalf("Flag* ops should never make it to codegen %v", v.LongString(psess.ssa))
	case ssa.Op386REPSTOSL:
		s.Prog(psess.gc, x86.AREP)
		s.Prog(psess.gc, x86.ASTOSL)
	case ssa.Op386REPMOVSL:
		s.Prog(psess.gc, x86.AREP)
		s.Prog(psess.gc, x86.AMOVSL)
	case ssa.Op386LoweredNilCheck:

		p := s.Prog(psess.gc, x86.ATESTB)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_AX
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
		if psess.gc.Debug_checknil != 0 && v.Pos.Line() > 1 {
			psess.gc.
				Warnl(v.Pos, "generated nil check")
		}
	case ssa.Op386FCHS:
		v.Fatalf("FCHS in non-387 mode")
	case ssa.OpClobber:
		p := s.Prog(psess.gc, x86.AMOVL)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 0xdeaddead
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = x86.REG_SP
		psess.gc.
			AddAux(&p.To, v)
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

		p := s.Prog(psess.gc, x86.ATESTL)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_AX
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_AX
		p = s.Prog(psess.gc, x86.AJNE)
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
		p := s.Prog(psess.gc, obj.AJMP)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = b.Aux.(*obj.LSym)

	case ssa.Block386EQF:
		s.FPJump(psess.gc, b, next, &psess.eqfJumps)

	case ssa.Block386NEF:
		s.FPJump(psess.gc, b, next, &psess.nefJumps)

	case ssa.Block386EQ, ssa.Block386NE,
		ssa.Block386LT, ssa.Block386GE,
		ssa.Block386LE, ssa.Block386GT,
		ssa.Block386ULT, ssa.Block386UGT,
		ssa.Block386ULE, ssa.Block386UGE:
		jmp := psess.blockJump[b.Kind]
		switch next {
		case b.Succs[0].Block():
			s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
		case b.Succs[1].Block():
			s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
		default:
			if b.Likely != ssa.BranchUnlikely {
				s.Br(psess.gc, jmp.asm, b.Succs[0].Block())
				s.Br(psess.gc, obj.AJMP, b.Succs[1].Block())
			} else {
				s.Br(psess.gc, jmp.invasm, b.Succs[1].Block())
				s.Br(psess.gc, obj.AJMP, b.Succs[0].Block())
			}
		}
	default:
		b.Fatalf("branch not implemented: %s. Control: %s", b.LongString(psess.ssa), b.Control.LongString(psess.ssa))
	}
}

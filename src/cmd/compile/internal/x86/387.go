package x86

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/x86"
	"math"
)

// Generates code for v using 387 instructions.
func (psess *PackageSession) ssaGenValue387(s *gc.SSAGenState, v *ssa.Value) {

	switch v.Op {
	case ssa.Op386MOVSSconst, ssa.Op386MOVSDconst:
		p := s.Prog(psess.gc, psess.loadPush(v.Type))
		p.From.Type = obj.TYPE_FCONST
		p.From.Val = math.Float64frombits(uint64(v.AuxInt))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		psess.
			popAndSave(s, v)

	case ssa.Op386MOVSSconst2, ssa.Op386MOVSDconst2:
		p := s.Prog(psess.gc, psess.loadPush(v.Type))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		psess.
			popAndSave(s, v)

	case ssa.Op386MOVSSload, ssa.Op386MOVSDload, ssa.Op386MOVSSloadidx1, ssa.Op386MOVSDloadidx1, ssa.Op386MOVSSloadidx4, ssa.Op386MOVSDloadidx8:
		p := s.Prog(psess.gc, psess.loadPush(v.Type))
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.From, v)
		switch v.Op {
		case ssa.Op386MOVSSloadidx1, ssa.Op386MOVSDloadidx1:
			p.From.Scale = 1
			p.From.Index = v.Args[1].Reg(psess.ssa)
			if p.From.Index == x86.REG_SP {
				p.From.Reg, p.From.Index = p.From.Index, p.From.Reg
			}
		case ssa.Op386MOVSSloadidx4:
			p.From.Scale = 4
			p.From.Index = v.Args[1].Reg(psess.ssa)
		case ssa.Op386MOVSDloadidx8:
			p.From.Scale = 8
			p.From.Index = v.Args[1].Reg(psess.ssa)
		}
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		psess.
			popAndSave(s, v)

	case ssa.Op386MOVSSstore, ssa.Op386MOVSDstore:
		psess.
			push(s, v.Args[1])

		// Pop and store value.
		var op obj.As
		switch v.Op {
		case ssa.Op386MOVSSstore:
			op = x86.AFMOVFP
		case ssa.Op386MOVSDstore:
			op = x86.AFMOVDP
		}
		p := s.Prog(psess.gc, op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)

	case ssa.Op386MOVSSstoreidx1, ssa.Op386MOVSDstoreidx1, ssa.Op386MOVSSstoreidx4, ssa.Op386MOVSDstoreidx8:
		psess.
			push(s, v.Args[2])
		var op obj.As
		switch v.Op {
		case ssa.Op386MOVSSstoreidx1, ssa.Op386MOVSSstoreidx4:
			op = x86.AFMOVFP
		case ssa.Op386MOVSDstoreidx1, ssa.Op386MOVSDstoreidx8:
			op = x86.AFMOVDP
		}
		p := s.Prog(psess.gc, op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg(psess.ssa)
		psess.gc.
			AddAux(&p.To, v)
		switch v.Op {
		case ssa.Op386MOVSSstoreidx1, ssa.Op386MOVSDstoreidx1:
			p.To.Scale = 1
			p.To.Index = v.Args[1].Reg(psess.ssa)
			if p.To.Index == x86.REG_SP {
				p.To.Reg, p.To.Index = p.To.Index, p.To.Reg
			}
		case ssa.Op386MOVSSstoreidx4:
			p.To.Scale = 4
			p.To.Index = v.Args[1].Reg(psess.ssa)
		case ssa.Op386MOVSDstoreidx8:
			p.To.Scale = 8
			p.To.Index = v.Args[1].Reg(psess.ssa)
		}

	case ssa.Op386ADDSS, ssa.Op386ADDSD, ssa.Op386SUBSS, ssa.Op386SUBSD,
		ssa.Op386MULSS, ssa.Op386MULSD, ssa.Op386DIVSS, ssa.Op386DIVSD:
		if v.Reg(psess.ssa) != v.Args[0].Reg(psess.ssa) {
			v.Fatalf("input[0] and output not in same register %s", v.LongString(psess.ssa))
		}
		psess.
			push(s, v.Args[1])

		switch v.Op {
		case ssa.Op386ADDSS, ssa.Op386SUBSS, ssa.Op386MULSS, ssa.Op386DIVSS:
			p := s.Prog(psess.gc, x86.AFSTCW)
			s.AddrScratch(psess.gc, &p.To)
			p = s.Prog(psess.gc, x86.AFLDCW)
			p.From.Type = obj.TYPE_MEM
			p.From.Name = obj.NAME_EXTERN
			p.From.Sym = psess.gc.ControlWord32
		}

		var op obj.As
		switch v.Op {
		case ssa.Op386ADDSS, ssa.Op386ADDSD:
			op = x86.AFADDDP
		case ssa.Op386SUBSS, ssa.Op386SUBSD:
			op = x86.AFSUBDP
		case ssa.Op386MULSS, ssa.Op386MULSD:
			op = x86.AFMULDP
		case ssa.Op386DIVSS, ssa.Op386DIVSD:
			op = x86.AFDIVDP
		}
		p := s.Prog(psess.gc, op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = s.SSEto387[v.Reg(psess.ssa)] + 1

		switch v.Op {
		case ssa.Op386ADDSS, ssa.Op386SUBSS, ssa.Op386MULSS, ssa.Op386DIVSS:
			p := s.Prog(psess.gc, x86.AFLDCW)
			s.AddrScratch(psess.gc, &p.From)
		}

	case ssa.Op386UCOMISS, ssa.Op386UCOMISD:
		psess.
			push(s, v.Args[0])

		p := s.Prog(psess.gc, x86.AFUCOMP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = s.SSEto387[v.Args[1].Reg(psess.ssa)] + 1

		p = s.Prog(psess.gc, x86.AMOVL)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_AX
		s.AddrScratch(psess.gc, &p.To)

		p = s.Prog(psess.gc, x86.AFSTSW)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_AX

		s.Prog(psess.gc, x86.ASAHF)

		p = s.Prog(psess.gc, x86.AMOVL)
		s.AddrScratch(psess.gc, &p.From)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_AX

	case ssa.Op386SQRTSD:
		psess.
			push(s, v.Args[0])
		s.Prog(psess.gc, x86.AFSQRT)
		psess.
			popAndSave(s, v)

	case ssa.Op386FCHS:
		psess.
			push(s, v.Args[0])
		s.Prog(psess.gc, x86.AFCHS)
		psess.
			popAndSave(s, v)

	case ssa.Op386CVTSL2SS, ssa.Op386CVTSL2SD:
		p := s.Prog(psess.gc, x86.AMOVL)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg(psess.ssa)
		s.AddrScratch(psess.gc, &p.To)
		p = s.Prog(psess.gc, x86.AFMOVL)
		s.AddrScratch(psess.gc, &p.From)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		psess.
			popAndSave(s, v)

	case ssa.Op386CVTTSD2SL, ssa.Op386CVTTSS2SL:
		psess.
			push(s, v.Args[0])

		p := s.Prog(psess.gc, x86.AFSTCW)
		s.AddrScratch(psess.gc, &p.To)
		p.To.Offset += 4

		p = s.Prog(psess.gc, x86.AFLDCW)
		p.From.Type = obj.TYPE_MEM
		p.From.Name = obj.NAME_EXTERN
		p.From.Sym = psess.gc.ControlWord64trunc

		p = s.Prog(psess.gc, x86.AFMOVLP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		s.AddrScratch(psess.gc, &p.To)
		p = s.Prog(psess.gc, x86.AMOVL)
		s.AddrScratch(psess.gc, &p.From)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg(psess.ssa)

		p = s.Prog(psess.gc, x86.AFLDCW)
		s.AddrScratch(psess.gc, &p.From)
		p.From.Offset += 4

	case ssa.Op386CVTSS2SD:
		psess.
			push(s, v.Args[0])
		psess.
			popAndSave(s, v)

	case ssa.Op386CVTSD2SS:
		psess.
			push(s, v.Args[0])
		p := s.Prog(psess.gc, x86.AFMOVFP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		s.AddrScratch(psess.gc, &p.To)
		p = s.Prog(psess.gc, x86.AFMOVF)
		s.AddrScratch(psess.gc, &p.From)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		psess.
			popAndSave(s, v)

	case ssa.OpLoadReg:
		if !v.Type.IsFloat() {
			psess.
				ssaGenValue(s, v)
			return
		}

		p := s.Prog(psess.gc, psess.loadPush(v.Type))
		psess.gc.
			AddrAuto(&p.From, v.Args[0])
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		psess.
			popAndSave(s, v)

	case ssa.OpStoreReg:
		if !v.Type.IsFloat() {
			psess.
				ssaGenValue(s, v)
			return
		}
		psess.
			push(s, v.Args[0])
		var op obj.As
		switch v.Type.Size(psess.types) {
		case 4:
			op = x86.AFMOVFP
		case 8:
			op = x86.AFMOVDP
		}
		p := s.Prog(psess.gc, op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		psess.gc.
			AddrAuto(&p.To, v)

	case ssa.OpCopy:
		if !v.Type.IsFloat() {
			psess.
				ssaGenValue(s, v)
			return
		}
		psess.
			push(s, v.Args[0])
		psess.
			popAndSave(s, v)

	case ssa.Op386CALLstatic, ssa.Op386CALLclosure, ssa.Op386CALLinter:
		psess.
			flush387(s)
		fallthrough
	default:
		psess.
			ssaGenValue(s, v)
	}
}

// push pushes v onto the floating-point stack.  v must be in a register.
func (psess *PackageSession) push(s *gc.SSAGenState, v *ssa.Value) {
	p := s.Prog(psess.gc, x86.AFMOVD)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = s.SSEto387[v.Reg(psess.ssa)]
	p.To.Type = obj.TYPE_REG
	p.To.Reg = x86.REG_F0
}

// popAndSave pops a value off of the floating-point stack and stores
// it in the reigster assigned to v.
func (psess *PackageSession) popAndSave(s *gc.SSAGenState, v *ssa.Value) {
	r := v.Reg(psess.ssa)
	if _, ok := s.SSEto387[r]; ok {

		p := s.Prog(psess.gc, x86.AFMOVDP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = s.SSEto387[v.Reg(psess.ssa)] + 1
	} else {

		for rSSE, r387 := range s.SSEto387 {
			s.SSEto387[rSSE] = r387 + 1
		}
		s.SSEto387[r] = x86.REG_F0
	}
}

// loadPush returns the opcode for load+push of the given type.
func (psess *PackageSession) loadPush(t *types.Type) obj.As {
	if t.Size(psess.types) == 4 {
		return x86.AFMOVF
	}
	return x86.AFMOVD
}

// flush387 removes all entries from the 387 floating-point stack.
func (psess *PackageSession) flush387(s *gc.SSAGenState) {
	for k := range s.SSEto387 {
		p := s.Prog(psess.gc, x86.AFMOVDP)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x86.REG_F0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = x86.REG_F0
		delete(s.SSEto387, k)
	}
}

func (psess *PackageSession) ssaGenBlock387(s *gc.SSAGenState, b, next *ssa.Block) {
	psess.
		flush387(s)
	psess.
		ssaGenBlock(s, b, next)
}

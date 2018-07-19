package wasm

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/wasm"
)

func (psess *PackageSession) Init(arch *gc.Arch) {
	arch.LinkArch = &psess.wasm.Linkwasm
	arch.REGSP = wasm.REG_SP
	arch.MAXWIDTH = 1 << 50

	arch.ZeroRange = psess.zeroRange
	arch.ZeroAuto = psess.zeroAuto
	arch.Ginsnop = psess.ginsnop

	arch.SSAMarkMoves = ssaMarkMoves
	arch.SSAGenValue = psess.ssaGenValue
	arch.SSAGenBlock = psess.ssaGenBlock
}

func (psess *PackageSession) zeroRange(pp *gc.Progs, p *obj.Prog, off, cnt int64, state *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	if cnt%8 != 0 {
		psess.gc.
			Fatalf("zerorange count not a multiple of widthptr %d", cnt)
	}

	for i := int64(0); i < cnt; i += 8 {
		p = pp.Appendpp(psess.gc, p, wasm.AGet, obj.TYPE_REG, wasm.REG_SP, 0, 0, 0, 0)
		p = pp.Appendpp(psess.gc, p, wasm.AI64Const, obj.TYPE_CONST, 0, 0, 0, 0, 0)
		p = pp.Appendpp(psess.gc, p, wasm.AI64Store, 0, 0, 0, obj.TYPE_CONST, 0, off+i)
	}

	return p
}

func (psess *PackageSession) zeroAuto(pp *gc.Progs, n *gc.Node) {
	sym := n.Sym.Linksym(psess.types)
	size := n.Type.Size(psess.types)
	for i := int64(0); i < size; i += 8 {
		p := pp.Prog(psess.gc, wasm.AGet)
		p.From = obj.Addr{Type: obj.TYPE_REG, Reg: wasm.REG_SP}

		p = pp.Prog(psess.gc, wasm.AI64Const)
		p.From = obj.Addr{Type: obj.TYPE_CONST, Offset: 0}

		p = pp.Prog(psess.gc, wasm.AI64Store)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_AUTO, Offset: n.Xoffset + i, Sym: sym}
	}
}

func (psess *PackageSession) ginsnop(pp *gc.Progs) {
	pp.Prog(psess.gc, wasm.ANop)
}

func ssaMarkMoves(s *gc.SSAGenState, b *ssa.Block) {
}

func (psess *PackageSession) ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	goToBlock := func(block *ssa.Block, canFallthrough bool) {
		if canFallthrough && block == next {
			return
		}
		s.Br(psess.gc, obj.AJMP, block)
	}

	switch b.Kind {
	case ssa.BlockPlain:
		goToBlock(b.Succs[0].Block(), true)

	case ssa.BlockIf:
		psess.
			getValue32(s, b.Control)
		s.Prog(psess.gc, wasm.AI32Eqz)
		s.Prog(psess.gc, wasm.AIf)
		goToBlock(b.Succs[1].Block(), false)
		s.Prog(psess.gc, wasm.AEnd)
		goToBlock(b.Succs[0].Block(), true)

	case ssa.BlockRet:
		s.Prog(psess.gc, obj.ARET)

	case ssa.BlockRetJmp:
		p := s.Prog(psess.gc, obj.ARET)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = b.Aux.(*obj.LSym)

	case ssa.BlockExit:
		s.Prog(psess.gc, obj.AUNDEF)

	case ssa.BlockDefer:
		p := s.Prog(psess.gc, wasm.AGet)
		p.From = obj.Addr{Type: obj.TYPE_REG, Reg: wasm.REG_RET0}
		s.Prog(psess.gc, wasm.AI64Eqz)
		s.Prog(psess.gc, wasm.AI32Eqz)
		s.Prog(psess.gc, wasm.AIf)
		goToBlock(b.Succs[1].Block(), false)
		s.Prog(psess.gc, wasm.AEnd)
		goToBlock(b.Succs[0].Block(), true)

	default:
		panic("unexpected block")
	}

	s.Prog(psess.gc, wasm.ARESUMEPOINT)

	if s.OnWasmStackSkipped != 0 {
		panic("wasm: bad stack")
	}
}

func (psess *PackageSession) ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpWasmLoweredStaticCall, ssa.OpWasmLoweredClosureCall, ssa.OpWasmLoweredInterCall:
		s.PrepareCall(psess.gc, v)
		if v.Aux == psess.gc.Deferreturn {

			s.Prog(psess.gc, wasm.ARESUMEPOINT)
		}
		if v.Op == ssa.OpWasmLoweredClosureCall {
			psess.
				getValue64(s, v.Args[1])
			psess.
				setReg(s, wasm.REG_CTXT)
		}
		if sym, ok := v.Aux.(*obj.LSym); ok {
			p := s.Prog(psess.gc, obj.ACALL)
			p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: sym}
		} else {
			psess.
				getValue64(s, v.Args[0])
			p := s.Prog(psess.gc, obj.ACALL)
			p.To = obj.Addr{Type: obj.TYPE_NONE}
		}

	case ssa.OpWasmLoweredMove:
		psess.
			getValue32(s, v.Args[0])
		psess.
			getValue32(s, v.Args[1])
		psess.
			i32Const(s, int32(v.AuxInt))
		p := s.Prog(psess.gc, wasm.ACall)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: psess.gc.WasmMove}

	case ssa.OpWasmLoweredZero:
		psess.
			getValue32(s, v.Args[0])
		psess.
			i32Const(s, int32(v.AuxInt))
		p := s.Prog(psess.gc, wasm.ACall)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: psess.gc.WasmZero}

	case ssa.OpWasmLoweredNilCheck:
		psess.
			getValue64(s, v.Args[0])
		s.Prog(psess.gc, wasm.AI64Eqz)
		s.Prog(psess.gc, wasm.AIf)
		p := s.Prog(psess.gc, wasm.ACALLNORESUME)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: psess.gc.SigPanic}
		s.Prog(psess.gc, wasm.AEnd)
		if psess.gc.Debug_checknil != 0 && v.Pos.Line() > 1 {
			psess.gc.
				Warnl(v.Pos, "generated nil check")
		}

	case ssa.OpWasmLoweredWB:
		psess.
			getValue64(s, v.Args[0])
		psess.
			getValue64(s, v.Args[1])
		p := s.Prog(psess.gc, wasm.ACALLNORESUME)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: v.Aux.(*obj.LSym)}

	case ssa.OpWasmI64Store8, ssa.OpWasmI64Store16, ssa.OpWasmI64Store32, ssa.OpWasmI64Store, ssa.OpWasmF32Store, ssa.OpWasmF64Store:
		psess.
			getValue32(s, v.Args[0])
		psess.
			getValue64(s, v.Args[1])
		if v.Op == ssa.OpWasmF32Store {
			s.Prog(psess.gc, wasm.AF32DemoteF64)
		}
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.To = obj.Addr{Type: obj.TYPE_CONST, Offset: v.AuxInt}

	case ssa.OpStoreReg:
		psess.
			getReg(s, wasm.REG_SP)
		psess.
			getValue64(s, v.Args[0])
		if v.Type.Etype == types.TFLOAT32 {
			s.Prog(psess.gc, wasm.AF32DemoteF64)
		}
		p := s.Prog(psess.gc, psess.storeOp(v.Type))
		psess.gc.
			AddrAuto(&p.To, v)

	default:
		if v.Type.IsMemory(psess.types) {
			return
		}
		if v.OnWasmStack {
			s.OnWasmStackSkipped++

			return
		}
		psess.
			ssaGenValueOnStack(s, v)
		if s.OnWasmStackSkipped != 0 {
			panic("wasm: bad stack")
		}
		psess.
			setReg(s, v.Reg(psess.ssa))
	}
}

func (psess *PackageSession) ssaGenValueOnStack(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpWasmLoweredGetClosurePtr:
		psess.
			getReg(s, wasm.REG_CTXT)

	case ssa.OpWasmLoweredGetCallerPC:
		p := s.Prog(psess.gc, wasm.AI64Load)

		p.From = obj.Addr{
			Type:   obj.TYPE_MEM,
			Name:   obj.NAME_PARAM,
			Offset: -8,
		}

	case ssa.OpWasmLoweredGetCallerSP:
		p := s.Prog(psess.gc, wasm.AGet)

		p.From = obj.Addr{
			Type:   obj.TYPE_ADDR,
			Name:   obj.NAME_PARAM,
			Reg:    wasm.REG_SP,
			Offset: 0,
		}

	case ssa.OpWasmLoweredAddr:
		p := s.Prog(psess.gc, wasm.AGet)
		switch n := v.Aux.(type) {
		case *obj.LSym:
			p.From = obj.Addr{Type: obj.TYPE_ADDR, Name: obj.NAME_EXTERN, Sym: n}
		case *gc.Node:
			p.From = obj.Addr{
				Type:   obj.TYPE_ADDR,
				Name:   obj.NAME_AUTO,
				Reg:    v.Args[0].Reg(psess.ssa),
				Offset: n.Xoffset,
			}
			if n.Class() == gc.PPARAM || n.Class() == gc.PPARAMOUT {
				p.From.Name = obj.NAME_PARAM
			}
		default:
			panic("wasm: bad LoweredAddr")
		}

	case ssa.OpWasmLoweredRound32F:
		psess.
			getValue64(s, v.Args[0])
		s.Prog(psess.gc, wasm.AF32DemoteF64)
		s.Prog(psess.gc, wasm.AF64PromoteF32)

	case ssa.OpWasmLoweredConvert:
		psess.
			getValue64(s, v.Args[0])

	case ssa.OpWasmSelect:
		psess.
			getValue64(s, v.Args[0])
		psess.
			getValue64(s, v.Args[1])
		psess.
			getValue64(s, v.Args[2])
		s.Prog(psess.gc, wasm.AI32WrapI64)
		s.Prog(psess.gc, v.Op.Asm(psess.ssa))

	case ssa.OpWasmI64AddConst:
		psess.
			getValue64(s, v.Args[0])
		psess.
			i64Const(s, v.AuxInt)
		s.Prog(psess.gc, v.Op.Asm(psess.ssa))

	case ssa.OpWasmI64Const:
		psess.
			i64Const(s, v.AuxInt)

	case ssa.OpWasmF64Const:
		psess.
			f64Const(s, v.AuxFloat(psess.ssa))

	case ssa.OpWasmI64Load8U, ssa.OpWasmI64Load8S, ssa.OpWasmI64Load16U, ssa.OpWasmI64Load16S, ssa.OpWasmI64Load32U, ssa.OpWasmI64Load32S, ssa.OpWasmI64Load, ssa.OpWasmF32Load, ssa.OpWasmF64Load:
		psess.
			getValue32(s, v.Args[0])
		p := s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		p.From = obj.Addr{Type: obj.TYPE_CONST, Offset: v.AuxInt}
		if v.Op == ssa.OpWasmF32Load {
			s.Prog(psess.gc, wasm.AF64PromoteF32)
		}

	case ssa.OpWasmI64Eqz:
		psess.
			getValue64(s, v.Args[0])
		s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		s.Prog(psess.gc, wasm.AI64ExtendUI32)

	case ssa.OpWasmI64Eq, ssa.OpWasmI64Ne, ssa.OpWasmI64LtS, ssa.OpWasmI64LtU, ssa.OpWasmI64GtS, ssa.OpWasmI64GtU, ssa.OpWasmI64LeS, ssa.OpWasmI64LeU, ssa.OpWasmI64GeS, ssa.OpWasmI64GeU, ssa.OpWasmF64Eq, ssa.OpWasmF64Ne, ssa.OpWasmF64Lt, ssa.OpWasmF64Gt, ssa.OpWasmF64Le, ssa.OpWasmF64Ge:
		psess.
			getValue64(s, v.Args[0])
		psess.
			getValue64(s, v.Args[1])
		s.Prog(psess.gc, v.Op.Asm(psess.ssa))
		s.Prog(psess.gc, wasm.AI64ExtendUI32)

	case ssa.OpWasmI64Add, ssa.OpWasmI64Sub, ssa.OpWasmI64Mul, ssa.OpWasmI64DivU, ssa.OpWasmI64RemS, ssa.OpWasmI64RemU, ssa.OpWasmI64And, ssa.OpWasmI64Or, ssa.OpWasmI64Xor, ssa.OpWasmI64Shl, ssa.OpWasmI64ShrS, ssa.OpWasmI64ShrU, ssa.OpWasmF64Add, ssa.OpWasmF64Sub, ssa.OpWasmF64Mul, ssa.OpWasmF64Div:
		psess.
			getValue64(s, v.Args[0])
		psess.
			getValue64(s, v.Args[1])
		s.Prog(psess.gc, v.Op.Asm(psess.ssa))

	case ssa.OpWasmI64DivS:
		psess.
			getValue64(s, v.Args[0])
		psess.
			getValue64(s, v.Args[1])
		if v.Type.Size(psess.types) == 8 {

			p := s.Prog(psess.gc, wasm.ACall)
			p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: psess.gc.WasmDiv}
			break
		}
		s.Prog(psess.gc, wasm.AI64DivS)

	case ssa.OpWasmI64TruncSF64:
		psess.
			getValue64(s, v.Args[0])
		p := s.Prog(psess.gc, wasm.ACall)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: psess.gc.WasmTruncS}

	case ssa.OpWasmI64TruncUF64:
		psess.
			getValue64(s, v.Args[0])
		p := s.Prog(psess.gc, wasm.ACall)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: psess.gc.WasmTruncU}

	case ssa.OpWasmF64Neg, ssa.OpWasmF64ConvertSI64, ssa.OpWasmF64ConvertUI64:
		psess.
			getValue64(s, v.Args[0])
		s.Prog(psess.gc, v.Op.Asm(psess.ssa))

	case ssa.OpLoadReg:
		p := s.Prog(psess.gc, psess.loadOp(v.Type))
		psess.gc.
			AddrAuto(&p.From, v.Args[0])
		if v.Type.Etype == types.TFLOAT32 {
			s.Prog(psess.gc, wasm.AF64PromoteF32)
		}

	case ssa.OpCopy:
		psess.
			getValue64(s, v.Args[0])

	default:
		v.Fatalf("unexpected op: %s", v.Op)

	}
}

func (psess *PackageSession) getValue32(s *gc.SSAGenState, v *ssa.Value) {
	if v.OnWasmStack {
		s.OnWasmStackSkipped--
		psess.
			ssaGenValueOnStack(s, v)
		s.Prog(psess.gc, wasm.AI32WrapI64)
		return
	}

	reg := v.Reg(psess.ssa)
	psess.
		getReg(s, reg)
	if reg != wasm.REG_SP {
		s.Prog(psess.gc, wasm.AI32WrapI64)
	}
}

func (psess *PackageSession) getValue64(s *gc.SSAGenState, v *ssa.Value) {
	if v.OnWasmStack {
		s.OnWasmStackSkipped--
		psess.
			ssaGenValueOnStack(s, v)
		return
	}

	reg := v.Reg(psess.ssa)
	psess.
		getReg(s, reg)
	if reg == wasm.REG_SP {
		s.Prog(psess.gc, wasm.AI64ExtendUI32)
	}
}

func (psess *PackageSession) i32Const(s *gc.SSAGenState, val int32) {
	p := s.Prog(psess.gc, wasm.AI32Const)
	p.From = obj.Addr{Type: obj.TYPE_CONST, Offset: int64(val)}
}

func (psess *PackageSession) i64Const(s *gc.SSAGenState, val int64) {
	p := s.Prog(psess.gc, wasm.AI64Const)
	p.From = obj.Addr{Type: obj.TYPE_CONST, Offset: val}
}

func (psess *PackageSession) f64Const(s *gc.SSAGenState, val float64) {
	p := s.Prog(psess.gc, wasm.AF64Const)
	p.From = obj.Addr{Type: obj.TYPE_FCONST, Val: val}
}

func (psess *PackageSession) getReg(s *gc.SSAGenState, reg int16) {
	p := s.Prog(psess.gc, wasm.AGet)
	p.From = obj.Addr{Type: obj.TYPE_REG, Reg: reg}
}

func (psess *PackageSession) setReg(s *gc.SSAGenState, reg int16) {
	p := s.Prog(psess.gc, wasm.ASet)
	p.To = obj.Addr{Type: obj.TYPE_REG, Reg: reg}
}

func (psess *PackageSession) loadOp(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(psess.types) {
		case 4:
			return wasm.AF32Load
		case 8:
			return wasm.AF64Load
		default:
			panic("bad load type")
		}
	}

	switch t.Size(psess.types) {
	case 1:
		if t.IsSigned() {
			return wasm.AI64Load8S
		}
		return wasm.AI64Load8U
	case 2:
		if t.IsSigned() {
			return wasm.AI64Load16S
		}
		return wasm.AI64Load16U
	case 4:
		if t.IsSigned() {
			return wasm.AI64Load32S
		}
		return wasm.AI64Load32U
	case 8:
		return wasm.AI64Load
	default:
		panic("bad load type")
	}
}

func (psess *PackageSession) storeOp(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(psess.types) {
		case 4:
			return wasm.AF32Store
		case 8:
			return wasm.AF64Store
		default:
			panic("bad store type")
		}
	}

	switch t.Size(psess.types) {
	case 1:
		return wasm.AI64Store8
	case 2:
		return wasm.AI64Store16
	case 4:
		return wasm.AI64Store32
	case 8:
		return wasm.AI64Store
	default:
		panic("bad store type")
	}
}

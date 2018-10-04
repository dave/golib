// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wasm

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/wasm"
)

func (pstate *PackageState) Init(arch *gc.Arch) {
	arch.LinkArch = &pstate.wasm.Linkwasm
	arch.REGSP = wasm.REG_SP
	arch.MAXWIDTH = 1 << 50

	arch.ZeroRange = pstate.zeroRange
	arch.ZeroAuto = pstate.zeroAuto
	arch.Ginsnop = pstate.ginsnop

	arch.SSAMarkMoves = ssaMarkMoves
	arch.SSAGenValue = pstate.ssaGenValue
	arch.SSAGenBlock = pstate.ssaGenBlock
}

func (pstate *PackageState) zeroRange(pp *gc.Progs, p *obj.Prog, off, cnt int64, state *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	if cnt%8 != 0 {
		pstate.gc.Fatalf("zerorange count not a multiple of widthptr %d", cnt)
	}

	for i := int64(0); i < cnt; i += 8 {
		p = pp.Appendpp(pstate.gc, p, wasm.AGet, obj.TYPE_REG, wasm.REG_SP, 0, 0, 0, 0)
		p = pp.Appendpp(pstate.gc, p, wasm.AI64Const, obj.TYPE_CONST, 0, 0, 0, 0, 0)
		p = pp.Appendpp(pstate.gc, p, wasm.AI64Store, 0, 0, 0, obj.TYPE_CONST, 0, off+i)
	}

	return p
}

func (pstate *PackageState) zeroAuto(pp *gc.Progs, n *gc.Node) {
	sym := n.Sym.Linksym(pstate.types)
	size := n.Type.Size(pstate.types)
	for i := int64(0); i < size; i += 8 {
		p := pp.Prog(pstate.gc, wasm.AGet)
		p.From = obj.Addr{Type: obj.TYPE_REG, Reg: wasm.REG_SP}

		p = pp.Prog(pstate.gc, wasm.AI64Const)
		p.From = obj.Addr{Type: obj.TYPE_CONST, Offset: 0}

		p = pp.Prog(pstate.gc, wasm.AI64Store)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_AUTO, Offset: n.Xoffset + i, Sym: sym}
	}
}

func (pstate *PackageState) ginsnop(pp *gc.Progs) {
	pp.Prog(pstate.gc, wasm.ANop)
}

func ssaMarkMoves(s *gc.SSAGenState, b *ssa.Block) {
}

func (pstate *PackageState) ssaGenBlock(s *gc.SSAGenState, b, next *ssa.Block) {
	goToBlock := func(block *ssa.Block, canFallthrough bool) {
		if canFallthrough && block == next {
			return
		}
		s.Br(pstate.gc, obj.AJMP, block)
	}

	switch b.Kind {
	case ssa.BlockPlain:
		goToBlock(b.Succs[0].Block(), true)

	case ssa.BlockIf:
		pstate.getValue32(s, b.Control)
		s.Prog(pstate.gc, wasm.AI32Eqz)
		s.Prog(pstate.gc, wasm.AIf)
		goToBlock(b.Succs[1].Block(), false)
		s.Prog(pstate.gc, wasm.AEnd)
		goToBlock(b.Succs[0].Block(), true)

	case ssa.BlockRet:
		s.Prog(pstate.gc, obj.ARET)

	case ssa.BlockRetJmp:
		p := s.Prog(pstate.gc, obj.ARET)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = b.Aux.(*obj.LSym)

	case ssa.BlockExit:
		s.Prog(pstate.gc, obj.AUNDEF)

	case ssa.BlockDefer:
		p := s.Prog(pstate.gc, wasm.AGet)
		p.From = obj.Addr{Type: obj.TYPE_REG, Reg: wasm.REG_RET0}
		s.Prog(pstate.gc, wasm.AI64Eqz)
		s.Prog(pstate.gc, wasm.AI32Eqz)
		s.Prog(pstate.gc, wasm.AIf)
		goToBlock(b.Succs[1].Block(), false)
		s.Prog(pstate.gc, wasm.AEnd)
		goToBlock(b.Succs[0].Block(), true)

	default:
		panic("unexpected block")
	}

	// Entry point for the next block. Used by the JMP in goToBlock.
	s.Prog(pstate.gc, wasm.ARESUMEPOINT)

	if s.OnWasmStackSkipped != 0 {
		panic("wasm: bad stack")
	}
}

func (pstate *PackageState) ssaGenValue(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpWasmLoweredStaticCall, ssa.OpWasmLoweredClosureCall, ssa.OpWasmLoweredInterCall:
		s.PrepareCall(pstate.gc, v)
		if v.Aux == pstate.gc.Deferreturn {
			// add a resume point before call to deferreturn so it can be called again via jmpdefer
			s.Prog(pstate.gc, wasm.ARESUMEPOINT)
		}
		if v.Op == ssa.OpWasmLoweredClosureCall {
			pstate.getValue64(s, v.Args[1])
			pstate.setReg(s, wasm.REG_CTXT)
		}
		if sym, ok := v.Aux.(*obj.LSym); ok {
			p := s.Prog(pstate.gc, obj.ACALL)
			p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: sym}
		} else {
			pstate.getValue64(s, v.Args[0])
			p := s.Prog(pstate.gc, obj.ACALL)
			p.To = obj.Addr{Type: obj.TYPE_NONE}
		}

	case ssa.OpWasmLoweredMove:
		pstate.getValue32(s, v.Args[0])
		pstate.getValue32(s, v.Args[1])
		pstate.i32Const(s, int32(v.AuxInt))
		p := s.Prog(pstate.gc, wasm.ACall)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: pstate.gc.WasmMove}

	case ssa.OpWasmLoweredZero:
		pstate.getValue32(s, v.Args[0])
		pstate.i32Const(s, int32(v.AuxInt))
		p := s.Prog(pstate.gc, wasm.ACall)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: pstate.gc.WasmZero}

	case ssa.OpWasmLoweredNilCheck:
		pstate.getValue64(s, v.Args[0])
		s.Prog(pstate.gc, wasm.AI64Eqz)
		s.Prog(pstate.gc, wasm.AIf)
		p := s.Prog(pstate.gc, wasm.ACALLNORESUME)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: pstate.gc.SigPanic}
		s.Prog(pstate.gc, wasm.AEnd)
		if pstate.gc.Debug_checknil != 0 && v.Pos.Line() > 1 { // v.Pos.Line()==1 in generated wrappers
			pstate.gc.Warnl(v.Pos, "generated nil check")
		}

	case ssa.OpWasmLoweredWB:
		pstate.getValue64(s, v.Args[0])
		pstate.getValue64(s, v.Args[1])
		p := s.Prog(pstate.gc, wasm.ACALLNORESUME) // TODO(neelance): If possible, turn this into a simple wasm.ACall).
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: v.Aux.(*obj.LSym)}

	case ssa.OpWasmI64Store8, ssa.OpWasmI64Store16, ssa.OpWasmI64Store32, ssa.OpWasmI64Store, ssa.OpWasmF32Store, ssa.OpWasmF64Store:
		pstate.getValue32(s, v.Args[0])
		pstate.getValue64(s, v.Args[1])
		if v.Op == ssa.OpWasmF32Store {
			s.Prog(pstate.gc, wasm.AF32DemoteF64)
		}
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.To = obj.Addr{Type: obj.TYPE_CONST, Offset: v.AuxInt}

	case ssa.OpStoreReg:
		pstate.getReg(s, wasm.REG_SP)
		pstate.getValue64(s, v.Args[0])
		if v.Type.Etype == types.TFLOAT32 {
			s.Prog(pstate.gc, wasm.AF32DemoteF64)
		}
		p := s.Prog(pstate.gc, pstate.storeOp(v.Type))
		pstate.gc.AddrAuto(&p.To, v)

	default:
		if v.Type.IsMemory(pstate.types) {
			return
		}
		if v.OnWasmStack {
			s.OnWasmStackSkipped++
			// If a Value is marked OnWasmStack, we don't generate the value and store it to a register now.
			// Instead, we delay the generation to when the value is used and then directly generate it on the WebAssembly stack.
			return
		}
		pstate.ssaGenValueOnStack(s, v)
		if s.OnWasmStackSkipped != 0 {
			panic("wasm: bad stack")
		}
		pstate.setReg(s, v.Reg(pstate.ssa))
	}
}

func (pstate *PackageState) ssaGenValueOnStack(s *gc.SSAGenState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpWasmLoweredGetClosurePtr:
		pstate.getReg(s, wasm.REG_CTXT)

	case ssa.OpWasmLoweredGetCallerPC:
		p := s.Prog(pstate.gc, wasm.AI64Load)
		// Caller PC is stored 8 bytes below first parameter.
		p.From = obj.Addr{
			Type:   obj.TYPE_MEM,
			Name:   obj.NAME_PARAM,
			Offset: -8,
		}

	case ssa.OpWasmLoweredGetCallerSP:
		p := s.Prog(pstate.gc, wasm.AGet)
		// Caller SP is the address of the first parameter.
		p.From = obj.Addr{
			Type:   obj.TYPE_ADDR,
			Name:   obj.NAME_PARAM,
			Reg:    wasm.REG_SP,
			Offset: 0,
		}

	case ssa.OpWasmLoweredAddr:
		p := s.Prog(pstate.gc, wasm.AGet)
		switch n := v.Aux.(type) {
		case *obj.LSym:
			p.From = obj.Addr{Type: obj.TYPE_ADDR, Name: obj.NAME_EXTERN, Sym: n}
		case *gc.Node:
			p.From = obj.Addr{
				Type:   obj.TYPE_ADDR,
				Name:   obj.NAME_AUTO,
				Reg:    v.Args[0].Reg(pstate.ssa),
				Offset: n.Xoffset,
			}
			if n.Class() == gc.PPARAM || n.Class() == gc.PPARAMOUT {
				p.From.Name = obj.NAME_PARAM
			}
		default:
			panic("wasm: bad LoweredAddr")
		}

	case ssa.OpWasmLoweredRound32F:
		pstate.getValue64(s, v.Args[0])
		s.Prog(pstate.gc, wasm.AF32DemoteF64)
		s.Prog(pstate.gc, wasm.AF64PromoteF32)

	case ssa.OpWasmLoweredConvert:
		pstate.getValue64(s, v.Args[0])

	case ssa.OpWasmSelect:
		pstate.getValue64(s, v.Args[0])
		pstate.getValue64(s, v.Args[1])
		pstate.getValue64(s, v.Args[2])
		s.Prog(pstate.gc, wasm.AI32WrapI64)
		s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))

	case ssa.OpWasmI64AddConst:
		pstate.getValue64(s, v.Args[0])
		pstate.i64Const(s, v.AuxInt)
		s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))

	case ssa.OpWasmI64Const:
		pstate.i64Const(s, v.AuxInt)

	case ssa.OpWasmF64Const:
		pstate.f64Const(s, v.AuxFloat(pstate.ssa))

	case ssa.OpWasmI64Load8U, ssa.OpWasmI64Load8S, ssa.OpWasmI64Load16U, ssa.OpWasmI64Load16S, ssa.OpWasmI64Load32U, ssa.OpWasmI64Load32S, ssa.OpWasmI64Load, ssa.OpWasmF32Load, ssa.OpWasmF64Load:
		pstate.getValue32(s, v.Args[0])
		p := s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		p.From = obj.Addr{Type: obj.TYPE_CONST, Offset: v.AuxInt}
		if v.Op == ssa.OpWasmF32Load {
			s.Prog(pstate.gc, wasm.AF64PromoteF32)
		}

	case ssa.OpWasmI64Eqz:
		pstate.getValue64(s, v.Args[0])
		s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		s.Prog(pstate.gc, wasm.AI64ExtendUI32)

	case ssa.OpWasmI64Eq, ssa.OpWasmI64Ne, ssa.OpWasmI64LtS, ssa.OpWasmI64LtU, ssa.OpWasmI64GtS, ssa.OpWasmI64GtU, ssa.OpWasmI64LeS, ssa.OpWasmI64LeU, ssa.OpWasmI64GeS, ssa.OpWasmI64GeU, ssa.OpWasmF64Eq, ssa.OpWasmF64Ne, ssa.OpWasmF64Lt, ssa.OpWasmF64Gt, ssa.OpWasmF64Le, ssa.OpWasmF64Ge:
		pstate.getValue64(s, v.Args[0])
		pstate.getValue64(s, v.Args[1])
		s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))
		s.Prog(pstate.gc, wasm.AI64ExtendUI32)

	case ssa.OpWasmI64Add, ssa.OpWasmI64Sub, ssa.OpWasmI64Mul, ssa.OpWasmI64DivU, ssa.OpWasmI64RemS, ssa.OpWasmI64RemU, ssa.OpWasmI64And, ssa.OpWasmI64Or, ssa.OpWasmI64Xor, ssa.OpWasmI64Shl, ssa.OpWasmI64ShrS, ssa.OpWasmI64ShrU, ssa.OpWasmF64Add, ssa.OpWasmF64Sub, ssa.OpWasmF64Mul, ssa.OpWasmF64Div:
		pstate.getValue64(s, v.Args[0])
		pstate.getValue64(s, v.Args[1])
		s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))

	case ssa.OpWasmI64DivS:
		pstate.getValue64(s, v.Args[0])
		pstate.getValue64(s, v.Args[1])
		if v.Type.Size(pstate.types) == 8 {
			// Division of int64 needs helper function wasmDiv to handle the MinInt64 / -1 case.
			p := s.Prog(pstate.gc, wasm.ACall)
			p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: pstate.gc.WasmDiv}
			break
		}
		s.Prog(pstate.gc, wasm.AI64DivS)

	case ssa.OpWasmI64TruncSF64:
		pstate.getValue64(s, v.Args[0])
		p := s.Prog(pstate.gc, wasm.ACall)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: pstate.gc.WasmTruncS}

	case ssa.OpWasmI64TruncUF64:
		pstate.getValue64(s, v.Args[0])
		p := s.Prog(pstate.gc, wasm.ACall)
		p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: pstate.gc.WasmTruncU}

	case ssa.OpWasmF64Neg, ssa.OpWasmF64ConvertSI64, ssa.OpWasmF64ConvertUI64:
		pstate.getValue64(s, v.Args[0])
		s.Prog(pstate.gc, v.Op.Asm(pstate.ssa))

	case ssa.OpLoadReg:
		p := s.Prog(pstate.gc, pstate.loadOp(v.Type))
		pstate.gc.AddrAuto(&p.From, v.Args[0])
		if v.Type.Etype == types.TFLOAT32 {
			s.Prog(pstate.gc, wasm.AF64PromoteF32)
		}

	case ssa.OpCopy:
		pstate.getValue64(s, v.Args[0])

	default:
		v.Fatalf("unexpected op: %s", v.Op)

	}
}

func (pstate *PackageState) getValue32(s *gc.SSAGenState, v *ssa.Value) {
	if v.OnWasmStack {
		s.OnWasmStackSkipped--
		pstate.ssaGenValueOnStack(s, v)
		s.Prog(pstate.gc, wasm.AI32WrapI64)
		return
	}

	reg := v.Reg(pstate.ssa)
	pstate.getReg(s, reg)
	if reg != wasm.REG_SP {
		s.Prog(pstate.gc, wasm.AI32WrapI64)
	}
}

func (pstate *PackageState) getValue64(s *gc.SSAGenState, v *ssa.Value) {
	if v.OnWasmStack {
		s.OnWasmStackSkipped--
		pstate.ssaGenValueOnStack(s, v)
		return
	}

	reg := v.Reg(pstate.ssa)
	pstate.getReg(s, reg)
	if reg == wasm.REG_SP {
		s.Prog(pstate.gc, wasm.AI64ExtendUI32)
	}
}

func (pstate *PackageState) i32Const(s *gc.SSAGenState, val int32) {
	p := s.Prog(pstate.gc, wasm.AI32Const)
	p.From = obj.Addr{Type: obj.TYPE_CONST, Offset: int64(val)}
}

func (pstate *PackageState) i64Const(s *gc.SSAGenState, val int64) {
	p := s.Prog(pstate.gc, wasm.AI64Const)
	p.From = obj.Addr{Type: obj.TYPE_CONST, Offset: val}
}

func (pstate *PackageState) f64Const(s *gc.SSAGenState, val float64) {
	p := s.Prog(pstate.gc, wasm.AF64Const)
	p.From = obj.Addr{Type: obj.TYPE_FCONST, Val: val}
}

func (pstate *PackageState) getReg(s *gc.SSAGenState, reg int16) {
	p := s.Prog(pstate.gc, wasm.AGet)
	p.From = obj.Addr{Type: obj.TYPE_REG, Reg: reg}
}

func (pstate *PackageState) setReg(s *gc.SSAGenState, reg int16) {
	p := s.Prog(pstate.gc, wasm.ASet)
	p.To = obj.Addr{Type: obj.TYPE_REG, Reg: reg}
}

func (pstate *PackageState) loadOp(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(pstate.types) {
		case 4:
			return wasm.AF32Load
		case 8:
			return wasm.AF64Load
		default:
			panic("bad load type")
		}
	}

	switch t.Size(pstate.types) {
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

func (pstate *PackageState) storeOp(t *types.Type) obj.As {
	if t.IsFloat() {
		switch t.Size(pstate.types) {
		case 4:
			return wasm.AF32Store
		case 8:
			return wasm.AF64Store
		default:
			panic("bad store type")
		}
	}

	switch t.Size(pstate.types) {
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

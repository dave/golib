// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wasm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"io"
	"math"
)

func (pstate *PackageState) init() {
	pstate.obj.RegisterRegister(MINREG, MAXREG, pstate.rconv)
	pstate.obj.RegisterOpcode(obj.ABaseWasm, pstate.Anames)

	pstate.registerNames = make([]string, MAXREG-MINREG)
	for name, reg := range pstate.Register {
		pstate.registerNames[reg-MINREG] = name
	}
}

func (pstate *PackageState) rconv(r int) string {
	return pstate.registerNames[r-MINREG]
}

const (
	/* mark flags */
	WasmImport = 1 << 0
)

func (pstate *PackageState) instinit(ctxt *obj.Link) {
	pstate.morestack = ctxt.Lookup("runtime.morestack")
	pstate.morestackNoCtxt = ctxt.Lookup("runtime.morestack_noctxt")
	pstate.gcWriteBarrier = ctxt.Lookup("runtime.gcWriteBarrier")
	pstate.sigpanic = ctxt.Lookup("runtime.sigpanic")
	pstate.deferreturn = ctxt.Lookup("runtime.deferreturn")
	pstate.jmpdefer = ctxt.Lookup("\"\".jmpdefer")
}

func (pstate *PackageState) preprocess(ctxt *obj.Link, s *obj.LSym, newprog obj.ProgAlloc) {
	appendp := func(p *obj.Prog, as obj.As, args ...obj.Addr) *obj.Prog {
		if p.As != obj.ANOP {
			p2 := obj.Appendp(p, newprog)
			p2.Pc = p.Pc
			p = p2
		}
		p.As = as
		switch len(args) {
		case 0:
			p.From = obj.Addr{}
			p.To = obj.Addr{}
		case 1:
			if pstate.unaryDst[as] {
				p.From = obj.Addr{}
				p.To = args[0]
			} else {
				p.From = args[0]
				p.To = obj.Addr{}
			}
		case 2:
			p.From = args[0]
			p.To = args[1]
		default:
			panic("bad args")
		}
		return p
	}

	framesize := s.Func.Text.To.Offset
	if framesize < 0 {
		panic("bad framesize")
	}
	s.Func.Args = s.Func.Text.To.Val.(int32)
	s.Func.Locals = int32(framesize)

	if s.Func.Text.From.Sym.Wrapper() {
		// if g._panic != nil && g._panic.argp == FP {
		//   g._panic.argp = bottom-of-frame
		// }
		//
		// MOVD g_panic(g), R0
		// Get R0
		// I64Eqz
		// Not
		// If
		//   Get SP
		//   I64ExtendUI32
		//   I64Const $framesize+8
		//   I64Add
		//   I64Load panic_argp(R0)
		//   I64Eq
		//   If
		//     MOVD SP, panic_argp(R0)
		//   End
		// End

		gpanic := obj.Addr{
			Type:   obj.TYPE_MEM,
			Reg:    REGG,
			Offset: 4 * 8, // g_panic
		}

		panicargp := obj.Addr{
			Type:   obj.TYPE_MEM,
			Reg:    REG_R0,
			Offset: 0, // panic.argp
		}

		p := s.Func.Text
		p = appendp(p, AMOVD, gpanic, regAddr(REG_R0))

		p = appendp(p, AGet, regAddr(REG_R0))
		p = appendp(p, AI64Eqz)
		p = appendp(p, ANot)
		p = appendp(p, AIf)

		p = appendp(p, AGet, regAddr(REG_SP))
		p = appendp(p, AI64ExtendUI32)
		p = appendp(p, AI64Const, constAddr(framesize+8))
		p = appendp(p, AI64Add)
		p = appendp(p, AI64Load, panicargp)

		p = appendp(p, AI64Eq)
		p = appendp(p, AIf)
		p = appendp(p, AMOVD, regAddr(REG_SP), panicargp)
		p = appendp(p, AEnd)

		p = appendp(p, AEnd)
	}

	if framesize > 0 {
		p := s.Func.Text
		p = appendp(p, AGet, regAddr(REG_SP))
		p = appendp(p, AI32Const, constAddr(framesize))
		p = appendp(p, AI32Sub)
		p = appendp(p, ASet, regAddr(REG_SP))
		p.Spadj = int32(framesize)
	}

	// Introduce resume points for CALL instructions
	// and collect other explicit resume points.
	numResumePoints := 0
	explicitBlockDepth := 0
	pc := int64(0) // pc is only incremented when necessary, this avoids bloat of the BrTable instruction
	var tableIdxs []uint64
	tablePC := int64(0)
	base := ctxt.PosTable.Pos(s.Func.Text.Pos).Base()
	for p := s.Func.Text; p != nil; p = p.Link {
		prevBase := base
		base = ctxt.PosTable.Pos(p.Pos).Base()

		switch p.As {
		case ABlock, ALoop, AIf:
			explicitBlockDepth++

		case AEnd:
			if explicitBlockDepth == 0 {
				panic("End without block")
			}
			explicitBlockDepth--

		case ARESUMEPOINT:
			if explicitBlockDepth != 0 {
				panic("RESUME can only be used on toplevel")
			}
			p.As = AEnd
			for tablePC <= pc {
				tableIdxs = append(tableIdxs, uint64(numResumePoints))
				tablePC++
			}
			numResumePoints++
			pc++

		case obj.ACALL:
			if explicitBlockDepth != 0 {
				panic("CALL can only be used on toplevel, try CALLNORESUME instead")
			}
			appendp(p, ARESUMEPOINT)
		}

		p.Pc = pc

		// Increase pc whenever some pc-value table needs a new entry. Don't increase it
		// more often to avoid bloat of the BrTable instruction.
		// The "base != prevBase" condition detects inlined instructions. They are an
		// implicit call, so entering and leaving this section affects the stack trace.
		if p.As == ACALLNORESUME || p.As == obj.ANOP || p.Spadj != 0 || base != prevBase {
			pc++
		}
	}
	tableIdxs = append(tableIdxs, uint64(numResumePoints))
	s.Size = pc + 1

	if !s.Func.Text.From.Sym.NoSplit() {
		p := s.Func.Text

		if framesize <= objabi.StackSmall {
			// small stack: SP <= stackguard
			// Get SP
			// Get g
			// I32WrapI64
			// I32Load $stackguard0
			// I32GtU

			p = appendp(p, AGet, regAddr(REG_SP))
			p = appendp(p, AGet, regAddr(REGG))
			p = appendp(p, AI32WrapI64)
			p = appendp(p, AI32Load, constAddr(2*int64(ctxt.Arch.PtrSize))) // G.stackguard0
			p = appendp(p, AI32LeU)
		} else {
			// large stack: SP-framesize <= stackguard-StackSmall
			//              SP <= stackguard+(framesize-StackSmall)
			// Get SP
			// Get g
			// I32WrapI64
			// I32Load $stackguard0
			// I32Const $(framesize-StackSmall)
			// I32Add
			// I32GtU

			p = appendp(p, AGet, regAddr(REG_SP))
			p = appendp(p, AGet, regAddr(REGG))
			p = appendp(p, AI32WrapI64)
			p = appendp(p, AI32Load, constAddr(2*int64(ctxt.Arch.PtrSize))) // G.stackguard0
			p = appendp(p, AI32Const, constAddr(int64(framesize)-objabi.StackSmall))
			p = appendp(p, AI32Add)
			p = appendp(p, AI32LeU)
		}
		// TODO(neelance): handle wraparound case

		p = appendp(p, AIf)
		p = appendp(p, obj.ACALL, constAddr(0))
		if s.Func.Text.From.Sym.NeedCtxt() {
			p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: pstate.morestack}
		} else {
			p.To = obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: pstate.morestackNoCtxt}
		}
		p = appendp(p, AEnd)
	}

	// Add Block instructions for resume points and BrTable to jump to selected resume point.
	if numResumePoints > 0 {
		p := s.Func.Text
		p = appendp(p, ALoop) // entryPointLoop, used to jump between basic blocks

		for i := 0; i < numResumePoints+1; i++ {
			p = appendp(p, ABlock)
		}
		p = appendp(p, AGet, regAddr(REG_PC_B)) // read next basic block from PC_B
		p = appendp(p, ABrTable, obj.Addr{Val: tableIdxs})
		p = appendp(p, AEnd) // end of Block

		for p.Link != nil {
			p = p.Link
		}

		p = appendp(p, AEnd) // end of entryPointLoop
		p = appendp(p, obj.AUNDEF)
	}

	p := s.Func.Text
	currentDepth := 0
	blockDepths := make(map[*obj.Prog]int)
	for p != nil {
		switch p.As {
		case ABlock, ALoop, AIf:
			currentDepth++
			blockDepths[p] = currentDepth
		case AEnd:
			currentDepth--
		}

		switch p.As {
		case ABr, ABrIf:
			if p.To.Type == obj.TYPE_BRANCH {
				blockDepth, ok := blockDepths[p.To.Val.(*obj.Prog)]
				if !ok {
					panic("label not at block")
				}
				p.To = constAddr(int64(currentDepth - blockDepth))
			}
		case obj.AJMP:
			jmp := *p
			p.As = obj.ANOP

			if jmp.To.Type == obj.TYPE_BRANCH {
				// jump to basic block
				p = appendp(p, AI32Const, constAddr(jmp.To.Val.(*obj.Prog).Pc))
				p = appendp(p, ASet, regAddr(REG_PC_B))               // write next basic block to PC_B
				p = appendp(p, ABr, constAddr(int64(currentDepth-1))) // jump to beginning of entryPointLoop
				break
			}

			// reset PC_B to function entry
			p = appendp(p, AI32Const, constAddr(0))
			p = appendp(p, ASet, regAddr(REG_PC_B))

			// low-level WebAssembly call to function
			switch jmp.To.Type {
			case obj.TYPE_MEM:
				p = appendp(p, ACall, jmp.To)
			case obj.TYPE_NONE:
				// (target PC is on stack)
				p = appendp(p, AI32WrapI64)
				p = appendp(p, AI32Const, constAddr(16)) // only needs PC_F bits (16-31), PC_B bits (0-15) are zero
				p = appendp(p, AI32ShrU)
				p = appendp(p, ACallIndirect)
			default:
				panic("bad target for JMP")
			}

			p = appendp(p, AReturn)

		case obj.ACALL, ACALLNORESUME:
			call := *p
			p.As = obj.ANOP

			pcAfterCall := call.Link.Pc
			if call.To.Sym == pstate.sigpanic {
				pcAfterCall-- // sigpanic expects to be called without advancing the pc
			}

			// jmpdefer manipulates the return address on the stack so deferreturn gets called repeatedly.
			// Model this in WebAssembly with a loop.
			if call.To.Sym == pstate.deferreturn {
				p = appendp(p, ALoop)
			}

			// SP -= 8
			p = appendp(p, AGet, regAddr(REG_SP))
			p = appendp(p, AI32Const, constAddr(8))
			p = appendp(p, AI32Sub)
			p = appendp(p, ASet, regAddr(REG_SP))

			// write return address to Go stack
			p = appendp(p, AGet, regAddr(REG_SP))
			p = appendp(p, AI64Const, obj.Addr{
				Type:   obj.TYPE_ADDR,
				Name:   obj.NAME_EXTERN,
				Sym:    s,           // PC_F
				Offset: pcAfterCall, // PC_B
			})
			p = appendp(p, AI64Store, constAddr(0))

			// reset PC_B to function entry
			p = appendp(p, AI32Const, constAddr(0))
			p = appendp(p, ASet, regAddr(REG_PC_B))

			// low-level WebAssembly call to function
			switch call.To.Type {
			case obj.TYPE_MEM:
				p = appendp(p, ACall, call.To)
			case obj.TYPE_NONE:
				// (target PC is on stack)
				p = appendp(p, AI32WrapI64)
				p = appendp(p, AI32Const, constAddr(16)) // only needs PC_F bits (16-31), PC_B bits (0-15) are zero
				p = appendp(p, AI32ShrU)
				p = appendp(p, ACallIndirect)
			default:
				panic("bad target for CALL")
			}

			// gcWriteBarrier has no return value, it never unwinds the stack
			if call.To.Sym == pstate.gcWriteBarrier {
				break
			}

			// jmpdefer removes the frame of deferreturn from the Go stack.
			// However, its WebAssembly function still returns normally,
			// so we need to return from deferreturn without removing its
			// stack frame (no RET), because the frame is already gone.
			if call.To.Sym == pstate.jmpdefer {
				p = appendp(p, AReturn)
				break
			}

			// return value of call is on the top of the stack, indicating whether to unwind the WebAssembly stack
			p = appendp(p, AIf)
			if call.As == ACALLNORESUME && call.To.Sym != pstate.sigpanic { // sigpanic unwinds the stack, but it never resumes
				// trying to unwind WebAssembly stack but call has no resume point, terminate with error
				p = appendp(p, obj.AUNDEF)
			} else {
				// unwinding WebAssembly stack to switch goroutine, return 1
				p = appendp(p, AI32Const, constAddr(1))
				p = appendp(p, AReturn)
			}
			p = appendp(p, AEnd)

			// jump to before the call if jmpdefer has reset the return address to the call's PC
			if call.To.Sym == pstate.deferreturn {
				p = appendp(p, AGet, regAddr(REG_PC_B))
				p = appendp(p, AI32Const, constAddr(call.Pc))
				p = appendp(p, AI32Eq)
				p = appendp(p, ABrIf, constAddr(0))
				p = appendp(p, AEnd) // end of Loop
			}

		case obj.ARET, ARETUNWIND:
			ret := *p
			p.As = obj.ANOP

			if framesize > 0 {
				// SP += framesize
				p = appendp(p, AGet, regAddr(REG_SP))
				p = appendp(p, AI32Const, constAddr(framesize))
				p = appendp(p, AI32Add)
				p = appendp(p, ASet, regAddr(REG_SP))
				// TODO(neelance): This should theoretically set Spadj, but it only works without.
				// p.Spadj = int32(-framesize)
			}

			if ret.To.Type == obj.TYPE_MEM {
				// reset PC_B to function entry
				p = appendp(p, AI32Const, constAddr(0))
				p = appendp(p, ASet, regAddr(REG_PC_B))

				// low-level WebAssembly call to function
				p = appendp(p, ACall, ret.To)
				p = appendp(p, AReturn)
				break
			}

			// read return PC_F from Go stack
			p = appendp(p, AGet, regAddr(REG_SP))
			p = appendp(p, AI32Load16U, constAddr(2))
			p = appendp(p, ASet, regAddr(REG_PC_F))

			// read return PC_B from Go stack
			p = appendp(p, AGet, regAddr(REG_SP))
			p = appendp(p, AI32Load16U, constAddr(0))
			p = appendp(p, ASet, regAddr(REG_PC_B))

			// SP += 8
			p = appendp(p, AGet, regAddr(REG_SP))
			p = appendp(p, AI32Const, constAddr(8))
			p = appendp(p, AI32Add)
			p = appendp(p, ASet, regAddr(REG_SP))

			if ret.As == ARETUNWIND {
				// function needs to unwind the WebAssembly stack, return 1
				p = appendp(p, AI32Const, constAddr(1))
				p = appendp(p, AReturn)
				break
			}

			// not unwinding the WebAssembly stack, return 0
			p = appendp(p, AI32Const, constAddr(0))
			p = appendp(p, AReturn)
		}

		p = p.Link
	}

	p = s.Func.Text
	for p != nil {
		switch p.From.Name {
		case obj.NAME_AUTO:
			p.From.Offset += int64(framesize)
		case obj.NAME_PARAM:
			p.From.Reg = REG_SP
			p.From.Offset += int64(framesize) + 8 // parameters are after the frame and the 8-byte return address
		}

		switch p.To.Name {
		case obj.NAME_AUTO:
			p.To.Offset += int64(framesize)
		case obj.NAME_PARAM:
			p.To.Reg = REG_SP
			p.To.Offset += int64(framesize) + 8 // parameters are after the frame and the 8-byte return address
		}

		switch p.As {
		case AGet:
			if p.From.Type == obj.TYPE_ADDR {
				get := *p
				p.As = obj.ANOP

				switch get.From.Name {
				case obj.NAME_EXTERN:
					p = appendp(p, AI64Const, get.From)
				case obj.NAME_AUTO, obj.NAME_PARAM:
					p = appendp(p, AGet, regAddr(get.From.Reg))
					if get.From.Reg == REG_SP {
						p = appendp(p, AI64ExtendUI32)
					}
					if get.From.Offset != 0 {
						p = appendp(p, AI64Const, constAddr(get.From.Offset))
						p = appendp(p, AI64Add)
					}
				default:
					panic("bad Get: invalid name")
				}
			}

		case AI32Load, AI64Load, AF32Load, AF64Load, AI32Load8S, AI32Load8U, AI32Load16S, AI32Load16U, AI64Load8S, AI64Load8U, AI64Load16S, AI64Load16U, AI64Load32S, AI64Load32U:
			if p.From.Type == obj.TYPE_MEM {
				as := p.As
				from := p.From

				p.As = AGet
				p.From = regAddr(from.Reg)

				if from.Reg != REG_SP {
					p = appendp(p, AI32WrapI64)
				}

				p = appendp(p, as, constAddr(from.Offset))
			}

		case AMOVB, AMOVH, AMOVW, AMOVD:
			mov := *p
			p.As = obj.ANOP

			var loadAs obj.As
			var storeAs obj.As
			switch mov.As {
			case AMOVB:
				loadAs = AI64Load8U
				storeAs = AI64Store8
			case AMOVH:
				loadAs = AI64Load16U
				storeAs = AI64Store16
			case AMOVW:
				loadAs = AI64Load32U
				storeAs = AI64Store32
			case AMOVD:
				loadAs = AI64Load
				storeAs = AI64Store
			}

			appendValue := func() {
				switch mov.From.Type {
				case obj.TYPE_CONST:
					p = appendp(p, AI64Const, constAddr(mov.From.Offset))

				case obj.TYPE_ADDR:
					switch mov.From.Name {
					case obj.NAME_NONE, obj.NAME_PARAM, obj.NAME_AUTO:
						p = appendp(p, AGet, regAddr(mov.From.Reg))
						if mov.From.Reg == REG_SP {
							p = appendp(p, AI64ExtendUI32)
						}
						p = appendp(p, AI64Const, constAddr(mov.From.Offset))
						p = appendp(p, AI64Add)
					case obj.NAME_EXTERN:
						p = appendp(p, AI64Const, mov.From)
					default:
						panic("bad name for MOV")
					}

				case obj.TYPE_REG:
					p = appendp(p, AGet, mov.From)
					if mov.From.Reg == REG_SP {
						p = appendp(p, AI64ExtendUI32)
					}

				case obj.TYPE_MEM:
					p = appendp(p, AGet, regAddr(mov.From.Reg))
					if mov.From.Reg != REG_SP {
						p = appendp(p, AI32WrapI64)
					}
					p = appendp(p, loadAs, constAddr(mov.From.Offset))

				default:
					panic("bad MOV type")
				}
			}

			switch mov.To.Type {
			case obj.TYPE_REG:
				appendValue()
				if mov.To.Reg == REG_SP {
					p = appendp(p, AI32WrapI64)
				}
				p = appendp(p, ASet, mov.To)

			case obj.TYPE_MEM:
				switch mov.To.Name {
				case obj.NAME_NONE, obj.NAME_PARAM:
					p = appendp(p, AGet, regAddr(mov.To.Reg))
					if mov.To.Reg != REG_SP {
						p = appendp(p, AI32WrapI64)
					}
				case obj.NAME_EXTERN:
					p = appendp(p, AI32Const, obj.Addr{Type: obj.TYPE_ADDR, Name: obj.NAME_EXTERN, Sym: mov.To.Sym})
				default:
					panic("bad MOV name")
				}
				appendValue()
				p = appendp(p, storeAs, constAddr(mov.To.Offset))

			default:
				panic("bad MOV type")
			}

		case ACallImport:
			p.As = obj.ANOP
			p = appendp(p, AGet, regAddr(REG_SP))
			p = appendp(p, ACall, obj.Addr{Type: obj.TYPE_MEM, Name: obj.NAME_EXTERN, Sym: s})
			p.Mark = WasmImport
		}

		p = p.Link
	}
}

func constAddr(value int64) obj.Addr {
	return obj.Addr{Type: obj.TYPE_CONST, Offset: value}
}

func regAddr(reg int16) obj.Addr {
	return obj.Addr{Type: obj.TYPE_REG, Reg: reg}
}

func assemble(ctxt *obj.Link, s *obj.LSym, newprog obj.ProgAlloc) {
	w := new(bytes.Buffer)

	// Function starts with declaration of locals: numbers and types.
	switch s.Name {
	case "memchr":
		writeUleb128(w, 1) // number of sets of locals
		writeUleb128(w, 3) // number of locals
		w.WriteByte(0x7F)  // i32
	case "memcmp":
		writeUleb128(w, 1) // number of sets of locals
		writeUleb128(w, 2) // number of locals
		w.WriteByte(0x7F)  // i32
	default:
		writeUleb128(w, 2)  // number of sets of locals
		writeUleb128(w, 16) // number of locals
		w.WriteByte(0x7E)   // i64
		writeUleb128(w, 16) // number of locals
		w.WriteByte(0x7C)   // f64
	}

	for p := s.Func.Text; p != nil; p = p.Link {
		switch p.As {
		case AGet:
			if p.From.Type != obj.TYPE_REG {
				panic("bad Get: argument is not a register")
			}
			reg := p.From.Reg
			switch {
			case reg >= REG_PC_F && reg <= REG_RUN:
				w.WriteByte(0x23) // get_global
				writeUleb128(w, uint64(reg-REG_PC_F))
			case reg >= REG_R0 && reg <= REG_F15:
				w.WriteByte(0x20) // get_local
				writeUleb128(w, uint64(reg-REG_R0))
			default:
				panic("bad Get: invalid register")
			}
			continue

		case ASet:
			if p.To.Type != obj.TYPE_REG {
				panic("bad Set: argument is not a register")
			}
			reg := p.To.Reg
			switch {
			case reg >= REG_PC_F && reg <= REG_RUN:
				w.WriteByte(0x24) // set_global
				writeUleb128(w, uint64(reg-REG_PC_F))
			case reg >= REG_R0 && reg <= REG_F15:
				if p.Link.As == AGet && p.Link.From.Reg == reg {
					w.WriteByte(0x22) // tee_local
					p = p.Link
				} else {
					w.WriteByte(0x21) // set_local
				}
				writeUleb128(w, uint64(reg-REG_R0))
			default:
				panic("bad Set: invalid register")
			}
			continue

		case ATee:
			if p.To.Type != obj.TYPE_REG {
				panic("bad Tee: argument is not a register")
			}
			reg := p.To.Reg
			switch {
			case reg >= REG_R0 && reg <= REG_F15:
				w.WriteByte(0x22) // tee_local
				writeUleb128(w, uint64(reg-REG_R0))
			default:
				panic("bad Tee: invalid register")
			}
			continue

		case ANot:
			w.WriteByte(0x45) // i32.eqz
			continue

		case obj.AUNDEF:
			w.WriteByte(0x00) // unreachable
			continue

		case obj.ANOP, obj.ATEXT, obj.AFUNCDATA, obj.APCDATA:
			// ignore
			continue
		}

		switch {
		case p.As < AUnreachable || p.As > AF64ReinterpretI64:
			panic(fmt.Sprintf("unexpected assembler op: %s", p.As))
		case p.As < AEnd:
			w.WriteByte(byte(p.As - AUnreachable + 0x00))
		case p.As < ADrop:
			w.WriteByte(byte(p.As - AEnd + 0x0B))
		case p.As < AI32Load:
			w.WriteByte(byte(p.As - ADrop + 0x1A))
		default:
			w.WriteByte(byte(p.As - AI32Load + 0x28))
		}

		switch p.As {
		case ABlock, ALoop, AIf:
			if p.From.Offset != 0 {
				// block type, rarely used, e.g. for code compiled with emscripten
				w.WriteByte(0x80 - byte(p.From.Offset))
				continue
			}
			w.WriteByte(0x40)

		case ABr, ABrIf:
			if p.To.Type != obj.TYPE_CONST {
				panic("bad Br/BrIf")
			}
			writeUleb128(w, uint64(p.To.Offset))

		case ABrTable:
			idxs := p.To.Val.([]uint64)
			writeUleb128(w, uint64(len(idxs)-1))
			for _, idx := range idxs {
				writeUleb128(w, idx)
			}

		case ACall:
			switch p.To.Type {
			case obj.TYPE_CONST:
				writeUleb128(w, uint64(p.To.Offset))

			case obj.TYPE_MEM:
				if p.To.Name != obj.NAME_EXTERN && p.To.Name != obj.NAME_STATIC {
					fmt.Println(p.To)
					panic("bad name for Call")
				}
				r := obj.Addrel(s)
				r.Off = int32(w.Len())
				r.Type = objabi.R_CALL
				if p.Mark&WasmImport != 0 {
					r.Type = objabi.R_WASMIMPORT
				}
				r.Sym = p.To.Sym

			default:
				panic("bad type for Call")
			}

		case ACallIndirect:
			writeUleb128(w, uint64(p.To.Offset))
			w.WriteByte(0x00) // reserved value

		case AI32Const, AI64Const:
			if p.From.Name == obj.NAME_EXTERN {
				r := obj.Addrel(s)
				r.Off = int32(w.Len())
				r.Type = objabi.R_ADDR
				r.Sym = p.From.Sym
				r.Add = p.From.Offset
				break
			}
			writeSleb128(w, p.From.Offset)

		case AF64Const:
			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b, math.Float64bits(p.From.Val.(float64)))
			w.Write(b)

		case AI32Load, AI64Load, AF32Load, AF64Load, AI32Load8S, AI32Load8U, AI32Load16S, AI32Load16U, AI64Load8S, AI64Load8U, AI64Load16S, AI64Load16U, AI64Load32S, AI64Load32U:
			if p.From.Offset < 0 {
				panic("negative offset for *Load")
			}
			if p.From.Type != obj.TYPE_CONST {
				panic("bad type for *Load")
			}
			if p.From.Offset > math.MaxUint32 {
				ctxt.Diag("bad offset in %v", p)
			}
			writeUleb128(w, align(p.As))
			writeUleb128(w, uint64(p.From.Offset))

		case AI32Store, AI64Store, AF32Store, AF64Store, AI32Store8, AI32Store16, AI64Store8, AI64Store16, AI64Store32:
			if p.To.Offset < 0 {
				panic("negative offset")
			}
			if p.From.Offset > math.MaxUint32 {
				ctxt.Diag("bad offset in %v", p)
			}
			writeUleb128(w, align(p.As))
			writeUleb128(w, uint64(p.To.Offset))

		case ACurrentMemory, AGrowMemory:
			w.WriteByte(0x00)

		}
	}

	w.WriteByte(0x0b) // end

	s.P = w.Bytes()
}

func align(as obj.As) uint64 {
	switch as {
	case AI32Load8S, AI32Load8U, AI64Load8S, AI64Load8U, AI32Store8, AI64Store8:
		return 0
	case AI32Load16S, AI32Load16U, AI64Load16S, AI64Load16U, AI32Store16, AI64Store16:
		return 1
	case AI32Load, AF32Load, AI64Load32S, AI64Load32U, AI32Store, AF32Store, AI64Store32:
		return 2
	case AI64Load, AF64Load, AI64Store, AF64Store:
		return 3
	default:
		panic("align: bad op")
	}
}

func writeUleb128(w io.ByteWriter, v uint64) {
	more := true
	for more {
		c := uint8(v & 0x7f)
		v >>= 7
		more = v != 0
		if more {
			c |= 0x80
		}
		w.WriteByte(c)
	}
}

func writeSleb128(w io.ByteWriter, v int64) {
	more := true
	for more {
		c := uint8(v & 0x7f)
		s := uint8(v & 0x40)
		v >>= 7
		more = !((v == 0 && s == 0) || (v == -1 && s != 0))
		if more {
			c |= 0x80
		}
		w.WriteByte(c)
	}
}

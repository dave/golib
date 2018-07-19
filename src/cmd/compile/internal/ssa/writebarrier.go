package ssa

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/src"
	"strings"
)

// needwb returns whether we need write barrier for store op v.
// v must be Store/Move/Zero.
func (psess *PackageSession) needwb(v *Value) bool {
	t, ok := v.Aux.(*types.Type)
	if !ok {
		v.Fatalf("store aux is not a type: %s", v.LongString(psess))
	}
	if !t.HasHeapPointer(psess.types) {
		return false
	}
	if IsStackAddr(v.Args[0]) {
		return false
	}
	return true
}

// writebarrier pass inserts write barriers for store ops (Store, Move, Zero)
// when necessary (the condition above). It rewrites store ops to branches
// and runtime calls, like
//
// if writeBarrier.enabled {
//   gcWriteBarrier(ptr, val)	// Not a regular Go call
// } else {
//   *ptr = val
// }
//
// A sequence of WB stores for many pointer fields of a single type will
// be emitted together, with a single branch.
func (psess *PackageSession) writebarrier(f *Func) {
	if !f.fe.UseWriteBarrier() {
		return
	}

	var sb, sp, wbaddr, const0 *Value
	var typedmemmove, typedmemclr, gcWriteBarrier *obj.LSym
	var stores, after []*Value
	var sset *sparseSet
	var storeNumber []int32

	for _, b := range f.Blocks {

		nWBops := 0
		for _, v := range b.Values {
			switch v.Op {
			case OpStore, OpMove, OpZero:
				if psess.needwb(v) {
					switch v.Op {
					case OpStore:
						v.Op = OpStoreWB
					case OpMove:
						v.Op = OpMoveWB
					case OpZero:
						v.Op = OpZeroWB
					}
					nWBops++
				}
			}
		}
		if nWBops == 0 {
			continue
		}

		if wbaddr == nil {

			initpos := f.Entry.Pos
			for _, v := range f.Entry.Values {
				if v.Op == OpSB {
					sb = v
				}
				if v.Op == OpSP {
					sp = v
				}
				if sb != nil && sp != nil {
					break
				}
			}
			if sb == nil {
				sb = f.Entry.NewValue0(initpos, OpSB, f.Config.Types.Uintptr)
			}
			if sp == nil {
				sp = f.Entry.NewValue0(initpos, OpSP, f.Config.Types.Uintptr)
			}
			wbsym := f.fe.Syslook("writeBarrier")
			wbaddr = f.Entry.NewValue1A(initpos, OpAddr, f.Config.Types.UInt32Ptr, wbsym, sb)
			gcWriteBarrier = f.fe.Syslook("gcWriteBarrier")
			typedmemmove = f.fe.Syslook("typedmemmove")
			typedmemclr = f.fe.Syslook("typedmemclr")
			const0 = f.ConstInt32(psess, f.Config.Types.UInt32, 0)

			sset = f.newSparseSet(f.NumValues())
			defer f.retSparseSet(sset)
			storeNumber = make([]int32, f.NumValues())
		}

		b.Values = psess.storeOrder(b.Values, sset, storeNumber)

		firstSplit := true
	again:
		// find the start and end of the last contiguous WB store sequence.
		// a branch will be inserted there. values after it will be moved
		// to a new block.
		var last *Value
		var start, end int
		values := b.Values
	FindSeq:
		for i := len(values) - 1; i >= 0; i-- {
			w := values[i]
			switch w.Op {
			case OpStoreWB, OpMoveWB, OpZeroWB:
				start = i
				if last == nil {
					last = w
					end = i + 1
				}
			case OpVarDef, OpVarLive, OpVarKill:
				continue
			default:
				if last == nil {
					continue
				}
				break FindSeq
			}
		}
		stores = append(stores[:0], b.Values[start:end]...)
		after = append(after[:0], b.Values[end:]...)
		b.Values = b.Values[:start]

		mem := stores[0].MemoryArg(psess)
		pos := stores[0].Pos
		bThen := f.NewBlock(BlockPlain)
		bElse := f.NewBlock(BlockPlain)
		bEnd := f.NewBlock(b.Kind)
		bThen.Pos = pos
		bElse.Pos = pos
		bEnd.Pos = b.Pos
		b.Pos = pos

		bEnd.SetControl(b.Control)
		bEnd.Likely = b.Likely
		for _, e := range b.Succs {
			bEnd.Succs = append(bEnd.Succs, e)
			e.b.Preds[e.i].b = bEnd
		}

		cfgtypes := &f.Config.Types
		flag := b.NewValue2(pos, OpLoad, cfgtypes.UInt32, wbaddr, mem)
		flag = b.NewValue2(pos, OpNeq32, cfgtypes.Bool, flag, const0)
		b.Kind = BlockIf
		b.SetControl(flag)
		b.Likely = BranchUnlikely
		b.Succs = b.Succs[:0]
		b.AddEdgeTo(bThen)
		b.AddEdgeTo(bElse)

		bThen.AddEdgeTo(bEnd)
		bElse.AddEdgeTo(bEnd)

		memThen := mem
		memElse := mem
		for _, w := range stores {
			ptr := w.Args[0]
			pos := w.Pos

			var fn *obj.LSym
			var typ *obj.LSym
			var val *Value
			switch w.Op {
			case OpStoreWB:
				val = w.Args[1]
				nWBops--
			case OpMoveWB:
				fn = typedmemmove
				val = w.Args[1]
				typ = w.Aux.(*types.Type).Symbol(psess.types)
				nWBops--
			case OpZeroWB:
				fn = typedmemclr
				typ = w.Aux.(*types.Type).Symbol(psess.types)
				nWBops--
			case OpVarDef, OpVarLive, OpVarKill:
			}

			switch w.Op {
			case OpStoreWB, OpMoveWB, OpZeroWB:
				volatile := w.Op == OpMoveWB && isVolatile(val)
				if w.Op == OpStoreWB {
					memThen = bThen.NewValue3A(pos, OpWB, psess.types.TypeMem, gcWriteBarrier, ptr, val, memThen)
				} else {
					memThen = psess.wbcall(pos, bThen, fn, typ, ptr, val, memThen, sp, sb, volatile)
				}

				f.fe.SetWBPos(pos)
			case OpVarDef, OpVarLive, OpVarKill:
				memThen = bThen.NewValue1A(pos, w.Op, psess.types.TypeMem, w.Aux, memThen)
			}

			switch w.Op {
			case OpStoreWB:
				memElse = bElse.NewValue3A(pos, OpStore, psess.types.TypeMem, w.Aux, ptr, val, memElse)
			case OpMoveWB:
				memElse = bElse.NewValue3I(pos, OpMove, psess.types.TypeMem, w.AuxInt, ptr, val, memElse)
				memElse.Aux = w.Aux
			case OpZeroWB:
				memElse = bElse.NewValue2I(pos, OpZero, psess.types.TypeMem, w.AuxInt, ptr, memElse)
				memElse.Aux = w.Aux
			case OpVarDef, OpVarLive, OpVarKill:
				memElse = bElse.NewValue1A(pos, w.Op, psess.types.TypeMem, w.Aux, memElse)
			}
		}

		bEnd.Values = append(bEnd.Values, last)
		last.Block = bEnd
		last.reset(OpPhi)
		last.Type = psess.types.TypeMem
		last.AddArg(memThen)
		last.AddArg(memElse)
		for _, w := range stores {
			if w != last {
				w.resetArgs()
			}
		}
		for _, w := range stores {
			if w != last {
				f.freeValue(psess, w)
			}
		}

		bEnd.Values = append(bEnd.Values, after...)
		for _, w := range after {
			w.Block = bEnd
		}

		if firstSplit {

			b.Func.WBLoads = append(b.Func.WBLoads, b)
			firstSplit = false
		} else {

			b.Func.WBLoads = append(b.Func.WBLoads, bEnd)
		}

		if nWBops > 0 {
			goto again
		}
	}
}

// wbcall emits write barrier runtime call in b, returns memory.
// if valIsVolatile, it moves val into temp space before making the call.
func (psess *PackageSession) wbcall(pos src.XPos, b *Block, fn, typ *obj.LSym, ptr, val, mem, sp, sb *Value, valIsVolatile bool) *Value {
	config := b.Func.Config

	var tmp GCNode
	if valIsVolatile {

		t := val.Type.Elem(psess.types)
		tmp = b.Func.fe.Auto(val.Pos, t)
		mem = b.NewValue1A(pos, OpVarDef, psess.types.TypeMem, tmp, mem)
		tmpaddr := b.NewValue1A(pos, OpAddr, t.PtrTo(psess.types), tmp, sp)
		siz := t.Size(psess.types)
		mem = b.NewValue3I(pos, OpMove, psess.types.TypeMem, siz, tmpaddr, val, mem)
		mem.Aux = t
		val = tmpaddr
	}

	off := config.ctxt.FixedFrameSize()

	if typ != nil {
		taddr := b.NewValue1A(pos, OpAddr, b.Func.Config.Types.Uintptr, typ, sb)
		off = round(off, taddr.Type.Alignment(psess.types))
		arg := b.NewValue1I(pos, OpOffPtr, taddr.Type.PtrTo(psess.types), off, sp)
		mem = b.NewValue3A(pos, OpStore, psess.types.TypeMem, ptr.Type, arg, taddr, mem)
		off += taddr.Type.Size(psess.types)
	}

	off = round(off, ptr.Type.Alignment(psess.types))
	arg := b.NewValue1I(pos, OpOffPtr, ptr.Type.PtrTo(psess.types), off, sp)
	mem = b.NewValue3A(pos, OpStore, psess.types.TypeMem, ptr.Type, arg, ptr, mem)
	off += ptr.Type.Size(psess.types)

	if val != nil {
		off = round(off, val.Type.Alignment(psess.types))
		arg = b.NewValue1I(pos, OpOffPtr, val.Type.PtrTo(psess.types), off, sp)
		mem = b.NewValue3A(pos, OpStore, psess.types.TypeMem, val.Type, arg, val, mem)
		off += val.Type.Size(psess.types)
	}
	off = round(off, config.PtrSize)

	mem = b.NewValue1A(pos, OpStaticCall, psess.types.TypeMem, fn, mem)
	mem.AuxInt = off - config.ctxt.FixedFrameSize()

	if valIsVolatile {
		mem = b.NewValue1A(pos, OpVarKill, psess.types.TypeMem, tmp, mem)
	}

	return mem
}

// round to a multiple of r, r is a power of 2
func round(o int64, r int64) int64 {
	return (o + r - 1) &^ (r - 1)
}

// IsStackAddr returns whether v is known to be an address of a stack slot
func IsStackAddr(v *Value) bool {
	for v.Op == OpOffPtr || v.Op == OpAddPtr || v.Op == OpPtrIndex || v.Op == OpCopy {
		v = v.Args[0]
	}
	switch v.Op {
	case OpSP:
		return true
	case OpAddr:
		return v.Args[0].Op == OpSP
	}
	return false
}

// IsSanitizerSafeAddr reports whether v is known to be an address
// that doesn't need instrumentation.
func IsSanitizerSafeAddr(v *Value) bool {
	for v.Op == OpOffPtr || v.Op == OpAddPtr || v.Op == OpPtrIndex || v.Op == OpCopy {
		v = v.Args[0]
	}
	switch v.Op {
	case OpSP:

		return true
	case OpITab, OpStringPtr, OpGetClosurePtr:

		return true
	case OpAddr:
		switch v.Args[0].Op {
		case OpSP:
			return true
		case OpSB:
			sym := v.Aux.(*obj.LSym)

			if strings.HasPrefix(sym.Name, "\"\".statictmp_") {
				return true
			}
		}
	}
	return false
}

// isVolatile returns whether v is a pointer to argument region on stack which
// will be clobbered by a function call.
func isVolatile(v *Value) bool {
	for v.Op == OpOffPtr || v.Op == OpAddPtr || v.Op == OpPtrIndex || v.Op == OpCopy {
		v = v.Args[0]
	}
	return v.Op == OpSP
}

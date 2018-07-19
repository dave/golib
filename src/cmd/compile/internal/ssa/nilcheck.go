package ssa

import (
	"github.com/dave/golib/src/cmd/internal/src"
)

// nilcheckelim eliminates unnecessary nil checks.
// runs on machine-independent code.
func (psess *PackageSession) nilcheckelim(f *Func) {

	sdom := f.sdom()

	type walkState int
	const (
		Work     walkState = iota // process nil checks and traverse to dominees
		ClearPtr                  // forget the fact that ptr is nil
	)

	type bp struct {
		block *Block // block, or nil in ClearPtr state
		ptr   *Value // if non-nil, ptr that is to be cleared in ClearPtr state
		op    walkState
	}

	work := make([]bp, 0, 256)
	work = append(work, bp{block: f.Entry})

	nonNilValues := make([]bool, f.NumValues())

	for _, b := range f.Blocks {
		for _, v := range b.Values {

			if v.Op == OpAddr || v.Op == OpAddPtr {
				nonNilValues[v.ID] = true
			}
		}
	}

	for changed := true; changed; {
		changed = false
		for _, b := range f.Blocks {
			for _, v := range b.Values {

				if v.Op == OpPhi {
					argsNonNil := true
					for _, a := range v.Args {
						if !nonNilValues[a.ID] {
							argsNonNil = false
							break
						}
					}
					if argsNonNil {
						if !nonNilValues[v.ID] {
							changed = true
						}
						nonNilValues[v.ID] = true
					}
				}
			}
		}
	}

	sset := f.newSparseSet(f.NumValues())
	defer f.retSparseSet(sset)
	storeNumber := make([]int32, f.NumValues())

	for len(work) > 0 {
		node := work[len(work)-1]
		work = work[:len(work)-1]

		switch node.op {
		case Work:
			b := node.block

			if len(b.Preds) == 1 {
				p := b.Preds[0].b
				if p.Kind == BlockIf && p.Control.Op == OpIsNonNil && p.Succs[0].b == b {
					ptr := p.Control.Args[0]
					if !nonNilValues[ptr.ID] {
						nonNilValues[ptr.ID] = true
						work = append(work, bp{op: ClearPtr, ptr: ptr})
					}
				}
			}

			b.Values = psess.storeOrder(b.Values, sset, storeNumber)

			pendingLines := f.cachedLineStarts
			pendingLines.clear()

			i := 0
			for _, v := range b.Values {
				b.Values[i] = v
				i++
				switch v.Op {
				case OpIsNonNil:
					ptr := v.Args[0]
					if nonNilValues[ptr.ID] {
						if v.Pos.IsStmt() == src.PosIsStmt {
							pendingLines.add(psess, v.Pos.Line())
							v.Pos = v.Pos.WithNotStmt()
						}

						v.reset(OpConstBool)
						v.AuxInt = 1
					}
				case OpNilCheck:
					ptr := v.Args[0]
					if nonNilValues[ptr.ID] {

						if f.fe.Debug_checknil() && v.Pos.Line() > 1 {
							f.Warnl(v.Pos, "removed nil check")
						}
						if v.Pos.IsStmt() == src.PosIsStmt {
							pendingLines.add(psess, v.Pos.Line())
						}
						v.reset(OpUnknown)
						f.freeValue(psess, v)
						i--
						continue
					}

					nonNilValues[ptr.ID] = true
					work = append(work, bp{op: ClearPtr, ptr: ptr})
					fallthrough
				default:
					if pendingLines.contains(v.Pos.Line()) && v.Pos.IsStmt() != src.PosNotStmt {
						v.Pos = v.Pos.WithIsStmt()
						pendingLines.remove(v.Pos.Line())
					}
				}
			}
			if pendingLines.contains(b.Pos.Line()) {
				b.Pos = b.Pos.WithIsStmt()
				pendingLines.remove(b.Pos.Line())
			}
			for j := i; j < len(b.Values); j++ {
				b.Values[j] = nil
			}
			b.Values = b.Values[:i]

			for w := sdom[node.block.ID].child; w != nil; w = sdom[w.ID].sibling {
				work = append(work, bp{op: Work, block: w})
			}

		case ClearPtr:
			nonNilValues[node.ptr.ID] = false
			continue
		}
	}
}

// All platforms are guaranteed to fault if we load/store to anything smaller than this address.
//
// This should agree with minLegalPointer in the runtime.
const minZeroPage = 4096

// nilcheckelim2 eliminates unnecessary nil checks.
// Runs after lowering and scheduling.
func (psess *PackageSession) nilcheckelim2(f *Func) {
	unnecessary := f.newSparseSet(f.NumValues())
	defer f.retSparseSet(unnecessary)

	pendingLines := f.cachedLineStarts

	for _, b := range f.Blocks {

		unnecessary.clear()
		pendingLines.clear()

		firstToRemove := len(b.Values)
		for i := len(b.Values) - 1; i >= 0; i-- {
			v := b.Values[i]
			if psess.opcodeTable[v.Op].nilCheck && unnecessary.contains(v.Args[0].ID) {
				if f.fe.Debug_checknil() && v.Pos.Line() > 1 {
					f.Warnl(v.Pos, "removed nil check")
				}
				if v.Pos.IsStmt() == src.PosIsStmt {
					pendingLines.add(psess, v.Pos.Line())
				}
				v.reset(OpUnknown)
				firstToRemove = i
				continue
			}
			if v.Type.IsMemory(psess.types) || v.Type.IsTuple() && v.Type.FieldType(psess.types, 1).IsMemory(psess.types) {
				if v.Op == OpVarDef || v.Op == OpVarKill || v.Op == OpVarLive {

					continue
				}

				unnecessary.clear()
			}

			// Find any pointers that this op is guaranteed to fault on if nil.
			var ptrstore [2]*Value
			ptrs := ptrstore[:0]
			if psess.opcodeTable[v.Op].faultOnNilArg0 {
				ptrs = append(ptrs, v.Args[0])
			}
			if psess.opcodeTable[v.Op].faultOnNilArg1 {
				ptrs = append(ptrs, v.Args[1])
			}
			for _, ptr := range ptrs {

				switch psess.opcodeTable[v.Op].auxType {
				case auxSymOff:
					if v.Aux != nil || v.AuxInt < 0 || v.AuxInt >= minZeroPage {
						continue
					}
				case auxSymValAndOff:
					off := ValAndOff(v.AuxInt).Off()
					if v.Aux != nil || off < 0 || off >= minZeroPage {
						continue
					}
				case auxInt32:

				case auxInt64:

				case auxNone:

				default:
					v.Fatalf("can't handle aux %s (type %d) yet\n", v.auxString(psess), int(psess.opcodeTable[v.Op].auxType))
				}

				unnecessary.add(ptr.ID)
			}
		}

		i := firstToRemove
		for j := i; j < len(b.Values); j++ {
			v := b.Values[j]
			if v.Op != OpUnknown {
				if v.Pos.IsStmt() != src.PosNotStmt && pendingLines.contains(v.Pos.Line()) {
					v.Pos = v.Pos.WithIsStmt()
					pendingLines.remove(v.Pos.Line())
				}
				b.Values[i] = v
				i++
			}
		}

		if pendingLines.contains(b.Pos.Line()) {
			b.Pos = b.Pos.WithIsStmt()
		}

		for j := i; j < len(b.Values); j++ {
			b.Values[j] = nil
		}
		b.Values = b.Values[:i]

	}
}

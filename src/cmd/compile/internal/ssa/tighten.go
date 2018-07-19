package ssa

// tighten moves Values closer to the Blocks in which they are used.
// This can reduce the amount of register spilling required,
// if it doesn't also create more live values.
// A Value can be moved to any block that
// dominates all blocks in which it is used.
func (psess *PackageSession) tighten(f *Func) {
	canMove := make([]bool, f.NumValues())
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			switch v.Op {
			case OpPhi, OpArg, OpSelect0, OpSelect1,
				OpAMD64LoweredGetClosurePtr, Op386LoweredGetClosurePtr,
				OpARMLoweredGetClosurePtr, OpARM64LoweredGetClosurePtr,
				OpMIPSLoweredGetClosurePtr, OpMIPS64LoweredGetClosurePtr,
				OpS390XLoweredGetClosurePtr, OpPPC64LoweredGetClosurePtr,
				OpWasmLoweredGetClosurePtr:

				continue
			}
			if v.MemoryArg(psess) != nil {

				continue
			}

			narg := 0
			for _, a := range v.Args {
				if !a.rematerializeable(psess) {
					narg++
				}
			}
			if narg >= 2 && !v.Type.IsFlags(psess.types) {

				continue
			}
			canMove[v.ID] = true
		}
	}

	lca := makeLCArange(f)

	target := make([]*Block, f.NumValues())

	idom := f.Idom()
	loops := f.loopnest(psess)
	loops.calculateDepths()

	changed := true
	for changed {
		changed = false

		for i := range target {
			target[i] = nil
		}

		for _, b := range f.Blocks {
			for _, v := range b.Values {
				for i, a := range v.Args {
					if !canMove[a.ID] {
						continue
					}
					use := b
					if v.Op == OpPhi {
						use = b.Preds[i].b
					}
					if target[a.ID] == nil {
						target[a.ID] = use
					} else {
						target[a.ID] = lca.find(target[a.ID], use)
					}
				}
			}
			if c := b.Control; c != nil {
				if !canMove[c.ID] {
					continue
				}
				if target[c.ID] == nil {
					target[c.ID] = b
				} else {
					target[c.ID] = lca.find(target[c.ID], b)
				}
			}
		}

		for _, b := range f.Blocks {
			origloop := loops.b2l[b.ID]
			for _, v := range b.Values {
				t := target[v.ID]
				if t == nil {
					continue
				}
				targetloop := loops.b2l[t.ID]
				for targetloop != nil && (origloop == nil || targetloop.depth > origloop.depth) {
					t = idom[targetloop.header.ID]
					target[v.ID] = t
					targetloop = loops.b2l[t.ID]
				}
			}
		}

		for _, b := range f.Blocks {
			for i := 0; i < len(b.Values); i++ {
				v := b.Values[i]
				t := target[v.ID]
				if t == nil || t == b {

					continue
				}

				t.Values = append(t.Values, v)
				v.Block = t
				last := len(b.Values) - 1
				b.Values[i] = b.Values[last]
				b.Values[last] = nil
				b.Values = b.Values[:last]
				changed = true
				i--
			}
		}
	}
}

// phiTighten moves constants closer to phi users.
// This pass avoids having lots of constants live for lots of the program.
// See issue 16407.
func (psess *PackageSession) phiTighten(f *Func) {
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Op != OpPhi {
				continue
			}
			for i, a := range v.Args {
				if !a.rematerializeable(psess) {
					continue
				}
				if a.Block == b.Preds[i].b {
					continue
				}

				v.SetArg(i, a.copyInto(psess, b.Preds[i].b))
			}
		}
	}
}

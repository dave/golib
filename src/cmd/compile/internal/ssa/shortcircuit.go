package ssa

// Shortcircuit finds situations where branch directions
// are always correlated and rewrites the CFG to take
// advantage of that fact.
// This optimization is useful for compiling && and || expressions.
func (psess *PackageSession) shortcircuit(f *Func) {
	// Step 1: Replace a phi arg with a constant if that arg
	// is the control value of a preceding If block.
	// b1:
	//    If a goto b2 else b3
	// b2: <- b1 ...
	//    x = phi(a, ...)
	//
	// We can replace the "a" in the phi with the constant true.
	var ct, cf *Value
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Op != OpPhi {
				continue
			}
			if !v.Type.IsBoolean() {
				continue
			}
			for i, a := range v.Args {
				e := b.Preds[i]
				p := e.b
				if p.Kind != BlockIf {
					continue
				}
				if p.Control != a {
					continue
				}
				if e.i == 0 {
					if ct == nil {
						ct = f.ConstBool(psess, f.Config.Types.Bool, true)
					}
					v.SetArg(i, ct)
				} else {
					if cf == nil {
						cf = f.ConstBool(psess, f.Config.Types.Bool, false)
					}
					v.SetArg(i, cf)
				}
			}
		}
	}

	live := make([]bool, f.NumValues())
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			for _, a := range v.Args {
				if a.Block != v.Block {
					live[a.ID] = true
				}
			}
		}
		if b.Control != nil && b.Control.Block != b {
			live[b.Control.ID] = true
		}
	}

	for _, b := range f.Blocks {
		if b.Kind != BlockIf {
			continue
		}
		if len(b.Values) != 1 {
			continue
		}
		v := b.Values[0]
		if v.Op != OpPhi {
			continue
		}
		if b.Control != v {
			continue
		}
		if live[v.ID] {
			continue
		}
		for i := 0; i < len(v.Args); i++ {
			a := v.Args[i]
			if a.Op != OpConstBool {
				continue
			}

			e1 := b.Preds[i]
			p := e1.b
			pi := e1.i

			e2 := b.Succs[1-a.AuxInt]
			t := e2.b
			ti := e2.i

			b.removePred(i)
			n := len(b.Preds)
			v.Args[i].Uses--
			v.Args[i] = v.Args[n]
			v.Args[n] = nil
			v.Args = v.Args[:n]

			p.Succs[pi] = Edge{t, len(t.Preds)}

			t.Preds = append(t.Preds, Edge{p, pi})
			for _, w := range t.Values {
				if w.Op != OpPhi {
					continue
				}
				w.AddArg(w.Args[ti])
			}

			if len(b.Preds) == 1 {
				v.Op = OpCopy

				break
			}
			i--
		}
	}
}

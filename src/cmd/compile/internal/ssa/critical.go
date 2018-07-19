package ssa

// critical splits critical edges (those that go from a block with
// more than one outedge to a block with more than one inedge).
// Regalloc wants a critical-edge-free CFG so it can implement phi values.
func critical(f *Func) {

	blocks := make([]*Block, f.NumValues())

	for j := 0; j < len(f.Blocks); j++ {
		b := f.Blocks[j]
		if len(b.Preds) <= 1 {
			continue
		}

		var phi *Value

		for _, v := range b.Values {
			if v.Op == OpPhi {
				if phi != nil {
					phi = nil
					break
				}
				phi = v
			}
		}

		if phi != nil {
			for _, v := range phi.Args {
				blocks[v.ID] = nil
			}
		}

		for i := 0; i < len(b.Preds); {
			e := b.Preds[i]
			p := e.b
			pi := e.i
			if p.Kind == BlockPlain {
				i++
				continue
			}

			var d *Block // new block used to remove critical edge
			reusedBlock := false
			if phi != nil {
				argID := phi.Args[i].ID

				if d = blocks[argID]; d == nil {

					d = f.NewBlock(BlockPlain)
					d.Pos = p.Pos
					blocks[argID] = d
					if f.pass.debug > 0 {
						f.Warnl(p.Pos, "split critical edge")
					}
				} else {
					reusedBlock = true
				}
			} else {

				d = f.NewBlock(BlockPlain)
				d.Pos = p.Pos
				if f.pass.debug > 0 {
					f.Warnl(p.Pos, "split critical edge")
				}
			}

			if reusedBlock {

				p.Succs[pi] = Edge{d, len(d.Preds)}
				d.Preds = append(d.Preds, Edge{p, pi})

				b.removePred(i)

				n := len(b.Preds)
				phi.Args[i].Uses--
				phi.Args[i] = phi.Args[n]
				phi.Args[n] = nil
				phi.Args = phi.Args[:n]

				if n == 1 {
					phi.Op = OpCopy
				}

			} else {

				p.Succs[pi] = Edge{d, 0}
				b.Preds[i] = Edge{d, 0}
				d.Preds = append(d.Preds, Edge{p, pi})
				d.Succs = append(d.Succs, Edge{b, i})
				i++
			}
		}
	}
}

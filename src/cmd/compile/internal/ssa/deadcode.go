package ssa

import (
	"github.com/dave/golib/src/cmd/internal/src"
)

// findlive returns the reachable blocks and live values in f.
func (psess *PackageSession) findlive(f *Func) (reachable []bool, live []bool) {
	reachable = ReachableBlocks(f)
	live, _ = psess.liveValues(f, reachable)
	return
}

// ReachableBlocks returns the reachable blocks in f.
func ReachableBlocks(f *Func) []bool {
	reachable := make([]bool, f.NumBlocks())
	reachable[f.Entry.ID] = true
	p := make([]*Block, 0, 64)
	p = append(p, f.Entry)
	for len(p) > 0 {

		b := p[len(p)-1]
		p = p[:len(p)-1]

		s := b.Succs
		if b.Kind == BlockFirst {
			s = s[:1]
		}
		for _, e := range s {
			c := e.b
			if int(c.ID) >= len(reachable) {
				f.Fatalf("block %s >= f.NumBlocks()=%d?", c, len(reachable))
			}
			if !reachable[c.ID] {
				reachable[c.ID] = true
				p = append(p, c)
			}
		}
	}
	return reachable
}

// liveValues returns the live values in f and a list of values that are eligible
// to be statements in reversed data flow order.
// The second result is used to help conserve statement boundaries for debugging.
// reachable is a map from block ID to whether the block is reachable.
func (psess *PackageSession) liveValues(f *Func, reachable []bool) (live []bool, liveOrderStmts []*Value) {
	live = make([]bool, f.NumValues())

	if f.RegAlloc != nil {
		for i := range live {
			live[i] = true
		}
		return
	}

	q := make([]*Value, 0, 64)

	for _, b := range f.Blocks {
		if !reachable[b.ID] {
			continue
		}
		if v := b.Control; v != nil && !live[v.ID] {
			live[v.ID] = true
			q = append(q, v)
			if v.Pos.IsStmt() != src.PosNotStmt {
				liveOrderStmts = append(liveOrderStmts, v)
			}
		}
		for _, v := range b.Values {
			if (psess.opcodeTable[v.Op].call || psess.opcodeTable[v.Op].hasSideEffects) && !live[v.ID] {
				live[v.ID] = true
				q = append(q, v)
				if v.Pos.IsStmt() != src.PosNotStmt {
					liveOrderStmts = append(liveOrderStmts, v)
				}
			}
			if v.Type.IsVoid(psess.types) && !live[v.ID] {

				live[v.ID] = true
				q = append(q, v)
				if v.Pos.IsStmt() != src.PosNotStmt {
					liveOrderStmts = append(liveOrderStmts, v)
				}
			}
		}
	}

	for len(q) > 0 {

		v := q[len(q)-1]
		q = q[:len(q)-1]
		for i, x := range v.Args {
			if v.Op == OpPhi && !reachable[v.Block.Preds[i].b.ID] {
				continue
			}
			if !live[x.ID] {
				live[x.ID] = true
				q = append(q, x)
				if x.Pos.IsStmt() != src.PosNotStmt {
					liveOrderStmts = append(liveOrderStmts, x)
				}
			}
		}
	}

	return
}

// deadcode removes dead code from f.
func (psess *PackageSession) deadcode(f *Func) {

	if f.RegAlloc != nil {
		f.Fatalf("deadcode after regalloc")
	}

	reachable := ReachableBlocks(f)

	for _, b := range f.Blocks {
		if reachable[b.ID] {
			continue
		}
		for i := 0; i < len(b.Succs); {
			e := b.Succs[i]
			if reachable[e.b.ID] {
				b.removeEdge(i)
			} else {
				i++
			}
		}
	}

	for _, b := range f.Blocks {
		if !reachable[b.ID] {
			continue
		}
		if b.Kind != BlockFirst {
			continue
		}
		b.removeEdge(1)
		b.Kind = BlockPlain
		b.Likely = BranchUnknown
	}

	copyelim(f)

	live, order := psess.liveValues(f, reachable)

	s := f.newSparseSet(f.NumValues())
	defer f.retSparseSet(s)
	i := 0
	for _, name := range f.Names {
		j := 0
		s.clear()
		values := f.NamedValues[name]
		for _, v := range values {
			if live[v.ID] && !s.contains(v.ID) {
				values[j] = v
				j++
				s.add(v.ID)
			}
		}
		if j == 0 {
			delete(f.NamedValues, name)
		} else {
			f.Names[i] = name
			i++
			for k := len(values) - 1; k >= j; k-- {
				values[k] = nil
			}
			f.NamedValues[name] = values[:j]
		}
	}
	for k := len(f.Names) - 1; k >= i; k-- {
		f.Names[k] = LocalSlot{}
	}
	f.Names = f.Names[:i]

	pendingLines := f.cachedLineStarts
	pendingLines.clear()

	for i, b := range f.Blocks {
		if !reachable[b.ID] {

			b.SetControl(nil)
		}
		for _, v := range b.Values {
			if !live[v.ID] {
				v.resetArgs()
				if v.Pos.IsStmt() == src.PosIsStmt && reachable[b.ID] {
					pendingLines.set(psess, v.Pos.Line(), int32(i))
				}
			}
		}
	}

	for i := len(order) - 1; i >= 0; i-- {
		w := order[i]
		if j := pendingLines.get(w.Pos.Line()); j > -1 && f.Blocks[j] == w.Block {
			w.Pos = w.Pos.WithIsStmt()
			pendingLines.remove(w.Pos.Line())
		}
	}

	for i := 0; i < pendingLines.size(); i++ {
		l, bi := pendingLines.getEntry(i)
		b := f.Blocks[bi]
		if b.Pos.Line() == l {
			b.Pos = b.Pos.WithIsStmt()
		}
	}

	for _, b := range f.Blocks {
		i := 0
		for _, v := range b.Values {
			if live[v.ID] {
				b.Values[i] = v
				i++
			} else {
				f.freeValue(psess, v)
			}
		}

		tail := b.Values[i:]
		for j := range tail {
			tail[j] = nil
		}
		b.Values = b.Values[:i]
	}

	i = 0
	for _, b := range f.WBLoads {
		if reachable[b.ID] {
			f.WBLoads[i] = b
			i++
		}
	}
	for j := i; j < len(f.WBLoads); j++ {
		f.WBLoads[j] = nil
	}
	f.WBLoads = f.WBLoads[:i]

	i = 0
	for _, b := range f.Blocks {
		if reachable[b.ID] {
			f.Blocks[i] = b
			i++
		} else {
			if len(b.Values) > 0 {
				b.Fatalf("live values in unreachable block %v: %v", b, b.Values)
			}
			f.freeBlock(b)
		}
	}

	tail := f.Blocks[i:]
	for j := range tail {
		tail[j] = nil
	}
	f.Blocks = f.Blocks[:i]
}

// removeEdge removes the i'th outgoing edge from b (and
// the corresponding incoming edge from b.Succs[i].b).
func (b *Block) removeEdge(i int) {
	e := b.Succs[i]
	c := e.b
	j := e.i

	b.removeSucc(i)

	c.removePred(j)

	n := len(c.Preds)
	for _, v := range c.Values {
		if v.Op != OpPhi {
			continue
		}
		v.Args[j].Uses--
		v.Args[j] = v.Args[n]
		v.Args[n] = nil
		v.Args = v.Args[:n]
		phielimValue(v)

	}
}

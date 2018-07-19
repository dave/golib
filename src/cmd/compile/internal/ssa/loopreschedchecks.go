package ssa

import (
	"fmt"
)

// an edgeMem records a backedge, together with the memory
// phi functions at the target of the backedge that must
// be updated when a rescheduling check replaces the backedge.
type edgeMem struct {
	e Edge
	m *Value // phi for memory at dest of e
}

// a rewriteTarget is a value-argindex pair indicating
// where a rewrite is applied.  Note that this is for values,
// not for block controls, because block controls are not targets
// for the rewrites performed in inserting rescheduling checks.
type rewriteTarget struct {
	v *Value
	i int
}

type rewrite struct {
	before, after *Value          // before is the expected value before rewrite, after is the new value installed.
	rewrites      []rewriteTarget // all the targets for this rewrite.
}

func (r *rewrite) String(psess *PackageSession) string {
	s := "\n\tbefore=" + r.before.String() + ", after=" + r.after.String()
	for _, rw := range r.rewrites {
		s += ", (i=" + fmt.Sprint(rw.i) + ", v=" + rw.v.LongString(psess) + ")"
	}
	s += "\n"
	return s
}

// insertLoopReschedChecks inserts rescheduling checks on loop backedges.
func (psess *PackageSession) insertLoopReschedChecks(f *Func) {

	if f.NoSplit {
		return
	}

	backedges := backedges(f)
	if len(backedges) == 0 {
		return
	}

	lastMems := psess.findLastMems(f)

	idom := f.Idom()
	po := f.postorder()

	sdom := newSparseOrderedTree(f, idom, po)

	if f.pass.debug > 1 {
		fmt.Printf("before %s = %s\n", f.Name, sdom.treestructure(f.Entry))
	}

	tofixBackedges := []edgeMem{}

	for _, e := range backedges {
		tofixBackedges = append(tofixBackedges, edgeMem{e, nil})
	}

	if lastMems[f.Entry.ID] == nil {
		lastMems[f.Entry.ID] = f.Entry.NewValue0(f.Entry.Pos, OpInitMem, psess.types.TypeMem)
	}

	memDefsAtBlockEnds := make([]*Value, f.NumBlocks())

	for i := len(po) - 1; i >= 0; i-- {
		b := po[i]
		mem := lastMems[b.ID]
		for j := 0; mem == nil; j++ {

			mem = memDefsAtBlockEnds[b.Preds[j].b.ID]
		}
		memDefsAtBlockEnds[b.ID] = mem
		if f.pass.debug > 2 {
			fmt.Printf("memDefsAtBlockEnds[%s] = %s\n", b, mem)
		}
	}

	newmemphis := make(map[*Block]rewrite)

	for i, emc := range tofixBackedges {
		e := emc.e
		h := e.b

		// find the phi function for the memory input at "h", if there is one.
		var headerMemPhi *Value // look for header mem phi

		for _, v := range h.Values {
			if v.Op == OpPhi && v.Type.IsMemory(psess.types) {
				headerMemPhi = v
			}
		}

		if headerMemPhi == nil {

			mem0 := memDefsAtBlockEnds[idom[h.ID].ID]
			headerMemPhi = newPhiFor(h, mem0)
			newmemphis[h] = rewrite{before: mem0, after: headerMemPhi}
			addDFphis(mem0, h, h, f, memDefsAtBlockEnds, newmemphis, sdom)

		}
		tofixBackedges[i].m = headerMemPhi

	}
	if f.pass.debug > 0 {
		for b, r := range newmemphis {
			fmt.Printf("before b=%s, rewrite=%s\n", b, r.String(psess))
		}
	}

	dfPhiTargets := make(map[rewriteTarget]bool)
	psess.
		rewriteNewPhis(f.Entry, f.Entry, f, memDefsAtBlockEnds, newmemphis, dfPhiTargets, sdom)

	if f.pass.debug > 0 {
		for b, r := range newmemphis {
			fmt.Printf("after b=%s, rewrite=%s\n", b, r.String(psess))
		}
	}

	for _, r := range newmemphis {
		for _, rw := range r.rewrites {
			rw.v.SetArg(rw.i, r.after)
		}
	}

	for _, emc := range tofixBackedges {
		e := emc.e
		headerMemPhi := emc.m
		h := e.b
		i := e.i
		p := h.Preds[i]
		bb := p.b
		mem0 := headerMemPhi.Args[i]

		likely := BranchLikely
		if p.i != 0 {
			likely = BranchUnlikely
		}
		if bb.Kind != BlockPlain {
			bb.Likely = likely
		}

		test := f.NewBlock(BlockIf)
		sched := f.NewBlock(BlockPlain)

		test.Pos = bb.Pos
		sched.Pos = bb.Pos

		cfgtypes := &f.Config.Types
		pt := cfgtypes.Uintptr
		g := test.NewValue1(bb.Pos, OpGetG, pt, mem0)
		sp := test.NewValue0(bb.Pos, OpSP, pt)
		cmpOp := OpLess64U
		if pt.Size(psess.types) == 4 {
			cmpOp = OpLess32U
		}
		limaddr := test.NewValue1I(bb.Pos, OpOffPtr, pt, 2*pt.Size(psess.types), g)
		lim := test.NewValue2(bb.Pos, OpLoad, pt, limaddr, mem0)
		cmp := test.NewValue2(bb.Pos, cmpOp, cfgtypes.Bool, sp, lim)
		test.SetControl(cmp)

		test.AddEdgeTo(sched)

		test.Succs = append(test.Succs, Edge{h, i})
		h.Preds[i] = Edge{test, 1}
		headerMemPhi.SetArg(i, mem0)

		test.Likely = BranchUnlikely

		resched := f.fe.Syslook("goschedguarded")
		mem1 := sched.NewValue1A(bb.Pos, OpStaticCall, psess.types.TypeMem, resched, mem0)
		sched.AddEdgeTo(h)
		headerMemPhi.AddArg(mem1)

		bb.Succs[p.i] = Edge{test, 0}
		test.Preds = append(test.Preds, Edge{bb, p.i})

		for _, v := range h.Values {
			if v.Op == OpPhi && v != headerMemPhi {
				v.AddArg(v.Args[i])
			}
		}
	}

	f.invalidateCFG()

	if f.pass.debug > 1 {
		sdom = newSparseTree(f, f.Idom())
		fmt.Printf("after %s = %s\n", f.Name, sdom.treestructure(f.Entry))
	}
}

// newPhiFor inserts a new Phi function into b,
// with all inputs set to v.
func newPhiFor(b *Block, v *Value) *Value {
	phiV := b.NewValue0(b.Pos, OpPhi, v.Type)

	for range b.Preds {
		phiV.AddArg(v)
	}
	return phiV
}

// rewriteNewPhis updates newphis[h] to record all places where the new phi function inserted
// in block h will replace a previous definition.  Block b is the block currently being processed;
// if b has its own phi definition then it takes the place of h.
// defsForUses provides information about other definitions of the variable that are present
// (and if nil, indicates that the variable is no longer live)
// sdom must yield a preorder of the flow graph if recursively walked, root-to-children.
// The result of newSparseOrderedTree with order supplied by a dfs-postorder satisfies this
// requirement.
func (psess *PackageSession) rewriteNewPhis(h, b *Block, f *Func, defsForUses []*Value, newphis map[*Block]rewrite, dfPhiTargets map[rewriteTarget]bool, sdom SparseTree) {

	if _, ok := newphis[b]; ok {
		h = b
	}
	change := newphis[h]
	x := change.before
	y := change.after

	if x != nil {
		p := &change.rewrites
		for _, v := range b.Values {
			if v == y {
				continue
			}
			for i, w := range v.Args {
				if w != x {
					continue
				}
				tgt := rewriteTarget{v, i}

				if dfPhiTargets[tgt] {
					continue
				}
				*p = append(*p, tgt)
				if f.pass.debug > 1 {
					fmt.Printf("added block target for h=%v, b=%v, x=%v, y=%v, tgt.v=%s, tgt.i=%d\n",
						h, b, x, y, v, i)
				}
			}
		}

		if dfu := defsForUses[b.ID]; dfu != nil && dfu.Block != b {
			for _, e := range b.Succs {
				s := e.b

				for _, v := range s.Values {
					if v.Op == OpPhi && v.Args[e.i] == x {
						tgt := rewriteTarget{v, e.i}
						*p = append(*p, tgt)
						dfPhiTargets[tgt] = true
						if f.pass.debug > 1 {
							fmt.Printf("added phi target for h=%v, b=%v, s=%v, x=%v, y=%v, tgt.v=%s, tgt.i=%d\n",
								h, b, s, x, y, v.LongString(psess), e.i)
						}
						break
					}
				}
			}
		}
		newphis[h] = change
	}

	for c := sdom[b.ID].child; c != nil; c = sdom[c.ID].sibling {
		psess.
			rewriteNewPhis(h, c, f, defsForUses, newphis, dfPhiTargets, sdom)
	}
}

// addDFphis creates new trivial phis that are necessary to correctly reflect (within SSA)
// a new definition for variable "x" inserted at h (usually but not necessarily a phi).
// These new phis can only occur at the dominance frontier of h; block s is in the dominance
// frontier of h if h does not strictly dominate s and if s is a successor of a block b where
// either b = h or h strictly dominates b.
// These newly created phis are themselves new definitions that may require addition of their
// own trivial phi functions in their own dominance frontier, and this is handled recursively.
func addDFphis(x *Value, h, b *Block, f *Func, defForUses []*Value, newphis map[*Block]rewrite, sdom SparseTree) {
	oldv := defForUses[b.ID]
	if oldv != x {
		return
	}
	idom := f.Idom()
outer:
	for _, e := range b.Succs {
		s := e.b

		if sdom.isAncestor(h, s) {
			continue
		}
		if _, ok := newphis[s]; ok {
			continue
		}
		if x != nil {
			for _, v := range s.Values {
				if v.Op == OpPhi && v.Args[e.i] == x {
					continue outer
				}
			}
		}

		old := defForUses[idom[s.ID].ID]
		headerPhi := newPhiFor(s, old)

		newphis[s] = rewrite{before: old, after: headerPhi}
		addDFphis(old, s, s, f, defForUses, newphis, sdom)
	}
	for c := sdom[b.ID].child; c != nil; c = sdom[c.ID].sibling {
		addDFphis(x, h, c, f, defForUses, newphis, sdom)
	}
}

// findLastMems maps block ids to last memory-output op in a block, if any
func (psess *PackageSession) findLastMems(f *Func) []*Value {

	var stores []*Value
	lastMems := make([]*Value, f.NumBlocks())
	storeUse := f.newSparseSet(f.NumValues())
	defer f.retSparseSet(storeUse)
	for _, b := range f.Blocks {

		storeUse.clear()
		stores = stores[:0]
		var memPhi *Value
		for _, v := range b.Values {
			if v.Op == OpPhi {
				if v.Type.IsMemory(psess.types) {
					memPhi = v
				}
				continue
			}
			if v.Type.IsMemory(psess.types) {
				stores = append(stores, v)
				for _, a := range v.Args {
					if a.Block == b && a.Type.IsMemory(psess.types) {
						storeUse.add(a.ID)
					}
				}
			}
		}
		if len(stores) == 0 {
			lastMems[b.ID] = memPhi
			continue
		}

		// find last store in the block
		var last *Value
		for _, v := range stores {
			if storeUse.contains(v.ID) {
				continue
			}
			if last != nil {
				b.Fatalf("two final stores - simultaneous live stores %s %s", last, v)
			}
			last = v
		}
		if last == nil {
			b.Fatalf("no last store found - cycle?")
		}
		lastMems[b.ID] = last
	}
	return lastMems
}

type backedgesState struct {
	b *Block
	i int
}

// backedges returns a slice of successor edges that are back
// edges.  For reducible loops, edge.b is the header.
func backedges(f *Func) []Edge {
	edges := []Edge{}
	mark := make([]markKind, f.NumBlocks())
	stack := []backedgesState{}

	mark[f.Entry.ID] = notExplored
	stack = append(stack, backedgesState{f.Entry, 0})

	for len(stack) > 0 {
		l := len(stack)
		x := stack[l-1]
		if x.i < len(x.b.Succs) {
			e := x.b.Succs[x.i]
			stack[l-1].i++
			s := e.b
			if mark[s.ID] == notFound {
				mark[s.ID] = notExplored
				stack = append(stack, backedgesState{s, 0})
			} else if mark[s.ID] == notExplored {
				edges = append(edges, e)
			}
		} else {
			mark[x.b.ID] = done
			stack = stack[0 : l-1]
		}
	}
	return edges
}

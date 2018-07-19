package ssa

import (
	"fmt"
	"math"
)

type branch int

const (
	unknown branch = iota
	positive
	negative
)

// relation represents the set of possible relations between
// pairs of variables (v, w). Without a priori knowledge the
// mask is lt | eq | gt meaning v can be less than, equal to or
// greater than w. When the execution path branches on the condition
// `v op w` the set of relations is updated to exclude any
// relation not possible due to `v op w` being true (or false).
//
// E.g.
//
// r := relation(...)
//
// if v < w {
//   newR := r & lt
// }
// if v >= w {
//   newR := r & (eq|gt)
// }
// if v != w {
//   newR := r & (lt|gt)
// }
type relation uint

const (
	lt relation = 1 << iota
	eq
	gt
)

func (r relation) String(psess *PackageSession) string {
	if r < relation(len(psess.relationStrings)) {
		return psess.relationStrings[r]
	}
	return fmt.Sprintf("relation(%d)", uint(r))
}

// domain represents the domain of a variable pair in which a set
// of relations is known.  For example, relations learned for unsigned
// pairs cannot be transferred to signed pairs because the same bit
// representation can mean something else.
type domain uint

const (
	signed domain = 1 << iota
	unsigned
	pointer
	boolean
)

func (d domain) String(psess *PackageSession) string {
	s := ""
	for i, ds := range psess.domainStrings {
		if d&(1<<uint(i)) != 0 {
			if len(s) != 0 {
				s += "|"
			}
			s += ds
			d &^= 1 << uint(i)
		}
	}
	if d != 0 {
		if len(s) != 0 {
			s += "|"
		}
		s += fmt.Sprintf("0x%x", uint(d))
	}
	return s
}

type pair struct {
	v, w *Value // a pair of values, ordered by ID.
	// v can be nil, to mean the zero value.
	// for booleans the zero value (v == nil) is false.
	d domain
}

// fact is a pair plus a relation for that pair.
type fact struct {
	p pair
	r relation
}

// a limit records known upper and lower bounds for a value.
type limit struct {
	min, max   int64  // min <= value <= max, signed
	umin, umax uint64 // umin <= value <= umax, unsigned
}

func (l limit) String() string {
	return fmt.Sprintf("sm,SM,um,UM=%d,%d,%d,%d", l.min, l.max, l.umin, l.umax)
}

func (l limit) intersect(l2 limit) limit {
	if l.min < l2.min {
		l.min = l2.min
	}
	if l.umin < l2.umin {
		l.umin = l2.umin
	}
	if l.max > l2.max {
		l.max = l2.max
	}
	if l.umax > l2.umax {
		l.umax = l2.umax
	}
	return l
}

// a limitFact is a limit known for a particular value.
type limitFact struct {
	vid   ID
	limit limit
}

// factsTable keeps track of relations between pairs of values.
//
// The fact table logic is sound, but incomplete. Outside of a few
// special cases, it performs no deduction or arithmetic. While there
// are known decision procedures for this, the ad hoc approach taken
// by the facts table is effective for real code while remaining very
// efficient.
type factsTable struct {
	// unsat is true if facts contains a contradiction.
	//
	// Note that the factsTable logic is incomplete, so if unsat
	// is false, the assertions in factsTable could be satisfiable
	// *or* unsatisfiable.
	unsat      bool // true if facts contains a contradiction
	unsatDepth int  // number of unsat checkpoints

	facts map[pair]relation // current known set of relation
	stack []fact            // previous sets of relations

	// order is a couple of partial order sets that record information
	// about relations between SSA values in the signed and unsigned
	// domain.
	order [2]*poset

	// known lower and upper bounds on individual values.
	limits     map[ID]limit
	limitStack []limitFact // previous entries

	// For each slice s, a map from s to a len(s)/cap(s) value (if any)
	// TODO: check if there are cases that matter where we have
	// more than one len(s) for a slice. We could keep a list if necessary.
	lens map[ID]*Value
	caps map[ID]*Value
}

// checkpointFact is an invalid value used for checkpointing
// and restoring factsTable.

func newFactsTable(f *Func) *factsTable {
	ft := &factsTable{}
	ft.order[0] = f.newPoset()
	ft.order[1] = f.newPoset()
	ft.order[0].SetUnsigned(false)
	ft.order[1].SetUnsigned(true)
	ft.facts = make(map[pair]relation)
	ft.stack = make([]fact, 4)
	ft.limits = make(map[ID]limit)
	ft.limitStack = make([]limitFact, 4)
	return ft
}

// update updates the set of relations between v and w in domain d
// restricting it to r.
func (ft *factsTable) update(psess *PackageSession, parent *Block, v, w *Value, d domain, r relation) {

	if ft.unsat {
		return
	}

	if v == w {
		if r&eq == 0 {
			ft.unsat = true
		}
		return
	}

	if d == signed || d == unsigned {
		var ok bool
		idx := 0
		if d == unsigned {
			idx = 1
		}
		switch r {
		case lt:
			ok = ft.order[idx].SetOrder(v, w)
		case gt:
			ok = ft.order[idx].SetOrder(w, v)
		case lt | eq:
			ok = ft.order[idx].SetOrderOrEqual(v, w)
		case gt | eq:
			ok = ft.order[idx].SetOrderOrEqual(w, v)
		case eq:
			ok = ft.order[idx].SetEqual(v, w)
		case lt | gt:
			ok = ft.order[idx].SetNonEqual(v, w)
		default:
			panic("unknown relation")
		}
		if !ok {
			ft.unsat = true
			return
		}
	} else {
		if lessByID(w, v) {
			v, w = w, v
			r = psess.reverseBits[r]
		}

		p := pair{v, w, d}
		oldR, ok := ft.facts[p]
		if !ok {
			if v == w {
				oldR = eq
			} else {
				oldR = lt | eq | gt
			}
		}

		if oldR == r {
			return
		}
		ft.stack = append(ft.stack, fact{p, oldR})
		ft.facts[p] = oldR & r

		if oldR&r == 0 {
			ft.unsat = true
			return
		}
	}

	if v.isGenericIntConst() {
		v, w = w, v
		r = psess.reverseBits[r]
	}
	if v != nil && w.isGenericIntConst() {

		old, ok := ft.limits[v.ID]
		if !ok {
			old = psess.noLimit
			if v.isGenericIntConst() {
				switch d {
				case signed:
					old.min, old.max = v.AuxInt, v.AuxInt
					if v.AuxInt >= 0 {
						old.umin, old.umax = uint64(v.AuxInt), uint64(v.AuxInt)
					}
				case unsigned:
					old.umin = v.AuxUnsigned()
					old.umax = old.umin
					if int64(old.umin) >= 0 {
						old.min, old.max = int64(old.umin), int64(old.umin)
					}
				}
			}
		}
		lim := psess.noLimit
		switch d {
		case signed:
			c := w.AuxInt
			switch r {
			case lt:
				lim.max = c - 1
			case lt | eq:
				lim.max = c
			case gt | eq:
				lim.min = c
			case gt:
				lim.min = c + 1
			case lt | gt:
				lim = old
				if c == lim.min {
					lim.min++
				}
				if c == lim.max {
					lim.max--
				}
			case eq:
				lim.min = c
				lim.max = c
			}
			if lim.min >= 0 {

				lim.umin = uint64(lim.min)
			}
			if lim.max != psess.noLimit.max && old.min >= 0 && lim.max >= 0 {

				lim.umax = uint64(lim.max)
			}
		case unsigned:
			uc := w.AuxUnsigned()
			switch r {
			case lt:
				lim.umax = uc - 1
			case lt | eq:
				lim.umax = uc
			case gt | eq:
				lim.umin = uc
			case gt:
				lim.umin = uc + 1
			case lt | gt:
				lim = old
				if uc == lim.umin {
					lim.umin++
				}
				if uc == lim.umax {
					lim.umax--
				}
			case eq:
				lim.umin = uc
				lim.umax = uc
			}

		}
		ft.limitStack = append(ft.limitStack, limitFact{v.ID, old})
		lim = old.intersect(lim)
		ft.limits[v.ID] = lim
		if v.Block.Func.pass.debug > 2 {
			v.Block.Func.Warnl(parent.Pos, "parent=%s, new limits %s %s %s", parent, v, w, lim.String())
		}
		if lim.min > lim.max || lim.umin > lim.umax {
			ft.unsat = true
			return
		}
	}

	if d != signed && d != unsigned {
		return
	}

	if v.Op == OpSliceLen && r&lt == 0 && ft.caps[v.Args[0].ID] != nil {

		ft.update(psess, parent, ft.caps[v.Args[0].ID], w, d, r|gt)
	}
	if w.Op == OpSliceLen && r&gt == 0 && ft.caps[w.Args[0].ID] != nil {

		ft.update(psess, parent, v, ft.caps[w.Args[0].ID], d, r|lt)
	}
	if v.Op == OpSliceCap && r&gt == 0 && ft.lens[v.Args[0].ID] != nil {

		ft.update(psess, parent, ft.lens[v.Args[0].ID], w, d, r|lt)
	}
	if w.Op == OpSliceCap && r&lt == 0 && ft.lens[w.Args[0].ID] != nil {

		ft.update(psess, parent, v, ft.lens[w.Args[0].ID], d, r|gt)
	}

	if r == lt || r == lt|eq {
		v, w = w, v
		r = psess.reverseBits[r]
	}
	switch r {
	case gt:
		if x, delta := isConstDelta(v); x != nil && delta == 1 {

			ft.update(psess, parent, x, w, d, gt|eq)
		} else if x, delta := isConstDelta(w); x != nil && delta == -1 {

			ft.update(psess, parent, v, x, d, gt|eq)
		}
	case gt | eq:
		if x, delta := isConstDelta(v); x != nil && delta == -1 {

			lim, ok := ft.limits[x.ID]
			if ok && lim.min > psess.opMin[v.Op] {
				ft.update(psess, parent, x, w, d, gt)
			}
		} else if x, delta := isConstDelta(w); x != nil && delta == 1 {

			lim, ok := ft.limits[x.ID]
			if ok && lim.max < psess.opMax[w.Op] {
				ft.update(psess, parent, v, x, d, gt)
			}
		}
	}

	if r == gt || r == gt|eq {
		if x, delta := isConstDelta(v); x != nil && d == signed {
			if parent.Func.pass.debug > 1 {
				parent.Func.Warnl(parent.Pos, "x+d >= w; x:%v %v delta:%v w:%v d:%v", x, parent.String(), delta, w.AuxInt, d)
			}
			if !w.isGenericIntConst() {

				if l, has := ft.limits[x.ID]; has && delta < 0 {
					if (x.Type.Size(psess.types) == 8 && l.min >= math.MinInt64-delta) ||
						(x.Type.Size(psess.types) == 4 && l.min >= math.MinInt32-delta) {
						ft.update(psess, parent, x, w, signed, r)
					}
				}
			} else {
				// With w,delta constants, we want to derive: x+delta > w  ⇒  x > w-delta
				//
				// We compute (using integers of the correct size):
				//    min = w - delta
				//    max = MaxInt - delta
				//
				// And we prove that:
				//    if min<max: min < x AND x <= max
				//    if min>max: min < x OR  x <= max
				//
				// This is always correct, even in case of overflow.
				//
				// If the initial fact is x+delta >= w instead, the derived conditions are:
				//    if min<max: min <= x AND x <= max
				//    if min>max: min <= x OR  x <= max
				//
				// Notice the conditions for max are still <=, as they handle overflows.
				var min, max int64
				var vmin, vmax *Value
				switch x.Type.Size(psess.types) {
				case 8:
					min = w.AuxInt - delta
					max = int64(^uint64(0)>>1) - delta

					vmin = parent.NewValue0I(parent.Pos, OpConst64, parent.Func.Config.Types.Int64, min)
					vmax = parent.NewValue0I(parent.Pos, OpConst64, parent.Func.Config.Types.Int64, max)

				case 4:
					min = int64(int32(w.AuxInt) - int32(delta))
					max = int64(int32(^uint32(0)>>1) - int32(delta))

					vmin = parent.NewValue0I(parent.Pos, OpConst32, parent.Func.Config.Types.Int32, min)
					vmax = parent.NewValue0I(parent.Pos, OpConst32, parent.Func.Config.Types.Int32, max)

				default:
					panic("unimplemented")
				}

				if min < max {

					ft.update(psess, parent, x, vmin, d, r)
					ft.update(psess, parent, vmax, x, d, r|eq)
				} else {

					if l, has := ft.limits[x.ID]; has {
						if l.max <= min {

							ft.update(psess, parent, vmax, x, d, r|eq)
						} else if l.min > max {

							ft.update(psess, parent, x, vmin, d, r)
						}
					}
				}
			}
		}
	}

}

// isNonNegative reports whether v is known to be non-negative.
func (ft *factsTable) isNonNegative(v *Value) bool {
	if isNonNegative(v) {
		return true
	}

	if l, has := ft.limits[v.ID]; has && (l.min >= 0 || l.umax <= math.MaxInt64) {
		return true
	}

	if x, delta := isConstDelta(v); x != nil {
		if l, has := ft.limits[x.ID]; has {
			if delta > 0 && l.min >= -delta && l.max <= math.MaxInt64-delta {
				return true
			}
			if delta < 0 && l.min >= -delta {
				return true
			}
		}
	}

	return false
}

// checkpoint saves the current state of known relations.
// Called when descending on a branch.
func (ft *factsTable) checkpoint(psess *PackageSession) {
	if ft.unsat {
		ft.unsatDepth++
	}
	ft.stack = append(ft.stack, psess.checkpointFact)
	ft.limitStack = append(ft.limitStack, psess.checkpointBound)
	ft.order[0].Checkpoint()
	ft.order[1].Checkpoint()
}

// restore restores known relation to the state just
// before the previous checkpoint.
// Called when backing up on a branch.
func (ft *factsTable) restore(psess *PackageSession) {
	if ft.unsatDepth > 0 {
		ft.unsatDepth--
	} else {
		ft.unsat = false
	}
	for {
		old := ft.stack[len(ft.stack)-1]
		ft.stack = ft.stack[:len(ft.stack)-1]
		if old == psess.checkpointFact {
			break
		}
		if old.r == lt|eq|gt {
			delete(ft.facts, old.p)
		} else {
			ft.facts[old.p] = old.r
		}
	}
	for {
		old := ft.limitStack[len(ft.limitStack)-1]
		ft.limitStack = ft.limitStack[:len(ft.limitStack)-1]
		if old.vid == 0 {
			break
		}
		if old.limit == psess.noLimit {
			delete(ft.limits, old.vid)
		} else {
			ft.limits[old.vid] = old.limit
		}
	}
	ft.order[0].Undo()
	ft.order[1].Undo()
}

func lessByID(v, w *Value) bool {
	if v == nil && w == nil {

		return false
	}
	if v == nil {
		return true
	}
	return w != nil && v.ID < w.ID
}

// maps what we learn when the positive branch is taken.
// For example:
//      OpLess8:   {signed, lt},
//	v1 = (OpLess8 v2 v3).
// If v1 branch is taken than we learn that the rangeMaks
// can be at most lt.

// prove removes redundant BlockIf branches that can be inferred
// from previous dominating comparisons.
//
// By far, the most common redundant pair are generated by bounds checking.
// For example for the code:
//
//    a[i] = 4
//    foo(a[i])
//
// The compiler will generate the following code:
//
//    if i >= len(a) {
//        panic("not in bounds")
//    }
//    a[i] = 4
//    if i >= len(a) {
//        panic("not in bounds")
//    }
//    foo(a[i])
//
// The second comparison i >= len(a) is clearly redundant because if the
// else branch of the first comparison is executed, we already know that i < len(a).
// The code for the second panic can be removed.
//
// prove works by finding contradictions and trimming branches whose
// conditions are unsatisfiable given the branches leading up to them.
// It tracks a "fact table" of branch conditions. For each branching
// block, it asserts the branch conditions that uniquely dominate that
// block, and then separately asserts the block's branch condition and
// its negation. If either leads to a contradiction, it can trim that
// successor.
func (psess *PackageSession) prove(f *Func) {
	ft := newFactsTable(f)
	ft.checkpoint(psess)

	// Find length and capacity ops.
	var zero *Value
	for _, b := range f.Blocks {
		for _, v := range b.Values {

			if zero == nil && v.Op == OpConst64 && v.AuxInt == 0 {
				zero = v
			}
			if v.Uses == 0 {

				continue
			}
			switch v.Op {
			case OpStringLen:
				if zero == nil {
					zero = b.NewValue0I(b.Pos, OpConst64, f.Config.Types.Int64, 0)
				}
				ft.update(psess, b, v, zero, signed, gt|eq)
			case OpSliceLen:
				if ft.lens == nil {
					ft.lens = map[ID]*Value{}
				}
				ft.lens[v.Args[0].ID] = v
				if zero == nil {
					zero = b.NewValue0I(b.Pos, OpConst64, f.Config.Types.Int64, 0)
				}
				ft.update(psess, b, v, zero, signed, gt|eq)
			case OpSliceCap:
				if ft.caps == nil {
					ft.caps = map[ID]*Value{}
				}
				ft.caps[v.Args[0].ID] = v
				if zero == nil {
					zero = b.NewValue0I(b.Pos, OpConst64, f.Config.Types.Int64, 0)
				}
				ft.update(psess, b, v, zero, signed, gt|eq)
			}
		}
	}

	// Find induction variables. Currently, findIndVars
	// is limited to one induction variable per block.
	var indVars map[*Block]indVar
	for _, v := range findIndVar(f) {
		if indVars == nil {
			indVars = make(map[*Block]indVar)
		}
		indVars[v.entry] = v
	}

	// current node state
	type walkState int
	const (
		descend walkState = iota
		simplify
	)
	// work maintains the DFS stack.
	type bp struct {
		block *Block    // current handled block
		state walkState // what's to do
	}
	work := make([]bp, 0, 256)
	work = append(work, bp{
		block: f.Entry,
		state: descend,
	})

	idom := f.Idom()
	sdom := f.sdom()

	for len(work) > 0 {
		node := work[len(work)-1]
		work = work[:len(work)-1]
		parent := idom[node.block.ID]
		branch := getBranch(sdom, parent, node.block)

		switch node.state {
		case descend:
			ft.checkpoint(psess)
			if iv, ok := indVars[node.block]; ok {
				psess.
					addIndVarRestrictions(ft, parent, iv)
			}

			if branch != unknown {
				psess.
					addBranchRestrictions(ft, parent, branch)
				if ft.unsat {

					removeBranch(parent, branch)
					ft.restore(psess)
					break
				}

			}
			psess.
				addLocalInductiveFacts(ft, node.block)

			work = append(work, bp{
				block: node.block,
				state: simplify,
			})
			for s := sdom.Child(node.block); s != nil; s = sdom.Sibling(s) {
				work = append(work, bp{
					block: s,
					state: descend,
				})
			}

		case simplify:
			psess.
				simplifyBlock(sdom, ft, node.block)
			ft.restore(psess)
		}
	}

	ft.restore(psess)

	for _, po := range ft.order {

		if psess.checkEnabled {
			if err := po.CheckEmpty(); err != nil {
				f.Fatalf("prove poset not empty after function %s: %v", f.Name, err)
			}
		}
		f.retPoset(po)
	}
}

// getBranch returns the range restrictions added by p
// when reaching b. p is the immediate dominator of b.
func getBranch(sdom SparseTree, p *Block, b *Block) branch {
	if p == nil || p.Kind != BlockIf {
		return unknown
	}

	if sdom.isAncestorEq(p.Succs[0].b, b) && len(p.Succs[0].b.Preds) == 1 {
		return positive
	}
	if sdom.isAncestorEq(p.Succs[1].b, b) && len(p.Succs[1].b.Preds) == 1 {
		return negative
	}
	return unknown
}

// addIndVarRestrictions updates the factsTables ft with the facts
// learned from the induction variable indVar which drives the loop
// starting in Block b.
func (psess *PackageSession) addIndVarRestrictions(ft *factsTable, b *Block, iv indVar) {
	d := signed
	if isNonNegative(iv.min) && isNonNegative(iv.max) {
		d |= unsigned
	}

	if iv.flags&indVarMinExc == 0 {
		psess.
			addRestrictions(b, ft, d, iv.min, iv.ind, lt|eq)
	} else {
		psess.
			addRestrictions(b, ft, d, iv.min, iv.ind, lt)
	}

	if iv.flags&indVarMaxInc == 0 {
		psess.
			addRestrictions(b, ft, d, iv.ind, iv.max, lt)
	} else {
		psess.
			addRestrictions(b, ft, d, iv.ind, iv.max, lt|eq)
	}
}

// addBranchRestrictions updates the factsTables ft with the facts learned when
// branching from Block b in direction br.
func (psess *PackageSession) addBranchRestrictions(ft *factsTable, b *Block, br branch) {
	c := b.Control
	switch br {
	case negative:
		psess.
			addRestrictions(b, ft, boolean, nil, c, eq)
	case positive:
		psess.
			addRestrictions(b, ft, boolean, nil, c, lt|gt)
	default:
		panic("unknown branch")
	}
	if tr, has := psess.domainRelationTable[b.Control.Op]; has {

		d := tr.d
		if d == signed && ft.isNonNegative(c.Args[0]) && ft.isNonNegative(c.Args[1]) {
			d |= unsigned
		}
		switch br {
		case negative:
			switch b.Control.Op {
			case OpIsInBounds, OpIsSliceInBounds:

				d = unsigned
				if ft.isNonNegative(c.Args[0]) {
					d |= signed
				}
			}
			psess.
				addRestrictions(b, ft, d, c.Args[0], c.Args[1], tr.r^(lt|gt|eq))
		case positive:
			psess.
				addRestrictions(b, ft, d, c.Args[0], c.Args[1], tr.r)
		}
	}
}

// addRestrictions updates restrictions from the immediate
// dominating block (p) using r.
func (psess *PackageSession) addRestrictions(parent *Block, ft *factsTable, t domain, v, w *Value, r relation) {
	if t == 0 {

		return
	}
	for i := domain(1); i <= t; i <<= 1 {
		if t&i == 0 {
			continue
		}
		ft.update(psess, parent, v, w, i, r)
	}
}

// addLocalInductiveFacts adds inductive facts when visiting b, where
// b is a join point in a loop. In contrast with findIndVar, this
// depends on facts established for b, which is why it happens when
// visiting b. addLocalInductiveFacts specifically targets the pattern
// created by OFORUNTIL, which isn't detected by findIndVar.
//
// TODO: It would be nice to combine this with findIndVar.
func (psess *PackageSession) addLocalInductiveFacts(ft *factsTable, b *Block) {

	for _, i1 := range b.Values {
		if i1.Op != OpPhi {
			continue
		}

		min, i2 := i1.Args[0], i1.Args[1]
		if i1q, delta := isConstDelta(i2); i1q != i1 || delta != 1 {
			continue
		}

		uniquePred := func(b *Block) *Block {
			if len(b.Preds) == 1 {
				return b.Preds[0].b
			}
			return nil
		}
		pred, child := b.Preds[1].b, b
		for ; pred != nil; pred = uniquePred(pred) {
			if pred.Kind != BlockIf {
				continue
			}

			br := unknown
			if pred.Succs[0].b == child {
				br = positive
			}
			if pred.Succs[1].b == child {
				if br != unknown {
					continue
				}
				br = negative
			}

			tr, has := psess.domainRelationTable[pred.Control.Op]
			if !has {
				continue
			}
			r := tr.r
			if br == negative {

				r = (lt | eq | gt) ^ r
			}

			// Check for i2 < max or max > i2.
			var max *Value
			if r == lt && pred.Control.Args[0] == i2 {
				max = pred.Control.Args[1]
			} else if r == gt && pred.Control.Args[1] == i2 {
				max = pred.Control.Args[0]
			} else {
				continue
			}

			ft.checkpoint(psess)
			ft.update(psess, b, min, max, tr.d, gt|eq)
			proved := ft.unsat
			ft.restore(psess)

			if proved {

				if b.Func.pass.debug > 0 {
					printIndVar(b, i1, min, max, 1, 0)
				}
				ft.update(psess, b, min, i1, tr.d, lt|eq)
				ft.update(psess, b, i1, max, tr.d, lt)
			}
		}
	}
}

// simplifyBlock simplifies some constant values in b and evaluates
// branches to non-uniquely dominated successors of b.
func (psess *PackageSession) simplifyBlock(sdom SparseTree, ft *factsTable, b *Block) {
	for _, v := range b.Values {
		switch v.Op {
		case OpSlicemask:

			x, delta := isConstDelta(v.Args[0])
			if x == nil {
				continue
			}

			lim, ok := ft.limits[x.ID]
			if !ok {
				continue
			}
			if lim.umin > uint64(-delta) {
				if v.Args[0].Op == OpAdd64 {
					v.reset(OpConst64)
				} else {
					v.reset(OpConst32)
				}
				if b.Func.pass.debug > 0 {
					b.Func.Warnl(v.Pos, "Proved slicemask not needed")
				}
				v.AuxInt = -1
			}
		case OpCtz8, OpCtz16, OpCtz32, OpCtz64:

			x := v.Args[0]
			lim, ok := ft.limits[x.ID]
			if !ok {
				continue
			}
			if lim.umin > 0 || lim.min > 0 || lim.max < 0 {
				if b.Func.pass.debug > 0 {
					b.Func.Warnl(v.Pos, "Proved %v non-zero", v.Op)
				}
				v.Op = psess.ctzNonZeroOp[v.Op]
			}

		case OpLsh8x8, OpLsh8x16, OpLsh8x32, OpLsh8x64,
			OpLsh16x8, OpLsh16x16, OpLsh16x32, OpLsh16x64,
			OpLsh32x8, OpLsh32x16, OpLsh32x32, OpLsh32x64,
			OpLsh64x8, OpLsh64x16, OpLsh64x32, OpLsh64x64,
			OpRsh8x8, OpRsh8x16, OpRsh8x32, OpRsh8x64,
			OpRsh16x8, OpRsh16x16, OpRsh16x32, OpRsh16x64,
			OpRsh32x8, OpRsh32x16, OpRsh32x32, OpRsh32x64,
			OpRsh64x8, OpRsh64x16, OpRsh64x32, OpRsh64x64,
			OpRsh8Ux8, OpRsh8Ux16, OpRsh8Ux32, OpRsh8Ux64,
			OpRsh16Ux8, OpRsh16Ux16, OpRsh16Ux32, OpRsh16Ux64,
			OpRsh32Ux8, OpRsh32Ux16, OpRsh32Ux32, OpRsh32Ux64,
			OpRsh64Ux8, OpRsh64Ux16, OpRsh64Ux32, OpRsh64Ux64:

			by := v.Args[1]
			lim, ok := ft.limits[by.ID]
			if !ok {
				continue
			}
			bits := 8 * v.Args[0].Type.Size(psess.types)
			if lim.umax < uint64(bits) || (lim.max < bits && ft.isNonNegative(by)) {
				v.AuxInt = 1
				if b.Func.pass.debug > 0 {
					b.Func.Warnl(v.Pos, "Proved %v bounded", v.Op)
				}
			}
		}
	}

	if b.Kind != BlockIf {
		return
	}

	parent := b
	for i, branch := range [...]branch{positive, negative} {
		child := parent.Succs[i].b
		if getBranch(sdom, parent, child) != unknown {

			continue
		}

		ft.checkpoint(psess)
		psess.
			addBranchRestrictions(ft, parent, branch)
		unsat := ft.unsat
		ft.restore(psess)
		if unsat {

			removeBranch(parent, branch)

			break
		}
	}
}

func removeBranch(b *Block, branch branch) {
	if b.Func.pass.debug > 0 {
		verb := "Proved"
		if branch == positive {
			verb = "Disproved"
		}
		c := b.Control
		if b.Func.pass.debug > 1 {
			b.Func.Warnl(b.Pos, "%s %s (%s)", verb, c.Op, c)
		} else {
			b.Func.Warnl(b.Pos, "%s %s", verb, c.Op)
		}
	}
	b.Kind = BlockFirst
	b.SetControl(nil)
	if branch == positive {
		b.swapSuccessors()
	}
}

// isNonNegative reports whether v is known to be greater or equal to zero.
func isNonNegative(v *Value) bool {
	switch v.Op {
	case OpConst64:
		return v.AuxInt >= 0

	case OpConst32:
		return int32(v.AuxInt) >= 0

	case OpStringLen, OpSliceLen, OpSliceCap,
		OpZeroExt8to64, OpZeroExt16to64, OpZeroExt32to64:
		return true

	case OpRsh64Ux64:
		by := v.Args[1]
		return by.Op == OpConst64 && by.AuxInt > 0

	case OpRsh64x64:
		return isNonNegative(v.Args[0])
	}
	return false
}

// isConstDelta returns non-nil if v is equivalent to w+delta (signed).
func isConstDelta(v *Value) (w *Value, delta int64) {
	cop := OpConst64
	switch v.Op {
	case OpAdd32, OpSub32:
		cop = OpConst32
	}
	switch v.Op {
	case OpAdd64, OpAdd32:
		if v.Args[0].Op == cop {
			return v.Args[1], v.Args[0].AuxInt
		}
		if v.Args[1].Op == cop {
			return v.Args[0], v.Args[1].AuxInt
		}
	case OpSub64, OpSub32:
		if v.Args[1].Op == cop {
			aux := v.Args[1].AuxInt
			if aux != -aux {
				return v.Args[0], -aux
			}
		}
	}
	return nil, 0
}

package gc

import (
	"container/heap"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/src"
)

const smallBlocks = 500

const debugPhi = false

// insertPhis finds all the places in the function where a phi is
// necessary and inserts them.
// Uses FwdRef ops to find all uses of variables, and s.defvars to find
// all definitions.
// Phi values are inserted, and all FwdRefs are changed to a Copy
// of the appropriate phi or definition.
// TODO: make this part of cmd/compile/internal/ssa somehow?
func (s *state) insertPhis(psess *PackageSession) {
	if len(s.f.Blocks) <= smallBlocks {
		sps := simplePhiState{s: s, f: s.f, defvars: s.defvars}
		sps.insertPhis()
		return
	}
	ps := phiState{s: s, f: s.f, defvars: s.defvars}
	ps.insertPhis(psess)
}

type phiState struct {
	s       *state                 // SSA state
	f       *ssa.Func              // function to work on
	defvars []map[*Node]*ssa.Value // defined variables at end of each block

	varnum map[*Node]int32 // variable numbering

	// properties of the dominator tree
	idom  []*ssa.Block // dominator parents
	tree  []domBlock   // dominator child+sibling
	level []int32      // level in dominator tree (0 = root or unreachable, 1 = children of root, ...)

	// scratch locations
	priq   blockHeap    // priority queue of blocks, higher level (toward leaves) = higher priority
	q      []*ssa.Block // inner loop queue
	queued *sparseSet   // has been put in q
	hasPhi *sparseSet   // has a phi
	hasDef *sparseSet   // has a write of the variable we're processing

	// miscellaneous
	placeholder *ssa.Value // dummy value to use as a "not set yet" placeholder.
}

func (s *phiState) insertPhis(psess *PackageSession) {
	if debugPhi {
		fmt.Println(s.f.String(psess.ssa))
	}

	s.varnum = map[*Node]int32{}
	var vars []*Node
	var vartypes []*types.Type
	for _, b := range s.f.Blocks {
		for _, v := range b.Values {
			if v.Op != ssa.OpFwdRef {
				continue
			}
			var_ := v.Aux.(*Node)

			if len(b.Preds) == 1 {
				c := b.Preds[0].Block()
				if w := s.defvars[c.ID][var_]; w != nil {
					v.Op = ssa.OpCopy
					v.Aux = nil
					v.AddArg(w)
					continue
				}
			}

			if _, ok := s.varnum[var_]; ok {
				continue
			}
			s.varnum[var_] = int32(len(vartypes))
			if debugPhi {
				fmt.Printf("var%d = %v\n", len(vartypes), var_)
			}
			vars = append(vars, var_)
			vartypes = append(vartypes, v.Type)
		}
	}

	if len(vartypes) == 0 {
		return
	}

	defs := make([][]*ssa.Block, len(vartypes))
	for _, b := range s.f.Blocks {
		for var_ := range s.defvars[b.ID] {
			if n, ok := s.varnum[var_]; ok {
				defs[n] = append(defs[n], b)
			}
		}
	}

	s.idom = s.f.Idom()
	s.tree = make([]domBlock, s.f.NumBlocks())
	for _, b := range s.f.Blocks {
		p := s.idom[b.ID]
		if p != nil {
			s.tree[b.ID].sibling = s.tree[p.ID].firstChild
			s.tree[p.ID].firstChild = b
		}
	}

	s.level = make([]int32, s.f.NumBlocks())
	b := s.f.Entry
levels:
	for {
		if p := s.idom[b.ID]; p != nil {
			s.level[b.ID] = s.level[p.ID] + 1
			if debugPhi {
				fmt.Printf("level %s = %d\n", b, s.level[b.ID])
			}
		}
		if c := s.tree[b.ID].firstChild; c != nil {
			b = c
			continue
		}
		for {
			if c := s.tree[b.ID].sibling; c != nil {
				b = c
				continue levels
			}
			b = s.idom[b.ID]
			if b == nil {
				break levels
			}
		}
	}

	s.priq.level = s.level
	s.q = make([]*ssa.Block, 0, s.f.NumBlocks())
	s.queued = newSparseSet(s.f.NumBlocks())
	s.hasPhi = newSparseSet(s.f.NumBlocks())
	s.hasDef = newSparseSet(s.f.NumBlocks())
	s.placeholder = s.s.entryNewValue0(psess, ssa.OpUnknown, psess.types.TypeInvalid)

	for n := range vartypes {
		s.insertVarPhis(n, vars[n], defs[n], vartypes[n])
	}

	s.resolveFwdRefs()

	for _, b := range s.f.Blocks {
		for _, v := range b.Values {
			if v.Op == ssa.OpPhi {
				v.AuxInt = 0
			}
		}
	}
}

func (s *phiState) insertVarPhis(n int, var_ *Node, defs []*ssa.Block, typ *types.Type) {
	priq := &s.priq
	q := s.q
	queued := s.queued
	queued.clear()
	hasPhi := s.hasPhi
	hasPhi.clear()
	hasDef := s.hasDef
	hasDef.clear()

	for _, b := range defs {
		priq.a = append(priq.a, b)
		hasDef.add(b.ID)
		if debugPhi {
			fmt.Printf("def of var%d in %s\n", n, b)
		}
	}
	heap.Init(priq)

	for len(priq.a) > 0 {
		currentRoot := heap.Pop(priq).(*ssa.Block)
		if debugPhi {
			fmt.Printf("currentRoot %s\n", currentRoot)
		}

		if queued.contains(currentRoot.ID) {
			s.s.Fatalf("root already in queue")
		}
		q = append(q, currentRoot)
		queued.add(currentRoot.ID)
		for len(q) > 0 {
			b := q[len(q)-1]
			q = q[:len(q)-1]
			if debugPhi {
				fmt.Printf("  processing %s\n", b)
			}

			currentRootLevel := s.level[currentRoot.ID]
			for _, e := range b.Succs {
				c := e.Block()

				if s.level[c.ID] > currentRootLevel {

					continue
				}
				if hasPhi.contains(c.ID) {
					continue
				}

				hasPhi.add(c.ID)
				v := c.NewValue0I(currentRoot.Pos, ssa.OpPhi, typ, int64(n))

				s.s.addNamedValue(var_, v)
				for range c.Preds {
					v.AddArg(s.placeholder)
				}
				if debugPhi {
					fmt.Printf("new phi for var%d in %s: %s\n", n, c, v)
				}
				if !hasDef.contains(c.ID) {

					heap.Push(priq, c)
					hasDef.add(c.ID)
				}
			}

			for c := s.tree[b.ID].firstChild; c != nil; c = s.tree[c.ID].sibling {
				if !queued.contains(c.ID) {
					q = append(q, c)
					queued.add(c.ID)
				}
			}
		}
	}
}

// resolveFwdRefs links all FwdRef uses up to their nearest dominating definition.
func (s *phiState) resolveFwdRefs() {

	values := make([]*ssa.Value, len(s.varnum))
	for i := range values {
		values[i] = s.placeholder
	}

	// Stack of work to do.
	type stackEntry struct {
		b *ssa.Block // block to explore

		// variable/value pair to reinstate on exit
		n int32 // variable ID
		v *ssa.Value
	}
	var stk []stackEntry

	stk = append(stk, stackEntry{b: s.f.Entry})
	for len(stk) > 0 {
		work := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		b := work.b
		if b == nil {

			values[work.n] = work.v
			continue
		}

		for _, v := range b.Values {
			if v.Op != ssa.OpPhi {
				continue
			}
			n := int32(v.AuxInt)

			stk = append(stk, stackEntry{n: n, v: values[n]})

			values[n] = v
		}

		for _, v := range b.Values {
			if v.Op != ssa.OpFwdRef {
				continue
			}
			n := s.varnum[v.Aux.(*Node)]
			v.Op = ssa.OpCopy
			v.Aux = nil
			v.AddArg(values[n])
		}

		for var_, v := range s.defvars[b.ID] {
			n, ok := s.varnum[var_]
			if !ok {

				continue
			}

			stk = append(stk, stackEntry{n: n, v: values[n]})

			values[n] = v
		}

		for _, e := range b.Succs {
			c, i := e.Block(), e.Index()
			for j := len(c.Values) - 1; j >= 0; j-- {
				v := c.Values[j]
				if v.Op != ssa.OpPhi {
					break
				}

				if w := values[v.AuxInt]; w.Op != ssa.OpUnknown {
					v.SetArg(i, w)
				}
			}
		}

		for c := s.tree[b.ID].firstChild; c != nil; c = s.tree[c.ID].sibling {
			stk = append(stk, stackEntry{b: c})
		}
	}
}

// domBlock contains extra per-block information to record the dominator tree.
type domBlock struct {
	firstChild *ssa.Block // first child of block in dominator tree
	sibling    *ssa.Block // next child of parent in dominator tree
}

// A block heap is used as a priority queue to implement the PiggyBank
// from Sreedhar and Gao.  That paper uses an array which is better
// asymptotically but worse in the common case when the PiggyBank
// holds a sparse set of blocks.
type blockHeap struct {
	a     []*ssa.Block // block IDs in heap
	level []int32      // depth in dominator tree (static, used for determining priority)
}

func (h *blockHeap) Len() int      { return len(h.a) }
func (h *blockHeap) Swap(i, j int) { a := h.a; a[i], a[j] = a[j], a[i] }

func (h *blockHeap) Push(x interface{}) {
	v := x.(*ssa.Block)
	h.a = append(h.a, v)
}
func (h *blockHeap) Pop() interface{} {
	old := h.a
	n := len(old)
	x := old[n-1]
	h.a = old[:n-1]
	return x
}
func (h *blockHeap) Less(i, j int) bool {
	return h.level[h.a[i].ID] > h.level[h.a[j].ID]
}

// copy of ../ssa/sparseset.go
// TODO: move this file to ../ssa, then use sparseSet there.
type sparseSet struct {
	dense  []ssa.ID
	sparse []int32
}

// newSparseSet returns a sparseSet that can represent
// integers between 0 and n-1
func newSparseSet(n int) *sparseSet {
	return &sparseSet{dense: nil, sparse: make([]int32, n)}
}

func (s *sparseSet) contains(x ssa.ID) bool {
	i := s.sparse[x]
	return i < int32(len(s.dense)) && s.dense[i] == x
}

func (s *sparseSet) add(x ssa.ID) {
	i := s.sparse[x]
	if i < int32(len(s.dense)) && s.dense[i] == x {
		return
	}
	s.dense = append(s.dense, x)
	s.sparse[x] = int32(len(s.dense)) - 1
}

func (s *sparseSet) clear() {
	s.dense = s.dense[:0]
}

// Variant to use for small functions.
type simplePhiState struct {
	s         *state                 // SSA state
	f         *ssa.Func              // function to work on
	fwdrefs   []*ssa.Value           // list of FwdRefs to be processed
	defvars   []map[*Node]*ssa.Value // defined variables at end of each block
	reachable []bool                 // which blocks are reachable
}

func (s *simplePhiState) insertPhis() {
	s.reachable = ssa.ReachableBlocks(s.f)

	for _, b := range s.f.Blocks {
		for _, v := range b.Values {
			if v.Op != ssa.OpFwdRef {
				continue
			}
			s.fwdrefs = append(s.fwdrefs, v)
			var_ := v.Aux.(*Node)
			if _, ok := s.defvars[b.ID][var_]; !ok {
				s.defvars[b.ID][var_] = v
			}
		}
	}

	var args []*ssa.Value

loop:
	for len(s.fwdrefs) > 0 {
		v := s.fwdrefs[len(s.fwdrefs)-1]
		s.fwdrefs = s.fwdrefs[:len(s.fwdrefs)-1]
		b := v.Block
		var_ := v.Aux.(*Node)
		if b == s.f.Entry {

			s.s.Fatalf("Value live at entry. It shouldn't be. func %s, node %v, value %v", s.f.Name, var_, v)
		}
		if !s.reachable[b.ID] {

			v.Op = ssa.OpUnknown
			v.Aux = nil
			continue
		}

		args = args[:0]
		for _, e := range b.Preds {
			args = append(args, s.lookupVarOutgoing(e.Block(), v.Type, var_, v.Pos))
		}

		// Decide if we need a phi or not. We need a phi if there
		// are two different args (which are both not v).
		var w *ssa.Value
		for _, a := range args {
			if a == v {
				continue
			}
			if a == w {
				continue
			}
			if w != nil {

				v.Op = ssa.OpPhi
				v.AddArgs(args...)
				v.Aux = nil
				continue loop
			}
			w = a
		}
		if w == nil {
			s.s.Fatalf("no witness for reachable phi %s", v)
		}

		v.Op = ssa.OpCopy
		v.Aux = nil
		v.AddArg(w)
	}
}

// lookupVarOutgoing finds the variable's value at the end of block b.
func (s *simplePhiState) lookupVarOutgoing(b *ssa.Block, t *types.Type, var_ *Node, line src.XPos) *ssa.Value {
	for {
		if v := s.defvars[b.ID][var_]; v != nil {
			return v
		}

		if len(b.Preds) != 1 {
			break
		}
		b = b.Preds[0].Block()
		if !s.reachable[b.ID] {

			break
		}
	}

	v := b.NewValue0A(line, ssa.OpFwdRef, t, var_)
	s.defvars[b.ID][var_] = v
	s.s.addNamedValue(var_, v)
	s.fwdrefs = append(s.fwdrefs, v)
	return v
}

package ssa

import "container/heap"

const (
	ScorePhi = iota // towards top of block
	ScoreNilCheck
	ScoreReadTuple
	ScoreVarDef
	ScoreMemory
	ScoreDefault
	ScoreFlags
	ScoreControl // towards bottom of block
)

type ValHeap struct {
	a     []*Value
	score []int8
}

func (h ValHeap) Len() int      { return len(h.a) }
func (h ValHeap) Swap(i, j int) { a := h.a; a[i], a[j] = a[j], a[i] }

func (h *ValHeap) Push(x interface{}) {

	v := x.(*Value)
	h.a = append(h.a, v)
}
func (h *ValHeap) Pop() interface{} {
	old := h.a
	n := len(old)
	x := old[n-1]
	h.a = old[0 : n-1]
	return x
}
func (h ValHeap) Less(i, j int) bool {
	x := h.a[i]
	y := h.a[j]
	sx := h.score[x.ID]
	sy := h.score[y.ID]
	if c := sx - sy; c != 0 {
		return c > 0
	}
	if x.Pos != y.Pos {
		return x.Pos.After(y.Pos)
	}
	if x.Op != OpPhi {
		if c := len(x.Args) - len(y.Args); c != 0 {
			return c < 0
		}
	}
	return x.ID > y.ID
}

// Schedule the Values in each Block. After this phase returns, the
// order of b.Values matters and is the order in which those values
// will appear in the assembly output. For now it generates a
// reasonable valid schedule using a priority queue. TODO(khr):
// schedule smarter.
func (psess *PackageSession) schedule(f *Func) {

	uses := make([]int32, f.NumValues())

	priq := new(ValHeap)

	score := make([]int8, f.NumValues())

	order := make([]*Value, 0, 64)

	nextMem := make([]*Value, f.NumValues())

	additionalArgs := make([][]*Value, f.NumValues())

	for _, b := range f.Blocks {

		for _, v := range b.Values {
			switch {
			case v.Op == OpAMD64LoweredGetClosurePtr || v.Op == OpPPC64LoweredGetClosurePtr ||
				v.Op == OpARMLoweredGetClosurePtr || v.Op == OpARM64LoweredGetClosurePtr ||
				v.Op == Op386LoweredGetClosurePtr || v.Op == OpMIPS64LoweredGetClosurePtr ||
				v.Op == OpS390XLoweredGetClosurePtr || v.Op == OpMIPSLoweredGetClosurePtr ||
				v.Op == OpWasmLoweredGetClosurePtr:

				if b != f.Entry {
					f.Fatalf("LoweredGetClosurePtr appeared outside of entry block, b=%s", b.String())
				}
				score[v.ID] = ScorePhi
			case v.Op == OpAMD64LoweredNilCheck || v.Op == OpPPC64LoweredNilCheck ||
				v.Op == OpARMLoweredNilCheck || v.Op == OpARM64LoweredNilCheck ||
				v.Op == Op386LoweredNilCheck || v.Op == OpMIPS64LoweredNilCheck ||
				v.Op == OpS390XLoweredNilCheck || v.Op == OpMIPSLoweredNilCheck ||
				v.Op == OpWasmLoweredNilCheck:

				score[v.ID] = ScoreNilCheck
			case v.Op == OpPhi:

				score[v.ID] = ScorePhi
			case v.Op == OpVarDef:

				score[v.ID] = ScoreVarDef
			case v.Type.IsMemory(psess.types):

				score[v.ID] = ScoreMemory
			case v.Op == OpSelect0 || v.Op == OpSelect1:

				score[v.ID] = ScoreReadTuple
			case v.Type.IsFlags(psess.types) || v.Type.IsTuple():

				score[v.ID] = ScoreFlags
			default:
				score[v.ID] = ScoreDefault
			}
		}
	}

	for _, b := range f.Blocks {

		for _, v := range b.Values {
			if v.Op != OpPhi && v.Type.IsMemory(psess.types) {
				for _, w := range v.Args {
					if w.Type.IsMemory(psess.types) {
						nextMem[w.ID] = v
					}
				}
			}
		}

		for _, v := range b.Values {
			if v.Op == OpPhi {

				continue
			}
			for _, w := range v.Args {
				if w.Block == b {
					uses[w.ID]++
				}

				if !v.Type.IsMemory(psess.types) && w.Type.IsMemory(psess.types) {

					s := nextMem[w.ID]
					if s == nil || s.Block != b {
						continue
					}
					additionalArgs[s.ID] = append(additionalArgs[s.ID], v)
					uses[v.ID]++
				}
			}
		}

		if b.Control != nil && b.Control.Op != OpPhi {

			score[b.Control.ID] = ScoreControl

			for _, v := range b.Values {
				if v.Op != OpPhi {
					for _, a := range v.Args {
						if a == b.Control {
							score[v.ID] = ScoreControl
						}
					}
				}
			}
		}

		priq.score = score
		priq.a = priq.a[:0]

		for _, v := range b.Values {
			if uses[v.ID] == 0 {
				heap.Push(priq, v)
			}
		}

		order = order[:0]
		tuples := make(map[ID][]*Value)
		for {

			if priq.Len() == 0 {
				break
			}

			v := heap.Pop(priq).(*Value)

			switch {
			case v.Op == OpSelect0:
				if tuples[v.Args[0].ID] == nil {
					tuples[v.Args[0].ID] = make([]*Value, 2)
				}
				tuples[v.Args[0].ID][0] = v
			case v.Op == OpSelect1:
				if tuples[v.Args[0].ID] == nil {
					tuples[v.Args[0].ID] = make([]*Value, 2)
				}
				tuples[v.Args[0].ID][1] = v
			case v.Type.IsTuple() && tuples[v.ID] != nil:
				if tuples[v.ID][1] != nil {
					order = append(order, tuples[v.ID][1])
				}
				if tuples[v.ID][0] != nil {
					order = append(order, tuples[v.ID][0])
				}
				delete(tuples, v.ID)
				fallthrough
			default:
				order = append(order, v)
			}

			for _, w := range v.Args {
				if w.Block != b {
					continue
				}
				uses[w.ID]--
				if uses[w.ID] == 0 {

					heap.Push(priq, w)
				}
			}
			for _, w := range additionalArgs[v.ID] {
				uses[w.ID]--
				if uses[w.ID] == 0 {

					heap.Push(priq, w)
				}
			}
		}
		if len(order) != len(b.Values) {
			f.Fatalf("schedule does not include all values in block %s", b)
		}
		for i := 0; i < len(b.Values); i++ {
			b.Values[i] = order[len(b.Values)-1-i]
		}
	}

	f.scheduled = true
}

// storeOrder orders values with respect to stores. That is,
// if v transitively depends on store s, v is ordered after s,
// otherwise v is ordered before s.
// Specifically, values are ordered like
//   store1
//   NilCheck that depends on store1
//   other values that depends on store1
//   store2
//   NilCheck that depends on store2
//   other values that depends on store2
//   ...
// The order of non-store and non-NilCheck values are undefined
// (not necessarily dependency order). This should be cheaper
// than a full scheduling as done above.
// Note that simple dependency order won't work: there is no
// dependency between NilChecks and values like IsNonNil.
// Auxiliary data structures are passed in as arguments, so
// that they can be allocated in the caller and be reused.
// This function takes care of reset them.
func (psess *PackageSession) storeOrder(values []*Value, sset *sparseSet, storeNumber []int32) []*Value {
	if len(values) == 0 {
		return values
	}

	f := values[0].Block.Func

	stores := make([]*Value, 0, 64)
	hasNilCheck := false
	sset.clear()
	for _, v := range values {
		if v.Type.IsMemory(psess.types) {
			stores = append(stores, v)
			if v.Op == OpInitMem || v.Op == OpPhi {
				continue
			}
			sset.add(v.MemoryArg(psess).ID)
		}
		if v.Op == OpNilCheck {
			hasNilCheck = true
		}
	}
	if len(stores) == 0 || !hasNilCheck && f.pass.name == "nilcheckelim" {

		return values
	}

	// find last store, which is the one that is not used by other stores
	var last *Value
	for _, v := range stores {
		if !sset.contains(v.ID) {
			if last != nil {
				f.Fatalf("two stores live simultaneously: %v and %v", v, last)
			}
			last = v
		}
	}

	count := make([]int32, 3*(len(stores)+1))
	sset.clear()
	for n, w := len(stores), last; n > 0; n-- {
		storeNumber[w.ID] = int32(3 * n)
		count[3*n]++
		sset.add(w.ID)
		if w.Op == OpInitMem || w.Op == OpPhi {
			if n != 1 {
				f.Fatalf("store order is wrong: there are stores before %v", w)
			}
			break
		}
		w = w.MemoryArg(psess)
	}
	var stack []*Value
	for _, v := range values {
		if sset.contains(v.ID) {

			continue
		}
		stack = append(stack, v)
		sset.add(v.ID)

		for len(stack) > 0 {
			w := stack[len(stack)-1]
			if storeNumber[w.ID] != 0 {
				stack = stack[:len(stack)-1]
				continue
			}
			if w.Op == OpPhi {

				storeNumber[w.ID] = 2
				count[2]++
				stack = stack[:len(stack)-1]
				continue
			}

			max := int32(0)
			argsdone := true
			for _, a := range w.Args {
				if a.Block != w.Block {
					continue
				}
				if !sset.contains(a.ID) {
					stack = append(stack, a)
					sset.add(a.ID)
					argsdone = false
					break
				}
				if storeNumber[a.ID]/3 > max {
					max = storeNumber[a.ID] / 3
				}
			}
			if !argsdone {
				continue
			}

			n := 3*max + 2
			if w.Op == OpNilCheck {
				n = 3*max + 1
			}
			storeNumber[w.ID] = n
			count[n]++
			stack = stack[:len(stack)-1]
		}
	}

	for i := range count {
		if i == 0 {
			continue
		}
		count[i] += count[i-1]
	}
	if count[len(count)-1] != int32(len(values)) {
		f.Fatalf("storeOrder: value is missing, total count = %d, values = %v", count[len(count)-1], values)
	}

	order := make([]*Value, len(values))
	for _, v := range values {
		s := storeNumber[v.ID]
		order[count[s-1]] = v
		count[s-1]++
	}

	return order
}

package ssa

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/src"
	"sort"
)

// cse does common-subexpression elimination on the Function.
// Values are just relinked, nothing is deleted. A subsequent deadcode
// pass is required to actually remove duplicate expressions.
func (psess *PackageSession) cse(f *Func) {

	a := make([]*Value, 0, f.NumValues())
	if f.auxmap == nil {
		f.auxmap = auxmap{}
	}
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if v.Type.IsMemory(psess.types) {
				continue
			}
			if f.auxmap[v.Aux] == 0 {
				f.auxmap[v.Aux] = int32(len(f.auxmap)) + 1
			}
			a = append(a, v)
		}
	}
	partition := psess.partitionValues(a, f.auxmap)

	valueEqClass := make([]ID, f.NumValues())
	for _, b := range f.Blocks {
		for _, v := range b.Values {

			valueEqClass[v.ID] = -v.ID
		}
	}
	var pNum ID = 1
	for _, e := range partition {
		if f.pass.debug > 1 && len(e) > 500 {
			fmt.Printf("CSE.large partition (%d): ", len(e))
			for j := 0; j < 3; j++ {
				fmt.Printf("%s ", e[j].LongString(psess))
			}
			fmt.Println()
		}

		for _, v := range e {
			valueEqClass[v.ID] = pNum
		}
		if f.pass.debug > 2 && len(e) > 1 {
			fmt.Printf("CSE.partition #%d:", pNum)
			for _, v := range e {
				fmt.Printf(" %s", v.String())
			}
			fmt.Printf("\n")
		}
		pNum++
	}

	// Split equivalence classes at points where they have
	// non-equivalent arguments.  Repeat until we can't find any
	// more splits.
	var splitPoints []int
	byArgClass := new(partitionByArgClass)
	for {
		changed := false

		for i := 0; i < len(partition); i++ {
			e := partition[i]

			if psess.opcodeTable[e[0].Op].commutative {

				for _, v := range e {
					if valueEqClass[v.Args[0].ID] > valueEqClass[v.Args[1].ID] {
						v.Args[0], v.Args[1] = v.Args[1], v.Args[0]
					}
				}
			}

			byArgClass.a = e
			byArgClass.eqClass = valueEqClass
			sort.Sort(byArgClass)

			splitPoints = append(splitPoints[:0], 0)
			for j := 1; j < len(e); j++ {
				v, w := e[j-1], e[j]

				eqArgs := true
				for k, a := range v.Args {
					b := w.Args[k]
					if valueEqClass[a.ID] != valueEqClass[b.ID] {
						eqArgs = false
						break
					}
				}
				if !eqArgs {
					splitPoints = append(splitPoints, j)
				}
			}
			if len(splitPoints) == 1 {
				continue
			}

			partition[i] = partition[len(partition)-1]
			partition = partition[:len(partition)-1]
			i--

			splitPoints = append(splitPoints, len(e))
			for j := 0; j < len(splitPoints)-1; j++ {
				f := e[splitPoints[j]:splitPoints[j+1]]
				if len(f) == 1 {

					valueEqClass[f[0].ID] = -f[0].ID
					continue
				}
				for _, v := range f {
					valueEqClass[v.ID] = pNum
				}
				pNum++
				partition = append(partition, f)
			}
			changed = true
		}

		if !changed {
			break
		}
	}

	sdom := f.sdom()

	rewrite := make([]*Value, f.NumValues())
	byDom := new(partitionByDom)
	for _, e := range partition {
		byDom.a = e
		byDom.sdom = sdom
		sort.Sort(byDom)
		for i := 0; i < len(e)-1; i++ {

			v := e[i]
			if v == nil {
				continue
			}

			e[i] = nil

			for j := i + 1; j < len(e); j++ {
				w := e[j]
				if w == nil {
					continue
				}
				if sdom.isAncestorEq(v.Block, w.Block) {
					rewrite[w.ID] = v
					e[j] = nil
				} else {

					break
				}
			}
		}
	}

	copiedSelects := make(map[ID][]*Value)
	for _, b := range f.Blocks {
	out:
		for _, v := range b.Values {

			if int(v.ID) >= len(rewrite) || rewrite[v.ID] != nil {
				continue
			}
			if v.Op != OpSelect0 && v.Op != OpSelect1 {
				continue
			}
			if !v.Args[0].Type.IsTuple() {
				f.Fatalf("arg of tuple selector %s is not a tuple: %s", v.String(), v.Args[0].LongString(psess))
			}
			t := rewrite[v.Args[0].ID]
			if t != nil && t.Block != b {

				for _, c := range copiedSelects[t.ID] {
					if v.Op == c.Op {

						rewrite[v.ID] = c
						continue out
					}
				}
				c := v.copyInto(psess, t.Block)
				rewrite[v.ID] = c
				copiedSelects[t.ID] = append(copiedSelects[t.ID], c)
			}
		}
	}

	rewrites := int64(0)

	for _, b := range f.Blocks {
		for _, v := range b.Values {
			for i, w := range v.Args {
				if x := rewrite[w.ID]; x != nil {
					if w.Pos.IsStmt() == src.PosIsStmt {

						if w.Block == v.Block && w.Pos.Line() == v.Pos.Line() {
							v.Pos = v.Pos.WithIsStmt()
							w.Pos = w.Pos.WithNotStmt()
						}
					}
					v.SetArg(i, x)
					rewrites++
				}
			}
		}
		if v := b.Control; v != nil {
			if x := rewrite[v.ID]; x != nil {
				if v.Op == OpNilCheck {

					continue
				}
				b.SetControl(x)
			}
		}
	}
	if f.pass.stats > 0 {
		f.LogStat("CSE REWRITES", rewrites)
	}
}

// An eqclass approximates an equivalence class. During the
// algorithm it may represent the union of several of the
// final equivalence classes.
type eqclass []*Value

// partitionValues partitions the values into equivalence classes
// based on having all the following features match:
//  - opcode
//  - type
//  - auxint
//  - aux
//  - nargs
//  - block # if a phi op
//  - first two arg's opcodes and auxint
//  - NOT first two arg's aux; that can break CSE.
// partitionValues returns a list of equivalence classes, each
// being a sorted by ID list of *Values. The eqclass slices are
// backed by the same storage as the input slice.
// Equivalence classes of size 1 are ignored.
func (psess *PackageSession) partitionValues(a []*Value, auxIDs auxmap) []eqclass {
	sort.Sort(sortvalues{a, auxIDs})

	var partition []eqclass
	for len(a) > 0 {
		v := a[0]
		j := 1
		for ; j < len(a); j++ {
			w := a[j]
			if psess.cmpVal(v, w, auxIDs) != types.CMPeq {
				break
			}
		}
		if j > 1 {
			partition = append(partition, a[:j])
		}
		a = a[j:]
	}

	return partition
}
func lt2Cmp(isLt bool) types.Cmp {
	if isLt {
		return types.CMPlt
	}
	return types.CMPgt
}

type auxmap map[interface{}]int32

func (psess *PackageSession) cmpVal(v, w *Value, auxIDs auxmap) types.Cmp {

	if v.Op != w.Op {
		return lt2Cmp(v.Op < w.Op)
	}
	if v.AuxInt != w.AuxInt {
		return lt2Cmp(v.AuxInt < w.AuxInt)
	}
	if len(v.Args) != len(w.Args) {
		return lt2Cmp(len(v.Args) < len(w.Args))
	}
	if v.Op == OpPhi && v.Block != w.Block {
		return lt2Cmp(v.Block.ID < w.Block.ID)
	}
	if v.Type.IsMemory(psess.types) {

		return lt2Cmp(v.ID < w.ID)
	}

	if v.Op != OpSelect0 && v.Op != OpSelect1 {
		if tc := v.Type.Compare(psess.types, w.Type); tc != types.CMPeq {
			return tc
		}
	}

	if v.Aux != w.Aux {
		if v.Aux == nil {
			return types.CMPlt
		}
		if w.Aux == nil {
			return types.CMPgt
		}
		return lt2Cmp(auxIDs[v.Aux] < auxIDs[w.Aux])
	}

	return types.CMPeq
}

// Sort values to make the initial partition.
type sortvalues struct {
	a      []*Value // array of values
	auxIDs auxmap   // aux -> aux ID map
}

func (sv sortvalues) Len() int      { return len(sv.a) }
func (sv sortvalues) Swap(i, j int) { sv.a[i], sv.a[j] = sv.a[j], sv.a[i] }
func (sv sortvalues) Less(psess *PackageSession, i, j int) bool {
	v := sv.a[i]
	w := sv.a[j]
	if cmp := psess.cmpVal(v, w, sv.auxIDs); cmp != types.CMPeq {
		return cmp == types.CMPlt
	}

	return v.ID < w.ID
}

type partitionByDom struct {
	a    []*Value // array of values
	sdom SparseTree
}

func (sv partitionByDom) Len() int      { return len(sv.a) }
func (sv partitionByDom) Swap(i, j int) { sv.a[i], sv.a[j] = sv.a[j], sv.a[i] }
func (sv partitionByDom) Less(i, j int) bool {
	v := sv.a[i]
	w := sv.a[j]
	return sv.sdom.domorder(v.Block) < sv.sdom.domorder(w.Block)
}

type partitionByArgClass struct {
	a       []*Value // array of values
	eqClass []ID     // equivalence class IDs of values
}

func (sv partitionByArgClass) Len() int      { return len(sv.a) }
func (sv partitionByArgClass) Swap(i, j int) { sv.a[i], sv.a[j] = sv.a[j], sv.a[i] }
func (sv partitionByArgClass) Less(i, j int) bool {
	v := sv.a[i]
	w := sv.a[j]
	for i, a := range v.Args {
		b := w.Args[i]
		if sv.eqClass[a.ID] < sv.eqClass[b.ID] {
			return true
		}
		if sv.eqClass[a.ID] > sv.eqClass[b.ID] {
			return false
		}
	}
	return false
}

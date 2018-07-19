package ssa

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
)

// dse does dead-store elimination on the Function.
// Dead stores are those which are unconditionally followed by
// another store to the same location, with no intervening load.
// This implementation only works within a basic block. TODO: use something more global.
func (psess *PackageSession) dse(f *Func) {
	var stores []*Value
	loadUse := f.newSparseSet(f.NumValues())
	defer f.retSparseSet(loadUse)
	storeUse := f.newSparseSet(f.NumValues())
	defer f.retSparseSet(storeUse)
	shadowed := f.newSparseMap(f.NumValues())
	defer f.retSparseMap(shadowed)
	for _, b := range f.Blocks {

		loadUse.clear()
		storeUse.clear()
		stores = stores[:0]
		for _, v := range b.Values {
			if v.Op == OpPhi {

				continue
			}
			if v.Type.IsMemory(psess.types) {
				stores = append(stores, v)
				for _, a := range v.Args {
					if a.Block == b && a.Type.IsMemory(psess.types) {
						storeUse.add(a.ID)
						if v.Op != OpStore && v.Op != OpZero && v.Op != OpVarDef && v.Op != OpVarKill {

							loadUse.add(a.ID)
						}
					}
				}
			} else {
				for _, a := range v.Args {
					if a.Block == b && a.Type.IsMemory(psess.types) {
						loadUse.add(a.ID)
					}
				}
			}
		}
		if len(stores) == 0 {
			continue
		}

		// find last store in the block
		var last *Value
		for _, v := range stores {
			if storeUse.contains(v.ID) {
				continue
			}
			if last != nil {
				b.Fatalf("two final stores - simultaneous live stores %s %s", last.LongString(psess), v.LongString(psess))
			}
			last = v
		}
		if last == nil {
			b.Fatalf("no last store found - cycle?")
		}

		shadowed.clear()
		v := last

	walkloop:
		if loadUse.contains(v.ID) {

			shadowed.clear()
		}
		if v.Op == OpStore || v.Op == OpZero {
			var sz int64
			if v.Op == OpStore {
				sz = v.Aux.(*types.Type).Size(psess.types)
			} else {
				sz = v.AuxInt
			}
			if shadowedSize := int64(shadowed.get(v.Args[0].ID)); shadowedSize != -1 && shadowedSize >= sz {

				if v.Op == OpStore {

					v.SetArgs1(v.Args[2])
				} else {

					typesz := v.Args[0].Type.Elem(psess.types).Size(psess.types)
					if sz != typesz {
						f.Fatalf("mismatched zero/store sizes: %d and %d [%s]",
							sz, typesz, v.LongString(psess))
					}
					v.SetArgs1(v.Args[1])
				}
				v.Aux = nil
				v.AuxInt = 0
				v.Op = OpCopy
			} else {
				if sz > 0x7fffffff {
					sz = 0x7fffffff
				}
				shadowed.set(v.Args[0].ID, int32(sz), psess.src.NoXPos)
			}
		}

		if v.Op == OpPhi {

			continue
		}
		for _, a := range v.Args {
			if a.Block == b && a.Type.IsMemory(psess.types) {
				v = a
				goto walkloop
			}
		}
	}
}

// elimDeadAutosGeneric deletes autos that are never accessed. To acheive this
// we track the operations that the address of each auto reaches and if it only
// reaches stores then we delete all the stores. The other operations will then
// be eliminated by the dead code elimination pass.
func (psess *PackageSession) elimDeadAutosGeneric(f *Func) {
	addr := make(map[*Value]GCNode)
	elim := make(map[*Value]GCNode)
	used := make(map[GCNode]bool)

	visit := func(v *Value) (changed bool) {
		args := v.Args
		switch v.Op {
		case OpAddr:

			n, ok := v.Aux.(GCNode)
			if !ok || n.StorageClass() != ClassAuto {
				return
			}
			if addr[v] == nil {
				addr[v] = n
				changed = true
			}
			return
		case OpVarDef, OpVarKill:

			n, ok := v.Aux.(GCNode)
			if !ok || n.StorageClass() != ClassAuto {
				return
			}
			if elim[v] == nil {
				elim[v] = n
				changed = true
			}
			return
		case OpVarLive:

			n, ok := v.Aux.(GCNode)
			if !ok || n.StorageClass() != ClassAuto {
				return
			}
			if !used[n] {
				used[n] = true
				changed = true
			}
			return
		case OpStore, OpMove, OpZero:

			n, ok := addr[args[0]]
			if ok && elim[v] == nil {
				elim[v] = n
				changed = true
			}

			args = args[1:]
		}

		if v.Op.SymEffect(psess) != SymNone && v.Op != OpArg {
			panic("unhandled op with sym effect")
		}

		if v.Uses == 0 || len(args) == 0 {
			return
		}

		if v.Type.IsMemory(psess.types) || v.Type.IsFlags(psess.types) || (v.Op != OpPhi && v.MemoryArg(psess) != nil) {
			for _, a := range args {
				if n, ok := addr[a]; ok {
					if !used[n] {
						used[n] = true
						changed = true
					}
				}
			}
			return
		}

		node := GCNode(nil)
		for _, a := range args {
			if n, ok := addr[a]; ok && !used[n] {
				if node == nil {
					node = n
				} else if node != n {

					used[n] = true
					changed = true
				}
			}
		}
		if node == nil {
			return
		}
		if addr[v] == nil {

			addr[v] = node
			changed = true
			return
		}
		if addr[v] != node {

			used[node] = true
			changed = true
		}
		return
	}

	iterations := 0
	for {
		if iterations == 4 {

			return
		}
		iterations++
		changed := false
		for _, b := range f.Blocks {
			for _, v := range b.Values {
				changed = visit(v) || changed
			}
		}
		if !changed {
			break
		}
	}

	for v, n := range elim {
		if used[n] {
			continue
		}

		v.SetArgs1(v.MemoryArg(psess))
		v.Aux = nil
		v.AuxInt = 0
		v.Op = OpCopy
	}
}

// elimUnreadAutos deletes stores (and associated bookkeeping ops VarDef and VarKill)
// to autos that are never read from.
func (psess *PackageSession) elimUnreadAutos(f *Func) {

	seen := make(map[GCNode]bool)
	var stores []*Value
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			n, ok := v.Aux.(GCNode)
			if !ok {
				continue
			}
			if n.StorageClass() != ClassAuto {
				continue
			}

			effect := v.Op.SymEffect(psess)
			switch effect {
			case SymNone, SymWrite:

				if !seen[n] {
					stores = append(stores, v)
				}
			default:

				if v.Uses > 0 {
					seen[n] = true
				}
			}
		}
	}

	for _, store := range stores {
		n, _ := store.Aux.(GCNode)
		if seen[n] {
			continue
		}

		store.SetArgs1(store.MemoryArg(psess))
		store.Aux = nil
		store.AuxInt = 0
		store.Op = OpCopy
	}
}

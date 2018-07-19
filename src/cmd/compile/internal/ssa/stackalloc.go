package ssa

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
)

type stackAllocState struct {
	f *Func

	// live is the output of stackalloc.
	// live[b.id] = live values at the end of block b.
	live [][]ID

	// The following slices are reused across multiple users
	// of stackAllocState.
	values    []stackValState
	interfere [][]ID // interfere[v.id] = values that interfere with v.
	names     []LocalSlot
	slots     []int
	used      []bool

	nArgSlot,
	nNotNeed,
	nNamedSlot,
	nReuse,
	nAuto,
	nSelfInterfere int32 // Number of self-interferences
}

func (psess *PackageSession) newStackAllocState(f *Func) *stackAllocState {
	s := f.Cache.stackAllocState
	if s == nil {
		return new(stackAllocState)
	}
	if s.f != nil {
		f.fe.Fatalf(psess.src.NoXPos, "newStackAllocState called without previous free")
	}
	return s
}

func putStackAllocState(s *stackAllocState) {
	for i := range s.values {
		s.values[i] = stackValState{}
	}
	for i := range s.interfere {
		s.interfere[i] = nil
	}
	for i := range s.names {
		s.names[i] = LocalSlot{}
	}
	for i := range s.slots {
		s.slots[i] = 0
	}
	for i := range s.used {
		s.used[i] = false
	}
	s.f.Cache.stackAllocState = s
	s.f = nil
	s.live = nil
	s.nArgSlot, s.nNotNeed, s.nNamedSlot, s.nReuse, s.nAuto, s.nSelfInterfere = 0, 0, 0, 0, 0, 0
}

type stackValState struct {
	typ      *types.Type
	spill    *Value
	needSlot bool
	isArg    bool
}

// stackalloc allocates storage in the stack frame for
// all Values that did not get a register.
// Returns a map from block ID to the stack values live at the end of that block.
func (psess *PackageSession) stackalloc(f *Func, spillLive [][]ID) [][]ID {
	if f.pass.debug > stackDebug {
		fmt.Println("before stackalloc")
		fmt.Println(f.String(psess))
	}
	s := psess.newStackAllocState(f)
	s.init(psess, f, spillLive)
	defer putStackAllocState(s)

	s.stackalloc(psess)
	if f.pass.stats > 0 {
		f.LogStat("stack_alloc_stats",
			s.nArgSlot, "arg_slots", s.nNotNeed, "slot_not_needed",
			s.nNamedSlot, "named_slots", s.nAuto, "auto_slots",
			s.nReuse, "reused_slots", s.nSelfInterfere, "self_interfering")
	}

	return s.live
}

func (s *stackAllocState) init(psess *PackageSession, f *Func, spillLive [][]ID) {
	s.f = f

	if n := f.NumValues(); cap(s.values) >= n {
		s.values = s.values[:n]
	} else {
		s.values = make([]stackValState, n)
	}
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			s.values[v.ID].typ = v.Type
			s.values[v.ID].needSlot = !v.Type.IsMemory(psess.types) && !v.Type.IsVoid(psess.types) && !v.Type.IsFlags(psess.types) && f.getHome(v.ID) == nil && !v.rematerializeable(psess) && !v.OnWasmStack
			s.values[v.ID].isArg = v.Op == OpArg
			if f.pass.debug > stackDebug && s.values[v.ID].needSlot {
				fmt.Printf("%s needs a stack slot\n", v)
			}
			if v.Op == OpStoreReg {
				s.values[v.Args[0].ID].spill = v
			}
		}
	}

	s.computeLive(psess, spillLive)

	s.buildInterferenceGraph(psess)
}

func (s *stackAllocState) stackalloc(psess *PackageSession) {
	f := s.f

	if n := f.NumValues(); cap(s.names) >= n {
		s.names = s.names[:n]
	} else {
		s.names = make([]LocalSlot, n)
	}
	names := s.names
	for _, name := range f.Names {

		for _, v := range f.NamedValues[name] {
			names[v.ID] = name
		}
	}

	for _, v := range f.Entry.Values {
		if v.Op != OpArg {
			continue
		}
		loc := LocalSlot{N: v.Aux.(GCNode), Type: v.Type, Off: v.AuxInt}
		if f.pass.debug > stackDebug {
			fmt.Printf("stackalloc %s to %s\n", v, loc)
		}
		f.setHome(v, loc)
	}

	locations := map[*types.Type][]LocalSlot{}

	slots := s.slots
	if n := f.NumValues(); cap(slots) >= n {
		slots = slots[:n]
	} else {
		slots = make([]int, n)
		s.slots = slots
	}
	for i := range slots {
		slots[i] = -1
	}

	// Pick a stack slot for each value needing one.
	var used []bool
	if n := f.NumValues(); cap(s.used) >= n {
		used = s.used[:n]
	} else {
		used = make([]bool, n)
		s.used = used
	}
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if !s.values[v.ID].needSlot {
				s.nNotNeed++
				continue
			}
			if v.Op == OpArg {
				s.nArgSlot++
				continue
			}

			// If this is a named value, try to use the name as
			// the spill location.
			var name LocalSlot
			if v.Op == OpStoreReg {
				name = names[v.Args[0].ID]
			} else {
				name = names[v.ID]
			}
			if name.N != nil && v.Type.Compare(psess.types, name.Type) == types.CMPeq {
				for _, id := range s.interfere[v.ID] {
					h := f.getHome(id)
					if h != nil && h.(LocalSlot).N == name.N && h.(LocalSlot).Off == name.Off {

						s.nSelfInterfere++
						goto noname
					}
				}
				if f.pass.debug > stackDebug {
					fmt.Printf("stackalloc %s to %s\n", v, name)
				}
				s.nNamedSlot++
				f.setHome(v, name)
				continue
			}

		noname:

			locs := locations[v.Type]

			for i := 0; i < len(locs); i++ {
				used[i] = false
			}
			for _, xid := range s.interfere[v.ID] {
				slot := slots[xid]
				if slot >= 0 {
					used[slot] = true
				}
			}
			// Find an unused stack slot.
			var i int
			for i = 0; i < len(locs); i++ {
				if !used[i] {
					s.nReuse++
					break
				}
			}

			if i == len(locs) {
				s.nAuto++
				locs = append(locs, LocalSlot{N: f.fe.Auto(v.Pos, v.Type), Type: v.Type, Off: 0})
				locations[v.Type] = locs
			}

			loc := locs[i]
			if f.pass.debug > stackDebug {
				fmt.Printf("stackalloc %s to %s\n", v, loc)
			}
			f.setHome(v, loc)
			slots[v.ID] = i
		}
	}
}

// computeLive computes a map from block ID to a list of
// stack-slot-needing value IDs live at the end of that block.
// TODO: this could be quadratic if lots of variables are live across lots of
// basic blocks. Figure out a way to make this function (or, more precisely, the user
// of this function) require only linear size & time.
func (s *stackAllocState) computeLive(psess *PackageSession, spillLive [][]ID) {
	s.live = make([][]ID, s.f.NumBlocks())
	var phis []*Value
	live := s.f.newSparseSet(s.f.NumValues())
	defer s.f.retSparseSet(live)
	t := s.f.newSparseSet(s.f.NumValues())
	defer s.f.retSparseSet(t)

	po := s.f.postorder()
	for {
		changed := false
		for _, b := range po {

			live.clear()
			live.addAll(s.live[b.ID])

			phis = phis[:0]
			for i := len(b.Values) - 1; i >= 0; i-- {
				v := b.Values[i]
				live.remove(v.ID)
				if v.Op == OpPhi {

					if !v.Type.IsMemory(psess.types) && !v.Type.IsVoid(psess.types) {
						phis = append(phis, v)
					}
					continue
				}
				for _, a := range v.Args {
					if s.values[a.ID].needSlot {
						live.add(a.ID)
					}
				}
			}

			for i, e := range b.Preds {
				p := e.b
				t.clear()
				t.addAll(s.live[p.ID])
				t.addAll(live.contents())
				t.addAll(spillLive[p.ID])
				for _, v := range phis {
					a := v.Args[i]
					if s.values[a.ID].needSlot {
						t.add(a.ID)
					}
					if spill := s.values[a.ID].spill; spill != nil {

						t.add(spill.ID)
					}
				}
				if t.size() == len(s.live[p.ID]) {
					continue
				}

				s.live[p.ID] = append(s.live[p.ID][:0], t.contents()...)
				changed = true
			}
		}

		if !changed {
			break
		}
	}
	if s.f.pass.debug > stackDebug {
		for _, b := range s.f.Blocks {
			fmt.Printf("stacklive %s %v\n", b, s.live[b.ID])
		}
	}
}

func (f *Func) getHome(vid ID) Location {
	if int(vid) >= len(f.RegAlloc) {
		return nil
	}
	return f.RegAlloc[vid]
}

func (f *Func) setHome(v *Value, loc Location) {
	for v.ID >= ID(len(f.RegAlloc)) {
		f.RegAlloc = append(f.RegAlloc, nil)
	}
	f.RegAlloc[v.ID] = loc
}

func (s *stackAllocState) buildInterferenceGraph(psess *PackageSession) {
	f := s.f
	if n := f.NumValues(); cap(s.interfere) >= n {
		s.interfere = s.interfere[:n]
	} else {
		s.interfere = make([][]ID, n)
	}
	live := f.newSparseSet(f.NumValues())
	defer f.retSparseSet(live)
	for _, b := range f.Blocks {

		live.clear()
		live.addAll(s.live[b.ID])
		for i := len(b.Values) - 1; i >= 0; i-- {
			v := b.Values[i]
			if s.values[v.ID].needSlot {
				live.remove(v.ID)
				for _, id := range live.contents() {

					if s.values[v.ID].typ.Compare(psess.types, s.values[id].typ) == types.CMPeq || v.Op == OpArg || s.values[id].isArg {
						s.interfere[v.ID] = append(s.interfere[v.ID], id)
						s.interfere[id] = append(s.interfere[id], v.ID)
					}
				}
			}
			for _, a := range v.Args {
				if s.values[a.ID].needSlot {
					live.add(a.ID)
				}
			}
			if v.Op == OpArg && s.values[v.ID].needSlot {

				live.add(v.ID)
			}
		}
	}
	if f.pass.debug > stackDebug {
		for vid, i := range s.interfere {
			if len(i) > 0 {
				fmt.Printf("v%d interferes with", vid)
				for _, x := range i {
					fmt.Printf(" v%d", x)
				}
				fmt.Println()
			}
		}
	}
}

package ssa

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"

	"github.com/dave/golib/src/cmd/internal/src"

	"unsafe"
)

const (
	moveSpills = iota
	logSpills
	regDebug
	stackDebug
)

// distance is a measure of how far into the future values are used.
// distance is measured in units of instructions.
const (
	likelyDistance   = 1
	normalDistance   = 10
	unlikelyDistance = 100
)

// regalloc performs register allocation on f. It sets f.RegAlloc
// to the resulting allocation.
func (psess *PackageSession) regalloc(f *Func) {
	var s regAllocState
	s.init(psess, f)
	s.regalloc(psess, f)
}

type register uint8

const noRegister register = 255

type regMask uint64

func (m regMask) String() string {
	s := ""
	for r := register(0); m != 0; r++ {
		if m>>r&1 == 0 {
			continue
		}
		m &^= regMask(1) << r
		if s != "" {
			s += " "
		}
		s += fmt.Sprintf("r%d", r)
	}
	return s
}

func (s *regAllocState) RegMaskString(m regMask) string {
	str := ""
	for r := register(0); m != 0; r++ {
		if m>>r&1 == 0 {
			continue
		}
		m &^= regMask(1) << r
		if str != "" {
			str += " "
		}
		str += s.registers[r].String()
	}
	return str
}

// countRegs returns the number of set bits in the register mask.
func countRegs(r regMask) int {
	n := 0
	for r != 0 {
		n += int(r & 1)
		r >>= 1
	}
	return n
}

// pickReg picks an arbitrary register from the register mask.
func pickReg(r regMask) register {

	if r == 0 {
		panic("can't pick a register from an empty set")
	}
	for i := register(0); ; i++ {
		if r&1 != 0 {
			return i
		}
		r >>= 1
	}
}

type use struct {
	dist int32    // distance from start of the block to a use of a value
	pos  src.XPos // source position of the use
	next *use     // linked list of uses of a value in nondecreasing dist order
}

// A valState records the register allocation state for a (pre-regalloc) value.
type valState struct {
	regs              regMask // the set of registers holding a Value (usually just one)
	uses              *use    // list of uses in this block
	spill             *Value  // spilled copy of the Value (if any)
	restoreMin        int32   // minimum of all restores' blocks' sdom.entry
	restoreMax        int32   // maximum of all restores' blocks' sdom.exit
	needReg           bool    // cached value of !v.Type.IsMemory() && !v.Type.IsVoid() && !.v.Type.IsFlags()
	rematerializeable bool    // cached value of v.rematerializeable()
}

type regState struct {
	v *Value // Original (preregalloc) Value stored in this register.
	c *Value // A Value equal to v which is currently in a register.  Might be v or a copy of it.

}

type regAllocState struct {
	f *Func

	sdom        SparseTree
	registers   []Register
	numRegs     register
	SPReg       register
	SBReg       register
	GReg        register
	allocatable regMask

	// for each block, its primary predecessor.
	// A predecessor of b is primary if it is the closest
	// predecessor that appears before b in the layout order.
	// We record the index in the Preds list where the primary predecessor sits.
	primary []int32

	// live values at the end of each block.  live[b.ID] is a list of value IDs
	// which are live at the end of b, together with a count of how many instructions
	// forward to the next use.
	live [][]liveInfo
	// desired register assignments at the end of each block.
	// Note that this is a static map computed before allocation occurs. Dynamic
	// register desires (from partially completed allocations) will trump
	// this information.
	desired []desiredState

	// current state of each (preregalloc) Value
	values []valState

	// ID of SP, SB values
	sp, sb ID

	// For each Value, map from its value ID back to the
	// preregalloc Value it was derived from.
	orig []*Value

	// current state of each register
	regs []regState

	// registers that contain values which can't be kicked out
	nospill regMask

	// mask of registers currently in use
	used regMask

	// mask of registers used in the current instruction
	tmpused regMask

	// current block we're working on
	curBlock *Block

	// cache of use records
	freeUseRecords *use

	// endRegs[blockid] is the register state at the end of each block.
	// encoded as a set of endReg records.
	endRegs [][]endReg

	// startRegs[blockid] is the register state at the start of merge blocks.
	// saved state does not include the state of phi ops in the block.
	startRegs [][]startReg

	// spillLive[blockid] is the set of live spills at the end of each block
	spillLive [][]ID

	// a set of copies we generated to move things around, and
	// whether it is used in shuffle. Unused copies will be deleted.
	copies map[*Value]bool

	loopnest *loopnest

	// choose a good order in which to visit blocks for allocation purposes.
	visitOrder []*Block
}

type endReg struct {
	r register
	v *Value // pre-regalloc value held in this register (TODO: can we use ID here?)
	c *Value // cached version of the value
}

type startReg struct {
	r   register
	v   *Value   // pre-regalloc value needed in this register
	c   *Value   // cached version of the value
	pos src.XPos // source position of use of this register
}

// freeReg frees up register r. Any current user of r is kicked out.
func (s *regAllocState) freeReg(r register) {
	v := s.regs[r].v
	if v == nil {
		s.f.Fatalf("tried to free an already free register %d\n", r)
	}

	if s.f.pass.debug > regDebug {
		fmt.Printf("freeReg %s (dump %s/%s)\n", &s.registers[r], v, s.regs[r].c)
	}
	s.regs[r] = regState{}
	s.values[v.ID].regs &^= regMask(1) << r
	s.used &^= regMask(1) << r
}

// freeRegs frees up all registers listed in m.
func (s *regAllocState) freeRegs(m regMask) {
	for m&s.used != 0 {
		s.freeReg(pickReg(m & s.used))
	}
}

// setOrig records that c's original value is the same as
// v's original value.
func (s *regAllocState) setOrig(c *Value, v *Value) {
	for int(c.ID) >= len(s.orig) {
		s.orig = append(s.orig, nil)
	}
	if s.orig[c.ID] != nil {
		s.f.Fatalf("orig value set twice %s %s", c, v)
	}
	s.orig[c.ID] = s.orig[v.ID]
}

// assignReg assigns register r to hold c, a copy of v.
// r must be unused.
func (s *regAllocState) assignReg(r register, v *Value, c *Value) {
	if s.f.pass.debug > regDebug {
		fmt.Printf("assignReg %s %s/%s\n", &s.registers[r], v, c)
	}
	if s.regs[r].v != nil {
		s.f.Fatalf("tried to assign register %d to %s/%s but it is already used by %s", r, v, c, s.regs[r].v)
	}

	s.regs[r] = regState{v, c}
	s.values[v.ID].regs |= regMask(1) << r
	s.used |= regMask(1) << r
	s.f.setHome(c, &s.registers[r])
}

// allocReg chooses a register from the set of registers in mask.
// If there is no unused register, a Value will be kicked out of
// a register to make room.
func (s *regAllocState) allocReg(psess *PackageSession, mask regMask, v *Value) register {
	if v.OnWasmStack {
		return noRegister
	}

	mask &= s.allocatable
	mask &^= s.nospill
	if mask == 0 {
		s.f.Fatalf("no register available for %s", v.LongString(psess))
	}

	if mask&^s.used != 0 {
		return pickReg(mask &^ s.used)
	}

	// Find a register to spill. We spill the register containing the value
	// whose next use is as far in the future as possible.
	// https://en.wikipedia.org/wiki/Page_replacement_algorithm#The_theoretically_optimal_page_replacement_algorithm
	var r register
	maxuse := int32(-1)
	for t := register(0); t < s.numRegs; t++ {
		if mask>>t&1 == 0 {
			continue
		}
		v := s.regs[t].v
		if n := s.values[v.ID].uses.dist; n > maxuse {

			r = t
			maxuse = n
		}
	}
	if maxuse == -1 {
		s.f.Fatalf("couldn't find register to spill")
	}

	if s.f.Config.ctxt.Arch.Arch == psess.sys.ArchWasm {

		s.freeReg(r)
		return r
	}

	v2 := s.regs[r].v
	m := s.compatRegs(psess, v2.Type) &^ s.used &^ s.tmpused &^ (regMask(1) << r)
	if m != 0 && !s.values[v2.ID].rematerializeable && countRegs(s.values[v2.ID].regs) == 1 {
		r2 := pickReg(m)
		c := s.curBlock.NewValue1(v2.Pos, OpCopy, v2.Type, s.regs[r].c)
		s.copies[c] = false
		if s.f.pass.debug > regDebug {
			fmt.Printf("copy %s to %s : %s\n", v2, c, &s.registers[r2])
		}
		s.setOrig(c, v2)
		s.assignReg(r2, v2, c)
	}
	s.freeReg(r)
	return r
}

// makeSpill returns a Value which represents the spilled value of v.
// b is the block in which the spill is used.
func (s *regAllocState) makeSpill(v *Value, b *Block) *Value {
	vi := &s.values[v.ID]
	if vi.spill != nil {

		vi.restoreMin = min32(vi.restoreMin, s.sdom[b.ID].entry)
		vi.restoreMax = max32(vi.restoreMax, s.sdom[b.ID].exit)
		return vi.spill
	}

	spill := s.f.newValueNoBlock(OpStoreReg, v.Type, v.Pos)

	s.setOrig(spill, v)
	vi.spill = spill
	vi.restoreMin = s.sdom[b.ID].entry
	vi.restoreMax = s.sdom[b.ID].exit
	return spill
}

// allocValToReg allocates v to a register selected from regMask and
// returns the register copy of v. Any previous user is kicked out and spilled
// (if necessary). Load code is added at the current pc. If nospill is set the
// allocated register is marked nospill so the assignment cannot be
// undone until the caller allows it by clearing nospill. Returns a
// *Value which is either v or a copy of v allocated to the chosen register.
func (s *regAllocState) allocValToReg(psess *PackageSession, v *Value, mask regMask, nospill bool, pos src.XPos) *Value {
	if s.f.Config.ctxt.Arch.Arch == psess.sys.ArchWasm && v.rematerializeable(psess) {
		c := v.copyIntoWithXPos(psess, s.curBlock, pos)
		c.OnWasmStack = true
		s.setOrig(c, v)
		return c
	}
	if v.OnWasmStack {
		return v
	}

	vi := &s.values[v.ID]
	pos = pos.WithNotStmt()

	if mask&vi.regs != 0 {
		r := pickReg(mask & vi.regs)
		if s.regs[r].v != v || s.regs[r].c == nil {
			panic("bad register state")
		}
		if nospill {
			s.nospill |= regMask(1) << r
		}
		return s.regs[r].c
	}

	var r register

	onWasmStack := nospill && s.f.Config.ctxt.Arch.Arch == psess.sys.ArchWasm
	if !onWasmStack {

		r = s.allocReg(psess, mask, v)
	}

	// Allocate v to the new register.
	var c *Value
	if vi.regs != 0 {

		r2 := pickReg(vi.regs)
		if s.regs[r2].v != v {
			panic("bad register state")
		}
		c = s.curBlock.NewValue1(pos, OpCopy, v.Type, s.regs[r2].c)
	} else if v.rematerializeable(psess) {

		c = v.copyIntoWithXPos(psess, s.curBlock, pos)
	} else {

		spill := s.makeSpill(v, s.curBlock)
		if s.f.pass.debug > logSpills {
			s.f.Warnl(vi.spill.Pos, "load spill for %v from %v", v, spill)
		}
		c = s.curBlock.NewValue1(pos, OpLoadReg, v.Type, spill)
	}

	s.setOrig(c, v)

	if onWasmStack {
		c.OnWasmStack = true
		return c
	}

	s.assignReg(r, v, c)
	if c.Op == OpLoadReg && s.isGReg(r) {
		s.f.Fatalf("allocValToReg.OpLoadReg targeting g: " + c.LongString(psess))
	}
	if nospill {
		s.nospill |= regMask(1) << r
	}
	return c
}

// isLeaf reports whether f performs any calls.
func (psess *PackageSession) isLeaf(f *Func) bool {
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if psess.opcodeTable[v.Op].call {
				return false
			}
		}
	}
	return true
}

func (s *regAllocState) init(psess *PackageSession, f *Func) {
	s.f = f
	s.f.RegAlloc = s.f.Cache.locs[:0]
	s.registers = f.Config.registers
	if nr := len(s.registers); nr == 0 || nr > int(noRegister) || nr > int(unsafe.Sizeof(regMask(0))*8) {
		s.f.Fatalf("bad number of registers: %d", nr)
	} else {
		s.numRegs = register(nr)
	}

	s.SPReg = noRegister
	s.SBReg = noRegister
	s.GReg = noRegister
	for r := register(0); r < s.numRegs; r++ {
		switch s.registers[r].String() {
		case "SP":
			s.SPReg = r
		case "SB":
			s.SBReg = r
		case "g":
			s.GReg = r
		}
	}

	switch noRegister {
	case s.SPReg:
		s.f.Fatalf("no SP register found")
	case s.SBReg:
		s.f.Fatalf("no SB register found")
	case s.GReg:
		if f.Config.hasGReg {
			s.f.Fatalf("no g register found")
		}
	}

	s.allocatable = s.f.Config.gpRegMask | s.f.Config.fpRegMask | s.f.Config.specialRegMask
	s.allocatable &^= 1 << s.SPReg
	s.allocatable &^= 1 << s.SBReg
	if s.f.Config.hasGReg {
		s.allocatable &^= 1 << s.GReg
	}
	if s.f.Config.ctxt.Framepointer_enabled && s.f.Config.FPReg >= 0 {
		s.allocatable &^= 1 << uint(s.f.Config.FPReg)
	}
	if s.f.Config.LinkReg != -1 {
		if psess.isLeaf(f) {

			s.allocatable &^= 1 << uint(s.f.Config.LinkReg)
		}
		if s.f.Config.arch == "arm" && psess.objabi.GOARM == 5 {

			s.allocatable &^= 1 << uint(s.f.Config.LinkReg)
		}
	}
	if s.f.Config.ctxt.Flag_dynlink {
		switch s.f.Config.arch {
		case "amd64":
			s.allocatable &^= 1 << 15
		case "arm":
			s.allocatable &^= 1 << 9
		case "ppc64le":

		case "arm64":

		case "386":

		case "s390x":
			s.allocatable &^= 1 << 11
		default:
			s.f.fe.Fatalf(psess.src.NoXPos, "arch %s not implemented", s.f.Config.arch)
		}
	}
	if s.f.Config.nacl {
		switch s.f.Config.arch {
		case "arm":
			s.allocatable &^= 1 << 9
		case "amd64p32":
			s.allocatable &^= 1 << 5
			s.allocatable &^= 1 << 15
		}
	}
	if s.f.Config.use387 {
		s.allocatable &^= 1 << 15
	}

	s.visitOrder = layoutRegallocOrder(f)

	blockOrder := make([]int32, f.NumBlocks())
	for i, b := range s.visitOrder {
		blockOrder[b.ID] = int32(i)
	}

	s.regs = make([]regState, s.numRegs)
	s.values = make([]valState, f.NumValues())
	s.orig = make([]*Value, f.NumValues())
	s.copies = make(map[*Value]bool)
	for _, b := range s.visitOrder {
		for _, v := range b.Values {
			if !v.Type.IsMemory(psess.types) && !v.Type.IsVoid(psess.types) && !v.Type.IsFlags(psess.types) && !v.Type.IsTuple() {
				s.values[v.ID].needReg = true
				s.values[v.ID].rematerializeable = v.rematerializeable(psess)
				s.orig[v.ID] = v
			}

		}
	}
	s.computeLive(psess)

	s.primary = make([]int32, f.NumBlocks())
	for _, b := range s.visitOrder {
		best := -1
		for i, e := range b.Preds {
			p := e.b
			if blockOrder[p.ID] >= blockOrder[b.ID] {
				continue
			}
			if best == -1 || blockOrder[p.ID] > blockOrder[b.Preds[best].b.ID] {
				best = i
			}
		}
		s.primary[b.ID] = int32(best)
	}

	s.endRegs = make([][]endReg, f.NumBlocks())
	s.startRegs = make([][]startReg, f.NumBlocks())
	s.spillLive = make([][]ID, f.NumBlocks())
	s.sdom = f.sdom()

	if f.Config.ctxt.Arch.Arch == psess.sys.ArchWasm {
		canLiveOnStack := f.newSparseSet(f.NumValues())
		defer f.retSparseSet(canLiveOnStack)
		for _, b := range f.Blocks {

			canLiveOnStack.clear()
			if b.Control != nil && b.Control.Uses == 1 && !psess.opcodeTable[b.Control.Op].generic {
				canLiveOnStack.add(b.Control.ID)
			}

			for i := len(b.Values) - 1; i >= 0; i-- {
				v := b.Values[i]
				if canLiveOnStack.contains(v.ID) {
					v.OnWasmStack = true
				} else {

					canLiveOnStack.clear()
				}
				for _, arg := range v.Args {

					if arg.Uses == 1 && arg.Block == v.Block && !arg.Type.IsMemory(psess.types) && !psess.opcodeTable[arg.Op].generic {
						canLiveOnStack.add(arg.ID)
					}
				}
			}
		}
	}
}

// Adds a use record for id at distance dist from the start of the block.
// All calls to addUse must happen with nonincreasing dist.
func (s *regAllocState) addUse(id ID, dist int32, pos src.XPos) {
	r := s.freeUseRecords
	if r != nil {
		s.freeUseRecords = r.next
	} else {
		r = &use{}
	}
	r.dist = dist
	r.pos = pos
	r.next = s.values[id].uses
	s.values[id].uses = r
	if r.next != nil && dist > r.next.dist {
		s.f.Fatalf("uses added in wrong order")
	}
}

// advanceUses advances the uses of v's args from the state before v to the state after v.
// Any values which have no more uses are deallocated from registers.
func (s *regAllocState) advanceUses(v *Value) {
	for _, a := range v.Args {
		if !s.values[a.ID].needReg {
			continue
		}
		ai := &s.values[a.ID]
		r := ai.uses
		ai.uses = r.next
		if r.next == nil {

			s.freeRegs(ai.regs)
		}
		r.next = s.freeUseRecords
		s.freeUseRecords = r
	}
}

// liveAfterCurrentInstruction reports whether v is live after
// the current instruction is completed.  v must be used by the
// current instruction.
func (s *regAllocState) liveAfterCurrentInstruction(v *Value) bool {
	u := s.values[v.ID].uses
	d := u.dist
	for u != nil && u.dist == d {
		u = u.next
	}
	return u != nil && u.dist > d
}

// Sets the state of the registers to that encoded in regs.
func (s *regAllocState) setState(regs []endReg) {
	s.freeRegs(s.used)
	for _, x := range regs {
		s.assignReg(x.r, x.v, x.c)
	}
}

// compatRegs returns the set of registers which can store a type t.
func (s *regAllocState) compatRegs(psess *PackageSession, t *types.Type) regMask {
	var m regMask
	if t.IsTuple() || t.IsFlags(psess.types) {
		return 0
	}
	if t.IsFloat() || t == psess.types.TypeInt128 {
		m = s.f.Config.fpRegMask
	} else {
		m = s.f.Config.gpRegMask
	}
	return m & s.allocatable
}

// regspec returns the regInfo for operation op.
func (s *regAllocState) regspec(psess *PackageSession, op Op) regInfo {
	if op == OpConvert {

		m := s.allocatable & s.f.Config.gpRegMask
		return regInfo{inputs: []inputInfo{{regs: m}}, outputs: []outputInfo{{regs: m}}}
	}
	return psess.opcodeTable[op].reg
}

func (s *regAllocState) isGReg(r register) bool {
	return s.f.Config.hasGReg && s.GReg == r
}

func (s *regAllocState) regalloc(psess *PackageSession, f *Func) {
	regValLiveSet := f.newSparseSet(f.NumValues())
	defer f.retSparseSet(regValLiveSet)
	var oldSched []*Value
	var phis []*Value
	var phiRegs []register
	var args []*Value

	// Data structure used for computing desired registers.
	var desired desiredState

	// Desired registers for inputs & outputs for each instruction in the block.
	type dentry struct {
		out [4]register    // desired output registers
		in  [3][4]register // desired input registers (for inputs 0,1, and 2)
	}
	var dinfo []dentry

	if f.Entry != f.Blocks[0] {
		f.Fatalf("entry block must be first")
	}

	for _, b := range s.visitOrder {
		if s.f.pass.debug > regDebug {
			fmt.Printf("Begin processing block %v\n", b)
		}
		s.curBlock = b

		regValLiveSet.clear()
		for _, e := range s.live[b.ID] {
			s.addUse(e.ID, int32(len(b.Values))+e.dist, e.pos)
			regValLiveSet.add(e.ID)
		}
		if v := b.Control; v != nil && s.values[v.ID].needReg {
			s.addUse(v.ID, int32(len(b.Values)), b.Pos)
			regValLiveSet.add(v.ID)
		}
		for i := len(b.Values) - 1; i >= 0; i-- {
			v := b.Values[i]
			regValLiveSet.remove(v.ID)
			if v.Op == OpPhi {

				continue
			}
			if psess.opcodeTable[v.Op].call {

				regValLiveSet.clear()
				if s.sp != 0 && s.values[s.sp].uses != nil {
					regValLiveSet.add(s.sp)
				}
				if s.sb != 0 && s.values[s.sb].uses != nil {
					regValLiveSet.add(s.sb)
				}
			}
			for _, a := range v.Args {
				if !s.values[a.ID].needReg {
					continue
				}
				s.addUse(a.ID, int32(i), v.Pos)
				regValLiveSet.add(a.ID)
			}
		}
		if s.f.pass.debug > regDebug {
			fmt.Printf("use distances for %s\n", b)
			for i := range s.values {
				vi := &s.values[i]
				u := vi.uses
				if u == nil {
					continue
				}
				fmt.Printf("  v%d:", i)
				for u != nil {
					fmt.Printf(" %d", u.dist)
					u = u.next
				}
				fmt.Println()
			}
		}

		nphi := 0
		for _, v := range b.Values {
			if v.Op != OpPhi {
				break
			}
			nphi++
		}
		phis = append(phis[:0], b.Values[:nphi]...)
		oldSched = append(oldSched[:0], b.Values[nphi:]...)
		b.Values = b.Values[:0]

		if b == f.Entry {

			if nphi > 0 {
				f.Fatalf("phis in entry block")
			}
		} else if len(b.Preds) == 1 {

			s.setState(s.endRegs[b.Preds[0].b.ID])
			if nphi > 0 {
				f.Fatalf("phis in single-predecessor block")
			}

			for r := register(0); r < s.numRegs; r++ {
				v := s.regs[r].v
				if v != nil && !regValLiveSet.contains(v.ID) {
					s.freeReg(r)
				}
			}
		} else {

			idx := s.primary[b.ID]
			if idx < 0 {
				f.Fatalf("block with no primary predecessor %s", b)
			}
			p := b.Preds[idx].b
			s.setState(s.endRegs[p.ID])

			if s.f.pass.debug > regDebug {
				fmt.Printf("starting merge block %s with end state of %s:\n", b, p)
				for _, x := range s.endRegs[p.ID] {
					fmt.Printf("  %s: orig:%s cache:%s\n", &s.registers[x.r], x.v, x.c)
				}
			}

			phiRegs = phiRegs[:0]
			var phiUsed regMask

			for _, v := range phis {
				if !s.values[v.ID].needReg {
					phiRegs = append(phiRegs, noRegister)
					continue
				}
				a := v.Args[idx]

				m := s.values[a.ID].regs &^ phiUsed & s.allocatable
				if m != 0 {
					r := pickReg(m)
					phiUsed |= regMask(1) << r
					phiRegs = append(phiRegs, r)
				} else {
					phiRegs = append(phiRegs, noRegister)
				}
			}

			for i, v := range phis {
				if !s.values[v.ID].needReg {
					continue
				}
				a := v.Args[idx]
				if !regValLiveSet.contains(a.ID) {

					s.freeRegs(s.values[a.ID].regs)
				} else {

					r := phiRegs[i]
					if r == noRegister {
						continue
					}

					m := s.compatRegs(psess, a.Type) &^ s.used &^ phiUsed
					if m != 0 && !s.values[a.ID].rematerializeable && countRegs(s.values[a.ID].regs) == 1 {
						r2 := pickReg(m)
						c := p.NewValue1(a.Pos, OpCopy, a.Type, s.regs[r].c)
						s.copies[c] = false
						if s.f.pass.debug > regDebug {
							fmt.Printf("copy %s to %s : %s\n", a, c, &s.registers[r2])
						}
						s.setOrig(c, a)
						s.assignReg(r2, a, c)
						s.endRegs[p.ID] = append(s.endRegs[p.ID], endReg{r2, a, c})
					}
					s.freeReg(r)
				}
			}

			b.Values = append(b.Values, phis...)

			for i, v := range phis {
				if !s.values[v.ID].needReg {
					continue
				}
				if phiRegs[i] != noRegister {
					continue
				}
				if s.f.Config.use387 && v.Type.IsFloat() {
					continue
				}
				m := s.compatRegs(psess, v.Type) &^ phiUsed &^ s.used
				if m != 0 {
					r := pickReg(m)
					phiRegs[i] = r
					phiUsed |= regMask(1) << r
				}
			}

			for i, v := range phis {
				if !s.values[v.ID].needReg {
					continue
				}
				r := phiRegs[i]
				if r == noRegister {

					s.values[v.ID].spill = v
					continue
				}

				s.assignReg(r, v, v)
			}

			for r := register(0); r < s.numRegs; r++ {
				if phiUsed>>r&1 != 0 {
					continue
				}
				v := s.regs[r].v
				if v != nil && !regValLiveSet.contains(v.ID) {
					s.freeReg(r)
				}
			}

			regList := make([]startReg, 0, 32)
			for r := register(0); r < s.numRegs; r++ {
				v := s.regs[r].v
				if v == nil {
					continue
				}
				if phiUsed>>r&1 != 0 {

					continue
				}
				regList = append(regList, startReg{r, v, s.regs[r].c, s.values[v.ID].uses.pos})
			}
			s.startRegs[b.ID] = make([]startReg, len(regList))
			copy(s.startRegs[b.ID], regList)

			if s.f.pass.debug > regDebug {
				fmt.Printf("after phis\n")
				for _, x := range s.startRegs[b.ID] {
					fmt.Printf("  %s: v%d\n", &s.registers[x.r], x.v.ID)
				}
			}
		}

		if l := len(oldSched); cap(dinfo) < l {
			dinfo = make([]dentry, l)
		} else {
			dinfo = dinfo[:l]
			for i := range dinfo {
				dinfo[i] = dentry{}
			}
		}

		desired.copy(&s.desired[b.ID])

		for _, e := range b.Succs {
			succ := e.b

			for _, x := range s.startRegs[succ.ID] {
				desired.add(x.v.ID, x.r)
			}

			pidx := e.i
			for _, v := range succ.Values {
				if v.Op != OpPhi {
					break
				}
				if !s.values[v.ID].needReg {
					continue
				}
				rp, ok := s.f.getHome(v.ID).(*Register)
				if !ok {
					continue
				}
				desired.add(v.Args[pidx].ID, register(rp.num))
			}
		}

		for i := len(oldSched) - 1; i >= 0; i-- {
			v := oldSched[i]
			prefs := desired.remove(v.ID)
			regspec := s.regspec(psess, v.Op)
			desired.clobber(regspec.clobbers)
			for _, j := range regspec.inputs {
				if countRegs(j.regs) != 1 {
					continue
				}
				desired.clobber(j.regs)
				desired.add(v.Args[j.idx].ID, pickReg(j.regs))
			}
			if psess.opcodeTable[v.Op].resultInArg0 {
				if psess.opcodeTable[v.Op].commutative {
					desired.addList(v.Args[1].ID, prefs)
				}
				desired.addList(v.Args[0].ID, prefs)
			}

			dinfo[i].out = prefs
			for j, a := range v.Args {
				if j >= len(dinfo[i].in) {
					break
				}
				dinfo[i].in[j] = desired.get(a.ID)
			}
		}

		for idx, v := range oldSched {
			if s.f.pass.debug > regDebug {
				fmt.Printf("  processing %s\n", v.LongString(psess))
			}
			regspec := s.regspec(psess, v.Op)
			if v.Op == OpPhi {
				f.Fatalf("phi %s not at start of block", v)
			}
			if v.Op == OpSP {
				s.assignReg(s.SPReg, v, v)
				b.Values = append(b.Values, v)
				s.advanceUses(v)
				s.sp = v.ID
				continue
			}
			if v.Op == OpSB {
				s.assignReg(s.SBReg, v, v)
				b.Values = append(b.Values, v)
				s.advanceUses(v)
				s.sb = v.ID
				continue
			}
			if v.Op == OpSelect0 || v.Op == OpSelect1 {
				if s.values[v.ID].needReg {
					var i = 0
					if v.Op == OpSelect1 {
						i = 1
					}
					s.assignReg(register(s.f.getHome(v.Args[0].ID).(LocPair)[i].(*Register).num), v, v)
				}
				b.Values = append(b.Values, v)
				s.advanceUses(v)
				goto issueSpill
			}
			if v.Op == OpGetG && s.f.Config.hasGReg {

				if s.regs[s.GReg].v != nil {
					s.freeReg(s.GReg)
				}
				s.assignReg(s.GReg, v, v)
				b.Values = append(b.Values, v)
				s.advanceUses(v)
				goto issueSpill
			}
			if v.Op == OpArg {

				s.values[v.ID].spill = v
				b.Values = append(b.Values, v)
				s.advanceUses(v)
				continue
			}
			if v.Op == OpKeepAlive {

				s.advanceUses(v)
				a := v.Args[0]
				vi := &s.values[a.ID]
				if vi.regs == 0 && !vi.rematerializeable {

					v.SetArg(0, s.makeSpill(a, b))
				} else {

					v.Op = OpCopy
					v.SetArgs1(v.Args[1])
				}
				b.Values = append(b.Values, v)
				continue
			}
			if len(regspec.inputs) == 0 && len(regspec.outputs) == 0 {

				s.freeRegs(regspec.clobbers)
				b.Values = append(b.Values, v)
				s.advanceUses(v)
				continue
			}

			if s.values[v.ID].rematerializeable {

				for _, a := range v.Args {
					a.Uses--
				}
				s.advanceUses(v)
				continue
			}

			if s.f.pass.debug > regDebug {
				fmt.Printf("value %s\n", v.LongString(psess))
				fmt.Printf("  out:")
				for _, r := range dinfo[idx].out {
					if r != noRegister {
						fmt.Printf(" %s", &s.registers[r])
					}
				}
				fmt.Println()
				for i := 0; i < len(v.Args) && i < 3; i++ {
					fmt.Printf("  in%d:", i)
					for _, r := range dinfo[idx].in[i] {
						if r != noRegister {
							fmt.Printf(" %s", &s.registers[r])
						}
					}
					fmt.Println()
				}
			}

			args = append(args[:0], v.Args...)
			for _, i := range regspec.inputs {
				mask := i.regs
				if mask&s.values[args[i.idx].ID].regs == 0 {

					mask &= s.allocatable
					mask &^= s.nospill

					if i.idx < 3 {
						for _, r := range dinfo[idx].in[i.idx] {
							if r != noRegister && (mask&^s.used)>>r&1 != 0 {

								mask = regMask(1) << r
								break
							}
						}
					}

					if mask&^desired.avoid != 0 {
						mask &^= desired.avoid
					}
				}
				args[i.idx] = s.allocValToReg(psess, args[i.idx], mask, true, v.Pos)
			}

			if psess.opcodeTable[v.Op].resultInArg0 {
				var m regMask
				if !s.liveAfterCurrentInstruction(v.Args[0]) {

					goto ok
				}
				if s.values[v.Args[0].ID].rematerializeable {

					goto ok
				}
				if countRegs(s.values[v.Args[0].ID].regs) >= 2 {

					goto ok
				}
				if psess.opcodeTable[v.Op].commutative {
					if !s.liveAfterCurrentInstruction(v.Args[1]) {
						args[0], args[1] = args[1], args[0]
						goto ok
					}
					if s.values[v.Args[1].ID].rematerializeable {
						args[0], args[1] = args[1], args[0]
						goto ok
					}
					if countRegs(s.values[v.Args[1].ID].regs) >= 2 {
						args[0], args[1] = args[1], args[0]
						goto ok
					}
				}

				m = s.compatRegs(psess, v.Args[0].Type) &^ s.used
				if m == 0 {

					goto ok
				}

				for _, r := range dinfo[idx].out {
					if r != noRegister && m>>r&1 != 0 {
						m = regMask(1) << r
						args[0] = s.allocValToReg(psess, v.Args[0], m, true, v.Pos)

						goto ok
					}
				}

				for _, r := range dinfo[idx].in[0] {
					if r != noRegister && m>>r&1 != 0 {
						m = regMask(1) << r
						c := s.allocValToReg(psess, v.Args[0], m, true, v.Pos)
						s.copies[c] = false

						goto ok
					}
				}
				if psess.opcodeTable[v.Op].commutative {
					for _, r := range dinfo[idx].in[1] {
						if r != noRegister && m>>r&1 != 0 {
							m = regMask(1) << r
							c := s.allocValToReg(psess, v.Args[1], m, true, v.Pos)
							s.copies[c] = false
							args[0], args[1] = args[1], args[0]
							goto ok
						}
					}
				}

				if m&^desired.avoid != 0 {
					m &^= desired.avoid
				}

				c := s.allocValToReg(psess, v.Args[0], m, true, v.Pos)
				s.copies[c] = false
			}

		ok:

			if !psess.opcodeTable[v.Op].resultNotInArgs {
				s.tmpused = s.nospill
				s.nospill = 0
				s.advanceUses(v)
			}

			s.freeRegs(regspec.clobbers)
			s.tmpused |= regspec.clobbers

			{
				outRegs := [2]register{noRegister, noRegister}
				var used regMask
				for _, out := range regspec.outputs {
					mask := out.regs & s.allocatable &^ used
					if mask == 0 {
						continue
					}
					if psess.opcodeTable[v.Op].resultInArg0 && out.idx == 0 {
						if !psess.opcodeTable[v.Op].commutative {

							r := register(s.f.getHome(args[0].ID).(*Register).num)
							mask = regMask(1) << r
						} else {

							r0 := register(s.f.getHome(args[0].ID).(*Register).num)
							r1 := register(s.f.getHome(args[1].ID).(*Register).num)

							found := false
							for _, r := range dinfo[idx].out {
								if (r == r0 || r == r1) && (mask&^s.used)>>r&1 != 0 {
									mask = regMask(1) << r
									found = true
									if r == r1 {
										args[0], args[1] = args[1], args[0]
									}
									break
								}
							}
							if !found {

								mask = regMask(1) << r0
							}
						}
					}
					for _, r := range dinfo[idx].out {
						if r != noRegister && (mask&^s.used)>>r&1 != 0 {

							mask = regMask(1) << r
							break
						}
					}

					if mask&^desired.avoid != 0 {
						mask &^= desired.avoid
					}
					r := s.allocReg(psess, mask, v)
					outRegs[out.idx] = r
					used |= regMask(1) << r
					s.tmpused |= regMask(1) << r
				}

				if v.Type.IsTuple() {
					var outLocs LocPair
					if r := outRegs[0]; r != noRegister {
						outLocs[0] = &s.registers[r]
					}
					if r := outRegs[1]; r != noRegister {
						outLocs[1] = &s.registers[r]
					}
					s.f.setHome(v, outLocs)

				} else {
					if r := outRegs[0]; r != noRegister {
						s.assignReg(r, v, v)
					}
				}
			}

			if psess.opcodeTable[v.Op].resultNotInArgs {
				s.nospill = 0
				s.advanceUses(v)
			}
			s.tmpused = 0

			for i, a := range args {
				v.SetArg(i, a)
			}
			b.Values = append(b.Values, v)

		issueSpill:
		}

		if v := b.Control; v != nil && s.values[v.ID].needReg {
			if s.f.pass.debug > regDebug {
				fmt.Printf("  processing control %s\n", v.LongString(psess))
			}

			b.Control = s.allocValToReg(psess, v, s.compatRegs(psess, v.Type), false, b.Pos)
			if b.Control != v {
				v.Uses--
				b.Control.Uses++
			}

			vi := &s.values[v.ID]
			u := vi.uses
			vi.uses = u.next
			if u.next == nil {
				s.freeRegs(vi.regs)
			}
			u.next = s.freeUseRecords
			s.freeUseRecords = u
		}

		if s.f.Config.use387 {
			s.freeRegs(s.f.Config.fpRegMask)
		}

		if len(b.Succs) == 1 {
			if s.f.Config.hasGReg && s.regs[s.GReg].v != nil {
				s.freeReg(s.GReg)
			}

			top := b.Succs[0].b
			loop := s.loopnest.b2l[top.ID]
			if loop == nil || loop.header != top || loop.containsUnavoidableCall {
				goto badloop
			}

			for _, live := range s.live[b.ID] {
				if live.dist >= unlikelyDistance {

					continue
				}
				vid := live.ID
				vi := &s.values[vid]
				if vi.regs != 0 {
					continue
				}
				if vi.rematerializeable {
					continue
				}
				v := s.orig[vid]
				if s.f.Config.use387 && v.Type.IsFloat() {
					continue
				}
				m := s.compatRegs(psess, v.Type) &^ s.used
				if m&^desired.avoid != 0 {
					m &^= desired.avoid
				}
				if m != 0 {
					s.allocValToReg(psess, v, m, false, b.Pos)
				}
			}
		}
	badloop:
		;

		k := 0
		for r := register(0); r < s.numRegs; r++ {
			v := s.regs[r].v
			if v == nil {
				continue
			}
			k++
		}
		regList := make([]endReg, 0, k)
		for r := register(0); r < s.numRegs; r++ {
			v := s.regs[r].v
			if v == nil {
				continue
			}
			regList = append(regList, endReg{r, v, s.regs[r].c})
		}
		s.endRegs[b.ID] = regList

		if psess.checkEnabled {
			regValLiveSet.clear()
			for _, x := range s.live[b.ID] {
				regValLiveSet.add(x.ID)
			}
			for r := register(0); r < s.numRegs; r++ {
				v := s.regs[r].v
				if v == nil {
					continue
				}
				if !regValLiveSet.contains(v.ID) {
					s.f.Fatalf("val %s is in reg but not live at end of %s", v, b)
				}
			}
		}

		for _, e := range s.live[b.ID] {
			vi := &s.values[e.ID]
			if vi.regs != 0 {

				continue
			}
			if vi.rematerializeable {

				continue
			}

			spill := s.makeSpill(s.orig[e.ID], b)
			s.spillLive[b.ID] = append(s.spillLive[b.ID], spill.ID)
		}

		for _, e := range s.live[b.ID] {
			u := s.values[e.ID].uses
			if u == nil {
				f.Fatalf("live at end, no uses v%d", e.ID)
			}
			if u.next != nil {
				f.Fatalf("live at end, too many uses v%d", e.ID)
			}
			s.values[e.ID].uses = nil
			u.next = s.freeUseRecords
			s.freeUseRecords = u
		}
	}

	s.placeSpills()

	stacklive := psess.stackalloc(s.f, s.spillLive)

	s.shuffle(psess, stacklive)

	for {
		progress := false
		for c, used := range s.copies {
			if !used && c.Uses == 0 {
				if s.f.pass.debug > regDebug {
					fmt.Printf("delete copied value %s\n", c.LongString(psess))
				}
				c.RemoveArg(0)
				f.freeValue(psess, c)
				delete(s.copies, c)
				progress = true
			}
		}
		if !progress {
			break
		}
	}

	for _, b := range s.visitOrder {
		i := 0
		for _, v := range b.Values {
			if v.Op == OpInvalid {
				continue
			}
			b.Values[i] = v
			i++
		}
		b.Values = b.Values[:i]
	}
}

func (s *regAllocState) placeSpills() {
	f := s.f

	phiRegs := make([]regMask, f.NumBlocks())
	for _, b := range s.visitOrder {
		var m regMask
		for _, v := range b.Values {
			if v.Op != OpPhi {
				break
			}
			if r, ok := f.getHome(v.ID).(*Register); ok {
				m |= regMask(1) << uint(r.num)
			}
		}
		phiRegs[b.ID] = m
	}

	start := map[ID][]*Value{}

	after := map[ID][]*Value{}

	for i := range s.values {
		vi := s.values[i]
		spill := vi.spill
		if spill == nil {
			continue
		}
		if spill.Block != nil {

			continue
		}
		v := s.orig[i]

		best := v.Block
		bestArg := v
		var bestDepth int16
		if l := s.loopnest.b2l[best.ID]; l != nil {
			bestDepth = l.depth
		}
		b := best
		const maxSpillSearch = 100
		for i := 0; i < maxSpillSearch; i++ {

			p := b
			b = nil
			for c := s.sdom.Child(p); c != nil && i < maxSpillSearch; c, i = s.sdom.Sibling(c), i+1 {
				if s.sdom[c.ID].entry <= vi.restoreMin && s.sdom[c.ID].exit >= vi.restoreMax {

					b = c
					break
				}
			}
			if b == nil {

				break
			}

			var depth int16
			if l := s.loopnest.b2l[b.ID]; l != nil {
				depth = l.depth
			}
			if depth > bestDepth {

				continue
			}

			if len(b.Preds) == 1 {
				for _, e := range s.endRegs[b.Preds[0].b.ID] {
					if e.v == v {

						best = b
						bestArg = e.c
						bestDepth = depth
						break
					}
				}
			} else {
				for _, e := range s.startRegs[b.ID] {
					if e.v == v {

						best = b
						bestArg = e.c
						bestDepth = depth
						break
					}
				}
			}
		}

		spill.Block = best
		spill.AddArg(bestArg)
		if best == v.Block && v.Op != OpPhi {

			after[v.ID] = append(after[v.ID], spill)
		} else {

			start[best.ID] = append(start[best.ID], spill)
		}
	}

	// Insert spill instructions into the block schedules.
	var oldSched []*Value
	for _, b := range s.visitOrder {
		nphi := 0
		for _, v := range b.Values {
			if v.Op != OpPhi {
				break
			}
			nphi++
		}
		oldSched = append(oldSched[:0], b.Values[nphi:]...)
		b.Values = b.Values[:nphi]
		b.Values = append(b.Values, start[b.ID]...)
		for _, v := range oldSched {
			b.Values = append(b.Values, v)
			b.Values = append(b.Values, after[v.ID]...)
		}
	}
}

// shuffle fixes up all the merge edges (those going into blocks of indegree > 1).
func (s *regAllocState) shuffle(psess *PackageSession, stacklive [][]ID) {
	var e edgeState
	e.s = s
	e.cache = map[ID][]*Value{}
	e.contents = map[Location]contentRecord{}
	if s.f.pass.debug > regDebug {
		fmt.Printf("shuffle %s\n", s.f.Name)
		fmt.Println(s.f.String(psess))
	}

	for _, b := range s.visitOrder {
		if len(b.Preds) <= 1 {
			continue
		}
		e.b = b
		for i, edge := range b.Preds {
			p := edge.b
			e.p = p
			e.setup(psess, i, s.endRegs[p.ID], s.startRegs[b.ID], stacklive[p.ID])
			e.process(psess)
		}
	}
}

type edgeState struct {
	s    *regAllocState
	p, b *Block // edge goes from p->b.

	// for each pre-regalloc value, a list of equivalent cached values
	cache      map[ID][]*Value
	cachedVals []ID // (superset of) keys of the above map, for deterministic iteration

	// map from location to the value it contains
	contents map[Location]contentRecord

	// desired destination locations
	destinations []dstRecord
	extra        []dstRecord

	usedRegs              regMask // registers currently holding something
	uniqueRegs            regMask // registers holding the only copy of a value
	finalRegs             regMask // registers holding final target
	rematerializeableRegs regMask // registers that hold rematerializeable values
}

type contentRecord struct {
	vid   ID       // pre-regalloc value
	c     *Value   // cached value
	final bool     // this is a satisfied destination
	pos   src.XPos // source position of use of the value
}

type dstRecord struct {
	loc    Location // register or stack slot
	vid    ID       // pre-regalloc value it should contain
	splice **Value  // place to store reference to the generating instruction
	pos    src.XPos // source position of use of this location
}

// setup initializes the edge state for shuffling.
func (e *edgeState) setup(psess *PackageSession, idx int, srcReg []endReg, dstReg []startReg, stacklive []ID) {
	if e.s.f.pass.debug > regDebug {
		fmt.Printf("edge %s->%s\n", e.p, e.b)
	}

	for _, vid := range e.cachedVals {
		delete(e.cache, vid)
	}
	e.cachedVals = e.cachedVals[:0]
	for k := range e.contents {
		delete(e.contents, k)
	}
	e.usedRegs = 0
	e.uniqueRegs = 0
	e.finalRegs = 0
	e.rematerializeableRegs = 0

	for _, x := range srcReg {
		e.set(psess, &e.s.registers[x.r], x.v.ID, x.c, false, psess.src.NoXPos)
	}

	for _, spillID := range stacklive {
		v := e.s.orig[spillID]
		spill := e.s.values[v.ID].spill
		if !e.s.sdom.isAncestorEq(spill.Block, e.p) {

			continue
		}
		e.set(psess, e.s.f.getHome(spillID), v.ID, spill, false, psess.src.NoXPos)
	}

	dsts := e.destinations[:0]
	for _, x := range dstReg {
		dsts = append(dsts, dstRecord{&e.s.registers[x.r], x.v.ID, nil, x.pos})
	}

	for _, v := range e.b.Values {
		if v.Op != OpPhi {
			break
		}
		loc := e.s.f.getHome(v.ID)
		if loc == nil {
			continue
		}
		dsts = append(dsts, dstRecord{loc, v.Args[idx].ID, &v.Args[idx], v.Pos})
	}
	e.destinations = dsts

	if e.s.f.pass.debug > regDebug {
		for _, vid := range e.cachedVals {
			a := e.cache[vid]
			for _, c := range a {
				fmt.Printf("src %s: v%d cache=%s\n", e.s.f.getHome(c.ID), vid, c)
			}
		}
		for _, d := range e.destinations {
			fmt.Printf("dst %s: v%d\n", d.loc, d.vid)
		}
	}
}

// process generates code to move all the values to the right destination locations.
func (e *edgeState) process(psess *PackageSession) {
	dsts := e.destinations

	for len(dsts) > 0 {
		i := 0
		for _, d := range dsts {
			if !e.processDest(psess, d.loc, d.vid, d.splice, d.pos) {

				dsts[i] = d
				i++
			}
		}
		if i < len(dsts) {

			dsts = dsts[:i]

			dsts = append(dsts, e.extra...)
			e.extra = e.extra[:0]
			continue
		}

		d := dsts[0]
		loc := d.loc
		vid := e.contents[loc].vid
		c := e.contents[loc].c
		r := e.findRegFor(psess, c.Type)
		if e.s.f.pass.debug > regDebug {
			fmt.Printf("breaking cycle with v%d in %s:%s\n", vid, loc, c)
		}
		e.erase(r)
		pos := d.pos.WithNotStmt()
		if _, isReg := loc.(*Register); isReg {
			c = e.p.NewValue1(pos, OpCopy, c.Type, c)
		} else {
			c = e.p.NewValue1(pos, OpLoadReg, c.Type, c)
		}
		e.set(psess, r, vid, c, false, pos)
		if c.Op == OpLoadReg && e.s.isGReg(register(r.(*Register).num)) {
			e.s.f.Fatalf("process.OpLoadReg targeting g: " + c.LongString(psess))
		}
	}
}

// processDest generates code to put value vid into location loc. Returns true
// if progress was made.
func (e *edgeState) processDest(psess *PackageSession, loc Location, vid ID, splice **Value, pos src.XPos) bool {
	pos = pos.WithNotStmt()
	occupant := e.contents[loc]
	if occupant.vid == vid {

		e.contents[loc] = contentRecord{vid, occupant.c, true, pos}
		if splice != nil {
			(*splice).Uses--
			*splice = occupant.c
			occupant.c.Uses++
		}

		if _, ok := e.s.copies[occupant.c]; ok {

			e.s.copies[occupant.c] = true
		}
		return true
	}

	if len(e.cache[occupant.vid]) == 1 && !e.s.values[occupant.vid].rematerializeable {

		return false
	}

	v := e.s.orig[vid]
	var c *Value
	var src Location
	if e.s.f.pass.debug > regDebug {
		fmt.Printf("moving v%d to %s\n", vid, loc)
		fmt.Printf("sources of v%d:", vid)
	}
	for _, w := range e.cache[vid] {
		h := e.s.f.getHome(w.ID)
		if e.s.f.pass.debug > regDebug {
			fmt.Printf(" %s:%s", h, w)
		}
		_, isreg := h.(*Register)
		if src == nil || isreg {
			c = w
			src = h
		}
	}
	if e.s.f.pass.debug > regDebug {
		if src != nil {
			fmt.Printf(" [use %s]\n", src)
		} else {
			fmt.Printf(" [no source]\n")
		}
	}
	_, dstReg := loc.(*Register)

	e.erase(loc)
	var x *Value
	if c == nil || e.s.values[vid].rematerializeable {
		if !e.s.values[vid].rematerializeable {
			e.s.f.Fatalf("can't find source for %s->%s: %s\n", e.p, e.b, v.LongString(psess))
		}
		if dstReg {
			x = v.copyInto(psess, e.p)
		} else {

			r := e.findRegFor(psess, v.Type)
			e.erase(r)
			x = v.copyIntoWithXPos(psess, e.p, pos)
			e.set(psess, r, vid, x, false, pos)

			x = e.p.NewValue1(pos, OpStoreReg, loc.(LocalSlot).Type, x)
		}
	} else {

		_, srcReg := src.(*Register)
		if srcReg {
			if dstReg {
				x = e.p.NewValue1(pos, OpCopy, c.Type, c)
			} else {
				x = e.p.NewValue1(pos, OpStoreReg, loc.(LocalSlot).Type, c)
			}
		} else {
			if dstReg {
				x = e.p.NewValue1(pos, OpLoadReg, c.Type, c)
			} else {

				r := e.findRegFor(psess, c.Type)
				e.erase(r)
				t := e.p.NewValue1(pos, OpLoadReg, c.Type, c)
				e.set(psess, r, vid, t, false, pos)
				x = e.p.NewValue1(pos, OpStoreReg, loc.(LocalSlot).Type, t)
			}
		}
	}
	e.set(psess, loc, vid, x, true, pos)
	if x.Op == OpLoadReg && e.s.isGReg(register(loc.(*Register).num)) {
		e.s.f.Fatalf("processDest.OpLoadReg targeting g: " + x.LongString(psess))
	}
	if splice != nil {
		(*splice).Uses--
		*splice = x
		x.Uses++
	}
	return true
}

// set changes the contents of location loc to hold the given value and its cached representative.
func (e *edgeState) set(psess *PackageSession, loc Location, vid ID, c *Value, final bool, pos src.XPos) {
	e.s.f.setHome(c, loc)
	e.contents[loc] = contentRecord{vid, c, final, pos}
	a := e.cache[vid]
	if len(a) == 0 {
		e.cachedVals = append(e.cachedVals, vid)
	}
	a = append(a, c)
	e.cache[vid] = a
	if r, ok := loc.(*Register); ok {
		e.usedRegs |= regMask(1) << uint(r.num)
		if final {
			e.finalRegs |= regMask(1) << uint(r.num)
		}
		if len(a) == 1 {
			e.uniqueRegs |= regMask(1) << uint(r.num)
		}
		if len(a) == 2 {
			if t, ok := e.s.f.getHome(a[0].ID).(*Register); ok {
				e.uniqueRegs &^= regMask(1) << uint(t.num)
			}
		}
		if e.s.values[vid].rematerializeable {
			e.rematerializeableRegs |= regMask(1) << uint(r.num)
		}
	}
	if e.s.f.pass.debug > regDebug {
		fmt.Printf("%s\n", c.LongString(psess))
		fmt.Printf("v%d now available in %s:%s\n", vid, loc, c)
	}
}

// erase removes any user of loc.
func (e *edgeState) erase(loc Location) {
	cr := e.contents[loc]
	if cr.c == nil {
		return
	}
	vid := cr.vid

	if cr.final {

		e.extra = append(e.extra, dstRecord{loc, cr.vid, nil, cr.pos})
	}

	a := e.cache[vid]
	for i, c := range a {
		if e.s.f.getHome(c.ID) == loc {
			if e.s.f.pass.debug > regDebug {
				fmt.Printf("v%d no longer available in %s:%s\n", vid, loc, c)
			}
			a[i], a = a[len(a)-1], a[:len(a)-1]
			break
		}
	}
	e.cache[vid] = a

	if r, ok := loc.(*Register); ok {
		e.usedRegs &^= regMask(1) << uint(r.num)
		if cr.final {
			e.finalRegs &^= regMask(1) << uint(r.num)
		}
		e.rematerializeableRegs &^= regMask(1) << uint(r.num)
	}
	if len(a) == 1 {
		if r, ok := e.s.f.getHome(a[0].ID).(*Register); ok {
			e.uniqueRegs |= regMask(1) << uint(r.num)
		}
	}
}

// findRegFor finds a register we can use to make a temp copy of type typ.
func (e *edgeState) findRegFor(psess *PackageSession, typ *types.Type) Location {
	// Which registers are possibilities.
	var m regMask
	types := &e.s.f.Config.Types
	if typ.IsFloat() {
		m = e.s.compatRegs(psess, types.Float64)
	} else {
		m = e.s.compatRegs(psess, types.Int64)
	}

	x := m &^ e.usedRegs
	if x != 0 {
		return &e.s.registers[pickReg(x)]
	}
	x = m &^ e.uniqueRegs &^ e.finalRegs
	if x != 0 {
		return &e.s.registers[pickReg(x)]
	}
	x = m &^ e.uniqueRegs
	if x != 0 {
		return &e.s.registers[pickReg(x)]
	}
	x = m & e.rematerializeableRegs
	if x != 0 {
		return &e.s.registers[pickReg(x)]
	}

	for _, vid := range e.cachedVals {
		a := e.cache[vid]
		for _, c := range a {
			if r, ok := e.s.f.getHome(c.ID).(*Register); ok && m>>uint(r.num)&1 != 0 {
				if !c.rematerializeable(psess) {
					x := e.p.NewValue1(c.Pos, OpStoreReg, c.Type, c)

					t := LocalSlot{N: e.s.f.fe.Auto(c.Pos, types.Int64), Type: types.Int64}

					e.set(psess, t, vid, x, false, c.Pos)
					if e.s.f.pass.debug > regDebug {
						fmt.Printf("  SPILL %s->%s %s\n", r, t, x.LongString(psess))
					}
				}

				return r
			}
		}
	}

	fmt.Printf("m:%d unique:%d final:%d rematerializable:%d\n", m, e.uniqueRegs, e.finalRegs, e.rematerializeableRegs)
	for _, vid := range e.cachedVals {
		a := e.cache[vid]
		for _, c := range a {
			fmt.Printf("v%d: %s %s\n", vid, c, e.s.f.getHome(c.ID))
		}
	}
	e.s.f.Fatalf("can't find empty register on edge %s->%s", e.p, e.b)
	return nil
}

// rematerializeable reports whether the register allocator should recompute
// a value instead of spilling/restoring it.
func (v *Value) rematerializeable(psess *PackageSession) bool {
	if !psess.opcodeTable[v.Op].rematerializeable {
		return false
	}
	for _, a := range v.Args {

		if a.Op != OpSP && a.Op != OpSB {
			return false
		}
	}
	return true
}

type liveInfo struct {
	ID   ID       // ID of value
	dist int32    // # of instructions before next use
	pos  src.XPos // source position of next use
}

// computeLive computes a map from block ID to a list of value IDs live at the end
// of that block. Together with the value ID is a count of how many instructions
// to the next use of that value. The resulting map is stored in s.live.
// computeLive also computes the desired register information at the end of each block.
// This desired register information is stored in s.desired.
// TODO: this could be quadratic if lots of variables are live across lots of
// basic blocks. Figure out a way to make this function (or, more precisely, the user
// of this function) require only linear size & time.
func (s *regAllocState) computeLive(psess *PackageSession) {
	f := s.f
	s.live = make([][]liveInfo, f.NumBlocks())
	s.desired = make([]desiredState, f.NumBlocks())
	var phis []*Value

	live := f.newSparseMap(f.NumValues())
	defer f.retSparseMap(live)
	t := f.newSparseMap(f.NumValues())
	defer f.retSparseMap(t)

	// Keep track of which value we want in each register.
	var desired desiredState

	po := f.postorder()
	s.loopnest = f.loopnest(psess)
	s.loopnest.calculateDepths()
	for {
		changed := false

		for _, b := range po {

			live.clear()
			for _, e := range s.live[b.ID] {
				live.set(e.ID, e.dist+int32(len(b.Values)), e.pos)
			}

			if b.Control != nil && s.values[b.Control.ID].needReg {
				live.set(b.Control.ID, int32(len(b.Values)), b.Pos)
			}

			phis = phis[:0]
			for i := len(b.Values) - 1; i >= 0; i-- {
				v := b.Values[i]
				live.remove(v.ID)
				if v.Op == OpPhi {

					phis = append(phis, v)
					continue
				}
				if psess.opcodeTable[v.Op].call {
					c := live.contents()
					for i := range c {
						c[i].val += unlikelyDistance
					}
				}
				for _, a := range v.Args {
					if s.values[a.ID].needReg {
						live.set(a.ID, int32(i), v.Pos)
					}
				}
			}

			desired.copy(&s.desired[b.ID])
			for i := len(b.Values) - 1; i >= 0; i-- {
				v := b.Values[i]
				prefs := desired.remove(v.ID)
				if v.Op == OpPhi {

					continue
				}
				regspec := s.regspec(psess, v.Op)

				desired.clobber(regspec.clobbers)

				for _, j := range regspec.inputs {
					if countRegs(j.regs) != 1 {
						continue
					}
					desired.clobber(j.regs)
					desired.add(v.Args[j.idx].ID, pickReg(j.regs))
				}

				if psess.opcodeTable[v.Op].resultInArg0 {
					if psess.opcodeTable[v.Op].commutative {
						desired.addList(v.Args[1].ID, prefs)
					}
					desired.addList(v.Args[0].ID, prefs)
				}
			}

			for i, e := range b.Preds {
				p := e.b

				delta := int32(normalDistance)
				if len(p.Succs) == 2 {
					if p.Succs[0].b == b && p.Likely == BranchLikely ||
						p.Succs[1].b == b && p.Likely == BranchUnlikely {
						delta = likelyDistance
					}
					if p.Succs[0].b == b && p.Likely == BranchUnlikely ||
						p.Succs[1].b == b && p.Likely == BranchLikely {
						delta = unlikelyDistance
					}
				}

				s.desired[p.ID].merge(&desired)

				t.clear()
				for _, e := range s.live[p.ID] {
					t.set(e.ID, e.dist, e.pos)
				}
				update := false

				for _, e := range live.contents() {
					d := e.val + delta
					if !t.contains(e.key) || d < t.get(e.key) {
						update = true
						t.set(e.key, d, e.aux)
					}
				}

				for _, v := range phis {
					id := v.Args[i].ID
					if s.values[id].needReg && (!t.contains(id) || delta < t.get(id)) {
						update = true
						t.set(id, delta, v.Pos)
					}
				}

				if !update {
					continue
				}

				l := s.live[p.ID][:0]
				if cap(l) < t.size() {
					l = make([]liveInfo, 0, t.size())
				}
				for _, e := range t.contents() {
					l = append(l, liveInfo{e.key, e.val, e.aux})
				}
				s.live[p.ID] = l
				changed = true
			}
		}

		if !changed {
			break
		}
	}
	if f.pass.debug > regDebug {
		fmt.Println("live values at end of each block")
		for _, b := range f.Blocks {
			fmt.Printf("  %s:", b)
			for _, x := range s.live[b.ID] {
				fmt.Printf(" v%d", x.ID)
				for _, e := range s.desired[b.ID].entries {
					if e.ID != x.ID {
						continue
					}
					fmt.Printf("[")
					first := true
					for _, r := range e.regs {
						if r == noRegister {
							continue
						}
						if !first {
							fmt.Printf(",")
						}
						fmt.Print(&s.registers[r])
						first = false
					}
					fmt.Printf("]")
				}
			}
			if avoid := s.desired[b.ID].avoid; avoid != 0 {
				fmt.Printf(" avoid=%v", s.RegMaskString(avoid))
			}
			fmt.Println()
		}
	}
}

// A desiredState represents desired register assignments.
type desiredState struct {
	// Desired assignments will be small, so we just use a list
	// of valueID+registers entries.
	entries []desiredStateEntry
	// Registers that other values want to be in.  This value will
	// contain at least the union of the regs fields of entries, but
	// may contain additional entries for values that were once in
	// this data structure but are no longer.
	avoid regMask
}
type desiredStateEntry struct {
	// (pre-regalloc) value
	ID ID
	// Registers it would like to be in, in priority order.
	// Unused slots are filled with noRegister.
	regs [4]register
}

func (d *desiredState) clear() {
	d.entries = d.entries[:0]
	d.avoid = 0
}

// get returns a list of desired registers for value vid.
func (d *desiredState) get(vid ID) [4]register {
	for _, e := range d.entries {
		if e.ID == vid {
			return e.regs
		}
	}
	return [4]register{noRegister, noRegister, noRegister, noRegister}
}

// add records that we'd like value vid to be in register r.
func (d *desiredState) add(vid ID, r register) {
	d.avoid |= regMask(1) << r
	for i := range d.entries {
		e := &d.entries[i]
		if e.ID != vid {
			continue
		}
		if e.regs[0] == r {

			return
		}
		for j := 1; j < len(e.regs); j++ {
			if e.regs[j] == r {

				copy(e.regs[1:], e.regs[:j])
				e.regs[0] = r
				return
			}
		}
		copy(e.regs[1:], e.regs[:])
		e.regs[0] = r
		return
	}
	d.entries = append(d.entries, desiredStateEntry{vid, [4]register{r, noRegister, noRegister, noRegister}})
}

func (d *desiredState) addList(vid ID, regs [4]register) {

	for i := len(regs) - 1; i >= 0; i-- {
		r := regs[i]
		if r != noRegister {
			d.add(vid, r)
		}
	}
}

// clobber erases any desired registers in the set m.
func (d *desiredState) clobber(m regMask) {
	for i := 0; i < len(d.entries); {
		e := &d.entries[i]
		j := 0
		for _, r := range e.regs {
			if r != noRegister && m>>r&1 == 0 {
				e.regs[j] = r
				j++
			}
		}
		if j == 0 {

			d.entries[i] = d.entries[len(d.entries)-1]
			d.entries = d.entries[:len(d.entries)-1]
			continue
		}
		for ; j < len(e.regs); j++ {
			e.regs[j] = noRegister
		}
		i++
	}
	d.avoid &^= m
}

// copy copies a desired state from another desiredState x.
func (d *desiredState) copy(x *desiredState) {
	d.entries = append(d.entries[:0], x.entries...)
	d.avoid = x.avoid
}

// remove removes the desired registers for vid and returns them.
func (d *desiredState) remove(vid ID) [4]register {
	for i := range d.entries {
		if d.entries[i].ID == vid {
			regs := d.entries[i].regs
			d.entries[i] = d.entries[len(d.entries)-1]
			d.entries = d.entries[:len(d.entries)-1]
			return regs
		}
	}
	return [4]register{noRegister, noRegister, noRegister, noRegister}
}

// merge merges another desired state x into d.
func (d *desiredState) merge(x *desiredState) {
	d.avoid |= x.avoid

	for _, e := range x.entries {
		d.addList(e.ID, e.regs)
	}
}

func min32(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}
func max32(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

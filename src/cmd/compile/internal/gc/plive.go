package gc

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"

	"os"
	"strings"
)

// BlockEffects summarizes the liveness effects on an SSA block.
type BlockEffects struct {
	// Computed during Liveness.prologue using only the content of
	// individual blocks:
	//
	//	uevar: upward exposed variables (used before set in block)
	//	varkill: killed variables (set in block)
	//	avarinit: addrtaken variables set or used (proof of initialization)
	uevar    varRegVec
	varkill  varRegVec
	avarinit bvec

	// Computed during Liveness.solve using control flow information:
	//
	//	livein: variables live at block entry
	//	liveout: variables live at block exit
	//	avarinitany: addrtaken variables possibly initialized at block exit
	//		(initialized in block or at exit from any predecessor block)
	//	avarinitall: addrtaken variables certainly initialized at block exit
	//		(initialized in block or at exit from all predecessor blocks)
	livein      varRegVec
	liveout     varRegVec
	avarinitany bvec
	avarinitall bvec
}

// A collection of global state used by liveness analysis.
type Liveness struct {
	fn         *Node
	f          *ssa.Func
	vars       []*Node
	idx        map[*Node]int32
	stkptrsize int64

	be []BlockEffects

	// unsafePoints bit i is set if Value ID i is not a safe point.
	unsafePoints bvec

	// An array with a bit vector for each safe point in the
	// current Block during Liveness.epilogue. Indexed in Value
	// order for that block. Additionally, for the entry block
	// livevars[0] is the entry bitmap. Liveness.compact moves
	// these to stackMaps and regMaps.
	livevars []varRegVec

	// livenessMap maps from safe points (i.e., CALLs) to their
	// liveness map indexes.
	livenessMap LivenessMap
	stackMapSet bvecSet
	stackMaps   []bvec
	regMapSet   map[liveRegMask]int
	regMaps     []liveRegMask

	cache progeffectscache
}

// LivenessMap maps from *ssa.Value to LivenessIndex.
type LivenessMap struct {
	m []LivenessIndex
}

func (m *LivenessMap) reset(psess *PackageSession, ids int) {
	m2 := m.m
	if ids > cap(m2) {
		m2 = make([]LivenessIndex, ids)
	} else {
		m2 = m2[:ids]
	}
	none := psess.LivenessInvalid
	for i := range m2 {
		m2[i] = none
	}
	m.m = m2
}

func (m *LivenessMap) set(v *ssa.Value, i LivenessIndex) {
	m.m[v.ID] = i
}

func (m LivenessMap) Get(psess *PackageSession, v *ssa.Value) LivenessIndex {
	if int(v.ID) < len(m.m) {
		return m.m[int(v.ID)]
	}

	return psess.LivenessInvalid
}

// LivenessIndex stores the liveness map index for a safe-point.
type LivenessIndex struct {
	stackMapIndex int
	regMapIndex   int
}

// LivenessInvalid indicates an unsafe point.
//
// We use index -2 because PCDATA tables conventionally start at -1,
// so -1 is used to mean the entry liveness map (which is actually at
// index 0; sigh). TODO(austin): Maybe we should use PCDATA+1 as the
// index into the liveness map so -1 uniquely refers to the entry
// liveness map.

func (idx LivenessIndex) Valid() bool {
	return idx.stackMapIndex >= 0
}

type progeffectscache struct {
	textavarinit []int32
	retuevar     []int32
	tailuevar    []int32
	initialized  bool
}

// varRegVec contains liveness bitmaps for variables and registers.
type varRegVec struct {
	vars bvec
	regs liveRegMask
}

func (v *varRegVec) Eq(psess *PackageSession, v2 varRegVec) bool {
	return v.vars.Eq(psess, v2.vars) && v.regs == v2.regs
}

func (v *varRegVec) Copy(v2 varRegVec) {
	v.vars.Copy(v2.vars)
	v.regs = v2.regs
}

func (v *varRegVec) Clear() {
	v.vars.Clear()
	v.regs = 0
}

func (v *varRegVec) Or(v1, v2 varRegVec) {
	v.vars.Or(v1.vars, v2.vars)
	v.regs = v1.regs | v2.regs
}

func (v *varRegVec) AndNot(v1, v2 varRegVec) {
	v.vars.AndNot(v1.vars, v2.vars)
	v.regs = v1.regs &^ v2.regs
}

// livenessShouldTrack reports whether the liveness analysis
// should track the variable n.
// We don't care about variables that have no pointers,
// nor do we care about non-local variables,
// nor do we care about empty structs (handled by the pointer check),
// nor do we care about the fake PAUTOHEAP variables.
func (psess *PackageSession) livenessShouldTrack(n *Node) bool {
	return n.Op == ONAME && (n.Class() == PAUTO || n.Class() == PPARAM || n.Class() == PPARAMOUT) && psess.types.Haspointers(n.Type)
}

// getvariables returns the list of on-stack variables that we need to track
// and a map for looking up indices by *Node.
func (psess *PackageSession) getvariables(fn *Node) ([]*Node, map[*Node]int32) {
	var vars []*Node
	for _, n := range fn.Func.Dcl {
		if psess.livenessShouldTrack(n) {
			vars = append(vars, n)
		}
	}
	idx := make(map[*Node]int32, len(vars))
	for i, n := range vars {
		idx[n] = int32(i)
	}
	return vars, idx
}

func (lv *Liveness) initcache(psess *PackageSession) {
	if lv.cache.initialized {
		psess.
			Fatalf("liveness cache initialized twice")
		return
	}
	lv.cache.initialized = true

	for i, node := range lv.vars {
		switch node.Class() {
		case PPARAM:

			lv.cache.tailuevar = append(lv.cache.tailuevar, int32(i))

			if node.Addrtaken() {
				lv.cache.textavarinit = append(lv.cache.textavarinit, int32(i))
			}

		case PPARAMOUT:

			if !node.Addrtaken() {
				lv.cache.retuevar = append(lv.cache.retuevar, int32(i))
			}
		}
	}
}

// A liveEffect is a set of flags that describe an instruction's
// liveness effects on a variable.
//
// The possible flags are:
//	uevar - used by the instruction
//	varkill - killed by the instruction
//		for variables without address taken, means variable was set
//		for variables with address taken, means variable was marked dead
//	avarinit - initialized or referred to by the instruction,
//		only for variables with address taken but not escaping to heap
//
// The avarinit output serves as a signal that the data has been
// initialized, because any use of a variable must come after its
// initialization.
type liveEffect int

const (
	uevar liveEffect = 1 << iota
	varkill
	avarinit
)

// valueEffects returns the index of a variable in lv.vars and the
// liveness effects v has on that variable.
// If v does not affect any tracked variables, it returns -1, 0.
func (lv *Liveness) valueEffects(psess *PackageSession, v *ssa.Value) (int32, liveEffect) {
	n, e := psess.affectedNode(v)
	if e == 0 || n == nil || n.Op != ONAME {
		return -1, 0
	}

	switch v.Op {
	case ssa.OpVarDef, ssa.OpVarKill, ssa.OpVarLive, ssa.OpKeepAlive:
		if !n.Name.Used() {
			return -1, 0
		}
	}

	var effect liveEffect
	if n.Addrtaken() {
		if v.Op != ssa.OpVarKill {
			effect |= avarinit
		}
		if v.Op == ssa.OpVarDef || v.Op == ssa.OpVarKill {
			effect |= varkill
		}
	} else {

		if e&ssa.SymRead != 0 || e&(ssa.SymAddr|ssa.SymWrite) == ssa.SymAddr {
			effect |= uevar
		}
		if e&ssa.SymWrite != 0 && (!isfat(n.Type) || v.Op == ssa.OpVarDef) {
			effect |= varkill
		}
	}

	if effect == 0 {
		return -1, 0
	}

	if pos, ok := lv.idx[n]; ok {
		return pos, effect
	}
	return -1, 0
}

// affectedNode returns the *Node affected by v
func (psess *PackageSession) affectedNode(v *ssa.Value) (*Node, ssa.SymEffect) {

	switch v.Op {
	case ssa.OpLoadReg:
		n, _ := psess.AutoVar(v.Args[0])
		return n, ssa.SymRead
	case ssa.OpStoreReg:
		n, _ := psess.AutoVar(v)
		return n, ssa.SymWrite

	case ssa.OpVarLive:
		return v.Aux.(*Node), ssa.SymRead
	case ssa.OpVarDef, ssa.OpVarKill:
		return v.Aux.(*Node), ssa.SymWrite
	case ssa.OpKeepAlive:
		n, _ := psess.AutoVar(v.Args[0])
		return n, ssa.SymRead
	}

	e := v.Op.SymEffect(psess.ssa)
	if e == 0 {
		return nil, 0
	}

	switch a := v.Aux.(type) {
	case nil, *obj.LSym:

		return nil, e
	case *Node:
		return a, e
	default:
		psess.
			Fatalf("weird aux: %s", v.LongString(psess.ssa))
		return nil, e
	}
}

// regEffects returns the registers affected by v.
func (lv *Liveness) regEffects(psess *PackageSession, v *ssa.Value) (uevar, kill liveRegMask) {
	if v.Op == ssa.OpPhi {

		return 0, 0
	}
	addLocs := func(mask liveRegMask, v *ssa.Value, ptrOnly bool) liveRegMask {
		if int(v.ID) >= len(lv.f.RegAlloc) {

			return mask
		}
		loc := lv.f.RegAlloc[v.ID]
		if loc == nil {

			return mask
		}
		if v.Op == ssa.OpGetG {

			return mask
		}

		// Collect registers and types from v's location.
		var regs [2]*ssa.Register
		nreg := 0
		switch loc := loc.(type) {
		case ssa.LocalSlot:
			return mask
		case *ssa.Register:
			if ptrOnly && !v.Type.HasHeapPointer(psess.types) {
				return mask
			}
			regs[0] = loc
			nreg = 1
		case ssa.LocPair:

			if v.Type.Etype != types.TTUPLE {
				v.Fatalf("location pair %s has non-tuple type %v", loc, v.Type)
			}
			for i, loc1 := range loc {
				if loc1 == nil {
					continue
				}
				if ptrOnly && !v.Type.FieldType(psess.types, i).HasHeapPointer(psess.types) {
					continue
				}
				regs[nreg] = loc1.(*ssa.Register)
				nreg++
			}
		default:
			v.Fatalf("weird RegAlloc location: %s (%T)", loc, loc)
		}

		for _, reg := range regs[:nreg] {
			if reg.GCNum() == -1 {
				if ptrOnly {
					v.Fatalf("pointer in non-pointer register %v", reg)
				} else {
					continue
				}
			}
			mask |= 1 << uint(reg.GCNum())
		}
		return mask
	}

	kill = addLocs(0, v, false)
	for _, arg := range v.Args {

		uevar = addLocs(uevar, arg, true)
	}
	return uevar, kill
}

type liveRegMask uint32

func (m liveRegMask) niceString(config *ssa.Config) string {
	if m == 0 {
		return "<none>"
	}
	str := ""
	for i, reg := range config.GCRegMap {
		if m&(1<<uint(i)) != 0 {
			if str != "" {
				str += ","
			}
			str += reg.String()
		}
	}
	return str
}

type livenessFuncCache struct {
	be          []BlockEffects
	livenessMap LivenessMap
}

// Constructs a new liveness structure used to hold the global state of the
// liveness computation. The cfg argument is a slice of *BasicBlocks and the
// vars argument is a slice of *Nodes.
func (psess *PackageSession) newliveness(fn *Node, f *ssa.Func, vars []*Node, idx map[*Node]int32, stkptrsize int64) *Liveness {
	lv := &Liveness{
		fn:         fn,
		f:          f,
		vars:       vars,
		idx:        idx,
		stkptrsize: stkptrsize,

		regMapSet: make(map[liveRegMask]int),
	}

	if lc, _ := f.Cache.Liveness.(*livenessFuncCache); lc == nil {

		f.Cache.Liveness = new(livenessFuncCache)
	} else {
		if cap(lc.be) >= f.NumBlocks() {
			lv.be = lc.be[:f.NumBlocks()]
		}
		lv.livenessMap = LivenessMap{lc.livenessMap.m[:0]}
	}
	if lv.be == nil {
		lv.be = make([]BlockEffects, f.NumBlocks())
	}

	nblocks := int32(len(f.Blocks))
	nvars := int32(len(vars))
	bulk := psess.bvbulkalloc(nvars, nblocks*7)
	for _, b := range f.Blocks {
		be := lv.blockEffects(b)

		be.uevar = varRegVec{vars: bulk.next()}
		be.varkill = varRegVec{vars: bulk.next()}
		be.livein = varRegVec{vars: bulk.next()}
		be.liveout = varRegVec{vars: bulk.next()}
		be.avarinit = bulk.next()
		be.avarinitany = bulk.next()
		be.avarinitall = bulk.next()
	}
	lv.livenessMap.reset(psess, lv.f.NumValues())

	lv.markUnsafePoints(psess)
	return lv
}

func (lv *Liveness) blockEffects(b *ssa.Block) *BlockEffects {
	return &lv.be[b.ID]
}

// NOTE: The bitmap for a specific type t could be cached in t after
// the first run and then simply copied into bv at the correct offset
// on future calls with the same type t.
func (psess *PackageSession) onebitwalktype1(t *types.Type, off int64, bv bvec) {
	if t.Align > 0 && off&int64(t.Align-1) != 0 {
		psess.
			Fatalf("onebitwalktype1: invalid initial alignment: type %v has alignment %d, but offset is %v", t, t.Align, off)
	}

	switch t.Etype {
	case TINT8, TUINT8, TINT16, TUINT16,
		TINT32, TUINT32, TINT64, TUINT64,
		TINT, TUINT, TUINTPTR, TBOOL,
		TFLOAT32, TFLOAT64, TCOMPLEX64, TCOMPLEX128:

	case TPTR32, TPTR64, TUNSAFEPTR, TFUNC, TCHAN, TMAP:
		if off&int64(psess.Widthptr-1) != 0 {
			psess.
				Fatalf("onebitwalktype1: invalid alignment, %v", t)
		}
		bv.Set(psess, int32(off/int64(psess.Widthptr)))

	case TSTRING:

		if off&int64(psess.Widthptr-1) != 0 {
			psess.
				Fatalf("onebitwalktype1: invalid alignment, %v", t)
		}
		bv.Set(psess, int32(off/int64(psess.Widthptr)))

	case TINTER:

		if off&int64(psess.Widthptr-1) != 0 {
			psess.
				Fatalf("onebitwalktype1: invalid alignment, %v", t)
		}

		bv.Set(psess, int32(off/int64(psess.Widthptr)+1))

	case TSLICE:

		if off&int64(psess.Widthptr-1) != 0 {
			psess.
				Fatalf("onebitwalktype1: invalid TARRAY alignment, %v", t)
		}
		bv.Set(psess, int32(off/int64(psess.Widthptr)))

	case TARRAY:
		elt := t.Elem(psess.types)
		if elt.Width == 0 {

			break
		}
		for i := int64(0); i < t.NumElem(psess.types); i++ {
			psess.
				onebitwalktype1(elt, off, bv)
			off += elt.Width
		}

	case TSTRUCT:
		for _, f := range t.Fields(psess.types).Slice() {
			psess.
				onebitwalktype1(f.Type, off+f.Offset, bv)
		}

	default:
		psess.
			Fatalf("onebitwalktype1: unexpected type, %v", t)
	}
}

// usedRegs returns the maximum width of the live register map.
func (lv *Liveness) usedRegs() int32 {
	var any liveRegMask
	for _, live := range lv.regMaps {
		any |= live
	}
	i := int32(0)
	for any != 0 {
		any >>= 1
		i++
	}
	return i
}

// Generates live pointer value maps for arguments and local variables. The
// this argument and the in arguments are always assumed live. The vars
// argument is a slice of *Nodes.
func (lv *Liveness) pointerMap(psess *PackageSession, liveout bvec, vars []*Node, args, locals bvec) {
	for i := int32(0); ; i++ {
		i = liveout.Next(i)
		if i < 0 {
			break
		}
		node := vars[i]
		switch node.Class() {
		case PAUTO:
			psess.
				onebitwalktype1(node.Type, node.Xoffset+lv.stkptrsize, locals)

		case PPARAM, PPARAMOUT:
			psess.
				onebitwalktype1(node.Type, node.Xoffset, args)
		}
	}
}

// markUnsafePoints finds unsafe points and computes lv.unsafePoints.
func (lv *Liveness) markUnsafePoints(psess *PackageSession) {
	if psess.compiling_runtime || lv.f.NoSplit {

		return
	}

	lv.unsafePoints = bvalloc(int32(lv.f.NumValues()))

	for _, wbBlock := range lv.f.WBLoads {
		if wbBlock.Kind == ssa.BlockPlain && len(wbBlock.Values) == 0 {

			continue
		}

		if len(wbBlock.Succs) != 2 {
			lv.f.Fatalf("expected branch at write barrier block %v", wbBlock)
		}
		s0, s1 := wbBlock.Succs[0].Block(), wbBlock.Succs[1].Block()
		if s0.Kind != ssa.BlockPlain || s1.Kind != ssa.BlockPlain {
			lv.f.Fatalf("expected successors of write barrier block %v to be plain", wbBlock)
		}
		if s0.Succs[0].Block() != s1.Succs[0].Block() {
			lv.f.Fatalf("expected successors of write barrier block %v to converge", wbBlock)
		}

		// Flow backwards from the control value to find the
		// flag load. We don't know what lowered ops we're
		// looking for, but all current arches produce a
		// single op that does the memory load from the flag
		// address, so we look for that.
		var load *ssa.Value
		v := wbBlock.Control
		for {
			if sym, ok := v.Aux.(*obj.LSym); ok && sym == psess.writeBarrier {
				load = v
				break
			}
			switch v.Op {
			case ssa.Op386TESTL:

				if v.Args[0] == v.Args[1] {
					v = v.Args[0]
					continue
				}
			case ssa.Op386MOVLload, ssa.OpARM64MOVWUload, ssa.OpPPC64MOVWZload, ssa.OpWasmI64Load32U:

				v = v.Args[0]
				continue
			}

			if len(v.Args) != 1 {
				v.Fatalf("write barrier control value has more than one argument: %s", v.LongString(psess.ssa))
			}
			v = v.Args[0]
		}

		found := false
		for _, v := range wbBlock.Values {
			found = found || v == load
			if found {
				lv.unsafePoints.Set(psess, int32(v.ID))
			}
		}

		for _, succ := range wbBlock.Succs {
			for _, v := range succ.Block().Values {
				lv.unsafePoints.Set(psess, int32(v.ID))
			}
		}
	}

	// Find uintptr -> unsafe.Pointer conversions and flood
	// unsafeness back to a call (which is always a safe point).
	//
	// Looking for the uintptr -> unsafe.Pointer conversion has a
	// few advantages over looking for unsafe.Pointer -> uintptr
	// conversions:
	//
	// 1. We avoid needlessly blocking safe-points for
	// unsafe.Pointer -> uintptr conversions that never go back to
	// a Pointer.
	//
	// 2. We don't have to detect calls to reflect.Value.Pointer,
	// reflect.Value.UnsafeAddr, and reflect.Value.InterfaceData,
	// which are implicit unsafe.Pointer -> uintptr conversions.
	// We can't even reliably detect this if there's an indirect
	// call to one of these methods.
	//
	// TODO: For trivial unsafe.Pointer arithmetic, it would be
	// nice to only flood as far as the unsafe.Pointer -> uintptr
	// conversion, but it's hard to know which argument of an Add
	// or Sub to follow.
	var flooded bvec
	var flood func(b *ssa.Block, vi int)
	flood = func(b *ssa.Block, vi int) {
		if flooded.n == 0 {
			flooded = bvalloc(int32(lv.f.NumBlocks()))
		}
		if flooded.Get(psess, int32(b.ID)) {
			return
		}
		for i := vi - 1; i >= 0; i-- {
			v := b.Values[i]
			if v.Op.IsCall(psess.ssa) {

				return
			}
			lv.unsafePoints.Set(psess, int32(v.ID))
		}
		if vi == len(b.Values) {

			flooded.Set(psess, int32(b.ID))
		}
		for _, pred := range b.Preds {
			flood(pred.Block(), len(pred.Block().Values))
		}
	}
	for _, b := range lv.f.Blocks {
		for i, v := range b.Values {
			if !(v.Op == ssa.OpConvert && v.Type.IsPtrShaped()) {
				continue
			}

			flood(b, i+1)
		}
	}
}

// Returns true for instructions that are safe points that must be annotated
// with liveness information.
func (lv *Liveness) issafepoint(psess *PackageSession, v *ssa.Value) bool {

	if psess.compiling_runtime || lv.f.NoSplit {
		return v.Op.IsCall(psess.ssa)
	}
	switch v.Op {
	case ssa.OpInitMem, ssa.OpArg, ssa.OpSP, ssa.OpSB,
		ssa.OpSelect0, ssa.OpSelect1, ssa.OpGetG,
		ssa.OpVarDef, ssa.OpVarLive, ssa.OpKeepAlive,
		ssa.OpPhi:

		return false
	}
	return !lv.unsafePoints.Get(psess, int32(v.ID))
}

// Initializes the sets for solving the live variables. Visits all the
// instructions in each basic block to summarizes the information at each basic
// block
func (lv *Liveness) prologue(psess *PackageSession) {
	lv.initcache(psess)

	for _, b := range lv.f.Blocks {
		be := lv.blockEffects(b)

		for j := len(b.Values) - 1; j >= 0; j-- {
			pos, e := lv.valueEffects(psess, b.Values[j])
			regUevar, regKill := lv.regEffects(psess, b.Values[j])
			if e&varkill != 0 {
				be.varkill.vars.Set(psess, pos)
				be.uevar.vars.Unset(psess, pos)
			}
			be.varkill.regs |= regKill
			be.uevar.regs &^= regKill
			if e&uevar != 0 {
				be.uevar.vars.Set(psess, pos)
			}
			be.uevar.regs |= regUevar
		}

		for _, val := range b.Values {
			pos, e := lv.valueEffects(psess, val)

			if e&varkill != 0 {
				be.avarinit.Unset(psess, pos)
			}
			if e&avarinit != 0 {
				be.avarinit.Set(psess, pos)
			}
		}
	}
}

// Solve the liveness dataflow equations.
func (lv *Liveness) solve(psess *PackageSession) {

	nvars := int32(len(lv.vars))
	newlivein := varRegVec{vars: bvalloc(nvars)}
	newliveout := varRegVec{vars: bvalloc(nvars)}
	any := bvalloc(nvars)
	all := bvalloc(nvars)

	for _, b := range lv.f.Blocks {
		be := lv.blockEffects(b)
		if b == lv.f.Entry {
			be.avarinitall.Copy(be.avarinit)
		} else {
			be.avarinitall.Clear()
			be.avarinitall.Not()
		}
		be.avarinitany.Copy(be.avarinit)
	}

	po := lv.f.Postorder()

	for change := true; change; {
		change = false
		for i := len(po) - 1; i >= 0; i-- {
			b := po[i]
			be := lv.blockEffects(b)
			lv.avarinitanyall(psess, b, any, all)

			any.AndNot(any, be.varkill.vars)
			all.AndNot(all, be.varkill.vars)
			any.Or(any, be.avarinit)
			all.Or(all, be.avarinit)
			if !any.Eq(psess, be.avarinitany) {
				change = true
				be.avarinitany.Copy(any)
			}

			if !all.Eq(psess, be.avarinitall) {
				change = true
				be.avarinitall.Copy(all)
			}
		}
	}

	for change := true; change; {
		change = false
		for _, b := range po {
			be := lv.blockEffects(b)

			newliveout.Clear()
			switch b.Kind {
			case ssa.BlockRet:
				for _, pos := range lv.cache.retuevar {
					newliveout.vars.Set(psess, pos)
				}
			case ssa.BlockRetJmp:
				for _, pos := range lv.cache.tailuevar {
					newliveout.vars.Set(psess, pos)
				}
			case ssa.BlockExit:

			default:

				newliveout.Copy(lv.blockEffects(b.Succs[0].Block()).livein)
				for _, succ := range b.Succs[1:] {
					newliveout.Or(newliveout, lv.blockEffects(succ.Block()).livein)
				}
			}

			if !be.liveout.Eq(psess, newliveout) {
				change = true
				be.liveout.Copy(newliveout)
			}

			newlivein.AndNot(be.liveout, be.varkill)
			be.livein.Or(newlivein, be.uevar)
		}
	}
}

// Visits all instructions in a basic block and computes a bit vector of live
// variables at each safe point locations.
func (lv *Liveness) epilogue(psess *PackageSession) {
	nvars := int32(len(lv.vars))
	liveout := varRegVec{vars: bvalloc(nvars)}
	any := bvalloc(nvars)
	all := bvalloc(nvars)
	livedefer := bvalloc(nvars)

	if lv.fn.Func.HasDefer() {
		for i, n := range lv.vars {
			if n.Class() == PPARAMOUT {
				if n.IsOutputParamHeapAddr() {
					psess.
						Fatalf("variable %v both output param and heap output param", n)
				}
				if n.Name.Param.Heapaddr != nil {

					continue
				}

				livedefer.Set(psess, int32(i))
			}
			if n.IsOutputParamHeapAddr() {
				n.Name.SetNeedzero(true)
				livedefer.Set(psess, int32(i))
			}
		}
	}

	if lv.f.Entry != lv.f.Blocks[0] {
		lv.f.Fatalf("entry block must be first")
	}

	{

		live := bvalloc(nvars)
		for _, pos := range lv.cache.textavarinit {
			live.Set(psess, pos)
		}
		lv.livevars = append(lv.livevars, varRegVec{vars: live})
	}

	for _, b := range lv.f.Blocks {
		be := lv.blockEffects(b)
		firstBitmapIndex := len(lv.livevars)

		lv.avarinitanyall(psess, b, any, all)

		for _, v := range b.Values {
			pos, e := lv.valueEffects(psess, v)

			if e&varkill != 0 {
				any.Unset(psess, pos)
				all.Unset(psess, pos)
			}
			if e&avarinit != 0 {
				any.Set(psess, pos)
				all.Set(psess, pos)
			}

			if !lv.issafepoint(psess, v) {
				continue
			}

			liveout.vars.AndNot(any, all)
			if !liveout.vars.IsEmpty() {
				for pos := int32(0); pos < liveout.vars.n; pos++ {
					if !liveout.vars.Get(psess, pos) {
						continue
					}
					all.Set(psess, pos)
					n := lv.vars[pos]
					if !n.Name.Needzero() {
						n.Name.SetNeedzero(true)
						if psess.debuglive >= 1 {
							psess.
								Warnl(v.Pos, "%v: %L is ambiguously live", lv.fn.Func.Nname, n)
						}
					}
				}
			}

			live := bvalloc(nvars)
			live.Copy(any)
			lv.livevars = append(lv.livevars, varRegVec{vars: live})
		}

		index := int32(len(lv.livevars) - 1)

		liveout.Copy(be.liveout)
		for i := len(b.Values) - 1; i >= 0; i-- {
			v := b.Values[i]

			if lv.issafepoint(psess, v) {

				live := &lv.livevars[index]
				live.Or(*live, liveout)
				live.vars.Or(live.vars, livedefer)
				index--
			}

			pos, e := lv.valueEffects(psess, v)
			regUevar, regKill := lv.regEffects(psess, v)
			if e&varkill != 0 {
				liveout.vars.Unset(psess, pos)
			}
			liveout.regs &^= regKill
			if e&uevar != 0 {
				liveout.vars.Set(psess, pos)
			}
			liveout.regs |= regUevar
		}

		if b == lv.f.Entry {
			if index != 0 {
				psess.
					Fatalf("bad index for entry point: %v", index)
			}

			live := &lv.livevars[index]
			live.Or(*live, liveout)
		}

		index = int32(firstBitmapIndex)
		for _, v := range b.Values {
			if lv.issafepoint(psess, v) {
				live := lv.livevars[index]
				if v.Op.IsCall(psess.ssa) && live.regs != 0 {
					lv.printDebug(psess)
					v.Fatalf("internal error: %v register %s recorded as live at call", lv.fn.Func.Nname, live.regs.niceString(lv.f.Config))
				}
				index++
			}
		}

		lv.compact(psess, b)
	}

	lv.stackMaps = lv.stackMapSet.extractUniqe()
	lv.stackMapSet = bvecSet{}

	for j, n := range lv.vars {
		if n.Class() != PPARAM && lv.stackMaps[0].Get(psess, int32(j)) {
			psess.
				Fatalf("internal error: %v %L recorded as live on entry", lv.fn.Func.Nname, n)
		}
	}

	if regs := lv.regMaps[0]; regs != 0 {
		lv.printDebug(psess)
		lv.f.Fatalf("internal error: %v register %s recorded as live on entry", lv.fn.Func.Nname, regs.niceString(lv.f.Config))
	}
}

func (lv *Liveness) clobber(psess *PackageSession) {

	if psess.objabi.Clobberdead_enabled == 0 {
		return
	}
	var varSize int64
	for _, n := range lv.vars {
		varSize += n.Type.Size(psess.types)
	}
	if len(lv.stackMaps) > 1000 || varSize > 10000 {

		return
	}
	if h := os.Getenv("GOCLOBBERDEADHASH"); h != "" {

		hstr := ""
		for _, b := range sha1.Sum([]byte(lv.fn.funcname())) {
			hstr += fmt.Sprintf("%08b", b)
		}
		if !strings.HasSuffix(hstr, h) {
			return
		}
		fmt.Printf("\t\t\tCLOBBERDEAD %s\n", lv.fn.funcname())
	}
	if lv.f.Name == "forkAndExecInChild" {

		return
	}

	var oldSched []*ssa.Value
	for _, b := range lv.f.Blocks {

		oldSched = append(oldSched[:0], b.Values...)
		b.Values = b.Values[:0]

		if b == lv.f.Entry {
			for len(oldSched) > 0 && len(oldSched[0].Args) == 0 {

				b.Values = append(b.Values, oldSched[0])
				oldSched = oldSched[1:]
			}
			psess.
				clobber(lv, b, lv.stackMaps[0])
		}

		for _, v := range oldSched {
			if !lv.issafepoint(psess, v) {
				b.Values = append(b.Values, v)
				continue
			}
			before := true
			if v.Op.IsCall(psess.ssa) && v.Aux != nil && v.Aux.(*obj.LSym) == psess.typedmemmove {

				before = false
			}
			if before {
				psess.
					clobber(lv, b, lv.stackMaps[lv.livenessMap.Get(psess, v).stackMapIndex])
			}
			b.Values = append(b.Values, v)
			psess.
				clobber(lv, b, lv.stackMaps[lv.livenessMap.Get(psess, v).stackMapIndex])
		}
	}
}

// clobber generates code to clobber all dead variables (those not marked in live).
// Clobbering instructions are added to the end of b.Values.
func (psess *PackageSession) clobber(lv *Liveness, b *ssa.Block, live bvec) {
	for i, n := range lv.vars {
		if !live.Get(psess, int32(i)) {
			psess.
				clobberVar(b, n)
		}
	}
}

// clobberVar generates code to trash the pointers in v.
// Clobbering instructions are added to the end of b.Values.
func (psess *PackageSession) clobberVar(b *ssa.Block, v *Node) {
	psess.
		clobberWalk(b, v, 0, v.Type)
}

// b = block to which we append instructions
// v = variable
// offset = offset of (sub-portion of) variable to clobber (in bytes)
// t = type of sub-portion of v.
func (psess *PackageSession) clobberWalk(b *ssa.Block, v *Node, offset int64, t *types.Type) {
	if !psess.types.Haspointers(t) {
		return
	}
	switch t.Etype {
	case TPTR32,
		TPTR64,
		TUNSAFEPTR,
		TFUNC,
		TCHAN,
		TMAP:
		psess.
			clobberPtr(b, v, offset)

	case TSTRING:
		psess.
			clobberPtr(b, v, offset)

	case TINTER:
		psess.
			clobberPtr(b, v, offset+int64(psess.Widthptr))

	case TSLICE:
		psess.
			clobberPtr(b, v, offset)

	case TARRAY:
		for i := int64(0); i < t.NumElem(psess.types); i++ {
			psess.
				clobberWalk(b, v, offset+i*t.Elem(psess.types).Size(psess.types), t.Elem(psess.types))
		}

	case TSTRUCT:
		for _, t1 := range t.Fields(psess.types).Slice() {
			psess.
				clobberWalk(b, v, offset+t1.Offset, t1.Type)
		}

	default:
		psess.
			Fatalf("clobberWalk: unexpected type, %v", t)
	}
}

// clobberPtr generates a clobber of the pointer at offset offset in v.
// The clobber instruction is added at the end of b.
func (psess *PackageSession) clobberPtr(b *ssa.Block, v *Node, offset int64) {
	b.NewValue0IA(psess.src.NoXPos, ssa.OpClobber, psess.types.TypeVoid, offset, v)
}

func (lv *Liveness) avarinitanyall(psess *PackageSession, b *ssa.Block, any, all bvec) {
	if len(b.Preds) == 0 {
		any.Clear()
		all.Clear()
		for _, pos := range lv.cache.textavarinit {
			any.Set(psess, pos)
			all.Set(psess, pos)
		}
		return
	}

	be := lv.blockEffects(b.Preds[0].Block())
	any.Copy(be.avarinitany)
	all.Copy(be.avarinitall)

	for _, pred := range b.Preds[1:] {
		be := lv.blockEffects(pred.Block())
		any.Or(any, be.avarinitany)
		all.And(all, be.avarinitall)
	}
}

// Compact coalesces identical bitmaps from lv.livevars into the sets
// lv.stackMapSet and lv.regMaps.
//
// Compact clears lv.livevars.
//
// There are actually two lists of bitmaps, one list for the local variables and one
// list for the function arguments. Both lists are indexed by the same PCDATA
// index, so the corresponding pairs must be considered together when
// merging duplicates. The argument bitmaps change much less often during
// function execution than the local variable bitmaps, so it is possible that
// we could introduce a separate PCDATA index for arguments vs locals and
// then compact the set of argument bitmaps separately from the set of
// local variable bitmaps. As of 2014-04-02, doing this to the godoc binary
// is actually a net loss: we save about 50k of argument bitmaps but the new
// PCDATA tables cost about 100k. So for now we keep using a single index for
// both bitmap lists.
func (lv *Liveness) compact(psess *PackageSession, b *ssa.Block) {
	add := func(live varRegVec) LivenessIndex {

		stackIndex := lv.stackMapSet.add(psess, live.vars)

		regIndex, ok := lv.regMapSet[live.regs]
		if !ok {
			regIndex = len(lv.regMapSet)
			lv.regMapSet[live.regs] = regIndex
			lv.regMaps = append(lv.regMaps, live.regs)
		}
		return LivenessIndex{stackIndex, regIndex}
	}
	pos := 0
	if b == lv.f.Entry {

		add(lv.livevars[0])
		pos++
	}
	for _, v := range b.Values {
		if lv.issafepoint(psess, v) {
			lv.livenessMap.set(v, add(lv.livevars[pos]))
			pos++
		}
	}

	lv.livevars = lv.livevars[:0]
}

func (lv *Liveness) showlive(psess *PackageSession, v *ssa.Value, live bvec) {
	if psess.debuglive == 0 || lv.fn.funcname() == "init" || strings.HasPrefix(lv.fn.funcname(), ".") {
		return
	}
	if !(v == nil || v.Op.IsCall(psess.ssa)) {

		return
	}
	if live.IsEmpty() {
		return
	}

	pos := lv.fn.Func.Nname.Pos
	if v != nil {
		pos = v.Pos
	}

	s := "live at "
	if v == nil {
		s += fmt.Sprintf("entry to %s:", lv.fn.funcname())
	} else if sym, ok := v.Aux.(*obj.LSym); ok {
		fn := sym.Name
		if pos := strings.Index(fn, "."); pos >= 0 {
			fn = fn[pos+1:]
		}
		s += fmt.Sprintf("call to %s:", fn)
	} else {
		s += "indirect call:"
	}

	for j, n := range lv.vars {
		if live.Get(psess, int32(j)) {
			s += fmt.Sprintf(" %v", n)
		}
	}
	psess.
		Warnl(pos, s)
}

func (lv *Liveness) printbvec(psess *PackageSession, printed bool, name string, live varRegVec) bool {
	if live.vars.IsEmpty() && live.regs == 0 {
		return printed
	}

	if !printed {
		fmt.Printf("\t")
	} else {
		fmt.Printf(" ")
	}
	fmt.Printf("%s=", name)

	comma := ""
	for i, n := range lv.vars {
		if !live.vars.Get(psess, int32(i)) {
			continue
		}
		fmt.Printf("%s%s", comma, n.Sym.Name)
		comma = ","
	}
	fmt.Printf("%s%s", comma, live.regs.niceString(lv.f.Config))
	return true
}

// printeffect is like printbvec, but for valueEffects and regEffects.
func (lv *Liveness) printeffect(printed bool, name string, pos int32, x bool, regMask liveRegMask) bool {
	if !x && regMask == 0 {
		return printed
	}
	if !printed {
		fmt.Printf("\t")
	} else {
		fmt.Printf(" ")
	}
	fmt.Printf("%s=", name)
	if x {
		fmt.Printf("%s", lv.vars[pos].Sym.Name)
	}
	for j, reg := range lv.f.Config.GCRegMap {
		if regMask&(1<<uint(j)) != 0 {
			if x {
				fmt.Printf(",")
			}
			x = true
			fmt.Printf("%v", reg)
		}
	}
	return true
}

// Prints the computed liveness information and inputs, for debugging.
// This format synthesizes the information used during the multiple passes
// into a single presentation.
func (lv *Liveness) printDebug(psess *PackageSession) {
	fmt.Printf("liveness: %s\n", lv.fn.funcname())

	pcdata := 0
	for i, b := range lv.f.Blocks {
		if i > 0 {
			fmt.Printf("\n")
		}

		fmt.Printf("bb#%d pred=", b.ID)
		for j, pred := range b.Preds {
			if j > 0 {
				fmt.Printf(",")
			}
			fmt.Printf("%d", pred.Block().ID)
		}
		fmt.Printf(" succ=")
		for j, succ := range b.Succs {
			if j > 0 {
				fmt.Printf(",")
			}
			fmt.Printf("%d", succ.Block().ID)
		}
		fmt.Printf("\n")

		be := lv.blockEffects(b)

		printed := false
		printed = lv.printbvec(psess, printed, "uevar", be.uevar)
		printed = lv.printbvec(psess, printed, "livein", be.livein)
		if printed {
			fmt.Printf("\n")
		}

		if b == lv.f.Entry {
			live := lv.stackMaps[pcdata]
			fmt.Printf("(%s) function entry\n", psess.linestr(lv.fn.Func.Nname.Pos))
			fmt.Printf("\tlive=")
			printed = false
			for j, n := range lv.vars {
				if !live.Get(psess, int32(j)) {
					continue
				}
				if printed {
					fmt.Printf(",")
				}
				fmt.Printf("%v", n)
				printed = true
			}
			fmt.Printf("\n")
		}

		for _, v := range b.Values {
			fmt.Printf("(%s) %v\n", psess.linestr(v.Pos), v.LongString(psess.ssa))

			if pos := lv.livenessMap.Get(psess, v); pos.Valid() {
				pcdata = pos.stackMapIndex
			}

			pos, effect := lv.valueEffects(psess, v)
			regUevar, regKill := lv.regEffects(psess, v)
			printed = false
			printed = lv.printeffect(printed, "uevar", pos, effect&uevar != 0, regUevar)
			printed = lv.printeffect(printed, "varkill", pos, effect&varkill != 0, regKill)
			printed = lv.printeffect(printed, "avarinit", pos, effect&avarinit != 0, 0)
			if printed {
				fmt.Printf("\n")
			}

			if !lv.issafepoint(psess, v) {
				continue
			}

			live := lv.stackMaps[pcdata]
			fmt.Printf("\tlive=")
			printed = false
			for j, n := range lv.vars {
				if !live.Get(psess, int32(j)) {
					continue
				}
				if printed {
					fmt.Printf(",")
				}
				fmt.Printf("%v", n)
				printed = true
			}
			regLive := lv.regMaps[lv.livenessMap.Get(psess, v).regMapIndex]
			if regLive != 0 {
				if printed {
					fmt.Printf(",")
				}
				fmt.Printf("%s", regLive.niceString(lv.f.Config))
			}
			fmt.Printf("\n")
		}

		fmt.Printf("end\n")
		printed = false
		printed = lv.printbvec(psess, printed, "varkill", be.varkill)
		printed = lv.printbvec(psess, printed, "liveout", be.liveout)
		printed = lv.printbvec(psess, printed, "avarinit", varRegVec{vars: be.avarinit})
		printed = lv.printbvec(psess, printed, "avarinitany", varRegVec{vars: be.avarinitany})
		printed = lv.printbvec(psess, printed, "avarinitall", varRegVec{vars: be.avarinitall})
		if printed {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")
}

// Dumps a slice of bitmaps to a symbol as a sequence of uint32 values. The
// first word dumped is the total number of bitmaps. The second word is the
// length of the bitmaps. All bitmaps are assumed to be of equal length. The
// remaining bytes are the raw bitmaps.
func (lv *Liveness) emit(psess *PackageSession, argssym, livesym, regssym *obj.LSym) {
	// Size args bitmaps to be just large enough to hold the largest pointer.
	// First, find the largest Xoffset node we care about.
	// (Nodes without pointers aren't in lv.vars; see livenessShouldTrack.)
	var maxArgNode *Node
	for _, n := range lv.vars {
		switch n.Class() {
		case PPARAM, PPARAMOUT:
			if maxArgNode == nil || n.Xoffset > maxArgNode.Xoffset {
				maxArgNode = n
			}
		}
	}
	// Next, find the offset of the largest pointer in the largest node.
	var maxArgs int64
	if maxArgNode != nil {
		maxArgs = maxArgNode.Xoffset + psess.typeptrdata(maxArgNode.Type)
	}

	maxLocals := lv.stkptrsize

	args := bvalloc(int32(maxArgs / int64(psess.Widthptr)))
	aoff := psess.duint32(argssym, 0, uint32(len(lv.stackMaps)))
	aoff = psess.duint32(argssym, aoff, uint32(args.n))

	locals := bvalloc(int32(maxLocals / int64(psess.Widthptr)))
	loff := psess.duint32(livesym, 0, uint32(len(lv.stackMaps)))
	loff = psess.duint32(livesym, loff, uint32(locals.n))

	for _, live := range lv.stackMaps {
		args.Clear()
		locals.Clear()

		lv.pointerMap(psess, live, lv.vars, args, locals)

		aoff = psess.dbvec(argssym, aoff, args)
		loff = psess.dbvec(livesym, loff, locals)
	}

	regs := bvalloc(lv.usedRegs())
	roff := psess.duint32(regssym, 0, uint32(len(lv.regMaps)))
	roff = psess.duint32(regssym, roff, uint32(regs.n))
	if regs.n > 32 {
		psess.
			Fatalf("GP registers overflow uint32")
	}

	if regs.n > 0 {
		for _, live := range lv.regMaps {
			regs.Clear()
			regs.b[0] = uint32(live)
			roff = psess.dbvec(regssym, roff, regs)
		}
	}

	argssym.Name = fmt.Sprintf("gclocals·%x", md5.Sum(argssym.P))
	livesym.Name = fmt.Sprintf("gclocals·%x", md5.Sum(livesym.P))
	regssym.Name = fmt.Sprintf("gclocals·%x", md5.Sum(regssym.P))
}

// Entry pointer for liveness analysis. Solves for the liveness of
// pointer variables in the function and emits a runtime data
// structure read by the garbage collector.
// Returns a map from GC safe points to their corresponding stack map index.
func (psess *PackageSession) liveness(e *ssafn, f *ssa.Func) LivenessMap {

	vars, idx := psess.getvariables(e.curfn)
	lv := psess.newliveness(e.curfn, f, vars, idx, e.stkptrsize)

	lv.prologue(psess)
	lv.solve(psess)
	lv.epilogue(psess)
	lv.clobber(psess)
	if psess.debuglive > 0 {
		lv.showlive(psess, nil, lv.stackMaps[0])
		for _, b := range f.Blocks {
			for _, val := range b.Values {
				if idx := lv.livenessMap.Get(psess, val); idx.Valid() {
					lv.showlive(psess, val, lv.stackMaps[idx.stackMapIndex])
				}
			}
		}
	}
	if psess.debuglive >= 2 {
		lv.printDebug(psess)
	}

	{
		cache := f.Cache.Liveness.(*livenessFuncCache)
		if cap(lv.be) < 2000 {
			for i := range lv.be {
				lv.be[i] = BlockEffects{}
			}
			cache.be = lv.be
		}
		if cap(lv.livenessMap.m) < 2000 {
			cache.livenessMap = lv.livenessMap
		}
	}

	if ls := e.curfn.Func.lsym; ls != nil {
		lv.emit(psess, &ls.Func.GCArgs, &ls.Func.GCLocals, &ls.Func.GCRegs)
	}
	return lv.livenessMap
}

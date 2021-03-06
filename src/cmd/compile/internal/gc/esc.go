// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"strconv"
	"strings"
)

// Run analysis on minimal sets of mutually recursive functions
// or single non-recursive functions, bottom up.
//
// Finding these sets is finding strongly connected components
// by reverse topological order in the static call graph.
// The algorithm (known as Tarjan's algorithm) for doing that is taken from
// Sedgewick, Algorithms, Second Edition, p. 482, with two adaptations.
//
// First, a hidden closure function (n.Func.IsHiddenClosure()) cannot be the
// root of a connected component. Refusing to use it as a root
// forces it into the component of the function in which it appears.
// This is more convenient for escape analysis.
//
// Second, each function becomes two virtual nodes in the graph,
// with numbers n and n+1. We record the function's node number as n
// but search from node n+1. If the search tells us that the component
// number (min) is n+1, we know that this is a trivial component: one function
// plus its closures. If the search tells us that the component number is
// n, then there was a path from node n+1 back to node n, meaning that
// the function set is mutually recursive. The escape analysis can be
// more precise when analyzing a single non-recursive function than
// when analyzing a set of mutually recursive functions.

type bottomUpVisitor struct {
	analyze  func([]*Node, bool)
	visitgen uint32
	nodeID   map[*Node]uint32
	stack    []*Node
}

// visitBottomUp invokes analyze on the ODCLFUNC nodes listed in list.
// It calls analyze with successive groups of functions, working from
// the bottom of the call graph upward. Each time analyze is called with
// a list of functions, every function on that list only calls other functions
// on the list or functions that have been passed in previous invocations of
// analyze. Closures appear in the same list as their outer functions.
// The lists are as short as possible while preserving those requirements.
// (In a typical program, many invocations of analyze will be passed just
// a single function.) The boolean argument 'recursive' passed to analyze
// specifies whether the functions on the list are mutually recursive.
// If recursive is false, the list consists of only a single function and its closures.
// If recursive is true, the list may still contain only a single function,
// if that function is itself recursive.
func (pstate *PackageState) visitBottomUp(list []*Node, analyze func(list []*Node, recursive bool)) {
	var v bottomUpVisitor
	v.analyze = analyze
	v.nodeID = make(map[*Node]uint32)
	for _, n := range list {
		if n.Op == ODCLFUNC && !n.Func.IsHiddenClosure() {
			v.visit(pstate, n)
		}
	}
}

func (v *bottomUpVisitor) visit(pstate *PackageState, n *Node) uint32 {
	if id := v.nodeID[n]; id > 0 {
		// already visited
		return id
	}

	v.visitgen++
	id := v.visitgen
	v.nodeID[n] = id
	v.visitgen++
	min := v.visitgen

	v.stack = append(v.stack, n)
	min = v.visitcodelist(pstate, n.Nbody, min)
	if (min == id || min == id+1) && !n.Func.IsHiddenClosure() {
		// This node is the root of a strongly connected component.

		// The original min passed to visitcodelist was v.nodeID[n]+1.
		// If visitcodelist found its way back to v.nodeID[n], then this
		// block is a set of mutually recursive functions.
		// Otherwise it's just a lone function that does not recurse.
		recursive := min == id

		// Remove connected component from stack.
		// Mark walkgen so that future visits return a large number
		// so as not to affect the caller's min.

		var i int
		for i = len(v.stack) - 1; i >= 0; i-- {
			x := v.stack[i]
			if x == n {
				break
			}
			v.nodeID[x] = ^uint32(0)
		}
		v.nodeID[n] = ^uint32(0)
		block := v.stack[i:]
		// Run escape analysis on this set of functions.
		v.stack = v.stack[:i]
		v.analyze(block, recursive)
	}

	return min
}

func (v *bottomUpVisitor) visitcodelist(pstate *PackageState, l Nodes, min uint32) uint32 {
	for _, n := range l.Slice() {
		min = v.visitcode(pstate, n, min)
	}
	return min
}

func (v *bottomUpVisitor) visitcode(pstate *PackageState, n *Node, min uint32) uint32 {
	if n == nil {
		return min
	}

	min = v.visitcodelist(pstate, n.Ninit, min)
	min = v.visitcode(pstate, n.Left, min)
	min = v.visitcode(pstate, n.Right, min)
	min = v.visitcodelist(pstate, n.List, min)
	min = v.visitcodelist(pstate, n.Nbody, min)
	min = v.visitcodelist(pstate, n.Rlist, min)

	switch n.Op {
	case OCALLFUNC, OCALLMETH:
		fn := asNode(n.Left.Type.Nname(pstate.types))
		if fn != nil && fn.Op == ONAME && fn.Class() == PFUNC && fn.Name.Defn != nil {
			m := v.visit(pstate, fn.Name.Defn)
			if m < min {
				min = m
			}
		}

	case OCLOSURE:
		m := v.visit(pstate, n.Func.Closure)
		if m < min {
			min = m
		}
	}

	return min
}

// Escape analysis.

// An escape analysis pass for a set of functions.
// The analysis assumes that closures and the functions in which they
// appear are analyzed together, so that the aliasing between their
// variables can be modeled more precisely.
//
// First escfunc, esc and escassign recurse over the ast of each
// function to dig out flow(dst,src) edges between any
// pointer-containing nodes and store them in e.nodeEscState(dst).Flowsrc. For
// variables assigned to a variable in an outer scope or used as a
// return value, they store a flow(theSink, src) edge to a fake node
// 'the Sink'.  For variables referenced in closures, an edge
// flow(closure, &var) is recorded and the flow of a closure itself to
// an outer scope is tracked the same way as other variables.
//
// Then escflood walks the graph starting at theSink and tags all
// variables of it can reach an & node as escaping and all function
// parameters it can reach as leaking.
//
// If a value's address is taken but the address does not escape,
// then the value can stay on the stack. If the value new(T) does
// not escape, then new(T) can be rewritten into a stack allocation.
// The same is true of slice literals.

func (pstate *PackageState) escapes(all []*Node) {
	pstate.visitBottomUp(all, pstate.escAnalyze)
}

const (
	EscFuncUnknown = 0 + iota
	EscFuncPlanned
	EscFuncStarted
	EscFuncTagged
)

// There appear to be some loops in the escape graph, causing
// arbitrary recursion into deeper and deeper levels.
// Cut this off safely by making minLevel sticky: once you
// get that deep, you cannot go down any further but you also
// cannot go up any further. This is a conservative fix.
// Making minLevel smaller (more negative) would handle more
// complex chains of indirections followed by address-of operations,
// at the cost of repeating the traversal once for each additional
// allowed level when a loop is encountered. Using -2 suffices to
// pass all the tests we have written so far, which we assume matches
// the level of complexity we want the escape analysis code to handle.
const MinLevel = -2

// A Level encodes the reference state and context applied to
// (stack, heap) allocated memory.
//
// value is the overall sum of *(1) and &(-1) operations encountered
// along a path from a destination (sink, return value) to a source
// (allocation, parameter).
//
// suffixValue is the maximum-copy-started-suffix-level applied to a sink.
// For example:
// sink = x.left.left --> level=2, x is dereferenced twice and does not escape to sink.
// sink = &Node{x} --> level=-1, x is accessible from sink via one "address of"
// sink = &Node{&Node{x}} --> level=-2, x is accessible from sink via two "address of"
// sink = &Node{&Node{x.left}} --> level=-1, but x is NOT accessible from sink because it was indirected and then copied.
// (The copy operations are sometimes implicit in the source code; in this case,
// value of x.left was copied into a field of a newly allocated Node)
//
// There's one of these for each Node, and the integer values
// rarely exceed even what can be stored in 4 bits, never mind 8.
type Level struct {
	value, suffixValue int8
}

func (l Level) int() int {
	return int(l.value)
}

func levelFrom(i int) Level {
	if i <= MinLevel {
		return Level{value: MinLevel}
	}
	return Level{value: int8(i)}
}

func satInc8(x int8) int8 {
	if x == 127 {
		return 127
	}
	return x + 1
}

func min8(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

func max8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

// inc returns the level l + 1, representing the effect of an indirect (*) operation.
func (l Level) inc() Level {
	if l.value <= MinLevel {
		return Level{value: MinLevel}
	}
	return Level{value: satInc8(l.value), suffixValue: satInc8(l.suffixValue)}
}

// dec returns the level l - 1, representing the effect of an address-of (&) operation.
func (l Level) dec() Level {
	if l.value <= MinLevel {
		return Level{value: MinLevel}
	}
	return Level{value: l.value - 1, suffixValue: l.suffixValue - 1}
}

// copy returns the level for a copy of a value with level l.
func (l Level) copy() Level {
	return Level{value: l.value, suffixValue: max8(l.suffixValue, 0)}
}

func (l1 Level) min(l2 Level) Level {
	return Level{
		value:       min8(l1.value, l2.value),
		suffixValue: min8(l1.suffixValue, l2.suffixValue)}
}

// guaranteedDereference returns the number of dereferences
// applied to a pointer before addresses are taken/generated.
// This is the maximum level computed from path suffixes starting
// with copies where paths flow from destination to source.
func (l Level) guaranteedDereference() int {
	return int(l.suffixValue)
}

// An EscStep documents one step in the path from memory
// that is heap allocated to the (alleged) reason for the
// heap allocation.
type EscStep struct {
	src, dst *Node    // the endpoints of this edge in the escape-to-heap chain.
	where    *Node    // sometimes the endpoints don't match source locations; set 'where' to make that right
	parent   *EscStep // used in flood to record path
	why      string   // explanation for this step in the escape-to-heap chain
	busy     bool     // used in prevent to snip cycles.
}

type NodeEscState struct {
	Curfn             *Node
	Flowsrc           []EscStep // flow(this, src)
	Retval            Nodes     // on OCALLxxx, list of dummy return values
	Loopdepth         int32     // -1: global, 0: return variables, 1:function top level, increased inside function for every loop or label to mark scopes
	Level             Level
	Walkgen           uint32
	Maxextraloopdepth int32
}

func (e *EscState) nodeEscState(pstate *PackageState, n *Node) *NodeEscState {
	if nE, ok := n.Opt().(*NodeEscState); ok {
		return nE
	}
	if n.Opt() != nil {
		pstate.Fatalf("nodeEscState: opt in use (%T)", n.Opt())
	}
	nE := &NodeEscState{
		Curfn: pstate.Curfn,
	}
	n.SetOpt(pstate, nE)
	e.opts = append(e.opts, n)
	return nE
}

func (e *EscState) track(pstate *PackageState, n *Node) {
	if pstate.Curfn == nil {
		pstate.Fatalf("EscState.track: Curfn nil")
	}
	n.Esc = EscNone // until proven otherwise
	nE := e.nodeEscState(pstate, n)
	nE.Loopdepth = e.loopdepth
	e.noesc = append(e.noesc, n)
}

// Escape constants are numbered in order of increasing "escapiness"
// to help make inferences be monotonic. With the exception of
// EscNever which is sticky, eX < eY means that eY is more exposed
// than eX, and hence replaces it in a conservative analysis.
const (
	EscUnknown        = iota
	EscNone           // Does not escape to heap, result, or parameters.
	EscReturn         // Is returned or reachable from returned.
	EscHeap           // Reachable from the heap
	EscNever          // By construction will not escape.
	EscBits           = 3
	EscMask           = (1 << EscBits) - 1
	EscContentEscapes = 1 << EscBits // value obtained by indirect of parameter escapes to heap
	EscReturnBits     = EscBits + 1

// Node.esc encoding = | escapeReturnEncoding:(width-4) | contentEscapes:1 | escEnum:3
)

// escMax returns the maximum of an existing escape value
// (and its additional parameter flow flags) and a new escape type.
func (pstate *PackageState) escMax(e, etype uint16) uint16 {
	if e&EscMask >= EscHeap {
		// normalize
		if e&^EscMask != 0 {
			pstate.Fatalf("Escape information had unexpected return encoding bits (w/ EscHeap, EscNever), e&EscMask=%v", e&EscMask)
		}
	}
	if e&EscMask > etype {
		return e
	}
	if etype == EscNone || etype == EscReturn {
		return (e &^ EscMask) | etype
	}
	return etype
}

// For each input parameter to a function, the escapeReturnEncoding describes
// how the parameter may leak to the function's outputs. This is currently the
// "level" of the leak where level is 0 or larger (negative level means stored into
// something whose address is returned -- but that implies stored into the heap,
// hence EscHeap, which means that the details are not currently relevant. )
const (
	bitsPerOutputInTag = 3                                 // For each output, the number of bits for a tag
	bitsMaskForTag     = uint16(1<<bitsPerOutputInTag) - 1 // The bit mask to extract a single tag.
	maxEncodedLevel    = int(bitsMaskForTag - 1)           // The largest level that can be stored in a tag.
)

type EscState struct {
	// Fake node that all
	//   - return values and output variables
	//   - parameters on imported functions not marked 'safe'
	//   - assignments to global variables
	// flow to.
	theSink Node

	dsts      []*Node // all dst nodes
	loopdepth int32   // for detecting nested loop scopes
	pdepth    int     // for debug printing in recursions.
	dstcount  int     // diagnostic
	edgecount int     // diagnostic
	noesc     []*Node // list of possible non-escaping nodes, for printing
	recursive bool    // recursive function or group of mutually recursive functions.
	opts      []*Node // nodes with .Opt initialized
	walkgen   uint32
}

func (pstate *PackageState) newEscState(recursive bool) *EscState {
	e := new(EscState)
	e.theSink.Op = ONAME
	e.theSink.Orig = &e.theSink
	e.theSink.SetClass(PEXTERN)
	e.theSink.Sym = pstate.lookup(".sink")
	e.nodeEscState(pstate, &e.theSink).Loopdepth = -1
	e.recursive = recursive
	return e
}

func (e *EscState) stepWalk(pstate *PackageState, dst, src *Node, why string, parent *EscStep) *EscStep {
	// TODO: keep a cache of these, mark entry/exit in escwalk to avoid allocation
	// Or perhaps never mind, since it is disabled unless printing is on.
	// We may want to revisit this, since the EscStep nodes would make
	// an excellent replacement for the poorly-separated graph-build/graph-flood
	// stages.
	if pstate.Debug['m'] == 0 {
		return nil
	}
	return &EscStep{src: src, dst: dst, why: why, parent: parent}
}

func (e *EscState) stepAssign(pstate *PackageState, step *EscStep, dst, src *Node, why string) *EscStep {
	if pstate.Debug['m'] == 0 {
		return nil
	}
	if step != nil { // Caller may have known better.
		if step.why == "" {
			step.why = why
		}
		if step.dst == nil {
			step.dst = dst
		}
		if step.src == nil {
			step.src = src
		}
		return step
	}
	return &EscStep{src: src, dst: dst, why: why}
}

func (e *EscState) stepAssignWhere(pstate *PackageState, dst, src *Node, why string, where *Node) *EscStep {
	if pstate.Debug['m'] == 0 {
		return nil
	}
	return &EscStep{src: src, dst: dst, why: why, where: where}
}

// funcSym returns fn.Func.Nname.Sym if no nils are encountered along the way.
func funcSym(fn *Node) *types.Sym {
	if fn == nil || fn.Func.Nname == nil {
		return nil
	}
	return fn.Func.Nname.Sym
}

// curfnSym returns n.Curfn.Nname.Sym if no nils are encountered along the way.
func (e *EscState) curfnSym(pstate *PackageState, n *Node) *types.Sym {
	nE := e.nodeEscState(pstate, n)
	return funcSym(nE.Curfn)
}

func (pstate *PackageState) escAnalyze(all []*Node, recursive bool) {
	e := pstate.newEscState(recursive)

	for _, n := range all {
		if n.Op == ODCLFUNC {
			n.Esc = EscFuncPlanned
			if pstate.Debug['m'] > 3 {
				Dump("escAnalyze", n)
			}

		}
	}

	// flow-analyze functions
	for _, n := range all {
		if n.Op == ODCLFUNC {
			e.escfunc(pstate, n)
		}
	}

	// print("escapes: %d e.dsts, %d edges\n", e.dstcount, e.edgecount);

	// visit the upstream of each dst, mark address nodes with
	// addrescapes, mark parameters unsafe
	escapes := make([]uint16, len(e.dsts))
	for i, n := range e.dsts {
		escapes[i] = n.Esc
	}
	for _, n := range e.dsts {
		e.escflood(pstate, n)
	}
	for {
		done := true
		for i, n := range e.dsts {
			if n.Esc != escapes[i] {
				done = false
				if pstate.Debug['m'] > 2 {
					pstate.Warnl(n.Pos, "Reflooding %v %S", e.curfnSym(pstate, n), n)
				}
				escapes[i] = n.Esc
				e.escflood(pstate, n)
			}
		}
		if done {
			break
		}
	}

	// for all top level functions, tag the typenodes corresponding to the param nodes
	for _, n := range all {
		if n.Op == ODCLFUNC {
			e.esctag(pstate, n)
		}
	}

	if pstate.Debug['m'] != 0 {
		for _, n := range e.noesc {
			if n.Esc == EscNone {
				pstate.Warnl(n.Pos, "%v %S does not escape", e.curfnSym(pstate, n), n)
			}
		}
	}

	for _, x := range e.opts {
		x.SetOpt(pstate, nil)
	}
}

func (e *EscState) escfunc(pstate *PackageState, fn *Node) {
	//	print("escfunc %N %s\n", fn.Func.Nname, e.recursive?"(recursive)":"");
	if fn.Esc != EscFuncPlanned {
		pstate.Fatalf("repeat escfunc %v", fn.Func.Nname)
	}
	fn.Esc = EscFuncStarted

	saveld := e.loopdepth
	e.loopdepth = 1
	savefn := pstate.Curfn
	pstate.Curfn = fn

	for _, ln := range pstate.Curfn.Func.Dcl {
		if ln.Op != ONAME {
			continue
		}
		lnE := e.nodeEscState(pstate, ln)
		switch ln.Class() {
		// out params are in a loopdepth between the sink and all local variables
		case PPARAMOUT:
			lnE.Loopdepth = 0

		case PPARAM:
			lnE.Loopdepth = 1
			if ln.Type != nil && !pstate.types.Haspointers(ln.Type) {
				break
			}
			if pstate.Curfn.Nbody.Len() == 0 && !pstate.Curfn.Noescape() {
				ln.Esc = EscHeap
			} else {
				ln.Esc = EscNone // prime for escflood later
			}
			e.noesc = append(e.noesc, ln)
		}
	}

	// in a mutually recursive group we lose track of the return values
	if e.recursive {
		for _, ln := range pstate.Curfn.Func.Dcl {
			if ln.Op == ONAME && ln.Class() == PPARAMOUT {
				e.escflows(pstate, &e.theSink, ln, e.stepAssign(pstate, nil, ln, ln, "returned from recursive function"))
			}
		}
	}

	e.escloopdepthlist(pstate, pstate.Curfn.Nbody)
	e.esclist(pstate, pstate.Curfn.Nbody, pstate.Curfn)
	pstate.Curfn = savefn
	e.loopdepth = saveld
}

func (e *EscState) escloopdepthlist(pstate *PackageState, l Nodes) {
	for _, n := range l.Slice() {
		e.escloopdepth(pstate, n)
	}
}

func (e *EscState) escloopdepth(pstate *PackageState, n *Node) {
	if n == nil {
		return
	}

	e.escloopdepthlist(pstate, n.Ninit)

	switch n.Op {
	case OLABEL:
		if n.Left == nil || n.Left.Sym == nil {
			pstate.Fatalf("esc:label without label: %+v", n)
		}

		// Walk will complain about this label being already defined, but that's not until
		// after escape analysis. in the future, maybe pull label & goto analysis out of walk and put before esc
		// if(n.Left.Sym.Label != nil)
		//	fatal("escape analysis messed up analyzing label: %+N", n);
		n.Left.Sym.Label = asTypesNode(&pstate.nonlooping)

	case OGOTO:
		if n.Left == nil || n.Left.Sym == nil {
			pstate.Fatalf("esc:goto without label: %+v", n)
		}

		// If we come past one that's uninitialized, this must be a (harmless) forward jump
		// but if it's set to nonlooping the label must have preceded this goto.
		if asNode(n.Left.Sym.Label) == &pstate.nonlooping {
			n.Left.Sym.Label = asTypesNode(&pstate.looping)
		}
	}

	e.escloopdepth(pstate, n.Left)
	e.escloopdepth(pstate, n.Right)
	e.escloopdepthlist(pstate, n.List)
	e.escloopdepthlist(pstate, n.Nbody)
	e.escloopdepthlist(pstate, n.Rlist)
}

func (e *EscState) esclist(pstate *PackageState, l Nodes, parent *Node) {
	for _, n := range l.Slice() {
		e.esc(pstate, n, parent)
	}
}

func (e *EscState) esc(pstate *PackageState, n *Node, parent *Node) {
	if n == nil {
		return
	}

	lno := pstate.setlineno(n)

	// ninit logically runs at a different loopdepth than the rest of the for loop.
	e.esclist(pstate, n.Ninit, n)

	if n.Op == OFOR || n.Op == OFORUNTIL || n.Op == ORANGE {
		e.loopdepth++
	}

	// type switch variables have no ODCL.
	// process type switch as declaration.
	// must happen before processing of switch body,
	// so before recursion.
	if n.Op == OSWITCH && n.Left != nil && n.Left.Op == OTYPESW {
		for _, cas := range n.List.Slice() { // cases
			// it.N().Rlist is the variable per case
			if cas.Rlist.Len() != 0 {
				e.nodeEscState(pstate, cas.Rlist.First()).Loopdepth = e.loopdepth
			}
		}
	}

	// Big stuff and non-constant-sized stuff escapes unconditionally.
	// "Big" conditions that were scattered around in walk have been
	// gathered here.
	if n.Esc != EscHeap && n.Type != nil &&
		(n.Type.Width > maxStackVarSize ||
			(n.Op == ONEW || n.Op == OPTRLIT) && n.Type.Elem(pstate.types).Width >= 1<<16 ||
			n.Op == OMAKESLICE && !pstate.isSmallMakeSlice(n)) {

		// isSmallMakeSlice returns false for non-constant len/cap.
		// If that's the case, print a more accurate escape reason.
		var msgVerb, escapeMsg string
		if n.Op == OMAKESLICE && (!pstate.Isconst(n.Left, CTINT) || !pstate.Isconst(n.Right, CTINT)) {
			msgVerb, escapeMsg = "has ", "non-constant size"
		} else {
			msgVerb, escapeMsg = "is ", "too large for stack"
		}

		if pstate.Debug['m'] > 2 {
			pstate.Warnl(n.Pos, "%v "+msgVerb+escapeMsg, n)
		}
		n.Esc = EscHeap
		pstate.addrescapes(n)
		e.escassignSinkWhy(pstate, n, n, escapeMsg) // TODO category: tooLarge
	}

	e.esc(pstate, n.Left, n)

	if n.Op == ORANGE {
		// ORANGE node's Right is evaluated before the loop
		e.loopdepth--
	}

	e.esc(pstate, n.Right, n)

	if n.Op == ORANGE {
		e.loopdepth++
	}

	e.esclist(pstate, n.Nbody, n)
	e.esclist(pstate, n.List, n)
	e.esclist(pstate, n.Rlist, n)

	if n.Op == OFOR || n.Op == OFORUNTIL || n.Op == ORANGE {
		e.loopdepth--
	}

	if pstate.Debug['m'] > 2 {
		fmt.Printf("%v:[%d] %v esc: %v\n", pstate.linestr(pstate.lineno), e.loopdepth, funcSym(pstate.Curfn), n)
	}

opSwitch:
	switch n.Op {
	// Record loop depth at declaration.
	case ODCL:
		if n.Left != nil {
			e.nodeEscState(pstate, n.Left).Loopdepth = e.loopdepth
		}

	case OLABEL:
		if asNode(n.Left.Sym.Label) == &pstate.nonlooping {
			if pstate.Debug['m'] > 2 {
				fmt.Printf("%v:%v non-looping label\n", pstate.linestr(pstate.lineno), n)
			}
		} else if asNode(n.Left.Sym.Label) == &pstate.looping {
			if pstate.Debug['m'] > 2 {
				fmt.Printf("%v: %v looping label\n", pstate.linestr(pstate.lineno), n)
			}
			e.loopdepth++
		}

		// See case OLABEL in escloopdepth above
		// else if(n.Left.Sym.Label == nil)
		//	fatal("escape analysis missed or messed up a label: %+N", n);

		n.Left.Sym.Label = nil

	case ORANGE:
		if n.List.Len() >= 2 {
			// Everything but fixed array is a dereference.

			// If fixed array is really the address of fixed array,
			// it is also a dereference, because it is implicitly
			// dereferenced (see #12588)
			if n.Type.IsArray() &&
				!(n.Right.Type.IsPtr() && pstate.eqtype(n.Right.Type.Elem(pstate.types), n.Type)) {
				e.escassignWhyWhere(pstate, n.List.Second(), n.Right, "range", n)
			} else {
				e.escassignDereference(pstate, n.List.Second(), n.Right, e.stepAssignWhere(pstate, n.List.Second(), n.Right, "range-deref", n))
			}
		}

	case OSWITCH:
		if n.Left != nil && n.Left.Op == OTYPESW {
			for _, cas := range n.List.Slice() {
				// cases
				// n.Left.Right is the argument of the .(type),
				// it.N().Rlist is the variable per case
				if cas.Rlist.Len() != 0 {
					e.escassignWhyWhere(pstate, cas.Rlist.First(), n.Left.Right, "switch case", n)
				}
			}
		}

	// Filter out the following special case.
	//
	//	func (b *Buffer) Foo() {
	//		n, m := ...
	//		b.buf = b.buf[n:m]
	//	}
	//
	// This assignment is a no-op for escape analysis,
	// it does not store any new pointers into b that were not already there.
	// However, without this special case b will escape, because we assign to OIND/ODOTPTR.
	case OAS, OASOP:
		if (n.Left.Op == OIND || n.Left.Op == ODOTPTR) && n.Left.Left.Op == ONAME && // dst is ONAME dereference
			(n.Right.Op == OSLICE || n.Right.Op == OSLICE3 || n.Right.Op == OSLICESTR) && // src is slice operation
			(n.Right.Left.Op == OIND || n.Right.Left.Op == ODOTPTR) && n.Right.Left.Left.Op == ONAME && // slice is applied to ONAME dereference
			n.Left.Left == n.Right.Left.Left { // dst and src reference the same base ONAME

			// Here we also assume that the statement will not contain calls,
			// that is, that order will move any calls to init.
			// Otherwise base ONAME value could change between the moments
			// when we evaluate it for dst and for src.
			//
			// Note, this optimization does not apply to OSLICEARR,
			// because it does introduce a new pointer into b that was not already there
			// (pointer to b itself). After such assignment, if b contents escape,
			// b escapes as well. If we ignore such OSLICEARR, we will conclude
			// that b does not escape when b contents do.
			if pstate.Debug['m'] != 0 {
				pstate.Warnl(n.Pos, "%v ignoring self-assignment to %S", e.curfnSym(pstate, n), n.Left)
			}

			break
		}

		e.escassign(pstate, n.Left, n.Right, e.stepAssignWhere(pstate, nil, nil, "", n))

	case OAS2: // x,y = a,b
		if n.List.Len() == n.Rlist.Len() {
			rs := n.Rlist.Slice()
			for i, n := range n.List.Slice() {
				e.escassignWhyWhere(pstate, n, rs[i], "assign-pair", n)
			}
		}

	case OAS2RECV: // v, ok = <-ch
		e.escassignWhyWhere(pstate, n.List.First(), n.Rlist.First(), "assign-pair-receive", n)
	case OAS2MAPR: // v, ok = m[k]
		e.escassignWhyWhere(pstate, n.List.First(), n.Rlist.First(), "assign-pair-mapr", n)
	case OAS2DOTTYPE: // v, ok = x.(type)
		e.escassignWhyWhere(pstate, n.List.First(), n.Rlist.First(), "assign-pair-dot-type", n)

	case OSEND: // ch <- x
		e.escassignSinkWhy(pstate, n, n.Right, "send")

	case ODEFER:
		if e.loopdepth == 1 { // top level
			break
		}
		// arguments leak out of scope
		// TODO: leak to a dummy node instead
		// defer f(x) - f and x escape
		e.escassignSinkWhy(pstate, n, n.Left.Left, "defer func")
		e.escassignSinkWhy(pstate, n, n.Left.Right, "defer func ...") // ODDDARG for call
		for _, arg := range n.Left.List.Slice() {
			e.escassignSinkWhy(pstate, n, arg, "defer func arg")
		}

	case OPROC:
		// go f(x) - f and x escape
		e.escassignSinkWhy(pstate, n, n.Left.Left, "go func")
		e.escassignSinkWhy(pstate, n, n.Left.Right, "go func ...") // ODDDARG for call
		for _, arg := range n.Left.List.Slice() {
			e.escassignSinkWhy(pstate, n, arg, "go func arg")
		}

	case OCALLMETH, OCALLFUNC, OCALLINTER:
		e.esccall(pstate, n, parent)

	// esccall already done on n.Rlist.First(). tie it's Retval to n.List
	case OAS2FUNC: // x,y = f()
		rs := e.nodeEscState(pstate, n.Rlist.First()).Retval.Slice()
		for i, n := range n.List.Slice() {
			if i >= len(rs) {
				break
			}
			e.escassignWhyWhere(pstate, n, rs[i], "assign-pair-func-call", n)
		}
		if n.List.Len() != len(rs) {
			pstate.Fatalf("esc oas2func")
		}

	case ORETURN:
		retList := n.List
		if retList.Len() == 1 && pstate.Curfn.Type.NumResults(pstate.types) > 1 {
			// OAS2FUNC in disguise
			// esccall already done on n.List.First()
			// tie e.nodeEscState(n.List.First()).Retval to Curfn.Func.Dcl PPARAMOUT's
			retList = e.nodeEscState(pstate, n.List.First()).Retval
		}

		i := 0
		for _, lrn := range pstate.Curfn.Func.Dcl {
			if i >= retList.Len() {
				break
			}
			if lrn.Op != ONAME || lrn.Class() != PPARAMOUT {
				continue
			}
			e.escassignWhyWhere(pstate, lrn, retList.Index(i), "return", n)
			i++
		}

		if i < retList.Len() {
			pstate.Fatalf("esc return list")
		}

	// Argument could leak through recover.
	case OPANIC:
		e.escassignSinkWhy(pstate, n, n.Left, "panic")

	case OAPPEND:
		if !n.Isddd() {
			for _, nn := range n.List.Slice()[1:] {
				e.escassignSinkWhy(pstate, n, nn, "appended to slice") // lose track of assign to dereference
			}
		} else {
			// append(slice1, slice2...) -- slice2 itself does not escape, but contents do.
			slice2 := n.List.Second()
			e.escassignDereference(pstate, &e.theSink, slice2, e.stepAssignWhere(pstate, n, slice2, "appended slice...", n)) // lose track of assign of dereference
			if pstate.Debug['m'] > 3 {
				pstate.Warnl(n.Pos, "%v special treatment of append(slice1, slice2...) %S", e.curfnSym(pstate, n), n)
			}
		}
		e.escassignDereference(pstate, &e.theSink, n.List.First(), e.stepAssignWhere(pstate, n, n.List.First(), "appendee slice", n)) // The original elements are now leaked, too

	case OCOPY:
		e.escassignDereference(pstate, &e.theSink, n.Right, e.stepAssignWhere(pstate, n, n.Right, "copied slice", n)) // lose track of assign of dereference

	case OCONV, OCONVNOP:
		e.escassignWhyWhere(pstate, n, n.Left, "converted", n)

	case OCONVIFACE:
		e.track(pstate, n)
		e.escassignWhyWhere(pstate, n, n.Left, "interface-converted", n)

	case OARRAYLIT:
		// Link values to array
		for _, elt := range n.List.Slice() {
			if elt.Op == OKEY {
				elt = elt.Right
			}
			e.escassign(pstate, n, elt, e.stepAssignWhere(pstate, n, elt, "array literal element", n))
		}

	case OSLICELIT:
		// Slice is not leaked until proven otherwise
		e.track(pstate, n)
		// Link values to slice
		for _, elt := range n.List.Slice() {
			if elt.Op == OKEY {
				elt = elt.Right
			}
			e.escassign(pstate, n, elt, e.stepAssignWhere(pstate, n, elt, "slice literal element", n))
		}

	// Link values to struct.
	case OSTRUCTLIT:
		for _, elt := range n.List.Slice() {
			e.escassignWhyWhere(pstate, n, elt.Left, "struct literal element", n)
		}

	case OPTRLIT:
		e.track(pstate, n)

		// Link OSTRUCTLIT to OPTRLIT; if OPTRLIT escapes, OSTRUCTLIT elements do too.
		e.escassignWhyWhere(pstate, n, n.Left, "pointer literal [assign]", n)

	case OCALLPART:
		e.track(pstate, n)

		// Contents make it to memory, lose track.
		e.escassignSinkWhy(pstate, n, n.Left, "call part")

	case OMAPLIT:
		e.track(pstate, n)
		// Keys and values make it to memory, lose track.
		for _, elt := range n.List.Slice() {
			e.escassignSinkWhy(pstate, n, elt.Left, "map literal key")
			e.escassignSinkWhy(pstate, n, elt.Right, "map literal value")
		}

	case OCLOSURE:
		// Link addresses of captured variables to closure.
		for _, v := range n.Func.Closure.Func.Cvars.Slice() {
			if v.Op == OXXX { // unnamed out argument; see dcl.go:/^funcargs
				continue
			}
			a := v.Name.Defn
			if !v.Name.Byval() {
				a = pstate.nod(OADDR, a, nil)
				a.Pos = v.Pos
				e.nodeEscState(pstate, a).Loopdepth = e.loopdepth
				a = pstate.typecheck(a, Erv)
			}

			e.escassignWhyWhere(pstate, n, a, "captured by a closure", n)
		}
		fallthrough

	case OMAKECHAN,
		OMAKEMAP,
		OMAKESLICE,
		ONEW,
		OARRAYRUNESTR,
		OARRAYBYTESTR,
		OSTRARRAYRUNE,
		OSTRARRAYBYTE,
		ORUNESTR:
		e.track(pstate, n)

	case OADDSTR:
		e.track(pstate, n)
	// Arguments of OADDSTR do not escape.

	case OADDR:
		// current loop depth is an upper bound on actual loop depth
		// of addressed value.
		e.track(pstate, n)

		// for &x, use loop depth of x if known.
		// it should always be known, but if not, be conservative
		// and keep the current loop depth.
		if n.Left.Op == ONAME {
			switch n.Left.Class() {
			// PPARAM is loop depth 1 always.
			// PPARAMOUT is loop depth 0 for writes
			// but considered loop depth 1 for address-of,
			// so that writing the address of one result
			// to another (or the same) result makes the
			// first result move to the heap.
			case PPARAM, PPARAMOUT:
				nE := e.nodeEscState(pstate, n)
				nE.Loopdepth = 1
				break opSwitch
			}
		}
		nE := e.nodeEscState(pstate, n)
		leftE := e.nodeEscState(pstate, n.Left)
		if leftE.Loopdepth != 0 {
			nE.Loopdepth = leftE.Loopdepth
		}

	case ODOT,
		ODOTPTR,
		OINDEX:
		// Propagate the loopdepth of t to t.field.
		if n.Left.Op != OLITERAL { // OLITERAL node doesn't have esc state
			e.nodeEscState(pstate, n).Loopdepth = e.nodeEscState(pstate, n.Left).Loopdepth
		}
	}

	pstate.lineno = lno
}

// escassignWhyWhere bundles a common case of
// escassign(e, dst, src, e.stepAssignWhere(dst, src, reason, where))
func (e *EscState) escassignWhyWhere(pstate *PackageState, dst, src *Node, reason string, where *Node) {
	var step *EscStep
	if pstate.Debug['m'] != 0 {
		step = e.stepAssignWhere(pstate, dst, src, reason, where)
	}
	e.escassign(pstate, dst, src, step)
}

// escassignSinkWhy bundles a common case of
// escassign(e, &e.theSink, src, e.stepAssign(nil, dst, src, reason))
func (e *EscState) escassignSinkWhy(pstate *PackageState, dst, src *Node, reason string) {
	var step *EscStep
	if pstate.Debug['m'] != 0 {
		step = e.stepAssign(pstate, nil, dst, src, reason)
	}
	e.escassign(pstate, &e.theSink, src, step)
}

// escassignSinkWhyWhere is escassignSinkWhy but includes a call site
// for accurate location reporting.
func (e *EscState) escassignSinkWhyWhere(pstate *PackageState, dst, src *Node, reason string, call *Node) {
	var step *EscStep
	if pstate.Debug['m'] != 0 {
		step = e.stepAssignWhere(pstate, dst, src, reason, call)
	}
	e.escassign(pstate, &e.theSink, src, step)
}

// Assert that expr somehow gets assigned to dst, if non nil.  for
// dst==nil, any name node expr still must be marked as being
// evaluated in curfn.	For expr==nil, dst must still be examined for
// evaluations inside it (e.g *f(x) = y)
func (e *EscState) escassign(pstate *PackageState, dst, src *Node, step *EscStep) {
	if dst.isBlank() || dst == nil || src == nil || src.Op == ONONAME || src.Op == OXXX {
		return
	}

	if pstate.Debug['m'] > 2 {
		fmt.Printf("%v:[%d] %v escassign: %S(%0j)[%v] = %S(%0j)[%v]\n",
			pstate.linestr(pstate.lineno), e.loopdepth, funcSym(pstate.Curfn),
			dst, dst, dst.Op,
			src, src, src.Op)
	}

	pstate.setlineno(dst)

	originalDst := dst
	dstwhy := "assigned"

	// Analyze lhs of assignment.
	// Replace dst with &e.theSink if we can't track it.
	switch dst.Op {
	default:
		Dump("dst", dst)
		pstate.Fatalf("escassign: unexpected dst")

	case OARRAYLIT,
		OSLICELIT,
		OCLOSURE,
		OCONV,
		OCONVIFACE,
		OCONVNOP,
		OMAPLIT,
		OSTRUCTLIT,
		OPTRLIT,
		ODDDARG,
		OCALLPART:

	case ONAME:
		if dst.Class() == PEXTERN {
			dstwhy = "assigned to top level variable"
			dst = &e.theSink
		}

	case ODOT: // treat "dst.x = src" as "dst = src"
		e.escassign(pstate, dst.Left, src, e.stepAssign(pstate, step, originalDst, src, "dot-equals"))
		return

	case OINDEX:
		if dst.Left.Type.IsArray() {
			e.escassign(pstate, dst.Left, src, e.stepAssign(pstate, step, originalDst, src, "array-element-equals"))
			return
		}

		dstwhy = "slice-element-equals"
		dst = &e.theSink // lose track of dereference

	case OIND:
		dstwhy = "star-equals"
		dst = &e.theSink // lose track of dereference

	case ODOTPTR:
		dstwhy = "star-dot-equals"
		dst = &e.theSink // lose track of dereference

	// lose track of key and value
	case OINDEXMAP:
		e.escassign(pstate, &e.theSink, dst.Right, e.stepAssign(pstate, nil, originalDst, src, "key of map put"))
		dstwhy = "value of map put"
		dst = &e.theSink
	}

	lno := pstate.setlineno(src)
	e.pdepth++

	switch src.Op {
	case OADDR, // dst = &x
		OIND,    // dst = *x
		ODOTPTR, // dst = (*x).f
		ONAME,
		ODDDARG,
		OPTRLIT,
		OARRAYLIT,
		OSLICELIT,
		OMAPLIT,
		OSTRUCTLIT,
		OMAKECHAN,
		OMAKEMAP,
		OMAKESLICE,
		OARRAYRUNESTR,
		OARRAYBYTESTR,
		OSTRARRAYRUNE,
		OSTRARRAYBYTE,
		OADDSTR,
		ONEW,
		OCALLPART,
		ORUNESTR,
		OCONVIFACE:
		e.escflows(pstate, dst, src, e.stepAssign(pstate, step, originalDst, src, dstwhy))

	case OCLOSURE:
		// OCLOSURE is lowered to OPTRLIT,
		// insert OADDR to account for the additional indirection.
		a := pstate.nod(OADDR, src, nil)
		a.Pos = src.Pos
		e.nodeEscState(pstate, a).Loopdepth = e.nodeEscState(pstate, src).Loopdepth
		a.Type = pstate.types.NewPtr(src.Type)
		e.escflows(pstate, dst, a, e.stepAssign(pstate, nil, originalDst, src, dstwhy))

	// Flowing multiple returns to a single dst happens when
	// analyzing "go f(g())": here g() flows to sink (issue 4529).
	case OCALLMETH, OCALLFUNC, OCALLINTER:
		for _, n := range e.nodeEscState(pstate, src).Retval.Slice() {
			e.escflows(pstate, dst, n, e.stepAssign(pstate, nil, originalDst, n, dstwhy))
		}

	// A non-pointer escaping from a struct does not concern us.
	case ODOT:
		if src.Type != nil && !pstate.types.Haspointers(src.Type) {
			break
		}
		fallthrough

	// Conversions, field access, slice all preserve the input value.
	case OCONV,
		OCONVNOP,
		ODOTMETH,
		// treat recv.meth as a value with recv in it, only happens in ODEFER and OPROC
		// iface.method already leaks iface in esccall, no need to put in extra ODOTINTER edge here
		OSLICE,
		OSLICE3,
		OSLICEARR,
		OSLICE3ARR,
		OSLICESTR:
		// Conversions, field access, slice all preserve the input value.
		e.escassign(pstate, dst, src.Left, e.stepAssign(pstate, step, originalDst, src, dstwhy))

	case ODOTTYPE,
		ODOTTYPE2:
		if src.Type != nil && !pstate.types.Haspointers(src.Type) {
			break
		}
		e.escassign(pstate, dst, src.Left, e.stepAssign(pstate, step, originalDst, src, dstwhy))

	case OAPPEND:
		// Append returns first argument.
		// Subsequent arguments are already leaked because they are operands to append.
		e.escassign(pstate, dst, src.List.First(), e.stepAssign(pstate, step, dst, src.List.First(), dstwhy))

	case OINDEX:
		// Index of array preserves input value.
		if src.Left.Type.IsArray() {
			e.escassign(pstate, dst, src.Left, e.stepAssign(pstate, step, originalDst, src, dstwhy))
		} else {
			e.escflows(pstate, dst, src, e.stepAssign(pstate, step, originalDst, src, dstwhy))
		}

	// Might be pointer arithmetic, in which case
	// the operands flow into the result.
	// TODO(rsc): Decide what the story is here. This is unsettling.
	case OADD,
		OSUB,
		OOR,
		OXOR,
		OMUL,
		ODIV,
		OMOD,
		OLSH,
		ORSH,
		OAND,
		OANDNOT,
		OPLUS,
		OMINUS,
		OCOM:
		e.escassign(pstate, dst, src.Left, e.stepAssign(pstate, step, originalDst, src, dstwhy))

		e.escassign(pstate, dst, src.Right, e.stepAssign(pstate, step, originalDst, src, dstwhy))
	}

	e.pdepth--
	pstate.lineno = lno
}

// mktag returns the string representation for an escape analysis tag.
func (pstate *PackageState) mktag(mask int) string {
	switch mask & EscMask {
	case EscNone, EscReturn:
	default:
		pstate.Fatalf("escape mktag")
	}

	if mask < len(pstate.tags) && pstate.tags[mask] != "" {
		return pstate.tags[mask]
	}

	s := fmt.Sprintf("esc:0x%x", mask)
	if mask < len(pstate.tags) {
		pstate.tags[mask] = s
	}
	return s
}

// parsetag decodes an escape analysis tag and returns the esc value.
func parsetag(note string) uint16 {
	if !strings.HasPrefix(note, "esc:") {
		return EscUnknown
	}
	n, _ := strconv.ParseInt(note[4:], 0, 0)
	em := uint16(n)
	if em == 0 {
		return EscNone
	}
	return em
}

// describeEscape returns a string describing the escape tag.
// The result is either one of {EscUnknown, EscNone, EscHeap} which all have no further annotation
// or a description of parameter flow, which takes the form of an optional "contentToHeap"
// indicating that the content of this parameter is leaked to the heap, followed by a sequence
// of level encodings separated by spaces, one for each parameter, where _ means no flow,
// = means direct flow, and N asterisks (*) encodes content (obtained by indirection) flow.
// e.g., "contentToHeap _ =" means that a parameter's content (one or more dereferences)
// escapes to the heap, the parameter does not leak to the first output, but does leak directly
// to the second output (and if there are more than two outputs, there is no flow to those.)
func describeEscape(em uint16) string {
	var s string
	switch em & EscMask {
	case EscUnknown:
		s = "EscUnknown"
	case EscNone:
		s = "EscNone"
	case EscHeap:
		s = "EscHeap"
	case EscReturn:
		s = "EscReturn"
	}
	if em&EscContentEscapes != 0 {
		if s != "" {
			s += " "
		}
		s += "contentToHeap"
	}
	for em >>= EscReturnBits; em != 0; em = em >> bitsPerOutputInTag {
		// See encoding description above
		if s != "" {
			s += " "
		}
		switch embits := em & bitsMaskForTag; embits {
		case 0:
			s += "_"
		case 1:
			s += "="
		default:
			for i := uint16(0); i < embits-1; i++ {
				s += "*"
			}
		}

	}
	return s
}

// escassignfromtag models the input-to-output assignment flow of one of a function
// calls arguments, where the flow is encoded in "note".
func (e *EscState) escassignfromtag(pstate *PackageState, note string, dsts Nodes, src, call *Node) uint16 {
	em := parsetag(note)
	if src.Op == OLITERAL {
		return em
	}

	if pstate.Debug['m'] > 3 {
		fmt.Printf("%v::assignfromtag:: src=%S, em=%s\n",
			pstate.linestr(pstate.lineno), src, describeEscape(em))
	}

	if em == EscUnknown {
		e.escassignSinkWhyWhere(pstate, src, src, "passed to call[argument escapes]", call)
		return em
	}

	if em == EscNone {
		return em
	}

	// If content inside parameter (reached via indirection)
	// escapes to heap, mark as such.
	if em&EscContentEscapes != 0 {
		e.escassign(pstate, &e.theSink, e.addDereference(pstate, src), e.stepAssignWhere(pstate, src, src, "passed to call[argument content escapes]", call))
	}

	em0 := em
	dstsi := 0
	for em >>= EscReturnBits; em != 0 && dstsi < dsts.Len(); em = em >> bitsPerOutputInTag {
		// Prefer the lowest-level path to the reference (for escape purposes).
		// Two-bit encoding (for example. 1, 3, and 4 bits are other options)
		//  01 = 0-level
		//  10 = 1-level, (content escapes),
		//  11 = 2-level, (content of content escapes),
		embits := em & bitsMaskForTag
		if embits > 0 {
			n := src
			for i := uint16(0); i < embits-1; i++ {
				n = e.addDereference(pstate, n) // encode level>0 as indirections
			}
			e.escassign(pstate, dsts.Index(dstsi), n, e.stepAssignWhere(pstate, dsts.Index(dstsi), src, "passed-to-and-returned-from-call", call))
		}
		dstsi++
	}
	// If there are too many outputs to fit in the tag,
	// that is handled at the encoding end as EscHeap,
	// so there is no need to check here.

	if em != 0 && dstsi >= dsts.Len() {
		pstate.Fatalf("corrupt esc tag %q or messed up escretval list\n", note)
	}
	return em0
}

func (e *EscState) escassignDereference(pstate *PackageState, dst *Node, src *Node, step *EscStep) {
	if src.Op == OLITERAL {
		return
	}
	e.escassign(pstate, dst, e.addDereference(pstate, src), step)
}

// addDereference constructs a suitable OIND note applied to src.
// Because this is for purposes of escape accounting, not execution,
// some semantically dubious node combinations are (currently) possible.
func (e *EscState) addDereference(pstate *PackageState, n *Node) *Node {
	ind := pstate.nod(OIND, n, nil)
	e.nodeEscState(pstate, ind).Loopdepth = e.nodeEscState(pstate, n).Loopdepth
	ind.Pos = n.Pos
	t := n.Type
	if t.IsKind(pstate.types.Tptr) || t.IsSlice() {
		// This should model our own sloppy use of OIND to encode
		// decreasing levels of indirection; i.e., "indirecting" a slice
		// yields the type of an element.
		t = t.Elem(pstate.types)
	} else if t.IsString() {
		t = pstate.types.Types[TUINT8]
	}
	ind.Type = t
	return ind
}

// escNoteOutputParamFlow encodes maxEncodedLevel/.../1/0-level flow to the vargen'th parameter.
// Levels greater than maxEncodedLevel are replaced with maxEncodedLevel.
// If the encoding cannot describe the modified input level and output number, then EscHeap is returned.
func (pstate *PackageState) escNoteOutputParamFlow(e uint16, vargen int32, level Level) uint16 {
	// Flow+level is encoded in two bits.
	// 00 = not flow, xx = level+1 for 0 <= level <= maxEncodedLevel
	// 16 bits for Esc allows 6x2bits or 4x3bits or 3x4bits if additional information would be useful.
	if level.int() <= 0 && level.guaranteedDereference() > 0 {
		return pstate.escMax(e|EscContentEscapes, EscNone) // At least one deref, thus only content.
	}
	if level.int() < 0 {
		return EscHeap
	}
	if level.int() > maxEncodedLevel {
		// Cannot encode larger values than maxEncodedLevel.
		level = levelFrom(maxEncodedLevel)
	}
	encoded := uint16(level.int() + 1)

	shift := uint(bitsPerOutputInTag*(vargen-1) + EscReturnBits)
	old := (e >> shift) & bitsMaskForTag
	if old == 0 || encoded != 0 && encoded < old {
		old = encoded
	}

	encodedFlow := old << shift
	if (encodedFlow>>shift)&bitsMaskForTag != old {
		// Encoding failure defaults to heap.
		return EscHeap
	}

	return (e &^ (bitsMaskForTag << shift)) | encodedFlow
}

func (e *EscState) initEscRetval(pstate *PackageState, call *Node, fntype *types.Type) {
	cE := e.nodeEscState(pstate, call)
	cE.Retval.Set(nil) // Suspect this is not nil for indirect calls.
	for i, f := range fntype.Results(pstate.types).Fields(pstate.types).Slice() {
		buf := fmt.Sprintf(".out%d", i)
		ret := pstate.newname(pstate.lookup(buf))
		ret.SetAddable(false) // TODO(mdempsky): Seems suspicious.
		ret.Type = f.Type
		ret.SetClass(PAUTO)
		ret.Name.Curfn = pstate.Curfn
		e.nodeEscState(pstate, ret).Loopdepth = e.loopdepth
		ret.Name.SetUsed(true)
		ret.Pos = call.Pos
		cE.Retval.Append(ret)
	}
}

// This is a bit messier than fortunate, pulled out of esc's big
// switch for clarity. We either have the paramnodes, which may be
// connected to other things through flows or we have the parameter type
// nodes, which may be marked "noescape". Navigating the ast is slightly
// different for methods vs plain functions and for imported vs
// this-package
func (e *EscState) esccall(pstate *PackageState, call *Node, parent *Node) {
	var fntype *types.Type
	var indirect bool
	var fn *Node
	switch call.Op {
	default:
		pstate.Fatalf("esccall")

	case OCALLFUNC:
		fn = call.Left
		fntype = fn.Type
		indirect = fn.Op != ONAME || fn.Class() != PFUNC

	case OCALLMETH:
		fn = asNode(call.Left.Sym.Def)
		if fn != nil {
			fntype = fn.Type
		} else {
			fntype = call.Left.Type
		}

	case OCALLINTER:
		fntype = call.Left.Type
		indirect = true
	}

	argList := call.List
	if argList.Len() == 1 {
		arg := argList.First()
		if arg.Type.IsFuncArgStruct() { // f(g())
			argList = e.nodeEscState(pstate, arg).Retval
		}
	}

	args := argList.Slice()

	if indirect {
		// We know nothing!
		// Leak all the parameters
		for _, arg := range args {
			e.escassignSinkWhy(pstate, call, arg, "parameter to indirect call")
			if pstate.Debug['m'] > 3 {
				fmt.Printf("%v::esccall:: indirect call <- %S, untracked\n", pstate.linestr(pstate.lineno), arg)
			}
		}
		// Set up bogus outputs
		e.initEscRetval(pstate, call, fntype)
		// If there is a receiver, it also leaks to heap.
		if call.Op != OCALLFUNC {
			rf := fntype.Recv(pstate.types)
			r := call.Left.Left
			if pstate.types.Haspointers(rf.Type) {
				e.escassignSinkWhy(pstate, call, r, "receiver in indirect call")
			}
		} else { // indirect and OCALLFUNC = could be captured variables, too. (#14409)
			rets := e.nodeEscState(pstate, call).Retval.Slice()
			for _, ret := range rets {
				e.escassignDereference(pstate, ret, fn, e.stepAssignWhere(pstate, ret, fn, "captured by called closure", call))
			}
		}
		return
	}

	cE := e.nodeEscState(pstate, call)
	if fn != nil && fn.Op == ONAME && fn.Class() == PFUNC &&
		fn.Name.Defn != nil && fn.Name.Defn.Nbody.Len() != 0 && fn.Name.Param.Ntype != nil && fn.Name.Defn.Esc < EscFuncTagged {
		if pstate.Debug['m'] > 3 {
			fmt.Printf("%v::esccall:: %S in recursive group\n", pstate.linestr(pstate.lineno), call)
		}

		// function in same mutually recursive group. Incorporate into flow graph.
		//		print("esc local fn: %N\n", fn.Func.Ntype);
		if fn.Name.Defn.Esc == EscFuncUnknown || cE.Retval.Len() != 0 {
			pstate.Fatalf("graph inconsistency")
		}

		sawRcvr := false
		for _, n := range fn.Name.Defn.Func.Dcl {
			switch n.Class() {
			case PPARAM:
				if call.Op != OCALLFUNC && !sawRcvr {
					e.escassignWhyWhere(pstate, n, call.Left.Left, "call receiver", call)
					sawRcvr = true
					continue
				}
				if len(args) == 0 {
					continue
				}
				arg := args[0]
				if n.Isddd() && !call.Isddd() {
					// Introduce ODDDARG node to represent ... allocation.
					arg = pstate.nod(ODDDARG, nil, nil)
					arr := pstate.types.NewArray(n.Type.Elem(pstate.types), int64(len(args)))
					arg.Type = pstate.types.NewPtr(arr) // make pointer so it will be tracked
					arg.Pos = call.Pos
					e.track(pstate, arg)
					call.Right = arg
				}
				e.escassignWhyWhere(pstate, n, arg, "arg to recursive call", call) // TODO this message needs help.
				if arg == args[0] {
					args = args[1:]
					continue
				}
				// "..." arguments are untracked
				for _, a := range args {
					if pstate.Debug['m'] > 3 {
						fmt.Printf("%v::esccall:: ... <- %S, untracked\n", pstate.linestr(pstate.lineno), a)
					}
					e.escassignSinkWhyWhere(pstate, arg, a, "... arg to recursive call", call)
				}
				// No more PPARAM processing, but keep
				// going for PPARAMOUT.
				args = nil

			case PPARAMOUT:
				cE.Retval.Append(n)
			}
		}

		return
	}

	// Imported or completely analyzed function. Use the escape tags.
	if cE.Retval.Len() != 0 {
		pstate.Fatalf("esc already decorated call %+v\n", call)
	}

	if pstate.Debug['m'] > 3 {
		fmt.Printf("%v::esccall:: %S not recursive\n", pstate.linestr(pstate.lineno), call)
	}

	// set up out list on this call node with dummy auto ONAMES in the current (calling) function.
	e.initEscRetval(pstate, call, fntype)

	//	print("esc analyzed fn: %#N (%+T) returning (%+H)\n", fn, fntype, e.nodeEscState(call).Retval);

	// Receiver.
	if call.Op != OCALLFUNC {
		rf := fntype.Recv(pstate.types)
		r := call.Left.Left
		if pstate.types.Haspointers(rf.Type) {
			e.escassignfromtag(pstate, rf.Note, cE.Retval, r, call)
		}
	}

	for i, param := range fntype.Params(pstate.types).FieldSlice(pstate.types) {
		note := param.Note
		var arg *Node
		if param.Isddd() && !call.Isddd() {
			rest := args[i:]
			if len(rest) == 0 {
				break
			}

			// Introduce ODDDARG node to represent ... allocation.
			arg = pstate.nod(ODDDARG, nil, nil)
			arg.Pos = call.Pos
			arr := pstate.types.NewArray(param.Type.Elem(pstate.types), int64(len(rest)))
			arg.Type = pstate.types.NewPtr(arr) // make pointer so it will be tracked
			e.track(pstate, arg)
			call.Right = arg

			// Store arguments into slice for ... arg.
			for _, a := range rest {
				if pstate.Debug['m'] > 3 {
					fmt.Printf("%v::esccall:: ... <- %S\n", pstate.linestr(pstate.lineno), a)
				}
				if note == uintptrEscapesTag {
					e.escassignSinkWhyWhere(pstate, arg, a, "arg to uintptrescapes ...", call)
				} else {
					e.escassignWhyWhere(pstate, arg, a, "arg to ...", call)
				}
			}
		} else {
			arg = args[i]
			if note == uintptrEscapesTag {
				e.escassignSinkWhy(pstate, arg, arg, "escaping uintptr")
			}
		}

		if pstate.types.Haspointers(param.Type) && e.escassignfromtag(pstate, note, cE.Retval, arg, call)&EscMask == EscNone && parent.Op != ODEFER && parent.Op != OPROC {
			a := arg
			for a.Op == OCONVNOP {
				a = a.Left
			}
			switch a.Op {
			// The callee has already been analyzed, so its arguments have esc tags.
			// The argument is marked as not escaping at all.
			// Record that fact so that any temporary used for
			// synthesizing this expression can be reclaimed when
			// the function returns.
			// This 'noescape' is even stronger than the usual esc == EscNone.
			// arg.Esc == EscNone means that arg does not escape the current function.
			// arg.SetNoescape(true) here means that arg does not escape this statement
			// in the current function.
			case OCALLPART, OCLOSURE, ODDDARG, OARRAYLIT, OSLICELIT, OPTRLIT, OSTRUCTLIT:
				a.SetNoescape(true)
			}
		}
	}
}

// escflows records the link src->dst in dst, throwing out some quick wins,
// and also ensuring that dst is noted as a flow destination.
func (e *EscState) escflows(pstate *PackageState, dst, src *Node, why *EscStep) {
	if dst == nil || src == nil || dst == src {
		return
	}

	// Don't bother building a graph for scalars.
	if src.Type != nil && !pstate.types.Haspointers(src.Type) && !pstate.isReflectHeaderDataField(src) {
		if pstate.Debug['m'] > 3 {
			fmt.Printf("%v::NOT flows:: %S <- %S\n", pstate.linestr(pstate.lineno), dst, src)
		}
		return
	}

	if pstate.Debug['m'] > 3 {
		fmt.Printf("%v::flows:: %S <- %S\n", pstate.linestr(pstate.lineno), dst, src)
	}

	dstE := e.nodeEscState(pstate, dst)
	if len(dstE.Flowsrc) == 0 {
		e.dsts = append(e.dsts, dst)
		e.dstcount++
	}

	e.edgecount++

	if why == nil {
		dstE.Flowsrc = append(dstE.Flowsrc, EscStep{src: src})
	} else {
		starwhy := *why
		starwhy.src = src // TODO: need to reconcile this w/ needs of explanations.
		dstE.Flowsrc = append(dstE.Flowsrc, starwhy)
	}
}

// Whenever we hit a reference node, the level goes up by one, and whenever
// we hit an OADDR, the level goes down by one. as long as we're on a level > 0
// finding an OADDR just means we're following the upstream of a dereference,
// so this address doesn't leak (yet).
// If level == 0, it means the /value/ of this node can reach the root of this flood.
// so if this node is an OADDR, its argument should be marked as escaping iff
// its currfn/e.loopdepth are different from the flood's root.
// Once an object has been moved to the heap, all of its upstream should be considered
// escaping to the global scope.
func (e *EscState) escflood(pstate *PackageState, dst *Node) {
	switch dst.Op {
	case ONAME, OCLOSURE:
	default:
		return
	}

	dstE := e.nodeEscState(pstate, dst)
	if pstate.Debug['m'] > 2 {
		fmt.Printf("\nescflood:%d: dst %S scope:%v[%d]\n", e.walkgen, dst, e.curfnSym(pstate, dst), dstE.Loopdepth)
	}

	for i := range dstE.Flowsrc {
		e.walkgen++
		s := &dstE.Flowsrc[i]
		s.parent = nil
		e.escwalk(pstate, levelFrom(0), dst, s.src, s)
	}
}

// funcOutputAndInput reports whether dst and src correspond to output and input parameters of the same function.
func funcOutputAndInput(dst, src *Node) bool {
	// Note if dst is marked as escaping, then "returned" is too weak.
	return dst.Op == ONAME && dst.Class() == PPARAMOUT &&
		src.Op == ONAME && src.Class() == PPARAM && src.Name.Curfn == dst.Name.Curfn
}

func (es *EscStep) describe(pstate *PackageState, src *Node) {
	if pstate.Debug['m'] < 2 {
		return
	}
	step0 := es
	for step := step0; step != nil && !step.busy; step = step.parent {
		// TODO: We get cycles. Trigger is i = &i (where var i interface{})
		step.busy = true
		// The trail is a little odd because of how the
		// graph is constructed.  The link to the current
		// Node is parent.src unless parent is nil in which
		// case it is step.dst.
		nextDest := step.parent
		dst := step.dst
		where := step.where
		if nextDest != nil {
			dst = nextDest.src
		}
		if where == nil {
			where = dst
		}
		pstate.Warnl(src.Pos, "\tfrom %v (%s) at %s", dst, step.why, where.Line(pstate))
	}
	for step := step0; step != nil && step.busy; step = step.parent {
		step.busy = false
	}
}

const NOTALOOPDEPTH = -1

func (e *EscState) escwalk(pstate *PackageState, level Level, dst *Node, src *Node, step *EscStep) {
	e.escwalkBody(pstate, level, dst, src, step, NOTALOOPDEPTH)
}

func (e *EscState) escwalkBody(pstate *PackageState, level Level, dst *Node, src *Node, step *EscStep, extraloopdepth int32) {
	if src.Op == OLITERAL {
		return
	}
	srcE := e.nodeEscState(pstate, src)
	if srcE.Walkgen == e.walkgen {
		// Esclevels are vectors, do not compare as integers,
		// and must use "min" of old and new to guarantee
		// convergence.
		level = level.min(srcE.Level)
		if level == srcE.Level {
			// Have we been here already with an extraloopdepth,
			// or is the extraloopdepth provided no improvement on
			// what's already been seen?
			if srcE.Maxextraloopdepth >= extraloopdepth || srcE.Loopdepth >= extraloopdepth {
				return
			}
			srcE.Maxextraloopdepth = extraloopdepth
		}
	} else { // srcE.Walkgen < e.walkgen -- first time, reset this.
		srcE.Maxextraloopdepth = NOTALOOPDEPTH
	}

	srcE.Walkgen = e.walkgen
	srcE.Level = level
	modSrcLoopdepth := srcE.Loopdepth

	if extraloopdepth > modSrcLoopdepth {
		modSrcLoopdepth = extraloopdepth
	}

	if pstate.Debug['m'] > 2 {
		fmt.Printf("escwalk: level:%d depth:%d %.*s op=%v %S(%0j) scope:%v[%d] extraloopdepth=%v\n",
			level, e.pdepth, e.pdepth, "\t\t\t\t\t\t\t\t\t\t", src.Op, src, src, e.curfnSym(pstate, src), srcE.Loopdepth, extraloopdepth)
	}

	e.pdepth++

	// Input parameter flowing to output parameter?
	var leaks bool
	var osrcesc uint16 // used to prevent duplicate error messages

	dstE := e.nodeEscState(pstate, dst)
	if funcOutputAndInput(dst, src) && src.Esc&EscMask < EscHeap && dst.Esc != EscHeap {
		// This case handles:
		// 1. return in
		// 2. return &in
		// 3. tmp := in; return &tmp
		// 4. return *in
		if pstate.Debug['m'] != 0 {
			if pstate.Debug['m'] <= 2 {
				pstate.Warnl(src.Pos, "leaking param: %S to result %v level=%v", src, dst.Sym, level.int())
				step.describe(pstate, src)
			} else {
				pstate.Warnl(src.Pos, "leaking param: %S to result %v level=%v", src, dst.Sym, level)
			}
		}
		if src.Esc&EscMask != EscReturn {
			src.Esc = EscReturn | src.Esc&EscContentEscapes
		}
		src.Esc = pstate.escNoteOutputParamFlow(src.Esc, dst.Name.Vargen, level)
		goto recurse
	}

	// If parameter content escapes to heap, set EscContentEscapes
	// Note minor confusion around escape from pointer-to-struct vs escape from struct
	if dst.Esc == EscHeap &&
		src.Op == ONAME && src.Class() == PPARAM && src.Esc&EscMask < EscHeap &&
		level.int() > 0 {
		src.Esc = pstate.escMax(EscContentEscapes|src.Esc, EscNone)
		if pstate.Debug['m'] != 0 {
			pstate.Warnl(src.Pos, "mark escaped content: %S", src)
			step.describe(pstate, src)
		}
	}

	leaks = level.int() <= 0 && level.guaranteedDereference() <= 0 && dstE.Loopdepth < modSrcLoopdepth
	leaks = leaks || level.int() <= 0 && dst.Esc&EscMask == EscHeap

	osrcesc = src.Esc
	switch src.Op {
	case ONAME:
		if src.Class() == PPARAM && (leaks || dstE.Loopdepth < 0) && src.Esc&EscMask < EscHeap {
			if level.guaranteedDereference() > 0 {
				src.Esc = pstate.escMax(EscContentEscapes|src.Esc, EscNone)
				if pstate.Debug['m'] != 0 {
					if pstate.Debug['m'] <= 2 {
						if osrcesc != src.Esc {
							pstate.Warnl(src.Pos, "leaking param content: %S", src)
							step.describe(pstate, src)
						}
					} else {
						pstate.Warnl(src.Pos, "leaking param content: %S level=%v dst.eld=%v src.eld=%v dst=%S",
							src, level, dstE.Loopdepth, modSrcLoopdepth, dst)
					}
				}
			} else {
				src.Esc = EscHeap
				if pstate.Debug['m'] != 0 {
					if pstate.Debug['m'] <= 2 {
						pstate.Warnl(src.Pos, "leaking param: %S", src)
						step.describe(pstate, src)
					} else {
						pstate.Warnl(src.Pos, "leaking param: %S level=%v dst.eld=%v src.eld=%v dst=%S",
							src, level, dstE.Loopdepth, modSrcLoopdepth, dst)
					}
				}
			}
		}

		// Treat a captured closure variable as equivalent to the
		// original variable.
		if src.IsClosureVar() {
			if leaks && pstate.Debug['m'] != 0 {
				pstate.Warnl(src.Pos, "leaking closure reference %S", src)
				step.describe(pstate, src)
			}
			e.escwalk(pstate, level, dst, src.Name.Defn, e.stepWalk(pstate, dst, src.Name.Defn, "closure-var", step))
		}

	case OPTRLIT, OADDR:
		why := "pointer literal"
		if src.Op == OADDR {
			why = "address-of"
		}
		if leaks {
			src.Esc = EscHeap
			if pstate.Debug['m'] != 0 && osrcesc != src.Esc {
				p := src
				if p.Left.Op == OCLOSURE {
					p = p.Left // merely to satisfy error messages in tests
				}
				if pstate.Debug['m'] > 2 {
					pstate.Warnl(src.Pos, "%S escapes to heap, level=%v, dst=%v dst.eld=%v, src.eld=%v",
						p, level, dst, dstE.Loopdepth, modSrcLoopdepth)
				} else {
					pstate.Warnl(src.Pos, "%S escapes to heap", p)
					step.describe(pstate, src)
				}
			}
			pstate.addrescapes(src.Left)
			e.escwalkBody(pstate, level.dec(), dst, src.Left, e.stepWalk(pstate, dst, src.Left, why, step), modSrcLoopdepth)
			extraloopdepth = modSrcLoopdepth // passes to recursive case, seems likely a no-op
		} else {
			e.escwalk(pstate, level.dec(), dst, src.Left, e.stepWalk(pstate, dst, src.Left, why, step))
		}

	case OAPPEND:
		e.escwalk(pstate, level, dst, src.List.First(), e.stepWalk(pstate, dst, src.List.First(), "append-first-arg", step))

	case ODDDARG:
		if leaks {
			src.Esc = EscHeap
			if pstate.Debug['m'] != 0 && osrcesc != src.Esc {
				pstate.Warnl(src.Pos, "%S escapes to heap", src)
				step.describe(pstate, src)
			}
			extraloopdepth = modSrcLoopdepth
		}
		// similar to a slice arraylit and its args.
		level = level.dec()

	case OSLICELIT:
		for _, elt := range src.List.Slice() {
			if elt.Op == OKEY {
				elt = elt.Right
			}
			e.escwalk(pstate, level.dec(), dst, elt, e.stepWalk(pstate, dst, elt, "slice-literal-element", step))
		}

		fallthrough

	case OMAKECHAN,
		OMAKEMAP,
		OMAKESLICE,
		OARRAYRUNESTR,
		OARRAYBYTESTR,
		OSTRARRAYRUNE,
		OSTRARRAYBYTE,
		OADDSTR,
		OMAPLIT,
		ONEW,
		OCLOSURE,
		OCALLPART,
		ORUNESTR,
		OCONVIFACE:
		if leaks {
			src.Esc = EscHeap
			if pstate.Debug['m'] != 0 && osrcesc != src.Esc {
				pstate.Warnl(src.Pos, "%S escapes to heap", src)
				step.describe(pstate, src)
			}
			extraloopdepth = modSrcLoopdepth
		}

	case ODOT,
		ODOTTYPE:
		e.escwalk(pstate, level, dst, src.Left, e.stepWalk(pstate, dst, src.Left, "dot", step))

	case
		OSLICE,
		OSLICEARR,
		OSLICE3,
		OSLICE3ARR,
		OSLICESTR:
		e.escwalk(pstate, level, dst, src.Left, e.stepWalk(pstate, dst, src.Left, "slice", step))

	case OINDEX:
		if src.Left.Type.IsArray() {
			e.escwalk(pstate, level, dst, src.Left, e.stepWalk(pstate, dst, src.Left, "fixed-array-index-of", step))
			break
		}
		fallthrough

	case ODOTPTR:
		e.escwalk(pstate, level.inc(), dst, src.Left, e.stepWalk(pstate, dst, src.Left, "dot of pointer", step))
	case OINDEXMAP:
		e.escwalk(pstate, level.inc(), dst, src.Left, e.stepWalk(pstate, dst, src.Left, "map index", step))
	case OIND:
		e.escwalk(pstate, level.inc(), dst, src.Left, e.stepWalk(pstate, dst, src.Left, "indirection", step))

	// In this case a link went directly to a call, but should really go
	// to the dummy .outN outputs that were created for the call that
	// themselves link to the inputs with levels adjusted.
	// See e.g. #10466
	// This can only happen with functions returning a single result.
	case OCALLMETH, OCALLFUNC, OCALLINTER:
		if srcE.Retval.Len() != 0 {
			if pstate.Debug['m'] > 2 {
				fmt.Printf("%v:[%d] dst %S escwalk replace src: %S with %S\n",
					pstate.linestr(pstate.lineno), e.loopdepth,
					dst, src, srcE.Retval.First())
			}
			src = srcE.Retval.First()
			srcE = e.nodeEscState(pstate, src)
		}
	}

recurse:
	level = level.copy()

	for i := range srcE.Flowsrc {
		s := &srcE.Flowsrc[i]
		s.parent = step
		e.escwalkBody(pstate, level, dst, s.src, s, extraloopdepth)
		s.parent = nil
	}

	e.pdepth--
}

// addrescapes tags node n as having had its address taken
// by "increasing" the "value" of n.Esc to EscHeap.
// Storage is allocated as necessary to allow the address
// to be taken.
func (pstate *PackageState) addrescapes(n *Node) {
	switch n.Op {
	default:
	// Unexpected Op, probably due to a previous type error. Ignore.

	case OIND, ODOTPTR:
	// Nothing to do.

	case ONAME:
		if n == pstate.nodfp {
			break
		}

		// if this is a tmpname (PAUTO), it was tagged by tmpname as not escaping.
		// on PPARAM it means something different.
		if n.Class() == PAUTO && n.Esc == EscNever {
			break
		}

		// If a closure reference escapes, mark the outer variable as escaping.
		if n.IsClosureVar() {
			pstate.addrescapes(n.Name.Defn)
			break
		}

		if n.Class() != PPARAM && n.Class() != PPARAMOUT && n.Class() != PAUTO {
			break
		}

		// This is a plain parameter or local variable that needs to move to the heap,
		// but possibly for the function outside the one we're compiling.
		// That is, if we have:
		//
		//	func f(x int) {
		//		func() {
		//			global = &x
		//		}
		//	}
		//
		// then we're analyzing the inner closure but we need to move x to the
		// heap in f, not in the inner closure. Flip over to f before calling moveToHeap.
		oldfn := pstate.Curfn
		pstate.Curfn = n.Name.Curfn
		if pstate.Curfn.Func.Closure != nil && pstate.Curfn.Op == OCLOSURE {
			pstate.Curfn = pstate.Curfn.Func.Closure
		}
		ln := pstate.lineno
		pstate.lineno = pstate.Curfn.Pos
		pstate.moveToHeap(n)
		pstate.Curfn = oldfn
		pstate.lineno = ln

	// ODOTPTR has already been introduced,
	// so these are the non-pointer ODOT and OINDEX.
	// In &x[0], if x is a slice, then x does not
	// escape--the pointer inside x does, but that
	// is always a heap pointer anyway.
	case ODOT, OINDEX, OPAREN, OCONVNOP:
		if !n.Left.Type.IsSlice() {
			pstate.addrescapes(n.Left)
		}
	}
}

// moveToHeap records the parameter or local variable n as moved to the heap.
func (pstate *PackageState) moveToHeap(n *Node) {
	if pstate.Debug['r'] != 0 {
		Dump("MOVE", n)
	}
	if pstate.compiling_runtime {
		pstate.yyerror("%v escapes to heap, not allowed in runtime.", n)
	}
	if n.Class() == PAUTOHEAP {
		Dump("n", n)
		pstate.Fatalf("double move to heap")
	}

	// Allocate a local stack variable to hold the pointer to the heap copy.
	// temp will add it to the function declaration list automatically.
	heapaddr := pstate.temp(pstate.types.NewPtr(n.Type))
	heapaddr.Sym = pstate.lookup("&" + n.Sym.Name)
	heapaddr.Orig.Sym = heapaddr.Sym
	heapaddr.Pos = n.Pos

	// Unset AutoTemp to persist the &foo variable name through SSA to
	// liveness analysis.
	// TODO(mdempsky/drchase): Cleaner solution?
	heapaddr.Name.SetAutoTemp(false)

	// Parameters have a local stack copy used at function start/end
	// in addition to the copy in the heap that may live longer than
	// the function.
	if n.Class() == PPARAM || n.Class() == PPARAMOUT {
		if n.Xoffset == BADWIDTH {
			pstate.Fatalf("addrescapes before param assignment")
		}

		// We rewrite n below to be a heap variable (indirection of heapaddr).
		// Preserve a copy so we can still write code referring to the original,
		// and substitute that copy into the function declaration list
		// so that analyses of the local (on-stack) variables use it.
		stackcopy := pstate.newname(n.Sym)
		stackcopy.SetAddable(false)
		stackcopy.Type = n.Type
		stackcopy.Xoffset = n.Xoffset
		stackcopy.SetClass(n.Class())
		stackcopy.Name.Param.Heapaddr = heapaddr
		if n.Class() == PPARAMOUT {
			// Make sure the pointer to the heap copy is kept live throughout the function.
			// The function could panic at any point, and then a defer could recover.
			// Thus, we need the pointer to the heap copy always available so the
			// post-deferreturn code can copy the return value back to the stack.
			// See issue 16095.
			heapaddr.SetIsOutputParamHeapAddr(true)
		}
		n.Name.Param.Stackcopy = stackcopy

		// Substitute the stackcopy into the function variable list so that
		// liveness and other analyses use the underlying stack slot
		// and not the now-pseudo-variable n.
		found := false
		for i, d := range pstate.Curfn.Func.Dcl {
			if d == n {
				pstate.Curfn.Func.Dcl[i] = stackcopy
				found = true
				break
			}
			// Parameters are before locals, so can stop early.
			// This limits the search even in functions with many local variables.
			if d.Class() == PAUTO {
				break
			}
		}
		if !found {
			pstate.Fatalf("cannot find %v in local variable list", n)
		}
		pstate.Curfn.Func.Dcl = append(pstate.Curfn.Func.Dcl, n)
	}

	// Modify n in place so that uses of n now mean indirection of the heapaddr.
	n.SetClass(PAUTOHEAP)
	n.Xoffset = 0
	n.Name.Param.Heapaddr = heapaddr
	n.Esc = EscHeap
	if pstate.Debug['m'] != 0 {
		fmt.Printf("%v: moved to heap: %v\n", n.Line(pstate), n)
	}
}

// This special tag is applied to uintptr variables
// that we believe may hold unsafe.Pointers for
// calls into assembly functions.
const unsafeUintptrTag = "unsafe-uintptr"

// This special tag is applied to uintptr parameters of functions
// marked go:uintptrescapes.
const uintptrEscapesTag = "uintptr-escapes"

func (e *EscState) esctag(pstate *PackageState, fn *Node) {
	fn.Esc = EscFuncTagged

	name := func(s *types.Sym, narg int) string {
		if s != nil {
			return s.Name
		}
		return fmt.Sprintf("arg#%d", narg)
	}

	// External functions are assumed unsafe,
	// unless //go:noescape is given before the declaration.
	if fn.Nbody.Len() == 0 {
		if fn.Noescape() {
			for _, f := range fn.Type.Params(pstate.types).Fields(pstate.types).Slice() {
				if pstate.types.Haspointers(f.Type) {
					f.Note = pstate.mktag(EscNone)
				}
			}
		}

		// Assume that uintptr arguments must be held live across the call.
		// This is most important for syscall.Syscall.
		// See golang.org/issue/13372.
		// This really doesn't have much to do with escape analysis per se,
		// but we are reusing the ability to annotate an individual function
		// argument and pass those annotations along to importing code.
		narg := 0
		for _, f := range fn.Type.Params(pstate.types).Fields(pstate.types).Slice() {
			narg++
			if f.Type.Etype == TUINTPTR {
				if pstate.Debug['m'] != 0 {
					pstate.Warnl(fn.Pos, "%v assuming %v is unsafe uintptr", funcSym(fn), name(f.Sym, narg))
				}
				f.Note = unsafeUintptrTag
			}
		}

		return
	}

	if fn.Func.Pragma&UintptrEscapes != 0 {
		narg := 0
		for _, f := range fn.Type.Params(pstate.types).Fields(pstate.types).Slice() {
			narg++
			if f.Type.Etype == TUINTPTR {
				if pstate.Debug['m'] != 0 {
					pstate.Warnl(fn.Pos, "%v marking %v as escaping uintptr", funcSym(fn), name(f.Sym, narg))
				}
				f.Note = uintptrEscapesTag
			}

			if f.Isddd() && f.Type.Elem(pstate.types).Etype == TUINTPTR {
				// final argument is ...uintptr.
				if pstate.Debug['m'] != 0 {
					pstate.Warnl(fn.Pos, "%v marking %v as escaping ...uintptr", funcSym(fn), name(f.Sym, narg))
				}
				f.Note = uintptrEscapesTag
			}
		}
	}

	for _, fs := range pstate.types.RecvsParams {
		for _, f := range fs(fn.Type).Fields(pstate.types).Slice() {
			if !pstate.types.Haspointers(f.Type) { // don't bother tagging for scalars
				continue
			}
			if f.Note == uintptrEscapesTag {
				// Note is already set in the loop above.
				continue
			}

			// Unnamed parameters are unused and therefore do not escape.
			if f.Sym == nil || f.Sym.IsBlank() {
				f.Note = pstate.mktag(EscNone)
				continue
			}

			switch esc := asNode(f.Nname).Esc; esc & EscMask {
			case EscNone, // not touched by escflood
				EscReturn:
				f.Note = pstate.mktag(int(esc))

			case EscHeap: // touched by escflood, moved to heap
			}
		}
	}
}

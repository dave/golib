package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"strconv"
	"strings"
)

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
func (psess *PackageSession) visitBottomUp(list []*Node, analyze func(list []*Node, recursive bool)) {
	var v bottomUpVisitor
	v.analyze = analyze
	v.nodeID = make(map[*Node]uint32)
	for _, n := range list {
		if n.Op == ODCLFUNC && !n.Func.IsHiddenClosure() {
			v.visit(psess, n)
		}
	}
}

func (v *bottomUpVisitor) visit(psess *PackageSession, n *Node) uint32 {
	if id := v.nodeID[n]; id > 0 {

		return id
	}

	v.visitgen++
	id := v.visitgen
	v.nodeID[n] = id
	v.visitgen++
	min := v.visitgen

	v.stack = append(v.stack, n)
	min = v.visitcodelist(psess, n.Nbody, min)
	if (min == id || min == id+1) && !n.Func.IsHiddenClosure() {

		recursive := min == id

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

		v.stack = v.stack[:i]
		v.analyze(block, recursive)
	}

	return min
}

func (v *bottomUpVisitor) visitcodelist(psess *PackageSession, l Nodes, min uint32) uint32 {
	for _, n := range l.Slice() {
		min = v.visitcode(psess, n, min)
	}
	return min
}

func (v *bottomUpVisitor) visitcode(psess *PackageSession, n *Node, min uint32) uint32 {
	if n == nil {
		return min
	}

	min = v.visitcodelist(psess, n.Ninit, min)
	min = v.visitcode(psess, n.Left, min)
	min = v.visitcode(psess, n.Right, min)
	min = v.visitcodelist(psess, n.List, min)
	min = v.visitcodelist(psess, n.Nbody, min)
	min = v.visitcodelist(psess, n.Rlist, min)

	switch n.Op {
	case OCALLFUNC, OCALLMETH:
		fn := asNode(n.Left.Type.Nname(psess.types))
		if fn != nil && fn.Op == ONAME && fn.Class() == PFUNC && fn.Name.Defn != nil {
			m := v.visit(psess, fn.Name.Defn)
			if m < min {
				min = m
			}
		}

	case OCLOSURE:
		m := v.visit(psess, n.Func.Closure)
		if m < min {
			min = m
		}
	}

	return min
}

func (psess *PackageSession) escapes(all []*Node) {
	psess.
		visitBottomUp(all, psess.escAnalyze)
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

func (e *EscState) nodeEscState(psess *PackageSession, n *Node) *NodeEscState {
	if nE, ok := n.Opt().(*NodeEscState); ok {
		return nE
	}
	if n.Opt() != nil {
		psess.
			Fatalf("nodeEscState: opt in use (%T)", n.Opt())
	}
	nE := &NodeEscState{
		Curfn: psess.Curfn,
	}
	n.SetOpt(psess, nE)
	e.opts = append(e.opts, n)
	return nE
}

func (e *EscState) track(psess *PackageSession, n *Node) {
	if psess.Curfn == nil {
		psess.
			Fatalf("EscState.track: Curfn nil")
	}
	n.Esc = EscNone
	nE := e.nodeEscState(psess, n)
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
)

// escMax returns the maximum of an existing escape value
// (and its additional parameter flow flags) and a new escape type.
func (psess *PackageSession) escMax(e, etype uint16) uint16 {
	if e&EscMask >= EscHeap {

		if e&^EscMask != 0 {
			psess.
				Fatalf("Escape information had unexpected return encoding bits (w/ EscHeap, EscNever), e&EscMask=%v", e&EscMask)
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

func (psess *PackageSession) newEscState(recursive bool) *EscState {
	e := new(EscState)
	e.theSink.Op = ONAME
	e.theSink.Orig = &e.theSink
	e.theSink.SetClass(PEXTERN)
	e.theSink.Sym = psess.lookup(".sink")
	e.nodeEscState(psess, &e.theSink).Loopdepth = -1
	e.recursive = recursive
	return e
}

func (e *EscState) stepWalk(psess *PackageSession, dst, src *Node, why string, parent *EscStep) *EscStep {

	if psess.Debug['m'] == 0 {
		return nil
	}
	return &EscStep{src: src, dst: dst, why: why, parent: parent}
}

func (e *EscState) stepAssign(psess *PackageSession, step *EscStep, dst, src *Node, why string) *EscStep {
	if psess.Debug['m'] == 0 {
		return nil
	}
	if step != nil {
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

func (e *EscState) stepAssignWhere(psess *PackageSession, dst, src *Node, why string, where *Node) *EscStep {
	if psess.Debug['m'] == 0 {
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
func (e *EscState) curfnSym(psess *PackageSession, n *Node) *types.Sym {
	nE := e.nodeEscState(psess, n)
	return funcSym(nE.Curfn)
}

func (psess *PackageSession) escAnalyze(all []*Node, recursive bool) {
	e := psess.newEscState(recursive)

	for _, n := range all {
		if n.Op == ODCLFUNC {
			n.Esc = EscFuncPlanned
			if psess.Debug['m'] > 3 {
				Dump("escAnalyze", n)
			}

		}
	}

	for _, n := range all {
		if n.Op == ODCLFUNC {
			e.escfunc(psess, n)
		}
	}

	escapes := make([]uint16, len(e.dsts))
	for i, n := range e.dsts {
		escapes[i] = n.Esc
	}
	for _, n := range e.dsts {
		e.escflood(psess, n)
	}
	for {
		done := true
		for i, n := range e.dsts {
			if n.Esc != escapes[i] {
				done = false
				if psess.Debug['m'] > 2 {
					psess.
						Warnl(n.Pos, "Reflooding %v %S", e.curfnSym(psess, n), n)
				}
				escapes[i] = n.Esc
				e.escflood(psess, n)
			}
		}
		if done {
			break
		}
	}

	for _, n := range all {
		if n.Op == ODCLFUNC {
			e.esctag(psess, n)
		}
	}

	if psess.Debug['m'] != 0 {
		for _, n := range e.noesc {
			if n.Esc == EscNone {
				psess.
					Warnl(n.Pos, "%v %S does not escape", e.curfnSym(psess, n), n)
			}
		}
	}

	for _, x := range e.opts {
		x.SetOpt(psess, nil)
	}
}

func (e *EscState) escfunc(psess *PackageSession, fn *Node) {

	if fn.Esc != EscFuncPlanned {
		psess.
			Fatalf("repeat escfunc %v", fn.Func.Nname)
	}
	fn.Esc = EscFuncStarted

	saveld := e.loopdepth
	e.loopdepth = 1
	savefn := psess.Curfn
	psess.
		Curfn = fn

	for _, ln := range psess.Curfn.Func.Dcl {
		if ln.Op != ONAME {
			continue
		}
		lnE := e.nodeEscState(psess, ln)
		switch ln.Class() {

		case PPARAMOUT:
			lnE.Loopdepth = 0

		case PPARAM:
			lnE.Loopdepth = 1
			if ln.Type != nil && !psess.types.Haspointers(ln.Type) {
				break
			}
			if psess.Curfn.Nbody.Len() == 0 && !psess.Curfn.Noescape() {
				ln.Esc = EscHeap
			} else {
				ln.Esc = EscNone
			}
			e.noesc = append(e.noesc, ln)
		}
	}

	if e.recursive {
		for _, ln := range psess.Curfn.Func.Dcl {
			if ln.Op == ONAME && ln.Class() == PPARAMOUT {
				e.escflows(psess, &e.theSink, ln, e.stepAssign(psess, nil, ln, ln, "returned from recursive function"))
			}
		}
	}

	e.escloopdepthlist(psess, psess.Curfn.Nbody)
	e.esclist(psess, psess.Curfn.Nbody, psess.Curfn)
	psess.
		Curfn = savefn
	e.loopdepth = saveld
}

// Mark labels that have no backjumps to them as not increasing e.loopdepth.
// Walk hasn't generated (goto|label).Left.Sym.Label yet, so we'll cheat
// and set it to one of the following two. Then in esc we'll clear it again.

func (e *EscState) escloopdepthlist(psess *PackageSession, l Nodes) {
	for _, n := range l.Slice() {
		e.escloopdepth(psess, n)
	}
}

func (e *EscState) escloopdepth(psess *PackageSession, n *Node) {
	if n == nil {
		return
	}

	e.escloopdepthlist(psess, n.Ninit)

	switch n.Op {
	case OLABEL:
		if n.Left == nil || n.Left.Sym == nil {
			psess.
				Fatalf("esc:label without label: %+v", n)
		}

		n.Left.Sym.Label = asTypesNode(&psess.nonlooping)

	case OGOTO:
		if n.Left == nil || n.Left.Sym == nil {
			psess.
				Fatalf("esc:goto without label: %+v", n)
		}

		if asNode(n.Left.Sym.Label) == &psess.nonlooping {
			n.Left.Sym.Label = asTypesNode(&psess.looping)
		}
	}

	e.escloopdepth(psess, n.Left)
	e.escloopdepth(psess, n.Right)
	e.escloopdepthlist(psess, n.List)
	e.escloopdepthlist(psess, n.Nbody)
	e.escloopdepthlist(psess, n.Rlist)
}

func (e *EscState) esclist(psess *PackageSession, l Nodes, parent *Node) {
	for _, n := range l.Slice() {
		e.esc(psess, n, parent)
	}
}

func (e *EscState) esc(psess *PackageSession, n *Node, parent *Node) {
	if n == nil {
		return
	}

	lno := psess.setlineno(n)

	e.esclist(psess, n.Ninit, n)

	if n.Op == OFOR || n.Op == OFORUNTIL || n.Op == ORANGE {
		e.loopdepth++
	}

	if n.Op == OSWITCH && n.Left != nil && n.Left.Op == OTYPESW {
		for _, cas := range n.List.Slice() {

			if cas.Rlist.Len() != 0 {
				e.nodeEscState(psess, cas.Rlist.First()).Loopdepth = e.loopdepth
			}
		}
	}

	if n.Esc != EscHeap && n.Type != nil &&
		(n.Type.Width > maxStackVarSize ||
			(n.Op == ONEW || n.Op == OPTRLIT) && n.Type.Elem(psess.types).Width >= 1<<16 ||
			n.Op == OMAKESLICE && !psess.isSmallMakeSlice(n)) {

		// isSmallMakeSlice returns false for non-constant len/cap.
		// If that's the case, print a more accurate escape reason.
		var msgVerb, escapeMsg string
		if n.Op == OMAKESLICE && (!psess.Isconst(n.Left, CTINT) || !psess.Isconst(n.Right, CTINT)) {
			msgVerb, escapeMsg = "has ", "non-constant size"
		} else {
			msgVerb, escapeMsg = "is ", "too large for stack"
		}

		if psess.Debug['m'] > 2 {
			psess.
				Warnl(n.Pos, "%v "+msgVerb+escapeMsg, n)
		}
		n.Esc = EscHeap
		psess.
			addrescapes(n)
		e.escassignSinkWhy(psess, n, n, escapeMsg)
	}

	e.esc(psess, n.Left, n)

	if n.Op == ORANGE {

		e.loopdepth--
	}

	e.esc(psess, n.Right, n)

	if n.Op == ORANGE {
		e.loopdepth++
	}

	e.esclist(psess, n.Nbody, n)
	e.esclist(psess, n.List, n)
	e.esclist(psess, n.Rlist, n)

	if n.Op == OFOR || n.Op == OFORUNTIL || n.Op == ORANGE {
		e.loopdepth--
	}

	if psess.Debug['m'] > 2 {
		fmt.Printf("%v:[%d] %v esc: %v\n", psess.linestr(psess.lineno), e.loopdepth, funcSym(psess.Curfn), n)
	}

opSwitch:
	switch n.Op {

	case ODCL:
		if n.Left != nil {
			e.nodeEscState(psess, n.Left).Loopdepth = e.loopdepth
		}

	case OLABEL:
		if asNode(n.Left.Sym.Label) == &psess.nonlooping {
			if psess.Debug['m'] > 2 {
				fmt.Printf("%v:%v non-looping label\n", psess.linestr(psess.lineno), n)
			}
		} else if asNode(n.Left.Sym.Label) == &psess.looping {
			if psess.Debug['m'] > 2 {
				fmt.Printf("%v: %v looping label\n", psess.linestr(psess.lineno), n)
			}
			e.loopdepth++
		}

		n.Left.Sym.Label = nil

	case ORANGE:
		if n.List.Len() >= 2 {

			if n.Type.IsArray() &&
				!(n.Right.Type.IsPtr() && psess.eqtype(n.Right.Type.Elem(psess.types), n.Type)) {
				e.escassignWhyWhere(psess, n.List.Second(), n.Right, "range", n)
			} else {
				e.escassignDereference(psess, n.List.Second(), n.Right, e.stepAssignWhere(psess, n.List.Second(), n.Right, "range-deref", n))
			}
		}

	case OSWITCH:
		if n.Left != nil && n.Left.Op == OTYPESW {
			for _, cas := range n.List.Slice() {

				if cas.Rlist.Len() != 0 {
					e.escassignWhyWhere(psess, cas.Rlist.First(), n.Left.Right, "switch case", n)
				}
			}
		}

	case OAS, OASOP:
		if (n.Left.Op == OIND || n.Left.Op == ODOTPTR) && n.Left.Left.Op == ONAME &&
			(n.Right.Op == OSLICE || n.Right.Op == OSLICE3 || n.Right.Op == OSLICESTR) &&
			(n.Right.Left.Op == OIND || n.Right.Left.Op == ODOTPTR) && n.Right.Left.Left.Op == ONAME &&
			n.Left.Left == n.Right.Left.Left {

			if psess.Debug['m'] != 0 {
				psess.
					Warnl(n.Pos, "%v ignoring self-assignment to %S", e.curfnSym(psess, n), n.Left)
			}

			break
		}

		e.escassign(psess, n.Left, n.Right, e.stepAssignWhere(psess, nil, nil, "", n))

	case OAS2:
		if n.List.Len() == n.Rlist.Len() {
			rs := n.Rlist.Slice()
			for i, n := range n.List.Slice() {
				e.escassignWhyWhere(psess, n, rs[i], "assign-pair", n)
			}
		}

	case OAS2RECV:
		e.escassignWhyWhere(psess, n.List.First(), n.Rlist.First(), "assign-pair-receive", n)
	case OAS2MAPR:
		e.escassignWhyWhere(psess, n.List.First(), n.Rlist.First(), "assign-pair-mapr", n)
	case OAS2DOTTYPE:
		e.escassignWhyWhere(psess, n.List.First(), n.Rlist.First(), "assign-pair-dot-type", n)

	case OSEND:
		e.escassignSinkWhy(psess, n, n.Right, "send")

	case ODEFER:
		if e.loopdepth == 1 {
			break
		}

		e.escassignSinkWhy(psess, n, n.Left.Left, "defer func")
		e.escassignSinkWhy(psess, n, n.Left.Right, "defer func ...")
		for _, arg := range n.Left.List.Slice() {
			e.escassignSinkWhy(psess, n, arg, "defer func arg")
		}

	case OPROC:

		e.escassignSinkWhy(psess, n, n.Left.Left, "go func")
		e.escassignSinkWhy(psess, n, n.Left.Right, "go func ...")
		for _, arg := range n.Left.List.Slice() {
			e.escassignSinkWhy(psess, n, arg, "go func arg")
		}

	case OCALLMETH, OCALLFUNC, OCALLINTER:
		e.esccall(psess, n, parent)

	case OAS2FUNC:
		rs := e.nodeEscState(psess, n.Rlist.First()).Retval.Slice()
		for i, n := range n.List.Slice() {
			if i >= len(rs) {
				break
			}
			e.escassignWhyWhere(psess, n, rs[i], "assign-pair-func-call", n)
		}
		if n.List.Len() != len(rs) {
			psess.
				Fatalf("esc oas2func")
		}

	case ORETURN:
		retList := n.List
		if retList.Len() == 1 && psess.Curfn.Type.NumResults(psess.types) > 1 {

			retList = e.nodeEscState(psess, n.List.First()).Retval
		}

		i := 0
		for _, lrn := range psess.Curfn.Func.Dcl {
			if i >= retList.Len() {
				break
			}
			if lrn.Op != ONAME || lrn.Class() != PPARAMOUT {
				continue
			}
			e.escassignWhyWhere(psess, lrn, retList.Index(i), "return", n)
			i++
		}

		if i < retList.Len() {
			psess.
				Fatalf("esc return list")
		}

	case OPANIC:
		e.escassignSinkWhy(psess, n, n.Left, "panic")

	case OAPPEND:
		if !n.Isddd() {
			for _, nn := range n.List.Slice()[1:] {
				e.escassignSinkWhy(psess, n, nn, "appended to slice")
			}
		} else {

			slice2 := n.List.Second()
			e.escassignDereference(psess, &e.theSink, slice2, e.stepAssignWhere(psess, n, slice2, "appended slice...", n))
			if psess.Debug['m'] > 3 {
				psess.
					Warnl(n.Pos, "%v special treatment of append(slice1, slice2...) %S", e.curfnSym(psess, n), n)
			}
		}
		e.escassignDereference(psess, &e.theSink, n.List.First(), e.stepAssignWhere(psess, n, n.List.First(), "appendee slice", n))

	case OCOPY:
		e.escassignDereference(psess, &e.theSink, n.Right, e.stepAssignWhere(psess, n, n.Right, "copied slice", n))

	case OCONV, OCONVNOP:
		e.escassignWhyWhere(psess, n, n.Left, "converted", n)

	case OCONVIFACE:
		e.track(psess, n)
		e.escassignWhyWhere(psess, n, n.Left, "interface-converted", n)

	case OARRAYLIT:

		for _, elt := range n.List.Slice() {
			if elt.Op == OKEY {
				elt = elt.Right
			}
			e.escassign(psess, n, elt, e.stepAssignWhere(psess, n, elt, "array literal element", n))
		}

	case OSLICELIT:

		e.track(psess, n)

		for _, elt := range n.List.Slice() {
			if elt.Op == OKEY {
				elt = elt.Right
			}
			e.escassign(psess, n, elt, e.stepAssignWhere(psess, n, elt, "slice literal element", n))
		}

	case OSTRUCTLIT:
		for _, elt := range n.List.Slice() {
			e.escassignWhyWhere(psess, n, elt.Left, "struct literal element", n)
		}

	case OPTRLIT:
		e.track(psess, n)

		e.escassignWhyWhere(psess, n, n.Left, "pointer literal [assign]", n)

	case OCALLPART:
		e.track(psess, n)

		e.escassignSinkWhy(psess, n, n.Left, "call part")

	case OMAPLIT:
		e.track(psess, n)

		for _, elt := range n.List.Slice() {
			e.escassignSinkWhy(psess, n, elt.Left, "map literal key")
			e.escassignSinkWhy(psess, n, elt.Right, "map literal value")
		}

	case OCLOSURE:

		for _, v := range n.Func.Closure.Func.Cvars.Slice() {
			if v.Op == OXXX {
				continue
			}
			a := v.Name.Defn
			if !v.Name.Byval() {
				a = psess.nod(OADDR, a, nil)
				a.Pos = v.Pos
				e.nodeEscState(psess, a).Loopdepth = e.loopdepth
				a = psess.typecheck(a, Erv)
			}

			e.escassignWhyWhere(psess, n, a, "captured by a closure", n)
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
		e.track(psess, n)

	case OADDSTR:
		e.track(psess, n)

	case OADDR:

		e.track(psess, n)

		if n.Left.Op == ONAME {
			switch n.Left.Class() {

			case PPARAM, PPARAMOUT:
				nE := e.nodeEscState(psess, n)
				nE.Loopdepth = 1
				break opSwitch
			}
		}
		nE := e.nodeEscState(psess, n)
		leftE := e.nodeEscState(psess, n.Left)
		if leftE.Loopdepth != 0 {
			nE.Loopdepth = leftE.Loopdepth
		}

	case ODOT,
		ODOTPTR,
		OINDEX:

		if n.Left.Op != OLITERAL {
			e.nodeEscState(psess, n).Loopdepth = e.nodeEscState(psess, n.Left).Loopdepth
		}
	}
	psess.
		lineno = lno
}

// escassignWhyWhere bundles a common case of
// escassign(e, dst, src, e.stepAssignWhere(dst, src, reason, where))
func (e *EscState) escassignWhyWhere(psess *PackageSession, dst, src *Node, reason string, where *Node) {
	var step *EscStep
	if psess.Debug['m'] != 0 {
		step = e.stepAssignWhere(psess, dst, src, reason, where)
	}
	e.escassign(psess, dst, src, step)
}

// escassignSinkWhy bundles a common case of
// escassign(e, &e.theSink, src, e.stepAssign(nil, dst, src, reason))
func (e *EscState) escassignSinkWhy(psess *PackageSession, dst, src *Node, reason string) {
	var step *EscStep
	if psess.Debug['m'] != 0 {
		step = e.stepAssign(psess, nil, dst, src, reason)
	}
	e.escassign(psess, &e.theSink, src, step)
}

// escassignSinkWhyWhere is escassignSinkWhy but includes a call site
// for accurate location reporting.
func (e *EscState) escassignSinkWhyWhere(psess *PackageSession, dst, src *Node, reason string, call *Node) {
	var step *EscStep
	if psess.Debug['m'] != 0 {
		step = e.stepAssignWhere(psess, dst, src, reason, call)
	}
	e.escassign(psess, &e.theSink, src, step)
}

// Assert that expr somehow gets assigned to dst, if non nil.  for
// dst==nil, any name node expr still must be marked as being
// evaluated in curfn.	For expr==nil, dst must still be examined for
// evaluations inside it (e.g *f(x) = y)
func (e *EscState) escassign(psess *PackageSession, dst, src *Node, step *EscStep) {
	if dst.isBlank() || dst == nil || src == nil || src.Op == ONONAME || src.Op == OXXX {
		return
	}

	if psess.Debug['m'] > 2 {
		fmt.Printf("%v:[%d] %v escassign: %S(%0j)[%v] = %S(%0j)[%v]\n", psess.
			linestr(psess.lineno), e.loopdepth, funcSym(psess.Curfn),
			dst, dst, dst.Op,
			src, src, src.Op)
	}
	psess.
		setlineno(dst)

	originalDst := dst
	dstwhy := "assigned"

	switch dst.Op {
	default:
		Dump("dst", dst)
		psess.
			Fatalf("escassign: unexpected dst")

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

	case ODOT:
		e.escassign(psess, dst.Left, src, e.stepAssign(psess, step, originalDst, src, "dot-equals"))
		return

	case OINDEX:
		if dst.Left.Type.IsArray() {
			e.escassign(psess, dst.Left, src, e.stepAssign(psess, step, originalDst, src, "array-element-equals"))
			return
		}

		dstwhy = "slice-element-equals"
		dst = &e.theSink

	case OIND:
		dstwhy = "star-equals"
		dst = &e.theSink

	case ODOTPTR:
		dstwhy = "star-dot-equals"
		dst = &e.theSink

	case OINDEXMAP:
		e.escassign(psess, &e.theSink, dst.Right, e.stepAssign(psess, nil, originalDst, src, "key of map put"))
		dstwhy = "value of map put"
		dst = &e.theSink
	}

	lno := psess.setlineno(src)
	e.pdepth++

	switch src.Op {
	case OADDR,
		OIND,
		ODOTPTR,
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
		e.escflows(psess, dst, src, e.stepAssign(psess, step, originalDst, src, dstwhy))

	case OCLOSURE:

		a := psess.nod(OADDR, src, nil)
		a.Pos = src.Pos
		e.nodeEscState(psess, a).Loopdepth = e.nodeEscState(psess, src).Loopdepth
		a.Type = psess.types.NewPtr(src.Type)
		e.escflows(psess, dst, a, e.stepAssign(psess, nil, originalDst, src, dstwhy))

	case OCALLMETH, OCALLFUNC, OCALLINTER:
		for _, n := range e.nodeEscState(psess, src).Retval.Slice() {
			e.escflows(psess, dst, n, e.stepAssign(psess, nil, originalDst, n, dstwhy))
		}

	case ODOT:
		if src.Type != nil && !psess.types.Haspointers(src.Type) {
			break
		}
		fallthrough

	case OCONV,
		OCONVNOP,
		ODOTMETH,

		OSLICE,
		OSLICE3,
		OSLICEARR,
		OSLICE3ARR,
		OSLICESTR:

		e.escassign(psess, dst, src.Left, e.stepAssign(psess, step, originalDst, src, dstwhy))

	case ODOTTYPE,
		ODOTTYPE2:
		if src.Type != nil && !psess.types.Haspointers(src.Type) {
			break
		}
		e.escassign(psess, dst, src.Left, e.stepAssign(psess, step, originalDst, src, dstwhy))

	case OAPPEND:

		e.escassign(psess, dst, src.List.First(), e.stepAssign(psess, step, dst, src.List.First(), dstwhy))

	case OINDEX:

		if src.Left.Type.IsArray() {
			e.escassign(psess, dst, src.Left, e.stepAssign(psess, step, originalDst, src, dstwhy))
		} else {
			e.escflows(psess, dst, src, e.stepAssign(psess, step, originalDst, src, dstwhy))
		}

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
		e.escassign(psess, dst, src.Left, e.stepAssign(psess, step, originalDst, src, dstwhy))

		e.escassign(psess, dst, src.Right, e.stepAssign(psess, step, originalDst, src, dstwhy))
	}

	e.pdepth--
	psess.
		lineno = lno
}

// Common case for escapes is 16 bits 000000000xxxEEEE
// where commonest cases for xxx encoding in-to-out pointer
//  flow are 000, 001, 010, 011  and EEEE is computed Esc bits.
// Note width of xxx depends on value of constant
// bitsPerOutputInTag -- expect 2 or 3, so in practice the
// tag cache array is 64 or 128 long. Some entries will
// never be populated.

// mktag returns the string representation for an escape analysis tag.
func (psess *PackageSession) mktag(mask int) string {
	switch mask & EscMask {
	case EscNone, EscReturn:
	default:
		psess.
			Fatalf("escape mktag")
	}

	if mask < len(psess.tags) && psess.tags[mask] != "" {
		return psess.tags[mask]
	}

	s := fmt.Sprintf("esc:0x%x", mask)
	if mask < len(psess.tags) {
		psess.
			tags[mask] = s
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
func (e *EscState) escassignfromtag(psess *PackageSession, note string, dsts Nodes, src, call *Node) uint16 {
	em := parsetag(note)
	if src.Op == OLITERAL {
		return em
	}

	if psess.Debug['m'] > 3 {
		fmt.Printf("%v::assignfromtag:: src=%S, em=%s\n", psess.
			linestr(psess.lineno), src, describeEscape(em))
	}

	if em == EscUnknown {
		e.escassignSinkWhyWhere(psess, src, src, "passed to call[argument escapes]", call)
		return em
	}

	if em == EscNone {
		return em
	}

	if em&EscContentEscapes != 0 {
		e.escassign(psess, &e.theSink, e.addDereference(psess, src), e.stepAssignWhere(psess, src, src, "passed to call[argument content escapes]", call))
	}

	em0 := em
	dstsi := 0
	for em >>= EscReturnBits; em != 0 && dstsi < dsts.Len(); em = em >> bitsPerOutputInTag {

		embits := em & bitsMaskForTag
		if embits > 0 {
			n := src
			for i := uint16(0); i < embits-1; i++ {
				n = e.addDereference(psess, n)
			}
			e.escassign(psess, dsts.Index(dstsi), n, e.stepAssignWhere(psess, dsts.Index(dstsi), src, "passed-to-and-returned-from-call", call))
		}
		dstsi++
	}

	if em != 0 && dstsi >= dsts.Len() {
		psess.
			Fatalf("corrupt esc tag %q or messed up escretval list\n", note)
	}
	return em0
}

func (e *EscState) escassignDereference(psess *PackageSession, dst *Node, src *Node, step *EscStep) {
	if src.Op == OLITERAL {
		return
	}
	e.escassign(psess, dst, e.addDereference(psess, src), step)
}

// addDereference constructs a suitable OIND note applied to src.
// Because this is for purposes of escape accounting, not execution,
// some semantically dubious node combinations are (currently) possible.
func (e *EscState) addDereference(psess *PackageSession, n *Node) *Node {
	ind := psess.nod(OIND, n, nil)
	e.nodeEscState(psess, ind).Loopdepth = e.nodeEscState(psess, n).Loopdepth
	ind.Pos = n.Pos
	t := n.Type
	if t.IsKind(psess.types.Tptr) || t.IsSlice() {

		t = t.Elem(psess.types)
	} else if t.IsString() {
		t = psess.types.Types[TUINT8]
	}
	ind.Type = t
	return ind
}

// escNoteOutputParamFlow encodes maxEncodedLevel/.../1/0-level flow to the vargen'th parameter.
// Levels greater than maxEncodedLevel are replaced with maxEncodedLevel.
// If the encoding cannot describe the modified input level and output number, then EscHeap is returned.
func (psess *PackageSession) escNoteOutputParamFlow(e uint16, vargen int32, level Level) uint16 {

	if level.int() <= 0 && level.guaranteedDereference() > 0 {
		return psess.escMax(e|EscContentEscapes, EscNone)
	}
	if level.int() < 0 {
		return EscHeap
	}
	if level.int() > maxEncodedLevel {

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

		return EscHeap
	}

	return (e &^ (bitsMaskForTag << shift)) | encodedFlow
}

func (e *EscState) initEscRetval(psess *PackageSession, call *Node, fntype *types.Type) {
	cE := e.nodeEscState(psess, call)
	cE.Retval.Set(nil)
	for i, f := range fntype.Results(psess.types).Fields(psess.types).Slice() {
		buf := fmt.Sprintf(".out%d", i)
		ret := psess.newname(psess.lookup(buf))
		ret.SetAddable(false)
		ret.Type = f.Type
		ret.SetClass(PAUTO)
		ret.Name.Curfn = psess.Curfn
		e.nodeEscState(psess, ret).Loopdepth = e.loopdepth
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
func (e *EscState) esccall(psess *PackageSession, call *Node, parent *Node) {
	var fntype *types.Type
	var indirect bool
	var fn *Node
	switch call.Op {
	default:
		psess.
			Fatalf("esccall")

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
		if arg.Type.IsFuncArgStruct() {
			argList = e.nodeEscState(psess, arg).Retval
		}
	}

	args := argList.Slice()

	if indirect {

		for _, arg := range args {
			e.escassignSinkWhy(psess, call, arg, "parameter to indirect call")
			if psess.Debug['m'] > 3 {
				fmt.Printf("%v::esccall:: indirect call <- %S, untracked\n", psess.linestr(psess.lineno), arg)
			}
		}

		e.initEscRetval(psess, call, fntype)

		if call.Op != OCALLFUNC {
			rf := fntype.Recv(psess.types)
			r := call.Left.Left
			if psess.types.Haspointers(rf.Type) {
				e.escassignSinkWhy(psess, call, r, "receiver in indirect call")
			}
		} else {
			rets := e.nodeEscState(psess, call).Retval.Slice()
			for _, ret := range rets {
				e.escassignDereference(psess, ret, fn, e.stepAssignWhere(psess, ret, fn, "captured by called closure", call))
			}
		}
		return
	}

	cE := e.nodeEscState(psess, call)
	if fn != nil && fn.Op == ONAME && fn.Class() == PFUNC &&
		fn.Name.Defn != nil && fn.Name.Defn.Nbody.Len() != 0 && fn.Name.Param.Ntype != nil && fn.Name.Defn.Esc < EscFuncTagged {
		if psess.Debug['m'] > 3 {
			fmt.Printf("%v::esccall:: %S in recursive group\n", psess.linestr(psess.lineno), call)
		}

		if fn.Name.Defn.Esc == EscFuncUnknown || cE.Retval.Len() != 0 {
			psess.
				Fatalf("graph inconsistency")
		}

		sawRcvr := false
		for _, n := range fn.Name.Defn.Func.Dcl {
			switch n.Class() {
			case PPARAM:
				if call.Op != OCALLFUNC && !sawRcvr {
					e.escassignWhyWhere(psess, n, call.Left.Left, "call receiver", call)
					sawRcvr = true
					continue
				}
				if len(args) == 0 {
					continue
				}
				arg := args[0]
				if n.Isddd() && !call.Isddd() {

					arg = psess.nod(ODDDARG, nil, nil)
					arr := psess.types.NewArray(n.Type.Elem(psess.types), int64(len(args)))
					arg.Type = psess.types.NewPtr(arr)
					arg.Pos = call.Pos
					e.track(psess, arg)
					call.Right = arg
				}
				e.escassignWhyWhere(psess, n, arg, "arg to recursive call", call)
				if arg == args[0] {
					args = args[1:]
					continue
				}

				for _, a := range args {
					if psess.Debug['m'] > 3 {
						fmt.Printf("%v::esccall:: ... <- %S, untracked\n", psess.linestr(psess.lineno), a)
					}
					e.escassignSinkWhyWhere(psess, arg, a, "... arg to recursive call", call)
				}

				args = nil

			case PPARAMOUT:
				cE.Retval.Append(n)
			}
		}

		return
	}

	if cE.Retval.Len() != 0 {
		psess.
			Fatalf("esc already decorated call %+v\n", call)
	}

	if psess.Debug['m'] > 3 {
		fmt.Printf("%v::esccall:: %S not recursive\n", psess.linestr(psess.lineno), call)
	}

	e.initEscRetval(psess, call, fntype)

	if call.Op != OCALLFUNC {
		rf := fntype.Recv(psess.types)
		r := call.Left.Left
		if psess.types.Haspointers(rf.Type) {
			e.escassignfromtag(psess, rf.Note, cE.Retval, r, call)
		}
	}

	for i, param := range fntype.Params(psess.types).FieldSlice(psess.types) {
		note := param.Note
		var arg *Node
		if param.Isddd() && !call.Isddd() {
			rest := args[i:]
			if len(rest) == 0 {
				break
			}

			arg = psess.nod(ODDDARG, nil, nil)
			arg.Pos = call.Pos
			arr := psess.types.NewArray(param.Type.Elem(psess.types), int64(len(rest)))
			arg.Type = psess.types.NewPtr(arr)
			e.track(psess, arg)
			call.Right = arg

			for _, a := range rest {
				if psess.Debug['m'] > 3 {
					fmt.Printf("%v::esccall:: ... <- %S\n", psess.linestr(psess.lineno), a)
				}
				if note == uintptrEscapesTag {
					e.escassignSinkWhyWhere(psess, arg, a, "arg to uintptrescapes ...", call)
				} else {
					e.escassignWhyWhere(psess, arg, a, "arg to ...", call)
				}
			}
		} else {
			arg = args[i]
			if note == uintptrEscapesTag {
				e.escassignSinkWhy(psess, arg, arg, "escaping uintptr")
			}
		}

		if psess.types.Haspointers(param.Type) && e.escassignfromtag(psess, note, cE.Retval, arg, call)&EscMask == EscNone && parent.Op != ODEFER && parent.Op != OPROC {
			a := arg
			for a.Op == OCONVNOP {
				a = a.Left
			}
			switch a.Op {

			case OCALLPART, OCLOSURE, ODDDARG, OARRAYLIT, OSLICELIT, OPTRLIT, OSTRUCTLIT:
				a.SetNoescape(true)
			}
		}
	}
}

// escflows records the link src->dst in dst, throwing out some quick wins,
// and also ensuring that dst is noted as a flow destination.
func (e *EscState) escflows(psess *PackageSession, dst, src *Node, why *EscStep) {
	if dst == nil || src == nil || dst == src {
		return
	}

	if src.Type != nil && !psess.types.Haspointers(src.Type) && !psess.isReflectHeaderDataField(src) {
		if psess.Debug['m'] > 3 {
			fmt.Printf("%v::NOT flows:: %S <- %S\n", psess.linestr(psess.lineno), dst, src)
		}
		return
	}

	if psess.Debug['m'] > 3 {
		fmt.Printf("%v::flows:: %S <- %S\n", psess.linestr(psess.lineno), dst, src)
	}

	dstE := e.nodeEscState(psess, dst)
	if len(dstE.Flowsrc) == 0 {
		e.dsts = append(e.dsts, dst)
		e.dstcount++
	}

	e.edgecount++

	if why == nil {
		dstE.Flowsrc = append(dstE.Flowsrc, EscStep{src: src})
	} else {
		starwhy := *why
		starwhy.src = src
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
func (e *EscState) escflood(psess *PackageSession, dst *Node) {
	switch dst.Op {
	case ONAME, OCLOSURE:
	default:
		return
	}

	dstE := e.nodeEscState(psess, dst)
	if psess.Debug['m'] > 2 {
		fmt.Printf("\nescflood:%d: dst %S scope:%v[%d]\n", e.walkgen, dst, e.curfnSym(psess, dst), dstE.Loopdepth)
	}

	for i := range dstE.Flowsrc {
		e.walkgen++
		s := &dstE.Flowsrc[i]
		s.parent = nil
		e.escwalk(psess, levelFrom(0), dst, s.src, s)
	}
}

// funcOutputAndInput reports whether dst and src correspond to output and input parameters of the same function.
func funcOutputAndInput(dst, src *Node) bool {

	return dst.Op == ONAME && dst.Class() == PPARAMOUT &&
		src.Op == ONAME && src.Class() == PPARAM && src.Name.Curfn == dst.Name.Curfn
}

func (es *EscStep) describe(psess *PackageSession, src *Node) {
	if psess.Debug['m'] < 2 {
		return
	}
	step0 := es
	for step := step0; step != nil && !step.busy; step = step.parent {

		step.busy = true

		nextDest := step.parent
		dst := step.dst
		where := step.where
		if nextDest != nil {
			dst = nextDest.src
		}
		if where == nil {
			where = dst
		}
		psess.
			Warnl(src.Pos, "\tfrom %v (%s) at %s", dst, step.why, where.Line(psess))
	}
	for step := step0; step != nil && step.busy; step = step.parent {
		step.busy = false
	}
}

const NOTALOOPDEPTH = -1

func (e *EscState) escwalk(psess *PackageSession, level Level, dst *Node, src *Node, step *EscStep) {
	e.escwalkBody(psess, level, dst, src, step, NOTALOOPDEPTH)
}

func (e *EscState) escwalkBody(psess *PackageSession, level Level, dst *Node, src *Node, step *EscStep, extraloopdepth int32) {
	if src.Op == OLITERAL {
		return
	}
	srcE := e.nodeEscState(psess, src)
	if srcE.Walkgen == e.walkgen {

		level = level.min(srcE.Level)
		if level == srcE.Level {

			if srcE.Maxextraloopdepth >= extraloopdepth || srcE.Loopdepth >= extraloopdepth {
				return
			}
			srcE.Maxextraloopdepth = extraloopdepth
		}
	} else {
		srcE.Maxextraloopdepth = NOTALOOPDEPTH
	}

	srcE.Walkgen = e.walkgen
	srcE.Level = level
	modSrcLoopdepth := srcE.Loopdepth

	if extraloopdepth > modSrcLoopdepth {
		modSrcLoopdepth = extraloopdepth
	}

	if psess.Debug['m'] > 2 {
		fmt.Printf("escwalk: level:%d depth:%d %.*s op=%v %S(%0j) scope:%v[%d] extraloopdepth=%v\n",
			level, e.pdepth, e.pdepth, "\t\t\t\t\t\t\t\t\t\t", src.Op, src, src, e.curfnSym(psess, src), srcE.Loopdepth, extraloopdepth)
	}

	e.pdepth++

	// Input parameter flowing to output parameter?
	var leaks bool
	var osrcesc uint16 // used to prevent duplicate error messages

	dstE := e.nodeEscState(psess, dst)
	if funcOutputAndInput(dst, src) && src.Esc&EscMask < EscHeap && dst.Esc != EscHeap {

		if psess.Debug['m'] != 0 {
			if psess.Debug['m'] <= 2 {
				psess.
					Warnl(src.Pos, "leaking param: %S to result %v level=%v", src, dst.Sym, level.int())
				step.describe(psess, src)
			} else {
				psess.
					Warnl(src.Pos, "leaking param: %S to result %v level=%v", src, dst.Sym, level)
			}
		}
		if src.Esc&EscMask != EscReturn {
			src.Esc = EscReturn | src.Esc&EscContentEscapes
		}
		src.Esc = psess.escNoteOutputParamFlow(src.Esc, dst.Name.Vargen, level)
		goto recurse
	}

	if dst.Esc == EscHeap &&
		src.Op == ONAME && src.Class() == PPARAM && src.Esc&EscMask < EscHeap &&
		level.int() > 0 {
		src.Esc = psess.escMax(EscContentEscapes|src.Esc, EscNone)
		if psess.Debug['m'] != 0 {
			psess.
				Warnl(src.Pos, "mark escaped content: %S", src)
			step.describe(psess, src)
		}
	}

	leaks = level.int() <= 0 && level.guaranteedDereference() <= 0 && dstE.Loopdepth < modSrcLoopdepth
	leaks = leaks || level.int() <= 0 && dst.Esc&EscMask == EscHeap

	osrcesc = src.Esc
	switch src.Op {
	case ONAME:
		if src.Class() == PPARAM && (leaks || dstE.Loopdepth < 0) && src.Esc&EscMask < EscHeap {
			if level.guaranteedDereference() > 0 {
				src.Esc = psess.escMax(EscContentEscapes|src.Esc, EscNone)
				if psess.Debug['m'] != 0 {
					if psess.Debug['m'] <= 2 {
						if osrcesc != src.Esc {
							psess.
								Warnl(src.Pos, "leaking param content: %S", src)
							step.describe(psess, src)
						}
					} else {
						psess.
							Warnl(src.Pos, "leaking param content: %S level=%v dst.eld=%v src.eld=%v dst=%S",
								src, level, dstE.Loopdepth, modSrcLoopdepth, dst)
					}
				}
			} else {
				src.Esc = EscHeap
				if psess.Debug['m'] != 0 {
					if psess.Debug['m'] <= 2 {
						psess.
							Warnl(src.Pos, "leaking param: %S", src)
						step.describe(psess, src)
					} else {
						psess.
							Warnl(src.Pos, "leaking param: %S level=%v dst.eld=%v src.eld=%v dst=%S",
								src, level, dstE.Loopdepth, modSrcLoopdepth, dst)
					}
				}
			}
		}

		if src.IsClosureVar() {
			if leaks && psess.Debug['m'] != 0 {
				psess.
					Warnl(src.Pos, "leaking closure reference %S", src)
				step.describe(psess, src)
			}
			e.escwalk(psess, level, dst, src.Name.Defn, e.stepWalk(psess, dst, src.Name.Defn, "closure-var", step))
		}

	case OPTRLIT, OADDR:
		why := "pointer literal"
		if src.Op == OADDR {
			why = "address-of"
		}
		if leaks {
			src.Esc = EscHeap
			if psess.Debug['m'] != 0 && osrcesc != src.Esc {
				p := src
				if p.Left.Op == OCLOSURE {
					p = p.Left
				}
				if psess.Debug['m'] > 2 {
					psess.
						Warnl(src.Pos, "%S escapes to heap, level=%v, dst=%v dst.eld=%v, src.eld=%v",
							p, level, dst, dstE.Loopdepth, modSrcLoopdepth)
				} else {
					psess.
						Warnl(src.Pos, "%S escapes to heap", p)
					step.describe(psess, src)
				}
			}
			psess.
				addrescapes(src.Left)
			e.escwalkBody(psess, level.dec(), dst, src.Left, e.stepWalk(psess, dst, src.Left, why, step), modSrcLoopdepth)
			extraloopdepth = modSrcLoopdepth
		} else {
			e.escwalk(psess, level.dec(), dst, src.Left, e.stepWalk(psess, dst, src.Left, why, step))
		}

	case OAPPEND:
		e.escwalk(psess, level, dst, src.List.First(), e.stepWalk(psess, dst, src.List.First(), "append-first-arg", step))

	case ODDDARG:
		if leaks {
			src.Esc = EscHeap
			if psess.Debug['m'] != 0 && osrcesc != src.Esc {
				psess.
					Warnl(src.Pos, "%S escapes to heap", src)
				step.describe(psess, src)
			}
			extraloopdepth = modSrcLoopdepth
		}

		level = level.dec()

	case OSLICELIT:
		for _, elt := range src.List.Slice() {
			if elt.Op == OKEY {
				elt = elt.Right
			}
			e.escwalk(psess, level.dec(), dst, elt, e.stepWalk(psess, dst, elt, "slice-literal-element", step))
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
			if psess.Debug['m'] != 0 && osrcesc != src.Esc {
				psess.
					Warnl(src.Pos, "%S escapes to heap", src)
				step.describe(psess, src)
			}
			extraloopdepth = modSrcLoopdepth
		}

	case ODOT,
		ODOTTYPE:
		e.escwalk(psess, level, dst, src.Left, e.stepWalk(psess, dst, src.Left, "dot", step))

	case
		OSLICE,
		OSLICEARR,
		OSLICE3,
		OSLICE3ARR,
		OSLICESTR:
		e.escwalk(psess, level, dst, src.Left, e.stepWalk(psess, dst, src.Left, "slice", step))

	case OINDEX:
		if src.Left.Type.IsArray() {
			e.escwalk(psess, level, dst, src.Left, e.stepWalk(psess, dst, src.Left, "fixed-array-index-of", step))
			break
		}
		fallthrough

	case ODOTPTR:
		e.escwalk(psess, level.inc(), dst, src.Left, e.stepWalk(psess, dst, src.Left, "dot of pointer", step))
	case OINDEXMAP:
		e.escwalk(psess, level.inc(), dst, src.Left, e.stepWalk(psess, dst, src.Left, "map index", step))
	case OIND:
		e.escwalk(psess, level.inc(), dst, src.Left, e.stepWalk(psess, dst, src.Left, "indirection", step))

	case OCALLMETH, OCALLFUNC, OCALLINTER:
		if srcE.Retval.Len() != 0 {
			if psess.Debug['m'] > 2 {
				fmt.Printf("%v:[%d] dst %S escwalk replace src: %S with %S\n", psess.
					linestr(psess.lineno), e.loopdepth,
					dst, src, srcE.Retval.First())
			}
			src = srcE.Retval.First()
			srcE = e.nodeEscState(psess, src)
		}
	}

recurse:
	level = level.copy()

	for i := range srcE.Flowsrc {
		s := &srcE.Flowsrc[i]
		s.parent = step
		e.escwalkBody(psess, level, dst, s.src, s, extraloopdepth)
		s.parent = nil
	}

	e.pdepth--
}

// addrescapes tags node n as having had its address taken
// by "increasing" the "value" of n.Esc to EscHeap.
// Storage is allocated as necessary to allow the address
// to be taken.
func (psess *PackageSession) addrescapes(n *Node) {
	switch n.Op {
	default:

	case OIND, ODOTPTR:

	case ONAME:
		if n == psess.nodfp {
			break
		}

		if n.Class() == PAUTO && n.Esc == EscNever {
			break
		}

		if n.IsClosureVar() {
			psess.
				addrescapes(n.Name.Defn)
			break
		}

		if n.Class() != PPARAM && n.Class() != PPARAMOUT && n.Class() != PAUTO {
			break
		}

		oldfn := psess.Curfn
		psess.
			Curfn = n.Name.Curfn
		if psess.Curfn.Func.Closure != nil && psess.Curfn.Op == OCLOSURE {
			psess.
				Curfn = psess.Curfn.Func.Closure
		}
		ln := psess.lineno
		psess.
			lineno = psess.Curfn.Pos
		psess.
			moveToHeap(n)
		psess.
			Curfn = oldfn
		psess.
			lineno = ln

	case ODOT, OINDEX, OPAREN, OCONVNOP:
		if !n.Left.Type.IsSlice() {
			psess.
				addrescapes(n.Left)
		}
	}
}

// moveToHeap records the parameter or local variable n as moved to the heap.
func (psess *PackageSession) moveToHeap(n *Node) {
	if psess.Debug['r'] != 0 {
		Dump("MOVE", n)
	}
	if psess.compiling_runtime {
		psess.
			yyerror("%v escapes to heap, not allowed in runtime.", n)
	}
	if n.Class() == PAUTOHEAP {
		Dump("n", n)
		psess.
			Fatalf("double move to heap")
	}

	heapaddr := psess.temp(psess.types.NewPtr(n.Type))
	heapaddr.Sym = psess.lookup("&" + n.Sym.Name)
	heapaddr.Orig.Sym = heapaddr.Sym
	heapaddr.Pos = n.Pos

	heapaddr.Name.SetAutoTemp(false)

	if n.Class() == PPARAM || n.Class() == PPARAMOUT {
		if n.Xoffset == BADWIDTH {
			psess.
				Fatalf("addrescapes before param assignment")
		}

		stackcopy := psess.newname(n.Sym)
		stackcopy.SetAddable(false)
		stackcopy.Type = n.Type
		stackcopy.Xoffset = n.Xoffset
		stackcopy.SetClass(n.Class())
		stackcopy.Name.Param.Heapaddr = heapaddr
		if n.Class() == PPARAMOUT {

			heapaddr.SetIsOutputParamHeapAddr(true)
		}
		n.Name.Param.Stackcopy = stackcopy

		found := false
		for i, d := range psess.Curfn.Func.Dcl {
			if d == n {
				psess.
					Curfn.Func.Dcl[i] = stackcopy
				found = true
				break
			}

			if d.Class() == PAUTO {
				break
			}
		}
		if !found {
			psess.
				Fatalf("cannot find %v in local variable list", n)
		}
		psess.
			Curfn.Func.Dcl = append(psess.Curfn.Func.Dcl, n)
	}

	n.SetClass(PAUTOHEAP)
	n.Xoffset = 0
	n.Name.Param.Heapaddr = heapaddr
	n.Esc = EscHeap
	if psess.Debug['m'] != 0 {
		fmt.Printf("%v: moved to heap: %v\n", n.Line(psess), n)
	}
}

// This special tag is applied to uintptr variables
// that we believe may hold unsafe.Pointers for
// calls into assembly functions.
const unsafeUintptrTag = "unsafe-uintptr"

// This special tag is applied to uintptr parameters of functions
// marked go:uintptrescapes.
const uintptrEscapesTag = "uintptr-escapes"

func (e *EscState) esctag(psess *PackageSession, fn *Node) {
	fn.Esc = EscFuncTagged

	name := func(s *types.Sym, narg int) string {
		if s != nil {
			return s.Name
		}
		return fmt.Sprintf("arg#%d", narg)
	}

	if fn.Nbody.Len() == 0 {
		if fn.Noescape() {
			for _, f := range fn.Type.Params(psess.types).Fields(psess.types).Slice() {
				if psess.types.Haspointers(f.Type) {
					f.Note = psess.mktag(EscNone)
				}
			}
		}

		narg := 0
		for _, f := range fn.Type.Params(psess.types).Fields(psess.types).Slice() {
			narg++
			if f.Type.Etype == TUINTPTR {
				if psess.Debug['m'] != 0 {
					psess.
						Warnl(fn.Pos, "%v assuming %v is unsafe uintptr", funcSym(fn), name(f.Sym, narg))
				}
				f.Note = unsafeUintptrTag
			}
		}

		return
	}

	if fn.Func.Pragma&UintptrEscapes != 0 {
		narg := 0
		for _, f := range fn.Type.Params(psess.types).Fields(psess.types).Slice() {
			narg++
			if f.Type.Etype == TUINTPTR {
				if psess.Debug['m'] != 0 {
					psess.
						Warnl(fn.Pos, "%v marking %v as escaping uintptr", funcSym(fn), name(f.Sym, narg))
				}
				f.Note = uintptrEscapesTag
			}

			if f.Isddd() && f.Type.Elem(psess.types).Etype == TUINTPTR {

				if psess.Debug['m'] != 0 {
					psess.
						Warnl(fn.Pos, "%v marking %v as escaping ...uintptr", funcSym(fn), name(f.Sym, narg))
				}
				f.Note = uintptrEscapesTag
			}
		}
	}

	for _, fs := range psess.types.RecvsParams {
		for _, f := range fs(fn.Type).Fields(psess.types).Slice() {
			if !psess.types.Haspointers(f.Type) {
				continue
			}
			if f.Note == uintptrEscapesTag {

				continue
			}

			if f.Sym == nil || f.Sym.IsBlank() {
				f.Note = psess.mktag(EscNone)
				continue
			}

			switch esc := asNode(f.Nname).Esc; esc & EscMask {
			case EscNone,
				EscReturn:
				f.Note = psess.mktag(int(esc))

			case EscHeap:
			}
		}
	}
}

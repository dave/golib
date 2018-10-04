// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
)

// Static initialization ordering state.
// These values are stored in two bits in Node.flags.
const (
	InitNotStarted = iota
	InitDone
	InitPending
)

type InitEntry struct {
	Xoffset int64 // struct, array only
	Expr    *Node // bytes of run-time computed expressions
}

type InitPlan struct {
	E []InitEntry
}

// init1 walks the AST starting at n, and accumulates in out
// the list of definitions needing init code in dependency order.
func (pstate *PackageState) init1(n *Node, out *[]*Node) {
	if n == nil {
		return
	}
	pstate.init1(n.Left, out)
	pstate.init1(n.Right, out)
	for _, n1 := range n.List.Slice() {
		pstate.init1(n1, out)
	}

	if n.isMethodExpression() {
		// Methods called as Type.Method(receiver, ...).
		// Definitions for method expressions are stored in type->nname.
		pstate.init1(asNode(n.Type.FuncType(pstate.types).Nname), out)
	}

	if n.Op != ONAME {
		return
	}
	switch n.Class() {
	case PEXTERN, PFUNC:
	default:
		if n.isBlank() && n.Name.Curfn == nil && n.Name.Defn != nil && n.Name.Defn.Initorder() == InitNotStarted {
			// blank names initialization is part of init() but not
			// when they are inside a function.
			break
		}
		return
	}

	if n.Initorder() == InitDone {
		return
	}
	if n.Initorder() == InitPending {
		// Since mutually recursive sets of functions are allowed,
		// we don't necessarily raise an error if n depends on a node
		// which is already waiting for its dependencies to be visited.
		//
		// initlist contains a cycle of identifiers referring to each other.
		// If this cycle contains a variable, then this variable refers to itself.
		// Conversely, if there exists an initialization cycle involving
		// a variable in the program, the tree walk will reach a cycle
		// involving that variable.
		if n.Class() != PFUNC {
			pstate.foundinitloop(n, n)
		}

		for i := len(pstate.initlist) - 1; i >= 0; i-- {
			x := pstate.initlist[i]
			if x == n {
				break
			}
			if x.Class() != PFUNC {
				pstate.foundinitloop(n, x)
			}
		}

		// The loop involves only functions, ok.
		return
	}

	// reached a new unvisited node.
	n.SetInitorder(InitPending)
	pstate.initlist = append(pstate.initlist, n)

	// make sure that everything n depends on is initialized.
	// n->defn is an assignment to n
	if defn := n.Name.Defn; defn != nil {
		switch defn.Op {
		default:
			Dump("defn", defn)
			pstate.Fatalf("init1: bad defn")

		case ODCLFUNC:
			pstate.init2list(defn.Nbody, out)

		case OAS:
			if defn.Left != n {
				Dump("defn", defn)
				pstate.Fatalf("init1: bad defn")
			}
			if defn.Left.isBlank() && pstate.candiscard(defn.Right) {
				defn.Op = OEMPTY
				defn.Left = nil
				defn.Right = nil
				break
			}

			pstate.init2(defn.Right, out)
			if pstate.Debug['j'] != 0 {
				fmt.Printf("%v\n", n.Sym)
			}
			if n.isBlank() || !pstate.staticinit(n, out) {
				if pstate.Debug['%'] != 0 {
					Dump("nonstatic", defn)
				}
				*out = append(*out, defn)
			}

		case OAS2FUNC, OAS2MAPR, OAS2DOTTYPE, OAS2RECV:
			if defn.Initorder() == InitDone {
				break
			}
			defn.SetInitorder(InitPending)
			for _, n2 := range defn.Rlist.Slice() {
				pstate.init1(n2, out)
			}
			if pstate.Debug['%'] != 0 {
				Dump("nonstatic", defn)
			}
			*out = append(*out, defn)
			defn.SetInitorder(InitDone)
		}
	}

	last := len(pstate.initlist) - 1
	if pstate.initlist[last] != n {
		pstate.Fatalf("bad initlist %v", pstate.initlist)
	}
	pstate.initlist[last] = nil // allow GC
	pstate.initlist = pstate.initlist[:last]

	n.SetInitorder(InitDone)
}

// foundinitloop prints an init loop error and exits.
func (pstate *PackageState) foundinitloop(node, visited *Node) {
	// If there have already been errors printed,
	// those errors probably confused us and
	// there might not be a loop. Let the user
	// fix those first.
	pstate.flusherrors()
	if pstate.nerrors > 0 {
		pstate.errorexit()
	}

	// Find the index of node and visited in the initlist.
	var nodeindex, visitedindex int
	for ; pstate.initlist[nodeindex] != node; nodeindex++ {
	}
	for ; pstate.initlist[visitedindex] != visited; visitedindex++ {
	}

	// There is a loop involving visited. We know about node and
	// initlist = n1 <- ... <- visited <- ... <- node <- ...
	fmt.Printf("%v: initialization loop:\n", visited.Line(pstate))

	// Print visited -> ... -> n1 -> node.
	for _, n := range pstate.initlist[visitedindex:] {
		fmt.Printf("\t%v %v refers to\n", n.Line(pstate), n.Sym)
	}

	// Print node -> ... -> visited.
	for _, n := range pstate.initlist[nodeindex:visitedindex] {
		fmt.Printf("\t%v %v refers to\n", n.Line(pstate), n.Sym)
	}

	fmt.Printf("\t%v %v\n", visited.Line(pstate), visited.Sym)
	pstate.errorexit()
}

// recurse over n, doing init1 everywhere.
func (pstate *PackageState) init2(n *Node, out *[]*Node) {
	if n == nil || n.Initorder() == InitDone {
		return
	}

	if n.Op == ONAME && n.Ninit.Len() != 0 {
		pstate.Fatalf("name %v with ninit: %+v\n", n.Sym, n)
	}

	pstate.init1(n, out)
	pstate.init2(n.Left, out)
	pstate.init2(n.Right, out)
	pstate.init2list(n.Ninit, out)
	pstate.init2list(n.List, out)
	pstate.init2list(n.Rlist, out)
	pstate.init2list(n.Nbody, out)

	switch n.Op {
	case OCLOSURE:
		pstate.init2list(n.Func.Closure.Nbody, out)
	case ODOTMETH, OCALLPART:
		pstate.init2(asNode(n.Type.FuncType(pstate.types).Nname), out)
	}
}

func (pstate *PackageState) init2list(l Nodes, out *[]*Node) {
	for _, n := range l.Slice() {
		pstate.init2(n, out)
	}
}

func (pstate *PackageState) initreorder(l []*Node, out *[]*Node) {
	for _, n := range l {
		switch n.Op {
		case ODCLFUNC, ODCLCONST, ODCLTYPE:
			continue
		}

		pstate.initreorder(n.Ninit.Slice(), out)
		n.Ninit.Set(nil)
		pstate.init1(n, out)
	}
}

// initfix computes initialization order for a list l of top-level
// declarations and outputs the corresponding list of statements
// to include in the init() function body.
func (pstate *PackageState) initfix(l []*Node) []*Node {
	var lout []*Node
	pstate.initplans = make(map[*Node]*InitPlan)
	lno := pstate.lineno
	pstate.initreorder(l, &lout)
	pstate.lineno = lno
	pstate.initplans = nil
	return lout
}

// compilation of top-level (static) assignments
// into DATA statements if at all possible.
func (pstate *PackageState) staticinit(n *Node, out *[]*Node) bool {
	if n.Op != ONAME || n.Class() != PEXTERN || n.Name.Defn == nil || n.Name.Defn.Op != OAS {
		pstate.Fatalf("staticinit")
	}

	pstate.lineno = n.Pos
	l := n.Name.Defn.Left
	r := n.Name.Defn.Right
	return pstate.staticassign(l, r, out)
}

// like staticassign but we are copying an already
// initialized value r.
func (pstate *PackageState) staticcopy(l *Node, r *Node, out *[]*Node) bool {
	if r.Op != ONAME {
		return false
	}
	if r.Class() == PFUNC {
		pstate.gdata(l, r, pstate.Widthptr)
		return true
	}
	if r.Class() != PEXTERN || r.Sym.Pkg != pstate.localpkg {
		return false
	}
	if r.Name.Defn == nil { // probably zeroed but perhaps supplied externally and of unknown value
		return false
	}
	if r.Name.Defn.Op != OAS {
		return false
	}
	orig := r
	r = r.Name.Defn.Right

	for r.Op == OCONVNOP && !pstate.eqtype(r.Type, l.Type) {
		r = r.Left
	}

	switch r.Op {
	case ONAME:
		if pstate.staticcopy(l, r, out) {
			return true
		}
		// We may have skipped past one or more OCONVNOPs, so
		// use conv to ensure r is assignable to l (#13263).
		*out = append(*out, pstate.nod(OAS, l, pstate.conv(r, l.Type)))
		return true

	case OLITERAL:
		if pstate.isZero(r) {
			return true
		}
		pstate.gdata(l, r, int(l.Type.Width))
		return true

	case OADDR:
		switch r.Left.Op {
		case ONAME:
			pstate.gdata(l, r, int(l.Type.Width))
			return true
		}

	case OPTRLIT:
		switch r.Left.Op {
		case OARRAYLIT, OSLICELIT, OSTRUCTLIT, OMAPLIT:
			// copy pointer
			pstate.gdata(l, pstate.nod(OADDR, pstate.inittemps[r], nil), int(l.Type.Width))
			return true
		}

	case OSLICELIT:
		// copy slice
		a := pstate.inittemps[r]

		n := l.copy()
		n.Xoffset = l.Xoffset + int64(pstate.array_array)
		pstate.gdata(n, pstate.nod(OADDR, a, nil), pstate.Widthptr)
		n.Xoffset = l.Xoffset + int64(pstate.array_nel)
		pstate.gdata(n, r.Right, pstate.Widthptr)
		n.Xoffset = l.Xoffset + int64(pstate.array_cap)
		pstate.gdata(n, r.Right, pstate.Widthptr)
		return true

	case OARRAYLIT, OSTRUCTLIT:
		p := pstate.initplans[r]

		n := l.copy()
		for i := range p.E {
			e := &p.E[i]
			n.Xoffset = l.Xoffset + e.Xoffset
			n.Type = e.Expr.Type
			if e.Expr.Op == OLITERAL {
				pstate.gdata(n, e.Expr, int(n.Type.Width))
				continue
			}
			ll := n.copy()
			ll.Orig = ll // completely separate copy
			if pstate.staticassign(ll, e.Expr, out) {
				continue
			}
			// Requires computation, but we're
			// copying someone else's computation.
			rr := orig.copy()
			rr.Orig = rr // completely separate copy
			rr.Type = ll.Type
			rr.Xoffset += e.Xoffset
			pstate.setlineno(rr)
			*out = append(*out, pstate.nod(OAS, ll, rr))
		}

		return true
	}

	return false
}

func (pstate *PackageState) staticassign(l *Node, r *Node, out *[]*Node) bool {
	for r.Op == OCONVNOP {
		r = r.Left
	}

	switch r.Op {
	case ONAME:
		return pstate.staticcopy(l, r, out)

	case OLITERAL:
		if pstate.isZero(r) {
			return true
		}
		pstate.gdata(l, r, int(l.Type.Width))
		return true

	case OADDR:
		var nam Node
		if pstate.stataddr(&nam, r.Left) {
			n := *r
			n.Left = &nam
			pstate.gdata(l, &n, int(l.Type.Width))
			return true
		}
		fallthrough

	case OPTRLIT:
		switch r.Left.Op {
		case OARRAYLIT, OSLICELIT, OMAPLIT, OSTRUCTLIT:
			// Init pointer.
			a := pstate.staticname(r.Left.Type)

			pstate.inittemps[r] = a
			pstate.gdata(l, pstate.nod(OADDR, a, nil), int(l.Type.Width))

			// Init underlying literal.
			if !pstate.staticassign(a, r.Left, out) {
				*out = append(*out, pstate.nod(OAS, a, r.Left))
			}
			return true
		}
	//dump("not static ptrlit", r);

	case OSTRARRAYBYTE:
		if l.Class() == PEXTERN && r.Left.Op == OLITERAL {
			sval := r.Left.Val().U.(string)
			pstate.slicebytes(l, sval, len(sval))
			return true
		}

	case OSLICELIT:
		pstate.initplan(r)
		// Init slice.
		bound := r.Right.Int64(pstate)
		ta := pstate.types.NewArray(r.Type.Elem(pstate.types), bound)
		a := pstate.staticname(ta)
		pstate.inittemps[r] = a
		n := l.copy()
		n.Xoffset = l.Xoffset + int64(pstate.array_array)
		pstate.gdata(n, pstate.nod(OADDR, a, nil), pstate.Widthptr)
		n.Xoffset = l.Xoffset + int64(pstate.array_nel)
		pstate.gdata(n, r.Right, pstate.Widthptr)
		n.Xoffset = l.Xoffset + int64(pstate.array_cap)
		pstate.gdata(n, r.Right, pstate.Widthptr)

		// Fall through to init underlying array.
		l = a
		fallthrough

	case OARRAYLIT, OSTRUCTLIT:
		pstate.initplan(r)

		p := pstate.initplans[r]
		n := l.copy()
		for i := range p.E {
			e := &p.E[i]
			n.Xoffset = l.Xoffset + e.Xoffset
			n.Type = e.Expr.Type
			if e.Expr.Op == OLITERAL {
				pstate.gdata(n, e.Expr, int(n.Type.Width))
				continue
			}
			pstate.setlineno(e.Expr)
			a := n.copy()
			a.Orig = a // completely separate copy
			if !pstate.staticassign(a, e.Expr, out) {
				*out = append(*out, pstate.nod(OAS, a, e.Expr))
			}
		}

		return true

	case OMAPLIT:
		break

	case OCLOSURE:
		if hasemptycvars(r) {
			if pstate.Debug_closure > 0 {
				pstate.Warnl(r.Pos, "closure converted to global")
			}
			// Closures with no captured variables are globals,
			// so the assignment can be done at link time.
			pstate.gdata(l, r.Func.Closure.Func.Nname, pstate.Widthptr)
			return true
		}
		pstate.closuredebugruntimecheck(r)

	case OCONVIFACE:
		// This logic is mirrored in isStaticCompositeLiteral.
		// If you change something here, change it there, and vice versa.

		// Determine the underlying concrete type and value we are converting from.
		val := r
		for val.Op == OCONVIFACE {
			val = val.Left
		}
		if val.Type.IsInterface() {
			// val is an interface type.
			// If val is nil, we can statically initialize l;
			// both words are zero and so there no work to do, so report success.
			// If val is non-nil, we have no concrete type to record,
			// and we won't be able to statically initialize its value, so report failure.
			return pstate.Isconst(val, CTNIL)
		}

		var itab *Node
		if l.Type.IsEmptyInterface(pstate.types) {
			itab = pstate.typename(val.Type)
		} else {
			itab = pstate.itabname(val.Type, l.Type)
		}

		// Create a copy of l to modify while we emit data.
		n := l.copy()

		// Emit itab, advance offset.
		pstate.gdata(n, itab, pstate.Widthptr)
		n.Xoffset += int64(pstate.Widthptr)

		// Emit data.
		if pstate.isdirectiface(val.Type) {
			if pstate.Isconst(val, CTNIL) {
				// Nil is zero, nothing to do.
				return true
			}
			// Copy val directly into n.
			n.Type = val.Type
			pstate.setlineno(val)
			a := n.copy()
			a.Orig = a
			if !pstate.staticassign(a, val, out) {
				*out = append(*out, pstate.nod(OAS, a, val))
			}
		} else {
			// Construct temp to hold val, write pointer to temp into n.
			a := pstate.staticname(val.Type)
			pstate.inittemps[val] = a
			if !pstate.staticassign(a, val, out) {
				*out = append(*out, pstate.nod(OAS, a, val))
			}
			ptr := pstate.nod(OADDR, a, nil)
			n.Type = pstate.types.NewPtr(val.Type)
			pstate.gdata(n, ptr, pstate.Widthptr)
		}

		return true
	}

	//dump("not static", r);
	return false
}

// initContext is the context in which static data is populated.
// It is either in an init function or in any other function.
// Static data populated in an init function will be written either
// zero times (as a readonly, static data symbol) or
// one time (during init function execution).
// Either way, there is no opportunity for races or further modification,
// so the data can be written to a (possibly readonly) data symbol.
// Static data populated in any other function needs to be local to
// that function to allow multiple instances of that function
// to execute concurrently without clobbering each others' data.
type initContext uint8

const (
	inInitFunction initContext = iota
	inNonInitFunction
)

// staticname returns a name backed by a static data symbol.
// Callers should call n.Name.SetReadonly(true) on the
// returned node for readonly nodes.
func (pstate *PackageState) staticname(t *types.Type) *Node {
	// Don't use lookupN; it interns the resulting string, but these are all unique.
	n := pstate.newname(pstate.lookup(fmt.Sprintf("statictmp_%d", pstate.statuniqgen)))
	pstate.statuniqgen++
	pstate.addvar(n, t, PEXTERN)
	return n
}

func (pstate *PackageState) isLiteral(n *Node) bool {
	// Treat nils as zeros rather than literals.
	return n.Op == OLITERAL && n.Val().Ctype(pstate) != CTNIL
}

func (n *Node) isSimpleName() bool {
	return n.Op == ONAME && n.Addable() && n.Class() != PAUTOHEAP && n.Class() != PEXTERN
}

func (pstate *PackageState) litas(l *Node, r *Node, init *Nodes) {
	a := pstate.nod(OAS, l, r)
	a = pstate.typecheck(a, Etop)
	a = pstate.walkexpr(a, init)
	init.Append(a)
}

// initGenType is a bitmap indicating the types of generation that will occur for a static value.
type initGenType uint8

const (
	initDynamic initGenType = 1 << iota // contains some dynamic values, for which init code will be generated
	initConst                           // contains some constant values, which may be written into data symbols
)

// getdyn calculates the initGenType for n.
// If top is false, getdyn is recursing.
func (pstate *PackageState) getdyn(n *Node, top bool) initGenType {
	switch n.Op {
	default:
		if pstate.isLiteral(n) {
			return initConst
		}
		return initDynamic

	case OSLICELIT:
		if !top {
			return initDynamic
		}

	case OARRAYLIT, OSTRUCTLIT:
	}

	var mode initGenType
	for _, n1 := range n.List.Slice() {
		switch n1.Op {
		case OKEY:
			n1 = n1.Right
		case OSTRUCTKEY:
			n1 = n1.Left
		}
		mode |= pstate.getdyn(n1, false)
		if mode == initDynamic|initConst {
			break
		}
	}
	return mode
}

// isStaticCompositeLiteral reports whether n is a compile-time constant.
func (pstate *PackageState) isStaticCompositeLiteral(n *Node) bool {
	switch n.Op {
	case OSLICELIT:
		return false
	case OARRAYLIT:
		for _, r := range n.List.Slice() {
			if r.Op == OKEY {
				r = r.Right
			}
			if !pstate.isStaticCompositeLiteral(r) {
				return false
			}
		}
		return true
	case OSTRUCTLIT:
		for _, r := range n.List.Slice() {
			if r.Op != OSTRUCTKEY {
				pstate.Fatalf("isStaticCompositeLiteral: rhs not OSTRUCTKEY: %v", r)
			}
			if !pstate.isStaticCompositeLiteral(r.Left) {
				return false
			}
		}
		return true
	case OLITERAL:
		return true
	case OCONVIFACE:
		// See staticassign's OCONVIFACE case for comments.
		val := n
		for val.Op == OCONVIFACE {
			val = val.Left
		}
		if val.Type.IsInterface() {
			return pstate.Isconst(val, CTNIL)
		}
		if pstate.isdirectiface(val.Type) && pstate.Isconst(val, CTNIL) {
			return true
		}
		return pstate.isStaticCompositeLiteral(val)
	}
	return false
}

// initKind is a kind of static initialization: static, dynamic, or local.
// Static initialization represents literals and
// literal components of composite literals.
// Dynamic initialization represents non-literals and
// non-literal components of composite literals.
// LocalCode initializion represents initialization
// that occurs purely in generated code local to the function of use.
// Initialization code is sometimes generated in passes,
// first static then dynamic.
type initKind uint8

const (
	initKindStatic initKind = iota + 1
	initKindDynamic
	initKindLocalCode
)

// fixedlit handles struct, array, and slice literals.
// TODO: expand documentation.
func (pstate *PackageState) fixedlit(ctxt initContext, kind initKind, n *Node, var_ *Node, init *Nodes) {
	var splitnode func(*Node) (a *Node, value *Node)
	switch n.Op {
	case OARRAYLIT, OSLICELIT:
		var k int64
		splitnode = func(r *Node) (*Node, *Node) {
			if r.Op == OKEY {
				k = pstate.nonnegintconst(r.Left)
				r = r.Right
			}
			a := pstate.nod(OINDEX, var_, pstate.nodintconst(k))
			k++
			return a, r
		}
	case OSTRUCTLIT:
		splitnode = func(r *Node) (*Node, *Node) {
			if r.Op != OSTRUCTKEY {
				pstate.Fatalf("fixedlit: rhs not OSTRUCTKEY: %v", r)
			}
			if r.Sym.IsBlank() {
				return pstate.nblank, r.Left
			}
			return pstate.nodSym(ODOT, var_, r.Sym), r.Left
		}
	default:
		pstate.Fatalf("fixedlit bad op: %v", n.Op)
	}

	for _, r := range n.List.Slice() {
		a, value := splitnode(r)

		switch value.Op {
		case OSLICELIT:
			if (kind == initKindStatic && ctxt == inNonInitFunction) || (kind == initKindDynamic && ctxt == inInitFunction) {
				pstate.slicelit(ctxt, value, a, init)
				continue
			}

		case OARRAYLIT, OSTRUCTLIT:
			pstate.fixedlit(ctxt, kind, value, a, init)
			continue
		}

		islit := pstate.isLiteral(value)
		if (kind == initKindStatic && !islit) || (kind == initKindDynamic && islit) {
			continue
		}

		// build list of assignments: var[index] = expr
		pstate.setlineno(value)
		a = pstate.nod(OAS, a, value)
		a = pstate.typecheck(a, Etop)
		switch kind {
		case initKindStatic:
			pstate.genAsStatic(a)
		case initKindDynamic, initKindLocalCode:
			a = pstate.orderStmtInPlace(a)
			a = pstate.walkstmt(a)
			init.Append(a)
		default:
			pstate.Fatalf("fixedlit: bad kind %d", kind)
		}

	}
}

func (pstate *PackageState) slicelit(ctxt initContext, n *Node, var_ *Node, init *Nodes) {
	// make an array type corresponding the number of elements we have
	t := pstate.types.NewArray(n.Type.Elem(pstate.types), n.Right.Int64(pstate))
	pstate.dowidth(t)

	if ctxt == inNonInitFunction {
		// put everything into static array
		vstat := pstate.staticname(t)

		pstate.fixedlit(ctxt, initKindStatic, n, vstat, init)
		pstate.fixedlit(ctxt, initKindDynamic, n, vstat, init)

		// copy static to slice
		var_ = pstate.typecheck(var_, Erv|Easgn)
		var nam Node
		if !pstate.stataddr(&nam, var_) || nam.Class() != PEXTERN {
			pstate.Fatalf("slicelit: %v", var_)
		}

		var v Node
		v.Type = pstate.types.Types[TINT]
		pstate.setintconst(&v, t.NumElem(pstate.types))

		nam.Xoffset += int64(pstate.array_array)
		pstate.gdata(&nam, pstate.nod(OADDR, vstat, nil), pstate.Widthptr)
		nam.Xoffset += int64(pstate.array_nel) - int64(pstate.array_array)
		pstate.gdata(&nam, &v, pstate.Widthptr)
		nam.Xoffset += int64(pstate.array_cap) - int64(pstate.array_nel)
		pstate.gdata(&nam, &v, pstate.Widthptr)

		return
	}

	// recipe for var = []t{...}
	// 1. make a static array
	//	var vstat [...]t
	// 2. assign (data statements) the constant part
	//	vstat = constpart{}
	// 3. make an auto pointer to array and allocate heap to it
	//	var vauto *[...]t = new([...]t)
	// 4. copy the static array to the auto array
	//	*vauto = vstat
	// 5. for each dynamic part assign to the array
	//	vauto[i] = dynamic part
	// 6. assign slice of allocated heap to var
	//	var = vauto[:]
	//
	// an optimization is done if there is no constant part
	//	3. var vauto *[...]t = new([...]t)
	//	5. vauto[i] = dynamic part
	//	6. var = vauto[:]

	// if the literal contains constants,
	// make static initialized array (1),(2)
	var vstat *Node

	mode := pstate.getdyn(n, true)
	if mode&initConst != 0 {
		vstat = pstate.staticname(t)
		if ctxt == inInitFunction {
			vstat.Name.SetReadonly(true)
		}
		pstate.fixedlit(ctxt, initKindStatic, n, vstat, init)
	}

	// make new auto *array (3 declare)
	vauto := pstate.temp(pstate.types.NewPtr(t))

	// set auto to point at new temp or heap (3 assign)
	var a *Node
	if x := pstate.prealloc[n]; x != nil {
		// temp allocated during order.go for dddarg
		x.Type = t

		if vstat == nil {
			a = pstate.nod(OAS, x, nil)
			a = pstate.typecheck(a, Etop)
			init.Append(a) // zero new temp
		}

		a = pstate.nod(OADDR, x, nil)
	} else if n.Esc == EscNone {
		a = pstate.temp(t)
		if vstat == nil {
			a = pstate.nod(OAS, pstate.temp(t), nil)
			a = pstate.typecheck(a, Etop)
			init.Append(a) // zero new temp
			a = a.Left
		}

		a = pstate.nod(OADDR, a, nil)
	} else {
		a = pstate.nod(ONEW, nil, nil)
		a.List.Set1(pstate.typenod(t))
	}

	a = pstate.nod(OAS, vauto, a)
	a = pstate.typecheck(a, Etop)
	a = pstate.walkexpr(a, init)
	init.Append(a)

	if vstat != nil {
		// copy static to heap (4)
		a = pstate.nod(OIND, vauto, nil)

		a = pstate.nod(OAS, a, vstat)
		a = pstate.typecheck(a, Etop)
		a = pstate.walkexpr(a, init)
		init.Append(a)
	}

	// put dynamics into array (5)
	var index int64
	for _, value := range n.List.Slice() {
		if value.Op == OKEY {
			index = pstate.nonnegintconst(value.Left)
			value = value.Right
		}
		a := pstate.nod(OINDEX, vauto, pstate.nodintconst(index))
		a.SetBounded(true)
		index++

		// TODO need to check bounds?

		switch value.Op {
		case OSLICELIT:
			break

		case OARRAYLIT, OSTRUCTLIT:
			pstate.fixedlit(ctxt, initKindDynamic, value, a, init)
			continue
		}

		if pstate.isLiteral(value) {
			continue
		}

		// build list of vauto[c] = expr
		pstate.setlineno(value)
		a = pstate.nod(OAS, a, value)

		a = pstate.typecheck(a, Etop)
		a = pstate.orderStmtInPlace(a)
		a = pstate.walkstmt(a)
		init.Append(a)
	}

	// make slice out of heap (6)
	a = pstate.nod(OAS, var_, pstate.nod(OSLICE, vauto, nil))

	a = pstate.typecheck(a, Etop)
	a = pstate.orderStmtInPlace(a)
	a = pstate.walkstmt(a)
	init.Append(a)
}

func (pstate *PackageState) maplit(n *Node, m *Node, init *Nodes) {
	// make the map var
	a := pstate.nod(OMAKE, nil, nil)
	a.Esc = n.Esc
	a.List.Set2(pstate.typenod(n.Type), pstate.nodintconst(int64(n.List.Len())))
	pstate.litas(m, a, init)

	// Split the initializers into static and dynamic.
	var stat, dyn []*Node
	for _, r := range n.List.Slice() {
		if r.Op != OKEY {
			pstate.Fatalf("maplit: rhs not OKEY: %v", r)
		}
		if pstate.isStaticCompositeLiteral(r.Left) && pstate.isStaticCompositeLiteral(r.Right) {
			stat = append(stat, r)
		} else {
			dyn = append(dyn, r)
		}
	}

	// Add static entries.
	if len(stat) > 25 {
		// For a large number of static entries, put them in an array and loop.

		// build types [count]Tindex and [count]Tvalue
		tk := pstate.types.NewArray(n.Type.Key(pstate.types), int64(len(stat)))
		tv := pstate.types.NewArray(n.Type.Elem(pstate.types), int64(len(stat)))

		// TODO(josharian): suppress alg generation for these types?
		pstate.dowidth(tk)
		pstate.dowidth(tv)

		// make and initialize static arrays
		vstatk := pstate.staticname(tk)
		vstatk.Name.SetReadonly(true)
		vstatv := pstate.staticname(tv)
		vstatv.Name.SetReadonly(true)

		datak := pstate.nod(OARRAYLIT, nil, nil)
		datav := pstate.nod(OARRAYLIT, nil, nil)
		for _, r := range stat {
			datak.List.Append(r.Left)
			datav.List.Append(r.Right)
		}
		pstate.fixedlit(inInitFunction, initKindStatic, datak, vstatk, init)
		pstate.fixedlit(inInitFunction, initKindStatic, datav, vstatv, init)

		// loop adding structure elements to map
		// for i = 0; i < len(vstatk); i++ {
		//	map[vstatk[i]] = vstatv[i]
		// }
		i := pstate.temp(pstate.types.Types[TINT])
		rhs := pstate.nod(OINDEX, vstatv, i)
		rhs.SetBounded(true)

		kidx := pstate.nod(OINDEX, vstatk, i)
		kidx.SetBounded(true)
		lhs := pstate.nod(OINDEX, m, kidx)

		zero := pstate.nod(OAS, i, pstate.nodintconst(0))
		cond := pstate.nod(OLT, i, pstate.nodintconst(tk.NumElem(pstate.types)))
		incr := pstate.nod(OAS, i, pstate.nod(OADD, i, pstate.nodintconst(1)))
		body := pstate.nod(OAS, lhs, rhs)

		loop := pstate.nod(OFOR, cond, incr)
		loop.Nbody.Set1(body)
		loop.Ninit.Set1(zero)

		loop = pstate.typecheck(loop, Etop)
		loop = pstate.walkstmt(loop)
		init.Append(loop)
	} else {
		// For a small number of static entries, just add them directly.
		pstate.addMapEntries(m, stat, init)
	}

	// Add dynamic entries.
	pstate.addMapEntries(m, dyn, init)
}

func (pstate *PackageState) addMapEntries(m *Node, dyn []*Node, init *Nodes) {
	if len(dyn) == 0 {
		return
	}

	nerr := pstate.nerrors

	// Build list of var[c] = expr.
	// Use temporaries so that mapassign1 can have addressable key, val.
	// TODO(josharian): avoid map key temporaries for mapfast_* assignments with literal keys.
	key := pstate.temp(m.Type.Key(pstate.types))
	val := pstate.temp(m.Type.Elem(pstate.types))

	for _, r := range dyn {
		index, value := r.Left, r.Right

		pstate.setlineno(index)
		a := pstate.nod(OAS, key, index)
		a = pstate.typecheck(a, Etop)
		a = pstate.walkstmt(a)
		init.Append(a)

		pstate.setlineno(value)
		a = pstate.nod(OAS, val, value)
		a = pstate.typecheck(a, Etop)
		a = pstate.walkstmt(a)
		init.Append(a)

		pstate.setlineno(val)
		a = pstate.nod(OAS, pstate.nod(OINDEX, m, key), val)
		a = pstate.typecheck(a, Etop)
		a = pstate.walkstmt(a)
		init.Append(a)

		if nerr != pstate.nerrors {
			break
		}
	}

	a := pstate.nod(OVARKILL, key, nil)
	a = pstate.typecheck(a, Etop)
	init.Append(a)
	a = pstate.nod(OVARKILL, val, nil)
	a = pstate.typecheck(a, Etop)
	init.Append(a)
}

func (pstate *PackageState) anylit(n *Node, var_ *Node, init *Nodes) {
	t := n.Type
	switch n.Op {
	default:
		pstate.Fatalf("anylit: not lit, op=%v node=%v", n.Op, n)

	case OPTRLIT:
		if !t.IsPtr() {
			pstate.Fatalf("anylit: not ptr")
		}

		var r *Node
		if n.Right != nil {
			// n.Right is stack temporary used as backing store.
			init.Append(pstate.nod(OAS, n.Right, nil)) // zero backing store, just in case (#18410)
			r = pstate.nod(OADDR, n.Right, nil)
			r = pstate.typecheck(r, Erv)
		} else {
			r = pstate.nod(ONEW, nil, nil)
			r.SetTypecheck(1)
			r.Type = t
			r.Esc = n.Esc
		}

		r = pstate.walkexpr(r, init)
		a := pstate.nod(OAS, var_, r)

		a = pstate.typecheck(a, Etop)
		init.Append(a)

		var_ = pstate.nod(OIND, var_, nil)
		var_ = pstate.typecheck(var_, Erv|Easgn)
		pstate.anylit(n.Left, var_, init)

	case OSTRUCTLIT, OARRAYLIT:
		if !t.IsStruct() && !t.IsArray() {
			pstate.Fatalf("anylit: not struct/array")
		}

		if var_.isSimpleName() && n.List.Len() > 4 {
			// lay out static data
			vstat := pstate.staticname(t)
			vstat.Name.SetReadonly(true)

			ctxt := inInitFunction
			if n.Op == OARRAYLIT {
				ctxt = inNonInitFunction
			}
			pstate.fixedlit(ctxt, initKindStatic, n, vstat, init)

			// copy static to var
			a := pstate.nod(OAS, var_, vstat)

			a = pstate.typecheck(a, Etop)
			a = pstate.walkexpr(a, init)
			init.Append(a)

			// add expressions to automatic
			pstate.fixedlit(inInitFunction, initKindDynamic, n, var_, init)
			break
		}

		var components int64
		if n.Op == OARRAYLIT {
			components = t.NumElem(pstate.types)
		} else {
			components = int64(t.NumFields(pstate.types))
		}
		// initialization of an array or struct with unspecified components (missing fields or arrays)
		if var_.isSimpleName() || int64(n.List.Len()) < components {
			a := pstate.nod(OAS, var_, nil)
			a = pstate.typecheck(a, Etop)
			a = pstate.walkexpr(a, init)
			init.Append(a)
		}

		pstate.fixedlit(inInitFunction, initKindLocalCode, n, var_, init)

	case OSLICELIT:
		pstate.slicelit(inInitFunction, n, var_, init)

	case OMAPLIT:
		if !t.IsMap() {
			pstate.Fatalf("anylit: not map")
		}
		pstate.maplit(n, var_, init)
	}
}

func (pstate *PackageState) oaslit(n *Node, init *Nodes) bool {
	if n.Left == nil || n.Right == nil {
		// not a special composite literal assignment
		return false
	}
	if n.Left.Type == nil || n.Right.Type == nil {
		// not a special composite literal assignment
		return false
	}
	if !n.Left.isSimpleName() {
		// not a special composite literal assignment
		return false
	}
	if !pstate.eqtype(n.Left.Type, n.Right.Type) {
		// not a special composite literal assignment
		return false
	}

	switch n.Right.Op {
	default:
		// not a special composite literal assignment
		return false

	case OSTRUCTLIT, OARRAYLIT, OSLICELIT, OMAPLIT:
		if vmatch1(n.Left, n.Right) {
			// not a special composite literal assignment
			return false
		}
		pstate.anylit(n.Right, n.Left, init)
	}

	n.Op = OEMPTY
	n.Right = nil
	return true
}

func (pstate *PackageState) getlit(lit *Node) int {
	if pstate.smallintconst(lit) {
		return int(lit.Int64(pstate))
	}
	return -1
}

// stataddr sets nam to the static address of n and reports whether it succeeded.
func (pstate *PackageState) stataddr(nam *Node, n *Node) bool {
	if n == nil {
		return false
	}

	switch n.Op {
	case ONAME:
		*nam = *n
		return n.Addable()

	case ODOT:
		if !pstate.stataddr(nam, n.Left) {
			break
		}
		nam.Xoffset += n.Xoffset
		nam.Type = n.Type
		return true

	case OINDEX:
		if n.Left.Type.IsSlice() {
			break
		}
		if !pstate.stataddr(nam, n.Left) {
			break
		}
		l := pstate.getlit(n.Right)
		if l < 0 {
			break
		}

		// Check for overflow.
		if n.Type.Width != 0 && pstate.thearch.MAXWIDTH/n.Type.Width <= int64(l) {
			break
		}
		nam.Xoffset += int64(l) * n.Type.Width
		nam.Type = n.Type
		return true
	}

	return false
}

func (pstate *PackageState) initplan(n *Node) {
	if pstate.initplans[n] != nil {
		return
	}
	p := new(InitPlan)
	pstate.initplans[n] = p
	switch n.Op {
	default:
		pstate.Fatalf("initplan")

	case OARRAYLIT, OSLICELIT:
		var k int64
		for _, a := range n.List.Slice() {
			if a.Op == OKEY {
				k = pstate.nonnegintconst(a.Left)
				a = a.Right
			}
			pstate.addvalue(p, k*n.Type.Elem(pstate.types).Width, a)
			k++
		}

	case OSTRUCTLIT:
		for _, a := range n.List.Slice() {
			if a.Op != OSTRUCTKEY {
				pstate.Fatalf("initplan fixedlit")
			}
			pstate.addvalue(p, a.Xoffset, a.Left)
		}

	case OMAPLIT:
		for _, a := range n.List.Slice() {
			if a.Op != OKEY {
				pstate.Fatalf("initplan maplit")
			}
			pstate.addvalue(p, -1, a.Right)
		}
	}
}

func (pstate *PackageState) addvalue(p *InitPlan, xoffset int64, n *Node) {
	// special case: zero can be dropped entirely
	if pstate.isZero(n) {
		return
	}

	// special case: inline struct and array (not slice) literals
	if isvaluelit(n) {
		pstate.initplan(n)
		q := pstate.initplans[n]
		for _, qe := range q.E {
			// qe is a copy; we are not modifying entries in q.E
			qe.Xoffset += xoffset
			p.E = append(p.E, qe)
		}
		return
	}

	// add to plan
	p.E = append(p.E, InitEntry{Xoffset: xoffset, Expr: n})
}

func (pstate *PackageState) isZero(n *Node) bool {
	switch n.Op {
	case OLITERAL:
		switch u := n.Val().U.(type) {
		default:
			Dump("unexpected literal", n)
			pstate.Fatalf("isZero")
		case *NilVal:
			return true
		case string:
			return u == ""
		case bool:
			return !u
		case *Mpint:
			return u.CmpInt64(0) == 0
		case *Mpflt:
			return u.CmpFloat64(0) == 0
		case *Mpcplx:
			return u.Real.CmpFloat64(0) == 0 && u.Imag.CmpFloat64(0) == 0
		}

	case OARRAYLIT:
		for _, n1 := range n.List.Slice() {
			if n1.Op == OKEY {
				n1 = n1.Right
			}
			if !pstate.isZero(n1) {
				return false
			}
		}
		return true

	case OSTRUCTLIT:
		for _, n1 := range n.List.Slice() {
			if !pstate.isZero(n1.Left) {
				return false
			}
		}
		return true
	}

	return false
}

func isvaluelit(n *Node) bool {
	return n.Op == OARRAYLIT || n.Op == OSTRUCTLIT
}

func (pstate *PackageState) genAsStatic(as *Node) {
	if as.Left.Type == nil {
		pstate.Fatalf("genAsStatic as.Left not typechecked")
	}

	var nam Node
	if !pstate.stataddr(&nam, as.Left) || (nam.Class() != PEXTERN && as.Left != pstate.nblank) {
		pstate.Fatalf("genAsStatic: lhs %v", as.Left)
	}

	switch {
	case as.Right.Op == OLITERAL:
	case as.Right.Op == ONAME && as.Right.Class() == PFUNC:
	default:
		pstate.Fatalf("genAsStatic: rhs %v", as.Right)
	}

	pstate.gdata(&nam, as.Right, int(as.Right.Type.Width))
}

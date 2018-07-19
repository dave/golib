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
func (psess *PackageSession) init1(n *Node, out *[]*Node) {
	if n == nil {
		return
	}
	psess.
		init1(n.Left, out)
	psess.
		init1(n.Right, out)
	for _, n1 := range n.List.Slice() {
		psess.
			init1(n1, out)
	}

	if n.isMethodExpression() {
		psess.
			init1(asNode(n.Type.FuncType(psess.types).Nname), out)
	}

	if n.Op != ONAME {
		return
	}
	switch n.Class() {
	case PEXTERN, PFUNC:
	default:
		if n.isBlank() && n.Name.Curfn == nil && n.Name.Defn != nil && n.Name.Defn.Initorder() == InitNotStarted {

			break
		}
		return
	}

	if n.Initorder() == InitDone {
		return
	}
	if n.Initorder() == InitPending {

		if n.Class() != PFUNC {
			psess.
				foundinitloop(n, n)
		}

		for i := len(psess.initlist) - 1; i >= 0; i-- {
			x := psess.initlist[i]
			if x == n {
				break
			}
			if x.Class() != PFUNC {
				psess.
					foundinitloop(n, x)
			}
		}

		return
	}

	n.SetInitorder(InitPending)
	psess.
		initlist = append(psess.initlist, n)

	if defn := n.Name.Defn; defn != nil {
		switch defn.Op {
		default:
			Dump("defn", defn)
			psess.
				Fatalf("init1: bad defn")

		case ODCLFUNC:
			psess.
				init2list(defn.Nbody, out)

		case OAS:
			if defn.Left != n {
				Dump("defn", defn)
				psess.
					Fatalf("init1: bad defn")
			}
			if defn.Left.isBlank() && psess.candiscard(defn.Right) {
				defn.Op = OEMPTY
				defn.Left = nil
				defn.Right = nil
				break
			}
			psess.
				init2(defn.Right, out)
			if psess.Debug['j'] != 0 {
				fmt.Printf("%v\n", n.Sym)
			}
			if n.isBlank() || !psess.staticinit(n, out) {
				if psess.Debug['%'] != 0 {
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
				psess.
					init1(n2, out)
			}
			if psess.Debug['%'] != 0 {
				Dump("nonstatic", defn)
			}
			*out = append(*out, defn)
			defn.SetInitorder(InitDone)
		}
	}

	last := len(psess.initlist) - 1
	if psess.initlist[last] != n {
		psess.
			Fatalf("bad initlist %v", psess.initlist)
	}
	psess.
		initlist[last] = nil
	psess.
		initlist = psess.initlist[:last]

	n.SetInitorder(InitDone)
}

// foundinitloop prints an init loop error and exits.
func (psess *PackageSession) foundinitloop(node, visited *Node) {
	psess.
		flusherrors()
	if psess.nerrors > 0 {
		psess.
			errorexit()
	}

	// Find the index of node and visited in the initlist.
	var nodeindex, visitedindex int
	for ; psess.initlist[nodeindex] != node; nodeindex++ {
	}
	for ; psess.initlist[visitedindex] != visited; visitedindex++ {
	}

	fmt.Printf("%v: initialization loop:\n", visited.Line(psess))

	for _, n := range psess.initlist[visitedindex:] {
		fmt.Printf("\t%v %v refers to\n", n.Line(psess), n.Sym)
	}

	for _, n := range psess.initlist[nodeindex:visitedindex] {
		fmt.Printf("\t%v %v refers to\n", n.Line(psess), n.Sym)
	}

	fmt.Printf("\t%v %v\n", visited.Line(psess), visited.Sym)
	psess.
		errorexit()
}

// recurse over n, doing init1 everywhere.
func (psess *PackageSession) init2(n *Node, out *[]*Node) {
	if n == nil || n.Initorder() == InitDone {
		return
	}

	if n.Op == ONAME && n.Ninit.Len() != 0 {
		psess.
			Fatalf("name %v with ninit: %+v\n", n.Sym, n)
	}
	psess.
		init1(n, out)
	psess.
		init2(n.Left, out)
	psess.
		init2(n.Right, out)
	psess.
		init2list(n.Ninit, out)
	psess.
		init2list(n.List, out)
	psess.
		init2list(n.Rlist, out)
	psess.
		init2list(n.Nbody, out)

	switch n.Op {
	case OCLOSURE:
		psess.
			init2list(n.Func.Closure.Nbody, out)
	case ODOTMETH, OCALLPART:
		psess.
			init2(asNode(n.Type.FuncType(psess.types).Nname), out)
	}
}

func (psess *PackageSession) init2list(l Nodes, out *[]*Node) {
	for _, n := range l.Slice() {
		psess.
			init2(n, out)
	}
}

func (psess *PackageSession) initreorder(l []*Node, out *[]*Node) {
	for _, n := range l {
		switch n.Op {
		case ODCLFUNC, ODCLCONST, ODCLTYPE:
			continue
		}
		psess.
			initreorder(n.Ninit.Slice(), out)
		n.Ninit.Set(nil)
		psess.
			init1(n, out)
	}
}

// initfix computes initialization order for a list l of top-level
// declarations and outputs the corresponding list of statements
// to include in the init() function body.
func (psess *PackageSession) initfix(l []*Node) []*Node {
	var lout []*Node
	psess.
		initplans = make(map[*Node]*InitPlan)
	lno := psess.lineno
	psess.
		initreorder(l, &lout)
	psess.
		lineno = lno
	psess.
		initplans = nil
	return lout
}

// compilation of top-level (static) assignments
// into DATA statements if at all possible.
func (psess *PackageSession) staticinit(n *Node, out *[]*Node) bool {
	if n.Op != ONAME || n.Class() != PEXTERN || n.Name.Defn == nil || n.Name.Defn.Op != OAS {
		psess.
			Fatalf("staticinit")
	}
	psess.
		lineno = n.Pos
	l := n.Name.Defn.Left
	r := n.Name.Defn.Right
	return psess.staticassign(l, r, out)
}

// like staticassign but we are copying an already
// initialized value r.
func (psess *PackageSession) staticcopy(l *Node, r *Node, out *[]*Node) bool {
	if r.Op != ONAME {
		return false
	}
	if r.Class() == PFUNC {
		psess.
			gdata(l, r, psess.Widthptr)
		return true
	}
	if r.Class() != PEXTERN || r.Sym.Pkg != psess.localpkg {
		return false
	}
	if r.Name.Defn == nil {
		return false
	}
	if r.Name.Defn.Op != OAS {
		return false
	}
	orig := r
	r = r.Name.Defn.Right

	for r.Op == OCONVNOP && !psess.eqtype(r.Type, l.Type) {
		r = r.Left
	}

	switch r.Op {
	case ONAME:
		if psess.staticcopy(l, r, out) {
			return true
		}

		*out = append(*out, psess.nod(OAS, l, psess.conv(r, l.Type)))
		return true

	case OLITERAL:
		if psess.isZero(r) {
			return true
		}
		psess.
			gdata(l, r, int(l.Type.Width))
		return true

	case OADDR:
		switch r.Left.Op {
		case ONAME:
			psess.
				gdata(l, r, int(l.Type.Width))
			return true
		}

	case OPTRLIT:
		switch r.Left.Op {
		case OARRAYLIT, OSLICELIT, OSTRUCTLIT, OMAPLIT:
			psess.
				gdata(l, psess.nod(OADDR, psess.inittemps[r], nil), int(l.Type.Width))
			return true
		}

	case OSLICELIT:

		a := psess.inittemps[r]

		n := l.copy()
		n.Xoffset = l.Xoffset + int64(psess.array_array)
		psess.
			gdata(n, psess.nod(OADDR, a, nil), psess.Widthptr)
		n.Xoffset = l.Xoffset + int64(psess.array_nel)
		psess.
			gdata(n, r.Right, psess.Widthptr)
		n.Xoffset = l.Xoffset + int64(psess.array_cap)
		psess.
			gdata(n, r.Right, psess.Widthptr)
		return true

	case OARRAYLIT, OSTRUCTLIT:
		p := psess.initplans[r]

		n := l.copy()
		for i := range p.E {
			e := &p.E[i]
			n.Xoffset = l.Xoffset + e.Xoffset
			n.Type = e.Expr.Type
			if e.Expr.Op == OLITERAL {
				psess.
					gdata(n, e.Expr, int(n.Type.Width))
				continue
			}
			ll := n.copy()
			ll.Orig = ll
			if psess.staticassign(ll, e.Expr, out) {
				continue
			}

			rr := orig.copy()
			rr.Orig = rr
			rr.Type = ll.Type
			rr.Xoffset += e.Xoffset
			psess.
				setlineno(rr)
			*out = append(*out, psess.nod(OAS, ll, rr))
		}

		return true
	}

	return false
}

func (psess *PackageSession) staticassign(l *Node, r *Node, out *[]*Node) bool {
	for r.Op == OCONVNOP {
		r = r.Left
	}

	switch r.Op {
	case ONAME:
		return psess.staticcopy(l, r, out)

	case OLITERAL:
		if psess.isZero(r) {
			return true
		}
		psess.
			gdata(l, r, int(l.Type.Width))
		return true

	case OADDR:
		var nam Node
		if psess.stataddr(&nam, r.Left) {
			n := *r
			n.Left = &nam
			psess.
				gdata(l, &n, int(l.Type.Width))
			return true
		}
		fallthrough

	case OPTRLIT:
		switch r.Left.Op {
		case OARRAYLIT, OSLICELIT, OMAPLIT, OSTRUCTLIT:

			a := psess.staticname(r.Left.Type)
			psess.
				inittemps[r] = a
			psess.
				gdata(l, psess.nod(OADDR, a, nil), int(l.Type.Width))

			if !psess.staticassign(a, r.Left, out) {
				*out = append(*out, psess.nod(OAS, a, r.Left))
			}
			return true
		}

	case OSTRARRAYBYTE:
		if l.Class() == PEXTERN && r.Left.Op == OLITERAL {
			sval := r.Left.Val().U.(string)
			psess.
				slicebytes(l, sval, len(sval))
			return true
		}

	case OSLICELIT:
		psess.
			initplan(r)

		bound := r.Right.Int64(psess)
		ta := psess.types.NewArray(r.Type.Elem(psess.types), bound)
		a := psess.staticname(ta)
		psess.
			inittemps[r] = a
		n := l.copy()
		n.Xoffset = l.Xoffset + int64(psess.array_array)
		psess.
			gdata(n, psess.nod(OADDR, a, nil), psess.Widthptr)
		n.Xoffset = l.Xoffset + int64(psess.array_nel)
		psess.
			gdata(n, r.Right, psess.Widthptr)
		n.Xoffset = l.Xoffset + int64(psess.array_cap)
		psess.
			gdata(n, r.Right, psess.Widthptr)

		l = a
		fallthrough

	case OARRAYLIT, OSTRUCTLIT:
		psess.
			initplan(r)

		p := psess.initplans[r]
		n := l.copy()
		for i := range p.E {
			e := &p.E[i]
			n.Xoffset = l.Xoffset + e.Xoffset
			n.Type = e.Expr.Type
			if e.Expr.Op == OLITERAL {
				psess.
					gdata(n, e.Expr, int(n.Type.Width))
				continue
			}
			psess.
				setlineno(e.Expr)
			a := n.copy()
			a.Orig = a
			if !psess.staticassign(a, e.Expr, out) {
				*out = append(*out, psess.nod(OAS, a, e.Expr))
			}
		}

		return true

	case OMAPLIT:
		break

	case OCLOSURE:
		if hasemptycvars(r) {
			if psess.Debug_closure > 0 {
				psess.
					Warnl(r.Pos, "closure converted to global")
			}
			psess.
				gdata(l, r.Func.Closure.Func.Nname, psess.Widthptr)
			return true
		}
		psess.
			closuredebugruntimecheck(r)

	case OCONVIFACE:

		val := r
		for val.Op == OCONVIFACE {
			val = val.Left
		}
		if val.Type.IsInterface() {

			return psess.Isconst(val, CTNIL)
		}

		var itab *Node
		if l.Type.IsEmptyInterface(psess.types) {
			itab = psess.typename(val.Type)
		} else {
			itab = psess.itabname(val.Type, l.Type)
		}

		n := l.copy()
		psess.
			gdata(n, itab, psess.Widthptr)
		n.Xoffset += int64(psess.Widthptr)

		if psess.isdirectiface(val.Type) {
			if psess.Isconst(val, CTNIL) {

				return true
			}

			n.Type = val.Type
			psess.
				setlineno(val)
			a := n.copy()
			a.Orig = a
			if !psess.staticassign(a, val, out) {
				*out = append(*out, psess.nod(OAS, a, val))
			}
		} else {

			a := psess.staticname(val.Type)
			psess.
				inittemps[val] = a
			if !psess.staticassign(a, val, out) {
				*out = append(*out, psess.nod(OAS, a, val))
			}
			ptr := psess.nod(OADDR, a, nil)
			n.Type = psess.types.NewPtr(val.Type)
			psess.
				gdata(n, ptr, psess.Widthptr)
		}

		return true
	}

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

// name generator for static temps

// staticname returns a name backed by a static data symbol.
// Callers should call n.Name.SetReadonly(true) on the
// returned node for readonly nodes.
func (psess *PackageSession) staticname(t *types.Type) *Node {

	n := psess.newname(psess.lookup(fmt.Sprintf("statictmp_%d", psess.statuniqgen)))
	psess.
		statuniqgen++
	psess.
		addvar(n, t, PEXTERN)
	return n
}

func (psess *PackageSession) isLiteral(n *Node) bool {

	return n.Op == OLITERAL && n.Val().Ctype(psess) != CTNIL
}

func (n *Node) isSimpleName() bool {
	return n.Op == ONAME && n.Addable() && n.Class() != PAUTOHEAP && n.Class() != PEXTERN
}

func (psess *PackageSession) litas(l *Node, r *Node, init *Nodes) {
	a := psess.nod(OAS, l, r)
	a = psess.typecheck(a, Etop)
	a = psess.walkexpr(a, init)
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
func (psess *PackageSession) getdyn(n *Node, top bool) initGenType {
	switch n.Op {
	default:
		if psess.isLiteral(n) {
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
		mode |= psess.getdyn(n1, false)
		if mode == initDynamic|initConst {
			break
		}
	}
	return mode
}

// isStaticCompositeLiteral reports whether n is a compile-time constant.
func (psess *PackageSession) isStaticCompositeLiteral(n *Node) bool {
	switch n.Op {
	case OSLICELIT:
		return false
	case OARRAYLIT:
		for _, r := range n.List.Slice() {
			if r.Op == OKEY {
				r = r.Right
			}
			if !psess.isStaticCompositeLiteral(r) {
				return false
			}
		}
		return true
	case OSTRUCTLIT:
		for _, r := range n.List.Slice() {
			if r.Op != OSTRUCTKEY {
				psess.
					Fatalf("isStaticCompositeLiteral: rhs not OSTRUCTKEY: %v", r)
			}
			if !psess.isStaticCompositeLiteral(r.Left) {
				return false
			}
		}
		return true
	case OLITERAL:
		return true
	case OCONVIFACE:

		val := n
		for val.Op == OCONVIFACE {
			val = val.Left
		}
		if val.Type.IsInterface() {
			return psess.Isconst(val, CTNIL)
		}
		if psess.isdirectiface(val.Type) && psess.Isconst(val, CTNIL) {
			return true
		}
		return psess.isStaticCompositeLiteral(val)
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
func (psess *PackageSession) fixedlit(ctxt initContext, kind initKind, n *Node, var_ *Node, init *Nodes) {
	var splitnode func(*Node) (a *Node, value *Node)
	switch n.Op {
	case OARRAYLIT, OSLICELIT:
		var k int64
		splitnode = func(r *Node) (*Node, *Node) {
			if r.Op == OKEY {
				k = psess.nonnegintconst(r.Left)
				r = r.Right
			}
			a := psess.nod(OINDEX, var_, psess.nodintconst(k))
			k++
			return a, r
		}
	case OSTRUCTLIT:
		splitnode = func(r *Node) (*Node, *Node) {
			if r.Op != OSTRUCTKEY {
				psess.
					Fatalf("fixedlit: rhs not OSTRUCTKEY: %v", r)
			}
			if r.Sym.IsBlank() {
				return psess.nblank, r.Left
			}
			return psess.nodSym(ODOT, var_, r.Sym), r.Left
		}
	default:
		psess.
			Fatalf("fixedlit bad op: %v", n.Op)
	}

	for _, r := range n.List.Slice() {
		a, value := splitnode(r)

		switch value.Op {
		case OSLICELIT:
			if (kind == initKindStatic && ctxt == inNonInitFunction) || (kind == initKindDynamic && ctxt == inInitFunction) {
				psess.
					slicelit(ctxt, value, a, init)
				continue
			}

		case OARRAYLIT, OSTRUCTLIT:
			psess.
				fixedlit(ctxt, kind, value, a, init)
			continue
		}

		islit := psess.isLiteral(value)
		if (kind == initKindStatic && !islit) || (kind == initKindDynamic && islit) {
			continue
		}
		psess.
			setlineno(value)
		a = psess.nod(OAS, a, value)
		a = psess.typecheck(a, Etop)
		switch kind {
		case initKindStatic:
			psess.
				genAsStatic(a)
		case initKindDynamic, initKindLocalCode:
			a = psess.orderStmtInPlace(a)
			a = psess.walkstmt(a)
			init.Append(a)
		default:
			psess.
				Fatalf("fixedlit: bad kind %d", kind)
		}

	}
}

func (psess *PackageSession) slicelit(ctxt initContext, n *Node, var_ *Node, init *Nodes) {

	t := psess.types.NewArray(n.Type.Elem(psess.types), n.Right.Int64(psess))
	psess.
		dowidth(t)

	if ctxt == inNonInitFunction {

		vstat := psess.staticname(t)
		psess.
			fixedlit(ctxt, initKindStatic, n, vstat, init)
		psess.
			fixedlit(ctxt, initKindDynamic, n, vstat, init)

		var_ = psess.typecheck(var_, Erv|Easgn)
		var nam Node
		if !psess.stataddr(&nam, var_) || nam.Class() != PEXTERN {
			psess.
				Fatalf("slicelit: %v", var_)
		}

		var v Node
		v.Type = psess.types.Types[TINT]
		psess.
			setintconst(&v, t.NumElem(psess.types))

		nam.Xoffset += int64(psess.array_array)
		psess.
			gdata(&nam, psess.nod(OADDR, vstat, nil), psess.Widthptr)
		nam.Xoffset += int64(psess.array_nel) - int64(psess.array_array)
		psess.
			gdata(&nam, &v, psess.Widthptr)
		nam.Xoffset += int64(psess.array_cap) - int64(psess.array_nel)
		psess.
			gdata(&nam, &v, psess.Widthptr)

		return
	}

	// if the literal contains constants,
	// make static initialized array (1),(2)
	var vstat *Node

	mode := psess.getdyn(n, true)
	if mode&initConst != 0 {
		vstat = psess.staticname(t)
		if ctxt == inInitFunction {
			vstat.Name.SetReadonly(true)
		}
		psess.
			fixedlit(ctxt, initKindStatic, n, vstat, init)
	}

	vauto := psess.temp(psess.types.NewPtr(t))

	// set auto to point at new temp or heap (3 assign)
	var a *Node
	if x := psess.prealloc[n]; x != nil {

		x.Type = t

		if vstat == nil {
			a = psess.nod(OAS, x, nil)
			a = psess.typecheck(a, Etop)
			init.Append(a)
		}

		a = psess.nod(OADDR, x, nil)
	} else if n.Esc == EscNone {
		a = psess.temp(t)
		if vstat == nil {
			a = psess.nod(OAS, psess.temp(t), nil)
			a = psess.typecheck(a, Etop)
			init.Append(a)
			a = a.Left
		}

		a = psess.nod(OADDR, a, nil)
	} else {
		a = psess.nod(ONEW, nil, nil)
		a.List.Set1(psess.typenod(t))
	}

	a = psess.nod(OAS, vauto, a)
	a = psess.typecheck(a, Etop)
	a = psess.walkexpr(a, init)
	init.Append(a)

	if vstat != nil {

		a = psess.nod(OIND, vauto, nil)

		a = psess.nod(OAS, a, vstat)
		a = psess.typecheck(a, Etop)
		a = psess.walkexpr(a, init)
		init.Append(a)
	}

	// put dynamics into array (5)
	var index int64
	for _, value := range n.List.Slice() {
		if value.Op == OKEY {
			index = psess.nonnegintconst(value.Left)
			value = value.Right
		}
		a := psess.nod(OINDEX, vauto, psess.nodintconst(index))
		a.SetBounded(true)
		index++

		switch value.Op {
		case OSLICELIT:
			break

		case OARRAYLIT, OSTRUCTLIT:
			psess.
				fixedlit(ctxt, initKindDynamic, value, a, init)
			continue
		}

		if psess.isLiteral(value) {
			continue
		}
		psess.
			setlineno(value)
		a = psess.nod(OAS, a, value)

		a = psess.typecheck(a, Etop)
		a = psess.orderStmtInPlace(a)
		a = psess.walkstmt(a)
		init.Append(a)
	}

	a = psess.nod(OAS, var_, psess.nod(OSLICE, vauto, nil))

	a = psess.typecheck(a, Etop)
	a = psess.orderStmtInPlace(a)
	a = psess.walkstmt(a)
	init.Append(a)
}

func (psess *PackageSession) maplit(n *Node, m *Node, init *Nodes) {

	a := psess.nod(OMAKE, nil, nil)
	a.Esc = n.Esc
	a.List.Set2(psess.typenod(n.Type), psess.nodintconst(int64(n.List.Len())))
	psess.
		litas(m, a, init)

	// Split the initializers into static and dynamic.
	var stat, dyn []*Node
	for _, r := range n.List.Slice() {
		if r.Op != OKEY {
			psess.
				Fatalf("maplit: rhs not OKEY: %v", r)
		}
		if psess.isStaticCompositeLiteral(r.Left) && psess.isStaticCompositeLiteral(r.Right) {
			stat = append(stat, r)
		} else {
			dyn = append(dyn, r)
		}
	}

	if len(stat) > 25 {

		tk := psess.types.NewArray(n.Type.Key(psess.types), int64(len(stat)))
		tv := psess.types.NewArray(n.Type.Elem(psess.types), int64(len(stat)))
		psess.
			dowidth(tk)
		psess.
			dowidth(tv)

		vstatk := psess.staticname(tk)
		vstatk.Name.SetReadonly(true)
		vstatv := psess.staticname(tv)
		vstatv.Name.SetReadonly(true)

		datak := psess.nod(OARRAYLIT, nil, nil)
		datav := psess.nod(OARRAYLIT, nil, nil)
		for _, r := range stat {
			datak.List.Append(r.Left)
			datav.List.Append(r.Right)
		}
		psess.
			fixedlit(inInitFunction, initKindStatic, datak, vstatk, init)
		psess.
			fixedlit(inInitFunction, initKindStatic, datav, vstatv, init)

		i := psess.temp(psess.types.Types[TINT])
		rhs := psess.nod(OINDEX, vstatv, i)
		rhs.SetBounded(true)

		kidx := psess.nod(OINDEX, vstatk, i)
		kidx.SetBounded(true)
		lhs := psess.nod(OINDEX, m, kidx)

		zero := psess.nod(OAS, i, psess.nodintconst(0))
		cond := psess.nod(OLT, i, psess.nodintconst(tk.NumElem(psess.types)))
		incr := psess.nod(OAS, i, psess.nod(OADD, i, psess.nodintconst(1)))
		body := psess.nod(OAS, lhs, rhs)

		loop := psess.nod(OFOR, cond, incr)
		loop.Nbody.Set1(body)
		loop.Ninit.Set1(zero)

		loop = psess.typecheck(loop, Etop)
		loop = psess.walkstmt(loop)
		init.Append(loop)
	} else {
		psess.
			addMapEntries(m, stat, init)
	}
	psess.
		addMapEntries(m, dyn, init)
}

func (psess *PackageSession) addMapEntries(m *Node, dyn []*Node, init *Nodes) {
	if len(dyn) == 0 {
		return
	}

	nerr := psess.nerrors

	key := psess.temp(m.Type.Key(psess.types))
	val := psess.temp(m.Type.Elem(psess.types))

	for _, r := range dyn {
		index, value := r.Left, r.Right
		psess.
			setlineno(index)
		a := psess.nod(OAS, key, index)
		a = psess.typecheck(a, Etop)
		a = psess.walkstmt(a)
		init.Append(a)
		psess.
			setlineno(value)
		a = psess.nod(OAS, val, value)
		a = psess.typecheck(a, Etop)
		a = psess.walkstmt(a)
		init.Append(a)
		psess.
			setlineno(val)
		a = psess.nod(OAS, psess.nod(OINDEX, m, key), val)
		a = psess.typecheck(a, Etop)
		a = psess.walkstmt(a)
		init.Append(a)

		if nerr != psess.nerrors {
			break
		}
	}

	a := psess.nod(OVARKILL, key, nil)
	a = psess.typecheck(a, Etop)
	init.Append(a)
	a = psess.nod(OVARKILL, val, nil)
	a = psess.typecheck(a, Etop)
	init.Append(a)
}

func (psess *PackageSession) anylit(n *Node, var_ *Node, init *Nodes) {
	t := n.Type
	switch n.Op {
	default:
		psess.
			Fatalf("anylit: not lit, op=%v node=%v", n.Op, n)

	case OPTRLIT:
		if !t.IsPtr() {
			psess.
				Fatalf("anylit: not ptr")
		}

		var r *Node
		if n.Right != nil {

			init.Append(psess.nod(OAS, n.Right, nil))
			r = psess.nod(OADDR, n.Right, nil)
			r = psess.typecheck(r, Erv)
		} else {
			r = psess.nod(ONEW, nil, nil)
			r.SetTypecheck(1)
			r.Type = t
			r.Esc = n.Esc
		}

		r = psess.walkexpr(r, init)
		a := psess.nod(OAS, var_, r)

		a = psess.typecheck(a, Etop)
		init.Append(a)

		var_ = psess.nod(OIND, var_, nil)
		var_ = psess.typecheck(var_, Erv|Easgn)
		psess.
			anylit(n.Left, var_, init)

	case OSTRUCTLIT, OARRAYLIT:
		if !t.IsStruct() && !t.IsArray() {
			psess.
				Fatalf("anylit: not struct/array")
		}

		if var_.isSimpleName() && n.List.Len() > 4 {

			vstat := psess.staticname(t)
			vstat.Name.SetReadonly(true)

			ctxt := inInitFunction
			if n.Op == OARRAYLIT {
				ctxt = inNonInitFunction
			}
			psess.
				fixedlit(ctxt, initKindStatic, n, vstat, init)

			a := psess.nod(OAS, var_, vstat)

			a = psess.typecheck(a, Etop)
			a = psess.walkexpr(a, init)
			init.Append(a)
			psess.
				fixedlit(inInitFunction, initKindDynamic, n, var_, init)
			break
		}

		var components int64
		if n.Op == OARRAYLIT {
			components = t.NumElem(psess.types)
		} else {
			components = int64(t.NumFields(psess.types))
		}

		if var_.isSimpleName() || int64(n.List.Len()) < components {
			a := psess.nod(OAS, var_, nil)
			a = psess.typecheck(a, Etop)
			a = psess.walkexpr(a, init)
			init.Append(a)
		}
		psess.
			fixedlit(inInitFunction, initKindLocalCode, n, var_, init)

	case OSLICELIT:
		psess.
			slicelit(inInitFunction, n, var_, init)

	case OMAPLIT:
		if !t.IsMap() {
			psess.
				Fatalf("anylit: not map")
		}
		psess.
			maplit(n, var_, init)
	}
}

func (psess *PackageSession) oaslit(n *Node, init *Nodes) bool {
	if n.Left == nil || n.Right == nil {

		return false
	}
	if n.Left.Type == nil || n.Right.Type == nil {

		return false
	}
	if !n.Left.isSimpleName() {

		return false
	}
	if !psess.eqtype(n.Left.Type, n.Right.Type) {

		return false
	}

	switch n.Right.Op {
	default:

		return false

	case OSTRUCTLIT, OARRAYLIT, OSLICELIT, OMAPLIT:
		if vmatch1(n.Left, n.Right) {

			return false
		}
		psess.
			anylit(n.Right, n.Left, init)
	}

	n.Op = OEMPTY
	n.Right = nil
	return true
}

func (psess *PackageSession) getlit(lit *Node) int {
	if psess.smallintconst(lit) {
		return int(lit.Int64(psess))
	}
	return -1
}

// stataddr sets nam to the static address of n and reports whether it succeeded.
func (psess *PackageSession) stataddr(nam *Node, n *Node) bool {
	if n == nil {
		return false
	}

	switch n.Op {
	case ONAME:
		*nam = *n
		return n.Addable()

	case ODOT:
		if !psess.stataddr(nam, n.Left) {
			break
		}
		nam.Xoffset += n.Xoffset
		nam.Type = n.Type
		return true

	case OINDEX:
		if n.Left.Type.IsSlice() {
			break
		}
		if !psess.stataddr(nam, n.Left) {
			break
		}
		l := psess.getlit(n.Right)
		if l < 0 {
			break
		}

		if n.Type.Width != 0 && psess.thearch.MAXWIDTH/n.Type.Width <= int64(l) {
			break
		}
		nam.Xoffset += int64(l) * n.Type.Width
		nam.Type = n.Type
		return true
	}

	return false
}

func (psess *PackageSession) initplan(n *Node) {
	if psess.initplans[n] != nil {
		return
	}
	p := new(InitPlan)
	psess.
		initplans[n] = p
	switch n.Op {
	default:
		psess.
			Fatalf("initplan")

	case OARRAYLIT, OSLICELIT:
		var k int64
		for _, a := range n.List.Slice() {
			if a.Op == OKEY {
				k = psess.nonnegintconst(a.Left)
				a = a.Right
			}
			psess.
				addvalue(p, k*n.Type.Elem(psess.types).Width, a)
			k++
		}

	case OSTRUCTLIT:
		for _, a := range n.List.Slice() {
			if a.Op != OSTRUCTKEY {
				psess.
					Fatalf("initplan fixedlit")
			}
			psess.
				addvalue(p, a.Xoffset, a.Left)
		}

	case OMAPLIT:
		for _, a := range n.List.Slice() {
			if a.Op != OKEY {
				psess.
					Fatalf("initplan maplit")
			}
			psess.
				addvalue(p, -1, a.Right)
		}
	}
}

func (psess *PackageSession) addvalue(p *InitPlan, xoffset int64, n *Node) {

	if psess.isZero(n) {
		return
	}

	if isvaluelit(n) {
		psess.
			initplan(n)
		q := psess.initplans[n]
		for _, qe := range q.E {

			qe.Xoffset += xoffset
			p.E = append(p.E, qe)
		}
		return
	}

	p.E = append(p.E, InitEntry{Xoffset: xoffset, Expr: n})
}

func (psess *PackageSession) isZero(n *Node) bool {
	switch n.Op {
	case OLITERAL:
		switch u := n.Val().U.(type) {
		default:
			Dump("unexpected literal", n)
			psess.
				Fatalf("isZero")
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
			if !psess.isZero(n1) {
				return false
			}
		}
		return true

	case OSTRUCTLIT:
		for _, n1 := range n.List.Slice() {
			if !psess.isZero(n1.Left) {
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

func (psess *PackageSession) genAsStatic(as *Node) {
	if as.Left.Type == nil {
		psess.
			Fatalf("genAsStatic as.Left not typechecked")
	}

	var nam Node
	if !psess.stataddr(&nam, as.Left) || (nam.Class() != PEXTERN && as.Left != psess.nblank) {
		psess.
			Fatalf("genAsStatic: lhs %v", as.Left)
	}

	switch {
	case as.Right.Op == OLITERAL:
	case as.Right.Op == ONAME && as.Right.Class() == PFUNC:
	default:
		psess.
			Fatalf("genAsStatic: rhs %v", as.Right)
	}
	psess.
		gdata(&nam, as.Right, int(as.Right.Type.Width))
}

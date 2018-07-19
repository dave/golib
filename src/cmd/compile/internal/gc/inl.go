package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/src"
	"strings"
)

// Inlining budget parameters, gathered in one place
const (
	inlineMaxBudget       = 80
	inlineExtraAppendCost = 0
	inlineExtraCallCost   = inlineMaxBudget // default is do not inline, -l=4 enables by using 1 instead.
	inlineExtraPanicCost  = 1               // do not penalize inlining panics.
	inlineExtraThrowCost  = inlineMaxBudget // with current (2018-05/1.11) code, inlining runtime.throw does not help.
)

// Get the function's package. For ordinary functions it's on the ->sym, but for imported methods
// the ->sym can be re-used in the local package, so peel it off the receiver's type.
func (psess *PackageSession) fnpkg(fn *Node) *types.Pkg {
	if fn.IsMethod(psess) {

		rcvr := fn.Type.Recv(psess.types).Type

		if rcvr.IsPtr() {
			rcvr = rcvr.Elem(psess.types)
		}
		if rcvr.Sym == nil {
			psess.
				Fatalf("receiver with no sym: [%v] %L  (%v)", fn.Sym, fn, rcvr)
		}
		return rcvr.Sym.Pkg
	}

	return fn.Sym.Pkg
}

// Lazy typechecking of imported bodies. For local functions, caninl will set ->typecheck
// because they're a copy of an already checked body.
func (psess *PackageSession) typecheckinl(fn *Node) {
	lno := psess.setlineno(fn)

	if psess.flagiexport {
		psess.
			expandInline(fn)
	}

	pkg := psess.fnpkg(fn)

	if pkg == psess.localpkg || pkg == nil {
		return
	}

	if psess.Debug['m'] > 2 || psess.Debug_export != 0 {
		fmt.Printf("typecheck import [%v] %L { %#v }\n", fn.Sym, fn, asNodes(fn.Func.Inl.Body))
	}

	save_safemode := psess.safemode
	psess.
		safemode = false

	savefn := psess.Curfn
	psess.
		Curfn = fn
	psess.
		typecheckslice(fn.Func.Inl.Body, Etop)
	psess.
		Curfn = savefn

	fn.Func.Inl.Dcl = append(fn.Func.Inl.Dcl, fn.Func.Dcl...)
	fn.Func.Dcl = nil
	psess.
		safemode = save_safemode
	psess.
		lineno = lno
}

// Caninl determines whether fn is inlineable.
// If so, caninl saves fn->nbody in fn->inl and substitutes it with a copy.
// fn and ->nbody will already have been typechecked.
func (psess *PackageSession) caninl(fn *Node) {
	if fn.Op != ODCLFUNC {
		psess.
			Fatalf("caninl %v", fn)
	}
	if fn.Func.Nname == nil {
		psess.
			Fatalf("caninl no nname %+v", fn)
	}

	var reason string // reason, if any, that the function was not inlined
	if psess.Debug['m'] > 1 {
		defer func() {
			if reason != "" {
				fmt.Printf("%v: cannot inline %v: %s\n", fn.Line(psess), fn.Func.Nname, reason)
			}
		}()
	}

	if fn.Func.Pragma&Noinline != 0 {
		reason = "marked go:noinline"
		return
	}

	if psess.flag_race && fn.Func.Pragma&Norace != 0 {
		reason = "marked go:norace with -race compilation"
		return
	}

	if fn.Func.Pragma&CgoUnsafeArgs != 0 {
		reason = "marked go:cgo_unsafe_args"
		return
	}

	if fn.Func.Pragma&Yeswritebarrierrec != 0 {
		reason = "marked go:yeswritebarrierrec"
		return
	}

	if fn.Nbody.Len() == 0 {
		reason = "no function body"
		return
	}

	if fn.Typecheck() == 0 {
		psess.
			Fatalf("caninl on non-typechecked function %v", fn)
	}

	n := fn.Func.Nname
	if n.Func.InlinabilityChecked() {
		return
	}
	defer n.Func.SetInlinabilityChecked(true)

	cc := int32(inlineExtraCallCost)
	if psess.Debug['l'] == 4 {
		cc = 1
	}

	visitor := hairyVisitor{
		budget:        inlineMaxBudget,
		extraCallCost: cc,
		usedLocals:    make(map[*Node]bool),
	}
	if visitor.visitList(psess, fn.Nbody) {
		reason = visitor.reason
		return
	}
	if visitor.budget < 0 {
		reason = fmt.Sprintf("function too complex: cost %d exceeds budget %d", inlineMaxBudget-visitor.budget, inlineMaxBudget)
		return
	}

	n.Func.Inl = &Inline{
		Cost: inlineMaxBudget - visitor.budget,
		Dcl:  psess.inlcopylist(pruneUnusedAutos(n.Name.Defn.Func.Dcl, &visitor)),
		Body: psess.inlcopylist(fn.Nbody.Slice()),
	}

	fn.Type.FuncType(psess.types).Nname = asTypesNode(n)

	if psess.Debug['m'] > 1 {
		fmt.Printf("%v: can inline %#v as: %#v { %#v }\n", fn.Line(psess), n, fn.Type, asNodes(n.Func.Inl.Body))
	} else if psess.Debug['m'] != 0 {
		fmt.Printf("%v: can inline %v\n", fn.Line(psess), n)
	}
}

// inlFlood marks n's inline body for export and recursively ensures
// all called functions are marked too.
func (psess *PackageSession) inlFlood(n *Node) {
	if n == nil {
		return
	}
	if n.Op != ONAME || n.Class() != PFUNC {
		psess.
			Fatalf("inlFlood: unexpected %v, %v, %v", n, n.Op, n.Class())
	}
	if n.Func == nil {
		psess.
			Fatalf("inlFlood: missing Func on %v", n)
	}
	if n.Func.Inl == nil {
		return
	}

	if n.Func.ExportInline() {
		return
	}
	n.Func.SetExportInline(true)
	psess.
		typecheckinl(n)

	inspectList(asNodes(n.Func.Inl.Body), func(n *Node) bool {
		switch n.Op {
		case ONAME:

			if n.Class() == PEXTERN || n.Class() == PFUNC && !n.isMethodExpression() {
				psess.
					exportsym(n)
			}

		case OCALLFUNC, OCALLMETH:
			psess.
				inlFlood(asNode(n.Left.Type.Nname(psess.types)))
		}
		return true
	})
}

// hairyVisitor visits a function body to determine its inlining
// hairiness and whether or not it can be inlined.
type hairyVisitor struct {
	budget        int32
	reason        string
	extraCallCost int32
	usedLocals    map[*Node]bool
}

// Look for anything we want to punt on.
func (v *hairyVisitor) visitList(psess *PackageSession, ll Nodes) bool {
	for _, n := range ll.Slice() {
		if v.visit(psess, n) {
			return true
		}
	}
	return false
}

func (v *hairyVisitor) visit(psess *PackageSession, n *Node) bool {
	if n == nil {
		return false
	}

	switch n.Op {

	case OCALLFUNC:
		if psess.isIntrinsicCall(n) {
			v.budget--
			break
		}

		if n.Left.Op == ONAME && n.Left.Class() == PFUNC && psess.isRuntimePkg(n.Left.Sym.Pkg) {
			fn := n.Left.Sym.Name
			if fn == "getcallerpc" || fn == "getcallersp" {
				v.reason = "call to " + fn
				return true
			}
			if fn == "throw" {
				v.budget -= inlineExtraThrowCost
				break
			}
		}

		if fn := n.Left.Func; fn != nil && fn.Inl != nil {
			v.budget -= fn.Inl.Cost
			break
		}
		if n.Left.isMethodExpression() {
			if d := asNode(n.Left.Sym.Def); d != nil && d.Func.Inl != nil {
				v.budget -= d.Func.Inl.Cost
				break
			}
		}

		v.budget -= v.extraCallCost

	case OCALLMETH:
		t := n.Left.Type
		if t == nil {
			psess.
				Fatalf("no function type for [%p] %+v\n", n.Left, n.Left)
		}
		if t.Nname(psess.types) == nil {
			psess.
				Fatalf("no function definition for [%p] %+v\n", t, t)
		}
		if psess.isRuntimePkg(n.Left.Sym.Pkg) {
			fn := n.Left.Sym.Name
			if fn == "heapBits.nextArena" {

				break
			}
		}
		if inlfn := asNode(t.FuncType(psess.types).Nname).Func; inlfn.Inl != nil {
			v.budget -= inlfn.Inl.Cost
			break
		}

		v.budget -= v.extraCallCost

	case OCALL, OCALLINTER:

		v.budget -= v.extraCallCost

	case OPANIC:
		v.budget -= inlineExtraPanicCost

	case ORECOVER:

		v.reason = "call to recover"
		return true

	case OCLOSURE,
		OCALLPART,
		ORANGE,
		OFOR,
		OFORUNTIL,
		OSELECT,
		OTYPESW,
		OPROC,
		ODEFER,
		ODCLTYPE,
		OBREAK,
		ORETJMP:
		v.reason = "unhandled op " + n.Op.String(psess)
		return true

	case OAPPEND:
		v.budget -= inlineExtraAppendCost

	case ODCLCONST, OEMPTY, OFALL, OLABEL:

		return false

	case OIF:
		if psess.Isconst(n.Left, CTBOOL) {

			return v.visitList(psess, n.Ninit) || v.visitList(psess, n.Nbody) ||
				v.visitList(psess, n.Rlist)
		}

	case ONAME:
		if n.Class() == PAUTO {
			v.usedLocals[n] = true
		}

	}

	v.budget--

	switch n.Op {
	case OSTRUCTKEY:
		v.budget--
	case OSLICE, OSLICEARR, OSLICESTR:
		v.budget--
	case OSLICE3, OSLICE3ARR:
		v.budget -= 2
	}

	if v.budget < 0 && psess.Debug['m'] < 2 {
		return true
	}

	return v.visit(psess, n.Left) || v.visit(psess, n.Right) ||
		v.visitList(psess, n.List) || v.visitList(psess, n.Rlist) ||
		v.visitList(psess, n.Ninit) || v.visitList(psess, n.Nbody)
}

// Inlcopy and inlcopylist recursively copy the body of a function.
// Any name-like node of non-local class is marked for re-export by adding it to
// the exportlist.
func (psess *PackageSession) inlcopylist(ll []*Node) []*Node {
	s := make([]*Node, 0, len(ll))
	for _, n := range ll {
		s = append(s, psess.inlcopy(n))
	}
	return s
}

func (psess *PackageSession) inlcopy(n *Node) *Node {
	if n == nil {
		return nil
	}

	switch n.Op {
	case ONAME, OTYPE, OLITERAL:
		return n
	}

	m := n.copy()
	if m.Func != nil {
		psess.
			Fatalf("unexpected Func: %v", m)
	}
	m.Left = psess.inlcopy(n.Left)
	m.Right = psess.inlcopy(n.Right)
	m.List.Set(psess.inlcopylist(n.List.Slice()))
	m.Rlist.Set(psess.inlcopylist(n.Rlist.Slice()))
	m.Ninit.Set(psess.inlcopylist(n.Ninit.Slice()))
	m.Nbody.Set(psess.inlcopylist(n.Nbody.Slice()))

	return m
}

// Inlcalls/nodelist/node walks fn's statements and expressions and substitutes any
// calls made to inlineable functions. This is the external entry point.
func (psess *PackageSession) inlcalls(fn *Node) {
	savefn := psess.Curfn
	psess.
		Curfn = fn
	fn = psess.inlnode(fn)
	if fn != psess.Curfn {
		psess.
			Fatalf("inlnode replaced curfn")
	}
	psess.
		Curfn = savefn
}

// Turn an OINLCALL into a statement.
func inlconv2stmt(n *Node) {
	n.Op = OBLOCK

	n.List.Set(n.Nbody.Slice())

	n.Nbody.Set(nil)
	n.Rlist.Set(nil)
}

// Turn an OINLCALL into a single valued expression.
// The result of inlconv2expr MUST be assigned back to n, e.g.
// 	n.Left = inlconv2expr(n.Left)
func (psess *PackageSession) inlconv2expr(n *Node) *Node {
	r := n.Rlist.First()
	return psess.addinit(r, append(n.Ninit.Slice(), n.Nbody.Slice()...))
}

// Turn the rlist (with the return values) of the OINLCALL in
// n into an expression list lumping the ninit and body
// containing the inlined statements on the first list element so
// order will be preserved Used in return, oas2func and call
// statements.
func (psess *PackageSession) inlconv2list(n *Node) []*Node {
	if n.Op != OINLCALL || n.Rlist.Len() == 0 {
		psess.
			Fatalf("inlconv2list %+v\n", n)
	}

	s := n.Rlist.Slice()
	s[0] = psess.addinit(s[0], append(n.Ninit.Slice(), n.Nbody.Slice()...))
	return s
}

func (psess *PackageSession) inlnodelist(l Nodes) {
	s := l.Slice()
	for i := range s {
		s[i] = psess.inlnode(s[i])
	}
}

// inlnode recurses over the tree to find inlineable calls, which will
// be turned into OINLCALLs by mkinlcall. When the recursion comes
// back up will examine left, right, list, rlist, ninit, ntest, nincr,
// nbody and nelse and use one of the 4 inlconv/glue functions above
// to turn the OINLCALL into an expression, a statement, or patch it
// in to this nodes list or rlist as appropriate.
// NOTE it makes no sense to pass the glue functions down the
// recursion to the level where the OINLCALL gets created because they
// have to edit /this/ n, so you'd have to push that one down as well,
// but then you may as well do it here.  so this is cleaner and
// shorter and less complicated.
// The result of inlnode MUST be assigned back to n, e.g.
// 	n.Left = inlnode(n.Left)
func (psess *PackageSession) inlnode(n *Node) *Node {
	if n == nil {
		return n
	}

	switch n.Op {

	case ODEFER, OPROC:
		switch n.Left.Op {
		case OCALLFUNC, OCALLMETH:
			n.Left.SetNoInline(true)
		}
		return n

	case OCLOSURE:
		return n
	}

	lno := psess.setlineno(n)
	psess.
		inlnodelist(n.Ninit)
	for _, n1 := range n.Ninit.Slice() {
		if n1.Op == OINLCALL {
			inlconv2stmt(n1)
		}
	}

	n.Left = psess.inlnode(n.Left)
	if n.Left != nil && n.Left.Op == OINLCALL {
		n.Left = psess.inlconv2expr(n.Left)
	}

	n.Right = psess.inlnode(n.Right)
	if n.Right != nil && n.Right.Op == OINLCALL {
		if n.Op == OFOR || n.Op == OFORUNTIL {
			inlconv2stmt(n.Right)
		} else {
			n.Right = psess.inlconv2expr(n.Right)
		}
	}
	psess.
		inlnodelist(n.List)
	switch n.Op {
	case OBLOCK:
		for _, n2 := range n.List.Slice() {
			if n2.Op == OINLCALL {
				inlconv2stmt(n2)
			}
		}

	case ORETURN, OCALLFUNC, OCALLMETH, OCALLINTER, OAPPEND, OCOMPLEX:

		if n.List.Len() == 1 && n.List.First().Op == OINLCALL && n.List.First().Rlist.Len() > 1 {
			n.List.Set(psess.inlconv2list(n.List.First()))
			break
		}
		fallthrough

	default:
		s := n.List.Slice()
		for i1, n1 := range s {
			if n1 != nil && n1.Op == OINLCALL {
				s[i1] = psess.inlconv2expr(s[i1])
			}
		}
	}
	psess.
		inlnodelist(n.Rlist)
	if n.Op == OAS2FUNC && n.Rlist.First().Op == OINLCALL {
		n.Rlist.Set(psess.inlconv2list(n.Rlist.First()))
		n.Op = OAS2
		n.SetTypecheck(0)
		n = psess.typecheck(n, Etop)
	} else {
		s := n.Rlist.Slice()
		for i1, n1 := range s {
			if n1.Op == OINLCALL {
				if n.Op == OIF {
					inlconv2stmt(n1)
				} else {
					s[i1] = psess.inlconv2expr(s[i1])
				}
			}
		}
	}
	psess.
		inlnodelist(n.Nbody)
	for _, n := range n.Nbody.Slice() {
		if n.Op == OINLCALL {
			inlconv2stmt(n)
		}
	}

	switch n.Op {
	case OCALLFUNC, OCALLMETH:
		if n.NoInline() {
			return n
		}
	}

	switch n.Op {
	case OCALLFUNC:
		if psess.Debug['m'] > 3 {
			fmt.Printf("%v:call to func %+v\n", n.Line(psess), n.Left)
		}
		if n.Left.Func != nil && n.Left.Func.Inl != nil && !psess.isIntrinsicCall(n) {
			n = psess.mkinlcall(n, n.Left)
		} else if n.Left.isMethodExpression() && asNode(n.Left.Sym.Def) != nil {
			n = psess.mkinlcall(n, asNode(n.Left.Sym.Def))
		} else if n.Left.Op == OCLOSURE {
			if f := psess.inlinableClosure(n.Left); f != nil {
				n = psess.mkinlcall(n, f)
			}
		} else if n.Left.Op == ONAME && n.Left.Name != nil && n.Left.Name.Defn != nil {
			if d := n.Left.Name.Defn; d.Op == OAS && d.Right.Op == OCLOSURE {
				if f := psess.inlinableClosure(d.Right); f != nil {

					if d.Left.Addrtaken() {
						if psess.Debug['m'] > 1 {
							fmt.Printf("%v: cannot inline escaping closure variable %v\n", n.Line(psess), n.Left)
						}
						break
					}

					if unsafe, a := psess.reassigned(n.Left); unsafe {
						if psess.Debug['m'] > 1 {
							if a != nil {
								fmt.Printf("%v: cannot inline re-assigned closure variable at %v: %v\n", n.Line(psess), a.Line(psess), a)
							} else {
								fmt.Printf("%v: cannot inline global closure variable %v\n", n.Line(psess), n.Left)
							}
						}
						break
					}
					n = psess.mkinlcall(n, f)
				}
			}
		}

	case OCALLMETH:
		if psess.Debug['m'] > 3 {
			fmt.Printf("%v:call to meth %L\n", n.Line(psess), n.Left.Right)
		}

		if n.Left.Type == nil {
			psess.
				Fatalf("no function type for [%p] %+v\n", n.Left, n.Left)
		}

		if n.Left.Type.Nname(psess.types) == nil {
			psess.
				Fatalf("no function definition for [%p] %+v\n", n.Left.Type, n.Left.Type)
		}

		n = psess.mkinlcall(n, asNode(n.Left.Type.FuncType(psess.types).Nname))
	}
	psess.
		lineno = lno
	return n
}

// inlinableClosure takes an OCLOSURE node and follows linkage to the matching ONAME with
// the inlinable body. Returns nil if the function is not inlinable.
func (psess *PackageSession) inlinableClosure(n *Node) *Node {
	c := n.Func.Closure
	psess.
		caninl(c)
	f := c.Func.Nname
	if f == nil || f.Func.Inl == nil {
		return nil
	}
	return f
}

// reassigned takes an ONAME node, walks the function in which it is defined, and returns a boolean
// indicating whether the name has any assignments other than its declaration.
// The second return value is the first such assignment encountered in the walk, if any. It is mostly
// useful for -m output documenting the reason for inhibited optimizations.
// NB: global variables are always considered to be re-assigned.
// TODO: handle initial declaration not including an assignment and followed by a single assignment?
func (psess *PackageSession) reassigned(n *Node) (bool, *Node) {
	if n.Op != ONAME {
		psess.
			Fatalf("reassigned %v", n)
	}

	if n.Name.Curfn == nil {
		return true, nil
	}
	f := n.Name.Curfn

	if f.Op == OCLOSURE {
		f = f.Func.Closure
	}
	v := reassignVisitor{name: n}
	a := v.visitList(f.Nbody)
	return a != nil, a
}

type reassignVisitor struct {
	name *Node
}

func (v *reassignVisitor) visit(n *Node) *Node {
	if n == nil {
		return nil
	}
	switch n.Op {
	case OAS:
		if n.Left == v.name && n != v.name.Name.Defn {
			return n
		}
		return nil
	case OAS2, OAS2FUNC, OAS2MAPR, OAS2DOTTYPE:
		for _, p := range n.List.Slice() {
			if p == v.name && n != v.name.Name.Defn {
				return n
			}
		}
		return nil
	}
	if a := v.visit(n.Left); a != nil {
		return a
	}
	if a := v.visit(n.Right); a != nil {
		return a
	}
	if a := v.visitList(n.List); a != nil {
		return a
	}
	if a := v.visitList(n.Rlist); a != nil {
		return a
	}
	if a := v.visitList(n.Ninit); a != nil {
		return a
	}
	if a := v.visitList(n.Nbody); a != nil {
		return a
	}
	return nil
}

func (v *reassignVisitor) visitList(l Nodes) *Node {
	for _, n := range l.Slice() {
		if a := v.visit(n); a != nil {
			return a
		}
	}
	return nil
}

// The result of mkinlcall MUST be assigned back to n, e.g.
// 	n.Left = mkinlcall(n.Left, fn, isddd)
func (psess *PackageSession) mkinlcall(n *Node, fn *Node) *Node {
	save_safemode := psess.safemode

	pkg := psess.fnpkg(fn)

	if pkg != psess.localpkg && pkg != nil {
		psess.
			safemode = false
	}
	n = psess.mkinlcall1(n, fn)
	psess.
		safemode = save_safemode
	return n
}

func (psess *PackageSession) tinlvar(t *types.Field, inlvars map[*Node]*Node) *Node {
	if n := asNode(t.Nname); n != nil && !n.isBlank() {
		inlvar := inlvars[n]
		if inlvar == nil {
			psess.
				Fatalf("missing inlvar for %v\n", n)
		}
		return inlvar
	}

	return psess.typecheck(psess.nblank, Erv|Easgn)
}

// If n is a call, and fn is a function with an inlinable body,
// return an OINLCALL.
// On return ninit has the parameter assignments, the nbody is the
// inlined function body and list, rlist contain the input, output
// parameters.
// The result of mkinlcall1 MUST be assigned back to n, e.g.
// 	n.Left = mkinlcall1(n.Left, fn, isddd)
func (psess *PackageSession) mkinlcall1(n, fn *Node) *Node {
	if fn.Func.Inl == nil {

		return n
	}

	if fn == psess.Curfn || fn.Name.Defn == psess.Curfn {

		return n
	}

	if psess.instrumenting && psess.isRuntimePkg(fn.Sym.Pkg) {

		return n
	}

	if psess.Debug_typecheckinl == 0 {
		psess.
			typecheckinl(fn)
	}

	if psess.Debug['m'] > 1 {
		fmt.Printf("%v: inlining call to %v %#v { %#v }\n", n.Line(psess), fn.Sym, fn.Type, asNodes(fn.Func.Inl.Body))
	} else if psess.Debug['m'] != 0 {
		fmt.Printf("%v: inlining call to %v\n", n.Line(psess), fn)
	}
	if psess.Debug['m'] > 2 {
		fmt.Printf("%v: Before inlining: %+v\n", n.Line(psess), n)
	}

	ninit := n.Ninit

	inlvars := make(map[*Node]*Node)

	// record formals/locals for later post-processing
	var inlfvars []*Node

	if fn.Name.Defn != nil {
		if c := fn.Name.Defn.Func.Closure; c != nil {
			for _, v := range c.Func.Closure.Func.Cvars.Slice() {
				if v.Op == OXXX {
					continue
				}

				o := v.Name.Param.Outer

				if o == nil || (o.Name.Curfn != psess.Curfn && o.Name.Curfn.Func.Closure != psess.Curfn) {
					psess.
						Fatalf("%v: unresolvable capture %v %v\n", n.Line(psess), fn, v)
				}

				if v.Name.Byval() {
					iv := psess.typecheck(psess.inlvar(v), Erv)
					ninit.Append(psess.nod(ODCL, iv, nil))
					ninit.Append(psess.typecheck(psess.nod(OAS, iv, o), Etop))
					inlvars[v] = iv
				} else {
					addr := psess.newname(psess.lookup("&" + v.Sym.Name))
					addr.Type = psess.types.NewPtr(v.Type)
					ia := psess.typecheck(psess.inlvar(addr), Erv)
					ninit.Append(psess.nod(ODCL, ia, nil))
					ninit.Append(psess.typecheck(psess.nod(OAS, ia, psess.nod(OADDR, o, nil)), Etop))
					inlvars[addr] = ia

					inlvars[v] = psess.typecheck(psess.nod(OIND, ia, nil), Erv)
				}
			}
		}
	}

	for _, ln := range fn.Func.Inl.Dcl {
		if ln.Op != ONAME {
			continue
		}
		if ln.Class() == PPARAMOUT {
			continue
		}
		if ln.isParamStackCopy() {
			continue
		}
		inlvars[ln] = psess.typecheck(psess.inlvar(ln), Erv)
		if ln.Class() == PPARAM || ln.Name.Param.Stackcopy != nil && ln.Name.Param.Stackcopy.Class() == PPARAM {
			ninit.Append(psess.nod(ODCL, inlvars[ln], nil))
		}
		if psess.genDwarfInline > 0 {
			inlf := inlvars[ln]
			if ln.Class() == PPARAM {
				inlf.SetInlFormal(true)
			} else {
				inlf.SetInlLocal(true)
			}
			inlf.Pos = ln.Pos
			inlfvars = append(inlfvars, inlf)
		}
	}

	// temporaries for return values.
	var retvars []*Node
	for i, t := range fn.Type.Results(psess.types).Fields(psess.types).Slice() {
		var m *Node
		mpos := t.Pos
		if n := asNode(t.Nname); n != nil && !n.isBlank() {
			m = psess.inlvar(n)
			m = psess.typecheck(m, Erv)
			inlvars[n] = m
		} else {

			m = psess.retvar(t, i)
		}

		if psess.genDwarfInline > 0 {

			if !strings.HasPrefix(m.Sym.Name, "~R") {
				m.SetInlFormal(true)
				m.Pos = mpos
				inlfvars = append(inlfvars, m)
			}
		}

		ninit.Append(psess.nod(ODCL, m, nil))
		retvars = append(retvars, m)
	}

	as := psess.nod(OAS2, nil, nil)
	as.Rlist.Set(n.List.Slice())

	// For non-dotted calls to variadic functions, we assign the
	// variadic parameter's temp name separately.
	var vas *Node

	if fn.IsMethod(psess) {
		rcv := fn.Type.Recv(psess.types)

		if n.Left.Op == ODOTMETH {

			if n.Left.Left == nil {
				psess.
					Fatalf("method call without receiver: %+v", n)
			}
			ras := psess.nod(OAS, psess.tinlvar(rcv, inlvars), n.Left.Left)
			ras = psess.typecheck(ras, Etop)
			ninit.Append(ras)
		} else {

			if as.Rlist.Len() == 0 {
				psess.
					Fatalf("non-method call to method without first arg: %+v", n)
			}
			as.List.Append(psess.tinlvar(rcv, inlvars))
		}
	}

	for _, param := range fn.Type.Params(psess.types).Fields(psess.types).Slice() {

		if !param.Isddd() || n.Isddd() {
			as.List.Append(psess.tinlvar(param, inlvars))
			continue
		}

		numvals := n.List.Len()
		if numvals == 1 && n.List.First().Type.IsFuncArgStruct() {
			numvals = n.List.First().Type.NumFields(psess.types)
		}

		x := as.List.Len()
		for as.List.Len() < numvals {
			as.List.Append(psess.argvar(param.Type, as.List.Len()))
		}
		varargs := as.List.Slice()[x:]

		vas = psess.nod(OAS, psess.tinlvar(param, inlvars), nil)
		if len(varargs) == 0 {
			vas.Right = psess.nodnil()
			vas.Right.Type = param.Type
		} else {
			vas.Right = psess.nod(OCOMPLIT, nil, psess.typenod(param.Type))
			vas.Right.List.Set(varargs)
		}
	}

	if as.Rlist.Len() != 0 {
		as = psess.typecheck(as, Etop)
		ninit.Append(as)
	}

	if vas != nil {
		vas = psess.typecheck(vas, Etop)
		ninit.Append(vas)
	}

	for _, n := range retvars {
		ras := psess.nod(OAS, n, nil)
		ras = psess.typecheck(ras, Etop)
		ninit.Append(ras)
	}

	retlabel := psess.autolabel(".i")
	psess.
		inlgen++

	parent := -1
	if b := psess.Ctxt.PosTable.Pos(n.Pos).Base(); b != nil {
		parent = b.InliningIndex()
	}
	newIndex := psess.Ctxt.InlTree.Add(parent, n.Pos, fn.Sym.Linksym(psess.types))

	if psess.genDwarfInline > 0 {
		if !fn.Sym.Linksym(psess.types).WasInlined() {
			psess.
				Ctxt.DwFixups.SetPrecursorFunc(fn.Sym.Linksym(psess.types), fn)
			fn.Sym.Linksym(psess.types).Set(obj.AttrWasInlined, true)
		}
	}

	subst := inlsubst{
		retlabel:    retlabel,
		retvars:     retvars,
		inlvars:     inlvars,
		bases:       make(map[*src.PosBase]*src.PosBase),
		newInlIndex: newIndex,
	}

	body := subst.list(psess, asNodes(fn.Func.Inl.Body))

	lab := psess.nod(OLABEL, retlabel, nil)
	body = append(body, lab)
	psess.
		typecheckslice(body, Etop)

	if psess.genDwarfInline > 0 {
		for _, v := range inlfvars {
			v.Pos = subst.updatedPos(psess, v.Pos)
		}
	}

	call := psess.nod(OINLCALL, nil, nil)
	call.Ninit.Set(ninit.Slice())
	call.Nbody.Set(body)
	call.Rlist.Set(retvars)
	call.Type = n.Type
	call.SetTypecheck(1)
	psess.
		inlnodelist(call.Nbody)
	for _, n := range call.Nbody.Slice() {
		if n.Op == OINLCALL {
			inlconv2stmt(n)
		}
	}

	if psess.Debug['m'] > 2 {
		fmt.Printf("%v: After inlining %+v\n\n", call.Line(psess), call)
	}

	return call
}

// Every time we expand a function we generate a new set of tmpnames,
// PAUTO's in the calling functions, and link them off of the
// PPARAM's, PAUTOS and PPARAMOUTs of the called function.
func (psess *PackageSession) inlvar(var_ *Node) *Node {
	if psess.Debug['m'] > 3 {
		fmt.Printf("inlvar %+v\n", var_)
	}

	n := psess.newname(var_.Sym)
	n.Type = var_.Type
	n.SetClass(PAUTO)
	n.Name.SetUsed(true)
	n.Name.Curfn = psess.Curfn
	n.SetAddrtaken(var_.Addrtaken())
	psess.
		Curfn.Func.Dcl = append(psess.Curfn.Func.Dcl, n)
	return n
}

// Synthesize a variable to store the inlined function's results in.
func (psess *PackageSession) retvar(t *types.Field, i int) *Node {
	n := psess.newname(psess.lookupN("~R", i))
	n.Type = t.Type
	n.SetClass(PAUTO)
	n.Name.SetUsed(true)
	n.Name.Curfn = psess.Curfn
	psess.
		Curfn.Func.Dcl = append(psess.Curfn.Func.Dcl, n)
	return n
}

// Synthesize a variable to store the inlined function's arguments
// when they come from a multiple return call.
func (psess *PackageSession) argvar(t *types.Type, i int) *Node {
	n := psess.newname(psess.lookupN("~arg", i))
	n.Type = t.Elem(psess.types)
	n.SetClass(PAUTO)
	n.Name.SetUsed(true)
	n.Name.Curfn = psess.Curfn
	psess.
		Curfn.Func.Dcl = append(psess.Curfn.Func.Dcl, n)
	return n
}

// The inlsubst type implements the actual inlining of a single
// function call.
type inlsubst struct {
	// Target of the goto substituted in place of a return.
	retlabel *Node

	// Temporary result variables.
	retvars []*Node

	inlvars map[*Node]*Node

	// bases maps from original PosBase to PosBase with an extra
	// inlined call frame.
	bases map[*src.PosBase]*src.PosBase

	// newInlIndex is the index of the inlined call frame to
	// insert for inlined nodes.
	newInlIndex int
}

// list inlines a list of nodes.
func (subst *inlsubst) list(psess *PackageSession, ll Nodes) []*Node {
	s := make([]*Node, 0, ll.Len())
	for _, n := range ll.Slice() {
		s = append(s, subst.node(psess, n))
	}
	return s
}

// node recursively copies a node from the saved pristine body of the
// inlined function, substituting references to input/output
// parameters with ones to the tmpnames, and substituting returns with
// assignments to the output.
func (subst *inlsubst) node(psess *PackageSession, n *Node) *Node {
	if n == nil {
		return nil
	}

	switch n.Op {
	case ONAME:
		if inlvar := subst.inlvars[n]; inlvar != nil {
			if psess.Debug['m'] > 2 {
				fmt.Printf("substituting name %+v  ->  %+v\n", n, inlvar)
			}
			return inlvar
		}

		if psess.Debug['m'] > 2 {
			fmt.Printf("not substituting name %+v\n", n)
		}
		return n

	case OLITERAL, OTYPE:

		if n.Sym != nil {
			return n
		}

	case ORETURN:
		m := psess.nod(OGOTO, subst.retlabel, nil)
		m.Ninit.Set(subst.list(psess, n.Ninit))

		if len(subst.retvars) != 0 && n.List.Len() != 0 {
			as := psess.nod(OAS2, nil, nil)

			for _, n := range subst.retvars {
				as.List.Append(n)
			}
			as.Rlist.Set(subst.list(psess, n.List))
			as = psess.typecheck(as, Etop)
			m.Ninit.Append(as)
		}
		psess.
			typecheckslice(m.Ninit.Slice(), Etop)
		m = psess.typecheck(m, Etop)

		return m

	case OGOTO, OLABEL:
		m := n.copy()
		m.Pos = subst.updatedPos(psess, m.Pos)
		m.Ninit.Set(nil)
		p := fmt.Sprintf("%s·%d", n.Left.Sym.Name, psess.inlgen)
		m.Left = psess.newname(psess.lookup(p))

		return m
	}

	m := n.copy()
	m.Pos = subst.updatedPos(psess, m.Pos)
	m.Ninit.Set(nil)

	if n.Op == OCLOSURE {
		psess.
			Fatalf("cannot inline function containing closure: %+v", n)
	}

	m.Left = subst.node(psess, n.Left)
	m.Right = subst.node(psess, n.Right)
	m.List.Set(subst.list(psess, n.List))
	m.Rlist.Set(subst.list(psess, n.Rlist))
	m.Ninit.Set(append(m.Ninit.Slice(), subst.list(psess, n.Ninit)...))
	m.Nbody.Set(subst.list(psess, n.Nbody))

	return m
}

func (subst *inlsubst) updatedPos(psess *PackageSession, xpos src.XPos) src.XPos {
	pos := psess.Ctxt.PosTable.Pos(xpos)
	oldbase := pos.Base()
	newbase := subst.bases[oldbase]
	if newbase == nil {
		newbase = src.NewInliningBase(oldbase, subst.newInlIndex)
		subst.bases[oldbase] = newbase
	}
	pos.SetBase(newbase)
	return psess.Ctxt.PosTable.XPos(pos)
}

func pruneUnusedAutos(ll []*Node, vis *hairyVisitor) []*Node {
	s := make([]*Node, 0, len(ll))
	for _, n := range ll {
		if n.Class() == PAUTO {
			if _, found := vis.usedLocals[n]; !found {
				continue
			}
		}
		s = append(s, n)
	}
	return s
}

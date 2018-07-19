package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
)

// Order holds state during the ordering process.
type Order struct {
	out  []*Node // list of generated statements
	temp []*Node // stack of temporary variables
}

// Order rewrites fn.Nbody to apply the ordering constraints
// described in the comment at the top of the file.
func (psess *PackageSession) order(fn *Node) {
	if psess.Debug['W'] > 1 {
		s := fmt.Sprintf("\nbefore order %v", fn.Func.Nname.Sym)
		dumplist(s, fn.Nbody)
	}
	psess.
		orderBlock(&fn.Nbody)
}

// newTemp allocates a new temporary with the given type,
// pushes it onto the temp stack, and returns it.
// If clear is true, newTemp emits code to zero the temporary.
func (o *Order) newTemp(psess *PackageSession, t *types.Type, clear bool) *Node {
	v := psess.temp(t)
	if clear {
		a := psess.nod(OAS, v, nil)
		a = psess.typecheck(a, Etop)
		o.out = append(o.out, a)
	}

	o.temp = append(o.temp, v)
	return v
}

// copyExpr behaves like ordertemp but also emits
// code to initialize the temporary to the value n.
//
// The clear argument is provided for use when the evaluation
// of tmp = n turns into a function call that is passed a pointer
// to the temporary as the output space. If the call blocks before
// tmp has been written, the garbage collector will still treat the
// temporary as live, so we must zero it before entering that call.
// Today, this only happens for channel receive operations.
// (The other candidate would be map access, but map access
// returns a pointer to the result data instead of taking a pointer
// to be filled in.)
func (o *Order) copyExpr(psess *PackageSession, n *Node, t *types.Type, clear bool) *Node {
	v := o.newTemp(psess, t, clear)
	a := psess.nod(OAS, v, n)
	a = psess.typecheck(a, Etop)
	o.out = append(o.out, a)
	return v
}

// cheapExpr returns a cheap version of n.
// The definition of cheap is that n is a variable or constant.
// If not, cheapExpr allocates a new tmp, emits tmp = n,
// and then returns tmp.
func (o *Order) cheapExpr(psess *PackageSession, n *Node) *Node {
	if n == nil {
		return nil
	}

	switch n.Op {
	case ONAME, OLITERAL:
		return n
	case OLEN, OCAP:
		l := o.cheapExpr(psess, n.Left)
		if l == n.Left {
			return n
		}
		a := n.copy()
		a.Orig = a
		a.Left = l
		return psess.typecheck(a, Erv)
	}

	return o.copyExpr(psess, n, n.Type, false)
}

// safeExpr returns a safe version of n.
// The definition of safe is that n can appear multiple times
// without violating the semantics of the original program,
// and that assigning to the safe version has the same effect
// as assigning to the original n.
//
// The intended use is to apply to x when rewriting x += y into x = x + y.
func (o *Order) safeExpr(psess *PackageSession, n *Node) *Node {
	switch n.Op {
	case ONAME, OLITERAL:
		return n

	case ODOT, OLEN, OCAP:
		l := o.safeExpr(psess, n.Left)
		if l == n.Left {
			return n
		}
		a := n.copy()
		a.Orig = a
		a.Left = l
		return psess.typecheck(a, Erv)

	case ODOTPTR, OIND:
		l := o.cheapExpr(psess, n.Left)
		if l == n.Left {
			return n
		}
		a := n.copy()
		a.Orig = a
		a.Left = l
		return psess.typecheck(a, Erv)

	case OINDEX, OINDEXMAP:
		var l *Node
		if n.Left.Type.IsArray() {
			l = o.safeExpr(psess, n.Left)
		} else {
			l = o.cheapExpr(psess, n.Left)
		}
		r := o.cheapExpr(psess, n.Right)
		if l == n.Left && r == n.Right {
			return n
		}
		a := n.copy()
		a.Orig = a
		a.Left = l
		a.Right = r
		return psess.typecheck(a, Erv)

	default:
		psess.
			Fatalf("ordersafeexpr %v", n.Op)
		return nil
	}
}

// Isaddrokay reports whether it is okay to pass n's address to runtime routines.
// Taking the address of a variable makes the liveness and optimization analyses
// lose track of where the variable's lifetime ends. To avoid hurting the analyses
// of ordinary stack variables, those are not 'isaddrokay'. Temporaries are okay,
// because we emit explicit VARKILL instructions marking the end of those
// temporaries' lifetimes.
func isaddrokay(n *Node) bool {
	return islvalue(n) && (n.Op != ONAME || n.Class() == PEXTERN || n.IsAutoTmp())
}

// addrTemp ensures that n is okay to pass by address to runtime routines.
// If the original argument n is not okay, addrTemp creates a tmp, emits
// tmp = n, and then returns tmp.
// The result of addrTemp MUST be assigned back to n, e.g.
// 	n.Left = o.addrTemp(n.Left)
func (o *Order) addrTemp(psess *PackageSession, n *Node) *Node {
	if psess.consttype(n) > 0 {

		n = psess.defaultlit(n, nil)
		psess.
			dowidth(n.Type)
		vstat := psess.staticname(n.Type)
		vstat.Name.SetReadonly(true)
		var out []*Node
		psess.
			staticassign(vstat, n, &out)
		if out != nil {
			psess.
				Fatalf("staticassign of const generated code: %+v", n)
		}
		vstat = psess.typecheck(vstat, Erv)
		return vstat
	}
	if isaddrokay(n) {
		return n
	}
	return o.copyExpr(psess, n, n.Type, false)
}

// mapKeyTemp prepares n to be a key in a map runtime call and returns n.
// It should only be used for map runtime calls which have *_fast* versions.
func (o *Order) mapKeyTemp(psess *PackageSession, t *types.Type, n *Node) *Node {

	if psess.mapfast(t) == mapslow {
		return o.addrTemp(psess, n)
	}
	return n
}

type ordermarker int

// Marktemp returns the top of the temporary variable stack.
func (o *Order) markTemp() ordermarker {
	return ordermarker(len(o.temp))
}

// Poptemp pops temporaries off the stack until reaching the mark,
// which must have been returned by marktemp.
func (o *Order) popTemp(mark ordermarker) {
	o.temp = o.temp[:mark]
}

// Cleantempnopop emits VARKILL and if needed VARLIVE instructions
// to *out for each temporary above the mark on the temporary stack.
// It does not pop the temporaries from the stack.
func (o *Order) cleanTempNoPop(psess *PackageSession, mark ordermarker) []*Node {
	var out []*Node
	for i := len(o.temp) - 1; i >= int(mark); i-- {
		n := o.temp[i]
		if n.Name.Keepalive() {
			n.Name.SetKeepalive(false)
			n.SetAddrtaken(true)
			live := psess.nod(OVARLIVE, n, nil)
			live = psess.typecheck(live, Etop)
			out = append(out, live)
		}
		kill := psess.nod(OVARKILL, n, nil)
		kill = psess.typecheck(kill, Etop)
		out = append(out, kill)
	}
	return out
}

// cleanTemp emits VARKILL instructions for each temporary above the
// mark on the temporary stack and removes them from the stack.
func (o *Order) cleanTemp(psess *PackageSession, top ordermarker) {
	o.out = append(o.out, o.cleanTempNoPop(psess, top)...)
	o.popTemp(top)
}

// stmtList orders each of the statements in the list.
func (o *Order) stmtList(psess *PackageSession, l Nodes) {
	for _, n := range l.Slice() {
		o.stmt(psess, n)
	}
}

// orderBlock orders the block of statements in n into a new slice,
// and then replaces the old slice in n with the new slice.
func (psess *PackageSession) orderBlock(n *Nodes) {
	var order Order
	mark := order.markTemp()
	order.stmtList(psess, *n)
	order.cleanTemp(psess, mark)
	n.Set(order.out)
}

// exprInPlace orders the side effects in *np and
// leaves them as the init list of the final *np.
// The result of exprInPlace MUST be assigned back to n, e.g.
// 	n.Left = o.exprInPlace(n.Left)
func (o *Order) exprInPlace(psess *PackageSession, n *Node) *Node {
	var order Order
	n = order.expr(psess, n, nil)
	n = psess.addinit(n, order.out)

	o.temp = append(o.temp, order.temp...)
	return n
}

// orderStmtInPlace orders the side effects of the single statement *np
// and replaces it with the resulting statement list.
// The result of orderStmtInPlace MUST be assigned back to n, e.g.
// 	n.Left = orderStmtInPlace(n.Left)
func (psess *PackageSession) orderStmtInPlace(n *Node) *Node {
	var order Order
	mark := order.markTemp()
	order.stmt(psess, n)
	order.cleanTemp(psess, mark)
	return psess.liststmt(order.out)
}

// init moves n's init list to o.out.
func (o *Order) init(psess *PackageSession, n *Node) {
	if n.mayBeShared() {

		if n.Ninit.Len() > 0 {
			psess.
				Fatalf("orderinit shared node with ninit")
		}
		return
	}
	o.stmtList(psess, n.Ninit)
	n.Ninit.Set(nil)
}

// Ismulticall reports whether the list l is f() for a multi-value function.
// Such an f() could appear as the lone argument to a multi-arg function.
func (psess *PackageSession) ismulticall(l Nodes) bool {

	if l.Len() != 1 {
		return false
	}
	n := l.First()

	switch n.Op {
	default:
		return false
	case OCALLFUNC, OCALLMETH, OCALLINTER:

		return n.Left.Type.NumResults(psess.types) > 1
	}
}

// copyRet emits t1, t2, ... = n, where n is a function call,
// and then returns the list t1, t2, ....
func (o *Order) copyRet(psess *PackageSession, n *Node) []*Node {
	if !n.Type.IsFuncArgStruct() {
		psess.
			Fatalf("copyret %v %d", n.Type, n.Left.Type.NumResults(psess.types))
	}

	var l1, l2 []*Node
	for _, f := range n.Type.Fields(psess.types).Slice() {
		tmp := psess.temp(f.Type)
		l1 = append(l1, tmp)
		l2 = append(l2, tmp)
	}

	as := psess.nod(OAS2, nil, nil)
	as.List.Set(l1)
	as.Rlist.Set1(n)
	as = psess.typecheck(as, Etop)
	o.stmt(psess, as)

	return l2
}

// callArgs orders the list of call arguments *l.
func (o *Order) callArgs(psess *PackageSession, l *Nodes) {
	if psess.ismulticall(*l) {

		l.Set(o.copyRet(psess, l.First()))
	} else {
		o.exprList(psess, *l)
	}
}

// call orders the call expression n.
// n.Op is OCALLMETH/OCALLFUNC/OCALLINTER or a builtin like OCOPY.
func (o *Order) call(psess *PackageSession, n *Node) {
	n.Left = o.expr(psess, n.Left, nil)
	n.Right = o.expr(psess, n.Right, nil)
	o.callArgs(psess, &n.List)

	if n.Op != OCALLFUNC {
		return
	}
	keepAlive := func(i int) {

		xp := n.List.Addr(i)
		for (*xp).Op == OCONVNOP && !(*xp).Type.IsUnsafePtr() {
			xp = &(*xp).Left
		}
		x := *xp
		if x.Type.IsUnsafePtr() {
			x = o.copyExpr(psess, x, x.Type, false)
			x.Name.SetKeepalive(true)
			*xp = x
		}
	}

	for i, t := range n.Left.Type.Params(psess.types).FieldSlice(psess.types) {

		if t.Isddd() && !n.Isddd() {
			if t.Note == uintptrEscapesTag {
				for ; i < n.List.Len(); i++ {
					keepAlive(i)
				}
			}
		} else {
			if t.Note == unsafeUintptrTag || t.Note == uintptrEscapesTag {
				keepAlive(i)
			}
		}
	}
}

// mapAssign appends n to o.out, introducing temporaries
// to make sure that all map assignments have the form m[k] = x.
// (Note: expr has already been called on n, so we know k is addressable.)
//
// If n is the multiple assignment form ..., m[k], ... = ..., x, ..., the rewrite is
//	t1 = m
//	t2 = k
//	...., t3, ... = ..., x, ...
//	t1[t2] = t3
//
// The temporaries t1, t2 are needed in case the ... being assigned
// contain m or k. They are usually unnecessary, but in the unnecessary
// cases they are also typically registerizable, so not much harm done.
// And this only applies to the multiple-assignment form.
// We could do a more precise analysis if needed, like in walk.go.
func (o *Order) mapAssign(psess *PackageSession, n *Node) {
	switch n.Op {
	default:
		psess.
			Fatalf("ordermapassign %v", n.Op)

	case OAS, OASOP:
		if n.Left.Op == OINDEXMAP {

			if n.Right.Op == OAPPEND {
				s := n.Right.List.Slice()[1:]
				for i, n := range s {
					s[i] = o.cheapExpr(psess, n)
				}
			} else {
				n.Right = o.cheapExpr(psess, n.Right)
			}
		}
		o.out = append(o.out, n)

	case OAS2, OAS2DOTTYPE, OAS2MAPR, OAS2FUNC:
		var post []*Node
		for i, m := range n.List.Slice() {
			switch {
			case m.Op == OINDEXMAP:
				if !m.Left.IsAutoTmp() {
					m.Left = o.copyExpr(psess, m.Left, m.Left.Type, false)
				}
				if !m.Right.IsAutoTmp() {
					m.Right = o.copyExpr(psess, m.Right, m.Right.Type, false)
				}
				fallthrough
			case psess.instrumenting && n.Op == OAS2FUNC && !m.isBlank():
				t := o.newTemp(psess, m.Type, false)
				n.List.SetIndex(i, t)
				a := psess.nod(OAS, m, t)
				a = psess.typecheck(a, Etop)
				post = append(post, a)
			}
		}

		o.out = append(o.out, n)
		o.out = append(o.out, post...)
	}
}

// stmt orders the statement n, appending to o.out.
// Temporaries created during the statement are cleaned
// up using VARKILL instructions as possible.
func (o *Order) stmt(psess *PackageSession, n *Node) {
	if n == nil {
		return
	}

	lno := psess.setlineno(n)
	o.init(psess, n)

	switch n.Op {
	default:
		psess.
			Fatalf("orderstmt %v", n.Op)

	case OVARKILL, OVARLIVE:
		o.out = append(o.out, n)

	case OAS:
		t := o.markTemp()
		n.Left = o.expr(psess, n.Left, nil)
		n.Right = o.expr(psess, n.Right, n.Left)
		o.mapAssign(psess, n)
		o.cleanTemp(psess, t)

	case OAS2,
		OCLOSE,
		OCOPY,
		OPRINT,
		OPRINTN,
		ORECOVER,
		ORECV:
		t := o.markTemp()
		n.Left = o.expr(psess, n.Left, nil)
		n.Right = o.expr(psess, n.Right, nil)
		o.exprList(psess, n.List)
		o.exprList(psess, n.Rlist)
		switch n.Op {
		case OAS2:
			o.mapAssign(psess, n)
		default:
			o.out = append(o.out, n)
		}
		o.cleanTemp(psess, t)

	case OASOP:
		t := o.markTemp()
		n.Left = o.expr(psess, n.Left, nil)
		n.Right = o.expr(psess, n.Right, nil)

		if psess.instrumenting || n.Left.Op == OINDEXMAP && (n.SubOp(psess) == ODIV || n.SubOp(psess) == OMOD) {

			n.Left = o.safeExpr(psess, n.Left)

			l := psess.treecopy(n.Left, psess.src.NoXPos)
			if l.Op == OINDEXMAP {
				l.SetIndexMapLValue(psess, false)
			}
			l = o.copyExpr(psess, l, n.Left.Type, false)
			n.Right = psess.nod(n.SubOp(psess), l, n.Right)
			n.Right = psess.typecheck(n.Right, Erv)
			n.Right = o.expr(psess, n.Right, nil)

			n.Op = OAS
			n.ResetAux()
		}

		o.mapAssign(psess, n)
		o.cleanTemp(psess, t)

	case OAS2MAPR:
		t := o.markTemp()
		o.exprList(psess, n.List)
		r := n.Rlist.First()
		r.Left = o.expr(psess, r.Left, nil)
		r.Right = o.expr(psess, r.Right, nil)

		if r.Right.Op == OARRAYBYTESTR {
			r.Right.Op = OARRAYBYTESTRTMP
		}
		r.Right = o.mapKeyTemp(psess, r.Left.Type, r.Right)
		o.okAs2(psess, n)
		o.cleanTemp(psess, t)

	case OAS2FUNC:
		t := o.markTemp()
		o.exprList(psess, n.List)
		o.call(psess, n.Rlist.First())
		o.as2(psess, n)
		o.cleanTemp(psess, t)

	case OAS2DOTTYPE:
		t := o.markTemp()
		o.exprList(psess, n.List)
		n.Rlist.First().Left = o.expr(psess, n.Rlist.First().Left, nil)
		o.okAs2(psess, n)
		o.cleanTemp(psess, t)

	case OAS2RECV:
		t := o.markTemp()
		o.exprList(psess, n.List)
		n.Rlist.First().Left = o.expr(psess, n.Rlist.First().Left, nil)
		ch := n.Rlist.First().Left.Type
		tmp1 := o.newTemp(psess, ch.Elem(psess.types), psess.types.Haspointers(ch.Elem(psess.types)))
		tmp2 := o.newTemp(psess, psess.types.Types[TBOOL], false)
		o.out = append(o.out, n)
		r := psess.nod(OAS, n.List.First(), tmp1)
		r = psess.typecheck(r, Etop)
		o.mapAssign(psess, r)
		r = psess.okas(n.List.Second(), tmp2)
		r = psess.typecheck(r, Etop)
		o.mapAssign(psess, r)
		n.List.Set2(tmp1, tmp2)
		o.cleanTemp(psess, t)

	case OBLOCK, OEMPTY:
		o.stmtList(psess, n.List)

	case OBREAK,
		OCONTINUE,
		ODCL,
		ODCLCONST,
		ODCLTYPE,
		OFALL,
		OGOTO,
		OLABEL,
		ORETJMP:
		o.out = append(o.out, n)

	case OCALLFUNC, OCALLINTER, OCALLMETH:
		t := o.markTemp()
		o.call(psess, n)
		o.out = append(o.out, n)
		o.cleanTemp(psess, t)

	case ODEFER, OPROC:
		t := o.markTemp()
		o.call(psess, n.Left)
		o.out = append(o.out, n)
		o.cleanTemp(psess, t)

	case ODELETE:
		t := o.markTemp()
		n.List.SetFirst(o.expr(psess, n.List.First(), nil))
		n.List.SetSecond(o.expr(psess, n.List.Second(), nil))
		n.List.SetSecond(o.mapKeyTemp(psess, n.List.First().Type, n.List.Second()))
		o.out = append(o.out, n)
		o.cleanTemp(psess, t)

	case OFOR:
		t := o.markTemp()
		n.Left = o.exprInPlace(psess, n.Left)
		n.Nbody.Prepend(o.cleanTempNoPop(psess, t)...)
		psess.
			orderBlock(&n.Nbody)
		n.Right = psess.orderStmtInPlace(n.Right)
		o.out = append(o.out, n)
		o.cleanTemp(psess, t)

	case OIF:
		t := o.markTemp()
		n.Left = o.exprInPlace(psess, n.Left)
		n.Nbody.Prepend(o.cleanTempNoPop(psess, t)...)
		n.Rlist.Prepend(o.cleanTempNoPop(psess, t)...)
		o.popTemp(t)
		psess.
			orderBlock(&n.Nbody)
		psess.
			orderBlock(&n.Rlist)
		o.out = append(o.out, n)

	case OPANIC:
		t := o.markTemp()
		n.Left = o.expr(psess, n.Left, nil)
		if !n.Left.Type.IsInterface() {
			n.Left = o.addrTemp(psess, n.Left)
		}
		o.out = append(o.out, n)
		o.cleanTemp(psess, t)

	case ORANGE:

		if n.Right.Op == OSTRARRAYBYTE {
			n.Right.Op = OSTRARRAYBYTETMP
		}

		t := o.markTemp()
		n.Right = o.expr(psess, n.Right, nil)

		orderBody := true
		switch n.Type.Etype {
		default:
			psess.
				Fatalf("orderstmt range %v", n.Type)

		case TARRAY, TSLICE:
			if n.List.Len() < 2 || n.List.Second().isBlank() {

				break
			}
			fallthrough

		case TCHAN, TSTRING:

			r := n.Right

			if r.Type.IsString() && r.Type != psess.types.Types[TSTRING] {
				r = psess.nod(OCONV, r, nil)
				r.Type = psess.types.Types[TSTRING]
				r = psess.typecheck(r, Erv)
			}

			n.Right = o.copyExpr(psess, r, r.Type, false)

		case TMAP:
			if psess.isMapClear(n) {

				orderBody = false
				break
			}

			r := n.Right
			n.Right = o.copyExpr(psess, r, r.Type, false)
			psess.
				prealloc[n] = o.newTemp(psess, psess.hiter(n.Type), true)
		}
		o.exprListInPlace(psess, n.List)
		if orderBody {
			psess.
				orderBlock(&n.Nbody)
		}
		o.out = append(o.out, n)
		o.cleanTemp(psess, t)

	case ORETURN:
		o.callArgs(psess, &n.List)
		o.out = append(o.out, n)

	case OSELECT:
		t := o.markTemp()

		for _, n2 := range n.List.Slice() {
			if n2.Op != OXCASE {
				psess.
					Fatalf("order select case %v", n2.Op)
			}
			r := n2.Left
			psess.
				setlineno(n2)

			if n2.Ninit.Len() != 0 {
				psess.
					Fatalf("order select ninit")
			}
			if r == nil {
				continue
			}
			switch r.Op {
			default:
				Dump("select case", r)
				psess.
					Fatalf("unknown op in select %v", r.Op)

			case OSELRECV, OSELRECV2:
				if r.Colas() {
					i := 0
					if r.Ninit.Len() != 0 && r.Ninit.First().Op == ODCL && r.Ninit.First().Left == r.Left {
						i++
					}
					if i < r.Ninit.Len() && r.Ninit.Index(i).Op == ODCL && r.List.Len() != 0 && r.Ninit.Index(i).Left == r.List.First() {
						i++
					}
					if i >= r.Ninit.Len() {
						r.Ninit.Set(nil)
					}
				}

				if r.Ninit.Len() != 0 {
					dumplist("ninit", r.Ninit)
					psess.
						Fatalf("ninit on select recv")
				}

				r.Right.Left = o.expr(psess, r.Right.Left, nil)

				if r.Right.Left.Op != ONAME {
					r.Right.Left = o.copyExpr(psess, r.Right.Left, r.Right.Left.Type, false)
				}

				if r.Left != nil && r.Left.isBlank() {
					r.Left = nil
				}
				if r.Left != nil {

					tmp1 := r.Left

					if r.Colas() {
						tmp2 := psess.nod(ODCL, tmp1, nil)
						tmp2 = psess.typecheck(tmp2, Etop)
						n2.Ninit.Append(tmp2)
					}

					r.Left = o.newTemp(psess, r.Right.Left.Type.Elem(psess.types), psess.types.Haspointers(r.Right.Left.Type.Elem(psess.types)))
					tmp2 := psess.nod(OAS, tmp1, r.Left)
					tmp2 = psess.typecheck(tmp2, Etop)
					n2.Ninit.Append(tmp2)
				}

				if r.List.Len() != 0 && r.List.First().isBlank() {
					r.List.Set(nil)
				}
				if r.List.Len() != 0 {
					tmp1 := r.List.First()
					if r.Colas() {
						tmp2 := psess.nod(ODCL, tmp1, nil)
						tmp2 = psess.typecheck(tmp2, Etop)
						n2.Ninit.Append(tmp2)
					}

					r.List.Set1(o.newTemp(psess, psess.types.Types[TBOOL], false))
					tmp2 := psess.okas(tmp1, r.List.First())
					tmp2 = psess.typecheck(tmp2, Etop)
					n2.Ninit.Append(tmp2)
				}
				psess.
					orderBlock(&n2.Ninit)

			case OSEND:
				if r.Ninit.Len() != 0 {
					dumplist("ninit", r.Ninit)
					psess.
						Fatalf("ninit on select send")
				}

				r.Left = o.expr(psess, r.Left, nil)

				if !r.Left.IsAutoTmp() {
					r.Left = o.copyExpr(psess, r.Left, r.Left.Type, false)
				}
				r.Right = o.expr(psess, r.Right, nil)
				if !r.Right.IsAutoTmp() {
					r.Right = o.copyExpr(psess, r.Right, r.Right.Type, false)
				}
			}
		}

		for _, n3 := range n.List.Slice() {
			psess.
				orderBlock(&n3.Nbody)
			n3.Nbody.Prepend(o.cleanTempNoPop(psess, t)...)

			n3.Nbody.Prepend(n3.Ninit.Slice()...)
			n3.Ninit.Set(nil)
		}

		o.out = append(o.out, n)
		o.popTemp(t)

	case OSEND:
		t := o.markTemp()
		n.Left = o.expr(psess, n.Left, nil)
		n.Right = o.expr(psess, n.Right, nil)
		if psess.instrumenting {

			n.Right = o.copyExpr(psess, n.Right, n.Right.Type, false)
		} else {
			n.Right = o.addrTemp(psess, n.Right)
		}
		o.out = append(o.out, n)
		o.cleanTemp(psess, t)

	case OSWITCH:
		t := o.markTemp()
		n.Left = o.expr(psess, n.Left, nil)
		for _, ncas := range n.List.Slice() {
			if ncas.Op != OXCASE {
				psess.
					Fatalf("order switch case %v", ncas.Op)
			}
			o.exprListInPlace(psess, ncas.List)
			psess.
				orderBlock(&ncas.Nbody)
		}

		o.out = append(o.out, n)
		o.cleanTemp(psess, t)
	}
	psess.
		lineno = lno
}

// exprList orders the expression list l into o.
func (o *Order) exprList(psess *PackageSession, l Nodes) {
	s := l.Slice()
	for i := range s {
		s[i] = o.expr(psess, s[i], nil)
	}
}

// exprListInPlace orders the expression list l but saves
// the side effects on the individual expression ninit lists.
func (o *Order) exprListInPlace(psess *PackageSession, l Nodes) {
	s := l.Slice()
	for i := range s {
		s[i] = o.exprInPlace(psess, s[i])
	}
}

// prealloc[x] records the allocation to use for x.

// expr orders a single expression, appending side
// effects to o.out as needed.
// If this is part of an assignment lhs = *np, lhs is given.
// Otherwise lhs == nil. (When lhs != nil it may be possible
// to avoid copying the result of the expression to a temporary.)
// The result of expr MUST be assigned back to n, e.g.
// 	n.Left = o.expr(n.Left, lhs)
func (o *Order) expr(psess *PackageSession, n, lhs *Node) *Node {
	if n == nil {
		return n
	}

	lno := psess.setlineno(n)
	o.init(psess, n)

	switch n.Op {
	default:
		n.Left = o.expr(psess, n.Left, nil)
		n.Right = o.expr(psess, n.Right, nil)
		o.exprList(psess, n.List)
		o.exprList(psess, n.Rlist)

	case OADDSTR:
		o.exprList(psess, n.List)

		if n.List.Len() > 5 {
			t := psess.types.NewArray(psess.types.Types[TSTRING], int64(n.List.Len()))
			psess.
				prealloc[n] = o.newTemp(psess, t, false)
		}

		hasbyte := false

		haslit := false
		for _, n1 := range n.List.Slice() {
			hasbyte = hasbyte || n1.Op == OARRAYBYTESTR
			haslit = haslit || n1.Op == OLITERAL && len(n1.Val().U.(string)) != 0
		}

		if haslit && hasbyte {
			for _, n2 := range n.List.Slice() {
				if n2.Op == OARRAYBYTESTR {
					n2.Op = OARRAYBYTESTRTMP
				}
			}
		}

	case OCMPSTR:
		n.Left = o.expr(psess, n.Left, nil)
		n.Right = o.expr(psess, n.Right, nil)

		if n.Left.Op == OARRAYBYTESTR {
			n.Left.Op = OARRAYBYTESTRTMP
		}
		if n.Right.Op == OARRAYBYTESTR {
			n.Right.Op = OARRAYBYTESTRTMP
		}

	case OINDEXMAP:
		n.Left = o.expr(psess, n.Left, nil)
		n.Right = o.expr(psess, n.Right, nil)
		needCopy := false

		if !n.IndexMapLValue(psess) && psess.instrumenting {

			needCopy = true
		}

		if !n.IndexMapLValue(psess) && n.Right.Op == OARRAYBYTESTR {
			n.Right.Op = OARRAYBYTESTRTMP
			needCopy = true
		}

		n.Right = o.mapKeyTemp(psess, n.Left.Type, n.Right)
		if needCopy {
			n = o.copyExpr(psess, n, n.Type, false)
		}

	case OCONVIFACE:
		n.Left = o.expr(psess, n.Left, nil)

		if !n.Left.Type.IsInterface() {
			n.Left = o.addrTemp(psess, n.Left)
		}

	case OCONVNOP:
		if n.Type.IsKind(TUNSAFEPTR) && n.Left.Type.IsKind(TUINTPTR) && (n.Left.Op == OCALLFUNC || n.Left.Op == OCALLINTER || n.Left.Op == OCALLMETH) {

			o.init(psess, n.Left)
			o.call(psess, n.Left)
			if lhs == nil || lhs.Op != ONAME || psess.instrumenting {
				n = o.copyExpr(psess, n, n.Type, false)
			}
		} else {
			n.Left = o.expr(psess, n.Left, nil)
		}

	case OANDAND, OOROR:
		mark := o.markTemp()
		n.Left = o.expr(psess, n.Left, nil)

		n.Right = psess.addinit(n.Right, o.cleanTempNoPop(psess, mark))
		n.Right = o.exprInPlace(psess, n.Right)

	case OCALLFUNC,
		OCALLINTER,
		OCALLMETH,
		OCAP,
		OCOMPLEX,
		OCOPY,
		OIMAG,
		OLEN,
		OMAKECHAN,
		OMAKEMAP,
		OMAKESLICE,
		ONEW,
		OREAL,
		ORECOVER,
		OSTRARRAYBYTE,
		OSTRARRAYBYTETMP,
		OSTRARRAYRUNE:

		if psess.isRuneCount(n) {

			n.Left.Left = o.expr(psess, n.Left.Left, nil)
		} else {
			o.call(psess, n)
		}

		if lhs == nil || lhs.Op != ONAME || psess.instrumenting {
			n = o.copyExpr(psess, n, n.Type, false)
		}

	case OAPPEND:

		if psess.isAppendOfMake(n) {
			n.List.SetFirst(o.expr(psess, n.List.First(), nil))
			n.List.Second().Left = o.expr(psess, n.List.Second().Left, nil)
		} else {
			o.callArgs(psess, &n.List)
		}

		if lhs == nil || lhs.Op != ONAME && !psess.samesafeexpr(lhs, n.List.First()) {
			n = o.copyExpr(psess, n, n.Type, false)
		}

	case OSLICE, OSLICEARR, OSLICESTR, OSLICE3, OSLICE3ARR:
		n.Left = o.expr(psess, n.Left, nil)
		low, high, max := n.SliceBounds(psess)
		low = o.expr(psess, low, nil)
		low = o.cheapExpr(psess, low)
		high = o.expr(psess, high, nil)
		high = o.cheapExpr(psess, high)
		max = o.expr(psess, max, nil)
		max = o.cheapExpr(psess, max)
		n.SetSliceBounds(psess, low, high, max)
		if lhs == nil || lhs.Op != ONAME && !psess.samesafeexpr(lhs, n.Left) {
			n = o.copyExpr(psess, n, n.Type, false)
		}

	case OCLOSURE:
		if n.Noescape() && n.Func.Closure.Func.Cvars.Len() > 0 {
			psess.
				prealloc[n] = o.newTemp(psess, psess.types.Types[TUINT8], false)
		}

	case OARRAYLIT, OSLICELIT, OCALLPART:
		n.Left = o.expr(psess, n.Left, nil)
		n.Right = o.expr(psess, n.Right, nil)
		o.exprList(psess, n.List)
		o.exprList(psess, n.Rlist)
		if n.Noescape() {
			psess.
				prealloc[n] = o.newTemp(psess, psess.types.Types[TUINT8], false)
		}

	case ODDDARG:
		if n.Noescape() {
			psess.
				prealloc[n] = o.newTemp(psess, n.Type.Elem(psess.types), false)
		}

	case ODOTTYPE, ODOTTYPE2:
		n.Left = o.expr(psess, n.Left, nil)

		if !psess.isdirectiface(n.Type) || isfat(n.Type) || psess.instrumenting {
			n = o.copyExpr(psess, n, n.Type, true)
		}

	case ORECV:
		n.Left = o.expr(psess, n.Left, nil)
		n = o.copyExpr(psess, n, n.Type, true)

	case OEQ, ONE:
		n.Left = o.expr(psess, n.Left, nil)
		n.Right = o.expr(psess, n.Right, nil)
		t := n.Left.Type
		if t.IsStruct() || t.IsArray() {

			n.Left = o.addrTemp(psess, n.Left)
			n.Right = o.addrTemp(psess, n.Right)
		}
	}
	psess.
		lineno = lno
	return n
}

// okas creates and returns an assignment of val to ok,
// including an explicit conversion if necessary.
func (psess *PackageSession) okas(ok, val *Node) *Node {
	if !ok.isBlank() {
		val = psess.conv(val, ok.Type)
	}
	return psess.nod(OAS, ok, val)
}

// as2 orders OAS2XXXX nodes. It creates temporaries to ensure left-to-right assignment.
// The caller should order the right-hand side of the assignment before calling orderas2.
// It rewrites,
// 	a, b, a = ...
// as
//	tmp1, tmp2, tmp3 = ...
// 	a, b, a = tmp1, tmp2, tmp3
// This is necessary to ensure left to right assignment order.
func (o *Order) as2(psess *PackageSession, n *Node) {
	tmplist := []*Node{}
	left := []*Node{}
	for _, l := range n.List.Slice() {
		if !l.isBlank() {
			tmp := o.newTemp(psess, l.Type, psess.types.Haspointers(l.Type))
			tmplist = append(tmplist, tmp)
			left = append(left, l)
		}
	}

	o.out = append(o.out, n)

	as := psess.nod(OAS2, nil, nil)
	as.List.Set(left)
	as.Rlist.Set(tmplist)
	as = psess.typecheck(as, Etop)
	o.stmt(psess, as)

	ti := 0
	for ni, l := range n.List.Slice() {
		if !l.isBlank() {
			n.List.SetIndex(ni, tmplist[ti])
			ti++
		}
	}
}

// okAs2 orders OAS2 with ok.
// Just like as2, this also adds temporaries to ensure left-to-right assignment.
func (o *Order) okAs2(psess *PackageSession, n *Node) {
	var tmp1, tmp2 *Node
	if !n.List.First().isBlank() {
		typ := n.Rlist.First().Type
		tmp1 = o.newTemp(psess, typ, psess.types.Haspointers(typ))
	}

	if !n.List.Second().isBlank() {
		tmp2 = o.newTemp(psess, psess.types.Types[TBOOL], false)
	}

	o.out = append(o.out, n)

	if tmp1 != nil {
		r := psess.nod(OAS, n.List.First(), tmp1)
		r = psess.typecheck(r, Etop)
		o.mapAssign(psess, r)
		n.List.SetFirst(tmp1)
	}
	if tmp2 != nil {
		r := psess.okas(n.List.Second(), tmp2)
		r = psess.typecheck(r, Etop)
		o.mapAssign(psess, r)
		n.List.SetSecond(tmp2)
	}
}

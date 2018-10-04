// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/src"
)

// Rewrite tree to use separate statements to enforce
// order of evaluation. Makes walk easier, because it
// can (after this runs) reorder at will within an expression.
//
// Rewrite m[k] op= r into m[k] = m[k] op r if op is / or %.
//
// Introduce temporaries as needed by runtime routines.
// For example, the map runtime routines take the map key
// by reference, so make sure all map keys are addressable
// by copying them to temporaries as needed.
// The same is true for channel operations.
//
// Arrange that map index expressions only appear in direct
// assignments x = m[k] or m[k] = x, never in larger expressions.
//
// Arrange that receive expressions only appear in direct assignments
// x = <-c or as standalone statements <-c, never in larger expressions.

// TODO(rsc): The temporary introduction during multiple assignments
// should be moved into this file, so that the temporaries can be cleaned
// and so that conversions implicit in the OAS2FUNC and OAS2RECV
// nodes can be made explicit and then have their temporaries cleaned.

// TODO(rsc): Goto and multilevel break/continue can jump over
// inserted VARKILL annotations. Work out a way to handle these.
// The current implementation is safe, in that it will execute correctly.
// But it won't reuse temporaries as aggressively as it might, and
// it can result in unnecessary zeroing of those variables in the function
// prologue.

// Order holds state during the ordering process.
type Order struct {
	out  []*Node // list of generated statements
	temp []*Node // stack of temporary variables
}

// Order rewrites fn.Nbody to apply the ordering constraints
// described in the comment at the top of the file.
func (pstate *PackageState) order(fn *Node) {
	if pstate.Debug['W'] > 1 {
		s := fmt.Sprintf("\nbefore order %v", fn.Func.Nname.Sym)
		dumplist(s, fn.Nbody)
	}

	pstate.orderBlock(&fn.Nbody)
}

// newTemp allocates a new temporary with the given type,
// pushes it onto the temp stack, and returns it.
// If clear is true, newTemp emits code to zero the temporary.
func (o *Order) newTemp(pstate *PackageState, t *types.Type, clear bool) *Node {
	v := pstate.temp(t)
	if clear {
		a := pstate.nod(OAS, v, nil)
		a = pstate.typecheck(a, Etop)
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
func (o *Order) copyExpr(pstate *PackageState, n *Node, t *types.Type, clear bool) *Node {
	v := o.newTemp(pstate, t, clear)
	a := pstate.nod(OAS, v, n)
	a = pstate.typecheck(a, Etop)
	o.out = append(o.out, a)
	return v
}

// cheapExpr returns a cheap version of n.
// The definition of cheap is that n is a variable or constant.
// If not, cheapExpr allocates a new tmp, emits tmp = n,
// and then returns tmp.
func (o *Order) cheapExpr(pstate *PackageState, n *Node) *Node {
	if n == nil {
		return nil
	}

	switch n.Op {
	case ONAME, OLITERAL:
		return n
	case OLEN, OCAP:
		l := o.cheapExpr(pstate, n.Left)
		if l == n.Left {
			return n
		}
		a := n.copy()
		a.Orig = a
		a.Left = l
		return pstate.typecheck(a, Erv)
	}

	return o.copyExpr(pstate, n, n.Type, false)
}

// safeExpr returns a safe version of n.
// The definition of safe is that n can appear multiple times
// without violating the semantics of the original program,
// and that assigning to the safe version has the same effect
// as assigning to the original n.
//
// The intended use is to apply to x when rewriting x += y into x = x + y.
func (o *Order) safeExpr(pstate *PackageState, n *Node) *Node {
	switch n.Op {
	case ONAME, OLITERAL:
		return n

	case ODOT, OLEN, OCAP:
		l := o.safeExpr(pstate, n.Left)
		if l == n.Left {
			return n
		}
		a := n.copy()
		a.Orig = a
		a.Left = l
		return pstate.typecheck(a, Erv)

	case ODOTPTR, OIND:
		l := o.cheapExpr(pstate, n.Left)
		if l == n.Left {
			return n
		}
		a := n.copy()
		a.Orig = a
		a.Left = l
		return pstate.typecheck(a, Erv)

	case OINDEX, OINDEXMAP:
		var l *Node
		if n.Left.Type.IsArray() {
			l = o.safeExpr(pstate, n.Left)
		} else {
			l = o.cheapExpr(pstate, n.Left)
		}
		r := o.cheapExpr(pstate, n.Right)
		if l == n.Left && r == n.Right {
			return n
		}
		a := n.copy()
		a.Orig = a
		a.Left = l
		a.Right = r
		return pstate.typecheck(a, Erv)

	default:
		pstate.Fatalf("ordersafeexpr %v", n.Op)
		return nil // not reached
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
func (o *Order) addrTemp(pstate *PackageState, n *Node) *Node {
	if pstate.consttype(n) > 0 {
		// TODO: expand this to all static composite literal nodes?
		n = pstate.defaultlit(n, nil)
		pstate.dowidth(n.Type)
		vstat := pstate.staticname(n.Type)
		vstat.Name.SetReadonly(true)
		var out []*Node
		pstate.staticassign(vstat, n, &out)
		if out != nil {
			pstate.Fatalf("staticassign of const generated code: %+v", n)
		}
		vstat = pstate.typecheck(vstat, Erv)
		return vstat
	}
	if isaddrokay(n) {
		return n
	}
	return o.copyExpr(pstate, n, n.Type, false)
}

// mapKeyTemp prepares n to be a key in a map runtime call and returns n.
// It should only be used for map runtime calls which have *_fast* versions.
func (o *Order) mapKeyTemp(pstate *PackageState, t *types.Type, n *Node) *Node {
	// Most map calls need to take the address of the key.
	// Exception: map*_fast* calls. See golang.org/issue/19015.
	if pstate.mapfast(t) == mapslow {
		return o.addrTemp(pstate, n)
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
func (o *Order) cleanTempNoPop(pstate *PackageState, mark ordermarker) []*Node {
	var out []*Node
	for i := len(o.temp) - 1; i >= int(mark); i-- {
		n := o.temp[i]
		if n.Name.Keepalive() {
			n.Name.SetKeepalive(false)
			n.SetAddrtaken(true) // ensure SSA keeps the n variable
			live := pstate.nod(OVARLIVE, n, nil)
			live = pstate.typecheck(live, Etop)
			out = append(out, live)
		}
		kill := pstate.nod(OVARKILL, n, nil)
		kill = pstate.typecheck(kill, Etop)
		out = append(out, kill)
	}
	return out
}

// cleanTemp emits VARKILL instructions for each temporary above the
// mark on the temporary stack and removes them from the stack.
func (o *Order) cleanTemp(pstate *PackageState, top ordermarker) {
	o.out = append(o.out, o.cleanTempNoPop(pstate, top)...)
	o.popTemp(top)
}

// stmtList orders each of the statements in the list.
func (o *Order) stmtList(pstate *PackageState, l Nodes) {
	for _, n := range l.Slice() {
		o.stmt(pstate, n)
	}
}

// orderBlock orders the block of statements in n into a new slice,
// and then replaces the old slice in n with the new slice.
func (pstate *PackageState) orderBlock(n *Nodes) {
	var order Order
	mark := order.markTemp()
	order.stmtList(pstate, *n)
	order.cleanTemp(pstate, mark)
	n.Set(order.out)
}

// exprInPlace orders the side effects in *np and
// leaves them as the init list of the final *np.
// The result of exprInPlace MUST be assigned back to n, e.g.
// 	n.Left = o.exprInPlace(n.Left)
func (o *Order) exprInPlace(pstate *PackageState, n *Node) *Node {
	var order Order
	n = order.expr(pstate, n, nil)
	n = pstate.addinit(n, order.out)

	// insert new temporaries from order
	// at head of outer list.
	o.temp = append(o.temp, order.temp...)
	return n
}

// orderStmtInPlace orders the side effects of the single statement *np
// and replaces it with the resulting statement list.
// The result of orderStmtInPlace MUST be assigned back to n, e.g.
// 	n.Left = orderStmtInPlace(n.Left)
func (pstate *PackageState) orderStmtInPlace(n *Node) *Node {
	var order Order
	mark := order.markTemp()
	order.stmt(pstate, n)
	order.cleanTemp(pstate, mark)
	return pstate.liststmt(order.out)
}

// init moves n's init list to o.out.
func (o *Order) init(pstate *PackageState, n *Node) {
	if n.mayBeShared() {
		// For concurrency safety, don't mutate potentially shared nodes.
		// First, ensure that no work is required here.
		if n.Ninit.Len() > 0 {
			pstate.Fatalf("orderinit shared node with ninit")
		}
		return
	}
	o.stmtList(pstate, n.Ninit)
	n.Ninit.Set(nil)
}

// Ismulticall reports whether the list l is f() for a multi-value function.
// Such an f() could appear as the lone argument to a multi-arg function.
func (pstate *PackageState) ismulticall(l Nodes) bool {
	// one arg only
	if l.Len() != 1 {
		return false
	}
	n := l.First()

	// must be call
	switch n.Op {
	default:
		return false
	case OCALLFUNC, OCALLMETH, OCALLINTER:
		// call must return multiple values
		return n.Left.Type.NumResults(pstate.types) > 1
	}
}

// copyRet emits t1, t2, ... = n, where n is a function call,
// and then returns the list t1, t2, ....
func (o *Order) copyRet(pstate *PackageState, n *Node) []*Node {
	if !n.Type.IsFuncArgStruct() {
		pstate.Fatalf("copyret %v %d", n.Type, n.Left.Type.NumResults(pstate.types))
	}

	var l1, l2 []*Node
	for _, f := range n.Type.Fields(pstate.types).Slice() {
		tmp := pstate.temp(f.Type)
		l1 = append(l1, tmp)
		l2 = append(l2, tmp)
	}

	as := pstate.nod(OAS2, nil, nil)
	as.List.Set(l1)
	as.Rlist.Set1(n)
	as = pstate.typecheck(as, Etop)
	o.stmt(pstate, as)

	return l2
}

// callArgs orders the list of call arguments *l.
func (o *Order) callArgs(pstate *PackageState, l *Nodes) {
	if pstate.ismulticall(*l) {
		// return f() where f() is multiple values.
		l.Set(o.copyRet(pstate, l.First()))
	} else {
		o.exprList(pstate, *l)
	}
}

// call orders the call expression n.
// n.Op is OCALLMETH/OCALLFUNC/OCALLINTER or a builtin like OCOPY.
func (o *Order) call(pstate *PackageState, n *Node) {
	n.Left = o.expr(pstate, n.Left, nil)
	n.Right = o.expr(pstate, n.Right, nil) // ODDDARG temp
	o.callArgs(pstate, &n.List)

	if n.Op != OCALLFUNC {
		return
	}
	keepAlive := func(i int) {
		// If the argument is really a pointer being converted to uintptr,
		// arrange for the pointer to be kept alive until the call returns,
		// by copying it into a temp and marking that temp
		// still alive when we pop the temp stack.
		xp := n.List.Addr(i)
		for (*xp).Op == OCONVNOP && !(*xp).Type.IsUnsafePtr() {
			xp = &(*xp).Left
		}
		x := *xp
		if x.Type.IsUnsafePtr() {
			x = o.copyExpr(pstate, x, x.Type, false)
			x.Name.SetKeepalive(true)
			*xp = x
		}
	}

	for i, t := range n.Left.Type.Params(pstate.types).FieldSlice(pstate.types) {
		// Check for "unsafe-uintptr" tag provided by escape analysis.
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
func (o *Order) mapAssign(pstate *PackageState, n *Node) {
	switch n.Op {
	default:
		pstate.Fatalf("ordermapassign %v", n.Op)

	case OAS, OASOP:
		if n.Left.Op == OINDEXMAP {
			// Make sure we evaluate the RHS before starting the map insert.
			// We need to make sure the RHS won't panic.  See issue 22881.
			if n.Right.Op == OAPPEND {
				s := n.Right.List.Slice()[1:]
				for i, n := range s {
					s[i] = o.cheapExpr(pstate, n)
				}
			} else {
				n.Right = o.cheapExpr(pstate, n.Right)
			}
		}
		o.out = append(o.out, n)

	case OAS2, OAS2DOTTYPE, OAS2MAPR, OAS2FUNC:
		var post []*Node
		for i, m := range n.List.Slice() {
			switch {
			case m.Op == OINDEXMAP:
				if !m.Left.IsAutoTmp() {
					m.Left = o.copyExpr(pstate, m.Left, m.Left.Type, false)
				}
				if !m.Right.IsAutoTmp() {
					m.Right = o.copyExpr(pstate, m.Right, m.Right.Type, false)
				}
				fallthrough
			case pstate.instrumenting && n.Op == OAS2FUNC && !m.isBlank():
				t := o.newTemp(pstate, m.Type, false)
				n.List.SetIndex(i, t)
				a := pstate.nod(OAS, m, t)
				a = pstate.typecheck(a, Etop)
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
func (o *Order) stmt(pstate *PackageState, n *Node) {
	if n == nil {
		return
	}

	lno := pstate.setlineno(n)
	o.init(pstate, n)

	switch n.Op {
	default:
		pstate.Fatalf("orderstmt %v", n.Op)

	case OVARKILL, OVARLIVE:
		o.out = append(o.out, n)

	case OAS:
		t := o.markTemp()
		n.Left = o.expr(pstate, n.Left, nil)
		n.Right = o.expr(pstate, n.Right, n.Left)
		o.mapAssign(pstate, n)
		o.cleanTemp(pstate, t)

	case OAS2,
		OCLOSE,
		OCOPY,
		OPRINT,
		OPRINTN,
		ORECOVER,
		ORECV:
		t := o.markTemp()
		n.Left = o.expr(pstate, n.Left, nil)
		n.Right = o.expr(pstate, n.Right, nil)
		o.exprList(pstate, n.List)
		o.exprList(pstate, n.Rlist)
		switch n.Op {
		case OAS2:
			o.mapAssign(pstate, n)
		default:
			o.out = append(o.out, n)
		}
		o.cleanTemp(pstate, t)

	case OASOP:
		t := o.markTemp()
		n.Left = o.expr(pstate, n.Left, nil)
		n.Right = o.expr(pstate, n.Right, nil)

		if pstate.instrumenting || n.Left.Op == OINDEXMAP && (n.SubOp(pstate) == ODIV || n.SubOp(pstate) == OMOD) {
			// Rewrite m[k] op= r into m[k] = m[k] op r so
			// that we can ensure that if op panics
			// because r is zero, the panic happens before
			// the map assignment.

			n.Left = o.safeExpr(pstate, n.Left)

			l := pstate.treecopy(n.Left, pstate.src.NoXPos)
			if l.Op == OINDEXMAP {
				l.SetIndexMapLValue(pstate, false)
			}
			l = o.copyExpr(pstate, l, n.Left.Type, false)
			n.Right = pstate.nod(n.SubOp(pstate), l, n.Right)
			n.Right = pstate.typecheck(n.Right, Erv)
			n.Right = o.expr(pstate, n.Right, nil)

			n.Op = OAS
			n.ResetAux()
		}

		o.mapAssign(pstate, n)
		o.cleanTemp(pstate, t)

	// Special: make sure key is addressable if needed,
	// and make sure OINDEXMAP is not copied out.
	case OAS2MAPR:
		t := o.markTemp()
		o.exprList(pstate, n.List)
		r := n.Rlist.First()
		r.Left = o.expr(pstate, r.Left, nil)
		r.Right = o.expr(pstate, r.Right, nil)

		// See case OINDEXMAP below.
		if r.Right.Op == OARRAYBYTESTR {
			r.Right.Op = OARRAYBYTESTRTMP
		}
		r.Right = o.mapKeyTemp(pstate, r.Left.Type, r.Right)
		o.okAs2(pstate, n)
		o.cleanTemp(pstate, t)

	// Special: avoid copy of func call n.Rlist.First().
	case OAS2FUNC:
		t := o.markTemp()
		o.exprList(pstate, n.List)
		o.call(pstate, n.Rlist.First())
		o.as2(pstate, n)
		o.cleanTemp(pstate, t)

	// Special: use temporary variables to hold result,
	// so that assertI2Tetc can take address of temporary.
	// No temporary for blank assignment.
	case OAS2DOTTYPE:
		t := o.markTemp()
		o.exprList(pstate, n.List)
		n.Rlist.First().Left = o.expr(pstate, n.Rlist.First().Left, nil) // i in i.(T)
		o.okAs2(pstate, n)
		o.cleanTemp(pstate, t)

	// Special: use temporary variables to hold result,
	// so that chanrecv can take address of temporary.
	case OAS2RECV:
		t := o.markTemp()
		o.exprList(pstate, n.List)
		n.Rlist.First().Left = o.expr(pstate, n.Rlist.First().Left, nil) // arg to recv
		ch := n.Rlist.First().Left.Type
		tmp1 := o.newTemp(pstate, ch.Elem(pstate.types), pstate.types.Haspointers(ch.Elem(pstate.types)))
		tmp2 := o.newTemp(pstate, pstate.types.Types[TBOOL], false)
		o.out = append(o.out, n)
		r := pstate.nod(OAS, n.List.First(), tmp1)
		r = pstate.typecheck(r, Etop)
		o.mapAssign(pstate, r)
		r = pstate.okas(n.List.Second(), tmp2)
		r = pstate.typecheck(r, Etop)
		o.mapAssign(pstate, r)
		n.List.Set2(tmp1, tmp2)
		o.cleanTemp(pstate, t)

	// Special: does not save n onto out.
	case OBLOCK, OEMPTY:
		o.stmtList(pstate, n.List)

	// Special: n->left is not an expression; save as is.
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

	// Special: handle call arguments.
	case OCALLFUNC, OCALLINTER, OCALLMETH:
		t := o.markTemp()
		o.call(pstate, n)
		o.out = append(o.out, n)
		o.cleanTemp(pstate, t)

	// Special: order arguments to inner call but not call itself.
	case ODEFER, OPROC:
		t := o.markTemp()
		o.call(pstate, n.Left)
		o.out = append(o.out, n)
		o.cleanTemp(pstate, t)

	case ODELETE:
		t := o.markTemp()
		n.List.SetFirst(o.expr(pstate, n.List.First(), nil))
		n.List.SetSecond(o.expr(pstate, n.List.Second(), nil))
		n.List.SetSecond(o.mapKeyTemp(pstate, n.List.First().Type, n.List.Second()))
		o.out = append(o.out, n)
		o.cleanTemp(pstate, t)

	// Clean temporaries from condition evaluation at
	// beginning of loop body and after for statement.
	case OFOR:
		t := o.markTemp()
		n.Left = o.exprInPlace(pstate, n.Left)
		n.Nbody.Prepend(o.cleanTempNoPop(pstate, t)...)
		pstate.orderBlock(&n.Nbody)
		n.Right = pstate.orderStmtInPlace(n.Right)
		o.out = append(o.out, n)
		o.cleanTemp(pstate, t)

	// Clean temporaries from condition at
	// beginning of both branches.
	case OIF:
		t := o.markTemp()
		n.Left = o.exprInPlace(pstate, n.Left)
		n.Nbody.Prepend(o.cleanTempNoPop(pstate, t)...)
		n.Rlist.Prepend(o.cleanTempNoPop(pstate, t)...)
		o.popTemp(t)
		pstate.orderBlock(&n.Nbody)
		pstate.orderBlock(&n.Rlist)
		o.out = append(o.out, n)

	// Special: argument will be converted to interface using convT2E
	// so make sure it is an addressable temporary.
	case OPANIC:
		t := o.markTemp()
		n.Left = o.expr(pstate, n.Left, nil)
		if !n.Left.Type.IsInterface() {
			n.Left = o.addrTemp(pstate, n.Left)
		}
		o.out = append(o.out, n)
		o.cleanTemp(pstate, t)

	case ORANGE:
		// n.Right is the expression being ranged over.
		// order it, and then make a copy if we need one.
		// We almost always do, to ensure that we don't
		// see any value changes made during the loop.
		// Usually the copy is cheap (e.g., array pointer,
		// chan, slice, string are all tiny).
		// The exception is ranging over an array value
		// (not a slice, not a pointer to array),
		// which must make a copy to avoid seeing updates made during
		// the range body. Ranging over an array value is uncommon though.

		// Mark []byte(str) range expression to reuse string backing storage.
		// It is safe because the storage cannot be mutated.
		if n.Right.Op == OSTRARRAYBYTE {
			n.Right.Op = OSTRARRAYBYTETMP
		}

		t := o.markTemp()
		n.Right = o.expr(pstate, n.Right, nil)

		orderBody := true
		switch n.Type.Etype {
		default:
			pstate.Fatalf("orderstmt range %v", n.Type)

		case TARRAY, TSLICE:
			if n.List.Len() < 2 || n.List.Second().isBlank() {
				// for i := range x will only use x once, to compute len(x).
				// No need to copy it.
				break
			}
			fallthrough

		case TCHAN, TSTRING:
			// chan, string, slice, array ranges use value multiple times.
			// make copy.
			r := n.Right

			if r.Type.IsString() && r.Type != pstate.types.Types[TSTRING] {
				r = pstate.nod(OCONV, r, nil)
				r.Type = pstate.types.Types[TSTRING]
				r = pstate.typecheck(r, Erv)
			}

			n.Right = o.copyExpr(pstate, r, r.Type, false)

		case TMAP:
			if pstate.isMapClear(n) {
				// Preserve the body of the map clear pattern so it can
				// be detected during walk. The loop body will not be used
				// when optimizing away the range loop to a runtime call.
				orderBody = false
				break
			}

			// copy the map value in case it is a map literal.
			// TODO(rsc): Make tmp = literal expressions reuse tmp.
			// For maps tmp is just one word so it hardly matters.
			r := n.Right
			n.Right = o.copyExpr(pstate, r, r.Type, false)

			// prealloc[n] is the temp for the iterator.
			// hiter contains pointers and needs to be zeroed.
			pstate.prealloc[n] = o.newTemp(pstate, pstate.hiter(n.Type), true)
		}
		o.exprListInPlace(pstate, n.List)
		if orderBody {
			pstate.orderBlock(&n.Nbody)
		}
		o.out = append(o.out, n)
		o.cleanTemp(pstate, t)

	case ORETURN:
		o.callArgs(pstate, &n.List)
		o.out = append(o.out, n)

	// Special: clean case temporaries in each block entry.
	// Select must enter one of its blocks, so there is no
	// need for a cleaning at the end.
	// Doubly special: evaluation order for select is stricter
	// than ordinary expressions. Even something like p.c
	// has to be hoisted into a temporary, so that it cannot be
	// reordered after the channel evaluation for a different
	// case (if p were nil, then the timing of the fault would
	// give this away).
	case OSELECT:
		t := o.markTemp()

		for _, n2 := range n.List.Slice() {
			if n2.Op != OXCASE {
				pstate.Fatalf("order select case %v", n2.Op)
			}
			r := n2.Left
			pstate.setlineno(n2)

			// Append any new body prologue to ninit.
			// The next loop will insert ninit into nbody.
			if n2.Ninit.Len() != 0 {
				pstate.Fatalf("order select ninit")
			}
			if r == nil {
				continue
			}
			switch r.Op {
			default:
				Dump("select case", r)
				pstate.Fatalf("unknown op in select %v", r.Op)

			// If this is case x := <-ch or case x, y := <-ch, the case has
			// the ODCL nodes to declare x and y. We want to delay that
			// declaration (and possible allocation) until inside the case body.
			// Delete the ODCL nodes here and recreate them inside the body below.
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
					pstate.Fatalf("ninit on select recv")
				}

				// case x = <-c
				// case x, ok = <-c
				// r->left is x, r->ntest is ok, r->right is ORECV, r->right->left is c.
				// r->left == N means 'case <-c'.
				// c is always evaluated; x and ok are only evaluated when assigned.
				r.Right.Left = o.expr(pstate, r.Right.Left, nil)

				if r.Right.Left.Op != ONAME {
					r.Right.Left = o.copyExpr(pstate, r.Right.Left, r.Right.Left.Type, false)
				}

				// Introduce temporary for receive and move actual copy into case body.
				// avoids problems with target being addressed, as usual.
				// NOTE: If we wanted to be clever, we could arrange for just one
				// temporary per distinct type, sharing the temp among all receives
				// with that temp. Similarly one ok bool could be shared among all
				// the x,ok receives. Not worth doing until there's a clear need.
				if r.Left != nil && r.Left.isBlank() {
					r.Left = nil
				}
				if r.Left != nil {
					// use channel element type for temporary to avoid conversions,
					// such as in case interfacevalue = <-intchan.
					// the conversion happens in the OAS instead.
					tmp1 := r.Left

					if r.Colas() {
						tmp2 := pstate.nod(ODCL, tmp1, nil)
						tmp2 = pstate.typecheck(tmp2, Etop)
						n2.Ninit.Append(tmp2)
					}

					r.Left = o.newTemp(pstate, r.Right.Left.Type.Elem(pstate.types), pstate.types.Haspointers(r.Right.Left.Type.Elem(pstate.types)))
					tmp2 := pstate.nod(OAS, tmp1, r.Left)
					tmp2 = pstate.typecheck(tmp2, Etop)
					n2.Ninit.Append(tmp2)
				}

				if r.List.Len() != 0 && r.List.First().isBlank() {
					r.List.Set(nil)
				}
				if r.List.Len() != 0 {
					tmp1 := r.List.First()
					if r.Colas() {
						tmp2 := pstate.nod(ODCL, tmp1, nil)
						tmp2 = pstate.typecheck(tmp2, Etop)
						n2.Ninit.Append(tmp2)
					}

					r.List.Set1(o.newTemp(pstate, pstate.types.Types[TBOOL], false))
					tmp2 := pstate.okas(tmp1, r.List.First())
					tmp2 = pstate.typecheck(tmp2, Etop)
					n2.Ninit.Append(tmp2)
				}
				pstate.orderBlock(&n2.Ninit)

			case OSEND:
				if r.Ninit.Len() != 0 {
					dumplist("ninit", r.Ninit)
					pstate.Fatalf("ninit on select send")
				}

				// case c <- x
				// r->left is c, r->right is x, both are always evaluated.
				r.Left = o.expr(pstate, r.Left, nil)

				if !r.Left.IsAutoTmp() {
					r.Left = o.copyExpr(pstate, r.Left, r.Left.Type, false)
				}
				r.Right = o.expr(pstate, r.Right, nil)
				if !r.Right.IsAutoTmp() {
					r.Right = o.copyExpr(pstate, r.Right, r.Right.Type, false)
				}
			}
		}
		// Now that we have accumulated all the temporaries, clean them.
		// Also insert any ninit queued during the previous loop.
		// (The temporary cleaning must follow that ninit work.)
		for _, n3 := range n.List.Slice() {
			pstate.orderBlock(&n3.Nbody)
			n3.Nbody.Prepend(o.cleanTempNoPop(pstate, t)...)

			// TODO(mdempsky): Is this actually necessary?
			// walkselect appears to walk Ninit.
			n3.Nbody.Prepend(n3.Ninit.Slice()...)
			n3.Ninit.Set(nil)
		}

		o.out = append(o.out, n)
		o.popTemp(t)

	// Special: value being sent is passed as a pointer; make it addressable.
	case OSEND:
		t := o.markTemp()
		n.Left = o.expr(pstate, n.Left, nil)
		n.Right = o.expr(pstate, n.Right, nil)
		if pstate.instrumenting {
			// Force copying to the stack so that (chan T)(nil) <- x
			// is still instrumented as a read of x.
			n.Right = o.copyExpr(pstate, n.Right, n.Right.Type, false)
		} else {
			n.Right = o.addrTemp(pstate, n.Right)
		}
		o.out = append(o.out, n)
		o.cleanTemp(pstate, t)

	// TODO(rsc): Clean temporaries more aggressively.
	// Note that because walkswitch will rewrite some of the
	// switch into a binary search, this is not as easy as it looks.
	// (If we ran that code here we could invoke orderstmt on
	// the if-else chain instead.)
	// For now just clean all the temporaries at the end.
	// In practice that's fine.
	case OSWITCH:
		t := o.markTemp()
		n.Left = o.expr(pstate, n.Left, nil)
		for _, ncas := range n.List.Slice() {
			if ncas.Op != OXCASE {
				pstate.Fatalf("order switch case %v", ncas.Op)
			}
			o.exprListInPlace(pstate, ncas.List)
			pstate.orderBlock(&ncas.Nbody)
		}

		o.out = append(o.out, n)
		o.cleanTemp(pstate, t)
	}

	pstate.lineno = lno
}

// exprList orders the expression list l into o.
func (o *Order) exprList(pstate *PackageState, l Nodes) {
	s := l.Slice()
	for i := range s {
		s[i] = o.expr(pstate, s[i], nil)
	}
}

// exprListInPlace orders the expression list l but saves
// the side effects on the individual expression ninit lists.
func (o *Order) exprListInPlace(pstate *PackageState, l Nodes) {
	s := l.Slice()
	for i := range s {
		s[i] = o.exprInPlace(pstate, s[i])
	}
}

// expr orders a single expression, appending side
// effects to o.out as needed.
// If this is part of an assignment lhs = *np, lhs is given.
// Otherwise lhs == nil. (When lhs != nil it may be possible
// to avoid copying the result of the expression to a temporary.)
// The result of expr MUST be assigned back to n, e.g.
// 	n.Left = o.expr(n.Left, lhs)
func (o *Order) expr(pstate *PackageState, n, lhs *Node) *Node {
	if n == nil {
		return n
	}

	lno := pstate.setlineno(n)
	o.init(pstate, n)

	switch n.Op {
	default:
		n.Left = o.expr(pstate, n.Left, nil)
		n.Right = o.expr(pstate, n.Right, nil)
		o.exprList(pstate, n.List)
		o.exprList(pstate, n.Rlist)

	// Addition of strings turns into a function call.
	// Allocate a temporary to hold the strings.
	// Fewer than 5 strings use direct runtime helpers.
	case OADDSTR:
		o.exprList(pstate, n.List)

		if n.List.Len() > 5 {
			t := pstate.types.NewArray(pstate.types.Types[TSTRING], int64(n.List.Len()))
			pstate.prealloc[n] = o.newTemp(pstate, t, false)
		}

		// Mark string(byteSlice) arguments to reuse byteSlice backing
		// buffer during conversion. String concatenation does not
		// memorize the strings for later use, so it is safe.
		// However, we can do it only if there is at least one non-empty string literal.
		// Otherwise if all other arguments are empty strings,
		// concatstrings will return the reference to the temp string
		// to the caller.
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
		n.Left = o.expr(pstate, n.Left, nil)
		n.Right = o.expr(pstate, n.Right, nil)

		// Mark string(byteSlice) arguments to reuse byteSlice backing
		// buffer during conversion. String comparison does not
		// memorize the strings for later use, so it is safe.
		if n.Left.Op == OARRAYBYTESTR {
			n.Left.Op = OARRAYBYTESTRTMP
		}
		if n.Right.Op == OARRAYBYTESTR {
			n.Right.Op = OARRAYBYTESTRTMP
		}

	// key must be addressable
	case OINDEXMAP:
		n.Left = o.expr(pstate, n.Left, nil)
		n.Right = o.expr(pstate, n.Right, nil)
		needCopy := false

		if !n.IndexMapLValue(pstate) && pstate.instrumenting {
			// Race detector needs the copy so it can
			// call treecopy on the result.
			needCopy = true
		}

		// For x = m[string(k)] where k is []byte, the allocation of
		// backing bytes for the string can be avoided by reusing
		// the []byte backing array. This is a special case that it
		// would be nice to handle more generally, but because
		// there are no []byte-keyed maps, this specific case comes
		// up in important cases in practice. See issue 3512.
		// Nothing can change the []byte we are not copying before
		// the map index, because the map access is going to
		// be forced to happen immediately following this
		// conversion (by the ordercopyexpr a few lines below).
		if !n.IndexMapLValue(pstate) && n.Right.Op == OARRAYBYTESTR {
			n.Right.Op = OARRAYBYTESTRTMP
			needCopy = true
		}

		n.Right = o.mapKeyTemp(pstate, n.Left.Type, n.Right)
		if needCopy {
			n = o.copyExpr(pstate, n, n.Type, false)
		}

	// concrete type (not interface) argument must be addressable
	// temporary to pass to runtime.
	case OCONVIFACE:
		n.Left = o.expr(pstate, n.Left, nil)

		if !n.Left.Type.IsInterface() {
			n.Left = o.addrTemp(pstate, n.Left)
		}

	case OCONVNOP:
		if n.Type.IsKind(TUNSAFEPTR) && n.Left.Type.IsKind(TUINTPTR) && (n.Left.Op == OCALLFUNC || n.Left.Op == OCALLINTER || n.Left.Op == OCALLMETH) {
			// When reordering unsafe.Pointer(f()) into a separate
			// statement, the conversion and function call must stay
			// together. See golang.org/issue/15329.
			o.init(pstate, n.Left)
			o.call(pstate, n.Left)
			if lhs == nil || lhs.Op != ONAME || pstate.instrumenting {
				n = o.copyExpr(pstate, n, n.Type, false)
			}
		} else {
			n.Left = o.expr(pstate, n.Left, nil)
		}

	case OANDAND, OOROR:
		mark := o.markTemp()
		n.Left = o.expr(pstate, n.Left, nil)

		// Clean temporaries from first branch at beginning of second.
		// Leave them on the stack so that they can be killed in the outer
		// context in case the short circuit is taken.
		n.Right = pstate.addinit(n.Right, o.cleanTempNoPop(pstate, mark))
		n.Right = o.exprInPlace(pstate, n.Right)

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

		if pstate.isRuneCount(n) {
			// len([]rune(s)) is rewritten to runtime.countrunes(s) later.
			n.Left.Left = o.expr(pstate, n.Left.Left, nil)
		} else {
			o.call(pstate, n)
		}

		if lhs == nil || lhs.Op != ONAME || pstate.instrumenting {
			n = o.copyExpr(pstate, n, n.Type, false)
		}

	case OAPPEND:
		// Check for append(x, make([]T, y)...) .
		if pstate.isAppendOfMake(n) {
			n.List.SetFirst(o.expr(pstate, n.List.First(), nil))             // order x
			n.List.Second().Left = o.expr(pstate, n.List.Second().Left, nil) // order y
		} else {
			o.callArgs(pstate, &n.List)
		}

		if lhs == nil || lhs.Op != ONAME && !pstate.samesafeexpr(lhs, n.List.First()) {
			n = o.copyExpr(pstate, n, n.Type, false)
		}

	case OSLICE, OSLICEARR, OSLICESTR, OSLICE3, OSLICE3ARR:
		n.Left = o.expr(pstate, n.Left, nil)
		low, high, max := n.SliceBounds(pstate)
		low = o.expr(pstate, low, nil)
		low = o.cheapExpr(pstate, low)
		high = o.expr(pstate, high, nil)
		high = o.cheapExpr(pstate, high)
		max = o.expr(pstate, max, nil)
		max = o.cheapExpr(pstate, max)
		n.SetSliceBounds(pstate, low, high, max)
		if lhs == nil || lhs.Op != ONAME && !pstate.samesafeexpr(lhs, n.Left) {
			n = o.copyExpr(pstate, n, n.Type, false)
		}

	case OCLOSURE:
		if n.Noescape() && n.Func.Closure.Func.Cvars.Len() > 0 {
			pstate.prealloc[n] = o.newTemp(pstate, pstate.types.Types[TUINT8], false) // walk will fill in correct type
		}

	case OARRAYLIT, OSLICELIT, OCALLPART:
		n.Left = o.expr(pstate, n.Left, nil)
		n.Right = o.expr(pstate, n.Right, nil)
		o.exprList(pstate, n.List)
		o.exprList(pstate, n.Rlist)
		if n.Noescape() {
			pstate.prealloc[n] = o.newTemp(pstate, pstate.types.Types[TUINT8], false) // walk will fill in correct type
		}

	case ODDDARG:
		if n.Noescape() {
			// The ddd argument does not live beyond the call it is created for.
			// Allocate a temporary that will be cleaned up when this statement
			// completes. We could be more aggressive and try to arrange for it
			// to be cleaned up when the call completes.
			pstate.prealloc[n] = o.newTemp(pstate, n.Type.Elem(pstate.types), false)
		}

	case ODOTTYPE, ODOTTYPE2:
		n.Left = o.expr(pstate, n.Left, nil)
		// TODO(rsc): The isfat is for consistency with componentgen and walkexpr.
		// It needs to be removed in all three places.
		// That would allow inlining x.(struct{*int}) the same as x.(*int).
		if !pstate.isdirectiface(n.Type) || isfat(n.Type) || pstate.instrumenting {
			n = o.copyExpr(pstate, n, n.Type, true)
		}

	case ORECV:
		n.Left = o.expr(pstate, n.Left, nil)
		n = o.copyExpr(pstate, n, n.Type, true)

	case OEQ, ONE:
		n.Left = o.expr(pstate, n.Left, nil)
		n.Right = o.expr(pstate, n.Right, nil)
		t := n.Left.Type
		if t.IsStruct() || t.IsArray() {
			// for complex comparisons, we need both args to be
			// addressable so we can pass them to the runtime.
			n.Left = o.addrTemp(pstate, n.Left)
			n.Right = o.addrTemp(pstate, n.Right)
		}
	}

	pstate.lineno = lno
	return n
}

// okas creates and returns an assignment of val to ok,
// including an explicit conversion if necessary.
func (pstate *PackageState) okas(ok, val *Node) *Node {
	if !ok.isBlank() {
		val = pstate.conv(val, ok.Type)
	}
	return pstate.nod(OAS, ok, val)
}

// as2 orders OAS2XXXX nodes. It creates temporaries to ensure left-to-right assignment.
// The caller should order the right-hand side of the assignment before calling orderas2.
// It rewrites,
// 	a, b, a = ...
// as
//	tmp1, tmp2, tmp3 = ...
// 	a, b, a = tmp1, tmp2, tmp3
// This is necessary to ensure left to right assignment order.
func (o *Order) as2(pstate *PackageState, n *Node) {
	tmplist := []*Node{}
	left := []*Node{}
	for _, l := range n.List.Slice() {
		if !l.isBlank() {
			tmp := o.newTemp(pstate, l.Type, pstate.types.Haspointers(l.Type))
			tmplist = append(tmplist, tmp)
			left = append(left, l)
		}
	}

	o.out = append(o.out, n)

	as := pstate.nod(OAS2, nil, nil)
	as.List.Set(left)
	as.Rlist.Set(tmplist)
	as = pstate.typecheck(as, Etop)
	o.stmt(pstate, as)

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
func (o *Order) okAs2(pstate *PackageState, n *Node) {
	var tmp1, tmp2 *Node
	if !n.List.First().isBlank() {
		typ := n.Rlist.First().Type
		tmp1 = o.newTemp(pstate, typ, pstate.types.Haspointers(typ))
	}

	if !n.List.Second().isBlank() {
		tmp2 = o.newTemp(pstate, pstate.types.Types[TBOOL], false)
	}

	o.out = append(o.out, n)

	if tmp1 != nil {
		r := pstate.nod(OAS, n.List.First(), tmp1)
		r = pstate.typecheck(r, Etop)
		o.mapAssign(pstate, r)
		n.List.SetFirst(tmp1)
	}
	if tmp2 != nil {
		r := pstate.okas(n.List.Second(), tmp2)
		r = pstate.typecheck(r, Etop)
		o.mapAssign(pstate, r)
		n.List.SetSecond(tmp2)
	}
}

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"math"
	"strings"
)

const (
	Etop      = 1 << iota // evaluated at statement level
	Erv                   // evaluated in value context
	Etype                 // evaluated in type context
	Ecall                 // call-only expressions are ok
	Efnstruct             // multivalue function returns are ok
	Easgn                 // assigning to expression
	Ecomplit              // type in composite literal
)

// resolve ONONAME to definition, if any.
func (pstate *PackageState) resolve(n *Node) *Node {
	if n == nil || n.Op != ONONAME {
		return n
	}

	if n.Sym.Pkg != pstate.localpkg {
		if pstate.inimport {
			pstate.Fatalf("recursive inimport")
		}
		pstate.inimport = true
		pstate.expandDecl(n)
		pstate.inimport = false
		return n
	}

	r := asNode(n.Sym.Def)
	if r == nil {
		return n
	}

	if r.Op == OIOTA {
		if i := len(pstate.typecheckdefstack); i > 0 {
			if x := pstate.typecheckdefstack[i-1]; x.Op == OLITERAL {
				return pstate.nodintconst(x.Iota())
			}
		}
		return n
	}

	return r
}

func (pstate *PackageState) typecheckslice(l []*Node, top int) {
	for i := range l {
		l[i] = pstate.typecheck(l[i], top)
	}
}

func (pstate *PackageState) typekind(t *types.Type) string {
	if t.IsSlice() {
		return "slice"
	}
	et := t.Etype
	if int(et) < len(pstate._typekind) {
		s := pstate._typekind[et]
		if s != "" {
			return s
		}
	}
	return fmt.Sprintf("etype=%d", et)
}

func (pstate *PackageState) cycleFor(start *Node) []*Node {
	// Find the start node in typecheck_tcstack.
	// We know that it must exist because each time we mark
	// a node with n.SetTypecheck(2) we push it on the stack,
	// and each time we mark a node with n.SetTypecheck(2) we
	// pop it from the stack. We hit a cycle when we encounter
	// a node marked 2 in which case is must be on the stack.
	i := len(pstate.typecheck_tcstack) - 1
	for i > 0 && pstate.typecheck_tcstack[i] != start {
		i--
	}

	// collect all nodes with same Op
	var cycle []*Node
	for _, n := range pstate.typecheck_tcstack[i:] {
		if n.Op == start.Op {
			cycle = append(cycle, n)
		}
	}

	return cycle
}

func (pstate *PackageState) cycleTrace(cycle []*Node) string {
	var s string
	for i, n := range cycle {
		s += fmt.Sprintf("\n\t%v: %v uses %v", n.Line(pstate), n, cycle[(i+1)%len(cycle)])
	}
	return s
}

// typecheck type checks node n.
// The result of typecheck MUST be assigned back to n, e.g.
// 	n.Left = typecheck(n.Left, top)
func (pstate *PackageState) typecheck(n *Node, top int) *Node {
	// cannot type check until all the source has been parsed
	if !pstate.typecheckok {
		pstate.Fatalf("early typecheck")
	}

	if n == nil {
		return nil
	}

	lno := pstate.setlineno(n)

	// Skip over parens.
	for n.Op == OPAREN {
		n = n.Left
	}

	// Resolve definition of name and value of iota lazily.
	n = pstate.resolve(n)

	// Skip typecheck if already done.
	// But re-typecheck ONAME/OTYPE/OLITERAL/OPACK node in case context has changed.
	if n.Typecheck() == 1 {
		switch n.Op {
		case ONAME, OTYPE, OLITERAL, OPACK:
			break

		default:
			pstate.lineno = lno
			return n
		}
	}

	if n.Typecheck() == 2 {
		// Typechecking loop. Trying printing a meaningful message,
		// otherwise a stack trace of typechecking.
		switch n.Op {
		// We can already diagnose variables used as types.
		case ONAME:
			if top&(Erv|Etype) == Etype {
				pstate.yyerror("%v is not a type", n)
			}

		case OTYPE:
			// Only report a type cycle if we are expecting a type.
			// Otherwise let other code report an error.
			if top&Etype == Etype {
				// A cycle containing only alias types is an error
				// since it would expand indefinitely when aliases
				// are substituted.
				cycle := pstate.cycleFor(n)
				for _, n := range cycle {
					if n.Name != nil && !n.Name.Param.Alias {
						pstate.lineno = lno
						return n
					}
				}
				pstate.yyerrorl(n.Pos, "invalid recursive type alias %v%s", n, pstate.cycleTrace(cycle))
			}

		case OLITERAL:
			if top&(Erv|Etype) == Etype {
				pstate.yyerror("%v is not a type", n)
				break
			}
			pstate.yyerrorl(n.Pos, "constant definition loop%s", pstate.cycleTrace(pstate.cycleFor(n)))
		}

		if pstate.nsavederrors+pstate.nerrors == 0 {
			var trace string
			for i := len(pstate.typecheck_tcstack) - 1; i >= 0; i-- {
				x := pstate.typecheck_tcstack[i]
				trace += fmt.Sprintf("\n\t%v %v", x.Line(pstate), x)
			}
			pstate.yyerror("typechecking loop involving %v%s", n, trace)
		}

		pstate.lineno = lno
		return n
	}

	n.SetTypecheck(2)

	pstate.typecheck_tcstack = append(pstate.typecheck_tcstack, n)
	n = pstate.typecheck1(n, top)

	n.SetTypecheck(1)

	last := len(pstate.typecheck_tcstack) - 1
	pstate.typecheck_tcstack[last] = nil
	pstate.typecheck_tcstack = pstate.typecheck_tcstack[:last]

	pstate.lineno = lno
	return n
}

// does n contain a call or receive operation?
func callrecv(n *Node) bool {
	if n == nil {
		return false
	}

	switch n.Op {
	case OCALL,
		OCALLMETH,
		OCALLINTER,
		OCALLFUNC,
		ORECV,
		OCAP,
		OLEN,
		OCOPY,
		ONEW,
		OAPPEND,
		ODELETE:
		return true
	}

	return callrecv(n.Left) || callrecv(n.Right) || callrecvlist(n.Ninit) || callrecvlist(n.Nbody) || callrecvlist(n.List) || callrecvlist(n.Rlist)
}

func callrecvlist(l Nodes) bool {
	for _, n := range l.Slice() {
		if callrecv(n) {
			return true
		}
	}
	return false
}

// indexlit implements typechecking of untyped values as
// array/slice indexes. It is almost equivalent to defaultlit
// but also accepts untyped numeric values representable as
// value of type int (see also checkmake for comparison).
// The result of indexlit MUST be assigned back to n, e.g.
// 	n.Left = indexlit(n.Left)
func (pstate *PackageState) indexlit(n *Node) *Node {
	if n != nil && n.Type != nil && n.Type.Etype == TIDEAL {
		return pstate.defaultlit(n, pstate.types.Types[TINT])
	}
	return n
}

// The result of typecheck1 MUST be assigned back to n, e.g.
// 	n.Left = typecheck1(n.Left, top)
func (pstate *PackageState) typecheck1(n *Node, top int) *Node {
	switch n.Op {
	case OXDOT, ODOT, ODOTPTR, ODOTMETH, ODOTINTER, ORETJMP:
	// n.Sym is a field/method name, not a variable.
	default:
		if n.Sym != nil {
			if n.Op == ONAME && n.SubOp(pstate) != 0 && top&Ecall == 0 {
				pstate.yyerror("use of builtin %v not in function call", n.Sym)
				n.Type = nil
				return n
			}

			pstate.typecheckdef(n)
			if n.Op == ONONAME {
				n.Type = nil
				return n
			}
		}
	}

	ok := 0
	switch n.Op {
	// until typecheck is complete, do nothing.
	default:
		Dump("typecheck", n)

		pstate.Fatalf("typecheck %v", n.Op)

	// names
	case OLITERAL:
		ok |= Erv

		if n.Type == nil && n.Val().Ctype(pstate) == CTSTR {
			n.Type = pstate.types.Idealstring
		}

	case ONONAME:
		ok |= Erv

	case ONAME:
		if n.Name.Decldepth == 0 {
			n.Name.Decldepth = pstate.decldepth
		}
		if n.SubOp(pstate) != 0 {
			ok |= Ecall
			break
		}

		if top&Easgn == 0 {
			// not a write to the variable
			if n.isBlank() {
				pstate.yyerror("cannot use _ as value")
				n.Type = nil
				return n
			}

			n.Name.SetUsed(true)
		}

		ok |= Erv

	case OPACK:
		pstate.yyerror("use of package %v without selector", n.Sym)
		n.Type = nil
		return n

	case ODDD:
		break

	// types (OIND is with exprs)
	case OTYPE:
		ok |= Etype

		if n.Type == nil {
			return n
		}

	case OTARRAY:
		ok |= Etype
		r := pstate.typecheck(n.Right, Etype)
		if r.Type == nil {
			n.Type = nil
			return n
		}

		var t *types.Type
		if n.Left == nil {
			t = pstate.types.NewSlice(r.Type)
		} else if n.Left.Op == ODDD {
			if top&Ecomplit == 0 {
				if !n.Diag() {
					n.SetDiag(true)
					pstate.yyerror("use of [...] array outside of array literal")
				}
				n.Type = nil
				return n
			}
			t = types.NewDDDArray(r.Type)
		} else {
			n.Left = pstate.indexlit(pstate.typecheck(n.Left, Erv))
			l := n.Left
			if pstate.consttype(l) != CTINT {
				switch {
				case l.Type == nil:
				// Error already reported elsewhere.
				case l.Type.IsInteger() && l.Op != OLITERAL:
					pstate.yyerror("non-constant array bound %v", l)
				default:
					pstate.yyerror("invalid array bound %v", l)
				}
				n.Type = nil
				return n
			}

			v := l.Val()
			if pstate.doesoverflow(v, pstate.types.Types[TINT]) {
				pstate.yyerror("array bound is too large")
				n.Type = nil
				return n
			}

			bound := v.U.(*Mpint).Int64(pstate)
			if bound < 0 {
				pstate.yyerror("array bound must be non-negative")
				n.Type = nil
				return n
			}
			t = pstate.types.NewArray(r.Type, bound)
		}

		n.Op = OTYPE
		n.Type = t
		n.Left = nil
		n.Right = nil
		if !t.IsDDDArray() {
			pstate.checkwidth(t)
		}

	case OTMAP:
		ok |= Etype
		n.Left = pstate.typecheck(n.Left, Etype)
		n.Right = pstate.typecheck(n.Right, Etype)
		l := n.Left
		r := n.Right
		if l.Type == nil || r.Type == nil {
			n.Type = nil
			return n
		}
		if l.Type.NotInHeap() {
			pstate.yyerror("go:notinheap map key not allowed")
		}
		if r.Type.NotInHeap() {
			pstate.yyerror("go:notinheap map value not allowed")
		}
		n.Op = OTYPE
		n.Type = pstate.types.NewMap(l.Type, r.Type)
		pstate.mapqueue = append(pstate.mapqueue, n) // check map keys when all types are settled
		n.Left = nil
		n.Right = nil

	case OTCHAN:
		ok |= Etype
		n.Left = pstate.typecheck(n.Left, Etype)
		l := n.Left
		if l.Type == nil {
			n.Type = nil
			return n
		}
		if l.Type.NotInHeap() {
			pstate.yyerror("chan of go:notinheap type not allowed")
		}
		t := pstate.types.NewChan(l.Type, n.TChanDir(pstate))
		n.Op = OTYPE
		n.Type = t
		n.Left = nil
		n.ResetAux()

	case OTSTRUCT:
		ok |= Etype
		n.Op = OTYPE
		n.Type = pstate.tostruct(n.List.Slice())
		if n.Type == nil || n.Type.Broke() {
			n.Type = nil
			return n
		}
		n.List.Set(nil)

	case OTINTER:
		ok |= Etype
		n.Op = OTYPE
		n.Type = pstate.tointerface(n.List.Slice())
		if n.Type == nil {
			return n
		}

	case OTFUNC:
		ok |= Etype
		n.Op = OTYPE
		n.Type = pstate.functype(n.Left, n.List.Slice(), n.Rlist.Slice())
		if n.Type == nil {
			return n
		}
		n.Left = nil
		n.List.Set(nil)
		n.Rlist.Set(nil)

	// type or expr
	case OIND:
		n.Left = pstate.typecheck(n.Left, Erv|Etype|top&Ecomplit)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if l.Op == OTYPE {
			ok |= Etype
			n.Op = OTYPE
			n.Type = pstate.types.NewPtr(l.Type)
			// Ensure l.Type gets dowidth'd for the backend. Issue 20174.
			// Don't checkwidth [...] arrays, though, since they
			// will be replaced by concrete-sized arrays. Issue 20333.
			if !l.Type.IsDDDArray() {
				pstate.checkwidth(l.Type)
			}
			n.Left = nil
			break
		}

		if !t.IsPtr() {
			if top&(Erv|Etop) != 0 {
				pstate.yyerror("invalid indirect of %L", n.Left)
				n.Type = nil
				return n
			}

			break
		}

		ok |= Erv
		n.Type = t.Elem(pstate.types)

	// arithmetic exprs
	case OASOP,
		OADD,
		OAND,
		OANDAND,
		OANDNOT,
		ODIV,
		OEQ,
		OGE,
		OGT,
		OLE,
		OLT,
		OLSH,
		ORSH,
		OMOD,
		OMUL,
		ONE,
		OOR,
		OOROR,
		OSUB,
		OXOR:
		var l *Node
		var op Op
		var r *Node
		if n.Op == OASOP {
			ok |= Etop
			n.Left = pstate.typecheck(n.Left, Erv)
			n.Right = pstate.typecheck(n.Right, Erv)
			l = n.Left
			r = n.Right
			pstate.checkassign(n, n.Left)
			if l.Type == nil || r.Type == nil {
				n.Type = nil
				return n
			}
			if n.Implicit() && !pstate.okforarith[l.Type.Etype] {
				pstate.yyerror("invalid operation: %v (non-numeric type %v)", n, l.Type)
				n.Type = nil
				return n
			}
			// TODO(marvin): Fix Node.EType type union.
			op = n.SubOp(pstate)
		} else {
			ok |= Erv
			n.Left = pstate.typecheck(n.Left, Erv)
			n.Right = pstate.typecheck(n.Right, Erv)
			l = n.Left
			r = n.Right
			if l.Type == nil || r.Type == nil {
				n.Type = nil
				return n
			}
			op = n.Op
		}
		if op == OLSH || op == ORSH {
			r = pstate.defaultlit(r, pstate.types.Types[TUINT])
			n.Right = r
			t := r.Type
			if !t.IsInteger() || t.IsSigned() {
				pstate.yyerror("invalid operation: %v (shift count type %v, must be unsigned integer)", n, r.Type)
				n.Type = nil
				return n
			}

			t = l.Type
			if t != nil && t.Etype != TIDEAL && !t.IsInteger() {
				pstate.yyerror("invalid operation: %v (shift of type %v)", n, t)
				n.Type = nil
				return n
			}

			// no defaultlit for left
			// the outer context gives the type
			n.Type = l.Type

			break
		}

		// ideal mixed with non-ideal
		l, r = pstate.defaultlit2(l, r, false)

		n.Left = l
		n.Right = r
		if l.Type == nil || r.Type == nil {
			n.Type = nil
			return n
		}
		t := l.Type
		if t.Etype == TIDEAL {
			t = r.Type
		}
		et := t.Etype
		if et == TIDEAL {
			et = TINT
		}
		aop := OXXX
		if pstate.iscmp[n.Op] && t.Etype != TIDEAL && !pstate.eqtype(l.Type, r.Type) {
			// comparison is okay as long as one side is
			// assignable to the other.  convert so they have
			// the same type.
			//
			// the only conversion that isn't a no-op is concrete == interface.
			// in that case, check comparability of the concrete type.
			// The conversion allocates, so only do it if the concrete type is huge.
			converted := false
			if r.Type.Etype != TBLANK {
				aop = pstate.assignop(l.Type, r.Type, nil)
				if aop != 0 {
					if r.Type.IsInterface() && !l.Type.IsInterface() && !pstate.IsComparable(l.Type) {
						pstate.yyerror("invalid operation: %v (operator %v not defined on %s)", n, op, pstate.typekind(l.Type))
						n.Type = nil
						return n
					}

					pstate.dowidth(l.Type)
					if r.Type.IsInterface() == l.Type.IsInterface() || l.Type.Width >= 1<<16 {
						l = pstate.nod(aop, l, nil)
						l.Type = r.Type
						l.SetTypecheck(1)
						n.Left = l
					}

					t = r.Type
					converted = true
				}
			}

			if !converted && l.Type.Etype != TBLANK {
				aop = pstate.assignop(r.Type, l.Type, nil)
				if aop != 0 {
					if l.Type.IsInterface() && !r.Type.IsInterface() && !pstate.IsComparable(r.Type) {
						pstate.yyerror("invalid operation: %v (operator %v not defined on %s)", n, op, pstate.typekind(r.Type))
						n.Type = nil
						return n
					}

					pstate.dowidth(r.Type)
					if r.Type.IsInterface() == l.Type.IsInterface() || r.Type.Width >= 1<<16 {
						r = pstate.nod(aop, r, nil)
						r.Type = l.Type
						r.SetTypecheck(1)
						n.Right = r
					}

					t = l.Type
				}
			}

			et = t.Etype
		}

		if t.Etype != TIDEAL && !pstate.eqtype(l.Type, r.Type) {
			l, r = pstate.defaultlit2(l, r, true)
			if r.Type.IsInterface() == l.Type.IsInterface() || aop == 0 {
				pstate.yyerror("invalid operation: %v (mismatched types %v and %v)", n, l.Type, r.Type)
				n.Type = nil
				return n
			}
		}

		if !pstate.okfor[op][et] {
			pstate.yyerror("invalid operation: %v (operator %v not defined on %s)", n, op, pstate.typekind(t))
			n.Type = nil
			return n
		}

		// okfor allows any array == array, map == map, func == func.
		// restrict to slice/map/func == nil and nil == slice/map/func.
		if l.Type.IsArray() && !pstate.IsComparable(l.Type) {
			pstate.yyerror("invalid operation: %v (%v cannot be compared)", n, l.Type)
			n.Type = nil
			return n
		}

		if l.Type.IsSlice() && !l.isNil(pstate) && !r.isNil(pstate) {
			pstate.yyerror("invalid operation: %v (slice can only be compared to nil)", n)
			n.Type = nil
			return n
		}

		if l.Type.IsMap() && !l.isNil(pstate) && !r.isNil(pstate) {
			pstate.yyerror("invalid operation: %v (map can only be compared to nil)", n)
			n.Type = nil
			return n
		}

		if l.Type.Etype == TFUNC && !l.isNil(pstate) && !r.isNil(pstate) {
			pstate.yyerror("invalid operation: %v (func can only be compared to nil)", n)
			n.Type = nil
			return n
		}

		if l.Type.IsStruct() {
			if f := pstate.IncomparableField(l.Type); f != nil {
				pstate.yyerror("invalid operation: %v (struct containing %v cannot be compared)", n, f.Type)
				n.Type = nil
				return n
			}
		}

		t = l.Type
		if pstate.iscmp[n.Op] {
			pstate.evconst(n)
			t = pstate.types.Idealbool
			if n.Op != OLITERAL {
				l, r = pstate.defaultlit2(l, r, true)
				n.Left = l
				n.Right = r
			}
		}

		if et == TSTRING {
			if pstate.iscmp[n.Op] {
				ot := n.Op
				n.Op = OCMPSTR
				n.SetSubOp(pstate, ot)
			} else if n.Op == OADD {
				// create OADDSTR node with list of strings in x + y + z + (w + v) + ...
				n.Op = OADDSTR

				if l.Op == OADDSTR {
					n.List.Set(l.List.Slice())
				} else {
					n.List.Set1(l)
				}
				if r.Op == OADDSTR {
					n.List.AppendNodes(&r.List)
				} else {
					n.List.Append(r)
				}
				n.Left = nil
				n.Right = nil
			}
		}

		if et == TINTER {
			if l.Op == OLITERAL && l.Val().Ctype(pstate) == CTNIL {
				// swap for back end
				n.Left = r

				n.Right = l
			} else if r.Op == OLITERAL && r.Val().Ctype(pstate) == CTNIL {
			} else // leave alone for back end
			if r.Type.IsInterface() == l.Type.IsInterface() {
				ot := n.Op
				n.Op = OCMPIFACE
				n.SetSubOp(pstate, ot)
			}
		}

		if (op == ODIV || op == OMOD) && pstate.Isconst(r, CTINT) {
			if r.Val().U.(*Mpint).CmpInt64(0) == 0 {
				pstate.yyerror("division by zero")
				n.Type = nil
				return n
			}
		}

		n.Type = t

	case OCOM, OMINUS, ONOT, OPLUS:
		ok |= Erv
		n.Left = pstate.typecheck(n.Left, Erv)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !pstate.okfor[n.Op][t.Etype] {
			pstate.yyerror("invalid operation: %v %v", n.Op, t)
			n.Type = nil
			return n
		}

		n.Type = t

	// exprs
	case OADDR:
		ok |= Erv

		n.Left = pstate.typecheck(n.Left, Erv)
		if n.Left.Type == nil {
			n.Type = nil
			return n
		}
		pstate.checklvalue(n.Left, "take the address of")
		r := pstate.outervalue(n.Left)
		var l *Node
		for l = n.Left; l != r; l = l.Left {
			l.SetAddrtaken(true)
			if l.IsClosureVar() && !pstate.capturevarscomplete {
				// Mark the original variable as Addrtaken so that capturevars
				// knows not to pass it by value.
				// But if the capturevars phase is complete, don't touch it,
				// in case l.Name's containing function has not yet been compiled.
				l.Name.Defn.SetAddrtaken(true)
			}
		}

		if l.Orig != l && l.Op == ONAME {
			pstate.Fatalf("found non-orig name node %v", l)
		}
		l.SetAddrtaken(true)
		if l.IsClosureVar() && !pstate.capturevarscomplete {
			// See comments above about closure variables.
			l.Name.Defn.SetAddrtaken(true)
		}
		n.Left = pstate.defaultlit(n.Left, nil)
		l = n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		n.Type = pstate.types.NewPtr(t)

	case OCOMPLIT:
		ok |= Erv
		n = pstate.typecheckcomplit(n)
		if n.Type == nil {
			return n
		}

	case OXDOT, ODOT:
		if n.Op == OXDOT {
			n = pstate.adddot(n)
			n.Op = ODOT
			if n.Left == nil {
				n.Type = nil
				return n
			}
		}

		n.Left = pstate.typecheck(n.Left, Erv|Etype)

		n.Left = pstate.defaultlit(n.Left, nil)

		t := n.Left.Type
		if t == nil {
			pstate.adderrorname(n)
			n.Type = nil
			return n
		}

		s := n.Sym

		if n.Left.Op == OTYPE {
			n = pstate.typecheckMethodExpr(n)
			if n.Type == nil {
				return n
			}
			ok = Erv
			break
		}

		if t.IsPtr() && !t.Elem(pstate.types).IsInterface() {
			t = t.Elem(pstate.types)
			if t == nil {
				n.Type = nil
				return n
			}
			n.Op = ODOTPTR
			pstate.checkwidth(t)
		}

		if n.Sym.IsBlank() {
			pstate.yyerror("cannot refer to blank field or method")
			n.Type = nil
			return n
		}

		if pstate.lookdot(n, t, 0) == nil {
			// Legitimate field or method lookup failed, try to explain the error
			switch {
			case t.IsEmptyInterface(pstate.types):
				pstate.yyerror("%v undefined (type %v is interface with no methods)", n, n.Left.Type)

			case t.IsPtr() && t.Elem(pstate.types).IsInterface():
				// Pointer to interface is almost always a mistake.
				pstate.yyerror("%v undefined (type %v is pointer to interface, not interface)", n, n.Left.Type)

			case pstate.lookdot(n, t, 1) != nil:
				// Field or method matches by name, but it is not exported.
				pstate.yyerror("%v undefined (cannot refer to unexported field or method %v)", n, n.Sym)

			default:
				if mt := pstate.lookdot(n, t, 2); mt != nil && pstate.visible(mt.Sym) { // Case-insensitive lookup.
					pstate.yyerror("%v undefined (type %v has no field or method %v, but does have %v)", n, n.Left.Type, n.Sym, mt.Sym)
				} else {
					pstate.yyerror("%v undefined (type %v has no field or method %v)", n, n.Left.Type, n.Sym)
				}
			}
			n.Type = nil
			return n
		}

		switch n.Op {
		case ODOTINTER, ODOTMETH:
			if top&Ecall != 0 {
				ok |= Ecall
			} else {
				pstate.typecheckpartialcall(n, s)
				ok |= Erv
			}

		default:
			ok |= Erv
		}

	case ODOTTYPE:
		ok |= Erv
		n.Left = pstate.typecheck(n.Left, Erv)
		n.Left = pstate.defaultlit(n.Left, nil)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsInterface() {
			pstate.yyerror("invalid type assertion: %v (non-interface type %v on left)", n, t)
			n.Type = nil
			return n
		}

		if n.Right != nil {
			n.Right = pstate.typecheck(n.Right, Etype)
			n.Type = n.Right.Type
			n.Right = nil
			if n.Type == nil {
				return n
			}
		}

		if n.Type != nil && !n.Type.IsInterface() {
			var missing, have *types.Field
			var ptr int
			if !pstate.implements(n.Type, t, &missing, &have, &ptr) {
				if have != nil && have.Sym == missing.Sym {
					pstate.yyerror("impossible type assertion:\n\t%v does not implement %v (wrong type for %v method)\n"+
						"\t\thave %v%0S\n\t\twant %v%0S", n.Type, t, missing.Sym, have.Sym, have.Type, missing.Sym, missing.Type)
				} else if ptr != 0 {
					pstate.yyerror("impossible type assertion:\n\t%v does not implement %v (%v method has pointer receiver)", n.Type, t, missing.Sym)
				} else if have != nil {
					pstate.yyerror("impossible type assertion:\n\t%v does not implement %v (missing %v method)\n"+
						"\t\thave %v%0S\n\t\twant %v%0S", n.Type, t, missing.Sym, have.Sym, have.Type, missing.Sym, missing.Type)
				} else {
					pstate.yyerror("impossible type assertion:\n\t%v does not implement %v (missing %v method)", n.Type, t, missing.Sym)
				}
				n.Type = nil
				return n
			}
		}

	case OINDEX:
		ok |= Erv
		n.Left = pstate.typecheck(n.Left, Erv)
		n.Left = pstate.defaultlit(n.Left, nil)
		n.Left = pstate.implicitstar(n.Left)
		l := n.Left
		n.Right = pstate.typecheck(n.Right, Erv)
		r := n.Right
		t := l.Type
		if t == nil || r.Type == nil {
			n.Type = nil
			return n
		}
		switch t.Etype {
		default:
			pstate.yyerror("invalid operation: %v (type %v does not support indexing)", n, t)
			n.Type = nil
			return n

		case TSTRING, TARRAY, TSLICE:
			n.Right = pstate.indexlit(n.Right)
			if t.IsString() {
				n.Type = pstate.types.Bytetype
			} else {
				n.Type = t.Elem(pstate.types)
			}
			why := "string"
			if t.IsArray() {
				why = "array"
			} else if t.IsSlice() {
				why = "slice"
			}

			if n.Right.Type != nil && !n.Right.Type.IsInteger() {
				pstate.yyerror("non-integer %s index %v", why, n.Right)
				break
			}

			if !n.Bounded() && pstate.Isconst(n.Right, CTINT) {
				x := n.Right.Int64(pstate)
				if x < 0 {
					pstate.yyerror("invalid %s index %v (index must be non-negative)", why, n.Right)
				} else if t.IsArray() && x >= t.NumElem(pstate.types) {
					pstate.yyerror("invalid array index %v (out of bounds for %d-element array)", n.Right, t.NumElem(pstate.types))
				} else if pstate.Isconst(n.Left, CTSTR) && x >= int64(len(n.Left.Val().U.(string))) {
					pstate.yyerror("invalid string index %v (out of bounds for %d-byte string)", n.Right, len(n.Left.Val().U.(string)))
				} else if n.Right.Val().U.(*Mpint).Cmp(pstate.maxintval[TINT]) > 0 {
					pstate.yyerror("invalid %s index %v (index too large)", why, n.Right)
				}
			}

		case TMAP:
			n.Right = pstate.defaultlit(n.Right, t.Key(pstate.types))
			if n.Right.Type != nil {
				n.Right = pstate.assignconv(n.Right, t.Key(pstate.types), "map index")
			}
			n.Type = t.Elem(pstate.types)
			n.Op = OINDEXMAP
			n.ResetAux()
		}

	case ORECV:
		ok |= Etop | Erv
		n.Left = pstate.typecheck(n.Left, Erv)
		n.Left = pstate.defaultlit(n.Left, nil)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsChan() {
			pstate.yyerror("invalid operation: %v (receive from non-chan type %v)", n, t)
			n.Type = nil
			return n
		}

		if !t.ChanDir(pstate.types).CanRecv() {
			pstate.yyerror("invalid operation: %v (receive from send-only type %v)", n, t)
			n.Type = nil
			return n
		}

		n.Type = t.Elem(pstate.types)

	case OSEND:
		ok |= Etop
		n.Left = pstate.typecheck(n.Left, Erv)
		n.Right = pstate.typecheck(n.Right, Erv)
		n.Left = pstate.defaultlit(n.Left, nil)
		t := n.Left.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsChan() {
			pstate.yyerror("invalid operation: %v (send to non-chan type %v)", n, t)
			n.Type = nil
			return n
		}

		if !t.ChanDir(pstate.types).CanSend() {
			pstate.yyerror("invalid operation: %v (send to receive-only type %v)", n, t)
			n.Type = nil
			return n
		}

		n.Right = pstate.defaultlit(n.Right, t.Elem(pstate.types))
		r := n.Right
		if r.Type == nil {
			n.Type = nil
			return n
		}
		n.Right = pstate.assignconv(r, t.Elem(pstate.types), "send")
		n.Type = nil

	case OSLICE, OSLICE3:
		ok |= Erv
		n.Left = pstate.typecheck(n.Left, Erv)
		low, high, max := n.SliceBounds(pstate)
		hasmax := n.Op.IsSlice3(pstate)
		low = pstate.typecheck(low, Erv)
		high = pstate.typecheck(high, Erv)
		max = pstate.typecheck(max, Erv)
		n.Left = pstate.defaultlit(n.Left, nil)
		low = pstate.indexlit(low)
		high = pstate.indexlit(high)
		max = pstate.indexlit(max)
		n.SetSliceBounds(pstate, low, high, max)
		l := n.Left
		if l.Type == nil {
			n.Type = nil
			return n
		}
		if l.Type.IsArray() {
			if !islvalue(n.Left) {
				pstate.yyerror("invalid operation %v (slice of unaddressable value)", n)
				n.Type = nil
				return n
			}

			n.Left = pstate.nod(OADDR, n.Left, nil)
			n.Left.SetImplicit(true)
			n.Left = pstate.typecheck(n.Left, Erv)
			l = n.Left
		}
		t := l.Type
		var tp *types.Type
		if t.IsString() {
			if hasmax {
				pstate.yyerror("invalid operation %v (3-index slice of string)", n)
				n.Type = nil
				return n
			}
			n.Type = t
			n.Op = OSLICESTR
		} else if t.IsPtr() && t.Elem(pstate.types).IsArray() {
			tp = t.Elem(pstate.types)
			n.Type = pstate.types.NewSlice(tp.Elem(pstate.types))
			pstate.dowidth(n.Type)
			if hasmax {
				n.Op = OSLICE3ARR
			} else {
				n.Op = OSLICEARR
			}
		} else if t.IsSlice() {
			n.Type = t
		} else {
			pstate.yyerror("cannot slice %v (type %v)", l, t)
			n.Type = nil
			return n
		}

		if low != nil && !pstate.checksliceindex(l, low, tp) {
			n.Type = nil
			return n
		}
		if high != nil && !pstate.checksliceindex(l, high, tp) {
			n.Type = nil
			return n
		}
		if max != nil && !pstate.checksliceindex(l, max, tp) {
			n.Type = nil
			return n
		}
		if !pstate.checksliceconst(low, high) || !pstate.checksliceconst(low, max) || !pstate.checksliceconst(high, max) {
			n.Type = nil
			return n
		}

	// call and call like
	case OCALL:
		n.Left = pstate.typecheck(n.Left, Erv|Etype|Ecall)
		if n.Left.Diag() {
			n.SetDiag(true)
		}

		l := n.Left

		if l.Op == ONAME && l.SubOp(pstate) != 0 {
			if n.Isddd() && l.SubOp(pstate) != OAPPEND {
				pstate.yyerror("invalid use of ... with builtin %v", l)
			}

			// builtin: OLEN, OCAP, etc.
			n.Op = l.SubOp(pstate)
			n.Left = n.Right
			n.Right = nil
			n = pstate.typecheck1(n, top)
			return n
		}

		n.Left = pstate.defaultlit(n.Left, nil)
		l = n.Left
		if l.Op == OTYPE {
			if n.Isddd() || l.Type.IsDDDArray() {
				if !l.Type.Broke() {
					pstate.yyerror("invalid use of ... in type conversion to %v", l.Type)
				}
				n.SetDiag(true)
			}

			// pick off before type-checking arguments
			ok |= Erv

			// turn CALL(type, arg) into CONV(arg) w/ type
			n.Left = nil

			n.Op = OCONV
			n.Type = l.Type
			if !pstate.onearg(n, "conversion to %v", l.Type) {
				n.Type = nil
				return n
			}
			n = pstate.typecheck1(n, top)
			return n
		}

		if n.List.Len() == 1 && !n.Isddd() {
			n.List.SetFirst(pstate.typecheck(n.List.First(), Erv|Efnstruct))
		} else {
			pstate.typecheckslice(n.List.Slice(), Erv)
		}
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		pstate.checkwidth(t)

		switch l.Op {
		case ODOTINTER:
			n.Op = OCALLINTER

		case ODOTMETH:
			n.Op = OCALLMETH

			// typecheckaste was used here but there wasn't enough
			// information further down the call chain to know if we
			// were testing a method receiver for unexported fields.
			// It isn't necessary, so just do a sanity check.
			tp := t.Recv(pstate.types).Type

			if l.Left == nil || !pstate.eqtype(l.Left.Type, tp) {
				pstate.Fatalf("method receiver")
			}

		default:
			n.Op = OCALLFUNC
			if t.Etype != TFUNC {
				name := l.String()
				if pstate.isBuiltinFuncName(name) {
					// be more specific when the function
					// name matches a predeclared function
					pstate.yyerror("cannot call non-function %s (type %v), declared at %s",
						name, t, pstate.linestr(l.Name.Defn.Pos))
				} else {
					pstate.yyerror("cannot call non-function %s (type %v)", name, t)
				}
				n.Type = nil
				return n
			}
		}

		pstate.typecheckaste(OCALL, n.Left, n.Isddd(), t.Params(pstate.types), n.List, func() string { return fmt.Sprintf("argument to %v", n.Left) })
		ok |= Etop
		if t.NumResults(pstate.types) == 0 {
			break
		}
		ok |= Erv
		if t.NumResults(pstate.types) == 1 {
			n.Type = l.Type.Results(pstate.types).Field(pstate.types, 0).Type

			if n.Op == OCALLFUNC && n.Left.Op == ONAME && pstate.isRuntimePkg(n.Left.Sym.Pkg) && n.Left.Sym.Name == "getg" {
				// Emit code for runtime.getg() directly instead of calling function.
				// Most such rewrites (for example the similar one for math.Sqrt) should be done in walk,
				// so that the ordering pass can make sure to preserve the semantics of the original code
				// (in particular, the exact time of the function call) by introducing temporaries.
				// In this case, we know getg() always returns the same result within a given function
				// and we want to avoid the temporaries, so we do the rewrite earlier than is typical.
				n.Op = OGETG
			}

			break
		}

		// multiple return
		if top&(Efnstruct|Etop) == 0 {
			pstate.yyerror("multiple-value %v() in single-value context", l)
			break
		}

		n.Type = l.Type.Results(pstate.types)

	case OALIGNOF, OOFFSETOF, OSIZEOF:
		ok |= Erv
		if !pstate.onearg(n, "%v", n.Op) {
			n.Type = nil
			return n
		}
		n.Type = pstate.types.Types[TUINTPTR]

		// any side effects disappear; ignore init
		pstate.setintconst(n, pstate.evalunsafe(n))

	case OCAP, OLEN:
		ok |= Erv
		if !pstate.onearg(n, "%v", n.Op) {
			n.Type = nil
			return n
		}

		n.Left = pstate.typecheck(n.Left, Erv)
		n.Left = pstate.defaultlit(n.Left, nil)
		n.Left = pstate.implicitstar(n.Left)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}

		var ok bool
		if n.Op == OLEN {
			ok = pstate.okforlen[t.Etype]
		} else {
			ok = pstate.okforcap[t.Etype]
		}
		if !ok {
			pstate.yyerror("invalid argument %L for %v", l, n.Op)
			n.Type = nil
			return n
		}

		n.Type = pstate.types.Types[TINT]

		// Result might be constant.
		var res int64 = -1 // valid if >= 0
		switch t.Etype {
		case TSTRING:
			if pstate.Isconst(l, CTSTR) {
				res = int64(len(l.Val().U.(string)))
			}

		case TARRAY:
			if !callrecv(l) {
				res = t.NumElem(pstate.types)
			}
		}
		if res >= 0 {
			pstate.setintconst(n, res)
		}

	case OREAL, OIMAG:
		ok |= Erv
		if !pstate.onearg(n, "%v", n.Op) {
			n.Type = nil
			return n
		}

		n.Left = pstate.typecheck(n.Left, Erv)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}

		// Determine result type.
		et := t.Etype
		switch et {
		case TIDEAL:
		// result is ideal
		case TCOMPLEX64:
			et = TFLOAT32
		case TCOMPLEX128:
			et = TFLOAT64
		default:
			pstate.yyerror("invalid argument %L for %v", l, n.Op)
			n.Type = nil
			return n
		}
		n.Type = pstate.types.Types[et]

		// if the argument is a constant, the result is a constant
		// (any untyped numeric constant can be represented as a
		// complex number)
		if l.Op == OLITERAL {
			var re, im *Mpflt
			switch pstate.consttype(l) {
			case CTINT, CTRUNE:
				re = newMpflt()
				re.SetInt(l.Val().U.(*Mpint))
			// im = 0
			case CTFLT:
				re = l.Val().U.(*Mpflt)
			// im = 0
			case CTCPLX:
				re = &l.Val().U.(*Mpcplx).Real
				im = &l.Val().U.(*Mpcplx).Imag
			default:
				pstate.yyerror("invalid argument %L for %v", l, n.Op)
				n.Type = nil
				return n
			}
			if n.Op == OIMAG {
				if im == nil {
					im = newMpflt()
				}
				re = im
			}
			pstate.setconst(n, Val{re})
		}

	case OCOMPLEX:
		ok |= Erv
		var r *Node
		var l *Node
		if n.List.Len() == 1 {
			pstate.typecheckslice(n.List.Slice(), Efnstruct)
			if n.List.First().Op != OCALLFUNC && n.List.First().Op != OCALLMETH {
				pstate.yyerror("invalid operation: complex expects two arguments")
				n.Type = nil
				return n
			}

			t := n.List.First().Left.Type
			if !t.IsKind(TFUNC) {
				// Bail. This error will be reported elsewhere.
				return n
			}
			if t.NumResults(pstate.types) != 2 {
				pstate.yyerror("invalid operation: complex expects two arguments, %v returns %d results", n.List.First(), t.NumResults(pstate.types))
				n.Type = nil
				return n
			}

			t = n.List.First().Type
			l = asNode(t.Field(pstate.types, 0).Nname)
			r = asNode(t.Field(pstate.types, 1).Nname)
		} else {
			if !pstate.twoarg(n) {
				n.Type = nil
				return n
			}
			n.Left = pstate.typecheck(n.Left, Erv)
			n.Right = pstate.typecheck(n.Right, Erv)
			l = n.Left
			r = n.Right
			if l.Type == nil || r.Type == nil {
				n.Type = nil
				return n
			}
			l, r = pstate.defaultlit2(l, r, false)
			if l.Type == nil || r.Type == nil {
				n.Type = nil
				return n
			}
			n.Left = l
			n.Right = r
		}

		if !pstate.eqtype(l.Type, r.Type) {
			pstate.yyerror("invalid operation: %v (mismatched types %v and %v)", n, l.Type, r.Type)
			n.Type = nil
			return n
		}

		var t *types.Type
		switch l.Type.Etype {
		default:
			pstate.yyerror("invalid operation: %v (arguments have type %v, expected floating-point)", n, l.Type)
			n.Type = nil
			return n

		case TIDEAL:
			t = pstate.types.Types[TIDEAL]

		case TFLOAT32:
			t = pstate.types.Types[TCOMPLEX64]

		case TFLOAT64:
			t = pstate.types.Types[TCOMPLEX128]
		}
		n.Type = t

		if l.Op == OLITERAL && r.Op == OLITERAL {
			// make it a complex literal
			c := new(Mpcplx)
			c.Real.Set(pstate.toflt(l.Val()).U.(*Mpflt))
			c.Imag.Set(pstate.toflt(r.Val()).U.(*Mpflt))
			pstate.setconst(n, Val{c})
		}

	case OCLOSE:
		if !pstate.onearg(n, "%v", n.Op) {
			n.Type = nil
			return n
		}
		n.Left = pstate.typecheck(n.Left, Erv)
		n.Left = pstate.defaultlit(n.Left, nil)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsChan() {
			pstate.yyerror("invalid operation: %v (non-chan type %v)", n, t)
			n.Type = nil
			return n
		}

		if !t.ChanDir(pstate.types).CanSend() {
			pstate.yyerror("invalid operation: %v (cannot close receive-only channel)", n)
			n.Type = nil
			return n
		}

		ok |= Etop

	case ODELETE:
		args := n.List
		if args.Len() == 0 {
			pstate.yyerror("missing arguments to delete")
			n.Type = nil
			return n
		}

		if args.Len() == 1 {
			pstate.yyerror("missing second (key) argument to delete")
			n.Type = nil
			return n
		}

		if args.Len() != 2 {
			pstate.yyerror("too many arguments to delete")
			n.Type = nil
			return n
		}

		ok |= Etop
		pstate.typecheckslice(args.Slice(), Erv)
		l := args.First()
		r := args.Second()
		if l.Type != nil && !l.Type.IsMap() {
			pstate.yyerror("first argument to delete must be map; have %L", l.Type)
			n.Type = nil
			return n
		}

		args.SetSecond(pstate.assignconv(r, l.Type.Key(pstate.types), "delete"))

	case OAPPEND:
		ok |= Erv
		args := n.List
		if args.Len() == 0 {
			pstate.yyerror("missing arguments to append")
			n.Type = nil
			return n
		}

		if args.Len() == 1 && !n.Isddd() {
			args.SetFirst(pstate.typecheck(args.First(), Erv|Efnstruct))
		} else {
			pstate.typecheckslice(args.Slice(), Erv)
		}

		t := args.First().Type
		if t == nil {
			n.Type = nil
			return n
		}

		// Unpack multiple-return result before type-checking.
		var funarg *types.Type
		if t.IsFuncArgStruct() {
			funarg = t
			t = t.Field(pstate.types, 0).Type
		}

		n.Type = t
		if !t.IsSlice() {
			if pstate.Isconst(args.First(), CTNIL) {
				pstate.yyerror("first argument to append must be typed slice; have untyped nil")
				n.Type = nil
				return n
			}

			pstate.yyerror("first argument to append must be slice; have %L", t)
			n.Type = nil
			return n
		}

		if n.Isddd() {
			if args.Len() == 1 {
				pstate.yyerror("cannot use ... on first argument to append")
				n.Type = nil
				return n
			}

			if args.Len() != 2 {
				pstate.yyerror("too many arguments to append")
				n.Type = nil
				return n
			}

			if t.Elem(pstate.types).IsKind(TUINT8) && args.Second().Type.IsString() {
				args.SetSecond(pstate.defaultlit(args.Second(), pstate.types.Types[TSTRING]))
				break
			}

			args.SetSecond(pstate.assignconv(args.Second(), t.Orig, "append"))
			break
		}

		if funarg != nil {
			for _, t := range funarg.FieldSlice(pstate.types)[1:] {
				if pstate.assignop(t.Type, n.Type.Elem(pstate.types), nil) == 0 {
					pstate.yyerror("cannot append %v value to []%v", t.Type, n.Type.Elem(pstate.types))
				}
			}
		} else {
			as := args.Slice()[1:]
			for i, n := range as {
				if n.Type == nil {
					continue
				}
				as[i] = pstate.assignconv(n, t.Elem(pstate.types), "append")
				pstate.checkwidth(as[i].Type) // ensure width is calculated for backend
			}
		}

	case OCOPY:
		ok |= Etop | Erv
		args := n.List
		if args.Len() < 2 {
			pstate.yyerror("missing arguments to copy")
			n.Type = nil
			return n
		}

		if args.Len() > 2 {
			pstate.yyerror("too many arguments to copy")
			n.Type = nil
			return n
		}

		n.Left = args.First()
		n.Right = args.Second()
		n.List.Set(nil)
		n.Type = pstate.types.Types[TINT]
		n.Left = pstate.typecheck(n.Left, Erv)
		n.Right = pstate.typecheck(n.Right, Erv)
		if n.Left.Type == nil || n.Right.Type == nil {
			n.Type = nil
			return n
		}
		n.Left = pstate.defaultlit(n.Left, nil)
		n.Right = pstate.defaultlit(n.Right, nil)
		if n.Left.Type == nil || n.Right.Type == nil {
			n.Type = nil
			return n
		}

		// copy([]byte, string)
		if n.Left.Type.IsSlice() && n.Right.Type.IsString() {
			if pstate.eqtype(n.Left.Type.Elem(pstate.types), pstate.types.Bytetype) {
				break
			}
			pstate.yyerror("arguments to copy have different element types: %L and string", n.Left.Type)
			n.Type = nil
			return n
		}

		if !n.Left.Type.IsSlice() || !n.Right.Type.IsSlice() {
			if !n.Left.Type.IsSlice() && !n.Right.Type.IsSlice() {
				pstate.yyerror("arguments to copy must be slices; have %L, %L", n.Left.Type, n.Right.Type)
			} else if !n.Left.Type.IsSlice() {
				pstate.yyerror("first argument to copy should be slice; have %L", n.Left.Type)
			} else {
				pstate.yyerror("second argument to copy should be slice or string; have %L", n.Right.Type)
			}
			n.Type = nil
			return n
		}

		if !pstate.eqtype(n.Left.Type.Elem(pstate.types), n.Right.Type.Elem(pstate.types)) {
			pstate.yyerror("arguments to copy have different element types: %L and %L", n.Left.Type, n.Right.Type)
			n.Type = nil
			return n
		}

	case OCONV:
		ok |= Erv
		pstate.checkwidth(n.Type) // ensure width is calculated for backend
		n.Left = pstate.typecheck(n.Left, Erv)
		n.Left = pstate.convlit1(n.Left, n.Type, true, noReuse)
		t := n.Left.Type
		if t == nil || n.Type == nil {
			n.Type = nil
			return n
		}
		var why string
		n.Op = pstate.convertop(t, n.Type, &why)
		if n.Op == 0 {
			if !n.Diag() && !n.Type.Broke() && !n.Left.Diag() {
				pstate.yyerror("cannot convert %L to type %v%s", n.Left, n.Type, why)
				n.SetDiag(true)
			}
			n.Op = OCONV
			n.Type = nil
			return n
		}

		switch n.Op {
		case OCONVNOP:
			if n.Left.Op == OLITERAL {
				n.Op = OCONV
				pstate.setconst(n, n.Left.Val())
			} else if t.Etype == n.Type.Etype {
				switch t.Etype {
				case TFLOAT32, TFLOAT64, TCOMPLEX64, TCOMPLEX128:
					// Floating point casts imply rounding and
					// so the conversion must be kept.
					n.Op = OCONV
				}
			}

		// do not use stringtoarraylit.
		// generated code and compiler memory footprint is better without it.
		case OSTRARRAYBYTE:
			break

		case OSTRARRAYRUNE:
			if n.Left.Op == OLITERAL {
				n = pstate.stringtoarraylit(n)
			}
		}

	case OMAKE:
		ok |= Erv
		args := n.List.Slice()
		if len(args) == 0 {
			pstate.yyerror("missing argument to make")
			n.Type = nil
			return n
		}

		n.List.Set(nil)
		l := args[0]
		l = pstate.typecheck(l, Etype)
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}

		i := 1
		switch t.Etype {
		default:
			pstate.yyerror("cannot make type %v", t)
			n.Type = nil
			return n

		case TSLICE:
			if i >= len(args) {
				pstate.yyerror("missing len argument to make(%v)", t)
				n.Type = nil
				return n
			}

			l = args[i]
			i++
			l = pstate.typecheck(l, Erv)
			var r *Node
			if i < len(args) {
				r = args[i]
				i++
				r = pstate.typecheck(r, Erv)
			}

			if l.Type == nil || (r != nil && r.Type == nil) {
				n.Type = nil
				return n
			}
			if !pstate.checkmake(t, "len", l) || r != nil && !pstate.checkmake(t, "cap", r) {
				n.Type = nil
				return n
			}
			if pstate.Isconst(l, CTINT) && r != nil && pstate.Isconst(r, CTINT) && l.Val().U.(*Mpint).Cmp(r.Val().U.(*Mpint)) > 0 {
				pstate.yyerror("len larger than cap in make(%v)", t)
				n.Type = nil
				return n
			}

			n.Left = l
			n.Right = r
			n.Op = OMAKESLICE

		case TMAP:
			if i < len(args) {
				l = args[i]
				i++
				l = pstate.typecheck(l, Erv)
				l = pstate.defaultlit(l, pstate.types.Types[TINT])
				if l.Type == nil {
					n.Type = nil
					return n
				}
				if !pstate.checkmake(t, "size", l) {
					n.Type = nil
					return n
				}
				n.Left = l
			} else {
				n.Left = pstate.nodintconst(0)
			}
			n.Op = OMAKEMAP

		case TCHAN:
			l = nil
			if i < len(args) {
				l = args[i]
				i++
				l = pstate.typecheck(l, Erv)
				l = pstate.defaultlit(l, pstate.types.Types[TINT])
				if l.Type == nil {
					n.Type = nil
					return n
				}
				if !pstate.checkmake(t, "buffer", l) {
					n.Type = nil
					return n
				}
				n.Left = l
			} else {
				n.Left = pstate.nodintconst(0)
			}
			n.Op = OMAKECHAN
		}

		if i < len(args) {
			pstate.yyerror("too many arguments to make(%v)", t)
			n.Op = OMAKE
			n.Type = nil
			return n
		}

		n.Type = t

	case ONEW:
		ok |= Erv
		args := n.List
		if args.Len() == 0 {
			pstate.yyerror("missing argument to new")
			n.Type = nil
			return n
		}

		l := args.First()
		l = pstate.typecheck(l, Etype)
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if args.Len() > 1 {
			pstate.yyerror("too many arguments to new(%v)", t)
			n.Type = nil
			return n
		}

		n.Left = l
		n.Type = pstate.types.NewPtr(t)

	case OPRINT, OPRINTN:
		ok |= Etop
		pstate.typecheckslice(n.List.Slice(), Erv)
		ls := n.List.Slice()
		for i1, n1 := range ls {
			// Special case for print: int constant is int64, not int.
			if pstate.Isconst(n1, CTINT) {
				ls[i1] = pstate.defaultlit(ls[i1], pstate.types.Types[TINT64])
			} else {
				ls[i1] = pstate.defaultlit(ls[i1], nil)
			}
		}

	case OPANIC:
		ok |= Etop
		if !pstate.onearg(n, "panic") {
			n.Type = nil
			return n
		}
		n.Left = pstate.typecheck(n.Left, Erv)
		n.Left = pstate.defaultlit(n.Left, pstate.types.Types[TINTER])
		if n.Left.Type == nil {
			n.Type = nil
			return n
		}

	case ORECOVER:
		ok |= Erv | Etop
		if n.List.Len() != 0 {
			pstate.yyerror("too many arguments to recover")
			n.Type = nil
			return n
		}

		n.Type = pstate.types.Types[TINTER]

	case OCLOSURE:
		ok |= Erv
		pstate.typecheckclosure(n, top)
		if n.Type == nil {
			return n
		}

	case OITAB:
		ok |= Erv
		n.Left = pstate.typecheck(n.Left, Erv)
		t := n.Left.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsInterface() {
			pstate.Fatalf("OITAB of %v", t)
		}
		n.Type = pstate.types.NewPtr(pstate.types.Types[TUINTPTR])

	case OIDATA:
		// Whoever creates the OIDATA node must know a priori the concrete type at that moment,
		// usually by just having checked the OITAB.
		pstate.Fatalf("cannot typecheck interface data %v", n)

	case OSPTR:
		ok |= Erv
		n.Left = pstate.typecheck(n.Left, Erv)
		t := n.Left.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsSlice() && !t.IsString() {
			pstate.Fatalf("OSPTR of %v", t)
		}
		if t.IsString() {
			n.Type = pstate.types.NewPtr(pstate.types.Types[TUINT8])
		} else {
			n.Type = pstate.types.NewPtr(t.Elem(pstate.types))
		}

	case OCLOSUREVAR:
		ok |= Erv

	case OCFUNC:
		ok |= Erv
		n.Left = pstate.typecheck(n.Left, Erv)
		n.Type = pstate.types.Types[TUINTPTR]

	case OCONVNOP:
		ok |= Erv
		n.Left = pstate.typecheck(n.Left, Erv)

	// statements
	case OAS:
		ok |= Etop

		pstate.typecheckas(n)

		// Code that creates temps does not bother to set defn, so do it here.
		if n.Left.Op == ONAME && n.Left.IsAutoTmp() {
			n.Left.Name.Defn = n
		}

	case OAS2:
		ok |= Etop
		pstate.typecheckas2(n)

	case OBREAK,
		OCONTINUE,
		ODCL,
		OEMPTY,
		OGOTO,
		OFALL,
		OVARKILL,
		OVARLIVE:
		ok |= Etop

	case OLABEL:
		ok |= Etop
		pstate.decldepth++
		if n.Left.Sym.IsBlank() {
			// Empty identifier is valid but useless.
			// Eliminate now to simplify life later.
			// See issues 7538, 11589, 11593.
			n.Op = OEMPTY
			n.Left = nil
		}

	case ODEFER:
		ok |= Etop
		n.Left = pstate.typecheck(n.Left, Etop|Erv)
		if !n.Left.Diag() {
			pstate.checkdefergo(n)
		}

	case OPROC:
		ok |= Etop
		n.Left = pstate.typecheck(n.Left, Etop|Erv)
		pstate.checkdefergo(n)

	case OFOR, OFORUNTIL:
		ok |= Etop
		pstate.typecheckslice(n.Ninit.Slice(), Etop)
		pstate.decldepth++
		n.Left = pstate.typecheck(n.Left, Erv)
		n.Left = pstate.defaultlit(n.Left, nil)
		if n.Left != nil {
			t := n.Left.Type
			if t != nil && !t.IsBoolean() {
				pstate.yyerror("non-bool %L used as for condition", n.Left)
			}
		}
		n.Right = pstate.typecheck(n.Right, Etop)
		if n.Op == OFORUNTIL {
			pstate.typecheckslice(n.List.Slice(), Etop)
		}
		pstate.typecheckslice(n.Nbody.Slice(), Etop)
		pstate.decldepth--

	case OIF:
		ok |= Etop
		pstate.typecheckslice(n.Ninit.Slice(), Etop)
		n.Left = pstate.typecheck(n.Left, Erv)
		n.Left = pstate.defaultlit(n.Left, nil)
		if n.Left != nil {
			t := n.Left.Type
			if t != nil && !t.IsBoolean() {
				pstate.yyerror("non-bool %L used as if condition", n.Left)
			}
		}
		pstate.typecheckslice(n.Nbody.Slice(), Etop)
		pstate.typecheckslice(n.Rlist.Slice(), Etop)

	case ORETURN:
		ok |= Etop
		if n.List.Len() == 1 {
			pstate.typecheckslice(n.List.Slice(), Erv|Efnstruct)
		} else {
			pstate.typecheckslice(n.List.Slice(), Erv)
		}
		if pstate.Curfn == nil {
			pstate.yyerror("return outside function")
			n.Type = nil
			return n
		}

		if pstate.Curfn.Type.FuncType(pstate.types).Outnamed && n.List.Len() == 0 {
			break
		}
		pstate.typecheckaste(ORETURN, nil, false, pstate.Curfn.Type.Results(pstate.types), n.List, func() string { return "return argument" })

	case ORETJMP:
		ok |= Etop

	case OSELECT:
		ok |= Etop
		pstate.typecheckselect(n)

	case OSWITCH:
		ok |= Etop
		pstate.typecheckswitch(n)

	case ORANGE:
		ok |= Etop
		pstate.typecheckrange(n)

	case OTYPESW:
		pstate.yyerror("use of .(type) outside type switch")
		n.Type = nil
		return n

	case OXCASE:
		ok |= Etop
		pstate.typecheckslice(n.List.Slice(), Erv)
		pstate.typecheckslice(n.Nbody.Slice(), Etop)

	case ODCLFUNC:
		ok |= Etop
		pstate.typecheckfunc(n)

	case ODCLCONST:
		ok |= Etop
		n.Left = pstate.typecheck(n.Left, Erv)

	case ODCLTYPE:
		ok |= Etop
		n.Left = pstate.typecheck(n.Left, Etype)
		pstate.checkwidth(n.Left.Type)
		if n.Left.Type != nil && n.Left.Type.NotInHeap() && n.Left.Name.Param.Pragma&NotInHeap == 0 {
			// The type contains go:notinheap types, so it
			// must be marked as such (alternatively, we
			// could silently propagate go:notinheap).
			pstate.yyerror("type %v must be go:notinheap", n.Left.Type)
		}
	}

	t := n.Type
	if t != nil && !t.IsFuncArgStruct() && n.Op != OTYPE {
		switch t.Etype {
		case TFUNC, // might have TANY; wait until it's called
			TANY, TFORW, TIDEAL, TNIL, TBLANK:
			break

		default:
			pstate.checkwidth(t)
		}
	}

	if pstate.safemode && !pstate.inimport && !pstate.compiling_wrappers && t != nil && t.Etype == TUNSAFEPTR {
		pstate.yyerror("cannot use unsafe.Pointer")
	}

	pstate.evconst(n)
	if n.Op == OTYPE && top&Etype == 0 {
		if !n.Type.Broke() {
			pstate.yyerror("type %v is not an expression", n.Type)
		}
		n.Type = nil
		return n
	}

	if top&(Erv|Etype) == Etype && n.Op != OTYPE {
		pstate.yyerror("%v is not a type", n)
		n.Type = nil
		return n
	}

	// TODO(rsc): simplify
	if (top&(Ecall|Erv|Etype) != 0) && top&Etop == 0 && ok&(Erv|Etype|Ecall) == 0 {
		pstate.yyerror("%v used as value", n)
		n.Type = nil
		return n
	}

	if (top&Etop != 0) && top&(Ecall|Erv|Etype) == 0 && ok&Etop == 0 {
		if !n.Diag() {
			pstate.yyerror("%v evaluated but not used", n)
			n.SetDiag(true)
		}

		n.Type = nil
		return n
	}

	return n
}

func (pstate *PackageState) checksliceindex(l *Node, r *Node, tp *types.Type) bool {
	t := r.Type
	if t == nil {
		return false
	}
	if !t.IsInteger() {
		pstate.yyerror("invalid slice index %v (type %v)", r, t)
		return false
	}

	if r.Op == OLITERAL {
		if r.Int64(pstate) < 0 {
			pstate.yyerror("invalid slice index %v (index must be non-negative)", r)
			return false
		} else if tp != nil && tp.NumElem(pstate.types) >= 0 && r.Int64(pstate) > tp.NumElem(pstate.types) {
			pstate.yyerror("invalid slice index %v (out of bounds for %d-element array)", r, tp.NumElem(pstate.types))
			return false
		} else if pstate.Isconst(l, CTSTR) && r.Int64(pstate) > int64(len(l.Val().U.(string))) {
			pstate.yyerror("invalid slice index %v (out of bounds for %d-byte string)", r, len(l.Val().U.(string)))
			return false
		} else if r.Val().U.(*Mpint).Cmp(pstate.maxintval[TINT]) > 0 {
			pstate.yyerror("invalid slice index %v (index too large)", r)
			return false
		}
	}

	return true
}

func (pstate *PackageState) checksliceconst(lo *Node, hi *Node) bool {
	if lo != nil && hi != nil && lo.Op == OLITERAL && hi.Op == OLITERAL && lo.Val().U.(*Mpint).Cmp(hi.Val().U.(*Mpint)) > 0 {
		pstate.yyerror("invalid slice index: %v > %v", lo, hi)
		return false
	}

	return true
}

func (pstate *PackageState) checkdefergo(n *Node) {
	what := "defer"
	if n.Op == OPROC {
		what = "go"
	}

	switch n.Left.Op {
	// ok
	case OCALLINTER,
		OCALLMETH,
		OCALLFUNC,
		OCLOSE,
		OCOPY,
		ODELETE,
		OPANIC,
		OPRINT,
		OPRINTN,
		ORECOVER:
		return

	case OAPPEND,
		OCAP,
		OCOMPLEX,
		OIMAG,
		OLEN,
		OMAKE,
		OMAKESLICE,
		OMAKECHAN,
		OMAKEMAP,
		ONEW,
		OREAL,
		OLITERAL: // conversion or unsafe.Alignof, Offsetof, Sizeof
		if n.Left.Orig != nil && n.Left.Orig.Op == OCONV {
			break
		}
		pstate.yyerror("%s discards result of %v", what, n.Left)
		return
	}

	// type is broken or missing, most likely a method call on a broken type
	// we will warn about the broken type elsewhere. no need to emit a potentially confusing error
	if n.Left.Type == nil || n.Left.Type.Broke() {
		return
	}

	if !n.Diag() {
		// The syntax made sure it was a call, so this must be
		// a conversion.
		n.SetDiag(true)
		pstate.yyerror("%s requires function call, not conversion", what)
	}
}

// The result of implicitstar MUST be assigned back to n, e.g.
// 	n.Left = implicitstar(n.Left)
func (pstate *PackageState) implicitstar(n *Node) *Node {
	// insert implicit * if needed for fixed array
	t := n.Type
	if t == nil || !t.IsPtr() {
		return n
	}
	t = t.Elem(pstate.types)
	if t == nil {
		return n
	}
	if !t.IsArray() {
		return n
	}
	n = pstate.nod(OIND, n, nil)
	n.SetImplicit(true)
	n = pstate.typecheck(n, Erv)
	return n
}

func (pstate *PackageState) onearg(n *Node, f string, args ...interface{}) bool {
	if n.Left != nil {
		return true
	}
	if n.List.Len() == 0 {
		p := fmt.Sprintf(f, args...)
		pstate.yyerror("missing argument to %s: %v", p, n)
		return false
	}

	if n.List.Len() > 1 {
		p := fmt.Sprintf(f, args...)
		pstate.yyerror("too many arguments to %s: %v", p, n)
		n.Left = n.List.First()
		n.List.Set(nil)
		return false
	}

	n.Left = n.List.First()
	n.List.Set(nil)
	return true
}

func (pstate *PackageState) twoarg(n *Node) bool {
	if n.Left != nil {
		return true
	}
	if n.List.Len() == 0 {
		pstate.yyerror("missing argument to %v - %v", n.Op, n)
		return false
	}

	n.Left = n.List.First()
	if n.List.Len() == 1 {
		pstate.yyerror("missing argument to %v - %v", n.Op, n)
		n.List.Set(nil)
		return false
	}

	if n.List.Len() > 2 {
		pstate.yyerror("too many arguments to %v - %v", n.Op, n)
		n.List.Set(nil)
		return false
	}

	n.Right = n.List.Second()
	n.List.Set(nil)
	return true
}

func (pstate *PackageState) lookdot1(errnode *Node, s *types.Sym, t *types.Type, fs *types.Fields, dostrcmp int) *types.Field {
	var r *types.Field
	for _, f := range fs.Slice() {
		if dostrcmp != 0 && f.Sym.Name == s.Name {
			return f
		}
		if dostrcmp == 2 && strings.EqualFold(f.Sym.Name, s.Name) {
			return f
		}
		if f.Sym != s {
			continue
		}
		if r != nil {
			if errnode != nil {
				pstate.yyerror("ambiguous selector %v", errnode)
			} else if t.IsPtr() {
				pstate.yyerror("ambiguous selector (%v).%v", t, s)
			} else {
				pstate.yyerror("ambiguous selector %v.%v", t, s)
			}
			break
		}

		r = f
	}

	return r
}

// typecheckMethodExpr checks selector expressions (ODOT) where the
// base expression is a type expression (OTYPE).
func (pstate *PackageState) typecheckMethodExpr(n *Node) *Node {
	t := n.Left.Type

	// Compute the method set for t.
	var ms *types.Fields
	if t.IsInterface() {
		ms = t.Fields(pstate.types)
	} else {
		mt := pstate.methtype(t)
		if mt == nil {
			pstate.yyerror("%v undefined (type %v has no method %v)", n, t, n.Sym)
			n.Type = nil
			return n
		}
		pstate.expandmeth(mt)
		ms = mt.AllMethods()

		// The method expression T.m requires a wrapper when T
		// is different from m's declared receiver type. We
		// normally generate these wrappers while writing out
		// runtime type descriptors, which is always done for
		// types declared at package scope. However, we need
		// to make sure to generate wrappers for anonymous
		// receiver types too.
		if mt.Sym == nil {
			pstate.addsignat(t)
		}
	}

	s := n.Sym
	m := pstate.lookdot1(n, s, t, ms, 0)
	if m == nil {
		if pstate.lookdot1(n, s, t, ms, 1) != nil {
			pstate.yyerror("%v undefined (cannot refer to unexported method %v)", n, s)
		} else if _, ambig := pstate.dotpath(s, t, nil, false); ambig {
			pstate.yyerror("%v undefined (ambiguous selector)", n) // method or field
		} else {
			pstate.yyerror("%v undefined (type %v has no method %v)", n, t, s)
		}
		n.Type = nil
		return n
	}

	if !pstate.isMethodApplicable(t, m) {
		pstate.yyerror("invalid method expression %v (needs pointer receiver: (*%v).%S)", n, t, s)
		n.Type = nil
		return n
	}

	n.Op = ONAME
	if n.Name == nil {
		n.Name = new(Name)
	}
	n.Right = pstate.newname(n.Sym)
	n.Sym = pstate.methodSym(t, n.Sym)
	n.Type = pstate.methodfunc(m.Type, n.Left.Type)
	n.Xoffset = 0
	n.SetClass(PFUNC)
	return n
}

// isMethodApplicable reports whether method m can be called on a
// value of type t. This is necessary because we compute a single
// method set for both T and *T, but some *T methods are not
// applicable to T receivers.
func (pstate *PackageState) isMethodApplicable(t *types.Type, m *types.Field) bool {
	return t.IsPtr() || !m.Type.Recv(pstate.types).Type.IsPtr() || pstate.isifacemethod(m.Type) || m.Embedded == 2
}

func (pstate *PackageState) derefall(t *types.Type) *types.Type {
	for t != nil && t.Etype == pstate.types.Tptr {
		t = t.Elem(pstate.types)
	}
	return t
}

type typeSymKey struct {
	t *types.Type
	s *types.Sym
}

func (pstate *PackageState) lookdot(n *Node, t *types.Type, dostrcmp int) *types.Field {
	s := n.Sym

	pstate.dowidth(t)
	var f1 *types.Field
	if t.IsStruct() || t.IsInterface() {
		f1 = pstate.lookdot1(n, s, t, t.Fields(pstate.types), dostrcmp)
	}

	var f2 *types.Field
	if n.Left.Type == t || n.Left.Type.Sym == nil {
		mt := pstate.methtype(t)
		if mt != nil {
			f2 = pstate.lookdot1(n, s, mt, mt.Methods(), dostrcmp)
		}
	}

	if f1 != nil {
		if dostrcmp > 1 {
			// Already in the process of diagnosing an error.
			return f1
		}
		if f2 != nil {
			pstate.yyerror("%v is both field and method", n.Sym)
		}
		if f1.Offset == BADWIDTH {
			pstate.Fatalf("lookdot badwidth %v %p", f1, f1)
		}
		n.Xoffset = f1.Offset
		n.Type = f1.Type
		if pstate.objabi.Fieldtrack_enabled > 0 {
			pstate.dotField[typeSymKey{t.Orig, s}] = f1
		}
		if t.IsInterface() {
			if n.Left.Type.IsPtr() {
				n.Left = pstate.nod(OIND, n.Left, nil) // implicitstar
				n.Left.SetImplicit(true)
				n.Left = pstate.typecheck(n.Left, Erv)
			}

			n.Op = ODOTINTER
		}

		return f1
	}

	if f2 != nil {
		if dostrcmp > 1 {
			// Already in the process of diagnosing an error.
			return f2
		}
		tt := n.Left.Type
		pstate.dowidth(tt)
		rcvr := f2.Type.Recv(pstate.types).Type
		if !pstate.eqtype(rcvr, tt) {
			if rcvr.Etype == pstate.types.Tptr && pstate.eqtype(rcvr.Elem(pstate.types), tt) {
				pstate.checklvalue(n.Left, "call pointer method on")
				n.Left = pstate.nod(OADDR, n.Left, nil)
				n.Left.SetImplicit(true)
				n.Left = pstate.typecheck(n.Left, Etype|Erv)
			} else if tt.Etype == pstate.types.Tptr && rcvr.Etype != pstate.types.Tptr && pstate.eqtype(tt.Elem(pstate.types), rcvr) {
				n.Left = pstate.nod(OIND, n.Left, nil)
				n.Left.SetImplicit(true)
				n.Left = pstate.typecheck(n.Left, Etype|Erv)
			} else if tt.Etype == pstate.types.Tptr && tt.Elem(pstate.types).Etype == pstate.types.Tptr && pstate.eqtype(pstate.derefall(tt), pstate.derefall(rcvr)) {
				pstate.yyerror("calling method %v with receiver %L requires explicit dereference", n.Sym, n.Left)
				for tt.Etype == pstate.types.Tptr {
					// Stop one level early for method with pointer receiver.
					if rcvr.Etype == pstate.types.Tptr && tt.Elem(pstate.types).Etype != pstate.types.Tptr {
						break
					}
					n.Left = pstate.nod(OIND, n.Left, nil)
					n.Left.SetImplicit(true)
					n.Left = pstate.typecheck(n.Left, Etype|Erv)
					tt = tt.Elem(pstate.types)
				}
			} else {
				pstate.Fatalf("method mismatch: %v for %v", rcvr, tt)
			}
		}

		pll := n
		ll := n.Left
		for ll.Left != nil && (ll.Op == ODOT || ll.Op == ODOTPTR || ll.Op == OIND) {
			pll = ll
			ll = ll.Left
		}
		if pll.Implicit() && ll.Type.IsPtr() && ll.Type.Sym != nil && asNode(ll.Type.Sym.Def) != nil && asNode(ll.Type.Sym.Def).Op == OTYPE {
			// It is invalid to automatically dereference a named pointer type when selecting a method.
			// Make n.Left == ll to clarify error message.
			n.Left = ll
			return nil
		}

		n.Sym = pstate.methodSym(n.Left.Type, f2.Sym)
		n.Xoffset = f2.Offset
		n.Type = f2.Type
		n.Op = ODOTMETH

		return f2
	}

	return nil
}

func nokeys(l Nodes) bool {
	for _, n := range l.Slice() {
		if n.Op == OKEY || n.Op == OSTRUCTKEY {
			return false
		}
	}
	return true
}

func (pstate *PackageState) hasddd(t *types.Type) bool {
	for _, tl := range t.Fields(pstate.types).Slice() {
		if tl.Isddd() {
			return true
		}
	}

	return false
}

// typecheck assignment: type list = expression list
func (pstate *PackageState) typecheckaste(op Op, call *Node, isddd bool, tstruct *types.Type, nl Nodes, desc func() string) {
	var t *types.Type
	var n1 int
	var n2 int
	var i int

	lno := pstate.lineno
	defer func() { pstate.lineno = lno }()

	if tstruct.Broke() {
		return
	}

	var n *Node
	if nl.Len() == 1 {
		n = nl.First()
		if n.Type != nil && n.Type.IsFuncArgStruct() {
			if !pstate.hasddd(tstruct) {
				n1 := tstruct.NumFields(pstate.types)
				n2 := n.Type.NumFields(pstate.types)
				if n2 > n1 {
					goto toomany
				}
				if n2 < n1 {
					goto notenough
				}
			}

			lfs := tstruct.FieldSlice(pstate.types)
			rfs := n.Type.FieldSlice(pstate.types)
			var why string
			for i, tl := range lfs {
				if tl.Isddd() {
					for _, tn := range rfs[i:] {
						if pstate.assignop(tn.Type, tl.Type.Elem(pstate.types), &why) == 0 {
							if call != nil {
								pstate.yyerror("cannot use %v as type %v in argument to %v%s", tn.Type, tl.Type.Elem(pstate.types), call, why)
							} else {
								pstate.yyerror("cannot use %v as type %v in %s%s", tn.Type, tl.Type.Elem(pstate.types), desc(), why)
							}
						}
					}
					return
				}

				if i >= len(rfs) {
					goto notenough
				}
				tn := rfs[i]
				if pstate.assignop(tn.Type, tl.Type, &why) == 0 {
					if call != nil {
						pstate.yyerror("cannot use %v as type %v in argument to %v%s", tn.Type, tl.Type, call, why)
					} else {
						pstate.yyerror("cannot use %v as type %v in %s%s", tn.Type, tl.Type, desc(), why)
					}
				}
			}

			if len(rfs) > len(lfs) {
				goto toomany
			}
			return
		}
	}

	n1 = tstruct.NumFields(pstate.types)
	n2 = nl.Len()
	if !pstate.hasddd(tstruct) {
		if n2 > n1 {
			goto toomany
		}
		if n2 < n1 {
			goto notenough
		}
	} else {
		if !isddd {
			if n2 < n1-1 {
				goto notenough
			}
		} else {
			if n2 > n1 {
				goto toomany
			}
			if n2 < n1 {
				goto notenough
			}
		}
	}

	i = 0
	for _, tl := range tstruct.Fields(pstate.types).Slice() {
		t = tl.Type
		if tl.Isddd() {
			if isddd {
				if i >= nl.Len() {
					goto notenough
				}
				if nl.Len()-i > 1 {
					goto toomany
				}
				n = nl.Index(i)
				pstate.setlineno(n)
				if n.Type != nil {
					nl.SetIndex(i, pstate.assignconvfn(n, t, desc))
				}
				return
			}

			for ; i < nl.Len(); i++ {
				n = nl.Index(i)
				pstate.setlineno(n)
				if n.Type != nil {
					nl.SetIndex(i, pstate.assignconvfn(n, t.Elem(pstate.types), desc))
				}
			}
			return
		}

		if i >= nl.Len() {
			goto notenough
		}
		n = nl.Index(i)
		pstate.setlineno(n)
		if n.Type != nil {
			nl.SetIndex(i, pstate.assignconvfn(n, t, desc))
		}
		i++
	}

	if i < nl.Len() {
		goto toomany
	}
	if isddd {
		if call != nil {
			pstate.yyerror("invalid use of ... in call to %v", call)
		} else {
			pstate.yyerror("invalid use of ... in %v", op)
		}
	}
	return

notenough:
	if n == nil || !n.Diag() {
		details := pstate.errorDetails(nl, tstruct, isddd)
		if call != nil {
			// call is the expression being called, not the overall call.
			// Method expressions have the form T.M, and the compiler has
			// rewritten those to ONAME nodes but left T in Left.
			if call.isMethodExpression() {
				pstate.yyerror("not enough arguments in call to method expression %v%s", call, details)
			} else {
				pstate.yyerror("not enough arguments in call to %v%s", call, details)
			}
		} else {
			pstate.yyerror("not enough arguments to %v%s", op, details)
		}
		if n != nil {
			n.SetDiag(true)
		}
	}
	return

toomany:
	details := pstate.errorDetails(nl, tstruct, isddd)
	if call != nil {
		pstate.yyerror("too many arguments in call to %v%s", call, details)
	} else {
		pstate.yyerror("too many arguments to %v%s", op, details)
	}
}

func (pstate *PackageState) errorDetails(nl Nodes, tstruct *types.Type, isddd bool) string {
	// If we don't know any type at a call site, let's suppress any return
	// message signatures. See Issue https://golang.org/issues/19012.
	if tstruct == nil {
		return ""
	}
	// If any node has an unknown type, suppress it as well
	for _, n := range nl.Slice() {
		if n.Type == nil {
			return ""
		}
	}
	return fmt.Sprintf("\n\thave %s\n\twant %v", nl.retsigerr(pstate, isddd), tstruct)
}

// sigrepr is a type's representation to the outside world,
// in string representations of return signatures
// e.g in error messages about wrong arguments to return.
func (pstate *PackageState) sigrepr(t *types.Type) string {
	switch t {
	default:
		return t.String(pstate.types)

	case pstate.types.Types[TIDEAL]:
		// "untyped number" is not commonly used
		// outside of the compiler, so let's use "number".
		return "number"

	case pstate.types.Idealstring:
		return "string"

	case pstate.types.Idealbool:
		return "bool"
	}
}

// retsigerr returns the signature of the types
// at the respective return call site of a function.
func (nl Nodes) retsigerr(pstate *PackageState, isddd bool) string {
	if nl.Len() < 1 {
		return "()"
	}

	var typeStrings []string
	if nl.Len() == 1 && nl.First().Type != nil && nl.First().Type.IsFuncArgStruct() {
		for _, f := range nl.First().Type.Fields(pstate.types).Slice() {
			typeStrings = append(typeStrings, pstate.sigrepr(f.Type))
		}
	} else {
		for _, n := range nl.Slice() {
			typeStrings = append(typeStrings, pstate.sigrepr(n.Type))
		}
	}

	ddd := ""
	if isddd {
		ddd = "..."
	}
	return fmt.Sprintf("(%s%s)", strings.Join(typeStrings, ", "), ddd)
}

// type check composite
func (pstate *PackageState) fielddup(name string, hash map[string]bool) {
	if hash[name] {
		pstate.yyerror("duplicate field name in struct literal: %s", name)
		return
	}
	hash[name] = true
}

func (pstate *PackageState) keydup(n *Node, hash map[uint32][]*Node) {
	orign := n
	if n.Op == OCONVIFACE {
		n = n.Left
	}
	pstate.evconst(n)
	if n.Op != OLITERAL {
		return // we don't check variables
	}

	const PRIME1 = 3

	var h uint32
	switch v := n.Val().U.(type) {
	default: // unknown, bool, nil
		h = 23

	case *Mpint:
		h = uint32(v.Int64(pstate))

	case *Mpflt:
		x := math.Float64bits(v.Float64(pstate))
		for i := 0; i < 8; i++ {
			h = h*PRIME1 + uint32(x&0xFF)
			x >>= 8
		}

	case string:
		for i := 0; i < len(v); i++ {
			h = h*PRIME1 + uint32(v[i])
		}
	}

	var cmp Node
	for _, a := range hash[h] {
		cmp.Op = OEQ
		cmp.Left = n
		if a.Op == OCONVIFACE && orign.Op == OCONVIFACE {
			a = a.Left
		}
		if !pstate.eqtype(a.Type, n.Type) {
			continue
		}
		cmp.Right = a
		pstate.evconst(&cmp)
		if cmp.Op != OLITERAL {
			// Sometimes evconst fails. See issue 12536.
			continue
		}
		if cmp.Val().U.(bool) {
			pstate.yyerror("duplicate key %v in map literal", n)
			return
		}
	}

	hash[h] = append(hash[h], orign)
}

// iscomptype reports whether type t is a composite literal type
// or a pointer to one.
func (pstate *PackageState) iscomptype(t *types.Type) bool {
	if t.IsPtr() {
		t = t.Elem(pstate.types)
	}

	switch t.Etype {
	case TARRAY, TSLICE, TSTRUCT, TMAP:
		return true
	default:
		return false
	}
}

func (pstate *PackageState) pushtype(n *Node, t *types.Type) {
	if n == nil || n.Op != OCOMPLIT || !pstate.iscomptype(t) {
		return
	}

	if n.Right == nil {
		n.Right = pstate.typenod(t)
		n.SetImplicit(true)       // don't print
		n.Right.SetImplicit(true) // * is okay
	} else if pstate.Debug['s'] != 0 {
		n.Right = pstate.typecheck(n.Right, Etype)
		if n.Right.Type != nil && pstate.eqtype(n.Right.Type, t) {
			fmt.Printf("%v: redundant type: %v\n", n.Line(pstate), t)
		}
	}
}

// The result of typecheckcomplit MUST be assigned back to n, e.g.
// 	n.Left = typecheckcomplit(n.Left)
func (pstate *PackageState) typecheckcomplit(n *Node) *Node {
	lno := pstate.lineno
	defer func() {
		pstate.lineno = lno
	}()

	if n.Right == nil {
		pstate.yyerrorl(n.Pos, "missing type in composite literal")
		n.Type = nil
		return n
	}

	// Save original node (including n.Right)
	norig := n.copy()

	pstate.setlineno(n.Right)
	n.Right = pstate.typecheck(n.Right, Etype|Ecomplit)
	l := n.Right // sic
	t := l.Type
	if t == nil {
		n.Type = nil
		return n
	}
	nerr := pstate.nerrors
	n.Type = t

	if t.IsPtr() {
		// For better or worse, we don't allow pointers as the composite literal type,
		// except when using the &T syntax, which sets implicit on the OIND.
		if !n.Right.Implicit() {
			pstate.yyerror("invalid pointer type %v for composite literal (use &%v instead)", t, t.Elem(pstate.types))
			n.Type = nil
			return n
		}

		// Also, the underlying type must be a struct, map, slice, or array.
		if !pstate.iscomptype(t) {
			pstate.yyerror("invalid pointer type %v for composite literal", t)
			n.Type = nil
			return n
		}

		t = t.Elem(pstate.types)
	}

	switch t.Etype {
	default:
		pstate.yyerror("invalid type for composite literal: %v", t)
		n.Type = nil

	case TARRAY, TSLICE:
		// If there are key/value pairs, create a map to keep seen
		// keys so we can check for duplicate indices.
		var indices map[int64]bool
		for _, n1 := range n.List.Slice() {
			if n1.Op == OKEY {
				indices = make(map[int64]bool)
				break
			}
		}

		var length, i int64
		checkBounds := t.IsArray() && !t.IsDDDArray()
		nl := n.List.Slice()
		for i2, l := range nl {
			pstate.setlineno(l)
			vp := &nl[i2]
			if l.Op == OKEY {
				l.Left = pstate.typecheck(l.Left, Erv)
				pstate.evconst(l.Left)
				i = pstate.nonnegintconst(l.Left)
				if i < 0 && !l.Left.Diag() {
					pstate.yyerror("index must be non-negative integer constant")
					l.Left.SetDiag(true)
					i = -(1 << 30) // stay negative for a while
				}
				vp = &l.Right
			}

			if i >= 0 && indices != nil {
				if indices[i] {
					pstate.yyerror("duplicate index in array literal: %d", i)
				} else {
					indices[i] = true
				}
			}

			r := *vp
			pstate.pushtype(r, t.Elem(pstate.types))
			r = pstate.typecheck(r, Erv)
			r = pstate.defaultlit(r, t.Elem(pstate.types))
			*vp = pstate.assignconv(r, t.Elem(pstate.types), "array or slice literal")

			i++
			if i > length {
				length = i
				if checkBounds && length > t.NumElem(pstate.types) {
					pstate.setlineno(l)
					pstate.yyerror("array index %d out of bounds [0:%d]", length-1, t.NumElem(pstate.types))
					checkBounds = false
				}
			}
		}

		if t.IsDDDArray() {
			t.SetNumElem(pstate.types, length)
		}
		if t.IsSlice() {
			n.Op = OSLICELIT
			n.Right = pstate.nodintconst(length)
		} else {
			n.Op = OARRAYLIT
			n.Right = nil
		}

	case TMAP:
		hash := make(map[uint32][]*Node)
		for i3, l := range n.List.Slice() {
			pstate.setlineno(l)
			if l.Op != OKEY {
				n.List.SetIndex(i3, pstate.typecheck(l, Erv))
				pstate.yyerror("missing key in map literal")
				continue
			}

			r := l.Left
			pstate.pushtype(r, t.Key(pstate.types))
			r = pstate.typecheck(r, Erv)
			r = pstate.defaultlit(r, t.Key(pstate.types))
			l.Left = pstate.assignconv(r, t.Key(pstate.types), "map key")
			if l.Left.Op != OCONV {
				pstate.keydup(l.Left, hash)
			}

			r = l.Right
			pstate.pushtype(r, t.Elem(pstate.types))
			r = pstate.typecheck(r, Erv)
			r = pstate.defaultlit(r, t.Elem(pstate.types))
			l.Right = pstate.assignconv(r, t.Elem(pstate.types), "map value")
		}

		n.Op = OMAPLIT
		n.Right = nil

	case TSTRUCT:
		// Need valid field offsets for Xoffset below.
		pstate.dowidth(t)

		errored := false
		if n.List.Len() != 0 && nokeys(n.List) {
			// simple list of variables
			ls := n.List.Slice()
			for i, n1 := range ls {
				pstate.setlineno(n1)
				n1 = pstate.typecheck(n1, Erv)
				ls[i] = n1
				if i >= t.NumFields(pstate.types) {
					if !errored {
						pstate.yyerror("too many values in %v", n)
						errored = true
					}
					continue
				}

				f := t.Field(pstate.types, i)
				s := f.Sym
				if s != nil && !types.IsExported(s.Name) && s.Pkg != pstate.localpkg {
					pstate.yyerror("implicit assignment of unexported field '%s' in %v literal", s.Name, t)
				}
				// No pushtype allowed here. Must name fields for that.
				n1 = pstate.assignconv(n1, f.Type, "field value")
				n1 = pstate.nodSym(OSTRUCTKEY, n1, f.Sym)
				n1.Xoffset = f.Offset
				ls[i] = n1
			}
			if len(ls) < t.NumFields(pstate.types) {
				pstate.yyerror("too few values in %v", n)
			}
		} else {
			hash := make(map[string]bool)

			// keyed list
			ls := n.List.Slice()
			for i, l := range ls {
				pstate.setlineno(l)

				if l.Op == OKEY {
					key := l.Left

					l.Op = OSTRUCTKEY
					l.Left = l.Right
					l.Right = nil

					// An OXDOT uses the Sym field to hold
					// the field to the right of the dot,
					// so s will be non-nil, but an OXDOT
					// is never a valid struct literal key.
					if key.Sym == nil || key.Op == OXDOT || key.Sym.IsBlank() {
						pstate.yyerror("invalid field name %v in struct initializer", key)
						l.Left = pstate.typecheck(l.Left, Erv)
						continue
					}

					// Sym might have resolved to name in other top-level
					// package, because of import dot. Redirect to correct sym
					// before we do the lookup.
					s := key.Sym
					if s.Pkg != pstate.localpkg && types.IsExported(s.Name) {
						s1 := pstate.lookup(s.Name)
						if s1.Origpkg == s.Pkg {
							s = s1
						}
					}
					l.Sym = s
				}

				if l.Op != OSTRUCTKEY {
					if !errored {
						pstate.yyerror("mixture of field:value and value initializers")
						errored = true
					}
					ls[i] = pstate.typecheck(ls[i], Erv)
					continue
				}

				f := pstate.lookdot1(nil, l.Sym, t, t.Fields(pstate.types), 0)
				if f == nil {
					if ci := pstate.lookdot1(nil, l.Sym, t, t.Fields(pstate.types), 2); ci != nil { // Case-insensitive lookup.
						if pstate.visible(ci.Sym) {
							pstate.yyerror("unknown field '%v' in struct literal of type %v (but does have %v)", l.Sym, t, ci.Sym)
						} else {
							pstate.yyerror("unknown field '%v' in struct literal of type %v", l.Sym, t)
						}
						continue
					}
					p, _ := pstate.dotpath(l.Sym, t, nil, true)
					if p == nil {
						pstate.yyerror("unknown field '%v' in struct literal of type %v", l.Sym, t)
						continue
					}
					// dotpath returns the parent embedded types in reverse order.
					var ep []string
					for ei := len(p) - 1; ei >= 0; ei-- {
						ep = append(ep, p[ei].field.Type.Sym.Name)
					}
					ep = append(ep, l.Sym.Name)
					pstate.yyerror("cannot use promoted field %v in struct literal of type %v", strings.Join(ep, "."), t)
					continue
				}
				pstate.fielddup(f.Sym.Name, hash)
				l.Xoffset = f.Offset

				// No pushtype allowed here. Tried and rejected.
				l.Left = pstate.typecheck(l.Left, Erv)
				l.Left = pstate.assignconv(l.Left, f.Type, "field value")
			}
		}

		n.Op = OSTRUCTLIT
		n.Right = nil
	}

	if nerr != pstate.nerrors {
		return n
	}

	n.Orig = norig
	if n.Type.IsPtr() {
		n = pstate.nod(OPTRLIT, n, nil)
		n.SetTypecheck(1)
		n.Type = n.Left.Type
		n.Left.Type = t
		n.Left.SetTypecheck(1)
	}

	n.Orig = norig
	return n
}

// visible reports whether sym is exported or locally defined.
func (pstate *PackageState) visible(sym *types.Sym) bool {
	return sym != nil && (types.IsExported(sym.Name) || sym.Pkg == pstate.localpkg)
}

// lvalue etc
func islvalue(n *Node) bool {
	switch n.Op {
	case OINDEX:
		if n.Left.Type != nil && n.Left.Type.IsArray() {
			return islvalue(n.Left)
		}
		if n.Left.Type != nil && n.Left.Type.IsString() {
			return false
		}
		fallthrough
	case OIND, ODOTPTR, OCLOSUREVAR:
		return true

	case ODOT:
		return islvalue(n.Left)

	case ONAME:
		if n.Class() == PFUNC {
			return false
		}
		return true
	}

	return false
}

func (pstate *PackageState) checklvalue(n *Node, verb string) {
	if !islvalue(n) {
		pstate.yyerror("cannot %s %v", verb, n)
	}
}

func (pstate *PackageState) checkassign(stmt *Node, n *Node) {
	// Variables declared in ORANGE are assigned on every iteration.
	if n.Name == nil || n.Name.Defn != stmt || stmt.Op == ORANGE {
		r := pstate.outervalue(n)
		var l *Node
		for l = n; l != r; l = l.Left {
			l.SetAssigned(true)
			if l.IsClosureVar() {
				l.Name.Defn.SetAssigned(true)
			}
		}

		l.SetAssigned(true)
		if l.IsClosureVar() {
			l.Name.Defn.SetAssigned(true)
		}
	}

	if islvalue(n) {
		return
	}
	if n.Op == OINDEXMAP {
		n.SetIndexMapLValue(pstate, true)
		return
	}

	// have already complained about n being invalid
	if n.Type == nil {
		return
	}

	if n.Op == ODOT && n.Left.Op == OINDEXMAP {
		pstate.yyerror("cannot assign to struct field %v in map", n)
	} else {
		pstate.yyerror("cannot assign to %v", n)
	}
	n.Type = nil
}

func (pstate *PackageState) checkassignlist(stmt *Node, l Nodes) {
	for _, n := range l.Slice() {
		pstate.checkassign(stmt, n)
	}
}

// samesafeexpr checks whether it is safe to reuse one of l and r
// instead of computing both. samesafeexpr assumes that l and r are
// used in the same statement or expression. In order for it to be
// safe to reuse l or r, they must:
// * be the same expression
// * not have side-effects (no function calls, no channel ops);
//   however, panics are ok
// * not cause inappropriate aliasing; e.g. two string to []byte
//   conversions, must result in two distinct slices
//
// The handling of OINDEXMAP is subtle. OINDEXMAP can occur both
// as an lvalue (map assignment) and an rvalue (map access). This is
// currently OK, since the only place samesafeexpr gets used on an
// lvalue expression is for OSLICE and OAPPEND optimizations, and it
// is correct in those settings.
func (pstate *PackageState) samesafeexpr(l *Node, r *Node) bool {
	if l.Op != r.Op || !pstate.eqtype(l.Type, r.Type) {
		return false
	}

	switch l.Op {
	case ONAME, OCLOSUREVAR:
		return l == r

	case ODOT, ODOTPTR:
		return l.Sym != nil && r.Sym != nil && l.Sym == r.Sym && pstate.samesafeexpr(l.Left, r.Left)

	case OIND, OCONVNOP:
		return pstate.samesafeexpr(l.Left, r.Left)

	case OCONV:
		// Some conversions can't be reused, such as []byte(str).
		// Allow only numeric-ish types. This is a bit conservative.
		return pstate.issimple[l.Type.Etype] && pstate.samesafeexpr(l.Left, r.Left)

	case OINDEX, OINDEXMAP:
		return pstate.samesafeexpr(l.Left, r.Left) && pstate.samesafeexpr(l.Right, r.Right)

	case OLITERAL:
		return pstate.eqval(l.Val(), r.Val())
	}

	return false
}

// type check assignment.
// if this assignment is the definition of a var on the left side,
// fill in the var's type.
func (pstate *PackageState) typecheckas(n *Node) {
	// delicate little dance.
	// the definition of n may refer to this assignment
	// as its definition, in which case it will call typecheckas.
	// in that case, do not call typecheck back, or it will cycle.
	// if the variable has a type (ntype) then typechecking
	// will not look at defn, so it is okay (and desirable,
	// so that the conversion below happens).
	n.Left = pstate.resolve(n.Left)

	if n.Left.Name == nil || n.Left.Name.Defn != n || n.Left.Name.Param.Ntype != nil {
		n.Left = pstate.typecheck(n.Left, Erv|Easgn)
	}

	n.Right = pstate.typecheck(n.Right, Erv)
	pstate.checkassign(n, n.Left)
	if n.Right != nil && n.Right.Type != nil {
		if n.Left.Type != nil {
			n.Right = pstate.assignconv(n.Right, n.Left.Type, "assignment")
		}
	}

	if n.Left.Name != nil && n.Left.Name.Defn == n && n.Left.Name.Param.Ntype == nil {
		n.Right = pstate.defaultlit(n.Right, nil)
		n.Left.Type = n.Right.Type
	}

	// second half of dance.
	// now that right is done, typecheck the left
	// just to get it over with.  see dance above.
	n.SetTypecheck(1)

	if n.Left.Typecheck() == 0 {
		n.Left = pstate.typecheck(n.Left, Erv|Easgn)
	}
	if !n.Left.isBlank() {
		pstate.checkwidth(n.Left.Type) // ensure width is calculated for backend
	}
}

func (pstate *PackageState) checkassignto(src *types.Type, dst *Node) {
	var why string

	if pstate.assignop(src, dst.Type, &why) == 0 {
		pstate.yyerror("cannot assign %v to %L in multiple assignment%s", src, dst, why)
		return
	}
}

func (pstate *PackageState) typecheckas2(n *Node) {
	ls := n.List.Slice()
	for i1, n1 := range ls {
		// delicate little dance.
		n1 = pstate.resolve(n1)
		ls[i1] = n1

		if n1.Name == nil || n1.Name.Defn != n || n1.Name.Param.Ntype != nil {
			ls[i1] = pstate.typecheck(ls[i1], Erv|Easgn)
		}
	}

	cl := n.List.Len()
	cr := n.Rlist.Len()
	if cl > 1 && cr == 1 {
		n.Rlist.SetFirst(pstate.typecheck(n.Rlist.First(), Erv|Efnstruct))
	} else {
		pstate.typecheckslice(n.Rlist.Slice(), Erv)
	}
	pstate.checkassignlist(n, n.List)

	var l *Node
	var r *Node
	if cl == cr {
		// easy
		ls := n.List.Slice()
		rs := n.Rlist.Slice()
		for il, nl := range ls {
			nr := rs[il]
			if nl.Type != nil && nr.Type != nil {
				rs[il] = pstate.assignconv(nr, nl.Type, "assignment")
			}
			if nl.Name != nil && nl.Name.Defn == n && nl.Name.Param.Ntype == nil {
				rs[il] = pstate.defaultlit(rs[il], nil)
				nl.Type = rs[il].Type
			}
		}

		goto out
	}

	l = n.List.First()
	r = n.Rlist.First()

	// x,y,z = f()
	if cr == 1 {
		if r.Type == nil {
			goto out
		}
		switch r.Op {
		case OCALLMETH, OCALLINTER, OCALLFUNC:
			if !r.Type.IsFuncArgStruct() {
				break
			}
			cr = r.Type.NumFields(pstate.types)
			if cr != cl {
				goto mismatch
			}
			n.Op = OAS2FUNC
			for i, l := range n.List.Slice() {
				f := r.Type.Field(pstate.types, i)
				if f.Type != nil && l.Type != nil {
					pstate.checkassignto(f.Type, l)
				}
				if l.Name != nil && l.Name.Defn == n && l.Name.Param.Ntype == nil {
					l.Type = f.Type
				}
			}
			goto out
		}
	}

	// x, ok = y
	if cl == 2 && cr == 1 {
		if r.Type == nil {
			goto out
		}
		switch r.Op {
		case OINDEXMAP, ORECV, ODOTTYPE:
			switch r.Op {
			case OINDEXMAP:
				n.Op = OAS2MAPR

			case ORECV:
				n.Op = OAS2RECV

			case ODOTTYPE:
				n.Op = OAS2DOTTYPE
				r.Op = ODOTTYPE2
			}

			if l.Type != nil {
				pstate.checkassignto(r.Type, l)
			}
			if l.Name != nil && l.Name.Defn == n {
				l.Type = r.Type
			}
			l := n.List.Second()
			if l.Type != nil && !l.Type.IsBoolean() {
				pstate.checkassignto(pstate.types.Types[TBOOL], l)
			}
			if l.Name != nil && l.Name.Defn == n && l.Name.Param.Ntype == nil {
				l.Type = pstate.types.Types[TBOOL]
			}
			goto out
		}
	}

mismatch:
	pstate.yyerror("assignment mismatch: %d variables but %d values", cl, cr)

	// second half of dance
out:
	n.SetTypecheck(1)
	ls = n.List.Slice()
	for i1, n1 := range ls {
		if n1.Typecheck() == 0 {
			ls[i1] = pstate.typecheck(ls[i1], Erv|Easgn)
		}
	}
}

// type check function definition
func (pstate *PackageState) typecheckfunc(n *Node) {
	for _, ln := range n.Func.Dcl {
		if ln.Op == ONAME && (ln.Class() == PPARAM || ln.Class() == PPARAMOUT) {
			ln.Name.Decldepth = 1
		}
	}

	n.Func.Nname = pstate.typecheck(n.Func.Nname, Erv|Easgn)
	t := n.Func.Nname.Type
	if t == nil {
		return
	}
	n.Type = t
	t.FuncType(pstate.types).Nname = asTypesNode(n.Func.Nname)
	rcvr := t.Recv(pstate.types)
	if rcvr != nil && n.Func.Shortname != nil {
		m := pstate.addmethod(n.Func.Shortname, t, true, n.Func.Pragma&Nointerface != 0)
		if m == nil {
			return
		}

		n.Func.Nname.Sym = pstate.methodSym(rcvr.Type, n.Func.Shortname)
		pstate.declare(n.Func.Nname, PFUNC)
	}

	if pstate.Ctxt.Flag_dynlink && !pstate.inimport && n.Func.Nname != nil {
		pstate.makefuncsym(n.Func.Nname.Sym)
	}
}

// The result of stringtoarraylit MUST be assigned back to n, e.g.
// 	n.Left = stringtoarraylit(n.Left)
func (pstate *PackageState) stringtoarraylit(n *Node) *Node {
	if n.Left.Op != OLITERAL || n.Left.Val().Ctype(pstate) != CTSTR {
		pstate.Fatalf("stringtoarraylit %v", n)
	}

	s := n.Left.Val().U.(string)
	var l []*Node
	if n.Type.Elem(pstate.types).Etype == TUINT8 {
		// []byte
		for i := 0; i < len(s); i++ {
			l = append(l, pstate.nod(OKEY, pstate.nodintconst(int64(i)), pstate.nodintconst(int64(s[0]))))
		}
	} else {
		// []rune
		i := 0
		for _, r := range s {
			l = append(l, pstate.nod(OKEY, pstate.nodintconst(int64(i)), pstate.nodintconst(int64(r))))
			i++
		}
	}

	nn := pstate.nod(OCOMPLIT, nil, pstate.typenod(n.Type))
	nn.List.Set(l)
	nn = pstate.typecheck(nn, Erv)
	return nn
}

func (pstate *PackageState) checkMapKeys() {
	for _, n := range pstate.mapqueue {
		k := n.Type.MapType(pstate.types).Key
		if !k.Broke() && !pstate.IsComparable(k) {
			pstate.yyerrorl(n.Pos, "invalid map key type %v", k)
		}
	}
	pstate.mapqueue = nil
}

func (pstate *PackageState) copytype(n *Node, t *types.Type) {
	if t.Etype == TFORW {
		// This type isn't computed yet; when it is, update n.
		t.ForwardType(pstate.types).Copyto = append(t.ForwardType(pstate.types).Copyto, asTypesNode(n))
		return
	}

	embedlineno := n.Type.ForwardType(pstate.types).Embedlineno
	l := n.Type.ForwardType(pstate.types).Copyto

	ptrBase := n.Type.PtrBase
	sliceOf := n.Type.SliceOf

	// TODO(mdempsky): Fix Type rekinding.
	*n.Type = *t

	t = n.Type
	t.Sym = n.Sym
	if n.Name != nil {
		t.Vargen = n.Name.Vargen
	}

	// spec: "The declared type does not inherit any methods bound
	// to the existing type, but the method set of an interface
	// type [...] remains unchanged."
	if !t.IsInterface() {
		*t.Methods() = types.Fields{}
		*t.AllMethods() = types.Fields{}
	}

	t.Nod = asTypesNode(n)
	t.SetDeferwidth(false)
	t.PtrBase = ptrBase
	t.SliceOf = sliceOf

	// Propagate go:notinheap pragma from the Name to the Type.
	if n.Name != nil && n.Name.Param != nil && n.Name.Param.Pragma&NotInHeap != 0 {
		t.SetNotInHeap(true)
	}

	// Update nodes waiting on this type.
	for _, n := range l {
		pstate.copytype(asNode(n), t)
	}

	// Double-check use of type as embedded type.
	lno := pstate.lineno

	if embedlineno.IsKnown() {
		pstate.lineno = embedlineno
		if t.IsPtr() || t.IsUnsafePtr() {
			pstate.yyerror("embedded type cannot be a pointer")
		}
	}

	pstate.lineno = lno
}

func (pstate *PackageState) typecheckdeftype(n *Node) {
	lno := pstate.lineno
	pstate.setlineno(n)
	n.Type.Sym = n.Sym
	n.SetTypecheck(1)
	n.Name.Param.Ntype = pstate.typecheck(n.Name.Param.Ntype, Etype)
	t := n.Name.Param.Ntype.Type
	if t == nil {
		n.SetDiag(true)
		n.Type = nil
	} else if n.Type == nil {
		n.SetDiag(true)
	} else {
		// copy new type and clear fields
		// that don't come along.
		pstate.copytype(n, t)
	}

	pstate.lineno = lno
}

func (pstate *PackageState) typecheckdef(n *Node) {
	lno := pstate.lineno
	pstate.setlineno(n)

	if n.Op == ONONAME {
		if !n.Diag() {
			n.SetDiag(true)
			if n.Pos.IsKnown() {
				pstate.lineno = n.Pos
			}

			// Note: adderrorname looks for this string and
			// adds context about the outer expression
			pstate.yyerror("undefined: %v", n.Sym)
		}

		return
	}

	if n.Walkdef() == 1 {
		return
	}

	pstate.typecheckdefstack = append(pstate.typecheckdefstack, n)
	if n.Walkdef() == 2 {
		pstate.flusherrors()
		fmt.Printf("typecheckdef loop:")
		for i := len(pstate.typecheckdefstack) - 1; i >= 0; i-- {
			n := pstate.typecheckdefstack[i]
			fmt.Printf(" %v", n.Sym)
		}
		fmt.Printf("\n")
		pstate.Fatalf("typecheckdef loop")
	}

	n.SetWalkdef(2)

	if n.Type != nil || n.Sym == nil { // builtin or no name
		goto ret
	}

	switch n.Op {
	default:
		pstate.Fatalf("typecheckdef %v", n.Op)

	case OGOTO, OLABEL, OPACK:
	// nothing to do here

	case OLITERAL:
		if n.Name.Param.Ntype != nil {
			n.Name.Param.Ntype = pstate.typecheck(n.Name.Param.Ntype, Etype)
			n.Type = n.Name.Param.Ntype.Type
			n.Name.Param.Ntype = nil
			if n.Type == nil {
				n.SetDiag(true)
				goto ret
			}
		}

		e := n.Name.Defn
		n.Name.Defn = nil
		if e == nil {
			pstate.lineno = n.Pos
			Dump("typecheckdef nil defn", n)
			pstate.yyerror("xxx")
		}

		e = pstate.typecheck(e, Erv)
		if pstate.Isconst(e, CTNIL) {
			pstate.yyerror("const initializer cannot be nil")
			goto ret
		}

		if e.Type != nil && e.Op != OLITERAL || !e.isGoConst(pstate) {
			if !e.Diag() {
				pstate.yyerror("const initializer %v is not a constant", e)
				e.SetDiag(true)
			}

			goto ret
		}

		t := n.Type
		if t != nil {
			if !pstate.okforconst[t.Etype] {
				pstate.yyerror("invalid constant type %v", t)
				goto ret
			}

			if !e.Type.IsUntyped(pstate.types) && !pstate.eqtype(t, e.Type) {
				pstate.yyerror("cannot use %L as type %v in const initializer", e, t)
				goto ret
			}

			e = pstate.convlit(e, t)
		}

		n.SetVal(pstate, e.Val())
		n.Type = e.Type

	case ONAME:
		if n.Name.Param.Ntype != nil {
			n.Name.Param.Ntype = pstate.typecheck(n.Name.Param.Ntype, Etype)
			n.Type = n.Name.Param.Ntype.Type
			if n.Type == nil {
				n.SetDiag(true)
				goto ret
			}
		}

		if n.Type != nil {
			break
		}
		if n.Name.Defn == nil {
			if n.SubOp(pstate) != 0 { // like OPRINTN
				break
			}
			if pstate.nsavederrors+pstate.nerrors > 0 {
				// Can have undefined variables in x := foo
				// that make x have an n.name.Defn == nil.
				// If there are other errors anyway, don't
				// bother adding to the noise.
				break
			}

			pstate.Fatalf("var without type, init: %v", n.Sym)
		}

		if n.Name.Defn.Op == ONAME {
			n.Name.Defn = pstate.typecheck(n.Name.Defn, Erv)
			n.Type = n.Name.Defn.Type
			break
		}

		n.Name.Defn = pstate.typecheck(n.Name.Defn, Etop) // fills in n.Type

	case OTYPE:
		if p := n.Name.Param; p.Alias {
			// Type alias declaration: Simply use the rhs type - no need
			// to create a new type.
			// If we have a syntax error, p.Ntype may be nil.
			if p.Ntype != nil {
				p.Ntype = pstate.typecheck(p.Ntype, Etype)
				n.Type = p.Ntype.Type
				if n.Type == nil {
					n.SetDiag(true)
					goto ret
				}
				n.Sym.Def = asTypesNode(p.Ntype)
			}
			break
		}

		// regular type declaration
		if pstate.Curfn != nil {
			pstate.defercheckwidth()
		}
		n.SetWalkdef(1)
		n.Type = types.New(TFORW)
		n.Type.Nod = asTypesNode(n)
		n.Type.Sym = n.Sym // TODO(gri) this also happens in typecheckdeftype(n) - where should it happen?
		nerrors0 := pstate.nerrors
		pstate.typecheckdeftype(n)
		if n.Type.Etype == TFORW && pstate.nerrors > nerrors0 {
			// Something went wrong during type-checking,
			// but it was reported. Silence future errors.
			n.Type.SetBroke(true)
		}
		if pstate.Curfn != nil {
			pstate.resumecheckwidth()
		}
	}

ret:
	if n.Op != OLITERAL && n.Type != nil && n.Type.IsUntyped(pstate.types) {
		pstate.Fatalf("got %v for %v", n.Type, n)
	}
	last := len(pstate.typecheckdefstack) - 1
	if pstate.typecheckdefstack[last] != n {
		pstate.Fatalf("typecheckdefstack mismatch")
	}
	pstate.typecheckdefstack[last] = nil
	pstate.typecheckdefstack = pstate.typecheckdefstack[:last]

	pstate.lineno = lno
	n.SetWalkdef(1)
}

func (pstate *PackageState) checkmake(t *types.Type, arg string, n *Node) bool {
	if !n.Type.IsInteger() && n.Type.Etype != TIDEAL {
		pstate.yyerror("non-integer %s argument in make(%v) - %v", arg, t, n.Type)
		return false
	}

	// Do range checks for constants before defaultlit
	// to avoid redundant "constant NNN overflows int" errors.
	switch pstate.consttype(n) {
	case CTINT, CTRUNE, CTFLT, CTCPLX:
		n.SetVal(pstate, pstate.toint(n.Val()))
		if n.Val().U.(*Mpint).CmpInt64(0) < 0 {
			pstate.yyerror("negative %s argument in make(%v)", arg, t)
			return false
		}
		if n.Val().U.(*Mpint).Cmp(pstate.maxintval[TINT]) > 0 {
			pstate.yyerror("%s argument too large in make(%v)", arg, t)
			return false
		}
	}

	// defaultlit is necessary for non-constants too: n might be 1.1<<k.
	// TODO(gri) The length argument requirements for (array/slice) make
	// are the same as for index expressions. Factor the code better;
	// for instance, indexlit might be called here and incorporate some
	// of the bounds checks done for make.
	n = pstate.defaultlit(n, pstate.types.Types[TINT])

	return true
}

func markbreak(n *Node, implicit *Node) {
	if n == nil {
		return
	}

	switch n.Op {
	case OBREAK:
		if n.Left == nil {
			if implicit != nil {
				implicit.SetHasBreak(true)
			}
		} else {
			lab := asNode(n.Left.Sym.Label)
			if lab != nil {
				lab.SetHasBreak(true)
			}
		}
	case OFOR, OFORUNTIL, OSWITCH, OTYPESW, OSELECT, ORANGE:
		implicit = n
		fallthrough
	default:
		markbreak(n.Left, implicit)
		markbreak(n.Right, implicit)
		markbreaklist(n.Ninit, implicit)
		markbreaklist(n.Nbody, implicit)
		markbreaklist(n.List, implicit)
		markbreaklist(n.Rlist, implicit)
	}
}

func markbreaklist(l Nodes, implicit *Node) {
	s := l.Slice()
	for i := 0; i < len(s); i++ {
		n := s[i]
		if n == nil {
			continue
		}
		if n.Op == OLABEL && i+1 < len(s) && n.Name.Defn == s[i+1] {
			switch n.Name.Defn.Op {
			case OFOR, OFORUNTIL, OSWITCH, OTYPESW, OSELECT, ORANGE:
				n.Left.Sym.Label = asTypesNode(n.Name.Defn)
				markbreak(n.Name.Defn, n.Name.Defn)
				n.Left.Sym.Label = nil
				i++
				continue
			}
		}

		markbreak(n, implicit)
	}
}

// isterminating reports whether the Nodes list ends with a terminating statement.
func (l Nodes) isterminating() bool {
	s := l.Slice()
	c := len(s)
	if c == 0 {
		return false
	}
	return s[c-1].isterminating()
}

// Isterminating reports whether the node n, the last one in a
// statement list, is a terminating statement.
func (n *Node) isterminating() bool {
	switch n.Op {
	// NOTE: OLABEL is treated as a separate statement,
	// not a separate prefix, so skipping to the last statement
	// in the block handles the labeled statement case by
	// skipping over the label. No case OLABEL here.

	case OBLOCK:
		return n.List.isterminating()

	case OGOTO, ORETURN, ORETJMP, OPANIC, OFALL:
		return true

	case OFOR, OFORUNTIL:
		if n.Left != nil {
			return false
		}
		if n.HasBreak() {
			return false
		}
		return true

	case OIF:
		return n.Nbody.isterminating() && n.Rlist.isterminating()

	case OSWITCH, OTYPESW, OSELECT:
		if n.HasBreak() {
			return false
		}
		def := false
		for _, n1 := range n.List.Slice() {
			if !n1.Nbody.isterminating() {
				return false
			}
			if n1.List.Len() == 0 { // default
				def = true
			}
		}

		if n.Op != OSELECT && !def {
			return false
		}
		return true
	}

	return false
}

// checkreturn makes sure that fn terminates appropriately.
func (pstate *PackageState) checkreturn(fn *Node) {
	if fn.Type.NumResults(pstate.types) != 0 && fn.Nbody.Len() != 0 {
		markbreaklist(fn.Nbody, nil)
		if !fn.Nbody.isterminating() {
			pstate.yyerrorl(fn.Func.Endlineno, "missing return at end of function")
		}
	}
}

func (pstate *PackageState) deadcode(fn *Node) {
	pstate.deadcodeslice(fn.Nbody)
}

func (pstate *PackageState) deadcodeslice(nn Nodes) {
	for i, n := range nn.Slice() {
		// Cut is set to true when all nodes after i'th position
		// should be removed.
		// In other words, it marks whole slice "tail" as dead.
		cut := false
		if n == nil {
			continue
		}
		if n.Op == OIF {
			n.Left = pstate.deadcodeexpr(n.Left)
			if pstate.Isconst(n.Left, CTBOOL) {
				var body Nodes
				if n.Left.Bool(pstate) {
					n.Rlist = Nodes{}
					body = n.Nbody
				} else {
					n.Nbody = Nodes{}
					body = n.Rlist
				}
				// If "then" or "else" branch ends with panic or return statement,
				// it is safe to remove all statements after this node.
				// isterminating is not used to avoid goto-related complications.
				if body := body.Slice(); len(body) != 0 {
					switch body[(len(body) - 1)].Op {
					case ORETURN, ORETJMP, OPANIC:
						cut = true
					}
				}
			}
		}

		pstate.deadcodeslice(n.Ninit)
		pstate.deadcodeslice(n.Nbody)
		pstate.deadcodeslice(n.List)
		pstate.deadcodeslice(n.Rlist)
		if cut {
			*nn.slice = nn.Slice()[:i+1]
			break
		}
	}
}

func (pstate *PackageState) deadcodeexpr(n *Node) *Node {
	// Perform dead-code elimination on short-circuited boolean
	// expressions involving constants with the intent of
	// producing a constant 'if' condition.
	switch n.Op {
	case OANDAND:
		n.Left = pstate.deadcodeexpr(n.Left)
		n.Right = pstate.deadcodeexpr(n.Right)
		if pstate.Isconst(n.Left, CTBOOL) {
			if n.Left.Bool(pstate) {
				return n.Right // true && x => x
			} else {
				return n.Left // false && x => false
			}
		}
	case OOROR:
		n.Left = pstate.deadcodeexpr(n.Left)
		n.Right = pstate.deadcodeexpr(n.Right)
		if pstate.Isconst(n.Left, CTBOOL) {
			if n.Left.Bool(pstate) {
				return n.Left // true || x => true
			} else {
				return n.Right // false || x => x
			}
		}
	}
	return n
}

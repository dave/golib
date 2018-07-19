package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"

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
func (psess *PackageSession) resolve(n *Node) *Node {
	if n == nil || n.Op != ONONAME {
		return n
	}

	if n.Sym.Pkg != psess.localpkg {
		if psess.inimport {
			psess.
				Fatalf("recursive inimport")
		}
		psess.
			inimport = true
		psess.
			expandDecl(n)
		psess.
			inimport = false
		return n
	}

	r := asNode(n.Sym.Def)
	if r == nil {
		return n
	}

	if r.Op == OIOTA {
		if i := len(psess.typecheckdefstack); i > 0 {
			if x := psess.typecheckdefstack[i-1]; x.Op == OLITERAL {
				return psess.nodintconst(x.Iota())
			}
		}
		return n
	}

	return r
}

func (psess *PackageSession) typecheckslice(l []*Node, top int) {
	for i := range l {
		l[i] = psess.typecheck(l[i], top)
	}
}

func (psess *PackageSession) typekind(t *types.Type) string {
	if t.IsSlice() {
		return "slice"
	}
	et := t.Etype
	if int(et) < len(psess._typekind) {
		s := psess._typekind[et]
		if s != "" {
			return s
		}
	}
	return fmt.Sprintf("etype=%d", et)
}

func (psess *PackageSession) cycleFor(start *Node) []*Node {

	i := len(psess.typecheck_tcstack) - 1
	for i > 0 && psess.typecheck_tcstack[i] != start {
		i--
	}

	// collect all nodes with same Op
	var cycle []*Node
	for _, n := range psess.typecheck_tcstack[i:] {
		if n.Op == start.Op {
			cycle = append(cycle, n)
		}
	}

	return cycle
}

func (psess *PackageSession) cycleTrace(cycle []*Node) string {
	var s string
	for i, n := range cycle {
		s += fmt.Sprintf("\n\t%v: %v uses %v", n.Line(psess), n, cycle[(i+1)%len(cycle)])
	}
	return s
}

// typecheck type checks node n.
// The result of typecheck MUST be assigned back to n, e.g.
// 	n.Left = typecheck(n.Left, top)
func (psess *PackageSession) typecheck(n *Node, top int) *Node {

	if !psess.typecheckok {
		psess.
			Fatalf("early typecheck")
	}

	if n == nil {
		return nil
	}

	lno := psess.setlineno(n)

	for n.Op == OPAREN {
		n = n.Left
	}

	n = psess.resolve(n)

	if n.Typecheck() == 1 {
		switch n.Op {
		case ONAME, OTYPE, OLITERAL, OPACK:
			break

		default:
			psess.
				lineno = lno
			return n
		}
	}

	if n.Typecheck() == 2 {

		switch n.Op {

		case ONAME:
			if top&(Erv|Etype) == Etype {
				psess.
					yyerror("%v is not a type", n)
			}

		case OTYPE:

			if top&Etype == Etype {

				cycle := psess.cycleFor(n)
				for _, n := range cycle {
					if n.Name != nil && !n.Name.Param.Alias {
						psess.
							lineno = lno
						return n
					}
				}
				psess.
					yyerrorl(n.Pos, "invalid recursive type alias %v%s", n, psess.cycleTrace(cycle))
			}

		case OLITERAL:
			if top&(Erv|Etype) == Etype {
				psess.
					yyerror("%v is not a type", n)
				break
			}
			psess.
				yyerrorl(n.Pos, "constant definition loop%s", psess.cycleTrace(psess.cycleFor(n)))
		}

		if psess.nsavederrors+psess.nerrors == 0 {
			var trace string
			for i := len(psess.typecheck_tcstack) - 1; i >= 0; i-- {
				x := psess.typecheck_tcstack[i]
				trace += fmt.Sprintf("\n\t%v %v", x.Line(psess), x)
			}
			psess.
				yyerror("typechecking loop involving %v%s", n, trace)
		}
		psess.
			lineno = lno
		return n
	}

	n.SetTypecheck(2)
	psess.
		typecheck_tcstack = append(psess.typecheck_tcstack, n)
	n = psess.typecheck1(n, top)

	n.SetTypecheck(1)

	last := len(psess.typecheck_tcstack) - 1
	psess.
		typecheck_tcstack[last] = nil
	psess.
		typecheck_tcstack = psess.typecheck_tcstack[:last]
	psess.
		lineno = lno
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
func (psess *PackageSession) indexlit(n *Node) *Node {
	if n != nil && n.Type != nil && n.Type.Etype == TIDEAL {
		return psess.defaultlit(n, psess.types.Types[TINT])
	}
	return n
}

// The result of typecheck1 MUST be assigned back to n, e.g.
// 	n.Left = typecheck1(n.Left, top)
func (psess *PackageSession) typecheck1(n *Node, top int) *Node {
	switch n.Op {
	case OXDOT, ODOT, ODOTPTR, ODOTMETH, ODOTINTER, ORETJMP:

	default:
		if n.Sym != nil {
			if n.Op == ONAME && n.SubOp(psess) != 0 && top&Ecall == 0 {
				psess.
					yyerror("use of builtin %v not in function call", n.Sym)
				n.Type = nil
				return n
			}
			psess.
				typecheckdef(n)
			if n.Op == ONONAME {
				n.Type = nil
				return n
			}
		}
	}

	ok := 0
	switch n.Op {

	default:
		Dump("typecheck", n)
		psess.
			Fatalf("typecheck %v", n.Op)

	case OLITERAL:
		ok |= Erv

		if n.Type == nil && n.Val().Ctype(psess) == CTSTR {
			n.Type = psess.types.Idealstring
		}

	case ONONAME:
		ok |= Erv

	case ONAME:
		if n.Name.Decldepth == 0 {
			n.Name.Decldepth = psess.decldepth
		}
		if n.SubOp(psess) != 0 {
			ok |= Ecall
			break
		}

		if top&Easgn == 0 {

			if n.isBlank() {
				psess.
					yyerror("cannot use _ as value")
				n.Type = nil
				return n
			}

			n.Name.SetUsed(true)
		}

		ok |= Erv

	case OPACK:
		psess.
			yyerror("use of package %v without selector", n.Sym)
		n.Type = nil
		return n

	case ODDD:
		break

	case OTYPE:
		ok |= Etype

		if n.Type == nil {
			return n
		}

	case OTARRAY:
		ok |= Etype
		r := psess.typecheck(n.Right, Etype)
		if r.Type == nil {
			n.Type = nil
			return n
		}

		var t *types.Type
		if n.Left == nil {
			t = psess.types.NewSlice(r.Type)
		} else if n.Left.Op == ODDD {
			if top&Ecomplit == 0 {
				if !n.Diag() {
					n.SetDiag(true)
					psess.
						yyerror("use of [...] array outside of array literal")
				}
				n.Type = nil
				return n
			}
			t = types.NewDDDArray(r.Type)
		} else {
			n.Left = psess.indexlit(psess.typecheck(n.Left, Erv))
			l := n.Left
			if psess.consttype(l) != CTINT {
				switch {
				case l.Type == nil:

				case l.Type.IsInteger() && l.Op != OLITERAL:
					psess.
						yyerror("non-constant array bound %v", l)
				default:
					psess.
						yyerror("invalid array bound %v", l)
				}
				n.Type = nil
				return n
			}

			v := l.Val()
			if psess.doesoverflow(v, psess.types.Types[TINT]) {
				psess.
					yyerror("array bound is too large")
				n.Type = nil
				return n
			}

			bound := v.U.(*Mpint).Int64(psess)
			if bound < 0 {
				psess.
					yyerror("array bound must be non-negative")
				n.Type = nil
				return n
			}
			t = psess.types.NewArray(r.Type, bound)
		}

		n.Op = OTYPE
		n.Type = t
		n.Left = nil
		n.Right = nil
		if !t.IsDDDArray() {
			psess.
				checkwidth(t)
		}

	case OTMAP:
		ok |= Etype
		n.Left = psess.typecheck(n.Left, Etype)
		n.Right = psess.typecheck(n.Right, Etype)
		l := n.Left
		r := n.Right
		if l.Type == nil || r.Type == nil {
			n.Type = nil
			return n
		}
		if l.Type.NotInHeap() {
			psess.
				yyerror("go:notinheap map key not allowed")
		}
		if r.Type.NotInHeap() {
			psess.
				yyerror("go:notinheap map value not allowed")
		}
		n.Op = OTYPE
		n.Type = psess.types.NewMap(l.Type, r.Type)
		psess.
			mapqueue = append(psess.mapqueue, n)
		n.Left = nil
		n.Right = nil

	case OTCHAN:
		ok |= Etype
		n.Left = psess.typecheck(n.Left, Etype)
		l := n.Left
		if l.Type == nil {
			n.Type = nil
			return n
		}
		if l.Type.NotInHeap() {
			psess.
				yyerror("chan of go:notinheap type not allowed")
		}
		t := psess.types.NewChan(l.Type, n.TChanDir(psess))
		n.Op = OTYPE
		n.Type = t
		n.Left = nil
		n.ResetAux()

	case OTSTRUCT:
		ok |= Etype
		n.Op = OTYPE
		n.Type = psess.tostruct(n.List.Slice())
		if n.Type == nil || n.Type.Broke() {
			n.Type = nil
			return n
		}
		n.List.Set(nil)

	case OTINTER:
		ok |= Etype
		n.Op = OTYPE
		n.Type = psess.tointerface(n.List.Slice())
		if n.Type == nil {
			return n
		}

	case OTFUNC:
		ok |= Etype
		n.Op = OTYPE
		n.Type = psess.functype(n.Left, n.List.Slice(), n.Rlist.Slice())
		if n.Type == nil {
			return n
		}
		n.Left = nil
		n.List.Set(nil)
		n.Rlist.Set(nil)

	case OIND:
		n.Left = psess.typecheck(n.Left, Erv|Etype|top&Ecomplit)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if l.Op == OTYPE {
			ok |= Etype
			n.Op = OTYPE
			n.Type = psess.types.NewPtr(l.Type)

			if !l.Type.IsDDDArray() {
				psess.
					checkwidth(l.Type)
			}
			n.Left = nil
			break
		}

		if !t.IsPtr() {
			if top&(Erv|Etop) != 0 {
				psess.
					yyerror("invalid indirect of %L", n.Left)
				n.Type = nil
				return n
			}

			break
		}

		ok |= Erv
		n.Type = t.Elem(psess.types)

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
			n.Left = psess.typecheck(n.Left, Erv)
			n.Right = psess.typecheck(n.Right, Erv)
			l = n.Left
			r = n.Right
			psess.
				checkassign(n, n.Left)
			if l.Type == nil || r.Type == nil {
				n.Type = nil
				return n
			}
			if n.Implicit() && !psess.okforarith[l.Type.Etype] {
				psess.
					yyerror("invalid operation: %v (non-numeric type %v)", n, l.Type)
				n.Type = nil
				return n
			}

			op = n.SubOp(psess)
		} else {
			ok |= Erv
			n.Left = psess.typecheck(n.Left, Erv)
			n.Right = psess.typecheck(n.Right, Erv)
			l = n.Left
			r = n.Right
			if l.Type == nil || r.Type == nil {
				n.Type = nil
				return n
			}
			op = n.Op
		}
		if op == OLSH || op == ORSH {
			r = psess.defaultlit(r, psess.types.Types[TUINT])
			n.Right = r
			t := r.Type
			if !t.IsInteger() || t.IsSigned() {
				psess.
					yyerror("invalid operation: %v (shift count type %v, must be unsigned integer)", n, r.Type)
				n.Type = nil
				return n
			}

			t = l.Type
			if t != nil && t.Etype != TIDEAL && !t.IsInteger() {
				psess.
					yyerror("invalid operation: %v (shift of type %v)", n, t)
				n.Type = nil
				return n
			}

			n.Type = l.Type

			break
		}

		l, r = psess.defaultlit2(l, r, false)

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
		if psess.iscmp[n.Op] && t.Etype != TIDEAL && !psess.eqtype(l.Type, r.Type) {

			converted := false
			if r.Type.Etype != TBLANK {
				aop = psess.assignop(l.Type, r.Type, nil)
				if aop != 0 {
					if r.Type.IsInterface() && !l.Type.IsInterface() && !psess.IsComparable(l.Type) {
						psess.
							yyerror("invalid operation: %v (operator %v not defined on %s)", n, op, psess.typekind(l.Type))
						n.Type = nil
						return n
					}
					psess.
						dowidth(l.Type)
					if r.Type.IsInterface() == l.Type.IsInterface() || l.Type.Width >= 1<<16 {
						l = psess.nod(aop, l, nil)
						l.Type = r.Type
						l.SetTypecheck(1)
						n.Left = l
					}

					t = r.Type
					converted = true
				}
			}

			if !converted && l.Type.Etype != TBLANK {
				aop = psess.assignop(r.Type, l.Type, nil)
				if aop != 0 {
					if l.Type.IsInterface() && !r.Type.IsInterface() && !psess.IsComparable(r.Type) {
						psess.
							yyerror("invalid operation: %v (operator %v not defined on %s)", n, op, psess.typekind(r.Type))
						n.Type = nil
						return n
					}
					psess.
						dowidth(r.Type)
					if r.Type.IsInterface() == l.Type.IsInterface() || r.Type.Width >= 1<<16 {
						r = psess.nod(aop, r, nil)
						r.Type = l.Type
						r.SetTypecheck(1)
						n.Right = r
					}

					t = l.Type
				}
			}

			et = t.Etype
		}

		if t.Etype != TIDEAL && !psess.eqtype(l.Type, r.Type) {
			l, r = psess.defaultlit2(l, r, true)
			if r.Type.IsInterface() == l.Type.IsInterface() || aop == 0 {
				psess.
					yyerror("invalid operation: %v (mismatched types %v and %v)", n, l.Type, r.Type)
				n.Type = nil
				return n
			}
		}

		if !psess.okfor[op][et] {
			psess.
				yyerror("invalid operation: %v (operator %v not defined on %s)", n, op, psess.typekind(t))
			n.Type = nil
			return n
		}

		if l.Type.IsArray() && !psess.IsComparable(l.Type) {
			psess.
				yyerror("invalid operation: %v (%v cannot be compared)", n, l.Type)
			n.Type = nil
			return n
		}

		if l.Type.IsSlice() && !l.isNil(psess) && !r.isNil(psess) {
			psess.
				yyerror("invalid operation: %v (slice can only be compared to nil)", n)
			n.Type = nil
			return n
		}

		if l.Type.IsMap() && !l.isNil(psess) && !r.isNil(psess) {
			psess.
				yyerror("invalid operation: %v (map can only be compared to nil)", n)
			n.Type = nil
			return n
		}

		if l.Type.Etype == TFUNC && !l.isNil(psess) && !r.isNil(psess) {
			psess.
				yyerror("invalid operation: %v (func can only be compared to nil)", n)
			n.Type = nil
			return n
		}

		if l.Type.IsStruct() {
			if f := psess.IncomparableField(l.Type); f != nil {
				psess.
					yyerror("invalid operation: %v (struct containing %v cannot be compared)", n, f.Type)
				n.Type = nil
				return n
			}
		}

		t = l.Type
		if psess.iscmp[n.Op] {
			psess.
				evconst(n)
			t = psess.types.Idealbool
			if n.Op != OLITERAL {
				l, r = psess.defaultlit2(l, r, true)
				n.Left = l
				n.Right = r
			}
		}

		if et == TSTRING {
			if psess.iscmp[n.Op] {
				ot := n.Op
				n.Op = OCMPSTR
				n.SetSubOp(psess, ot)
			} else if n.Op == OADD {

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
			if l.Op == OLITERAL && l.Val().Ctype(psess) == CTNIL {

				n.Left = r

				n.Right = l
			} else if r.Op == OLITERAL && r.Val().Ctype(psess) == CTNIL {
			} else if r.Type.IsInterface() == l.Type.IsInterface() {
				ot := n.Op
				n.Op = OCMPIFACE
				n.SetSubOp(psess, ot)
			}
		}

		if (op == ODIV || op == OMOD) && psess.Isconst(r, CTINT) {
			if r.Val().U.(*Mpint).CmpInt64(0) == 0 {
				psess.
					yyerror("division by zero")
				n.Type = nil
				return n
			}
		}

		n.Type = t

	case OCOM, OMINUS, ONOT, OPLUS:
		ok |= Erv
		n.Left = psess.typecheck(n.Left, Erv)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !psess.okfor[n.Op][t.Etype] {
			psess.
				yyerror("invalid operation: %v %v", n.Op, t)
			n.Type = nil
			return n
		}

		n.Type = t

	case OADDR:
		ok |= Erv

		n.Left = psess.typecheck(n.Left, Erv)
		if n.Left.Type == nil {
			n.Type = nil
			return n
		}
		psess.
			checklvalue(n.Left, "take the address of")
		r := psess.outervalue(n.Left)
		var l *Node
		for l = n.Left; l != r; l = l.Left {
			l.SetAddrtaken(true)
			if l.IsClosureVar() && !psess.capturevarscomplete {

				l.Name.Defn.SetAddrtaken(true)
			}
		}

		if l.Orig != l && l.Op == ONAME {
			psess.
				Fatalf("found non-orig name node %v", l)
		}
		l.SetAddrtaken(true)
		if l.IsClosureVar() && !psess.capturevarscomplete {

			l.Name.Defn.SetAddrtaken(true)
		}
		n.Left = psess.defaultlit(n.Left, nil)
		l = n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		n.Type = psess.types.NewPtr(t)

	case OCOMPLIT:
		ok |= Erv
		n = psess.typecheckcomplit(n)
		if n.Type == nil {
			return n
		}

	case OXDOT, ODOT:
		if n.Op == OXDOT {
			n = psess.adddot(n)
			n.Op = ODOT
			if n.Left == nil {
				n.Type = nil
				return n
			}
		}

		n.Left = psess.typecheck(n.Left, Erv|Etype)

		n.Left = psess.defaultlit(n.Left, nil)

		t := n.Left.Type
		if t == nil {
			psess.
				adderrorname(n)
			n.Type = nil
			return n
		}

		s := n.Sym

		if n.Left.Op == OTYPE {
			n = psess.typecheckMethodExpr(n)
			if n.Type == nil {
				return n
			}
			ok = Erv
			break
		}

		if t.IsPtr() && !t.Elem(psess.types).IsInterface() {
			t = t.Elem(psess.types)
			if t == nil {
				n.Type = nil
				return n
			}
			n.Op = ODOTPTR
			psess.
				checkwidth(t)
		}

		if n.Sym.IsBlank() {
			psess.
				yyerror("cannot refer to blank field or method")
			n.Type = nil
			return n
		}

		if psess.lookdot(n, t, 0) == nil {

			switch {
			case t.IsEmptyInterface(psess.types):
				psess.
					yyerror("%v undefined (type %v is interface with no methods)", n, n.Left.Type)

			case t.IsPtr() && t.Elem(psess.types).IsInterface():
				psess.
					yyerror("%v undefined (type %v is pointer to interface, not interface)", n, n.Left.Type)

			case psess.lookdot(n, t, 1) != nil:
				psess.
					yyerror("%v undefined (cannot refer to unexported field or method %v)", n, n.Sym)

			default:
				if mt := psess.lookdot(n, t, 2); mt != nil && psess.visible(mt.Sym) {
					psess.
						yyerror("%v undefined (type %v has no field or method %v, but does have %v)", n, n.Left.Type, n.Sym, mt.Sym)
				} else {
					psess.
						yyerror("%v undefined (type %v has no field or method %v)", n, n.Left.Type, n.Sym)
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
				psess.
					typecheckpartialcall(n, s)
				ok |= Erv
			}

		default:
			ok |= Erv
		}

	case ODOTTYPE:
		ok |= Erv
		n.Left = psess.typecheck(n.Left, Erv)
		n.Left = psess.defaultlit(n.Left, nil)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsInterface() {
			psess.
				yyerror("invalid type assertion: %v (non-interface type %v on left)", n, t)
			n.Type = nil
			return n
		}

		if n.Right != nil {
			n.Right = psess.typecheck(n.Right, Etype)
			n.Type = n.Right.Type
			n.Right = nil
			if n.Type == nil {
				return n
			}
		}

		if n.Type != nil && !n.Type.IsInterface() {
			var missing, have *types.Field
			var ptr int
			if !psess.implements(n.Type, t, &missing, &have, &ptr) {
				if have != nil && have.Sym == missing.Sym {
					psess.
						yyerror("impossible type assertion:\n\t%v does not implement %v (wrong type for %v method)\n"+
							"\t\thave %v%0S\n\t\twant %v%0S", n.Type, t, missing.Sym, have.Sym, have.Type, missing.Sym, missing.Type)
				} else if ptr != 0 {
					psess.
						yyerror("impossible type assertion:\n\t%v does not implement %v (%v method has pointer receiver)", n.Type, t, missing.Sym)
				} else if have != nil {
					psess.
						yyerror("impossible type assertion:\n\t%v does not implement %v (missing %v method)\n"+
							"\t\thave %v%0S\n\t\twant %v%0S", n.Type, t, missing.Sym, have.Sym, have.Type, missing.Sym, missing.Type)
				} else {
					psess.
						yyerror("impossible type assertion:\n\t%v does not implement %v (missing %v method)", n.Type, t, missing.Sym)
				}
				n.Type = nil
				return n
			}
		}

	case OINDEX:
		ok |= Erv
		n.Left = psess.typecheck(n.Left, Erv)
		n.Left = psess.defaultlit(n.Left, nil)
		n.Left = psess.implicitstar(n.Left)
		l := n.Left
		n.Right = psess.typecheck(n.Right, Erv)
		r := n.Right
		t := l.Type
		if t == nil || r.Type == nil {
			n.Type = nil
			return n
		}
		switch t.Etype {
		default:
			psess.
				yyerror("invalid operation: %v (type %v does not support indexing)", n, t)
			n.Type = nil
			return n

		case TSTRING, TARRAY, TSLICE:
			n.Right = psess.indexlit(n.Right)
			if t.IsString() {
				n.Type = psess.types.Bytetype
			} else {
				n.Type = t.Elem(psess.types)
			}
			why := "string"
			if t.IsArray() {
				why = "array"
			} else if t.IsSlice() {
				why = "slice"
			}

			if n.Right.Type != nil && !n.Right.Type.IsInteger() {
				psess.
					yyerror("non-integer %s index %v", why, n.Right)
				break
			}

			if !n.Bounded() && psess.Isconst(n.Right, CTINT) {
				x := n.Right.Int64(psess)
				if x < 0 {
					psess.
						yyerror("invalid %s index %v (index must be non-negative)", why, n.Right)
				} else if t.IsArray() && x >= t.NumElem(psess.types) {
					psess.
						yyerror("invalid array index %v (out of bounds for %d-element array)", n.Right, t.NumElem(psess.types))
				} else if psess.Isconst(n.Left, CTSTR) && x >= int64(len(n.Left.Val().U.(string))) {
					psess.
						yyerror("invalid string index %v (out of bounds for %d-byte string)", n.Right, len(n.Left.Val().U.(string)))
				} else if n.Right.Val().U.(*Mpint).Cmp(psess.maxintval[TINT]) > 0 {
					psess.
						yyerror("invalid %s index %v (index too large)", why, n.Right)
				}
			}

		case TMAP:
			n.Right = psess.defaultlit(n.Right, t.Key(psess.types))
			if n.Right.Type != nil {
				n.Right = psess.assignconv(n.Right, t.Key(psess.types), "map index")
			}
			n.Type = t.Elem(psess.types)
			n.Op = OINDEXMAP
			n.ResetAux()
		}

	case ORECV:
		ok |= Etop | Erv
		n.Left = psess.typecheck(n.Left, Erv)
		n.Left = psess.defaultlit(n.Left, nil)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsChan() {
			psess.
				yyerror("invalid operation: %v (receive from non-chan type %v)", n, t)
			n.Type = nil
			return n
		}

		if !t.ChanDir(psess.types).CanRecv() {
			psess.
				yyerror("invalid operation: %v (receive from send-only type %v)", n, t)
			n.Type = nil
			return n
		}

		n.Type = t.Elem(psess.types)

	case OSEND:
		ok |= Etop
		n.Left = psess.typecheck(n.Left, Erv)
		n.Right = psess.typecheck(n.Right, Erv)
		n.Left = psess.defaultlit(n.Left, nil)
		t := n.Left.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsChan() {
			psess.
				yyerror("invalid operation: %v (send to non-chan type %v)", n, t)
			n.Type = nil
			return n
		}

		if !t.ChanDir(psess.types).CanSend() {
			psess.
				yyerror("invalid operation: %v (send to receive-only type %v)", n, t)
			n.Type = nil
			return n
		}

		n.Right = psess.defaultlit(n.Right, t.Elem(psess.types))
		r := n.Right
		if r.Type == nil {
			n.Type = nil
			return n
		}
		n.Right = psess.assignconv(r, t.Elem(psess.types), "send")
		n.Type = nil

	case OSLICE, OSLICE3:
		ok |= Erv
		n.Left = psess.typecheck(n.Left, Erv)
		low, high, max := n.SliceBounds(psess)
		hasmax := n.Op.IsSlice3(psess)
		low = psess.typecheck(low, Erv)
		high = psess.typecheck(high, Erv)
		max = psess.typecheck(max, Erv)
		n.Left = psess.defaultlit(n.Left, nil)
		low = psess.indexlit(low)
		high = psess.indexlit(high)
		max = psess.indexlit(max)
		n.SetSliceBounds(psess, low, high, max)
		l := n.Left
		if l.Type == nil {
			n.Type = nil
			return n
		}
		if l.Type.IsArray() {
			if !islvalue(n.Left) {
				psess.
					yyerror("invalid operation %v (slice of unaddressable value)", n)
				n.Type = nil
				return n
			}

			n.Left = psess.nod(OADDR, n.Left, nil)
			n.Left.SetImplicit(true)
			n.Left = psess.typecheck(n.Left, Erv)
			l = n.Left
		}
		t := l.Type
		var tp *types.Type
		if t.IsString() {
			if hasmax {
				psess.
					yyerror("invalid operation %v (3-index slice of string)", n)
				n.Type = nil
				return n
			}
			n.Type = t
			n.Op = OSLICESTR
		} else if t.IsPtr() && t.Elem(psess.types).IsArray() {
			tp = t.Elem(psess.types)
			n.Type = psess.types.NewSlice(tp.Elem(psess.types))
			psess.
				dowidth(n.Type)
			if hasmax {
				n.Op = OSLICE3ARR
			} else {
				n.Op = OSLICEARR
			}
		} else if t.IsSlice() {
			n.Type = t
		} else {
			psess.
				yyerror("cannot slice %v (type %v)", l, t)
			n.Type = nil
			return n
		}

		if low != nil && !psess.checksliceindex(l, low, tp) {
			n.Type = nil
			return n
		}
		if high != nil && !psess.checksliceindex(l, high, tp) {
			n.Type = nil
			return n
		}
		if max != nil && !psess.checksliceindex(l, max, tp) {
			n.Type = nil
			return n
		}
		if !psess.checksliceconst(low, high) || !psess.checksliceconst(low, max) || !psess.checksliceconst(high, max) {
			n.Type = nil
			return n
		}

	case OCALL:
		n.Left = psess.typecheck(n.Left, Erv|Etype|Ecall)
		if n.Left.Diag() {
			n.SetDiag(true)
		}

		l := n.Left

		if l.Op == ONAME && l.SubOp(psess) != 0 {
			if n.Isddd() && l.SubOp(psess) != OAPPEND {
				psess.
					yyerror("invalid use of ... with builtin %v", l)
			}

			n.Op = l.SubOp(psess)
			n.Left = n.Right
			n.Right = nil
			n = psess.typecheck1(n, top)
			return n
		}

		n.Left = psess.defaultlit(n.Left, nil)
		l = n.Left
		if l.Op == OTYPE {
			if n.Isddd() || l.Type.IsDDDArray() {
				if !l.Type.Broke() {
					psess.
						yyerror("invalid use of ... in type conversion to %v", l.Type)
				}
				n.SetDiag(true)
			}

			ok |= Erv

			n.Left = nil

			n.Op = OCONV
			n.Type = l.Type
			if !psess.onearg(n, "conversion to %v", l.Type) {
				n.Type = nil
				return n
			}
			n = psess.typecheck1(n, top)
			return n
		}

		if n.List.Len() == 1 && !n.Isddd() {
			n.List.SetFirst(psess.typecheck(n.List.First(), Erv|Efnstruct))
		} else {
			psess.
				typecheckslice(n.List.Slice(), Erv)
		}
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		psess.
			checkwidth(t)

		switch l.Op {
		case ODOTINTER:
			n.Op = OCALLINTER

		case ODOTMETH:
			n.Op = OCALLMETH

			tp := t.Recv(psess.types).Type

			if l.Left == nil || !psess.eqtype(l.Left.Type, tp) {
				psess.
					Fatalf("method receiver")
			}

		default:
			n.Op = OCALLFUNC
			if t.Etype != TFUNC {
				name := l.String()
				if psess.isBuiltinFuncName(name) {
					psess.
						yyerror("cannot call non-function %s (type %v), declared at %s",
							name, t, psess.linestr(l.Name.Defn.Pos))
				} else {
					psess.
						yyerror("cannot call non-function %s (type %v)", name, t)
				}
				n.Type = nil
				return n
			}
		}
		psess.
			typecheckaste(OCALL, n.Left, n.Isddd(), t.Params(psess.types), n.List, func() string { return fmt.Sprintf("argument to %v", n.Left) })
		ok |= Etop
		if t.NumResults(psess.types) == 0 {
			break
		}
		ok |= Erv
		if t.NumResults(psess.types) == 1 {
			n.Type = l.Type.Results(psess.types).Field(psess.types, 0).Type

			if n.Op == OCALLFUNC && n.Left.Op == ONAME && psess.isRuntimePkg(n.Left.Sym.Pkg) && n.Left.Sym.Name == "getg" {

				n.Op = OGETG
			}

			break
		}

		if top&(Efnstruct|Etop) == 0 {
			psess.
				yyerror("multiple-value %v() in single-value context", l)
			break
		}

		n.Type = l.Type.Results(psess.types)

	case OALIGNOF, OOFFSETOF, OSIZEOF:
		ok |= Erv
		if !psess.onearg(n, "%v", n.Op) {
			n.Type = nil
			return n
		}
		n.Type = psess.types.Types[TUINTPTR]
		psess.
			setintconst(n, psess.evalunsafe(n))

	case OCAP, OLEN:
		ok |= Erv
		if !psess.onearg(n, "%v", n.Op) {
			n.Type = nil
			return n
		}

		n.Left = psess.typecheck(n.Left, Erv)
		n.Left = psess.defaultlit(n.Left, nil)
		n.Left = psess.implicitstar(n.Left)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}

		var ok bool
		if n.Op == OLEN {
			ok = psess.okforlen[t.Etype]
		} else {
			ok = psess.okforcap[t.Etype]
		}
		if !ok {
			psess.
				yyerror("invalid argument %L for %v", l, n.Op)
			n.Type = nil
			return n
		}

		n.Type = psess.types.Types[TINT]

		// Result might be constant.
		var res int64 = -1 // valid if >= 0
		switch t.Etype {
		case TSTRING:
			if psess.Isconst(l, CTSTR) {
				res = int64(len(l.Val().U.(string)))
			}

		case TARRAY:
			if !callrecv(l) {
				res = t.NumElem(psess.types)
			}
		}
		if res >= 0 {
			psess.
				setintconst(n, res)
		}

	case OREAL, OIMAG:
		ok |= Erv
		if !psess.onearg(n, "%v", n.Op) {
			n.Type = nil
			return n
		}

		n.Left = psess.typecheck(n.Left, Erv)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}

		et := t.Etype
		switch et {
		case TIDEAL:

		case TCOMPLEX64:
			et = TFLOAT32
		case TCOMPLEX128:
			et = TFLOAT64
		default:
			psess.
				yyerror("invalid argument %L for %v", l, n.Op)
			n.Type = nil
			return n
		}
		n.Type = psess.types.Types[et]

		if l.Op == OLITERAL {
			var re, im *Mpflt
			switch psess.consttype(l) {
			case CTINT, CTRUNE:
				re = newMpflt()
				re.SetInt(l.Val().U.(*Mpint))

			case CTFLT:
				re = l.Val().U.(*Mpflt)

			case CTCPLX:
				re = &l.Val().U.(*Mpcplx).Real
				im = &l.Val().U.(*Mpcplx).Imag
			default:
				psess.
					yyerror("invalid argument %L for %v", l, n.Op)
				n.Type = nil
				return n
			}
			if n.Op == OIMAG {
				if im == nil {
					im = newMpflt()
				}
				re = im
			}
			psess.
				setconst(n, Val{re})
		}

	case OCOMPLEX:
		ok |= Erv
		var r *Node
		var l *Node
		if n.List.Len() == 1 {
			psess.
				typecheckslice(n.List.Slice(), Efnstruct)
			if n.List.First().Op != OCALLFUNC && n.List.First().Op != OCALLMETH {
				psess.
					yyerror("invalid operation: complex expects two arguments")
				n.Type = nil
				return n
			}

			t := n.List.First().Left.Type
			if !t.IsKind(TFUNC) {

				return n
			}
			if t.NumResults(psess.types) != 2 {
				psess.
					yyerror("invalid operation: complex expects two arguments, %v returns %d results", n.List.First(), t.NumResults(psess.types))
				n.Type = nil
				return n
			}

			t = n.List.First().Type
			l = asNode(t.Field(psess.types, 0).Nname)
			r = asNode(t.Field(psess.types, 1).Nname)
		} else {
			if !psess.twoarg(n) {
				n.Type = nil
				return n
			}
			n.Left = psess.typecheck(n.Left, Erv)
			n.Right = psess.typecheck(n.Right, Erv)
			l = n.Left
			r = n.Right
			if l.Type == nil || r.Type == nil {
				n.Type = nil
				return n
			}
			l, r = psess.defaultlit2(l, r, false)
			if l.Type == nil || r.Type == nil {
				n.Type = nil
				return n
			}
			n.Left = l
			n.Right = r
		}

		if !psess.eqtype(l.Type, r.Type) {
			psess.
				yyerror("invalid operation: %v (mismatched types %v and %v)", n, l.Type, r.Type)
			n.Type = nil
			return n
		}

		var t *types.Type
		switch l.Type.Etype {
		default:
			psess.
				yyerror("invalid operation: %v (arguments have type %v, expected floating-point)", n, l.Type)
			n.Type = nil
			return n

		case TIDEAL:
			t = psess.types.Types[TIDEAL]

		case TFLOAT32:
			t = psess.types.Types[TCOMPLEX64]

		case TFLOAT64:
			t = psess.types.Types[TCOMPLEX128]
		}
		n.Type = t

		if l.Op == OLITERAL && r.Op == OLITERAL {

			c := new(Mpcplx)
			c.Real.Set(psess.toflt(l.Val()).U.(*Mpflt))
			c.Imag.Set(psess.toflt(r.Val()).U.(*Mpflt))
			psess.
				setconst(n, Val{c})
		}

	case OCLOSE:
		if !psess.onearg(n, "%v", n.Op) {
			n.Type = nil
			return n
		}
		n.Left = psess.typecheck(n.Left, Erv)
		n.Left = psess.defaultlit(n.Left, nil)
		l := n.Left
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsChan() {
			psess.
				yyerror("invalid operation: %v (non-chan type %v)", n, t)
			n.Type = nil
			return n
		}

		if !t.ChanDir(psess.types).CanSend() {
			psess.
				yyerror("invalid operation: %v (cannot close receive-only channel)", n)
			n.Type = nil
			return n
		}

		ok |= Etop

	case ODELETE:
		args := n.List
		if args.Len() == 0 {
			psess.
				yyerror("missing arguments to delete")
			n.Type = nil
			return n
		}

		if args.Len() == 1 {
			psess.
				yyerror("missing second (key) argument to delete")
			n.Type = nil
			return n
		}

		if args.Len() != 2 {
			psess.
				yyerror("too many arguments to delete")
			n.Type = nil
			return n
		}

		ok |= Etop
		psess.
			typecheckslice(args.Slice(), Erv)
		l := args.First()
		r := args.Second()
		if l.Type != nil && !l.Type.IsMap() {
			psess.
				yyerror("first argument to delete must be map; have %L", l.Type)
			n.Type = nil
			return n
		}

		args.SetSecond(psess.assignconv(r, l.Type.Key(psess.types), "delete"))

	case OAPPEND:
		ok |= Erv
		args := n.List
		if args.Len() == 0 {
			psess.
				yyerror("missing arguments to append")
			n.Type = nil
			return n
		}

		if args.Len() == 1 && !n.Isddd() {
			args.SetFirst(psess.typecheck(args.First(), Erv|Efnstruct))
		} else {
			psess.
				typecheckslice(args.Slice(), Erv)
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
			t = t.Field(psess.types, 0).Type
		}

		n.Type = t
		if !t.IsSlice() {
			if psess.Isconst(args.First(), CTNIL) {
				psess.
					yyerror("first argument to append must be typed slice; have untyped nil")
				n.Type = nil
				return n
			}
			psess.
				yyerror("first argument to append must be slice; have %L", t)
			n.Type = nil
			return n
		}

		if n.Isddd() {
			if args.Len() == 1 {
				psess.
					yyerror("cannot use ... on first argument to append")
				n.Type = nil
				return n
			}

			if args.Len() != 2 {
				psess.
					yyerror("too many arguments to append")
				n.Type = nil
				return n
			}

			if t.Elem(psess.types).IsKind(TUINT8) && args.Second().Type.IsString() {
				args.SetSecond(psess.defaultlit(args.Second(), psess.types.Types[TSTRING]))
				break
			}

			args.SetSecond(psess.assignconv(args.Second(), t.Orig, "append"))
			break
		}

		if funarg != nil {
			for _, t := range funarg.FieldSlice(psess.types)[1:] {
				if psess.assignop(t.Type, n.Type.Elem(psess.types), nil) == 0 {
					psess.
						yyerror("cannot append %v value to []%v", t.Type, n.Type.Elem(psess.types))
				}
			}
		} else {
			as := args.Slice()[1:]
			for i, n := range as {
				if n.Type == nil {
					continue
				}
				as[i] = psess.assignconv(n, t.Elem(psess.types), "append")
				psess.
					checkwidth(as[i].Type)
			}
		}

	case OCOPY:
		ok |= Etop | Erv
		args := n.List
		if args.Len() < 2 {
			psess.
				yyerror("missing arguments to copy")
			n.Type = nil
			return n
		}

		if args.Len() > 2 {
			psess.
				yyerror("too many arguments to copy")
			n.Type = nil
			return n
		}

		n.Left = args.First()
		n.Right = args.Second()
		n.List.Set(nil)
		n.Type = psess.types.Types[TINT]
		n.Left = psess.typecheck(n.Left, Erv)
		n.Right = psess.typecheck(n.Right, Erv)
		if n.Left.Type == nil || n.Right.Type == nil {
			n.Type = nil
			return n
		}
		n.Left = psess.defaultlit(n.Left, nil)
		n.Right = psess.defaultlit(n.Right, nil)
		if n.Left.Type == nil || n.Right.Type == nil {
			n.Type = nil
			return n
		}

		if n.Left.Type.IsSlice() && n.Right.Type.IsString() {
			if psess.eqtype(n.Left.Type.Elem(psess.types), psess.types.Bytetype) {
				break
			}
			psess.
				yyerror("arguments to copy have different element types: %L and string", n.Left.Type)
			n.Type = nil
			return n
		}

		if !n.Left.Type.IsSlice() || !n.Right.Type.IsSlice() {
			if !n.Left.Type.IsSlice() && !n.Right.Type.IsSlice() {
				psess.
					yyerror("arguments to copy must be slices; have %L, %L", n.Left.Type, n.Right.Type)
			} else if !n.Left.Type.IsSlice() {
				psess.
					yyerror("first argument to copy should be slice; have %L", n.Left.Type)
			} else {
				psess.
					yyerror("second argument to copy should be slice or string; have %L", n.Right.Type)
			}
			n.Type = nil
			return n
		}

		if !psess.eqtype(n.Left.Type.Elem(psess.types), n.Right.Type.Elem(psess.types)) {
			psess.
				yyerror("arguments to copy have different element types: %L and %L", n.Left.Type, n.Right.Type)
			n.Type = nil
			return n
		}

	case OCONV:
		ok |= Erv
		psess.
			checkwidth(n.Type)
		n.Left = psess.typecheck(n.Left, Erv)
		n.Left = psess.convlit1(n.Left, n.Type, true, noReuse)
		t := n.Left.Type
		if t == nil || n.Type == nil {
			n.Type = nil
			return n
		}
		var why string
		n.Op = psess.convertop(t, n.Type, &why)
		if n.Op == 0 {
			if !n.Diag() && !n.Type.Broke() && !n.Left.Diag() {
				psess.
					yyerror("cannot convert %L to type %v%s", n.Left, n.Type, why)
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
				psess.
					setconst(n, n.Left.Val())
			} else if t.Etype == n.Type.Etype {
				switch t.Etype {
				case TFLOAT32, TFLOAT64, TCOMPLEX64, TCOMPLEX128:

					n.Op = OCONV
				}
			}

		case OSTRARRAYBYTE:
			break

		case OSTRARRAYRUNE:
			if n.Left.Op == OLITERAL {
				n = psess.stringtoarraylit(n)
			}
		}

	case OMAKE:
		ok |= Erv
		args := n.List.Slice()
		if len(args) == 0 {
			psess.
				yyerror("missing argument to make")
			n.Type = nil
			return n
		}

		n.List.Set(nil)
		l := args[0]
		l = psess.typecheck(l, Etype)
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}

		i := 1
		switch t.Etype {
		default:
			psess.
				yyerror("cannot make type %v", t)
			n.Type = nil
			return n

		case TSLICE:
			if i >= len(args) {
				psess.
					yyerror("missing len argument to make(%v)", t)
				n.Type = nil
				return n
			}

			l = args[i]
			i++
			l = psess.typecheck(l, Erv)
			var r *Node
			if i < len(args) {
				r = args[i]
				i++
				r = psess.typecheck(r, Erv)
			}

			if l.Type == nil || (r != nil && r.Type == nil) {
				n.Type = nil
				return n
			}
			if !psess.checkmake(t, "len", l) || r != nil && !psess.checkmake(t, "cap", r) {
				n.Type = nil
				return n
			}
			if psess.Isconst(l, CTINT) && r != nil && psess.Isconst(r, CTINT) && l.Val().U.(*Mpint).Cmp(r.Val().U.(*Mpint)) > 0 {
				psess.
					yyerror("len larger than cap in make(%v)", t)
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
				l = psess.typecheck(l, Erv)
				l = psess.defaultlit(l, psess.types.Types[TINT])
				if l.Type == nil {
					n.Type = nil
					return n
				}
				if !psess.checkmake(t, "size", l) {
					n.Type = nil
					return n
				}
				n.Left = l
			} else {
				n.Left = psess.nodintconst(0)
			}
			n.Op = OMAKEMAP

		case TCHAN:
			l = nil
			if i < len(args) {
				l = args[i]
				i++
				l = psess.typecheck(l, Erv)
				l = psess.defaultlit(l, psess.types.Types[TINT])
				if l.Type == nil {
					n.Type = nil
					return n
				}
				if !psess.checkmake(t, "buffer", l) {
					n.Type = nil
					return n
				}
				n.Left = l
			} else {
				n.Left = psess.nodintconst(0)
			}
			n.Op = OMAKECHAN
		}

		if i < len(args) {
			psess.
				yyerror("too many arguments to make(%v)", t)
			n.Op = OMAKE
			n.Type = nil
			return n
		}

		n.Type = t

	case ONEW:
		ok |= Erv
		args := n.List
		if args.Len() == 0 {
			psess.
				yyerror("missing argument to new")
			n.Type = nil
			return n
		}

		l := args.First()
		l = psess.typecheck(l, Etype)
		t := l.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if args.Len() > 1 {
			psess.
				yyerror("too many arguments to new(%v)", t)
			n.Type = nil
			return n
		}

		n.Left = l
		n.Type = psess.types.NewPtr(t)

	case OPRINT, OPRINTN:
		ok |= Etop
		psess.
			typecheckslice(n.List.Slice(), Erv)
		ls := n.List.Slice()
		for i1, n1 := range ls {

			if psess.Isconst(n1, CTINT) {
				ls[i1] = psess.defaultlit(ls[i1], psess.types.Types[TINT64])
			} else {
				ls[i1] = psess.defaultlit(ls[i1], nil)
			}
		}

	case OPANIC:
		ok |= Etop
		if !psess.onearg(n, "panic") {
			n.Type = nil
			return n
		}
		n.Left = psess.typecheck(n.Left, Erv)
		n.Left = psess.defaultlit(n.Left, psess.types.Types[TINTER])
		if n.Left.Type == nil {
			n.Type = nil
			return n
		}

	case ORECOVER:
		ok |= Erv | Etop
		if n.List.Len() != 0 {
			psess.
				yyerror("too many arguments to recover")
			n.Type = nil
			return n
		}

		n.Type = psess.types.Types[TINTER]

	case OCLOSURE:
		ok |= Erv
		psess.
			typecheckclosure(n, top)
		if n.Type == nil {
			return n
		}

	case OITAB:
		ok |= Erv
		n.Left = psess.typecheck(n.Left, Erv)
		t := n.Left.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsInterface() {
			psess.
				Fatalf("OITAB of %v", t)
		}
		n.Type = psess.types.NewPtr(psess.types.Types[TUINTPTR])

	case OIDATA:
		psess.
			Fatalf("cannot typecheck interface data %v", n)

	case OSPTR:
		ok |= Erv
		n.Left = psess.typecheck(n.Left, Erv)
		t := n.Left.Type
		if t == nil {
			n.Type = nil
			return n
		}
		if !t.IsSlice() && !t.IsString() {
			psess.
				Fatalf("OSPTR of %v", t)
		}
		if t.IsString() {
			n.Type = psess.types.NewPtr(psess.types.Types[TUINT8])
		} else {
			n.Type = psess.types.NewPtr(t.Elem(psess.types))
		}

	case OCLOSUREVAR:
		ok |= Erv

	case OCFUNC:
		ok |= Erv
		n.Left = psess.typecheck(n.Left, Erv)
		n.Type = psess.types.Types[TUINTPTR]

	case OCONVNOP:
		ok |= Erv
		n.Left = psess.typecheck(n.Left, Erv)

	case OAS:
		ok |= Etop
		psess.
			typecheckas(n)

		if n.Left.Op == ONAME && n.Left.IsAutoTmp() {
			n.Left.Name.Defn = n
		}

	case OAS2:
		ok |= Etop
		psess.
			typecheckas2(n)

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
		psess.
			decldepth++
		if n.Left.Sym.IsBlank() {

			n.Op = OEMPTY
			n.Left = nil
		}

	case ODEFER:
		ok |= Etop
		n.Left = psess.typecheck(n.Left, Etop|Erv)
		if !n.Left.Diag() {
			psess.
				checkdefergo(n)
		}

	case OPROC:
		ok |= Etop
		n.Left = psess.typecheck(n.Left, Etop|Erv)
		psess.
			checkdefergo(n)

	case OFOR, OFORUNTIL:
		ok |= Etop
		psess.
			typecheckslice(n.Ninit.Slice(), Etop)
		psess.
			decldepth++
		n.Left = psess.typecheck(n.Left, Erv)
		n.Left = psess.defaultlit(n.Left, nil)
		if n.Left != nil {
			t := n.Left.Type
			if t != nil && !t.IsBoolean() {
				psess.
					yyerror("non-bool %L used as for condition", n.Left)
			}
		}
		n.Right = psess.typecheck(n.Right, Etop)
		if n.Op == OFORUNTIL {
			psess.
				typecheckslice(n.List.Slice(), Etop)
		}
		psess.
			typecheckslice(n.Nbody.Slice(), Etop)
		psess.
			decldepth--

	case OIF:
		ok |= Etop
		psess.
			typecheckslice(n.Ninit.Slice(), Etop)
		n.Left = psess.typecheck(n.Left, Erv)
		n.Left = psess.defaultlit(n.Left, nil)
		if n.Left != nil {
			t := n.Left.Type
			if t != nil && !t.IsBoolean() {
				psess.
					yyerror("non-bool %L used as if condition", n.Left)
			}
		}
		psess.
			typecheckslice(n.Nbody.Slice(), Etop)
		psess.
			typecheckslice(n.Rlist.Slice(), Etop)

	case ORETURN:
		ok |= Etop
		if n.List.Len() == 1 {
			psess.
				typecheckslice(n.List.Slice(), Erv|Efnstruct)
		} else {
			psess.
				typecheckslice(n.List.Slice(), Erv)
		}
		if psess.Curfn == nil {
			psess.
				yyerror("return outside function")
			n.Type = nil
			return n
		}

		if psess.Curfn.Type.FuncType(psess.types).Outnamed && n.List.Len() == 0 {
			break
		}
		psess.
			typecheckaste(ORETURN, nil, false, psess.Curfn.Type.Results(psess.types), n.List, func() string { return "return argument" })

	case ORETJMP:
		ok |= Etop

	case OSELECT:
		ok |= Etop
		psess.
			typecheckselect(n)

	case OSWITCH:
		ok |= Etop
		psess.
			typecheckswitch(n)

	case ORANGE:
		ok |= Etop
		psess.
			typecheckrange(n)

	case OTYPESW:
		psess.
			yyerror("use of .(type) outside type switch")
		n.Type = nil
		return n

	case OXCASE:
		ok |= Etop
		psess.
			typecheckslice(n.List.Slice(), Erv)
		psess.
			typecheckslice(n.Nbody.Slice(), Etop)

	case ODCLFUNC:
		ok |= Etop
		psess.
			typecheckfunc(n)

	case ODCLCONST:
		ok |= Etop
		n.Left = psess.typecheck(n.Left, Erv)

	case ODCLTYPE:
		ok |= Etop
		n.Left = psess.typecheck(n.Left, Etype)
		psess.
			checkwidth(n.Left.Type)
		if n.Left.Type != nil && n.Left.Type.NotInHeap() && n.Left.Name.Param.Pragma&NotInHeap == 0 {
			psess.
				yyerror("type %v must be go:notinheap", n.Left.Type)
		}
	}

	t := n.Type
	if t != nil && !t.IsFuncArgStruct() && n.Op != OTYPE {
		switch t.Etype {
		case TFUNC,
			TANY, TFORW, TIDEAL, TNIL, TBLANK:
			break

		default:
			psess.
				checkwidth(t)
		}
	}

	if psess.safemode && !psess.inimport && !psess.compiling_wrappers && t != nil && t.Etype == TUNSAFEPTR {
		psess.
			yyerror("cannot use unsafe.Pointer")
	}
	psess.
		evconst(n)
	if n.Op == OTYPE && top&Etype == 0 {
		if !n.Type.Broke() {
			psess.
				yyerror("type %v is not an expression", n.Type)
		}
		n.Type = nil
		return n
	}

	if top&(Erv|Etype) == Etype && n.Op != OTYPE {
		psess.
			yyerror("%v is not a type", n)
		n.Type = nil
		return n
	}

	if (top&(Ecall|Erv|Etype) != 0) && top&Etop == 0 && ok&(Erv|Etype|Ecall) == 0 {
		psess.
			yyerror("%v used as value", n)
		n.Type = nil
		return n
	}

	if (top&Etop != 0) && top&(Ecall|Erv|Etype) == 0 && ok&Etop == 0 {
		if !n.Diag() {
			psess.
				yyerror("%v evaluated but not used", n)
			n.SetDiag(true)
		}

		n.Type = nil
		return n
	}

	return n
}

func (psess *PackageSession) checksliceindex(l *Node, r *Node, tp *types.Type) bool {
	t := r.Type
	if t == nil {
		return false
	}
	if !t.IsInteger() {
		psess.
			yyerror("invalid slice index %v (type %v)", r, t)
		return false
	}

	if r.Op == OLITERAL {
		if r.Int64(psess) < 0 {
			psess.
				yyerror("invalid slice index %v (index must be non-negative)", r)
			return false
		} else if tp != nil && tp.NumElem(psess.types) >= 0 && r.Int64(psess) > tp.NumElem(psess.types) {
			psess.
				yyerror("invalid slice index %v (out of bounds for %d-element array)", r, tp.NumElem(psess.types))
			return false
		} else if psess.Isconst(l, CTSTR) && r.Int64(psess) > int64(len(l.Val().U.(string))) {
			psess.
				yyerror("invalid slice index %v (out of bounds for %d-byte string)", r, len(l.Val().U.(string)))
			return false
		} else if r.Val().U.(*Mpint).Cmp(psess.maxintval[TINT]) > 0 {
			psess.
				yyerror("invalid slice index %v (index too large)", r)
			return false
		}
	}

	return true
}

func (psess *PackageSession) checksliceconst(lo *Node, hi *Node) bool {
	if lo != nil && hi != nil && lo.Op == OLITERAL && hi.Op == OLITERAL && lo.Val().U.(*Mpint).Cmp(hi.Val().U.(*Mpint)) > 0 {
		psess.
			yyerror("invalid slice index: %v > %v", lo, hi)
		return false
	}

	return true
}

func (psess *PackageSession) checkdefergo(n *Node) {
	what := "defer"
	if n.Op == OPROC {
		what = "go"
	}

	switch n.Left.Op {

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
		OLITERAL:
		if n.Left.Orig != nil && n.Left.Orig.Op == OCONV {
			break
		}
		psess.
			yyerror("%s discards result of %v", what, n.Left)
		return
	}

	if n.Left.Type == nil || n.Left.Type.Broke() {
		return
	}

	if !n.Diag() {

		n.SetDiag(true)
		psess.
			yyerror("%s requires function call, not conversion", what)
	}
}

// The result of implicitstar MUST be assigned back to n, e.g.
// 	n.Left = implicitstar(n.Left)
func (psess *PackageSession) implicitstar(n *Node) *Node {

	t := n.Type
	if t == nil || !t.IsPtr() {
		return n
	}
	t = t.Elem(psess.types)
	if t == nil {
		return n
	}
	if !t.IsArray() {
		return n
	}
	n = psess.nod(OIND, n, nil)
	n.SetImplicit(true)
	n = psess.typecheck(n, Erv)
	return n
}

func (psess *PackageSession) onearg(n *Node, f string, args ...interface{}) bool {
	if n.Left != nil {
		return true
	}
	if n.List.Len() == 0 {
		p := fmt.Sprintf(f, args...)
		psess.
			yyerror("missing argument to %s: %v", p, n)
		return false
	}

	if n.List.Len() > 1 {
		p := fmt.Sprintf(f, args...)
		psess.
			yyerror("too many arguments to %s: %v", p, n)
		n.Left = n.List.First()
		n.List.Set(nil)
		return false
	}

	n.Left = n.List.First()
	n.List.Set(nil)
	return true
}

func (psess *PackageSession) twoarg(n *Node) bool {
	if n.Left != nil {
		return true
	}
	if n.List.Len() == 0 {
		psess.
			yyerror("missing argument to %v - %v", n.Op, n)
		return false
	}

	n.Left = n.List.First()
	if n.List.Len() == 1 {
		psess.
			yyerror("missing argument to %v - %v", n.Op, n)
		n.List.Set(nil)
		return false
	}

	if n.List.Len() > 2 {
		psess.
			yyerror("too many arguments to %v - %v", n.Op, n)
		n.List.Set(nil)
		return false
	}

	n.Right = n.List.Second()
	n.List.Set(nil)
	return true
}

func (psess *PackageSession) lookdot1(errnode *Node, s *types.Sym, t *types.Type, fs *types.Fields, dostrcmp int) *types.Field {
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
				psess.
					yyerror("ambiguous selector %v", errnode)
			} else if t.IsPtr() {
				psess.
					yyerror("ambiguous selector (%v).%v", t, s)
			} else {
				psess.
					yyerror("ambiguous selector %v.%v", t, s)
			}
			break
		}

		r = f
	}

	return r
}

// typecheckMethodExpr checks selector expressions (ODOT) where the
// base expression is a type expression (OTYPE).
func (psess *PackageSession) typecheckMethodExpr(n *Node) *Node {
	t := n.Left.Type

	// Compute the method set for t.
	var ms *types.Fields
	if t.IsInterface() {
		ms = t.Fields(psess.types)
	} else {
		mt := psess.methtype(t)
		if mt == nil {
			psess.
				yyerror("%v undefined (type %v has no method %v)", n, t, n.Sym)
			n.Type = nil
			return n
		}
		psess.
			expandmeth(mt)
		ms = mt.AllMethods()

		if mt.Sym == nil {
			psess.
				addsignat(t)
		}
	}

	s := n.Sym
	m := psess.lookdot1(n, s, t, ms, 0)
	if m == nil {
		if psess.lookdot1(n, s, t, ms, 1) != nil {
			psess.
				yyerror("%v undefined (cannot refer to unexported method %v)", n, s)
		} else if _, ambig := psess.dotpath(s, t, nil, false); ambig {
			psess.
				yyerror("%v undefined (ambiguous selector)", n)
		} else {
			psess.
				yyerror("%v undefined (type %v has no method %v)", n, t, s)
		}
		n.Type = nil
		return n
	}

	if !psess.isMethodApplicable(t, m) {
		psess.
			yyerror("invalid method expression %v (needs pointer receiver: (*%v).%S)", n, t, s)
		n.Type = nil
		return n
	}

	n.Op = ONAME
	if n.Name == nil {
		n.Name = new(Name)
	}
	n.Right = psess.newname(n.Sym)
	n.Sym = psess.methodSym(t, n.Sym)
	n.Type = psess.methodfunc(m.Type, n.Left.Type)
	n.Xoffset = 0
	n.SetClass(PFUNC)
	return n
}

// isMethodApplicable reports whether method m can be called on a
// value of type t. This is necessary because we compute a single
// method set for both T and *T, but some *T methods are not
// applicable to T receivers.
func (psess *PackageSession) isMethodApplicable(t *types.Type, m *types.Field) bool {
	return t.IsPtr() || !m.Type.Recv(psess.types).Type.IsPtr() || psess.isifacemethod(m.Type) || m.Embedded == 2
}

func (psess *PackageSession) derefall(t *types.Type) *types.Type {
	for t != nil && t.Etype == psess.types.Tptr {
		t = t.Elem(psess.types)
	}
	return t
}

type typeSymKey struct {
	t *types.Type
	s *types.Sym
}

// dotField maps (*types.Type, *types.Sym) pairs to the corresponding struct field (*types.Type with Etype==TFIELD).
// It is a cache for use during usefield in walk.go, only enabled when field tracking.

func (psess *PackageSession) lookdot(n *Node, t *types.Type, dostrcmp int) *types.Field {
	s := n.Sym
	psess.
		dowidth(t)
	var f1 *types.Field
	if t.IsStruct() || t.IsInterface() {
		f1 = psess.lookdot1(n, s, t, t.Fields(psess.types), dostrcmp)
	}

	var f2 *types.Field
	if n.Left.Type == t || n.Left.Type.Sym == nil {
		mt := psess.methtype(t)
		if mt != nil {
			f2 = psess.lookdot1(n, s, mt, mt.Methods(), dostrcmp)
		}
	}

	if f1 != nil {
		if dostrcmp > 1 {

			return f1
		}
		if f2 != nil {
			psess.
				yyerror("%v is both field and method", n.Sym)
		}
		if f1.Offset == BADWIDTH {
			psess.
				Fatalf("lookdot badwidth %v %p", f1, f1)
		}
		n.Xoffset = f1.Offset
		n.Type = f1.Type
		if psess.objabi.Fieldtrack_enabled > 0 {
			psess.
				dotField[typeSymKey{t.Orig, s}] = f1
		}
		if t.IsInterface() {
			if n.Left.Type.IsPtr() {
				n.Left = psess.nod(OIND, n.Left, nil)
				n.Left.SetImplicit(true)
				n.Left = psess.typecheck(n.Left, Erv)
			}

			n.Op = ODOTINTER
		}

		return f1
	}

	if f2 != nil {
		if dostrcmp > 1 {

			return f2
		}
		tt := n.Left.Type
		psess.
			dowidth(tt)
		rcvr := f2.Type.Recv(psess.types).Type
		if !psess.eqtype(rcvr, tt) {
			if rcvr.Etype == psess.types.Tptr && psess.eqtype(rcvr.Elem(psess.types), tt) {
				psess.
					checklvalue(n.Left, "call pointer method on")
				n.Left = psess.nod(OADDR, n.Left, nil)
				n.Left.SetImplicit(true)
				n.Left = psess.typecheck(n.Left, Etype|Erv)
			} else if tt.Etype == psess.types.Tptr && rcvr.Etype != psess.types.Tptr && psess.eqtype(tt.Elem(psess.types), rcvr) {
				n.Left = psess.nod(OIND, n.Left, nil)
				n.Left.SetImplicit(true)
				n.Left = psess.typecheck(n.Left, Etype|Erv)
			} else if tt.Etype == psess.types.Tptr && tt.Elem(psess.types).Etype == psess.types.Tptr && psess.eqtype(psess.derefall(tt), psess.derefall(rcvr)) {
				psess.
					yyerror("calling method %v with receiver %L requires explicit dereference", n.Sym, n.Left)
				for tt.Etype == psess.types.Tptr {

					if rcvr.Etype == psess.types.Tptr && tt.Elem(psess.types).Etype != psess.types.Tptr {
						break
					}
					n.Left = psess.nod(OIND, n.Left, nil)
					n.Left.SetImplicit(true)
					n.Left = psess.typecheck(n.Left, Etype|Erv)
					tt = tt.Elem(psess.types)
				}
			} else {
				psess.
					Fatalf("method mismatch: %v for %v", rcvr, tt)
			}
		}

		pll := n
		ll := n.Left
		for ll.Left != nil && (ll.Op == ODOT || ll.Op == ODOTPTR || ll.Op == OIND) {
			pll = ll
			ll = ll.Left
		}
		if pll.Implicit() && ll.Type.IsPtr() && ll.Type.Sym != nil && asNode(ll.Type.Sym.Def) != nil && asNode(ll.Type.Sym.Def).Op == OTYPE {

			n.Left = ll
			return nil
		}

		n.Sym = psess.methodSym(n.Left.Type, f2.Sym)
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

func (psess *PackageSession) hasddd(t *types.Type) bool {
	for _, tl := range t.Fields(psess.types).Slice() {
		if tl.Isddd() {
			return true
		}
	}

	return false
}

// typecheck assignment: type list = expression list
func (psess *PackageSession) typecheckaste(op Op, call *Node, isddd bool, tstruct *types.Type, nl Nodes, desc func() string) {
	var t *types.Type
	var n1 int
	var n2 int
	var i int

	lno := psess.lineno
	defer func() {
		psess.lineno = lno
	}()

	if tstruct.Broke() {
		return
	}

	var n *Node
	if nl.Len() == 1 {
		n = nl.First()
		if n.Type != nil && n.Type.IsFuncArgStruct() {
			if !psess.hasddd(tstruct) {
				n1 := tstruct.NumFields(psess.types)
				n2 := n.Type.NumFields(psess.types)
				if n2 > n1 {
					goto toomany
				}
				if n2 < n1 {
					goto notenough
				}
			}

			lfs := tstruct.FieldSlice(psess.types)
			rfs := n.Type.FieldSlice(psess.types)
			var why string
			for i, tl := range lfs {
				if tl.Isddd() {
					for _, tn := range rfs[i:] {
						if psess.assignop(tn.Type, tl.Type.Elem(psess.types), &why) == 0 {
							if call != nil {
								psess.
									yyerror("cannot use %v as type %v in argument to %v%s", tn.Type, tl.Type.Elem(psess.types), call, why)
							} else {
								psess.
									yyerror("cannot use %v as type %v in %s%s", tn.Type, tl.Type.Elem(psess.types), desc(), why)
							}
						}
					}
					return
				}

				if i >= len(rfs) {
					goto notenough
				}
				tn := rfs[i]
				if psess.assignop(tn.Type, tl.Type, &why) == 0 {
					if call != nil {
						psess.
							yyerror("cannot use %v as type %v in argument to %v%s", tn.Type, tl.Type, call, why)
					} else {
						psess.
							yyerror("cannot use %v as type %v in %s%s", tn.Type, tl.Type, desc(), why)
					}
				}
			}

			if len(rfs) > len(lfs) {
				goto toomany
			}
			return
		}
	}

	n1 = tstruct.NumFields(psess.types)
	n2 = nl.Len()
	if !psess.hasddd(tstruct) {
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
	for _, tl := range tstruct.Fields(psess.types).Slice() {
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
				psess.
					setlineno(n)
				if n.Type != nil {
					nl.SetIndex(i, psess.assignconvfn(n, t, desc))
				}
				return
			}

			for ; i < nl.Len(); i++ {
				n = nl.Index(i)
				psess.
					setlineno(n)
				if n.Type != nil {
					nl.SetIndex(i, psess.assignconvfn(n, t.Elem(psess.types), desc))
				}
			}
			return
		}

		if i >= nl.Len() {
			goto notenough
		}
		n = nl.Index(i)
		psess.
			setlineno(n)
		if n.Type != nil {
			nl.SetIndex(i, psess.assignconvfn(n, t, desc))
		}
		i++
	}

	if i < nl.Len() {
		goto toomany
	}
	if isddd {
		if call != nil {
			psess.
				yyerror("invalid use of ... in call to %v", call)
		} else {
			psess.
				yyerror("invalid use of ... in %v", op)
		}
	}
	return

notenough:
	if n == nil || !n.Diag() {
		details := psess.errorDetails(nl, tstruct, isddd)
		if call != nil {

			if call.isMethodExpression() {
				psess.
					yyerror("not enough arguments in call to method expression %v%s", call, details)
			} else {
				psess.
					yyerror("not enough arguments in call to %v%s", call, details)
			}
		} else {
			psess.
				yyerror("not enough arguments to %v%s", op, details)
		}
		if n != nil {
			n.SetDiag(true)
		}
	}
	return

toomany:
	details := psess.errorDetails(nl, tstruct, isddd)
	if call != nil {
		psess.
			yyerror("too many arguments in call to %v%s", call, details)
	} else {
		psess.
			yyerror("too many arguments to %v%s", op, details)
	}
}

func (psess *PackageSession) errorDetails(nl Nodes, tstruct *types.Type, isddd bool) string {

	if tstruct == nil {
		return ""
	}

	for _, n := range nl.Slice() {
		if n.Type == nil {
			return ""
		}
	}
	return fmt.Sprintf("\n\thave %s\n\twant %v", nl.retsigerr(psess, isddd), tstruct)
}

// sigrepr is a type's representation to the outside world,
// in string representations of return signatures
// e.g in error messages about wrong arguments to return.
func (psess *PackageSession) sigrepr(t *types.Type) string {
	switch t {
	default:
		return t.String(psess.types)

	case psess.types.Types[TIDEAL]:

		return "number"

	case psess.types.Idealstring:
		return "string"

	case psess.types.Idealbool:
		return "bool"
	}
}

// retsigerr returns the signature of the types
// at the respective return call site of a function.
func (nl Nodes) retsigerr(psess *PackageSession, isddd bool) string {
	if nl.Len() < 1 {
		return "()"
	}

	var typeStrings []string
	if nl.Len() == 1 && nl.First().Type != nil && nl.First().Type.IsFuncArgStruct() {
		for _, f := range nl.First().Type.Fields(psess.types).Slice() {
			typeStrings = append(typeStrings, psess.sigrepr(f.Type))
		}
	} else {
		for _, n := range nl.Slice() {
			typeStrings = append(typeStrings, psess.sigrepr(n.Type))
		}
	}

	ddd := ""
	if isddd {
		ddd = "..."
	}
	return fmt.Sprintf("(%s%s)", strings.Join(typeStrings, ", "), ddd)
}

// type check composite
func (psess *PackageSession) fielddup(name string, hash map[string]bool) {
	if hash[name] {
		psess.
			yyerror("duplicate field name in struct literal: %s", name)
		return
	}
	hash[name] = true
}

func (psess *PackageSession) keydup(n *Node, hash map[uint32][]*Node) {
	orign := n
	if n.Op == OCONVIFACE {
		n = n.Left
	}
	psess.
		evconst(n)
	if n.Op != OLITERAL {
		return
	}

	const PRIME1 = 3

	var h uint32
	switch v := n.Val().U.(type) {
	default:
		h = 23

	case *Mpint:
		h = uint32(v.Int64(psess))

	case *Mpflt:
		x := math.Float64bits(v.Float64(psess))
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
		if !psess.eqtype(a.Type, n.Type) {
			continue
		}
		cmp.Right = a
		psess.
			evconst(&cmp)
		if cmp.Op != OLITERAL {

			continue
		}
		if cmp.Val().U.(bool) {
			psess.
				yyerror("duplicate key %v in map literal", n)
			return
		}
	}

	hash[h] = append(hash[h], orign)
}

// iscomptype reports whether type t is a composite literal type
// or a pointer to one.
func (psess *PackageSession) iscomptype(t *types.Type) bool {
	if t.IsPtr() {
		t = t.Elem(psess.types)
	}

	switch t.Etype {
	case TARRAY, TSLICE, TSTRUCT, TMAP:
		return true
	default:
		return false
	}
}

func (psess *PackageSession) pushtype(n *Node, t *types.Type) {
	if n == nil || n.Op != OCOMPLIT || !psess.iscomptype(t) {
		return
	}

	if n.Right == nil {
		n.Right = psess.typenod(t)
		n.SetImplicit(true)
		n.Right.SetImplicit(true)
	} else if psess.Debug['s'] != 0 {
		n.Right = psess.typecheck(n.Right, Etype)
		if n.Right.Type != nil && psess.eqtype(n.Right.Type, t) {
			fmt.Printf("%v: redundant type: %v\n", n.Line(psess), t)
		}
	}
}

// The result of typecheckcomplit MUST be assigned back to n, e.g.
// 	n.Left = typecheckcomplit(n.Left)
func (psess *PackageSession) typecheckcomplit(n *Node) *Node {
	lno := psess.lineno
	defer func() {
		psess.
			lineno = lno
	}()

	if n.Right == nil {
		psess.
			yyerrorl(n.Pos, "missing type in composite literal")
		n.Type = nil
		return n
	}

	norig := n.copy()
	psess.
		setlineno(n.Right)
	n.Right = psess.typecheck(n.Right, Etype|Ecomplit)
	l := n.Right
	t := l.Type
	if t == nil {
		n.Type = nil
		return n
	}
	nerr := psess.nerrors
	n.Type = t

	if t.IsPtr() {

		if !n.Right.Implicit() {
			psess.
				yyerror("invalid pointer type %v for composite literal (use &%v instead)", t, t.Elem(psess.types))
			n.Type = nil
			return n
		}

		if !psess.iscomptype(t) {
			psess.
				yyerror("invalid pointer type %v for composite literal", t)
			n.Type = nil
			return n
		}

		t = t.Elem(psess.types)
	}

	switch t.Etype {
	default:
		psess.
			yyerror("invalid type for composite literal: %v", t)
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
			psess.
				setlineno(l)
			vp := &nl[i2]
			if l.Op == OKEY {
				l.Left = psess.typecheck(l.Left, Erv)
				psess.
					evconst(l.Left)
				i = psess.nonnegintconst(l.Left)
				if i < 0 && !l.Left.Diag() {
					psess.
						yyerror("index must be non-negative integer constant")
					l.Left.SetDiag(true)
					i = -(1 << 30)
				}
				vp = &l.Right
			}

			if i >= 0 && indices != nil {
				if indices[i] {
					psess.
						yyerror("duplicate index in array literal: %d", i)
				} else {
					indices[i] = true
				}
			}

			r := *vp
			psess.
				pushtype(r, t.Elem(psess.types))
			r = psess.typecheck(r, Erv)
			r = psess.defaultlit(r, t.Elem(psess.types))
			*vp = psess.assignconv(r, t.Elem(psess.types), "array or slice literal")

			i++
			if i > length {
				length = i
				if checkBounds && length > t.NumElem(psess.types) {
					psess.
						setlineno(l)
					psess.
						yyerror("array index %d out of bounds [0:%d]", length-1, t.NumElem(psess.types))
					checkBounds = false
				}
			}
		}

		if t.IsDDDArray() {
			t.SetNumElem(psess.types, length)
		}
		if t.IsSlice() {
			n.Op = OSLICELIT
			n.Right = psess.nodintconst(length)
		} else {
			n.Op = OARRAYLIT
			n.Right = nil
		}

	case TMAP:
		hash := make(map[uint32][]*Node)
		for i3, l := range n.List.Slice() {
			psess.
				setlineno(l)
			if l.Op != OKEY {
				n.List.SetIndex(i3, psess.typecheck(l, Erv))
				psess.
					yyerror("missing key in map literal")
				continue
			}

			r := l.Left
			psess.
				pushtype(r, t.Key(psess.types))
			r = psess.typecheck(r, Erv)
			r = psess.defaultlit(r, t.Key(psess.types))
			l.Left = psess.assignconv(r, t.Key(psess.types), "map key")
			if l.Left.Op != OCONV {
				psess.
					keydup(l.Left, hash)
			}

			r = l.Right
			psess.
				pushtype(r, t.Elem(psess.types))
			r = psess.typecheck(r, Erv)
			r = psess.defaultlit(r, t.Elem(psess.types))
			l.Right = psess.assignconv(r, t.Elem(psess.types), "map value")
		}

		n.Op = OMAPLIT
		n.Right = nil

	case TSTRUCT:
		psess.
			dowidth(t)

		errored := false
		if n.List.Len() != 0 && nokeys(n.List) {

			ls := n.List.Slice()
			for i, n1 := range ls {
				psess.
					setlineno(n1)
				n1 = psess.typecheck(n1, Erv)
				ls[i] = n1
				if i >= t.NumFields(psess.types) {
					if !errored {
						psess.
							yyerror("too many values in %v", n)
						errored = true
					}
					continue
				}

				f := t.Field(psess.types, i)
				s := f.Sym
				if s != nil && !types.IsExported(s.Name) && s.Pkg != psess.localpkg {
					psess.
						yyerror("implicit assignment of unexported field '%s' in %v literal", s.Name, t)
				}

				n1 = psess.assignconv(n1, f.Type, "field value")
				n1 = psess.nodSym(OSTRUCTKEY, n1, f.Sym)
				n1.Xoffset = f.Offset
				ls[i] = n1
			}
			if len(ls) < t.NumFields(psess.types) {
				psess.
					yyerror("too few values in %v", n)
			}
		} else {
			hash := make(map[string]bool)

			ls := n.List.Slice()
			for i, l := range ls {
				psess.
					setlineno(l)

				if l.Op == OKEY {
					key := l.Left

					l.Op = OSTRUCTKEY
					l.Left = l.Right
					l.Right = nil

					if key.Sym == nil || key.Op == OXDOT || key.Sym.IsBlank() {
						psess.
							yyerror("invalid field name %v in struct initializer", key)
						l.Left = psess.typecheck(l.Left, Erv)
						continue
					}

					s := key.Sym
					if s.Pkg != psess.localpkg && types.IsExported(s.Name) {
						s1 := psess.lookup(s.Name)
						if s1.Origpkg == s.Pkg {
							s = s1
						}
					}
					l.Sym = s
				}

				if l.Op != OSTRUCTKEY {
					if !errored {
						psess.
							yyerror("mixture of field:value and value initializers")
						errored = true
					}
					ls[i] = psess.typecheck(ls[i], Erv)
					continue
				}

				f := psess.lookdot1(nil, l.Sym, t, t.Fields(psess.types), 0)
				if f == nil {
					if ci := psess.lookdot1(nil, l.Sym, t, t.Fields(psess.types), 2); ci != nil {
						if psess.visible(ci.Sym) {
							psess.
								yyerror("unknown field '%v' in struct literal of type %v (but does have %v)", l.Sym, t, ci.Sym)
						} else {
							psess.
								yyerror("unknown field '%v' in struct literal of type %v", l.Sym, t)
						}
						continue
					}
					p, _ := psess.dotpath(l.Sym, t, nil, true)
					if p == nil {
						psess.
							yyerror("unknown field '%v' in struct literal of type %v", l.Sym, t)
						continue
					}
					// dotpath returns the parent embedded types in reverse order.
					var ep []string
					for ei := len(p) - 1; ei >= 0; ei-- {
						ep = append(ep, p[ei].field.Type.Sym.Name)
					}
					ep = append(ep, l.Sym.Name)
					psess.
						yyerror("cannot use promoted field %v in struct literal of type %v", strings.Join(ep, "."), t)
					continue
				}
				psess.
					fielddup(f.Sym.Name, hash)
				l.Xoffset = f.Offset

				l.Left = psess.typecheck(l.Left, Erv)
				l.Left = psess.assignconv(l.Left, f.Type, "field value")
			}
		}

		n.Op = OSTRUCTLIT
		n.Right = nil
	}

	if nerr != psess.nerrors {
		return n
	}

	n.Orig = norig
	if n.Type.IsPtr() {
		n = psess.nod(OPTRLIT, n, nil)
		n.SetTypecheck(1)
		n.Type = n.Left.Type
		n.Left.Type = t
		n.Left.SetTypecheck(1)
	}

	n.Orig = norig
	return n
}

// visible reports whether sym is exported or locally defined.
func (psess *PackageSession) visible(sym *types.Sym) bool {
	return sym != nil && (types.IsExported(sym.Name) || sym.Pkg == psess.localpkg)
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

func (psess *PackageSession) checklvalue(n *Node, verb string) {
	if !islvalue(n) {
		psess.
			yyerror("cannot %s %v", verb, n)
	}
}

func (psess *PackageSession) checkassign(stmt *Node, n *Node) {

	if n.Name == nil || n.Name.Defn != stmt || stmt.Op == ORANGE {
		r := psess.outervalue(n)
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
		n.SetIndexMapLValue(psess, true)
		return
	}

	if n.Type == nil {
		return
	}

	if n.Op == ODOT && n.Left.Op == OINDEXMAP {
		psess.
			yyerror("cannot assign to struct field %v in map", n)
	} else {
		psess.
			yyerror("cannot assign to %v", n)
	}
	n.Type = nil
}

func (psess *PackageSession) checkassignlist(stmt *Node, l Nodes) {
	for _, n := range l.Slice() {
		psess.
			checkassign(stmt, n)
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
func (psess *PackageSession) samesafeexpr(l *Node, r *Node) bool {
	if l.Op != r.Op || !psess.eqtype(l.Type, r.Type) {
		return false
	}

	switch l.Op {
	case ONAME, OCLOSUREVAR:
		return l == r

	case ODOT, ODOTPTR:
		return l.Sym != nil && r.Sym != nil && l.Sym == r.Sym && psess.samesafeexpr(l.Left, r.Left)

	case OIND, OCONVNOP:
		return psess.samesafeexpr(l.Left, r.Left)

	case OCONV:

		return psess.issimple[l.Type.Etype] && psess.samesafeexpr(l.Left, r.Left)

	case OINDEX, OINDEXMAP:
		return psess.samesafeexpr(l.Left, r.Left) && psess.samesafeexpr(l.Right, r.Right)

	case OLITERAL:
		return psess.eqval(l.Val(), r.Val())
	}

	return false
}

// type check assignment.
// if this assignment is the definition of a var on the left side,
// fill in the var's type.
func (psess *PackageSession) typecheckas(n *Node) {

	n.Left = psess.resolve(n.Left)

	if n.Left.Name == nil || n.Left.Name.Defn != n || n.Left.Name.Param.Ntype != nil {
		n.Left = psess.typecheck(n.Left, Erv|Easgn)
	}

	n.Right = psess.typecheck(n.Right, Erv)
	psess.
		checkassign(n, n.Left)
	if n.Right != nil && n.Right.Type != nil {
		if n.Left.Type != nil {
			n.Right = psess.assignconv(n.Right, n.Left.Type, "assignment")
		}
	}

	if n.Left.Name != nil && n.Left.Name.Defn == n && n.Left.Name.Param.Ntype == nil {
		n.Right = psess.defaultlit(n.Right, nil)
		n.Left.Type = n.Right.Type
	}

	n.SetTypecheck(1)

	if n.Left.Typecheck() == 0 {
		n.Left = psess.typecheck(n.Left, Erv|Easgn)
	}
	if !n.Left.isBlank() {
		psess.
			checkwidth(n.Left.Type)
	}
}

func (psess *PackageSession) checkassignto(src *types.Type, dst *Node) {
	var why string

	if psess.assignop(src, dst.Type, &why) == 0 {
		psess.
			yyerror("cannot assign %v to %L in multiple assignment%s", src, dst, why)
		return
	}
}

func (psess *PackageSession) typecheckas2(n *Node) {
	ls := n.List.Slice()
	for i1, n1 := range ls {

		n1 = psess.resolve(n1)
		ls[i1] = n1

		if n1.Name == nil || n1.Name.Defn != n || n1.Name.Param.Ntype != nil {
			ls[i1] = psess.typecheck(ls[i1], Erv|Easgn)
		}
	}

	cl := n.List.Len()
	cr := n.Rlist.Len()
	if cl > 1 && cr == 1 {
		n.Rlist.SetFirst(psess.typecheck(n.Rlist.First(), Erv|Efnstruct))
	} else {
		psess.
			typecheckslice(n.Rlist.Slice(), Erv)
	}
	psess.
		checkassignlist(n, n.List)

	var l *Node
	var r *Node
	if cl == cr {

		ls := n.List.Slice()
		rs := n.Rlist.Slice()
		for il, nl := range ls {
			nr := rs[il]
			if nl.Type != nil && nr.Type != nil {
				rs[il] = psess.assignconv(nr, nl.Type, "assignment")
			}
			if nl.Name != nil && nl.Name.Defn == n && nl.Name.Param.Ntype == nil {
				rs[il] = psess.defaultlit(rs[il], nil)
				nl.Type = rs[il].Type
			}
		}

		goto out
	}

	l = n.List.First()
	r = n.Rlist.First()

	if cr == 1 {
		if r.Type == nil {
			goto out
		}
		switch r.Op {
		case OCALLMETH, OCALLINTER, OCALLFUNC:
			if !r.Type.IsFuncArgStruct() {
				break
			}
			cr = r.Type.NumFields(psess.types)
			if cr != cl {
				goto mismatch
			}
			n.Op = OAS2FUNC
			for i, l := range n.List.Slice() {
				f := r.Type.Field(psess.types, i)
				if f.Type != nil && l.Type != nil {
					psess.
						checkassignto(f.Type, l)
				}
				if l.Name != nil && l.Name.Defn == n && l.Name.Param.Ntype == nil {
					l.Type = f.Type
				}
			}
			goto out
		}
	}

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
				psess.
					checkassignto(r.Type, l)
			}
			if l.Name != nil && l.Name.Defn == n {
				l.Type = r.Type
			}
			l := n.List.Second()
			if l.Type != nil && !l.Type.IsBoolean() {
				psess.
					checkassignto(psess.types.Types[TBOOL], l)
			}
			if l.Name != nil && l.Name.Defn == n && l.Name.Param.Ntype == nil {
				l.Type = psess.types.Types[TBOOL]
			}
			goto out
		}
	}

mismatch:
	psess.
		yyerror("assignment mismatch: %d variables but %d values", cl, cr)

out:
	n.SetTypecheck(1)
	ls = n.List.Slice()
	for i1, n1 := range ls {
		if n1.Typecheck() == 0 {
			ls[i1] = psess.typecheck(ls[i1], Erv|Easgn)
		}
	}
}

// type check function definition
func (psess *PackageSession) typecheckfunc(n *Node) {
	for _, ln := range n.Func.Dcl {
		if ln.Op == ONAME && (ln.Class() == PPARAM || ln.Class() == PPARAMOUT) {
			ln.Name.Decldepth = 1
		}
	}

	n.Func.Nname = psess.typecheck(n.Func.Nname, Erv|Easgn)
	t := n.Func.Nname.Type
	if t == nil {
		return
	}
	n.Type = t
	t.FuncType(psess.types).Nname = asTypesNode(n.Func.Nname)
	rcvr := t.Recv(psess.types)
	if rcvr != nil && n.Func.Shortname != nil {
		m := psess.addmethod(n.Func.Shortname, t, true, n.Func.Pragma&Nointerface != 0)
		if m == nil {
			return
		}

		n.Func.Nname.Sym = psess.methodSym(rcvr.Type, n.Func.Shortname)
		psess.
			declare(n.Func.Nname, PFUNC)
	}

	if psess.Ctxt.Flag_dynlink && !psess.inimport && n.Func.Nname != nil {
		psess.
			makefuncsym(n.Func.Nname.Sym)
	}
}

// The result of stringtoarraylit MUST be assigned back to n, e.g.
// 	n.Left = stringtoarraylit(n.Left)
func (psess *PackageSession) stringtoarraylit(n *Node) *Node {
	if n.Left.Op != OLITERAL || n.Left.Val().Ctype(psess) != CTSTR {
		psess.
			Fatalf("stringtoarraylit %v", n)
	}

	s := n.Left.Val().U.(string)
	var l []*Node
	if n.Type.Elem(psess.types).Etype == TUINT8 {

		for i := 0; i < len(s); i++ {
			l = append(l, psess.nod(OKEY, psess.nodintconst(int64(i)), psess.nodintconst(int64(s[0]))))
		}
	} else {

		i := 0
		for _, r := range s {
			l = append(l, psess.nod(OKEY, psess.nodintconst(int64(i)), psess.nodintconst(int64(r))))
			i++
		}
	}

	nn := psess.nod(OCOMPLIT, nil, psess.typenod(n.Type))
	nn.List.Set(l)
	nn = psess.typecheck(nn, Erv)
	return nn
}

func (psess *PackageSession) checkMapKeys() {
	for _, n := range psess.mapqueue {
		k := n.Type.MapType(psess.types).Key
		if !k.Broke() && !psess.IsComparable(k) {
			psess.
				yyerrorl(n.Pos, "invalid map key type %v", k)
		}
	}
	psess.
		mapqueue = nil
}

func (psess *PackageSession) copytype(n *Node, t *types.Type) {
	if t.Etype == TFORW {

		t.ForwardType(psess.types).Copyto = append(t.ForwardType(psess.types).Copyto, asTypesNode(n))
		return
	}

	embedlineno := n.Type.ForwardType(psess.types).Embedlineno
	l := n.Type.ForwardType(psess.types).Copyto

	ptrBase := n.Type.PtrBase
	sliceOf := n.Type.SliceOf

	*n.Type = *t

	t = n.Type
	t.Sym = n.Sym
	if n.Name != nil {
		t.Vargen = n.Name.Vargen
	}

	if !t.IsInterface() {
		*t.Methods() = types.Fields{}
		*t.AllMethods() = types.Fields{}
	}

	t.Nod = asTypesNode(n)
	t.SetDeferwidth(false)
	t.PtrBase = ptrBase
	t.SliceOf = sliceOf

	if n.Name != nil && n.Name.Param != nil && n.Name.Param.Pragma&NotInHeap != 0 {
		t.SetNotInHeap(true)
	}

	for _, n := range l {
		psess.
			copytype(asNode(n), t)
	}

	lno := psess.lineno

	if embedlineno.IsKnown() {
		psess.
			lineno = embedlineno
		if t.IsPtr() || t.IsUnsafePtr() {
			psess.
				yyerror("embedded type cannot be a pointer")
		}
	}
	psess.
		lineno = lno
}

func (psess *PackageSession) typecheckdeftype(n *Node) {
	lno := psess.lineno
	psess.
		setlineno(n)
	n.Type.Sym = n.Sym
	n.SetTypecheck(1)
	n.Name.Param.Ntype = psess.typecheck(n.Name.Param.Ntype, Etype)
	t := n.Name.Param.Ntype.Type
	if t == nil {
		n.SetDiag(true)
		n.Type = nil
	} else if n.Type == nil {
		n.SetDiag(true)
	} else {
		psess.
			copytype(n, t)
	}
	psess.
		lineno = lno
}

func (psess *PackageSession) typecheckdef(n *Node) {
	lno := psess.lineno
	psess.
		setlineno(n)

	if n.Op == ONONAME {
		if !n.Diag() {
			n.SetDiag(true)
			if n.Pos.IsKnown() {
				psess.
					lineno = n.Pos
			}
			psess.
				yyerror("undefined: %v", n.Sym)
		}

		return
	}

	if n.Walkdef() == 1 {
		return
	}
	psess.
		typecheckdefstack = append(psess.typecheckdefstack, n)
	if n.Walkdef() == 2 {
		psess.
			flusherrors()
		fmt.Printf("typecheckdef loop:")
		for i := len(psess.typecheckdefstack) - 1; i >= 0; i-- {
			n := psess.typecheckdefstack[i]
			fmt.Printf(" %v", n.Sym)
		}
		fmt.Printf("\n")
		psess.
			Fatalf("typecheckdef loop")
	}

	n.SetWalkdef(2)

	if n.Type != nil || n.Sym == nil {
		goto ret
	}

	switch n.Op {
	default:
		psess.
			Fatalf("typecheckdef %v", n.Op)

	case OGOTO, OLABEL, OPACK:

	case OLITERAL:
		if n.Name.Param.Ntype != nil {
			n.Name.Param.Ntype = psess.typecheck(n.Name.Param.Ntype, Etype)
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
			psess.
				lineno = n.Pos
			Dump("typecheckdef nil defn", n)
			psess.
				yyerror("xxx")
		}

		e = psess.typecheck(e, Erv)
		if psess.Isconst(e, CTNIL) {
			psess.
				yyerror("const initializer cannot be nil")
			goto ret
		}

		if e.Type != nil && e.Op != OLITERAL || !e.isGoConst(psess) {
			if !e.Diag() {
				psess.
					yyerror("const initializer %v is not a constant", e)
				e.SetDiag(true)
			}

			goto ret
		}

		t := n.Type
		if t != nil {
			if !psess.okforconst[t.Etype] {
				psess.
					yyerror("invalid constant type %v", t)
				goto ret
			}

			if !e.Type.IsUntyped(psess.types) && !psess.eqtype(t, e.Type) {
				psess.
					yyerror("cannot use %L as type %v in const initializer", e, t)
				goto ret
			}

			e = psess.convlit(e, t)
		}

		n.SetVal(psess, e.Val())
		n.Type = e.Type

	case ONAME:
		if n.Name.Param.Ntype != nil {
			n.Name.Param.Ntype = psess.typecheck(n.Name.Param.Ntype, Etype)
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
			if n.SubOp(psess) != 0 {
				break
			}
			if psess.nsavederrors+psess.nerrors > 0 {

				break
			}
			psess.
				Fatalf("var without type, init: %v", n.Sym)
		}

		if n.Name.Defn.Op == ONAME {
			n.Name.Defn = psess.typecheck(n.Name.Defn, Erv)
			n.Type = n.Name.Defn.Type
			break
		}

		n.Name.Defn = psess.typecheck(n.Name.Defn, Etop)

	case OTYPE:
		if p := n.Name.Param; p.Alias {

			if p.Ntype != nil {
				p.Ntype = psess.typecheck(p.Ntype, Etype)
				n.Type = p.Ntype.Type
				if n.Type == nil {
					n.SetDiag(true)
					goto ret
				}
				n.Sym.Def = asTypesNode(p.Ntype)
			}
			break
		}

		if psess.Curfn != nil {
			psess.
				defercheckwidth()
		}
		n.SetWalkdef(1)
		n.Type = types.New(TFORW)
		n.Type.Nod = asTypesNode(n)
		n.Type.Sym = n.Sym
		nerrors0 := psess.nerrors
		psess.
			typecheckdeftype(n)
		if n.Type.Etype == TFORW && psess.nerrors > nerrors0 {

			n.Type.SetBroke(true)
		}
		if psess.Curfn != nil {
			psess.
				resumecheckwidth()
		}
	}

ret:
	if n.Op != OLITERAL && n.Type != nil && n.Type.IsUntyped(psess.types) {
		psess.
			Fatalf("got %v for %v", n.Type, n)
	}
	last := len(psess.typecheckdefstack) - 1
	if psess.typecheckdefstack[last] != n {
		psess.
			Fatalf("typecheckdefstack mismatch")
	}
	psess.
		typecheckdefstack[last] = nil
	psess.
		typecheckdefstack = psess.typecheckdefstack[:last]
	psess.
		lineno = lno
	n.SetWalkdef(1)
}

func (psess *PackageSession) checkmake(t *types.Type, arg string, n *Node) bool {
	if !n.Type.IsInteger() && n.Type.Etype != TIDEAL {
		psess.
			yyerror("non-integer %s argument in make(%v) - %v", arg, t, n.Type)
		return false
	}

	switch psess.consttype(n) {
	case CTINT, CTRUNE, CTFLT, CTCPLX:
		n.SetVal(psess, psess.toint(n.Val()))
		if n.Val().U.(*Mpint).CmpInt64(0) < 0 {
			psess.
				yyerror("negative %s argument in make(%v)", arg, t)
			return false
		}
		if n.Val().U.(*Mpint).Cmp(psess.maxintval[TINT]) > 0 {
			psess.
				yyerror("%s argument too large in make(%v)", arg, t)
			return false
		}
	}

	n = psess.defaultlit(n, psess.types.Types[TINT])

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
			if n1.List.Len() == 0 {
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
func (psess *PackageSession) checkreturn(fn *Node) {
	if fn.Type.NumResults(psess.types) != 0 && fn.Nbody.Len() != 0 {
		markbreaklist(fn.Nbody, nil)
		if !fn.Nbody.isterminating() {
			psess.
				yyerrorl(fn.Func.Endlineno, "missing return at end of function")
		}
	}
}

func (psess *PackageSession) deadcode(fn *Node) {
	psess.
		deadcodeslice(fn.Nbody)
}

func (psess *PackageSession) deadcodeslice(nn Nodes) {
	for i, n := range nn.Slice() {

		cut := false
		if n == nil {
			continue
		}
		if n.Op == OIF {
			n.Left = psess.deadcodeexpr(n.Left)
			if psess.Isconst(n.Left, CTBOOL) {
				var body Nodes
				if n.Left.Bool(psess) {
					n.Rlist = Nodes{}
					body = n.Nbody
				} else {
					n.Nbody = Nodes{}
					body = n.Rlist
				}

				if body := body.Slice(); len(body) != 0 {
					switch body[(len(body) - 1)].Op {
					case ORETURN, ORETJMP, OPANIC:
						cut = true
					}
				}
			}
		}
		psess.
			deadcodeslice(n.Ninit)
		psess.
			deadcodeslice(n.Nbody)
		psess.
			deadcodeslice(n.List)
		psess.
			deadcodeslice(n.Rlist)
		if cut {
			*nn.slice = nn.Slice()[:i+1]
			break
		}
	}
}

func (psess *PackageSession) deadcodeexpr(n *Node) *Node {

	switch n.Op {
	case OANDAND:
		n.Left = psess.deadcodeexpr(n.Left)
		n.Right = psess.deadcodeexpr(n.Right)
		if psess.Isconst(n.Left, CTBOOL) {
			if n.Left.Bool(psess) {
				return n.Right
			} else {
				return n.Left
			}
		}
	case OOROR:
		n.Left = psess.deadcodeexpr(n.Left)
		n.Right = psess.deadcodeexpr(n.Right)
		if psess.Isconst(n.Left, CTBOOL) {
			if n.Left.Bool(psess) {
				return n.Left
			} else {
				return n.Right
			}
		}
	}
	return n
}

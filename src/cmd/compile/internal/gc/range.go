package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/sys"
	"unicode/utf8"
)

// range
func (psess *PackageSession) typecheckrange(n *Node) {
	psess.
		typecheckrangeExpr(n)

	n.SetTypecheck(1)
	ls := n.List.Slice()
	for i1, n1 := range ls {
		if n1.Typecheck() == 0 {
			ls[i1] = psess.typecheck(ls[i1], Erv|Easgn)
		}
	}
	psess.
		decldepth++
	psess.
		typecheckslice(n.Nbody.Slice(), Etop)
	psess.
		decldepth--
}

func (psess *PackageSession) typecheckrangeExpr(n *Node) {
	n.Right = psess.typecheck(n.Right, Erv)

	t := n.Right.Type
	if t == nil {
		return
	}

	ls := n.List.Slice()
	for i1, n1 := range ls {
		if n1.Name == nil || n1.Name.Defn != n {
			ls[i1] = psess.typecheck(ls[i1], Erv|Easgn)
		}
	}

	if t.IsPtr() && t.Elem(psess.types).IsArray() {
		t = t.Elem(psess.types)
	}
	n.Type = t

	var t1, t2 *types.Type
	toomany := false
	switch t.Etype {
	default:
		psess.
			yyerrorl(n.Pos, "cannot range over %L", n.Right)
		return

	case TARRAY, TSLICE:
		t1 = psess.types.Types[TINT]
		t2 = t.Elem(psess.types)

	case TMAP:
		t1 = t.Key(psess.types)
		t2 = t.Elem(psess.types)

	case TCHAN:
		if !t.ChanDir(psess.types).CanRecv() {
			psess.
				yyerrorl(n.Pos, "invalid operation: range %v (receive from send-only type %v)", n.Right, n.Right.Type)
			return
		}

		t1 = t.Elem(psess.types)
		t2 = nil
		if n.List.Len() == 2 {
			toomany = true
		}

	case TSTRING:
		t1 = psess.types.Types[TINT]
		t2 = psess.types.Runetype
	}

	if n.List.Len() > 2 || toomany {
		psess.
			yyerrorl(n.Pos, "too many variables in range")
	}

	var v1, v2 *Node
	if n.List.Len() != 0 {
		v1 = n.List.First()
	}
	if n.List.Len() > 1 {
		v2 = n.List.Second()
	}

	if v2.isBlank() {
		if v1 != nil {
			n.List.Set1(v1)
		}
		v2 = nil
	}

	var why string
	if v1 != nil {
		if v1.Name != nil && v1.Name.Defn == n {
			v1.Type = t1
		} else if v1.Type != nil && psess.assignop(t1, v1.Type, &why) == 0 {
			psess.
				yyerrorl(n.Pos, "cannot assign type %v to %L in range%s", t1, v1, why)
		}
		psess.
			checkassign(n, v1)
	}

	if v2 != nil {
		if v2.Name != nil && v2.Name.Defn == n {
			v2.Type = t2
		} else if v2.Type != nil && psess.assignop(t2, v2.Type, &why) == 0 {
			psess.
				yyerrorl(n.Pos, "cannot assign type %v to %L in range%s", t2, v2, why)
		}
		psess.
			checkassign(n, v2)
	}
}

func (psess *PackageSession) cheapComputableIndex(width int64) bool {
	switch psess.thearch.LinkArch.Family {

	case sys.PPC64, sys.S390X:
		return width == 1
	case sys.AMD64, sys.I386, sys.ARM64, sys.ARM:
		switch width {
		case 1, 2, 4, 8:
			return true
		}
	}
	return false
}

// walkrange transforms various forms of ORANGE into
// simpler forms.  The result must be assigned back to n.
// Node n may also be modified in place, and may also be
// the returned node.
func (psess *PackageSession) walkrange(n *Node) *Node {
	if psess.isMapClear(n) {
		m := n.Right
		lno := psess.setlineno(m)
		n = psess.mapClear(m)
		psess.
			lineno = lno
		return n
	}

	t := n.Type

	a := n.Right
	lno := psess.setlineno(a)
	n.Right = nil

	var v1, v2 *Node
	l := n.List.Len()
	if l > 0 {
		v1 = n.List.First()
	}

	if l > 1 {
		v2 = n.List.Second()
	}

	if v2.isBlank() {
		v2 = nil
	}

	if v1.isBlank() && v2 == nil {
		v1 = nil
	}

	if v1 == nil && v2 != nil {
		psess.
			Fatalf("walkrange: v2 != nil while v1 == nil")
	}

	n.List.Set(nil)

	var ifGuard *Node

	translatedLoopOp := OFOR

	var body []*Node
	var init []*Node
	switch t.Etype {
	default:
		psess.
			Fatalf("walkrange")

	case TARRAY, TSLICE:
		if psess.arrayClear(n, v1, v2, a) {
			psess.
				lineno = lno
			return n
		}

		ha := a

		hv1 := psess.temp(psess.types.Types[TINT])
		hn := psess.temp(psess.types.Types[TINT])

		init = append(init, psess.nod(OAS, hv1, nil))
		init = append(init, psess.nod(OAS, hn, psess.nod(OLEN, ha, nil)))

		n.Left = psess.nod(OLT, hv1, hn)
		n.Right = psess.nod(OAS, hv1, psess.nod(OADD, hv1, psess.nodintconst(1)))

		if v1 == nil {
			break
		}

		if v2 == nil {
			body = []*Node{psess.nod(OAS, v1, hv1)}
			break
		}

		if psess.cheapComputableIndex(n.Type.Elem(psess.types).Width) {

			tmp := psess.nod(OINDEX, ha, hv1)
			tmp.SetBounded(true)

			a := psess.nod(OAS2, nil, nil)
			a.List.Set2(v1, v2)
			a.Rlist.Set2(hv1, tmp)
			body = []*Node{a}
			break
		}

		ifGuard = psess.nod(OIF, nil, nil)
		ifGuard.Left = psess.nod(OLT, hv1, hn)
		translatedLoopOp = OFORUNTIL

		hp := psess.temp(psess.types.NewPtr(n.Type.Elem(psess.types)))
		tmp := psess.nod(OINDEX, ha, psess.nodintconst(0))
		tmp.SetBounded(true)
		init = append(init, psess.nod(OAS, hp, psess.nod(OADDR, tmp, nil)))

		a := psess.nod(OAS2, nil, nil)
		a.List.Set2(v1, v2)
		a.Rlist.Set2(hv1, psess.nod(OIND, hp, nil))
		body = append(body, a)

		tmp = psess.nod(OADD, hp, psess.nodintconst(t.Elem(psess.types).Width))

		tmp.Type = hp.Type
		tmp.SetTypecheck(1)
		tmp.Right.Type = psess.types.Types[psess.types.Tptr]
		tmp.Right.SetTypecheck(1)
		a = psess.nod(OAS, hp, tmp)
		a = psess.typecheck(a, Etop)
		n.List.Set1(a)

	case TMAP:

		ha := a

		hit := psess.prealloc[n]
		th := hit.Type
		n.Left = nil
		keysym := th.Field(psess.types, 0).Sym
		valsym := th.Field(psess.types, 1).Sym

		fn := psess.syslook("mapiterinit")

		fn = psess.substArgTypes(fn, t.Key(psess.types), t.Elem(psess.types), th)
		init = append(init, psess.mkcall1(fn, nil, nil, psess.typename(t), ha, psess.nod(OADDR, hit, nil)))
		n.Left = psess.nod(ONE, psess.nodSym(ODOT, hit, keysym), psess.nodnil())

		fn = psess.syslook("mapiternext")
		fn = psess.substArgTypes(fn, th)
		n.Right = psess.mkcall1(fn, nil, nil, psess.nod(OADDR, hit, nil))

		key := psess.nodSym(ODOT, hit, keysym)
		key = psess.nod(OIND, key, nil)
		if v1 == nil {
			body = nil
		} else if v2 == nil {
			body = []*Node{psess.nod(OAS, v1, key)}
		} else {
			val := psess.nodSym(ODOT, hit, valsym)
			val = psess.nod(OIND, val, nil)
			a := psess.nod(OAS2, nil, nil)
			a.List.Set2(v1, v2)
			a.Rlist.Set2(key, val)
			body = []*Node{a}
		}

	case TCHAN:

		ha := a

		n.Left = nil

		hv1 := psess.temp(t.Elem(psess.types))
		hv1.SetTypecheck(1)
		if psess.types.Haspointers(t.Elem(psess.types)) {
			init = append(init, psess.nod(OAS, hv1, nil))
		}
		hb := psess.temp(psess.types.Types[TBOOL])

		n.Left = psess.nod(ONE, hb, psess.nodbool(false))
		a := psess.nod(OAS2RECV, nil, nil)
		a.SetTypecheck(1)
		a.List.Set2(hv1, hb)
		a.Rlist.Set1(psess.nod(ORECV, ha, nil))
		n.Left.Ninit.Set1(a)
		if v1 == nil {
			body = nil
		} else {
			body = []*Node{psess.nod(OAS, v1, hv1)}
		}

		body = append(body, psess.nod(OAS, hv1, nil))

	case TSTRING:

		ha := a

		hv1 := psess.temp(psess.types.Types[TINT])
		hv1t := psess.temp(psess.types.Types[TINT])
		hv2 := psess.temp(psess.types.Runetype)

		init = append(init, psess.nod(OAS, hv1, nil))

		n.Left = psess.nod(OLT, hv1, psess.nod(OLEN, ha, nil))

		if v1 != nil {

			body = append(body, psess.nod(OAS, hv1t, hv1))
		}

		nind := psess.nod(OINDEX, ha, hv1)
		nind.SetBounded(true)
		body = append(body, psess.nod(OAS, hv2, psess.conv(nind, psess.types.Runetype)))

		nif := psess.nod(OIF, nil, nil)
		nif.Left = psess.nod(OLT, hv2, psess.nodintconst(utf8.RuneSelf))

		nif.Nbody.Set1(psess.nod(OAS, hv1, psess.nod(OADD, hv1, psess.nodintconst(1))))

		eif := psess.nod(OAS2, nil, nil)
		nif.Rlist.Set1(eif)

		eif.List.Set2(hv2, hv1)
		fn := psess.syslook("decoderune")
		eif.Rlist.Set1(psess.mkcall1(fn, fn.Type.Results(psess.types), nil, ha, hv1))

		body = append(body, nif)

		if v1 != nil {
			if v2 != nil {

				a := psess.nod(OAS2, nil, nil)
				a.List.Set2(v1, v2)
				a.Rlist.Set2(hv1t, hv2)
				body = append(body, a)
			} else {

				body = append(body, psess.nod(OAS, v1, hv1t))
			}
		}
	}

	n.Op = translatedLoopOp
	psess.
		typecheckslice(init, Etop)

	if ifGuard != nil {
		ifGuard.Ninit.Append(init...)
		ifGuard = psess.typecheck(ifGuard, Etop)
	} else {
		n.Ninit.Append(init...)
	}
	psess.
		typecheckslice(n.Left.Ninit.Slice(), Etop)

	n.Left = psess.typecheck(n.Left, Erv)
	n.Left = psess.defaultlit(n.Left, nil)
	n.Right = psess.typecheck(n.Right, Etop)
	psess.
		typecheckslice(body, Etop)
	n.Nbody.Prepend(body...)

	if ifGuard != nil {
		ifGuard.Nbody.Set1(n)
		n = ifGuard
	}

	n = psess.walkstmt(n)
	psess.
		lineno = lno
	return n
}

// isMapClear checks if n is of the form:
//
// for k := range m {
//   delete(m, k)
// }
//
// where == for keys of map m is reflexive.
func (psess *PackageSession) isMapClear(n *Node) bool {
	if psess.Debug['N'] != 0 || psess.instrumenting {
		return false
	}

	if n.Op != ORANGE || n.Type.Etype != TMAP || n.List.Len() != 1 {
		return false
	}

	k := n.List.First()
	if k == nil || k.isBlank() {
		return false
	}

	if k.Name == nil || k.Name.Defn != n {
		return false
	}

	if n.Nbody.Len() != 1 {
		return false
	}

	stmt := n.Nbody.First()
	if stmt == nil || stmt.Op != ODELETE {
		return false
	}

	m := n.Right
	if !psess.samesafeexpr(stmt.List.First(), m) || !psess.samesafeexpr(stmt.List.Second(), k) {
		return false
	}

	if !psess.isreflexive(m.Type.Key(psess.types)) {
		return false
	}

	return true
}

// mapClear constructs a call to runtime.mapclear for the map m.
func (psess *PackageSession) mapClear(m *Node) *Node {
	t := m.Type

	fn := psess.syslook("mapclear")
	fn = psess.substArgTypes(fn, t.Key(psess.types), t.Elem(psess.types))
	n := psess.mkcall1(fn, nil, nil, psess.typename(t), m)

	n = psess.typecheck(n, Etop)
	n = psess.walkstmt(n)

	return n
}

// Lower n into runtime·memclr if possible, for
// fast zeroing of slices and arrays (issue 5373).
// Look for instances of
//
// for i := range a {
// 	a[i] = zero
// }
//
// in which the evaluation of a is side-effect-free.
//
// Parameters are as in walkrange: "for v1, v2 = range a".
func (psess *PackageSession) arrayClear(n, v1, v2, a *Node) bool {
	if psess.Debug['N'] != 0 || psess.instrumenting {
		return false
	}

	if v1 == nil || v2 != nil {
		return false
	}

	if n.Nbody.Len() != 1 || n.Nbody.First() == nil {
		return false
	}

	stmt := n.Nbody.First()
	if stmt.Op != OAS || stmt.Left.Op != OINDEX {
		return false
	}

	if !psess.samesafeexpr(stmt.Left.Left, a) || !psess.samesafeexpr(stmt.Left.Right, v1) {
		return false
	}

	elemsize := n.Type.Elem(psess.types).Width
	if elemsize <= 0 || !psess.isZero(stmt.Right) {
		return false
	}

	n.Op = OIF

	n.Nbody.Set(nil)
	n.Left = psess.nod(ONE, psess.nod(OLEN, a, nil), psess.nodintconst(0))

	hp := psess.temp(psess.types.Types[TUNSAFEPTR])

	tmp := psess.nod(OINDEX, a, psess.nodintconst(0))
	tmp.SetBounded(true)
	tmp = psess.nod(OADDR, tmp, nil)
	tmp = psess.nod(OCONVNOP, tmp, nil)
	tmp.Type = psess.types.Types[TUNSAFEPTR]
	n.Nbody.Append(psess.nod(OAS, hp, tmp))

	hn := psess.temp(psess.types.Types[TUINTPTR])

	tmp = psess.nod(OLEN, a, nil)
	tmp = psess.nod(OMUL, tmp, psess.nodintconst(elemsize))
	tmp = psess.conv(tmp, psess.types.Types[TUINTPTR])
	n.Nbody.Append(psess.nod(OAS, hn, tmp))

	var fn *Node
	if psess.types.Haspointers(a.Type.Elem(psess.types)) {

		fn = psess.mkcall("memclrHasPointers", nil, nil, hp, hn)
	} else {

		fn = psess.mkcall("memclrNoHeapPointers", nil, nil, hp, hn)
	}

	n.Nbody.Append(fn)

	v1 = psess.nod(OAS, v1, psess.nod(OSUB, psess.nod(OLEN, a, nil), psess.nodintconst(1)))

	n.Nbody.Append(v1)

	n.Left = psess.typecheck(n.Left, Erv)
	n.Left = psess.defaultlit(n.Left, nil)
	psess.
		typecheckslice(n.Nbody.Slice(), Etop)
	n = psess.walkstmt(n)
	return true
}

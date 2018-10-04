// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/sys"
	"unicode/utf8"
)

// range
func (pstate *PackageState) typecheckrange(n *Node) {
	// Typechecking order is important here:
	// 0. first typecheck range expression (slice/map/chan),
	//	it is evaluated only once and so logically it is not part of the loop.
	// 1. typcheck produced values,
	//	this part can declare new vars and so it must be typechecked before body,
	//	because body can contain a closure that captures the vars.
	// 2. decldepth++ to denote loop body.
	// 3. typecheck body.
	// 4. decldepth--.
	pstate.typecheckrangeExpr(n)

	// second half of dance, the first half being typecheckrangeExpr
	n.SetTypecheck(1)
	ls := n.List.Slice()
	for i1, n1 := range ls {
		if n1.Typecheck() == 0 {
			ls[i1] = pstate.typecheck(ls[i1], Erv|Easgn)
		}
	}

	pstate.decldepth++
	pstate.typecheckslice(n.Nbody.Slice(), Etop)
	pstate.decldepth--
}

func (pstate *PackageState) typecheckrangeExpr(n *Node) {
	n.Right = pstate.typecheck(n.Right, Erv)

	t := n.Right.Type
	if t == nil {
		return
	}
	// delicate little dance.  see typecheckas2
	ls := n.List.Slice()
	for i1, n1 := range ls {
		if n1.Name == nil || n1.Name.Defn != n {
			ls[i1] = pstate.typecheck(ls[i1], Erv|Easgn)
		}
	}

	if t.IsPtr() && t.Elem(pstate.types).IsArray() {
		t = t.Elem(pstate.types)
	}
	n.Type = t

	var t1, t2 *types.Type
	toomany := false
	switch t.Etype {
	default:
		pstate.yyerrorl(n.Pos, "cannot range over %L", n.Right)
		return

	case TARRAY, TSLICE:
		t1 = pstate.types.Types[TINT]
		t2 = t.Elem(pstate.types)

	case TMAP:
		t1 = t.Key(pstate.types)
		t2 = t.Elem(pstate.types)

	case TCHAN:
		if !t.ChanDir(pstate.types).CanRecv() {
			pstate.yyerrorl(n.Pos, "invalid operation: range %v (receive from send-only type %v)", n.Right, n.Right.Type)
			return
		}

		t1 = t.Elem(pstate.types)
		t2 = nil
		if n.List.Len() == 2 {
			toomany = true
		}

	case TSTRING:
		t1 = pstate.types.Types[TINT]
		t2 = pstate.types.Runetype
	}

	if n.List.Len() > 2 || toomany {
		pstate.yyerrorl(n.Pos, "too many variables in range")
	}

	var v1, v2 *Node
	if n.List.Len() != 0 {
		v1 = n.List.First()
	}
	if n.List.Len() > 1 {
		v2 = n.List.Second()
	}

	// this is not only a optimization but also a requirement in the spec.
	// "if the second iteration variable is the blank identifier, the range
	// clause is equivalent to the same clause with only the first variable
	// present."
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
		} else if v1.Type != nil && pstate.assignop(t1, v1.Type, &why) == 0 {
			pstate.yyerrorl(n.Pos, "cannot assign type %v to %L in range%s", t1, v1, why)
		}
		pstate.checkassign(n, v1)
	}

	if v2 != nil {
		if v2.Name != nil && v2.Name.Defn == n {
			v2.Type = t2
		} else if v2.Type != nil && pstate.assignop(t2, v2.Type, &why) == 0 {
			pstate.yyerrorl(n.Pos, "cannot assign type %v to %L in range%s", t2, v2, why)
		}
		pstate.checkassign(n, v2)
	}
}

func (pstate *PackageState) cheapComputableIndex(width int64) bool {
	switch pstate.thearch.LinkArch.Family {
	// MIPS does not have R+R addressing
	// Arm64 may lack ability to generate this code in our assembler,
	// but the architecture supports it.
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
func (pstate *PackageState) walkrange(n *Node) *Node {
	if pstate.isMapClear(n) {
		m := n.Right
		lno := pstate.setlineno(m)
		n = pstate.mapClear(m)
		pstate.lineno = lno
		return n
	}

	// variable name conventions:
	//	ohv1, hv1, hv2: hidden (old) val 1, 2
	//	ha, hit: hidden aggregate, iterator
	//	hn, hp: hidden len, pointer
	//	hb: hidden bool
	//	a, v1, v2: not hidden aggregate, val 1, 2

	t := n.Type

	a := n.Right
	lno := pstate.setlineno(a)
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
		pstate.Fatalf("walkrange: v2 != nil while v1 == nil")
	}

	// n.List has no meaning anymore, clear it
	// to avoid erroneous processing by racewalk.
	n.List.Set(nil)

	var ifGuard *Node

	translatedLoopOp := OFOR

	var body []*Node
	var init []*Node
	switch t.Etype {
	default:
		pstate.Fatalf("walkrange")

	case TARRAY, TSLICE:
		if pstate.arrayClear(n, v1, v2, a) {
			pstate.lineno = lno
			return n
		}

		// orderstmt arranged for a copy of the array/slice variable if needed.
		ha := a

		hv1 := pstate.temp(pstate.types.Types[TINT])
		hn := pstate.temp(pstate.types.Types[TINT])

		init = append(init, pstate.nod(OAS, hv1, nil))
		init = append(init, pstate.nod(OAS, hn, pstate.nod(OLEN, ha, nil)))

		n.Left = pstate.nod(OLT, hv1, hn)
		n.Right = pstate.nod(OAS, hv1, pstate.nod(OADD, hv1, pstate.nodintconst(1)))

		// for range ha { body }
		if v1 == nil {
			break
		}

		// for v1 := range ha { body }
		if v2 == nil {
			body = []*Node{pstate.nod(OAS, v1, hv1)}
			break
		}

		// for v1, v2 := range ha { body }
		if pstate.cheapComputableIndex(n.Type.Elem(pstate.types).Width) {
			// v1, v2 = hv1, ha[hv1]
			tmp := pstate.nod(OINDEX, ha, hv1)
			tmp.SetBounded(true)
			// Use OAS2 to correctly handle assignments
			// of the form "v1, a[v1] := range".
			a := pstate.nod(OAS2, nil, nil)
			a.List.Set2(v1, v2)
			a.Rlist.Set2(hv1, tmp)
			body = []*Node{a}
			break
		}

		// TODO(austin): OFORUNTIL is a strange beast, but is
		// necessary for expressing the control flow we need
		// while also making "break" and "continue" work. It
		// would be nice to just lower ORANGE during SSA, but
		// racewalk needs to see many of the operations
		// involved in ORANGE's implementation. If racewalk
		// moves into SSA, consider moving ORANGE into SSA and
		// eliminating OFORUNTIL.

		// TODO(austin): OFORUNTIL inhibits bounds-check
		// elimination on the index variable (see #20711).
		// Enhance the prove pass to understand this.
		ifGuard = pstate.nod(OIF, nil, nil)
		ifGuard.Left = pstate.nod(OLT, hv1, hn)
		translatedLoopOp = OFORUNTIL

		hp := pstate.temp(pstate.types.NewPtr(n.Type.Elem(pstate.types)))
		tmp := pstate.nod(OINDEX, ha, pstate.nodintconst(0))
		tmp.SetBounded(true)
		init = append(init, pstate.nod(OAS, hp, pstate.nod(OADDR, tmp, nil)))

		// Use OAS2 to correctly handle assignments
		// of the form "v1, a[v1] := range".
		a := pstate.nod(OAS2, nil, nil)
		a.List.Set2(v1, v2)
		a.Rlist.Set2(hv1, pstate.nod(OIND, hp, nil))
		body = append(body, a)

		// Advance pointer as part of the late increment.
		//
		// This runs *after* the condition check, so we know
		// advancing the pointer is safe and won't go past the
		// end of the allocation.
		tmp = pstate.nod(OADD, hp, pstate.nodintconst(t.Elem(pstate.types).Width))

		tmp.Type = hp.Type
		tmp.SetTypecheck(1)
		tmp.Right.Type = pstate.types.Types[pstate.types.Tptr]
		tmp.Right.SetTypecheck(1)
		a = pstate.nod(OAS, hp, tmp)
		a = pstate.typecheck(a, Etop)
		n.List.Set1(a)

	case TMAP:
		// orderstmt allocated the iterator for us.
		// we only use a once, so no copy needed.
		ha := a

		hit := pstate.prealloc[n]
		th := hit.Type
		n.Left = nil
		keysym := th.Field(pstate.types, 0).Sym // depends on layout of iterator struct.  See reflect.go:hiter
		valsym := th.Field(pstate.types, 1).Sym // ditto

		fn := pstate.syslook("mapiterinit")

		fn = pstate.substArgTypes(fn, t.Key(pstate.types), t.Elem(pstate.types), th)
		init = append(init, pstate.mkcall1(fn, nil, nil, pstate.typename(t), ha, pstate.nod(OADDR, hit, nil)))
		n.Left = pstate.nod(ONE, pstate.nodSym(ODOT, hit, keysym), pstate.nodnil())

		fn = pstate.syslook("mapiternext")
		fn = pstate.substArgTypes(fn, th)
		n.Right = pstate.mkcall1(fn, nil, nil, pstate.nod(OADDR, hit, nil))

		key := pstate.nodSym(ODOT, hit, keysym)
		key = pstate.nod(OIND, key, nil)
		if v1 == nil {
			body = nil
		} else if v2 == nil {
			body = []*Node{pstate.nod(OAS, v1, key)}
		} else {
			val := pstate.nodSym(ODOT, hit, valsym)
			val = pstate.nod(OIND, val, nil)
			a := pstate.nod(OAS2, nil, nil)
			a.List.Set2(v1, v2)
			a.Rlist.Set2(key, val)
			body = []*Node{a}
		}

	case TCHAN:
		// orderstmt arranged for a copy of the channel variable.
		ha := a

		n.Left = nil

		hv1 := pstate.temp(t.Elem(pstate.types))
		hv1.SetTypecheck(1)
		if pstate.types.Haspointers(t.Elem(pstate.types)) {
			init = append(init, pstate.nod(OAS, hv1, nil))
		}
		hb := pstate.temp(pstate.types.Types[TBOOL])

		n.Left = pstate.nod(ONE, hb, pstate.nodbool(false))
		a := pstate.nod(OAS2RECV, nil, nil)
		a.SetTypecheck(1)
		a.List.Set2(hv1, hb)
		a.Rlist.Set1(pstate.nod(ORECV, ha, nil))
		n.Left.Ninit.Set1(a)
		if v1 == nil {
			body = nil
		} else {
			body = []*Node{pstate.nod(OAS, v1, hv1)}
		}
		// Zero hv1. This prevents hv1 from being the sole, inaccessible
		// reference to an otherwise GC-able value during the next channel receive.
		// See issue 15281.
		body = append(body, pstate.nod(OAS, hv1, nil))

	case TSTRING:
		// Transform string range statements like "for v1, v2 = range a" into
		//
		// ha := a
		// for hv1 := 0; hv1 < len(ha); {
		//   hv1t := hv1
		//   hv2 := rune(ha[hv1])
		//   if hv2 < utf8.RuneSelf {
		//      hv1++
		//   } else {
		//      hv2, hv1 = decoderune(ha, hv1)
		//   }
		//   v1, v2 = hv1t, hv2
		//   // original body
		// }

		// orderstmt arranged for a copy of the string variable.
		ha := a

		hv1 := pstate.temp(pstate.types.Types[TINT])
		hv1t := pstate.temp(pstate.types.Types[TINT])
		hv2 := pstate.temp(pstate.types.Runetype)

		// hv1 := 0
		init = append(init, pstate.nod(OAS, hv1, nil))

		// hv1 < len(ha)
		n.Left = pstate.nod(OLT, hv1, pstate.nod(OLEN, ha, nil))

		if v1 != nil {
			// hv1t = hv1
			body = append(body, pstate.nod(OAS, hv1t, hv1))
		}

		// hv2 := rune(ha[hv1])
		nind := pstate.nod(OINDEX, ha, hv1)
		nind.SetBounded(true)
		body = append(body, pstate.nod(OAS, hv2, pstate.conv(nind, pstate.types.Runetype)))

		// if hv2 < utf8.RuneSelf
		nif := pstate.nod(OIF, nil, nil)
		nif.Left = pstate.nod(OLT, hv2, pstate.nodintconst(utf8.RuneSelf))

		// hv1++
		nif.Nbody.Set1(pstate.nod(OAS, hv1, pstate.nod(OADD, hv1, pstate.nodintconst(1))))

		// } else {
		eif := pstate.nod(OAS2, nil, nil)
		nif.Rlist.Set1(eif)

		// hv2, hv1 = decoderune(ha, hv1)
		eif.List.Set2(hv2, hv1)
		fn := pstate.syslook("decoderune")
		eif.Rlist.Set1(pstate.mkcall1(fn, fn.Type.Results(pstate.types), nil, ha, hv1))

		body = append(body, nif)

		if v1 != nil {
			if v2 != nil {
				// v1, v2 = hv1t, hv2
				a := pstate.nod(OAS2, nil, nil)
				a.List.Set2(v1, v2)
				a.Rlist.Set2(hv1t, hv2)
				body = append(body, a)
			} else {
				// v1 = hv1t
				body = append(body, pstate.nod(OAS, v1, hv1t))
			}
		}
	}

	n.Op = translatedLoopOp
	pstate.typecheckslice(init, Etop)

	if ifGuard != nil {
		ifGuard.Ninit.Append(init...)
		ifGuard = pstate.typecheck(ifGuard, Etop)
	} else {
		n.Ninit.Append(init...)
	}

	pstate.typecheckslice(n.Left.Ninit.Slice(), Etop)

	n.Left = pstate.typecheck(n.Left, Erv)
	n.Left = pstate.defaultlit(n.Left, nil)
	n.Right = pstate.typecheck(n.Right, Etop)
	pstate.typecheckslice(body, Etop)
	n.Nbody.Prepend(body...)

	if ifGuard != nil {
		ifGuard.Nbody.Set1(n)
		n = ifGuard
	}

	n = pstate.walkstmt(n)

	pstate.lineno = lno
	return n
}

// isMapClear checks if n is of the form:
//
// for k := range m {
//   delete(m, k)
// }
//
// where == for keys of map m is reflexive.
func (pstate *PackageState) isMapClear(n *Node) bool {
	if pstate.Debug['N'] != 0 || pstate.instrumenting {
		return false
	}

	if n.Op != ORANGE || n.Type.Etype != TMAP || n.List.Len() != 1 {
		return false
	}

	k := n.List.First()
	if k == nil || k.isBlank() {
		return false
	}

	// Require k to be a new variable name.
	if k.Name == nil || k.Name.Defn != n {
		return false
	}

	if n.Nbody.Len() != 1 {
		return false
	}

	stmt := n.Nbody.First() // only stmt in body
	if stmt == nil || stmt.Op != ODELETE {
		return false
	}

	m := n.Right
	if !pstate.samesafeexpr(stmt.List.First(), m) || !pstate.samesafeexpr(stmt.List.Second(), k) {
		return false
	}

	// Keys where equality is not reflexive can not be deleted from maps.
	if !pstate.isreflexive(m.Type.Key(pstate.types)) {
		return false
	}

	return true
}

// mapClear constructs a call to runtime.mapclear for the map m.
func (pstate *PackageState) mapClear(m *Node) *Node {
	t := m.Type

	// instantiate mapclear(typ *type, hmap map[any]any)
	fn := pstate.syslook("mapclear")
	fn = pstate.substArgTypes(fn, t.Key(pstate.types), t.Elem(pstate.types))
	n := pstate.mkcall1(fn, nil, nil, pstate.typename(t), m)

	n = pstate.typecheck(n, Etop)
	n = pstate.walkstmt(n)

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
func (pstate *PackageState) arrayClear(n, v1, v2, a *Node) bool {
	if pstate.Debug['N'] != 0 || pstate.instrumenting {
		return false
	}

	if v1 == nil || v2 != nil {
		return false
	}

	if n.Nbody.Len() != 1 || n.Nbody.First() == nil {
		return false
	}

	stmt := n.Nbody.First() // only stmt in body
	if stmt.Op != OAS || stmt.Left.Op != OINDEX {
		return false
	}

	if !pstate.samesafeexpr(stmt.Left.Left, a) || !pstate.samesafeexpr(stmt.Left.Right, v1) {
		return false
	}

	elemsize := n.Type.Elem(pstate.types).Width
	if elemsize <= 0 || !pstate.isZero(stmt.Right) {
		return false
	}

	// Convert to
	// if len(a) != 0 {
	// 	hp = &a[0]
	// 	hn = len(a)*sizeof(elem(a))
	// 	memclr{NoHeap,Has}Pointers(hp, hn)
	// 	i = len(a) - 1
	// }
	n.Op = OIF

	n.Nbody.Set(nil)
	n.Left = pstate.nod(ONE, pstate.nod(OLEN, a, nil), pstate.nodintconst(0))

	// hp = &a[0]
	hp := pstate.temp(pstate.types.Types[TUNSAFEPTR])

	tmp := pstate.nod(OINDEX, a, pstate.nodintconst(0))
	tmp.SetBounded(true)
	tmp = pstate.nod(OADDR, tmp, nil)
	tmp = pstate.nod(OCONVNOP, tmp, nil)
	tmp.Type = pstate.types.Types[TUNSAFEPTR]
	n.Nbody.Append(pstate.nod(OAS, hp, tmp))

	// hn = len(a) * sizeof(elem(a))
	hn := pstate.temp(pstate.types.Types[TUINTPTR])

	tmp = pstate.nod(OLEN, a, nil)
	tmp = pstate.nod(OMUL, tmp, pstate.nodintconst(elemsize))
	tmp = pstate.conv(tmp, pstate.types.Types[TUINTPTR])
	n.Nbody.Append(pstate.nod(OAS, hn, tmp))

	var fn *Node
	if pstate.types.Haspointers(a.Type.Elem(pstate.types)) {
		// memclrHasPointers(hp, hn)
		fn = pstate.mkcall("memclrHasPointers", nil, nil, hp, hn)
	} else {
		// memclrNoHeapPointers(hp, hn)
		fn = pstate.mkcall("memclrNoHeapPointers", nil, nil, hp, hn)
	}

	n.Nbody.Append(fn)

	// i = len(a) - 1
	v1 = pstate.nod(OAS, v1, pstate.nod(OSUB, pstate.nod(OLEN, a, nil), pstate.nodintconst(1)))

	n.Nbody.Append(v1)

	n.Left = pstate.typecheck(n.Left, Erv)
	n.Left = pstate.defaultlit(n.Left, nil)
	pstate.typecheckslice(n.Nbody.Slice(), Etop)
	n = pstate.walkstmt(n)
	return true
}

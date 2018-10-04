// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import "github.com/dave/golib/src/cmd/compile/internal/types"

// select
func (pstate *PackageState) typecheckselect(sel *Node) {
	var def *Node
	lno := pstate.setlineno(sel)
	pstate.typecheckslice(sel.Ninit.Slice(), Etop)
	for _, ncase := range sel.List.Slice() {
		if ncase.Op != OXCASE {
			pstate.setlineno(ncase)
			pstate.Fatalf("typecheckselect %v", ncase.Op)
		}

		if ncase.List.Len() == 0 {
			// default
			if def != nil {
				pstate.yyerrorl(ncase.Pos, "multiple defaults in select (first at %v)", def.Line(pstate))
			} else {
				def = ncase
			}
		} else if ncase.List.Len() > 1 {
			pstate.yyerrorl(ncase.Pos, "select cases cannot be lists")
		} else {
			ncase.List.SetFirst(pstate.typecheck(ncase.List.First(), Etop))
			n := ncase.List.First()
			ncase.Left = n
			ncase.List.Set(nil)
			switch n.Op {
			default:
				pos := n.Pos
				if n.Op == ONAME {
					// We don't have the right position for ONAME nodes (see #15459 and
					// others). Using ncase.Pos for now as it will provide the correct
					// line number (assuming the expression follows the "case" keyword
					// on the same line). This matches the approach before 1.10.
					pos = ncase.Pos
				}
				pstate.yyerrorl(pos, "select case must be receive, send or assign recv")

			// convert x = <-c into OSELRECV(x, <-c).
			// remove implicit conversions; the eventual assignment
			// will reintroduce them.
			case OAS:
				if (n.Right.Op == OCONVNOP || n.Right.Op == OCONVIFACE) && n.Right.Implicit() {
					n.Right = n.Right.Left
				}

				if n.Right.Op != ORECV {
					pstate.yyerrorl(n.Pos, "select assignment must have receive on right hand side")
					break
				}

				n.Op = OSELRECV

			// convert x, ok = <-c into OSELRECV2(x, <-c) with ntest=ok
			case OAS2RECV:
				if n.Rlist.First().Op != ORECV {
					pstate.yyerrorl(n.Pos, "select assignment must have receive on right hand side")
					break
				}

				n.Op = OSELRECV2
				n.Left = n.List.First()
				n.List.Set1(n.List.Second())
				n.Right = n.Rlist.First()
				n.Rlist.Set(nil)

			// convert <-c into OSELRECV(N, <-c)
			case ORECV:
				n = pstate.nodl(n.Pos, OSELRECV, nil, n)

				n.SetTypecheck(1)
				ncase.Left = n

			case OSEND:
				break
			}
		}

		pstate.typecheckslice(ncase.Nbody.Slice(), Etop)
	}

	pstate.lineno = lno
}

func (pstate *PackageState) walkselect(sel *Node) {
	lno := pstate.setlineno(sel)
	if sel.Nbody.Len() != 0 {
		pstate.Fatalf("double walkselect")
	}

	init := sel.Ninit.Slice()
	sel.Ninit.Set(nil)

	init = append(init, pstate.walkselectcases(&sel.List)...)
	sel.List.Set(nil)

	sel.Nbody.Set(init)
	pstate.walkstmtlist(sel.Nbody.Slice())

	pstate.lineno = lno
}

func (pstate *PackageState) walkselectcases(cases *Nodes) []*Node {
	n := cases.Len()
	sellineno := pstate.lineno

	// optimization: zero-case select
	if n == 0 {
		return []*Node{pstate.mkcall("block", nil, nil)}
	}

	// optimization: one-case select: single op.
	// TODO(rsc): Reenable optimization once order.go can handle it.
	// golang.org/issue/7672.
	if n == 1 {
		cas := cases.First()
		pstate.setlineno(cas)
		l := cas.Ninit.Slice()
		if cas.Left != nil { // not default:
			n := cas.Left
			l = append(l, n.Ninit.Slice()...)
			n.Ninit.Set(nil)
			var ch *Node
			switch n.Op {
			default:
				pstate.Fatalf("select %v", n.Op)

			// ok already
			case OSEND:
				ch = n.Left

			case OSELRECV, OSELRECV2:
				ch = n.Right.Left
				if n.Op == OSELRECV || n.List.Len() == 0 {
					if n.Left == nil {
						n = n.Right
					} else {
						n.Op = OAS
					}
					break
				}

				if n.Left == nil {
					pstate.nblank = pstate.typecheck(pstate.nblank, Erv|Easgn)
					n.Left = pstate.nblank
				}

				n.Op = OAS2
				n.List.Prepend(n.Left)
				n.Rlist.Set1(n.Right)
				n.Right = nil
				n.Left = nil
				n.SetTypecheck(0)
				n = pstate.typecheck(n, Etop)
			}

			// if ch == nil { block() }; n;
			a := pstate.nod(OIF, nil, nil)

			a.Left = pstate.nod(OEQ, ch, pstate.nodnil())
			var ln Nodes
			ln.Set(l)
			a.Nbody.Set1(pstate.mkcall("block", nil, &ln))
			l = ln.Slice()
			a = pstate.typecheck(a, Etop)
			l = append(l, a, n)
		}

		l = append(l, cas.Nbody.Slice()...)
		l = append(l, pstate.nod(OBREAK, nil, nil))
		return l
	}

	// convert case value arguments to addresses.
	// this rewrite is used by both the general code and the next optimization.
	for _, cas := range cases.Slice() {
		pstate.setlineno(cas)
		n := cas.Left
		if n == nil {
			continue
		}
		switch n.Op {
		case OSEND:
			n.Right = pstate.nod(OADDR, n.Right, nil)
			n.Right = pstate.typecheck(n.Right, Erv)

		case OSELRECV, OSELRECV2:
			if n.Op == OSELRECV2 && n.List.Len() == 0 {
				n.Op = OSELRECV
			}

			if n.Left != nil {
				n.Left = pstate.nod(OADDR, n.Left, nil)
				n.Left = pstate.typecheck(n.Left, Erv)
			}
		}
	}

	// optimization: two-case select but one is default: single non-blocking op.
	if n == 2 && (cases.First().Left == nil || cases.Second().Left == nil) {
		var cas *Node
		var dflt *Node
		if cases.First().Left == nil {
			cas = cases.Second()
			dflt = cases.First()
		} else {
			dflt = cases.Second()
			cas = cases.First()
		}

		n := cas.Left
		pstate.setlineno(n)
		r := pstate.nod(OIF, nil, nil)
		r.Ninit.Set(cas.Ninit.Slice())
		switch n.Op {
		default:
			pstate.Fatalf("select %v", n.Op)

		case OSEND:
			// if selectnbsend(c, v) { body } else { default body }
			ch := n.Left
			r.Left = pstate.mkcall1(pstate.chanfn("selectnbsend", 2, ch.Type), pstate.types.Types[TBOOL], &r.Ninit, ch, n.Right)

		case OSELRECV:
			// if selectnbrecv(&v, c) { body } else { default body }
			r = pstate.nod(OIF, nil, nil)
			r.Ninit.Set(cas.Ninit.Slice())
			ch := n.Right.Left
			elem := n.Left
			if elem == nil {
				elem = pstate.nodnil()
			}
			r.Left = pstate.mkcall1(pstate.chanfn("selectnbrecv", 2, ch.Type), pstate.types.Types[TBOOL], &r.Ninit, elem, ch)

		case OSELRECV2:
			// if selectnbrecv2(&v, &received, c) { body } else { default body }
			r = pstate.nod(OIF, nil, nil)
			r.Ninit.Set(cas.Ninit.Slice())
			ch := n.Right.Left
			elem := n.Left
			if elem == nil {
				elem = pstate.nodnil()
			}
			receivedp := pstate.nod(OADDR, n.List.First(), nil)
			receivedp = pstate.typecheck(receivedp, Erv)
			r.Left = pstate.mkcall1(pstate.chanfn("selectnbrecv2", 2, ch.Type), pstate.types.Types[TBOOL], &r.Ninit, elem, receivedp, ch)
		}

		r.Left = pstate.typecheck(r.Left, Erv)
		r.Nbody.Set(cas.Nbody.Slice())
		r.Rlist.Set(append(dflt.Ninit.Slice(), dflt.Nbody.Slice()...))
		return []*Node{r, pstate.nod(OBREAK, nil, nil)}
	}

	var init []*Node

	// generate sel-struct
	pstate.lineno = sellineno
	selv := pstate.temp(pstate.types.NewArray(pstate.scasetype(), int64(n)))
	r := pstate.nod(OAS, selv, nil)
	r = pstate.typecheck(r, Etop)
	init = append(init, r)

	order := pstate.temp(pstate.types.NewArray(pstate.types.Types[TUINT16], 2*int64(n)))
	r = pstate.nod(OAS, order, nil)
	r = pstate.typecheck(r, Etop)
	init = append(init, r)

	// register cases
	for i, cas := range cases.Slice() {
		pstate.setlineno(cas)

		init = append(init, cas.Ninit.Slice()...)
		cas.Ninit.Set(nil)

		// Keep in sync with runtime/select.go.
		const (
			caseNil = iota
			caseRecv
			caseSend
			caseDefault
		)

		var c, elem *Node
		var kind int64 = caseDefault

		if n := cas.Left; n != nil {
			init = append(init, n.Ninit.Slice()...)

			switch n.Op {
			default:
				pstate.Fatalf("select %v", n.Op)
			case OSEND:
				kind = caseSend
				c = n.Left
				elem = n.Right
			case OSELRECV, OSELRECV2:
				kind = caseRecv
				c = n.Right.Left
				elem = n.Left
			}
		}

		setField := func(f string, val *Node) {
			r := pstate.nod(OAS, pstate.nodSym(ODOT, pstate.nod(OINDEX, selv, pstate.nodintconst(int64(i))), pstate.lookup(f)), val)
			r = pstate.typecheck(r, Etop)
			init = append(init, r)
		}

		setField("kind", pstate.nodintconst(kind))
		if c != nil {
			c = pstate.nod(OCONVNOP, c, nil)
			c.Type = pstate.types.Types[TUNSAFEPTR]
			setField("c", c)
		}
		if elem != nil {
			elem = pstate.nod(OCONVNOP, elem, nil)
			elem.Type = pstate.types.Types[TUNSAFEPTR]
			setField("elem", elem)
		}

		// TODO(mdempsky): There should be a cleaner way to
		// handle this.
		if pstate.instrumenting {
			r = pstate.mkcall("selectsetpc", nil, nil, pstate.bytePtrToIndex(selv, int64(i)))
			init = append(init, r)
		}
	}

	// run the select
	pstate.lineno = sellineno
	chosen := pstate.temp(pstate.types.Types[TINT])
	recvOK := pstate.temp(pstate.types.Types[TBOOL])
	r = pstate.nod(OAS2, nil, nil)
	r.List.Set2(chosen, recvOK)
	fn := pstate.syslook("selectgo")
	r.Rlist.Set1(pstate.mkcall1(fn, fn.Type.Results(pstate.types), nil, pstate.bytePtrToIndex(selv, 0), pstate.bytePtrToIndex(order, 0), pstate.nodintconst(int64(n))))
	r = pstate.typecheck(r, Etop)
	init = append(init, r)

	// selv and order are no longer alive after selectgo.
	init = append(init, pstate.nod(OVARKILL, selv, nil))
	init = append(init, pstate.nod(OVARKILL, order, nil))

	// dispatch cases
	for i, cas := range cases.Slice() {
		pstate.setlineno(cas)

		cond := pstate.nod(OEQ, chosen, pstate.nodintconst(int64(i)))
		cond = pstate.typecheck(cond, Erv)
		cond = pstate.defaultlit(cond, nil)

		r = pstate.nod(OIF, cond, nil)

		if n := cas.Left; n != nil && n.Op == OSELRECV2 {
			x := pstate.nod(OAS, n.List.First(), recvOK)
			x = pstate.typecheck(x, Etop)
			r.Nbody.Append(x)
		}

		r.Nbody.AppendNodes(&cas.Nbody)
		r.Nbody.Append(pstate.nod(OBREAK, nil, nil))
		init = append(init, r)
	}

	return init
}

// bytePtrToIndex returns a Node representing "(*byte)(&n[i])".
func (pstate *PackageState) bytePtrToIndex(n *Node, i int64) *Node {
	s := pstate.nod(OCONVNOP, pstate.nod(OADDR, pstate.nod(OINDEX, n, pstate.nodintconst(i)), nil), nil)
	s.Type = pstate.types.NewPtr(pstate.types.Types[TUINT8])
	s = pstate.typecheck(s, Erv)
	return s
}

// Keep in sync with src/runtime/select.go.
func (pstate *PackageState) scasetype() *types.Type {
	if pstate.scase == nil {
		pstate.scase = pstate.tostruct([]*Node{
			pstate.namedfield("c", pstate.types.Types[TUNSAFEPTR]),
			pstate.namedfield("elem", pstate.types.Types[TUNSAFEPTR]),
			pstate.namedfield("kind", pstate.types.Types[TUINT16]),
			pstate.namedfield("pc", pstate.types.Types[TUINTPTR]),
			pstate.namedfield("releasetime", pstate.types.Types[TUINT64]),
		})
		pstate.scase.SetNoalg(true)
	}
	return pstate.scase
}

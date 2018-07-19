package gc

import "github.com/dave/golib/src/cmd/compile/internal/types"

// select
func (psess *PackageSession) typecheckselect(sel *Node) {
	var def *Node
	lno := psess.setlineno(sel)
	psess.
		typecheckslice(sel.Ninit.Slice(), Etop)
	for _, ncase := range sel.List.Slice() {
		if ncase.Op != OXCASE {
			psess.
				setlineno(ncase)
			psess.
				Fatalf("typecheckselect %v", ncase.Op)
		}

		if ncase.List.Len() == 0 {

			if def != nil {
				psess.
					yyerrorl(ncase.Pos, "multiple defaults in select (first at %v)", def.Line(psess))
			} else {
				def = ncase
			}
		} else if ncase.List.Len() > 1 {
			psess.
				yyerrorl(ncase.Pos, "select cases cannot be lists")
		} else {
			ncase.List.SetFirst(psess.typecheck(ncase.List.First(), Etop))
			n := ncase.List.First()
			ncase.Left = n
			ncase.List.Set(nil)
			switch n.Op {
			default:
				pos := n.Pos
				if n.Op == ONAME {

					pos = ncase.Pos
				}
				psess.
					yyerrorl(pos, "select case must be receive, send or assign recv")

			case OAS:
				if (n.Right.Op == OCONVNOP || n.Right.Op == OCONVIFACE) && n.Right.Implicit() {
					n.Right = n.Right.Left
				}

				if n.Right.Op != ORECV {
					psess.
						yyerrorl(n.Pos, "select assignment must have receive on right hand side")
					break
				}

				n.Op = OSELRECV

			case OAS2RECV:
				if n.Rlist.First().Op != ORECV {
					psess.
						yyerrorl(n.Pos, "select assignment must have receive on right hand side")
					break
				}

				n.Op = OSELRECV2
				n.Left = n.List.First()
				n.List.Set1(n.List.Second())
				n.Right = n.Rlist.First()
				n.Rlist.Set(nil)

			case ORECV:
				n = psess.nodl(n.Pos, OSELRECV, nil, n)

				n.SetTypecheck(1)
				ncase.Left = n

			case OSEND:
				break
			}
		}
		psess.
			typecheckslice(ncase.Nbody.Slice(), Etop)
	}
	psess.
		lineno = lno
}

func (psess *PackageSession) walkselect(sel *Node) {
	lno := psess.setlineno(sel)
	if sel.Nbody.Len() != 0 {
		psess.
			Fatalf("double walkselect")
	}

	init := sel.Ninit.Slice()
	sel.Ninit.Set(nil)

	init = append(init, psess.walkselectcases(&sel.List)...)
	sel.List.Set(nil)

	sel.Nbody.Set(init)
	psess.
		walkstmtlist(sel.Nbody.Slice())
	psess.
		lineno = lno
}

func (psess *PackageSession) walkselectcases(cases *Nodes) []*Node {
	n := cases.Len()
	sellineno := psess.lineno

	if n == 0 {
		return []*Node{psess.mkcall("block", nil, nil)}
	}

	if n == 1 {
		cas := cases.First()
		psess.
			setlineno(cas)
		l := cas.Ninit.Slice()
		if cas.Left != nil {
			n := cas.Left
			l = append(l, n.Ninit.Slice()...)
			n.Ninit.Set(nil)
			var ch *Node
			switch n.Op {
			default:
				psess.
					Fatalf("select %v", n.Op)

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
					psess.
						nblank = psess.typecheck(psess.nblank, Erv|Easgn)
					n.Left = psess.nblank
				}

				n.Op = OAS2
				n.List.Prepend(n.Left)
				n.Rlist.Set1(n.Right)
				n.Right = nil
				n.Left = nil
				n.SetTypecheck(0)
				n = psess.typecheck(n, Etop)
			}

			a := psess.nod(OIF, nil, nil)

			a.Left = psess.nod(OEQ, ch, psess.nodnil())
			var ln Nodes
			ln.Set(l)
			a.Nbody.Set1(psess.mkcall("block", nil, &ln))
			l = ln.Slice()
			a = psess.typecheck(a, Etop)
			l = append(l, a, n)
		}

		l = append(l, cas.Nbody.Slice()...)
		l = append(l, psess.nod(OBREAK, nil, nil))
		return l
	}

	for _, cas := range cases.Slice() {
		psess.
			setlineno(cas)
		n := cas.Left
		if n == nil {
			continue
		}
		switch n.Op {
		case OSEND:
			n.Right = psess.nod(OADDR, n.Right, nil)
			n.Right = psess.typecheck(n.Right, Erv)

		case OSELRECV, OSELRECV2:
			if n.Op == OSELRECV2 && n.List.Len() == 0 {
				n.Op = OSELRECV
			}

			if n.Left != nil {
				n.Left = psess.nod(OADDR, n.Left, nil)
				n.Left = psess.typecheck(n.Left, Erv)
			}
		}
	}

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
		psess.
			setlineno(n)
		r := psess.nod(OIF, nil, nil)
		r.Ninit.Set(cas.Ninit.Slice())
		switch n.Op {
		default:
			psess.
				Fatalf("select %v", n.Op)

		case OSEND:

			ch := n.Left
			r.Left = psess.mkcall1(psess.chanfn("selectnbsend", 2, ch.Type), psess.types.Types[TBOOL], &r.Ninit, ch, n.Right)

		case OSELRECV:

			r = psess.nod(OIF, nil, nil)
			r.Ninit.Set(cas.Ninit.Slice())
			ch := n.Right.Left
			elem := n.Left
			if elem == nil {
				elem = psess.nodnil()
			}
			r.Left = psess.mkcall1(psess.chanfn("selectnbrecv", 2, ch.Type), psess.types.Types[TBOOL], &r.Ninit, elem, ch)

		case OSELRECV2:

			r = psess.nod(OIF, nil, nil)
			r.Ninit.Set(cas.Ninit.Slice())
			ch := n.Right.Left
			elem := n.Left
			if elem == nil {
				elem = psess.nodnil()
			}
			receivedp := psess.nod(OADDR, n.List.First(), nil)
			receivedp = psess.typecheck(receivedp, Erv)
			r.Left = psess.mkcall1(psess.chanfn("selectnbrecv2", 2, ch.Type), psess.types.Types[TBOOL], &r.Ninit, elem, receivedp, ch)
		}

		r.Left = psess.typecheck(r.Left, Erv)
		r.Nbody.Set(cas.Nbody.Slice())
		r.Rlist.Set(append(dflt.Ninit.Slice(), dflt.Nbody.Slice()...))
		return []*Node{r, psess.nod(OBREAK, nil, nil)}
	}

	var init []*Node
	psess.
		lineno = sellineno
	selv := psess.temp(psess.types.NewArray(psess.scasetype(), int64(n)))
	r := psess.nod(OAS, selv, nil)
	r = psess.typecheck(r, Etop)
	init = append(init, r)

	order := psess.temp(psess.types.NewArray(psess.types.Types[TUINT16], 2*int64(n)))
	r = psess.nod(OAS, order, nil)
	r = psess.typecheck(r, Etop)
	init = append(init, r)

	for i, cas := range cases.Slice() {
		psess.
			setlineno(cas)

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
				psess.
					Fatalf("select %v", n.Op)
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
			r := psess.nod(OAS, psess.nodSym(ODOT, psess.nod(OINDEX, selv, psess.nodintconst(int64(i))), psess.lookup(f)), val)
			r = psess.typecheck(r, Etop)
			init = append(init, r)
		}

		setField("kind", psess.nodintconst(kind))
		if c != nil {
			c = psess.nod(OCONVNOP, c, nil)
			c.Type = psess.types.Types[TUNSAFEPTR]
			setField("c", c)
		}
		if elem != nil {
			elem = psess.nod(OCONVNOP, elem, nil)
			elem.Type = psess.types.Types[TUNSAFEPTR]
			setField("elem", elem)
		}

		if psess.instrumenting {
			r = psess.mkcall("selectsetpc", nil, nil, psess.bytePtrToIndex(selv, int64(i)))
			init = append(init, r)
		}
	}
	psess.
		lineno = sellineno
	chosen := psess.temp(psess.types.Types[TINT])
	recvOK := psess.temp(psess.types.Types[TBOOL])
	r = psess.nod(OAS2, nil, nil)
	r.List.Set2(chosen, recvOK)
	fn := psess.syslook("selectgo")
	r.Rlist.Set1(psess.mkcall1(fn, fn.Type.Results(psess.types), nil, psess.bytePtrToIndex(selv, 0), psess.bytePtrToIndex(order, 0), psess.nodintconst(int64(n))))
	r = psess.typecheck(r, Etop)
	init = append(init, r)

	init = append(init, psess.nod(OVARKILL, selv, nil))
	init = append(init, psess.nod(OVARKILL, order, nil))

	for i, cas := range cases.Slice() {
		psess.
			setlineno(cas)

		cond := psess.nod(OEQ, chosen, psess.nodintconst(int64(i)))
		cond = psess.typecheck(cond, Erv)
		cond = psess.defaultlit(cond, nil)

		r = psess.nod(OIF, cond, nil)

		if n := cas.Left; n != nil && n.Op == OSELRECV2 {
			x := psess.nod(OAS, n.List.First(), recvOK)
			x = psess.typecheck(x, Etop)
			r.Nbody.Append(x)
		}

		r.Nbody.AppendNodes(&cas.Nbody)
		r.Nbody.Append(psess.nod(OBREAK, nil, nil))
		init = append(init, r)
	}

	return init
}

// bytePtrToIndex returns a Node representing "(*byte)(&n[i])".
func (psess *PackageSession) bytePtrToIndex(n *Node, i int64) *Node {
	s := psess.nod(OCONVNOP, psess.nod(OADDR, psess.nod(OINDEX, n, psess.nodintconst(i)), nil), nil)
	s.Type = psess.types.NewPtr(psess.types.Types[TUINT8])
	s = psess.typecheck(s, Erv)
	return s
}

// Keep in sync with src/runtime/select.go.
func (psess *PackageSession) scasetype() *types.Type {
	if psess.scase == nil {
		psess.
			scase = psess.tostruct([]*Node{psess.
			namedfield("c", psess.types.Types[TUNSAFEPTR]), psess.
			namedfield("elem", psess.types.Types[TUNSAFEPTR]), psess.
			namedfield("kind", psess.types.Types[TUINT16]), psess.
			namedfield("pc", psess.types.Types[TUINTPTR]), psess.
			namedfield("releasetime", psess.types.Types[TUINT64]),
		})
		psess.
			scase.SetNoalg(true)
	}
	return psess.scase
}

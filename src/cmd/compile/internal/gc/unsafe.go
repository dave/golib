package gc

// evalunsafe evaluates a package unsafe operation and returns the result.
func (psess *PackageSession) evalunsafe(n *Node) int64 {
	switch n.Op {
	case OALIGNOF, OSIZEOF:
		n.Left = psess.typecheck(n.Left, Erv)
		n.Left = psess.defaultlit(n.Left, nil)
		tr := n.Left.Type
		if tr == nil {
			return 0
		}
		psess.
			dowidth(tr)
		if n.Op == OALIGNOF {
			return int64(tr.Align)
		}
		return tr.Width

	case OOFFSETOF:

		if n.Left.Op != OXDOT {
			psess.
				yyerror("invalid expression %v", n)
			return 0
		}

		n.Left.Left = psess.typecheck(n.Left.Left, Erv)
		base := n.Left.Left

		n.Left = psess.typecheck(n.Left, Erv)
		if n.Left.Type == nil {
			return 0
		}
		switch n.Left.Op {
		case ODOT, ODOTPTR:
			break
		case OCALLPART:
			psess.
				yyerror("invalid expression %v: argument is a method value", n)
			return 0
		default:
			psess.
				yyerror("invalid expression %v", n)
			return 0
		}

		// Sum offsets for dots until we reach base.
		var v int64
		for r := n.Left; r != base; r = r.Left {
			switch r.Op {
			case ODOTPTR:

				if r.Left != base {
					psess.
						yyerror("invalid expression %v: selector implies indirection of embedded %v", n, r.Left)
					return 0
				}
				fallthrough
			case ODOT:
				v += r.Xoffset
			default:
				Dump("unsafenmagic", n.Left)
				psess.
					Fatalf("impossible %#v node after dot insertion", r.Op)
			}
		}
		return v
	}
	psess.
		Fatalf("unexpected op %v", n.Op)
	return 0
}

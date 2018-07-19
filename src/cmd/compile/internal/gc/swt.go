package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"sort"
)

const (
	// expression switch
	switchKindExpr  = iota // switch a {...} or switch 5 {...}
	switchKindTrue         // switch true {...} or switch {...}
	switchKindFalse        // switch false {...}
)

const (
	binarySearchMin = 4 // minimum number of cases for binary search
	integerRangeMin = 2 // minimum size of integer ranges
)

// An exprSwitch walks an expression switch.
type exprSwitch struct {
	exprname *Node // node for the expression being switched on
	kind     int   // kind of switch statement (switchKind*)
}

// A typeSwitch walks a type switch.
type typeSwitch struct {
	hashname *Node // node for the hash of the type of the variable being switched on
	facename *Node // node for the concrete type of the variable being switched on
	okname   *Node // boolean node used for comma-ok type assertions
}

// A caseClause is a single case clause in a switch statement.
type caseClause struct {
	node    *Node  // points at case statement
	ordinal int    // position in switch
	hash    uint32 // hash of a type switch
	// isconst indicates whether this case clause is a constant,
	// for the purposes of the switch code generation.
	// For expression switches, that's generally literals (case 5:, not case x:).
	// For type switches, that's concrete types (case time.Time:), not interfaces (case io.Reader:).
	isconst bool
}

// caseClauses are all the case clauses in a switch statement.
type caseClauses struct {
	list   []caseClause // general cases
	defjmp *Node        // OGOTO for default case or OBREAK if no default case present
	niljmp *Node        // OGOTO for nil type case in a type switch
}

// typecheckswitch typechecks a switch statement.
func (psess *PackageSession) typecheckswitch(n *Node) {
	psess.
		typecheckslice(n.Ninit.Slice(), Etop)

	var nilonly string
	var top int
	var t *types.Type

	if n.Left != nil && n.Left.Op == OTYPESW {

		top = Etype
		n.Left.Right = psess.typecheck(n.Left.Right, Erv)
		t = n.Left.Right.Type
		if t != nil && !t.IsInterface() {
			psess.
				yyerrorl(n.Pos, "cannot type switch on non-interface value %L", n.Left.Right)
		}
		if v := n.Left.Left; v != nil && !v.isBlank() && n.List.Len() == 0 {
			psess.
				yyerrorl(v.Pos, "%v declared and not used", v.Sym)
		}
	} else {

		top = Erv
		if n.Left != nil {
			n.Left = psess.typecheck(n.Left, Erv)
			n.Left = psess.defaultlit(n.Left, nil)
			t = n.Left.Type
		} else {
			t = psess.types.Types[TBOOL]
		}
		if t != nil {
			switch {
			case !psess.okforeq[t.Etype]:
				psess.
					yyerrorl(n.Pos, "cannot switch on %L", n.Left)
			case t.IsSlice():
				nilonly = "slice"
			case t.IsArray() && !psess.IsComparable(t):
				psess.
					yyerrorl(n.Pos, "cannot switch on %L", n.Left)
			case t.IsStruct():
				if f := psess.IncomparableField(t); f != nil {
					psess.
						yyerrorl(n.Pos, "cannot switch on %L (struct containing %v cannot be compared)", n.Left, f.Type)
				}
			case t.Etype == TFUNC:
				nilonly = "func"
			case t.IsMap():
				nilonly = "map"
			}
		}
	}

	n.Type = t

	var def, niltype *Node
	for _, ncase := range n.List.Slice() {
		if ncase.List.Len() == 0 {

			if def != nil {
				psess.
					setlineno(ncase)
				psess.
					yyerrorl(ncase.Pos, "multiple defaults in switch (first at %v)", def.Line(psess))
			} else {
				def = ncase
			}
		} else {
			ls := ncase.List.Slice()
			for i1, n1 := range ls {
				psess.
					setlineno(n1)
				ls[i1] = psess.typecheck(ls[i1], Erv|Etype)
				n1 = ls[i1]
				if n1.Type == nil || t == nil {
					continue
				}
				psess.
					setlineno(ncase)
				switch top {

				case Erv:
					ls[i1] = psess.defaultlit(ls[i1], t)
					n1 = ls[i1]
					switch {
					case n1.Op == OTYPE:
						psess.
							yyerrorl(ncase.Pos, "type %v is not an expression", n1.Type)
					case n1.Type != nil && psess.assignop(n1.Type, t, nil) == 0 && psess.assignop(t, n1.Type, nil) == 0:
						if n.Left != nil {
							psess.
								yyerrorl(ncase.Pos, "invalid case %v in switch on %v (mismatched types %v and %v)", n1, n.Left, n1.Type, t)
						} else {
							psess.
								yyerrorl(ncase.Pos, "invalid case %v in switch (mismatched types %v and bool)", n1, n1.Type)
						}
					case nilonly != "" && !n1.isNil(psess):
						psess.
							yyerrorl(ncase.Pos, "invalid case %v in switch (can only compare %s %v to nil)", n1, nilonly, n.Left)
					case t.IsInterface() && !n1.Type.IsInterface() && !psess.IsComparable(n1.Type):
						psess.
							yyerrorl(ncase.Pos, "invalid case %L in switch (incomparable type)", n1)
					}

				case Etype:
					var missing, have *types.Field
					var ptr int
					switch {
					case n1.Op == OLITERAL && n1.Type.IsKind(TNIL):

						if niltype != nil {
							psess.
								yyerrorl(ncase.Pos, "multiple nil cases in type switch (first at %v)", niltype.Line(psess))
						} else {
							niltype = ncase
						}
					case n1.Op != OTYPE && n1.Type != nil:
						psess.
							yyerrorl(ncase.Pos, "%L is not a type", n1)

						n1 = n.Left.Right
						ls[i1] = n1
					case !n1.Type.IsInterface() && t.IsInterface() && !psess.implements(n1.Type, t, &missing, &have, &ptr):
						if have != nil && !missing.Broke() && !have.Broke() {
							psess.
								yyerrorl(ncase.Pos, "impossible type switch case: %L cannot have dynamic type %v"+
									" (wrong type for %v method)\n\thave %v%S\n\twant %v%S", n.Left.Right, n1.Type, missing.Sym, have.Sym, have.Type, missing.Sym, missing.Type)
						} else if !missing.Broke() {
							if ptr != 0 {
								psess.
									yyerrorl(ncase.Pos, "impossible type switch case: %L cannot have dynamic type %v"+
										" (%v method has pointer receiver)", n.Left.Right, n1.Type, missing.Sym)
							} else {
								psess.
									yyerrorl(ncase.Pos, "impossible type switch case: %L cannot have dynamic type %v"+
										" (missing %v method)", n.Left.Right, n1.Type, missing.Sym)
							}
						}
					}
				}
			}
		}

		if n.Type == nil || n.Type.IsUntyped(psess.types) {

			return
		}

		if top == Etype {
			ll := ncase.List
			if ncase.Rlist.Len() != 0 {
				nvar := ncase.Rlist.First()
				if ll.Len() == 1 && ll.First().Type != nil && !ll.First().Type.IsKind(TNIL) {

					nvar.Type = ll.First().Type
				} else {

					nvar.Type = n.Type
				}

				nvar = psess.typecheck(nvar, Erv|Easgn)
				ncase.Rlist.SetFirst(nvar)
			}
		}
		psess.
			typecheckslice(ncase.Nbody.Slice(), Etop)
	}
	switch top {

	case Erv:
		psess.
			checkDupExprCases(n.Left, n.List.Slice())
	}
}

// walkswitch walks a switch statement.
func (psess *PackageSession) walkswitch(sw *Node) {

	if sw.Left == nil {
		sw.Left = psess.nodbool(true)
		sw.Left = psess.typecheck(sw.Left, Erv)
		sw.Left = psess.defaultlit(sw.Left, nil)
	}

	if sw.Left.Op == OTYPESW {
		var s typeSwitch
		s.walk(psess, sw)
	} else {
		var s exprSwitch
		s.walk(psess, sw)
	}
}

// walk generates an AST implementing sw.
// sw is an expression switch.
// The AST is generally of the form of a linear
// search using if..goto, although binary search
// is used with long runs of constants.
func (s *exprSwitch) walk(psess *PackageSession, sw *Node) {

	if sw.List.Len() == 0 && sw.Nbody.Len() > 0 {
		psess.
			Fatalf("second walk of switch")
	}
	psess.
		casebody(sw, nil)

	cond := sw.Left
	sw.Left = nil

	s.kind = switchKindExpr
	if psess.Isconst(cond, CTBOOL) {
		s.kind = switchKindTrue
		if !cond.Val().U.(bool) {
			s.kind = switchKindFalse
		}
	}

	if cond.Op == OARRAYBYTESTR {
		ok := true
		for _, cas := range sw.List.Slice() {
			if cas.Op != OCASE {
				psess.
					Fatalf("switch string(byteslice) bad op: %v", cas.Op)
			}
			if cas.Left != nil && !psess.Isconst(cas.Left, CTSTR) {
				ok = false
				break
			}
		}
		if ok {
			cond.Op = OARRAYBYTESTRTMP
		}
	}

	cond = psess.walkexpr(cond, &sw.Ninit)
	t := sw.Type
	if t == nil {
		return
	}

	// convert the switch into OIF statements
	var cas []*Node
	if s.kind == switchKindTrue || s.kind == switchKindFalse {
		s.exprname = psess.nodbool(s.kind == switchKindTrue)
	} else if psess.consttype(cond) > 0 {

		s.exprname = cond
	} else {
		s.exprname = psess.temp(cond.Type)
		cas = []*Node{psess.nod(OAS, s.exprname, cond)}
		psess.
			typecheckslice(cas, Etop)
	}

	clauses := s.genCaseClauses(psess, sw.List.Slice())
	sw.List.Set(nil)
	cc := clauses.list

	for len(cc) > 0 {
		run := 1
		if psess.okforcmp[t.Etype] && cc[0].isconst {

			for ; run < len(cc) && cc[run].isconst; run++ {
			}

			sort.Sort(caseClauseByConstVal(cc[:run]))
		}

		a := s.walkCases(psess, cc[:run])
		cas = append(cas, a)
		cc = cc[run:]
	}

	if psess.nerrors == 0 {
		cas = append(cas, clauses.defjmp)
		sw.Nbody.Prepend(cas...)
		psess.
			walkstmtlist(sw.Nbody.Slice())
	}
}

// walkCases generates an AST implementing the cases in cc.
func (s *exprSwitch) walkCases(psess *PackageSession, cc []caseClause) *Node {
	if len(cc) < binarySearchMin {
		// linear search
		var cas []*Node
		for _, c := range cc {
			n := c.node
			lno := psess.setlineno(n)

			a := psess.nod(OIF, nil, nil)
			if rng := n.List.Slice(); rng != nil {

				low := psess.nod(OGE, s.exprname, rng[0])
				high := psess.nod(OLE, s.exprname, rng[1])
				a.Left = psess.nod(OANDAND, low, high)
			} else if (s.kind != switchKindTrue && s.kind != switchKindFalse) || psess.assignop(n.Left.Type, s.exprname.Type, nil) == OCONVIFACE || psess.assignop(s.exprname.Type, n.Left.Type, nil) == OCONVIFACE {
				a.Left = psess.nod(OEQ, s.exprname, n.Left)
			} else if s.kind == switchKindTrue {
				a.Left = n.Left
			} else {

				a.Left = psess.nod(ONOT, n.Left, nil)
			}
			a.Left = psess.typecheck(a.Left, Erv)
			a.Left = psess.defaultlit(a.Left, nil)
			a.Nbody.Set1(n.Right)

			cas = append(cas, a)
			psess.
				lineno = lno
		}
		return psess.liststmt(cas)
	}

	half := len(cc) / 2
	a := psess.nod(OIF, nil, nil)
	n := cc[half-1].node
	var mid *Node
	if rng := n.List.Slice(); rng != nil {
		mid = rng[1]
	} else {
		mid = n.Left
	}
	le := psess.nod(OLE, s.exprname, mid)
	if psess.Isconst(mid, CTSTR) {

		lenlt := psess.nod(OLT, psess.nod(OLEN, s.exprname, nil), psess.nod(OLEN, mid, nil))
		leneq := psess.nod(OEQ, psess.nod(OLEN, s.exprname, nil), psess.nod(OLEN, mid, nil))
		a.Left = psess.nod(OOROR, lenlt, psess.nod(OANDAND, leneq, le))
	} else {
		a.Left = le
	}
	a.Left = psess.typecheck(a.Left, Erv)
	a.Left = psess.defaultlit(a.Left, nil)
	a.Nbody.Set1(s.walkCases(psess, cc[:half]))
	a.Rlist.Set1(s.walkCases(psess, cc[half:]))
	return a
}

// casebody builds separate lists of statements and cases.
// It makes labels between cases and statements
// and deals with fallthrough, break, and unreachable statements.
func (psess *PackageSession) casebody(sw *Node, typeswvar *Node) {
	if sw.List.Len() == 0 {
		return
	}

	lno := psess.setlineno(sw)

	var cas []*Node  // cases
	var stat []*Node // statements
	var def *Node    // defaults
	br := psess.nod(OBREAK, nil, nil)

	for _, n := range sw.List.Slice() {
		psess.
			setlineno(n)
		if n.Op != OXCASE {
			psess.
				Fatalf("casebody %v", n.Op)
		}
		n.Op = OCASE
		needvar := n.List.Len() != 1 || n.List.First().Op == OLITERAL

		jmp := psess.nod(OGOTO, psess.autolabel(".s"), nil)
		switch n.List.Len() {
		case 0:

			if def != nil {
				psess.
					yyerrorl(n.Pos, "more than one default case")
			}

			n.Right = jmp
			def = n
		case 1:

			n.Left = n.List.First()
			n.Right = jmp
			n.List.Set(nil)
			cas = append(cas, n)
		default:

			if typeswvar != nil || sw.Left.Type.IsInterface() || !n.List.First().Type.IsInteger() || n.List.Len() < integerRangeMin {

				for _, n1 := range n.List.Slice() {
					cas = append(cas, psess.nod(OCASE, n1, jmp))
				}
				break
			}

			s := n.List.Slice()
			j := 0
			for j < len(s) {
				// Find a run of constants.
				var run int
				for run = j; run < len(s) && psess.Isconst(s[run], CTINT); run++ {
				}
				if run-j >= integerRangeMin {

					search := s[j:run]
					sort.Sort(constIntNodesByVal(search))
					for beg, end := 0, 1; end <= len(search); end++ {
						if end < len(search) && search[end].Int64(psess) == search[end-1].Int64(psess)+1 {
							continue
						}
						if end-beg >= integerRangeMin {

							c := psess.nod(OCASE, nil, jmp)
							c.List.Set2(search[beg], search[end-1])
							cas = append(cas, c)
						} else {

							for _, n := range search[beg:end] {
								cas = append(cas, psess.nod(OCASE, n, jmp))
							}
						}
						beg = end
					}
					j = run
				}

				for ; j < len(s) && (j < run || !psess.Isconst(s[j], CTINT)); j++ {
					cas = append(cas, psess.nod(OCASE, s[j], jmp))
				}
			}
		}

		stat = append(stat, psess.nod(OLABEL, jmp.Left, nil))
		if typeswvar != nil && needvar && n.Rlist.Len() != 0 {
			l := []*Node{psess.
				nod(ODCL, n.Rlist.First(), nil), psess.
				nod(OAS, n.Rlist.First(), typeswvar),
			}
			psess.
				typecheckslice(l, Etop)
			stat = append(stat, l...)
		}
		stat = append(stat, n.Nbody.Slice()...)

		fallIndex := len(stat) - 1
		for stat[fallIndex].Op == OVARKILL {
			fallIndex--
		}
		last := stat[fallIndex]
		if last.Op != OFALL {
			stat = append(stat, br)
		}
	}

	stat = append(stat, br)
	if def != nil {
		cas = append(cas, def)
	}

	sw.List.Set(cas)
	sw.Nbody.Set(stat)
	psess.
		lineno = lno
}

// genCaseClauses generates the caseClauses value for clauses.
func (s *exprSwitch) genCaseClauses(psess *PackageSession, clauses []*Node) caseClauses {
	var cc caseClauses
	for _, n := range clauses {
		if n.Left == nil && n.List.Len() == 0 {

			if cc.defjmp != nil {
				psess.
					Fatalf("duplicate default case not detected during typechecking")
			}
			cc.defjmp = n.Right
			continue
		}
		c := caseClause{node: n, ordinal: len(cc.list)}
		if n.List.Len() > 0 {
			c.isconst = true
		}
		switch psess.consttype(n.Left) {
		case CTFLT, CTINT, CTRUNE, CTSTR:
			c.isconst = true
		}
		cc.list = append(cc.list, c)
	}

	if cc.defjmp == nil {
		cc.defjmp = psess.nod(OBREAK, nil, nil)
	}
	return cc
}

// genCaseClauses generates the caseClauses value for clauses.
func (s *typeSwitch) genCaseClauses(psess *PackageSession, clauses []*Node) caseClauses {
	var cc caseClauses
	for _, n := range clauses {
		switch {
		case n.Left == nil:

			if cc.defjmp != nil {
				psess.
					Fatalf("duplicate default case not detected during typechecking")
			}
			cc.defjmp = n.Right
			continue
		case n.Left.Op == OLITERAL:

			if cc.niljmp != nil {
				psess.
					Fatalf("duplicate nil case not detected during typechecking")
			}
			cc.niljmp = n.Right
			continue
		}

		c := caseClause{
			node:    n,
			ordinal: len(cc.list),
			isconst: !n.Left.Type.IsInterface(),
			hash:    psess.typehash(n.Left.Type),
		}
		cc.list = append(cc.list, c)
	}

	if cc.defjmp == nil {
		cc.defjmp = psess.nod(OBREAK, nil, nil)
	}

	s.checkDupCases(psess, cc.list)
	return cc
}

func (s *typeSwitch) checkDupCases(psess *PackageSession, cc []caseClause) {
	if len(cc) < 2 {
		return
	}

	seen := make(map[uint32][]*Node)

	nn := make([]*Node, 0, len(cc))
Outer:
	for _, c := range cc {
		prev, ok := seen[c.hash]
		if !ok {

			nn = append(nn, c.node)
			seen[c.hash] = nn[len(nn)-1 : len(nn) : len(nn)]
			continue
		}
		for _, n := range prev {
			if psess.eqtype(n.Left.Type, c.node.Left.Type) {
				psess.
					yyerrorl(c.node.Pos, "duplicate case %v in type switch\n\tprevious case at %v", c.node.Left.Type, n.Line(psess))

				continue Outer
			}
		}
		seen[c.hash] = append(seen[c.hash], c.node)
	}
}

func (psess *PackageSession) checkDupExprCases(exprname *Node, clauses []*Node) {

	if exprname == nil {
		return
	}

	if !exprname.Type.IsInterface() {
		seen := make(map[interface{}]*Node)
		for _, ncase := range clauses {
			for _, n := range ncase.List.Slice() {

				if ct := psess.consttype(n); ct == 0 || ct == CTBOOL {
					continue
				}

				val := n.Val().Interface(psess)
				prev, dup := seen[val]
				if !dup {
					seen[val] = n
					continue
				}
				psess.
					yyerrorl(ncase.Pos, "duplicate case %s in switch\n\tprevious case at %v", psess.
						nodeAndVal(n), prev.Line(psess))
			}
		}
		return
	}

	// s's expression is an interface. This is fairly rare, so
	// keep this simple. Case expressions are only duplicates if
	// they have the same value and identical types.
	//
	// In general, we have to use eqtype to test type identity,
	// because == gives false negatives for anonymous types and
	// the byte/uint8 and rune/int32 builtin type aliases.
	// However, this is not a problem here, because constant
	// expressions are always untyped or have a named type, and we
	// explicitly handle the builtin type aliases below.
	//
	// This approach may need to be revisited though if we fix
	// #21866 by treating all type aliases like byte/uint8 and
	// rune/int32.
	type typeVal struct {
		typ *types.Type
		val interface{}
	}
	seen := make(map[typeVal]*Node)
	for _, ncase := range clauses {
		for _, n := range ncase.List.Slice() {
			if ct := psess.consttype(n); ct == 0 || ct == CTBOOL {
				continue
			}
			tv := typeVal{
				typ: n.Type,
				val: n.Val().Interface(psess),
			}
			switch tv.typ {
			case psess.types.Bytetype:
				tv.typ = psess.types.Types[TUINT8]
			case psess.types.Runetype:
				tv.typ = psess.types.Types[TINT32]
			}
			prev, dup := seen[tv]
			if !dup {
				seen[tv] = n
				continue
			}
			psess.
				yyerrorl(ncase.Pos, "duplicate case %s in switch\n\tprevious case at %v", psess.
					nodeAndVal(n), prev.Line(psess))
		}
	}
}

func (psess *PackageSession) nodeAndVal(n *Node) string {
	show := n.String()
	val := n.Val().Interface(psess)
	if s := fmt.Sprintf("%#v", val); show != s {
		show += " (value " + s + ")"
	}
	return show
}

// walk generates an AST that implements sw,
// where sw is a type switch.
// The AST is generally of the form of a linear
// search using if..goto, although binary search
// is used with long runs of concrete types.
func (s *typeSwitch) walk(psess *PackageSession, sw *Node) {
	cond := sw.Left
	sw.Left = nil

	if cond == nil {
		sw.List.Set(nil)
		return
	}
	if cond.Right == nil {
		psess.
			yyerrorl(sw.Pos, "type switch must have an assignment")
		return
	}

	cond.Right = psess.walkexpr(cond.Right, &sw.Ninit)
	if !cond.Right.Type.IsInterface() {
		psess.
			yyerrorl(sw.Pos, "type switch must be on an interface")
		return
	}

	var cas []*Node

	s.facename = psess.temp(cond.Right.Type)

	a := psess.nod(OAS, s.facename, cond.Right)
	a = psess.typecheck(a, Etop)
	cas = append(cas, a)

	s.okname = psess.temp(psess.types.Types[TBOOL])
	s.okname = psess.typecheck(s.okname, Erv)

	s.hashname = psess.temp(psess.types.Types[TUINT32])
	s.hashname = psess.typecheck(s.hashname, Erv)
	psess.
		casebody(sw, s.facename)

	clauses := s.genCaseClauses(psess, sw.List.Slice())
	sw.List.Set(nil)
	def := clauses.defjmp

	itab := psess.nod(OITAB, s.facename, nil)

	i := psess.nod(OIF, nil, nil)
	i.Left = psess.nod(OEQ, itab, psess.nodnil())
	if clauses.niljmp != nil {

		i.Nbody.Set1(clauses.niljmp)
	} else {

		lbl := psess.autolabel(".s")
		i.Nbody.Set1(psess.nod(OGOTO, lbl, nil))

		blk := psess.nod(OBLOCK, nil, nil)
		blk.List.Set2(psess.nod(OLABEL, lbl, nil), def)
		def = blk
	}
	i.Left = psess.typecheck(i.Left, Erv)
	i.Left = psess.defaultlit(i.Left, nil)
	cas = append(cas, i)

	h := psess.nodSym(ODOTPTR, itab, nil)
	h.Type = psess.types.Types[TUINT32]
	h.SetTypecheck(1)
	if cond.Right.Type.IsEmptyInterface(psess.types) {
		h.Xoffset = int64(2 * psess.Widthptr)
	} else {
		h.Xoffset = int64(2 * psess.Widthptr)
	}
	h.SetBounded(true)
	a = psess.nod(OAS, s.hashname, h)
	a = psess.typecheck(a, Etop)
	cas = append(cas, a)

	cc := clauses.list

	for _, c := range cc {
		c.node.Right = s.typeone(psess, c.node)
	}

	for len(cc) > 0 {
		if !cc[0].isconst {
			n := cc[0].node
			cas = append(cas, n.Right)
			cc = cc[1:]
			continue
		}

		// identify run of constants
		var run int
		for run = 1; run < len(cc) && cc[run].isconst; run++ {
		}

		sort.Sort(caseClauseByType(cc[:run]))

		if false {
			for i := 0; i < run; i++ {
				n := cc[i].node
				cas = append(cas, n.Right)
			}
			continue
		}

		// combine adjacent cases with the same hash
		var batch []caseClause
		for i, j := 0, 0; i < run; i = j {
			hash := []*Node{cc[i].node.Right}
			for j = i + 1; j < run && cc[i].hash == cc[j].hash; j++ {
				hash = append(hash, cc[j].node.Right)
			}
			cc[i].node.Right = psess.liststmt(hash)
			batch = append(batch, cc[i])
		}

		cas = append(cas, s.walkCases(psess, batch))
		cc = cc[run:]
	}

	if psess.nerrors == 0 {
		cas = append(cas, def)
		sw.Nbody.Prepend(cas...)
		sw.List.Set(nil)
		psess.
			walkstmtlist(sw.Nbody.Slice())
	}
}

// typeone generates an AST that jumps to the
// case body if the variable is of type t.
func (s *typeSwitch) typeone(psess *PackageSession, t *Node) *Node {
	var name *Node
	var init Nodes
	if t.Rlist.Len() == 0 {
		name = psess.nblank
		psess.
			nblank = psess.typecheck(psess.nblank, Erv|Easgn)
	} else {
		name = t.Rlist.First()
		init.Append(psess.nod(ODCL, name, nil))
		a := psess.nod(OAS, name, nil)
		a = psess.typecheck(a, Etop)
		init.Append(a)
	}

	a := psess.nod(OAS2, nil, nil)
	a.List.Set2(name, s.okname)
	b := psess.nod(ODOTTYPE, s.facename, nil)
	b.Type = t.Left.Type
	a.Rlist.Set1(b)
	a = psess.typecheck(a, Etop)
	a = psess.walkexpr(a, &init)
	init.Append(a)

	c := psess.nod(OIF, nil, nil)
	c.Left = s.okname
	c.Nbody.Set1(t.Right)

	init.Append(c)
	return init.asblock(psess)
}

// walkCases generates an AST implementing the cases in cc.
func (s *typeSwitch) walkCases(psess *PackageSession, cc []caseClause) *Node {
	if len(cc) < binarySearchMin {
		var cas []*Node
		for _, c := range cc {
			n := c.node
			if !c.isconst {
				psess.
					Fatalf("typeSwitch walkCases")
			}
			a := psess.nod(OIF, nil, nil)
			a.Left = psess.nod(OEQ, s.hashname, psess.nodintconst(int64(c.hash)))
			a.Left = psess.typecheck(a.Left, Erv)
			a.Left = psess.defaultlit(a.Left, nil)
			a.Nbody.Set1(n.Right)
			cas = append(cas, a)
		}
		return psess.liststmt(cas)
	}

	half := len(cc) / 2
	a := psess.nod(OIF, nil, nil)
	a.Left = psess.nod(OLE, s.hashname, psess.nodintconst(int64(cc[half-1].hash)))
	a.Left = psess.typecheck(a.Left, Erv)
	a.Left = psess.defaultlit(a.Left, nil)
	a.Nbody.Set1(s.walkCases(psess, cc[:half]))
	a.Rlist.Set1(s.walkCases(psess, cc[half:]))
	return a
}

// caseClauseByConstVal sorts clauses by constant value to enable binary search.
type caseClauseByConstVal []caseClause

func (x caseClauseByConstVal) Len() int      { return len(x) }
func (x caseClauseByConstVal) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x caseClauseByConstVal) Less(psess *PackageSession, i, j int) bool {

	n1 := x[i].node
	var v1 interface{}
	if s := n1.List.Slice(); s != nil {
		v1 = s[0].Val().U
	} else {
		v1 = n1.Left.Val().U
	}

	n2 := x[j].node
	var v2 interface{}
	if s := n2.List.Slice(); s != nil {
		v2 = s[0].Val().U
	} else {
		v2 = n2.Left.Val().U
	}

	switch v1 := v1.(type) {
	case *Mpflt:
		return v1.Cmp(v2.(*Mpflt)) < 0
	case *Mpint:
		return v1.Cmp(v2.(*Mpint)) < 0
	case string:

		a := v1
		b := v2.(string)
		if len(a) != len(b) {
			return len(a) < len(b)
		}
		return a < b
	}
	psess.
		Fatalf("caseClauseByConstVal passed bad clauses %v < %v", x[i].node.Left, x[j].node.Left)
	return false
}

type caseClauseByType []caseClause

func (x caseClauseByType) Len() int      { return len(x) }
func (x caseClauseByType) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x caseClauseByType) Less(i, j int) bool {
	c1, c2 := x[i], x[j]

	if c1.hash != c2.hash {
		return c1.hash < c2.hash
	}
	return c1.ordinal < c2.ordinal
}

type constIntNodesByVal []*Node

func (x constIntNodesByVal) Len() int      { return len(x) }
func (x constIntNodesByVal) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x constIntNodesByVal) Less(i, j int) bool {
	return x[i].Val().U.(*Mpint).Cmp(x[j].Val().U.(*Mpint)) < 0
}

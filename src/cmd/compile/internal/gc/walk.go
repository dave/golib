package gc

import (
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"

	"github.com/dave/golib/src/cmd/internal/sys"
	"strings"
)

// The constant is known to runtime.
const tmpstringbufsize = 32

func (psess *PackageSession) walk(fn *Node) {
	psess.
		Curfn = fn

	if psess.Debug['W'] != 0 {
		s := fmt.Sprintf("\nbefore walk %v", psess.Curfn.Func.Nname.Sym)
		dumplist(s, psess.Curfn.Nbody)
	}

	lno := psess.lineno

	for i, ln := range fn.Func.Dcl {
		if ln.Op == ONAME && (ln.Class() == PAUTO || ln.Class() == PAUTOHEAP) {
			ln = psess.typecheck(ln, Erv|Easgn)
			fn.Func.Dcl[i] = ln
		}
	}

	for _, ln := range fn.Func.Dcl {
		if ln.Op == ONAME && (ln.Class() == PAUTO || ln.Class() == PAUTOHEAP) && ln.Name.Defn != nil && ln.Name.Defn.Op == OTYPESW && ln.Name.Used() {
			ln.Name.Defn.Left.Name.SetUsed(true)
		}
	}

	for _, ln := range fn.Func.Dcl {
		if ln.Op != ONAME || (ln.Class() != PAUTO && ln.Class() != PAUTOHEAP) || ln.Sym.Name[0] == '&' || ln.Name.Used() {
			continue
		}
		if defn := ln.Name.Defn; defn != nil && defn.Op == OTYPESW {
			if defn.Left.Name.Used() {
				continue
			}
			psess.
				yyerrorl(defn.Left.Pos, "%v declared and not used", ln.Sym)
			defn.Left.Name.SetUsed(true)
		} else {
			psess.
				yyerrorl(ln.Pos, "%v declared and not used", ln.Sym)
		}
	}
	psess.
		lineno = lno
	if psess.nerrors != 0 {
		return
	}
	psess.
		walkstmtlist(psess.Curfn.Nbody.Slice())
	if psess.Debug['W'] != 0 {
		s := fmt.Sprintf("after walk %v", psess.Curfn.Func.Nname.Sym)
		dumplist(s, psess.Curfn.Nbody)
	}
	psess.
		zeroResults()
	psess.
		heapmoves()
	if psess.Debug['W'] != 0 && psess.Curfn.Func.Enter.Len() > 0 {
		s := fmt.Sprintf("enter %v", psess.Curfn.Func.Nname.Sym)
		dumplist(s, psess.Curfn.Func.Enter)
	}
}

func (psess *PackageSession) walkstmtlist(s []*Node) {
	for i := range s {
		s[i] = psess.walkstmt(s[i])
	}
}

func samelist(a, b []*Node) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := range a {
		if n != b[i] {
			return false
		}
	}
	return true
}

func paramoutheap(fn *Node) bool {
	for _, ln := range fn.Func.Dcl {
		switch ln.Class() {
		case PPARAMOUT:
			if ln.isParamStackCopy() || ln.Addrtaken() {
				return true
			}

		case PAUTO:

			return false
		}
	}

	return false
}

// adds "adjust" to all the argument locations for the call n.
// n must be a defer or go node that has already been walked.
func (psess *PackageSession) adjustargs(n *Node, adjust int) {
	callfunc := n.Left
	for _, arg := range callfunc.List.Slice() {
		if arg.Op != OAS {
			psess.
				Fatalf("call arg not assignment")
		}
		lhs := arg.Left
		if lhs.Op == ONAME {

			continue
		}

		if lhs.Op != OINDREGSP {
			psess.
				Fatalf("call argument store does not use OINDREGSP")
		}

		lhs.Xoffset += int64(adjust)
	}
}

// The result of walkstmt MUST be assigned back to n, e.g.
// 	n.Left = walkstmt(n.Left)
func (psess *PackageSession) walkstmt(n *Node) *Node {
	if n == nil {
		return n
	}
	psess.
		setlineno(n)
	psess.
		walkstmtlist(n.Ninit.Slice())

	switch n.Op {
	default:
		if n.Op == ONAME {
			psess.
				yyerror("%v is not a top level statement", n.Sym)
		} else {
			psess.
				yyerror("%v is not a top level statement", n.Op)
		}
		Dump("nottop", n)

	case OAS,
		OASOP,
		OAS2,
		OAS2DOTTYPE,
		OAS2RECV,
		OAS2FUNC,
		OAS2MAPR,
		OCLOSE,
		OCOPY,
		OCALLMETH,
		OCALLINTER,
		OCALL,
		OCALLFUNC,
		ODELETE,
		OSEND,
		OPRINT,
		OPRINTN,
		OPANIC,
		OEMPTY,
		ORECOVER,
		OGETG:
		if n.Typecheck() == 0 {
			psess.
				Fatalf("missing typecheck: %+v", n)
		}
		wascopy := n.Op == OCOPY
		init := n.Ninit
		n.Ninit.Set(nil)
		n = psess.walkexpr(n, &init)
		n = psess.addinit(n, init.Slice())
		if wascopy && n.Op == OCONVNOP {
			n.Op = OEMPTY
		}

	case ORECV:
		if n.Typecheck() == 0 {
			psess.
				Fatalf("missing typecheck: %+v", n)
		}
		init := n.Ninit
		n.Ninit.Set(nil)

		n.Left = psess.walkexpr(n.Left, &init)
		n = psess.mkcall1(psess.chanfn("chanrecv1", 2, n.Left.Type), nil, &init, n.Left, psess.nodnil())
		n = psess.walkexpr(n, &init)

		n = psess.addinit(n, init.Slice())

	case OBREAK,
		OCONTINUE,
		OFALL,
		OGOTO,
		OLABEL,
		ODCLCONST,
		ODCLTYPE,
		OCHECKNIL,
		OVARKILL,
		OVARLIVE:
		break

	case ODCL:
		v := n.Left
		if v.Class() == PAUTOHEAP {
			if psess.compiling_runtime {
				psess.
					yyerror("%v escapes to heap, not allowed in runtime.", v)
			}
			if psess.prealloc[v] == nil {
				psess.
					prealloc[v] = psess.callnew(v.Type)
			}
			nn := psess.nod(OAS, v.Name.Param.Heapaddr, psess.prealloc[v])
			nn.SetColas(true)
			nn = psess.typecheck(nn, Etop)
			return psess.walkstmt(nn)
		}

	case OBLOCK:
		psess.
			walkstmtlist(n.List.Slice())

	case OXCASE:
		psess.
			yyerror("case statement out of place")
		n.Op = OCASE
		fallthrough

	case OCASE:
		n.Right = psess.walkstmt(n.Right)

	case ODEFER:
		psess.
			Curfn.Func.SetHasDefer(true)
		fallthrough
	case OPROC:
		switch n.Left.Op {
		case OPRINT, OPRINTN:
			n.Left = psess.wrapCall(n.Left, &n.Ninit)

		case ODELETE:
			if psess.mapfast(n.Left.List.First().Type) == mapslow {
				n.Left = psess.wrapCall(n.Left, &n.Ninit)
			} else {
				n.Left = psess.walkexpr(n.Left, &n.Ninit)
			}

		case OCOPY:
			n.Left = psess.copyany(n.Left, &n.Ninit, true)

		default:
			n.Left = psess.walkexpr(n.Left, &n.Ninit)
		}
		psess.
			adjustargs(n, 2*psess.Widthptr)

	case OFOR, OFORUNTIL:
		if n.Left != nil {
			psess.
				walkstmtlist(n.Left.Ninit.Slice())
			init := n.Left.Ninit
			n.Left.Ninit.Set(nil)
			n.Left = psess.walkexpr(n.Left, &init)
			n.Left = psess.addinit(n.Left, init.Slice())
		}

		n.Right = psess.walkstmt(n.Right)
		if n.Op == OFORUNTIL {
			psess.
				walkstmtlist(n.List.Slice())
		}
		psess.
			walkstmtlist(n.Nbody.Slice())

	case OIF:
		n.Left = psess.walkexpr(n.Left, &n.Ninit)
		psess.
			walkstmtlist(n.Nbody.Slice())
		psess.
			walkstmtlist(n.Rlist.Slice())

	case ORETURN:
		if n.List.Len() == 0 {
			break
		}
		if (psess.Curfn.Type.FuncType(psess.types).Outnamed && n.List.Len() > 1) || paramoutheap(psess.Curfn) {
			// assign to the function out parameters,
			// so that reorder3 can fix up conflicts
			var rl []*Node

			for _, ln := range psess.Curfn.Func.Dcl {
				cl := ln.Class()
				if cl == PAUTO || cl == PAUTOHEAP {
					break
				}
				if cl == PPARAMOUT {
					if ln.isParamStackCopy() {
						ln = psess.walkexpr(psess.typecheck(psess.nod(OIND, ln.Name.Param.Heapaddr, nil), Erv), nil)
					}
					rl = append(rl, ln)
				}
			}

			if got, want := n.List.Len(), len(rl); got != want {
				psess.
					Fatalf("expected %v return arguments, have %v", want, got)
			}

			if samelist(rl, n.List.Slice()) {
				psess.
					walkexprlist(n.List.Slice(), &n.Ninit)
				n.List.Set(nil)

				break
			}
			psess.
				walkexprlistsafe(n.List.Slice(), &n.Ninit)

			ll := psess.ascompatee(n.Op, rl, n.List.Slice(), &n.Ninit)
			n.List.Set(psess.reorder3(ll))
			break
		}
		psess.
			walkexprlist(n.List.Slice(), &n.Ninit)

		ll := psess.ascompatte(nil, false, psess.Curfn.Type.Results(psess.types), n.List.Slice(), 1, &n.Ninit)
		n.List.Set(ll)

	case ORETJMP:
		break

	case OSELECT:
		psess.
			walkselect(n)

	case OSWITCH:
		psess.
			walkswitch(n)

	case ORANGE:
		n = psess.walkrange(n)
	}

	if n.Op == ONAME {
		psess.
			Fatalf("walkstmt ended up with name: %+v", n)
	}
	return n
}

func (psess *PackageSession) isSmallMakeSlice(n *Node) bool {
	if n.Op != OMAKESLICE {
		return false
	}
	l := n.Left
	r := n.Right
	if r == nil {
		r = l
	}
	t := n.Type

	return psess.smallintconst(l) && psess.smallintconst(r) && (t.Elem(psess.types).Width == 0 || r.Int64(psess) < (1<<16)/t.Elem(psess.types).Width)
}

// walk the whole tree of the body of an
// expression or simple statement.
// the types expressions are calculated.
// compile-time constants are evaluated.
// complex side effects like statements are appended to init
func (psess *PackageSession) walkexprlist(s []*Node, init *Nodes) {
	for i := range s {
		s[i] = psess.walkexpr(s[i], init)
	}
}

func (psess *PackageSession) walkexprlistsafe(s []*Node, init *Nodes) {
	for i, n := range s {
		s[i] = psess.safeexpr(n, init)
		s[i] = psess.walkexpr(s[i], init)
	}
}

func (psess *PackageSession) walkexprlistcheap(s []*Node, init *Nodes) {
	for i, n := range s {
		s[i] = psess.cheapexpr(n, init)
		s[i] = psess.walkexpr(s[i], init)
	}
}

// convFuncName builds the runtime function name for interface conversion.
// It also reports whether the function expects the data by address.
// Not all names are possible. For example, we never generate convE2E or convE2I.
func (psess *PackageSession) convFuncName(from, to *types.Type) (fnname string, needsaddr bool) {
	tkind := to.Tie(psess.types)
	switch from.Tie(psess.types) {
	case 'I':
		switch tkind {
		case 'I':
			return "convI2I", false
		}
	case 'T':
		switch tkind {
		case 'E':
			switch {
			case from.Size(psess.types) == 2 && from.Align == 2:
				return "convT2E16", false
			case from.Size(psess.types) == 4 && from.Align == 4 && !psess.types.Haspointers(from):
				return "convT2E32", false
			case from.Size(psess.types) == 8 && from.Align == psess.types.Types[TUINT64].Align && !psess.types.Haspointers(from):
				return "convT2E64", false
			case from.IsString():
				return "convT2Estring", true
			case from.IsSlice():
				return "convT2Eslice", true
			case !psess.types.Haspointers(from):
				return "convT2Enoptr", true
			}
			return "convT2E", true
		case 'I':
			switch {
			case from.Size(psess.types) == 2 && from.Align == 2:
				return "convT2I16", false
			case from.Size(psess.types) == 4 && from.Align == 4 && !psess.types.Haspointers(from):
				return "convT2I32", false
			case from.Size(psess.types) == 8 && from.Align == psess.types.Types[TUINT64].Align && !psess.types.Haspointers(from):
				return "convT2I64", false
			case from.IsString():
				return "convT2Istring", true
			case from.IsSlice():
				return "convT2Islice", true
			case !psess.types.Haspointers(from):
				return "convT2Inoptr", true
			}
			return "convT2I", true
		}
	}
	psess.
		Fatalf("unknown conv func %c2%c", from.Tie(psess.types), to.Tie(psess.types))
	panic("unreachable")
}

// The result of walkexpr MUST be assigned back to n, e.g.
// 	n.Left = walkexpr(n.Left, init)
func (psess *PackageSession) walkexpr(n *Node, init *Nodes) *Node {
	if n == nil {
		return n
	}

	if n.Type != nil && !n.Type.WidthCalculated() {
		switch n.Type.Etype {
		case TBLANK, TNIL, TIDEAL:
		default:
			psess.
				checkwidth(n.Type)
		}
	}

	if init == &n.Ninit {
		psess.
			Fatalf("walkexpr init == &n->ninit")
	}

	if n.Ninit.Len() != 0 {
		psess.
			walkstmtlist(n.Ninit.Slice())
		init.AppendNodes(&n.Ninit)
	}

	lno := psess.setlineno(n)

	if psess.Debug['w'] > 1 {
		Dump("before walk expr", n)
	}

	if n.Typecheck() != 1 {
		psess.
			Fatalf("missed typecheck: %+v", n)
	}

	if n.Type.IsUntyped(psess.types) {
		psess.
			Fatalf("expression has untyped type: %+v", n)
	}

	if n.Op == ONAME && n.Class() == PAUTOHEAP {
		nn := psess.nod(OIND, n.Name.Param.Heapaddr, nil)
		nn = psess.typecheck(nn, Erv)
		nn = psess.walkexpr(nn, init)
		nn.Left.SetNonNil(true)
		return nn
	}

opswitch:
	switch n.Op {
	default:
		Dump("walk", n)
		psess.
			Fatalf("walkexpr: switch 1 unknown op %+S", n)

	case ONONAME, OINDREGSP, OEMPTY, OGETG:

	case OTYPE, ONAME, OLITERAL:

	case ONOT, OMINUS, OPLUS, OCOM, OREAL, OIMAG, ODOTMETH, ODOTINTER,
		OIND, OSPTR, OITAB, OIDATA, OADDR:
		n.Left = psess.walkexpr(n.Left, init)

	case OEFACE, OAND, OSUB, OMUL, OLT, OLE, OGE, OGT, OADD, OOR, OXOR:
		n.Left = psess.walkexpr(n.Left, init)
		n.Right = psess.walkexpr(n.Right, init)

	case ODOT:
		psess.
			usefield(n)
		n.Left = psess.walkexpr(n.Left, init)

	case ODOTTYPE, ODOTTYPE2:
		n.Left = psess.walkexpr(n.Left, init)

		n.Right = psess.typename(n.Type)
		if n.Op == ODOTTYPE {
			n.Right.Right = psess.typename(n.Left.Type)
		}
		if !n.Type.IsInterface() && !n.Left.Type.IsEmptyInterface(psess.types) {
			n.List.Set1(psess.itabname(n.Type, n.Left.Type))
		}

	case ODOTPTR:
		psess.
			usefield(n)
		if n.Op == ODOTPTR && n.Left.Type.Elem(psess.types).Width == 0 {

			n.Left = psess.cheapexpr(n.Left, init)
			psess.
				checknil(n.Left, init)
		}

		n.Left = psess.walkexpr(n.Left, init)

	case OLEN, OCAP:
		if psess.isRuneCount(n) {

			n = psess.mkcall("countrunes", n.Type, init, psess.conv(n.Left.Left, psess.types.Types[TSTRING]))
			break
		}

		n.Left = psess.walkexpr(n.Left, init)

		t := n.Left.Type

		if t.IsPtr() {
			t = t.Elem(psess.types)
		}
		if t.IsArray() {
			psess.
				safeexpr(n.Left, init)
			psess.
				setintconst(n, t.NumElem(psess.types))
			n.SetTypecheck(1)
		}

	case OLSH, ORSH:
		n.Left = psess.walkexpr(n.Left, init)
		n.Right = psess.walkexpr(n.Right, init)
		t := n.Left.Type
		n.SetBounded(psess.bounded(n.Right, 8*t.Width))
		if psess.Debug['m'] != 0 && n.Bounded() && !psess.Isconst(n.Right, CTINT) {
			psess.
				Warn("shift bounds check elided")
		}

	case OCOMPLEX:

		if n.Left == nil && n.Right == nil {
			n.Left = n.List.First()
			n.Right = n.List.Second()
		}
		n.Left = psess.walkexpr(n.Left, init)
		n.Right = psess.walkexpr(n.Right, init)

	case OEQ, ONE:
		n.Left = psess.walkexpr(n.Left, init)
		n.Right = psess.walkexpr(n.Right, init)

		old_safemode := psess.safemode
		psess.
			safemode = false
		n = psess.walkcompare(n, init)
		psess.
			safemode = old_safemode

	case OANDAND, OOROR:
		n.Left = psess.walkexpr(n.Left, init)

		// cannot put side effects from n.Right on init,
		// because they cannot run before n.Left is checked.
		// save elsewhere and store on the eventual n.Right.
		var ll Nodes

		n.Right = psess.walkexpr(n.Right, &ll)
		n.Right = psess.addinit(n.Right, ll.Slice())
		n = psess.walkinrange(n, init)

	case OPRINT, OPRINTN:
		psess.
			walkexprlist(n.List.Slice(), init)
		n = psess.walkprint(n, init)

	case OPANIC:
		n = psess.mkcall("gopanic", nil, init, n.Left)

	case ORECOVER:
		n = psess.mkcall("gorecover", n.Type, init, psess.nod(OADDR, psess.nodfp, nil))

	case OCLOSUREVAR, OCFUNC:
		n.SetAddable(true)

	case OCALLINTER:
		psess.
			usemethod(n)
		t := n.Left.Type
		if n.List.Len() != 0 && n.List.First().Op == OAS {
			break
		}
		n.Left = psess.walkexpr(n.Left, init)
		psess.
			walkexprlist(n.List.Slice(), init)
		ll := psess.ascompatte(n, n.Isddd(), t.Params(psess.types), n.List.Slice(), 0, init)
		n.List.Set(psess.reorder1(ll))

	case OCALLFUNC:
		if n.Left.Op == OCLOSURE {

			n.List.Prepend(n.Left.Func.Enter.Slice()...)

			n.Left.Func.Enter.Set(nil)

			n.Left = n.Left.Func.Closure.Func.Nname

			if n.Left.Type.NumResults(psess.types) == 1 {
				n.Type = n.Left.Type.Results(psess.types).Field(psess.types, 0).Type
			} else {
				n.Type = n.Left.Type.Results(psess.types)
			}
		}

		t := n.Left.Type
		if n.List.Len() != 0 && n.List.First().Op == OAS {
			break
		}

		n.Left = psess.walkexpr(n.Left, init)
		psess.
			walkexprlist(n.List.Slice(), init)

		ll := psess.ascompatte(n, n.Isddd(), t.Params(psess.types), n.List.Slice(), 0, init)
		n.List.Set(psess.reorder1(ll))

	case OCALLMETH:
		t := n.Left.Type
		if n.List.Len() != 0 && n.List.First().Op == OAS {
			break
		}
		n.Left = psess.walkexpr(n.Left, init)
		psess.
			walkexprlist(n.List.Slice(), init)
		ll := psess.ascompatte(n, false, t.Recvs(psess.types), []*Node{n.Left.Left}, 0, init)
		lr := psess.ascompatte(n, n.Isddd(), t.Params(psess.types), n.List.Slice(), 0, init)
		ll = append(ll, lr...)
		n.Left.Left = nil
		psess.
			updateHasCall(n.Left)
		n.List.Set(psess.reorder1(ll))

	case OAS, OASOP:
		init.AppendNodes(&n.Ninit)

		mapAppend := n.Left.Op == OINDEXMAP && n.Right.Op == OAPPEND
		if mapAppend && !psess.samesafeexpr(n.Left, n.Right.List.First()) {
			psess.
				Fatalf("not same expressions: %v != %v", n.Left, n.Right.List.First())
		}

		n.Left = psess.walkexpr(n.Left, init)
		n.Left = psess.safeexpr(n.Left, init)

		if mapAppend {
			n.Right.List.SetFirst(n.Left)
		}

		if n.Op == OASOP {

			n.Right = psess.nod(n.SubOp(psess), n.Left, n.Right)
			n.Right = psess.typecheck(n.Right, Erv)

			n.Op = OAS
			n.ResetAux()
		}

		if psess.oaslit(n, init) {
			break
		}

		if n.Right == nil {

			break
		}

		if !psess.instrumenting && psess.isZero(n.Right) {
			break
		}

		switch n.Right.Op {
		default:
			n.Right = psess.walkexpr(n.Right, init)

		case ORECV:

			n.Right.Left = psess.walkexpr(n.Right.Left, init)

			n1 := psess.nod(OADDR, n.Left, nil)
			r := n.Right.Left
			n = psess.mkcall1(psess.chanfn("chanrecv1", 2, r.Type), nil, init, r, n1)
			n = psess.walkexpr(n, init)
			break opswitch

		case OAPPEND:

			r := n.Right
			if r.Type.Elem(psess.types).NotInHeap() {
				psess.
					yyerror("%v is go:notinheap; heap allocation disallowed", r.Type.Elem(psess.types))
			}
			switch {
			case psess.isAppendOfMake(r):

				r = psess.extendslice(r, init)
			case r.Isddd():
				r = psess.appendslice(r, init)
			default:
				r = psess.walkappend(r, init, n)
			}
			n.Right = r
			if r.Op == OAPPEND {

				r.Left = psess.typename(r.Type.Elem(psess.types))
				break opswitch
			}

		}

		if n.Left != nil && n.Right != nil {
			n = psess.convas(n, init)
		}

	case OAS2:
		init.AppendNodes(&n.Ninit)
		psess.
			walkexprlistsafe(n.List.Slice(), init)
		psess.
			walkexprlistsafe(n.Rlist.Slice(), init)
		ll := psess.ascompatee(OAS, n.List.Slice(), n.Rlist.Slice(), init)
		ll = psess.reorder3(ll)
		n = psess.liststmt(ll)

	case OAS2FUNC:
		init.AppendNodes(&n.Ninit)

		r := n.Rlist.First()
		psess.
			walkexprlistsafe(n.List.Slice(), init)
		r = psess.walkexpr(r, init)

		if psess.isIntrinsicCall(r) {
			n.Rlist.Set1(r)
			break
		}
		init.Append(r)

		ll := psess.ascompatet(n.List, r.Type)
		n = psess.liststmt(ll)

	case OAS2RECV:
		init.AppendNodes(&n.Ninit)

		r := n.Rlist.First()
		psess.
			walkexprlistsafe(n.List.Slice(), init)
		r.Left = psess.walkexpr(r.Left, init)
		var n1 *Node
		if n.List.First().isBlank() {
			n1 = psess.nodnil()
		} else {
			n1 = psess.nod(OADDR, n.List.First(), nil)
		}
		fn := psess.chanfn("chanrecv2", 2, r.Left.Type)
		ok := n.List.Second()
		call := psess.mkcall1(fn, ok.Type, init, r.Left, n1)
		n = psess.nod(OAS, ok, call)
		n = psess.typecheck(n, Etop)

	case OAS2MAPR:
		init.AppendNodes(&n.Ninit)

		r := n.Rlist.First()
		psess.
			walkexprlistsafe(n.List.Slice(), init)
		r.Left = psess.walkexpr(r.Left, init)
		r.Right = psess.walkexpr(r.Right, init)
		t := r.Left.Type

		fast := psess.mapfast(t)
		var key *Node
		if fast != mapslow {

			key = r.Right
		} else {

			key = psess.nod(OADDR, r.Right, nil)
		}

		a := n.List.First()

		if w := t.Elem(psess.types).Width; w <= 1024 {
			fn := psess.mapfn(psess.mapaccess2[fast], t)
			r = psess.mkcall1(fn, fn.Type.Results(psess.types), init, psess.typename(t), r.Left, key)
		} else {
			fn := psess.mapfn("mapaccess2_fat", t)
			z := psess.zeroaddr(w)
			r = psess.mkcall1(fn, fn.Type.Results(psess.types), init, psess.typename(t), r.Left, key, z)
		}

		if ok := n.List.Second(); !ok.isBlank() && ok.Type.IsBoolean() {
			r.Type.Field(psess.types, 1).Type = ok.Type
		}
		n.Rlist.Set1(r)
		n.Op = OAS2FUNC

		if !a.isBlank() {
			var_ := psess.temp(psess.types.NewPtr(t.Elem(psess.types)))
			var_.SetTypecheck(1)
			var_.SetNonNil(true)
			n.List.SetFirst(var_)
			n = psess.walkexpr(n, init)
			init.Append(n)
			n = psess.nod(OAS, a, psess.nod(OIND, var_, nil))
		}

		n = psess.typecheck(n, Etop)
		n = psess.walkexpr(n, init)

	case ODELETE:
		init.AppendNodes(&n.Ninit)
		map_ := n.List.First()
		key := n.List.Second()
		map_ = psess.walkexpr(map_, init)
		key = psess.walkexpr(key, init)

		t := map_.Type
		fast := psess.mapfast(t)
		if fast == mapslow {

			key = psess.nod(OADDR, key, nil)
		}
		n = psess.mkcall1(psess.mapfndel(psess.mapdelete[fast], t), nil, init, psess.typename(t), map_, key)

	case OAS2DOTTYPE:
		psess.
			walkexprlistsafe(n.List.Slice(), init)
		n.Rlist.SetFirst(psess.walkexpr(n.Rlist.First(), init))

	case OCONVIFACE:
		n.Left = psess.walkexpr(n.Left, init)

		if psess.isdirectiface(n.Left.Type) {
			var t *Node
			if n.Type.IsEmptyInterface(psess.types) {
				t = psess.typename(n.Left.Type)
			} else {
				t = psess.itabname(n.Left.Type, n.Type)
			}
			l := psess.nod(OEFACE, t, n.Left)
			l.Type = n.Type
			l.SetTypecheck(n.Typecheck())
			n = l
			break
		}

		if psess.staticbytes == nil {
			psess.
				staticbytes = psess.newname(psess.Runtimepkg.Lookup(psess.types, "staticbytes"))
			psess.
				staticbytes.SetClass(PEXTERN)
			psess.
				staticbytes.Type = psess.types.NewArray(psess.types.Types[TUINT8], 256)
			psess.
				zerobase = psess.newname(psess.Runtimepkg.Lookup(psess.types, "zerobase"))
			psess.
				zerobase.SetClass(PEXTERN)
			psess.
				zerobase.Type = psess.types.Types[TUINTPTR]
		}

		// Optimize convT2{E,I} for many cases in which T is not pointer-shaped,
		// by using an existing addressable value identical to n.Left
		// or creating one on the stack.
		var value *Node
		switch {
		case n.Left.Type.Size(psess.types) == 0:
			psess.
				cheapexpr(n.Left, init)
			value = psess.zerobase
		case n.Left.Type.IsBoolean() || (n.Left.Type.Size(psess.types) == 1 && n.Left.Type.IsInteger()):

			n.Left = psess.cheapexpr(n.Left, init)
			value = psess.nod(OINDEX, psess.staticbytes, psess.byteindex(n.Left))
			value.SetBounded(true)
		case n.Left.Class() == PEXTERN && n.Left.Name != nil && n.Left.Name.Readonly():

			value = n.Left
		case !n.Left.Type.IsInterface() && n.Esc == EscNone && n.Left.Type.Width <= 1024:

			value = psess.temp(n.Left.Type)
			init.Append(psess.typecheck(psess.nod(OAS, value, n.Left), Etop))
		}

		if value != nil {
			// Value is identical to n.Left.
			// Construct the interface directly: {type/itab, &value}.
			var t *Node
			if n.Type.IsEmptyInterface(psess.types) {
				t = psess.typename(n.Left.Type)
			} else {
				t = psess.itabname(n.Left.Type, n.Type)
			}
			l := psess.nod(OEFACE, t, psess.typecheck(psess.nod(OADDR, value, nil), Erv))
			l.Type = n.Type
			l.SetTypecheck(n.Typecheck())
			n = l
			break
		}

		if n.Type.IsEmptyInterface(psess.types) && n.Left.Type.IsInterface() && !n.Left.Type.IsEmptyInterface(psess.types) {

			c := psess.temp(n.Left.Type)
			init.Append(psess.nod(OAS, c, n.Left))

			tmp := psess.temp(psess.types.NewPtr(psess.types.Types[TUINT8]))
			init.Append(psess.nod(OAS, tmp, psess.typecheck(psess.nod(OITAB, c, nil), Erv)))

			nif := psess.nod(OIF, psess.typecheck(psess.nod(ONE, tmp, psess.nodnil()), Erv), nil)
			nif.Nbody.Set1(psess.nod(OAS, tmp, psess.itabType(tmp)))
			init.Append(nif)

			e := psess.nod(OEFACE, tmp, psess.ifaceData(c, psess.types.NewPtr(psess.types.Types[TUINT8])))
			e.Type = n.Type
			e.SetTypecheck(1)
			n = e
			break
		}

		var ll []*Node
		if n.Type.IsEmptyInterface(psess.types) {
			if !n.Left.Type.IsInterface() {
				ll = append(ll, psess.typename(n.Left.Type))
			}
		} else {
			if n.Left.Type.IsInterface() {
				ll = append(ll, psess.typename(n.Type))
			} else {
				ll = append(ll, psess.itabname(n.Left.Type, n.Type))
			}
		}

		fnname, needsaddr := psess.convFuncName(n.Left.Type, n.Type)
		v := n.Left
		if needsaddr {

			if !islvalue(v) {
				v = psess.copyexpr(v, v.Type, init)
			}
			v = psess.nod(OADDR, v, nil)
		}
		ll = append(ll, v)
		psess.
			dowidth(n.Left.Type)
		fn := psess.syslook(fnname)
		fn = psess.substArgTypes(fn, n.Left.Type, n.Type)
		psess.
			dowidth(fn.Type)
		n = psess.nod(OCALL, fn, nil)
		n.List.Set(ll)
		n = psess.typecheck(n, Erv)
		n = psess.walkexpr(n, init)

	case OCONV, OCONVNOP:
		if psess.thearch.SoftFloat {

			n.Left = psess.walkexpr(n.Left, init)
			break
		}
		switch psess.thearch.LinkArch.Family {
		case sys.ARM, sys.MIPS:
			if n.Left.Type.IsFloat() {
				switch n.Type.Etype {
				case TINT64:
					n = psess.mkcall("float64toint64", n.Type, init, psess.conv(n.Left, psess.types.Types[TFLOAT64]))
					break opswitch
				case TUINT64:
					n = psess.mkcall("float64touint64", n.Type, init, psess.conv(n.Left, psess.types.Types[TFLOAT64]))
					break opswitch
				}
			}

			if n.Type.IsFloat() {
				switch n.Left.Type.Etype {
				case TINT64:
					n = psess.conv(psess.mkcall("int64tofloat64", psess.types.Types[TFLOAT64], init, psess.conv(n.Left, psess.types.Types[TINT64])), n.Type)
					break opswitch
				case TUINT64:
					n = psess.conv(psess.mkcall("uint64tofloat64", psess.types.Types[TFLOAT64], init, psess.conv(n.Left, psess.types.Types[TUINT64])), n.Type)
					break opswitch
				}
			}

		case sys.I386:
			if n.Left.Type.IsFloat() {
				switch n.Type.Etype {
				case TINT64:
					n = psess.mkcall("float64toint64", n.Type, init, psess.conv(n.Left, psess.types.Types[TFLOAT64]))
					break opswitch
				case TUINT64:
					n = psess.mkcall("float64touint64", n.Type, init, psess.conv(n.Left, psess.types.Types[TFLOAT64]))
					break opswitch
				case TUINT32, TUINT, TUINTPTR:
					n = psess.mkcall("float64touint32", n.Type, init, psess.conv(n.Left, psess.types.Types[TFLOAT64]))
					break opswitch
				}
			}
			if n.Type.IsFloat() {
				switch n.Left.Type.Etype {
				case TINT64:
					n = psess.conv(psess.mkcall("int64tofloat64", psess.types.Types[TFLOAT64], init, psess.conv(n.Left, psess.types.Types[TINT64])), n.Type)
					break opswitch
				case TUINT64:
					n = psess.conv(psess.mkcall("uint64tofloat64", psess.types.Types[TFLOAT64], init, psess.conv(n.Left, psess.types.Types[TUINT64])), n.Type)
					break opswitch
				case TUINT32, TUINT, TUINTPTR:
					n = psess.conv(psess.mkcall("uint32tofloat64", psess.types.Types[TFLOAT64], init, psess.conv(n.Left, psess.types.Types[TUINT32])), n.Type)
					break opswitch
				}
			}
		}
		n.Left = psess.walkexpr(n.Left, init)

	case OANDNOT:
		n.Left = psess.walkexpr(n.Left, init)
		n.Op = OAND
		n.Right = psess.nod(OCOM, n.Right, nil)
		n.Right = psess.typecheck(n.Right, Erv)
		n.Right = psess.walkexpr(n.Right, init)

	case ODIV, OMOD:
		n.Left = psess.walkexpr(n.Left, init)
		n.Right = psess.walkexpr(n.Right, init)

		et := n.Left.Type.Etype

		if psess.isComplex[et] && n.Op == ODIV {
			t := n.Type
			n = psess.mkcall("complex128div", psess.types.Types[TCOMPLEX128], init, psess.conv(n.Left, psess.types.Types[TCOMPLEX128]), psess.conv(n.Right, psess.types.Types[TCOMPLEX128]))
			n = psess.conv(n, t)
			break
		}

		if psess.isFloat[et] {
			break
		}

		if psess.Widthreg < 8 && (et == TINT64 || et == TUINT64) {
			if n.Right.Op == OLITERAL {

				switch et {
				case TINT64:
					c := n.Right.Int64(psess)
					if c < 0 {
						c = -c
					}
					if c != 0 && c&(c-1) == 0 {
						break opswitch
					}
				case TUINT64:
					c := uint64(n.Right.Int64(psess))
					if c != 0 && c&(c-1) == 0 {
						break opswitch
					}
				}
			}
			var fn string
			if et == TINT64 {
				fn = "int64"
			} else {
				fn = "uint64"
			}
			if n.Op == ODIV {
				fn += "div"
			} else {
				fn += "mod"
			}
			n = psess.mkcall(fn, n.Type, init, psess.conv(n.Left, psess.types.Types[et]), psess.conv(n.Right, psess.types.Types[et]))
		}

	case OINDEX:
		n.Left = psess.walkexpr(n.Left, init)

		r := n.Right

		n.Right = psess.walkexpr(n.Right, init)

		if n.Bounded() {
			break
		}
		t := n.Left.Type
		if t != nil && t.IsPtr() {
			t = t.Elem(psess.types)
		}
		if t.IsArray() {
			n.SetBounded(psess.bounded(r, t.NumElem(psess.types)))
			if psess.Debug['m'] != 0 && n.Bounded() && !psess.Isconst(n.Right, CTINT) {
				psess.
					Warn("index bounds check elided")
			}
			if psess.smallintconst(n.Right) && !n.Bounded() {
				psess.
					yyerror("index out of bounds")
			}
		} else if psess.Isconst(n.Left, CTSTR) {
			n.SetBounded(psess.bounded(r, int64(len(n.Left.Val().U.(string)))))
			if psess.Debug['m'] != 0 && n.Bounded() && !psess.Isconst(n.Right, CTINT) {
				psess.
					Warn("index bounds check elided")
			}
			if psess.smallintconst(n.Right) && !n.Bounded() {
				psess.
					yyerror("index out of bounds")
			}
		}

		if psess.Isconst(n.Right, CTINT) {
			if n.Right.Val().U.(*Mpint).CmpInt64(0) < 0 || n.Right.Val().U.(*Mpint).Cmp(psess.maxintval[TINT]) > 0 {
				psess.
					yyerror("index out of bounds")
			}
		}

	case OINDEXMAP:

		n.Left = psess.walkexpr(n.Left, init)
		n.Right = psess.walkexpr(n.Right, init)
		map_ := n.Left
		key := n.Right
		t := map_.Type
		if n.IndexMapLValue(psess) {

			fast := psess.mapfast(t)
			if fast == mapslow {

				key = psess.nod(OADDR, key, nil)
			}
			n = psess.mkcall1(psess.mapfn(psess.mapassign[fast], t), nil, init, psess.typename(t), map_, key)
		} else {

			fast := psess.mapfast(t)
			if fast == mapslow {

				key = psess.nod(OADDR, key, nil)
			}

			if w := t.Elem(psess.types).Width; w <= 1024 {
				n = psess.mkcall1(psess.mapfn(psess.mapaccess1[fast], t), psess.types.NewPtr(t.Elem(psess.types)), init, psess.typename(t), map_, key)
			} else {
				z := psess.zeroaddr(w)
				n = psess.mkcall1(psess.mapfn("mapaccess1_fat", t), psess.types.NewPtr(t.Elem(psess.types)), init, psess.typename(t), map_, key, z)
			}
		}
		n.Type = psess.types.NewPtr(t.Elem(psess.types))
		n.SetNonNil(true)
		n = psess.nod(OIND, n, nil)
		n.Type = t.Elem(psess.types)
		n.SetTypecheck(1)

	case ORECV:
		psess.
			Fatalf("walkexpr ORECV")

	case OSLICE, OSLICEARR, OSLICESTR, OSLICE3, OSLICE3ARR:
		n.Left = psess.walkexpr(n.Left, init)
		low, high, max := n.SliceBounds(psess)
		low = psess.walkexpr(low, init)
		if low != nil && psess.isZero(low) {

			low = nil
		}
		high = psess.walkexpr(high, init)
		max = psess.walkexpr(max, init)
		n.SetSliceBounds(psess, low, high, max)
		if n.Op.IsSlice3(psess) {
			if max != nil && max.Op == OCAP && psess.samesafeexpr(n.Left, max.Left) {

				if n.Op == OSLICE3 {
					n.Op = OSLICE
				} else {
					n.Op = OSLICEARR
				}
				n = psess.reduceSlice(n)
			}
		} else {
			n = psess.reduceSlice(n)
		}

	case ONEW:
		if n.Esc == EscNone {
			if n.Type.Elem(psess.types).Width >= 1<<16 {
				psess.
					Fatalf("large ONEW with EscNone: %v", n)
			}
			r := psess.temp(n.Type.Elem(psess.types))
			r = psess.nod(OAS, r, nil)
			r = psess.typecheck(r, Etop)
			init.Append(r)
			r = psess.nod(OADDR, r.Left, nil)
			r = psess.typecheck(r, Erv)
			n = r
		} else {
			n = psess.callnew(n.Type.Elem(psess.types))
		}

	case OCMPSTR:

		if (n.SubOp(psess) == OEQ || n.SubOp(psess) == ONE) && psess.Isconst(n.Right, CTSTR) && n.Left.Op == OADDSTR && n.Left.List.Len() == 2 && psess.Isconst(n.Left.List.Second(), CTSTR) && strlit(n.Right) == strlit(n.Left.List.Second()) {
			r := psess.nod(n.SubOp(psess), psess.nod(OLEN, n.Left.List.First(), nil), psess.nodintconst(0))
			n = psess.finishcompare(n, r, init)
			break
		}

		// Rewrite comparisons to short constant strings as length+byte-wise comparisons.
		var cs, ncs *Node // const string, non-const string
		switch {
		case psess.Isconst(n.Left, CTSTR) && psess.Isconst(n.Right, CTSTR):

		case psess.Isconst(n.Left, CTSTR):
			cs = n.Left
			ncs = n.Right
		case psess.Isconst(n.Right, CTSTR):
			cs = n.Right
			ncs = n.Left
		}
		if cs != nil {
			cmp := n.SubOp(psess)

			if psess.Isconst(n.Left, CTSTR) {
				cmp = psess.brrev(cmp)
			}

			maxRewriteLen := 6

			canCombineLoads := psess.canMergeLoads()
			combine64bit := false
			if canCombineLoads {

				maxRewriteLen = 2 * psess.thearch.LinkArch.RegSize
				combine64bit = psess.thearch.LinkArch.RegSize >= 8
			}

			var and Op
			switch cmp {
			case OEQ:
				and = OANDAND
			case ONE:
				and = OOROR
			default:

				maxRewriteLen = 0
			}
			if s := cs.Val().U.(string); len(s) <= maxRewriteLen {
				if len(s) > 0 {
					ncs = psess.safeexpr(ncs, init)
				}
				r := psess.nod(cmp, psess.nod(OLEN, ncs, nil), psess.nodintconst(int64(len(s))))
				remains := len(s)
				for i := 0; remains > 0; {
					if remains == 1 || !canCombineLoads {
						cb := psess.nodintconst(int64(s[i]))
						ncb := psess.nod(OINDEX, ncs, psess.nodintconst(int64(i)))
						r = psess.nod(and, r, psess.nod(cmp, ncb, cb))
						remains--
						i++
						continue
					}
					var step int
					var convType *types.Type
					switch {
					case remains >= 8 && combine64bit:
						convType = psess.types.Types[TINT64]
						step = 8
					case remains >= 4:
						convType = psess.types.Types[TUINT32]
						step = 4
					case remains >= 2:
						convType = psess.types.Types[TUINT16]
						step = 2
					}
					ncsubstr := psess.nod(OINDEX, ncs, psess.nodintconst(int64(i)))
					ncsubstr = psess.conv(ncsubstr, convType)
					csubstr := int64(s[i])

					for offset := 1; offset < step; offset++ {
						b := psess.nod(OINDEX, ncs, psess.nodintconst(int64(i+offset)))
						b = psess.conv(b, convType)
						b = psess.nod(OLSH, b, psess.nodintconst(int64(8*offset)))
						ncsubstr = psess.nod(OOR, ncsubstr, b)
						csubstr = csubstr | int64(s[i+offset])<<uint8(8*offset)
					}
					csubstrPart := psess.nodintconst(csubstr)

					r = psess.nod(and, r, psess.nod(cmp, csubstrPart, ncsubstr))
					remains -= step
					i += step
				}
				n = psess.finishcompare(n, r, init)
				break
			}
		}

		var r *Node
		if n.SubOp(psess) == OEQ || n.SubOp(psess) == ONE {

			n.Left = psess.cheapexpr(n.Left, init)
			n.Right = psess.cheapexpr(n.Right, init)

			lstr := psess.conv(n.Left, psess.types.Types[TSTRING])
			rstr := psess.conv(n.Right, psess.types.Types[TSTRING])
			lptr := psess.nod(OSPTR, lstr, nil)
			rptr := psess.nod(OSPTR, rstr, nil)
			llen := psess.conv(psess.nod(OLEN, lstr, nil), psess.types.Types[TUINTPTR])
			rlen := psess.conv(psess.nod(OLEN, rstr, nil), psess.types.Types[TUINTPTR])

			fn := psess.syslook("memequal")
			fn = psess.substArgTypes(fn, psess.types.Types[TUINT8], psess.types.Types[TUINT8])
			r = psess.mkcall1(fn, psess.types.Types[TBOOL], init, lptr, rptr, llen)

			if n.SubOp(psess) == OEQ {

				r = psess.nod(OANDAND, psess.nod(OEQ, llen, rlen), r)
			} else {

				r = psess.nod(ONOT, r, nil)
				r = psess.nod(OOROR, psess.nod(ONE, llen, rlen), r)
			}
		} else {

			r = psess.mkcall("cmpstring", psess.types.Types[TINT], init, psess.conv(n.Left, psess.types.Types[TSTRING]), psess.conv(n.Right, psess.types.Types[TSTRING]))
			r = psess.nod(n.SubOp(psess), r, psess.nodintconst(0))
		}

		n = psess.finishcompare(n, r, init)

	case OADDSTR:
		n = psess.addstr(n, init)

	case OAPPEND:
		psess.
			Fatalf("append outside assignment")

	case OCOPY:
		n = psess.copyany(n, init, psess.instrumenting && !psess.compiling_runtime)

	case OCLOSE:
		fn := psess.syslook("closechan")

		fn = psess.substArgTypes(fn, n.Left.Type)
		n = psess.mkcall1(fn, nil, init, n.Left)

	case OMAKECHAN:

		size := n.Left
		fnname := "makechan64"
		argtype := psess.types.Types[TINT64]

		if size.Type.IsKind(TIDEAL) || psess.maxintval[size.Type.Etype].Cmp(psess.maxintval[TUINT]) <= 0 {
			fnname = "makechan"
			argtype = psess.types.Types[TINT]
		}

		n = psess.mkcall1(psess.chanfn(fnname, 1, n.Type), n.Type, init, psess.typename(n.Type), psess.conv(size, argtype))

	case OMAKEMAP:
		t := n.Type
		hmapType := psess.hmap(t)
		hint := n.Left

		// var h *hmap
		var h *Node
		if n.Esc == EscNone {

			hv := psess.temp(hmapType)
			zero := psess.nod(OAS, hv, nil)
			zero = psess.typecheck(zero, Etop)
			init.Append(zero)

			h = psess.nod(OADDR, hv, nil)

			if !psess.Isconst(hint, CTINT) ||
				!(hint.Val().U.(*Mpint).CmpInt64(BUCKETSIZE) > 0) {

				bv := psess.temp(psess.bmap(t))

				zero = psess.nod(OAS, bv, nil)
				zero = psess.typecheck(zero, Etop)
				init.Append(zero)

				b := psess.nod(OADDR, bv, nil)

				bsym := hmapType.Field(psess.types, 5).Sym
				na := psess.nod(OAS, psess.nodSym(ODOT, h, bsym), b)
				na = psess.typecheck(na, Etop)
				init.Append(na)
			}
		}

		if psess.Isconst(hint, CTINT) && hint.Val().U.(*Mpint).CmpInt64(BUCKETSIZE) <= 0 {

			if n.Esc == EscNone {

				rand := psess.mkcall("fastrand", psess.types.Types[TUINT32], init)
				hashsym := hmapType.Field(psess.types, 4).Sym
				a := psess.nod(OAS, psess.nodSym(ODOT, h, hashsym), rand)
				a = psess.typecheck(a, Etop)
				a = psess.walkexpr(a, init)
				init.Append(a)
				n = psess.nod(OCONVNOP, h, nil)
				n.Type = t
				n = psess.typecheck(n, Erv)
			} else {

				fn := psess.syslook("makemap_small")
				fn = psess.substArgTypes(fn, t.Key(psess.types), t.Elem(psess.types))
				n = psess.mkcall1(fn, n.Type, init)
			}
		} else {
			if n.Esc != EscNone {
				h = psess.nodnil()
			}

			fnname := "makemap64"
			argtype := psess.types.Types[TINT64]

			if hint.Type.IsKind(TIDEAL) || psess.maxintval[hint.Type.Etype].Cmp(psess.maxintval[TUINT]) <= 0 {
				fnname = "makemap"
				argtype = psess.types.Types[TINT]
			}

			fn := psess.syslook(fnname)
			fn = psess.substArgTypes(fn, hmapType, t.Key(psess.types), t.Elem(psess.types))
			n = psess.mkcall1(fn, n.Type, init, psess.typename(n.Type), psess.conv(hint, argtype), h)
		}

	case OMAKESLICE:
		l := n.Left
		r := n.Right
		if r == nil {
			r = psess.safeexpr(l, init)
			l = r
		}
		t := n.Type
		if n.Esc == EscNone {
			if !psess.isSmallMakeSlice(n) {
				psess.
					Fatalf("non-small OMAKESLICE with EscNone: %v", n)
			}

			t = psess.types.NewArray(t.Elem(psess.types), psess.nonnegintconst(r))
			var_ := psess.temp(t)
			a := psess.nod(OAS, var_, nil)
			a = psess.typecheck(a, Etop)
			init.Append(a)
			r := psess.nod(OSLICE, var_, nil)
			r.SetSliceBounds(psess, nil, l, nil)
			r = psess.conv(r, n.Type)
			r = psess.typecheck(r, Erv)
			r = psess.walkexpr(r, init)
			n = r
		} else {

			if t.Elem(psess.types).NotInHeap() {
				psess.
					yyerror("%v is go:notinheap; heap allocation disallowed", t.Elem(psess.types))
			}

			len, cap := l, r

			fnname := "makeslice64"
			argtype := psess.types.Types[TINT64]

			if (len.Type.IsKind(TIDEAL) || psess.maxintval[len.Type.Etype].Cmp(psess.maxintval[TUINT]) <= 0) &&
				(cap.Type.IsKind(TIDEAL) || psess.maxintval[cap.Type.Etype].Cmp(psess.maxintval[TUINT]) <= 0) {
				fnname = "makeslice"
				argtype = psess.types.Types[TINT]
			}

			fn := psess.syslook(fnname)
			fn = psess.substArgTypes(fn, t.Elem(psess.types))
			n = psess.mkcall1(fn, t, init, psess.typename(t.Elem(psess.types)), psess.conv(len, argtype), psess.conv(cap, argtype))
		}

	case ORUNESTR:
		a := psess.nodnil()
		if n.Esc == EscNone {
			t := psess.types.NewArray(psess.types.Types[TUINT8], 4)
			var_ := psess.temp(t)
			a = psess.nod(OADDR, var_, nil)
		}

		n = psess.mkcall("intstring", n.Type, init, a, psess.conv(n.Left, psess.types.Types[TINT64]))

	case OARRAYBYTESTR:
		a := psess.nodnil()
		if n.Esc == EscNone {

			t := psess.types.NewArray(psess.types.Types[TUINT8], tmpstringbufsize)

			a = psess.nod(OADDR, psess.temp(t), nil)
		}

		n = psess.mkcall("slicebytetostring", n.Type, init, a, n.Left)

	case OARRAYBYTESTRTMP:
		n.Left = psess.walkexpr(n.Left, init)

		if !psess.instrumenting {

			break
		}

		n = psess.mkcall("slicebytetostringtmp", n.Type, init, n.Left)

	case OARRAYRUNESTR:
		a := psess.nodnil()

		if n.Esc == EscNone {

			t := psess.types.NewArray(psess.types.Types[TUINT8], tmpstringbufsize)

			a = psess.nod(OADDR, psess.temp(t), nil)
		}

		n = psess.mkcall("slicerunetostring", n.Type, init, a, n.Left)

	case OSTRARRAYBYTE:
		a := psess.nodnil()

		if n.Esc == EscNone {

			t := psess.types.NewArray(psess.types.Types[TUINT8], tmpstringbufsize)

			a = psess.nod(OADDR, psess.temp(t), nil)
		}

		n = psess.mkcall("stringtoslicebyte", n.Type, init, a, psess.conv(n.Left, psess.types.Types[TSTRING]))

	case OSTRARRAYBYTETMP:

		n.Left = psess.walkexpr(n.Left, init)

	case OSTRARRAYRUNE:
		a := psess.nodnil()

		if n.Esc == EscNone {

			t := psess.types.NewArray(psess.types.Types[TINT32], tmpstringbufsize)

			a = psess.nod(OADDR, psess.temp(t), nil)
		}

		n = psess.mkcall("stringtoslicerune", n.Type, init, a, psess.conv(n.Left, psess.types.Types[TSTRING]))

	case OCMPIFACE:
		if !psess.eqtype(n.Left.Type, n.Right.Type) {
			psess.
				Fatalf("ifaceeq %v %v %v", n.Op, n.Left.Type, n.Right.Type)
		}
		var fn *Node
		if n.Left.Type.IsEmptyInterface(psess.types) {
			fn = psess.syslook("efaceeq")
		} else {
			fn = psess.syslook("ifaceeq")
		}

		n.Right = psess.cheapexpr(n.Right, init)
		n.Left = psess.cheapexpr(n.Left, init)
		lt := psess.nod(OITAB, n.Left, nil)
		rt := psess.nod(OITAB, n.Right, nil)
		ld := psess.nod(OIDATA, n.Left, nil)
		rd := psess.nod(OIDATA, n.Right, nil)
		ld.Type = psess.types.Types[TUNSAFEPTR]
		rd.Type = psess.types.Types[TUNSAFEPTR]
		ld.SetTypecheck(1)
		rd.SetTypecheck(1)
		call := psess.mkcall1(fn, n.Type, init, lt, ld, rd)

		// Check itable/type before full compare.
		// Note: short-circuited because order matters.
		var cmp *Node
		if n.SubOp(psess) == OEQ {
			cmp = psess.nod(OANDAND, psess.nod(OEQ, lt, rt), call)
		} else {
			cmp = psess.nod(OOROR, psess.nod(ONE, lt, rt), psess.nod(ONOT, call, nil))
		}
		n = psess.finishcompare(n, cmp, init)

	case OARRAYLIT, OSLICELIT, OMAPLIT, OSTRUCTLIT, OPTRLIT:
		if psess.isStaticCompositeLiteral(n) && !psess.canSSAType(n.Type) {

			vstat := psess.staticname(n.Type)
			vstat.Name.SetReadonly(true)
			psess.
				fixedlit(inInitFunction, initKindStatic, n, vstat, init)
			n = vstat
			n = psess.typecheck(n, Erv)
			break
		}
		var_ := psess.temp(n.Type)
		psess.
			anylit(n, var_, init)
		n = var_

	case OSEND:
		n1 := n.Right
		n1 = psess.assignconv(n1, n.Left.Type.Elem(psess.types), "chan send")
		n1 = psess.walkexpr(n1, init)
		n1 = psess.nod(OADDR, n1, nil)
		n = psess.mkcall1(psess.chanfn("chansend1", 2, n.Left.Type), nil, init, n.Left, n1)

	case OCLOSURE:
		n = psess.walkclosure(n, init)

	case OCALLPART:
		n = psess.walkpartialcall(n, init)
	}

	t := n.Type
	psess.
		evconst(n)
	if n.Type != t {
		psess.
			Fatalf("evconst changed Type: %v had type %v, now %v", n, t, n.Type)
	}
	if n.Op == OLITERAL {
		n = psess.typecheck(n, Erv)

		if s, ok := n.Val().U.(string); ok {
			_ = psess.stringsym(n.Pos, s)
		}
	}
	psess.
		updateHasCall(n)

	if psess.Debug['w'] != 0 && n != nil {
		Dump("after walk expr", n)
	}
	psess.
		lineno = lno
	return n
}

// TODO(josharian): combine this with its caller and simplify
func (psess *PackageSession) reduceSlice(n *Node) *Node {
	low, high, max := n.SliceBounds(psess)
	if high != nil && high.Op == OLEN && psess.samesafeexpr(n.Left, high.Left) {

		high = nil
	}
	n.SetSliceBounds(psess, low, high, max)
	if (n.Op == OSLICE || n.Op == OSLICESTR) && low == nil && high == nil {

		if psess.Debug_slice > 0 {
			psess.
				Warn("slice: omit slice operation")
		}
		return n.Left
	}
	return n
}

func (psess *PackageSession) ascompatee1(l *Node, r *Node, init *Nodes) *Node {

	n := psess.nod(OAS, l, r)

	if l.Op == OINDEXMAP {
		return n
	}

	return psess.convas(n, init)
}

func (psess *PackageSession) ascompatee(op Op, nl, nr []*Node, init *Nodes) []*Node {

	for i := range nl {
		nl[i] = psess.safeexpr(nl[i], init)
	}
	for i1 := range nr {
		nr[i1] = psess.safeexpr(nr[i1], init)
	}

	var nn []*Node
	i := 0
	for ; i < len(nl); i++ {
		if i >= len(nr) {
			break
		}

		if op == ORETURN && psess.samesafeexpr(nl[i], nr[i]) {
			continue
		}
		nn = append(nn, psess.ascompatee1(nl[i], nr[i], init))
	}

	if i < len(nl) || i < len(nr) {
		var nln, nrn Nodes
		nln.Set(nl)
		nrn.Set(nr)
		psess.
			Fatalf("error in shape across %+v %v %+v / %d %d [%s]", nln, op, nrn, len(nl), len(nr), psess.Curfn.funcname())
	}
	return nn
}

// fncall reports whether assigning an rvalue of type rt to an lvalue l might involve a function call.
func (psess *PackageSession) fncall(l *Node, rt *types.Type) bool {
	if l.HasCall() || l.Op == OINDEXMAP {
		return true
	}
	if psess.eqtype(l.Type, rt) {
		return false
	}

	return true
}

// check assign type list to
// an expression list. called in
//	expr-list = func()
func (psess *PackageSession) ascompatet(nl Nodes, nr *types.Type) []*Node {
	if nl.Len() != nr.NumFields(psess.types) {
		psess.
			Fatalf("ascompatet: assignment count mismatch: %d = %d", nl.Len(), nr.NumFields(psess.types))
	}

	var nn, mm Nodes
	for i, l := range nl.Slice() {
		if l.isBlank() {
			continue
		}
		r := nr.Field(psess.types, i)

		if psess.fncall(l, r.Type) {
			tmp := psess.temp(r.Type)
			tmp = psess.typecheck(tmp, Erv)
			a := psess.nod(OAS, l, tmp)
			a = psess.convas(a, &mm)
			mm.Append(a)
			l = tmp
		}

		a := psess.nod(OAS, l, psess.nodarg(r, 0))
		a = psess.convas(a, &nn)
		psess.
			updateHasCall(a)
		if a.HasCall() {
			Dump("ascompatet ucount", a)
			psess.
				Fatalf("ascompatet: too many function calls evaluating parameters")
		}

		nn.Append(a)
	}
	return append(nn.Slice(), mm.Slice()...)
}

// nodarg returns a Node for the function argument denoted by t,
// which is either the entire function argument or result struct (t is a  struct *types.Type)
// or a specific argument (t is a *types.Field within a struct *types.Type).
//
// If fp is 0, the node is for use by a caller invoking the given
// function, preparing the arguments before the call
// or retrieving the results after the call.
// In this case, the node will correspond to an outgoing argument
// slot like 8(SP).
//
// If fp is 1, the node is for use by the function itself
// (the callee), to retrieve its arguments or write its results.
// In this case the node will be an ONAME with an appropriate
// type and offset.
func (psess *PackageSession) nodarg(t interface{}, fp int) *Node {
	var n *Node

	switch t := t.(type) {
	default:
		psess.
			Fatalf("bad nodarg %T(%v)", t, t)

	case *types.Type:

		if !t.IsFuncArgStruct() {
			psess.
				Fatalf("nodarg: bad type %v", t)
		}

		n = psess.newname(psess.lookup(".args"))
		n.Type = t
		first := t.Field(psess.types, 0)
		if first == nil {
			psess.
				Fatalf("nodarg: bad struct")
		}
		if first.Offset == BADWIDTH {
			psess.
				Fatalf("nodarg: offset not computed for %v", t)
		}
		n.Xoffset = first.Offset

	case *types.Field:
		if fp == 1 {

			expect := asNode(t.Nname)
			if expect.isParamHeapCopy() {
				expect = expect.Name.Param.Stackcopy
			}

			for _, n := range psess.Curfn.Func.Dcl {
				if (n.Class() == PPARAM || n.Class() == PPARAMOUT) && !t.Sym.IsBlank() && n.Sym == t.Sym {
					if n != expect {
						psess.
							Fatalf("nodarg: unexpected node: %v (%p %v) vs %v (%p %v)", n, n, n.Op, asNode(t.Nname), asNode(t.Nname), asNode(t.Nname).Op)
					}
					return n
				}
			}

			if !expect.Sym.IsBlank() {
				psess.
					Fatalf("nodarg: did not find node in dcl list: %v", expect)
			}
		}

		n = psess.newname(psess.lookup("__"))
		n.Type = t.Type
		if t.Offset == BADWIDTH {
			psess.
				Fatalf("nodarg: offset not computed for %v", t)
		}
		n.Xoffset = t.Offset
		n.Orig = asNode(t.Nname)
	}

	if n.isBlank() {
		n.Sym = psess.lookup("__")
	}

	if fp != 0 {
		psess.
			Fatalf("bad fp: %v", fp)
	}

	n.Op = OINDREGSP
	n.Xoffset += psess.Ctxt.FixedFrameSize()
	n.SetTypecheck(1)
	n.SetAddrtaken(true)
	return n
}

// package all the arguments that match a ... T parameter into a []T.
func (psess *PackageSession) mkdotargslice(typ *types.Type, args []*Node, init *Nodes, ddd *Node) *Node {
	esc := uint16(EscUnknown)
	if ddd != nil {
		esc = ddd.Esc
	}

	if len(args) == 0 {
		n := psess.nodnil()
		n.Type = typ
		return n
	}

	n := psess.nod(OCOMPLIT, nil, psess.typenod(typ))
	if ddd != nil && psess.prealloc[ddd] != nil {
		psess.
			prealloc[n] = psess.prealloc[ddd]
	}
	n.List.Set(args)
	n.Esc = esc
	n = psess.typecheck(n, Erv)
	if n.Type == nil {
		psess.
			Fatalf("mkdotargslice: typecheck failed")
	}
	n = psess.walkexpr(n, init)
	return n
}

// check assign expression list to
// a type list. called in
//	return expr-list
//	func(expr-list)
func (psess *PackageSession) ascompatte(call *Node, isddd bool, lhs *types.Type, rhs []*Node, fp int, init *Nodes) []*Node {

	if len(rhs) == 1 && rhs[0].Type.IsFuncArgStruct() {

		if psess.eqtypenoname(rhs[0].Type, lhs) {
			nl := psess.nodarg(lhs, fp)
			nr := psess.nod(OCONVNOP, rhs[0], nil)
			nr.Type = nl.Type
			n := psess.convas(psess.nod(OAS, nl, nr), init)
			n.SetTypecheck(1)
			return []*Node{n}
		}

		// conversions involved.
		// copy into temporaries.
		var tmps []*Node
		for _, nr := range rhs[0].Type.FieldSlice(psess.types) {
			tmps = append(tmps, psess.temp(nr.Type))
		}

		a := psess.nod(OAS2, nil, nil)
		a.List.Set(tmps)
		a.Rlist.Set(rhs)
		a = psess.typecheck(a, Etop)
		a = psess.walkstmt(a)
		init.Append(a)

		rhs = tmps
	}

	// For each parameter (LHS), assign its corresponding argument (RHS).
	// If there's a ... parameter (which is only valid as the final
	// parameter) and this is not a ... call expression,
	// then assign the remaining arguments as a slice.
	var nn []*Node
	for i, nl := range lhs.FieldSlice(psess.types) {
		var nr *Node
		if nl.Isddd() && !isddd {
			nr = psess.mkdotargslice(nl.Type, rhs[i:], init, call.Right)
		} else {
			nr = rhs[i]
		}

		a := psess.nod(OAS, psess.nodarg(nl, fp), nr)
		a = psess.convas(a, init)
		a.SetTypecheck(1)
		nn = append(nn, a)
	}

	return nn
}

// generate code for print
func (psess *PackageSession) walkprint(nn *Node, init *Nodes) *Node {
	psess.
		walkexprlistcheap(nn.List.Slice(), init)

	if nn.Op == OPRINTN {
		s := nn.List.Slice()
		t := make([]*Node, 0, len(s)*2)
		for i, n := range s {
			if i != 0 {
				t = append(t, psess.nodstr(" "))
			}
			t = append(t, n)
		}
		t = append(t, psess.nodstr("\n"))
		nn.List.Set(t)
	}

	s := nn.List.Slice()
	t := make([]*Node, 0, len(s))
	for i := 0; i < len(s); {
		var strs []string
		for i < len(s) && psess.Isconst(s[i], CTSTR) {
			strs = append(strs, s[i].Val().U.(string))
			i++
		}
		if len(strs) > 0 {
			t = append(t, psess.nodstr(strings.Join(strs, "")))
		}
		if i < len(s) {
			t = append(t, s[i])
			i++
		}
	}
	nn.List.Set(t)

	calls := []*Node{psess.mkcall("printlock", nil, init)}
	for i, n := range nn.List.Slice() {
		if n.Op == OLITERAL {
			switch n.Val().Ctype(psess) {
			case CTRUNE:
				n = psess.defaultlit(n, psess.types.Runetype)

			case CTINT:
				n = psess.defaultlit(n, psess.types.Types[TINT64])

			case CTFLT:
				n = psess.defaultlit(n, psess.types.Types[TFLOAT64])
			}
		}

		if n.Op != OLITERAL && n.Type != nil && n.Type.Etype == TIDEAL {
			n = psess.defaultlit(n, psess.types.Types[TINT64])
		}
		n = psess.defaultlit(n, nil)
		nn.List.SetIndex(i, n)
		if n.Type == nil || n.Type.Etype == TFORW {
			continue
		}

		var on *Node
		switch n.Type.Etype {
		case TINTER:
			if n.Type.IsEmptyInterface(psess.types) {
				on = psess.syslook("printeface")
			} else {
				on = psess.syslook("printiface")
			}
			on = psess.substArgTypes(on, n.Type)
		case TPTR32, TPTR64, TCHAN, TMAP, TFUNC, TUNSAFEPTR:
			on = psess.syslook("printpointer")
			on = psess.substArgTypes(on, n.Type)
		case TSLICE:
			on = psess.syslook("printslice")
			on = psess.substArgTypes(on, n.Type)
		case TUINT, TUINT8, TUINT16, TUINT32, TUINT64, TUINTPTR:
			if psess.isRuntimePkg(n.Type.Sym.Pkg) && n.Type.Sym.Name == "hex" {
				on = psess.syslook("printhex")
			} else {
				on = psess.syslook("printuint")
			}
		case TINT, TINT8, TINT16, TINT32, TINT64:
			on = psess.syslook("printint")
		case TFLOAT32, TFLOAT64:
			on = psess.syslook("printfloat")
		case TCOMPLEX64, TCOMPLEX128:
			on = psess.syslook("printcomplex")
		case TBOOL:
			on = psess.syslook("printbool")
		case TSTRING:
			cs := ""
			if psess.Isconst(n, CTSTR) {
				cs = n.Val().U.(string)
			}
			switch cs {
			case " ":
				on = psess.syslook("printsp")
			case "\n":
				on = psess.syslook("printnl")
			default:
				on = psess.syslook("printstring")
			}
		default:
			psess.
				badtype(OPRINT, n.Type, nil)
			continue
		}

		r := psess.nod(OCALL, on, nil)
		if params := on.Type.Params(psess.types).FieldSlice(psess.types); len(params) > 0 {
			t := params[0].Type
			if !psess.eqtype(t, n.Type) {
				n = psess.nod(OCONV, n, nil)
				n.Type = t
			}
			r.List.Append(n)
		}
		calls = append(calls, r)
	}

	calls = append(calls, psess.mkcall("printunlock", nil, init))
	psess.
		typecheckslice(calls, Etop)
	psess.
		walkexprlist(calls, init)

	r := psess.nod(OEMPTY, nil, nil)
	r = psess.typecheck(r, Etop)
	r = psess.walkexpr(r, init)
	r.Ninit.Set(calls)
	return r
}

func (psess *PackageSession) callnew(t *types.Type) *Node {
	if t.NotInHeap() {
		psess.
			yyerror("%v is go:notinheap; heap allocation disallowed", t)
	}
	psess.
		dowidth(t)
	fn := psess.syslook("newobject")
	fn = psess.substArgTypes(fn, t)
	v := psess.mkcall1(fn, psess.types.NewPtr(t), nil, psess.typename(t))
	v.SetNonNil(true)
	return v
}

func (psess *PackageSession) iscallret(n *Node) bool {
	if n == nil {
		return false
	}
	n = psess.outervalue(n)
	return n.Op == OINDREGSP
}

// isReflectHeaderDataField reports whether l is an expression p.Data
// where p has type reflect.SliceHeader or reflect.StringHeader.
func (psess *PackageSession) isReflectHeaderDataField(l *Node) bool {
	if l.Type != psess.types.Types[TUINTPTR] {
		return false
	}

	var tsym *types.Sym
	switch l.Op {
	case ODOT:
		tsym = l.Left.Type.Sym
	case ODOTPTR:
		tsym = l.Left.Type.Elem(psess.types).Sym
	default:
		return false
	}

	if tsym == nil || l.Sym.Name != "Data" || tsym.Pkg.Path != "reflect" {
		return false
	}
	return tsym.Name == "SliceHeader" || tsym.Name == "StringHeader"
}

func (psess *PackageSession) convas(n *Node, init *Nodes) *Node {
	if n.Op != OAS {
		psess.
			Fatalf("convas: not OAS %v", n.Op)
	}
	defer psess.updateHasCall(n)

	n.SetTypecheck(1)

	if n.Left == nil || n.Right == nil {
		return n
	}

	lt := n.Left.Type
	rt := n.Right.Type
	if lt == nil || rt == nil {
		return n
	}

	if n.Left.isBlank() {
		n.Right = psess.defaultlit(n.Right, nil)
		return n
	}

	if !psess.eqtype(lt, rt) {
		n.Right = psess.assignconv(n.Right, lt, "assignment")
		n.Right = psess.walkexpr(n.Right, init)
	}
	psess.
		dowidth(n.Right.Type)

	return n
}

// from ascompat[te]
// evaluating actual function arguments.
//	f(a,b)
// if there is exactly one function expr,
// then it is done first. otherwise must
// make temp variables
func (psess *PackageSession) reorder1(all []*Node) []*Node {

	funcCalls := 0
	if !psess.instrumenting {
		if len(all) == 1 {
			return all
		}

		for _, n := range all {
			psess.
				updateHasCall(n)
			if n.HasCall() {
				funcCalls++
			}
		}
		if funcCalls == 0 {
			return all
		}
	}

	var g []*Node // fncalls assigned to tempnames
	var f *Node   // last fncall assigned to stack
	var r []*Node // non fncalls and tempnames assigned to stack
	d := 0
	for _, n := range all {
		if !psess.instrumenting {
			if !n.HasCall() {
				r = append(r, n)
				continue
			}

			d++
			if d == funcCalls {
				f = n
				continue
			}
		}

		a := psess.temp(n.Right.Type)

		a = psess.nod(OAS, a, n.Right)
		g = append(g, a)

		n.Right = a.Left

		r = append(r, n)
	}

	if f != nil {
		g = append(g, f)
	}
	return append(g, r...)
}

// from ascompat[ee]
//	a,b = c,d
// simultaneous assignment. there cannot
// be later use of an earlier lvalue.
//
// function calls have been removed.
func (psess *PackageSession) reorder3(all []*Node) []*Node {
	// If a needed expression may be affected by an
	// earlier assignment, make an early copy of that
	// expression and use the copy instead.
	var early []*Node

	var mapinit Nodes
	for i, n := range all {
		l := n.Left

		for {
			if l.Op == ODOT || l.Op == OPAREN {
				l = l.Left
				continue
			}

			if l.Op == OINDEX && l.Left.Type.IsArray() {
				l.Right = psess.reorder3save(l.Right, all, i, &early)
				l = l.Left
				continue
			}

			break
		}

		switch l.Op {
		default:
			psess.
				Fatalf("reorder3 unexpected lvalue %#v", l.Op)

		case ONAME:
			break

		case OINDEX, OINDEXMAP:
			l.Left = psess.reorder3save(l.Left, all, i, &early)
			l.Right = psess.reorder3save(l.Right, all, i, &early)
			if l.Op == OINDEXMAP {
				all[i] = psess.convas(all[i], &mapinit)
			}

		case OIND, ODOTPTR:
			l.Left = psess.reorder3save(l.Left, all, i, &early)
		}

		all[i].Right = psess.reorder3save(all[i].Right, all, i, &early)
	}

	early = append(mapinit.Slice(), early...)
	return append(early, all...)
}

// if the evaluation of *np would be affected by the
// assignments in all up to but not including the ith assignment,
// copy into a temporary during *early and
// replace *np with that temp.
// The result of reorder3save MUST be assigned back to n, e.g.
// 	n.Left = reorder3save(n.Left, all, i, early)
func (psess *PackageSession) reorder3save(n *Node, all []*Node, i int, early *[]*Node) *Node {
	if !psess.aliased(n, all, i) {
		return n
	}

	q := psess.temp(n.Type)
	q = psess.nod(OAS, q, n)
	q = psess.typecheck(q, Etop)
	*early = append(*early, q)
	return q.Left
}

// what's the outer value that a write to n affects?
// outer value means containing struct or array.
func (psess *PackageSession) outervalue(n *Node) *Node {
	for {
		switch n.Op {
		case OXDOT:
			psess.
				Fatalf("OXDOT in walk")
		case ODOT, OPAREN, OCONVNOP:
			n = n.Left
			continue
		case OINDEX:
			if n.Left.Type != nil && n.Left.Type.IsArray() {
				n = n.Left
				continue
			}
		}

		return n
	}
}

// Is it possible that the computation of n might be
// affected by writes in as up to but not including the ith element?
func (psess *PackageSession) aliased(n *Node, all []*Node, i int) bool {
	if n == nil {
		return false
	}

	for n.Op == ODOT {
		n = n.Left
	}

	memwrite := false
	varwrite := false
	for _, an := range all[:i] {
		a := psess.outervalue(an.Left)

		for a.Op == ODOT {
			a = a.Left
		}

		if a.Op != ONAME {
			memwrite = true
			continue
		}

		switch n.Class() {
		default:
			varwrite = true
			continue

		case PAUTO, PPARAM, PPARAMOUT:
			if n.Addrtaken() {
				varwrite = true
				continue
			}

			if vmatch2(a, n) {

				return true
			}
		}
	}

	if !memwrite && !varwrite {
		return false
	}

	if psess.varexpr(n) {
		return false
	}

	return true
}

// does the evaluation of n only refer to variables
// whose addresses have not been taken?
// (and no other memory)
func (psess *PackageSession) varexpr(n *Node) bool {
	if n == nil {
		return true
	}

	switch n.Op {
	case OLITERAL:
		return true

	case ONAME:
		switch n.Class() {
		case PAUTO, PPARAM, PPARAMOUT:
			if !n.Addrtaken() {
				return true
			}
		}

		return false

	case OADD,
		OSUB,
		OOR,
		OXOR,
		OMUL,
		ODIV,
		OMOD,
		OLSH,
		ORSH,
		OAND,
		OANDNOT,
		OPLUS,
		OMINUS,
		OCOM,
		OPAREN,
		OANDAND,
		OOROR,
		OCONV,
		OCONVNOP,
		OCONVIFACE,
		ODOTTYPE:
		return psess.varexpr(n.Left) && psess.varexpr(n.Right)

	case ODOT:
		psess.
			Fatalf("varexpr unexpected ODOT")
	}

	return false
}

// is the name l mentioned in r?
func vmatch2(l *Node, r *Node) bool {
	if r == nil {
		return false
	}
	switch r.Op {

	case ONAME:
		return l == r

	case OLITERAL:
		return false
	}

	if vmatch2(l, r.Left) {
		return true
	}
	if vmatch2(l, r.Right) {
		return true
	}
	for _, n := range r.List.Slice() {
		if vmatch2(l, n) {
			return true
		}
	}
	return false
}

// is any name mentioned in l also mentioned in r?
// called by sinit.go
func vmatch1(l *Node, r *Node) bool {

	if l == nil || r == nil {
		return false
	}
	switch l.Op {
	case ONAME:
		switch l.Class() {
		case PPARAM, PAUTO:
			break

		default:

			if r.HasCall() {
				return true
			}
		}

		return vmatch2(l, r)

	case OLITERAL:
		return false
	}

	if vmatch1(l.Left, r) {
		return true
	}
	if vmatch1(l.Right, r) {
		return true
	}
	for _, n := range l.List.Slice() {
		if vmatch1(n, r) {
			return true
		}
	}
	return false
}

// paramstoheap returns code to allocate memory for heap-escaped parameters
// and to copy non-result parameters' values from the stack.
func (psess *PackageSession) paramstoheap(params *types.Type) []*Node {
	var nn []*Node
	for _, t := range params.Fields(psess.types).Slice() {
		v := asNode(t.Nname)
		if v != nil && v.Sym != nil && strings.HasPrefix(v.Sym.Name, "~r") {
			v = nil
		}
		if v == nil {
			continue
		}

		if stackcopy := v.Name.Param.Stackcopy; stackcopy != nil {
			nn = append(nn, psess.walkstmt(psess.nod(ODCL, v, nil)))
			if stackcopy.Class() == PPARAM {
				nn = append(nn, psess.walkstmt(psess.typecheck(psess.nod(OAS, v, stackcopy), Etop)))
			}
		}
	}

	return nn
}

// zeroResults zeros the return values at the start of the function.
// We need to do this very early in the function.  Defer might stop a
// panic and show the return values as they exist at the time of
// panic.  For precise stacks, the garbage collector assumes results
// are always live, so we need to zero them before any allocations,
// even allocations to move params/results to the heap.
// The generated code is added to Curfn's Enter list.
func (psess *PackageSession) zeroResults() {
	for _, f := range psess.Curfn.Type.Results(psess.types).Fields(psess.types).Slice() {
		if v := asNode(f.Nname); v != nil && v.Name.Param.Heapaddr != nil {

			continue
		}
		psess.
			Curfn.Func.Enter.Append(psess.nodl(psess.Curfn.Pos, OAS, psess.nodarg(f, 1), nil))
	}
}

// returnsfromheap returns code to copy values for heap-escaped parameters
// back to the stack.
func (psess *PackageSession) returnsfromheap(params *types.Type) []*Node {
	var nn []*Node
	for _, t := range params.Fields(psess.types).Slice() {
		v := asNode(t.Nname)
		if v == nil {
			continue
		}
		if stackcopy := v.Name.Param.Stackcopy; stackcopy != nil && stackcopy.Class() == PPARAMOUT {
			nn = append(nn, psess.walkstmt(psess.typecheck(psess.nod(OAS, stackcopy, v), Etop)))
		}
	}

	return nn
}

// heapmoves generates code to handle migrating heap-escaped parameters
// between the stack and the heap. The generated code is added to Curfn's
// Enter and Exit lists.
func (psess *PackageSession) heapmoves() {
	lno := psess.lineno
	psess.
		lineno = psess.Curfn.Pos
	nn := psess.paramstoheap(psess.Curfn.Type.Recvs(psess.types))
	nn = append(nn, psess.paramstoheap(psess.Curfn.Type.Params(psess.types))...)
	nn = append(nn, psess.paramstoheap(psess.Curfn.Type.Results(psess.types))...)
	psess.
		Curfn.Func.Enter.Append(nn...)
	psess.
		lineno = psess.Curfn.Func.Endlineno
	psess.
		Curfn.Func.Exit.Append(psess.returnsfromheap(psess.Curfn.Type.Results(psess.types))...)
	psess.
		lineno = lno
}

func (psess *PackageSession) vmkcall(fn *Node, t *types.Type, init *Nodes, va []*Node) *Node {
	if fn.Type == nil || fn.Type.Etype != TFUNC {
		psess.
			Fatalf("mkcall %v %v", fn, fn.Type)
	}

	n := fn.Type.NumParams(psess.types)
	if n != len(va) {
		psess.
			Fatalf("vmkcall %v needs %v args got %v", fn, n, len(va))
	}

	r := psess.nod(OCALL, fn, nil)
	r.List.Set(va)
	if fn.Type.NumResults(psess.types) > 0 {
		r = psess.typecheck(r, Erv|Efnstruct)
	} else {
		r = psess.typecheck(r, Etop)
	}
	r = psess.walkexpr(r, init)
	r.Type = t
	return r
}

func (psess *PackageSession) mkcall(name string, t *types.Type, init *Nodes, args ...*Node) *Node {
	return psess.vmkcall(psess.syslook(name), t, init, args)
}

func (psess *PackageSession) mkcall1(fn *Node, t *types.Type, init *Nodes, args ...*Node) *Node {
	return psess.vmkcall(fn, t, init, args)
}

func (psess *PackageSession) conv(n *Node, t *types.Type) *Node {
	if psess.eqtype(n.Type, t) {
		return n
	}
	n = psess.nod(OCONV, n, nil)
	n.Type = t
	n = psess.typecheck(n, Erv)
	return n
}

// byteindex converts n, which is byte-sized, to a uint8.
// We cannot use conv, because we allow converting bool to uint8 here,
// which is forbidden in user code.
func (psess *PackageSession) byteindex(n *Node) *Node {
	if psess.eqtype(n.Type, psess.types.Types[TUINT8]) {
		return n
	}
	n = psess.nod(OCONV, n, nil)
	n.Type = psess.types.Types[TUINT8]
	n.SetTypecheck(1)
	return n
}

func (psess *PackageSession) chanfn(name string, n int, t *types.Type) *Node {
	if !t.IsChan() {
		psess.
			Fatalf("chanfn %v", t)
	}
	fn := psess.syslook(name)
	switch n {
	default:
		psess.
			Fatalf("chanfn %d", n)
	case 1:
		fn = psess.substArgTypes(fn, t.Elem(psess.types))
	case 2:
		fn = psess.substArgTypes(fn, t.Elem(psess.types), t.Elem(psess.types))
	}
	return fn
}

func (psess *PackageSession) mapfn(name string, t *types.Type) *Node {
	if !t.IsMap() {
		psess.
			Fatalf("mapfn %v", t)
	}
	fn := psess.syslook(name)
	fn = psess.substArgTypes(fn, t.Key(psess.types), t.Elem(psess.types), t.Key(psess.types), t.Elem(psess.types))
	return fn
}

func (psess *PackageSession) mapfndel(name string, t *types.Type) *Node {
	if !t.IsMap() {
		psess.
			Fatalf("mapfn %v", t)
	}
	fn := psess.syslook(name)
	fn = psess.substArgTypes(fn, t.Key(psess.types), t.Elem(psess.types), t.Key(psess.types))
	return fn
}

const (
	mapslow = iota
	mapfast32
	mapfast32ptr
	mapfast64
	mapfast64ptr
	mapfaststr
	nmapfast
)

type mapnames [nmapfast]string

func mkmapnames(base string, ptr string) mapnames {
	return mapnames{base, base + "_fast32", base + "_fast32" + ptr, base + "_fast64", base + "_fast64" + ptr, base + "_faststr"}
}

func (psess *PackageSession) mapfast(t *types.Type) int {

	if t.Elem(psess.types).Width > 128 {
		return mapslow
	}
	switch psess.algtype(t.Key(psess.types)) {
	case AMEM32:
		if !t.Key(psess.types).HasHeapPointer(psess.types) {
			return mapfast32
		}
		if psess.Widthptr == 4 {
			return mapfast32ptr
		}
		psess.
			Fatalf("small pointer %v", t.Key(psess.types))
	case AMEM64:
		if !t.Key(psess.types).HasHeapPointer(psess.types) {
			return mapfast64
		}
		if psess.Widthptr == 8 {
			return mapfast64ptr
		}

	case ASTRING:
		return mapfaststr
	}
	return mapslow
}

func (psess *PackageSession) writebarrierfn(name string, l *types.Type, r *types.Type) *Node {
	fn := psess.syslook(name)
	fn = psess.substArgTypes(fn, l, r)
	return fn
}

func (psess *PackageSession) addstr(n *Node, init *Nodes) *Node {

	c := n.List.Len()

	if c < 2 {
		psess.
			Fatalf("addstr count %d too small", c)
	}

	buf := psess.nodnil()
	if n.Esc == EscNone {
		sz := int64(0)
		for _, n1 := range n.List.Slice() {
			if n1.Op == OLITERAL {
				sz += int64(len(n1.Val().U.(string)))
			}
		}

		if sz < tmpstringbufsize {

			t := psess.types.NewArray(psess.types.Types[TUINT8], tmpstringbufsize)

			buf = psess.nod(OADDR, psess.temp(t), nil)
		}
	}

	args := []*Node{buf}
	for _, n2 := range n.List.Slice() {
		args = append(args, psess.conv(n2, psess.types.Types[TSTRING]))
	}

	var fn string
	if c <= 5 {

		fn = fmt.Sprintf("concatstring%d", c)
	} else {

		fn = "concatstrings"

		t := psess.types.NewSlice(psess.types.Types[TSTRING])
		slice := psess.nod(OCOMPLIT, nil, psess.typenod(t))
		if psess.prealloc[n] != nil {
			psess.
				prealloc[slice] = psess.prealloc[n]
		}
		slice.List.Set(args[1:])
		args = []*Node{buf, slice}
		slice.Esc = EscNone
	}

	cat := psess.syslook(fn)
	r := psess.nod(OCALL, cat, nil)
	r.List.Set(args)
	r = psess.typecheck(r, Erv)
	r = psess.walkexpr(r, init)
	r.Type = n.Type

	return r
}

func (psess *PackageSession) walkAppendArgs(n *Node, init *Nodes) {
	psess.
		walkexprlistsafe(n.List.Slice(), init)

	ls := n.List.Slice()
	for i1, n1 := range ls {
		ls[i1] = psess.cheapexpr(n1, init)
	}
}

// expand append(l1, l2...) to
//   init {
//     s := l1
//     n := len(s) + len(l2)
//     // Compare as uint so growslice can panic on overflow.
//     if uint(n) > uint(cap(s)) {
//       s = growslice(s, n)
//     }
//     s = s[:n]
//     memmove(&s[len(l1)], &l2[0], len(l2)*sizeof(T))
//   }
//   s
//
// l2 is allowed to be a string.
func (psess *PackageSession) appendslice(n *Node, init *Nodes) *Node {
	psess.
		walkAppendArgs(n, init)

	l1 := n.List.First()
	l2 := n.List.Second()

	var l []*Node

	s := psess.temp(l1.Type)
	l = append(l, psess.nod(OAS, s, l1))

	nn := psess.temp(psess.types.Types[TINT])
	l = append(l, psess.nod(OAS, nn, psess.nod(OADD, psess.nod(OLEN, s, nil), psess.nod(OLEN, l2, nil))))

	nif := psess.nod(OIF, nil, nil)
	nif.Left = psess.nod(OGT, psess.nod(OCONV, nn, nil), psess.nod(OCONV, psess.nod(OCAP, s, nil), nil))
	nif.Left.Left.Type = psess.types.Types[TUINT]
	nif.Left.Right.Type = psess.types.Types[TUINT]

	fn := psess.syslook("growslice")
	fn = psess.substArgTypes(fn, s.Type.Elem(psess.types), s.Type.Elem(psess.types))

	nif.Nbody.Set1(psess.nod(OAS, s, psess.mkcall1(fn, s.Type, &nif.Ninit, psess.typename(s.Type.Elem(psess.types)), s, nn)))
	l = append(l, nif)

	nt := psess.nod(OSLICE, s, nil)
	nt.SetSliceBounds(psess, nil, nn, nil)
	l = append(l, psess.nod(OAS, s, nt))

	if l1.Type.Elem(psess.types).HasHeapPointer(psess.types) {

		nptr1 := psess.nod(OSLICE, s, nil)
		nptr1.SetSliceBounds(psess, psess.nod(OLEN, l1, nil), nil, nil)
		nptr2 := l2
		psess.
			Curfn.Func.setWBPos(psess, n.Pos)
		fn := psess.syslook("typedslicecopy")
		fn = psess.substArgTypes(fn, l1.Type, l2.Type)
		var ln Nodes
		ln.Set(l)
		nt := psess.mkcall1(fn, psess.types.Types[TINT], &ln, psess.typename(l1.Type.Elem(psess.types)), nptr1, nptr2)
		l = append(ln.Slice(), nt)
	} else if psess.instrumenting && !psess.compiling_runtime {

		nptr1 := psess.nod(OSLICE, s, nil)
		nptr1.SetSliceBounds(psess, psess.nod(OLEN, l1, nil), nil, nil)
		nptr2 := l2

		var ln Nodes
		ln.Set(l)
		var nt *Node
		if l2.Type.IsString() {
			fn := psess.syslook("slicestringcopy")
			fn = psess.substArgTypes(fn, l1.Type, l2.Type)
			nt = psess.mkcall1(fn, psess.types.Types[TINT], &ln, nptr1, nptr2)
		} else {
			fn := psess.syslook("slicecopy")
			fn = psess.substArgTypes(fn, l1.Type, l2.Type)
			nt = psess.mkcall1(fn, psess.types.Types[TINT], &ln, nptr1, nptr2, psess.nodintconst(s.Type.Elem(psess.types).Width))
		}

		l = append(ln.Slice(), nt)
	} else {

		nptr1 := psess.nod(OINDEX, s, psess.nod(OLEN, l1, nil))
		nptr1.SetBounded(true)

		nptr1 = psess.nod(OADDR, nptr1, nil)

		nptr2 := psess.nod(OSPTR, l2, nil)

		fn := psess.syslook("memmove")
		fn = psess.substArgTypes(fn, s.Type.Elem(psess.types), s.Type.Elem(psess.types))

		var ln Nodes
		ln.Set(l)
		nwid := psess.cheapexpr(psess.conv(psess.nod(OLEN, l2, nil), psess.types.Types[TUINTPTR]), &ln)

		nwid = psess.nod(OMUL, nwid, psess.nodintconst(s.Type.Elem(psess.types).Width))
		nt := psess.mkcall1(fn, nil, &ln, nptr1, nptr2, nwid)
		l = append(ln.Slice(), nt)
	}
	psess.
		typecheckslice(l, Etop)
	psess.
		walkstmtlist(l)
	init.Append(l...)
	return s
}

// isAppendOfMake reports whether n is of the form append(x , make([]T, y)...).
// isAppendOfMake assumes n has already been typechecked.
func (psess *PackageSession) isAppendOfMake(n *Node) bool {
	if psess.Debug['N'] != 0 || psess.instrumenting {
		return false
	}

	if n.Typecheck() == 0 {
		psess.
			Fatalf("missing typecheck: %+v", n)
	}

	if n.Op != OAPPEND || !n.Isddd() || n.List.Len() != 2 {
		return false
	}

	second := n.List.Second()
	if second.Op != OMAKESLICE || second.Right != nil {
		return false
	}

	y := second.Left
	if !psess.Isconst(y, CTINT) && y.Type.Etype != TINT {
		return false
	}

	return true
}

// extendslice rewrites append(l1, make([]T, l2)...) to
//   init {
//     if l2 < 0 {
//       panicmakeslicelen()
//     }
//     s := l1
//     n := len(s) + l2
//     // Compare n and s as uint so growslice can panic on overflow of len(s) + l2.
//     // cap is a positive int and n can become negative when len(s) + l2
//     // overflows int. Interpreting n when negative as uint makes it larger
//     // than cap(s). growslice will check the int n arg and panic if n is
//     // negative. This prevents the overflow from being undetected.
//     if uint(n) > uint(cap(s)) {
//       s = growslice(T, s, n)
//     }
//     s = s[:n]
//     lptr := &l1[0]
//     sptr := &s[0]
//     if lptr == sptr || !hasPointers(T) {
//       // growslice did not clear the whole underlying array (or did not get called)
//       hp := &s[len(l1)]
//       hn := l2 * sizeof(T)
//       memclr(hp, hn)
//     }
//   }
//   s
func (psess *PackageSession) extendslice(n *Node, init *Nodes) *Node {

	l2 := psess.conv(n.List.Second().Left, psess.types.Types[TINT])
	l2 = psess.typecheck(l2, Erv)
	n.List.SetSecond(l2)
	psess.
		walkAppendArgs(n, init)

	l1 := n.List.First()
	l2 = n.List.Second()

	var nodes []*Node

	nifneg := psess.nod(OIF, psess.nod(OLT, l2, psess.nodintconst(0)), nil)
	nifneg.SetLikely(false)

	nifneg.Nbody.Set1(psess.mkcall("panicmakeslicelen", nil, init))
	nodes = append(nodes, nifneg)

	s := psess.temp(l1.Type)
	nodes = append(nodes, psess.nod(OAS, s, l1))

	elemtype := s.Type.Elem(psess.types)

	nn := psess.temp(psess.types.Types[TINT])
	nodes = append(nodes, psess.nod(OAS, nn, psess.nod(OADD, psess.nod(OLEN, s, nil), l2)))

	nuint := psess.conv(nn, psess.types.Types[TUINT])
	capuint := psess.conv(psess.nod(OCAP, s, nil), psess.types.Types[TUINT])
	nif := psess.nod(OIF, psess.nod(OGT, nuint, capuint), nil)

	fn := psess.syslook("growslice")
	fn = psess.substArgTypes(fn, elemtype, elemtype)

	nif.Nbody.Set1(psess.nod(OAS, s, psess.mkcall1(fn, s.Type, &nif.Ninit, psess.typename(elemtype), s, nn)))
	nodes = append(nodes, nif)

	nt := psess.nod(OSLICE, s, nil)
	nt.SetSliceBounds(psess, nil, nn, nil)
	nodes = append(nodes, psess.nod(OAS, s, nt))

	l1ptr := psess.temp(l1.Type.Elem(psess.types).PtrTo(psess.types))
	tmp := psess.nod(OSPTR, l1, nil)
	nodes = append(nodes, psess.nod(OAS, l1ptr, tmp))

	sptr := psess.temp(elemtype.PtrTo(psess.types))
	tmp = psess.nod(OSPTR, s, nil)
	nodes = append(nodes, psess.nod(OAS, sptr, tmp))

	hp := psess.nod(OINDEX, s, psess.nod(OLEN, l1, nil))
	hp.SetBounded(true)
	hp = psess.nod(OADDR, hp, nil)
	hp = psess.nod(OCONVNOP, hp, nil)
	hp.Type = psess.types.Types[TUNSAFEPTR]

	hn := psess.nod(OMUL, l2, psess.nodintconst(elemtype.Width))
	hn = psess.conv(hn, psess.types.Types[TUINTPTR])

	clrname := "memclrNoHeapPointers"
	hasPointers := psess.types.Haspointers(elemtype)
	if hasPointers {
		clrname = "memclrHasPointers"
	}

	var clr Nodes
	clrfn := psess.mkcall(clrname, nil, &clr, hp, hn)
	clr.Append(clrfn)

	if hasPointers {

		nifclr := psess.nod(OIF, psess.nod(OEQ, l1ptr, sptr), nil)
		nifclr.Nbody = clr
		nodes = append(nodes, nifclr)
	} else {
		nodes = append(nodes, clr.Slice()...)
	}
	psess.
		typecheckslice(nodes, Etop)
	psess.
		walkstmtlist(nodes)
	init.Append(nodes...)
	return s
}

// Rewrite append(src, x, y, z) so that any side effects in
// x, y, z (including runtime panics) are evaluated in
// initialization statements before the append.
// For normal code generation, stop there and leave the
// rest to cgen_append.
//
// For race detector, expand append(src, a [, b]* ) to
//
//   init {
//     s := src
//     const argc = len(args) - 1
//     if cap(s) - len(s) < argc {
//	    s = growslice(s, len(s)+argc)
//     }
//     n := len(s)
//     s = s[:n+argc]
//     s[n] = a
//     s[n+1] = b
//     ...
//   }
//   s
func (psess *PackageSession) walkappend(n *Node, init *Nodes, dst *Node) *Node {
	if !psess.samesafeexpr(dst, n.List.First()) {
		n.List.SetFirst(psess.safeexpr(n.List.First(), init))
		n.List.SetFirst(psess.walkexpr(n.List.First(), init))
	}
	psess.
		walkexprlistsafe(n.List.Slice()[1:], init)

	ls := n.List.Slice()[1:]
	for i, n := range ls {
		ls[i] = psess.cheapexpr(n, init)
	}

	nsrc := n.List.First()

	argc := n.List.Len() - 1
	if argc < 1 {
		return nsrc
	}

	if !psess.instrumenting || psess.compiling_runtime {
		return n
	}

	var l []*Node

	ns := psess.temp(nsrc.Type)
	l = append(l, psess.nod(OAS, ns, nsrc))

	na := psess.nodintconst(int64(argc))
	nx := psess.nod(OIF, nil, nil)
	nx.Left = psess.nod(OLT, psess.nod(OSUB, psess.nod(OCAP, ns, nil), psess.nod(OLEN, ns, nil)), na)

	fn := psess.syslook("growslice")
	fn = psess.substArgTypes(fn, ns.Type.Elem(psess.types), ns.Type.Elem(psess.types))

	nx.Nbody.Set1(psess.nod(OAS, ns, psess.
		mkcall1(fn, ns.Type, &nx.Ninit, psess.typename(ns.Type.Elem(psess.types)), ns, psess.
			nod(OADD, psess.nod(OLEN, ns, nil), na))))

	l = append(l, nx)

	nn := psess.temp(psess.types.Types[TINT])
	l = append(l, psess.nod(OAS, nn, psess.nod(OLEN, ns, nil)))

	nx = psess.nod(OSLICE, ns, nil)
	nx.SetSliceBounds(psess, nil, psess.nod(OADD, nn, na), nil)
	l = append(l, psess.nod(OAS, ns, nx))

	ls = n.List.Slice()[1:]
	for i, n := range ls {
		nx = psess.nod(OINDEX, ns, nn)
		nx.SetBounded(true)
		l = append(l, psess.nod(OAS, nx, n))
		if i+1 < len(ls) {
			l = append(l, psess.nod(OAS, nn, psess.nod(OADD, nn, psess.nodintconst(1))))
		}
	}
	psess.
		typecheckslice(l, Etop)
	psess.
		walkstmtlist(l)
	init.Append(l...)
	return ns
}

// Lower copy(a, b) to a memmove call or a runtime call.
//
// init {
//   n := len(a)
//   if n > len(b) { n = len(b) }
//   if a.ptr != b.ptr { memmove(a.ptr, b.ptr, n*sizeof(elem(a))) }
// }
// n;
//
// Also works if b is a string.
//
func (psess *PackageSession) copyany(n *Node, init *Nodes, runtimecall bool) *Node {
	if n.Left.Type.Elem(psess.types).HasHeapPointer(psess.types) {
		psess.
			Curfn.Func.setWBPos(psess, n.Pos)
		fn := psess.writebarrierfn("typedslicecopy", n.Left.Type, n.Right.Type)
		return psess.mkcall1(fn, n.Type, init, psess.typename(n.Left.Type.Elem(psess.types)), n.Left, n.Right)
	}

	if runtimecall {
		if n.Right.Type.IsString() {
			fn := psess.syslook("slicestringcopy")
			fn = psess.substArgTypes(fn, n.Left.Type, n.Right.Type)
			return psess.mkcall1(fn, n.Type, init, n.Left, n.Right)
		}

		fn := psess.syslook("slicecopy")
		fn = psess.substArgTypes(fn, n.Left.Type, n.Right.Type)
		return psess.mkcall1(fn, n.Type, init, n.Left, n.Right, psess.nodintconst(n.Left.Type.Elem(psess.types).Width))
	}

	n.Left = psess.walkexpr(n.Left, init)
	n.Right = psess.walkexpr(n.Right, init)
	nl := psess.temp(n.Left.Type)
	nr := psess.temp(n.Right.Type)
	var l []*Node
	l = append(l, psess.nod(OAS, nl, n.Left))
	l = append(l, psess.nod(OAS, nr, n.Right))

	nfrm := psess.nod(OSPTR, nr, nil)
	nto := psess.nod(OSPTR, nl, nil)

	nlen := psess.temp(psess.types.Types[TINT])

	l = append(l, psess.nod(OAS, nlen, psess.nod(OLEN, nl, nil)))

	nif := psess.nod(OIF, nil, nil)

	nif.Left = psess.nod(OGT, nlen, psess.nod(OLEN, nr, nil))
	nif.Nbody.Append(psess.nod(OAS, nlen, psess.nod(OLEN, nr, nil)))
	l = append(l, nif)

	ne := psess.nod(OIF, psess.nod(ONE, nto, nfrm), nil)
	ne.SetLikely(true)
	l = append(l, ne)

	fn := psess.syslook("memmove")
	fn = psess.substArgTypes(fn, nl.Type.Elem(psess.types), nl.Type.Elem(psess.types))
	nwid := psess.temp(psess.types.Types[TUINTPTR])
	setwid := psess.nod(OAS, nwid, psess.conv(nlen, psess.types.Types[TUINTPTR]))
	ne.Nbody.Append(setwid)
	nwid = psess.nod(OMUL, nwid, psess.nodintconst(nl.Type.Elem(psess.types).Width))
	call := psess.mkcall1(fn, nil, init, nto, nfrm, nwid)
	ne.Nbody.Append(call)
	psess.
		typecheckslice(l, Etop)
	psess.
		walkstmtlist(l)
	init.Append(l...)
	return nlen
}

func (psess *PackageSession) eqfor(t *types.Type) (n *Node, needsize bool) {

	switch a, _ := psess.algtype1(t); a {
	case AMEM:
		n := psess.syslook("memequal")
		n = psess.substArgTypes(n, t, t)
		return n, true
	case ASPECIAL:
		sym := psess.typesymprefix(".eq", t)
		n := psess.newname(sym)
		n.SetClass(PFUNC)
		n.Type = psess.functype(nil, []*Node{psess.
			anonfield(psess.types.NewPtr(t)), psess.
			anonfield(psess.types.NewPtr(t)),
		}, []*Node{psess.
			anonfield(psess.types.Types[TBOOL]),
		})
		return n, false
	}
	psess.
		Fatalf("eqfor %v", t)
	return nil, false
}

// The result of walkcompare MUST be assigned back to n, e.g.
// 	n.Left = walkcompare(n.Left, init)
func (psess *PackageSession) walkcompare(n *Node, init *Nodes) *Node {
	// Given interface value l and concrete value r, rewrite
	//   l == r
	// into types-equal && data-equal.
	// This is efficient, avoids allocations, and avoids runtime calls.
	var l, r *Node
	if n.Left.Type.IsInterface() && !n.Right.Type.IsInterface() {
		l = n.Left
		r = n.Right
	} else if !n.Left.Type.IsInterface() && n.Right.Type.IsInterface() {
		l = n.Right
		r = n.Left
	}

	if l != nil {

		eq := n.Op
		var andor Op
		if eq == OEQ {
			andor = OANDAND
		} else {
			andor = OOROR
		}
		// Check for types equal.
		// For empty interface, this is:
		//   l.tab == type(r)
		// For non-empty interface, this is:
		//   l.tab != nil && l.tab._type == type(r)
		var eqtype *Node
		tab := psess.nod(OITAB, l, nil)
		rtyp := psess.typename(r.Type)
		if l.Type.IsEmptyInterface(psess.types) {
			tab.Type = psess.types.NewPtr(psess.types.Types[TUINT8])
			tab.SetTypecheck(1)
			eqtype = psess.nod(eq, tab, rtyp)
		} else {
			nonnil := psess.nod(psess.brcom(eq), psess.nodnil(), tab)
			match := psess.nod(eq, psess.itabType(tab), rtyp)
			eqtype = psess.nod(andor, nonnil, match)
		}

		eqdata := psess.nod(eq, psess.ifaceData(l, r.Type), r)

		expr := psess.nod(andor, eqtype, eqdata)
		n = psess.finishcompare(n, expr, init)
		return n
	}

	t := n.Left.Type
	var inline bool

	maxcmpsize := int64(4)
	unalignedLoad := psess.canMergeLoads()
	if unalignedLoad {

		maxcmpsize = 2 * int64(psess.thearch.LinkArch.RegSize)
	}

	switch t.Etype {
	default:
		return n
	case TARRAY:

		inline = t.NumElem(psess.types) <= 1 || (psess.issimple[t.Elem(psess.types).Etype] && (t.NumElem(psess.types) <= 4 || t.Elem(psess.types).Width*t.NumElem(psess.types) <= maxcmpsize))
	case TSTRUCT:
		inline = t.NumComponents(psess.types, types.IgnoreBlankFields) <= 4
	}

	cmpl := n.Left
	for cmpl != nil && cmpl.Op == OCONVNOP {
		cmpl = cmpl.Left
	}
	cmpr := n.Right
	for cmpr != nil && cmpr.Op == OCONVNOP {
		cmpr = cmpr.Left
	}

	if !inline {
		if isvaluelit(cmpl) {
			var_ := psess.temp(cmpl.Type)
			psess.
				anylit(cmpl, var_, init)
			cmpl = var_
		}
		if isvaluelit(cmpr) {
			var_ := psess.temp(cmpr.Type)
			psess.
				anylit(cmpr, var_, init)
			cmpr = var_
		}
		if !islvalue(cmpl) || !islvalue(cmpr) {
			psess.
				Fatalf("arguments of comparison must be lvalues - %v %v", cmpl, cmpr)
		}

		pl := psess.temp(psess.types.NewPtr(t))
		al := psess.nod(OAS, pl, psess.nod(OADDR, cmpl, nil))
		al = psess.typecheck(al, Etop)
		init.Append(al)

		pr := psess.temp(psess.types.NewPtr(t))
		ar := psess.nod(OAS, pr, psess.nod(OADDR, cmpr, nil))
		ar = psess.typecheck(ar, Etop)
		init.Append(ar)

		fn, needsize := psess.eqfor(t)
		call := psess.nod(OCALL, fn, nil)
		call.List.Append(pl)
		call.List.Append(pr)
		if needsize {
			call.List.Append(psess.nodintconst(t.Width))
		}
		res := call
		if n.Op != OEQ {
			res = psess.nod(ONOT, res, nil)
		}
		n = psess.finishcompare(n, res, init)
		return n
	}

	andor := OANDAND
	if n.Op == ONE {
		andor = OOROR
	}
	var expr *Node
	compare := func(el, er *Node) {
		a := psess.nod(n.Op, el, er)
		if expr == nil {
			expr = a
		} else {
			expr = psess.nod(andor, expr, a)
		}
	}
	cmpl = psess.safeexpr(cmpl, init)
	cmpr = psess.safeexpr(cmpr, init)
	if t.IsStruct() {
		for _, f := range t.Fields(psess.types).Slice() {
			sym := f.Sym
			if sym.IsBlank() {
				continue
			}
			compare(psess.
				nodSym(OXDOT, cmpl, sym), psess.
				nodSym(OXDOT, cmpr, sym),
			)
		}
	} else {
		step := int64(1)
		remains := t.NumElem(psess.types) * t.Elem(psess.types).Width
		combine64bit := unalignedLoad && psess.Widthreg == 8 && t.Elem(psess.types).Width <= 4 && t.Elem(psess.types).IsInteger()
		combine32bit := unalignedLoad && t.Elem(psess.types).Width <= 2 && t.Elem(psess.types).IsInteger()
		combine16bit := unalignedLoad && t.Elem(psess.types).Width == 1 && t.Elem(psess.types).IsInteger()
		for i := int64(0); remains > 0; {
			var convType *types.Type
			switch {
			case remains >= 8 && combine64bit:
				convType = psess.types.Types[TINT64]
				step = 8 / t.Elem(psess.types).Width
			case remains >= 4 && combine32bit:
				convType = psess.types.Types[TUINT32]
				step = 4 / t.Elem(psess.types).Width
			case remains >= 2 && combine16bit:
				convType = psess.types.Types[TUINT16]
				step = 2 / t.Elem(psess.types).Width
			default:
				step = 1
			}
			if step == 1 {
				compare(psess.
					nod(OINDEX, cmpl, psess.nodintconst(i)), psess.
					nod(OINDEX, cmpr, psess.nodintconst(i)),
				)
				i++
				remains -= t.Elem(psess.types).Width
			} else {
				elemType := t.Elem(psess.types).ToUnsigned(psess.types)
				cmplw := psess.nod(OINDEX, cmpl, psess.nodintconst(i))
				cmplw = psess.conv(cmplw, elemType)
				cmplw = psess.conv(cmplw, convType)
				cmprw := psess.nod(OINDEX, cmpr, psess.nodintconst(i))
				cmprw = psess.conv(cmprw, elemType)
				cmprw = psess.conv(cmprw, convType)

				for offset := int64(1); offset < step; offset++ {
					lb := psess.nod(OINDEX, cmpl, psess.nodintconst(i+offset))
					lb = psess.conv(lb, elemType)
					lb = psess.conv(lb, convType)
					lb = psess.nod(OLSH, lb, psess.nodintconst(8*t.Elem(psess.types).Width*offset))
					cmplw = psess.nod(OOR, cmplw, lb)
					rb := psess.nod(OINDEX, cmpr, psess.nodintconst(i+offset))
					rb = psess.conv(rb, elemType)
					rb = psess.conv(rb, convType)
					rb = psess.nod(OLSH, rb, psess.nodintconst(8*t.Elem(psess.types).Width*offset))
					cmprw = psess.nod(OOR, cmprw, rb)
				}
				compare(cmplw, cmprw)
				i += step
				remains -= step * t.Elem(psess.types).Width
			}
		}
	}
	if expr == nil {
		expr = psess.nodbool(n.Op == OEQ)
	}
	n = psess.finishcompare(n, expr, init)
	return n
}

// The result of finishcompare MUST be assigned back to n, e.g.
// 	n.Left = finishcompare(n.Left, x, r, init)
func (psess *PackageSession) finishcompare(n, r *Node, init *Nodes) *Node {
	r = psess.typecheck(r, Erv)
	r = psess.conv(r, n.Type)
	r = psess.walkexpr(r, init)
	return r
}

// isIntOrdering reports whether n is a <, ≤, >, or ≥ ordering between integers.
func (n *Node) isIntOrdering() bool {
	switch n.Op {
	case OLE, OLT, OGE, OGT:
	default:
		return false
	}
	return n.Left.Type.IsInteger() && n.Right.Type.IsInteger()
}

// walkinrange optimizes integer-in-range checks, such as 4 <= x && x < 10.
// n must be an OANDAND or OOROR node.
// The result of walkinrange MUST be assigned back to n, e.g.
// 	n.Left = walkinrange(n.Left)
func (psess *PackageSession) walkinrange(n *Node, init *Nodes) *Node {

	l := n.Left
	r := n.Right
	if !l.isIntOrdering() || !r.isIntOrdering() {
		return n
	}

	a, opl, b := l.Left, l.Op, l.Right
	x, opr, c := r.Left, r.Op, r.Right
	for i := 0; ; i++ {
		if psess.samesafeexpr(b, x) {
			break
		}
		if i == 3 {

			return n
		}
		if i&1 == 0 {
			a, opl, b = b, psess.brrev(opl), a
		} else {
			x, opr, c = c, psess.brrev(opr), x
		}
	}

	negateResult := n.Op == OOROR
	if negateResult {
		opl = psess.brcom(opl)
		opr = psess.brcom(opr)
	}

	cmpdir := func(o Op) int {
		switch o {
		case OLE, OLT:
			return -1
		case OGE, OGT:
			return +1
		}
		psess.
			Fatalf("walkinrange cmpdir %v", o)
		return 0
	}
	if cmpdir(opl) != cmpdir(opr) {

		return n
	}

	switch opl {
	case OGE, OGT:

		a, c = c, a
		opl, opr = psess.brrev(opr), psess.brrev(opl)
	}

	if !psess.Isconst(a, CTINT) || !psess.Isconst(c, CTINT) {
		return n
	}

	if opl == OLT {

		if a.Int64(psess) >= psess.maxintval[b.Type.Etype].Int64(psess) {
			return n
		}
		a = psess.nodintconst(a.Int64(psess) + 1)
		opl = OLE
	}

	bound := c.Int64(psess) - a.Int64(psess)
	if bound < 0 {

		return n
	}

	ut := b.Type.ToUnsigned(psess.types)
	lhs := psess.conv(psess.nod(OSUB, b, a), ut)
	rhs := psess.nodintconst(bound)
	if negateResult {

		opr = psess.brcom(opr)
	}
	cmp := psess.nod(opr, lhs, rhs)
	cmp.Pos = n.Pos
	cmp = psess.addinit(cmp, l.Ninit.Slice())
	cmp = psess.addinit(cmp, r.Ninit.Slice())

	cmp = psess.typecheck(cmp, Erv)

	cmp.Type = n.Type
	cmp = psess.walkexpr(cmp, init)
	return cmp
}

// return 1 if integer n must be in range [0, max), 0 otherwise
func (psess *PackageSession) bounded(n *Node, max int64) bool {
	if n.Type == nil || !n.Type.IsInteger() {
		return false
	}

	sign := n.Type.IsSigned()
	bits := int32(8 * n.Type.Width)

	if psess.smallintconst(n) {
		v := n.Int64(psess)
		return 0 <= v && v < max
	}

	switch n.Op {
	case OAND:
		v := int64(-1)
		if psess.smallintconst(n.Left) {
			v = n.Left.Int64(psess)
		} else if psess.smallintconst(n.Right) {
			v = n.Right.Int64(psess)
		}

		if 0 <= v && v < max {
			return true
		}

	case OMOD:
		if !sign && psess.smallintconst(n.Right) {
			v := n.Right.Int64(psess)
			if 0 <= v && v <= max {
				return true
			}
		}

	case ODIV:
		if !sign && psess.smallintconst(n.Right) {
			v := n.Right.Int64(psess)
			for bits > 0 && v >= 2 {
				bits--
				v >>= 1
			}
		}

	case ORSH:
		if !sign && psess.smallintconst(n.Right) {
			v := n.Right.Int64(psess)
			if v > int64(bits) {
				return true
			}
			bits -= int32(v)
		}
	}

	if !sign && bits <= 62 && 1<<uint(bits) <= max {
		return true
	}

	return false
}

// usemethod checks interface method calls for uses of reflect.Type.Method.
func (psess *PackageSession) usemethod(n *Node) {
	t := n.Left.Type

	if n := t.NumParams(psess.types); n != 1 {
		return
	}
	if n := t.NumResults(psess.types); n != 1 && n != 2 {
		return
	}
	p0 := t.Params(psess.types).Field(psess.types, 0)
	res0 := t.Results(psess.types).Field(psess.types, 0)
	var res1 *types.Field
	if t.NumResults(psess.types) == 2 {
		res1 = t.Results(psess.types).Field(psess.types, 1)
	}

	if res1 == nil {
		if p0.Type.Etype != TINT {
			return
		}
	} else {
		if !p0.Type.IsString() {
			return
		}
		if !res1.Type.IsBoolean() {
			return
		}
	}

	if s := res0.Type.Sym; s != nil && s.Name == "Method" && s.Pkg != nil && s.Pkg.Path == "reflect" {
		psess.
			Curfn.Func.SetReflectMethod(true)
	}
}

func (psess *PackageSession) usefield(n *Node) {
	if psess.objabi.Fieldtrack_enabled == 0 {
		return
	}

	switch n.Op {
	default:
		psess.
			Fatalf("usefield %v", n.Op)

	case ODOT, ODOTPTR:
		break
	}
	if n.Sym == nil {

		return
	}

	t := n.Left.Type
	if t.IsPtr() {
		t = t.Elem(psess.types)
	}
	field := psess.dotField[typeSymKey{t.Orig, n.Sym}]
	if field == nil {
		psess.
			Fatalf("usefield %v %v without paramfld", n.Left.Type, n.Sym)
	}
	if !strings.Contains(field.Note, "go:\"track\"") {
		return
	}

	outer := n.Left.Type
	if outer.IsPtr() {
		outer = outer.Elem(psess.types)
	}
	if outer.Sym == nil {
		psess.
			yyerror("tracked field must be in named struct type")
	}
	if !types.IsExported(field.Sym.Name) {
		psess.
			yyerror("tracked field must be exported (upper case)")
	}

	sym := psess.tracksym(outer, field)
	if psess.Curfn.Func.FieldTrack == nil {
		psess.
			Curfn.Func.FieldTrack = make(map[*types.Sym]struct{})
	}
	psess.
		Curfn.Func.FieldTrack[sym] = struct{}{}
}

func (psess *PackageSession) candiscardlist(l Nodes) bool {
	for _, n := range l.Slice() {
		if !psess.candiscard(n) {
			return false
		}
	}
	return true
}

func (psess *PackageSession) candiscard(n *Node) bool {
	if n == nil {
		return true
	}

	switch n.Op {
	default:
		return false

	case ONAME,
		ONONAME,
		OTYPE,
		OPACK,
		OLITERAL,
		OADD,
		OSUB,
		OOR,
		OXOR,
		OADDSTR,
		OADDR,
		OANDAND,
		OARRAYBYTESTR,
		OARRAYRUNESTR,
		OSTRARRAYBYTE,
		OSTRARRAYRUNE,
		OCAP,
		OCMPIFACE,
		OCMPSTR,
		OCOMPLIT,
		OMAPLIT,
		OSTRUCTLIT,
		OARRAYLIT,
		OSLICELIT,
		OPTRLIT,
		OCONV,
		OCONVIFACE,
		OCONVNOP,
		ODOT,
		OEQ,
		ONE,
		OLT,
		OLE,
		OGT,
		OGE,
		OKEY,
		OSTRUCTKEY,
		OLEN,
		OMUL,
		OLSH,
		ORSH,
		OAND,
		OANDNOT,
		ONEW,
		ONOT,
		OCOM,
		OPLUS,
		OMINUS,
		OOROR,
		OPAREN,
		ORUNESTR,
		OREAL,
		OIMAG,
		OCOMPLEX:
		break

	case ODIV, OMOD:
		if psess.Isconst(n.Right, CTINT) && n.Right.Val().U.(*Mpint).CmpInt64(0) != 0 {
			break
		}
		if psess.Isconst(n.Right, CTFLT) && n.Right.Val().U.(*Mpflt).CmpFloat64(0) != 0 {
			break
		}
		return false

	case OMAKECHAN, OMAKEMAP:
		if psess.Isconst(n.Left, CTINT) && n.Left.Val().U.(*Mpint).CmpInt64(0) == 0 {
			break
		}
		return false

	case OMAKESLICE:
		return false
	}

	if !psess.candiscard(n.Left) || !psess.candiscard(n.Right) || !psess.candiscardlist(n.Ninit) || !psess.candiscardlist(n.Nbody) || !psess.candiscardlist(n.List) || !psess.candiscardlist(n.Rlist) {
		return false
	}

	return true
}

// The result of wrapCall MUST be assigned back to n, e.g.
// 	n.Left = wrapCall(n.Left, init)
func (psess *PackageSession) wrapCall(n *Node, init *Nodes) *Node {
	if n.Ninit.Len() != 0 {
		psess.
			walkstmtlist(n.Ninit.Slice())
		init.AppendNodes(&n.Ninit)
	}

	t := psess.nod(OTFUNC, nil, nil)
	for i, arg := range n.List.Slice() {
		s := psess.lookupN("a", i)
		t.List.Append(psess.symfield(s, arg.Type))
	}
	psess.
		wrapCall_prgen++
	sym := psess.lookupN("wrap·", psess.wrapCall_prgen)
	fn := psess.dclfunc(sym, t)

	a := psess.nod(n.Op, nil, nil)
	a.List.Set(psess.paramNnames(t.Type))
	a = psess.typecheck(a, Etop)
	fn.Nbody.Set1(a)
	psess.
		funcbody()

	fn = psess.typecheck(fn, Etop)
	psess.
		typecheckslice(fn.Nbody.Slice(), Etop)
	psess.
		xtop = append(psess.xtop, fn)

	a = psess.nod(OCALL, nil, nil)
	a.Left = fn.Func.Nname
	a.List.Set(n.List.Slice())
	a = psess.typecheck(a, Etop)
	a = psess.walkexpr(a, init)
	return a
}

// substArgTypes substitutes the given list of types for
// successive occurrences of the "any" placeholder in the
// type syntax expression n.Type.
// The result of substArgTypes MUST be assigned back to old, e.g.
// 	n.Left = substArgTypes(n.Left, t1, t2)
func (psess *PackageSession) substArgTypes(old *Node, types_ ...*types.Type) *Node {
	n := old.copy()

	for _, t := range types_ {
		psess.
			dowidth(t)
	}
	n.Type = psess.types.SubstAny(n.Type, &types_)
	if len(types_) > 0 {
		psess.
			Fatalf("substArgTypes: too many argument types")
	}
	return n
}

// canMergeLoads reports whether the backend optimization passes for
// the current architecture can combine adjacent loads into a single
// larger, possibly unaligned, load. Note that currently the
// optimizations must be able to handle little endian byte order.
func (psess *PackageSession) canMergeLoads() bool {
	switch psess.thearch.LinkArch.Family {
	case sys.ARM64, sys.AMD64, sys.I386, sys.S390X:
		return true
	case sys.PPC64:

		return psess.thearch.LinkArch.ByteOrder == binary.LittleEndian
	}
	return false
}

// isRuneCount reports whether n is of the form len([]rune(string)).
// These are optimized into a call to runtime.countrunes.
func (psess *PackageSession) isRuneCount(n *Node) bool {
	return psess.Debug['N'] == 0 && !psess.instrumenting && n.Op == OLEN && n.Left.Op == OSTRARRAYRUNE
}

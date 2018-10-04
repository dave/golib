// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"strings"
)

// The constant is known to runtime.
const tmpstringbufsize = 32

func (pstate *PackageState) walk(fn *Node) {
	pstate.Curfn = fn

	if pstate.Debug['W'] != 0 {
		s := fmt.Sprintf("\nbefore walk %v", pstate.Curfn.Func.Nname.Sym)
		dumplist(s, pstate.Curfn.Nbody)
	}

	lno := pstate.lineno

	// Final typecheck for any unused variables.
	for i, ln := range fn.Func.Dcl {
		if ln.Op == ONAME && (ln.Class() == PAUTO || ln.Class() == PAUTOHEAP) {
			ln = pstate.typecheck(ln, Erv|Easgn)
			fn.Func.Dcl[i] = ln
		}
	}

	// Propagate the used flag for typeswitch variables up to the NONAME in it's definition.
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
			pstate.yyerrorl(defn.Left.Pos, "%v declared and not used", ln.Sym)
			defn.Left.Name.SetUsed(true) // suppress repeats
		} else {
			pstate.yyerrorl(ln.Pos, "%v declared and not used", ln.Sym)
		}
	}

	pstate.lineno = lno
	if pstate.nerrors != 0 {
		return
	}
	pstate.walkstmtlist(pstate.Curfn.Nbody.Slice())
	if pstate.Debug['W'] != 0 {
		s := fmt.Sprintf("after walk %v", pstate.Curfn.Func.Nname.Sym)
		dumplist(s, pstate.Curfn.Nbody)
	}

	pstate.zeroResults()
	pstate.heapmoves()
	if pstate.Debug['W'] != 0 && pstate.Curfn.Func.Enter.Len() > 0 {
		s := fmt.Sprintf("enter %v", pstate.Curfn.Func.Nname.Sym)
		dumplist(s, pstate.Curfn.Func.Enter)
	}
}

func (pstate *PackageState) walkstmtlist(s []*Node) {
	for i := range s {
		s[i] = pstate.walkstmt(s[i])
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
			// stop early - parameters are over
			return false
		}
	}

	return false
}

// adds "adjust" to all the argument locations for the call n.
// n must be a defer or go node that has already been walked.
func (pstate *PackageState) adjustargs(n *Node, adjust int) {
	callfunc := n.Left
	for _, arg := range callfunc.List.Slice() {
		if arg.Op != OAS {
			pstate.Fatalf("call arg not assignment")
		}
		lhs := arg.Left
		if lhs.Op == ONAME {
			// This is a temporary introduced by reorder1.
			// The real store to the stack appears later in the arg list.
			continue
		}

		if lhs.Op != OINDREGSP {
			pstate.Fatalf("call argument store does not use OINDREGSP")
		}

		// can't really check this in machine-indep code.
		//if(lhs->val.u.reg != D_SP)
		//      Fatalf("call arg assign not indreg(SP)")
		lhs.Xoffset += int64(adjust)
	}
}

// The result of walkstmt MUST be assigned back to n, e.g.
// 	n.Left = walkstmt(n.Left)
func (pstate *PackageState) walkstmt(n *Node) *Node {
	if n == nil {
		return n
	}

	pstate.setlineno(n)

	pstate.walkstmtlist(n.Ninit.Slice())

	switch n.Op {
	default:
		if n.Op == ONAME {
			pstate.yyerror("%v is not a top level statement", n.Sym)
		} else {
			pstate.yyerror("%v is not a top level statement", n.Op)
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
			pstate.Fatalf("missing typecheck: %+v", n)
		}
		wascopy := n.Op == OCOPY
		init := n.Ninit
		n.Ninit.Set(nil)
		n = pstate.walkexpr(n, &init)
		n = pstate.addinit(n, init.Slice())
		if wascopy && n.Op == OCONVNOP {
			n.Op = OEMPTY // don't leave plain values as statements.
		}

	// special case for a receive where we throw away
	// the value received.
	case ORECV:
		if n.Typecheck() == 0 {
			pstate.Fatalf("missing typecheck: %+v", n)
		}
		init := n.Ninit
		n.Ninit.Set(nil)

		n.Left = pstate.walkexpr(n.Left, &init)
		n = pstate.mkcall1(pstate.chanfn("chanrecv1", 2, n.Left.Type), nil, &init, n.Left, pstate.nodnil())
		n = pstate.walkexpr(n, &init)

		n = pstate.addinit(n, init.Slice())

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
			if pstate.compiling_runtime {
				pstate.yyerror("%v escapes to heap, not allowed in runtime.", v)
			}
			if pstate.prealloc[v] == nil {
				pstate.prealloc[v] = pstate.callnew(v.Type)
			}
			nn := pstate.nod(OAS, v.Name.Param.Heapaddr, pstate.prealloc[v])
			nn.SetColas(true)
			nn = pstate.typecheck(nn, Etop)
			return pstate.walkstmt(nn)
		}

	case OBLOCK:
		pstate.walkstmtlist(n.List.Slice())

	case OXCASE:
		pstate.yyerror("case statement out of place")
		n.Op = OCASE
		fallthrough

	case OCASE:
		n.Right = pstate.walkstmt(n.Right)

	case ODEFER:
		pstate.Curfn.Func.SetHasDefer(true)
		fallthrough
	case OPROC:
		switch n.Left.Op {
		case OPRINT, OPRINTN:
			n.Left = pstate.wrapCall(n.Left, &n.Ninit)

		case ODELETE:
			if pstate.mapfast(n.Left.List.First().Type) == mapslow {
				n.Left = pstate.wrapCall(n.Left, &n.Ninit)
			} else {
				n.Left = pstate.walkexpr(n.Left, &n.Ninit)
			}

		case OCOPY:
			n.Left = pstate.copyany(n.Left, &n.Ninit, true)

		default:
			n.Left = pstate.walkexpr(n.Left, &n.Ninit)
		}

		// make room for size & fn arguments.
		pstate.adjustargs(n, 2*pstate.Widthptr)

	case OFOR, OFORUNTIL:
		if n.Left != nil {
			pstate.walkstmtlist(n.Left.Ninit.Slice())
			init := n.Left.Ninit
			n.Left.Ninit.Set(nil)
			n.Left = pstate.walkexpr(n.Left, &init)
			n.Left = pstate.addinit(n.Left, init.Slice())
		}

		n.Right = pstate.walkstmt(n.Right)
		if n.Op == OFORUNTIL {
			pstate.walkstmtlist(n.List.Slice())
		}
		pstate.walkstmtlist(n.Nbody.Slice())

	case OIF:
		n.Left = pstate.walkexpr(n.Left, &n.Ninit)
		pstate.walkstmtlist(n.Nbody.Slice())
		pstate.walkstmtlist(n.Rlist.Slice())

	case ORETURN:
		if n.List.Len() == 0 {
			break
		}
		if (pstate.Curfn.Type.FuncType(pstate.types).Outnamed && n.List.Len() > 1) || paramoutheap(pstate.Curfn) {
			// assign to the function out parameters,
			// so that reorder3 can fix up conflicts
			var rl []*Node

			for _, ln := range pstate.Curfn.Func.Dcl {
				cl := ln.Class()
				if cl == PAUTO || cl == PAUTOHEAP {
					break
				}
				if cl == PPARAMOUT {
					if ln.isParamStackCopy() {
						ln = pstate.walkexpr(pstate.typecheck(pstate.nod(OIND, ln.Name.Param.Heapaddr, nil), Erv), nil)
					}
					rl = append(rl, ln)
				}
			}

			if got, want := n.List.Len(), len(rl); got != want {
				// order should have rewritten multi-value function calls
				// with explicit OAS2FUNC nodes.
				pstate.Fatalf("expected %v return arguments, have %v", want, got)
			}

			if samelist(rl, n.List.Slice()) {
				// special return in disguise
				// TODO(josharian, 1.12): is "special return" still relevant?
				// Tests still pass w/o this. See comments on https://go-review.googlesource.com/c/go/+/118318
				pstate.walkexprlist(n.List.Slice(), &n.Ninit)
				n.List.Set(nil)

				break
			}

			// move function calls out, to make reorder3's job easier.
			pstate.walkexprlistsafe(n.List.Slice(), &n.Ninit)

			ll := pstate.ascompatee(n.Op, rl, n.List.Slice(), &n.Ninit)
			n.List.Set(pstate.reorder3(ll))
			break
		}
		pstate.walkexprlist(n.List.Slice(), &n.Ninit)

		ll := pstate.ascompatte(nil, false, pstate.Curfn.Type.Results(pstate.types), n.List.Slice(), 1, &n.Ninit)
		n.List.Set(ll)

	case ORETJMP:
		break

	case OSELECT:
		pstate.walkselect(n)

	case OSWITCH:
		pstate.walkswitch(n)

	case ORANGE:
		n = pstate.walkrange(n)
	}

	if n.Op == ONAME {
		pstate.Fatalf("walkstmt ended up with name: %+v", n)
	}
	return n
}

func (pstate *PackageState) isSmallMakeSlice(n *Node) bool {
	if n.Op != OMAKESLICE {
		return false
	}
	l := n.Left
	r := n.Right
	if r == nil {
		r = l
	}
	t := n.Type

	return pstate.smallintconst(l) && pstate.smallintconst(r) && (t.Elem(pstate.types).Width == 0 || r.Int64(pstate) < (1<<16)/t.Elem(pstate.types).Width)
}

// walk the whole tree of the body of an
// expression or simple statement.
// the types expressions are calculated.
// compile-time constants are evaluated.
// complex side effects like statements are appended to init
func (pstate *PackageState) walkexprlist(s []*Node, init *Nodes) {
	for i := range s {
		s[i] = pstate.walkexpr(s[i], init)
	}
}

func (pstate *PackageState) walkexprlistsafe(s []*Node, init *Nodes) {
	for i, n := range s {
		s[i] = pstate.safeexpr(n, init)
		s[i] = pstate.walkexpr(s[i], init)
	}
}

func (pstate *PackageState) walkexprlistcheap(s []*Node, init *Nodes) {
	for i, n := range s {
		s[i] = pstate.cheapexpr(n, init)
		s[i] = pstate.walkexpr(s[i], init)
	}
}

// convFuncName builds the runtime function name for interface conversion.
// It also reports whether the function expects the data by address.
// Not all names are possible. For example, we never generate convE2E or convE2I.
func (pstate *PackageState) convFuncName(from, to *types.Type) (fnname string, needsaddr bool) {
	tkind := to.Tie(pstate.types)
	switch from.Tie(pstate.types) {
	case 'I':
		switch tkind {
		case 'I':
			return "convI2I", false
		}
	case 'T':
		switch tkind {
		case 'E':
			switch {
			case from.Size(pstate.types) == 2 && from.Align == 2:
				return "convT2E16", false
			case from.Size(pstate.types) == 4 && from.Align == 4 && !pstate.types.Haspointers(from):
				return "convT2E32", false
			case from.Size(pstate.types) == 8 && from.Align == pstate.types.Types[TUINT64].Align && !pstate.types.Haspointers(from):
				return "convT2E64", false
			case from.IsString():
				return "convT2Estring", true
			case from.IsSlice():
				return "convT2Eslice", true
			case !pstate.types.Haspointers(from):
				return "convT2Enoptr", true
			}
			return "convT2E", true
		case 'I':
			switch {
			case from.Size(pstate.types) == 2 && from.Align == 2:
				return "convT2I16", false
			case from.Size(pstate.types) == 4 && from.Align == 4 && !pstate.types.Haspointers(from):
				return "convT2I32", false
			case from.Size(pstate.types) == 8 && from.Align == pstate.types.Types[TUINT64].Align && !pstate.types.Haspointers(from):
				return "convT2I64", false
			case from.IsString():
				return "convT2Istring", true
			case from.IsSlice():
				return "convT2Islice", true
			case !pstate.types.Haspointers(from):
				return "convT2Inoptr", true
			}
			return "convT2I", true
		}
	}
	pstate.Fatalf("unknown conv func %c2%c", from.Tie(pstate.types), to.Tie(pstate.types))
	panic("unreachable")
}

// The result of walkexpr MUST be assigned back to n, e.g.
// 	n.Left = walkexpr(n.Left, init)
func (pstate *PackageState) walkexpr(n *Node, init *Nodes) *Node {
	if n == nil {
		return n
	}

	// Eagerly checkwidth all expressions for the back end.
	if n.Type != nil && !n.Type.WidthCalculated() {
		switch n.Type.Etype {
		case TBLANK, TNIL, TIDEAL:
		default:
			pstate.checkwidth(n.Type)
		}
	}

	if init == &n.Ninit {
		// not okay to use n->ninit when walking n,
		// because we might replace n with some other node
		// and would lose the init list.
		pstate.Fatalf("walkexpr init == &n->ninit")
	}

	if n.Ninit.Len() != 0 {
		pstate.walkstmtlist(n.Ninit.Slice())
		init.AppendNodes(&n.Ninit)
	}

	lno := pstate.setlineno(n)

	if pstate.Debug['w'] > 1 {
		Dump("before walk expr", n)
	}

	if n.Typecheck() != 1 {
		pstate.Fatalf("missed typecheck: %+v", n)
	}

	if n.Type.IsUntyped(pstate.types) {
		pstate.Fatalf("expression has untyped type: %+v", n)
	}

	if n.Op == ONAME && n.Class() == PAUTOHEAP {
		nn := pstate.nod(OIND, n.Name.Param.Heapaddr, nil)
		nn = pstate.typecheck(nn, Erv)
		nn = pstate.walkexpr(nn, init)
		nn.Left.SetNonNil(true)
		return nn
	}

opswitch:
	switch n.Op {
	default:
		Dump("walk", n)
		pstate.Fatalf("walkexpr: switch 1 unknown op %+S", n)

	case ONONAME, OINDREGSP, OEMPTY, OGETG:

	case OTYPE, ONAME, OLITERAL:
	// TODO(mdempsky): Just return n; see discussion on CL 38655.
	// Perhaps refactor to use Node.mayBeShared for these instead.
	// If these return early, make sure to still call
	// stringsym for constant strings.

	case ONOT, OMINUS, OPLUS, OCOM, OREAL, OIMAG, ODOTMETH, ODOTINTER,
		OIND, OSPTR, OITAB, OIDATA, OADDR:
		n.Left = pstate.walkexpr(n.Left, init)

	case OEFACE, OAND, OSUB, OMUL, OLT, OLE, OGE, OGT, OADD, OOR, OXOR:
		n.Left = pstate.walkexpr(n.Left, init)
		n.Right = pstate.walkexpr(n.Right, init)

	case ODOT:
		pstate.usefield(n)
		n.Left = pstate.walkexpr(n.Left, init)

	case ODOTTYPE, ODOTTYPE2:
		n.Left = pstate.walkexpr(n.Left, init)
		// Set up interface type addresses for back end.
		n.Right = pstate.typename(n.Type)
		if n.Op == ODOTTYPE {
			n.Right.Right = pstate.typename(n.Left.Type)
		}
		if !n.Type.IsInterface() && !n.Left.Type.IsEmptyInterface(pstate.types) {
			n.List.Set1(pstate.itabname(n.Type, n.Left.Type))
		}

	case ODOTPTR:
		pstate.usefield(n)
		if n.Op == ODOTPTR && n.Left.Type.Elem(pstate.types).Width == 0 {
			// No actual copy will be generated, so emit an explicit nil check.
			n.Left = pstate.cheapexpr(n.Left, init)

			pstate.checknil(n.Left, init)
		}

		n.Left = pstate.walkexpr(n.Left, init)

	case OLEN, OCAP:
		if pstate.isRuneCount(n) {
			// Replace len([]rune(string)) with runtime.countrunes(string).
			n = pstate.mkcall("countrunes", n.Type, init, pstate.conv(n.Left.Left, pstate.types.Types[TSTRING]))
			break
		}

		n.Left = pstate.walkexpr(n.Left, init)

		// replace len(*[10]int) with 10.
		// delayed until now to preserve side effects.
		t := n.Left.Type

		if t.IsPtr() {
			t = t.Elem(pstate.types)
		}
		if t.IsArray() {
			pstate.safeexpr(n.Left, init)
			pstate.setintconst(n, t.NumElem(pstate.types))
			n.SetTypecheck(1)
		}

	case OLSH, ORSH:
		n.Left = pstate.walkexpr(n.Left, init)
		n.Right = pstate.walkexpr(n.Right, init)
		t := n.Left.Type
		n.SetBounded(pstate.bounded(n.Right, 8*t.Width))
		if pstate.Debug['m'] != 0 && n.Bounded() && !pstate.Isconst(n.Right, CTINT) {
			pstate.Warn("shift bounds check elided")
		}

	case OCOMPLEX:
		// Use results from call expression as arguments for complex.
		if n.Left == nil && n.Right == nil {
			n.Left = n.List.First()
			n.Right = n.List.Second()
		}
		n.Left = pstate.walkexpr(n.Left, init)
		n.Right = pstate.walkexpr(n.Right, init)

	case OEQ, ONE:
		n.Left = pstate.walkexpr(n.Left, init)
		n.Right = pstate.walkexpr(n.Right, init)

		// Disable safemode while compiling this code: the code we
		// generate internally can refer to unsafe.Pointer.
		// In this case it can happen if we need to generate an ==
		// for a struct containing a reflect.Value, which itself has
		// an unexported field of type unsafe.Pointer.
		old_safemode := pstate.safemode
		pstate.safemode = false
		n = pstate.walkcompare(n, init)
		pstate.safemode = old_safemode

	case OANDAND, OOROR:
		n.Left = pstate.walkexpr(n.Left, init)

		// cannot put side effects from n.Right on init,
		// because they cannot run before n.Left is checked.
		// save elsewhere and store on the eventual n.Right.
		var ll Nodes

		n.Right = pstate.walkexpr(n.Right, &ll)
		n.Right = pstate.addinit(n.Right, ll.Slice())
		n = pstate.walkinrange(n, init)

	case OPRINT, OPRINTN:
		pstate.walkexprlist(n.List.Slice(), init)
		n = pstate.walkprint(n, init)

	case OPANIC:
		n = pstate.mkcall("gopanic", nil, init, n.Left)

	case ORECOVER:
		n = pstate.mkcall("gorecover", n.Type, init, pstate.nod(OADDR, pstate.nodfp, nil))

	case OCLOSUREVAR, OCFUNC:
		n.SetAddable(true)

	case OCALLINTER:
		pstate.usemethod(n)
		t := n.Left.Type
		if n.List.Len() != 0 && n.List.First().Op == OAS {
			break
		}
		n.Left = pstate.walkexpr(n.Left, init)
		pstate.walkexprlist(n.List.Slice(), init)
		ll := pstate.ascompatte(n, n.Isddd(), t.Params(pstate.types), n.List.Slice(), 0, init)
		n.List.Set(pstate.reorder1(ll))

	case OCALLFUNC:
		if n.Left.Op == OCLOSURE {
			// Transform direct call of a closure to call of a normal function.
			// transformclosure already did all preparation work.

			// Prepend captured variables to argument list.
			n.List.Prepend(n.Left.Func.Enter.Slice()...)

			n.Left.Func.Enter.Set(nil)

			// Replace OCLOSURE with ONAME/PFUNC.
			n.Left = n.Left.Func.Closure.Func.Nname

			// Update type of OCALLFUNC node.
			// Output arguments had not changed, but their offsets could.
			if n.Left.Type.NumResults(pstate.types) == 1 {
				n.Type = n.Left.Type.Results(pstate.types).Field(pstate.types, 0).Type
			} else {
				n.Type = n.Left.Type.Results(pstate.types)
			}
		}

		t := n.Left.Type
		if n.List.Len() != 0 && n.List.First().Op == OAS {
			break
		}

		n.Left = pstate.walkexpr(n.Left, init)
		pstate.walkexprlist(n.List.Slice(), init)

		ll := pstate.ascompatte(n, n.Isddd(), t.Params(pstate.types), n.List.Slice(), 0, init)
		n.List.Set(pstate.reorder1(ll))

	case OCALLMETH:
		t := n.Left.Type
		if n.List.Len() != 0 && n.List.First().Op == OAS {
			break
		}
		n.Left = pstate.walkexpr(n.Left, init)
		pstate.walkexprlist(n.List.Slice(), init)
		ll := pstate.ascompatte(n, false, t.Recvs(pstate.types), []*Node{n.Left.Left}, 0, init)
		lr := pstate.ascompatte(n, n.Isddd(), t.Params(pstate.types), n.List.Slice(), 0, init)
		ll = append(ll, lr...)
		n.Left.Left = nil
		pstate.updateHasCall(n.Left)
		n.List.Set(pstate.reorder1(ll))

	case OAS, OASOP:
		init.AppendNodes(&n.Ninit)

		// Recognize m[k] = append(m[k], ...) so we can reuse
		// the mapassign call.
		mapAppend := n.Left.Op == OINDEXMAP && n.Right.Op == OAPPEND
		if mapAppend && !pstate.samesafeexpr(n.Left, n.Right.List.First()) {
			pstate.Fatalf("not same expressions: %v != %v", n.Left, n.Right.List.First())
		}

		n.Left = pstate.walkexpr(n.Left, init)
		n.Left = pstate.safeexpr(n.Left, init)

		if mapAppend {
			n.Right.List.SetFirst(n.Left)
		}

		if n.Op == OASOP {
			// Rewrite x op= y into x = x op y.
			n.Right = pstate.nod(n.SubOp(pstate), n.Left, n.Right)
			n.Right = pstate.typecheck(n.Right, Erv)

			n.Op = OAS
			n.ResetAux()
		}

		if pstate.oaslit(n, init) {
			break
		}

		if n.Right == nil {
			// TODO(austin): Check all "implicit zeroing"
			break
		}

		if !pstate.instrumenting && pstate.isZero(n.Right) {
			break
		}

		switch n.Right.Op {
		default:
			n.Right = pstate.walkexpr(n.Right, init)

		case ORECV:
			// x = <-c; n.Left is x, n.Right.Left is c.
			// orderstmt made sure x is addressable.
			n.Right.Left = pstate.walkexpr(n.Right.Left, init)

			n1 := pstate.nod(OADDR, n.Left, nil)
			r := n.Right.Left // the channel
			n = pstate.mkcall1(pstate.chanfn("chanrecv1", 2, r.Type), nil, init, r, n1)
			n = pstate.walkexpr(n, init)
			break opswitch

		case OAPPEND:
			// x = append(...)
			r := n.Right
			if r.Type.Elem(pstate.types).NotInHeap() {
				pstate.yyerror("%v is go:notinheap; heap allocation disallowed", r.Type.Elem(pstate.types))
			}
			switch {
			case pstate.isAppendOfMake(r):
				// x = append(y, make([]T, y)...)
				r = pstate.extendslice(r, init)
			case r.Isddd():
				r = pstate.appendslice(r, init) // also works for append(slice, string).
			default:
				r = pstate.walkappend(r, init, n)
			}
			n.Right = r
			if r.Op == OAPPEND {
				// Left in place for back end.
				// Do not add a new write barrier.
				// Set up address of type for back end.
				r.Left = pstate.typename(r.Type.Elem(pstate.types))
				break opswitch
			}
			// Otherwise, lowered for race detector.
			// Treat as ordinary assignment.
		}

		if n.Left != nil && n.Right != nil {
			n = pstate.convas(n, init)
		}

	case OAS2:
		init.AppendNodes(&n.Ninit)
		pstate.walkexprlistsafe(n.List.Slice(), init)
		pstate.walkexprlistsafe(n.Rlist.Slice(), init)
		ll := pstate.ascompatee(OAS, n.List.Slice(), n.Rlist.Slice(), init)
		ll = pstate.reorder3(ll)
		n = pstate.liststmt(ll)

	// a,b,... = fn()
	case OAS2FUNC:
		init.AppendNodes(&n.Ninit)

		r := n.Rlist.First()
		pstate.walkexprlistsafe(n.List.Slice(), init)
		r = pstate.walkexpr(r, init)

		if pstate.isIntrinsicCall(r) {
			n.Rlist.Set1(r)
			break
		}
		init.Append(r)

		ll := pstate.ascompatet(n.List, r.Type)
		n = pstate.liststmt(ll)

	// x, y = <-c
	// orderstmt made sure x is addressable.
	case OAS2RECV:
		init.AppendNodes(&n.Ninit)

		r := n.Rlist.First()
		pstate.walkexprlistsafe(n.List.Slice(), init)
		r.Left = pstate.walkexpr(r.Left, init)
		var n1 *Node
		if n.List.First().isBlank() {
			n1 = pstate.nodnil()
		} else {
			n1 = pstate.nod(OADDR, n.List.First(), nil)
		}
		fn := pstate.chanfn("chanrecv2", 2, r.Left.Type)
		ok := n.List.Second()
		call := pstate.mkcall1(fn, ok.Type, init, r.Left, n1)
		n = pstate.nod(OAS, ok, call)
		n = pstate.typecheck(n, Etop)

	// a,b = m[i]
	case OAS2MAPR:
		init.AppendNodes(&n.Ninit)

		r := n.Rlist.First()
		pstate.walkexprlistsafe(n.List.Slice(), init)
		r.Left = pstate.walkexpr(r.Left, init)
		r.Right = pstate.walkexpr(r.Right, init)
		t := r.Left.Type

		fast := pstate.mapfast(t)
		var key *Node
		if fast != mapslow {
			// fast versions take key by value
			key = r.Right
		} else {
			// standard version takes key by reference
			// orderexpr made sure key is addressable.
			key = pstate.nod(OADDR, r.Right, nil)
		}

		// from:
		//   a,b = m[i]
		// to:
		//   var,b = mapaccess2*(t, m, i)
		//   a = *var
		a := n.List.First()

		if w := t.Elem(pstate.types).Width; w <= 1024 { // 1024 must match runtime/map.go:maxZero
			fn := pstate.mapfn(pstate.mapaccess2[fast], t)
			r = pstate.mkcall1(fn, fn.Type.Results(pstate.types), init, pstate.typename(t), r.Left, key)
		} else {
			fn := pstate.mapfn("mapaccess2_fat", t)
			z := pstate.zeroaddr(w)
			r = pstate.mkcall1(fn, fn.Type.Results(pstate.types), init, pstate.typename(t), r.Left, key, z)
		}

		// mapaccess2* returns a typed bool, but due to spec changes,
		// the boolean result of i.(T) is now untyped so we make it the
		// same type as the variable on the lhs.
		if ok := n.List.Second(); !ok.isBlank() && ok.Type.IsBoolean() {
			r.Type.Field(pstate.types, 1).Type = ok.Type
		}
		n.Rlist.Set1(r)
		n.Op = OAS2FUNC

		// don't generate a = *var if a is _
		if !a.isBlank() {
			var_ := pstate.temp(pstate.types.NewPtr(t.Elem(pstate.types)))
			var_.SetTypecheck(1)
			var_.SetNonNil(true) // mapaccess always returns a non-nil pointer
			n.List.SetFirst(var_)
			n = pstate.walkexpr(n, init)
			init.Append(n)
			n = pstate.nod(OAS, a, pstate.nod(OIND, var_, nil))
		}

		n = pstate.typecheck(n, Etop)
		n = pstate.walkexpr(n, init)

	case ODELETE:
		init.AppendNodes(&n.Ninit)
		map_ := n.List.First()
		key := n.List.Second()
		map_ = pstate.walkexpr(map_, init)
		key = pstate.walkexpr(key, init)

		t := map_.Type
		fast := pstate.mapfast(t)
		if fast == mapslow {
			// orderstmt made sure key is addressable.
			key = pstate.nod(OADDR, key, nil)
		}
		n = pstate.mkcall1(pstate.mapfndel(pstate.mapdelete[fast], t), nil, init, pstate.typename(t), map_, key)

	case OAS2DOTTYPE:
		pstate.walkexprlistsafe(n.List.Slice(), init)
		n.Rlist.SetFirst(pstate.walkexpr(n.Rlist.First(), init))

	case OCONVIFACE:
		n.Left = pstate.walkexpr(n.Left, init)

		// Optimize convT2E or convT2I as a two-word copy when T is pointer-shaped.
		if pstate.isdirectiface(n.Left.Type) {
			var t *Node
			if n.Type.IsEmptyInterface(pstate.types) {
				t = pstate.typename(n.Left.Type)
			} else {
				t = pstate.itabname(n.Left.Type, n.Type)
			}
			l := pstate.nod(OEFACE, t, n.Left)
			l.Type = n.Type
			l.SetTypecheck(n.Typecheck())
			n = l
			break
		}

		if pstate.staticbytes == nil {
			pstate.staticbytes = pstate.newname(pstate.Runtimepkg.Lookup(pstate.types, "staticbytes"))
			pstate.staticbytes.SetClass(PEXTERN)
			pstate.staticbytes.Type = pstate.types.NewArray(pstate.types.Types[TUINT8], 256)
			pstate.zerobase = pstate.newname(pstate.Runtimepkg.Lookup(pstate.types, "zerobase"))
			pstate.zerobase.SetClass(PEXTERN)
			pstate.zerobase.Type = pstate.types.Types[TUINTPTR]
		}

		// Optimize convT2{E,I} for many cases in which T is not pointer-shaped,
		// by using an existing addressable value identical to n.Left
		// or creating one on the stack.
		var value *Node
		switch {
		case n.Left.Type.Size(pstate.types) == 0:
			// n.Left is zero-sized. Use zerobase.
			pstate.cheapexpr(n.Left, init) // Evaluate n.Left for side-effects. See issue 19246.
			value = pstate.zerobase
		case n.Left.Type.IsBoolean() || (n.Left.Type.Size(pstate.types) == 1 && n.Left.Type.IsInteger()):
			// n.Left is a bool/byte. Use staticbytes[n.Left].
			n.Left = pstate.cheapexpr(n.Left, init)
			value = pstate.nod(OINDEX, pstate.staticbytes, pstate.byteindex(n.Left))
			value.SetBounded(true)
		case n.Left.Class() == PEXTERN && n.Left.Name != nil && n.Left.Name.Readonly():
			// n.Left is a readonly global; use it directly.
			value = n.Left
		case !n.Left.Type.IsInterface() && n.Esc == EscNone && n.Left.Type.Width <= 1024:
			// n.Left does not escape. Use a stack temporary initialized to n.Left.
			value = pstate.temp(n.Left.Type)
			init.Append(pstate.typecheck(pstate.nod(OAS, value, n.Left), Etop))
		}

		if value != nil {
			// Value is identical to n.Left.
			// Construct the interface directly: {type/itab, &value}.
			var t *Node
			if n.Type.IsEmptyInterface(pstate.types) {
				t = pstate.typename(n.Left.Type)
			} else {
				t = pstate.itabname(n.Left.Type, n.Type)
			}
			l := pstate.nod(OEFACE, t, pstate.typecheck(pstate.nod(OADDR, value, nil), Erv))
			l.Type = n.Type
			l.SetTypecheck(n.Typecheck())
			n = l
			break
		}

		// Implement interface to empty interface conversion.
		// tmp = i.itab
		// if tmp != nil {
		//    tmp = tmp.type
		// }
		// e = iface{tmp, i.data}
		if n.Type.IsEmptyInterface(pstate.types) && n.Left.Type.IsInterface() && !n.Left.Type.IsEmptyInterface(pstate.types) {
			// Evaluate the input interface.
			c := pstate.temp(n.Left.Type)
			init.Append(pstate.nod(OAS, c, n.Left))

			// Get the itab out of the interface.
			tmp := pstate.temp(pstate.types.NewPtr(pstate.types.Types[TUINT8]))
			init.Append(pstate.nod(OAS, tmp, pstate.typecheck(pstate.nod(OITAB, c, nil), Erv)))

			// Get the type out of the itab.
			nif := pstate.nod(OIF, pstate.typecheck(pstate.nod(ONE, tmp, pstate.nodnil()), Erv), nil)
			nif.Nbody.Set1(pstate.nod(OAS, tmp, pstate.itabType(tmp)))
			init.Append(nif)

			// Build the result.
			e := pstate.nod(OEFACE, tmp, pstate.ifaceData(c, pstate.types.NewPtr(pstate.types.Types[TUINT8])))
			e.Type = n.Type // assign type manually, typecheck doesn't understand OEFACE.
			e.SetTypecheck(1)
			n = e
			break
		}

		var ll []*Node
		if n.Type.IsEmptyInterface(pstate.types) {
			if !n.Left.Type.IsInterface() {
				ll = append(ll, pstate.typename(n.Left.Type))
			}
		} else {
			if n.Left.Type.IsInterface() {
				ll = append(ll, pstate.typename(n.Type))
			} else {
				ll = append(ll, pstate.itabname(n.Left.Type, n.Type))
			}
		}

		fnname, needsaddr := pstate.convFuncName(n.Left.Type, n.Type)
		v := n.Left
		if needsaddr {
			// Types of large or unknown size are passed by reference.
			// Orderexpr arranged for n.Left to be a temporary for all
			// the conversions it could see. Comparison of an interface
			// with a non-interface, especially in a switch on interface value
			// with non-interface cases, is not visible to orderstmt, so we
			// have to fall back on allocating a temp here.
			if !islvalue(v) {
				v = pstate.copyexpr(v, v.Type, init)
			}
			v = pstate.nod(OADDR, v, nil)
		}
		ll = append(ll, v)

		pstate.dowidth(n.Left.Type)
		fn := pstate.syslook(fnname)
		fn = pstate.substArgTypes(fn, n.Left.Type, n.Type)
		pstate.dowidth(fn.Type)
		n = pstate.nod(OCALL, fn, nil)
		n.List.Set(ll)
		n = pstate.typecheck(n, Erv)
		n = pstate.walkexpr(n, init)

	case OCONV, OCONVNOP:
		if pstate.thearch.SoftFloat {
			// For the soft-float case, ssa.go handles these conversions.
			n.Left = pstate.walkexpr(n.Left, init)
			break
		}
		switch pstate.thearch.LinkArch.Family {
		case sys.ARM, sys.MIPS:
			if n.Left.Type.IsFloat() {
				switch n.Type.Etype {
				case TINT64:
					n = pstate.mkcall("float64toint64", n.Type, init, pstate.conv(n.Left, pstate.types.Types[TFLOAT64]))
					break opswitch
				case TUINT64:
					n = pstate.mkcall("float64touint64", n.Type, init, pstate.conv(n.Left, pstate.types.Types[TFLOAT64]))
					break opswitch
				}
			}

			if n.Type.IsFloat() {
				switch n.Left.Type.Etype {
				case TINT64:
					n = pstate.conv(pstate.mkcall("int64tofloat64", pstate.types.Types[TFLOAT64], init, pstate.conv(n.Left, pstate.types.Types[TINT64])), n.Type)
					break opswitch
				case TUINT64:
					n = pstate.conv(pstate.mkcall("uint64tofloat64", pstate.types.Types[TFLOAT64], init, pstate.conv(n.Left, pstate.types.Types[TUINT64])), n.Type)
					break opswitch
				}
			}

		case sys.I386:
			if n.Left.Type.IsFloat() {
				switch n.Type.Etype {
				case TINT64:
					n = pstate.mkcall("float64toint64", n.Type, init, pstate.conv(n.Left, pstate.types.Types[TFLOAT64]))
					break opswitch
				case TUINT64:
					n = pstate.mkcall("float64touint64", n.Type, init, pstate.conv(n.Left, pstate.types.Types[TFLOAT64]))
					break opswitch
				case TUINT32, TUINT, TUINTPTR:
					n = pstate.mkcall("float64touint32", n.Type, init, pstate.conv(n.Left, pstate.types.Types[TFLOAT64]))
					break opswitch
				}
			}
			if n.Type.IsFloat() {
				switch n.Left.Type.Etype {
				case TINT64:
					n = pstate.conv(pstate.mkcall("int64tofloat64", pstate.types.Types[TFLOAT64], init, pstate.conv(n.Left, pstate.types.Types[TINT64])), n.Type)
					break opswitch
				case TUINT64:
					n = pstate.conv(pstate.mkcall("uint64tofloat64", pstate.types.Types[TFLOAT64], init, pstate.conv(n.Left, pstate.types.Types[TUINT64])), n.Type)
					break opswitch
				case TUINT32, TUINT, TUINTPTR:
					n = pstate.conv(pstate.mkcall("uint32tofloat64", pstate.types.Types[TFLOAT64], init, pstate.conv(n.Left, pstate.types.Types[TUINT32])), n.Type)
					break opswitch
				}
			}
		}
		n.Left = pstate.walkexpr(n.Left, init)

	case OANDNOT:
		n.Left = pstate.walkexpr(n.Left, init)
		n.Op = OAND
		n.Right = pstate.nod(OCOM, n.Right, nil)
		n.Right = pstate.typecheck(n.Right, Erv)
		n.Right = pstate.walkexpr(n.Right, init)

	case ODIV, OMOD:
		n.Left = pstate.walkexpr(n.Left, init)
		n.Right = pstate.walkexpr(n.Right, init)

		// rewrite complex div into function call.
		et := n.Left.Type.Etype

		if pstate.isComplex[et] && n.Op == ODIV {
			t := n.Type
			n = pstate.mkcall("complex128div", pstate.types.Types[TCOMPLEX128], init, pstate.conv(n.Left, pstate.types.Types[TCOMPLEX128]), pstate.conv(n.Right, pstate.types.Types[TCOMPLEX128]))
			n = pstate.conv(n, t)
			break
		}

		// Nothing to do for float divisions.
		if pstate.isFloat[et] {
			break
		}

		// rewrite 64-bit div and mod on 32-bit architectures.
		// TODO: Remove this code once we can introduce
		// runtime calls late in SSA processing.
		if pstate.Widthreg < 8 && (et == TINT64 || et == TUINT64) {
			if n.Right.Op == OLITERAL {
				// Leave div/mod by constant powers of 2.
				// The SSA backend will handle those.
				switch et {
				case TINT64:
					c := n.Right.Int64(pstate)
					if c < 0 {
						c = -c
					}
					if c != 0 && c&(c-1) == 0 {
						break opswitch
					}
				case TUINT64:
					c := uint64(n.Right.Int64(pstate))
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
			n = pstate.mkcall(fn, n.Type, init, pstate.conv(n.Left, pstate.types.Types[et]), pstate.conv(n.Right, pstate.types.Types[et]))
		}

	case OINDEX:
		n.Left = pstate.walkexpr(n.Left, init)

		// save the original node for bounds checking elision.
		// If it was a ODIV/OMOD walk might rewrite it.
		r := n.Right

		n.Right = pstate.walkexpr(n.Right, init)

		// if range of type cannot exceed static array bound,
		// disable bounds check.
		if n.Bounded() {
			break
		}
		t := n.Left.Type
		if t != nil && t.IsPtr() {
			t = t.Elem(pstate.types)
		}
		if t.IsArray() {
			n.SetBounded(pstate.bounded(r, t.NumElem(pstate.types)))
			if pstate.Debug['m'] != 0 && n.Bounded() && !pstate.Isconst(n.Right, CTINT) {
				pstate.Warn("index bounds check elided")
			}
			if pstate.smallintconst(n.Right) && !n.Bounded() {
				pstate.yyerror("index out of bounds")
			}
		} else if pstate.Isconst(n.Left, CTSTR) {
			n.SetBounded(pstate.bounded(r, int64(len(n.Left.Val().U.(string)))))
			if pstate.Debug['m'] != 0 && n.Bounded() && !pstate.Isconst(n.Right, CTINT) {
				pstate.Warn("index bounds check elided")
			}
			if pstate.smallintconst(n.Right) && !n.Bounded() {
				pstate.yyerror("index out of bounds")
			}
		}

		if pstate.Isconst(n.Right, CTINT) {
			if n.Right.Val().U.(*Mpint).CmpInt64(0) < 0 || n.Right.Val().U.(*Mpint).Cmp(pstate.maxintval[TINT]) > 0 {
				pstate.yyerror("index out of bounds")
			}
		}

	case OINDEXMAP:
		// Replace m[k] with *map{access1,assign}(maptype, m, &k)
		n.Left = pstate.walkexpr(n.Left, init)
		n.Right = pstate.walkexpr(n.Right, init)
		map_ := n.Left
		key := n.Right
		t := map_.Type
		if n.IndexMapLValue(pstate) {
			// This m[k] expression is on the left-hand side of an assignment.
			fast := pstate.mapfast(t)
			if fast == mapslow {
				// standard version takes key by reference.
				// orderexpr made sure key is addressable.
				key = pstate.nod(OADDR, key, nil)
			}
			n = pstate.mkcall1(pstate.mapfn(pstate.mapassign[fast], t), nil, init, pstate.typename(t), map_, key)
		} else {
			// m[k] is not the target of an assignment.
			fast := pstate.mapfast(t)
			if fast == mapslow {
				// standard version takes key by reference.
				// orderexpr made sure key is addressable.
				key = pstate.nod(OADDR, key, nil)
			}

			if w := t.Elem(pstate.types).Width; w <= 1024 { // 1024 must match runtime/map.go:maxZero
				n = pstate.mkcall1(pstate.mapfn(pstate.mapaccess1[fast], t), pstate.types.NewPtr(t.Elem(pstate.types)), init, pstate.typename(t), map_, key)
			} else {
				z := pstate.zeroaddr(w)
				n = pstate.mkcall1(pstate.mapfn("mapaccess1_fat", t), pstate.types.NewPtr(t.Elem(pstate.types)), init, pstate.typename(t), map_, key, z)
			}
		}
		n.Type = pstate.types.NewPtr(t.Elem(pstate.types))
		n.SetNonNil(true) // mapaccess1* and mapassign always return non-nil pointers.
		n = pstate.nod(OIND, n, nil)
		n.Type = t.Elem(pstate.types)
		n.SetTypecheck(1)

	case ORECV:
		pstate.Fatalf("walkexpr ORECV") // should see inside OAS only

	case OSLICE, OSLICEARR, OSLICESTR, OSLICE3, OSLICE3ARR:
		n.Left = pstate.walkexpr(n.Left, init)
		low, high, max := n.SliceBounds(pstate)
		low = pstate.walkexpr(low, init)
		if low != nil && pstate.isZero(low) {
			// Reduce x[0:j] to x[:j] and x[0:j:k] to x[:j:k].
			low = nil
		}
		high = pstate.walkexpr(high, init)
		max = pstate.walkexpr(max, init)
		n.SetSliceBounds(pstate, low, high, max)
		if n.Op.IsSlice3(pstate) {
			if max != nil && max.Op == OCAP && pstate.samesafeexpr(n.Left, max.Left) {
				// Reduce x[i:j:cap(x)] to x[i:j].
				if n.Op == OSLICE3 {
					n.Op = OSLICE
				} else {
					n.Op = OSLICEARR
				}
				n = pstate.reduceSlice(n)
			}
		} else {
			n = pstate.reduceSlice(n)
		}

	case ONEW:
		if n.Esc == EscNone {
			if n.Type.Elem(pstate.types).Width >= 1<<16 {
				pstate.Fatalf("large ONEW with EscNone: %v", n)
			}
			r := pstate.temp(n.Type.Elem(pstate.types))
			r = pstate.nod(OAS, r, nil) // zero temp
			r = pstate.typecheck(r, Etop)
			init.Append(r)
			r = pstate.nod(OADDR, r.Left, nil)
			r = pstate.typecheck(r, Erv)
			n = r
		} else {
			n = pstate.callnew(n.Type.Elem(pstate.types))
		}

	case OCMPSTR:
		// s + "badgerbadgerbadger" == "badgerbadgerbadger"
		if (n.SubOp(pstate) == OEQ || n.SubOp(pstate) == ONE) && pstate.Isconst(n.Right, CTSTR) && n.Left.Op == OADDSTR && n.Left.List.Len() == 2 && pstate.Isconst(n.Left.List.Second(), CTSTR) && strlit(n.Right) == strlit(n.Left.List.Second()) {
			r := pstate.nod(n.SubOp(pstate), pstate.nod(OLEN, n.Left.List.First(), nil), pstate.nodintconst(0))
			n = pstate.finishcompare(n, r, init)
			break
		}

		// Rewrite comparisons to short constant strings as length+byte-wise comparisons.
		var cs, ncs *Node // const string, non-const string
		switch {
		case pstate.Isconst(n.Left, CTSTR) && pstate.Isconst(n.Right, CTSTR):
		// ignore; will be constant evaluated
		case pstate.Isconst(n.Left, CTSTR):
			cs = n.Left
			ncs = n.Right
		case pstate.Isconst(n.Right, CTSTR):
			cs = n.Right
			ncs = n.Left
		}
		if cs != nil {
			cmp := n.SubOp(pstate)
			// Our comparison below assumes that the non-constant string
			// is on the left hand side, so rewrite "" cmp x to x cmp "".
			// See issue 24817.
			if pstate.Isconst(n.Left, CTSTR) {
				cmp = pstate.brrev(cmp)
			}

			// maxRewriteLen was chosen empirically.
			// It is the value that minimizes cmd/go file size
			// across most architectures.
			// See the commit description for CL 26758 for details.
			maxRewriteLen := 6
			// Some architectures can load unaligned byte sequence as 1 word.
			// So we can cover longer strings with the same amount of code.
			canCombineLoads := pstate.canMergeLoads()
			combine64bit := false
			if canCombineLoads {
				// Keep this low enough to generate less code than a function call.
				maxRewriteLen = 2 * pstate.thearch.LinkArch.RegSize
				combine64bit = pstate.thearch.LinkArch.RegSize >= 8
			}

			var and Op
			switch cmp {
			case OEQ:
				and = OANDAND
			case ONE:
				and = OOROR
			default:
				// Don't do byte-wise comparisons for <, <=, etc.
				// They're fairly complicated.
				// Length-only checks are ok, though.
				maxRewriteLen = 0
			}
			if s := cs.Val().U.(string); len(s) <= maxRewriteLen {
				if len(s) > 0 {
					ncs = pstate.safeexpr(ncs, init)
				}
				r := pstate.nod(cmp, pstate.nod(OLEN, ncs, nil), pstate.nodintconst(int64(len(s))))
				remains := len(s)
				for i := 0; remains > 0; {
					if remains == 1 || !canCombineLoads {
						cb := pstate.nodintconst(int64(s[i]))
						ncb := pstate.nod(OINDEX, ncs, pstate.nodintconst(int64(i)))
						r = pstate.nod(and, r, pstate.nod(cmp, ncb, cb))
						remains--
						i++
						continue
					}
					var step int
					var convType *types.Type
					switch {
					case remains >= 8 && combine64bit:
						convType = pstate.types.Types[TINT64]
						step = 8
					case remains >= 4:
						convType = pstate.types.Types[TUINT32]
						step = 4
					case remains >= 2:
						convType = pstate.types.Types[TUINT16]
						step = 2
					}
					ncsubstr := pstate.nod(OINDEX, ncs, pstate.nodintconst(int64(i)))
					ncsubstr = pstate.conv(ncsubstr, convType)
					csubstr := int64(s[i])
					// Calculate large constant from bytes as sequence of shifts and ors.
					// Like this:  uint32(s[0]) | uint32(s[1])<<8 | uint32(s[2])<<16 ...
					// ssa will combine this into a single large load.
					for offset := 1; offset < step; offset++ {
						b := pstate.nod(OINDEX, ncs, pstate.nodintconst(int64(i+offset)))
						b = pstate.conv(b, convType)
						b = pstate.nod(OLSH, b, pstate.nodintconst(int64(8*offset)))
						ncsubstr = pstate.nod(OOR, ncsubstr, b)
						csubstr = csubstr | int64(s[i+offset])<<uint8(8*offset)
					}
					csubstrPart := pstate.nodintconst(csubstr)
					// Compare "step" bytes as once
					r = pstate.nod(and, r, pstate.nod(cmp, csubstrPart, ncsubstr))
					remains -= step
					i += step
				}
				n = pstate.finishcompare(n, r, init)
				break
			}
		}

		var r *Node
		if n.SubOp(pstate) == OEQ || n.SubOp(pstate) == ONE {
			// prepare for rewrite below
			n.Left = pstate.cheapexpr(n.Left, init)
			n.Right = pstate.cheapexpr(n.Right, init)

			lstr := pstate.conv(n.Left, pstate.types.Types[TSTRING])
			rstr := pstate.conv(n.Right, pstate.types.Types[TSTRING])
			lptr := pstate.nod(OSPTR, lstr, nil)
			rptr := pstate.nod(OSPTR, rstr, nil)
			llen := pstate.conv(pstate.nod(OLEN, lstr, nil), pstate.types.Types[TUINTPTR])
			rlen := pstate.conv(pstate.nod(OLEN, rstr, nil), pstate.types.Types[TUINTPTR])

			fn := pstate.syslook("memequal")
			fn = pstate.substArgTypes(fn, pstate.types.Types[TUINT8], pstate.types.Types[TUINT8])
			r = pstate.mkcall1(fn, pstate.types.Types[TBOOL], init, lptr, rptr, llen)

			// quick check of len before full compare for == or !=.
			// memequal then tests equality up to length len.
			if n.SubOp(pstate) == OEQ {
				// len(left) == len(right) && memequal(left, right, len)
				r = pstate.nod(OANDAND, pstate.nod(OEQ, llen, rlen), r)
			} else {
				// len(left) != len(right) || !memequal(left, right, len)
				r = pstate.nod(ONOT, r, nil)
				r = pstate.nod(OOROR, pstate.nod(ONE, llen, rlen), r)
			}
		} else {
			// sys_cmpstring(s1, s2) :: 0
			r = pstate.mkcall("cmpstring", pstate.types.Types[TINT], init, pstate.conv(n.Left, pstate.types.Types[TSTRING]), pstate.conv(n.Right, pstate.types.Types[TSTRING]))
			r = pstate.nod(n.SubOp(pstate), r, pstate.nodintconst(0))
		}

		n = pstate.finishcompare(n, r, init)

	case OADDSTR:
		n = pstate.addstr(n, init)

	case OAPPEND:
		// order should make sure we only see OAS(node, OAPPEND), which we handle above.
		pstate.Fatalf("append outside assignment")

	case OCOPY:
		n = pstate.copyany(n, init, pstate.instrumenting && !pstate.compiling_runtime)

	// cannot use chanfn - closechan takes any, not chan any
	case OCLOSE:
		fn := pstate.syslook("closechan")

		fn = pstate.substArgTypes(fn, n.Left.Type)
		n = pstate.mkcall1(fn, nil, init, n.Left)

	case OMAKECHAN:
		// When size fits into int, use makechan instead of
		// makechan64, which is faster and shorter on 32 bit platforms.
		size := n.Left
		fnname := "makechan64"
		argtype := pstate.types.Types[TINT64]

		// Type checking guarantees that TIDEAL size is positive and fits in an int.
		// The case of size overflow when converting TUINT or TUINTPTR to TINT
		// will be handled by the negative range checks in makechan during runtime.
		if size.Type.IsKind(TIDEAL) || pstate.maxintval[size.Type.Etype].Cmp(pstate.maxintval[TUINT]) <= 0 {
			fnname = "makechan"
			argtype = pstate.types.Types[TINT]
		}

		n = pstate.mkcall1(pstate.chanfn(fnname, 1, n.Type), n.Type, init, pstate.typename(n.Type), pstate.conv(size, argtype))

	case OMAKEMAP:
		t := n.Type
		hmapType := pstate.hmap(t)
		hint := n.Left

		// var h *hmap
		var h *Node
		if n.Esc == EscNone {
			// Allocate hmap on stack.

			// var hv hmap
			hv := pstate.temp(hmapType)
			zero := pstate.nod(OAS, hv, nil)
			zero = pstate.typecheck(zero, Etop)
			init.Append(zero)
			// h = &hv
			h = pstate.nod(OADDR, hv, nil)

			// Allocate one bucket pointed to by hmap.buckets on stack if hint
			// is not larger than BUCKETSIZE. In case hint is larger than
			// BUCKETSIZE runtime.makemap will allocate the buckets on the heap.
			// Maximum key and value size is 128 bytes, larger objects
			// are stored with an indirection. So max bucket size is 2048+eps.
			if !pstate.Isconst(hint, CTINT) ||
				!(hint.Val().U.(*Mpint).CmpInt64(BUCKETSIZE) > 0) {
				// var bv bmap
				bv := pstate.temp(pstate.bmap(t))

				zero = pstate.nod(OAS, bv, nil)
				zero = pstate.typecheck(zero, Etop)
				init.Append(zero)

				// b = &bv
				b := pstate.nod(OADDR, bv, nil)

				// h.buckets = b
				bsym := hmapType.Field(pstate.types, 5).Sym // hmap.buckets see reflect.go:hmap
				na := pstate.nod(OAS, pstate.nodSym(ODOT, h, bsym), b)
				na = pstate.typecheck(na, Etop)
				init.Append(na)
			}
		}

		if pstate.Isconst(hint, CTINT) && hint.Val().U.(*Mpint).CmpInt64(BUCKETSIZE) <= 0 {
			// Handling make(map[any]any) and
			// make(map[any]any, hint) where hint <= BUCKETSIZE
			// special allows for faster map initialization and
			// improves binary size by using calls with fewer arguments.
			// For hint <= BUCKETSIZE overLoadFactor(hint, 0) is false
			// and no buckets will be allocated by makemap. Therefore,
			// no buckets need to be allocated in this code path.
			if n.Esc == EscNone {
				// Only need to initialize h.hash0 since
				// hmap h has been allocated on the stack already.
				// h.hash0 = fastrand()
				rand := pstate.mkcall("fastrand", pstate.types.Types[TUINT32], init)
				hashsym := hmapType.Field(pstate.types, 4).Sym // hmap.hash0 see reflect.go:hmap
				a := pstate.nod(OAS, pstate.nodSym(ODOT, h, hashsym), rand)
				a = pstate.typecheck(a, Etop)
				a = pstate.walkexpr(a, init)
				init.Append(a)
				n = pstate.nod(OCONVNOP, h, nil)
				n.Type = t
				n = pstate.typecheck(n, Erv)
			} else {
				// Call runtime.makehmap to allocate an
				// hmap on the heap and initialize hmap's hash0 field.
				fn := pstate.syslook("makemap_small")
				fn = pstate.substArgTypes(fn, t.Key(pstate.types), t.Elem(pstate.types))
				n = pstate.mkcall1(fn, n.Type, init)
			}
		} else {
			if n.Esc != EscNone {
				h = pstate.nodnil()
			}
			// Map initialization with a variable or large hint is
			// more complicated. We therefore generate a call to
			// runtime.makemap to intialize hmap and allocate the
			// map buckets.

			// When hint fits into int, use makemap instead of
			// makemap64, which is faster and shorter on 32 bit platforms.
			fnname := "makemap64"
			argtype := pstate.types.Types[TINT64]

			// Type checking guarantees that TIDEAL hint is positive and fits in an int.
			// See checkmake call in TMAP case of OMAKE case in OpSwitch in typecheck1 function.
			// The case of hint overflow when converting TUINT or TUINTPTR to TINT
			// will be handled by the negative range checks in makemap during runtime.
			if hint.Type.IsKind(TIDEAL) || pstate.maxintval[hint.Type.Etype].Cmp(pstate.maxintval[TUINT]) <= 0 {
				fnname = "makemap"
				argtype = pstate.types.Types[TINT]
			}

			fn := pstate.syslook(fnname)
			fn = pstate.substArgTypes(fn, hmapType, t.Key(pstate.types), t.Elem(pstate.types))
			n = pstate.mkcall1(fn, n.Type, init, pstate.typename(n.Type), pstate.conv(hint, argtype), h)
		}

	case OMAKESLICE:
		l := n.Left
		r := n.Right
		if r == nil {
			r = pstate.safeexpr(l, init)
			l = r
		}
		t := n.Type
		if n.Esc == EscNone {
			if !pstate.isSmallMakeSlice(n) {
				pstate.Fatalf("non-small OMAKESLICE with EscNone: %v", n)
			}
			// var arr [r]T
			// n = arr[:l]
			t = pstate.types.NewArray(t.Elem(pstate.types), pstate.nonnegintconst(r)) // [r]T
			var_ := pstate.temp(t)
			a := pstate.nod(OAS, var_, nil) // zero temp
			a = pstate.typecheck(a, Etop)
			init.Append(a)
			r := pstate.nod(OSLICE, var_, nil) // arr[:l]
			r.SetSliceBounds(pstate, nil, l, nil)
			r = pstate.conv(r, n.Type) // in case n.Type is named.
			r = pstate.typecheck(r, Erv)
			r = pstate.walkexpr(r, init)
			n = r
		} else {
			// n escapes; set up a call to makeslice.
			// When len and cap can fit into int, use makeslice instead of
			// makeslice64, which is faster and shorter on 32 bit platforms.

			if t.Elem(pstate.types).NotInHeap() {
				pstate.yyerror("%v is go:notinheap; heap allocation disallowed", t.Elem(pstate.types))
			}

			len, cap := l, r

			fnname := "makeslice64"
			argtype := pstate.types.Types[TINT64]

			// Type checking guarantees that TIDEAL len/cap are positive and fit in an int.
			// The case of len or cap overflow when converting TUINT or TUINTPTR to TINT
			// will be handled by the negative range checks in makeslice during runtime.
			if (len.Type.IsKind(TIDEAL) || pstate.maxintval[len.Type.Etype].Cmp(pstate.maxintval[TUINT]) <= 0) &&
				(cap.Type.IsKind(TIDEAL) || pstate.maxintval[cap.Type.Etype].Cmp(pstate.maxintval[TUINT]) <= 0) {
				fnname = "makeslice"
				argtype = pstate.types.Types[TINT]
			}

			fn := pstate.syslook(fnname)
			fn = pstate.substArgTypes(fn, t.Elem(pstate.types)) // any-1
			n = pstate.mkcall1(fn, t, init, pstate.typename(t.Elem(pstate.types)), pstate.conv(len, argtype), pstate.conv(cap, argtype))
		}

	case ORUNESTR:
		a := pstate.nodnil()
		if n.Esc == EscNone {
			t := pstate.types.NewArray(pstate.types.Types[TUINT8], 4)
			var_ := pstate.temp(t)
			a = pstate.nod(OADDR, var_, nil)
		}

		// intstring(*[4]byte, rune)
		n = pstate.mkcall("intstring", n.Type, init, a, pstate.conv(n.Left, pstate.types.Types[TINT64]))

	case OARRAYBYTESTR:
		a := pstate.nodnil()
		if n.Esc == EscNone {
			// Create temporary buffer for string on stack.
			t := pstate.types.NewArray(pstate.types.Types[TUINT8], tmpstringbufsize)

			a = pstate.nod(OADDR, pstate.temp(t), nil)
		}

		// slicebytetostring(*[32]byte, []byte) string;
		n = pstate.mkcall("slicebytetostring", n.Type, init, a, n.Left)

	// slicebytetostringtmp([]byte) string;
	case OARRAYBYTESTRTMP:
		n.Left = pstate.walkexpr(n.Left, init)

		if !pstate.instrumenting {
			// Let the backend handle OARRAYBYTESTRTMP directly
			// to avoid a function call to slicebytetostringtmp.
			break
		}

		n = pstate.mkcall("slicebytetostringtmp", n.Type, init, n.Left)

	// slicerunetostring(*[32]byte, []rune) string;
	case OARRAYRUNESTR:
		a := pstate.nodnil()

		if n.Esc == EscNone {
			// Create temporary buffer for string on stack.
			t := pstate.types.NewArray(pstate.types.Types[TUINT8], tmpstringbufsize)

			a = pstate.nod(OADDR, pstate.temp(t), nil)
		}

		n = pstate.mkcall("slicerunetostring", n.Type, init, a, n.Left)

	// stringtoslicebyte(*32[byte], string) []byte;
	case OSTRARRAYBYTE:
		a := pstate.nodnil()

		if n.Esc == EscNone {
			// Create temporary buffer for slice on stack.
			t := pstate.types.NewArray(pstate.types.Types[TUINT8], tmpstringbufsize)

			a = pstate.nod(OADDR, pstate.temp(t), nil)
		}

		n = pstate.mkcall("stringtoslicebyte", n.Type, init, a, pstate.conv(n.Left, pstate.types.Types[TSTRING]))

	case OSTRARRAYBYTETMP:
		// []byte(string) conversion that creates a slice
		// referring to the actual string bytes.
		// This conversion is handled later by the backend and
		// is only for use by internal compiler optimizations
		// that know that the slice won't be mutated.
		// The only such case today is:
		// for i, c := range []byte(string)
		n.Left = pstate.walkexpr(n.Left, init)

	// stringtoslicerune(*[32]rune, string) []rune
	case OSTRARRAYRUNE:
		a := pstate.nodnil()

		if n.Esc == EscNone {
			// Create temporary buffer for slice on stack.
			t := pstate.types.NewArray(pstate.types.Types[TINT32], tmpstringbufsize)

			a = pstate.nod(OADDR, pstate.temp(t), nil)
		}

		n = pstate.mkcall("stringtoslicerune", n.Type, init, a, pstate.conv(n.Left, pstate.types.Types[TSTRING]))

	// ifaceeq(i1 any-1, i2 any-2) (ret bool);
	case OCMPIFACE:
		if !pstate.eqtype(n.Left.Type, n.Right.Type) {
			pstate.Fatalf("ifaceeq %v %v %v", n.Op, n.Left.Type, n.Right.Type)
		}
		var fn *Node
		if n.Left.Type.IsEmptyInterface(pstate.types) {
			fn = pstate.syslook("efaceeq")
		} else {
			fn = pstate.syslook("ifaceeq")
		}

		n.Right = pstate.cheapexpr(n.Right, init)
		n.Left = pstate.cheapexpr(n.Left, init)
		lt := pstate.nod(OITAB, n.Left, nil)
		rt := pstate.nod(OITAB, n.Right, nil)
		ld := pstate.nod(OIDATA, n.Left, nil)
		rd := pstate.nod(OIDATA, n.Right, nil)
		ld.Type = pstate.types.Types[TUNSAFEPTR]
		rd.Type = pstate.types.Types[TUNSAFEPTR]
		ld.SetTypecheck(1)
		rd.SetTypecheck(1)
		call := pstate.mkcall1(fn, n.Type, init, lt, ld, rd)

		// Check itable/type before full compare.
		// Note: short-circuited because order matters.
		var cmp *Node
		if n.SubOp(pstate) == OEQ {
			cmp = pstate.nod(OANDAND, pstate.nod(OEQ, lt, rt), call)
		} else {
			cmp = pstate.nod(OOROR, pstate.nod(ONE, lt, rt), pstate.nod(ONOT, call, nil))
		}
		n = pstate.finishcompare(n, cmp, init)

	case OARRAYLIT, OSLICELIT, OMAPLIT, OSTRUCTLIT, OPTRLIT:
		if pstate.isStaticCompositeLiteral(n) && !pstate.canSSAType(n.Type) {
			// n can be directly represented in the read-only data section.
			// Make direct reference to the static data. See issue 12841.
			vstat := pstate.staticname(n.Type)
			vstat.Name.SetReadonly(true)
			pstate.fixedlit(inInitFunction, initKindStatic, n, vstat, init)
			n = vstat
			n = pstate.typecheck(n, Erv)
			break
		}
		var_ := pstate.temp(n.Type)
		pstate.anylit(n, var_, init)
		n = var_

	case OSEND:
		n1 := n.Right
		n1 = pstate.assignconv(n1, n.Left.Type.Elem(pstate.types), "chan send")
		n1 = pstate.walkexpr(n1, init)
		n1 = pstate.nod(OADDR, n1, nil)
		n = pstate.mkcall1(pstate.chanfn("chansend1", 2, n.Left.Type), nil, init, n.Left, n1)

	case OCLOSURE:
		n = pstate.walkclosure(n, init)

	case OCALLPART:
		n = pstate.walkpartialcall(n, init)
	}

	// Expressions that are constant at run time but not
	// considered const by the language spec are not turned into
	// constants until walk. For example, if n is y%1 == 0, the
	// walk of y%1 may have replaced it by 0.
	// Check whether n with its updated args is itself now a constant.
	t := n.Type
	pstate.evconst(n)
	if n.Type != t {
		pstate.Fatalf("evconst changed Type: %v had type %v, now %v", n, t, n.Type)
	}
	if n.Op == OLITERAL {
		n = pstate.typecheck(n, Erv)
		// Emit string symbol now to avoid emitting
		// any concurrently during the backend.
		if s, ok := n.Val().U.(string); ok {
			_ = pstate.stringsym(n.Pos, s)
		}
	}

	pstate.updateHasCall(n)

	if pstate.Debug['w'] != 0 && n != nil {
		Dump("after walk expr", n)
	}

	pstate.lineno = lno
	return n
}

// TODO(josharian): combine this with its caller and simplify
func (pstate *PackageState) reduceSlice(n *Node) *Node {
	low, high, max := n.SliceBounds(pstate)
	if high != nil && high.Op == OLEN && pstate.samesafeexpr(n.Left, high.Left) {
		// Reduce x[i:len(x)] to x[i:].
		high = nil
	}
	n.SetSliceBounds(pstate, low, high, max)
	if (n.Op == OSLICE || n.Op == OSLICESTR) && low == nil && high == nil {
		// Reduce x[:] to x.
		if pstate.Debug_slice > 0 {
			pstate.Warn("slice: omit slice operation")
		}
		return n.Left
	}
	return n
}

func (pstate *PackageState) ascompatee1(l *Node, r *Node, init *Nodes) *Node {
	// convas will turn map assigns into function calls,
	// making it impossible for reorder3 to work.
	n := pstate.nod(OAS, l, r)

	if l.Op == OINDEXMAP {
		return n
	}

	return pstate.convas(n, init)
}

func (pstate *PackageState) ascompatee(op Op, nl, nr []*Node, init *Nodes) []*Node {
	// check assign expression list to
	// an expression list. called in
	//	expr-list = expr-list

	// ensure order of evaluation for function calls
	for i := range nl {
		nl[i] = pstate.safeexpr(nl[i], init)
	}
	for i1 := range nr {
		nr[i1] = pstate.safeexpr(nr[i1], init)
	}

	var nn []*Node
	i := 0
	for ; i < len(nl); i++ {
		if i >= len(nr) {
			break
		}
		// Do not generate 'x = x' during return. See issue 4014.
		if op == ORETURN && pstate.samesafeexpr(nl[i], nr[i]) {
			continue
		}
		nn = append(nn, pstate.ascompatee1(nl[i], nr[i], init))
	}

	// cannot happen: caller checked that lists had same length
	if i < len(nl) || i < len(nr) {
		var nln, nrn Nodes
		nln.Set(nl)
		nrn.Set(nr)
		pstate.Fatalf("error in shape across %+v %v %+v / %d %d [%s]", nln, op, nrn, len(nl), len(nr), pstate.Curfn.funcname())
	}
	return nn
}

// fncall reports whether assigning an rvalue of type rt to an lvalue l might involve a function call.
func (pstate *PackageState) fncall(l *Node, rt *types.Type) bool {
	if l.HasCall() || l.Op == OINDEXMAP {
		return true
	}
	if pstate.eqtype(l.Type, rt) {
		return false
	}
	// There might be a conversion required, which might involve a runtime call.
	return true
}

// check assign type list to
// an expression list. called in
//	expr-list = func()
func (pstate *PackageState) ascompatet(nl Nodes, nr *types.Type) []*Node {
	if nl.Len() != nr.NumFields(pstate.types) {
		pstate.Fatalf("ascompatet: assignment count mismatch: %d = %d", nl.Len(), nr.NumFields(pstate.types))
	}

	var nn, mm Nodes
	for i, l := range nl.Slice() {
		if l.isBlank() {
			continue
		}
		r := nr.Field(pstate.types, i)

		// Any assignment to an lvalue that might cause a function call must be
		// deferred until all the returned values have been read.
		if pstate.fncall(l, r.Type) {
			tmp := pstate.temp(r.Type)
			tmp = pstate.typecheck(tmp, Erv)
			a := pstate.nod(OAS, l, tmp)
			a = pstate.convas(a, &mm)
			mm.Append(a)
			l = tmp
		}

		a := pstate.nod(OAS, l, pstate.nodarg(r, 0))
		a = pstate.convas(a, &nn)
		pstate.updateHasCall(a)
		if a.HasCall() {
			Dump("ascompatet ucount", a)
			pstate.Fatalf("ascompatet: too many function calls evaluating parameters")
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
func (pstate *PackageState) nodarg(t interface{}, fp int) *Node {
	var n *Node

	switch t := t.(type) {
	default:
		pstate.Fatalf("bad nodarg %T(%v)", t, t)

	case *types.Type:
		// Entire argument struct, not just one arg
		if !t.IsFuncArgStruct() {
			pstate.Fatalf("nodarg: bad type %v", t)
		}

		// Build fake variable name for whole arg struct.
		n = pstate.newname(pstate.lookup(".args"))
		n.Type = t
		first := t.Field(pstate.types, 0)
		if first == nil {
			pstate.Fatalf("nodarg: bad struct")
		}
		if first.Offset == BADWIDTH {
			pstate.Fatalf("nodarg: offset not computed for %v", t)
		}
		n.Xoffset = first.Offset

	case *types.Field:
		if fp == 1 {
			// NOTE(rsc): This should be using t.Nname directly,
			// except in the case where t.Nname.Sym is the blank symbol and
			// so the assignment would be discarded during code generation.
			// In that case we need to make a new node, and there is no harm
			// in optimization passes to doing so. But otherwise we should
			// definitely be using the actual declaration and not a newly built node.
			// The extra Fatalf checks here are verifying that this is the case,
			// without changing the actual logic (at time of writing, it's getting
			// toward time for the Go 1.7 beta).
			// At some quieter time (assuming we've never seen these Fatalfs happen)
			// we could change this code to use "expect" directly.
			expect := asNode(t.Nname)
			if expect.isParamHeapCopy() {
				expect = expect.Name.Param.Stackcopy
			}

			for _, n := range pstate.Curfn.Func.Dcl {
				if (n.Class() == PPARAM || n.Class() == PPARAMOUT) && !t.Sym.IsBlank() && n.Sym == t.Sym {
					if n != expect {
						pstate.Fatalf("nodarg: unexpected node: %v (%p %v) vs %v (%p %v)", n, n, n.Op, asNode(t.Nname), asNode(t.Nname), asNode(t.Nname).Op)
					}
					return n
				}
			}

			if !expect.Sym.IsBlank() {
				pstate.Fatalf("nodarg: did not find node in dcl list: %v", expect)
			}
		}

		// Build fake name for individual variable.
		// This is safe because if there was a real declared name
		// we'd have used it above.
		n = pstate.newname(pstate.lookup("__"))
		n.Type = t.Type
		if t.Offset == BADWIDTH {
			pstate.Fatalf("nodarg: offset not computed for %v", t)
		}
		n.Xoffset = t.Offset
		n.Orig = asNode(t.Nname)
	}

	// Rewrite argument named _ to __,
	// or else the assignment to _ will be
	// discarded during code generation.
	if n.isBlank() {
		n.Sym = pstate.lookup("__")
	}

	if fp != 0 {
		pstate.Fatalf("bad fp: %v", fp)
	}

	// preparing arguments for call
	n.Op = OINDREGSP
	n.Xoffset += pstate.Ctxt.FixedFrameSize()
	n.SetTypecheck(1)
	n.SetAddrtaken(true) // keep optimizers at bay
	return n
}

// package all the arguments that match a ... T parameter into a []T.
func (pstate *PackageState) mkdotargslice(typ *types.Type, args []*Node, init *Nodes, ddd *Node) *Node {
	esc := uint16(EscUnknown)
	if ddd != nil {
		esc = ddd.Esc
	}

	if len(args) == 0 {
		n := pstate.nodnil()
		n.Type = typ
		return n
	}

	n := pstate.nod(OCOMPLIT, nil, pstate.typenod(typ))
	if ddd != nil && pstate.prealloc[ddd] != nil {
		pstate.prealloc[n] = pstate.prealloc[ddd] // temporary to use
	}
	n.List.Set(args)
	n.Esc = esc
	n = pstate.typecheck(n, Erv)
	if n.Type == nil {
		pstate.Fatalf("mkdotargslice: typecheck failed")
	}
	n = pstate.walkexpr(n, init)
	return n
}

// check assign expression list to
// a type list. called in
//	return expr-list
//	func(expr-list)
func (pstate *PackageState) ascompatte(call *Node, isddd bool, lhs *types.Type, rhs []*Node, fp int, init *Nodes) []*Node {
	// f(g()) where g has multiple return values
	if len(rhs) == 1 && rhs[0].Type.IsFuncArgStruct() {
		// optimization - can do block copy
		if pstate.eqtypenoname(rhs[0].Type, lhs) {
			nl := pstate.nodarg(lhs, fp)
			nr := pstate.nod(OCONVNOP, rhs[0], nil)
			nr.Type = nl.Type
			n := pstate.convas(pstate.nod(OAS, nl, nr), init)
			n.SetTypecheck(1)
			return []*Node{n}
		}

		// conversions involved.
		// copy into temporaries.
		var tmps []*Node
		for _, nr := range rhs[0].Type.FieldSlice(pstate.types) {
			tmps = append(tmps, pstate.temp(nr.Type))
		}

		a := pstate.nod(OAS2, nil, nil)
		a.List.Set(tmps)
		a.Rlist.Set(rhs)
		a = pstate.typecheck(a, Etop)
		a = pstate.walkstmt(a)
		init.Append(a)

		rhs = tmps
	}

	// For each parameter (LHS), assign its corresponding argument (RHS).
	// If there's a ... parameter (which is only valid as the final
	// parameter) and this is not a ... call expression,
	// then assign the remaining arguments as a slice.
	var nn []*Node
	for i, nl := range lhs.FieldSlice(pstate.types) {
		var nr *Node
		if nl.Isddd() && !isddd {
			nr = pstate.mkdotargslice(nl.Type, rhs[i:], init, call.Right)
		} else {
			nr = rhs[i]
		}

		a := pstate.nod(OAS, pstate.nodarg(nl, fp), nr)
		a = pstate.convas(a, init)
		a.SetTypecheck(1)
		nn = append(nn, a)
	}

	return nn
}

// generate code for print
func (pstate *PackageState) walkprint(nn *Node, init *Nodes) *Node {
	// Hoist all the argument evaluation up before the lock.
	pstate.walkexprlistcheap(nn.List.Slice(), init)

	// For println, add " " between elements and "\n" at the end.
	if nn.Op == OPRINTN {
		s := nn.List.Slice()
		t := make([]*Node, 0, len(s)*2)
		for i, n := range s {
			if i != 0 {
				t = append(t, pstate.nodstr(" "))
			}
			t = append(t, n)
		}
		t = append(t, pstate.nodstr("\n"))
		nn.List.Set(t)
	}

	// Collapse runs of constant strings.
	s := nn.List.Slice()
	t := make([]*Node, 0, len(s))
	for i := 0; i < len(s); {
		var strs []string
		for i < len(s) && pstate.Isconst(s[i], CTSTR) {
			strs = append(strs, s[i].Val().U.(string))
			i++
		}
		if len(strs) > 0 {
			t = append(t, pstate.nodstr(strings.Join(strs, "")))
		}
		if i < len(s) {
			t = append(t, s[i])
			i++
		}
	}
	nn.List.Set(t)

	calls := []*Node{pstate.mkcall("printlock", nil, init)}
	for i, n := range nn.List.Slice() {
		if n.Op == OLITERAL {
			switch n.Val().Ctype(pstate) {
			case CTRUNE:
				n = pstate.defaultlit(n, pstate.types.Runetype)

			case CTINT:
				n = pstate.defaultlit(n, pstate.types.Types[TINT64])

			case CTFLT:
				n = pstate.defaultlit(n, pstate.types.Types[TFLOAT64])
			}
		}

		if n.Op != OLITERAL && n.Type != nil && n.Type.Etype == TIDEAL {
			n = pstate.defaultlit(n, pstate.types.Types[TINT64])
		}
		n = pstate.defaultlit(n, nil)
		nn.List.SetIndex(i, n)
		if n.Type == nil || n.Type.Etype == TFORW {
			continue
		}

		var on *Node
		switch n.Type.Etype {
		case TINTER:
			if n.Type.IsEmptyInterface(pstate.types) {
				on = pstate.syslook("printeface")
			} else {
				on = pstate.syslook("printiface")
			}
			on = pstate.substArgTypes(on, n.Type) // any-1
		case TPTR32, TPTR64, TCHAN, TMAP, TFUNC, TUNSAFEPTR:
			on = pstate.syslook("printpointer")
			on = pstate.substArgTypes(on, n.Type) // any-1
		case TSLICE:
			on = pstate.syslook("printslice")
			on = pstate.substArgTypes(on, n.Type) // any-1
		case TUINT, TUINT8, TUINT16, TUINT32, TUINT64, TUINTPTR:
			if pstate.isRuntimePkg(n.Type.Sym.Pkg) && n.Type.Sym.Name == "hex" {
				on = pstate.syslook("printhex")
			} else {
				on = pstate.syslook("printuint")
			}
		case TINT, TINT8, TINT16, TINT32, TINT64:
			on = pstate.syslook("printint")
		case TFLOAT32, TFLOAT64:
			on = pstate.syslook("printfloat")
		case TCOMPLEX64, TCOMPLEX128:
			on = pstate.syslook("printcomplex")
		case TBOOL:
			on = pstate.syslook("printbool")
		case TSTRING:
			cs := ""
			if pstate.Isconst(n, CTSTR) {
				cs = n.Val().U.(string)
			}
			switch cs {
			case " ":
				on = pstate.syslook("printsp")
			case "\n":
				on = pstate.syslook("printnl")
			default:
				on = pstate.syslook("printstring")
			}
		default:
			pstate.badtype(OPRINT, n.Type, nil)
			continue
		}

		r := pstate.nod(OCALL, on, nil)
		if params := on.Type.Params(pstate.types).FieldSlice(pstate.types); len(params) > 0 {
			t := params[0].Type
			if !pstate.eqtype(t, n.Type) {
				n = pstate.nod(OCONV, n, nil)
				n.Type = t
			}
			r.List.Append(n)
		}
		calls = append(calls, r)
	}

	calls = append(calls, pstate.mkcall("printunlock", nil, init))

	pstate.typecheckslice(calls, Etop)
	pstate.walkexprlist(calls, init)

	r := pstate.nod(OEMPTY, nil, nil)
	r = pstate.typecheck(r, Etop)
	r = pstate.walkexpr(r, init)
	r.Ninit.Set(calls)
	return r
}

func (pstate *PackageState) callnew(t *types.Type) *Node {
	if t.NotInHeap() {
		pstate.yyerror("%v is go:notinheap; heap allocation disallowed", t)
	}
	pstate.dowidth(t)
	fn := pstate.syslook("newobject")
	fn = pstate.substArgTypes(fn, t)
	v := pstate.mkcall1(fn, pstate.types.NewPtr(t), nil, pstate.typename(t))
	v.SetNonNil(true)
	return v
}

func (pstate *PackageState) iscallret(n *Node) bool {
	if n == nil {
		return false
	}
	n = pstate.outervalue(n)
	return n.Op == OINDREGSP
}

// isReflectHeaderDataField reports whether l is an expression p.Data
// where p has type reflect.SliceHeader or reflect.StringHeader.
func (pstate *PackageState) isReflectHeaderDataField(l *Node) bool {
	if l.Type != pstate.types.Types[TUINTPTR] {
		return false
	}

	var tsym *types.Sym
	switch l.Op {
	case ODOT:
		tsym = l.Left.Type.Sym
	case ODOTPTR:
		tsym = l.Left.Type.Elem(pstate.types).Sym
	default:
		return false
	}

	if tsym == nil || l.Sym.Name != "Data" || tsym.Pkg.Path != "reflect" {
		return false
	}
	return tsym.Name == "SliceHeader" || tsym.Name == "StringHeader"
}

func (pstate *PackageState) convas(n *Node, init *Nodes) *Node {
	if n.Op != OAS {
		pstate.Fatalf("convas: not OAS %v", n.Op)
	}
	defer pstate.updateHasCall(n)

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
		n.Right = pstate.defaultlit(n.Right, nil)
		return n
	}

	if !pstate.eqtype(lt, rt) {
		n.Right = pstate.assignconv(n.Right, lt, "assignment")
		n.Right = pstate.walkexpr(n.Right, init)
	}
	pstate.dowidth(n.Right.Type)

	return n
}

// from ascompat[te]
// evaluating actual function arguments.
//	f(a,b)
// if there is exactly one function expr,
// then it is done first. otherwise must
// make temp variables
func (pstate *PackageState) reorder1(all []*Node) []*Node {
	// When instrumenting, force all arguments into temporary
	// variables to prevent instrumentation calls from clobbering
	// arguments already on the stack.

	funcCalls := 0
	if !pstate.instrumenting {
		if len(all) == 1 {
			return all
		}

		for _, n := range all {
			pstate.updateHasCall(n)
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
		if !pstate.instrumenting {
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

		// make assignment of fncall to tempname
		a := pstate.temp(n.Right.Type)

		a = pstate.nod(OAS, a, n.Right)
		g = append(g, a)

		// put normal arg assignment on list
		// with fncall replaced by tempname
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
func (pstate *PackageState) reorder3(all []*Node) []*Node {
	// If a needed expression may be affected by an
	// earlier assignment, make an early copy of that
	// expression and use the copy instead.
	var early []*Node

	var mapinit Nodes
	for i, n := range all {
		l := n.Left

		// Save subexpressions needed on left side.
		// Drill through non-dereferences.
		for {
			if l.Op == ODOT || l.Op == OPAREN {
				l = l.Left
				continue
			}

			if l.Op == OINDEX && l.Left.Type.IsArray() {
				l.Right = pstate.reorder3save(l.Right, all, i, &early)
				l = l.Left
				continue
			}

			break
		}

		switch l.Op {
		default:
			pstate.Fatalf("reorder3 unexpected lvalue %#v", l.Op)

		case ONAME:
			break

		case OINDEX, OINDEXMAP:
			l.Left = pstate.reorder3save(l.Left, all, i, &early)
			l.Right = pstate.reorder3save(l.Right, all, i, &early)
			if l.Op == OINDEXMAP {
				all[i] = pstate.convas(all[i], &mapinit)
			}

		case OIND, ODOTPTR:
			l.Left = pstate.reorder3save(l.Left, all, i, &early)
		}

		// Save expression on right side.
		all[i].Right = pstate.reorder3save(all[i].Right, all, i, &early)
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
func (pstate *PackageState) reorder3save(n *Node, all []*Node, i int, early *[]*Node) *Node {
	if !pstate.aliased(n, all, i) {
		return n
	}

	q := pstate.temp(n.Type)
	q = pstate.nod(OAS, q, n)
	q = pstate.typecheck(q, Etop)
	*early = append(*early, q)
	return q.Left
}

// what's the outer value that a write to n affects?
// outer value means containing struct or array.
func (pstate *PackageState) outervalue(n *Node) *Node {
	for {
		switch n.Op {
		case OXDOT:
			pstate.Fatalf("OXDOT in walk")
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
func (pstate *PackageState) aliased(n *Node, all []*Node, i int) bool {
	if n == nil {
		return false
	}

	// Treat all fields of a struct as referring to the whole struct.
	// We could do better but we would have to keep track of the fields.
	for n.Op == ODOT {
		n = n.Left
	}

	// Look for obvious aliasing: a variable being assigned
	// during the all list and appearing in n.
	// Also record whether there are any writes to main memory.
	// Also record whether there are any writes to variables
	// whose addresses have been taken.
	memwrite := false
	varwrite := false
	for _, an := range all[:i] {
		a := pstate.outervalue(an.Left)

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
				// Direct hit.
				return true
			}
		}
	}

	// The variables being written do not appear in n.
	// However, n might refer to computed addresses
	// that are being written.

	// If no computed addresses are affected by the writes, no aliasing.
	if !memwrite && !varwrite {
		return false
	}

	// If n does not refer to computed addresses
	// (that is, if n only refers to variables whose addresses
	// have not been taken), no aliasing.
	if pstate.varexpr(n) {
		return false
	}

	// Otherwise, both the writes and n refer to computed memory addresses.
	// Assume that they might conflict.
	return true
}

// does the evaluation of n only refer to variables
// whose addresses have not been taken?
// (and no other memory)
func (pstate *PackageState) varexpr(n *Node) bool {
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
		return pstate.varexpr(n.Left) && pstate.varexpr(n.Right)

	case ODOT: // but not ODOTPTR
		// Should have been handled in aliased.
		pstate.Fatalf("varexpr unexpected ODOT")
	}

	// Be conservative.
	return false
}

// is the name l mentioned in r?
func vmatch2(l *Node, r *Node) bool {
	if r == nil {
		return false
	}
	switch r.Op {
	// match each right given left
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
	// isolate all left sides
	if l == nil || r == nil {
		return false
	}
	switch l.Op {
	case ONAME:
		switch l.Class() {
		case PPARAM, PAUTO:
			break

		default:
			// assignment to non-stack variable must be
			// delayed if right has function calls.
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
func (pstate *PackageState) paramstoheap(params *types.Type) []*Node {
	var nn []*Node
	for _, t := range params.Fields(pstate.types).Slice() {
		v := asNode(t.Nname)
		if v != nil && v.Sym != nil && strings.HasPrefix(v.Sym.Name, "~r") { // unnamed result
			v = nil
		}
		if v == nil {
			continue
		}

		if stackcopy := v.Name.Param.Stackcopy; stackcopy != nil {
			nn = append(nn, pstate.walkstmt(pstate.nod(ODCL, v, nil)))
			if stackcopy.Class() == PPARAM {
				nn = append(nn, pstate.walkstmt(pstate.typecheck(pstate.nod(OAS, v, stackcopy), Etop)))
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
func (pstate *PackageState) zeroResults() {
	for _, f := range pstate.Curfn.Type.Results(pstate.types).Fields(pstate.types).Slice() {
		if v := asNode(f.Nname); v != nil && v.Name.Param.Heapaddr != nil {
			// The local which points to the return value is the
			// thing that needs zeroing. This is already handled
			// by a Needzero annotation in plive.go:livenessepilogue.
			continue
		}
		// Zero the stack location containing f.
		pstate.Curfn.Func.Enter.Append(pstate.nodl(pstate.Curfn.Pos, OAS, pstate.nodarg(f, 1), nil))
	}
}

// returnsfromheap returns code to copy values for heap-escaped parameters
// back to the stack.
func (pstate *PackageState) returnsfromheap(params *types.Type) []*Node {
	var nn []*Node
	for _, t := range params.Fields(pstate.types).Slice() {
		v := asNode(t.Nname)
		if v == nil {
			continue
		}
		if stackcopy := v.Name.Param.Stackcopy; stackcopy != nil && stackcopy.Class() == PPARAMOUT {
			nn = append(nn, pstate.walkstmt(pstate.typecheck(pstate.nod(OAS, stackcopy, v), Etop)))
		}
	}

	return nn
}

// heapmoves generates code to handle migrating heap-escaped parameters
// between the stack and the heap. The generated code is added to Curfn's
// Enter and Exit lists.
func (pstate *PackageState) heapmoves() {
	lno := pstate.lineno
	pstate.lineno = pstate.Curfn.Pos
	nn := pstate.paramstoheap(pstate.Curfn.Type.Recvs(pstate.types))
	nn = append(nn, pstate.paramstoheap(pstate.Curfn.Type.Params(pstate.types))...)
	nn = append(nn, pstate.paramstoheap(pstate.Curfn.Type.Results(pstate.types))...)
	pstate.Curfn.Func.Enter.Append(nn...)
	pstate.lineno = pstate.Curfn.Func.Endlineno
	pstate.Curfn.Func.Exit.Append(pstate.returnsfromheap(pstate.Curfn.Type.Results(pstate.types))...)
	pstate.lineno = lno
}

func (pstate *PackageState) vmkcall(fn *Node, t *types.Type, init *Nodes, va []*Node) *Node {
	if fn.Type == nil || fn.Type.Etype != TFUNC {
		pstate.Fatalf("mkcall %v %v", fn, fn.Type)
	}

	n := fn.Type.NumParams(pstate.types)
	if n != len(va) {
		pstate.Fatalf("vmkcall %v needs %v args got %v", fn, n, len(va))
	}

	r := pstate.nod(OCALL, fn, nil)
	r.List.Set(va)
	if fn.Type.NumResults(pstate.types) > 0 {
		r = pstate.typecheck(r, Erv|Efnstruct)
	} else {
		r = pstate.typecheck(r, Etop)
	}
	r = pstate.walkexpr(r, init)
	r.Type = t
	return r
}

func (pstate *PackageState) mkcall(name string, t *types.Type, init *Nodes, args ...*Node) *Node {
	return pstate.vmkcall(pstate.syslook(name), t, init, args)
}

func (pstate *PackageState) mkcall1(fn *Node, t *types.Type, init *Nodes, args ...*Node) *Node {
	return pstate.vmkcall(fn, t, init, args)
}

func (pstate *PackageState) conv(n *Node, t *types.Type) *Node {
	if pstate.eqtype(n.Type, t) {
		return n
	}
	n = pstate.nod(OCONV, n, nil)
	n.Type = t
	n = pstate.typecheck(n, Erv)
	return n
}

// byteindex converts n, which is byte-sized, to a uint8.
// We cannot use conv, because we allow converting bool to uint8 here,
// which is forbidden in user code.
func (pstate *PackageState) byteindex(n *Node) *Node {
	if pstate.eqtype(n.Type, pstate.types.Types[TUINT8]) {
		return n
	}
	n = pstate.nod(OCONV, n, nil)
	n.Type = pstate.types.Types[TUINT8]
	n.SetTypecheck(1)
	return n
}

func (pstate *PackageState) chanfn(name string, n int, t *types.Type) *Node {
	if !t.IsChan() {
		pstate.Fatalf("chanfn %v", t)
	}
	fn := pstate.syslook(name)
	switch n {
	default:
		pstate.Fatalf("chanfn %d", n)
	case 1:
		fn = pstate.substArgTypes(fn, t.Elem(pstate.types))
	case 2:
		fn = pstate.substArgTypes(fn, t.Elem(pstate.types), t.Elem(pstate.types))
	}
	return fn
}

func (pstate *PackageState) mapfn(name string, t *types.Type) *Node {
	if !t.IsMap() {
		pstate.Fatalf("mapfn %v", t)
	}
	fn := pstate.syslook(name)
	fn = pstate.substArgTypes(fn, t.Key(pstate.types), t.Elem(pstate.types), t.Key(pstate.types), t.Elem(pstate.types))
	return fn
}

func (pstate *PackageState) mapfndel(name string, t *types.Type) *Node {
	if !t.IsMap() {
		pstate.Fatalf("mapfn %v", t)
	}
	fn := pstate.syslook(name)
	fn = pstate.substArgTypes(fn, t.Key(pstate.types), t.Elem(pstate.types), t.Key(pstate.types))
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

func (pstate *PackageState) mapfast(t *types.Type) int {
	// Check runtime/map.go:maxValueSize before changing.
	if t.Elem(pstate.types).Width > 128 {
		return mapslow
	}
	switch pstate.algtype(t.Key(pstate.types)) {
	case AMEM32:
		if !t.Key(pstate.types).HasHeapPointer(pstate.types) {
			return mapfast32
		}
		if pstate.Widthptr == 4 {
			return mapfast32ptr
		}
		pstate.Fatalf("small pointer %v", t.Key(pstate.types))
	case AMEM64:
		if !t.Key(pstate.types).HasHeapPointer(pstate.types) {
			return mapfast64
		}
		if pstate.Widthptr == 8 {
			return mapfast64ptr
		}
	// Two-word object, at least one of which is a pointer.
	// Use the slow path.
	case ASTRING:
		return mapfaststr
	}
	return mapslow
}

func (pstate *PackageState) writebarrierfn(name string, l *types.Type, r *types.Type) *Node {
	fn := pstate.syslook(name)
	fn = pstate.substArgTypes(fn, l, r)
	return fn
}

func (pstate *PackageState) addstr(n *Node, init *Nodes) *Node {
	// orderexpr rewrote OADDSTR to have a list of strings.
	c := n.List.Len()

	if c < 2 {
		pstate.Fatalf("addstr count %d too small", c)
	}

	buf := pstate.nodnil()
	if n.Esc == EscNone {
		sz := int64(0)
		for _, n1 := range n.List.Slice() {
			if n1.Op == OLITERAL {
				sz += int64(len(n1.Val().U.(string)))
			}
		}

		// Don't allocate the buffer if the result won't fit.
		if sz < tmpstringbufsize {
			// Create temporary buffer for result string on stack.
			t := pstate.types.NewArray(pstate.types.Types[TUINT8], tmpstringbufsize)

			buf = pstate.nod(OADDR, pstate.temp(t), nil)
		}
	}

	// build list of string arguments
	args := []*Node{buf}
	for _, n2 := range n.List.Slice() {
		args = append(args, pstate.conv(n2, pstate.types.Types[TSTRING]))
	}

	var fn string
	if c <= 5 {
		// small numbers of strings use direct runtime helpers.
		// note: orderexpr knows this cutoff too.
		fn = fmt.Sprintf("concatstring%d", c)
	} else {
		// large numbers of strings are passed to the runtime as a slice.
		fn = "concatstrings"

		t := pstate.types.NewSlice(pstate.types.Types[TSTRING])
		slice := pstate.nod(OCOMPLIT, nil, pstate.typenod(t))
		if pstate.prealloc[n] != nil {
			pstate.prealloc[slice] = pstate.prealloc[n]
		}
		slice.List.Set(args[1:]) // skip buf arg
		args = []*Node{buf, slice}
		slice.Esc = EscNone
	}

	cat := pstate.syslook(fn)
	r := pstate.nod(OCALL, cat, nil)
	r.List.Set(args)
	r = pstate.typecheck(r, Erv)
	r = pstate.walkexpr(r, init)
	r.Type = n.Type

	return r
}

func (pstate *PackageState) walkAppendArgs(n *Node, init *Nodes) {
	pstate.walkexprlistsafe(n.List.Slice(), init)

	// walkexprlistsafe will leave OINDEX (s[n]) alone if both s
	// and n are name or literal, but those may index the slice we're
	// modifying here. Fix explicitly.
	ls := n.List.Slice()
	for i1, n1 := range ls {
		ls[i1] = pstate.cheapexpr(n1, init)
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
func (pstate *PackageState) appendslice(n *Node, init *Nodes) *Node {
	pstate.walkAppendArgs(n, init)

	l1 := n.List.First()
	l2 := n.List.Second()

	var l []*Node

	// var s []T
	s := pstate.temp(l1.Type)
	l = append(l, pstate.nod(OAS, s, l1)) // s = l1

	// n := len(s) + len(l2)
	nn := pstate.temp(pstate.types.Types[TINT])
	l = append(l, pstate.nod(OAS, nn, pstate.nod(OADD, pstate.nod(OLEN, s, nil), pstate.nod(OLEN, l2, nil))))

	// if uint(n) > uint(cap(s))
	nif := pstate.nod(OIF, nil, nil)
	nif.Left = pstate.nod(OGT, pstate.nod(OCONV, nn, nil), pstate.nod(OCONV, pstate.nod(OCAP, s, nil), nil))
	nif.Left.Left.Type = pstate.types.Types[TUINT]
	nif.Left.Right.Type = pstate.types.Types[TUINT]

	// instantiate growslice(Type*, []any, int) []any
	fn := pstate.syslook("growslice")
	fn = pstate.substArgTypes(fn, s.Type.Elem(pstate.types), s.Type.Elem(pstate.types))

	// s = growslice(T, s, n)
	nif.Nbody.Set1(pstate.nod(OAS, s, pstate.mkcall1(fn, s.Type, &nif.Ninit, pstate.typename(s.Type.Elem(pstate.types)), s, nn)))
	l = append(l, nif)

	// s = s[:n]
	nt := pstate.nod(OSLICE, s, nil)
	nt.SetSliceBounds(pstate, nil, nn, nil)
	l = append(l, pstate.nod(OAS, s, nt))

	if l1.Type.Elem(pstate.types).HasHeapPointer(pstate.types) {
		// copy(s[len(l1):], l2)
		nptr1 := pstate.nod(OSLICE, s, nil)
		nptr1.SetSliceBounds(pstate, pstate.nod(OLEN, l1, nil), nil, nil)
		nptr2 := l2
		pstate.Curfn.Func.setWBPos(pstate, n.Pos)
		fn := pstate.syslook("typedslicecopy")
		fn = pstate.substArgTypes(fn, l1.Type, l2.Type)
		var ln Nodes
		ln.Set(l)
		nt := pstate.mkcall1(fn, pstate.types.Types[TINT], &ln, pstate.typename(l1.Type.Elem(pstate.types)), nptr1, nptr2)
		l = append(ln.Slice(), nt)
	} else if pstate.instrumenting && !pstate.compiling_runtime {
		// rely on runtime to instrument copy.
		// copy(s[len(l1):], l2)
		nptr1 := pstate.nod(OSLICE, s, nil)
		nptr1.SetSliceBounds(pstate, pstate.nod(OLEN, l1, nil), nil, nil)
		nptr2 := l2

		var ln Nodes
		ln.Set(l)
		var nt *Node
		if l2.Type.IsString() {
			fn := pstate.syslook("slicestringcopy")
			fn = pstate.substArgTypes(fn, l1.Type, l2.Type)
			nt = pstate.mkcall1(fn, pstate.types.Types[TINT], &ln, nptr1, nptr2)
		} else {
			fn := pstate.syslook("slicecopy")
			fn = pstate.substArgTypes(fn, l1.Type, l2.Type)
			nt = pstate.mkcall1(fn, pstate.types.Types[TINT], &ln, nptr1, nptr2, pstate.nodintconst(s.Type.Elem(pstate.types).Width))
		}

		l = append(ln.Slice(), nt)
	} else {
		// memmove(&s[len(l1)], &l2[0], len(l2)*sizeof(T))
		nptr1 := pstate.nod(OINDEX, s, pstate.nod(OLEN, l1, nil))
		nptr1.SetBounded(true)

		nptr1 = pstate.nod(OADDR, nptr1, nil)

		nptr2 := pstate.nod(OSPTR, l2, nil)

		fn := pstate.syslook("memmove")
		fn = pstate.substArgTypes(fn, s.Type.Elem(pstate.types), s.Type.Elem(pstate.types))

		var ln Nodes
		ln.Set(l)
		nwid := pstate.cheapexpr(pstate.conv(pstate.nod(OLEN, l2, nil), pstate.types.Types[TUINTPTR]), &ln)

		nwid = pstate.nod(OMUL, nwid, pstate.nodintconst(s.Type.Elem(pstate.types).Width))
		nt := pstate.mkcall1(fn, nil, &ln, nptr1, nptr2, nwid)
		l = append(ln.Slice(), nt)
	}

	pstate.typecheckslice(l, Etop)
	pstate.walkstmtlist(l)
	init.Append(l...)
	return s
}

// isAppendOfMake reports whether n is of the form append(x , make([]T, y)...).
// isAppendOfMake assumes n has already been typechecked.
func (pstate *PackageState) isAppendOfMake(n *Node) bool {
	if pstate.Debug['N'] != 0 || pstate.instrumenting {
		return false
	}

	if n.Typecheck() == 0 {
		pstate.Fatalf("missing typecheck: %+v", n)
	}

	if n.Op != OAPPEND || !n.Isddd() || n.List.Len() != 2 {
		return false
	}

	second := n.List.Second()
	if second.Op != OMAKESLICE || second.Right != nil {
		return false
	}

	// y must be either an integer constant or a variable of type int.
	// typecheck checks that constant arguments to make are not negative and
	// fit into an int.
	// runtime.growslice uses int as type for the newcap argument.
	// Constraining variables to be type int avoids the need for runtime checks
	// that e.g. check if an int64 value fits into an int.
	// TODO(moehrmann): support other integer types that always fit in an int
	y := second.Left
	if !pstate.Isconst(y, CTINT) && y.Type.Etype != TINT {
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
func (pstate *PackageState) extendslice(n *Node, init *Nodes) *Node {
	// isAppendOfMake made sure l2 fits in an int.
	l2 := pstate.conv(n.List.Second().Left, pstate.types.Types[TINT])
	l2 = pstate.typecheck(l2, Erv)
	n.List.SetSecond(l2) // walkAppendArgs expects l2 in n.List.Second().

	pstate.walkAppendArgs(n, init)

	l1 := n.List.First()
	l2 = n.List.Second() // re-read l2, as it may have been updated by walkAppendArgs

	var nodes []*Node

	// if l2 < 0
	nifneg := pstate.nod(OIF, pstate.nod(OLT, l2, pstate.nodintconst(0)), nil)
	nifneg.SetLikely(false)

	// panicmakeslicelen()
	nifneg.Nbody.Set1(pstate.mkcall("panicmakeslicelen", nil, init))
	nodes = append(nodes, nifneg)

	// s := l1
	s := pstate.temp(l1.Type)
	nodes = append(nodes, pstate.nod(OAS, s, l1))

	elemtype := s.Type.Elem(pstate.types)

	// n := len(s) + l2
	nn := pstate.temp(pstate.types.Types[TINT])
	nodes = append(nodes, pstate.nod(OAS, nn, pstate.nod(OADD, pstate.nod(OLEN, s, nil), l2)))

	// if uint(n) > uint(cap(s))
	nuint := pstate.conv(nn, pstate.types.Types[TUINT])
	capuint := pstate.conv(pstate.nod(OCAP, s, nil), pstate.types.Types[TUINT])
	nif := pstate.nod(OIF, pstate.nod(OGT, nuint, capuint), nil)

	// instantiate growslice(typ *type, old []any, newcap int) []any
	fn := pstate.syslook("growslice")
	fn = pstate.substArgTypes(fn, elemtype, elemtype)

	// s = growslice(T, s, n)
	nif.Nbody.Set1(pstate.nod(OAS, s, pstate.mkcall1(fn, s.Type, &nif.Ninit, pstate.typename(elemtype), s, nn)))
	nodes = append(nodes, nif)

	// s = s[:n]
	nt := pstate.nod(OSLICE, s, nil)
	nt.SetSliceBounds(pstate, nil, nn, nil)
	nodes = append(nodes, pstate.nod(OAS, s, nt))

	// lptr := &l1[0]
	l1ptr := pstate.temp(l1.Type.Elem(pstate.types).PtrTo(pstate.types))
	tmp := pstate.nod(OSPTR, l1, nil)
	nodes = append(nodes, pstate.nod(OAS, l1ptr, tmp))

	// sptr := &s[0]
	sptr := pstate.temp(elemtype.PtrTo(pstate.types))
	tmp = pstate.nod(OSPTR, s, nil)
	nodes = append(nodes, pstate.nod(OAS, sptr, tmp))

	// hp := &s[len(l1)]
	hp := pstate.nod(OINDEX, s, pstate.nod(OLEN, l1, nil))
	hp.SetBounded(true)
	hp = pstate.nod(OADDR, hp, nil)
	hp = pstate.nod(OCONVNOP, hp, nil)
	hp.Type = pstate.types.Types[TUNSAFEPTR]

	// hn := l2 * sizeof(elem(s))
	hn := pstate.nod(OMUL, l2, pstate.nodintconst(elemtype.Width))
	hn = pstate.conv(hn, pstate.types.Types[TUINTPTR])

	clrname := "memclrNoHeapPointers"
	hasPointers := pstate.types.Haspointers(elemtype)
	if hasPointers {
		clrname = "memclrHasPointers"
	}

	var clr Nodes
	clrfn := pstate.mkcall(clrname, nil, &clr, hp, hn)
	clr.Append(clrfn)

	if hasPointers {
		// if l1ptr == sptr
		nifclr := pstate.nod(OIF, pstate.nod(OEQ, l1ptr, sptr), nil)
		nifclr.Nbody = clr
		nodes = append(nodes, nifclr)
	} else {
		nodes = append(nodes, clr.Slice()...)
	}

	pstate.typecheckslice(nodes, Etop)
	pstate.walkstmtlist(nodes)
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
func (pstate *PackageState) walkappend(n *Node, init *Nodes, dst *Node) *Node {
	if !pstate.samesafeexpr(dst, n.List.First()) {
		n.List.SetFirst(pstate.safeexpr(n.List.First(), init))
		n.List.SetFirst(pstate.walkexpr(n.List.First(), init))
	}
	pstate.walkexprlistsafe(n.List.Slice()[1:], init)

	// walkexprlistsafe will leave OINDEX (s[n]) alone if both s
	// and n are name or literal, but those may index the slice we're
	// modifying here. Fix explicitly.
	// Using cheapexpr also makes sure that the evaluation
	// of all arguments (and especially any panics) happen
	// before we begin to modify the slice in a visible way.
	ls := n.List.Slice()[1:]
	for i, n := range ls {
		ls[i] = pstate.cheapexpr(n, init)
	}

	nsrc := n.List.First()

	argc := n.List.Len() - 1
	if argc < 1 {
		return nsrc
	}

	// General case, with no function calls left as arguments.
	// Leave for gen, except that instrumentation requires old form.
	if !pstate.instrumenting || pstate.compiling_runtime {
		return n
	}

	var l []*Node

	ns := pstate.temp(nsrc.Type)
	l = append(l, pstate.nod(OAS, ns, nsrc)) // s = src

	na := pstate.nodintconst(int64(argc)) // const argc
	nx := pstate.nod(OIF, nil, nil)       // if cap(s) - len(s) < argc
	nx.Left = pstate.nod(OLT, pstate.nod(OSUB, pstate.nod(OCAP, ns, nil), pstate.nod(OLEN, ns, nil)), na)

	fn := pstate.syslook("growslice") //   growslice(<type>, old []T, mincap int) (ret []T)
	fn = pstate.substArgTypes(fn, ns.Type.Elem(pstate.types), ns.Type.Elem(pstate.types))

	nx.Nbody.Set1(pstate.nod(OAS, ns,
		pstate.mkcall1(fn, ns.Type, &nx.Ninit, pstate.typename(ns.Type.Elem(pstate.types)), ns,
			pstate.nod(OADD, pstate.nod(OLEN, ns, nil), na))))

	l = append(l, nx)

	nn := pstate.temp(pstate.types.Types[TINT])
	l = append(l, pstate.nod(OAS, nn, pstate.nod(OLEN, ns, nil))) // n = len(s)

	nx = pstate.nod(OSLICE, ns, nil) // ...s[:n+argc]
	nx.SetSliceBounds(pstate, nil, pstate.nod(OADD, nn, na), nil)
	l = append(l, pstate.nod(OAS, ns, nx)) // s = s[:n+argc]

	ls = n.List.Slice()[1:]
	for i, n := range ls {
		nx = pstate.nod(OINDEX, ns, nn) // s[n] ...
		nx.SetBounded(true)
		l = append(l, pstate.nod(OAS, nx, n)) // s[n] = arg
		if i+1 < len(ls) {
			l = append(l, pstate.nod(OAS, nn, pstate.nod(OADD, nn, pstate.nodintconst(1)))) // n = n + 1
		}
	}

	pstate.typecheckslice(l, Etop)
	pstate.walkstmtlist(l)
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
func (pstate *PackageState) copyany(n *Node, init *Nodes, runtimecall bool) *Node {
	if n.Left.Type.Elem(pstate.types).HasHeapPointer(pstate.types) {
		pstate.Curfn.Func.setWBPos(pstate, n.Pos)
		fn := pstate.writebarrierfn("typedslicecopy", n.Left.Type, n.Right.Type)
		return pstate.mkcall1(fn, n.Type, init, pstate.typename(n.Left.Type.Elem(pstate.types)), n.Left, n.Right)
	}

	if runtimecall {
		if n.Right.Type.IsString() {
			fn := pstate.syslook("slicestringcopy")
			fn = pstate.substArgTypes(fn, n.Left.Type, n.Right.Type)
			return pstate.mkcall1(fn, n.Type, init, n.Left, n.Right)
		}

		fn := pstate.syslook("slicecopy")
		fn = pstate.substArgTypes(fn, n.Left.Type, n.Right.Type)
		return pstate.mkcall1(fn, n.Type, init, n.Left, n.Right, pstate.nodintconst(n.Left.Type.Elem(pstate.types).Width))
	}

	n.Left = pstate.walkexpr(n.Left, init)
	n.Right = pstate.walkexpr(n.Right, init)
	nl := pstate.temp(n.Left.Type)
	nr := pstate.temp(n.Right.Type)
	var l []*Node
	l = append(l, pstate.nod(OAS, nl, n.Left))
	l = append(l, pstate.nod(OAS, nr, n.Right))

	nfrm := pstate.nod(OSPTR, nr, nil)
	nto := pstate.nod(OSPTR, nl, nil)

	nlen := pstate.temp(pstate.types.Types[TINT])

	// n = len(to)
	l = append(l, pstate.nod(OAS, nlen, pstate.nod(OLEN, nl, nil)))

	// if n > len(frm) { n = len(frm) }
	nif := pstate.nod(OIF, nil, nil)

	nif.Left = pstate.nod(OGT, nlen, pstate.nod(OLEN, nr, nil))
	nif.Nbody.Append(pstate.nod(OAS, nlen, pstate.nod(OLEN, nr, nil)))
	l = append(l, nif)

	// if to.ptr != frm.ptr { memmove( ... ) }
	ne := pstate.nod(OIF, pstate.nod(ONE, nto, nfrm), nil)
	ne.SetLikely(true)
	l = append(l, ne)

	fn := pstate.syslook("memmove")
	fn = pstate.substArgTypes(fn, nl.Type.Elem(pstate.types), nl.Type.Elem(pstate.types))
	nwid := pstate.temp(pstate.types.Types[TUINTPTR])
	setwid := pstate.nod(OAS, nwid, pstate.conv(nlen, pstate.types.Types[TUINTPTR]))
	ne.Nbody.Append(setwid)
	nwid = pstate.nod(OMUL, nwid, pstate.nodintconst(nl.Type.Elem(pstate.types).Width))
	call := pstate.mkcall1(fn, nil, init, nto, nfrm, nwid)
	ne.Nbody.Append(call)

	pstate.typecheckslice(l, Etop)
	pstate.walkstmtlist(l)
	init.Append(l...)
	return nlen
}

func (pstate *PackageState) eqfor(t *types.Type) (n *Node, needsize bool) {
	// Should only arrive here with large memory or
	// a struct/array containing a non-memory field/element.
	// Small memory is handled inline, and single non-memory
	// is handled during type check (OCMPSTR etc).
	switch a, _ := pstate.algtype1(t); a {
	case AMEM:
		n := pstate.syslook("memequal")
		n = pstate.substArgTypes(n, t, t)
		return n, true
	case ASPECIAL:
		sym := pstate.typesymprefix(".eq", t)
		n := pstate.newname(sym)
		n.SetClass(PFUNC)
		n.Type = pstate.functype(nil, []*Node{
			pstate.anonfield(pstate.types.NewPtr(t)),
			pstate.anonfield(pstate.types.NewPtr(t)),
		}, []*Node{
			pstate.anonfield(pstate.types.Types[TBOOL]),
		})
		return n, false
	}
	pstate.Fatalf("eqfor %v", t)
	return nil, false
}

// The result of walkcompare MUST be assigned back to n, e.g.
// 	n.Left = walkcompare(n.Left, init)
func (pstate *PackageState) walkcompare(n *Node, init *Nodes) *Node {
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
		// Handle both == and !=.
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
		tab := pstate.nod(OITAB, l, nil)
		rtyp := pstate.typename(r.Type)
		if l.Type.IsEmptyInterface(pstate.types) {
			tab.Type = pstate.types.NewPtr(pstate.types.Types[TUINT8])
			tab.SetTypecheck(1)
			eqtype = pstate.nod(eq, tab, rtyp)
		} else {
			nonnil := pstate.nod(pstate.brcom(eq), pstate.nodnil(), tab)
			match := pstate.nod(eq, pstate.itabType(tab), rtyp)
			eqtype = pstate.nod(andor, nonnil, match)
		}
		// Check for data equal.
		eqdata := pstate.nod(eq, pstate.ifaceData(l, r.Type), r)
		// Put it all together.
		expr := pstate.nod(andor, eqtype, eqdata)
		n = pstate.finishcompare(n, expr, init)
		return n
	}

	// Must be comparison of array or struct.
	// Otherwise back end handles it.
	// While we're here, decide whether to
	// inline or call an eq alg.
	t := n.Left.Type
	var inline bool

	maxcmpsize := int64(4)
	unalignedLoad := pstate.canMergeLoads()
	if unalignedLoad {
		// Keep this low enough to generate less code than a function call.
		maxcmpsize = 2 * int64(pstate.thearch.LinkArch.RegSize)
	}

	switch t.Etype {
	default:
		return n
	case TARRAY:
		// We can compare several elements at once with 2/4/8 byte integer compares
		inline = t.NumElem(pstate.types) <= 1 || (pstate.issimple[t.Elem(pstate.types).Etype] && (t.NumElem(pstate.types) <= 4 || t.Elem(pstate.types).Width*t.NumElem(pstate.types) <= maxcmpsize))
	case TSTRUCT:
		inline = t.NumComponents(pstate.types, types.IgnoreBlankFields) <= 4
	}

	cmpl := n.Left
	for cmpl != nil && cmpl.Op == OCONVNOP {
		cmpl = cmpl.Left
	}
	cmpr := n.Right
	for cmpr != nil && cmpr.Op == OCONVNOP {
		cmpr = cmpr.Left
	}

	// Chose not to inline. Call equality function directly.
	if !inline {
		if isvaluelit(cmpl) {
			var_ := pstate.temp(cmpl.Type)
			pstate.anylit(cmpl, var_, init)
			cmpl = var_
		}
		if isvaluelit(cmpr) {
			var_ := pstate.temp(cmpr.Type)
			pstate.anylit(cmpr, var_, init)
			cmpr = var_
		}
		if !islvalue(cmpl) || !islvalue(cmpr) {
			pstate.Fatalf("arguments of comparison must be lvalues - %v %v", cmpl, cmpr)
		}

		// eq algs take pointers
		pl := pstate.temp(pstate.types.NewPtr(t))
		al := pstate.nod(OAS, pl, pstate.nod(OADDR, cmpl, nil))
		al = pstate.typecheck(al, Etop)
		init.Append(al)

		pr := pstate.temp(pstate.types.NewPtr(t))
		ar := pstate.nod(OAS, pr, pstate.nod(OADDR, cmpr, nil))
		ar = pstate.typecheck(ar, Etop)
		init.Append(ar)

		fn, needsize := pstate.eqfor(t)
		call := pstate.nod(OCALL, fn, nil)
		call.List.Append(pl)
		call.List.Append(pr)
		if needsize {
			call.List.Append(pstate.nodintconst(t.Width))
		}
		res := call
		if n.Op != OEQ {
			res = pstate.nod(ONOT, res, nil)
		}
		n = pstate.finishcompare(n, res, init)
		return n
	}

	// inline: build boolean expression comparing element by element
	andor := OANDAND
	if n.Op == ONE {
		andor = OOROR
	}
	var expr *Node
	compare := func(el, er *Node) {
		a := pstate.nod(n.Op, el, er)
		if expr == nil {
			expr = a
		} else {
			expr = pstate.nod(andor, expr, a)
		}
	}
	cmpl = pstate.safeexpr(cmpl, init)
	cmpr = pstate.safeexpr(cmpr, init)
	if t.IsStruct() {
		for _, f := range t.Fields(pstate.types).Slice() {
			sym := f.Sym
			if sym.IsBlank() {
				continue
			}
			compare(
				pstate.nodSym(OXDOT, cmpl, sym),
				pstate.nodSym(OXDOT, cmpr, sym),
			)
		}
	} else {
		step := int64(1)
		remains := t.NumElem(pstate.types) * t.Elem(pstate.types).Width
		combine64bit := unalignedLoad && pstate.Widthreg == 8 && t.Elem(pstate.types).Width <= 4 && t.Elem(pstate.types).IsInteger()
		combine32bit := unalignedLoad && t.Elem(pstate.types).Width <= 2 && t.Elem(pstate.types).IsInteger()
		combine16bit := unalignedLoad && t.Elem(pstate.types).Width == 1 && t.Elem(pstate.types).IsInteger()
		for i := int64(0); remains > 0; {
			var convType *types.Type
			switch {
			case remains >= 8 && combine64bit:
				convType = pstate.types.Types[TINT64]
				step = 8 / t.Elem(pstate.types).Width
			case remains >= 4 && combine32bit:
				convType = pstate.types.Types[TUINT32]
				step = 4 / t.Elem(pstate.types).Width
			case remains >= 2 && combine16bit:
				convType = pstate.types.Types[TUINT16]
				step = 2 / t.Elem(pstate.types).Width
			default:
				step = 1
			}
			if step == 1 {
				compare(
					pstate.nod(OINDEX, cmpl, pstate.nodintconst(i)),
					pstate.nod(OINDEX, cmpr, pstate.nodintconst(i)),
				)
				i++
				remains -= t.Elem(pstate.types).Width
			} else {
				elemType := t.Elem(pstate.types).ToUnsigned(pstate.types)
				cmplw := pstate.nod(OINDEX, cmpl, pstate.nodintconst(i))
				cmplw = pstate.conv(cmplw, elemType) // convert to unsigned
				cmplw = pstate.conv(cmplw, convType) // widen
				cmprw := pstate.nod(OINDEX, cmpr, pstate.nodintconst(i))
				cmprw = pstate.conv(cmprw, elemType)
				cmprw = pstate.conv(cmprw, convType)
				// For code like this:  uint32(s[0]) | uint32(s[1])<<8 | uint32(s[2])<<16 ...
				// ssa will generate a single large load.
				for offset := int64(1); offset < step; offset++ {
					lb := pstate.nod(OINDEX, cmpl, pstate.nodintconst(i+offset))
					lb = pstate.conv(lb, elemType)
					lb = pstate.conv(lb, convType)
					lb = pstate.nod(OLSH, lb, pstate.nodintconst(8*t.Elem(pstate.types).Width*offset))
					cmplw = pstate.nod(OOR, cmplw, lb)
					rb := pstate.nod(OINDEX, cmpr, pstate.nodintconst(i+offset))
					rb = pstate.conv(rb, elemType)
					rb = pstate.conv(rb, convType)
					rb = pstate.nod(OLSH, rb, pstate.nodintconst(8*t.Elem(pstate.types).Width*offset))
					cmprw = pstate.nod(OOR, cmprw, rb)
				}
				compare(cmplw, cmprw)
				i += step
				remains -= step * t.Elem(pstate.types).Width
			}
		}
	}
	if expr == nil {
		expr = pstate.nodbool(n.Op == OEQ)
	}
	n = pstate.finishcompare(n, expr, init)
	return n
}

// The result of finishcompare MUST be assigned back to n, e.g.
// 	n.Left = finishcompare(n.Left, x, r, init)
func (pstate *PackageState) finishcompare(n, r *Node, init *Nodes) *Node {
	r = pstate.typecheck(r, Erv)
	r = pstate.conv(r, n.Type)
	r = pstate.walkexpr(r, init)
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
func (pstate *PackageState) walkinrange(n *Node, init *Nodes) *Node {
	// We are looking for something equivalent to a opl b OP b opr c, where:
	// * a, b, and c have integer type
	// * b is side-effect-free
	// * opl and opr are each < or ≤
	// * OP is &&
	l := n.Left
	r := n.Right
	if !l.isIntOrdering() || !r.isIntOrdering() {
		return n
	}

	// Find b, if it exists, and rename appropriately.
	// Input is: l.Left l.Op l.Right ANDAND/OROR r.Left r.Op r.Right
	// Output is: a opl b(==x) ANDAND/OROR b(==x) opr c
	a, opl, b := l.Left, l.Op, l.Right
	x, opr, c := r.Left, r.Op, r.Right
	for i := 0; ; i++ {
		if pstate.samesafeexpr(b, x) {
			break
		}
		if i == 3 {
			// Tried all permutations and couldn't find an appropriate b == x.
			return n
		}
		if i&1 == 0 {
			a, opl, b = b, pstate.brrev(opl), a
		} else {
			x, opr, c = c, pstate.brrev(opr), x
		}
	}

	// If n.Op is ||, apply de Morgan.
	// Negate the internal ops now; we'll negate the top level op at the end.
	// Henceforth assume &&.
	negateResult := n.Op == OOROR
	if negateResult {
		opl = pstate.brcom(opl)
		opr = pstate.brcom(opr)
	}

	cmpdir := func(o Op) int {
		switch o {
		case OLE, OLT:
			return -1
		case OGE, OGT:
			return +1
		}
		pstate.Fatalf("walkinrange cmpdir %v", o)
		return 0
	}
	if cmpdir(opl) != cmpdir(opr) {
		// Not a range check; something like b < a && b < c.
		return n
	}

	switch opl {
	case OGE, OGT:
		// We have something like a > b && b ≥ c.
		// Switch and reverse ops and rename constants,
		// to make it look like a ≤ b && b < c.
		a, c = c, a
		opl, opr = pstate.brrev(opr), pstate.brrev(opl)
	}

	// We must ensure that c-a is non-negative.
	// For now, require a and c to be constants.
	// In the future, we could also support a == 0 and c == len/cap(...).
	// Unfortunately, by this point, most len/cap expressions have been
	// stored into temporary variables.
	if !pstate.Isconst(a, CTINT) || !pstate.Isconst(c, CTINT) {
		return n
	}

	if opl == OLT {
		// We have a < b && ...
		// We need a ≤ b && ... to safely use unsigned comparison tricks.
		// If a is not the maximum constant for b's type,
		// we can increment a and switch to ≤.
		if a.Int64(pstate) >= pstate.maxintval[b.Type.Etype].Int64(pstate) {
			return n
		}
		a = pstate.nodintconst(a.Int64(pstate) + 1)
		opl = OLE
	}

	bound := c.Int64(pstate) - a.Int64(pstate)
	if bound < 0 {
		// Bad news. Something like 5 <= x && x < 3.
		// Rare in practice, and we still need to generate side-effects,
		// so just leave it alone.
		return n
	}

	// We have a ≤ b && b < c (or a ≤ b && b ≤ c).
	// This is equivalent to (a-a) ≤ (b-a) && (b-a) < (c-a),
	// which is equivalent to 0 ≤ (b-a) && (b-a) < (c-a),
	// which is equivalent to uint(b-a) < uint(c-a).
	ut := b.Type.ToUnsigned(pstate.types)
	lhs := pstate.conv(pstate.nod(OSUB, b, a), ut)
	rhs := pstate.nodintconst(bound)
	if negateResult {
		// Negate top level.
		opr = pstate.brcom(opr)
	}
	cmp := pstate.nod(opr, lhs, rhs)
	cmp.Pos = n.Pos
	cmp = pstate.addinit(cmp, l.Ninit.Slice())
	cmp = pstate.addinit(cmp, r.Ninit.Slice())
	// Typecheck the AST rooted at cmp...
	cmp = pstate.typecheck(cmp, Erv)
	// ...but then reset cmp's type to match n's type.
	cmp.Type = n.Type
	cmp = pstate.walkexpr(cmp, init)
	return cmp
}

// return 1 if integer n must be in range [0, max), 0 otherwise
func (pstate *PackageState) bounded(n *Node, max int64) bool {
	if n.Type == nil || !n.Type.IsInteger() {
		return false
	}

	sign := n.Type.IsSigned()
	bits := int32(8 * n.Type.Width)

	if pstate.smallintconst(n) {
		v := n.Int64(pstate)
		return 0 <= v && v < max
	}

	switch n.Op {
	case OAND:
		v := int64(-1)
		if pstate.smallintconst(n.Left) {
			v = n.Left.Int64(pstate)
		} else if pstate.smallintconst(n.Right) {
			v = n.Right.Int64(pstate)
		}

		if 0 <= v && v < max {
			return true
		}

	case OMOD:
		if !sign && pstate.smallintconst(n.Right) {
			v := n.Right.Int64(pstate)
			if 0 <= v && v <= max {
				return true
			}
		}

	case ODIV:
		if !sign && pstate.smallintconst(n.Right) {
			v := n.Right.Int64(pstate)
			for bits > 0 && v >= 2 {
				bits--
				v >>= 1
			}
		}

	case ORSH:
		if !sign && pstate.smallintconst(n.Right) {
			v := n.Right.Int64(pstate)
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
func (pstate *PackageState) usemethod(n *Node) {
	t := n.Left.Type

	// Looking for either of:
	//	Method(int) reflect.Method
	//	MethodByName(string) (reflect.Method, bool)
	//
	// TODO(crawshaw): improve precision of match by working out
	//                 how to check the method name.
	if n := t.NumParams(pstate.types); n != 1 {
		return
	}
	if n := t.NumResults(pstate.types); n != 1 && n != 2 {
		return
	}
	p0 := t.Params(pstate.types).Field(pstate.types, 0)
	res0 := t.Results(pstate.types).Field(pstate.types, 0)
	var res1 *types.Field
	if t.NumResults(pstate.types) == 2 {
		res1 = t.Results(pstate.types).Field(pstate.types, 1)
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

	// Note: Don't rely on res0.Type.String() since its formatting depends on multiple factors
	//       (including global variables such as numImports - was issue #19028).
	if s := res0.Type.Sym; s != nil && s.Name == "Method" && s.Pkg != nil && s.Pkg.Path == "reflect" {
		pstate.Curfn.Func.SetReflectMethod(true)
	}
}

func (pstate *PackageState) usefield(n *Node) {
	if pstate.objabi.Fieldtrack_enabled == 0 {
		return
	}

	switch n.Op {
	default:
		pstate.Fatalf("usefield %v", n.Op)

	case ODOT, ODOTPTR:
		break
	}
	if n.Sym == nil {
		// No field name.  This DOTPTR was built by the compiler for access
		// to runtime data structures.  Ignore.
		return
	}

	t := n.Left.Type
	if t.IsPtr() {
		t = t.Elem(pstate.types)
	}
	field := pstate.dotField[typeSymKey{t.Orig, n.Sym}]
	if field == nil {
		pstate.Fatalf("usefield %v %v without paramfld", n.Left.Type, n.Sym)
	}
	if !strings.Contains(field.Note, "go:\"track\"") {
		return
	}

	outer := n.Left.Type
	if outer.IsPtr() {
		outer = outer.Elem(pstate.types)
	}
	if outer.Sym == nil {
		pstate.yyerror("tracked field must be in named struct type")
	}
	if !types.IsExported(field.Sym.Name) {
		pstate.yyerror("tracked field must be exported (upper case)")
	}

	sym := pstate.tracksym(outer, field)
	if pstate.Curfn.Func.FieldTrack == nil {
		pstate.Curfn.Func.FieldTrack = make(map[*types.Sym]struct{})
	}
	pstate.Curfn.Func.FieldTrack[sym] = struct{}{}
}

func (pstate *PackageState) candiscardlist(l Nodes) bool {
	for _, n := range l.Slice() {
		if !pstate.candiscard(n) {
			return false
		}
	}
	return true
}

func (pstate *PackageState) candiscard(n *Node) bool {
	if n == nil {
		return true
	}

	switch n.Op {
	default:
		return false

	// Discardable as long as the subpieces are.
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

	// Discardable as long as we know it's not division by zero.
	case ODIV, OMOD:
		if pstate.Isconst(n.Right, CTINT) && n.Right.Val().U.(*Mpint).CmpInt64(0) != 0 {
			break
		}
		if pstate.Isconst(n.Right, CTFLT) && n.Right.Val().U.(*Mpflt).CmpFloat64(0) != 0 {
			break
		}
		return false

	// Discardable as long as we know it won't fail because of a bad size.
	case OMAKECHAN, OMAKEMAP:
		if pstate.Isconst(n.Left, CTINT) && n.Left.Val().U.(*Mpint).CmpInt64(0) == 0 {
			break
		}
		return false

	// Difficult to tell what sizes are okay.
	case OMAKESLICE:
		return false
	}

	if !pstate.candiscard(n.Left) || !pstate.candiscard(n.Right) || !pstate.candiscardlist(n.Ninit) || !pstate.candiscardlist(n.Nbody) || !pstate.candiscardlist(n.List) || !pstate.candiscardlist(n.Rlist) {
		return false
	}

	return true
}

// The result of wrapCall MUST be assigned back to n, e.g.
// 	n.Left = wrapCall(n.Left, init)
func (pstate *PackageState) wrapCall(n *Node, init *Nodes) *Node {
	if n.Ninit.Len() != 0 {
		pstate.walkstmtlist(n.Ninit.Slice())
		init.AppendNodes(&n.Ninit)
	}

	t := pstate.nod(OTFUNC, nil, nil)
	for i, arg := range n.List.Slice() {
		s := pstate.lookupN("a", i)
		t.List.Append(pstate.symfield(s, arg.Type))
	}

	pstate.wrapCall_prgen++
	sym := pstate.lookupN("wrap·", pstate.wrapCall_prgen)
	fn := pstate.dclfunc(sym, t)

	a := pstate.nod(n.Op, nil, nil)
	a.List.Set(pstate.paramNnames(t.Type))
	a = pstate.typecheck(a, Etop)
	fn.Nbody.Set1(a)

	pstate.funcbody()

	fn = pstate.typecheck(fn, Etop)
	pstate.typecheckslice(fn.Nbody.Slice(), Etop)
	pstate.xtop = append(pstate.xtop, fn)

	a = pstate.nod(OCALL, nil, nil)
	a.Left = fn.Func.Nname
	a.List.Set(n.List.Slice())
	a = pstate.typecheck(a, Etop)
	a = pstate.walkexpr(a, init)
	return a
}

// substArgTypes substitutes the given list of types for
// successive occurrences of the "any" placeholder in the
// type syntax expression n.Type.
// The result of substArgTypes MUST be assigned back to old, e.g.
// 	n.Left = substArgTypes(n.Left, t1, t2)
func (pstate *PackageState) substArgTypes(old *Node, types_ ...*types.Type) *Node {
	n := old.copy() // make shallow copy

	for _, t := range types_ {
		pstate.dowidth(t)
	}
	n.Type = pstate.types.SubstAny(n.Type, &types_)
	if len(types_) > 0 {
		pstate.Fatalf("substArgTypes: too many argument types")
	}
	return n
}

// canMergeLoads reports whether the backend optimization passes for
// the current architecture can combine adjacent loads into a single
// larger, possibly unaligned, load. Note that currently the
// optimizations must be able to handle little endian byte order.
func (pstate *PackageState) canMergeLoads() bool {
	switch pstate.thearch.LinkArch.Family {
	case sys.ARM64, sys.AMD64, sys.I386, sys.S390X:
		return true
	case sys.PPC64:
		// Load combining only supported on ppc64le.
		return pstate.thearch.LinkArch.ByteOrder == binary.LittleEndian
	}
	return false
}

// isRuneCount reports whether n is of the form len([]rune(string)).
// These are optimized into a call to runtime.countrunes.
func (pstate *PackageState) isRuneCount(n *Node) bool {
	return pstate.Debug['N'] == 0 && !pstate.instrumenting && n.Op == OLEN && n.Left.Op == OSTRARRAYRUNE
}

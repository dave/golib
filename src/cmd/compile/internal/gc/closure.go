package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/syntax"
	"github.com/dave/golib/src/cmd/compile/internal/types"
)

func (p *noder) funcLit(psess *PackageSession, expr *syntax.FuncLit) *Node {
	xtype := p.typeExpr(psess, expr.Type)
	ntype := p.typeExpr(psess, expr.Type)

	xfunc := p.nod(psess, expr, ODCLFUNC, nil, nil)
	xfunc.Func.SetIsHiddenClosure(psess.Curfn != nil)
	xfunc.Func.Nname = p.setlineno(psess, expr, psess.newfuncname(psess.nblank.Sym))
	xfunc.Func.Nname.Name.Param.Ntype = xtype
	xfunc.Func.Nname.Name.Defn = xfunc

	clo := p.nod(psess, expr, OCLOSURE, nil, nil)
	clo.Func.Ntype = ntype

	xfunc.Func.Closure = clo
	clo.Func.Closure = xfunc

	p.funcBody(psess, xfunc, expr.Body)

	for _, v := range xfunc.Func.Cvars.Slice() {

		v1 := v.Name.Defn
		v1.Name.Param.Innermost = v.Name.Param.Outer

		v.Name.Param.Outer = psess.oldname(v.Sym)
	}

	return clo
}

func (psess *PackageSession) typecheckclosure(clo *Node, top int) {
	xfunc := clo.Func.Closure

	for _, ln := range xfunc.Func.Cvars.Slice() {
		n := ln.Name.Defn
		if !n.Name.Captured() {
			n.Name.SetCaptured(true)
			if n.Name.Decldepth == 0 {
				psess.
					Fatalf("typecheckclosure: var %S does not have decldepth assigned", n)
			}

			if n.Name.Decldepth == psess.decldepth {
				n.SetAssigned(false)
			}
		}
	}

	xfunc.Func.Nname.Sym = psess.closurename(psess.Curfn)
	disableExport(xfunc.Func.Nname.Sym)
	psess.
		declare(xfunc.Func.Nname, PFUNC)
	xfunc = psess.typecheck(xfunc, Etop)

	clo.Func.Ntype = psess.typecheck(clo.Func.Ntype, Etype)
	clo.Type = clo.Func.Ntype.Type
	clo.Func.Top = top

	if psess.Curfn != nil && clo.Type != nil {
		oldfn := psess.Curfn
		psess.
			Curfn = xfunc
		olddd := psess.decldepth
		psess.
			decldepth = 1
		psess.
			typecheckslice(xfunc.Nbody.Slice(), Etop)
		psess.
			decldepth = olddd
		psess.
			Curfn = oldfn
	}
	psess.
		xtop = append(psess.xtop, xfunc)
}

// globClosgen is like Func.Closgen, but for the global scope.

// closurename generates a new unique name for a closure within
// outerfunc.
func (psess *PackageSession) closurename(outerfunc *Node) *types.Sym {
	outer := "glob."
	prefix := "func"
	gen := &psess.globClosgen

	if outerfunc != nil {
		if outerfunc.Func.Closure != nil {
			prefix = ""
		}

		outer = outerfunc.funcname()

		if !outerfunc.Func.Nname.isBlank() {
			gen = &outerfunc.Func.Closgen
		}
	}

	*gen++
	return psess.lookup(fmt.Sprintf("%s.%s%d", outer, prefix, *gen))
}

// capturevarscomplete is set to true when the capturevars phase is done.

// capturevars is called in a separate phase after all typechecking is done.
// It decides whether each variable captured by a closure should be captured
// by value or by reference.
// We use value capturing for values <= 128 bytes that are never reassigned
// after capturing (effectively constant).
func (psess *PackageSession) capturevars(xfunc *Node) {
	lno := psess.lineno
	psess.
		lineno = xfunc.Pos

	clo := xfunc.Func.Closure
	cvars := xfunc.Func.Cvars.Slice()
	out := cvars[:0]
	for _, v := range cvars {
		if v.Type == nil {

			continue
		}
		out = append(out, v)
		psess.
			dowidth(v.Type)

		outer := v.Name.Param.Outer
		outermost := v.Name.Defn

		if outer.Class() != PPARAMOUT && !outermost.Addrtaken() && !outermost.Assigned() && v.Type.Width <= 128 {
			v.Name.SetByval(true)
		} else {
			outermost.SetAddrtaken(true)
			outer = psess.nod(OADDR, outer, nil)
		}

		if psess.Debug['m'] > 1 {
			var name *types.Sym
			if v.Name.Curfn != nil && v.Name.Curfn.Func.Nname != nil {
				name = v.Name.Curfn.Func.Nname.Sym
			}
			how := "ref"
			if v.Name.Byval() {
				how = "value"
			}
			psess.
				Warnl(v.Pos, "%v capturing by %s: %v (addr=%v assign=%v width=%d)", name, how, v.Sym, outermost.Addrtaken(), outermost.Assigned(), int32(v.Type.Width))
		}

		outer = psess.typecheck(outer, Erv)
		clo.Func.Enter.Append(outer)
	}

	xfunc.Func.Cvars.Set(out)
	psess.
		lineno = lno
}

// transformclosure is called in a separate phase after escape analysis.
// It transform closure bodies to properly reference captured variables.
func (psess *PackageSession) transformclosure(xfunc *Node) {
	lno := psess.lineno
	psess.
		lineno = xfunc.Pos
	clo := xfunc.Func.Closure

	if clo.Func.Top&Ecall != 0 {

		f := xfunc.Func.Nname

		// We are going to insert captured variables before input args.
		var params []*types.Field
		var decls []*Node
		for _, v := range xfunc.Func.Cvars.Slice() {
			if !v.Name.Byval() {

				addr := psess.newname(psess.lookup("&" + v.Sym.Name))
				addr.Type = psess.types.NewPtr(v.Type)
				v.Name.Param.Heapaddr = addr
				v = addr
			}

			v.SetClass(PPARAM)
			decls = append(decls, v)

			fld := types.NewField()
			fld.Nname = asTypesNode(v)
			fld.Type = v.Type
			fld.Sym = v.Sym
			params = append(params, fld)
		}

		if len(params) > 0 {

			f.Type.Params(psess.types).SetFields(psess.types, append(params, f.Type.Params(psess.types).FieldSlice(psess.types)...))
			xfunc.Func.Dcl = append(decls, xfunc.Func.Dcl...)
		}
		psess.
			dowidth(f.Type)
		xfunc.Type = f.Type
	} else {
		// The closure is not called, so it is going to stay as closure.
		var body []*Node
		offset := int64(psess.Widthptr)
		for _, v := range xfunc.Func.Cvars.Slice() {

			cv := psess.nod(OCLOSUREVAR, nil, nil)

			cv.Type = v.Type
			if !v.Name.Byval() {
				cv.Type = psess.types.NewPtr(v.Type)
			}
			offset = psess.Rnd(offset, int64(cv.Type.Align))
			cv.Xoffset = offset
			offset += cv.Type.Width

			if v.Name.Byval() && v.Type.Width <= int64(2*psess.Widthptr) {

				v.SetClass(PAUTO)
				xfunc.Func.Dcl = append(xfunc.Func.Dcl, v)
				body = append(body, psess.nod(OAS, v, cv))
			} else {

				addr := psess.newname(psess.lookup("&" + v.Sym.Name))
				addr.Type = psess.types.NewPtr(v.Type)
				addr.SetClass(PAUTO)
				addr.Name.SetUsed(true)
				addr.Name.Curfn = xfunc
				xfunc.Func.Dcl = append(xfunc.Func.Dcl, addr)
				v.Name.Param.Heapaddr = addr
				if v.Name.Byval() {
					cv = psess.nod(OADDR, cv, nil)
				}
				body = append(body, psess.nod(OAS, addr, cv))
			}
		}

		if len(body) > 0 {
			psess.
				typecheckslice(body, Etop)
			xfunc.Func.Enter.Set(body)
			xfunc.Func.SetNeedctxt(true)
		}
	}
	psess.
		lineno = lno
}

// hasemptycvars returns true iff closure clo has an
// empty list of captured vars.
func hasemptycvars(clo *Node) bool {
	xfunc := clo.Func.Closure
	return xfunc.Func.Cvars.Len() == 0
}

// closuredebugruntimecheck applies boilerplate checks for debug flags
// and compiling runtime
func (psess *PackageSession) closuredebugruntimecheck(clo *Node) {
	if psess.Debug_closure > 0 {
		xfunc := clo.Func.Closure
		if clo.Esc == EscHeap {
			psess.
				Warnl(clo.Pos, "heap closure, captured vars = %v", xfunc.Func.Cvars)
		} else {
			psess.
				Warnl(clo.Pos, "stack closure, captured vars = %v", xfunc.Func.Cvars)
		}
	}
	if psess.compiling_runtime && clo.Esc == EscHeap {
		psess.
			yyerrorl(clo.Pos, "heap-allocated closure, not allowed in runtime.")
	}
}

func (psess *PackageSession) walkclosure(clo *Node, init *Nodes) *Node {
	xfunc := clo.Func.Closure

	if hasemptycvars(clo) {
		if psess.Debug_closure > 0 {
			psess.
				Warnl(clo.Pos, "closure converted to global")
		}
		return xfunc.Func.Nname
	}
	psess.
		closuredebugruntimecheck(clo)

	fields := []*Node{psess.
		namedfield(".F", psess.types.Types[TUINTPTR]),
	}
	for _, v := range xfunc.Func.Cvars.Slice() {
		typ := v.Type
		if !v.Name.Byval() {
			typ = psess.types.NewPtr(typ)
		}
		fields = append(fields, psess.symfield(v.Sym, typ))
	}
	typ := psess.tostruct(fields)
	typ.SetNoalg(true)

	clos := psess.nod(OCOMPLIT, nil, psess.nod(OIND, psess.typenod(typ), nil))
	clos.Esc = clo.Esc
	clos.Right.SetImplicit(true)
	clos.List.Set(append([]*Node{psess.nod(OCFUNC, xfunc.Func.Nname, nil)}, clo.Func.Enter.Slice()...))

	clos = psess.nod(OCONVNOP, clos, nil)
	clos.Type = clo.Type

	clos = psess.typecheck(clos, Erv)

	clos.Left.Esc = clo.Esc

	if x := psess.prealloc[clo]; x != nil {
		x.Type = clos.Left.Left.Type
		x.Orig.Type = x.Type
		clos.Left.Right = x
		delete(psess.prealloc, clo)
	}

	return psess.walkexpr(clos, init)
}

func (psess *PackageSession) typecheckpartialcall(fn *Node, sym *types.Sym) {
	switch fn.Op {
	case ODOTINTER, ODOTMETH:
		break

	default:
		psess.
			Fatalf("invalid typecheckpartialcall")
	}

	xfunc := psess.makepartialcall(fn, fn.Type, sym)
	fn.Func = xfunc.Func
	fn.Right = psess.newname(sym)
	fn.Op = OCALLPART
	fn.Type = xfunc.Type
}

func (psess *PackageSession) makepartialcall(fn *Node, t0 *types.Type, meth *types.Sym) *Node {
	rcvrtype := fn.Left.Type
	sym := psess.methodSymSuffix(rcvrtype, meth, "-fm")

	if sym.Uniq() {
		return asNode(sym.Def)
	}
	sym.SetUniq(true)

	savecurfn := psess.Curfn
	psess.
		Curfn = nil

	tfn := psess.nod(OTFUNC, nil, nil)
	tfn.List.Set(psess.structargs(t0.Params(psess.types), true))
	tfn.Rlist.Set(psess.structargs(t0.Results(psess.types), false))

	disableExport(sym)
	xfunc := psess.dclfunc(sym, tfn)
	xfunc.Func.SetDupok(true)
	xfunc.Func.SetNeedctxt(true)

	tfn.Type.SetPkg(psess.types, t0.Pkg(psess.types))

	cv := psess.nod(OCLOSUREVAR, nil, nil)
	cv.Type = rcvrtype
	cv.Xoffset = psess.Rnd(int64(psess.Widthptr), int64(cv.Type.Align))

	ptr := psess.newname(psess.lookup(".this"))
	psess.
		declare(ptr, PAUTO)
	ptr.Name.SetUsed(true)
	var body []*Node
	if rcvrtype.IsPtr() || rcvrtype.IsInterface() {
		ptr.Type = rcvrtype
		body = append(body, psess.nod(OAS, ptr, cv))
	} else {
		ptr.Type = psess.types.NewPtr(rcvrtype)
		body = append(body, psess.nod(OAS, ptr, psess.nod(OADDR, cv, nil)))
	}

	call := psess.nod(OCALL, psess.nodSym(OXDOT, ptr, meth), nil)
	call.List.Set(psess.paramNnames(tfn.Type))
	call.SetIsddd(tfn.Type.IsVariadic(psess.types))
	if t0.NumResults(psess.types) != 0 {
		n := psess.nod(ORETURN, nil, nil)
		n.List.Set1(call)
		call = n
	}
	body = append(body, call)

	xfunc.Nbody.Set(body)
	psess.
		funcbody()

	xfunc = psess.typecheck(xfunc, Etop)
	sym.Def = asTypesNode(xfunc)
	psess.
		xtop = append(psess.xtop, xfunc)
	psess.
		Curfn = savecurfn

	return xfunc
}

func (psess *PackageSession) walkpartialcall(n *Node, init *Nodes) *Node {

	if n.Left.Type.IsInterface() {

		n.Left = psess.cheapexpr(n.Left, init)
		psess.
			checknil(n.Left, init)
	}

	typ := psess.tostruct([]*Node{psess.
		namedfield("F", psess.types.Types[TUINTPTR]), psess.
		namedfield("R", n.Left.Type),
	})
	typ.SetNoalg(true)

	clos := psess.nod(OCOMPLIT, nil, psess.nod(OIND, psess.typenod(typ), nil))
	clos.Esc = n.Esc
	clos.Right.SetImplicit(true)
	clos.List.Set1(psess.nod(OCFUNC, n.Func.Nname, nil))
	clos.List.Append(n.Left)

	clos = psess.nod(OCONVNOP, clos, nil)
	clos.Type = n.Type

	clos = psess.typecheck(clos, Erv)

	clos.Left.Esc = n.Esc

	if x := psess.prealloc[n]; x != nil {
		x.Type = clos.Left.Left.Type
		x.Orig.Type = x.Type
		clos.Left.Right = x
		delete(psess.prealloc, n)
	}

	return psess.walkexpr(clos, init)
}

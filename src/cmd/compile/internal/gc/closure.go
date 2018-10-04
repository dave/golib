// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/syntax"
	"github.com/dave/golib/src/cmd/compile/internal/types"
)

func (p *noder) funcLit(pstate *PackageState, expr *syntax.FuncLit) *Node {
	xtype := p.typeExpr(pstate, expr.Type)
	ntype := p.typeExpr(pstate, expr.Type)

	xfunc := p.nod(pstate, expr, ODCLFUNC, nil, nil)
	xfunc.Func.SetIsHiddenClosure(pstate.Curfn != nil)
	xfunc.Func.Nname = p.setlineno(pstate, expr, pstate.newfuncname(pstate.nblank.Sym)) // filled in by typecheckclosure
	xfunc.Func.Nname.Name.Param.Ntype = xtype
	xfunc.Func.Nname.Name.Defn = xfunc

	clo := p.nod(pstate, expr, OCLOSURE, nil, nil)
	clo.Func.Ntype = ntype

	xfunc.Func.Closure = clo
	clo.Func.Closure = xfunc

	p.funcBody(pstate, xfunc, expr.Body)

	// closure-specific variables are hanging off the
	// ordinary ones in the symbol table; see oldname.
	// unhook them.
	// make the list of pointers for the closure call.
	for _, v := range xfunc.Func.Cvars.Slice() {
		// Unlink from v1; see comment in syntax.go type Param for these fields.
		v1 := v.Name.Defn
		v1.Name.Param.Innermost = v.Name.Param.Outer

		// If the closure usage of v is not dense,
		// we need to make it dense; now that we're out
		// of the function in which v appeared,
		// look up v.Sym in the enclosing function
		// and keep it around for use in the compiled code.
		//
		// That is, suppose we just finished parsing the innermost
		// closure f4 in this code:
		//
		//	func f() {
		//		v := 1
		//		func() { // f2
		//			use(v)
		//			func() { // f3
		//				func() { // f4
		//					use(v)
		//				}()
		//			}()
		//		}()
		//	}
		//
		// At this point v.Outer is f2's v; there is no f3's v.
		// To construct the closure f4 from within f3,
		// we need to use f3's v and in this case we need to create f3's v.
		// We are now in the context of f3, so calling oldname(v.Sym)
		// obtains f3's v, creating it if necessary (as it is in the example).
		//
		// capturevars will decide whether to use v directly or &v.
		v.Name.Param.Outer = pstate.oldname(v.Sym)
	}

	return clo
}

func (pstate *PackageState) typecheckclosure(clo *Node, top int) {
	xfunc := clo.Func.Closure

	for _, ln := range xfunc.Func.Cvars.Slice() {
		n := ln.Name.Defn
		if !n.Name.Captured() {
			n.Name.SetCaptured(true)
			if n.Name.Decldepth == 0 {
				pstate.Fatalf("typecheckclosure: var %S does not have decldepth assigned", n)
			}

			// Ignore assignments to the variable in straightline code
			// preceding the first capturing by a closure.
			if n.Name.Decldepth == pstate.decldepth {
				n.SetAssigned(false)
			}
		}
	}

	xfunc.Func.Nname.Sym = pstate.closurename(pstate.Curfn)
	disableExport(xfunc.Func.Nname.Sym)
	pstate.declare(xfunc.Func.Nname, PFUNC)
	xfunc = pstate.typecheck(xfunc, Etop)

	clo.Func.Ntype = pstate.typecheck(clo.Func.Ntype, Etype)
	clo.Type = clo.Func.Ntype.Type
	clo.Func.Top = top

	// Type check the body now, but only if we're inside a function.
	// At top level (in a variable initialization: curfn==nil) we're not
	// ready to type check code yet; we'll check it later, because the
	// underlying closure function we create is added to xtop.
	if pstate.Curfn != nil && clo.Type != nil {
		oldfn := pstate.Curfn
		pstate.Curfn = xfunc
		olddd := pstate.decldepth
		pstate.decldepth = 1
		pstate.typecheckslice(xfunc.Nbody.Slice(), Etop)
		pstate.decldepth = olddd
		pstate.Curfn = oldfn
	}

	pstate.xtop = append(pstate.xtop, xfunc)
}

// closurename generates a new unique name for a closure within
// outerfunc.
func (pstate *PackageState) closurename(outerfunc *Node) *types.Sym {
	outer := "glob."
	prefix := "func"
	gen := &pstate.globClosgen

	if outerfunc != nil {
		if outerfunc.Func.Closure != nil {
			prefix = ""
		}

		outer = outerfunc.funcname()

		// There may be multiple functions named "_". In those
		// cases, we can't use their individual Closgens as it
		// would lead to name clashes.
		if !outerfunc.Func.Nname.isBlank() {
			gen = &outerfunc.Func.Closgen
		}
	}

	*gen++
	return pstate.lookup(fmt.Sprintf("%s.%s%d", outer, prefix, *gen))
}

// capturevars is called in a separate phase after all typechecking is done.
// It decides whether each variable captured by a closure should be captured
// by value or by reference.
// We use value capturing for values <= 128 bytes that are never reassigned
// after capturing (effectively constant).
func (pstate *PackageState) capturevars(xfunc *Node) {
	lno := pstate.lineno
	pstate.lineno = xfunc.Pos

	clo := xfunc.Func.Closure
	cvars := xfunc.Func.Cvars.Slice()
	out := cvars[:0]
	for _, v := range cvars {
		if v.Type == nil {
			// If v.Type is nil, it means v looked like it
			// was going to be used in the closure, but
			// isn't. This happens in struct literals like
			// s{f: x} where we can't distinguish whether
			// f is a field identifier or expression until
			// resolving s.
			continue
		}
		out = append(out, v)

		// type check the & of closed variables outside the closure,
		// so that the outer frame also grabs them and knows they escape.
		pstate.dowidth(v.Type)

		outer := v.Name.Param.Outer
		outermost := v.Name.Defn

		// out parameters will be assigned to implicitly upon return.
		if outer.Class() != PPARAMOUT && !outermost.Addrtaken() && !outermost.Assigned() && v.Type.Width <= 128 {
			v.Name.SetByval(true)
		} else {
			outermost.SetAddrtaken(true)
			outer = pstate.nod(OADDR, outer, nil)
		}

		if pstate.Debug['m'] > 1 {
			var name *types.Sym
			if v.Name.Curfn != nil && v.Name.Curfn.Func.Nname != nil {
				name = v.Name.Curfn.Func.Nname.Sym
			}
			how := "ref"
			if v.Name.Byval() {
				how = "value"
			}
			pstate.Warnl(v.Pos, "%v capturing by %s: %v (addr=%v assign=%v width=%d)", name, how, v.Sym, outermost.Addrtaken(), outermost.Assigned(), int32(v.Type.Width))
		}

		outer = pstate.typecheck(outer, Erv)
		clo.Func.Enter.Append(outer)
	}

	xfunc.Func.Cvars.Set(out)
	pstate.lineno = lno
}

// transformclosure is called in a separate phase after escape analysis.
// It transform closure bodies to properly reference captured variables.
func (pstate *PackageState) transformclosure(xfunc *Node) {
	lno := pstate.lineno
	pstate.lineno = xfunc.Pos
	clo := xfunc.Func.Closure

	if clo.Func.Top&Ecall != 0 {
		// If the closure is directly called, we transform it to a plain function call
		// with variables passed as args. This avoids allocation of a closure object.
		// Here we do only a part of the transformation. Walk of OCALLFUNC(OCLOSURE)
		// will complete the transformation later.
		// For illustration, the following closure:
		//	func(a int) {
		//		println(byval)
		//		byref++
		//	}(42)
		// becomes:
		//	func(byval int, &byref *int, a int) {
		//		println(byval)
		//		(*&byref)++
		//	}(byval, &byref, 42)

		// f is ONAME of the actual function.
		f := xfunc.Func.Nname

		// We are going to insert captured variables before input args.
		var params []*types.Field
		var decls []*Node
		for _, v := range xfunc.Func.Cvars.Slice() {
			if !v.Name.Byval() {
				// If v of type T is captured by reference,
				// we introduce function param &v *T
				// and v remains PAUTOHEAP with &v heapaddr
				// (accesses will implicitly deref &v).
				addr := pstate.newname(pstate.lookup("&" + v.Sym.Name))
				addr.Type = pstate.types.NewPtr(v.Type)
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
			// Prepend params and decls.
			f.Type.Params(pstate.types).SetFields(pstate.types, append(params, f.Type.Params(pstate.types).FieldSlice(pstate.types)...))
			xfunc.Func.Dcl = append(decls, xfunc.Func.Dcl...)
		}

		pstate.dowidth(f.Type)
		xfunc.Type = f.Type // update type of ODCLFUNC
	} else {
		// The closure is not called, so it is going to stay as closure.
		var body []*Node
		offset := int64(pstate.Widthptr)
		for _, v := range xfunc.Func.Cvars.Slice() {
			// cv refers to the field inside of closure OSTRUCTLIT.
			cv := pstate.nod(OCLOSUREVAR, nil, nil)

			cv.Type = v.Type
			if !v.Name.Byval() {
				cv.Type = pstate.types.NewPtr(v.Type)
			}
			offset = pstate.Rnd(offset, int64(cv.Type.Align))
			cv.Xoffset = offset
			offset += cv.Type.Width

			if v.Name.Byval() && v.Type.Width <= int64(2*pstate.Widthptr) {
				// If it is a small variable captured by value, downgrade it to PAUTO.
				v.SetClass(PAUTO)
				xfunc.Func.Dcl = append(xfunc.Func.Dcl, v)
				body = append(body, pstate.nod(OAS, v, cv))
			} else {
				// Declare variable holding addresses taken from closure
				// and initialize in entry prologue.
				addr := pstate.newname(pstate.lookup("&" + v.Sym.Name))
				addr.Type = pstate.types.NewPtr(v.Type)
				addr.SetClass(PAUTO)
				addr.Name.SetUsed(true)
				addr.Name.Curfn = xfunc
				xfunc.Func.Dcl = append(xfunc.Func.Dcl, addr)
				v.Name.Param.Heapaddr = addr
				if v.Name.Byval() {
					cv = pstate.nod(OADDR, cv, nil)
				}
				body = append(body, pstate.nod(OAS, addr, cv))
			}
		}

		if len(body) > 0 {
			pstate.typecheckslice(body, Etop)
			xfunc.Func.Enter.Set(body)
			xfunc.Func.SetNeedctxt(true)
		}
	}

	pstate.lineno = lno
}

// hasemptycvars returns true iff closure clo has an
// empty list of captured vars.
func hasemptycvars(clo *Node) bool {
	xfunc := clo.Func.Closure
	return xfunc.Func.Cvars.Len() == 0
}

// closuredebugruntimecheck applies boilerplate checks for debug flags
// and compiling runtime
func (pstate *PackageState) closuredebugruntimecheck(clo *Node) {
	if pstate.Debug_closure > 0 {
		xfunc := clo.Func.Closure
		if clo.Esc == EscHeap {
			pstate.Warnl(clo.Pos, "heap closure, captured vars = %v", xfunc.Func.Cvars)
		} else {
			pstate.Warnl(clo.Pos, "stack closure, captured vars = %v", xfunc.Func.Cvars)
		}
	}
	if pstate.compiling_runtime && clo.Esc == EscHeap {
		pstate.yyerrorl(clo.Pos, "heap-allocated closure, not allowed in runtime.")
	}
}

func (pstate *PackageState) walkclosure(clo *Node, init *Nodes) *Node {
	xfunc := clo.Func.Closure

	// If no closure vars, don't bother wrapping.
	if hasemptycvars(clo) {
		if pstate.Debug_closure > 0 {
			pstate.Warnl(clo.Pos, "closure converted to global")
		}
		return xfunc.Func.Nname
	}
	pstate.closuredebugruntimecheck(clo)

	// Create closure in the form of a composite literal.
	// supposing the closure captures an int i and a string s
	// and has one float64 argument and no results,
	// the generated code looks like:
	//
	//	clos = &struct{.F uintptr; i *int; s *string}{func.1, &i, &s}
	//
	// The use of the struct provides type information to the garbage
	// collector so that it can walk the closure. We could use (in this case)
	// [3]unsafe.Pointer instead, but that would leave the gc in the dark.
	// The information appears in the binary in the form of type descriptors;
	// the struct is unnamed so that closures in multiple packages with the
	// same struct type can share the descriptor.

	fields := []*Node{
		pstate.namedfield(".F", pstate.types.Types[TUINTPTR]),
	}
	for _, v := range xfunc.Func.Cvars.Slice() {
		typ := v.Type
		if !v.Name.Byval() {
			typ = pstate.types.NewPtr(typ)
		}
		fields = append(fields, pstate.symfield(v.Sym, typ))
	}
	typ := pstate.tostruct(fields)
	typ.SetNoalg(true)

	clos := pstate.nod(OCOMPLIT, nil, pstate.nod(OIND, pstate.typenod(typ), nil))
	clos.Esc = clo.Esc
	clos.Right.SetImplicit(true)
	clos.List.Set(append([]*Node{pstate.nod(OCFUNC, xfunc.Func.Nname, nil)}, clo.Func.Enter.Slice()...))

	// Force type conversion from *struct to the func type.
	clos = pstate.nod(OCONVNOP, clos, nil)
	clos.Type = clo.Type

	clos = pstate.typecheck(clos, Erv)

	// typecheck will insert a PTRLIT node under CONVNOP,
	// tag it with escape analysis result.
	clos.Left.Esc = clo.Esc

	// non-escaping temp to use, if any.
	// orderexpr did not compute the type; fill it in now.
	if x := pstate.prealloc[clo]; x != nil {
		x.Type = clos.Left.Left.Type
		x.Orig.Type = x.Type
		clos.Left.Right = x
		delete(pstate.prealloc, clo)
	}

	return pstate.walkexpr(clos, init)
}

func (pstate *PackageState) typecheckpartialcall(fn *Node, sym *types.Sym) {
	switch fn.Op {
	case ODOTINTER, ODOTMETH:
		break

	default:
		pstate.Fatalf("invalid typecheckpartialcall")
	}

	// Create top-level function.
	xfunc := pstate.makepartialcall(fn, fn.Type, sym)
	fn.Func = xfunc.Func
	fn.Right = pstate.newname(sym)
	fn.Op = OCALLPART
	fn.Type = xfunc.Type
}

func (pstate *PackageState) makepartialcall(fn *Node, t0 *types.Type, meth *types.Sym) *Node {
	rcvrtype := fn.Left.Type
	sym := pstate.methodSymSuffix(rcvrtype, meth, "-fm")

	if sym.Uniq() {
		return asNode(sym.Def)
	}
	sym.SetUniq(true)

	savecurfn := pstate.Curfn
	pstate.Curfn = nil

	tfn := pstate.nod(OTFUNC, nil, nil)
	tfn.List.Set(pstate.structargs(t0.Params(pstate.types), true))
	tfn.Rlist.Set(pstate.structargs(t0.Results(pstate.types), false))

	disableExport(sym)
	xfunc := pstate.dclfunc(sym, tfn)
	xfunc.Func.SetDupok(true)
	xfunc.Func.SetNeedctxt(true)

	tfn.Type.SetPkg(pstate.types, t0.Pkg(pstate.types))

	// Declare and initialize variable holding receiver.

	cv := pstate.nod(OCLOSUREVAR, nil, nil)
	cv.Type = rcvrtype
	cv.Xoffset = pstate.Rnd(int64(pstate.Widthptr), int64(cv.Type.Align))

	ptr := pstate.newname(pstate.lookup(".this"))
	pstate.declare(ptr, PAUTO)
	ptr.Name.SetUsed(true)
	var body []*Node
	if rcvrtype.IsPtr() || rcvrtype.IsInterface() {
		ptr.Type = rcvrtype
		body = append(body, pstate.nod(OAS, ptr, cv))
	} else {
		ptr.Type = pstate.types.NewPtr(rcvrtype)
		body = append(body, pstate.nod(OAS, ptr, pstate.nod(OADDR, cv, nil)))
	}

	call := pstate.nod(OCALL, pstate.nodSym(OXDOT, ptr, meth), nil)
	call.List.Set(pstate.paramNnames(tfn.Type))
	call.SetIsddd(tfn.Type.IsVariadic(pstate.types))
	if t0.NumResults(pstate.types) != 0 {
		n := pstate.nod(ORETURN, nil, nil)
		n.List.Set1(call)
		call = n
	}
	body = append(body, call)

	xfunc.Nbody.Set(body)
	pstate.funcbody()

	xfunc = pstate.typecheck(xfunc, Etop)
	sym.Def = asTypesNode(xfunc)
	pstate.xtop = append(pstate.xtop, xfunc)
	pstate.Curfn = savecurfn

	return xfunc
}

func (pstate *PackageState) walkpartialcall(n *Node, init *Nodes) *Node {
	// Create closure in the form of a composite literal.
	// For x.M with receiver (x) type T, the generated code looks like:
	//
	//	clos = &struct{F uintptr; R T}{M.T·f, x}
	//
	// Like walkclosure above.

	if n.Left.Type.IsInterface() {
		// Trigger panic for method on nil interface now.
		// Otherwise it happens in the wrapper and is confusing.
		n.Left = pstate.cheapexpr(n.Left, init)

		pstate.checknil(n.Left, init)
	}

	typ := pstate.tostruct([]*Node{
		pstate.namedfield("F", pstate.types.Types[TUINTPTR]),
		pstate.namedfield("R", n.Left.Type),
	})
	typ.SetNoalg(true)

	clos := pstate.nod(OCOMPLIT, nil, pstate.nod(OIND, pstate.typenod(typ), nil))
	clos.Esc = n.Esc
	clos.Right.SetImplicit(true)
	clos.List.Set1(pstate.nod(OCFUNC, n.Func.Nname, nil))
	clos.List.Append(n.Left)

	// Force type conversion from *struct to the func type.
	clos = pstate.nod(OCONVNOP, clos, nil)
	clos.Type = n.Type

	clos = pstate.typecheck(clos, Erv)

	// typecheck will insert a PTRLIT node under CONVNOP,
	// tag it with escape analysis result.
	clos.Left.Esc = n.Esc

	// non-escaping temp to use, if any.
	// orderexpr did not compute the type; fill it in now.
	if x := pstate.prealloc[n]; x != nil {
		x.Type = clos.Left.Left.Type
		x.Orig.Type = x.Type
		clos.Left.Right = x
		delete(pstate.prealloc, n)
	}

	return pstate.walkexpr(clos, init)
}

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"bytes"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/src"
	"strings"
)

func (pstate *PackageState) testdclstack() {
	if !pstate.types.IsDclstackValid() {
		if pstate.nerrors != 0 {
			pstate.errorexit()
		}
		pstate.Fatalf("mark left on the dclstack")
	}
}

// redeclare emits a diagnostic about symbol s being redeclared at pos.
func (pstate *PackageState) redeclare(pos src.XPos, s *types.Sym, where string) {
	if !s.Lastlineno.IsKnown() {
		pkg := s.Origpkg
		if pkg == nil {
			pkg = s.Pkg
		}
		pstate.yyerrorl(pos, "%v redeclared %s\n"+
			"\tprevious declaration during import %q", s, where, pkg.Path)
	} else {
		prevPos := s.Lastlineno

		// When an import and a declaration collide in separate files,
		// present the import as the "redeclared", because the declaration
		// is visible where the import is, but not vice versa.
		// See issue 4510.
		if s.Def == nil {
			pos, prevPos = prevPos, pos
		}

		pstate.yyerrorl(pos, "%v redeclared %s\n"+
			"\tprevious declaration at %v", s, where, pstate.linestr(prevPos))
	}
}

// declare records that Node n declares symbol n.Sym in the specified
// declaration context.
func (pstate *PackageState) declare(n *Node, ctxt Class) {
	if ctxt == PDISCARD {
		return
	}

	if n.isBlank() {
		return
	}

	if n.Name == nil {
		// named OLITERAL needs Name; most OLITERALs don't.
		n.Name = new(Name)
	}

	s := n.Sym

	// kludgy: typecheckok means we're past parsing. Eg genwrapper may declare out of package names later.
	if !pstate.inimport && !pstate.typecheckok && s.Pkg != pstate.localpkg {
		pstate.yyerrorl(n.Pos, "cannot declare name %v", s)
	}

	gen := 0
	if ctxt == PEXTERN {
		if s.Name == "init" {
			pstate.yyerrorl(n.Pos, "cannot declare init - must be func")
		}
		if s.Name == "main" && s.Pkg.Name == "main" {
			pstate.yyerrorl(n.Pos, "cannot declare main - must be func")
		}
		pstate.externdcl = append(pstate.externdcl, n)
	} else {
		if pstate.Curfn == nil && ctxt == PAUTO {
			pstate.lineno = n.Pos
			pstate.Fatalf("automatic outside function")
		}
		if pstate.Curfn != nil {
			pstate.Curfn.Func.Dcl = append(pstate.Curfn.Func.Dcl, n)
		}
		if n.Op == OTYPE {
			pstate.declare_typegen++
			gen = pstate.declare_typegen
		} else if n.Op == ONAME && ctxt == PAUTO && !strings.Contains(s.Name, "·") {
			pstate.vargen++
			gen = pstate.vargen
		}
		pstate.types.Pushdcl(s)
		n.Name.Curfn = pstate.Curfn
	}

	if ctxt == PAUTO {
		n.Xoffset = 0
	}

	if s.Block == pstate.types.Block {
		// functype will print errors about duplicate function arguments.
		// Don't repeat the error here.
		if ctxt != PPARAM && ctxt != PPARAMOUT {
			pstate.redeclare(n.Pos, s, "in this block")
		}
	}

	s.Block = pstate.types.Block
	s.Lastlineno = pstate.lineno
	s.Def = asTypesNode(n)
	n.Name.Vargen = int32(gen)
	n.SetClass(ctxt)

	pstate.autoexport(n, ctxt)
}

func (pstate *PackageState) addvar(n *Node, t *types.Type, ctxt Class) {
	if n == nil || n.Sym == nil || (n.Op != ONAME && n.Op != ONONAME) || t == nil {
		pstate.Fatalf("addvar: n=%v t=%v nil", n, t)
	}

	n.Op = ONAME
	pstate.declare(n, ctxt)
	n.Type = t
}

// declare variables from grammar
// new_name_list (type | [type] = expr_list)
func (pstate *PackageState) variter(vl []*Node, t *Node, el []*Node) []*Node {
	var init []*Node
	doexpr := len(el) > 0

	if len(el) == 1 && len(vl) > 1 {
		e := el[0]
		as2 := pstate.nod(OAS2, nil, nil)
		as2.List.Set(vl)
		as2.Rlist.Set1(e)
		for _, v := range vl {
			v.Op = ONAME
			pstate.declare(v, pstate.dclcontext)
			v.Name.Param.Ntype = t
			v.Name.Defn = as2
			if pstate.Curfn != nil {
				init = append(init, pstate.nod(ODCL, v, nil))
			}
		}

		return append(init, as2)
	}

	for _, v := range vl {
		var e *Node
		if doexpr {
			if len(el) == 0 {
				pstate.yyerror("missing expression in var declaration")
				break
			}
			e = el[0]
			el = el[1:]
		}

		v.Op = ONAME
		pstate.declare(v, pstate.dclcontext)
		v.Name.Param.Ntype = t

		if e != nil || pstate.Curfn != nil || v.isBlank() {
			if pstate.Curfn != nil {
				init = append(init, pstate.nod(ODCL, v, nil))
			}
			e = pstate.nod(OAS, v, e)
			init = append(init, e)
			if e.Right != nil {
				v.Name.Defn = e
			}
		}
	}

	if len(el) != 0 {
		pstate.yyerror("extra expression in var declaration")
	}
	return init
}

// newnoname returns a new ONONAME Node associated with symbol s.
func (pstate *PackageState) newnoname(s *types.Sym) *Node {
	if s == nil {
		pstate.Fatalf("newnoname nil")
	}
	n := pstate.nod(ONONAME, nil, nil)
	n.Sym = s
	n.SetAddable(true)
	n.Xoffset = 0
	return n
}

// newfuncname generates a new name node for a function or method.
// TODO(rsc): Use an ODCLFUNC node instead. See comment in CL 7360.
func (pstate *PackageState) newfuncname(s *types.Sym) *Node {
	return pstate.newfuncnamel(pstate.lineno, s)
}

// newfuncnamel generates a new name node for a function or method.
// TODO(rsc): Use an ODCLFUNC node instead. See comment in CL 7360.
func (pstate *PackageState) newfuncnamel(pos src.XPos, s *types.Sym) *Node {
	n := pstate.newnamel(pos, s)
	n.Func = new(Func)
	n.Func.SetIsHiddenClosure(pstate.Curfn != nil)
	return n
}

// this generates a new name node for a name
// being declared.
func (pstate *PackageState) dclname(s *types.Sym) *Node {
	n := pstate.newname(s)
	n.Op = ONONAME // caller will correct it
	return n
}

func (pstate *PackageState) typenod(t *types.Type) *Node {
	return pstate.typenodl(pstate.src.NoXPos, t)
}

func (pstate *PackageState) typenodl(pos src.XPos, t *types.Type) *Node {
	// if we copied another type with *t = *u
	// then t->nod might be out of date, so
	// check t->nod->type too
	if asNode(t.Nod) == nil || asNode(t.Nod).Type != t {
		t.Nod = asTypesNode(pstate.nodl(pos, OTYPE, nil, nil))
		asNode(t.Nod).Type = t
		asNode(t.Nod).Sym = t.Sym
	}

	return asNode(t.Nod)
}

func (pstate *PackageState) anonfield(typ *types.Type) *Node {
	return pstate.symfield(nil, typ)
}

func (pstate *PackageState) namedfield(s string, typ *types.Type) *Node {
	return pstate.symfield(pstate.lookup(s), typ)
}

func (pstate *PackageState) symfield(s *types.Sym, typ *types.Type) *Node {
	n := pstate.nodSym(ODCLFIELD, nil, s)
	n.Type = typ
	return n
}

// oldname returns the Node that declares symbol s in the current scope.
// If no such Node currently exists, an ONONAME Node is returned instead.
func (pstate *PackageState) oldname(s *types.Sym) *Node {
	n := asNode(s.Def)
	if n == nil {
		// Maybe a top-level declaration will come along later to
		// define s. resolve will check s.Def again once all input
		// source has been processed.
		return pstate.newnoname(s)
	}

	if pstate.Curfn != nil && n.Op == ONAME && n.Name.Curfn != nil && n.Name.Curfn != pstate.Curfn {
		// Inner func is referring to var in outer func.
		//
		// TODO(rsc): If there is an outer variable x and we
		// are parsing x := 5 inside the closure, until we get to
		// the := it looks like a reference to the outer x so we'll
		// make x a closure variable unnecessarily.
		c := n.Name.Param.Innermost
		if c == nil || c.Name.Curfn != pstate.Curfn {
			// Do not have a closure var for the active closure yet; make one.
			c = pstate.newname(s)
			c.SetClass(PAUTOHEAP)
			c.SetIsClosureVar(true)
			c.SetIsddd(n.Isddd())
			c.Name.Defn = n
			c.SetAddable(false)

			// Link into list of active closure variables.
			// Popped from list in func closurebody.
			c.Name.Param.Outer = n.Name.Param.Innermost
			n.Name.Param.Innermost = c

			pstate.Curfn.Func.Cvars.Append(c)
		}

		// return ref to closure var, not original
		return c
	}

	return n
}

// := declarations
func colasname(n *Node) bool {
	switch n.Op {
	case ONAME,
		ONONAME,
		OPACK,
		OTYPE,
		OLITERAL:
		return n.Sym != nil
	}

	return false
}

func (pstate *PackageState) colasdefn(left []*Node, defn *Node) {
	for _, n := range left {
		if n.Sym != nil {
			n.Sym.SetUniq(true)
		}
	}

	var nnew, nerr int
	for i, n := range left {
		if n.isBlank() {
			continue
		}
		if !colasname(n) {
			pstate.yyerrorl(defn.Pos, "non-name %v on left side of :=", n)
			nerr++
			continue
		}

		if !n.Sym.Uniq() {
			pstate.yyerrorl(defn.Pos, "%v repeated on left side of :=", n.Sym)
			n.SetDiag(true)
			nerr++
			continue
		}

		n.Sym.SetUniq(false)
		if n.Sym.Block == pstate.types.Block {
			continue
		}

		nnew++
		n = pstate.newname(n.Sym)
		pstate.declare(n, pstate.dclcontext)
		n.Name.Defn = defn
		defn.Ninit.Append(pstate.nod(ODCL, n, nil))
		left[i] = n
	}

	if nnew == 0 && nerr == 0 {
		pstate.yyerrorl(defn.Pos, "no new variables on left side of :=")
	}
}

// declare the arguments in an
// interface field declaration.
func (pstate *PackageState) ifacedcl(n *Node) {
	if n.Op != ODCLFIELD || n.Left == nil {
		pstate.Fatalf("ifacedcl")
	}

	if n.Sym.IsBlank() {
		pstate.yyerror("methods must have a unique non-blank name")
	}
}

// declare the function proper
// and declare the arguments.
// called in extern-declaration context
// returns in auto-declaration context.
func (pstate *PackageState) funchdr(n *Node) {
	// change the declaration context from extern to auto
	if pstate.Curfn == nil && pstate.dclcontext != PEXTERN {
		pstate.Fatalf("funchdr: dclcontext = %d", pstate.dclcontext)
	}

	pstate.dclcontext = PAUTO
	pstate.types.Markdcl()
	pstate.funcstack = append(pstate.funcstack, pstate.Curfn)
	pstate.Curfn = n

	if n.Func.Nname != nil {
		pstate.funcargs(n.Func.Nname.Name.Param.Ntype)
	} else if n.Func.Ntype != nil {
		pstate.funcargs(n.Func.Ntype)
	} else {
		pstate.funcargs2(n.Type)
	}
}

func (pstate *PackageState) funcargs(nt *Node) {
	if nt.Op != OTFUNC {
		pstate.Fatalf("funcargs %v", nt.Op)
	}

	// re-start the variable generation number
	// we want to use small numbers for the return variables,
	// so let them have the chunk starting at 1.
	//
	// TODO(mdempsky): This is ugly, and only necessary because
	// esc.go uses Vargen to figure out result parameters' index
	// within the result tuple.
	pstate.vargen = nt.Rlist.Len()

	// declare the receiver and in arguments.
	if nt.Left != nil {
		pstate.funcarg(nt.Left, PPARAM)
	}
	for _, n := range nt.List.Slice() {
		pstate.funcarg(n, PPARAM)
	}

	oldvargen := pstate.vargen
	pstate.vargen = 0

	// declare the out arguments.
	gen := nt.List.Len()
	for _, n := range nt.Rlist.Slice() {
		if n.Sym == nil {
			// Name so that escape analysis can track it. ~r stands for 'result'.
			n.Sym = pstate.lookupN("~r", gen)
			gen++
		}
		if n.Sym.IsBlank() {
			// Give it a name so we can assign to it during return. ~b stands for 'blank'.
			// The name must be different from ~r above because if you have
			//	func f() (_ int)
			//	func g() int
			// f is allowed to use a plain 'return' with no arguments, while g is not.
			// So the two cases must be distinguished.
			n.Sym = pstate.lookupN("~b", gen)
			gen++
		}

		pstate.funcarg(n, PPARAMOUT)
	}

	pstate.vargen = oldvargen
}

func (pstate *PackageState) funcarg(n *Node, ctxt Class) {
	if n.Op != ODCLFIELD {
		pstate.Fatalf("funcarg %v", n.Op)
	}
	if n.Sym == nil {
		return
	}

	n.Right = pstate.newnamel(n.Pos, n.Sym)
	n.Right.Name.Param.Ntype = n.Left
	n.Right.SetIsddd(n.Isddd())
	pstate.declare(n.Right, ctxt)

	pstate.vargen++
	n.Right.Name.Vargen = int32(pstate.vargen)
}

// Same as funcargs, except run over an already constructed TFUNC.
// This happens during import, where the hidden_fndcl rule has
// used functype directly to parse the function's type.
func (pstate *PackageState) funcargs2(t *types.Type) {
	if t.Etype != TFUNC {
		pstate.Fatalf("funcargs2 %v", t)
	}

	for _, f := range t.Recvs(pstate.types).Fields(pstate.types).Slice() {
		pstate.funcarg2(f, PPARAM)
	}
	for _, f := range t.Params(pstate.types).Fields(pstate.types).Slice() {
		pstate.funcarg2(f, PPARAM)
	}
	for _, f := range t.Results(pstate.types).Fields(pstate.types).Slice() {
		pstate.funcarg2(f, PPARAMOUT)
	}
}

func (pstate *PackageState) funcarg2(f *types.Field, ctxt Class) {
	if f.Sym == nil {
		return
	}
	n := pstate.newnamel(f.Pos, f.Sym)
	f.Nname = asTypesNode(n)
	n.Type = f.Type
	n.SetIsddd(f.Isddd())
	pstate.declare(n, ctxt)
}

// finish the body.
// called in auto-declaration context.
// returns in extern-declaration context.
func (pstate *PackageState) funcbody() {
	// change the declaration context from auto to extern
	if pstate.dclcontext != PAUTO {
		pstate.Fatalf("funcbody: unexpected dclcontext %d", pstate.dclcontext)
	}
	pstate.types.Popdcl()
	pstate.funcstack, pstate.Curfn = pstate.funcstack[:len(pstate.funcstack)-1], pstate.funcstack[len(pstate.funcstack)-1]
	if pstate.Curfn == nil {
		pstate.dclcontext = PEXTERN
	}
}

// structs, functions, and methods.
// they don't belong here, but where do they belong?
func (pstate *PackageState) checkembeddedtype(t *types.Type) {
	if t == nil {
		return
	}

	if t.Sym == nil && t.IsPtr() {
		t = t.Elem(pstate.types)
		if t.IsInterface() {
			pstate.yyerror("embedded type cannot be a pointer to interface")
		}
	}

	if t.IsPtr() || t.IsUnsafePtr() {
		pstate.yyerror("embedded type cannot be a pointer")
	} else if t.Etype == TFORW && !t.ForwardType(pstate.types).Embedlineno.IsKnown() {
		t.ForwardType(pstate.types).Embedlineno = pstate.lineno
	}
}

func (pstate *PackageState) structfield(n *Node) *types.Field {
	lno := pstate.lineno
	pstate.lineno = n.Pos

	if n.Op != ODCLFIELD {
		pstate.Fatalf("structfield: oops %v\n", n)
	}

	f := types.NewField()
	f.Pos = n.Pos
	f.Sym = n.Sym

	if n.Left != nil {
		n.Left = pstate.typecheck(n.Left, Etype)
		n.Type = n.Left.Type
		n.Left = nil
	}

	f.Type = n.Type
	if f.Type == nil {
		f.SetBroke(true)
	}

	if n.Embedded() {
		pstate.checkembeddedtype(n.Type)
		f.Embedded = 1
	} else {
		f.Embedded = 0
	}

	switch u := n.Val().U.(type) {
	case string:
		f.Note = u
	default:
		pstate.yyerror("field tag must be a string")
	case nil:
		// no-op
	}

	pstate.lineno = lno
	return f
}

// checkdupfields emits errors for duplicately named fields or methods in
// a list of struct or interface types.
func (pstate *PackageState) checkdupfields(what string, ts ...*types.Type) {
	seen := make(map[*types.Sym]bool)
	for _, t := range ts {
		for _, f := range t.Fields(pstate.types).Slice() {
			if f.Sym == nil || f.Sym.IsBlank() {
				continue
			}
			if seen[f.Sym] {
				pstate.yyerrorl(f.Pos, "duplicate %s %s", what, f.Sym.Name)
				continue
			}
			seen[f.Sym] = true
		}
	}
}

// convert a parsed id/type list into
// a type for struct/interface/arglist
func (pstate *PackageState) tostruct(l []*Node) *types.Type {
	t := types.New(TSTRUCT)
	pstate.tostruct0(t, l)
	return t
}

func (pstate *PackageState) tostruct0(t *types.Type, l []*Node) {
	if t == nil || !t.IsStruct() {
		pstate.Fatalf("struct expected")
	}

	fields := make([]*types.Field, len(l))
	for i, n := range l {
		f := pstate.structfield(n)
		if f.Broke() {
			t.SetBroke(true)
		}
		fields[i] = f
	}
	t.SetFields(pstate.types, fields)

	pstate.checkdupfields("field", t)

	if !t.Broke() {
		pstate.checkwidth(t)
	}
}

func (pstate *PackageState) tofunargs(l []*Node, funarg types.Funarg) *types.Type {
	t := types.New(TSTRUCT)
	t.StructType(pstate.types).Funarg = funarg

	fields := make([]*types.Field, len(l))
	for i, n := range l {
		f := pstate.structfield(n)
		f.SetIsddd(n.Isddd())
		if n.Right != nil {
			n.Right.Type = f.Type
			f.Nname = asTypesNode(n.Right)
		}
		if f.Broke() {
			t.SetBroke(true)
		}
		fields[i] = f
	}
	t.SetFields(pstate.types, fields)
	return t
}

func (pstate *PackageState) tofunargsfield(fields []*types.Field, funarg types.Funarg) *types.Type {
	t := types.New(TSTRUCT)
	t.StructType(pstate.types).Funarg = funarg
	t.SetFields(pstate.types, fields)
	return t
}

func (pstate *PackageState) interfacefield(n *Node) *types.Field {
	lno := pstate.lineno
	pstate.lineno = n.Pos

	if n.Op != ODCLFIELD {
		pstate.Fatalf("interfacefield: oops %v\n", n)
	}

	if n.Val().Ctype(pstate) != CTxxx {
		pstate.yyerror("interface method cannot have annotation")
	}

	// MethodSpec = MethodName Signature | InterfaceTypeName .
	//
	// If Sym != nil, then Sym is MethodName and Left is Signature.
	// Otherwise, Left is InterfaceTypeName.

	if n.Left != nil {
		n.Left = pstate.typecheck(n.Left, Etype)
		n.Type = n.Left.Type
		n.Left = nil
	}

	f := types.NewField()
	f.Pos = n.Pos
	f.Sym = n.Sym
	f.Type = n.Type
	if f.Type == nil {
		f.SetBroke(true)
	}

	pstate.lineno = lno
	return f
}

func (pstate *PackageState) tointerface(l []*Node) *types.Type {
	if len(l) == 0 {
		return pstate.types.Types[TINTER]
	}
	t := types.New(TINTER)
	pstate.tointerface0(t, l)
	return t
}

func (pstate *PackageState) tointerface0(t *types.Type, l []*Node) {
	if t == nil || !t.IsInterface() {
		pstate.Fatalf("interface expected")
	}

	var fields []*types.Field
	for _, n := range l {
		f := pstate.interfacefield(n)
		if f.Broke() {
			t.SetBroke(true)
		}
		fields = append(fields, f)
	}
	t.SetInterface(pstate.types, fields)
}

func (pstate *PackageState) fakeRecv() *Node {
	return pstate.anonfield(pstate.types.FakeRecvType())
}

func (pstate *PackageState) fakeRecvField() *types.Field {
	f := types.NewField()
	f.Type = pstate.types.FakeRecvType()
	return f
}

// isifacemethod reports whether (field) m is
// an interface method. Such methods have the
// special receiver type types.FakeRecvType().
func (pstate *PackageState) isifacemethod(f *types.Type) bool {
	return f.Recv(pstate.types).Type == pstate.types.FakeRecvType()
}

// turn a parsed function declaration into a type
func (pstate *PackageState) functype(this *Node, in, out []*Node) *types.Type {
	t := types.New(TFUNC)
	pstate.functype0(t, this, in, out)
	return t
}

func (pstate *PackageState) functype0(t *types.Type, this *Node, in, out []*Node) {
	if t == nil || t.Etype != TFUNC {
		pstate.Fatalf("function type expected")
	}

	var rcvr []*Node
	if this != nil {
		rcvr = []*Node{this}
	}
	t.FuncType(pstate.types).Receiver = pstate.tofunargs(rcvr, types.FunargRcvr)
	t.FuncType(pstate.types).Params = pstate.tofunargs(in, types.FunargParams)
	t.FuncType(pstate.types).Results = pstate.tofunargs(out, types.FunargResults)

	pstate.checkdupfields("argument", t.Recvs(pstate.types), t.Params(pstate.types), t.Results(pstate.types))

	if t.Recvs(pstate.types).Broke() || t.Results(pstate.types).Broke() || t.Params(pstate.types).Broke() {
		t.SetBroke(true)
	}

	t.FuncType(pstate.types).Outnamed = t.NumResults(pstate.types) > 0 && pstate.origSym(t.Results(pstate.types).Field(pstate.types, 0).Sym) != nil
}

func (pstate *PackageState) functypefield(this *types.Field, in, out []*types.Field) *types.Type {
	t := types.New(TFUNC)
	pstate.functypefield0(t, this, in, out)
	return t
}

func (pstate *PackageState) functypefield0(t *types.Type, this *types.Field, in, out []*types.Field) {
	var rcvr []*types.Field
	if this != nil {
		rcvr = []*types.Field{this}
	}
	t.FuncType(pstate.types).Receiver = pstate.tofunargsfield(rcvr, types.FunargRcvr)
	t.FuncType(pstate.types).Params = pstate.tofunargsfield(in, types.FunargParams)
	t.FuncType(pstate.types).Results = pstate.tofunargsfield(out, types.FunargResults)

	t.FuncType(pstate.types).Outnamed = t.NumResults(pstate.types) > 0 && pstate.origSym(t.Results(pstate.types).Field(pstate.types, 0).Sym) != nil
}

// origSym returns the original symbol written by the user.
func (pstate *PackageState) origSym(s *types.Sym) *types.Sym {
	if s == nil {
		return nil
	}

	if len(s.Name) > 1 && s.Name[0] == '~' {
		switch s.Name[1] {
		case 'r': // originally an unnamed result
			return nil
		case 'b': // originally the blank identifier _
			// TODO(mdempsky): Does s.Pkg matter here?
			return pstate.nblank.Sym
		}
		return s
	}

	if strings.HasPrefix(s.Name, ".anon") {
		// originally an unnamed or _ name (see subr.go: structargs)
		return nil
	}

	return s
}

// methodSym returns the method symbol representing a method name
// associated with a specific receiver type.
//
// Method symbols can be used to distinguish the same method appearing
// in different method sets. For example, T.M and (*T).M have distinct
// method symbols.
func (pstate *PackageState) methodSym(recv *types.Type, msym *types.Sym) *types.Sym {
	return pstate.methodSymSuffix(recv, msym, "")
}

// methodSymSuffix is like methodsym, but allows attaching a
// distinguisher suffix. To avoid collisions, the suffix must not
// start with a letter, number, or period.
func (pstate *PackageState) methodSymSuffix(recv *types.Type, msym *types.Sym, suffix string) *types.Sym {
	if msym.IsBlank() {
		pstate.Fatalf("blank method name")
	}

	rsym := recv.Sym
	if recv.IsPtr() {
		if rsym != nil {
			pstate.Fatalf("declared pointer receiver type: %v", recv)
		}
		rsym = recv.Elem(pstate.types).Sym
	}

	// Find the package the receiver type appeared in. For
	// anonymous receiver types (i.e., anonymous structs with
	// embedded fields), use the "go" pseudo-package instead.
	rpkg := pstate.gopkg
	if rsym != nil {
		rpkg = rsym.Pkg
	}

	var b bytes.Buffer
	if recv.IsPtr() {
		// The parentheses aren't really necessary, but
		// they're pretty traditional at this point.
		fmt.Fprintf(&b, "(%-S)", recv)
	} else {
		fmt.Fprintf(&b, "%-S", recv)
	}

	// A particular receiver type may have multiple non-exported
	// methods with the same name. To disambiguate them, include a
	// package qualifier for names that came from a different
	// package than the receiver type.
	if !types.IsExported(msym.Name) && msym.Pkg != rpkg {
		b.WriteString(".")
		b.WriteString(msym.Pkg.Prefix)
	}

	b.WriteString(".")
	b.WriteString(msym.Name)
	b.WriteString(suffix)

	return rpkg.LookupBytes(pstate.types, b.Bytes())
}

// Add a method, declared as a function.
// - msym is the method symbol
// - t is function type (with receiver)
// Returns a pointer to the existing or added Field.
func (pstate *PackageState) addmethod(msym *types.Sym, t *types.Type, local, nointerface bool) *types.Field {
	if msym == nil {
		pstate.Fatalf("no method symbol")
	}

	// get parent type sym
	rf := t.Recv(pstate.types) // ptr to this structure
	if rf == nil {
		pstate.yyerror("missing receiver")
		return nil
	}

	mt := pstate.methtype(rf.Type)
	if mt == nil || mt.Sym == nil {
		pa := rf.Type
		t := pa
		if t != nil && t.IsPtr() {
			if t.Sym != nil {
				pstate.yyerror("invalid receiver type %v (%v is a pointer type)", pa, t)
				return nil
			}
			t = t.Elem(pstate.types)
		}

		switch {
		case t == nil || t.Broke():
		// rely on typecheck having complained before
		case t.Sym == nil:
			pstate.yyerror("invalid receiver type %v (%v is not a defined type)", pa, t)
		case t.IsPtr():
			pstate.yyerror("invalid receiver type %v (%v is a pointer type)", pa, t)
		case t.IsInterface():
			pstate.yyerror("invalid receiver type %v (%v is an interface type)", pa, t)
		default:
			// Should have picked off all the reasons above,
			// but just in case, fall back to generic error.
			pstate.yyerror("invalid receiver type %v (%L / %L)", pa, pa, t)
		}
		return nil
	}

	if local && mt.Sym.Pkg != pstate.localpkg {
		pstate.yyerror("cannot define new methods on non-local type %v", mt)
		return nil
	}

	if msym.IsBlank() {
		return nil
	}

	if mt.IsStruct() {
		for _, f := range mt.Fields(pstate.types).Slice() {
			if f.Sym == msym {
				pstate.yyerror("type %v has both field and method named %v", mt, msym)
				return nil
			}
		}
	}

	for _, f := range mt.Methods().Slice() {
		if msym.Name != f.Sym.Name {
			continue
		}
		// eqtype only checks that incoming and result parameters match,
		// so explicitly check that the receiver parameters match too.
		if !pstate.eqtype(t, f.Type) || !pstate.eqtype(t.Recv(pstate.types).Type, f.Type.Recv(pstate.types).Type) {
			pstate.yyerror("method redeclared: %v.%v\n\t%v\n\t%v", mt, msym, f.Type, t)
		}
		return f
	}

	f := types.NewField()
	f.Pos = pstate.lineno
	f.Sym = msym
	f.Type = t
	f.SetNointerface(nointerface)

	mt.Methods().Append(f)
	return f
}

func funcsymname(s *types.Sym) string {
	return s.Name + "·f"
}

// funcsym returns s·f.
func (pstate *PackageState) funcsym(s *types.Sym) *types.Sym {
	// funcsymsmu here serves to protect not just mutations of funcsyms (below),
	// but also the package lookup of the func sym name,
	// since this function gets called concurrently from the backend.
	// There are no other concurrent package lookups in the backend,
	// except for the types package, which is protected separately.
	// Reusing funcsymsmu to also cover this package lookup
	// avoids a general, broader, expensive package lookup mutex.
	// Note makefuncsym also does package look-up of func sym names,
	// but that it is only called serially, from the front end.
	pstate.funcsymsmu.Lock()
	sf, existed := s.Pkg.LookupOK(pstate.types, funcsymname(s))
	// Don't export s·f when compiling for dynamic linking.
	// When dynamically linking, the necessary function
	// symbols will be created explicitly with makefuncsym.
	// See the makefuncsym comment for details.
	if !pstate.Ctxt.Flag_dynlink && !existed {
		pstate.funcsyms = append(pstate.funcsyms, s)
	}
	pstate.funcsymsmu.Unlock()
	return sf
}

// makefuncsym ensures that s·f is exported.
// It is only used with -dynlink.
// When not compiling for dynamic linking,
// the funcsyms are created as needed by
// the packages that use them.
// Normally we emit the s·f stubs as DUPOK syms,
// but DUPOK doesn't work across shared library boundaries.
// So instead, when dynamic linking, we only create
// the s·f stubs in s's package.
func (pstate *PackageState) makefuncsym(s *types.Sym) {
	if !pstate.Ctxt.Flag_dynlink {
		pstate.Fatalf("makefuncsym dynlink")
	}
	if s.IsBlank() {
		return
	}
	if pstate.compiling_runtime && (s.Name == "getg" || s.Name == "getclosureptr" || s.Name == "getcallerpc" || s.Name == "getcallersp") {
		// runtime.getg(), getclosureptr(), getcallerpc(), and
		// getcallersp() are not real functions and so do not
		// get funcsyms.
		return
	}
	if _, existed := s.Pkg.LookupOK(pstate.types, funcsymname(s)); !existed {
		pstate.funcsyms = append(pstate.funcsyms, s)
	}
}

// disableExport prevents sym from being included in package export
// data. To be effectual, it must be called before declare.
func disableExport(sym *types.Sym) {
	sym.SetOnExportList(true)
}

func (pstate *PackageState) dclfunc(sym *types.Sym, tfn *Node) *Node {
	if tfn.Op != OTFUNC {
		pstate.Fatalf("expected OTFUNC node, got %v", tfn)
	}

	fn := pstate.nod(ODCLFUNC, nil, nil)
	fn.Func.Nname = pstate.newfuncname(sym)
	fn.Func.Nname.Name.Defn = fn
	fn.Func.Nname.Name.Param.Ntype = tfn
	pstate.declare(fn.Func.Nname, PFUNC)
	pstate.funchdr(fn)
	fn.Func.Nname.Name.Param.Ntype = pstate.typecheck(fn.Func.Nname.Name.Param.Ntype, Etype)
	return fn
}

type nowritebarrierrecChecker struct {
	// extraCalls contains extra function calls that may not be
	// visible during later analysis. It maps from the ODCLFUNC of
	// the caller to a list of callees.
	extraCalls map[*Node][]nowritebarrierrecCall

	// curfn is the current function during AST walks.
	curfn *Node
}

type nowritebarrierrecCall struct {
	target *Node    // ODCLFUNC of caller or callee
	lineno src.XPos // line of call
}

type nowritebarrierrecCallSym struct {
	target *obj.LSym // LSym of callee
	lineno src.XPos  // line of call
}

// newNowritebarrierrecChecker creates a nowritebarrierrecChecker. It
// must be called before transformclosure and walk.
func (pstate *PackageState) newNowritebarrierrecChecker() *nowritebarrierrecChecker {
	c := &nowritebarrierrecChecker{
		extraCalls: make(map[*Node][]nowritebarrierrecCall),
	}

	// Find all systemstack calls and record their targets. In
	// general, flow analysis can't see into systemstack, but it's
	// important to handle it for this check, so we model it
	// directly. This has to happen before transformclosure since
	// it's a lot harder to work out the argument after.
	for _, n := range pstate.xtop {
		if n.Op != ODCLFUNC {
			continue
		}
		c.curfn = n
		inspect(n, c.findExtraCalls)
	}
	c.curfn = nil
	return c
}

func (c *nowritebarrierrecChecker) findExtraCalls(pstate *PackageState, n *Node) bool {
	if n.Op != OCALLFUNC {
		return true
	}
	fn := n.Left
	if fn == nil || fn.Op != ONAME || fn.Class() != PFUNC || fn.Name.Defn == nil {
		return true
	}
	if !pstate.isRuntimePkg(fn.Sym.Pkg) || fn.Sym.Name != "systemstack" {
		return true
	}

	var callee *Node
	arg := n.List.First()
	switch arg.Op {
	case ONAME:
		callee = arg.Name.Defn
	case OCLOSURE:
		callee = arg.Func.Closure
	default:
		pstate.Fatalf("expected ONAME or OCLOSURE node, got %+v", arg)
	}
	if callee.Op != ODCLFUNC {
		pstate.Fatalf("expected ODCLFUNC node, got %+v", callee)
	}
	c.extraCalls[c.curfn] = append(c.extraCalls[c.curfn], nowritebarrierrecCall{callee, n.Pos})
	return true
}

// recordCall records a call from ODCLFUNC node "from", to function
// symbol "to" at position pos.
//
// This should be done as late as possible during compilation to
// capture precise call graphs. The target of the call is an LSym
// because that's all we know after we start SSA.
//
// This can be called concurrently for different from Nodes.
func (c *nowritebarrierrecChecker) recordCall(pstate *PackageState, from *Node, to *obj.LSym, pos src.XPos) {
	if from.Op != ODCLFUNC {
		pstate.Fatalf("expected ODCLFUNC, got %v", from)
	}
	// We record this information on the *Func so this is
	// concurrent-safe.
	fn := from.Func
	if fn.nwbrCalls == nil {
		fn.nwbrCalls = new([]nowritebarrierrecCallSym)
	}
	*fn.nwbrCalls = append(*fn.nwbrCalls, nowritebarrierrecCallSym{to, pos})
}

func (c *nowritebarrierrecChecker) check(pstate *PackageState) {
	// We walk the call graph as late as possible so we can
	// capture all calls created by lowering, but this means we
	// only get to see the obj.LSyms of calls. symToFunc lets us
	// get back to the ODCLFUNCs.
	symToFunc := make(map[*obj.LSym]*Node)
	// funcs records the back-edges of the BFS call graph walk. It
	// maps from the ODCLFUNC of each function that must not have
	// write barriers to the call that inhibits them. Functions
	// that are directly marked go:nowritebarrierrec are in this
	// map with a zero-valued nowritebarrierrecCall. This also
	// acts as the set of marks for the BFS of the call graph.
	funcs := make(map[*Node]nowritebarrierrecCall)
	// q is the queue of ODCLFUNC Nodes to visit in BFS order.
	var q nodeQueue

	for _, n := range pstate.xtop {
		if n.Op != ODCLFUNC {
			continue
		}

		symToFunc[n.Func.lsym] = n

		// Make nowritebarrierrec functions BFS roots.
		if n.Func.Pragma&Nowritebarrierrec != 0 {
			funcs[n] = nowritebarrierrecCall{}
			q.pushRight(n)
		}
		// Check go:nowritebarrier functions.
		if n.Func.Pragma&Nowritebarrier != 0 && n.Func.WBPos.IsKnown() {
			pstate.yyerrorl(n.Func.WBPos, "write barrier prohibited")
		}
	}

	// Perform a BFS of the call graph from all
	// go:nowritebarrierrec functions.
	enqueue := func(src, target *Node, pos src.XPos) {
		if target.Func.Pragma&Yeswritebarrierrec != 0 {
			// Don't flow into this function.
			return
		}
		if _, ok := funcs[target]; ok {
			// Already found a path to target.
			return
		}

		// Record the path.
		funcs[target] = nowritebarrierrecCall{target: src, lineno: pos}
		q.pushRight(target)
	}
	for !q.empty() {
		fn := q.popLeft()

		// Check fn.
		if fn.Func.WBPos.IsKnown() {
			var err bytes.Buffer
			call := funcs[fn]
			for call.target != nil {
				fmt.Fprintf(&err, "\n\t%v: called by %v", pstate.linestr(call.lineno), call.target.Func.Nname)
				call = funcs[call.target]
			}
			pstate.yyerrorl(fn.Func.WBPos, "write barrier prohibited by caller; %v%s", fn.Func.Nname, err.String())
			continue
		}

		// Enqueue fn's calls.
		for _, callee := range c.extraCalls[fn] {
			enqueue(fn, callee.target, callee.lineno)
		}
		if fn.Func.nwbrCalls == nil {
			continue
		}
		for _, callee := range *fn.Func.nwbrCalls {
			target := symToFunc[callee.target]
			if target != nil {
				enqueue(fn, target, callee.lineno)
			}
		}
	}
}

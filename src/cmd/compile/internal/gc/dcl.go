package gc

import (
	"bytes"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/src"
	"strings"
)

func (psess *PackageSession) testdclstack() {
	if !psess.types.IsDclstackValid() {
		if psess.nerrors != 0 {
			psess.
				errorexit()
		}
		psess.
			Fatalf("mark left on the dclstack")
	}
}

// redeclare emits a diagnostic about symbol s being redeclared at pos.
func (psess *PackageSession) redeclare(pos src.XPos, s *types.Sym, where string) {
	if !s.Lastlineno.IsKnown() {
		pkg := s.Origpkg
		if pkg == nil {
			pkg = s.Pkg
		}
		psess.
			yyerrorl(pos, "%v redeclared %s\n"+
				"\tprevious declaration during import %q", s, where, pkg.Path)
	} else {
		prevPos := s.Lastlineno

		if s.Def == nil {
			pos, prevPos = prevPos, pos
		}
		psess.
			yyerrorl(pos, "%v redeclared %s\n"+
				"\tprevious declaration at %v", s, where, psess.linestr(prevPos))
	}
}

// declare records that Node n declares symbol n.Sym in the specified
// declaration context.
func (psess *PackageSession) declare(n *Node, ctxt Class) {
	if ctxt == PDISCARD {
		return
	}

	if n.isBlank() {
		return
	}

	if n.Name == nil {

		n.Name = new(Name)
	}

	s := n.Sym

	if !psess.inimport && !psess.typecheckok && s.Pkg != psess.localpkg {
		psess.
			yyerrorl(n.Pos, "cannot declare name %v", s)
	}

	gen := 0
	if ctxt == PEXTERN {
		if s.Name == "init" {
			psess.
				yyerrorl(n.Pos, "cannot declare init - must be func")
		}
		if s.Name == "main" && s.Pkg.Name == "main" {
			psess.
				yyerrorl(n.Pos, "cannot declare main - must be func")
		}
		psess.
			externdcl = append(psess.externdcl, n)
	} else {
		if psess.Curfn == nil && ctxt == PAUTO {
			psess.
				lineno = n.Pos
			psess.
				Fatalf("automatic outside function")
		}
		if psess.Curfn != nil {
			psess.
				Curfn.Func.Dcl = append(psess.Curfn.Func.Dcl, n)
		}
		if n.Op == OTYPE {
			psess.
				declare_typegen++
			gen = psess.declare_typegen
		} else if n.Op == ONAME && ctxt == PAUTO && !strings.Contains(s.Name, "·") {
			psess.
				vargen++
			gen = psess.vargen
		}
		psess.types.
			Pushdcl(s)
		n.Name.Curfn = psess.Curfn
	}

	if ctxt == PAUTO {
		n.Xoffset = 0
	}

	if s.Block == psess.types.Block {

		if ctxt != PPARAM && ctxt != PPARAMOUT {
			psess.
				redeclare(n.Pos, s, "in this block")
		}
	}

	s.Block = psess.types.Block
	s.Lastlineno = psess.lineno
	s.Def = asTypesNode(n)
	n.Name.Vargen = int32(gen)
	n.SetClass(ctxt)
	psess.
		autoexport(n, ctxt)
}

func (psess *PackageSession) addvar(n *Node, t *types.Type, ctxt Class) {
	if n == nil || n.Sym == nil || (n.Op != ONAME && n.Op != ONONAME) || t == nil {
		psess.
			Fatalf("addvar: n=%v t=%v nil", n, t)
	}

	n.Op = ONAME
	psess.
		declare(n, ctxt)
	n.Type = t
}

// declare variables from grammar
// new_name_list (type | [type] = expr_list)
func (psess *PackageSession) variter(vl []*Node, t *Node, el []*Node) []*Node {
	var init []*Node
	doexpr := len(el) > 0

	if len(el) == 1 && len(vl) > 1 {
		e := el[0]
		as2 := psess.nod(OAS2, nil, nil)
		as2.List.Set(vl)
		as2.Rlist.Set1(e)
		for _, v := range vl {
			v.Op = ONAME
			psess.
				declare(v, psess.dclcontext)
			v.Name.Param.Ntype = t
			v.Name.Defn = as2
			if psess.Curfn != nil {
				init = append(init, psess.nod(ODCL, v, nil))
			}
		}

		return append(init, as2)
	}

	for _, v := range vl {
		var e *Node
		if doexpr {
			if len(el) == 0 {
				psess.
					yyerror("missing expression in var declaration")
				break
			}
			e = el[0]
			el = el[1:]
		}

		v.Op = ONAME
		psess.
			declare(v, psess.dclcontext)
		v.Name.Param.Ntype = t

		if e != nil || psess.Curfn != nil || v.isBlank() {
			if psess.Curfn != nil {
				init = append(init, psess.nod(ODCL, v, nil))
			}
			e = psess.nod(OAS, v, e)
			init = append(init, e)
			if e.Right != nil {
				v.Name.Defn = e
			}
		}
	}

	if len(el) != 0 {
		psess.
			yyerror("extra expression in var declaration")
	}
	return init
}

// newnoname returns a new ONONAME Node associated with symbol s.
func (psess *PackageSession) newnoname(s *types.Sym) *Node {
	if s == nil {
		psess.
			Fatalf("newnoname nil")
	}
	n := psess.nod(ONONAME, nil, nil)
	n.Sym = s
	n.SetAddable(true)
	n.Xoffset = 0
	return n
}

// newfuncname generates a new name node for a function or method.
// TODO(rsc): Use an ODCLFUNC node instead. See comment in CL 7360.
func (psess *PackageSession) newfuncname(s *types.Sym) *Node {
	return psess.newfuncnamel(psess.lineno, s)
}

// newfuncnamel generates a new name node for a function or method.
// TODO(rsc): Use an ODCLFUNC node instead. See comment in CL 7360.
func (psess *PackageSession) newfuncnamel(pos src.XPos, s *types.Sym) *Node {
	n := psess.newnamel(pos, s)
	n.Func = new(Func)
	n.Func.SetIsHiddenClosure(psess.Curfn != nil)
	return n
}

// this generates a new name node for a name
// being declared.
func (psess *PackageSession) dclname(s *types.Sym) *Node {
	n := psess.newname(s)
	n.Op = ONONAME
	return n
}

func (psess *PackageSession) typenod(t *types.Type) *Node {
	return psess.typenodl(psess.src.NoXPos, t)
}

func (psess *PackageSession) typenodl(pos src.XPos, t *types.Type) *Node {

	if asNode(t.Nod) == nil || asNode(t.Nod).Type != t {
		t.Nod = asTypesNode(psess.nodl(pos, OTYPE, nil, nil))
		asNode(t.Nod).Type = t
		asNode(t.Nod).Sym = t.Sym
	}

	return asNode(t.Nod)
}

func (psess *PackageSession) anonfield(typ *types.Type) *Node {
	return psess.symfield(nil, typ)
}

func (psess *PackageSession) namedfield(s string, typ *types.Type) *Node {
	return psess.symfield(psess.lookup(s), typ)
}

func (psess *PackageSession) symfield(s *types.Sym, typ *types.Type) *Node {
	n := psess.nodSym(ODCLFIELD, nil, s)
	n.Type = typ
	return n
}

// oldname returns the Node that declares symbol s in the current scope.
// If no such Node currently exists, an ONONAME Node is returned instead.
func (psess *PackageSession) oldname(s *types.Sym) *Node {
	n := asNode(s.Def)
	if n == nil {

		return psess.newnoname(s)
	}

	if psess.Curfn != nil && n.Op == ONAME && n.Name.Curfn != nil && n.Name.Curfn != psess.Curfn {

		c := n.Name.Param.Innermost
		if c == nil || c.Name.Curfn != psess.Curfn {

			c = psess.newname(s)
			c.SetClass(PAUTOHEAP)
			c.SetIsClosureVar(true)
			c.SetIsddd(n.Isddd())
			c.Name.Defn = n
			c.SetAddable(false)

			c.Name.Param.Outer = n.Name.Param.Innermost
			n.Name.Param.Innermost = c
			psess.
				Curfn.Func.Cvars.Append(c)
		}

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

func (psess *PackageSession) colasdefn(left []*Node, defn *Node) {
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
			psess.
				yyerrorl(defn.Pos, "non-name %v on left side of :=", n)
			nerr++
			continue
		}

		if !n.Sym.Uniq() {
			psess.
				yyerrorl(defn.Pos, "%v repeated on left side of :=", n.Sym)
			n.SetDiag(true)
			nerr++
			continue
		}

		n.Sym.SetUniq(false)
		if n.Sym.Block == psess.types.Block {
			continue
		}

		nnew++
		n = psess.newname(n.Sym)
		psess.
			declare(n, psess.dclcontext)
		n.Name.Defn = defn
		defn.Ninit.Append(psess.nod(ODCL, n, nil))
		left[i] = n
	}

	if nnew == 0 && nerr == 0 {
		psess.
			yyerrorl(defn.Pos, "no new variables on left side of :=")
	}
}

// declare the arguments in an
// interface field declaration.
func (psess *PackageSession) ifacedcl(n *Node) {
	if n.Op != ODCLFIELD || n.Left == nil {
		psess.
			Fatalf("ifacedcl")
	}

	if n.Sym.IsBlank() {
		psess.
			yyerror("methods must have a unique non-blank name")
	}
}

// declare the function proper
// and declare the arguments.
// called in extern-declaration context
// returns in auto-declaration context.
func (psess *PackageSession) funchdr(n *Node) {

	if psess.Curfn == nil && psess.dclcontext != PEXTERN {
		psess.
			Fatalf("funchdr: dclcontext = %d", psess.dclcontext)
	}
	psess.
		dclcontext = PAUTO
	psess.types.
		Markdcl()
	psess.
		funcstack = append(psess.funcstack, psess.Curfn)
	psess.
		Curfn = n

	if n.Func.Nname != nil {
		psess.
			funcargs(n.Func.Nname.Name.Param.Ntype)
	} else if n.Func.Ntype != nil {
		psess.
			funcargs(n.Func.Ntype)
	} else {
		psess.
			funcargs2(n.Type)
	}
}

func (psess *PackageSession) funcargs(nt *Node) {
	if nt.Op != OTFUNC {
		psess.
			Fatalf("funcargs %v", nt.Op)
	}
	psess.
		vargen = nt.Rlist.Len()

	if nt.Left != nil {
		psess.
			funcarg(nt.Left, PPARAM)
	}
	for _, n := range nt.List.Slice() {
		psess.
			funcarg(n, PPARAM)
	}

	oldvargen := psess.vargen
	psess.
		vargen = 0

	gen := nt.List.Len()
	for _, n := range nt.Rlist.Slice() {
		if n.Sym == nil {

			n.Sym = psess.lookupN("~r", gen)
			gen++
		}
		if n.Sym.IsBlank() {

			n.Sym = psess.lookupN("~b", gen)
			gen++
		}
		psess.
			funcarg(n, PPARAMOUT)
	}
	psess.
		vargen = oldvargen
}

func (psess *PackageSession) funcarg(n *Node, ctxt Class) {
	if n.Op != ODCLFIELD {
		psess.
			Fatalf("funcarg %v", n.Op)
	}
	if n.Sym == nil {
		return
	}

	n.Right = psess.newnamel(n.Pos, n.Sym)
	n.Right.Name.Param.Ntype = n.Left
	n.Right.SetIsddd(n.Isddd())
	psess.
		declare(n.Right, ctxt)
	psess.
		vargen++
	n.Right.Name.Vargen = int32(psess.vargen)
}

// Same as funcargs, except run over an already constructed TFUNC.
// This happens during import, where the hidden_fndcl rule has
// used functype directly to parse the function's type.
func (psess *PackageSession) funcargs2(t *types.Type) {
	if t.Etype != TFUNC {
		psess.
			Fatalf("funcargs2 %v", t)
	}

	for _, f := range t.Recvs(psess.types).Fields(psess.types).Slice() {
		psess.
			funcarg2(f, PPARAM)
	}
	for _, f := range t.Params(psess.types).Fields(psess.types).Slice() {
		psess.
			funcarg2(f, PPARAM)
	}
	for _, f := range t.Results(psess.types).Fields(psess.types).Slice() {
		psess.
			funcarg2(f, PPARAMOUT)
	}
}

func (psess *PackageSession) funcarg2(f *types.Field, ctxt Class) {
	if f.Sym == nil {
		return
	}
	n := psess.newnamel(f.Pos, f.Sym)
	f.Nname = asTypesNode(n)
	n.Type = f.Type
	n.SetIsddd(f.Isddd())
	psess.
		declare(n, ctxt)
}

// stack of previous values of Curfn

// finish the body.
// called in auto-declaration context.
// returns in extern-declaration context.
func (psess *PackageSession) funcbody() {

	if psess.dclcontext != PAUTO {
		psess.
			Fatalf("funcbody: unexpected dclcontext %d", psess.dclcontext)
	}
	psess.types.
		Popdcl()
	psess.
		funcstack, psess.Curfn = psess.funcstack[:len(psess.funcstack)-1], psess.funcstack[len(psess.funcstack)-1]
	if psess.Curfn == nil {
		psess.
			dclcontext = PEXTERN
	}
}

// structs, functions, and methods.
// they don't belong here, but where do they belong?
func (psess *PackageSession) checkembeddedtype(t *types.Type) {
	if t == nil {
		return
	}

	if t.Sym == nil && t.IsPtr() {
		t = t.Elem(psess.types)
		if t.IsInterface() {
			psess.
				yyerror("embedded type cannot be a pointer to interface")
		}
	}

	if t.IsPtr() || t.IsUnsafePtr() {
		psess.
			yyerror("embedded type cannot be a pointer")
	} else if t.Etype == TFORW && !t.ForwardType(psess.types).Embedlineno.IsKnown() {
		t.ForwardType(psess.types).Embedlineno = psess.lineno
	}
}

func (psess *PackageSession) structfield(n *Node) *types.Field {
	lno := psess.lineno
	psess.
		lineno = n.Pos

	if n.Op != ODCLFIELD {
		psess.
			Fatalf("structfield: oops %v\n", n)
	}

	f := types.NewField()
	f.Pos = n.Pos
	f.Sym = n.Sym

	if n.Left != nil {
		n.Left = psess.typecheck(n.Left, Etype)
		n.Type = n.Left.Type
		n.Left = nil
	}

	f.Type = n.Type
	if f.Type == nil {
		f.SetBroke(true)
	}

	if n.Embedded() {
		psess.
			checkembeddedtype(n.Type)
		f.Embedded = 1
	} else {
		f.Embedded = 0
	}

	switch u := n.Val().U.(type) {
	case string:
		f.Note = u
	default:
		psess.
			yyerror("field tag must be a string")
	case nil:

	}
	psess.
		lineno = lno
	return f
}

// checkdupfields emits errors for duplicately named fields or methods in
// a list of struct or interface types.
func (psess *PackageSession) checkdupfields(what string, ts ...*types.Type) {
	seen := make(map[*types.Sym]bool)
	for _, t := range ts {
		for _, f := range t.Fields(psess.types).Slice() {
			if f.Sym == nil || f.Sym.IsBlank() {
				continue
			}
			if seen[f.Sym] {
				psess.
					yyerrorl(f.Pos, "duplicate %s %s", what, f.Sym.Name)
				continue
			}
			seen[f.Sym] = true
		}
	}
}

// convert a parsed id/type list into
// a type for struct/interface/arglist
func (psess *PackageSession) tostruct(l []*Node) *types.Type {
	t := types.New(TSTRUCT)
	psess.
		tostruct0(t, l)
	return t
}

func (psess *PackageSession) tostruct0(t *types.Type, l []*Node) {
	if t == nil || !t.IsStruct() {
		psess.
			Fatalf("struct expected")
	}

	fields := make([]*types.Field, len(l))
	for i, n := range l {
		f := psess.structfield(n)
		if f.Broke() {
			t.SetBroke(true)
		}
		fields[i] = f
	}
	t.SetFields(psess.types, fields)
	psess.
		checkdupfields("field", t)

	if !t.Broke() {
		psess.
			checkwidth(t)
	}
}

func (psess *PackageSession) tofunargs(l []*Node, funarg types.Funarg) *types.Type {
	t := types.New(TSTRUCT)
	t.StructType(psess.types).Funarg = funarg

	fields := make([]*types.Field, len(l))
	for i, n := range l {
		f := psess.structfield(n)
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
	t.SetFields(psess.types, fields)
	return t
}

func (psess *PackageSession) tofunargsfield(fields []*types.Field, funarg types.Funarg) *types.Type {
	t := types.New(TSTRUCT)
	t.StructType(psess.types).Funarg = funarg
	t.SetFields(psess.types, fields)
	return t
}

func (psess *PackageSession) interfacefield(n *Node) *types.Field {
	lno := psess.lineno
	psess.
		lineno = n.Pos

	if n.Op != ODCLFIELD {
		psess.
			Fatalf("interfacefield: oops %v\n", n)
	}

	if n.Val().Ctype(psess) != CTxxx {
		psess.
			yyerror("interface method cannot have annotation")
	}

	if n.Left != nil {
		n.Left = psess.typecheck(n.Left, Etype)
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
	psess.
		lineno = lno
	return f
}

func (psess *PackageSession) tointerface(l []*Node) *types.Type {
	if len(l) == 0 {
		return psess.types.Types[TINTER]
	}
	t := types.New(TINTER)
	psess.
		tointerface0(t, l)
	return t
}

func (psess *PackageSession) tointerface0(t *types.Type, l []*Node) {
	if t == nil || !t.IsInterface() {
		psess.
			Fatalf("interface expected")
	}

	var fields []*types.Field
	for _, n := range l {
		f := psess.interfacefield(n)
		if f.Broke() {
			t.SetBroke(true)
		}
		fields = append(fields, f)
	}
	t.SetInterface(psess.types, fields)
}

func (psess *PackageSession) fakeRecv() *Node {
	return psess.anonfield(psess.types.FakeRecvType())
}

func (psess *PackageSession) fakeRecvField() *types.Field {
	f := types.NewField()
	f.Type = psess.types.FakeRecvType()
	return f
}

// isifacemethod reports whether (field) m is
// an interface method. Such methods have the
// special receiver type types.FakeRecvType().
func (psess *PackageSession) isifacemethod(f *types.Type) bool {
	return f.Recv(psess.types).Type == psess.types.FakeRecvType()
}

// turn a parsed function declaration into a type
func (psess *PackageSession) functype(this *Node, in, out []*Node) *types.Type {
	t := types.New(TFUNC)
	psess.
		functype0(t, this, in, out)
	return t
}

func (psess *PackageSession) functype0(t *types.Type, this *Node, in, out []*Node) {
	if t == nil || t.Etype != TFUNC {
		psess.
			Fatalf("function type expected")
	}

	var rcvr []*Node
	if this != nil {
		rcvr = []*Node{this}
	}
	t.FuncType(psess.types).Receiver = psess.tofunargs(rcvr, types.FunargRcvr)
	t.FuncType(psess.types).Params = psess.tofunargs(in, types.FunargParams)
	t.FuncType(psess.types).Results = psess.tofunargs(out, types.FunargResults)
	psess.
		checkdupfields("argument", t.Recvs(psess.types), t.Params(psess.types), t.Results(psess.types))

	if t.Recvs(psess.types).Broke() || t.Results(psess.types).Broke() || t.Params(psess.types).Broke() {
		t.SetBroke(true)
	}

	t.FuncType(psess.types).Outnamed = t.NumResults(psess.types) > 0 && psess.origSym(t.Results(psess.types).Field(psess.types, 0).Sym) != nil
}

func (psess *PackageSession) functypefield(this *types.Field, in, out []*types.Field) *types.Type {
	t := types.New(TFUNC)
	psess.
		functypefield0(t, this, in, out)
	return t
}

func (psess *PackageSession) functypefield0(t *types.Type, this *types.Field, in, out []*types.Field) {
	var rcvr []*types.Field
	if this != nil {
		rcvr = []*types.Field{this}
	}
	t.FuncType(psess.types).Receiver = psess.tofunargsfield(rcvr, types.FunargRcvr)
	t.FuncType(psess.types).Params = psess.tofunargsfield(in, types.FunargParams)
	t.FuncType(psess.types).Results = psess.tofunargsfield(out, types.FunargResults)

	t.FuncType(psess.types).Outnamed = t.NumResults(psess.types) > 0 && psess.origSym(t.Results(psess.types).Field(psess.types, 0).Sym) != nil
}

// origSym returns the original symbol written by the user.
func (psess *PackageSession) origSym(s *types.Sym) *types.Sym {
	if s == nil {
		return nil
	}

	if len(s.Name) > 1 && s.Name[0] == '~' {
		switch s.Name[1] {
		case 'r':
			return nil
		case 'b':

			return psess.nblank.Sym
		}
		return s
	}

	if strings.HasPrefix(s.Name, ".anon") {

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
func (psess *PackageSession) methodSym(recv *types.Type, msym *types.Sym) *types.Sym {
	return psess.methodSymSuffix(recv, msym, "")
}

// methodSymSuffix is like methodsym, but allows attaching a
// distinguisher suffix. To avoid collisions, the suffix must not
// start with a letter, number, or period.
func (psess *PackageSession) methodSymSuffix(recv *types.Type, msym *types.Sym, suffix string) *types.Sym {
	if msym.IsBlank() {
		psess.
			Fatalf("blank method name")
	}

	rsym := recv.Sym
	if recv.IsPtr() {
		if rsym != nil {
			psess.
				Fatalf("declared pointer receiver type: %v", recv)
		}
		rsym = recv.Elem(psess.types).Sym
	}

	rpkg := psess.gopkg
	if rsym != nil {
		rpkg = rsym.Pkg
	}

	var b bytes.Buffer
	if recv.IsPtr() {

		fmt.Fprintf(&b, "(%-S)", recv)
	} else {
		fmt.Fprintf(&b, "%-S", recv)
	}

	if !types.IsExported(msym.Name) && msym.Pkg != rpkg {
		b.WriteString(".")
		b.WriteString(msym.Pkg.Prefix)
	}

	b.WriteString(".")
	b.WriteString(msym.Name)
	b.WriteString(suffix)

	return rpkg.LookupBytes(psess.types, b.Bytes())
}

// Add a method, declared as a function.
// - msym is the method symbol
// - t is function type (with receiver)
// Returns a pointer to the existing or added Field.
func (psess *PackageSession) addmethod(msym *types.Sym, t *types.Type, local, nointerface bool) *types.Field {
	if msym == nil {
		psess.
			Fatalf("no method symbol")
	}

	rf := t.Recv(psess.types)
	if rf == nil {
		psess.
			yyerror("missing receiver")
		return nil
	}

	mt := psess.methtype(rf.Type)
	if mt == nil || mt.Sym == nil {
		pa := rf.Type
		t := pa
		if t != nil && t.IsPtr() {
			if t.Sym != nil {
				psess.
					yyerror("invalid receiver type %v (%v is a pointer type)", pa, t)
				return nil
			}
			t = t.Elem(psess.types)
		}

		switch {
		case t == nil || t.Broke():

		case t.Sym == nil:
			psess.
				yyerror("invalid receiver type %v (%v is not a defined type)", pa, t)
		case t.IsPtr():
			psess.
				yyerror("invalid receiver type %v (%v is a pointer type)", pa, t)
		case t.IsInterface():
			psess.
				yyerror("invalid receiver type %v (%v is an interface type)", pa, t)
		default:
			psess.
				yyerror("invalid receiver type %v (%L / %L)", pa, pa, t)
		}
		return nil
	}

	if local && mt.Sym.Pkg != psess.localpkg {
		psess.
			yyerror("cannot define new methods on non-local type %v", mt)
		return nil
	}

	if msym.IsBlank() {
		return nil
	}

	if mt.IsStruct() {
		for _, f := range mt.Fields(psess.types).Slice() {
			if f.Sym == msym {
				psess.
					yyerror("type %v has both field and method named %v", mt, msym)
				return nil
			}
		}
	}

	for _, f := range mt.Methods().Slice() {
		if msym.Name != f.Sym.Name {
			continue
		}

		if !psess.eqtype(t, f.Type) || !psess.eqtype(t.Recv(psess.types).Type, f.Type.Recv(psess.types).Type) {
			psess.
				yyerror("method redeclared: %v.%v\n\t%v\n\t%v", mt, msym, f.Type, t)
		}
		return f
	}

	f := types.NewField()
	f.Pos = psess.lineno
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
func (psess *PackageSession) funcsym(s *types.Sym) *types.Sym {
	psess.
		funcsymsmu.Lock()
	sf, existed := s.Pkg.LookupOK(psess.types, funcsymname(s))

	if !psess.Ctxt.Flag_dynlink && !existed {
		psess.
			funcsyms = append(psess.funcsyms, s)
	}
	psess.
		funcsymsmu.Unlock()
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
func (psess *PackageSession) makefuncsym(s *types.Sym) {
	if !psess.Ctxt.Flag_dynlink {
		psess.
			Fatalf("makefuncsym dynlink")
	}
	if s.IsBlank() {
		return
	}
	if psess.compiling_runtime && (s.Name == "getg" || s.Name == "getclosureptr" || s.Name == "getcallerpc" || s.Name == "getcallersp") {

		return
	}
	if _, existed := s.Pkg.LookupOK(psess.types, funcsymname(s)); !existed {
		psess.
			funcsyms = append(psess.funcsyms, s)
	}
}

// disableExport prevents sym from being included in package export
// data. To be effectual, it must be called before declare.
func disableExport(sym *types.Sym) {
	sym.SetOnExportList(true)
}

func (psess *PackageSession) dclfunc(sym *types.Sym, tfn *Node) *Node {
	if tfn.Op != OTFUNC {
		psess.
			Fatalf("expected OTFUNC node, got %v", tfn)
	}

	fn := psess.nod(ODCLFUNC, nil, nil)
	fn.Func.Nname = psess.newfuncname(sym)
	fn.Func.Nname.Name.Defn = fn
	fn.Func.Nname.Name.Param.Ntype = tfn
	psess.
		declare(fn.Func.Nname, PFUNC)
	psess.
		funchdr(fn)
	fn.Func.Nname.Name.Param.Ntype = psess.typecheck(fn.Func.Nname.Name.Param.Ntype, Etype)
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
func (psess *PackageSession) newNowritebarrierrecChecker() *nowritebarrierrecChecker {
	c := &nowritebarrierrecChecker{
		extraCalls: make(map[*Node][]nowritebarrierrecCall),
	}

	for _, n := range psess.xtop {
		if n.Op != ODCLFUNC {
			continue
		}
		c.curfn = n
		inspect(n, c.findExtraCalls)
	}
	c.curfn = nil
	return c
}

func (c *nowritebarrierrecChecker) findExtraCalls(psess *PackageSession, n *Node) bool {
	if n.Op != OCALLFUNC {
		return true
	}
	fn := n.Left
	if fn == nil || fn.Op != ONAME || fn.Class() != PFUNC || fn.Name.Defn == nil {
		return true
	}
	if !psess.isRuntimePkg(fn.Sym.Pkg) || fn.Sym.Name != "systemstack" {
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
		psess.
			Fatalf("expected ONAME or OCLOSURE node, got %+v", arg)
	}
	if callee.Op != ODCLFUNC {
		psess.
			Fatalf("expected ODCLFUNC node, got %+v", callee)
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
func (c *nowritebarrierrecChecker) recordCall(psess *PackageSession, from *Node, to *obj.LSym, pos src.XPos) {
	if from.Op != ODCLFUNC {
		psess.
			Fatalf("expected ODCLFUNC, got %v", from)
	}

	fn := from.Func
	if fn.nwbrCalls == nil {
		fn.nwbrCalls = new([]nowritebarrierrecCallSym)
	}
	*fn.nwbrCalls = append(*fn.nwbrCalls, nowritebarrierrecCallSym{to, pos})
}

func (c *nowritebarrierrecChecker) check(psess *PackageSession) {

	symToFunc := make(map[*obj.LSym]*Node)

	funcs := make(map[*Node]nowritebarrierrecCall)
	// q is the queue of ODCLFUNC Nodes to visit in BFS order.
	var q nodeQueue

	for _, n := range psess.xtop {
		if n.Op != ODCLFUNC {
			continue
		}

		symToFunc[n.Func.lsym] = n

		if n.Func.Pragma&Nowritebarrierrec != 0 {
			funcs[n] = nowritebarrierrecCall{}
			q.pushRight(n)
		}

		if n.Func.Pragma&Nowritebarrier != 0 && n.Func.WBPos.IsKnown() {
			psess.
				yyerrorl(n.Func.WBPos, "write barrier prohibited")
		}
	}

	enqueue := func(src, target *Node, pos src.XPos) {
		if target.Func.Pragma&Yeswritebarrierrec != 0 {

			return
		}
		if _, ok := funcs[target]; ok {

			return
		}

		funcs[target] = nowritebarrierrecCall{target: src, lineno: pos}
		q.pushRight(target)
	}
	for !q.empty() {
		fn := q.popLeft()

		if fn.Func.WBPos.IsKnown() {
			var err bytes.Buffer
			call := funcs[fn]
			for call.target != nil {
				fmt.Fprintf(&err, "\n\t%v: called by %v", psess.linestr(call.lineno), call.target.Func.Nname)
				call = funcs[call.target]
			}
			psess.
				yyerrorl(fn.Func.WBPos, "write barrier prohibited by caller; %v%s", fn.Func.Nname, err.String())
			continue
		}

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

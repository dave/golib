package gc

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/dave/golib/src/cmd/compile/internal/syntax"
	"github.com/dave/golib/src/cmd/compile/internal/types"

	"github.com/dave/golib/src/cmd/internal/src"
)

// parseFiles concurrently parses files into *syntax.File structures.
// Each declaration in every *syntax.File is converted to a syntax tree
// and its root represented by *Node is appended to xtop.
// Returns the total count of parsed lines.
func (psess *PackageSession) parseFiles(filenames []string) uint {
	var noders []*noder

	sem := make(chan struct{}, runtime.GOMAXPROCS(0)+10)

	for _, filename := range filenames {
		p := &noder{
			basemap: make(map[*syntax.PosBase]*src.PosBase),
			err:     make(chan syntax.Error),
		}
		noders = append(noders, p)

		go func(filename string) {
			sem <- struct{}{}
			defer func() { <-sem }()
			defer close(p.err)
			base := syntax.NewFileBase(filename)

			f, err := os.Open(filename)
			if err != nil {
				p.error(syntax.Error{Pos: syntax.MakePos(base, 0, 0), Msg: err.Error()})
				return
			}
			defer f.Close()

			p.file, _ = psess.syntax.Parse(base, f, p.error, p.pragma, syntax.CheckBranches)
		}(filename)
	}

	var lines uint
	for _, p := range noders {
		for e := range p.err {
			p.yyerrorpos(psess, e.Pos, "%s", e.Msg)
		}

		p.node(psess)
		lines += p.file.Lines
		p.file = nil

		if psess.nsyntaxerrors != 0 {
			psess.
				errorexit()
		}
		psess.
			testdclstack()
	}
	psess.
		localpkg.Height = psess.myheight

	return lines
}

// makeSrcPosBase translates from a *syntax.PosBase to a *src.PosBase.
func (p *noder) makeSrcPosBase(psess *PackageSession, b0 *syntax.PosBase) *src.PosBase {

	if p.basecache.last == b0 {
		return p.basecache.base
	}

	b1, ok := p.basemap[b0]
	if !ok {
		fn := b0.Filename()
		if b0.IsFileBase() {
			b1 = src.NewFileBase(fn, psess.absFilename(fn))
		} else {

			p0 := b0.Pos()
			p1 := src.MakePos(p.makeSrcPosBase(psess, p0.Base()), p0.Line(), p0.Col())
			b1 = src.NewLinePragmaBase(p1, fn, psess.fileh(fn), b0.Line(), b0.Col())
		}
		p.basemap[b0] = b1
	}

	p.basecache.last = b0
	p.basecache.base = b1

	return b1
}

func (p *noder) makeXPos(psess *PackageSession, pos syntax.Pos) (_ src.XPos) {
	return psess.Ctxt.PosTable.XPos(src.MakePos(p.makeSrcPosBase(psess, pos.Base()), pos.Line(), pos.Col()))
}

func (p *noder) yyerrorpos(psess *PackageSession, pos syntax.Pos, format string, args ...interface{}) {
	psess.
		yyerrorl(p.makeXPos(psess, pos), format, args...)
}

// TODO(gri) Can we eliminate fileh in favor of absFilename?
func (psess *PackageSession) fileh(name string) string {
	return psess.objabi.AbsFile("", name, psess.pathPrefix)
}

func (psess *PackageSession) absFilename(name string) string {
	return psess.objabi.AbsFile(psess.Ctxt.Pathname, name, psess.pathPrefix)
}

// noder transforms package syntax's AST into a Node tree.
type noder struct {
	basemap   map[*syntax.PosBase]*src.PosBase
	basecache struct {
		last *syntax.PosBase
		base *src.PosBase
	}

	file       *syntax.File
	linknames  []linkname
	pragcgobuf [][]string
	err        chan syntax.Error
	scope      ScopeID

	// scopeVars is a stack tracking the number of variables declared in the
	// current function at the moment each open scope was opened.
	scopeVars []int

	lastCloseScopePos syntax.Pos
}

func (p *noder) funcBody(psess *PackageSession, fn *Node, block *syntax.BlockStmt) {
	oldScope := p.scope
	p.scope = 0
	psess.
		funchdr(fn)

	if block != nil {
		body := p.stmts(psess, block.List)
		if body == nil {
			body = []*Node{psess.nod(OEMPTY, nil, nil)}
		}
		fn.Nbody.Set(body)
		psess.
			lineno = p.makeXPos(psess, block.Rbrace)
		fn.Func.Endlineno = psess.lineno
	}
	psess.
		funcbody()
	p.scope = oldScope
}

func (p *noder) openScope(psess *PackageSession, pos syntax.Pos) {
	psess.types.
		Markdcl()

	if psess.trackScopes {
		psess.
			Curfn.Func.Parents = append(psess.Curfn.Func.Parents, p.scope)
		p.scopeVars = append(p.scopeVars, len(psess.Curfn.Func.Dcl))
		p.scope = ScopeID(len(psess.Curfn.Func.Parents))

		p.markScope(psess, pos)
	}
}

func (p *noder) closeScope(psess *PackageSession, pos syntax.Pos) {
	p.lastCloseScopePos = pos
	psess.types.
		Popdcl()

	if psess.trackScopes {
		scopeVars := p.scopeVars[len(p.scopeVars)-1]
		p.scopeVars = p.scopeVars[:len(p.scopeVars)-1]
		if scopeVars == len(psess.Curfn.Func.Dcl) {

			if int(p.scope) != len(psess.Curfn.Func.Parents) {
				psess.
					Fatalf("scope tracking inconsistency, no variables declared but scopes were not retracted")
			}

			p.scope = psess.Curfn.Func.Parents[p.scope-1]
			psess.
				Curfn.Func.Parents = psess.Curfn.Func.Parents[:len(psess.Curfn.Func.Parents)-1]

			nmarks := len(psess.Curfn.Func.Marks)
			psess.
				Curfn.Func.Marks[nmarks-1].Scope = p.scope
			prevScope := ScopeID(0)
			if nmarks >= 2 {
				prevScope = psess.Curfn.Func.Marks[nmarks-2].Scope
			}
			if psess.Curfn.Func.Marks[nmarks-1].Scope == prevScope {
				psess.
					Curfn.Func.Marks = psess.Curfn.Func.Marks[:nmarks-1]
			}
			return
		}

		p.scope = psess.Curfn.Func.Parents[p.scope-1]

		p.markScope(psess, pos)
	}
}

func (p *noder) markScope(psess *PackageSession, pos syntax.Pos) {
	xpos := p.makeXPos(psess, pos)
	if i := len(psess.Curfn.Func.Marks); i > 0 && psess.Curfn.Func.Marks[i-1].Pos == xpos {
		psess.
			Curfn.Func.Marks[i-1].Scope = p.scope
	} else {
		psess.
			Curfn.Func.Marks = append(psess.Curfn.Func.Marks, Mark{xpos, p.scope})
	}
}

// closeAnotherScope is like closeScope, but it reuses the same mark
// position as the last closeScope call. This is useful for "for" and
// "if" statements, as their implicit blocks always end at the same
// position as an explicit block.
func (p *noder) closeAnotherScope(psess *PackageSession) {
	p.closeScope(psess, p.lastCloseScopePos)
}

// linkname records a //go:linkname directive.
type linkname struct {
	pos    syntax.Pos
	local  string
	remote string
}

func (p *noder) node(psess *PackageSession) {
	psess.types.
		Block = 1
	psess.
		imported_unsafe = false

	p.lineno(psess, p.file.PkgName)
	psess.
		mkpackage(p.file.PkgName.Value)
	psess.
		xtop = append(psess.xtop, p.decls(psess, p.file.DeclList)...)

	for _, n := range p.linknames {
		if psess.imported_unsafe {
			psess.
				lookup(n.local).Linkname = n.remote
		} else {
			p.yyerrorpos(psess, n.pos, "//go:linkname only allowed in Go files that import \"unsafe\"")
		}
	}
	psess.
		pragcgobuf = append(psess.pragcgobuf, p.pragcgobuf...)
	psess.
		lineno = psess.src.NoXPos
	psess.
		clearImports()
}

func (p *noder) decls(psess *PackageSession, decls []syntax.Decl) (l []*Node) {
	var cs constState

	for _, decl := range decls {
		p.lineno(psess, decl)
		switch decl := decl.(type) {
		case *syntax.ImportDecl:
			p.importDecl(psess, decl)

		case *syntax.VarDecl:
			l = append(l, p.varDecl(psess, decl)...)

		case *syntax.ConstDecl:
			l = append(l, p.constDecl(psess, decl, &cs)...)

		case *syntax.TypeDecl:
			l = append(l, p.typeDecl(psess, decl))

		case *syntax.FuncDecl:
			l = append(l, p.funcDecl(psess, decl))

		default:
			panic("unhandled Decl")
		}
	}

	return
}

func (p *noder) importDecl(psess *PackageSession, imp *syntax.ImportDecl) {
	val := p.basicLit(psess, imp.Path)
	ipkg := psess.importfile(&val)

	if ipkg == nil {
		if psess.nerrors == 0 {
			psess.
				Fatalf("phase error in import")
		}
		return
	}

	ipkg.Direct = true

	var my *types.Sym
	if imp.LocalPkgName != nil {
		my = p.name(psess, imp.LocalPkgName)
	} else {
		my = psess.lookup(ipkg.Name)
	}

	pack := p.nod(psess, imp, OPACK, nil, nil)
	pack.Sym = my
	pack.Name.Pkg = ipkg

	switch my.Name {
	case ".":
		psess.
			importdot(ipkg, pack)
		return
	case "init":
		psess.
			yyerrorl(pack.Pos, "cannot import package as init - init must be a func")
		return
	case "_":
		return
	}
	if my.Def != nil {
		psess.
			redeclare(pack.Pos, my, "as imported package name")
	}
	my.Def = asTypesNode(pack)
	my.Lastlineno = pack.Pos
	my.Block = 1
}

func (p *noder) varDecl(psess *PackageSession, decl *syntax.VarDecl) []*Node {
	names := p.declNames(psess, decl.NameList)
	typ := p.typeExprOrNil(psess, decl.Type)

	var exprs []*Node
	if decl.Values != nil {
		exprs = p.exprList(psess, decl.Values)
	}

	p.lineno(psess, decl)
	return psess.variter(names, typ, exprs)
}

// constState tracks state between constant specifiers within a
// declaration group. This state is kept separate from noder so nested
// constant declarations are handled correctly (e.g., issue 15550).
type constState struct {
	group  *syntax.Group
	typ    *Node
	values []*Node
	iota   int64
}

func (p *noder) constDecl(psess *PackageSession, decl *syntax.ConstDecl, cs *constState) []*Node {
	if decl.Group == nil || decl.Group != cs.group {
		*cs = constState{
			group: decl.Group,
		}
	}

	names := p.declNames(psess, decl.NameList)
	typ := p.typeExprOrNil(psess, decl.Type)

	var values []*Node
	if decl.Values != nil {
		values = p.exprList(psess, decl.Values)
		cs.typ, cs.values = typ, values
	} else {
		if typ != nil {
			psess.
				yyerror("const declaration cannot have type without expression")
		}
		typ, values = cs.typ, cs.values
	}

	var nn []*Node
	for i, n := range names {
		if i >= len(values) {
			psess.
				yyerror("missing value in const declaration")
			break
		}
		v := values[i]
		if decl.Values == nil {
			v = psess.treecopy(v, n.Pos)
		}

		n.Op = OLITERAL
		psess.
			declare(n, psess.dclcontext)

		n.Name.Param.Ntype = typ
		n.Name.Defn = v
		n.SetIota(cs.iota)

		nn = append(nn, p.nod(psess, decl, ODCLCONST, n, nil))
	}

	if len(values) > len(names) {
		psess.
			yyerror("extra expression in const declaration")
	}

	cs.iota++

	return nn
}

func (p *noder) typeDecl(psess *PackageSession, decl *syntax.TypeDecl) *Node {
	n := p.declName(psess, decl.Name)
	n.Op = OTYPE
	psess.
		declare(n, psess.dclcontext)

	typ := p.typeExprOrNil(psess, decl.Type)

	param := n.Name.Param
	param.Ntype = typ
	param.Pragma = decl.Pragma
	param.Alias = decl.Alias
	if param.Alias && param.Pragma != 0 {
		psess.
			yyerror("cannot specify directive with type alias")
		param.Pragma = 0
	}

	return p.nod(psess, decl, ODCLTYPE, n, nil)

}

func (p *noder) declNames(psess *PackageSession, names []*syntax.Name) []*Node {
	var nodes []*Node
	for _, name := range names {
		nodes = append(nodes, p.declName(psess, name))
	}
	return nodes
}

func (p *noder) declName(psess *PackageSession, name *syntax.Name) *Node {
	return p.setlineno(psess, name, psess.dclname(p.name(psess, name)))
}

func (p *noder) funcDecl(psess *PackageSession, fun *syntax.FuncDecl) *Node {
	name := p.name(psess, fun.Name)
	t := p.signature(psess, fun.Recv, fun.Type)
	f := p.nod(psess, fun, ODCLFUNC, nil, nil)

	if fun.Recv == nil {
		if name.Name == "init" {
			name = psess.renameinit()
			if t.List.Len() > 0 || t.Rlist.Len() > 0 {
				psess.
					yyerrorl(f.Pos, "func init must have no arguments and no return values")
			}
		}

		if psess.localpkg.Name == "main" && name.Name == "main" {
			if t.List.Len() > 0 || t.Rlist.Len() > 0 {
				psess.
					yyerrorl(f.Pos, "func main must have no arguments and no return values")
			}
		}
	} else {
		f.Func.Shortname = name
		name = psess.nblank.Sym
	}

	f.Func.Nname = p.setlineno(psess, fun.Name, psess.newfuncname(name))
	f.Func.Nname.Name.Defn = f
	f.Func.Nname.Name.Param.Ntype = t

	pragma := fun.Pragma
	f.Func.Pragma = fun.Pragma
	f.SetNoescape(pragma&Noescape != 0)
	if pragma&Systemstack != 0 && pragma&Nosplit != 0 {
		psess.
			yyerrorl(f.Pos, "go:nosplit and go:systemstack cannot be combined")
	}

	if fun.Recv == nil {
		psess.
			declare(f.Func.Nname, PFUNC)
	}

	p.funcBody(psess, f, fun.Body)

	if fun.Body != nil {
		if f.Noescape() {
			psess.
				yyerrorl(f.Pos, "can only use //go:noescape with external func implementations")
		}
	} else {
		if psess.pure_go || strings.HasPrefix(f.funcname(), "init.") {
			psess.
				yyerrorl(f.Pos, "missing function body")
		}
	}

	return f
}

func (p *noder) signature(psess *PackageSession, recv *syntax.Field, typ *syntax.FuncType) *Node {
	n := p.nod(psess, typ, OTFUNC, nil, nil)
	if recv != nil {
		n.Left = p.param(psess, recv, false, false)
	}
	n.List.Set(p.params(psess, typ.ParamList, true))
	n.Rlist.Set(p.params(psess, typ.ResultList, false))
	return n
}

func (p *noder) params(psess *PackageSession, params []*syntax.Field, dddOk bool) []*Node {
	var nodes []*Node
	for i, param := range params {
		p.lineno(psess, param)
		nodes = append(nodes, p.param(psess, param, dddOk, i+1 == len(params)))
	}
	return nodes
}

func (p *noder) param(psess *PackageSession, param *syntax.Field, dddOk, final bool) *Node {
	var name *types.Sym
	if param.Name != nil {
		name = p.name(psess, param.Name)
	}

	typ := p.typeExpr(psess, param.Type)
	n := p.nodSym(psess, param, ODCLFIELD, typ, name)

	if typ.Op == ODDD {
		if !dddOk {
			psess.
				yyerror("cannot use ... in receiver or result parameter list")
		} else if !final {
			psess.
				yyerror("can only use ... with final parameter in list")
		}
		typ.Op = OTARRAY
		typ.Right = typ.Left
		typ.Left = nil
		n.SetIsddd(true)
		if n.Left != nil {
			n.Left.SetIsddd(true)
		}
	}

	return n
}

func (p *noder) exprList(psess *PackageSession, expr syntax.Expr) []*Node {
	if list, ok := expr.(*syntax.ListExpr); ok {
		return p.exprs(psess, list.ElemList)
	}
	return []*Node{p.expr(psess, expr)}
}

func (p *noder) exprs(psess *PackageSession, exprs []syntax.Expr) []*Node {
	var nodes []*Node
	for _, expr := range exprs {
		nodes = append(nodes, p.expr(psess, expr))
	}
	return nodes
}

func (p *noder) expr(psess *PackageSession, expr syntax.Expr) *Node {
	p.lineno(psess, expr)
	switch expr := expr.(type) {
	case nil, *syntax.BadExpr:
		return nil
	case *syntax.Name:
		return p.mkname(psess, expr)
	case *syntax.BasicLit:
		return p.setlineno(psess, expr, psess.nodlit(p.basicLit(psess, expr)))

	case *syntax.CompositeLit:
		n := p.nod(psess, expr, OCOMPLIT, nil, nil)
		if expr.Type != nil {
			n.Right = p.expr(psess, expr.Type)
		}
		l := p.exprs(psess, expr.ElemList)
		for i, e := range l {
			l[i] = p.wrapname(psess, expr.ElemList[i], e)
		}
		n.List.Set(l)
		psess.
			lineno = p.makeXPos(psess, expr.Rbrace)
		return n
	case *syntax.KeyValueExpr:

		return p.nod(psess, expr.Key, OKEY, p.expr(psess, expr.Key), p.wrapname(psess, expr.Value, p.expr(psess, expr.Value)))
	case *syntax.FuncLit:
		return p.funcLit(psess, expr)
	case *syntax.ParenExpr:
		return p.nod(psess, expr, OPAREN, p.expr(psess, expr.X), nil)
	case *syntax.SelectorExpr:

		obj := p.expr(psess, expr.X)
		if obj.Op == OPACK {
			obj.Name.SetUsed(true)
			return psess.oldname(psess.restrictlookup(expr.Sel.Value, obj.Name.Pkg))
		}
		return p.setlineno(psess, expr, psess.nodSym(OXDOT, obj, p.name(psess, expr.Sel)))
	case *syntax.IndexExpr:
		return p.nod(psess, expr, OINDEX, p.expr(psess, expr.X), p.expr(psess, expr.Index))
	case *syntax.SliceExpr:
		op := OSLICE
		if expr.Full {
			op = OSLICE3
		}
		n := p.nod(psess, expr, op, p.expr(psess, expr.X), nil)
		var index [3]*Node
		for i, x := range expr.Index {
			if x != nil {
				index[i] = p.expr(psess, x)
			}
		}
		n.SetSliceBounds(psess, index[0], index[1], index[2])
		return n
	case *syntax.AssertExpr:
		return p.nod(psess, expr, ODOTTYPE, p.expr(psess, expr.X), p.typeExpr(psess, expr.Type))
	case *syntax.Operation:
		if expr.Op == syntax.Add && expr.Y != nil {
			return p.sum(psess, expr)
		}
		x := p.expr(psess, expr.X)
		if expr.Y == nil {
			if expr.Op == syntax.And {
				x = unparen(x)
				if x.Op == OCOMPLIT {

					x.Right = p.nod(psess, expr, OIND, x.Right, nil)
					x.Right.SetImplicit(true)
					return x
				}
			}
			return p.nod(psess, expr, p.unOp(psess, expr.Op), x, nil)
		}
		return p.nod(psess, expr, p.binOp(psess, expr.Op), x, p.expr(psess, expr.Y))
	case *syntax.CallExpr:
		n := p.nod(psess, expr, OCALL, p.expr(psess, expr.Fun), nil)
		n.List.Set(p.exprs(psess, expr.ArgList))
		n.SetIsddd(expr.HasDots)
		return n

	case *syntax.ArrayType:
		var len *Node
		if expr.Len != nil {
			len = p.expr(psess, expr.Len)
		} else {
			len = p.nod(psess, expr, ODDD, nil, nil)
		}
		return p.nod(psess, expr, OTARRAY, len, p.typeExpr(psess, expr.Elem))
	case *syntax.SliceType:
		return p.nod(psess, expr, OTARRAY, nil, p.typeExpr(psess, expr.Elem))
	case *syntax.DotsType:
		return p.nod(psess, expr, ODDD, p.typeExpr(psess, expr.Elem), nil)
	case *syntax.StructType:
		return p.structType(psess, expr)
	case *syntax.InterfaceType:
		return p.interfaceType(psess, expr)
	case *syntax.FuncType:
		return p.signature(psess, nil, expr)
	case *syntax.MapType:
		return p.nod(psess, expr, OTMAP, p.typeExpr(psess, expr.Key), p.typeExpr(psess, expr.Value))
	case *syntax.ChanType:
		n := p.nod(psess, expr, OTCHAN, p.typeExpr(psess, expr.Elem), nil)
		n.SetTChanDir(psess, p.chanDir(expr.Dir))
		return n

	case *syntax.TypeSwitchGuard:
		n := p.nod(psess, expr, OTYPESW, nil, p.expr(psess, expr.X))
		if expr.Lhs != nil {
			n.Left = p.declName(psess, expr.Lhs)
			if n.Left.isBlank() {
				psess.
					yyerror("invalid variable name %v in type switch", n.Left)
			}
		}
		return n
	}
	panic("unhandled Expr")
}

// sum efficiently handles very large summation expressions (such as
// in issue #16394). In particular, it avoids left recursion and
// collapses string literals.
func (p *noder) sum(psess *PackageSession, x syntax.Expr) *Node {

	adds := make([]*syntax.Operation, 0, 2)
	for {
		add, ok := x.(*syntax.Operation)
		if !ok || add.Op != syntax.Add || add.Y == nil {
			break
		}
		adds = append(adds, add)
		x = add.X
	}

	// nstr is the current rightmost string literal in the
	// summation (if any), and chunks holds its accumulated
	// substrings.
	//
	// Consider the expression x + "a" + "b" + "c" + y. When we
	// reach the string literal "a", we assign nstr to point to
	// its corresponding Node and initialize chunks to {"a"}.
	// Visiting the subsequent string literals "b" and "c", we
	// simply append their values to chunks. Finally, when we
	// reach the non-constant operand y, we'll join chunks to form
	// "abc" and reassign the "a" string literal's value.
	//
	// N.B., we need to be careful about named string constants
	// (indicated by Sym != nil) because 1) we can't modify their
	// value, as doing so would affect other uses of the string
	// constant, and 2) they may have types, which we need to
	// handle correctly. For now, we avoid these problems by
	// treating named string constants the same as non-constant
	// operands.
	var nstr *Node
	chunks := make([]string, 0, 1)

	n := p.expr(psess, x)
	if psess.Isconst(n, CTSTR) && n.Sym == nil {
		nstr = n
		chunks = append(chunks, nstr.Val().U.(string))
	}

	for i := len(adds) - 1; i >= 0; i-- {
		add := adds[i]

		r := p.expr(psess, add.Y)
		if psess.Isconst(r, CTSTR) && r.Sym == nil {
			if nstr != nil {

				chunks = append(chunks, r.Val().U.(string))
				continue
			}

			nstr = r
			chunks = append(chunks, nstr.Val().U.(string))
		} else {
			if len(chunks) > 1 {
				nstr.SetVal(psess, Val{U: strings.Join(chunks, "")})
			}
			nstr = nil
			chunks = chunks[:0]
		}
		n = p.nod(psess, add, OADD, n, r)
	}
	if len(chunks) > 1 {
		nstr.SetVal(psess, Val{U: strings.Join(chunks, "")})
	}

	return n
}

func (p *noder) typeExpr(psess *PackageSession, typ syntax.Expr) *Node {

	return p.expr(psess, typ)
}

func (p *noder) typeExprOrNil(psess *PackageSession, typ syntax.Expr) *Node {
	if typ != nil {
		return p.expr(psess, typ)
	}
	return nil
}

func (p *noder) chanDir(dir syntax.ChanDir) types.ChanDir {
	switch dir {
	case 0:
		return types.Cboth
	case syntax.SendOnly:
		return types.Csend
	case syntax.RecvOnly:
		return types.Crecv
	}
	panic("unhandled ChanDir")
}

func (p *noder) structType(psess *PackageSession, expr *syntax.StructType) *Node {
	var l []*Node
	for i, field := range expr.FieldList {
		p.lineno(psess, field)
		var n *Node
		if field.Name == nil {
			n = p.embedded(psess, field.Type)
		} else {
			n = p.nodSym(psess, field, ODCLFIELD, p.typeExpr(psess, field.Type), p.name(psess, field.Name))
		}
		if i < len(expr.TagList) && expr.TagList[i] != nil {
			n.SetVal(psess, p.basicLit(psess, expr.TagList[i]))
		}
		l = append(l, n)
	}

	p.lineno(psess, expr)
	n := p.nod(psess, expr, OTSTRUCT, nil, nil)
	n.List.Set(l)
	return n
}

func (p *noder) interfaceType(psess *PackageSession, expr *syntax.InterfaceType) *Node {
	var l []*Node
	for _, method := range expr.MethodList {
		p.lineno(psess, method)
		var n *Node
		if method.Name == nil {
			n = p.nodSym(psess, method, ODCLFIELD, psess.oldname(p.packname(psess, method.Type)), nil)
		} else {
			mname := p.name(psess, method.Name)
			sig := p.typeExpr(psess, method.Type)
			sig.Left = psess.fakeRecv()
			n = p.nodSym(psess, method, ODCLFIELD, sig, mname)
			psess.
				ifacedcl(n)
		}
		l = append(l, n)
	}

	n := p.nod(psess, expr, OTINTER, nil, nil)
	n.List.Set(l)
	return n
}

func (p *noder) packname(psess *PackageSession, expr syntax.Expr) *types.Sym {
	switch expr := expr.(type) {
	case *syntax.Name:
		name := p.name(psess, expr)
		if n := psess.oldname(name); n.Name != nil && n.Name.Pack != nil {
			n.Name.Pack.Name.SetUsed(true)
		}
		return name
	case *syntax.SelectorExpr:
		name := p.name(psess, expr.X.(*syntax.Name))
		var pkg *types.Pkg
		if asNode(name.Def) == nil || asNode(name.Def).Op != OPACK {
			psess.
				yyerror("%v is not a package", name)
			pkg = psess.localpkg
		} else {
			asNode(name.Def).Name.SetUsed(true)
			pkg = asNode(name.Def).Name.Pkg
		}
		return psess.restrictlookup(expr.Sel.Value, pkg)
	}
	panic(fmt.Sprintf("unexpected packname: %#v", expr))
}

func (p *noder) embedded(psess *PackageSession, typ syntax.Expr) *Node {
	op, isStar := typ.(*syntax.Operation)
	if isStar {
		if op.Op != syntax.Mul || op.Y != nil {
			panic("unexpected Operation")
		}
		typ = op.X
	}

	sym := p.packname(psess, typ)
	n := p.nodSym(psess, typ, ODCLFIELD, psess.oldname(sym), psess.lookup(sym.Name))
	n.SetEmbedded(true)

	if isStar {
		n.Left = p.nod(psess, op, OIND, n.Left, nil)
	}
	return n
}

func (p *noder) stmts(psess *PackageSession, stmts []syntax.Stmt) []*Node {
	return p.stmtsFall(psess, stmts, false)
}

func (p *noder) stmtsFall(psess *PackageSession, stmts []syntax.Stmt, fallOK bool) []*Node {
	var nodes []*Node
	for i, stmt := range stmts {
		s := p.stmtFall(psess, stmt, fallOK && i+1 == len(stmts))
		if s == nil {
		} else if s.Op == OBLOCK && s.Ninit.Len() == 0 {
			nodes = append(nodes, s.List.Slice()...)
		} else {
			nodes = append(nodes, s)
		}
	}
	return nodes
}

func (p *noder) stmt(psess *PackageSession, stmt syntax.Stmt) *Node {
	return p.stmtFall(psess, stmt, false)
}

func (p *noder) stmtFall(psess *PackageSession, stmt syntax.Stmt, fallOK bool) *Node {
	p.lineno(psess, stmt)
	switch stmt := stmt.(type) {
	case *syntax.EmptyStmt:
		return nil
	case *syntax.LabeledStmt:
		return p.labeledStmt(psess, stmt, fallOK)
	case *syntax.BlockStmt:
		l := p.blockStmt(psess, stmt)
		if len(l) == 0 {

			return psess.nod(OEMPTY, nil, nil)
		}
		return psess.liststmt(l)
	case *syntax.ExprStmt:
		return p.wrapname(psess, stmt, p.expr(psess, stmt.X))
	case *syntax.SendStmt:
		return p.nod(psess, stmt, OSEND, p.expr(psess, stmt.Chan), p.expr(psess, stmt.Value))
	case *syntax.DeclStmt:
		return psess.liststmt(p.decls(psess, stmt.DeclList))
	case *syntax.AssignStmt:
		if stmt.Op != 0 && stmt.Op != syntax.Def {
			n := p.nod(psess, stmt, OASOP, p.expr(psess, stmt.Lhs), p.expr(psess, stmt.Rhs))
			n.SetImplicit(stmt.Rhs == psess.syntax.ImplicitOne)
			n.SetSubOp(psess, p.binOp(psess, stmt.Op))
			return n
		}

		n := p.nod(psess, stmt, OAS, nil, nil)

		rhs := p.exprList(psess, stmt.Rhs)
		lhs := p.assignList(psess, stmt.Lhs, n, stmt.Op == syntax.Def)

		if len(lhs) == 1 && len(rhs) == 1 {

			n.Left = lhs[0]
			n.Right = rhs[0]
		} else {
			n.Op = OAS2
			n.List.Set(lhs)
			n.Rlist.Set(rhs)
		}
		return n

	case *syntax.BranchStmt:
		var op Op
		switch stmt.Tok {
		case syntax.Break:
			op = OBREAK
		case syntax.Continue:
			op = OCONTINUE
		case syntax.Fallthrough:
			if !fallOK {
				psess.
					yyerror("fallthrough statement out of place")
			}
			op = OFALL
		case syntax.Goto:
			op = OGOTO
		default:
			panic("unhandled BranchStmt")
		}
		n := p.nod(psess, stmt, op, nil, nil)
		if stmt.Label != nil {
			n.Left = p.newname(psess, stmt.Label)
		}
		return n
	case *syntax.CallStmt:
		var op Op
		switch stmt.Tok {
		case syntax.Defer:
			op = ODEFER
		case syntax.Go:
			op = OPROC
		default:
			panic("unhandled CallStmt")
		}
		return p.nod(psess, stmt, op, p.expr(psess, stmt.Call), nil)
	case *syntax.ReturnStmt:
		var results []*Node
		if stmt.Results != nil {
			results = p.exprList(psess, stmt.Results)
		}
		n := p.nod(psess, stmt, ORETURN, nil, nil)
		n.List.Set(results)
		if n.List.Len() == 0 && psess.Curfn != nil {
			for _, ln := range psess.Curfn.Func.Dcl {
				if ln.Class() == PPARAM {
					continue
				}
				if ln.Class() != PPARAMOUT {
					break
				}
				if asNode(ln.Sym.Def) != ln {
					psess.
						yyerror("%s is shadowed during return", ln.Sym.Name)
				}
			}
		}
		return n
	case *syntax.IfStmt:
		return p.ifStmt(psess, stmt)
	case *syntax.ForStmt:
		return p.forStmt(psess, stmt)
	case *syntax.SwitchStmt:
		return p.switchStmt(psess, stmt)
	case *syntax.SelectStmt:
		return p.selectStmt(psess, stmt)
	}
	panic("unhandled Stmt")
}

func (p *noder) assignList(psess *PackageSession, expr syntax.Expr, defn *Node, colas bool) []*Node {
	if !colas {
		return p.exprList(psess, expr)
	}

	defn.SetColas(true)

	var exprs []syntax.Expr
	if list, ok := expr.(*syntax.ListExpr); ok {
		exprs = list.ElemList
	} else {
		exprs = []syntax.Expr{expr}
	}

	res := make([]*Node, len(exprs))
	seen := make(map[*types.Sym]bool, len(exprs))

	newOrErr := false
	for i, expr := range exprs {
		p.lineno(psess, expr)
		res[i] = psess.nblank

		name, ok := expr.(*syntax.Name)
		if !ok {
			p.yyerrorpos(psess, expr.Pos(), "non-name %v on left side of :=", p.expr(psess, expr))
			newOrErr = true
			continue
		}

		sym := p.name(psess, name)
		if sym.IsBlank() {
			continue
		}

		if seen[sym] {
			p.yyerrorpos(psess, expr.Pos(), "%v repeated on left side of :=", sym)
			newOrErr = true
			continue
		}
		seen[sym] = true

		if sym.Block == psess.types.Block {
			res[i] = psess.oldname(sym)
			continue
		}

		newOrErr = true
		n := psess.newname(sym)
		psess.
			declare(n, psess.dclcontext)
		n.Name.Defn = defn
		defn.Ninit.Append(psess.nod(ODCL, n, nil))
		res[i] = n
	}

	if !newOrErr {
		psess.
			yyerrorl(defn.Pos, "no new variables on left side of :=")
	}
	return res
}

func (p *noder) blockStmt(psess *PackageSession, stmt *syntax.BlockStmt) []*Node {
	p.openScope(psess, stmt.Pos())
	nodes := p.stmts(psess, stmt.List)
	p.closeScope(psess, stmt.Rbrace)
	return nodes
}

func (p *noder) ifStmt(psess *PackageSession, stmt *syntax.IfStmt) *Node {
	p.openScope(psess, stmt.Pos())
	n := p.nod(psess, stmt, OIF, nil, nil)
	if stmt.Init != nil {
		n.Ninit.Set1(p.stmt(psess, stmt.Init))
	}
	if stmt.Cond != nil {
		n.Left = p.expr(psess, stmt.Cond)
	}
	n.Nbody.Set(p.blockStmt(psess, stmt.Then))
	if stmt.Else != nil {
		e := p.stmt(psess, stmt.Else)
		if e.Op == OBLOCK && e.Ninit.Len() == 0 {
			n.Rlist.Set(e.List.Slice())
		} else {
			n.Rlist.Set1(e)
		}
	}
	p.closeAnotherScope(psess)
	return n
}

func (p *noder) forStmt(psess *PackageSession, stmt *syntax.ForStmt) *Node {
	p.openScope(psess, stmt.Pos())
	var n *Node
	if r, ok := stmt.Init.(*syntax.RangeClause); ok {
		if stmt.Cond != nil || stmt.Post != nil {
			panic("unexpected RangeClause")
		}

		n = p.nod(psess, r, ORANGE, nil, p.expr(psess, r.X))
		if r.Lhs != nil {
			n.List.Set(p.assignList(psess, r.Lhs, n, r.Def))
		}
	} else {
		n = p.nod(psess, stmt, OFOR, nil, nil)
		if stmt.Init != nil {
			n.Ninit.Set1(p.stmt(psess, stmt.Init))
		}
		if stmt.Cond != nil {
			n.Left = p.expr(psess, stmt.Cond)
		}
		if stmt.Post != nil {
			n.Right = p.stmt(psess, stmt.Post)
		}
	}
	n.Nbody.Set(p.blockStmt(psess, stmt.Body))
	p.closeAnotherScope(psess)
	return n
}

func (p *noder) switchStmt(psess *PackageSession, stmt *syntax.SwitchStmt) *Node {
	p.openScope(psess, stmt.Pos())
	n := p.nod(psess, stmt, OSWITCH, nil, nil)
	if stmt.Init != nil {
		n.Ninit.Set1(p.stmt(psess, stmt.Init))
	}
	if stmt.Tag != nil {
		n.Left = p.expr(psess, stmt.Tag)
	}

	tswitch := n.Left
	if tswitch != nil && tswitch.Op != OTYPESW {
		tswitch = nil
	}
	n.List.Set(p.caseClauses(psess, stmt.Body, tswitch, stmt.Rbrace))

	p.closeScope(psess, stmt.Rbrace)
	return n
}

func (p *noder) caseClauses(psess *PackageSession, clauses []*syntax.CaseClause, tswitch *Node, rbrace syntax.Pos) []*Node {
	var nodes []*Node
	for i, clause := range clauses {
		p.lineno(psess, clause)
		if i > 0 {
			p.closeScope(psess, clause.Pos())
		}
		p.openScope(psess, clause.Pos())

		n := p.nod(psess, clause, OXCASE, nil, nil)
		if clause.Cases != nil {
			n.List.Set(p.exprList(psess, clause.Cases))
		}
		if tswitch != nil && tswitch.Left != nil {
			nn := psess.newname(tswitch.Left.Sym)
			psess.
				declare(nn, psess.dclcontext)
			n.Rlist.Set1(nn)

			nn.Name.Defn = tswitch
		}

		body := clause.Body
		for len(body) > 0 {
			if _, ok := body[len(body)-1].(*syntax.EmptyStmt); !ok {
				break
			}
			body = body[:len(body)-1]
		}

		n.Nbody.Set(p.stmtsFall(psess, body, true))
		if l := n.Nbody.Len(); l > 0 && n.Nbody.Index(l-1).Op == OFALL {
			if tswitch != nil {
				psess.
					yyerror("cannot fallthrough in type switch")
			}
			if i+1 == len(clauses) {
				psess.
					yyerror("cannot fallthrough final case in switch")
			}
		}

		nodes = append(nodes, n)
	}
	if len(clauses) > 0 {
		p.closeScope(psess, rbrace)
	}
	return nodes
}

func (p *noder) selectStmt(psess *PackageSession, stmt *syntax.SelectStmt) *Node {
	n := p.nod(psess, stmt, OSELECT, nil, nil)
	n.List.Set(p.commClauses(psess, stmt.Body, stmt.Rbrace))
	return n
}

func (p *noder) commClauses(psess *PackageSession, clauses []*syntax.CommClause, rbrace syntax.Pos) []*Node {
	var nodes []*Node
	for i, clause := range clauses {
		p.lineno(psess, clause)
		if i > 0 {
			p.closeScope(psess, clause.Pos())
		}
		p.openScope(psess, clause.Pos())

		n := p.nod(psess, clause, OXCASE, nil, nil)
		if clause.Comm != nil {
			n.List.Set1(p.stmt(psess, clause.Comm))
		}
		n.Nbody.Set(p.stmts(psess, clause.Body))
		nodes = append(nodes, n)
	}
	if len(clauses) > 0 {
		p.closeScope(psess, rbrace)
	}
	return nodes
}

func (p *noder) labeledStmt(psess *PackageSession, label *syntax.LabeledStmt, fallOK bool) *Node {
	lhs := p.nod(psess, label, OLABEL, p.newname(psess, label.Label), nil)

	var ls *Node
	if label.Stmt != nil {
		ls = p.stmtFall(psess, label.Stmt, fallOK)
	}

	lhs.Name.Defn = ls
	l := []*Node{lhs}
	if ls != nil {
		if ls.Op == OBLOCK && ls.Ninit.Len() == 0 {
			l = append(l, ls.List.Slice()...)
		} else {
			l = append(l, ls)
		}
	}
	return psess.liststmt(l)
}

func (p *noder) unOp(psess *PackageSession, op syntax.Operator) Op {
	if uint64(op) >= uint64(len(psess.unOps)) || psess.unOps[op] == 0 {
		panic("invalid Operator")
	}
	return psess.unOps[op]
}

func (p *noder) binOp(psess *PackageSession, op syntax.Operator) Op {
	if uint64(op) >= uint64(len(psess.binOps)) || psess.binOps[op] == 0 {
		panic("invalid Operator")
	}
	return psess.binOps[op]
}

func (p *noder) basicLit(psess *PackageSession, lit *syntax.BasicLit) Val {

	switch s := lit.Value; lit.Kind {
	case syntax.IntLit:
		x := new(Mpint)
		x.SetString(psess, s)
		return Val{U: x}

	case syntax.FloatLit:
		x := newMpflt()
		x.SetString(psess, s)
		return Val{U: x}

	case syntax.ImagLit:
		x := new(Mpcplx)
		x.Imag.SetString(psess, strings.TrimSuffix(s, "i"))
		return Val{U: x}

	case syntax.RuneLit:
		var r rune
		if u, err := strconv.Unquote(s); err == nil && len(u) > 0 {

			if len(u) == 1 {
				r = rune(u[0])
			} else {
				r, _ = utf8.DecodeRuneInString(u)
			}
		}
		x := new(Mpint)
		x.SetInt64(int64(r))
		x.Rune = true
		return Val{U: x}

	case syntax.StringLit:
		if len(s) > 0 && s[0] == '`' {

			s = strings.Replace(s, "\r", "", -1)
		}

		u, _ := strconv.Unquote(s)
		return Val{U: u}

	default:
		panic("unhandled BasicLit kind")
	}
}

func (p *noder) name(psess *PackageSession, name *syntax.Name) *types.Sym {
	return psess.lookup(name.Value)
}

func (p *noder) mkname(psess *PackageSession, name *syntax.Name) *Node {

	return psess.mkname(p.name(psess, name))
}

func (p *noder) newname(psess *PackageSession, name *syntax.Name) *Node {

	return psess.newname(p.name(psess, name))
}

func (p *noder) wrapname(psess *PackageSession, n syntax.Node, x *Node) *Node {

	switch x.Op {
	case OTYPE, OLITERAL:
		if x.Sym == nil {
			break
		}
		fallthrough
	case ONAME, ONONAME, OPACK:
		x = p.nod(psess, n, OPAREN, x, nil)
		x.SetImplicit(true)
	}
	return x
}

func (p *noder) nod(psess *PackageSession, orig syntax.Node, op Op, left, right *Node) *Node {
	return p.setlineno(psess, orig, psess.nod(op, left, right))
}

func (p *noder) nodSym(psess *PackageSession, orig syntax.Node, op Op, left *Node, sym *types.Sym) *Node {
	return p.setlineno(psess, orig, psess.nodSym(op, left, sym))
}

func (p *noder) setlineno(psess *PackageSession, src_ syntax.Node, dst *Node) *Node {
	pos := src_.Pos()
	if !pos.IsKnown() {

		return dst
	}
	dst.Pos = p.makeXPos(psess, pos)
	return dst
}

func (p *noder) lineno(psess *PackageSession, n syntax.Node) {
	if n == nil {
		return
	}
	pos := n.Pos()
	if !pos.IsKnown() {

		return
	}
	psess.
		lineno = p.makeXPos(psess, pos)
}

// error is called concurrently if files are parsed concurrently.
func (p *noder) error(err error) {
	p.err <- err.(syntax.Error)
}

// pragmas that are allowed in the std lib, but don't have
// a syntax.Pragma value (see lex.go) associated with them.

// pragma is called concurrently if files are parsed concurrently.
func (p *noder) pragma(psess *PackageSession, pos syntax.Pos, text string) syntax.Pragma {
	switch {
	case strings.HasPrefix(text, "line "):

		panic("unreachable")

	case strings.HasPrefix(text, "go:linkname "):
		f := strings.Fields(text)
		if len(f) != 3 {
			p.error(syntax.Error{Pos: pos, Msg: "usage: //go:linkname localname linkname"})
			break
		}
		p.linknames = append(p.linknames, linkname{pos, f[1], f[2]})

	case strings.HasPrefix(text, "go:cgo_import_dynamic "):

		fields := pragmaFields(text)
		if len(fields) >= 4 {
			lib := strings.Trim(fields[3], "\"")
			if lib != "" && !safeArg(lib) && !psess.isCgoGeneratedFile(pos) {
				p.error(syntax.Error{Pos: pos, Msg: fmt.Sprintf("invalid library name %q in cgo_import_dynamic directive", lib)})
			}
			p.pragcgo(pos, text)
			return psess.pragmaValue("go:cgo_import_dynamic")
		}
		fallthrough
	case strings.HasPrefix(text, "go:cgo_"):

		if !psess.isCgoGeneratedFile(pos) && !psess.compiling_std {
			p.error(syntax.Error{Pos: pos, Msg: fmt.Sprintf("//%s only allowed in cgo-generated code", text)})
		}
		p.pragcgo(pos, text)
		fallthrough
	default:
		verb := text
		if i := strings.Index(text, " "); i >= 0 {
			verb = verb[:i]
		}
		prag := psess.pragmaValue(verb)
		const runtimePragmas = Systemstack | Nowritebarrier | Nowritebarrierrec | Yeswritebarrierrec
		if !psess.compiling_runtime && prag&runtimePragmas != 0 {
			p.error(syntax.Error{Pos: pos, Msg: fmt.Sprintf("//%s only allowed in runtime", verb)})
		}
		if prag == 0 && !psess.allowedStdPragmas[verb] && psess.compiling_std {
			p.error(syntax.Error{Pos: pos, Msg: fmt.Sprintf("//%s is not allowed in the standard library", verb)})
		}
		return prag
	}

	return 0
}

// isCgoGeneratedFile reports whether pos is in a file
// generated by cgo, which is to say a file with name
// beginning with "_cgo_". Such files are allowed to
// contain cgo directives, and for security reasons
// (primarily misuse of linker flags), other files are not.
// See golang.org/issue/23672.
func (psess *PackageSession) isCgoGeneratedFile(pos syntax.Pos) bool {
	return strings.HasPrefix(filepath.Base(filepath.Clean(psess.fileh(pos.Base().Filename()))), "_cgo_")
}

// safeArg reports whether arg is a "safe" command-line argument,
// meaning that when it appears in a command-line, it probably
// doesn't have some special meaning other than its own name.
// This is copied from SafeArg in cmd/go/internal/load/pkg.go.
func safeArg(name string) bool {
	if name == "" {
		return false
	}
	c := name[0]
	return '0' <= c && c <= '9' || 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || c == '.' || c == '_' || c == '/' || c >= utf8.RuneSelf
}

func (psess *PackageSession) mkname(sym *types.Sym) *Node {
	n := psess.oldname(sym)
	if n.Name != nil && n.Name.Pack != nil {
		n.Name.Pack.Name.SetUsed(true)
	}
	return n
}

func unparen(x *Node) *Node {
	for x.Op == OPAREN {
		x = x.Left
	}
	return x
}

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
)

// parseFiles concurrently parses files into *syntax.File structures.
// Each declaration in every *syntax.File is converted to a syntax tree
// and its root represented by *Node is appended to xtop.
// Returns the total count of parsed lines.
func (pstate *PackageState) parseFiles(filenames []string) uint {
	var noders []*noder
	// Limit the number of simultaneously open files.
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

			p.file, _ = pstate.syntax.Parse(base, f, p.error, p.pragma, syntax.CheckBranches) // errors are tracked via p.error
		}(filename)
	}

	var lines uint
	for _, p := range noders {
		for e := range p.err {
			p.yyerrorpos(pstate, e.Pos, "%s", e.Msg)
		}

		p.node(pstate)
		lines += p.file.Lines
		p.file = nil // release memory

		if pstate.nsyntaxerrors != 0 {
			pstate.errorexit()
		}
		// Always run testdclstack here, even when debug_dclstack is not set, as a sanity measure.
		pstate.testdclstack()
	}

	pstate.localpkg.Height = pstate.myheight

	return lines
}

// makeSrcPosBase translates from a *syntax.PosBase to a *src.PosBase.
func (p *noder) makeSrcPosBase(pstate *PackageState, b0 *syntax.PosBase) *src.PosBase {
	// fast path: most likely PosBase hasn't changed
	if p.basecache.last == b0 {
		return p.basecache.base
	}

	b1, ok := p.basemap[b0]
	if !ok {
		fn := b0.Filename()
		if b0.IsFileBase() {
			b1 = src.NewFileBase(fn, pstate.absFilename(fn))
		} else {
			// line directive base
			p0 := b0.Pos()
			p1 := src.MakePos(p.makeSrcPosBase(pstate, p0.Base()), p0.Line(), p0.Col())
			b1 = src.NewLinePragmaBase(p1, fn, pstate.fileh(fn), b0.Line(), b0.Col())
		}
		p.basemap[b0] = b1
	}

	// update cache
	p.basecache.last = b0
	p.basecache.base = b1

	return b1
}

func (p *noder) makeXPos(pstate *PackageState, pos syntax.Pos) (_ src.XPos) {
	return pstate.Ctxt.PosTable.XPos(src.MakePos(p.makeSrcPosBase(pstate, pos.Base()), pos.Line(), pos.Col()))
}

func (p *noder) yyerrorpos(pstate *PackageState, pos syntax.Pos, format string, args ...interface{}) {
	pstate.yyerrorl(p.makeXPos(pstate, pos), format, args...)
}

// TODO(gri) Can we eliminate fileh in favor of absFilename?
func (pstate *PackageState) fileh(name string) string {
	return pstate.objabi.AbsFile("", name, pstate.pathPrefix)
}

func (pstate *PackageState) absFilename(name string) string {
	return pstate.objabi.AbsFile(pstate.Ctxt.Pathname, name, pstate.pathPrefix)
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

func (p *noder) funcBody(pstate *PackageState, fn *Node, block *syntax.BlockStmt) {
	oldScope := p.scope
	p.scope = 0
	pstate.funchdr(fn)

	if block != nil {
		body := p.stmts(pstate, block.List)
		if body == nil {
			body = []*Node{pstate.nod(OEMPTY, nil, nil)}
		}
		fn.Nbody.Set(body)

		pstate.lineno = p.makeXPos(pstate, block.Rbrace)
		fn.Func.Endlineno = pstate.lineno
	}

	pstate.funcbody()
	p.scope = oldScope
}

func (p *noder) openScope(pstate *PackageState, pos syntax.Pos) {
	pstate.types.Markdcl()

	if pstate.trackScopes {
		pstate.Curfn.Func.Parents = append(pstate.Curfn.Func.Parents, p.scope)
		p.scopeVars = append(p.scopeVars, len(pstate.Curfn.Func.Dcl))
		p.scope = ScopeID(len(pstate.Curfn.Func.Parents))

		p.markScope(pstate, pos)
	}
}

func (p *noder) closeScope(pstate *PackageState, pos syntax.Pos) {
	p.lastCloseScopePos = pos
	pstate.types.Popdcl()

	if pstate.trackScopes {
		scopeVars := p.scopeVars[len(p.scopeVars)-1]
		p.scopeVars = p.scopeVars[:len(p.scopeVars)-1]
		if scopeVars == len(pstate.Curfn.Func.Dcl) {
			// no variables were declared in this scope, so we can retract it.

			if int(p.scope) != len(pstate.Curfn.Func.Parents) {
				pstate.Fatalf("scope tracking inconsistency, no variables declared but scopes were not retracted")
			}

			p.scope = pstate.Curfn.Func.Parents[p.scope-1]
			pstate.Curfn.Func.Parents = pstate.Curfn.Func.Parents[:len(pstate.Curfn.Func.Parents)-1]

			nmarks := len(pstate.Curfn.Func.Marks)
			pstate.Curfn.Func.Marks[nmarks-1].Scope = p.scope
			prevScope := ScopeID(0)
			if nmarks >= 2 {
				prevScope = pstate.Curfn.Func.Marks[nmarks-2].Scope
			}
			if pstate.Curfn.Func.Marks[nmarks-1].Scope == prevScope {
				pstate.Curfn.Func.Marks = pstate.Curfn.Func.Marks[:nmarks-1]
			}
			return
		}

		p.scope = pstate.Curfn.Func.Parents[p.scope-1]

		p.markScope(pstate, pos)
	}
}

func (p *noder) markScope(pstate *PackageState, pos syntax.Pos) {
	xpos := p.makeXPos(pstate, pos)
	if i := len(pstate.Curfn.Func.Marks); i > 0 && pstate.Curfn.Func.Marks[i-1].Pos == xpos {
		pstate.Curfn.Func.Marks[i-1].Scope = p.scope
	} else {
		pstate.Curfn.Func.Marks = append(pstate.Curfn.Func.Marks, Mark{xpos, p.scope})
	}
}

// closeAnotherScope is like closeScope, but it reuses the same mark
// position as the last closeScope call. This is useful for "for" and
// "if" statements, as their implicit blocks always end at the same
// position as an explicit block.
func (p *noder) closeAnotherScope(pstate *PackageState) {
	p.closeScope(pstate, p.lastCloseScopePos)
}

// linkname records a //go:linkname directive.
type linkname struct {
	pos    syntax.Pos
	local  string
	remote string
}

func (p *noder) node(pstate *PackageState) {
	pstate.types.Block = 1
	pstate.imported_unsafe = false

	p.lineno(pstate, p.file.PkgName)
	pstate.mkpackage(p.file.PkgName.Value)

	pstate.xtop = append(pstate.xtop, p.decls(pstate, p.file.DeclList)...)

	for _, n := range p.linknames {
		if pstate.imported_unsafe {
			pstate.lookup(n.local).Linkname = n.remote
		} else {
			p.yyerrorpos(pstate, n.pos, "//go:linkname only allowed in Go files that import \"unsafe\"")
		}
	}

	pstate.pragcgobuf = append(pstate.pragcgobuf, p.pragcgobuf...)
	pstate.lineno = pstate.src.NoXPos
	pstate.clearImports()
}

func (p *noder) decls(pstate *PackageState, decls []syntax.Decl) (l []*Node) {
	var cs constState

	for _, decl := range decls {
		p.lineno(pstate, decl)
		switch decl := decl.(type) {
		case *syntax.ImportDecl:
			p.importDecl(pstate, decl)

		case *syntax.VarDecl:
			l = append(l, p.varDecl(pstate, decl)...)

		case *syntax.ConstDecl:
			l = append(l, p.constDecl(pstate, decl, &cs)...)

		case *syntax.TypeDecl:
			l = append(l, p.typeDecl(pstate, decl))

		case *syntax.FuncDecl:
			l = append(l, p.funcDecl(pstate, decl))

		default:
			panic("unhandled Decl")
		}
	}

	return
}

func (p *noder) importDecl(pstate *PackageState, imp *syntax.ImportDecl) {
	val := p.basicLit(pstate, imp.Path)
	ipkg := pstate.importfile(&val)

	if ipkg == nil {
		if pstate.nerrors == 0 {
			pstate.Fatalf("phase error in import")
		}
		return
	}

	ipkg.Direct = true

	var my *types.Sym
	if imp.LocalPkgName != nil {
		my = p.name(pstate, imp.LocalPkgName)
	} else {
		my = pstate.lookup(ipkg.Name)
	}

	pack := p.nod(pstate, imp, OPACK, nil, nil)
	pack.Sym = my
	pack.Name.Pkg = ipkg

	switch my.Name {
	case ".":
		pstate.importdot(ipkg, pack)
		return
	case "init":
		pstate.yyerrorl(pack.Pos, "cannot import package as init - init must be a func")
		return
	case "_":
		return
	}
	if my.Def != nil {
		pstate.redeclare(pack.Pos, my, "as imported package name")
	}
	my.Def = asTypesNode(pack)
	my.Lastlineno = pack.Pos
	my.Block = 1 // at top level
}

func (p *noder) varDecl(pstate *PackageState, decl *syntax.VarDecl) []*Node {
	names := p.declNames(pstate, decl.NameList)
	typ := p.typeExprOrNil(pstate, decl.Type)

	var exprs []*Node
	if decl.Values != nil {
		exprs = p.exprList(pstate, decl.Values)
	}

	p.lineno(pstate, decl)
	return pstate.variter(names, typ, exprs)
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

func (p *noder) constDecl(pstate *PackageState, decl *syntax.ConstDecl, cs *constState) []*Node {
	if decl.Group == nil || decl.Group != cs.group {
		*cs = constState{
			group: decl.Group,
		}
	}

	names := p.declNames(pstate, decl.NameList)
	typ := p.typeExprOrNil(pstate, decl.Type)

	var values []*Node
	if decl.Values != nil {
		values = p.exprList(pstate, decl.Values)
		cs.typ, cs.values = typ, values
	} else {
		if typ != nil {
			pstate.yyerror("const declaration cannot have type without expression")
		}
		typ, values = cs.typ, cs.values
	}

	var nn []*Node
	for i, n := range names {
		if i >= len(values) {
			pstate.yyerror("missing value in const declaration")
			break
		}
		v := values[i]
		if decl.Values == nil {
			v = pstate.treecopy(v, n.Pos)
		}

		n.Op = OLITERAL
		pstate.declare(n, pstate.dclcontext)

		n.Name.Param.Ntype = typ
		n.Name.Defn = v
		n.SetIota(cs.iota)

		nn = append(nn, p.nod(pstate, decl, ODCLCONST, n, nil))
	}

	if len(values) > len(names) {
		pstate.yyerror("extra expression in const declaration")
	}

	cs.iota++

	return nn
}

func (p *noder) typeDecl(pstate *PackageState, decl *syntax.TypeDecl) *Node {
	n := p.declName(pstate, decl.Name)
	n.Op = OTYPE
	pstate.declare(n, pstate.dclcontext)

	// decl.Type may be nil but in that case we got a syntax error during parsing
	typ := p.typeExprOrNil(pstate, decl.Type)

	param := n.Name.Param
	param.Ntype = typ
	param.Pragma = decl.Pragma
	param.Alias = decl.Alias
	if param.Alias && param.Pragma != 0 {
		pstate.yyerror("cannot specify directive with type alias")
		param.Pragma = 0
	}

	return p.nod(pstate, decl, ODCLTYPE, n, nil)

}

func (p *noder) declNames(pstate *PackageState, names []*syntax.Name) []*Node {
	var nodes []*Node
	for _, name := range names {
		nodes = append(nodes, p.declName(pstate, name))
	}
	return nodes
}

func (p *noder) declName(pstate *PackageState, name *syntax.Name) *Node {
	return p.setlineno(pstate, name, pstate.dclname(p.name(pstate, name)))
}

func (p *noder) funcDecl(pstate *PackageState, fun *syntax.FuncDecl) *Node {
	name := p.name(pstate, fun.Name)
	t := p.signature(pstate, fun.Recv, fun.Type)
	f := p.nod(pstate, fun, ODCLFUNC, nil, nil)

	if fun.Recv == nil {
		if name.Name == "init" {
			name = pstate.renameinit()
			if t.List.Len() > 0 || t.Rlist.Len() > 0 {
				pstate.yyerrorl(f.Pos, "func init must have no arguments and no return values")
			}
		}

		if pstate.localpkg.Name == "main" && name.Name == "main" {
			if t.List.Len() > 0 || t.Rlist.Len() > 0 {
				pstate.yyerrorl(f.Pos, "func main must have no arguments and no return values")
			}
		}
	} else {
		f.Func.Shortname = name
		name = pstate.nblank.Sym // filled in by typecheckfunc
	}

	f.Func.Nname = p.setlineno(pstate, fun.Name, pstate.newfuncname(name))
	f.Func.Nname.Name.Defn = f
	f.Func.Nname.Name.Param.Ntype = t

	pragma := fun.Pragma
	f.Func.Pragma = fun.Pragma
	f.SetNoescape(pragma&Noescape != 0)
	if pragma&Systemstack != 0 && pragma&Nosplit != 0 {
		pstate.yyerrorl(f.Pos, "go:nosplit and go:systemstack cannot be combined")
	}

	if fun.Recv == nil {
		pstate.declare(f.Func.Nname, PFUNC)
	}

	p.funcBody(pstate, f, fun.Body)

	if fun.Body != nil {
		if f.Noescape() {
			pstate.yyerrorl(f.Pos, "can only use //go:noescape with external func implementations")
		}
	} else {
		if pstate.pure_go || strings.HasPrefix(f.funcname(), "init.") {
			pstate.yyerrorl(f.Pos, "missing function body")
		}
	}

	return f
}

func (p *noder) signature(pstate *PackageState, recv *syntax.Field, typ *syntax.FuncType) *Node {
	n := p.nod(pstate, typ, OTFUNC, nil, nil)
	if recv != nil {
		n.Left = p.param(pstate, recv, false, false)
	}
	n.List.Set(p.params(pstate, typ.ParamList, true))
	n.Rlist.Set(p.params(pstate, typ.ResultList, false))
	return n
}

func (p *noder) params(pstate *PackageState, params []*syntax.Field, dddOk bool) []*Node {
	var nodes []*Node
	for i, param := range params {
		p.lineno(pstate, param)
		nodes = append(nodes, p.param(pstate, param, dddOk, i+1 == len(params)))
	}
	return nodes
}

func (p *noder) param(pstate *PackageState, param *syntax.Field, dddOk, final bool) *Node {
	var name *types.Sym
	if param.Name != nil {
		name = p.name(pstate, param.Name)
	}

	typ := p.typeExpr(pstate, param.Type)
	n := p.nodSym(pstate, param, ODCLFIELD, typ, name)

	// rewrite ...T parameter
	if typ.Op == ODDD {
		if !dddOk {
			pstate.yyerror("cannot use ... in receiver or result parameter list")
		} else if !final {
			pstate.yyerror("can only use ... with final parameter in list")
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

func (p *noder) exprList(pstate *PackageState, expr syntax.Expr) []*Node {
	if list, ok := expr.(*syntax.ListExpr); ok {
		return p.exprs(pstate, list.ElemList)
	}
	return []*Node{p.expr(pstate, expr)}
}

func (p *noder) exprs(pstate *PackageState, exprs []syntax.Expr) []*Node {
	var nodes []*Node
	for _, expr := range exprs {
		nodes = append(nodes, p.expr(pstate, expr))
	}
	return nodes
}

func (p *noder) expr(pstate *PackageState, expr syntax.Expr) *Node {
	p.lineno(pstate, expr)
	switch expr := expr.(type) {
	case nil, *syntax.BadExpr:
		return nil
	case *syntax.Name:
		return p.mkname(pstate, expr)
	case *syntax.BasicLit:
		return p.setlineno(pstate, expr, pstate.nodlit(p.basicLit(pstate, expr)))

	case *syntax.CompositeLit:
		n := p.nod(pstate, expr, OCOMPLIT, nil, nil)
		if expr.Type != nil {
			n.Right = p.expr(pstate, expr.Type)
		}
		l := p.exprs(pstate, expr.ElemList)
		for i, e := range l {
			l[i] = p.wrapname(pstate, expr.ElemList[i], e)
		}
		n.List.Set(l)
		pstate.lineno = p.makeXPos(pstate, expr.Rbrace)
		return n
	case *syntax.KeyValueExpr:
		// use position of expr.Key rather than of expr (which has position of ':')
		return p.nod(pstate, expr.Key, OKEY, p.expr(pstate, expr.Key), p.wrapname(pstate, expr.Value, p.expr(pstate, expr.Value)))
	case *syntax.FuncLit:
		return p.funcLit(pstate, expr)
	case *syntax.ParenExpr:
		return p.nod(pstate, expr, OPAREN, p.expr(pstate, expr.X), nil)
	case *syntax.SelectorExpr:
		// parser.new_dotname
		obj := p.expr(pstate, expr.X)
		if obj.Op == OPACK {
			obj.Name.SetUsed(true)
			return pstate.oldname(pstate.restrictlookup(expr.Sel.Value, obj.Name.Pkg))
		}
		return p.setlineno(pstate, expr, pstate.nodSym(OXDOT, obj, p.name(pstate, expr.Sel)))
	case *syntax.IndexExpr:
		return p.nod(pstate, expr, OINDEX, p.expr(pstate, expr.X), p.expr(pstate, expr.Index))
	case *syntax.SliceExpr:
		op := OSLICE
		if expr.Full {
			op = OSLICE3
		}
		n := p.nod(pstate, expr, op, p.expr(pstate, expr.X), nil)
		var index [3]*Node
		for i, x := range expr.Index {
			if x != nil {
				index[i] = p.expr(pstate, x)
			}
		}
		n.SetSliceBounds(pstate, index[0], index[1], index[2])
		return n
	case *syntax.AssertExpr:
		return p.nod(pstate, expr, ODOTTYPE, p.expr(pstate, expr.X), p.typeExpr(pstate, expr.Type))
	case *syntax.Operation:
		if expr.Op == syntax.Add && expr.Y != nil {
			return p.sum(pstate, expr)
		}
		x := p.expr(pstate, expr.X)
		if expr.Y == nil {
			if expr.Op == syntax.And {
				x = unparen(x) // TODO(mdempsky): Needed?
				if x.Op == OCOMPLIT {
					// Special case for &T{...}: turn into (*T){...}.
					x.Right = p.nod(pstate, expr, OIND, x.Right, nil)
					x.Right.SetImplicit(true)
					return x
				}
			}
			return p.nod(pstate, expr, p.unOp(pstate, expr.Op), x, nil)
		}
		return p.nod(pstate, expr, p.binOp(pstate, expr.Op), x, p.expr(pstate, expr.Y))
	case *syntax.CallExpr:
		n := p.nod(pstate, expr, OCALL, p.expr(pstate, expr.Fun), nil)
		n.List.Set(p.exprs(pstate, expr.ArgList))
		n.SetIsddd(expr.HasDots)
		return n

	case *syntax.ArrayType:
		var len *Node
		if expr.Len != nil {
			len = p.expr(pstate, expr.Len)
		} else {
			len = p.nod(pstate, expr, ODDD, nil, nil)
		}
		return p.nod(pstate, expr, OTARRAY, len, p.typeExpr(pstate, expr.Elem))
	case *syntax.SliceType:
		return p.nod(pstate, expr, OTARRAY, nil, p.typeExpr(pstate, expr.Elem))
	case *syntax.DotsType:
		return p.nod(pstate, expr, ODDD, p.typeExpr(pstate, expr.Elem), nil)
	case *syntax.StructType:
		return p.structType(pstate, expr)
	case *syntax.InterfaceType:
		return p.interfaceType(pstate, expr)
	case *syntax.FuncType:
		return p.signature(pstate, nil, expr)
	case *syntax.MapType:
		return p.nod(pstate, expr, OTMAP, p.typeExpr(pstate, expr.Key), p.typeExpr(pstate, expr.Value))
	case *syntax.ChanType:
		n := p.nod(pstate, expr, OTCHAN, p.typeExpr(pstate, expr.Elem), nil)
		n.SetTChanDir(pstate, p.chanDir(expr.Dir))
		return n

	case *syntax.TypeSwitchGuard:
		n := p.nod(pstate, expr, OTYPESW, nil, p.expr(pstate, expr.X))
		if expr.Lhs != nil {
			n.Left = p.declName(pstate, expr.Lhs)
			if n.Left.isBlank() {
				pstate.yyerror("invalid variable name %v in type switch", n.Left)
			}
		}
		return n
	}
	panic("unhandled Expr")
}

// sum efficiently handles very large summation expressions (such as
// in issue #16394). In particular, it avoids left recursion and
// collapses string literals.
func (p *noder) sum(pstate *PackageState, x syntax.Expr) *Node {
	// While we need to handle long sums with asymptotic
	// efficiency, the vast majority of sums are very small: ~95%
	// have only 2 or 3 operands, and ~99% of string literals are
	// never concatenated.

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

	n := p.expr(pstate, x)
	if pstate.Isconst(n, CTSTR) && n.Sym == nil {
		nstr = n
		chunks = append(chunks, nstr.Val().U.(string))
	}

	for i := len(adds) - 1; i >= 0; i-- {
		add := adds[i]

		r := p.expr(pstate, add.Y)
		if pstate.Isconst(r, CTSTR) && r.Sym == nil {
			if nstr != nil {
				// Collapse r into nstr instead of adding to n.
				chunks = append(chunks, r.Val().U.(string))
				continue
			}

			nstr = r
			chunks = append(chunks, nstr.Val().U.(string))
		} else {
			if len(chunks) > 1 {
				nstr.SetVal(pstate, Val{U: strings.Join(chunks, "")})
			}
			nstr = nil
			chunks = chunks[:0]
		}
		n = p.nod(pstate, add, OADD, n, r)
	}
	if len(chunks) > 1 {
		nstr.SetVal(pstate, Val{U: strings.Join(chunks, "")})
	}

	return n
}

func (p *noder) typeExpr(pstate *PackageState, typ syntax.Expr) *Node {
	// TODO(mdempsky): Be stricter? typecheck should handle errors anyway.
	return p.expr(pstate, typ)
}

func (p *noder) typeExprOrNil(pstate *PackageState, typ syntax.Expr) *Node {
	if typ != nil {
		return p.expr(pstate, typ)
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

func (p *noder) structType(pstate *PackageState, expr *syntax.StructType) *Node {
	var l []*Node
	for i, field := range expr.FieldList {
		p.lineno(pstate, field)
		var n *Node
		if field.Name == nil {
			n = p.embedded(pstate, field.Type)
		} else {
			n = p.nodSym(pstate, field, ODCLFIELD, p.typeExpr(pstate, field.Type), p.name(pstate, field.Name))
		}
		if i < len(expr.TagList) && expr.TagList[i] != nil {
			n.SetVal(pstate, p.basicLit(pstate, expr.TagList[i]))
		}
		l = append(l, n)
	}

	p.lineno(pstate, expr)
	n := p.nod(pstate, expr, OTSTRUCT, nil, nil)
	n.List.Set(l)
	return n
}

func (p *noder) interfaceType(pstate *PackageState, expr *syntax.InterfaceType) *Node {
	var l []*Node
	for _, method := range expr.MethodList {
		p.lineno(pstate, method)
		var n *Node
		if method.Name == nil {
			n = p.nodSym(pstate, method, ODCLFIELD, pstate.oldname(p.packname(pstate, method.Type)), nil)
		} else {
			mname := p.name(pstate, method.Name)
			sig := p.typeExpr(pstate, method.Type)
			sig.Left = pstate.fakeRecv()
			n = p.nodSym(pstate, method, ODCLFIELD, sig, mname)
			pstate.ifacedcl(n)
		}
		l = append(l, n)
	}

	n := p.nod(pstate, expr, OTINTER, nil, nil)
	n.List.Set(l)
	return n
}

func (p *noder) packname(pstate *PackageState, expr syntax.Expr) *types.Sym {
	switch expr := expr.(type) {
	case *syntax.Name:
		name := p.name(pstate, expr)
		if n := pstate.oldname(name); n.Name != nil && n.Name.Pack != nil {
			n.Name.Pack.Name.SetUsed(true)
		}
		return name
	case *syntax.SelectorExpr:
		name := p.name(pstate, expr.X.(*syntax.Name))
		var pkg *types.Pkg
		if asNode(name.Def) == nil || asNode(name.Def).Op != OPACK {
			pstate.yyerror("%v is not a package", name)
			pkg = pstate.localpkg
		} else {
			asNode(name.Def).Name.SetUsed(true)
			pkg = asNode(name.Def).Name.Pkg
		}
		return pstate.restrictlookup(expr.Sel.Value, pkg)
	}
	panic(fmt.Sprintf("unexpected packname: %#v", expr))
}

func (p *noder) embedded(pstate *PackageState, typ syntax.Expr) *Node {
	op, isStar := typ.(*syntax.Operation)
	if isStar {
		if op.Op != syntax.Mul || op.Y != nil {
			panic("unexpected Operation")
		}
		typ = op.X
	}

	sym := p.packname(pstate, typ)
	n := p.nodSym(pstate, typ, ODCLFIELD, pstate.oldname(sym), pstate.lookup(sym.Name))
	n.SetEmbedded(true)

	if isStar {
		n.Left = p.nod(pstate, op, OIND, n.Left, nil)
	}
	return n
}

func (p *noder) stmts(pstate *PackageState, stmts []syntax.Stmt) []*Node {
	return p.stmtsFall(pstate, stmts, false)
}

func (p *noder) stmtsFall(pstate *PackageState, stmts []syntax.Stmt, fallOK bool) []*Node {
	var nodes []*Node
	for i, stmt := range stmts {
		s := p.stmtFall(pstate, stmt, fallOK && i+1 == len(stmts))
		if s == nil {
		} else if s.Op == OBLOCK && s.Ninit.Len() == 0 {
			nodes = append(nodes, s.List.Slice()...)
		} else {
			nodes = append(nodes, s)
		}
	}
	return nodes
}

func (p *noder) stmt(pstate *PackageState, stmt syntax.Stmt) *Node {
	return p.stmtFall(pstate, stmt, false)
}

func (p *noder) stmtFall(pstate *PackageState, stmt syntax.Stmt, fallOK bool) *Node {
	p.lineno(pstate, stmt)
	switch stmt := stmt.(type) {
	case *syntax.EmptyStmt:
		return nil
	case *syntax.LabeledStmt:
		return p.labeledStmt(pstate, stmt, fallOK)
	case *syntax.BlockStmt:
		l := p.blockStmt(pstate, stmt)
		if len(l) == 0 {
			// TODO(mdempsky): Line number?
			return pstate.nod(OEMPTY, nil, nil)
		}
		return pstate.liststmt(l)
	case *syntax.ExprStmt:
		return p.wrapname(pstate, stmt, p.expr(pstate, stmt.X))
	case *syntax.SendStmt:
		return p.nod(pstate, stmt, OSEND, p.expr(pstate, stmt.Chan), p.expr(pstate, stmt.Value))
	case *syntax.DeclStmt:
		return pstate.liststmt(p.decls(pstate, stmt.DeclList))
	case *syntax.AssignStmt:
		if stmt.Op != 0 && stmt.Op != syntax.Def {
			n := p.nod(pstate, stmt, OASOP, p.expr(pstate, stmt.Lhs), p.expr(pstate, stmt.Rhs))
			n.SetImplicit(stmt.Rhs == pstate.syntax.ImplicitOne)
			n.SetSubOp(pstate, p.binOp(pstate, stmt.Op))
			return n
		}

		n := p.nod(pstate, stmt, OAS, nil, nil) // assume common case

		rhs := p.exprList(pstate, stmt.Rhs)
		lhs := p.assignList(pstate, stmt.Lhs, n, stmt.Op == syntax.Def)

		if len(lhs) == 1 && len(rhs) == 1 {
			// common case
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
				pstate.yyerror("fallthrough statement out of place")
			}
			op = OFALL
		case syntax.Goto:
			op = OGOTO
		default:
			panic("unhandled BranchStmt")
		}
		n := p.nod(pstate, stmt, op, nil, nil)
		if stmt.Label != nil {
			n.Left = p.newname(pstate, stmt.Label)
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
		return p.nod(pstate, stmt, op, p.expr(pstate, stmt.Call), nil)
	case *syntax.ReturnStmt:
		var results []*Node
		if stmt.Results != nil {
			results = p.exprList(pstate, stmt.Results)
		}
		n := p.nod(pstate, stmt, ORETURN, nil, nil)
		n.List.Set(results)
		if n.List.Len() == 0 && pstate.Curfn != nil {
			for _, ln := range pstate.Curfn.Func.Dcl {
				if ln.Class() == PPARAM {
					continue
				}
				if ln.Class() != PPARAMOUT {
					break
				}
				if asNode(ln.Sym.Def) != ln {
					pstate.yyerror("%s is shadowed during return", ln.Sym.Name)
				}
			}
		}
		return n
	case *syntax.IfStmt:
		return p.ifStmt(pstate, stmt)
	case *syntax.ForStmt:
		return p.forStmt(pstate, stmt)
	case *syntax.SwitchStmt:
		return p.switchStmt(pstate, stmt)
	case *syntax.SelectStmt:
		return p.selectStmt(pstate, stmt)
	}
	panic("unhandled Stmt")
}

func (p *noder) assignList(pstate *PackageState, expr syntax.Expr, defn *Node, colas bool) []*Node {
	if !colas {
		return p.exprList(pstate, expr)
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
		p.lineno(pstate, expr)
		res[i] = pstate.nblank

		name, ok := expr.(*syntax.Name)
		if !ok {
			p.yyerrorpos(pstate, expr.Pos(), "non-name %v on left side of :=", p.expr(pstate, expr))
			newOrErr = true
			continue
		}

		sym := p.name(pstate, name)
		if sym.IsBlank() {
			continue
		}

		if seen[sym] {
			p.yyerrorpos(pstate, expr.Pos(), "%v repeated on left side of :=", sym)
			newOrErr = true
			continue
		}
		seen[sym] = true

		if sym.Block == pstate.types.Block {
			res[i] = pstate.oldname(sym)
			continue
		}

		newOrErr = true
		n := pstate.newname(sym)
		pstate.declare(n, pstate.dclcontext)
		n.Name.Defn = defn
		defn.Ninit.Append(pstate.nod(ODCL, n, nil))
		res[i] = n
	}

	if !newOrErr {
		pstate.yyerrorl(defn.Pos, "no new variables on left side of :=")
	}
	return res
}

func (p *noder) blockStmt(pstate *PackageState, stmt *syntax.BlockStmt) []*Node {
	p.openScope(pstate, stmt.Pos())
	nodes := p.stmts(pstate, stmt.List)
	p.closeScope(pstate, stmt.Rbrace)
	return nodes
}

func (p *noder) ifStmt(pstate *PackageState, stmt *syntax.IfStmt) *Node {
	p.openScope(pstate, stmt.Pos())
	n := p.nod(pstate, stmt, OIF, nil, nil)
	if stmt.Init != nil {
		n.Ninit.Set1(p.stmt(pstate, stmt.Init))
	}
	if stmt.Cond != nil {
		n.Left = p.expr(pstate, stmt.Cond)
	}
	n.Nbody.Set(p.blockStmt(pstate, stmt.Then))
	if stmt.Else != nil {
		e := p.stmt(pstate, stmt.Else)
		if e.Op == OBLOCK && e.Ninit.Len() == 0 {
			n.Rlist.Set(e.List.Slice())
		} else {
			n.Rlist.Set1(e)
		}
	}
	p.closeAnotherScope(pstate)
	return n
}

func (p *noder) forStmt(pstate *PackageState, stmt *syntax.ForStmt) *Node {
	p.openScope(pstate, stmt.Pos())
	var n *Node
	if r, ok := stmt.Init.(*syntax.RangeClause); ok {
		if stmt.Cond != nil || stmt.Post != nil {
			panic("unexpected RangeClause")
		}

		n = p.nod(pstate, r, ORANGE, nil, p.expr(pstate, r.X))
		if r.Lhs != nil {
			n.List.Set(p.assignList(pstate, r.Lhs, n, r.Def))
		}
	} else {
		n = p.nod(pstate, stmt, OFOR, nil, nil)
		if stmt.Init != nil {
			n.Ninit.Set1(p.stmt(pstate, stmt.Init))
		}
		if stmt.Cond != nil {
			n.Left = p.expr(pstate, stmt.Cond)
		}
		if stmt.Post != nil {
			n.Right = p.stmt(pstate, stmt.Post)
		}
	}
	n.Nbody.Set(p.blockStmt(pstate, stmt.Body))
	p.closeAnotherScope(pstate)
	return n
}

func (p *noder) switchStmt(pstate *PackageState, stmt *syntax.SwitchStmt) *Node {
	p.openScope(pstate, stmt.Pos())
	n := p.nod(pstate, stmt, OSWITCH, nil, nil)
	if stmt.Init != nil {
		n.Ninit.Set1(p.stmt(pstate, stmt.Init))
	}
	if stmt.Tag != nil {
		n.Left = p.expr(pstate, stmt.Tag)
	}

	tswitch := n.Left
	if tswitch != nil && tswitch.Op != OTYPESW {
		tswitch = nil
	}
	n.List.Set(p.caseClauses(pstate, stmt.Body, tswitch, stmt.Rbrace))

	p.closeScope(pstate, stmt.Rbrace)
	return n
}

func (p *noder) caseClauses(pstate *PackageState, clauses []*syntax.CaseClause, tswitch *Node, rbrace syntax.Pos) []*Node {
	var nodes []*Node
	for i, clause := range clauses {
		p.lineno(pstate, clause)
		if i > 0 {
			p.closeScope(pstate, clause.Pos())
		}
		p.openScope(pstate, clause.Pos())

		n := p.nod(pstate, clause, OXCASE, nil, nil)
		if clause.Cases != nil {
			n.List.Set(p.exprList(pstate, clause.Cases))
		}
		if tswitch != nil && tswitch.Left != nil {
			nn := pstate.newname(tswitch.Left.Sym)
			pstate.declare(nn, pstate.dclcontext)
			n.Rlist.Set1(nn)
			// keep track of the instances for reporting unused
			nn.Name.Defn = tswitch
		}

		// Trim trailing empty statements. We omit them from
		// the Node AST anyway, and it's easier to identify
		// out-of-place fallthrough statements without them.
		body := clause.Body
		for len(body) > 0 {
			if _, ok := body[len(body)-1].(*syntax.EmptyStmt); !ok {
				break
			}
			body = body[:len(body)-1]
		}

		n.Nbody.Set(p.stmtsFall(pstate, body, true))
		if l := n.Nbody.Len(); l > 0 && n.Nbody.Index(l-1).Op == OFALL {
			if tswitch != nil {
				pstate.yyerror("cannot fallthrough in type switch")
			}
			if i+1 == len(clauses) {
				pstate.yyerror("cannot fallthrough final case in switch")
			}
		}

		nodes = append(nodes, n)
	}
	if len(clauses) > 0 {
		p.closeScope(pstate, rbrace)
	}
	return nodes
}

func (p *noder) selectStmt(pstate *PackageState, stmt *syntax.SelectStmt) *Node {
	n := p.nod(pstate, stmt, OSELECT, nil, nil)
	n.List.Set(p.commClauses(pstate, stmt.Body, stmt.Rbrace))
	return n
}

func (p *noder) commClauses(pstate *PackageState, clauses []*syntax.CommClause, rbrace syntax.Pos) []*Node {
	var nodes []*Node
	for i, clause := range clauses {
		p.lineno(pstate, clause)
		if i > 0 {
			p.closeScope(pstate, clause.Pos())
		}
		p.openScope(pstate, clause.Pos())

		n := p.nod(pstate, clause, OXCASE, nil, nil)
		if clause.Comm != nil {
			n.List.Set1(p.stmt(pstate, clause.Comm))
		}
		n.Nbody.Set(p.stmts(pstate, clause.Body))
		nodes = append(nodes, n)
	}
	if len(clauses) > 0 {
		p.closeScope(pstate, rbrace)
	}
	return nodes
}

func (p *noder) labeledStmt(pstate *PackageState, label *syntax.LabeledStmt, fallOK bool) *Node {
	lhs := p.nod(pstate, label, OLABEL, p.newname(pstate, label.Label), nil)

	var ls *Node
	if label.Stmt != nil { // TODO(mdempsky): Should always be present.
		ls = p.stmtFall(pstate, label.Stmt, fallOK)
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
	return pstate.liststmt(l)
}

func (p *noder) unOp(pstate *PackageState, op syntax.Operator) Op {
	if uint64(op) >= uint64(len(pstate.unOps)) || pstate.unOps[op] == 0 {
		panic("invalid Operator")
	}
	return pstate.unOps[op]
}

func (p *noder) binOp(pstate *PackageState, op syntax.Operator) Op {
	if uint64(op) >= uint64(len(pstate.binOps)) || pstate.binOps[op] == 0 {
		panic("invalid Operator")
	}
	return pstate.binOps[op]
}

func (p *noder) basicLit(pstate *PackageState, lit *syntax.BasicLit) Val {
	// TODO: Don't try to convert if we had syntax errors (conversions may fail).
	//       Use dummy values so we can continue to compile. Eventually, use a
	//       form of "unknown" literals that are ignored during type-checking so
	//       we can continue type-checking w/o spurious follow-up errors.
	switch s := lit.Value; lit.Kind {
	case syntax.IntLit:
		x := new(Mpint)
		x.SetString(pstate, s)
		return Val{U: x}

	case syntax.FloatLit:
		x := newMpflt()
		x.SetString(pstate, s)
		return Val{U: x}

	case syntax.ImagLit:
		x := new(Mpcplx)
		x.Imag.SetString(pstate, strings.TrimSuffix(s, "i"))
		return Val{U: x}

	case syntax.RuneLit:
		var r rune
		if u, err := strconv.Unquote(s); err == nil && len(u) > 0 {
			// Package syntax already reported any errors.
			// Check for them again though because 0 is a
			// better fallback value for invalid rune
			// literals than 0xFFFD.
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
			// strip carriage returns from raw string
			s = strings.Replace(s, "\r", "", -1)
		}
		// Ignore errors because package syntax already reported them.
		u, _ := strconv.Unquote(s)
		return Val{U: u}

	default:
		panic("unhandled BasicLit kind")
	}
}

func (p *noder) name(pstate *PackageState, name *syntax.Name) *types.Sym {
	return pstate.lookup(name.Value)
}

func (p *noder) mkname(pstate *PackageState, name *syntax.Name) *Node {
	// TODO(mdempsky): Set line number?
	return pstate.mkname(p.name(pstate, name))
}

func (p *noder) newname(pstate *PackageState, name *syntax.Name) *Node {
	// TODO(mdempsky): Set line number?
	return pstate.newname(p.name(pstate, name))
}

func (p *noder) wrapname(pstate *PackageState, n syntax.Node, x *Node) *Node {
	// These nodes do not carry line numbers.
	// Introduce a wrapper node to give them the correct line.
	switch x.Op {
	case OTYPE, OLITERAL:
		if x.Sym == nil {
			break
		}
		fallthrough
	case ONAME, ONONAME, OPACK:
		x = p.nod(pstate, n, OPAREN, x, nil)
		x.SetImplicit(true)
	}
	return x
}

func (p *noder) nod(pstate *PackageState, orig syntax.Node, op Op, left, right *Node) *Node {
	return p.setlineno(pstate, orig, pstate.nod(op, left, right))
}

func (p *noder) nodSym(pstate *PackageState, orig syntax.Node, op Op, left *Node, sym *types.Sym) *Node {
	return p.setlineno(pstate, orig, pstate.nodSym(op, left, sym))
}

func (p *noder) setlineno(pstate *PackageState, src_ syntax.Node, dst *Node) *Node {
	pos := src_.Pos()
	if !pos.IsKnown() {
		// TODO(mdempsky): Shouldn't happen. Fix package syntax.
		return dst
	}
	dst.Pos = p.makeXPos(pstate, pos)
	return dst
}

func (p *noder) lineno(pstate *PackageState, n syntax.Node) {
	if n == nil {
		return
	}
	pos := n.Pos()
	if !pos.IsKnown() {
		// TODO(mdempsky): Shouldn't happen. Fix package syntax.
		return
	}
	pstate.lineno = p.makeXPos(pstate, pos)
}

// error is called concurrently if files are parsed concurrently.
func (p *noder) error(err error) {
	p.err <- err.(syntax.Error)
}

// pragma is called concurrently if files are parsed concurrently.
func (p *noder) pragma(pstate *PackageState, pos syntax.Pos, text string) syntax.Pragma {
	switch {
	case strings.HasPrefix(text, "line "):
		// line directives are handled by syntax package
		panic("unreachable")

	case strings.HasPrefix(text, "go:linkname "):
		f := strings.Fields(text)
		if len(f) != 3 {
			p.error(syntax.Error{Pos: pos, Msg: "usage: //go:linkname localname linkname"})
			break
		}
		p.linknames = append(p.linknames, linkname{pos, f[1], f[2]})

	case strings.HasPrefix(text, "go:cgo_import_dynamic "):
		// This is permitted for general use because Solaris
		// code relies on it in golang.org/x/sys/unix and others.
		fields := pragmaFields(text)
		if len(fields) >= 4 {
			lib := strings.Trim(fields[3], "\"")
			if lib != "" && !safeArg(lib) && !pstate.isCgoGeneratedFile(pos) {
				p.error(syntax.Error{Pos: pos, Msg: fmt.Sprintf("invalid library name %q in cgo_import_dynamic directive", lib)})
			}
			p.pragcgo(pos, text)
			return pstate.pragmaValue("go:cgo_import_dynamic")
		}
		fallthrough
	case strings.HasPrefix(text, "go:cgo_"):
		// For security, we disallow //go:cgo_* directives other
		// than cgo_import_dynamic outside cgo-generated files.
		// Exception: they are allowed in the standard library, for runtime and syscall.
		if !pstate.isCgoGeneratedFile(pos) && !pstate.compiling_std {
			p.error(syntax.Error{Pos: pos, Msg: fmt.Sprintf("//%s only allowed in cgo-generated code", text)})
		}
		p.pragcgo(pos, text)
		fallthrough // because of //go:cgo_unsafe_args
	default:
		verb := text
		if i := strings.Index(text, " "); i >= 0 {
			verb = verb[:i]
		}
		prag := pstate.pragmaValue(verb)
		const runtimePragmas = Systemstack | Nowritebarrier | Nowritebarrierrec | Yeswritebarrierrec
		if !pstate.compiling_runtime && prag&runtimePragmas != 0 {
			p.error(syntax.Error{Pos: pos, Msg: fmt.Sprintf("//%s only allowed in runtime", verb)})
		}
		if prag == 0 && !pstate.allowedStdPragmas[verb] && pstate.compiling_std {
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
func (pstate *PackageState) isCgoGeneratedFile(pos syntax.Pos) bool {
	return strings.HasPrefix(filepath.Base(filepath.Clean(pstate.fileh(pos.Base().Filename()))), "_cgo_")
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

func (pstate *PackageState) mkname(sym *types.Sym) *Node {
	n := pstate.oldname(sym)
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

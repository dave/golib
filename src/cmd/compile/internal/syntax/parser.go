package syntax

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

const debug = false
const trace = false

type parser struct {
	file *PosBase
	errh ErrorHandler
	mode Mode
	scanner

	base   *PosBase // current position base
	first  error    // first error encountered
	errcnt int      // number of errors encountered
	pragma Pragma   // pragma flags

	fnest  int    // function nesting level (for error handling)
	xnest  int    // expression nesting level (for complit ambiguity resolution)
	indent []byte // tracing support
}

func (p *parser) init(file *PosBase, r io.Reader, errh ErrorHandler, pragh PragmaHandler, mode Mode) {
	p.file = file
	p.errh = errh
	p.mode = mode
	p.scanner.init(
		r,

		func(line, col uint, msg string) {
			if msg[0] != '/' {
				p.errorAt(p.posAt(line, col), msg)
				return
			}

			text := commentText(msg)
			if strings.HasPrefix(text, "line ") {
				var pos Pos // position immediately following the comment
				if msg[1] == '/' {

					pos = MakePos(p.file, line+1, colbase)
				} else {

					pos = MakePos(p.file, line, col+uint(len(msg)))
				}
				p.updateBase(pos, line, col+2+5, text[5:])
				return
			}

			if pragh != nil && strings.HasPrefix(text, "go:") {
				p.pragma |= pragh(p.posAt(line, col+2), text)
			}
		},
		directives,
	)

	p.base = file
	p.first = nil
	p.errcnt = 0
	p.pragma = 0

	p.fnest = 0
	p.xnest = 0
	p.indent = nil
}

// updateBase sets the current position base to a new line base at pos.
// The base's filename, line, and column values are extracted from text
// which is positioned at (tline, tcol) (only needed for error messages).
func (p *parser) updateBase(pos Pos, tline, tcol uint, text string) {
	i, n, ok := trailingDigits(text)
	if i == 0 {
		return
	}

	if !ok {

		p.errorAt(p.posAt(tline, tcol+i), "invalid line number: "+text[i:])
		return
	}

	var line, col uint
	i2, n2, ok2 := trailingDigits(text[:i-1])
	if ok2 {

		i, i2 = i2, i
		line, col = n2, n
		if col == 0 || col > PosMax {
			p.errorAt(p.posAt(tline, tcol+i2), "invalid column number: "+text[i2:])
			return
		}
		text = text[:i2-1]
	} else {

		line = n
	}

	if line == 0 || line > PosMax {
		p.errorAt(p.posAt(tline, tcol+i), "invalid line number: "+text[i:])
		return
	}

	filename := text[:i-1]
	if filename == "" && ok2 {
		filename = p.base.Filename()
	}

	p.base = NewLineBase(pos, filename, line, col)
}

func commentText(s string) string {
	if s[:2] == "/*" {
		return s[2 : len(s)-2]
	}

	i := len(s)
	if s[i-1] == '\r' {
		i--
	}
	return s[2:i]
}

func trailingDigits(text string) (uint, uint, bool) {

	i := strings.LastIndex(text, ":")
	if i < 0 {
		return 0, 0, false
	}

	n, err := strconv.ParseUint(text[i+1:], 10, 0)
	return uint(i + 1), uint(n), err == nil
}

func (p *parser) got(psess *PackageSession, tok token) bool {
	if p.tok == tok {
		p.next(psess)
		return true
	}
	return false
}

func (p *parser) want(psess *PackageSession, tok token) {
	if !p.got(psess, tok) {
		p.syntaxError(psess, "expecting "+psess.tokstring(tok))
		p.advance(psess)
	}
}

// posAt returns the Pos value for (line, col) and the current position base.
func (p *parser) posAt(line, col uint) Pos {
	return MakePos(p.base, line, col)
}

// error reports an error at the given position.
func (p *parser) errorAt(pos Pos, msg string) {
	err := Error{pos, msg}
	if p.first == nil {
		p.first = err
	}
	p.errcnt++
	if p.errh == nil {
		panic(p.first)
	}
	p.errh(err)
}

// syntaxErrorAt reports a syntax error at the given position.
func (p *parser) syntaxErrorAt(psess *PackageSession, pos Pos, msg string) {
	if trace {
		p.print("syntax error: " + msg)
	}

	if p.tok == _EOF && p.first != nil {
		return
	}

	switch {
	case msg == "":

	case strings.HasPrefix(msg, "in "), strings.HasPrefix(msg, "at "), strings.HasPrefix(msg, "after "):
		msg = " " + msg
	case strings.HasPrefix(msg, "expecting "):
		msg = ", " + msg
	default:

		p.errorAt(pos, "syntax error: "+msg)
		return
	}

	// determine token string
	var tok string
	switch p.tok {
	case _Name, _Semi:
		tok = p.lit
	case _Literal:
		tok = "literal " + p.lit
	case _Operator:
		tok = p.op.String(psess)
	case _AssignOp:
		tok = p.op.String(psess) + "="
	case _IncOp:
		tok = p.op.String(psess)
		tok += tok
	default:
		tok = psess.tokstring(p.tok)
	}

	p.errorAt(pos, "syntax error: unexpected "+tok+msg)
}

// tokstring returns the English word for selected punctuation tokens
// for more readable error messages.
func (psess *PackageSession) tokstring(tok token) string {
	switch tok {
	case _Comma:
		return "comma"
	case _Semi:
		return "semicolon or newline"
	}
	return tok.String(psess)
}

// Convenience methods using the current token position.
func (p *parser) pos() Pos                                      { return p.posAt(p.line, p.col) }
func (p *parser) syntaxError(psess *PackageSession, msg string) { p.syntaxErrorAt(psess, p.pos(), msg) }

// The stopset contains keywords that start a statement.
// They are good synchronization points in case of syntax
// errors and (usually) shouldn't be skipped over.
const stopset uint64 = 1<<_Break |
	1<<_Const |
	1<<_Continue |
	1<<_Defer |
	1<<_Fallthrough |
	1<<_For |
	1<<_Go |
	1<<_Goto |
	1<<_If |
	1<<_Return |
	1<<_Select |
	1<<_Switch |
	1<<_Type |
	1<<_Var

// Advance consumes tokens until it finds a token of the stopset or followlist.
// The stopset is only considered if we are inside a function (p.fnest > 0).
// The followlist is the list of valid tokens that can follow a production;
// if it is empty, exactly one (non-EOF) token is consumed to ensure progress.
func (p *parser) advance(psess *PackageSession, followlist ...token) {
	if trace {
		p.print(fmt.Sprintf("advance %s", followlist))
	}

	// compute follow set
	// (not speed critical, advance is only called in error situations)
	var followset uint64 = 1 << _EOF // don't skip over EOF
	if len(followlist) > 0 {
		if p.fnest > 0 {
			followset |= stopset
		}
		for _, tok := range followlist {
			followset |= 1 << tok
		}
	}

	for !contains(followset, p.tok) {
		if trace {
			p.print("skip " + p.tok.String(psess))
		}
		p.next(psess)
		if len(followlist) == 0 {
			break
		}
	}

	if trace {
		p.print("next " + p.tok.String(psess))
	}
}

// usage: defer p.trace(msg)()
func (p *parser) trace(msg string) func() {
	p.print(msg + " (")
	const tab = ". "
	p.indent = append(p.indent, tab...)
	return func() {
		p.indent = p.indent[:len(p.indent)-len(tab)]
		if x := recover(); x != nil {
			panic(x)
		}
		p.print(")")
	}
}

func (p *parser) print(msg string) {
	fmt.Printf("%5d: %s%s\n", p.line, p.indent, msg)
}

// SourceFile = PackageClause ";" { ImportDecl ";" } { TopLevelDecl ";" } .
func (p *parser) fileOrNil(psess *PackageSession) *File {
	if trace {
		defer p.trace("file")()
	}

	f := new(File)
	f.pos = p.pos()

	if !p.got(psess, _Package) {
		p.syntaxError(psess, "package statement must be first")
		return nil
	}
	f.PkgName = p.name(psess)
	p.want(psess, _Semi)

	if p.first != nil {
		return nil
	}

	for p.got(psess, _Import) {
		f.DeclList = p.appendGroup(psess, f.DeclList, p.importDecl)
		p.want(psess, _Semi)
	}

	for p.tok != _EOF {
		switch p.tok {
		case _Const:
			p.next(psess)
			f.DeclList = p.appendGroup(psess, f.DeclList, p.constDecl)

		case _Type:
			p.next(psess)
			f.DeclList = p.appendGroup(psess, f.DeclList, p.typeDecl)

		case _Var:
			p.next(psess)
			f.DeclList = p.appendGroup(psess, f.DeclList, p.varDecl)

		case _Func:
			p.next(psess)
			if d := p.funcDeclOrNil(psess); d != nil {
				f.DeclList = append(f.DeclList, d)
			}

		default:
			if p.tok == _Lbrace && len(f.DeclList) > 0 && isEmptyFuncDecl(f.DeclList[len(f.DeclList)-1]) {

				p.syntaxError(psess, "unexpected semicolon or newline before {")
			} else {
				p.syntaxError(psess, "non-declaration statement outside function body")
			}
			p.advance(psess, _Const, _Type, _Var, _Func)
			continue
		}

		p.pragma = 0

		if p.tok != _EOF && !p.got(psess, _Semi) {
			p.syntaxError(psess, "after top level declaration")
			p.advance(psess, _Const, _Type, _Var, _Func)
		}
	}

	f.Lines = p.source.line

	return f
}

func isEmptyFuncDecl(dcl Decl) bool {
	f, ok := dcl.(*FuncDecl)
	return ok && f.Body == nil
}

// list parses a possibly empty, sep-separated list, optionally
// followed by sep and enclosed by ( and ) or { and }. open is
// one of _Lparen, or _Lbrace, sep is one of _Comma or _Semi,
// and close is expected to be the (closing) opposite of open.
// For each list element, f is called. After f returns true, no
// more list elements are accepted. list returns the position
// of the closing token.
//
// list = "(" { f sep } ")" |
//        "{" { f sep } "}" . // sep is optional before ")" or "}"
//
func (p *parser) list(psess *PackageSession, open, sep, close token, f func() bool) Pos {
	p.want(psess, open)

	var done bool
	for p.tok != _EOF && p.tok != close && !done {
		done = f()

		if !p.got(psess, sep) && p.tok != close {
			p.syntaxError(psess, fmt.Sprintf("expecting %s or %s", psess.tokstring(sep), psess.tokstring(close)))
			p.advance(psess, _Rparen, _Rbrack, _Rbrace)
			if p.tok != close {

				return p.pos()
			}
		}
	}

	pos := p.pos()
	p.want(psess, close)
	return pos
}

// appendGroup(f) = f | "(" { f ";" } ")" . // ";" is optional before ")"
func (p *parser) appendGroup(psess *PackageSession, list []Decl, f func(*Group) Decl) []Decl {
	if p.tok == _Lparen {
		g := new(Group)
		p.list(psess, _Lparen, _Semi, _Rparen, func() bool {
			list = append(list, f(g))
			return false
		})
	} else {
		list = append(list, f(nil))
	}

	if debug {
		for _, d := range list {
			if d == nil {
				panic("nil list entry")
			}
		}
	}

	return list
}

// ImportSpec = [ "." | PackageName ] ImportPath .
// ImportPath = string_lit .
func (p *parser) importDecl(psess *PackageSession, group *Group) Decl {
	if trace {
		defer p.trace("importDecl")()
	}

	d := new(ImportDecl)
	d.pos = p.pos()

	switch p.tok {
	case _Name:
		d.LocalPkgName = p.name(psess)
	case _Dot:
		d.LocalPkgName = p.newName(".")
		p.next(psess)
	}
	d.Path = p.oliteral(psess)
	if d.Path == nil {
		p.syntaxError(psess, "missing import path")
		p.advance(psess, _Semi, _Rparen)
		return nil
	}
	d.Group = group

	return d
}

// ConstSpec = IdentifierList [ [ Type ] "=" ExpressionList ] .
func (p *parser) constDecl(psess *PackageSession, group *Group) Decl {
	if trace {
		defer p.trace("constDecl")()
	}

	d := new(ConstDecl)
	d.pos = p.pos()

	d.NameList = p.nameList(psess, p.name(psess))
	if p.tok != _EOF && p.tok != _Semi && p.tok != _Rparen {
		d.Type = p.typeOrNil(psess)
		if p.got(psess, _Assign) {
			d.Values = p.exprList(psess)
		}
	}
	d.Group = group

	return d
}

// TypeSpec = identifier [ "=" ] Type .
func (p *parser) typeDecl(psess *PackageSession, group *Group) Decl {
	if trace {
		defer p.trace("typeDecl")()
	}

	d := new(TypeDecl)
	d.pos = p.pos()

	d.Name = p.name(psess)
	d.Alias = p.got(psess, _Assign)
	d.Type = p.typeOrNil(psess)
	if d.Type == nil {
		d.Type = p.bad()
		p.syntaxError(psess, "in type declaration")
		p.advance(psess, _Semi, _Rparen)
	}
	d.Group = group
	d.Pragma = p.pragma

	return d
}

// VarSpec = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .
func (p *parser) varDecl(psess *PackageSession, group *Group) Decl {
	if trace {
		defer p.trace("varDecl")()
	}

	d := new(VarDecl)
	d.pos = p.pos()

	d.NameList = p.nameList(psess, p.name(psess))
	if p.got(psess, _Assign) {
		d.Values = p.exprList(psess)
	} else {
		d.Type = p.type_(psess)
		if p.got(psess, _Assign) {
			d.Values = p.exprList(psess)
		}
	}
	d.Group = group

	return d
}

// FunctionDecl = "func" FunctionName ( Function | Signature ) .
// FunctionName = identifier .
// Function     = Signature FunctionBody .
// MethodDecl   = "func" Receiver MethodName ( Function | Signature ) .
// Receiver     = Parameters .
func (p *parser) funcDeclOrNil(psess *PackageSession) *FuncDecl {
	if trace {
		defer p.trace("funcDecl")()
	}

	f := new(FuncDecl)
	f.pos = p.pos()

	if p.tok == _Lparen {
		rcvr := p.paramList(psess)
		switch len(rcvr) {
		case 0:
			p.error("method has no receiver")
		default:
			p.error("method has multiple receivers")
			fallthrough
		case 1:
			f.Recv = rcvr[0]
		}
	}

	if p.tok != _Name {
		p.syntaxError(psess, "expecting name or (")
		p.advance(psess, _Lbrace, _Semi)
		return nil
	}

	f.Name = p.name(psess)
	f.Type = p.funcType(psess)
	if p.tok == _Lbrace {
		f.Body = p.funcBody(psess)
	}
	f.Pragma = p.pragma

	return f
}

func (p *parser) funcBody(psess *PackageSession) *BlockStmt {
	p.fnest++
	errcnt := p.errcnt
	body := p.blockStmt(psess, "")
	p.fnest--

	if p.mode&CheckBranches != 0 && errcnt == p.errcnt {
		psess.
			checkBranches(body, p.errh)
	}

	return body
}

func (p *parser) expr(psess *PackageSession) Expr {
	if trace {
		defer p.trace("expr")()
	}

	return p.binaryExpr(psess, 0)
}

// Expression = UnaryExpr | Expression binary_op Expression .
func (p *parser) binaryExpr(psess *PackageSession, prec int) Expr {

	x := p.unaryExpr(psess)
	for (p.tok == _Operator || p.tok == _Star) && p.prec > prec {
		t := new(Operation)
		t.pos = p.pos()
		t.Op = p.op
		t.X = x
		tprec := p.prec
		p.next(psess)
		t.Y = p.binaryExpr(psess, tprec)
		x = t
	}
	return x
}

// UnaryExpr = PrimaryExpr | unary_op UnaryExpr .
func (p *parser) unaryExpr(psess *PackageSession) Expr {
	if trace {
		defer p.trace("unaryExpr")()
	}

	switch p.tok {
	case _Operator, _Star:
		switch p.op {
		case Mul, Add, Sub, Not, Xor:
			x := new(Operation)
			x.pos = p.pos()
			x.Op = p.op
			p.next(psess)
			x.X = p.unaryExpr(psess)
			return x

		case And:
			x := new(Operation)
			x.pos = p.pos()
			x.Op = And
			p.next(psess)

			x.X = unparen(p.unaryExpr(psess))
			return x
		}

	case _Arrow:

		pos := p.pos()
		p.next(psess)

		x := p.unaryExpr(psess)

		if _, ok := x.(*ChanType); ok {

			dir := SendOnly
			t := x
			for dir == SendOnly {
				c, ok := t.(*ChanType)
				if !ok {
					break
				}
				dir = c.Dir
				if dir == RecvOnly {

					p.syntaxError(psess, "unexpected <-, expecting chan")

				}
				c.Dir = RecvOnly
				t = c.Elem
			}
			if dir == SendOnly {

				p.syntaxError(psess, fmt.Sprintf("unexpected %s, expecting chan", psess.String(t)))

			}
			return x
		}

		o := new(Operation)
		o.pos = pos
		o.Op = Recv
		o.X = x
		return o
	}

	return p.pexpr(psess, true)
}

// callStmt parses call-like statements that can be preceded by 'defer' and 'go'.
func (p *parser) callStmt(psess *PackageSession) *CallStmt {
	if trace {
		defer p.trace("callStmt")()
	}

	s := new(CallStmt)
	s.pos = p.pos()
	s.Tok = p.tok
	p.next(psess)

	x := p.pexpr(psess, p.tok == _Lparen)
	if t := unparen(x); t != x {
		p.errorAt(x.Pos(), fmt.Sprintf("expression in %s must not be parenthesized", s.Tok))

		x = t
	}

	cx, ok := x.(*CallExpr)
	if !ok {
		p.errorAt(x.Pos(), fmt.Sprintf("expression in %s must be function call", s.Tok))

		cx = new(CallExpr)
		cx.pos = x.Pos()
		cx.Fun = x
	}

	s.Call = cx
	return s
}

// Operand     = Literal | OperandName | MethodExpr | "(" Expression ")" .
// Literal     = BasicLit | CompositeLit | FunctionLit .
// BasicLit    = int_lit | float_lit | imaginary_lit | rune_lit | string_lit .
// OperandName = identifier | QualifiedIdent.
func (p *parser) operand(psess *PackageSession, keep_parens bool) Expr {
	if trace {
		defer p.trace("operand " + p.tok.String(psess))()
	}

	switch p.tok {
	case _Name:
		return p.name(psess)

	case _Literal:
		return p.oliteral(psess)

	case _Lparen:
		pos := p.pos()
		p.next(psess)
		p.xnest++
		x := p.expr(psess)
		p.xnest--
		p.want(psess, _Rparen)

		if p.tok == _Lbrace {
			keep_parens = true
		}

		if keep_parens {
			px := new(ParenExpr)
			px.pos = pos
			px.X = x
			x = px
		}
		return x

	case _Func:
		pos := p.pos()
		p.next(psess)
		t := p.funcType(psess)
		if p.tok == _Lbrace {
			p.xnest++

			f := new(FuncLit)
			f.pos = pos
			f.Type = t
			f.Body = p.funcBody(psess)

			p.xnest--
			return f
		}
		return t

	case _Lbrack, _Chan, _Map, _Struct, _Interface:
		return p.type_(psess)

	default:
		x := p.bad()
		p.syntaxError(psess, "expecting expression")
		p.advance(psess)
		return x
	}

}

// PrimaryExpr =
// 	Operand |
// 	Conversion |
// 	PrimaryExpr Selector |
// 	PrimaryExpr Index |
// 	PrimaryExpr Slice |
// 	PrimaryExpr TypeAssertion |
// 	PrimaryExpr Arguments .
//
// Selector       = "." identifier .
// Index          = "[" Expression "]" .
// Slice          = "[" ( [ Expression ] ":" [ Expression ] ) |
//                      ( [ Expression ] ":" Expression ":" Expression )
//                  "]" .
// TypeAssertion  = "." "(" Type ")" .
// Arguments      = "(" [ ( ExpressionList | Type [ "," ExpressionList ] ) [ "..." ] [ "," ] ] ")" .
func (p *parser) pexpr(psess *PackageSession, keep_parens bool) Expr {
	if trace {
		defer p.trace("pexpr")()
	}

	x := p.operand(psess, keep_parens)

loop:
	for {
		pos := p.pos()
		switch p.tok {
		case _Dot:
			p.next(psess)
			switch p.tok {
			case _Name:

				t := new(SelectorExpr)
				t.pos = pos
				t.X = x
				t.Sel = p.name(psess)
				x = t

			case _Lparen:
				p.next(psess)
				if p.got(psess, _Type) {
					t := new(TypeSwitchGuard)

					t.pos = pos
					t.X = x
					x = t
				} else {
					t := new(AssertExpr)
					t.pos = pos
					t.X = x
					t.Type = p.type_(psess)
					x = t
				}
				p.want(psess, _Rparen)

			default:
				p.syntaxError(psess, "expecting name or (")
				p.advance(psess, _Semi, _Rparen)
			}

		case _Lbrack:
			p.next(psess)
			p.xnest++

			var i Expr
			if p.tok != _Colon {
				i = p.expr(psess)
				if p.got(psess, _Rbrack) {

					t := new(IndexExpr)
					t.pos = pos
					t.X = x
					t.Index = i
					x = t
					p.xnest--
					break
				}
			}

			t := new(SliceExpr)
			t.pos = pos
			t.X = x
			t.Index[0] = i
			p.want(psess, _Colon)
			if p.tok != _Colon && p.tok != _Rbrack {

				t.Index[1] = p.expr(psess)
			}
			if p.got(psess, _Colon) {
				t.Full = true

				if t.Index[1] == nil {
					p.error("middle index required in 3-index slice")
				}
				if p.tok != _Rbrack {

					t.Index[2] = p.expr(psess)
				} else {
					p.error("final index required in 3-index slice")
				}
			}
			p.want(psess, _Rbrack)

			x = t
			p.xnest--

		case _Lparen:
			t := new(CallExpr)
			t.pos = pos
			t.Fun = x
			t.ArgList, t.HasDots = p.argList(psess)
			x = t

		case _Lbrace:

			t := unparen(x)

			complit_ok := false
			switch t.(type) {
			case *Name, *SelectorExpr:
				if p.xnest >= 0 {

					complit_ok = true
				}
			case *ArrayType, *SliceType, *StructType, *MapType:

				complit_ok = true
			}
			if !complit_ok {
				break loop
			}
			if t != x {
				p.syntaxError(psess, "cannot parenthesize type in composite literal")

			}
			n := p.complitexpr(psess)
			n.Type = x
			x = n

		default:
			break loop
		}
	}

	return x
}

// Element = Expression | LiteralValue .
func (p *parser) bare_complitexpr(psess *PackageSession) Expr {
	if trace {
		defer p.trace("bare_complitexpr")()
	}

	if p.tok == _Lbrace {

		return p.complitexpr(psess)
	}

	return p.expr(psess)
}

// LiteralValue = "{" [ ElementList [ "," ] ] "}" .
func (p *parser) complitexpr(psess *PackageSession) *CompositeLit {
	if trace {
		defer p.trace("complitexpr")()
	}

	x := new(CompositeLit)
	x.pos = p.pos()

	p.xnest++
	x.Rbrace = p.list(psess, _Lbrace, _Comma, _Rbrace, func() bool {

		e := p.bare_complitexpr(psess)
		if p.tok == _Colon {

			l := new(KeyValueExpr)
			l.pos = p.pos()
			p.next(psess)
			l.Key = e
			l.Value = p.bare_complitexpr(psess)
			e = l
			x.NKeys++
		}
		x.ElemList = append(x.ElemList, e)
		return false
	})
	p.xnest--

	return x
}

func (p *parser) type_(psess *PackageSession) Expr {
	if trace {
		defer p.trace("type_")()
	}

	typ := p.typeOrNil(psess)
	if typ == nil {
		typ = p.bad()
		p.syntaxError(psess, "expecting type")
		p.advance(psess, _Comma, _Colon, _Semi, _Rparen, _Rbrack, _Rbrace)
	}

	return typ
}

func newIndirect(pos Pos, typ Expr) Expr {
	o := new(Operation)
	o.pos = pos
	o.Op = Mul
	o.X = typ
	return o
}

// typeOrNil is like type_ but it returns nil if there was no type
// instead of reporting an error.
//
// Type     = TypeName | TypeLit | "(" Type ")" .
// TypeName = identifier | QualifiedIdent .
// TypeLit  = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
// 	      SliceType | MapType | Channel_Type .
func (p *parser) typeOrNil(psess *PackageSession) Expr {
	if trace {
		defer p.trace("typeOrNil")()
	}

	pos := p.pos()
	switch p.tok {
	case _Star:

		p.next(psess)
		return newIndirect(pos, p.type_(psess))

	case _Arrow:

		p.next(psess)
		p.want(psess, _Chan)
		t := new(ChanType)
		t.pos = pos
		t.Dir = RecvOnly
		t.Elem = p.chanElem(psess)
		return t

	case _Func:

		p.next(psess)
		return p.funcType(psess)

	case _Lbrack:

		p.next(psess)
		p.xnest++
		if p.got(psess, _Rbrack) {

			p.xnest--
			t := new(SliceType)
			t.pos = pos
			t.Elem = p.type_(psess)
			return t
		}

		t := new(ArrayType)
		t.pos = pos
		if !p.got(psess, _DotDotDot) {
			t.Len = p.expr(psess)
		}
		p.want(psess, _Rbrack)
		p.xnest--
		t.Elem = p.type_(psess)
		return t

	case _Chan:

		p.next(psess)
		t := new(ChanType)
		t.pos = pos
		if p.got(psess, _Arrow) {
			t.Dir = SendOnly
		}
		t.Elem = p.chanElem(psess)
		return t

	case _Map:

		p.next(psess)
		p.want(psess, _Lbrack)
		t := new(MapType)
		t.pos = pos
		t.Key = p.type_(psess)
		p.want(psess, _Rbrack)
		t.Value = p.type_(psess)
		return t

	case _Struct:
		return p.structType(psess)

	case _Interface:
		return p.interfaceType(psess)

	case _Name:
		return p.dotname(psess, p.name(psess))

	case _Lparen:
		p.next(psess)
		t := p.type_(psess)
		p.want(psess, _Rparen)
		return t
	}

	return nil
}

func (p *parser) funcType(psess *PackageSession) *FuncType {
	if trace {
		defer p.trace("funcType")()
	}

	typ := new(FuncType)
	typ.pos = p.pos()
	typ.ParamList = p.paramList(psess)
	typ.ResultList = p.funcResult(psess)

	return typ
}

func (p *parser) chanElem(psess *PackageSession) Expr {
	if trace {
		defer p.trace("chanElem")()
	}

	typ := p.typeOrNil(psess)
	if typ == nil {
		typ = p.bad()
		p.syntaxError(psess, "missing channel element type")

	}

	return typ
}

func (p *parser) dotname(psess *PackageSession, name *Name) Expr {
	if trace {
		defer p.trace("dotname")()
	}

	if p.tok == _Dot {
		s := new(SelectorExpr)
		s.pos = p.pos()
		p.next(psess)
		s.X = name
		s.Sel = p.name(psess)
		return s
	}
	return name
}

// StructType = "struct" "{" { FieldDecl ";" } "}" .
func (p *parser) structType(psess *PackageSession) *StructType {
	if trace {
		defer p.trace("structType")()
	}

	typ := new(StructType)
	typ.pos = p.pos()

	p.want(psess, _Struct)
	p.list(psess, _Lbrace, _Semi, _Rbrace, func() bool {
		p.fieldDecl(psess, typ)
		return false
	})

	return typ
}

// InterfaceType = "interface" "{" { MethodSpec ";" } "}" .
func (p *parser) interfaceType(psess *PackageSession) *InterfaceType {
	if trace {
		defer p.trace("interfaceType")()
	}

	typ := new(InterfaceType)
	typ.pos = p.pos()

	p.want(psess, _Interface)
	p.list(psess, _Lbrace, _Semi, _Rbrace, func() bool {
		if m := p.methodDecl(psess); m != nil {
			typ.MethodList = append(typ.MethodList, m)
		}
		return false
	})

	return typ
}

// Result = Parameters | Type .
func (p *parser) funcResult(psess *PackageSession) []*Field {
	if trace {
		defer p.trace("funcResult")()
	}

	if p.tok == _Lparen {
		return p.paramList(psess)
	}

	pos := p.pos()
	if typ := p.typeOrNil(psess); typ != nil {
		f := new(Field)
		f.pos = pos
		f.Type = typ
		return []*Field{f}
	}

	return nil
}

func (p *parser) addField(styp *StructType, pos Pos, name *Name, typ Expr, tag *BasicLit) {
	if tag != nil {
		for i := len(styp.FieldList) - len(styp.TagList); i > 0; i-- {
			styp.TagList = append(styp.TagList, nil)
		}
		styp.TagList = append(styp.TagList, tag)
	}

	f := new(Field)
	f.pos = pos
	f.Name = name
	f.Type = typ
	styp.FieldList = append(styp.FieldList, f)

	if debug && tag != nil && len(styp.FieldList) != len(styp.TagList) {
		panic("inconsistent struct field list")
	}
}

// FieldDecl      = (IdentifierList Type | AnonymousField) [ Tag ] .
// AnonymousField = [ "*" ] TypeName .
// Tag            = string_lit .
func (p *parser) fieldDecl(psess *PackageSession, styp *StructType) {
	if trace {
		defer p.trace("fieldDecl")()
	}

	pos := p.pos()
	switch p.tok {
	case _Name:
		name := p.name(psess)
		if p.tok == _Dot || p.tok == _Literal || p.tok == _Semi || p.tok == _Rbrace {

			typ := p.qualifiedName(psess, name)
			tag := p.oliteral(psess)
			p.addField(styp, pos, nil, typ, tag)
			return
		}

		names := p.nameList(psess, name)
		typ := p.type_(psess)
		tag := p.oliteral(psess)

		for _, name := range names {
			p.addField(styp, name.Pos(), name, typ, tag)
		}

	case _Lparen:
		p.next(psess)
		if p.tok == _Star {

			pos := p.pos()
			p.next(psess)
			typ := newIndirect(pos, p.qualifiedName(psess, nil))
			p.want(psess, _Rparen)
			tag := p.oliteral(psess)
			p.addField(styp, pos, nil, typ, tag)
			p.syntaxError(psess, "cannot parenthesize embedded type")

		} else {

			typ := p.qualifiedName(psess, nil)
			p.want(psess, _Rparen)
			tag := p.oliteral(psess)
			p.addField(styp, pos, nil, typ, tag)
			p.syntaxError(psess, "cannot parenthesize embedded type")
		}

	case _Star:
		p.next(psess)
		if p.got(psess, _Lparen) {

			typ := newIndirect(pos, p.qualifiedName(psess, nil))
			p.want(psess, _Rparen)
			tag := p.oliteral(psess)
			p.addField(styp, pos, nil, typ, tag)
			p.syntaxError(psess, "cannot parenthesize embedded type")

		} else {

			typ := newIndirect(pos, p.qualifiedName(psess, nil))
			tag := p.oliteral(psess)
			p.addField(styp, pos, nil, typ, tag)
		}

	default:
		p.syntaxError(psess, "expecting field name or embedded type")
		p.advance(psess, _Semi, _Rbrace)
	}
}

func (p *parser) oliteral(psess *PackageSession) *BasicLit {
	if p.tok == _Literal {
		b := new(BasicLit)
		b.pos = p.pos()
		b.Value = p.lit
		b.Kind = p.kind
		p.next(psess)
		return b
	}
	return nil
}

// MethodSpec        = MethodName Signature | InterfaceTypeName .
// MethodName        = identifier .
// InterfaceTypeName = TypeName .
func (p *parser) methodDecl(psess *PackageSession) *Field {
	if trace {
		defer p.trace("methodDecl")()
	}

	switch p.tok {
	case _Name:
		name := p.name(psess)

		hasNameList := false
		for p.got(psess, _Comma) {
			p.name(psess)
			hasNameList = true
		}
		if hasNameList {
			p.syntaxError(psess, "name list not allowed in interface type")

		}

		f := new(Field)
		f.pos = name.Pos()
		if p.tok != _Lparen {

			f.Type = p.qualifiedName(psess, name)
			return f
		}

		f.Name = name
		f.Type = p.funcType(psess)
		return f

	case _Lparen:
		p.syntaxError(psess, "cannot parenthesize embedded type")
		f := new(Field)
		f.pos = p.pos()
		p.next(psess)
		f.Type = p.qualifiedName(psess, nil)
		p.want(psess, _Rparen)
		return f

	default:
		p.syntaxError(psess, "expecting method or interface name")
		p.advance(psess, _Semi, _Rbrace)
		return nil
	}
}

// ParameterDecl = [ IdentifierList ] [ "..." ] Type .
func (p *parser) paramDeclOrNil(psess *PackageSession) *Field {
	if trace {
		defer p.trace("paramDecl")()
	}

	f := new(Field)
	f.pos = p.pos()

	switch p.tok {
	case _Name:
		f.Name = p.name(psess)
		switch p.tok {
		case _Name, _Star, _Arrow, _Func, _Lbrack, _Chan, _Map, _Struct, _Interface, _Lparen:

			f.Type = p.type_(psess)

		case _DotDotDot:

			f.Type = p.dotsType(psess)

		case _Dot:

			f.Type = p.dotname(psess, f.Name)
			f.Name = nil
		}

	case _Arrow, _Star, _Func, _Lbrack, _Chan, _Map, _Struct, _Interface, _Lparen:

		f.Type = p.type_(psess)

	case _DotDotDot:

		f.Type = p.dotsType(psess)

	default:
		p.syntaxError(psess, "expecting )")
		p.advance(psess, _Comma, _Rparen)
		return nil
	}

	return f
}

// ...Type
func (p *parser) dotsType(psess *PackageSession) *DotsType {
	if trace {
		defer p.trace("dotsType")()
	}

	t := new(DotsType)
	t.pos = p.pos()

	p.want(psess, _DotDotDot)
	t.Elem = p.typeOrNil(psess)
	if t.Elem == nil {
		t.Elem = p.bad()
		p.syntaxError(psess, "final argument in variadic function missing type")
	}

	return t
}

// Parameters    = "(" [ ParameterList [ "," ] ] ")" .
// ParameterList = ParameterDecl { "," ParameterDecl } .
func (p *parser) paramList(psess *PackageSession) (list []*Field) {
	if trace {
		defer p.trace("paramList")()
	}

	pos := p.pos()

	var named int // number of parameters that have an explicit name and type
	p.list(psess, _Lparen, _Comma, _Rparen, func() bool {
		if par := p.paramDeclOrNil(psess); par != nil {
			if debug && par.Name == nil && par.Type == nil {
				panic("parameter without name or type")
			}
			if par.Name != nil && par.Type != nil {
				named++
			}
			list = append(list, par)
		}
		return false
	})

	if named == 0 {

		for _, par := range list {
			if typ := par.Name; typ != nil {
				par.Type = typ
				par.Name = nil
			}
		}
	} else if named != len(list) {

		ok := true
		var typ Expr
		for i := len(list) - 1; i >= 0; i-- {
			if par := list[i]; par.Type != nil {
				typ = par.Type
				if par.Name == nil {
					ok = false
					n := p.newName("_")
					n.pos = typ.Pos()
					par.Name = n
				}
			} else if typ != nil {
				par.Type = typ
			} else {

				ok = false
				t := p.bad()
				t.pos = par.Name.Pos()
				par.Type = t
			}
		}
		if !ok {
			p.syntaxErrorAt(psess, pos, "mixed named and unnamed function parameters")
		}
	}

	return
}

func (p *parser) bad() *BadExpr {
	b := new(BadExpr)
	b.pos = p.pos()
	return b
}

// We represent x++, x-- as assignments x += ImplicitOne, x -= ImplicitOne.
// ImplicitOne should not be used elsewhere.

// SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
func (p *parser) simpleStmt(psess *PackageSession, lhs Expr, keyword token) SimpleStmt {
	if trace {
		defer p.trace("simpleStmt")()
	}

	if keyword == _For && p.tok == _Range {

		if debug && lhs != nil {
			panic("invalid call of simpleStmt")
		}
		return p.newRangeClause(psess, nil, false)
	}

	if lhs == nil {
		lhs = p.exprList(psess)
	}

	if _, ok := lhs.(*ListExpr); !ok && p.tok != _Assign && p.tok != _Define {

		pos := p.pos()
		switch p.tok {
		case _AssignOp:

			op := p.op
			p.next(psess)
			return p.newAssignStmt(pos, op, lhs, p.expr(psess))

		case _IncOp:

			op := p.op
			p.next(psess)
			return p.newAssignStmt(pos, op, lhs, psess.ImplicitOne)

		case _Arrow:

			s := new(SendStmt)
			s.pos = pos
			p.next(psess)
			s.Chan = lhs
			s.Value = p.expr(psess)
			return s

		default:

			s := new(ExprStmt)
			s.pos = lhs.Pos()
			s.X = lhs
			return s
		}
	}

	switch p.tok {
	case _Assign, _Define:
		pos := p.pos()
		var op Operator
		if p.tok == _Define {
			op = Def
		}
		p.next(psess)

		if keyword == _For && p.tok == _Range {

			return p.newRangeClause(psess, lhs, op == Def)
		}

		rhs := p.exprList(psess)

		if x, ok := rhs.(*TypeSwitchGuard); ok && keyword == _Switch && op == Def {
			if lhs, ok := lhs.(*Name); ok {

				x.Lhs = lhs
				s := new(ExprStmt)
				s.pos = x.Pos()
				s.X = x
				return s
			}
		}

		return p.newAssignStmt(pos, op, lhs, rhs)

	default:
		p.syntaxError(psess, "expecting := or = or comma")
		p.advance(psess, _Semi, _Rbrace)

		if x, ok := lhs.(*ListExpr); ok {
			lhs = x.ElemList[0]
		}
		s := new(ExprStmt)
		s.pos = lhs.Pos()
		s.X = lhs
		return s
	}
}

func (p *parser) newRangeClause(psess *PackageSession, lhs Expr, def bool) *RangeClause {
	r := new(RangeClause)
	r.pos = p.pos()
	p.next(psess)
	r.Lhs = lhs
	r.Def = def
	r.X = p.expr(psess)
	return r
}

func (p *parser) newAssignStmt(pos Pos, op Operator, lhs, rhs Expr) *AssignStmt {
	a := new(AssignStmt)
	a.pos = pos
	a.Op = op
	a.Lhs = lhs
	a.Rhs = rhs
	return a
}

func (p *parser) labeledStmtOrNil(psess *PackageSession, label *Name) Stmt {
	if trace {
		defer p.trace("labeledStmt")()
	}

	s := new(LabeledStmt)
	s.pos = p.pos()
	s.Label = label

	p.want(psess, _Colon)

	if p.tok == _Rbrace {

		e := new(EmptyStmt)
		e.pos = p.pos()
		s.Stmt = e
		return s
	}

	s.Stmt = p.stmtOrNil(psess)
	if s.Stmt != nil {
		return s
	}

	p.syntaxErrorAt(psess, s.pos, "missing statement after label")

	return nil
}

// context must be a non-empty string unless we know that p.tok == _Lbrace.
func (p *parser) blockStmt(psess *PackageSession, context string) *BlockStmt {
	if trace {
		defer p.trace("blockStmt")()
	}

	s := new(BlockStmt)
	s.pos = p.pos()

	if !p.got(psess, _Lbrace) {
		p.syntaxError(psess, "expecting { after "+context)
		p.advance(psess, _Name, _Rbrace)
		s.Rbrace = p.pos()
		if p.got(psess, _Rbrace) {
			return s
		}
	}

	s.List = p.stmtList(psess)
	s.Rbrace = p.pos()
	p.want(psess, _Rbrace)

	return s
}

func (p *parser) declStmt(psess *PackageSession, f func(*Group) Decl) *DeclStmt {
	if trace {
		defer p.trace("declStmt")()
	}

	s := new(DeclStmt)
	s.pos = p.pos()

	p.next(psess)
	s.DeclList = p.appendGroup(psess, nil, f)

	return s
}

func (p *parser) forStmt(psess *PackageSession) Stmt {
	if trace {
		defer p.trace("forStmt")()
	}

	s := new(ForStmt)
	s.pos = p.pos()

	s.Init, s.Cond, s.Post = p.header(psess, _For)
	s.Body = p.blockStmt(psess, "for clause")

	return s
}

func (p *parser) header(psess *PackageSession, keyword token) (init SimpleStmt, cond Expr, post SimpleStmt) {
	p.want(psess, keyword)

	if p.tok == _Lbrace {
		if keyword == _If {
			p.syntaxError(psess, "missing condition in if statement")
		}
		return
	}

	outer := p.xnest
	p.xnest = -1

	if p.tok != _Semi {

		if p.got(psess, _Var) {
			p.syntaxError(psess, fmt.Sprintf("var declaration not allowed in %s initializer", keyword.String(psess)))
		}
		init = p.simpleStmt(psess, nil, keyword)

		if _, ok := init.(*RangeClause); ok {
			p.xnest = outer
			return
		}
	}

	var condStmt SimpleStmt
	var semi struct {
		pos Pos
		lit string // valid if pos.IsKnown()
	}
	if p.tok != _Lbrace {
		if p.tok == _Semi {
			semi.pos = p.pos()
			semi.lit = p.lit
			p.next(psess)
		} else {

			p.want(psess, _Lbrace)
		}
		if keyword == _For {
			if p.tok != _Semi {
				if p.tok == _Lbrace {
					p.syntaxError(psess, "expecting for loop condition")
					goto done
				}
				condStmt = p.simpleStmt(psess, nil, 0)
			}
			p.want(psess, _Semi)
			if p.tok != _Lbrace {
				post = p.simpleStmt(psess, nil, 0)
				if a, _ := post.(*AssignStmt); a != nil && a.Op == Def {
					p.syntaxErrorAt(psess, a.Pos(), "cannot declare in post statement of for loop")
				}
			}
		} else if p.tok != _Lbrace {
			condStmt = p.simpleStmt(psess, nil, keyword)
		}
	} else {
		condStmt = init
		init = nil
	}

done:

	switch s := condStmt.(type) {
	case nil:
		if keyword == _If && semi.pos.IsKnown() {
			if semi.lit != "semicolon" {
				p.syntaxErrorAt(psess, semi.pos, fmt.Sprintf("unexpected %s, expecting { after if clause", semi.lit))
			} else {
				p.syntaxErrorAt(psess, semi.pos, "missing condition in if statement")
			}
		}
	case *ExprStmt:
		cond = s.X
	default:

		str := psess.String(s)
		if as, ok := s.(*AssignStmt); ok && as.Op == 0 {
			str = "assignment " + str
		}
		p.syntaxError(psess, fmt.Sprintf("%s used as value", str))
	}

	p.xnest = outer
	return
}

func (p *parser) ifStmt(psess *PackageSession) *IfStmt {
	if trace {
		defer p.trace("ifStmt")()
	}

	s := new(IfStmt)
	s.pos = p.pos()

	s.Init, s.Cond, _ = p.header(psess, _If)
	s.Then = p.blockStmt(psess, "if clause")

	if p.got(psess, _Else) {
		switch p.tok {
		case _If:
			s.Else = p.ifStmt(psess)
		case _Lbrace:
			s.Else = p.blockStmt(psess, "")
		default:
			p.syntaxError(psess, "else must be followed by if or statement block")
			p.advance(psess, _Name, _Rbrace)
		}
	}

	return s
}

func (p *parser) switchStmt(psess *PackageSession) *SwitchStmt {
	if trace {
		defer p.trace("switchStmt")()
	}

	s := new(SwitchStmt)
	s.pos = p.pos()

	s.Init, s.Tag, _ = p.header(psess, _Switch)

	if !p.got(psess, _Lbrace) {
		p.syntaxError(psess, "missing { after switch clause")
		p.advance(psess, _Case, _Default, _Rbrace)
	}
	for p.tok != _EOF && p.tok != _Rbrace {
		s.Body = append(s.Body, p.caseClause(psess))
	}
	s.Rbrace = p.pos()
	p.want(psess, _Rbrace)

	return s
}

func (p *parser) selectStmt(psess *PackageSession) *SelectStmt {
	if trace {
		defer p.trace("selectStmt")()
	}

	s := new(SelectStmt)
	s.pos = p.pos()

	p.want(psess, _Select)
	if !p.got(psess, _Lbrace) {
		p.syntaxError(psess, "missing { after select clause")
		p.advance(psess, _Case, _Default, _Rbrace)
	}
	for p.tok != _EOF && p.tok != _Rbrace {
		s.Body = append(s.Body, p.commClause(psess))
	}
	s.Rbrace = p.pos()
	p.want(psess, _Rbrace)

	return s
}

func (p *parser) caseClause(psess *PackageSession) *CaseClause {
	if trace {
		defer p.trace("caseClause")()
	}

	c := new(CaseClause)
	c.pos = p.pos()

	switch p.tok {
	case _Case:
		p.next(psess)
		c.Cases = p.exprList(psess)

	case _Default:
		p.next(psess)

	default:
		p.syntaxError(psess, "expecting case or default or }")
		p.advance(psess, _Colon, _Case, _Default, _Rbrace)
	}

	c.Colon = p.pos()
	p.want(psess, _Colon)
	c.Body = p.stmtList(psess)

	return c
}

func (p *parser) commClause(psess *PackageSession) *CommClause {
	if trace {
		defer p.trace("commClause")()
	}

	c := new(CommClause)
	c.pos = p.pos()

	switch p.tok {
	case _Case:
		p.next(psess)
		c.Comm = p.simpleStmt(psess, nil, 0)

	case _Default:
		p.next(psess)

	default:
		p.syntaxError(psess, "expecting case or default or }")
		p.advance(psess, _Colon, _Case, _Default, _Rbrace)
	}

	c.Colon = p.pos()
	p.want(psess, _Colon)
	c.Body = p.stmtList(psess)

	return c
}

// Statement =
// 	Declaration | LabeledStmt | SimpleStmt |
// 	GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
// 	FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
// 	DeferStmt .
func (p *parser) stmtOrNil(psess *PackageSession) Stmt {
	if trace {
		defer p.trace("stmt " + p.tok.String(psess))()
	}

	if p.tok == _Name {
		lhs := p.exprList(psess)
		if label, ok := lhs.(*Name); ok && p.tok == _Colon {
			return p.labeledStmtOrNil(psess, label)
		}
		return p.simpleStmt(psess, lhs, 0)
	}

	switch p.tok {
	case _Lbrace:
		return p.blockStmt(psess, "")

	case _Var:
		return p.declStmt(psess, p.varDecl)

	case _Const:
		return p.declStmt(psess, p.constDecl)

	case _Type:
		return p.declStmt(psess, p.typeDecl)

	case _Operator, _Star:
		switch p.op {
		case Add, Sub, Mul, And, Xor, Not:
			return p.simpleStmt(psess, nil, 0)
		}

	case _Literal, _Func, _Lparen,
		_Lbrack, _Struct, _Map, _Chan, _Interface,
		_Arrow:
		return p.simpleStmt(psess, nil, 0)

	case _For:
		return p.forStmt(psess)

	case _Switch:
		return p.switchStmt(psess)

	case _Select:
		return p.selectStmt(psess)

	case _If:
		return p.ifStmt(psess)

	case _Fallthrough:
		s := new(BranchStmt)
		s.pos = p.pos()
		p.next(psess)
		s.Tok = _Fallthrough
		return s

	case _Break, _Continue:
		s := new(BranchStmt)
		s.pos = p.pos()
		s.Tok = p.tok
		p.next(psess)
		if p.tok == _Name {
			s.Label = p.name(psess)
		}
		return s

	case _Go, _Defer:
		return p.callStmt(psess)

	case _Goto:
		s := new(BranchStmt)
		s.pos = p.pos()
		s.Tok = _Goto
		p.next(psess)
		s.Label = p.name(psess)
		return s

	case _Return:
		s := new(ReturnStmt)
		s.pos = p.pos()
		p.next(psess)
		if p.tok != _Semi && p.tok != _Rbrace {
			s.Results = p.exprList(psess)
		}
		return s

	case _Semi:
		s := new(EmptyStmt)
		s.pos = p.pos()
		return s
	}

	return nil
}

// StatementList = { Statement ";" } .
func (p *parser) stmtList(psess *PackageSession) (l []Stmt) {
	if trace {
		defer p.trace("stmtList")()
	}

	for p.tok != _EOF && p.tok != _Rbrace && p.tok != _Case && p.tok != _Default {
		s := p.stmtOrNil(psess)
		if s == nil {
			break
		}
		l = append(l, s)

		if !p.got(psess, _Semi) && p.tok != _Rbrace {
			p.syntaxError(psess, "at end of statement")
			p.advance(psess, _Semi, _Rbrace, _Case, _Default)
			p.got(psess, _Semi)
		}
	}
	return
}

// Arguments = "(" [ ( ExpressionList | Type [ "," ExpressionList ] ) [ "..." ] [ "," ] ] ")" .
func (p *parser) argList(psess *PackageSession) (list []Expr, hasDots bool) {
	if trace {
		defer p.trace("argList")()
	}

	p.xnest++
	p.list(psess, _Lparen, _Comma, _Rparen, func() bool {
		list = append(list, p.expr(psess))
		hasDots = p.got(psess, _DotDotDot)
		return hasDots
	})
	p.xnest--

	return
}

func (p *parser) newName(value string) *Name {
	n := new(Name)
	n.pos = p.pos()
	n.Value = value
	return n
}

func (p *parser) name(psess *PackageSession) *Name {

	if p.tok == _Name {
		n := p.newName(p.lit)
		p.next(psess)
		return n
	}

	n := p.newName("_")
	p.syntaxError(psess, "expecting name")
	p.advance(psess)
	return n
}

// IdentifierList = identifier { "," identifier } .
// The first name must be provided.
func (p *parser) nameList(psess *PackageSession, first *Name) []*Name {
	if trace {
		defer p.trace("nameList")()
	}

	if debug && first == nil {
		panic("first name not provided")
	}

	l := []*Name{first}
	for p.got(psess, _Comma) {
		l = append(l, p.name(psess))
	}

	return l
}

// The first name may be provided, or nil.
func (p *parser) qualifiedName(psess *PackageSession, name *Name) Expr {
	if trace {
		defer p.trace("qualifiedName")()
	}

	switch {
	case name != nil:

	case p.tok == _Name:
		name = p.name(psess)
	default:
		name = p.newName("_")
		p.syntaxError(psess, "expecting name")
		p.advance(psess, _Dot, _Semi, _Rbrace)
	}

	return p.dotname(psess, name)
}

// ExpressionList = Expression { "," Expression } .
func (p *parser) exprList(psess *PackageSession) Expr {
	if trace {
		defer p.trace("exprList")()
	}

	x := p.expr(psess)
	if p.got(psess, _Comma) {
		list := []Expr{x, p.expr(psess)}
		for p.got(psess, _Comma) {
			list = append(list, p.expr(psess))
		}
		t := new(ListExpr)
		t.pos = x.Pos()
		t.ElemList = list
		x = t
	}
	return x
}

// unparen removes all parentheses around an expression.
func unparen(x Expr) Expr {
	for {
		p, ok := x.(*ParenExpr)
		if !ok {
			break
		}
		x = p.X
	}
	return x
}

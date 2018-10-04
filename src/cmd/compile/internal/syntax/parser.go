// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
		// Error and directive handler for scanner.
		// Because the (line, col) positions passed to the
		// handler is always at or after the current reading
		// position, it is safe to use the most recent position
		// base to compute the corresponding Pos value.
		func(line, col uint, msg string) {
			if msg[0] != '/' {
				p.errorAt(p.posAt(line, col), msg)
				return
			}

			// otherwise it must be a comment containing a line or go: directive
			text := commentText(msg)
			if strings.HasPrefix(text, "line ") {
				var pos Pos // position immediately following the comment
				if msg[1] == '/' {
					// line comment (newline is part of the comment)
					pos = MakePos(p.file, line+1, colbase)
				} else {
					// regular comment
					// (if the comment spans multiple lines it's not
					// a valid line directive and will be discarded
					// by updateBase)
					pos = MakePos(p.file, line, col+uint(len(msg)))
				}
				p.updateBase(pos, line, col+2+5, text[5:]) // +2 to skip over // or /*
				return
			}

			// go: directive (but be conservative and test)
			if pragh != nil && strings.HasPrefix(text, "go:") {
				p.pragma |= pragh(p.posAt(line, col+2), text) // +2 to skip over // or /*
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
		return // ignore (not a line directive)
	}
	// i > 0

	if !ok {
		// text has a suffix :xxx but xxx is not a number
		p.errorAt(p.posAt(tline, tcol+i), "invalid line number: "+text[i:])
		return
	}

	var line, col uint
	i2, n2, ok2 := trailingDigits(text[:i-1])
	if ok2 {
		//line filename:line:col
		i, i2 = i2, i
		line, col = n2, n
		if col == 0 || col > PosMax {
			p.errorAt(p.posAt(tline, tcol+i2), "invalid column number: "+text[i2:])
			return
		}
		text = text[:i2-1] // lop off ":col"
	} else {
		//line filename:line
		line = n
	}

	if line == 0 || line > PosMax {
		p.errorAt(p.posAt(tline, tcol+i), "invalid line number: "+text[i:])
		return
	}

	// If we have a column (//line filename:line:col form),
	// an empty filename means to use the previous filename.
	filename := text[:i-1] // lop off ":line"
	if filename == "" && ok2 {
		filename = p.base.Filename()
	}

	p.base = NewLineBase(pos, filename, line, col)
}

func commentText(s string) string {
	if s[:2] == "/*" {
		return s[2 : len(s)-2] // lop off /* and */
	}

	// line comment (does not include newline)
	// (on Windows, the line comment may end in \r\n)
	i := len(s)
	if s[i-1] == '\r' {
		i--
	}
	return s[2:i] // lop off //, and \r at end, if any
}

func trailingDigits(text string) (uint, uint, bool) {
	// Want to use LastIndexByte below but it's not defined in Go1.4 and bootstrap fails.
	i := strings.LastIndex(text, ":") // look from right (Windows filenames may contain ':')
	if i < 0 {
		return 0, 0, false // no ":"
	}
	// i >= 0
	n, err := strconv.ParseUint(text[i+1:], 10, 0)
	return uint(i + 1), uint(n), err == nil
}

func (p *parser) got(pstate *PackageState, tok token) bool {
	if p.tok == tok {
		p.next(pstate)
		return true
	}
	return false
}

func (p *parser) want(pstate *PackageState, tok token) {
	if !p.got(pstate, tok) {
		p.syntaxError(pstate, "expecting "+pstate.tokstring(tok))
		p.advance(pstate)
	}
}

// ----------------------------------------------------------------------------
// Error handling

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
func (p *parser) syntaxErrorAt(pstate *PackageState, pos Pos, msg string) {
	if trace {
		p.print("syntax error: " + msg)
	}

	if p.tok == _EOF && p.first != nil {
		return // avoid meaningless follow-up errors
	}

	// add punctuation etc. as needed to msg
	switch {
	case msg == "":
	// nothing to do
	case strings.HasPrefix(msg, "in "), strings.HasPrefix(msg, "at "), strings.HasPrefix(msg, "after "):
		msg = " " + msg
	case strings.HasPrefix(msg, "expecting "):
		msg = ", " + msg
	default:
		// plain error - we don't care about current token
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
		tok = p.op.String(pstate)
	case _AssignOp:
		tok = p.op.String(pstate) + "="
	case _IncOp:
		tok = p.op.String(pstate)
		tok += tok
	default:
		tok = pstate.tokstring(p.tok)
	}

	p.errorAt(pos, "syntax error: unexpected "+tok+msg)
}

// tokstring returns the English word for selected punctuation tokens
// for more readable error messages.
func (pstate *PackageState) tokstring(tok token) string {
	switch tok {
	case _Comma:
		return "comma"
	case _Semi:
		return "semicolon or newline"
	}
	return tok.String(pstate)
}

// Convenience methods using the current token position.
func (p *parser) pos() Pos                                     { return p.posAt(p.line, p.col) }
func (p *parser) syntaxError(pstate *PackageState, msg string) { p.syntaxErrorAt(pstate, p.pos(), msg) }

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
func (p *parser) advance(pstate *PackageState, followlist ...token) {
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
			p.print("skip " + p.tok.String(pstate))
		}
		p.next(pstate)
		if len(followlist) == 0 {
			break
		}
	}

	if trace {
		p.print("next " + p.tok.String(pstate))
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
			panic(x) // skip print_trace
		}
		p.print(")")
	}
}

func (p *parser) print(msg string) {
	fmt.Printf("%5d: %s%s\n", p.line, p.indent, msg)
}

// ----------------------------------------------------------------------------
// Package files
//
// Parse methods are annotated with matching Go productions as appropriate.
// The annotations are intended as guidelines only since a single Go grammar
// rule may be covered by multiple parse methods and vice versa.
//
// Excluding methods returning slices, parse methods named xOrNil may return
// nil; all others are expected to return a valid non-nil node.

// SourceFile = PackageClause ";" { ImportDecl ";" } { TopLevelDecl ";" } .
func (p *parser) fileOrNil(pstate *PackageState) *File {
	if trace {
		defer p.trace("file")()
	}

	f := new(File)
	f.pos = p.pos()

	// PackageClause
	if !p.got(pstate, _Package) {
		p.syntaxError(pstate, "package statement must be first")
		return nil
	}
	f.PkgName = p.name(pstate)
	p.want(pstate, _Semi)

	// don't bother continuing if package clause has errors
	if p.first != nil {
		return nil
	}

	// { ImportDecl ";" }
	for p.got(pstate, _Import) {
		f.DeclList = p.appendGroup(pstate, f.DeclList, p.importDecl)
		p.want(pstate, _Semi)
	}

	// { TopLevelDecl ";" }
	for p.tok != _EOF {
		switch p.tok {
		case _Const:
			p.next(pstate)
			f.DeclList = p.appendGroup(pstate, f.DeclList, p.constDecl)

		case _Type:
			p.next(pstate)
			f.DeclList = p.appendGroup(pstate, f.DeclList, p.typeDecl)

		case _Var:
			p.next(pstate)
			f.DeclList = p.appendGroup(pstate, f.DeclList, p.varDecl)

		case _Func:
			p.next(pstate)
			if d := p.funcDeclOrNil(pstate); d != nil {
				f.DeclList = append(f.DeclList, d)
			}

		default:
			if p.tok == _Lbrace && len(f.DeclList) > 0 && isEmptyFuncDecl(f.DeclList[len(f.DeclList)-1]) {
				// opening { of function declaration on next line
				p.syntaxError(pstate, "unexpected semicolon or newline before {")
			} else {
				p.syntaxError(pstate, "non-declaration statement outside function body")
			}
			p.advance(pstate, _Const, _Type, _Var, _Func)
			continue
		}

		// Reset p.pragma BEFORE advancing to the next token (consuming ';')
		// since comments before may set pragmas for the next function decl.
		p.pragma = 0

		if p.tok != _EOF && !p.got(pstate, _Semi) {
			p.syntaxError(pstate, "after top level declaration")
			p.advance(pstate, _Const, _Type, _Var, _Func)
		}
	}
	// p.tok == _EOF

	f.Lines = p.source.line

	return f
}

func isEmptyFuncDecl(dcl Decl) bool {
	f, ok := dcl.(*FuncDecl)
	return ok && f.Body == nil
}

// ----------------------------------------------------------------------------
// Declarations

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
func (p *parser) list(pstate *PackageState, open, sep, close token, f func() bool) Pos {
	p.want(pstate, open)

	var done bool
	for p.tok != _EOF && p.tok != close && !done {
		done = f()
		// sep is optional before close
		if !p.got(pstate, sep) && p.tok != close {
			p.syntaxError(pstate, fmt.Sprintf("expecting %s or %s", pstate.tokstring(sep), pstate.tokstring(close)))
			p.advance(pstate, _Rparen, _Rbrack, _Rbrace)
			if p.tok != close {
				// position could be better but we had an error so we don't care
				return p.pos()
			}
		}
	}

	pos := p.pos()
	p.want(pstate, close)
	return pos
}

// appendGroup(f) = f | "(" { f ";" } ")" . // ";" is optional before ")"
func (p *parser) appendGroup(pstate *PackageState, list []Decl, f func(*Group) Decl) []Decl {
	if p.tok == _Lparen {
		g := new(Group)
		p.list(pstate, _Lparen, _Semi, _Rparen, func() bool {
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
func (p *parser) importDecl(pstate *PackageState, group *Group) Decl {
	if trace {
		defer p.trace("importDecl")()
	}

	d := new(ImportDecl)
	d.pos = p.pos()

	switch p.tok {
	case _Name:
		d.LocalPkgName = p.name(pstate)
	case _Dot:
		d.LocalPkgName = p.newName(".")
		p.next(pstate)
	}
	d.Path = p.oliteral(pstate)
	if d.Path == nil {
		p.syntaxError(pstate, "missing import path")
		p.advance(pstate, _Semi, _Rparen)
		return nil
	}
	d.Group = group

	return d
}

// ConstSpec = IdentifierList [ [ Type ] "=" ExpressionList ] .
func (p *parser) constDecl(pstate *PackageState, group *Group) Decl {
	if trace {
		defer p.trace("constDecl")()
	}

	d := new(ConstDecl)
	d.pos = p.pos()

	d.NameList = p.nameList(pstate, p.name(pstate))
	if p.tok != _EOF && p.tok != _Semi && p.tok != _Rparen {
		d.Type = p.typeOrNil(pstate)
		if p.got(pstate, _Assign) {
			d.Values = p.exprList(pstate)
		}
	}
	d.Group = group

	return d
}

// TypeSpec = identifier [ "=" ] Type .
func (p *parser) typeDecl(pstate *PackageState, group *Group) Decl {
	if trace {
		defer p.trace("typeDecl")()
	}

	d := new(TypeDecl)
	d.pos = p.pos()

	d.Name = p.name(pstate)
	d.Alias = p.got(pstate, _Assign)
	d.Type = p.typeOrNil(pstate)
	if d.Type == nil {
		d.Type = p.bad()
		p.syntaxError(pstate, "in type declaration")
		p.advance(pstate, _Semi, _Rparen)
	}
	d.Group = group
	d.Pragma = p.pragma

	return d
}

// VarSpec = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .
func (p *parser) varDecl(pstate *PackageState, group *Group) Decl {
	if trace {
		defer p.trace("varDecl")()
	}

	d := new(VarDecl)
	d.pos = p.pos()

	d.NameList = p.nameList(pstate, p.name(pstate))
	if p.got(pstate, _Assign) {
		d.Values = p.exprList(pstate)
	} else {
		d.Type = p.type_(pstate)
		if p.got(pstate, _Assign) {
			d.Values = p.exprList(pstate)
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
func (p *parser) funcDeclOrNil(pstate *PackageState) *FuncDecl {
	if trace {
		defer p.trace("funcDecl")()
	}

	f := new(FuncDecl)
	f.pos = p.pos()

	if p.tok == _Lparen {
		rcvr := p.paramList(pstate)
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
		p.syntaxError(pstate, "expecting name or (")
		p.advance(pstate, _Lbrace, _Semi)
		return nil
	}

	f.Name = p.name(pstate)
	f.Type = p.funcType(pstate)
	if p.tok == _Lbrace {
		f.Body = p.funcBody(pstate)
	}
	f.Pragma = p.pragma

	return f
}

func (p *parser) funcBody(pstate *PackageState) *BlockStmt {
	p.fnest++
	errcnt := p.errcnt
	body := p.blockStmt(pstate, "")
	p.fnest--

	// Don't check branches if there were syntax errors in the function
	// as it may lead to spurious errors (e.g., see test/switch2.go) or
	// possibly crashes due to incomplete syntax trees.
	if p.mode&CheckBranches != 0 && errcnt == p.errcnt {
		pstate.checkBranches(body, p.errh)
	}

	return body
}

// ----------------------------------------------------------------------------
// Expressions

func (p *parser) expr(pstate *PackageState) Expr {
	if trace {
		defer p.trace("expr")()
	}

	return p.binaryExpr(pstate, 0)
}

// Expression = UnaryExpr | Expression binary_op Expression .
func (p *parser) binaryExpr(pstate *PackageState, prec int) Expr {
	// don't trace binaryExpr - only leads to overly nested trace output

	x := p.unaryExpr(pstate)
	for (p.tok == _Operator || p.tok == _Star) && p.prec > prec {
		t := new(Operation)
		t.pos = p.pos()
		t.Op = p.op
		t.X = x
		tprec := p.prec
		p.next(pstate)
		t.Y = p.binaryExpr(pstate, tprec)
		x = t
	}
	return x
}

// UnaryExpr = PrimaryExpr | unary_op UnaryExpr .
func (p *parser) unaryExpr(pstate *PackageState) Expr {
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
			p.next(pstate)
			x.X = p.unaryExpr(pstate)
			return x

		case And:
			x := new(Operation)
			x.pos = p.pos()
			x.Op = And
			p.next(pstate)
			// unaryExpr may have returned a parenthesized composite literal
			// (see comment in operand) - remove parentheses if any
			x.X = unparen(p.unaryExpr(pstate))
			return x
		}

	case _Arrow:
		// receive op (<-x) or receive-only channel (<-chan E)
		pos := p.pos()
		p.next(pstate)

		// If the next token is _Chan we still don't know if it is
		// a channel (<-chan int) or a receive op (<-chan int(ch)).
		// We only know once we have found the end of the unaryExpr.

		x := p.unaryExpr(pstate)

		// There are two cases:
		//
		//   <-chan...  => <-x is a channel type
		//   <-x        => <-x is a receive operation
		//
		// In the first case, <- must be re-associated with
		// the channel type parsed already:
		//
		//   <-(chan E)   =>  (<-chan E)
		//   <-(chan<-E)  =>  (<-chan (<-E))

		if _, ok := x.(*ChanType); ok {
			// x is a channel type => re-associate <-
			dir := SendOnly
			t := x
			for dir == SendOnly {
				c, ok := t.(*ChanType)
				if !ok {
					break
				}
				dir = c.Dir
				if dir == RecvOnly {
					// t is type <-chan E but <-<-chan E is not permitted
					// (report same error as for "type _ <-<-chan E")
					p.syntaxError(pstate, "unexpected <-, expecting chan")
					// already progressed, no need to advance
				}
				c.Dir = RecvOnly
				t = c.Elem
			}
			if dir == SendOnly {
				// channel dir is <- but channel element E is not a channel
				// (report same error as for "type _ <-chan<-E")
				p.syntaxError(pstate, fmt.Sprintf("unexpected %s, expecting chan", pstate.String(t)))
				// already progressed, no need to advance
			}
			return x
		}

		// x is not a channel type => we have a receive op
		o := new(Operation)
		o.pos = pos
		o.Op = Recv
		o.X = x
		return o
	}

	// TODO(mdempsky): We need parens here so we can report an
	// error for "(x) := true". It should be possible to detect
	// and reject that more efficiently though.
	return p.pexpr(pstate, true)
}

// callStmt parses call-like statements that can be preceded by 'defer' and 'go'.
func (p *parser) callStmt(pstate *PackageState) *CallStmt {
	if trace {
		defer p.trace("callStmt")()
	}

	s := new(CallStmt)
	s.pos = p.pos()
	s.Tok = p.tok // _Defer or _Go
	p.next(pstate)

	x := p.pexpr(pstate, p.tok == _Lparen) // keep_parens so we can report error below
	if t := unparen(x); t != x {
		p.errorAt(x.Pos(), fmt.Sprintf("expression in %s must not be parenthesized", s.Tok))
		// already progressed, no need to advance
		x = t
	}

	cx, ok := x.(*CallExpr)
	if !ok {
		p.errorAt(x.Pos(), fmt.Sprintf("expression in %s must be function call", s.Tok))
		// already progressed, no need to advance
		cx = new(CallExpr)
		cx.pos = x.Pos()
		cx.Fun = x // assume common error of missing parentheses (function invocation)
	}

	s.Call = cx
	return s
}

// Operand     = Literal | OperandName | MethodExpr | "(" Expression ")" .
// Literal     = BasicLit | CompositeLit | FunctionLit .
// BasicLit    = int_lit | float_lit | imaginary_lit | rune_lit | string_lit .
// OperandName = identifier | QualifiedIdent.
func (p *parser) operand(pstate *PackageState, keep_parens bool) Expr {
	if trace {
		defer p.trace("operand " + p.tok.String(pstate))()
	}

	switch p.tok {
	case _Name:
		return p.name(pstate)

	case _Literal:
		return p.oliteral(pstate)

	case _Lparen:
		pos := p.pos()
		p.next(pstate)
		p.xnest++
		x := p.expr(pstate)
		p.xnest--
		p.want(pstate, _Rparen)

		// Optimization: Record presence of ()'s only where needed
		// for error reporting. Don't bother in other cases; it is
		// just a waste of memory and time.

		// Parentheses are not permitted on lhs of := .
		// switch x.Op {
		// case ONAME, ONONAME, OPACK, OTYPE, OLITERAL, OTYPESW:
		// 	keep_parens = true
		// }

		// Parentheses are not permitted around T in a composite
		// literal T{}. If the next token is a {, assume x is a
		// composite literal type T (it may not be, { could be
		// the opening brace of a block, but we don't know yet).
		if p.tok == _Lbrace {
			keep_parens = true
		}

		// Parentheses are also not permitted around the expression
		// in a go/defer statement. In that case, operand is called
		// with keep_parens set.
		if keep_parens {
			px := new(ParenExpr)
			px.pos = pos
			px.X = x
			x = px
		}
		return x

	case _Func:
		pos := p.pos()
		p.next(pstate)
		t := p.funcType(pstate)
		if p.tok == _Lbrace {
			p.xnest++

			f := new(FuncLit)
			f.pos = pos
			f.Type = t
			f.Body = p.funcBody(pstate)

			p.xnest--
			return f
		}
		return t

	case _Lbrack, _Chan, _Map, _Struct, _Interface:
		return p.type_(pstate) // othertype

	default:
		x := p.bad()
		p.syntaxError(pstate, "expecting expression")
		p.advance(pstate)
		return x
	}

	// Syntactically, composite literals are operands. Because a complit
	// type may be a qualified identifier which is handled by pexpr
	// (together with selector expressions), complits are parsed there
	// as well (operand is only called from pexpr).
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
func (p *parser) pexpr(pstate *PackageState, keep_parens bool) Expr {
	if trace {
		defer p.trace("pexpr")()
	}

	x := p.operand(pstate, keep_parens)

loop:
	for {
		pos := p.pos()
		switch p.tok {
		case _Dot:
			p.next(pstate)
			switch p.tok {
			case _Name:
				// pexpr '.' sym
				t := new(SelectorExpr)
				t.pos = pos
				t.X = x
				t.Sel = p.name(pstate)
				x = t

			case _Lparen:
				p.next(pstate)
				if p.got(pstate, _Type) {
					t := new(TypeSwitchGuard)
					// t.Lhs is filled in by parser.simpleStmt
					t.pos = pos
					t.X = x
					x = t
				} else {
					t := new(AssertExpr)
					t.pos = pos
					t.X = x
					t.Type = p.type_(pstate)
					x = t
				}
				p.want(pstate, _Rparen)

			default:
				p.syntaxError(pstate, "expecting name or (")
				p.advance(pstate, _Semi, _Rparen)
			}

		case _Lbrack:
			p.next(pstate)
			p.xnest++

			var i Expr
			if p.tok != _Colon {
				i = p.expr(pstate)
				if p.got(pstate, _Rbrack) {
					// x[i]
					t := new(IndexExpr)
					t.pos = pos
					t.X = x
					t.Index = i
					x = t
					p.xnest--
					break
				}
			}

			// x[i:...
			t := new(SliceExpr)
			t.pos = pos
			t.X = x
			t.Index[0] = i
			p.want(pstate, _Colon)
			if p.tok != _Colon && p.tok != _Rbrack {
				// x[i:j...
				t.Index[1] = p.expr(pstate)
			}
			if p.got(pstate, _Colon) {
				t.Full = true
				// x[i:j:...]
				if t.Index[1] == nil {
					p.error("middle index required in 3-index slice")
				}
				if p.tok != _Rbrack {
					// x[i:j:k...
					t.Index[2] = p.expr(pstate)
				} else {
					p.error("final index required in 3-index slice")
				}
			}
			p.want(pstate, _Rbrack)

			x = t
			p.xnest--

		case _Lparen:
			t := new(CallExpr)
			t.pos = pos
			t.Fun = x
			t.ArgList, t.HasDots = p.argList(pstate)
			x = t

		case _Lbrace:
			// operand may have returned a parenthesized complit
			// type; accept it but complain if we have a complit
			t := unparen(x)
			// determine if '{' belongs to a composite literal or a block statement
			complit_ok := false
			switch t.(type) {
			case *Name, *SelectorExpr:
				if p.xnest >= 0 {
					// x is considered a composite literal type
					complit_ok = true
				}
			case *ArrayType, *SliceType, *StructType, *MapType:
				// x is a comptype
				complit_ok = true
			}
			if !complit_ok {
				break loop
			}
			if t != x {
				p.syntaxError(pstate, "cannot parenthesize type in composite literal")
				// already progressed, no need to advance
			}
			n := p.complitexpr(pstate)
			n.Type = x
			x = n

		default:
			break loop
		}
	}

	return x
}

// Element = Expression | LiteralValue .
func (p *parser) bare_complitexpr(pstate *PackageState) Expr {
	if trace {
		defer p.trace("bare_complitexpr")()
	}

	if p.tok == _Lbrace {
		// '{' start_complit braced_keyval_list '}'
		return p.complitexpr(pstate)
	}

	return p.expr(pstate)
}

// LiteralValue = "{" [ ElementList [ "," ] ] "}" .
func (p *parser) complitexpr(pstate *PackageState) *CompositeLit {
	if trace {
		defer p.trace("complitexpr")()
	}

	x := new(CompositeLit)
	x.pos = p.pos()

	p.xnest++
	x.Rbrace = p.list(pstate, _Lbrace, _Comma, _Rbrace, func() bool {
		// value
		e := p.bare_complitexpr(pstate)
		if p.tok == _Colon {
			// key ':' value
			l := new(KeyValueExpr)
			l.pos = p.pos()
			p.next(pstate)
			l.Key = e
			l.Value = p.bare_complitexpr(pstate)
			e = l
			x.NKeys++
		}
		x.ElemList = append(x.ElemList, e)
		return false
	})
	p.xnest--

	return x
}

// ----------------------------------------------------------------------------
// Types

func (p *parser) type_(pstate *PackageState) Expr {
	if trace {
		defer p.trace("type_")()
	}

	typ := p.typeOrNil(pstate)
	if typ == nil {
		typ = p.bad()
		p.syntaxError(pstate, "expecting type")
		p.advance(pstate, _Comma, _Colon, _Semi, _Rparen, _Rbrack, _Rbrace)
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
func (p *parser) typeOrNil(pstate *PackageState) Expr {
	if trace {
		defer p.trace("typeOrNil")()
	}

	pos := p.pos()
	switch p.tok {
	case _Star:
		// ptrtype
		p.next(pstate)
		return newIndirect(pos, p.type_(pstate))

	case _Arrow:
		// recvchantype
		p.next(pstate)
		p.want(pstate, _Chan)
		t := new(ChanType)
		t.pos = pos
		t.Dir = RecvOnly
		t.Elem = p.chanElem(pstate)
		return t

	case _Func:
		// fntype
		p.next(pstate)
		return p.funcType(pstate)

	case _Lbrack:
		// '[' oexpr ']' ntype
		// '[' _DotDotDot ']' ntype
		p.next(pstate)
		p.xnest++
		if p.got(pstate, _Rbrack) {
			// []T
			p.xnest--
			t := new(SliceType)
			t.pos = pos
			t.Elem = p.type_(pstate)
			return t
		}

		// [n]T
		t := new(ArrayType)
		t.pos = pos
		if !p.got(pstate, _DotDotDot) {
			t.Len = p.expr(pstate)
		}
		p.want(pstate, _Rbrack)
		p.xnest--
		t.Elem = p.type_(pstate)
		return t

	case _Chan:
		// _Chan non_recvchantype
		// _Chan _Comm ntype
		p.next(pstate)
		t := new(ChanType)
		t.pos = pos
		if p.got(pstate, _Arrow) {
			t.Dir = SendOnly
		}
		t.Elem = p.chanElem(pstate)
		return t

	case _Map:
		// _Map '[' ntype ']' ntype
		p.next(pstate)
		p.want(pstate, _Lbrack)
		t := new(MapType)
		t.pos = pos
		t.Key = p.type_(pstate)
		p.want(pstate, _Rbrack)
		t.Value = p.type_(pstate)
		return t

	case _Struct:
		return p.structType(pstate)

	case _Interface:
		return p.interfaceType(pstate)

	case _Name:
		return p.dotname(pstate, p.name(pstate))

	case _Lparen:
		p.next(pstate)
		t := p.type_(pstate)
		p.want(pstate, _Rparen)
		return t
	}

	return nil
}

func (p *parser) funcType(pstate *PackageState) *FuncType {
	if trace {
		defer p.trace("funcType")()
	}

	typ := new(FuncType)
	typ.pos = p.pos()
	typ.ParamList = p.paramList(pstate)
	typ.ResultList = p.funcResult(pstate)

	return typ
}

func (p *parser) chanElem(pstate *PackageState) Expr {
	if trace {
		defer p.trace("chanElem")()
	}

	typ := p.typeOrNil(pstate)
	if typ == nil {
		typ = p.bad()
		p.syntaxError(pstate, "missing channel element type")
		// assume element type is simply absent - don't advance
	}

	return typ
}

func (p *parser) dotname(pstate *PackageState, name *Name) Expr {
	if trace {
		defer p.trace("dotname")()
	}

	if p.tok == _Dot {
		s := new(SelectorExpr)
		s.pos = p.pos()
		p.next(pstate)
		s.X = name
		s.Sel = p.name(pstate)
		return s
	}
	return name
}

// StructType = "struct" "{" { FieldDecl ";" } "}" .
func (p *parser) structType(pstate *PackageState) *StructType {
	if trace {
		defer p.trace("structType")()
	}

	typ := new(StructType)
	typ.pos = p.pos()

	p.want(pstate, _Struct)
	p.list(pstate, _Lbrace, _Semi, _Rbrace, func() bool {
		p.fieldDecl(pstate, typ)
		return false
	})

	return typ
}

// InterfaceType = "interface" "{" { MethodSpec ";" } "}" .
func (p *parser) interfaceType(pstate *PackageState) *InterfaceType {
	if trace {
		defer p.trace("interfaceType")()
	}

	typ := new(InterfaceType)
	typ.pos = p.pos()

	p.want(pstate, _Interface)
	p.list(pstate, _Lbrace, _Semi, _Rbrace, func() bool {
		if m := p.methodDecl(pstate); m != nil {
			typ.MethodList = append(typ.MethodList, m)
		}
		return false
	})

	return typ
}

// Result = Parameters | Type .
func (p *parser) funcResult(pstate *PackageState) []*Field {
	if trace {
		defer p.trace("funcResult")()
	}

	if p.tok == _Lparen {
		return p.paramList(pstate)
	}

	pos := p.pos()
	if typ := p.typeOrNil(pstate); typ != nil {
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
func (p *parser) fieldDecl(pstate *PackageState, styp *StructType) {
	if trace {
		defer p.trace("fieldDecl")()
	}

	pos := p.pos()
	switch p.tok {
	case _Name:
		name := p.name(pstate)
		if p.tok == _Dot || p.tok == _Literal || p.tok == _Semi || p.tok == _Rbrace {
			// embed oliteral
			typ := p.qualifiedName(pstate, name)
			tag := p.oliteral(pstate)
			p.addField(styp, pos, nil, typ, tag)
			return
		}

		// new_name_list ntype oliteral
		names := p.nameList(pstate, name)
		typ := p.type_(pstate)
		tag := p.oliteral(pstate)

		for _, name := range names {
			p.addField(styp, name.Pos(), name, typ, tag)
		}

	case _Lparen:
		p.next(pstate)
		if p.tok == _Star {
			// '(' '*' embed ')' oliteral
			pos := p.pos()
			p.next(pstate)
			typ := newIndirect(pos, p.qualifiedName(pstate, nil))
			p.want(pstate, _Rparen)
			tag := p.oliteral(pstate)
			p.addField(styp, pos, nil, typ, tag)
			p.syntaxError(pstate, "cannot parenthesize embedded type")

		} else {
			// '(' embed ')' oliteral
			typ := p.qualifiedName(pstate, nil)
			p.want(pstate, _Rparen)
			tag := p.oliteral(pstate)
			p.addField(styp, pos, nil, typ, tag)
			p.syntaxError(pstate, "cannot parenthesize embedded type")
		}

	case _Star:
		p.next(pstate)
		if p.got(pstate, _Lparen) {
			// '*' '(' embed ')' oliteral
			typ := newIndirect(pos, p.qualifiedName(pstate, nil))
			p.want(pstate, _Rparen)
			tag := p.oliteral(pstate)
			p.addField(styp, pos, nil, typ, tag)
			p.syntaxError(pstate, "cannot parenthesize embedded type")

		} else {
			// '*' embed oliteral
			typ := newIndirect(pos, p.qualifiedName(pstate, nil))
			tag := p.oliteral(pstate)
			p.addField(styp, pos, nil, typ, tag)
		}

	default:
		p.syntaxError(pstate, "expecting field name or embedded type")
		p.advance(pstate, _Semi, _Rbrace)
	}
}

func (p *parser) oliteral(pstate *PackageState) *BasicLit {
	if p.tok == _Literal {
		b := new(BasicLit)
		b.pos = p.pos()
		b.Value = p.lit
		b.Kind = p.kind
		p.next(pstate)
		return b
	}
	return nil
}

// MethodSpec        = MethodName Signature | InterfaceTypeName .
// MethodName        = identifier .
// InterfaceTypeName = TypeName .
func (p *parser) methodDecl(pstate *PackageState) *Field {
	if trace {
		defer p.trace("methodDecl")()
	}

	switch p.tok {
	case _Name:
		name := p.name(pstate)

		// accept potential name list but complain
		hasNameList := false
		for p.got(pstate, _Comma) {
			p.name(pstate)
			hasNameList = true
		}
		if hasNameList {
			p.syntaxError(pstate, "name list not allowed in interface type")
			// already progressed, no need to advance
		}

		f := new(Field)
		f.pos = name.Pos()
		if p.tok != _Lparen {
			// packname
			f.Type = p.qualifiedName(pstate, name)
			return f
		}

		f.Name = name
		f.Type = p.funcType(pstate)
		return f

	case _Lparen:
		p.syntaxError(pstate, "cannot parenthesize embedded type")
		f := new(Field)
		f.pos = p.pos()
		p.next(pstate)
		f.Type = p.qualifiedName(pstate, nil)
		p.want(pstate, _Rparen)
		return f

	default:
		p.syntaxError(pstate, "expecting method or interface name")
		p.advance(pstate, _Semi, _Rbrace)
		return nil
	}
}

// ParameterDecl = [ IdentifierList ] [ "..." ] Type .
func (p *parser) paramDeclOrNil(pstate *PackageState) *Field {
	if trace {
		defer p.trace("paramDecl")()
	}

	f := new(Field)
	f.pos = p.pos()

	switch p.tok {
	case _Name:
		f.Name = p.name(pstate)
		switch p.tok {
		case _Name, _Star, _Arrow, _Func, _Lbrack, _Chan, _Map, _Struct, _Interface, _Lparen:
			// sym name_or_type
			f.Type = p.type_(pstate)

		case _DotDotDot:
			// sym dotdotdot
			f.Type = p.dotsType(pstate)

		case _Dot:
			// name_or_type
			// from dotname
			f.Type = p.dotname(pstate, f.Name)
			f.Name = nil
		}

	case _Arrow, _Star, _Func, _Lbrack, _Chan, _Map, _Struct, _Interface, _Lparen:
		// name_or_type
		f.Type = p.type_(pstate)

	case _DotDotDot:
		// dotdotdot
		f.Type = p.dotsType(pstate)

	default:
		p.syntaxError(pstate, "expecting )")
		p.advance(pstate, _Comma, _Rparen)
		return nil
	}

	return f
}

// ...Type
func (p *parser) dotsType(pstate *PackageState) *DotsType {
	if trace {
		defer p.trace("dotsType")()
	}

	t := new(DotsType)
	t.pos = p.pos()

	p.want(pstate, _DotDotDot)
	t.Elem = p.typeOrNil(pstate)
	if t.Elem == nil {
		t.Elem = p.bad()
		p.syntaxError(pstate, "final argument in variadic function missing type")
	}

	return t
}

// Parameters    = "(" [ ParameterList [ "," ] ] ")" .
// ParameterList = ParameterDecl { "," ParameterDecl } .
func (p *parser) paramList(pstate *PackageState) (list []*Field) {
	if trace {
		defer p.trace("paramList")()
	}

	pos := p.pos()

	var named int // number of parameters that have an explicit name and type
	p.list(pstate, _Lparen, _Comma, _Rparen, func() bool {
		if par := p.paramDeclOrNil(pstate); par != nil {
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

	// distribute parameter types
	if named == 0 {
		// all unnamed => found names are named types
		for _, par := range list {
			if typ := par.Name; typ != nil {
				par.Type = typ
				par.Name = nil
			}
		}
	} else if named != len(list) {
		// some named => all must be named
		ok := true
		var typ Expr
		for i := len(list) - 1; i >= 0; i-- {
			if par := list[i]; par.Type != nil {
				typ = par.Type
				if par.Name == nil {
					ok = false
					n := p.newName("_")
					n.pos = typ.Pos() // correct position
					par.Name = n
				}
			} else if typ != nil {
				par.Type = typ
			} else {
				// par.Type == nil && typ == nil => we only have a par.Name
				ok = false
				t := p.bad()
				t.pos = par.Name.Pos() // correct position
				par.Type = t
			}
		}
		if !ok {
			p.syntaxErrorAt(pstate, pos, "mixed named and unnamed function parameters")
		}
	}

	return
}

func (p *parser) bad() *BadExpr {
	b := new(BadExpr)
	b.pos = p.pos()
	return b
}

// SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
func (p *parser) simpleStmt(pstate *PackageState, lhs Expr, keyword token) SimpleStmt {
	if trace {
		defer p.trace("simpleStmt")()
	}

	if keyword == _For && p.tok == _Range {
		// _Range expr
		if debug && lhs != nil {
			panic("invalid call of simpleStmt")
		}
		return p.newRangeClause(pstate, nil, false)
	}

	if lhs == nil {
		lhs = p.exprList(pstate)
	}

	if _, ok := lhs.(*ListExpr); !ok && p.tok != _Assign && p.tok != _Define {
		// expr
		pos := p.pos()
		switch p.tok {
		case _AssignOp:
			// lhs op= rhs
			op := p.op
			p.next(pstate)
			return p.newAssignStmt(pos, op, lhs, p.expr(pstate))

		case _IncOp:
			// lhs++ or lhs--
			op := p.op
			p.next(pstate)
			return p.newAssignStmt(pos, op, lhs, pstate.ImplicitOne)

		case _Arrow:
			// lhs <- rhs
			s := new(SendStmt)
			s.pos = pos
			p.next(pstate)
			s.Chan = lhs
			s.Value = p.expr(pstate)
			return s

		default:
			// expr
			s := new(ExprStmt)
			s.pos = lhs.Pos()
			s.X = lhs
			return s
		}
	}

	// expr_list
	switch p.tok {
	case _Assign, _Define:
		pos := p.pos()
		var op Operator
		if p.tok == _Define {
			op = Def
		}
		p.next(pstate)

		if keyword == _For && p.tok == _Range {
			// expr_list op= _Range expr
			return p.newRangeClause(pstate, lhs, op == Def)
		}

		// expr_list op= expr_list
		rhs := p.exprList(pstate)

		if x, ok := rhs.(*TypeSwitchGuard); ok && keyword == _Switch && op == Def {
			if lhs, ok := lhs.(*Name); ok {
				// switch … lhs := rhs.(type)
				x.Lhs = lhs
				s := new(ExprStmt)
				s.pos = x.Pos()
				s.X = x
				return s
			}
		}

		return p.newAssignStmt(pos, op, lhs, rhs)

	default:
		p.syntaxError(pstate, "expecting := or = or comma")
		p.advance(pstate, _Semi, _Rbrace)
		// make the best of what we have
		if x, ok := lhs.(*ListExpr); ok {
			lhs = x.ElemList[0]
		}
		s := new(ExprStmt)
		s.pos = lhs.Pos()
		s.X = lhs
		return s
	}
}

func (p *parser) newRangeClause(pstate *PackageState, lhs Expr, def bool) *RangeClause {
	r := new(RangeClause)
	r.pos = p.pos()
	p.next(pstate) // consume _Range
	r.Lhs = lhs
	r.Def = def
	r.X = p.expr(pstate)
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

func (p *parser) labeledStmtOrNil(pstate *PackageState, label *Name) Stmt {
	if trace {
		defer p.trace("labeledStmt")()
	}

	s := new(LabeledStmt)
	s.pos = p.pos()
	s.Label = label

	p.want(pstate, _Colon)

	if p.tok == _Rbrace {
		// We expect a statement (incl. an empty statement), which must be
		// terminated by a semicolon. Because semicolons may be omitted before
		// an _Rbrace, seeing an _Rbrace implies an empty statement.
		e := new(EmptyStmt)
		e.pos = p.pos()
		s.Stmt = e
		return s
	}

	s.Stmt = p.stmtOrNil(pstate)
	if s.Stmt != nil {
		return s
	}

	// report error at line of ':' token
	p.syntaxErrorAt(pstate, s.pos, "missing statement after label")
	// we are already at the end of the labeled statement - no need to advance
	return nil // avoids follow-on errors (see e.g., fixedbugs/bug274.go)
}

// context must be a non-empty string unless we know that p.tok == _Lbrace.
func (p *parser) blockStmt(pstate *PackageState, context string) *BlockStmt {
	if trace {
		defer p.trace("blockStmt")()
	}

	s := new(BlockStmt)
	s.pos = p.pos()

	// people coming from C may forget that braces are mandatory in Go
	if !p.got(pstate, _Lbrace) {
		p.syntaxError(pstate, "expecting { after "+context)
		p.advance(pstate, _Name, _Rbrace)
		s.Rbrace = p.pos() // in case we found "}"
		if p.got(pstate, _Rbrace) {
			return s
		}
	}

	s.List = p.stmtList(pstate)
	s.Rbrace = p.pos()
	p.want(pstate, _Rbrace)

	return s
}

func (p *parser) declStmt(pstate *PackageState, f func(*Group) Decl) *DeclStmt {
	if trace {
		defer p.trace("declStmt")()
	}

	s := new(DeclStmt)
	s.pos = p.pos()

	p.next(pstate) // _Const, _Type, or _Var
	s.DeclList = p.appendGroup(pstate, nil, f)

	return s
}

func (p *parser) forStmt(pstate *PackageState) Stmt {
	if trace {
		defer p.trace("forStmt")()
	}

	s := new(ForStmt)
	s.pos = p.pos()

	s.Init, s.Cond, s.Post = p.header(pstate, _For)
	s.Body = p.blockStmt(pstate, "for clause")

	return s
}

func (p *parser) header(pstate *PackageState, keyword token) (init SimpleStmt, cond Expr, post SimpleStmt) {
	p.want(pstate, keyword)

	if p.tok == _Lbrace {
		if keyword == _If {
			p.syntaxError(pstate, "missing condition in if statement")
		}
		return
	}
	// p.tok != _Lbrace

	outer := p.xnest
	p.xnest = -1

	if p.tok != _Semi {
		// accept potential varDecl but complain
		if p.got(pstate, _Var) {
			p.syntaxError(pstate, fmt.Sprintf("var declaration not allowed in %s initializer", keyword.String(pstate)))
		}
		init = p.simpleStmt(pstate, nil, keyword)
		// If we have a range clause, we are done (can only happen for keyword == _For).
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
			p.next(pstate)
		} else {
			// asking for a '{' rather than a ';' here leads to a better error message
			p.want(pstate, _Lbrace)
		}
		if keyword == _For {
			if p.tok != _Semi {
				if p.tok == _Lbrace {
					p.syntaxError(pstate, "expecting for loop condition")
					goto done
				}
				condStmt = p.simpleStmt(pstate, nil, 0 /* range not permitted */)
			}
			p.want(pstate, _Semi)
			if p.tok != _Lbrace {
				post = p.simpleStmt(pstate, nil, 0 /* range not permitted */)
				if a, _ := post.(*AssignStmt); a != nil && a.Op == Def {
					p.syntaxErrorAt(pstate, a.Pos(), "cannot declare in post statement of for loop")
				}
			}
		} else if p.tok != _Lbrace {
			condStmt = p.simpleStmt(pstate, nil, keyword)
		}
	} else {
		condStmt = init
		init = nil
	}

done:
	// unpack condStmt
	switch s := condStmt.(type) {
	case nil:
		if keyword == _If && semi.pos.IsKnown() {
			if semi.lit != "semicolon" {
				p.syntaxErrorAt(pstate, semi.pos, fmt.Sprintf("unexpected %s, expecting { after if clause", semi.lit))
			} else {
				p.syntaxErrorAt(pstate, semi.pos, "missing condition in if statement")
			}
		}
	case *ExprStmt:
		cond = s.X
	default:
		// A common syntax error is to write '=' instead of '==',
		// which turns an expression into an assignment. Provide
		// a more explicit error message in that case to prevent
		// further confusion.
		str := pstate.String(s)
		if as, ok := s.(*AssignStmt); ok && as.Op == 0 {
			str = "assignment " + str
		}
		p.syntaxError(pstate, fmt.Sprintf("%s used as value", str))
	}

	p.xnest = outer
	return
}

func (p *parser) ifStmt(pstate *PackageState) *IfStmt {
	if trace {
		defer p.trace("ifStmt")()
	}

	s := new(IfStmt)
	s.pos = p.pos()

	s.Init, s.Cond, _ = p.header(pstate, _If)
	s.Then = p.blockStmt(pstate, "if clause")

	if p.got(pstate, _Else) {
		switch p.tok {
		case _If:
			s.Else = p.ifStmt(pstate)
		case _Lbrace:
			s.Else = p.blockStmt(pstate, "")
		default:
			p.syntaxError(pstate, "else must be followed by if or statement block")
			p.advance(pstate, _Name, _Rbrace)
		}
	}

	return s
}

func (p *parser) switchStmt(pstate *PackageState) *SwitchStmt {
	if trace {
		defer p.trace("switchStmt")()
	}

	s := new(SwitchStmt)
	s.pos = p.pos()

	s.Init, s.Tag, _ = p.header(pstate, _Switch)

	if !p.got(pstate, _Lbrace) {
		p.syntaxError(pstate, "missing { after switch clause")
		p.advance(pstate, _Case, _Default, _Rbrace)
	}
	for p.tok != _EOF && p.tok != _Rbrace {
		s.Body = append(s.Body, p.caseClause(pstate))
	}
	s.Rbrace = p.pos()
	p.want(pstate, _Rbrace)

	return s
}

func (p *parser) selectStmt(pstate *PackageState) *SelectStmt {
	if trace {
		defer p.trace("selectStmt")()
	}

	s := new(SelectStmt)
	s.pos = p.pos()

	p.want(pstate, _Select)
	if !p.got(pstate, _Lbrace) {
		p.syntaxError(pstate, "missing { after select clause")
		p.advance(pstate, _Case, _Default, _Rbrace)
	}
	for p.tok != _EOF && p.tok != _Rbrace {
		s.Body = append(s.Body, p.commClause(pstate))
	}
	s.Rbrace = p.pos()
	p.want(pstate, _Rbrace)

	return s
}

func (p *parser) caseClause(pstate *PackageState) *CaseClause {
	if trace {
		defer p.trace("caseClause")()
	}

	c := new(CaseClause)
	c.pos = p.pos()

	switch p.tok {
	case _Case:
		p.next(pstate)
		c.Cases = p.exprList(pstate)

	case _Default:
		p.next(pstate)

	default:
		p.syntaxError(pstate, "expecting case or default or }")
		p.advance(pstate, _Colon, _Case, _Default, _Rbrace)
	}

	c.Colon = p.pos()
	p.want(pstate, _Colon)
	c.Body = p.stmtList(pstate)

	return c
}

func (p *parser) commClause(pstate *PackageState) *CommClause {
	if trace {
		defer p.trace("commClause")()
	}

	c := new(CommClause)
	c.pos = p.pos()

	switch p.tok {
	case _Case:
		p.next(pstate)
		c.Comm = p.simpleStmt(pstate, nil, 0)

	// The syntax restricts the possible simple statements here to:
	//
	//     lhs <- x (send statement)
	//     <-x
	//     lhs = <-x
	//     lhs := <-x
	//
	// All these (and more) are recognized by simpleStmt and invalid
	// syntax trees are flagged later, during type checking.
	// TODO(gri) eventually may want to restrict valid syntax trees
	// here.

	case _Default:
		p.next(pstate)

	default:
		p.syntaxError(pstate, "expecting case or default or }")
		p.advance(pstate, _Colon, _Case, _Default, _Rbrace)
	}

	c.Colon = p.pos()
	p.want(pstate, _Colon)
	c.Body = p.stmtList(pstate)

	return c
}

// Statement =
// 	Declaration | LabeledStmt | SimpleStmt |
// 	GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
// 	FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
// 	DeferStmt .
func (p *parser) stmtOrNil(pstate *PackageState) Stmt {
	if trace {
		defer p.trace("stmt " + p.tok.String(pstate))()
	}

	// Most statements (assignments) start with an identifier;
	// look for it first before doing anything more expensive.
	if p.tok == _Name {
		lhs := p.exprList(pstate)
		if label, ok := lhs.(*Name); ok && p.tok == _Colon {
			return p.labeledStmtOrNil(pstate, label)
		}
		return p.simpleStmt(pstate, lhs, 0)
	}

	switch p.tok {
	case _Lbrace:
		return p.blockStmt(pstate, "")

	case _Var:
		return p.declStmt(pstate, p.varDecl)

	case _Const:
		return p.declStmt(pstate, p.constDecl)

	case _Type:
		return p.declStmt(pstate, p.typeDecl)

	case _Operator, _Star:
		switch p.op {
		case Add, Sub, Mul, And, Xor, Not:
			return p.simpleStmt(pstate, nil, 0) // unary operators
		}

	case _Literal, _Func, _Lparen, // operands
		_Lbrack, _Struct, _Map, _Chan, _Interface, // composite types
		_Arrow: // receive operator
		return p.simpleStmt(pstate, nil, 0)

	case _For:
		return p.forStmt(pstate)

	case _Switch:
		return p.switchStmt(pstate)

	case _Select:
		return p.selectStmt(pstate)

	case _If:
		return p.ifStmt(pstate)

	case _Fallthrough:
		s := new(BranchStmt)
		s.pos = p.pos()
		p.next(pstate)
		s.Tok = _Fallthrough
		return s

	case _Break, _Continue:
		s := new(BranchStmt)
		s.pos = p.pos()
		s.Tok = p.tok
		p.next(pstate)
		if p.tok == _Name {
			s.Label = p.name(pstate)
		}
		return s

	case _Go, _Defer:
		return p.callStmt(pstate)

	case _Goto:
		s := new(BranchStmt)
		s.pos = p.pos()
		s.Tok = _Goto
		p.next(pstate)
		s.Label = p.name(pstate)
		return s

	case _Return:
		s := new(ReturnStmt)
		s.pos = p.pos()
		p.next(pstate)
		if p.tok != _Semi && p.tok != _Rbrace {
			s.Results = p.exprList(pstate)
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
func (p *parser) stmtList(pstate *PackageState) (l []Stmt) {
	if trace {
		defer p.trace("stmtList")()
	}

	for p.tok != _EOF && p.tok != _Rbrace && p.tok != _Case && p.tok != _Default {
		s := p.stmtOrNil(pstate)
		if s == nil {
			break
		}
		l = append(l, s)
		// ";" is optional before "}"
		if !p.got(pstate, _Semi) && p.tok != _Rbrace {
			p.syntaxError(pstate, "at end of statement")
			p.advance(pstate, _Semi, _Rbrace, _Case, _Default)
			p.got(pstate, _Semi) // avoid spurious empty statement
		}
	}
	return
}

// Arguments = "(" [ ( ExpressionList | Type [ "," ExpressionList ] ) [ "..." ] [ "," ] ] ")" .
func (p *parser) argList(pstate *PackageState) (list []Expr, hasDots bool) {
	if trace {
		defer p.trace("argList")()
	}

	p.xnest++
	p.list(pstate, _Lparen, _Comma, _Rparen, func() bool {
		list = append(list, p.expr(pstate))
		hasDots = p.got(pstate, _DotDotDot)
		return hasDots
	})
	p.xnest--

	return
}

// ----------------------------------------------------------------------------
// Common productions

func (p *parser) newName(value string) *Name {
	n := new(Name)
	n.pos = p.pos()
	n.Value = value
	return n
}

func (p *parser) name(pstate *PackageState) *Name {
	// no tracing to avoid overly verbose output

	if p.tok == _Name {
		n := p.newName(p.lit)
		p.next(pstate)
		return n
	}

	n := p.newName("_")
	p.syntaxError(pstate, "expecting name")
	p.advance(pstate)
	return n
}

// IdentifierList = identifier { "," identifier } .
// The first name must be provided.
func (p *parser) nameList(pstate *PackageState, first *Name) []*Name {
	if trace {
		defer p.trace("nameList")()
	}

	if debug && first == nil {
		panic("first name not provided")
	}

	l := []*Name{first}
	for p.got(pstate, _Comma) {
		l = append(l, p.name(pstate))
	}

	return l
}

// The first name may be provided, or nil.
func (p *parser) qualifiedName(pstate *PackageState, name *Name) Expr {
	if trace {
		defer p.trace("qualifiedName")()
	}

	switch {
	case name != nil:
	// name is provided
	case p.tok == _Name:
		name = p.name(pstate)
	default:
		name = p.newName("_")
		p.syntaxError(pstate, "expecting name")
		p.advance(pstate, _Dot, _Semi, _Rbrace)
	}

	return p.dotname(pstate, name)
}

// ExpressionList = Expression { "," Expression } .
func (p *parser) exprList(pstate *PackageState) Expr {
	if trace {
		defer p.trace("exprList")()
	}

	x := p.expr(pstate)
	if p.got(pstate, _Comma) {
		list := []Expr{x, p.expr(pstate)}
		for p.got(pstate, _Comma) {
			list = append(list, p.expr(pstate))
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

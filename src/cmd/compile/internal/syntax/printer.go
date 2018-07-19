package syntax

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func (psess *PackageSession) Fprint(w io.Writer, x Node, linebreaks bool) (n int, err error) {
	p := printer{
		output:     w,
		linebreaks: linebreaks,
	}

	defer func() {
		n = p.written
		if e := recover(); e != nil {
			err = e.(localError).err
		}
	}()

	p.print(psess, x)
	p.flush(psess, _EOF)

	return
}

func (psess *PackageSession) String(n Node) string {
	var buf bytes.Buffer
	_, err := psess.Fprint(&buf, n, false)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

type ctrlSymbol int

const (
	none ctrlSymbol = iota
	semi
	blank
	newline
	indent
	outdent
)

type whitespace struct {
	last token
	kind ctrlSymbol
}

type printer struct {
	output     io.Writer
	written    int  // number of bytes written
	linebreaks bool // print linebreaks instead of semis

	indent  int // current indentation level
	nlcount int // number of consecutive newlines

	pending []whitespace // pending whitespace
	lastTok token        // last token (after any pending semi) processed by print
}

// write is a thin wrapper around p.output.Write
// that takes care of accounting and error handling.
func (p *printer) write(data []byte) {
	n, err := p.output.Write(data)
	p.written += n
	if err != nil {
		panic(localError{err})
	}
}

func (p *printer) writeBytes(psess *PackageSession, data []byte) {
	if len(data) == 0 {
		panic("expected non-empty []byte")
	}
	if p.nlcount > 0 && p.indent > 0 {

		n := p.indent
		for n > len(psess.tabBytes) {
			p.write(psess.tabBytes)
			n -= len(psess.tabBytes)
		}
		p.write(psess.tabBytes[:n])
	}
	p.write(data)
	p.nlcount = 0
}

func (p *printer) writeString(psess *PackageSession, s string) {
	p.writeBytes(psess, []byte(s))
}

// If impliesSemi returns true for a non-blank line's final token tok,
// a semicolon is automatically inserted. Vice versa, a semicolon may
// be omitted in those cases.
func impliesSemi(tok token) bool {
	switch tok {
	case _Name,
		_Break, _Continue, _Fallthrough, _Return,
		_Rparen, _Rbrack, _Rbrace:
		return true
	}
	return false
}

func lineComment(text string) bool {
	return strings.HasPrefix(text, "//")
}

func (p *printer) addWhitespace(kind ctrlSymbol, text string) {
	p.pending = append(p.pending, whitespace{p.lastTok, kind})
	switch kind {
	case semi:
		p.lastTok = _Semi
	case newline:
		p.lastTok = 0

	}
}

func (p *printer) flush(psess *PackageSession, next token) {

	sawNewline := next == _EOF
	sawParen := next == _Rparen || next == _Rbrace
	for i := len(p.pending) - 1; i >= 0; i-- {
		switch p.pending[i].kind {
		case semi:
			k := semi
			if sawParen {
				sawParen = false
				k = none
			} else if sawNewline && impliesSemi(p.pending[i].last) {
				sawNewline = false
				k = none
			}
			p.pending[i].kind = k
		case newline:
			sawNewline = true
		case blank, indent, outdent:

		default:
			panic("unreachable")
		}
	}

	prev := none
	for i := range p.pending {
		switch p.pending[i].kind {
		case none:

		case semi:
			p.writeString(psess, ";")
			p.nlcount = 0
			prev = semi
		case blank:
			if prev != blank {

				p.writeBytes(psess, psess.blankByte)
				p.nlcount = 0
				prev = blank
			}
		case newline:
			const maxEmptyLines = 1
			if p.nlcount <= maxEmptyLines {
				p.write(psess.newlineByte)
				p.nlcount++
				prev = newline
			}
		case indent:
			p.indent++
		case outdent:
			p.indent--
			if p.indent < 0 {
				panic("negative indentation")
			}

		default:
			panic("unreachable")
		}
	}

	p.pending = p.pending[:0]
}

func mayCombine(prev token, next byte) (b bool) {
	return

}

func (p *printer) print(psess *PackageSession, args ...interface{}) {
	for i := 0; i < len(args); i++ {
		switch x := args[i].(type) {
		case nil:

		case Node:
			p.printNode(psess, x)

		case token:
			// _Name implies an immediately following string
			// argument which is the actual value to print.
			var s string
			if x == _Name {
				i++
				if i >= len(args) {
					panic("missing string argument after _Name")
				}
				s = args[i].(string)
			} else {
				s = x.String(psess)
			}

			if mayCombine(p.lastTok, s[0]) {
				panic("adjacent tokens combine without whitespace")
			}

			if x == _Semi {

				p.addWhitespace(semi, "")
			} else {
				p.flush(psess, x)
				p.writeString(psess, s)
				p.nlcount = 0
				p.lastTok = x
			}

		case Operator:
			if x != 0 {
				p.flush(psess, _Operator)
				p.writeString(psess, x.String(psess))
			}

		case ctrlSymbol:
			switch x {
			case none, semi:
				panic("unreachable")
			case newline:

				if !p.linebreaks {
					x = blank
				}
			}
			p.addWhitespace(x, "")

		default:
			panic(fmt.Sprintf("unexpected argument %v (%T)", x, x))
		}
	}
}

func (p *printer) printNode(psess *PackageSession, n Node) {

	p.printRawNode(psess, n)

}

func (p *printer) printRawNode(psess *PackageSession, n Node) {
	switch n := n.(type) {
	case nil:

	case *BadExpr:
		p.print(psess, _Name, "<bad expr>")

	case *Name:
		p.print(psess, _Name, n.Value)

	case *BasicLit:
		p.print(psess, _Name, n.Value)

	case *FuncLit:
		p.print(psess, n.Type, blank, n.Body)

	case *CompositeLit:
		if n.Type != nil {
			p.print(psess, n.Type)
		}
		p.print(psess, _Lbrace)
		if n.NKeys > 0 && n.NKeys == len(n.ElemList) {
			p.printExprLines(psess, n.ElemList)
		} else {
			p.printExprList(psess, n.ElemList)
		}
		p.print(psess, _Rbrace)

	case *ParenExpr:
		p.print(psess, _Lparen, n.X, _Rparen)

	case *SelectorExpr:
		p.print(psess, n.X, _Dot, n.Sel)

	case *IndexExpr:
		p.print(psess, n.X, _Lbrack, n.Index, _Rbrack)

	case *SliceExpr:
		p.print(psess, n.X, _Lbrack)
		if i := n.Index[0]; i != nil {
			p.printNode(psess, i)
		}
		p.print(psess, _Colon)
		if j := n.Index[1]; j != nil {
			p.printNode(psess, j)
		}
		if k := n.Index[2]; k != nil {
			p.print(psess, _Colon, k)
		}
		p.print(psess, _Rbrack)

	case *AssertExpr:
		p.print(psess, n.X, _Dot, _Lparen, n.Type, _Rparen)

	case *TypeSwitchGuard:
		if n.Lhs != nil {
			p.print(psess, n.Lhs, blank, _Define, blank)
		}
		p.print(psess, n.X, _Dot, _Lparen, _Type, _Rparen)

	case *CallExpr:
		p.print(psess, n.Fun, _Lparen)
		p.printExprList(psess, n.ArgList)
		if n.HasDots {
			p.print(psess, _DotDotDot)
		}
		p.print(psess, _Rparen)

	case *Operation:
		if n.Y == nil {

			p.print(psess, n.Op)

			p.print(psess, n.X)
		} else {

			p.print(psess, n.X, blank, n.Op, blank, n.Y)
		}

	case *KeyValueExpr:
		p.print(psess, n.Key, _Colon, blank, n.Value)

	case *ListExpr:
		p.printExprList(psess, n.ElemList)

	case *ArrayType:
		var len interface{} = _DotDotDot
		if n.Len != nil {
			len = n.Len
		}
		p.print(psess, _Lbrack, len, _Rbrack, n.Elem)

	case *SliceType:
		p.print(psess, _Lbrack, _Rbrack, n.Elem)

	case *DotsType:
		p.print(psess, _DotDotDot, n.Elem)

	case *StructType:
		p.print(psess, _Struct)
		if len(n.FieldList) > 0 && p.linebreaks {
			p.print(psess, blank)
		}
		p.print(psess, _Lbrace)
		if len(n.FieldList) > 0 {
			p.print(psess, newline, indent)
			p.printFieldList(psess, n.FieldList, n.TagList)
			p.print(psess, outdent, newline)
		}
		p.print(psess, _Rbrace)

	case *FuncType:
		p.print(psess, _Func)
		p.printSignature(psess, n)

	case *InterfaceType:
		p.print(psess, _Interface)
		if len(n.MethodList) > 0 && p.linebreaks {
			p.print(psess, blank)
		}
		p.print(psess, _Lbrace)
		if len(n.MethodList) > 0 {
			p.print(psess, newline, indent)
			p.printMethodList(psess, n.MethodList)
			p.print(psess, outdent, newline)
		}
		p.print(psess, _Rbrace)

	case *MapType:
		p.print(psess, _Map, _Lbrack, n.Key, _Rbrack, n.Value)

	case *ChanType:
		if n.Dir == RecvOnly {
			p.print(psess, _Arrow)
		}
		p.print(psess, _Chan)
		if n.Dir == SendOnly {
			p.print(psess, _Arrow)
		}
		p.print(psess, blank, n.Elem)

	case *DeclStmt:
		p.printDecl(psess, n.DeclList)

	case *EmptyStmt:

	case *LabeledStmt:
		p.print(psess, outdent, n.Label, _Colon, indent, newline, n.Stmt)

	case *ExprStmt:
		p.print(psess, n.X)

	case *SendStmt:
		p.print(psess, n.Chan, blank, _Arrow, blank, n.Value)

	case *AssignStmt:
		p.print(psess, n.Lhs)
		if n.Rhs == psess.ImplicitOne {

			p.print(psess, n.Op, n.Op)
		} else {
			p.print(psess, blank, n.Op, _Assign, blank)
			p.print(psess, n.Rhs)
		}

	case *CallStmt:
		p.print(psess, n.Tok, blank, n.Call)

	case *ReturnStmt:
		p.print(psess, _Return)
		if n.Results != nil {
			p.print(psess, blank, n.Results)
		}

	case *BranchStmt:
		p.print(psess, n.Tok)
		if n.Label != nil {
			p.print(psess, blank, n.Label)
		}

	case *BlockStmt:
		p.print(psess, _Lbrace)
		if len(n.List) > 0 {
			p.print(psess, newline, indent)
			p.printStmtList(psess, n.List, true)
			p.print(psess, outdent, newline)
		}
		p.print(psess, _Rbrace)

	case *IfStmt:
		p.print(psess, _If, blank)
		if n.Init != nil {
			p.print(psess, n.Init, _Semi, blank)
		}
		p.print(psess, n.Cond, blank, n.Then)
		if n.Else != nil {
			p.print(psess, blank, _Else, blank, n.Else)
		}

	case *SwitchStmt:
		p.print(psess, _Switch, blank)
		if n.Init != nil {
			p.print(psess, n.Init, _Semi, blank)
		}
		if n.Tag != nil {
			p.print(psess, n.Tag, blank)
		}
		p.printSwitchBody(psess, n.Body)

	case *SelectStmt:
		p.print(psess, _Select, blank)
		p.printSelectBody(psess, n.Body)

	case *RangeClause:
		if n.Lhs != nil {
			tok := _Assign
			if n.Def {
				tok = _Define
			}
			p.print(psess, n.Lhs, blank, tok, blank)
		}
		p.print(psess, _Range, blank, n.X)

	case *ForStmt:
		p.print(psess, _For, blank)
		if n.Init == nil && n.Post == nil {
			if n.Cond != nil {
				p.print(psess, n.Cond, blank)
			}
		} else {
			if n.Init != nil {
				p.print(psess, n.Init)

				if _, ok := n.Init.(*RangeClause); ok {
					p.print(psess, blank, n.Body)
					break
				}
			}
			p.print(psess, _Semi, blank)
			if n.Cond != nil {
				p.print(psess, n.Cond)
			}
			p.print(psess, _Semi, blank)
			if n.Post != nil {
				p.print(psess, n.Post, blank)
			}
		}
		p.print(psess, n.Body)

	case *ImportDecl:
		if n.Group == nil {
			p.print(psess, _Import, blank)
		}
		if n.LocalPkgName != nil {
			p.print(psess, n.LocalPkgName, blank)
		}
		p.print(psess, n.Path)

	case *ConstDecl:
		if n.Group == nil {
			p.print(psess, _Const, blank)
		}
		p.printNameList(psess, n.NameList)
		if n.Type != nil {
			p.print(psess, blank, n.Type)
		}
		if n.Values != nil {
			p.print(psess, blank, _Assign, blank, n.Values)
		}

	case *TypeDecl:
		if n.Group == nil {
			p.print(psess, _Type, blank)
		}
		p.print(psess, n.Name, blank)
		if n.Alias {
			p.print(psess, _Assign, blank)
		}
		p.print(psess, n.Type)

	case *VarDecl:
		if n.Group == nil {
			p.print(psess, _Var, blank)
		}
		p.printNameList(psess, n.NameList)
		if n.Type != nil {
			p.print(psess, blank, n.Type)
		}
		if n.Values != nil {
			p.print(psess, blank, _Assign, blank, n.Values)
		}

	case *FuncDecl:
		p.print(psess, _Func, blank)
		if r := n.Recv; r != nil {
			p.print(psess, _Lparen)
			if r.Name != nil {
				p.print(psess, r.Name, blank)
			}
			p.printNode(psess, r.Type)
			p.print(psess, _Rparen, blank)
		}
		p.print(psess, n.Name)
		p.printSignature(psess, n.Type)
		if n.Body != nil {
			p.print(psess, blank, n.Body)
		}

	case *printGroup:
		p.print(psess, n.Tok, blank, _Lparen)
		if len(n.Decls) > 0 {
			p.print(psess, newline, indent)
			for _, d := range n.Decls {
				p.printNode(psess, d)
				p.print(psess, _Semi, newline)
			}
			p.print(psess, outdent)
		}
		p.print(psess, _Rparen)

	case *File:
		p.print(psess, _Package, blank, n.PkgName)
		if len(n.DeclList) > 0 {
			p.print(psess, _Semi, newline, newline)
			p.printDeclList(psess, n.DeclList)
		}

	default:
		panic(fmt.Sprintf("syntax.Iterate: unexpected node type %T", n))
	}
}

func (p *printer) printFields(psess *PackageSession, fields []*Field, tags []*BasicLit, i, j int) {
	if i+1 == j && fields[i].Name == nil {

		p.printNode(psess, fields[i].Type)
	} else {
		for k, f := range fields[i:j] {
			if k > 0 {
				p.print(psess, _Comma, blank)
			}
			p.printNode(psess, f.Name)
		}
		p.print(psess, blank)
		p.printNode(psess, fields[i].Type)
	}
	if i < len(tags) && tags[i] != nil {
		p.print(psess, blank)
		p.printNode(psess, tags[i])
	}
}

func (p *printer) printFieldList(psess *PackageSession, fields []*Field, tags []*BasicLit) {
	i0 := 0
	var typ Expr
	for i, f := range fields {
		if f.Name == nil || f.Type != typ {
			if i0 < i {
				p.printFields(psess, fields, tags, i0, i)
				p.print(psess, _Semi, newline)
				i0 = i
			}
			typ = f.Type
		}
	}
	p.printFields(psess, fields, tags, i0, len(fields))
}

func (p *printer) printMethodList(psess *PackageSession, methods []*Field) {
	for i, m := range methods {
		if i > 0 {
			p.print(psess, _Semi, newline)
		}
		if m.Name != nil {
			p.printNode(psess, m.Name)
			p.printSignature(psess, m.Type.(*FuncType))
		} else {
			p.printNode(psess, m.Type)
		}
	}
}

func (p *printer) printNameList(psess *PackageSession, list []*Name) {
	for i, x := range list {
		if i > 0 {
			p.print(psess, _Comma, blank)
		}
		p.printNode(psess, x)
	}
}

func (p *printer) printExprList(psess *PackageSession, list []Expr) {
	for i, x := range list {
		if i > 0 {
			p.print(psess, _Comma, blank)
		}
		p.printNode(psess, x)
	}
}

func (p *printer) printExprLines(psess *PackageSession, list []Expr) {
	if len(list) > 0 {
		p.print(psess, newline, indent)
		for _, x := range list {
			p.print(psess, x, _Comma, newline)
		}
		p.print(psess, outdent)
	}
}

func groupFor(d Decl) (token, *Group) {
	switch d := d.(type) {
	case *ImportDecl:
		return _Import, d.Group
	case *ConstDecl:
		return _Const, d.Group
	case *TypeDecl:
		return _Type, d.Group
	case *VarDecl:
		return _Var, d.Group
	case *FuncDecl:
		return _Func, nil
	default:
		panic("unreachable")
	}
}

type printGroup struct {
	node
	Tok   token
	Decls []Decl
}

func (p *printer) printDecl(psess *PackageSession, list []Decl) {
	tok, group := groupFor(list[0])

	if group == nil {
		if len(list) != 1 {
			panic("unreachable")
		}
		p.printNode(psess, list[0])
		return
	}

	// printGroup is here for consistent comment handling
	// (this is not yet used)
	var pg printGroup

	pg.Tok = tok
	pg.Decls = list
	p.printNode(psess, &pg)
}

func (p *printer) printDeclList(psess *PackageSession, list []Decl) {
	i0 := 0
	var tok token
	var group *Group
	for i, x := range list {
		if s, g := groupFor(x); g == nil || g != group {
			if i0 < i {
				p.printDecl(psess, list[i0:i])
				p.print(psess, _Semi, newline)

				if g != group || s != tok || s == _Func {
					p.print(psess, newline)
				}
				i0 = i
			}
			tok, group = s, g
		}
	}
	p.printDecl(psess, list[i0:])
}

func (p *printer) printSignature(psess *PackageSession, sig *FuncType) {
	p.printParameterList(psess, sig.ParamList)
	if list := sig.ResultList; list != nil {
		p.print(psess, blank)
		if len(list) == 1 && list[0].Name == nil {
			p.printNode(psess, list[0].Type)
		} else {
			p.printParameterList(psess, list)
		}
	}
}

func (p *printer) printParameterList(psess *PackageSession, list []*Field) {
	p.print(psess, _Lparen)
	if len(list) > 0 {
		for i, f := range list {
			if i > 0 {
				p.print(psess, _Comma, blank)
			}
			if f.Name != nil {
				p.printNode(psess, f.Name)
				if i+1 < len(list) {
					f1 := list[i+1]
					if f1.Name != nil && f1.Type == f.Type {
						continue
					}
				}
				p.print(psess, blank)
			}
			p.printNode(psess, f.Type)
		}
	}
	p.print(psess, _Rparen)
}

func (p *printer) printStmtList(psess *PackageSession, list []Stmt, braces bool) {
	for i, x := range list {
		p.print(psess, x, _Semi)
		if i+1 < len(list) {
			p.print(psess, newline)
		} else if braces {

			if _, ok := x.(*EmptyStmt); ok {
				p.print(psess, x, _Semi)
			}
		}
	}
}

func (p *printer) printSwitchBody(psess *PackageSession, list []*CaseClause) {
	p.print(psess, _Lbrace)
	if len(list) > 0 {
		p.print(psess, newline)
		for i, c := range list {
			p.printCaseClause(psess, c, i+1 == len(list))
			p.print(psess, newline)
		}
	}
	p.print(psess, _Rbrace)
}

func (p *printer) printSelectBody(psess *PackageSession, list []*CommClause) {
	p.print(psess, _Lbrace)
	if len(list) > 0 {
		p.print(psess, newline)
		for i, c := range list {
			p.printCommClause(psess, c, i+1 == len(list))
			p.print(psess, newline)
		}
	}
	p.print(psess, _Rbrace)
}

func (p *printer) printCaseClause(psess *PackageSession, c *CaseClause, braces bool) {
	if c.Cases != nil {
		p.print(psess, _Case, blank, c.Cases)
	} else {
		p.print(psess, _Default)
	}
	p.print(psess, _Colon)
	if len(c.Body) > 0 {
		p.print(psess, newline, indent)
		p.printStmtList(psess, c.Body, braces)
		p.print(psess, outdent)
	}
}

func (p *printer) printCommClause(psess *PackageSession, c *CommClause, braces bool) {
	if c.Comm != nil {
		p.print(psess, _Case, blank)
		p.print(psess, c.Comm)
	} else {
		p.print(psess, _Default)
	}
	p.print(psess, _Colon)
	if len(c.Body) > 0 {
		p.print(psess, newline, indent)
		p.printStmtList(psess, c.Body, braces)
		p.print(psess, outdent)
	}
}

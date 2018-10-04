// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements printing of syntax trees in source format.

package syntax

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// TODO(gri) Consider removing the linebreaks flag from this signature.
// Its likely rarely used in common cases.

func (pstate *PackageState) Fprint(w io.Writer, x Node, linebreaks bool) (n int, err error) {
	p := printer{
		output:     w,
		linebreaks: linebreaks,
	}

	defer func() {
		n = p.written
		if e := recover(); e != nil {
			err = e.(localError).err // re-panics if it's not a localError
		}
	}()

	p.print(pstate, x)
	p.flush(pstate, _EOF)

	return
}

func (pstate *PackageState) String(n Node) string {
	var buf bytes.Buffer
	_, err := pstate.Fprint(&buf, n, false)
	if err != nil {
		panic(err) // TODO(gri) print something sensible into buf instead
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

// comment
// eolComment
)

type whitespace struct {
	last token
	kind ctrlSymbol
	//text string // comment text (possibly ""); valid if kind == comment
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

func (p *printer) writeBytes(pstate *PackageState, data []byte) {
	if len(data) == 0 {
		panic("expected non-empty []byte")
	}
	if p.nlcount > 0 && p.indent > 0 {
		// write indentation
		n := p.indent
		for n > len(pstate.tabBytes) {
			p.write(pstate.tabBytes)
			n -= len(pstate.tabBytes)
		}
		p.write(pstate.tabBytes[:n])
	}
	p.write(data)
	p.nlcount = 0
}

func (p *printer) writeString(pstate *PackageState, s string) {
	p.writeBytes(pstate, []byte(s))
}

// If impliesSemi returns true for a non-blank line's final token tok,
// a semicolon is automatically inserted. Vice versa, a semicolon may
// be omitted in those cases.
func impliesSemi(tok token) bool {
	switch tok {
	case _Name,
		_Break, _Continue, _Fallthrough, _Return,
		/*_Inc, _Dec,*/ _Rparen, _Rbrack, _Rbrace: // TODO(gri) fix this
		return true
	}
	return false
}

// TODO(gri) provide table of []byte values for all tokens to avoid repeated string conversion

func lineComment(text string) bool {
	return strings.HasPrefix(text, "//")
}

func (p *printer) addWhitespace(kind ctrlSymbol, text string) {
	p.pending = append(p.pending, whitespace{p.lastTok, kind /*text*/})
	switch kind {
	case semi:
		p.lastTok = _Semi
	case newline:
		p.lastTok = 0
		// TODO(gri) do we need to handle /*-style comments containing newlines here?
	}
}

func (p *printer) flush(pstate *PackageState, next token) {
	// eliminate semis and redundant whitespace
	sawNewline := next == _EOF
	sawParen := next == _Rparen || next == _Rbrace
	for i := len(p.pending) - 1; i >= 0; i-- {
		switch p.pending[i].kind {
		case semi:
			k := semi
			if sawParen {
				sawParen = false
				k = none // eliminate semi
			} else if sawNewline && impliesSemi(p.pending[i].last) {
				sawNewline = false
				k = none // eliminate semi
			}
			p.pending[i].kind = k
		case newline:
			sawNewline = true
		case blank, indent, outdent:
		// nothing to do
		// case comment:
		// 	// A multi-line comment acts like a newline; and a ""
		// 	// comment implies by definition at least one newline.
		// 	if text := p.pending[i].text; strings.HasPrefix(text, "/*") && strings.ContainsRune(text, '\n') {
		// 		sawNewline = true
		// 	}
		// case eolComment:
		// 	// TODO(gri) act depending on sawNewline
		default:
			panic("unreachable")
		}
	}

	// print pending
	prev := none
	for i := range p.pending {
		switch p.pending[i].kind {
		case none:
		// nothing to do
		case semi:
			p.writeString(pstate, ";")
			p.nlcount = 0
			prev = semi
		case blank:
			if prev != blank {
				// at most one blank
				p.writeBytes(pstate, pstate.blankByte)
				p.nlcount = 0
				prev = blank
			}
		case newline:
			const maxEmptyLines = 1
			if p.nlcount <= maxEmptyLines {
				p.write(pstate.newlineByte)
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
		// case comment:
		// 	if text := p.pending[i].text; text != "" {
		// 		p.writeString(text)
		// 		p.nlcount = 0
		// 		prev = comment
		// 	}
		// 	// TODO(gri) should check that line comments are always followed by newline
		default:
			panic("unreachable")
		}
	}

	p.pending = p.pending[:0] // re-use underlying array
}

func mayCombine(prev token, next byte) (b bool) {
	return // for now
	// switch prev {
	// case lexical.Int:
	// 	b = next == '.' // 1.
	// case lexical.Add:
	// 	b = next == '+' // ++
	// case lexical.Sub:
	// 	b = next == '-' // --
	// case lexical.Quo:
	// 	b = next == '*' // /*
	// case lexical.Lss:
	// 	b = next == '-' || next == '<' // <- or <<
	// case lexical.And:
	// 	b = next == '&' || next == '^' // && or &^
	// }
	// return
}

func (p *printer) print(pstate *PackageState, args ...interface{}) {
	for i := 0; i < len(args); i++ {
		switch x := args[i].(type) {
		case nil:
		// we should not reach here but don't crash

		case Node:
			p.printNode(pstate, x)

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
				s = x.String(pstate)
			}

			// TODO(gri) This check seems at the wrong place since it doesn't
			//           take into account pending white space.
			if mayCombine(p.lastTok, s[0]) {
				panic("adjacent tokens combine without whitespace")
			}

			if x == _Semi {
				// delay printing of semi
				p.addWhitespace(semi, "")
			} else {
				p.flush(pstate, x)
				p.writeString(pstate, s)
				p.nlcount = 0
				p.lastTok = x
			}

		case Operator:
			if x != 0 {
				p.flush(pstate, _Operator)
				p.writeString(pstate, x.String(pstate))
			}

		case ctrlSymbol:
			switch x {
			case none, semi /*, comment*/ :
				panic("unreachable")
			case newline:
				// TODO(gri) need to handle mandatory newlines after a //-style comment
				if !p.linebreaks {
					x = blank
				}
			}
			p.addWhitespace(x, "")

		// case *Comment: // comments are not Nodes
		// 	p.addWhitespace(comment, x.Text)

		default:
			panic(fmt.Sprintf("unexpected argument %v (%T)", x, x))
		}
	}
}

func (p *printer) printNode(pstate *PackageState, n Node) {
	// ncom := *n.Comments()
	// if ncom != nil {
	// 	// TODO(gri) in general we cannot make assumptions about whether
	// 	// a comment is a /*- or a //-style comment since the syntax
	// 	// tree may have been manipulated. Need to make sure the correct
	// 	// whitespace is emitted.
	// 	for _, c := range ncom.Alone {
	// 		p.print(c, newline)
	// 	}
	// 	for _, c := range ncom.Before {
	// 		if c.Text == "" || lineComment(c.Text) {
	// 			panic("unexpected empty line or //-style 'before' comment")
	// 		}
	// 		p.print(c, blank)
	// 	}
	// }

	p.printRawNode(pstate, n)

	// if ncom != nil && len(ncom.After) > 0 {
	// 	for i, c := range ncom.After {
	// 		if i+1 < len(ncom.After) {
	// 			if c.Text == "" || lineComment(c.Text) {
	// 				panic("unexpected empty line or //-style non-final 'after' comment")
	// 			}
	// 		}
	// 		p.print(blank, c)
	// 	}
	// 	//p.print(newline)
	// }
}

func (p *printer) printRawNode(pstate *PackageState, n Node) {
	switch n := n.(type) {
	case nil:
	// we should not reach here but don't crash

	// expressions and types
	case *BadExpr:
		p.print(pstate, _Name, "<bad expr>")

	case *Name:
		p.print(pstate, _Name, n.Value) // _Name requires actual value following immediately

	case *BasicLit:
		p.print(pstate, _Name, n.Value) // _Name requires actual value following immediately

	case *FuncLit:
		p.print(pstate, n.Type, blank, n.Body)

	case *CompositeLit:
		if n.Type != nil {
			p.print(pstate, n.Type)
		}
		p.print(pstate, _Lbrace)
		if n.NKeys > 0 && n.NKeys == len(n.ElemList) {
			p.printExprLines(pstate, n.ElemList)
		} else {
			p.printExprList(pstate, n.ElemList)
		}
		p.print(pstate, _Rbrace)

	case *ParenExpr:
		p.print(pstate, _Lparen, n.X, _Rparen)

	case *SelectorExpr:
		p.print(pstate, n.X, _Dot, n.Sel)

	case *IndexExpr:
		p.print(pstate, n.X, _Lbrack, n.Index, _Rbrack)

	case *SliceExpr:
		p.print(pstate, n.X, _Lbrack)
		if i := n.Index[0]; i != nil {
			p.printNode(pstate, i)
		}
		p.print(pstate, _Colon)
		if j := n.Index[1]; j != nil {
			p.printNode(pstate, j)
		}
		if k := n.Index[2]; k != nil {
			p.print(pstate, _Colon, k)
		}
		p.print(pstate, _Rbrack)

	case *AssertExpr:
		p.print(pstate, n.X, _Dot, _Lparen, n.Type, _Rparen)

	case *TypeSwitchGuard:
		if n.Lhs != nil {
			p.print(pstate, n.Lhs, blank, _Define, blank)
		}
		p.print(pstate, n.X, _Dot, _Lparen, _Type, _Rparen)

	case *CallExpr:
		p.print(pstate, n.Fun, _Lparen)
		p.printExprList(pstate, n.ArgList)
		if n.HasDots {
			p.print(pstate, _DotDotDot)
		}
		p.print(pstate, _Rparen)

	case *Operation:
		if n.Y == nil {
			// unary expr
			p.print(pstate, n.Op)
			// if n.Op == lexical.Range {
			// 	p.print(blank)
			// }
			p.print(pstate, n.X)
		} else {
			// binary expr
			// TODO(gri) eventually take precedence into account
			// to control possibly missing parentheses
			p.print(pstate, n.X, blank, n.Op, blank, n.Y)
		}

	case *KeyValueExpr:
		p.print(pstate, n.Key, _Colon, blank, n.Value)

	case *ListExpr:
		p.printExprList(pstate, n.ElemList)

	case *ArrayType:
		var len interface{} = _DotDotDot
		if n.Len != nil {
			len = n.Len
		}
		p.print(pstate, _Lbrack, len, _Rbrack, n.Elem)

	case *SliceType:
		p.print(pstate, _Lbrack, _Rbrack, n.Elem)

	case *DotsType:
		p.print(pstate, _DotDotDot, n.Elem)

	case *StructType:
		p.print(pstate, _Struct)
		if len(n.FieldList) > 0 && p.linebreaks {
			p.print(pstate, blank)
		}
		p.print(pstate, _Lbrace)
		if len(n.FieldList) > 0 {
			p.print(pstate, newline, indent)
			p.printFieldList(pstate, n.FieldList, n.TagList)
			p.print(pstate, outdent, newline)
		}
		p.print(pstate, _Rbrace)

	case *FuncType:
		p.print(pstate, _Func)
		p.printSignature(pstate, n)

	case *InterfaceType:
		p.print(pstate, _Interface)
		if len(n.MethodList) > 0 && p.linebreaks {
			p.print(pstate, blank)
		}
		p.print(pstate, _Lbrace)
		if len(n.MethodList) > 0 {
			p.print(pstate, newline, indent)
			p.printMethodList(pstate, n.MethodList)
			p.print(pstate, outdent, newline)
		}
		p.print(pstate, _Rbrace)

	case *MapType:
		p.print(pstate, _Map, _Lbrack, n.Key, _Rbrack, n.Value)

	case *ChanType:
		if n.Dir == RecvOnly {
			p.print(pstate, _Arrow)
		}
		p.print(pstate, _Chan)
		if n.Dir == SendOnly {
			p.print(pstate, _Arrow)
		}
		p.print(pstate, blank, n.Elem)

	// statements
	case *DeclStmt:
		p.printDecl(pstate, n.DeclList)

	case *EmptyStmt:
	// nothing to print

	case *LabeledStmt:
		p.print(pstate, outdent, n.Label, _Colon, indent, newline, n.Stmt)

	case *ExprStmt:
		p.print(pstate, n.X)

	case *SendStmt:
		p.print(pstate, n.Chan, blank, _Arrow, blank, n.Value)

	case *AssignStmt:
		p.print(pstate, n.Lhs)
		if n.Rhs == pstate.ImplicitOne {
			// TODO(gri) This is going to break the mayCombine
			//           check once we enable that again.
			p.print(pstate, n.Op, n.Op) // ++ or --
		} else {
			p.print(pstate, blank, n.Op, _Assign, blank)
			p.print(pstate, n.Rhs)
		}

	case *CallStmt:
		p.print(pstate, n.Tok, blank, n.Call)

	case *ReturnStmt:
		p.print(pstate, _Return)
		if n.Results != nil {
			p.print(pstate, blank, n.Results)
		}

	case *BranchStmt:
		p.print(pstate, n.Tok)
		if n.Label != nil {
			p.print(pstate, blank, n.Label)
		}

	case *BlockStmt:
		p.print(pstate, _Lbrace)
		if len(n.List) > 0 {
			p.print(pstate, newline, indent)
			p.printStmtList(pstate, n.List, true)
			p.print(pstate, outdent, newline)
		}
		p.print(pstate, _Rbrace)

	case *IfStmt:
		p.print(pstate, _If, blank)
		if n.Init != nil {
			p.print(pstate, n.Init, _Semi, blank)
		}
		p.print(pstate, n.Cond, blank, n.Then)
		if n.Else != nil {
			p.print(pstate, blank, _Else, blank, n.Else)
		}

	case *SwitchStmt:
		p.print(pstate, _Switch, blank)
		if n.Init != nil {
			p.print(pstate, n.Init, _Semi, blank)
		}
		if n.Tag != nil {
			p.print(pstate, n.Tag, blank)
		}
		p.printSwitchBody(pstate, n.Body)

	case *SelectStmt:
		p.print(pstate, _Select, blank) // for now
		p.printSelectBody(pstate, n.Body)

	case *RangeClause:
		if n.Lhs != nil {
			tok := _Assign
			if n.Def {
				tok = _Define
			}
			p.print(pstate, n.Lhs, blank, tok, blank)
		}
		p.print(pstate, _Range, blank, n.X)

	case *ForStmt:
		p.print(pstate, _For, blank)
		if n.Init == nil && n.Post == nil {
			if n.Cond != nil {
				p.print(pstate, n.Cond, blank)
			}
		} else {
			if n.Init != nil {
				p.print(pstate, n.Init)
				// TODO(gri) clean this up
				if _, ok := n.Init.(*RangeClause); ok {
					p.print(pstate, blank, n.Body)
					break
				}
			}
			p.print(pstate, _Semi, blank)
			if n.Cond != nil {
				p.print(pstate, n.Cond)
			}
			p.print(pstate, _Semi, blank)
			if n.Post != nil {
				p.print(pstate, n.Post, blank)
			}
		}
		p.print(pstate, n.Body)

	case *ImportDecl:
		if n.Group == nil {
			p.print(pstate, _Import, blank)
		}
		if n.LocalPkgName != nil {
			p.print(pstate, n.LocalPkgName, blank)
		}
		p.print(pstate, n.Path)

	case *ConstDecl:
		if n.Group == nil {
			p.print(pstate, _Const, blank)
		}
		p.printNameList(pstate, n.NameList)
		if n.Type != nil {
			p.print(pstate, blank, n.Type)
		}
		if n.Values != nil {
			p.print(pstate, blank, _Assign, blank, n.Values)
		}

	case *TypeDecl:
		if n.Group == nil {
			p.print(pstate, _Type, blank)
		}
		p.print(pstate, n.Name, blank)
		if n.Alias {
			p.print(pstate, _Assign, blank)
		}
		p.print(pstate, n.Type)

	case *VarDecl:
		if n.Group == nil {
			p.print(pstate, _Var, blank)
		}
		p.printNameList(pstate, n.NameList)
		if n.Type != nil {
			p.print(pstate, blank, n.Type)
		}
		if n.Values != nil {
			p.print(pstate, blank, _Assign, blank, n.Values)
		}

	case *FuncDecl:
		p.print(pstate, _Func, blank)
		if r := n.Recv; r != nil {
			p.print(pstate, _Lparen)
			if r.Name != nil {
				p.print(pstate, r.Name, blank)
			}
			p.printNode(pstate, r.Type)
			p.print(pstate, _Rparen, blank)
		}
		p.print(pstate, n.Name)
		p.printSignature(pstate, n.Type)
		if n.Body != nil {
			p.print(pstate, blank, n.Body)
		}

	case *printGroup:
		p.print(pstate, n.Tok, blank, _Lparen)
		if len(n.Decls) > 0 {
			p.print(pstate, newline, indent)
			for _, d := range n.Decls {
				p.printNode(pstate, d)
				p.print(pstate, _Semi, newline)
			}
			p.print(pstate, outdent)
		}
		p.print(pstate, _Rparen)

	// files
	case *File:
		p.print(pstate, _Package, blank, n.PkgName)
		if len(n.DeclList) > 0 {
			p.print(pstate, _Semi, newline, newline)
			p.printDeclList(pstate, n.DeclList)
		}

	default:
		panic(fmt.Sprintf("syntax.Iterate: unexpected node type %T", n))
	}
}

func (p *printer) printFields(pstate *PackageState, fields []*Field, tags []*BasicLit, i, j int) {
	if i+1 == j && fields[i].Name == nil {
		// anonymous field
		p.printNode(pstate, fields[i].Type)
	} else {
		for k, f := range fields[i:j] {
			if k > 0 {
				p.print(pstate, _Comma, blank)
			}
			p.printNode(pstate, f.Name)
		}
		p.print(pstate, blank)
		p.printNode(pstate, fields[i].Type)
	}
	if i < len(tags) && tags[i] != nil {
		p.print(pstate, blank)
		p.printNode(pstate, tags[i])
	}
}

func (p *printer) printFieldList(pstate *PackageState, fields []*Field, tags []*BasicLit) {
	i0 := 0
	var typ Expr
	for i, f := range fields {
		if f.Name == nil || f.Type != typ {
			if i0 < i {
				p.printFields(pstate, fields, tags, i0, i)
				p.print(pstate, _Semi, newline)
				i0 = i
			}
			typ = f.Type
		}
	}
	p.printFields(pstate, fields, tags, i0, len(fields))
}

func (p *printer) printMethodList(pstate *PackageState, methods []*Field) {
	for i, m := range methods {
		if i > 0 {
			p.print(pstate, _Semi, newline)
		}
		if m.Name != nil {
			p.printNode(pstate, m.Name)
			p.printSignature(pstate, m.Type.(*FuncType))
		} else {
			p.printNode(pstate, m.Type)
		}
	}
}

func (p *printer) printNameList(pstate *PackageState, list []*Name) {
	for i, x := range list {
		if i > 0 {
			p.print(pstate, _Comma, blank)
		}
		p.printNode(pstate, x)
	}
}

func (p *printer) printExprList(pstate *PackageState, list []Expr) {
	for i, x := range list {
		if i > 0 {
			p.print(pstate, _Comma, blank)
		}
		p.printNode(pstate, x)
	}
}

func (p *printer) printExprLines(pstate *PackageState, list []Expr) {
	if len(list) > 0 {
		p.print(pstate, newline, indent)
		for _, x := range list {
			p.print(pstate, x, _Comma, newline)
		}
		p.print(pstate, outdent)
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

func (p *printer) printDecl(pstate *PackageState, list []Decl) {
	tok, group := groupFor(list[0])

	if group == nil {
		if len(list) != 1 {
			panic("unreachable")
		}
		p.printNode(pstate, list[0])
		return
	}

	// if _, ok := list[0].(*EmptyDecl); ok {
	// 	if len(list) != 1 {
	// 		panic("unreachable")
	// 	}
	// 	// TODO(gri) if there are comments inside the empty
	// 	// group, we may need to keep the list non-nil
	// 	list = nil
	// }

	// printGroup is here for consistent comment handling
	// (this is not yet used)
	var pg printGroup
	// *pg.Comments() = *group.Comments()
	pg.Tok = tok
	pg.Decls = list
	p.printNode(pstate, &pg)
}

func (p *printer) printDeclList(pstate *PackageState, list []Decl) {
	i0 := 0
	var tok token
	var group *Group
	for i, x := range list {
		if s, g := groupFor(x); g == nil || g != group {
			if i0 < i {
				p.printDecl(pstate, list[i0:i])
				p.print(pstate, _Semi, newline)
				// print empty line between different declaration groups,
				// different kinds of declarations, or between functions
				if g != group || s != tok || s == _Func {
					p.print(pstate, newline)
				}
				i0 = i
			}
			tok, group = s, g
		}
	}
	p.printDecl(pstate, list[i0:])
}

func (p *printer) printSignature(pstate *PackageState, sig *FuncType) {
	p.printParameterList(pstate, sig.ParamList)
	if list := sig.ResultList; list != nil {
		p.print(pstate, blank)
		if len(list) == 1 && list[0].Name == nil {
			p.printNode(pstate, list[0].Type)
		} else {
			p.printParameterList(pstate, list)
		}
	}
}

func (p *printer) printParameterList(pstate *PackageState, list []*Field) {
	p.print(pstate, _Lparen)
	if len(list) > 0 {
		for i, f := range list {
			if i > 0 {
				p.print(pstate, _Comma, blank)
			}
			if f.Name != nil {
				p.printNode(pstate, f.Name)
				if i+1 < len(list) {
					f1 := list[i+1]
					if f1.Name != nil && f1.Type == f.Type {
						continue // no need to print type
					}
				}
				p.print(pstate, blank)
			}
			p.printNode(pstate, f.Type)
		}
	}
	p.print(pstate, _Rparen)
}

func (p *printer) printStmtList(pstate *PackageState, list []Stmt, braces bool) {
	for i, x := range list {
		p.print(pstate, x, _Semi)
		if i+1 < len(list) {
			p.print(pstate, newline)
		} else if braces {
			// Print an extra semicolon if the last statement is
			// an empty statement and we are in a braced block
			// because one semicolon is automatically removed.
			if _, ok := x.(*EmptyStmt); ok {
				p.print(pstate, x, _Semi)
			}
		}
	}
}

func (p *printer) printSwitchBody(pstate *PackageState, list []*CaseClause) {
	p.print(pstate, _Lbrace)
	if len(list) > 0 {
		p.print(pstate, newline)
		for i, c := range list {
			p.printCaseClause(pstate, c, i+1 == len(list))
			p.print(pstate, newline)
		}
	}
	p.print(pstate, _Rbrace)
}

func (p *printer) printSelectBody(pstate *PackageState, list []*CommClause) {
	p.print(pstate, _Lbrace)
	if len(list) > 0 {
		p.print(pstate, newline)
		for i, c := range list {
			p.printCommClause(pstate, c, i+1 == len(list))
			p.print(pstate, newline)
		}
	}
	p.print(pstate, _Rbrace)
}

func (p *printer) printCaseClause(pstate *PackageState, c *CaseClause, braces bool) {
	if c.Cases != nil {
		p.print(pstate, _Case, blank, c.Cases)
	} else {
		p.print(pstate, _Default)
	}
	p.print(pstate, _Colon)
	if len(c.Body) > 0 {
		p.print(pstate, newline, indent)
		p.printStmtList(pstate, c.Body, braces)
		p.print(pstate, outdent)
	}
}

func (p *printer) printCommClause(pstate *PackageState, c *CommClause, braces bool) {
	if c.Comm != nil {
		p.print(pstate, _Case, blank)
		p.print(pstate, c.Comm)
	} else {
		p.print(pstate, _Default)
	}
	p.print(pstate, _Colon)
	if len(c.Body) > 0 {
		p.print(pstate, newline, indent)
		p.printStmtList(pstate, c.Body, braces)
		p.print(pstate, outdent)
	}
}

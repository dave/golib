// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"strconv"
	"strings"
	"unicode/utf8"
)

// A FmtFlag value is a set of flags (or 0).
// They control how the Xconv functions format their values.
// See the respective function's documentation for details.
type FmtFlag int

const ( //                                 fmt.Format flag/prec or verb
	FmtLeft     FmtFlag = 1 << iota // '-'
	FmtSharp                        // '#'
	FmtSign                         // '+'
	FmtUnsigned                     // internal use only (historic: u flag)
	FmtShort                        // verb == 'S'       (historic: h flag)
	FmtLong                         // verb == 'L'       (historic: l flag)
	FmtComma                        // '.' (== hasPrec)  (historic: , flag)
	FmtByte                         // '0'               (historic: hh flag)
)

// fmtFlag computes the (internal) FmtFlag
// value given the fmt.State and format verb.
func (pstate *PackageState) fmtFlag(s fmt.State, verb rune) FmtFlag {
	var flag FmtFlag
	if s.Flag('-') {
		flag |= FmtLeft
	}
	if s.Flag('#') {
		flag |= FmtSharp
	}
	if s.Flag('+') {
		flag |= FmtSign
	}
	if s.Flag(' ') {
		pstate.Fatalf("FmtUnsigned in format string")
	}
	if _, ok := s.Precision(); ok {
		flag |= FmtComma
	}
	if s.Flag('0') {
		flag |= FmtByte
	}
	switch verb {
	case 'S':
		flag |= FmtShort
	case 'L':
		flag |= FmtLong
	}
	return flag
}

// Format conversions:
// TODO(gri) verify these; eliminate those not used anymore
//
//	%v Op		Node opcodes
//		Flags:  #: print Go syntax (automatic unless mode == FDbg)
//
//	%j *Node	Node details
//		Flags:  0: suppresses things not relevant until walk
//
//	%v *Val		Constant values
//
//	%v *types.Sym		Symbols
//	%S              unqualified identifier in any mode
//		Flags:  +,- #: mode (see below)
//			0: in export mode: unqualified identifier if exported, qualified if not
//
//	%v *types.Type	Types
//	%S              omit "func" and receiver in function types
//	%L              definition instead of name.
//		Flags:  +,- #: mode (see below)
//			' ' (only in -/Sym mode) print type identifiers wit package name instead of prefix.
//
//	%v *Node	Nodes
//	%S              (only in +/debug mode) suppress recursion
//	%L              (only in Error mode) print "foo (type Bar)"
//		Flags:  +,- #: mode (see below)
//
//	%v Nodes	Node lists
//		Flags:  those of *Node
//			.: separate items with ',' instead of ';'

// *types.Sym, *types.Type, and *Node types use the flags below to set the format mode
const (
	FErr = iota
	FDbg
	FTypeId
	FTypeIdName // same as FTypeId, but use package name instead of prefix
)

// The mode flags '+', '-', and '#' are sticky; they persist through
// recursions of *Node, *types.Type, and *types.Sym values. The ' ' flag is
// sticky only on *types.Type recursions and only used in %-/*types.Sym mode.
//
// Example: given a *types.Sym: %+v %#v %-v print an identifier properly qualified for debug/export/internal mode

// Useful format combinations:
// TODO(gri): verify these
//
// *Node, Nodes:
//   %+v    multiline recursive debug dump of *Node/Nodes
//   %+S    non-recursive debug dump
//
// *Node:
//   %#v    Go format
//   %L     "foo (type Bar)" for error messages
//
// *types.Type:
//   %#v    Go format
//   %#L    type definition instead of name
//   %#S    omit"func" and receiver in function signature
//
//   %-v    type identifiers
//   %-S    type identifiers without "func" and arg names in type signatures (methodsym)
//   %- v   type identifiers with package name instead of prefix (typesym, dcommontype, typehash)

// update returns the results of applying f to mode.
func (f FmtFlag) update(mode fmtMode) (FmtFlag, fmtMode) {
	switch {
	case f&FmtSign != 0:
		mode = FDbg
	case f&FmtSharp != 0:
	// ignore (textual export format no longer supported)
	case f&FmtUnsigned != 0:
		mode = FTypeIdName
	case f&FmtLeft != 0:
		mode = FTypeId
	}

	f &^= FmtSharp | FmtLeft | FmtSign
	return f, mode
}

func (o Op) GoString() string {
	return fmt.Sprintf("%#v", o)
}

func (o Op) format(pstate *PackageState, s fmt.State, verb rune, mode fmtMode) {
	switch verb {
	case 'v':
		o.oconv(pstate, s, pstate.fmtFlag(s, verb), mode)

	default:
		fmt.Fprintf(s, "%%!%c(Op=%d)", verb, int(o))
	}
}

func (o Op) oconv(pstate *PackageState, s fmt.State, flag FmtFlag, mode fmtMode) {
	if flag&FmtSharp != 0 || mode != FDbg {
		if int(o) < len(pstate.goopnames) && pstate.goopnames[o] != "" {
			fmt.Fprint(s, pstate.goopnames[o])
			return
		}
	}

	// 'o.String()' instead of just 'o' to avoid infinite recursion
	fmt.Fprint(s, o.String(pstate))
}

type (
	fmtMode int

	fmtNodeErr        Node
	fmtNodeDbg        Node
	fmtNodeTypeId     Node
	fmtNodeTypeIdName Node

	fmtOpErr        Op
	fmtOpDbg        Op
	fmtOpTypeId     Op
	fmtOpTypeIdName Op

	fmtTypeErr        types.Type
	fmtTypeDbg        types.Type
	fmtTypeTypeId     types.Type
	fmtTypeTypeIdName types.Type

	fmtSymErr        types.Sym
	fmtSymDbg        types.Sym
	fmtSymTypeId     types.Sym
	fmtSymTypeIdName types.Sym

	fmtNodesErr        Nodes
	fmtNodesDbg        Nodes
	fmtNodesTypeId     Nodes
	fmtNodesTypeIdName Nodes
)

func (n *fmtNodeErr) Format(pstate *PackageState, s fmt.State, verb rune) {
	(*Node)(n).format(pstate, s, verb, FErr)
}
func (n *fmtNodeDbg) Format(pstate *PackageState, s fmt.State, verb rune) {
	(*Node)(n).format(pstate, s, verb, FDbg)
}
func (n *fmtNodeTypeId) Format(pstate *PackageState, s fmt.State, verb rune) {
	(*Node)(n).format(pstate, s, verb, FTypeId)
}
func (n *fmtNodeTypeIdName) Format(pstate *PackageState, s fmt.State, verb rune) {
	(*Node)(n).format(pstate, s, verb, FTypeIdName)
}
func (n *Node) Format(pstate *PackageState, s fmt.State, verb rune) { n.format(pstate, s, verb, FErr) }

func (o fmtOpErr) Format(pstate *PackageState, s fmt.State, verb rune) {
	Op(o).format(pstate, s, verb, FErr)
}
func (o fmtOpDbg) Format(pstate *PackageState, s fmt.State, verb rune) {
	Op(o).format(pstate, s, verb, FDbg)
}
func (o fmtOpTypeId) Format(pstate *PackageState, s fmt.State, verb rune) {
	Op(o).format(pstate, s, verb, FTypeId)
}
func (o fmtOpTypeIdName) Format(pstate *PackageState, s fmt.State, verb rune) {
	Op(o).format(pstate, s, verb, FTypeIdName)
}
func (o Op) Format(pstate *PackageState, s fmt.State, verb rune) { o.format(pstate, s, verb, FErr) }

func (t *fmtTypeErr) Format(pstate *PackageState, s fmt.State, verb rune) {
	pstate.typeFormat((*types.Type)(t), s, verb, FErr)
}
func (t *fmtTypeDbg) Format(pstate *PackageState, s fmt.State, verb rune) {
	pstate.typeFormat((*types.Type)(t), s, verb, FDbg)
}
func (t *fmtTypeTypeId) Format(pstate *PackageState, s fmt.State, verb rune) {
	pstate.typeFormat((*types.Type)(t), s, verb, FTypeId)
}
func (t *fmtTypeTypeIdName) Format(pstate *PackageState, s fmt.State, verb rune) {
	pstate.typeFormat((*types.Type)(t), s, verb, FTypeIdName)
}

// func (t *types.Type) Format(s fmt.State, verb rune)     // in package types

func (y *fmtSymErr) Format(pstate *PackageState, s fmt.State, verb rune) {
	pstate.symFormat((*types.Sym)(y), s, verb, FErr)
}
func (y *fmtSymDbg) Format(pstate *PackageState, s fmt.State, verb rune) {
	pstate.symFormat((*types.Sym)(y), s, verb, FDbg)
}
func (y *fmtSymTypeId) Format(pstate *PackageState, s fmt.State, verb rune) {
	pstate.symFormat((*types.Sym)(y), s, verb, FTypeId)
}
func (y *fmtSymTypeIdName) Format(pstate *PackageState, s fmt.State, verb rune) {
	pstate.symFormat((*types.Sym)(y), s, verb, FTypeIdName)
}

// func (y *types.Sym) Format(s fmt.State, verb rune)            // in package types  { y.format(s, verb, FErr) }

func (n fmtNodesErr) Format(pstate *PackageState, s fmt.State, verb rune) {
	(Nodes)(n).format(pstate, s, verb, FErr)
}
func (n fmtNodesDbg) Format(pstate *PackageState, s fmt.State, verb rune) {
	(Nodes)(n).format(pstate, s, verb, FDbg)
}
func (n fmtNodesTypeId) Format(pstate *PackageState, s fmt.State, verb rune) {
	(Nodes)(n).format(pstate, s, verb, FTypeId)
}
func (n fmtNodesTypeIdName) Format(pstate *PackageState, s fmt.State, verb rune) {
	(Nodes)(n).format(pstate, s, verb, FTypeIdName)
}
func (n Nodes) Format(pstate *PackageState, s fmt.State, verb rune) { n.format(pstate, s, verb, FErr) }

func (m fmtMode) Fprintf(pstate *PackageState, s fmt.State, format string, args ...interface{}) {
	m.prepareArgs(pstate, args)
	fmt.Fprintf(s, format, args...)
}

func (m fmtMode) Sprintf(pstate *PackageState, format string, args ...interface{}) string {
	m.prepareArgs(pstate, args)
	return fmt.Sprintf(format, args...)
}

func (m fmtMode) Sprint(pstate *PackageState, args ...interface{}) string {
	m.prepareArgs(pstate, args)
	return fmt.Sprint(args...)
}

func (m fmtMode) prepareArgs(pstate *PackageState, args []interface{}) {
	switch m {
	case FErr:
		for i, arg := range args {
			switch arg := arg.(type) {
			case Op:
				args[i] = fmtOpErr(arg)
			case *Node:
				args[i] = (*fmtNodeErr)(arg)
			case *types.Type:
				args[i] = (*fmtTypeErr)(arg)
			case *types.Sym:
				args[i] = (*fmtSymErr)(arg)
			case Nodes:
				args[i] = fmtNodesErr(arg)
			case Val, int32, int64, string, types.EType:
			// OK: printing these types doesn't depend on mode
			default:
				pstate.Fatalf("mode.prepareArgs type %T", arg)
			}
		}
	case FDbg:
		for i, arg := range args {
			switch arg := arg.(type) {
			case Op:
				args[i] = fmtOpDbg(arg)
			case *Node:
				args[i] = (*fmtNodeDbg)(arg)
			case *types.Type:
				args[i] = (*fmtTypeDbg)(arg)
			case *types.Sym:
				args[i] = (*fmtSymDbg)(arg)
			case Nodes:
				args[i] = fmtNodesDbg(arg)
			case Val, int32, int64, string, types.EType:
			// OK: printing these types doesn't depend on mode
			default:
				pstate.Fatalf("mode.prepareArgs type %T", arg)
			}
		}
	case FTypeId:
		for i, arg := range args {
			switch arg := arg.(type) {
			case Op:
				args[i] = fmtOpTypeId(arg)
			case *Node:
				args[i] = (*fmtNodeTypeId)(arg)
			case *types.Type:
				args[i] = (*fmtTypeTypeId)(arg)
			case *types.Sym:
				args[i] = (*fmtSymTypeId)(arg)
			case Nodes:
				args[i] = fmtNodesTypeId(arg)
			case Val, int32, int64, string, types.EType:
			// OK: printing these types doesn't depend on mode
			default:
				pstate.Fatalf("mode.prepareArgs type %T", arg)
			}
		}
	case FTypeIdName:
		for i, arg := range args {
			switch arg := arg.(type) {
			case Op:
				args[i] = fmtOpTypeIdName(arg)
			case *Node:
				args[i] = (*fmtNodeTypeIdName)(arg)
			case *types.Type:
				args[i] = (*fmtTypeTypeIdName)(arg)
			case *types.Sym:
				args[i] = (*fmtSymTypeIdName)(arg)
			case Nodes:
				args[i] = fmtNodesTypeIdName(arg)
			case Val, int32, int64, string, types.EType:
			// OK: printing these types doesn't depend on mode
			default:
				pstate.Fatalf("mode.prepareArgs type %T", arg)
			}
		}
	default:
		pstate.Fatalf("mode.prepareArgs mode %d", m)
	}
}

func (n *Node) format(pstate *PackageState, s fmt.State, verb rune, mode fmtMode) {
	switch verb {
	case 'v', 'S', 'L':
		n.nconv(pstate, s, pstate.fmtFlag(s, verb), mode)

	case 'j':
		n.jconv(s, pstate.fmtFlag(s, verb))

	default:
		fmt.Fprintf(s, "%%!%c(*Node=%p)", verb, n)
	}
}

// *Node details
func (n *Node) jconv(s fmt.State, flag FmtFlag) {
	c := flag & FmtShort

	if c == 0 && n.Addable() {
		fmt.Fprintf(s, " a(%v)", n.Addable())
	}

	if c == 0 && n.Name != nil && n.Name.Vargen != 0 {
		fmt.Fprintf(s, " g(%d)", n.Name.Vargen)
	}

	if n.Pos.IsKnown() {
		fmt.Fprintf(s, " l(%d)", n.Pos.Line())
	}

	if c == 0 && n.Xoffset != BADWIDTH {
		fmt.Fprintf(s, " x(%d)", n.Xoffset)
	}

	if n.Class() != 0 {
		fmt.Fprintf(s, " class(%v)", n.Class())
	}

	if n.Colas() {
		fmt.Fprintf(s, " colas(%v)", n.Colas())
	}

	switch n.Esc {
	case EscUnknown:
		break

	case EscHeap:
		fmt.Fprint(s, " esc(h)")

	case EscNone:
		fmt.Fprint(s, " esc(no)")

	case EscNever:
		if c == 0 {
			fmt.Fprint(s, " esc(N)")
		}

	default:
		fmt.Fprintf(s, " esc(%d)", n.Esc)
	}

	if e, ok := n.Opt().(*NodeEscState); ok && e.Loopdepth != 0 {
		fmt.Fprintf(s, " ld(%d)", e.Loopdepth)
	}

	if c == 0 && n.Typecheck() != 0 {
		fmt.Fprintf(s, " tc(%d)", n.Typecheck())
	}

	if n.Isddd() {
		fmt.Fprintf(s, " isddd(%v)", n.Isddd())
	}

	if n.Implicit() {
		fmt.Fprintf(s, " implicit(%v)", n.Implicit())
	}

	if n.Embedded() {
		fmt.Fprintf(s, " embedded")
	}

	if n.Addrtaken() {
		fmt.Fprint(s, " addrtaken")
	}

	if n.Assigned() {
		fmt.Fprint(s, " assigned")
	}
	if n.Bounded() {
		fmt.Fprint(s, " bounded")
	}
	if n.NonNil() {
		fmt.Fprint(s, " nonnil")
	}

	if c == 0 && n.HasCall() {
		fmt.Fprint(s, " hascall")
	}

	if c == 0 && n.Name != nil && n.Name.Used() {
		fmt.Fprint(s, " used")
	}
}

func (v Val) Format(pstate *PackageState, s fmt.State, verb rune) {
	switch verb {
	case 'v':
		v.vconv(pstate, s, pstate.fmtFlag(s, verb))

	default:
		fmt.Fprintf(s, "%%!%c(Val=%T)", verb, v)
	}
}

func (v Val) vconv(pstate *PackageState, s fmt.State, flag FmtFlag) {
	switch u := v.U.(type) {
	case *Mpint:
		if !u.Rune {
			if flag&FmtSharp != 0 {
				fmt.Fprint(s, bconv(u, FmtSharp))
				return
			}
			fmt.Fprint(s, bconv(u, 0))
			return
		}

		switch x := u.Int64(pstate); {
		case ' ' <= x && x < utf8.RuneSelf && x != '\\' && x != '\'':
			fmt.Fprintf(s, "'%c'", int(x))

		case 0 <= x && x < 1<<16:
			fmt.Fprintf(s, "'\\u%04x'", uint(int(x)))

		case 0 <= x && x <= utf8.MaxRune:
			fmt.Fprintf(s, "'\\U%08x'", uint64(x))

		default:
			fmt.Fprintf(s, "('\\x00' + %v)", u)
		}

	case *Mpflt:
		if flag&FmtSharp != 0 {
			fmt.Fprint(s, fconv(u, 0))
			return
		}
		fmt.Fprint(s, fconv(u, FmtSharp))
		return

	case *Mpcplx:
		switch {
		case flag&FmtSharp != 0:
			fmt.Fprintf(s, "(%v+%vi)", &u.Real, &u.Imag)

		case v.U.(*Mpcplx).Real.CmpFloat64(0) == 0:
			fmt.Fprintf(s, "%vi", fconv(&u.Imag, FmtSharp))

		case v.U.(*Mpcplx).Imag.CmpFloat64(0) == 0:
			fmt.Fprint(s, fconv(&u.Real, FmtSharp))

		case v.U.(*Mpcplx).Imag.CmpFloat64(0) < 0:
			fmt.Fprintf(s, "(%v%vi)", fconv(&u.Real, FmtSharp), fconv(&u.Imag, FmtSharp))

		default:
			fmt.Fprintf(s, "(%v+%vi)", fconv(&u.Real, FmtSharp), fconv(&u.Imag, FmtSharp))
		}

	case string:
		fmt.Fprint(s, strconv.Quote(u))

	case bool:
		fmt.Fprint(s, u)

	case *NilVal:
		fmt.Fprint(s, "nil")

	default:
		fmt.Fprintf(s, "<ctype=%d>", v.Ctype(pstate))
	}
}

/*
s%,%,\n%g
s%\n+%\n%g
s%^[	]*T%%g
s%,.*%%g
s%.+%	[T&]		= "&",%g
s%^	........*\]%&~%g
s%~	%%g
*/

func (pstate *PackageState) symfmt(s *types.Sym, flag FmtFlag, mode fmtMode) string {
	if s.Pkg != nil && flag&FmtShort == 0 {
		switch mode {
		case FErr: // This is for the user
			if s.Pkg == pstate.builtinpkg || s.Pkg == pstate.localpkg {
				return s.Name
			}

			// If the name was used by multiple packages, display the full path,
			if s.Pkg.Name != "" && pstate.numImport[s.Pkg.Name] > 1 {
				return fmt.Sprintf("%q.%s", s.Pkg.Path, s.Name)
			}
			return s.Pkg.Name + "." + s.Name

		case FDbg:
			return s.Pkg.Name + "." + s.Name

		case FTypeIdName:
			return s.Pkg.Name + "." + s.Name // dcommontype, typehash

		case FTypeId:
			return s.Pkg.Prefix + "." + s.Name // (methodsym), typesym, weaksym
		}
	}

	if flag&FmtByte != 0 {
		// FmtByte (hh) implies FmtShort (h)
		// skip leading "type." in method name
		name := s.Name
		if i := strings.LastIndex(name, "."); i >= 0 {
			name = name[i+1:]
		}

		if mode == FDbg {
			return fmt.Sprintf("@%q.%s", s.Pkg.Path, name)
		}

		return name
	}

	return s.Name
}

func (pstate *PackageState) typefmt(t *types.Type, flag FmtFlag, mode fmtMode, depth int) string {
	if t == nil {
		return "<T>"
	}

	if t == pstate.types.Bytetype || t == pstate.types.Runetype {
		// in %-T mode collapse rune and byte with their originals.
		switch mode {
		case FTypeIdName, FTypeId:
			t = pstate.types.Types[t.Etype]
		default:
			return pstate.sconv(t.Sym, FmtShort, mode)
		}
	}

	if t == pstate.types.Errortype {
		return "error"
	}

	// Unless the 'l' flag was specified, if the type has a name, just print that name.
	if flag&FmtLong == 0 && t.Sym != nil && t != pstate.types.Types[t.Etype] {
		switch mode {
		case FTypeId, FTypeIdName:
			if flag&FmtShort != 0 {
				if t.Vargen != 0 {
					return mode.Sprintf(pstate, "%v·%d", pstate.sconv(t.Sym, FmtShort, mode), t.Vargen)
				}
				return pstate.sconv(t.Sym, FmtShort, mode)
			}

			if mode == FTypeIdName {
				return pstate.sconv(t.Sym, FmtUnsigned, mode)
			}

			if t.Sym.Pkg == pstate.localpkg && t.Vargen != 0 {
				return mode.Sprintf(pstate, "%v·%d", t.Sym, t.Vargen)
			}
		}

		return pstate.smodeString(t.Sym, mode)
	}

	if int(t.Etype) < len(pstate.basicnames) && pstate.basicnames[t.Etype] != "" {
		name := pstate.basicnames[t.Etype]
		if t == pstate.types.Idealbool || t == pstate.types.Idealstring {
			name = "untyped " + name
		}
		return name
	}

	if mode == FDbg {
		return t.Etype.String(pstate.types) + "-" + pstate.typefmt(t, flag, FErr, depth)
	}

	switch t.Etype {
	case TPTR32, TPTR64:
		switch mode {
		case FTypeId, FTypeIdName:
			if flag&FmtShort != 0 {
				return "*" + pstate.tconv(t.Elem(pstate.types), FmtShort, mode, depth)
			}
		}
		return "*" + pstate.tmodeString(t.Elem(pstate.types), mode, depth)

	case TARRAY:
		if t.IsDDDArray() {
			return "[...]" + pstate.tmodeString(t.Elem(pstate.types), mode, depth)
		}
		return "[" + strconv.FormatInt(t.NumElem(pstate.types), 10) + "]" + pstate.tmodeString(t.Elem(pstate.types), mode, depth)

	case TSLICE:
		return "[]" + pstate.tmodeString(t.Elem(pstate.types), mode, depth)

	case TCHAN:
		switch t.ChanDir(pstate.types) {
		case types.Crecv:
			return "<-chan " + pstate.tmodeString(t.Elem(pstate.types), mode, depth)

		case types.Csend:
			return "chan<- " + pstate.tmodeString(t.Elem(pstate.types), mode, depth)
		}

		if t.Elem(pstate.types) != nil && t.Elem(pstate.types).IsChan() && t.Elem(pstate.types).Sym == nil && t.Elem(pstate.types).ChanDir(pstate.types) == types.Crecv {
			return "chan (" + pstate.tmodeString(t.Elem(pstate.types), mode, depth) + ")"
		}
		return "chan " + pstate.tmodeString(t.Elem(pstate.types), mode, depth)

	case TMAP:
		return "map[" + pstate.tmodeString(t.Key(pstate.types), mode, depth) + "]" + pstate.tmodeString(t.Elem(pstate.types), mode, depth)

	case TINTER:
		if t.IsEmptyInterface(pstate.types) {
			return "interface {}"
		}
		buf := make([]byte, 0, 64)
		buf = append(buf, "interface {"...)
		for i, f := range t.Fields(pstate.types).Slice() {
			if i != 0 {
				buf = append(buf, ';')
			}
			buf = append(buf, ' ')
			switch {
			case f.Sym == nil:
				// Check first that a symbol is defined for this type.
				// Wrong interface definitions may have types lacking a symbol.
				break
			case types.IsExported(f.Sym.Name):
				buf = append(buf, pstate.sconv(f.Sym, FmtShort, mode)...)
			default:
				buf = append(buf, pstate.sconv(f.Sym, FmtUnsigned, mode)...)
			}
			buf = append(buf, pstate.tconv(f.Type, FmtShort, mode, depth)...)
		}
		if t.NumFields(pstate.types) != 0 {
			buf = append(buf, ' ')
		}
		buf = append(buf, '}')
		return string(buf)

	case TFUNC:
		buf := make([]byte, 0, 64)
		if flag&FmtShort != 0 {
			// no leading func
		} else {
			if t.Recv(pstate.types) != nil {
				buf = append(buf, "method"...)
				buf = append(buf, pstate.tmodeString(t.Recvs(pstate.types), mode, depth)...)
				buf = append(buf, ' ')
			}
			buf = append(buf, "func"...)
		}
		buf = append(buf, pstate.tmodeString(t.Params(pstate.types), mode, depth)...)

		switch t.NumResults(pstate.types) {
		case 0:
		// nothing to do

		case 1:
			buf = append(buf, ' ')
			buf = append(buf, pstate.tmodeString(t.Results(pstate.types).Field(pstate.types, 0).Type, mode, depth)...) // struct->field->field's type

		default:
			buf = append(buf, ' ')
			buf = append(buf, pstate.tmodeString(t.Results(pstate.types), mode, depth)...)
		}
		return string(buf)

	case TSTRUCT:
		if m := t.StructType(pstate.types).Map; m != nil {
			mt := m.MapType(pstate.types)
			// Format the bucket struct for map[x]y as map.bucket[x]y.
			// This avoids a recursive print that generates very long names.
			var subtype string
			switch t {
			case mt.Bucket:
				subtype = "bucket"
			case mt.Hmap:
				subtype = "hdr"
			case mt.Hiter:
				subtype = "iter"
			default:
				pstate.Fatalf("unknown internal map type")
			}
			return fmt.Sprintf("map.%s[%s]%s", subtype, pstate.tmodeString(m.Key(pstate.types), mode, depth), pstate.tmodeString(m.Elem(pstate.types), mode, depth))
		}

		buf := make([]byte, 0, 64)
		if funarg := t.StructType(pstate.types).Funarg; funarg != types.FunargNone {
			buf = append(buf, '(')
			var flag1 FmtFlag
			switch mode {
			case FTypeId, FTypeIdName, FErr:
				// no argument names on function signature, and no "noescape"/"nosplit" tags
				flag1 = FmtShort
			}
			for i, f := range t.Fields(pstate.types).Slice() {
				if i != 0 {
					buf = append(buf, ", "...)
				}
				buf = append(buf, pstate.fldconv(f, flag1, mode, depth, funarg)...)
			}
			buf = append(buf, ')')
		} else {
			buf = append(buf, "struct {"...)
			for i, f := range t.Fields(pstate.types).Slice() {
				if i != 0 {
					buf = append(buf, ';')
				}
				buf = append(buf, ' ')
				buf = append(buf, pstate.fldconv(f, FmtLong, mode, depth, funarg)...)
			}
			if t.NumFields(pstate.types) != 0 {
				buf = append(buf, ' ')
			}
			buf = append(buf, '}')
		}
		return string(buf)

	case TFORW:
		if t.Sym != nil {
			return "undefined " + pstate.smodeString(t.Sym, mode)
		}
		return "undefined"

	case TUNSAFEPTR:
		return "unsafe.Pointer"

	case TDDDFIELD:
		return mode.Sprintf(pstate, "%v <%v> %v", t.Etype, t.Sym, t.DDDField(pstate.types))

	case Txxx:
		return "Txxx"
	}

	// Don't know how to handle - fall back to detailed prints.
	return mode.Sprintf(pstate, "%v <%v>", t.Etype, t.Sym)
}

// Statements which may be rendered with a simplestmt as init.
func stmtwithinit(op Op) bool {
	switch op {
	case OIF, OFOR, OFORUNTIL, OSWITCH:
		return true
	}

	return false
}

func (n *Node) stmtfmt(pstate *PackageState, s fmt.State, mode fmtMode) {
	// some statements allow for an init, but at most one,
	// but we may have an arbitrary number added, eg by typecheck
	// and inlining. If it doesn't fit the syntax, emit an enclosing
	// block starting with the init statements.

	// if we can just say "for" n->ninit; ... then do so
	simpleinit := n.Ninit.Len() == 1 && n.Ninit.First().Ninit.Len() == 0 && stmtwithinit(n.Op)

	// otherwise, print the inits as separate statements
	complexinit := n.Ninit.Len() != 0 && !simpleinit && (mode != FErr)

	// but if it was for if/for/switch, put in an extra surrounding block to limit the scope
	extrablock := complexinit && stmtwithinit(n.Op)

	if extrablock {
		fmt.Fprint(s, "{")
	}

	if complexinit {
		mode.Fprintf(pstate, s, " %v; ", n.Ninit)
	}

	switch n.Op {
	case ODCL:
		mode.Fprintf(pstate, s, "var %v %v", n.Left.Sym, n.Left.Type)

	case ODCLFIELD:
		if n.Sym != nil {
			mode.Fprintf(pstate, s, "%v %v", n.Sym, n.Left)
		} else {
			mode.Fprintf(pstate, s, "%v", n.Left)
		}

	// Don't export "v = <N>" initializing statements, hope they're always
	// preceded by the DCL which will be re-parsed and typechecked to reproduce
	// the "v = <N>" again.
	case OAS:
		if n.Colas() && !complexinit {
			mode.Fprintf(pstate, s, "%v := %v", n.Left, n.Right)
		} else {
			mode.Fprintf(pstate, s, "%v = %v", n.Left, n.Right)
		}

	case OASOP:
		if n.Implicit() {
			if n.SubOp(pstate) == OADD {
				mode.Fprintf(pstate, s, "%v++", n.Left)
			} else {
				mode.Fprintf(pstate, s, "%v--", n.Left)
			}
			break
		}

		mode.Fprintf(pstate, s, "%v %#v= %v", n.Left, n.SubOp(pstate), n.Right)

	case OAS2:
		if n.Colas() && !complexinit {
			mode.Fprintf(pstate, s, "%.v := %.v", n.List, n.Rlist)
			break
		}
		fallthrough

	case OAS2DOTTYPE, OAS2FUNC, OAS2MAPR, OAS2RECV:
		mode.Fprintf(pstate, s, "%.v = %.v", n.List, n.Rlist)

	case ORETURN:
		mode.Fprintf(pstate, s, "return %.v", n.List)

	case ORETJMP:
		mode.Fprintf(pstate, s, "retjmp %v", n.Sym)

	case OPROC:
		mode.Fprintf(pstate, s, "go %v", n.Left)

	case ODEFER:
		mode.Fprintf(pstate, s, "defer %v", n.Left)

	case OIF:
		if simpleinit {
			mode.Fprintf(pstate, s, "if %v; %v { %v }", n.Ninit.First(), n.Left, n.Nbody)
		} else {
			mode.Fprintf(pstate, s, "if %v { %v }", n.Left, n.Nbody)
		}
		if n.Rlist.Len() != 0 {
			mode.Fprintf(pstate, s, " else { %v }", n.Rlist)
		}

	case OFOR, OFORUNTIL:
		opname := "for"
		if n.Op == OFORUNTIL {
			opname = "foruntil"
		}
		if mode == FErr { // TODO maybe only if FmtShort, same below
			fmt.Fprintf(s, "%s loop", opname)
			break
		}

		fmt.Fprint(s, opname)
		if simpleinit {
			mode.Fprintf(pstate, s, " %v;", n.Ninit.First())
		} else if n.Right != nil {
			fmt.Fprint(s, " ;")
		}

		if n.Left != nil {
			mode.Fprintf(pstate, s, " %v", n.Left)
		}

		if n.Right != nil {
			mode.Fprintf(pstate, s, "; %v", n.Right)
		} else if simpleinit {
			fmt.Fprint(s, ";")
		}

		if n.Op == OFORUNTIL && n.List.Len() != 0 {
			mode.Fprintf(pstate, s, "; %v", n.List)
		}

		mode.Fprintf(pstate, s, " { %v }", n.Nbody)

	case ORANGE:
		if mode == FErr {
			fmt.Fprint(s, "for loop")
			break
		}

		if n.List.Len() == 0 {
			mode.Fprintf(pstate, s, "for range %v { %v }", n.Right, n.Nbody)
			break
		}

		mode.Fprintf(pstate, s, "for %.v = range %v { %v }", n.List, n.Right, n.Nbody)

	case OSELECT, OSWITCH:
		if mode == FErr {
			mode.Fprintf(pstate, s, "%v statement", n.Op)
			break
		}

		mode.Fprintf(pstate, s, "%#v", n.Op)
		if simpleinit {
			mode.Fprintf(pstate, s, " %v;", n.Ninit.First())
		}
		if n.Left != nil {
			mode.Fprintf(pstate, s, " %v ", n.Left)
		}

		mode.Fprintf(pstate, s, " { %v }", n.List)

	case OXCASE:
		if n.List.Len() != 0 {
			mode.Fprintf(pstate, s, "case %.v", n.List)
		} else {
			fmt.Fprint(s, "default")
		}
		mode.Fprintf(pstate, s, ": %v", n.Nbody)

	case OCASE:
		switch {
		case n.Left != nil:
			// single element
			mode.Fprintf(pstate, s, "case %v", n.Left)
		case n.List.Len() > 0:
			// range
			if n.List.Len() != 2 {
				pstate.Fatalf("bad OCASE list length %d", n.List.Len())
			}
			mode.Fprintf(pstate, s, "case %v..%v", n.List.First(), n.List.Second())
		default:
			fmt.Fprint(s, "default")
		}
		mode.Fprintf(pstate, s, ": %v", n.Nbody)

	case OBREAK, OCONTINUE, OGOTO, OFALL:
		if n.Left != nil {
			mode.Fprintf(pstate, s, "%#v %v", n.Op, n.Left)
		} else {
			mode.Fprintf(pstate, s, "%#v", n.Op)
		}

	case OEMPTY:
		break

	case OLABEL:
		mode.Fprintf(pstate, s, "%v: ", n.Left)
	}

	if extrablock {
		fmt.Fprint(s, "}")
	}
}

func (n *Node) exprfmt(pstate *PackageState, s fmt.State, prec int, mode fmtMode) {
	for n != nil && n.Implicit() && (n.Op == OIND || n.Op == OADDR) {
		n = n.Left
	}

	if n == nil {
		fmt.Fprint(s, "<N>")
		return
	}

	nprec := pstate.opprec[n.Op]
	if n.Op == OTYPE && n.Sym != nil {
		nprec = 8
	}

	if prec > nprec {
		mode.Fprintf(pstate, s, "(%v)", n)
		return
	}

	switch n.Op {
	case OPAREN:
		mode.Fprintf(pstate, s, "(%v)", n.Left)

	case ODDDARG:
		fmt.Fprint(s, "... argument")

	case OLITERAL: // this is a bit of a mess
		if mode == FErr {
			if n.Orig != nil && n.Orig != n {
				n.Orig.exprfmt(pstate, s, prec, mode)
				return
			}
			if n.Sym != nil {
				fmt.Fprint(s, pstate.smodeString(n.Sym, mode))
				return
			}
		}
		if n.Val().Ctype(pstate) == CTNIL && n.Orig != nil && n.Orig != n {
			n.Orig.exprfmt(pstate, s, prec, mode)
			return
		}
		if n.Type != nil && n.Type.Etype != TIDEAL && n.Type.Etype != TNIL && n.Type != pstate.types.Idealbool && n.Type != pstate.types.Idealstring {
			// Need parens when type begins with what might
			// be misinterpreted as a unary operator: * or <-.
			if n.Type.IsPtr() || (n.Type.IsChan() && n.Type.ChanDir(pstate.types) == types.Crecv) {
				mode.Fprintf(pstate, s, "(%v)(%v)", n.Type, n.Val())
				return
			} else {
				mode.Fprintf(pstate, s, "%v(%v)", n.Type, n.Val())
				return
			}
		}

		mode.Fprintf(pstate, s, "%v", n.Val())

	// Special case: name used as local variable in export.
	// _ becomes ~b%d internally; print as _ for export
	case ONAME:
		if mode == FErr && n.Sym != nil && n.Sym.Name[0] == '~' && n.Sym.Name[1] == 'b' {
			fmt.Fprint(s, "_")
			return
		}
		fallthrough
	case OPACK, ONONAME:
		fmt.Fprint(s, pstate.smodeString(n.Sym, mode))

	case OTYPE:
		if n.Type == nil && n.Sym != nil {
			fmt.Fprint(s, pstate.smodeString(n.Sym, mode))
			return
		}
		mode.Fprintf(pstate, s, "%v", n.Type)

	case OTARRAY:
		if n.Left != nil {
			mode.Fprintf(pstate, s, "[%v]%v", n.Left, n.Right)
			return
		}
		mode.Fprintf(pstate, s, "[]%v", n.Right) // happens before typecheck

	case OTMAP:
		mode.Fprintf(pstate, s, "map[%v]%v", n.Left, n.Right)

	case OTCHAN:
		switch n.TChanDir(pstate) {
		case types.Crecv:
			mode.Fprintf(pstate, s, "<-chan %v", n.Left)

		case types.Csend:
			mode.Fprintf(pstate, s, "chan<- %v", n.Left)

		default:
			if n.Left != nil && n.Left.Op == OTCHAN && n.Left.Sym == nil && n.Left.TChanDir(pstate) == types.Crecv {
				mode.Fprintf(pstate, s, "chan (%v)", n.Left)
			} else {
				mode.Fprintf(pstate, s, "chan %v", n.Left)
			}
		}

	case OTSTRUCT:
		fmt.Fprint(s, "<struct>")

	case OTINTER:
		fmt.Fprint(s, "<inter>")

	case OTFUNC:
		fmt.Fprint(s, "<func>")

	case OCLOSURE:
		if mode == FErr {
			fmt.Fprint(s, "func literal")
			return
		}
		if n.Nbody.Len() != 0 {
			mode.Fprintf(pstate, s, "%v { %v }", n.Type, n.Nbody)
			return
		}
		mode.Fprintf(pstate, s, "%v { %v }", n.Type, n.Func.Closure.Nbody)

	case OCOMPLIT:
		ptrlit := n.Right != nil && n.Right.Implicit() && n.Right.Type != nil && n.Right.Type.IsPtr()
		if mode == FErr {
			if n.Right != nil && n.Right.Type != nil && !n.Implicit() {
				if ptrlit {
					mode.Fprintf(pstate, s, "&%v literal", n.Right.Type.Elem(pstate.types))
					return
				} else {
					mode.Fprintf(pstate, s, "%v literal", n.Right.Type)
					return
				}
			}

			fmt.Fprint(s, "composite literal")
			return
		}
		mode.Fprintf(pstate, s, "(%v{ %.v })", n.Right, n.List)

	case OPTRLIT:
		mode.Fprintf(pstate, s, "&%v", n.Left)

	case OSTRUCTLIT, OARRAYLIT, OSLICELIT, OMAPLIT:
		if mode == FErr {
			mode.Fprintf(pstate, s, "%v literal", n.Type)
			return
		}
		mode.Fprintf(pstate, s, "(%v{ %.v })", n.Type, n.List)

	case OKEY:
		if n.Left != nil && n.Right != nil {
			mode.Fprintf(pstate, s, "%v:%v", n.Left, n.Right)
			return
		}

		if n.Left == nil && n.Right != nil {
			mode.Fprintf(pstate, s, ":%v", n.Right)
			return
		}
		if n.Left != nil && n.Right == nil {
			mode.Fprintf(pstate, s, "%v:", n.Left)
			return
		}
		fmt.Fprint(s, ":")

	case OSTRUCTKEY:
		mode.Fprintf(pstate, s, "%v:%v", n.Sym, n.Left)

	case OCALLPART:
		n.Left.exprfmt(pstate, s, nprec, mode)
		if n.Right == nil || n.Right.Sym == nil {
			fmt.Fprint(s, ".<nil>")
			return
		}
		mode.Fprintf(pstate, s, ".%0S", n.Right.Sym)

	case OXDOT, ODOT, ODOTPTR, ODOTINTER, ODOTMETH:
		n.Left.exprfmt(pstate, s, nprec, mode)
		if n.Sym == nil {
			fmt.Fprint(s, ".<nil>")
			return
		}
		mode.Fprintf(pstate, s, ".%0S", n.Sym)

	case ODOTTYPE, ODOTTYPE2:
		n.Left.exprfmt(pstate, s, nprec, mode)
		if n.Right != nil {
			mode.Fprintf(pstate, s, ".(%v)", n.Right)
			return
		}
		mode.Fprintf(pstate, s, ".(%v)", n.Type)

	case OINDEX, OINDEXMAP:
		n.Left.exprfmt(pstate, s, nprec, mode)
		mode.Fprintf(pstate, s, "[%v]", n.Right)

	case OSLICE, OSLICESTR, OSLICEARR, OSLICE3, OSLICE3ARR:
		n.Left.exprfmt(pstate, s, nprec, mode)
		fmt.Fprint(s, "[")
		low, high, max := n.SliceBounds(pstate)
		if low != nil {
			fmt.Fprint(s, low.modeString(pstate, mode))
		}
		fmt.Fprint(s, ":")
		if high != nil {
			fmt.Fprint(s, high.modeString(pstate, mode))
		}
		if n.Op.IsSlice3(pstate) {
			fmt.Fprint(s, ":")
			if max != nil {
				fmt.Fprint(s, max.modeString(pstate, mode))
			}
		}
		fmt.Fprint(s, "]")

	case OCOPY, OCOMPLEX:
		mode.Fprintf(pstate, s, "%#v(%v, %v)", n.Op, n.Left, n.Right)

	case OCONV,
		OCONVIFACE,
		OCONVNOP,
		OARRAYBYTESTR,
		OARRAYRUNESTR,
		OSTRARRAYBYTE,
		OSTRARRAYRUNE,
		ORUNESTR:
		if n.Type == nil || n.Type.Sym == nil {
			mode.Fprintf(pstate, s, "(%v)", n.Type)
		} else {
			mode.Fprintf(pstate, s, "%v", n.Type)
		}
		if n.Left != nil {
			mode.Fprintf(pstate, s, "(%v)", n.Left)
		} else {
			mode.Fprintf(pstate, s, "(%.v)", n.List)
		}

	case OREAL,
		OIMAG,
		OAPPEND,
		OCAP,
		OCLOSE,
		ODELETE,
		OLEN,
		OMAKE,
		ONEW,
		OPANIC,
		ORECOVER,
		OALIGNOF,
		OOFFSETOF,
		OSIZEOF,
		OPRINT,
		OPRINTN:
		if n.Left != nil {
			mode.Fprintf(pstate, s, "%#v(%v)", n.Op, n.Left)
			return
		}
		if n.Isddd() {
			mode.Fprintf(pstate, s, "%#v(%.v...)", n.Op, n.List)
			return
		}
		mode.Fprintf(pstate, s, "%#v(%.v)", n.Op, n.List)

	case OCALL, OCALLFUNC, OCALLINTER, OCALLMETH, OGETG:
		n.Left.exprfmt(pstate, s, nprec, mode)
		if n.Isddd() {
			mode.Fprintf(pstate, s, "(%.v...)", n.List)
			return
		}
		mode.Fprintf(pstate, s, "(%.v)", n.List)

	case OMAKEMAP, OMAKECHAN, OMAKESLICE:
		if n.List.Len() != 0 { // pre-typecheck
			mode.Fprintf(pstate, s, "make(%v, %.v)", n.Type, n.List)
			return
		}
		if n.Right != nil {
			mode.Fprintf(pstate, s, "make(%v, %v, %v)", n.Type, n.Left, n.Right)
			return
		}
		if n.Left != nil && (n.Op == OMAKESLICE || !n.Left.Type.IsUntyped(pstate.types)) {
			mode.Fprintf(pstate, s, "make(%v, %v)", n.Type, n.Left)
			return
		}
		mode.Fprintf(pstate, s, "make(%v)", n.Type)

	case OPLUS, OMINUS, OADDR, OCOM, OIND, ONOT, ORECV:
		// Unary
		mode.Fprintf(pstate, s, "%#v", n.Op)
		if n.Left != nil && n.Left.Op == n.Op {
			fmt.Fprint(s, " ")
		}
		n.Left.exprfmt(pstate, s, nprec+1, mode)

	// Binary
	case OADD,
		OAND,
		OANDAND,
		OANDNOT,
		ODIV,
		OEQ,
		OGE,
		OGT,
		OLE,
		OLT,
		OLSH,
		OMOD,
		OMUL,
		ONE,
		OOR,
		OOROR,
		ORSH,
		OSEND,
		OSUB,
		OXOR:
		n.Left.exprfmt(pstate, s, nprec, mode)
		mode.Fprintf(pstate, s, " %#v ", n.Op)
		n.Right.exprfmt(pstate, s, nprec+1, mode)

	case OADDSTR:
		for i, n1 := range n.List.Slice() {
			if i != 0 {
				fmt.Fprint(s, " + ")
			}
			n1.exprfmt(pstate, s, nprec, mode)
		}

	case OCMPSTR, OCMPIFACE:
		n.Left.exprfmt(pstate, s, nprec, mode)
		mode.Fprintf(pstate, s, " %#v ", n.SubOp(pstate))
		n.Right.exprfmt(pstate, s, nprec+1, mode)

	default:
		mode.Fprintf(pstate, s, "<node %v>", n.Op)
	}
}

func (n *Node) nodefmt(pstate *PackageState, s fmt.State, flag FmtFlag, mode fmtMode) {
	t := n.Type

	// We almost always want the original, except in export mode for literals.
	// This saves the importer some work, and avoids us having to redo some
	// special casing for package unsafe.
	if n.Op != OLITERAL && n.Orig != nil {
		n = n.Orig
	}

	if flag&FmtLong != 0 && t != nil {
		if t.Etype == TNIL {
			fmt.Fprint(s, "nil")
		} else {
			mode.Fprintf(pstate, s, "%v (type %v)", n, t)
		}
		return
	}

	// TODO inlining produces expressions with ninits. we can't print these yet.

	if pstate.opprec[n.Op] < 0 {
		n.stmtfmt(pstate, s, mode)
		return
	}

	n.exprfmt(pstate, s, 0, mode)
}

func (n *Node) nodedump(pstate *PackageState, s fmt.State, flag FmtFlag, mode fmtMode) {
	recur := flag&FmtShort == 0

	if recur {
		pstate.indent(s)
		if pstate.dumpdepth > 40 {
			fmt.Fprint(s, "...")
			return
		}

		if n.Ninit.Len() != 0 {
			mode.Fprintf(pstate, s, "%v-init%v", n.Op, n.Ninit)
			pstate.indent(s)
		}
	}

	switch n.Op {
	default:
		mode.Fprintf(pstate, s, "%v%j", n.Op, n)

	case OINDREGSP:
		mode.Fprintf(pstate, s, "%v-SP%j", n.Op, n)

	case OLITERAL:
		mode.Fprintf(pstate, s, "%v-%v%j", n.Op, n.Val(), n)

	case ONAME, ONONAME:
		if n.Sym != nil {
			mode.Fprintf(pstate, s, "%v-%v%j", n.Op, n.Sym, n)
		} else {
			mode.Fprintf(pstate, s, "%v%j", n.Op, n)
		}
		if recur && n.Type == nil && n.Name != nil && n.Name.Param != nil && n.Name.Param.Ntype != nil {
			pstate.indent(s)
			mode.Fprintf(pstate, s, "%v-ntype%v", n.Op, n.Name.Param.Ntype)
		}

	case OASOP:
		mode.Fprintf(pstate, s, "%v-%v%j", n.Op, n.SubOp(pstate), n)

	case OTYPE:
		mode.Fprintf(pstate, s, "%v %v%j type=%v", n.Op, n.Sym, n, n.Type)
		if recur && n.Type == nil && n.Name != nil && n.Name.Param != nil && n.Name.Param.Ntype != nil {
			pstate.indent(s)
			mode.Fprintf(pstate, s, "%v-ntype%v", n.Op, n.Name.Param.Ntype)
		}
	}

	if n.Sym != nil && n.Op != ONAME {
		mode.Fprintf(pstate, s, " %v", n.Sym)
	}

	if n.Type != nil {
		mode.Fprintf(pstate, s, " %v", n.Type)
	}

	if recur {
		if n.Left != nil {
			mode.Fprintf(pstate, s, "%v", n.Left)
		}
		if n.Right != nil {
			mode.Fprintf(pstate, s, "%v", n.Right)
		}
		if n.List.Len() != 0 {
			pstate.indent(s)
			mode.Fprintf(pstate, s, "%v-list%v", n.Op, n.List)
		}

		if n.Rlist.Len() != 0 {
			pstate.indent(s)
			mode.Fprintf(pstate, s, "%v-rlist%v", n.Op, n.Rlist)
		}

		if n.Nbody.Len() != 0 {
			pstate.indent(s)
			mode.Fprintf(pstate, s, "%v-body%v", n.Op, n.Nbody)
		}
	}
}

// "%S" suppresses qualifying with package
func (pstate *PackageState) symFormat(s *types.Sym, f fmt.State, verb rune, mode fmtMode) {
	switch verb {
	case 'v', 'S':
		fmt.Fprint(f, pstate.sconv(s, pstate.fmtFlag(f, verb), mode))

	default:
		fmt.Fprintf(f, "%%!%c(*types.Sym=%p)", verb, s)
	}
}

func (pstate *PackageState) smodeString(s *types.Sym, mode fmtMode) string {
	return pstate.sconv(s, 0, mode)
}

// See #16897 before changing the implementation of sconv.
func (pstate *PackageState) sconv(s *types.Sym, flag FmtFlag, mode fmtMode) string {
	if flag&FmtLong != 0 {
		panic("linksymfmt")
	}

	if s == nil {
		return "<S>"
	}

	if s.Name == "_" {
		return "_"
	}

	flag, mode = flag.update(mode)
	return pstate.symfmt(s, flag, mode)
}

func (pstate *PackageState) tmodeString(t *types.Type, mode fmtMode, depth int) string {
	return pstate.tconv(t, 0, mode, depth)
}

func (pstate *PackageState) fldconv(f *types.Field, flag FmtFlag, mode fmtMode, depth int, funarg types.Funarg) string {
	if f == nil {
		return "<T>"
	}

	flag, mode = flag.update(mode)
	if mode == FTypeIdName {
		flag |= FmtUnsigned
	}

	var name string
	if flag&FmtShort == 0 {
		s := f.Sym

		// Take the name from the original.
		if mode == FErr {
			s = pstate.origSym(s)
		}

		if s != nil && f.Embedded == 0 {
			if funarg != types.FunargNone {
				name = asNode(f.Nname).modeString(pstate, mode)
			} else if flag&FmtLong != 0 {
				name = mode.Sprintf(pstate, "%0S", s)
				if !types.IsExported(name) && flag&FmtUnsigned == 0 {
					name = pstate.smodeString(s, mode) // qualify non-exported names (used on structs, not on funarg)
				}
			} else {
				name = pstate.smodeString(s, mode)
			}
		}
	}

	var typ string
	if f.Isddd() {
		var et *types.Type
		if f.Type != nil {
			et = f.Type.Elem(pstate.types)
		}
		typ = "..." + pstate.tmodeString(et, mode, depth)
	} else {
		typ = pstate.tmodeString(f.Type, mode, depth)
	}

	str := typ
	if name != "" {
		str = name + " " + typ
	}

	if flag&FmtShort == 0 && funarg == types.FunargNone && f.Note != "" {
		str += " " + strconv.Quote(f.Note)
	}

	return str
}

// "%L"  print definition, not name
// "%S"  omit 'func' and receiver from function types, short type names
func (pstate *PackageState) typeFormat(t *types.Type, s fmt.State, verb rune, mode fmtMode) {
	switch verb {
	case 'v', 'S', 'L':
		// This is an external entry point, so we pass depth 0 to tconv.
		// See comments in Type.String.
		fmt.Fprint(s, pstate.tconv(t, pstate.fmtFlag(s, verb), mode, 0))

	default:
		fmt.Fprintf(s, "%%!%c(*Type=%p)", verb, t)
	}
}

// See #16897 before changing the implementation of tconv.
func (pstate *PackageState) tconv(t *types.Type, flag FmtFlag, mode fmtMode, depth int) string {
	if t == nil {
		return "<T>"
	}
	if t.Etype == types.TSSA {
		return t.Extra.(string)
	}
	if t.Etype == types.TTUPLE {
		return t.FieldType(pstate.types, 0).String(pstate.types) + "," + t.FieldType(pstate.types, 1).String(pstate.types)
	}

	if depth > 100 {
		return "<...>"
	}

	flag, mode = flag.update(mode)
	if mode == FTypeIdName {
		flag |= FmtUnsigned
	}

	str := pstate.typefmt(t, flag, mode, depth+1)

	return str
}

func (n *Node) String() string                                       { return fmt.Sprint(n) }
func (n *Node) modeString(pstate *PackageState, mode fmtMode) string { return mode.Sprint(pstate, n) }

// "%L"  suffix with "(type %T)" where possible
// "%+S" in debug mode, don't recurse, no multiline output
func (n *Node) nconv(pstate *PackageState, s fmt.State, flag FmtFlag, mode fmtMode) {
	if n == nil {
		fmt.Fprint(s, "<N>")
		return
	}

	flag, mode = flag.update(mode)

	switch mode {
	case FErr:
		n.nodefmt(pstate, s, flag, mode)

	case FDbg:
		pstate.dumpdepth++
		n.nodedump(pstate, s, flag, mode)
		pstate.dumpdepth--

	default:
		pstate.Fatalf("unhandled %%N mode: %d", mode)
	}
}

func (l Nodes) format(pstate *PackageState, s fmt.State, verb rune, mode fmtMode) {
	switch verb {
	case 'v':
		l.hconv(pstate, s, pstate.fmtFlag(s, verb), mode)

	default:
		fmt.Fprintf(s, "%%!%c(Nodes)", verb)
	}
}

func (n Nodes) String() string {
	return fmt.Sprint(n)
}

// Flags: all those of %N plus '.': separate with comma's instead of semicolons.
func (l Nodes) hconv(pstate *PackageState, s fmt.State, flag FmtFlag, mode fmtMode) {
	if l.Len() == 0 && mode == FDbg {
		fmt.Fprint(s, "<nil>")
		return
	}

	flag, mode = flag.update(mode)
	sep := "; "
	if mode == FDbg {
		sep = "\n"
	} else if flag&FmtComma != 0 {
		sep = ", "
	}

	for i, n := range l.Slice() {
		fmt.Fprint(s, n.modeString(pstate, mode))
		if i+1 < l.Len() {
			fmt.Fprint(s, sep)
		}
	}
}

func dumplist(s string, l Nodes) {
	fmt.Printf("%s%+v\n", s, l)
}

func Dump(s string, n *Node) {
	fmt.Printf("%s [%p]%+v\n", s, n, n)
}

// indent prints indentation to s.
func (pstate *PackageState) indent(s fmt.State) {
	fmt.Fprint(s, "\n")
	for i := 0; i < pstate.dumpdepth; i++ {
		fmt.Fprint(s, ".   ")
	}
}

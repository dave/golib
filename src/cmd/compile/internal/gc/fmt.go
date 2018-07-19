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

const (
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
func (psess *PackageSession) fmtFlag(s fmt.State, verb rune) FmtFlag {
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
		psess.
			Fatalf("FmtUnsigned in format string")
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

// *types.Sym, *types.Type, and *Node types use the flags below to set the format mode
const (
	FErr = iota
	FDbg
	FTypeId
	FTypeIdName // same as FTypeId, but use package name instead of prefix
)

// update returns the results of applying f to mode.
func (f FmtFlag) update(mode fmtMode) (FmtFlag, fmtMode) {
	switch {
	case f&FmtSign != 0:
		mode = FDbg
	case f&FmtSharp != 0:

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

func (o Op) format(psess *PackageSession, s fmt.State, verb rune, mode fmtMode) {
	switch verb {
	case 'v':
		o.oconv(psess, s, psess.fmtFlag(s, verb), mode)

	default:
		fmt.Fprintf(s, "%%!%c(Op=%d)", verb, int(o))
	}
}

func (o Op) oconv(psess *PackageSession, s fmt.State, flag FmtFlag, mode fmtMode) {
	if flag&FmtSharp != 0 || mode != FDbg {
		if int(o) < len(psess.goopnames) && psess.goopnames[o] != "" {
			fmt.Fprint(s, psess.goopnames[o])
			return
		}
	}

	fmt.Fprint(s, o.String(psess))
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

func (n *fmtNodeErr) Format(psess *PackageSession, s fmt.State, verb rune) {
	(*Node)(n).format(psess, s, verb, FErr)
}
func (n *fmtNodeDbg) Format(psess *PackageSession, s fmt.State, verb rune) {
	(*Node)(n).format(psess, s, verb, FDbg)
}
func (n *fmtNodeTypeId) Format(psess *PackageSession, s fmt.State, verb rune) {
	(*Node)(n).format(psess, s, verb, FTypeId)
}
func (n *fmtNodeTypeIdName) Format(psess *PackageSession, s fmt.State, verb rune) {
	(*Node)(n).format(psess, s, verb, FTypeIdName)
}
func (n *Node) Format(psess *PackageSession, s fmt.State, verb rune) { n.format(psess, s, verb, FErr) }

func (o fmtOpErr) Format(psess *PackageSession, s fmt.State, verb rune) {
	Op(o).format(psess, s, verb, FErr)
}
func (o fmtOpDbg) Format(psess *PackageSession, s fmt.State, verb rune) {
	Op(o).format(psess, s, verb, FDbg)
}
func (o fmtOpTypeId) Format(psess *PackageSession, s fmt.State, verb rune) {
	Op(o).format(psess, s, verb, FTypeId)
}
func (o fmtOpTypeIdName) Format(psess *PackageSession, s fmt.State, verb rune) {
	Op(o).format(psess, s, verb, FTypeIdName)
}
func (o Op) Format(psess *PackageSession, s fmt.State, verb rune) { o.format(psess, s, verb, FErr) }

func (t *fmtTypeErr) Format(psess *PackageSession, s fmt.State, verb rune) {
	psess.typeFormat((*types.Type)(t), s, verb, FErr)
}
func (t *fmtTypeDbg) Format(psess *PackageSession, s fmt.State, verb rune) {
	psess.typeFormat((*types.Type)(t), s, verb, FDbg)
}
func (t *fmtTypeTypeId) Format(psess *PackageSession, s fmt.State, verb rune) {
	psess.typeFormat((*types.Type)(t), s, verb, FTypeId)
}
func (t *fmtTypeTypeIdName) Format(psess *PackageSession, s fmt.State, verb rune) {
	psess.
		typeFormat((*types.Type)(t), s, verb, FTypeIdName)
}

func (y *fmtSymErr) Format(psess *PackageSession, s fmt.State, verb rune) {
	psess.symFormat((*types.Sym)(y), s, verb, FErr)
}
func (y *fmtSymDbg) Format(psess *PackageSession, s fmt.State, verb rune) {
	psess.symFormat((*types.Sym)(y), s, verb, FDbg)
}
func (y *fmtSymTypeId) Format(psess *PackageSession, s fmt.State, verb rune) {
	psess.symFormat((*types.Sym)(y), s, verb, FTypeId)
}
func (y *fmtSymTypeIdName) Format(psess *PackageSession, s fmt.State, verb rune) {
	psess.
		symFormat((*types.Sym)(y), s, verb, FTypeIdName)
}

func (n fmtNodesErr) Format(psess *PackageSession, s fmt.State, verb rune) {
	(Nodes)(n).format(psess, s, verb, FErr)
}
func (n fmtNodesDbg) Format(psess *PackageSession, s fmt.State, verb rune) {
	(Nodes)(n).format(psess, s, verb, FDbg)
}
func (n fmtNodesTypeId) Format(psess *PackageSession, s fmt.State, verb rune) {
	(Nodes)(n).format(psess, s, verb, FTypeId)
}
func (n fmtNodesTypeIdName) Format(psess *PackageSession, s fmt.State, verb rune) {
	(Nodes)(n).format(psess, s, verb, FTypeIdName)
}
func (n Nodes) Format(psess *PackageSession, s fmt.State, verb rune) { n.format(psess, s, verb, FErr) }

func (m fmtMode) Fprintf(psess *PackageSession, s fmt.State, format string, args ...interface{}) {
	m.prepareArgs(psess, args)
	fmt.Fprintf(s, format, args...)
}

func (m fmtMode) Sprintf(psess *PackageSession, format string, args ...interface{}) string {
	m.prepareArgs(psess, args)
	return fmt.Sprintf(format, args...)
}

func (m fmtMode) Sprint(psess *PackageSession, args ...interface{}) string {
	m.prepareArgs(psess, args)
	return fmt.Sprint(args...)
}

func (m fmtMode) prepareArgs(psess *PackageSession, args []interface{}) {
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

			default:
				psess.
					Fatalf("mode.prepareArgs type %T", arg)
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

			default:
				psess.
					Fatalf("mode.prepareArgs type %T", arg)
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

			default:
				psess.
					Fatalf("mode.prepareArgs type %T", arg)
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

			default:
				psess.
					Fatalf("mode.prepareArgs type %T", arg)
			}
		}
	default:
		psess.
			Fatalf("mode.prepareArgs mode %d", m)
	}
}

func (n *Node) format(psess *PackageSession, s fmt.State, verb rune, mode fmtMode) {
	switch verb {
	case 'v', 'S', 'L':
		n.nconv(psess, s, psess.fmtFlag(s, verb), mode)

	case 'j':
		n.jconv(s, psess.fmtFlag(s, verb))

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

func (v Val) Format(psess *PackageSession, s fmt.State, verb rune) {
	switch verb {
	case 'v':
		v.vconv(psess, s, psess.fmtFlag(s, verb))

	default:
		fmt.Fprintf(s, "%%!%c(Val=%T)", verb, v)
	}
}

func (v Val) vconv(psess *PackageSession, s fmt.State, flag FmtFlag) {
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

		switch x := u.Int64(psess); {
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
		fmt.Fprintf(s, "<ctype=%d>", v.Ctype(psess))
	}
}

func (psess *PackageSession) symfmt(s *types.Sym, flag FmtFlag, mode fmtMode) string {
	if s.Pkg != nil && flag&FmtShort == 0 {
		switch mode {
		case FErr:
			if s.Pkg == psess.builtinpkg || s.Pkg == psess.localpkg {
				return s.Name
			}

			if s.Pkg.Name != "" && psess.numImport[s.Pkg.Name] > 1 {
				return fmt.Sprintf("%q.%s", s.Pkg.Path, s.Name)
			}
			return s.Pkg.Name + "." + s.Name

		case FDbg:
			return s.Pkg.Name + "." + s.Name

		case FTypeIdName:
			return s.Pkg.Name + "." + s.Name

		case FTypeId:
			return s.Pkg.Prefix + "." + s.Name
		}
	}

	if flag&FmtByte != 0 {

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

func (psess *PackageSession) typefmt(t *types.Type, flag FmtFlag, mode fmtMode, depth int) string {
	if t == nil {
		return "<T>"
	}

	if t == psess.types.Bytetype || t == psess.types.Runetype {

		switch mode {
		case FTypeIdName, FTypeId:
			t = psess.types.Types[t.Etype]
		default:
			return psess.sconv(t.Sym, FmtShort, mode)
		}
	}

	if t == psess.types.Errortype {
		return "error"
	}

	if flag&FmtLong == 0 && t.Sym != nil && t != psess.types.Types[t.Etype] {
		switch mode {
		case FTypeId, FTypeIdName:
			if flag&FmtShort != 0 {
				if t.Vargen != 0 {
					return mode.Sprintf(psess, "%v·%d", psess.sconv(t.Sym, FmtShort, mode), t.Vargen)
				}
				return psess.sconv(t.Sym, FmtShort, mode)
			}

			if mode == FTypeIdName {
				return psess.sconv(t.Sym, FmtUnsigned, mode)
			}

			if t.Sym.Pkg == psess.localpkg && t.Vargen != 0 {
				return mode.Sprintf(psess, "%v·%d", t.Sym, t.Vargen)
			}
		}

		return psess.smodeString(t.Sym, mode)
	}

	if int(t.Etype) < len(psess.basicnames) && psess.basicnames[t.Etype] != "" {
		name := psess.basicnames[t.Etype]
		if t == psess.types.Idealbool || t == psess.types.Idealstring {
			name = "untyped " + name
		}
		return name
	}

	if mode == FDbg {
		return t.Etype.String(psess.types) + "-" + psess.typefmt(t, flag, FErr, depth)
	}

	switch t.Etype {
	case TPTR32, TPTR64:
		switch mode {
		case FTypeId, FTypeIdName:
			if flag&FmtShort != 0 {
				return "*" + psess.tconv(t.Elem(psess.types), FmtShort, mode, depth)
			}
		}
		return "*" + psess.tmodeString(t.Elem(psess.types), mode, depth)

	case TARRAY:
		if t.IsDDDArray() {
			return "[...]" + psess.tmodeString(t.Elem(psess.types), mode, depth)
		}
		return "[" + strconv.FormatInt(t.NumElem(psess.types), 10) + "]" + psess.tmodeString(t.Elem(psess.types), mode, depth)

	case TSLICE:
		return "[]" + psess.tmodeString(t.Elem(psess.types), mode, depth)

	case TCHAN:
		switch t.ChanDir(psess.types) {
		case types.Crecv:
			return "<-chan " + psess.tmodeString(t.Elem(psess.types), mode, depth)

		case types.Csend:
			return "chan<- " + psess.tmodeString(t.Elem(psess.types), mode, depth)
		}

		if t.Elem(psess.types) != nil && t.Elem(psess.types).IsChan() && t.Elem(psess.types).Sym == nil && t.Elem(psess.types).ChanDir(psess.types) == types.Crecv {
			return "chan (" + psess.tmodeString(t.Elem(psess.types), mode, depth) + ")"
		}
		return "chan " + psess.tmodeString(t.Elem(psess.types), mode, depth)

	case TMAP:
		return "map[" + psess.tmodeString(t.Key(psess.types), mode, depth) + "]" + psess.tmodeString(t.Elem(psess.types), mode, depth)

	case TINTER:
		if t.IsEmptyInterface(psess.types) {
			return "interface {}"
		}
		buf := make([]byte, 0, 64)
		buf = append(buf, "interface {"...)
		for i, f := range t.Fields(psess.types).Slice() {
			if i != 0 {
				buf = append(buf, ';')
			}
			buf = append(buf, ' ')
			switch {
			case f.Sym == nil:

				break
			case types.IsExported(f.Sym.Name):
				buf = append(buf, psess.sconv(f.Sym, FmtShort, mode)...)
			default:
				buf = append(buf, psess.sconv(f.Sym, FmtUnsigned, mode)...)
			}
			buf = append(buf, psess.tconv(f.Type, FmtShort, mode, depth)...)
		}
		if t.NumFields(psess.types) != 0 {
			buf = append(buf, ' ')
		}
		buf = append(buf, '}')
		return string(buf)

	case TFUNC:
		buf := make([]byte, 0, 64)
		if flag&FmtShort != 0 {

		} else {
			if t.Recv(psess.types) != nil {
				buf = append(buf, "method"...)
				buf = append(buf, psess.tmodeString(t.Recvs(psess.types), mode, depth)...)
				buf = append(buf, ' ')
			}
			buf = append(buf, "func"...)
		}
		buf = append(buf, psess.tmodeString(t.Params(psess.types), mode, depth)...)

		switch t.NumResults(psess.types) {
		case 0:

		case 1:
			buf = append(buf, ' ')
			buf = append(buf, psess.tmodeString(t.Results(psess.types).Field(psess.types, 0).Type, mode, depth)...)

		default:
			buf = append(buf, ' ')
			buf = append(buf, psess.tmodeString(t.Results(psess.types), mode, depth)...)
		}
		return string(buf)

	case TSTRUCT:
		if m := t.StructType(psess.types).Map; m != nil {
			mt := m.MapType(psess.
				// Format the bucket struct for map[x]y as map.bucket[x]y.
				// This avoids a recursive print that generates very long names.
				types)

			var subtype string
			switch t {
			case mt.Bucket:
				subtype = "bucket"
			case mt.Hmap:
				subtype = "hdr"
			case mt.Hiter:
				subtype = "iter"
			default:
				psess.
					Fatalf("unknown internal map type")
			}
			return fmt.Sprintf("map.%s[%s]%s", subtype, psess.tmodeString(m.Key(psess.types), mode, depth), psess.tmodeString(m.Elem(psess.types), mode, depth))
		}

		buf := make([]byte, 0, 64)
		if funarg := t.StructType(psess.types).Funarg; funarg != types.FunargNone {
			buf = append(buf, '(')
			var flag1 FmtFlag
			switch mode {
			case FTypeId, FTypeIdName, FErr:

				flag1 = FmtShort
			}
			for i, f := range t.Fields(psess.types).Slice() {
				if i != 0 {
					buf = append(buf, ", "...)
				}
				buf = append(buf, psess.fldconv(f, flag1, mode, depth, funarg)...)
			}
			buf = append(buf, ')')
		} else {
			buf = append(buf, "struct {"...)
			for i, f := range t.Fields(psess.types).Slice() {
				if i != 0 {
					buf = append(buf, ';')
				}
				buf = append(buf, ' ')
				buf = append(buf, psess.fldconv(f, FmtLong, mode, depth, funarg)...)
			}
			if t.NumFields(psess.types) != 0 {
				buf = append(buf, ' ')
			}
			buf = append(buf, '}')
		}
		return string(buf)

	case TFORW:
		if t.Sym != nil {
			return "undefined " + psess.smodeString(t.Sym, mode)
		}
		return "undefined"

	case TUNSAFEPTR:
		return "unsafe.Pointer"

	case TDDDFIELD:
		return mode.Sprintf(psess, "%v <%v> %v", t.Etype, t.Sym, t.DDDField(psess.types))

	case Txxx:
		return "Txxx"
	}

	return mode.Sprintf(psess, "%v <%v>", t.Etype, t.Sym)
}

// Statements which may be rendered with a simplestmt as init.
func stmtwithinit(op Op) bool {
	switch op {
	case OIF, OFOR, OFORUNTIL, OSWITCH:
		return true
	}

	return false
}

func (n *Node) stmtfmt(psess *PackageSession, s fmt.State, mode fmtMode) {

	simpleinit := n.Ninit.Len() == 1 && n.Ninit.First().Ninit.Len() == 0 && stmtwithinit(n.Op)

	complexinit := n.Ninit.Len() != 0 && !simpleinit && (mode != FErr)

	extrablock := complexinit && stmtwithinit(n.Op)

	if extrablock {
		fmt.Fprint(s, "{")
	}

	if complexinit {
		mode.Fprintf(psess, s, " %v; ", n.Ninit)
	}

	switch n.Op {
	case ODCL:
		mode.Fprintf(psess, s, "var %v %v", n.Left.Sym, n.Left.Type)

	case ODCLFIELD:
		if n.Sym != nil {
			mode.Fprintf(psess, s, "%v %v", n.Sym, n.Left)
		} else {
			mode.Fprintf(psess, s, "%v", n.Left)
		}

	case OAS:
		if n.Colas() && !complexinit {
			mode.Fprintf(psess, s, "%v := %v", n.Left, n.Right)
		} else {
			mode.Fprintf(psess, s, "%v = %v", n.Left, n.Right)
		}

	case OASOP:
		if n.Implicit() {
			if n.SubOp(psess) == OADD {
				mode.Fprintf(psess, s, "%v++", n.Left)
			} else {
				mode.Fprintf(psess, s, "%v--", n.Left)
			}
			break
		}

		mode.Fprintf(psess, s, "%v %#v= %v", n.Left, n.SubOp(psess), n.Right)

	case OAS2:
		if n.Colas() && !complexinit {
			mode.Fprintf(psess, s, "%.v := %.v", n.List, n.Rlist)
			break
		}
		fallthrough

	case OAS2DOTTYPE, OAS2FUNC, OAS2MAPR, OAS2RECV:
		mode.Fprintf(psess, s, "%.v = %.v", n.List, n.Rlist)

	case ORETURN:
		mode.Fprintf(psess, s, "return %.v", n.List)

	case ORETJMP:
		mode.Fprintf(psess, s, "retjmp %v", n.Sym)

	case OPROC:
		mode.Fprintf(psess, s, "go %v", n.Left)

	case ODEFER:
		mode.Fprintf(psess, s, "defer %v", n.Left)

	case OIF:
		if simpleinit {
			mode.Fprintf(psess, s, "if %v; %v { %v }", n.Ninit.First(), n.Left, n.Nbody)
		} else {
			mode.Fprintf(psess, s, "if %v { %v }", n.Left, n.Nbody)
		}
		if n.Rlist.Len() != 0 {
			mode.Fprintf(psess, s, " else { %v }", n.Rlist)
		}

	case OFOR, OFORUNTIL:
		opname := "for"
		if n.Op == OFORUNTIL {
			opname = "foruntil"
		}
		if mode == FErr {
			fmt.Fprintf(s, "%s loop", opname)
			break
		}

		fmt.Fprint(s, opname)
		if simpleinit {
			mode.Fprintf(psess, s, " %v;", n.Ninit.First())
		} else if n.Right != nil {
			fmt.Fprint(s, " ;")
		}

		if n.Left != nil {
			mode.Fprintf(psess, s, " %v", n.Left)
		}

		if n.Right != nil {
			mode.Fprintf(psess, s, "; %v", n.Right)
		} else if simpleinit {
			fmt.Fprint(s, ";")
		}

		if n.Op == OFORUNTIL && n.List.Len() != 0 {
			mode.Fprintf(psess, s, "; %v", n.List)
		}

		mode.Fprintf(psess, s, " { %v }", n.Nbody)

	case ORANGE:
		if mode == FErr {
			fmt.Fprint(s, "for loop")
			break
		}

		if n.List.Len() == 0 {
			mode.Fprintf(psess, s, "for range %v { %v }", n.Right, n.Nbody)
			break
		}

		mode.Fprintf(psess, s, "for %.v = range %v { %v }", n.List, n.Right, n.Nbody)

	case OSELECT, OSWITCH:
		if mode == FErr {
			mode.Fprintf(psess, s, "%v statement", n.Op)
			break
		}

		mode.Fprintf(psess, s, "%#v", n.Op)
		if simpleinit {
			mode.Fprintf(psess, s, " %v;", n.Ninit.First())
		}
		if n.Left != nil {
			mode.Fprintf(psess, s, " %v ", n.Left)
		}

		mode.Fprintf(psess, s, " { %v }", n.List)

	case OXCASE:
		if n.List.Len() != 0 {
			mode.Fprintf(psess, s, "case %.v", n.List)
		} else {
			fmt.Fprint(s, "default")
		}
		mode.Fprintf(psess, s, ": %v", n.Nbody)

	case OCASE:
		switch {
		case n.Left != nil:

			mode.Fprintf(psess, s, "case %v", n.Left)
		case n.List.Len() > 0:

			if n.List.Len() != 2 {
				psess.
					Fatalf("bad OCASE list length %d", n.List.Len())
			}
			mode.Fprintf(psess, s, "case %v..%v", n.List.First(), n.List.Second())
		default:
			fmt.Fprint(s, "default")
		}
		mode.Fprintf(psess, s, ": %v", n.Nbody)

	case OBREAK, OCONTINUE, OGOTO, OFALL:
		if n.Left != nil {
			mode.Fprintf(psess, s, "%#v %v", n.Op, n.Left)
		} else {
			mode.Fprintf(psess, s, "%#v", n.Op)
		}

	case OEMPTY:
		break

	case OLABEL:
		mode.Fprintf(psess, s, "%v: ", n.Left)
	}

	if extrablock {
		fmt.Fprint(s, "}")
	}
}

func (n *Node) exprfmt(psess *PackageSession, s fmt.State, prec int, mode fmtMode) {
	for n != nil && n.Implicit() && (n.Op == OIND || n.Op == OADDR) {
		n = n.Left
	}

	if n == nil {
		fmt.Fprint(s, "<N>")
		return
	}

	nprec := psess.opprec[n.Op]
	if n.Op == OTYPE && n.Sym != nil {
		nprec = 8
	}

	if prec > nprec {
		mode.Fprintf(psess, s, "(%v)", n)
		return
	}

	switch n.Op {
	case OPAREN:
		mode.Fprintf(psess, s, "(%v)", n.Left)

	case ODDDARG:
		fmt.Fprint(s, "... argument")

	case OLITERAL:
		if mode == FErr {
			if n.Orig != nil && n.Orig != n {
				n.Orig.exprfmt(psess, s, prec, mode)
				return
			}
			if n.Sym != nil {
				fmt.Fprint(s, psess.smodeString(n.Sym, mode))
				return
			}
		}
		if n.Val().Ctype(psess) == CTNIL && n.Orig != nil && n.Orig != n {
			n.Orig.exprfmt(psess, s, prec, mode)
			return
		}
		if n.Type != nil && n.Type.Etype != TIDEAL && n.Type.Etype != TNIL && n.Type != psess.types.Idealbool && n.Type != psess.types.Idealstring {

			if n.Type.IsPtr() || (n.Type.IsChan() && n.Type.ChanDir(psess.types) == types.Crecv) {
				mode.Fprintf(psess, s, "(%v)(%v)", n.Type, n.Val())
				return
			} else {
				mode.Fprintf(psess, s, "%v(%v)", n.Type, n.Val())
				return
			}
		}

		mode.Fprintf(psess, s, "%v", n.Val())

	case ONAME:
		if mode == FErr && n.Sym != nil && n.Sym.Name[0] == '~' && n.Sym.Name[1] == 'b' {
			fmt.Fprint(s, "_")
			return
		}
		fallthrough
	case OPACK, ONONAME:
		fmt.Fprint(s, psess.smodeString(n.Sym, mode))

	case OTYPE:
		if n.Type == nil && n.Sym != nil {
			fmt.Fprint(s, psess.smodeString(n.Sym, mode))
			return
		}
		mode.Fprintf(psess, s, "%v", n.Type)

	case OTARRAY:
		if n.Left != nil {
			mode.Fprintf(psess, s, "[%v]%v", n.Left, n.Right)
			return
		}
		mode.Fprintf(psess, s, "[]%v", n.Right)

	case OTMAP:
		mode.Fprintf(psess, s, "map[%v]%v", n.Left, n.Right)

	case OTCHAN:
		switch n.TChanDir(psess) {
		case types.Crecv:
			mode.Fprintf(psess, s, "<-chan %v", n.Left)

		case types.Csend:
			mode.Fprintf(psess, s, "chan<- %v", n.Left)

		default:
			if n.Left != nil && n.Left.Op == OTCHAN && n.Left.Sym == nil && n.Left.TChanDir(psess) == types.Crecv {
				mode.Fprintf(psess, s, "chan (%v)", n.Left)
			} else {
				mode.Fprintf(psess, s, "chan %v", n.Left)
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
			mode.Fprintf(psess, s, "%v { %v }", n.Type, n.Nbody)
			return
		}
		mode.Fprintf(psess, s, "%v { %v }", n.Type, n.Func.Closure.Nbody)

	case OCOMPLIT:
		ptrlit := n.Right != nil && n.Right.Implicit() && n.Right.Type != nil && n.Right.Type.IsPtr()
		if mode == FErr {
			if n.Right != nil && n.Right.Type != nil && !n.Implicit() {
				if ptrlit {
					mode.Fprintf(psess, s, "&%v literal", n.Right.Type.Elem(psess.types))
					return
				} else {
					mode.Fprintf(psess, s, "%v literal", n.Right.Type)
					return
				}
			}

			fmt.Fprint(s, "composite literal")
			return
		}
		mode.Fprintf(psess, s, "(%v{ %.v })", n.Right, n.List)

	case OPTRLIT:
		mode.Fprintf(psess, s, "&%v", n.Left)

	case OSTRUCTLIT, OARRAYLIT, OSLICELIT, OMAPLIT:
		if mode == FErr {
			mode.Fprintf(psess, s, "%v literal", n.Type)
			return
		}
		mode.Fprintf(psess, s, "(%v{ %.v })", n.Type, n.List)

	case OKEY:
		if n.Left != nil && n.Right != nil {
			mode.Fprintf(psess, s, "%v:%v", n.Left, n.Right)
			return
		}

		if n.Left == nil && n.Right != nil {
			mode.Fprintf(psess, s, ":%v", n.Right)
			return
		}
		if n.Left != nil && n.Right == nil {
			mode.Fprintf(psess, s, "%v:", n.Left)
			return
		}
		fmt.Fprint(s, ":")

	case OSTRUCTKEY:
		mode.Fprintf(psess, s, "%v:%v", n.Sym, n.Left)

	case OCALLPART:
		n.Left.exprfmt(psess, s, nprec, mode)
		if n.Right == nil || n.Right.Sym == nil {
			fmt.Fprint(s, ".<nil>")
			return
		}
		mode.Fprintf(psess, s, ".%0S", n.Right.Sym)

	case OXDOT, ODOT, ODOTPTR, ODOTINTER, ODOTMETH:
		n.Left.exprfmt(psess, s, nprec, mode)
		if n.Sym == nil {
			fmt.Fprint(s, ".<nil>")
			return
		}
		mode.Fprintf(psess, s, ".%0S", n.Sym)

	case ODOTTYPE, ODOTTYPE2:
		n.Left.exprfmt(psess, s, nprec, mode)
		if n.Right != nil {
			mode.Fprintf(psess, s, ".(%v)", n.Right)
			return
		}
		mode.Fprintf(psess, s, ".(%v)", n.Type)

	case OINDEX, OINDEXMAP:
		n.Left.exprfmt(psess, s, nprec, mode)
		mode.Fprintf(psess, s, "[%v]", n.Right)

	case OSLICE, OSLICESTR, OSLICEARR, OSLICE3, OSLICE3ARR:
		n.Left.exprfmt(psess, s, nprec, mode)
		fmt.Fprint(s, "[")
		low, high, max := n.SliceBounds(psess)
		if low != nil {
			fmt.Fprint(s, low.modeString(psess, mode))
		}
		fmt.Fprint(s, ":")
		if high != nil {
			fmt.Fprint(s, high.modeString(psess, mode))
		}
		if n.Op.IsSlice3(psess) {
			fmt.Fprint(s, ":")
			if max != nil {
				fmt.Fprint(s, max.modeString(psess, mode))
			}
		}
		fmt.Fprint(s, "]")

	case OCOPY, OCOMPLEX:
		mode.Fprintf(psess, s, "%#v(%v, %v)", n.Op, n.Left, n.Right)

	case OCONV,
		OCONVIFACE,
		OCONVNOP,
		OARRAYBYTESTR,
		OARRAYRUNESTR,
		OSTRARRAYBYTE,
		OSTRARRAYRUNE,
		ORUNESTR:
		if n.Type == nil || n.Type.Sym == nil {
			mode.Fprintf(psess, s, "(%v)", n.Type)
		} else {
			mode.Fprintf(psess, s, "%v", n.Type)
		}
		if n.Left != nil {
			mode.Fprintf(psess, s, "(%v)", n.Left)
		} else {
			mode.Fprintf(psess, s, "(%.v)", n.List)
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
			mode.Fprintf(psess, s, "%#v(%v)", n.Op, n.Left)
			return
		}
		if n.Isddd() {
			mode.Fprintf(psess, s, "%#v(%.v...)", n.Op, n.List)
			return
		}
		mode.Fprintf(psess, s, "%#v(%.v)", n.Op, n.List)

	case OCALL, OCALLFUNC, OCALLINTER, OCALLMETH, OGETG:
		n.Left.exprfmt(psess, s, nprec, mode)
		if n.Isddd() {
			mode.Fprintf(psess, s, "(%.v...)", n.List)
			return
		}
		mode.Fprintf(psess, s, "(%.v)", n.List)

	case OMAKEMAP, OMAKECHAN, OMAKESLICE:
		if n.List.Len() != 0 {
			mode.Fprintf(psess, s, "make(%v, %.v)", n.Type, n.List)
			return
		}
		if n.Right != nil {
			mode.Fprintf(psess, s, "make(%v, %v, %v)", n.Type, n.Left, n.Right)
			return
		}
		if n.Left != nil && (n.Op == OMAKESLICE || !n.Left.Type.IsUntyped(psess.types)) {
			mode.Fprintf(psess, s, "make(%v, %v)", n.Type, n.Left)
			return
		}
		mode.Fprintf(psess, s, "make(%v)", n.Type)

	case OPLUS, OMINUS, OADDR, OCOM, OIND, ONOT, ORECV:

		mode.Fprintf(psess, s, "%#v", n.Op)
		if n.Left != nil && n.Left.Op == n.Op {
			fmt.Fprint(s, " ")
		}
		n.Left.exprfmt(psess, s, nprec+1, mode)

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
		n.Left.exprfmt(psess, s, nprec, mode)
		mode.Fprintf(psess, s, " %#v ", n.Op)
		n.Right.exprfmt(psess, s, nprec+1, mode)

	case OADDSTR:
		for i, n1 := range n.List.Slice() {
			if i != 0 {
				fmt.Fprint(s, " + ")
			}
			n1.exprfmt(psess, s, nprec, mode)
		}

	case OCMPSTR, OCMPIFACE:
		n.Left.exprfmt(psess, s, nprec, mode)
		mode.Fprintf(psess, s, " %#v ", n.SubOp(psess))
		n.Right.exprfmt(psess, s, nprec+1, mode)

	default:
		mode.Fprintf(psess, s, "<node %v>", n.Op)
	}
}

func (n *Node) nodefmt(psess *PackageSession, s fmt.State, flag FmtFlag, mode fmtMode) {
	t := n.Type

	if n.Op != OLITERAL && n.Orig != nil {
		n = n.Orig
	}

	if flag&FmtLong != 0 && t != nil {
		if t.Etype == TNIL {
			fmt.Fprint(s, "nil")
		} else {
			mode.Fprintf(psess, s, "%v (type %v)", n, t)
		}
		return
	}

	if psess.opprec[n.Op] < 0 {
		n.stmtfmt(psess, s, mode)
		return
	}

	n.exprfmt(psess, s, 0, mode)
}

func (n *Node) nodedump(psess *PackageSession, s fmt.State, flag FmtFlag, mode fmtMode) {
	recur := flag&FmtShort == 0

	if recur {
		psess.
			indent(s)
		if psess.dumpdepth > 40 {
			fmt.Fprint(s, "...")
			return
		}

		if n.Ninit.Len() != 0 {
			mode.Fprintf(psess, s, "%v-init%v", n.Op, n.Ninit)
			psess.
				indent(s)
		}
	}

	switch n.Op {
	default:
		mode.Fprintf(psess, s, "%v%j", n.Op, n)

	case OINDREGSP:
		mode.Fprintf(psess, s, "%v-SP%j", n.Op, n)

	case OLITERAL:
		mode.Fprintf(psess, s, "%v-%v%j", n.Op, n.Val(), n)

	case ONAME, ONONAME:
		if n.Sym != nil {
			mode.Fprintf(psess, s, "%v-%v%j", n.Op, n.Sym, n)
		} else {
			mode.Fprintf(psess, s, "%v%j", n.Op, n)
		}
		if recur && n.Type == nil && n.Name != nil && n.Name.Param != nil && n.Name.Param.Ntype != nil {
			psess.
				indent(s)
			mode.Fprintf(psess, s, "%v-ntype%v", n.Op, n.Name.Param.Ntype)
		}

	case OASOP:
		mode.Fprintf(psess, s, "%v-%v%j", n.Op, n.SubOp(psess), n)

	case OTYPE:
		mode.Fprintf(psess, s, "%v %v%j type=%v", n.Op, n.Sym, n, n.Type)
		if recur && n.Type == nil && n.Name != nil && n.Name.Param != nil && n.Name.Param.Ntype != nil {
			psess.
				indent(s)
			mode.Fprintf(psess, s, "%v-ntype%v", n.Op, n.Name.Param.Ntype)
		}
	}

	if n.Sym != nil && n.Op != ONAME {
		mode.Fprintf(psess, s, " %v", n.Sym)
	}

	if n.Type != nil {
		mode.Fprintf(psess, s, " %v", n.Type)
	}

	if recur {
		if n.Left != nil {
			mode.Fprintf(psess, s, "%v", n.Left)
		}
		if n.Right != nil {
			mode.Fprintf(psess, s, "%v", n.Right)
		}
		if n.List.Len() != 0 {
			psess.
				indent(s)
			mode.Fprintf(psess, s, "%v-list%v", n.Op, n.List)
		}

		if n.Rlist.Len() != 0 {
			psess.
				indent(s)
			mode.Fprintf(psess, s, "%v-rlist%v", n.Op, n.Rlist)
		}

		if n.Nbody.Len() != 0 {
			psess.
				indent(s)
			mode.Fprintf(psess, s, "%v-body%v", n.Op, n.Nbody)
		}
	}
}

// "%S" suppresses qualifying with package
func (psess *PackageSession) symFormat(s *types.Sym, f fmt.State, verb rune, mode fmtMode) {
	switch verb {
	case 'v', 'S':
		fmt.Fprint(f, psess.sconv(s, psess.fmtFlag(f, verb), mode))

	default:
		fmt.Fprintf(f, "%%!%c(*types.Sym=%p)", verb, s)
	}
}

func (psess *PackageSession) smodeString(s *types.Sym, mode fmtMode) string {
	return psess.sconv(s, 0, mode)
}

// See #16897 before changing the implementation of sconv.
func (psess *PackageSession) sconv(s *types.Sym, flag FmtFlag, mode fmtMode) string {
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
	return psess.symfmt(s, flag, mode)
}

func (psess *PackageSession) tmodeString(t *types.Type, mode fmtMode, depth int) string {
	return psess.tconv(t, 0, mode, depth)
}

func (psess *PackageSession) fldconv(f *types.Field, flag FmtFlag, mode fmtMode, depth int, funarg types.Funarg) string {
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

		if mode == FErr {
			s = psess.origSym(s)
		}

		if s != nil && f.Embedded == 0 {
			if funarg != types.FunargNone {
				name = asNode(f.Nname).modeString(psess, mode)
			} else if flag&FmtLong != 0 {
				name = mode.Sprintf(psess, "%0S", s)
				if !types.IsExported(name) && flag&FmtUnsigned == 0 {
					name = psess.smodeString(s, mode)
				}
			} else {
				name = psess.smodeString(s, mode)
			}
		}
	}

	var typ string
	if f.Isddd() {
		var et *types.Type
		if f.Type != nil {
			et = f.Type.Elem(psess.types)
		}
		typ = "..." + psess.tmodeString(et, mode, depth)
	} else {
		typ = psess.tmodeString(f.Type, mode, depth)
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
func (psess *PackageSession) typeFormat(t *types.Type, s fmt.State, verb rune, mode fmtMode) {
	switch verb {
	case 'v', 'S', 'L':

		fmt.Fprint(s, psess.tconv(t, psess.fmtFlag(s, verb), mode, 0))

	default:
		fmt.Fprintf(s, "%%!%c(*Type=%p)", verb, t)
	}
}

// See #16897 before changing the implementation of tconv.
func (psess *PackageSession) tconv(t *types.Type, flag FmtFlag, mode fmtMode, depth int) string {
	if t == nil {
		return "<T>"
	}
	if t.Etype == types.TSSA {
		return t.Extra.(string)
	}
	if t.Etype == types.TTUPLE {
		return t.FieldType(psess.types, 0).String(psess.types) + "," + t.FieldType(psess.types, 1).String(psess.types)
	}

	if depth > 100 {
		return "<...>"
	}

	flag, mode = flag.update(mode)
	if mode == FTypeIdName {
		flag |= FmtUnsigned
	}

	str := psess.typefmt(t, flag, mode, depth+1)

	return str
}

func (n *Node) String() string                                        { return fmt.Sprint(n) }
func (n *Node) modeString(psess *PackageSession, mode fmtMode) string { return mode.Sprint(psess, n) }

// "%L"  suffix with "(type %T)" where possible
// "%+S" in debug mode, don't recurse, no multiline output
func (n *Node) nconv(psess *PackageSession, s fmt.State, flag FmtFlag, mode fmtMode) {
	if n == nil {
		fmt.Fprint(s, "<N>")
		return
	}

	flag, mode = flag.update(mode)

	switch mode {
	case FErr:
		n.nodefmt(psess, s, flag, mode)

	case FDbg:
		psess.
			dumpdepth++
		n.nodedump(psess, s, flag, mode)
		psess.
			dumpdepth--

	default:
		psess.
			Fatalf("unhandled %%N mode: %d", mode)
	}
}

func (l Nodes) format(psess *PackageSession, s fmt.State, verb rune, mode fmtMode) {
	switch verb {
	case 'v':
		l.hconv(psess, s, psess.fmtFlag(s, verb), mode)

	default:
		fmt.Fprintf(s, "%%!%c(Nodes)", verb)
	}
}

func (n Nodes) String() string {
	return fmt.Sprint(n)
}

// Flags: all those of %N plus '.': separate with comma's instead of semicolons.
func (l Nodes) hconv(psess *PackageSession, s fmt.State, flag FmtFlag, mode fmtMode) {
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
		fmt.Fprint(s, n.modeString(psess, mode))
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

// TODO(gri) make variable local somehow

// indent prints indentation to s.
func (psess *PackageSession) indent(s fmt.State) {
	fmt.Fprint(s, "\n")
	for i := 0; i < psess.dumpdepth; i++ {
		fmt.Fprint(s, ".   ")
	}
}

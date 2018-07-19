package gc

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"

	"github.com/dave/golib/src/cmd/internal/src"
	"os"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"

	"unicode"
	"unicode/utf8"
)

type Error struct {
	pos src.XPos
	msg string
}

// protects largeStackFrames
// positions of functions whose stack frames are too large (rare)

func (psess *PackageSession) errorexit() {
	psess.
		flusherrors()
	if psess.outfile != "" {
		os.Remove(psess.outfile)
	}
	os.Exit(2)
}

func (psess *PackageSession) adderrorname(n *Node) {
	if n.Op != ODOT {
		return
	}
	old := fmt.Sprintf("%v: undefined: %v\n", n.Line(psess), n.Left)
	if len(psess.errors) > 0 && psess.errors[len(psess.errors)-1].pos.Line() == n.Pos.Line() && psess.errors[len(psess.errors)-1].msg == old {
		psess.
			errors[len(psess.errors)-1].msg = fmt.Sprintf("%v: undefined: %v in %v\n", n.Line(psess), n.Left, n)
	}
}

func (psess *PackageSession) adderr(pos src.XPos, format string, args ...interface{}) {
	psess.
		errors = append(psess.errors, Error{
		pos: pos,
		msg: fmt.Sprintf("%v: %s\n", psess.linestr(pos), fmt.Sprintf(format, args...)),
	})
}

// byPos sorts errors by source position.
type byPos []Error

func (x byPos) Len() int           { return len(x) }
func (x byPos) Less(i, j int) bool { return x[i].pos.Before(x[j].pos) }
func (x byPos) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// flusherrors sorts errors seen so far by line number, prints them to stdout,
// and empties the errors array.
func (psess *PackageSession) flusherrors() {
	psess.
		Ctxt.Bso.Flush()
	if len(psess.errors) == 0 {
		return
	}
	sort.Stable(byPos(psess.errors))
	for i, err := range psess.errors {
		if i == 0 || err.msg != psess.errors[i-1].msg {
			fmt.Printf("%s", err.msg)
		}
	}
	psess.
		errors = psess.errors[:0]
}

func (psess *PackageSession) hcrash() {
	if psess.Debug['h'] != 0 {
		psess.
			flusherrors()
		if psess.outfile != "" {
			os.Remove(psess.outfile)
		}
		var x *int
		*x = 0
	}
}

func (psess *PackageSession) linestr(pos src.XPos) string {
	return psess.Ctxt.OutermostPos(pos).Format(psess.src, psess.Debug['C'] == 0, psess.Debug['L'] == 1)
}

// lasterror keeps track of the most recently issued error.
// It is used to avoid multiple error messages on the same
// line.

// source position of last syntax error
// source position of last non-syntax error
// error message of last non-syntax error

// sameline reports whether two positions a, b are on the same line.
func (psess *PackageSession) sameline(a, b src.XPos) bool {
	p := psess.Ctxt.PosTable.Pos(a)
	q := psess.Ctxt.PosTable.Pos(b)
	return p.Base() == q.Base() && p.Line() == q.Line()
}

func (psess *PackageSession) yyerrorl(pos src.XPos, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	if strings.HasPrefix(msg, "syntax error") {
		psess.
			nsyntaxerrors++

		if psess.sameline(psess.lasterror.syntax, pos) {
			return
		}
		psess.
			lasterror.syntax = pos
	} else {

		if psess.sameline(psess.lasterror.other, pos) && psess.lasterror.msg == msg {
			return
		}
		psess.
			lasterror.other = pos
		psess.
			lasterror.msg = msg
	}
	psess.
		adderr(pos, "%s", msg)
	psess.
		hcrash()
	psess.
		nerrors++
	if psess.nsavederrors+psess.nerrors >= 10 && psess.Debug['e'] == 0 {
		psess.
			flusherrors()
		fmt.Printf("%v: too many errors\n", psess.linestr(pos))
		psess.
			errorexit()
	}
}

func (psess *PackageSession) yyerror(format string, args ...interface{}) {
	psess.
		yyerrorl(psess.lineno, format, args...)
}

func (psess *PackageSession) Warn(fmt_ string, args ...interface{}) {
	psess.
		adderr(psess.lineno, fmt_, args...)
	psess.
		hcrash()
}

func (psess *PackageSession) Warnl(line src.XPos, fmt_ string, args ...interface{}) {
	psess.
		adderr(line, fmt_, args...)
	if psess.Debug['m'] != 0 {
		psess.
			flusherrors()
	}
}

func (psess *PackageSession) Fatalf(fmt_ string, args ...interface{}) {
	psess.
		flusherrors()

	if psess.Debug_panic != 0 || psess.nsavederrors+psess.nerrors == 0 {
		fmt.Printf("%v: internal compiler error: ", psess.linestr(psess.lineno))
		fmt.Printf(fmt_, args...)
		fmt.Printf("\n")

		if strings.HasPrefix(psess.objabi.Version, "go") {
			fmt.Printf("\n")
			fmt.Printf("Please file a bug report including a short program that triggers the error.\n")
			fmt.Printf("https://golang.org/issue/new\n")
		} else {

			fmt.Println()
			os.Stdout.Write(debug.Stack())
			fmt.Println()
		}
	}
	psess.
		hcrash()
	psess.
		errorexit()
}

func (psess *PackageSession) setlineno(n *Node) src.XPos {
	lno := psess.lineno
	if n != nil {
		switch n.Op {
		case ONAME, OPACK:
			break

		case OLITERAL, OTYPE:
			if n.Sym != nil {
				break
			}
			fallthrough

		default:
			psess.
				lineno = n.Pos
			if !psess.lineno.IsKnown() {
				if psess.Debug['K'] != 0 {
					psess.
						Warn("setlineno: unknown position (line 0)")
				}
				psess.
					lineno = lno
			}
		}
	}

	return lno
}

func (psess *PackageSession) lookup(name string) *types.Sym {
	return psess.localpkg.Lookup(psess.types, name)
}

// lookupN looks up the symbol starting with prefix and ending with
// the decimal n. If prefix is too long, lookupN panics.
func (psess *PackageSession) lookupN(prefix string, n int) *types.Sym {
	var buf [20]byte // plenty long enough for all current users
	copy(buf[:], prefix)
	b := strconv.AppendInt(buf[:len(prefix)], int64(n), 10)
	return psess.localpkg.LookupBytes(psess.types, b)
}

// autolabel generates a new Name node for use with
// an automatically generated label.
// prefix is a short mnemonic (e.g. ".s" for switch)
// to help with debugging.
// It should begin with "." to avoid conflicts with
// user labels.
func (psess *PackageSession) autolabel(prefix string) *Node {
	if prefix[0] != '.' {
		psess.
			Fatalf("autolabel prefix must start with '.', have %q", prefix)
	}
	fn := psess.Curfn
	if psess.Curfn == nil {
		psess.
			Fatalf("autolabel outside function")
	}
	n := fn.Func.Label
	fn.Func.Label++
	return psess.newname(psess.lookupN(prefix, int(n)))
}

func (psess *PackageSession) restrictlookup(name string, pkg *types.Pkg) *types.Sym {
	if !types.IsExported(name) && pkg != psess.localpkg {
		psess.
			yyerror("cannot refer to unexported name %s.%s", pkg.Name, name)
	}
	return pkg.Lookup(psess.types, name)
}

// find all the exported symbols in package opkg
// and make them available in the current package
func (psess *PackageSession) importdot(opkg *types.Pkg, pack *Node) {
	n := 0
	for _, s := range opkg.Syms {
		if s.Def == nil {
			continue
		}
		if !types.IsExported(s.Name) || strings.ContainsRune(s.Name, 0xb7) {
			continue
		}
		s1 := psess.lookup(s.Name)
		if s1.Def != nil {
			pkgerror := fmt.Sprintf("during import %q", opkg.Path)
			psess.
				redeclare(psess.lineno, s1, pkgerror)
			continue
		}

		s1.Def = s.Def
		s1.Block = s.Block
		if asNode(s1.Def).Name == nil {
			Dump("s1def", asNode(s1.Def))
			psess.
				Fatalf("missing Name")
		}
		asNode(s1.Def).Name.Pack = pack
		s1.Origpkg = opkg
		n++
	}

	if n == 0 {
		psess.
			yyerrorl(pack.Pos, "imported and not used: %q", opkg.Path)
	}
}

func (psess *PackageSession) nod(op Op, nleft, nright *Node) *Node {
	return psess.nodl(psess.lineno, op, nleft, nright)
}

func (psess *PackageSession) nodl(pos src.XPos, op Op, nleft, nright *Node) *Node {
	var n *Node
	switch op {
	case OCLOSURE, ODCLFUNC:
		var x struct {
			Node
			Func
		}
		n = &x.Node
		n.Func = &x.Func
	case ONAME:
		psess.
			Fatalf("use newname instead")
	case OLABEL, OPACK:
		var x struct {
			Node
			Name
		}
		n = &x.Node
		n.Name = &x.Name
	default:
		n = new(Node)
	}
	n.Op = op
	n.Left = nleft
	n.Right = nright
	n.Pos = pos
	n.Xoffset = BADWIDTH
	n.Orig = n
	return n
}

// newname returns a new ONAME Node associated with symbol s.
func (psess *PackageSession) newname(s *types.Sym) *Node {
	n := psess.newnamel(psess.lineno, s)
	n.Name.Curfn = psess.Curfn
	return n
}

// newname returns a new ONAME Node associated with symbol s at position pos.
// The caller is responsible for setting n.Name.Curfn.
func (psess *PackageSession) newnamel(pos src.XPos, s *types.Sym) *Node {
	if s == nil {
		psess.
			Fatalf("newnamel nil")
	}

	var x struct {
		Node
		Name
		Param
	}
	n := &x.Node
	n.Name = &x.Name
	n.Name.Param = &x.Param

	n.Op = ONAME
	n.Pos = pos
	n.Orig = n

	n.Sym = s
	n.SetAddable(true)
	return n
}

// nodSym makes a Node with Op op and with the Left field set to left
// and the Sym field set to sym. This is for ODOT and friends.
func (psess *PackageSession) nodSym(op Op, left *Node, sym *types.Sym) *Node {
	n := psess.nod(op, left, nil)
	n.Sym = sym
	return n
}

func (n *Node) copy() *Node {
	n2 := *n
	return &n2
}

// methcmp sorts methods by symbol.
type methcmp []*types.Field

func (x methcmp) Len() int           { return len(x) }
func (x methcmp) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x methcmp) Less(i, j int) bool { return x[i].Sym.Less(x[j].Sym) }

func (psess *PackageSession) nodintconst(v int64) *Node {
	u := new(Mpint)
	u.SetInt64(v)
	return psess.nodlit(Val{u})
}

func (psess *PackageSession) nodfltconst(v *Mpflt) *Node {
	u := newMpflt()
	u.Set(v)
	return psess.nodlit(Val{u})
}

func (psess *PackageSession) nodnil() *Node {
	return psess.nodlit(Val{new(NilVal)})
}

func (psess *PackageSession) nodbool(b bool) *Node {
	return psess.nodlit(Val{b})
}

func (psess *PackageSession) nodstr(s string) *Node {
	return psess.nodlit(Val{s})
}

// treecopy recursively copies n, with the exception of
// ONAME, OLITERAL, OTYPE, and non-iota ONONAME leaves.
// Copies of iota ONONAME nodes are assigned the current
// value of iota_. If pos.IsKnown(), it sets the source
// position of newly allocated nodes to pos.
func (psess *PackageSession) treecopy(n *Node, pos src.XPos) *Node {
	if n == nil {
		return nil
	}

	switch n.Op {
	default:
		m := n.copy()
		m.Orig = m
		m.Left = psess.treecopy(n.Left, pos)
		m.Right = psess.treecopy(n.Right, pos)
		m.List.Set(psess.listtreecopy(n.List.Slice(), pos))
		if pos.IsKnown() {
			m.Pos = pos
		}
		if m.Name != nil && n.Op != ODCLFIELD {
			Dump("treecopy", n)
			psess.
				Fatalf("treecopy Name")
		}
		return m

	case OPACK:

		fallthrough

	case ONAME, ONONAME, OLITERAL, OTYPE:
		return n

	}
}

// isNil reports whether n represents the universal untyped zero value "nil".
func (n *Node) isNil(psess *PackageSession) bool {

	return psess.Isconst(n.Orig, CTNIL)
}

func (psess *PackageSession) isptrto(t *types.Type, et types.EType) bool {
	if t == nil {
		return false
	}
	if !t.IsPtr() {
		return false
	}
	t = t.Elem(psess.types)
	if t == nil {
		return false
	}
	if t.Etype != et {
		return false
	}
	return true
}

func (n *Node) isBlank() bool {
	if n == nil {
		return false
	}
	return n.Sym.IsBlank()
}

// methtype returns the underlying type, if any,
// that owns methods with receiver parameter t.
// The result is either a named type or an anonymous struct.
func (psess *PackageSession) methtype(t *types.Type) *types.Type {
	if t == nil {
		return nil
	}

	if t.IsPtr() {
		if t.Sym != nil {
			return nil
		}
		t = t.Elem(psess.types)
		if t == nil {
			return nil
		}
	}

	if t.Sym == nil && !t.IsStruct() {
		return nil
	}

	if psess.issimple[t.Etype] {
		return t
	}
	switch t.Etype {
	case TARRAY, TCHAN, TFUNC, TMAP, TSLICE, TSTRING, TSTRUCT:
		return t
	}
	return nil
}

// eqtype reports whether t1 and t2 are identical, following the spec rules.
//
// Any cyclic type must go through a named type, and if one is
// named, it is only identical to the other if they are the same
// pointer (t1 == t2), so there's no chance of chasing cycles
// ad infinitum, so no need for a depth counter.
func (psess *PackageSession) eqtype(t1, t2 *types.Type) bool {
	return psess.eqtype1(t1, t2, true, nil)
}

// eqtypeIgnoreTags is like eqtype but it ignores struct tags for struct identity.
func (psess *PackageSession) eqtypeIgnoreTags(t1, t2 *types.Type) bool {
	return psess.eqtype1(t1, t2, false, nil)
}

type typePair struct {
	t1 *types.Type
	t2 *types.Type
}

func (psess *PackageSession) eqtype1(t1, t2 *types.Type, cmpTags bool, assumedEqual map[typePair]struct{}) bool {
	if t1 == t2 {
		return true
	}
	if t1 == nil || t2 == nil || t1.Etype != t2.Etype || t1.Broke() || t2.Broke() {
		return false
	}
	if t1.Sym != nil || t2.Sym != nil {

		switch t1.Etype {
		case TUINT8:
			return (t1 == psess.types.Types[TUINT8] || t1 == psess.types.Bytetype) && (t2 == psess.types.Types[TUINT8] || t2 == psess.types.Bytetype)
		case TINT32:
			return (t1 == psess.types.Types[TINT32] || t1 == psess.types.Runetype) && (t2 == psess.types.Types[TINT32] || t2 == psess.types.Runetype)
		default:
			return false
		}
	}

	if assumedEqual == nil {
		assumedEqual = make(map[typePair]struct{})
	} else if _, ok := assumedEqual[typePair{t1, t2}]; ok {
		return true
	}
	assumedEqual[typePair{t1, t2}] = struct{}{}

	switch t1.Etype {
	case TINTER:
		if t1.NumFields(psess.types) != t2.NumFields(psess.types) {
			return false
		}
		for i, f1 := range t1.FieldSlice(psess.types) {
			f2 := t2.Field(psess.types, i)
			if f1.Sym != f2.Sym || !psess.eqtype1(f1.Type, f2.Type, cmpTags, assumedEqual) {
				return false
			}
		}
		return true

	case TSTRUCT:
		if t1.NumFields(psess.types) != t2.NumFields(psess.types) {
			return false
		}
		for i, f1 := range t1.FieldSlice(psess.types) {
			f2 := t2.Field(psess.types, i)
			if f1.Sym != f2.Sym || f1.Embedded != f2.Embedded || !psess.eqtype1(f1.Type, f2.Type, cmpTags, assumedEqual) {
				return false
			}
			if cmpTags && f1.Note != f2.Note {
				return false
			}
		}
		return true

	case TFUNC:

		for _, f := range psess.types.ParamsResults {

			fs1, fs2 := f(t1).FieldSlice(psess.types), f(t2).FieldSlice(psess.types)
			if len(fs1) != len(fs2) {
				return false
			}
			for i, f1 := range fs1 {
				f2 := fs2[i]
				if f1.Isddd() != f2.Isddd() || !psess.eqtype1(f1.Type, f2.Type, cmpTags, assumedEqual) {
					return false
				}
			}
		}
		return true

	case TARRAY:
		if t1.NumElem(psess.types) != t2.NumElem(psess.types) {
			return false
		}

	case TCHAN:
		if t1.ChanDir(psess.types) != t2.ChanDir(psess.types) {
			return false
		}

	case TMAP:
		if !psess.eqtype1(t1.Key(psess.types), t2.Key(psess.types), cmpTags, assumedEqual) {
			return false
		}
	}

	return psess.eqtype1(t1.Elem(psess.types), t2.Elem(psess.types), cmpTags, assumedEqual)
}

// Are t1 and t2 equal struct types when field names are ignored?
// For deciding whether the result struct from g can be copied
// directly when compiling f(g()).
func (psess *PackageSession) eqtypenoname(t1 *types.Type, t2 *types.Type) bool {
	if t1 == nil || t2 == nil || !t1.IsStruct() || !t2.IsStruct() {
		return false
	}

	if t1.NumFields(psess.types) != t2.NumFields(psess.types) {
		return false
	}
	for i, f1 := range t1.FieldSlice(psess.types) {
		f2 := t2.Field(psess.types, i)
		if !psess.eqtype(f1.Type, f2.Type) {
			return false
		}
	}
	return true
}

// Is type src assignment compatible to type dst?
// If so, return op code to use in conversion.
// If not, return 0.
func (psess *PackageSession) assignop(src *types.Type, dst *types.Type, why *string) Op {
	if why != nil {
		*why = ""
	}

	if psess.safemode && !psess.inimport && src != nil && src.Etype == TUNSAFEPTR {
		psess.
			yyerror("cannot use unsafe.Pointer")
		psess.
			errorexit()
	}

	if src == dst {
		return OCONVNOP
	}
	if src == nil || dst == nil || src.Etype == TFORW || dst.Etype == TFORW || src.Orig == nil || dst.Orig == nil {
		return 0
	}

	if psess.eqtype(src, dst) {
		return OCONVNOP
	}

	if psess.eqtype(src.Orig, dst.Orig) {
		if src.IsEmptyInterface(psess.types) {

			return OCONVNOP
		}
		if (src.Sym == nil || dst.Sym == nil) && !src.IsInterface() {

			return OCONVNOP
		}
	}

	if dst.IsInterface() && src.Etype != TNIL {
		var missing, have *types.Field
		var ptr int
		if psess.implements(src, dst, &missing, &have, &ptr) {
			return OCONVIFACE
		}

		if have != nil && have.Sym == missing.Sym && (have.Type.Broke() || missing.Type.Broke()) {
			return OCONVIFACE
		}

		if why != nil {
			if psess.isptrto(src, TINTER) {
				*why = fmt.Sprintf(":\n\t%v is pointer to interface, not interface", src)
			} else if have != nil && have.Sym == missing.Sym && have.Nointerface() {
				*why = fmt.Sprintf(":\n\t%v does not implement %v (%v method is marked 'nointerface')", src, dst, missing.Sym)
			} else if have != nil && have.Sym == missing.Sym {
				*why = fmt.Sprintf(":\n\t%v does not implement %v (wrong type for %v method)\n"+
					"\t\thave %v%0S\n\t\twant %v%0S", src, dst, missing.Sym, have.Sym, have.Type, missing.Sym, missing.Type)
			} else if ptr != 0 {
				*why = fmt.Sprintf(":\n\t%v does not implement %v (%v method has pointer receiver)", src, dst, missing.Sym)
			} else if have != nil {
				*why = fmt.Sprintf(":\n\t%v does not implement %v (missing %v method)\n"+
					"\t\thave %v%0S\n\t\twant %v%0S", src, dst, missing.Sym, have.Sym, have.Type, missing.Sym, missing.Type)
			} else {
				*why = fmt.Sprintf(":\n\t%v does not implement %v (missing %v method)", src, dst, missing.Sym)
			}
		}

		return 0
	}

	if psess.isptrto(dst, TINTER) {
		if why != nil {
			*why = fmt.Sprintf(":\n\t%v is pointer to interface, not interface", dst)
		}
		return 0
	}

	if src.IsInterface() && dst.Etype != TBLANK {
		var missing, have *types.Field
		var ptr int
		if why != nil && psess.implements(dst, src, &missing, &have, &ptr) {
			*why = ": need type assertion"
		}
		return 0
	}

	if src.IsChan() && src.ChanDir(psess.types) == types.Cboth && dst.IsChan() {
		if psess.eqtype(src.Elem(psess.types), dst.Elem(psess.types)) && (src.Sym == nil || dst.Sym == nil) {
			return OCONVNOP
		}
	}

	if src.Etype == TNIL {
		switch dst.Etype {
		case TPTR32,
			TPTR64,
			TFUNC,
			TMAP,
			TCHAN,
			TINTER,
			TSLICE:
			return OCONVNOP
		}
	}

	if dst.Etype == TBLANK {
		return OCONVNOP
	}

	return 0
}

// Can we convert a value of type src to a value of type dst?
// If so, return op code to use in conversion (maybe OCONVNOP).
// If not, return 0.
func (psess *PackageSession) convertop(src *types.Type, dst *types.Type, why *string) Op {
	if why != nil {
		*why = ""
	}

	if src == dst {
		return OCONVNOP
	}
	if src == nil || dst == nil {
		return 0
	}

	if src.IsPtr() && dst.IsPtr() && dst.Elem(psess.types).NotInHeap() && !src.Elem(psess.types).NotInHeap() {
		if why != nil {
			*why = fmt.Sprintf(":\n\t%v is go:notinheap, but %v is not", dst.Elem(psess.types), src.Elem(psess.types))
		}
		return 0
	}

	op := psess.assignop(src, dst, why)
	if op != 0 {
		return op
	}

	if src.IsInterface() || dst.IsInterface() {
		return 0
	}
	if why != nil {
		*why = ""
	}

	if psess.eqtypeIgnoreTags(src.Orig, dst.Orig) {
		return OCONVNOP
	}

	if src.IsPtr() && dst.IsPtr() && src.Sym == nil && dst.Sym == nil {
		if psess.eqtypeIgnoreTags(src.Elem(psess.types).Orig, dst.Elem(psess.types).Orig) {
			return OCONVNOP
		}
	}

	if (src.IsInteger() || src.IsFloat()) && (dst.IsInteger() || dst.IsFloat()) {
		if psess.simtype[src.Etype] == psess.simtype[dst.Etype] {
			return OCONVNOP
		}
		return OCONV
	}

	if src.IsComplex() && dst.IsComplex() {
		if psess.simtype[src.Etype] == psess.simtype[dst.Etype] {
			return OCONVNOP
		}
		return OCONV
	}

	if src.IsInteger() && dst.IsString() {
		return ORUNESTR
	}

	if src.IsSlice() && dst.IsString() {
		if src.Elem(psess.types).Etype == psess.types.Bytetype.Etype {
			return OARRAYBYTESTR
		}
		if src.Elem(psess.types).Etype == psess.types.Runetype.Etype {
			return OARRAYRUNESTR
		}
	}

	if src.IsString() && dst.IsSlice() {
		if dst.Elem(psess.types).Etype == psess.types.Bytetype.Etype {
			return OSTRARRAYBYTE
		}
		if dst.Elem(psess.types).Etype == psess.types.Runetype.Etype {
			return OSTRARRAYRUNE
		}
	}

	if (src.IsPtr() || src.Etype == TUINTPTR) && dst.Etype == TUNSAFEPTR {
		return OCONVNOP
	}

	if src.Etype == TUNSAFEPTR && (dst.IsPtr() || dst.Etype == TUINTPTR) {
		return OCONVNOP
	}

	if src.Etype == TMAP && dst.IsPtr() &&
		src.MapType(psess.types).Hmap == dst.Elem(psess.types) {
		return OCONVNOP
	}

	return 0
}

func (psess *PackageSession) assignconv(n *Node, t *types.Type, context string) *Node {
	return psess.assignconvfn(n, t, func() string { return context })
}

// Convert node n for assignment to type t.
func (psess *PackageSession) assignconvfn(n *Node, t *types.Type, context func() string) *Node {
	if n == nil || n.Type == nil || n.Type.Broke() {
		return n
	}

	if t.Etype == TBLANK && n.Type.Etype == TNIL {
		psess.
			yyerror("use of untyped nil")
	}

	old := n
	od := old.Diag()
	old.SetDiag(true)
	n = psess.defaultlit(n, t)
	old.SetDiag(od)
	if t.Etype == TBLANK {
		return n
	}

	if n.Type == psess.types.Idealbool && !t.IsBoolean() {
		if n.Op == ONAME || n.Op == OLITERAL {
			r := psess.nod(OCONVNOP, n, nil)
			r.Type = psess.types.Types[TBOOL]
			r.SetTypecheck(1)
			r.SetImplicit(true)
			n = r
		}
	}

	if psess.eqtype(n.Type, t) {
		return n
	}

	var why string
	op := psess.assignop(n.Type, t, &why)
	if op == 0 {
		if !old.Diag() {
			psess.
				yyerror("cannot use %L as type %v in %s%s", n, t, context(), why)
		}
		op = OCONV
	}

	r := psess.nod(op, n, nil)
	r.Type = t
	r.SetTypecheck(1)
	r.SetImplicit(true)
	r.Orig = n.Orig
	return r
}

// IsMethod reports whether n is a method.
// n must be a function or a method.
func (n *Node) IsMethod(psess *PackageSession) bool {
	return n.Type.Recv(psess.types) != nil
}

// SliceBounds returns n's slice bounds: low, high, and max in expr[low:high:max].
// n must be a slice expression. max is nil if n is a simple slice expression.
func (n *Node) SliceBounds(psess *PackageSession) (low, high, max *Node) {
	if n.List.Len() == 0 {
		return nil, nil, nil
	}

	switch n.Op {
	case OSLICE, OSLICEARR, OSLICESTR:
		s := n.List.Slice()
		return s[0], s[1], nil
	case OSLICE3, OSLICE3ARR:
		s := n.List.Slice()
		return s[0], s[1], s[2]
	}
	psess.
		Fatalf("SliceBounds op %v: %v", n.Op, n)
	return nil, nil, nil
}

// SetSliceBounds sets n's slice bounds, where n is a slice expression.
// n must be a slice expression. If max is non-nil, n must be a full slice expression.
func (n *Node) SetSliceBounds(psess *PackageSession, low, high, max *Node) {
	switch n.Op {
	case OSLICE, OSLICEARR, OSLICESTR:
		if max != nil {
			psess.
				Fatalf("SetSliceBounds %v given three bounds", n.Op)
		}
		s := n.List.Slice()
		if s == nil {
			if low == nil && high == nil {
				return
			}
			n.List.Set2(low, high)
			return
		}
		s[0] = low
		s[1] = high
		return
	case OSLICE3, OSLICE3ARR:
		s := n.List.Slice()
		if s == nil {
			if low == nil && high == nil && max == nil {
				return
			}
			n.List.Set3(low, high, max)
			return
		}
		s[0] = low
		s[1] = high
		s[2] = max
		return
	}
	psess.
		Fatalf("SetSliceBounds op %v: %v", n.Op, n)
}

// IsSlice3 reports whether o is a slice3 op (OSLICE3, OSLICE3ARR).
// o must be a slicing op.
func (o Op) IsSlice3(psess *PackageSession) bool {
	switch o {
	case OSLICE, OSLICEARR, OSLICESTR:
		return false
	case OSLICE3, OSLICE3ARR:
		return true
	}
	psess.
		Fatalf("IsSlice3 op %v", o)
	return false
}

// labeledControl returns the control flow Node (for, switch, select)
// associated with the label n, if any.
func (n *Node) labeledControl(psess *PackageSession) *Node {
	if n.Op != OLABEL {
		psess.
			Fatalf("labeledControl %v", n.Op)
	}
	ctl := n.Name.Defn
	if ctl == nil {
		return nil
	}
	switch ctl.Op {
	case OFOR, OFORUNTIL, OSWITCH, OSELECT:
		return ctl
	}
	return nil
}

func (psess *PackageSession) syslook(name string) *Node {
	s := psess.Runtimepkg.Lookup(psess.types, name)
	if s == nil || s.Def == nil {
		psess.
			Fatalf("syslook: can't find runtime.%s", name)
	}
	return asNode(s.Def)
}

// typehash computes a hash value for type t to use in type switch statements.
func (psess *PackageSession) typehash(t *types.Type) uint32 {
	p := t.LongString(psess.types)

	h := md5.Sum([]byte(p))
	return binary.LittleEndian.Uint32(h[:4])
}

func (psess *PackageSession) frame(context int) {
	if context != 0 {
		fmt.Printf("--- external frame ---\n")
		for _, n := range psess.externdcl {
			printframenode(n)
		}
		return
	}

	if psess.Curfn != nil {
		fmt.Printf("--- %v frame ---\n", psess.Curfn.Func.Nname.Sym)
		for _, ln := range psess.Curfn.Func.Dcl {
			printframenode(ln)
		}
	}
}

func printframenode(n *Node) {
	w := int64(-1)
	if n.Type != nil {
		w = n.Type.Width
	}
	switch n.Op {
	case ONAME:
		fmt.Printf("%v %v G%d %v width=%d\n", n.Op, n.Sym, n.Name.Vargen, n.Type, w)
	case OTYPE:
		fmt.Printf("%v %v width=%d\n", n.Op, n.Type, w)
	}
}

// updateHasCall checks whether expression n contains any function
// calls and sets the n.HasCall flag if so.
func (psess *PackageSession) updateHasCall(n *Node) {
	if n == nil {
		return
	}
	n.SetHasCall(psess.calcHasCall(n))
}

func (psess *PackageSession) calcHasCall(n *Node) bool {
	if n.Ninit.Len() != 0 {

		return true
	}

	switch n.Op {
	case OLITERAL, ONAME, OTYPE:
		if n.HasCall() {
			psess.
				Fatalf("OLITERAL/ONAME/OTYPE should never have calls: %+v", n)
		}
		return false
	case OCALL, OCALLFUNC, OCALLMETH, OCALLINTER:
		return true
	case OANDAND, OOROR:

		if psess.instrumenting {
			return true
		}
	case OINDEX, OSLICE, OSLICEARR, OSLICE3, OSLICE3ARR, OSLICESTR,
		OIND, ODOTPTR, ODOTTYPE, ODIV, OMOD:

		return true

	case OADD, OSUB, OMINUS:
		if psess.thearch.SoftFloat && (psess.isFloat[n.Type.Etype] || psess.isComplex[n.Type.Etype]) {
			return true
		}
	case OLT, OEQ, ONE, OLE, OGE, OGT:
		if psess.thearch.SoftFloat && (psess.isFloat[n.Left.Type.Etype] || psess.isComplex[n.Left.Type.Etype]) {
			return true
		}
	case OCONV:
		if psess.thearch.SoftFloat && ((psess.isFloat[n.Type.Etype] || psess.isComplex[n.Type.Etype]) || (psess.isFloat[n.Left.Type.Etype] || psess.isComplex[n.Left.Type.Etype])) {
			return true
		}
	}

	if n.Left != nil && n.Left.HasCall() {
		return true
	}
	if n.Right != nil && n.Right.HasCall() {
		return true
	}
	return false
}

func (psess *PackageSession) badtype(op Op, tl *types.Type, tr *types.Type) {
	fmt_ := ""
	if tl != nil {
		fmt_ += fmt.Sprintf("\n\t%v", tl)
	}
	if tr != nil {
		fmt_ += fmt.Sprintf("\n\t%v", tr)
	}

	if tl != nil && tr != nil && tl.IsPtr() && tr.IsPtr() {
		if tl.Elem(psess.types).IsStruct() && tr.Elem(psess.types).IsInterface() {
			fmt_ += "\n\t(*struct vs *interface)"
		} else if tl.Elem(psess.types).IsInterface() && tr.Elem(psess.types).IsStruct() {
			fmt_ += "\n\t(*interface vs *struct)"
		}
	}

	s := fmt_
	psess.
		yyerror("illegal types for operand: %v%s", op, s)
}

// brcom returns !(op).
// For example, brcom(==) is !=.
func (psess *PackageSession) brcom(op Op) Op {
	switch op {
	case OEQ:
		return ONE
	case ONE:
		return OEQ
	case OLT:
		return OGE
	case OGT:
		return OLE
	case OLE:
		return OGT
	case OGE:
		return OLT
	}
	psess.
		Fatalf("brcom: no com for %v\n", op)
	return op
}

// brrev returns reverse(op).
// For example, Brrev(<) is >.
func (psess *PackageSession) brrev(op Op) Op {
	switch op {
	case OEQ:
		return OEQ
	case ONE:
		return ONE
	case OLT:
		return OGT
	case OGT:
		return OLT
	case OLE:
		return OGE
	case OGE:
		return OLE
	}
	psess.
		Fatalf("brrev: no rev for %v\n", op)
	return op
}

// return side effect-free n, appending side effects to init.
// result is assignable if n is.
func (psess *PackageSession) safeexpr(n *Node, init *Nodes) *Node {
	if n == nil {
		return nil
	}

	if n.Ninit.Len() != 0 {
		psess.
			walkstmtlist(n.Ninit.Slice())
		init.AppendNodes(&n.Ninit)
	}

	switch n.Op {
	case ONAME, OLITERAL:
		return n

	case ODOT, OLEN, OCAP:
		l := psess.safeexpr(n.Left, init)
		if l == n.Left {
			return n
		}
		r := n.copy()
		r.Left = l
		r = psess.typecheck(r, Erv)
		r = psess.walkexpr(r, init)
		return r

	case ODOTPTR, OIND:
		l := psess.safeexpr(n.Left, init)
		if l == n.Left {
			return n
		}
		a := n.copy()
		a.Left = l
		a = psess.walkexpr(a, init)
		return a

	case OINDEX, OINDEXMAP:
		l := psess.safeexpr(n.Left, init)
		r := psess.safeexpr(n.Right, init)
		if l == n.Left && r == n.Right {
			return n
		}
		a := n.copy()
		a.Left = l
		a.Right = r
		a = psess.walkexpr(a, init)
		return a

	case OSTRUCTLIT, OARRAYLIT, OSLICELIT:
		if psess.isStaticCompositeLiteral(n) {
			return n
		}
	}

	if islvalue(n) {
		psess.
			Fatalf("missing lvalue case in safeexpr: %v", n)
	}
	return psess.cheapexpr(n, init)
}

func (psess *PackageSession) copyexpr(n *Node, t *types.Type, init *Nodes) *Node {
	l := psess.temp(t)
	a := psess.nod(OAS, l, n)
	a = psess.typecheck(a, Etop)
	a = psess.walkexpr(a, init)
	init.Append(a)
	return l
}

// return side-effect free and cheap n, appending side effects to init.
// result may not be assignable.
func (psess *PackageSession) cheapexpr(n *Node, init *Nodes) *Node {
	switch n.Op {
	case ONAME, OLITERAL:
		return n
	}

	return psess.copyexpr(n, n.Type, init)
}

// A Dlist stores a pointer to a TFIELD Type embedded within
// a TSTRUCT or TINTER Type.
type Dlist struct {
	field *types.Field
}

// dotlist is used by adddot1 to record the path of embedded fields
// used to access a target field or method.
// Must be non-nil so that dotpath returns a non-nil slice even if d is zero.

// lookdot0 returns the number of fields or methods named s associated
// with Type t. If exactly one exists, it will be returned in *save
// (if save is not nil).
func (psess *PackageSession) lookdot0(s *types.Sym, t *types.Type, save **types.Field, ignorecase bool) int {
	u := t
	if u.IsPtr() {
		u = u.Elem(psess.types)
	}

	c := 0
	if u.IsStruct() || u.IsInterface() {
		for _, f := range u.Fields(psess.types).Slice() {
			if f.Sym == s || (ignorecase && f.Type.Etype == TFUNC && f.Type.Recv(psess.types) != nil && strings.EqualFold(f.Sym.Name, s.Name)) {
				if save != nil {
					*save = f
				}
				c++
			}
		}
	}

	u = psess.methtype(t)
	if u != nil {
		for _, f := range u.Methods().Slice() {
			if f.Embedded == 0 && (f.Sym == s || (ignorecase && strings.EqualFold(f.Sym.Name, s.Name))) {
				if save != nil {
					*save = f
				}
				c++
			}
		}
	}

	return c
}

// adddot1 returns the number of fields or methods named s at depth d in Type t.
// If exactly one exists, it will be returned in *save (if save is not nil),
// and dotlist will contain the path of embedded fields traversed to find it,
// in reverse order. If none exist, more will indicate whether t contains any
// embedded fields at depth d, so callers can decide whether to retry at
// a greater depth.
func (psess *PackageSession) adddot1(s *types.Sym, t *types.Type, d int, save **types.Field, ignorecase bool) (c int, more bool) {
	if t.Recur() {
		return
	}
	t.SetRecur(true)
	defer t.SetRecur(false)

	var u *types.Type
	d--
	if d < 0 {

		c = psess.lookdot0(s, t, save, ignorecase)
		if c != 0 {
			return c, false
		}
	}

	u = t
	if u.IsPtr() {
		u = u.Elem(psess.types)
	}
	if !u.IsStruct() && !u.IsInterface() {
		return c, false
	}

	for _, f := range u.Fields(psess.types).Slice() {
		if f.Embedded == 0 || f.Sym == nil {
			continue
		}
		if d < 0 {

			return c, true
		}
		a, more1 := psess.adddot1(s, f.Type, d, save, ignorecase)
		if a != 0 && c == 0 {
			psess.
				dotlist[d].field = f
		}
		c += a
		if more1 {
			more = true
		}
	}

	return c, more
}

// dotpath computes the unique shortest explicit selector path to fully qualify
// a selection expression x.f, where x is of type t and f is the symbol s.
// If no such path exists, dotpath returns nil.
// If there are multiple shortest paths to the same depth, ambig is true.
func (psess *PackageSession) dotpath(s *types.Sym, t *types.Type, save **types.Field, ignorecase bool) (path []Dlist, ambig bool) {

	for d := 0; ; d++ {
		if d > len(psess.dotlist) {
			psess.
				dotlist = append(psess.dotlist, Dlist{})
		}
		if c, more := psess.adddot1(s, t, d, save, ignorecase); c == 1 {
			return psess.dotlist[:d], false
		} else if c > 1 {
			return nil, true
		} else if !more {
			return nil, false
		}
	}
}

// in T.field
// find missing fields that
// will give shortest unique addressing.
// modify the tree with missing type names.
func (psess *PackageSession) adddot(n *Node) *Node {
	n.Left = psess.typecheck(n.Left, Etype|Erv)
	if n.Left.Diag() {
		n.SetDiag(true)
	}
	t := n.Left.Type
	if t == nil {
		return n
	}

	if n.Left.Op == OTYPE {
		return n
	}

	s := n.Sym
	if s == nil {
		return n
	}

	switch path, ambig := psess.dotpath(s, t, nil, false); {
	case path != nil:

		for c := len(path) - 1; c >= 0; c-- {
			n.Left = psess.nodSym(ODOT, n.Left, path[c].field.Sym)
			n.Left.SetImplicit(true)
		}
	case ambig:
		psess.
			yyerror("ambiguous selector %v", n)
		n.Left = nil
	}

	return n
}

type Symlink struct {
	field *types.Field
}

func (psess *PackageSession) expand0(t *types.Type) {
	u := t
	if u.IsPtr() {
		u = u.Elem(psess.types)
	}

	if u.IsInterface() {
		for _, f := range u.Fields(psess.types).Slice() {
			if f.Sym.Uniq() {
				continue
			}
			f.Sym.SetUniq(true)
			psess.
				slist = append(psess.slist, Symlink{field: f})
		}

		return
	}

	u = psess.methtype(t)
	if u != nil {
		for _, f := range u.Methods().Slice() {
			if f.Sym.Uniq() {
				continue
			}
			f.Sym.SetUniq(true)
			psess.
				slist = append(psess.slist, Symlink{field: f})
		}
	}
}

func (psess *PackageSession) expand1(t *types.Type, top bool) {
	if t.Recur() {
		return
	}
	t.SetRecur(true)

	if !top {
		psess.
			expand0(t)
	}

	u := t
	if u.IsPtr() {
		u = u.Elem(psess.types)
	}

	if u.IsStruct() || u.IsInterface() {
		for _, f := range u.Fields(psess.types).Slice() {
			if f.Embedded == 0 {
				continue
			}
			if f.Sym == nil {
				continue
			}
			psess.
				expand1(f.Type, false)
		}
	}

	t.SetRecur(false)
}

func (psess *PackageSession) expandmeth(t *types.Type) {
	if t == nil || t.AllMethods().Len() != 0 {
		return
	}

	for _, f := range t.Methods().Slice() {
		f.Sym.SetUniq(true)
	}
	psess.
		slist = psess.slist[:0]
	psess.
		expand1(t, true)

	// check each method to be uniquely reachable
	var ms []*types.Field
	for i, sl := range psess.slist {
		psess.
			slist[i].field = nil
		sl.field.Sym.SetUniq(false)

		var f *types.Field
		path, _ := psess.dotpath(sl.field.Sym, t, &f, false)
		if path == nil {
			continue
		}

		if f.Type.Etype != TFUNC || f.Type.Recv(psess.types) == nil {
			continue
		}

		f = f.Copy()
		f.Embedded = 1
		for _, d := range path {
			if d.field.Type.IsPtr() {
				f.Embedded = 2
				break
			}
		}
		ms = append(ms, f)
	}

	for _, f := range t.Methods().Slice() {
		f.Sym.SetUniq(false)
	}

	ms = append(ms, t.Methods().Slice()...)
	sort.Sort(methcmp(ms))
	t.AllMethods().Set(ms)
}

// Given funarg struct list, return list of ODCLFIELD Node fn args.
func (psess *PackageSession) structargs(tl *types.Type, mustname bool) []*Node {
	var args []*Node
	gen := 0
	for _, t := range tl.Fields(psess.types).Slice() {
		s := t.Sym
		if mustname && (s == nil || s.Name == "_") {

			s = psess.lookupN(".anon", gen)
			gen++
		}
		a := psess.symfield(s, t.Type)
		a.Pos = t.Pos
		a.SetIsddd(t.Isddd())
		args = append(args, a)
	}

	return args
}

// Generate a wrapper function to convert from
// a receiver of type T to a receiver of type U.
// That is,
//
//	func (t T) M() {
//		...
//	}
//
// already exists; this function generates
//
//	func (u U) M() {
//		u.M()
//	}
//
// where the types T and U are such that u.M() is valid
// and calls the T.M method.
// The resulting function is for use in method tables.
//
//	rcvr - U
//	method - M func (t T)(), a TFIELD type struct
//	newnam - the eventual mangled name of this function
func (psess *PackageSession) genwrapper(rcvr *types.Type, method *types.Field, newnam *types.Sym) {
	if false && psess.Debug['r'] != 0 {
		fmt.Printf("genwrapper rcvrtype=%v method=%v newnam=%v\n", rcvr, method, newnam)
	}

	if rcvr.IsPtr() && rcvr.Elem(psess.types) == method.Type.Recv(psess.types).Type &&
		rcvr.Elem(psess.types).Sym != nil && rcvr.Elem(psess.types).Sym.Pkg != psess.localpkg {
		return
	}

	if rcvr.IsInterface() && rcvr.Sym != nil && rcvr.Sym.Pkg != psess.localpkg {
		return
	}
	psess.
		lineno = psess.autogeneratedPos
	psess.
		dclcontext = PEXTERN

	tfn := psess.nod(OTFUNC, nil, nil)
	tfn.Left = psess.namedfield(".this", rcvr)
	tfn.List.Set(psess.structargs(method.Type.Params(psess.types), true))
	tfn.Rlist.Set(psess.structargs(method.Type.Results(psess.types), false))

	disableExport(newnam)
	fn := psess.dclfunc(newnam, tfn)
	fn.Func.SetDupok(true)

	nthis := asNode(tfn.Type.Recv(psess.types).Nname)

	methodrcvr := method.Type.Recv(psess.types).Type

	if rcvr.IsPtr() && rcvr.Elem(psess.types) == methodrcvr {

		n := psess.nod(OIF, nil, nil)
		n.Left = psess.nod(OEQ, nthis, psess.nodnil())
		call := psess.nod(OCALL, psess.syslook("panicwrap"), nil)
		n.Nbody.Set1(call)
		fn.Nbody.Append(n)
	}

	dot := psess.adddot(psess.nodSym(OXDOT, nthis, method.Sym))

	if !psess.instrumenting && rcvr.IsPtr() && methodrcvr.IsPtr() && method.Embedded != 0 && !psess.isifacemethod(method.Type) && !(psess.thearch.LinkArch.Name == "ppc64le" && psess.Ctxt.Flag_dynlink) {

		dot = dot.Left

		if !psess.dotlist[0].field.Type.IsPtr() {
			dot = psess.nod(OADDR, dot, nil)
		}
		as := psess.nod(OAS, nthis, psess.nod(OCONVNOP, dot, nil))
		as.Right.Type = rcvr
		fn.Nbody.Append(as)
		fn.Nbody.Append(psess.nodSym(ORETJMP, nil, psess.methodSym(methodrcvr, method.Sym)))
	} else {
		fn.Func.SetWrapper(true)
		call := psess.nod(OCALL, dot, nil)
		call.List.Set(psess.paramNnames(tfn.Type))
		call.SetIsddd(tfn.Type.IsVariadic(psess.types))
		if method.Type.NumResults(psess.types) > 0 {
			n := psess.nod(ORETURN, nil, nil)
			n.List.Set1(call)
			call = n
		}
		fn.Nbody.Append(call)
	}

	if false && psess.Debug['r'] != 0 {
		dumplist("genwrapper body", fn.Nbody)
	}
	psess.
		funcbody()
	if psess.debug_dclstack != 0 {
		psess.
			testdclstack()
	}

	fn = psess.typecheck(fn, Etop)
	psess.
		Curfn = fn
	psess.
		typecheckslice(fn.Nbody.Slice(), Etop)

	if false {
		psess.
			inlcalls(fn)
	}
	psess.
		escAnalyze([]*Node{fn}, false)
	psess.
		Curfn = nil
	psess.
		funccompile(fn)
}

func (psess *PackageSession) paramNnames(ft *types.Type) []*Node {
	args := make([]*Node, ft.NumParams(psess.types))
	for i, f := range ft.Params(psess.types).FieldSlice(psess.types) {
		args[i] = asNode(f.Nname)
	}
	return args
}

func (psess *PackageSession) hashmem(t *types.Type) *Node {
	sym := psess.Runtimepkg.Lookup(psess.types, "memhash")

	n := psess.newname(sym)
	n.SetClass(PFUNC)
	n.Type = psess.functype(nil, []*Node{psess.
		anonfield(psess.types.NewPtr(t)), psess.
		anonfield(psess.types.Types[TUINTPTR]), psess.
		anonfield(psess.types.Types[TUINTPTR]),
	}, []*Node{psess.
		anonfield(psess.types.Types[TUINTPTR]),
	})
	return n
}

func (psess *PackageSession) ifacelookdot(s *types.Sym, t *types.Type, ignorecase bool) (m *types.Field, followptr bool) {
	if t == nil {
		return nil, false
	}

	path, ambig := psess.dotpath(s, t, &m, ignorecase)
	if path == nil {
		if ambig {
			psess.
				yyerror("%v.%v is ambiguous", t, s)
		}
		return nil, false
	}

	for _, d := range path {
		if d.field.Type.IsPtr() {
			followptr = true
			break
		}
	}

	if m.Type.Etype != TFUNC || m.Type.Recv(psess.types) == nil {
		psess.
			yyerror("%v.%v is a field, not a method", t, s)
		return nil, followptr
	}

	return m, followptr
}

func (psess *PackageSession) implements(t, iface *types.Type, m, samename **types.Field, ptr *int) bool {
	t0 := t
	if t == nil {
		return false
	}

	if t.IsInterface() {
		i := 0
		tms := t.Fields(psess.types).Slice()
		for _, im := range iface.Fields(psess.types).Slice() {
			for i < len(tms) && tms[i].Sym != im.Sym {
				i++
			}
			if i == len(tms) {
				*m = im
				*samename = nil
				*ptr = 0
				return false
			}
			tm := tms[i]
			if !psess.eqtype(tm.Type, im.Type) {
				*m = im
				*samename = tm
				*ptr = 0
				return false
			}
		}

		return true
	}

	t = psess.methtype(t)
	var tms []*types.Field
	if t != nil {
		psess.
			expandmeth(t)
		tms = t.AllMethods().Slice()
	}
	i := 0
	for _, im := range iface.Fields(psess.types).Slice() {
		if im.Broke() {
			continue
		}
		for i < len(tms) && tms[i].Sym != im.Sym {
			i++
		}
		if i == len(tms) {
			*m = im
			*samename, _ = psess.ifacelookdot(im.Sym, t, true)
			*ptr = 0
			return false
		}
		tm := tms[i]
		if tm.Nointerface() || !psess.eqtype(tm.Type, im.Type) {
			*m = im
			*samename = tm
			*ptr = 0
			return false
		}
		followptr := tm.Embedded == 2

		rcvr := tm.Type.Recv(psess.types).Type
		if rcvr.IsPtr() && !t0.IsPtr() && !followptr && !psess.isifacemethod(tm.Type) {
			if false && psess.Debug['r'] != 0 {
				psess.
					yyerror("interface pointer mismatch")
			}

			*m = im
			*samename = nil
			*ptr = 1
			return false
		}
	}

	if psess.isdirectiface(t0) && !iface.IsEmptyInterface(psess.types) {
		psess.
			itabname(t0, iface)
	}
	return true
}

func (psess *PackageSession) listtreecopy(l []*Node, pos src.XPos) []*Node {
	var out []*Node
	for _, n := range l {
		out = append(out, psess.treecopy(n, pos))
	}
	return out
}

func (psess *PackageSession) liststmt(l []*Node) *Node {
	n := psess.nod(OBLOCK, nil, nil)
	n.List.Set(l)
	if len(l) != 0 {
		n.Pos = l[0].Pos
	}
	return n
}

func (l Nodes) asblock(psess *PackageSession) *Node {
	n := psess.nod(OBLOCK, nil, nil)
	n.List = l
	if l.Len() != 0 {
		n.Pos = l.First().Pos
	}
	return n
}

func (psess *PackageSession) ngotype(n *Node) *types.Sym {
	if n.Type != nil {
		return psess.typenamesym(n.Type)
	}
	return nil
}

// The result of addinit MUST be assigned back to n, e.g.
// 	n.Left = addinit(n.Left, init)
func (psess *PackageSession) addinit(n *Node, init []*Node) *Node {
	if len(init) == 0 {
		return n
	}
	if n.mayBeShared() {

		n = psess.nod(OCONVNOP, n, nil)
		n.Type = n.Left.Type
		n.SetTypecheck(1)
	}

	n.Ninit.Prepend(init...)
	n.SetHasCall(true)
	return n
}

// The linker uses the magic symbol prefixes "go." and "type."
// Avoid potential confusion between import paths and symbols
// by rejecting these reserved imports for now. Also, people
// "can do weird things in GOPATH and we'd prefer they didn't
// do _that_ weird thing" (per rsc). See also #4257.

func (psess *PackageSession) isbadimport(path string, allowSpace bool) bool {
	if strings.Contains(path, "\x00") {
		psess.
			yyerror("import path contains NUL")
		return true
	}

	for _, ri := range psess.reservedimports {
		if path == ri {
			psess.
				yyerror("import path %q is reserved and cannot be used", path)
			return true
		}
	}

	for _, r := range path {
		if r == utf8.RuneError {
			psess.
				yyerror("import path contains invalid UTF-8 sequence: %q", path)
			return true
		}

		if r < 0x20 || r == 0x7f {
			psess.
				yyerror("import path contains control character: %q", path)
			return true
		}

		if r == '\\' {
			psess.
				yyerror("import path contains backslash; use slash: %q", path)
			return true
		}

		if !allowSpace && unicode.IsSpace(r) {
			psess.
				yyerror("import path contains space character: %q", path)
			return true
		}

		if strings.ContainsRune("!\"#$%&'()*,:;<=>?[]^`{|}", r) {
			psess.
				yyerror("import path contains invalid character '%c': %q", r, path)
			return true
		}
	}

	return false
}

func (psess *PackageSession) checknil(x *Node, init *Nodes) {
	x = psess.walkexpr(x, nil)
	if x.Type.IsInterface() {
		x = psess.nod(OITAB, x, nil)
		x = psess.typecheck(x, Erv)
	}

	n := psess.nod(OCHECKNIL, x, nil)
	n.SetTypecheck(1)
	init.Append(n)
}

// Can this type be stored directly in an interface word?
// Yes, if the representation is a single pointer.
func (psess *PackageSession) isdirectiface(t *types.Type) bool {
	if t.Broke() {
		return false
	}

	switch t.Etype {
	case TPTR32,
		TPTR64,
		TCHAN,
		TMAP,
		TFUNC,
		TUNSAFEPTR:
		return true

	case TARRAY:

		return t.NumElem(psess.types) == 1 && psess.isdirectiface(t.Elem(psess.types))

	case TSTRUCT:

		return t.NumFields(psess.types) == 1 && psess.isdirectiface(t.Field(psess.types, 0).Type)
	}

	return false
}

// itabType loads the _type field from a runtime.itab struct.
func (psess *PackageSession) itabType(itab *Node) *Node {
	typ := psess.nodSym(ODOTPTR, itab, nil)
	typ.Type = psess.types.NewPtr(psess.types.Types[TUINT8])
	typ.SetTypecheck(1)
	typ.Xoffset = int64(psess.Widthptr)
	typ.SetBounded(true)
	return typ
}

// ifaceData loads the data field from an interface.
// The concrete type must be known to have type t.
// It follows the pointer if !isdirectiface(t).
func (psess *PackageSession) ifaceData(n *Node, t *types.Type) *Node {
	ptr := psess.nodSym(OIDATA, n, nil)
	if psess.isdirectiface(t) {
		ptr.Type = t
		ptr.SetTypecheck(1)
		return ptr
	}
	ptr.Type = psess.types.NewPtr(t)
	ptr.SetBounded(true)
	ptr.SetTypecheck(1)
	ind := psess.nod(OIND, ptr, nil)
	ind.Type = t
	ind.SetTypecheck(1)
	return ind
}

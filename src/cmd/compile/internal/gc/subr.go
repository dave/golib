// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"os"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

type Error struct {
	pos src.XPos
	msg string
}

func (pstate *PackageState) errorexit() {
	pstate.flusherrors()
	if pstate.outfile != "" {
		os.Remove(pstate.outfile)
	}
	os.Exit(2)
}

func (pstate *PackageState) adderrorname(n *Node) {
	if n.Op != ODOT {
		return
	}
	old := fmt.Sprintf("%v: undefined: %v\n", n.Line(pstate), n.Left)
	if len(pstate.errors) > 0 && pstate.errors[len(pstate.errors)-1].pos.Line() == n.Pos.Line() && pstate.errors[len(pstate.errors)-1].msg == old {
		pstate.errors[len(pstate.errors)-1].msg = fmt.Sprintf("%v: undefined: %v in %v\n", n.Line(pstate), n.Left, n)
	}
}

func (pstate *PackageState) adderr(pos src.XPos, format string, args ...interface{}) {
	pstate.errors = append(pstate.errors, Error{
		pos: pos,
		msg: fmt.Sprintf("%v: %s\n", pstate.linestr(pos), fmt.Sprintf(format, args...)),
	})
}

// byPos sorts errors by source position.
type byPos []Error

func (x byPos) Len() int           { return len(x) }
func (x byPos) Less(i, j int) bool { return x[i].pos.Before(x[j].pos) }
func (x byPos) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// flusherrors sorts errors seen so far by line number, prints them to stdout,
// and empties the errors array.
func (pstate *PackageState) flusherrors() {
	pstate.Ctxt.Bso.Flush()
	if len(pstate.errors) == 0 {
		return
	}
	sort.Stable(byPos(pstate.errors))
	for i, err := range pstate.errors {
		if i == 0 || err.msg != pstate.errors[i-1].msg {
			fmt.Printf("%s", err.msg)
		}
	}
	pstate.errors = pstate.errors[:0]
}

func (pstate *PackageState) hcrash() {
	if pstate.Debug['h'] != 0 {
		pstate.flusherrors()
		if pstate.outfile != "" {
			os.Remove(pstate.outfile)
		}
		var x *int
		*x = 0
	}
}

func (pstate *PackageState) linestr(pos src.XPos) string {
	return pstate.Ctxt.OutermostPos(pos).Format(pstate.src, pstate.Debug['C'] == 0, pstate.Debug['L'] == 1)
}

// sameline reports whether two positions a, b are on the same line.
func (pstate *PackageState) sameline(a, b src.XPos) bool {
	p := pstate.Ctxt.PosTable.Pos(a)
	q := pstate.Ctxt.PosTable.Pos(b)
	return p.Base() == q.Base() && p.Line() == q.Line()
}

func (pstate *PackageState) yyerrorl(pos src.XPos, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	if strings.HasPrefix(msg, "syntax error") {
		pstate.nsyntaxerrors++
		// only one syntax error per line, no matter what error
		if pstate.sameline(pstate.lasterror.syntax, pos) {
			return
		}
		pstate.lasterror.syntax = pos
	} else {
		// only one of multiple equal non-syntax errors per line
		// (flusherrors shows only one of them, so we filter them
		// here as best as we can (they may not appear in order)
		// so that we don't count them here and exit early, and
		// then have nothing to show for.)
		if pstate.sameline(pstate.lasterror.other, pos) && pstate.lasterror.msg == msg {
			return
		}
		pstate.lasterror.other = pos
		pstate.lasterror.msg = msg
	}

	pstate.adderr(pos, "%s", msg)

	pstate.hcrash()
	pstate.nerrors++
	if pstate.nsavederrors+pstate.nerrors >= 10 && pstate.Debug['e'] == 0 {
		pstate.flusherrors()
		fmt.Printf("%v: too many errors\n", pstate.linestr(pos))
		pstate.errorexit()
	}
}

func (pstate *PackageState) yyerror(format string, args ...interface{}) {
	pstate.yyerrorl(pstate.lineno, format, args...)
}

func (pstate *PackageState) Warn(fmt_ string, args ...interface{}) {
	pstate.adderr(pstate.lineno, fmt_, args...)

	pstate.hcrash()
}

func (pstate *PackageState) Warnl(line src.XPos, fmt_ string, args ...interface{}) {
	pstate.adderr(line, fmt_, args...)
	if pstate.Debug['m'] != 0 {
		pstate.flusherrors()
	}
}

func (pstate *PackageState) Fatalf(fmt_ string, args ...interface{}) {
	pstate.flusherrors()

	if pstate.Debug_panic != 0 || pstate.nsavederrors+pstate.nerrors == 0 {
		fmt.Printf("%v: internal compiler error: ", pstate.linestr(pstate.lineno))
		fmt.Printf(fmt_, args...)
		fmt.Printf("\n")

		// If this is a released compiler version, ask for a bug report.
		if strings.HasPrefix(pstate.objabi.Version, "go") {
			fmt.Printf("\n")
			fmt.Printf("Please file a bug report including a short program that triggers the error.\n")
			fmt.Printf("https://golang.org/issue/new\n")
		} else {
			// Not a release; dump a stack trace, too.
			fmt.Println()
			os.Stdout.Write(debug.Stack())
			fmt.Println()
		}
	}

	pstate.hcrash()
	pstate.errorexit()
}

func (pstate *PackageState) setlineno(n *Node) src.XPos {
	lno := pstate.lineno
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
			pstate.lineno = n.Pos
			if !pstate.lineno.IsKnown() {
				if pstate.Debug['K'] != 0 {
					pstate.Warn("setlineno: unknown position (line 0)")
				}
				pstate.lineno = lno
			}
		}
	}

	return lno
}

func (pstate *PackageState) lookup(name string) *types.Sym {
	return pstate.localpkg.Lookup(pstate.types, name)
}

// lookupN looks up the symbol starting with prefix and ending with
// the decimal n. If prefix is too long, lookupN panics.
func (pstate *PackageState) lookupN(prefix string, n int) *types.Sym {
	var buf [20]byte // plenty long enough for all current users
	copy(buf[:], prefix)
	b := strconv.AppendInt(buf[:len(prefix)], int64(n), 10)
	return pstate.localpkg.LookupBytes(pstate.types, b)
}

// autolabel generates a new Name node for use with
// an automatically generated label.
// prefix is a short mnemonic (e.g. ".s" for switch)
// to help with debugging.
// It should begin with "." to avoid conflicts with
// user labels.
func (pstate *PackageState) autolabel(prefix string) *Node {
	if prefix[0] != '.' {
		pstate.Fatalf("autolabel prefix must start with '.', have %q", prefix)
	}
	fn := pstate.Curfn
	if pstate.Curfn == nil {
		pstate.Fatalf("autolabel outside function")
	}
	n := fn.Func.Label
	fn.Func.Label++
	return pstate.newname(pstate.lookupN(prefix, int(n)))
}

func (pstate *PackageState) restrictlookup(name string, pkg *types.Pkg) *types.Sym {
	if !types.IsExported(name) && pkg != pstate.localpkg {
		pstate.yyerror("cannot refer to unexported name %s.%s", pkg.Name, name)
	}
	return pkg.Lookup(pstate.types, name)
}

// find all the exported symbols in package opkg
// and make them available in the current package
func (pstate *PackageState) importdot(opkg *types.Pkg, pack *Node) {
	n := 0
	for _, s := range opkg.Syms {
		if s.Def == nil {
			continue
		}
		if !types.IsExported(s.Name) || strings.ContainsRune(s.Name, 0xb7) { // 0xb7 = center dot
			continue
		}
		s1 := pstate.lookup(s.Name)
		if s1.Def != nil {
			pkgerror := fmt.Sprintf("during import %q", opkg.Path)
			pstate.redeclare(pstate.lineno, s1, pkgerror)
			continue
		}

		s1.Def = s.Def
		s1.Block = s.Block
		if asNode(s1.Def).Name == nil {
			Dump("s1def", asNode(s1.Def))
			pstate.Fatalf("missing Name")
		}
		asNode(s1.Def).Name.Pack = pack
		s1.Origpkg = opkg
		n++
	}

	if n == 0 {
		// can't possibly be used - there were no symbols
		pstate.yyerrorl(pack.Pos, "imported and not used: %q", opkg.Path)
	}
}

func (pstate *PackageState) nod(op Op, nleft, nright *Node) *Node {
	return pstate.nodl(pstate.lineno, op, nleft, nright)
}

func (pstate *PackageState) nodl(pos src.XPos, op Op, nleft, nright *Node) *Node {
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
		pstate.Fatalf("use newname instead")
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
func (pstate *PackageState) newname(s *types.Sym) *Node {
	n := pstate.newnamel(pstate.lineno, s)
	n.Name.Curfn = pstate.Curfn
	return n
}

// newname returns a new ONAME Node associated with symbol s at position pos.
// The caller is responsible for setting n.Name.Curfn.
func (pstate *PackageState) newnamel(pos src.XPos, s *types.Sym) *Node {
	if s == nil {
		pstate.Fatalf("newnamel nil")
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
func (pstate *PackageState) nodSym(op Op, left *Node, sym *types.Sym) *Node {
	n := pstate.nod(op, left, nil)
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

func (pstate *PackageState) nodintconst(v int64) *Node {
	u := new(Mpint)
	u.SetInt64(v)
	return pstate.nodlit(Val{u})
}

func (pstate *PackageState) nodfltconst(v *Mpflt) *Node {
	u := newMpflt()
	u.Set(v)
	return pstate.nodlit(Val{u})
}

func (pstate *PackageState) nodnil() *Node {
	return pstate.nodlit(Val{new(NilVal)})
}

func (pstate *PackageState) nodbool(b bool) *Node {
	return pstate.nodlit(Val{b})
}

func (pstate *PackageState) nodstr(s string) *Node {
	return pstate.nodlit(Val{s})
}

// treecopy recursively copies n, with the exception of
// ONAME, OLITERAL, OTYPE, and non-iota ONONAME leaves.
// Copies of iota ONONAME nodes are assigned the current
// value of iota_. If pos.IsKnown(), it sets the source
// position of newly allocated nodes to pos.
func (pstate *PackageState) treecopy(n *Node, pos src.XPos) *Node {
	if n == nil {
		return nil
	}

	switch n.Op {
	default:
		m := n.copy()
		m.Orig = m
		m.Left = pstate.treecopy(n.Left, pos)
		m.Right = pstate.treecopy(n.Right, pos)
		m.List.Set(pstate.listtreecopy(n.List.Slice(), pos))
		if pos.IsKnown() {
			m.Pos = pos
		}
		if m.Name != nil && n.Op != ODCLFIELD {
			Dump("treecopy", n)
			pstate.Fatalf("treecopy Name")
		}
		return m

	case OPACK:
		// OPACK nodes are never valid in const value declarations,
		// but allow them like any other declared symbol to avoid
		// crashing (golang.org/issue/11361).
		fallthrough

	case ONAME, ONONAME, OLITERAL, OTYPE:
		return n

	}
}

// isNil reports whether n represents the universal untyped zero value "nil".
func (n *Node) isNil(pstate *PackageState) bool {
	// Check n.Orig because constant propagation may produce typed nil constants,
	// which don't exist in the Go spec.
	return pstate.Isconst(n.Orig, CTNIL)
}

func (pstate *PackageState) isptrto(t *types.Type, et types.EType) bool {
	if t == nil {
		return false
	}
	if !t.IsPtr() {
		return false
	}
	t = t.Elem(pstate.types)
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
func (pstate *PackageState) methtype(t *types.Type) *types.Type {
	if t == nil {
		return nil
	}

	// Strip away pointer if it's there.
	if t.IsPtr() {
		if t.Sym != nil {
			return nil
		}
		t = t.Elem(pstate.types)
		if t == nil {
			return nil
		}
	}

	// Must be a named type or anonymous struct.
	if t.Sym == nil && !t.IsStruct() {
		return nil
	}

	// Check types.
	if pstate.issimple[t.Etype] {
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
func (pstate *PackageState) eqtype(t1, t2 *types.Type) bool {
	return pstate.eqtype1(t1, t2, true, nil)
}

// eqtypeIgnoreTags is like eqtype but it ignores struct tags for struct identity.
func (pstate *PackageState) eqtypeIgnoreTags(t1, t2 *types.Type) bool {
	return pstate.eqtype1(t1, t2, false, nil)
}

type typePair struct {
	t1 *types.Type
	t2 *types.Type
}

func (pstate *PackageState) eqtype1(t1, t2 *types.Type, cmpTags bool, assumedEqual map[typePair]struct{}) bool {
	if t1 == t2 {
		return true
	}
	if t1 == nil || t2 == nil || t1.Etype != t2.Etype || t1.Broke() || t2.Broke() {
		return false
	}
	if t1.Sym != nil || t2.Sym != nil {
		// Special case: we keep byte/uint8 and rune/int32
		// separate for error messages. Treat them as equal.
		switch t1.Etype {
		case TUINT8:
			return (t1 == pstate.types.Types[TUINT8] || t1 == pstate.types.Bytetype) && (t2 == pstate.types.Types[TUINT8] || t2 == pstate.types.Bytetype)
		case TINT32:
			return (t1 == pstate.types.Types[TINT32] || t1 == pstate.types.Runetype) && (t2 == pstate.types.Types[TINT32] || t2 == pstate.types.Runetype)
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
		if t1.NumFields(pstate.types) != t2.NumFields(pstate.types) {
			return false
		}
		for i, f1 := range t1.FieldSlice(pstate.types) {
			f2 := t2.Field(pstate.types, i)
			if f1.Sym != f2.Sym || !pstate.eqtype1(f1.Type, f2.Type, cmpTags, assumedEqual) {
				return false
			}
		}
		return true

	case TSTRUCT:
		if t1.NumFields(pstate.types) != t2.NumFields(pstate.types) {
			return false
		}
		for i, f1 := range t1.FieldSlice(pstate.types) {
			f2 := t2.Field(pstate.types, i)
			if f1.Sym != f2.Sym || f1.Embedded != f2.Embedded || !pstate.eqtype1(f1.Type, f2.Type, cmpTags, assumedEqual) {
				return false
			}
			if cmpTags && f1.Note != f2.Note {
				return false
			}
		}
		return true

	case TFUNC:
		// Check parameters and result parameters for type equality.
		// We intentionally ignore receiver parameters for type
		// equality, because they're never relevant.
		for _, f := range pstate.types.ParamsResults {
			// Loop over fields in structs, ignoring argument names.
			fs1, fs2 := f(t1).FieldSlice(pstate.types), f(t2).FieldSlice(pstate.types)
			if len(fs1) != len(fs2) {
				return false
			}
			for i, f1 := range fs1 {
				f2 := fs2[i]
				if f1.Isddd() != f2.Isddd() || !pstate.eqtype1(f1.Type, f2.Type, cmpTags, assumedEqual) {
					return false
				}
			}
		}
		return true

	case TARRAY:
		if t1.NumElem(pstate.types) != t2.NumElem(pstate.types) {
			return false
		}

	case TCHAN:
		if t1.ChanDir(pstate.types) != t2.ChanDir(pstate.types) {
			return false
		}

	case TMAP:
		if !pstate.eqtype1(t1.Key(pstate.types), t2.Key(pstate.types), cmpTags, assumedEqual) {
			return false
		}
	}

	return pstate.eqtype1(t1.Elem(pstate.types), t2.Elem(pstate.types), cmpTags, assumedEqual)
}

// Are t1 and t2 equal struct types when field names are ignored?
// For deciding whether the result struct from g can be copied
// directly when compiling f(g()).
func (pstate *PackageState) eqtypenoname(t1 *types.Type, t2 *types.Type) bool {
	if t1 == nil || t2 == nil || !t1.IsStruct() || !t2.IsStruct() {
		return false
	}

	if t1.NumFields(pstate.types) != t2.NumFields(pstate.types) {
		return false
	}
	for i, f1 := range t1.FieldSlice(pstate.types) {
		f2 := t2.Field(pstate.types, i)
		if !pstate.eqtype(f1.Type, f2.Type) {
			return false
		}
	}
	return true
}

// Is type src assignment compatible to type dst?
// If so, return op code to use in conversion.
// If not, return 0.
func (pstate *PackageState) assignop(src *types.Type, dst *types.Type, why *string) Op {
	if why != nil {
		*why = ""
	}

	// TODO(rsc,lvd): This behaves poorly in the presence of inlining.
	// https://golang.org/issue/2795
	if pstate.safemode && !pstate.inimport && src != nil && src.Etype == TUNSAFEPTR {
		pstate.yyerror("cannot use unsafe.Pointer")
		pstate.errorexit()
	}

	if src == dst {
		return OCONVNOP
	}
	if src == nil || dst == nil || src.Etype == TFORW || dst.Etype == TFORW || src.Orig == nil || dst.Orig == nil {
		return 0
	}

	// 1. src type is identical to dst.
	if pstate.eqtype(src, dst) {
		return OCONVNOP
	}

	// 2. src and dst have identical underlying types
	// and either src or dst is not a named type or
	// both are empty interface types.
	// For assignable but different non-empty interface types,
	// we want to recompute the itab. Recomputing the itab ensures
	// that itabs are unique (thus an interface with a compile-time
	// type I has an itab with interface type I).
	if pstate.eqtype(src.Orig, dst.Orig) {
		if src.IsEmptyInterface(pstate.types) {
			// Conversion between two empty interfaces
			// requires no code.
			return OCONVNOP
		}
		if (src.Sym == nil || dst.Sym == nil) && !src.IsInterface() {
			// Conversion between two types, at least one unnamed,
			// needs no conversion. The exception is nonempty interfaces
			// which need to have their itab updated.
			return OCONVNOP
		}
	}

	// 3. dst is an interface type and src implements dst.
	if dst.IsInterface() && src.Etype != TNIL {
		var missing, have *types.Field
		var ptr int
		if pstate.implements(src, dst, &missing, &have, &ptr) {
			return OCONVIFACE
		}

		// we'll have complained about this method anyway, suppress spurious messages.
		if have != nil && have.Sym == missing.Sym && (have.Type.Broke() || missing.Type.Broke()) {
			return OCONVIFACE
		}

		if why != nil {
			if pstate.isptrto(src, TINTER) {
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

	if pstate.isptrto(dst, TINTER) {
		if why != nil {
			*why = fmt.Sprintf(":\n\t%v is pointer to interface, not interface", dst)
		}
		return 0
	}

	if src.IsInterface() && dst.Etype != TBLANK {
		var missing, have *types.Field
		var ptr int
		if why != nil && pstate.implements(dst, src, &missing, &have, &ptr) {
			*why = ": need type assertion"
		}
		return 0
	}

	// 4. src is a bidirectional channel value, dst is a channel type,
	// src and dst have identical element types, and
	// either src or dst is not a named type.
	if src.IsChan() && src.ChanDir(pstate.types) == types.Cboth && dst.IsChan() {
		if pstate.eqtype(src.Elem(pstate.types), dst.Elem(pstate.types)) && (src.Sym == nil || dst.Sym == nil) {
			return OCONVNOP
		}
	}

	// 5. src is the predeclared identifier nil and dst is a nillable type.
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

	// 6. rule about untyped constants - already converted by defaultlit.

	// 7. Any typed value can be assigned to the blank identifier.
	if dst.Etype == TBLANK {
		return OCONVNOP
	}

	return 0
}

// Can we convert a value of type src to a value of type dst?
// If so, return op code to use in conversion (maybe OCONVNOP).
// If not, return 0.
func (pstate *PackageState) convertop(src *types.Type, dst *types.Type, why *string) Op {
	if why != nil {
		*why = ""
	}

	if src == dst {
		return OCONVNOP
	}
	if src == nil || dst == nil {
		return 0
	}

	// Conversions from regular to go:notinheap are not allowed
	// (unless it's unsafe.Pointer). This is a runtime-specific
	// rule.
	if src.IsPtr() && dst.IsPtr() && dst.Elem(pstate.types).NotInHeap() && !src.Elem(pstate.types).NotInHeap() {
		if why != nil {
			*why = fmt.Sprintf(":\n\t%v is go:notinheap, but %v is not", dst.Elem(pstate.types), src.Elem(pstate.types))
		}
		return 0
	}

	// 1. src can be assigned to dst.
	op := pstate.assignop(src, dst, why)
	if op != 0 {
		return op
	}

	// The rules for interfaces are no different in conversions
	// than assignments. If interfaces are involved, stop now
	// with the good message from assignop.
	// Otherwise clear the error.
	if src.IsInterface() || dst.IsInterface() {
		return 0
	}
	if why != nil {
		*why = ""
	}

	// 2. Ignoring struct tags, src and dst have identical underlying types.
	if pstate.eqtypeIgnoreTags(src.Orig, dst.Orig) {
		return OCONVNOP
	}

	// 3. src and dst are unnamed pointer types and, ignoring struct tags,
	// their base types have identical underlying types.
	if src.IsPtr() && dst.IsPtr() && src.Sym == nil && dst.Sym == nil {
		if pstate.eqtypeIgnoreTags(src.Elem(pstate.types).Orig, dst.Elem(pstate.types).Orig) {
			return OCONVNOP
		}
	}

	// 4. src and dst are both integer or floating point types.
	if (src.IsInteger() || src.IsFloat()) && (dst.IsInteger() || dst.IsFloat()) {
		if pstate.simtype[src.Etype] == pstate.simtype[dst.Etype] {
			return OCONVNOP
		}
		return OCONV
	}

	// 5. src and dst are both complex types.
	if src.IsComplex() && dst.IsComplex() {
		if pstate.simtype[src.Etype] == pstate.simtype[dst.Etype] {
			return OCONVNOP
		}
		return OCONV
	}

	// 6. src is an integer or has type []byte or []rune
	// and dst is a string type.
	if src.IsInteger() && dst.IsString() {
		return ORUNESTR
	}

	if src.IsSlice() && dst.IsString() {
		if src.Elem(pstate.types).Etype == pstate.types.Bytetype.Etype {
			return OARRAYBYTESTR
		}
		if src.Elem(pstate.types).Etype == pstate.types.Runetype.Etype {
			return OARRAYRUNESTR
		}
	}

	// 7. src is a string and dst is []byte or []rune.
	// String to slice.
	if src.IsString() && dst.IsSlice() {
		if dst.Elem(pstate.types).Etype == pstate.types.Bytetype.Etype {
			return OSTRARRAYBYTE
		}
		if dst.Elem(pstate.types).Etype == pstate.types.Runetype.Etype {
			return OSTRARRAYRUNE
		}
	}

	// 8. src is a pointer or uintptr and dst is unsafe.Pointer.
	if (src.IsPtr() || src.Etype == TUINTPTR) && dst.Etype == TUNSAFEPTR {
		return OCONVNOP
	}

	// 9. src is unsafe.Pointer and dst is a pointer or uintptr.
	if src.Etype == TUNSAFEPTR && (dst.IsPtr() || dst.Etype == TUINTPTR) {
		return OCONVNOP
	}

	// src is map and dst is a pointer to corresponding hmap.
	// This rule is needed for the implementation detail that
	// go gc maps are implemented as a pointer to a hmap struct.
	if src.Etype == TMAP && dst.IsPtr() &&
		src.MapType(pstate.types).Hmap == dst.Elem(pstate.types) {
		return OCONVNOP
	}

	return 0
}

func (pstate *PackageState) assignconv(n *Node, t *types.Type, context string) *Node {
	return pstate.assignconvfn(n, t, func() string { return context })
}

// Convert node n for assignment to type t.
func (pstate *PackageState) assignconvfn(n *Node, t *types.Type, context func() string) *Node {
	if n == nil || n.Type == nil || n.Type.Broke() {
		return n
	}

	if t.Etype == TBLANK && n.Type.Etype == TNIL {
		pstate.yyerror("use of untyped nil")
	}

	old := n
	od := old.Diag()
	old.SetDiag(true) // silence errors about n; we'll issue one below
	n = pstate.defaultlit(n, t)
	old.SetDiag(od)
	if t.Etype == TBLANK {
		return n
	}

	// Convert ideal bool from comparison to plain bool
	// if the next step is non-bool (like interface{}).
	if n.Type == pstate.types.Idealbool && !t.IsBoolean() {
		if n.Op == ONAME || n.Op == OLITERAL {
			r := pstate.nod(OCONVNOP, n, nil)
			r.Type = pstate.types.Types[TBOOL]
			r.SetTypecheck(1)
			r.SetImplicit(true)
			n = r
		}
	}

	if pstate.eqtype(n.Type, t) {
		return n
	}

	var why string
	op := pstate.assignop(n.Type, t, &why)
	if op == 0 {
		if !old.Diag() {
			pstate.yyerror("cannot use %L as type %v in %s%s", n, t, context(), why)
		}
		op = OCONV
	}

	r := pstate.nod(op, n, nil)
	r.Type = t
	r.SetTypecheck(1)
	r.SetImplicit(true)
	r.Orig = n.Orig
	return r
}

// IsMethod reports whether n is a method.
// n must be a function or a method.
func (n *Node) IsMethod(pstate *PackageState) bool {
	return n.Type.Recv(pstate.types) != nil
}

// SliceBounds returns n's slice bounds: low, high, and max in expr[low:high:max].
// n must be a slice expression. max is nil if n is a simple slice expression.
func (n *Node) SliceBounds(pstate *PackageState) (low, high, max *Node) {
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
	pstate.Fatalf("SliceBounds op %v: %v", n.Op, n)
	return nil, nil, nil
}

// SetSliceBounds sets n's slice bounds, where n is a slice expression.
// n must be a slice expression. If max is non-nil, n must be a full slice expression.
func (n *Node) SetSliceBounds(pstate *PackageState, low, high, max *Node) {
	switch n.Op {
	case OSLICE, OSLICEARR, OSLICESTR:
		if max != nil {
			pstate.Fatalf("SetSliceBounds %v given three bounds", n.Op)
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
	pstate.Fatalf("SetSliceBounds op %v: %v", n.Op, n)
}

// IsSlice3 reports whether o is a slice3 op (OSLICE3, OSLICE3ARR).
// o must be a slicing op.
func (o Op) IsSlice3(pstate *PackageState) bool {
	switch o {
	case OSLICE, OSLICEARR, OSLICESTR:
		return false
	case OSLICE3, OSLICE3ARR:
		return true
	}
	pstate.Fatalf("IsSlice3 op %v", o)
	return false
}

// labeledControl returns the control flow Node (for, switch, select)
// associated with the label n, if any.
func (n *Node) labeledControl(pstate *PackageState) *Node {
	if n.Op != OLABEL {
		pstate.Fatalf("labeledControl %v", n.Op)
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

func (pstate *PackageState) syslook(name string) *Node {
	s := pstate.Runtimepkg.Lookup(pstate.types, name)
	if s == nil || s.Def == nil {
		pstate.Fatalf("syslook: can't find runtime.%s", name)
	}
	return asNode(s.Def)
}

// typehash computes a hash value for type t to use in type switch statements.
func (pstate *PackageState) typehash(t *types.Type) uint32 {
	p := t.LongString(pstate.types)

	// Using MD5 is overkill, but reduces accidental collisions.
	h := md5.Sum([]byte(p))
	return binary.LittleEndian.Uint32(h[:4])
}

func (pstate *PackageState) frame(context int) {
	if context != 0 {
		fmt.Printf("--- external frame ---\n")
		for _, n := range pstate.externdcl {
			printframenode(n)
		}
		return
	}

	if pstate.Curfn != nil {
		fmt.Printf("--- %v frame ---\n", pstate.Curfn.Func.Nname.Sym)
		for _, ln := range pstate.Curfn.Func.Dcl {
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
func (pstate *PackageState) updateHasCall(n *Node) {
	if n == nil {
		return
	}
	n.SetHasCall(pstate.calcHasCall(n))
}

func (pstate *PackageState) calcHasCall(n *Node) bool {
	if n.Ninit.Len() != 0 {
		// TODO(mdempsky): This seems overly conservative.
		return true
	}

	switch n.Op {
	case OLITERAL, ONAME, OTYPE:
		if n.HasCall() {
			pstate.Fatalf("OLITERAL/ONAME/OTYPE should never have calls: %+v", n)
		}
		return false
	case OCALL, OCALLFUNC, OCALLMETH, OCALLINTER:
		return true
	case OANDAND, OOROR:
		// hard with instrumented code
		if pstate.instrumenting {
			return true
		}
	case OINDEX, OSLICE, OSLICEARR, OSLICE3, OSLICE3ARR, OSLICESTR,
		OIND, ODOTPTR, ODOTTYPE, ODIV, OMOD:
		// These ops might panic, make sure they are done
		// before we start marshaling args for a call. See issue 16760.
		return true

	// When using soft-float, these ops might be rewritten to function calls
	// so we ensure they are evaluated first.
	case OADD, OSUB, OMINUS:
		if pstate.thearch.SoftFloat && (pstate.isFloat[n.Type.Etype] || pstate.isComplex[n.Type.Etype]) {
			return true
		}
	case OLT, OEQ, ONE, OLE, OGE, OGT:
		if pstate.thearch.SoftFloat && (pstate.isFloat[n.Left.Type.Etype] || pstate.isComplex[n.Left.Type.Etype]) {
			return true
		}
	case OCONV:
		if pstate.thearch.SoftFloat && ((pstate.isFloat[n.Type.Etype] || pstate.isComplex[n.Type.Etype]) || (pstate.isFloat[n.Left.Type.Etype] || pstate.isComplex[n.Left.Type.Etype])) {
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

func (pstate *PackageState) badtype(op Op, tl *types.Type, tr *types.Type) {
	fmt_ := ""
	if tl != nil {
		fmt_ += fmt.Sprintf("\n\t%v", tl)
	}
	if tr != nil {
		fmt_ += fmt.Sprintf("\n\t%v", tr)
	}

	// common mistake: *struct and *interface.
	if tl != nil && tr != nil && tl.IsPtr() && tr.IsPtr() {
		if tl.Elem(pstate.types).IsStruct() && tr.Elem(pstate.types).IsInterface() {
			fmt_ += "\n\t(*struct vs *interface)"
		} else if tl.Elem(pstate.types).IsInterface() && tr.Elem(pstate.types).IsStruct() {
			fmt_ += "\n\t(*interface vs *struct)"
		}
	}

	s := fmt_
	pstate.yyerror("illegal types for operand: %v%s", op, s)
}

// brcom returns !(op).
// For example, brcom(==) is !=.
func (pstate *PackageState) brcom(op Op) Op {
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
	pstate.Fatalf("brcom: no com for %v\n", op)
	return op
}

// brrev returns reverse(op).
// For example, Brrev(<) is >.
func (pstate *PackageState) brrev(op Op) Op {
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
	pstate.Fatalf("brrev: no rev for %v\n", op)
	return op
}

// return side effect-free n, appending side effects to init.
// result is assignable if n is.
func (pstate *PackageState) safeexpr(n *Node, init *Nodes) *Node {
	if n == nil {
		return nil
	}

	if n.Ninit.Len() != 0 {
		pstate.walkstmtlist(n.Ninit.Slice())
		init.AppendNodes(&n.Ninit)
	}

	switch n.Op {
	case ONAME, OLITERAL:
		return n

	case ODOT, OLEN, OCAP:
		l := pstate.safeexpr(n.Left, init)
		if l == n.Left {
			return n
		}
		r := n.copy()
		r.Left = l
		r = pstate.typecheck(r, Erv)
		r = pstate.walkexpr(r, init)
		return r

	case ODOTPTR, OIND:
		l := pstate.safeexpr(n.Left, init)
		if l == n.Left {
			return n
		}
		a := n.copy()
		a.Left = l
		a = pstate.walkexpr(a, init)
		return a

	case OINDEX, OINDEXMAP:
		l := pstate.safeexpr(n.Left, init)
		r := pstate.safeexpr(n.Right, init)
		if l == n.Left && r == n.Right {
			return n
		}
		a := n.copy()
		a.Left = l
		a.Right = r
		a = pstate.walkexpr(a, init)
		return a

	case OSTRUCTLIT, OARRAYLIT, OSLICELIT:
		if pstate.isStaticCompositeLiteral(n) {
			return n
		}
	}

	// make a copy; must not be used as an lvalue
	if islvalue(n) {
		pstate.Fatalf("missing lvalue case in safeexpr: %v", n)
	}
	return pstate.cheapexpr(n, init)
}

func (pstate *PackageState) copyexpr(n *Node, t *types.Type, init *Nodes) *Node {
	l := pstate.temp(t)
	a := pstate.nod(OAS, l, n)
	a = pstate.typecheck(a, Etop)
	a = pstate.walkexpr(a, init)
	init.Append(a)
	return l
}

// return side-effect free and cheap n, appending side effects to init.
// result may not be assignable.
func (pstate *PackageState) cheapexpr(n *Node, init *Nodes) *Node {
	switch n.Op {
	case ONAME, OLITERAL:
		return n
	}

	return pstate.copyexpr(n, n.Type, init)
}

// Code to resolve elided DOTs in embedded types.

// A Dlist stores a pointer to a TFIELD Type embedded within
// a TSTRUCT or TINTER Type.
type Dlist struct {
	field *types.Field
}

// lookdot0 returns the number of fields or methods named s associated
// with Type t. If exactly one exists, it will be returned in *save
// (if save is not nil).
func (pstate *PackageState) lookdot0(s *types.Sym, t *types.Type, save **types.Field, ignorecase bool) int {
	u := t
	if u.IsPtr() {
		u = u.Elem(pstate.types)
	}

	c := 0
	if u.IsStruct() || u.IsInterface() {
		for _, f := range u.Fields(pstate.types).Slice() {
			if f.Sym == s || (ignorecase && f.Type.Etype == TFUNC && f.Type.Recv(pstate.types) != nil && strings.EqualFold(f.Sym.Name, s.Name)) {
				if save != nil {
					*save = f
				}
				c++
			}
		}
	}

	u = pstate.methtype(t)
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
func (pstate *PackageState) adddot1(s *types.Sym, t *types.Type, d int, save **types.Field, ignorecase bool) (c int, more bool) {
	if t.Recur() {
		return
	}
	t.SetRecur(true)
	defer t.SetRecur(false)

	var u *types.Type
	d--
	if d < 0 {
		// We've reached our target depth. If t has any fields/methods
		// named s, then we're done. Otherwise, we still need to check
		// below for embedded fields.
		c = pstate.lookdot0(s, t, save, ignorecase)
		if c != 0 {
			return c, false
		}
	}

	u = t
	if u.IsPtr() {
		u = u.Elem(pstate.types)
	}
	if !u.IsStruct() && !u.IsInterface() {
		return c, false
	}

	for _, f := range u.Fields(pstate.types).Slice() {
		if f.Embedded == 0 || f.Sym == nil {
			continue
		}
		if d < 0 {
			// Found an embedded field at target depth.
			return c, true
		}
		a, more1 := pstate.adddot1(s, f.Type, d, save, ignorecase)
		if a != 0 && c == 0 {
			pstate.dotlist[d].field = f
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
func (pstate *PackageState) dotpath(s *types.Sym, t *types.Type, save **types.Field, ignorecase bool) (path []Dlist, ambig bool) {
	// The embedding of types within structs imposes a tree structure onto
	// types: structs parent the types they embed, and types parent their
	// fields or methods. Our goal here is to find the shortest path to
	// a field or method named s in the subtree rooted at t. To accomplish
	// that, we iteratively perform depth-first searches of increasing depth
	// until we either find the named field/method or exhaust the tree.
	for d := 0; ; d++ {
		if d > len(pstate.dotlist) {
			pstate.dotlist = append(pstate.dotlist, Dlist{})
		}
		if c, more := pstate.adddot1(s, t, d, save, ignorecase); c == 1 {
			return pstate.dotlist[:d], false
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
func (pstate *PackageState) adddot(n *Node) *Node {
	n.Left = pstate.typecheck(n.Left, Etype|Erv)
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

	switch path, ambig := pstate.dotpath(s, t, nil, false); {
	case path != nil:
		// rebuild elided dots
		for c := len(path) - 1; c >= 0; c-- {
			n.Left = pstate.nodSym(ODOT, n.Left, path[c].field.Sym)
			n.Left.SetImplicit(true)
		}
	case ambig:
		pstate.yyerror("ambiguous selector %v", n)
		n.Left = nil
	}

	return n
}

// Code to help generate trampoline functions for methods on embedded
// types. These are approx the same as the corresponding adddot
// routines except that they expect to be called with unique tasks and
// they return the actual methods.

type Symlink struct {
	field *types.Field
}

func (pstate *PackageState) expand0(t *types.Type) {
	u := t
	if u.IsPtr() {
		u = u.Elem(pstate.types)
	}

	if u.IsInterface() {
		for _, f := range u.Fields(pstate.types).Slice() {
			if f.Sym.Uniq() {
				continue
			}
			f.Sym.SetUniq(true)
			pstate.slist = append(pstate.slist, Symlink{field: f})
		}

		return
	}

	u = pstate.methtype(t)
	if u != nil {
		for _, f := range u.Methods().Slice() {
			if f.Sym.Uniq() {
				continue
			}
			f.Sym.SetUniq(true)
			pstate.slist = append(pstate.slist, Symlink{field: f})
		}
	}
}

func (pstate *PackageState) expand1(t *types.Type, top bool) {
	if t.Recur() {
		return
	}
	t.SetRecur(true)

	if !top {
		pstate.expand0(t)
	}

	u := t
	if u.IsPtr() {
		u = u.Elem(pstate.types)
	}

	if u.IsStruct() || u.IsInterface() {
		for _, f := range u.Fields(pstate.types).Slice() {
			if f.Embedded == 0 {
				continue
			}
			if f.Sym == nil {
				continue
			}
			pstate.expand1(f.Type, false)
		}
	}

	t.SetRecur(false)
}

func (pstate *PackageState) expandmeth(t *types.Type) {
	if t == nil || t.AllMethods().Len() != 0 {
		return
	}

	// mark top-level method symbols
	// so that expand1 doesn't consider them.
	for _, f := range t.Methods().Slice() {
		f.Sym.SetUniq(true)
	}

	// generate all reachable methods
	pstate.slist = pstate.slist[:0]
	pstate.expand1(t, true)

	// check each method to be uniquely reachable
	var ms []*types.Field
	for i, sl := range pstate.slist {
		pstate.slist[i].field = nil
		sl.field.Sym.SetUniq(false)

		var f *types.Field
		path, _ := pstate.dotpath(sl.field.Sym, t, &f, false)
		if path == nil {
			continue
		}

		// dotpath may have dug out arbitrary fields, we only want methods.
		if f.Type.Etype != TFUNC || f.Type.Recv(pstate.types) == nil {
			continue
		}

		// add it to the base type method list
		f = f.Copy()
		f.Embedded = 1 // needs a trampoline
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
func (pstate *PackageState) structargs(tl *types.Type, mustname bool) []*Node {
	var args []*Node
	gen := 0
	for _, t := range tl.Fields(pstate.types).Slice() {
		s := t.Sym
		if mustname && (s == nil || s.Name == "_") {
			// invent a name so that we can refer to it in the trampoline
			s = pstate.lookupN(".anon", gen)
			gen++
		}
		a := pstate.symfield(s, t.Type)
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
func (pstate *PackageState) genwrapper(rcvr *types.Type, method *types.Field, newnam *types.Sym) {
	if false && pstate.Debug['r'] != 0 {
		fmt.Printf("genwrapper rcvrtype=%v method=%v newnam=%v\n", rcvr, method, newnam)
	}

	// Only generate (*T).M wrappers for T.M in T's own package.
	if rcvr.IsPtr() && rcvr.Elem(pstate.types) == method.Type.Recv(pstate.types).Type &&
		rcvr.Elem(pstate.types).Sym != nil && rcvr.Elem(pstate.types).Sym.Pkg != pstate.localpkg {
		return
	}

	// Only generate I.M wrappers for I in I's own package.
	if rcvr.IsInterface() && rcvr.Sym != nil && rcvr.Sym.Pkg != pstate.localpkg {
		return
	}

	pstate.lineno = pstate.autogeneratedPos
	pstate.dclcontext = PEXTERN

	tfn := pstate.nod(OTFUNC, nil, nil)
	tfn.Left = pstate.namedfield(".this", rcvr)
	tfn.List.Set(pstate.structargs(method.Type.Params(pstate.types), true))
	tfn.Rlist.Set(pstate.structargs(method.Type.Results(pstate.types), false))

	disableExport(newnam)
	fn := pstate.dclfunc(newnam, tfn)
	fn.Func.SetDupok(true)

	nthis := asNode(tfn.Type.Recv(pstate.types).Nname)

	methodrcvr := method.Type.Recv(pstate.types).Type

	// generate nil pointer check for better error
	if rcvr.IsPtr() && rcvr.Elem(pstate.types) == methodrcvr {
		// generating wrapper from *T to T.
		n := pstate.nod(OIF, nil, nil)
		n.Left = pstate.nod(OEQ, nthis, pstate.nodnil())
		call := pstate.nod(OCALL, pstate.syslook("panicwrap"), nil)
		n.Nbody.Set1(call)
		fn.Nbody.Append(n)
	}

	dot := pstate.adddot(pstate.nodSym(OXDOT, nthis, method.Sym))

	// generate call
	// It's not possible to use a tail call when dynamic linking on ppc64le. The
	// bad scenario is when a local call is made to the wrapper: the wrapper will
	// call the implementation, which might be in a different module and so set
	// the TOC to the appropriate value for that module. But if it returns
	// directly to the wrapper's caller, nothing will reset it to the correct
	// value for that function.
	if !pstate.instrumenting && rcvr.IsPtr() && methodrcvr.IsPtr() && method.Embedded != 0 && !pstate.isifacemethod(method.Type) && !(pstate.thearch.LinkArch.Name == "ppc64le" && pstate.Ctxt.Flag_dynlink) {
		// generate tail call: adjust pointer receiver and jump to embedded method.
		dot = dot.Left // skip final .M
		// TODO(mdempsky): Remove dependency on dotlist.
		if !pstate.dotlist[0].field.Type.IsPtr() {
			dot = pstate.nod(OADDR, dot, nil)
		}
		as := pstate.nod(OAS, nthis, pstate.nod(OCONVNOP, dot, nil))
		as.Right.Type = rcvr
		fn.Nbody.Append(as)
		fn.Nbody.Append(pstate.nodSym(ORETJMP, nil, pstate.methodSym(methodrcvr, method.Sym)))
	} else {
		fn.Func.SetWrapper(true) // ignore frame for panic+recover matching
		call := pstate.nod(OCALL, dot, nil)
		call.List.Set(pstate.paramNnames(tfn.Type))
		call.SetIsddd(tfn.Type.IsVariadic(pstate.types))
		if method.Type.NumResults(pstate.types) > 0 {
			n := pstate.nod(ORETURN, nil, nil)
			n.List.Set1(call)
			call = n
		}
		fn.Nbody.Append(call)
	}

	if false && pstate.Debug['r'] != 0 {
		dumplist("genwrapper body", fn.Nbody)
	}

	pstate.funcbody()
	if pstate.debug_dclstack != 0 {
		pstate.testdclstack()
	}

	fn = pstate.typecheck(fn, Etop)

	pstate.Curfn = fn
	pstate.typecheckslice(fn.Nbody.Slice(), Etop)

	// TODO(mdempsky): Investigate why this doesn't work with
	// indexed export. For now, we disable even in non-indexed
	// mode to ensure fair benchmark comparisons and to track down
	// unintended compilation differences.
	if false {
		pstate.inlcalls(fn)
	}
	pstate.escAnalyze([]*Node{fn}, false)

	pstate.Curfn = nil
	pstate.funccompile(fn)
}

func (pstate *PackageState) paramNnames(ft *types.Type) []*Node {
	args := make([]*Node, ft.NumParams(pstate.types))
	for i, f := range ft.Params(pstate.types).FieldSlice(pstate.types) {
		args[i] = asNode(f.Nname)
	}
	return args
}

func (pstate *PackageState) hashmem(t *types.Type) *Node {
	sym := pstate.Runtimepkg.Lookup(pstate.types, "memhash")

	n := pstate.newname(sym)
	n.SetClass(PFUNC)
	n.Type = pstate.functype(nil, []*Node{
		pstate.anonfield(pstate.types.NewPtr(t)),
		pstate.anonfield(pstate.types.Types[TUINTPTR]),
		pstate.anonfield(pstate.types.Types[TUINTPTR]),
	}, []*Node{
		pstate.anonfield(pstate.types.Types[TUINTPTR]),
	})
	return n
}

func (pstate *PackageState) ifacelookdot(s *types.Sym, t *types.Type, ignorecase bool) (m *types.Field, followptr bool) {
	if t == nil {
		return nil, false
	}

	path, ambig := pstate.dotpath(s, t, &m, ignorecase)
	if path == nil {
		if ambig {
			pstate.yyerror("%v.%v is ambiguous", t, s)
		}
		return nil, false
	}

	for _, d := range path {
		if d.field.Type.IsPtr() {
			followptr = true
			break
		}
	}

	if m.Type.Etype != TFUNC || m.Type.Recv(pstate.types) == nil {
		pstate.yyerror("%v.%v is a field, not a method", t, s)
		return nil, followptr
	}

	return m, followptr
}

func (pstate *PackageState) implements(t, iface *types.Type, m, samename **types.Field, ptr *int) bool {
	t0 := t
	if t == nil {
		return false
	}

	if t.IsInterface() {
		i := 0
		tms := t.Fields(pstate.types).Slice()
		for _, im := range iface.Fields(pstate.types).Slice() {
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
			if !pstate.eqtype(tm.Type, im.Type) {
				*m = im
				*samename = tm
				*ptr = 0
				return false
			}
		}

		return true
	}

	t = pstate.methtype(t)
	var tms []*types.Field
	if t != nil {
		pstate.expandmeth(t)
		tms = t.AllMethods().Slice()
	}
	i := 0
	for _, im := range iface.Fields(pstate.types).Slice() {
		if im.Broke() {
			continue
		}
		for i < len(tms) && tms[i].Sym != im.Sym {
			i++
		}
		if i == len(tms) {
			*m = im
			*samename, _ = pstate.ifacelookdot(im.Sym, t, true)
			*ptr = 0
			return false
		}
		tm := tms[i]
		if tm.Nointerface() || !pstate.eqtype(tm.Type, im.Type) {
			*m = im
			*samename = tm
			*ptr = 0
			return false
		}
		followptr := tm.Embedded == 2

		// if pointer receiver in method,
		// the method does not exist for value types.
		rcvr := tm.Type.Recv(pstate.types).Type
		if rcvr.IsPtr() && !t0.IsPtr() && !followptr && !pstate.isifacemethod(tm.Type) {
			if false && pstate.Debug['r'] != 0 {
				pstate.yyerror("interface pointer mismatch")
			}

			*m = im
			*samename = nil
			*ptr = 1
			return false
		}
	}

	// We're going to emit an OCONVIFACE.
	// Call itabname so that (t, iface)
	// gets added to itabs early, which allows
	// us to de-virtualize calls through this
	// type/interface pair later. See peekitabs in reflect.go
	if pstate.isdirectiface(t0) && !iface.IsEmptyInterface(pstate.types) {
		pstate.itabname(t0, iface)
	}
	return true
}

func (pstate *PackageState) listtreecopy(l []*Node, pos src.XPos) []*Node {
	var out []*Node
	for _, n := range l {
		out = append(out, pstate.treecopy(n, pos))
	}
	return out
}

func (pstate *PackageState) liststmt(l []*Node) *Node {
	n := pstate.nod(OBLOCK, nil, nil)
	n.List.Set(l)
	if len(l) != 0 {
		n.Pos = l[0].Pos
	}
	return n
}

func (l Nodes) asblock(pstate *PackageState) *Node {
	n := pstate.nod(OBLOCK, nil, nil)
	n.List = l
	if l.Len() != 0 {
		n.Pos = l.First().Pos
	}
	return n
}

func (pstate *PackageState) ngotype(n *Node) *types.Sym {
	if n.Type != nil {
		return pstate.typenamesym(n.Type)
	}
	return nil
}

// The result of addinit MUST be assigned back to n, e.g.
// 	n.Left = addinit(n.Left, init)
func (pstate *PackageState) addinit(n *Node, init []*Node) *Node {
	if len(init) == 0 {
		return n
	}
	if n.mayBeShared() {
		// Introduce OCONVNOP to hold init list.
		n = pstate.nod(OCONVNOP, n, nil)
		n.Type = n.Left.Type
		n.SetTypecheck(1)
	}

	n.Ninit.Prepend(init...)
	n.SetHasCall(true)
	return n
}

func (pstate *PackageState) isbadimport(path string, allowSpace bool) bool {
	if strings.Contains(path, "\x00") {
		pstate.yyerror("import path contains NUL")
		return true
	}

	for _, ri := range pstate.reservedimports {
		if path == ri {
			pstate.yyerror("import path %q is reserved and cannot be used", path)
			return true
		}
	}

	for _, r := range path {
		if r == utf8.RuneError {
			pstate.yyerror("import path contains invalid UTF-8 sequence: %q", path)
			return true
		}

		if r < 0x20 || r == 0x7f {
			pstate.yyerror("import path contains control character: %q", path)
			return true
		}

		if r == '\\' {
			pstate.yyerror("import path contains backslash; use slash: %q", path)
			return true
		}

		if !allowSpace && unicode.IsSpace(r) {
			pstate.yyerror("import path contains space character: %q", path)
			return true
		}

		if strings.ContainsRune("!\"#$%&'()*,:;<=>?[]^`{|}", r) {
			pstate.yyerror("import path contains invalid character '%c': %q", r, path)
			return true
		}
	}

	return false
}

func (pstate *PackageState) checknil(x *Node, init *Nodes) {
	x = pstate.walkexpr(x, nil) // caller has not done this yet
	if x.Type.IsInterface() {
		x = pstate.nod(OITAB, x, nil)
		x = pstate.typecheck(x, Erv)
	}

	n := pstate.nod(OCHECKNIL, x, nil)
	n.SetTypecheck(1)
	init.Append(n)
}

// Can this type be stored directly in an interface word?
// Yes, if the representation is a single pointer.
func (pstate *PackageState) isdirectiface(t *types.Type) bool {
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
		// Array of 1 direct iface type can be direct.
		return t.NumElem(pstate.types) == 1 && pstate.isdirectiface(t.Elem(pstate.types))

	case TSTRUCT:
		// Struct with 1 field of direct iface type can be direct.
		return t.NumFields(pstate.types) == 1 && pstate.isdirectiface(t.Field(pstate.types, 0).Type)
	}

	return false
}

// itabType loads the _type field from a runtime.itab struct.
func (pstate *PackageState) itabType(itab *Node) *Node {
	typ := pstate.nodSym(ODOTPTR, itab, nil)
	typ.Type = pstate.types.NewPtr(pstate.types.Types[TUINT8])
	typ.SetTypecheck(1)
	typ.Xoffset = int64(pstate.Widthptr) // offset of _type in runtime.itab
	typ.SetBounded(true)                 // guaranteed not to fault
	return typ
}

// ifaceData loads the data field from an interface.
// The concrete type must be known to have type t.
// It follows the pointer if !isdirectiface(t).
func (pstate *PackageState) ifaceData(n *Node, t *types.Type) *Node {
	ptr := pstate.nodSym(OIDATA, n, nil)
	if pstate.isdirectiface(t) {
		ptr.Type = t
		ptr.SetTypecheck(1)
		return ptr
	}
	ptr.Type = pstate.types.NewPtr(t)
	ptr.SetBounded(true)
	ptr.SetTypecheck(1)
	ind := pstate.nod(OIND, ptr, nil)
	ind.Type = t
	ind.SetTypecheck(1)
	return ind
}

// This file implements TestFormats; a test that verifies
// format strings in the compiler (this directory and all
// subdirectories, recursively).
//
// TestFormats finds potential (Printf, etc.) format strings.
// If they are used in a call, the format verbs are verified
// based on the matching argument type against a precomputed
// table of valid formats. The knownFormats table can be used
// to automatically rewrite format strings with the -u flag.
//
// A new knownFormats table based on the found formats is printed
// when the test is run in verbose mode (-v flag). The table
// needs to be updated whenever a new (type, format) combination
// is found and the format verb is not 'v' or 'T' (as in "%v" or
// "%T").
//
// Run as: go test -run Formats [-u][-v]
//
// Known bugs:
// - indexed format strings ("%[2]s", etc.) are not supported
//   (the test will fail)
// - format strings that are not simple string literals cannot
//   be updated automatically
//   (the test will fail with respective warnings)
// - format strings in _test packages outside the current
//   package are not processed
//   (the test will report those files)
//
package main_test

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/dave/golib/src/internal/testenv"
	"go/ast"
	"go/build"
	"go/constant"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

var update = flag.Bool("u", false, "update format strings")

// The following variables collect information across all processed files.
var (
	fset          = token.NewFileSet()
	formatStrings = make(map[*ast.BasicLit]bool)      // set of all potential format strings found
	foundFormats  = make(map[string]bool)             // set of all formats found
	callSites     = make(map[*ast.CallExpr]*callSite) // map of all calls
)

// A File is a corresponding (filename, ast) pair.
type File struct {
	name string
	ast  *ast.File
}

func TestFormats(t *testing.T) {
	testenv.MustHaveGoBuild(t)

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() == "testdata" {
				return filepath.SkipDir
			}

			importPath := filepath.Join("github.com/dave/golib/src/cmd/compile", path)
			if blacklistedPackages[filepath.ToSlash(importPath)] {
				return filepath.SkipDir
			}

			pkg, err := build.Import(importPath, path, 0)
			if err != nil {
				if _, ok := err.(*build.NoGoError); ok {
					return nil
				}
				t.Fatal(err)
			}
			collectPkgFormats(t, pkg)
		}
		return nil
	})

	updatedFiles := make(map[string]File)
	for _, p := range callSites {

		out := formatReplace(p.str, func(index int, in string) string {
			if in == "*" {
				return in
			}

			typ := p.types[index]
			format := typ + " " + in

			out, known := knownFormats[format]

			_, found := foundFormats[format]
			if !found {
				foundFormats[format] = true
			}

			if !known && !found && in != "%v" && in != "%T" {
				t.Errorf("%s: unknown format %q for %s argument", posString(p.arg), in, typ)
			}

			if out == "" {
				out = in
			}
			return out
		})

		if out != p.str {

			lit, ok := p.arg.(*ast.BasicLit)
			if !ok {
				delete(callSites, p.call)
				continue
			}

			if testing.Verbose() {
				fmt.Printf("%s:\n\t- %q\n\t+ %q\n", posString(p.arg), p.str, out)
			}

			index := -1
			for i, arg := range p.call.Args {
				if p.arg == arg {
					index = i
					break
				}
			}
			if index < 0 {

				panic("internal error: matching argument not found")
			}

			new := *lit
			new.Value = strconv.Quote(out)
			p.call.Args[index] = &new
			updatedFiles[p.file.name] = p.file
		}
	}

	// write dirty files back
	var filesUpdated bool
	if len(updatedFiles) > 0 && *update {
		for _, file := range updatedFiles {
			var buf bytes.Buffer
			if err := format.Node(&buf, fset, file.ast); err != nil {
				t.Errorf("WARNING: formatting %s failed: %v", file.name, err)
				continue
			}
			if err := ioutil.WriteFile(file.name, buf.Bytes(), 0x666); err != nil {
				t.Errorf("WARNING: writing %s failed: %v", file.name, err)
				continue
			}
			fmt.Printf("updated %s\n", file.name)
			filesUpdated = true
		}
	}

	if len(callSites) > 0 && testing.Verbose() {
		set := make(map[string]bool)
		for _, p := range callSites {
			set[nodeString(p.call.Fun)] = true
		}
		var list []string
		for s := range set {
			list = append(list, s)
		}
		fmt.Println("\nFunctions")
		printList(list)
	}

	if len(foundFormats) > 0 && testing.Verbose() {
		var list []string
		for s := range foundFormats {
			list = append(list, fmt.Sprintf("%q: \"\",", s))
		}
		fmt.Println("\nvar knownFormats = map[string]string{")
		printList(list)
		fmt.Println("}")
	}

	if !testing.Verbose() && !*update {
		var mismatch bool
		for s := range foundFormats {
			if _, ok := knownFormats[s]; !ok {
				mismatch = true
				break
			}
		}
		if !mismatch {
			for s := range knownFormats {
				if _, ok := foundFormats[s]; !ok {
					mismatch = true
					break
				}
			}
		}
		if mismatch {
			t.Errorf("knownFormats is out of date; please 'go test -v fmt_test.go > foo', then extract new definition of knownFormats from foo")
		}
	}

	for _, p := range callSites {
		if lit, ok := p.arg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			if formatStrings[lit] {

				delete(formatStrings, lit)
			} else {

				panic(fmt.Sprintf("internal error: format string not found (%s)", posString(lit)))
			}
		}
	}

	if len(formatStrings) > 0 && filesUpdated {
		var list []string
		for lit := range formatStrings {
			list = append(list, fmt.Sprintf("%s: %s", posString(lit), nodeString(lit)))
		}
		fmt.Println("\nWARNING: Potentially missed format strings")
		printList(list)
		t.Fail()
	}

	fmt.Println()
}

// A callSite describes a function call that appears to contain
// a format string.
type callSite struct {
	file  File
	call  *ast.CallExpr // call containing the format string
	arg   ast.Expr      // format argument (string literal or constant)
	str   string        // unquoted format string
	types []string      // argument types
}

func collectPkgFormats(t *testing.T, pkg *build.Package) {
	// collect all files
	var filenames []string
	filenames = append(filenames, pkg.GoFiles...)
	filenames = append(filenames, pkg.CgoFiles...)
	filenames = append(filenames, pkg.TestGoFiles...)

	for _, name := range pkg.XTestGoFiles {

		if name != "fmt_test.go" && testing.Verbose() {
			fmt.Printf("WARNING: %s not processed\n", filepath.Join(pkg.Dir, name))
		}
	}

	for i, name := range filenames {
		filenames[i] = filepath.Join(pkg.Dir, name)
	}

	files := make([]*ast.File, len(filenames))
	for i, filename := range filenames {
		f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
		if err != nil {
			t.Fatal(err)
		}
		files[i] = f
	}

	conf := types.Config{Importer: importer.Default()}
	etypes := make(map[ast.Expr]types.TypeAndValue)
	if _, err := conf.Check(pkg.ImportPath, fset, files, &types.Info{Types: etypes}); err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		ast.Inspect(file, func(n ast.Node) bool {
			if s, ok := stringLit(n); ok && isFormat(s) {
				formatStrings[n.(*ast.BasicLit)] = true
			}
			return true
		})
	}

	for index, file := range files {
		ast.Inspect(file, func(n ast.Node) bool {
			if call, ok := n.(*ast.CallExpr); ok {

				if blacklistedFunctions[nodeString(call.Fun)] {
					return true
				}

				for i, arg := range call.Args {
					if s, ok := stringVal(etypes[arg]); ok && isFormat(s) {

						n := numFormatArgs(s)
						if i+1+n > len(call.Args) {
							t.Errorf("%s: not enough format args (blacklist %s?)", posString(call), nodeString(call.Fun))
							break
						}

						argTypes := make([]string, n)
						for i, arg := range call.Args[len(call.Args)-n:] {
							if tv, ok := etypes[arg]; ok {
								argTypes[i] = typeString(tv.Type)
							}
						}

						if callSites[call] != nil {
							panic("internal error: file processed twice?")
						}
						callSites[call] = &callSite{
							file:  File{filenames[index], file},
							call:  call,
							arg:   arg,
							str:   s,
							types: argTypes,
						}
						break
					}
				}
			}
			return true
		})
	}
}

// printList prints list in sorted order.
func printList(list []string) {
	sort.Strings(list)
	for _, s := range list {
		fmt.Println("\t", s)
	}
}

// posString returns a string representation of n's position
// in the form filename:line:col: .
func posString(n ast.Node) string {
	if n == nil {
		return ""
	}
	return fset.Position(n.Pos()).String()
}

// nodeString returns a string representation of n.
func nodeString(n ast.Node) string {
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, n); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

// typeString returns a string representation of n.
func typeString(typ types.Type) string {
	return filepath.ToSlash(typ.String())
}

// stringLit returns the unquoted string value and true if
// n represents a string literal; otherwise it returns ""
// and false.
func stringLit(n ast.Node) (string, bool) {
	if lit, ok := n.(*ast.BasicLit); ok && lit.Kind == token.STRING {
		s, err := strconv.Unquote(lit.Value)
		if err != nil {
			log.Fatal(err)
		}
		return s, true
	}
	return "", false
}

// stringVal returns the (unquoted) string value and true if
// tv is a string constant; otherwise it returns "" and false.
func stringVal(tv types.TypeAndValue) (string, bool) {
	if tv.IsValue() && tv.Value != nil && tv.Value.Kind() == constant.String {
		return constant.StringVal(tv.Value), true
	}
	return "", false
}

// formatIter iterates through the string s in increasing
// index order and calls f for each format specifier '%..v'.
// The arguments for f describe the specifier's index range.
// If a format specifier contains a "*", f is called with
// the index range for "*" alone, before being called for
// the entire specifier. The result of f is the index of
// the rune at which iteration continues.
func formatIter(s string, f func(i, j int) int) {
	i := 0
	var r rune // current rune

	next := func() {
		r1, w := utf8.DecodeRuneInString(s[i:])
		if w == 0 {
			r1 = -1
		}
		r = r1
		i += w
	}

	flags := func() {
		for r == ' ' || r == '#' || r == '+' || r == '-' || r == '0' {
			next()
		}
	}

	index := func() {
		if r == '[' {
			log.Fatalf("cannot handle indexed arguments: %s", s)
		}
	}

	digits := func() {
		index()
		if r == '*' {
			i = f(i-1, i)
			next()
			return
		}
		for '0' <= r && r <= '9' {
			next()
		}
	}

	for next(); r >= 0; next() {
		if r == '%' {
			i0 := i
			next()
			flags()
			digits()
			if r == '.' {
				next()
				digits()
			}
			index()

			if 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' {
				i = f(i0-1, i)
			}
		}
	}
}

// isFormat reports whether s contains format specifiers.
func isFormat(s string) (yes bool) {
	formatIter(s, func(i, j int) int {
		yes = true
		return len(s)
	})
	return
}

// oneFormat reports whether s is exactly one format specifier.
func oneFormat(s string) (yes bool) {
	formatIter(s, func(i, j int) int {
		yes = i == 0 && j == len(s)
		return j
	})
	return
}

// numFormatArgs returns the number of format specifiers in s.
func numFormatArgs(s string) int {
	count := 0
	formatIter(s, func(i, j int) int {
		count++
		return j
	})
	return count
}

// formatReplace replaces the i'th format specifier s in the incoming
// string in with the result of f(i, s) and returns the new string.
func formatReplace(in string, f func(i int, s string) string) string {
	var buf []byte
	i0 := 0
	index := 0
	formatIter(in, func(i, j int) int {
		if sub := in[i:j]; sub != "*" {
			buf = append(buf, in[i0:i]...)
			buf = append(buf, f(index, sub)...)
			i0 = j
		}
		index++
		return j
	})
	return string(append(buf, in[i0:]...))
}

// blacklistedPackages is the set of packages which can
// be ignored.
var blacklistedPackages = map[string]bool{}

// blacklistedFunctions is the set of functions which may have
// format-like arguments but which don't do any formatting and
// thus may be ignored.
var blacklistedFunctions = map[string]bool{}

func init() {

	for key, val := range knownFormats {

		i := strings.Index(key, "%")
		if i < 0 || !oneFormat(key[i:]) {
			log.Fatalf("incorrect knownFormats key: %q", key)
		}

		if val != "" && !oneFormat(val) {
			log.Fatalf("incorrect knownFormats value: %q (key = %q)", val, key)
		}
	}
}

// knownFormats entries are of the form "typename format" -> "newformat".
// An absent entry means that the format is not recognized as valid.
// An empty new format means that the format should remain unchanged.
// To print out a new table, run: go test -run Formats -v.
var knownFormats = map[string]string{
	"*bytes.Buffer %s":                                                          "",
	"*github.com/dave/golib/src/cmd/compile/internal/gc.Mpflt %v":               "",
	"*github.com/dave/golib/src/cmd/compile/internal/gc.Mpint %v":               "",
	"*github.com/dave/golib/src/cmd/compile/internal/gc.Node %#v":               "",
	"*github.com/dave/golib/src/cmd/compile/internal/gc.Node %+S":               "",
	"*github.com/dave/golib/src/cmd/compile/internal/gc.Node %+v":               "",
	"*github.com/dave/golib/src/cmd/compile/internal/gc.Node %0j":               "",
	"*github.com/dave/golib/src/cmd/compile/internal/gc.Node %L":                "",
	"*github.com/dave/golib/src/cmd/compile/internal/gc.Node %S":                "",
	"*github.com/dave/golib/src/cmd/compile/internal/gc.Node %j":                "",
	"*github.com/dave/golib/src/cmd/compile/internal/gc.Node %p":                "",
	"*github.com/dave/golib/src/cmd/compile/internal/gc.Node %v":                "",
	"*github.com/dave/golib/src/cmd/compile/internal/ssa.Block %s":              "",
	"*github.com/dave/golib/src/cmd/compile/internal/ssa.Block %v":              "",
	"*github.com/dave/golib/src/cmd/compile/internal/ssa.Func %s":               "",
	"*github.com/dave/golib/src/cmd/compile/internal/ssa.Func %v":               "",
	"*github.com/dave/golib/src/cmd/compile/internal/ssa.Register %s":           "",
	"*github.com/dave/golib/src/cmd/compile/internal/ssa.Register %v":           "",
	"*github.com/dave/golib/src/cmd/compile/internal/ssa.SparseTreeNode %v":     "",
	"*github.com/dave/golib/src/cmd/compile/internal/ssa.Value %s":              "",
	"*github.com/dave/golib/src/cmd/compile/internal/ssa.Value %v":              "",
	"*github.com/dave/golib/src/cmd/compile/internal/ssa.sparseTreeMapEntry %v": "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Field %p":            "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Field %v":            "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Sym %+v":             "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Sym %0S":             "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Sym %S":              "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Sym %p":              "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Sym %v":              "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Type %#v":            "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Type %+v":            "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Type %-S":            "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Type %0S":            "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Type %L":             "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Type %S":             "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Type %p":             "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Type %s":             "",
	"*github.com/dave/golib/src/cmd/compile/internal/types.Type %v":             "",
	"*github.com/dave/golib/src/cmd/internal/obj.Addr %v":                       "",
	"*github.com/dave/golib/src/cmd/internal/obj.LSym %v":                       "",
	"*math/big.Int %#x":                                                         "",
	"*math/big.Int %s":                                                          "",
	"*math/big.Int %v":                                                          "",
	"[16]byte %x":                                                               "",
	"[]*github.com/dave/golib/src/cmd/compile/internal/gc.Node %v":   "",
	"[]*github.com/dave/golib/src/cmd/compile/internal/ssa.Block %v": "",
	"[]*github.com/dave/golib/src/cmd/compile/internal/ssa.Value %v": "",
	"[][]string %q": "",
	"[]byte %s":     "",
	"[]byte %x":     "",
	"[]github.com/dave/golib/src/cmd/compile/internal/ssa.Edge %v":      "",
	"[]github.com/dave/golib/src/cmd/compile/internal/ssa.ID %v":        "",
	"[]github.com/dave/golib/src/cmd/compile/internal/ssa.posetNode %v": "",
	"[]github.com/dave/golib/src/cmd/compile/internal/ssa.posetUndo %v": "",
	"[]github.com/dave/golib/src/cmd/compile/internal/syntax.token %s":  "",
	"[]string %v": "",
	"[]uint32 %v": "",
	"bool %v":     "",
	"byte %08b":   "",
	"byte %c":     "",
	"byte %v":     "",
	"github.com/dave/golib/src/cmd/compile/internal/arm.shift %d":            "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Class %d":             "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Class %s":             "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Class %v":             "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Ctype %d":             "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Ctype %v":             "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Level %d":             "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Level %v":             "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Nodes %#v":            "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Nodes %+v":            "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Nodes %.v":            "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Nodes %v":             "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Op %#v":               "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Op %v":                "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Val %#v":              "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Val %T":               "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.Val %v":               "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.fmtMode %d":           "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.initKind %d":          "",
	"github.com/dave/golib/src/cmd/compile/internal/gc.itag %v":              "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.BranchPrediction %d": "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.Edge %v":             "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.GCNode %v":           "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.ID %d":               "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.ID %v":               "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.LocPair %s":          "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.LocalSlot %s":        "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.LocalSlot %v":        "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.Location %T":         "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.Location %s":         "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.Op %s":               "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.Op %v":               "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.ValAndOff %s":        "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.domain %v":           "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.posetNode %v":        "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.posetTestOp %v":      "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.rbrank %d":           "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.regMask %d":          "",
	"github.com/dave/golib/src/cmd/compile/internal/ssa.register %d":         "",
	"github.com/dave/golib/src/cmd/compile/internal/syntax.Expr %#v":         "",
	"github.com/dave/golib/src/cmd/compile/internal/syntax.Node %T":          "",
	"github.com/dave/golib/src/cmd/compile/internal/syntax.Operator %s":      "",
	"github.com/dave/golib/src/cmd/compile/internal/syntax.Pos %s":           "",
	"github.com/dave/golib/src/cmd/compile/internal/syntax.Pos %v":           "",
	"github.com/dave/golib/src/cmd/compile/internal/syntax.position %s":      "",
	"github.com/dave/golib/src/cmd/compile/internal/syntax.token %q":         "",
	"github.com/dave/golib/src/cmd/compile/internal/syntax.token %s":         "",
	"github.com/dave/golib/src/cmd/compile/internal/types.EType %d":          "",
	"github.com/dave/golib/src/cmd/compile/internal/types.EType %s":          "",
	"github.com/dave/golib/src/cmd/compile/internal/types.EType %v":          "",
	"error %v":        "",
	"float64 %.2f":    "",
	"float64 %.3f":    "",
	"float64 %.6g":    "",
	"float64 %g":      "",
	"int %-12d":       "",
	"int %-6d":        "",
	"int %-8o":        "",
	"int %02d":        "",
	"int %6d":         "",
	"int %c":          "",
	"int %d":          "",
	"int %v":          "",
	"int %x":          "",
	"int16 %d":        "",
	"int16 %x":        "",
	"int32 %d":        "",
	"int32 %v":        "",
	"int32 %x":        "",
	"int64 %+d":       "",
	"int64 %-10d":     "",
	"int64 %.5d":      "",
	"int64 %X":        "",
	"int64 %d":        "",
	"int64 %v":        "",
	"int64 %x":        "",
	"int8 %d":         "",
	"int8 %x":         "",
	"interface{} %#v": "",
	"interface{} %T":  "",
	"interface{} %q":  "",
	"interface{} %s":  "",
	"interface{} %v":  "",
	"map[*github.com/dave/golib/src/cmd/compile/internal/gc.Node]*github.com/dave/golib/src/cmd/compile/internal/ssa.Value %v": "",
	"map[github.com/dave/golib/src/cmd/compile/internal/ssa.ID]uint32 %v":                                                      "",
	"reflect.Type %s":  "",
	"rune %#U":         "",
	"rune %c":          "",
	"string %-*s":      "",
	"string %-16s":     "",
	"string %-6s":      "",
	"string %.*s":      "",
	"string %q":        "",
	"string %s":        "",
	"string %v":        "",
	"time.Duration %d": "",
	"time.Duration %v": "",
	"uint %04x":        "",
	"uint %5d":         "",
	"uint %d":          "",
	"uint %x":          "",
	"uint16 %d":        "",
	"uint16 %v":        "",
	"uint16 %x":        "",
	"uint32 %d":        "",
	"uint32 %v":        "",
	"uint32 %x":        "",
	"uint64 %08x":      "",
	"uint64 %d":        "",
	"uint64 %x":        "",
	"uint8 %d":         "",
	"uint8 %x":         "",
	"uintptr %d":       "",
}

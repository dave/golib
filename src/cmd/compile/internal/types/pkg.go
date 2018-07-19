package types

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"sort"
)

// pkgMap maps a package path to a package.

// MaxPkgHeight is a height greater than any likely package height.
const MaxPkgHeight = 1e9

type Pkg struct {
	Path    string // string literal used in import statement, e.g. "runtime/internal/sys"
	Name    string // package name, e.g. "sys"
	Prefix  string // escaped path for use in symbol table
	Syms    map[string]*Sym
	Pathsym *obj.LSym

	// Height is the package's height in the import graph. Leaf
	// packages (i.e., packages with no imports) have height 0,
	// and all other packages have height 1 plus the maximum
	// height of their imported packages.
	Height int

	Imported bool // export data of this package was parsed
	Direct   bool // imported directly
}

// NewPkg returns a new Pkg for the given package path and name.
// Unless name is the empty string, if the package exists already,
// the existing package name and the provided name must match.
func (psess *PackageSession) NewPkg(path, name string) *Pkg {
	if p := psess.pkgMap[path]; p != nil {
		if name != "" && p.Name != name {
			panic(fmt.Sprintf("conflicting package names %s and %s for path %q", p.Name, name, path))
		}
		return p
	}

	p := new(Pkg)
	p.Path = path
	p.Name = name
	p.Prefix = objabi.PathToPrefix(path)
	p.Syms = make(map[string]*Sym)
	psess.
		pkgMap[path] = p

	return p
}

// ImportedPkgList returns the list of directly imported packages.
// The list is sorted by package path.
func (psess *PackageSession) ImportedPkgList() []*Pkg {
	var list []*Pkg
	for _, p := range psess.pkgMap {
		if p.Direct {
			list = append(list, p)
		}
	}
	sort.Sort(byPath(list))
	return list
}

type byPath []*Pkg

func (a byPath) Len() int           { return len(a) }
func (a byPath) Less(i, j int) bool { return a[i].Path < a[j].Path }
func (a byPath) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (pkg *Pkg) Lookup(psess *PackageSession, name string) *Sym {
	s, _ := pkg.LookupOK(psess, name)
	return s
}

// LookupOK looks up name in pkg and reports whether it previously existed.
func (pkg *Pkg) LookupOK(psess *PackageSession, name string) (s *Sym, existed bool) {

	if pkg == nil {
		pkg = psess.nopkg
	}
	if s := pkg.Syms[name]; s != nil {
		return s, true
	}

	s = &Sym{
		Name: name,
		Pkg:  pkg,
	}
	if name == "init" {
		psess.
			InitSyms = append(psess.InitSyms, s)
	}
	pkg.Syms[name] = s
	return s, false
}

func (pkg *Pkg) LookupBytes(psess *PackageSession, name []byte) *Sym {

	if pkg == nil {
		pkg = psess.nopkg
	}
	if s := pkg.Syms[string(name)]; s != nil {
		return s
	}
	str := psess.InternString(name)
	return pkg.Lookup(psess, str)
}

// protects internedStrings

func (psess *PackageSession) InternString(b []byte) string {
	psess.
		internedStringsmu.Lock()
	s, ok := psess.internedStrings[string(b)]
	if !ok {
		s = string(b)
		psess.
			internedStrings[s] = s
	}
	psess.
		internedStringsmu.Unlock()
	return s
}

// CleanroomDo invokes f in an environment with with no preexisting packages.
// For testing of import/export only.
func (psess *PackageSession) CleanroomDo(f func()) {
	saved := psess.pkgMap
	psess.
		pkgMap = make(map[string]*Pkg)
	f()
	psess.
		pkgMap = saved
}

// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"sort"
	"sync"
)

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
func (pstate *PackageState) NewPkg(path, name string) *Pkg {
	if p := pstate.pkgMap[path]; p != nil {
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
	pstate.pkgMap[path] = p

	return p
}

// ImportedPkgList returns the list of directly imported packages.
// The list is sorted by package path.
func (pstate *PackageState) ImportedPkgList() []*Pkg {
	var list []*Pkg
	for _, p := range pstate.pkgMap {
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

func (pkg *Pkg) Lookup(pstate *PackageState, name string) *Sym {
	s, _ := pkg.LookupOK(pstate, name)
	return s
}

// LookupOK looks up name in pkg and reports whether it previously existed.
func (pkg *Pkg) LookupOK(pstate *PackageState, name string) (s *Sym, existed bool) {
	// TODO(gri) remove this check in favor of specialized lookup
	if pkg == nil {
		pkg = pstate.nopkg
	}
	if s := pkg.Syms[name]; s != nil {
		return s, true
	}

	s = &Sym{
		Name: name,
		Pkg:  pkg,
	}
	if name == "init" {
		pstate.InitSyms = append(pstate.InitSyms, s)
	}
	pkg.Syms[name] = s
	return s, false
}

func (pkg *Pkg) LookupBytes(pstate *PackageState, name []byte) *Sym {
	// TODO(gri) remove this check in favor of specialized lookup
	if pkg == nil {
		pkg = pstate.nopkg
	}
	if s := pkg.Syms[string(name)]; s != nil {
		return s
	}
	str := pstate.InternString(name)
	return pkg.Lookup(pstate, str)
}

func (pstate *PackageState) InternString(b []byte) string {
	pstate.internedStringsmu.Lock()
	s, ok := pstate.internedStrings[string(b)] // string(b) here doesn't allocate
	if !ok {
		s = string(b)
		pstate.internedStrings[s] = s
	}
	pstate.internedStringsmu.Unlock()
	return s
}

// CleanroomDo invokes f in an environment with with no preexisting packages.
// For testing of import/export only.
func (pstate *PackageState) CleanroomDo(f func()) {
	saved := pstate.pkgMap
	pstate.pkgMap = make(map[string]*Pkg)
	f()
	pstate.pkgMap = saved
}

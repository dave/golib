// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import "github.com/dave/golib/src/cmd/internal/src"

// A dsym stores a symbol's shadowed declaration so that it can be
// restored once the block scope ends.
type dsym struct {
	sym        *Sym // sym == nil indicates stack mark
	def        *Node
	block      int32
	lastlineno src.XPos // last declaration for diagnostic
}

// Pushdcl pushes the current declaration for symbol s (if any) so that
// it can be shadowed by a new declaration within a nested block scope.
func (pstate *PackageState) Pushdcl(s *Sym) {
	pstate.dclstack = append(pstate.dclstack, dsym{
		sym:        s,
		def:        s.Def,
		block:      s.Block,
		lastlineno: s.Lastlineno,
	})
}

// Popdcl pops the innermost block scope and restores all symbol declarations
// to their previous state.
func (pstate *PackageState) Popdcl() {
	for i := len(pstate.dclstack); i > 0; i-- {
		d := &pstate.dclstack[i-1]
		s := d.sym
		if s == nil {
			// pop stack mark
			pstate.Block = d.block
			pstate.dclstack = pstate.dclstack[:i-1]
			return
		}

		s.Def = d.def
		s.Block = d.block
		s.Lastlineno = d.lastlineno

		// Clear dead pointer fields.
		d.sym = nil
		d.def = nil
	}
	pstate.Fatalf("popdcl: no stack mark")
}

// Markdcl records the start of a new block scope for declarations.
func (pstate *PackageState) Markdcl() {
	pstate.dclstack = append(pstate.dclstack, dsym{
		sym:   nil, // stack mark
		block: pstate.Block,
	})
	pstate.blockgen++
	pstate.Block = pstate.blockgen
}

func (pstate *PackageState) IsDclstackValid() bool {
	for _, d := range pstate.dclstack {
		if d.sym == nil {
			return false
		}
	}
	return true
}

// PkgDef returns the definition associated with s at package scope.
func (s *Sym) PkgDef(pstate *PackageState) *Node {
	// Look for outermost saved declaration, which must be the
	// package scope definition, if present.
	for _, d := range pstate.dclstack {
		if s == d.sym {
			return d.def
		}
	}

	// Otherwise, the declaration hasn't been shadowed within a
	// function scope.
	return s.Def
}

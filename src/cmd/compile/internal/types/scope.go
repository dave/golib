package types

import "github.com/dave/golib/src/cmd/internal/src"

// max block number
// current block number

// A dsym stores a symbol's shadowed declaration so that it can be
// restored once the block scope ends.
type dsym struct {
	sym        *Sym // sym == nil indicates stack mark
	def        *Node
	block      int32
	lastlineno src.XPos // last declaration for diagnostic
}

// dclstack maintains a stack of shadowed symbol declarations so that
// Popdcl can restore their declarations when a block scope ends.

// Pushdcl pushes the current declaration for symbol s (if any) so that
// it can be shadowed by a new declaration within a nested block scope.
func (psess *PackageSession) Pushdcl(s *Sym) {
	psess.
		dclstack = append(psess.dclstack, dsym{
		sym:        s,
		def:        s.Def,
		block:      s.Block,
		lastlineno: s.Lastlineno,
	})
}

// Popdcl pops the innermost block scope and restores all symbol declarations
// to their previous state.
func (psess *PackageSession) Popdcl() {
	for i := len(psess.dclstack); i > 0; i-- {
		d := &psess.dclstack[i-1]
		s := d.sym
		if s == nil {
			psess.
				Block = d.block
			psess.
				dclstack = psess.dclstack[:i-1]
			return
		}

		s.Def = d.def
		s.Block = d.block
		s.Lastlineno = d.lastlineno

		d.sym = nil
		d.def = nil
	}
	psess.
		Fatalf("popdcl: no stack mark")
}

// Markdcl records the start of a new block scope for declarations.
func (psess *PackageSession) Markdcl() {
	psess.
		dclstack = append(psess.dclstack, dsym{
		sym:   nil,
		block: psess.Block,
	})
	psess.
		blockgen++
	psess.
		Block = psess.blockgen
}

func (psess *PackageSession) IsDclstackValid() bool {
	for _, d := range psess.dclstack {
		if d.sym == nil {
			return false
		}
	}
	return true
}

// PkgDef returns the definition associated with s at package scope.
func (s *Sym) PkgDef(psess *PackageSession,) *Node {

	for _, d := range psess.dclstack {
		if s == d.sym {
			return d.def
		}
	}

	return s.Def
}

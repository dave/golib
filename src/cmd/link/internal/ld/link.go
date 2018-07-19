package ld

import (
	"bufio"
	"debug/elf"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/sym"
)

type Shlib struct {
	Path            string
	Hash            []byte
	Deps            []string
	File            *elf.File
	gcdataAddresses map[*sym.Symbol]uint64
}

// Link holds the context for writing object code from a compiler
// or for reading that input into the linker.
type Link struct {
	Out *OutBuf

	Syms *sym.Symbols

	Arch      *sys.Arch
	Debugvlog int
	Bso       *bufio.Writer

	Loaded bool // set after all inputs have been loaded as symbols

	IsELF    bool
	HeadType objabi.HeadType

	linkShared bool // link against installed Go shared libraries
	LinkMode   LinkMode
	BuildMode  BuildMode

	Tlsg         *sym.Symbol
	Libdir       []string
	Library      []*sym.Library
	LibraryByPkg map[string]*sym.Library
	Shlibs       []Shlib
	Tlsoffset    int
	Textp        []*sym.Symbol
	Filesyms     []*sym.Symbol
	Moduledata   *sym.Symbol

	PackageFile  map[string]string
	PackageShlib map[string]string

	tramps []*sym.Symbol // trampolines

	// unresolvedSymSet is a set of erroneous unresolved references.
	// Used to avoid duplicated error messages.
	unresolvedSymSet map[unresolvedSymKey]bool
}

type unresolvedSymKey struct {
	from *sym.Symbol // Symbol that referenced unresolved "to"
	to   *sym.Symbol // Unresolved symbol referenced by "from"
}

// ErrorUnresolved prints unresolved symbol error for r.Sym that is referenced from s.
func (ctxt *Link) ErrorUnresolved(psess *PackageSession, s *sym.Symbol, r *sym.Reloc) {
	if ctxt.unresolvedSymSet == nil {
		ctxt.unresolvedSymSet = make(map[unresolvedSymKey]bool)
	}

	k := unresolvedSymKey{from: s, to: r.Sym}
	if !ctxt.unresolvedSymSet[k] {
		ctxt.unresolvedSymSet[k] = true

		if r.Sym.Name == "main.main" {
			psess.
				Errorf(s, "function main is undeclared in the main package")
		} else {
			psess.
				Errorf(s, "relocation target %s not defined", r.Sym.Name)
		}
	}
}

// The smallest possible offset from the hardware stack pointer to a local
// variable on the stack. Architectures that use a link register save its value
// on the stack in the function prologue and so always have a pointer between
// the hardware stack pointer and the local variable area.
func (ctxt *Link) FixedFrameSize() int64 {
	switch ctxt.Arch.Family {
	case sys.AMD64, sys.I386:
		return 0
	case sys.PPC64:

		return int64(4 * ctxt.Arch.PtrSize)
	default:
		return int64(ctxt.Arch.PtrSize)
	}
}

func (ctxt *Link) Logf(format string, args ...interface{}) {
	fmt.Fprintf(ctxt.Bso, format, args...)
	ctxt.Bso.Flush()
}

func (psess *PackageSession) addImports(ctxt *Link, l *sym.Library, pn string) {
	pkg := objabi.PathToPrefix(l.Pkg)
	for _, importStr := range l.ImportStrings {
		lib := psess.addlib(ctxt, pkg, pn, importStr)
		if lib != nil {
			l.Imports = append(l.Imports, lib)
		}
	}
	l.ImportStrings = nil
}

type Pciter struct {
	d       sym.Pcdata
	p       []byte
	pc      uint32
	nextpc  uint32
	pcscale uint32
	value   int32
	start   int
	done    int
}

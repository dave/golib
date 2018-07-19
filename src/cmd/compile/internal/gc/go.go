package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
)

const (
	BADWIDTH        = types.BADWIDTH
	maxStackVarSize = 10 * 1024 * 1024
)

// isRuntimePkg reports whether p is package runtime.
func (psess *PackageSession) isRuntimePkg(p *types.Pkg) bool {
	if psess.compiling_runtime && p == psess.localpkg {
		return true
	}
	return p.Path == "runtime"
}

// The Class of a variable/function describes the "storage class"
// of a variable or function. During parsing, storage classes are
// called declaration contexts.
type Class uint8

//go:generate stringer -type=Class
const (
	Pxxx      Class = iota // no class; used during ssa conversion to indicate pseudo-variables
	PEXTERN                // global variable
	PAUTO                  // local variables
	PAUTOHEAP              // local variable or parameter moved to heap
	PPARAM                 // input arguments
	PPARAMOUT              // output results
	PFUNC                  // global function

	PDISCARD // discard during parse of duplicate import

)

func init() {
	if PDISCARD != 7 {
		panic("PDISCARD changed; does all Class values still fit in three bits?")
	}
}

// note this is the runtime representation
// of the compilers arrays.
//
// typedef	struct
// {				// must not move anything
// 	uchar	array[8];	// pointer to data
// 	uchar	nel[4];		// number of elements
// 	uchar	cap[4];		// allocated number of elements
// } Array;
// runtime offsetof(Array,array) - same for String

// runtime offsetof(Array,nel) - same for String

// runtime offsetof(Array,cap)

// runtime sizeof(Array)

// note this is the runtime representation
// of the compilers strings.
//
// typedef	struct
// {				// must not move anything
// 	uchar	array[8];	// pointer to data
// 	uchar	nel[4];		// number of elements
// } String;
// runtime sizeof(String)

// nerrors is the number of compiler errors reported
// since the last call to saveerrors.

// nsavederrors is the total number of compiler errors
// reported before the last call to saveerrors.

// package being compiled

// set during import

// fake pkg for itab entries

// fake package for runtime itab entries

// fake package runtime

// package runtime/race

// package runtime/msan

// package unsafe

// fake package for field tracking

// fake package for map zero value

// pseudo-package for method symbols on anonymous receiver types

// imported functions and methods with inlinable bodies

// protects funcsyms and associated package lookups (see func funcsym)

// PEXTERN/PAUTO

// Compiling the standard library

// Whether we are adding any sort of code instrumentation, such as
// when the race detector is enabled.

// Whether we are tracking lexical scopes for DWARF.

// Controls generation of DWARF inlined instance records. Zero
// disables, 1 emits inlined routines but suppresses var info,
// and 2 emits inlined routines with tracking of formals/locals.

type Arch struct {
	LinkArch *obj.LinkArch

	REGSP     int
	MAXWIDTH  int64
	Use387    bool // should 386 backend use 387 FP instructions instead of sse2.
	SoftFloat bool

	PadFrame  func(int64) int64
	ZeroRange func(*Progs, *obj.Prog, int64, int64, *uint32) *obj.Prog
	Ginsnop   func(*Progs)

	// SSAMarkMoves marks any MOVXconst ops that need to avoid clobbering flags.
	SSAMarkMoves func(*SSAGenState, *ssa.Block)

	// SSAGenValue emits Prog(s) for the Value.
	SSAGenValue func(*SSAGenState, *ssa.Value)

	// SSAGenBlock emits end-of-block Progs. SSAGenValue should be called
	// for all values in the block before SSAGenBlock.
	SSAGenBlock func(s *SSAGenState, b, next *ssa.Block)

	// ZeroAuto emits code to zero the given auto stack variable.
	// ZeroAuto must not use any non-temporary registers.
	// ZeroAuto will only be called for variables which contain a pointer.
	ZeroAuto func(*Progs, *Node)
}

// GO386=387

// Wasm

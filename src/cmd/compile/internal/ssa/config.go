// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"os"
	"strconv"
)

// A Config holds readonly compilation information.
// It is created once, early during compilation,
// and shared across all compilations.
type Config struct {
	arch            string // "amd64", etc.
	PtrSize         int64  // 4 or 8; copy of cmd/internal/sys.Arch.PtrSize
	RegSize         int64  // 4 or 8; copy of cmd/internal/sys.Arch.RegSize
	Types           Types
	lowerBlock      blockRewriter // lowering function
	lowerValue      valueRewriter // lowering function
	registers       []Register    // machine registers
	gpRegMask       regMask       // general purpose integer register mask
	fpRegMask       regMask       // floating point register mask
	specialRegMask  regMask       // special register mask
	GCRegMap        []*Register   // garbage collector register map, by GC register index
	FPReg           int8          // register number of frame pointer, -1 if not used
	LinkReg         int8          // register number of link register if it is a general purpose register, -1 if not used
	hasGReg         bool          // has hardware g register
	ctxt            *obj.Link     // Generic arch information
	optimize        bool          // Do optimization
	noDuffDevice    bool          // Don't use Duff's device
	useSSE          bool          // Use SSE for non-float operations
	useAvg          bool          // Use optimizations that need Avg* operations
	useHmul         bool          // Use optimizations that need Hmul* operations
	nacl            bool          // GOOS=nacl
	use387          bool          // GO386=387
	SoftFloat       bool          //
	NeedsFpScratch  bool          // No direct move between GP and FP register sets
	BigEndian       bool          //
	sparsePhiCutoff uint64        // Sparse phi location algorithm used above this #blocks*#variables score
}

type (
	blockRewriter func(*Block) bool
	valueRewriter func(*Value) bool
)

type Types struct {
	Bool       *types.Type
	Int8       *types.Type
	Int16      *types.Type
	Int32      *types.Type
	Int64      *types.Type
	UInt8      *types.Type
	UInt16     *types.Type
	UInt32     *types.Type
	UInt64     *types.Type
	Int        *types.Type
	Float32    *types.Type
	Float64    *types.Type
	UInt       *types.Type
	Uintptr    *types.Type
	String     *types.Type
	BytePtr    *types.Type // TODO: use unsafe.Pointer instead?
	Int32Ptr   *types.Type
	UInt32Ptr  *types.Type
	IntPtr     *types.Type
	UintptrPtr *types.Type
	Float32Ptr *types.Type
	Float64Ptr *types.Type
	BytePtrPtr *types.Type
}

// NewTypes creates and populates a Types.
func (pstate *PackageState) NewTypes() *Types {
	t := new(Types)
	t.SetTypPtrs(pstate)
	return t
}

// SetTypPtrs populates t.
func (t *Types) SetTypPtrs(pstate *PackageState) {
	t.Bool = pstate.types.Types[types.TBOOL]
	t.Int8 = pstate.types.Types[types.TINT8]
	t.Int16 = pstate.types.Types[types.TINT16]
	t.Int32 = pstate.types.Types[types.TINT32]
	t.Int64 = pstate.types.Types[types.TINT64]
	t.UInt8 = pstate.types.Types[types.TUINT8]
	t.UInt16 = pstate.types.Types[types.TUINT16]
	t.UInt32 = pstate.types.Types[types.TUINT32]
	t.UInt64 = pstate.types.Types[types.TUINT64]
	t.Int = pstate.types.Types[types.TINT]
	t.Float32 = pstate.types.Types[types.TFLOAT32]
	t.Float64 = pstate.types.Types[types.TFLOAT64]
	t.UInt = pstate.types.Types[types.TUINT]
	t.Uintptr = pstate.types.Types[types.TUINTPTR]
	t.String = pstate.types.Types[types.TSTRING]
	t.BytePtr = pstate.types.NewPtr(pstate.types.Types[types.TUINT8])
	t.Int32Ptr = pstate.types.NewPtr(pstate.types.Types[types.TINT32])
	t.UInt32Ptr = pstate.types.NewPtr(pstate.types.Types[types.TUINT32])
	t.IntPtr = pstate.types.NewPtr(pstate.types.Types[types.TINT])
	t.UintptrPtr = pstate.types.NewPtr(pstate.types.Types[types.TUINTPTR])
	t.Float32Ptr = pstate.types.NewPtr(pstate.types.Types[types.TFLOAT32])
	t.Float64Ptr = pstate.types.NewPtr(pstate.types.Types[types.TFLOAT64])
	t.BytePtrPtr = pstate.types.NewPtr(pstate.types.NewPtr(pstate.types.Types[types.TUINT8]))
}

type Logger interface {
	// Logf logs a message from the compiler.
	Logf(string, ...interface{})

	// Log returns true if logging is not a no-op
	// some logging calls account for more than a few heap allocations.
	Log() bool

	// Fatal reports a compiler error and exits.
	Fatalf(pos src.XPos, msg string, args ...interface{})

	// Warnl writes compiler messages in the form expected by "errorcheck" tests
	Warnl(pos src.XPos, fmt_ string, args ...interface{})

	// Forwards the Debug flags from gc
	Debug_checknil() bool
}

type Frontend interface {
	CanSSA(t *types.Type) bool

	Logger

	// StringData returns a symbol pointing to the given string's contents.
	StringData(string) interface{} // returns *gc.Sym

	// Auto returns a Node for an auto variable of the given type.
	// The SSA compiler uses this function to allocate space for spills.
	Auto(src.XPos, *types.Type) GCNode

	// Given the name for a compound type, returns the name we should use
	// for the parts of that compound type.
	SplitString(LocalSlot) (LocalSlot, LocalSlot)
	SplitInterface(LocalSlot) (LocalSlot, LocalSlot)
	SplitSlice(LocalSlot) (LocalSlot, LocalSlot, LocalSlot)
	SplitComplex(LocalSlot) (LocalSlot, LocalSlot)
	SplitStruct(LocalSlot, int) LocalSlot
	SplitArray(LocalSlot) LocalSlot              // array must be length 1
	SplitInt64(LocalSlot) (LocalSlot, LocalSlot) // returns (hi, lo)

	// DerefItab dereferences an itab function
	// entry, given the symbol of the itab and
	// the byte offset of the function pointer.
	// It may return nil.
	DerefItab(sym *obj.LSym, offset int64) *obj.LSym

	// Line returns a string describing the given position.
	Line(src.XPos) string

	// AllocFrame assigns frame offsets to all live auto variables.
	AllocFrame(f *Func)

	// Syslook returns a symbol of the runtime function/variable with the
	// given name.
	Syslook(string) *obj.LSym

	// UseWriteBarrier returns whether write barrier is enabled
	UseWriteBarrier() bool

	// SetWBPos indicates that a write barrier has been inserted
	// in this function at position pos.
	SetWBPos(pos src.XPos)
}

// interface used to hold a *gc.Node (a stack variable).
// We'd use *gc.Node directly but that would lead to an import cycle.
type GCNode interface {
	Typ() *types.Type
	String() string
	IsSynthetic() bool
	StorageClass() StorageClass
}

type StorageClass uint8

const (
	ClassAuto     StorageClass = iota // local stack variable
	ClassParam                        // argument
	ClassParamOut                     // return value
)

// NewConfig returns a new configuration object for the given architecture.
func (pstate *PackageState) NewConfig(arch string, types Types, ctxt *obj.Link, optimize bool) *Config {
	c := &Config{arch: arch, Types: types}
	c.useAvg = true
	c.useHmul = true
	switch arch {
	case "amd64":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = pstate.rewriteBlockAMD64
		c.lowerValue = pstate.rewriteValueAMD64
		c.registers = pstate.registersAMD64[:]
		c.gpRegMask = pstate.gpRegMaskAMD64
		c.fpRegMask = pstate.fpRegMaskAMD64
		c.FPReg = pstate.framepointerRegAMD64
		c.LinkReg = pstate.linkRegAMD64
		c.hasGReg = false
	case "amd64p32":
		c.PtrSize = 4
		c.RegSize = 8
		c.lowerBlock = pstate.rewriteBlockAMD64
		c.lowerValue = pstate.rewriteValueAMD64
		c.registers = pstate.registersAMD64[:]
		c.gpRegMask = pstate.gpRegMaskAMD64
		c.fpRegMask = pstate.fpRegMaskAMD64
		c.FPReg = pstate.framepointerRegAMD64
		c.LinkReg = pstate.linkRegAMD64
		c.hasGReg = false
		c.noDuffDevice = true
	case "386":
		c.PtrSize = 4
		c.RegSize = 4
		c.lowerBlock = pstate.rewriteBlock386
		c.lowerValue = pstate.rewriteValue386
		c.registers = pstate.registers386[:]
		c.gpRegMask = pstate.gpRegMask386
		c.fpRegMask = pstate.fpRegMask386
		c.FPReg = pstate.framepointerReg386
		c.LinkReg = pstate.linkReg386
		c.hasGReg = false
	case "arm":
		c.PtrSize = 4
		c.RegSize = 4
		c.lowerBlock = pstate.rewriteBlockARM
		c.lowerValue = pstate.rewriteValueARM
		c.registers = pstate.registersARM[:]
		c.gpRegMask = pstate.gpRegMaskARM
		c.fpRegMask = pstate.fpRegMaskARM
		c.FPReg = pstate.framepointerRegARM
		c.LinkReg = pstate.linkRegARM
		c.hasGReg = true
	case "arm64":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = pstate.rewriteBlockARM64
		c.lowerValue = pstate.rewriteValueARM64
		c.registers = pstate.registersARM64[:]
		c.gpRegMask = pstate.gpRegMaskARM64
		c.fpRegMask = pstate.fpRegMaskARM64
		c.FPReg = pstate.framepointerRegARM64
		c.LinkReg = pstate.linkRegARM64
		c.hasGReg = true
		c.noDuffDevice = pstate.objabi.GOOS == "darwin" // darwin linker cannot handle BR26 reloc with non-zero addend
	case "ppc64":
		c.BigEndian = true
		fallthrough
	case "ppc64le":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = pstate.rewriteBlockPPC64
		c.lowerValue = pstate.rewriteValuePPC64
		c.registers = pstate.registersPPC64[:]
		c.gpRegMask = pstate.gpRegMaskPPC64
		c.fpRegMask = pstate.fpRegMaskPPC64
		c.FPReg = pstate.framepointerRegPPC64
		c.LinkReg = pstate.linkRegPPC64
		c.noDuffDevice = true // TODO: Resolve PPC64 DuffDevice (has zero, but not copy)
		c.hasGReg = true
	case "mips64":
		c.BigEndian = true
		fallthrough
	case "mips64le":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockMIPS64
		c.lowerValue = pstate.rewriteValueMIPS64
		c.registers = pstate.registersMIPS64[:]
		c.gpRegMask = pstate.gpRegMaskMIPS64
		c.fpRegMask = pstate.fpRegMaskMIPS64
		c.specialRegMask = pstate.specialRegMaskMIPS64
		c.FPReg = pstate.framepointerRegMIPS64
		c.LinkReg = pstate.linkRegMIPS64
		c.hasGReg = true
	case "s390x":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = pstate.rewriteBlockS390X
		c.lowerValue = pstate.rewriteValueS390X
		c.registers = pstate.registersS390X[:]
		c.gpRegMask = pstate.gpRegMaskS390X
		c.fpRegMask = pstate.fpRegMaskS390X
		c.FPReg = pstate.framepointerRegS390X
		c.LinkReg = pstate.linkRegS390X
		c.hasGReg = true
		c.noDuffDevice = true
		c.BigEndian = true
	case "mips":
		c.BigEndian = true
		fallthrough
	case "mipsle":
		c.PtrSize = 4
		c.RegSize = 4
		c.lowerBlock = rewriteBlockMIPS
		c.lowerValue = pstate.rewriteValueMIPS
		c.registers = pstate.registersMIPS[:]
		c.gpRegMask = pstate.gpRegMaskMIPS
		c.fpRegMask = pstate.fpRegMaskMIPS
		c.specialRegMask = pstate.specialRegMaskMIPS
		c.FPReg = pstate.framepointerRegMIPS
		c.LinkReg = pstate.linkRegMIPS
		c.hasGReg = true
		c.noDuffDevice = true
	case "wasm":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockWasm
		c.lowerValue = pstate.rewriteValueWasm
		c.registers = pstate.registersWasm[:]
		c.gpRegMask = pstate.gpRegMaskWasm
		c.fpRegMask = pstate.fpRegMaskWasm
		c.FPReg = pstate.framepointerRegWasm
		c.LinkReg = pstate.linkRegWasm
		c.hasGReg = true
		c.noDuffDevice = true
		c.useAvg = false
		c.useHmul = false
	default:
		ctxt.Diag("arch %s not implemented", arch)
	}
	c.ctxt = ctxt
	c.optimize = optimize
	c.nacl = pstate.objabi.GOOS == "nacl"
	c.useSSE = true

	// Don't use Duff's device nor SSE on Plan 9 AMD64, because
	// floating point operations are not allowed in note handler.
	if pstate.objabi.GOOS == "plan9" && arch == "amd64" {
		c.noDuffDevice = true
		c.useSSE = false
	}

	if c.nacl {
		c.noDuffDevice = true // Don't use Duff's device on NaCl

		// Returns clobber BP on nacl/386, so the write
		// barrier does.
		pstate.opcodeTable[Op386LoweredWB].reg.clobbers |= 1 << 5 // BP

		// ... and SI on nacl/amd64.
		pstate.opcodeTable[OpAMD64LoweredWB].reg.clobbers |= 1 << 6 // SI
	}

	if ctxt.Flag_shared {
		// LoweredWB is secretly a CALL and CALLs on 386 in
		// shared mode get rewritten by obj6.go to go through
		// the GOT, which clobbers BX.
		pstate.opcodeTable[Op386LoweredWB].reg.clobbers |= 1 << 3 // BX
	}

	// cutoff is compared with product of numblocks and numvalues,
	// if product is smaller than cutoff, use old non-sparse method.
	// cutoff == 0 implies all sparse.
	// cutoff == -1 implies none sparse.
	// Good cutoff values seem to be O(million) depending on constant factor cost of sparse.
	// TODO: get this from a flag, not an environment variable
	c.sparsePhiCutoff = 2500000 // 0 for testing. // 2500000 determined with crude experiments w/ make.bash
	ev := os.Getenv("GO_SSA_PHI_LOC_CUTOFF")
	if ev != "" {
		v, err := strconv.ParseInt(ev, 10, 64)
		if err != nil {
			ctxt.Diag("Environment variable GO_SSA_PHI_LOC_CUTOFF (value '%s') did not parse as a number", ev)
		}
		c.sparsePhiCutoff = uint64(v) // convert -1 to maxint, for never use sparse
	}

	// Create the GC register map index.
	// TODO: This is only used for debug printing. Maybe export config.registers?
	gcRegMapSize := int16(0)
	for _, r := range c.registers {
		if r.gcNum+1 > gcRegMapSize {
			gcRegMapSize = r.gcNum + 1
		}
	}
	c.GCRegMap = make([]*Register, gcRegMapSize)
	for i, r := range c.registers {
		if r.gcNum != -1 {
			c.GCRegMap[r.gcNum] = &c.registers[i]
		}
	}

	return c
}

func (c *Config) Set387(b bool) {
	c.NeedsFpScratch = b
	c.use387 = b
}

func (c *Config) SparsePhiCutoff() uint64 { return c.sparsePhiCutoff }
func (c *Config) Ctxt() *obj.Link         { return c.ctxt }

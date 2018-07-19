package ssa

import (
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"

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
func (psess *PackageSession) NewTypes() *Types {
	t := new(Types)
	t.SetTypPtrs(psess)
	return t
}

// SetTypPtrs populates t.
func (t *Types) SetTypPtrs(psess *PackageSession) {
	t.Bool = psess.types.Types[types.TBOOL]
	t.Int8 = psess.types.Types[types.TINT8]
	t.Int16 = psess.types.Types[types.TINT16]
	t.Int32 = psess.types.Types[types.TINT32]
	t.Int64 = psess.types.Types[types.TINT64]
	t.UInt8 = psess.types.Types[types.TUINT8]
	t.UInt16 = psess.types.Types[types.TUINT16]
	t.UInt32 = psess.types.Types[types.TUINT32]
	t.UInt64 = psess.types.Types[types.TUINT64]
	t.Int = psess.types.Types[types.TINT]
	t.Float32 = psess.types.Types[types.TFLOAT32]
	t.Float64 = psess.types.Types[types.TFLOAT64]
	t.UInt = psess.types.Types[types.TUINT]
	t.Uintptr = psess.types.Types[types.TUINTPTR]
	t.String = psess.types.Types[types.TSTRING]
	t.BytePtr = psess.types.NewPtr(psess.types.Types[types.TUINT8])
	t.Int32Ptr = psess.types.NewPtr(psess.types.Types[types.TINT32])
	t.UInt32Ptr = psess.types.NewPtr(psess.types.Types[types.TUINT32])
	t.IntPtr = psess.types.NewPtr(psess.types.Types[types.TINT])
	t.UintptrPtr = psess.types.NewPtr(psess.types.Types[types.TUINTPTR])
	t.Float32Ptr = psess.types.NewPtr(psess.types.Types[types.TFLOAT32])
	t.Float64Ptr = psess.types.NewPtr(psess.types.Types[types.TFLOAT64])
	t.BytePtrPtr = psess.types.NewPtr(psess.types.NewPtr(psess.types.Types[types.TUINT8]))
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
func (psess *PackageSession) NewConfig(arch string, types Types, ctxt *obj.Link, optimize bool) *Config {
	c := &Config{arch: arch, Types: types}
	c.useAvg = true
	c.useHmul = true
	switch arch {
	case "amd64":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = psess.rewriteBlockAMD64
		c.lowerValue = psess.rewriteValueAMD64
		c.registers = psess.registersAMD64[:]
		c.gpRegMask = psess.gpRegMaskAMD64
		c.fpRegMask = psess.fpRegMaskAMD64
		c.FPReg = psess.framepointerRegAMD64
		c.LinkReg = psess.linkRegAMD64
		c.hasGReg = false
	case "amd64p32":
		c.PtrSize = 4
		c.RegSize = 8
		c.lowerBlock = psess.rewriteBlockAMD64
		c.lowerValue = psess.rewriteValueAMD64
		c.registers = psess.registersAMD64[:]
		c.gpRegMask = psess.gpRegMaskAMD64
		c.fpRegMask = psess.fpRegMaskAMD64
		c.FPReg = psess.framepointerRegAMD64
		c.LinkReg = psess.linkRegAMD64
		c.hasGReg = false
		c.noDuffDevice = true
	case "386":
		c.PtrSize = 4
		c.RegSize = 4
		c.lowerBlock = psess.rewriteBlock386
		c.lowerValue = psess.rewriteValue386
		c.registers = psess.registers386[:]
		c.gpRegMask = psess.gpRegMask386
		c.fpRegMask = psess.fpRegMask386
		c.FPReg = psess.framepointerReg386
		c.LinkReg = psess.linkReg386
		c.hasGReg = false
	case "arm":
		c.PtrSize = 4
		c.RegSize = 4
		c.lowerBlock = psess.rewriteBlockARM
		c.lowerValue = psess.rewriteValueARM
		c.registers = psess.registersARM[:]
		c.gpRegMask = psess.gpRegMaskARM
		c.fpRegMask = psess.fpRegMaskARM
		c.FPReg = psess.framepointerRegARM
		c.LinkReg = psess.linkRegARM
		c.hasGReg = true
	case "arm64":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = psess.rewriteBlockARM64
		c.lowerValue = psess.rewriteValueARM64
		c.registers = psess.registersARM64[:]
		c.gpRegMask = psess.gpRegMaskARM64
		c.fpRegMask = psess.fpRegMaskARM64
		c.FPReg = psess.framepointerRegARM64
		c.LinkReg = psess.linkRegARM64
		c.hasGReg = true
		c.noDuffDevice = psess.objabi.GOOS == "darwin"
	case "ppc64":
		c.BigEndian = true
		fallthrough
	case "ppc64le":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = psess.rewriteBlockPPC64
		c.lowerValue = psess.rewriteValuePPC64
		c.registers = psess.registersPPC64[:]
		c.gpRegMask = psess.gpRegMaskPPC64
		c.fpRegMask = psess.fpRegMaskPPC64
		c.FPReg = psess.framepointerRegPPC64
		c.LinkReg = psess.linkRegPPC64
		c.noDuffDevice = true
		c.hasGReg = true
	case "mips64":
		c.BigEndian = true
		fallthrough
	case "mips64le":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockMIPS64
		c.lowerValue = psess.rewriteValueMIPS64
		c.registers = psess.registersMIPS64[:]
		c.gpRegMask = psess.gpRegMaskMIPS64
		c.fpRegMask = psess.fpRegMaskMIPS64
		c.specialRegMask = psess.specialRegMaskMIPS64
		c.FPReg = psess.framepointerRegMIPS64
		c.LinkReg = psess.linkRegMIPS64
		c.hasGReg = true
	case "s390x":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = psess.rewriteBlockS390X
		c.lowerValue = psess.rewriteValueS390X
		c.registers = psess.registersS390X[:]
		c.gpRegMask = psess.gpRegMaskS390X
		c.fpRegMask = psess.fpRegMaskS390X
		c.FPReg = psess.framepointerRegS390X
		c.LinkReg = psess.linkRegS390X
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
		c.lowerValue = psess.rewriteValueMIPS
		c.registers = psess.registersMIPS[:]
		c.gpRegMask = psess.gpRegMaskMIPS
		c.fpRegMask = psess.fpRegMaskMIPS
		c.specialRegMask = psess.specialRegMaskMIPS
		c.FPReg = psess.framepointerRegMIPS
		c.LinkReg = psess.linkRegMIPS
		c.hasGReg = true
		c.noDuffDevice = true
	case "wasm":
		c.PtrSize = 8
		c.RegSize = 8
		c.lowerBlock = rewriteBlockWasm
		c.lowerValue = psess.rewriteValueWasm
		c.registers = psess.registersWasm[:]
		c.gpRegMask = psess.gpRegMaskWasm
		c.fpRegMask = psess.fpRegMaskWasm
		c.FPReg = psess.framepointerRegWasm
		c.LinkReg = psess.linkRegWasm
		c.hasGReg = true
		c.noDuffDevice = true
		c.useAvg = false
		c.useHmul = false
	default:
		ctxt.Diag("arch %s not implemented", arch)
	}
	c.ctxt = ctxt
	c.optimize = optimize
	c.nacl = psess.objabi.GOOS == "nacl"
	c.useSSE = true

	if psess.objabi.GOOS == "plan9" && arch == "amd64" {
		c.noDuffDevice = true
		c.useSSE = false
	}

	if c.nacl {
		c.noDuffDevice = true
		psess.
			opcodeTable[Op386LoweredWB].reg.clobbers |= 1 << 5
		psess.
			opcodeTable[OpAMD64LoweredWB].reg.clobbers |= 1 << 6
	}

	if ctxt.Flag_shared {
		psess.
			opcodeTable[Op386LoweredWB].reg.clobbers |= 1 << 3
	}

	c.sparsePhiCutoff = 2500000
	ev := os.Getenv("GO_SSA_PHI_LOC_CUTOFF")
	if ev != "" {
		v, err := strconv.ParseInt(ev, 10, 64)
		if err != nil {
			ctxt.Diag("Environment variable GO_SSA_PHI_LOC_CUTOFF (value '%s') did not parse as a number", ev)
		}
		c.sparsePhiCutoff = uint64(v)
	}

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

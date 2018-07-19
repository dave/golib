package obj

import (
	"bufio"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/dwarf"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"github.com/dave/golib/src/cmd/internal/sys"
	"sync"
)

type Addr struct {
	Reg    int16
	Index  int16
	Scale  int16 // Sometimes holds a register.
	Type   AddrType
	Name   AddrName
	Class  int8
	Offset int64
	Sym    *LSym

	// argument value:
	//	for TYPE_SCONST, a string
	//	for TYPE_FCONST, a float64
	//	for TYPE_BRANCH, a *Prog (optional)
	//	for TYPE_TEXTSIZE, an int32 (optional)
	Val interface{}
}

type AddrName int8

const (
	NAME_NONE AddrName = iota
	NAME_EXTERN
	NAME_STATIC
	NAME_AUTO
	NAME_PARAM
	// A reference to name@GOT(SB) is a reference to the entry in the global offset
	// table for 'name'.
	NAME_GOTREF
	// Indicates auto that was optimized away, but whose type
	// we want to preserve in the DWARF debug info.
	NAME_DELETED_AUTO
)

type AddrType uint8

const (
	TYPE_NONE AddrType = iota
	TYPE_BRANCH
	TYPE_TEXTSIZE
	TYPE_MEM
	TYPE_CONST
	TYPE_FCONST
	TYPE_SCONST
	TYPE_REG
	TYPE_ADDR
	TYPE_SHIFT
	TYPE_REGREG
	TYPE_REGREG2
	TYPE_INDIR
	TYPE_REGLIST
)

// Prog describes a single machine instruction.
//
// The general instruction form is:
//
//	(1) As.Scond From [, ...RestArgs], To
//	(2) As.Scond From, Reg [, ...RestArgs], To, RegTo2
//
// where As is an opcode and the others are arguments:
// From, Reg are sources, and To, RegTo2 are destinations.
// RestArgs can hold additional sources and destinations.
// Usually, not all arguments are present.
// For example, MOVL R1, R2 encodes using only As=MOVL, From=R1, To=R2.
// The Scond field holds additional condition bits for systems (like arm)
// that have generalized conditional execution.
// (2) form is present for compatibility with older code,
// to avoid too much changes in a single swing.
// (1) scheme is enough to express any kind of operand combination.
//
// Jump instructions use the Pcond field to point to the target instruction,
// which must be in the same linked list as the jump instruction.
//
// The Progs for a given function are arranged in a list linked through the Link field.
//
// Each Prog is charged to a specific source line in the debug information,
// specified by Pos.Line().
// Every Prog has a Ctxt field that defines its context.
// For performance reasons, Progs usually are usually bulk allocated, cached, and reused;
// those bulk allocators should always be used, rather than new(Prog).
//
// The other fields not yet mentioned are for use by the back ends and should
// be left zeroed by creators of Prog lists.
type Prog struct {
	Ctxt     *Link    // linker context
	Link     *Prog    // next Prog in linked list
	From     Addr     // first source operand
	RestArgs []Addr   // can pack any operands that not fit into {Prog.From, Prog.To}
	To       Addr     // destination operand (second is RegTo2 below)
	Pcond    *Prog    // target of conditional jump
	Forwd    *Prog    // for x86 back end
	Rel      *Prog    // for x86, arm back ends
	Pc       int64    // for back ends or assembler: virtual or actual program counter, depending on phase
	Pos      src.XPos // source position of this instruction
	Spadj    int32    // effect of instruction on stack pointer (increment or decrement amount)
	As       As       // assembler opcode
	Reg      int16    // 2nd source operand
	RegTo2   int16    // 2nd destination operand
	Mark     uint16   // bitmask of arch-specific items
	Optab    uint16   // arch-specific opcode index
	Scond    uint8    // bits that describe instruction suffixes (e.g. ARM conditions)
	Back     uint8    // for x86 back end: backwards branch state
	Ft       uint8    // for x86 back end: type index of Prog.From
	Tt       uint8    // for x86 back end: type index of Prog.To
	Isize    uint8    // for x86 back end: size of the instruction in bytes
}

// From3Type returns p.GetFrom3().Type, or TYPE_NONE when
// p.GetFrom3() returns nil.
//
// Deprecated: for the same reasons as Prog.GetFrom3.
func (p *Prog) From3Type() AddrType {
	if p.RestArgs == nil {
		return TYPE_NONE
	}
	return p.RestArgs[0].Type
}

// GetFrom3 returns second source operand (the first is Prog.From).
// In combination with Prog.From and Prog.To it makes common 3 operand
// case easier to use.
//
// Should be used only when RestArgs is set with SetFrom3.
//
// Deprecated: better use RestArgs directly or define backend-specific getters.
// Introduced to simplify transition to []Addr.
// Usage of this is discouraged due to fragility and lack of guarantees.
func (p *Prog) GetFrom3() *Addr {
	if p.RestArgs == nil {
		return nil
	}
	return &p.RestArgs[0]
}

// SetFrom3 assigns []Addr{a} to p.RestArgs.
// In pair with Prog.GetFrom3 it can help in emulation of Prog.From3.
//
// Deprecated: for the same reasons as Prog.GetFrom3.
func (p *Prog) SetFrom3(a Addr) {
	p.RestArgs = []Addr{a}
}

// An As denotes an assembler opcode.
// There are some portable opcodes, declared here in package obj,
// that are common to all architectures.
// However, the majority of opcodes are arch-specific
// and are declared in their respective architecture's subpackage.
type As int16

// These are the portable opcodes.
const (
	AXXX As = iota
	ACALL
	ADUFFCOPY
	ADUFFZERO
	AEND
	AFUNCDATA
	AJMP
	ANOP
	APCDATA
	ARET
	AGETCALLERPC
	ATEXT
	AUNDEF
	A_ARCHSPECIFIC
)

// Each architecture is allotted a distinct subspace of opcode values
// for declaring its arch-specific opcodes.
// Within this subspace, the first arch-specific opcode should be
// at offset A_ARCHSPECIFIC.
//
// Subspaces are aligned to a power of two so opcodes can be masked
// with AMask and used as compact array indices.
const (
	ABase386 = (1 + iota) << 11
	ABaseARM
	ABaseAMD64
	ABasePPC64
	ABaseARM64
	ABaseMIPS
	ABaseS390X
	ABaseWasm

	AllowedOpCodes = 1 << 11            // The number of opcodes available for any given architecture.
	AMask          = AllowedOpCodes - 1 // AND with this to use the opcode as an array index.
)

// An LSym is the sort of symbol that is written to an object file.
type LSym struct {
	Name string
	Type objabi.SymKind
	Attribute

	RefIdx int // Index of this symbol in the symbol reference list.
	Size   int64
	Gotype *LSym
	P      []byte
	R      []Reloc

	Func *FuncInfo
}

// A FuncInfo contains extra fields for STEXT symbols.
type FuncInfo struct {
	Args   int32
	Locals int32
	Text   *Prog
	Autom  []*Auto
	Pcln   Pcln

	dwarfInfoSym   *LSym
	dwarfLocSym    *LSym
	dwarfRangesSym *LSym
	dwarfAbsFnSym  *LSym
	dwarfIsStmtSym *LSym

	GCArgs   LSym
	GCLocals LSym
	GCRegs   LSym
}

// Attribute is a set of symbol attributes.
type Attribute int16

const (
	AttrDuplicateOK Attribute = 1 << iota
	AttrCFunc
	AttrNoSplit
	AttrLeaf
	AttrWrapper
	AttrNeedCtxt
	AttrNoFrame
	AttrSeenGlobl
	AttrOnList
	AttrStatic

	// MakeTypelink means that the type should have an entry in the typelink table.
	AttrMakeTypelink

	// ReflectMethod means the function may call reflect.Type.Method or
	// reflect.Type.MethodByName. Matching is imprecise (as reflect.Type
	// can be used through a custom interface), so ReflectMethod may be
	// set in some cases when the reflect package is not called.
	//
	// Used by the linker to determine what methods can be pruned.
	AttrReflectMethod

	// Local means make the symbol local even when compiling Go code to reference Go
	// symbols in other shared libraries, as in this mode symbols are global by
	// default. "local" here means in the sense of the dynamic linker, i.e. not
	// visible outside of the module (shared library or executable) that contains its
	// definition. (When not compiling to support Go shared libraries, all symbols are
	// local in this sense unless there is a cgo_export_* directive).
	AttrLocal

	// For function symbols; indicates that the specified function was the
	// target of an inline during compilation
	AttrWasInlined
)

func (a Attribute) DuplicateOK() bool   { return a&AttrDuplicateOK != 0 }
func (a Attribute) MakeTypelink() bool  { return a&AttrMakeTypelink != 0 }
func (a Attribute) CFunc() bool         { return a&AttrCFunc != 0 }
func (a Attribute) NoSplit() bool       { return a&AttrNoSplit != 0 }
func (a Attribute) Leaf() bool          { return a&AttrLeaf != 0 }
func (a Attribute) SeenGlobl() bool     { return a&AttrSeenGlobl != 0 }
func (a Attribute) OnList() bool        { return a&AttrOnList != 0 }
func (a Attribute) ReflectMethod() bool { return a&AttrReflectMethod != 0 }
func (a Attribute) Local() bool         { return a&AttrLocal != 0 }
func (a Attribute) Wrapper() bool       { return a&AttrWrapper != 0 }
func (a Attribute) NeedCtxt() bool      { return a&AttrNeedCtxt != 0 }
func (a Attribute) NoFrame() bool       { return a&AttrNoFrame != 0 }
func (a Attribute) Static() bool        { return a&AttrStatic != 0 }
func (a Attribute) WasInlined() bool    { return a&AttrWasInlined != 0 }

func (a *Attribute) Set(flag Attribute, value bool) {
	if value {
		*a |= flag
	} else {
		*a &^= flag
	}
}

// TextAttrString formats a for printing in as part of a TEXT prog.
func (a Attribute) TextAttrString(psess *PackageSession) string {
	var s string
	for _, x := range psess.textAttrStrings {
		if a&x.bit != 0 {
			if x.s != "" {
				s += x.s + "|"
			}
			a &^= x.bit
		}
	}
	if a != 0 {
		s += fmt.Sprintf("UnknownAttribute(%d)|", a)
	}

	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

// The compiler needs LSym to satisfy fmt.Stringer, because it stores
// an LSym in ssa.ExternSymbol.
func (s *LSym) String() string {
	return s.Name
}

type Pcln struct {
	Pcsp        Pcdata
	Pcfile      Pcdata
	Pcline      Pcdata
	Pcinline    Pcdata
	Pcdata      []Pcdata
	Funcdata    []*LSym
	Funcdataoff []int64
	File        []string
	Lastfile    string
	Lastindex   int
	InlTree     InlTree // per-function inlining tree extracted from the global tree
}

type Reloc struct {
	Off  int32
	Siz  uint8
	Type objabi.RelocType
	Add  int64
	Sym  *LSym
}

type Auto struct {
	Asym    *LSym
	Aoffset int32
	Name    AddrName
	Gotype  *LSym
}

type Pcdata struct {
	P []byte
}

// Link holds the context for writing object code from a compiler
// to be linker input or for reading that input into the linker.
type Link struct {
	Headtype           objabi.HeadType
	Arch               *LinkArch
	Debugasm           bool
	Debugvlog          bool
	Debugpcln          string
	Flag_shared        bool
	Flag_dynlink       bool
	Flag_optimize      bool
	Flag_locationlists bool
	Bso                *bufio.Writer
	Pathname           string
	hashmu             sync.Mutex       // protects hash
	hash               map[string]*LSym // name -> sym mapping
	statichash         map[string]*LSym // name -> sym mapping for static syms
	PosTable           src.PosTable
	InlTree            InlTree // global inlining tree used by gc/inl.go
	DwFixups           *DwarfFixupTable
	Imports            []string
	DiagFunc           func(string, ...interface{})
	DiagFlush          func()
	DebugInfo          func(fn *LSym, curfn interface{}) ([]dwarf.Scope, dwarf.InlCalls) // if non-nil, curfn is a *gc.Node
	GenAbstractFunc    func(fn *LSym)
	Errors             int

	InParallel           bool // parallel backend phase in effect
	Framepointer_enabled bool

	// state for writing objects
	Text []*LSym
	Data []*LSym
}

func (ctxt *Link) Diag(format string, args ...interface{}) {
	ctxt.Errors++
	ctxt.DiagFunc(format, args...)
}

func (ctxt *Link) Logf(format string, args ...interface{}) {
	fmt.Fprintf(ctxt.Bso, format, args...)
	ctxt.Bso.Flush()
}

// The smallest possible offset from the hardware stack pointer to a local
// variable on the stack. Architectures that use a link register save its value
// on the stack in the function prologue and so always have a pointer between
// the hardware stack pointer and the local variable area.
func (ctxt *Link) FixedFrameSize() int64 {
	switch ctxt.Arch.Family {
	case sys.AMD64, sys.I386, sys.Wasm:
		return 0
	case sys.PPC64:

		return int64(4 * ctxt.Arch.PtrSize)
	default:
		return int64(ctxt.Arch.PtrSize)
	}
}

// LinkArch is the definition of a single architecture.
type LinkArch struct {
	*sys.Arch
	Init           func(*Link)
	Preprocess     func(*Link, *LSym, ProgAlloc)
	Assemble       func(*Link, *LSym, ProgAlloc)
	Progedit       func(*Link, *Prog, ProgAlloc)
	UnaryDst       map[As]bool // Instruction takes one operand, a destination.
	DWARFRegisters map[int16]int16
}

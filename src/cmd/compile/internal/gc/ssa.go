package gc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"html"
	"os"
	"sort"

	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/src"
	"github.com/dave/golib/src/cmd/internal/sys"
)

func (psess *PackageSession) initssaconfig() {
	types_ := psess.ssa.NewTypes()

	if psess.thearch.SoftFloat {
		psess.
			softfloatInit()
	}

	_ = psess.types.NewPtr(psess.types.Types[TINTER])
	_ = psess.types.NewPtr(psess.types.NewPtr(psess.types.Types[TSTRING]))
	_ = psess.types.NewPtr(psess.types.NewPtr(psess.types.Idealstring))
	_ = psess.types.NewPtr(psess.types.NewSlice(psess.types.Types[TINTER]))
	_ = psess.types.NewPtr(psess.types.NewPtr(psess.types.Bytetype))
	_ = psess.types.NewPtr(psess.types.NewSlice(psess.types.Bytetype))
	_ = psess.types.NewPtr(psess.types.NewSlice(psess.types.Types[TSTRING]))
	_ = psess.types.NewPtr(psess.types.NewSlice(psess.types.Idealstring))
	_ = psess.types.NewPtr(psess.types.NewPtr(psess.types.NewPtr(psess.types.Types[TUINT8])))
	_ = psess.types.NewPtr(psess.types.Types[TINT16])
	_ = psess.types.NewPtr(psess.types.Types[TINT64])
	_ = psess.types.NewPtr(psess.types.Errortype)
	psess.types.
		NewPtrCacheEnabled = false
	psess.
		ssaConfig = psess.ssa.NewConfig(psess.thearch.LinkArch.Name, *types_, psess.Ctxt, psess.Debug['N'] == 0)
	if psess.thearch.LinkArch.Name == "386" {
		psess.
			ssaConfig.Set387(psess.thearch.Use387)
	}
	psess.
		ssaConfig.SoftFloat = psess.thearch.SoftFloat
	psess.
		ssaCaches = make([]ssa.Cache, psess.nBackendWorkers)
	psess.
		assertE2I = psess.sysfunc("assertE2I")
	psess.
		assertE2I2 = psess.sysfunc("assertE2I2")
	psess.
		assertI2I = psess.sysfunc("assertI2I")
	psess.
		assertI2I2 = psess.sysfunc("assertI2I2")
	psess.
		Deferproc = psess.sysfunc("deferproc")
	psess.
		Deferreturn = psess.sysfunc("deferreturn")
	psess.
		Duffcopy = psess.sysfunc("duffcopy")
	psess.
		Duffzero = psess.sysfunc("duffzero")
	psess.
		gcWriteBarrier = psess.sysfunc("gcWriteBarrier")
	psess.
		goschedguarded = psess.sysfunc("goschedguarded")
	psess.
		growslice = psess.sysfunc("growslice")
	psess.
		msanread = psess.sysfunc("msanread")
	psess.
		msanwrite = psess.sysfunc("msanwrite")
	psess.
		Newproc = psess.sysfunc("newproc")
	psess.
		panicdivide = psess.sysfunc("panicdivide")
	psess.
		panicdottypeE = psess.sysfunc("panicdottypeE")
	psess.
		panicdottypeI = psess.sysfunc("panicdottypeI")
	psess.
		panicindex = psess.sysfunc("panicindex")
	psess.
		panicnildottype = psess.sysfunc("panicnildottype")
	psess.
		panicslice = psess.sysfunc("panicslice")
	psess.
		raceread = psess.sysfunc("raceread")
	psess.
		racereadrange = psess.sysfunc("racereadrange")
	psess.
		racewrite = psess.sysfunc("racewrite")
	psess.
		racewriterange = psess.sysfunc("racewriterange")
	psess.
		supportPopcnt = psess.sysfunc("support_popcnt")
	psess.
		supportSSE41 = psess.sysfunc("support_sse41")
	psess.
		arm64SupportAtomics = psess.sysfunc("arm64_support_atomics")
	psess.
		typedmemclr = psess.sysfunc("typedmemclr")
	psess.
		typedmemmove = psess.sysfunc("typedmemmove")
	psess.
		Udiv = psess.sysfunc("udiv")
	psess.
		writeBarrier = psess.sysfunc("writeBarrier")
	psess.
		ControlWord64trunc = psess.sysfunc("controlWord64trunc")
	psess.
		ControlWord32 = psess.sysfunc("controlWord32")
	psess.
		WasmMove = psess.sysfunc("wasmMove")
	psess.
		WasmZero = psess.sysfunc("wasmZero")
	psess.
		WasmDiv = psess.sysfunc("wasmDiv")
	psess.
		WasmTruncS = psess.sysfunc("wasmTruncS")
	psess.
		WasmTruncU = psess.sysfunc("wasmTruncU")
	psess.
		SigPanic = psess.sysfunc("sigpanic")
}

// buildssa builds an SSA function for fn.
// worker indicates which of the backend workers is doing the processing.
func (psess *PackageSession) buildssa(fn *Node, worker int) *ssa.Func {
	name := fn.funcname()
	printssa := name == os.Getenv("GOSSAFUNC")
	if printssa {
		fmt.Println("generating SSA for", name)
		dumplist("buildssa-enter", fn.Func.Enter)
		dumplist("buildssa-body", fn.Nbody)
		dumplist("buildssa-exit", fn.Func.Exit)
	}

	var s state
	s.pushLine(psess, fn.Pos)
	defer s.popLine()

	s.hasdefer = fn.Func.HasDefer()
	if fn.Func.Pragma&CgoUnsafeArgs != 0 {
		s.cgoUnsafeArgs = true
	}

	fe := ssafn{
		curfn: fn,
		log:   printssa,
	}
	s.curfn = fn

	s.f = ssa.NewFunc(&fe)
	s.config = psess.ssaConfig
	s.f.Type = fn.Type
	s.f.Config = psess.ssaConfig
	s.f.Cache = &psess.ssaCaches[worker]
	s.f.Cache.Reset()
	s.f.DebugTest = s.f.DebugHashMatch("GOSSAHASH", name)
	s.f.Name = name
	if fn.Func.Pragma&Nosplit != 0 {
		s.f.NoSplit = true
	}
	s.panics = map[funcLine]*ssa.Block{}
	s.softFloat = s.config.SoftFloat

	if name == os.Getenv("GOSSAFUNC") {
		s.f.HTMLWriter = psess.ssa.NewHTMLWriter("ssa.html", s.f.Frontend(), name)

	}

	s.f.Entry = s.f.NewBlock(ssa.BlockPlain)

	s.labels = map[string]*ssaLabel{}
	s.labeledNodes = map[*Node]*ssaLabel{}
	s.fwdVars = map[*Node]*ssa.Value{}
	s.startmem = s.entryNewValue0(psess, ssa.OpInitMem, psess.types.TypeMem)
	s.sp = s.entryNewValue0(psess, ssa.OpSP, psess.types.Types[TUINTPTR])
	s.sb = s.entryNewValue0(psess, ssa.OpSB, psess.types.Types[TUINTPTR])

	s.startBlock(s.f.Entry)
	s.vars[&psess.memVar] = s.startmem

	s.decladdrs = map[*Node]*ssa.Value{}
	for _, n := range fn.Func.Dcl {
		switch n.Class() {
		case PPARAM, PPARAMOUT:
			s.decladdrs[n] = s.entryNewValue1A(psess, ssa.OpAddr, psess.types.NewPtr(n.Type), n, s.sp)
			if n.Class() == PPARAMOUT && s.canSSA(psess, n) {

				s.returns = append(s.returns, n)
			}
		case PAUTO:

		case PAUTOHEAP:

		case PFUNC:

		default:
			s.Fatalf("local variable with class %v unimplemented", n.Class())
		}
	}

	for _, n := range fn.Func.Dcl {
		if n.Class() == PPARAM && s.canSSA(psess, n) {
			s.vars[n] = s.newValue0A(ssa.OpArg, n.Type, n)
		}
	}

	s.stmtList(psess, fn.Func.Enter)
	s.stmtList(psess, fn.Nbody)

	if s.curBlock != nil {
		s.pushLine(psess, fn.Func.Endlineno)
		s.exit(psess)
		s.popLine()
	}

	for _, b := range s.f.Blocks {
		if b.Pos != psess.src.NoXPos {
			s.updateUnsetPredPos(psess, b)
		}
	}

	s.insertPhis(psess)
	psess.ssa.
		Compile(s.f)
	return s.f
}

// updateUnsetPredPos propagates the earliest-value position information for b
// towards all of b's predecessors that need a position, and recurs on that
// predecessor if its position is updated. B should have a non-empty position.
func (s *state) updateUnsetPredPos(psess *PackageSession, b *ssa.Block) {
	if b.Pos == psess.src.NoXPos {
		s.Fatalf("Block %s should have a position", b)
	}
	bestPos := psess.src.NoXPos
	for _, e := range b.Preds {
		p := e.Block()
		if !p.LackingPos(psess.ssa) {
			continue
		}
		if bestPos == psess.src.NoXPos {
			bestPos = b.Pos
			for _, v := range b.Values {
				if v.LackingPos(psess.ssa) {
					continue
				}
				if v.Pos != psess.src.NoXPos {

					bestPos = v.Pos
					break
				}
			}
		}
		p.Pos = bestPos
		s.updateUnsetPredPos(psess, p)
	}
}

type state struct {
	// configuration (arch) information
	config *ssa.Config

	// function we're building
	f *ssa.Func

	// Node for function
	curfn *Node

	// labels and labeled control flow nodes (OFOR, OFORUNTIL, OSWITCH, OSELECT) in f
	labels       map[string]*ssaLabel
	labeledNodes map[*Node]*ssaLabel

	// unlabeled break and continue statement tracking
	breakTo    *ssa.Block // current target for plain break statement
	continueTo *ssa.Block // current target for plain continue statement

	// current location where we're interpreting the AST
	curBlock *ssa.Block

	// variable assignments in the current block (map from variable symbol to ssa value)
	// *Node is the unique identifier (an ONAME Node) for the variable.
	// TODO: keep a single varnum map, then make all of these maps slices instead?
	vars map[*Node]*ssa.Value

	// fwdVars are variables that are used before they are defined in the current block.
	// This map exists just to coalesce multiple references into a single FwdRef op.
	// *Node is the unique identifier (an ONAME Node) for the variable.
	fwdVars map[*Node]*ssa.Value

	// all defined variables at the end of each block. Indexed by block ID.
	defvars []map[*Node]*ssa.Value

	// addresses of PPARAM and PPARAMOUT variables.
	decladdrs map[*Node]*ssa.Value

	// starting values. Memory, stack pointer, and globals pointer
	startmem *ssa.Value
	sp       *ssa.Value
	sb       *ssa.Value

	// line number stack. The current line number is top of stack
	line []src.XPos
	// the last line number processed; it may have been popped
	lastPos src.XPos

	// list of panic calls by function name and line number.
	// Used to deduplicate panic calls.
	panics map[funcLine]*ssa.Block

	// list of PPARAMOUT (return) variables.
	returns []*Node

	cgoUnsafeArgs bool
	hasdefer      bool // whether the function contains a defer statement
	softFloat     bool
}

type funcLine struct {
	f    *obj.LSym
	base *src.PosBase
	line uint
}

type ssaLabel struct {
	target         *ssa.Block // block identified by this label
	breakTarget    *ssa.Block // block to break to in control flow node identified by this label
	continueTarget *ssa.Block // block to continue to in control flow node identified by this label
}

// label returns the label associated with sym, creating it if necessary.
func (s *state) label(sym *types.Sym) *ssaLabel {
	lab := s.labels[sym.Name]
	if lab == nil {
		lab = new(ssaLabel)
		s.labels[sym.Name] = lab
	}
	return lab
}

func (s *state) Logf(msg string, args ...interface{}) { s.f.Logf(msg, args...) }
func (s *state) Log() bool                            { return s.f.Log() }
func (s *state) Fatalf(msg string, args ...interface{}) {
	s.f.Frontend().Fatalf(s.peekPos(), msg, args...)
}
func (s *state) Warnl(pos src.XPos, msg string, args ...interface{}) { s.f.Warnl(pos, msg, args...) }
func (s *state) Debug_checknil() bool                                { return s.f.Frontend().Debug_checknil() }

// dummy node for the memory variable

// dummy nodes for temporary variables

// startBlock sets the current block we're generating code in to b.
func (s *state) startBlock(b *ssa.Block) {
	if s.curBlock != nil {
		s.Fatalf("starting block %v when block %v has not ended", b, s.curBlock)
	}
	s.curBlock = b
	s.vars = map[*Node]*ssa.Value{}
	for n := range s.fwdVars {
		delete(s.fwdVars, n)
	}
}

// endBlock marks the end of generating code for the current block.
// Returns the (former) current block. Returns nil if there is no current
// block, i.e. if no code flows to the current execution point.
func (s *state) endBlock(psess *PackageSession) *ssa.Block {
	b := s.curBlock
	if b == nil {
		return nil
	}
	for len(s.defvars) <= int(b.ID) {
		s.defvars = append(s.defvars, nil)
	}
	s.defvars[b.ID] = s.vars
	s.curBlock = nil
	s.vars = nil
	if b.LackingPos(psess.ssa) {

		b.Pos = psess.src.NoXPos
	} else {
		b.Pos = s.lastPos
	}
	return b
}

// pushLine pushes a line number on the line number stack.
func (s *state) pushLine(psess *PackageSession, line src.XPos) {
	if !line.IsKnown() {

		line = s.peekPos()
		if psess.Debug['K'] != 0 {
			psess.
				Warn("buildssa: unknown position (line 0)")
		}
	} else {
		s.lastPos = line
	}

	s.line = append(s.line, line)
}

// popLine pops the top of the line number stack.
func (s *state) popLine() {
	s.line = s.line[:len(s.line)-1]
}

// peekPos peeks the top of the line number stack.
func (s *state) peekPos() src.XPos {
	return s.line[len(s.line)-1]
}

// newValue0 adds a new value with no arguments to the current block.
func (s *state) newValue0(op ssa.Op, t *types.Type) *ssa.Value {
	return s.curBlock.NewValue0(s.peekPos(), op, t)
}

// newValue0A adds a new value with no arguments and an aux value to the current block.
func (s *state) newValue0A(op ssa.Op, t *types.Type, aux interface{}) *ssa.Value {
	return s.curBlock.NewValue0A(s.peekPos(), op, t, aux)
}

// newValue0I adds a new value with no arguments and an auxint value to the current block.
func (s *state) newValue0I(op ssa.Op, t *types.Type, auxint int64) *ssa.Value {
	return s.curBlock.NewValue0I(s.peekPos(), op, t, auxint)
}

// newValue1 adds a new value with one argument to the current block.
func (s *state) newValue1(op ssa.Op, t *types.Type, arg *ssa.Value) *ssa.Value {
	return s.curBlock.NewValue1(s.peekPos(), op, t, arg)
}

// newValue1A adds a new value with one argument and an aux value to the current block.
func (s *state) newValue1A(op ssa.Op, t *types.Type, aux interface{}, arg *ssa.Value) *ssa.Value {
	return s.curBlock.NewValue1A(s.peekPos(), op, t, aux, arg)
}

// newValue1Apos adds a new value with one argument and an aux value to the current block.
// isStmt determines whether the created values may be a statement or not
// (i.e., false means never, yes means maybe).
func (s *state) newValue1Apos(op ssa.Op, t *types.Type, aux interface{}, arg *ssa.Value, isStmt bool) *ssa.Value {
	if isStmt {
		return s.curBlock.NewValue1A(s.peekPos(), op, t, aux, arg)
	}
	return s.curBlock.NewValue1A(s.peekPos().WithNotStmt(), op, t, aux, arg)
}

// newValue1I adds a new value with one argument and an auxint value to the current block.
func (s *state) newValue1I(op ssa.Op, t *types.Type, aux int64, arg *ssa.Value) *ssa.Value {
	return s.curBlock.NewValue1I(s.peekPos(), op, t, aux, arg)
}

// newValue2 adds a new value with two arguments to the current block.
func (s *state) newValue2(op ssa.Op, t *types.Type, arg0, arg1 *ssa.Value) *ssa.Value {
	return s.curBlock.NewValue2(s.peekPos(), op, t, arg0, arg1)
}

// newValue2I adds a new value with two arguments and an auxint value to the current block.
func (s *state) newValue2I(op ssa.Op, t *types.Type, aux int64, arg0, arg1 *ssa.Value) *ssa.Value {
	return s.curBlock.NewValue2I(s.peekPos(), op, t, aux, arg0, arg1)
}

// newValue3 adds a new value with three arguments to the current block.
func (s *state) newValue3(op ssa.Op, t *types.Type, arg0, arg1, arg2 *ssa.Value) *ssa.Value {
	return s.curBlock.NewValue3(s.peekPos(), op, t, arg0, arg1, arg2)
}

// newValue3I adds a new value with three arguments and an auxint value to the current block.
func (s *state) newValue3I(op ssa.Op, t *types.Type, aux int64, arg0, arg1, arg2 *ssa.Value) *ssa.Value {
	return s.curBlock.NewValue3I(s.peekPos(), op, t, aux, arg0, arg1, arg2)
}

// newValue3A adds a new value with three arguments and an aux value to the current block.
func (s *state) newValue3A(op ssa.Op, t *types.Type, aux interface{}, arg0, arg1, arg2 *ssa.Value) *ssa.Value {
	return s.curBlock.NewValue3A(s.peekPos(), op, t, aux, arg0, arg1, arg2)
}

// newValue3Apos adds a new value with three arguments and an aux value to the current block.
// isStmt determines whether the created values may be a statement or not
// (i.e., false means never, yes means maybe).
func (s *state) newValue3Apos(op ssa.Op, t *types.Type, aux interface{}, arg0, arg1, arg2 *ssa.Value, isStmt bool) *ssa.Value {
	if isStmt {
		return s.curBlock.NewValue3A(s.peekPos(), op, t, aux, arg0, arg1, arg2)
	}
	return s.curBlock.NewValue3A(s.peekPos().WithNotStmt(), op, t, aux, arg0, arg1, arg2)
}

// newValue4 adds a new value with four arguments to the current block.
func (s *state) newValue4(op ssa.Op, t *types.Type, arg0, arg1, arg2, arg3 *ssa.Value) *ssa.Value {
	return s.curBlock.NewValue4(s.peekPos(), op, t, arg0, arg1, arg2, arg3)
}

// entryNewValue0 adds a new value with no arguments to the entry block.
func (s *state) entryNewValue0(psess *PackageSession, op ssa.Op, t *types.Type) *ssa.Value {
	return s.f.Entry.NewValue0(psess.src.NoXPos, op, t)
}

// entryNewValue0A adds a new value with no arguments and an aux value to the entry block.
func (s *state) entryNewValue0A(psess *PackageSession, op ssa.Op, t *types.Type, aux interface{}) *ssa.Value {
	return s.f.Entry.NewValue0A(psess.src.NoXPos, op, t, aux)
}

// entryNewValue1 adds a new value with one argument to the entry block.
func (s *state) entryNewValue1(psess *PackageSession, op ssa.Op, t *types.Type, arg *ssa.Value) *ssa.Value {
	return s.f.Entry.NewValue1(psess.src.NoXPos, op, t, arg)
}

// entryNewValue1 adds a new value with one argument and an auxint value to the entry block.
func (s *state) entryNewValue1I(psess *PackageSession, op ssa.Op, t *types.Type, auxint int64, arg *ssa.Value) *ssa.Value {
	return s.f.Entry.NewValue1I(psess.src.NoXPos, op, t, auxint, arg)
}

// entryNewValue1A adds a new value with one argument and an aux value to the entry block.
func (s *state) entryNewValue1A(psess *PackageSession, op ssa.Op, t *types.Type, aux interface{}, arg *ssa.Value) *ssa.Value {
	return s.f.Entry.NewValue1A(psess.src.NoXPos, op, t, aux, arg)
}

// entryNewValue2 adds a new value with two arguments to the entry block.
func (s *state) entryNewValue2(psess *PackageSession, op ssa.Op, t *types.Type, arg0, arg1 *ssa.Value) *ssa.Value {
	return s.f.Entry.NewValue2(psess.src.NoXPos, op, t, arg0, arg1)
}

// const* routines add a new const value to the entry block.
func (s *state) constSlice(psess *PackageSession, t *types.Type) *ssa.Value {
	return s.f.ConstSlice(psess.ssa, t)
}
func (s *state) constInterface(psess *PackageSession, t *types.Type) *ssa.Value {
	return s.f.ConstInterface(psess.ssa, t)
}
func (s *state) constNil(psess *PackageSession, t *types.Type) *ssa.Value {
	return s.f.ConstNil(psess.ssa, t)
}
func (s *state) constEmptyString(psess *PackageSession, t *types.Type) *ssa.Value {
	return s.f.ConstEmptyString(psess.ssa, t)
}
func (s *state) constBool(psess *PackageSession, c bool) *ssa.Value {
	return s.f.ConstBool(psess.ssa, psess.types.Types[TBOOL], c)
}
func (s *state) constInt8(psess *PackageSession, t *types.Type, c int8) *ssa.Value {
	return s.f.ConstInt8(psess.ssa, t, c)
}
func (s *state) constInt16(psess *PackageSession, t *types.Type, c int16) *ssa.Value {
	return s.f.ConstInt16(psess.ssa, t, c)
}
func (s *state) constInt32(psess *PackageSession, t *types.Type, c int32) *ssa.Value {
	return s.f.ConstInt32(psess.ssa, t, c)
}
func (s *state) constInt64(psess *PackageSession, t *types.Type, c int64) *ssa.Value {
	return s.f.ConstInt64(psess.ssa, t, c)
}
func (s *state) constFloat32(psess *PackageSession, t *types.Type, c float64) *ssa.Value {
	return s.f.ConstFloat32(psess.ssa, t, c)
}
func (s *state) constFloat64(psess *PackageSession, t *types.Type, c float64) *ssa.Value {
	return s.f.ConstFloat64(psess.ssa, t, c)
}
func (s *state) constInt(psess *PackageSession, t *types.Type, c int64) *ssa.Value {
	if s.config.PtrSize == 8 {
		return s.constInt64(psess, t, c)
	}
	if int64(int32(c)) != c {
		s.Fatalf("integer constant too big %d", c)
	}
	return s.constInt32(psess, t, int32(c))
}
func (s *state) constOffPtrSP(psess *PackageSession, t *types.Type, c int64) *ssa.Value {
	return s.f.ConstOffPtrSP(psess.ssa, t, c, s.sp)
}

// newValueOrSfCall* are wrappers around newValue*, which may create a call to a
// soft-float runtime function instead (when emitting soft-float code).
func (s *state) newValueOrSfCall1(psess *PackageSession, op ssa.Op, t *types.Type, arg *ssa.Value) *ssa.Value {
	if s.softFloat {
		if c, ok := s.sfcall(psess, op, arg); ok {
			return c
		}
	}
	return s.newValue1(op, t, arg)
}
func (s *state) newValueOrSfCall2(psess *PackageSession, op ssa.Op, t *types.Type, arg0, arg1 *ssa.Value) *ssa.Value {
	if s.softFloat {
		if c, ok := s.sfcall(psess, op, arg0, arg1); ok {
			return c
		}
	}
	return s.newValue2(op, t, arg0, arg1)
}

func (s *state) instrument(psess *PackageSession, t *types.Type, addr *ssa.Value, wr bool) {
	if !s.curfn.Func.InstrumentBody() {
		return
	}

	w := t.Size(psess.types)
	if w == 0 {
		return
	}

	if ssa.IsSanitizerSafeAddr(addr) {
		return
	}

	var fn *obj.LSym
	needWidth := false

	if psess.flag_msan {
		fn = psess.msanread
		if wr {
			fn = psess.msanwrite
		}
		needWidth = true
	} else if psess.flag_race && t.NumComponents(psess.types, types.CountBlankFields) > 1 {

		fn = psess.racereadrange
		if wr {
			fn = psess.racewriterange
		}
		needWidth = true
	} else if psess.flag_race {

		fn = psess.raceread
		if wr {
			fn = psess.racewrite
		}
	} else {
		panic("unreachable")
	}

	args := []*ssa.Value{addr}
	if needWidth {
		args = append(args, s.constInt(psess, psess.types.Types[TUINTPTR], w))
	}
	s.rtcall(psess, fn, true, nil, args...)
}

func (s *state) load(psess *PackageSession, t *types.Type, src *ssa.Value) *ssa.Value {
	s.instrument(psess, t, src, false)
	return s.rawLoad(psess, t, src)
}

func (s *state) rawLoad(psess *PackageSession, t *types.Type, src *ssa.Value) *ssa.Value {
	return s.newValue2(ssa.OpLoad, t, src, s.mem(psess))
}

func (s *state) store(psess *PackageSession, t *types.Type, dst, val *ssa.Value) {
	s.vars[&psess.memVar] = s.newValue3A(ssa.OpStore, psess.types.TypeMem, t, dst, val, s.mem(psess))
}

func (s *state) zero(psess *PackageSession, t *types.Type, dst *ssa.Value) {
	s.instrument(psess, t, dst, true)
	store := s.newValue2I(ssa.OpZero, psess.types.TypeMem, t.Size(psess.types), dst, s.mem(psess))
	store.Aux = t
	s.vars[&psess.memVar] = store
}

func (s *state) move(psess *PackageSession, t *types.Type, dst, src *ssa.Value) {
	s.instrument(psess, t, src, false)
	s.instrument(psess, t, dst, true)
	store := s.newValue3I(ssa.OpMove, psess.types.TypeMem, t.Size(psess.types), dst, src, s.mem(psess))
	store.Aux = t
	s.vars[&psess.memVar] = store
}

// stmtList converts the statement list n to SSA and adds it to s.
func (s *state) stmtList(psess *PackageSession, l Nodes) {
	for _, n := range l.Slice() {
		s.stmt(psess, n)
	}
}

// stmt converts the statement n to SSA and adds it to s.
func (s *state) stmt(psess *PackageSession, n *Node) {
	if !(n.Op == OVARKILL || n.Op == OVARLIVE) {

		s.pushLine(psess, n.Pos)
		defer s.popLine()
	}

	if s.curBlock == nil && n.Op != OLABEL {
		return
	}

	s.stmtList(psess, n.Ninit)
	switch n.Op {

	case OBLOCK:
		s.stmtList(psess, n.List)

	case OEMPTY, ODCLCONST, ODCLTYPE, OFALL:

	case OCALLFUNC:
		if psess.isIntrinsicCall(n) {
			s.intrinsicCall(psess, n)
			return
		}
		fallthrough

	case OCALLMETH, OCALLINTER:
		s.call(psess, n, callNormal)
		if n.Op == OCALLFUNC && n.Left.Op == ONAME && n.Left.Class() == PFUNC {
			if fn := n.Left.Sym.Name; psess.compiling_runtime && fn == "throw" ||
				n.Left.Sym.Pkg == psess.Runtimepkg && (fn == "throwinit" || fn == "gopanic" || fn == "panicwrap" || fn == "block" || fn == "panicmakeslicelen" || fn == "panicmakeslicecap") {
				m := s.mem(psess)
				b := s.endBlock(psess)
				b.Kind = ssa.BlockExit
				b.SetControl(m)

			}
		}
	case ODEFER:
		s.call(psess, n.Left, callDefer)
	case OPROC:
		s.call(psess, n.Left, callGo)

	case OAS2DOTTYPE:
		res, resok := s.dottype(psess, n.Rlist.First(), true)
		deref := false
		if !psess.canSSAType(n.Rlist.First().Type) {
			if res.Op != ssa.OpLoad {
				s.Fatalf("dottype of non-load")
			}
			mem := s.mem(psess)
			if mem.Op == ssa.OpVarKill {
				mem = mem.Args[0]
			}
			if res.Args[1] != mem {
				s.Fatalf("memory no longer live from 2-result dottype load")
			}
			deref = true
			res = res.Args[0]
		}
		s.assign(psess, n.List.First(), res, deref, 0)
		s.assign(psess, n.List.Second(), resok, false, 0)
		return

	case OAS2FUNC:

		if !psess.isIntrinsicCall(n.Rlist.First()) {
			s.Fatalf("non-intrinsic AS2FUNC not expanded %v", n.Rlist.First())
		}
		v := s.intrinsicCall(psess, n.Rlist.First())
		v1 := s.newValue1(ssa.OpSelect0, n.List.First().Type, v)
		v2 := s.newValue1(ssa.OpSelect1, n.List.Second().Type, v)
		s.assign(psess, n.List.First(), v1, false, 0)
		s.assign(psess, n.List.Second(), v2, false, 0)
		return

	case ODCL:
		if n.Left.Class() == PAUTOHEAP {
			psess.
				Fatalf("DCL %v", n)
		}

	case OLABEL:
		sym := n.Left.Sym
		lab := s.label(sym)

		if ctl := n.labeledControl(psess); ctl != nil {
			s.labeledNodes[ctl] = lab
		}

		if lab.target == nil {
			lab.target = s.f.NewBlock(ssa.BlockPlain)
		}

		if s.curBlock != nil {
			b := s.endBlock(psess)
			b.AddEdgeTo(lab.target)
		}
		s.startBlock(lab.target)

	case OGOTO:
		sym := n.Left.Sym

		lab := s.label(sym)
		if lab.target == nil {
			lab.target = s.f.NewBlock(ssa.BlockPlain)
		}

		b := s.endBlock(psess)
		b.Pos = s.lastPos.WithIsStmt()
		b.AddEdgeTo(lab.target)

	case OAS:
		if n.Left == n.Right && n.Left.Op == ONAME {

			return
		}

		rhs := n.Right
		if rhs != nil {
			switch rhs.Op {
			case OSTRUCTLIT, OARRAYLIT, OSLICELIT:

				if !psess.isZero(rhs) {
					psess.
						Fatalf("literal with nonzero value in SSA: %v", rhs)
				}
				rhs = nil
			case OAPPEND:

				if !psess.samesafeexpr(n.Left, rhs.List.First()) || psess.Debug['N'] != 0 {
					break
				}

				if s.canSSA(psess, n.Left) {
					if psess.Debug_append > 0 {
						psess.
							Warnl(n.Pos, "append: len-only update (in local slice)")
					}
					break
				}
				if psess.Debug_append > 0 {
					psess.
						Warnl(n.Pos, "append: len-only update")
				}
				s.append(psess, rhs, true)
				return
			}
		}

		if n.Left.isBlank() {

			if rhs != nil {
				s.expr(psess, rhs)
			}
			return
		}

		var t *types.Type
		if n.Right != nil {
			t = n.Right.Type
		} else {
			t = n.Left.Type
		}

		var r *ssa.Value
		deref := !psess.canSSAType(t)
		if deref {
			if rhs == nil {
				r = nil
			} else {
				r = s.addr(psess, rhs, false)
			}
		} else {
			if rhs == nil {
				r = s.zeroVal(psess, t)
			} else {
				r = s.expr(psess, rhs)
			}
		}

		var skip skipMask
		if rhs != nil && (rhs.Op == OSLICE || rhs.Op == OSLICE3 || rhs.Op == OSLICESTR) && psess.samesafeexpr(rhs.Left, n.Left) {

			i, j, k := rhs.SliceBounds(psess)
			if i != nil && (i.Op == OLITERAL && i.Val().Ctype(psess) == CTINT && i.Int64(psess) == 0) {

				i = nil
			}

			if i == nil {
				skip |= skipPtr
				if j == nil {
					skip |= skipLen
				}
				if k == nil {
					skip |= skipCap
				}
			}
		}

		s.assign(psess, n.Left, r, deref, skip)

	case OIF:
		bThen := s.f.NewBlock(ssa.BlockPlain)
		bEnd := s.f.NewBlock(ssa.BlockPlain)
		var bElse *ssa.Block
		var likely int8
		if n.Likely() {
			likely = 1
		}
		if n.Rlist.Len() != 0 {
			bElse = s.f.NewBlock(ssa.BlockPlain)
			s.condBranch(psess, n.Left, bThen, bElse, likely)
		} else {
			s.condBranch(psess, n.Left, bThen, bEnd, likely)
		}

		s.startBlock(bThen)
		s.stmtList(psess, n.Nbody)
		if b := s.endBlock(psess); b != nil {
			b.AddEdgeTo(bEnd)
		}

		if n.Rlist.Len() != 0 {
			s.startBlock(bElse)
			s.stmtList(psess, n.Rlist)
			if b := s.endBlock(psess); b != nil {
				b.AddEdgeTo(bEnd)
			}
		}
		s.startBlock(bEnd)

	case ORETURN:
		s.stmtList(psess, n.List)
		b := s.exit(psess)
		b.Pos = s.lastPos.WithIsStmt()

	case ORETJMP:
		s.stmtList(psess, n.List)
		b := s.exit(psess)
		b.Kind = ssa.BlockRetJmp
		b.Aux = n.Sym.Linksym(psess.types)

	case OCONTINUE, OBREAK:
		var to *ssa.Block
		if n.Left == nil {

			switch n.Op {
			case OCONTINUE:
				to = s.continueTo
			case OBREAK:
				to = s.breakTo
			}
		} else {

			sym := n.Left.Sym
			lab := s.label(sym)
			switch n.Op {
			case OCONTINUE:
				to = lab.continueTarget
			case OBREAK:
				to = lab.breakTarget
			}
		}

		b := s.endBlock(psess)
		b.Pos = s.lastPos.WithIsStmt()
		b.AddEdgeTo(to)

	case OFOR, OFORUNTIL:

		bCond := s.f.NewBlock(ssa.BlockPlain)
		bBody := s.f.NewBlock(ssa.BlockPlain)
		bIncr := s.f.NewBlock(ssa.BlockPlain)
		bEnd := s.f.NewBlock(ssa.BlockPlain)

		b := s.endBlock(psess)
		if n.Op == OFOR {
			b.AddEdgeTo(bCond)

			s.startBlock(bCond)
			if n.Left != nil {
				s.condBranch(psess, n.Left, bBody, bEnd, 1)
			} else {
				b := s.endBlock(psess)
				b.Kind = ssa.BlockPlain
				b.AddEdgeTo(bBody)
			}

		} else {
			b.AddEdgeTo(bBody)
		}

		prevContinue := s.continueTo
		prevBreak := s.breakTo
		s.continueTo = bIncr
		s.breakTo = bEnd
		lab := s.labeledNodes[n]
		if lab != nil {

			lab.continueTarget = bIncr
			lab.breakTarget = bEnd
		}

		s.startBlock(bBody)
		s.stmtList(psess, n.Nbody)

		s.continueTo = prevContinue
		s.breakTo = prevBreak
		if lab != nil {
			lab.continueTarget = nil
			lab.breakTarget = nil
		}

		if b := s.endBlock(psess); b != nil {
			b.AddEdgeTo(bIncr)
		}

		s.startBlock(bIncr)
		if n.Right != nil {
			s.stmt(psess, n.Right)
		}
		if n.Op == OFOR {
			if b := s.endBlock(psess); b != nil {
				b.AddEdgeTo(bCond)

				if n.Op != OFORUNTIL && b.Pos == psess.src.NoXPos {
					b.Pos = bCond.Pos
				}
			}
		} else {

			bLateIncr := bCond

			s.condBranch(psess, n.Left, bLateIncr, bEnd, 1)

			s.startBlock(bLateIncr)
			s.stmtList(psess, n.List)
			s.endBlock(psess).AddEdgeTo(bBody)
		}

		s.startBlock(bEnd)

	case OSWITCH, OSELECT:

		bEnd := s.f.NewBlock(ssa.BlockPlain)

		prevBreak := s.breakTo
		s.breakTo = bEnd
		lab := s.labeledNodes[n]
		if lab != nil {

			lab.breakTarget = bEnd
		}

		s.stmtList(psess, n.Nbody)

		s.breakTo = prevBreak
		if lab != nil {
			lab.breakTarget = nil
		}

		if s.curBlock != nil {
			m := s.mem(psess)
			b := s.endBlock(psess)
			b.Kind = ssa.BlockExit
			b.SetControl(m)
		}
		s.startBlock(bEnd)

	case OVARKILL:

		if !s.canSSA(psess, n.Left) {
			s.vars[&psess.memVar] = s.newValue1Apos(ssa.OpVarKill, psess.types.TypeMem, n.Left, s.mem(psess), false)
		}

	case OVARLIVE:

		if !n.Left.Addrtaken() {
			s.Fatalf("VARLIVE variable %v must have Addrtaken set", n.Left)
		}
		switch n.Left.Class() {
		case PAUTO, PPARAM, PPARAMOUT:
		default:
			s.Fatalf("VARLIVE variable %v must be Auto or Arg", n.Left)
		}
		s.vars[&psess.memVar] = s.newValue1A(ssa.OpVarLive, psess.types.TypeMem, n.Left, s.mem(psess))

	case OCHECKNIL:
		p := s.expr(psess, n.Left)
		s.nilCheck(psess, p)

	default:
		s.Fatalf("unhandled stmt %v", n.Op)
	}
}

// exit processes any code that needs to be generated just before returning.
// It returns a BlockRet block that ends the control flow. Its control value
// will be set to the final memory state.
func (s *state) exit(psess *PackageSession) *ssa.Block {
	if s.hasdefer {
		s.rtcall(psess, psess.Deferreturn, true, nil)
	}

	s.stmtList(psess, s.curfn.Func.Exit)

	for _, n := range s.returns {
		addr := s.decladdrs[n]
		val := s.variable(n, n.Type)
		s.vars[&psess.memVar] = s.newValue1A(ssa.OpVarDef, psess.types.TypeMem, n, s.mem(psess))
		s.store(psess, n.Type, addr, val)

	}

	m := s.mem(psess)
	b := s.endBlock(psess)
	b.Kind = ssa.BlockRet
	b.SetControl(m)
	return b
}

type opAndType struct {
	op    Op
	etype types.EType
}

func (s *state) concreteEtype(t *types.Type) types.EType {
	e := t.Etype
	switch e {
	default:
		return e
	case TINT:
		if s.config.PtrSize == 8 {
			return TINT64
		}
		return TINT32
	case TUINT:
		if s.config.PtrSize == 8 {
			return TUINT64
		}
		return TUINT32
	case TUINTPTR:
		if s.config.PtrSize == 8 {
			return TUINT64
		}
		return TUINT32
	}
}

func (s *state) ssaOp(psess *PackageSession, op Op, t *types.Type) ssa.Op {
	etype := s.concreteEtype(t)
	x, ok := psess.opToSSA[opAndType{op, etype}]
	if !ok {
		s.Fatalf("unhandled binary op %v %s", op, etype)
	}
	return x
}

func (psess *PackageSession) floatForComplex(t *types.Type) *types.Type {
	if t.Size(psess.types) == 8 {
		return psess.types.Types[TFLOAT32]
	} else {
		return psess.types.Types[TFLOAT64]
	}
}

type opAndTwoTypes struct {
	op     Op
	etype1 types.EType
	etype2 types.EType
}

type twoTypes struct {
	etype1 types.EType
	etype2 types.EType
}

type twoOpsAndType struct {
	op1              ssa.Op
	op2              ssa.Op
	intermediateType types.EType
}

// this map is used only for 32-bit arch, and only includes the difference
// on 32-bit arch, don't use int64<->float conversion for uint32

// uint64<->float conversions, only on machines that have intructions for that

func (s *state) ssaShiftOp(psess *PackageSession, op Op, t *types.Type, u *types.Type) ssa.Op {
	etype1 := s.concreteEtype(t)
	etype2 := s.concreteEtype(u)
	x, ok := psess.shiftOpToSSA[opAndTwoTypes{op, etype1, etype2}]
	if !ok {
		s.Fatalf("unhandled shift op %v etype=%s/%s", op, etype1, etype2)
	}
	return x
}

// expr converts the expression n to ssa, adds it to s and returns the ssa result.
func (s *state) expr(psess *PackageSession, n *Node) *ssa.Value {
	if !(n.Op == ONAME || n.Op == OLITERAL && n.Sym != nil) {

		s.pushLine(psess, n.Pos)
		defer s.popLine()
	}

	s.stmtList(psess, n.Ninit)
	switch n.Op {
	case OARRAYBYTESTRTMP:
		slice := s.expr(psess, n.Left)
		ptr := s.newValue1(ssa.OpSlicePtr, s.f.Config.Types.BytePtr, slice)
		len := s.newValue1(ssa.OpSliceLen, psess.types.Types[TINT], slice)
		return s.newValue2(ssa.OpStringMake, n.Type, ptr, len)
	case OSTRARRAYBYTETMP:
		str := s.expr(psess, n.Left)
		ptr := s.newValue1(ssa.OpStringPtr, s.f.Config.Types.BytePtr, str)
		len := s.newValue1(ssa.OpStringLen, psess.types.Types[TINT], str)
		return s.newValue3(ssa.OpSliceMake, n.Type, ptr, len, len)
	case OCFUNC:
		aux := n.Left.Sym.Linksym(psess.types)
		return s.entryNewValue1A(psess, ssa.OpAddr, n.Type, aux, s.sb)
	case ONAME:
		if n.Class() == PFUNC {

			sym := psess.funcsym(n.Sym).Linksym(psess.types)
			return s.entryNewValue1A(psess, ssa.OpAddr, psess.types.NewPtr(n.Type), sym, s.sb)
		}
		if s.canSSA(psess, n) {
			return s.variable(n, n.Type)
		}
		addr := s.addr(psess, n, false)
		return s.load(psess, n.Type, addr)
	case OCLOSUREVAR:
		addr := s.addr(psess, n, false)
		return s.load(psess, n.Type, addr)
	case OLITERAL:
		switch u := n.Val().U.(type) {
		case *Mpint:
			i := u.Int64(psess)
			switch n.Type.Size(psess.types) {
			case 1:
				return s.constInt8(psess, n.Type, int8(i))
			case 2:
				return s.constInt16(psess, n.Type, int16(i))
			case 4:
				return s.constInt32(psess, n.Type, int32(i))
			case 8:
				return s.constInt64(psess, n.Type, i)
			default:
				s.Fatalf("bad integer size %d", n.Type.Size(psess.types))
				return nil
			}
		case string:
			if u == "" {
				return s.constEmptyString(psess, n.Type)
			}
			return s.entryNewValue0A(psess, ssa.OpConstString, n.Type, u)
		case bool:
			return s.constBool(psess, u)
		case *NilVal:
			t := n.Type
			switch {
			case t.IsSlice():
				return s.constSlice(psess, t)
			case t.IsInterface():
				return s.constInterface(psess, t)
			default:
				return s.constNil(psess, t)
			}
		case *Mpflt:
			switch n.Type.Size(psess.types) {
			case 4:
				return s.constFloat32(psess, n.Type, u.Float32(psess))
			case 8:
				return s.constFloat64(psess, n.Type, u.Float64(psess))
			default:
				s.Fatalf("bad float size %d", n.Type.Size(psess.types))
				return nil
			}
		case *Mpcplx:
			r := &u.Real
			i := &u.Imag
			switch n.Type.Size(psess.types) {
			case 8:
				pt := psess.types.Types[TFLOAT32]
				return s.newValue2(ssa.OpComplexMake, n.Type,
					s.constFloat32(psess, pt, r.Float32(psess)),
					s.constFloat32(psess, pt, i.Float32(psess)))
			case 16:
				pt := psess.types.Types[TFLOAT64]
				return s.newValue2(ssa.OpComplexMake, n.Type,
					s.constFloat64(psess, pt, r.Float64(psess)),
					s.constFloat64(psess, pt, i.Float64(psess)))
			default:
				s.Fatalf("bad float size %d", n.Type.Size(psess.types))
				return nil
			}

		default:
			s.Fatalf("unhandled OLITERAL %v", n.Val().Ctype(psess))
			return nil
		}
	case OCONVNOP:
		to := n.Type
		from := n.Left.Type

		x := s.expr(psess, n.Left)

		if to.IsPtrShaped() != from.IsPtrShaped() {
			return s.newValue2(ssa.OpConvert, to, x, s.mem(psess))
		}

		v := s.newValue1(ssa.OpCopy, to, x)

		if to.Etype == TFUNC && from.IsPtrShaped() {
			return v
		}

		if from.Etype == to.Etype {
			return v
		}

		if to.Etype == TUNSAFEPTR && from.IsPtrShaped() || from.Etype == TUNSAFEPTR && to.IsPtrShaped() {
			return v
		}

		if to.Etype == TMAP && from.IsPtr() &&
			to.MapType(psess.types).Hmap == from.Elem(psess.types) {
			return v
		}
		psess.
			dowidth(from)
		psess.
			dowidth(to)
		if from.Width != to.Width {
			s.Fatalf("CONVNOP width mismatch %v (%d) -> %v (%d)\n", from, from.Width, to, to.Width)
			return nil
		}
		if etypesign(from.Etype) != etypesign(to.Etype) {
			s.Fatalf("CONVNOP sign mismatch %v (%s) -> %v (%s)\n", from, from.Etype, to, to.Etype)
			return nil
		}

		if psess.instrumenting {

			return v
		}

		if etypesign(from.Etype) == 0 {
			s.Fatalf("CONVNOP unrecognized non-integer %v -> %v\n", from, to)
			return nil
		}

		return v

	case OCONV:
		x := s.expr(psess, n.Left)
		ft := n.Left.Type
		tt := n.Type
		if ft.IsBoolean() && tt.IsKind(TUINT8) {

			return s.newValue1(ssa.OpCopy, n.Type, x)
		}
		if ft.IsInteger() && tt.IsInteger() {
			var op ssa.Op
			if tt.Size(psess.types) == ft.Size(psess.types) {
				op = ssa.OpCopy
			} else if tt.Size(psess.types) < ft.Size(psess.types) {

				switch 10*ft.Size(psess.types) + tt.Size(psess.types) {
				case 21:
					op = ssa.OpTrunc16to8
				case 41:
					op = ssa.OpTrunc32to8
				case 42:
					op = ssa.OpTrunc32to16
				case 81:
					op = ssa.OpTrunc64to8
				case 82:
					op = ssa.OpTrunc64to16
				case 84:
					op = ssa.OpTrunc64to32
				default:
					s.Fatalf("weird integer truncation %v -> %v", ft, tt)
				}
			} else if ft.IsSigned() {

				switch 10*ft.Size(psess.types) + tt.Size(psess.types) {
				case 12:
					op = ssa.OpSignExt8to16
				case 14:
					op = ssa.OpSignExt8to32
				case 18:
					op = ssa.OpSignExt8to64
				case 24:
					op = ssa.OpSignExt16to32
				case 28:
					op = ssa.OpSignExt16to64
				case 48:
					op = ssa.OpSignExt32to64
				default:
					s.Fatalf("bad integer sign extension %v -> %v", ft, tt)
				}
			} else {

				switch 10*ft.Size(psess.types) + tt.Size(psess.types) {
				case 12:
					op = ssa.OpZeroExt8to16
				case 14:
					op = ssa.OpZeroExt8to32
				case 18:
					op = ssa.OpZeroExt8to64
				case 24:
					op = ssa.OpZeroExt16to32
				case 28:
					op = ssa.OpZeroExt16to64
				case 48:
					op = ssa.OpZeroExt32to64
				default:
					s.Fatalf("weird integer sign extension %v -> %v", ft, tt)
				}
			}
			return s.newValue1(op, n.Type, x)
		}

		if ft.IsFloat() || tt.IsFloat() {
			conv, ok := psess.fpConvOpToSSA[twoTypes{s.concreteEtype(ft), s.concreteEtype(tt)}]
			if s.config.RegSize == 4 && psess.thearch.LinkArch.Family != sys.MIPS && !s.softFloat {
				if conv1, ok1 := psess.fpConvOpToSSA32[twoTypes{s.concreteEtype(ft), s.concreteEtype(tt)}]; ok1 {
					conv = conv1
				}
			}
			if psess.thearch.LinkArch.Family == sys.ARM64 || psess.thearch.LinkArch.Family == sys.Wasm || s.softFloat {
				if conv1, ok1 := psess.uint64fpConvOpToSSA[twoTypes{s.concreteEtype(ft), s.concreteEtype(tt)}]; ok1 {
					conv = conv1
				}
			}

			if psess.thearch.LinkArch.Family == sys.MIPS && !s.softFloat {
				if ft.Size(psess.types) == 4 && ft.IsInteger() && !ft.IsSigned() {

					if tt.Size(psess.types) == 4 {
						return s.uint32Tofloat32(psess, n, x, ft, tt)
					}
					if tt.Size(psess.types) == 8 {
						return s.uint32Tofloat64(psess, n, x, ft, tt)
					}
				} else if tt.Size(psess.types) == 4 && tt.IsInteger() && !tt.IsSigned() {

					if ft.Size(psess.types) == 4 {
						return s.float32ToUint32(psess, n, x, ft, tt)
					}
					if ft.Size(psess.types) == 8 {
						return s.float64ToUint32(psess, n, x, ft, tt)
					}
				}
			}

			if !ok {
				s.Fatalf("weird float conversion %v -> %v", ft, tt)
			}
			op1, op2, it := conv.op1, conv.op2, conv.intermediateType

			if op1 != ssa.OpInvalid && op2 != ssa.OpInvalid {

				if op1 == ssa.OpCopy {
					if op2 == ssa.OpCopy {
						return x
					}
					return s.newValueOrSfCall1(psess, op2, n.Type, x)
				}
				if op2 == ssa.OpCopy {
					return s.newValueOrSfCall1(psess, op1, n.Type, x)
				}
				return s.newValueOrSfCall1(psess, op2, n.Type, s.newValueOrSfCall1(psess, op1, psess.types.Types[it], x))
			}

			if ft.IsInteger() {

				if tt.Size(psess.types) == 4 {
					return s.uint64Tofloat32(psess, n, x, ft, tt)
				}
				if tt.Size(psess.types) == 8 {
					return s.uint64Tofloat64(psess, n, x, ft, tt)
				}
				s.Fatalf("weird unsigned integer to float conversion %v -> %v", ft, tt)
			}

			if ft.Size(psess.types) == 4 {
				return s.float32ToUint64(psess, n, x, ft, tt)
			}
			if ft.Size(psess.types) == 8 {
				return s.float64ToUint64(psess, n, x, ft, tt)
			}
			s.Fatalf("weird float to unsigned integer conversion %v -> %v", ft, tt)
			return nil
		}

		if ft.IsComplex() && tt.IsComplex() {
			var op ssa.Op
			if ft.Size(psess.types) == tt.Size(psess.types) {
				switch ft.Size(psess.types) {
				case 8:
					op = ssa.OpRound32F
				case 16:
					op = ssa.OpRound64F
				default:
					s.Fatalf("weird complex conversion %v -> %v", ft, tt)
				}
			} else if ft.Size(psess.types) == 8 && tt.Size(psess.types) == 16 {
				op = ssa.OpCvt32Fto64F
			} else if ft.Size(psess.types) == 16 && tt.Size(psess.types) == 8 {
				op = ssa.OpCvt64Fto32F
			} else {
				s.Fatalf("weird complex conversion %v -> %v", ft, tt)
			}
			ftp := psess.floatForComplex(ft)
			ttp := psess.floatForComplex(tt)
			return s.newValue2(ssa.OpComplexMake, tt,
				s.newValueOrSfCall1(psess, op, ttp, s.newValue1(ssa.OpComplexReal, ftp, x)),
				s.newValueOrSfCall1(psess, op, ttp, s.newValue1(ssa.OpComplexImag, ftp, x)))
		}

		s.Fatalf("unhandled OCONV %s -> %s", n.Left.Type.Etype, n.Type.Etype)
		return nil

	case ODOTTYPE:
		res, _ := s.dottype(psess, n, false)
		return res

	case OLT, OEQ, ONE, OLE, OGE, OGT:
		a := s.expr(psess, n.Left)
		b := s.expr(psess, n.Right)
		if n.Left.Type.IsComplex() {
			pt := psess.floatForComplex(n.Left.Type)
			op := s.ssaOp(psess, OEQ, pt)
			r := s.newValueOrSfCall2(psess, op, psess.types.Types[TBOOL], s.newValue1(ssa.OpComplexReal, pt, a), s.newValue1(ssa.OpComplexReal, pt, b))
			i := s.newValueOrSfCall2(psess, op, psess.types.Types[TBOOL], s.newValue1(ssa.OpComplexImag, pt, a), s.newValue1(ssa.OpComplexImag, pt, b))
			c := s.newValue2(ssa.OpAndB, psess.types.Types[TBOOL], r, i)
			switch n.Op {
			case OEQ:
				return c
			case ONE:
				return s.newValue1(ssa.OpNot, psess.types.Types[TBOOL], c)
			default:
				s.Fatalf("ordered complex compare %v", n.Op)
			}
		}
		if n.Left.Type.IsFloat() {
			return s.newValueOrSfCall2(psess, s.ssaOp(psess, n.Op, n.Left.Type), psess.types.Types[TBOOL], a, b)
		}
		return s.newValue2(s.ssaOp(psess, n.Op, n.Left.Type), psess.types.Types[TBOOL], a, b)
	case OMUL:
		a := s.expr(psess, n.Left)
		b := s.expr(psess, n.Right)
		if n.Type.IsComplex() {
			mulop := ssa.OpMul64F
			addop := ssa.OpAdd64F
			subop := ssa.OpSub64F
			pt := psess.floatForComplex(n.Type)
			wt := psess.types.Types[TFLOAT64]

			areal := s.newValue1(ssa.OpComplexReal, pt, a)
			breal := s.newValue1(ssa.OpComplexReal, pt, b)
			aimag := s.newValue1(ssa.OpComplexImag, pt, a)
			bimag := s.newValue1(ssa.OpComplexImag, pt, b)

			if pt != wt {
				areal = s.newValueOrSfCall1(psess, ssa.OpCvt32Fto64F, wt, areal)
				breal = s.newValueOrSfCall1(psess, ssa.OpCvt32Fto64F, wt, breal)
				aimag = s.newValueOrSfCall1(psess, ssa.OpCvt32Fto64F, wt, aimag)
				bimag = s.newValueOrSfCall1(psess, ssa.OpCvt32Fto64F, wt, bimag)
			}

			xreal := s.newValueOrSfCall2(psess, subop, wt, s.newValueOrSfCall2(psess, mulop, wt, areal, breal), s.newValueOrSfCall2(psess, mulop, wt, aimag, bimag))
			ximag := s.newValueOrSfCall2(psess, addop, wt, s.newValueOrSfCall2(psess, mulop, wt, areal, bimag), s.newValueOrSfCall2(psess, mulop, wt, aimag, breal))

			if pt != wt {
				xreal = s.newValueOrSfCall1(psess, ssa.OpCvt64Fto32F, pt, xreal)
				ximag = s.newValueOrSfCall1(psess, ssa.OpCvt64Fto32F, pt, ximag)
			}

			return s.newValue2(ssa.OpComplexMake, n.Type, xreal, ximag)
		}

		if n.Type.IsFloat() {
			return s.newValueOrSfCall2(psess, s.ssaOp(psess, n.Op, n.Type), a.Type, a, b)
		}

		return s.newValue2(s.ssaOp(psess, n.Op, n.Type), a.Type, a, b)

	case ODIV:
		a := s.expr(psess, n.Left)
		b := s.expr(psess, n.Right)
		if n.Type.IsComplex() {

			mulop := ssa.OpMul64F
			addop := ssa.OpAdd64F
			subop := ssa.OpSub64F
			divop := ssa.OpDiv64F
			pt := psess.floatForComplex(n.Type)
			wt := psess.types.Types[TFLOAT64]

			areal := s.newValue1(ssa.OpComplexReal, pt, a)
			breal := s.newValue1(ssa.OpComplexReal, pt, b)
			aimag := s.newValue1(ssa.OpComplexImag, pt, a)
			bimag := s.newValue1(ssa.OpComplexImag, pt, b)

			if pt != wt {
				areal = s.newValueOrSfCall1(psess, ssa.OpCvt32Fto64F, wt, areal)
				breal = s.newValueOrSfCall1(psess, ssa.OpCvt32Fto64F, wt, breal)
				aimag = s.newValueOrSfCall1(psess, ssa.OpCvt32Fto64F, wt, aimag)
				bimag = s.newValueOrSfCall1(psess, ssa.OpCvt32Fto64F, wt, bimag)
			}

			denom := s.newValueOrSfCall2(psess, addop, wt, s.newValueOrSfCall2(psess, mulop, wt, breal, breal), s.newValueOrSfCall2(psess, mulop, wt, bimag, bimag))
			xreal := s.newValueOrSfCall2(psess, addop, wt, s.newValueOrSfCall2(psess, mulop, wt, areal, breal), s.newValueOrSfCall2(psess, mulop, wt, aimag, bimag))
			ximag := s.newValueOrSfCall2(psess, subop, wt, s.newValueOrSfCall2(psess, mulop, wt, aimag, breal), s.newValueOrSfCall2(psess, mulop, wt, areal, bimag))

			xreal = s.newValueOrSfCall2(psess, divop, wt, xreal, denom)
			ximag = s.newValueOrSfCall2(psess, divop, wt, ximag, denom)

			if pt != wt {
				xreal = s.newValueOrSfCall1(psess, ssa.OpCvt64Fto32F, pt, xreal)
				ximag = s.newValueOrSfCall1(psess, ssa.OpCvt64Fto32F, pt, ximag)
			}
			return s.newValue2(ssa.OpComplexMake, n.Type, xreal, ximag)
		}
		if n.Type.IsFloat() {
			return s.newValueOrSfCall2(psess, s.ssaOp(psess, n.Op, n.Type), a.Type, a, b)
		}
		return s.intDivide(psess, n, a, b)
	case OMOD:
		a := s.expr(psess, n.Left)
		b := s.expr(psess, n.Right)
		return s.intDivide(psess, n, a, b)
	case OADD, OSUB:
		a := s.expr(psess, n.Left)
		b := s.expr(psess, n.Right)
		if n.Type.IsComplex() {
			pt := psess.floatForComplex(n.Type)
			op := s.ssaOp(psess, n.Op, pt)
			return s.newValue2(ssa.OpComplexMake, n.Type,
				s.newValueOrSfCall2(psess, op, pt, s.newValue1(ssa.OpComplexReal, pt, a), s.newValue1(ssa.OpComplexReal, pt, b)),
				s.newValueOrSfCall2(psess, op, pt, s.newValue1(ssa.OpComplexImag, pt, a), s.newValue1(ssa.OpComplexImag, pt, b)))
		}
		if n.Type.IsFloat() {
			return s.newValueOrSfCall2(psess, s.ssaOp(psess, n.Op, n.Type), a.Type, a, b)
		}
		return s.newValue2(s.ssaOp(psess, n.Op, n.Type), a.Type, a, b)
	case OAND, OOR, OXOR:
		a := s.expr(psess, n.Left)
		b := s.expr(psess, n.Right)
		return s.newValue2(s.ssaOp(psess, n.Op, n.Type), a.Type, a, b)
	case OLSH, ORSH:
		a := s.expr(psess, n.Left)
		b := s.expr(psess, n.Right)
		return s.newValue2(s.ssaShiftOp(psess, n.Op, n.Type, n.Right.Type), a.Type, a, b)
	case OANDAND, OOROR:

		el := s.expr(psess, n.Left)
		s.vars[n] = el

		b := s.endBlock(psess)
		b.Kind = ssa.BlockIf
		b.SetControl(el)

		bRight := s.f.NewBlock(ssa.BlockPlain)
		bResult := s.f.NewBlock(ssa.BlockPlain)
		if n.Op == OANDAND {
			b.AddEdgeTo(bRight)
			b.AddEdgeTo(bResult)
		} else if n.Op == OOROR {
			b.AddEdgeTo(bResult)
			b.AddEdgeTo(bRight)
		}

		s.startBlock(bRight)
		er := s.expr(psess, n.Right)
		s.vars[n] = er

		b = s.endBlock(psess)
		b.AddEdgeTo(bResult)

		s.startBlock(bResult)
		return s.variable(n, psess.types.Types[TBOOL])
	case OCOMPLEX:
		r := s.expr(psess, n.Left)
		i := s.expr(psess, n.Right)
		return s.newValue2(ssa.OpComplexMake, n.Type, r, i)

	case OMINUS:
		a := s.expr(psess, n.Left)
		if n.Type.IsComplex() {
			tp := psess.floatForComplex(n.Type)
			negop := s.ssaOp(psess, n.Op, tp)
			return s.newValue2(ssa.OpComplexMake, n.Type,
				s.newValue1(negop, tp, s.newValue1(ssa.OpComplexReal, tp, a)),
				s.newValue1(negop, tp, s.newValue1(ssa.OpComplexImag, tp, a)))
		}
		return s.newValue1(s.ssaOp(psess, n.Op, n.Type), a.Type, a)
	case ONOT, OCOM:
		a := s.expr(psess, n.Left)
		return s.newValue1(s.ssaOp(psess, n.Op, n.Type), a.Type, a)
	case OIMAG, OREAL:
		a := s.expr(psess, n.Left)
		return s.newValue1(s.ssaOp(psess, n.Op, n.Left.Type), n.Type, a)
	case OPLUS:
		return s.expr(psess, n.Left)

	case OADDR:
		return s.addr(psess, n.Left, n.Bounded())

	case OINDREGSP:
		addr := s.constOffPtrSP(psess, psess.types.NewPtr(n.Type), n.Xoffset)
		return s.load(psess, n.Type, addr)

	case OIND:
		p := s.exprPtr(psess, n.Left, false, n.Pos)
		return s.load(psess, n.Type, p)

	case ODOT:
		if n.Left.Op == OSTRUCTLIT {

			if !psess.isZero(n.Left) {
				psess.
					Fatalf("literal with nonzero value in SSA: %v", n.Left)
			}
			return s.zeroVal(psess, n.Type)
		}

		if islvalue(n) && !s.canSSA(psess, n) {
			p := s.addr(psess, n, false)
			return s.load(psess, n.Type, p)
		}
		v := s.expr(psess, n.Left)
		return s.newValue1I(ssa.OpStructSelect, n.Type, int64(psess.fieldIdx(n)), v)

	case ODOTPTR:
		p := s.exprPtr(psess, n.Left, false, n.Pos)
		p = s.newValue1I(ssa.OpOffPtr, psess.types.NewPtr(n.Type), n.Xoffset, p)
		return s.load(psess, n.Type, p)

	case OINDEX:
		switch {
		case n.Left.Type.IsString():
			if n.Bounded() && psess.Isconst(n.Left, CTSTR) && psess.Isconst(n.Right, CTINT) {

				return s.newValue0I(ssa.OpConst8, psess.types.Types[TUINT8], int64(int8(n.Left.Val().U.(string)[n.Right.Int64(psess)])))
			}
			a := s.expr(psess, n.Left)
			i := s.expr(psess, n.Right)
			i = s.extendIndex(psess, i, psess.panicindex)
			if !n.Bounded() {
				len := s.newValue1(ssa.OpStringLen, psess.types.Types[TINT], a)
				s.boundsCheck(psess, i, len)
			}
			ptrtyp := s.f.Config.Types.BytePtr
			ptr := s.newValue1(ssa.OpStringPtr, ptrtyp, a)
			if psess.Isconst(n.Right, CTINT) {
				ptr = s.newValue1I(ssa.OpOffPtr, ptrtyp, n.Right.Int64(psess), ptr)
			} else {
				ptr = s.newValue2(ssa.OpAddPtr, ptrtyp, ptr, i)
			}
			return s.load(psess, psess.types.Types[TUINT8], ptr)
		case n.Left.Type.IsSlice():
			p := s.addr(psess, n, false)
			return s.load(psess, n.Left.Type.Elem(psess.types), p)
		case n.Left.Type.IsArray():
			if bound := n.Left.Type.NumElem(psess.types); bound <= 1 {

				a := s.expr(psess, n.Left)
				i := s.expr(psess, n.Right)
				if bound == 0 {

					z := s.constInt(psess, psess.types.Types[TINT], 0)
					s.boundsCheck(psess, z, z)

					return s.newValue0(ssa.OpUnknown, n.Type)
				}
				i = s.extendIndex(psess, i, psess.panicindex)
				if !n.Bounded() {
					s.boundsCheck(psess, i, s.constInt(psess, psess.types.Types[TINT], bound))
				}
				return s.newValue1I(ssa.OpArraySelect, n.Type, 0, a)
			}
			p := s.addr(psess, n, false)
			return s.load(psess, n.Left.Type.Elem(psess.types), p)
		default:
			s.Fatalf("bad type for index %v", n.Left.Type)
			return nil
		}

	case OLEN, OCAP:
		switch {
		case n.Left.Type.IsSlice():
			op := ssa.OpSliceLen
			if n.Op == OCAP {
				op = ssa.OpSliceCap
			}
			return s.newValue1(op, psess.types.Types[TINT], s.expr(psess, n.Left))
		case n.Left.Type.IsString():
			return s.newValue1(ssa.OpStringLen, psess.types.Types[TINT], s.expr(psess, n.Left))
		case n.Left.Type.IsMap(), n.Left.Type.IsChan():
			return s.referenceTypeBuiltin(psess, n, s.expr(psess, n.Left))
		default:
			return s.constInt(psess, psess.types.Types[TINT], n.Left.Type.NumElem(psess.types))
		}

	case OSPTR:
		a := s.expr(psess, n.Left)
		if n.Left.Type.IsSlice() {
			return s.newValue1(ssa.OpSlicePtr, n.Type, a)
		} else {
			return s.newValue1(ssa.OpStringPtr, n.Type, a)
		}

	case OITAB:
		a := s.expr(psess, n.Left)
		return s.newValue1(ssa.OpITab, n.Type, a)

	case OIDATA:
		a := s.expr(psess, n.Left)
		return s.newValue1(ssa.OpIData, n.Type, a)

	case OEFACE:
		tab := s.expr(psess, n.Left)
		data := s.expr(psess, n.Right)
		return s.newValue2(ssa.OpIMake, n.Type, tab, data)

	case OSLICE, OSLICEARR, OSLICE3, OSLICE3ARR:
		v := s.expr(psess, n.Left)
		var i, j, k *ssa.Value
		low, high, max := n.SliceBounds(psess)
		if low != nil {
			i = s.extendIndex(psess, s.expr(psess, low), psess.panicslice)
		}
		if high != nil {
			j = s.extendIndex(psess, s.expr(psess, high), psess.panicslice)
		}
		if max != nil {
			k = s.extendIndex(psess, s.expr(psess, max), psess.panicslice)
		}
		p, l, c := s.slice(psess, n.Left.Type, v, i, j, k)
		return s.newValue3(ssa.OpSliceMake, n.Type, p, l, c)

	case OSLICESTR:
		v := s.expr(psess, n.Left)
		var i, j *ssa.Value
		low, high, _ := n.SliceBounds(psess)
		if low != nil {
			i = s.extendIndex(psess, s.expr(psess, low), psess.panicslice)
		}
		if high != nil {
			j = s.extendIndex(psess, s.expr(psess, high), psess.panicslice)
		}
		p, l, _ := s.slice(psess, n.Left.Type, v, i, j, nil)
		return s.newValue2(ssa.OpStringMake, n.Type, p, l)

	case OCALLFUNC:
		if psess.isIntrinsicCall(n) {
			return s.intrinsicCall(psess, n)
		}
		fallthrough

	case OCALLINTER, OCALLMETH:
		a := s.call(psess, n, callNormal)
		return s.load(psess, n.Type, a)

	case OGETG:
		return s.newValue1(ssa.OpGetG, n.Type, s.mem(psess))

	case OAPPEND:
		return s.append(psess, n, false)

	case OSTRUCTLIT, OARRAYLIT:

		if !psess.isZero(n) {
			psess.
				Fatalf("literal with nonzero value in SSA: %v", n)
		}
		return s.zeroVal(psess, n.Type)

	default:
		s.Fatalf("unhandled expr %v", n.Op)
		return nil
	}
}

// append converts an OAPPEND node to SSA.
// If inplace is false, it converts the OAPPEND expression n to an ssa.Value,
// adds it to s, and returns the Value.
// If inplace is true, it writes the result of the OAPPEND expression n
// back to the slice being appended to, and returns nil.
// inplace MUST be set to false if the slice can be SSA'd.
func (s *state) append(psess *PackageSession, n *Node, inplace bool) *ssa.Value {

	et := n.Type.Elem(psess.types)
	pt := psess.types.NewPtr(et)

	sn := n.List.First()

	var slice, addr *ssa.Value
	if inplace {
		addr = s.addr(psess, sn, false)
		slice = s.load(psess, n.Type, addr)
	} else {
		slice = s.expr(psess, sn)
	}

	grow := s.f.NewBlock(ssa.BlockPlain)
	assign := s.f.NewBlock(ssa.BlockPlain)

	nargs := int64(n.List.Len() - 1)
	p := s.newValue1(ssa.OpSlicePtr, pt, slice)
	l := s.newValue1(ssa.OpSliceLen, psess.types.Types[TINT], slice)
	c := s.newValue1(ssa.OpSliceCap, psess.types.Types[TINT], slice)
	nl := s.newValue2(s.ssaOp(psess, OADD, psess.types.Types[TINT]), psess.types.Types[TINT], l, s.constInt(psess, psess.types.Types[TINT], nargs))

	cmp := s.newValue2(s.ssaOp(psess, OGT, psess.types.Types[TINT]), psess.types.Types[TBOOL], nl, c)
	s.vars[&psess.ptrVar] = p

	if !inplace {
		s.vars[&psess.newlenVar] = nl
		s.vars[&psess.capVar] = c
	} else {
		s.vars[&psess.lenVar] = l
	}

	b := s.endBlock(psess)
	b.Kind = ssa.BlockIf
	b.Likely = ssa.BranchUnlikely
	b.SetControl(cmp)
	b.AddEdgeTo(grow)
	b.AddEdgeTo(assign)

	s.startBlock(grow)
	taddr := s.expr(psess, n.Left)
	r := s.rtcall(psess, psess.growslice, true, []*types.Type{pt, psess.types.Types[TINT], psess.types.Types[TINT]}, taddr, p, l, c, nl)

	if inplace {
		if sn.Op == ONAME && sn.Class() != PEXTERN {

			s.vars[&psess.memVar] = s.newValue1A(ssa.OpVarDef, psess.types.TypeMem, sn, s.mem(psess))
		}
		capaddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.IntPtr, int64(psess.array_cap), addr)
		s.store(psess, psess.types.Types[TINT], capaddr, r[2])
		s.store(psess, pt, addr, r[0])

		s.vars[&psess.ptrVar] = s.load(psess, pt, addr)
		s.vars[&psess.lenVar] = r[1]
	} else {
		s.vars[&psess.ptrVar] = r[0]
		s.vars[&psess.newlenVar] = s.newValue2(s.ssaOp(psess, OADD, psess.types.Types[TINT]), psess.types.Types[TINT], r[1], s.constInt(psess, psess.types.Types[TINT], nargs))
		s.vars[&psess.capVar] = r[2]
	}

	b = s.endBlock(psess)
	b.AddEdgeTo(assign)

	s.startBlock(assign)

	if inplace {
		l = s.variable(&psess.lenVar, psess.types.Types[TINT])
		nl = s.newValue2(s.ssaOp(psess, OADD, psess.types.Types[TINT]), psess.types.Types[TINT], l, s.constInt(psess, psess.types.Types[TINT], nargs))
		lenaddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.IntPtr, int64(psess.array_nel), addr)
		s.store(psess, psess.types.Types[TINT], lenaddr, nl)
	}

	// Evaluate args
	type argRec struct {
		// if store is true, we're appending the value v.  If false, we're appending the
		// value at *v.
		v     *ssa.Value
		store bool
	}
	args := make([]argRec, 0, nargs)
	for _, n := range n.List.Slice()[1:] {
		if psess.canSSAType(n.Type) {
			args = append(args, argRec{v: s.expr(psess, n), store: true})
		} else {
			v := s.addr(psess, n, false)
			args = append(args, argRec{v: v})
		}
	}

	p = s.variable(&psess.ptrVar, pt)
	if !inplace {
		nl = s.variable(&psess.newlenVar, psess.types.Types[TINT])
		c = s.variable(&psess.capVar, psess.types.Types[TINT])
	}
	p2 := s.newValue2(ssa.OpPtrIndex, pt, p, l)
	for i, arg := range args {
		addr := s.newValue2(ssa.OpPtrIndex, pt, p2, s.constInt(psess, psess.types.Types[TINT], int64(i)))
		if arg.store {
			s.storeType(psess, et, addr, arg.v, 0, true)
		} else {
			s.move(psess, et, addr, arg.v)
		}
	}

	delete(s.vars, &psess.ptrVar)
	if inplace {
		delete(s.vars, &psess.lenVar)
		return nil
	}
	delete(s.vars, &psess.newlenVar)
	delete(s.vars, &psess.capVar)

	return s.newValue3(ssa.OpSliceMake, n.Type, p, nl, c)
}

// condBranch evaluates the boolean expression cond and branches to yes
// if cond is true and no if cond is false.
// This function is intended to handle && and || better than just calling
// s.expr(cond) and branching on the result.
func (s *state) condBranch(psess *PackageSession, cond *Node, yes, no *ssa.Block, likely int8) {
	switch cond.Op {
	case OANDAND:
		mid := s.f.NewBlock(ssa.BlockPlain)
		s.stmtList(psess, cond.Ninit)
		s.condBranch(psess, cond.Left, mid, no, max8(likely, 0))
		s.startBlock(mid)
		s.condBranch(psess, cond.Right, yes, no, likely)
		return

	case OOROR:
		mid := s.f.NewBlock(ssa.BlockPlain)
		s.stmtList(psess, cond.Ninit)
		s.condBranch(psess, cond.Left, yes, mid, min8(likely, 0))
		s.startBlock(mid)
		s.condBranch(psess, cond.Right, yes, no, likely)
		return

	case ONOT:
		s.stmtList(psess, cond.Ninit)
		s.condBranch(psess, cond.Left, no, yes, -likely)
		return
	}
	c := s.expr(psess, cond)
	b := s.endBlock(psess)
	b.Kind = ssa.BlockIf
	b.SetControl(c)
	b.Likely = ssa.BranchPrediction(likely)
	b.AddEdgeTo(yes)
	b.AddEdgeTo(no)
}

type skipMask uint8

const (
	skipPtr skipMask = 1 << iota
	skipLen
	skipCap
)

// assign does left = right.
// Right has already been evaluated to ssa, left has not.
// If deref is true, then we do left = *right instead (and right has already been nil-checked).
// If deref is true and right == nil, just do left = 0.
// skip indicates assignments (at the top level) that can be avoided.
func (s *state) assign(psess *PackageSession, left *Node, right *ssa.Value, deref bool, skip skipMask) {
	if left.Op == ONAME && left.isBlank() {
		return
	}
	t := left.Type
	psess.
		dowidth(t)
	if s.canSSA(psess, left) {
		if deref {
			s.Fatalf("can SSA LHS %v but not RHS %s", left, right)
		}
		if left.Op == ODOT {

			t := left.Left.Type
			nf := t.NumFields(psess.types)
			idx := psess.fieldIdx(left)

			old := s.expr(psess, left.Left)

			new := s.newValue0(ssa.StructMakeOp(t.NumFields(psess.types)), t)

			for i := 0; i < nf; i++ {
				if i == idx {
					new.AddArg(right)
				} else {
					new.AddArg(s.newValue1I(ssa.OpStructSelect, t.FieldType(psess.types, i), int64(i), old))
				}
			}

			s.assign(psess, left.Left, new, false, 0)

			return
		}
		if left.Op == OINDEX && left.Left.Type.IsArray() {

			t := left.Left.Type
			n := t.NumElem(psess.types)

			i := s.expr(psess, left.Right)
			if n == 0 {

				z := s.constInt(psess, psess.types.Types[TINT], 0)
				s.boundsCheck(psess, z, z)
				return
			}
			if n != 1 {
				s.Fatalf("assigning to non-1-length array")
			}

			i = s.extendIndex(psess, i, psess.panicindex)
			s.boundsCheck(psess, i, s.constInt(psess, psess.types.Types[TINT], 1))
			v := s.newValue1(ssa.OpArrayMake1, t, right)
			s.assign(psess, left.Left, v, false, 0)
			return
		}

		s.vars[left] = right
		s.addNamedValue(left, right)
		return
	}

	addr := s.addr(psess, left, false)
	if left.Op == ONAME && left.Class() != PEXTERN && skip == 0 {
		s.vars[&psess.memVar] = s.newValue1Apos(ssa.OpVarDef, psess.types.TypeMem, left, s.mem(psess), !left.IsAutoTmp())
	}
	if psess.isReflectHeaderDataField(left) {

		t = psess.types.Types[TUNSAFEPTR]
	}
	if deref {

		if right == nil {
			s.zero(psess, t, addr)
		} else {
			s.move(psess, t, addr, right)
		}
		return
	}

	s.storeType(psess, t, addr, right, skip, !left.IsAutoTmp())
}

// zeroVal returns the zero value for type t.
func (s *state) zeroVal(psess *PackageSession, t *types.Type) *ssa.Value {
	switch {
	case t.IsInteger():
		switch t.Size(psess.types) {
		case 1:
			return s.constInt8(psess, t, 0)
		case 2:
			return s.constInt16(psess, t, 0)
		case 4:
			return s.constInt32(psess, t, 0)
		case 8:
			return s.constInt64(psess, t, 0)
		default:
			s.Fatalf("bad sized integer type %v", t)
		}
	case t.IsFloat():
		switch t.Size(psess.types) {
		case 4:
			return s.constFloat32(psess, t, 0)
		case 8:
			return s.constFloat64(psess, t, 0)
		default:
			s.Fatalf("bad sized float type %v", t)
		}
	case t.IsComplex():
		switch t.Size(psess.types) {
		case 8:
			z := s.constFloat32(psess, psess.types.Types[TFLOAT32], 0)
			return s.entryNewValue2(psess, ssa.OpComplexMake, t, z, z)
		case 16:
			z := s.constFloat64(psess, psess.types.Types[TFLOAT64], 0)
			return s.entryNewValue2(psess, ssa.OpComplexMake, t, z, z)
		default:
			s.Fatalf("bad sized complex type %v", t)
		}

	case t.IsString():
		return s.constEmptyString(psess, t)
	case t.IsPtrShaped():
		return s.constNil(psess, t)
	case t.IsBoolean():
		return s.constBool(psess, false)
	case t.IsInterface():
		return s.constInterface(psess, t)
	case t.IsSlice():
		return s.constSlice(psess, t)
	case t.IsStruct():
		n := t.NumFields(psess.types)
		v := s.entryNewValue0(psess, ssa.StructMakeOp(t.NumFields(psess.types)), t)
		for i := 0; i < n; i++ {
			v.AddArg(s.zeroVal(psess, t.FieldType(psess.types, i)))
		}
		return v
	case t.IsArray():
		switch t.NumElem(psess.types) {
		case 0:
			return s.entryNewValue0(psess, ssa.OpArrayMake0, t)
		case 1:
			return s.entryNewValue1(psess, ssa.OpArrayMake1, t, s.zeroVal(psess, t.Elem(psess.types)))
		}
	}
	s.Fatalf("zero for type %v not implemented", t)
	return nil
}

type callKind int8

const (
	callNormal callKind = iota
	callDefer
	callGo
)

type sfRtCallDef struct {
	rtfn  *obj.LSym
	rtype types.EType
}

func (psess *PackageSession) softfloatInit() {
	psess.
		softFloatOps = map[ssa.Op]sfRtCallDef{
		ssa.OpAdd32F: sfRtCallDef{psess.sysfunc("fadd32"), TFLOAT32},
		ssa.OpAdd64F: sfRtCallDef{psess.sysfunc("fadd64"), TFLOAT64},
		ssa.OpSub32F: sfRtCallDef{psess.sysfunc("fadd32"), TFLOAT32},
		ssa.OpSub64F: sfRtCallDef{psess.sysfunc("fadd64"), TFLOAT64},
		ssa.OpMul32F: sfRtCallDef{psess.sysfunc("fmul32"), TFLOAT32},
		ssa.OpMul64F: sfRtCallDef{psess.sysfunc("fmul64"), TFLOAT64},
		ssa.OpDiv32F: sfRtCallDef{psess.sysfunc("fdiv32"), TFLOAT32},
		ssa.OpDiv64F: sfRtCallDef{psess.sysfunc("fdiv64"), TFLOAT64},

		ssa.OpEq64F:      sfRtCallDef{psess.sysfunc("feq64"), TBOOL},
		ssa.OpEq32F:      sfRtCallDef{psess.sysfunc("feq32"), TBOOL},
		ssa.OpNeq64F:     sfRtCallDef{psess.sysfunc("feq64"), TBOOL},
		ssa.OpNeq32F:     sfRtCallDef{psess.sysfunc("feq32"), TBOOL},
		ssa.OpLess64F:    sfRtCallDef{psess.sysfunc("fgt64"), TBOOL},
		ssa.OpLess32F:    sfRtCallDef{psess.sysfunc("fgt32"), TBOOL},
		ssa.OpGreater64F: sfRtCallDef{psess.sysfunc("fgt64"), TBOOL},
		ssa.OpGreater32F: sfRtCallDef{psess.sysfunc("fgt32"), TBOOL},
		ssa.OpLeq64F:     sfRtCallDef{psess.sysfunc("fge64"), TBOOL},
		ssa.OpLeq32F:     sfRtCallDef{psess.sysfunc("fge32"), TBOOL},
		ssa.OpGeq64F:     sfRtCallDef{psess.sysfunc("fge64"), TBOOL},
		ssa.OpGeq32F:     sfRtCallDef{psess.sysfunc("fge32"), TBOOL},

		ssa.OpCvt32to32F:  sfRtCallDef{psess.sysfunc("fint32to32"), TFLOAT32},
		ssa.OpCvt32Fto32:  sfRtCallDef{psess.sysfunc("f32toint32"), TINT32},
		ssa.OpCvt64to32F:  sfRtCallDef{psess.sysfunc("fint64to32"), TFLOAT32},
		ssa.OpCvt32Fto64:  sfRtCallDef{psess.sysfunc("f32toint64"), TINT64},
		ssa.OpCvt64Uto32F: sfRtCallDef{psess.sysfunc("fuint64to32"), TFLOAT32},
		ssa.OpCvt32Fto64U: sfRtCallDef{psess.sysfunc("f32touint64"), TUINT64},
		ssa.OpCvt32to64F:  sfRtCallDef{psess.sysfunc("fint32to64"), TFLOAT64},
		ssa.OpCvt64Fto32:  sfRtCallDef{psess.sysfunc("f64toint32"), TINT32},
		ssa.OpCvt64to64F:  sfRtCallDef{psess.sysfunc("fint64to64"), TFLOAT64},
		ssa.OpCvt64Fto64:  sfRtCallDef{psess.sysfunc("f64toint64"), TINT64},
		ssa.OpCvt64Uto64F: sfRtCallDef{psess.sysfunc("fuint64to64"), TFLOAT64},
		ssa.OpCvt64Fto64U: sfRtCallDef{psess.sysfunc("f64touint64"), TUINT64},
		ssa.OpCvt32Fto64F: sfRtCallDef{psess.sysfunc("f32to64"), TFLOAT64},
		ssa.OpCvt64Fto32F: sfRtCallDef{psess.sysfunc("f64to32"), TFLOAT32},
	}
}

// TODO: do not emit sfcall if operation can be optimized to constant in later
// opt phase
func (s *state) sfcall(psess *PackageSession, op ssa.Op, args ...*ssa.Value) (*ssa.Value, bool) {
	if callDef, ok := psess.softFloatOps[op]; ok {
		switch op {
		case ssa.OpLess32F,
			ssa.OpLess64F,
			ssa.OpLeq32F,
			ssa.OpLeq64F:
			args[0], args[1] = args[1], args[0]
		case ssa.OpSub32F,
			ssa.OpSub64F:
			args[1] = s.newValue1(s.ssaOp(psess, OMINUS, psess.types.Types[callDef.rtype]), args[1].Type, args[1])
		}

		result := s.rtcall(psess, callDef.rtfn, true, []*types.Type{psess.types.Types[callDef.rtype]}, args...)[0]
		if op == ssa.OpNeq32F || op == ssa.OpNeq64F {
			result = s.newValue1(ssa.OpNot, result.Type, result)
		}
		return result, true
	}
	return nil, false
}

// An intrinsicBuilder converts a call node n into an ssa value that
// implements that call as an intrinsic. args is a list of arguments to the func.
type intrinsicBuilder func(s *state, n *Node, args []*ssa.Value) *ssa.Value

type intrinsicKey struct {
	arch *sys.Arch
	pkg  string
	fn   string
}

func (psess *PackageSession) init() {
	psess.
		intrinsics = map[intrinsicKey]intrinsicBuilder{}

	var all []*sys.Arch
	var p4 []*sys.Arch
	var p8 []*sys.Arch
	for _, a := range psess.sys.Archs {
		all = append(all, a)
		if a.PtrSize == 4 {
			p4 = append(p4, a)
		} else {
			p8 = append(p8, a)
		}
	}

	add := func(pkg, fn string, b intrinsicBuilder, archs ...*sys.Arch) {
		for _, a := range archs {
			psess.
				intrinsics[intrinsicKey{a, pkg, fn}] = b
		}
	}

	addF := func(pkg, fn string, b intrinsicBuilder, archFamilies ...sys.ArchFamily) {
		m := 0
		for _, f := range archFamilies {
			if f >= 32 {
				panic("too many architecture families")
			}
			m |= 1 << uint(f)
		}
		for _, a := range all {
			if m>>uint(a.Family)&1 != 0 {
				psess.
					intrinsics[intrinsicKey{a, pkg, fn}] = b
			}
		}
	}

	alias := func(pkg, fn, pkg2, fn2 string, archs ...*sys.Arch) {
		for _, a := range archs {
			if b, ok := psess.intrinsics[intrinsicKey{a, pkg2, fn2}]; ok {
				psess.
					intrinsics[intrinsicKey{a, pkg, fn}] = b
			}
		}
	}

	if !psess.instrumenting {
		add("runtime", "slicebytetostringtmp",
			func(s *state, n *Node, args []*ssa.Value) *ssa.Value {

				slice := args[0]
				ptr := s.newValue1(ssa.OpSlicePtr, s.f.Config.Types.BytePtr, slice)
				len := s.newValue1(ssa.OpSliceLen, psess.types.Types[TINT], slice)
				return s.newValue2(ssa.OpStringMake, n.Type, ptr, len)
			},
			all...)
	}
	add("runtime", "KeepAlive",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			data := s.newValue1(ssa.OpIData, s.f.Config.Types.BytePtr, args[0])
			s.vars[&psess.memVar] = s.newValue2(ssa.OpKeepAlive, psess.types.TypeMem, data, s.mem(psess))
			return nil
		},
		all...)
	add("runtime", "getclosureptr",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue0(ssa.OpGetClosurePtr, s.f.Config.Types.Uintptr)
		},
		all...)

	add("runtime", "getcallerpc",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue0(ssa.OpGetCallerPC, s.f.Config.Types.Uintptr)
		},
		all...)

	add("runtime", "getcallersp",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue0(ssa.OpGetCallerSP, s.f.Config.Types.Uintptr)
		},
		all...)

	addF("runtime/internal/sys", "Ctz32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz32, psess.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("runtime/internal/sys", "Ctz64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz64, psess.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("runtime/internal/sys", "Bswap32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBswap32, psess.types.Types[TUINT32], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X)
	addF("runtime/internal/sys", "Bswap64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBswap64, psess.types.Types[TUINT64], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X)

	addF("runtime/internal/atomic", "Load",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue2(ssa.OpAtomicLoad32, types.NewTuple(psess.types.Types[TUINT32], psess.types.TypeMem), args[0], s.mem(psess))
			s.vars[&psess.memVar] = s.newValue1(ssa.OpSelect1, psess.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, psess.types.Types[TUINT32], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Load64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue2(ssa.OpAtomicLoad64, types.NewTuple(psess.types.Types[TUINT64], psess.types.TypeMem), args[0], s.mem(psess))
			s.vars[&psess.memVar] = s.newValue1(ssa.OpSelect1, psess.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, psess.types.Types[TUINT64], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Loadp",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue2(ssa.OpAtomicLoadPtr, types.NewTuple(s.f.Config.Types.BytePtr, psess.types.TypeMem), args[0], s.mem(psess))
			s.vars[&psess.memVar] = s.newValue1(ssa.OpSelect1, psess.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, s.f.Config.Types.BytePtr, v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)

	addF("runtime/internal/atomic", "Store",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			s.vars[&psess.memVar] = s.newValue3(ssa.OpAtomicStore32, psess.types.TypeMem, args[0], args[1], s.mem(psess))
			return nil
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Store64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			s.vars[&psess.memVar] = s.newValue3(ssa.OpAtomicStore64, psess.types.TypeMem, args[0], args[1], s.mem(psess))
			return nil
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "StorepNoWB",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			s.vars[&psess.memVar] = s.newValue3(ssa.OpAtomicStorePtrNoWB, psess.types.TypeMem, args[0], args[1], s.mem(psess))
			return nil
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64)

	addF("runtime/internal/atomic", "Xchg",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue3(ssa.OpAtomicExchange32, types.NewTuple(psess.types.Types[TUINT32], psess.types.TypeMem), args[0], args[1], s.mem(psess))
			s.vars[&psess.memVar] = s.newValue1(ssa.OpSelect1, psess.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, psess.types.Types[TUINT32], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Xchg64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue3(ssa.OpAtomicExchange64, types.NewTuple(psess.types.Types[TUINT64], psess.types.TypeMem), args[0], args[1], s.mem(psess))
			s.vars[&psess.memVar] = s.newValue1(ssa.OpSelect1, psess.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, psess.types.Types[TUINT64], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS64, sys.PPC64)

	addF("runtime/internal/atomic", "Xadd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue3(ssa.OpAtomicAdd32, types.NewTuple(psess.types.Types[TUINT32], psess.types.TypeMem), args[0], args[1], s.mem(psess))
			s.vars[&psess.memVar] = s.newValue1(ssa.OpSelect1, psess.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, psess.types.Types[TUINT32], v)
		},
		sys.AMD64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Xadd64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue3(ssa.OpAtomicAdd64, types.NewTuple(psess.types.Types[TUINT64], psess.types.TypeMem), args[0], args[1], s.mem(psess))
			s.vars[&psess.memVar] = s.newValue1(ssa.OpSelect1, psess.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, psess.types.Types[TUINT64], v)
		},
		sys.AMD64, sys.S390X, sys.MIPS64, sys.PPC64)

	makeXaddARM64 := func(op0 ssa.Op, op1 ssa.Op, ty types.EType) func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return func(s *state, n *Node, args []*ssa.Value) *ssa.Value {

			addr := s.entryNewValue1A(psess, ssa.OpAddr, psess.types.Types[TBOOL].PtrTo(psess.types), psess.arm64SupportAtomics, s.sb)
			v := s.load(psess, psess.types.Types[TBOOL], addr)
			b := s.endBlock(psess)
			b.Kind = ssa.BlockIf
			b.SetControl(v)
			bTrue := s.f.NewBlock(ssa.BlockPlain)
			bFalse := s.f.NewBlock(ssa.BlockPlain)
			bEnd := s.f.NewBlock(ssa.BlockPlain)
			b.AddEdgeTo(bTrue)
			b.AddEdgeTo(bFalse)
			b.Likely = ssa.BranchUnlikely

			s.startBlock(bTrue)
			v0 := s.newValue3(op1, types.NewTuple(psess.types.Types[ty], psess.types.TypeMem), args[0], args[1], s.mem(psess))
			s.vars[&psess.memVar] = s.newValue1(ssa.OpSelect1, psess.types.TypeMem, v0)
			s.vars[n] = s.newValue1(ssa.OpSelect0, psess.types.Types[ty], v0)
			s.endBlock(psess).AddEdgeTo(bEnd)

			s.startBlock(bFalse)
			v1 := s.newValue3(op0, types.NewTuple(psess.types.Types[ty], psess.types.TypeMem), args[0], args[1], s.mem(psess))
			s.vars[&psess.memVar] = s.newValue1(ssa.OpSelect1, psess.types.TypeMem, v1)
			s.vars[n] = s.newValue1(ssa.OpSelect0, psess.types.Types[ty], v1)
			s.endBlock(psess).AddEdgeTo(bEnd)

			s.startBlock(bEnd)
			return s.variable(n, psess.types.Types[ty])
		}
	}

	addF("runtime/internal/atomic", "Xadd",
		makeXaddARM64(ssa.OpAtomicAdd32, ssa.OpAtomicAdd32Variant, TUINT32),
		sys.ARM64)
	addF("runtime/internal/atomic", "Xadd64",
		makeXaddARM64(ssa.OpAtomicAdd64, ssa.OpAtomicAdd64Variant, TUINT64),
		sys.ARM64)

	addF("runtime/internal/atomic", "Cas",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue4(ssa.OpAtomicCompareAndSwap32, types.NewTuple(psess.types.Types[TBOOL], psess.types.TypeMem), args[0], args[1], args[2], s.mem(psess))
			s.vars[&psess.memVar] = s.newValue1(ssa.OpSelect1, psess.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, psess.types.Types[TBOOL], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Cas64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue4(ssa.OpAtomicCompareAndSwap64, types.NewTuple(psess.types.Types[TBOOL], psess.types.TypeMem), args[0], args[1], args[2], s.mem(psess))
			s.vars[&psess.memVar] = s.newValue1(ssa.OpSelect1, psess.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, psess.types.Types[TBOOL], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS64, sys.PPC64)

	addF("runtime/internal/atomic", "And8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			s.vars[&psess.memVar] = s.newValue3(ssa.OpAtomicAnd8, psess.types.TypeMem, args[0], args[1], s.mem(psess))
			return nil
		},
		sys.AMD64, sys.ARM64, sys.MIPS, sys.PPC64)
	addF("runtime/internal/atomic", "Or8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			s.vars[&psess.memVar] = s.newValue3(ssa.OpAtomicOr8, psess.types.TypeMem, args[0], args[1], s.mem(psess))
			return nil
		},
		sys.AMD64, sys.ARM64, sys.MIPS, sys.PPC64)

	alias("runtime/internal/atomic", "Loadint64", "runtime/internal/atomic", "Load64", all...)
	alias("runtime/internal/atomic", "Xaddint64", "runtime/internal/atomic", "Xadd64", all...)
	alias("runtime/internal/atomic", "Loaduint", "runtime/internal/atomic", "Load", p4...)
	alias("runtime/internal/atomic", "Loaduint", "runtime/internal/atomic", "Load64", p8...)
	alias("runtime/internal/atomic", "Loaduintptr", "runtime/internal/atomic", "Load", p4...)
	alias("runtime/internal/atomic", "Loaduintptr", "runtime/internal/atomic", "Load64", p8...)
	alias("runtime/internal/atomic", "Storeuintptr", "runtime/internal/atomic", "Store", p4...)
	alias("runtime/internal/atomic", "Storeuintptr", "runtime/internal/atomic", "Store64", p8...)
	alias("runtime/internal/atomic", "Xchguintptr", "runtime/internal/atomic", "Xchg", p4...)
	alias("runtime/internal/atomic", "Xchguintptr", "runtime/internal/atomic", "Xchg64", p8...)
	alias("runtime/internal/atomic", "Xadduintptr", "runtime/internal/atomic", "Xadd", p4...)
	alias("runtime/internal/atomic", "Xadduintptr", "runtime/internal/atomic", "Xadd64", p8...)
	alias("runtime/internal/atomic", "Casuintptr", "runtime/internal/atomic", "Cas", p4...)
	alias("runtime/internal/atomic", "Casuintptr", "runtime/internal/atomic", "Cas64", p8...)
	alias("runtime/internal/atomic", "Casp1", "runtime/internal/atomic", "Cas", p4...)
	alias("runtime/internal/atomic", "Casp1", "runtime/internal/atomic", "Cas64", p8...)

	alias("runtime/internal/sys", "Ctz8", "math/bits", "TrailingZeros8", all...)

	addF("math", "Sqrt",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpSqrt, psess.types.Types[TFLOAT64], args[0])
		},
		sys.I386, sys.AMD64, sys.ARM, sys.ARM64, sys.MIPS, sys.MIPS64, sys.PPC64, sys.S390X)
	addF("math", "Trunc",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpTrunc, psess.types.Types[TFLOAT64], args[0])
		},
		sys.ARM64, sys.PPC64, sys.S390X)
	addF("math", "Ceil",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCeil, psess.types.Types[TFLOAT64], args[0])
		},
		sys.ARM64, sys.PPC64, sys.S390X)
	addF("math", "Floor",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpFloor, psess.types.Types[TFLOAT64], args[0])
		},
		sys.ARM64, sys.PPC64, sys.S390X)
	addF("math", "Round",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpRound, psess.types.Types[TFLOAT64], args[0])
		},
		sys.ARM64, sys.PPC64, sys.S390X)
	addF("math", "RoundToEven",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpRoundToEven, psess.types.Types[TFLOAT64], args[0])
		},
		sys.S390X)
	addF("math", "Abs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpAbs, psess.types.Types[TFLOAT64], args[0])
		},
		sys.PPC64)
	addF("math", "Copysign",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue2(ssa.OpCopysign, psess.types.Types[TFLOAT64], args[0], args[1])
		},
		sys.PPC64)

	makeRoundAMD64 := func(op ssa.Op) func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			addr := s.entryNewValue1A(psess, ssa.OpAddr, psess.types.Types[TBOOL].PtrTo(psess.types), psess.supportSSE41, s.sb)
			v := s.load(psess, psess.types.Types[TBOOL], addr)
			b := s.endBlock(psess)
			b.Kind = ssa.BlockIf
			b.SetControl(v)
			bTrue := s.f.NewBlock(ssa.BlockPlain)
			bFalse := s.f.NewBlock(ssa.BlockPlain)
			bEnd := s.f.NewBlock(ssa.BlockPlain)
			b.AddEdgeTo(bTrue)
			b.AddEdgeTo(bFalse)
			b.Likely = ssa.BranchLikely

			s.startBlock(bTrue)
			s.vars[n] = s.newValue1(op, psess.types.Types[TFLOAT64], args[0])
			s.endBlock(psess).AddEdgeTo(bEnd)

			s.startBlock(bFalse)
			a := s.call(psess, n, callNormal)
			s.vars[n] = s.load(psess, psess.types.Types[TFLOAT64], a)
			s.endBlock(psess).AddEdgeTo(bEnd)

			s.startBlock(bEnd)
			return s.variable(n, psess.types.Types[TFLOAT64])
		}
	}
	addF("math", "RoundToEven",
		makeRoundAMD64(ssa.OpRoundToEven),
		sys.AMD64)
	addF("math", "Floor",
		makeRoundAMD64(ssa.OpFloor),
		sys.AMD64)
	addF("math", "Ceil",
		makeRoundAMD64(ssa.OpCeil),
		sys.AMD64)
	addF("math", "Trunc",
		makeRoundAMD64(ssa.OpTrunc),
		sys.AMD64)

	addF("math/bits", "TrailingZeros64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz64, psess.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "TrailingZeros32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz32, psess.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "TrailingZeros16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			x := s.newValue1(ssa.OpZeroExt16to32, psess.types.Types[TUINT32], args[0])
			c := s.constInt32(psess, psess.types.Types[TUINT32], 1<<16)
			y := s.newValue2(ssa.OpOr32, psess.types.Types[TUINT32], x, c)
			return s.newValue1(ssa.OpCtz32, psess.types.Types[TINT], y)
		},
		sys.ARM, sys.MIPS)
	addF("math/bits", "TrailingZeros16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz16, psess.types.Types[TINT], args[0])
		},
		sys.AMD64)
	addF("math/bits", "TrailingZeros16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			x := s.newValue1(ssa.OpZeroExt16to64, psess.types.Types[TUINT64], args[0])
			c := s.constInt64(psess, psess.types.Types[TUINT64], 1<<16)
			y := s.newValue2(ssa.OpOr64, psess.types.Types[TUINT64], x, c)
			return s.newValue1(ssa.OpCtz64, psess.types.Types[TINT], y)
		},
		sys.ARM64, sys.S390X)
	addF("math/bits", "TrailingZeros8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			x := s.newValue1(ssa.OpZeroExt8to32, psess.types.Types[TUINT32], args[0])
			c := s.constInt32(psess, psess.types.Types[TUINT32], 1<<8)
			y := s.newValue2(ssa.OpOr32, psess.types.Types[TUINT32], x, c)
			return s.newValue1(ssa.OpCtz32, psess.types.Types[TINT], y)
		},
		sys.ARM, sys.MIPS)
	addF("math/bits", "TrailingZeros8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz8, psess.types.Types[TINT], args[0])
		},
		sys.AMD64)
	addF("math/bits", "TrailingZeros8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			x := s.newValue1(ssa.OpZeroExt8to64, psess.types.Types[TUINT64], args[0])
			c := s.constInt64(psess, psess.types.Types[TUINT64], 1<<8)
			y := s.newValue2(ssa.OpOr64, psess.types.Types[TUINT64], x, c)
			return s.newValue1(ssa.OpCtz64, psess.types.Types[TINT], y)
		},
		sys.ARM64, sys.S390X)
	alias("math/bits", "ReverseBytes64", "runtime/internal/sys", "Bswap64", all...)
	alias("math/bits", "ReverseBytes32", "runtime/internal/sys", "Bswap32", all...)

	addF("math/bits", "Len64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitLen64, psess.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "Len32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitLen32, psess.types.Types[TINT], args[0])
		},
		sys.AMD64)
	addF("math/bits", "Len32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			if s.config.PtrSize == 4 {
				return s.newValue1(ssa.OpBitLen32, psess.types.Types[TINT], args[0])
			}
			x := s.newValue1(ssa.OpZeroExt32to64, psess.types.Types[TUINT64], args[0])
			return s.newValue1(ssa.OpBitLen64, psess.types.Types[TINT], x)
		},
		sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "Len16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			if s.config.PtrSize == 4 {
				x := s.newValue1(ssa.OpZeroExt16to32, psess.types.Types[TUINT32], args[0])
				return s.newValue1(ssa.OpBitLen32, psess.types.Types[TINT], x)
			}
			x := s.newValue1(ssa.OpZeroExt16to64, psess.types.Types[TUINT64], args[0])
			return s.newValue1(ssa.OpBitLen64, psess.types.Types[TINT], x)
		},
		sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "Len16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitLen16, psess.types.Types[TINT], args[0])
		},
		sys.AMD64)
	addF("math/bits", "Len8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			if s.config.PtrSize == 4 {
				x := s.newValue1(ssa.OpZeroExt8to32, psess.types.Types[TUINT32], args[0])
				return s.newValue1(ssa.OpBitLen32, psess.types.Types[TINT], x)
			}
			x := s.newValue1(ssa.OpZeroExt8to64, psess.types.Types[TUINT64], args[0])
			return s.newValue1(ssa.OpBitLen64, psess.types.Types[TINT], x)
		},
		sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "Len8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitLen8, psess.types.Types[TINT], args[0])
		},
		sys.AMD64)
	addF("math/bits", "Len",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			if s.config.PtrSize == 4 {
				return s.newValue1(ssa.OpBitLen32, psess.types.Types[TINT], args[0])
			}
			return s.newValue1(ssa.OpBitLen64, psess.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)

	addF("math/bits", "Reverse64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitRev64, psess.types.Types[TINT], args[0])
		},
		sys.ARM64)
	addF("math/bits", "Reverse32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitRev32, psess.types.Types[TINT], args[0])
		},
		sys.ARM64)
	addF("math/bits", "Reverse16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitRev16, psess.types.Types[TINT], args[0])
		},
		sys.ARM64)
	addF("math/bits", "Reverse8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitRev8, psess.types.Types[TINT], args[0])
		},
		sys.ARM64)
	addF("math/bits", "Reverse",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			if s.config.PtrSize == 4 {
				return s.newValue1(ssa.OpBitRev32, psess.types.Types[TINT], args[0])
			}
			return s.newValue1(ssa.OpBitRev64, psess.types.Types[TINT], args[0])
		},
		sys.ARM64)
	makeOnesCountAMD64 := func(op64 ssa.Op, op32 ssa.Op) func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			addr := s.entryNewValue1A(psess, ssa.OpAddr, psess.types.Types[TBOOL].PtrTo(psess.types), psess.supportPopcnt, s.sb)
			v := s.load(psess, psess.types.Types[TBOOL], addr)
			b := s.endBlock(psess)
			b.Kind = ssa.BlockIf
			b.SetControl(v)
			bTrue := s.f.NewBlock(ssa.BlockPlain)
			bFalse := s.f.NewBlock(ssa.BlockPlain)
			bEnd := s.f.NewBlock(ssa.BlockPlain)
			b.AddEdgeTo(bTrue)
			b.AddEdgeTo(bFalse)
			b.Likely = ssa.BranchLikely

			s.startBlock(bTrue)
			op := op64
			if s.config.PtrSize == 4 {
				op = op32
			}
			s.vars[n] = s.newValue1(op, psess.types.Types[TINT], args[0])
			s.endBlock(psess).AddEdgeTo(bEnd)

			s.startBlock(bFalse)
			a := s.call(psess, n, callNormal)
			s.vars[n] = s.load(psess, psess.types.Types[TINT], a)
			s.endBlock(psess).AddEdgeTo(bEnd)

			s.startBlock(bEnd)
			return s.variable(n, psess.types.Types[TINT])
		}
	}
	addF("math/bits", "OnesCount64",
		makeOnesCountAMD64(ssa.OpPopCount64, ssa.OpPopCount64),
		sys.AMD64)
	addF("math/bits", "OnesCount64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpPopCount64, psess.types.Types[TINT], args[0])
		},
		sys.PPC64, sys.ARM64)
	addF("math/bits", "OnesCount32",
		makeOnesCountAMD64(ssa.OpPopCount32, ssa.OpPopCount32),
		sys.AMD64)
	addF("math/bits", "OnesCount32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpPopCount32, psess.types.Types[TINT], args[0])
		},
		sys.PPC64, sys.ARM64)
	addF("math/bits", "OnesCount16",
		makeOnesCountAMD64(ssa.OpPopCount16, ssa.OpPopCount16),
		sys.AMD64)
	addF("math/bits", "OnesCount16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpPopCount16, psess.types.Types[TINT], args[0])
		},
		sys.ARM64)

	addF("math/bits", "OnesCount",
		makeOnesCountAMD64(ssa.OpPopCount64, ssa.OpPopCount32),
		sys.AMD64)

	alias("sync/atomic", "LoadInt32", "runtime/internal/atomic", "Load", all...)
	alias("sync/atomic", "LoadInt64", "runtime/internal/atomic", "Load64", all...)
	alias("sync/atomic", "LoadPointer", "runtime/internal/atomic", "Loadp", all...)
	alias("sync/atomic", "LoadUint32", "runtime/internal/atomic", "Load", all...)
	alias("sync/atomic", "LoadUint64", "runtime/internal/atomic", "Load64", all...)
	alias("sync/atomic", "LoadUintptr", "runtime/internal/atomic", "Load", p4...)
	alias("sync/atomic", "LoadUintptr", "runtime/internal/atomic", "Load64", p8...)

	alias("sync/atomic", "StoreInt32", "runtime/internal/atomic", "Store", all...)
	alias("sync/atomic", "StoreInt64", "runtime/internal/atomic", "Store64", all...)

	alias("sync/atomic", "StoreUint32", "runtime/internal/atomic", "Store", all...)
	alias("sync/atomic", "StoreUint64", "runtime/internal/atomic", "Store64", all...)
	alias("sync/atomic", "StoreUintptr", "runtime/internal/atomic", "Store", p4...)
	alias("sync/atomic", "StoreUintptr", "runtime/internal/atomic", "Store64", p8...)

	alias("sync/atomic", "SwapInt32", "runtime/internal/atomic", "Xchg", all...)
	alias("sync/atomic", "SwapInt64", "runtime/internal/atomic", "Xchg64", all...)
	alias("sync/atomic", "SwapUint32", "runtime/internal/atomic", "Xchg", all...)
	alias("sync/atomic", "SwapUint64", "runtime/internal/atomic", "Xchg64", all...)
	alias("sync/atomic", "SwapUintptr", "runtime/internal/atomic", "Xchg", p4...)
	alias("sync/atomic", "SwapUintptr", "runtime/internal/atomic", "Xchg64", p8...)

	alias("sync/atomic", "CompareAndSwapInt32", "runtime/internal/atomic", "Cas", all...)
	alias("sync/atomic", "CompareAndSwapInt64", "runtime/internal/atomic", "Cas64", all...)
	alias("sync/atomic", "CompareAndSwapUint32", "runtime/internal/atomic", "Cas", all...)
	alias("sync/atomic", "CompareAndSwapUint64", "runtime/internal/atomic", "Cas64", all...)
	alias("sync/atomic", "CompareAndSwapUintptr", "runtime/internal/atomic", "Cas", p4...)
	alias("sync/atomic", "CompareAndSwapUintptr", "runtime/internal/atomic", "Cas64", p8...)

	alias("sync/atomic", "AddInt32", "runtime/internal/atomic", "Xadd", all...)
	alias("sync/atomic", "AddInt64", "runtime/internal/atomic", "Xadd64", all...)
	alias("sync/atomic", "AddUint32", "runtime/internal/atomic", "Xadd", all...)
	alias("sync/atomic", "AddUint64", "runtime/internal/atomic", "Xadd64", all...)
	alias("sync/atomic", "AddUintptr", "runtime/internal/atomic", "Xadd", p4...)
	alias("sync/atomic", "AddUintptr", "runtime/internal/atomic", "Xadd64", p8...)

	add("math/big", "mulWW",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue2(ssa.OpMul64uhilo, types.NewTuple(psess.types.Types[TUINT64], psess.types.Types[TUINT64]), args[0], args[1])
		}, psess.sys.
			ArchAMD64, psess.sys.ArchARM64)
	add("math/big", "divWW",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue3(ssa.OpDiv128u, types.NewTuple(psess.types.Types[TUINT64], psess.types.Types[TUINT64]), args[0], args[1], args[2])
		}, psess.sys.
			ArchAMD64)
}

// findIntrinsic returns a function which builds the SSA equivalent of the
// function identified by the symbol sym.  If sym is not an intrinsic call, returns nil.
func (psess *PackageSession) findIntrinsic(sym *types.Sym) intrinsicBuilder {
	if psess.ssa.IntrinsicsDisable {
		return nil
	}
	if sym == nil || sym.Pkg == nil {
		return nil
	}
	pkg := sym.Pkg.Path
	if sym.Pkg == psess.localpkg {
		pkg = psess.myimportpath
	}
	if psess.flag_race && pkg == "sync/atomic" {

		return nil
	}

	if psess.thearch.SoftFloat && pkg == "math" {
		return nil
	}

	fn := sym.Name
	return psess.intrinsics[intrinsicKey{psess.thearch.LinkArch.Arch, pkg, fn}]
}

func (psess *PackageSession) isIntrinsicCall(n *Node) bool {
	if n == nil || n.Left == nil {
		return false
	}
	return psess.findIntrinsic(n.Left.Sym) != nil
}

// intrinsicCall converts a call to a recognized intrinsic function into the intrinsic SSA operation.
func (s *state) intrinsicCall(psess *PackageSession, n *Node) *ssa.Value {
	v := psess.findIntrinsic(n.Left.Sym)(s, n, s.intrinsicArgs(psess, n))
	if psess.ssa.IntrinsicsDebug > 0 {
		x := v
		if x == nil {
			x = s.mem(psess)
		}
		if x.Op == ssa.OpSelect0 || x.Op == ssa.OpSelect1 {
			x = x.Args[0]
		}
		psess.
			Warnl(n.Pos, "intrinsic substitution for %v with %s", n.Left.Sym.Name, x.LongString(psess.ssa))
	}
	return v
}

type callArg struct {
	offset int64
	v      *ssa.Value
}
type byOffset []callArg

func (x byOffset) Len() int      { return len(x) }
func (x byOffset) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
func (x byOffset) Less(i, j int) bool {
	return x[i].offset < x[j].offset
}

// intrinsicArgs extracts args from n, evaluates them to SSA values, and returns them.
func (s *state) intrinsicArgs(psess *PackageSession, n *Node) []*ssa.Value {
	// This code is complicated because of how walk transforms calls. For a call node,
	// each entry in n.List is either an assignment to OINDREGSP which actually
	// stores an arg, or an assignment to a temporary which computes an arg
	// which is later assigned.
	// The args can also be out of order.
	// TODO: when walk goes away someday, this code can go away also.
	var args []callArg
	temps := map[*Node]*ssa.Value{}
	for _, a := range n.List.Slice() {
		if a.Op != OAS {
			s.Fatalf("non-assignment as a function argument %v", a.Op)
		}
		l, r := a.Left, a.Right
		switch l.Op {
		case ONAME:

			temps[l] = s.expr(psess, r)
		case OINDREGSP:
			// Store a value to an argument slot.
			var v *ssa.Value
			if x, ok := temps[r]; ok {

				v = x
			} else {

				v = s.expr(psess, r)
			}
			args = append(args, callArg{l.Xoffset, v})
		default:
			s.Fatalf("function argument assignment target not allowed: %v", l.Op)
		}
	}
	sort.Sort(byOffset(args))
	res := make([]*ssa.Value, len(args))
	for i, a := range args {
		res[i] = a.v
	}
	return res
}

// Calls the function n using the specified call type.
// Returns the address of the return value (or nil if none).
func (s *state) call(psess *PackageSession, n *Node, k callKind) *ssa.Value {
	var sym *types.Sym     // target symbol (if static)
	var closure *ssa.Value // ptr to closure to run (if dynamic)
	var codeptr *ssa.Value // ptr to target code (if dynamic)
	var rcvr *ssa.Value    // receiver to set
	fn := n.Left
	switch n.Op {
	case OCALLFUNC:
		if k == callNormal && fn.Op == ONAME && fn.Class() == PFUNC {
			sym = fn.Sym
			break
		}
		closure = s.expr(psess, fn)
	case OCALLMETH:
		if fn.Op != ODOTMETH {
			psess.
				Fatalf("OCALLMETH: n.Left not an ODOTMETH: %v", fn)
		}
		if k == callNormal {
			sym = fn.Sym
			break
		}

		n2 := psess.newnamel(fn.Pos, fn.Sym)
		n2.Name.Curfn = s.curfn
		n2.SetClass(PFUNC)
		n2.Pos = fn.Pos
		n2.Type = psess.types.Types[TUINT8]
		closure = s.expr(psess, n2)

	case OCALLINTER:
		if fn.Op != ODOTINTER {
			psess.
				Fatalf("OCALLINTER: n.Left not an ODOTINTER: %v", fn.Op)
		}
		i := s.expr(psess, fn.Left)
		itab := s.newValue1(ssa.OpITab, psess.types.Types[TUINTPTR], i)
		s.nilCheck(psess, itab)
		itabidx := fn.Xoffset + 2*int64(psess.Widthptr) + 8
		itab = s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.UintptrPtr, itabidx, itab)
		if k == callNormal {
			codeptr = s.load(psess, psess.types.Types[TUINTPTR], itab)
		} else {
			closure = itab
		}
		rcvr = s.newValue1(ssa.OpIData, psess.types.Types[TUINTPTR], i)
	}
	psess.
		dowidth(fn.Type)
	stksize := fn.Type.ArgWidth(psess.types)

	s.stmtList(psess, n.List)

	if rcvr != nil {
		argStart := psess.Ctxt.FixedFrameSize()
		if k != callNormal {
			argStart += int64(2 * psess.Widthptr)
		}
		addr := s.constOffPtrSP(psess, s.f.Config.Types.UintptrPtr, argStart)
		s.store(psess, psess.types.Types[TUINTPTR], addr, rcvr)
	}

	if k != callNormal {

		argStart := psess.Ctxt.FixedFrameSize()
		argsize := s.constInt32(psess, psess.types.Types[TUINT32], int32(stksize))
		addr := s.constOffPtrSP(psess, s.f.Config.Types.UInt32Ptr, argStart)
		s.store(psess, psess.types.Types[TUINT32], addr, argsize)
		addr = s.constOffPtrSP(psess, s.f.Config.Types.UintptrPtr, argStart+int64(psess.Widthptr))
		s.store(psess, psess.types.Types[TUINTPTR], addr, closure)
		stksize += 2 * int64(psess.Widthptr)
	}

	// call target
	var call *ssa.Value
	switch {
	case k == callDefer:
		call = s.newValue1A(ssa.OpStaticCall, psess.types.TypeMem, psess.Deferproc, s.mem(psess))
	case k == callGo:
		call = s.newValue1A(ssa.OpStaticCall, psess.types.TypeMem, psess.Newproc, s.mem(psess))
	case closure != nil:

		codeptr = s.rawLoad(psess, psess.types.Types[TUINTPTR], closure)
		call = s.newValue3(ssa.OpClosureCall, psess.types.TypeMem, codeptr, closure, s.mem(psess))
	case codeptr != nil:
		call = s.newValue2(ssa.OpInterCall, psess.types.TypeMem, codeptr, s.mem(psess))
	case sym != nil:
		call = s.newValue1A(ssa.OpStaticCall, psess.types.TypeMem, sym.Linksym(psess.types), s.mem(psess))
	default:
		psess.
			Fatalf("bad call type %v %v", n.Op, n)
	}
	call.AuxInt = stksize
	s.vars[&psess.memVar] = call

	if k == callDefer {
		b := s.endBlock(psess)
		b.Kind = ssa.BlockDefer
		b.SetControl(call)
		bNext := s.f.NewBlock(ssa.BlockPlain)
		b.AddEdgeTo(bNext)

		r := s.f.NewBlock(ssa.BlockPlain)
		s.startBlock(r)
		s.exit(psess)
		b.AddEdgeTo(r)
		b.Likely = ssa.BranchLikely
		s.startBlock(bNext)
	}

	res := n.Left.Type.Results(psess.types)
	if res.NumFields(psess.types) == 0 || k != callNormal {

		return nil
	}
	fp := res.Field(psess.types, 0)
	return s.constOffPtrSP(psess, psess.types.NewPtr(fp.Type), fp.Offset+psess.Ctxt.FixedFrameSize())
}

// etypesign returns the signed-ness of e, for integer/pointer etypes.
// -1 means signed, +1 means unsigned, 0 means non-integer/non-pointer.
func etypesign(e types.EType) int8 {
	switch e {
	case TINT8, TINT16, TINT32, TINT64, TINT:
		return -1
	case TUINT8, TUINT16, TUINT32, TUINT64, TUINT, TUINTPTR, TUNSAFEPTR:
		return +1
	}
	return 0
}

// addr converts the address of the expression n to SSA, adds it to s and returns the SSA result.
// The value that the returned Value represents is guaranteed to be non-nil.
// If bounded is true then this address does not require a nil check for its operand
// even if that would otherwise be implied.
func (s *state) addr(psess *PackageSession, n *Node, bounded bool) *ssa.Value {
	t := psess.types.NewPtr(n.Type)
	switch n.Op {
	case ONAME:
		switch n.Class() {
		case PEXTERN:

			v := s.entryNewValue1A(psess, ssa.OpAddr, t, n.Sym.Linksym(psess.types), s.sb)

			if n.Xoffset != 0 {
				v = s.entryNewValue1I(psess, ssa.OpOffPtr, v.Type, n.Xoffset, v)
			}
			return v
		case PPARAM:

			v := s.decladdrs[n]
			if v != nil {
				return v
			}
			if n == psess.nodfp {

				return s.entryNewValue1A(psess, ssa.OpAddr, t, n, s.sp)
			}
			s.Fatalf("addr of undeclared ONAME %v. declared: %v", n, s.decladdrs)
			return nil
		case PAUTO:
			return s.newValue1Apos(ssa.OpAddr, t, n, s.sp, !n.IsAutoTmp())
		case PPARAMOUT:

			return s.newValue1A(ssa.OpAddr, t, n, s.sp)
		default:
			s.Fatalf("variable address class %v not implemented", n.Class())
			return nil
		}
	case OINDREGSP:

		return s.constOffPtrSP(psess, t, n.Xoffset)
	case OINDEX:
		if n.Left.Type.IsSlice() {
			a := s.expr(psess, n.Left)
			i := s.expr(psess, n.Right)
			i = s.extendIndex(psess, i, psess.panicindex)
			len := s.newValue1(ssa.OpSliceLen, psess.types.Types[TINT], a)
			if !n.Bounded() {
				s.boundsCheck(psess, i, len)
			}
			p := s.newValue1(ssa.OpSlicePtr, t, a)
			return s.newValue2(ssa.OpPtrIndex, t, p, i)
		} else {
			a := s.addr(psess, n.Left, bounded)
			i := s.expr(psess, n.Right)
			i = s.extendIndex(psess, i, psess.panicindex)
			len := s.constInt(psess, psess.types.Types[TINT], n.Left.Type.NumElem(psess.types))
			if !n.Bounded() {
				s.boundsCheck(psess, i, len)
			}
			return s.newValue2(ssa.OpPtrIndex, psess.types.NewPtr(n.Left.Type.Elem(psess.types)), a, i)
		}
	case OIND:
		return s.exprPtr(psess, n.Left, bounded, n.Pos)
	case ODOT:
		p := s.addr(psess, n.Left, bounded)
		return s.newValue1I(ssa.OpOffPtr, t, n.Xoffset, p)
	case ODOTPTR:
		p := s.exprPtr(psess, n.Left, bounded, n.Pos)
		return s.newValue1I(ssa.OpOffPtr, t, n.Xoffset, p)
	case OCLOSUREVAR:
		return s.newValue1I(ssa.OpOffPtr, t, n.Xoffset,
			s.entryNewValue0(psess, ssa.OpGetClosurePtr, s.f.Config.Types.BytePtr))
	case OCONVNOP:
		addr := s.addr(psess, n.Left, bounded)
		return s.newValue1(ssa.OpCopy, t, addr)
	case OCALLFUNC, OCALLINTER, OCALLMETH:
		return s.call(psess, n, callNormal)
	case ODOTTYPE:
		v, _ := s.dottype(psess, n, false)
		if v.Op != ssa.OpLoad {
			s.Fatalf("dottype of non-load")
		}
		if v.Args[1] != s.mem(psess) {
			s.Fatalf("memory no longer live from dottype load")
		}
		return v.Args[0]
	default:
		s.Fatalf("unhandled addr %v", n.Op)
		return nil
	}
}

// canSSA reports whether n is SSA-able.
// n must be an ONAME (or an ODOT sequence with an ONAME base).
func (s *state) canSSA(psess *PackageSession, n *Node) bool {
	if psess.Debug['N'] != 0 {
		return false
	}
	for n.Op == ODOT || (n.Op == OINDEX && n.Left.Type.IsArray()) {
		n = n.Left
	}
	if n.Op != ONAME {
		return false
	}
	if n.Addrtaken() {
		return false
	}
	if n.isParamHeapCopy() {
		return false
	}
	if n.Class() == PAUTOHEAP {
		psess.
			Fatalf("canSSA of PAUTOHEAP %v", n)
	}
	switch n.Class() {
	case PEXTERN:
		return false
	case PPARAMOUT:
		if s.hasdefer {

			return false
		}
		if s.cgoUnsafeArgs {

			return false
		}
	}
	if n.Class() == PPARAM && n.Sym != nil && n.Sym.Name == ".this" {

		return false
	}
	return psess.canSSAType(n.Type)

}

// canSSA reports whether variables of type t are SSA-able.
func (psess *PackageSession) canSSAType(t *types.Type) bool {
	psess.
		dowidth(t)
	if t.Width > int64(4*psess.Widthptr) {

		return false
	}
	switch t.Etype {
	case TARRAY:

		if t.NumElem(psess.types) <= 1 {
			return psess.canSSAType(t.Elem(psess.types))
		}
		return false
	case TSTRUCT:
		if t.NumFields(psess.types) > ssa.MaxStruct {
			return false
		}
		for _, t1 := range t.Fields(psess.types).Slice() {
			if !psess.canSSAType(t1.Type) {
				return false
			}
		}
		return true
	default:
		return true
	}
}

// exprPtr evaluates n to a pointer and nil-checks it.
func (s *state) exprPtr(psess *PackageSession, n *Node, bounded bool, lineno src.XPos) *ssa.Value {
	p := s.expr(psess, n)
	if bounded || n.NonNil() {
		if s.f.Frontend().Debug_checknil() && lineno.Line() > 1 {
			s.f.Warnl(lineno, "removed nil check")
		}
		return p
	}
	s.nilCheck(psess, p)
	return p
}

// nilCheck generates nil pointer checking code.
// Used only for automatically inserted nil checks,
// not for user code like 'x != nil'.
func (s *state) nilCheck(psess *PackageSession, ptr *ssa.Value) {
	if psess.disable_checknil != 0 || s.curfn.Func.NilCheckDisabled() {
		return
	}
	s.newValue2(ssa.OpNilCheck, psess.types.TypeVoid, ptr, s.mem(psess))
}

// boundsCheck generates bounds checking code. Checks if 0 <= idx < len, branches to exit if not.
// Starts a new block on return.
// idx is already converted to full int width.
func (s *state) boundsCheck(psess *PackageSession, idx, len *ssa.Value) {
	if psess.Debug['B'] != 0 {
		return
	}

	cmp := s.newValue2(ssa.OpIsInBounds, psess.types.Types[TBOOL], idx, len)
	s.check(psess, cmp, psess.panicindex)
}

// sliceBoundsCheck generates slice bounds checking code. Checks if 0 <= idx <= len, branches to exit if not.
// Starts a new block on return.
// idx and len are already converted to full int width.
func (s *state) sliceBoundsCheck(psess *PackageSession, idx, len *ssa.Value) {
	if psess.Debug['B'] != 0 {
		return
	}

	cmp := s.newValue2(ssa.OpIsSliceInBounds, psess.types.Types[TBOOL], idx, len)
	s.check(psess, cmp, psess.panicslice)
}

// If cmp (a bool) is false, panic using the given function.
func (s *state) check(psess *PackageSession, cmp *ssa.Value, fn *obj.LSym) {
	b := s.endBlock(psess)
	b.Kind = ssa.BlockIf
	b.SetControl(cmp)
	b.Likely = ssa.BranchLikely
	bNext := s.f.NewBlock(ssa.BlockPlain)
	line := s.peekPos()
	pos := psess.Ctxt.PosTable.Pos(line)
	fl := funcLine{f: fn, base: pos.Base(), line: pos.Line()}
	bPanic := s.panics[fl]
	if bPanic == nil {
		bPanic = s.f.NewBlock(ssa.BlockPlain)
		s.panics[fl] = bPanic
		s.startBlock(bPanic)

		s.rtcall(psess, fn, false, nil)
	}
	b.AddEdgeTo(bNext)
	b.AddEdgeTo(bPanic)
	s.startBlock(bNext)
}

func (s *state) intDivide(psess *PackageSession, n *Node, a, b *ssa.Value) *ssa.Value {
	needcheck := true
	switch b.Op {
	case ssa.OpConst8, ssa.OpConst16, ssa.OpConst32, ssa.OpConst64:
		if b.AuxInt != 0 {
			needcheck = false
		}
	}
	if needcheck {

		cmp := s.newValue2(s.ssaOp(psess, ONE, n.Type), psess.types.Types[TBOOL], b, s.zeroVal(psess, n.Type))
		s.check(psess, cmp, psess.panicdivide)
	}
	return s.newValue2(s.ssaOp(psess, n.Op, n.Type), a.Type, a, b)
}

// rtcall issues a call to the given runtime function fn with the listed args.
// Returns a slice of results of the given result types.
// The call is added to the end of the current block.
// If returns is false, the block is marked as an exit block.
func (s *state) rtcall(psess *PackageSession, fn *obj.LSym, returns bool, results []*types.Type, args ...*ssa.Value) []*ssa.Value {

	off := psess.Ctxt.FixedFrameSize()
	for _, arg := range args {
		t := arg.Type
		off = psess.Rnd(off, t.Alignment(psess.types))
		ptr := s.constOffPtrSP(psess, t.PtrTo(psess.types), off)
		size := t.Size(psess.types)
		s.store(psess, t, ptr, arg)
		off += size
	}
	off = psess.Rnd(off, int64(psess.Widthreg))

	call := s.newValue1A(ssa.OpStaticCall, psess.types.TypeMem, fn, s.mem(psess))
	s.vars[&psess.memVar] = call

	if !returns {

		b := s.endBlock(psess)
		b.Kind = ssa.BlockExit
		b.SetControl(call)
		call.AuxInt = off - psess.Ctxt.FixedFrameSize()
		if len(results) > 0 {
			psess.
				Fatalf("panic call can't have results")
		}
		return nil
	}

	res := make([]*ssa.Value, len(results))
	for i, t := range results {
		off = psess.Rnd(off, t.Alignment(psess.types))
		ptr := s.constOffPtrSP(psess, psess.types.NewPtr(t), off)
		res[i] = s.load(psess, t, ptr)
		off += t.Size(psess.types)
	}
	off = psess.Rnd(off, int64(psess.Widthptr))

	call.AuxInt = off

	return res
}

// do *left = right for type t.
func (s *state) storeType(psess *PackageSession, t *types.Type, left, right *ssa.Value, skip skipMask, leftIsStmt bool) {
	s.instrument(psess, t, left, true)

	if skip == 0 && (!psess.types.Haspointers(t) || ssa.IsStackAddr(left)) {

		s.vars[&psess.memVar] = s.newValue3Apos(ssa.OpStore, psess.types.TypeMem, t, left, right, s.mem(psess), leftIsStmt)
		return
	}

	s.storeTypeScalars(psess, t, left, right, skip)
	if skip&skipPtr == 0 && psess.types.Haspointers(t) {
		s.storeTypePtrs(psess, t, left, right)
	}
}

// do *left = right for all scalar (non-pointer) parts of t.
func (s *state) storeTypeScalars(psess *PackageSession, t *types.Type, left, right *ssa.Value, skip skipMask) {
	switch {
	case t.IsBoolean() || t.IsInteger() || t.IsFloat() || t.IsComplex():
		s.store(psess, t, left, right)
	case t.IsPtrShaped():

	case t.IsString():
		if skip&skipLen != 0 {
			return
		}
		len := s.newValue1(ssa.OpStringLen, psess.types.Types[TINT], right)
		lenAddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.IntPtr, s.config.PtrSize, left)
		s.store(psess, psess.types.Types[TINT], lenAddr, len)
	case t.IsSlice():
		if skip&skipLen == 0 {
			len := s.newValue1(ssa.OpSliceLen, psess.types.Types[TINT], right)
			lenAddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.IntPtr, s.config.PtrSize, left)
			s.store(psess, psess.types.Types[TINT], lenAddr, len)
		}
		if skip&skipCap == 0 {
			cap := s.newValue1(ssa.OpSliceCap, psess.types.Types[TINT], right)
			capAddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.IntPtr, 2*s.config.PtrSize, left)
			s.store(psess, psess.types.Types[TINT], capAddr, cap)
		}
	case t.IsInterface():

		itab := s.newValue1(ssa.OpITab, s.f.Config.Types.BytePtr, right)
		s.store(psess, psess.types.Types[TUINTPTR], left, itab)
	case t.IsStruct():
		n := t.NumFields(psess.types)
		for i := 0; i < n; i++ {
			ft := t.FieldType(psess.types, i)
			addr := s.newValue1I(ssa.OpOffPtr, ft.PtrTo(psess.types), t.FieldOff(psess.types, i), left)
			val := s.newValue1I(ssa.OpStructSelect, ft, int64(i), right)
			s.storeTypeScalars(psess, ft, addr, val, 0)
		}
	case t.IsArray() && t.NumElem(psess.types) == 0:

	case t.IsArray() && t.NumElem(psess.types) == 1:
		s.storeTypeScalars(psess, t.Elem(psess.types), left, s.newValue1I(ssa.OpArraySelect, t.Elem(psess.types), 0, right), 0)
	default:
		s.Fatalf("bad write barrier type %v", t)
	}
}

// do *left = right for all pointer parts of t.
func (s *state) storeTypePtrs(psess *PackageSession, t *types.Type, left, right *ssa.Value) {
	switch {
	case t.IsPtrShaped():
		s.store(psess, t, left, right)
	case t.IsString():
		ptr := s.newValue1(ssa.OpStringPtr, s.f.Config.Types.BytePtr, right)
		s.store(psess, s.f.Config.Types.BytePtr, left, ptr)
	case t.IsSlice():
		elType := psess.types.NewPtr(t.Elem(psess.types))
		ptr := s.newValue1(ssa.OpSlicePtr, elType, right)
		s.store(psess, elType, left, ptr)
	case t.IsInterface():

		idata := s.newValue1(ssa.OpIData, s.f.Config.Types.BytePtr, right)
		idataAddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.BytePtrPtr, s.config.PtrSize, left)
		s.store(psess, s.f.Config.Types.BytePtr, idataAddr, idata)
	case t.IsStruct():
		n := t.NumFields(psess.types)
		for i := 0; i < n; i++ {
			ft := t.FieldType(psess.types, i)
			if !psess.types.Haspointers(ft) {
				continue
			}
			addr := s.newValue1I(ssa.OpOffPtr, ft.PtrTo(psess.types), t.FieldOff(psess.types, i), left)
			val := s.newValue1I(ssa.OpStructSelect, ft, int64(i), right)
			s.storeTypePtrs(psess, ft, addr, val)
		}
	case t.IsArray() && t.NumElem(psess.types) == 0:

	case t.IsArray() && t.NumElem(psess.types) == 1:
		s.storeTypePtrs(psess, t.Elem(psess.types), left, s.newValue1I(ssa.OpArraySelect, t.Elem(psess.types), 0, right))
	default:
		s.Fatalf("bad write barrier type %v", t)
	}
}

// slice computes the slice v[i:j:k] and returns ptr, len, and cap of result.
// i,j,k may be nil, in which case they are set to their default value.
// t is a slice, ptr to array, or string type.
func (s *state) slice(psess *PackageSession, t *types.Type, v, i, j, k *ssa.Value) (p, l, c *ssa.Value) {
	var elemtype *types.Type
	var ptrtype *types.Type
	var ptr *ssa.Value
	var len *ssa.Value
	var cap *ssa.Value
	zero := s.constInt(psess, psess.types.Types[TINT], 0)
	switch {
	case t.IsSlice():
		elemtype = t.Elem(psess.types)
		ptrtype = psess.types.NewPtr(elemtype)
		ptr = s.newValue1(ssa.OpSlicePtr, ptrtype, v)
		len = s.newValue1(ssa.OpSliceLen, psess.types.Types[TINT], v)
		cap = s.newValue1(ssa.OpSliceCap, psess.types.Types[TINT], v)
	case t.IsString():
		elemtype = psess.types.Types[TUINT8]
		ptrtype = psess.types.NewPtr(elemtype)
		ptr = s.newValue1(ssa.OpStringPtr, ptrtype, v)
		len = s.newValue1(ssa.OpStringLen, psess.types.Types[TINT], v)
		cap = len
	case t.IsPtr():
		if !t.Elem(psess.types).IsArray() {
			s.Fatalf("bad ptr to array in slice %v\n", t)
		}
		elemtype = t.Elem(psess.types).Elem(psess.types)
		ptrtype = psess.types.NewPtr(elemtype)
		s.nilCheck(psess, v)
		ptr = v
		len = s.constInt(psess, psess.types.Types[TINT], t.Elem(psess.types).NumElem(psess.types))
		cap = len
	default:
		s.Fatalf("bad type in slice %v\n", t)
	}

	if i == nil {
		i = zero
	}
	if j == nil {
		j = len
	}
	if k == nil {
		k = cap
	}

	s.sliceBoundsCheck(psess, i, j)
	if j != k {
		s.sliceBoundsCheck(psess, j, k)
	}
	if k != cap {
		s.sliceBoundsCheck(psess, k, cap)
	}

	subOp := s.ssaOp(psess, OSUB, psess.types.Types[TINT])
	mulOp := s.ssaOp(psess, OMUL, psess.types.Types[TINT])
	andOp := s.ssaOp(psess, OAND, psess.types.Types[TINT])
	rlen := s.newValue2(subOp, psess.types.Types[TINT], j, i)
	var rcap *ssa.Value
	switch {
	case t.IsString():

		rcap = rlen
	case j == k:
		rcap = rlen
	default:
		rcap = s.newValue2(subOp, psess.types.Types[TINT], k, i)
	}

	var rptr *ssa.Value
	if (i.Op == ssa.OpConst64 || i.Op == ssa.OpConst32) && i.AuxInt == 0 {

		rptr = ptr
	} else {

		delta := s.newValue2(mulOp, psess.types.Types[TINT], i, s.constInt(psess, psess.types.Types[TINT], elemtype.Width))

		mask := s.newValue1(ssa.OpSlicemask, psess.types.Types[TINT], rcap)
		delta = s.newValue2(andOp, psess.types.Types[TINT], delta, mask)

		rptr = s.newValue2(ssa.OpAddPtr, ptrtype, ptr, delta)
	}

	return rptr, rlen, rcap
}

type u642fcvtTab struct {
	geq, cvt2F, and, rsh, or, add ssa.Op
	one                           func(*state, *types.Type, int64) *ssa.Value
}

func (s *state) uint64Tofloat64(psess *PackageSession, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.uint64Tofloat(psess, &psess.u64_f64, n, x, ft, tt)
}

func (s *state) uint64Tofloat32(psess *PackageSession, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.uint64Tofloat(psess, &psess.u64_f32, n, x, ft, tt)
}

func (s *state) uint64Tofloat(psess *PackageSession, cvttab *u642fcvtTab, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {

	cmp := s.newValue2(cvttab.geq, psess.types.Types[TBOOL], x, s.zeroVal(psess, ft))
	b := s.endBlock(psess)
	b.Kind = ssa.BlockIf
	b.SetControl(cmp)
	b.Likely = ssa.BranchLikely

	bThen := s.f.NewBlock(ssa.BlockPlain)
	bElse := s.f.NewBlock(ssa.BlockPlain)
	bAfter := s.f.NewBlock(ssa.BlockPlain)

	b.AddEdgeTo(bThen)
	s.startBlock(bThen)
	a0 := s.newValue1(cvttab.cvt2F, tt, x)
	s.vars[n] = a0
	s.endBlock(psess)
	bThen.AddEdgeTo(bAfter)

	b.AddEdgeTo(bElse)
	s.startBlock(bElse)
	one := cvttab.one(s, ft, 1)
	y := s.newValue2(cvttab.and, ft, x, one)
	z := s.newValue2(cvttab.rsh, ft, x, one)
	z = s.newValue2(cvttab.or, ft, z, y)
	a := s.newValue1(cvttab.cvt2F, tt, z)
	a1 := s.newValue2(cvttab.add, tt, a, a)
	s.vars[n] = a1
	s.endBlock(psess)
	bElse.AddEdgeTo(bAfter)

	s.startBlock(bAfter)
	return s.variable(n, n.Type)
}

type u322fcvtTab struct {
	cvtI2F, cvtF2F ssa.Op
}

func (s *state) uint32Tofloat64(psess *PackageSession, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.uint32Tofloat(psess, &psess.u32_f64, n, x, ft, tt)
}

func (s *state) uint32Tofloat32(psess *PackageSession, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.uint32Tofloat(psess, &psess.u32_f32, n, x, ft, tt)
}

func (s *state) uint32Tofloat(psess *PackageSession, cvttab *u322fcvtTab, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {

	cmp := s.newValue2(ssa.OpGeq32, psess.types.Types[TBOOL], x, s.zeroVal(psess, ft))
	b := s.endBlock(psess)
	b.Kind = ssa.BlockIf
	b.SetControl(cmp)
	b.Likely = ssa.BranchLikely

	bThen := s.f.NewBlock(ssa.BlockPlain)
	bElse := s.f.NewBlock(ssa.BlockPlain)
	bAfter := s.f.NewBlock(ssa.BlockPlain)

	b.AddEdgeTo(bThen)
	s.startBlock(bThen)
	a0 := s.newValue1(cvttab.cvtI2F, tt, x)
	s.vars[n] = a0
	s.endBlock(psess)
	bThen.AddEdgeTo(bAfter)

	b.AddEdgeTo(bElse)
	s.startBlock(bElse)
	a1 := s.newValue1(ssa.OpCvt32to64F, psess.types.Types[TFLOAT64], x)
	twoToThe32 := s.constFloat64(psess, psess.types.Types[TFLOAT64], float64(1<<32))
	a2 := s.newValue2(ssa.OpAdd64F, psess.types.Types[TFLOAT64], a1, twoToThe32)
	a3 := s.newValue1(cvttab.cvtF2F, tt, a2)

	s.vars[n] = a3
	s.endBlock(psess)
	bElse.AddEdgeTo(bAfter)

	s.startBlock(bAfter)
	return s.variable(n, n.Type)
}

// referenceTypeBuiltin generates code for the len/cap builtins for maps and channels.
func (s *state) referenceTypeBuiltin(psess *PackageSession, n *Node, x *ssa.Value) *ssa.Value {
	if !n.Left.Type.IsMap() && !n.Left.Type.IsChan() {
		s.Fatalf("node must be a map or a channel")
	}

	lenType := n.Type
	nilValue := s.constNil(psess, psess.types.Types[TUINTPTR])
	cmp := s.newValue2(ssa.OpEqPtr, psess.types.Types[TBOOL], x, nilValue)
	b := s.endBlock(psess)
	b.Kind = ssa.BlockIf
	b.SetControl(cmp)
	b.Likely = ssa.BranchUnlikely

	bThen := s.f.NewBlock(ssa.BlockPlain)
	bElse := s.f.NewBlock(ssa.BlockPlain)
	bAfter := s.f.NewBlock(ssa.BlockPlain)

	b.AddEdgeTo(bThen)
	s.startBlock(bThen)
	s.vars[n] = s.zeroVal(psess, lenType)
	s.endBlock(psess)
	bThen.AddEdgeTo(bAfter)

	b.AddEdgeTo(bElse)
	s.startBlock(bElse)
	switch n.Op {
	case OLEN:

		s.vars[n] = s.load(psess, lenType, x)
	case OCAP:

		sw := s.newValue1I(ssa.OpOffPtr, lenType.PtrTo(psess.types), lenType.Width, x)
		s.vars[n] = s.load(psess, lenType, sw)
	default:
		s.Fatalf("op must be OLEN or OCAP")
	}
	s.endBlock(psess)
	bElse.AddEdgeTo(bAfter)

	s.startBlock(bAfter)
	return s.variable(n, lenType)
}

type f2uCvtTab struct {
	ltf, cvt2U, subf, or ssa.Op
	floatValue           func(*state, *types.Type, float64) *ssa.Value
	intValue             func(*state, *types.Type, int64) *ssa.Value
	cutoff               uint64
}

func (s *state) float32ToUint64(psess *PackageSession, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.floatToUint(psess, &psess.f32_u64, n, x, ft, tt)
}
func (s *state) float64ToUint64(psess *PackageSession, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.floatToUint(psess, &psess.f64_u64, n, x, ft, tt)
}

func (s *state) float32ToUint32(psess *PackageSession, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.floatToUint(psess, &psess.f32_u32, n, x, ft, tt)
}

func (s *state) float64ToUint32(psess *PackageSession, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.floatToUint(psess, &psess.f64_u32, n, x, ft, tt)
}

func (s *state) floatToUint(psess *PackageSession, cvttab *f2uCvtTab, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {

	cutoff := cvttab.floatValue(s, ft, float64(cvttab.cutoff))
	cmp := s.newValue2(cvttab.ltf, psess.types.Types[TBOOL], x, cutoff)
	b := s.endBlock(psess)
	b.Kind = ssa.BlockIf
	b.SetControl(cmp)
	b.Likely = ssa.BranchLikely

	bThen := s.f.NewBlock(ssa.BlockPlain)
	bElse := s.f.NewBlock(ssa.BlockPlain)
	bAfter := s.f.NewBlock(ssa.BlockPlain)

	b.AddEdgeTo(bThen)
	s.startBlock(bThen)
	a0 := s.newValue1(cvttab.cvt2U, tt, x)
	s.vars[n] = a0
	s.endBlock(psess)
	bThen.AddEdgeTo(bAfter)

	b.AddEdgeTo(bElse)
	s.startBlock(bElse)
	y := s.newValue2(cvttab.subf, ft, x, cutoff)
	y = s.newValue1(cvttab.cvt2U, tt, y)
	z := cvttab.intValue(s, tt, int64(-cvttab.cutoff))
	a1 := s.newValue2(cvttab.or, tt, y, z)
	s.vars[n] = a1
	s.endBlock(psess)
	bElse.AddEdgeTo(bAfter)

	s.startBlock(bAfter)
	return s.variable(n, n.Type)
}

// dottype generates SSA for a type assertion node.
// commaok indicates whether to panic or return a bool.
// If commaok is false, resok will be nil.
func (s *state) dottype(psess *PackageSession, n *Node, commaok bool) (res, resok *ssa.Value) {
	iface := s.expr(psess, n.Left)
	target := s.expr(psess, n.Right)
	byteptr := s.f.Config.Types.BytePtr

	if n.Type.IsInterface() {
		if n.Type.IsEmptyInterface(psess.types) {

			if psess.Debug_typeassert > 0 {
				psess.
					Warnl(n.Pos, "type assertion inlined")
			}

			itab := s.newValue1(ssa.OpITab, byteptr, iface)

			cond := s.newValue2(ssa.OpNeqPtr, psess.types.Types[TBOOL], itab, s.constNil(psess, byteptr))

			if n.Left.Type.IsEmptyInterface(psess.types) && commaok {

				return iface, cond
			}

			b := s.endBlock(psess)
			b.Kind = ssa.BlockIf
			b.SetControl(cond)
			b.Likely = ssa.BranchLikely
			bOk := s.f.NewBlock(ssa.BlockPlain)
			bFail := s.f.NewBlock(ssa.BlockPlain)
			b.AddEdgeTo(bOk)
			b.AddEdgeTo(bFail)

			if !commaok {

				s.startBlock(bFail)
				s.rtcall(psess, psess.panicnildottype, false, nil, target)

				s.startBlock(bOk)
				if n.Left.Type.IsEmptyInterface(psess.types) {
					res = iface
					return
				}

				off := s.newValue1I(ssa.OpOffPtr, byteptr, int64(psess.Widthptr), itab)
				typ := s.load(psess, byteptr, off)
				idata := s.newValue1(ssa.OpIData, n.Type, iface)
				res = s.newValue2(ssa.OpIMake, n.Type, typ, idata)
				return
			}

			s.startBlock(bOk)

			off := s.newValue1I(ssa.OpOffPtr, byteptr, int64(psess.Widthptr), itab)
			s.vars[&psess.typVar] = s.load(psess, byteptr, off)
			s.endBlock(psess)

			s.startBlock(bFail)
			s.vars[&psess.typVar] = itab
			s.endBlock(psess)

			bEnd := s.f.NewBlock(ssa.BlockPlain)
			bOk.AddEdgeTo(bEnd)
			bFail.AddEdgeTo(bEnd)
			s.startBlock(bEnd)
			idata := s.newValue1(ssa.OpIData, n.Type, iface)
			res = s.newValue2(ssa.OpIMake, n.Type, s.variable(&psess.typVar, byteptr), idata)
			resok = cond
			delete(s.vars, &psess.typVar)
			return
		}

		if psess.Debug_typeassert > 0 {
			psess.
				Warnl(n.Pos, "type assertion not inlined")
		}
		if n.Left.Type.IsEmptyInterface(psess.types) {
			if commaok {
				call := s.rtcall(psess, psess.assertE2I2, true, []*types.Type{n.Type, psess.types.Types[TBOOL]}, target, iface)
				return call[0], call[1]
			}
			return s.rtcall(psess, psess.assertE2I, true, []*types.Type{n.Type}, target, iface)[0], nil
		}
		if commaok {
			call := s.rtcall(psess, psess.assertI2I2, true, []*types.Type{n.Type, psess.types.Types[TBOOL]}, target, iface)
			return call[0], call[1]
		}
		return s.rtcall(psess, psess.assertI2I, true, []*types.Type{n.Type}, target, iface)[0], nil
	}

	if psess.Debug_typeassert > 0 {
		psess.
			Warnl(n.Pos, "type assertion inlined")
	}

	direct := psess.isdirectiface(n.Type)
	itab := s.newValue1(ssa.OpITab, byteptr, iface)
	if psess.Debug_typeassert > 0 {
		psess.
			Warnl(n.Pos, "type assertion inlined")
	}
	var targetITab *ssa.Value
	if n.Left.Type.IsEmptyInterface(psess.types) {

		targetITab = target
	} else {

		targetITab = s.expr(psess, n.List.First())
	}

	var tmp *Node       // temporary for use with large types
	var addr *ssa.Value // address of tmp
	if commaok && !psess.canSSAType(n.Type) {

		tmp = psess.tempAt(n.Pos, s.curfn, n.Type)
		addr = s.addr(psess, tmp, false)
		s.vars[&psess.memVar] = s.newValue1A(ssa.OpVarDef, psess.types.TypeMem, tmp, s.mem(psess))
	}

	cond := s.newValue2(ssa.OpEqPtr, psess.types.Types[TBOOL], itab, targetITab)
	b := s.endBlock(psess)
	b.Kind = ssa.BlockIf
	b.SetControl(cond)
	b.Likely = ssa.BranchLikely

	bOk := s.f.NewBlock(ssa.BlockPlain)
	bFail := s.f.NewBlock(ssa.BlockPlain)
	b.AddEdgeTo(bOk)
	b.AddEdgeTo(bFail)

	if !commaok {

		s.startBlock(bFail)
		taddr := s.expr(psess, n.Right.Right)
		if n.Left.Type.IsEmptyInterface(psess.types) {
			s.rtcall(psess, psess.panicdottypeE, false, nil, itab, target, taddr)
		} else {
			s.rtcall(psess, psess.panicdottypeI, false, nil, itab, target, taddr)
		}

		s.startBlock(bOk)
		if direct {
			return s.newValue1(ssa.OpIData, n.Type, iface), nil
		}
		p := s.newValue1(ssa.OpIData, psess.types.NewPtr(n.Type), iface)
		return s.load(psess, n.Type, p), nil
	}

	bEnd := s.f.NewBlock(ssa.BlockPlain)

	valVar := &Node{Op: ONAME, Sym: &types.Sym{Name: "val"}}

	s.startBlock(bOk)
	if tmp == nil {
		if direct {
			s.vars[valVar] = s.newValue1(ssa.OpIData, n.Type, iface)
		} else {
			p := s.newValue1(ssa.OpIData, psess.types.NewPtr(n.Type), iface)
			s.vars[valVar] = s.load(psess, n.Type, p)
		}
	} else {
		p := s.newValue1(ssa.OpIData, psess.types.NewPtr(n.Type), iface)
		s.move(psess, n.Type, addr, p)
	}
	s.vars[&psess.okVar] = s.constBool(psess, true)
	s.endBlock(psess)
	bOk.AddEdgeTo(bEnd)

	s.startBlock(bFail)
	if tmp == nil {
		s.vars[valVar] = s.zeroVal(psess, n.Type)
	} else {
		s.zero(psess, n.Type, addr)
	}
	s.vars[&psess.okVar] = s.constBool(psess, false)
	s.endBlock(psess)
	bFail.AddEdgeTo(bEnd)

	s.startBlock(bEnd)
	if tmp == nil {
		res = s.variable(valVar, n.Type)
		delete(s.vars, valVar)
	} else {
		res = s.load(psess, n.Type, addr)
		s.vars[&psess.memVar] = s.newValue1A(ssa.OpVarKill, psess.types.TypeMem, tmp, s.mem(psess))
	}
	resok = s.variable(&psess.okVar, psess.types.Types[TBOOL])
	delete(s.vars, &psess.okVar)
	return res, resok
}

// variable returns the value of a variable at the current location.
func (s *state) variable(name *Node, t *types.Type) *ssa.Value {
	v := s.vars[name]
	if v != nil {
		return v
	}
	v = s.fwdVars[name]
	if v != nil {
		return v
	}

	if s.curBlock == s.f.Entry {

		s.Fatalf("Value live at entry. It shouldn't be. func %s, node %v, value %v", s.f.Name, name, v)
	}

	v = s.newValue0A(ssa.OpFwdRef, t, name)
	s.fwdVars[name] = v
	s.addNamedValue(name, v)
	return v
}

func (s *state) mem(psess *PackageSession) *ssa.Value {
	return s.variable(&psess.memVar, psess.types.TypeMem)
}

func (s *state) addNamedValue(n *Node, v *ssa.Value) {
	if n.Class() == Pxxx {

		return
	}
	if n.IsAutoTmp() {

		return
	}
	if n.Class() == PPARAMOUT {

		return
	}
	if n.Class() == PAUTO && n.Xoffset != 0 {
		s.Fatalf("AUTO var with offset %v %d", n, n.Xoffset)
	}
	loc := ssa.LocalSlot{N: n, Type: n.Type, Off: 0}
	values, ok := s.f.NamedValues[loc]
	if !ok {
		s.f.Names = append(s.f.Names, loc)
	}
	s.f.NamedValues[loc] = append(values, v)
}

// Branch is an unresolved branch.
type Branch struct {
	P *obj.Prog  // branch instruction
	B *ssa.Block // target
}

// SSAGenState contains state needed during Prog generation.
type SSAGenState struct {
	pp *Progs

	// Branches remembers all the branch instructions we've seen
	// and where they would like to go.
	Branches []Branch

	// bstart remembers where each block starts (indexed by block ID)
	bstart []*obj.Prog

	// 387 port: maps from SSE registers (REG_X?) to 387 registers (REG_F?)
	SSEto387 map[int16]int16
	// Some architectures require a 64-bit temporary for FP-related register shuffling. Examples include x86-387, PPC, and Sparc V8.
	ScratchFpMem *Node

	maxarg int64 // largest frame size for arguments to calls made by the function

	// Map from GC safe points to liveness index, generated by
	// liveness analysis.
	livenessMap LivenessMap

	// lineRunStart records the beginning of the current run of instructions
	// within a single block sharing the same line number
	// Used to move statement marks to the beginning of such runs.
	lineRunStart *obj.Prog

	// wasm: The number of values on the WebAssembly stack. This is only used as a safeguard.
	OnWasmStackSkipped int
}

// Prog appends a new Prog.
func (s *SSAGenState) Prog(psess *PackageSession, as obj.As) *obj.Prog {
	p := s.pp.Prog(psess, as)
	if ssa.LosesStmtMark(as) {
		return p
	}

	if s.lineRunStart == nil || s.lineRunStart.Pos.Line() != p.Pos.Line() {
		s.lineRunStart = p
	} else if p.Pos.IsStmt() == src.PosIsStmt {
		s.lineRunStart.Pos = s.lineRunStart.Pos.WithIsStmt()
		p.Pos = p.Pos.WithNotStmt()
	}
	return p
}

// Pc returns the current Prog.
func (s *SSAGenState) Pc() *obj.Prog {
	return s.pp.next
}

// SetPos sets the current source position.
func (s *SSAGenState) SetPos(pos src.XPos) {
	s.pp.pos = pos
}

// Br emits a single branch instruction and returns the instruction.
// Not all architectures need the returned instruction, but otherwise
// the boilerplate is common to all.
func (s *SSAGenState) Br(psess *PackageSession, op obj.As, target *ssa.Block) *obj.Prog {
	p := s.Prog(psess, op)
	p.To.Type = obj.TYPE_BRANCH
	s.Branches = append(s.Branches, Branch{P: p, B: target})
	return p
}

// DebugFriendlySetPos adjusts Pos.IsStmt subject to heuristics
// that reduce "jumpy" line number churn when debugging.
// Spill/fill/copy instructions from the register allocator,
// phi functions, and instructions with a no-pos position
// are examples of instructions that can cause churn.
func (s *SSAGenState) DebugFriendlySetPosFrom(psess *PackageSession, v *ssa.Value) {
	switch v.Op {
	case ssa.OpPhi, ssa.OpCopy, ssa.OpLoadReg, ssa.OpStoreReg:

		s.SetPos(v.Pos.WithNotStmt())
	default:
		p := v.Pos
		if p != psess.src.NoXPos {

			if p.IsStmt() != src.PosIsStmt {
				p = p.WithNotStmt()
			}
			s.SetPos(p)
		}
	}
}

// genssa appends entries to pp for each instruction in f.
func (psess *PackageSession) genssa(f *ssa.Func, pp *Progs) {
	var s SSAGenState

	e := f.Frontend().(*ssafn)

	s.livenessMap = psess.liveness(e, f)

	s.bstart = make([]*obj.Prog, f.NumBlocks())
	s.pp = pp
	var progToValue map[*obj.Prog]*ssa.Value
	var progToBlock map[*obj.Prog]*ssa.Block
	var valueToProgAfter []*obj.Prog // The first Prog following computation of a value v; v is visible at this point.
	var logProgs = e.log
	if logProgs {
		progToValue = make(map[*obj.Prog]*ssa.Value, f.NumValues())
		progToBlock = make(map[*obj.Prog]*ssa.Block, f.NumBlocks())
		f.Logf("genssa %s\n", f.Name)
		progToBlock[s.pp.next] = f.Blocks[0]
	}

	if psess.thearch.Use387 {
		s.SSEto387 = map[int16]int16{}
	}

	s.ScratchFpMem = e.scratchFpMem

	if psess.Ctxt.Flag_locationlists {
		if cap(f.Cache.ValueToProgAfter) < f.NumValues() {
			f.Cache.ValueToProgAfter = make([]*obj.Prog, f.NumValues())
		}
		valueToProgAfter = f.Cache.ValueToProgAfter[:f.NumValues()]
		for i := range valueToProgAfter {
			valueToProgAfter[i] = nil
		}
	}

	firstPos := psess.src.NoXPos
	for _, v := range f.Entry.Values {
		if v.Pos.IsStmt() == src.PosIsStmt {
			firstPos = v.Pos
			v.Pos = firstPos.WithDefaultStmt()
			break
		}
	}

	for i, b := range f.Blocks {
		s.bstart[b.ID] = s.pp.next
		s.pp.nextLive = psess.LivenessInvalid
		s.lineRunStart = nil
		psess.
			thearch.SSAMarkMoves(&s, b)
		for _, v := range b.Values {
			x := s.pp.next
			s.DebugFriendlySetPosFrom(psess, v)

			s.pp.nextLive = s.livenessMap.Get(psess, v)
			switch v.Op {
			case ssa.OpInitMem:

			case ssa.OpArg:

			case ssa.OpSP, ssa.OpSB:

			case ssa.OpSelect0, ssa.OpSelect1:

			case ssa.OpGetG:

			case ssa.OpVarDef, ssa.OpVarLive, ssa.OpKeepAlive:

			case ssa.OpVarKill:

				n := v.Aux.(*Node)
				if n.Name.Needzero() {
					if n.Class() != PAUTO {
						v.Fatalf("zero of variable which isn't PAUTO %v", n)
					}
					if n.Type.Size(psess.types)%int64(psess.Widthptr) != 0 {
						v.Fatalf("zero of variable not a multiple of ptr size %v", n)
					}
					psess.
						thearch.ZeroAuto(s.pp, n)
				}
			case ssa.OpPhi:
				psess.
					CheckLoweredPhi(v)
			case ssa.OpConvert:

				if v.Args[0].Reg(psess.ssa) != v.Reg(psess.ssa) {
					v.Fatalf("OpConvert should be a no-op: %s; %s", v.Args[0].LongString(psess.ssa), v.LongString(psess.ssa))
				}
			default:

				if firstPos != psess.src.NoXPos {
					s.SetPos(firstPos)
					firstPos = psess.src.NoXPos
				}
				psess.
					thearch.SSAGenValue(&s, v)
			}

			if psess.Ctxt.Flag_locationlists {
				valueToProgAfter[v.ID] = s.pp.next
			}

			if logProgs {
				for ; x != s.pp.next; x = x.Link {
					progToValue[x] = v
				}
			}
		}

		// Emit control flow instructions for block
		var next *ssa.Block
		if i < len(f.Blocks)-1 && psess.Debug['N'] == 0 {

			next = f.Blocks[i+1]
		}
		x := s.pp.next
		s.SetPos(b.Pos)
		psess.
			thearch.SSAGenBlock(&s, b, next)
		if logProgs {
			for ; x != s.pp.next; x = x.Link {
				progToBlock[x] = b
			}
		}
	}

	if psess.Ctxt.Flag_locationlists {
		e.curfn.Func.DebugInfo = psess.ssa.BuildFuncDebug(psess.Ctxt, f, psess.Debug_locationlist > 1, psess.stackOffset)
		bstart := s.bstart

		e.curfn.Func.DebugInfo.GetPC = func(b, v ssa.ID) int64 {
			switch v {
			case psess.ssa.BlockStart.ID:
				return bstart[b].Pc
			case psess.ssa.BlockEnd.ID:
				return e.curfn.Func.lsym.Size
			default:
				return valueToProgAfter[v].Pc
			}
		}
	}

	for _, br := range s.Branches {
		br.P.To.Val = s.bstart[br.B.ID]
		if br.P.Pos.IsStmt() != src.PosIsStmt {
			br.P.Pos = br.P.Pos.WithNotStmt()
		}
	}

	if logProgs {
		filename := ""
		for p := pp.Text; p != nil; p = p.Link {
			if p.Pos.IsKnown() && p.InnermostFilename(psess.obj) != filename {
				filename = p.InnermostFilename(psess.obj)
				f.Logf("# %s\n", filename)
			}

			var s string
			if v, ok := progToValue[p]; ok {
				s = v.String()
			} else if b, ok := progToBlock[p]; ok {
				s = b.String()
			} else {
				s = "   "
			}
			f.Logf(" %-6s\t%.5d (%s)\t%s\n", s, p.Pc, p.InnermostLineNumber(), p.InstructionString(psess.obj))
		}
		if f.HTMLWriter != nil {
			// LineHist is defunct now - this code won't do
			// anything.
			// TODO: fix this (ideally without a global variable)
			// saved := pp.Text.Ctxt.LineHist.PrintFilenameOnly
			// pp.Text.Ctxt.LineHist.PrintFilenameOnly = true
			var buf bytes.Buffer
			buf.WriteString("<code>")
			buf.WriteString("<dl class=\"ssa-gen\">")
			filename := ""
			for p := pp.Text; p != nil; p = p.Link {

				if p.Pos.IsKnown() && p.InnermostFilename(psess.obj) != filename {
					filename = p.InnermostFilename(psess.obj)
					buf.WriteString("<dt class=\"ssa-prog-src\"></dt><dd class=\"ssa-prog\">")
					buf.WriteString(html.EscapeString("# " + filename))
					buf.WriteString("</dd>")
				}

				buf.WriteString("<dt class=\"ssa-prog-src\">")
				if v, ok := progToValue[p]; ok {
					buf.WriteString(v.HTML())
				} else if b, ok := progToBlock[p]; ok {
					buf.WriteString("<b>" + b.HTML() + "</b>")
				}
				buf.WriteString("</dt>")
				buf.WriteString("<dd class=\"ssa-prog\">")
				buf.WriteString(fmt.Sprintf("%.5d <span class=\"line-number\">(%s)</span> %s", p.Pc, p.InnermostLineNumberHTML(), html.EscapeString(p.InstructionString(psess.obj))))
				buf.WriteString("</dd>")
			}
			buf.WriteString("</dl>")
			buf.WriteString("</code>")
			f.HTMLWriter.WriteColumn(psess.ssa, "genssa", "genssa", "ssa-prog", buf.String())

		}
	}
	psess.
		defframe(&s, e)
	if psess.Debug['f'] != 0 {
		psess.
			frame(0)
	}

	f.HTMLWriter.Close()
	f.HTMLWriter = nil
}

func (psess *PackageSession) defframe(s *SSAGenState, e *ssafn) {
	pp := s.pp

	frame := psess.Rnd(s.maxarg+e.stksize, int64(psess.Widthreg))
	if psess.thearch.PadFrame != nil {
		frame = psess.thearch.PadFrame(frame)
	}

	pp.Text.To.Type = obj.TYPE_TEXTSIZE
	pp.Text.To.Val = int32(psess.Rnd(e.curfn.Type.ArgWidth(psess.types), int64(psess.Widthreg)))
	pp.Text.To.Offset = frame

	p := pp.Text
	var lo, hi int64

	// Opaque state for backend to use. Current backends use it to
	// keep track of which helper registers have been zeroed.
	var state uint32

	for _, n := range e.curfn.Func.Dcl {
		if !n.Name.Needzero() {
			continue
		}
		if n.Class() != PAUTO {
			psess.
				Fatalf("needzero class %d", n.Class())
		}
		if n.Type.Size(psess.types)%int64(psess.Widthptr) != 0 || n.Xoffset%int64(psess.Widthptr) != 0 || n.Type.Size(psess.types) == 0 {
			psess.
				Fatalf("var %L has size %d offset %d", n, n.Type.Size(psess.types), n.Xoffset)
		}

		if lo != hi && n.Xoffset+n.Type.Size(psess.types) >= lo-int64(2*psess.Widthreg) {

			lo = n.Xoffset
			continue
		}

		p = psess.thearch.ZeroRange(pp, p, frame+lo, hi-lo, &state)

		lo = n.Xoffset
		hi = lo + n.Type.Size(psess.types)
	}
	psess.
		thearch.ZeroRange(pp, p, frame+lo, hi-lo, &state)
}

type FloatingEQNEJump struct {
	Jump  obj.As
	Index int
}

func (s *SSAGenState) oneFPJump(psess *PackageSession, b *ssa.Block, jumps *FloatingEQNEJump) {
	p := s.Prog(psess, jumps.Jump)
	p.To.Type = obj.TYPE_BRANCH
	p.Pos = b.Pos
	to := jumps.Index
	s.Branches = append(s.Branches, Branch{p, b.Succs[to].Block()})
}

func (s *SSAGenState) FPJump(psess *PackageSession, b, next *ssa.Block, jumps *[2][2]FloatingEQNEJump) {
	switch next {
	case b.Succs[0].Block():
		s.oneFPJump(psess, b, &jumps[0][0])
		s.oneFPJump(psess, b, &jumps[0][1])
	case b.Succs[1].Block():
		s.oneFPJump(psess, b, &jumps[1][0])
		s.oneFPJump(psess, b, &jumps[1][1])
	default:
		s.oneFPJump(psess, b, &jumps[1][0])
		s.oneFPJump(psess, b, &jumps[1][1])
		q := s.Prog(psess, obj.AJMP)
		q.Pos = b.Pos
		q.To.Type = obj.TYPE_BRANCH
		s.Branches = append(s.Branches, Branch{q, b.Succs[1].Block()})
	}
}

func (psess *PackageSession) AuxOffset(v *ssa.Value) (offset int64) {
	if v.Aux == nil {
		return 0
	}
	n, ok := v.Aux.(*Node)
	if !ok {
		v.Fatalf("bad aux type in %s\n", v.LongString(psess.ssa))
	}
	if n.Class() == PAUTO {
		return n.Xoffset
	}
	return 0
}

// AddAux adds the offset in the aux fields (AuxInt and Aux) of v to a.
func (psess *PackageSession) AddAux(a *obj.Addr, v *ssa.Value) {
	psess.
		AddAux2(a, v, v.AuxInt)
}
func (psess *PackageSession) AddAux2(a *obj.Addr, v *ssa.Value, offset int64) {
	if a.Type != obj.TYPE_MEM && a.Type != obj.TYPE_ADDR {
		v.Fatalf("bad AddAux addr %v", a)
	}

	a.Offset += offset

	if v.Aux == nil {
		return
	}

	switch n := v.Aux.(type) {
	case *obj.LSym:
		a.Name = obj.NAME_EXTERN
		a.Sym = n
	case *Node:
		if n.Class() == PPARAM || n.Class() == PPARAMOUT {
			a.Name = obj.NAME_PARAM
			a.Sym = n.Orig.Sym.Linksym(psess.types)
			a.Offset += n.Xoffset
			break
		}
		a.Name = obj.NAME_AUTO
		a.Sym = n.Sym.Linksym(psess.types)
		a.Offset += n.Xoffset
	default:
		v.Fatalf("aux in %s not implemented %#v", v, v.Aux)
	}
}

// extendIndex extends v to a full int width.
// panic using the given function if v does not fit in an int (only on 32-bit archs).
func (s *state) extendIndex(psess *PackageSession, v *ssa.Value, panicfn *obj.LSym) *ssa.Value {
	size := v.Type.Size(psess.types)
	if size == s.config.PtrSize {
		return v
	}
	if size > s.config.PtrSize {

		if psess.Debug['B'] == 0 {
			hi := s.newValue1(ssa.OpInt64Hi, psess.types.Types[TUINT32], v)
			cmp := s.newValue2(ssa.OpEq32, psess.types.Types[TBOOL], hi, s.constInt32(psess, psess.types.Types[TUINT32], 0))
			s.check(psess, cmp, panicfn)
		}
		return s.newValue1(ssa.OpTrunc64to32, psess.types.Types[TINT], v)
	}

	// Extend value to the required size
	var op ssa.Op
	if v.Type.IsSigned() {
		switch 10*size + s.config.PtrSize {
		case 14:
			op = ssa.OpSignExt8to32
		case 18:
			op = ssa.OpSignExt8to64
		case 24:
			op = ssa.OpSignExt16to32
		case 28:
			op = ssa.OpSignExt16to64
		case 48:
			op = ssa.OpSignExt32to64
		default:
			s.Fatalf("bad signed index extension %s", v.Type)
		}
	} else {
		switch 10*size + s.config.PtrSize {
		case 14:
			op = ssa.OpZeroExt8to32
		case 18:
			op = ssa.OpZeroExt8to64
		case 24:
			op = ssa.OpZeroExt16to32
		case 28:
			op = ssa.OpZeroExt16to64
		case 48:
			op = ssa.OpZeroExt32to64
		default:
			s.Fatalf("bad unsigned index extension %s", v.Type)
		}
	}
	return s.newValue1(op, psess.types.Types[TINT], v)
}

// CheckLoweredPhi checks that regalloc and stackalloc correctly handled phi values.
// Called during ssaGenValue.
func (psess *PackageSession) CheckLoweredPhi(v *ssa.Value) {
	if v.Op != ssa.OpPhi {
		v.Fatalf("CheckLoweredPhi called with non-phi value: %v", v.LongString(psess.ssa))
	}
	if v.Type.IsMemory(psess.types) {
		return
	}
	f := v.Block.Func
	loc := f.RegAlloc[v.ID]
	for _, a := range v.Args {
		if aloc := f.RegAlloc[a.ID]; aloc != loc {
			v.Fatalf("phi arg at different location than phi: %v @ %s, but arg %v @ %s\n%s\n", v, loc, a, aloc, v.Block.Func)
		}
	}
}

// CheckLoweredGetClosurePtr checks that v is the first instruction in the function's entry block.
// The output of LoweredGetClosurePtr is generally hardwired to the correct register.
// That register contains the closure pointer on closure entry.
func (psess *PackageSession) CheckLoweredGetClosurePtr(v *ssa.Value) {
	entry := v.Block.Func.Entry
	if entry != v.Block || entry.Values[0] != v {
		psess.
			Fatalf("in %s, badly placed LoweredGetClosurePtr: %v %v", v.Block.Func.Name, v.Block, v)
	}
}

// AutoVar returns a *Node and int64 representing the auto variable and offset within it
// where v should be spilled.
func (psess *PackageSession) AutoVar(v *ssa.Value) (*Node, int64) {
	loc := v.Block.Func.RegAlloc[v.ID].(ssa.LocalSlot)
	if v.Type.Size(psess.types) > loc.Type.Size(psess.types) {
		v.Fatalf("spill/restore type %s doesn't fit in slot type %s", v.Type, loc.Type)
	}
	return loc.N.(*Node), loc.Off
}

func (psess *PackageSession) AddrAuto(a *obj.Addr, v *ssa.Value) {
	n, off := psess.AutoVar(v)
	a.Type = obj.TYPE_MEM
	a.Sym = n.Sym.Linksym(psess.types)
	a.Reg = int16(psess.thearch.REGSP)
	a.Offset = n.Xoffset + off
	if n.Class() == PPARAM || n.Class() == PPARAMOUT {
		a.Name = obj.NAME_PARAM
	} else {
		a.Name = obj.NAME_AUTO
	}
}

func (s *SSAGenState) AddrScratch(psess *PackageSession, a *obj.Addr) {
	if s.ScratchFpMem == nil {
		panic("no scratch memory available; forgot to declare usesScratch for Op?")
	}
	a.Type = obj.TYPE_MEM
	a.Name = obj.NAME_AUTO
	a.Sym = s.ScratchFpMem.Sym.Linksym(psess.types)
	a.Reg = int16(psess.thearch.REGSP)
	a.Offset = s.ScratchFpMem.Xoffset
}

// Call returns a new CALL instruction for the SSA value v.
// It uses PrepareCall to prepare the call.
func (s *SSAGenState) Call(psess *PackageSession, v *ssa.Value) *obj.Prog {
	s.PrepareCall(psess, v)

	p := s.Prog(psess, obj.ACALL)
	if sym, ok := v.Aux.(*obj.LSym); ok {
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = sym
	} else {

		switch psess.thearch.LinkArch.Family {
		case sys.AMD64, sys.I386, sys.PPC64, sys.S390X, sys.Wasm:
			p.To.Type = obj.TYPE_REG
		case sys.ARM, sys.ARM64, sys.MIPS, sys.MIPS64:
			p.To.Type = obj.TYPE_MEM
		default:
			psess.
				Fatalf("unknown indirect call family")
		}
		p.To.Reg = v.Args[0].Reg(psess.ssa)
	}
	return p
}

// PrepareCall prepares to emit a CALL instruction for v and does call-related bookkeeping.
// It must be called immediately before emitting the actual CALL instruction,
// since it emits PCDATA for the stack map at the call (calls are safe points).
func (s *SSAGenState) PrepareCall(psess *PackageSession, v *ssa.Value) {
	idx := s.livenessMap.Get(psess, v)
	if !idx.Valid() {

		if sym, _ := v.Aux.(*obj.LSym); !(sym == psess.typedmemclr || sym == psess.typedmemmove) {
			psess.
				Fatalf("missing stack map index for %v", v.LongString(psess.ssa))
		}
	}

	if sym, _ := v.Aux.(*obj.LSym); sym == psess.Deferreturn {
		psess.
			thearch.Ginsnop(s.pp)
	}

	if sym, ok := v.Aux.(*obj.LSym); ok {

		if psess.nowritebarrierrecCheck != nil {
			psess.
				nowritebarrierrecCheck.recordCall(psess, s.pp.curfn, sym, v.Pos)
		}
	}

	if s.maxarg < v.AuxInt {
		s.maxarg = v.AuxInt
	}
}

// fieldIdx finds the index of the field referred to by the ODOT node n.
func (psess *PackageSession) fieldIdx(n *Node) int {
	t := n.Left.Type
	f := n.Sym
	if !t.IsStruct() {
		panic("ODOT's LHS is not a struct")
	}

	var i int
	for _, t1 := range t.Fields(psess.types).Slice() {
		if t1.Sym != f {
			i++
			continue
		}
		if t1.Offset != n.Xoffset {
			panic("field offset doesn't match")
		}
		return i
	}
	panic(fmt.Sprintf("can't find field in expr %v\n", n))

}

// ssafn holds frontend information about a function that the backend is processing.
// It also exports a bunch of compiler services for the ssa backend.
type ssafn struct {
	curfn        *Node
	strings      map[string]interface{} // map from constant string to data symbols
	scratchFpMem *Node                  // temp for floating point register / memory moves on some architectures
	stksize      int64                  // stack size for current frame
	stkptrsize   int64                  // prefix of stack containing pointers
	log          bool
}

// StringData returns a symbol (a *types.Sym wrapped in an interface) which
// is the data component of a global string constant containing s.
func (e *ssafn) StringData(psess *PackageSession, s string) interface{} {
	if aux, ok := e.strings[s]; ok {
		return aux
	}
	if e.strings == nil {
		e.strings = make(map[string]interface{})
	}
	data := psess.stringsym(e.curfn.Pos, s)
	e.strings[s] = data
	return data
}

func (e *ssafn) Auto(psess *PackageSession, pos src.XPos, t *types.Type) ssa.GCNode {
	n := psess.tempAt(pos, e.curfn, t)
	return n
}

func (e *ssafn) SplitString(psess *PackageSession, name ssa.LocalSlot) (ssa.LocalSlot, ssa.LocalSlot) {
	n := name.N.(*Node)
	ptrType := psess.types.NewPtr(psess.types.Types[TUINT8])
	lenType := psess.types.Types[TINT]
	if n.Class() == PAUTO && !n.Addrtaken() {

		p := e.splitSlot(psess, &name, ".ptr", 0, ptrType)
		l := e.splitSlot(psess, &name, ".len", ptrType.Size(psess.types), lenType)
		return p, l
	}

	return ssa.LocalSlot{N: n, Type: ptrType, Off: name.Off}, ssa.LocalSlot{N: n, Type: lenType, Off: name.Off + int64(psess.Widthptr)}
}

func (e *ssafn) SplitInterface(psess *PackageSession, name ssa.LocalSlot) (ssa.LocalSlot, ssa.LocalSlot) {
	n := name.N.(*Node)
	u := psess.types.Types[TUINTPTR]
	t := psess.types.NewPtr(psess.types.Types[TUINT8])
	if n.Class() == PAUTO && !n.Addrtaken() {

		f := ".itab"
		if n.Type.IsEmptyInterface(psess.types) {
			f = ".type"
		}
		c := e.splitSlot(psess, &name, f, 0, u)
		d := e.splitSlot(psess, &name, ".data", u.Size(psess.types), t)
		return c, d
	}

	return ssa.LocalSlot{N: n, Type: u, Off: name.Off}, ssa.LocalSlot{N: n, Type: t, Off: name.Off + int64(psess.Widthptr)}
}

func (e *ssafn) SplitSlice(psess *PackageSession, name ssa.LocalSlot) (ssa.LocalSlot, ssa.LocalSlot, ssa.LocalSlot) {
	n := name.N.(*Node)
	ptrType := psess.types.NewPtr(name.Type.Elem(psess.types))
	lenType := psess.types.Types[TINT]
	if n.Class() == PAUTO && !n.Addrtaken() {

		p := e.splitSlot(psess, &name, ".ptr", 0, ptrType)
		l := e.splitSlot(psess, &name, ".len", ptrType.Size(psess.types), lenType)
		c := e.splitSlot(psess, &name, ".cap", ptrType.Size(psess.types)+lenType.Size(psess.types), lenType)
		return p, l, c
	}

	return ssa.LocalSlot{N: n, Type: ptrType, Off: name.Off},
		ssa.LocalSlot{N: n, Type: lenType, Off: name.Off + int64(psess.Widthptr)},
		ssa.LocalSlot{N: n, Type: lenType, Off: name.Off + int64(2*psess.Widthptr)}
}

func (e *ssafn) SplitComplex(psess *PackageSession, name ssa.LocalSlot) (ssa.LocalSlot, ssa.LocalSlot) {
	n := name.N.(*Node)
	s := name.Type.Size(psess.types) / 2
	var t *types.Type
	if s == 8 {
		t = psess.types.Types[TFLOAT64]
	} else {
		t = psess.types.Types[TFLOAT32]
	}
	if n.Class() == PAUTO && !n.Addrtaken() {

		r := e.splitSlot(psess, &name, ".real", 0, t)
		i := e.splitSlot(psess, &name, ".imag", t.Size(psess.types), t)
		return r, i
	}

	return ssa.LocalSlot{N: n, Type: t, Off: name.Off}, ssa.LocalSlot{N: n, Type: t, Off: name.Off + s}
}

func (e *ssafn) SplitInt64(psess *PackageSession, name ssa.LocalSlot) (ssa.LocalSlot, ssa.LocalSlot) {
	n := name.N.(*Node)
	var t *types.Type
	if name.Type.IsSigned() {
		t = psess.types.Types[TINT32]
	} else {
		t = psess.types.Types[TUINT32]
	}
	if n.Class() == PAUTO && !n.Addrtaken() {

		if psess.thearch.LinkArch.ByteOrder == binary.BigEndian {
			return e.splitSlot(psess, &name, ".hi", 0, t), e.splitSlot(psess, &name, ".lo", t.Size(psess.types), psess.types.Types[TUINT32])
		}
		return e.splitSlot(psess, &name, ".hi", t.Size(psess.types), t), e.splitSlot(psess, &name, ".lo", 0, psess.types.Types[TUINT32])
	}

	if psess.thearch.LinkArch.ByteOrder == binary.BigEndian {
		return ssa.LocalSlot{N: n, Type: t, Off: name.Off}, ssa.LocalSlot{N: n, Type: psess.types.Types[TUINT32], Off: name.Off + 4}
	}
	return ssa.LocalSlot{N: n, Type: t, Off: name.Off + 4}, ssa.LocalSlot{N: n, Type: psess.types.Types[TUINT32], Off: name.Off}
}

func (e *ssafn) SplitStruct(psess *PackageSession, name ssa.LocalSlot, i int) ssa.LocalSlot {
	n := name.N.(*Node)
	st := name.Type
	ft := st.FieldType(psess.types, i)
	var offset int64
	for f := 0; f < i; f++ {
		offset += st.FieldType(psess.types, f).Size(psess.types)
	}
	if n.Class() == PAUTO && !n.Addrtaken() {

		return e.splitSlot(psess, &name, "."+st.FieldName(psess.types, i), offset, ft)
	}
	return ssa.LocalSlot{N: n, Type: ft, Off: name.Off + st.FieldOff(psess.types, i)}
}

func (e *ssafn) SplitArray(psess *PackageSession, name ssa.LocalSlot) ssa.LocalSlot {
	n := name.N.(*Node)
	at := name.Type
	if at.NumElem(psess.types) != 1 {
		psess.
			Fatalf("bad array size")
	}
	et := at.Elem(psess.types)
	if n.Class() == PAUTO && !n.Addrtaken() {
		return e.splitSlot(psess, &name, "[0]", 0, et)
	}
	return ssa.LocalSlot{N: n, Type: et, Off: name.Off}
}

func (e *ssafn) DerefItab(psess *PackageSession, it *obj.LSym, offset int64) *obj.LSym {
	return psess.itabsym(it, offset)
}

// splitSlot returns a slot representing the data of parent starting at offset.
func (e *ssafn) splitSlot(psess *PackageSession, parent *ssa.LocalSlot, suffix string, offset int64, t *types.Type) ssa.LocalSlot {
	s := &types.Sym{Name: parent.N.(*Node).Sym.Name + suffix, Pkg: psess.localpkg}

	n := &Node{
		Name: new(Name),
		Op:   ONAME,
		Pos:  parent.N.(*Node).Pos,
	}
	n.Orig = n

	s.Def = asTypesNode(n)
	asNode(s.Def).Name.SetUsed(true)
	n.Sym = s
	n.Type = t
	n.SetClass(PAUTO)
	n.SetAddable(true)
	n.Esc = EscNever
	n.Name.Curfn = e.curfn
	e.curfn.Func.Dcl = append(e.curfn.Func.Dcl, n)
	psess.
		dowidth(t)
	return ssa.LocalSlot{N: n, Type: t, Off: 0, SplitOf: parent, SplitOffset: offset}
}

func (e *ssafn) CanSSA(psess *PackageSession, t *types.Type) bool {
	return psess.canSSAType(t)
}

func (e *ssafn) Line(psess *PackageSession, pos src.XPos) string {
	return psess.linestr(pos)
}

// Log logs a message from the compiler.
func (e *ssafn) Logf(msg string, args ...interface{}) {
	if e.log {
		fmt.Printf(msg, args...)
	}
}

func (e *ssafn) Log() bool {
	return e.log
}

// Fatal reports a compiler error and exits.
func (e *ssafn) Fatalf(psess *PackageSession, pos src.XPos, msg string, args ...interface{}) {
	psess.
		lineno = pos
	psess.
		Fatalf(msg, args...)
}

// Warnl reports a "warning", which is usually flag-triggered
// logging output for the benefit of tests.
func (e *ssafn) Warnl(psess *PackageSession, pos src.XPos, fmt_ string, args ...interface{}) {
	psess.
		Warnl(pos, fmt_, args...)
}

func (e *ssafn) Debug_checknil(psess *PackageSession) bool {
	return psess.Debug_checknil != 0
}

func (e *ssafn) UseWriteBarrier(psess *PackageSession) bool {
	return psess.use_writebarrier
}

func (e *ssafn) Syslook(psess *PackageSession, name string) *obj.LSym {
	switch name {
	case "goschedguarded":
		return psess.goschedguarded
	case "writeBarrier":
		return psess.writeBarrier
	case "gcWriteBarrier":
		return psess.gcWriteBarrier
	case "typedmemmove":
		return psess.typedmemmove
	case "typedmemclr":
		return psess.typedmemclr
	}
	psess.
		Fatalf("unknown Syslook func %v", name)
	return nil
}

func (e *ssafn) SetWBPos(psess *PackageSession, pos src.XPos) {
	e.curfn.Func.setWBPos(psess, pos)
}

func (n *Node) Typ() *types.Type {
	return n.Type
}
func (n *Node) StorageClass(psess *PackageSession) ssa.StorageClass {
	switch n.Class() {
	case PPARAM:
		return ssa.ClassParam
	case PPARAMOUT:
		return ssa.ClassParamOut
	case PAUTO:
		return ssa.ClassAuto
	default:
		psess.
			Fatalf("untranslateable storage class for %v: %s", n, n.Class())
		return 0
	}
}

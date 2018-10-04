// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

func (pstate *PackageState) initssaconfig() {
	types_ := pstate.ssa.NewTypes()

	if pstate.thearch.SoftFloat {
		pstate.softfloatInit()
	}

	// Generate a few pointer types that are uncommon in the frontend but common in the backend.
	// Caching is disabled in the backend, so generating these here avoids allocations.
	_ = pstate.types.NewPtr(pstate.types.Types[TINTER])                                           // *interface{}
	_ = pstate.types.NewPtr(pstate.types.NewPtr(pstate.types.Types[TSTRING]))                     // **string
	_ = pstate.types.NewPtr(pstate.types.NewPtr(pstate.types.Idealstring))                        // **string
	_ = pstate.types.NewPtr(pstate.types.NewSlice(pstate.types.Types[TINTER]))                    // *[]interface{}
	_ = pstate.types.NewPtr(pstate.types.NewPtr(pstate.types.Bytetype))                           // **byte
	_ = pstate.types.NewPtr(pstate.types.NewSlice(pstate.types.Bytetype))                         // *[]byte
	_ = pstate.types.NewPtr(pstate.types.NewSlice(pstate.types.Types[TSTRING]))                   // *[]string
	_ = pstate.types.NewPtr(pstate.types.NewSlice(pstate.types.Idealstring))                      // *[]string
	_ = pstate.types.NewPtr(pstate.types.NewPtr(pstate.types.NewPtr(pstate.types.Types[TUINT8]))) // ***uint8
	_ = pstate.types.NewPtr(pstate.types.Types[TINT16])                                           // *int16
	_ = pstate.types.NewPtr(pstate.types.Types[TINT64])                                           // *int64
	_ = pstate.types.NewPtr(pstate.types.Errortype)                                               // *error
	pstate.types.NewPtrCacheEnabled = false
	pstate.ssaConfig = pstate.ssa.NewConfig(pstate.thearch.LinkArch.Name, *types_, pstate.Ctxt, pstate.Debug['N'] == 0)
	if pstate.thearch.LinkArch.Name == "386" {
		pstate.ssaConfig.Set387(pstate.thearch.Use387)
	}
	pstate.ssaConfig.SoftFloat = pstate.thearch.SoftFloat
	pstate.ssaCaches = make([]ssa.Cache, pstate.nBackendWorkers)

	// Set up some runtime functions we'll need to call.
	pstate.assertE2I = pstate.sysfunc("assertE2I")
	pstate.assertE2I2 = pstate.sysfunc("assertE2I2")
	pstate.assertI2I = pstate.sysfunc("assertI2I")
	pstate.assertI2I2 = pstate.sysfunc("assertI2I2")
	pstate.Deferproc = pstate.sysfunc("deferproc")
	pstate.Deferreturn = pstate.sysfunc("deferreturn")
	pstate.Duffcopy = pstate.sysfunc("duffcopy")
	pstate.Duffzero = pstate.sysfunc("duffzero")
	pstate.gcWriteBarrier = pstate.sysfunc("gcWriteBarrier")
	pstate.goschedguarded = pstate.sysfunc("goschedguarded")
	pstate.growslice = pstate.sysfunc("growslice")
	pstate.msanread = pstate.sysfunc("msanread")
	pstate.msanwrite = pstate.sysfunc("msanwrite")
	pstate.Newproc = pstate.sysfunc("newproc")
	pstate.panicdivide = pstate.sysfunc("panicdivide")
	pstate.panicdottypeE = pstate.sysfunc("panicdottypeE")
	pstate.panicdottypeI = pstate.sysfunc("panicdottypeI")
	pstate.panicindex = pstate.sysfunc("panicindex")
	pstate.panicnildottype = pstate.sysfunc("panicnildottype")
	pstate.panicslice = pstate.sysfunc("panicslice")
	pstate.raceread = pstate.sysfunc("raceread")
	pstate.racereadrange = pstate.sysfunc("racereadrange")
	pstate.racewrite = pstate.sysfunc("racewrite")
	pstate.racewriterange = pstate.sysfunc("racewriterange")
	pstate.supportPopcnt = pstate.sysfunc("support_popcnt")
	pstate.supportSSE41 = pstate.sysfunc("support_sse41")
	pstate.arm64SupportAtomics = pstate.sysfunc("arm64_support_atomics")
	pstate.typedmemclr = pstate.sysfunc("typedmemclr")
	pstate.typedmemmove = pstate.sysfunc("typedmemmove")
	pstate.Udiv = pstate.sysfunc("udiv")
	pstate.writeBarrier = pstate.sysfunc("writeBarrier")

	// GO386=387 runtime functions
	pstate.ControlWord64trunc = pstate.sysfunc("controlWord64trunc")
	pstate.ControlWord32 = pstate.sysfunc("controlWord32")

	// Wasm
	pstate.WasmMove = pstate.sysfunc("wasmMove")
	pstate.WasmZero = pstate.sysfunc("wasmZero")
	pstate.WasmDiv = pstate.sysfunc("wasmDiv")
	pstate.WasmTruncS = pstate.sysfunc("wasmTruncS")
	pstate.WasmTruncU = pstate.sysfunc("wasmTruncU")
	pstate.SigPanic = pstate.sysfunc("sigpanic")
}

// buildssa builds an SSA function for fn.
// worker indicates which of the backend workers is doing the processing.
func (pstate *PackageState) buildssa(fn *Node, worker int) *ssa.Func {
	name := fn.funcname()
	printssa := name == os.Getenv("GOSSAFUNC")
	if printssa {
		fmt.Println("generating SSA for", name)
		dumplist("buildssa-enter", fn.Func.Enter)
		dumplist("buildssa-body", fn.Nbody)
		dumplist("buildssa-exit", fn.Func.Exit)
	}

	var s state
	s.pushLine(pstate, fn.Pos)
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
	s.config = pstate.ssaConfig
	s.f.Type = fn.Type
	s.f.Config = pstate.ssaConfig
	s.f.Cache = &pstate.ssaCaches[worker]
	s.f.Cache.Reset()
	s.f.DebugTest = s.f.DebugHashMatch("GOSSAHASH", name)
	s.f.Name = name
	if fn.Func.Pragma&Nosplit != 0 {
		s.f.NoSplit = true
	}
	s.panics = map[funcLine]*ssa.Block{}
	s.softFloat = s.config.SoftFloat

	if name == os.Getenv("GOSSAFUNC") {
		s.f.HTMLWriter = pstate.ssa.NewHTMLWriter("ssa.html", s.f.Frontend(), name)
		// TODO: generate and print a mapping from nodes to values and blocks
	}

	// Allocate starting block
	s.f.Entry = s.f.NewBlock(ssa.BlockPlain)

	// Allocate starting values
	s.labels = map[string]*ssaLabel{}
	s.labeledNodes = map[*Node]*ssaLabel{}
	s.fwdVars = map[*Node]*ssa.Value{}
	s.startmem = s.entryNewValue0(pstate, ssa.OpInitMem, pstate.types.TypeMem)
	s.sp = s.entryNewValue0(pstate, ssa.OpSP, pstate.types.Types[TUINTPTR]) // TODO: use generic pointer type (unsafe.Pointer?) instead
	s.sb = s.entryNewValue0(pstate, ssa.OpSB, pstate.types.Types[TUINTPTR])

	s.startBlock(s.f.Entry)
	s.vars[&pstate.memVar] = s.startmem

	// Generate addresses of local declarations
	s.decladdrs = map[*Node]*ssa.Value{}
	for _, n := range fn.Func.Dcl {
		switch n.Class() {
		case PPARAM, PPARAMOUT:
			s.decladdrs[n] = s.entryNewValue1A(pstate, ssa.OpAddr, pstate.types.NewPtr(n.Type), n, s.sp)
			if n.Class() == PPARAMOUT && s.canSSA(pstate, n) {
				// Save ssa-able PPARAMOUT variables so we can
				// store them back to the stack at the end of
				// the function.
				s.returns = append(s.returns, n)
			}
		case PAUTO:
		// processed at each use, to prevent Addr coming
		// before the decl.
		case PAUTOHEAP:
		// moved to heap - already handled by frontend
		case PFUNC:
		// local function - already handled by frontend
		default:
			s.Fatalf("local variable with class %v unimplemented", n.Class())
		}
	}

	// Populate SSAable arguments.
	for _, n := range fn.Func.Dcl {
		if n.Class() == PPARAM && s.canSSA(pstate, n) {
			s.vars[n] = s.newValue0A(ssa.OpArg, n.Type, n)
		}
	}

	// Convert the AST-based IR to the SSA-based IR
	s.stmtList(pstate, fn.Func.Enter)
	s.stmtList(pstate, fn.Nbody)

	// fallthrough to exit
	if s.curBlock != nil {
		s.pushLine(pstate, fn.Func.Endlineno)
		s.exit(pstate)
		s.popLine()
	}

	for _, b := range s.f.Blocks {
		if b.Pos != pstate.src.NoXPos {
			s.updateUnsetPredPos(pstate, b)
		}
	}

	s.insertPhis(pstate)

	// Main call to ssa package to compile function
	pstate.ssa.Compile(s.f)
	return s.f
}

// updateUnsetPredPos propagates the earliest-value position information for b
// towards all of b's predecessors that need a position, and recurs on that
// predecessor if its position is updated. B should have a non-empty position.
func (s *state) updateUnsetPredPos(pstate *PackageState, b *ssa.Block) {
	if b.Pos == pstate.src.NoXPos {
		s.Fatalf("Block %s should have a position", b)
	}
	bestPos := pstate.src.NoXPos
	for _, e := range b.Preds {
		p := e.Block()
		if !p.LackingPos(pstate.ssa) {
			continue
		}
		if bestPos == pstate.src.NoXPos {
			bestPos = b.Pos
			for _, v := range b.Values {
				if v.LackingPos(pstate.ssa) {
					continue
				}
				if v.Pos != pstate.src.NoXPos {
					// Assume values are still in roughly textual order;
					// TODO: could also seek minimum position?
					bestPos = v.Pos
					break
				}
			}
		}
		p.Pos = bestPos
		s.updateUnsetPredPos(pstate, p) // We do not expect long chains of these, thus recursion is okay.
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
func (s *state) endBlock(pstate *PackageState) *ssa.Block {
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
	if b.LackingPos(pstate.ssa) {
		// Empty plain blocks get the line of their successor (handled after all blocks created),
		// except for increment blocks in For statements (handled in ssa conversion of OFOR),
		// and for blocks ending in GOTO/BREAK/CONTINUE.
		b.Pos = pstate.src.NoXPos
	} else {
		b.Pos = s.lastPos
	}
	return b
}

// pushLine pushes a line number on the line number stack.
func (s *state) pushLine(pstate *PackageState, line src.XPos) {
	if !line.IsKnown() {
		// the frontend may emit node with line number missing,
		// use the parent line number in this case.
		line = s.peekPos()
		if pstate.Debug['K'] != 0 {
			pstate.Warn("buildssa: unknown position (line 0)")
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
func (s *state) entryNewValue0(pstate *PackageState, op ssa.Op, t *types.Type) *ssa.Value {
	return s.f.Entry.NewValue0(pstate.src.NoXPos, op, t)
}

// entryNewValue0A adds a new value with no arguments and an aux value to the entry block.
func (s *state) entryNewValue0A(pstate *PackageState, op ssa.Op, t *types.Type, aux interface{}) *ssa.Value {
	return s.f.Entry.NewValue0A(pstate.src.NoXPos, op, t, aux)
}

// entryNewValue1 adds a new value with one argument to the entry block.
func (s *state) entryNewValue1(pstate *PackageState, op ssa.Op, t *types.Type, arg *ssa.Value) *ssa.Value {
	return s.f.Entry.NewValue1(pstate.src.NoXPos, op, t, arg)
}

// entryNewValue1 adds a new value with one argument and an auxint value to the entry block.
func (s *state) entryNewValue1I(pstate *PackageState, op ssa.Op, t *types.Type, auxint int64, arg *ssa.Value) *ssa.Value {
	return s.f.Entry.NewValue1I(pstate.src.NoXPos, op, t, auxint, arg)
}

// entryNewValue1A adds a new value with one argument and an aux value to the entry block.
func (s *state) entryNewValue1A(pstate *PackageState, op ssa.Op, t *types.Type, aux interface{}, arg *ssa.Value) *ssa.Value {
	return s.f.Entry.NewValue1A(pstate.src.NoXPos, op, t, aux, arg)
}

// entryNewValue2 adds a new value with two arguments to the entry block.
func (s *state) entryNewValue2(pstate *PackageState, op ssa.Op, t *types.Type, arg0, arg1 *ssa.Value) *ssa.Value {
	return s.f.Entry.NewValue2(pstate.src.NoXPos, op, t, arg0, arg1)
}

// const* routines add a new const value to the entry block.
func (s *state) constSlice(pstate *PackageState, t *types.Type) *ssa.Value {
	return s.f.ConstSlice(pstate.ssa, t)
}
func (s *state) constInterface(pstate *PackageState, t *types.Type) *ssa.Value {
	return s.f.ConstInterface(pstate.ssa, t)
}
func (s *state) constNil(pstate *PackageState, t *types.Type) *ssa.Value {
	return s.f.ConstNil(pstate.ssa, t)
}
func (s *state) constEmptyString(pstate *PackageState, t *types.Type) *ssa.Value {
	return s.f.ConstEmptyString(pstate.ssa, t)
}
func (s *state) constBool(pstate *PackageState, c bool) *ssa.Value {
	return s.f.ConstBool(pstate.ssa, pstate.types.Types[TBOOL], c)
}
func (s *state) constInt8(pstate *PackageState, t *types.Type, c int8) *ssa.Value {
	return s.f.ConstInt8(pstate.ssa, t, c)
}
func (s *state) constInt16(pstate *PackageState, t *types.Type, c int16) *ssa.Value {
	return s.f.ConstInt16(pstate.ssa, t, c)
}
func (s *state) constInt32(pstate *PackageState, t *types.Type, c int32) *ssa.Value {
	return s.f.ConstInt32(pstate.ssa, t, c)
}
func (s *state) constInt64(pstate *PackageState, t *types.Type, c int64) *ssa.Value {
	return s.f.ConstInt64(pstate.ssa, t, c)
}
func (s *state) constFloat32(pstate *PackageState, t *types.Type, c float64) *ssa.Value {
	return s.f.ConstFloat32(pstate.ssa, t, c)
}
func (s *state) constFloat64(pstate *PackageState, t *types.Type, c float64) *ssa.Value {
	return s.f.ConstFloat64(pstate.ssa, t, c)
}
func (s *state) constInt(pstate *PackageState, t *types.Type, c int64) *ssa.Value {
	if s.config.PtrSize == 8 {
		return s.constInt64(pstate, t, c)
	}
	if int64(int32(c)) != c {
		s.Fatalf("integer constant too big %d", c)
	}
	return s.constInt32(pstate, t, int32(c))
}
func (s *state) constOffPtrSP(pstate *PackageState, t *types.Type, c int64) *ssa.Value {
	return s.f.ConstOffPtrSP(pstate.ssa, t, c, s.sp)
}

// newValueOrSfCall* are wrappers around newValue*, which may create a call to a
// soft-float runtime function instead (when emitting soft-float code).
func (s *state) newValueOrSfCall1(pstate *PackageState, op ssa.Op, t *types.Type, arg *ssa.Value) *ssa.Value {
	if s.softFloat {
		if c, ok := s.sfcall(pstate, op, arg); ok {
			return c
		}
	}
	return s.newValue1(op, t, arg)
}
func (s *state) newValueOrSfCall2(pstate *PackageState, op ssa.Op, t *types.Type, arg0, arg1 *ssa.Value) *ssa.Value {
	if s.softFloat {
		if c, ok := s.sfcall(pstate, op, arg0, arg1); ok {
			return c
		}
	}
	return s.newValue2(op, t, arg0, arg1)
}

func (s *state) instrument(pstate *PackageState, t *types.Type, addr *ssa.Value, wr bool) {
	if !s.curfn.Func.InstrumentBody() {
		return
	}

	w := t.Size(pstate.types)
	if w == 0 {
		return // can't race on zero-sized things
	}

	if ssa.IsSanitizerSafeAddr(addr) {
		return
	}

	var fn *obj.LSym
	needWidth := false

	if pstate.flag_msan {
		fn = pstate.msanread
		if wr {
			fn = pstate.msanwrite
		}
		needWidth = true
	} else if pstate.flag_race && t.NumComponents(pstate.types, types.CountBlankFields) > 1 {
		// for composite objects we have to write every address
		// because a write might happen to any subobject.
		// composites with only one element don't have subobjects, though.
		fn = pstate.racereadrange
		if wr {
			fn = pstate.racewriterange
		}
		needWidth = true
	} else if pstate.flag_race {
		// for non-composite objects we can write just the start
		// address, as any write must write the first byte.
		fn = pstate.raceread
		if wr {
			fn = pstate.racewrite
		}
	} else {
		panic("unreachable")
	}

	args := []*ssa.Value{addr}
	if needWidth {
		args = append(args, s.constInt(pstate, pstate.types.Types[TUINTPTR], w))
	}
	s.rtcall(pstate, fn, true, nil, args...)
}

func (s *state) load(pstate *PackageState, t *types.Type, src *ssa.Value) *ssa.Value {
	s.instrument(pstate, t, src, false)
	return s.rawLoad(pstate, t, src)
}

func (s *state) rawLoad(pstate *PackageState, t *types.Type, src *ssa.Value) *ssa.Value {
	return s.newValue2(ssa.OpLoad, t, src, s.mem(pstate))
}

func (s *state) store(pstate *PackageState, t *types.Type, dst, val *ssa.Value) {
	s.vars[&pstate.memVar] = s.newValue3A(ssa.OpStore, pstate.types.TypeMem, t, dst, val, s.mem(pstate))
}

func (s *state) zero(pstate *PackageState, t *types.Type, dst *ssa.Value) {
	s.instrument(pstate, t, dst, true)
	store := s.newValue2I(ssa.OpZero, pstate.types.TypeMem, t.Size(pstate.types), dst, s.mem(pstate))
	store.Aux = t
	s.vars[&pstate.memVar] = store
}

func (s *state) move(pstate *PackageState, t *types.Type, dst, src *ssa.Value) {
	s.instrument(pstate, t, src, false)
	s.instrument(pstate, t, dst, true)
	store := s.newValue3I(ssa.OpMove, pstate.types.TypeMem, t.Size(pstate.types), dst, src, s.mem(pstate))
	store.Aux = t
	s.vars[&pstate.memVar] = store
}

// stmtList converts the statement list n to SSA and adds it to s.
func (s *state) stmtList(pstate *PackageState, l Nodes) {
	for _, n := range l.Slice() {
		s.stmt(pstate, n)
	}
}

// stmt converts the statement n to SSA and adds it to s.
func (s *state) stmt(pstate *PackageState, n *Node) {
	if !(n.Op == OVARKILL || n.Op == OVARLIVE) {
		// OVARKILL and OVARLIVE are invisible to the programmer, so we don't use their line numbers to avoid confusion in debugging.
		s.pushLine(pstate, n.Pos)
		defer s.popLine()
	}

	// If s.curBlock is nil, and n isn't a label (which might have an associated goto somewhere),
	// then this code is dead. Stop here.
	if s.curBlock == nil && n.Op != OLABEL {
		return
	}

	s.stmtList(pstate, n.Ninit)
	switch n.Op {

	case OBLOCK:
		s.stmtList(pstate, n.List)

	// No-ops
	case OEMPTY, ODCLCONST, ODCLTYPE, OFALL:

	// Expression statements
	case OCALLFUNC:
		if pstate.isIntrinsicCall(n) {
			s.intrinsicCall(pstate, n)
			return
		}
		fallthrough

	case OCALLMETH, OCALLINTER:
		s.call(pstate, n, callNormal)
		if n.Op == OCALLFUNC && n.Left.Op == ONAME && n.Left.Class() == PFUNC {
			if fn := n.Left.Sym.Name; pstate.compiling_runtime && fn == "throw" ||
				n.Left.Sym.Pkg == pstate.Runtimepkg && (fn == "throwinit" || fn == "gopanic" || fn == "panicwrap" || fn == "block" || fn == "panicmakeslicelen" || fn == "panicmakeslicecap") {
				m := s.mem(pstate)
				b := s.endBlock(pstate)
				b.Kind = ssa.BlockExit
				b.SetControl(m)
				// TODO: never rewrite OPANIC to OCALLFUNC in the
				// first place. Need to wait until all backends
				// go through SSA.
			}
		}
	case ODEFER:
		s.call(pstate, n.Left, callDefer)
	case OPROC:
		s.call(pstate, n.Left, callGo)

	case OAS2DOTTYPE:
		res, resok := s.dottype(pstate, n.Rlist.First(), true)
		deref := false
		if !pstate.canSSAType(n.Rlist.First().Type) {
			if res.Op != ssa.OpLoad {
				s.Fatalf("dottype of non-load")
			}
			mem := s.mem(pstate)
			if mem.Op == ssa.OpVarKill {
				mem = mem.Args[0]
			}
			if res.Args[1] != mem {
				s.Fatalf("memory no longer live from 2-result dottype load")
			}
			deref = true
			res = res.Args[0]
		}
		s.assign(pstate, n.List.First(), res, deref, 0)
		s.assign(pstate, n.List.Second(), resok, false, 0)
		return

	case OAS2FUNC:
		// We come here only when it is an intrinsic call returning two values.
		if !pstate.isIntrinsicCall(n.Rlist.First()) {
			s.Fatalf("non-intrinsic AS2FUNC not expanded %v", n.Rlist.First())
		}
		v := s.intrinsicCall(pstate, n.Rlist.First())
		v1 := s.newValue1(ssa.OpSelect0, n.List.First().Type, v)
		v2 := s.newValue1(ssa.OpSelect1, n.List.Second().Type, v)
		s.assign(pstate, n.List.First(), v1, false, 0)
		s.assign(pstate, n.List.Second(), v2, false, 0)
		return

	case ODCL:
		if n.Left.Class() == PAUTOHEAP {
			pstate.Fatalf("DCL %v", n)
		}

	case OLABEL:
		sym := n.Left.Sym
		lab := s.label(sym)

		// Associate label with its control flow node, if any
		if ctl := n.labeledControl(pstate); ctl != nil {
			s.labeledNodes[ctl] = lab
		}

		// The label might already have a target block via a goto.
		if lab.target == nil {
			lab.target = s.f.NewBlock(ssa.BlockPlain)
		}

		// Go to that label.
		// (We pretend "label:" is preceded by "goto label", unless the predecessor is unreachable.)
		if s.curBlock != nil {
			b := s.endBlock(pstate)
			b.AddEdgeTo(lab.target)
		}
		s.startBlock(lab.target)

	case OGOTO:
		sym := n.Left.Sym

		lab := s.label(sym)
		if lab.target == nil {
			lab.target = s.f.NewBlock(ssa.BlockPlain)
		}

		b := s.endBlock(pstate)
		b.Pos = s.lastPos.WithIsStmt() // Do this even if b is an empty block.
		b.AddEdgeTo(lab.target)

	case OAS:
		if n.Left == n.Right && n.Left.Op == ONAME {
			// An x=x assignment. No point in doing anything
			// here. In addition, skipping this assignment
			// prevents generating:
			//   VARDEF x
			//   COPY x -> x
			// which is bad because x is incorrectly considered
			// dead before the vardef. See issue #14904.
			return
		}

		// Evaluate RHS.
		rhs := n.Right
		if rhs != nil {
			switch rhs.Op {
			case OSTRUCTLIT, OARRAYLIT, OSLICELIT:
				// All literals with nonzero fields have already been
				// rewritten during walk. Any that remain are just T{}
				// or equivalents. Use the zero value.
				if !pstate.isZero(rhs) {
					pstate.Fatalf("literal with nonzero value in SSA: %v", rhs)
				}
				rhs = nil
			case OAPPEND:
				// Check whether we're writing the result of an append back to the same slice.
				// If so, we handle it specially to avoid write barriers on the fast
				// (non-growth) path.
				if !pstate.samesafeexpr(n.Left, rhs.List.First()) || pstate.Debug['N'] != 0 {
					break
				}
				// If the slice can be SSA'd, it'll be on the stack,
				// so there will be no write barriers,
				// so there's no need to attempt to prevent them.
				if s.canSSA(pstate, n.Left) {
					if pstate.Debug_append > 0 { // replicating old diagnostic message
						pstate.Warnl(n.Pos, "append: len-only update (in local slice)")
					}
					break
				}
				if pstate.Debug_append > 0 {
					pstate.Warnl(n.Pos, "append: len-only update")
				}
				s.append(pstate, rhs, true)
				return
			}
		}

		if n.Left.isBlank() {
			// _ = rhs
			// Just evaluate rhs for side-effects.
			if rhs != nil {
				s.expr(pstate, rhs)
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
		deref := !pstate.canSSAType(t)
		if deref {
			if rhs == nil {
				r = nil // Signal assign to use OpZero.
			} else {
				r = s.addr(pstate, rhs, false)
			}
		} else {
			if rhs == nil {
				r = s.zeroVal(pstate, t)
			} else {
				r = s.expr(pstate, rhs)
			}
		}

		var skip skipMask
		if rhs != nil && (rhs.Op == OSLICE || rhs.Op == OSLICE3 || rhs.Op == OSLICESTR) && pstate.samesafeexpr(rhs.Left, n.Left) {
			// We're assigning a slicing operation back to its source.
			// Don't write back fields we aren't changing. See issue #14855.
			i, j, k := rhs.SliceBounds(pstate)
			if i != nil && (i.Op == OLITERAL && i.Val().Ctype(pstate) == CTINT && i.Int64(pstate) == 0) {
				// [0:...] is the same as [:...]
				i = nil
			}
			// TODO: detect defaults for len/cap also.
			// Currently doesn't really work because (*p)[:len(*p)] appears here as:
			//    tmp = len(*p)
			//    (*p)[:tmp]
			//if j != nil && (j.Op == OLEN && samesafeexpr(j.Left, n.Left)) {
			//      j = nil
			//}
			//if k != nil && (k.Op == OCAP && samesafeexpr(k.Left, n.Left)) {
			//      k = nil
			//}
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

		s.assign(pstate, n.Left, r, deref, skip)

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
			s.condBranch(pstate, n.Left, bThen, bElse, likely)
		} else {
			s.condBranch(pstate, n.Left, bThen, bEnd, likely)
		}

		s.startBlock(bThen)
		s.stmtList(pstate, n.Nbody)
		if b := s.endBlock(pstate); b != nil {
			b.AddEdgeTo(bEnd)
		}

		if n.Rlist.Len() != 0 {
			s.startBlock(bElse)
			s.stmtList(pstate, n.Rlist)
			if b := s.endBlock(pstate); b != nil {
				b.AddEdgeTo(bEnd)
			}
		}
		s.startBlock(bEnd)

	case ORETURN:
		s.stmtList(pstate, n.List)
		b := s.exit(pstate)
		b.Pos = s.lastPos.WithIsStmt()

	case ORETJMP:
		s.stmtList(pstate, n.List)
		b := s.exit(pstate)
		b.Kind = ssa.BlockRetJmp // override BlockRet
		b.Aux = n.Sym.Linksym(pstate.types)

	case OCONTINUE, OBREAK:
		var to *ssa.Block
		if n.Left == nil {
			// plain break/continue
			switch n.Op {
			case OCONTINUE:
				to = s.continueTo
			case OBREAK:
				to = s.breakTo
			}
		} else {
			// labeled break/continue; look up the target
			sym := n.Left.Sym
			lab := s.label(sym)
			switch n.Op {
			case OCONTINUE:
				to = lab.continueTarget
			case OBREAK:
				to = lab.breakTarget
			}
		}

		b := s.endBlock(pstate)
		b.Pos = s.lastPos.WithIsStmt() // Do this even if b is an empty block.
		b.AddEdgeTo(to)

	case OFOR, OFORUNTIL:
		// OFOR: for Ninit; Left; Right { Nbody }
		// cond (Left); body (Nbody); incr (Right)
		//
		// OFORUNTIL: for Ninit; Left; Right; List { Nbody }
		// => body: { Nbody }; incr: Right; if Left { lateincr: List; goto body }; end:
		bCond := s.f.NewBlock(ssa.BlockPlain)
		bBody := s.f.NewBlock(ssa.BlockPlain)
		bIncr := s.f.NewBlock(ssa.BlockPlain)
		bEnd := s.f.NewBlock(ssa.BlockPlain)

		// first, jump to condition test (OFOR) or body (OFORUNTIL)
		b := s.endBlock(pstate)
		if n.Op == OFOR {
			b.AddEdgeTo(bCond)
			// generate code to test condition
			s.startBlock(bCond)
			if n.Left != nil {
				s.condBranch(pstate, n.Left, bBody, bEnd, 1)
			} else {
				b := s.endBlock(pstate)
				b.Kind = ssa.BlockPlain
				b.AddEdgeTo(bBody)
			}

		} else {
			b.AddEdgeTo(bBody)
		}

		// set up for continue/break in body
		prevContinue := s.continueTo
		prevBreak := s.breakTo
		s.continueTo = bIncr
		s.breakTo = bEnd
		lab := s.labeledNodes[n]
		if lab != nil {
			// labeled for loop
			lab.continueTarget = bIncr
			lab.breakTarget = bEnd
		}

		// generate body
		s.startBlock(bBody)
		s.stmtList(pstate, n.Nbody)

		// tear down continue/break
		s.continueTo = prevContinue
		s.breakTo = prevBreak
		if lab != nil {
			lab.continueTarget = nil
			lab.breakTarget = nil
		}

		// done with body, goto incr
		if b := s.endBlock(pstate); b != nil {
			b.AddEdgeTo(bIncr)
		}

		// generate incr (and, for OFORUNTIL, condition)
		s.startBlock(bIncr)
		if n.Right != nil {
			s.stmt(pstate, n.Right)
		}
		if n.Op == OFOR {
			if b := s.endBlock(pstate); b != nil {
				b.AddEdgeTo(bCond)
				// It can happen that bIncr ends in a block containing only VARKILL,
				// and that muddles the debugging experience.
				if n.Op != OFORUNTIL && b.Pos == pstate.src.NoXPos {
					b.Pos = bCond.Pos
				}
			}
		} else {
			// bCond is unused in OFORUNTIL, so repurpose it.
			bLateIncr := bCond
			// test condition
			s.condBranch(pstate, n.Left, bLateIncr, bEnd, 1)
			// generate late increment
			s.startBlock(bLateIncr)
			s.stmtList(pstate, n.List)
			s.endBlock(pstate).AddEdgeTo(bBody)
		}

		s.startBlock(bEnd)

	case OSWITCH, OSELECT:
		// These have been mostly rewritten by the front end into their Nbody fields.
		// Our main task is to correctly hook up any break statements.
		bEnd := s.f.NewBlock(ssa.BlockPlain)

		prevBreak := s.breakTo
		s.breakTo = bEnd
		lab := s.labeledNodes[n]
		if lab != nil {
			// labeled
			lab.breakTarget = bEnd
		}

		// generate body code
		s.stmtList(pstate, n.Nbody)

		s.breakTo = prevBreak
		if lab != nil {
			lab.breakTarget = nil
		}

		// walk adds explicit OBREAK nodes to the end of all reachable code paths.
		// If we still have a current block here, then mark it unreachable.
		if s.curBlock != nil {
			m := s.mem(pstate)
			b := s.endBlock(pstate)
			b.Kind = ssa.BlockExit
			b.SetControl(m)
		}
		s.startBlock(bEnd)

	case OVARKILL:
		// Insert a varkill op to record that a variable is no longer live.
		// We only care about liveness info at call sites, so putting the
		// varkill in the store chain is enough to keep it correctly ordered
		// with respect to call ops.
		if !s.canSSA(pstate, n.Left) {
			s.vars[&pstate.memVar] = s.newValue1Apos(ssa.OpVarKill, pstate.types.TypeMem, n.Left, s.mem(pstate), false)
		}

	case OVARLIVE:
		// Insert a varlive op to record that a variable is still live.
		if !n.Left.Addrtaken() {
			s.Fatalf("VARLIVE variable %v must have Addrtaken set", n.Left)
		}
		switch n.Left.Class() {
		case PAUTO, PPARAM, PPARAMOUT:
		default:
			s.Fatalf("VARLIVE variable %v must be Auto or Arg", n.Left)
		}
		s.vars[&pstate.memVar] = s.newValue1A(ssa.OpVarLive, pstate.types.TypeMem, n.Left, s.mem(pstate))

	case OCHECKNIL:
		p := s.expr(pstate, n.Left)
		s.nilCheck(pstate, p)

	default:
		s.Fatalf("unhandled stmt %v", n.Op)
	}
}

// exit processes any code that needs to be generated just before returning.
// It returns a BlockRet block that ends the control flow. Its control value
// will be set to the final memory state.
func (s *state) exit(pstate *PackageState) *ssa.Block {
	if s.hasdefer {
		s.rtcall(pstate, pstate.Deferreturn, true, nil)
	}

	// Run exit code. Typically, this code copies heap-allocated PPARAMOUT
	// variables back to the stack.
	s.stmtList(pstate, s.curfn.Func.Exit)

	// Store SSAable PPARAMOUT variables back to stack locations.
	for _, n := range s.returns {
		addr := s.decladdrs[n]
		val := s.variable(n, n.Type)
		s.vars[&pstate.memVar] = s.newValue1A(ssa.OpVarDef, pstate.types.TypeMem, n, s.mem(pstate))
		s.store(pstate, n.Type, addr, val)
		// TODO: if val is ever spilled, we'd like to use the
		// PPARAMOUT slot for spilling it. That won't happen
		// currently.
	}

	// Do actual return.
	m := s.mem(pstate)
	b := s.endBlock(pstate)
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

func (s *state) ssaOp(pstate *PackageState, op Op, t *types.Type) ssa.Op {
	etype := s.concreteEtype(t)
	x, ok := pstate.opToSSA[opAndType{op, etype}]
	if !ok {
		s.Fatalf("unhandled binary op %v %s", op, etype)
	}
	return x
}

func (pstate *PackageState) floatForComplex(t *types.Type) *types.Type {
	if t.Size(pstate.types) == 8 {
		return pstate.types.Types[TFLOAT32]
	} else {
		return pstate.types.Types[TFLOAT64]
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

func (s *state) ssaShiftOp(pstate *PackageState, op Op, t *types.Type, u *types.Type) ssa.Op {
	etype1 := s.concreteEtype(t)
	etype2 := s.concreteEtype(u)
	x, ok := pstate.shiftOpToSSA[opAndTwoTypes{op, etype1, etype2}]
	if !ok {
		s.Fatalf("unhandled shift op %v etype=%s/%s", op, etype1, etype2)
	}
	return x
}

// expr converts the expression n to ssa, adds it to s and returns the ssa result.
func (s *state) expr(pstate *PackageState, n *Node) *ssa.Value {
	if !(n.Op == ONAME || n.Op == OLITERAL && n.Sym != nil) {
		// ONAMEs and named OLITERALs have the line number
		// of the decl, not the use. See issue 14742.
		s.pushLine(pstate, n.Pos)
		defer s.popLine()
	}

	s.stmtList(pstate, n.Ninit)
	switch n.Op {
	case OARRAYBYTESTRTMP:
		slice := s.expr(pstate, n.Left)
		ptr := s.newValue1(ssa.OpSlicePtr, s.f.Config.Types.BytePtr, slice)
		len := s.newValue1(ssa.OpSliceLen, pstate.types.Types[TINT], slice)
		return s.newValue2(ssa.OpStringMake, n.Type, ptr, len)
	case OSTRARRAYBYTETMP:
		str := s.expr(pstate, n.Left)
		ptr := s.newValue1(ssa.OpStringPtr, s.f.Config.Types.BytePtr, str)
		len := s.newValue1(ssa.OpStringLen, pstate.types.Types[TINT], str)
		return s.newValue3(ssa.OpSliceMake, n.Type, ptr, len, len)
	case OCFUNC:
		aux := n.Left.Sym.Linksym(pstate.types)
		return s.entryNewValue1A(pstate, ssa.OpAddr, n.Type, aux, s.sb)
	case ONAME:
		if n.Class() == PFUNC {
			// "value" of a function is the address of the function's closure
			sym := pstate.funcsym(n.Sym).Linksym(pstate.types)
			return s.entryNewValue1A(pstate, ssa.OpAddr, pstate.types.NewPtr(n.Type), sym, s.sb)
		}
		if s.canSSA(pstate, n) {
			return s.variable(n, n.Type)
		}
		addr := s.addr(pstate, n, false)
		return s.load(pstate, n.Type, addr)
	case OCLOSUREVAR:
		addr := s.addr(pstate, n, false)
		return s.load(pstate, n.Type, addr)
	case OLITERAL:
		switch u := n.Val().U.(type) {
		case *Mpint:
			i := u.Int64(pstate)
			switch n.Type.Size(pstate.types) {
			case 1:
				return s.constInt8(pstate, n.Type, int8(i))
			case 2:
				return s.constInt16(pstate, n.Type, int16(i))
			case 4:
				return s.constInt32(pstate, n.Type, int32(i))
			case 8:
				return s.constInt64(pstate, n.Type, i)
			default:
				s.Fatalf("bad integer size %d", n.Type.Size(pstate.types))
				return nil
			}
		case string:
			if u == "" {
				return s.constEmptyString(pstate, n.Type)
			}
			return s.entryNewValue0A(pstate, ssa.OpConstString, n.Type, u)
		case bool:
			return s.constBool(pstate, u)
		case *NilVal:
			t := n.Type
			switch {
			case t.IsSlice():
				return s.constSlice(pstate, t)
			case t.IsInterface():
				return s.constInterface(pstate, t)
			default:
				return s.constNil(pstate, t)
			}
		case *Mpflt:
			switch n.Type.Size(pstate.types) {
			case 4:
				return s.constFloat32(pstate, n.Type, u.Float32(pstate))
			case 8:
				return s.constFloat64(pstate, n.Type, u.Float64(pstate))
			default:
				s.Fatalf("bad float size %d", n.Type.Size(pstate.types))
				return nil
			}
		case *Mpcplx:
			r := &u.Real
			i := &u.Imag
			switch n.Type.Size(pstate.types) {
			case 8:
				pt := pstate.types.Types[TFLOAT32]
				return s.newValue2(ssa.OpComplexMake, n.Type,
					s.constFloat32(pstate, pt, r.Float32(pstate)),
					s.constFloat32(pstate, pt, i.Float32(pstate)))
			case 16:
				pt := pstate.types.Types[TFLOAT64]
				return s.newValue2(ssa.OpComplexMake, n.Type,
					s.constFloat64(pstate, pt, r.Float64(pstate)),
					s.constFloat64(pstate, pt, i.Float64(pstate)))
			default:
				s.Fatalf("bad float size %d", n.Type.Size(pstate.types))
				return nil
			}

		default:
			s.Fatalf("unhandled OLITERAL %v", n.Val().Ctype(pstate))
			return nil
		}
	case OCONVNOP:
		to := n.Type
		from := n.Left.Type

		// Assume everything will work out, so set up our return value.
		// Anything interesting that happens from here is a fatal.
		x := s.expr(pstate, n.Left)

		// Special case for not confusing GC and liveness.
		// We don't want pointers accidentally classified
		// as not-pointers or vice-versa because of copy
		// elision.
		if to.IsPtrShaped() != from.IsPtrShaped() {
			return s.newValue2(ssa.OpConvert, to, x, s.mem(pstate))
		}

		v := s.newValue1(ssa.OpCopy, to, x) // ensure that v has the right type

		// CONVNOP closure
		if to.Etype == TFUNC && from.IsPtrShaped() {
			return v
		}

		// named <--> unnamed type or typed <--> untyped const
		if from.Etype == to.Etype {
			return v
		}

		// unsafe.Pointer <--> *T
		if to.Etype == TUNSAFEPTR && from.IsPtrShaped() || from.Etype == TUNSAFEPTR && to.IsPtrShaped() {
			return v
		}

		// map <--> *hmap
		if to.Etype == TMAP && from.IsPtr() &&
			to.MapType(pstate.types).Hmap == from.Elem(pstate.types) {
			return v
		}

		pstate.dowidth(from)
		pstate.dowidth(to)
		if from.Width != to.Width {
			s.Fatalf("CONVNOP width mismatch %v (%d) -> %v (%d)\n", from, from.Width, to, to.Width)
			return nil
		}
		if etypesign(from.Etype) != etypesign(to.Etype) {
			s.Fatalf("CONVNOP sign mismatch %v (%s) -> %v (%s)\n", from, from.Etype, to, to.Etype)
			return nil
		}

		if pstate.instrumenting {
			// These appear to be fine, but they fail the
			// integer constraint below, so okay them here.
			// Sample non-integer conversion: map[string]string -> *uint8
			return v
		}

		if etypesign(from.Etype) == 0 {
			s.Fatalf("CONVNOP unrecognized non-integer %v -> %v\n", from, to)
			return nil
		}

		// integer, same width, same sign
		return v

	case OCONV:
		x := s.expr(pstate, n.Left)
		ft := n.Left.Type // from type
		tt := n.Type      // to type
		if ft.IsBoolean() && tt.IsKind(TUINT8) {
			// Bool -> uint8 is generated internally when indexing into runtime.staticbyte.
			return s.newValue1(ssa.OpCopy, n.Type, x)
		}
		if ft.IsInteger() && tt.IsInteger() {
			var op ssa.Op
			if tt.Size(pstate.types) == ft.Size(pstate.types) {
				op = ssa.OpCopy
			} else if tt.Size(pstate.types) < ft.Size(pstate.types) {
				// truncation
				switch 10*ft.Size(pstate.types) + tt.Size(pstate.types) {
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
				// sign extension
				switch 10*ft.Size(pstate.types) + tt.Size(pstate.types) {
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
				// zero extension
				switch 10*ft.Size(pstate.types) + tt.Size(pstate.types) {
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
			conv, ok := pstate.fpConvOpToSSA[twoTypes{s.concreteEtype(ft), s.concreteEtype(tt)}]
			if s.config.RegSize == 4 && pstate.thearch.LinkArch.Family != sys.MIPS && !s.softFloat {
				if conv1, ok1 := pstate.fpConvOpToSSA32[twoTypes{s.concreteEtype(ft), s.concreteEtype(tt)}]; ok1 {
					conv = conv1
				}
			}
			if pstate.thearch.LinkArch.Family == sys.ARM64 || pstate.thearch.LinkArch.Family == sys.Wasm || s.softFloat {
				if conv1, ok1 := pstate.uint64fpConvOpToSSA[twoTypes{s.concreteEtype(ft), s.concreteEtype(tt)}]; ok1 {
					conv = conv1
				}
			}

			if pstate.thearch.LinkArch.Family == sys.MIPS && !s.softFloat {
				if ft.Size(pstate.types) == 4 && ft.IsInteger() && !ft.IsSigned() {
					// tt is float32 or float64, and ft is also unsigned
					if tt.Size(pstate.types) == 4 {
						return s.uint32Tofloat32(pstate, n, x, ft, tt)
					}
					if tt.Size(pstate.types) == 8 {
						return s.uint32Tofloat64(pstate, n, x, ft, tt)
					}
				} else if tt.Size(pstate.types) == 4 && tt.IsInteger() && !tt.IsSigned() {
					// ft is float32 or float64, and tt is unsigned integer
					if ft.Size(pstate.types) == 4 {
						return s.float32ToUint32(pstate, n, x, ft, tt)
					}
					if ft.Size(pstate.types) == 8 {
						return s.float64ToUint32(pstate, n, x, ft, tt)
					}
				}
			}

			if !ok {
				s.Fatalf("weird float conversion %v -> %v", ft, tt)
			}
			op1, op2, it := conv.op1, conv.op2, conv.intermediateType

			if op1 != ssa.OpInvalid && op2 != ssa.OpInvalid {
				// normal case, not tripping over unsigned 64
				if op1 == ssa.OpCopy {
					if op2 == ssa.OpCopy {
						return x
					}
					return s.newValueOrSfCall1(pstate, op2, n.Type, x)
				}
				if op2 == ssa.OpCopy {
					return s.newValueOrSfCall1(pstate, op1, n.Type, x)
				}
				return s.newValueOrSfCall1(pstate, op2, n.Type, s.newValueOrSfCall1(pstate, op1, pstate.types.Types[it], x))
			}
			// Tricky 64-bit unsigned cases.
			if ft.IsInteger() {
				// tt is float32 or float64, and ft is also unsigned
				if tt.Size(pstate.types) == 4 {
					return s.uint64Tofloat32(pstate, n, x, ft, tt)
				}
				if tt.Size(pstate.types) == 8 {
					return s.uint64Tofloat64(pstate, n, x, ft, tt)
				}
				s.Fatalf("weird unsigned integer to float conversion %v -> %v", ft, tt)
			}
			// ft is float32 or float64, and tt is unsigned integer
			if ft.Size(pstate.types) == 4 {
				return s.float32ToUint64(pstate, n, x, ft, tt)
			}
			if ft.Size(pstate.types) == 8 {
				return s.float64ToUint64(pstate, n, x, ft, tt)
			}
			s.Fatalf("weird float to unsigned integer conversion %v -> %v", ft, tt)
			return nil
		}

		if ft.IsComplex() && tt.IsComplex() {
			var op ssa.Op
			if ft.Size(pstate.types) == tt.Size(pstate.types) {
				switch ft.Size(pstate.types) {
				case 8:
					op = ssa.OpRound32F
				case 16:
					op = ssa.OpRound64F
				default:
					s.Fatalf("weird complex conversion %v -> %v", ft, tt)
				}
			} else if ft.Size(pstate.types) == 8 && tt.Size(pstate.types) == 16 {
				op = ssa.OpCvt32Fto64F
			} else if ft.Size(pstate.types) == 16 && tt.Size(pstate.types) == 8 {
				op = ssa.OpCvt64Fto32F
			} else {
				s.Fatalf("weird complex conversion %v -> %v", ft, tt)
			}
			ftp := pstate.floatForComplex(ft)
			ttp := pstate.floatForComplex(tt)
			return s.newValue2(ssa.OpComplexMake, tt,
				s.newValueOrSfCall1(pstate, op, ttp, s.newValue1(ssa.OpComplexReal, ftp, x)),
				s.newValueOrSfCall1(pstate, op, ttp, s.newValue1(ssa.OpComplexImag, ftp, x)))
		}

		s.Fatalf("unhandled OCONV %s -> %s", n.Left.Type.Etype, n.Type.Etype)
		return nil

	case ODOTTYPE:
		res, _ := s.dottype(pstate, n, false)
		return res

	// binary ops
	case OLT, OEQ, ONE, OLE, OGE, OGT:
		a := s.expr(pstate, n.Left)
		b := s.expr(pstate, n.Right)
		if n.Left.Type.IsComplex() {
			pt := pstate.floatForComplex(n.Left.Type)
			op := s.ssaOp(pstate, OEQ, pt)
			r := s.newValueOrSfCall2(pstate, op, pstate.types.Types[TBOOL], s.newValue1(ssa.OpComplexReal, pt, a), s.newValue1(ssa.OpComplexReal, pt, b))
			i := s.newValueOrSfCall2(pstate, op, pstate.types.Types[TBOOL], s.newValue1(ssa.OpComplexImag, pt, a), s.newValue1(ssa.OpComplexImag, pt, b))
			c := s.newValue2(ssa.OpAndB, pstate.types.Types[TBOOL], r, i)
			switch n.Op {
			case OEQ:
				return c
			case ONE:
				return s.newValue1(ssa.OpNot, pstate.types.Types[TBOOL], c)
			default:
				s.Fatalf("ordered complex compare %v", n.Op)
			}
		}
		if n.Left.Type.IsFloat() {
			return s.newValueOrSfCall2(pstate, s.ssaOp(pstate, n.Op, n.Left.Type), pstate.types.Types[TBOOL], a, b)
		}
		return s.newValue2(s.ssaOp(pstate, n.Op, n.Left.Type), pstate.types.Types[TBOOL], a, b)
	case OMUL:
		a := s.expr(pstate, n.Left)
		b := s.expr(pstate, n.Right)
		if n.Type.IsComplex() {
			mulop := ssa.OpMul64F
			addop := ssa.OpAdd64F
			subop := ssa.OpSub64F
			pt := pstate.floatForComplex(n.Type) // Could be Float32 or Float64
			wt := pstate.types.Types[TFLOAT64]   // Compute in Float64 to minimize cancelation error

			areal := s.newValue1(ssa.OpComplexReal, pt, a)
			breal := s.newValue1(ssa.OpComplexReal, pt, b)
			aimag := s.newValue1(ssa.OpComplexImag, pt, a)
			bimag := s.newValue1(ssa.OpComplexImag, pt, b)

			if pt != wt { // Widen for calculation
				areal = s.newValueOrSfCall1(pstate, ssa.OpCvt32Fto64F, wt, areal)
				breal = s.newValueOrSfCall1(pstate, ssa.OpCvt32Fto64F, wt, breal)
				aimag = s.newValueOrSfCall1(pstate, ssa.OpCvt32Fto64F, wt, aimag)
				bimag = s.newValueOrSfCall1(pstate, ssa.OpCvt32Fto64F, wt, bimag)
			}

			xreal := s.newValueOrSfCall2(pstate, subop, wt, s.newValueOrSfCall2(pstate, mulop, wt, areal, breal), s.newValueOrSfCall2(pstate, mulop, wt, aimag, bimag))
			ximag := s.newValueOrSfCall2(pstate, addop, wt, s.newValueOrSfCall2(pstate, mulop, wt, areal, bimag), s.newValueOrSfCall2(pstate, mulop, wt, aimag, breal))

			if pt != wt { // Narrow to store back
				xreal = s.newValueOrSfCall1(pstate, ssa.OpCvt64Fto32F, pt, xreal)
				ximag = s.newValueOrSfCall1(pstate, ssa.OpCvt64Fto32F, pt, ximag)
			}

			return s.newValue2(ssa.OpComplexMake, n.Type, xreal, ximag)
		}

		if n.Type.IsFloat() {
			return s.newValueOrSfCall2(pstate, s.ssaOp(pstate, n.Op, n.Type), a.Type, a, b)
		}

		return s.newValue2(s.ssaOp(pstate, n.Op, n.Type), a.Type, a, b)

	case ODIV:
		a := s.expr(pstate, n.Left)
		b := s.expr(pstate, n.Right)
		if n.Type.IsComplex() {
			// TODO this is not executed because the front-end substitutes a runtime call.
			// That probably ought to change; with modest optimization the widen/narrow
			// conversions could all be elided in larger expression trees.
			mulop := ssa.OpMul64F
			addop := ssa.OpAdd64F
			subop := ssa.OpSub64F
			divop := ssa.OpDiv64F
			pt := pstate.floatForComplex(n.Type) // Could be Float32 or Float64
			wt := pstate.types.Types[TFLOAT64]   // Compute in Float64 to minimize cancelation error

			areal := s.newValue1(ssa.OpComplexReal, pt, a)
			breal := s.newValue1(ssa.OpComplexReal, pt, b)
			aimag := s.newValue1(ssa.OpComplexImag, pt, a)
			bimag := s.newValue1(ssa.OpComplexImag, pt, b)

			if pt != wt { // Widen for calculation
				areal = s.newValueOrSfCall1(pstate, ssa.OpCvt32Fto64F, wt, areal)
				breal = s.newValueOrSfCall1(pstate, ssa.OpCvt32Fto64F, wt, breal)
				aimag = s.newValueOrSfCall1(pstate, ssa.OpCvt32Fto64F, wt, aimag)
				bimag = s.newValueOrSfCall1(pstate, ssa.OpCvt32Fto64F, wt, bimag)
			}

			denom := s.newValueOrSfCall2(pstate, addop, wt, s.newValueOrSfCall2(pstate, mulop, wt, breal, breal), s.newValueOrSfCall2(pstate, mulop, wt, bimag, bimag))
			xreal := s.newValueOrSfCall2(pstate, addop, wt, s.newValueOrSfCall2(pstate, mulop, wt, areal, breal), s.newValueOrSfCall2(pstate, mulop, wt, aimag, bimag))
			ximag := s.newValueOrSfCall2(pstate, subop, wt, s.newValueOrSfCall2(pstate, mulop, wt, aimag, breal), s.newValueOrSfCall2(pstate, mulop, wt, areal, bimag))

			// TODO not sure if this is best done in wide precision or narrow
			// Double-rounding might be an issue.
			// Note that the pre-SSA implementation does the entire calculation
			// in wide format, so wide is compatible.
			xreal = s.newValueOrSfCall2(pstate, divop, wt, xreal, denom)
			ximag = s.newValueOrSfCall2(pstate, divop, wt, ximag, denom)

			if pt != wt { // Narrow to store back
				xreal = s.newValueOrSfCall1(pstate, ssa.OpCvt64Fto32F, pt, xreal)
				ximag = s.newValueOrSfCall1(pstate, ssa.OpCvt64Fto32F, pt, ximag)
			}
			return s.newValue2(ssa.OpComplexMake, n.Type, xreal, ximag)
		}
		if n.Type.IsFloat() {
			return s.newValueOrSfCall2(pstate, s.ssaOp(pstate, n.Op, n.Type), a.Type, a, b)
		}
		return s.intDivide(pstate, n, a, b)
	case OMOD:
		a := s.expr(pstate, n.Left)
		b := s.expr(pstate, n.Right)
		return s.intDivide(pstate, n, a, b)
	case OADD, OSUB:
		a := s.expr(pstate, n.Left)
		b := s.expr(pstate, n.Right)
		if n.Type.IsComplex() {
			pt := pstate.floatForComplex(n.Type)
			op := s.ssaOp(pstate, n.Op, pt)
			return s.newValue2(ssa.OpComplexMake, n.Type,
				s.newValueOrSfCall2(pstate, op, pt, s.newValue1(ssa.OpComplexReal, pt, a), s.newValue1(ssa.OpComplexReal, pt, b)),
				s.newValueOrSfCall2(pstate, op, pt, s.newValue1(ssa.OpComplexImag, pt, a), s.newValue1(ssa.OpComplexImag, pt, b)))
		}
		if n.Type.IsFloat() {
			return s.newValueOrSfCall2(pstate, s.ssaOp(pstate, n.Op, n.Type), a.Type, a, b)
		}
		return s.newValue2(s.ssaOp(pstate, n.Op, n.Type), a.Type, a, b)
	case OAND, OOR, OXOR:
		a := s.expr(pstate, n.Left)
		b := s.expr(pstate, n.Right)
		return s.newValue2(s.ssaOp(pstate, n.Op, n.Type), a.Type, a, b)
	case OLSH, ORSH:
		a := s.expr(pstate, n.Left)
		b := s.expr(pstate, n.Right)
		return s.newValue2(s.ssaShiftOp(pstate, n.Op, n.Type, n.Right.Type), a.Type, a, b)
	case OANDAND, OOROR:
		// To implement OANDAND (and OOROR), we introduce a
		// new temporary variable to hold the result. The
		// variable is associated with the OANDAND node in the
		// s.vars table (normally variables are only
		// associated with ONAME nodes). We convert
		//     A && B
		// to
		//     var = A
		//     if var {
		//         var = B
		//     }
		// Using var in the subsequent block introduces the
		// necessary phi variable.
		el := s.expr(pstate, n.Left)
		s.vars[n] = el

		b := s.endBlock(pstate)
		b.Kind = ssa.BlockIf
		b.SetControl(el)
		// In theory, we should set b.Likely here based on context.
		// However, gc only gives us likeliness hints
		// in a single place, for plain OIF statements,
		// and passing around context is finnicky, so don't bother for now.

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
		er := s.expr(pstate, n.Right)
		s.vars[n] = er

		b = s.endBlock(pstate)
		b.AddEdgeTo(bResult)

		s.startBlock(bResult)
		return s.variable(n, pstate.types.Types[TBOOL])
	case OCOMPLEX:
		r := s.expr(pstate, n.Left)
		i := s.expr(pstate, n.Right)
		return s.newValue2(ssa.OpComplexMake, n.Type, r, i)

	// unary ops
	case OMINUS:
		a := s.expr(pstate, n.Left)
		if n.Type.IsComplex() {
			tp := pstate.floatForComplex(n.Type)
			negop := s.ssaOp(pstate, n.Op, tp)
			return s.newValue2(ssa.OpComplexMake, n.Type,
				s.newValue1(negop, tp, s.newValue1(ssa.OpComplexReal, tp, a)),
				s.newValue1(negop, tp, s.newValue1(ssa.OpComplexImag, tp, a)))
		}
		return s.newValue1(s.ssaOp(pstate, n.Op, n.Type), a.Type, a)
	case ONOT, OCOM:
		a := s.expr(pstate, n.Left)
		return s.newValue1(s.ssaOp(pstate, n.Op, n.Type), a.Type, a)
	case OIMAG, OREAL:
		a := s.expr(pstate, n.Left)
		return s.newValue1(s.ssaOp(pstate, n.Op, n.Left.Type), n.Type, a)
	case OPLUS:
		return s.expr(pstate, n.Left)

	case OADDR:
		return s.addr(pstate, n.Left, n.Bounded())

	case OINDREGSP:
		addr := s.constOffPtrSP(pstate, pstate.types.NewPtr(n.Type), n.Xoffset)
		return s.load(pstate, n.Type, addr)

	case OIND:
		p := s.exprPtr(pstate, n.Left, false, n.Pos)
		return s.load(pstate, n.Type, p)

	case ODOT:
		if n.Left.Op == OSTRUCTLIT {
			// All literals with nonzero fields have already been
			// rewritten during walk. Any that remain are just T{}
			// or equivalents. Use the zero value.
			if !pstate.isZero(n.Left) {
				pstate.Fatalf("literal with nonzero value in SSA: %v", n.Left)
			}
			return s.zeroVal(pstate, n.Type)
		}
		// If n is addressable and can't be represented in
		// SSA, then load just the selected field. This
		// prevents false memory dependencies in race/msan
		// instrumentation.
		if islvalue(n) && !s.canSSA(pstate, n) {
			p := s.addr(pstate, n, false)
			return s.load(pstate, n.Type, p)
		}
		v := s.expr(pstate, n.Left)
		return s.newValue1I(ssa.OpStructSelect, n.Type, int64(pstate.fieldIdx(n)), v)

	case ODOTPTR:
		p := s.exprPtr(pstate, n.Left, false, n.Pos)
		p = s.newValue1I(ssa.OpOffPtr, pstate.types.NewPtr(n.Type), n.Xoffset, p)
		return s.load(pstate, n.Type, p)

	case OINDEX:
		switch {
		case n.Left.Type.IsString():
			if n.Bounded() && pstate.Isconst(n.Left, CTSTR) && pstate.Isconst(n.Right, CTINT) {
				// Replace "abc"[1] with 'b'.
				// Delayed until now because "abc"[1] is not an ideal constant.
				// See test/fixedbugs/issue11370.go.
				return s.newValue0I(ssa.OpConst8, pstate.types.Types[TUINT8], int64(int8(n.Left.Val().U.(string)[n.Right.Int64(pstate)])))
			}
			a := s.expr(pstate, n.Left)
			i := s.expr(pstate, n.Right)
			i = s.extendIndex(pstate, i, pstate.panicindex)
			if !n.Bounded() {
				len := s.newValue1(ssa.OpStringLen, pstate.types.Types[TINT], a)
				s.boundsCheck(pstate, i, len)
			}
			ptrtyp := s.f.Config.Types.BytePtr
			ptr := s.newValue1(ssa.OpStringPtr, ptrtyp, a)
			if pstate.Isconst(n.Right, CTINT) {
				ptr = s.newValue1I(ssa.OpOffPtr, ptrtyp, n.Right.Int64(pstate), ptr)
			} else {
				ptr = s.newValue2(ssa.OpAddPtr, ptrtyp, ptr, i)
			}
			return s.load(pstate, pstate.types.Types[TUINT8], ptr)
		case n.Left.Type.IsSlice():
			p := s.addr(pstate, n, false)
			return s.load(pstate, n.Left.Type.Elem(pstate.types), p)
		case n.Left.Type.IsArray():
			if bound := n.Left.Type.NumElem(pstate.types); bound <= 1 {
				// SSA can handle arrays of length at most 1.
				a := s.expr(pstate, n.Left)
				i := s.expr(pstate, n.Right)
				if bound == 0 {
					// Bounds check will never succeed.  Might as well
					// use constants for the bounds check.
					z := s.constInt(pstate, pstate.types.Types[TINT], 0)
					s.boundsCheck(pstate, z, z)
					// The return value won't be live, return junk.
					return s.newValue0(ssa.OpUnknown, n.Type)
				}
				i = s.extendIndex(pstate, i, pstate.panicindex)
				if !n.Bounded() {
					s.boundsCheck(pstate, i, s.constInt(pstate, pstate.types.Types[TINT], bound))
				}
				return s.newValue1I(ssa.OpArraySelect, n.Type, 0, a)
			}
			p := s.addr(pstate, n, false)
			return s.load(pstate, n.Left.Type.Elem(pstate.types), p)
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
			return s.newValue1(op, pstate.types.Types[TINT], s.expr(pstate, n.Left))
		case n.Left.Type.IsString(): // string; not reachable for OCAP
			return s.newValue1(ssa.OpStringLen, pstate.types.Types[TINT], s.expr(pstate, n.Left))
		case n.Left.Type.IsMap(), n.Left.Type.IsChan():
			return s.referenceTypeBuiltin(pstate, n, s.expr(pstate, n.Left))
		default: // array
			return s.constInt(pstate, pstate.types.Types[TINT], n.Left.Type.NumElem(pstate.types))
		}

	case OSPTR:
		a := s.expr(pstate, n.Left)
		if n.Left.Type.IsSlice() {
			return s.newValue1(ssa.OpSlicePtr, n.Type, a)
		} else {
			return s.newValue1(ssa.OpStringPtr, n.Type, a)
		}

	case OITAB:
		a := s.expr(pstate, n.Left)
		return s.newValue1(ssa.OpITab, n.Type, a)

	case OIDATA:
		a := s.expr(pstate, n.Left)
		return s.newValue1(ssa.OpIData, n.Type, a)

	case OEFACE:
		tab := s.expr(pstate, n.Left)
		data := s.expr(pstate, n.Right)
		return s.newValue2(ssa.OpIMake, n.Type, tab, data)

	case OSLICE, OSLICEARR, OSLICE3, OSLICE3ARR:
		v := s.expr(pstate, n.Left)
		var i, j, k *ssa.Value
		low, high, max := n.SliceBounds(pstate)
		if low != nil {
			i = s.extendIndex(pstate, s.expr(pstate, low), pstate.panicslice)
		}
		if high != nil {
			j = s.extendIndex(pstate, s.expr(pstate, high), pstate.panicslice)
		}
		if max != nil {
			k = s.extendIndex(pstate, s.expr(pstate, max), pstate.panicslice)
		}
		p, l, c := s.slice(pstate, n.Left.Type, v, i, j, k)
		return s.newValue3(ssa.OpSliceMake, n.Type, p, l, c)

	case OSLICESTR:
		v := s.expr(pstate, n.Left)
		var i, j *ssa.Value
		low, high, _ := n.SliceBounds(pstate)
		if low != nil {
			i = s.extendIndex(pstate, s.expr(pstate, low), pstate.panicslice)
		}
		if high != nil {
			j = s.extendIndex(pstate, s.expr(pstate, high), pstate.panicslice)
		}
		p, l, _ := s.slice(pstate, n.Left.Type, v, i, j, nil)
		return s.newValue2(ssa.OpStringMake, n.Type, p, l)

	case OCALLFUNC:
		if pstate.isIntrinsicCall(n) {
			return s.intrinsicCall(pstate, n)
		}
		fallthrough

	case OCALLINTER, OCALLMETH:
		a := s.call(pstate, n, callNormal)
		return s.load(pstate, n.Type, a)

	case OGETG:
		return s.newValue1(ssa.OpGetG, n.Type, s.mem(pstate))

	case OAPPEND:
		return s.append(pstate, n, false)

	case OSTRUCTLIT, OARRAYLIT:
		// All literals with nonzero fields have already been
		// rewritten during walk. Any that remain are just T{}
		// or equivalents. Use the zero value.
		if !pstate.isZero(n) {
			pstate.Fatalf("literal with nonzero value in SSA: %v", n)
		}
		return s.zeroVal(pstate, n.Type)

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
func (s *state) append(pstate *PackageState, n *Node, inplace bool) *ssa.Value {
	// If inplace is false, process as expression "append(s, e1, e2, e3)":
	//
	// ptr, len, cap := s
	// newlen := len + 3
	// if newlen > cap {
	//     ptr, len, cap = growslice(s, newlen)
	//     newlen = len + 3 // recalculate to avoid a spill
	// }
	// // with write barriers, if needed:
	// *(ptr+len) = e1
	// *(ptr+len+1) = e2
	// *(ptr+len+2) = e3
	// return makeslice(ptr, newlen, cap)
	//
	//
	// If inplace is true, process as statement "s = append(s, e1, e2, e3)":
	//
	// a := &s
	// ptr, len, cap := s
	// newlen := len + 3
	// if newlen > cap {
	//    newptr, len, newcap = growslice(ptr, len, cap, newlen)
	//    vardef(a)       // if necessary, advise liveness we are writing a new a
	//    *a.cap = newcap // write before ptr to avoid a spill
	//    *a.ptr = newptr // with write barrier
	// }
	// newlen = len + 3 // recalculate to avoid a spill
	// *a.len = newlen
	// // with write barriers, if needed:
	// *(ptr+len) = e1
	// *(ptr+len+1) = e2
	// *(ptr+len+2) = e3

	et := n.Type.Elem(pstate.types)
	pt := pstate.types.NewPtr(et)

	// Evaluate slice
	sn := n.List.First() // the slice node is the first in the list

	var slice, addr *ssa.Value
	if inplace {
		addr = s.addr(pstate, sn, false)
		slice = s.load(pstate, n.Type, addr)
	} else {
		slice = s.expr(pstate, sn)
	}

	// Allocate new blocks
	grow := s.f.NewBlock(ssa.BlockPlain)
	assign := s.f.NewBlock(ssa.BlockPlain)

	// Decide if we need to grow
	nargs := int64(n.List.Len() - 1)
	p := s.newValue1(ssa.OpSlicePtr, pt, slice)
	l := s.newValue1(ssa.OpSliceLen, pstate.types.Types[TINT], slice)
	c := s.newValue1(ssa.OpSliceCap, pstate.types.Types[TINT], slice)
	nl := s.newValue2(s.ssaOp(pstate, OADD, pstate.types.Types[TINT]), pstate.types.Types[TINT], l, s.constInt(pstate, pstate.types.Types[TINT], nargs))

	cmp := s.newValue2(s.ssaOp(pstate, OGT, pstate.types.Types[TINT]), pstate.types.Types[TBOOL], nl, c)
	s.vars[&pstate.ptrVar] = p

	if !inplace {
		s.vars[&pstate.newlenVar] = nl
		s.vars[&pstate.capVar] = c
	} else {
		s.vars[&pstate.lenVar] = l
	}

	b := s.endBlock(pstate)
	b.Kind = ssa.BlockIf
	b.Likely = ssa.BranchUnlikely
	b.SetControl(cmp)
	b.AddEdgeTo(grow)
	b.AddEdgeTo(assign)

	// Call growslice
	s.startBlock(grow)
	taddr := s.expr(pstate, n.Left)
	r := s.rtcall(pstate, pstate.growslice, true, []*types.Type{pt, pstate.types.Types[TINT], pstate.types.Types[TINT]}, taddr, p, l, c, nl)

	if inplace {
		if sn.Op == ONAME && sn.Class() != PEXTERN {
			// Tell liveness we're about to build a new slice
			s.vars[&pstate.memVar] = s.newValue1A(ssa.OpVarDef, pstate.types.TypeMem, sn, s.mem(pstate))
		}
		capaddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.IntPtr, int64(pstate.array_cap), addr)
		s.store(pstate, pstate.types.Types[TINT], capaddr, r[2])
		s.store(pstate, pt, addr, r[0])
		// load the value we just stored to avoid having to spill it
		s.vars[&pstate.ptrVar] = s.load(pstate, pt, addr)
		s.vars[&pstate.lenVar] = r[1] // avoid a spill in the fast path
	} else {
		s.vars[&pstate.ptrVar] = r[0]
		s.vars[&pstate.newlenVar] = s.newValue2(s.ssaOp(pstate, OADD, pstate.types.Types[TINT]), pstate.types.Types[TINT], r[1], s.constInt(pstate, pstate.types.Types[TINT], nargs))
		s.vars[&pstate.capVar] = r[2]
	}

	b = s.endBlock(pstate)
	b.AddEdgeTo(assign)

	// assign new elements to slots
	s.startBlock(assign)

	if inplace {
		l = s.variable(&pstate.lenVar, pstate.types.Types[TINT]) // generates phi for len
		nl = s.newValue2(s.ssaOp(pstate, OADD, pstate.types.Types[TINT]), pstate.types.Types[TINT], l, s.constInt(pstate, pstate.types.Types[TINT], nargs))
		lenaddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.IntPtr, int64(pstate.array_nel), addr)
		s.store(pstate, pstate.types.Types[TINT], lenaddr, nl)
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
		if pstate.canSSAType(n.Type) {
			args = append(args, argRec{v: s.expr(pstate, n), store: true})
		} else {
			v := s.addr(pstate, n, false)
			args = append(args, argRec{v: v})
		}
	}

	p = s.variable(&pstate.ptrVar, pt) // generates phi for ptr
	if !inplace {
		nl = s.variable(&pstate.newlenVar, pstate.types.Types[TINT]) // generates phi for nl
		c = s.variable(&pstate.capVar, pstate.types.Types[TINT])     // generates phi for cap
	}
	p2 := s.newValue2(ssa.OpPtrIndex, pt, p, l)
	for i, arg := range args {
		addr := s.newValue2(ssa.OpPtrIndex, pt, p2, s.constInt(pstate, pstate.types.Types[TINT], int64(i)))
		if arg.store {
			s.storeType(pstate, et, addr, arg.v, 0, true)
		} else {
			s.move(pstate, et, addr, arg.v)
		}
	}

	delete(s.vars, &pstate.ptrVar)
	if inplace {
		delete(s.vars, &pstate.lenVar)
		return nil
	}
	delete(s.vars, &pstate.newlenVar)
	delete(s.vars, &pstate.capVar)
	// make result
	return s.newValue3(ssa.OpSliceMake, n.Type, p, nl, c)
}

// condBranch evaluates the boolean expression cond and branches to yes
// if cond is true and no if cond is false.
// This function is intended to handle && and || better than just calling
// s.expr(cond) and branching on the result.
func (s *state) condBranch(pstate *PackageState, cond *Node, yes, no *ssa.Block, likely int8) {
	switch cond.Op {
	case OANDAND:
		mid := s.f.NewBlock(ssa.BlockPlain)
		s.stmtList(pstate, cond.Ninit)
		s.condBranch(pstate, cond.Left, mid, no, max8(likely, 0))
		s.startBlock(mid)
		s.condBranch(pstate, cond.Right, yes, no, likely)
		return
	// Note: if likely==1, then both recursive calls pass 1.
	// If likely==-1, then we don't have enough information to decide
	// whether the first branch is likely or not. So we pass 0 for
	// the likeliness of the first branch.
	// TODO: have the frontend give us branch prediction hints for
	// OANDAND and OOROR nodes (if it ever has such info).
	case OOROR:
		mid := s.f.NewBlock(ssa.BlockPlain)
		s.stmtList(pstate, cond.Ninit)
		s.condBranch(pstate, cond.Left, yes, mid, min8(likely, 0))
		s.startBlock(mid)
		s.condBranch(pstate, cond.Right, yes, no, likely)
		return
	// Note: if likely==-1, then both recursive calls pass -1.
	// If likely==1, then we don't have enough info to decide
	// the likelihood of the first branch.
	case ONOT:
		s.stmtList(pstate, cond.Ninit)
		s.condBranch(pstate, cond.Left, no, yes, -likely)
		return
	}
	c := s.expr(pstate, cond)
	b := s.endBlock(pstate)
	b.Kind = ssa.BlockIf
	b.SetControl(c)
	b.Likely = ssa.BranchPrediction(likely) // gc and ssa both use -1/0/+1 for likeliness
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
func (s *state) assign(pstate *PackageState, left *Node, right *ssa.Value, deref bool, skip skipMask) {
	if left.Op == ONAME && left.isBlank() {
		return
	}
	t := left.Type
	pstate.dowidth(t)
	if s.canSSA(pstate, left) {
		if deref {
			s.Fatalf("can SSA LHS %v but not RHS %s", left, right)
		}
		if left.Op == ODOT {
			// We're assigning to a field of an ssa-able value.
			// We need to build a new structure with the new value for the
			// field we're assigning and the old values for the other fields.
			// For instance:
			//   type T struct {a, b, c int}
			//   var T x
			//   x.b = 5
			// For the x.b = 5 assignment we want to generate x = T{x.a, 5, x.c}

			// Grab information about the structure type.
			t := left.Left.Type
			nf := t.NumFields(pstate.types)
			idx := pstate.fieldIdx(left)

			// Grab old value of structure.
			old := s.expr(pstate, left.Left)

			// Make new structure.
			new := s.newValue0(ssa.StructMakeOp(t.NumFields(pstate.types)), t)

			// Add fields as args.
			for i := 0; i < nf; i++ {
				if i == idx {
					new.AddArg(right)
				} else {
					new.AddArg(s.newValue1I(ssa.OpStructSelect, t.FieldType(pstate.types, i), int64(i), old))
				}
			}

			// Recursively assign the new value we've made to the base of the dot op.
			s.assign(pstate, left.Left, new, false, 0)
			// TODO: do we need to update named values here?
			return
		}
		if left.Op == OINDEX && left.Left.Type.IsArray() {
			// We're assigning to an element of an ssa-able array.
			// a[i] = v
			t := left.Left.Type
			n := t.NumElem(pstate.types)

			i := s.expr(pstate, left.Right) // index
			if n == 0 {
				// The bounds check must fail.  Might as well
				// ignore the actual index and just use zeros.
				z := s.constInt(pstate, pstate.types.Types[TINT], 0)
				s.boundsCheck(pstate, z, z)
				return
			}
			if n != 1 {
				s.Fatalf("assigning to non-1-length array")
			}
			// Rewrite to a = [1]{v}
			i = s.extendIndex(pstate, i, pstate.panicindex)
			s.boundsCheck(pstate, i, s.constInt(pstate, pstate.types.Types[TINT], 1))
			v := s.newValue1(ssa.OpArrayMake1, t, right)
			s.assign(pstate, left.Left, v, false, 0)
			return
		}
		// Update variable assignment.
		s.vars[left] = right
		s.addNamedValue(left, right)
		return
	}
	// Left is not ssa-able. Compute its address.
	addr := s.addr(pstate, left, false)
	if left.Op == ONAME && left.Class() != PEXTERN && skip == 0 {
		s.vars[&pstate.memVar] = s.newValue1Apos(ssa.OpVarDef, pstate.types.TypeMem, left, s.mem(pstate), !left.IsAutoTmp())
	}
	if pstate.isReflectHeaderDataField(left) {
		// Package unsafe's documentation says storing pointers into
		// reflect.SliceHeader and reflect.StringHeader's Data fields
		// is valid, even though they have type uintptr (#19168).
		// Mark it pointer type to signal the writebarrier pass to
		// insert a write barrier.
		t = pstate.types.Types[TUNSAFEPTR]
	}
	if deref {
		// Treat as a mem->mem move.
		if right == nil {
			s.zero(pstate, t, addr)
		} else {
			s.move(pstate, t, addr, right)
		}
		return
	}
	// Treat as a store.
	s.storeType(pstate, t, addr, right, skip, !left.IsAutoTmp())
}

// zeroVal returns the zero value for type t.
func (s *state) zeroVal(pstate *PackageState, t *types.Type) *ssa.Value {
	switch {
	case t.IsInteger():
		switch t.Size(pstate.types) {
		case 1:
			return s.constInt8(pstate, t, 0)
		case 2:
			return s.constInt16(pstate, t, 0)
		case 4:
			return s.constInt32(pstate, t, 0)
		case 8:
			return s.constInt64(pstate, t, 0)
		default:
			s.Fatalf("bad sized integer type %v", t)
		}
	case t.IsFloat():
		switch t.Size(pstate.types) {
		case 4:
			return s.constFloat32(pstate, t, 0)
		case 8:
			return s.constFloat64(pstate, t, 0)
		default:
			s.Fatalf("bad sized float type %v", t)
		}
	case t.IsComplex():
		switch t.Size(pstate.types) {
		case 8:
			z := s.constFloat32(pstate, pstate.types.Types[TFLOAT32], 0)
			return s.entryNewValue2(pstate, ssa.OpComplexMake, t, z, z)
		case 16:
			z := s.constFloat64(pstate, pstate.types.Types[TFLOAT64], 0)
			return s.entryNewValue2(pstate, ssa.OpComplexMake, t, z, z)
		default:
			s.Fatalf("bad sized complex type %v", t)
		}

	case t.IsString():
		return s.constEmptyString(pstate, t)
	case t.IsPtrShaped():
		return s.constNil(pstate, t)
	case t.IsBoolean():
		return s.constBool(pstate, false)
	case t.IsInterface():
		return s.constInterface(pstate, t)
	case t.IsSlice():
		return s.constSlice(pstate, t)
	case t.IsStruct():
		n := t.NumFields(pstate.types)
		v := s.entryNewValue0(pstate, ssa.StructMakeOp(t.NumFields(pstate.types)), t)
		for i := 0; i < n; i++ {
			v.AddArg(s.zeroVal(pstate, t.FieldType(pstate.types, i)))
		}
		return v
	case t.IsArray():
		switch t.NumElem(pstate.types) {
		case 0:
			return s.entryNewValue0(pstate, ssa.OpArrayMake0, t)
		case 1:
			return s.entryNewValue1(pstate, ssa.OpArrayMake1, t, s.zeroVal(pstate, t.Elem(pstate.types)))
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

func (pstate *PackageState) softfloatInit() {
	// Some of these operations get transformed by sfcall.
	pstate.softFloatOps = map[ssa.Op]sfRtCallDef{
		ssa.OpAdd32F: sfRtCallDef{pstate.sysfunc("fadd32"), TFLOAT32},
		ssa.OpAdd64F: sfRtCallDef{pstate.sysfunc("fadd64"), TFLOAT64},
		ssa.OpSub32F: sfRtCallDef{pstate.sysfunc("fadd32"), TFLOAT32},
		ssa.OpSub64F: sfRtCallDef{pstate.sysfunc("fadd64"), TFLOAT64},
		ssa.OpMul32F: sfRtCallDef{pstate.sysfunc("fmul32"), TFLOAT32},
		ssa.OpMul64F: sfRtCallDef{pstate.sysfunc("fmul64"), TFLOAT64},
		ssa.OpDiv32F: sfRtCallDef{pstate.sysfunc("fdiv32"), TFLOAT32},
		ssa.OpDiv64F: sfRtCallDef{pstate.sysfunc("fdiv64"), TFLOAT64},

		ssa.OpEq64F:      sfRtCallDef{pstate.sysfunc("feq64"), TBOOL},
		ssa.OpEq32F:      sfRtCallDef{pstate.sysfunc("feq32"), TBOOL},
		ssa.OpNeq64F:     sfRtCallDef{pstate.sysfunc("feq64"), TBOOL},
		ssa.OpNeq32F:     sfRtCallDef{pstate.sysfunc("feq32"), TBOOL},
		ssa.OpLess64F:    sfRtCallDef{pstate.sysfunc("fgt64"), TBOOL},
		ssa.OpLess32F:    sfRtCallDef{pstate.sysfunc("fgt32"), TBOOL},
		ssa.OpGreater64F: sfRtCallDef{pstate.sysfunc("fgt64"), TBOOL},
		ssa.OpGreater32F: sfRtCallDef{pstate.sysfunc("fgt32"), TBOOL},
		ssa.OpLeq64F:     sfRtCallDef{pstate.sysfunc("fge64"), TBOOL},
		ssa.OpLeq32F:     sfRtCallDef{pstate.sysfunc("fge32"), TBOOL},
		ssa.OpGeq64F:     sfRtCallDef{pstate.sysfunc("fge64"), TBOOL},
		ssa.OpGeq32F:     sfRtCallDef{pstate.sysfunc("fge32"), TBOOL},

		ssa.OpCvt32to32F:  sfRtCallDef{pstate.sysfunc("fint32to32"), TFLOAT32},
		ssa.OpCvt32Fto32:  sfRtCallDef{pstate.sysfunc("f32toint32"), TINT32},
		ssa.OpCvt64to32F:  sfRtCallDef{pstate.sysfunc("fint64to32"), TFLOAT32},
		ssa.OpCvt32Fto64:  sfRtCallDef{pstate.sysfunc("f32toint64"), TINT64},
		ssa.OpCvt64Uto32F: sfRtCallDef{pstate.sysfunc("fuint64to32"), TFLOAT32},
		ssa.OpCvt32Fto64U: sfRtCallDef{pstate.sysfunc("f32touint64"), TUINT64},
		ssa.OpCvt32to64F:  sfRtCallDef{pstate.sysfunc("fint32to64"), TFLOAT64},
		ssa.OpCvt64Fto32:  sfRtCallDef{pstate.sysfunc("f64toint32"), TINT32},
		ssa.OpCvt64to64F:  sfRtCallDef{pstate.sysfunc("fint64to64"), TFLOAT64},
		ssa.OpCvt64Fto64:  sfRtCallDef{pstate.sysfunc("f64toint64"), TINT64},
		ssa.OpCvt64Uto64F: sfRtCallDef{pstate.sysfunc("fuint64to64"), TFLOAT64},
		ssa.OpCvt64Fto64U: sfRtCallDef{pstate.sysfunc("f64touint64"), TUINT64},
		ssa.OpCvt32Fto64F: sfRtCallDef{pstate.sysfunc("f32to64"), TFLOAT64},
		ssa.OpCvt64Fto32F: sfRtCallDef{pstate.sysfunc("f64to32"), TFLOAT32},
	}
}

// TODO: do not emit sfcall if operation can be optimized to constant in later
// opt phase
func (s *state) sfcall(pstate *PackageState, op ssa.Op, args ...*ssa.Value) (*ssa.Value, bool) {
	if callDef, ok := pstate.softFloatOps[op]; ok {
		switch op {
		case ssa.OpLess32F,
			ssa.OpLess64F,
			ssa.OpLeq32F,
			ssa.OpLeq64F:
			args[0], args[1] = args[1], args[0]
		case ssa.OpSub32F,
			ssa.OpSub64F:
			args[1] = s.newValue1(s.ssaOp(pstate, OMINUS, pstate.types.Types[callDef.rtype]), args[1].Type, args[1])
		}

		result := s.rtcall(pstate, callDef.rtfn, true, []*types.Type{pstate.types.Types[callDef.rtype]}, args...)[0]
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

func (pstate *PackageState) init() {
	pstate.intrinsics = map[intrinsicKey]intrinsicBuilder{}

	var all []*sys.Arch
	var p4 []*sys.Arch
	var p8 []*sys.Arch
	for _, a := range pstate.sys.Archs {
		all = append(all, a)
		if a.PtrSize == 4 {
			p4 = append(p4, a)
		} else {
			p8 = append(p8, a)
		}
	}

	// add adds the intrinsic b for pkg.fn for the given list of architectures.
	add := func(pkg, fn string, b intrinsicBuilder, archs ...*sys.Arch) {
		for _, a := range archs {
			pstate.intrinsics[intrinsicKey{a, pkg, fn}] = b
		}
	}
	// addF does the same as add but operates on architecture families.
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
				pstate.intrinsics[intrinsicKey{a, pkg, fn}] = b
			}
		}
	}
	// alias defines pkg.fn = pkg2.fn2 for all architectures in archs for which pkg2.fn2 exists.
	alias := func(pkg, fn, pkg2, fn2 string, archs ...*sys.Arch) {
		for _, a := range archs {
			if b, ok := pstate.intrinsics[intrinsicKey{a, pkg2, fn2}]; ok {
				pstate.intrinsics[intrinsicKey{a, pkg, fn}] = b
			}
		}
	}

	/******** runtime ********/
	if !pstate.instrumenting {
		add("runtime", "slicebytetostringtmp",
			func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
				// Compiler frontend optimizations emit OARRAYBYTESTRTMP nodes
				// for the backend instead of slicebytetostringtmp calls
				// when not instrumenting.
				slice := args[0]
				ptr := s.newValue1(ssa.OpSlicePtr, s.f.Config.Types.BytePtr, slice)
				len := s.newValue1(ssa.OpSliceLen, pstate.types.Types[TINT], slice)
				return s.newValue2(ssa.OpStringMake, n.Type, ptr, len)
			},
			all...)
	}
	add("runtime", "KeepAlive",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			data := s.newValue1(ssa.OpIData, s.f.Config.Types.BytePtr, args[0])
			s.vars[&pstate.memVar] = s.newValue2(ssa.OpKeepAlive, pstate.types.TypeMem, data, s.mem(pstate))
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

	/******** runtime/internal/sys ********/
	addF("runtime/internal/sys", "Ctz32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz32, pstate.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("runtime/internal/sys", "Ctz64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz64, pstate.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("runtime/internal/sys", "Bswap32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBswap32, pstate.types.Types[TUINT32], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X)
	addF("runtime/internal/sys", "Bswap64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBswap64, pstate.types.Types[TUINT64], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X)

	/******** runtime/internal/atomic ********/
	addF("runtime/internal/atomic", "Load",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue2(ssa.OpAtomicLoad32, types.NewTuple(pstate.types.Types[TUINT32], pstate.types.TypeMem), args[0], s.mem(pstate))
			s.vars[&pstate.memVar] = s.newValue1(ssa.OpSelect1, pstate.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, pstate.types.Types[TUINT32], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Load64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue2(ssa.OpAtomicLoad64, types.NewTuple(pstate.types.Types[TUINT64], pstate.types.TypeMem), args[0], s.mem(pstate))
			s.vars[&pstate.memVar] = s.newValue1(ssa.OpSelect1, pstate.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, pstate.types.Types[TUINT64], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Loadp",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue2(ssa.OpAtomicLoadPtr, types.NewTuple(s.f.Config.Types.BytePtr, pstate.types.TypeMem), args[0], s.mem(pstate))
			s.vars[&pstate.memVar] = s.newValue1(ssa.OpSelect1, pstate.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, s.f.Config.Types.BytePtr, v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)

	addF("runtime/internal/atomic", "Store",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			s.vars[&pstate.memVar] = s.newValue3(ssa.OpAtomicStore32, pstate.types.TypeMem, args[0], args[1], s.mem(pstate))
			return nil
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Store64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			s.vars[&pstate.memVar] = s.newValue3(ssa.OpAtomicStore64, pstate.types.TypeMem, args[0], args[1], s.mem(pstate))
			return nil
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "StorepNoWB",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			s.vars[&pstate.memVar] = s.newValue3(ssa.OpAtomicStorePtrNoWB, pstate.types.TypeMem, args[0], args[1], s.mem(pstate))
			return nil
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64)

	addF("runtime/internal/atomic", "Xchg",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue3(ssa.OpAtomicExchange32, types.NewTuple(pstate.types.Types[TUINT32], pstate.types.TypeMem), args[0], args[1], s.mem(pstate))
			s.vars[&pstate.memVar] = s.newValue1(ssa.OpSelect1, pstate.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, pstate.types.Types[TUINT32], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Xchg64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue3(ssa.OpAtomicExchange64, types.NewTuple(pstate.types.Types[TUINT64], pstate.types.TypeMem), args[0], args[1], s.mem(pstate))
			s.vars[&pstate.memVar] = s.newValue1(ssa.OpSelect1, pstate.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, pstate.types.Types[TUINT64], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS64, sys.PPC64)

	addF("runtime/internal/atomic", "Xadd",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue3(ssa.OpAtomicAdd32, types.NewTuple(pstate.types.Types[TUINT32], pstate.types.TypeMem), args[0], args[1], s.mem(pstate))
			s.vars[&pstate.memVar] = s.newValue1(ssa.OpSelect1, pstate.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, pstate.types.Types[TUINT32], v)
		},
		sys.AMD64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Xadd64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue3(ssa.OpAtomicAdd64, types.NewTuple(pstate.types.Types[TUINT64], pstate.types.TypeMem), args[0], args[1], s.mem(pstate))
			s.vars[&pstate.memVar] = s.newValue1(ssa.OpSelect1, pstate.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, pstate.types.Types[TUINT64], v)
		},
		sys.AMD64, sys.S390X, sys.MIPS64, sys.PPC64)

	makeXaddARM64 := func(op0 ssa.Op, op1 ssa.Op, ty types.EType) func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			// Target Atomic feature is identified by dynamic detection
			addr := s.entryNewValue1A(pstate, ssa.OpAddr, pstate.types.Types[TBOOL].PtrTo(pstate.types), pstate.arm64SupportAtomics, s.sb)
			v := s.load(pstate, pstate.types.Types[TBOOL], addr)
			b := s.endBlock(pstate)
			b.Kind = ssa.BlockIf
			b.SetControl(v)
			bTrue := s.f.NewBlock(ssa.BlockPlain)
			bFalse := s.f.NewBlock(ssa.BlockPlain)
			bEnd := s.f.NewBlock(ssa.BlockPlain)
			b.AddEdgeTo(bTrue)
			b.AddEdgeTo(bFalse)
			b.Likely = ssa.BranchUnlikely // most machines don't have Atomics nowadays

			// We have atomic instructions - use it directly.
			s.startBlock(bTrue)
			v0 := s.newValue3(op1, types.NewTuple(pstate.types.Types[ty], pstate.types.TypeMem), args[0], args[1], s.mem(pstate))
			s.vars[&pstate.memVar] = s.newValue1(ssa.OpSelect1, pstate.types.TypeMem, v0)
			s.vars[n] = s.newValue1(ssa.OpSelect0, pstate.types.Types[ty], v0)
			s.endBlock(pstate).AddEdgeTo(bEnd)

			// Use original instruction sequence.
			s.startBlock(bFalse)
			v1 := s.newValue3(op0, types.NewTuple(pstate.types.Types[ty], pstate.types.TypeMem), args[0], args[1], s.mem(pstate))
			s.vars[&pstate.memVar] = s.newValue1(ssa.OpSelect1, pstate.types.TypeMem, v1)
			s.vars[n] = s.newValue1(ssa.OpSelect0, pstate.types.Types[ty], v1)
			s.endBlock(pstate).AddEdgeTo(bEnd)

			// Merge results.
			s.startBlock(bEnd)
			return s.variable(n, pstate.types.Types[ty])
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
			v := s.newValue4(ssa.OpAtomicCompareAndSwap32, types.NewTuple(pstate.types.Types[TBOOL], pstate.types.TypeMem), args[0], args[1], args[2], s.mem(pstate))
			s.vars[&pstate.memVar] = s.newValue1(ssa.OpSelect1, pstate.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, pstate.types.Types[TBOOL], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS, sys.MIPS64, sys.PPC64)
	addF("runtime/internal/atomic", "Cas64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			v := s.newValue4(ssa.OpAtomicCompareAndSwap64, types.NewTuple(pstate.types.Types[TBOOL], pstate.types.TypeMem), args[0], args[1], args[2], s.mem(pstate))
			s.vars[&pstate.memVar] = s.newValue1(ssa.OpSelect1, pstate.types.TypeMem, v)
			return s.newValue1(ssa.OpSelect0, pstate.types.Types[TBOOL], v)
		},
		sys.AMD64, sys.ARM64, sys.S390X, sys.MIPS64, sys.PPC64)

	addF("runtime/internal/atomic", "And8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			s.vars[&pstate.memVar] = s.newValue3(ssa.OpAtomicAnd8, pstate.types.TypeMem, args[0], args[1], s.mem(pstate))
			return nil
		},
		sys.AMD64, sys.ARM64, sys.MIPS, sys.PPC64)
	addF("runtime/internal/atomic", "Or8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			s.vars[&pstate.memVar] = s.newValue3(ssa.OpAtomicOr8, pstate.types.TypeMem, args[0], args[1], s.mem(pstate))
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

	/******** math ********/
	addF("math", "Sqrt",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpSqrt, pstate.types.Types[TFLOAT64], args[0])
		},
		sys.I386, sys.AMD64, sys.ARM, sys.ARM64, sys.MIPS, sys.MIPS64, sys.PPC64, sys.S390X)
	addF("math", "Trunc",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpTrunc, pstate.types.Types[TFLOAT64], args[0])
		},
		sys.ARM64, sys.PPC64, sys.S390X)
	addF("math", "Ceil",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCeil, pstate.types.Types[TFLOAT64], args[0])
		},
		sys.ARM64, sys.PPC64, sys.S390X)
	addF("math", "Floor",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpFloor, pstate.types.Types[TFLOAT64], args[0])
		},
		sys.ARM64, sys.PPC64, sys.S390X)
	addF("math", "Round",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpRound, pstate.types.Types[TFLOAT64], args[0])
		},
		sys.ARM64, sys.PPC64, sys.S390X)
	addF("math", "RoundToEven",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpRoundToEven, pstate.types.Types[TFLOAT64], args[0])
		},
		sys.S390X)
	addF("math", "Abs",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpAbs, pstate.types.Types[TFLOAT64], args[0])
		},
		sys.PPC64)
	addF("math", "Copysign",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue2(ssa.OpCopysign, pstate.types.Types[TFLOAT64], args[0], args[1])
		},
		sys.PPC64)

	makeRoundAMD64 := func(op ssa.Op) func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			addr := s.entryNewValue1A(pstate, ssa.OpAddr, pstate.types.Types[TBOOL].PtrTo(pstate.types), pstate.supportSSE41, s.sb)
			v := s.load(pstate, pstate.types.Types[TBOOL], addr)
			b := s.endBlock(pstate)
			b.Kind = ssa.BlockIf
			b.SetControl(v)
			bTrue := s.f.NewBlock(ssa.BlockPlain)
			bFalse := s.f.NewBlock(ssa.BlockPlain)
			bEnd := s.f.NewBlock(ssa.BlockPlain)
			b.AddEdgeTo(bTrue)
			b.AddEdgeTo(bFalse)
			b.Likely = ssa.BranchLikely // most machines have sse4.1 nowadays

			// We have the intrinsic - use it directly.
			s.startBlock(bTrue)
			s.vars[n] = s.newValue1(op, pstate.types.Types[TFLOAT64], args[0])
			s.endBlock(pstate).AddEdgeTo(bEnd)

			// Call the pure Go version.
			s.startBlock(bFalse)
			a := s.call(pstate, n, callNormal)
			s.vars[n] = s.load(pstate, pstate.types.Types[TFLOAT64], a)
			s.endBlock(pstate).AddEdgeTo(bEnd)

			// Merge results.
			s.startBlock(bEnd)
			return s.variable(n, pstate.types.Types[TFLOAT64])
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

	/******** math/bits ********/
	addF("math/bits", "TrailingZeros64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz64, pstate.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "TrailingZeros32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz32, pstate.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "TrailingZeros16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			x := s.newValue1(ssa.OpZeroExt16to32, pstate.types.Types[TUINT32], args[0])
			c := s.constInt32(pstate, pstate.types.Types[TUINT32], 1<<16)
			y := s.newValue2(ssa.OpOr32, pstate.types.Types[TUINT32], x, c)
			return s.newValue1(ssa.OpCtz32, pstate.types.Types[TINT], y)
		},
		sys.ARM, sys.MIPS)
	addF("math/bits", "TrailingZeros16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz16, pstate.types.Types[TINT], args[0])
		},
		sys.AMD64)
	addF("math/bits", "TrailingZeros16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			x := s.newValue1(ssa.OpZeroExt16to64, pstate.types.Types[TUINT64], args[0])
			c := s.constInt64(pstate, pstate.types.Types[TUINT64], 1<<16)
			y := s.newValue2(ssa.OpOr64, pstate.types.Types[TUINT64], x, c)
			return s.newValue1(ssa.OpCtz64, pstate.types.Types[TINT], y)
		},
		sys.ARM64, sys.S390X)
	addF("math/bits", "TrailingZeros8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			x := s.newValue1(ssa.OpZeroExt8to32, pstate.types.Types[TUINT32], args[0])
			c := s.constInt32(pstate, pstate.types.Types[TUINT32], 1<<8)
			y := s.newValue2(ssa.OpOr32, pstate.types.Types[TUINT32], x, c)
			return s.newValue1(ssa.OpCtz32, pstate.types.Types[TINT], y)
		},
		sys.ARM, sys.MIPS)
	addF("math/bits", "TrailingZeros8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpCtz8, pstate.types.Types[TINT], args[0])
		},
		sys.AMD64)
	addF("math/bits", "TrailingZeros8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			x := s.newValue1(ssa.OpZeroExt8to64, pstate.types.Types[TUINT64], args[0])
			c := s.constInt64(pstate, pstate.types.Types[TUINT64], 1<<8)
			y := s.newValue2(ssa.OpOr64, pstate.types.Types[TUINT64], x, c)
			return s.newValue1(ssa.OpCtz64, pstate.types.Types[TINT], y)
		},
		sys.ARM64, sys.S390X)
	alias("math/bits", "ReverseBytes64", "runtime/internal/sys", "Bswap64", all...)
	alias("math/bits", "ReverseBytes32", "runtime/internal/sys", "Bswap32", all...)
	// ReverseBytes inlines correctly, no need to intrinsify it.
	// ReverseBytes16 lowers to a rotate, no need for anything special here.
	addF("math/bits", "Len64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitLen64, pstate.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "Len32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitLen32, pstate.types.Types[TINT], args[0])
		},
		sys.AMD64)
	addF("math/bits", "Len32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			if s.config.PtrSize == 4 {
				return s.newValue1(ssa.OpBitLen32, pstate.types.Types[TINT], args[0])
			}
			x := s.newValue1(ssa.OpZeroExt32to64, pstate.types.Types[TUINT64], args[0])
			return s.newValue1(ssa.OpBitLen64, pstate.types.Types[TINT], x)
		},
		sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "Len16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			if s.config.PtrSize == 4 {
				x := s.newValue1(ssa.OpZeroExt16to32, pstate.types.Types[TUINT32], args[0])
				return s.newValue1(ssa.OpBitLen32, pstate.types.Types[TINT], x)
			}
			x := s.newValue1(ssa.OpZeroExt16to64, pstate.types.Types[TUINT64], args[0])
			return s.newValue1(ssa.OpBitLen64, pstate.types.Types[TINT], x)
		},
		sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "Len16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitLen16, pstate.types.Types[TINT], args[0])
		},
		sys.AMD64)
	addF("math/bits", "Len8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			if s.config.PtrSize == 4 {
				x := s.newValue1(ssa.OpZeroExt8to32, pstate.types.Types[TUINT32], args[0])
				return s.newValue1(ssa.OpBitLen32, pstate.types.Types[TINT], x)
			}
			x := s.newValue1(ssa.OpZeroExt8to64, pstate.types.Types[TUINT64], args[0])
			return s.newValue1(ssa.OpBitLen64, pstate.types.Types[TINT], x)
		},
		sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	addF("math/bits", "Len8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitLen8, pstate.types.Types[TINT], args[0])
		},
		sys.AMD64)
	addF("math/bits", "Len",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			if s.config.PtrSize == 4 {
				return s.newValue1(ssa.OpBitLen32, pstate.types.Types[TINT], args[0])
			}
			return s.newValue1(ssa.OpBitLen64, pstate.types.Types[TINT], args[0])
		},
		sys.AMD64, sys.ARM64, sys.ARM, sys.S390X, sys.MIPS, sys.PPC64)
	// LeadingZeros is handled because it trivially calls Len.
	addF("math/bits", "Reverse64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitRev64, pstate.types.Types[TINT], args[0])
		},
		sys.ARM64)
	addF("math/bits", "Reverse32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitRev32, pstate.types.Types[TINT], args[0])
		},
		sys.ARM64)
	addF("math/bits", "Reverse16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitRev16, pstate.types.Types[TINT], args[0])
		},
		sys.ARM64)
	addF("math/bits", "Reverse8",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpBitRev8, pstate.types.Types[TINT], args[0])
		},
		sys.ARM64)
	addF("math/bits", "Reverse",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			if s.config.PtrSize == 4 {
				return s.newValue1(ssa.OpBitRev32, pstate.types.Types[TINT], args[0])
			}
			return s.newValue1(ssa.OpBitRev64, pstate.types.Types[TINT], args[0])
		},
		sys.ARM64)
	makeOnesCountAMD64 := func(op64 ssa.Op, op32 ssa.Op) func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
		return func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			addr := s.entryNewValue1A(pstate, ssa.OpAddr, pstate.types.Types[TBOOL].PtrTo(pstate.types), pstate.supportPopcnt, s.sb)
			v := s.load(pstate, pstate.types.Types[TBOOL], addr)
			b := s.endBlock(pstate)
			b.Kind = ssa.BlockIf
			b.SetControl(v)
			bTrue := s.f.NewBlock(ssa.BlockPlain)
			bFalse := s.f.NewBlock(ssa.BlockPlain)
			bEnd := s.f.NewBlock(ssa.BlockPlain)
			b.AddEdgeTo(bTrue)
			b.AddEdgeTo(bFalse)
			b.Likely = ssa.BranchLikely // most machines have popcnt nowadays

			// We have the intrinsic - use it directly.
			s.startBlock(bTrue)
			op := op64
			if s.config.PtrSize == 4 {
				op = op32
			}
			s.vars[n] = s.newValue1(op, pstate.types.Types[TINT], args[0])
			s.endBlock(pstate).AddEdgeTo(bEnd)

			// Call the pure Go version.
			s.startBlock(bFalse)
			a := s.call(pstate, n, callNormal)
			s.vars[n] = s.load(pstate, pstate.types.Types[TINT], a)
			s.endBlock(pstate).AddEdgeTo(bEnd)

			// Merge results.
			s.startBlock(bEnd)
			return s.variable(n, pstate.types.Types[TINT])
		}
	}
	addF("math/bits", "OnesCount64",
		makeOnesCountAMD64(ssa.OpPopCount64, ssa.OpPopCount64),
		sys.AMD64)
	addF("math/bits", "OnesCount64",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpPopCount64, pstate.types.Types[TINT], args[0])
		},
		sys.PPC64, sys.ARM64)
	addF("math/bits", "OnesCount32",
		makeOnesCountAMD64(ssa.OpPopCount32, ssa.OpPopCount32),
		sys.AMD64)
	addF("math/bits", "OnesCount32",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpPopCount32, pstate.types.Types[TINT], args[0])
		},
		sys.PPC64, sys.ARM64)
	addF("math/bits", "OnesCount16",
		makeOnesCountAMD64(ssa.OpPopCount16, ssa.OpPopCount16),
		sys.AMD64)
	addF("math/bits", "OnesCount16",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue1(ssa.OpPopCount16, pstate.types.Types[TINT], args[0])
		},
		sys.ARM64)
	// Note: no OnesCount8, the Go implementation is faster - just a table load.
	addF("math/bits", "OnesCount",
		makeOnesCountAMD64(ssa.OpPopCount64, ssa.OpPopCount32),
		sys.AMD64)

	/******** sync/atomic ********/

	// Note: these are disabled by flag_race in findIntrinsic below.
	alias("sync/atomic", "LoadInt32", "runtime/internal/atomic", "Load", all...)
	alias("sync/atomic", "LoadInt64", "runtime/internal/atomic", "Load64", all...)
	alias("sync/atomic", "LoadPointer", "runtime/internal/atomic", "Loadp", all...)
	alias("sync/atomic", "LoadUint32", "runtime/internal/atomic", "Load", all...)
	alias("sync/atomic", "LoadUint64", "runtime/internal/atomic", "Load64", all...)
	alias("sync/atomic", "LoadUintptr", "runtime/internal/atomic", "Load", p4...)
	alias("sync/atomic", "LoadUintptr", "runtime/internal/atomic", "Load64", p8...)

	alias("sync/atomic", "StoreInt32", "runtime/internal/atomic", "Store", all...)
	alias("sync/atomic", "StoreInt64", "runtime/internal/atomic", "Store64", all...)
	// Note: not StorePointer, that needs a write barrier.  Same below for {CompareAnd}Swap.
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

	/******** math/big ********/
	add("math/big", "mulWW",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue2(ssa.OpMul64uhilo, types.NewTuple(pstate.types.Types[TUINT64], pstate.types.Types[TUINT64]), args[0], args[1])
		},
		pstate.sys.ArchAMD64, pstate.sys.ArchARM64)
	add("math/big", "divWW",
		func(s *state, n *Node, args []*ssa.Value) *ssa.Value {
			return s.newValue3(ssa.OpDiv128u, types.NewTuple(pstate.types.Types[TUINT64], pstate.types.Types[TUINT64]), args[0], args[1], args[2])
		},
		pstate.sys.ArchAMD64)
}

// findIntrinsic returns a function which builds the SSA equivalent of the
// function identified by the symbol sym.  If sym is not an intrinsic call, returns nil.
func (pstate *PackageState) findIntrinsic(sym *types.Sym) intrinsicBuilder {
	if pstate.ssa.IntrinsicsDisable {
		return nil
	}
	if sym == nil || sym.Pkg == nil {
		return nil
	}
	pkg := sym.Pkg.Path
	if sym.Pkg == pstate.localpkg {
		pkg = pstate.myimportpath
	}
	if pstate.flag_race && pkg == "sync/atomic" {
		// The race detector needs to be able to intercept these calls.
		// We can't intrinsify them.
		return nil
	}
	// Skip intrinsifying math functions (which may contain hard-float
	// instructions) when soft-float
	if pstate.thearch.SoftFloat && pkg == "math" {
		return nil
	}

	fn := sym.Name
	return pstate.intrinsics[intrinsicKey{pstate.thearch.LinkArch.Arch, pkg, fn}]
}

func (pstate *PackageState) isIntrinsicCall(n *Node) bool {
	if n == nil || n.Left == nil {
		return false
	}
	return pstate.findIntrinsic(n.Left.Sym) != nil
}

// intrinsicCall converts a call to a recognized intrinsic function into the intrinsic SSA operation.
func (s *state) intrinsicCall(pstate *PackageState, n *Node) *ssa.Value {
	v := pstate.findIntrinsic(n.Left.Sym)(s, n, s.intrinsicArgs(pstate, n))
	if pstate.ssa.IntrinsicsDebug > 0 {
		x := v
		if x == nil {
			x = s.mem(pstate)
		}
		if x.Op == ssa.OpSelect0 || x.Op == ssa.OpSelect1 {
			x = x.Args[0]
		}
		pstate.Warnl(n.Pos, "intrinsic substitution for %v with %s", n.Left.Sym.Name, x.LongString(pstate.ssa))
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
func (s *state) intrinsicArgs(pstate *PackageState, n *Node) []*ssa.Value {
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
			// Evaluate and store to "temporary".
			// Walk ensures these temporaries are dead outside of n.
			temps[l] = s.expr(pstate, r)
		case OINDREGSP:
			// Store a value to an argument slot.
			var v *ssa.Value
			if x, ok := temps[r]; ok {
				// This is a previously computed temporary.
				v = x
			} else {
				// This is an explicit value; evaluate it.
				v = s.expr(pstate, r)
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
func (s *state) call(pstate *PackageState, n *Node, k callKind) *ssa.Value {
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
		closure = s.expr(pstate, fn)
	case OCALLMETH:
		if fn.Op != ODOTMETH {
			pstate.Fatalf("OCALLMETH: n.Left not an ODOTMETH: %v", fn)
		}
		if k == callNormal {
			sym = fn.Sym
			break
		}
		// Make a name n2 for the function.
		// fn.Sym might be sync.(*Mutex).Unlock.
		// Make a PFUNC node out of that, then evaluate it.
		// We get back an SSA value representing &sync.(*Mutex).Unlock·f.
		// We can then pass that to defer or go.
		n2 := pstate.newnamel(fn.Pos, fn.Sym)
		n2.Name.Curfn = s.curfn
		n2.SetClass(PFUNC)
		n2.Pos = fn.Pos
		n2.Type = pstate.types.Types[TUINT8] // dummy type for a static closure. Could use runtime.funcval if we had it.
		closure = s.expr(pstate, n2)
	// Note: receiver is already assigned in n.List, so we don't
	// want to set it here.
	case OCALLINTER:
		if fn.Op != ODOTINTER {
			pstate.Fatalf("OCALLINTER: n.Left not an ODOTINTER: %v", fn.Op)
		}
		i := s.expr(pstate, fn.Left)
		itab := s.newValue1(ssa.OpITab, pstate.types.Types[TUINTPTR], i)
		s.nilCheck(pstate, itab)
		itabidx := fn.Xoffset + 2*int64(pstate.Widthptr) + 8 // offset of fun field in runtime.itab
		itab = s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.UintptrPtr, itabidx, itab)
		if k == callNormal {
			codeptr = s.load(pstate, pstate.types.Types[TUINTPTR], itab)
		} else {
			closure = itab
		}
		rcvr = s.newValue1(ssa.OpIData, pstate.types.Types[TUINTPTR], i)
	}
	pstate.dowidth(fn.Type)
	stksize := fn.Type.ArgWidth(pstate.types) // includes receiver

	// Run all argument assignments. The arg slots have already
	// been offset by the appropriate amount (+2*widthptr for go/defer,
	// +widthptr for interface calls).
	// For OCALLMETH, the receiver is set in these statements.
	s.stmtList(pstate, n.List)

	// Set receiver (for interface calls)
	if rcvr != nil {
		argStart := pstate.Ctxt.FixedFrameSize()
		if k != callNormal {
			argStart += int64(2 * pstate.Widthptr)
		}
		addr := s.constOffPtrSP(pstate, s.f.Config.Types.UintptrPtr, argStart)
		s.store(pstate, pstate.types.Types[TUINTPTR], addr, rcvr)
	}

	// Defer/go args
	if k != callNormal {
		// Write argsize and closure (args to Newproc/Deferproc).
		argStart := pstate.Ctxt.FixedFrameSize()
		argsize := s.constInt32(pstate, pstate.types.Types[TUINT32], int32(stksize))
		addr := s.constOffPtrSP(pstate, s.f.Config.Types.UInt32Ptr, argStart)
		s.store(pstate, pstate.types.Types[TUINT32], addr, argsize)
		addr = s.constOffPtrSP(pstate, s.f.Config.Types.UintptrPtr, argStart+int64(pstate.Widthptr))
		s.store(pstate, pstate.types.Types[TUINTPTR], addr, closure)
		stksize += 2 * int64(pstate.Widthptr)
	}

	// call target
	var call *ssa.Value
	switch {
	case k == callDefer:
		call = s.newValue1A(ssa.OpStaticCall, pstate.types.TypeMem, pstate.Deferproc, s.mem(pstate))
	case k == callGo:
		call = s.newValue1A(ssa.OpStaticCall, pstate.types.TypeMem, pstate.Newproc, s.mem(pstate))
	case closure != nil:
		// rawLoad because loading the code pointer from a
		// closure is always safe, but IsSanitizerSafeAddr
		// can't always figure that out currently, and it's
		// critical that we not clobber any arguments already
		// stored onto the stack.
		codeptr = s.rawLoad(pstate, pstate.types.Types[TUINTPTR], closure)
		call = s.newValue3(ssa.OpClosureCall, pstate.types.TypeMem, codeptr, closure, s.mem(pstate))
	case codeptr != nil:
		call = s.newValue2(ssa.OpInterCall, pstate.types.TypeMem, codeptr, s.mem(pstate))
	case sym != nil:
		call = s.newValue1A(ssa.OpStaticCall, pstate.types.TypeMem, sym.Linksym(pstate.types), s.mem(pstate))
	default:
		pstate.Fatalf("bad call type %v %v", n.Op, n)
	}
	call.AuxInt = stksize // Call operations carry the argsize of the callee along with them
	s.vars[&pstate.memVar] = call

	// Finish block for defers
	if k == callDefer {
		b := s.endBlock(pstate)
		b.Kind = ssa.BlockDefer
		b.SetControl(call)
		bNext := s.f.NewBlock(ssa.BlockPlain)
		b.AddEdgeTo(bNext)
		// Add recover edge to exit code.
		r := s.f.NewBlock(ssa.BlockPlain)
		s.startBlock(r)
		s.exit(pstate)
		b.AddEdgeTo(r)
		b.Likely = ssa.BranchLikely
		s.startBlock(bNext)
	}

	res := n.Left.Type.Results(pstate.types)
	if res.NumFields(pstate.types) == 0 || k != callNormal {
		// call has no return value. Continue with the next statement.
		return nil
	}
	fp := res.Field(pstate.types, 0)
	return s.constOffPtrSP(pstate, pstate.types.NewPtr(fp.Type), fp.Offset+pstate.Ctxt.FixedFrameSize())
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
func (s *state) addr(pstate *PackageState, n *Node, bounded bool) *ssa.Value {
	t := pstate.types.NewPtr(n.Type)
	switch n.Op {
	case ONAME:
		switch n.Class() {
		case PEXTERN:
			// global variable
			v := s.entryNewValue1A(pstate, ssa.OpAddr, t, n.Sym.Linksym(pstate.types), s.sb)
			// TODO: Make OpAddr use AuxInt as well as Aux.
			if n.Xoffset != 0 {
				v = s.entryNewValue1I(pstate, ssa.OpOffPtr, v.Type, n.Xoffset, v)
			}
			return v
		case PPARAM:
			// parameter slot
			v := s.decladdrs[n]
			if v != nil {
				return v
			}
			if n == pstate.nodfp {
				// Special arg that points to the frame pointer (Used by ORECOVER).
				return s.entryNewValue1A(pstate, ssa.OpAddr, t, n, s.sp)
			}
			s.Fatalf("addr of undeclared ONAME %v. declared: %v", n, s.decladdrs)
			return nil
		case PAUTO:
			return s.newValue1Apos(ssa.OpAddr, t, n, s.sp, !n.IsAutoTmp())
		case PPARAMOUT: // Same as PAUTO -- cannot generate LEA early.
			// ensure that we reuse symbols for out parameters so
			// that cse works on their addresses
			return s.newValue1A(ssa.OpAddr, t, n, s.sp)
		default:
			s.Fatalf("variable address class %v not implemented", n.Class())
			return nil
		}
	case OINDREGSP:
		// indirect off REGSP
		// used for storing/loading arguments/returns to/from callees
		return s.constOffPtrSP(pstate, t, n.Xoffset)
	case OINDEX:
		if n.Left.Type.IsSlice() {
			a := s.expr(pstate, n.Left)
			i := s.expr(pstate, n.Right)
			i = s.extendIndex(pstate, i, pstate.panicindex)
			len := s.newValue1(ssa.OpSliceLen, pstate.types.Types[TINT], a)
			if !n.Bounded() {
				s.boundsCheck(pstate, i, len)
			}
			p := s.newValue1(ssa.OpSlicePtr, t, a)
			return s.newValue2(ssa.OpPtrIndex, t, p, i)
		} else { // array
			a := s.addr(pstate, n.Left, bounded)
			i := s.expr(pstate, n.Right)
			i = s.extendIndex(pstate, i, pstate.panicindex)
			len := s.constInt(pstate, pstate.types.Types[TINT], n.Left.Type.NumElem(pstate.types))
			if !n.Bounded() {
				s.boundsCheck(pstate, i, len)
			}
			return s.newValue2(ssa.OpPtrIndex, pstate.types.NewPtr(n.Left.Type.Elem(pstate.types)), a, i)
		}
	case OIND:
		return s.exprPtr(pstate, n.Left, bounded, n.Pos)
	case ODOT:
		p := s.addr(pstate, n.Left, bounded)
		return s.newValue1I(ssa.OpOffPtr, t, n.Xoffset, p)
	case ODOTPTR:
		p := s.exprPtr(pstate, n.Left, bounded, n.Pos)
		return s.newValue1I(ssa.OpOffPtr, t, n.Xoffset, p)
	case OCLOSUREVAR:
		return s.newValue1I(ssa.OpOffPtr, t, n.Xoffset,
			s.entryNewValue0(pstate, ssa.OpGetClosurePtr, s.f.Config.Types.BytePtr))
	case OCONVNOP:
		addr := s.addr(pstate, n.Left, bounded)
		return s.newValue1(ssa.OpCopy, t, addr) // ensure that addr has the right type
	case OCALLFUNC, OCALLINTER, OCALLMETH:
		return s.call(pstate, n, callNormal)
	case ODOTTYPE:
		v, _ := s.dottype(pstate, n, false)
		if v.Op != ssa.OpLoad {
			s.Fatalf("dottype of non-load")
		}
		if v.Args[1] != s.mem(pstate) {
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
func (s *state) canSSA(pstate *PackageState, n *Node) bool {
	if pstate.Debug['N'] != 0 {
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
		pstate.Fatalf("canSSA of PAUTOHEAP %v", n)
	}
	switch n.Class() {
	case PEXTERN:
		return false
	case PPARAMOUT:
		if s.hasdefer {
			// TODO: handle this case? Named return values must be
			// in memory so that the deferred function can see them.
			// Maybe do: if !strings.HasPrefix(n.String(), "~") { return false }
			// Or maybe not, see issue 18860.  Even unnamed return values
			// must be written back so if a defer recovers, the caller can see them.
			return false
		}
		if s.cgoUnsafeArgs {
			// Cgo effectively takes the address of all result args,
			// but the compiler can't see that.
			return false
		}
	}
	if n.Class() == PPARAM && n.Sym != nil && n.Sym.Name == ".this" {
		// wrappers generated by genwrapper need to update
		// the .this pointer in place.
		// TODO: treat as a PPARMOUT?
		return false
	}
	return pstate.canSSAType(n.Type)
	// TODO: try to make more variables SSAable?
}

// canSSA reports whether variables of type t are SSA-able.
func (pstate *PackageState) canSSAType(t *types.Type) bool {
	pstate.dowidth(t)
	if t.Width > int64(4*pstate.Widthptr) {
		// 4*Widthptr is an arbitrary constant. We want it
		// to be at least 3*Widthptr so slices can be registerized.
		// Too big and we'll introduce too much register pressure.
		return false
	}
	switch t.Etype {
	case TARRAY:
		// We can't do larger arrays because dynamic indexing is
		// not supported on SSA variables.
		// TODO: allow if all indexes are constant.
		if t.NumElem(pstate.types) <= 1 {
			return pstate.canSSAType(t.Elem(pstate.types))
		}
		return false
	case TSTRUCT:
		if t.NumFields(pstate.types) > ssa.MaxStruct {
			return false
		}
		for _, t1 := range t.Fields(pstate.types).Slice() {
			if !pstate.canSSAType(t1.Type) {
				return false
			}
		}
		return true
	default:
		return true
	}
}

// exprPtr evaluates n to a pointer and nil-checks it.
func (s *state) exprPtr(pstate *PackageState, n *Node, bounded bool, lineno src.XPos) *ssa.Value {
	p := s.expr(pstate, n)
	if bounded || n.NonNil() {
		if s.f.Frontend().Debug_checknil() && lineno.Line() > 1 {
			s.f.Warnl(lineno, "removed nil check")
		}
		return p
	}
	s.nilCheck(pstate, p)
	return p
}

// nilCheck generates nil pointer checking code.
// Used only for automatically inserted nil checks,
// not for user code like 'x != nil'.
func (s *state) nilCheck(pstate *PackageState, ptr *ssa.Value) {
	if pstate.disable_checknil != 0 || s.curfn.Func.NilCheckDisabled() {
		return
	}
	s.newValue2(ssa.OpNilCheck, pstate.types.TypeVoid, ptr, s.mem(pstate))
}

// boundsCheck generates bounds checking code. Checks if 0 <= idx < len, branches to exit if not.
// Starts a new block on return.
// idx is already converted to full int width.
func (s *state) boundsCheck(pstate *PackageState, idx, len *ssa.Value) {
	if pstate.Debug['B'] != 0 {
		return
	}

	// bounds check
	cmp := s.newValue2(ssa.OpIsInBounds, pstate.types.Types[TBOOL], idx, len)
	s.check(pstate, cmp, pstate.panicindex)
}

// sliceBoundsCheck generates slice bounds checking code. Checks if 0 <= idx <= len, branches to exit if not.
// Starts a new block on return.
// idx and len are already converted to full int width.
func (s *state) sliceBoundsCheck(pstate *PackageState, idx, len *ssa.Value) {
	if pstate.Debug['B'] != 0 {
		return
	}

	// bounds check
	cmp := s.newValue2(ssa.OpIsSliceInBounds, pstate.types.Types[TBOOL], idx, len)
	s.check(pstate, cmp, pstate.panicslice)
}

// If cmp (a bool) is false, panic using the given function.
func (s *state) check(pstate *PackageState, cmp *ssa.Value, fn *obj.LSym) {
	b := s.endBlock(pstate)
	b.Kind = ssa.BlockIf
	b.SetControl(cmp)
	b.Likely = ssa.BranchLikely
	bNext := s.f.NewBlock(ssa.BlockPlain)
	line := s.peekPos()
	pos := pstate.Ctxt.PosTable.Pos(line)
	fl := funcLine{f: fn, base: pos.Base(), line: pos.Line()}
	bPanic := s.panics[fl]
	if bPanic == nil {
		bPanic = s.f.NewBlock(ssa.BlockPlain)
		s.panics[fl] = bPanic
		s.startBlock(bPanic)
		// The panic call takes/returns memory to ensure that the right
		// memory state is observed if the panic happens.
		s.rtcall(pstate, fn, false, nil)
	}
	b.AddEdgeTo(bNext)
	b.AddEdgeTo(bPanic)
	s.startBlock(bNext)
}

func (s *state) intDivide(pstate *PackageState, n *Node, a, b *ssa.Value) *ssa.Value {
	needcheck := true
	switch b.Op {
	case ssa.OpConst8, ssa.OpConst16, ssa.OpConst32, ssa.OpConst64:
		if b.AuxInt != 0 {
			needcheck = false
		}
	}
	if needcheck {
		// do a size-appropriate check for zero
		cmp := s.newValue2(s.ssaOp(pstate, ONE, n.Type), pstate.types.Types[TBOOL], b, s.zeroVal(pstate, n.Type))
		s.check(pstate, cmp, pstate.panicdivide)
	}
	return s.newValue2(s.ssaOp(pstate, n.Op, n.Type), a.Type, a, b)
}

// rtcall issues a call to the given runtime function fn with the listed args.
// Returns a slice of results of the given result types.
// The call is added to the end of the current block.
// If returns is false, the block is marked as an exit block.
func (s *state) rtcall(pstate *PackageState, fn *obj.LSym, returns bool, results []*types.Type, args ...*ssa.Value) []*ssa.Value {
	// Write args to the stack
	off := pstate.Ctxt.FixedFrameSize()
	for _, arg := range args {
		t := arg.Type
		off = pstate.Rnd(off, t.Alignment(pstate.types))
		ptr := s.constOffPtrSP(pstate, t.PtrTo(pstate.types), off)
		size := t.Size(pstate.types)
		s.store(pstate, t, ptr, arg)
		off += size
	}
	off = pstate.Rnd(off, int64(pstate.Widthreg))

	// Issue call
	call := s.newValue1A(ssa.OpStaticCall, pstate.types.TypeMem, fn, s.mem(pstate))
	s.vars[&pstate.memVar] = call

	if !returns {
		// Finish block
		b := s.endBlock(pstate)
		b.Kind = ssa.BlockExit
		b.SetControl(call)
		call.AuxInt = off - pstate.Ctxt.FixedFrameSize()
		if len(results) > 0 {
			pstate.Fatalf("panic call can't have results")
		}
		return nil
	}

	// Load results
	res := make([]*ssa.Value, len(results))
	for i, t := range results {
		off = pstate.Rnd(off, t.Alignment(pstate.types))
		ptr := s.constOffPtrSP(pstate, pstate.types.NewPtr(t), off)
		res[i] = s.load(pstate, t, ptr)
		off += t.Size(pstate.types)
	}
	off = pstate.Rnd(off, int64(pstate.Widthptr))

	// Remember how much callee stack space we needed.
	call.AuxInt = off

	return res
}

// do *left = right for type t.
func (s *state) storeType(pstate *PackageState, t *types.Type, left, right *ssa.Value, skip skipMask, leftIsStmt bool) {
	s.instrument(pstate, t, left, true)

	if skip == 0 && (!pstate.types.Haspointers(t) || ssa.IsStackAddr(left)) {
		// Known to not have write barrier. Store the whole type.
		s.vars[&pstate.memVar] = s.newValue3Apos(ssa.OpStore, pstate.types.TypeMem, t, left, right, s.mem(pstate), leftIsStmt)
		return
	}

	// store scalar fields first, so write barrier stores for
	// pointer fields can be grouped together, and scalar values
	// don't need to be live across the write barrier call.
	// TODO: if the writebarrier pass knows how to reorder stores,
	// we can do a single store here as long as skip==0.
	s.storeTypeScalars(pstate, t, left, right, skip)
	if skip&skipPtr == 0 && pstate.types.Haspointers(t) {
		s.storeTypePtrs(pstate, t, left, right)
	}
}

// do *left = right for all scalar (non-pointer) parts of t.
func (s *state) storeTypeScalars(pstate *PackageState, t *types.Type, left, right *ssa.Value, skip skipMask) {
	switch {
	case t.IsBoolean() || t.IsInteger() || t.IsFloat() || t.IsComplex():
		s.store(pstate, t, left, right)
	case t.IsPtrShaped():
	// no scalar fields.
	case t.IsString():
		if skip&skipLen != 0 {
			return
		}
		len := s.newValue1(ssa.OpStringLen, pstate.types.Types[TINT], right)
		lenAddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.IntPtr, s.config.PtrSize, left)
		s.store(pstate, pstate.types.Types[TINT], lenAddr, len)
	case t.IsSlice():
		if skip&skipLen == 0 {
			len := s.newValue1(ssa.OpSliceLen, pstate.types.Types[TINT], right)
			lenAddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.IntPtr, s.config.PtrSize, left)
			s.store(pstate, pstate.types.Types[TINT], lenAddr, len)
		}
		if skip&skipCap == 0 {
			cap := s.newValue1(ssa.OpSliceCap, pstate.types.Types[TINT], right)
			capAddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.IntPtr, 2*s.config.PtrSize, left)
			s.store(pstate, pstate.types.Types[TINT], capAddr, cap)
		}
	case t.IsInterface():
		// itab field doesn't need a write barrier (even though it is a pointer).
		itab := s.newValue1(ssa.OpITab, s.f.Config.Types.BytePtr, right)
		s.store(pstate, pstate.types.Types[TUINTPTR], left, itab)
	case t.IsStruct():
		n := t.NumFields(pstate.types)
		for i := 0; i < n; i++ {
			ft := t.FieldType(pstate.types, i)
			addr := s.newValue1I(ssa.OpOffPtr, ft.PtrTo(pstate.types), t.FieldOff(pstate.types, i), left)
			val := s.newValue1I(ssa.OpStructSelect, ft, int64(i), right)
			s.storeTypeScalars(pstate, ft, addr, val, 0)
		}
	case t.IsArray() && t.NumElem(pstate.types) == 0:
	// nothing
	case t.IsArray() && t.NumElem(pstate.types) == 1:
		s.storeTypeScalars(pstate, t.Elem(pstate.types), left, s.newValue1I(ssa.OpArraySelect, t.Elem(pstate.types), 0, right), 0)
	default:
		s.Fatalf("bad write barrier type %v", t)
	}
}

// do *left = right for all pointer parts of t.
func (s *state) storeTypePtrs(pstate *PackageState, t *types.Type, left, right *ssa.Value) {
	switch {
	case t.IsPtrShaped():
		s.store(pstate, t, left, right)
	case t.IsString():
		ptr := s.newValue1(ssa.OpStringPtr, s.f.Config.Types.BytePtr, right)
		s.store(pstate, s.f.Config.Types.BytePtr, left, ptr)
	case t.IsSlice():
		elType := pstate.types.NewPtr(t.Elem(pstate.types))
		ptr := s.newValue1(ssa.OpSlicePtr, elType, right)
		s.store(pstate, elType, left, ptr)
	case t.IsInterface():
		// itab field is treated as a scalar.
		idata := s.newValue1(ssa.OpIData, s.f.Config.Types.BytePtr, right)
		idataAddr := s.newValue1I(ssa.OpOffPtr, s.f.Config.Types.BytePtrPtr, s.config.PtrSize, left)
		s.store(pstate, s.f.Config.Types.BytePtr, idataAddr, idata)
	case t.IsStruct():
		n := t.NumFields(pstate.types)
		for i := 0; i < n; i++ {
			ft := t.FieldType(pstate.types, i)
			if !pstate.types.Haspointers(ft) {
				continue
			}
			addr := s.newValue1I(ssa.OpOffPtr, ft.PtrTo(pstate.types), t.FieldOff(pstate.types, i), left)
			val := s.newValue1I(ssa.OpStructSelect, ft, int64(i), right)
			s.storeTypePtrs(pstate, ft, addr, val)
		}
	case t.IsArray() && t.NumElem(pstate.types) == 0:
	// nothing
	case t.IsArray() && t.NumElem(pstate.types) == 1:
		s.storeTypePtrs(pstate, t.Elem(pstate.types), left, s.newValue1I(ssa.OpArraySelect, t.Elem(pstate.types), 0, right))
	default:
		s.Fatalf("bad write barrier type %v", t)
	}
}

// slice computes the slice v[i:j:k] and returns ptr, len, and cap of result.
// i,j,k may be nil, in which case they are set to their default value.
// t is a slice, ptr to array, or string type.
func (s *state) slice(pstate *PackageState, t *types.Type, v, i, j, k *ssa.Value) (p, l, c *ssa.Value) {
	var elemtype *types.Type
	var ptrtype *types.Type
	var ptr *ssa.Value
	var len *ssa.Value
	var cap *ssa.Value
	zero := s.constInt(pstate, pstate.types.Types[TINT], 0)
	switch {
	case t.IsSlice():
		elemtype = t.Elem(pstate.types)
		ptrtype = pstate.types.NewPtr(elemtype)
		ptr = s.newValue1(ssa.OpSlicePtr, ptrtype, v)
		len = s.newValue1(ssa.OpSliceLen, pstate.types.Types[TINT], v)
		cap = s.newValue1(ssa.OpSliceCap, pstate.types.Types[TINT], v)
	case t.IsString():
		elemtype = pstate.types.Types[TUINT8]
		ptrtype = pstate.types.NewPtr(elemtype)
		ptr = s.newValue1(ssa.OpStringPtr, ptrtype, v)
		len = s.newValue1(ssa.OpStringLen, pstate.types.Types[TINT], v)
		cap = len
	case t.IsPtr():
		if !t.Elem(pstate.types).IsArray() {
			s.Fatalf("bad ptr to array in slice %v\n", t)
		}
		elemtype = t.Elem(pstate.types).Elem(pstate.types)
		ptrtype = pstate.types.NewPtr(elemtype)
		s.nilCheck(pstate, v)
		ptr = v
		len = s.constInt(pstate, pstate.types.Types[TINT], t.Elem(pstate.types).NumElem(pstate.types))
		cap = len
	default:
		s.Fatalf("bad type in slice %v\n", t)
	}

	// Set default values
	if i == nil {
		i = zero
	}
	if j == nil {
		j = len
	}
	if k == nil {
		k = cap
	}

	// Panic if slice indices are not in bounds.
	s.sliceBoundsCheck(pstate, i, j)
	if j != k {
		s.sliceBoundsCheck(pstate, j, k)
	}
	if k != cap {
		s.sliceBoundsCheck(pstate, k, cap)
	}

	// Generate the following code assuming that indexes are in bounds.
	// The masking is to make sure that we don't generate a slice
	// that points to the next object in memory.
	// rlen = j - i
	// rcap = k - i
	// delta = i * elemsize
	// rptr = p + delta&mask(rcap)
	// result = (SliceMake rptr rlen rcap)
	// where mask(x) is 0 if x==0 and -1 if x>0.
	subOp := s.ssaOp(pstate, OSUB, pstate.types.Types[TINT])
	mulOp := s.ssaOp(pstate, OMUL, pstate.types.Types[TINT])
	andOp := s.ssaOp(pstate, OAND, pstate.types.Types[TINT])
	rlen := s.newValue2(subOp, pstate.types.Types[TINT], j, i)
	var rcap *ssa.Value
	switch {
	case t.IsString():
		// Capacity of the result is unimportant. However, we use
		// rcap to test if we've generated a zero-length slice.
		// Use length of strings for that.
		rcap = rlen
	case j == k:
		rcap = rlen
	default:
		rcap = s.newValue2(subOp, pstate.types.Types[TINT], k, i)
	}

	var rptr *ssa.Value
	if (i.Op == ssa.OpConst64 || i.Op == ssa.OpConst32) && i.AuxInt == 0 {
		// No pointer arithmetic necessary.
		rptr = ptr
	} else {
		// delta = # of bytes to offset pointer by.
		delta := s.newValue2(mulOp, pstate.types.Types[TINT], i, s.constInt(pstate, pstate.types.Types[TINT], elemtype.Width))
		// If we're slicing to the point where the capacity is zero,
		// zero out the delta.
		mask := s.newValue1(ssa.OpSlicemask, pstate.types.Types[TINT], rcap)
		delta = s.newValue2(andOp, pstate.types.Types[TINT], delta, mask)
		// Compute rptr = ptr + delta
		rptr = s.newValue2(ssa.OpAddPtr, ptrtype, ptr, delta)
	}

	return rptr, rlen, rcap
}

type u642fcvtTab struct {
	geq, cvt2F, and, rsh, or, add ssa.Op
	one                           func(*state, *types.Type, int64) *ssa.Value
}

func (s *state) uint64Tofloat64(pstate *PackageState, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.uint64Tofloat(pstate, &pstate.u64_f64, n, x, ft, tt)
}

func (s *state) uint64Tofloat32(pstate *PackageState, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.uint64Tofloat(pstate, &pstate.u64_f32, n, x, ft, tt)
}

func (s *state) uint64Tofloat(pstate *PackageState, cvttab *u642fcvtTab, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	// if x >= 0 {
	//    result = (floatY) x
	// } else {
	// 	  y = uintX(x) ; y = x & 1
	// 	  z = uintX(x) ; z = z >> 1
	// 	  z = z >> 1
	// 	  z = z | y
	// 	  result = floatY(z)
	// 	  result = result + result
	// }
	//
	// Code borrowed from old code generator.
	// What's going on: large 64-bit "unsigned" looks like
	// negative number to hardware's integer-to-float
	// conversion. However, because the mantissa is only
	// 63 bits, we don't need the LSB, so instead we do an
	// unsigned right shift (divide by two), convert, and
	// double. However, before we do that, we need to be
	// sure that we do not lose a "1" if that made the
	// difference in the resulting rounding. Therefore, we
	// preserve it, and OR (not ADD) it back in. The case
	// that matters is when the eleven discarded bits are
	// equal to 10000000001; that rounds up, and the 1 cannot
	// be lost else it would round down if the LSB of the
	// candidate mantissa is 0.
	cmp := s.newValue2(cvttab.geq, pstate.types.Types[TBOOL], x, s.zeroVal(pstate, ft))
	b := s.endBlock(pstate)
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
	s.endBlock(pstate)
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
	s.endBlock(pstate)
	bElse.AddEdgeTo(bAfter)

	s.startBlock(bAfter)
	return s.variable(n, n.Type)
}

type u322fcvtTab struct {
	cvtI2F, cvtF2F ssa.Op
}

func (s *state) uint32Tofloat64(pstate *PackageState, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.uint32Tofloat(pstate, &pstate.u32_f64, n, x, ft, tt)
}

func (s *state) uint32Tofloat32(pstate *PackageState, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.uint32Tofloat(pstate, &pstate.u32_f32, n, x, ft, tt)
}

func (s *state) uint32Tofloat(pstate *PackageState, cvttab *u322fcvtTab, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	// if x >= 0 {
	// 	result = floatY(x)
	// } else {
	// 	result = floatY(float64(x) + (1<<32))
	// }
	cmp := s.newValue2(ssa.OpGeq32, pstate.types.Types[TBOOL], x, s.zeroVal(pstate, ft))
	b := s.endBlock(pstate)
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
	s.endBlock(pstate)
	bThen.AddEdgeTo(bAfter)

	b.AddEdgeTo(bElse)
	s.startBlock(bElse)
	a1 := s.newValue1(ssa.OpCvt32to64F, pstate.types.Types[TFLOAT64], x)
	twoToThe32 := s.constFloat64(pstate, pstate.types.Types[TFLOAT64], float64(1<<32))
	a2 := s.newValue2(ssa.OpAdd64F, pstate.types.Types[TFLOAT64], a1, twoToThe32)
	a3 := s.newValue1(cvttab.cvtF2F, tt, a2)

	s.vars[n] = a3
	s.endBlock(pstate)
	bElse.AddEdgeTo(bAfter)

	s.startBlock(bAfter)
	return s.variable(n, n.Type)
}

// referenceTypeBuiltin generates code for the len/cap builtins for maps and channels.
func (s *state) referenceTypeBuiltin(pstate *PackageState, n *Node, x *ssa.Value) *ssa.Value {
	if !n.Left.Type.IsMap() && !n.Left.Type.IsChan() {
		s.Fatalf("node must be a map or a channel")
	}
	// if n == nil {
	//   return 0
	// } else {
	//   // len
	//   return *((*int)n)
	//   // cap
	//   return *(((*int)n)+1)
	// }
	lenType := n.Type
	nilValue := s.constNil(pstate, pstate.types.Types[TUINTPTR])
	cmp := s.newValue2(ssa.OpEqPtr, pstate.types.Types[TBOOL], x, nilValue)
	b := s.endBlock(pstate)
	b.Kind = ssa.BlockIf
	b.SetControl(cmp)
	b.Likely = ssa.BranchUnlikely

	bThen := s.f.NewBlock(ssa.BlockPlain)
	bElse := s.f.NewBlock(ssa.BlockPlain)
	bAfter := s.f.NewBlock(ssa.BlockPlain)

	// length/capacity of a nil map/chan is zero
	b.AddEdgeTo(bThen)
	s.startBlock(bThen)
	s.vars[n] = s.zeroVal(pstate, lenType)
	s.endBlock(pstate)
	bThen.AddEdgeTo(bAfter)

	b.AddEdgeTo(bElse)
	s.startBlock(bElse)
	switch n.Op {
	case OLEN:
		// length is stored in the first word for map/chan
		s.vars[n] = s.load(pstate, lenType, x)
	case OCAP:
		// capacity is stored in the second word for chan
		sw := s.newValue1I(ssa.OpOffPtr, lenType.PtrTo(pstate.types), lenType.Width, x)
		s.vars[n] = s.load(pstate, lenType, sw)
	default:
		s.Fatalf("op must be OLEN or OCAP")
	}
	s.endBlock(pstate)
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

func (s *state) float32ToUint64(pstate *PackageState, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.floatToUint(pstate, &pstate.f32_u64, n, x, ft, tt)
}
func (s *state) float64ToUint64(pstate *PackageState, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.floatToUint(pstate, &pstate.f64_u64, n, x, ft, tt)
}

func (s *state) float32ToUint32(pstate *PackageState, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.floatToUint(pstate, &pstate.f32_u32, n, x, ft, tt)
}

func (s *state) float64ToUint32(pstate *PackageState, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	return s.floatToUint(pstate, &pstate.f64_u32, n, x, ft, tt)
}

func (s *state) floatToUint(pstate *PackageState, cvttab *f2uCvtTab, n *Node, x *ssa.Value, ft, tt *types.Type) *ssa.Value {
	// cutoff:=1<<(intY_Size-1)
	// if x < floatX(cutoff) {
	// 	result = uintY(x)
	// } else {
	// 	y = x - floatX(cutoff)
	// 	z = uintY(y)
	// 	result = z | -(cutoff)
	// }
	cutoff := cvttab.floatValue(s, ft, float64(cvttab.cutoff))
	cmp := s.newValue2(cvttab.ltf, pstate.types.Types[TBOOL], x, cutoff)
	b := s.endBlock(pstate)
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
	s.endBlock(pstate)
	bThen.AddEdgeTo(bAfter)

	b.AddEdgeTo(bElse)
	s.startBlock(bElse)
	y := s.newValue2(cvttab.subf, ft, x, cutoff)
	y = s.newValue1(cvttab.cvt2U, tt, y)
	z := cvttab.intValue(s, tt, int64(-cvttab.cutoff))
	a1 := s.newValue2(cvttab.or, tt, y, z)
	s.vars[n] = a1
	s.endBlock(pstate)
	bElse.AddEdgeTo(bAfter)

	s.startBlock(bAfter)
	return s.variable(n, n.Type)
}

// dottype generates SSA for a type assertion node.
// commaok indicates whether to panic or return a bool.
// If commaok is false, resok will be nil.
func (s *state) dottype(pstate *PackageState, n *Node, commaok bool) (res, resok *ssa.Value) {
	iface := s.expr(pstate, n.Left)   // input interface
	target := s.expr(pstate, n.Right) // target type
	byteptr := s.f.Config.Types.BytePtr

	if n.Type.IsInterface() {
		if n.Type.IsEmptyInterface(pstate.types) {
			// Converting to an empty interface.
			// Input could be an empty or nonempty interface.
			if pstate.Debug_typeassert > 0 {
				pstate.Warnl(n.Pos, "type assertion inlined")
			}

			// Get itab/type field from input.
			itab := s.newValue1(ssa.OpITab, byteptr, iface)
			// Conversion succeeds iff that field is not nil.
			cond := s.newValue2(ssa.OpNeqPtr, pstate.types.Types[TBOOL], itab, s.constNil(pstate, byteptr))

			if n.Left.Type.IsEmptyInterface(pstate.types) && commaok {
				// Converting empty interface to empty interface with ,ok is just a nil check.
				return iface, cond
			}

			// Branch on nilness.
			b := s.endBlock(pstate)
			b.Kind = ssa.BlockIf
			b.SetControl(cond)
			b.Likely = ssa.BranchLikely
			bOk := s.f.NewBlock(ssa.BlockPlain)
			bFail := s.f.NewBlock(ssa.BlockPlain)
			b.AddEdgeTo(bOk)
			b.AddEdgeTo(bFail)

			if !commaok {
				// On failure, panic by calling panicnildottype.
				s.startBlock(bFail)
				s.rtcall(pstate, pstate.panicnildottype, false, nil, target)

				// On success, return (perhaps modified) input interface.
				s.startBlock(bOk)
				if n.Left.Type.IsEmptyInterface(pstate.types) {
					res = iface // Use input interface unchanged.
					return
				}
				// Load type out of itab, build interface with existing idata.
				off := s.newValue1I(ssa.OpOffPtr, byteptr, int64(pstate.Widthptr), itab)
				typ := s.load(pstate, byteptr, off)
				idata := s.newValue1(ssa.OpIData, n.Type, iface)
				res = s.newValue2(ssa.OpIMake, n.Type, typ, idata)
				return
			}

			s.startBlock(bOk)
			// nonempty -> empty
			// Need to load type from itab
			off := s.newValue1I(ssa.OpOffPtr, byteptr, int64(pstate.Widthptr), itab)
			s.vars[&pstate.typVar] = s.load(pstate, byteptr, off)
			s.endBlock(pstate)

			// itab is nil, might as well use that as the nil result.
			s.startBlock(bFail)
			s.vars[&pstate.typVar] = itab
			s.endBlock(pstate)

			// Merge point.
			bEnd := s.f.NewBlock(ssa.BlockPlain)
			bOk.AddEdgeTo(bEnd)
			bFail.AddEdgeTo(bEnd)
			s.startBlock(bEnd)
			idata := s.newValue1(ssa.OpIData, n.Type, iface)
			res = s.newValue2(ssa.OpIMake, n.Type, s.variable(&pstate.typVar, byteptr), idata)
			resok = cond
			delete(s.vars, &pstate.typVar)
			return
		}
		// converting to a nonempty interface needs a runtime call.
		if pstate.Debug_typeassert > 0 {
			pstate.Warnl(n.Pos, "type assertion not inlined")
		}
		if n.Left.Type.IsEmptyInterface(pstate.types) {
			if commaok {
				call := s.rtcall(pstate, pstate.assertE2I2, true, []*types.Type{n.Type, pstate.types.Types[TBOOL]}, target, iface)
				return call[0], call[1]
			}
			return s.rtcall(pstate, pstate.assertE2I, true, []*types.Type{n.Type}, target, iface)[0], nil
		}
		if commaok {
			call := s.rtcall(pstate, pstate.assertI2I2, true, []*types.Type{n.Type, pstate.types.Types[TBOOL]}, target, iface)
			return call[0], call[1]
		}
		return s.rtcall(pstate, pstate.assertI2I, true, []*types.Type{n.Type}, target, iface)[0], nil
	}

	if pstate.Debug_typeassert > 0 {
		pstate.Warnl(n.Pos, "type assertion inlined")
	}

	// Converting to a concrete type.
	direct := pstate.isdirectiface(n.Type)
	itab := s.newValue1(ssa.OpITab, byteptr, iface) // type word of interface
	if pstate.Debug_typeassert > 0 {
		pstate.Warnl(n.Pos, "type assertion inlined")
	}
	var targetITab *ssa.Value
	if n.Left.Type.IsEmptyInterface(pstate.types) {
		// Looking for pointer to target type.
		targetITab = target
	} else {
		// Looking for pointer to itab for target type and source interface.
		targetITab = s.expr(pstate, n.List.First())
	}

	var tmp *Node       // temporary for use with large types
	var addr *ssa.Value // address of tmp
	if commaok && !pstate.canSSAType(n.Type) {
		// unSSAable type, use temporary.
		// TODO: get rid of some of these temporaries.
		tmp = pstate.tempAt(n.Pos, s.curfn, n.Type)
		addr = s.addr(pstate, tmp, false)
		s.vars[&pstate.memVar] = s.newValue1A(ssa.OpVarDef, pstate.types.TypeMem, tmp, s.mem(pstate))
	}

	cond := s.newValue2(ssa.OpEqPtr, pstate.types.Types[TBOOL], itab, targetITab)
	b := s.endBlock(pstate)
	b.Kind = ssa.BlockIf
	b.SetControl(cond)
	b.Likely = ssa.BranchLikely

	bOk := s.f.NewBlock(ssa.BlockPlain)
	bFail := s.f.NewBlock(ssa.BlockPlain)
	b.AddEdgeTo(bOk)
	b.AddEdgeTo(bFail)

	if !commaok {
		// on failure, panic by calling panicdottype
		s.startBlock(bFail)
		taddr := s.expr(pstate, n.Right.Right)
		if n.Left.Type.IsEmptyInterface(pstate.types) {
			s.rtcall(pstate, pstate.panicdottypeE, false, nil, itab, target, taddr)
		} else {
			s.rtcall(pstate, pstate.panicdottypeI, false, nil, itab, target, taddr)
		}

		// on success, return data from interface
		s.startBlock(bOk)
		if direct {
			return s.newValue1(ssa.OpIData, n.Type, iface), nil
		}
		p := s.newValue1(ssa.OpIData, pstate.types.NewPtr(n.Type), iface)
		return s.load(pstate, n.Type, p), nil
	}

	// commaok is the more complicated case because we have
	// a control flow merge point.
	bEnd := s.f.NewBlock(ssa.BlockPlain)
	// Note that we need a new valVar each time (unlike okVar where we can
	// reuse the variable) because it might have a different type every time.
	valVar := &Node{Op: ONAME, Sym: &types.Sym{Name: "val"}}

	// type assertion succeeded
	s.startBlock(bOk)
	if tmp == nil {
		if direct {
			s.vars[valVar] = s.newValue1(ssa.OpIData, n.Type, iface)
		} else {
			p := s.newValue1(ssa.OpIData, pstate.types.NewPtr(n.Type), iface)
			s.vars[valVar] = s.load(pstate, n.Type, p)
		}
	} else {
		p := s.newValue1(ssa.OpIData, pstate.types.NewPtr(n.Type), iface)
		s.move(pstate, n.Type, addr, p)
	}
	s.vars[&pstate.okVar] = s.constBool(pstate, true)
	s.endBlock(pstate)
	bOk.AddEdgeTo(bEnd)

	// type assertion failed
	s.startBlock(bFail)
	if tmp == nil {
		s.vars[valVar] = s.zeroVal(pstate, n.Type)
	} else {
		s.zero(pstate, n.Type, addr)
	}
	s.vars[&pstate.okVar] = s.constBool(pstate, false)
	s.endBlock(pstate)
	bFail.AddEdgeTo(bEnd)

	// merge point
	s.startBlock(bEnd)
	if tmp == nil {
		res = s.variable(valVar, n.Type)
		delete(s.vars, valVar)
	} else {
		res = s.load(pstate, n.Type, addr)
		s.vars[&pstate.memVar] = s.newValue1A(ssa.OpVarKill, pstate.types.TypeMem, tmp, s.mem(pstate))
	}
	resok = s.variable(&pstate.okVar, pstate.types.Types[TBOOL])
	delete(s.vars, &pstate.okVar)
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
		// No variable should be live at entry.
		s.Fatalf("Value live at entry. It shouldn't be. func %s, node %v, value %v", s.f.Name, name, v)
	}
	// Make a FwdRef, which records a value that's live on block input.
	// We'll find the matching definition as part of insertPhis.
	v = s.newValue0A(ssa.OpFwdRef, t, name)
	s.fwdVars[name] = v
	s.addNamedValue(name, v)
	return v
}

func (s *state) mem(pstate *PackageState) *ssa.Value {
	return s.variable(&pstate.memVar, pstate.types.TypeMem)
}

func (s *state) addNamedValue(n *Node, v *ssa.Value) {
	if n.Class() == Pxxx {
		// Don't track our dummy nodes (&memVar etc.).
		return
	}
	if n.IsAutoTmp() {
		// Don't track temporary variables.
		return
	}
	if n.Class() == PPARAMOUT {
		// Don't track named output values.  This prevents return values
		// from being assigned too early. See #14591 and #14762. TODO: allow this.
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
func (s *SSAGenState) Prog(pstate *PackageState, as obj.As) *obj.Prog {
	p := s.pp.Prog(pstate, as)
	if ssa.LosesStmtMark(as) {
		return p
	}
	// Float a statement start to the beginning of any same-line run.
	// lineRunStart is reset at block boundaries, which appears to work well.
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
func (s *SSAGenState) Br(pstate *PackageState, op obj.As, target *ssa.Block) *obj.Prog {
	p := s.Prog(pstate, op)
	p.To.Type = obj.TYPE_BRANCH
	s.Branches = append(s.Branches, Branch{P: p, B: target})
	return p
}

// DebugFriendlySetPos adjusts Pos.IsStmt subject to heuristics
// that reduce "jumpy" line number churn when debugging.
// Spill/fill/copy instructions from the register allocator,
// phi functions, and instructions with a no-pos position
// are examples of instructions that can cause churn.
func (s *SSAGenState) DebugFriendlySetPosFrom(pstate *PackageState, v *ssa.Value) {
	switch v.Op {
	case ssa.OpPhi, ssa.OpCopy, ssa.OpLoadReg, ssa.OpStoreReg:
		// These are not statements
		s.SetPos(v.Pos.WithNotStmt())
	default:
		p := v.Pos
		if p != pstate.src.NoXPos {
			// If the position is defined, update the position.
			// Also convert default IsStmt to NotStmt; only
			// explicit statement boundaries should appear
			// in the generated code.
			if p.IsStmt() != src.PosIsStmt {
				p = p.WithNotStmt()
			}
			s.SetPos(p)
		}
	}
}

// genssa appends entries to pp for each instruction in f.
func (pstate *PackageState) genssa(f *ssa.Func, pp *Progs) {
	var s SSAGenState

	e := f.Frontend().(*ssafn)

	s.livenessMap = pstate.liveness(e, f)

	// Remember where each block starts.
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

	if pstate.thearch.Use387 {
		s.SSEto387 = map[int16]int16{}
	}

	s.ScratchFpMem = e.scratchFpMem

	if pstate.Ctxt.Flag_locationlists {
		if cap(f.Cache.ValueToProgAfter) < f.NumValues() {
			f.Cache.ValueToProgAfter = make([]*obj.Prog, f.NumValues())
		}
		valueToProgAfter = f.Cache.ValueToProgAfter[:f.NumValues()]
		for i := range valueToProgAfter {
			valueToProgAfter[i] = nil
		}
	}

	// If the very first instruction is not tagged as a statement,
	// debuggers may attribute it to previous function in program.
	firstPos := pstate.src.NoXPos
	for _, v := range f.Entry.Values {
		if v.Pos.IsStmt() == src.PosIsStmt {
			firstPos = v.Pos
			v.Pos = firstPos.WithDefaultStmt()
			break
		}
	}

	// Emit basic blocks
	for i, b := range f.Blocks {
		s.bstart[b.ID] = s.pp.next
		s.pp.nextLive = pstate.LivenessInvalid
		s.lineRunStart = nil

		// Emit values in block
		pstate.thearch.SSAMarkMoves(&s, b)
		for _, v := range b.Values {
			x := s.pp.next
			s.DebugFriendlySetPosFrom(pstate, v)
			// Attach this safe point to the next
			// instruction.
			s.pp.nextLive = s.livenessMap.Get(pstate, v)
			switch v.Op {
			case ssa.OpInitMem:
			// memory arg needs no code
			case ssa.OpArg:
			// input args need no code
			case ssa.OpSP, ssa.OpSB:
			// nothing to do
			case ssa.OpSelect0, ssa.OpSelect1:
			// nothing to do
			case ssa.OpGetG:
			// nothing to do when there's a g register,
			// and checkLower complains if there's not
			case ssa.OpVarDef, ssa.OpVarLive, ssa.OpKeepAlive:
			// nothing to do; already used by liveness
			case ssa.OpVarKill:
				// Zero variable if it is ambiguously live.
				// After the VARKILL anything this variable references
				// might be collected. If it were to become live again later,
				// the GC will see references to already-collected objects.
				// See issue 20029.
				n := v.Aux.(*Node)
				if n.Name.Needzero() {
					if n.Class() != PAUTO {
						v.Fatalf("zero of variable which isn't PAUTO %v", n)
					}
					if n.Type.Size(pstate.types)%int64(pstate.Widthptr) != 0 {
						v.Fatalf("zero of variable not a multiple of ptr size %v", n)
					}
					pstate.thearch.ZeroAuto(s.pp, n)
				}
			case ssa.OpPhi:
				pstate.CheckLoweredPhi(v)
			case ssa.OpConvert:
				// nothing to do; no-op conversion for liveness
				if v.Args[0].Reg(pstate.ssa) != v.Reg(pstate.ssa) {
					v.Fatalf("OpConvert should be a no-op: %s; %s", v.Args[0].LongString(pstate.ssa), v.LongString(pstate.ssa))
				}
			default:
				// let the backend handle it
				// Special case for first line in function; move it to the start.
				if firstPos != pstate.src.NoXPos {
					s.SetPos(firstPos)
					firstPos = pstate.src.NoXPos
				}
				pstate.thearch.SSAGenValue(&s, v)
			}

			if pstate.Ctxt.Flag_locationlists {
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
		if i < len(f.Blocks)-1 && pstate.Debug['N'] == 0 {
			// If -N, leave next==nil so every block with successors
			// ends in a JMP (except call blocks - plive doesn't like
			// select{send,recv} followed by a JMP call).  Helps keep
			// line numbers for otherwise empty blocks.
			next = f.Blocks[i+1]
		}
		x := s.pp.next
		s.SetPos(b.Pos)
		pstate.thearch.SSAGenBlock(&s, b, next)
		if logProgs {
			for ; x != s.pp.next; x = x.Link {
				progToBlock[x] = b
			}
		}
	}

	if pstate.Ctxt.Flag_locationlists {
		e.curfn.Func.DebugInfo = pstate.ssa.BuildFuncDebug(pstate.Ctxt, f, pstate.Debug_locationlist > 1, pstate.stackOffset)
		bstart := s.bstart
		// Note that at this moment, Prog.Pc is a sequence number; it's
		// not a real PC until after assembly, so this mapping has to
		// be done later.
		e.curfn.Func.DebugInfo.GetPC = func(b, v ssa.ID) int64 {
			switch v {
			case pstate.ssa.BlockStart.ID:
				return bstart[b].Pc
			case pstate.ssa.BlockEnd.ID:
				return e.curfn.Func.lsym.Size
			default:
				return valueToProgAfter[v].Pc
			}
		}
	}

	// Resolove branchers, and relax DefaultStmt into NotStmt
	for _, br := range s.Branches {
		br.P.To.Val = s.bstart[br.B.ID]
		if br.P.Pos.IsStmt() != src.PosIsStmt {
			br.P.Pos = br.P.Pos.WithNotStmt()
		}
	}

	if logProgs {
		filename := ""
		for p := pp.Text; p != nil; p = p.Link {
			if p.Pos.IsKnown() && p.InnermostFilename(pstate.obj) != filename {
				filename = p.InnermostFilename(pstate.obj)
				f.Logf("# %s\n", filename)
			}

			var s string
			if v, ok := progToValue[p]; ok {
				s = v.String()
			} else if b, ok := progToBlock[p]; ok {
				s = b.String()
			} else {
				s = "   " // most value and branch strings are 2-3 characters long
			}
			f.Logf(" %-6s\t%.5d (%s)\t%s\n", s, p.Pc, p.InnermostLineNumber(), p.InstructionString(pstate.obj))
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
				// Don't spam every line with the file name, which is often huge.
				// Only print changes, and "unknown" is not a change.
				if p.Pos.IsKnown() && p.InnermostFilename(pstate.obj) != filename {
					filename = p.InnermostFilename(pstate.obj)
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
				buf.WriteString(fmt.Sprintf("%.5d <span class=\"line-number\">(%s)</span> %s", p.Pc, p.InnermostLineNumberHTML(), html.EscapeString(p.InstructionString(pstate.obj))))
				buf.WriteString("</dd>")
			}
			buf.WriteString("</dl>")
			buf.WriteString("</code>")
			f.HTMLWriter.WriteColumn(pstate.ssa, "genssa", "genssa", "ssa-prog", buf.String())
			// pp.Text.Ctxt.LineHist.PrintFilenameOnly = saved
		}
	}

	pstate.defframe(&s, e)
	if pstate.Debug['f'] != 0 {
		pstate.frame(0)
	}

	f.HTMLWriter.Close()
	f.HTMLWriter = nil
}

func (pstate *PackageState) defframe(s *SSAGenState, e *ssafn) {
	pp := s.pp

	frame := pstate.Rnd(s.maxarg+e.stksize, int64(pstate.Widthreg))
	if pstate.thearch.PadFrame != nil {
		frame = pstate.thearch.PadFrame(frame)
	}

	// Fill in argument and frame size.
	pp.Text.To.Type = obj.TYPE_TEXTSIZE
	pp.Text.To.Val = int32(pstate.Rnd(e.curfn.Type.ArgWidth(pstate.types), int64(pstate.Widthreg)))
	pp.Text.To.Offset = frame

	// Insert code to zero ambiguously live variables so that the
	// garbage collector only sees initialized values when it
	// looks for pointers.
	p := pp.Text
	var lo, hi int64

	// Opaque state for backend to use. Current backends use it to
	// keep track of which helper registers have been zeroed.
	var state uint32

	// Iterate through declarations. They are sorted in decreasing Xoffset order.
	for _, n := range e.curfn.Func.Dcl {
		if !n.Name.Needzero() {
			continue
		}
		if n.Class() != PAUTO {
			pstate.Fatalf("needzero class %d", n.Class())
		}
		if n.Type.Size(pstate.types)%int64(pstate.Widthptr) != 0 || n.Xoffset%int64(pstate.Widthptr) != 0 || n.Type.Size(pstate.types) == 0 {
			pstate.Fatalf("var %L has size %d offset %d", n, n.Type.Size(pstate.types), n.Xoffset)
		}

		if lo != hi && n.Xoffset+n.Type.Size(pstate.types) >= lo-int64(2*pstate.Widthreg) {
			// Merge with range we already have.
			lo = n.Xoffset
			continue
		}

		// Zero old range
		p = pstate.thearch.ZeroRange(pp, p, frame+lo, hi-lo, &state)

		// Set new range.
		lo = n.Xoffset
		hi = lo + n.Type.Size(pstate.types)
	}

	// Zero final range.
	pstate.thearch.ZeroRange(pp, p, frame+lo, hi-lo, &state)
}

type FloatingEQNEJump struct {
	Jump  obj.As
	Index int
}

func (s *SSAGenState) oneFPJump(pstate *PackageState, b *ssa.Block, jumps *FloatingEQNEJump) {
	p := s.Prog(pstate, jumps.Jump)
	p.To.Type = obj.TYPE_BRANCH
	p.Pos = b.Pos
	to := jumps.Index
	s.Branches = append(s.Branches, Branch{p, b.Succs[to].Block()})
}

func (s *SSAGenState) FPJump(pstate *PackageState, b, next *ssa.Block, jumps *[2][2]FloatingEQNEJump) {
	switch next {
	case b.Succs[0].Block():
		s.oneFPJump(pstate, b, &jumps[0][0])
		s.oneFPJump(pstate, b, &jumps[0][1])
	case b.Succs[1].Block():
		s.oneFPJump(pstate, b, &jumps[1][0])
		s.oneFPJump(pstate, b, &jumps[1][1])
	default:
		s.oneFPJump(pstate, b, &jumps[1][0])
		s.oneFPJump(pstate, b, &jumps[1][1])
		q := s.Prog(pstate, obj.AJMP)
		q.Pos = b.Pos
		q.To.Type = obj.TYPE_BRANCH
		s.Branches = append(s.Branches, Branch{q, b.Succs[1].Block()})
	}
}

func (pstate *PackageState) AuxOffset(v *ssa.Value) (offset int64) {
	if v.Aux == nil {
		return 0
	}
	n, ok := v.Aux.(*Node)
	if !ok {
		v.Fatalf("bad aux type in %s\n", v.LongString(pstate.ssa))
	}
	if n.Class() == PAUTO {
		return n.Xoffset
	}
	return 0
}

// AddAux adds the offset in the aux fields (AuxInt and Aux) of v to a.
func (pstate *PackageState) AddAux(a *obj.Addr, v *ssa.Value) {
	pstate.AddAux2(a, v, v.AuxInt)
}
func (pstate *PackageState) AddAux2(a *obj.Addr, v *ssa.Value, offset int64) {
	if a.Type != obj.TYPE_MEM && a.Type != obj.TYPE_ADDR {
		v.Fatalf("bad AddAux addr %v", a)
	}
	// add integer offset
	a.Offset += offset

	// If no additional symbol offset, we're done.
	if v.Aux == nil {
		return
	}
	// Add symbol's offset from its base register.
	switch n := v.Aux.(type) {
	case *obj.LSym:
		a.Name = obj.NAME_EXTERN
		a.Sym = n
	case *Node:
		if n.Class() == PPARAM || n.Class() == PPARAMOUT {
			a.Name = obj.NAME_PARAM
			a.Sym = n.Orig.Sym.Linksym(pstate.types)
			a.Offset += n.Xoffset
			break
		}
		a.Name = obj.NAME_AUTO
		a.Sym = n.Sym.Linksym(pstate.types)
		a.Offset += n.Xoffset
	default:
		v.Fatalf("aux in %s not implemented %#v", v, v.Aux)
	}
}

// extendIndex extends v to a full int width.
// panic using the given function if v does not fit in an int (only on 32-bit archs).
func (s *state) extendIndex(pstate *PackageState, v *ssa.Value, panicfn *obj.LSym) *ssa.Value {
	size := v.Type.Size(pstate.types)
	if size == s.config.PtrSize {
		return v
	}
	if size > s.config.PtrSize {
		// truncate 64-bit indexes on 32-bit pointer archs. Test the
		// high word and branch to out-of-bounds failure if it is not 0.
		if pstate.Debug['B'] == 0 {
			hi := s.newValue1(ssa.OpInt64Hi, pstate.types.Types[TUINT32], v)
			cmp := s.newValue2(ssa.OpEq32, pstate.types.Types[TBOOL], hi, s.constInt32(pstate, pstate.types.Types[TUINT32], 0))
			s.check(pstate, cmp, panicfn)
		}
		return s.newValue1(ssa.OpTrunc64to32, pstate.types.Types[TINT], v)
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
	return s.newValue1(op, pstate.types.Types[TINT], v)
}

// CheckLoweredPhi checks that regalloc and stackalloc correctly handled phi values.
// Called during ssaGenValue.
func (pstate *PackageState) CheckLoweredPhi(v *ssa.Value) {
	if v.Op != ssa.OpPhi {
		v.Fatalf("CheckLoweredPhi called with non-phi value: %v", v.LongString(pstate.ssa))
	}
	if v.Type.IsMemory(pstate.types) {
		return
	}
	f := v.Block.Func
	loc := f.RegAlloc[v.ID]
	for _, a := range v.Args {
		if aloc := f.RegAlloc[a.ID]; aloc != loc { // TODO: .Equal() instead?
			v.Fatalf("phi arg at different location than phi: %v @ %s, but arg %v @ %s\n%s\n", v, loc, a, aloc, v.Block.Func)
		}
	}
}

// CheckLoweredGetClosurePtr checks that v is the first instruction in the function's entry block.
// The output of LoweredGetClosurePtr is generally hardwired to the correct register.
// That register contains the closure pointer on closure entry.
func (pstate *PackageState) CheckLoweredGetClosurePtr(v *ssa.Value) {
	entry := v.Block.Func.Entry
	if entry != v.Block || entry.Values[0] != v {
		pstate.Fatalf("in %s, badly placed LoweredGetClosurePtr: %v %v", v.Block.Func.Name, v.Block, v)
	}
}

// AutoVar returns a *Node and int64 representing the auto variable and offset within it
// where v should be spilled.
func (pstate *PackageState) AutoVar(v *ssa.Value) (*Node, int64) {
	loc := v.Block.Func.RegAlloc[v.ID].(ssa.LocalSlot)
	if v.Type.Size(pstate.types) > loc.Type.Size(pstate.types) {
		v.Fatalf("spill/restore type %s doesn't fit in slot type %s", v.Type, loc.Type)
	}
	return loc.N.(*Node), loc.Off
}

func (pstate *PackageState) AddrAuto(a *obj.Addr, v *ssa.Value) {
	n, off := pstate.AutoVar(v)
	a.Type = obj.TYPE_MEM
	a.Sym = n.Sym.Linksym(pstate.types)
	a.Reg = int16(pstate.thearch.REGSP)
	a.Offset = n.Xoffset + off
	if n.Class() == PPARAM || n.Class() == PPARAMOUT {
		a.Name = obj.NAME_PARAM
	} else {
		a.Name = obj.NAME_AUTO
	}
}

func (s *SSAGenState) AddrScratch(pstate *PackageState, a *obj.Addr) {
	if s.ScratchFpMem == nil {
		panic("no scratch memory available; forgot to declare usesScratch for Op?")
	}
	a.Type = obj.TYPE_MEM
	a.Name = obj.NAME_AUTO
	a.Sym = s.ScratchFpMem.Sym.Linksym(pstate.types)
	a.Reg = int16(pstate.thearch.REGSP)
	a.Offset = s.ScratchFpMem.Xoffset
}

// Call returns a new CALL instruction for the SSA value v.
// It uses PrepareCall to prepare the call.
func (s *SSAGenState) Call(pstate *PackageState, v *ssa.Value) *obj.Prog {
	s.PrepareCall(pstate, v)

	p := s.Prog(pstate, obj.ACALL)
	if sym, ok := v.Aux.(*obj.LSym); ok {
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		p.To.Sym = sym
	} else {
		// TODO(mdempsky): Can these differences be eliminated?
		switch pstate.thearch.LinkArch.Family {
		case sys.AMD64, sys.I386, sys.PPC64, sys.S390X, sys.Wasm:
			p.To.Type = obj.TYPE_REG
		case sys.ARM, sys.ARM64, sys.MIPS, sys.MIPS64:
			p.To.Type = obj.TYPE_MEM
		default:
			pstate.Fatalf("unknown indirect call family")
		}
		p.To.Reg = v.Args[0].Reg(pstate.ssa)
	}
	return p
}

// PrepareCall prepares to emit a CALL instruction for v and does call-related bookkeeping.
// It must be called immediately before emitting the actual CALL instruction,
// since it emits PCDATA for the stack map at the call (calls are safe points).
func (s *SSAGenState) PrepareCall(pstate *PackageState, v *ssa.Value) {
	idx := s.livenessMap.Get(pstate, v)
	if !idx.Valid() {
		// typedmemclr and typedmemmove are write barriers and
		// deeply non-preemptible. They are unsafe points and
		// hence should not have liveness maps.
		if sym, _ := v.Aux.(*obj.LSym); !(sym == pstate.typedmemclr || sym == pstate.typedmemmove) {
			pstate.Fatalf("missing stack map index for %v", v.LongString(pstate.ssa))
		}
	}

	if sym, _ := v.Aux.(*obj.LSym); sym == pstate.Deferreturn {
		// Deferred calls will appear to be returning to
		// the CALL deferreturn(SB) that we are about to emit.
		// However, the stack trace code will show the line
		// of the instruction byte before the return PC.
		// To avoid that being an unrelated instruction,
		// insert an actual hardware NOP that will have the right line number.
		// This is different from obj.ANOP, which is a virtual no-op
		// that doesn't make it into the instruction stream.
		pstate.thearch.Ginsnop(s.pp)
	}

	if sym, ok := v.Aux.(*obj.LSym); ok {
		// Record call graph information for nowritebarrierrec
		// analysis.
		if pstate.nowritebarrierrecCheck != nil {
			pstate.nowritebarrierrecCheck.recordCall(pstate, s.pp.curfn, sym, v.Pos)
		}
	}

	if s.maxarg < v.AuxInt {
		s.maxarg = v.AuxInt
	}
}

// fieldIdx finds the index of the field referred to by the ODOT node n.
func (pstate *PackageState) fieldIdx(n *Node) int {
	t := n.Left.Type
	f := n.Sym
	if !t.IsStruct() {
		panic("ODOT's LHS is not a struct")
	}

	var i int
	for _, t1 := range t.Fields(pstate.types).Slice() {
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

	// TODO: keep the result of this function somewhere in the ODOT Node
	// so we don't have to recompute it each time we need it.
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
func (e *ssafn) StringData(pstate *PackageState, s string) interface{} {
	if aux, ok := e.strings[s]; ok {
		return aux
	}
	if e.strings == nil {
		e.strings = make(map[string]interface{})
	}
	data := pstate.stringsym(e.curfn.Pos, s)
	e.strings[s] = data
	return data
}

func (e *ssafn) Auto(pstate *PackageState, pos src.XPos, t *types.Type) ssa.GCNode {
	n := pstate.tempAt(pos, e.curfn, t) // Note: adds new auto to e.curfn.Func.Dcl list
	return n
}

func (e *ssafn) SplitString(pstate *PackageState, name ssa.LocalSlot) (ssa.LocalSlot, ssa.LocalSlot) {
	n := name.N.(*Node)
	ptrType := pstate.types.NewPtr(pstate.types.Types[TUINT8])
	lenType := pstate.types.Types[TINT]
	if n.Class() == PAUTO && !n.Addrtaken() {
		// Split this string up into two separate variables.
		p := e.splitSlot(pstate, &name, ".ptr", 0, ptrType)
		l := e.splitSlot(pstate, &name, ".len", ptrType.Size(pstate.types), lenType)
		return p, l
	}
	// Return the two parts of the larger variable.
	return ssa.LocalSlot{N: n, Type: ptrType, Off: name.Off}, ssa.LocalSlot{N: n, Type: lenType, Off: name.Off + int64(pstate.Widthptr)}
}

func (e *ssafn) SplitInterface(pstate *PackageState, name ssa.LocalSlot) (ssa.LocalSlot, ssa.LocalSlot) {
	n := name.N.(*Node)
	u := pstate.types.Types[TUINTPTR]
	t := pstate.types.NewPtr(pstate.types.Types[TUINT8])
	if n.Class() == PAUTO && !n.Addrtaken() {
		// Split this interface up into two separate variables.
		f := ".itab"
		if n.Type.IsEmptyInterface(pstate.types) {
			f = ".type"
		}
		c := e.splitSlot(pstate, &name, f, 0, u) // see comment in plive.go:onebitwalktype1.
		d := e.splitSlot(pstate, &name, ".data", u.Size(pstate.types), t)
		return c, d
	}
	// Return the two parts of the larger variable.
	return ssa.LocalSlot{N: n, Type: u, Off: name.Off}, ssa.LocalSlot{N: n, Type: t, Off: name.Off + int64(pstate.Widthptr)}
}

func (e *ssafn) SplitSlice(pstate *PackageState, name ssa.LocalSlot) (ssa.LocalSlot, ssa.LocalSlot, ssa.LocalSlot) {
	n := name.N.(*Node)
	ptrType := pstate.types.NewPtr(name.Type.Elem(pstate.types))
	lenType := pstate.types.Types[TINT]
	if n.Class() == PAUTO && !n.Addrtaken() {
		// Split this slice up into three separate variables.
		p := e.splitSlot(pstate, &name, ".ptr", 0, ptrType)
		l := e.splitSlot(pstate, &name, ".len", ptrType.Size(pstate.types), lenType)
		c := e.splitSlot(pstate, &name, ".cap", ptrType.Size(pstate.types)+lenType.Size(pstate.types), lenType)
		return p, l, c
	}
	// Return the three parts of the larger variable.
	return ssa.LocalSlot{N: n, Type: ptrType, Off: name.Off},
		ssa.LocalSlot{N: n, Type: lenType, Off: name.Off + int64(pstate.Widthptr)},
		ssa.LocalSlot{N: n, Type: lenType, Off: name.Off + int64(2*pstate.Widthptr)}
}

func (e *ssafn) SplitComplex(pstate *PackageState, name ssa.LocalSlot) (ssa.LocalSlot, ssa.LocalSlot) {
	n := name.N.(*Node)
	s := name.Type.Size(pstate.types) / 2
	var t *types.Type
	if s == 8 {
		t = pstate.types.Types[TFLOAT64]
	} else {
		t = pstate.types.Types[TFLOAT32]
	}
	if n.Class() == PAUTO && !n.Addrtaken() {
		// Split this complex up into two separate variables.
		r := e.splitSlot(pstate, &name, ".real", 0, t)
		i := e.splitSlot(pstate, &name, ".imag", t.Size(pstate.types), t)
		return r, i
	}
	// Return the two parts of the larger variable.
	return ssa.LocalSlot{N: n, Type: t, Off: name.Off}, ssa.LocalSlot{N: n, Type: t, Off: name.Off + s}
}

func (e *ssafn) SplitInt64(pstate *PackageState, name ssa.LocalSlot) (ssa.LocalSlot, ssa.LocalSlot) {
	n := name.N.(*Node)
	var t *types.Type
	if name.Type.IsSigned() {
		t = pstate.types.Types[TINT32]
	} else {
		t = pstate.types.Types[TUINT32]
	}
	if n.Class() == PAUTO && !n.Addrtaken() {
		// Split this int64 up into two separate variables.
		if pstate.thearch.LinkArch.ByteOrder == binary.BigEndian {
			return e.splitSlot(pstate, &name, ".hi", 0, t), e.splitSlot(pstate, &name, ".lo", t.Size(pstate.types), pstate.types.Types[TUINT32])
		}
		return e.splitSlot(pstate, &name, ".hi", t.Size(pstate.types), t), e.splitSlot(pstate, &name, ".lo", 0, pstate.types.Types[TUINT32])
	}
	// Return the two parts of the larger variable.
	if pstate.thearch.LinkArch.ByteOrder == binary.BigEndian {
		return ssa.LocalSlot{N: n, Type: t, Off: name.Off}, ssa.LocalSlot{N: n, Type: pstate.types.Types[TUINT32], Off: name.Off + 4}
	}
	return ssa.LocalSlot{N: n, Type: t, Off: name.Off + 4}, ssa.LocalSlot{N: n, Type: pstate.types.Types[TUINT32], Off: name.Off}
}

func (e *ssafn) SplitStruct(pstate *PackageState, name ssa.LocalSlot, i int) ssa.LocalSlot {
	n := name.N.(*Node)
	st := name.Type
	ft := st.FieldType(pstate.types, i)
	var offset int64
	for f := 0; f < i; f++ {
		offset += st.FieldType(pstate.types, f).Size(pstate.types)
	}
	if n.Class() == PAUTO && !n.Addrtaken() {
		// Note: the _ field may appear several times.  But
		// have no fear, identically-named but distinct Autos are
		// ok, albeit maybe confusing for a debugger.
		return e.splitSlot(pstate, &name, "."+st.FieldName(pstate.types, i), offset, ft)
	}
	return ssa.LocalSlot{N: n, Type: ft, Off: name.Off + st.FieldOff(pstate.types, i)}
}

func (e *ssafn) SplitArray(pstate *PackageState, name ssa.LocalSlot) ssa.LocalSlot {
	n := name.N.(*Node)
	at := name.Type
	if at.NumElem(pstate.types) != 1 {
		pstate.Fatalf("bad array size")
	}
	et := at.Elem(pstate.types)
	if n.Class() == PAUTO && !n.Addrtaken() {
		return e.splitSlot(pstate, &name, "[0]", 0, et)
	}
	return ssa.LocalSlot{N: n, Type: et, Off: name.Off}
}

func (e *ssafn) DerefItab(pstate *PackageState, it *obj.LSym, offset int64) *obj.LSym {
	return pstate.itabsym(it, offset)
}

// splitSlot returns a slot representing the data of parent starting at offset.
func (e *ssafn) splitSlot(pstate *PackageState, parent *ssa.LocalSlot, suffix string, offset int64, t *types.Type) ssa.LocalSlot {
	s := &types.Sym{Name: parent.N.(*Node).Sym.Name + suffix, Pkg: pstate.localpkg}

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
	pstate.dowidth(t)
	return ssa.LocalSlot{N: n, Type: t, Off: 0, SplitOf: parent, SplitOffset: offset}
}

func (e *ssafn) CanSSA(pstate *PackageState, t *types.Type) bool {
	return pstate.canSSAType(t)
}

func (e *ssafn) Line(pstate *PackageState, pos src.XPos) string {
	return pstate.linestr(pos)
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
func (e *ssafn) Fatalf(pstate *PackageState, pos src.XPos, msg string, args ...interface{}) {
	pstate.lineno = pos
	pstate.Fatalf(msg, args...)
}

// Warnl reports a "warning", which is usually flag-triggered
// logging output for the benefit of tests.
func (e *ssafn) Warnl(pstate *PackageState, pos src.XPos, fmt_ string, args ...interface{}) {
	pstate.Warnl(pos, fmt_, args...)
}

func (e *ssafn) Debug_checknil(pstate *PackageState) bool {
	return pstate.Debug_checknil != 0
}

func (e *ssafn) UseWriteBarrier(pstate *PackageState) bool {
	return pstate.use_writebarrier
}

func (e *ssafn) Syslook(pstate *PackageState, name string) *obj.LSym {
	switch name {
	case "goschedguarded":
		return pstate.goschedguarded
	case "writeBarrier":
		return pstate.writeBarrier
	case "gcWriteBarrier":
		return pstate.gcWriteBarrier
	case "typedmemmove":
		return pstate.typedmemmove
	case "typedmemclr":
		return pstate.typedmemclr
	}
	pstate.Fatalf("unknown Syslook func %v", name)
	return nil
}

func (e *ssafn) SetWBPos(pstate *PackageState, pos src.XPos) {
	e.curfn.Func.setWBPos(pstate, pos)
}

func (n *Node) Typ() *types.Type {
	return n.Type
}
func (n *Node) StorageClass(pstate *PackageState) ssa.StorageClass {
	switch n.Class() {
	case PPARAM:
		return ssa.ClassParam
	case PPARAMOUT:
		return ssa.ClassParamOut
	case PAUTO:
		return ssa.ClassAuto
	default:
		pstate.Fatalf("untranslateable storage class for %v: %s", n, n.Class())
		return 0
	}
}

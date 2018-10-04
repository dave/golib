// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/dwarf"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"github.com/dave/golib/src/cmd/internal/sys"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
)

func (pstate *PackageState) emitptrargsmap(fn *Node) {
	if fn.funcname() == "_" {
		return
	}
	sym := pstate.lookup(fmt.Sprintf("%s.args_stackmap", fn.funcname()))
	lsym := sym.Linksym(pstate.types)

	nptr := int(fn.Type.ArgWidth(pstate.types) / int64(pstate.Widthptr))
	bv := bvalloc(int32(nptr) * 2)
	nbitmap := 1
	if fn.Type.NumResults(pstate.types) > 0 {
		nbitmap = 2
	}
	off := pstate.duint32(lsym, 0, uint32(nbitmap))
	off = pstate.duint32(lsym, off, uint32(bv.n))

	if fn.IsMethod(pstate) {
		pstate.onebitwalktype1(fn.Type.Recvs(pstate.types), 0, bv)
	}
	if fn.Type.NumParams(pstate.types) > 0 {
		pstate.onebitwalktype1(fn.Type.Params(pstate.types), 0, bv)
	}
	off = pstate.dbvec(lsym, off, bv)

	if fn.Type.NumResults(pstate.types) > 0 {
		pstate.onebitwalktype1(fn.Type.Results(pstate.types), 0, bv)
		off = pstate.dbvec(lsym, off, bv)
	}

	pstate.ggloblsym(lsym, int32(off), obj.RODATA|obj.LOCAL)
}

// cmpstackvarlt reports whether the stack variable a sorts before b.
//
// Sort the list of stack variables. Autos after anything else,
// within autos, unused after used, within used, things with
// pointers first, zeroed things first, and then decreasing size.
// Because autos are laid out in decreasing addresses
// on the stack, pointers first, zeroed things first and decreasing size
// really means, in memory, things with pointers needing zeroing at
// the top of the stack and increasing in size.
// Non-autos sort on offset.
func (pstate *PackageState) cmpstackvarlt(a, b *Node) bool {
	if (a.Class() == PAUTO) != (b.Class() == PAUTO) {
		return b.Class() == PAUTO
	}

	if a.Class() != PAUTO {
		return a.Xoffset < b.Xoffset
	}

	if a.Name.Used() != b.Name.Used() {
		return a.Name.Used()
	}

	ap := pstate.types.Haspointers(a.Type)
	bp := pstate.types.Haspointers(b.Type)
	if ap != bp {
		return ap
	}

	ap = a.Name.Needzero()
	bp = b.Name.Needzero()
	if ap != bp {
		return ap
	}

	if a.Type.Width != b.Type.Width {
		return a.Type.Width > b.Type.Width
	}

	return a.Sym.Name < b.Sym.Name
}

// byStackvar implements sort.Interface for []*Node using cmpstackvarlt.
type byStackVar []*Node

func (s byStackVar) Len() int                                 { return len(s) }
func (s byStackVar) Less(pstate *PackageState, i, j int) bool { return pstate.cmpstackvarlt(s[i], s[j]) }
func (s byStackVar) Swap(i, j int)                            { s[i], s[j] = s[j], s[i] }

func (s *ssafn) AllocFrame(pstate *PackageState, f *ssa.Func) {
	s.stksize = 0
	s.stkptrsize = 0
	fn := s.curfn.Func

	// Mark the PAUTO's unused.
	for _, ln := range fn.Dcl {
		if ln.Class() == PAUTO {
			ln.Name.SetUsed(false)
		}
	}

	for _, l := range f.RegAlloc {
		if ls, ok := l.(ssa.LocalSlot); ok {
			ls.N.(*Node).Name.SetUsed(true)
		}
	}

	scratchUsed := false
	for _, b := range f.Blocks {
		for _, v := range b.Values {
			if n, ok := v.Aux.(*Node); ok {
				switch n.Class() {
				case PPARAM, PPARAMOUT:
					// Don't modify nodfp; it is a global.
					if n != pstate.nodfp {
						n.Name.SetUsed(true)
					}
				case PAUTO:
					n.Name.SetUsed(true)
				}
			}
			if !scratchUsed {
				scratchUsed = v.Op.UsesScratch(pstate.ssa)
			}

		}
	}

	if f.Config.NeedsFpScratch && scratchUsed {
		s.scratchFpMem = pstate.tempAt(pstate.src.NoXPos, s.curfn, pstate.types.Types[TUINT64])
	}

	sort.Sort(byStackVar(fn.Dcl))

	// Reassign stack offsets of the locals that are used.
	for i, n := range fn.Dcl {
		if n.Op != ONAME || n.Class() != PAUTO {
			continue
		}
		if !n.Name.Used() {
			fn.Dcl = fn.Dcl[:i]
			break
		}

		pstate.dowidth(n.Type)
		w := n.Type.Width
		if w >= pstate.thearch.MAXWIDTH || w < 0 {
			pstate.Fatalf("bad width")
		}
		s.stksize += w
		s.stksize = pstate.Rnd(s.stksize, int64(n.Type.Align))
		if pstate.types.Haspointers(n.Type) {
			s.stkptrsize = s.stksize
		}
		if pstate.thearch.LinkArch.InFamily(sys.MIPS, sys.MIPS64, sys.ARM, sys.ARM64, sys.PPC64, sys.S390X) {
			s.stksize = pstate.Rnd(s.stksize, int64(pstate.Widthptr))
		}
		n.Xoffset = -s.stksize
	}

	s.stksize = pstate.Rnd(s.stksize, int64(pstate.Widthreg))
	s.stkptrsize = pstate.Rnd(s.stkptrsize, int64(pstate.Widthreg))
}

func (pstate *PackageState) funccompile(fn *Node) {
	if pstate.Curfn != nil {
		pstate.Fatalf("funccompile %v inside %v", fn.Func.Nname.Sym, pstate.Curfn.Func.Nname.Sym)
	}

	if fn.Type == nil {
		if pstate.nerrors == 0 {
			pstate.Fatalf("funccompile missing type")
		}
		return
	}

	// assign parameter offsets
	pstate.dowidth(fn.Type)

	if fn.Nbody.Len() == 0 {
		pstate.emitptrargsmap(fn)
		return
	}

	pstate.dclcontext = PAUTO
	pstate.Curfn = fn

	pstate.compile(fn)

	pstate.Curfn = nil
	pstate.dclcontext = PEXTERN
}

func (pstate *PackageState) compile(fn *Node) {
	pstate.saveerrors()

	pstate.order(fn)
	if pstate.nerrors != 0 {
		return
	}

	pstate.walk(fn)
	if pstate.nerrors != 0 {
		return
	}
	if pstate.instrumenting {
		pstate.instrument(fn)
	}

	// From this point, there should be no uses of Curfn. Enforce that.
	pstate.Curfn = nil

	// Set up the function's LSym early to avoid data races with the assemblers.
	fn.Func.initLSym(pstate)

	if pstate.compilenow() {
		pstate.compileSSA(fn, 0)
	} else {
		pstate.compilequeue = append(pstate.compilequeue, fn)
	}
}

// compilenow reports whether to compile immediately.
// If functions are not compiled immediately,
// they are enqueued in compilequeue,
// which is drained by compileFunctions.
func (pstate *PackageState) compilenow() bool {
	return pstate.nBackendWorkers == 1 && pstate.Debug_compilelater == 0
}

const maxStackSize = 1 << 30

// compileSSA builds an SSA backend function,
// uses it to generate a plist,
// and flushes that plist to machine code.
// worker indicates which of the backend workers is doing the processing.
func (pstate *PackageState) compileSSA(fn *Node, worker int) {
	f := pstate.buildssa(fn, worker)
	// Note: check arg size to fix issue 25507.
	if f.Frontend().(*ssafn).stksize >= maxStackSize || fn.Type.ArgWidth(pstate.types) >= maxStackSize {
		pstate.largeStackFramesMu.Lock()
		pstate.largeStackFrames = append(pstate.largeStackFrames, fn.Pos)
		pstate.largeStackFramesMu.Unlock()
		return
	}
	pp := pstate.newProgs(fn, worker)
	defer pp.Free(pstate)
	pstate.genssa(f, pp)
	// Check frame size again.
	// The check above included only the space needed for local variables.
	// After genssa, the space needed includes local variables and the callee arg region.
	// We must do this check prior to calling pp.Flush.
	// If there are any oversized stack frames,
	// the assembler may emit inscrutable complaints about invalid instructions.
	if pp.Text.To.Offset >= maxStackSize {
		pstate.largeStackFramesMu.Lock()
		pstate.largeStackFrames = append(pstate.largeStackFrames, fn.Pos)
		pstate.largeStackFramesMu.Unlock()
		return
	}

	pp.Flush(pstate) // assemble, fill in boilerplate, etc.
	// fieldtrack must be called after pp.Flush. See issue 20014.
	pstate.fieldtrack(pp.Text.From.Sym, fn.Func.FieldTrack)
}

func init() {
	if raceEnabled {
		rand.Seed(time.Now().UnixNano())
	}
}

// compileFunctions compiles all functions in compilequeue.
// It fans out nBackendWorkers to do the work
// and waits for them to complete.
func (pstate *PackageState) compileFunctions() {
	if len(pstate.compilequeue) != 0 {
		pstate.sizeCalculationDisabled = true // not safe to calculate sizes concurrently
		if raceEnabled {
			// Randomize compilation order to try to shake out races.
			tmp := make([]*Node, len(pstate.compilequeue))
			perm := rand.Perm(len(pstate.compilequeue))
			for i, v := range perm {
				tmp[v] = pstate.compilequeue[i]
			}
			copy(pstate.compilequeue, tmp)
		} else {
			// Compile the longest functions first,
			// since they're most likely to be the slowest.
			// This helps avoid stragglers.
			obj.SortSlice(pstate.compilequeue, func(i, j int) bool {
				return pstate.compilequeue[i].Nbody.Len() > pstate.compilequeue[j].Nbody.Len()
			})
		}
		var wg sync.WaitGroup
		pstate.Ctxt.InParallel = true
		c := make(chan *Node, pstate.nBackendWorkers)
		for i := 0; i < pstate.nBackendWorkers; i++ {
			wg.Add(1)
			go func(worker int) {
				for fn := range c {
					pstate.compileSSA(fn, worker)
				}
				wg.Done()
			}(i)
		}
		for _, fn := range pstate.compilequeue {
			c <- fn
		}
		close(c)
		pstate.compilequeue = nil
		wg.Wait()
		pstate.Ctxt.InParallel = false
		pstate.sizeCalculationDisabled = false
	}
}

func (pstate *PackageState) debuginfo(fnsym *obj.LSym, curfn interface{}) ([]dwarf.Scope, dwarf.InlCalls) {
	fn := curfn.(*Node)
	if fn.Func.Nname != nil {
		if expect := fn.Func.Nname.Sym.Linksym(pstate.types); fnsym != expect {
			pstate.Fatalf("unexpected fnsym: %v != %v", fnsym, expect)
		}
	}

	var automDecls []*Node
	// Populate Automs for fn.
	for _, n := range fn.Func.Dcl {
		if n.Op != ONAME { // might be OTYPE or OLITERAL
			continue
		}
		var name obj.AddrName
		switch n.Class() {
		case PAUTO:
			if !n.Name.Used() {
				// Text == nil -> generating abstract function
				if fnsym.Func.Text != nil {
					pstate.Fatalf("debuginfo unused node (AllocFrame should truncate fn.Func.Dcl)")
				}
				continue
			}
			name = obj.NAME_AUTO
		case PPARAM, PPARAMOUT:
			name = obj.NAME_PARAM
		default:
			continue
		}
		automDecls = append(automDecls, n)
		gotype := pstate.ngotype(n).Linksym(pstate.types)
		fnsym.Func.Autom = append(fnsym.Func.Autom, &obj.Auto{
			Asym:    pstate.Ctxt.Lookup(n.Sym.Name),
			Aoffset: int32(n.Xoffset),
			Name:    name,
			Gotype:  gotype,
		})
	}

	decls, dwarfVars := pstate.createDwarfVars(fnsym, fn.Func, automDecls)

	var varScopes []ScopeID
	for _, decl := range decls {
		pos := decl.Pos
		if decl.Name.Defn != nil && (decl.Name.Captured() || decl.Name.Byval()) {
			// It's not clear which position is correct for captured variables here:
			// * decl.Pos is the wrong position for captured variables, in the inner
			//   function, but it is the right position in the outer function.
			// * decl.Name.Defn is nil for captured variables that were arguments
			//   on the outer function, however the decl.Pos for those seems to be
			//   correct.
			// * decl.Name.Defn is the "wrong" thing for variables declared in the
			//   header of a type switch, it's their position in the header, rather
			//   than the position of the case statement. In principle this is the
			//   right thing, but here we prefer the latter because it makes each
			//   instance of the header variable local to the lexical block of its
			//   case statement.
			// This code is probably wrong for type switch variables that are also
			// captured.
			pos = decl.Name.Defn.Pos
		}
		varScopes = append(varScopes, pstate.findScope(fn.Func.Marks, pos))
	}

	scopes := pstate.assembleScopes(fnsym, fn, dwarfVars, varScopes)
	var inlcalls dwarf.InlCalls
	if pstate.genDwarfInline > 0 {
		inlcalls = pstate.assembleInlines(fnsym, dwarfVars)
	}
	return scopes, inlcalls
}

// createSimpleVars creates a DWARF entry for every variable declared in the
// function, claiming that they are permanently on the stack.
func (pstate *PackageState) createSimpleVars(automDecls []*Node) ([]*Node, []*dwarf.Var, map[*Node]bool) {
	var vars []*dwarf.Var
	var decls []*Node
	selected := make(map[*Node]bool)
	for _, n := range automDecls {
		if n.IsAutoTmp() {
			continue
		}
		var abbrev int
		offs := n.Xoffset

		switch n.Class() {
		case PAUTO:
			abbrev = dwarf.DW_ABRV_AUTO
			if pstate.Ctxt.FixedFrameSize() == 0 {
				offs -= int64(pstate.Widthptr)
			}
			if pstate.objabi.Framepointer_enabled(pstate.objabi.GOOS, pstate.objabi.GOARCH) {
				offs -= int64(pstate.Widthptr)
			}

		case PPARAM, PPARAMOUT:
			abbrev = dwarf.DW_ABRV_PARAM
			offs += pstate.Ctxt.FixedFrameSize()
		default:
			pstate.Fatalf("createSimpleVars unexpected type %v for node %v", n.Class(), n)
		}

		selected[n] = true
		typename := dwarf.InfoPrefix + pstate.typesymname(n.Type)
		decls = append(decls, n)
		inlIndex := 0
		if pstate.genDwarfInline > 1 {
			if n.InlFormal() || n.InlLocal() {
				inlIndex = pstate.posInlIndex(n.Pos) + 1
				if n.InlFormal() {
					abbrev = dwarf.DW_ABRV_PARAM
				}
			}
		}
		declpos := pstate.Ctxt.InnermostPos(n.Pos)
		vars = append(vars, &dwarf.Var{
			Name:          n.Sym.Name,
			IsReturnValue: n.Class() == PPARAMOUT,
			IsInlFormal:   n.InlFormal(),
			Abbrev:        abbrev,
			StackOffset:   int32(offs),
			Type:          pstate.Ctxt.Lookup(typename),
			DeclFile:      declpos.RelFilename(),
			DeclLine:      declpos.RelLine(pstate.src),
			DeclCol:       declpos.Col(),
			InlIndex:      int32(inlIndex),
			ChildIndex:    -1,
		})
	}
	return decls, vars, selected
}

// createComplexVars creates recomposed DWARF vars with location lists,
// suitable for describing optimized code.
func (pstate *PackageState) createComplexVars(fn *Func) ([]*Node, []*dwarf.Var, map[*Node]bool) {
	debugInfo := fn.DebugInfo

	// Produce a DWARF variable entry for each user variable.
	var decls []*Node
	var vars []*dwarf.Var
	ssaVars := make(map[*Node]bool)

	for varID, dvar := range debugInfo.Vars {
		n := dvar.(*Node)
		ssaVars[n] = true
		for _, slot := range debugInfo.VarSlots[varID] {
			ssaVars[debugInfo.Slots[slot].N.(*Node)] = true
		}

		if dvar := pstate.createComplexVar(fn, ssa.VarID(varID)); dvar != nil {
			decls = append(decls, n)
			vars = append(vars, dvar)
		}
	}

	return decls, vars, ssaVars
}

// createDwarfVars process fn, returning a list of DWARF variables and the
// Nodes they represent.
func (pstate *PackageState) createDwarfVars(fnsym *obj.LSym, fn *Func, automDecls []*Node) ([]*Node, []*dwarf.Var) {
	// Collect a raw list of DWARF vars.
	var vars []*dwarf.Var
	var decls []*Node
	var selected map[*Node]bool
	if pstate.Ctxt.Flag_locationlists && pstate.Ctxt.Flag_optimize && fn.DebugInfo != nil {
		decls, vars, selected = pstate.createComplexVars(fn)
	} else {
		decls, vars, selected = pstate.createSimpleVars(automDecls)
	}

	var dcl []*Node
	if fnsym.WasInlined() {
		dcl = pstate.preInliningDcls(fnsym)
	} else {
		dcl = automDecls
	}

	// If optimization is enabled, the list above will typically be
	// missing some of the original pre-optimization variables in the
	// function (they may have been promoted to registers, folded into
	// constants, dead-coded away, etc). Here we add back in entries
	// for selected missing vars. Note that the recipe below creates a
	// conservative location. The idea here is that we want to
	// communicate to the user that "yes, there is a variable named X
	// in this function, but no, I don't have enough information to
	// reliably report its contents."
	for _, n := range dcl {
		if _, found := selected[n]; found {
			continue
		}
		c := n.Sym.Name[0]
		if c == '.' || n.Type.IsUntyped(pstate.types) {
			continue
		}
		typename := dwarf.InfoPrefix + pstate.typesymname(n.Type)
		decls = append(decls, n)
		abbrev := dwarf.DW_ABRV_AUTO_LOCLIST
		if n.Class() == PPARAM || n.Class() == PPARAMOUT {
			abbrev = dwarf.DW_ABRV_PARAM_LOCLIST
		}
		inlIndex := 0
		if pstate.genDwarfInline > 1 {
			if n.InlFormal() || n.InlLocal() {
				inlIndex = pstate.posInlIndex(n.Pos) + 1
				if n.InlFormal() {
					abbrev = dwarf.DW_ABRV_PARAM_LOCLIST
				}
			}
		}
		declpos := pstate.Ctxt.InnermostPos(n.Pos)
		vars = append(vars, &dwarf.Var{
			Name:          n.Sym.Name,
			IsReturnValue: n.Class() == PPARAMOUT,
			Abbrev:        abbrev,
			StackOffset:   int32(n.Xoffset),
			Type:          pstate.Ctxt.Lookup(typename),
			DeclFile:      declpos.RelFilename(),
			DeclLine:      declpos.RelLine(pstate.src),
			DeclCol:       declpos.Col(),
			InlIndex:      int32(inlIndex),
			ChildIndex:    -1,
		})
		// Append a "deleted auto" entry to the autom list so as to
		// insure that the type in question is picked up by the linker.
		// See issue 22941.
		gotype := pstate.ngotype(n).Linksym(pstate.types)
		fnsym.Func.Autom = append(fnsym.Func.Autom, &obj.Auto{
			Asym:    pstate.Ctxt.Lookup(n.Sym.Name),
			Aoffset: int32(-1),
			Name:    obj.NAME_DELETED_AUTO,
			Gotype:  gotype,
		})

	}

	return decls, vars
}

// Given a function that was inlined at some point during the
// compilation, return a sorted list of nodes corresponding to the
// autos/locals in that function prior to inlining. If this is a
// function that is not local to the package being compiled, then the
// names of the variables may have been "versioned" to avoid conflicts
// with local vars; disregard this versioning when sorting.
func (pstate *PackageState) preInliningDcls(fnsym *obj.LSym) []*Node {
	fn := pstate.Ctxt.DwFixups.GetPrecursorFunc(fnsym).(*Node)
	var rdcl []*Node
	for _, n := range fn.Func.Inl.Dcl {
		c := n.Sym.Name[0]
		// Avoid reporting "_" parameters, since if there are more than
		// one, it can result in a collision later on, as in #23179.
		if unversion(n.Sym.Name) == "_" || c == '.' || n.Type.IsUntyped(pstate.types) {
			continue
		}
		rdcl = append(rdcl, n)
	}
	sort.Sort(byNodeName(rdcl))
	return rdcl
}

func cmpNodeName(a, b *Node) bool {
	aart := 0
	if strings.HasPrefix(a.Sym.Name, "~") {
		aart = 1
	}
	bart := 0
	if strings.HasPrefix(b.Sym.Name, "~") {
		bart = 1
	}
	if aart != bart {
		return aart < bart
	}

	aname := unversion(a.Sym.Name)
	bname := unversion(b.Sym.Name)
	return aname < bname
}

// byNodeName implements sort.Interface for []*Node using cmpNodeName.
type byNodeName []*Node

func (s byNodeName) Len() int           { return len(s) }
func (s byNodeName) Less(i, j int) bool { return cmpNodeName(s[i], s[j]) }
func (s byNodeName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// stackOffset returns the stack location of a LocalSlot relative to the
// stack pointer, suitable for use in a DWARF location entry. This has nothing
// to do with its offset in the user variable.
func (pstate *PackageState) stackOffset(slot ssa.LocalSlot) int32 {
	n := slot.N.(*Node)
	var base int64
	switch n.Class() {
	case PAUTO:
		if pstate.Ctxt.FixedFrameSize() == 0 {
			base -= int64(pstate.Widthptr)
		}
		if pstate.objabi.Framepointer_enabled(pstate.objabi.GOOS, pstate.objabi.GOARCH) {
			base -= int64(pstate.Widthptr)
		}
	case PPARAM, PPARAMOUT:
		base += pstate.Ctxt.FixedFrameSize()
	}
	return int32(base + n.Xoffset + slot.Off)
}

// createComplexVar builds a single DWARF variable entry and location list.
func (pstate *PackageState) createComplexVar(fn *Func, varID ssa.VarID) *dwarf.Var {
	debug := fn.DebugInfo
	n := debug.Vars[varID].(*Node)

	var abbrev int
	switch n.Class() {
	case PAUTO:
		abbrev = dwarf.DW_ABRV_AUTO_LOCLIST
	case PPARAM, PPARAMOUT:
		abbrev = dwarf.DW_ABRV_PARAM_LOCLIST
	default:
		return nil
	}

	gotype := pstate.ngotype(n).Linksym(pstate.types)
	typename := dwarf.InfoPrefix + gotype.Name[len("type."):]
	inlIndex := 0
	if pstate.genDwarfInline > 1 {
		if n.InlFormal() || n.InlLocal() {
			inlIndex = pstate.posInlIndex(n.Pos) + 1
			if n.InlFormal() {
				abbrev = dwarf.DW_ABRV_PARAM_LOCLIST
			}
		}
	}
	declpos := pstate.Ctxt.InnermostPos(n.Pos)
	dvar := &dwarf.Var{
		Name:          n.Sym.Name,
		IsReturnValue: n.Class() == PPARAMOUT,
		IsInlFormal:   n.InlFormal(),
		Abbrev:        abbrev,
		Type:          pstate.Ctxt.Lookup(typename),
		// The stack offset is used as a sorting key, so for decomposed
		// variables just give it the first one. It's not used otherwise.
		// This won't work well if the first slot hasn't been assigned a stack
		// location, but it's not obvious how to do better.
		StackOffset: pstate.stackOffset(debug.Slots[debug.VarSlots[varID][0]]),
		DeclFile:    declpos.RelFilename(),
		DeclLine:    declpos.RelLine(pstate.src),
		DeclCol:     declpos.Col(),
		InlIndex:    int32(inlIndex),
		ChildIndex:  -1,
	}
	list := debug.LocationLists[varID]
	if len(list) != 0 {
		dvar.PutLocationList = func(listSym, startPC dwarf.Sym) {
			debug.PutLocationList(list, pstate.Ctxt, listSym.(*obj.LSym), startPC.(*obj.LSym))
		}
	}
	return dvar
}

// fieldtrack adds R_USEFIELD relocations to fnsym to record any
// struct fields that it used.
func (pstate *PackageState) fieldtrack(fnsym *obj.LSym, tracked map[*types.Sym]struct{}) {
	if fnsym == nil {
		return
	}
	if pstate.objabi.Fieldtrack_enabled == 0 || len(tracked) == 0 {
		return
	}

	trackSyms := make([]*types.Sym, 0, len(tracked))
	for sym := range tracked {
		trackSyms = append(trackSyms, sym)
	}
	sort.Sort(symByName(trackSyms))
	for _, sym := range trackSyms {
		r := obj.Addrel(fnsym)
		r.Sym = sym.Linksym(pstate.types)
		r.Type = objabi.R_USEFIELD
	}
}

type symByName []*types.Sym

func (a symByName) Len() int           { return len(a) }
func (a symByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a symByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

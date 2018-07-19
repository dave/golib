package gc

import (
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/dwarf"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"

	"github.com/dave/golib/src/cmd/internal/sys"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
)

// number of concurrent backend workers, set by a compiler flag
// functions waiting to be compiled

func (psess *PackageSession) emitptrargsmap(fn *Node) {
	if fn.funcname() == "_" {
		return
	}
	sym := psess.lookup(fmt.Sprintf("%s.args_stackmap", fn.funcname()))
	lsym := sym.Linksym(psess.types)

	nptr := int(fn.Type.ArgWidth(psess.types) / int64(psess.Widthptr))
	bv := bvalloc(int32(nptr) * 2)
	nbitmap := 1
	if fn.Type.NumResults(psess.types) > 0 {
		nbitmap = 2
	}
	off := psess.duint32(lsym, 0, uint32(nbitmap))
	off = psess.duint32(lsym, off, uint32(bv.n))

	if fn.IsMethod(psess) {
		psess.
			onebitwalktype1(fn.Type.Recvs(psess.types), 0, bv)
	}
	if fn.Type.NumParams(psess.types) > 0 {
		psess.
			onebitwalktype1(fn.Type.Params(psess.types), 0, bv)
	}
	off = psess.dbvec(lsym, off, bv)

	if fn.Type.NumResults(psess.types) > 0 {
		psess.
			onebitwalktype1(fn.Type.Results(psess.types), 0, bv)
		off = psess.dbvec(lsym, off, bv)
	}
	psess.
		ggloblsym(lsym, int32(off), obj.RODATA|obj.LOCAL)
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
func (psess *PackageSession) cmpstackvarlt(a, b *Node) bool {
	if (a.Class() == PAUTO) != (b.Class() == PAUTO) {
		return b.Class() == PAUTO
	}

	if a.Class() != PAUTO {
		return a.Xoffset < b.Xoffset
	}

	if a.Name.Used() != b.Name.Used() {
		return a.Name.Used()
	}

	ap := psess.types.Haspointers(a.Type)
	bp := psess.types.Haspointers(b.Type)
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

func (s byStackVar) Len() int                                  { return len(s) }
func (s byStackVar) Less(psess *PackageSession, i, j int) bool { return psess.cmpstackvarlt(s[i], s[j]) }
func (s byStackVar) Swap(i, j int)                             { s[i], s[j] = s[j], s[i] }

func (s *ssafn) AllocFrame(psess *PackageSession, f *ssa.Func) {
	s.stksize = 0
	s.stkptrsize = 0
	fn := s.curfn.Func

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

					if n != psess.nodfp {
						n.Name.SetUsed(true)
					}
				case PAUTO:
					n.Name.SetUsed(true)
				}
			}
			if !scratchUsed {
				scratchUsed = v.Op.UsesScratch(psess.ssa)
			}

		}
	}

	if f.Config.NeedsFpScratch && scratchUsed {
		s.scratchFpMem = psess.tempAt(psess.src.NoXPos, s.curfn, psess.types.Types[TUINT64])
	}

	sort.Sort(byStackVar(fn.Dcl))

	for i, n := range fn.Dcl {
		if n.Op != ONAME || n.Class() != PAUTO {
			continue
		}
		if !n.Name.Used() {
			fn.Dcl = fn.Dcl[:i]
			break
		}
		psess.
			dowidth(n.Type)
		w := n.Type.Width
		if w >= psess.thearch.MAXWIDTH || w < 0 {
			psess.
				Fatalf("bad width")
		}
		s.stksize += w
		s.stksize = psess.Rnd(s.stksize, int64(n.Type.Align))
		if psess.types.Haspointers(n.Type) {
			s.stkptrsize = s.stksize
		}
		if psess.thearch.LinkArch.InFamily(sys.MIPS, sys.MIPS64, sys.ARM, sys.ARM64, sys.PPC64, sys.S390X) {
			s.stksize = psess.Rnd(s.stksize, int64(psess.Widthptr))
		}
		n.Xoffset = -s.stksize
	}

	s.stksize = psess.Rnd(s.stksize, int64(psess.Widthreg))
	s.stkptrsize = psess.Rnd(s.stkptrsize, int64(psess.Widthreg))
}

func (psess *PackageSession) funccompile(fn *Node) {
	if psess.Curfn != nil {
		psess.
			Fatalf("funccompile %v inside %v", fn.Func.Nname.Sym, psess.Curfn.Func.Nname.Sym)
	}

	if fn.Type == nil {
		if psess.nerrors == 0 {
			psess.
				Fatalf("funccompile missing type")
		}
		return
	}
	psess.
		dowidth(fn.Type)

	if fn.Nbody.Len() == 0 {
		psess.
			emitptrargsmap(fn)
		return
	}
	psess.
		dclcontext = PAUTO
	psess.
		Curfn = fn
	psess.
		compile(fn)
	psess.
		Curfn = nil
	psess.
		dclcontext = PEXTERN
}

func (psess *PackageSession) compile(fn *Node) {
	psess.
		saveerrors()
	psess.
		order(fn)
	if psess.nerrors != 0 {
		return
	}
	psess.
		walk(fn)
	if psess.nerrors != 0 {
		return
	}
	if psess.instrumenting {
		psess.
			instrument(fn)
	}
	psess.
		Curfn = nil

	fn.Func.initLSym(psess)

	if psess.compilenow() {
		psess.
			compileSSA(fn, 0)
	} else {
		psess.
			compilequeue = append(psess.compilequeue, fn)
	}
}

// compilenow reports whether to compile immediately.
// If functions are not compiled immediately,
// they are enqueued in compilequeue,
// which is drained by compileFunctions.
func (psess *PackageSession) compilenow() bool {
	return psess.nBackendWorkers == 1 && psess.Debug_compilelater == 0
}

const maxStackSize = 1 << 30

// compileSSA builds an SSA backend function,
// uses it to generate a plist,
// and flushes that plist to machine code.
// worker indicates which of the backend workers is doing the processing.
func (psess *PackageSession) compileSSA(fn *Node, worker int) {
	f := psess.buildssa(fn, worker)

	if f.Frontend().(*ssafn).stksize >= maxStackSize || fn.Type.ArgWidth(psess.types) >= maxStackSize {
		psess.
			largeStackFramesMu.Lock()
		psess.
			largeStackFrames = append(psess.largeStackFrames, fn.Pos)
		psess.
			largeStackFramesMu.Unlock()
		return
	}
	pp := psess.newProgs(fn, worker)
	defer pp.Free(psess)
	psess.
		genssa(f, pp)

	if pp.Text.To.Offset >= maxStackSize {
		psess.
			largeStackFramesMu.Lock()
		psess.
			largeStackFrames = append(psess.largeStackFrames, fn.Pos)
		psess.
			largeStackFramesMu.Unlock()
		return
	}

	pp.Flush(psess)
	psess.
		fieldtrack(pp.Text.From.Sym, fn.Func.FieldTrack)
}

func init() {
	if raceEnabled {
		rand.Seed(time.Now().UnixNano())
	}
}

// compileFunctions compiles all functions in compilequeue.
// It fans out nBackendWorkers to do the work
// and waits for them to complete.
func (psess *PackageSession) compileFunctions() {
	if len(psess.compilequeue) != 0 {
		psess.
			sizeCalculationDisabled = true
		if raceEnabled {

			tmp := make([]*Node, len(psess.compilequeue))
			perm := rand.Perm(len(psess.compilequeue))
			for i, v := range perm {
				tmp[v] = psess.compilequeue[i]
			}
			copy(psess.compilequeue, tmp)
		} else {

			obj.SortSlice(psess.compilequeue, func(i, j int) bool {
				return psess.compilequeue[i].Nbody.Len() > psess.compilequeue[j].Nbody.Len()
			})
		}
		var wg sync.WaitGroup
		psess.
			Ctxt.InParallel = true
		c := make(chan *Node, psess.nBackendWorkers)
		for i := 0; i < psess.nBackendWorkers; i++ {
			wg.Add(1)
			go func(worker int) {
				for fn := range c {
					psess.
						compileSSA(fn, worker)
				}
				wg.Done()
			}(i)
		}
		for _, fn := range psess.compilequeue {
			c <- fn
		}
		close(c)
		psess.
			compilequeue = nil
		wg.Wait()
		psess.
			Ctxt.InParallel = false
		psess.
			sizeCalculationDisabled = false
	}
}

func (psess *PackageSession) debuginfo(fnsym *obj.LSym, curfn interface{}) ([]dwarf.Scope, dwarf.InlCalls) {
	fn := curfn.(*Node)
	if fn.Func.Nname != nil {
		if expect := fn.Func.Nname.Sym.Linksym(psess.types); fnsym != expect {
			psess.
				Fatalf("unexpected fnsym: %v != %v", fnsym, expect)
		}
	}

	var automDecls []*Node

	for _, n := range fn.Func.Dcl {
		if n.Op != ONAME {
			continue
		}
		var name obj.AddrName
		switch n.Class() {
		case PAUTO:
			if !n.Name.Used() {

				if fnsym.Func.Text != nil {
					psess.
						Fatalf("debuginfo unused node (AllocFrame should truncate fn.Func.Dcl)")
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
		gotype := psess.ngotype(n).Linksym(psess.types)
		fnsym.Func.Autom = append(fnsym.Func.Autom, &obj.Auto{
			Asym:    psess.Ctxt.Lookup(n.Sym.Name),
			Aoffset: int32(n.Xoffset),
			Name:    name,
			Gotype:  gotype,
		})
	}

	decls, dwarfVars := psess.createDwarfVars(fnsym, fn.Func, automDecls)

	var varScopes []ScopeID
	for _, decl := range decls {
		pos := decl.Pos
		if decl.Name.Defn != nil && (decl.Name.Captured() || decl.Name.Byval()) {

			pos = decl.Name.Defn.Pos
		}
		varScopes = append(varScopes, psess.findScope(fn.Func.Marks, pos))
	}

	scopes := psess.assembleScopes(fnsym, fn, dwarfVars, varScopes)
	var inlcalls dwarf.InlCalls
	if psess.genDwarfInline > 0 {
		inlcalls = psess.assembleInlines(fnsym, dwarfVars)
	}
	return scopes, inlcalls
}

// createSimpleVars creates a DWARF entry for every variable declared in the
// function, claiming that they are permanently on the stack.
func (psess *PackageSession) createSimpleVars(automDecls []*Node) ([]*Node, []*dwarf.Var, map[*Node]bool) {
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
			if psess.Ctxt.FixedFrameSize() == 0 {
				offs -= int64(psess.Widthptr)
			}
			if psess.objabi.Framepointer_enabled(psess.objabi.GOOS, psess.objabi.GOARCH) {
				offs -= int64(psess.Widthptr)
			}

		case PPARAM, PPARAMOUT:
			abbrev = dwarf.DW_ABRV_PARAM
			offs += psess.Ctxt.FixedFrameSize()
		default:
			psess.
				Fatalf("createSimpleVars unexpected type %v for node %v", n.Class(), n)
		}

		selected[n] = true
		typename := dwarf.InfoPrefix + psess.typesymname(n.Type)
		decls = append(decls, n)
		inlIndex := 0
		if psess.genDwarfInline > 1 {
			if n.InlFormal() || n.InlLocal() {
				inlIndex = psess.posInlIndex(n.Pos) + 1
				if n.InlFormal() {
					abbrev = dwarf.DW_ABRV_PARAM
				}
			}
		}
		declpos := psess.Ctxt.InnermostPos(n.Pos)
		vars = append(vars, &dwarf.Var{
			Name:          n.Sym.Name,
			IsReturnValue: n.Class() == PPARAMOUT,
			IsInlFormal:   n.InlFormal(),
			Abbrev:        abbrev,
			StackOffset:   int32(offs),
			Type:          psess.Ctxt.Lookup(typename),
			DeclFile:      declpos.RelFilename(),
			DeclLine:      declpos.RelLine(psess.src),
			DeclCol:       declpos.Col(),
			InlIndex:      int32(inlIndex),
			ChildIndex:    -1,
		})
	}
	return decls, vars, selected
}

// createComplexVars creates recomposed DWARF vars with location lists,
// suitable for describing optimized code.
func (psess *PackageSession) createComplexVars(fn *Func) ([]*Node, []*dwarf.Var, map[*Node]bool) {
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

		if dvar := psess.createComplexVar(fn, ssa.VarID(varID)); dvar != nil {
			decls = append(decls, n)
			vars = append(vars, dvar)
		}
	}

	return decls, vars, ssaVars
}

// createDwarfVars process fn, returning a list of DWARF variables and the
// Nodes they represent.
func (psess *PackageSession) createDwarfVars(fnsym *obj.LSym, fn *Func, automDecls []*Node) ([]*Node, []*dwarf.Var) {
	// Collect a raw list of DWARF vars.
	var vars []*dwarf.Var
	var decls []*Node
	var selected map[*Node]bool
	if psess.Ctxt.Flag_locationlists && psess.Ctxt.Flag_optimize && fn.DebugInfo != nil {
		decls, vars, selected = psess.createComplexVars(fn)
	} else {
		decls, vars, selected = psess.createSimpleVars(automDecls)
	}

	var dcl []*Node
	if fnsym.WasInlined() {
		dcl = psess.preInliningDcls(fnsym)
	} else {
		dcl = automDecls
	}

	for _, n := range dcl {
		if _, found := selected[n]; found {
			continue
		}
		c := n.Sym.Name[0]
		if c == '.' || n.Type.IsUntyped(psess.types) {
			continue
		}
		typename := dwarf.InfoPrefix + psess.typesymname(n.Type)
		decls = append(decls, n)
		abbrev := dwarf.DW_ABRV_AUTO_LOCLIST
		if n.Class() == PPARAM || n.Class() == PPARAMOUT {
			abbrev = dwarf.DW_ABRV_PARAM_LOCLIST
		}
		inlIndex := 0
		if psess.genDwarfInline > 1 {
			if n.InlFormal() || n.InlLocal() {
				inlIndex = psess.posInlIndex(n.Pos) + 1
				if n.InlFormal() {
					abbrev = dwarf.DW_ABRV_PARAM_LOCLIST
				}
			}
		}
		declpos := psess.Ctxt.InnermostPos(n.Pos)
		vars = append(vars, &dwarf.Var{
			Name:          n.Sym.Name,
			IsReturnValue: n.Class() == PPARAMOUT,
			Abbrev:        abbrev,
			StackOffset:   int32(n.Xoffset),
			Type:          psess.Ctxt.Lookup(typename),
			DeclFile:      declpos.RelFilename(),
			DeclLine:      declpos.RelLine(psess.src),
			DeclCol:       declpos.Col(),
			InlIndex:      int32(inlIndex),
			ChildIndex:    -1,
		})

		gotype := psess.ngotype(n).Linksym(psess.types)
		fnsym.Func.Autom = append(fnsym.Func.Autom, &obj.Auto{
			Asym:    psess.Ctxt.Lookup(n.Sym.Name),
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
func (psess *PackageSession) preInliningDcls(fnsym *obj.LSym) []*Node {
	fn := psess.Ctxt.DwFixups.GetPrecursorFunc(fnsym).(*Node)
	var rdcl []*Node
	for _, n := range fn.Func.Inl.Dcl {
		c := n.Sym.Name[0]

		if unversion(n.Sym.Name) == "_" || c == '.' || n.Type.IsUntyped(psess.types) {
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
func (psess *PackageSession) stackOffset(slot ssa.LocalSlot) int32 {
	n := slot.N.(*Node)
	var base int64
	switch n.Class() {
	case PAUTO:
		if psess.Ctxt.FixedFrameSize() == 0 {
			base -= int64(psess.Widthptr)
		}
		if psess.objabi.Framepointer_enabled(psess.objabi.GOOS, psess.objabi.GOARCH) {
			base -= int64(psess.Widthptr)
		}
	case PPARAM, PPARAMOUT:
		base += psess.Ctxt.FixedFrameSize()
	}
	return int32(base + n.Xoffset + slot.Off)
}

// createComplexVar builds a single DWARF variable entry and location list.
func (psess *PackageSession) createComplexVar(fn *Func, varID ssa.VarID) *dwarf.Var {
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

	gotype := psess.ngotype(n).Linksym(psess.types)
	typename := dwarf.InfoPrefix + gotype.Name[len("type."):]
	inlIndex := 0
	if psess.genDwarfInline > 1 {
		if n.InlFormal() || n.InlLocal() {
			inlIndex = psess.posInlIndex(n.Pos) + 1
			if n.InlFormal() {
				abbrev = dwarf.DW_ABRV_PARAM_LOCLIST
			}
		}
	}
	declpos := psess.Ctxt.InnermostPos(n.Pos)
	dvar := &dwarf.Var{
		Name:          n.Sym.Name,
		IsReturnValue: n.Class() == PPARAMOUT,
		IsInlFormal:   n.InlFormal(),
		Abbrev:        abbrev,
		Type:          psess.Ctxt.Lookup(typename),

		StackOffset: psess.stackOffset(debug.Slots[debug.VarSlots[varID][0]]),
		DeclFile:    declpos.RelFilename(),
		DeclLine:    declpos.RelLine(psess.src),
		DeclCol:     declpos.Col(),
		InlIndex:    int32(inlIndex),
		ChildIndex:  -1,
	}
	list := debug.LocationLists[varID]
	if len(list) != 0 {
		dvar.PutLocationList = func(listSym, startPC dwarf.Sym) {
			debug.PutLocationList(list, psess.Ctxt, listSym.(*obj.LSym), startPC.(*obj.LSym))
		}
	}
	return dvar
}

// fieldtrack adds R_USEFIELD relocations to fnsym to record any
// struct fields that it used.
func (psess *PackageSession) fieldtrack(fnsym *obj.LSym, tracked map[*types.Sym]struct{}) {
	if fnsym == nil {
		return
	}
	if psess.objabi.Fieldtrack_enabled == 0 || len(tracked) == 0 {
		return
	}

	trackSyms := make([]*types.Sym, 0, len(tracked))
	for sym := range tracked {
		trackSyms = append(trackSyms, sym)
	}
	sort.Sort(symByName(trackSyms))
	for _, sym := range trackSyms {
		r := obj.Addrel(fnsym)
		r.Sym = sym.Linksym(psess.types)
		r.Type = objabi.R_USEFIELD
	}
}

type symByName []*types.Sym

func (a symByName) Len() int           { return len(a) }
func (a symByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a symByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

package gc

import (
	"github.com/dave/golib/src/cmd/internal/dwarf"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/src"
	"sort"
	"strings"
)

// To identify variables by original source position.
type varPos struct {
	DeclName string
	DeclFile string
	DeclLine uint
	DeclCol  uint
}

// This is the main entry point for collection of raw material to
// drive generation of DWARF "inlined subroutine" DIEs. See proposal
// 22080 for more details and background info.
func (psess *PackageSession) assembleInlines(fnsym *obj.LSym, dwVars []*dwarf.Var) dwarf.InlCalls {
	var inlcalls dwarf.InlCalls

	if psess.Debug_gendwarfinl != 0 {
		psess.
			Ctxt.Logf("assembling DWARF inlined routine info for %v\n", fnsym.Name)
	}

	imap := make(map[int]int)

	// Walk progs to build up the InlCalls data structure
	var prevpos src.XPos
	for p := fnsym.Func.Text; p != nil; p = p.Link {
		if p.Pos == prevpos {
			continue
		}
		ii := psess.posInlIndex(p.Pos)
		if ii >= 0 {
			psess.
				insertInlCall(&inlcalls, ii, imap)
		}
		prevpos = p.Pos
	}

	vmap := make(map[int32][]*dwarf.Var)

	for _, dwv := range dwVars {

		vmap[dwv.InlIndex] = append(vmap[dwv.InlIndex], dwv)

		if dwv.InlIndex == 0 {
			continue
		}

		ii := int(dwv.InlIndex) - 1
		idx, ok := imap[ii]
		if !ok {

			idx = psess.insertInlCall(&inlcalls, ii, imap)
		}
		inlcalls.Calls[idx].InlVars =
			append(inlcalls.Calls[idx].InlVars, dwv)
	}

	for ii, sl := range vmap {
		sort.Sort(byClassThenName(sl))
		var m map[varPos]int
		if ii == 0 {
			if !fnsym.WasInlined() {
				for j, v := range sl {
					v.ChildIndex = int32(j)
				}
				continue
			}
			m = psess.makePreinlineDclMap(fnsym)
		} else {
			ifnlsym := psess.Ctxt.InlTree.InlinedFunction(int(ii - 1))
			m = psess.makePreinlineDclMap(ifnlsym)
		}

		synthCount := len(m)
		for _, v := range sl {
			canonName := unversion(v.Name)
			vp := varPos{
				DeclName: canonName,
				DeclFile: v.DeclFile,
				DeclLine: v.DeclLine,
				DeclCol:  v.DeclCol,
			}
			synthesized := strings.HasPrefix(v.Name, "~r") || canonName == "_"
			if idx, found := m[vp]; found {
				v.ChildIndex = int32(idx)
				v.IsInAbstract = !synthesized
				v.Name = canonName
			} else {

				v.ChildIndex = int32(synthCount)
				synthCount += 1
			}
		}
	}

	curii := -1
	var crange *dwarf.Range
	var prevp *obj.Prog
	for p := fnsym.Func.Text; p != nil; prevp, p = p, p.Link {
		if prevp != nil && p.Pos == prevp.Pos {
			continue
		}
		ii := psess.posInlIndex(p.Pos)
		if ii == curii {
			continue
		} else {

			endRange(crange, p)

			crange = psess.beginRange(inlcalls.Calls, p, ii, imap)
			curii = ii
		}
	}
	if crange != nil {
		crange.End = fnsym.Size
	}

	if psess.Debug_gendwarfinl != 0 {
		psess.
			dumpInlCalls(inlcalls)
		psess.
			dumpInlVars(dwVars)
	}

	return inlcalls
}

// Secondary hook for DWARF inlined subroutine generation. This is called
// late in the compilation when it is determined that we need an
// abstract function DIE for an inlined routine imported from a
// previously compiled package.
func (psess *PackageSession) genAbstractFunc(fn *obj.LSym) {
	ifn := psess.Ctxt.DwFixups.GetPrecursorFunc(fn)
	if ifn == nil {
		psess.
			Ctxt.Diag("failed to locate precursor fn for %v", fn)
		return
	}
	if psess.Debug_gendwarfinl != 0 {
		psess.
			Ctxt.Logf("DwarfAbstractFunc(%v)\n", fn.Name)
	}
	psess.
		Ctxt.DwarfAbstractFunc(psess.obj, ifn, fn, psess.myimportpath)
}

// Undo any versioning performed when a name was written
// out as part of export data.
func unversion(name string) string {
	if i := strings.Index(name, "·"); i > 0 {
		name = name[:i]
	}
	return name
}

// Given a function that was inlined as part of the compilation, dig
// up the pre-inlining DCL list for the function and create a map that
// supports lookup of pre-inline dcl index, based on variable
// position/name. NB: the recipe for computing variable pos/file/line
// needs to be kept in sync with the similar code in gc.createSimpleVars
// and related functions.
func (psess *PackageSession) makePreinlineDclMap(fnsym *obj.LSym) map[varPos]int {
	dcl := psess.preInliningDcls(fnsym)
	m := make(map[varPos]int)
	for i, n := range dcl {
		pos := psess.Ctxt.InnermostPos(n.Pos)
		vp := varPos{
			DeclName: unversion(n.Sym.Name),
			DeclFile: pos.RelFilename(),
			DeclLine: pos.RelLine(psess.src),
			DeclCol:  pos.Col(),
		}
		if _, found := m[vp]; found {
			psess.
				Fatalf("child dcl collision on symbol %s within %v\n", n.Sym.Name, fnsym.Name)
		}
		m[vp] = i
	}
	return m
}

func (psess *PackageSession) insertInlCall(dwcalls *dwarf.InlCalls, inlIdx int, imap map[int]int) int {
	callIdx, found := imap[inlIdx]
	if found {
		return callIdx
	}

	parCallIdx := -1
	parInlIdx := psess.Ctxt.InlTree.Parent(inlIdx)
	if parInlIdx >= 0 {
		parCallIdx = psess.insertInlCall(dwcalls, parInlIdx, imap)
	}

	inlinedFn := psess.Ctxt.InlTree.InlinedFunction(inlIdx)
	callXPos := psess.Ctxt.InlTree.CallPos(inlIdx)
	absFnSym := psess.Ctxt.DwFixups.AbsFuncDwarfSym(inlinedFn)
	pb := psess.Ctxt.PosTable.Pos(callXPos).Base()
	callFileSym := psess.Ctxt.Lookup(pb.SymFilename())
	ic := dwarf.InlCall{
		InlIndex:  inlIdx,
		CallFile:  callFileSym,
		CallLine:  uint32(callXPos.Line()),
		AbsFunSym: absFnSym,
		Root:      parCallIdx == -1,
	}
	dwcalls.Calls = append(dwcalls.Calls, ic)
	callIdx = len(dwcalls.Calls) - 1
	imap[inlIdx] = callIdx

	if parCallIdx != -1 {

		dwcalls.Calls[parCallIdx].Children = append(dwcalls.Calls[parCallIdx].Children, callIdx)
	}

	return callIdx
}

// Given a src.XPos, return its associated inlining index if it
// corresponds to something created as a result of an inline, or -1 if
// there is no inline info. Note that the index returned will refer to
// the deepest call in the inlined stack, e.g. if you have "A calls B
// calls C calls D" and all three callees are inlined (B, C, and D),
// the index for a node from the inlined body of D will refer to the
// call to D from C. Whew.
func (psess *PackageSession) posInlIndex(xpos src.XPos) int {
	pos := psess.Ctxt.PosTable.Pos(xpos)
	if b := pos.Base(); b != nil {
		ii := b.InliningIndex()
		if ii >= 0 {
			return ii
		}
	}
	return -1
}

func endRange(crange *dwarf.Range, p *obj.Prog) {
	if crange == nil {
		return
	}
	crange.End = p.Pc
}

func (psess *PackageSession) beginRange(calls []dwarf.InlCall, p *obj.Prog, ii int, imap map[int]int) *dwarf.Range {
	if ii == -1 {
		return nil
	}
	callIdx, found := imap[ii]
	if !found {
		psess.
			Fatalf("internal error: can't find inlIndex %d in imap for prog at %d\n", ii, p.Pc)
	}
	call := &calls[callIdx]

	call.Ranges = append(call.Ranges, dwarf.Range{Start: p.Pc, End: -1})
	return &call.Ranges[len(call.Ranges)-1]
}

func cmpDwarfVar(a, b *dwarf.Var) bool {

	aart := 0
	if strings.HasPrefix(a.Name, "~r") {
		aart = 1
	}
	bart := 0
	if strings.HasPrefix(b.Name, "~r") {
		bart = 1
	}
	if aart != bart {
		return aart < bart
	}

	return a.Name < b.Name
}

// byClassThenName implements sort.Interface for []*dwarf.Var using cmpDwarfVar.
type byClassThenName []*dwarf.Var

func (s byClassThenName) Len() int           { return len(s) }
func (s byClassThenName) Less(i, j int) bool { return cmpDwarfVar(s[i], s[j]) }
func (s byClassThenName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (psess *PackageSession) dumpInlCall(inlcalls dwarf.InlCalls, idx, ilevel int) {
	for i := 0; i < ilevel; i++ {
		psess.
			Ctxt.Logf("  ")
	}
	ic := inlcalls.Calls[idx]
	callee := psess.Ctxt.InlTree.InlinedFunction(ic.InlIndex)
	psess.
		Ctxt.Logf("  %d: II:%d (%s) V: (", idx, ic.InlIndex, callee.Name)
	for _, f := range ic.InlVars {
		psess.
			Ctxt.Logf(" %v", f.Name)
	}
	psess.
		Ctxt.Logf(" ) C: (")
	for _, k := range ic.Children {
		psess.
			Ctxt.Logf(" %v", k)
	}
	psess.
		Ctxt.Logf(" ) R:")
	for _, r := range ic.Ranges {
		psess.
			Ctxt.Logf(" [%d,%d)", r.Start, r.End)
	}
	psess.
		Ctxt.Logf("\n")
	for _, k := range ic.Children {
		psess.
			dumpInlCall(inlcalls, k, ilevel+1)
	}

}

func (psess *PackageSession) dumpInlCalls(inlcalls dwarf.InlCalls) {
	for k, c := range inlcalls.Calls {
		if c.Root {
			psess.
				dumpInlCall(inlcalls, k, 0)
		}
	}
}

func (psess *PackageSession) dumpInlVars(dwvars []*dwarf.Var) {
	for i, dwv := range dwvars {
		typ := "local"
		if dwv.Abbrev == dwarf.DW_ABRV_PARAM_LOCLIST || dwv.Abbrev == dwarf.DW_ABRV_PARAM {
			typ = "param"
		}
		ia := 0
		if dwv.IsInAbstract {
			ia = 1
		}
		psess.
			Ctxt.Logf("V%d: %s CI:%d II:%d IA:%d %s\n", i, dwv.Name, dwv.ChildIndex, dwv.InlIndex-1, ia, typ)
	}
}

// Derived from Inferno utils/6c/txt.c
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/6c/txt.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
)

// Progs accumulates Progs for a function and converts them into machine code.
type Progs struct {
	Text      *obj.Prog  // ATEXT Prog for this function
	next      *obj.Prog  // next Prog
	pc        int64      // virtual PC; count of Progs
	pos       src.XPos   // position to use for new Progs
	curfn     *Node      // fn these Progs are for
	progcache []obj.Prog // local progcache
	cacheidx  int        // first free element of progcache

	nextLive LivenessIndex // liveness index for the next Prog
	prevLive LivenessIndex // last emitted liveness index
}

// newProgs returns a new Progs for fn.
// worker indicates which of the backend workers will use the Progs.
func (pstate *PackageState) newProgs(fn *Node, worker int) *Progs {
	pp := new(Progs)
	if pstate.Ctxt.CanReuseProgs() {
		sz := len(pstate.sharedProgArray) / pstate.nBackendWorkers
		pp.progcache = pstate.sharedProgArray[sz*worker : sz*(worker+1)]
	}
	pp.curfn = fn

	// prime the pump
	pp.next = pp.NewProg(pstate)
	pp.clearp(pp.next)

	pp.pos = fn.Pos
	pp.settext(pstate, fn)
	pp.nextLive = pstate.LivenessInvalid
	pp.prevLive = pstate.LivenessInvalid
	return pp
}

func (pp *Progs) NewProg(pstate *PackageState) *obj.Prog {
	var p *obj.Prog
	if pp.cacheidx < len(pp.progcache) {
		p = &pp.progcache[pp.cacheidx]
		pp.cacheidx++
	} else {
		p = new(obj.Prog)
	}
	p.Ctxt = pstate.Ctxt
	return p
}

// Flush converts from pp to machine code.
func (pp *Progs) Flush(pstate *PackageState) {
	plist := &obj.Plist{Firstpc: pp.Text, Curfn: pp.curfn}
	pstate.obj.Flushplist(pstate.Ctxt, plist, pp.NewProg, pstate.myimportpath)
}

// Free clears pp and any associated resources.
func (pp *Progs) Free(pstate *PackageState) {
	if pstate.Ctxt.CanReuseProgs() {
		// Clear progs to enable GC and avoid abuse.
		s := pp.progcache[:pp.cacheidx]
		for i := range s {
			s[i] = obj.Prog{}
		}
	}
	// Clear pp to avoid abuse.
	*pp = Progs{}
}

// Prog adds a Prog with instruction As to pp.
func (pp *Progs) Prog(pstate *PackageState, as obj.As) *obj.Prog {
	if pp.nextLive.stackMapIndex != pp.prevLive.stackMapIndex {
		// Emit stack map index change.
		idx := pp.nextLive.stackMapIndex
		pp.prevLive.stackMapIndex = idx
		p := pp.Prog(pstate, obj.APCDATA)
		Addrconst(&p.From, objabi.PCDATA_StackMapIndex)
		Addrconst(&p.To, int64(idx))
	}
	if pp.nextLive.regMapIndex != pp.prevLive.regMapIndex {
		// Emit register map index change.
		idx := pp.nextLive.regMapIndex
		pp.prevLive.regMapIndex = idx
		p := pp.Prog(pstate, obj.APCDATA)
		Addrconst(&p.From, objabi.PCDATA_RegMapIndex)
		Addrconst(&p.To, int64(idx))
	}

	p := pp.next
	pp.next = pp.NewProg(pstate)
	pp.clearp(pp.next)
	p.Link = pp.next

	if !pp.pos.IsKnown() && pstate.Debug['K'] != 0 {
		pstate.Warn("prog: unknown position (line 0)")
	}

	p.As = as
	p.Pos = pp.pos
	if pp.pos.IsStmt() == src.PosIsStmt {
		// Clear IsStmt for later Progs at this pos provided that as can be marked as a stmt
		if ssa.LosesStmtMark(as) {
			return p
		}
		pp.pos = pp.pos.WithNotStmt()
	}
	return p
}

func (pp *Progs) clearp(p *obj.Prog) {
	obj.Nopout(p)
	p.As = obj.AEND
	p.Pc = pp.pc
	pp.pc++
}

func (pp *Progs) Appendpp(pstate *PackageState, p *obj.Prog, as obj.As, ftype obj.AddrType, freg int16, foffset int64, ttype obj.AddrType, treg int16, toffset int64) *obj.Prog {
	q := pp.NewProg(pstate)
	pp.clearp(q)
	q.As = as
	q.Pos = p.Pos
	q.From.Type = ftype
	q.From.Reg = freg
	q.From.Offset = foffset
	q.To.Type = ttype
	q.To.Reg = treg
	q.To.Offset = toffset
	q.Link = p.Link
	p.Link = q
	return q
}

func (pp *Progs) settext(pstate *PackageState, fn *Node) {
	if pp.Text != nil {
		pstate.Fatalf("Progs.settext called twice")
	}
	ptxt := pp.Prog(pstate, obj.ATEXT)
	pp.Text = ptxt

	if fn.Func.lsym == nil {
		// func _() { }
		return
	}

	fn.Func.lsym.Func.Text = ptxt
	ptxt.From.Type = obj.TYPE_MEM
	ptxt.From.Name = obj.NAME_EXTERN
	ptxt.From.Sym = fn.Func.lsym

	p := pp.Prog(pstate, obj.AFUNCDATA)
	Addrconst(&p.From, objabi.FUNCDATA_ArgsPointerMaps)
	p.To.Type = obj.TYPE_MEM
	p.To.Name = obj.NAME_EXTERN
	p.To.Sym = &fn.Func.lsym.Func.GCArgs

	p = pp.Prog(pstate, obj.AFUNCDATA)
	Addrconst(&p.From, objabi.FUNCDATA_LocalsPointerMaps)
	p.To.Type = obj.TYPE_MEM
	p.To.Name = obj.NAME_EXTERN
	p.To.Sym = &fn.Func.lsym.Func.GCLocals

	p = pp.Prog(pstate, obj.AFUNCDATA)
	Addrconst(&p.From, objabi.FUNCDATA_RegPointerMaps)
	p.To.Type = obj.TYPE_MEM
	p.To.Name = obj.NAME_EXTERN
	p.To.Sym = &fn.Func.lsym.Func.GCRegs
}

func (f *Func) initLSym(pstate *PackageState) {
	if f.lsym != nil {
		pstate.Fatalf("Func.initLSym called twice")
	}

	if nam := f.Nname; !nam.isBlank() {
		f.lsym = nam.Sym.Linksym(pstate.types)
		if f.Pragma&Systemstack != 0 {
			f.lsym.Set(obj.AttrCFunc, true)
		}
	}

	var flag int
	if f.Dupok() {
		flag |= obj.DUPOK
	}
	if f.Wrapper() {
		flag |= obj.WRAPPER
	}
	if f.Needctxt() {
		flag |= obj.NEEDCTXT
	}
	if f.Pragma&Nosplit != 0 {
		flag |= obj.NOSPLIT
	}
	if f.ReflectMethod() {
		flag |= obj.REFLECTMETHOD
	}

	// Clumsy but important.
	// See test/recover.go for test cases and src/reflect/value.go
	// for the actual functions being considered.
	if pstate.myimportpath == "reflect" {
		switch f.Nname.Sym.Name {
		case "callReflect", "callMethod":
			flag |= obj.WRAPPER
		}
	}

	pstate.Ctxt.InitTextSym(f.lsym, flag)
}

func (pstate *PackageState) ggloblnod(nam *Node) {
	s := nam.Sym.Linksym(pstate.types)
	s.Gotype = pstate.ngotype(nam).Linksym(pstate.types)
	flags := 0
	if nam.Name.Readonly() {
		flags = obj.RODATA
	}
	if nam.Type != nil && !pstate.types.Haspointers(nam.Type) {
		flags |= obj.NOPTR
	}
	pstate.Ctxt.Globl(s, nam.Type.Width, flags)
}

func (pstate *PackageState) ggloblsym(s *obj.LSym, width int32, flags int16) {
	if flags&obj.LOCAL != 0 {
		s.Set(obj.AttrLocal, true)
		flags &^= obj.LOCAL
	}
	pstate.Ctxt.Globl(s, int64(width), int(flags))
}

func isfat(t *types.Type) bool {
	if t != nil {
		switch t.Etype {
		case TSTRUCT, TARRAY, TSLICE, TSTRING,
			TINTER: // maybe remove later
			return true
		}
	}

	return false
}

func Addrconst(a *obj.Addr, v int64) {
	a.Sym = nil
	a.Type = obj.TYPE_CONST
	a.Offset = v
}

func (pstate *PackageState) Patch(p *obj.Prog, to *obj.Prog) {
	if p.To.Type != obj.TYPE_BRANCH {
		pstate.Fatalf("patch: not a branch")
	}
	p.To.Val = to
	p.To.Offset = to.Pc
}

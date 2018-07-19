package gc

import (
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
)

// *T instead of T to work around issue 19839

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
func (psess *PackageSession) newProgs(fn *Node, worker int) *Progs {
	pp := new(Progs)
	if psess.Ctxt.CanReuseProgs() {
		sz := len(psess.sharedProgArray) / psess.nBackendWorkers
		pp.progcache = psess.sharedProgArray[sz*worker : sz*(worker+1)]
	}
	pp.curfn = fn

	pp.next = pp.NewProg(psess)
	pp.clearp(pp.next)

	pp.pos = fn.Pos
	pp.settext(psess, fn)
	pp.nextLive = psess.LivenessInvalid
	pp.prevLive = psess.LivenessInvalid
	return pp
}

func (pp *Progs) NewProg(psess *PackageSession) *obj.Prog {
	var p *obj.Prog
	if pp.cacheidx < len(pp.progcache) {
		p = &pp.progcache[pp.cacheidx]
		pp.cacheidx++
	} else {
		p = new(obj.Prog)
	}
	p.Ctxt = psess.Ctxt
	return p
}

// Flush converts from pp to machine code.
func (pp *Progs) Flush(psess *PackageSession) {
	plist := &obj.Plist{Firstpc: pp.Text, Curfn: pp.curfn}
	psess.obj.
		Flushplist(psess.Ctxt, plist, pp.NewProg, psess.myimportpath)
}

// Free clears pp and any associated resources.
func (pp *Progs) Free(psess *PackageSession) {
	if psess.Ctxt.CanReuseProgs() {

		s := pp.progcache[:pp.cacheidx]
		for i := range s {
			s[i] = obj.Prog{}
		}
	}

	*pp = Progs{}
}

// Prog adds a Prog with instruction As to pp.
func (pp *Progs) Prog(psess *PackageSession, as obj.As) *obj.Prog {
	if pp.nextLive.stackMapIndex != pp.prevLive.stackMapIndex {

		idx := pp.nextLive.stackMapIndex
		pp.prevLive.stackMapIndex = idx
		p := pp.Prog(psess, obj.APCDATA)
		Addrconst(&p.From, objabi.PCDATA_StackMapIndex)
		Addrconst(&p.To, int64(idx))
	}
	if pp.nextLive.regMapIndex != pp.prevLive.regMapIndex {

		idx := pp.nextLive.regMapIndex
		pp.prevLive.regMapIndex = idx
		p := pp.Prog(psess, obj.APCDATA)
		Addrconst(&p.From, objabi.PCDATA_RegMapIndex)
		Addrconst(&p.To, int64(idx))
	}

	p := pp.next
	pp.next = pp.NewProg(psess)
	pp.clearp(pp.next)
	p.Link = pp.next

	if !pp.pos.IsKnown() && psess.Debug['K'] != 0 {
		psess.
			Warn("prog: unknown position (line 0)")
	}

	p.As = as
	p.Pos = pp.pos
	if pp.pos.IsStmt() == src.PosIsStmt {

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

func (pp *Progs) Appendpp(psess *PackageSession, p *obj.Prog, as obj.As, ftype obj.AddrType, freg int16, foffset int64, ttype obj.AddrType, treg int16, toffset int64) *obj.Prog {
	q := pp.NewProg(psess)
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

func (pp *Progs) settext(psess *PackageSession, fn *Node) {
	if pp.Text != nil {
		psess.
			Fatalf("Progs.settext called twice")
	}
	ptxt := pp.Prog(psess, obj.ATEXT)
	pp.Text = ptxt

	if fn.Func.lsym == nil {

		return
	}

	fn.Func.lsym.Func.Text = ptxt
	ptxt.From.Type = obj.TYPE_MEM
	ptxt.From.Name = obj.NAME_EXTERN
	ptxt.From.Sym = fn.Func.lsym

	p := pp.Prog(psess, obj.AFUNCDATA)
	Addrconst(&p.From, objabi.FUNCDATA_ArgsPointerMaps)
	p.To.Type = obj.TYPE_MEM
	p.To.Name = obj.NAME_EXTERN
	p.To.Sym = &fn.Func.lsym.Func.GCArgs

	p = pp.Prog(psess, obj.AFUNCDATA)
	Addrconst(&p.From, objabi.FUNCDATA_LocalsPointerMaps)
	p.To.Type = obj.TYPE_MEM
	p.To.Name = obj.NAME_EXTERN
	p.To.Sym = &fn.Func.lsym.Func.GCLocals

	p = pp.Prog(psess, obj.AFUNCDATA)
	Addrconst(&p.From, objabi.FUNCDATA_RegPointerMaps)
	p.To.Type = obj.TYPE_MEM
	p.To.Name = obj.NAME_EXTERN
	p.To.Sym = &fn.Func.lsym.Func.GCRegs
}

func (f *Func) initLSym(psess *PackageSession) {
	if f.lsym != nil {
		psess.
			Fatalf("Func.initLSym called twice")
	}

	if nam := f.Nname; !nam.isBlank() {
		f.lsym = nam.Sym.Linksym(psess.types)
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

	if psess.myimportpath == "reflect" {
		switch f.Nname.Sym.Name {
		case "callReflect", "callMethod":
			flag |= obj.WRAPPER
		}
	}
	psess.
		Ctxt.InitTextSym(f.lsym, flag)
}

func (psess *PackageSession) ggloblnod(nam *Node) {
	s := nam.Sym.Linksym(psess.types)
	s.Gotype = psess.ngotype(nam).Linksym(psess.types)
	flags := 0
	if nam.Name.Readonly() {
		flags = obj.RODATA
	}
	if nam.Type != nil && !psess.types.Haspointers(nam.Type) {
		flags |= obj.NOPTR
	}
	psess.
		Ctxt.Globl(s, nam.Type.Width, flags)
}

func (psess *PackageSession) ggloblsym(s *obj.LSym, width int32, flags int16) {
	if flags&obj.LOCAL != 0 {
		s.Set(obj.AttrLocal, true)
		flags &^= obj.LOCAL
	}
	psess.
		Ctxt.Globl(s, int64(width), int(flags))
}

func isfat(t *types.Type) bool {
	if t != nil {
		switch t.Etype {
		case TSTRUCT, TARRAY, TSLICE, TSTRING,
			TINTER:
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

func (psess *PackageSession) Patch(p *obj.Prog, to *obj.Prog) {
	if p.To.Type != obj.TYPE_BRANCH {
		psess.
			Fatalf("patch: not a branch")
	}
	p.To.Val = to
	p.To.Offset = to.Pc
}

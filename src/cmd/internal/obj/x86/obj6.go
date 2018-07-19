package x86

import (
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"github.com/dave/golib/src/cmd/internal/sys"
	"math"
	"strings"
)

func (psess *PackageSession) CanUse1InsnTLS(ctxt *obj.Link) bool {
	if psess.isAndroid {

		return true
	}

	if ctxt.Arch.Family == sys.I386 {
		switch ctxt.Headtype {
		case objabi.Hlinux,
			objabi.Hnacl,
			objabi.Hplan9,
			objabi.Hwindows:
			return false
		}

		return true
	}

	switch ctxt.Headtype {
	case objabi.Hplan9, objabi.Hwindows:
		return false
	case objabi.Hlinux, objabi.Hfreebsd:
		return !ctxt.Flag_shared
	}

	return true
}

func (psess *PackageSession) progedit(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {

	if psess.CanUse1InsnTLS(ctxt) {

		if (p.As == AMOVQ || p.As == AMOVL) && p.From.Type == obj.TYPE_REG && p.From.Reg == REG_TLS && p.To.Type == obj.TYPE_REG && REG_AX <= p.To.Reg && p.To.Reg <= REG_R15 && ctxt.Headtype != objabi.Hsolaris {
			obj.Nopout(p)
		}
		if p.From.Type == obj.TYPE_MEM && p.From.Index == REG_TLS && REG_AX <= p.From.Reg && p.From.Reg <= REG_R15 {
			p.From.Reg = REG_TLS
			p.From.Scale = 0
			p.From.Index = REG_NONE
		}

		if p.To.Type == obj.TYPE_MEM && p.To.Index == REG_TLS && REG_AX <= p.To.Reg && p.To.Reg <= REG_R15 {
			p.To.Reg = REG_TLS
			p.To.Scale = 0
			p.To.Index = REG_NONE
		}
	} else {

		if (p.As == AMOVQ || p.As == AMOVL) && p.From.Type == obj.TYPE_MEM && p.From.Reg == REG_TLS && p.To.Type == obj.TYPE_REG && REG_AX <= p.To.Reg && p.To.Reg <= REG_R15 {
			q := obj.Appendp(p, newprog)
			q.As = p.As
			q.From = p.From
			q.From.Type = obj.TYPE_MEM
			q.From.Reg = p.To.Reg
			q.From.Index = REG_TLS
			q.From.Scale = 2
			q.To = p.To
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REG_TLS
			p.From.Index = REG_NONE
			p.From.Offset = 0
		}
	}

	if ctxt.Headtype == objabi.Hwindows && ctxt.Arch.Family == sys.AMD64 || ctxt.Headtype == objabi.Hplan9 {
		if p.From.Scale == 1 && p.From.Index == REG_TLS {
			p.From.Scale = 2
		}
		if p.To.Scale == 1 && p.To.Index == REG_TLS {
			p.To.Scale = 2
		}
	}

	switch p.As {
	case ACMPPD, ACMPPS, ACMPSD, ACMPSS:
		if p.To.Type == obj.TYPE_MEM && p.To.Name == obj.NAME_NONE && p.To.Reg == REG_NONE && p.To.Index == REG_NONE && p.To.Sym == nil {
			p.To.Type = obj.TYPE_CONST
		}
	}

	switch p.As {
	case obj.ACALL, obj.AJMP, obj.ARET:
		if p.To.Type == obj.TYPE_MEM && (p.To.Name == obj.NAME_EXTERN || p.To.Name == obj.NAME_STATIC) && p.To.Sym != nil {
			p.To.Type = obj.TYPE_BRANCH
		}
	}

	if p.From.Type == obj.TYPE_ADDR && (ctxt.Arch.Family == sys.AMD64 || p.From.Name != obj.NAME_EXTERN && p.From.Name != obj.NAME_STATIC) {
		switch p.As {
		case AMOVL:
			p.As = ALEAL
			p.From.Type = obj.TYPE_MEM
		case AMOVQ:
			p.As = ALEAQ
			p.From.Type = obj.TYPE_MEM
		}
	}

	if ctxt.Headtype == objabi.Hnacl && ctxt.Arch.Family == sys.AMD64 {
		if p.GetFrom3() != nil {
			nacladdr(ctxt, p, p.GetFrom3())
		}
		nacladdr(ctxt, p, &p.From)
		nacladdr(ctxt, p, &p.To)
	}

	switch p.As {

	case AMOVSS:
		if p.From.Type == obj.TYPE_FCONST {

			if f := p.From.Val.(float64); math.Float64bits(f) == 0 {
				if p.To.Type == obj.TYPE_REG && REG_X0 <= p.To.Reg && p.To.Reg <= REG_X15 {
					p.As = AXORPS
					p.From = p.To
					break
				}
			}
		}
		fallthrough

	case AFMOVF,
		AFADDF,
		AFSUBF,
		AFSUBRF,
		AFMULF,
		AFDIVF,
		AFDIVRF,
		AFCOMF,
		AFCOMFP,
		AADDSS,
		ASUBSS,
		AMULSS,
		ADIVSS,
		ACOMISS,
		AUCOMISS:
		if p.From.Type == obj.TYPE_FCONST {
			f32 := float32(p.From.Val.(float64))
			p.From.Type = obj.TYPE_MEM
			p.From.Name = obj.NAME_EXTERN
			p.From.Sym = ctxt.Float32Sym(f32)
			p.From.Offset = 0
		}

	case AMOVSD:

		if p.From.Type == obj.TYPE_FCONST {

			if f := p.From.Val.(float64); math.Float64bits(f) == 0 {
				if p.To.Type == obj.TYPE_REG && REG_X0 <= p.To.Reg && p.To.Reg <= REG_X15 {
					p.As = AXORPS
					p.From = p.To
					break
				}
			}
		}
		fallthrough

	case AFMOVD,
		AFADDD,
		AFSUBD,
		AFSUBRD,
		AFMULD,
		AFDIVD,
		AFDIVRD,
		AFCOMD,
		AFCOMDP,
		AADDSD,
		ASUBSD,
		AMULSD,
		ADIVSD,
		ACOMISD,
		AUCOMISD:
		if p.From.Type == obj.TYPE_FCONST {
			f64 := p.From.Val.(float64)
			p.From.Type = obj.TYPE_MEM
			p.From.Name = obj.NAME_EXTERN
			p.From.Sym = ctxt.Float64Sym(f64)
			p.From.Offset = 0
		}
	}

	if ctxt.Flag_dynlink {
		rewriteToUseGot(ctxt, p, newprog)
	}

	if ctxt.Flag_shared && ctxt.Arch.Family == sys.I386 {
		psess.
			rewriteToPcrel(ctxt, p, newprog)
	}
}

// Rewrite p, if necessary, to access global data via the global offset table.
func rewriteToUseGot(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {
	var lea, mov obj.As
	var reg int16
	if ctxt.Arch.Family == sys.AMD64 {
		lea = ALEAQ
		mov = AMOVQ
		reg = REG_R15
	} else {
		lea = ALEAL
		mov = AMOVL
		reg = REG_CX
		if p.As == ALEAL && p.To.Reg != p.From.Reg && p.To.Reg != p.From.Index {

			reg = p.To.Reg
		}
	}

	if p.As == obj.ADUFFCOPY || p.As == obj.ADUFFZERO {
		//     ADUFFxxx $offset
		// becomes
		//     $MOV runtime.duffxxx@GOT, $reg
		//     $LEA $offset($reg), $reg
		//     CALL $reg
		// (we use LEAx rather than ADDx because ADDx clobbers
		// flags and duffzero on 386 does not otherwise do so).
		var sym *obj.LSym
		if p.As == obj.ADUFFZERO {
			sym = ctxt.Lookup("runtime.duffzero")
		} else {
			sym = ctxt.Lookup("runtime.duffcopy")
		}
		offset := p.To.Offset
		p.As = mov
		p.From.Type = obj.TYPE_MEM
		p.From.Name = obj.NAME_GOTREF
		p.From.Sym = sym
		p.To.Type = obj.TYPE_REG
		p.To.Reg = reg
		p.To.Offset = 0
		p.To.Sym = nil
		p1 := obj.Appendp(p, newprog)
		p1.As = lea
		p1.From.Type = obj.TYPE_MEM
		p1.From.Offset = offset
		p1.From.Reg = reg
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = reg
		p2 := obj.Appendp(p1, newprog)
		p2.As = obj.ACALL
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = reg
	}

	if p.As == lea && p.From.Type == obj.TYPE_MEM && p.From.Name == obj.NAME_EXTERN && !p.From.Sym.Local() {

		p.As = mov
		p.From.Type = obj.TYPE_ADDR
	}
	if p.From.Type == obj.TYPE_ADDR && p.From.Name == obj.NAME_EXTERN && !p.From.Sym.Local() {

		cmplxdest := false
		pAs := p.As
		var dest obj.Addr
		if p.To.Type != obj.TYPE_REG || pAs != mov {
			if ctxt.Arch.Family == sys.AMD64 {
				ctxt.Diag("do not know how to handle LEA-type insn to non-register in %v with -dynlink", p)
			}
			cmplxdest = true
			dest = p.To
			p.As = mov
			p.To.Type = obj.TYPE_REG
			p.To.Reg = reg
			p.To.Sym = nil
			p.To.Name = obj.NAME_NONE
		}
		p.From.Type = obj.TYPE_MEM
		p.From.Name = obj.NAME_GOTREF
		q := p
		if p.From.Offset != 0 {
			q = obj.Appendp(p, newprog)
			q.As = lea
			q.From.Type = obj.TYPE_MEM
			q.From.Reg = p.To.Reg
			q.From.Offset = p.From.Offset
			q.To = p.To
			p.From.Offset = 0
		}
		if cmplxdest {
			q = obj.Appendp(q, newprog)
			q.As = pAs
			q.To = dest
			q.From.Type = obj.TYPE_REG
			q.From.Reg = reg
		}
	}
	if p.GetFrom3() != nil && p.GetFrom3().Name == obj.NAME_EXTERN {
		ctxt.Diag("don't know how to handle %v with -dynlink", p)
	}
	var source *obj.Addr

	if p.From.Name == obj.NAME_EXTERN && !p.From.Sym.Local() {
		if p.To.Name == obj.NAME_EXTERN && !p.To.Sym.Local() {
			ctxt.Diag("cannot handle NAME_EXTERN on both sides in %v with -dynlink", p)
		}
		source = &p.From
	} else if p.To.Name == obj.NAME_EXTERN && !p.To.Sym.Local() {
		source = &p.To
	} else {
		return
	}
	if p.As == obj.ACALL {

		if ctxt.Arch.Family == sys.AMD64 || (p.To.Sym != nil && p.To.Sym.Local()) || p.RegTo2 != 0 {
			return
		}
		p1 := obj.Appendp(p, newprog)
		p2 := obj.Appendp(p1, newprog)

		p1.As = ALEAL
		p1.From.Type = obj.TYPE_MEM
		p1.From.Name = obj.NAME_STATIC
		p1.From.Sym = ctxt.Lookup("_GLOBAL_OFFSET_TABLE_")
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = REG_BX

		p2.As = p.As
		p2.Scond = p.Scond
		p2.From = p.From
		if p.RestArgs != nil {
			p2.RestArgs = append(p2.RestArgs, p.RestArgs...)
		}
		p2.Reg = p.Reg
		p2.To = p.To

		p2.To.Type = obj.TYPE_MEM
		p2.RegTo2 = 1

		obj.Nopout(p)
		return

	}
	if p.As == obj.ATEXT || p.As == obj.AFUNCDATA || p.As == obj.ARET || p.As == obj.AJMP {
		return
	}
	if source.Type != obj.TYPE_MEM {
		ctxt.Diag("don't know how to handle %v with -dynlink", p)
	}
	p1 := obj.Appendp(p, newprog)
	p2 := obj.Appendp(p1, newprog)

	p1.As = mov
	p1.From.Type = obj.TYPE_MEM
	p1.From.Sym = source.Sym
	p1.From.Name = obj.NAME_GOTREF
	p1.To.Type = obj.TYPE_REG
	p1.To.Reg = reg

	p2.As = p.As
	p2.From = p.From
	p2.To = p.To
	if p.From.Name == obj.NAME_EXTERN {
		p2.From.Reg = reg
		p2.From.Name = obj.NAME_NONE
		p2.From.Sym = nil
	} else if p.To.Name == obj.NAME_EXTERN {
		p2.To.Reg = reg
		p2.To.Name = obj.NAME_NONE
		p2.To.Sym = nil
	} else {
		return
	}
	obj.Nopout(p)
}

func (psess *PackageSession) rewriteToPcrel(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {

	if p.RegTo2 != 0 {
		return
	}
	if p.As == obj.ATEXT || p.As == obj.AFUNCDATA || p.As == obj.ACALL || p.As == obj.ARET || p.As == obj.AJMP {
		return
	}

	isName := func(a *obj.Addr) bool {
		if a.Sym == nil || (a.Type != obj.TYPE_MEM && a.Type != obj.TYPE_ADDR) || a.Reg != 0 {
			return false
		}
		if a.Sym.Type == objabi.STLSBSS {
			return false
		}
		return a.Name == obj.NAME_EXTERN || a.Name == obj.NAME_STATIC || a.Name == obj.NAME_GOTREF
	}

	if isName(&p.From) && p.From.Type == obj.TYPE_ADDR {

		if p.To.Type != obj.TYPE_REG {
			q := obj.Appendp(p, newprog)
			q.As = p.As
			q.From.Type = obj.TYPE_REG
			q.From.Reg = REG_CX
			q.To = p.To
			p.As = AMOVL
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REG_CX
			p.To.Sym = nil
			p.To.Name = obj.NAME_NONE
		}
	}

	if !isName(&p.From) && !isName(&p.To) && (p.GetFrom3() == nil || !isName(p.GetFrom3())) {
		return
	}
	var dst int16 = REG_CX
	if (p.As == ALEAL || p.As == AMOVL) && p.To.Reg != p.From.Reg && p.To.Reg != p.From.Index {
		dst = p.To.Reg

	}
	q := obj.Appendp(p, newprog)
	q.RegTo2 = 1
	r := obj.Appendp(q, newprog)
	r.RegTo2 = 1
	q.As = obj.ACALL
	thunkname := "__x86.get_pc_thunk." + strings.ToLower(psess.rconv(int(dst)))
	q.To.Sym = ctxt.LookupInit(thunkname, func(s *obj.LSym) { s.Set(obj.AttrLocal, true) })
	q.To.Type = obj.TYPE_MEM
	q.To.Name = obj.NAME_EXTERN
	r.As = p.As
	r.Scond = p.Scond
	r.From = p.From
	r.RestArgs = p.RestArgs
	r.Reg = p.Reg
	r.To = p.To
	if isName(&p.From) {
		r.From.Reg = dst
	}
	if isName(&p.To) {
		r.To.Reg = dst
	}
	if p.GetFrom3() != nil && isName(p.GetFrom3()) {
		r.GetFrom3().Reg = dst
	}
	obj.Nopout(p)
}

func nacladdr(ctxt *obj.Link, p *obj.Prog, a *obj.Addr) {
	if p.As == ALEAL || p.As == ALEAQ {
		return
	}

	if a.Reg == REG_BP {
		ctxt.Diag("invalid address: %v", p)
		return
	}

	if a.Reg == REG_TLS {
		a.Reg = REG_BP
	}
	if a.Type == obj.TYPE_MEM && a.Name == obj.NAME_NONE {
		switch a.Reg {

		case REG_BP, REG_SP, REG_R15:
			break

		default:
			if a.Index != REG_NONE {
				ctxt.Diag("invalid address %v", p)
			}
			a.Index = a.Reg
			if a.Index != REG_NONE {
				a.Scale = 1
			}
			a.Reg = REG_R15
		}
	}
}

func (psess *PackageSession) preprocess(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	if cursym.Func.Text == nil || cursym.Func.Text.Link == nil {
		return
	}

	p := cursym.Func.Text
	autoffset := int32(p.To.Offset)
	if autoffset < 0 {
		autoffset = 0
	}

	hasCall := false
	for q := p; q != nil; q = q.Link {
		if q.As == obj.ACALL || q.As == obj.ADUFFCOPY || q.As == obj.ADUFFZERO {
			hasCall = true
			break
		}
	}

	var bpsize int
	if ctxt.Arch.Family == sys.AMD64 && ctxt.Framepointer_enabled &&
		!p.From.Sym.NoFrame() &&
		!(autoffset == 0 && p.From.Sym.NoSplit()) &&
		!(autoffset == 0 && !hasCall) {

		bpsize = ctxt.Arch.PtrSize
		autoffset += int32(bpsize)
		p.To.Offset += int64(bpsize)
	} else {
		bpsize = 0
	}

	textarg := int64(p.To.Val.(int32))
	cursym.Func.Args = int32(textarg)
	cursym.Func.Locals = int32(p.To.Offset)

	if ctxt.Arch.Family == sys.I386 && cursym.Func.Locals < 0 {
		cursym.Func.Locals = 0
	}

	if ctxt.Arch.Family == sys.AMD64 && autoffset < objabi.StackSmall && !p.From.Sym.NoSplit() {
		leaf := true
	LeafSearch:
		for q := p; q != nil; q = q.Link {
			switch q.As {
			case obj.ACALL:

				if !isZeroArgRuntimeCall(q.To.Sym) {
					leaf = false
					break LeafSearch
				}
				fallthrough
			case obj.ADUFFCOPY, obj.ADUFFZERO:
				if autoffset >= objabi.StackSmall-8 {
					leaf = false
					break LeafSearch
				}
			}
		}

		if leaf {
			p.From.Sym.Set(obj.AttrNoSplit, true)
		}
	}

	if !p.From.Sym.NoSplit() || p.From.Sym.Wrapper() {
		p = obj.Appendp(p, newprog)
		p = psess.load_g_cx(ctxt, p, newprog)
	}

	if !cursym.Func.Text.From.Sym.NoSplit() {
		p = psess.stacksplit(ctxt, cursym, p, newprog, autoffset, int32(textarg))
	}

	markedPrologue := false

	if autoffset != 0 {
		if autoffset%int32(ctxt.Arch.RegSize) != 0 {
			ctxt.Diag("unaligned stack size %d", autoffset)
		}
		p = obj.Appendp(p, newprog)
		p.As = AADJSP
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = int64(autoffset)
		p.Spadj = autoffset
		p.Pos = p.Pos.WithXlogue(src.PosPrologueEnd)
		markedPrologue = true
	}

	deltasp := autoffset

	if bpsize > 0 {

		p = obj.Appendp(p, newprog)

		p.As = AMOVQ
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_BP
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = REG_SP
		p.To.Scale = 1
		p.To.Offset = int64(autoffset) - int64(bpsize)
		if !markedPrologue {
			p.Pos = p.Pos.WithXlogue(src.PosPrologueEnd)
		}

		p = obj.Appendp(p, newprog)

		p.As = ALEAQ
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = REG_SP
		p.From.Scale = 1
		p.From.Offset = int64(autoffset) - int64(bpsize)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_BP
	}

	if cursym.Func.Text.From.Sym.Wrapper() {

		p = obj.Appendp(p, newprog)
		p.As = AMOVQ
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = REG_CX
		p.From.Offset = 4 * int64(ctxt.Arch.PtrSize)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_BX
		if ctxt.Headtype == objabi.Hnacl && ctxt.Arch.Family == sys.AMD64 {
			p.As = AMOVL
			p.From.Type = obj.TYPE_MEM
			p.From.Reg = REG_R15
			p.From.Scale = 1
			p.From.Index = REG_CX
		}
		if ctxt.Arch.Family == sys.I386 {
			p.As = AMOVL
		}

		p = obj.Appendp(p, newprog)
		p.As = ATESTQ
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_BX
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_BX
		if ctxt.Headtype == objabi.Hnacl || ctxt.Arch.Family == sys.I386 {
			p.As = ATESTL
		}

		jne := obj.Appendp(p, newprog)
		jne.As = AJNE
		jne.To.Type = obj.TYPE_BRANCH

		end := obj.Appendp(jne, newprog)
		end.As = obj.ANOP

		// Fast forward to end of function.
		var last *obj.Prog
		for last = end; last.Link != nil; last = last.Link {
		}

		p = obj.Appendp(last, newprog)
		p.As = ALEAQ
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = REG_SP
		p.From.Offset = int64(autoffset) + int64(ctxt.Arch.RegSize)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_DI
		if ctxt.Headtype == objabi.Hnacl || ctxt.Arch.Family == sys.I386 {
			p.As = ALEAL
		}

		jne.Pcond = p

		p = obj.Appendp(p, newprog)
		p.As = ACMPQ
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = REG_BX
		p.From.Offset = 0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_DI
		if ctxt.Headtype == objabi.Hnacl && ctxt.Arch.Family == sys.AMD64 {
			p.As = ACMPL
			p.From.Type = obj.TYPE_MEM
			p.From.Reg = REG_R15
			p.From.Scale = 1
			p.From.Index = REG_BX
		}
		if ctxt.Arch.Family == sys.I386 {
			p.As = ACMPL
		}

		p = obj.Appendp(p, newprog)
		p.As = AJNE
		p.To.Type = obj.TYPE_BRANCH
		p.Pcond = end

		p = obj.Appendp(p, newprog)
		p.As = AMOVQ
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_SP
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = REG_BX
		p.To.Offset = 0
		if ctxt.Headtype == objabi.Hnacl && ctxt.Arch.Family == sys.AMD64 {
			p.As = AMOVL
			p.To.Type = obj.TYPE_MEM
			p.To.Reg = REG_R15
			p.To.Scale = 1
			p.To.Index = REG_BX
		}
		if ctxt.Arch.Family == sys.I386 {
			p.As = AMOVL
		}

		p = obj.Appendp(p, newprog)
		p.As = obj.AJMP
		p.To.Type = obj.TYPE_BRANCH
		p.Pcond = end

		p = end
	}

	for ; p != nil; p = p.Link {
		pcsize := ctxt.Arch.RegSize
		switch p.From.Name {
		case obj.NAME_AUTO:
			p.From.Offset += int64(deltasp) - int64(bpsize)
		case obj.NAME_PARAM:
			p.From.Offset += int64(deltasp) + int64(pcsize)
		}
		if p.GetFrom3() != nil {
			switch p.GetFrom3().Name {
			case obj.NAME_AUTO:
				p.GetFrom3().Offset += int64(deltasp) - int64(bpsize)
			case obj.NAME_PARAM:
				p.GetFrom3().Offset += int64(deltasp) + int64(pcsize)
			}
		}
		switch p.To.Name {
		case obj.NAME_AUTO:
			p.To.Offset += int64(deltasp) - int64(bpsize)
		case obj.NAME_PARAM:
			p.To.Offset += int64(deltasp) + int64(pcsize)
		}

		switch p.As {
		default:
			continue

		case APUSHL, APUSHFL:
			deltasp += 4
			p.Spadj = 4
			continue

		case APUSHQ, APUSHFQ:
			deltasp += 8
			p.Spadj = 8
			continue

		case APUSHW, APUSHFW:
			deltasp += 2
			p.Spadj = 2
			continue

		case APOPL, APOPFL:
			deltasp -= 4
			p.Spadj = -4
			continue

		case APOPQ, APOPFQ:
			deltasp -= 8
			p.Spadj = -8
			continue

		case APOPW, APOPFW:
			deltasp -= 2
			p.Spadj = -2
			continue

		case obj.ARET:

		}

		if autoffset != deltasp {
			ctxt.Diag("unbalanced PUSH/POP")
		}

		if autoffset != 0 {
			to := p.To
			p.To = obj.Addr{}
			if bpsize > 0 {

				p.As = AMOVQ

				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REG_SP
				p.From.Scale = 1
				p.From.Offset = int64(autoffset) - int64(bpsize)
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_BP
				p = obj.Appendp(p, newprog)
			}

			p.As = AADJSP
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = int64(-autoffset)
			p.Spadj = -autoffset
			p = obj.Appendp(p, newprog)
			p.As = obj.ARET
			p.To = to

			p.Spadj = +autoffset
		}

		if p.To.Sym != nil {
			p.As = obj.AJMP
		}
	}
}

func isZeroArgRuntimeCall(s *obj.LSym) bool {
	if s == nil {
		return false
	}
	switch s.Name {
	case "runtime.panicindex", "runtime.panicslice", "runtime.panicdivide", "runtime.panicwrap":
		return true
	}
	return false
}

func indir_cx(ctxt *obj.Link, a *obj.Addr) {
	if ctxt.Headtype == objabi.Hnacl && ctxt.Arch.Family == sys.AMD64 {
		a.Type = obj.TYPE_MEM
		a.Reg = REG_R15
		a.Index = REG_CX
		a.Scale = 1
		return
	}

	a.Type = obj.TYPE_MEM
	a.Reg = REG_CX
}

// Append code to p to load g into cx.
// Overwrites p with the first instruction (no first appendp).
// Overwriting p is unusual but it lets use this in both the
// prologue (caller must call appendp first) and in the epilogue.
// Returns last new instruction.
func (psess *PackageSession) load_g_cx(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) *obj.Prog {
	p.As = AMOVQ
	if ctxt.Arch.PtrSize == 4 {
		p.As = AMOVL
	}
	p.From.Type = obj.TYPE_MEM
	p.From.Reg = REG_TLS
	p.From.Offset = 0
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_CX

	next := p.Link
	psess.
		progedit(ctxt, p, newprog)
	for p.Link != next {
		p = p.Link
	}

	if p.From.Index == REG_TLS {
		p.From.Scale = 2
	}

	return p
}

// Append code to p to check for stack split.
// Appends to (does not overwrite) p.
// Assumes g is in CX.
// Returns last new instruction.
func (psess *PackageSession) stacksplit(ctxt *obj.Link, cursym *obj.LSym, p *obj.Prog, newprog obj.ProgAlloc, framesize int32, textarg int32) *obj.Prog {
	cmp := ACMPQ
	lea := ALEAQ
	mov := AMOVQ
	sub := ASUBQ

	if ctxt.Headtype == objabi.Hnacl || ctxt.Arch.Family == sys.I386 {
		cmp = ACMPL
		lea = ALEAL
		mov = AMOVL
		sub = ASUBL
	}

	var q1 *obj.Prog
	if framesize <= objabi.StackSmall {

		p = obj.Appendp(p, newprog)

		p.As = cmp
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_SP
		indir_cx(ctxt, &p.To)
		p.To.Offset = 2 * int64(ctxt.Arch.PtrSize)
		if cursym.CFunc() {
			p.To.Offset = 3 * int64(ctxt.Arch.PtrSize)
		}
	} else if framesize <= objabi.StackBig {

		p = obj.Appendp(p, newprog)

		p.As = lea
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = REG_SP
		p.From.Offset = -(int64(framesize) - objabi.StackSmall)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_AX

		p = obj.Appendp(p, newprog)
		p.As = cmp
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_AX
		indir_cx(ctxt, &p.To)
		p.To.Offset = 2 * int64(ctxt.Arch.PtrSize)
		if cursym.CFunc() {
			p.To.Offset = 3 * int64(ctxt.Arch.PtrSize)
		}
	} else {

		p = obj.Appendp(p, newprog)

		p.As = mov
		indir_cx(ctxt, &p.From)
		p.From.Offset = 2 * int64(ctxt.Arch.PtrSize)
		if cursym.CFunc() {
			p.From.Offset = 3 * int64(ctxt.Arch.PtrSize)
		}
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_SI

		p = obj.Appendp(p, newprog)
		p.As = cmp
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_SI
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = objabi.StackPreempt
		if ctxt.Arch.Family == sys.I386 {
			p.To.Offset = int64(uint32(objabi.StackPreempt & (1<<32 - 1)))
		}

		p = obj.Appendp(p, newprog)
		p.As = AJEQ
		p.To.Type = obj.TYPE_BRANCH
		q1 = p

		p = obj.Appendp(p, newprog)
		p.As = lea
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = REG_SP
		p.From.Offset = objabi.StackGuard
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_AX

		p = obj.Appendp(p, newprog)
		p.As = sub
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_SI
		p.To.Type = obj.TYPE_REG
		p.To.Reg = REG_AX

		p = obj.Appendp(p, newprog)
		p.As = cmp
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_AX
		p.To.Type = obj.TYPE_CONST
		p.To.Offset = int64(framesize) + (objabi.StackGuard - objabi.StackSmall)
	}

	jls := obj.Appendp(p, newprog)
	jls.As = AJLS
	jls.To.Type = obj.TYPE_BRANCH

	var last *obj.Prog
	for last = cursym.Func.Text; last.Link != nil; last = last.Link {
	}

	spfix := obj.Appendp(last, newprog)
	spfix.As = obj.ANOP
	spfix.Spadj = -framesize

	pcdata := ctxt.EmitEntryLiveness(cursym, spfix, newprog)

	call := obj.Appendp(pcdata, newprog)
	call.Pos = cursym.Func.Text.Pos
	call.As = obj.ACALL
	call.To.Type = obj.TYPE_BRANCH
	call.To.Name = obj.NAME_EXTERN
	morestack := "runtime.morestack"
	switch {
	case cursym.CFunc():
		morestack = "runtime.morestackc"
	case !cursym.Func.Text.From.Sym.NeedCtxt():
		morestack = "runtime.morestack_noctxt"
	}
	call.To.Sym = ctxt.Lookup(morestack)

	callend := call
	psess.
		progedit(ctxt, callend, newprog)
	for ; callend.Link != nil; callend = callend.Link {
		psess.
			progedit(ctxt, callend.Link, newprog)
	}

	jmp := obj.Appendp(callend, newprog)
	jmp.As = obj.AJMP
	jmp.To.Type = obj.TYPE_BRANCH
	jmp.Pcond = cursym.Func.Text.Link
	jmp.Spadj = +framesize

	jls.Pcond = call
	if q1 != nil {
		q1.Pcond = call
	}

	return jls
}

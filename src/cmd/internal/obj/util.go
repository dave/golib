package obj

import (
	"bytes"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"strings"
)

const REG_NONE = 0

// Line returns a string containing the filename and line number for p
func (p *Prog) Line(psess *PackageSession) string {
	return p.Ctxt.OutermostPos(p.Pos).Format(psess.src, false, true)
}

// InnermostLineNumber returns a string containing the line number for the
// innermost inlined function (if any inlining) at p's position
func (p *Prog) InnermostLineNumber() string {
	return p.Ctxt.InnermostPos(p.Pos).LineNumber()
}

// InnermostLineNumberHTML returns a string containing the line number for the
// innermost inlined function (if any inlining) at p's position
func (p *Prog) InnermostLineNumberHTML() string {
	return p.Ctxt.InnermostPos(p.Pos).LineNumberHTML()
}

// InnermostFilename returns a string containing the innermost
// (in inlining) filename at p's position
func (p *Prog) InnermostFilename(psess *PackageSession) string {

	pos := p.Ctxt.InnermostPos(p.Pos)
	if !pos.IsKnown() {
		return "<unknown file name>"
	}
	return pos.Filename(psess.src)
}

/* ARM scond byte */
const (
	C_SCOND     = (1 << 4) - 1
	C_SBIT      = 1 << 4
	C_PBIT      = 1 << 5
	C_WBIT      = 1 << 6
	C_FBIT      = 1 << 7
	C_UBIT      = 1 << 7
	C_SCOND_XOR = 14
)

// CConv formats opcode suffix bits (Prog.Scond).
func (psess *PackageSession) CConv(s uint8) string {
	if s == 0 {
		return ""
	}
	for i := range psess.opSuffixSpace {
		sset := &psess.opSuffixSpace[i]
		if sset.arch == psess.objabi.GOARCH {
			return sset.cconv(s)
		}
	}
	return fmt.Sprintf("SC???%d", s)
}

// CConvARM formats ARM opcode suffix bits (mostly condition codes).
func (psess *PackageSession) CConvARM(s uint8) string {

	sc := psess.armCondCode[(s&C_SCOND)^C_SCOND_XOR]
	if s&C_SBIT != 0 {
		sc += ".S"
	}
	if s&C_PBIT != 0 {
		sc += ".P"
	}
	if s&C_WBIT != 0 {
		sc += ".W"
	}
	if s&C_UBIT != 0 {
		sc += ".U"
	}
	return sc
}

func (p *Prog) String(psess *PackageSession) string {
	if p == nil {
		return "<nil Prog>"
	}
	if p.Ctxt == nil {
		return "<Prog without ctxt>"
	}
	return fmt.Sprintf("%.5d (%v)\t%s", p.Pc, p.Line(psess), p.InstructionString(psess))
}

// InstructionString returns a string representation of the instruction without preceding
// program counter or file and line number.
func (p *Prog) InstructionString(psess *PackageSession) string {
	if p == nil {
		return "<nil Prog>"
	}

	if p.Ctxt == nil {
		return "<Prog without ctxt>"
	}

	sc := psess.CConv(p.Scond)

	var buf bytes.Buffer

	fmt.Fprintf(&buf, "%v%s", p.As, sc)
	sep := "\t"

	if p.From.Type != TYPE_NONE {
		fmt.Fprintf(&buf, "%s%v", sep, psess.Dconv(p, &p.From))
		sep = ", "
	}
	if p.Reg != REG_NONE {

		fmt.Fprintf(&buf, "%s%v", sep, psess.Rconv(int(p.Reg)))
		sep = ", "
	}
	for i := range p.RestArgs {
		fmt.Fprintf(&buf, "%s%v", sep, psess.Dconv(p, &p.RestArgs[i]))
		sep = ", "
	}

	if p.As == ATEXT {

		s := p.From.Sym.Attribute.TextAttrString(psess)
		if s != "" {
			fmt.Fprintf(&buf, "%s%s", sep, s)
			sep = ", "
		}
	}
	if p.To.Type != TYPE_NONE {
		fmt.Fprintf(&buf, "%s%v", sep, psess.Dconv(p, &p.To))
	}
	if p.RegTo2 != REG_NONE {
		fmt.Fprintf(&buf, "%s%v", sep, psess.Rconv(int(p.RegTo2)))
	}
	return buf.String()
}

func (ctxt *Link) NewProg() *Prog {
	p := new(Prog)
	p.Ctxt = ctxt
	return p
}

func (ctxt *Link) CanReuseProgs() bool {
	return !ctxt.Debugasm
}

func (psess *PackageSession) Dconv(p *Prog, a *Addr) string {
	var str string

	switch a.Type {
	default:
		str = fmt.Sprintf("type=%d", a.Type)

	case TYPE_NONE:
		str = ""
		if a.Name != NAME_NONE || a.Reg != 0 || a.Sym != nil {
			str = fmt.Sprintf("%v(%v)(NONE)", psess.Mconv(a), psess.Rconv(int(a.Reg)))
		}

	case TYPE_REG:

		if a.Offset != 0 && (a.Reg < RBaseARM64 || a.Reg >= RBaseMIPS) {
			str = fmt.Sprintf("$%d,%v", a.Offset, psess.Rconv(int(a.Reg)))
			break
		}

		str = psess.Rconv(int(a.Reg))
		if a.Name != NAME_NONE || a.Sym != nil {
			str = fmt.Sprintf("%v(%v)(REG)", psess.Mconv(a), psess.Rconv(int(a.Reg)))
		}
		if (RBaseARM64+1<<10+1<<9) <= a.Reg &&
			a.Reg < (RBaseARM64+1<<11) {
			str += fmt.Sprintf("[%d]", a.Index)
		}

	case TYPE_BRANCH:
		if a.Sym != nil {
			str = fmt.Sprintf("%s(SB)", a.Sym.Name)
		} else if p != nil && p.Pcond != nil {
			str = fmt.Sprint(p.Pcond.Pc)
		} else if a.Val != nil {
			str = fmt.Sprint(a.Val.(*Prog).Pc)
		} else {
			str = fmt.Sprintf("%d(PC)", a.Offset)
		}

	case TYPE_INDIR:
		str = fmt.Sprintf("*%s", psess.Mconv(a))

	case TYPE_MEM:
		str = psess.Mconv(a)
		if a.Index != REG_NONE {
			if a.Scale == 0 {

				str += fmt.Sprintf("(%v)", psess.Rconv(int(a.Index)))
			} else {
				str += fmt.Sprintf("(%v*%d)", psess.Rconv(int(a.Index)), int(a.Scale))
			}
		}

	case TYPE_CONST:
		if a.Reg != 0 {
			str = fmt.Sprintf("$%v(%v)", psess.Mconv(a), psess.Rconv(int(a.Reg)))
		} else {
			str = fmt.Sprintf("$%v", psess.Mconv(a))
		}

	case TYPE_TEXTSIZE:
		if a.Val.(int32) == objabi.ArgsSizeUnknown {
			str = fmt.Sprintf("$%d", a.Offset)
		} else {
			str = fmt.Sprintf("$%d-%d", a.Offset, a.Val.(int32))
		}

	case TYPE_FCONST:
		str = fmt.Sprintf("%.17g", a.Val.(float64))

		if !strings.ContainsAny(str, ".e") {
			str += ".0"
		}
		str = fmt.Sprintf("$(%s)", str)

	case TYPE_SCONST:
		str = fmt.Sprintf("$%q", a.Val.(string))

	case TYPE_ADDR:
		str = fmt.Sprintf("$%s", psess.Mconv(a))

	case TYPE_SHIFT:
		v := int(a.Offset)
		ops := "<<>>->@>"
		switch psess.objabi.GOARCH {
		case "arm":
			op := ops[((v>>5)&3)<<1:]
			if v&(1<<4) != 0 {
				str = fmt.Sprintf("R%d%c%cR%d", v&15, op[0], op[1], (v>>8)&15)
			} else {
				str = fmt.Sprintf("R%d%c%c%d", v&15, op[0], op[1], (v>>7)&31)
			}
			if a.Reg != 0 {
				str += fmt.Sprintf("(%v)", psess.Rconv(int(a.Reg)))
			}
		case "arm64":
			op := ops[((v>>22)&3)<<1:]
			r := (v >> 16) & 31
			str = fmt.Sprintf("%s%c%c%d", psess.Rconv(r+RBaseARM64), op[0], op[1], (v>>10)&63)
		default:
			panic("TYPE_SHIFT is not supported on " + psess.objabi.GOARCH)
		}

	case TYPE_REGREG:
		str = fmt.Sprintf("(%v, %v)", psess.Rconv(int(a.Reg)), psess.Rconv(int(a.Offset)))

	case TYPE_REGREG2:
		str = fmt.Sprintf("%v, %v", psess.Rconv(int(a.Offset)), psess.Rconv(int(a.Reg)))

	case TYPE_REGLIST:
		str = psess.RLconv(a.Offset)
	}

	return str
}

func (psess *PackageSession) Mconv(a *Addr) string {
	var str string

	switch a.Name {
	default:
		str = fmt.Sprintf("name=%d", a.Name)

	case NAME_NONE:
		switch {
		case a.Reg == REG_NONE:
			str = fmt.Sprint(a.Offset)
		case a.Offset == 0:
			str = fmt.Sprintf("(%v)", psess.Rconv(int(a.Reg)))
		case a.Offset != 0:
			str = fmt.Sprintf("%d(%v)", a.Offset, psess.Rconv(int(a.Reg)))
		}

	case NAME_EXTERN:
		reg := "SB"
		if a.Reg != REG_NONE {
			reg = psess.Rconv(int(a.Reg))
		}
		if a.Sym != nil {
			str = fmt.Sprintf("%s%s(%s)", a.Sym.Name, offConv(a.Offset), reg)
		} else {
			str = fmt.Sprintf("%s(%s)", offConv(a.Offset), reg)
		}

	case NAME_GOTREF:
		reg := "SB"
		if a.Reg != REG_NONE {
			reg = psess.Rconv(int(a.Reg))
		}
		if a.Sym != nil {
			str = fmt.Sprintf("%s%s@GOT(%s)", a.Sym.Name, offConv(a.Offset), reg)
		} else {
			str = fmt.Sprintf("%s@GOT(%s)", offConv(a.Offset), reg)
		}

	case NAME_STATIC:
		reg := "SB"
		if a.Reg != REG_NONE {
			reg = psess.Rconv(int(a.Reg))
		}
		if a.Sym != nil {
			str = fmt.Sprintf("%s<>%s(%s)", a.Sym.Name, offConv(a.Offset), reg)
		} else {
			str = fmt.Sprintf("<>%s(%s)", offConv(a.Offset), reg)
		}

	case NAME_AUTO:
		reg := "SP"
		if a.Reg != REG_NONE {
			reg = psess.Rconv(int(a.Reg))
		}
		if a.Sym != nil {
			str = fmt.Sprintf("%s%s(%s)", a.Sym.Name, offConv(a.Offset), reg)
		} else {
			str = fmt.Sprintf("%s(%s)", offConv(a.Offset), reg)
		}

	case NAME_PARAM:
		reg := "FP"
		if a.Reg != REG_NONE {
			reg = psess.Rconv(int(a.Reg))
		}
		if a.Sym != nil {
			str = fmt.Sprintf("%s%s(%s)", a.Sym.Name, offConv(a.Offset), reg)
		} else {
			str = fmt.Sprintf("%s(%s)", offConv(a.Offset), reg)
		}
	}
	return str
}

func offConv(off int64) string {
	if off == 0 {
		return ""
	}
	return fmt.Sprintf("%+d", off)
}

// opSuffixSet is like regListSet, but for opcode suffixes.
//
// Unlike some other similar structures, uint8 space is not
// divided by it's own values set (because the're only 256 of them).
// Instead, every arch may interpret/format all 8 bits as they like,
// as long as they register proper cconv function for it.
type opSuffixSet struct {
	arch  string
	cconv func(suffix uint8) string
}

// RegisterOpSuffix assigns cconv function for formatting opcode suffixes
// when compiling for GOARCH=arch.
//
// cconv is never called with 0 argument.
func (psess *PackageSession) RegisterOpSuffix(arch string, cconv func(uint8) string) {
	psess.
		opSuffixSpace = append(psess.opSuffixSpace, opSuffixSet{
		arch:  arch,
		cconv: cconv,
	})
}

type regSet struct {
	lo    int
	hi    int
	Rconv func(int) string
}

// Few enough architectures that a linear scan is fastest.
// Not even worth sorting.

const (
	// Because of masking operations in the encodings, each register
	// space should start at 0 modulo some power of 2.
	RBase386   = 1 * 1024
	RBaseAMD64 = 2 * 1024
	RBaseARM   = 3 * 1024
	RBasePPC64 = 4 * 1024  // range [4k, 8k)
	RBaseARM64 = 8 * 1024  // range [8k, 13k)
	RBaseMIPS  = 13 * 1024 // range [13k, 14k)
	RBaseS390X = 14 * 1024 // range [14k, 15k)
	RBaseWasm  = 16 * 1024
)

// RegisterRegister binds a pretty-printer (Rconv) for register
// numbers to a given register number range. Lo is inclusive,
// hi exclusive (valid registers are lo through hi-1).
func (psess *PackageSession) RegisterRegister(lo, hi int, Rconv func(int) string) {
	psess.
		regSpace = append(psess.regSpace, regSet{lo, hi, Rconv})
}

func (psess *PackageSession) Rconv(reg int) string {
	if reg == REG_NONE {
		return "NONE"
	}
	for i := range psess.regSpace {
		rs := &psess.regSpace[i]
		if rs.lo <= reg && reg < rs.hi {
			return rs.Rconv(reg)
		}
	}
	return fmt.Sprintf("R???%d", reg)
}

type regListSet struct {
	lo     int64
	hi     int64
	RLconv func(int64) string
}

// Each architecture is allotted a distinct subspace: [Lo, Hi) for declaring its
// arch-specific register list numbers.
const (
	RegListARMLo = 0
	RegListARMHi = 1 << 16

	// arm64 uses the 60th bit to differentiate from other archs
	RegListARM64Lo = 1 << 60
	RegListARM64Hi = 1<<61 - 1

	// x86 uses the 61th bit to differentiate from other archs
	RegListX86Lo = 1 << 61
	RegListX86Hi = 1<<62 - 1
)

// RegisterRegisterList binds a pretty-printer (RLconv) for register list
// numbers to a given register list number range. Lo is inclusive,
// hi exclusive (valid register list are lo through hi-1).
func (psess *PackageSession) RegisterRegisterList(lo, hi int64, rlconv func(int64) string) {
	psess.
		regListSpace = append(psess.regListSpace, regListSet{lo, hi, rlconv})
}

func (psess *PackageSession) RLconv(list int64) string {
	for i := range psess.regListSpace {
		rls := &psess.regListSpace[i]
		if rls.lo <= list && list < rls.hi {
			return rls.RLconv(list)
		}
	}
	return fmt.Sprintf("RL???%d", list)
}

type opSet struct {
	lo    As
	names []string
}

// Not even worth sorting

// RegisterOpcode binds a list of instruction names
// to a given instruction number range.
func (psess *PackageSession) RegisterOpcode(lo As, Anames []string) {
	if len(Anames) > AllowedOpCodes {
		panic(fmt.Sprintf("too many instructions, have %d max %d", len(Anames), AllowedOpCodes))
	}
	psess.
		aSpace = append(psess.aSpace, opSet{lo, Anames})
}

func (a As) String(psess *PackageSession) string {
	if 0 <= a && int(a) < len(psess.Anames) {
		return psess.Anames[a]
	}
	for i := range psess.aSpace {
		as := &psess.aSpace[i]
		if as.lo <= a && int(a-as.lo) < len(as.names) {
			return as.names[a-as.lo]
		}
	}
	return fmt.Sprintf("A???%d", a)
}

func Bool2int(b bool) int {
	// The compiler currently only optimizes this form.
	// See issue 6011.
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

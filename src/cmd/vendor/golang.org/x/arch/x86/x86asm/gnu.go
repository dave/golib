package x86asm

import (
	"fmt"
	"strings"
)

// GNUSyntax returns the GNU assembler syntax for the instruction, as defined by GNU binutils.
// This general form is often called ``AT&T syntax'' as a reference to AT&T System V Unix.
func GNUSyntax(inst Inst, pc uint64, symname SymLookup) string {

	if symname == nil {
		symname = func(uint64) (string, uint64) { return "", 0 }
	}

	switch inst.Op {
	case FDIV, FDIVR, FSUB, FSUBR, FDIVP, FDIVRP, FSUBP, FSUBRP:

		_, reg1 := inst.Args[0].(Reg)
		_, reg2 := inst.Args[1].(Reg)
		if reg1 && reg2 && (inst.Opcode>>24 == 0xDC || inst.Opcode>>24 == 0xDE) {
			switch inst.Op {
			case FDIV:
				inst.Op = FDIVR
			case FDIVR:
				inst.Op = FDIV
			case FSUB:
				inst.Op = FSUBR
			case FSUBR:
				inst.Op = FSUB
			case FDIVP:
				inst.Op = FDIVRP
			case FDIVRP:
				inst.Op = FDIVP
			case FSUBP:
				inst.Op = FSUBRP
			case FSUBRP:
				inst.Op = FSUBP
			}
		}

	case MOVNTSD:

		if countPrefix(&inst, 0xF3) > 0 {
			found := false
			for i := len(inst.Prefix) - 1; i >= 0; i-- {
				switch inst.Prefix[i] & 0xFF {
				case 0xF3:
					if !found {
						found = true
						inst.Prefix[i] |= PrefixImplicit
					}
				case 0xF2:
					inst.Prefix[i] &^= PrefixImplicit
				}
			}
			inst.Op = MOVNTSS
		}
	}

	switch inst.Op {
	case MONITOR:
		inst.Args[0] = EDX
		inst.Args[1] = ECX
		inst.Args[2] = EAX
		if inst.AddrSize == 16 {
			inst.Args[2] = AX
		}

	case MWAIT:
		if inst.Mode == 64 {
			inst.Args[0] = RCX
			inst.Args[1] = RAX
		} else {
			inst.Args[0] = ECX
			inst.Args[1] = EAX
		}
	}

	switch inst.Op {
	case CRC32:

		if countPrefix(&inst, 0xF2) > 1 {
			unmarkImplicit(&inst, 0xF2)
			markLastImplicit(&inst, 0xF2)
		}

		unmarkImplicit(&inst, PrefixDataSize)
		markLastImplicit(&inst, PrefixDataSize)

	case CVTSI2SD, CVTSI2SS:
		if !isMem(inst.Args[1]) {
			markLastImplicit(&inst, PrefixDataSize)
		}

	case CVTSD2SI, CVTSS2SI, CVTTSD2SI, CVTTSS2SI,
		ENTER, FLDENV, FNSAVE, FNSTENV, FRSTOR, LGDT, LIDT, LRET,
		POP, PUSH, RET, SGDT, SIDT, SYSRET, XBEGIN:
		markLastImplicit(&inst, PrefixDataSize)

	case LOOP, LOOPE, LOOPNE, MONITOR:
		markLastImplicit(&inst, PrefixAddrSize)

	case MOV:

		dst, _ := inst.Args[0].(Reg)
		src, _ := inst.Args[1].(Reg)
		if ES <= src && src <= GS && isMem(inst.Args[0]) || ES <= dst && dst <= GS && isMem(inst.Args[1]) {
			unmarkImplicit(&inst, PrefixDataSize)
		}

	case MOVDQU:
		if countPrefix(&inst, 0xF3) > 1 {
			unmarkImplicit(&inst, 0xF3)
			markLastImplicit(&inst, 0xF3)
		}

	case MOVQ2DQ:
		markLastImplicit(&inst, PrefixDataSize)

	case SLDT, SMSW, STR, FXRSTOR, XRSTOR, XSAVE, XSAVEOPT, CMPXCHG8B:
		if isMem(inst.Args[0]) {
			unmarkImplicit(&inst, PrefixDataSize)
		}

	case SYSEXIT:
		unmarkImplicit(&inst, PrefixDataSize)
	}

	if isCondJmp[inst.Op] || isLoop[inst.Op] || inst.Op == JCXZ || inst.Op == JECXZ || inst.Op == JRCXZ {
		if countPrefix(&inst, PrefixCS) > 0 && countPrefix(&inst, PrefixDS) > 0 {
			for i, p := range inst.Prefix {
				switch p & 0xFFF {
				case PrefixPN, PrefixPT:
					inst.Prefix[i] &= 0xF0FF
				}
			}
		}
	}

	if inst.Op == MOV {

		haveREPN := false
		for i := len(inst.Prefix) - 1; i >= 0; i-- {
			switch inst.Prefix[i] &^ PrefixIgnored {
			case PrefixREPN:
				haveREPN = true
			case PrefixXRELEASE:
				if haveREPN {
					inst.Prefix[i] = PrefixREP
				}
			}
		}
	}

	haveXA := false
	haveXR := false
	for i := len(inst.Prefix) - 1; i >= 0; i-- {
		switch inst.Prefix[i] &^ PrefixIgnored {
		case PrefixXRELEASE:
			if !haveXR {
				haveXR = true
			} else {
				inst.Prefix[i] = PrefixREP
			}

		case PrefixXACQUIRE:
			if !haveXA {
				haveXA = true
			} else {
				inst.Prefix[i] = PrefixREPN
			}
		}
	}

	op := strings.ToLower(inst.Op.String())
	if alt := gnuOp[inst.Op]; alt != "" {
		op = alt
	}

	needSuffix := true
SuffixLoop:
	for i, a := range inst.Args {
		if a == nil {
			break
		}
		switch a := a.(type) {
		case Reg:
			switch inst.Op {
			case MOVSX, MOVZX:
				continue

			case SHL, SHR, RCL, RCR, ROL, ROR, SAR:
				if i == 1 {

					continue
				}

			case CRC32:

				continue

			case PUSH, POP:

				if ES <= a && a <= GS {
					continue
				}

			case CVTSI2SD, CVTSI2SS:

				if X0 <= a && a <= X15 {
					continue
				}
			}

			if AL <= a && a <= R15 || ES <= a && a <= GS || X0 <= a && a <= X15 || M0 <= a && a <= M7 {
				needSuffix = false
				break SuffixLoop
			}
		}
	}

	if needSuffix {
		switch inst.Op {
		case CMPXCHG8B, FLDCW, FNSTCW, FNSTSW, LDMXCSR, LLDT, LMSW, LTR, PCLMULQDQ,
			SETA, SETAE, SETB, SETBE, SETE, SETG, SETGE, SETL, SETLE, SETNE, SETNO, SETNP, SETNS, SETO, SETP, SETS,
			SLDT, SMSW, STMXCSR, STR, VERR, VERW:

		case CRC32:
			op += byteSizeSuffix(argBytes(&inst, inst.Args[1]))

		case LGDT, LIDT, SGDT, SIDT:
			op += byteSizeSuffix(inst.DataSize / 8)

		case MOVZX, MOVSX:

			op = op[:4] + byteSizeSuffix(argBytes(&inst, inst.Args[1])) + byteSizeSuffix(argBytes(&inst, inst.Args[0]))

		case LOOP, LOOPE, LOOPNE:

			if inst.AddrSize == 16 {
				op += "w"
			}

		case CALL, ENTER, JMP, LCALL, LEAVE, LJMP, LRET, RET, SYSRET, XBEGIN:

			if inst.Opcode>>24 == 0xEB {
				break
			}
			if inst.DataSize == 16 && inst.Mode != 16 {
				markLastImplicit(&inst, PrefixDataSize)
				op += "w"
			} else if inst.Mode == 64 {
				op += "q"
			}

		case FRSTOR, FNSAVE, FNSTENV, FLDENV:

			if inst.DataSize == 16 {
				op += "s"
			}

		case PUSH, POP:
			if markLastImplicit(&inst, PrefixDataSize) {
				op += byteSizeSuffix(inst.DataSize / 8)
			} else if inst.Mode == 64 {
				op += "q"
			} else {
				op += byteSizeSuffix(inst.MemBytes)
			}

		default:
			if isFloat(inst.Op) {

				switch inst.MemBytes {
				default:
					if (inst.Op == FLD || inst.Op == FSTP) && isMem(inst.Args[0]) {
						op += "t"
					}
				case 4:
					if isFloatInt(inst.Op) {
						op += "l"
					} else {
						op += "s"
					}
				case 8:
					if isFloatInt(inst.Op) {
						op += "ll"
					} else {
						op += "l"
					}
				}
				break
			}

			op += byteSizeSuffix(inst.MemBytes)
		}
	}

	switch inst.Op {
	case 0:
		if inst.Prefix[0] != 0 {
			return strings.ToLower(inst.Prefix[0].String())
		}

	case INT:
		if inst.Opcode>>24 == 0xCC {
			inst.Args[0] = nil
			op = "int3"
		}

	case CMPPS, CMPPD, CMPSD_XMM, CMPSS:
		imm, ok := inst.Args[2].(Imm)
		if ok && 0 <= imm && imm < 8 {
			inst.Args[2] = nil
			op = cmppsOps[imm] + op[3:]
		}

	case PCLMULQDQ:
		imm, ok := inst.Args[2].(Imm)
		if ok && imm&^0x11 == 0 {
			inst.Args[2] = nil
			op = pclmulqOps[(imm&0x10)>>3|(imm&1)]
		}

	case XLATB:
		if markLastImplicit(&inst, PrefixAddrSize) {
			op = "xlat"
		}
	}

	// Build list of argument strings.
	var (
		usedPrefixes bool     // segment prefixes consumed by Mem formatting
		args         []string // formatted arguments
	)
	for i, a := range inst.Args {
		if a == nil {
			break
		}
		switch inst.Op {
		case MOVSB, MOVSW, MOVSD, MOVSQ, OUTSB, OUTSW, OUTSD:
			if i == 0 {
				usedPrefixes = true
			} else {
				usedPrefixes = false
			}
		}
		if a == Imm(1) && (inst.Opcode>>24)&^1 == 0xD0 {
			continue
		}
		args = append(args, gnuArg(&inst, pc, symname, a, &usedPrefixes))
	}

	switch inst.Op {
	case BOUND, LCALL, ENTER, LJMP:

	default:

		for i, j := 0, len(args)-1; i < j; i, j = i+1, j-1 {
			args[i], args[j] = args[j], args[i]
		}
	}

	// Build prefix string.
	// Must be after argument formatting, which can turn off segment prefixes.
	var (
		prefix       = "" // output string
		numAddr      = 0
		numData      = 0
		implicitData = false
	)
	for _, p := range inst.Prefix {
		if p&0xFF == PrefixDataSize && p&PrefixImplicit != 0 {
			implicitData = true
		}
	}
	for _, p := range inst.Prefix {
		if p == 0 || p.IsVEX() {
			break
		}
		if p&PrefixImplicit != 0 {
			continue
		}
		switch p &^ (PrefixIgnored | PrefixInvalid) {
		default:
			if p.IsREX() {
				if p&0xFF == PrefixREX {
					prefix += "rex "
				} else {
					prefix += "rex." + p.String()[4:] + " "
				}
				break
			}
			prefix += strings.ToLower(p.String()) + " "

		case PrefixPN:
			op += ",pn"
			continue

		case PrefixPT:
			op += ",pt"
			continue

		case PrefixAddrSize, PrefixAddr16, PrefixAddr32:

			n := 32
			if inst.Mode == 32 {
				n = 16
			}
			numAddr++
			if countPrefix(&inst, PrefixAddrSize) > numAddr {
				n = inst.Mode
			}
			prefix += fmt.Sprintf("addr%d ", n)
			continue

		case PrefixData16, PrefixData32:
			if implicitData && countPrefix(&inst, PrefixDataSize) > 1 {

				n := 16
				if inst.Mode == 16 {
					n = 32
				}
				numData++
				if countPrefix(&inst, PrefixDataSize) > numData {
					if inst.Mode == 16 {
						n = 16
					} else {
						n = 32
					}
				}
				prefix += fmt.Sprintf("data%d ", n)
				continue
			}
			prefix += strings.ToLower(p.String()) + " "
		}
	}

	text := prefix + op
	if args != nil {
		text += " "

		if (inst.Op == CALL || inst.Op == JMP || inst.Op == LJMP || inst.Op == LCALL) && (isMem(inst.Args[0]) || isReg(inst.Args[0])) {
			text += "*"
		}
		text += strings.Join(args, ",")
	}
	return text
}

// gnuArg returns the GNU syntax for the argument x from the instruction inst.
// If *usedPrefixes is false and x is a Mem, then the formatting
// includes any segment prefixes and sets *usedPrefixes to true.
func gnuArg(inst *Inst, pc uint64, symname SymLookup, x Arg, usedPrefixes *bool) string {
	if x == nil {
		return "<nil>"
	}
	switch x := x.(type) {
	case Reg:
		switch inst.Op {
		case CVTSI2SS, CVTSI2SD, CVTSS2SI, CVTSD2SI, CVTTSD2SI, CVTTSS2SI:
			if inst.DataSize == 16 && EAX <= x && x <= R15L {
				x -= EAX - AX
			}

		case IN, INSB, INSW, INSD, OUT, OUTSB, OUTSW, OUTSD:

			if x == DX {
				return "(%dx)"
			}
		case VMOVDQA, VMOVDQU, VMOVNTDQA, VMOVNTDQ:
			return strings.Replace(gccRegName[x], "xmm", "ymm", -1)
		}
		return gccRegName[x]
	case Mem:
		if s, disp := memArgToSymbol(x, pc, inst.Len, symname); s != "" {
			suffix := ""
			if disp != 0 {
				suffix = fmt.Sprintf("%+d", disp)
			}
			return fmt.Sprintf("%s%s", s, suffix)
		}
		seg := ""
		var haveCS, haveDS, haveES, haveFS, haveGS, haveSS bool
		switch x.Segment {
		case CS:
			haveCS = true
		case DS:
			haveDS = true
		case ES:
			haveES = true
		case FS:
			haveFS = true
		case GS:
			haveGS = true
		case SS:
			haveSS = true
		}
		switch inst.Op {
		case INSB, INSW, INSD, STOSB, STOSW, STOSD, STOSQ, SCASB, SCASW, SCASD, SCASQ:

		default:
			if *usedPrefixes {
				break
			}
			for i := len(inst.Prefix) - 1; i >= 0; i-- {
				p := inst.Prefix[i] &^ PrefixIgnored
				if p == 0 {
					continue
				}
				switch p {
				case PrefixCS:
					if !haveCS {
						haveCS = true
						inst.Prefix[i] |= PrefixImplicit
					}
				case PrefixDS:
					if !haveDS {
						haveDS = true
						inst.Prefix[i] |= PrefixImplicit
					}
				case PrefixES:
					if !haveES {
						haveES = true
						inst.Prefix[i] |= PrefixImplicit
					}
				case PrefixFS:
					if !haveFS {
						haveFS = true
						inst.Prefix[i] |= PrefixImplicit
					}
				case PrefixGS:
					if !haveGS {
						haveGS = true
						inst.Prefix[i] |= PrefixImplicit
					}
				case PrefixSS:
					if !haveSS {
						haveSS = true
						inst.Prefix[i] |= PrefixImplicit
					}
				}
			}
			*usedPrefixes = true
		}
		if haveCS {
			seg += "%cs:"
		}
		if haveDS {
			seg += "%ds:"
		}
		if haveSS {
			seg += "%ss:"
		}
		if haveES {
			seg += "%es:"
		}
		if haveFS {
			seg += "%fs:"
		}
		if haveGS {
			seg += "%gs:"
		}
		disp := ""
		if x.Disp != 0 {
			disp = fmt.Sprintf("%#x", x.Disp)
		}
		if x.Scale == 0 || x.Index == 0 && x.Scale == 1 && (x.Base == ESP || x.Base == RSP || x.Base == 0 && inst.Mode == 64) {
			if x.Base == 0 {
				return seg + disp
			}
			return fmt.Sprintf("%s%s(%s)", seg, disp, gccRegName[x.Base])
		}
		base := gccRegName[x.Base]
		if x.Base == 0 {
			base = ""
		}
		index := gccRegName[x.Index]
		if x.Index == 0 {
			if inst.AddrSize == 64 {
				index = "%riz"
			} else {
				index = "%eiz"
			}
		}
		if AX <= x.Base && x.Base <= DI {

			return fmt.Sprintf("%s%s(%s,%s)", seg, disp, base, index)
		}
		return fmt.Sprintf("%s%s(%s,%s,%d)", seg, disp, base, index, x.Scale)
	case Rel:
		if pc == 0 {
			return fmt.Sprintf(".%+#x", int64(x))
		} else {
			addr := pc + uint64(inst.Len) + uint64(x)
			if s, base := symname(addr); s != "" && addr == base {
				return fmt.Sprintf("%s", s)
			} else {
				addr := pc + uint64(inst.Len) + uint64(x)
				return fmt.Sprintf("%#x", addr)
			}
		}
	case Imm:
		if s, base := symname(uint64(x)); s != "" {
			suffix := ""
			if uint64(x) != base {
				suffix = fmt.Sprintf("%+d", uint64(x)-base)
			}
			return fmt.Sprintf("$%s%s", s, suffix)
		}
		if inst.Mode == 32 {
			return fmt.Sprintf("$%#x", uint32(x))
		}
		return fmt.Sprintf("$%#x", int64(x))
	}
	return x.String()
}

var gccRegName = [...]string{
	0:    "REG0",
	AL:   "%al",
	CL:   "%cl",
	BL:   "%bl",
	DL:   "%dl",
	AH:   "%ah",
	CH:   "%ch",
	BH:   "%bh",
	DH:   "%dh",
	SPB:  "%spl",
	BPB:  "%bpl",
	SIB:  "%sil",
	DIB:  "%dil",
	R8B:  "%r8b",
	R9B:  "%r9b",
	R10B: "%r10b",
	R11B: "%r11b",
	R12B: "%r12b",
	R13B: "%r13b",
	R14B: "%r14b",
	R15B: "%r15b",
	AX:   "%ax",
	CX:   "%cx",
	BX:   "%bx",
	DX:   "%dx",
	SP:   "%sp",
	BP:   "%bp",
	SI:   "%si",
	DI:   "%di",
	R8W:  "%r8w",
	R9W:  "%r9w",
	R10W: "%r10w",
	R11W: "%r11w",
	R12W: "%r12w",
	R13W: "%r13w",
	R14W: "%r14w",
	R15W: "%r15w",
	EAX:  "%eax",
	ECX:  "%ecx",
	EDX:  "%edx",
	EBX:  "%ebx",
	ESP:  "%esp",
	EBP:  "%ebp",
	ESI:  "%esi",
	EDI:  "%edi",
	R8L:  "%r8d",
	R9L:  "%r9d",
	R10L: "%r10d",
	R11L: "%r11d",
	R12L: "%r12d",
	R13L: "%r13d",
	R14L: "%r14d",
	R15L: "%r15d",
	RAX:  "%rax",
	RCX:  "%rcx",
	RDX:  "%rdx",
	RBX:  "%rbx",
	RSP:  "%rsp",
	RBP:  "%rbp",
	RSI:  "%rsi",
	RDI:  "%rdi",
	R8:   "%r8",
	R9:   "%r9",
	R10:  "%r10",
	R11:  "%r11",
	R12:  "%r12",
	R13:  "%r13",
	R14:  "%r14",
	R15:  "%r15",
	IP:   "%ip",
	EIP:  "%eip",
	RIP:  "%rip",
	F0:   "%st",
	F1:   "%st(1)",
	F2:   "%st(2)",
	F3:   "%st(3)",
	F4:   "%st(4)",
	F5:   "%st(5)",
	F6:   "%st(6)",
	F7:   "%st(7)",
	M0:   "%mm0",
	M1:   "%mm1",
	M2:   "%mm2",
	M3:   "%mm3",
	M4:   "%mm4",
	M5:   "%mm5",
	M6:   "%mm6",
	M7:   "%mm7",
	X0:   "%xmm0",
	X1:   "%xmm1",
	X2:   "%xmm2",
	X3:   "%xmm3",
	X4:   "%xmm4",
	X5:   "%xmm5",
	X6:   "%xmm6",
	X7:   "%xmm7",
	X8:   "%xmm8",
	X9:   "%xmm9",
	X10:  "%xmm10",
	X11:  "%xmm11",
	X12:  "%xmm12",
	X13:  "%xmm13",
	X14:  "%xmm14",
	X15:  "%xmm15",
	CS:   "%cs",
	SS:   "%ss",
	DS:   "%ds",
	ES:   "%es",
	FS:   "%fs",
	GS:   "%gs",
	GDTR: "%gdtr",
	IDTR: "%idtr",
	LDTR: "%ldtr",
	MSW:  "%msw",
	TASK: "%task",
	CR0:  "%cr0",
	CR1:  "%cr1",
	CR2:  "%cr2",
	CR3:  "%cr3",
	CR4:  "%cr4",
	CR5:  "%cr5",
	CR6:  "%cr6",
	CR7:  "%cr7",
	CR8:  "%cr8",
	CR9:  "%cr9",
	CR10: "%cr10",
	CR11: "%cr11",
	CR12: "%cr12",
	CR13: "%cr13",
	CR14: "%cr14",
	CR15: "%cr15",
	DR0:  "%db0",
	DR1:  "%db1",
	DR2:  "%db2",
	DR3:  "%db3",
	DR4:  "%db4",
	DR5:  "%db5",
	DR6:  "%db6",
	DR7:  "%db7",
	TR0:  "%tr0",
	TR1:  "%tr1",
	TR2:  "%tr2",
	TR3:  "%tr3",
	TR4:  "%tr4",
	TR5:  "%tr5",
	TR6:  "%tr6",
	TR7:  "%tr7",
}

var gnuOp = map[Op]string{
	CBW:       "cbtw",
	CDQ:       "cltd",
	CMPSD:     "cmpsl",
	CMPSD_XMM: "cmpsd",
	CWD:       "cwtd",
	CWDE:      "cwtl",
	CQO:       "cqto",
	INSD:      "insl",
	IRET:      "iretw",
	IRETD:     "iret",
	IRETQ:     "iretq",
	LODSB:     "lods",
	LODSD:     "lods",
	LODSQ:     "lods",
	LODSW:     "lods",
	MOVSD:     "movsl",
	MOVSD_XMM: "movsd",
	OUTSD:     "outsl",
	POPA:      "popaw",
	POPAD:     "popa",
	POPF:      "popfw",
	POPFD:     "popf",
	PUSHA:     "pushaw",
	PUSHAD:    "pusha",
	PUSHF:     "pushfw",
	PUSHFD:    "pushf",
	SCASB:     "scas",
	SCASD:     "scas",
	SCASQ:     "scas",
	SCASW:     "scas",
	STOSB:     "stos",
	STOSD:     "stos",
	STOSQ:     "stos",
	STOSW:     "stos",
	XLATB:     "xlat",
}

var cmppsOps = []string{
	"cmpeq",
	"cmplt",
	"cmple",
	"cmpunord",
	"cmpneq",
	"cmpnlt",
	"cmpnle",
	"cmpord",
}

var pclmulqOps = []string{
	"pclmullqlqdq",
	"pclmulhqlqdq",
	"pclmullqhqdq",
	"pclmulhqhqdq",
}

func countPrefix(inst *Inst, target Prefix) int {
	n := 0
	for _, p := range inst.Prefix {
		if p&0xFF == target&0xFF {
			n++
		}
	}
	return n
}

func markLastImplicit(inst *Inst, prefix Prefix) bool {
	for i := len(inst.Prefix) - 1; i >= 0; i-- {
		p := inst.Prefix[i]
		if p&0xFF == prefix {
			inst.Prefix[i] |= PrefixImplicit
			return true
		}
	}
	return false
}

func unmarkImplicit(inst *Inst, prefix Prefix) {
	for i := len(inst.Prefix) - 1; i >= 0; i-- {
		p := inst.Prefix[i]
		if p&0xFF == prefix {
			inst.Prefix[i] &^= PrefixImplicit
		}
	}
}

func byteSizeSuffix(b int) string {
	switch b {
	case 1:
		return "b"
	case 2:
		return "w"
	case 4:
		return "l"
	case 8:
		return "q"
	}
	return ""
}

func argBytes(inst *Inst, arg Arg) int {
	if isMem(arg) {
		return inst.MemBytes
	}
	return regBytes(arg)
}

func isFloat(op Op) bool {
	switch op {
	case FADD, FCOM, FCOMP, FDIV, FDIVR, FIADD, FICOM, FICOMP, FIDIV, FIDIVR, FILD, FIMUL, FIST, FISTP, FISTTP, FISUB, FISUBR, FLD, FMUL, FST, FSTP, FSUB, FSUBR:
		return true
	}
	return false
}

func isFloatInt(op Op) bool {
	switch op {
	case FIADD, FICOM, FICOMP, FIDIV, FIDIVR, FILD, FIMUL, FIST, FISTP, FISTTP, FISUB, FISUBR:
		return true
	}
	return false
}

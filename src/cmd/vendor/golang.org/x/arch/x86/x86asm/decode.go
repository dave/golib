package x86asm

import (
	"encoding/binary"
	"errors"
	"fmt"
	"runtime"
)

// Set trace to true to cause the decoder to print the PC sequence
// of the executed instruction codes. This is typically only useful
// when you are running a test of a single input case.
const trace = false

type decodeOp uint16

const (
	xFail  decodeOp = iota // invalid instruction (return)
	xMatch                 // completed match
	xJump                  // jump to pc

	xCondByte     // switch on instruction byte value
	xCondSlashR   // read and switch on instruction /r value
	xCondPrefix   // switch on presence of instruction prefix
	xCondIs64     // switch on 64-bit processor mode
	xCondDataSize // switch on operand size
	xCondAddrSize // switch on address size
	xCondIsMem    // switch on memory vs register argument

	xSetOp // set instruction opcode

	xReadSlashR // read /r
	xReadIb     // read ib
	xReadIw     // read iw
	xReadId     // read id
	xReadIo     // read io
	xReadCb     // read cb
	xReadCw     // read cw
	xReadCd     // read cd
	xReadCp     // read cp
	xReadCm     // read cm

	xArg1            // arg 1
	xArg3            // arg 3
	xArgAL           // arg AL
	xArgAX           // arg AX
	xArgCL           // arg CL
	xArgCR0dashCR7   // arg CR0-CR7
	xArgCS           // arg CS
	xArgDR0dashDR7   // arg DR0-DR7
	xArgDS           // arg DS
	xArgDX           // arg DX
	xArgEAX          // arg EAX
	xArgEDX          // arg EDX
	xArgES           // arg ES
	xArgFS           // arg FS
	xArgGS           // arg GS
	xArgImm16        // arg imm16
	xArgImm32        // arg imm32
	xArgImm64        // arg imm64
	xArgImm8         // arg imm8
	xArgImm8u        // arg imm8 but record as unsigned
	xArgImm16u       // arg imm8 but record as unsigned
	xArgM            // arg m
	xArgM128         // arg m128
	xArgM256         // arg m256
	xArgM1428byte    // arg m14/28byte
	xArgM16          // arg m16
	xArgM16and16     // arg m16&16
	xArgM16and32     // arg m16&32
	xArgM16and64     // arg m16&64
	xArgM16colon16   // arg m16:16
	xArgM16colon32   // arg m16:32
	xArgM16colon64   // arg m16:64
	xArgM16int       // arg m16int
	xArgM2byte       // arg m2byte
	xArgM32          // arg m32
	xArgM32and32     // arg m32&32
	xArgM32fp        // arg m32fp
	xArgM32int       // arg m32int
	xArgM512byte     // arg m512byte
	xArgM64          // arg m64
	xArgM64fp        // arg m64fp
	xArgM64int       // arg m64int
	xArgM8           // arg m8
	xArgM80bcd       // arg m80bcd
	xArgM80dec       // arg m80dec
	xArgM80fp        // arg m80fp
	xArgM94108byte   // arg m94/108byte
	xArgMm           // arg mm
	xArgMm1          // arg mm1
	xArgMm2          // arg mm2
	xArgMm2M64       // arg mm2/m64
	xArgMmM32        // arg mm/m32
	xArgMmM64        // arg mm/m64
	xArgMem          // arg mem
	xArgMoffs16      // arg moffs16
	xArgMoffs32      // arg moffs32
	xArgMoffs64      // arg moffs64
	xArgMoffs8       // arg moffs8
	xArgPtr16colon16 // arg ptr16:16
	xArgPtr16colon32 // arg ptr16:32
	xArgR16          // arg r16
	xArgR16op        // arg r16 with +rw in opcode
	xArgR32          // arg r32
	xArgR32M16       // arg r32/m16
	xArgR32M8        // arg r32/m8
	xArgR32op        // arg r32 with +rd in opcode
	xArgR64          // arg r64
	xArgR64M16       // arg r64/m16
	xArgR64op        // arg r64 with +rd in opcode
	xArgR8           // arg r8
	xArgR8op         // arg r8 with +rb in opcode
	xArgRAX          // arg RAX
	xArgRDX          // arg RDX
	xArgRM           // arg r/m
	xArgRM16         // arg r/m16
	xArgRM32         // arg r/m32
	xArgRM64         // arg r/m64
	xArgRM8          // arg r/m8
	xArgReg          // arg reg
	xArgRegM16       // arg reg/m16
	xArgRegM32       // arg reg/m32
	xArgRegM8        // arg reg/m8
	xArgRel16        // arg rel16
	xArgRel32        // arg rel32
	xArgRel8         // arg rel8
	xArgSS           // arg SS
	xArgST           // arg ST, aka ST(0)
	xArgSTi          // arg ST(i) with +i in opcode
	xArgSreg         // arg Sreg
	xArgTR0dashTR7   // arg TR0-TR7
	xArgXmm          // arg xmm
	xArgXMM0         // arg <XMM0>
	xArgXmm1         // arg xmm1
	xArgXmm2         // arg xmm2
	xArgXmm2M128     // arg xmm2/m128
	xArgYmm2M256     // arg ymm2/m256
	xArgXmm2M16      // arg xmm2/m16
	xArgXmm2M32      // arg xmm2/m32
	xArgXmm2M64      // arg xmm2/m64
	xArgXmmM128      // arg xmm/m128
	xArgXmmM32       // arg xmm/m32
	xArgXmmM64       // arg xmm/m64
	xArgYmm1         // arg ymm1
	xArgRmf16        // arg r/m16 but force mod=3
	xArgRmf32        // arg r/m32 but force mod=3
	xArgRmf64        // arg r/m64 but force mod=3
)

// instPrefix returns an Inst describing just one prefix byte.
// It is only used if there is a prefix followed by an unintelligible
// or invalid instruction byte sequence.
func instPrefix(b byte, mode int) (Inst, error) {

	if trace {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line)
	}
	p := Prefix(b)
	switch p {
	case PrefixDataSize:
		if mode == 16 {
			p = PrefixData32
		} else {
			p = PrefixData16
		}
	case PrefixAddrSize:
		if mode == 32 {
			p = PrefixAddr16
		} else {
			p = PrefixAddr32
		}
	}

	inst := Inst{Len: 1}
	inst.Prefix = Prefixes{p}
	return inst, nil
}

// truncated reports a truncated instruction.
// For now we use instPrefix but perhaps later we will return
// a specific error here.
func truncated(src []byte, mode int) (Inst, error) {

	return instPrefix(src[0], mode)
}

// These are the errors returned by Decode.
var (
	ErrInvalidMode  = errors.New("invalid x86 mode in Decode")
	ErrTruncated    = errors.New("truncated instruction")
	ErrUnrecognized = errors.New("unrecognized instruction")
)

// decoderCover records coverage information for which parts
// of the byte code have been executed.
// TODO(rsc): This is for testing. Only use this if a flag is given.
var decoderCover []bool

// Decode decodes the leading bytes in src as a single instruction.
// The mode arguments specifies the assumed processor mode:
// 16, 32, or 64 for 16-, 32-, and 64-bit execution modes.
func Decode(src []byte, mode int) (inst Inst, err error) {
	return decode1(src, mode, false)
}

// decode1 is the implementation of Decode but takes an extra
// gnuCompat flag to cause it to change its behavior to mimic
// bugs (or at least unique features) of GNU libopcodes as used
// by objdump. We don't believe that logic is the right thing to do
// in general, but when testing against libopcodes it simplifies the
// comparison if we adjust a few small pieces of logic.
// The affected logic is in the conditional branch for "mandatory" prefixes,
// case xCondPrefix.
func decode1(src []byte, mode int, gnuCompat bool) (Inst, error) {
	switch mode {
	case 16, 32, 64:

	default:
		return Inst{}, ErrInvalidMode
	}

	if len(src) > 15 {
		src = src[:15]
	}

	var (
		// prefix decoding information
		pos           = 0    // position reading src
		nprefix       = 0    // number of prefixes
		lockIndex     = -1   // index of LOCK prefix in src and inst.Prefix
		repIndex      = -1   // index of REP/REPN prefix in src and inst.Prefix
		segIndex      = -1   // index of Group 2 prefix in src and inst.Prefix
		dataSizeIndex = -1   // index of Group 3 prefix in src and inst.Prefix
		addrSizeIndex = -1   // index of Group 4 prefix in src and inst.Prefix
		rex           Prefix // rex byte if present (or 0)
		rexUsed       Prefix // bits used in rex byte
		rexIndex      = -1   // index of rex byte
		vex           Prefix // use vex encoding
		vexIndex      = -1   // index of vex prefix

		addrMode = mode // address mode (width in bits)
		dataMode = mode // operand mode (width in bits)

		// decoded ModR/M fields
		haveModrm bool
		modrm     int
		mod       int
		regop     int
		rm        int

		// if ModR/M is memory reference, Mem form
		mem     Mem
		haveMem bool

		// decoded SIB fields
		haveSIB bool
		sib     int
		scale   int
		index   int
		base    int
		displen int
		dispoff int

		// decoded immediate values
		imm     int64
		imm8    int8
		immc    int64
		immcpos int

		// output
		opshift int
		inst    Inst
		narg    int // number of arguments written to inst
	)

	if mode == 64 {
		dataMode = 32
	}

ReadPrefixes:
	for ; pos < len(src); pos++ {
		p := Prefix(src[pos])
		switch p {
		default:
			nprefix = pos
			break ReadPrefixes

		case 0xF0:
			if lockIndex >= 0 {
				inst.Prefix[lockIndex] |= PrefixIgnored
			}
			lockIndex = pos
		case 0xF2, 0xF3:
			if repIndex >= 0 {
				inst.Prefix[repIndex] |= PrefixIgnored
			}
			repIndex = pos

		case 0x26, 0x2E, 0x36, 0x3E:
			if mode == 64 {
				p |= PrefixIgnored
				break
			}
			fallthrough
		case 0x64, 0x65:
			if segIndex >= 0 {
				inst.Prefix[segIndex] |= PrefixIgnored
			}
			segIndex = pos

		case 0x66:
			if mode == 16 {
				dataMode = 32
				p = PrefixData32
			} else {
				dataMode = 16
				p = PrefixData16
			}
			if dataSizeIndex >= 0 {
				inst.Prefix[dataSizeIndex] |= PrefixIgnored
			}
			dataSizeIndex = pos

		case 0x67:
			if mode == 32 {
				addrMode = 16
				p = PrefixAddr16
			} else {
				addrMode = 32
				p = PrefixAddr32
			}
			if addrSizeIndex >= 0 {
				inst.Prefix[addrSizeIndex] |= PrefixIgnored
			}
			addrSizeIndex = pos

		case 0xC5:
			if pos == 0 && (mode == 64 || (mode == 32 && pos+1 < len(src) && src[pos+1]&0xc0 == 0xc0)) {
				vex = p
				vexIndex = pos
				inst.Prefix[pos] = p
				inst.Prefix[pos+1] = Prefix(src[pos+1])
				pos += 1
				continue
			} else {
				nprefix = pos
				break ReadPrefixes
			}
		case 0xC4:
			if pos == 0 && (mode == 64 || (mode == 32 && pos+2 < len(src) && src[pos+1]&0xc0 == 0xc0)) {
				vex = p
				vexIndex = pos
				inst.Prefix[pos] = p
				inst.Prefix[pos+1] = Prefix(src[pos+1])
				inst.Prefix[pos+2] = Prefix(src[pos+2])
				pos += 2
				continue
			} else {
				nprefix = pos
				break ReadPrefixes
			}
		}

		if pos >= len(inst.Prefix) {
			return instPrefix(src[0], mode)
		}

		inst.Prefix[pos] = p
	}

	if pos < len(src) && mode == 64 && Prefix(src[pos]).IsREX() && vex == 0 {
		rex = Prefix(src[pos])
		rexIndex = pos
		if pos >= len(inst.Prefix) {
			return instPrefix(src[0], mode)
		}
		inst.Prefix[pos] = rex
		pos++
		if rex&PrefixREXW != 0 {
			dataMode = 64
			if dataSizeIndex >= 0 {
				inst.Prefix[dataSizeIndex] |= PrefixIgnored
			}
		}
	}

	opshift = 24
	if decoderCover == nil {
		decoderCover = make([]bool, len(decoder))
	}

	// Decode loop, executing decoder program.
	var oldPC, prevPC int
Decode:
	for pc := 1; ; {
		oldPC = prevPC
		prevPC = pc
		if trace {
			println("run", pc)
		}
		x := decoder[pc]
		decoderCover[pc] = true
		pc++

		switch decodeOp(x) {
		case xCondSlashR, xReadSlashR:
			if haveModrm {
				return Inst{Len: pos}, errInternal
			}
			haveModrm = true
			if pos >= len(src) {
				return truncated(src, mode)
			}
			modrm = int(src[pos])
			pos++
			if opshift >= 0 {
				inst.Opcode |= uint32(modrm) << uint(opshift)
				opshift -= 8
			}
			mod = modrm >> 6
			regop = (modrm >> 3) & 07
			rm = modrm & 07
			if rex&PrefixREXR != 0 {
				rexUsed |= PrefixREXR
				regop |= 8
			}
			if addrMode == 16 {

				if mod != 3 {
					haveMem = true
					mem = addr16[rm]
					if rm == 6 && mod == 0 {
						mem.Base = 0
					}

					if mod == 0 && rm == 6 || mod == 2 {
						if pos+2 > len(src) {
							return truncated(src, mode)
						}
						mem.Disp = int64(binary.LittleEndian.Uint16(src[pos:]))
						pos += 2
					}

					if mod == 1 {
						if pos >= len(src) {
							return truncated(src, mode)
						}
						mem.Disp = int64(int8(src[pos]))
						pos++
					}
				}
			} else {
				haveMem = mod != 3

				if rm == 4 && mod != 3 {
					haveSIB = true
					if pos >= len(src) {
						return truncated(src, mode)
					}
					sib = int(src[pos])
					pos++
					if opshift >= 0 {
						inst.Opcode |= uint32(sib) << uint(opshift)
						opshift -= 8
					}
					scale = sib >> 6
					index = (sib >> 3) & 07
					base = sib & 07
					if rex&PrefixREXB != 0 || vex == 0xC4 && inst.Prefix[vexIndex+1]&0x20 == 0 {
						rexUsed |= PrefixREXB
						base |= 8
					}
					if rex&PrefixREXX != 0 || vex == 0xC4 && inst.Prefix[vexIndex+1]&0x40 == 0 {
						rexUsed |= PrefixREXX
						index |= 8
					}

					mem.Scale = 1 << uint(scale)
					if index == 4 {

					} else {
						mem.Index = baseRegForBits(addrMode) + Reg(index)
					}
					if base&7 == 5 && mod == 0 {

					} else {
						mem.Base = baseRegForBits(addrMode) + Reg(base)
					}
				} else {
					if rex&PrefixREXB != 0 {
						rexUsed |= PrefixREXB
						rm |= 8
					}
					if mod == 0 && rm&7 == 5 || rm&7 == 4 {

					} else if mod != 3 {
						mem.Base = baseRegForBits(addrMode) + Reg(rm)
					}
				}

				if mod == 0 && (rm&7 == 5 || haveSIB && base&7 == 5) || mod == 2 {
					if pos+4 > len(src) {
						return truncated(src, mode)
					}
					dispoff = pos
					displen = 4
					mem.Disp = int64(binary.LittleEndian.Uint32(src[pos:]))
					pos += 4
				}

				if mod == 1 {
					if pos >= len(src) {
						return truncated(src, mode)
					}
					dispoff = pos
					displen = 1
					mem.Disp = int64(int8(src[pos]))
					pos++
				}

				if mode == 64 && mod == 0 && rm&7 == 5 {
					if addrMode == 32 {
						mem.Base = EIP
					} else {
						mem.Base = RIP
					}
				}
			}

			if segIndex >= 0 {
				mem.Segment = prefixToSegment(inst.Prefix[segIndex])
			}
		}

		switch decodeOp(x) {
		default:
			println("bad op", x, "at", pc-1, "from", oldPC)
			return Inst{Len: pos}, errInternal

		case xFail:
			inst.Op = 0
			break Decode

		case xMatch:
			break Decode

		case xJump:
			pc = int(decoder[pc])

		case xCondByte:
			if pos >= len(src) {
				return truncated(src, mode)
			}
			b := src[pos]
			n := int(decoder[pc])
			pc++
			for i := 0; i < n; i++ {
				xb, xpc := decoder[pc], int(decoder[pc+1])
				pc += 2
				if b == byte(xb) {
					pc = xpc
					pos++
					if opshift >= 0 {
						inst.Opcode |= uint32(b) << uint(opshift)
						opshift -= 8
					}
					continue Decode
				}
			}

			if decodeOp(decoder[pc]) == xJump {
				pc = int(decoder[pc+1])
			}
			if decodeOp(decoder[pc]) == xFail {
				pos++
			}

		case xCondIs64:
			if mode == 64 {
				pc = int(decoder[pc+1])
			} else {
				pc = int(decoder[pc])
			}

		case xCondIsMem:
			mem := haveMem
			if !haveModrm {
				if pos >= len(src) {
					return instPrefix(src[0], mode)
				}
				mem = src[pos]>>6 != 3
			}
			if mem {
				pc = int(decoder[pc+1])
			} else {
				pc = int(decoder[pc])
			}

		case xCondDataSize:
			switch dataMode {
			case 16:
				if dataSizeIndex >= 0 {
					inst.Prefix[dataSizeIndex] |= PrefixImplicit
				}
				pc = int(decoder[pc])
			case 32:
				if dataSizeIndex >= 0 {
					inst.Prefix[dataSizeIndex] |= PrefixImplicit
				}
				pc = int(decoder[pc+1])
			case 64:
				rexUsed |= PrefixREXW
				pc = int(decoder[pc+2])
			}

		case xCondAddrSize:
			switch addrMode {
			case 16:
				if addrSizeIndex >= 0 {
					inst.Prefix[addrSizeIndex] |= PrefixImplicit
				}
				pc = int(decoder[pc])
			case 32:
				if addrSizeIndex >= 0 {
					inst.Prefix[addrSizeIndex] |= PrefixImplicit
				}
				pc = int(decoder[pc+1])
			case 64:
				pc = int(decoder[pc+2])
			}

		case xCondPrefix:

			n := int(decoder[pc])
			pc++
			sawF3 := false
			for j := 0; j < n; j++ {
				prefix := Prefix(decoder[pc+2*j])
				if prefix.IsREX() {
					rexUsed |= prefix
					if rex&prefix == prefix {
						pc = int(decoder[pc+2*j+1])
						continue Decode
					}
					continue
				}
				ok := false
				if prefix == 0 {
					ok = true
				} else if prefix.IsREX() {
					rexUsed |= prefix
					if rex&prefix == prefix {
						ok = true
					}
				} else if prefix == 0xC5 || prefix == 0xC4 {
					if vex == prefix {
						ok = true
					}
				} else if vex != 0 && (prefix == 0x0F || prefix == 0x0F38 || prefix == 0x0F3A ||
					prefix == 0x66 || prefix == 0xF2 || prefix == 0xF3) {
					var vexM, vexP Prefix
					if vex == 0xC5 {
						vexM = 1
						vexP = inst.Prefix[vexIndex+1]
					} else {
						vexM = inst.Prefix[vexIndex+1]
						vexP = inst.Prefix[vexIndex+2]
					}
					switch prefix {
					case 0x66:
						ok = vexP&3 == 1
					case 0xF3:
						ok = vexP&3 == 2
					case 0xF2:
						ok = vexP&3 == 3
					case 0x0F:
						ok = vexM&3 == 1
					case 0x0F38:
						ok = vexM&3 == 2
					case 0x0F3A:
						ok = vexM&3 == 3
					}
				} else {
					if prefix == 0xF3 {
						sawF3 = true
					}
					switch prefix {
					case PrefixLOCK:
						if lockIndex >= 0 {
							inst.Prefix[lockIndex] |= PrefixImplicit
							ok = true
						}
					case PrefixREP, PrefixREPN:
						if repIndex >= 0 && inst.Prefix[repIndex]&0xFF == prefix {
							inst.Prefix[repIndex] |= PrefixImplicit
							ok = true
						}
						if gnuCompat && !ok && prefix == 0xF3 && repIndex >= 0 && (j+1 >= n || decoder[pc+2*(j+1)] != 0xF2) {

							for i := repIndex - 1; i >= 0; i-- {
								if inst.Prefix[i]&0xFF == prefix {
									inst.Prefix[i] |= PrefixImplicit
									ok = true
								}
							}
						}
						if gnuCompat && !ok && prefix == 0xF2 && repIndex >= 0 && !sawF3 && inst.Prefix[repIndex]&0xFF == 0xF3 {

							for i := repIndex - 1; i >= 0; i-- {
								if inst.Prefix[i]&0xFF == prefix {
									inst.Prefix[i] |= PrefixImplicit
									ok = true
								}
							}
						}
					case PrefixCS, PrefixDS, PrefixES, PrefixFS, PrefixGS, PrefixSS:
						if segIndex >= 0 && inst.Prefix[segIndex]&0xFF == prefix {
							inst.Prefix[segIndex] |= PrefixImplicit
							ok = true
						}
					case PrefixDataSize:

						if repIndex >= 0 && !gnuCompat {
							inst.Op = 0
							break Decode
						}
						if dataSizeIndex >= 0 {
							inst.Prefix[dataSizeIndex] |= PrefixImplicit
							ok = true
						}
					case PrefixAddrSize:
						if addrSizeIndex >= 0 {
							inst.Prefix[addrSizeIndex] |= PrefixImplicit
							ok = true
						}
					}
				}
				if ok {
					pc = int(decoder[pc+2*j+1])
					continue Decode
				}
			}
			inst.Op = 0
			break Decode

		case xCondSlashR:
			pc = int(decoder[pc+regop&7])

		case xReadSlashR:

		case xReadIb:
			if pos >= len(src) {
				return truncated(src, mode)
			}
			imm8 = int8(src[pos])
			pos++

		case xReadIw:
			if pos+2 > len(src) {
				return truncated(src, mode)
			}
			imm = int64(binary.LittleEndian.Uint16(src[pos:]))
			pos += 2

		case xReadId:
			if pos+4 > len(src) {
				return truncated(src, mode)
			}
			imm = int64(binary.LittleEndian.Uint32(src[pos:]))
			pos += 4

		case xReadIo:
			if pos+8 > len(src) {
				return truncated(src, mode)
			}
			imm = int64(binary.LittleEndian.Uint64(src[pos:]))
			pos += 8

		case xReadCb:
			if pos >= len(src) {
				return truncated(src, mode)
			}
			immcpos = pos
			immc = int64(src[pos])
			pos++

		case xReadCw:
			if pos+2 > len(src) {
				return truncated(src, mode)
			}
			immcpos = pos
			immc = int64(binary.LittleEndian.Uint16(src[pos:]))
			pos += 2

		case xReadCm:
			immcpos = pos
			if addrMode == 16 {
				if pos+2 > len(src) {
					return truncated(src, mode)
				}
				immc = int64(binary.LittleEndian.Uint16(src[pos:]))
				pos += 2
			} else if addrMode == 32 {
				if pos+4 > len(src) {
					return truncated(src, mode)
				}
				immc = int64(binary.LittleEndian.Uint32(src[pos:]))
				pos += 4
			} else {
				if pos+8 > len(src) {
					return truncated(src, mode)
				}
				immc = int64(binary.LittleEndian.Uint64(src[pos:]))
				pos += 8
			}
		case xReadCd:
			immcpos = pos
			if pos+4 > len(src) {
				return truncated(src, mode)
			}
			immc = int64(binary.LittleEndian.Uint32(src[pos:]))
			pos += 4

		case xReadCp:
			immcpos = pos
			if pos+6 > len(src) {
				return truncated(src, mode)
			}
			w := binary.LittleEndian.Uint32(src[pos:])
			w2 := binary.LittleEndian.Uint16(src[pos+4:])
			immc = int64(w2)<<32 | int64(w)
			pos += 6

		case xSetOp:
			inst.Op = Op(decoder[pc])
			pc++

		case xArg1,
			xArg3,
			xArgAL,
			xArgAX,
			xArgCL,
			xArgCS,
			xArgDS,
			xArgDX,
			xArgEAX,
			xArgEDX,
			xArgES,
			xArgFS,
			xArgGS,
			xArgRAX,
			xArgRDX,
			xArgSS,
			xArgST,
			xArgXMM0:
			inst.Args[narg] = fixedArg[x]
			narg++

		case xArgImm8:
			inst.Args[narg] = Imm(imm8)
			narg++

		case xArgImm8u:
			inst.Args[narg] = Imm(uint8(imm8))
			narg++

		case xArgImm16:
			inst.Args[narg] = Imm(int16(imm))
			narg++

		case xArgImm16u:
			inst.Args[narg] = Imm(uint16(imm))
			narg++

		case xArgImm32:
			inst.Args[narg] = Imm(int32(imm))
			narg++

		case xArgImm64:
			inst.Args[narg] = Imm(imm)
			narg++

		case xArgM,
			xArgM128,
			xArgM256,
			xArgM1428byte,
			xArgM16,
			xArgM16and16,
			xArgM16and32,
			xArgM16and64,
			xArgM16colon16,
			xArgM16colon32,
			xArgM16colon64,
			xArgM16int,
			xArgM2byte,
			xArgM32,
			xArgM32and32,
			xArgM32fp,
			xArgM32int,
			xArgM512byte,
			xArgM64,
			xArgM64fp,
			xArgM64int,
			xArgM8,
			xArgM80bcd,
			xArgM80dec,
			xArgM80fp,
			xArgM94108byte,
			xArgMem:
			if !haveMem {
				inst.Op = 0
				break Decode
			}
			inst.Args[narg] = mem
			inst.MemBytes = int(memBytes[decodeOp(x)])
			if mem.Base == RIP {
				inst.PCRel = displen
				inst.PCRelOff = dispoff
			}
			narg++

		case xArgPtr16colon16:
			inst.Args[narg] = Imm(immc >> 16)
			inst.Args[narg+1] = Imm(immc & (1<<16 - 1))
			narg += 2

		case xArgPtr16colon32:
			inst.Args[narg] = Imm(immc >> 32)
			inst.Args[narg+1] = Imm(immc & (1<<32 - 1))
			narg += 2

		case xArgMoffs8, xArgMoffs16, xArgMoffs32, xArgMoffs64:

			mem = Mem{Disp: int64(immc)}
			if segIndex >= 0 {
				mem.Segment = prefixToSegment(inst.Prefix[segIndex])
				inst.Prefix[segIndex] |= PrefixImplicit
			}
			inst.Args[narg] = mem
			inst.MemBytes = int(memBytes[decodeOp(x)])
			if mem.Base == RIP {
				inst.PCRel = displen
				inst.PCRelOff = dispoff
			}
			narg++

		case xArgYmm1:
			base := baseReg[x]
			index := Reg(regop)
			if inst.Prefix[vexIndex+1]&0x80 == 0 {
				index += 8
			}
			inst.Args[narg] = base + index
			narg++

		case xArgR8, xArgR16, xArgR32, xArgR64, xArgXmm, xArgXmm1, xArgDR0dashDR7:
			base := baseReg[x]
			index := Reg(regop)
			if rex != 0 && base == AL && index >= 4 {
				rexUsed |= PrefixREX
				index -= 4
				base = SPB
			}
			inst.Args[narg] = base + index
			narg++

		case xArgMm, xArgMm1, xArgTR0dashTR7:
			inst.Args[narg] = baseReg[x] + Reg(regop&7)
			narg++

		case xArgCR0dashCR7:

			if lockIndex >= 0 {
				inst.Prefix[lockIndex] |= PrefixImplicit
				regop += 8
			}
			inst.Args[narg] = CR0 + Reg(regop)
			narg++

		case xArgSreg:
			regop &= 7
			if regop >= 6 {
				inst.Op = 0
				break Decode
			}
			inst.Args[narg] = ES + Reg(regop)
			narg++

		case xArgRmf16, xArgRmf32, xArgRmf64:
			base := baseReg[x]
			index := Reg(modrm & 07)
			if rex&PrefixREXB != 0 {
				rexUsed |= PrefixREXB
				index += 8
			}
			inst.Args[narg] = base + index
			narg++

		case xArgR8op, xArgR16op, xArgR32op, xArgR64op, xArgSTi:
			n := inst.Opcode >> uint(opshift+8) & 07
			base := baseReg[x]
			index := Reg(n)
			if rex&PrefixREXB != 0 && decodeOp(x) != xArgSTi {
				rexUsed |= PrefixREXB
				index += 8
			}
			if rex != 0 && base == AL && index >= 4 {
				rexUsed |= PrefixREX
				index -= 4
				base = SPB
			}
			inst.Args[narg] = base + index
			narg++
		case xArgRM8, xArgRM16, xArgRM32, xArgRM64, xArgR32M16, xArgR32M8, xArgR64M16,
			xArgMmM32, xArgMmM64, xArgMm2M64,
			xArgXmm2M16, xArgXmm2M32, xArgXmm2M64, xArgXmmM64, xArgXmmM128, xArgXmmM32, xArgXmm2M128,
			xArgYmm2M256:
			if haveMem {
				inst.Args[narg] = mem
				inst.MemBytes = int(memBytes[decodeOp(x)])
				if mem.Base == RIP {
					inst.PCRel = displen
					inst.PCRelOff = dispoff
				}
			} else {
				base := baseReg[x]
				index := Reg(rm)
				switch decodeOp(x) {
				case xArgMmM32, xArgMmM64, xArgMm2M64:

					index &= 7
				case xArgRM8:
					if rex != 0 && index >= 4 {
						rexUsed |= PrefixREX
						index -= 4
						base = SPB
					}
				case xArgYmm2M256:
					if vex == 0xC4 && inst.Prefix[vexIndex+1]&0x40 == 0x40 {
						index += 8
					}
				}
				inst.Args[narg] = base + index
			}
			narg++

		case xArgMm2:
			if haveMem {
				inst.Op = 0
				break Decode
			}
			inst.Args[narg] = baseReg[x] + Reg(rm&7)
			narg++

		case xArgXmm2:
			if haveMem {
				inst.Op = 0
				break Decode
			}
			inst.Args[narg] = baseReg[x] + Reg(rm)
			narg++

		case xArgRel8:
			inst.PCRelOff = immcpos
			inst.PCRel = 1
			inst.Args[narg] = Rel(int8(immc))
			narg++

		case xArgRel16:
			inst.PCRelOff = immcpos
			inst.PCRel = 2
			inst.Args[narg] = Rel(int16(immc))
			narg++

		case xArgRel32:
			inst.PCRelOff = immcpos
			inst.PCRel = 4
			inst.Args[narg] = Rel(int32(immc))
			narg++
		}
	}

	if inst.Op == 0 {

		if nprefix > 0 {
			return instPrefix(src[0], mode)
		}
		return Inst{Len: pos}, ErrUnrecognized
	}

	if inst.Op == XCHG && inst.Opcode>>24 == 0x90 {
		if inst.Args[0] == RAX || inst.Args[0] == EAX || inst.Args[0] == AX {
			inst.Op = NOP
			if dataSizeIndex >= 0 {
				inst.Prefix[dataSizeIndex] &^= PrefixImplicit
			}
			inst.Args[0] = nil
			inst.Args[1] = nil
		}
		if repIndex >= 0 && inst.Prefix[repIndex] == 0xF3 {
			inst.Prefix[repIndex] |= PrefixImplicit
			inst.Op = PAUSE
			inst.Args[0] = nil
			inst.Args[1] = nil
		} else if gnuCompat {
			for i := nprefix - 1; i >= 0; i-- {
				if inst.Prefix[i]&0xFF == 0xF3 {
					inst.Prefix[i] |= PrefixImplicit
					inst.Op = PAUSE
					inst.Args[0] = nil
					inst.Args[1] = nil
					break
				}
			}
		}
	}

	defaultSeg := func() Reg {
		if segIndex >= 0 {
			inst.Prefix[segIndex] |= PrefixImplicit
			return prefixToSegment(inst.Prefix[segIndex])
		}
		return DS
	}

	usedAddrSize := false
	switch inst.Op {
	case INSB, INSW, INSD:
		inst.Args[0] = Mem{Segment: ES, Base: baseRegForBits(addrMode) + DI - AX}
		inst.Args[1] = DX
		usedAddrSize = true

	case OUTSB, OUTSW, OUTSD:
		inst.Args[0] = DX
		inst.Args[1] = Mem{Segment: defaultSeg(), Base: baseRegForBits(addrMode) + SI - AX}
		usedAddrSize = true

	case MOVSB, MOVSW, MOVSD, MOVSQ:
		inst.Args[0] = Mem{Segment: ES, Base: baseRegForBits(addrMode) + DI - AX}
		inst.Args[1] = Mem{Segment: defaultSeg(), Base: baseRegForBits(addrMode) + SI - AX}
		usedAddrSize = true

	case CMPSB, CMPSW, CMPSD, CMPSQ:
		inst.Args[0] = Mem{Segment: defaultSeg(), Base: baseRegForBits(addrMode) + SI - AX}
		inst.Args[1] = Mem{Segment: ES, Base: baseRegForBits(addrMode) + DI - AX}
		usedAddrSize = true

	case LODSB, LODSW, LODSD, LODSQ:
		switch inst.Op {
		case LODSB:
			inst.Args[0] = AL
		case LODSW:
			inst.Args[0] = AX
		case LODSD:
			inst.Args[0] = EAX
		case LODSQ:
			inst.Args[0] = RAX
		}
		inst.Args[1] = Mem{Segment: defaultSeg(), Base: baseRegForBits(addrMode) + SI - AX}
		usedAddrSize = true

	case STOSB, STOSW, STOSD, STOSQ:
		inst.Args[0] = Mem{Segment: ES, Base: baseRegForBits(addrMode) + DI - AX}
		switch inst.Op {
		case STOSB:
			inst.Args[1] = AL
		case STOSW:
			inst.Args[1] = AX
		case STOSD:
			inst.Args[1] = EAX
		case STOSQ:
			inst.Args[1] = RAX
		}
		usedAddrSize = true

	case SCASB, SCASW, SCASD, SCASQ:
		inst.Args[1] = Mem{Segment: ES, Base: baseRegForBits(addrMode) + DI - AX}
		switch inst.Op {
		case SCASB:
			inst.Args[0] = AL
		case SCASW:
			inst.Args[0] = AX
		case SCASD:
			inst.Args[0] = EAX
		case SCASQ:
			inst.Args[0] = RAX
		}
		usedAddrSize = true

	case XLATB:
		inst.Args[0] = Mem{Segment: defaultSeg(), Base: baseRegForBits(addrMode) + BX - AX}
		usedAddrSize = true
	}

	if haveMem || usedAddrSize {
		if addrSizeIndex >= 0 {
			inst.Prefix[addrSizeIndex] |= PrefixImplicit
		}
	}

	if haveMem {
		if segIndex >= 0 {
			inst.Prefix[segIndex] |= PrefixImplicit
		}
	}

	if isCondJmp[inst.Op] || isLoop[inst.Op] || inst.Op == JCXZ || inst.Op == JECXZ || inst.Op == JRCXZ {
	PredictLoop:
		for i := nprefix - 1; i >= 0; i-- {
			p := inst.Prefix[i]
			switch p & 0xFF {
			case PrefixCS:
				inst.Prefix[i] = PrefixPN
				break PredictLoop
			case PrefixDS:
				inst.Prefix[i] = PrefixPT
				break PredictLoop
			}
		}
	}

	if isCondJmp[inst.Op] || inst.Op == JMP || inst.Op == CALL || inst.Op == RET {
		for i := nprefix - 1; i >= 0; i-- {
			p := inst.Prefix[i]
			if p&^PrefixIgnored == PrefixREPN {
				inst.Prefix[i] = PrefixBND
				break
			}
		}
	}

	hasLock := false
	if lockIndex >= 0 && inst.Prefix[lockIndex]&PrefixImplicit == 0 {
		switch inst.Op {

		case ADD, ADC, AND, BTC, BTR, BTS, CMPXCHG, CMPXCHG8B, CMPXCHG16B, DEC, INC, NEG, NOT, OR, SBB, SUB, XOR, XADD, XCHG:
			if isMem(inst.Args[0]) {
				hasLock = true
				break
			}
			fallthrough
		default:
			inst.Prefix[lockIndex] |= PrefixInvalid
		}
	}

	if isMem(inst.Args[0]) {
		if inst.Op == XCHG {
			hasLock = true
		}

		for i := len(inst.Prefix) - 1; i >= 0; i-- {
			p := inst.Prefix[i] &^ PrefixIgnored
			switch p {
			case PrefixREPN:
				if hasLock {
					inst.Prefix[i] = inst.Prefix[i]&PrefixIgnored | PrefixXACQUIRE
				}

			case PrefixREP:
				if hasLock {
					inst.Prefix[i] = inst.Prefix[i]&PrefixIgnored | PrefixXRELEASE
				}

				if inst.Op == MOV {
					op := (inst.Opcode >> 24) &^ 1
					if op == 0x88 || op == 0xC6 {
						inst.Prefix[i] = inst.Prefix[i]&PrefixIgnored | PrefixXRELEASE
					}
				}
			}
		}
	}

	if repIndex >= 0 {
		switch inst.Prefix[repIndex] {
		case PrefixREP, PrefixREPN:
			switch inst.Op {

			case INSB, INSW, INSD,
				MOVSB, MOVSW, MOVSD, MOVSQ,
				OUTSB, OUTSW, OUTSD,
				LODSB, LODSW, LODSD, LODSQ,
				CMPSB, CMPSW, CMPSD, CMPSQ,
				SCASB, SCASW, SCASD, SCASQ,
				STOSB, STOSW, STOSD, STOSQ:

			default:
				inst.Prefix[repIndex] |= PrefixIgnored
			}
		}
	}

	if rexIndex >= 0 {
		if rexUsed != 0 {
			rexUsed |= PrefixREX
		}
		if rex&^rexUsed == 0 {
			inst.Prefix[rexIndex] |= PrefixImplicit
		}
	}

	inst.DataSize = dataMode
	inst.AddrSize = addrMode
	inst.Mode = mode
	inst.Len = pos
	return inst, nil
}

var errInternal = errors.New("internal error")

// addr16 records the eight 16-bit addressing modes.
var addr16 = [8]Mem{
	{Base: BX, Scale: 1, Index: SI},
	{Base: BX, Scale: 1, Index: DI},
	{Base: BP, Scale: 1, Index: SI},
	{Base: BP, Scale: 1, Index: DI},
	{Base: SI},
	{Base: DI},
	{Base: BP},
	{Base: BX},
}

// baseReg returns the base register for a given register size in bits.
func baseRegForBits(bits int) Reg {
	switch bits {
	case 8:
		return AL
	case 16:
		return AX
	case 32:
		return EAX
	case 64:
		return RAX
	}
	return 0
}

// baseReg records the base register for argument types that specify
// a range of registers indexed by op, regop, or rm.
var baseReg = [...]Reg{
	xArgDR0dashDR7: DR0,
	xArgMm1:        M0,
	xArgMm2:        M0,
	xArgMm2M64:     M0,
	xArgMm:         M0,
	xArgMmM32:      M0,
	xArgMmM64:      M0,
	xArgR16:        AX,
	xArgR16op:      AX,
	xArgR32:        EAX,
	xArgR32M16:     EAX,
	xArgR32M8:      EAX,
	xArgR32op:      EAX,
	xArgR64:        RAX,
	xArgR64M16:     RAX,
	xArgR64op:      RAX,
	xArgR8:         AL,
	xArgR8op:       AL,
	xArgRM16:       AX,
	xArgRM32:       EAX,
	xArgRM64:       RAX,
	xArgRM8:        AL,
	xArgRmf16:      AX,
	xArgRmf32:      EAX,
	xArgRmf64:      RAX,
	xArgSTi:        F0,
	xArgTR0dashTR7: TR0,
	xArgXmm1:       X0,
	xArgYmm1:       X0,
	xArgXmm2:       X0,
	xArgXmm2M128:   X0,
	xArgYmm2M256:   X0,
	xArgXmm2M16:    X0,
	xArgXmm2M32:    X0,
	xArgXmm2M64:    X0,
	xArgXmm:        X0,
	xArgXmmM128:    X0,
	xArgXmmM32:     X0,
	xArgXmmM64:     X0,
}

// prefixToSegment returns the segment register
// corresponding to a particular segment prefix.
func prefixToSegment(p Prefix) Reg {
	switch p &^ PrefixImplicit {
	case PrefixCS:
		return CS
	case PrefixDS:
		return DS
	case PrefixES:
		return ES
	case PrefixFS:
		return FS
	case PrefixGS:
		return GS
	case PrefixSS:
		return SS
	}
	return 0
}

// fixedArg records the fixed arguments corresponding to the given bytecodes.
var fixedArg = [...]Arg{
	xArg1:    Imm(1),
	xArg3:    Imm(3),
	xArgAL:   AL,
	xArgAX:   AX,
	xArgDX:   DX,
	xArgEAX:  EAX,
	xArgEDX:  EDX,
	xArgRAX:  RAX,
	xArgRDX:  RDX,
	xArgCL:   CL,
	xArgCS:   CS,
	xArgDS:   DS,
	xArgES:   ES,
	xArgFS:   FS,
	xArgGS:   GS,
	xArgSS:   SS,
	xArgST:   F0,
	xArgXMM0: X0,
}

// memBytes records the size of the memory pointed at
// by a memory argument of the given form.
var memBytes = [...]int8{
	xArgM128:       128 / 8,
	xArgM256:       256 / 8,
	xArgM16:        16 / 8,
	xArgM16and16:   (16 + 16) / 8,
	xArgM16colon16: (16 + 16) / 8,
	xArgM16colon32: (16 + 32) / 8,
	xArgM16int:     16 / 8,
	xArgM2byte:     2,
	xArgM32:        32 / 8,
	xArgM32and32:   (32 + 32) / 8,
	xArgM32fp:      32 / 8,
	xArgM32int:     32 / 8,
	xArgM64:        64 / 8,
	xArgM64fp:      64 / 8,
	xArgM64int:     64 / 8,
	xArgMm2M64:     64 / 8,
	xArgMmM32:      32 / 8,
	xArgMmM64:      64 / 8,
	xArgMoffs16:    16 / 8,
	xArgMoffs32:    32 / 8,
	xArgMoffs64:    64 / 8,
	xArgMoffs8:     8 / 8,
	xArgR32M16:     16 / 8,
	xArgR32M8:      8 / 8,
	xArgR64M16:     16 / 8,
	xArgRM16:       16 / 8,
	xArgRM32:       32 / 8,
	xArgRM64:       64 / 8,
	xArgRM8:        8 / 8,
	xArgXmm2M128:   128 / 8,
	xArgYmm2M256:   256 / 8,
	xArgXmm2M16:    16 / 8,
	xArgXmm2M32:    32 / 8,
	xArgXmm2M64:    64 / 8,
	xArgXmm:        128 / 8,
	xArgXmmM128:    128 / 8,
	xArgXmmM32:     32 / 8,
	xArgXmmM64:     64 / 8,
}

// isCondJmp records the conditional jumps.
var isCondJmp = [maxOp + 1]bool{
	JA:  true,
	JAE: true,
	JB:  true,
	JBE: true,
	JE:  true,
	JG:  true,
	JGE: true,
	JL:  true,
	JLE: true,
	JNE: true,
	JNO: true,
	JNP: true,
	JNS: true,
	JO:  true,
	JP:  true,
	JS:  true,
}

// isLoop records the loop operators.
var isLoop = [maxOp + 1]bool{
	LOOP:   true,
	LOOPE:  true,
	LOOPNE: true,
	JECXZ:  true,
	JRCXZ:  true,
}

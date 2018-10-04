// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package amd64

import (
	"github.com/dave/golib/src/cmd/compile/internal/gc"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/obj/x86"
	"github.com/dave/golib/src/cmd/internal/objabi"
)

// DUFFZERO consists of repeated blocks of 4 MOVUPSs + LEAQ,
// See runtime/mkduff.go.
const (
	dzBlocks    = 16 // number of MOV/ADD blocks
	dzBlockLen  = 4  // number of clears per block
	dzBlockSize = 19 // size of instructions in a single block
	dzMovSize   = 4  // size of single MOV instruction w/ offset
	dzLeaqSize  = 4  // size of single LEAQ instruction
	dzClearStep = 16 // number of bytes cleared by each MOV instruction

	dzClearLen = dzClearStep * dzBlockLen // bytes cleared by one block
	dzSize     = dzBlocks * dzBlockSize
)

// dzOff returns the offset for a jump into DUFFZERO.
// b is the number of bytes to zero.
func dzOff(b int64) int64 {
	off := int64(dzSize)
	off -= b / dzClearLen * dzBlockSize
	tailLen := b % dzClearLen
	if tailLen >= dzClearStep {
		off -= dzLeaqSize + dzMovSize*(tailLen/dzClearStep)
	}
	return off
}

// duffzeroDI returns the pre-adjustment to DI for a call to DUFFZERO.
// b is the number of bytes to zero.
func dzDI(b int64) int64 {
	tailLen := b % dzClearLen
	if tailLen < dzClearStep {
		return 0
	}
	tailSteps := tailLen / dzClearStep
	return -dzClearStep * (dzBlockLen - tailSteps)
}

func (pstate *PackageState) zerorange(pp *gc.Progs, p *obj.Prog, off, cnt int64, state *uint32) *obj.Prog {
	const (
		ax = 1 << iota
		x0
	)

	if cnt == 0 {
		return p
	}

	if cnt%int64(pstate.gc.Widthreg) != 0 {
		// should only happen with nacl
		if cnt%int64(pstate.gc.Widthptr) != 0 {
			pstate.gc.Fatalf("zerorange count not a multiple of widthptr %d", cnt)
		}
		if *state&ax == 0 {
			p = pp.Appendpp(pstate.gc, p, x86.AMOVQ, obj.TYPE_CONST, 0, 0, obj.TYPE_REG, x86.REG_AX, 0)
			*state |= ax
		}
		p = pp.Appendpp(pstate.gc, p, x86.AMOVL, obj.TYPE_REG, x86.REG_AX, 0, obj.TYPE_MEM, x86.REG_SP, off)
		off += int64(pstate.gc.Widthptr)
		cnt -= int64(pstate.gc.Widthptr)
	}

	if cnt == 8 {
		if *state&ax == 0 {
			p = pp.Appendpp(pstate.gc, p, x86.AMOVQ, obj.TYPE_CONST, 0, 0, obj.TYPE_REG, x86.REG_AX, 0)
			*state |= ax
		}
		p = pp.Appendpp(pstate.gc, p, x86.AMOVQ, obj.TYPE_REG, x86.REG_AX, 0, obj.TYPE_MEM, x86.REG_SP, off)
	} else if !pstate.isPlan9 && cnt <= int64(8*pstate.gc.Widthreg) {
		if *state&x0 == 0 {
			p = pp.Appendpp(pstate.gc, p, x86.AXORPS, obj.TYPE_REG, x86.REG_X0, 0, obj.TYPE_REG, x86.REG_X0, 0)
			*state |= x0
		}

		for i := int64(0); i < cnt/16; i++ {
			p = pp.Appendpp(pstate.gc, p, x86.AMOVUPS, obj.TYPE_REG, x86.REG_X0, 0, obj.TYPE_MEM, x86.REG_SP, off+i*16)
		}

		if cnt%16 != 0 {
			p = pp.Appendpp(pstate.gc, p, x86.AMOVUPS, obj.TYPE_REG, x86.REG_X0, 0, obj.TYPE_MEM, x86.REG_SP, off+cnt-int64(16))
		}
	} else if !pstate.gc.Nacl && !pstate.isPlan9 && (cnt <= int64(128*pstate.gc.Widthreg)) {
		if *state&x0 == 0 {
			p = pp.Appendpp(pstate.gc, p, x86.AXORPS, obj.TYPE_REG, x86.REG_X0, 0, obj.TYPE_REG, x86.REG_X0, 0)
			*state |= x0
		}
		p = pp.Appendpp(pstate.gc, p, pstate.leaptr, obj.TYPE_MEM, x86.REG_SP, off+dzDI(cnt), obj.TYPE_REG, x86.REG_DI, 0)
		p = pp.Appendpp(pstate.gc, p, obj.ADUFFZERO, obj.TYPE_NONE, 0, 0, obj.TYPE_ADDR, 0, dzOff(cnt))
		p.To.Sym = pstate.gc.Duffzero

		if cnt%16 != 0 {
			p = pp.Appendpp(pstate.gc, p, x86.AMOVUPS, obj.TYPE_REG, x86.REG_X0, 0, obj.TYPE_MEM, x86.REG_DI, -int64(8))
		}
	} else {
		if *state&ax == 0 {
			p = pp.Appendpp(pstate.gc, p, x86.AMOVQ, obj.TYPE_CONST, 0, 0, obj.TYPE_REG, x86.REG_AX, 0)
			*state |= ax
		}

		p = pp.Appendpp(pstate.gc, p, x86.AMOVQ, obj.TYPE_CONST, 0, cnt/int64(pstate.gc.Widthreg), obj.TYPE_REG, x86.REG_CX, 0)
		p = pp.Appendpp(pstate.gc, p, pstate.leaptr, obj.TYPE_MEM, x86.REG_SP, off, obj.TYPE_REG, x86.REG_DI, 0)
		p = pp.Appendpp(pstate.gc, p, x86.AREP, obj.TYPE_NONE, 0, 0, obj.TYPE_NONE, 0, 0)
		p = pp.Appendpp(pstate.gc, p, x86.ASTOSQ, obj.TYPE_NONE, 0, 0, obj.TYPE_NONE, 0, 0)
	}

	return p
}

func (pstate *PackageState) zeroAuto(pp *gc.Progs, n *gc.Node) {
	// Note: this code must not clobber any registers.
	op := x86.AMOVQ
	if pstate.gc.Widthptr == 4 {
		op = x86.AMOVL
	}
	sym := n.Sym.Linksym(pstate.types)
	size := n.Type.Size(pstate.types)
	for i := int64(0); i < size; i += int64(pstate.gc.Widthptr) {
		p := pp.Prog(pstate.gc, op)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 0
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_AUTO
		p.To.Reg = x86.REG_SP
		p.To.Offset = n.Xoffset + i
		p.To.Sym = sym
	}
}

func (pstate *PackageState) ginsnop(pp *gc.Progs) {
	// This is actually not the x86 NOP anymore,
	// but at the point where it gets used, AX is dead
	// so it's okay if we lose the high bits.
	p := pp.Prog(pstate.gc, x86.AXCHGL)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = x86.REG_AX
	p.To.Type = obj.TYPE_REG
	p.To.Reg = x86.REG_AX
}

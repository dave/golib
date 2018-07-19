package ld

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/gcprog"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// isRuntimeDepPkg returns whether pkg is the runtime package or its dependency
func isRuntimeDepPkg(pkg string) bool {
	switch pkg {
	case "runtime",
		"sync/atomic",
		"internal/bytealg",
		"internal/cpu":
		return true
	}
	return strings.HasPrefix(pkg, "runtime/internal/") && !strings.HasSuffix(pkg, "_test")
}

// Estimate the max size needed to hold any new trampolines created for this function. This
// is used to determine when the section can be split if it becomes too large, to ensure that
// the trampolines are in the same section as the function that uses them.
func (psess *PackageSession) maxSizeTrampolinesPPC64(s *sym.Symbol, isTramp bool) uint64 {

	if psess.thearch.Trampoline == nil || isTramp {
		return 0
	}

	n := uint64(0)
	for ri := range s.R {
		r := &s.R[ri]
		if r.Type.IsDirectJump() {
			n++
		}
	}

	return n * 16
}

// detect too-far jumps in function s, and add trampolines if necessary
// ARM, PPC64 & PPC64LE support trampoline insertion for internal and external linking
// On PPC64 & PPC64LE the text sections might be split but will still insert trampolines
// where necessary.
func (psess *PackageSession) trampoline(ctxt *Link, s *sym.Symbol) {
	if psess.thearch.Trampoline == nil {
		return
	}

	for ri := range s.R {
		r := &s.R[ri]
		if !r.Type.IsDirectJump() {
			continue
		}
		if psess.Symaddr(r.Sym) == 0 && r.Sym.Type != sym.SDYNIMPORT {
			if r.Sym.File != s.File {
				if !isRuntimeDepPkg(s.File) || !isRuntimeDepPkg(r.Sym.File) {
					ctxt.ErrorUnresolved(psess, s, r)
				}

			}
			continue
		}
		psess.
			thearch.Trampoline(ctxt, r, s)
	}

}

// resolve relocations in s.
func (psess *PackageSession) relocsym(ctxt *Link, s *sym.Symbol) {
	for ri := int32(0); ri < int32(len(s.R)); ri++ {
		r := &s.R[ri]
		if r.Done {

			continue
		}
		r.Done = true
		off := r.Off
		siz := int32(r.Siz)
		if off < 0 || off+siz > int32(len(s.P)) {
			rname := ""
			if r.Sym != nil {
				rname = r.Sym.Name
			}
			psess.
				Errorf(s, "invalid relocation %s: %d+%d not in [%d,%d)", rname, off, siz, 0, len(s.P))
			continue
		}

		if r.Sym != nil && ((r.Sym.Type == sym.Sxxx && !r.Sym.Attr.VisibilityHidden()) || r.Sym.Type == sym.SXREF) {

			if ctxt.BuildMode == BuildModeShared {
				if r.Sym.Name == "main.main" || r.Sym.Name == "main.init" {
					r.Sym.Type = sym.SDYNIMPORT
				} else if strings.HasPrefix(r.Sym.Name, "go.info.") {

					continue
				}
			} else {
				ctxt.ErrorUnresolved(psess, s, r)
				continue
			}
		}

		if r.Type >= 256 {
			continue
		}
		if r.Siz == 0 {
			continue
		}
		if r.Type == objabi.R_DWARFFILEREF {
			psess.
				Errorf(s, "orphan R_DWARFFILEREF reloc to %v", r.Sym.Name)
			continue
		}

		if ctxt.HeadType != objabi.Hsolaris && ctxt.HeadType != objabi.Hdarwin && r.Sym != nil && r.Sym.Type == sym.SDYNIMPORT && !ctxt.DynlinkingGo() && !r.Sym.Attr.SubSymbol() {
			if !(ctxt.Arch.Family == sys.PPC64 && ctxt.LinkMode == LinkExternal && r.Sym.Name == ".TOC.") {
				psess.
					Errorf(s, "unhandled relocation for %s (type %d (%s) rtype %d (%s))", r.Sym.Name, r.Sym.Type, r.Sym.Type, r.Type, psess.sym.RelocName(ctxt.Arch, r.Type))
			}
		}
		if r.Sym != nil && r.Sym.Type != sym.STLSBSS && r.Type != objabi.R_WEAKADDROFF && !r.Sym.Attr.Reachable() {
			psess.
				Errorf(s, "unreachable sym in relocation: %s", r.Sym.Name)
		}

		if ctxt.Arch.Family == sys.S390X {
			switch r.Type {
			case objabi.R_PCRELDBL:
				r.Type = objabi.R_PCREL
				r.Variant = sym.RV_390_DBL
			case objabi.R_CALL:
				r.Variant = sym.RV_390_DBL
			}
		}

		var o int64
		switch r.Type {
		default:
			switch siz {
			default:
				psess.
					Errorf(s, "bad reloc size %#x for %s", uint32(siz), r.Sym.Name)
			case 1:
				o = int64(s.P[off])
			case 2:
				o = int64(ctxt.Arch.ByteOrder.Uint16(s.P[off:]))
			case 4:
				o = int64(ctxt.Arch.ByteOrder.Uint32(s.P[off:]))
			case 8:
				o = int64(ctxt.Arch.ByteOrder.Uint64(s.P[off:]))
			}
			if !psess.thearch.Archreloc(ctxt, r, s, &o) {
				psess.
					Errorf(s, "unknown reloc to %v: %d (%s)", r.Sym.Name, r.Type, psess.sym.RelocName(ctxt.Arch, r.Type))
			}
		case objabi.R_TLS_LE:
			isAndroidX86 := psess.objabi.GOOS == "android" && (ctxt.Arch.InFamily(sys.AMD64, sys.I386))

			if ctxt.LinkMode == LinkExternal && ctxt.IsELF && !isAndroidX86 {
				r.Done = false
				if r.Sym == nil {
					r.Sym = ctxt.Tlsg
				}
				r.Xsym = r.Sym
				r.Xadd = r.Add
				o = 0
				if ctxt.Arch.Family != sys.AMD64 {
					o = r.Add
				}
				break
			}

			if ctxt.IsELF && ctxt.Arch.Family == sys.ARM {

				o = 8 + r.Sym.Value
			} else if ctxt.IsELF || ctxt.HeadType == objabi.Hplan9 || ctxt.HeadType == objabi.Hdarwin || isAndroidX86 {
				o = int64(ctxt.Tlsoffset) + r.Add
			} else if ctxt.HeadType == objabi.Hwindows {
				o = r.Add
			} else {
				log.Fatalf("unexpected R_TLS_LE relocation for %v", ctxt.HeadType)
			}
		case objabi.R_TLS_IE:
			isAndroidX86 := psess.objabi.GOOS == "android" && (ctxt.Arch.InFamily(sys.AMD64, sys.I386))

			if ctxt.LinkMode == LinkExternal && ctxt.IsELF && !isAndroidX86 {
				r.Done = false
				if r.Sym == nil {
					r.Sym = ctxt.Tlsg
				}
				r.Xsym = r.Sym
				r.Xadd = r.Add
				o = 0
				if ctxt.Arch.Family != sys.AMD64 {
					o = r.Add
				}
				break
			}
			if ctxt.BuildMode == BuildModePIE && ctxt.IsELF {

				if psess.thearch.TLSIEtoLE == nil {
					log.Fatalf("internal linking of TLS IE not supported on %v", ctxt.Arch.Family)
				}
				psess.
					thearch.TLSIEtoLE(s, int(off), int(r.Siz))
				o = int64(ctxt.Tlsoffset)

			} else {
				log.Fatalf("cannot handle R_TLS_IE (sym %s) when linking internally", s.Name)
			}
		case objabi.R_ADDR:
			if ctxt.LinkMode == LinkExternal && r.Sym.Type != sym.SCONST {
				r.Done = false

				rs := r.Sym

				r.Xadd = r.Add
				for rs.Outer != nil {
					r.Xadd += psess.Symaddr(rs) - psess.Symaddr(rs.Outer)
					rs = rs.Outer
				}

				if rs.Type != sym.SHOSTOBJ && rs.Type != sym.SDYNIMPORT && rs.Sect == nil {
					psess.
						Errorf(s, "missing section for relocation target %s", rs.Name)
				}
				r.Xsym = rs

				o = r.Xadd
				if ctxt.IsELF {
					if ctxt.Arch.Family == sys.AMD64 {
						o = 0
					}
				} else if ctxt.HeadType == objabi.Hdarwin {
					if rs.Type != sym.SHOSTOBJ {
						o += psess.Symaddr(rs)
					}
				} else if ctxt.HeadType == objabi.Hwindows {

				} else {
					psess.
						Errorf(s, "unhandled pcrel relocation to %s on %v", rs.Name, ctxt.HeadType)
				}

				break
			}

			o = psess.Symaddr(r.Sym) + r.Add

			if int32(o) < 0 && ctxt.Arch.PtrSize > 4 && siz == 4 {
				psess.
					Errorf(s, "non-pc-relative relocation address for %s is too big: %#x (%#x + %#x)", r.Sym.Name, uint64(o), psess.Symaddr(r.Sym), r.Add)
				psess.
					errorexit()
			}
		case objabi.R_DWARFSECREF:
			if r.Sym.Sect == nil {
				psess.
					Errorf(s, "missing DWARF section for relocation target %s", r.Sym.Name)
			}

			if ctxt.LinkMode == LinkExternal {
				r.Done = false

				if ctxt.HeadType == objabi.Hdarwin {
					r.Done = true
				}

				if ctxt.HeadType != objabi.Hwindows {
					r.Type = objabi.R_ADDR
				}

				r.Xsym = ctxt.Syms.ROLookup(r.Sym.Sect.Name, 0)
				r.Xadd = r.Add + psess.Symaddr(r.Sym) - int64(r.Sym.Sect.Vaddr)

				o = r.Xadd
				if ctxt.IsELF && ctxt.Arch.Family == sys.AMD64 {
					o = 0
				}
				break
			}
			o = psess.Symaddr(r.Sym) + r.Add - int64(r.Sym.Sect.Vaddr)
		case objabi.R_WEAKADDROFF:
			if !r.Sym.Attr.Reachable() {
				continue
			}
			fallthrough
		case objabi.R_ADDROFF:

			if r.Sym.Sect.Name == ".text" {
				o = psess.Symaddr(r.Sym) - int64(psess.Segtext.Sections[0].Vaddr) + r.Add
			} else {
				o = psess.Symaddr(r.Sym) - int64(r.Sym.Sect.Vaddr) + r.Add
			}

		case objabi.R_ADDRCUOFF:

			o = psess.Symaddr(r.Sym) + r.Add - psess.Symaddr(r.Sym.Lib.Textp[0])

		case objabi.R_GOTPCREL:
			if ctxt.DynlinkingGo() && ctxt.HeadType == objabi.Hdarwin && r.Sym != nil && r.Sym.Type != sym.SCONST {
				r.Done = false
				r.Xadd = r.Add
				r.Xadd -= int64(r.Siz)
				r.Xsym = r.Sym

				o = r.Xadd
				o += int64(r.Siz)
				break
			}
			fallthrough
		case objabi.R_CALL, objabi.R_PCREL:
			if ctxt.LinkMode == LinkExternal && r.Sym != nil && r.Sym.Type != sym.SCONST && (r.Sym.Sect != s.Sect || r.Type == objabi.R_GOTPCREL) {
				r.Done = false

				rs := r.Sym

				r.Xadd = r.Add
				for rs.Outer != nil {
					r.Xadd += psess.Symaddr(rs) - psess.Symaddr(rs.Outer)
					rs = rs.Outer
				}

				r.Xadd -= int64(r.Siz)
				if rs.Type != sym.SHOSTOBJ && rs.Type != sym.SDYNIMPORT && rs.Sect == nil {
					psess.
						Errorf(s, "missing section for relocation target %s", rs.Name)
				}
				r.Xsym = rs

				o = r.Xadd
				if ctxt.IsELF {
					if ctxt.Arch.Family == sys.AMD64 {
						o = 0
					}
				} else if ctxt.HeadType == objabi.Hdarwin {
					if r.Type == objabi.R_CALL {
						if ctxt.LinkMode == LinkExternal && rs.Type == sym.SDYNIMPORT {
							switch ctxt.Arch.Family {
							case sys.AMD64:

								o += int64(r.Siz)
							case sys.I386:

								o -= int64(r.Off)
								o -= int64(s.Value - int64(s.Sect.Vaddr))
							}
						} else {
							if rs.Type != sym.SHOSTOBJ {
								o += int64(uint64(psess.Symaddr(rs)) - rs.Sect.Vaddr)
							}
							o -= int64(r.Off)
						}
					} else if ctxt.Arch.Family == sys.ARM {

						o += psess.Symaddr(rs) - s.Value - int64(r.Off)
					} else {
						o += int64(r.Siz)
					}
				} else if ctxt.HeadType == objabi.Hwindows && ctxt.Arch.Family == sys.AMD64 {

					o += int64(r.Siz)
				} else {
					psess.
						Errorf(s, "unhandled pcrel relocation to %s on %v", rs.Name, ctxt.HeadType)
				}

				break
			}

			o = 0
			if r.Sym != nil {
				o += psess.Symaddr(r.Sym)
			}

			o += r.Add - (s.Value + int64(r.Off) + int64(r.Siz))
		case objabi.R_SIZE:
			o = r.Sym.Size + r.Add
		}

		if r.Variant != sym.RV_NONE {
			o = psess.thearch.Archrelocvariant(ctxt, r, s, o)
		}

		if false {
			nam := "<nil>"
			var addr int64
			if r.Sym != nil {
				nam = r.Sym.Name
				addr = psess.Symaddr(r.Sym)
			}
			xnam := "<nil>"
			if r.Xsym != nil {
				xnam = r.Xsym.Name
			}
			fmt.Printf("relocate %s %#x (%#x+%#x, size %d) => %s %#x +%#x (xsym: %s +%#x) [type %d (%s)/%d, %x]\n", s.Name, s.Value+int64(off), s.Value, r.Off, r.Siz, nam, addr, r.Add, xnam, r.Xadd, r.Type, psess.sym.RelocName(ctxt.Arch, r.Type), r.Variant, o)
		}
		switch siz {
		default:
			psess.
				Errorf(s, "bad reloc size %#x for %s", uint32(siz), r.Sym.Name)
			fallthrough

		case 1:
			s.P[off] = byte(int8(o))
		case 2:
			if o != int64(int16(o)) {
				psess.
					Errorf(s, "relocation address for %s is too big: %#x", r.Sym.Name, o)
			}
			i16 := int16(o)
			ctxt.Arch.ByteOrder.PutUint16(s.P[off:], uint16(i16))
		case 4:
			if r.Type == objabi.R_PCREL || r.Type == objabi.R_CALL {
				if o != int64(int32(o)) {
					psess.
						Errorf(s, "pc-relative relocation address for %s is too big: %#x", r.Sym.Name, o)
				}
			} else {
				if o != int64(int32(o)) && o != int64(uint32(o)) {
					psess.
						Errorf(s, "non-pc-relative relocation address for %s is too big: %#x", r.Sym.Name, uint64(o))
				}
			}

			fl := int32(o)
			ctxt.Arch.ByteOrder.PutUint32(s.P[off:], uint32(fl))
		case 8:
			ctxt.Arch.ByteOrder.PutUint64(s.P[off:], uint64(o))
		}
	}
}

func (ctxt *Link) reloc(psess *PackageSession) {
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f reloc\n", psess.Cputime())
	}

	for _, s := range ctxt.Textp {
		psess.
			relocsym(ctxt, s)
	}
	for _, s := range psess.datap {
		psess.
			relocsym(ctxt, s)
	}
	for _, s := range psess.dwarfp {
		psess.
			relocsym(ctxt, s)
	}
}

func (psess *PackageSession) windynrelocsym(ctxt *Link, s *sym.Symbol) {
	rel := ctxt.Syms.Lookup(".rel", 0)
	if s == rel {
		return
	}
	for ri := range s.R {
		r := &s.R[ri]
		targ := r.Sym
		if targ == nil {
			continue
		}
		if !targ.Attr.Reachable() {
			if r.Type == objabi.R_WEAKADDROFF {
				continue
			}
			psess.
				Errorf(s, "dynamic relocation to unreachable symbol %s", targ.Name)
		}
		if r.Sym.Plt == -2 && r.Sym.Got != -2 {
			targ.Plt = int32(rel.Size)
			r.Sym = rel
			r.Add = int64(targ.Plt)

			if ctxt.Arch.Family == sys.I386 {
				rel.AddUint8(0xff)
				rel.AddUint8(0x25)
				rel.AddAddr(ctxt.Arch, targ)
				rel.AddUint8(0x90)
				rel.AddUint8(0x90)
			} else {
				rel.AddUint8(0xff)
				rel.AddUint8(0x24)
				rel.AddUint8(0x25)
				rel.AddAddrPlus4(targ, 0)
				rel.AddUint8(0x90)
			}
		} else if r.Sym.Plt >= 0 {
			r.Sym = rel
			r.Add = int64(targ.Plt)
		}
	}
}

func (psess *PackageSession) dynrelocsym(ctxt *Link, s *sym.Symbol) {
	if ctxt.HeadType == objabi.Hwindows {
		if ctxt.LinkMode == LinkInternal {
			psess.
				windynrelocsym(ctxt, s)
		}
		return
	}

	for ri := range s.R {
		r := &s.R[ri]
		if ctxt.BuildMode == BuildModePIE && ctxt.LinkMode == LinkInternal {
			psess.
				thearch.Adddynrel(ctxt, s, r)
			continue
		}
		if r.Sym != nil && r.Sym.Type == sym.SDYNIMPORT || r.Type >= 256 {
			if r.Sym != nil && !r.Sym.Attr.Reachable() {
				psess.
					Errorf(s, "dynamic relocation to unreachable symbol %s", r.Sym.Name)
			}
			if !psess.thearch.Adddynrel(ctxt, s, r) {
				psess.
					Errorf(s, "unsupported dynamic relocation for symbol %s (type=%d (%s) stype=%d (%s))", r.Sym.Name, r.Type, psess.sym.RelocName(ctxt.Arch, r.Type), r.Sym.Type, r.Sym.Type)
			}
		}
	}
}

func (psess *PackageSession) dynreloc(ctxt *Link, data *[sym.SXREF][]*sym.Symbol) {

	if *psess.FlagD && ctxt.HeadType != objabi.Hwindows {
		return
	}
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f dynreloc\n", psess.Cputime())
	}

	for _, s := range ctxt.Textp {
		psess.
			dynrelocsym(ctxt, s)
	}
	for _, syms := range data {
		for _, s := range syms {
			psess.
				dynrelocsym(ctxt, s)
		}
	}
	if ctxt.IsELF {
		psess.
			elfdynhash(ctxt)
	}
}

func (psess *PackageSession) Codeblk(ctxt *Link, addr int64, size int64) {
	psess.
		CodeblkPad(ctxt, addr, size, psess.zeros[:])
}
func (psess *PackageSession) CodeblkPad(ctxt *Link, addr int64, size int64, pad []byte) {
	if *psess.flagA {
		ctxt.Logf("codeblk [%#x,%#x) at offset %#x\n", addr, addr+size, ctxt.Out.Offset())
	}
	psess.
		blk(ctxt, ctxt.Textp, addr, size, pad)

	if !*psess.flagA {
		return
	}

	syms := ctxt.Textp
	for i, s := range syms {
		if !s.Attr.Reachable() {
			continue
		}
		if s.Value >= addr {
			syms = syms[i:]
			break
		}
	}

	eaddr := addr + size
	for _, s := range syms {
		if !s.Attr.Reachable() {
			continue
		}
		if s.Value >= eaddr {
			break
		}

		if addr < s.Value {
			ctxt.Logf("%-20s %.8x|", "_", uint64(addr))
			for ; addr < s.Value; addr++ {
				ctxt.Logf(" %.2x", 0)
			}
			ctxt.Logf("\n")
		}

		ctxt.Logf("%.6x\t%-20s\n", uint64(addr), s.Name)
		q := s.P

		for len(q) >= 16 {
			ctxt.Logf("%.6x\t% x\n", uint64(addr), q[:16])
			addr += 16
			q = q[16:]
		}

		if len(q) > 0 {
			ctxt.Logf("%.6x\t% x\n", uint64(addr), q)
			addr += int64(len(q))
		}
	}

	if addr < eaddr {
		ctxt.Logf("%-20s %.8x|", "_", uint64(addr))
		for ; addr < eaddr; addr++ {
			ctxt.Logf(" %.2x", 0)
		}
	}
}

func (psess *PackageSession) blk(ctxt *Link, syms []*sym.Symbol, addr, size int64, pad []byte) {
	for i, s := range syms {
		if !s.Attr.SubSymbol() && s.Value >= addr {
			syms = syms[i:]
			break
		}
	}

	eaddr := addr + size
	for _, s := range syms {
		if s.Attr.SubSymbol() {
			continue
		}
		if s.Value >= eaddr {
			break
		}
		if s.Value < addr {
			psess.
				Errorf(s, "phase error: addr=%#x but sym=%#x type=%d", addr, s.Value, s.Type)
			psess.
				errorexit()
		}
		if addr < s.Value {
			ctxt.Out.WriteStringPad("", int(s.Value-addr), pad)
			addr = s.Value
		}
		ctxt.Out.Write(s.P)
		addr += int64(len(s.P))
		if addr < s.Value+s.Size {
			ctxt.Out.WriteStringPad("", int(s.Value+s.Size-addr), pad)
			addr = s.Value + s.Size
		}
		if addr != s.Value+s.Size {
			psess.
				Errorf(s, "phase error: addr=%#x value+size=%#x", addr, s.Value+s.Size)
			psess.
				errorexit()
		}
		if s.Value+s.Size >= eaddr {
			break
		}
	}

	if addr < eaddr {
		ctxt.Out.WriteStringPad("", int(eaddr-addr), pad)
	}
	ctxt.Out.Flush(psess)
}

func (psess *PackageSession) Datblk(ctxt *Link, addr int64, size int64) {
	if *psess.flagA {
		ctxt.Logf("datblk [%#x,%#x) at offset %#x\n", addr, addr+size, ctxt.Out.Offset())
	}
	psess.
		blk(ctxt, psess.datap, addr, size, psess.zeros[:])

	if !*psess.flagA {
		return
	}

	syms := psess.datap
	for i, sym := range syms {
		if sym.Value >= addr {
			syms = syms[i:]
			break
		}
	}

	eaddr := addr + size
	for _, sym := range syms {
		if sym.Value >= eaddr {
			break
		}
		if addr < sym.Value {
			ctxt.Logf("\t%.8x| 00 ...\n", uint64(addr))
			addr = sym.Value
		}

		ctxt.Logf("%s\n\t%.8x|", sym.Name, uint64(addr))
		for i, b := range sym.P {
			if i > 0 && i%16 == 0 {
				ctxt.Logf("\n\t%.8x|", uint64(addr)+uint64(i))
			}
			ctxt.Logf(" %.2x", b)
		}

		addr += int64(len(sym.P))
		for ; addr < sym.Value+sym.Size; addr++ {
			ctxt.Logf(" %.2x", 0)
		}
		ctxt.Logf("\n")

		if ctxt.LinkMode != LinkExternal {
			continue
		}
		for _, r := range sym.R {
			rsname := ""
			if r.Sym != nil {
				rsname = r.Sym.Name
			}
			typ := "?"
			switch r.Type {
			case objabi.R_ADDR:
				typ = "addr"
			case objabi.R_PCREL:
				typ = "pcrel"
			case objabi.R_CALL:
				typ = "call"
			}
			ctxt.Logf("\treloc %.8x/%d %s %s+%#x [%#x]\n", uint(sym.Value+int64(r.Off)), r.Siz, typ, rsname, r.Add, r.Sym.Value+r.Add)
		}
	}

	if addr < eaddr {
		ctxt.Logf("\t%.8x| 00 ...\n", uint(addr))
	}
	ctxt.Logf("\t%.8x|\n", uint(eaddr))
}

func (psess *PackageSession) Dwarfblk(ctxt *Link, addr int64, size int64) {
	if *psess.flagA {
		ctxt.Logf("dwarfblk [%#x,%#x) at offset %#x\n", addr, addr+size, ctxt.Out.Offset())
	}
	psess.
		blk(ctxt, psess.dwarfp, addr, size, psess.zeros[:])
}

func (psess *PackageSession) addstrdata1(ctxt *Link, arg string) {
	eq := strings.Index(arg, "=")
	dot := strings.LastIndex(arg[:eq+1], ".")
	if eq < 0 || dot < 0 {
		psess.
			Exitf("-X flag requires argument of the form importpath.name=value")
	}
	pkg := arg[:dot]
	if ctxt.BuildMode == BuildModePlugin && pkg == "main" {
		pkg = *psess.flagPluginPath
	}
	pkg = objabi.PathToPrefix(pkg)
	name := pkg + arg[dot:eq]
	value := arg[eq+1:]
	if _, ok := psess.strdata[name]; !ok {
		psess.
			strnames = append(psess.strnames, name)
	}
	psess.
		strdata[name] = value
}

func (psess *PackageSession) addstrdata(ctxt *Link, name, value string) {
	s := ctxt.Syms.ROLookup(name, 0)
	if s == nil || s.Gotype == nil {

		return
	}
	if s.Gotype.Name != "type.string" {
		psess.
			Errorf(s, "cannot set with -X: not a var of type string (%s)", s.Gotype.Name)
		return
	}
	if s.Type == sym.SBSS {
		s.Type = sym.SDATA
	}

	p := fmt.Sprintf("%s.str", s.Name)
	sp := ctxt.Syms.Lookup(p, 0)
	psess.
		Addstring(sp, value)
	sp.Type = sym.SRODATA

	s.Size = 0
	s.P = s.P[:0]
	s.R = s.R[:0]
	reachable := s.Attr.Reachable()
	s.AddAddr(ctxt.Arch, sp)
	s.AddUint(ctxt.Arch, uint64(len(value)))

	s.Attr.Set(sym.AttrReachable, reachable)

	sp.Attr.Set(sym.AttrReachable, reachable)
}

func (ctxt *Link) dostrdata(psess *PackageSession) {
	for _, name := range psess.strnames {
		psess.
			addstrdata(ctxt, name, psess.strdata[name])
	}
}

func (psess *PackageSession) Addstring(s *sym.Symbol, str string) int64 {
	if s.Type == 0 {
		s.Type = sym.SNOPTRDATA
	}
	s.Attr |= sym.AttrReachable
	r := s.Size
	if s.Name == ".shstrtab" {
		psess.
			elfsetstring(s, str, int(r))
	}
	s.P = append(s.P, str...)
	s.P = append(s.P, 0)
	s.Size = int64(len(s.P))
	return r
}

// addgostring adds str, as a Go string value, to s. symname is the name of the
// symbol used to define the string data and must be unique per linked object.
func (psess *PackageSession) addgostring(ctxt *Link, s *sym.Symbol, symname, str string) {
	sdata := ctxt.Syms.Lookup(symname, 0)
	if sdata.Type != sym.Sxxx {
		psess.
			Errorf(s, "duplicate symname in addgostring: %s", symname)
	}
	sdata.Attr |= sym.AttrReachable
	sdata.Attr |= sym.AttrLocal
	sdata.Type = sym.SRODATA
	sdata.Size = int64(len(str))
	sdata.P = []byte(str)
	s.AddAddr(ctxt.Arch, sdata)
	s.AddUint(ctxt.Arch, uint64(len(str)))
}

func addinitarrdata(ctxt *Link, s *sym.Symbol) {
	p := s.Name + ".ptr"
	sp := ctxt.Syms.Lookup(p, 0)
	sp.Type = sym.SINITARR
	sp.Size = 0
	sp.Attr |= sym.AttrDuplicateOK
	sp.AddAddr(ctxt.Arch, s)
}

func (psess *PackageSession) dosymtype(ctxt *Link) {
	switch ctxt.BuildMode {
	case BuildModeCArchive, BuildModeCShared:
		for _, s := range ctxt.Syms.Allsym {

			if s.Name == *psess.flagEntrySymbol {
				addinitarrdata(ctxt, s)
			}
		}
	}
}

// symalign returns the required alignment for the given symbol s.
func (psess *PackageSession) symalign(s *sym.Symbol) int32 {
	min := int32(psess.thearch.Minalign)
	if s.Align >= min {
		return s.Align
	} else if s.Align != 0 {
		return min
	}
	if strings.HasPrefix(s.Name, "go.string.") || strings.HasPrefix(s.Name, "type..namedata.") {

		return min
	}
	align := int32(psess.thearch.Maxalign)
	for int64(align) > s.Size && align > min {
		align >>= 1
	}
	return align
}

func (psess *PackageSession) aligndatsize(datsize int64, s *sym.Symbol) int64 {
	return Rnd(datsize, int64(psess.symalign(s)))
}

const debugGCProg = false

type GCProg struct {
	ctxt *Link
	sym  *sym.Symbol
	w    gcprog.Writer
}

func (p *GCProg) Init(ctxt *Link, name string) {
	p.ctxt = ctxt
	p.sym = ctxt.Syms.Lookup(name, 0)
	p.w.Init(p.writeByte(ctxt))
	if debugGCProg {
		fmt.Fprintf(os.Stderr, "ld: start GCProg %s\n", name)
		p.w.Debug(os.Stderr)
	}
}

func (p *GCProg) writeByte(ctxt *Link) func(x byte) {
	return func(x byte) {
		p.sym.AddUint8(x)
	}
}

func (p *GCProg) End(size int64) {
	p.w.ZeroUntil(size / int64(p.ctxt.Arch.PtrSize))
	p.w.End()
	if debugGCProg {
		fmt.Fprintf(os.Stderr, "ld: end GCProg\n")
	}
}

func (p *GCProg) AddSym(psess *PackageSession, s *sym.Symbol) {
	typ := s.Gotype

	if typ == nil {
		switch s.Name {
		case "runtime.data", "runtime.edata", "runtime.bss", "runtime.ebss":

			return
		}
		psess.
			Errorf(s, "missing Go type information for global symbol: size %d", s.Size)
		return
	}

	ptrsize := int64(p.ctxt.Arch.PtrSize)
	nptr := psess.decodetypePtrdata(p.ctxt.Arch, typ) / ptrsize

	if debugGCProg {
		fmt.Fprintf(os.Stderr, "gcprog sym: %s at %d (ptr=%d+%d)\n", s.Name, s.Value, s.Value/ptrsize, nptr)
	}

	if decodetypeUsegcprog(p.ctxt.Arch, typ) == 0 {

		mask := psess.decodetypeGcmask(p.ctxt, typ)
		for i := int64(0); i < nptr; i++ {
			if (mask[i/8]>>uint(i%8))&1 != 0 {
				p.w.Ptr(s.Value/ptrsize + i)
			}
		}
		return
	}

	prog := psess.decodetypeGcprog(p.ctxt, typ)
	p.w.ZeroUntil(s.Value / ptrsize)
	p.w.Append(prog[4:], nptr)
}

// dataSortKey is used to sort a slice of data symbol *sym.Symbol pointers.
// The sort keys are kept inline to improve cache behavior while sorting.
type dataSortKey struct {
	size int64
	name string
	sym  *sym.Symbol
}

type bySizeAndName []dataSortKey

func (d bySizeAndName) Len() int      { return len(d) }
func (d bySizeAndName) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d bySizeAndName) Less(i, j int) bool {
	s1, s2 := d[i], d[j]
	if s1.size != s2.size {
		return s1.size < s2.size
	}
	return s1.name < s2.name
}

// cutoff is the maximum data section size permitted by the linker
// (see issue #9862).
const cutoff = 2e9 // 2 GB (or so; looks better in errors than 2^31)

func (psess *PackageSession) checkdatsize(ctxt *Link, datsize int64, symn sym.SymKind) {
	if datsize > cutoff {
		psess.
			Errorf(nil, "too much data in section %v (over %v bytes)", symn, cutoff)
	}
}

// datap is a collection of reachable data symbols in address order.
// Generated by dodata.

func (ctxt *Link) dodata(psess *PackageSession) {
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f dodata\n", psess.Cputime())
	}

	if ctxt.DynlinkingGo() && ctxt.HeadType == objabi.Hdarwin {

		bss := ctxt.Syms.Lookup("runtime.bss", 0)
		bss.Size = 8
		bss.Attr.Set(sym.AttrSpecial, false)

		ctxt.Syms.Lookup("runtime.ebss", 0).Attr.Set(sym.AttrSpecial, false)

		data := ctxt.Syms.Lookup("runtime.data", 0)
		data.Size = 8
		data.Attr.Set(sym.AttrSpecial, false)

		ctxt.Syms.Lookup("runtime.edata", 0).Attr.Set(sym.AttrSpecial, false)

		types := ctxt.Syms.Lookup("runtime.types", 0)
		types.Type = sym.STYPE
		types.Size = 8
		types.Attr.Set(sym.AttrSpecial, false)

		etypes := ctxt.Syms.Lookup("runtime.etypes", 0)
		etypes.Type = sym.SFUNCTAB
		etypes.Attr.Set(sym.AttrSpecial, false)
	}

	// Collect data symbols by type into data.
	var data [sym.SXREF][]*sym.Symbol
	for _, s := range ctxt.Syms.Allsym {
		if !s.Attr.Reachable() || s.Attr.Special() || s.Attr.SubSymbol() {
			continue
		}
		if s.Type <= sym.STEXT || s.Type >= sym.SXREF {
			continue
		}
		data[s.Type] = append(data[s.Type], s)
	}

	if ctxt.HeadType == objabi.Hdarwin {
		psess.
			machosymorder(ctxt)
	}
	psess.
		dynreloc(ctxt, &data)

	if ctxt.UseRelro() {

		for _, symnro := range psess.sym.ReadOnly {
			symnrelro := psess.sym.RelROMap[symnro]

			ro := []*sym.Symbol{}
			relro := data[symnrelro]

			for _, s := range data[symnro] {
				isRelro := len(s.R) > 0
				switch s.Type {
				case sym.STYPE, sym.STYPERELRO, sym.SGOFUNCRELRO:

					isRelro = true
				}
				if isRelro {
					s.Type = symnrelro
					if s.Outer != nil {
						s.Outer.Type = s.Type
					}
					relro = append(relro, s)
				} else {
					ro = append(ro, s)
				}
			}

			for _, s := range relro {
				if s.Outer != nil && s.Outer.Type != s.Type {
					psess.
						Errorf(s, "inconsistent types for symbol and its Outer %s (%v != %v)",
							s.Outer.Name, s.Type, s.Outer.Type)
				}
			}

			data[symnro] = ro
			data[symnrelro] = relro
		}
	}

	// Sort symbols.
	var dataMaxAlign [sym.SXREF]int32
	var wg sync.WaitGroup
	for symn := range data {
		symn := sym.SymKind(symn)
		wg.Add(1)
		go func() {
			data[symn], dataMaxAlign[symn] = psess.dodataSect(ctxt, symn, data[symn])
			wg.Done()
		}()
	}
	wg.Wait()

	datsize := int64(0)

	writable := []sym.SymKind{
		sym.SELFSECT,
		sym.SMACHO,
		sym.SMACHOGOT,
		sym.SWINDOWS,
	}
	for _, symn := range writable {
		for _, s := range data[symn] {
			sect := addsection(ctxt.Arch, &psess.Segdata, s.Name, 06)
			sect.Align = psess.symalign(s)
			datsize = Rnd(datsize, int64(sect.Align))
			sect.Vaddr = uint64(datsize)
			s.Sect = sect
			s.Type = sym.SDATA
			s.Value = int64(uint64(datsize) - sect.Vaddr)
			datsize += s.Size
			sect.Length = uint64(datsize) - sect.Vaddr
		}
		psess.
			checkdatsize(ctxt, datsize, symn)
	}

	if len(data[sym.SELFGOT]) > 0 {
		sect := addsection(ctxt.Arch, &psess.Segdata, ".got", 06)
		sect.Align = dataMaxAlign[sym.SELFGOT]
		datsize = Rnd(datsize, int64(sect.Align))
		sect.Vaddr = uint64(datsize)
		for _, s := range data[sym.SELFGOT] {
			datsize = psess.aligndatsize(datsize, s)
			s.Sect = sect
			s.Type = sym.SDATA
			s.Value = int64(uint64(datsize) - sect.Vaddr)

			toc := ctxt.Syms.ROLookup(".TOC.", int(s.Version))
			if toc != nil {
				toc.Sect = sect
				toc.Outer = s
				toc.Sub = s.Sub
				s.Sub = toc

				toc.Value = 0x8000
			}

			datsize += s.Size
		}
		psess.
			checkdatsize(ctxt, datsize, sym.SELFGOT)
		sect.Length = uint64(datsize) - sect.Vaddr
	}

	sect := addsection(ctxt.Arch, &psess.Segdata, ".noptrdata", 06)
	sect.Align = dataMaxAlign[sym.SNOPTRDATA]
	datsize = Rnd(datsize, int64(sect.Align))
	sect.Vaddr = uint64(datsize)
	ctxt.Syms.Lookup("runtime.noptrdata", 0).Sect = sect
	ctxt.Syms.Lookup("runtime.enoptrdata", 0).Sect = sect
	for _, s := range data[sym.SNOPTRDATA] {
		datsize = psess.aligndatsize(datsize, s)
		s.Sect = sect
		s.Type = sym.SDATA
		s.Value = int64(uint64(datsize) - sect.Vaddr)
		datsize += s.Size
	}
	psess.
		checkdatsize(ctxt, datsize, sym.SNOPTRDATA)
	sect.Length = uint64(datsize) - sect.Vaddr

	hasinitarr := ctxt.linkShared

	switch ctxt.BuildMode {
	case BuildModeCArchive, BuildModeCShared, BuildModeShared, BuildModePlugin:
		hasinitarr = true
	}
	if hasinitarr {
		sect := addsection(ctxt.Arch, &psess.Segdata, ".init_array", 06)
		sect.Align = dataMaxAlign[sym.SINITARR]
		datsize = Rnd(datsize, int64(sect.Align))
		sect.Vaddr = uint64(datsize)
		for _, s := range data[sym.SINITARR] {
			datsize = psess.aligndatsize(datsize, s)
			s.Sect = sect
			s.Value = int64(uint64(datsize) - sect.Vaddr)
			datsize += s.Size
		}
		sect.Length = uint64(datsize) - sect.Vaddr
		psess.
			checkdatsize(ctxt, datsize, sym.SINITARR)
	}

	sect = addsection(ctxt.Arch, &psess.Segdata, ".data", 06)
	sect.Align = dataMaxAlign[sym.SDATA]
	datsize = Rnd(datsize, int64(sect.Align))
	sect.Vaddr = uint64(datsize)
	ctxt.Syms.Lookup("runtime.data", 0).Sect = sect
	ctxt.Syms.Lookup("runtime.edata", 0).Sect = sect
	var gc GCProg
	gc.Init(ctxt, "runtime.gcdata")
	for _, s := range data[sym.SDATA] {
		s.Sect = sect
		s.Type = sym.SDATA
		datsize = psess.aligndatsize(datsize, s)
		s.Value = int64(uint64(datsize) - sect.Vaddr)
		gc.AddSym(psess, s)
		datsize += s.Size
	}
	psess.
		checkdatsize(ctxt, datsize, sym.SDATA)
	sect.Length = uint64(datsize) - sect.Vaddr
	gc.End(int64(sect.Length))

	sect = addsection(ctxt.Arch, &psess.Segdata, ".bss", 06)
	sect.Align = dataMaxAlign[sym.SBSS]
	datsize = Rnd(datsize, int64(sect.Align))
	sect.Vaddr = uint64(datsize)
	ctxt.Syms.Lookup("runtime.bss", 0).Sect = sect
	ctxt.Syms.Lookup("runtime.ebss", 0).Sect = sect
	gc = GCProg{}
	gc.Init(ctxt, "runtime.gcbss")
	for _, s := range data[sym.SBSS] {
		s.Sect = sect
		datsize = psess.aligndatsize(datsize, s)
		s.Value = int64(uint64(datsize) - sect.Vaddr)
		gc.AddSym(psess, s)
		datsize += s.Size
	}
	psess.
		checkdatsize(ctxt, datsize, sym.SBSS)
	sect.Length = uint64(datsize) - sect.Vaddr
	gc.End(int64(sect.Length))

	sect = addsection(ctxt.Arch, &psess.Segdata, ".noptrbss", 06)
	sect.Align = dataMaxAlign[sym.SNOPTRBSS]
	datsize = Rnd(datsize, int64(sect.Align))
	sect.Vaddr = uint64(datsize)
	ctxt.Syms.Lookup("runtime.noptrbss", 0).Sect = sect
	ctxt.Syms.Lookup("runtime.enoptrbss", 0).Sect = sect
	for _, s := range data[sym.SNOPTRBSS] {
		datsize = psess.aligndatsize(datsize, s)
		s.Sect = sect
		s.Value = int64(uint64(datsize) - sect.Vaddr)
		datsize += s.Size
	}

	sect.Length = uint64(datsize) - sect.Vaddr
	ctxt.Syms.Lookup("runtime.end", 0).Sect = sect
	psess.
		checkdatsize(ctxt, datsize, sym.SNOPTRBSS)

	if len(data[sym.STLSBSS]) > 0 {
		var sect *sym.Section
		if ctxt.IsELF && (ctxt.LinkMode == LinkExternal || !*psess.FlagD) {
			sect = addsection(ctxt.Arch, &psess.Segdata, ".tbss", 06)
			sect.Align = int32(ctxt.Arch.PtrSize)
			sect.Vaddr = 0
		}
		datsize = 0

		for _, s := range data[sym.STLSBSS] {
			datsize = psess.aligndatsize(datsize, s)
			s.Sect = sect
			s.Value = datsize
			datsize += s.Size
		}
		psess.
			checkdatsize(ctxt, datsize, sym.STLSBSS)

		if sect != nil {
			sect.Length = uint64(datsize)
		}
	}

	/*
	 * We finished data, begin read-only data.
	 * Not all systems support a separate read-only non-executable data section.
	 * ELF and Windows PE systems do.
	 * OS X and Plan 9 do not.
	 * And if we're using external linking mode, the point is moot,
	 * since it's not our decision; that code expects the sections in
	 * segtext.
	 */
	var segro *sym.Segment
	if ctxt.IsELF && ctxt.LinkMode == LinkInternal {
		segro = &psess.Segrodata
	} else if ctxt.HeadType == objabi.Hwindows {
		segro = &psess.Segrodata
	} else {
		segro = &psess.Segtext
	}

	datsize = 0

	if len(data[sym.STEXT]) != 0 {
		psess.
			Errorf(nil, "dodata found an sym.STEXT symbol: %s", data[sym.STEXT][0].Name)
	}
	for _, s := range data[sym.SELFRXSECT] {
		sect := addsection(ctxt.Arch, &psess.Segtext, s.Name, 04)
		sect.Align = psess.symalign(s)
		datsize = Rnd(datsize, int64(sect.Align))
		sect.Vaddr = uint64(datsize)
		s.Sect = sect
		s.Type = sym.SRODATA
		s.Value = int64(uint64(datsize) - sect.Vaddr)
		datsize += s.Size
		sect.Length = uint64(datsize) - sect.Vaddr
		psess.
			checkdatsize(ctxt, datsize, sym.SELFRXSECT)
	}

	sect = addsection(ctxt.Arch, segro, ".rodata", 04)

	sect.Vaddr = 0
	ctxt.Syms.Lookup("runtime.rodata", 0).Sect = sect
	ctxt.Syms.Lookup("runtime.erodata", 0).Sect = sect
	if !ctxt.UseRelro() {
		ctxt.Syms.Lookup("runtime.types", 0).Sect = sect
		ctxt.Syms.Lookup("runtime.etypes", 0).Sect = sect
	}
	for _, symn := range psess.sym.ReadOnly {
		align := dataMaxAlign[symn]
		if sect.Align < align {
			sect.Align = align
		}
	}
	datsize = Rnd(datsize, int64(sect.Align))
	for _, symn := range psess.sym.ReadOnly {
		for _, s := range data[symn] {
			datsize = psess.aligndatsize(datsize, s)
			s.Sect = sect
			s.Type = sym.SRODATA
			s.Value = int64(uint64(datsize) - sect.Vaddr)
			datsize += s.Size
		}
		psess.
			checkdatsize(ctxt, datsize, symn)
	}
	sect.Length = uint64(datsize) - sect.Vaddr

	for _, s := range data[sym.SELFROSECT] {
		sect = addsection(ctxt.Arch, segro, s.Name, 04)
		sect.Align = psess.symalign(s)
		datsize = Rnd(datsize, int64(sect.Align))
		sect.Vaddr = uint64(datsize)
		s.Sect = sect
		s.Type = sym.SRODATA
		s.Value = int64(uint64(datsize) - sect.Vaddr)
		datsize += s.Size
		sect.Length = uint64(datsize) - sect.Vaddr
	}
	psess.
		checkdatsize(ctxt, datsize, sym.SELFROSECT)

	for _, s := range data[sym.SMACHOPLT] {
		sect = addsection(ctxt.Arch, segro, s.Name, 04)
		sect.Align = psess.symalign(s)
		datsize = Rnd(datsize, int64(sect.Align))
		sect.Vaddr = uint64(datsize)
		s.Sect = sect
		s.Type = sym.SRODATA
		s.Value = int64(uint64(datsize) - sect.Vaddr)
		datsize += s.Size
		sect.Length = uint64(datsize) - sect.Vaddr
	}
	psess.
		checkdatsize(ctxt, datsize, sym.SMACHOPLT)

	addrelrosection := func(suffix string) *sym.Section {
		return addsection(ctxt.Arch, segro, suffix, 04)
	}

	if ctxt.UseRelro() {
		addrelrosection = func(suffix string) *sym.Section {
			seg := &psess.Segrelrodata
			if ctxt.LinkMode == LinkExternal {

				seg = &psess.Segrodata
			}
			return addsection(ctxt.Arch, seg, ".data.rel.ro"+suffix, 06)
		}

		sect = addrelrosection("")

		sect.Vaddr = 0
		ctxt.Syms.Lookup("runtime.types", 0).Sect = sect
		ctxt.Syms.Lookup("runtime.etypes", 0).Sect = sect
		for _, symnro := range psess.sym.ReadOnly {
			symn := psess.sym.RelROMap[symnro]
			align := dataMaxAlign[symn]
			if sect.Align < align {
				sect.Align = align
			}
		}
		datsize = Rnd(datsize, int64(sect.Align))
		for _, symnro := range psess.sym.ReadOnly {
			symn := psess.sym.RelROMap[symnro]
			for _, s := range data[symn] {
				datsize = psess.aligndatsize(datsize, s)
				if s.Outer != nil && s.Outer.Sect != nil && s.Outer.Sect != sect {
					psess.
						Errorf(s, "s.Outer (%s) in different section from s, %s != %s", s.Outer.Name, s.Outer.Sect.Name, sect.Name)
				}
				s.Sect = sect
				s.Type = sym.SRODATA
				s.Value = int64(uint64(datsize) - sect.Vaddr)
				datsize += s.Size
			}
			psess.
				checkdatsize(ctxt, datsize, symn)
		}

		sect.Length = uint64(datsize) - sect.Vaddr
	}

	sect = addrelrosection(".typelink")
	sect.Align = dataMaxAlign[sym.STYPELINK]
	datsize = Rnd(datsize, int64(sect.Align))
	sect.Vaddr = uint64(datsize)
	typelink := ctxt.Syms.Lookup("runtime.typelink", 0)
	typelink.Sect = sect
	typelink.Type = sym.SRODATA
	datsize += typelink.Size
	psess.
		checkdatsize(ctxt, datsize, sym.STYPELINK)
	sect.Length = uint64(datsize) - sect.Vaddr

	sect = addrelrosection(".itablink")
	sect.Align = dataMaxAlign[sym.SITABLINK]
	datsize = Rnd(datsize, int64(sect.Align))
	sect.Vaddr = uint64(datsize)
	ctxt.Syms.Lookup("runtime.itablink", 0).Sect = sect
	ctxt.Syms.Lookup("runtime.eitablink", 0).Sect = sect
	for _, s := range data[sym.SITABLINK] {
		datsize = psess.aligndatsize(datsize, s)
		s.Sect = sect
		s.Type = sym.SRODATA
		s.Value = int64(uint64(datsize) - sect.Vaddr)
		datsize += s.Size
	}
	psess.
		checkdatsize(ctxt, datsize, sym.SITABLINK)
	sect.Length = uint64(datsize) - sect.Vaddr

	sect = addrelrosection(".gosymtab")
	sect.Align = dataMaxAlign[sym.SSYMTAB]
	datsize = Rnd(datsize, int64(sect.Align))
	sect.Vaddr = uint64(datsize)
	ctxt.Syms.Lookup("runtime.symtab", 0).Sect = sect
	ctxt.Syms.Lookup("runtime.esymtab", 0).Sect = sect
	for _, s := range data[sym.SSYMTAB] {
		datsize = psess.aligndatsize(datsize, s)
		s.Sect = sect
		s.Type = sym.SRODATA
		s.Value = int64(uint64(datsize) - sect.Vaddr)
		datsize += s.Size
	}
	psess.
		checkdatsize(ctxt, datsize, sym.SSYMTAB)
	sect.Length = uint64(datsize) - sect.Vaddr

	sect = addrelrosection(".gopclntab")
	sect.Align = dataMaxAlign[sym.SPCLNTAB]
	datsize = Rnd(datsize, int64(sect.Align))
	sect.Vaddr = uint64(datsize)
	ctxt.Syms.Lookup("runtime.pclntab", 0).Sect = sect
	ctxt.Syms.Lookup("runtime.epclntab", 0).Sect = sect
	for _, s := range data[sym.SPCLNTAB] {
		datsize = psess.aligndatsize(datsize, s)
		s.Sect = sect
		s.Type = sym.SRODATA
		s.Value = int64(uint64(datsize) - sect.Vaddr)
		datsize += s.Size
	}
	psess.
		checkdatsize(ctxt, datsize, sym.SRODATA)
	sect.Length = uint64(datsize) - sect.Vaddr

	if datsize != int64(uint32(datsize)) {
		psess.
			Errorf(nil, "read-only data segment too large: %d", datsize)
	}

	for symn := sym.SELFRXSECT; symn < sym.SXREF; symn++ {
		psess.
			datap = append(psess.datap, data[symn]...)
	}
	psess.
		dwarfgeneratedebugsyms(ctxt)

	var i int
	for ; i < len(psess.dwarfp); i++ {
		s := psess.dwarfp[i]
		if s.Type != sym.SDWARFSECT {
			break
		}

		sect = addsection(ctxt.Arch, &psess.Segdwarf, s.Name, 04)
		sect.Align = 1
		datsize = Rnd(datsize, int64(sect.Align))
		sect.Vaddr = uint64(datsize)
		s.Sect = sect
		s.Type = sym.SRODATA
		s.Value = int64(uint64(datsize) - sect.Vaddr)
		datsize += s.Size
		sect.Length = uint64(datsize) - sect.Vaddr
	}
	psess.
		checkdatsize(ctxt, datsize, sym.SDWARFSECT)

	for i < len(psess.dwarfp) {
		curType := psess.dwarfp[i].Type
		var sect *sym.Section
		switch curType {
		case sym.SDWARFINFO:
			sect = addsection(ctxt.Arch, &psess.Segdwarf, ".debug_info", 04)
		case sym.SDWARFRANGE:
			sect = addsection(ctxt.Arch, &psess.Segdwarf, ".debug_ranges", 04)
		case sym.SDWARFLOC:
			sect = addsection(ctxt.Arch, &psess.Segdwarf, ".debug_loc", 04)
		default:
			psess.
				Errorf(psess.dwarfp[i], "unknown DWARF section %v", curType)
		}

		sect.Align = 1
		datsize = Rnd(datsize, int64(sect.Align))
		sect.Vaddr = uint64(datsize)
		for ; i < len(psess.dwarfp); i++ {
			s := psess.dwarfp[i]
			if s.Type != curType {
				break
			}
			s.Sect = sect
			s.Type = sym.SRODATA
			s.Value = int64(uint64(datsize) - sect.Vaddr)
			s.Attr |= sym.AttrLocal
			datsize += s.Size
		}
		sect.Length = uint64(datsize) - sect.Vaddr
		psess.
			checkdatsize(ctxt, datsize, curType)
	}

	n := int32(1)

	for _, sect := range psess.Segtext.Sections {
		sect.Extnum = int16(n)
		n++
	}
	for _, sect := range psess.Segrodata.Sections {
		sect.Extnum = int16(n)
		n++
	}
	for _, sect := range psess.Segrelrodata.Sections {
		sect.Extnum = int16(n)
		n++
	}
	for _, sect := range psess.Segdata.Sections {
		sect.Extnum = int16(n)
		n++
	}
	for _, sect := range psess.Segdwarf.Sections {
		sect.Extnum = int16(n)
		n++
	}
}

func (psess *PackageSession) dodataSect(ctxt *Link, symn sym.SymKind, syms []*sym.Symbol) (result []*sym.Symbol, maxAlign int32) {
	if ctxt.HeadType == objabi.Hdarwin {

		newSyms := make([]*sym.Symbol, 0, len(syms))
		for _, s := range syms {
			if s.Type == symn {
				newSyms = append(newSyms, s)
			}
		}
		syms = newSyms
	}

	var head, tail *sym.Symbol
	symsSort := make([]dataSortKey, 0, len(syms))
	for _, s := range syms {
		if s.Attr.OnList() {
			log.Fatalf("symbol %s listed multiple times", s.Name)
		}
		s.Attr |= sym.AttrOnList
		switch {
		case s.Size < int64(len(s.P)):
			psess.
				Errorf(s, "initialize bounds (%d < %d)", s.Size, len(s.P))
		case s.Size < 0:
			psess.
				Errorf(s, "negative size (%d bytes)", s.Size)
		case s.Size > cutoff:
			psess.
				Errorf(s, "symbol too large (%d bytes)", s.Size)
		}

		if ctxt.DynlinkingGo() && ctxt.HeadType == objabi.Hdarwin {
			switch s.Name {
			case "runtime.text", "runtime.bss", "runtime.data", "runtime.types":
				head = s
				continue
			case "runtime.etext", "runtime.ebss", "runtime.edata", "runtime.etypes":
				tail = s
				continue
			}
		}

		key := dataSortKey{
			size: s.Size,
			name: s.Name,
			sym:  s,
		}

		switch s.Type {
		case sym.SELFGOT:

			key.size = 0
		}

		symsSort = append(symsSort, key)
	}

	sort.Sort(bySizeAndName(symsSort))

	off := 0
	if head != nil {
		syms[0] = head
		off++
	}
	for i, symSort := range symsSort {
		syms[i+off] = symSort.sym
		align := psess.symalign(symSort.sym)
		if maxAlign < align {
			maxAlign = align
		}
	}
	if tail != nil {
		syms[len(syms)-1] = tail
	}

	if ctxt.IsELF && symn == sym.SELFROSECT {

		reli, plti := -1, -1
		for i, s := range syms {
			switch s.Name {
			case ".rel.plt", ".rela.plt":
				plti = i
			case ".rel", ".rela":
				reli = i
			}
		}
		if reli >= 0 && plti >= 0 && plti != reli+1 {
			var first, second int
			if plti > reli {
				first, second = reli, plti
			} else {
				first, second = plti, reli
			}
			rel, plt := syms[reli], syms[plti]
			copy(syms[first+2:], syms[first+1:second])
			syms[first+0] = rel
			syms[first+1] = plt

			rel.Align = int32(ctxt.Arch.RegSize)
			plt.Align = int32(ctxt.Arch.RegSize)
		}
	}

	return syms, maxAlign
}

// Add buildid to beginning of text segment, on non-ELF systems.
// Non-ELF binary formats are not always flexible enough to
// give us a place to put the Go build ID. On those systems, we put it
// at the very beginning of the text segment.
// This ``header'' is read by cmd/go.
func (ctxt *Link) textbuildid(psess *PackageSession) {
	if ctxt.IsELF || ctxt.BuildMode == BuildModePlugin || *psess.flagBuildid == "" {
		return
	}

	s := ctxt.Syms.Lookup("go.buildid", 0)
	s.Attr |= sym.AttrReachable

	data := "\xff Go build ID: " + strconv.Quote(*psess.flagBuildid) + "\n \xff"
	s.Type = sym.STEXT
	s.P = []byte(data)
	s.Size = int64(len(s.P))

	ctxt.Textp = append(ctxt.Textp, nil)
	copy(ctxt.Textp[1:], ctxt.Textp)
	ctxt.Textp[0] = s
}

// assign addresses to text
func (ctxt *Link) textaddress(psess *PackageSession) {
	addsection(ctxt.Arch, &psess.Segtext, ".text", 05)

	sect := psess.Segtext.Sections[0]

	sect.Align = int32(psess.Funcalign)

	text := ctxt.Syms.Lookup("runtime.text", 0)
	text.Sect = sect

	if ctxt.DynlinkingGo() && ctxt.HeadType == objabi.Hdarwin {
		etext := ctxt.Syms.Lookup("runtime.etext", 0)
		etext.Sect = sect

		ctxt.Textp = append(ctxt.Textp, etext, nil)
		copy(ctxt.Textp[1:], ctxt.Textp)
		ctxt.Textp[0] = text
	}

	va := uint64(*psess.FlagTextAddr)
	n := 1
	sect.Vaddr = va
	ntramps := 0
	for _, s := range ctxt.Textp {
		sect, n, va = psess.assignAddress(ctxt, sect, n, s, va, false)
		psess.
			trampoline(ctxt, s)

		for ; ntramps < len(ctxt.tramps); ntramps++ {
			tramp := ctxt.tramps[ntramps]
			sect, n, va = psess.assignAddress(ctxt, sect, n, tramp, va, true)
		}
	}

	sect.Length = va - sect.Vaddr
	ctxt.Syms.Lookup("runtime.etext", 0).Sect = sect

	if ntramps != 0 {
		newtextp := make([]*sym.Symbol, 0, len(ctxt.Textp)+ntramps)
		i := 0
		for _, s := range ctxt.Textp {
			for ; i < ntramps && ctxt.tramps[i].Value < s.Value; i++ {
				newtextp = append(newtextp, ctxt.tramps[i])
			}
			newtextp = append(newtextp, s)
		}
		newtextp = append(newtextp, ctxt.tramps[i:ntramps]...)

		ctxt.Textp = newtextp
	}
}

// assigns address for a text symbol, returns (possibly new) section, its number, and the address
// Note: once we have trampoline insertion support for external linking, this function
// will not need to create new text sections, and so no need to return sect and n.
func (psess *PackageSession) assignAddress(ctxt *Link, sect *sym.Section, n int, s *sym.Symbol, va uint64, isTramp bool) (*sym.Section, int, uint64) {
	if psess.thearch.AssignAddress != nil {
		return psess.thearch.AssignAddress(ctxt, sect, n, s, va, isTramp)
	}

	s.Sect = sect
	if s.Attr.SubSymbol() {
		return sect, n, va
	}
	if s.Align != 0 {
		va = uint64(Rnd(int64(va), int64(s.Align)))
	} else {
		va = uint64(Rnd(int64(va), int64(psess.Funcalign)))
	}
	s.Value = 0
	for sub := s; sub != nil; sub = sub.Sub {
		sub.Value += int64(va)
	}

	funcsize := uint64(MINFUNC)
	if s.Size > MINFUNC {
		funcsize = uint64(s.Size)
	}

	if ctxt.Arch.InFamily(sys.PPC64) && s.Outer == nil && ctxt.IsELF && ctxt.LinkMode == LinkExternal && va-sect.Vaddr+funcsize+psess.maxSizeTrampolinesPPC64(s, isTramp) > 0x1c00000 {

		sect.Length = va - sect.Vaddr

		sect = addsection(ctxt.Arch, &psess.Segtext, ".text", 05)
		sect.Vaddr = va
		s.Sect = sect

		ctxt.Syms.Lookup(fmt.Sprintf("runtime.text.%d", n), 0).Sect = sect
		n++
	}
	va += funcsize

	return sect, n, va
}

// address assigns virtual addresses to all segments and sections and
// returns all segments in file order.
func (ctxt *Link) address(psess *PackageSession) []*sym.Segment {
	var order []*sym.Segment // Layout order

	va := uint64(*psess.FlagTextAddr)
	order = append(order, &psess.Segtext)
	psess.
		Segtext.Rwx = 05
	psess.
		Segtext.Vaddr = va
	for _, s := range psess.Segtext.Sections {
		va = uint64(Rnd(int64(va), int64(s.Align)))
		s.Vaddr = va
		va += s.Length
	}
	psess.
		Segtext.Length = va - uint64(*psess.FlagTextAddr)
	if ctxt.HeadType == objabi.Hnacl {
		va += 32
	}

	if len(psess.Segrodata.Sections) > 0 {

		va = uint64(Rnd(int64(va), int64(*psess.FlagRound)))

		order = append(order, &psess.Segrodata)
		psess.
			Segrodata.Rwx = 04
		psess.
			Segrodata.Vaddr = va
		for _, s := range psess.Segrodata.Sections {
			va = uint64(Rnd(int64(va), int64(s.Align)))
			s.Vaddr = va
			va += s.Length
		}
		psess.
			Segrodata.Length = va - psess.Segrodata.Vaddr
	}
	if len(psess.Segrelrodata.Sections) > 0 {

		va = uint64(Rnd(int64(va), int64(*psess.FlagRound)))

		order = append(order, &psess.Segrelrodata)
		psess.
			Segrelrodata.Rwx = 06
		psess.
			Segrelrodata.Vaddr = va
		for _, s := range psess.Segrelrodata.Sections {
			va = uint64(Rnd(int64(va), int64(s.Align)))
			s.Vaddr = va
			va += s.Length
		}
		psess.
			Segrelrodata.Length = va - psess.Segrelrodata.Vaddr
	}

	va = uint64(Rnd(int64(va), int64(*psess.FlagRound)))
	order = append(order, &psess.Segdata)
	psess.
		Segdata.Rwx = 06
	psess.
		Segdata.Vaddr = va
	var data *sym.Section
	var noptr *sym.Section
	var bss *sym.Section
	var noptrbss *sym.Section
	for i, s := range psess.Segdata.Sections {
		if ctxt.IsELF && s.Name == ".tbss" {
			continue
		}
		vlen := int64(s.Length)
		if i+1 < len(psess.Segdata.Sections) && !(ctxt.IsELF && psess.Segdata.Sections[i+1].Name == ".tbss") {
			vlen = int64(psess.Segdata.Sections[i+1].Vaddr - s.Vaddr)
		}
		s.Vaddr = va
		va += uint64(vlen)
		psess.
			Segdata.Length = va - psess.Segdata.Vaddr
		if s.Name == ".data" {
			data = s
		}
		if s.Name == ".noptrdata" {
			noptr = s
		}
		if s.Name == ".bss" {
			bss = s
		}
		if s.Name == ".noptrbss" {
			noptrbss = s
		}
	}
	psess.
		Segdata.Filelen = bss.Vaddr - psess.Segdata.Vaddr

	va = uint64(Rnd(int64(va), int64(*psess.FlagRound)))
	order = append(order, &psess.Segdwarf)
	psess.
		Segdwarf.Rwx = 06
	psess.
		Segdwarf.Vaddr = va
	for i, s := range psess.Segdwarf.Sections {
		vlen := int64(s.Length)
		if i+1 < len(psess.Segdwarf.Sections) {
			vlen = int64(psess.Segdwarf.Sections[i+1].Vaddr - s.Vaddr)
		}
		s.Vaddr = va
		va += uint64(vlen)
		if ctxt.HeadType == objabi.Hwindows {
			va = uint64(Rnd(int64(va), psess.PEFILEALIGN))
		}
		psess.
			Segdwarf.Length = va - psess.Segdwarf.Vaddr
	}

	var (
		text     = psess.Segtext.Sections[0]
		rodata   = ctxt.Syms.Lookup("runtime.rodata", 0).Sect
		itablink = ctxt.Syms.Lookup("runtime.itablink", 0).Sect
		symtab   = ctxt.Syms.Lookup("runtime.symtab", 0).Sect
		pclntab  = ctxt.Syms.Lookup("runtime.pclntab", 0).Sect
		types    = ctxt.Syms.Lookup("runtime.types", 0).Sect
	)
	lasttext := text

	for _, sect := range psess.Segtext.Sections {
		if sect.Name == ".text" {
			lasttext = sect
		}
	}

	for _, s := range psess.datap {
		if s.Sect != nil {
			s.Value += int64(s.Sect.Vaddr)
		}
		for sub := s.Sub; sub != nil; sub = sub.Sub {
			sub.Value += s.Value
		}
	}

	for _, s := range psess.dwarfp {
		if s.Sect != nil {
			s.Value += int64(s.Sect.Vaddr)
		}
		for sub := s.Sub; sub != nil; sub = sub.Sub {
			sub.Value += s.Value
		}
	}

	if ctxt.BuildMode == BuildModeShared {
		s := ctxt.Syms.Lookup("go.link.abihashbytes", 0)
		sectSym := ctxt.Syms.Lookup(".note.go.abihash", 0)
		s.Sect = sectSym.Sect
		s.Value = int64(sectSym.Sect.Vaddr + 16)
	}

	ctxt.xdefine("runtime.text", sym.STEXT, int64(text.Vaddr))
	ctxt.xdefine("runtime.etext", sym.STEXT, int64(lasttext.Vaddr+lasttext.Length))

	n := 1
	for _, sect := range psess.Segtext.Sections[1:] {
		if sect.Name != ".text" {
			break
		}
		symname := fmt.Sprintf("runtime.text.%d", n)
		ctxt.xdefine(symname, sym.STEXT, int64(sect.Vaddr))
		n++
	}

	ctxt.xdefine("runtime.rodata", sym.SRODATA, int64(rodata.Vaddr))
	ctxt.xdefine("runtime.erodata", sym.SRODATA, int64(rodata.Vaddr+rodata.Length))
	ctxt.xdefine("runtime.types", sym.SRODATA, int64(types.Vaddr))
	ctxt.xdefine("runtime.etypes", sym.SRODATA, int64(types.Vaddr+types.Length))
	ctxt.xdefine("runtime.itablink", sym.SRODATA, int64(itablink.Vaddr))
	ctxt.xdefine("runtime.eitablink", sym.SRODATA, int64(itablink.Vaddr+itablink.Length))

	s := ctxt.Syms.Lookup("runtime.gcdata", 0)
	s.Attr |= sym.AttrLocal
	ctxt.xdefine("runtime.egcdata", sym.SRODATA, psess.Symaddr(s)+s.Size)
	ctxt.Syms.Lookup("runtime.egcdata", 0).Sect = s.Sect

	s = ctxt.Syms.Lookup("runtime.gcbss", 0)
	s.Attr |= sym.AttrLocal
	ctxt.xdefine("runtime.egcbss", sym.SRODATA, psess.Symaddr(s)+s.Size)
	ctxt.Syms.Lookup("runtime.egcbss", 0).Sect = s.Sect

	ctxt.xdefine("runtime.symtab", sym.SRODATA, int64(symtab.Vaddr))
	ctxt.xdefine("runtime.esymtab", sym.SRODATA, int64(symtab.Vaddr+symtab.Length))
	ctxt.xdefine("runtime.pclntab", sym.SRODATA, int64(pclntab.Vaddr))
	ctxt.xdefine("runtime.epclntab", sym.SRODATA, int64(pclntab.Vaddr+pclntab.Length))
	ctxt.xdefine("runtime.noptrdata", sym.SNOPTRDATA, int64(noptr.Vaddr))
	ctxt.xdefine("runtime.enoptrdata", sym.SNOPTRDATA, int64(noptr.Vaddr+noptr.Length))
	ctxt.xdefine("runtime.bss", sym.SBSS, int64(bss.Vaddr))
	ctxt.xdefine("runtime.ebss", sym.SBSS, int64(bss.Vaddr+bss.Length))
	ctxt.xdefine("runtime.data", sym.SDATA, int64(data.Vaddr))
	ctxt.xdefine("runtime.edata", sym.SDATA, int64(data.Vaddr+data.Length))
	ctxt.xdefine("runtime.noptrbss", sym.SNOPTRBSS, int64(noptrbss.Vaddr))
	ctxt.xdefine("runtime.enoptrbss", sym.SNOPTRBSS, int64(noptrbss.Vaddr+noptrbss.Length))
	ctxt.xdefine("runtime.end", sym.SBSS, int64(psess.Segdata.Vaddr+psess.Segdata.Length))

	return order
}

// layout assigns file offsets and lengths to the segments in order.
func (ctxt *Link) layout(psess *PackageSession, order []*sym.Segment) {
	var prev *sym.Segment
	for _, seg := range order {
		if prev == nil {
			seg.Fileoff = uint64(psess.HEADR)
		} else {
			switch ctxt.HeadType {
			default:

				seg.Fileoff = uint64(Rnd(int64(prev.Fileoff+prev.Filelen), int64(*psess.FlagRound)))
				if seg.Vaddr%uint64(*psess.FlagRound) != seg.Fileoff%uint64(*psess.FlagRound) {
					psess.
						Exitf("bad segment rounding (Vaddr=%#x Fileoff=%#x FlagRound=%#x)", seg.Vaddr, seg.Fileoff, *psess.FlagRound)
				}
			case objabi.Hwindows:
				seg.Fileoff = prev.Fileoff + uint64(Rnd(int64(prev.Filelen), psess.PEFILEALIGN))
			case objabi.Hplan9:
				seg.Fileoff = prev.Fileoff + prev.Filelen
			}
		}
		if seg != &psess.Segdata {

			seg.Filelen = seg.Length
		}
		prev = seg
	}

}

// add a trampoline with symbol s (to be laid down after the current function)
func (ctxt *Link) AddTramp(psess *PackageSession, s *sym.Symbol) {
	s.Type = sym.STEXT
	s.Attr |= sym.AttrReachable
	s.Attr |= sym.AttrOnList
	ctxt.tramps = append(ctxt.tramps, s)
	if *psess.FlagDebugTramp > 0 && ctxt.Debugvlog > 0 {
		ctxt.Logf("trampoline %s inserted\n", s)
	}
}

// compressSyms compresses syms and returns the contents of the
// compressed section. If the section would get larger, it returns nil.
func (psess *PackageSession) compressSyms(ctxt *Link, syms []*sym.Symbol) []byte {
	var total int64
	for _, sym := range syms {
		total += sym.Size
	}

	var buf bytes.Buffer
	buf.Write([]byte("ZLIB"))
	var sizeBytes [8]byte
	binary.BigEndian.PutUint64(sizeBytes[:], uint64(total))
	buf.Write(sizeBytes[:])

	z := zlib.NewWriter(&buf)
	for _, sym := range syms {
		if _, err := z.Write(sym.P); err != nil {
			log.Fatalf("compression failed: %s", err)
		}
		for i := sym.Size - int64(len(sym.P)); i > 0; {
			b := psess.zeros[:]
			if i < int64(len(b)) {
				b = b[:i]
			}
			n, err := z.Write(b)
			if err != nil {
				log.Fatalf("compression failed: %s", err)
			}
			i -= int64(n)
		}
	}
	if err := z.Close(); err != nil {
		log.Fatalf("compression failed: %s", err)
	}
	if int64(buf.Len()) >= total {

		return nil
	}
	return buf.Bytes()
}

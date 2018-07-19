package ppc64

import (
	"debug/elf"
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"log"
)

func (psess *PackageSession) genplt(ctxt *ld.Link) {

	// Find all R_PPC64_REL24 relocations that reference dynamic
	// imports. Reserve PLT entries for these symbols and
	// generate call stubs. The call stubs need to live in .text,
	// which is why we need to do this pass this early.
	//
	// This assumes "case 1" from the ABI, where the caller needs
	// us to save and restore the TOC pointer.
	var stubs []*sym.Symbol
	for _, s := range ctxt.Textp {
		for i := range s.R {
			r := &s.R[i]
			if r.Type != 256+objabi.RelocType(elf.R_PPC64_REL24) || r.Sym.Type != sym.SDYNIMPORT {
				continue
			}
			psess.
				addpltsym(ctxt, r.Sym)

			n := fmt.Sprintf("%s.%s", s.Name, r.Sym.Name)

			stub := ctxt.Syms.Lookup(n, 0)
			if s.Attr.Reachable() {
				stub.Attr |= sym.AttrReachable
			}
			if stub.Size == 0 {

				stub.Outer = s
				stubs = append(stubs, stub)
				gencallstub(ctxt, 1, stub, r.Sym)
			}

			r.Sym = stub

			// Restore TOC after bl. The compiler put a
			// nop here for us to overwrite.
			const o1 = 0xe8410018 // ld r2,24(r1)
			ctxt.Arch.ByteOrder.PutUint32(s.P[r.Off+4:], o1)
		}
	}

	ctxt.Textp = append(stubs, ctxt.Textp...)
}

func genaddmoduledata(ctxt *ld.Link) {
	addmoduledata := ctxt.Syms.ROLookup("runtime.addmoduledata", 0)
	if addmoduledata.Type == sym.STEXT && ctxt.BuildMode != ld.BuildModePlugin {
		return
	}
	addmoduledata.Attr |= sym.AttrReachable
	initfunc := ctxt.Syms.Lookup("go.link.addmoduledata", 0)
	initfunc.Type = sym.STEXT
	initfunc.Attr |= sym.AttrLocal
	initfunc.Attr |= sym.AttrReachable
	o := func(op uint32) {
		initfunc.AddUint32(ctxt.Arch, op)
	}

	rel := initfunc.AddRel()
	rel.Off = int32(initfunc.Size)
	rel.Siz = 8
	rel.Sym = ctxt.Syms.Lookup(".TOC.", 0)
	rel.Sym.Attr |= sym.AttrReachable
	rel.Type = objabi.R_ADDRPOWER_PCREL
	o(0x3c4c0000)

	o(0x38420000)

	o(0x7c0802a6)

	o(0xf801ffe1)

	rel = initfunc.AddRel()
	rel.Off = int32(initfunc.Size)
	rel.Siz = 8
	if s := ctxt.Syms.ROLookup("local.moduledata", 0); s != nil {
		rel.Sym = s
	} else if s := ctxt.Syms.ROLookup("local.pluginmoduledata", 0); s != nil {
		rel.Sym = s
	} else {
		rel.Sym = ctxt.Syms.Lookup("runtime.firstmoduledata", 0)
	}
	rel.Sym.Attr |= sym.AttrReachable
	rel.Sym.Attr |= sym.AttrLocal
	rel.Type = objabi.R_ADDRPOWER_GOT
	o(0x3c620000)

	o(0xe8630000)

	rel = initfunc.AddRel()
	rel.Off = int32(initfunc.Size)
	rel.Siz = 4
	rel.Sym = addmoduledata
	rel.Type = objabi.R_CALLPOWER
	o(0x48000001)

	o(0x60000000)

	o(0xe8010000)

	o(0x7c0803a6)

	o(0x38210020)

	o(0x4e800020)

	if ctxt.BuildMode == ld.BuildModePlugin {
		ctxt.Textp = append(ctxt.Textp, addmoduledata)
	}
	initarray_entry := ctxt.Syms.Lookup("go.link.addmoduledatainit", 0)
	ctxt.Textp = append(ctxt.Textp, initfunc)
	initarray_entry.Attr |= sym.AttrReachable
	initarray_entry.Attr |= sym.AttrLocal
	initarray_entry.Type = sym.SINITARR
	initarray_entry.AddAddr(ctxt.Arch, initfunc)
}

func (psess *PackageSession) gentext(ctxt *ld.Link) {
	if ctxt.DynlinkingGo() {
		genaddmoduledata(ctxt)
	}

	if ctxt.LinkMode == ld.LinkInternal {
		psess.
			genplt(ctxt)
	}
}

// Construct a call stub in stub that calls symbol targ via its PLT
// entry.
func gencallstub(ctxt *ld.Link, abicase int, stub *sym.Symbol, targ *sym.Symbol) {
	if abicase != 1 {

		log.Fatalf("gencallstub only implements case 1 calls")
	}

	plt := ctxt.Syms.Lookup(".plt", 0)

	stub.Type = sym.STEXT

	stub.AddUint32(ctxt.Arch, 0xf8410018)

	r := stub.AddRel()

	r.Off = int32(stub.Size)
	r.Sym = plt
	r.Add = int64(targ.Plt)
	r.Siz = 2
	if ctxt.Arch.ByteOrder == binary.BigEndian {
		r.Off += int32(r.Siz)
	}
	r.Type = objabi.R_POWER_TOC
	r.Variant = sym.RV_POWER_HA
	stub.AddUint32(ctxt.Arch, 0x3d820000)
	r = stub.AddRel()
	r.Off = int32(stub.Size)
	r.Sym = plt
	r.Add = int64(targ.Plt)
	r.Siz = 2
	if ctxt.Arch.ByteOrder == binary.BigEndian {
		r.Off += int32(r.Siz)
	}
	r.Type = objabi.R_POWER_TOC
	r.Variant = sym.RV_POWER_LO
	stub.AddUint32(ctxt.Arch, 0xe98c0000)

	stub.AddUint32(ctxt.Arch, 0x7d8903a6)
	stub.AddUint32(ctxt.Arch, 0x4e800420)
}

func (psess *PackageSession) adddynrel(ctxt *ld.Link, s *sym.Symbol, r *sym.Reloc) bool {
	targ := r.Sym

	switch r.Type {
	default:
		if r.Type >= 256 {
			psess.ld.
				Errorf(s, "unexpected relocation type %d (%s)", r.Type, psess.sym.RelocName(ctxt.Arch, r.Type))
			return false
		}

	case 256 + objabi.RelocType(elf.R_PPC64_REL24):
		r.Type = objabi.R_CALLPOWER

		r.Add += int64(r.Sym.Localentry) * 4

		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_PPC64_REL24 for dyn import")
		}

		return true

	case 256 + objabi.RelocType(elf.R_PPC_REL32):
		r.Type = objabi.R_PCREL
		r.Add += 4

		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_PPC_REL32 for dyn import")
		}

		return true

	case 256 + objabi.RelocType(elf.R_PPC64_ADDR64):
		r.Type = objabi.R_ADDR
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Adddynsym(ctxt, targ)

			rela := ctxt.Syms.Lookup(".rela", 0)
			rela.AddAddrPlus(ctxt.Arch, s, int64(r.Off))
			rela.AddUint64(ctxt.Arch, ld.ELF64_R_INFO(uint32(targ.Dynid), uint32(elf.R_PPC64_ADDR64)))
			rela.AddUint64(ctxt.Arch, uint64(r.Add))
			r.Type = 256
		}

		return true

	case 256 + objabi.RelocType(elf.R_PPC64_TOC16):
		r.Type = objabi.R_POWER_TOC
		r.Variant = sym.RV_POWER_LO | sym.RV_CHECK_OVERFLOW
		return true

	case 256 + objabi.RelocType(elf.R_PPC64_TOC16_LO):
		r.Type = objabi.R_POWER_TOC
		r.Variant = sym.RV_POWER_LO
		return true

	case 256 + objabi.RelocType(elf.R_PPC64_TOC16_HA):
		r.Type = objabi.R_POWER_TOC
		r.Variant = sym.RV_POWER_HA | sym.RV_CHECK_OVERFLOW
		return true

	case 256 + objabi.RelocType(elf.R_PPC64_TOC16_HI):
		r.Type = objabi.R_POWER_TOC
		r.Variant = sym.RV_POWER_HI | sym.RV_CHECK_OVERFLOW
		return true

	case 256 + objabi.RelocType(elf.R_PPC64_TOC16_DS):
		r.Type = objabi.R_POWER_TOC
		r.Variant = sym.RV_POWER_DS | sym.RV_CHECK_OVERFLOW
		return true

	case 256 + objabi.RelocType(elf.R_PPC64_TOC16_LO_DS):
		r.Type = objabi.R_POWER_TOC
		r.Variant = sym.RV_POWER_DS
		return true

	case 256 + objabi.RelocType(elf.R_PPC64_REL16_LO):
		r.Type = objabi.R_PCREL
		r.Variant = sym.RV_POWER_LO
		r.Add += 2
		return true

	case 256 + objabi.RelocType(elf.R_PPC64_REL16_HI):
		r.Type = objabi.R_PCREL
		r.Variant = sym.RV_POWER_HI | sym.RV_CHECK_OVERFLOW
		r.Add += 2
		return true

	case 256 + objabi.RelocType(elf.R_PPC64_REL16_HA):
		r.Type = objabi.R_PCREL
		r.Variant = sym.RV_POWER_HA | sym.RV_CHECK_OVERFLOW
		r.Add += 2
		return true
	}

	if targ.Type != sym.SDYNIMPORT {
		return true
	}

	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	ctxt.Out.Write64(uint64(sectoff))

	elfsym := r.Xsym.ElfsymForReloc()
	switch r.Type {
	default:
		return false
	case objabi.R_ADDR:
		switch r.Siz {
		case 4:
			ctxt.Out.Write64(uint64(elf.R_PPC64_ADDR32) | uint64(elfsym)<<32)
		case 8:
			ctxt.Out.Write64(uint64(elf.R_PPC64_ADDR64) | uint64(elfsym)<<32)
		default:
			return false
		}
	case objabi.R_POWER_TLS:
		ctxt.Out.Write64(uint64(elf.R_PPC64_TLS) | uint64(elfsym)<<32)
	case objabi.R_POWER_TLS_LE:
		ctxt.Out.Write64(uint64(elf.R_PPC64_TPREL16) | uint64(elfsym)<<32)
	case objabi.R_POWER_TLS_IE:
		ctxt.Out.Write64(uint64(elf.R_PPC64_GOT_TPREL16_HA) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd))
		ctxt.Out.Write64(uint64(sectoff + 4))
		ctxt.Out.Write64(uint64(elf.R_PPC64_GOT_TPREL16_LO_DS) | uint64(elfsym)<<32)
	case objabi.R_ADDRPOWER:
		ctxt.Out.Write64(uint64(elf.R_PPC64_ADDR16_HA) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd))
		ctxt.Out.Write64(uint64(sectoff + 4))
		ctxt.Out.Write64(uint64(elf.R_PPC64_ADDR16_LO) | uint64(elfsym)<<32)
	case objabi.R_ADDRPOWER_DS:
		ctxt.Out.Write64(uint64(elf.R_PPC64_ADDR16_HA) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd))
		ctxt.Out.Write64(uint64(sectoff + 4))
		ctxt.Out.Write64(uint64(elf.R_PPC64_ADDR16_LO_DS) | uint64(elfsym)<<32)
	case objabi.R_ADDRPOWER_GOT:
		ctxt.Out.Write64(uint64(elf.R_PPC64_GOT16_HA) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd))
		ctxt.Out.Write64(uint64(sectoff + 4))
		ctxt.Out.Write64(uint64(elf.R_PPC64_GOT16_LO_DS) | uint64(elfsym)<<32)
	case objabi.R_ADDRPOWER_PCREL:
		ctxt.Out.Write64(uint64(elf.R_PPC64_REL16_HA) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd))
		ctxt.Out.Write64(uint64(sectoff + 4))
		ctxt.Out.Write64(uint64(elf.R_PPC64_REL16_LO) | uint64(elfsym)<<32)
		r.Xadd += 4
	case objabi.R_ADDRPOWER_TOCREL:
		ctxt.Out.Write64(uint64(elf.R_PPC64_TOC16_HA) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd))
		ctxt.Out.Write64(uint64(sectoff + 4))
		ctxt.Out.Write64(uint64(elf.R_PPC64_TOC16_LO) | uint64(elfsym)<<32)
	case objabi.R_ADDRPOWER_TOCREL_DS:
		ctxt.Out.Write64(uint64(elf.R_PPC64_TOC16_HA) | uint64(elfsym)<<32)
		ctxt.Out.Write64(uint64(r.Xadd))
		ctxt.Out.Write64(uint64(sectoff + 4))
		ctxt.Out.Write64(uint64(elf.R_PPC64_TOC16_LO_DS) | uint64(elfsym)<<32)
	case objabi.R_CALLPOWER:
		if r.Siz != 4 {
			return false
		}
		ctxt.Out.Write64(uint64(elf.R_PPC64_REL24) | uint64(elfsym)<<32)

	}
	ctxt.Out.Write64(uint64(r.Xadd))

	return true
}

func elfsetupplt(ctxt *ld.Link) {
	plt := ctxt.Syms.Lookup(".plt", 0)
	if plt.Size == 0 {

		plt.Size = 16
	}
}

func machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	return false
}

// Return the value of .TOC. for symbol s
func (psess *PackageSession) symtoc(ctxt *ld.Link, s *sym.Symbol) int64 {
	var toc *sym.Symbol

	if s.Outer != nil {
		toc = ctxt.Syms.ROLookup(".TOC.", int(s.Outer.Version))
	} else {
		toc = ctxt.Syms.ROLookup(".TOC.", int(s.Version))
	}

	if toc == nil {
		psess.ld.
			Errorf(s, "TOC-relative relocation in object without .TOC.")
		return 0
	}

	return toc.Value
}

func (psess *PackageSession) archrelocaddr(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val *int64) bool {
	var o1, o2 uint32
	if ctxt.Arch.ByteOrder == binary.BigEndian {
		o1 = uint32(*val >> 32)
		o2 = uint32(*val)
	} else {
		o1 = uint32(*val)
		o2 = uint32(*val >> 32)
	}

	t := psess.ld.Symaddr(r.Sym) + r.Add
	if t < 0 || t >= 1<<31 {
		psess.ld.
			Errorf(s, "relocation for %s is too big (>=2G): %d", s.Name, psess.ld.Symaddr(r.Sym))
	}
	if t&0x8000 != 0 {
		t += 0x10000
	}

	switch r.Type {
	case objabi.R_ADDRPOWER:
		o1 |= (uint32(t) >> 16) & 0xffff
		o2 |= uint32(t) & 0xffff
	case objabi.R_ADDRPOWER_DS:
		o1 |= (uint32(t) >> 16) & 0xffff
		if t&3 != 0 {
			psess.ld.
				Errorf(s, "bad DS reloc for %s: %d", s.Name, psess.ld.Symaddr(r.Sym))
		}
		o2 |= uint32(t) & 0xfffc
	default:
		return false
	}

	if ctxt.Arch.ByteOrder == binary.BigEndian {
		*val = int64(o1)<<32 | int64(o2)
	} else {
		*val = int64(o2)<<32 | int64(o1)
	}
	return true
}

// resolve direct jump relocation r in s, and add trampoline if necessary
func (psess *PackageSession) trampoline(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol) {

	if ctxt.LinkMode == ld.LinkExternal && (ctxt.DynlinkingGo() || ctxt.BuildMode == ld.BuildModeCArchive || ctxt.BuildMode == ld.BuildModeCShared || ctxt.BuildMode == ld.BuildModePIE) {

		return
	}

	t := psess.ld.Symaddr(r.Sym) + r.Add - (s.Value + int64(r.Off))
	switch r.Type {
	case objabi.R_CALLPOWER:

		if (ctxt.LinkMode == ld.LinkExternal && s.Sect != r.Sym.Sect) || (ctxt.LinkMode == ld.LinkInternal && int64(int32(t<<6)>>6) != t) || (*psess.ld.FlagDebugTramp > 1 && s.File != r.Sym.File) {
			var tramp *sym.Symbol
			for i := 0; ; i++ {

				name := r.Sym.Name
				if r.Add == 0 {
					name = name + fmt.Sprintf("-tramp%d", i)
				} else {
					name = name + fmt.Sprintf("%+x-tramp%d", r.Add, i)
				}

				tramp = ctxt.Syms.Lookup(name, int(r.Sym.Version))
				if tramp.Value == 0 {
					break
				}

				t = psess.ld.Symaddr(tramp) + r.Add - (s.Value + int64(r.Off))

				if (ctxt.LinkMode == ld.LinkInternal && int64(int32(t<<6)>>6) == t) || (ctxt.LinkMode == ld.LinkExternal && s.Sect == tramp.Sect) {
					break
				}
			}
			if tramp.Type == 0 {
				if ctxt.DynlinkingGo() || ctxt.BuildMode == ld.BuildModeCArchive || ctxt.BuildMode == ld.BuildModeCShared || ctxt.BuildMode == ld.BuildModePIE {
					psess.ld.
						Errorf(s, "unexpected trampoline for shared or dynamic linking\n")
				} else {
					ctxt.AddTramp(psess.ld, tramp)
					psess.
						gentramp(ctxt.Arch, ctxt.LinkMode, tramp, r.Sym, r.Add)
				}
			}
			r.Sym = tramp
			r.Add = 0
			r.Done = false
		}
	default:
		psess.ld.
			Errorf(s, "trampoline called with non-jump reloc: %d (%s)", r.Type, psess.sym.RelocName(ctxt.Arch, r.Type))
	}
}

func (psess *PackageSession) gentramp(arch *sys.Arch, linkmode ld.LinkMode, tramp, target *sym.Symbol, offset int64) {

	tramp.Size = 16
	tramp.P = make([]byte, tramp.Size)
	t := psess.ld.Symaddr(target) + offset
	o1 := uint32(0x3fe00000)
	o2 := uint32(0x3bff0000)

	if linkmode == ld.LinkExternal {
		tr := tramp.AddRel()
		tr.Off = 0
		tr.Type = objabi.R_ADDRPOWER
		tr.Siz = 8
		tr.Sym = target
		tr.Add = offset
	} else {

		val := uint32((t & 0xffff0000) >> 16)
		if t&0x8000 != 0 {
			val += 1
		}
		o1 |= val
		o2 |= uint32(t & 0xffff)
	}
	o3 := uint32(0x7fe903a6)
	o4 := uint32(0x4e800420)
	arch.ByteOrder.PutUint32(tramp.P, o1)
	arch.ByteOrder.PutUint32(tramp.P[4:], o2)
	arch.ByteOrder.PutUint32(tramp.P[8:], o3)
	arch.ByteOrder.PutUint32(tramp.P[12:], o4)
}

func (psess *PackageSession) archreloc(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val *int64) bool {
	if ctxt.LinkMode == ld.LinkExternal {
		switch r.Type {
		default:
			return false
		case objabi.R_POWER_TLS, objabi.R_POWER_TLS_LE, objabi.R_POWER_TLS_IE:
			r.Done = false

			r.Xadd = r.Add
			r.Xsym = r.Sym
			return true
		case objabi.R_ADDRPOWER,
			objabi.R_ADDRPOWER_DS,
			objabi.R_ADDRPOWER_TOCREL,
			objabi.R_ADDRPOWER_TOCREL_DS,
			objabi.R_ADDRPOWER_GOT,
			objabi.R_ADDRPOWER_PCREL:
			r.Done = false

			rs := r.Sym
			r.Xadd = r.Add
			for rs.Outer != nil {
				r.Xadd += psess.ld.Symaddr(rs) - psess.ld.Symaddr(rs.Outer)
				rs = rs.Outer
			}

			if rs.Type != sym.SHOSTOBJ && rs.Type != sym.SDYNIMPORT && rs.Sect == nil {
				psess.ld.
					Errorf(s, "missing section for %s", rs.Name)
			}
			r.Xsym = rs

			return true
		case objabi.R_CALLPOWER:
			r.Done = false
			r.Xsym = r.Sym
			r.Xadd = r.Add
			return true
		}
	}

	switch r.Type {
	case objabi.R_CONST:
		*val = r.Add
		return true
	case objabi.R_GOTOFF:
		*val = psess.ld.Symaddr(r.Sym) + r.Add - psess.ld.Symaddr(ctxt.Syms.Lookup(".got", 0))
		return true
	case objabi.R_ADDRPOWER, objabi.R_ADDRPOWER_DS:
		return psess.archrelocaddr(ctxt, r, s, val)
	case objabi.R_CALLPOWER:

		t := psess.ld.Symaddr(r.Sym) + r.Add - (s.Value + int64(r.Off))

		if t&3 != 0 {
			psess.ld.
				Errorf(s, "relocation for %s+%d is not aligned: %d", r.Sym.Name, r.Off, t)
		}

		if int64(int32(t<<6)>>6) != t {
			psess.ld.
				Errorf(s, "direct call too far: %s %x", r.Sym.Name, t)
		}
		*val |= int64(uint32(t) &^ 0xfc000003)
		return true
	case objabi.R_POWER_TOC:
		*val = psess.ld.Symaddr(r.Sym) + r.Add - psess.symtoc(ctxt, s)

		return true
	case objabi.R_POWER_TLS_LE:

		v := r.Sym.Value - 0x7000
		if int64(int16(v)) != v {
			psess.ld.
				Errorf(s, "TLS offset out of range %d", v)
		}
		*val = (*val &^ 0xffff) | (v & 0xffff)
		return true
	}

	return false
}

func (psess *PackageSession) archrelocvariant(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	switch r.Variant & sym.RV_TYPE_MASK {
	default:
		psess.ld.
			Errorf(s, "unexpected relocation variant %d", r.Variant)
		fallthrough

	case sym.RV_NONE:
		return t

	case sym.RV_POWER_LO:
		if r.Variant&sym.RV_CHECK_OVERFLOW != 0 {
			// Whether to check for signed or unsigned
			// overflow depends on the instruction
			var o1 uint32
			if ctxt.Arch.ByteOrder == binary.BigEndian {
				o1 = ld.Be32(s.P[r.Off-2:])
			} else {
				o1 = ld.Le32(s.P[r.Off:])
			}
			switch o1 >> 26 {
			case 24,
				26,
				28:
				if t>>16 != 0 {
					goto overflow
				}

			default:
				if int64(int16(t)) != t {
					goto overflow
				}
			}
		}

		return int64(int16(t))

	case sym.RV_POWER_HA:
		t += 0x8000
		fallthrough

	case sym.RV_POWER_HI:
		t >>= 16

		if r.Variant&sym.RV_CHECK_OVERFLOW != 0 {
			// Whether to check for signed or unsigned
			// overflow depends on the instruction
			var o1 uint32
			if ctxt.Arch.ByteOrder == binary.BigEndian {
				o1 = ld.Be32(s.P[r.Off-2:])
			} else {
				o1 = ld.Le32(s.P[r.Off:])
			}
			switch o1 >> 26 {
			case 25,
				27,
				29:
				if t>>16 != 0 {
					goto overflow
				}

			default:
				if int64(int16(t)) != t {
					goto overflow
				}
			}
		}

		return int64(int16(t))

	case sym.RV_POWER_DS:
		var o1 uint32
		if ctxt.Arch.ByteOrder == binary.BigEndian {
			o1 = uint32(ld.Be16(s.P[r.Off:]))
		} else {
			o1 = uint32(ld.Le16(s.P[r.Off:]))
		}
		if t&3 != 0 {
			psess.ld.
				Errorf(s, "relocation for %s+%d is not aligned: %d", r.Sym.Name, r.Off, t)
		}
		if (r.Variant&sym.RV_CHECK_OVERFLOW != 0) && int64(int16(t)) != t {
			goto overflow
		}
		return int64(o1)&0x3 | int64(int16(t))
	}

overflow:
	psess.ld.
		Errorf(s, "relocation for %s+%d is too big: %d", r.Sym.Name, r.Off, t)
	return t
}

func (psess *PackageSession) addpltsym(ctxt *ld.Link, s *sym.Symbol) {
	if s.Plt >= 0 {
		return
	}
	psess.ld.
		Adddynsym(ctxt, s)

	if ctxt.IsELF {
		plt := ctxt.Syms.Lookup(".plt", 0)
		rela := ctxt.Syms.Lookup(".rela.plt", 0)
		if plt.Size == 0 {
			elfsetupplt(ctxt)
		}

		glink := psess.ensureglinkresolver(ctxt)

		r := glink.AddRel()

		r.Sym = glink
		r.Off = int32(glink.Size)
		r.Siz = 4
		r.Type = objabi.R_CALLPOWER
		glink.AddUint32(ctxt.Arch, 0x48000000)

		s.Plt = int32(plt.Size)

		plt.Size += 8

		rela.AddAddrPlus(ctxt.Arch, plt, int64(s.Plt))
		rela.AddUint64(ctxt.Arch, ld.ELF64_R_INFO(uint32(s.Dynid), uint32(elf.R_PPC64_JMP_SLOT)))
		rela.AddUint64(ctxt.Arch, 0)
	} else {
		psess.ld.
			Errorf(s, "addpltsym: unsupported binary format")
	}
}

// Generate the glink resolver stub if necessary and return the .glink section
func (psess *PackageSession) ensureglinkresolver(ctxt *ld.Link) *sym.Symbol {
	glink := ctxt.Syms.Lookup(".glink", 0)
	if glink.Size != 0 {
		return glink
	}

	glink.AddUint32(ctxt.Arch, 0x7c0802a6)
	glink.AddUint32(ctxt.Arch, 0x429f0005)
	glink.AddUint32(ctxt.Arch, 0x7d6802a6)
	glink.AddUint32(ctxt.Arch, 0x7c0803a6)

	glink.AddUint32(ctxt.Arch, 0x3800ffd0)
	glink.AddUint32(ctxt.Arch, 0x7c006214)
	glink.AddUint32(ctxt.Arch, 0x7c0b0050)
	glink.AddUint32(ctxt.Arch, 0x7800f082)

	r := glink.AddRel()

	r.Off = int32(glink.Size)
	r.Sym = ctxt.Syms.Lookup(".plt", 0)
	r.Siz = 8
	r.Type = objabi.R_ADDRPOWER

	glink.AddUint32(ctxt.Arch, 0x3d600000)
	glink.AddUint32(ctxt.Arch, 0x396b0000)

	glink.AddUint32(ctxt.Arch, 0xe98b0000)
	glink.AddUint32(ctxt.Arch, 0xe96b0008)

	glink.AddUint32(ctxt.Arch, 0x7d8903a6)
	glink.AddUint32(ctxt.Arch, 0x4e800420)

	s := ctxt.Syms.Lookup(".dynamic", 0)
	psess.ld.
		Elfwritedynentsymplus(ctxt, s, ld.DT_PPC64_GLINK, glink, glink.Size-32)

	return glink
}

func (psess *PackageSession) asmb(ctxt *ld.Link) {
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f asmb\n", psess.ld.Cputime())
	}

	if ctxt.IsELF {
		psess.ld.
			Asmbelfsetup()
	}

	for _, sect := range psess.ld.Segtext.Sections {
		ctxt.Out.SeekSet(psess.ld, int64(sect.Vaddr-psess.ld.Segtext.Vaddr+psess.ld.Segtext.Fileoff))

		if sect.Name == ".text" {
			psess.ld.
				Codeblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
		} else {
			psess.ld.
				Datblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
		}
	}

	if psess.ld.Segrodata.Filelen > 0 {
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f rodatblk\n", psess.ld.Cputime())
		}
		ctxt.Out.SeekSet(psess.ld, int64(psess.ld.Segrodata.Fileoff))
		psess.ld.
			Datblk(ctxt, int64(psess.ld.Segrodata.Vaddr), int64(psess.ld.Segrodata.Filelen))
	}
	if psess.ld.Segrelrodata.Filelen > 0 {
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f relrodatblk\n", psess.ld.Cputime())
		}
		ctxt.Out.SeekSet(psess.ld, int64(psess.ld.Segrelrodata.Fileoff))
		psess.ld.
			Datblk(ctxt, int64(psess.ld.Segrelrodata.Vaddr), int64(psess.ld.Segrelrodata.Filelen))
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f datblk\n", psess.ld.Cputime())
	}

	ctxt.Out.SeekSet(psess.ld, int64(psess.ld.Segdata.Fileoff))
	psess.ld.
		Datblk(ctxt, int64(psess.ld.Segdata.Vaddr), int64(psess.ld.Segdata.Filelen))

	ctxt.Out.SeekSet(psess.ld, int64(psess.ld.Segdwarf.Fileoff))
	psess.ld.
		Dwarfblk(ctxt, int64(psess.ld.Segdwarf.Vaddr), int64(psess.ld.Segdwarf.Filelen))
	psess.ld.
		Symsize = 0
	psess.ld.
		Lcsize = 0
	symo := uint32(0)
	if !*psess.ld.FlagS {

		if ctxt.Debugvlog != 0 {
			ctxt.Logf("%5.2f sym\n", psess.ld.Cputime())
		}
		switch ctxt.HeadType {
		default:
			if ctxt.IsELF {
				symo = uint32(psess.ld.Segdwarf.Fileoff + psess.ld.Segdwarf.Filelen)
				symo = uint32(ld.Rnd(int64(symo), int64(*psess.ld.FlagRound)))
			}

		case objabi.Hplan9:
			symo = uint32(psess.ld.Segdata.Fileoff + psess.ld.Segdata.Filelen)
		}

		ctxt.Out.SeekSet(psess.ld, int64(symo))
		switch ctxt.HeadType {
		default:
			if ctxt.IsELF {
				if ctxt.Debugvlog != 0 {
					ctxt.Logf("%5.2f elfsym\n", psess.ld.Cputime())
				}
				psess.ld.
					Asmelfsym(ctxt)
				ctxt.Out.Flush(psess.ld)
				ctxt.Out.Write(psess.ld.Elfstrdat)

				if ctxt.LinkMode == ld.LinkExternal {
					psess.ld.
						Elfemitreloc(ctxt)
				}
			}

		case objabi.Hplan9:
			psess.ld.
				Asmplan9sym(ctxt)
			ctxt.Out.Flush(psess.ld)

			sym := ctxt.Syms.Lookup("pclntab", 0)
			if sym != nil {
				psess.ld.
					Lcsize = int32(len(sym.P))
				ctxt.Out.Write(sym.P)
				ctxt.Out.Flush(psess.ld)
			}
		}
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f header\n", psess.ld.Cputime())
	}
	ctxt.Out.SeekSet(psess.ld, 0)
	switch ctxt.HeadType {
	default:
	case objabi.Hplan9:
		ctxt.Out.Write32(0x647)
		ctxt.Out.Write32(uint32(psess.ld.Segtext.Filelen))
		ctxt.Out.Write32(uint32(psess.ld.Segdata.Filelen))
		ctxt.Out.Write32(uint32(psess.ld.Segdata.Length - psess.ld.Segdata.Filelen))
		ctxt.Out.Write32(uint32(psess.ld.Symsize))
		ctxt.Out.Write32(uint32(psess.ld.Entryvalue(ctxt)))
		ctxt.Out.Write32(0)
		ctxt.Out.Write32(uint32(psess.ld.Lcsize))

	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd,
		objabi.Hnacl:
		psess.ld.
			Asmbelf(ctxt, int64(symo))
	}

	ctxt.Out.Flush(psess.ld)
	if *psess.ld.FlagC {
		fmt.Printf("textsize=%d\n", psess.ld.Segtext.Filelen)
		fmt.Printf("datsize=%d\n", psess.ld.Segdata.Filelen)
		fmt.Printf("bsssize=%d\n", psess.ld.Segdata.Length-psess.ld.Segdata.Filelen)
		fmt.Printf("symsize=%d\n", psess.ld.Symsize)
		fmt.Printf("lcsize=%d\n", psess.ld.Lcsize)
		fmt.Printf("total=%d\n", psess.ld.Segtext.Filelen+psess.ld.Segdata.Length+uint64(psess.ld.Symsize)+uint64(psess.ld.Lcsize))
	}
}

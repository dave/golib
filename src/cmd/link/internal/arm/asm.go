package arm

import (
	"debug/elf"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/ld"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"log"
)

func gentext(ctxt *ld.Link) {
	if !ctxt.DynlinkingGo() {
		return
	}
	addmoduledata := ctxt.Syms.Lookup("runtime.addmoduledata", 0)
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
	o(0xe59f0004)
	o(0xe08f0000)

	o(0xeafffffe)
	rel := initfunc.AddRel()
	rel.Off = 8
	rel.Siz = 4
	rel.Sym = ctxt.Syms.Lookup("runtime.addmoduledata", 0)
	rel.Type = objabi.R_CALLARM
	rel.Add = 0xeafffffe

	o(0x00000000)
	rel = initfunc.AddRel()
	rel.Off = 12
	rel.Siz = 4
	rel.Sym = ctxt.Moduledata
	rel.Type = objabi.R_PCREL
	rel.Add = 4

	if ctxt.BuildMode == ld.BuildModePlugin {
		ctxt.Textp = append(ctxt.Textp, addmoduledata)
	}
	ctxt.Textp = append(ctxt.Textp, initfunc)
	initarray_entry := ctxt.Syms.Lookup("go.link.addmoduledatainit", 0)
	initarray_entry.Attr |= sym.AttrReachable
	initarray_entry.Attr |= sym.AttrLocal
	initarray_entry.Type = sym.SINITARR
	initarray_entry.AddAddr(ctxt.Arch, initfunc)
}

// Preserve highest 8 bits of a, and do addition to lower 24-bit
// of a and b; used to adjust ARM branch instruction's target
func braddoff(a int32, b int32) int32 {
	return int32((uint32(a))&0xff000000 | 0x00ffffff&uint32(a+b))
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

	case 256 + objabi.RelocType(elf.R_ARM_PLT32):
		r.Type = objabi.R_CALLARM

		if targ.Type == sym.SDYNIMPORT {
			psess.
				addpltsym(ctxt, targ)
			r.Sym = ctxt.Syms.Lookup(".plt", 0)
			r.Add = int64(braddoff(int32(r.Add), targ.Plt/4))
		}

		return true

	case 256 + objabi.RelocType(elf.R_ARM_THM_PC22):
		psess.ld.
			Exitf("R_ARM_THM_CALL, are you using -marm?")
		return false

	case 256 + objabi.RelocType(elf.R_ARM_GOT32):
		if targ.Type != sym.SDYNIMPORT {
			psess.
				addgotsyminternal(ctxt, targ)
		} else {
			psess.
				addgotsym(ctxt, targ)
		}

		r.Type = objabi.R_CONST
		r.Sym = nil
		r.Add += int64(targ.Got)
		return true

	case 256 + objabi.RelocType(elf.R_ARM_GOT_PREL):
		if targ.Type != sym.SDYNIMPORT {
			psess.
				addgotsyminternal(ctxt, targ)
		} else {
			psess.
				addgotsym(ctxt, targ)
		}

		r.Type = objabi.R_PCREL
		r.Sym = ctxt.Syms.Lookup(".got", 0)
		r.Add += int64(targ.Got) + 4
		return true

	case 256 + objabi.RelocType(elf.R_ARM_GOTOFF):
		r.Type = objabi.R_GOTOFF

		return true

	case 256 + objabi.RelocType(elf.R_ARM_GOTPC):
		r.Type = objabi.R_PCREL

		r.Sym = ctxt.Syms.Lookup(".got", 0)
		r.Add += 4
		return true

	case 256 + objabi.RelocType(elf.R_ARM_CALL):
		r.Type = objabi.R_CALLARM
		if targ.Type == sym.SDYNIMPORT {
			psess.
				addpltsym(ctxt, targ)
			r.Sym = ctxt.Syms.Lookup(".plt", 0)
			r.Add = int64(braddoff(int32(r.Add), targ.Plt/4))
		}

		return true

	case 256 + objabi.RelocType(elf.R_ARM_REL32):
		r.Type = objabi.R_PCREL

		r.Add += 4
		return true

	case 256 + objabi.RelocType(elf.R_ARM_ABS32):
		if targ.Type == sym.SDYNIMPORT {
			psess.ld.
				Errorf(s, "unexpected R_ARM_ABS32 relocation for dynamic symbol %s", targ.Name)
		}
		r.Type = objabi.R_ADDR
		return true

	case 256 + objabi.RelocType(elf.R_ARM_V4BX):
		if r.Sym != nil {

			r.Sym.Type = 0
		}

		r.Sym = nil
		return true

	case 256 + objabi.RelocType(elf.R_ARM_PC24),
		256 + objabi.RelocType(elf.R_ARM_JUMP24):
		r.Type = objabi.R_CALLARM
		if targ.Type == sym.SDYNIMPORT {
			psess.
				addpltsym(ctxt, targ)
			r.Sym = ctxt.Syms.Lookup(".plt", 0)
			r.Add = int64(braddoff(int32(r.Add), targ.Plt/4))
		}

		return true
	}

	if targ.Type != sym.SDYNIMPORT {
		return true
	}

	switch r.Type {
	case objabi.R_CALLARM:
		if ctxt.LinkMode == ld.LinkExternal {

			return true
		}
		psess.
			addpltsym(ctxt, targ)
		r.Sym = ctxt.Syms.Lookup(".plt", 0)
		r.Add = int64(targ.Plt)
		return true

	case objabi.R_ADDR:
		if s.Type != sym.SDATA {
			break
		}
		if ctxt.IsELF {
			psess.ld.
				Adddynsym(ctxt, targ)
			rel := ctxt.Syms.Lookup(".rel", 0)
			rel.AddAddrPlus(ctxt.Arch, s, int64(r.Off))
			rel.AddUint32(ctxt.Arch, ld.ELF32_R_INFO(uint32(targ.Dynid), uint32(elf.R_ARM_GLOB_DAT)))
			r.Type = objabi.R_CONST
			r.Sym = nil
			return true
		}
	}

	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	ctxt.Out.Write32(uint32(sectoff))

	elfsym := r.Xsym.ElfsymForReloc()
	switch r.Type {
	default:
		return false
	case objabi.R_ADDR:
		if r.Siz == 4 {
			ctxt.Out.Write32(uint32(elf.R_ARM_ABS32) | uint32(elfsym)<<8)
		} else {
			return false
		}
	case objabi.R_PCREL:
		if r.Siz == 4 {
			ctxt.Out.Write32(uint32(elf.R_ARM_REL32) | uint32(elfsym)<<8)
		} else {
			return false
		}
	case objabi.R_CALLARM:
		if r.Siz == 4 {
			if r.Add&0xff000000 == 0xeb000000 {
				ctxt.Out.Write32(uint32(elf.R_ARM_CALL) | uint32(elfsym)<<8)
			} else {
				ctxt.Out.Write32(uint32(elf.R_ARM_JUMP24) | uint32(elfsym)<<8)
			}
		} else {
			return false
		}
	case objabi.R_TLS_LE:
		ctxt.Out.Write32(uint32(elf.R_ARM_TLS_LE32) | uint32(elfsym)<<8)
	case objabi.R_TLS_IE:
		ctxt.Out.Write32(uint32(elf.R_ARM_TLS_IE32) | uint32(elfsym)<<8)
	case objabi.R_GOTPCREL:
		if r.Siz == 4 {
			ctxt.Out.Write32(uint32(elf.R_ARM_GOT_PREL) | uint32(elfsym)<<8)
		} else {
			return false
		}
	}

	return true
}

func elfsetupplt(ctxt *ld.Link) {
	plt := ctxt.Syms.Lookup(".plt", 0)
	got := ctxt.Syms.Lookup(".got.plt", 0)
	if plt.Size == 0 {

		plt.AddUint32(ctxt.Arch, 0xe52de004)

		plt.AddUint32(ctxt.Arch, 0xe59fe004)

		plt.AddUint32(ctxt.Arch, 0xe08fe00e)

		plt.AddUint32(ctxt.Arch, 0xe5bef008)

		plt.AddPCRelPlus(ctxt.Arch, got, 4)

		got.AddUint32(ctxt.Arch, 0)

		got.AddUint32(ctxt.Arch, 0)
		got.AddUint32(ctxt.Arch, 0)
	}
}

func (psess *PackageSession) machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	var v uint32

	rs := r.Xsym

	if r.Type == objabi.R_PCREL {
		if rs.Type == sym.SHOSTOBJ {
			psess.ld.
				Errorf(s, "pc-relative relocation of external symbol is not supported")
			return false
		}
		if r.Siz != 4 {
			return false
		}

		o1 := uint32(sectoff)
		o1 |= 1 << 31
		o1 |= ld.MACHO_ARM_RELOC_SECTDIFF << 24
		o1 |= 2 << 28

		o2 := uint32(0)
		o2 |= 1 << 31
		o2 |= ld.MACHO_ARM_RELOC_PAIR << 24
		o2 |= 2 << 28

		out.Write32(o1)
		out.Write32(uint32(psess.ld.Symaddr(rs)))
		out.Write32(o2)
		out.Write32(uint32(s.Value + int64(r.Off)))
		return true
	}

	if rs.Type == sym.SHOSTOBJ || r.Type == objabi.R_CALLARM {
		if rs.Dynid < 0 {
			psess.ld.
				Errorf(s, "reloc %d (%s) to non-macho symbol %s type=%d (%s)", r.Type, psess.sym.RelocName(arch, r.Type), rs.Name, rs.Type, rs.Type)
			return false
		}

		v = uint32(rs.Dynid)
		v |= 1 << 27
	} else {
		v = uint32(rs.Sect.Extnum)
		if v == 0 {
			psess.ld.
				Errorf(s, "reloc %d (%s) to symbol %s in non-macho section %s type=%d (%s)", r.Type, psess.sym.RelocName(arch, r.Type), rs.Name, rs.Sect.Name, rs.Type, rs.Type)
			return false
		}
	}

	switch r.Type {
	default:
		return false

	case objabi.R_ADDR:
		v |= ld.MACHO_GENERIC_RELOC_VANILLA << 28

	case objabi.R_CALLARM:
		v |= 1 << 24
		v |= ld.MACHO_ARM_RELOC_BR24 << 28
	}

	switch r.Siz {
	default:
		return false
	case 1:
		v |= 0 << 25

	case 2:
		v |= 1 << 25

	case 4:
		v |= 2 << 25

	case 8:
		v |= 3 << 25
	}

	out.Write32(uint32(sectoff))
	out.Write32(v)
	return true
}

// sign extend a 24-bit integer
func signext24(x int64) int32 {
	return (int32(x) << 8) >> 8
}

// encode an immediate in ARM's imm12 format. copied from ../../../internal/obj/arm/asm5.go
func immrot(v uint32) uint32 {
	for i := 0; i < 16; i++ {
		if v&^0xff == 0 {
			return uint32(i<<8) | v | 1<<25
		}
		v = v<<2 | v>>30
	}
	return 0
}

// Convert the direct jump relocation r to refer to a trampoline if the target is too far
func (psess *PackageSession) trampoline(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol) {
	switch r.Type {
	case objabi.R_CALLARM:

		t := (psess.ld.Symaddr(r.Sym) + int64(signext24(r.Add&0xffffff)*4) - (s.Value + int64(r.Off))) / 4
		if t > 0x7fffff || t < -0x800000 || (*psess.ld.FlagDebugTramp > 1 && s.File != r.Sym.File) {

			offset := (signext24(r.Add&0xffffff) + 2) * 4
			var tramp *sym.Symbol
			for i := 0; ; i++ {
				name := r.Sym.Name + fmt.Sprintf("%+d-tramp%d", offset, i)
				tramp = ctxt.Syms.Lookup(name, int(r.Sym.Version))
				if tramp.Type == sym.SDYNIMPORT {

					continue
				}
				if tramp.Value == 0 {

					break
				}

				t = (psess.ld.Symaddr(tramp) - 8 - (s.Value + int64(r.Off))) / 4
				if t >= -0x800000 && t < 0x7fffff {

					break
				}
			}
			if tramp.Type == 0 {

				ctxt.AddTramp(psess.ld, tramp)
				if ctxt.DynlinkingGo() {
					if immrot(uint32(offset)) == 0 {
						psess.ld.
							Errorf(s, "odd offset in dynlink direct call: %v+%d", r.Sym, offset)
					}
					gentrampdyn(ctxt.Arch, tramp, r.Sym, int64(offset))
				} else if ctxt.BuildMode == ld.BuildModeCArchive || ctxt.BuildMode == ld.BuildModeCShared || ctxt.BuildMode == ld.BuildModePIE {
					gentramppic(ctxt.Arch, tramp, r.Sym, int64(offset))
				} else {
					psess.
						gentramp(ctxt.Arch, ctxt.LinkMode, tramp, r.Sym, int64(offset))
				}
			}

			r.Sym = tramp
			r.Add = r.Add&0xff000000 | 0xfffffe
			r.Done = false
		}
	default:
		psess.ld.
			Errorf(s, "trampoline called with non-jump reloc: %d (%s)", r.Type, psess.sym.RelocName(ctxt.Arch, r.Type))
	}
}

// generate a trampoline to target+offset
func (psess *PackageSession) gentramp(arch *sys.Arch, linkmode ld.LinkMode, tramp, target *sym.Symbol, offset int64) {
	tramp.Size = 12
	tramp.P = make([]byte, tramp.Size)
	t := psess.ld.Symaddr(target) + offset
	o1 := uint32(0xe5900000 | 11<<12 | 15<<16)
	o2 := uint32(0xe12fff10 | 11)
	o3 := uint32(t)
	arch.ByteOrder.PutUint32(tramp.P, o1)
	arch.ByteOrder.PutUint32(tramp.P[4:], o2)
	arch.ByteOrder.PutUint32(tramp.P[8:], o3)

	if linkmode == ld.LinkExternal {
		r := tramp.AddRel()
		r.Off = 8
		r.Type = objabi.R_ADDR
		r.Siz = 4
		r.Sym = target
		r.Add = offset
	}
}

// generate a trampoline to target+offset in position independent code
func gentramppic(arch *sys.Arch, tramp, target *sym.Symbol, offset int64) {
	tramp.Size = 16
	tramp.P = make([]byte, tramp.Size)
	o1 := uint32(0xe5900000 | 11<<12 | 15<<16 | 4)
	o2 := uint32(0xe0800000 | 11<<12 | 15<<16 | 11)
	o3 := uint32(0xe12fff10 | 11)
	o4 := uint32(0)
	arch.ByteOrder.PutUint32(tramp.P, o1)
	arch.ByteOrder.PutUint32(tramp.P[4:], o2)
	arch.ByteOrder.PutUint32(tramp.P[8:], o3)
	arch.ByteOrder.PutUint32(tramp.P[12:], o4)

	r := tramp.AddRel()
	r.Off = 12
	r.Type = objabi.R_PCREL
	r.Siz = 4
	r.Sym = target
	r.Add = offset + 4
}

// generate a trampoline to target+offset in dynlink mode (using GOT)
func gentrampdyn(arch *sys.Arch, tramp, target *sym.Symbol, offset int64) {
	tramp.Size = 20
	o1 := uint32(0xe5900000 | 11<<12 | 15<<16 | 8)
	o2 := uint32(0xe0800000 | 11<<12 | 15<<16 | 11)
	o3 := uint32(0xe5900000 | 11<<12 | 11<<16)
	o4 := uint32(0xe12fff10 | 11)
	o5 := uint32(0)
	o6 := uint32(0)
	if offset != 0 {

		tramp.Size = 24
		o6 = o5
		o5 = o4
		o4 = 0xe2800000 | 11<<12 | 11<<16 | immrot(uint32(offset))
		o1 = uint32(0xe5900000 | 11<<12 | 15<<16 | 12)
	}
	tramp.P = make([]byte, tramp.Size)
	arch.ByteOrder.PutUint32(tramp.P, o1)
	arch.ByteOrder.PutUint32(tramp.P[4:], o2)
	arch.ByteOrder.PutUint32(tramp.P[8:], o3)
	arch.ByteOrder.PutUint32(tramp.P[12:], o4)
	arch.ByteOrder.PutUint32(tramp.P[16:], o5)
	if offset != 0 {
		arch.ByteOrder.PutUint32(tramp.P[20:], o6)
	}

	r := tramp.AddRel()
	r.Off = 16
	r.Type = objabi.R_GOTPCREL
	r.Siz = 4
	r.Sym = target
	r.Add = 8
	if offset != 0 {

		r.Off = 20
		r.Add = 12
	}
}

func (psess *PackageSession) archreloc(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val *int64) bool {
	if ctxt.LinkMode == ld.LinkExternal {
		switch r.Type {
		case objabi.R_CALLARM:
			r.Done = false

			rs := r.Sym

			r.Xadd = int64(signext24(r.Add & 0xffffff))
			r.Xadd *= 4
			for rs.Outer != nil {
				r.Xadd += psess.ld.Symaddr(rs) - psess.ld.Symaddr(rs.Outer)
				rs = rs.Outer
			}

			if rs.Type != sym.SHOSTOBJ && rs.Type != sym.SDYNIMPORT && rs.Sect == nil {
				psess.ld.
					Errorf(s, "missing section for %s", rs.Name)
			}
			r.Xsym = rs

			if ctxt.HeadType == objabi.Hdarwin {
				r.Xadd -= psess.ld.Symaddr(s) + int64(r.Off)
			}

			if r.Xadd/4 > 0x7fffff || r.Xadd/4 < -0x800000 {
				psess.ld.
					Errorf(s, "direct call too far %d", r.Xadd/4)
			}

			*val = int64(braddoff(int32(0xff000000&uint32(r.Add)), int32(0xffffff&uint32(r.Xadd/4))))
			return true
		}

		return false
	}

	switch r.Type {
	case objabi.R_CONST:
		*val = r.Add
		return true
	case objabi.R_GOTOFF:
		*val = psess.ld.Symaddr(r.Sym) + r.Add - psess.ld.Symaddr(ctxt.Syms.Lookup(".got", 0))
		return true

	case objabi.R_PLT0:
		if psess.ld.Symaddr(ctxt.Syms.Lookup(".got.plt", 0)) < psess.ld.Symaddr(ctxt.Syms.Lookup(".plt", 0)) {
			psess.ld.
				Errorf(s, ".got.plt should be placed after .plt section.")
		}
		*val = 0xe28fc600 + (0xff & (int64(uint32(psess.ld.Symaddr(r.Sym)-(psess.ld.Symaddr(ctxt.Syms.Lookup(".plt", 0))+int64(r.Off))+r.Add)) >> 20))
		return true
	case objabi.R_PLT1:
		*val = 0xe28cca00 + (0xff & (int64(uint32(psess.ld.Symaddr(r.Sym)-(psess.ld.Symaddr(ctxt.Syms.Lookup(".plt", 0))+int64(r.Off))+r.Add+4)) >> 12))

		return true
	case objabi.R_PLT2:
		*val = 0xe5bcf000 + (0xfff & int64(uint32(psess.ld.Symaddr(r.Sym)-(psess.ld.Symaddr(ctxt.Syms.Lookup(".plt", 0))+int64(r.Off))+r.Add+8)))

		return true
	case objabi.R_CALLARM:

		t := (psess.ld.Symaddr(r.Sym) + int64(signext24(r.Add&0xffffff)*4) - (s.Value + int64(r.Off))) / 4
		if t > 0x7fffff || t < -0x800000 {
			psess.ld.
				Errorf(s, "direct call too far: %s %x", r.Sym.Name, t)
		}
		*val = int64(braddoff(int32(0xff000000&uint32(r.Add)), int32(0xffffff&t)))

		return true
	}

	return false
}

func archrelocvariant(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	log.Fatalf("unexpected relocation variant")
	return t
}

func addpltreloc(ctxt *ld.Link, plt *sym.Symbol, got *sym.Symbol, s *sym.Symbol, typ objabi.RelocType) {
	r := plt.AddRel()
	r.Sym = got
	r.Off = int32(plt.Size)
	r.Siz = 4
	r.Type = typ
	r.Add = int64(s.Got) - 8

	plt.Attr |= sym.AttrReachable
	plt.Size += 4
	plt.Grow(plt.Size)
}

func (psess *PackageSession) addpltsym(ctxt *ld.Link, s *sym.Symbol) {
	if s.Plt >= 0 {
		return
	}
	psess.ld.
		Adddynsym(ctxt, s)

	if ctxt.IsELF {
		plt := ctxt.Syms.Lookup(".plt", 0)
		got := ctxt.Syms.Lookup(".got.plt", 0)
		rel := ctxt.Syms.Lookup(".rel.plt", 0)
		if plt.Size == 0 {
			elfsetupplt(ctxt)
		}

		s.Got = int32(got.Size)

		got.AddAddrPlus(ctxt.Arch, plt, 0)

		s.Plt = int32(plt.Size)

		addpltreloc(ctxt, plt, got, s, objabi.R_PLT0)
		addpltreloc(ctxt, plt, got, s, objabi.R_PLT1)
		addpltreloc(ctxt, plt, got, s, objabi.R_PLT2)

		rel.AddAddrPlus(ctxt.Arch, got, int64(s.Got))

		rel.AddUint32(ctxt.Arch, ld.ELF32_R_INFO(uint32(s.Dynid), uint32(elf.R_ARM_JUMP_SLOT)))
	} else {
		psess.ld.
			Errorf(s, "addpltsym: unsupported binary format")
	}
}

func (psess *PackageSession) addgotsyminternal(ctxt *ld.Link, s *sym.Symbol) {
	if s.Got >= 0 {
		return
	}

	got := ctxt.Syms.Lookup(".got", 0)
	s.Got = int32(got.Size)

	got.AddAddrPlus(ctxt.Arch, s, 0)

	if ctxt.IsELF {
	} else {
		psess.ld.
			Errorf(s, "addgotsyminternal: unsupported binary format")
	}
}

func (psess *PackageSession) addgotsym(ctxt *ld.Link, s *sym.Symbol) {
	if s.Got >= 0 {
		return
	}
	psess.ld.
		Adddynsym(ctxt, s)
	got := ctxt.Syms.Lookup(".got", 0)
	s.Got = int32(got.Size)
	got.AddUint32(ctxt.Arch, 0)

	if ctxt.IsELF {
		rel := ctxt.Syms.Lookup(".rel", 0)
		rel.AddAddrPlus(ctxt.Arch, got, int64(s.Got))
		rel.AddUint32(ctxt.Arch, ld.ELF32_R_INFO(uint32(s.Dynid), uint32(elf.R_ARM_GLOB_DAT)))
	} else {
		psess.ld.
			Errorf(s, "addgotsym: unsupported binary format")
	}
}

func (psess *PackageSession) asmb(ctxt *ld.Link) {
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f asmb\n", psess.ld.Cputime())
	}

	if ctxt.IsELF {
		psess.ld.
			Asmbelfsetup()
	}

	sect := psess.ld.Segtext.Sections[0]
	ctxt.Out.SeekSet(psess.ld, int64(sect.Vaddr-psess.ld.Segtext.Vaddr+psess.ld.Segtext.Fileoff))
	psess.ld.
		Codeblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	for _, sect = range psess.ld.Segtext.Sections[1:] {
		ctxt.Out.SeekSet(psess.ld, int64(sect.Vaddr-psess.ld.Segtext.Vaddr+psess.ld.Segtext.Fileoff))
		psess.ld.
			Datblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
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

	machlink := uint32(0)
	if ctxt.HeadType == objabi.Hdarwin {
		machlink = uint32(psess.ld.Domacholink(ctxt))
	}
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

		case objabi.Hdarwin:
			symo = uint32(psess.ld.Segdwarf.Fileoff + uint64(ld.Rnd(int64(psess.ld.Segdwarf.Filelen), int64(*psess.ld.FlagRound))) + uint64(machlink))
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

		case objabi.Hdarwin:
			if ctxt.LinkMode == ld.LinkExternal {
				psess.ld.
					Machoemitreloc(ctxt)
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
		ctxt.Out.Write32b(0x647)
		ctxt.Out.Write32b(uint32(psess.ld.Segtext.Filelen))
		ctxt.Out.Write32b(uint32(psess.ld.Segdata.Filelen))
		ctxt.Out.Write32b(uint32(psess.ld.Segdata.Length - psess.ld.Segdata.Filelen))
		ctxt.Out.Write32b(uint32(psess.ld.Symsize))
		ctxt.Out.Write32b(uint32(psess.ld.Entryvalue(ctxt)))
		ctxt.Out.Write32b(0)
		ctxt.Out.Write32b(uint32(psess.ld.Lcsize))

	case objabi.Hlinux,
		objabi.Hfreebsd,
		objabi.Hnetbsd,
		objabi.Hopenbsd,
		objabi.Hnacl:
		psess.ld.
			Asmbelf(ctxt, int64(symo))

	case objabi.Hdarwin:
		psess.ld.
			Asmbmacho(ctxt)
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

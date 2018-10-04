// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"sort"
	"strings"
)

type MachoHdr struct {
	cpu    uint32
	subcpu uint32
}

type MachoSect struct {
	name    string
	segname string
	addr    uint64
	size    uint64
	off     uint32
	align   uint32
	reloc   uint32
	nreloc  uint32
	flag    uint32
	res1    uint32
	res2    uint32
}

type MachoSeg struct {
	name       string
	vsize      uint64
	vaddr      uint64
	fileoffset uint64
	filesize   uint64
	prot1      uint32
	prot2      uint32
	nsect      uint32
	msect      uint32
	sect       []MachoSect
	flag       uint32
}

type MachoLoad struct {
	type_ uint32
	data  []uint32
}

/*
 * Total amount of space to reserve at the start of the file
 * for Header, PHeaders, and SHeaders.
 * May waste some.
 */
const (
	INITIAL_MACHO_HEADR = 4 * 1024
)

const (
	MACHO_CPU_AMD64               = 1<<24 | 7
	MACHO_CPU_386                 = 7
	MACHO_SUBCPU_X86              = 3
	MACHO_CPU_ARM                 = 12
	MACHO_SUBCPU_ARM              = 0
	MACHO_SUBCPU_ARMV7            = 9
	MACHO_CPU_ARM64               = 1<<24 | 12
	MACHO_SUBCPU_ARM64_ALL        = 0
	MACHO32SYMSIZE                = 12
	MACHO64SYMSIZE                = 16
	MACHO_X86_64_RELOC_UNSIGNED   = 0
	MACHO_X86_64_RELOC_SIGNED     = 1
	MACHO_X86_64_RELOC_BRANCH     = 2
	MACHO_X86_64_RELOC_GOT_LOAD   = 3
	MACHO_X86_64_RELOC_GOT        = 4
	MACHO_X86_64_RELOC_SUBTRACTOR = 5
	MACHO_X86_64_RELOC_SIGNED_1   = 6
	MACHO_X86_64_RELOC_SIGNED_2   = 7
	MACHO_X86_64_RELOC_SIGNED_4   = 8
	MACHO_ARM_RELOC_VANILLA       = 0
	MACHO_ARM_RELOC_PAIR          = 1
	MACHO_ARM_RELOC_SECTDIFF      = 2
	MACHO_ARM_RELOC_BR24          = 5
	MACHO_ARM64_RELOC_UNSIGNED    = 0
	MACHO_ARM64_RELOC_BRANCH26    = 2
	MACHO_ARM64_RELOC_PAGE21      = 3
	MACHO_ARM64_RELOC_PAGEOFF12   = 4
	MACHO_ARM64_RELOC_ADDEND      = 10
	MACHO_GENERIC_RELOC_VANILLA   = 0
	MACHO_FAKE_GOTPCREL           = 100
)

const (
	MH_MAGIC    = 0xfeedface
	MH_MAGIC_64 = 0xfeedfacf

	MH_OBJECT  = 0x1
	MH_EXECUTE = 0x2

	MH_NOUNDEFS = 0x1
)

const (
	LC_SEGMENT                  = 0x1
	LC_SYMTAB                   = 0x2
	LC_SYMSEG                   = 0x3
	LC_THREAD                   = 0x4
	LC_UNIXTHREAD               = 0x5
	LC_LOADFVMLIB               = 0x6
	LC_IDFVMLIB                 = 0x7
	LC_IDENT                    = 0x8
	LC_FVMFILE                  = 0x9
	LC_PREPAGE                  = 0xa
	LC_DYSYMTAB                 = 0xb
	LC_LOAD_DYLIB               = 0xc
	LC_ID_DYLIB                 = 0xd
	LC_LOAD_DYLINKER            = 0xe
	LC_ID_DYLINKER              = 0xf
	LC_PREBOUND_DYLIB           = 0x10
	LC_ROUTINES                 = 0x11
	LC_SUB_FRAMEWORK            = 0x12
	LC_SUB_UMBRELLA             = 0x13
	LC_SUB_CLIENT               = 0x14
	LC_SUB_LIBRARY              = 0x15
	LC_TWOLEVEL_HINTS           = 0x16
	LC_PREBIND_CKSUM            = 0x17
	LC_LOAD_WEAK_DYLIB          = 0x18
	LC_SEGMENT_64               = 0x19
	LC_ROUTINES_64              = 0x1a
	LC_UUID                     = 0x1b
	LC_RPATH                    = 0x8000001c
	LC_CODE_SIGNATURE           = 0x1d
	LC_SEGMENT_SPLIT_INFO       = 0x1e
	LC_REEXPORT_DYLIB           = 0x8000001f
	LC_LAZY_LOAD_DYLIB          = 0x20
	LC_ENCRYPTION_INFO          = 0x21
	LC_DYLD_INFO                = 0x22
	LC_DYLD_INFO_ONLY           = 0x80000022
	LC_LOAD_UPWARD_DYLIB        = 0x80000023
	LC_VERSION_MIN_MACOSX       = 0x24
	LC_VERSION_MIN_IPHONEOS     = 0x25
	LC_FUNCTION_STARTS          = 0x26
	LC_DYLD_ENVIRONMENT         = 0x27
	LC_MAIN                     = 0x80000028
	LC_DATA_IN_CODE             = 0x29
	LC_SOURCE_VERSION           = 0x2A
	LC_DYLIB_CODE_SIGN_DRS      = 0x2B
	LC_ENCRYPTION_INFO_64       = 0x2C
	LC_LINKER_OPTION            = 0x2D
	LC_LINKER_OPTIMIZATION_HINT = 0x2E
	LC_VERSION_MIN_TVOS         = 0x2F
	LC_VERSION_MIN_WATCHOS      = 0x30
	LC_VERSION_NOTE             = 0x31
	LC_BUILD_VERSION            = 0x32
)

const (
	S_REGULAR                  = 0x0
	S_ZEROFILL                 = 0x1
	S_NON_LAZY_SYMBOL_POINTERS = 0x6
	S_SYMBOL_STUBS             = 0x8
	S_MOD_INIT_FUNC_POINTERS   = 0x9
	S_ATTR_PURE_INSTRUCTIONS   = 0x80000000
	S_ATTR_DEBUG               = 0x02000000
	S_ATTR_SOME_INSTRUCTIONS   = 0x00000400
)

const (
	SymKindLocal = 0 + iota
	SymKindExtdef
	SymKindUndef
	NumSymKind
)

func (pstate *PackageState) getMachoHdr() *MachoHdr {
	return &pstate.machohdr
}

func (pstate *PackageState) newMachoLoad(arch *sys.Arch, type_ uint32, ndata uint32) *MachoLoad {
	if arch.PtrSize == 8 && (ndata&1 != 0) {
		ndata++
	}

	pstate.load = append(pstate.load, MachoLoad{})
	l := &pstate.load[len(pstate.load)-1]
	l.type_ = type_
	l.data = make([]uint32, ndata)
	return l
}

func (pstate *PackageState) newMachoSeg(name string, msect int) *MachoSeg {
	if pstate.nseg >= len(pstate.seg) {
		pstate.Exitf("too many segs")
	}

	s := &pstate.seg[pstate.nseg]
	pstate.nseg++
	s.name = name
	s.msect = uint32(msect)
	s.sect = make([]MachoSect, msect)
	return s
}

func (pstate *PackageState) newMachoSect(seg *MachoSeg, name string, segname string) *MachoSect {
	if seg.nsect >= seg.msect {
		pstate.Exitf("too many sects in segment %s", seg.name)
	}

	s := &seg.sect[seg.nsect]
	seg.nsect++
	s.name = name
	s.segname = segname
	pstate.nsect++
	return s
}

func (pstate *PackageState) machowrite(arch *sys.Arch, out *OutBuf, linkmode LinkMode) int {
	o1 := out.Offset()

	loadsize := 4 * 4 * pstate.ndebug
	for i := range pstate.load {
		loadsize += 4 * (len(pstate.load[i].data) + 2)
	}
	if arch.PtrSize == 8 {
		loadsize += 18 * 4 * pstate.nseg
		loadsize += 20 * 4 * pstate.nsect
	} else {
		loadsize += 14 * 4 * pstate.nseg
		loadsize += 17 * 4 * pstate.nsect
	}

	if arch.PtrSize == 8 {
		out.Write32(MH_MAGIC_64)
	} else {
		out.Write32(MH_MAGIC)
	}
	out.Write32(pstate.machohdr.cpu)
	out.Write32(pstate.machohdr.subcpu)
	if linkmode == LinkExternal {
		out.Write32(MH_OBJECT) /* file type - mach object */
	} else {
		out.Write32(MH_EXECUTE) /* file type - mach executable */
	}
	out.Write32(uint32(len(pstate.load)) + uint32(pstate.nseg) + uint32(pstate.ndebug))
	out.Write32(uint32(loadsize))
	if pstate.nkind[SymKindUndef] == 0 {
		out.Write32(MH_NOUNDEFS) /* flags - no undefines */
	} else {
		out.Write32(0) /* flags */
	}
	if arch.PtrSize == 8 {
		out.Write32(0) /* reserved */
	}

	for i := 0; i < pstate.nseg; i++ {
		s := &pstate.seg[i]
		if arch.PtrSize == 8 {
			out.Write32(LC_SEGMENT_64)
			out.Write32(72 + 80*s.nsect)
			out.WriteStringN(pstate, s.name, 16)
			out.Write64(s.vaddr)
			out.Write64(s.vsize)
			out.Write64(s.fileoffset)
			out.Write64(s.filesize)
			out.Write32(s.prot1)
			out.Write32(s.prot2)
			out.Write32(s.nsect)
			out.Write32(s.flag)
		} else {
			out.Write32(LC_SEGMENT)
			out.Write32(56 + 68*s.nsect)
			out.WriteStringN(pstate, s.name, 16)
			out.Write32(uint32(s.vaddr))
			out.Write32(uint32(s.vsize))
			out.Write32(uint32(s.fileoffset))
			out.Write32(uint32(s.filesize))
			out.Write32(s.prot1)
			out.Write32(s.prot2)
			out.Write32(s.nsect)
			out.Write32(s.flag)
		}

		for j := uint32(0); j < s.nsect; j++ {
			t := &s.sect[j]
			if arch.PtrSize == 8 {
				out.WriteStringN(pstate, t.name, 16)
				out.WriteStringN(pstate, t.segname, 16)
				out.Write64(t.addr)
				out.Write64(t.size)
				out.Write32(t.off)
				out.Write32(t.align)
				out.Write32(t.reloc)
				out.Write32(t.nreloc)
				out.Write32(t.flag)
				out.Write32(t.res1) /* reserved */
				out.Write32(t.res2) /* reserved */
				out.Write32(0)      /* reserved */
			} else {
				out.WriteStringN(pstate, t.name, 16)
				out.WriteStringN(pstate, t.segname, 16)
				out.Write32(uint32(t.addr))
				out.Write32(uint32(t.size))
				out.Write32(t.off)
				out.Write32(t.align)
				out.Write32(t.reloc)
				out.Write32(t.nreloc)
				out.Write32(t.flag)
				out.Write32(t.res1) /* reserved */
				out.Write32(t.res2) /* reserved */
			}
		}
	}

	for i := range pstate.load {
		l := &pstate.load[i]
		out.Write32(l.type_)
		out.Write32(4 * (uint32(len(l.data)) + 2))
		for j := 0; j < len(l.data); j++ {
			out.Write32(l.data[j])
		}
	}

	return int(out.Offset() - o1)
}

func (ctxt *Link) domacho(pstate *PackageState) {
	if *pstate.FlagD {
		return
	}

	// empirically, string table must begin with " \x00".
	s := ctxt.Syms.Lookup(".machosymstr", 0)

	s.Type = sym.SMACHOSYMSTR
	s.Attr |= sym.AttrReachable
	s.AddUint8(' ')
	s.AddUint8('\x00')

	s = ctxt.Syms.Lookup(".machosymtab", 0)
	s.Type = sym.SMACHOSYMTAB
	s.Attr |= sym.AttrReachable

	if ctxt.LinkMode != LinkExternal {
		s := ctxt.Syms.Lookup(".plt", 0) // will be __symbol_stub
		s.Type = sym.SMACHOPLT
		s.Attr |= sym.AttrReachable

		s = ctxt.Syms.Lookup(".got", 0) // will be __nl_symbol_ptr
		s.Type = sym.SMACHOGOT
		s.Attr |= sym.AttrReachable
		s.Align = 4

		s = ctxt.Syms.Lookup(".linkedit.plt", 0) // indirect table for .plt
		s.Type = sym.SMACHOINDIRECTPLT
		s.Attr |= sym.AttrReachable

		s = ctxt.Syms.Lookup(".linkedit.got", 0) // indirect table for .got
		s.Type = sym.SMACHOINDIRECTGOT
		s.Attr |= sym.AttrReachable
	}
}

func (pstate *PackageState) machoadddynlib(lib string, linkmode LinkMode) {
	if pstate.seenlib[lib] || linkmode == LinkExternal {
		return
	}
	pstate.seenlib[lib] = true

	// Will need to store the library name rounded up
	// and 24 bytes of header metadata. If not enough
	// space, grab another page of initial space at the
	// beginning of the output file.
	pstate.loadBudget -= (len(lib)+7)/8*8 + 24

	if pstate.loadBudget < 0 {
		pstate.HEADR += 4096
		*pstate.FlagTextAddr += 4096
		pstate.loadBudget += 4096
	}

	pstate.dylib = append(pstate.dylib, lib)
}

func (pstate *PackageState) machoshbits(ctxt *Link, mseg *MachoSeg, sect *sym.Section, segname string) {
	buf := "__" + strings.Replace(sect.Name[1:], ".", "_", -1)

	var msect *MachoSect
	if sect.Rwx&1 == 0 && segname != "__DWARF" && (ctxt.Arch.Family == sys.ARM64 ||
		(ctxt.Arch.Family == sys.AMD64 && ctxt.BuildMode != BuildModeExe) ||
		(ctxt.Arch.Family == sys.ARM && ctxt.BuildMode != BuildModeExe)) {
		// Darwin external linker on arm64 and on amd64 and arm in c-shared/c-archive buildmode
		// complains about absolute relocs in __TEXT, so if the section is not
		// executable, put it in __DATA segment.
		msect = pstate.newMachoSect(mseg, buf, "__DATA")
	} else {
		msect = pstate.newMachoSect(mseg, buf, segname)
	}

	if sect.Rellen > 0 {
		msect.reloc = uint32(sect.Reloff)
		msect.nreloc = uint32(sect.Rellen / 8)
	}

	for 1<<msect.align < sect.Align {
		msect.align++
	}
	msect.addr = sect.Vaddr
	msect.size = sect.Length

	if sect.Vaddr < sect.Seg.Vaddr+sect.Seg.Filelen {
		// data in file
		if sect.Length > sect.Seg.Vaddr+sect.Seg.Filelen-sect.Vaddr {
			pstate.Errorf(nil, "macho cannot represent section %s crossing data and bss", sect.Name)
		}
		msect.off = uint32(sect.Seg.Fileoff + sect.Vaddr - sect.Seg.Vaddr)
	} else {
		msect.off = 0
		msect.flag |= S_ZEROFILL
	}

	if sect.Rwx&1 != 0 {
		msect.flag |= S_ATTR_SOME_INSTRUCTIONS
	}

	if sect.Name == ".text" {
		msect.flag |= S_ATTR_PURE_INSTRUCTIONS
	}

	if sect.Name == ".plt" {
		msect.name = "__symbol_stub1"
		msect.flag = S_ATTR_PURE_INSTRUCTIONS | S_ATTR_SOME_INSTRUCTIONS | S_SYMBOL_STUBS
		msect.res1 = 0 //nkind[SymKindLocal];
		msect.res2 = 6
	}

	if sect.Name == ".got" {
		msect.name = "__nl_symbol_ptr"
		msect.flag = S_NON_LAZY_SYMBOL_POINTERS
		msect.res1 = uint32(ctxt.Syms.Lookup(".linkedit.plt", 0).Size / 4) /* offset into indirect symbol table */
	}

	if sect.Name == ".init_array" {
		msect.name = "__mod_init_func"
		msect.flag = S_MOD_INIT_FUNC_POINTERS
	}

	if segname == "__DWARF" {
		msect.flag |= S_ATTR_DEBUG
	}
}

func (pstate *PackageState) Asmbmacho(ctxt *Link) {
	/* apple MACH */
	va := *pstate.FlagTextAddr - int64(pstate.HEADR)

	mh := pstate.getMachoHdr()
	switch ctxt.Arch.Family {
	default:
		pstate.Exitf("unknown macho architecture: %v", ctxt.Arch.Family)

	case sys.ARM:
		mh.cpu = MACHO_CPU_ARM
		mh.subcpu = MACHO_SUBCPU_ARMV7

	case sys.AMD64:
		mh.cpu = MACHO_CPU_AMD64
		mh.subcpu = MACHO_SUBCPU_X86

	case sys.ARM64:
		mh.cpu = MACHO_CPU_ARM64
		mh.subcpu = MACHO_SUBCPU_ARM64_ALL

	case sys.I386:
		mh.cpu = MACHO_CPU_386
		mh.subcpu = MACHO_SUBCPU_X86
	}

	var ms *MachoSeg
	if ctxt.LinkMode == LinkExternal {
		/* segment for entire file */
		ms = pstate.newMachoSeg("", 40)

		ms.fileoffset = pstate.Segtext.Fileoff
		if ctxt.Arch.Family == sys.ARM || ctxt.BuildMode == BuildModeCArchive {
			ms.filesize = pstate.Segdata.Fileoff + pstate.Segdata.Filelen - pstate.Segtext.Fileoff
		} else {
			ms.filesize = pstate.Segdwarf.Fileoff + pstate.Segdwarf.Filelen - pstate.Segtext.Fileoff
			ms.vsize = pstate.Segdwarf.Vaddr + pstate.Segdwarf.Length - pstate.Segtext.Vaddr
		}
	}

	/* segment for zero page */
	if ctxt.LinkMode != LinkExternal {
		ms = pstate.newMachoSeg("__PAGEZERO", 0)
		ms.vsize = uint64(va)
	}

	/* text */
	v := Rnd(int64(uint64(pstate.HEADR)+pstate.Segtext.Length), int64(*pstate.FlagRound))

	if ctxt.LinkMode != LinkExternal {
		ms = pstate.newMachoSeg("__TEXT", 20)
		ms.vaddr = uint64(va)
		ms.vsize = uint64(v)
		ms.fileoffset = 0
		ms.filesize = uint64(v)
		ms.prot1 = 7
		ms.prot2 = 5
	}

	for _, sect := range pstate.Segtext.Sections {
		pstate.machoshbits(ctxt, ms, sect, "__TEXT")
	}

	/* data */
	if ctxt.LinkMode != LinkExternal {
		w := int64(pstate.Segdata.Length)
		ms = pstate.newMachoSeg("__DATA", 20)
		ms.vaddr = uint64(va) + uint64(v)
		ms.vsize = uint64(w)
		ms.fileoffset = uint64(v)
		ms.filesize = pstate.Segdata.Filelen
		ms.prot1 = 3
		ms.prot2 = 3
	}

	for _, sect := range pstate.Segdata.Sections {
		pstate.machoshbits(ctxt, ms, sect, "__DATA")
	}

	/* dwarf */
	if !*pstate.FlagW {
		if ctxt.LinkMode != LinkExternal {
			ms = pstate.newMachoSeg("__DWARF", 20)
			ms.vaddr = pstate.Segdwarf.Vaddr
			ms.vsize = 0
			ms.fileoffset = pstate.Segdwarf.Fileoff
			ms.filesize = pstate.Segdwarf.Filelen
		}
		for _, sect := range pstate.Segdwarf.Sections {
			pstate.machoshbits(ctxt, ms, sect, "__DWARF")
		}
	}

	if ctxt.LinkMode != LinkExternal {
		switch ctxt.Arch.Family {
		default:
			pstate.Exitf("unknown macho architecture: %v", ctxt.Arch.Family)

		case sys.ARM:
			ml := pstate.newMachoLoad(ctxt.Arch, LC_UNIXTHREAD, 17+2)
			ml.data[0] = 1                                  /* thread type */
			ml.data[1] = 17                                 /* word count */
			ml.data[2+15] = uint32(pstate.Entryvalue(ctxt)) /* start pc */

		case sys.AMD64:
			ml := pstate.newMachoLoad(ctxt.Arch, LC_UNIXTHREAD, 42+2)
			ml.data[0] = 4                                  /* thread type */
			ml.data[1] = 42                                 /* word count */
			ml.data[2+32] = uint32(pstate.Entryvalue(ctxt)) /* start pc */
			ml.data[2+32+1] = uint32(pstate.Entryvalue(ctxt) >> 32)

		case sys.ARM64:
			ml := pstate.newMachoLoad(ctxt.Arch, LC_UNIXTHREAD, 68+2)
			ml.data[0] = 6                                  /* thread type */
			ml.data[1] = 68                                 /* word count */
			ml.data[2+64] = uint32(pstate.Entryvalue(ctxt)) /* start pc */
			ml.data[2+64+1] = uint32(pstate.Entryvalue(ctxt) >> 32)

		case sys.I386:
			ml := pstate.newMachoLoad(ctxt.Arch, LC_UNIXTHREAD, 16+2)
			ml.data[0] = 1                                  /* thread type */
			ml.data[1] = 16                                 /* word count */
			ml.data[2+10] = uint32(pstate.Entryvalue(ctxt)) /* start pc */
		}
	}

	if !*pstate.FlagD {
		// must match domacholink below
		s1 := ctxt.Syms.Lookup(".machosymtab", 0)
		s2 := ctxt.Syms.Lookup(".linkedit.plt", 0)
		s3 := ctxt.Syms.Lookup(".linkedit.got", 0)
		s4 := ctxt.Syms.Lookup(".machosymstr", 0)

		if ctxt.LinkMode != LinkExternal {
			ms := pstate.newMachoSeg("__LINKEDIT", 0)
			ms.vaddr = uint64(va) + uint64(v) + uint64(Rnd(int64(pstate.Segdata.Length), int64(*pstate.FlagRound)))
			ms.vsize = uint64(s1.Size) + uint64(s2.Size) + uint64(s3.Size) + uint64(s4.Size)
			ms.fileoffset = uint64(pstate.linkoff)
			ms.filesize = ms.vsize
			ms.prot1 = 7
			ms.prot2 = 3
		}

		ml := pstate.newMachoLoad(ctxt.Arch, LC_SYMTAB, 4)
		ml.data[0] = uint32(pstate.linkoff)                               /* symoff */
		ml.data[1] = uint32(pstate.nsortsym)                              /* nsyms */
		ml.data[2] = uint32(pstate.linkoff + s1.Size + s2.Size + s3.Size) /* stroff */
		ml.data[3] = uint32(s4.Size)                                      /* strsize */

		pstate.machodysymtab(ctxt)

		if ctxt.LinkMode != LinkExternal {
			ml := pstate.newMachoLoad(ctxt.Arch, LC_LOAD_DYLINKER, 6)
			ml.data[0] = 12 /* offset to string */
			stringtouint32(ml.data[1:], "/usr/lib/dyld")

			for _, lib := range pstate.dylib {
				ml = pstate.newMachoLoad(ctxt.Arch, LC_LOAD_DYLIB, 4+(uint32(len(lib))+1+7)/8*2)
				ml.data[0] = 24 /* offset of string from beginning of load */
				ml.data[1] = 0  /* time stamp */
				ml.data[2] = 0  /* version */
				ml.data[3] = 0  /* compatibility version */
				stringtouint32(ml.data[4:], lib)
			}
		}
	}

	if ctxt.LinkMode == LinkInternal {
		// For lldb, must say LC_VERSION_MIN_MACOSX or else
		// it won't know that this Mach-O binary is from OS X
		// (could be iOS or WatchOS instead).
		// Go on iOS uses linkmode=external, and linkmode=external
		// adds this itself. So we only need this code for linkmode=internal
		// and we can assume OS X.
		//
		// See golang.org/issues/12941.
		ml := pstate.newMachoLoad(ctxt.Arch, LC_VERSION_MIN_MACOSX, 2)
		ml.data[0] = 10<<16 | 7<<8 | 0<<0 // OS X version 10.7.0
		ml.data[1] = 10<<16 | 7<<8 | 0<<0 // SDK 10.7.0
	}

	a := pstate.machowrite(ctxt.Arch, ctxt.Out, ctxt.LinkMode)
	if int32(a) > pstate.HEADR {
		pstate.Exitf("HEADR too small: %d > %d", a, pstate.HEADR)
	}
}

func symkind(s *sym.Symbol) int {
	if s.Type == sym.SDYNIMPORT {
		return SymKindUndef
	}
	if s.Attr.CgoExport() {
		return SymKindExtdef
	}
	return SymKindLocal
}

func (pstate *PackageState) addsym(ctxt *Link, s *sym.Symbol, name string, type_ SymbolType, addr int64, gotype *sym.Symbol) {
	if s == nil {
		return
	}

	switch type_ {
	default:
		return

	case DataSym, BSSSym, TextSym:
		break
	}

	if pstate.sortsym != nil {
		pstate.sortsym[pstate.nsortsym] = s
		pstate.nkind[symkind(s)]++
	}

	pstate.nsortsym++
}

type machoscmp []*sym.Symbol

func (x machoscmp) Len() int {
	return len(x)
}

func (x machoscmp) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x machoscmp) Less(i, j int) bool {
	s1 := x[i]
	s2 := x[j]

	k1 := symkind(s1)
	k2 := symkind(s2)
	if k1 != k2 {
		return k1 < k2
	}

	return s1.Extname < s2.Extname
}

func (pstate *PackageState) machogenasmsym(ctxt *Link) {
	pstate.genasmsym(ctxt, pstate.addsym)
	for _, s := range ctxt.Syms.Allsym {
		if s.Type == sym.SDYNIMPORT || s.Type == sym.SHOSTOBJ {
			if s.Attr.Reachable() {
				pstate.addsym(ctxt, s, "", DataSym, 0, nil)
			}
		}
	}
}

func (pstate *PackageState) machosymorder(ctxt *Link) {
	// On Mac OS X Mountain Lion, we must sort exported symbols
	// So we sort them here and pre-allocate dynid for them
	// See https://golang.org/issue/4029
	for i := range pstate.dynexp {
		pstate.dynexp[i].Attr |= sym.AttrReachable
	}
	pstate.machogenasmsym(ctxt)
	pstate.sortsym = make([]*sym.Symbol, pstate.nsortsym)
	pstate.nsortsym = 0
	pstate.machogenasmsym(ctxt)
	sort.Sort(machoscmp(pstate.sortsym[:pstate.nsortsym]))
	for i := 0; i < pstate.nsortsym; i++ {
		pstate.sortsym[i].Dynid = int32(i)
	}
}

// machoShouldExport reports whether a symbol needs to be exported.
//
// When dynamically linking, all non-local variables and plugin-exported
// symbols need to be exported.
func (pstate *PackageState) machoShouldExport(ctxt *Link, s *sym.Symbol) bool {
	if !ctxt.DynlinkingGo() || s.Attr.Local() {
		return false
	}
	if ctxt.BuildMode == BuildModePlugin && strings.HasPrefix(s.Extname, objabi.PathToPrefix(*pstate.flagPluginPath)) {
		return true
	}
	if strings.HasPrefix(s.Name, "go.itab.") {
		return true
	}
	if strings.HasPrefix(s.Name, "type.") && !strings.HasPrefix(s.Name, "type..") {
		// reduce runtime typemap pressure, but do not
		// export alg functions (type..*), as these
		// appear in pclntable.
		return true
	}
	if strings.HasPrefix(s.Name, "go.link.pkghash") {
		return true
	}
	return s.Type >= sym.SELFSECT // only writable sections
}

func (pstate *PackageState) machosymtab(ctxt *Link) {
	symtab := ctxt.Syms.Lookup(".machosymtab", 0)
	symstr := ctxt.Syms.Lookup(".machosymstr", 0)

	for i := 0; i < pstate.nsortsym; i++ {
		s := pstate.sortsym[i]
		symtab.AddUint32(ctxt.Arch, uint32(symstr.Size))

		export := pstate.machoShouldExport(ctxt, s)

		// In normal buildmodes, only add _ to C symbols, as
		// Go symbols have dot in the name.
		//
		// Do not export C symbols in plugins, as runtime C
		// symbols like crosscall2 are in pclntab and end up
		// pointing at the host binary, breaking unwinding.
		// See Issue #18190.
		cexport := !strings.Contains(s.Extname, ".") && (ctxt.BuildMode != BuildModePlugin || onlycsymbol(s))
		if cexport || export {
			symstr.AddUint8('_')
		}

		// replace "·" as ".", because DTrace cannot handle it.
		pstate.Addstring(symstr, strings.Replace(s.Extname, "·", ".", -1))

		if s.Type == sym.SDYNIMPORT || s.Type == sym.SHOSTOBJ {
			symtab.AddUint8(0x01)                             // type N_EXT, external symbol
			symtab.AddUint8(0)                                // no section
			symtab.AddUint16(ctxt.Arch, 0)                    // desc
			symtab.AddUintXX(ctxt.Arch, 0, ctxt.Arch.PtrSize) // no value
		} else {
			if s.Attr.CgoExport() || export {
				symtab.AddUint8(0x0f)
			} else {
				symtab.AddUint8(0x0e)
			}
			o := s
			for o.Outer != nil {
				o = o.Outer
			}
			if o.Sect == nil {
				pstate.Errorf(s, "missing section for symbol")
				symtab.AddUint8(0)
			} else {
				symtab.AddUint8(uint8(o.Sect.Extnum))
			}
			symtab.AddUint16(ctxt.Arch, 0) // desc
			symtab.AddUintXX(ctxt.Arch, uint64(pstate.Symaddr(s)), ctxt.Arch.PtrSize)
		}
	}
}

func (pstate *PackageState) machodysymtab(ctxt *Link) {
	ml := pstate.newMachoLoad(ctxt.Arch, LC_DYSYMTAB, 18)

	n := 0
	ml.data[0] = uint32(n)                          /* ilocalsym */
	ml.data[1] = uint32(pstate.nkind[SymKindLocal]) /* nlocalsym */
	n += pstate.nkind[SymKindLocal]

	ml.data[2] = uint32(n)                           /* iextdefsym */
	ml.data[3] = uint32(pstate.nkind[SymKindExtdef]) /* nextdefsym */
	n += pstate.nkind[SymKindExtdef]

	ml.data[4] = uint32(n)                          /* iundefsym */
	ml.data[5] = uint32(pstate.nkind[SymKindUndef]) /* nundefsym */

	ml.data[6] = 0  /* tocoffset */
	ml.data[7] = 0  /* ntoc */
	ml.data[8] = 0  /* modtaboff */
	ml.data[9] = 0  /* nmodtab */
	ml.data[10] = 0 /* extrefsymoff */
	ml.data[11] = 0 /* nextrefsyms */

	// must match domacholink below
	s1 := ctxt.Syms.Lookup(".machosymtab", 0)

	s2 := ctxt.Syms.Lookup(".linkedit.plt", 0)
	s3 := ctxt.Syms.Lookup(".linkedit.got", 0)
	ml.data[12] = uint32(pstate.linkoff + s1.Size) /* indirectsymoff */
	ml.data[13] = uint32((s2.Size + s3.Size) / 4)  /* nindirectsyms */

	ml.data[14] = 0 /* extreloff */
	ml.data[15] = 0 /* nextrel */
	ml.data[16] = 0 /* locreloff */
	ml.data[17] = 0 /* nlocrel */
}

func (pstate *PackageState) Domacholink(ctxt *Link) int64 {
	pstate.machosymtab(ctxt)

	// write data that will be linkedit section
	s1 := ctxt.Syms.Lookup(".machosymtab", 0)

	s2 := ctxt.Syms.Lookup(".linkedit.plt", 0)
	s3 := ctxt.Syms.Lookup(".linkedit.got", 0)
	s4 := ctxt.Syms.Lookup(".machosymstr", 0)

	// Force the linkedit section to end on a 16-byte
	// boundary. This allows pure (non-cgo) Go binaries
	// to be code signed correctly.
	//
	// Apple's codesign_allocate (a helper utility for
	// the codesign utility) can do this fine itself if
	// it is run on a dynamic Mach-O binary. However,
	// when it is run on a pure (non-cgo) Go binary, where
	// the linkedit section is mostly empty, it fails to
	// account for the extra padding that it itself adds
	// when adding the LC_CODE_SIGNATURE load command
	// (which must be aligned on a 16-byte boundary).
	//
	// By forcing the linkedit section to end on a 16-byte
	// boundary, codesign_allocate will not need to apply
	// any alignment padding itself, working around the
	// issue.
	for s4.Size%16 != 0 {
		s4.AddUint8(0)
	}

	size := int(s1.Size + s2.Size + s3.Size + s4.Size)

	if size > 0 {
		pstate.linkoff = Rnd(int64(uint64(pstate.HEADR)+pstate.Segtext.Length), int64(*pstate.FlagRound)) + Rnd(int64(pstate.Segdata.Filelen), int64(*pstate.FlagRound)) + Rnd(int64(pstate.Segdwarf.Filelen), int64(*pstate.FlagRound))
		ctxt.Out.SeekSet(pstate, pstate.linkoff)

		ctxt.Out.Write(s1.P[:s1.Size])
		ctxt.Out.Write(s2.P[:s2.Size])
		ctxt.Out.Write(s3.P[:s3.Size])
		ctxt.Out.Write(s4.P[:s4.Size])
	}

	return Rnd(int64(size), int64(*pstate.FlagRound))
}

func (pstate *PackageState) machorelocsect(ctxt *Link, sect *sym.Section, syms []*sym.Symbol) {
	// If main section has no bits, nothing to relocate.
	if sect.Vaddr >= sect.Seg.Vaddr+sect.Seg.Filelen {
		return
	}

	sect.Reloff = uint64(ctxt.Out.Offset())
	for i, s := range syms {
		if !s.Attr.Reachable() {
			continue
		}
		if uint64(s.Value) >= sect.Vaddr {
			syms = syms[i:]
			break
		}
	}

	eaddr := int32(sect.Vaddr + sect.Length)
	for _, s := range syms {
		if !s.Attr.Reachable() {
			continue
		}
		if s.Value >= int64(eaddr) {
			break
		}
		for ri := range s.R {
			r := &s.R[ri]
			if r.Done {
				continue
			}
			if r.Xsym == nil {
				pstate.Errorf(s, "missing xsym in relocation")
				continue
			}
			if !r.Xsym.Attr.Reachable() {
				pstate.Errorf(s, "unreachable reloc %d (%s) target %v", r.Type, pstate.sym.RelocName(ctxt.Arch, r.Type), r.Xsym.Name)
			}
			if !pstate.thearch.Machoreloc1(ctxt.Arch, ctxt.Out, s, r, int64(uint64(s.Value+int64(r.Off))-sect.Vaddr)) {
				pstate.Errorf(s, "unsupported obj reloc %d (%s)/%d to %s", r.Type, pstate.sym.RelocName(ctxt.Arch, r.Type), r.Siz, r.Sym.Name)
			}
		}
	}

	sect.Rellen = uint64(ctxt.Out.Offset()) - sect.Reloff
}

func (pstate *PackageState) Machoemitreloc(ctxt *Link) {
	for ctxt.Out.Offset()&7 != 0 {
		ctxt.Out.Write8(0)
	}

	pstate.machorelocsect(ctxt, pstate.Segtext.Sections[0], ctxt.Textp)
	for _, sect := range pstate.Segtext.Sections[1:] {
		pstate.machorelocsect(ctxt, sect, pstate.datap)
	}
	for _, sect := range pstate.Segdata.Sections {
		pstate.machorelocsect(ctxt, sect, pstate.datap)
	}
	for _, sect := range pstate.Segdwarf.Sections {
		pstate.machorelocsect(ctxt, sect, pstate.dwarfp)
	}
}

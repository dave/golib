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

// Amount of space left for adding load commands
// that refer to dynamic libraries. Because these have
// to go in the Mach-O header, we can't just pick a
// "big enough" header size. The initial header is
// one page, the non-dynamic library stuff takes
// up about 1300 bytes; we overestimate that as 2k.

func (psess *PackageSession) getMachoHdr() *MachoHdr {
	return &psess.machohdr
}

func (psess *PackageSession) newMachoLoad(arch *sys.Arch, type_ uint32, ndata uint32) *MachoLoad {
	if arch.PtrSize == 8 && (ndata&1 != 0) {
		ndata++
	}
	psess.
		load = append(psess.load, MachoLoad{})
	l := &psess.load[len(psess.load)-1]
	l.type_ = type_
	l.data = make([]uint32, ndata)
	return l
}

func (psess *PackageSession) newMachoSeg(name string, msect int) *MachoSeg {
	if psess.nseg >= len(psess.seg) {
		psess.
			Exitf("too many segs")
	}

	s := &psess.seg[psess.nseg]
	psess.
		nseg++
	s.name = name
	s.msect = uint32(msect)
	s.sect = make([]MachoSect, msect)
	return s
}

func (psess *PackageSession) newMachoSect(seg *MachoSeg, name string, segname string) *MachoSect {
	if seg.nsect >= seg.msect {
		psess.
			Exitf("too many sects in segment %s", seg.name)
	}

	s := &seg.sect[seg.nsect]
	seg.nsect++
	s.name = name
	s.segname = segname
	psess.
		nsect++
	return s
}

func (psess *PackageSession) machowrite(arch *sys.Arch, out *OutBuf, linkmode LinkMode) int {
	o1 := out.Offset()

	loadsize := 4 * 4 * psess.ndebug
	for i := range psess.load {
		loadsize += 4 * (len(psess.load[i].data) + 2)
	}
	if arch.PtrSize == 8 {
		loadsize += 18 * 4 * psess.nseg
		loadsize += 20 * 4 * psess.nsect
	} else {
		loadsize += 14 * 4 * psess.nseg
		loadsize += 17 * 4 * psess.nsect
	}

	if arch.PtrSize == 8 {
		out.Write32(MH_MAGIC_64)
	} else {
		out.Write32(MH_MAGIC)
	}
	out.Write32(psess.machohdr.cpu)
	out.Write32(psess.machohdr.subcpu)
	if linkmode == LinkExternal {
		out.Write32(MH_OBJECT)
	} else {
		out.Write32(MH_EXECUTE)
	}
	out.Write32(uint32(len(psess.load)) + uint32(psess.nseg) + uint32(psess.ndebug))
	out.Write32(uint32(loadsize))
	if psess.nkind[SymKindUndef] == 0 {
		out.Write32(MH_NOUNDEFS)
	} else {
		out.Write32(0)
	}
	if arch.PtrSize == 8 {
		out.Write32(0)
	}

	for i := 0; i < psess.nseg; i++ {
		s := &psess.seg[i]
		if arch.PtrSize == 8 {
			out.Write32(LC_SEGMENT_64)
			out.Write32(72 + 80*s.nsect)
			out.WriteStringN(psess, s.name, 16)
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
			out.WriteStringN(psess, s.name, 16)
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
				out.WriteStringN(psess, t.name, 16)
				out.WriteStringN(psess, t.segname, 16)
				out.Write64(t.addr)
				out.Write64(t.size)
				out.Write32(t.off)
				out.Write32(t.align)
				out.Write32(t.reloc)
				out.Write32(t.nreloc)
				out.Write32(t.flag)
				out.Write32(t.res1)
				out.Write32(t.res2)
				out.Write32(0)
			} else {
				out.WriteStringN(psess, t.name, 16)
				out.WriteStringN(psess, t.segname, 16)
				out.Write32(uint32(t.addr))
				out.Write32(uint32(t.size))
				out.Write32(t.off)
				out.Write32(t.align)
				out.Write32(t.reloc)
				out.Write32(t.nreloc)
				out.Write32(t.flag)
				out.Write32(t.res1)
				out.Write32(t.res2)
			}
		}
	}

	for i := range psess.load {
		l := &psess.load[i]
		out.Write32(l.type_)
		out.Write32(4 * (uint32(len(l.data)) + 2))
		for j := 0; j < len(l.data); j++ {
			out.Write32(l.data[j])
		}
	}

	return int(out.Offset() - o1)
}

func (ctxt *Link) domacho(psess *PackageSession) {
	if *psess.FlagD {
		return
	}

	s := ctxt.Syms.Lookup(".machosymstr", 0)

	s.Type = sym.SMACHOSYMSTR
	s.Attr |= sym.AttrReachable
	s.AddUint8(' ')
	s.AddUint8('\x00')

	s = ctxt.Syms.Lookup(".machosymtab", 0)
	s.Type = sym.SMACHOSYMTAB
	s.Attr |= sym.AttrReachable

	if ctxt.LinkMode != LinkExternal {
		s := ctxt.Syms.Lookup(".plt", 0)
		s.Type = sym.SMACHOPLT
		s.Attr |= sym.AttrReachable

		s = ctxt.Syms.Lookup(".got", 0)
		s.Type = sym.SMACHOGOT
		s.Attr |= sym.AttrReachable
		s.Align = 4

		s = ctxt.Syms.Lookup(".linkedit.plt", 0)
		s.Type = sym.SMACHOINDIRECTPLT
		s.Attr |= sym.AttrReachable

		s = ctxt.Syms.Lookup(".linkedit.got", 0)
		s.Type = sym.SMACHOINDIRECTGOT
		s.Attr |= sym.AttrReachable
	}
}

func (psess *PackageSession) machoadddynlib(lib string, linkmode LinkMode) {
	if psess.seenlib[lib] || linkmode == LinkExternal {
		return
	}
	psess.
		seenlib[lib] = true
	psess.
		loadBudget -= (len(lib)+7)/8*8 + 24

	if psess.loadBudget < 0 {
		psess.
			HEADR += 4096
		*psess.FlagTextAddr += 4096
		psess.
			loadBudget += 4096
	}
	psess.
		dylib = append(psess.dylib, lib)
}

func (psess *PackageSession) machoshbits(ctxt *Link, mseg *MachoSeg, sect *sym.Section, segname string) {
	buf := "__" + strings.Replace(sect.Name[1:], ".", "_", -1)

	var msect *MachoSect
	if sect.Rwx&1 == 0 && segname != "__DWARF" && (ctxt.Arch.Family == sys.ARM64 ||
		(ctxt.Arch.Family == sys.AMD64 && ctxt.BuildMode != BuildModeExe) ||
		(ctxt.Arch.Family == sys.ARM && ctxt.BuildMode != BuildModeExe)) {

		msect = psess.newMachoSect(mseg, buf, "__DATA")
	} else {
		msect = psess.newMachoSect(mseg, buf, segname)
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

		if sect.Length > sect.Seg.Vaddr+sect.Seg.Filelen-sect.Vaddr {
			psess.
				Errorf(nil, "macho cannot represent section %s crossing data and bss", sect.Name)
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
		msect.res1 = 0
		msect.res2 = 6
	}

	if sect.Name == ".got" {
		msect.name = "__nl_symbol_ptr"
		msect.flag = S_NON_LAZY_SYMBOL_POINTERS
		msect.res1 = uint32(ctxt.Syms.Lookup(".linkedit.plt", 0).Size / 4)
	}

	if sect.Name == ".init_array" {
		msect.name = "__mod_init_func"
		msect.flag = S_MOD_INIT_FUNC_POINTERS
	}

	if segname == "__DWARF" {
		msect.flag |= S_ATTR_DEBUG
	}
}

func (psess *PackageSession) Asmbmacho(ctxt *Link) {

	va := *psess.FlagTextAddr - int64(psess.HEADR)

	mh := psess.getMachoHdr()
	switch ctxt.Arch.Family {
	default:
		psess.
			Exitf("unknown macho architecture: %v", ctxt.Arch.Family)

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

		ms = psess.newMachoSeg("", 40)

		ms.fileoffset = psess.Segtext.Fileoff
		if ctxt.Arch.Family == sys.ARM || ctxt.BuildMode == BuildModeCArchive {
			ms.filesize = psess.Segdata.Fileoff + psess.Segdata.Filelen - psess.Segtext.Fileoff
		} else {
			ms.filesize = psess.Segdwarf.Fileoff + psess.Segdwarf.Filelen - psess.Segtext.Fileoff
			ms.vsize = psess.Segdwarf.Vaddr + psess.Segdwarf.Length - psess.Segtext.Vaddr
		}
	}

	if ctxt.LinkMode != LinkExternal {
		ms = psess.newMachoSeg("__PAGEZERO", 0)
		ms.vsize = uint64(va)
	}

	v := Rnd(int64(uint64(psess.HEADR)+psess.Segtext.Length), int64(*psess.FlagRound))

	if ctxt.LinkMode != LinkExternal {
		ms = psess.newMachoSeg("__TEXT", 20)
		ms.vaddr = uint64(va)
		ms.vsize = uint64(v)
		ms.fileoffset = 0
		ms.filesize = uint64(v)
		ms.prot1 = 7
		ms.prot2 = 5
	}

	for _, sect := range psess.Segtext.Sections {
		psess.
			machoshbits(ctxt, ms, sect, "__TEXT")
	}

	if ctxt.LinkMode != LinkExternal {
		w := int64(psess.Segdata.Length)
		ms = psess.newMachoSeg("__DATA", 20)
		ms.vaddr = uint64(va) + uint64(v)
		ms.vsize = uint64(w)
		ms.fileoffset = uint64(v)
		ms.filesize = psess.Segdata.Filelen
		ms.prot1 = 3
		ms.prot2 = 3
	}

	for _, sect := range psess.Segdata.Sections {
		psess.
			machoshbits(ctxt, ms, sect, "__DATA")
	}

	if !*psess.FlagW {
		if ctxt.LinkMode != LinkExternal {
			ms = psess.newMachoSeg("__DWARF", 20)
			ms.vaddr = psess.Segdwarf.Vaddr
			ms.vsize = 0
			ms.fileoffset = psess.Segdwarf.Fileoff
			ms.filesize = psess.Segdwarf.Filelen
		}
		for _, sect := range psess.Segdwarf.Sections {
			psess.
				machoshbits(ctxt, ms, sect, "__DWARF")
		}
	}

	if ctxt.LinkMode != LinkExternal {
		switch ctxt.Arch.Family {
		default:
			psess.
				Exitf("unknown macho architecture: %v", ctxt.Arch.Family)

		case sys.ARM:
			ml := psess.newMachoLoad(ctxt.Arch, LC_UNIXTHREAD, 17+2)
			ml.data[0] = 1
			ml.data[1] = 17
			ml.data[2+15] = uint32(psess.Entryvalue(ctxt))

		case sys.AMD64:
			ml := psess.newMachoLoad(ctxt.Arch, LC_UNIXTHREAD, 42+2)
			ml.data[0] = 4
			ml.data[1] = 42
			ml.data[2+32] = uint32(psess.Entryvalue(ctxt))
			ml.data[2+32+1] = uint32(psess.Entryvalue(ctxt) >> 32)

		case sys.ARM64:
			ml := psess.newMachoLoad(ctxt.Arch, LC_UNIXTHREAD, 68+2)
			ml.data[0] = 6
			ml.data[1] = 68
			ml.data[2+64] = uint32(psess.Entryvalue(ctxt))
			ml.data[2+64+1] = uint32(psess.Entryvalue(ctxt) >> 32)

		case sys.I386:
			ml := psess.newMachoLoad(ctxt.Arch, LC_UNIXTHREAD, 16+2)
			ml.data[0] = 1
			ml.data[1] = 16
			ml.data[2+10] = uint32(psess.Entryvalue(ctxt))
		}
	}

	if !*psess.FlagD {

		s1 := ctxt.Syms.Lookup(".machosymtab", 0)
		s2 := ctxt.Syms.Lookup(".linkedit.plt", 0)
		s3 := ctxt.Syms.Lookup(".linkedit.got", 0)
		s4 := ctxt.Syms.Lookup(".machosymstr", 0)

		if ctxt.LinkMode != LinkExternal {
			ms := psess.newMachoSeg("__LINKEDIT", 0)
			ms.vaddr = uint64(va) + uint64(v) + uint64(Rnd(int64(psess.Segdata.Length), int64(*psess.FlagRound)))
			ms.vsize = uint64(s1.Size) + uint64(s2.Size) + uint64(s3.Size) + uint64(s4.Size)
			ms.fileoffset = uint64(psess.linkoff)
			ms.filesize = ms.vsize
			ms.prot1 = 7
			ms.prot2 = 3
		}

		ml := psess.newMachoLoad(ctxt.Arch, LC_SYMTAB, 4)
		ml.data[0] = uint32(psess.linkoff)
		ml.data[1] = uint32(psess.nsortsym)
		ml.data[2] = uint32(psess.linkoff + s1.Size + s2.Size + s3.Size)
		ml.data[3] = uint32(s4.Size)
		psess.
			machodysymtab(ctxt)

		if ctxt.LinkMode != LinkExternal {
			ml := psess.newMachoLoad(ctxt.Arch, LC_LOAD_DYLINKER, 6)
			ml.data[0] = 12
			stringtouint32(ml.data[1:], "/usr/lib/dyld")

			for _, lib := range psess.dylib {
				ml = psess.newMachoLoad(ctxt.Arch, LC_LOAD_DYLIB, 4+(uint32(len(lib))+1+7)/8*2)
				ml.data[0] = 24
				ml.data[1] = 0
				ml.data[2] = 0
				ml.data[3] = 0
				stringtouint32(ml.data[4:], lib)
			}
		}
	}

	if ctxt.LinkMode == LinkInternal {

		ml := psess.newMachoLoad(ctxt.Arch, LC_VERSION_MIN_MACOSX, 2)
		ml.data[0] = 10<<16 | 7<<8 | 0<<0
		ml.data[1] = 10<<16 | 7<<8 | 0<<0
	}

	a := psess.machowrite(ctxt.Arch, ctxt.Out, ctxt.LinkMode)
	if int32(a) > psess.HEADR {
		psess.
			Exitf("HEADR too small: %d > %d", a, psess.HEADR)
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

func (psess *PackageSession) addsym(ctxt *Link, s *sym.Symbol, name string, type_ SymbolType, addr int64, gotype *sym.Symbol) {
	if s == nil {
		return
	}

	switch type_ {
	default:
		return

	case DataSym, BSSSym, TextSym:
		break
	}

	if psess.sortsym != nil {
		psess.
			sortsym[psess.nsortsym] = s
		psess.
			nkind[symkind(s)]++
	}
	psess.
		nsortsym++
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

func (psess *PackageSession) machogenasmsym(ctxt *Link) {
	psess.
		genasmsym(ctxt, psess.addsym)
	for _, s := range ctxt.Syms.Allsym {
		if s.Type == sym.SDYNIMPORT || s.Type == sym.SHOSTOBJ {
			if s.Attr.Reachable() {
				psess.
					addsym(ctxt, s, "", DataSym, 0, nil)
			}
		}
	}
}

func (psess *PackageSession) machosymorder(ctxt *Link) {

	for i := range psess.dynexp {
		psess.
			dynexp[i].Attr |= sym.AttrReachable
	}
	psess.
		machogenasmsym(ctxt)
	psess.
		sortsym = make([]*sym.Symbol, psess.nsortsym)
	psess.
		nsortsym = 0
	psess.
		machogenasmsym(ctxt)
	sort.Sort(machoscmp(psess.sortsym[:psess.nsortsym]))
	for i := 0; i < psess.nsortsym; i++ {
		psess.
			sortsym[i].Dynid = int32(i)
	}
}

// machoShouldExport reports whether a symbol needs to be exported.
//
// When dynamically linking, all non-local variables and plugin-exported
// symbols need to be exported.
func (psess *PackageSession) machoShouldExport(ctxt *Link, s *sym.Symbol) bool {
	if !ctxt.DynlinkingGo() || s.Attr.Local() {
		return false
	}
	if ctxt.BuildMode == BuildModePlugin && strings.HasPrefix(s.Extname, objabi.PathToPrefix(*psess.flagPluginPath)) {
		return true
	}
	if strings.HasPrefix(s.Name, "go.itab.") {
		return true
	}
	if strings.HasPrefix(s.Name, "type.") && !strings.HasPrefix(s.Name, "type..") {

		return true
	}
	if strings.HasPrefix(s.Name, "go.link.pkghash") {
		return true
	}
	return s.Type >= sym.SELFSECT
}

func (psess *PackageSession) machosymtab(ctxt *Link) {
	symtab := ctxt.Syms.Lookup(".machosymtab", 0)
	symstr := ctxt.Syms.Lookup(".machosymstr", 0)

	for i := 0; i < psess.nsortsym; i++ {
		s := psess.sortsym[i]
		symtab.AddUint32(ctxt.Arch, uint32(symstr.Size))

		export := psess.machoShouldExport(ctxt, s)

		cexport := !strings.Contains(s.Extname, ".") && (ctxt.BuildMode != BuildModePlugin || onlycsymbol(s))
		if cexport || export {
			symstr.AddUint8('_')
		}
		psess.
			Addstring(symstr, strings.Replace(s.Extname, "·", ".", -1))

		if s.Type == sym.SDYNIMPORT || s.Type == sym.SHOSTOBJ {
			symtab.AddUint8(0x01)
			symtab.AddUint8(0)
			symtab.AddUint16(ctxt.Arch, 0)
			symtab.AddUintXX(ctxt.Arch, 0, ctxt.Arch.PtrSize)
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
				psess.
					Errorf(s, "missing section for symbol")
				symtab.AddUint8(0)
			} else {
				symtab.AddUint8(uint8(o.Sect.Extnum))
			}
			symtab.AddUint16(ctxt.Arch, 0)
			symtab.AddUintXX(ctxt.Arch, uint64(psess.Symaddr(s)), ctxt.Arch.PtrSize)
		}
	}
}

func (psess *PackageSession) machodysymtab(ctxt *Link) {
	ml := psess.newMachoLoad(ctxt.Arch, LC_DYSYMTAB, 18)

	n := 0
	ml.data[0] = uint32(n)
	ml.data[1] = uint32(psess.nkind[SymKindLocal])
	n += psess.nkind[SymKindLocal]

	ml.data[2] = uint32(n)
	ml.data[3] = uint32(psess.nkind[SymKindExtdef])
	n += psess.nkind[SymKindExtdef]

	ml.data[4] = uint32(n)
	ml.data[5] = uint32(psess.nkind[SymKindUndef])

	ml.data[6] = 0
	ml.data[7] = 0
	ml.data[8] = 0
	ml.data[9] = 0
	ml.data[10] = 0
	ml.data[11] = 0

	s1 := ctxt.Syms.Lookup(".machosymtab", 0)

	s2 := ctxt.Syms.Lookup(".linkedit.plt", 0)
	s3 := ctxt.Syms.Lookup(".linkedit.got", 0)
	ml.data[12] = uint32(psess.linkoff + s1.Size)
	ml.data[13] = uint32((s2.Size + s3.Size) / 4)

	ml.data[14] = 0
	ml.data[15] = 0
	ml.data[16] = 0
	ml.data[17] = 0
}

func (psess *PackageSession) Domacholink(ctxt *Link) int64 {
	psess.
		machosymtab(ctxt)

	s1 := ctxt.Syms.Lookup(".machosymtab", 0)

	s2 := ctxt.Syms.Lookup(".linkedit.plt", 0)
	s3 := ctxt.Syms.Lookup(".linkedit.got", 0)
	s4 := ctxt.Syms.Lookup(".machosymstr", 0)

	for s4.Size%16 != 0 {
		s4.AddUint8(0)
	}

	size := int(s1.Size + s2.Size + s3.Size + s4.Size)

	if size > 0 {
		psess.
			linkoff = Rnd(int64(uint64(psess.HEADR)+psess.Segtext.Length), int64(*psess.FlagRound)) + Rnd(int64(psess.Segdata.Filelen), int64(*psess.FlagRound)) + Rnd(int64(psess.Segdwarf.Filelen), int64(*psess.FlagRound))
		ctxt.Out.SeekSet(psess, psess.linkoff)

		ctxt.Out.Write(s1.P[:s1.Size])
		ctxt.Out.Write(s2.P[:s2.Size])
		ctxt.Out.Write(s3.P[:s3.Size])
		ctxt.Out.Write(s4.P[:s4.Size])
	}

	return Rnd(int64(size), int64(*psess.FlagRound))
}

func (psess *PackageSession) machorelocsect(ctxt *Link, sect *sym.Section, syms []*sym.Symbol) {

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
				psess.
					Errorf(s, "missing xsym in relocation")
				continue
			}
			if !r.Xsym.Attr.Reachable() {
				psess.
					Errorf(s, "unreachable reloc %d (%s) target %v", r.Type, psess.sym.RelocName(ctxt.Arch, r.Type), r.Xsym.Name)
			}
			if !psess.thearch.Machoreloc1(ctxt.Arch, ctxt.Out, s, r, int64(uint64(s.Value+int64(r.Off))-sect.Vaddr)) {
				psess.
					Errorf(s, "unsupported obj reloc %d (%s)/%d to %s", r.Type, psess.sym.RelocName(ctxt.Arch, r.Type), r.Siz, r.Sym.Name)
			}
		}
	}

	sect.Rellen = uint64(ctxt.Out.Offset()) - sect.Reloff
}

func (psess *PackageSession) Machoemitreloc(ctxt *Link) {
	for ctxt.Out.Offset()&7 != 0 {
		ctxt.Out.Write8(0)
	}
	psess.
		machorelocsect(ctxt, psess.Segtext.Sections[0], ctxt.Textp)
	for _, sect := range psess.Segtext.Sections[1:] {
		psess.
			machorelocsect(ctxt, sect, psess.datap)
	}
	for _, sect := range psess.Segdata.Sections {
		psess.
			machorelocsect(ctxt, sect, psess.datap)
	}
	for _, sect := range psess.Segdwarf.Sections {
		psess.
			machorelocsect(ctxt, sect, psess.dwarfp)
	}
}

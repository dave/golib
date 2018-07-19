package ld

import (
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"io"
	"path/filepath"
	"sort"
	"strings"
)

/*
 * Note header.  The ".note" section contains an array of notes.  Each
 * begins with this header, aligned to a word boundary.  Immediately
 * following the note header is n_namesz bytes of name, padded to the
 * next word boundary.  Then comes n_descsz bytes of descriptor, again
 * padded to a word boundary.  The values of n_namesz and n_descsz do
 * not include the padding.
 */
type elfNote struct {
	nNamesz uint32
	nDescsz uint32
	nType   uint32
}

const (
	EI_MAG0              = 0
	EI_MAG1              = 1
	EI_MAG2              = 2
	EI_MAG3              = 3
	EI_CLASS             = 4
	EI_DATA              = 5
	EI_VERSION           = 6
	EI_OSABI             = 7
	EI_ABIVERSION        = 8
	OLD_EI_BRAND         = 8
	EI_PAD               = 9
	EI_NIDENT            = 16
	ELFMAG0              = 0x7f
	ELFMAG1              = 'E'
	ELFMAG2              = 'L'
	ELFMAG3              = 'F'
	SELFMAG              = 4
	EV_NONE              = 0
	EV_CURRENT           = 1
	ELFCLASSNONE         = 0
	ELFCLASS32           = 1
	ELFCLASS64           = 2
	ELFDATANONE          = 0
	ELFDATA2LSB          = 1
	ELFDATA2MSB          = 2
	ELFOSABI_NONE        = 0
	ELFOSABI_HPUX        = 1
	ELFOSABI_NETBSD      = 2
	ELFOSABI_LINUX       = 3
	ELFOSABI_HURD        = 4
	ELFOSABI_86OPEN      = 5
	ELFOSABI_SOLARIS     = 6
	ELFOSABI_AIX         = 7
	ELFOSABI_IRIX        = 8
	ELFOSABI_FREEBSD     = 9
	ELFOSABI_TRU64       = 10
	ELFOSABI_MODESTO     = 11
	ELFOSABI_OPENBSD     = 12
	ELFOSABI_OPENVMS     = 13
	ELFOSABI_NSK         = 14
	ELFOSABI_ARM         = 97
	ELFOSABI_STANDALONE  = 255
	ELFOSABI_SYSV        = ELFOSABI_NONE
	ELFOSABI_MONTEREY    = ELFOSABI_AIX
	ET_NONE              = 0
	ET_REL               = 1
	ET_EXEC              = 2
	ET_DYN               = 3
	ET_CORE              = 4
	ET_LOOS              = 0xfe00
	ET_HIOS              = 0xfeff
	ET_LOPROC            = 0xff00
	ET_HIPROC            = 0xffff
	EM_NONE              = 0
	EM_M32               = 1
	EM_SPARC             = 2
	EM_386               = 3
	EM_68K               = 4
	EM_88K               = 5
	EM_860               = 7
	EM_MIPS              = 8
	EM_S370              = 9
	EM_MIPS_RS3_LE       = 10
	EM_PARISC            = 15
	EM_VPP500            = 17
	EM_SPARC32PLUS       = 18
	EM_960               = 19
	EM_PPC               = 20
	EM_PPC64             = 21
	EM_S390              = 22
	EM_V800              = 36
	EM_FR20              = 37
	EM_RH32              = 38
	EM_RCE               = 39
	EM_ARM               = 40
	EM_SH                = 42
	EM_SPARCV9           = 43
	EM_TRICORE           = 44
	EM_ARC               = 45
	EM_H8_300            = 46
	EM_H8_300H           = 47
	EM_H8S               = 48
	EM_H8_500            = 49
	EM_IA_64             = 50
	EM_MIPS_X            = 51
	EM_COLDFIRE          = 52
	EM_68HC12            = 53
	EM_MMA               = 54
	EM_PCP               = 55
	EM_NCPU              = 56
	EM_NDR1              = 57
	EM_STARCORE          = 58
	EM_ME16              = 59
	EM_ST100             = 60
	EM_TINYJ             = 61
	EM_X86_64            = 62
	EM_AARCH64           = 183
	EM_486               = 6
	EM_MIPS_RS4_BE       = 10
	EM_ALPHA_STD         = 41
	EM_ALPHA             = 0x9026
	SHN_UNDEF            = 0
	SHN_LORESERVE        = 0xff00
	SHN_LOPROC           = 0xff00
	SHN_HIPROC           = 0xff1f
	SHN_LOOS             = 0xff20
	SHN_HIOS             = 0xff3f
	SHN_ABS              = 0xfff1
	SHN_COMMON           = 0xfff2
	SHN_XINDEX           = 0xffff
	SHN_HIRESERVE        = 0xffff
	SHT_NULL             = 0
	SHT_PROGBITS         = 1
	SHT_SYMTAB           = 2
	SHT_STRTAB           = 3
	SHT_RELA             = 4
	SHT_HASH             = 5
	SHT_DYNAMIC          = 6
	SHT_NOTE             = 7
	SHT_NOBITS           = 8
	SHT_REL              = 9
	SHT_SHLIB            = 10
	SHT_DYNSYM           = 11
	SHT_INIT_ARRAY       = 14
	SHT_FINI_ARRAY       = 15
	SHT_PREINIT_ARRAY    = 16
	SHT_GROUP            = 17
	SHT_SYMTAB_SHNDX     = 18
	SHT_LOOS             = 0x60000000
	SHT_HIOS             = 0x6fffffff
	SHT_GNU_VERDEF       = 0x6ffffffd
	SHT_GNU_VERNEED      = 0x6ffffffe
	SHT_GNU_VERSYM       = 0x6fffffff
	SHT_LOPROC           = 0x70000000
	SHT_ARM_ATTRIBUTES   = 0x70000003
	SHT_HIPROC           = 0x7fffffff
	SHT_LOUSER           = 0x80000000
	SHT_HIUSER           = 0xffffffff
	SHF_WRITE            = 0x1
	SHF_ALLOC            = 0x2
	SHF_EXECINSTR        = 0x4
	SHF_MERGE            = 0x10
	SHF_STRINGS          = 0x20
	SHF_INFO_LINK        = 0x40
	SHF_LINK_ORDER       = 0x80
	SHF_OS_NONCONFORMING = 0x100
	SHF_GROUP            = 0x200
	SHF_TLS              = 0x400
	SHF_MASKOS           = 0x0ff00000
	SHF_MASKPROC         = 0xf0000000
	PT_NULL              = 0
	PT_LOAD              = 1
	PT_DYNAMIC           = 2
	PT_INTERP            = 3
	PT_NOTE              = 4
	PT_SHLIB             = 5
	PT_PHDR              = 6
	PT_TLS               = 7
	PT_LOOS              = 0x60000000
	PT_HIOS              = 0x6fffffff
	PT_LOPROC            = 0x70000000
	PT_HIPROC            = 0x7fffffff
	PT_GNU_STACK         = 0x6474e551
	PT_GNU_RELRO         = 0x6474e552
	PT_PAX_FLAGS         = 0x65041580
	PT_SUNWSTACK         = 0x6ffffffb
	PF_X                 = 0x1
	PF_W                 = 0x2
	PF_R                 = 0x4
	PF_MASKOS            = 0x0ff00000
	PF_MASKPROC          = 0xf0000000
	DT_NULL              = 0
	DT_NEEDED            = 1
	DT_PLTRELSZ          = 2
	DT_PLTGOT            = 3
	DT_HASH              = 4
	DT_STRTAB            = 5
	DT_SYMTAB            = 6
	DT_RELA              = 7
	DT_RELASZ            = 8
	DT_RELAENT           = 9
	DT_STRSZ             = 10
	DT_SYMENT            = 11
	DT_INIT              = 12
	DT_FINI              = 13
	DT_SONAME            = 14
	DT_RPATH             = 15
	DT_SYMBOLIC          = 16
	DT_REL               = 17
	DT_RELSZ             = 18
	DT_RELENT            = 19
	DT_PLTREL            = 20
	DT_DEBUG             = 21
	DT_TEXTREL           = 22
	DT_JMPREL            = 23
	DT_BIND_NOW          = 24
	DT_INIT_ARRAY        = 25
	DT_FINI_ARRAY        = 26
	DT_INIT_ARRAYSZ      = 27
	DT_FINI_ARRAYSZ      = 28
	DT_RUNPATH           = 29
	DT_FLAGS             = 30
	DT_ENCODING          = 32
	DT_PREINIT_ARRAY     = 32
	DT_PREINIT_ARRAYSZ   = 33
	DT_LOOS              = 0x6000000d
	DT_HIOS              = 0x6ffff000
	DT_LOPROC            = 0x70000000
	DT_HIPROC            = 0x7fffffff
	DT_VERNEED           = 0x6ffffffe
	DT_VERNEEDNUM        = 0x6fffffff
	DT_VERSYM            = 0x6ffffff0
	DT_PPC64_GLINK       = DT_LOPROC + 0
	DT_PPC64_OPT         = DT_LOPROC + 3
	DF_ORIGIN            = 0x0001
	DF_SYMBOLIC          = 0x0002
	DF_TEXTREL           = 0x0004
	DF_BIND_NOW          = 0x0008
	DF_STATIC_TLS        = 0x0010
	NT_PRSTATUS          = 1
	NT_FPREGSET          = 2
	NT_PRPSINFO          = 3
	STB_LOCAL            = 0
	STB_GLOBAL           = 1
	STB_WEAK             = 2
	STB_LOOS             = 10
	STB_HIOS             = 12
	STB_LOPROC           = 13
	STB_HIPROC           = 15
	STT_NOTYPE           = 0
	STT_OBJECT           = 1
	STT_FUNC             = 2
	STT_SECTION          = 3
	STT_FILE             = 4
	STT_COMMON           = 5
	STT_TLS              = 6
	STT_LOOS             = 10
	STT_HIOS             = 12
	STT_LOPROC           = 13
	STT_HIPROC           = 15
	STV_DEFAULT          = 0x0
	STV_INTERNAL         = 0x1
	STV_HIDDEN           = 0x2
	STV_PROTECTED        = 0x3
	STN_UNDEF            = 0
)

/*
 * Relocation types.
 */
const (
	ARM_MAGIC_TRAMP_NUMBER = 0x5c000003
)

/*
 * ELF header.
 */
type ElfEhdr struct {
	ident     [EI_NIDENT]uint8
	type_     uint16
	machine   uint16
	version   uint32
	entry     uint64
	phoff     uint64
	shoff     uint64
	flags     uint32
	ehsize    uint16
	phentsize uint16
	phnum     uint16
	shentsize uint16
	shnum     uint16
	shstrndx  uint16
}

/*
 * Section header.
 */
type ElfShdr struct {
	name      uint32
	type_     uint32
	flags     uint64
	addr      uint64
	off       uint64
	size      uint64
	link      uint32
	info      uint32
	addralign uint64
	entsize   uint64
	shnum     int
}

/*
 * Program header.
 */
type ElfPhdr struct {
	type_  uint32
	flags  uint32
	off    uint64
	vaddr  uint64
	paddr  uint64
	filesz uint64
	memsz  uint64
	align  uint64
}

/*
 * Go linker interface
 */
const (
	ELF64HDRSIZE  = 64
	ELF64PHDRSIZE = 56
	ELF64SHDRSIZE = 64
	ELF64RELSIZE  = 16
	ELF64RELASIZE = 24
	ELF64SYMSIZE  = 24
	ELF32HDRSIZE  = 52
	ELF32PHDRSIZE = 32
	ELF32SHDRSIZE = 40
	ELF32SYMSIZE  = 16
	ELF32RELSIZE  = 8
)

/*
 * Total amount of space to reserve at the start of the file
 * for Header, PHeaders, SHeaders, and interp.
 * May waste some.
 * On FreeBSD, cannot be larger than a page.
 */
const (
	ELFRESERVE = 4096
)

/*
 * We use the 64-bit data structures on both 32- and 64-bit machines
 * in order to write the code just once.  The 64-bit data structure is
 * written in the 32-bit format on the 32-bit machines.
 */
const (
	NSECT = 400
)

// Either ".rel" or ".rela" depending on which type of relocation the
// target platform uses.

type Elfstring struct {
	s   string
	off int
}

/*
 Initialize the global variable that describes the ELF header. It will be updated as
 we write section and prog headers.
*/
func (psess *PackageSession) Elfinit(ctxt *Link) {
	ctxt.IsELF = true

	if ctxt.Arch.InFamily(sys.AMD64, sys.ARM64, sys.MIPS64, sys.PPC64, sys.S390X) {
		psess.
			elfRelType = ".rela"
	} else {
		psess.
			elfRelType = ".rel"
	}

	switch ctxt.Arch.Family {

	case sys.PPC64, sys.S390X:
		if ctxt.Arch.ByteOrder == binary.BigEndian {
			psess.
				ehdr.flags = 1
		} else {
			psess.
				ehdr.flags = 2
		}
		fallthrough
	case sys.AMD64, sys.ARM64, sys.MIPS64:
		if ctxt.Arch.Family == sys.MIPS64 {
			psess.
				ehdr.flags = 0x20000004
		}
		psess.
			elf64 = true
		psess.
			ehdr.phoff = ELF64HDRSIZE
		psess.
			ehdr.shoff = ELF64HDRSIZE
		psess.
			ehdr.ehsize = ELF64HDRSIZE
		psess.
			ehdr.phentsize = ELF64PHDRSIZE
		psess.
			ehdr.shentsize = ELF64SHDRSIZE

	case sys.ARM, sys.MIPS:
		if ctxt.Arch.Family == sys.ARM {

			if ctxt.HeadType == objabi.Hlinux || ctxt.HeadType == objabi.Hfreebsd || ctxt.HeadType == objabi.Hnetbsd {
				psess.
					ehdr.flags = 0x5000002
			}
		} else if ctxt.Arch.Family == sys.MIPS {
			psess.
				ehdr.flags = 0x50001004
		}
		fallthrough
	default:
		psess.
			ehdr.phoff = ELF32HDRSIZE
		psess.
			ehdr.shoff = ELF32HDRSIZE
		psess.
			ehdr.ehsize = ELF32HDRSIZE
		psess.
			ehdr.phentsize = ELF32PHDRSIZE
		psess.
			ehdr.shentsize = ELF32SHDRSIZE
	}
}

// Make sure PT_LOAD is aligned properly and
// that there is no gap,
// correct ELF loaders will do this implicitly,
// but buggy ELF loaders like the one in some
// versions of QEMU and UPX won't.
func fixElfPhdr(e *ElfPhdr) {
	frag := int(e.vaddr & (e.align - 1))

	e.off -= uint64(frag)
	e.vaddr -= uint64(frag)
	e.paddr -= uint64(frag)
	e.filesz += uint64(frag)
	e.memsz += uint64(frag)
}

func elf64phdr(out *OutBuf, e *ElfPhdr) {
	if e.type_ == PT_LOAD {
		fixElfPhdr(e)
	}

	out.Write32(e.type_)
	out.Write32(e.flags)
	out.Write64(e.off)
	out.Write64(e.vaddr)
	out.Write64(e.paddr)
	out.Write64(e.filesz)
	out.Write64(e.memsz)
	out.Write64(e.align)
}

func elf32phdr(out *OutBuf, e *ElfPhdr) {
	if e.type_ == PT_LOAD {
		fixElfPhdr(e)
	}

	out.Write32(e.type_)
	out.Write32(uint32(e.off))
	out.Write32(uint32(e.vaddr))
	out.Write32(uint32(e.paddr))
	out.Write32(uint32(e.filesz))
	out.Write32(uint32(e.memsz))
	out.Write32(e.flags)
	out.Write32(uint32(e.align))
}

func elf64shdr(out *OutBuf, e *ElfShdr) {
	out.Write32(e.name)
	out.Write32(e.type_)
	out.Write64(e.flags)
	out.Write64(e.addr)
	out.Write64(e.off)
	out.Write64(e.size)
	out.Write32(e.link)
	out.Write32(e.info)
	out.Write64(e.addralign)
	out.Write64(e.entsize)
}

func elf32shdr(out *OutBuf, e *ElfShdr) {
	out.Write32(e.name)
	out.Write32(e.type_)
	out.Write32(uint32(e.flags))
	out.Write32(uint32(e.addr))
	out.Write32(uint32(e.off))
	out.Write32(uint32(e.size))
	out.Write32(e.link)
	out.Write32(e.info)
	out.Write32(uint32(e.addralign))
	out.Write32(uint32(e.entsize))
}

func (psess *PackageSession) elfwriteshdrs(out *OutBuf) uint32 {
	if psess.elf64 {
		for i := 0; i < int(psess.ehdr.shnum); i++ {
			elf64shdr(out, psess.shdr[i])
		}
		return uint32(psess.ehdr.shnum) * ELF64SHDRSIZE
	}

	for i := 0; i < int(psess.ehdr.shnum); i++ {
		elf32shdr(out, psess.shdr[i])
	}
	return uint32(psess.ehdr.shnum) * ELF32SHDRSIZE
}

func (psess *PackageSession) elfsetstring(s *sym.Symbol, str string, off int) {
	if psess.nelfstr >= len(psess.elfstr) {
		psess.
			Errorf(s, "too many elf strings")
		psess.
			errorexit()
	}
	psess.
		elfstr[psess.nelfstr].s = str
	psess.
		elfstr[psess.nelfstr].off = off
	psess.
		nelfstr++
}

func (psess *PackageSession) elfwritephdrs(out *OutBuf) uint32 {
	if psess.elf64 {
		for i := 0; i < int(psess.ehdr.phnum); i++ {
			elf64phdr(out, psess.phdr[i])
		}
		return uint32(psess.ehdr.phnum) * ELF64PHDRSIZE
	}

	for i := 0; i < int(psess.ehdr.phnum); i++ {
		elf32phdr(out, psess.phdr[i])
	}
	return uint32(psess.ehdr.phnum) * ELF32PHDRSIZE
}

func (psess *PackageSession) newElfPhdr() *ElfPhdr {
	e := new(ElfPhdr)
	if psess.ehdr.phnum >= NSECT {
		psess.
			Errorf(nil, "too many phdrs")
	} else {
		psess.
			phdr[psess.ehdr.phnum] = e
		psess.
			ehdr.phnum++
	}
	if psess.elf64 {
		psess.
			ehdr.shoff += ELF64PHDRSIZE
	} else {
		psess.
			ehdr.shoff += ELF32PHDRSIZE
	}
	return e
}

func (psess *PackageSession) newElfShdr(name int64) *ElfShdr {
	e := new(ElfShdr)
	e.name = uint32(name)
	e.shnum = int(psess.ehdr.shnum)
	if psess.ehdr.shnum >= NSECT {
		psess.
			Errorf(nil, "too many shdrs")
	} else {
		psess.
			shdr[psess.ehdr.shnum] = e
		psess.
			ehdr.shnum++
	}

	return e
}

func (psess *PackageSession) getElfEhdr() *ElfEhdr {
	return &psess.ehdr
}

func (psess *PackageSession) elf64writehdr(out *OutBuf) uint32 {
	out.Write(psess.ehdr.ident[:])
	out.Write16(psess.ehdr.type_)
	out.Write16(psess.ehdr.machine)
	out.Write32(psess.ehdr.version)
	out.Write64(psess.ehdr.entry)
	out.Write64(psess.ehdr.phoff)
	out.Write64(psess.ehdr.shoff)
	out.Write32(psess.ehdr.flags)
	out.Write16(psess.ehdr.ehsize)
	out.Write16(psess.ehdr.phentsize)
	out.Write16(psess.ehdr.phnum)
	out.Write16(psess.ehdr.shentsize)
	out.Write16(psess.ehdr.shnum)
	out.Write16(psess.ehdr.shstrndx)
	return ELF64HDRSIZE
}

func (psess *PackageSession) elf32writehdr(out *OutBuf) uint32 {
	out.Write(psess.ehdr.ident[:])
	out.Write16(psess.ehdr.type_)
	out.Write16(psess.ehdr.machine)
	out.Write32(psess.ehdr.version)
	out.Write32(uint32(psess.ehdr.entry))
	out.Write32(uint32(psess.ehdr.phoff))
	out.Write32(uint32(psess.ehdr.shoff))
	out.Write32(psess.ehdr.flags)
	out.Write16(psess.ehdr.ehsize)
	out.Write16(psess.ehdr.phentsize)
	out.Write16(psess.ehdr.phnum)
	out.Write16(psess.ehdr.shentsize)
	out.Write16(psess.ehdr.shnum)
	out.Write16(psess.ehdr.shstrndx)
	return ELF32HDRSIZE
}

func (psess *PackageSession) elfwritehdr(out *OutBuf) uint32 {
	if psess.elf64 {
		return psess.elf64writehdr(out)
	}
	return psess.elf32writehdr(out)
}

/* Taken directly from the definition document for ELF64 */
func elfhash(name string) uint32 {
	var h uint32
	for i := 0; i < len(name); i++ {
		h = (h << 4) + uint32(name[i])
		if g := h & 0xf0000000; g != 0 {
			h ^= g >> 24
		}
		h &= 0x0fffffff
	}
	return h
}

func (psess *PackageSession) Elfwritedynent(ctxt *Link, s *sym.Symbol, tag int, val uint64) {
	if psess.elf64 {
		s.AddUint64(ctxt.Arch, uint64(tag))
		s.AddUint64(ctxt.Arch, val)
	} else {
		s.AddUint32(ctxt.Arch, uint32(tag))
		s.AddUint32(ctxt.Arch, uint32(val))
	}
}

func (psess *PackageSession) elfwritedynentsym(ctxt *Link, s *sym.Symbol, tag int, t *sym.Symbol) {
	psess.
		Elfwritedynentsymplus(ctxt, s, tag, t, 0)
}

func (psess *PackageSession) Elfwritedynentsymplus(ctxt *Link, s *sym.Symbol, tag int, t *sym.Symbol, add int64) {
	if psess.elf64 {
		s.AddUint64(ctxt.Arch, uint64(tag))
	} else {
		s.AddUint32(ctxt.Arch, uint32(tag))
	}
	s.AddAddrPlus(ctxt.Arch, t, add)
}

func (psess *PackageSession) elfwritedynentsymsize(ctxt *Link, s *sym.Symbol, tag int, t *sym.Symbol) {
	if psess.elf64 {
		s.AddUint64(ctxt.Arch, uint64(tag))
	} else {
		s.AddUint32(ctxt.Arch, uint32(tag))
	}
	s.AddSize(ctxt.Arch, t)
}

func (psess *PackageSession) elfinterp(sh *ElfShdr, startva uint64, resoff uint64, p string) int {
	psess.
		interp = p
	n := len(psess.interp) + 1
	sh.addr = startva + resoff - uint64(n)
	sh.off = resoff - uint64(n)
	sh.size = uint64(n)

	return n
}

func (psess *PackageSession) elfwriteinterp(out *OutBuf) int {
	sh := psess.elfshname(".interp")
	out.SeekSet(psess, int64(sh.off))
	out.WriteString(psess.interp)
	out.Write8(0)
	return int(sh.size)
}

func elfnote(sh *ElfShdr, startva uint64, resoff uint64, sz int) int {
	n := 3*4 + uint64(sz) + resoff%4

	sh.type_ = SHT_NOTE
	sh.flags = SHF_ALLOC
	sh.addralign = 4
	sh.addr = startva + resoff - n
	sh.off = resoff - n
	sh.size = n - resoff%4

	return int(n)
}

func (psess *PackageSession) elfwritenotehdr(out *OutBuf, str string, namesz uint32, descsz uint32, tag uint32) *ElfShdr {
	sh := psess.elfshname(str)

	out.SeekSet(psess, int64(sh.off))

	out.Write32(namesz)
	out.Write32(descsz)
	out.Write32(tag)

	return sh
}

// NetBSD Signature (as per sys/exec_elf.h)
const (
	ELF_NOTE_NETBSD_NAMESZ  = 7
	ELF_NOTE_NETBSD_DESCSZ  = 4
	ELF_NOTE_NETBSD_TAG     = 1
	ELF_NOTE_NETBSD_VERSION = 599000000 /* NetBSD 5.99 */
)

func elfnetbsdsig(sh *ElfShdr, startva uint64, resoff uint64) int {
	n := int(Rnd(ELF_NOTE_NETBSD_NAMESZ, 4) + Rnd(ELF_NOTE_NETBSD_DESCSZ, 4))
	return elfnote(sh, startva, resoff, n)
}

func (psess *PackageSession) elfwritenetbsdsig(out *OutBuf) int {

	sh := psess.elfwritenotehdr(out, ".note.netbsd.ident", ELF_NOTE_NETBSD_NAMESZ, ELF_NOTE_NETBSD_DESCSZ, ELF_NOTE_NETBSD_TAG)

	if sh == nil {
		return 0
	}

	out.Write(psess.ELF_NOTE_NETBSD_NAME)
	out.Write8(0)
	out.Write32(ELF_NOTE_NETBSD_VERSION)

	return int(sh.size)
}

// OpenBSD Signature
const (
	ELF_NOTE_OPENBSD_NAMESZ  = 8
	ELF_NOTE_OPENBSD_DESCSZ  = 4
	ELF_NOTE_OPENBSD_TAG     = 1
	ELF_NOTE_OPENBSD_VERSION = 0
)

func elfopenbsdsig(sh *ElfShdr, startva uint64, resoff uint64) int {
	n := ELF_NOTE_OPENBSD_NAMESZ + ELF_NOTE_OPENBSD_DESCSZ
	return elfnote(sh, startva, resoff, n)
}

func (psess *PackageSession) elfwriteopenbsdsig(out *OutBuf) int {

	sh := psess.elfwritenotehdr(out, ".note.openbsd.ident", ELF_NOTE_OPENBSD_NAMESZ, ELF_NOTE_OPENBSD_DESCSZ, ELF_NOTE_OPENBSD_TAG)

	if sh == nil {
		return 0
	}

	out.Write(psess.ELF_NOTE_OPENBSD_NAME)

	out.Write32(ELF_NOTE_OPENBSD_VERSION)

	return int(sh.size)
}

func (psess *PackageSession) addbuildinfo(val string) {
	if !strings.HasPrefix(val, "0x") {
		psess.
			Exitf("-B argument must start with 0x: %s", val)
	}

	ov := val
	val = val[2:]

	const maxLen = 32
	if hex.DecodedLen(len(val)) > maxLen {
		psess.
			Exitf("-B option too long (max %d digits): %s", maxLen, ov)
	}

	b, err := hex.DecodeString(val)
	if err != nil {
		if err == hex.ErrLength {
			psess.
				Exitf("-B argument must have even number of digits: %s", ov)
		}
		if inv, ok := err.(hex.InvalidByteError); ok {
			psess.
				Exitf("-B argument contains invalid hex digit %c: %s", byte(inv), ov)
		}
		psess.
			Exitf("-B argument contains invalid hex: %s", ov)
	}
	psess.
		buildinfo = b
}

// Build info note
const (
	ELF_NOTE_BUILDINFO_NAMESZ = 4
	ELF_NOTE_BUILDINFO_TAG    = 3
)

func (psess *PackageSession) elfbuildinfo(sh *ElfShdr, startva uint64, resoff uint64) int {
	n := int(ELF_NOTE_BUILDINFO_NAMESZ + Rnd(int64(len(psess.buildinfo)), 4))
	return elfnote(sh, startva, resoff, n)
}

func (psess *PackageSession) elfgobuildid(sh *ElfShdr, startva uint64, resoff uint64) int {
	n := len(psess.ELF_NOTE_GO_NAME) + int(Rnd(int64(len(*psess.flagBuildid)), 4))
	return elfnote(sh, startva, resoff, n)
}

func (psess *PackageSession) elfwritebuildinfo(out *OutBuf) int {
	sh := psess.elfwritenotehdr(out, ".note.gnu.build-id", ELF_NOTE_BUILDINFO_NAMESZ, uint32(len(psess.buildinfo)), ELF_NOTE_BUILDINFO_TAG)
	if sh == nil {
		return 0
	}

	out.Write(psess.ELF_NOTE_BUILDINFO_NAME)
	out.Write(psess.buildinfo)
	var zero = make([]byte, 4)
	out.Write(zero[:int(Rnd(int64(len(psess.buildinfo)), 4)-int64(len(psess.buildinfo)))])

	return int(sh.size)
}

func (psess *PackageSession) elfwritegobuildid(out *OutBuf) int {
	sh := psess.elfwritenotehdr(out, ".note.go.buildid", uint32(len(psess.ELF_NOTE_GO_NAME)), uint32(len(*psess.flagBuildid)), ELF_NOTE_GOBUILDID_TAG)
	if sh == nil {
		return 0
	}

	out.Write(psess.ELF_NOTE_GO_NAME)
	out.Write([]byte(*psess.flagBuildid))
	var zero = make([]byte, 4)
	out.Write(zero[:int(Rnd(int64(len(*psess.flagBuildid)), 4)-int64(len(*psess.flagBuildid)))])

	return int(sh.size)
}

// Go specific notes
const (
	ELF_NOTE_GOPKGLIST_TAG = 1
	ELF_NOTE_GOABIHASH_TAG = 2
	ELF_NOTE_GODEPS_TAG    = 3
	ELF_NOTE_GOBUILDID_TAG = 4
)

type Elfaux struct {
	next *Elfaux
	num  int
	vers string
}

type Elflib struct {
	next *Elflib
	aux  *Elfaux
	file string
}

func addelflib(list **Elflib, file string, vers string) *Elfaux {
	var lib *Elflib

	for lib = *list; lib != nil; lib = lib.next {
		if lib.file == file {
			goto havelib
		}
	}
	lib = new(Elflib)
	lib.next = *list
	lib.file = file
	*list = lib

havelib:
	for aux := lib.aux; aux != nil; aux = aux.next {
		if aux.vers == vers {
			return aux
		}
	}
	aux := new(Elfaux)
	aux.next = lib.aux
	aux.vers = vers
	lib.aux = aux

	return aux
}

func (psess *PackageSession) elfdynhash(ctxt *Link) {
	if !ctxt.IsELF {
		return
	}

	nsym := psess.Nelfsym
	s := ctxt.Syms.Lookup(".hash", 0)
	s.Type = sym.SELFROSECT
	s.Attr |= sym.AttrReachable

	i := nsym
	nbucket := 1
	for i > 0 {
		nbucket++
		i >>= 1
	}

	var needlib *Elflib
	need := make([]*Elfaux, nsym)
	chain := make([]uint32, nsym)
	buckets := make([]uint32, nbucket)

	for _, sy := range ctxt.Syms.Allsym {
		if sy.Dynid <= 0 {
			continue
		}

		if sy.Dynimpvers != "" {
			need[sy.Dynid] = addelflib(&needlib, sy.Dynimplib, sy.Dynimpvers)
		}

		name := sy.Extname
		hc := elfhash(name)

		b := hc % uint32(nbucket)
		chain[sy.Dynid] = buckets[b]
		buckets[b] = uint32(sy.Dynid)
	}

	if ctxt.Arch.Family == sys.S390X {
		s.AddUint64(ctxt.Arch, uint64(nbucket))
		s.AddUint64(ctxt.Arch, uint64(nsym))
		for i := 0; i < nbucket; i++ {
			s.AddUint64(ctxt.Arch, uint64(buckets[i]))
		}
		for i := 0; i < nsym; i++ {
			s.AddUint64(ctxt.Arch, uint64(chain[i]))
		}
	} else {
		s.AddUint32(ctxt.Arch, uint32(nbucket))
		s.AddUint32(ctxt.Arch, uint32(nsym))
		for i := 0; i < nbucket; i++ {
			s.AddUint32(ctxt.Arch, buckets[i])
		}
		for i := 0; i < nsym; i++ {
			s.AddUint32(ctxt.Arch, chain[i])
		}
	}

	dynstr := ctxt.Syms.Lookup(".dynstr", 0)

	s = ctxt.Syms.Lookup(".gnu.version_r", 0)
	i = 2
	nfile := 0
	for l := needlib; l != nil; l = l.next {
		nfile++

		s.AddUint16(ctxt.Arch, 1)
		j := 0
		for x := l.aux; x != nil; x = x.next {
			j++
		}
		s.AddUint16(ctxt.Arch, uint16(j))
		s.AddUint32(ctxt.Arch, uint32(psess.Addstring(dynstr, l.file)))
		s.AddUint32(ctxt.Arch, 16)
		if l.next != nil {
			s.AddUint32(ctxt.Arch, 16+uint32(j)*16)
		} else {
			s.AddUint32(ctxt.Arch, 0)
		}

		for x := l.aux; x != nil; x = x.next {
			x.num = i
			i++

			s.AddUint32(ctxt.Arch, elfhash(x.vers))
			s.AddUint16(ctxt.Arch, 0)
			s.AddUint16(ctxt.Arch, uint16(x.num))
			s.AddUint32(ctxt.Arch, uint32(psess.Addstring(dynstr, x.vers)))
			if x.next != nil {
				s.AddUint32(ctxt.Arch, 16)
			} else {
				s.AddUint32(ctxt.Arch, 0)
			}
		}
	}

	s = ctxt.Syms.Lookup(".gnu.version", 0)

	for i := 0; i < nsym; i++ {
		if i == 0 {
			s.AddUint16(ctxt.Arch, 0)
		} else if need[i] == nil {
			s.AddUint16(ctxt.Arch, 1)
		} else {
			s.AddUint16(ctxt.Arch, uint16(need[i].num))
		}
	}

	s = ctxt.Syms.Lookup(".dynamic", 0)
	psess.
		elfverneed = nfile
	if psess.elfverneed != 0 {
		psess.
			elfwritedynentsym(ctxt, s, DT_VERNEED, ctxt.Syms.Lookup(".gnu.version_r", 0))
		psess.
			Elfwritedynent(ctxt, s, DT_VERNEEDNUM, uint64(nfile))
		psess.
			elfwritedynentsym(ctxt, s, DT_VERSYM, ctxt.Syms.Lookup(".gnu.version", 0))
	}

	sy := ctxt.Syms.Lookup(psess.elfRelType+".plt", 0)
	if sy.Size > 0 {
		if psess.elfRelType == ".rela" {
			psess.
				Elfwritedynent(ctxt, s, DT_PLTREL, DT_RELA)
		} else {
			psess.
				Elfwritedynent(ctxt, s, DT_PLTREL, DT_REL)
		}
		psess.
			elfwritedynentsymsize(ctxt, s, DT_PLTRELSZ, sy)
		psess.
			elfwritedynentsym(ctxt, s, DT_JMPREL, sy)
	}
	psess.
		Elfwritedynent(ctxt, s, DT_NULL, 0)
}

func (psess *PackageSession) elfphload(seg *sym.Segment) *ElfPhdr {
	ph := psess.newElfPhdr()
	ph.type_ = PT_LOAD
	if seg.Rwx&4 != 0 {
		ph.flags |= PF_R
	}
	if seg.Rwx&2 != 0 {
		ph.flags |= PF_W
	}
	if seg.Rwx&1 != 0 {
		ph.flags |= PF_X
	}
	ph.vaddr = seg.Vaddr
	ph.paddr = seg.Vaddr
	ph.memsz = seg.Length
	ph.off = seg.Fileoff
	ph.filesz = seg.Filelen
	ph.align = uint64(*psess.FlagRound)

	return ph
}

func (psess *PackageSession) elfphrelro(seg *sym.Segment) {
	ph := psess.newElfPhdr()
	ph.type_ = PT_GNU_RELRO
	ph.vaddr = seg.Vaddr
	ph.paddr = seg.Vaddr
	ph.memsz = seg.Length
	ph.off = seg.Fileoff
	ph.filesz = seg.Filelen
	ph.align = uint64(*psess.FlagRound)
}

func (psess *PackageSession) elfshname(name string) *ElfShdr {
	for i := 0; i < psess.nelfstr; i++ {
		if name != psess.elfstr[i].s {
			continue
		}
		off := psess.elfstr[i].off
		for i = 0; i < int(psess.ehdr.shnum); i++ {
			sh := psess.shdr[i]
			if sh.name == uint32(off) {
				return sh
			}
		}
		return psess.newElfShdr(int64(off))
	}
	psess.
		Exitf("cannot find elf name %s", name)
	return nil
}

// Create an ElfShdr for the section with name.
// Create a duplicate if one already exists with that name
func (psess *PackageSession) elfshnamedup(name string) *ElfShdr {
	for i := 0; i < psess.nelfstr; i++ {
		if name == psess.elfstr[i].s {
			off := psess.elfstr[i].off
			return psess.newElfShdr(int64(off))
		}
	}
	psess.
		Errorf(nil, "cannot find elf name %s", name)
	psess.
		errorexit()
	return nil
}

func (psess *PackageSession) elfshalloc(sect *sym.Section) *ElfShdr {
	sh := psess.elfshname(sect.Name)
	sect.Elfsect = sh
	return sh
}

func (psess *PackageSession) elfshbits(linkmode LinkMode, sect *sym.Section) *ElfShdr {
	var sh *ElfShdr

	if sect.Name == ".text" {
		if sect.Elfsect == nil {
			sect.Elfsect = psess.elfshnamedup(sect.Name)
		}
		sh = sect.Elfsect.(*ElfShdr)
	} else {
		sh = psess.elfshalloc(sect)
	}

	if sh.type_ == SHT_NOTE {
		if linkmode != LinkExternal {
			psess.
				Errorf(nil, "sh.type_ == SHT_NOTE in elfshbits when linking internally")
		}
		sh.addralign = uint64(sect.Align)
		sh.size = sect.Length
		sh.off = sect.Seg.Fileoff + sect.Vaddr - sect.Seg.Vaddr
		return sh
	}
	if sh.type_ > 0 {
		return sh
	}

	if sect.Vaddr < sect.Seg.Vaddr+sect.Seg.Filelen {
		sh.type_ = SHT_PROGBITS
	} else {
		sh.type_ = SHT_NOBITS
	}
	sh.flags = SHF_ALLOC
	if sect.Rwx&1 != 0 {
		sh.flags |= SHF_EXECINSTR
	}
	if sect.Rwx&2 != 0 {
		sh.flags |= SHF_WRITE
	}
	if sect.Name == ".tbss" {
		sh.flags |= SHF_TLS
		sh.type_ = SHT_NOBITS
	}
	if strings.HasPrefix(sect.Name, ".debug") || strings.HasPrefix(sect.Name, ".zdebug") {
		sh.flags = 0
	}

	if linkmode != LinkExternal {
		sh.addr = sect.Vaddr
	}
	sh.addralign = uint64(sect.Align)
	sh.size = sect.Length
	if sect.Name != ".tbss" {
		sh.off = sect.Seg.Fileoff + sect.Vaddr - sect.Seg.Vaddr
	}

	return sh
}

func (psess *PackageSession) elfshreloc(arch *sys.Arch, sect *sym.Section) *ElfShdr {

	if sect.Vaddr >= sect.Seg.Vaddr+sect.Seg.Filelen {
		return nil
	}
	if sect.Name == ".shstrtab" || sect.Name == ".tbss" {
		return nil
	}
	if sect.Elfsect.(*ElfShdr).type_ == SHT_NOTE {
		return nil
	}

	var typ int
	if psess.elfRelType == ".rela" {
		typ = SHT_RELA
	} else {
		typ = SHT_REL
	}

	sh := psess.elfshname(psess.elfRelType + sect.Name)

	if sect.Name == ".text" {
		if sh.info != 0 && sh.info != uint32(sect.Elfsect.(*ElfShdr).shnum) {
			sh = psess.elfshnamedup(psess.elfRelType + sect.Name)
		}
	}

	sh.type_ = uint32(typ)
	sh.entsize = uint64(arch.RegSize) * 2
	if typ == SHT_RELA {
		sh.entsize += uint64(arch.RegSize)
	}
	sh.link = uint32(psess.elfshname(".symtab").shnum)
	sh.info = uint32(sect.Elfsect.(*ElfShdr).shnum)
	sh.off = sect.Reloff
	sh.size = sect.Rellen
	sh.addralign = uint64(arch.RegSize)
	return sh
}

func (psess *PackageSession) elfrelocsect(ctxt *Link, sect *sym.Section, syms []*sym.Symbol) {

	if sect.Vaddr >= sect.Seg.Vaddr+sect.Seg.Filelen {
		return
	}
	if sect.Name == ".shstrtab" {
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
					Errorf(s, "missing xsym in relocation %#v %#v", r.Sym.Name, s)
				continue
			}
			if r.Xsym.ElfsymForReloc() == 0 {
				psess.
					Errorf(s, "reloc %d (%s) to non-elf symbol %s (outer=%s) %d (%s)", r.Type, psess.sym.RelocName(ctxt.Arch, r.Type), r.Sym.Name, r.Xsym.Name, r.Sym.Type, r.Sym.Type)
			}
			if !r.Xsym.Attr.Reachable() {
				psess.
					Errorf(s, "unreachable reloc %d (%s) target %v", r.Type, psess.sym.RelocName(ctxt.Arch, r.Type), r.Xsym.Name)
			}
			if !psess.thearch.Elfreloc1(ctxt, r, int64(uint64(s.Value+int64(r.Off))-sect.Vaddr)) {
				psess.
					Errorf(s, "unsupported obj reloc %d (%s)/%d to %s", r.Type, psess.sym.RelocName(ctxt.Arch, r.Type), r.Siz, r.Sym.Name)
			}
		}
	}

	sect.Rellen = uint64(ctxt.Out.Offset()) - sect.Reloff
}

func (psess *PackageSession) Elfemitreloc(ctxt *Link) {
	for ctxt.Out.Offset()&7 != 0 {
		ctxt.Out.Write8(0)
	}

	for _, sect := range psess.Segtext.Sections {
		if sect.Name == ".text" {
			psess.
				elfrelocsect(ctxt, sect, ctxt.Textp)
		} else {
			psess.
				elfrelocsect(ctxt, sect, psess.datap)
		}
	}

	for _, sect := range psess.Segrodata.Sections {
		psess.
			elfrelocsect(ctxt, sect, psess.datap)
	}
	for _, sect := range psess.Segrelrodata.Sections {
		psess.
			elfrelocsect(ctxt, sect, psess.datap)
	}
	for _, sect := range psess.Segdata.Sections {
		psess.
			elfrelocsect(ctxt, sect, psess.datap)
	}
	for _, sect := range psess.Segdwarf.Sections {
		psess.
			elfrelocsect(ctxt, sect, psess.dwarfp)
	}
}

func (psess *PackageSession) addgonote(ctxt *Link, sectionName string, tag uint32, desc []byte) {
	s := ctxt.Syms.Lookup(sectionName, 0)
	s.Attr |= sym.AttrReachable
	s.Type = sym.SELFROSECT

	s.AddUint32(ctxt.Arch, uint32(len(psess.ELF_NOTE_GO_NAME)))

	s.AddUint32(ctxt.Arch, uint32(len(desc)))

	s.AddUint32(ctxt.Arch, tag)

	s.P = append(s.P, psess.ELF_NOTE_GO_NAME...)
	for len(s.P)%4 != 0 {
		s.P = append(s.P, 0)
	}

	s.P = append(s.P, desc...)
	for len(s.P)%4 != 0 {
		s.P = append(s.P, 0)
	}
	s.Size = int64(len(s.P))
	s.Align = 4
}

func (ctxt *Link) doelf(psess *PackageSession) {
	if !ctxt.IsELF {
		return
	}

	shstrtab := ctxt.Syms.Lookup(".shstrtab", 0)

	shstrtab.Type = sym.SELFROSECT
	shstrtab.Attr |= sym.AttrReachable
	psess.
		Addstring(shstrtab, "")
	psess.
		Addstring(shstrtab, ".text")
	psess.
		Addstring(shstrtab, ".noptrdata")
	psess.
		Addstring(shstrtab, ".data")
	psess.
		Addstring(shstrtab, ".bss")
	psess.
		Addstring(shstrtab, ".noptrbss")

	if !*psess.FlagD || ctxt.LinkMode == LinkExternal {
		psess.
			Addstring(shstrtab, ".tbss")
	}
	if ctxt.HeadType == objabi.Hnetbsd {
		psess.
			Addstring(shstrtab, ".note.netbsd.ident")
	}
	if ctxt.HeadType == objabi.Hopenbsd {
		psess.
			Addstring(shstrtab, ".note.openbsd.ident")
	}
	if len(psess.buildinfo) > 0 {
		psess.
			Addstring(shstrtab, ".note.gnu.build-id")
	}
	if *psess.flagBuildid != "" {
		psess.
			Addstring(shstrtab, ".note.go.buildid")
	}
	psess.
		Addstring(shstrtab, ".elfdata")
	psess.
		Addstring(shstrtab, ".rodata")

	relro_prefix := ""
	if ctxt.UseRelro() {
		psess.
			Addstring(shstrtab, ".data.rel.ro")
		relro_prefix = ".data.rel.ro"
	}
	psess.
		Addstring(shstrtab, relro_prefix+".typelink")
	psess.
		Addstring(shstrtab, relro_prefix+".itablink")
	psess.
		Addstring(shstrtab, relro_prefix+".gosymtab")
	psess.
		Addstring(shstrtab, relro_prefix+".gopclntab")

	if ctxt.LinkMode == LinkExternal {
		*psess.FlagD = true
		psess.
			Addstring(shstrtab, psess.elfRelType+".text")
		psess.
			Addstring(shstrtab, psess.elfRelType+".rodata")
		psess.
			Addstring(shstrtab, psess.elfRelType+relro_prefix+".typelink")
		psess.
			Addstring(shstrtab, psess.elfRelType+relro_prefix+".itablink")
		psess.
			Addstring(shstrtab, psess.elfRelType+relro_prefix+".gosymtab")
		psess.
			Addstring(shstrtab, psess.elfRelType+relro_prefix+".gopclntab")
		psess.
			Addstring(shstrtab, psess.elfRelType+".noptrdata")
		psess.
			Addstring(shstrtab, psess.elfRelType+".data")
		if ctxt.UseRelro() {
			psess.
				Addstring(shstrtab, psess.elfRelType+".data.rel.ro")
		}
		psess.
			Addstring(shstrtab, ".note.GNU-stack")

		if ctxt.BuildMode == BuildModeShared {
			psess.
				Addstring(shstrtab, ".note.go.abihash")
			psess.
				Addstring(shstrtab, ".note.go.pkg-list")
			psess.
				Addstring(shstrtab, ".note.go.deps")
		}
	}

	hasinitarr := ctxt.linkShared

	switch ctxt.BuildMode {
	case BuildModeCArchive, BuildModeCShared, BuildModeShared, BuildModePlugin:
		hasinitarr = true
	}

	if hasinitarr {
		psess.
			Addstring(shstrtab, ".init_array")
		psess.
			Addstring(shstrtab, psess.elfRelType+".init_array")
	}

	if !*psess.FlagS {
		psess.
			Addstring(shstrtab, ".symtab")
		psess.
			Addstring(shstrtab, ".strtab")
		psess.
			dwarfaddshstrings(ctxt, shstrtab)
	}
	psess.
		Addstring(shstrtab, ".shstrtab")

	if !*psess.FlagD {
		psess.
			Addstring(shstrtab, ".interp")
		psess.
			Addstring(shstrtab, ".hash")
		psess.
			Addstring(shstrtab, ".got")
		if ctxt.Arch.Family == sys.PPC64 {
			psess.
				Addstring(shstrtab, ".glink")
		}
		psess.
			Addstring(shstrtab, ".got.plt")
		psess.
			Addstring(shstrtab, ".dynamic")
		psess.
			Addstring(shstrtab, ".dynsym")
		psess.
			Addstring(shstrtab, ".dynstr")
		psess.
			Addstring(shstrtab, psess.elfRelType)
		psess.
			Addstring(shstrtab, psess.elfRelType+".plt")
		psess.
			Addstring(shstrtab, ".plt")
		psess.
			Addstring(shstrtab, ".gnu.version")
		psess.
			Addstring(shstrtab, ".gnu.version_r")

		s := ctxt.Syms.Lookup(".dynsym", 0)

		s.Type = sym.SELFROSECT
		s.Attr |= sym.AttrReachable
		if psess.elf64 {
			s.Size += ELF64SYMSIZE
		} else {
			s.Size += ELF32SYMSIZE
		}

		s = ctxt.Syms.Lookup(".dynstr", 0)

		s.Type = sym.SELFROSECT
		s.Attr |= sym.AttrReachable
		if s.Size == 0 {
			psess.
				Addstring(s, "")
		}
		dynstr := s

		s = ctxt.Syms.Lookup(psess.elfRelType, 0)
		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFROSECT

		s = ctxt.Syms.Lookup(".got", 0)

		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFGOT

		if ctxt.Arch.Family == sys.PPC64 {
			s := ctxt.Syms.Lookup(".glink", 0)
			s.Attr |= sym.AttrReachable
			s.Type = sym.SELFRXSECT
		}

		s = ctxt.Syms.Lookup(".hash", 0)

		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFROSECT

		s = ctxt.Syms.Lookup(".got.plt", 0)
		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFSECT

		s = ctxt.Syms.Lookup(".plt", 0)

		s.Attr |= sym.AttrReachable
		if ctxt.Arch.Family == sys.PPC64 {

			s.Type = sym.SELFSECT
		} else {
			s.Type = sym.SELFRXSECT
		}
		psess.
			thearch.Elfsetupplt(ctxt)

		s = ctxt.Syms.Lookup(psess.elfRelType+".plt", 0)
		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFROSECT

		s = ctxt.Syms.Lookup(".gnu.version", 0)
		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFROSECT

		s = ctxt.Syms.Lookup(".gnu.version_r", 0)
		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFROSECT

		s = ctxt.Syms.Lookup(".dynamic", 0)

		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFSECT
		psess.
			elfwritedynentsym(ctxt, s, DT_HASH, ctxt.Syms.Lookup(".hash", 0))
		psess.
			elfwritedynentsym(ctxt, s, DT_SYMTAB, ctxt.Syms.Lookup(".dynsym", 0))
		if psess.elf64 {
			psess.
				Elfwritedynent(ctxt, s, DT_SYMENT, ELF64SYMSIZE)
		} else {
			psess.
				Elfwritedynent(ctxt, s, DT_SYMENT, ELF32SYMSIZE)
		}
		psess.
			elfwritedynentsym(ctxt, s, DT_STRTAB, ctxt.Syms.Lookup(".dynstr", 0))
		psess.
			elfwritedynentsymsize(ctxt, s, DT_STRSZ, ctxt.Syms.Lookup(".dynstr", 0))
		if psess.elfRelType == ".rela" {
			psess.
				elfwritedynentsym(ctxt, s, DT_RELA, ctxt.Syms.Lookup(".rela", 0))
			psess.
				elfwritedynentsymsize(ctxt, s, DT_RELASZ, ctxt.Syms.Lookup(".rela", 0))
			psess.
				Elfwritedynent(ctxt, s, DT_RELAENT, ELF64RELASIZE)
		} else {
			psess.
				elfwritedynentsym(ctxt, s, DT_REL, ctxt.Syms.Lookup(".rel", 0))
			psess.
				elfwritedynentsymsize(ctxt, s, DT_RELSZ, ctxt.Syms.Lookup(".rel", 0))
			psess.
				Elfwritedynent(ctxt, s, DT_RELENT, ELF32RELSIZE)
		}

		if psess.rpath.val != "" {
			psess.
				Elfwritedynent(ctxt, s, DT_RUNPATH, uint64(psess.Addstring(dynstr, psess.rpath.val)))
		}

		if ctxt.Arch.Family == sys.PPC64 {
			psess.
				elfwritedynentsym(ctxt, s, DT_PLTGOT, ctxt.Syms.Lookup(".plt", 0))
		} else if ctxt.Arch.Family == sys.S390X {
			psess.
				elfwritedynentsym(ctxt, s, DT_PLTGOT, ctxt.Syms.Lookup(".got", 0))
		} else {
			psess.
				elfwritedynentsym(ctxt, s, DT_PLTGOT, ctxt.Syms.Lookup(".got.plt", 0))
		}

		if ctxt.Arch.Family == sys.PPC64 {
			psess.
				Elfwritedynent(ctxt, s, DT_PPC64_OPT, 0)
		}
		psess.
			Elfwritedynent(ctxt, s, DT_DEBUG, 0)
	}

	if ctxt.BuildMode == BuildModeShared {

		s := ctxt.Syms.Lookup("go.link.abihashbytes", 0)
		s.Attr |= sym.AttrLocal
		s.Type = sym.SRODATA
		s.Attr |= sym.AttrSpecial
		s.Attr |= sym.AttrReachable
		s.Size = int64(sha1.Size)

		sort.Sort(byPkg(ctxt.Library))
		h := sha1.New()
		for _, l := range ctxt.Library {
			io.WriteString(h, l.Hash)
		}
		psess.
			addgonote(ctxt, ".note.go.abihash", ELF_NOTE_GOABIHASH_TAG, h.Sum([]byte{}))
		psess.
			addgonote(ctxt, ".note.go.pkg-list", ELF_NOTE_GOPKGLIST_TAG, psess.pkglistfornote)
		var deplist []string
		for _, shlib := range ctxt.Shlibs {
			deplist = append(deplist, filepath.Base(shlib.Path))
		}
		psess.
			addgonote(ctxt, ".note.go.deps", ELF_NOTE_GODEPS_TAG, []byte(strings.Join(deplist, "\n")))
	}

	if ctxt.LinkMode == LinkExternal && *psess.flagBuildid != "" {
		psess.
			addgonote(ctxt, ".note.go.buildid", ELF_NOTE_GOBUILDID_TAG, []byte(*psess.flagBuildid))
	}
}

// Do not write DT_NULL.  elfdynhash will finish it.
func (psess *PackageSession) shsym(sh *ElfShdr, s *sym.Symbol) {
	addr := psess.Symaddr(s)
	if sh.flags&SHF_ALLOC != 0 {
		sh.addr = uint64(addr)
	}
	sh.off = uint64(psess.datoff(s, addr))
	sh.size = uint64(s.Size)
}

func phsh(ph *ElfPhdr, sh *ElfShdr) {
	ph.vaddr = sh.addr
	ph.paddr = ph.vaddr
	ph.off = sh.off
	ph.filesz = sh.size
	ph.memsz = sh.size
	ph.align = sh.addralign
}

func (psess *PackageSession) Asmbelfsetup() {
	psess.
		elfshname("")

	for _, sect := range psess.Segtext.Sections {

		if sect.Name == ".text" {
			if sect.Elfsect == nil {
				sect.Elfsect = psess.elfshnamedup(sect.Name)
			}
		} else {
			psess.
				elfshalloc(sect)
		}
	}
	for _, sect := range psess.Segrodata.Sections {
		psess.
			elfshalloc(sect)
	}
	for _, sect := range psess.Segrelrodata.Sections {
		psess.
			elfshalloc(sect)
	}
	for _, sect := range psess.Segdata.Sections {
		psess.
			elfshalloc(sect)
	}
	for _, sect := range psess.Segdwarf.Sections {
		psess.
			elfshalloc(sect)
	}
}

func (psess *PackageSession) Asmbelf(ctxt *Link, symo int64) {
	eh := psess.getElfEhdr()
	switch ctxt.Arch.Family {
	default:
		psess.
			Exitf("unknown architecture in asmbelf: %v", ctxt.Arch.Family)
	case sys.MIPS, sys.MIPS64:
		eh.machine = EM_MIPS
	case sys.ARM:
		eh.machine = EM_ARM
	case sys.AMD64:
		eh.machine = EM_X86_64
	case sys.ARM64:
		eh.machine = EM_AARCH64
	case sys.I386:
		eh.machine = EM_386
	case sys.PPC64:
		eh.machine = EM_PPC64
	case sys.S390X:
		eh.machine = EM_S390
	}

	elfreserve := int64(ELFRESERVE)

	numtext := int64(0)
	for _, sect := range psess.Segtext.Sections {
		if sect.Name == ".text" {
			numtext++
		}
	}

	if numtext > 4 {
		elfreserve += elfreserve + numtext*64*2
	}

	startva := *psess.FlagTextAddr - int64(psess.HEADR)
	resoff := elfreserve

	var pph *ElfPhdr
	var pnote *ElfPhdr
	if ctxt.LinkMode == LinkExternal {

		eh.phoff = 0

		eh.phentsize = 0

		if ctxt.BuildMode == BuildModeShared {
			sh := psess.elfshname(".note.go.pkg-list")
			sh.type_ = SHT_NOTE
			sh = psess.elfshname(".note.go.abihash")
			sh.type_ = SHT_NOTE
			sh.flags = SHF_ALLOC
			sh = psess.elfshname(".note.go.deps")
			sh.type_ = SHT_NOTE
		}

		if *psess.flagBuildid != "" {
			sh := psess.elfshname(".note.go.buildid")
			sh.type_ = SHT_NOTE
			sh.flags = SHF_ALLOC
		}

		goto elfobj
	}

	pph = psess.newElfPhdr()

	pph.type_ = PT_PHDR
	pph.flags = PF_R
	pph.off = uint64(eh.ehsize)
	pph.vaddr = uint64(*psess.FlagTextAddr) - uint64(psess.HEADR) + pph.off
	pph.paddr = uint64(*psess.FlagTextAddr) - uint64(psess.HEADR) + pph.off
	pph.align = uint64(*psess.FlagRound)

	if ctxt.HeadType != objabi.Hnacl {
		o := int64(psess.Segtext.Vaddr - pph.vaddr)
		psess.
			Segtext.Vaddr -= uint64(o)
		psess.
			Segtext.Length += uint64(o)
		o = int64(psess.Segtext.Fileoff - pph.off)
		psess.
			Segtext.Fileoff -= uint64(o)
		psess.
			Segtext.Filelen += uint64(o)
	}

	if !*psess.FlagD {

		sh := psess.elfshname(".interp")

		sh.type_ = SHT_PROGBITS
		sh.flags = SHF_ALLOC
		sh.addralign = 1
		if psess.interpreter == "" {
			switch ctxt.HeadType {
			case objabi.Hlinux:
				psess.
					interpreter = psess.thearch.Linuxdynld

			case objabi.Hfreebsd:
				psess.
					interpreter = psess.thearch.Freebsddynld

			case objabi.Hnetbsd:
				psess.
					interpreter = psess.thearch.Netbsddynld

			case objabi.Hopenbsd:
				psess.
					interpreter = psess.thearch.Openbsddynld

			case objabi.Hdragonfly:
				psess.
					interpreter = psess.thearch.Dragonflydynld

			case objabi.Hsolaris:
				psess.
					interpreter = psess.thearch.Solarisdynld
			}
		}

		resoff -= int64(psess.elfinterp(sh, uint64(startva), uint64(resoff), psess.interpreter))

		ph := psess.newElfPhdr()
		ph.type_ = PT_INTERP
		ph.flags = PF_R
		phsh(ph, sh)
	}

	pnote = nil
	if ctxt.HeadType == objabi.Hnetbsd || ctxt.HeadType == objabi.Hopenbsd {
		var sh *ElfShdr
		switch ctxt.HeadType {
		case objabi.Hnetbsd:
			sh = psess.elfshname(".note.netbsd.ident")
			resoff -= int64(elfnetbsdsig(sh, uint64(startva), uint64(resoff)))

		case objabi.Hopenbsd:
			sh = psess.elfshname(".note.openbsd.ident")
			resoff -= int64(elfopenbsdsig(sh, uint64(startva), uint64(resoff)))
		}

		pnote = psess.newElfPhdr()
		pnote.type_ = PT_NOTE
		pnote.flags = PF_R
		phsh(pnote, sh)
	}

	if len(psess.buildinfo) > 0 {
		sh := psess.elfshname(".note.gnu.build-id")
		resoff -= int64(psess.elfbuildinfo(sh, uint64(startva), uint64(resoff)))

		if pnote == nil {
			pnote = psess.newElfPhdr()
			pnote.type_ = PT_NOTE
			pnote.flags = PF_R
		}

		phsh(pnote, sh)
	}

	if *psess.flagBuildid != "" {
		sh := psess.elfshname(".note.go.buildid")
		resoff -= int64(psess.elfgobuildid(sh, uint64(startva), uint64(resoff)))

		pnote := psess.newElfPhdr()
		pnote.type_ = PT_NOTE
		pnote.flags = PF_R
		phsh(pnote, sh)
	}
	psess.
		elfphload(&psess.Segtext)
	if len(psess.Segrodata.Sections) > 0 {
		psess.
			elfphload(&psess.Segrodata)
	}
	if len(psess.Segrelrodata.Sections) > 0 {
		psess.
			elfphload(&psess.Segrelrodata)
		psess.
			elfphrelro(&psess.Segrelrodata)
	}
	psess.
		elfphload(&psess.Segdata)

	if !*psess.FlagD {
		sh := psess.elfshname(".dynsym")
		sh.type_ = SHT_DYNSYM
		sh.flags = SHF_ALLOC
		if psess.elf64 {
			sh.entsize = ELF64SYMSIZE
		} else {
			sh.entsize = ELF32SYMSIZE
		}
		sh.addralign = uint64(ctxt.Arch.RegSize)
		sh.link = uint32(psess.elfshname(".dynstr").shnum)
		psess.
			shsym(sh, ctxt.Syms.Lookup(".dynsym", 0))

		sh = psess.elfshname(".dynstr")
		sh.type_ = SHT_STRTAB
		sh.flags = SHF_ALLOC
		sh.addralign = 1
		psess.
			shsym(sh, ctxt.Syms.Lookup(".dynstr", 0))

		if psess.elfverneed != 0 {
			sh := psess.elfshname(".gnu.version")
			sh.type_ = SHT_GNU_VERSYM
			sh.flags = SHF_ALLOC
			sh.addralign = 2
			sh.link = uint32(psess.elfshname(".dynsym").shnum)
			sh.entsize = 2
			psess.
				shsym(sh, ctxt.Syms.Lookup(".gnu.version", 0))

			sh = psess.elfshname(".gnu.version_r")
			sh.type_ = SHT_GNU_VERNEED
			sh.flags = SHF_ALLOC
			sh.addralign = uint64(ctxt.Arch.RegSize)
			sh.info = uint32(psess.elfverneed)
			sh.link = uint32(psess.elfshname(".dynstr").shnum)
			psess.
				shsym(sh, ctxt.Syms.Lookup(".gnu.version_r", 0))
		}

		if psess.elfRelType == ".rela" {
			sh := psess.elfshname(".rela.plt")
			sh.type_ = SHT_RELA
			sh.flags = SHF_ALLOC
			sh.entsize = ELF64RELASIZE
			sh.addralign = uint64(ctxt.Arch.RegSize)
			sh.link = uint32(psess.elfshname(".dynsym").shnum)
			sh.info = uint32(psess.elfshname(".plt").shnum)
			psess.
				shsym(sh, ctxt.Syms.Lookup(".rela.plt", 0))

			sh = psess.elfshname(".rela")
			sh.type_ = SHT_RELA
			sh.flags = SHF_ALLOC
			sh.entsize = ELF64RELASIZE
			sh.addralign = 8
			sh.link = uint32(psess.elfshname(".dynsym").shnum)
			psess.
				shsym(sh, ctxt.Syms.Lookup(".rela", 0))
		} else {
			sh := psess.elfshname(".rel.plt")
			sh.type_ = SHT_REL
			sh.flags = SHF_ALLOC
			sh.entsize = ELF32RELSIZE
			sh.addralign = 4
			sh.link = uint32(psess.elfshname(".dynsym").shnum)
			psess.
				shsym(sh, ctxt.Syms.Lookup(".rel.plt", 0))

			sh = psess.elfshname(".rel")
			sh.type_ = SHT_REL
			sh.flags = SHF_ALLOC
			sh.entsize = ELF32RELSIZE
			sh.addralign = 4
			sh.link = uint32(psess.elfshname(".dynsym").shnum)
			psess.
				shsym(sh, ctxt.Syms.Lookup(".rel", 0))
		}

		if eh.machine == EM_PPC64 {
			sh := psess.elfshname(".glink")
			sh.type_ = SHT_PROGBITS
			sh.flags = SHF_ALLOC + SHF_EXECINSTR
			sh.addralign = 4
			psess.
				shsym(sh, ctxt.Syms.Lookup(".glink", 0))
		}

		sh = psess.elfshname(".plt")
		sh.type_ = SHT_PROGBITS
		sh.flags = SHF_ALLOC + SHF_EXECINSTR
		if eh.machine == EM_X86_64 {
			sh.entsize = 16
		} else if eh.machine == EM_S390 {
			sh.entsize = 32
		} else if eh.machine == EM_PPC64 {

			sh.type_ = SHT_NOBITS

			sh.flags = SHF_ALLOC + SHF_WRITE
			sh.entsize = 8
		} else {
			sh.entsize = 4
		}
		sh.addralign = sh.entsize
		psess.
			shsym(sh, ctxt.Syms.Lookup(".plt", 0))

		if eh.machine != EM_PPC64 {
			sh := psess.elfshname(".got")
			sh.type_ = SHT_PROGBITS
			sh.flags = SHF_ALLOC + SHF_WRITE
			sh.entsize = uint64(ctxt.Arch.RegSize)
			sh.addralign = uint64(ctxt.Arch.RegSize)
			psess.
				shsym(sh, ctxt.Syms.Lookup(".got", 0))

			sh = psess.elfshname(".got.plt")
			sh.type_ = SHT_PROGBITS
			sh.flags = SHF_ALLOC + SHF_WRITE
			sh.entsize = uint64(ctxt.Arch.RegSize)
			sh.addralign = uint64(ctxt.Arch.RegSize)
			psess.
				shsym(sh, ctxt.Syms.Lookup(".got.plt", 0))
		}

		sh = psess.elfshname(".hash")
		sh.type_ = SHT_HASH
		sh.flags = SHF_ALLOC
		sh.entsize = 4
		sh.addralign = uint64(ctxt.Arch.RegSize)
		sh.link = uint32(psess.elfshname(".dynsym").shnum)
		psess.
			shsym(sh, ctxt.Syms.Lookup(".hash", 0))

		sh = psess.elfshname(".dynamic")

		sh.type_ = SHT_DYNAMIC
		sh.flags = SHF_ALLOC + SHF_WRITE
		sh.entsize = 2 * uint64(ctxt.Arch.RegSize)
		sh.addralign = uint64(ctxt.Arch.RegSize)
		sh.link = uint32(psess.elfshname(".dynstr").shnum)
		psess.
			shsym(sh, ctxt.Syms.Lookup(".dynamic", 0))
		ph := psess.newElfPhdr()
		ph.type_ = PT_DYNAMIC
		ph.flags = PF_R + PF_W
		phsh(ph, sh)

		tlssize := uint64(0)
		for _, sect := range psess.Segdata.Sections {
			if sect.Name == ".tbss" {
				tlssize = sect.Length
			}
		}
		if tlssize != 0 {
			ph := psess.newElfPhdr()
			ph.type_ = PT_TLS
			ph.flags = PF_R
			ph.memsz = tlssize
			ph.align = uint64(ctxt.Arch.RegSize)
		}
	}

	if ctxt.HeadType == objabi.Hlinux {
		ph := psess.newElfPhdr()
		ph.type_ = PT_GNU_STACK
		ph.flags = PF_W + PF_R
		ph.align = uint64(ctxt.Arch.RegSize)

		ph = psess.newElfPhdr()
		ph.type_ = PT_PAX_FLAGS
		ph.flags = 0x2a00
		ph.align = uint64(ctxt.Arch.RegSize)
	} else if ctxt.HeadType == objabi.Hsolaris {
		ph := psess.newElfPhdr()
		ph.type_ = PT_SUNWSTACK
		ph.flags = PF_W + PF_R
	}

elfobj:
	sh := psess.elfshname(".shstrtab")
	sh.type_ = SHT_STRTAB
	sh.addralign = 1
	psess.
		shsym(sh, ctxt.Syms.Lookup(".shstrtab", 0))
	eh.shstrndx = uint16(sh.shnum)

	if !*psess.FlagS {
		psess.
			elfshname(".symtab")
		psess.
			elfshname(".strtab")
	}

	for _, sect := range psess.Segtext.Sections {
		psess.
			elfshbits(ctxt.LinkMode, sect)
	}
	for _, sect := range psess.Segrodata.Sections {
		psess.
			elfshbits(ctxt.LinkMode, sect)
	}
	for _, sect := range psess.Segrelrodata.Sections {
		psess.
			elfshbits(ctxt.LinkMode, sect)
	}
	for _, sect := range psess.Segdata.Sections {
		psess.
			elfshbits(ctxt.LinkMode, sect)
	}
	for _, sect := range psess.Segdwarf.Sections {
		psess.
			elfshbits(ctxt.LinkMode, sect)
	}

	if ctxt.LinkMode == LinkExternal {
		for _, sect := range psess.Segtext.Sections {
			psess.
				elfshreloc(ctxt.Arch, sect)
		}
		for _, sect := range psess.Segrodata.Sections {
			psess.
				elfshreloc(ctxt.Arch, sect)
		}
		for _, sect := range psess.Segrelrodata.Sections {
			psess.
				elfshreloc(ctxt.Arch, sect)
		}
		for _, sect := range psess.Segdata.Sections {
			psess.
				elfshreloc(ctxt.Arch, sect)
		}
		for _, s := range psess.dwarfp {
			if len(s.R) > 0 || s.Type == sym.SDWARFINFO || s.Type == sym.SDWARFLOC {
				psess.
					elfshreloc(ctxt.Arch, s.Sect)
			}
		}

		sh := psess.elfshname(".note.GNU-stack")

		sh.type_ = SHT_PROGBITS
		sh.addralign = 1
		sh.flags = 0
	}

	if !*psess.FlagS {
		sh := psess.elfshname(".symtab")
		sh.type_ = SHT_SYMTAB
		sh.off = uint64(symo)
		sh.size = uint64(psess.Symsize)
		sh.addralign = uint64(ctxt.Arch.RegSize)
		sh.entsize = 8 + 2*uint64(ctxt.Arch.RegSize)
		sh.link = uint32(psess.elfshname(".strtab").shnum)
		sh.info = uint32(psess.elfglobalsymndx)

		sh = psess.elfshname(".strtab")
		sh.type_ = SHT_STRTAB
		sh.off = uint64(symo) + uint64(psess.Symsize)
		sh.size = uint64(len(psess.Elfstrdat))
		sh.addralign = 1
	}

	eh.ident[EI_MAG0] = '\177'

	eh.ident[EI_MAG1] = 'E'
	eh.ident[EI_MAG2] = 'L'
	eh.ident[EI_MAG3] = 'F'
	if ctxt.HeadType == objabi.Hfreebsd {
		eh.ident[EI_OSABI] = ELFOSABI_FREEBSD
	} else if ctxt.HeadType == objabi.Hnetbsd {
		eh.ident[EI_OSABI] = ELFOSABI_NETBSD
	} else if ctxt.HeadType == objabi.Hopenbsd {
		eh.ident[EI_OSABI] = ELFOSABI_OPENBSD
	} else if ctxt.HeadType == objabi.Hdragonfly {
		eh.ident[EI_OSABI] = ELFOSABI_NONE
	}
	if psess.elf64 {
		eh.ident[EI_CLASS] = ELFCLASS64
	} else {
		eh.ident[EI_CLASS] = ELFCLASS32
	}
	if ctxt.Arch.ByteOrder == binary.BigEndian {
		eh.ident[EI_DATA] = ELFDATA2MSB
	} else {
		eh.ident[EI_DATA] = ELFDATA2LSB
	}
	eh.ident[EI_VERSION] = EV_CURRENT

	if ctxt.LinkMode == LinkExternal {
		eh.type_ = ET_REL
	} else if ctxt.BuildMode == BuildModePIE {
		eh.type_ = ET_DYN
	} else {
		eh.type_ = ET_EXEC
	}

	if ctxt.LinkMode != LinkExternal {
		eh.entry = uint64(psess.Entryvalue(ctxt))
	}

	eh.version = EV_CURRENT

	if pph != nil {
		pph.filesz = uint64(eh.phnum) * uint64(eh.phentsize)
		pph.memsz = pph.filesz
	}

	ctxt.Out.SeekSet(psess, 0)
	a := int64(0)
	a += int64(psess.elfwritehdr(ctxt.Out))
	a += int64(psess.elfwritephdrs(ctxt.Out))
	a += int64(psess.elfwriteshdrs(ctxt.Out))
	if !*psess.FlagD {
		a += int64(psess.elfwriteinterp(ctxt.Out))
	}
	if ctxt.LinkMode != LinkExternal {
		if ctxt.HeadType == objabi.Hnetbsd {
			a += int64(psess.elfwritenetbsdsig(ctxt.Out))
		}
		if ctxt.HeadType == objabi.Hopenbsd {
			a += int64(psess.elfwriteopenbsdsig(ctxt.Out))
		}
		if len(psess.buildinfo) > 0 {
			a += int64(psess.elfwritebuildinfo(ctxt.Out))
		}
		if *psess.flagBuildid != "" {
			a += int64(psess.elfwritegobuildid(ctxt.Out))
		}
	}

	if a > elfreserve {
		psess.
			Errorf(nil, "ELFRESERVE too small: %d > %d with %d text sections", a, elfreserve, numtext)
	}
}

func (psess *PackageSession) elfadddynsym(ctxt *Link, s *sym.Symbol) {
	if psess.elf64 {
		s.Dynid = int32(psess.Nelfsym)
		psess.
			Nelfsym++

		d := ctxt.Syms.Lookup(".dynsym", 0)

		name := s.Extname
		d.AddUint32(ctxt.Arch, uint32(psess.Addstring(ctxt.Syms.Lookup(".dynstr", 0), name)))

		t := STB_GLOBAL << 4

		if s.Attr.CgoExport() && s.Type == sym.STEXT {
			t |= STT_FUNC
		} else {
			t |= STT_OBJECT
		}
		d.AddUint8(uint8(t))

		d.AddUint8(0)

		if s.Type == sym.SDYNIMPORT {
			d.AddUint16(ctxt.Arch, SHN_UNDEF)
		} else {
			d.AddUint16(ctxt.Arch, 1)
		}

		if s.Type == sym.SDYNIMPORT {
			d.AddUint64(ctxt.Arch, 0)
		} else {
			d.AddAddr(ctxt.Arch, s)
		}

		d.AddUint64(ctxt.Arch, uint64(s.Size))

		if ctxt.Arch.Family == sys.AMD64 && !s.Attr.CgoExportDynamic() && s.Dynimplib != "" && !psess.seenlib[s.Dynimplib] {
			psess.
				Elfwritedynent(ctxt, ctxt.Syms.Lookup(".dynamic", 0), DT_NEEDED, uint64(psess.Addstring(ctxt.Syms.Lookup(".dynstr", 0), s.Dynimplib)))
		}
	} else {
		s.Dynid = int32(psess.Nelfsym)
		psess.
			Nelfsym++

		d := ctxt.Syms.Lookup(".dynsym", 0)

		name := s.Extname

		d.AddUint32(ctxt.Arch, uint32(psess.Addstring(ctxt.Syms.Lookup(".dynstr", 0), name)))

		if s.Type == sym.SDYNIMPORT {
			d.AddUint32(ctxt.Arch, 0)
		} else {
			d.AddAddr(ctxt.Arch, s)
		}

		d.AddUint32(ctxt.Arch, uint32(s.Size))

		t := STB_GLOBAL << 4

		if ctxt.Arch.Family == sys.I386 && s.Attr.CgoExport() && s.Type == sym.STEXT {
			t |= STT_FUNC
		} else if ctxt.Arch.Family == sys.ARM && s.Attr.CgoExportDynamic() && s.Type == sym.STEXT {
			t |= STT_FUNC
		} else {
			t |= STT_OBJECT
		}
		d.AddUint8(uint8(t))
		d.AddUint8(0)

		if s.Type == sym.SDYNIMPORT {
			d.AddUint16(ctxt.Arch, SHN_UNDEF)
		} else {
			d.AddUint16(ctxt.Arch, 1)
		}
	}
}

func ELF32_R_SYM(info uint32) uint32 {
	return info >> 8
}

func ELF32_R_TYPE(info uint32) uint32 {
	return uint32(uint8(info))
}

func ELF32_R_INFO(sym uint32, type_ uint32) uint32 {
	return sym<<8 | type_
}

func ELF32_ST_BIND(info uint8) uint8 {
	return info >> 4
}

func ELF32_ST_TYPE(info uint8) uint8 {
	return info & 0xf
}

func ELF32_ST_INFO(bind uint8, type_ uint8) uint8 {
	return bind<<4 | type_&0xf
}

func ELF32_ST_VISIBILITY(oth uint8) uint8 {
	return oth & 3
}

func ELF64_R_SYM(info uint64) uint32 {
	return uint32(info >> 32)
}

func ELF64_R_TYPE(info uint64) uint32 {
	return uint32(info)
}

func ELF64_R_INFO(sym uint32, type_ uint32) uint64 {
	return uint64(sym)<<32 | uint64(type_)
}

func ELF64_ST_BIND(info uint8) uint8 {
	return info >> 4
}

func ELF64_ST_TYPE(info uint8) uint8 {
	return info & 0xf
}

func ELF64_ST_INFO(bind uint8, type_ uint8) uint8 {
	return bind<<4 | type_&0xf
}

func ELF64_ST_VISIBILITY(oth uint8) uint8 {
	return oth & 3
}

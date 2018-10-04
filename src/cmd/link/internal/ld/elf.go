// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
 * Derived from:
 * $FreeBSD: src/sys/sys/elf32.h,v 1.8.14.1 2005/12/30 22:13:58 marcel Exp $
 * $FreeBSD: src/sys/sys/elf64.h,v 1.10.14.1 2005/12/30 22:13:58 marcel Exp $
 * $FreeBSD: src/sys/sys/elf_common.h,v 1.15.8.1 2005/12/30 22:13:58 marcel Exp $
 * $FreeBSD: src/sys/alpha/include/elf.h,v 1.14 2003/09/25 01:10:22 peter Exp $
 * $FreeBSD: src/sys/amd64/include/elf.h,v 1.18 2004/08/03 08:21:48 dfr Exp $
 * $FreeBSD: src/sys/arm/include/elf.h,v 1.5.2.1 2006/06/30 21:42:52 cognet Exp $
 * $FreeBSD: src/sys/i386/include/elf.h,v 1.16 2004/08/02 19:12:17 dfr Exp $
 * $FreeBSD: src/sys/powerpc/include/elf.h,v 1.7 2004/11/02 09:47:01 ssouhlal Exp $
 * $FreeBSD: src/sys/sparc64/include/elf.h,v 1.12 2003/09/25 01:10:26 peter Exp $
 *
 * Copyright (c) 1996-1998 John D. Polstra.  All rights reserved.
 * Copyright (c) 2001 David E. O'Brien
 * Portions Copyright 2009 The Go Authors. All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 */

/*
 * ELF definitions that are independent of architecture or word size.
 */

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

/* For accessing the fields of r_info. */

/* For constructing r_info from field values. */

/*
 * Relocation types.
 */
const (
	ARM_MAGIC_TRAMP_NUMBER = 0x5c000003
)

/*
 * Symbol table entries.
 */

/* For accessing the fields of st_info. */

/* For constructing st_info from field values. */

/* For accessing the fields of st_other. */

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

/* For accessing the fields of r_info. */

/* For constructing r_info from field values. */

/*
 * Symbol table entries.
 */

/* For accessing the fields of st_info. */

/* For constructing st_info from field values. */

/* For accessing the fields of st_other. */

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

type Elfstring struct {
	s   string
	off int
}

/*
 Initialize the global variable that describes the ELF header. It will be updated as
 we write section and prog headers.
*/
func (pstate *PackageState) Elfinit(ctxt *Link) {
	ctxt.IsELF = true

	if ctxt.Arch.InFamily(sys.AMD64, sys.ARM64, sys.MIPS64, sys.PPC64, sys.S390X) {
		pstate.elfRelType = ".rela"
	} else {
		pstate.elfRelType = ".rel"
	}

	switch ctxt.Arch.Family {
	// 64-bit architectures
	case sys.PPC64, sys.S390X:
		if ctxt.Arch.ByteOrder == binary.BigEndian {
			pstate.ehdr.flags = 1 /* Version 1 ABI */
		} else {
			pstate.ehdr.flags = 2 /* Version 2 ABI */
		}
		fallthrough
	case sys.AMD64, sys.ARM64, sys.MIPS64:
		if ctxt.Arch.Family == sys.MIPS64 {
			pstate.ehdr.flags = 0x20000004 /* MIPS 3 CPIC */
		}
		pstate.elf64 = true

		pstate.ehdr.phoff = ELF64HDRSIZE      /* Must be be ELF64HDRSIZE: first PHdr must follow ELF header */
		pstate.ehdr.shoff = ELF64HDRSIZE      /* Will move as we add PHeaders */
		pstate.ehdr.ehsize = ELF64HDRSIZE     /* Must be ELF64HDRSIZE */
		pstate.ehdr.phentsize = ELF64PHDRSIZE /* Must be ELF64PHDRSIZE */
		pstate.ehdr.shentsize = ELF64SHDRSIZE /* Must be ELF64SHDRSIZE */

	// 32-bit architectures
	case sys.ARM, sys.MIPS:
		if ctxt.Arch.Family == sys.ARM {
			// we use EABI on linux/arm, freebsd/arm, netbsd/arm.
			if ctxt.HeadType == objabi.Hlinux || ctxt.HeadType == objabi.Hfreebsd || ctxt.HeadType == objabi.Hnetbsd {
				// We set a value here that makes no indication of which
				// float ABI the object uses, because this is information
				// used by the dynamic linker to compare executables and
				// shared libraries -- so it only matters for cgo calls, and
				// the information properly comes from the object files
				// produced by the host C compiler. parseArmAttributes in
				// ldelf.go reads that information and updates this field as
				// appropriate.
				pstate.ehdr.flags = 0x5000002 // has entry point, Version5 EABI
			}
		} else if ctxt.Arch.Family == sys.MIPS {
			pstate.ehdr.flags = 0x50001004 /* MIPS 32 CPIC O32*/
		}
		fallthrough
	default:
		pstate.ehdr.phoff = ELF32HDRSIZE
		/* Must be be ELF32HDRSIZE: first PHdr must follow ELF header */
		pstate.ehdr.shoff = ELF32HDRSIZE      /* Will move as we add PHeaders */
		pstate.ehdr.ehsize = ELF32HDRSIZE     /* Must be ELF32HDRSIZE */
		pstate.ehdr.phentsize = ELF32PHDRSIZE /* Must be ELF32PHDRSIZE */
		pstate.ehdr.shentsize = ELF32SHDRSIZE /* Must be ELF32SHDRSIZE */
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

func (pstate *PackageState) elfwriteshdrs(out *OutBuf) uint32 {
	if pstate.elf64 {
		for i := 0; i < int(pstate.ehdr.shnum); i++ {
			elf64shdr(out, pstate.shdr[i])
		}
		return uint32(pstate.ehdr.shnum) * ELF64SHDRSIZE
	}

	for i := 0; i < int(pstate.ehdr.shnum); i++ {
		elf32shdr(out, pstate.shdr[i])
	}
	return uint32(pstate.ehdr.shnum) * ELF32SHDRSIZE
}

func (pstate *PackageState) elfsetstring(s *sym.Symbol, str string, off int) {
	if pstate.nelfstr >= len(pstate.elfstr) {
		pstate.Errorf(s, "too many elf strings")
		pstate.errorexit()
	}

	pstate.elfstr[pstate.nelfstr].s = str
	pstate.elfstr[pstate.nelfstr].off = off
	pstate.nelfstr++
}

func (pstate *PackageState) elfwritephdrs(out *OutBuf) uint32 {
	if pstate.elf64 {
		for i := 0; i < int(pstate.ehdr.phnum); i++ {
			elf64phdr(out, pstate.phdr[i])
		}
		return uint32(pstate.ehdr.phnum) * ELF64PHDRSIZE
	}

	for i := 0; i < int(pstate.ehdr.phnum); i++ {
		elf32phdr(out, pstate.phdr[i])
	}
	return uint32(pstate.ehdr.phnum) * ELF32PHDRSIZE
}

func (pstate *PackageState) newElfPhdr() *ElfPhdr {
	e := new(ElfPhdr)
	if pstate.ehdr.phnum >= NSECT {
		pstate.Errorf(nil, "too many phdrs")
	} else {
		pstate.phdr[pstate.ehdr.phnum] = e
		pstate.ehdr.phnum++
	}
	if pstate.elf64 {
		pstate.ehdr.shoff += ELF64PHDRSIZE
	} else {
		pstate.ehdr.shoff += ELF32PHDRSIZE
	}
	return e
}

func (pstate *PackageState) newElfShdr(name int64) *ElfShdr {
	e := new(ElfShdr)
	e.name = uint32(name)
	e.shnum = int(pstate.ehdr.shnum)
	if pstate.ehdr.shnum >= NSECT {
		pstate.Errorf(nil, "too many shdrs")
	} else {
		pstate.shdr[pstate.ehdr.shnum] = e
		pstate.ehdr.shnum++
	}

	return e
}

func (pstate *PackageState) getElfEhdr() *ElfEhdr {
	return &pstate.ehdr
}

func (pstate *PackageState) elf64writehdr(out *OutBuf) uint32 {
	out.Write(pstate.ehdr.ident[:])
	out.Write16(pstate.ehdr.type_)
	out.Write16(pstate.ehdr.machine)
	out.Write32(pstate.ehdr.version)
	out.Write64(pstate.ehdr.entry)
	out.Write64(pstate.ehdr.phoff)
	out.Write64(pstate.ehdr.shoff)
	out.Write32(pstate.ehdr.flags)
	out.Write16(pstate.ehdr.ehsize)
	out.Write16(pstate.ehdr.phentsize)
	out.Write16(pstate.ehdr.phnum)
	out.Write16(pstate.ehdr.shentsize)
	out.Write16(pstate.ehdr.shnum)
	out.Write16(pstate.ehdr.shstrndx)
	return ELF64HDRSIZE
}

func (pstate *PackageState) elf32writehdr(out *OutBuf) uint32 {
	out.Write(pstate.ehdr.ident[:])
	out.Write16(pstate.ehdr.type_)
	out.Write16(pstate.ehdr.machine)
	out.Write32(pstate.ehdr.version)
	out.Write32(uint32(pstate.ehdr.entry))
	out.Write32(uint32(pstate.ehdr.phoff))
	out.Write32(uint32(pstate.ehdr.shoff))
	out.Write32(pstate.ehdr.flags)
	out.Write16(pstate.ehdr.ehsize)
	out.Write16(pstate.ehdr.phentsize)
	out.Write16(pstate.ehdr.phnum)
	out.Write16(pstate.ehdr.shentsize)
	out.Write16(pstate.ehdr.shnum)
	out.Write16(pstate.ehdr.shstrndx)
	return ELF32HDRSIZE
}

func (pstate *PackageState) elfwritehdr(out *OutBuf) uint32 {
	if pstate.elf64 {
		return pstate.elf64writehdr(out)
	}
	return pstate.elf32writehdr(out)
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

func (pstate *PackageState) Elfwritedynent(ctxt *Link, s *sym.Symbol, tag int, val uint64) {
	if pstate.elf64 {
		s.AddUint64(ctxt.Arch, uint64(tag))
		s.AddUint64(ctxt.Arch, val)
	} else {
		s.AddUint32(ctxt.Arch, uint32(tag))
		s.AddUint32(ctxt.Arch, uint32(val))
	}
}

func (pstate *PackageState) elfwritedynentsym(ctxt *Link, s *sym.Symbol, tag int, t *sym.Symbol) {
	pstate.Elfwritedynentsymplus(ctxt, s, tag, t, 0)
}

func (pstate *PackageState) Elfwritedynentsymplus(ctxt *Link, s *sym.Symbol, tag int, t *sym.Symbol, add int64) {
	if pstate.elf64 {
		s.AddUint64(ctxt.Arch, uint64(tag))
	} else {
		s.AddUint32(ctxt.Arch, uint32(tag))
	}
	s.AddAddrPlus(ctxt.Arch, t, add)
}

func (pstate *PackageState) elfwritedynentsymsize(ctxt *Link, s *sym.Symbol, tag int, t *sym.Symbol) {
	if pstate.elf64 {
		s.AddUint64(ctxt.Arch, uint64(tag))
	} else {
		s.AddUint32(ctxt.Arch, uint32(tag))
	}
	s.AddSize(ctxt.Arch, t)
}

func (pstate *PackageState) elfinterp(sh *ElfShdr, startva uint64, resoff uint64, p string) int {
	pstate.interp = p
	n := len(pstate.interp) + 1
	sh.addr = startva + resoff - uint64(n)
	sh.off = resoff - uint64(n)
	sh.size = uint64(n)

	return n
}

func (pstate *PackageState) elfwriteinterp(out *OutBuf) int {
	sh := pstate.elfshname(".interp")
	out.SeekSet(pstate, int64(sh.off))
	out.WriteString(pstate.interp)
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

func (pstate *PackageState) elfwritenotehdr(out *OutBuf, str string, namesz uint32, descsz uint32, tag uint32) *ElfShdr {
	sh := pstate.elfshname(str)

	// Write Elf_Note header.
	out.SeekSet(pstate, int64(sh.off))

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

func (pstate *PackageState) elfwritenetbsdsig(out *OutBuf) int {
	// Write Elf_Note header.
	sh := pstate.elfwritenotehdr(out, ".note.netbsd.ident", ELF_NOTE_NETBSD_NAMESZ, ELF_NOTE_NETBSD_DESCSZ, ELF_NOTE_NETBSD_TAG)

	if sh == nil {
		return 0
	}

	// Followed by NetBSD string and version.
	out.Write(pstate.ELF_NOTE_NETBSD_NAME)
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

func (pstate *PackageState) elfwriteopenbsdsig(out *OutBuf) int {
	// Write Elf_Note header.
	sh := pstate.elfwritenotehdr(out, ".note.openbsd.ident", ELF_NOTE_OPENBSD_NAMESZ, ELF_NOTE_OPENBSD_DESCSZ, ELF_NOTE_OPENBSD_TAG)

	if sh == nil {
		return 0
	}

	// Followed by OpenBSD string and version.
	out.Write(pstate.ELF_NOTE_OPENBSD_NAME)

	out.Write32(ELF_NOTE_OPENBSD_VERSION)

	return int(sh.size)
}

func (pstate *PackageState) addbuildinfo(val string) {
	if !strings.HasPrefix(val, "0x") {
		pstate.Exitf("-B argument must start with 0x: %s", val)
	}

	ov := val
	val = val[2:]

	const maxLen = 32
	if hex.DecodedLen(len(val)) > maxLen {
		pstate.Exitf("-B option too long (max %d digits): %s", maxLen, ov)
	}

	b, err := hex.DecodeString(val)
	if err != nil {
		if err == hex.ErrLength {
			pstate.Exitf("-B argument must have even number of digits: %s", ov)
		}
		if inv, ok := err.(hex.InvalidByteError); ok {
			pstate.Exitf("-B argument contains invalid hex digit %c: %s", byte(inv), ov)
		}
		pstate.Exitf("-B argument contains invalid hex: %s", ov)
	}

	pstate.buildinfo = b
}

// Build info note
const (
	ELF_NOTE_BUILDINFO_NAMESZ = 4
	ELF_NOTE_BUILDINFO_TAG    = 3
)

func (pstate *PackageState) elfbuildinfo(sh *ElfShdr, startva uint64, resoff uint64) int {
	n := int(ELF_NOTE_BUILDINFO_NAMESZ + Rnd(int64(len(pstate.buildinfo)), 4))
	return elfnote(sh, startva, resoff, n)
}

func (pstate *PackageState) elfgobuildid(sh *ElfShdr, startva uint64, resoff uint64) int {
	n := len(pstate.ELF_NOTE_GO_NAME) + int(Rnd(int64(len(*pstate.flagBuildid)), 4))
	return elfnote(sh, startva, resoff, n)
}

func (pstate *PackageState) elfwritebuildinfo(out *OutBuf) int {
	sh := pstate.elfwritenotehdr(out, ".note.gnu.build-id", ELF_NOTE_BUILDINFO_NAMESZ, uint32(len(pstate.buildinfo)), ELF_NOTE_BUILDINFO_TAG)
	if sh == nil {
		return 0
	}

	out.Write(pstate.ELF_NOTE_BUILDINFO_NAME)
	out.Write(pstate.buildinfo)
	var zero = make([]byte, 4)
	out.Write(zero[:int(Rnd(int64(len(pstate.buildinfo)), 4)-int64(len(pstate.buildinfo)))])

	return int(sh.size)
}

func (pstate *PackageState) elfwritegobuildid(out *OutBuf) int {
	sh := pstate.elfwritenotehdr(out, ".note.go.buildid", uint32(len(pstate.ELF_NOTE_GO_NAME)), uint32(len(*pstate.flagBuildid)), ELF_NOTE_GOBUILDID_TAG)
	if sh == nil {
		return 0
	}

	out.Write(pstate.ELF_NOTE_GO_NAME)
	out.Write([]byte(*pstate.flagBuildid))
	var zero = make([]byte, 4)
	out.Write(zero[:int(Rnd(int64(len(*pstate.flagBuildid)), 4)-int64(len(*pstate.flagBuildid)))])

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

func (pstate *PackageState) elfdynhash(ctxt *Link) {
	if !ctxt.IsELF {
		return
	}

	nsym := pstate.Nelfsym
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

	// s390x (ELF64) hash table entries are 8 bytes
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

	// version symbols
	dynstr := ctxt.Syms.Lookup(".dynstr", 0)

	s = ctxt.Syms.Lookup(".gnu.version_r", 0)
	i = 2
	nfile := 0
	for l := needlib; l != nil; l = l.next {
		nfile++

		// header
		s.AddUint16(ctxt.Arch, 1) // table version
		j := 0
		for x := l.aux; x != nil; x = x.next {
			j++
		}
		s.AddUint16(ctxt.Arch, uint16(j))                                // aux count
		s.AddUint32(ctxt.Arch, uint32(pstate.Addstring(dynstr, l.file))) // file string offset
		s.AddUint32(ctxt.Arch, 16)                                       // offset from header to first aux
		if l.next != nil {
			s.AddUint32(ctxt.Arch, 16+uint32(j)*16) // offset from this header to next
		} else {
			s.AddUint32(ctxt.Arch, 0)
		}

		for x := l.aux; x != nil; x = x.next {
			x.num = i
			i++

			// aux struct
			s.AddUint32(ctxt.Arch, elfhash(x.vers))                          // hash
			s.AddUint16(ctxt.Arch, 0)                                        // flags
			s.AddUint16(ctxt.Arch, uint16(x.num))                            // other - index we refer to this by
			s.AddUint32(ctxt.Arch, uint32(pstate.Addstring(dynstr, x.vers))) // version string offset
			if x.next != nil {
				s.AddUint32(ctxt.Arch, 16) // offset from this aux to next
			} else {
				s.AddUint32(ctxt.Arch, 0)
			}
		}
	}

	// version references
	s = ctxt.Syms.Lookup(".gnu.version", 0)

	for i := 0; i < nsym; i++ {
		if i == 0 {
			s.AddUint16(ctxt.Arch, 0) // first entry - no symbol
		} else if need[i] == nil {
			s.AddUint16(ctxt.Arch, 1) // global
		} else {
			s.AddUint16(ctxt.Arch, uint16(need[i].num))
		}
	}

	s = ctxt.Syms.Lookup(".dynamic", 0)
	pstate.elfverneed = nfile
	if pstate.elfverneed != 0 {
		pstate.elfwritedynentsym(ctxt, s, DT_VERNEED, ctxt.Syms.Lookup(".gnu.version_r", 0))
		pstate.Elfwritedynent(ctxt, s, DT_VERNEEDNUM, uint64(nfile))
		pstate.elfwritedynentsym(ctxt, s, DT_VERSYM, ctxt.Syms.Lookup(".gnu.version", 0))
	}

	sy := ctxt.Syms.Lookup(pstate.elfRelType+".plt", 0)
	if sy.Size > 0 {
		if pstate.elfRelType == ".rela" {
			pstate.Elfwritedynent(ctxt, s, DT_PLTREL, DT_RELA)
		} else {
			pstate.Elfwritedynent(ctxt, s, DT_PLTREL, DT_REL)
		}
		pstate.elfwritedynentsymsize(ctxt, s, DT_PLTRELSZ, sy)
		pstate.elfwritedynentsym(ctxt, s, DT_JMPREL, sy)
	}

	pstate.Elfwritedynent(ctxt, s, DT_NULL, 0)
}

func (pstate *PackageState) elfphload(seg *sym.Segment) *ElfPhdr {
	ph := pstate.newElfPhdr()
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
	ph.align = uint64(*pstate.FlagRound)

	return ph
}

func (pstate *PackageState) elfphrelro(seg *sym.Segment) {
	ph := pstate.newElfPhdr()
	ph.type_ = PT_GNU_RELRO
	ph.vaddr = seg.Vaddr
	ph.paddr = seg.Vaddr
	ph.memsz = seg.Length
	ph.off = seg.Fileoff
	ph.filesz = seg.Filelen
	ph.align = uint64(*pstate.FlagRound)
}

func (pstate *PackageState) elfshname(name string) *ElfShdr {
	for i := 0; i < pstate.nelfstr; i++ {
		if name != pstate.elfstr[i].s {
			continue
		}
		off := pstate.elfstr[i].off
		for i = 0; i < int(pstate.ehdr.shnum); i++ {
			sh := pstate.shdr[i]
			if sh.name == uint32(off) {
				return sh
			}
		}
		return pstate.newElfShdr(int64(off))
	}
	pstate.Exitf("cannot find elf name %s", name)
	return nil
}

// Create an ElfShdr for the section with name.
// Create a duplicate if one already exists with that name
func (pstate *PackageState) elfshnamedup(name string) *ElfShdr {
	for i := 0; i < pstate.nelfstr; i++ {
		if name == pstate.elfstr[i].s {
			off := pstate.elfstr[i].off
			return pstate.newElfShdr(int64(off))
		}
	}

	pstate.Errorf(nil, "cannot find elf name %s", name)
	pstate.errorexit()
	return nil
}

func (pstate *PackageState) elfshalloc(sect *sym.Section) *ElfShdr {
	sh := pstate.elfshname(sect.Name)
	sect.Elfsect = sh
	return sh
}

func (pstate *PackageState) elfshbits(linkmode LinkMode, sect *sym.Section) *ElfShdr {
	var sh *ElfShdr

	if sect.Name == ".text" {
		if sect.Elfsect == nil {
			sect.Elfsect = pstate.elfshnamedup(sect.Name)
		}
		sh = sect.Elfsect.(*ElfShdr)
	} else {
		sh = pstate.elfshalloc(sect)
	}

	// If this section has already been set up as a note, we assume type_ and
	// flags are already correct, but the other fields still need filling in.
	if sh.type_ == SHT_NOTE {
		if linkmode != LinkExternal {
			// TODO(mwhudson): the approach here will work OK when
			// linking internally for notes that we want to be included
			// in a loadable segment (e.g. the abihash note) but not for
			// notes that we do not want to be mapped (e.g. the package
			// list note). The real fix is probably to define new values
			// for Symbol.Type corresponding to mapped and unmapped notes
			// and handle them in dodata().
			pstate.Errorf(nil, "sh.type_ == SHT_NOTE in elfshbits when linking internally")
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

func (pstate *PackageState) elfshreloc(arch *sys.Arch, sect *sym.Section) *ElfShdr {
	// If main section is SHT_NOBITS, nothing to relocate.
	// Also nothing to relocate in .shstrtab or notes.
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
	if pstate.elfRelType == ".rela" {
		typ = SHT_RELA
	} else {
		typ = SHT_REL
	}

	sh := pstate.elfshname(pstate.elfRelType + sect.Name)
	// There could be multiple text sections but each needs
	// its own .rela.text.

	if sect.Name == ".text" {
		if sh.info != 0 && sh.info != uint32(sect.Elfsect.(*ElfShdr).shnum) {
			sh = pstate.elfshnamedup(pstate.elfRelType + sect.Name)
		}
	}

	sh.type_ = uint32(typ)
	sh.entsize = uint64(arch.RegSize) * 2
	if typ == SHT_RELA {
		sh.entsize += uint64(arch.RegSize)
	}
	sh.link = uint32(pstate.elfshname(".symtab").shnum)
	sh.info = uint32(sect.Elfsect.(*ElfShdr).shnum)
	sh.off = sect.Reloff
	sh.size = sect.Rellen
	sh.addralign = uint64(arch.RegSize)
	return sh
}

func (pstate *PackageState) elfrelocsect(ctxt *Link, sect *sym.Section, syms []*sym.Symbol) {
	// If main section is SHT_NOBITS, nothing to relocate.
	// Also nothing to relocate in .shstrtab.
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
				pstate.Errorf(s, "missing xsym in relocation %#v %#v", r.Sym.Name, s)
				continue
			}
			if r.Xsym.ElfsymForReloc() == 0 {
				pstate.Errorf(s, "reloc %d (%s) to non-elf symbol %s (outer=%s) %d (%s)", r.Type, pstate.sym.RelocName(ctxt.Arch, r.Type), r.Sym.Name, r.Xsym.Name, r.Sym.Type, r.Sym.Type)
			}
			if !r.Xsym.Attr.Reachable() {
				pstate.Errorf(s, "unreachable reloc %d (%s) target %v", r.Type, pstate.sym.RelocName(ctxt.Arch, r.Type), r.Xsym.Name)
			}
			if !pstate.thearch.Elfreloc1(ctxt, r, int64(uint64(s.Value+int64(r.Off))-sect.Vaddr)) {
				pstate.Errorf(s, "unsupported obj reloc %d (%s)/%d to %s", r.Type, pstate.sym.RelocName(ctxt.Arch, r.Type), r.Siz, r.Sym.Name)
			}
		}
	}

	sect.Rellen = uint64(ctxt.Out.Offset()) - sect.Reloff
}

func (pstate *PackageState) Elfemitreloc(ctxt *Link) {
	for ctxt.Out.Offset()&7 != 0 {
		ctxt.Out.Write8(0)
	}

	for _, sect := range pstate.Segtext.Sections {
		if sect.Name == ".text" {
			pstate.elfrelocsect(ctxt, sect, ctxt.Textp)
		} else {
			pstate.elfrelocsect(ctxt, sect, pstate.datap)
		}
	}

	for _, sect := range pstate.Segrodata.Sections {
		pstate.elfrelocsect(ctxt, sect, pstate.datap)
	}
	for _, sect := range pstate.Segrelrodata.Sections {
		pstate.elfrelocsect(ctxt, sect, pstate.datap)
	}
	for _, sect := range pstate.Segdata.Sections {
		pstate.elfrelocsect(ctxt, sect, pstate.datap)
	}
	for _, sect := range pstate.Segdwarf.Sections {
		pstate.elfrelocsect(ctxt, sect, pstate.dwarfp)
	}
}

func (pstate *PackageState) addgonote(ctxt *Link, sectionName string, tag uint32, desc []byte) {
	s := ctxt.Syms.Lookup(sectionName, 0)
	s.Attr |= sym.AttrReachable
	s.Type = sym.SELFROSECT
	// namesz
	s.AddUint32(ctxt.Arch, uint32(len(pstate.ELF_NOTE_GO_NAME)))
	// descsz
	s.AddUint32(ctxt.Arch, uint32(len(desc)))
	// tag
	s.AddUint32(ctxt.Arch, tag)
	// name + padding
	s.P = append(s.P, pstate.ELF_NOTE_GO_NAME...)
	for len(s.P)%4 != 0 {
		s.P = append(s.P, 0)
	}
	// desc + padding
	s.P = append(s.P, desc...)
	for len(s.P)%4 != 0 {
		s.P = append(s.P, 0)
	}
	s.Size = int64(len(s.P))
	s.Align = 4
}

func (ctxt *Link) doelf(pstate *PackageState) {
	if !ctxt.IsELF {
		return
	}

	/* predefine strings we need for section headers */
	shstrtab := ctxt.Syms.Lookup(".shstrtab", 0)

	shstrtab.Type = sym.SELFROSECT
	shstrtab.Attr |= sym.AttrReachable

	pstate.Addstring(shstrtab, "")
	pstate.Addstring(shstrtab, ".text")
	pstate.Addstring(shstrtab, ".noptrdata")
	pstate.Addstring(shstrtab, ".data")
	pstate.Addstring(shstrtab, ".bss")
	pstate.Addstring(shstrtab, ".noptrbss")

	// generate .tbss section for dynamic internal linker or external
	// linking, so that various binutils could correctly calculate
	// PT_TLS size. See https://golang.org/issue/5200.
	if !*pstate.FlagD || ctxt.LinkMode == LinkExternal {
		pstate.Addstring(shstrtab, ".tbss")
	}
	if ctxt.HeadType == objabi.Hnetbsd {
		pstate.Addstring(shstrtab, ".note.netbsd.ident")
	}
	if ctxt.HeadType == objabi.Hopenbsd {
		pstate.Addstring(shstrtab, ".note.openbsd.ident")
	}
	if len(pstate.buildinfo) > 0 {
		pstate.Addstring(shstrtab, ".note.gnu.build-id")
	}
	if *pstate.flagBuildid != "" {
		pstate.Addstring(shstrtab, ".note.go.buildid")
	}
	pstate.Addstring(shstrtab, ".elfdata")
	pstate.Addstring(shstrtab, ".rodata")
	// See the comment about data.rel.ro.FOO section names in data.go.
	relro_prefix := ""
	if ctxt.UseRelro() {
		pstate.Addstring(shstrtab, ".data.rel.ro")
		relro_prefix = ".data.rel.ro"
	}
	pstate.Addstring(shstrtab, relro_prefix+".typelink")
	pstate.Addstring(shstrtab, relro_prefix+".itablink")
	pstate.Addstring(shstrtab, relro_prefix+".gosymtab")
	pstate.Addstring(shstrtab, relro_prefix+".gopclntab")

	if ctxt.LinkMode == LinkExternal {
		*pstate.FlagD = true

		pstate.Addstring(shstrtab, pstate.elfRelType+".text")
		pstate.Addstring(shstrtab, pstate.elfRelType+".rodata")
		pstate.Addstring(shstrtab, pstate.elfRelType+relro_prefix+".typelink")
		pstate.Addstring(shstrtab, pstate.elfRelType+relro_prefix+".itablink")
		pstate.Addstring(shstrtab, pstate.elfRelType+relro_prefix+".gosymtab")
		pstate.Addstring(shstrtab, pstate.elfRelType+relro_prefix+".gopclntab")
		pstate.Addstring(shstrtab, pstate.elfRelType+".noptrdata")
		pstate.Addstring(shstrtab, pstate.elfRelType+".data")
		if ctxt.UseRelro() {
			pstate.Addstring(shstrtab, pstate.elfRelType+".data.rel.ro")
		}

		// add a .note.GNU-stack section to mark the stack as non-executable
		pstate.Addstring(shstrtab, ".note.GNU-stack")

		if ctxt.BuildMode == BuildModeShared {
			pstate.Addstring(shstrtab, ".note.go.abihash")
			pstate.Addstring(shstrtab, ".note.go.pkg-list")
			pstate.Addstring(shstrtab, ".note.go.deps")
		}
	}

	hasinitarr := ctxt.linkShared

	/* shared library initializer */
	switch ctxt.BuildMode {
	case BuildModeCArchive, BuildModeCShared, BuildModeShared, BuildModePlugin:
		hasinitarr = true
	}

	if hasinitarr {
		pstate.Addstring(shstrtab, ".init_array")
		pstate.Addstring(shstrtab, pstate.elfRelType+".init_array")
	}

	if !*pstate.FlagS {
		pstate.Addstring(shstrtab, ".symtab")
		pstate.Addstring(shstrtab, ".strtab")
		pstate.dwarfaddshstrings(ctxt, shstrtab)
	}

	pstate.Addstring(shstrtab, ".shstrtab")

	if !*pstate.FlagD { /* -d suppresses dynamic loader format */
		pstate.Addstring(shstrtab, ".interp")
		pstate.Addstring(shstrtab, ".hash")
		pstate.Addstring(shstrtab, ".got")
		if ctxt.Arch.Family == sys.PPC64 {
			pstate.Addstring(shstrtab, ".glink")
		}
		pstate.Addstring(shstrtab, ".got.plt")
		pstate.Addstring(shstrtab, ".dynamic")
		pstate.Addstring(shstrtab, ".dynsym")
		pstate.Addstring(shstrtab, ".dynstr")
		pstate.Addstring(shstrtab, pstate.elfRelType)
		pstate.Addstring(shstrtab, pstate.elfRelType+".plt")

		pstate.Addstring(shstrtab, ".plt")
		pstate.Addstring(shstrtab, ".gnu.version")
		pstate.Addstring(shstrtab, ".gnu.version_r")

		/* dynamic symbol table - first entry all zeros */
		s := ctxt.Syms.Lookup(".dynsym", 0)

		s.Type = sym.SELFROSECT
		s.Attr |= sym.AttrReachable
		if pstate.elf64 {
			s.Size += ELF64SYMSIZE
		} else {
			s.Size += ELF32SYMSIZE
		}

		/* dynamic string table */
		s = ctxt.Syms.Lookup(".dynstr", 0)

		s.Type = sym.SELFROSECT
		s.Attr |= sym.AttrReachable
		if s.Size == 0 {
			pstate.Addstring(s, "")
		}
		dynstr := s

		/* relocation table */
		s = ctxt.Syms.Lookup(pstate.elfRelType, 0)
		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFROSECT

		/* global offset table */
		s = ctxt.Syms.Lookup(".got", 0)

		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFGOT // writable

		/* ppc64 glink resolver */
		if ctxt.Arch.Family == sys.PPC64 {
			s := ctxt.Syms.Lookup(".glink", 0)
			s.Attr |= sym.AttrReachable
			s.Type = sym.SELFRXSECT
		}

		/* hash */
		s = ctxt.Syms.Lookup(".hash", 0)

		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFROSECT

		s = ctxt.Syms.Lookup(".got.plt", 0)
		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFSECT // writable

		s = ctxt.Syms.Lookup(".plt", 0)

		s.Attr |= sym.AttrReachable
		if ctxt.Arch.Family == sys.PPC64 {
			// In the ppc64 ABI, .plt is a data section
			// written by the dynamic linker.
			s.Type = sym.SELFSECT
		} else {
			s.Type = sym.SELFRXSECT
		}

		pstate.thearch.Elfsetupplt(ctxt)

		s = ctxt.Syms.Lookup(pstate.elfRelType+".plt", 0)
		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFROSECT

		s = ctxt.Syms.Lookup(".gnu.version", 0)
		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFROSECT

		s = ctxt.Syms.Lookup(".gnu.version_r", 0)
		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFROSECT

		/* define dynamic elf table */
		s = ctxt.Syms.Lookup(".dynamic", 0)

		s.Attr |= sym.AttrReachable
		s.Type = sym.SELFSECT // writable

		/*
		 * .dynamic table
		 */
		pstate.elfwritedynentsym(ctxt, s, DT_HASH, ctxt.Syms.Lookup(".hash", 0))

		pstate.elfwritedynentsym(ctxt, s, DT_SYMTAB, ctxt.Syms.Lookup(".dynsym", 0))
		if pstate.elf64 {
			pstate.Elfwritedynent(ctxt, s, DT_SYMENT, ELF64SYMSIZE)
		} else {
			pstate.Elfwritedynent(ctxt, s, DT_SYMENT, ELF32SYMSIZE)
		}
		pstate.elfwritedynentsym(ctxt, s, DT_STRTAB, ctxt.Syms.Lookup(".dynstr", 0))
		pstate.elfwritedynentsymsize(ctxt, s, DT_STRSZ, ctxt.Syms.Lookup(".dynstr", 0))
		if pstate.elfRelType == ".rela" {
			pstate.elfwritedynentsym(ctxt, s, DT_RELA, ctxt.Syms.Lookup(".rela", 0))
			pstate.elfwritedynentsymsize(ctxt, s, DT_RELASZ, ctxt.Syms.Lookup(".rela", 0))
			pstate.Elfwritedynent(ctxt, s, DT_RELAENT, ELF64RELASIZE)
		} else {
			pstate.elfwritedynentsym(ctxt, s, DT_REL, ctxt.Syms.Lookup(".rel", 0))
			pstate.elfwritedynentsymsize(ctxt, s, DT_RELSZ, ctxt.Syms.Lookup(".rel", 0))
			pstate.Elfwritedynent(ctxt, s, DT_RELENT, ELF32RELSIZE)
		}

		if pstate.rpath.val != "" {
			pstate.Elfwritedynent(ctxt, s, DT_RUNPATH, uint64(pstate.Addstring(dynstr, pstate.rpath.val)))
		}

		if ctxt.Arch.Family == sys.PPC64 {
			pstate.elfwritedynentsym(ctxt, s, DT_PLTGOT, ctxt.Syms.Lookup(".plt", 0))
		} else if ctxt.Arch.Family == sys.S390X {
			pstate.elfwritedynentsym(ctxt, s, DT_PLTGOT, ctxt.Syms.Lookup(".got", 0))
		} else {
			pstate.elfwritedynentsym(ctxt, s, DT_PLTGOT, ctxt.Syms.Lookup(".got.plt", 0))
		}

		if ctxt.Arch.Family == sys.PPC64 {
			pstate.Elfwritedynent(ctxt, s, DT_PPC64_OPT, 0)
		}

		// Solaris dynamic linker can't handle an empty .rela.plt if
		// DT_JMPREL is emitted so we have to defer generation of DT_PLTREL,
		// DT_PLTRELSZ, and DT_JMPREL dynamic entries until after we know the
		// size of .rel(a).plt section.
		pstate.Elfwritedynent(ctxt, s, DT_DEBUG, 0)
	}

	if ctxt.BuildMode == BuildModeShared {
		// The go.link.abihashbytes symbol will be pointed at the appropriate
		// part of the .note.go.abihash section in data.go:func address().
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
		pstate.addgonote(ctxt, ".note.go.abihash", ELF_NOTE_GOABIHASH_TAG, h.Sum([]byte{}))
		pstate.addgonote(ctxt, ".note.go.pkg-list", ELF_NOTE_GOPKGLIST_TAG, pstate.pkglistfornote)
		var deplist []string
		for _, shlib := range ctxt.Shlibs {
			deplist = append(deplist, filepath.Base(shlib.Path))
		}
		pstate.addgonote(ctxt, ".note.go.deps", ELF_NOTE_GODEPS_TAG, []byte(strings.Join(deplist, "\n")))
	}

	if ctxt.LinkMode == LinkExternal && *pstate.flagBuildid != "" {
		pstate.addgonote(ctxt, ".note.go.buildid", ELF_NOTE_GOBUILDID_TAG, []byte(*pstate.flagBuildid))
	}
}

// Do not write DT_NULL.  elfdynhash will finish it.
func (pstate *PackageState) shsym(sh *ElfShdr, s *sym.Symbol) {
	addr := pstate.Symaddr(s)
	if sh.flags&SHF_ALLOC != 0 {
		sh.addr = uint64(addr)
	}
	sh.off = uint64(pstate.datoff(s, addr))
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

func (pstate *PackageState) Asmbelfsetup() {
	/* This null SHdr must appear before all others */
	pstate.elfshname("")

	for _, sect := range pstate.Segtext.Sections {
		// There could be multiple .text sections. Instead check the Elfsect
		// field to determine if already has an ElfShdr and if not, create one.
		if sect.Name == ".text" {
			if sect.Elfsect == nil {
				sect.Elfsect = pstate.elfshnamedup(sect.Name)
			}
		} else {
			pstate.elfshalloc(sect)
		}
	}
	for _, sect := range pstate.Segrodata.Sections {
		pstate.elfshalloc(sect)
	}
	for _, sect := range pstate.Segrelrodata.Sections {
		pstate.elfshalloc(sect)
	}
	for _, sect := range pstate.Segdata.Sections {
		pstate.elfshalloc(sect)
	}
	for _, sect := range pstate.Segdwarf.Sections {
		pstate.elfshalloc(sect)
	}
}

func (pstate *PackageState) Asmbelf(ctxt *Link, symo int64) {
	eh := pstate.getElfEhdr()
	switch ctxt.Arch.Family {
	default:
		pstate.Exitf("unknown architecture in asmbelf: %v", ctxt.Arch.Family)
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
	for _, sect := range pstate.Segtext.Sections {
		if sect.Name == ".text" {
			numtext++
		}
	}

	// If there are multiple text sections, extra space is needed
	// in the elfreserve for the additional .text and .rela.text
	// section headers.  It can handle 4 extra now. Headers are
	// 64 bytes.

	if numtext > 4 {
		elfreserve += elfreserve + numtext*64*2
	}

	startva := *pstate.FlagTextAddr - int64(pstate.HEADR)
	resoff := elfreserve

	var pph *ElfPhdr
	var pnote *ElfPhdr
	if ctxt.LinkMode == LinkExternal {
		/* skip program headers */
		eh.phoff = 0

		eh.phentsize = 0

		if ctxt.BuildMode == BuildModeShared {
			sh := pstate.elfshname(".note.go.pkg-list")
			sh.type_ = SHT_NOTE
			sh = pstate.elfshname(".note.go.abihash")
			sh.type_ = SHT_NOTE
			sh.flags = SHF_ALLOC
			sh = pstate.elfshname(".note.go.deps")
			sh.type_ = SHT_NOTE
		}

		if *pstate.flagBuildid != "" {
			sh := pstate.elfshname(".note.go.buildid")
			sh.type_ = SHT_NOTE
			sh.flags = SHF_ALLOC
		}

		goto elfobj
	}

	/* program header info */
	pph = pstate.newElfPhdr()

	pph.type_ = PT_PHDR
	pph.flags = PF_R
	pph.off = uint64(eh.ehsize)
	pph.vaddr = uint64(*pstate.FlagTextAddr) - uint64(pstate.HEADR) + pph.off
	pph.paddr = uint64(*pstate.FlagTextAddr) - uint64(pstate.HEADR) + pph.off
	pph.align = uint64(*pstate.FlagRound)

	/*
	 * PHDR must be in a loaded segment. Adjust the text
	 * segment boundaries downwards to include it.
	 * Except on NaCl where it must not be loaded.
	 */
	if ctxt.HeadType != objabi.Hnacl {
		o := int64(pstate.Segtext.Vaddr - pph.vaddr)
		pstate.Segtext.Vaddr -= uint64(o)
		pstate.Segtext.Length += uint64(o)
		o = int64(pstate.Segtext.Fileoff - pph.off)
		pstate.Segtext.Fileoff -= uint64(o)
		pstate.Segtext.Filelen += uint64(o)
	}

	if !*pstate.FlagD { /* -d suppresses dynamic loader format */
		/* interpreter */
		sh := pstate.elfshname(".interp")

		sh.type_ = SHT_PROGBITS
		sh.flags = SHF_ALLOC
		sh.addralign = 1
		if pstate.interpreter == "" {
			switch ctxt.HeadType {
			case objabi.Hlinux:
				pstate.interpreter = pstate.thearch.Linuxdynld

			case objabi.Hfreebsd:
				pstate.interpreter = pstate.thearch.Freebsddynld

			case objabi.Hnetbsd:
				pstate.interpreter = pstate.thearch.Netbsddynld

			case objabi.Hopenbsd:
				pstate.interpreter = pstate.thearch.Openbsddynld

			case objabi.Hdragonfly:
				pstate.interpreter = pstate.thearch.Dragonflydynld

			case objabi.Hsolaris:
				pstate.interpreter = pstate.thearch.Solarisdynld
			}
		}

		resoff -= int64(pstate.elfinterp(sh, uint64(startva), uint64(resoff), pstate.interpreter))

		ph := pstate.newElfPhdr()
		ph.type_ = PT_INTERP
		ph.flags = PF_R
		phsh(ph, sh)
	}

	pnote = nil
	if ctxt.HeadType == objabi.Hnetbsd || ctxt.HeadType == objabi.Hopenbsd {
		var sh *ElfShdr
		switch ctxt.HeadType {
		case objabi.Hnetbsd:
			sh = pstate.elfshname(".note.netbsd.ident")
			resoff -= int64(elfnetbsdsig(sh, uint64(startva), uint64(resoff)))

		case objabi.Hopenbsd:
			sh = pstate.elfshname(".note.openbsd.ident")
			resoff -= int64(elfopenbsdsig(sh, uint64(startva), uint64(resoff)))
		}

		pnote = pstate.newElfPhdr()
		pnote.type_ = PT_NOTE
		pnote.flags = PF_R
		phsh(pnote, sh)
	}

	if len(pstate.buildinfo) > 0 {
		sh := pstate.elfshname(".note.gnu.build-id")
		resoff -= int64(pstate.elfbuildinfo(sh, uint64(startva), uint64(resoff)))

		if pnote == nil {
			pnote = pstate.newElfPhdr()
			pnote.type_ = PT_NOTE
			pnote.flags = PF_R
		}

		phsh(pnote, sh)
	}

	if *pstate.flagBuildid != "" {
		sh := pstate.elfshname(".note.go.buildid")
		resoff -= int64(pstate.elfgobuildid(sh, uint64(startva), uint64(resoff)))

		pnote := pstate.newElfPhdr()
		pnote.type_ = PT_NOTE
		pnote.flags = PF_R
		phsh(pnote, sh)
	}

	// Additions to the reserved area must be above this line.

	pstate.elfphload(&pstate.Segtext)
	if len(pstate.Segrodata.Sections) > 0 {
		pstate.elfphload(&pstate.Segrodata)
	}
	if len(pstate.Segrelrodata.Sections) > 0 {
		pstate.elfphload(&pstate.Segrelrodata)
		pstate.elfphrelro(&pstate.Segrelrodata)
	}
	pstate.elfphload(&pstate.Segdata)

	/* Dynamic linking sections */
	if !*pstate.FlagD {
		sh := pstate.elfshname(".dynsym")
		sh.type_ = SHT_DYNSYM
		sh.flags = SHF_ALLOC
		if pstate.elf64 {
			sh.entsize = ELF64SYMSIZE
		} else {
			sh.entsize = ELF32SYMSIZE
		}
		sh.addralign = uint64(ctxt.Arch.RegSize)
		sh.link = uint32(pstate.elfshname(".dynstr").shnum)

		// sh->info = index of first non-local symbol (number of local symbols)
		pstate.shsym(sh, ctxt.Syms.Lookup(".dynsym", 0))

		sh = pstate.elfshname(".dynstr")
		sh.type_ = SHT_STRTAB
		sh.flags = SHF_ALLOC
		sh.addralign = 1
		pstate.shsym(sh, ctxt.Syms.Lookup(".dynstr", 0))

		if pstate.elfverneed != 0 {
			sh := pstate.elfshname(".gnu.version")
			sh.type_ = SHT_GNU_VERSYM
			sh.flags = SHF_ALLOC
			sh.addralign = 2
			sh.link = uint32(pstate.elfshname(".dynsym").shnum)
			sh.entsize = 2
			pstate.shsym(sh, ctxt.Syms.Lookup(".gnu.version", 0))

			sh = pstate.elfshname(".gnu.version_r")
			sh.type_ = SHT_GNU_VERNEED
			sh.flags = SHF_ALLOC
			sh.addralign = uint64(ctxt.Arch.RegSize)
			sh.info = uint32(pstate.elfverneed)
			sh.link = uint32(pstate.elfshname(".dynstr").shnum)
			pstate.shsym(sh, ctxt.Syms.Lookup(".gnu.version_r", 0))
		}

		if pstate.elfRelType == ".rela" {
			sh := pstate.elfshname(".rela.plt")
			sh.type_ = SHT_RELA
			sh.flags = SHF_ALLOC
			sh.entsize = ELF64RELASIZE
			sh.addralign = uint64(ctxt.Arch.RegSize)
			sh.link = uint32(pstate.elfshname(".dynsym").shnum)
			sh.info = uint32(pstate.elfshname(".plt").shnum)
			pstate.shsym(sh, ctxt.Syms.Lookup(".rela.plt", 0))

			sh = pstate.elfshname(".rela")
			sh.type_ = SHT_RELA
			sh.flags = SHF_ALLOC
			sh.entsize = ELF64RELASIZE
			sh.addralign = 8
			sh.link = uint32(pstate.elfshname(".dynsym").shnum)
			pstate.shsym(sh, ctxt.Syms.Lookup(".rela", 0))
		} else {
			sh := pstate.elfshname(".rel.plt")
			sh.type_ = SHT_REL
			sh.flags = SHF_ALLOC
			sh.entsize = ELF32RELSIZE
			sh.addralign = 4
			sh.link = uint32(pstate.elfshname(".dynsym").shnum)
			pstate.shsym(sh, ctxt.Syms.Lookup(".rel.plt", 0))

			sh = pstate.elfshname(".rel")
			sh.type_ = SHT_REL
			sh.flags = SHF_ALLOC
			sh.entsize = ELF32RELSIZE
			sh.addralign = 4
			sh.link = uint32(pstate.elfshname(".dynsym").shnum)
			pstate.shsym(sh, ctxt.Syms.Lookup(".rel", 0))
		}

		if eh.machine == EM_PPC64 {
			sh := pstate.elfshname(".glink")
			sh.type_ = SHT_PROGBITS
			sh.flags = SHF_ALLOC + SHF_EXECINSTR
			sh.addralign = 4
			pstate.shsym(sh, ctxt.Syms.Lookup(".glink", 0))
		}

		sh = pstate.elfshname(".plt")
		sh.type_ = SHT_PROGBITS
		sh.flags = SHF_ALLOC + SHF_EXECINSTR
		if eh.machine == EM_X86_64 {
			sh.entsize = 16
		} else if eh.machine == EM_S390 {
			sh.entsize = 32
		} else if eh.machine == EM_PPC64 {
			// On ppc64, this is just a table of addresses
			// filled by the dynamic linker
			sh.type_ = SHT_NOBITS

			sh.flags = SHF_ALLOC + SHF_WRITE
			sh.entsize = 8
		} else {
			sh.entsize = 4
		}
		sh.addralign = sh.entsize
		pstate.shsym(sh, ctxt.Syms.Lookup(".plt", 0))

		// On ppc64, .got comes from the input files, so don't
		// create it here, and .got.plt is not used.
		if eh.machine != EM_PPC64 {
			sh := pstate.elfshname(".got")
			sh.type_ = SHT_PROGBITS
			sh.flags = SHF_ALLOC + SHF_WRITE
			sh.entsize = uint64(ctxt.Arch.RegSize)
			sh.addralign = uint64(ctxt.Arch.RegSize)
			pstate.shsym(sh, ctxt.Syms.Lookup(".got", 0))

			sh = pstate.elfshname(".got.plt")
			sh.type_ = SHT_PROGBITS
			sh.flags = SHF_ALLOC + SHF_WRITE
			sh.entsize = uint64(ctxt.Arch.RegSize)
			sh.addralign = uint64(ctxt.Arch.RegSize)
			pstate.shsym(sh, ctxt.Syms.Lookup(".got.plt", 0))
		}

		sh = pstate.elfshname(".hash")
		sh.type_ = SHT_HASH
		sh.flags = SHF_ALLOC
		sh.entsize = 4
		sh.addralign = uint64(ctxt.Arch.RegSize)
		sh.link = uint32(pstate.elfshname(".dynsym").shnum)
		pstate.shsym(sh, ctxt.Syms.Lookup(".hash", 0))

		/* sh and PT_DYNAMIC for .dynamic section */
		sh = pstate.elfshname(".dynamic")

		sh.type_ = SHT_DYNAMIC
		sh.flags = SHF_ALLOC + SHF_WRITE
		sh.entsize = 2 * uint64(ctxt.Arch.RegSize)
		sh.addralign = uint64(ctxt.Arch.RegSize)
		sh.link = uint32(pstate.elfshname(".dynstr").shnum)
		pstate.shsym(sh, ctxt.Syms.Lookup(".dynamic", 0))
		ph := pstate.newElfPhdr()
		ph.type_ = PT_DYNAMIC
		ph.flags = PF_R + PF_W
		phsh(ph, sh)

		/*
		 * Thread-local storage segment (really just size).
		 */
		tlssize := uint64(0)
		for _, sect := range pstate.Segdata.Sections {
			if sect.Name == ".tbss" {
				tlssize = sect.Length
			}
		}
		if tlssize != 0 {
			ph := pstate.newElfPhdr()
			ph.type_ = PT_TLS
			ph.flags = PF_R
			ph.memsz = tlssize
			ph.align = uint64(ctxt.Arch.RegSize)
		}
	}

	if ctxt.HeadType == objabi.Hlinux {
		ph := pstate.newElfPhdr()
		ph.type_ = PT_GNU_STACK
		ph.flags = PF_W + PF_R
		ph.align = uint64(ctxt.Arch.RegSize)

		ph = pstate.newElfPhdr()
		ph.type_ = PT_PAX_FLAGS
		ph.flags = 0x2a00 // mprotect, randexec, emutramp disabled
		ph.align = uint64(ctxt.Arch.RegSize)
	} else if ctxt.HeadType == objabi.Hsolaris {
		ph := pstate.newElfPhdr()
		ph.type_ = PT_SUNWSTACK
		ph.flags = PF_W + PF_R
	}

elfobj:
	sh := pstate.elfshname(".shstrtab")
	sh.type_ = SHT_STRTAB
	sh.addralign = 1
	pstate.shsym(sh, ctxt.Syms.Lookup(".shstrtab", 0))
	eh.shstrndx = uint16(sh.shnum)

	// put these sections early in the list
	if !*pstate.FlagS {
		pstate.elfshname(".symtab")
		pstate.elfshname(".strtab")
	}

	for _, sect := range pstate.Segtext.Sections {
		pstate.elfshbits(ctxt.LinkMode, sect)
	}
	for _, sect := range pstate.Segrodata.Sections {
		pstate.elfshbits(ctxt.LinkMode, sect)
	}
	for _, sect := range pstate.Segrelrodata.Sections {
		pstate.elfshbits(ctxt.LinkMode, sect)
	}
	for _, sect := range pstate.Segdata.Sections {
		pstate.elfshbits(ctxt.LinkMode, sect)
	}
	for _, sect := range pstate.Segdwarf.Sections {
		pstate.elfshbits(ctxt.LinkMode, sect)
	}

	if ctxt.LinkMode == LinkExternal {
		for _, sect := range pstate.Segtext.Sections {
			pstate.elfshreloc(ctxt.Arch, sect)
		}
		for _, sect := range pstate.Segrodata.Sections {
			pstate.elfshreloc(ctxt.Arch, sect)
		}
		for _, sect := range pstate.Segrelrodata.Sections {
			pstate.elfshreloc(ctxt.Arch, sect)
		}
		for _, sect := range pstate.Segdata.Sections {
			pstate.elfshreloc(ctxt.Arch, sect)
		}
		for _, s := range pstate.dwarfp {
			if len(s.R) > 0 || s.Type == sym.SDWARFINFO || s.Type == sym.SDWARFLOC {
				pstate.elfshreloc(ctxt.Arch, s.Sect)
			}
		}
		// add a .note.GNU-stack section to mark the stack as non-executable
		sh := pstate.elfshname(".note.GNU-stack")

		sh.type_ = SHT_PROGBITS
		sh.addralign = 1
		sh.flags = 0
	}

	if !*pstate.FlagS {
		sh := pstate.elfshname(".symtab")
		sh.type_ = SHT_SYMTAB
		sh.off = uint64(symo)
		sh.size = uint64(pstate.Symsize)
		sh.addralign = uint64(ctxt.Arch.RegSize)
		sh.entsize = 8 + 2*uint64(ctxt.Arch.RegSize)
		sh.link = uint32(pstate.elfshname(".strtab").shnum)
		sh.info = uint32(pstate.elfglobalsymndx)

		sh = pstate.elfshname(".strtab")
		sh.type_ = SHT_STRTAB
		sh.off = uint64(symo) + uint64(pstate.Symsize)
		sh.size = uint64(len(pstate.Elfstrdat))
		sh.addralign = 1
	}

	/* Main header */
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
	if pstate.elf64 {
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
		eh.entry = uint64(pstate.Entryvalue(ctxt))
	}

	eh.version = EV_CURRENT

	if pph != nil {
		pph.filesz = uint64(eh.phnum) * uint64(eh.phentsize)
		pph.memsz = pph.filesz
	}

	ctxt.Out.SeekSet(pstate, 0)
	a := int64(0)
	a += int64(pstate.elfwritehdr(ctxt.Out))
	a += int64(pstate.elfwritephdrs(ctxt.Out))
	a += int64(pstate.elfwriteshdrs(ctxt.Out))
	if !*pstate.FlagD {
		a += int64(pstate.elfwriteinterp(ctxt.Out))
	}
	if ctxt.LinkMode != LinkExternal {
		if ctxt.HeadType == objabi.Hnetbsd {
			a += int64(pstate.elfwritenetbsdsig(ctxt.Out))
		}
		if ctxt.HeadType == objabi.Hopenbsd {
			a += int64(pstate.elfwriteopenbsdsig(ctxt.Out))
		}
		if len(pstate.buildinfo) > 0 {
			a += int64(pstate.elfwritebuildinfo(ctxt.Out))
		}
		if *pstate.flagBuildid != "" {
			a += int64(pstate.elfwritegobuildid(ctxt.Out))
		}
	}

	if a > elfreserve {
		pstate.Errorf(nil, "ELFRESERVE too small: %d > %d with %d text sections", a, elfreserve, numtext)
	}
}

func (pstate *PackageState) elfadddynsym(ctxt *Link, s *sym.Symbol) {
	if pstate.elf64 {
		s.Dynid = int32(pstate.Nelfsym)
		pstate.Nelfsym++

		d := ctxt.Syms.Lookup(".dynsym", 0)

		name := s.Extname
		d.AddUint32(ctxt.Arch, uint32(pstate.Addstring(ctxt.Syms.Lookup(".dynstr", 0), name)))

		/* type */
		t := STB_GLOBAL << 4

		if s.Attr.CgoExport() && s.Type == sym.STEXT {
			t |= STT_FUNC
		} else {
			t |= STT_OBJECT
		}
		d.AddUint8(uint8(t))

		/* reserved */
		d.AddUint8(0)

		/* section where symbol is defined */
		if s.Type == sym.SDYNIMPORT {
			d.AddUint16(ctxt.Arch, SHN_UNDEF)
		} else {
			d.AddUint16(ctxt.Arch, 1)
		}

		/* value */
		if s.Type == sym.SDYNIMPORT {
			d.AddUint64(ctxt.Arch, 0)
		} else {
			d.AddAddr(ctxt.Arch, s)
		}

		/* size of object */
		d.AddUint64(ctxt.Arch, uint64(s.Size))

		if ctxt.Arch.Family == sys.AMD64 && !s.Attr.CgoExportDynamic() && s.Dynimplib != "" && !pstate.seenlib[s.Dynimplib] {
			pstate.Elfwritedynent(ctxt, ctxt.Syms.Lookup(".dynamic", 0), DT_NEEDED, uint64(pstate.Addstring(ctxt.Syms.Lookup(".dynstr", 0), s.Dynimplib)))
		}
	} else {
		s.Dynid = int32(pstate.Nelfsym)
		pstate.Nelfsym++

		d := ctxt.Syms.Lookup(".dynsym", 0)

		/* name */
		name := s.Extname

		d.AddUint32(ctxt.Arch, uint32(pstate.Addstring(ctxt.Syms.Lookup(".dynstr", 0), name)))

		/* value */
		if s.Type == sym.SDYNIMPORT {
			d.AddUint32(ctxt.Arch, 0)
		} else {
			d.AddAddr(ctxt.Arch, s)
		}

		/* size of object */
		d.AddUint32(ctxt.Arch, uint32(s.Size))

		/* type */
		t := STB_GLOBAL << 4

		// TODO(mwhudson): presumably the behavior should actually be the same on both arm and 386.
		if ctxt.Arch.Family == sys.I386 && s.Attr.CgoExport() && s.Type == sym.STEXT {
			t |= STT_FUNC
		} else if ctxt.Arch.Family == sys.ARM && s.Attr.CgoExportDynamic() && s.Type == sym.STEXT {
			t |= STT_FUNC
		} else {
			t |= STT_OBJECT
		}
		d.AddUint8(uint8(t))
		d.AddUint8(0)

		/* shndx */
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

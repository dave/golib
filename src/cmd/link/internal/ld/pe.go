// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"debug/pe"
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"sort"
	"strconv"
	"strings"
)

type IMAGE_IMPORT_DESCRIPTOR struct {
	OriginalFirstThunk uint32
	TimeDateStamp      uint32
	ForwarderChain     uint32
	Name               uint32
	FirstThunk         uint32
}

type IMAGE_EXPORT_DIRECTORY struct {
	Characteristics       uint32
	TimeDateStamp         uint32
	MajorVersion          uint16
	MinorVersion          uint16
	Name                  uint32
	Base                  uint32
	NumberOfFunctions     uint32
	NumberOfNames         uint32
	AddressOfFunctions    uint32
	AddressOfNames        uint32
	AddressOfNameOrdinals uint32
}

const (
	PEBASE = 0x00400000
)

const (
	IMAGE_FILE_MACHINE_I386              = 0x14c
	IMAGE_FILE_MACHINE_AMD64             = 0x8664
	IMAGE_FILE_RELOCS_STRIPPED           = 0x0001
	IMAGE_FILE_EXECUTABLE_IMAGE          = 0x0002
	IMAGE_FILE_LINE_NUMS_STRIPPED        = 0x0004
	IMAGE_FILE_LARGE_ADDRESS_AWARE       = 0x0020
	IMAGE_FILE_32BIT_MACHINE             = 0x0100
	IMAGE_FILE_DEBUG_STRIPPED            = 0x0200
	IMAGE_SCN_CNT_CODE                   = 0x00000020
	IMAGE_SCN_CNT_INITIALIZED_DATA       = 0x00000040
	IMAGE_SCN_CNT_UNINITIALIZED_DATA     = 0x00000080
	IMAGE_SCN_MEM_EXECUTE                = 0x20000000
	IMAGE_SCN_MEM_READ                   = 0x40000000
	IMAGE_SCN_MEM_WRITE                  = 0x80000000
	IMAGE_SCN_MEM_DISCARDABLE            = 0x2000000
	IMAGE_SCN_LNK_NRELOC_OVFL            = 0x1000000
	IMAGE_SCN_ALIGN_32BYTES              = 0x600000
	IMAGE_DIRECTORY_ENTRY_EXPORT         = 0
	IMAGE_DIRECTORY_ENTRY_IMPORT         = 1
	IMAGE_DIRECTORY_ENTRY_RESOURCE       = 2
	IMAGE_DIRECTORY_ENTRY_EXCEPTION      = 3
	IMAGE_DIRECTORY_ENTRY_SECURITY       = 4
	IMAGE_DIRECTORY_ENTRY_BASERELOC      = 5
	IMAGE_DIRECTORY_ENTRY_DEBUG          = 6
	IMAGE_DIRECTORY_ENTRY_COPYRIGHT      = 7
	IMAGE_DIRECTORY_ENTRY_ARCHITECTURE   = 7
	IMAGE_DIRECTORY_ENTRY_GLOBALPTR      = 8
	IMAGE_DIRECTORY_ENTRY_TLS            = 9
	IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG    = 10
	IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT   = 11
	IMAGE_DIRECTORY_ENTRY_IAT            = 12
	IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT   = 13
	IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR = 14
	IMAGE_SUBSYSTEM_WINDOWS_GUI          = 2
	IMAGE_SUBSYSTEM_WINDOWS_CUI          = 3
)

// TODO(crawshaw): add these constants to debug/pe.
const (
	// TODO: the Microsoft doco says IMAGE_SYM_DTYPE_ARRAY is 3 and IMAGE_SYM_DTYPE_FUNCTION is 2
	IMAGE_SYM_TYPE_NULL      = 0
	IMAGE_SYM_TYPE_STRUCT    = 8
	IMAGE_SYM_DTYPE_FUNCTION = 0x20
	IMAGE_SYM_DTYPE_ARRAY    = 0x30
	IMAGE_SYM_CLASS_EXTERNAL = 2
	IMAGE_SYM_CLASS_STATIC   = 3

	IMAGE_REL_I386_DIR32  = 0x0006
	IMAGE_REL_I386_SECREL = 0x000B
	IMAGE_REL_I386_REL32  = 0x0014

	IMAGE_REL_AMD64_ADDR64 = 0x0001
	IMAGE_REL_AMD64_ADDR32 = 0x0002
	IMAGE_REL_AMD64_REL32  = 0x0004
	IMAGE_REL_AMD64_SECREL = 0x000B
)

type Imp struct {
	s       *sym.Symbol
	off     uint64
	next    *Imp
	argsize int
}

type Dll struct {
	name     string
	nameoff  uint64
	thunkoff uint64
	ms       *Imp
	next     *Dll
}

// peStringTable is a COFF string table.
type peStringTable struct {
	strings    []string
	stringsLen int
}

// size resturns size of string table t.
func (t *peStringTable) size() int {
	// string table starts with 4-byte length at the beginning
	return t.stringsLen + 4
}

// add adds string str to string table t.
func (t *peStringTable) add(str string) int {
	off := t.size()
	t.strings = append(t.strings, str)
	t.stringsLen += len(str) + 1 // each string will have 0 appended to it
	return off
}

// write writes string table t into the output file.
func (t *peStringTable) write(out *OutBuf) {
	out.Write32(uint32(t.size()))
	for _, s := range t.strings {
		out.WriteString(s)
		out.Write8(0)
	}
}

// peSection represents section from COFF section table.
type peSection struct {
	name                 string
	shortName            string
	index                int // one-based index into the Section Table
	virtualSize          uint32
	virtualAddress       uint32
	sizeOfRawData        uint32
	pointerToRawData     uint32
	pointerToRelocations uint32
	numberOfRelocations  uint16
	characteristics      uint32
}

// checkOffset verifies COFF section sect offset in the file.
func (sect *peSection) checkOffset(pstate *PackageState, off int64) {
	if off != int64(sect.pointerToRawData) {
		pstate.Errorf(nil, "%s.PointerToRawData = %#x, want %#x", sect.name, uint64(int64(sect.pointerToRawData)), uint64(off))
		pstate.errorexit()
	}
}

// checkSegment verifies COFF section sect matches address
// and file offset provided in segment seg.
func (sect *peSection) checkSegment(pstate *PackageState, seg *sym.Segment) {
	if seg.Vaddr-PEBASE != uint64(sect.virtualAddress) {
		pstate.Errorf(nil, "%s.VirtualAddress = %#x, want %#x", sect.name, uint64(int64(sect.virtualAddress)), uint64(int64(seg.Vaddr-PEBASE)))
		pstate.errorexit()
	}
	if seg.Fileoff != uint64(sect.pointerToRawData) {
		pstate.Errorf(nil, "%s.PointerToRawData = %#x, want %#x", sect.name, uint64(int64(sect.pointerToRawData)), uint64(int64(seg.Fileoff)))
		pstate.errorexit()
	}
}

// pad adds zeros to the section sect. It writes as many bytes
// as necessary to make section sect.SizeOfRawData bytes long.
// It assumes that n bytes are already written to the file.
func (sect *peSection) pad(pstate *PackageState, out *OutBuf, n uint32) {
	out.WriteStringN(pstate, "", int(sect.sizeOfRawData-n))
}

// write writes COFF section sect into the output file.
func (sect *peSection) write(out *OutBuf, linkmode LinkMode) error {
	h := pe.SectionHeader32{
		VirtualSize:          sect.virtualSize,
		SizeOfRawData:        sect.sizeOfRawData,
		PointerToRawData:     sect.pointerToRawData,
		PointerToRelocations: sect.pointerToRelocations,
		NumberOfRelocations:  sect.numberOfRelocations,
		Characteristics:      sect.characteristics,
	}
	if linkmode != LinkExternal {
		h.VirtualAddress = sect.virtualAddress
	}
	copy(h.Name[:], sect.shortName)
	return binary.Write(out, binary.LittleEndian, h)
}

// emitRelocations emits the relocation entries for the sect.
// The actual relocations are emitted by relocfn.
// This updates the corresponding PE section table entry
// with the relocation offset and count.
func (sect *peSection) emitRelocations(pstate *PackageState, out *OutBuf, relocfn func() int) {
	sect.pointerToRelocations = uint32(out.Offset())
	// first entry: extended relocs
	out.Write32(0) // placeholder for number of relocation + 1
	out.Write32(0)
	out.Write16(0)

	n := relocfn() + 1

	cpos := out.Offset()
	out.SeekSet(pstate, int64(sect.pointerToRelocations))
	out.Write32(uint32(n))
	out.SeekSet(pstate, cpos)
	if n > 0x10000 {
		n = 0x10000
		sect.characteristics |= IMAGE_SCN_LNK_NRELOC_OVFL
	} else {
		sect.pointerToRelocations += 10 // skip the extend reloc entry
	}
	sect.numberOfRelocations = uint16(n - 1)
}

// peFile is used to build COFF file.
type peFile struct {
	sections       []*peSection
	stringTable    peStringTable
	textSect       *peSection
	rdataSect      *peSection
	dataSect       *peSection
	bssSect        *peSection
	ctorsSect      *peSection
	nextSectOffset uint32
	nextFileOffset uint32
	symtabOffset   int64 // offset to the start of symbol table
	symbolCount    int   // number of symbol table records written
	dataDirectory  [16]pe.DataDirectory
}

// addSection adds section to the COFF file f.
func (f *peFile) addSection(pstate *PackageState, name string, sectsize int, filesize int) *peSection {
	sect := &peSection{
		name:             name,
		shortName:        name,
		index:            len(f.sections) + 1,
		virtualSize:      uint32(sectsize),
		virtualAddress:   f.nextSectOffset,
		pointerToRawData: f.nextFileOffset,
	}
	f.nextSectOffset = uint32(Rnd(int64(f.nextSectOffset)+int64(sectsize), pstate.PESECTALIGN))
	if filesize > 0 {
		sect.sizeOfRawData = uint32(Rnd(int64(filesize), pstate.PEFILEALIGN))
		f.nextFileOffset += sect.sizeOfRawData
	}
	f.sections = append(f.sections, sect)
	return sect
}

// addDWARFSection adds DWARF section to the COFF file f.
// This function is similar to addSection, but DWARF section names are
// longer than 8 characters, so they need to be stored in the string table.
func (f *peFile) addDWARFSection(pstate *PackageState, name string, size int) *peSection {
	if size == 0 {
		pstate.Exitf("DWARF section %q is empty", name)
	}
	// DWARF section names are longer than 8 characters.
	// PE format requires such names to be stored in string table,
	// and section names replaced with slash (/) followed by
	// correspondent string table index.
	// see http://www.microsoft.com/whdc/system/platform/firmware/PECOFFdwn.mspx
	// for details
	off := f.stringTable.add(name)
	h := f.addSection(pstate, name, size, size)
	h.shortName = fmt.Sprintf("/%d", off)
	h.characteristics = IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_DISCARDABLE
	return h
}

// addDWARF adds DWARF information to the COFF file f.
func (f *peFile) addDWARF(pstate *PackageState) {
	if *pstate.FlagS { // disable symbol table
		return
	}
	if *pstate.FlagW { // disable dwarf
		return
	}
	for _, sect := range pstate.Segdwarf.Sections {
		h := f.addDWARFSection(pstate, sect.Name, int(sect.Length))
		fileoff := sect.Vaddr - pstate.Segdwarf.Vaddr + pstate.Segdwarf.Fileoff
		if uint64(h.pointerToRawData) != fileoff {
			pstate.Exitf("%s.PointerToRawData = %#x, want %#x", sect.Name, h.pointerToRawData, fileoff)
		}
	}
}

// addInitArray adds .ctors COFF section to the file f.
func (f *peFile) addInitArray(pstate *PackageState, ctxt *Link) *peSection {
	// The size below was determined by the specification for array relocations,
	// and by observing what GCC writes here. If the initarray section grows to
	// contain more than one constructor entry, the size will need to be 8 * constructor_count.
	// However, the entire Go runtime is initialized from just one function, so it is unlikely
	// that this will need to grow in the future.
	var size int
	switch pstate.objabi.GOARCH {
	default:
		pstate.Exitf("peFile.addInitArray: unsupported GOARCH=%q\n", pstate.objabi.GOARCH)
	case "386":
		size = 4
	case "amd64":
		size = 8
	}
	sect := f.addSection(pstate, ".ctors", size, size)
	sect.characteristics = IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ
	sect.sizeOfRawData = uint32(size)
	ctxt.Out.SeekSet(pstate, int64(sect.pointerToRawData))
	sect.checkOffset(pstate, ctxt.Out.Offset())

	init_entry := ctxt.Syms.Lookup(*pstate.flagEntrySymbol, 0)
	addr := uint64(init_entry.Value) - init_entry.Sect.Vaddr
	switch pstate.objabi.GOARCH {
	case "386":
		ctxt.Out.Write32(uint32(addr))
	case "amd64":
		ctxt.Out.Write64(addr)
	}
	return sect
}

// emitRelocations emits relocation entries for go.o in external linking.
func (f *peFile) emitRelocations(pstate *PackageState, ctxt *Link) {
	for ctxt.Out.Offset()&7 != 0 {
		ctxt.Out.Write8(0)
	}

	// relocsect relocates symbols from first in section sect, and returns
	// the total number of relocations emitted.
	relocsect := func(sect *sym.Section, syms []*sym.Symbol, base uint64) int {
		// If main section has no bits, nothing to relocate.
		if sect.Vaddr >= sect.Seg.Vaddr+sect.Seg.Filelen {
			return 0
		}
		relocs := 0
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
		for _, sym := range syms {
			if !sym.Attr.Reachable() {
				continue
			}
			if sym.Value >= int64(eaddr) {
				break
			}
			for ri := range sym.R {
				r := &sym.R[ri]
				if r.Done {
					continue
				}
				if r.Xsym == nil {
					pstate.Errorf(sym, "missing xsym in relocation")
					continue
				}
				if r.Xsym.Dynid < 0 {
					pstate.Errorf(sym, "reloc %d to non-coff symbol %s (outer=%s) %d", r.Type, r.Sym.Name, r.Xsym.Name, r.Sym.Type)
				}
				if !pstate.thearch.PEreloc1(ctxt.Arch, ctxt.Out, sym, r, int64(uint64(sym.Value+int64(r.Off))-base)) {
					pstate.Errorf(sym, "unsupported obj reloc %d/%d to %s", r.Type, r.Siz, r.Sym.Name)
				}
				relocs++
			}
		}
		sect.Rellen = uint64(ctxt.Out.Offset()) - sect.Reloff
		return relocs
	}

	sects := []struct {
		peSect *peSection
		seg    *sym.Segment
		syms   []*sym.Symbol
	}{
		{f.textSect, &pstate.Segtext, ctxt.Textp},
		{f.rdataSect, &pstate.Segrodata, pstate.datap},
		{f.dataSect, &pstate.Segdata, pstate.datap},
	}
	for _, s := range sects {
		s.peSect.emitRelocations(pstate, ctxt.Out, func() int {
			var n int
			for _, sect := range s.seg.Sections {
				n += relocsect(sect, s.syms, s.seg.Vaddr)
			}
			return n
		})
	}

dwarfLoop:
	for _, sect := range pstate.Segdwarf.Sections {
		for _, pesect := range f.sections {
			if sect.Name == pesect.name {
				pesect.emitRelocations(pstate, ctxt.Out, func() int {
					return relocsect(sect, pstate.dwarfp, sect.Vaddr)
				})
				continue dwarfLoop
			}
		}
		pstate.Errorf(nil, "emitRelocations: could not find %q section", sect.Name)
	}

	f.ctorsSect.emitRelocations(pstate, ctxt.Out, func() int {
		dottext := ctxt.Syms.Lookup(".text", 0)
		ctxt.Out.Write32(0)
		ctxt.Out.Write32(uint32(dottext.Dynid))
		switch pstate.objabi.GOARCH {
		default:
			pstate.Errorf(dottext, "unknown architecture for PE: %q\n", pstate.objabi.GOARCH)
		case "386":
			ctxt.Out.Write16(IMAGE_REL_I386_DIR32)
		case "amd64":
			ctxt.Out.Write16(IMAGE_REL_AMD64_ADDR64)
		}
		return 1
	})
}

// writeSymbol appends symbol s to file f symbol table.
// It also sets s.Dynid to written symbol number.
func (f *peFile) writeSymbol(pstate *PackageState, out *OutBuf, s *sym.Symbol, value int64, sectidx int, typ uint16, class uint8) {
	if len(s.Name) > 8 {
		out.Write32(0)
		out.Write32(uint32(f.stringTable.add(s.Name)))
	} else {
		out.WriteStringN(pstate, s.Name, 8)
	}
	out.Write32(uint32(value))
	out.Write16(uint16(sectidx))
	out.Write16(typ)
	out.Write8(class)
	out.Write8(0) // no aux entries

	s.Dynid = int32(f.symbolCount)

	f.symbolCount++
}

// mapToPESection searches peFile f for s symbol's location.
// It returns PE section index, and offset within that section.
func (f *peFile) mapToPESection(pstate *PackageState, s *sym.Symbol, linkmode LinkMode) (pesectidx int, offset int64, err error) {
	if s.Sect == nil {
		return 0, 0, fmt.Errorf("could not map %s symbol with no section", s.Name)
	}
	if s.Sect.Seg == &pstate.Segtext {
		return f.textSect.index, int64(uint64(s.Value) - pstate.Segtext.Vaddr), nil
	}
	if s.Sect.Seg == &pstate.Segrodata {
		return f.rdataSect.index, int64(uint64(s.Value) - pstate.Segrodata.Vaddr), nil
	}
	if s.Sect.Seg != &pstate.Segdata {
		return 0, 0, fmt.Errorf("could not map %s symbol with non .text or .rdata or .data section", s.Name)
	}
	v := uint64(s.Value) - pstate.Segdata.Vaddr
	if linkmode != LinkExternal {
		return f.dataSect.index, int64(v), nil
	}
	if s.Type == sym.SDATA {
		return f.dataSect.index, int64(v), nil
	}
	// Note: although address of runtime.edata (type sym.SDATA) is at the start of .bss section
	// it still belongs to the .data section, not the .bss section.
	if v < pstate.Segdata.Filelen {
		return f.dataSect.index, int64(v), nil
	}
	return f.bssSect.index, int64(v - pstate.Segdata.Filelen), nil
}

// writeSymbols writes all COFF symbol table records.
func (f *peFile) writeSymbols(pstate *PackageState, ctxt *Link) {

	put := func(ctxt *Link, s *sym.Symbol, name string, type_ SymbolType, addr int64, gotype *sym.Symbol) {
		if s == nil {
			return
		}
		if s.Sect == nil && type_ != UndefinedSym {
			return
		}
		switch type_ {
		default:
			return
		case DataSym, BSSSym, TextSym, UndefinedSym:
		}

		// Only windows/386 requires underscore prefix on external symbols.
		if ctxt.Arch.Family == sys.I386 &&
			ctxt.LinkMode == LinkExternal &&
			(s.Type == sym.SHOSTOBJ || s.Attr.CgoExport()) {
			s.Name = "_" + s.Name
		}

		var typ uint16
		if ctxt.LinkMode == LinkExternal {
			typ = IMAGE_SYM_TYPE_NULL
		} else {
			// TODO: fix IMAGE_SYM_DTYPE_ARRAY value and use following expression, instead of 0x0308
			typ = IMAGE_SYM_DTYPE_ARRAY<<8 + IMAGE_SYM_TYPE_STRUCT
			typ = 0x0308 // "array of structs"
		}
		sect, value, err := f.mapToPESection(pstate, s, ctxt.LinkMode)
		if err != nil {
			if type_ == UndefinedSym {
				typ = IMAGE_SYM_DTYPE_FUNCTION
			} else {
				pstate.Errorf(s, "addpesym: %v", err)
			}
		}
		class := IMAGE_SYM_CLASS_EXTERNAL
		if s.Version != 0 || s.Attr.VisibilityHidden() || s.Attr.Local() {
			class = IMAGE_SYM_CLASS_STATIC
		}
		f.writeSymbol(pstate, ctxt.Out, s, value, sect, typ, uint8(class))
	}

	if ctxt.LinkMode == LinkExternal {
		// Include section symbols as external, because
		// .ctors and .debug_* section relocations refer to it.
		for _, pesect := range f.sections {
			sym := ctxt.Syms.Lookup(pesect.name, 0)
			f.writeSymbol(pstate, ctxt.Out, sym, 0, pesect.index, IMAGE_SYM_TYPE_NULL, IMAGE_SYM_CLASS_STATIC)
		}
	}

	pstate.genasmsym(ctxt, put)
}

// writeSymbolTableAndStringTable writes out symbol and string tables for peFile f.
func (f *peFile) writeSymbolTableAndStringTable(pstate *PackageState, ctxt *Link) {
	f.symtabOffset = ctxt.Out.Offset()

	// write COFF symbol table
	if !*pstate.FlagS || ctxt.LinkMode == LinkExternal {
		f.writeSymbols(pstate, ctxt)
	}

	// update COFF file header and section table
	size := f.stringTable.size() + 18*f.symbolCount
	var h *peSection
	if ctxt.LinkMode != LinkExternal {
		// We do not really need .symtab for go.o, and if we have one, ld
		// will also include it in the exe, and that will confuse windows.
		h = f.addSection(pstate, ".symtab", size, size)
		h.characteristics = IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_DISCARDABLE
		h.checkOffset(pstate, f.symtabOffset)
	}

	// write COFF string table
	f.stringTable.write(ctxt.Out)
	if ctxt.LinkMode != LinkExternal {
		h.pad(pstate, ctxt.Out, uint32(size))
	}
}

// writeFileHeader writes COFF file header for peFile f.
func (f *peFile) writeFileHeader(pstate *PackageState, arch *sys.Arch, out *OutBuf, linkmode LinkMode) {
	var fh pe.FileHeader

	switch arch.Family {
	default:
		pstate.Exitf("unknown PE architecture: %v", arch.Family)
	case sys.AMD64:
		fh.Machine = IMAGE_FILE_MACHINE_AMD64
	case sys.I386:
		fh.Machine = IMAGE_FILE_MACHINE_I386
	}

	fh.NumberOfSections = uint16(len(f.sections))

	// Being able to produce identical output for identical input is
	// much more beneficial than having build timestamp in the header.
	fh.TimeDateStamp = 0

	if linkmode == LinkExternal {
		fh.Characteristics = IMAGE_FILE_LINE_NUMS_STRIPPED
	} else {
		fh.Characteristics = IMAGE_FILE_RELOCS_STRIPPED | IMAGE_FILE_EXECUTABLE_IMAGE | IMAGE_FILE_DEBUG_STRIPPED
	}
	if pstate.pe64 != 0 {
		var oh64 pe.OptionalHeader64
		fh.SizeOfOptionalHeader = uint16(binary.Size(&oh64))
		fh.Characteristics |= IMAGE_FILE_LARGE_ADDRESS_AWARE
	} else {
		var oh pe.OptionalHeader32
		fh.SizeOfOptionalHeader = uint16(binary.Size(&oh))
		fh.Characteristics |= IMAGE_FILE_32BIT_MACHINE
	}

	fh.PointerToSymbolTable = uint32(f.symtabOffset)
	fh.NumberOfSymbols = uint32(f.symbolCount)

	binary.Write(out, binary.LittleEndian, &fh)
}

// writeOptionalHeader writes COFF optional header for peFile f.
func (f *peFile) writeOptionalHeader(pstate *PackageState, ctxt *Link) {
	var oh pe.OptionalHeader32
	var oh64 pe.OptionalHeader64

	if pstate.pe64 != 0 {
		oh64.Magic = 0x20b // PE32+
	} else {
		oh.Magic = 0x10b // PE32
		oh.BaseOfData = f.dataSect.virtualAddress
	}

	// Fill out both oh64 and oh. We only use one. Oh well.
	oh64.MajorLinkerVersion = 3
	oh.MajorLinkerVersion = 3
	oh64.MinorLinkerVersion = 0
	oh.MinorLinkerVersion = 0
	oh64.SizeOfCode = f.textSect.sizeOfRawData
	oh.SizeOfCode = f.textSect.sizeOfRawData
	oh64.SizeOfInitializedData = f.dataSect.sizeOfRawData
	oh.SizeOfInitializedData = f.dataSect.sizeOfRawData
	oh64.SizeOfUninitializedData = 0
	oh.SizeOfUninitializedData = 0
	if ctxt.LinkMode != LinkExternal {
		oh64.AddressOfEntryPoint = uint32(pstate.Entryvalue(ctxt) - PEBASE)
		oh.AddressOfEntryPoint = uint32(pstate.Entryvalue(ctxt) - PEBASE)
	}
	oh64.BaseOfCode = f.textSect.virtualAddress
	oh.BaseOfCode = f.textSect.virtualAddress
	oh64.ImageBase = PEBASE
	oh.ImageBase = PEBASE
	oh64.SectionAlignment = uint32(pstate.PESECTALIGN)
	oh.SectionAlignment = uint32(pstate.PESECTALIGN)
	oh64.FileAlignment = uint32(pstate.PEFILEALIGN)
	oh.FileAlignment = uint32(pstate.PEFILEALIGN)
	oh64.MajorOperatingSystemVersion = 4
	oh.MajorOperatingSystemVersion = 4
	oh64.MinorOperatingSystemVersion = 0
	oh.MinorOperatingSystemVersion = 0
	oh64.MajorImageVersion = 1
	oh.MajorImageVersion = 1
	oh64.MinorImageVersion = 0
	oh.MinorImageVersion = 0
	oh64.MajorSubsystemVersion = 4
	oh.MajorSubsystemVersion = 4
	oh64.MinorSubsystemVersion = 0
	oh.MinorSubsystemVersion = 0
	oh64.SizeOfImage = f.nextSectOffset
	oh.SizeOfImage = f.nextSectOffset
	oh64.SizeOfHeaders = uint32(pstate.PEFILEHEADR)
	oh.SizeOfHeaders = uint32(pstate.PEFILEHEADR)
	if pstate.windowsgui {
		oh64.Subsystem = IMAGE_SUBSYSTEM_WINDOWS_GUI
		oh.Subsystem = IMAGE_SUBSYSTEM_WINDOWS_GUI
	} else {
		oh64.Subsystem = IMAGE_SUBSYSTEM_WINDOWS_CUI
		oh.Subsystem = IMAGE_SUBSYSTEM_WINDOWS_CUI
	}

	// Disable stack growth as we don't want Windows to
	// fiddle with the thread stack limits, which we set
	// ourselves to circumvent the stack checks in the
	// Windows exception dispatcher.
	// Commit size must be strictly less than reserve
	// size otherwise reserve will be rounded up to a
	// larger size, as verified with VMMap.

	// On 64-bit, we always reserve 2MB stacks. "Pure" Go code is
	// okay with much smaller stacks, but the syscall package
	// makes it easy to call into arbitrary C code without cgo,
	// and system calls even in "pure" Go code are actually C
	// calls that may need more stack than we think.
	//
	// The default stack reserve size affects only the main
	// thread, ctrlhandler thread, and profileloop thread. For
	// these, it must be greater than the stack size assumed by
	// externalthreadhandler.
	//
	// For other threads we specify stack size in runtime explicitly.
	// For these, the reserve must match STACKSIZE in
	// runtime/cgo/gcc_windows_{386,amd64}.c and the correspondent
	// CreateThread parameter in runtime.newosproc.
	oh64.SizeOfStackReserve = 0x00200000
	if !pstate.iscgo {
		oh64.SizeOfStackCommit = 0x00001000
	} else {
		// TODO(brainman): Maybe remove optional header writing altogether for cgo.
		// For cgo it is the external linker that is building final executable.
		// And it probably does not use any information stored in optional header.
		oh64.SizeOfStackCommit = 0x00200000 - 0x2000 // account for 2 guard pages
	}

	oh.SizeOfStackReserve = 0x00100000
	if !pstate.iscgo {
		oh.SizeOfStackCommit = 0x00001000
	} else {
		oh.SizeOfStackCommit = 0x00100000 - 0x2000 // account for 2 guard pages
	}

	oh64.SizeOfHeapReserve = 0x00100000
	oh.SizeOfHeapReserve = 0x00100000
	oh64.SizeOfHeapCommit = 0x00001000
	oh.SizeOfHeapCommit = 0x00001000
	oh64.NumberOfRvaAndSizes = 16
	oh.NumberOfRvaAndSizes = 16

	if pstate.pe64 != 0 {
		oh64.DataDirectory = f.dataDirectory
	} else {
		oh.DataDirectory = f.dataDirectory
	}

	if pstate.pe64 != 0 {
		binary.Write(ctxt.Out, binary.LittleEndian, &oh64)
	} else {
		binary.Write(ctxt.Out, binary.LittleEndian, &oh)
	}
}

func (pstate *PackageState) Peinit(ctxt *Link) {
	var l int

	switch ctxt.Arch.Family {
	// 64-bit architectures
	case sys.AMD64:
		pstate.pe64 = 1
		var oh64 pe.OptionalHeader64
		l = binary.Size(&oh64)

	// 32-bit architectures
	default:
		var oh pe.OptionalHeader32
		l = binary.Size(&oh)

	}

	if ctxt.LinkMode == LinkExternal {
		// .rdata section will contain "masks" and "shifts" symbols, and they
		// need to be aligned to 16-bytes. So make all sections aligned
		// to 32-byte and mark them all IMAGE_SCN_ALIGN_32BYTES so external
		// linker will honour that requirement.
		pstate.PESECTALIGN = 32
		pstate.PEFILEALIGN = 0
	}

	var sh [16]pe.SectionHeader32
	var fh pe.FileHeader
	pstate.PEFILEHEADR = int32(Rnd(int64(len(pstate.dosstub)+binary.Size(&fh)+l+binary.Size(&sh)), pstate.PEFILEALIGN))
	if ctxt.LinkMode != LinkExternal {
		pstate.PESECTHEADR = int32(Rnd(int64(pstate.PEFILEHEADR), pstate.PESECTALIGN))
	} else {
		pstate.PESECTHEADR = 0
	}
	pstate.pefile.nextSectOffset = uint32(pstate.PESECTHEADR)
	pstate.pefile.nextFileOffset = uint32(pstate.PEFILEHEADR)

	if ctxt.LinkMode == LinkInternal {
		// some mingw libs depend on this symbol, for example, FindPESectionByName
		ctxt.xdefine("__image_base__", sym.SDATA, PEBASE)
		ctxt.xdefine("_image_base__", sym.SDATA, PEBASE)
	}

	pstate.HEADR = pstate.PEFILEHEADR
	if *pstate.FlagTextAddr == -1 {
		*pstate.FlagTextAddr = PEBASE + int64(pstate.PESECTHEADR)
	}
	if *pstate.FlagDataAddr == -1 {
		*pstate.FlagDataAddr = 0
	}
	if *pstate.FlagRound == -1 {
		*pstate.FlagRound = int(pstate.PESECTALIGN)
	}
	if *pstate.FlagDataAddr != 0 && *pstate.FlagRound != 0 {
		fmt.Printf("warning: -D0x%x is ignored because of -R0x%x\n", uint64(*pstate.FlagDataAddr), uint32(*pstate.FlagRound))
	}
}

func (pstate *PackageState) pewrite(ctxt *Link) {
	ctxt.Out.SeekSet(pstate, 0)
	if ctxt.LinkMode != LinkExternal {
		ctxt.Out.Write(pstate.dosstub)
		ctxt.Out.WriteStringN(pstate, "PE", 4)
	}

	pstate.pefile.writeFileHeader(pstate, ctxt.Arch, ctxt.Out, ctxt.LinkMode)

	pstate.pefile.writeOptionalHeader(pstate, ctxt)

	for _, sect := range pstate.pefile.sections {
		sect.write(ctxt.Out, ctxt.LinkMode)
	}
}

func strput(out *OutBuf, s string) {
	out.WriteString(s)
	out.Write8(0)
	// string must be padded to even size
	if (len(s)+1)%2 != 0 {
		out.Write8(0)
	}
}

func (pstate *PackageState) initdynimport(ctxt *Link) *Dll {
	var d *Dll

	pstate.dr = nil
	var m *Imp
	for _, s := range ctxt.Syms.Allsym {
		if !s.Attr.Reachable() || s.Type != sym.SDYNIMPORT {
			continue
		}
		for d = pstate.dr; d != nil; d = d.next {
			if d.name == s.Dynimplib {
				m = new(Imp)
				break
			}
		}

		if d == nil {
			d = new(Dll)
			d.name = s.Dynimplib
			d.next = pstate.dr
			pstate.dr = d
			m = new(Imp)
		}

		// Because external link requires properly stdcall decorated name,
		// all external symbols in runtime use %n to denote that the number
		// of uinptrs this function consumes. Store the argsize and discard
		// the %n suffix if any.
		m.argsize = -1
		if i := strings.IndexByte(s.Extname, '%'); i >= 0 {
			var err error
			m.argsize, err = strconv.Atoi(s.Extname[i+1:])
			if err != nil {
				pstate.Errorf(s, "failed to parse stdcall decoration: %v", err)
			}
			m.argsize *= ctxt.Arch.PtrSize
			s.Extname = s.Extname[:i]
		}

		m.s = s
		m.next = d.ms
		d.ms = m
	}

	if ctxt.LinkMode == LinkExternal {
		// Add real symbol name
		for d := pstate.dr; d != nil; d = d.next {
			for m = d.ms; m != nil; m = m.next {
				m.s.Type = sym.SDATA
				m.s.Grow(int64(ctxt.Arch.PtrSize))
				dynName := m.s.Extname
				// only windows/386 requires stdcall decoration
				if ctxt.Arch.Family == sys.I386 && m.argsize >= 0 {
					dynName += fmt.Sprintf("@%d", m.argsize)
				}
				dynSym := ctxt.Syms.Lookup(dynName, 0)
				dynSym.Attr |= sym.AttrReachable
				dynSym.Type = sym.SHOSTOBJ
				r := m.s.AddRel()
				r.Sym = dynSym
				r.Off = 0
				r.Siz = uint8(ctxt.Arch.PtrSize)
				r.Type = objabi.R_ADDR
			}
		}
	} else {
		dynamic := ctxt.Syms.Lookup(".windynamic", 0)
		dynamic.Attr |= sym.AttrReachable
		dynamic.Type = sym.SWINDOWS
		for d := pstate.dr; d != nil; d = d.next {
			for m = d.ms; m != nil; m = m.next {
				m.s.Type = sym.SWINDOWS
				m.s.Attr |= sym.AttrSubSymbol
				m.s.Sub = dynamic.Sub
				dynamic.Sub = m.s
				m.s.Value = dynamic.Size
				dynamic.Size += int64(ctxt.Arch.PtrSize)
			}

			dynamic.Size += int64(ctxt.Arch.PtrSize)
		}
	}

	return pstate.dr
}

// peimporteddlls returns the gcc command line argument to link all imported
// DLLs.
func (pstate *PackageState) peimporteddlls() []string {
	var dlls []string

	for d := pstate.dr; d != nil; d = d.next {
		dlls = append(dlls, "-l"+strings.TrimSuffix(d.name, ".dll"))
	}

	return dlls
}

func (pstate *PackageState) addimports(ctxt *Link, datsect *peSection) {
	startoff := ctxt.Out.Offset()
	dynamic := ctxt.Syms.Lookup(".windynamic", 0)

	// skip import descriptor table (will write it later)
	n := uint64(0)

	for d := pstate.dr; d != nil; d = d.next {
		n++
	}
	ctxt.Out.SeekSet(pstate, startoff+int64(binary.Size(&IMAGE_IMPORT_DESCRIPTOR{}))*int64(n+1))

	// write dll names
	for d := pstate.dr; d != nil; d = d.next {
		d.nameoff = uint64(ctxt.Out.Offset()) - uint64(startoff)
		strput(ctxt.Out, d.name)
	}

	// write function names
	for d := pstate.dr; d != nil; d = d.next {
		for m := d.ms; m != nil; m = m.next {
			m.off = uint64(pstate.pefile.nextSectOffset) + uint64(ctxt.Out.Offset()) - uint64(startoff)
			ctxt.Out.Write16(0) // hint
			strput(ctxt.Out, m.s.Extname)
		}
	}

	// write OriginalFirstThunks
	oftbase := uint64(ctxt.Out.Offset()) - uint64(startoff)

	n = uint64(ctxt.Out.Offset())
	for d := pstate.dr; d != nil; d = d.next {
		d.thunkoff = uint64(ctxt.Out.Offset()) - n
		for m := d.ms; m != nil; m = m.next {
			if pstate.pe64 != 0 {
				ctxt.Out.Write64(m.off)
			} else {
				ctxt.Out.Write32(uint32(m.off))
			}
		}

		if pstate.pe64 != 0 {
			ctxt.Out.Write64(0)
		} else {
			ctxt.Out.Write32(0)
		}
	}

	// add pe section and pad it at the end
	n = uint64(ctxt.Out.Offset()) - uint64(startoff)

	isect := pstate.pefile.addSection(pstate, ".idata", int(n), int(n))
	isect.characteristics = IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_WRITE
	isect.checkOffset(pstate, startoff)
	isect.pad(pstate, ctxt.Out, uint32(n))
	endoff := ctxt.Out.Offset()

	// write FirstThunks (allocated in .data section)
	ftbase := uint64(dynamic.Value) - uint64(datsect.virtualAddress) - PEBASE

	ctxt.Out.SeekSet(pstate, int64(uint64(datsect.pointerToRawData)+ftbase))
	for d := pstate.dr; d != nil; d = d.next {
		for m := d.ms; m != nil; m = m.next {
			if pstate.pe64 != 0 {
				ctxt.Out.Write64(m.off)
			} else {
				ctxt.Out.Write32(uint32(m.off))
			}
		}

		if pstate.pe64 != 0 {
			ctxt.Out.Write64(0)
		} else {
			ctxt.Out.Write32(0)
		}
	}

	// finally write import descriptor table
	out := ctxt.Out
	out.SeekSet(pstate, startoff)

	for d := pstate.dr; d != nil; d = d.next {
		out.Write32(uint32(uint64(isect.virtualAddress) + oftbase + d.thunkoff))
		out.Write32(0)
		out.Write32(0)
		out.Write32(uint32(uint64(isect.virtualAddress) + d.nameoff))
		out.Write32(uint32(uint64(datsect.virtualAddress) + ftbase + d.thunkoff))
	}

	out.Write32(0) //end
	out.Write32(0)
	out.Write32(0)
	out.Write32(0)
	out.Write32(0)

	// update data directory
	pstate.pefile.dataDirectory[IMAGE_DIRECTORY_ENTRY_IMPORT].VirtualAddress = isect.virtualAddress
	pstate.pefile.dataDirectory[IMAGE_DIRECTORY_ENTRY_IMPORT].Size = isect.virtualSize
	pstate.pefile.dataDirectory[IMAGE_DIRECTORY_ENTRY_IAT].VirtualAddress = uint32(dynamic.Value - PEBASE)
	pstate.pefile.dataDirectory[IMAGE_DIRECTORY_ENTRY_IAT].Size = uint32(dynamic.Size)

	out.SeekSet(pstate, endoff)
}

type byExtname []*sym.Symbol

func (s byExtname) Len() int           { return len(s) }
func (s byExtname) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byExtname) Less(i, j int) bool { return s[i].Extname < s[j].Extname }

func (pstate *PackageState) initdynexport(ctxt *Link) {
	pstate.nexport = 0
	for _, s := range ctxt.Syms.Allsym {
		if !s.Attr.Reachable() || !s.Attr.CgoExportDynamic() {
			continue
		}
		if pstate.nexport+1 > len(pstate.dexport) {
			pstate.Errorf(s, "pe dynexport table is full")
			pstate.errorexit()
		}

		pstate.dexport[pstate.nexport] = s
		pstate.nexport++
	}

	sort.Sort(byExtname(pstate.dexport[:pstate.nexport]))
}

func (pstate *PackageState) addexports(ctxt *Link) {
	var e IMAGE_EXPORT_DIRECTORY

	size := binary.Size(&e) + 10*pstate.nexport + len(*pstate.flagOutfile) + 1
	for i := 0; i < pstate.nexport; i++ {
		size += len(pstate.dexport[i].Extname) + 1
	}

	if pstate.nexport == 0 {
		return
	}

	sect := pstate.pefile.addSection(pstate, ".edata", size, size)
	sect.characteristics = IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ
	sect.checkOffset(pstate, ctxt.Out.Offset())
	va := int(sect.virtualAddress)
	pstate.pefile.dataDirectory[IMAGE_DIRECTORY_ENTRY_EXPORT].VirtualAddress = uint32(va)
	pstate.pefile.dataDirectory[IMAGE_DIRECTORY_ENTRY_EXPORT].Size = sect.virtualSize

	vaName := va + binary.Size(&e) + pstate.nexport*4
	vaAddr := va + binary.Size(&e)
	vaNa := va + binary.Size(&e) + pstate.nexport*8

	e.Characteristics = 0
	e.MajorVersion = 0
	e.MinorVersion = 0
	e.NumberOfFunctions = uint32(pstate.nexport)
	e.NumberOfNames = uint32(pstate.nexport)
	e.Name = uint32(va+binary.Size(&e)) + uint32(pstate.nexport)*10 // Program names.
	e.Base = 1
	e.AddressOfFunctions = uint32(vaAddr)
	e.AddressOfNames = uint32(vaName)
	e.AddressOfNameOrdinals = uint32(vaNa)

	out := ctxt.Out

	// put IMAGE_EXPORT_DIRECTORY
	binary.Write(out, binary.LittleEndian, &e)

	// put EXPORT Address Table
	for i := 0; i < pstate.nexport; i++ {
		out.Write32(uint32(pstate.dexport[i].Value - PEBASE))
	}

	// put EXPORT Name Pointer Table
	v := int(e.Name + uint32(len(*pstate.flagOutfile)) + 1)

	for i := 0; i < pstate.nexport; i++ {
		out.Write32(uint32(v))
		v += len(pstate.dexport[i].Extname) + 1
	}

	// put EXPORT Ordinal Table
	for i := 0; i < pstate.nexport; i++ {
		out.Write16(uint16(i))
	}

	// put Names
	out.WriteStringN(pstate, *pstate.flagOutfile, len(*pstate.flagOutfile)+1)

	for i := 0; i < pstate.nexport; i++ {
		out.WriteStringN(pstate, pstate.dexport[i].Extname, len(pstate.dexport[i].Extname)+1)
	}
	sect.pad(pstate, out, uint32(size))
}

func (ctxt *Link) dope(pstate *PackageState) {
	/* relocation table */
	rel := ctxt.Syms.Lookup(".rel", 0)

	rel.Attr |= sym.AttrReachable
	rel.Type = sym.SELFROSECT

	pstate.initdynimport(ctxt)
	pstate.initdynexport(ctxt)
}

func (pstate *PackageState) setpersrc(ctxt *Link, sym *sym.Symbol) {
	if pstate.rsrcsym != nil {
		pstate.Errorf(sym, "too many .rsrc sections")
	}

	pstate.rsrcsym = sym
}

func (pstate *PackageState) addpersrc(ctxt *Link) {
	if pstate.rsrcsym == nil {
		return
	}

	h := pstate.pefile.addSection(pstate, ".rsrc", int(pstate.rsrcsym.Size), int(pstate.rsrcsym.Size))
	h.characteristics = IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_WRITE | IMAGE_SCN_CNT_INITIALIZED_DATA
	h.checkOffset(pstate, ctxt.Out.Offset())

	// relocation
	for ri := range pstate.rsrcsym.R {
		r := &pstate.rsrcsym.R[ri]
		p := pstate.rsrcsym.P[r.Off:]
		val := uint32(int64(h.virtualAddress) + r.Add)

		// 32-bit little-endian
		p[0] = byte(val)

		p[1] = byte(val >> 8)
		p[2] = byte(val >> 16)
		p[3] = byte(val >> 24)
	}

	ctxt.Out.Write(pstate.rsrcsym.P)
	h.pad(pstate, ctxt.Out, uint32(pstate.rsrcsym.Size))

	// update data directory
	pstate.pefile.dataDirectory[IMAGE_DIRECTORY_ENTRY_RESOURCE].VirtualAddress = h.virtualAddress

	pstate.pefile.dataDirectory[IMAGE_DIRECTORY_ENTRY_RESOURCE].Size = h.virtualSize
}

func (pstate *PackageState) Asmbpe(ctxt *Link) {
	switch ctxt.Arch.Family {
	default:
		pstate.Exitf("unknown PE architecture: %v", ctxt.Arch.Family)
	case sys.AMD64, sys.I386:
	}

	t := pstate.pefile.addSection(pstate, ".text", int(pstate.Segtext.Length), int(pstate.Segtext.Length))
	t.characteristics = IMAGE_SCN_CNT_CODE | IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_EXECUTE | IMAGE_SCN_MEM_READ
	if ctxt.LinkMode == LinkExternal {
		// some data symbols (e.g. masks) end up in the .text section, and they normally
		// expect larger alignment requirement than the default text section alignment.
		t.characteristics |= IMAGE_SCN_ALIGN_32BYTES
	}
	t.checkSegment(pstate, &pstate.Segtext)
	pstate.pefile.textSect = t

	ro := pstate.pefile.addSection(pstate, ".rdata", int(pstate.Segrodata.Length), int(pstate.Segrodata.Length))
	ro.characteristics = IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ
	if ctxt.LinkMode == LinkExternal {
		// some data symbols (e.g. masks) end up in the .rdata section, and they normally
		// expect larger alignment requirement than the default text section alignment.
		ro.characteristics |= IMAGE_SCN_ALIGN_32BYTES
	} else {
		// TODO(brainman): should not need IMAGE_SCN_MEM_EXECUTE, but I do not know why it carshes without it
		ro.characteristics |= IMAGE_SCN_MEM_EXECUTE
	}
	ro.checkSegment(pstate, &pstate.Segrodata)
	pstate.pefile.rdataSect = ro

	var d *peSection
	if ctxt.LinkMode != LinkExternal {
		d = pstate.pefile.addSection(pstate, ".data", int(pstate.Segdata.Length), int(pstate.Segdata.Filelen))
		d.characteristics = IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_WRITE
		d.checkSegment(pstate, &pstate.Segdata)
		pstate.pefile.dataSect = d
	} else {
		d = pstate.pefile.addSection(pstate, ".data", int(pstate.Segdata.Filelen), int(pstate.Segdata.Filelen))
		d.characteristics = IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_WRITE | IMAGE_SCN_ALIGN_32BYTES
		d.checkSegment(pstate, &pstate.Segdata)
		pstate.pefile.dataSect = d

		b := pstate.pefile.addSection(pstate, ".bss", int(pstate.Segdata.Length-pstate.Segdata.Filelen), 0)
		b.characteristics = IMAGE_SCN_CNT_UNINITIALIZED_DATA | IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_WRITE | IMAGE_SCN_ALIGN_32BYTES
		b.pointerToRawData = 0
		pstate.pefile.bssSect = b
	}

	pstate.pefile.addDWARF(pstate)

	if ctxt.LinkMode == LinkExternal {
		pstate.pefile.ctorsSect = pstate.pefile.addInitArray(pstate, ctxt)
	}

	ctxt.Out.SeekSet(pstate, int64(pstate.pefile.nextFileOffset))
	if ctxt.LinkMode != LinkExternal {
		pstate.addimports(ctxt, d)
		pstate.addexports(ctxt)
	}
	pstate.pefile.writeSymbolTableAndStringTable(pstate, ctxt)
	pstate.addpersrc(ctxt)
	if ctxt.LinkMode == LinkExternal {
		pstate.pefile.emitRelocations(pstate, ctxt)
	}

	pstate.pewrite(ctxt)
}

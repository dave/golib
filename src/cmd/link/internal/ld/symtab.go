package ld

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"path/filepath"
	"strings"
)

func (psess *PackageSession) putelfstr(s string) int {
	if len(psess.Elfstrdat) == 0 && s != "" {
		psess.
			putelfstr("")
	}

	off := len(psess.Elfstrdat)
	psess.
		Elfstrdat = append(psess.Elfstrdat, s...)
	psess.
		Elfstrdat = append(psess.Elfstrdat, 0)
	return off
}

func (psess *PackageSession) putelfsyment(out *OutBuf, off int, addr int64, size int64, info int, shndx int, other int) {
	if psess.elf64 {
		out.Write32(uint32(off))
		out.Write8(uint8(info))
		out.Write8(uint8(other))
		out.Write16(uint16(shndx))
		out.Write64(uint64(addr))
		out.Write64(uint64(size))
		psess.
			Symsize += ELF64SYMSIZE
	} else {
		out.Write32(uint32(off))
		out.Write32(uint32(addr))
		out.Write32(uint32(size))
		out.Write8(uint8(info))
		out.Write8(uint8(other))
		out.Write16(uint16(shndx))
		psess.
			Symsize += ELF32SYMSIZE
	}
}

// 0 is reserved

func (psess *PackageSession) putelfsym(ctxt *Link, x *sym.Symbol, s string, t SymbolType, addr int64, go_ *sym.Symbol) {
	var typ int

	switch t {
	default:
		return

	case TextSym:
		typ = STT_FUNC

	case DataSym, BSSSym:
		typ = STT_OBJECT

	case UndefinedSym:

		typ = int(x.ElfType)

	case TLSSym:
		typ = STT_TLS
	}

	size := x.Size
	if t == UndefinedSym {
		size = 0
	}

	xo := x
	for xo.Outer != nil {
		xo = xo.Outer
	}

	var elfshnum int
	if xo.Type == sym.SDYNIMPORT || xo.Type == sym.SHOSTOBJ {
		elfshnum = SHN_UNDEF
	} else {
		if xo.Sect == nil {
			psess.
				Errorf(x, "missing section in putelfsym")
			return
		}
		if xo.Sect.Elfsect == nil {
			psess.
				Errorf(x, "missing ELF section in putelfsym")
			return
		}
		elfshnum = xo.Sect.Elfsect.(*ElfShdr).shnum
	}

	bind := STB_GLOBAL

	if x.Version != 0 || x.Attr.VisibilityHidden() || x.Attr.Local() {
		bind = STB_LOCAL
	}

	if !ctxt.DynlinkingGo() && ctxt.LinkMode == LinkExternal && !x.Attr.CgoExportStatic() && elfshnum != SHN_UNDEF {
		bind = STB_LOCAL
	}

	if ctxt.LinkMode == LinkExternal && elfshnum != SHN_UNDEF {
		addr -= int64(xo.Sect.Vaddr)
	}
	other := STV_DEFAULT
	if x.Attr.VisibilityHidden() {

		other = STV_HIDDEN
	}
	if ctxt.Arch.Family == sys.PPC64 && typ == STT_FUNC && x.Attr.Shared() && x.Name != "runtime.duffzero" && x.Name != "runtime.duffcopy" {

		other |= 3 << 5
	}

	if !ctxt.DynlinkingGo() {

		s = strings.Replace(s, "·", ".", -1)
	}

	if ctxt.DynlinkingGo() && bind == STB_GLOBAL && psess.elfbind == STB_LOCAL && x.Type == sym.STEXT {
		psess.
			putelfsyment(ctxt.Out, psess.putelfstr("local."+s), addr, size, STB_LOCAL<<4|typ&0xf, elfshnum, other)
		x.LocalElfsym = int32(psess.numelfsym)
		psess.
			numelfsym++
		return
	} else if bind != psess.elfbind {
		return
	}
	psess.
		putelfsyment(ctxt.Out, psess.putelfstr(s), addr, size, bind<<4|typ&0xf, elfshnum, other)
	x.Elfsym = int32(psess.numelfsym)
	psess.
		numelfsym++
}

func (psess *PackageSession) putelfsectionsym(out *OutBuf, s *sym.Symbol, shndx int) {
	psess.
		putelfsyment(out, 0, 0, 0, STB_LOCAL<<4|STT_SECTION, shndx, 0)
	s.Elfsym = int32(psess.numelfsym)
	psess.
		numelfsym++
}

func (psess *PackageSession) Asmelfsym(ctxt *Link) {
	psess.
		putelfsyment(ctxt.Out, 0, 0, 0, STB_LOCAL<<4|STT_NOTYPE, 0, 0)
	psess.
		dwarfaddelfsectionsyms(ctxt)
	psess.
		putelfsyment(ctxt.Out, psess.putelfstr("go.go"), 0, 0, STB_LOCAL<<4|STT_FILE, SHN_ABS, 0)
	psess.
		numelfsym++
	psess.
		elfbind = STB_LOCAL
	psess.
		genasmsym(ctxt, psess.putelfsym)
	psess.
		elfbind = STB_GLOBAL
	psess.
		elfglobalsymndx = psess.numelfsym
	psess.
		genasmsym(ctxt, psess.putelfsym)
}

func (psess *PackageSession) putplan9sym(ctxt *Link, x *sym.Symbol, s string, typ SymbolType, addr int64, go_ *sym.Symbol) {
	t := int(typ)
	switch typ {
	case TextSym, DataSym, BSSSym:
		if x.Version != 0 {
			t += 'a' - 'A'
		}
		fallthrough

	case AutoSym, ParamSym, FrameSym:
		l := 4
		if ctxt.HeadType == objabi.Hplan9 && ctxt.Arch.Family == sys.AMD64 && !psess.Flag8 {
			ctxt.Out.Write32b(uint32(addr >> 32))
			l = 8
		}

		ctxt.Out.Write32b(uint32(addr))
		ctxt.Out.Write8(uint8(t + 0x80))

		ctxt.Out.WriteString(s)
		ctxt.Out.Write8(0)
		psess.
			Symsize += int32(l) + 1 + int32(len(s)) + 1

	default:
		return
	}
}

func (psess *PackageSession) Asmplan9sym(ctxt *Link) {
	psess.
		genasmsym(ctxt, psess.putplan9sym)
}

type byPkg []*sym.Library

func (libs byPkg) Len() int {
	return len(libs)
}

func (libs byPkg) Less(a, b int) bool {
	return libs[a].Pkg < libs[b].Pkg
}

func (libs byPkg) Swap(a, b int) {
	libs[a], libs[b] = libs[b], libs[a]
}

func (psess *PackageSession) textsectionmap(ctxt *Link) uint32 {

	t := ctxt.Syms.Lookup("runtime.textsectionmap", 0)
	t.Type = sym.SRODATA
	t.Attr |= sym.AttrReachable
	nsections := int64(0)

	for _, sect := range psess.Segtext.Sections {
		if sect.Name == ".text" {
			nsections++
		} else {
			break
		}
	}
	t.Grow(3 * nsections * int64(ctxt.Arch.PtrSize))

	off := int64(0)
	n := 0

	textbase := psess.Segtext.Sections[0].Vaddr
	for _, sect := range psess.Segtext.Sections {
		if sect.Name != ".text" {
			break
		}
		off = t.SetUint(ctxt.Arch, off, sect.Vaddr-textbase)
		off = t.SetUint(ctxt.Arch, off, sect.Length)
		if n == 0 {
			s := ctxt.Syms.ROLookup("runtime.text", 0)
			if s == nil {
				psess.
					Errorf(nil, "Unable to find symbol runtime.text\n")
			}
			off = t.SetAddr(ctxt.Arch, off, s)

		} else {
			s := ctxt.Syms.Lookup(fmt.Sprintf("runtime.text.%d", n), 0)
			if s == nil {
				psess.
					Errorf(nil, "Unable to find symbol runtime.text.%d\n", n)
			}
			off = t.SetAddr(ctxt.Arch, off, s)
		}
		n++
	}
	return uint32(n)
}

func (ctxt *Link) symtab(psess *PackageSession) {
	psess.
		dosymtype(ctxt)

	ctxt.xdefine("runtime.text", sym.STEXT, 0)

	ctxt.xdefine("runtime.etext", sym.STEXT, 0)
	ctxt.xdefine("runtime.itablink", sym.SRODATA, 0)
	ctxt.xdefine("runtime.eitablink", sym.SRODATA, 0)
	ctxt.xdefine("runtime.rodata", sym.SRODATA, 0)
	ctxt.xdefine("runtime.erodata", sym.SRODATA, 0)
	ctxt.xdefine("runtime.types", sym.SRODATA, 0)
	ctxt.xdefine("runtime.etypes", sym.SRODATA, 0)
	ctxt.xdefine("runtime.noptrdata", sym.SNOPTRDATA, 0)
	ctxt.xdefine("runtime.enoptrdata", sym.SNOPTRDATA, 0)
	ctxt.xdefine("runtime.data", sym.SDATA, 0)
	ctxt.xdefine("runtime.edata", sym.SDATA, 0)
	ctxt.xdefine("runtime.bss", sym.SBSS, 0)
	ctxt.xdefine("runtime.ebss", sym.SBSS, 0)
	ctxt.xdefine("runtime.noptrbss", sym.SNOPTRBSS, 0)
	ctxt.xdefine("runtime.enoptrbss", sym.SNOPTRBSS, 0)
	ctxt.xdefine("runtime.end", sym.SBSS, 0)
	ctxt.xdefine("runtime.epclntab", sym.SRODATA, 0)
	ctxt.xdefine("runtime.esymtab", sym.SRODATA, 0)

	s := ctxt.Syms.Lookup("runtime.gcdata", 0)

	s.Type = sym.SRODATA
	s.Size = 0
	s.Attr |= sym.AttrReachable
	ctxt.xdefine("runtime.egcdata", sym.SRODATA, 0)

	s = ctxt.Syms.Lookup("runtime.gcbss", 0)
	s.Type = sym.SRODATA
	s.Size = 0
	s.Attr |= sym.AttrReachable
	ctxt.xdefine("runtime.egcbss", sym.SRODATA, 0)

	// pseudo-symbols to mark locations of type, string, and go string data.
	var symtype *sym.Symbol
	var symtyperel *sym.Symbol
	if ctxt.UseRelro() && (ctxt.BuildMode == BuildModeCArchive || ctxt.BuildMode == BuildModeCShared || ctxt.BuildMode == BuildModePIE) {
		s = ctxt.Syms.Lookup("type.*", 0)

		s.Type = sym.STYPE
		s.Size = 0
		s.Attr |= sym.AttrReachable
		symtype = s

		s = ctxt.Syms.Lookup("typerel.*", 0)

		s.Type = sym.STYPERELRO
		s.Size = 0
		s.Attr |= sym.AttrReachable
		symtyperel = s
	} else if !ctxt.DynlinkingGo() {
		s = ctxt.Syms.Lookup("type.*", 0)

		s.Type = sym.STYPE
		s.Size = 0
		s.Attr |= sym.AttrReachable
		symtype = s
		symtyperel = s
	}

	groupSym := func(name string, t sym.SymKind) *sym.Symbol {
		s := ctxt.Syms.Lookup(name, 0)
		s.Type = t
		s.Size = 0
		s.Attr |= sym.AttrLocal | sym.AttrReachable
		return s
	}
	var (
		symgostring = groupSym("go.string.*", sym.SGOSTRING)
		symgofunc   = groupSym("go.func.*", sym.SGOFUNC)
		symgcbits   = groupSym("runtime.gcbits.*", sym.SGCBITS)
	)

	var symgofuncrel *sym.Symbol
	if !ctxt.DynlinkingGo() {
		if ctxt.UseRelro() {
			symgofuncrel = groupSym("go.funcrel.*", sym.SGOFUNCRELRO)
		} else {
			symgofuncrel = symgofunc
		}
	}

	symitablink := ctxt.Syms.Lookup("runtime.itablink", 0)
	symitablink.Type = sym.SITABLINK
	psess.
		symt = ctxt.Syms.Lookup("runtime.symtab", 0)
	psess.
		symt.Attr |= sym.AttrLocal
	psess.
		symt.Type = sym.SSYMTAB
	psess.
		symt.Size = 0
	psess.
		symt.Attr |= sym.AttrReachable

	nitablinks := 0

	for _, s := range ctxt.Syms.Allsym {
		if !s.Attr.Reachable() || s.Attr.Special() || s.Type != sym.SRODATA {
			continue
		}

		switch {
		case strings.HasPrefix(s.Name, "type."):
			if !ctxt.DynlinkingGo() {
				s.Attr |= sym.AttrNotInSymbolTable
			}
			if ctxt.UseRelro() {
				s.Type = sym.STYPERELRO
				s.Outer = symtyperel
			} else {
				s.Type = sym.STYPE
				s.Outer = symtype
			}

		case strings.HasPrefix(s.Name, "go.importpath.") && ctxt.UseRelro():

			s.Type = sym.STYPERELRO

		case strings.HasPrefix(s.Name, "go.itablink."):
			nitablinks++
			s.Type = sym.SITABLINK
			s.Attr |= sym.AttrNotInSymbolTable
			s.Outer = symitablink

		case strings.HasPrefix(s.Name, "go.string."):
			s.Type = sym.SGOSTRING
			s.Attr |= sym.AttrNotInSymbolTable
			s.Outer = symgostring

		case strings.HasPrefix(s.Name, "runtime.gcbits."):
			s.Type = sym.SGCBITS
			s.Attr |= sym.AttrNotInSymbolTable
			s.Outer = symgcbits

		case strings.HasSuffix(s.Name, "·f"):
			if !ctxt.DynlinkingGo() {
				s.Attr |= sym.AttrNotInSymbolTable
			}
			if ctxt.UseRelro() {
				s.Type = sym.SGOFUNCRELRO
				s.Outer = symgofuncrel
			} else {
				s.Type = sym.SGOFUNC
				s.Outer = symgofunc
			}

		case strings.HasPrefix(s.Name, "gcargs."),
			strings.HasPrefix(s.Name, "gclocals."),
			strings.HasPrefix(s.Name, "gclocals·"),
			strings.HasPrefix(s.Name, "inltree."):
			s.Type = sym.SGOFUNC
			s.Attr |= sym.AttrNotInSymbolTable
			s.Outer = symgofunc
			s.Align = 4
			psess.
				liveness += (s.Size + int64(s.Align) - 1) &^ (int64(s.Align) - 1)
		}
	}

	if ctxt.BuildMode == BuildModeShared {
		abihashgostr := ctxt.Syms.Lookup("go.link.abihash."+filepath.Base(*psess.flagOutfile), 0)
		abihashgostr.Attr |= sym.AttrReachable
		abihashgostr.Type = sym.SRODATA
		hashsym := ctxt.Syms.Lookup("go.link.abihashbytes", 0)
		abihashgostr.AddAddr(ctxt.Arch, hashsym)
		abihashgostr.AddUint(ctxt.Arch, uint64(hashsym.Size))
	}
	if ctxt.BuildMode == BuildModePlugin || ctxt.Syms.ROLookup("plugin.Open", 0) != nil {
		for _, l := range ctxt.Library {
			s := ctxt.Syms.Lookup("go.link.pkghashbytes."+l.Pkg, 0)
			s.Attr |= sym.AttrReachable
			s.Type = sym.SRODATA
			s.Size = int64(len(l.Hash))
			s.P = []byte(l.Hash)
			str := ctxt.Syms.Lookup("go.link.pkghash."+l.Pkg, 0)
			str.Attr |= sym.AttrReachable
			str.Type = sym.SRODATA
			str.AddAddr(ctxt.Arch, s)
			str.AddUint(ctxt.Arch, uint64(len(l.Hash)))
		}
	}

	nsections := psess.textsectionmap(ctxt)

	moduledata := ctxt.Moduledata

	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.pclntab", 0))
	moduledata.AddUint(ctxt.Arch, uint64(ctxt.Syms.Lookup("runtime.pclntab", 0).Size))
	moduledata.AddUint(ctxt.Arch, uint64(ctxt.Syms.Lookup("runtime.pclntab", 0).Size))

	moduledata.AddAddrPlus(ctxt.Arch, ctxt.Syms.Lookup("runtime.pclntab", 0), int64(psess.pclntabPclntabOffset))
	moduledata.AddUint(ctxt.Arch, uint64(psess.pclntabNfunc+1))
	moduledata.AddUint(ctxt.Arch, uint64(psess.pclntabNfunc+1))

	moduledata.AddAddrPlus(ctxt.Arch, ctxt.Syms.Lookup("runtime.pclntab", 0), int64(psess.pclntabFiletabOffset))
	moduledata.AddUint(ctxt.Arch, uint64(len(ctxt.Filesyms))+1)
	moduledata.AddUint(ctxt.Arch, uint64(len(ctxt.Filesyms))+1)

	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.findfunctab", 0))

	moduledata.AddAddr(ctxt.Arch, psess.pclntabFirstFunc)
	moduledata.AddAddrPlus(ctxt.Arch, psess.pclntabLastFunc, psess.pclntabLastFunc.Size)

	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.text", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.etext", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.noptrdata", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.enoptrdata", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.data", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.edata", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.bss", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.ebss", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.noptrbss", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.enoptrbss", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.end", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.gcdata", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.gcbss", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.types", 0))
	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.etypes", 0))

	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.textsectionmap", 0))
	moduledata.AddUint(ctxt.Arch, uint64(nsections))
	moduledata.AddUint(ctxt.Arch, uint64(nsections))

	typelinkSym := ctxt.Syms.Lookup("runtime.typelink", 0)
	ntypelinks := uint64(typelinkSym.Size) / 4
	moduledata.AddAddr(ctxt.Arch, typelinkSym)
	moduledata.AddUint(ctxt.Arch, ntypelinks)
	moduledata.AddUint(ctxt.Arch, ntypelinks)

	moduledata.AddAddr(ctxt.Arch, ctxt.Syms.Lookup("runtime.itablink", 0))
	moduledata.AddUint(ctxt.Arch, uint64(nitablinks))
	moduledata.AddUint(ctxt.Arch, uint64(nitablinks))

	if ptab := ctxt.Syms.ROLookup("go.plugin.tabs", 0); ptab != nil && ptab.Attr.Reachable() {
		ptab.Attr |= sym.AttrLocal
		ptab.Type = sym.SRODATA

		nentries := uint64(len(ptab.P) / 8)
		moduledata.AddAddr(ctxt.Arch, ptab)
		moduledata.AddUint(ctxt.Arch, nentries)
		moduledata.AddUint(ctxt.Arch, nentries)
	} else {
		moduledata.AddUint(ctxt.Arch, 0)
		moduledata.AddUint(ctxt.Arch, 0)
		moduledata.AddUint(ctxt.Arch, 0)
	}
	if ctxt.BuildMode == BuildModePlugin {
		psess.
			addgostring(ctxt, moduledata, "go.link.thispluginpath", objabi.PathToPrefix(*psess.flagPluginPath))

		pkghashes := ctxt.Syms.Lookup("go.link.pkghashes", 0)
		pkghashes.Attr |= sym.AttrReachable
		pkghashes.Attr |= sym.AttrLocal
		pkghashes.Type = sym.SRODATA

		for i, l := range ctxt.Library {
			psess.
				addgostring(ctxt, pkghashes, fmt.Sprintf("go.link.pkgname.%d", i), l.Pkg)
			psess.
				addgostring(ctxt, pkghashes, fmt.Sprintf("go.link.pkglinkhash.%d", i), l.Hash)

			hash := ctxt.Syms.ROLookup("go.link.pkghash."+l.Pkg, 0)
			pkghashes.AddAddr(ctxt.Arch, hash)
		}
		moduledata.AddAddr(ctxt.Arch, pkghashes)
		moduledata.AddUint(ctxt.Arch, uint64(len(ctxt.Library)))
		moduledata.AddUint(ctxt.Arch, uint64(len(ctxt.Library)))
	} else {
		moduledata.AddUint(ctxt.Arch, 0)
		moduledata.AddUint(ctxt.Arch, 0)
		moduledata.AddUint(ctxt.Arch, 0)
		moduledata.AddUint(ctxt.Arch, 0)
		moduledata.AddUint(ctxt.Arch, 0)
	}
	if len(ctxt.Shlibs) > 0 {
		thismodulename := filepath.Base(*psess.flagOutfile)
		switch ctxt.BuildMode {
		case BuildModeExe, BuildModePIE:

			thismodulename = "the executable"
		}
		psess.
			addgostring(ctxt, moduledata, "go.link.thismodulename", thismodulename)

		modulehashes := ctxt.Syms.Lookup("go.link.abihashes", 0)
		modulehashes.Attr |= sym.AttrReachable
		modulehashes.Attr |= sym.AttrLocal
		modulehashes.Type = sym.SRODATA

		for i, shlib := range ctxt.Shlibs {

			modulename := filepath.Base(shlib.Path)
			psess.
				addgostring(ctxt, modulehashes, fmt.Sprintf("go.link.libname.%d", i), modulename)
			psess.
				addgostring(ctxt, modulehashes, fmt.Sprintf("go.link.linkhash.%d", i), string(shlib.Hash))

			abihash := ctxt.Syms.Lookup("go.link.abihash."+modulename, 0)
			abihash.Attr |= sym.AttrReachable
			modulehashes.AddAddr(ctxt.Arch, abihash)
		}

		moduledata.AddAddr(ctxt.Arch, modulehashes)
		moduledata.AddUint(ctxt.Arch, uint64(len(ctxt.Shlibs)))
		moduledata.AddUint(ctxt.Arch, uint64(len(ctxt.Shlibs)))
	} else {
		moduledata.AddUint(ctxt.Arch, 0)
		moduledata.AddUint(ctxt.Arch, 0)
		moduledata.AddUint(ctxt.Arch, 0)
		moduledata.AddUint(ctxt.Arch, 0)
		moduledata.AddUint(ctxt.Arch, 0)
	}

	hasmain := ctxt.BuildMode == BuildModeExe || ctxt.BuildMode == BuildModePIE
	if hasmain {
		moduledata.AddUint8(1)
	} else {
		moduledata.AddUint8(0)
	}

	moduledatatype := ctxt.Syms.ROLookup("type.runtime.moduledata", 0)
	moduledata.Size = psess.decodetypeSize(ctxt.Arch, moduledatatype)
	moduledata.Grow(moduledata.Size)

	lastmoduledatap := ctxt.Syms.Lookup("runtime.lastmoduledatap", 0)
	if lastmoduledatap.Type != sym.SDYNIMPORT {
		lastmoduledatap.Type = sym.SNOPTRDATA
		lastmoduledatap.Size = 0
		lastmoduledatap.AddAddr(ctxt.Arch, moduledata)
	}
}

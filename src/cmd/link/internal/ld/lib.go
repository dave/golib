package ld

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"debug/elf"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/bio"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/loadelf"
	"github.com/dave/golib/src/cmd/link/internal/loadmacho"
	"github.com/dave/golib/src/cmd/link/internal/loadpe"

	"github.com/dave/golib/src/cmd/link/internal/sym"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

type Arch struct {
	Funcalign        int
	Maxalign         int
	Minalign         int
	Dwarfregsp       int
	Dwarfreglr       int
	Linuxdynld       string
	Freebsddynld     string
	Netbsddynld      string
	Openbsddynld     string
	Dragonflydynld   string
	Solarisdynld     string
	Adddynrel        func(*Link, *sym.Symbol, *sym.Reloc) bool
	Archinit         func(*Link)
	Archreloc        func(*Link, *sym.Reloc, *sym.Symbol, *int64) bool
	Archrelocvariant func(*Link, *sym.Reloc, *sym.Symbol, int64) int64
	Trampoline       func(*Link, *sym.Reloc, *sym.Symbol)
	Asmb             func(*Link)
	Elfreloc1        func(*Link, *sym.Reloc, int64) bool
	Elfsetupplt      func(*Link)
	Gentext          func(*Link)
	Machoreloc1      func(*sys.Arch, *OutBuf, *sym.Symbol, *sym.Reloc, int64) bool
	PEreloc1         func(*sys.Arch, *OutBuf, *sym.Symbol, *sym.Reloc, int64) bool

	// TLSIEtoLE converts a TLS Initial Executable relocation to
	// a TLS Local Executable relocation.
	//
	// This is possible when a TLS IE relocation refers to a local
	// symbol in an executable, which is typical when internally
	// linking PIE binaries.
	TLSIEtoLE func(s *sym.Symbol, off, size int)

	// optional override for assignAddress
	AssignAddress func(ctxt *Link, sect *sym.Section, n int, s *sym.Symbol, va uint64, isTramp bool) (*sym.Section, int, uint64)
}

const (
	MINFUNC = 16 // minimum size for a function
)

// DynlinkingGo returns whether we are producing Go code that can live
// in separate shared libraries linked together at runtime.
func (ctxt *Link) DynlinkingGo() bool {
	if !ctxt.Loaded {
		panic("DynlinkingGo called before all symbols loaded")
	}
	return ctxt.BuildMode == BuildModeShared || ctxt.linkShared || ctxt.BuildMode == BuildModePlugin || ctxt.CanUsePlugins()
}

// CanUsePlugins returns whether a plugins can be used
func (ctxt *Link) CanUsePlugins() bool {
	return ctxt.Syms.ROLookup("plugin.Open", 0) != nil
}

// UseRelro returns whether to make use of "read only relocations" aka
// relro.
func (ctxt *Link) UseRelro() bool {
	switch ctxt.BuildMode {
	case BuildModeCArchive, BuildModeCShared, BuildModeShared, BuildModePIE, BuildModePlugin:
		return ctxt.IsELF
	default:
		return ctxt.linkShared
	}
}

// backup old value of debug['s']

const pkgdef = "__.PKGDEF"

// Set if we see an object compiled by the host compiler that is not
// from a package that is known to support internal linking mode.

func Lflag(ctxt *Link, arg string) {
	ctxt.Libdir = append(ctxt.Libdir, arg)
}

/*
 * Unix doesn't like it when we write to a running (or, sometimes,
 * recently run) binary, so remove the output file before writing it.
 * On Windows 7, remove() can force a subsequent create() to fail.
 * S_ISREG() does not exist on Plan 9.
 */
func (psess *PackageSession) mayberemoveoutfile() {
	if fi, err := os.Lstat(*psess.flagOutfile); err == nil && !fi.Mode().IsRegular() {
		return
	}
	os.Remove(*psess.flagOutfile)
}

func (psess *PackageSession) libinit(ctxt *Link) {
	psess.
		Funcalign = psess.thearch.Funcalign

	suffix := ""

	suffixsep := ""
	if *psess.flagInstallSuffix != "" {
		suffixsep = "_"
		suffix = *psess.flagInstallSuffix
	} else if *psess.flagRace {
		suffixsep = "_"
		suffix = "race"
	} else if *psess.flagMsan {
		suffixsep = "_"
		suffix = "msan"
	}

	Lflag(ctxt, filepath.Join(psess.objabi.GOROOT, "pkg", fmt.Sprintf("%s_%s%s%s", psess.objabi.GOOS, psess.objabi.GOARCH, suffixsep, suffix)))
	psess.
		mayberemoveoutfile()
	f, err := os.OpenFile(*psess.flagOutfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0775)
	if err != nil {
		psess.
			Exitf("cannot create %s: %v", *psess.flagOutfile, err)
	}

	ctxt.Out.w = bufio.NewWriter(f)
	ctxt.Out.f = f

	if *psess.flagEntrySymbol == "" {
		switch ctxt.BuildMode {
		case BuildModeCShared, BuildModeCArchive:
			*psess.flagEntrySymbol = fmt.Sprintf("_rt0_%s_%s_lib", psess.objabi.GOARCH, psess.objabi.GOOS)
		case BuildModeExe, BuildModePIE:
			*psess.flagEntrySymbol = fmt.Sprintf("_rt0_%s_%s", psess.objabi.GOARCH, psess.objabi.GOOS)
		case BuildModeShared, BuildModePlugin:

		default:
			psess.
				Errorf(nil, "unknown *flagEntrySymbol for buildmode %v", ctxt.BuildMode)
		}
	}
}

func (psess *PackageSession) errorexit() {
	if psess.nerrors != 0 {
		psess.
			Exit(2)
	}
	psess.
		Exit(0)
}

func (psess *PackageSession) loadinternal(ctxt *Link, name string) *sym.Library {
	if ctxt.linkShared && ctxt.PackageShlib != nil {
		if shlib := ctxt.PackageShlib[name]; shlib != "" {
			return psess.addlibpath(ctxt, "internal", "internal", "", name, shlib)
		}
	}
	if ctxt.PackageFile != nil {
		if pname := ctxt.PackageFile[name]; pname != "" {
			return psess.addlibpath(ctxt, "internal", "internal", pname, name, "")
		}
		ctxt.Logf("loadinternal: cannot find %s\n", name)
		return nil
	}

	for _, libdir := range ctxt.Libdir {
		if ctxt.linkShared {
			shlibname := filepath.Join(libdir, name+".shlibname")
			if ctxt.Debugvlog != 0 {
				ctxt.Logf("searching for %s.a in %s\n", name, shlibname)
			}
			if _, err := os.Stat(shlibname); err == nil {
				return psess.addlibpath(ctxt, "internal", "internal", "", name, shlibname)
			}
		}
		pname := filepath.Join(libdir, name+".a")
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("searching for %s.a in %s\n", name, pname)
		}
		if _, err := os.Stat(pname); err == nil {
			return psess.addlibpath(ctxt, "internal", "internal", pname, name, "")
		}
	}

	ctxt.Logf("warning: unable to find %s.a\n", name)
	return nil
}

// findLibPathCmd uses cmd command to find gcc library libname.
// It returns library full path if found, or "none" if not found.
func (ctxt *Link) findLibPathCmd(psess *PackageSession, cmd, libname string) string {
	if *psess.flagExtld == "" {
		*psess.flagExtld = "gcc"
	}
	args := hostlinkArchArgs(ctxt.Arch)
	args = append(args, cmd)
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%s %v\n", *psess.flagExtld, args)
	}
	out, err := exec.Command(*psess.flagExtld, args...).Output()
	if err != nil {
		if ctxt.Debugvlog != 0 {
			ctxt.Logf("not using a %s file because compiler failed\n%v\n%s\n", libname, err, out)
		}
		return "none"
	}
	return strings.TrimSpace(string(out))
}

// findLibPath searches for library libname.
// It returns library full path if found, or "none" if not found.
func (ctxt *Link) findLibPath(psess *PackageSession, libname string) string {
	return ctxt.findLibPathCmd(psess, "--print-file-name="+libname, libname)
}

func (ctxt *Link) loadlib(psess *PackageSession) {
	switch ctxt.BuildMode {
	case BuildModeCShared, BuildModePlugin:
		s := ctxt.Syms.Lookup("runtime.islibrary", 0)
		s.Attr |= sym.AttrDuplicateOK
		s.AddUint8(1)
	case BuildModeCArchive:
		s := ctxt.Syms.Lookup("runtime.isarchive", 0)
		s.Attr |= sym.AttrDuplicateOK
		s.AddUint8(1)
	}
	psess.
		loadinternal(ctxt, "runtime")
	if ctxt.Arch.Family == sys.ARM {
		psess.
			loadinternal(ctxt, "math")
	}
	if *psess.flagRace {
		psess.
			loadinternal(ctxt, "runtime/race")
	}
	if *psess.flagMsan {
		psess.
			loadinternal(ctxt, "runtime/msan")
	}

	for i := 0; i < len(ctxt.Library); i++ {
		lib := ctxt.Library[i]
		if lib.Shlib == "" {
			if ctxt.Debugvlog > 1 {
				ctxt.Logf("%5.2f autolib: %s (from %s)\n", psess.Cputime(), lib.File, lib.Objref)
			}
			psess.
				loadobjfile(ctxt, lib)
		}
	}

	for _, lib := range ctxt.Library {
		if lib.Shlib != "" {
			if ctxt.Debugvlog > 1 {
				ctxt.Logf("%5.2f autolib: %s (from %s)\n", psess.Cputime(), lib.Shlib, lib.Objref)
			}
			psess.
				ldshlibsyms(ctxt, lib.Shlib)
		}
	}
	psess.
		iscgo = ctxt.Syms.ROLookup("x_cgo_init", 0) != nil
	psess.
		determineLinkMode(ctxt)

	if ctxt.HeadType == objabi.Hwindows {
		psess.
			Peinit(ctxt)
	}

	if ctxt.HeadType == objabi.Hdarwin && ctxt.LinkMode == LinkExternal {
		*psess.FlagTextAddr = 0
	}

	if ctxt.LinkMode == LinkExternal && ctxt.Arch.Family == sys.PPC64 {
		toc := ctxt.Syms.Lookup(".TOC.", 0)
		toc.Type = sym.SDYNIMPORT
	}

	if ctxt.LinkMode == LinkExternal && !psess.iscgo && ctxt.LibraryByPkg["runtime/cgo"] == nil && !(psess.objabi.GOOS == "darwin" && (ctxt.Arch.Family == sys.AMD64 || ctxt.Arch.Family == sys.I386)) {

		if lib := psess.loadinternal(ctxt, "runtime/cgo"); lib != nil {
			if lib.Shlib != "" {
				psess.
					ldshlibsyms(ctxt, lib.Shlib)
			} else {
				if ctxt.BuildMode == BuildModeShared || ctxt.linkShared {
					psess.
						Exitf("cannot implicitly include runtime/cgo in a shared library")
				}
				psess.
					loadobjfile(ctxt, lib)
			}
		}
	}

	if ctxt.LinkMode == LinkInternal {

		for _, s := range ctxt.Syms.Allsym {
			if s.Type == sym.SHOSTOBJ {

				if s.Extname != "" && s.Dynimplib != "" && !s.Attr.CgoExport() {
					s.Type = sym.SDYNIMPORT
				} else {
					s.Type = 0
				}
			}
		}
	}

	tlsg := ctxt.Syms.Lookup("runtime.tlsg", 0)

	if tlsg.Type == 0 {
		tlsg.Type = sym.STLSBSS
		tlsg.Size = int64(ctxt.Arch.PtrSize)
	} else if tlsg.Type != sym.SDYNIMPORT {
		psess.
			Errorf(nil, "runtime declared tlsg variable %v", tlsg.Type)
	}
	tlsg.Attr |= sym.AttrReachable
	ctxt.Tlsg = tlsg

	var moduledata *sym.Symbol
	if ctxt.BuildMode == BuildModePlugin {
		moduledata = ctxt.Syms.Lookup("local.pluginmoduledata", 0)
		moduledata.Attr |= sym.AttrLocal
	} else {
		moduledata = ctxt.Syms.Lookup("runtime.firstmoduledata", 0)
	}
	if moduledata.Type != 0 && moduledata.Type != sym.SDYNIMPORT {

		moduledata.Size = 0

		if ctxt.Arch.Family == sys.ARM {
			s := ctxt.Syms.Lookup("runtime.goarm", 0)
			s.Type = sym.SDATA
			s.Size = 0
			s.AddUint8(uint8(psess.objabi.GOARM))
		}

		if psess.objabi.Framepointer_enabled(psess.objabi.GOOS, psess.objabi.GOARCH) {
			s := ctxt.Syms.Lookup("runtime.framepointer_enabled", 0)
			s.Type = sym.SDATA
			s.Size = 0
			s.AddUint8(1)
		}
	} else {

		moduledata = ctxt.Syms.Lookup("local.moduledata", 0)
		moduledata.Attr |= sym.AttrLocal
	}

	moduledata.Type = sym.SNOPTRDATA
	moduledata.Attr |= sym.AttrReachable
	ctxt.Moduledata = moduledata

	x := sym.AttrCgoExportDynamic

	if ctxt.LinkMode == LinkExternal {
		x = sym.AttrCgoExportStatic
	}
	w := 0
	for i := range psess.dynexp {
		if psess.dynexp[i].Attr&x != 0 {
			psess.
				dynexp[w] = psess.dynexp[i]
			w++
		}
	}
	psess.
		dynexp = psess.dynexp[:w]

	if ctxt.LinkMode == LinkInternal {
		psess.
			hostobjs(ctxt)

		any := false
		for _, s := range ctxt.Syms.Allsym {
			for _, r := range s.R {
				if r.Sym != nil && r.Sym.Type == sym.SXREF && r.Sym.Name != ".got" {
					any = true
					break
				}
			}
		}
		if any {
			if *psess.flagLibGCC == "" {
				*psess.flagLibGCC = ctxt.findLibPathCmd(psess, "--print-libgcc-file-name", "libgcc")
			}
			if *psess.flagLibGCC != "none" {
				psess.
					hostArchive(ctxt, *psess.flagLibGCC)
			}
			if ctxt.HeadType == objabi.Hwindows {
				if p := ctxt.findLibPath(psess, "libmingwex.a"); p != "none" {
					psess.
						hostArchive(ctxt, p)
				}
				if p := ctxt.findLibPath(psess, "libmingw32.a"); p != "none" {
					psess.
						hostArchive(ctxt, p)
				}

			}
		}
	} else {
		psess.
			hostlinksetup(ctxt)
	}

	ctxt.Loaded = true

	if ctxt.BuildMode == BuildModeExe {
		if psess.havedynamic == 0 && ctxt.HeadType != objabi.Hdarwin && ctxt.HeadType != objabi.Hsolaris {
			*psess.FlagD = true
		}
	}

	if typeSymbolMangling(ctxt) {
		*psess.FlagW = true
		for _, s := range ctxt.Syms.Allsym {
			newName := typeSymbolMangle(s.Name)
			if newName != s.Name {
				ctxt.Syms.Rename(s.Name, newName, int(s.Version))
			}
		}
	}

	if ctxt.BuildMode == BuildModeShared || ctxt.BuildMode == BuildModePlugin || ctxt.CanUsePlugins() {
		for _, lib := range ctxt.Library {
			if lib.Shlib == "" {
				psess.
					genhash(ctxt, lib)
			}
		}
	}

	if ctxt.Arch == psess.sys.Arch386 {
		if (ctxt.BuildMode == BuildModeCArchive && ctxt.IsELF) || (ctxt.BuildMode == BuildModeCShared && ctxt.HeadType != objabi.Hwindows) || ctxt.BuildMode == BuildModePIE || ctxt.DynlinkingGo() {
			got := ctxt.Syms.Lookup("_GLOBAL_OFFSET_TABLE_", 0)
			got.Type = sym.SDYNIMPORT
			got.Attr |= sym.AttrReachable
		}
	}
	psess.
		importcycles()

	ctxt.Library = postorder(ctxt.Library)
	for _, doInternal := range [2]bool{true, false} {
		for _, lib := range ctxt.Library {
			if isRuntimeDepPkg(lib.Pkg) != doInternal {
				continue
			}
			ctxt.Textp = append(ctxt.Textp, lib.Textp...)
			for _, s := range lib.DupTextSyms {
				if !s.Attr.OnList() {
					ctxt.Textp = append(ctxt.Textp, s)
					s.Attr |= sym.AttrOnList

					s.File = objabi.PathToPrefix(lib.Pkg)
				}
			}
		}
	}

	if len(ctxt.Shlibs) > 0 {

		textp := make([]*sym.Symbol, 0, len(ctxt.Textp))
		for _, s := range ctxt.Textp {
			if s.Type != sym.SDYNIMPORT {
				textp = append(textp, s)
			}
		}
		ctxt.Textp = textp
	}
}

// typeSymbolMangling reports whether the linker should shorten the
// names of symbols that represent Go types.
//
// As the names of these symbols are derived from the string of
// the type, they can run to many kilobytes long. So we shorten
// them using a SHA-1 when the name appears in the final binary.
//
// These are the symbols that begin with the prefix 'type.' and
// contain run-time type information used by the runtime and reflect
// packages. All Go binaries contain these symbols, but only only
// those programs loaded dynamically in multiple parts need these
// symbols to have entries in the symbol table.
func typeSymbolMangling(ctxt *Link) bool {
	return ctxt.BuildMode == BuildModeShared || ctxt.linkShared || ctxt.BuildMode == BuildModePlugin || ctxt.Syms.ROLookup("plugin.Open", 0) != nil
}

// typeSymbolMangle mangles the given symbol name into something shorter.
func typeSymbolMangle(name string) string {
	if !strings.HasPrefix(name, "type.") {
		return name
	}
	if strings.HasPrefix(name, "type.runtime.") {
		return name
	}
	if len(name) <= 14 && !strings.Contains(name, "@") {
		return name
	}
	hash := sha1.Sum([]byte(name))
	prefix := "type."
	if name[5] == '.' {
		prefix = "type.."
	}
	return prefix + base64.StdEncoding.EncodeToString(hash[:6])
}

/*
 * look for the next file in an archive.
 * adapted from libmach.
 */
func nextar(bp *bio.Reader, off int64, a *ArHdr) int64 {
	if off&1 != 0 {
		off++
	}
	bp.Seek(off, 0)
	var buf [SAR_HDR]byte
	if n, err := io.ReadFull(bp, buf[:]); err != nil {
		if n == 0 && err != io.EOF {
			return -1
		}
		return 0
	}

	a.name = artrim(buf[0:16])
	a.date = artrim(buf[16:28])
	a.uid = artrim(buf[28:34])
	a.gid = artrim(buf[34:40])
	a.mode = artrim(buf[40:48])
	a.size = artrim(buf[48:58])
	a.fmag = artrim(buf[58:60])

	arsize := atolwhex(a.size)
	if arsize&1 != 0 {
		arsize++
	}
	return arsize + SAR_HDR
}

func (psess *PackageSession) genhash(ctxt *Link, lib *sym.Library) {
	f, err := bio.Open(lib.File)
	if err != nil {
		psess.
			Errorf(nil, "cannot open file %s for hash generation: %v", lib.File, err)
		return
	}
	defer f.Close()

	var magbuf [len(ARMAG)]byte
	if _, err := io.ReadFull(f, magbuf[:]); err != nil {
		psess.
			Exitf("file %s too short", lib.File)
	}

	if string(magbuf[:]) != ARMAG {
		psess.
			Exitf("%s is not an archive file", lib.File)
	}

	var arhdr ArHdr
	l := nextar(f, f.Offset(), &arhdr)
	if l <= 0 {
		psess.
			Errorf(nil, "%s: short read on archive file symbol header", lib.File)
		return
	}
	if arhdr.name != pkgdef {
		psess.
			Errorf(nil, "%s: missing package data entry", lib.File)
		return
	}

	h := sha1.New()

	pkgDefBytes := make([]byte, atolwhex(arhdr.size))
	_, err = io.ReadFull(f, pkgDefBytes)
	if err != nil {
		psess.
			Errorf(nil, "%s: error reading package data: %v", lib.File, err)
		return
	}
	firstEOL := bytes.IndexByte(pkgDefBytes, '\n')
	if firstEOL < 0 {
		psess.
			Errorf(nil, "cannot parse package data of %s for hash generation, no newline found", lib.File)
		return
	}
	firstDoubleDollar := bytes.Index(pkgDefBytes, []byte("\n$$"))
	if firstDoubleDollar < 0 {
		psess.
			Errorf(nil, "cannot parse package data of %s for hash generation, no \\n$$ found", lib.File)
		return
	}
	secondDoubleDollar := bytes.Index(pkgDefBytes[firstDoubleDollar+1:], []byte("\n$$"))
	if secondDoubleDollar < 0 {
		psess.
			Errorf(nil, "cannot parse package data of %s for hash generation, only one \\n$$ found", lib.File)
		return
	}
	h.Write(pkgDefBytes[0:firstEOL])
	h.Write(pkgDefBytes[firstDoubleDollar : firstDoubleDollar+secondDoubleDollar])
	lib.Hash = hex.EncodeToString(h.Sum(nil))
}

func (psess *PackageSession) loadobjfile(ctxt *Link, lib *sym.Library) {
	pkg := objabi.PathToPrefix(lib.Pkg)

	if ctxt.Debugvlog > 1 {
		ctxt.Logf("%5.2f ldobj: %s (%s)\n", psess.Cputime(), lib.File, pkg)
	}
	f, err := bio.Open(lib.File)
	if err != nil {
		psess.
			Exitf("cannot open file %s: %v", lib.File, err)
	}
	defer f.Close()
	defer func() {
		if pkg == "main" && !lib.Main {
			psess.
				Exitf("%s: not package main", lib.File)
		}

		if *psess.flagU && !lib.Safe {
			psess.
				Exitf("%s: load of unsafe package %s", lib.File, pkg)
		}
	}()

	for i := 0; i < len(ARMAG); i++ {
		if c, err := f.ReadByte(); err == nil && c == ARMAG[i] {
			continue
		}

		l := f.Seek(0, 2)
		f.Seek(0, 0)
		psess.
			ldobj(ctxt, f, lib, l, lib.File, lib.File)
		return
	}

	/*
	 * load all the object files from the archive now.
	 * this gives us sequential file access and keeps us
	 * from needing to come back later to pick up more
	 * objects.  it breaks the usual C archive model, but
	 * this is Go, not C.  the common case in Go is that
	 * we need to load all the objects, and then we throw away
	 * the individual symbols that are unused.
	 *
	 * loading every object will also make it possible to
	 * load foreign objects not referenced by __.PKGDEF.
	 */
	var arhdr ArHdr
	off := f.Offset()
	for {
		l := nextar(f, off, &arhdr)
		if l == 0 {
			break
		}
		if l < 0 {
			psess.
				Exitf("%s: malformed archive", lib.File)
		}
		off += l

		if arhdr.name == pkgdef {
			continue
		}

		pname := fmt.Sprintf("%s(%s)", lib.File, arhdr.name)
		l = atolwhex(arhdr.size)
		psess.
			ldobj(ctxt, f, lib, l, pname, lib.File)
	}
}

type Hostobj struct {
	ld     func(*Link, *bio.Reader, string, int64, string)
	pkg    string
	pn     string
	file   string
	off    int64
	length int64
}

// These packages can use internal linking mode.
// Others trigger external mode.

func (psess *PackageSession) ldhostobj(ld func(*Link, *bio.Reader, string, int64, string), headType objabi.HeadType, f *bio.Reader, pkg string, length int64, pn string, file string) *Hostobj {
	isinternal := false
	for _, intpkg := range psess.internalpkg {
		if pkg == intpkg {
			isinternal = true
			break
		}
	}

	if headType == objabi.Hdragonfly {
		if pkg == "net" || pkg == "os/user" {
			isinternal = false
		}
	}

	if !isinternal {
		psess.
			externalobj = true
	}
	psess.
		hostobj = append(psess.hostobj, Hostobj{})
	h := &psess.hostobj[len(psess.hostobj)-1]
	h.ld = ld
	h.pkg = pkg
	h.pn = pn
	h.file = file
	h.off = f.Offset()
	h.length = length
	return h
}

func (psess *PackageSession) hostobjs(ctxt *Link) {
	var h *Hostobj

	for i := 0; i < len(psess.hostobj); i++ {
		h = &psess.hostobj[i]
		f, err := bio.Open(h.file)
		if err != nil {
			psess.
				Exitf("cannot reopen %s: %v", h.pn, err)
		}

		f.Seek(h.off, 0)
		h.ld(ctxt, f, h.pkg, h.length, h.pn)
		f.Close()
	}
}

func (psess *PackageSession) hostlinksetup(ctxt *Link) {
	if ctxt.LinkMode != LinkExternal {
		return
	}
	psess.
		debug_s = *psess.FlagS
	*psess.FlagS = false

	if *psess.flagTmpdir == "" {
		dir, err := ioutil.TempDir("", "go-link-")
		if err != nil {
			log.Fatal(err)
		}
		*psess.flagTmpdir = dir
		psess.
			AtExit(func() {
				ctxt.Out.f.Close()
				os.RemoveAll(*psess.flagTmpdir)
			})
	}

	ctxt.Out.f.Close()
	psess.
		mayberemoveoutfile()

	p := filepath.Join(*psess.flagTmpdir, "go.o")
	var err error
	f, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0775)
	if err != nil {
		psess.
			Exitf("cannot create %s: %v", p, err)
	}

	ctxt.Out.w = bufio.NewWriter(f)
	ctxt.Out.f = f
	ctxt.Out.off = 0
}

// hostobjCopy creates a copy of the object files in hostobj in a
// temporary directory.
func (psess *PackageSession) hostobjCopy() (paths []string) {
	var wg sync.WaitGroup
	sema := make(chan struct{}, runtime.NumCPU())
	for i, h := range psess.hostobj {
		h := h
		dst := filepath.Join(*psess.flagTmpdir, fmt.Sprintf("%06d.o", i))
		paths = append(paths, dst)

		wg.Add(1)
		go func() {
			sema <- struct{}{}
			defer func() {
				<-sema
				wg.Done()
			}()
			f, err := os.Open(h.file)
			if err != nil {
				psess.
					Exitf("cannot reopen %s: %v", h.pn, err)
			}
			if _, err := f.Seek(h.off, 0); err != nil {
				psess.
					Exitf("cannot seek %s: %v", h.pn, err)
			}

			w, err := os.Create(dst)
			if err != nil {
				psess.
					Exitf("cannot create %s: %v", dst, err)
			}
			if _, err := io.CopyN(w, f, h.length); err != nil {
				psess.
					Exitf("cannot write %s: %v", dst, err)
			}
			if err := w.Close(); err != nil {
				psess.
					Exitf("cannot close %s: %v", dst, err)
			}
		}()
	}
	wg.Wait()
	return paths
}

// writeGDBLinkerScript creates gcc linker script file in temp
// directory. writeGDBLinkerScript returns created file path.
// The script is used to work around gcc bug
// (see https://golang.org/issue/20183 for details).
func (psess *PackageSession) writeGDBLinkerScript() string {
	name := "fix_debug_gdb_scripts.ld"
	path := filepath.Join(*psess.flagTmpdir, name)
	src := "SECTIONS\n{\n  .debug_gdb_scripts BLOCK(__section_alignment__) (NOLOAD) :\n  {\n    *(.debug_gdb_scripts)\n  }\n}\nINSERT AFTER .debug_types;\n"

	err := ioutil.WriteFile(path, []byte(src), 0666)
	if err != nil {
		psess.
			Errorf(nil, "WriteFile %s failed: %v", name, err)
	}
	return path
}

// archive builds a .a archive from the hostobj object files.
func (ctxt *Link) archive(psess *PackageSession) {
	if ctxt.BuildMode != BuildModeCArchive {
		return
	}

	if *psess.flagExtar == "" {
		*psess.flagExtar = "ar"
	}
	psess.
		mayberemoveoutfile()

	ctxt.Out.Flush(psess)
	if err := ctxt.Out.f.Close(); err != nil {
		psess.
			Exitf("close: %v", err)
	}
	ctxt.Out.f = nil

	argv := []string{*psess.flagExtar, "-q", "-c", "-s", *psess.flagOutfile}
	argv = append(argv, filepath.Join(*psess.flagTmpdir, "go.o"))
	argv = append(argv, psess.hostobjCopy()...)

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("archive: %s\n", strings.Join(argv, " "))
	}

	if out, err := exec.Command(argv[0], argv[1:]...).CombinedOutput(); err != nil {
		psess.
			Exitf("running %s failed: %v\n%s", argv[0], err, out)
	}
}

func (ctxt *Link) hostlink(psess *PackageSession) {
	if ctxt.LinkMode != LinkExternal || psess.nerrors > 0 {
		return
	}
	if ctxt.BuildMode == BuildModeCArchive {
		return
	}

	if *psess.flagExtld == "" {
		*psess.flagExtld = "gcc"
	}

	var argv []string
	argv = append(argv, *psess.flagExtld)
	argv = append(argv, hostlinkArchArgs(ctxt.Arch)...)

	if *psess.FlagS || psess.debug_s {
		if ctxt.HeadType == objabi.Hdarwin {

		} else {
			argv = append(argv, "-s")
		}
	}

	switch ctxt.HeadType {
	case objabi.Hdarwin:
		argv = append(argv, "-Wl,-headerpad,1144")
		if ctxt.DynlinkingGo() {
			argv = append(argv, "-Wl,-flat_namespace")
		}
		if ctxt.BuildMode == BuildModeExe && !ctxt.Arch.InFamily(sys.ARM64) {
			argv = append(argv, "-Wl,-no_pie")
		}
	case objabi.Hopenbsd:
		argv = append(argv, "-Wl,-nopie")
	case objabi.Hwindows:
		if psess.windowsgui {
			argv = append(argv, "-mwindows")
		} else {
			argv = append(argv, "-mconsole")
		}
	}

	switch ctxt.BuildMode {
	case BuildModeExe:
		if ctxt.HeadType == objabi.Hdarwin {
			if ctxt.Arch.Family == sys.ARM64 {

				argv = append(argv, "-Wl,-pagezero_size,100000000")
			} else {
				argv = append(argv, "-Wl,-pagezero_size,4000000")
			}
		}
	case BuildModePIE:

		if ctxt.HeadType != objabi.Hdarwin {
			if ctxt.UseRelro() {
				argv = append(argv, "-Wl,-z,relro")
			}
			argv = append(argv, "-pie")
		}
	case BuildModeCShared:
		if ctxt.HeadType == objabi.Hdarwin {
			argv = append(argv, "-dynamiclib")
			if ctxt.Arch.Family != sys.AMD64 {
				argv = append(argv, "-Wl,-read_only_relocs,suppress")
			}
		} else {

			argv = append(argv, "-Wl,-Bsymbolic")
			if ctxt.UseRelro() {
				argv = append(argv, "-Wl,-z,relro")
			}
			argv = append(argv, "-shared")
			if ctxt.HeadType != objabi.Hwindows {

				argv = append(argv, "-Wl,-z,nodelete")
			}
		}
	case BuildModeShared:
		if ctxt.UseRelro() {
			argv = append(argv, "-Wl,-z,relro")
		}
		argv = append(argv, "-shared")
	case BuildModePlugin:
		if ctxt.HeadType == objabi.Hdarwin {
			argv = append(argv, "-dynamiclib")
		} else {
			if ctxt.UseRelro() {
				argv = append(argv, "-Wl,-z,relro")
			}
			argv = append(argv, "-shared")
		}
	}

	if ctxt.IsELF && ctxt.DynlinkingGo() {

		argv = append(argv, "-Wl,-znow")

		argv = append(argv, "-Wl,-znocopyreloc")

		if ctxt.Arch.InFamily(sys.ARM, sys.ARM64) {

			argv = append(argv, "-fuse-ld=gold")

			cmd := exec.Command(*psess.flagExtld, "-fuse-ld=gold", "-Wl,--version")
			if out, err := cmd.CombinedOutput(); err == nil {
				if !bytes.Contains(out, []byte("GNU gold")) {
					log.Fatalf("ARM external linker must be gold (issue #15696), but is not: %s", out)
				}
			}
		}
	}

	if ctxt.IsELF && len(psess.buildinfo) > 0 {
		argv = append(argv, fmt.Sprintf("-Wl,--build-id=0x%x", psess.buildinfo))
	}

	outopt := *psess.flagOutfile
	if psess.objabi.GOOS == "windows" && runtime.GOOS == "windows" && filepath.Ext(outopt) == "" {
		outopt += "."
	}
	argv = append(argv, "-o")
	argv = append(argv, outopt)

	if psess.rpath.val != "" {
		argv = append(argv, fmt.Sprintf("-Wl,-rpath,%s", psess.rpath.val))
	}

	if ctxt.IsELF {
		argv = append(argv, "-rdynamic")
	}

	if strings.Contains(argv[0], "clang") {
		argv = append(argv, "-Qunused-arguments")
	}

	const compressDWARF = "-Wl,--compress-debug-sections=zlib-gnu"
	if psess.linkerFlagSupported(argv[0], compressDWARF) {
		argv = append(argv, compressDWARF)
	}

	argv = append(argv, filepath.Join(*psess.flagTmpdir, "go.o"))
	argv = append(argv, psess.hostobjCopy()...)

	if ctxt.linkShared {
		seenDirs := make(map[string]bool)
		seenLibs := make(map[string]bool)
		addshlib := func(path string) {
			dir, base := filepath.Split(path)
			if !seenDirs[dir] {
				argv = append(argv, "-L"+dir)
				if !psess.rpath.set {
					argv = append(argv, "-Wl,-rpath="+dir)
				}
				seenDirs[dir] = true
			}
			base = strings.TrimSuffix(base, ".so")
			base = strings.TrimPrefix(base, "lib")
			if !seenLibs[base] {
				argv = append(argv, "-l"+base)
				seenLibs[base] = true
			}
		}
		for _, shlib := range ctxt.Shlibs {
			addshlib(shlib.Path)
			for _, dep := range shlib.Deps {
				if dep == "" {
					continue
				}
				libpath := psess.findshlib(ctxt, dep)
				if libpath != "" {
					addshlib(libpath)
				}
			}
		}
	}

	argv = append(argv, psess.ldflag...)

	if ctxt.BuildMode == BuildModeExe && !ctxt.linkShared {

		for _, nopie := range []string{"-no-pie", "-nopie"} {
			if psess.linkerFlagSupported(argv[0], nopie) {
				argv = append(argv, nopie)
				break
			}
		}
	}

	for _, p := range strings.Fields(*psess.flagExtldflags) {
		argv = append(argv, p)

		if ctxt.IsELF && p == "-static" {
			for i := range argv {
				if argv[i] == "-rdynamic" {
					argv[i] = "-static"
				}
			}
		}
	}
	if ctxt.HeadType == objabi.Hwindows {

		p := psess.writeGDBLinkerScript()
		argv = append(argv, "-Wl,-T,"+p)

		argv = append(argv, "-Wl,--start-group", "-lmingwex", "-lmingw32", "-Wl,--end-group")
		argv = append(argv, psess.peimporteddlls()...)
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f host link:", psess.Cputime())
		for _, v := range argv {
			ctxt.Logf(" %q", v)
		}
		ctxt.Logf("\n")
	}

	if out, err := exec.Command(argv[0], argv[1:]...).CombinedOutput(); err != nil {
		psess.
			Exitf("running %s failed: %v\n%s", argv[0], err, out)
	} else if len(out) > 0 {

		ctxt.Logf("%s", out)
	}

	if !*psess.FlagS && !*psess.FlagW && !psess.debug_s && ctxt.HeadType == objabi.Hdarwin {
		dsym := filepath.Join(*psess.flagTmpdir, "go.dwarf")
		if out, err := exec.Command("dsymutil", "-f", *psess.flagOutfile, "-o", dsym).CombinedOutput(); err != nil {
			psess.
				Exitf("%s: running dsymutil failed: %v\n%s", os.Args[0], err, out)
		}

		if _, err := os.Stat(dsym); os.IsNotExist(err) {
			return
		}

		combinedOutput := *psess.flagOutfile + "~"
		isIOS, err := psess.machoCombineDwarf(*psess.flagOutfile, dsym, combinedOutput, ctxt.BuildMode)
		if err != nil {
			psess.
				Exitf("%s: combining dwarf failed: %v", os.Args[0], err)
		}
		if !isIOS {
			os.Remove(*psess.flagOutfile)
			if err := os.Rename(combinedOutput, *psess.flagOutfile); err != nil {
				psess.
					Exitf("%s: %v", os.Args[0], err)
			}
		}
	}
}

func (psess *PackageSession) linkerFlagSupported(linker, flag string) bool {
	psess.
		createTrivialCOnce.Do(func() {
		src := filepath.Join(*psess.flagTmpdir, "trivial.c")
		if err := ioutil.WriteFile(src, []byte("int main() { return 0; }"), 0666); err != nil {
			psess.
				Errorf(nil, "WriteFile trivial.c failed: %v", err)
		}
	})

	cmd := exec.Command(linker, flag, "trivial.c")
	cmd.Dir = *psess.flagTmpdir
	cmd.Env = append([]string{"LC_ALL=C"}, os.Environ()...)
	out, err := cmd.CombinedOutput()

	return err == nil && !bytes.Contains(out, []byte("unrecognized")) && !bytes.Contains(out, []byte("unknown"))
}

// hostlinkArchArgs returns arguments to pass to the external linker
// based on the architecture.
func hostlinkArchArgs(arch *sys.Arch) []string {
	switch arch.Family {
	case sys.I386:
		return []string{"-m32"}
	case sys.AMD64, sys.PPC64, sys.S390X:
		return []string{"-m64"}
	case sys.ARM:
		return []string{"-marm"}
	case sys.ARM64:

	case sys.MIPS64:
		return []string{"-mabi=64"}
	case sys.MIPS:
		return []string{"-mabi=32"}
	}
	return nil
}

// ldobj loads an input object. If it is a host object (an object
// compiled by a non-Go compiler) it returns the Hostobj pointer. If
// it is a Go object, it returns nil.
func (psess *PackageSession) ldobj(ctxt *Link, f *bio.Reader, lib *sym.Library, length int64, pn string, file string) *Hostobj {
	pkg := objabi.PathToPrefix(lib.Pkg)

	eof := f.Offset() + length
	start := f.Offset()
	c1 := bgetc(f)
	c2 := bgetc(f)
	c3 := bgetc(f)
	c4 := bgetc(f)
	f.Seek(start, 0)

	magic := uint32(c1)<<24 | uint32(c2)<<16 | uint32(c3)<<8 | uint32(c4)
	if magic == 0x7f454c46 {
		ldelf := func(ctxt *Link, f *bio.Reader, pkg string, length int64, pn string) {
			textp, flags, err := loadelf.Load(ctxt.Arch, ctxt.Syms, f, pkg, length, pn, psess.ehdr.flags)
			if err != nil {
				psess.
					Errorf(nil, "%v", err)
				return
			}
			psess.
				ehdr.flags = flags
			ctxt.Textp = append(ctxt.Textp, textp...)
		}
		return psess.ldhostobj(ldelf, ctxt.HeadType, f, pkg, length, pn, file)
	}

	if magic&^1 == 0xfeedface || magic&^0x01000000 == 0xcefaedfe {
		ldmacho := func(ctxt *Link, f *bio.Reader, pkg string, length int64, pn string) {
			textp, err := loadmacho.Load(ctxt.Arch, ctxt.Syms, f, pkg, length, pn)
			if err != nil {
				psess.
					Errorf(nil, "%v", err)
				return
			}
			ctxt.Textp = append(ctxt.Textp, textp...)
		}
		return psess.ldhostobj(ldmacho, ctxt.HeadType, f, pkg, length, pn, file)
	}

	if c1 == 0x4c && c2 == 0x01 || c1 == 0x64 && c2 == 0x86 {
		ldpe := func(ctxt *Link, f *bio.Reader, pkg string, length int64, pn string) {
			textp, rsrc, err := loadpe.Load(ctxt.Arch, ctxt.Syms, f, pkg, length, pn)
			if err != nil {
				psess.
					Errorf(nil, "%v", err)
				return
			}
			if rsrc != nil {
				psess.
					setpersrc(ctxt, rsrc)
			}
			ctxt.Textp = append(ctxt.Textp, textp...)
		}
		return psess.ldhostobj(ldpe, ctxt.HeadType, f, pkg, length, pn, file)
	}

	line, err := f.ReadString('\n')
	if err != nil {
		psess.
			Errorf(nil, "truncated object file: %s: %v", pn, err)
		return nil
	}

	if !strings.HasPrefix(line, "go object ") {
		if strings.HasSuffix(pn, ".go") {
			psess.
				Exitf("%s: uncompiled .go source file", pn)
			return nil
		}

		if line == ctxt.Arch.Name {
			psess.
				Errorf(nil, "%s: stale object file", pn)
			return nil
		}
		psess.
			Errorf(nil, "%s: not an object file", pn)
		return nil
	}

	t := fmt.Sprintf("%s %s %s ", psess.objabi.GOOS, psess.objabi.GOARCH, psess.objabi.Version)

	line = strings.TrimRight(line, "\n")
	if !strings.HasPrefix(line[10:]+" ", t) && !*psess.flagF {
		psess.
			Errorf(nil, "%s: object is [%s] expected [%s]", pn, line[10:], t)
		return nil
	}

	if len(line) >= len(t)+10 {
		if psess.theline == "" {
			psess.
				theline = line[10:]
		} else if psess.theline != line[10:] {
			psess.
				Errorf(nil, "%s: object is [%s] expected [%s]", pn, line[10:], psess.theline)
			return nil
		}
	}

	import0 := f.Offset()

	c1 = '\n'
	c2 = bgetc(f)
	c3 = bgetc(f)
	markers := 0
	for {
		if c1 == '\n' {
			if markers%2 == 0 && c2 == '!' && c3 == '\n' {
				break
			}
			if c2 == '$' && c3 == '$' {
				markers++
			}
		}

		c1 = c2
		c2 = c3
		c3 = bgetc(f)
		if c3 == -1 {
			psess.
				Errorf(nil, "truncated object file: %s", pn)
			return nil
		}
	}

	import1 := f.Offset()

	f.Seek(import0, 0)
	psess.
		ldpkg(ctxt, f, lib, import1-import0-2, pn)
	f.Seek(import1, 0)
	psess.objfile.
		Load(ctxt.Arch, ctxt.Syms, f, lib, eof-f.Offset(), pn)
	psess.
		addImports(ctxt, lib, pn)
	return nil
}

func (psess *PackageSession) readelfsymboldata(ctxt *Link, f *elf.File, sym *elf.Symbol) []byte {
	data := make([]byte, sym.Size)
	sect := f.Sections[sym.Section]
	if sect.Type != elf.SHT_PROGBITS && sect.Type != elf.SHT_NOTE {
		psess.
			Errorf(nil, "reading %s from non-data section", sym.Name)
	}
	n, err := sect.ReadAt(data, int64(sym.Value-sect.Addr))
	if uint64(n) != sym.Size {
		psess.
			Errorf(nil, "reading contents of %s: %v", sym.Name, err)
	}
	return data
}

func readwithpad(r io.Reader, sz int32) ([]byte, error) {
	data := make([]byte, Rnd(int64(sz), 4))
	_, err := io.ReadFull(r, data)
	if err != nil {
		return nil, err
	}
	data = data[:sz]
	return data, nil
}

func readnote(f *elf.File, name []byte, typ int32) ([]byte, error) {
	for _, sect := range f.Sections {
		if sect.Type != elf.SHT_NOTE {
			continue
		}
		r := sect.Open()
		for {
			var namesize, descsize, noteType int32
			err := binary.Read(r, f.ByteOrder, &namesize)
			if err != nil {
				if err == io.EOF {
					break
				}
				return nil, fmt.Errorf("read namesize failed: %v", err)
			}
			err = binary.Read(r, f.ByteOrder, &descsize)
			if err != nil {
				return nil, fmt.Errorf("read descsize failed: %v", err)
			}
			err = binary.Read(r, f.ByteOrder, &noteType)
			if err != nil {
				return nil, fmt.Errorf("read type failed: %v", err)
			}
			noteName, err := readwithpad(r, namesize)
			if err != nil {
				return nil, fmt.Errorf("read name failed: %v", err)
			}
			desc, err := readwithpad(r, descsize)
			if err != nil {
				return nil, fmt.Errorf("read desc failed: %v", err)
			}
			if string(name) == string(noteName) && typ == noteType {
				return desc, nil
			}
		}
	}
	return nil, nil
}

func (psess *PackageSession) findshlib(ctxt *Link, shlib string) string {
	if filepath.IsAbs(shlib) {
		return shlib
	}
	for _, libdir := range ctxt.Libdir {
		libpath := filepath.Join(libdir, shlib)
		if _, err := os.Stat(libpath); err == nil {
			return libpath
		}
	}
	psess.
		Errorf(nil, "cannot find shared library: %s", shlib)
	return ""
}

func (psess *PackageSession) ldshlibsyms(ctxt *Link, shlib string) {
	var libpath string
	if filepath.IsAbs(shlib) {
		libpath = shlib
		shlib = filepath.Base(shlib)
	} else {
		libpath = psess.findshlib(ctxt, shlib)
		if libpath == "" {
			return
		}
	}
	for _, processedlib := range ctxt.Shlibs {
		if processedlib.Path == libpath {
			return
		}
	}
	if ctxt.Debugvlog > 1 {
		ctxt.Logf("%5.2f ldshlibsyms: found library with name %s at %s\n", psess.Cputime(), shlib, libpath)
	}

	f, err := elf.Open(libpath)
	if err != nil {
		psess.
			Errorf(nil, "cannot open shared library: %s", libpath)
		return
	}
	defer f.Close()

	hash, err := readnote(f, psess.ELF_NOTE_GO_NAME, ELF_NOTE_GOABIHASH_TAG)
	if err != nil {
		psess.
			Errorf(nil, "cannot read ABI hash from shared library %s: %v", libpath, err)
		return
	}

	depsbytes, err := readnote(f, psess.ELF_NOTE_GO_NAME, ELF_NOTE_GODEPS_TAG)
	if err != nil {
		psess.
			Errorf(nil, "cannot read dep list from shared library %s: %v", libpath, err)
		return
	}
	var deps []string
	for _, dep := range strings.Split(string(depsbytes), "\n") {
		if dep == "" {
			continue
		}
		if !filepath.IsAbs(dep) {

			abs := filepath.Join(filepath.Dir(libpath), dep)
			if _, err := os.Stat(abs); err == nil {
				dep = abs
			}
		}
		deps = append(deps, dep)
	}

	syms, err := f.DynamicSymbols()
	if err != nil {
		psess.
			Errorf(nil, "cannot read symbols from shared library: %s", libpath)
		return
	}
	gcdataLocations := make(map[uint64]*sym.Symbol)
	for _, elfsym := range syms {
		if elf.ST_TYPE(elfsym.Info) == elf.STT_NOTYPE || elf.ST_TYPE(elfsym.Info) == elf.STT_SECTION {
			continue
		}
		lsym := ctxt.Syms.Lookup(elfsym.Name, 0)

		if lsym.Type != 0 && lsym.Type != sym.SDYNIMPORT {
			continue
		}
		lsym.Type = sym.SDYNIMPORT
		lsym.ElfType = elf.ST_TYPE(elfsym.Info)
		lsym.Size = int64(elfsym.Size)
		if elfsym.Section != elf.SHN_UNDEF {

			lsym.File = libpath

			if strings.HasPrefix(lsym.Name, "type.") && !strings.HasPrefix(lsym.Name, "type..") {
				lsym.P = psess.readelfsymboldata(ctxt, f, &elfsym)
				gcdataLocations[elfsym.Value+2*uint64(ctxt.Arch.PtrSize)+8+1*uint64(ctxt.Arch.PtrSize)] = lsym
			}
		}
	}
	gcdataAddresses := make(map[*sym.Symbol]uint64)
	if ctxt.Arch.Family == sys.ARM64 {
		for _, sect := range f.Sections {
			if sect.Type == elf.SHT_RELA {
				var rela elf.Rela64
				rdr := sect.Open()
				for {
					err := binary.Read(rdr, f.ByteOrder, &rela)
					if err == io.EOF {
						break
					} else if err != nil {
						psess.
							Errorf(nil, "reading relocation failed %v", err)
						return
					}
					t := elf.R_AARCH64(rela.Info & 0xffff)
					if t != elf.R_AARCH64_RELATIVE {
						continue
					}
					if lsym, ok := gcdataLocations[rela.Off]; ok {
						gcdataAddresses[lsym] = uint64(rela.Addend)
					}
				}
			}
		}
	}

	ctxt.Shlibs = append(ctxt.Shlibs, Shlib{Path: libpath, Hash: hash, Deps: deps, File: f, gcdataAddresses: gcdataAddresses})
}

func addsection(arch *sys.Arch, seg *sym.Segment, name string, rwx int) *sym.Section {
	sect := new(sym.Section)
	sect.Rwx = uint8(rwx)
	sect.Name = name
	sect.Seg = seg
	sect.Align = int32(arch.PtrSize)
	seg.Sections = append(seg.Sections, sect)
	return sect
}

func Le16(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

func Le32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func Le64(b []byte) uint64 {
	return uint64(Le32(b)) | uint64(Le32(b[4:]))<<32
}

func Be16(b []byte) uint16 {
	return uint16(b[0])<<8 | uint16(b[1])
}

func Be32(b []byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

type chain struct {
	sym   *sym.Symbol
	up    *chain
	limit int // limit on entry to sym
}

func haslinkregister(ctxt *Link) bool {
	return ctxt.FixedFrameSize() != 0
}

func callsize(ctxt *Link) int {
	if haslinkregister(ctxt) {
		return 0
	}
	return ctxt.Arch.RegSize
}

func (ctxt *Link) dostkcheck(psess *PackageSession) {
	var ch chain
	psess.
		morestack = ctxt.Syms.Lookup("runtime.morestack", 0)

	ch.up = nil

	ch.limit = objabi.StackLimit - callsize(ctxt)

	for _, s := range ctxt.Textp {

		if s.Name == "runtime.racesymbolizethunk" {
			continue
		}

		if s.Attr.NoSplit() {
			ch.sym = s
			psess.
				stkcheck(ctxt, &ch, 0)
		}
	}

	for _, s := range ctxt.Textp {
		if !s.Attr.NoSplit() {
			ch.sym = s
			psess.
				stkcheck(ctxt, &ch, 0)
		}
	}
}

func (psess *PackageSession) stkcheck(ctxt *Link, up *chain, depth int) int {
	limit := up.limit
	s := up.sym

	top := limit == objabi.StackLimit-callsize(ctxt)
	if top {
		if s.Attr.StackCheck() {
			return 0
		}
		s.Attr |= sym.AttrStackCheck
	}

	if depth > 100 {
		psess.
			Errorf(s, "nosplit stack check too deep")
		psess.
			stkbroke(ctxt, up, 0)
		return -1
	}

	if s.Attr.External() || s.FuncInfo == nil {

		if depth == 1 && s.Type != sym.SXREF && !ctxt.DynlinkingGo() &&
			ctxt.BuildMode != BuildModeCArchive && ctxt.BuildMode != BuildModePIE && ctxt.BuildMode != BuildModeCShared && ctxt.BuildMode != BuildModePlugin {

		}
		return -1
	}

	if limit < 0 {
		psess.
			stkbroke(ctxt, up, limit)
		return -1
	}

	if s == psess.morestack {
		return 0
	}

	var ch chain
	ch.up = up

	if !s.Attr.NoSplit() {

		ch.limit = limit - callsize(ctxt)
		ch.sym = psess.morestack
		if psess.stkcheck(ctxt, &ch, depth+1) < 0 {
			return -1
		}
		if !top {
			return 0
		}

		locals := int32(0)
		if s.FuncInfo != nil {
			locals = s.FuncInfo.Locals
		}
		limit = int(objabi.StackLimit+locals) + int(ctxt.FixedFrameSize())
	}

	ri := 0

	endr := len(s.R)
	var ch1 chain
	var pcsp Pciter
	var r *sym.Reloc
	for pciterinit(ctxt, &pcsp, &s.FuncInfo.Pcsp); pcsp.done == 0; pciternext(&pcsp) {

		if int32(limit)-pcsp.value < 0 {
			psess.
				stkbroke(ctxt, up, int(int32(limit)-pcsp.value))
			return -1
		}

		for ; ri < endr && uint32(s.R[ri].Off) < pcsp.nextpc; ri++ {
			r = &s.R[ri]
			switch r.Type {

			case objabi.R_CALL, objabi.R_CALLARM, objabi.R_CALLARM64, objabi.R_CALLPOWER, objabi.R_CALLMIPS:
				ch.limit = int(int32(limit) - pcsp.value - int32(callsize(ctxt)))
				ch.sym = r.Sym
				if psess.stkcheck(ctxt, &ch, depth+1) < 0 {
					return -1
				}

			case objabi.R_CALLIND:
				ch.limit = int(int32(limit) - pcsp.value - int32(callsize(ctxt)))

				ch.sym = nil
				ch1.limit = ch.limit - callsize(ctxt)
				ch1.up = &ch
				ch1.sym = psess.morestack
				if psess.stkcheck(ctxt, &ch1, depth+2) < 0 {
					return -1
				}
			}
		}
	}

	return 0
}

func (psess *PackageSession) stkbroke(ctxt *Link, ch *chain, limit int) {
	psess.
		Errorf(ch.sym, "nosplit stack overflow")
	stkprint(ctxt, ch, limit)
}

func stkprint(ctxt *Link, ch *chain, limit int) {
	var name string

	if ch.sym != nil {
		name = ch.sym.Name
		if ch.sym.Attr.NoSplit() {
			name += " (nosplit)"
		}
	} else {
		name = "function pointer"
	}

	if ch.up == nil {

		if ch.sym.Attr.NoSplit() {
			fmt.Printf("\t%d\tassumed on entry to %s\n", ch.limit, name)
		} else {
			fmt.Printf("\t%d\tguaranteed after split check in %s\n", ch.limit, name)
		}
	} else {
		stkprint(ctxt, ch.up, ch.limit+callsize(ctxt))
		if !haslinkregister(ctxt) {
			fmt.Printf("\t%d\ton entry to %s\n", ch.limit, name)
		}
	}

	if ch.limit != limit {
		fmt.Printf("\t%d\tafter %s uses %d\n", limit, name, ch.limit-limit)
	}
}

func (psess *PackageSession) usage() {
	fmt.Fprintf(os.Stderr, "usage: link [options] main.o\n")
	objabi.Flagprint(os.Stderr)
	psess.
		Exit(2)
}

type SymbolType int8

const (
	// see also https://9p.io/magic/man2html/1/nm
	TextSym      SymbolType = 'T'
	DataSym      SymbolType = 'D'
	BSSSym       SymbolType = 'B'
	UndefinedSym SymbolType = 'U'
	TLSSym       SymbolType = 't'
	FrameSym     SymbolType = 'm'
	ParamSym     SymbolType = 'p'
	AutoSym      SymbolType = 'a'

	// Deleted auto (not a real sym, just placeholder for type)
	DeletedAutoSym = 'x'
)

func (psess *PackageSession) genasmsym(ctxt *Link, put func(*Link, *sym.Symbol, string, SymbolType, int64, *sym.Symbol)) {

	s := ctxt.Syms.Lookup("runtime.text", 0)
	if s.Type == sym.STEXT {

		if !(ctxt.DynlinkingGo() && ctxt.HeadType == objabi.Hdarwin) {
			put(ctxt, s, s.Name, TextSym, s.Value, nil)
		}
	}

	n := 0

	for _, sect := range psess.Segtext.Sections {
		if n == 0 {
			n++
			continue
		}
		if sect.Name != ".text" {
			break
		}
		s = ctxt.Syms.ROLookup(fmt.Sprintf("runtime.text.%d", n), 0)
		if s == nil {
			break
		}
		if s.Type == sym.STEXT {
			put(ctxt, s, s.Name, TextSym, s.Value, nil)
		}
		n++
	}

	s = ctxt.Syms.Lookup("runtime.etext", 0)
	if s.Type == sym.STEXT {

		if !(ctxt.DynlinkingGo() && ctxt.HeadType == objabi.Hdarwin) {
			put(ctxt, s, s.Name, TextSym, s.Value, nil)
		}
	}

	for _, s := range ctxt.Syms.Allsym {
		if s.Attr.NotInSymbolTable() {
			continue
		}
		if (s.Name == "" || s.Name[0] == '.') && s.Version == 0 && s.Name != ".rathole" && s.Name != ".TOC." {
			continue
		}
		switch s.Type {
		case sym.SCONST,
			sym.SRODATA,
			sym.SSYMTAB,
			sym.SPCLNTAB,
			sym.SINITARR,
			sym.SDATA,
			sym.SNOPTRDATA,
			sym.SELFROSECT,
			sym.SMACHOGOT,
			sym.STYPE,
			sym.SSTRING,
			sym.SGOSTRING,
			sym.SGOFUNC,
			sym.SGCBITS,
			sym.STYPERELRO,
			sym.SSTRINGRELRO,
			sym.SGOSTRINGRELRO,
			sym.SGOFUNCRELRO,
			sym.SGCBITSRELRO,
			sym.SRODATARELRO,
			sym.STYPELINK,
			sym.SITABLINK,
			sym.SWINDOWS:
			if !s.Attr.Reachable() {
				continue
			}
			put(ctxt, s, s.Name, DataSym, psess.Symaddr(s), s.Gotype)

		case sym.SBSS, sym.SNOPTRBSS:
			if !s.Attr.Reachable() {
				continue
			}
			if len(s.P) > 0 {
				psess.
					Errorf(s, "should not be bss (size=%d type=%v special=%v)", len(s.P), s.Type, s.Attr.Special())
			}
			put(ctxt, s, s.Name, BSSSym, psess.Symaddr(s), s.Gotype)

		case sym.SHOSTOBJ:
			if ctxt.HeadType == objabi.Hwindows || ctxt.IsELF {
				put(ctxt, s, s.Name, UndefinedSym, s.Value, nil)
			}

		case sym.SDYNIMPORT:
			if !s.Attr.Reachable() {
				continue
			}
			put(ctxt, s, s.Extname, UndefinedSym, 0, nil)

		case sym.STLSBSS:
			if ctxt.LinkMode == LinkExternal {
				put(ctxt, s, s.Name, TLSSym, psess.Symaddr(s), s.Gotype)
			}
		}
	}

	var off int32
	for _, s := range ctxt.Textp {
		put(ctxt, s, s.Name, TextSym, s.Value, s.Gotype)

		locals := int32(0)
		if s.FuncInfo != nil {
			locals = s.FuncInfo.Locals
		}

		put(ctxt, nil, ".frame", FrameSym, int64(locals)+int64(ctxt.Arch.PtrSize), nil)

		if s.FuncInfo == nil {
			continue
		}
		for _, a := range s.FuncInfo.Autom {
			if a.Name == objabi.A_DELETED_AUTO {
				put(ctxt, nil, "", DeletedAutoSym, 0, a.Gotype)
				continue
			}

			if a.Name != objabi.A_AUTO && a.Name != objabi.A_PARAM {
				continue
			}

			if a.Name == objabi.A_PARAM {
				off = a.Aoffset
			} else {
				off = a.Aoffset - int32(ctxt.Arch.PtrSize)
			}

			if off >= 0 {
				put(ctxt, nil, a.Asym.Name, ParamSym, int64(off), a.Gotype)
				continue
			}

			if off <= int32(-ctxt.Arch.PtrSize) {
				put(ctxt, nil, a.Asym.Name, AutoSym, -(int64(off) + int64(ctxt.Arch.PtrSize)), a.Gotype)
				continue
			}

		}
	}

	if ctxt.Debugvlog != 0 || *psess.flagN {
		ctxt.Logf("%5.2f symsize = %d\n", psess.Cputime(), uint32(psess.Symsize))
	}
}

func (psess *PackageSession) Symaddr(s *sym.Symbol) int64 {
	if !s.Attr.Reachable() {
		psess.
			Errorf(s, "unreachable symbol in symaddr")
	}
	return s.Value
}

func (ctxt *Link) xdefine(p string, t sym.SymKind, v int64) {
	s := ctxt.Syms.Lookup(p, 0)
	s.Type = t
	s.Value = v
	s.Attr |= sym.AttrReachable
	s.Attr |= sym.AttrSpecial
	s.Attr |= sym.AttrLocal
}

func (psess *PackageSession) datoff(s *sym.Symbol, addr int64) int64 {
	if uint64(addr) >= psess.Segdata.Vaddr {
		return int64(uint64(addr) - psess.Segdata.Vaddr + psess.Segdata.Fileoff)
	}
	if uint64(addr) >= psess.Segtext.Vaddr {
		return int64(uint64(addr) - psess.Segtext.Vaddr + psess.Segtext.Fileoff)
	}
	psess.
		Errorf(s, "invalid datoff %#x", addr)
	return 0
}

func (psess *PackageSession) Entryvalue(ctxt *Link) int64 {
	a := *psess.flagEntrySymbol
	if a[0] >= '0' && a[0] <= '9' {
		return atolwhex(a)
	}
	s := ctxt.Syms.Lookup(a, 0)
	if s.Type == 0 {
		return *psess.FlagTextAddr
	}
	if s.Type != sym.STEXT {
		psess.
			Errorf(s, "entry not text")
	}
	return s.Value
}

func (psess *PackageSession) undefsym(ctxt *Link, s *sym.Symbol) {
	var r *sym.Reloc

	for i := 0; i < len(s.R); i++ {
		r = &s.R[i]
		if r.Sym == nil {
			continue
		}

		if (r.Sym.Type == sym.Sxxx || r.Sym.Type == sym.SXREF) && !r.Sym.Attr.VisibilityHidden() {
			psess.
				Errorf(s, "undefined: %q", r.Sym.Name)
		}
		if !r.Sym.Attr.Reachable() && r.Type != objabi.R_WEAKADDROFF {
			psess.
				Errorf(s, "relocation target %q", r.Sym.Name)
		}
	}
}

func (ctxt *Link) undef(psess *PackageSession) {

	if psess.nerrors > 0 {
		return
	}

	for _, s := range ctxt.Textp {
		psess.
			undefsym(ctxt, s)
	}
	for _, s := range psess.datap {
		psess.
			undefsym(ctxt, s)
	}
	if psess.nerrors > 0 {
		psess.
			errorexit()
	}
}

func (ctxt *Link) callgraph(psess *PackageSession) {
	if !*psess.FlagC {
		return
	}

	var i int
	var r *sym.Reloc
	for _, s := range ctxt.Textp {
		for i = 0; i < len(s.R); i++ {
			r = &s.R[i]
			if r.Sym == nil {
				continue
			}
			if (r.Type == objabi.R_CALL || r.Type == objabi.R_CALLARM || r.Type == objabi.R_CALLPOWER || r.Type == objabi.R_CALLMIPS) && r.Sym.Type == sym.STEXT {
				ctxt.Logf("%s calls %s\n", s.Name, r.Sym.Name)
			}
		}
	}
}

func Rnd(v int64, r int64) int64 {
	if r <= 0 {
		return v
	}
	v += r - 1
	c := v % r
	if c < 0 {
		c += r
	}
	v -= c
	return v
}

func bgetc(r *bio.Reader) int {
	c, err := r.ReadByte()
	if err != nil {
		if err != io.EOF {
			log.Fatalf("reading input: %v", err)
		}
		return -1
	}
	return int(c)
}

type markKind uint8 // for postorder traversal
const (
	_ markKind = iota
	visiting
	visited
)

func postorder(libs []*sym.Library) []*sym.Library {
	order := make([]*sym.Library, 0, len(libs))
	mark := make(map[*sym.Library]markKind, len(libs))
	for _, lib := range libs {
		dfs(lib, mark, &order)
	}
	return order
}

func dfs(lib *sym.Library, mark map[*sym.Library]markKind, order *[]*sym.Library) {
	if mark[lib] == visited {
		return
	}
	if mark[lib] == visiting {
		panic("found import cycle while visiting " + lib.Pkg)
	}
	mark[lib] = visiting
	for _, i := range lib.Imports {
		dfs(i, mark, order)
	}
	mark[lib] = visited
	*order = append(*order, lib)
}

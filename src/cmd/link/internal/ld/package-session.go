package ld

import (
	"flag"
	"github.com/dave/golib/src/cmd/internal/bio"
	"github.com/dave/golib/src/cmd/internal/dwarf"
	"github.com/dave/golib/src/cmd/internal/gcprog"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/loadelf"
	"github.com/dave/golib/src/cmd/link/internal/loadmacho"
	"github.com/dave/golib/src/cmd/link/internal/loadpe"
	"github.com/dave/golib/src/cmd/link/internal/objfile"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"time"
)

type PackageSession struct {
	bio       *bio.PackageSession
	dwarf     *dwarf.PackageSession
	gcprog    *gcprog.PackageSession
	loadelf   *loadelf.PackageSession
	loadmacho *loadmacho.PackageSession
	loadpe    *loadpe.PackageSession
	obj       *obj.PackageSession
	objabi    *objabi.PackageSession
	objfile   *objfile.PackageSession
	src       *src.PackageSession
	sym       *sym.PackageSession
	sys       *sys.PackageSession

	ELF_NOTE_BUILDINFO_NAME []byte

	ELF_NOTE_GO_NAME     []byte
	ELF_NOTE_NETBSD_NAME []byte

	ELF_NOTE_OPENBSD_NAME []byte
	Elfstrdat             []byte
	Flag8                 bool
	FlagC                 *bool
	FlagD                 *bool

	FlagDataAddr   *int64
	FlagDebugTramp *int

	FlagRound *int
	FlagS     *bool

	FlagTextAddr *int64
	FlagW        *bool

	Funcalign int

	HEADR  int32
	Lcsize int32

	Nelfsym     int
	PEFILEALIGN int64

	PEFILEHEADR int32
	PESECTALIGN int64

	PESECTHEADR  int32
	Segdata      sym.Segment
	Segdwarf     sym.Segment
	Segrelrodata sym.Segment
	Segrodata    sym.Segment
	Segtext      sym.Segment
	Spsize       int32
	Symsize      int32
	atExitFuncs  []func()

	buildinfo  []byte
	cpuprofile *string

	createTrivialCOnce sync.Once
	datap              []*sym.Symbol
	debug_s            bool

	dexport [1024]*sym.Symbol
	dosstub []uint8

	dr        *Dll
	dwarfaddr int64

	dwarfp                []*sym.Symbol
	dwarfstart, linkstart int64

	dwglobals dwarf.DWDie
	dwroot    dwarf.DWDie

	dwtypes dwarf.DWDie

	dylib  []string
	dynexp []*sym.Symbol
	dynlib []string

	ehdr  ElfEhdr
	elf64 bool

	elfRelType string
	elfbind    int

	elfglobalsymndx int

	elfstr [100]Elfstring

	elfverneed  int
	externalobj bool
	flagA       *bool
	flagBuildid *string

	flagDumpDep *bool

	flagEntrySymbol *string
	flagExtar       *string
	flagExtld       *string
	flagExtldflags  *string

	flagF          *bool
	flagFieldTrack *string

	flagG             *bool
	flagH             *bool
	flagInstallSuffix *string

	flagInterpreter *string
	flagLibGCC      *string
	flagMsan        *bool

	flagN          *bool
	flagOutfile    *string
	flagPluginPath *string

	flagRace *bool

	flagTmpdir *string

	flagU *bool

	gdbscript string

	havedynamic int

	hostobj []Hostobj

	internalpkg []string
	interp      string
	interpreter string
	iscgo       bool
	ldflag      []string

	linkoff    int64
	linkoffset uint32

	liveness int64

	load []MachoLoad

	loadBudget     int
	machohdr       MachoHdr
	memprofile     *string
	memprofilerate *int64

	morestack *sym.Symbol
	ndebug    int

	nelfstr int
	nerrors int

	nexport int
	nkind   [3]int
	nsect   int
	nseg    int

	nsortsym  int
	numelfsym int

	pclntabFiletabOffset int32

	pclntabFirstFunc *sym.Symbol
	pclntabLastFunc  *sym.Symbol
	pclntabNfunc     int32

	pclntabPclntabOffset int32
	pclntabZpcln         sym.FuncInfo

	pe64 int

	pefile         peFile
	phdr           [400]*ElfPhdr
	pkgall         []*Pkg
	pkglistfornote []byte

	prefixBuf []byte

	prototypedies      map[string]*dwarf.DWDie
	realdwarf, linkseg *macho.Segment

	rpath Rpath

	rsrcsym *sym.Symbol
	seenlib map[string]bool
	seg     [16]MachoSeg

	shdr      [400]*ElfShdr
	sortsym   []*sym.Symbol
	start     time.Time
	startTime time.Time

	strdata  map[string]string
	strnames []string
	symt     *sym.Symbol
	thearch  Arch

	theline    string
	windowsgui bool

	zeros [512]byte
}

func NewPackageSession(bio_psess *bio.PackageSession, objabi_psess *objabi.PackageSession, sys_psess *sys.PackageSession, loadelf_psess *loadelf.PackageSession, loadmacho_psess *loadmacho.PackageSession, loadpe_psess *loadpe.PackageSession, objfile_psess *objfile.PackageSession, sym_psess *sym.PackageSession, gcprog_psess *gcprog.PackageSession, src_psess *src.PackageSession, dwarf_psess *dwarf.PackageSession, obj_psess *obj.PackageSession) *PackageSession {
	psess := &PackageSession{}
	psess.bio = bio_psess
	psess.objabi = objabi_psess
	psess.sys = sys_psess
	psess.loadelf = loadelf_psess
	psess.loadmacho = loadmacho_psess
	psess.loadpe = loadpe_psess
	psess.objfile = objfile_psess
	psess.sym = sym_psess
	psess.gcprog = gcprog_psess
	psess.src = src_psess
	psess.dwarf = dwarf_psess
	psess.obj = obj_psess
	psess.externalobj = false
	psess.internalpkg = []string{
		"crypto/x509",
		"net",
		"os/user",
		"runtime/cgo",
		"runtime/race",
		"runtime/msan",
	}
	psess.strdata = make(map[string]string)
	psess.loadBudget = INITIAL_MACHO_HEADR - 2*1024
	psess.start = time.Now()
	psess.flagBuildid = flag.String("buildid", "", "record `id` as Go toolchain build id")
	psess.flagOutfile = flag.String("o", "", "write output to `file`")
	psess.flagPluginPath = flag.String("pluginpath", "", "full path name for plugin")
	psess.flagInstallSuffix = flag.String("installsuffix", "", "set package directory `suffix`")
	psess.flagDumpDep = flag.Bool("dumpdep", false, "dump symbol dependency graph")
	psess.flagRace = flag.Bool("race", false, "enable race detector")
	psess.flagMsan = flag.Bool("msan", false, "enable MSan interface")
	psess.flagFieldTrack = flag.String("k", "", "set field tracking `symbol`")
	psess.flagLibGCC = flag.String("libgcc", "", "compiler support lib for internal linking; use \"none\" to disable")
	psess.flagTmpdir = flag.String("tmpdir", "", "use `directory` for temporary files")
	psess.flagExtld = flag.String("extld", "", "use `linker` when linking in external mode")
	psess.flagExtldflags = flag.String("extldflags", "", "pass `flags` to external linker")
	psess.flagExtar = flag.String("extar", "", "archive program for buildmode=c-archive")
	psess.flagA = flag.Bool("a", false, "disassemble output")
	psess.FlagC = flag.Bool("c", false, "dump call graph")
	psess.FlagD = flag.Bool("d", false, "disable dynamic executable")
	psess.flagF = flag.Bool("f", false, "ignore version mismatch")
	psess.flagG = flag.Bool("g", false, "disable go package data checks")
	psess.flagH = flag.Bool("h", false, "halt on error")
	psess.flagN = flag.Bool("n", false, "dump symbol table")
	psess.FlagS = flag.Bool("s", false, "disable symbol table")
	psess.flagU = flag.Bool("u", false, "reject unsafe packages")
	psess.FlagW = flag.Bool("w", false, "disable DWARF generation")
	psess.flagInterpreter = flag.String("I", "", "use `linker` as ELF dynamic linker")
	psess.FlagDebugTramp = flag.Int("debugtramp", 0, "debug trampolines")
	psess.FlagRound = flag.Int("R", -1, "set address rounding `quantum`")
	psess.FlagTextAddr = flag.Int64("T", -1, "set text segment `address`")
	psess.FlagDataAddr = flag.Int64("D", -1, "set data segment `address`")
	psess.flagEntrySymbol = flag.String("E", "", "set `entry` symbol name")
	psess.cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	psess.memprofile = flag.String("memprofile", "", "write memory profile to `file`")
	psess.memprofilerate = flag.Int64("memprofilerate", 0, "set runtime.MemProfileRate to `rate`")
	psess.seenlib = make(map[string]bool)
	psess.numelfsym = 1
	psess.Nelfsym = 1
	psess.ELF_NOTE_NETBSD_NAME = []byte("NetBSD\x00")
	psess.ELF_NOTE_OPENBSD_NAME = []byte("OpenBSD\x00")
	psess.ELF_NOTE_BUILDINFO_NAME = []byte("GNU\x00")
	psess.ELF_NOTE_GO_NAME = []byte("Go\x00\x00")
	psess.PESECTALIGN = 0x1000
	psess.PEFILEALIGN = 2 << 8
	psess.dosstub = []uint8{
		0x4d,
		0x5a,
		0x90,
		0x00,
		0x03,
		0x00,
		0x04,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0xff,
		0xff,
		0x00,
		0x00,
		0x8b,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x40,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x80,
		0x00,
		0x00,
		0x00,
		0x0e,
		0x1f,
		0xba,
		0x0e,
		0x00,
		0xb4,
		0x09,
		0xcd,
		0x21,
		0xb8,
		0x01,
		0x4c,
		0xcd,
		0x21,
		0x54,
		0x68,
		0x69,
		0x73,
		0x20,
		0x70,
		0x72,
		0x6f,
		0x67,
		0x72,
		0x61,
		0x6d,
		0x20,
		0x63,
		0x61,
		0x6e,
		0x6e,
		0x6f,
		0x74,
		0x20,
		0x62,
		0x65,
		0x20,
		0x72,
		0x75,
		0x6e,
		0x20,
		0x69,
		0x6e,
		0x20,
		0x44,
		0x4f,
		0x53,
		0x20,
		0x6d,
		0x6f,
		0x64,
		0x65,
		0x2e,
		0x0d,
		0x0d,
		0x0a,
		0x24,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
		0x00,
	}
	psess.prefixBuf = []byte(dwarf.InfoPrefix)
	return psess
}

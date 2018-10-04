package ld

type PackageState struct {
	bio                     *bio.PackageState
	dwarf                   *dwarf.PackageState
	gcprog                  *gcprog.PackageState
	loadelf                 *loadelf.PackageState
	loadmacho               *loadmacho.PackageState
	loadpe                  *loadpe.PackageState
	obj                     *obj.PackageState
	objabi                  *objabi.PackageState
	objfile                 *objfile.PackageState
	src                     *src.PackageState
	sym                     *sym.PackageState
	sys                     *sys.PackageState
	ELF_NOTE_BUILDINFO_NAME []byte
	ELF_NOTE_GO_NAME        []byte
	ELF_NOTE_NETBSD_NAME    []byte
	ELF_NOTE_OPENBSD_NAME   []byte
	Elfstrdat               []byte
	Flag8                   bool
	FlagC                   *bool
	FlagD                   *bool
	FlagDataAddr            *int64
	FlagDebugTramp          *int
	FlagRound               *int
	FlagS                   *bool
	FlagTextAddr            *int64
	FlagW                   *bool
	Funcalign               int
	HEADR                   int32
	Lcsize                  int32
	Nelfsym                 int
	PEFILEALIGN             int64
	PEFILEHEADR             int32
	PESECTALIGN             int64
	PESECTHEADR             int32
	Segdata                 sym.Segment
	Segdwarf                sym.Segment
	Segrelrodata            sym.Segment
	Segrodata               sym.Segment
	Segtext                 sym.Segment
	Spsize                  int32
	Symsize                 int32
	atExitFuncs             []func()
	buildinfo               []byte
	cpuprofile              *string
	createTrivialCOnce      sync.Once
	datap                   []*sym.Symbol
	debug_s                 bool
	dexport                 [1024]*sym.Symbol
	dosstub                 []uint8
	dr                      *Dll
	dwarfaddr               int64
	dwarfp                  []*sym.Symbol
	dwarfstart, linkstart   int64
	dwglobals               dwarf.DWDie
	dwroot                  dwarf.DWDie
	dwtypes                 dwarf.DWDie
	dylib                   []string
	dynexp                  []*sym.Symbol
	dynlib                  []string
	ehdr                    ElfEhdr
	elf64                   bool
	elfRelType              string
	elfbind                 int
	elfglobalsymndx         int
	elfstr                  [100]Elfstring
	elfverneed              int
	externalobj             bool
	flagA                   *bool
	flagBuildid             *string
	flagDumpDep             *bool
	flagEntrySymbol         *string
	flagExtar               *string
	flagExtld               *string
	flagExtldflags          *string
	flagF                   *bool
	flagFieldTrack          *string
	flagG                   *bool
	flagH                   *bool
	flagInstallSuffix       *string
	flagInterpreter         *string
	flagLibGCC              *string
	flagMsan                *bool
	flagN                   *bool
	flagOutfile             *string
	flagPluginPath          *string
	flagRace                *bool
	flagTmpdir              *string
	flagU                   *bool
	gdbscript               string
	havedynamic             int
	hostobj                 []Hostobj
	internalpkg             []string
	interp                  string
	interpreter             string
	iscgo                   bool
	ldflag                  []string
	linkoff                 int64
	linkoffset              uint32
	liveness                int64
	load                    []MachoLoad
	loadBudget              int
	machohdr                MachoHdr
	memprofile              *string
	memprofilerate          *int64
	morestack               *sym.Symbol
	ndebug                  int
	nelfstr                 int
	nerrors                 int
	nexport                 int
	nkind                   [3]int
	nsect                   int
	nseg                    int
	nsortsym                int
	numelfsym               int
	pclntabFiletabOffset    int32
	pclntabFirstFunc        *sym.Symbol
	pclntabLastFunc         *sym.Symbol
	pclntabNfunc            int32
	pclntabPclntabOffset    int32
	pclntabZpcln            sym.FuncInfo
	pe64                    int
	pefile                  peFile
	phdr                    [400]*ElfPhdr
	pkgall                  []*Pkg
	pkglistfornote          []byte
	prefixBuf               []byte
	prototypedies           map[string]*dwarf.DWDie
	realdwarf, linkseg      *macho.Segment
	rpath                   Rpath
	rsrcsym                 *sym.Symbol
	seenlib                 map[string]bool
	seg                     [16]MachoSeg
	shdr                    [400]*ElfShdr
	sortsym                 []*sym.Symbol
	start                   time.Time
	startTime               time.Time
	strdata                 map[string]string
	strnames                []string
	symt                    *sym.Symbol
	thearch                 Arch
	theline                 string
	windowsgui              bool
	zeros                   [512]byte
}

func NewPackageState(dwarf_pstate *dwarf.PackageState, obj_pstate *obj.PackageState, objabi_pstate *objabi.PackageState, sys_pstate *sys.PackageState, sym_pstate *sym.PackageState, bio_pstate *bio.PackageState, loadelf_pstate *loadelf.PackageState, loadmacho_pstate *loadmacho.PackageState, loadpe_pstate *loadpe.PackageState, objfile_pstate *objfile.PackageState, src_pstate *src.PackageState, gcprog_pstate *gcprog.PackageState) *PackageState {
	pstate := &PackageState{}
	pstate.dwarf = dwarf_pstate
	pstate.obj = obj_pstate
	pstate.objabi = objabi_pstate
	pstate.sys = sys_pstate
	pstate.sym = sym_pstate
	pstate.bio = bio_pstate
	pstate.loadelf = loadelf_pstate
	pstate.loadmacho = loadmacho_pstate
	pstate.loadpe = loadpe_pstate
	pstate.objfile = objfile_pstate
	pstate.src = src_pstate
	pstate.gcprog = gcprog_pstate
	pstate.prefixBuf = []byte(dwarf.InfoPrefix)
	pstate.externalobj = false
	pstate.internalpkg = []string{
		"crypto/x509",
		"net",
		"os/user",
		"runtime/cgo",
		"runtime/race",
		"runtime/msan",
	}
	pstate.start = time.Now()
	pstate.numelfsym = 1
	pstate.loadBudget = INITIAL_MACHO_HEADR - 2*1024
	pstate.flagBuildid = flag.String("buildid", "", "record `id` as Go toolchain build id")
	pstate.flagOutfile = flag.String("o", "", "write output to `file`")
	pstate.flagPluginPath = flag.String("pluginpath", "", "full path name for plugin")
	pstate.flagInstallSuffix = flag.String("installsuffix", "", "set package directory `suffix`")
	pstate.flagDumpDep = flag.Bool("dumpdep", false, "dump symbol dependency graph")
	pstate.flagRace = flag.Bool("race", false, "enable race detector")
	pstate.flagMsan = flag.Bool("msan", false, "enable MSan interface")
	pstate.flagFieldTrack = flag.String("k", "", "set field tracking `symbol`")
	pstate.flagLibGCC = flag.String("libgcc", "", "compiler support lib for internal linking; use \"none\" to disable")
	pstate.flagTmpdir = flag.String("tmpdir", "", "use `directory` for temporary files")
	pstate.flagExtld = flag.String("extld", "", "use `linker` when linking in external mode")
	pstate.flagExtldflags = flag.String("extldflags", "", "pass `flags` to external linker")
	pstate.flagExtar = flag.String("extar", "", "archive program for buildmode=c-archive")
	pstate.flagA = flag.Bool("a", false, "disassemble output")
	pstate.FlagC = flag.Bool("c", false, "dump call graph")
	pstate.FlagD = flag.Bool("d", false, "disable dynamic executable")
	pstate.flagF = flag.Bool("f", false, "ignore version mismatch")
	pstate.flagG = flag.Bool("g", false, "disable go package data checks")
	pstate.flagH = flag.Bool("h", false, "halt on error")
	pstate.flagN = flag.Bool("n", false, "dump symbol table")
	pstate.FlagS = flag.Bool("s", false, "disable symbol table")
	pstate.flagU = flag.Bool("u", false, "reject unsafe packages")
	pstate.FlagW = flag.Bool("w", false, "disable DWARF generation")
	pstate.flagInterpreter = flag.String("I", "", "use `linker` as ELF dynamic linker")
	pstate.FlagDebugTramp = flag.Int("debugtramp", 0, "debug trampolines")
	pstate.FlagRound = flag.Int("R", -1, "set address rounding `quantum`")
	pstate.FlagTextAddr = flag.Int64("T", -1, "set text segment `address`")
	pstate.FlagDataAddr = flag.Int64("D", -1, "set data segment `address`")
	pstate.flagEntrySymbol = flag.String("E", "", "set `entry` symbol name")
	pstate.cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	pstate.memprofile = flag.String("memprofile", "", "write memory profile to `file`")
	pstate.memprofilerate = flag.Int64("memprofilerate", 0, "set runtime.MemProfileRate to `rate`")
	pstate.Nelfsym = 1
	pstate.ELF_NOTE_NETBSD_NAME = []byte("NetBSD\x00")
	pstate.ELF_NOTE_OPENBSD_NAME = []byte("OpenBSD\x00")
	pstate.ELF_NOTE_BUILDINFO_NAME = []byte("GNU\x00")
	pstate.ELF_NOTE_GO_NAME = []byte("Go\x00\x00")
	pstate.seenlib = make(map[string]bool)
	pstate.strdata = make(map[string]string)
	pstate.PESECTALIGN = 0x1000
	pstate.PEFILEALIGN = 2 << 8
	pstate.dosstub = []uint8{
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
	return pstate
}

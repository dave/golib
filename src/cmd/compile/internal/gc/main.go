package gc

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"

	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/bio"
	"github.com/dave/golib/src/cmd/internal/dwarf"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"github.com/dave/golib/src/cmd/internal/sys"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

// Debug arguments.
// These can be specified with the -d flag, as in "-d nil"
// to set the debug_checknil variable.
// Multiple options can be comma-separated.
// Each option accepts an optional argument, as in "gcprog=2"

// must be *int or *string

const debugHelpHeader = "usage: -d arg[,arg]* and arg is <key>[=<value>]\n\n<key> is one of:\n\n"

const debugHelpFooter = "\n<value> is key-specific.\n\nKey \"pctab\" supports values:\n\t\"pctospadj\", \"pctofile\", \"pctoline\", \"pctoinline\", \"pctopcdata\"\n"

func (psess *PackageSession) usage() {
	fmt.Fprintf(os.Stderr, "usage: compile [options] file.go...\n")
	objabi.Flagprint(os.Stderr)
	psess.
		Exit(2)
}

func (psess *PackageSession) hidePanic() {
	if psess.Debug_panic == 0 && psess.nsavederrors+psess.nerrors > 0 {

		if err := recover(); err != nil {
			psess.
				errorexit()
		}
	}
}

// supportsDynlink reports whether or not the code generator for the given
// architecture supports the -shared and -dynlink flags.
func supportsDynlink(arch *sys.Arch) bool {
	return arch.InFamily(sys.AMD64, sys.ARM, sys.ARM64, sys.I386, sys.PPC64, sys.S390X)
}

// timing data for compiler phases

// Main parses flags and Go source files specified in the command-line
// arguments, type-checks the parsed Go package, compiles functions to machine
// code, and finally writes the compiled package definition to disk.
func (psess *PackageSession) Main(archInit func(*Arch)) {
	psess.
		timings.Start("fe", "init")

	defer psess.hidePanic()

	archInit(&psess.thearch)
	psess.
		Ctxt = psess.obj.Linknew(psess.thearch.LinkArch)
	psess.
		Ctxt.DiagFunc = psess.yyerror
	psess.
		Ctxt.DiagFlush = psess.flusherrors
	psess.
		Ctxt.Bso = bufio.NewWriter(os.Stdout)
	psess.
		localpkg = psess.types.NewPkg("", "")
	psess.
		localpkg.Prefix = "\"\""
	psess.
		localpkg.Height = types.MaxPkgHeight
	psess.
		builtinpkg = psess.types.NewPkg("go.builtin", "")
	psess.
		builtinpkg.Prefix = "go.builtin"
	psess.
		unsafepkg = psess.types.NewPkg("unsafe", "unsafe")
	psess.
		Runtimepkg = psess.types.NewPkg("go.runtime", "runtime")
	psess.
		Runtimepkg.Prefix = "runtime"
	psess.
		itabpkg = psess.types.NewPkg("go.itab", "go.itab")
	psess.
		itabpkg.Prefix = "go.itab"
	psess.
		itablinkpkg = psess.types.NewPkg("go.itablink", "go.itablink")
	psess.
		itablinkpkg.Prefix = "go.itablink"
	psess.
		trackpkg = psess.types.NewPkg("go.track", "go.track")
	psess.
		trackpkg.Prefix = "go.track"
	psess.
		mappkg = psess.types.NewPkg("go.map", "go.map")
	psess.
		mappkg.Prefix = "go.map"
	psess.
		gopkg = psess.types.NewPkg("go", "")
	psess.
		Nacl = psess.objabi.GOOS == "nacl"
	Wasm := psess.objabi.GOARCH == "wasm"

	flag.BoolVar(&psess.compiling_runtime, "+", false, "compiling runtime")
	flag.BoolVar(&psess.compiling_std, "std", false, "compiling standard library")
	objabi.Flagcount("%", "debug non-static initializers", &psess.Debug['%'])
	objabi.Flagcount("B", "disable bounds checking", &psess.Debug['B'])
	objabi.Flagcount("C", "disable printing of columns in error messages", &psess.Debug['C'])
	flag.StringVar(&psess.localimport, "D", "", "set relative `path` for local imports")
	objabi.Flagcount("E", "debug symbol export", &psess.Debug['E'])
	objabi.Flagfn1("I", "add `directory` to import search path", psess.addidir)
	objabi.Flagcount("K", "debug missing line numbers", &psess.Debug['K'])
	objabi.Flagcount("L", "show full file names in error messages", &psess.Debug['L'])
	objabi.Flagcount("N", "disable optimizations", &psess.Debug['N'])
	flag.BoolVar(&psess.Debug_asm, "S", false, "print assembly listing")
	objabi.AddVersionFlag()
	objabi.Flagcount("W", "debug parse tree after type checking", &psess.Debug['W'])
	flag.StringVar(&psess.asmhdr, "asmhdr", "", "write assembly header to `file`")
	flag.StringVar(&psess.buildid, "buildid", "", "record `id` as the build id in the export metadata")
	flag.IntVar(&psess.nBackendWorkers, "c", 1, "concurrency during compilation, 1 means no concurrency")
	flag.BoolVar(&psess.pure_go, "complete", false, "compiling complete package (no C or assembly)")
	flag.StringVar(&psess.debugstr, "d", "", "print debug information about items in `list`; try -d help")
	flag.BoolVar(&psess.flagDWARF, "dwarf", !Wasm, "generate DWARF symbols")
	flag.BoolVar(&psess.Ctxt.Flag_locationlists, "dwarflocationlists", true, "add location lists to DWARF in optimized mode")
	flag.IntVar(&psess.genDwarfInline, "gendwarfinl", 2, "generate DWARF inline info records")
	objabi.Flagcount("e", "no limit on number of errors reported", &psess.Debug['e'])
	objabi.Flagcount("f", "debug stack frames", &psess.Debug['f'])
	objabi.Flagcount("h", "halt on error", &psess.Debug['h'])
	objabi.Flagcount("i", "debug line number stack", &psess.Debug['i'])
	objabi.Flagfn1("importmap", "add `definition` of the form source=actual to import map", psess.addImportMap)
	objabi.Flagfn1("importcfg", "read import configuration from `file`", psess.readImportCfg)
	flag.StringVar(&psess.flag_installsuffix, "installsuffix", "", "set pkg directory `suffix`")
	objabi.Flagcount("j", "debug runtime-initialized variables", &psess.Debug['j'])
	objabi.Flagcount("l", "disable inlining", &psess.Debug['l'])
	flag.StringVar(&psess.linkobj, "linkobj", "", "write linker-specific object to `file`")
	objabi.Flagcount("live", "debug liveness analysis", &psess.debuglive)
	objabi.Flagcount("m", "print optimization decisions", &psess.Debug['m'])
	flag.BoolVar(&psess.flag_msan, "msan", false, "build code compatible with C/C++ memory sanitizer")
	flag.BoolVar(&psess.dolinkobj, "dolinkobj", true, "generate linker-specific objects; if false, some invalid code may compile")
	flag.BoolVar(&psess.nolocalimports, "nolocalimports", false, "reject local (relative) imports")
	flag.StringVar(&psess.outfile, "o", "", "write output to `file`")
	flag.StringVar(&psess.myimportpath, "p", "", "set expected package import `path`")
	flag.BoolVar(&psess.writearchive, "pack", false, "write to file.a instead of file.o")
	objabi.Flagcount("r", "debug generated wrappers", &psess.Debug['r'])
	flag.BoolVar(&psess.flag_race, "race", false, "enable race detector")
	objabi.Flagcount("s", "warn about composite literals that can be simplified", &psess.Debug['s'])
	flag.StringVar(&psess.pathPrefix, "trimpath", "", "remove `prefix` from recorded source file paths")
	flag.BoolVar(&psess.safemode, "u", false, "reject unsafe code")
	flag.BoolVar(&psess.Debug_vlog, "v", false, "increase debug verbosity")
	objabi.Flagcount("w", "debug type checking", &psess.Debug['w'])
	flag.BoolVar(&psess.use_writebarrier, "wb", true, "enable write barrier")
	var flag_shared bool
	var flag_dynlink bool
	if supportsDynlink(psess.thearch.LinkArch.Arch) {
		flag.BoolVar(&flag_shared, "shared", false, "generate code that can be linked into a shared library")
		flag.BoolVar(&flag_dynlink, "dynlink", false, "support references to Go symbols defined in other shared libraries")
	}
	flag.StringVar(&psess.cpuprofile, "cpuprofile", "", "write cpu profile to `file`")
	flag.StringVar(&psess.memprofile, "memprofile", "", "write memory profile to `file`")
	flag.Int64Var(&psess.memprofilerate, "memprofilerate", 0, "set runtime.MemProfileRate to `rate`")
	var goversion string
	flag.StringVar(&goversion, "goversion", "", "required version of the runtime")
	flag.StringVar(&psess.traceprofile, "traceprofile", "", "write an execution trace to `file`")
	flag.StringVar(&psess.blockprofile, "blockprofile", "", "write block profile to `file`")
	flag.StringVar(&psess.mutexprofile, "mutexprofile", "", "write mutex profile to `file`")
	flag.StringVar(&psess.benchfile, "bench", "", "append benchmark times to `file`")
	flag.BoolVar(&psess.flagiexport, "iexport", true, "export indexed package data")
	objabi.Flagparse(psess.usage)
	psess.
		recordFlags("B", "N", "l", "msan", "race", "shared", "dynlink", "dwarflocationlists")
	psess.
		Ctxt.Flag_shared = flag_dynlink || flag_shared
	psess.
		Ctxt.Flag_dynlink = flag_dynlink
	psess.
		Ctxt.Flag_optimize = psess.Debug['N'] == 0
	psess.
		Ctxt.Debugasm = psess.Debug_asm
	psess.
		Ctxt.Debugvlog = psess.Debug_vlog
	if psess.flagDWARF {
		psess.
			Ctxt.DebugInfo = psess.debuginfo
		psess.
			Ctxt.GenAbstractFunc = psess.genAbstractFunc
		psess.
			Ctxt.DwFixups = obj.NewDwarfFixupTable(psess.Ctxt)
	} else {
		psess.
			genDwarfInline = 0
		psess.
			Ctxt.Flag_locationlists = false
	}

	if flag.NArg() < 1 && psess.debugstr != "help" && psess.debugstr != "ssa/help" {
		psess.
			usage()
	}

	if goversion != "" && goversion != runtime.Version() {
		fmt.Printf("compile: version %q does not match go tool version %q\n", runtime.Version(), goversion)
		psess.
			Exit(2)
	}
	psess.
		thearch.LinkArch.Init(psess.Ctxt)

	if psess.outfile == "" {
		p := flag.Arg(0)
		if i := strings.LastIndex(p, "/"); i >= 0 {
			p = p[i+1:]
		}
		if runtime.GOOS == "windows" {
			if i := strings.LastIndex(p, "\\"); i >= 0 {
				p = p[i+1:]
			}
		}
		if i := strings.LastIndex(p, "."); i >= 0 {
			p = p[:i]
		}
		suffix := ".o"
		if psess.writearchive {
			suffix = ".a"
		}
		psess.
			outfile = p + suffix
	}
	psess.
		startProfile()

	if psess.flag_race && psess.flag_msan {
		log.Fatal("cannot use both -race and -msan")
	}
	if psess.ispkgin(psess.omit_pkgs) {
		psess.
			flag_race = false
		psess.
			flag_msan = false
	}
	if psess.flag_race {
		psess.
			racepkg = psess.types.NewPkg("runtime/race", "")
	}
	if psess.flag_msan {
		psess.
			msanpkg = psess.types.NewPkg("runtime/msan", "")
	}
	if psess.flag_race || psess.flag_msan {
		psess.
			instrumenting = true
	}

	if psess.compiling_runtime && psess.Debug['N'] != 0 {
		log.Fatal("cannot disable optimizations while compiling runtime")
	}
	if psess.nBackendWorkers < 1 {
		log.Fatalf("-c must be at least 1, got %d", psess.nBackendWorkers)
	}
	if psess.nBackendWorkers > 1 && !psess.concurrentBackendAllowed() {
		log.Fatalf("cannot use concurrent backend compilation with provided flags; invoked as %v", os.Args)
	}
	if psess.Ctxt.Flag_locationlists && len(psess.Ctxt.Arch.DWARFRegisters) == 0 {
		log.Fatalf("location lists requested but register mapping not available on %v", psess.Ctxt.Arch.Name)
	}

	if psess.debugstr != "" {
	Split:
		for _, name := range strings.Split(psess.debugstr, ",") {
			if name == "" {
				continue
			}

			if name == "help" {
				fmt.Printf(debugHelpHeader)
				maxLen := len("ssa/help")
				for _, t := range psess.debugtab {
					if len(t.name) > maxLen {
						maxLen = len(t.name)
					}
				}
				for _, t := range psess.debugtab {
					fmt.Printf("\t%-*s\t%s\n", maxLen, t.name, t.help)
				}

				fmt.Printf("\t%-*s\t%s\n", maxLen, "ssa/help", "print help about SSA debugging")
				fmt.Printf(debugHelpFooter)
				os.Exit(0)
			}
			val, valstring, haveInt := 1, "", true
			if i := strings.IndexAny(name, "=:"); i >= 0 {
				var err error
				name, valstring = name[:i], name[i+1:]
				val, err = strconv.Atoi(valstring)
				if err != nil {
					val, haveInt = 1, false
				}
			}
			for _, t := range psess.debugtab {
				if t.name != name {
					continue
				}
				switch vp := t.val.(type) {
				case nil:

				case *string:
					*vp = valstring
				case *int:
					if !haveInt {
						log.Fatalf("invalid debug value %v", name)
					}
					*vp = val
				default:
					panic("bad debugtab type")
				}
				continue Split
			}

			if strings.HasPrefix(name, "ssa/") {

				phase := name[4:]
				flag := "debug"
				if i := strings.Index(phase, "/"); i >= 0 {
					flag = phase[i+1:]
					phase = phase[:i]
				}
				err := psess.ssa.PhaseOption(phase, flag, val, valstring)
				if err != "" {
					log.Fatalf(err)
				}
				continue Split
			}
			log.Fatalf("unknown debug key -d %s\n", name)
		}
	}
	psess.
		Ctxt.Debugpcln = psess.Debug_pctab
	if psess.flagDWARF {
		psess.dwarf.
			EnableLogging(psess.Debug_gendwarfinl != 0)
	}

	if psess.Debug_softfloat != 0 {
		psess.
			thearch.SoftFloat = true
	}

	if psess.Debug['l'] <= 1 {
		psess.
			Debug['l'] = 1 - psess.Debug['l']
	}
	psess.
		trackScopes = psess.flagDWARF
	psess.
		Widthptr = psess.thearch.LinkArch.PtrSize
	psess.
		Widthreg = psess.thearch.LinkArch.RegSize
	psess.types.
		Widthptr = psess.Widthptr
	psess.types.
		Dowidth = psess.dowidth
	psess.types.
		Fatalf = psess.Fatalf
	psess.types.
		Sconv = func(s *types.Sym, flag, mode int) string {
		return psess.sconv(s, FmtFlag(flag), fmtMode(mode))
	}
	psess.types.
		Tconv = func(t *types.Type, flag, mode, depth int) string {
		return psess.tconv(t, FmtFlag(flag), fmtMode(mode), depth)
	}
	psess.types.
		FormatSym = func(sym *types.Sym, s fmt.State, verb rune, mode int) {
		psess.
			symFormat(sym, s, verb, fmtMode(mode))
	}
	psess.types.
		FormatType = func(t *types.Type, s fmt.State, verb rune, mode int) {
		psess.
			typeFormat(t, s, verb, fmtMode(mode))
	}
	psess.types.
		TypeLinkSym = func(t *types.Type) *obj.LSym {
		return psess.typenamesym(t).Linksym(psess.types)
	}
	psess.types.
		FmtLeft = int(FmtLeft)
	psess.types.
		FmtUnsigned = int(FmtUnsigned)
	psess.types.
		FErr = FErr
	psess.types.
		Ctxt = psess.Ctxt
	psess.
		initUniverse()
	psess.
		dclcontext = PEXTERN
	psess.
		nerrors = 0
	psess.
		autogeneratedPos = psess.makePos(src.NewFileBase("<autogenerated>", "<autogenerated>"), 1, 0)
	psess.
		timings.Start("fe", "loadsys")
	psess.
		loadsys()
	psess.
		timings.Start("fe", "parse")
	lines := psess.parseFiles(flag.Args())
	psess.
		timings.Stop()
	psess.
		timings.AddEvent(int64(lines), "lines")
	psess.
		finishUniverse()
	psess.
		typecheckok = true
	if psess.Debug['f'] != 0 {
		psess.
			frame(1)
	}
	psess.
		defercheckwidth()
	psess.
		timings.Start("fe", "typecheck", "top1")
	for i := 0; i < len(psess.xtop); i++ {
		n := psess.xtop[i]
		if op := n.Op; op != ODCL && op != OAS && op != OAS2 {
			psess.
				xtop[i] = psess.typecheck(n, Etop)
		}
	}
	psess.
		timings.Start("fe", "typecheck", "top2")
	for i := 0; i < len(psess.xtop); i++ {
		n := psess.xtop[i]
		if op := n.Op; op == ODCL || op == OAS || op == OAS2 {
			psess.
				xtop[i] = psess.typecheck(n, Etop)
		}
	}
	psess.
		resumecheckwidth()
	psess.
		timings.Start("fe", "typecheck", "func")
	var fcount int64
	for i := 0; i < len(psess.xtop); i++ {
		n := psess.xtop[i]
		if op := n.Op; op == ODCLFUNC || op == OCLOSURE {
			psess.
				Curfn = n
			psess.
				decldepth = 1
			psess.
				saveerrors()
			psess.
				typecheckslice(psess.Curfn.Nbody.Slice(), Etop)
			psess.
				checkreturn(psess.Curfn)
			if psess.nerrors != 0 {
				psess.
					Curfn.Nbody.Set(nil)
			}
			psess.
				deadcode(psess.Curfn)
			fcount++
		}
	}
	psess.
		checkMapKeys()
	psess.
		timings.AddEvent(fcount, "funcs")

	if psess.nsavederrors+psess.nerrors != 0 {
		psess.
			errorexit()
	}
	psess.
		timings.Start("fe", "capturevars")
	for _, n := range psess.xtop {
		if n.Op == ODCLFUNC && n.Func.Closure != nil {
			psess.
				Curfn = n
			psess.
				capturevars(n)
		}
	}
	psess.
		capturevarscomplete = true
	psess.
		Curfn = nil

	if psess.nsavederrors+psess.nerrors != 0 {
		psess.
			errorexit()
	}
	psess.
		timings.Start("fe", "inlining")
	if psess.Debug_typecheckinl != 0 {

		for _, n := range psess.importlist {
			if n.Func.Inl != nil {
				psess.
					saveerrors()
				psess.
					typecheckinl(n)
			}
		}

		if psess.nsavederrors+psess.nerrors != 0 {
			psess.
				errorexit()
		}
	}

	if psess.Debug['l'] != 0 {
		psess.
			visitBottomUp(psess.xtop, func(list []*Node, recursive bool) {
				for _, n := range list {
					if !recursive {
						psess.
							caninl(n)
					} else {
						if psess.Debug['m'] > 1 {
							fmt.Printf("%v: cannot inline %v: recursive\n", n.Line(psess), n.Func.Nname)
						}
					}
					psess.
						inlcalls(n)
				}
			})
	}
	psess.
		timings.Start("fe", "escapes")
	psess.
		escapes(psess.xtop)

	if psess.dolinkobj {

		if psess.compiling_runtime {
			psess.
				nowritebarrierrecCheck = psess.newNowritebarrierrecChecker()
		}
		psess.
			timings.Start("fe", "xclosures")
		for _, n := range psess.xtop {
			if n.Op == ODCLFUNC && n.Func.Closure != nil {
				psess.
					Curfn = n
				psess.
					transformclosure(n)
			}
		}
		psess.
			initssaconfig()
		psess.
			Curfn = nil
		psess.
			peekitabs()
		psess.
			timings.Start("be", "compilefuncs")
		fcount = 0
		for i := 0; i < len(psess.xtop); i++ {
			n := psess.xtop[i]
			if n.Op == ODCLFUNC {
				psess.
					funccompile(n)
				fcount++
			}
		}
		psess.
			timings.AddEvent(fcount, "funcs")

		if psess.nsavederrors+psess.nerrors == 0 {
			psess.
				fninit(psess.xtop)
		}
		psess.
			compileFunctions()

		if psess.nowritebarrierrecCheck != nil {
			psess.
				nowritebarrierrecCheck.check(psess)
			psess.
				nowritebarrierrecCheck = nil
		}

		if psess.Ctxt.DwFixups != nil {
			psess.
				Ctxt.DwFixups.Finalize(psess.myimportpath, psess.Debug_gendwarfinl != 0)
			psess.
				Ctxt.DwFixups = nil
			psess.
				genDwarfInline = 0
		}
	}
	psess.
		timings.Start("be", "externaldcls")
	for i, n := range psess.externdcl {
		if n.Op == ONAME {
			psess.
				externdcl[i] = psess.typecheck(psess.externdcl[i], Erv)
		}
	}

	if psess.nerrors+psess.nsavederrors != 0 {
		psess.
			errorexit()
	}
	psess.
		timings.Start("be", "dumpobj")
	psess.
		dumpobj()
	if psess.asmhdr != "" {
		psess.
			dumpasmhdr()
	}

	obj.SortSlice(psess.largeStackFrames, func(i, j int) bool {
		return psess.largeStackFrames[i].Before(psess.largeStackFrames[j])
	})
	for _, largePos := range psess.largeStackFrames {
		psess.
			yyerrorl(largePos, "stack frame too large (>1GB)")
	}

	if len(psess.compilequeue) != 0 {
		psess.
			Fatalf("%d uncompiled functions", len(psess.compilequeue))
	}

	if psess.nerrors+psess.nsavederrors != 0 {
		psess.
			errorexit()
	}
	psess.
		flusherrors()
	psess.
		timings.Stop()

	if psess.benchfile != "" {
		if err := psess.writebench(psess.benchfile); err != nil {
			log.Fatalf("cannot write benchmark data: %v", err)
		}
	}
}

func (psess *PackageSession) writebench(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	fmt.Fprintln(&buf, "commit:", psess.objabi.Version)
	fmt.Fprintln(&buf, "goos:", runtime.GOOS)
	fmt.Fprintln(&buf, "goarch:", runtime.GOARCH)
	psess.
		timings.Write(&buf, "BenchmarkCompile:"+psess.myimportpath+":")

	n, err := f.Write(buf.Bytes())
	if err != nil {
		return err
	}
	if n != buf.Len() {
		panic("bad writer")
	}

	return f.Close()
}

// nil means not in use

func (psess *PackageSession) addImportMap(s string) {
	if strings.Count(s, "=") != 1 {
		log.Fatal("-importmap argument must be of the form source=actual")
	}
	i := strings.Index(s, "=")
	source, actual := s[:i], s[i+1:]
	if source == "" || actual == "" {
		log.Fatal("-importmap argument must be of the form source=actual; source and actual must be non-empty")
	}
	psess.
		importMap[source] = actual
}

func (psess *PackageSession) readImportCfg(file string) {
	psess.
		packageFile = map[string]string{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("-importcfg: %v", err)
	}

	for lineNum, line := range strings.Split(string(data), "\n") {
		lineNum++
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		var verb, args string
		if i := strings.Index(line, " "); i < 0 {
			verb = line
		} else {
			verb, args = line[:i], strings.TrimSpace(line[i+1:])
		}
		var before, after string
		if i := strings.Index(args, "="); i >= 0 {
			before, after = args[:i], args[i+1:]
		}
		switch verb {
		default:
			log.Fatalf("%s:%d: unknown directive %q", file, lineNum, verb)
		case "importmap":
			if before == "" || after == "" {
				log.Fatalf("%s:%d: invalid importmap: syntax is \"importmap old=new\"", file, lineNum)
			}
			psess.
				importMap[before] = after
		case "packagefile":
			if before == "" || after == "" {
				log.Fatalf("%s:%d: invalid packagefile: syntax is \"packagefile path=filename\"", file, lineNum)
			}
			psess.
				packageFile[before] = after
		}
	}
}

func (psess *PackageSession) saveerrors() {
	psess.
		nsavederrors += psess.nerrors
	psess.
		nerrors = 0
}

func arsize(b *bufio.Reader, name string) int {
	var buf [ArhdrSize]byte
	if _, err := io.ReadFull(b, buf[:]); err != nil {
		return -1
	}
	aname := strings.Trim(string(buf[0:16]), " ")
	if !strings.HasPrefix(aname, name) {
		return -1
	}
	asize := strings.Trim(string(buf[48:58]), " ")
	i, _ := strconv.Atoi(asize)
	return i
}

func (psess *PackageSession) addidir(dir string) {
	if dir != "" {
		psess.
			idirs = append(psess.idirs, dir)
	}
}

func isDriveLetter(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z'
}

// is this path a local name? begins with ./ or ../ or /
func islocalname(name string) bool {
	return strings.HasPrefix(name, "/") ||
		runtime.GOOS == "windows" && len(name) >= 3 && isDriveLetter(name[0]) && name[1] == ':' && name[2] == '/' ||
		strings.HasPrefix(name, "./") || name == "." ||
		strings.HasPrefix(name, "../") || name == ".."
}

func (psess *PackageSession) findpkg(name string) (file string, ok bool) {
	if islocalname(name) {
		if psess.safemode || psess.nolocalimports {
			return "", false
		}

		if psess.packageFile != nil {
			file, ok = psess.packageFile[name]
			return file, ok
		}

		file = fmt.Sprintf("%s.a", name)
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
		file = fmt.Sprintf("%s.o", name)
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
		return "", false
	}

	if q := path.Clean(name); q != name {
		psess.
			yyerror("non-canonical import path %q (should be %q)", name, q)
		return "", false
	}

	if psess.packageFile != nil {
		file, ok = psess.packageFile[name]
		return file, ok
	}

	for _, dir := range psess.idirs {
		file = fmt.Sprintf("%s/%s.a", dir, name)
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
		file = fmt.Sprintf("%s/%s.o", dir, name)
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
	}

	if psess.objabi.GOROOT != "" {
		suffix := ""
		suffixsep := ""
		if psess.flag_installsuffix != "" {
			suffixsep = "_"
			suffix = psess.flag_installsuffix
		} else if psess.flag_race {
			suffixsep = "_"
			suffix = "race"
		} else if psess.flag_msan {
			suffixsep = "_"
			suffix = "msan"
		}

		file = fmt.Sprintf("%s/pkg/%s_%s%s%s/%s.a", psess.objabi.GOROOT, psess.objabi.GOOS, psess.objabi.GOARCH, suffixsep, suffix, name)
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
		file = fmt.Sprintf("%s/pkg/%s_%s%s%s/%s.o", psess.objabi.GOROOT, psess.objabi.GOOS, psess.objabi.GOARCH, suffixsep, suffix, name)
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
	}

	return "", false
}

// loadsys loads the definitions for the low-level runtime functions,
// so that the compiler can generate calls to them,
// but does not make them visible to user code.
func (psess *PackageSession) loadsys() {
	psess.types.
		Block = 1
	psess.
		inimport = true
	psess.
		typecheckok = true
	psess.
		defercheckwidth()

	typs := psess.runtimeTypes()
	for _, d := range psess.runtimeDecls {
		sym := psess.Runtimepkg.Lookup(psess.types, d.name)
		typ := typs[d.typ]
		switch d.tag {
		case funcTag:
			psess.
				importfunc(psess.Runtimepkg, psess.src.NoXPos, sym, typ)
		case varTag:
			psess.
				importvar(psess.Runtimepkg, psess.src.NoXPos, sym, typ)
		default:
			psess.
				Fatalf("unhandled declaration tag %v", d.tag)
		}
	}
	psess.
		typecheckok = false
	psess.
		resumecheckwidth()
	psess.
		inimport = false
}

// myheight tracks the local package's height based on packages
// imported so far.

func (psess *PackageSession) importfile(f *Val) *types.Pkg {
	path_, ok := f.U.(string)
	if !ok {
		psess.
			yyerror("import path must be a string")
		return nil
	}

	if len(path_) == 0 {
		psess.
			yyerror("import path is empty")
		return nil
	}

	if psess.isbadimport(path_, false) {
		return nil
	}

	if path_ == "main" {
		psess.
			yyerror("cannot import \"main\"")
		psess.
			errorexit()
	}

	if psess.myimportpath != "" && path_ == psess.myimportpath {
		psess.
			yyerror("import %q while compiling that package (import cycle)", path_)
		psess.
			errorexit()
	}

	if mapped, ok := psess.importMap[path_]; ok {
		path_ = mapped
	}

	if path_ == "unsafe" {
		if psess.safemode {
			psess.
				yyerror("cannot import package unsafe")
			psess.
				errorexit()
		}
		psess.
			imported_unsafe = true
		return psess.unsafepkg
	}

	if islocalname(path_) {
		if path_[0] == '/' {
			psess.
				yyerror("import path cannot be absolute path")
			return nil
		}

		prefix := psess.Ctxt.Pathname
		if psess.localimport != "" {
			prefix = psess.localimport
		}
		path_ = path.Join(prefix, path_)

		if psess.isbadimport(path_, true) {
			return nil
		}
	}

	file, found := psess.findpkg(path_)
	if !found {
		psess.
			yyerror("can't find import: %q", path_)
		psess.
			errorexit()
	}

	importpkg := psess.types.NewPkg(path_, "")
	if importpkg.Imported {
		return importpkg
	}

	importpkg.Imported = true

	imp, err := bio.Open(file)
	if err != nil {
		psess.
			yyerror("can't open import: %q: %v", path_, err)
		psess.
			errorexit()
	}
	defer imp.Close()

	p, err := imp.ReadString('\n')
	if err != nil {
		psess.
			yyerror("import %s: reading input: %v", file, err)
		psess.
			errorexit()
	}

	if p == "!<arch>\n" {

		sz := arsize(imp.Reader, "__.PKGDEF")
		if sz <= 0 {
			psess.
				yyerror("import %s: not a package file", file)
			psess.
				errorexit()
		}
		p, err = imp.ReadString('\n')
		if err != nil {
			psess.
				yyerror("import %s: reading input: %v", file, err)
			psess.
				errorexit()
		}
	}

	if !strings.HasPrefix(p, "go object ") {
		psess.
			yyerror("import %s: not a go object file: %s", file, p)
		psess.
			errorexit()
	}
	q := fmt.Sprintf("%s %s %s %s\n", psess.objabi.GOOS, psess.objabi.GOARCH, psess.objabi.Version, psess.objabi.Expstring())
	if p[10:] != q {
		psess.
			yyerror("import %s: object is [%s] expected [%s]", file, p[10:], q)
		psess.
			errorexit()
	}

	safe := false
	for {
		p, err = imp.ReadString('\n')
		if err != nil {
			psess.
				yyerror("import %s: reading input: %v", file, err)
			psess.
				errorexit()
		}
		if p == "\n" {
			break
		}
		if strings.HasPrefix(p, "safe") {
			safe = true
			break
		}
	}
	if psess.safemode && !safe {
		psess.
			yyerror("cannot import unsafe package %q", importpkg.Path)
	}

	if psess.packageFile != nil {
		psess.
			Ctxt.AddImport(path_)
	} else {
		psess.
			Ctxt.AddImport(file[len(file)-len(path_)-len(".a"):])
	}

	// look for $$
	var c byte
	for {
		c, err = imp.ReadByte()
		if err != nil {
			break
		}
		if c == '$' {
			c, err = imp.ReadByte()
			if c == '$' || err != nil {
				break
			}
		}
	}

	if err == nil {
		c, _ = imp.ReadByte()
	}

	switch c {
	case '\n':
		psess.
			yyerror("cannot import %s: old export format no longer supported (recompile library)", path_)
		return nil

	case 'B':
		if psess.Debug_export != 0 {
			fmt.Printf("importing %s (%s)\n", path_, file)
		}
		imp.ReadByte()

		c, err = imp.ReadByte()
		if err != nil {
			psess.
				yyerror("import %s: reading input: %v", file, err)
			psess.
				errorexit()
		}

		if c == 'i' {
			if !psess.flagiexport {
				psess.
					yyerror("import %s: cannot import package compiled with -iexport=true", file)
				psess.
					errorexit()
			}
			psess.
				iimport(importpkg, imp)
		} else {
			if psess.flagiexport {
				psess.
					yyerror("import %s: cannot import package compiled with -iexport=false", file)
				psess.
					errorexit()
			}

			imp.UnreadByte()
			psess.
				Import(importpkg, imp.Reader)
		}

	default:
		psess.
			yyerror("no import in %q", path_)
		psess.
			errorexit()
	}

	if importpkg.Height >= psess.myheight {
		psess.
			myheight = importpkg.Height + 1
	}

	return importpkg
}

func (psess *PackageSession) pkgnotused(lineno src.XPos, path string, name string) {

	elem := path
	if i := strings.LastIndex(elem, "/"); i >= 0 {
		elem = elem[i+1:]
	}
	if name == "" || elem == name {
		psess.
			yyerrorl(lineno, "imported and not used: %q", path)
	} else {
		psess.
			yyerrorl(lineno, "imported and not used: %q as %s", path, name)
	}
}

func (psess *PackageSession) mkpackage(pkgname string) {
	if psess.localpkg.Name == "" {
		if pkgname == "_" {
			psess.
				yyerror("invalid package name _")
		}
		psess.
			localpkg.Name = pkgname
	} else {
		if pkgname != psess.localpkg.Name {
			psess.
				yyerror("package %s; expected %s", pkgname, psess.localpkg.Name)
		}
	}
}

func (psess *PackageSession) clearImports() {
	type importedPkg struct {
		pos  src.XPos
		path string
		name string
	}
	var unused []importedPkg

	for _, s := range psess.localpkg.Syms {
		n := asNode(s.Def)
		if n == nil {
			continue
		}
		if n.Op == OPACK {

			if !n.Name.Used() && psess.nsyntaxerrors == 0 {
				unused = append(unused, importedPkg{n.Pos, n.Name.Pkg.Path, s.Name})
			}
			s.Def = nil
			continue
		}
		if IsAlias(s) {

			if n.Name != nil && n.Name.Pack != nil && !n.Name.Pack.Name.Used() && psess.nsyntaxerrors == 0 {
				unused = append(unused, importedPkg{n.Name.Pack.Pos, n.Name.Pack.Name.Pkg.Path, ""})
				n.Name.Pack.Name.SetUsed(true)
			}
			s.Def = nil
			continue
		}
	}

	obj.SortSlice(unused, func(i, j int) bool { return unused[i].pos.Before(unused[j].pos) })
	for _, pkg := range unused {
		psess.
			pkgnotused(pkg.pos, pkg.path, pkg.name)
	}
}

func IsAlias(sym *types.Sym) bool {
	return sym.Def != nil && asNode(sym.Def).Sym != sym
}

// By default, assume any debug flags are incompatible with concurrent compilation.
// A few are safe and potentially in common use for normal compiles, though; mark them as such here.

func (psess *PackageSession) concurrentBackendAllowed() bool {
	for i, x := range psess.Debug {
		if x != 0 && !psess.concurrentFlagOK[i] {
			return false
		}
	}

	if psess.Debug_vlog || psess.debugstr != "" || psess.debuglive > 0 {
		return false
	}

	if psess.objabi.Fieldtrack_enabled != 0 || psess.objabi.Clobberdead_enabled != 0 {
		return false
	}

	if psess.Ctxt.Flag_shared || psess.Ctxt.Flag_dynlink || psess.flag_race {
		return false
	}
	return true
}

// recordFlags records the specified command-line flags to be placed
// in the DWARF info.
func (psess *PackageSession) recordFlags(flags ...string) {
	if psess.myimportpath == "" {

		return
	}

	type BoolFlag interface {
		IsBoolFlag() bool
	}
	type CountFlag interface {
		IsCountFlag() bool
	}
	var cmd bytes.Buffer
	for _, name := range flags {
		f := flag.Lookup(name)
		if f == nil {
			continue
		}
		getter := f.Value.(flag.Getter)
		if getter.String() == f.DefValue {

			continue
		}
		if bf, ok := f.Value.(BoolFlag); ok && bf.IsBoolFlag() {
			val, ok := getter.Get().(bool)
			if ok && val {
				fmt.Fprintf(&cmd, " -%s", f.Name)
				continue
			}
		}
		if cf, ok := f.Value.(CountFlag); ok && cf.IsCountFlag() {
			val, ok := getter.Get().(int)
			if ok && val == 1 {
				fmt.Fprintf(&cmd, " -%s", f.Name)
				continue
			}
		}
		fmt.Fprintf(&cmd, " -%s=%v", f.Name, getter.Get())
	}

	if cmd.Len() == 0 {
		return
	}
	s := psess.Ctxt.Lookup(dwarf.CUInfoPrefix + "producer." + psess.myimportpath)
	s.Type = objabi.SDWARFINFO

	s.Set(obj.AttrDuplicateOK, true)
	psess.
		Ctxt.Data = append(psess.Ctxt.Data, s)
	s.P = cmd.Bytes()[1:]
}

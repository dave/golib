// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run mkbuiltin.go

package gc

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/ssa"
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

const debugHelpHeader = "usage: -d arg[,arg]* and arg is <key>[=<value>]\n\n<key> is one of:\n\n"

const debugHelpFooter = "\n<value> is key-specific.\n\nKey \"pctab\" supports values:\n\t\"pctospadj\", \"pctofile\", \"pctoline\", \"pctoinline\", \"pctopcdata\"\n"

func (pstate *PackageState) usage() {
	fmt.Fprintf(os.Stderr, "usage: compile [options] file.go...\n")
	objabi.Flagprint(os.Stderr)
	pstate.Exit(2)
}

func (pstate *PackageState) hidePanic() {
	if pstate.Debug_panic == 0 && pstate.nsavederrors+pstate.nerrors > 0 {
		// If we've already complained about things
		// in the program, don't bother complaining
		// about a panic too; let the user clean up
		// the code and try again.
		if err := recover(); err != nil {
			pstate.errorexit()
		}
	}
}

// supportsDynlink reports whether or not the code generator for the given
// architecture supports the -shared and -dynlink flags.
func supportsDynlink(arch *sys.Arch) bool {
	return arch.InFamily(sys.AMD64, sys.ARM, sys.ARM64, sys.I386, sys.PPC64, sys.S390X)
}

// Main parses flags and Go source files specified in the command-line
// arguments, type-checks the parsed Go package, compiles functions to machine
// code, and finally writes the compiled package definition to disk.
func (pstate *PackageState) Main(archInit func(*Arch)) {
	pstate.timings.Start("fe", "init")

	defer pstate.hidePanic()

	archInit(&pstate.thearch)

	pstate.Ctxt = pstate.obj.Linknew(pstate.thearch.LinkArch)
	pstate.Ctxt.DiagFunc = pstate.yyerror
	pstate.Ctxt.DiagFlush = pstate.flusherrors
	pstate.Ctxt.Bso = bufio.NewWriter(os.Stdout)

	pstate.localpkg = pstate.types.NewPkg("", "")
	pstate.localpkg.Prefix = "\"\""

	// We won't know localpkg's height until after import
	// processing. In the mean time, set to MaxPkgHeight to ensure
	// height comparisons at least work until then.
	pstate.localpkg.Height = types.MaxPkgHeight

	// pseudo-package, for scoping
	pstate.builtinpkg = pstate.types.NewPkg("go.builtin", "") // TODO(gri) name this package go.builtin?
	pstate.builtinpkg.Prefix = "go.builtin"                   // not go%2ebuiltin

	// pseudo-package, accessed by import "unsafe"
	pstate.unsafepkg = pstate.types.NewPkg("unsafe", "unsafe")

	// Pseudo-package that contains the compiler's builtin
	// declarations for package runtime. These are declared in a
	// separate package to avoid conflicts with package runtime's
	// actual declarations, which may differ intentionally but
	// insignificantly.
	pstate.Runtimepkg = pstate.types.NewPkg("go.runtime", "runtime")
	pstate.Runtimepkg.Prefix = "runtime"

	// pseudo-packages used in symbol tables
	pstate.itabpkg = pstate.types.NewPkg("go.itab", "go.itab")
	pstate.itabpkg.Prefix = "go.itab" // not go%2eitab

	pstate.itablinkpkg = pstate.types.NewPkg("go.itablink", "go.itablink")
	pstate.itablinkpkg.Prefix = "go.itablink" // not go%2eitablink

	pstate.trackpkg = pstate.types.NewPkg("go.track", "go.track")
	pstate.trackpkg.Prefix = "go.track" // not go%2etrack

	// pseudo-package used for map zero values
	pstate.mappkg = pstate.types.NewPkg("go.map", "go.map")
	pstate.mappkg.Prefix = "go.map"

	// pseudo-package used for methods with anonymous receivers
	pstate.gopkg = pstate.types.NewPkg("go", "")

	pstate.Nacl = pstate.objabi.GOOS == "nacl"
	Wasm := pstate.objabi.GOARCH == "wasm"

	flag.BoolVar(&pstate.compiling_runtime, "+", false, "compiling runtime")
	flag.BoolVar(&pstate.compiling_std, "std", false, "compiling standard library")
	objabi.Flagcount("%", "debug non-static initializers", &pstate.Debug['%'])
	objabi.Flagcount("B", "disable bounds checking", &pstate.Debug['B'])
	objabi.Flagcount("C", "disable printing of columns in error messages", &pstate.Debug['C']) // TODO(gri) remove eventually
	flag.StringVar(&pstate.localimport, "D", "", "set relative `path` for local imports")
	objabi.Flagcount("E", "debug symbol export", &pstate.Debug['E'])
	objabi.Flagfn1("I", "add `directory` to import search path", pstate.addidir)
	objabi.Flagcount("K", "debug missing line numbers", &pstate.Debug['K'])
	objabi.Flagcount("L", "show full file names in error messages", &pstate.Debug['L'])
	objabi.Flagcount("N", "disable optimizations", &pstate.Debug['N'])
	flag.BoolVar(&pstate.Debug_asm, "S", false, "print assembly listing")
	objabi.AddVersionFlag() // -V
	objabi.Flagcount("W", "debug parse tree after type checking", &pstate.Debug['W'])
	flag.StringVar(&pstate.asmhdr, "asmhdr", "", "write assembly header to `file`")
	flag.StringVar(&pstate.buildid, "buildid", "", "record `id` as the build id in the export metadata")
	flag.IntVar(&pstate.nBackendWorkers, "c", 1, "concurrency during compilation, 1 means no concurrency")
	flag.BoolVar(&pstate.pure_go, "complete", false, "compiling complete package (no C or assembly)")
	flag.StringVar(&pstate.debugstr, "d", "", "print debug information about items in `list`; try -d help")
	flag.BoolVar(&pstate.flagDWARF, "dwarf", !Wasm, "generate DWARF symbols")
	flag.BoolVar(&pstate.Ctxt.Flag_locationlists, "dwarflocationlists", true, "add location lists to DWARF in optimized mode")
	flag.IntVar(&pstate.genDwarfInline, "gendwarfinl", 2, "generate DWARF inline info records")
	objabi.Flagcount("e", "no limit on number of errors reported", &pstate.Debug['e'])
	objabi.Flagcount("f", "debug stack frames", &pstate.Debug['f'])
	objabi.Flagcount("h", "halt on error", &pstate.Debug['h'])
	objabi.Flagcount("i", "debug line number stack", &pstate.Debug['i'])
	objabi.Flagfn1("importmap", "add `definition` of the form source=actual to import map", pstate.addImportMap)
	objabi.Flagfn1("importcfg", "read import configuration from `file`", pstate.readImportCfg)
	flag.StringVar(&pstate.flag_installsuffix, "installsuffix", "", "set pkg directory `suffix`")
	objabi.Flagcount("j", "debug runtime-initialized variables", &pstate.Debug['j'])
	objabi.Flagcount("l", "disable inlining", &pstate.Debug['l'])
	flag.StringVar(&pstate.linkobj, "linkobj", "", "write linker-specific object to `file`")
	objabi.Flagcount("live", "debug liveness analysis", &pstate.debuglive)
	objabi.Flagcount("m", "print optimization decisions", &pstate.Debug['m'])
	flag.BoolVar(&pstate.flag_msan, "msan", false, "build code compatible with C/C++ memory sanitizer")
	flag.BoolVar(&pstate.dolinkobj, "dolinkobj", true, "generate linker-specific objects; if false, some invalid code may compile")
	flag.BoolVar(&pstate.nolocalimports, "nolocalimports", false, "reject local (relative) imports")
	flag.StringVar(&pstate.outfile, "o", "", "write output to `file`")
	flag.StringVar(&pstate.myimportpath, "p", "", "set expected package import `path`")
	flag.BoolVar(&pstate.writearchive, "pack", false, "write to file.a instead of file.o")
	objabi.Flagcount("r", "debug generated wrappers", &pstate.Debug['r'])
	flag.BoolVar(&pstate.flag_race, "race", false, "enable race detector")
	objabi.Flagcount("s", "warn about composite literals that can be simplified", &pstate.Debug['s'])
	flag.StringVar(&pstate.pathPrefix, "trimpath", "", "remove `prefix` from recorded source file paths")
	flag.BoolVar(&pstate.safemode, "u", false, "reject unsafe code")
	flag.BoolVar(&pstate.Debug_vlog, "v", false, "increase debug verbosity")
	objabi.Flagcount("w", "debug type checking", &pstate.Debug['w'])
	flag.BoolVar(&pstate.use_writebarrier, "wb", true, "enable write barrier")
	var flag_shared bool
	var flag_dynlink bool
	if supportsDynlink(pstate.thearch.LinkArch.Arch) {
		flag.BoolVar(&flag_shared, "shared", false, "generate code that can be linked into a shared library")
		flag.BoolVar(&flag_dynlink, "dynlink", false, "support references to Go symbols defined in other shared libraries")
	}
	flag.StringVar(&pstate.cpuprofile, "cpuprofile", "", "write cpu profile to `file`")
	flag.StringVar(&pstate.memprofile, "memprofile", "", "write memory profile to `file`")
	flag.Int64Var(&pstate.memprofilerate, "memprofilerate", 0, "set runtime.MemProfileRate to `rate`")
	var goversion string
	flag.StringVar(&goversion, "goversion", "", "required version of the runtime")
	flag.StringVar(&pstate.traceprofile, "traceprofile", "", "write an execution trace to `file`")
	flag.StringVar(&pstate.blockprofile, "blockprofile", "", "write block profile to `file`")
	flag.StringVar(&pstate.mutexprofile, "mutexprofile", "", "write mutex profile to `file`")
	flag.StringVar(&pstate.benchfile, "bench", "", "append benchmark times to `file`")
	flag.BoolVar(&pstate.flagiexport, "iexport", true, "export indexed package data")
	objabi.Flagparse(pstate.usage)

	// Record flags that affect the build result. (And don't
	// record flags that don't, since that would cause spurious
	// changes in the binary.)
	pstate.recordFlags("B", "N", "l", "msan", "race", "shared", "dynlink", "dwarflocationlists")

	pstate.Ctxt.Flag_shared = flag_dynlink || flag_shared
	pstate.Ctxt.Flag_dynlink = flag_dynlink
	pstate.Ctxt.Flag_optimize = pstate.Debug['N'] == 0

	pstate.Ctxt.Debugasm = pstate.Debug_asm
	pstate.Ctxt.Debugvlog = pstate.Debug_vlog
	if pstate.flagDWARF {
		pstate.Ctxt.DebugInfo = pstate.debuginfo
		pstate.Ctxt.GenAbstractFunc = pstate.genAbstractFunc
		pstate.Ctxt.DwFixups = obj.NewDwarfFixupTable(pstate.Ctxt)
	} else {
		// turn off inline generation if no dwarf at all
		pstate.genDwarfInline = 0
		pstate.Ctxt.Flag_locationlists = false
	}

	if flag.NArg() < 1 && pstate.debugstr != "help" && pstate.debugstr != "ssa/help" {
		pstate.usage()
	}

	if goversion != "" && goversion != runtime.Version() {
		fmt.Printf("compile: version %q does not match go tool version %q\n", runtime.Version(), goversion)
		pstate.Exit(2)
	}

	pstate.thearch.LinkArch.Init(pstate.Ctxt)

	if pstate.outfile == "" {
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
		if pstate.writearchive {
			suffix = ".a"
		}
		pstate.outfile = p + suffix
	}

	pstate.startProfile()

	if pstate.flag_race && pstate.flag_msan {
		log.Fatal("cannot use both -race and -msan")
	}
	if pstate.ispkgin(pstate.omit_pkgs) {
		pstate.flag_race = false
		pstate.flag_msan = false
	}
	if pstate.flag_race {
		pstate.racepkg = pstate.types.NewPkg("runtime/race", "")
	}
	if pstate.flag_msan {
		pstate.msanpkg = pstate.types.NewPkg("runtime/msan", "")
	}
	if pstate.flag_race || pstate.flag_msan {
		pstate.instrumenting = true
	}

	if pstate.compiling_runtime && pstate.Debug['N'] != 0 {
		log.Fatal("cannot disable optimizations while compiling runtime")
	}
	if pstate.nBackendWorkers < 1 {
		log.Fatalf("-c must be at least 1, got %d", pstate.nBackendWorkers)
	}
	if pstate.nBackendWorkers > 1 && !pstate.concurrentBackendAllowed() {
		log.Fatalf("cannot use concurrent backend compilation with provided flags; invoked as %v", os.Args)
	}
	if pstate.Ctxt.Flag_locationlists && len(pstate.Ctxt.Arch.DWARFRegisters) == 0 {
		log.Fatalf("location lists requested but register mapping not available on %v", pstate.Ctxt.Arch.Name)
	}

	// parse -d argument
	if pstate.debugstr != "" {
	Split:
		for _, name := range strings.Split(pstate.debugstr, ",") {
			if name == "" {
				continue
			}
			// display help about the -d option itself and quit
			if name == "help" {
				fmt.Printf(debugHelpHeader)
				maxLen := len("ssa/help")
				for _, t := range pstate.debugtab {
					if len(t.name) > maxLen {
						maxLen = len(t.name)
					}
				}
				for _, t := range pstate.debugtab {
					fmt.Printf("\t%-*s\t%s\n", maxLen, t.name, t.help)
				}
				// ssa options have their own help
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
			for _, t := range pstate.debugtab {
				if t.name != name {
					continue
				}
				switch vp := t.val.(type) {
				case nil:
				// Ignore
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
			// special case for ssa for now
			if strings.HasPrefix(name, "ssa/") {
				// expect form ssa/phase/flag
				// e.g. -d=ssa/generic_cse/time
				// _ in phase name also matches space
				phase := name[4:]
				flag := "debug" // default flag is debug
				if i := strings.Index(phase, "/"); i >= 0 {
					flag = phase[i+1:]
					phase = phase[:i]
				}
				err := pstate.ssa.PhaseOption(phase, flag, val, valstring)
				if err != "" {
					log.Fatalf(err)
				}
				continue Split
			}
			log.Fatalf("unknown debug key -d %s\n", name)
		}
	}

	// set via a -d flag
	pstate.Ctxt.Debugpcln = pstate.Debug_pctab
	if pstate.flagDWARF {
		pstate.dwarf.EnableLogging(pstate.Debug_gendwarfinl != 0)
	}

	if pstate.Debug_softfloat != 0 {
		pstate.thearch.SoftFloat = true
	}

	// enable inlining.  for now:
	//	default: inlining on.  (debug['l'] == 1)
	//	-l: inlining off  (debug['l'] == 0)
	//	-l=2, -l=3: inlining on again, with extra debugging (debug['l'] > 1)
	if pstate.Debug['l'] <= 1 {
		pstate.Debug['l'] = 1 - pstate.Debug['l']
	}

	pstate.trackScopes = pstate.flagDWARF

	pstate.Widthptr = pstate.thearch.LinkArch.PtrSize
	pstate.Widthreg = pstate.thearch.LinkArch.RegSize

	// initialize types package
	// (we need to do this to break dependencies that otherwise
	// would lead to import cycles)
	pstate.types.Widthptr = pstate.Widthptr
	pstate.types.Dowidth = pstate.dowidth
	pstate.types.Fatalf = pstate.Fatalf
	pstate.types.Sconv = func(s *types.Sym, flag, mode int) string {
		return pstate.sconv(s, FmtFlag(flag), fmtMode(mode))
	}
	pstate.types.Tconv = func(t *types.Type, flag, mode, depth int) string {
		return pstate.tconv(t, FmtFlag(flag), fmtMode(mode), depth)
	}
	pstate.types.FormatSym = func(sym *types.Sym, s fmt.State, verb rune, mode int) {
		pstate.symFormat(sym, s, verb, fmtMode(mode))
	}
	pstate.types.FormatType = func(t *types.Type, s fmt.State, verb rune, mode int) {
		pstate.typeFormat(t, s, verb, fmtMode(mode))
	}
	pstate.types.TypeLinkSym = func(t *types.Type) *obj.LSym {
		return pstate.typenamesym(t).Linksym(pstate.types)
	}
	pstate.types.FmtLeft = int(FmtLeft)
	pstate.types.FmtUnsigned = int(FmtUnsigned)
	pstate.types.FErr = FErr
	pstate.types.Ctxt = pstate.Ctxt

	pstate.initUniverse()

	pstate.dclcontext = PEXTERN
	pstate.nerrors = 0

	pstate.autogeneratedPos = pstate.makePos(src.NewFileBase("<autogenerated>", "<autogenerated>"), 1, 0)

	pstate.timings.Start("fe", "loadsys")
	pstate.loadsys()

	pstate.timings.Start("fe", "parse")
	lines := pstate.parseFiles(flag.Args())
	pstate.timings.Stop()
	pstate.timings.AddEvent(int64(lines), "lines")

	pstate.finishUniverse()

	pstate.typecheckok = true
	if pstate.Debug['f'] != 0 {
		pstate.frame(1)
	}

	// Process top-level declarations in phases.

	// Phase 1: const, type, and names and types of funcs.
	//   This will gather all the information about types
	//   and methods but doesn't depend on any of it.
	pstate.defercheckwidth()

	// Don't use range--typecheck can add closures to xtop.
	pstate.timings.Start("fe", "typecheck", "top1")
	for i := 0; i < len(pstate.xtop); i++ {
		n := pstate.xtop[i]
		if op := n.Op; op != ODCL && op != OAS && op != OAS2 {
			pstate.xtop[i] = pstate.typecheck(n, Etop)
		}
	}

	// Phase 2: Variable assignments.
	//   To check interface assignments, depends on phase 1.

	// Don't use range--typecheck can add closures to xtop.
	pstate.timings.Start("fe", "typecheck", "top2")
	for i := 0; i < len(pstate.xtop); i++ {
		n := pstate.xtop[i]
		if op := n.Op; op == ODCL || op == OAS || op == OAS2 {
			pstate.xtop[i] = pstate.typecheck(n, Etop)
		}
	}
	pstate.resumecheckwidth()

	// Phase 3: Type check function bodies.
	// Don't use range--typecheck can add closures to xtop.
	pstate.timings.Start("fe", "typecheck", "func")
	var fcount int64
	for i := 0; i < len(pstate.xtop); i++ {
		n := pstate.xtop[i]
		if op := n.Op; op == ODCLFUNC || op == OCLOSURE {
			pstate.Curfn = n
			pstate.decldepth = 1
			pstate.saveerrors()
			pstate.typecheckslice(pstate.Curfn.Nbody.Slice(), Etop)
			pstate.checkreturn(pstate.Curfn)
			if pstate.nerrors != 0 {
				pstate.Curfn.Nbody.Set(nil) // type errors; do not compile
			}
			// Now that we've checked whether n terminates,
			// we can eliminate some obviously dead code.
			pstate.deadcode(pstate.Curfn)
			fcount++
		}
	}
	// With all types ckecked, it's now safe to verify map keys.
	pstate.checkMapKeys()
	pstate.timings.AddEvent(fcount, "funcs")

	if pstate.nsavederrors+pstate.nerrors != 0 {
		pstate.errorexit()
	}

	// Phase 4: Decide how to capture closed variables.
	// This needs to run before escape analysis,
	// because variables captured by value do not escape.
	pstate.timings.Start("fe", "capturevars")
	for _, n := range pstate.xtop {
		if n.Op == ODCLFUNC && n.Func.Closure != nil {
			pstate.Curfn = n
			pstate.capturevars(n)
		}
	}
	pstate.capturevarscomplete = true

	pstate.Curfn = nil

	if pstate.nsavederrors+pstate.nerrors != 0 {
		pstate.errorexit()
	}

	// Phase 5: Inlining
	pstate.timings.Start("fe", "inlining")
	if pstate.Debug_typecheckinl != 0 {
		// Typecheck imported function bodies if debug['l'] > 1,
		// otherwise lazily when used or re-exported.
		for _, n := range pstate.importlist {
			if n.Func.Inl != nil {
				pstate.saveerrors()
				pstate.typecheckinl(n)
			}
		}

		if pstate.nsavederrors+pstate.nerrors != 0 {
			pstate.errorexit()
		}
	}

	if pstate.Debug['l'] != 0 {
		// Find functions that can be inlined and clone them before walk expands them.
		pstate.visitBottomUp(pstate.xtop, func(list []*Node, recursive bool) {
			for _, n := range list {
				if !recursive {
					pstate.caninl(n)
				} else {
					if pstate.Debug['m'] > 1 {
						fmt.Printf("%v: cannot inline %v: recursive\n", n.Line(pstate), n.Func.Nname)
					}
				}
				pstate.inlcalls(n)
			}
		})
	}

	// Phase 6: Escape analysis.
	// Required for moving heap allocations onto stack,
	// which in turn is required by the closure implementation,
	// which stores the addresses of stack variables into the closure.
	// If the closure does not escape, it needs to be on the stack
	// or else the stack copier will not update it.
	// Large values are also moved off stack in escape analysis;
	// because large values may contain pointers, it must happen early.
	pstate.timings.Start("fe", "escapes")
	pstate.escapes(pstate.xtop)

	if pstate.dolinkobj {
		// Collect information for go:nowritebarrierrec
		// checking. This must happen before transformclosure.
		// We'll do the final check after write barriers are
		// inserted.
		if pstate.compiling_runtime {
			pstate.nowritebarrierrecCheck = pstate.newNowritebarrierrecChecker()
		}

		// Phase 7: Transform closure bodies to properly reference captured variables.
		// This needs to happen before walk, because closures must be transformed
		// before walk reaches a call of a closure.
		pstate.timings.Start("fe", "xclosures")
		for _, n := range pstate.xtop {
			if n.Op == ODCLFUNC && n.Func.Closure != nil {
				pstate.Curfn = n
				pstate.transformclosure(n)
			}
		}

		// Prepare for SSA compilation.
		// This must be before peekitabs, because peekitabs
		// can trigger function compilation.
		pstate.initssaconfig()

		// Just before compilation, compile itabs found on
		// the right side of OCONVIFACE so that methods
		// can be de-virtualized during compilation.
		pstate.Curfn = nil
		pstate.peekitabs()

		// Phase 8: Compile top level functions.
		// Don't use range--walk can add functions to xtop.
		pstate.timings.Start("be", "compilefuncs")
		fcount = 0
		for i := 0; i < len(pstate.xtop); i++ {
			n := pstate.xtop[i]
			if n.Op == ODCLFUNC {
				pstate.funccompile(n)
				fcount++
			}
		}
		pstate.timings.AddEvent(fcount, "funcs")

		if pstate.nsavederrors+pstate.nerrors == 0 {
			pstate.fninit(pstate.xtop)
		}

		pstate.compileFunctions()

		if pstate.nowritebarrierrecCheck != nil {
			// Write barriers are now known. Check the
			// call graph.
			pstate.nowritebarrierrecCheck.check(pstate)
			pstate.nowritebarrierrecCheck = nil
		}

		// Finalize DWARF inline routine DIEs, then explicitly turn off
		// DWARF inlining gen so as to avoid problems with generated
		// method wrappers.
		if pstate.Ctxt.DwFixups != nil {
			pstate.Ctxt.DwFixups.Finalize(pstate.myimportpath, pstate.Debug_gendwarfinl != 0)
			pstate.Ctxt.DwFixups = nil
			pstate.genDwarfInline = 0
		}
	}

	// Phase 9: Check external declarations.
	pstate.timings.Start("be", "externaldcls")
	for i, n := range pstate.externdcl {
		if n.Op == ONAME {
			pstate.externdcl[i] = pstate.typecheck(pstate.externdcl[i], Erv)
		}
	}

	if pstate.nerrors+pstate.nsavederrors != 0 {
		pstate.errorexit()
	}

	// Write object data to disk.
	pstate.timings.Start("be", "dumpobj")
	pstate.dumpobj()
	if pstate.asmhdr != "" {
		pstate.dumpasmhdr()
	}

	// Check whether any of the functions we have compiled have gigantic stack frames.
	obj.SortSlice(pstate.largeStackFrames, func(i, j int) bool {
		return pstate.largeStackFrames[i].Before(pstate.largeStackFrames[j])
	})
	for _, largePos := range pstate.largeStackFrames {
		pstate.yyerrorl(largePos, "stack frame too large (>1GB)")
	}

	if len(pstate.compilequeue) != 0 {
		pstate.Fatalf("%d uncompiled functions", len(pstate.compilequeue))
	}

	if pstate.nerrors+pstate.nsavederrors != 0 {
		pstate.errorexit()
	}

	pstate.flusherrors()
	pstate.timings.Stop()

	if pstate.benchfile != "" {
		if err := pstate.writebench(pstate.benchfile); err != nil {
			log.Fatalf("cannot write benchmark data: %v", err)
		}
	}
}

func (pstate *PackageState) writebench(filename string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	fmt.Fprintln(&buf, "commit:", pstate.objabi.Version)
	fmt.Fprintln(&buf, "goos:", runtime.GOOS)
	fmt.Fprintln(&buf, "goarch:", runtime.GOARCH)
	pstate.timings.Write(&buf, "BenchmarkCompile:"+pstate.myimportpath+":")

	n, err := f.Write(buf.Bytes())
	if err != nil {
		return err
	}
	if n != buf.Len() {
		panic("bad writer")
	}

	return f.Close()
}

func (pstate *PackageState) addImportMap(s string) {
	if strings.Count(s, "=") != 1 {
		log.Fatal("-importmap argument must be of the form source=actual")
	}
	i := strings.Index(s, "=")
	source, actual := s[:i], s[i+1:]
	if source == "" || actual == "" {
		log.Fatal("-importmap argument must be of the form source=actual; source and actual must be non-empty")
	}
	pstate.importMap[source] = actual
}

func (pstate *PackageState) readImportCfg(file string) {
	pstate.packageFile = map[string]string{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("-importcfg: %v", err)
	}

	for lineNum, line := range strings.Split(string(data), "\n") {
		lineNum++ // 1-based
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
			pstate.importMap[before] = after
		case "packagefile":
			if before == "" || after == "" {
				log.Fatalf("%s:%d: invalid packagefile: syntax is \"packagefile path=filename\"", file, lineNum)
			}
			pstate.packageFile[before] = after
		}
	}
}

func (pstate *PackageState) saveerrors() {
	pstate.nsavederrors += pstate.nerrors
	pstate.nerrors = 0
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

func (pstate *PackageState) addidir(dir string) {
	if dir != "" {
		pstate.idirs = append(pstate.idirs, dir)
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

func (pstate *PackageState) findpkg(name string) (file string, ok bool) {
	if islocalname(name) {
		if pstate.safemode || pstate.nolocalimports {
			return "", false
		}

		if pstate.packageFile != nil {
			file, ok = pstate.packageFile[name]
			return file, ok
		}

		// try .a before .6.  important for building libraries:
		// if there is an array.6 in the array.a library,
		// want to find all of array.a, not just array.6.
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

	// local imports should be canonicalized already.
	// don't want to see "encoding/../encoding/base64"
	// as different from "encoding/base64".
	if q := path.Clean(name); q != name {
		pstate.yyerror("non-canonical import path %q (should be %q)", name, q)
		return "", false
	}

	if pstate.packageFile != nil {
		file, ok = pstate.packageFile[name]
		return file, ok
	}

	for _, dir := range pstate.idirs {
		file = fmt.Sprintf("%s/%s.a", dir, name)
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
		file = fmt.Sprintf("%s/%s.o", dir, name)
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
	}

	if pstate.objabi.GOROOT != "" {
		suffix := ""
		suffixsep := ""
		if pstate.flag_installsuffix != "" {
			suffixsep = "_"
			suffix = pstate.flag_installsuffix
		} else if pstate.flag_race {
			suffixsep = "_"
			suffix = "race"
		} else if pstate.flag_msan {
			suffixsep = "_"
			suffix = "msan"
		}

		file = fmt.Sprintf("%s/pkg/%s_%s%s%s/%s.a", pstate.objabi.GOROOT, pstate.objabi.GOOS, pstate.objabi.GOARCH, suffixsep, suffix, name)
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
		file = fmt.Sprintf("%s/pkg/%s_%s%s%s/%s.o", pstate.objabi.GOROOT, pstate.objabi.GOOS, pstate.objabi.GOARCH, suffixsep, suffix, name)
		if _, err := os.Stat(file); err == nil {
			return file, true
		}
	}

	return "", false
}

// loadsys loads the definitions for the low-level runtime functions,
// so that the compiler can generate calls to them,
// but does not make them visible to user code.
func (pstate *PackageState) loadsys() {
	pstate.types.Block = 1

	pstate.inimport = true
	pstate.typecheckok = true
	pstate.defercheckwidth()

	typs := pstate.runtimeTypes()
	for _, d := range pstate.runtimeDecls {
		sym := pstate.Runtimepkg.Lookup(pstate.types, d.name)
		typ := typs[d.typ]
		switch d.tag {
		case funcTag:
			pstate.importfunc(pstate.Runtimepkg, pstate.src.NoXPos, sym, typ)
		case varTag:
			pstate.importvar(pstate.Runtimepkg, pstate.src.NoXPos, sym, typ)
		default:
			pstate.Fatalf("unhandled declaration tag %v", d.tag)
		}
	}

	pstate.typecheckok = false
	pstate.resumecheckwidth()
	pstate.inimport = false
}

func (pstate *PackageState) importfile(f *Val) *types.Pkg {
	path_, ok := f.U.(string)
	if !ok {
		pstate.yyerror("import path must be a string")
		return nil
	}

	if len(path_) == 0 {
		pstate.yyerror("import path is empty")
		return nil
	}

	if pstate.isbadimport(path_, false) {
		return nil
	}

	// The package name main is no longer reserved,
	// but we reserve the import path "main" to identify
	// the main package, just as we reserve the import
	// path "math" to identify the standard math package.
	if path_ == "main" {
		pstate.yyerror("cannot import \"main\"")
		pstate.errorexit()
	}

	if pstate.myimportpath != "" && path_ == pstate.myimportpath {
		pstate.yyerror("import %q while compiling that package (import cycle)", path_)
		pstate.errorexit()
	}

	if mapped, ok := pstate.importMap[path_]; ok {
		path_ = mapped
	}

	if path_ == "unsafe" {
		if pstate.safemode {
			pstate.yyerror("cannot import package unsafe")
			pstate.errorexit()
		}

		pstate.imported_unsafe = true
		return pstate.unsafepkg
	}

	if islocalname(path_) {
		if path_[0] == '/' {
			pstate.yyerror("import path cannot be absolute path")
			return nil
		}

		prefix := pstate.Ctxt.Pathname
		if pstate.localimport != "" {
			prefix = pstate.localimport
		}
		path_ = path.Join(prefix, path_)

		if pstate.isbadimport(path_, true) {
			return nil
		}
	}

	file, found := pstate.findpkg(path_)
	if !found {
		pstate.yyerror("can't find import: %q", path_)
		pstate.errorexit()
	}

	importpkg := pstate.types.NewPkg(path_, "")
	if importpkg.Imported {
		return importpkg
	}

	importpkg.Imported = true

	imp, err := bio.Open(file)
	if err != nil {
		pstate.yyerror("can't open import: %q: %v", path_, err)
		pstate.errorexit()
	}
	defer imp.Close()

	// check object header
	p, err := imp.ReadString('\n')
	if err != nil {
		pstate.yyerror("import %s: reading input: %v", file, err)
		pstate.errorexit()
	}

	if p == "!<arch>\n" { // package archive
		// package export block should be first
		sz := arsize(imp.Reader, "__.PKGDEF")
		if sz <= 0 {
			pstate.yyerror("import %s: not a package file", file)
			pstate.errorexit()
		}
		p, err = imp.ReadString('\n')
		if err != nil {
			pstate.yyerror("import %s: reading input: %v", file, err)
			pstate.errorexit()
		}
	}

	if !strings.HasPrefix(p, "go object ") {
		pstate.yyerror("import %s: not a go object file: %s", file, p)
		pstate.errorexit()
	}
	q := fmt.Sprintf("%s %s %s %s\n", pstate.objabi.GOOS, pstate.objabi.GOARCH, pstate.objabi.Version, pstate.objabi.Expstring())
	if p[10:] != q {
		pstate.yyerror("import %s: object is [%s] expected [%s]", file, p[10:], q)
		pstate.errorexit()
	}

	// process header lines
	safe := false
	for {
		p, err = imp.ReadString('\n')
		if err != nil {
			pstate.yyerror("import %s: reading input: %v", file, err)
			pstate.errorexit()
		}
		if p == "\n" {
			break // header ends with blank line
		}
		if strings.HasPrefix(p, "safe") {
			safe = true
			break // ok to ignore rest
		}
	}
	if pstate.safemode && !safe {
		pstate.yyerror("cannot import unsafe package %q", importpkg.Path)
	}

	// assume files move (get installed) so don't record the full path
	if pstate.packageFile != nil {
		// If using a packageFile map, assume path_ can be recorded directly.
		pstate.Ctxt.AddImport(path_)
	} else {
		// For file "/Users/foo/go/pkg/darwin_amd64/math.a" record "math.a".
		pstate.Ctxt.AddImport(file[len(file)-len(path_)-len(".a"):])
	}

	// In the importfile, if we find:
	// $$\n  (textual format): not supported anymore
	// $$B\n (binary format) : import directly, then feed the lexer a dummy statement

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

	// get character after $$
	if err == nil {
		c, _ = imp.ReadByte()
	}

	switch c {
	case '\n':
		pstate.yyerror("cannot import %s: old export format no longer supported (recompile library)", path_)
		return nil

	case 'B':
		if pstate.Debug_export != 0 {
			fmt.Printf("importing %s (%s)\n", path_, file)
		}
		imp.ReadByte() // skip \n after $$B

		c, err = imp.ReadByte()
		if err != nil {
			pstate.yyerror("import %s: reading input: %v", file, err)
			pstate.errorexit()
		}

		// New indexed format is distinguished by an 'i' byte,
		// whereas old export format always starts with 'c', 'd', or 'v'.
		if c == 'i' {
			if !pstate.flagiexport {
				pstate.yyerror("import %s: cannot import package compiled with -iexport=true", file)
				pstate.errorexit()
			}

			pstate.iimport(importpkg, imp)
		} else {
			if pstate.flagiexport {
				pstate.yyerror("import %s: cannot import package compiled with -iexport=false", file)
				pstate.errorexit()
			}

			imp.UnreadByte()
			pstate.Import(importpkg, imp.Reader)
		}

	default:
		pstate.yyerror("no import in %q", path_)
		pstate.errorexit()
	}

	if importpkg.Height >= pstate.myheight {
		pstate.myheight = importpkg.Height + 1
	}

	return importpkg
}

func (pstate *PackageState) pkgnotused(lineno src.XPos, path string, name string) {
	// If the package was imported with a name other than the final
	// import path element, show it explicitly in the error message.
	// Note that this handles both renamed imports and imports of
	// packages containing unconventional package declarations.
	// Note that this uses / always, even on Windows, because Go import
	// paths always use forward slashes.
	elem := path
	if i := strings.LastIndex(elem, "/"); i >= 0 {
		elem = elem[i+1:]
	}
	if name == "" || elem == name {
		pstate.yyerrorl(lineno, "imported and not used: %q", path)
	} else {
		pstate.yyerrorl(lineno, "imported and not used: %q as %s", path, name)
	}
}

func (pstate *PackageState) mkpackage(pkgname string) {
	if pstate.localpkg.Name == "" {
		if pkgname == "_" {
			pstate.yyerror("invalid package name _")
		}
		pstate.localpkg.Name = pkgname
	} else {
		if pkgname != pstate.localpkg.Name {
			pstate.yyerror("package %s; expected %s", pkgname, pstate.localpkg.Name)
		}
	}
}

func (pstate *PackageState) clearImports() {
	type importedPkg struct {
		pos  src.XPos
		path string
		name string
	}
	var unused []importedPkg

	for _, s := range pstate.localpkg.Syms {
		n := asNode(s.Def)
		if n == nil {
			continue
		}
		if n.Op == OPACK {
			// throw away top-level package name left over
			// from previous file.
			// leave s->block set to cause redeclaration
			// errors if a conflicting top-level name is
			// introduced by a different file.
			if !n.Name.Used() && pstate.nsyntaxerrors == 0 {
				unused = append(unused, importedPkg{n.Pos, n.Name.Pkg.Path, s.Name})
			}
			s.Def = nil
			continue
		}
		if IsAlias(s) {
			// throw away top-level name left over
			// from previous import . "x"
			if n.Name != nil && n.Name.Pack != nil && !n.Name.Pack.Name.Used() && pstate.nsyntaxerrors == 0 {
				unused = append(unused, importedPkg{n.Name.Pack.Pos, n.Name.Pack.Name.Pkg.Path, ""})
				n.Name.Pack.Name.SetUsed(true)
			}
			s.Def = nil
			continue
		}
	}

	obj.SortSlice(unused, func(i, j int) bool { return unused[i].pos.Before(unused[j].pos) })
	for _, pkg := range unused {
		pstate.pkgnotused(pkg.pos, pkg.path, pkg.name)
	}
}

func IsAlias(sym *types.Sym) bool {
	return sym.Def != nil && asNode(sym.Def).Sym != sym
}

func (pstate *PackageState) concurrentBackendAllowed() bool {
	for i, x := range pstate.Debug {
		if x != 0 && !pstate.concurrentFlagOK[i] {
			return false
		}
	}
	// Debug_asm by itself is ok, because all printing occurs
	// while writing the object file, and that is non-concurrent.
	// Adding Debug_vlog, however, causes Debug_asm to also print
	// while flushing the plist, which happens concurrently.
	if pstate.Debug_vlog || pstate.debugstr != "" || pstate.debuglive > 0 {
		return false
	}
	// TODO: Test and delete these conditions.
	if pstate.objabi.Fieldtrack_enabled != 0 || pstate.objabi.Clobberdead_enabled != 0 {
		return false
	}
	// TODO: fix races and enable the following flags
	if pstate.Ctxt.Flag_shared || pstate.Ctxt.Flag_dynlink || pstate.flag_race {
		return false
	}
	return true
}

// recordFlags records the specified command-line flags to be placed
// in the DWARF info.
func (pstate *PackageState) recordFlags(flags ...string) {
	if pstate.myimportpath == "" {
		// We can't record the flags if we don't know what the
		// package name is.
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
			// Flag has default value, so omit it.
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
	s := pstate.Ctxt.Lookup(dwarf.CUInfoPrefix + "producer." + pstate.myimportpath)
	s.Type = objabi.SDWARFINFO
	// Sometimes (for example when building tests) we can link
	// together two package main archives. So allow dups.
	s.Set(obj.AttrDuplicateOK, true)
	pstate.Ctxt.Data = append(pstate.Ctxt.Data, s)
	s.P = cmd.Bytes()[1:]
}

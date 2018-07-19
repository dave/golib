package ld

import (
	"bufio"
	"flag"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
)

// writes a "GUI binary" instead of a "console binary"

func (psess *PackageSession) init() {
	flag.Var(&psess.rpath, "r", "set the ELF dynamic linker search `path` to dir1:dir2:...")
}

// Flags used by the linker. The exported flags are used by the architecture-specific packages.

// use 64-bit addresses in symbol table

// Main is the main entry point for the linker code.
func (psess *PackageSession) Main(arch *sys.Arch, theArch Arch) {
	psess.
		thearch = theArch
	ctxt := psess.linknew(arch)
	ctxt.Bso = bufio.NewWriter(os.Stdout)

	for _, arg := range os.Args {
		if arg == "-crash_for_testing" {
			os.Exit(2)
		}
	}

	final := psess.gorootFinal()
	psess.
		addstrdata1(ctxt, "runtime/internal/sys.DefaultGoroot="+final)
	psess.
		addstrdata1(ctxt, "github.com/dave/golib/src/cmd/internal/objabi.defaultGOROOT="+final)

	if ctxt.Arch.Family == sys.AMD64 && psess.objabi.GOOS == "plan9" {
		flag.BoolVar(&psess.Flag8, "8", false, "use 64-bit addresses in symbol table")
	}
	flagHeadType := flag.String("H", "", "set header `type`")
	flag.BoolVar(&ctxt.linkShared, "linkshared", false, "link against installed Go shared libraries")
	flag.Var(&ctxt.LinkMode, "linkmode", "set link `mode`")
	flag.Var(&ctxt.BuildMode, "buildmode", "set build `mode`")
	objabi.Flagfn1("B", "add an ELF NT_GNU_BUILD_ID `note` when using ELF", psess.addbuildinfo)
	objabi.Flagfn1("L", "add specified `directory` to library path", func(a string) { Lflag(ctxt, a) })
	objabi.AddVersionFlag()
	objabi.Flagfn1("X", "add string value `definition` of the form importpath.name=value", func(s string) {
		psess.addstrdata1(ctxt, s)
	})
	objabi.Flagcount("v", "print link trace", &ctxt.Debugvlog)
	objabi.Flagfn1("importcfg", "read import configuration from `file`", ctxt.readImportCfg)

	objabi.Flagparse(psess.usage)

	switch *flagHeadType {
	case "":
	case "windowsgui":
		ctxt.HeadType = objabi.Hwindows
		psess.
			windowsgui = true
	default:
		if err := ctxt.HeadType.Set(*flagHeadType); err != nil {
			psess.
				Errorf(nil, "%v", err)
			psess.
				usage()
		}
	}
	psess.
		startProfile()
	if ctxt.BuildMode == BuildModeUnset {
		ctxt.BuildMode = BuildModeExe
	}

	if ctxt.BuildMode != BuildModeShared && flag.NArg() != 1 {
		psess.
			usage()
	}

	if *psess.flagOutfile == "" {
		*psess.flagOutfile = "a.out"
		if ctxt.HeadType == objabi.Hwindows {
			*psess.flagOutfile += ".exe"
		}
	}
	psess.
		interpreter = *psess.flagInterpreter
	psess.
		libinit(ctxt)

	if ctxt.HeadType == objabi.Hunknown {
		ctxt.HeadType.Set(psess.objabi.GOOS)
	}

	ctxt.computeTLSOffset(psess)
	psess.
		thearch.Archinit(ctxt)

	if ctxt.linkShared && !ctxt.IsELF {
		psess.
			Exitf("-linkshared can only be used on elf systems")
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("HEADER = -H%d -T0x%x -D0x%x -R0x%x\n", ctxt.HeadType, uint64(*psess.FlagTextAddr), uint64(*psess.FlagDataAddr), uint32(*psess.FlagRound))
	}

	switch ctxt.BuildMode {
	case BuildModeShared:
		for i := 0; i < flag.NArg(); i++ {
			arg := flag.Arg(i)
			parts := strings.SplitN(arg, "=", 2)
			var pkgpath, file string
			if len(parts) == 1 {
				pkgpath, file = "main", arg
			} else {
				pkgpath, file = parts[0], parts[1]
			}
			psess.
				pkglistfornote = append(psess.pkglistfornote, pkgpath...)
			psess.
				pkglistfornote = append(psess.pkglistfornote, '\n')
			psess.
				addlibpath(ctxt, "command line", "command line", file, pkgpath, "")
		}
	case BuildModePlugin:
		psess.
			addlibpath(ctxt, "command line", "command line", flag.Arg(0), *psess.flagPluginPath, "")
	default:
		psess.
			addlibpath(ctxt, "command line", "command line", flag.Arg(0), "main", "")
	}
	ctxt.loadlib(psess)

	ctxt.dostrdata(psess)
	psess.
		deadcode(ctxt)
	if psess.objabi.Fieldtrack_enabled != 0 {
		psess.
			fieldtrack(ctxt)
	}
	ctxt.callgraph(psess)

	ctxt.doelf(psess)
	if ctxt.HeadType == objabi.Hdarwin {
		ctxt.domacho(psess)
	}
	ctxt.dostkcheck(psess)
	if ctxt.HeadType == objabi.Hwindows {
		ctxt.dope(psess)
	}
	ctxt.addexport(psess)
	psess.
		thearch.Gentext(ctxt)
	ctxt.textbuildid(psess)
	ctxt.textaddress(psess)
	ctxt.pclntab(psess)
	ctxt.findfunctab(psess)
	ctxt.typelink()
	ctxt.symtab(psess)
	ctxt.dodata(psess)
	order := ctxt.address(psess)
	ctxt.reloc(psess)
	psess.
		dwarfcompress(ctxt)
	ctxt.layout(psess, order)
	psess.
		thearch.Asmb(ctxt)
	ctxt.undef(psess)
	ctxt.hostlink(psess)
	ctxt.archive(psess)
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f cpu time\n", psess.Cputime())
		ctxt.Logf("%d symbols\n", len(ctxt.Syms.Allsym))
		ctxt.Logf("%d liveness data\n", psess.liveness)
	}

	ctxt.Bso.Flush()
	psess.
		errorexit()
}

type Rpath struct {
	set bool
	val string
}

func (r *Rpath) Set(val string) error {
	r.set = true
	r.val = val
	return nil
}

func (r *Rpath) String() string {
	return r.val
}

func (psess *PackageSession) startProfile() {
	if *psess.cpuprofile != "" {
		f, err := os.Create(*psess.cpuprofile)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatalf("%v", err)
		}
		psess.
			AtExit(pprof.StopCPUProfile)
	}
	if *psess.memprofile != "" {
		if *psess.memprofilerate != 0 {
			runtime.MemProfileRate = int(*psess.memprofilerate)
		}
		f, err := os.Create(*psess.memprofile)
		if err != nil {
			log.Fatalf("%v", err)
		}
		psess.
			AtExit(func() {
				runtime.GC()
				if err := pprof.WriteHeapProfile(f); err != nil {
					log.Fatalf("%v", err)
				}
			})
	}
}

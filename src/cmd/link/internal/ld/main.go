// Inferno utils/6l/obj.c
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/6l/obj.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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

func (pstate *PackageState) init() {
	flag.Var(&pstate.rpath, "r", "set the ELF dynamic linker search `path` to dir1:dir2:...")
}

// Main is the main entry point for the linker code.
func (pstate *PackageState) Main(arch *sys.Arch, theArch Arch) {
	pstate.thearch = theArch
	ctxt := pstate.linknew(arch)
	ctxt.Bso = bufio.NewWriter(os.Stdout)

	// For testing behavior of go command when tools crash silently.
	// Undocumented, not in standard flag parser to avoid
	// exposing in usage message.
	for _, arg := range os.Args {
		if arg == "-crash_for_testing" {
			os.Exit(2)
		}
	}

	final := pstate.gorootFinal()
	pstate.addstrdata1(ctxt, "runtime/internal/sys.DefaultGoroot="+final)
	pstate.addstrdata1(ctxt, "github.com/dave/golib/src/cmd/internal/objabi.defaultGOROOT="+final)

	// TODO(matloob): define these above and then check flag values here
	if ctxt.Arch.Family == sys.AMD64 && pstate.objabi.GOOS == "plan9" {
		flag.BoolVar(&pstate.Flag8, "8", false, "use 64-bit addresses in symbol table")
	}
	flagHeadType := flag.String("H", "", "set header `type`")
	flag.BoolVar(&ctxt.linkShared, "linkshared", false, "link against installed Go shared libraries")
	flag.Var(&ctxt.LinkMode, "linkmode", "set link `mode`")
	flag.Var(&ctxt.BuildMode, "buildmode", "set build `mode`")
	objabi.Flagfn1("B", "add an ELF NT_GNU_BUILD_ID `note` when using ELF", pstate.addbuildinfo)
	objabi.Flagfn1("L", "add specified `directory` to library path", func(a string) { Lflag(ctxt, a) })
	objabi.AddVersionFlag() // -V
	objabi.Flagfn1("X", "add string value `definition` of the form importpath.name=value", func(s string) { pstate.addstrdata1(ctxt, s) })
	objabi.Flagcount("v", "print link trace", &ctxt.Debugvlog)
	objabi.Flagfn1("importcfg", "read import configuration from `file`", ctxt.readImportCfg)

	objabi.Flagparse(pstate.usage)

	switch *flagHeadType {
	case "":
	case "windowsgui":
		ctxt.HeadType = objabi.Hwindows
		pstate.windowsgui = true
	default:
		if err := ctxt.HeadType.Set(*flagHeadType); err != nil {
			pstate.Errorf(nil, "%v", err)
			pstate.usage()
		}
	}

	pstate.startProfile()
	if ctxt.BuildMode == BuildModeUnset {
		ctxt.BuildMode = BuildModeExe
	}

	if ctxt.BuildMode != BuildModeShared && flag.NArg() != 1 {
		pstate.usage()
	}

	if *pstate.flagOutfile == "" {
		*pstate.flagOutfile = "a.out"
		if ctxt.HeadType == objabi.Hwindows {
			*pstate.flagOutfile += ".exe"
		}
	}

	pstate.interpreter = *pstate.flagInterpreter

	pstate.libinit(ctxt) // creates outfile

	if ctxt.HeadType == objabi.Hunknown {
		ctxt.HeadType.Set(pstate.objabi.GOOS)
	}

	ctxt.computeTLSOffset(pstate)
	pstate.thearch.Archinit(ctxt)

	if ctxt.linkShared && !ctxt.IsELF {
		pstate.Exitf("-linkshared can only be used on elf systems")
	}

	if ctxt.Debugvlog != 0 {
		ctxt.Logf("HEADER = -H%d -T0x%x -D0x%x -R0x%x\n", ctxt.HeadType, uint64(*pstate.FlagTextAddr), uint64(*pstate.FlagDataAddr), uint32(*pstate.FlagRound))
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
			pstate.pkglistfornote = append(pstate.pkglistfornote, pkgpath...)
			pstate.pkglistfornote = append(pstate.pkglistfornote, '\n')
			pstate.addlibpath(ctxt, "command line", "command line", file, pkgpath, "")
		}
	case BuildModePlugin:
		pstate.addlibpath(ctxt, "command line", "command line", flag.Arg(0), *pstate.flagPluginPath, "")
	default:
		pstate.addlibpath(ctxt, "command line", "command line", flag.Arg(0), "main", "")
	}
	ctxt.loadlib(pstate)

	ctxt.dostrdata(pstate)
	pstate.deadcode(ctxt)
	if pstate.objabi.Fieldtrack_enabled != 0 {
		pstate.fieldtrack(ctxt)
	}
	ctxt.callgraph(pstate)

	ctxt.doelf(pstate)
	if ctxt.HeadType == objabi.Hdarwin {
		ctxt.domacho(pstate)
	}
	ctxt.dostkcheck(pstate)
	if ctxt.HeadType == objabi.Hwindows {
		ctxt.dope(pstate)
	}
	ctxt.addexport(pstate)
	pstate.thearch.Gentext(ctxt) // trampolines, call stubs, etc.
	ctxt.textbuildid(pstate)
	ctxt.textaddress(pstate)
	ctxt.pclntab(pstate)
	ctxt.findfunctab(pstate)
	ctxt.typelink()
	ctxt.symtab(pstate)
	ctxt.dodata(pstate)
	order := ctxt.address(pstate)
	ctxt.reloc(pstate)
	pstate.dwarfcompress(ctxt)
	ctxt.layout(pstate, order)
	pstate.thearch.Asmb(ctxt)
	ctxt.undef(pstate)
	ctxt.hostlink(pstate)
	ctxt.archive(pstate)
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f cpu time\n", pstate.Cputime())
		ctxt.Logf("%d symbols\n", len(ctxt.Syms.Allsym))
		ctxt.Logf("%d liveness data\n", pstate.liveness)
	}

	ctxt.Bso.Flush()

	pstate.errorexit()
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

func (pstate *PackageState) startProfile() {
	if *pstate.cpuprofile != "" {
		f, err := os.Create(*pstate.cpuprofile)
		if err != nil {
			log.Fatalf("%v", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatalf("%v", err)
		}
		pstate.AtExit(pprof.StopCPUProfile)
	}
	if *pstate.memprofile != "" {
		if *pstate.memprofilerate != 0 {
			runtime.MemProfileRate = int(*pstate.memprofilerate)
		}
		f, err := os.Create(*pstate.memprofile)
		if err != nil {
			log.Fatalf("%v", err)
		}
		pstate.AtExit(func() {
			runtime.GC() // profile all outstanding allocations
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatalf("%v", err)
			}
		})
	}
}

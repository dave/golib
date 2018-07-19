package gc

// Do not instrument the following packages at all,
// at best instrumentation would cause infinite recursion.

// Only insert racefuncenterfp/racefuncexit into the following packages.
// Memory accesses in the packages are either uninteresting or will cause false positives.

func (psess *PackageSession) ispkgin(pkgs []string) bool {
	if psess.myimportpath != "" {
		for _, p := range pkgs {
			if psess.myimportpath == p {
				return true
			}
		}
	}

	return false
}

func (psess *PackageSession) instrument(fn *Node) {
	if fn.Func.Pragma&Norace != 0 {
		return
	}

	if !psess.flag_race || !psess.ispkgin(psess.norace_inst_pkgs) {
		fn.Func.SetInstrumentBody(true)
	}

	if psess.flag_race {
		lno := psess.lineno
		psess.
			lineno = psess.src.NoXPos

		if psess.thearch.LinkArch.Arch == psess.sys.ArchPPC64LE {
			fn.Func.Enter.Prepend(psess.mkcall("racefuncenterfp", nil, nil))
			fn.Func.Exit.Append(psess.mkcall("racefuncexit", nil, nil))
		} else {

			nodpc := psess.nodfp.copy()
			nodpc.Type = psess.types.Types[TUINTPTR]
			nodpc.Xoffset = int64(-psess.Widthptr)
			fn.Func.Dcl = append(fn.Func.Dcl, nodpc)
			fn.Func.Enter.Prepend(psess.mkcall("racefuncenter", nil, nil, nodpc))
			fn.Func.Exit.Append(psess.mkcall("racefuncexit", nil, nil))
		}
		psess.
			lineno = lno
	}
}

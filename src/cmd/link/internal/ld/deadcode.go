package ld

import (
	"fmt"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/sys"
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"strings"
	"unicode"
)

// deadcode marks all reachable symbols.
//
// The basis of the dead code elimination is a flood fill of symbols,
// following their relocations, beginning at *flagEntrySymbol.
//
// This flood fill is wrapped in logic for pruning unused methods.
// All methods are mentioned by relocations on their receiver's *rtype.
// These relocations are specially defined as R_METHODOFF by the compiler
// so we can detect and manipulated them here.
//
// There are three ways a method of a reachable type can be invoked:
//
//	1. direct call
//	2. through a reachable interface type
//	3. reflect.Value.Call, .Method, or reflect.Method.Func
//
// The first case is handled by the flood fill, a directly called method
// is marked as reachable.
//
// The second case is handled by decomposing all reachable interface
// types into method signatures. Each encountered method is compared
// against the interface method signatures, if it matches it is marked
// as reachable. This is extremely conservative, but easy and correct.
//
// The third case is handled by looking to see if any of:
//	- reflect.Value.Call is reachable
//	- reflect.Value.Method is reachable
// 	- reflect.Type.Method or MethodByName is called.
// If any of these happen, all bets are off and all exported methods
// of reachable types are marked reachable.
//
// Any unreached text symbols are removed from ctxt.Textp.
func (psess *PackageSession) deadcode(ctxt *Link) {
	if ctxt.Debugvlog != 0 {
		ctxt.Logf("%5.2f deadcode\n", psess.Cputime())
	}

	d := &deadcodepass{
		ctxt:        ctxt,
		ifaceMethod: make(map[methodsig]bool),
	}

	d.init(psess)
	d.flood(psess)

	callSym := ctxt.Syms.ROLookup("reflect.Value.Call", 0)
	methSym := ctxt.Syms.ROLookup("reflect.Value.Method", 0)
	reflectSeen := false

	if ctxt.DynlinkingGo() {

		reflectSeen = true
	}

	for {
		if !reflectSeen {
			if d.reflectMethod || (callSym != nil && callSym.Attr.Reachable()) || (methSym != nil && methSym.Attr.Reachable()) {

				reflectSeen = true
			}
		}

		// Mark all methods that could satisfy a discovered
		// interface as reachable. We recheck old marked interfaces
		// as new types (with new methods) may have been discovered
		// in the last pass.
		var rem []methodref
		for _, m := range d.markableMethods {
			if (reflectSeen && m.isExported()) || d.ifaceMethod[m.m] {
				d.markMethod(psess, m)
			} else {
				rem = append(rem, m)
			}
		}
		d.markableMethods = rem

		if len(d.markQueue) == 0 {

			break
		}
		d.flood(psess)
	}

	for _, m := range d.markableMethods {
		for _, r := range m.r {
			d.cleanupReloc(r)
		}
	}

	if ctxt.BuildMode != BuildModeShared {

		for _, s := range ctxt.Syms.Allsym {
			if strings.HasPrefix(s.Name, "go.itablink.") {
				s.Attr.Set(sym.AttrReachable, len(s.R) == 1 && s.R[0].Sym.Attr.Reachable())
			}
		}
	}

	for _, lib := range ctxt.Library {
		lib.Textp = lib.Textp[:0]
	}

	textp := make([]*sym.Symbol, 0, len(ctxt.Textp))
	for _, s := range ctxt.Textp {
		if s.Attr.Reachable() {
			if s.Lib != nil {
				s.Lib.Textp = append(s.Lib.Textp, s)
			}
			textp = append(textp, s)
		}
	}
	ctxt.Textp = textp
}

// methodref holds the relocations from a receiver type symbol to its
// method. There are three relocations, one for each of the fields in
// the reflect.method struct: mtyp, ifn, and tfn.
type methodref struct {
	m   methodsig
	src *sym.Symbol   // receiver type symbol
	r   [3]*sym.Reloc // R_METHODOFF relocations to fields of runtime.method
}

func (m methodref) ifn() *sym.Symbol { return m.r[1].Sym }

func (m methodref) isExported() bool {
	for _, r := range m.m {
		return unicode.IsUpper(r)
	}
	panic("methodref has no signature")
}

// deadcodepass holds state for the deadcode flood fill.
type deadcodepass struct {
	ctxt            *Link
	markQueue       []*sym.Symbol      // symbols to flood fill in next pass
	ifaceMethod     map[methodsig]bool // methods declared in reached interfaces
	markableMethods []methodref        // methods of reached types
	reflectMethod   bool
}

func (d *deadcodepass) cleanupReloc(r *sym.Reloc) {
	if r.Sym.Attr.Reachable() {
		r.Type = objabi.R_ADDROFF
	} else {
		if d.ctxt.Debugvlog > 1 {
			d.ctxt.Logf("removing method %s\n", r.Sym.Name)
		}
		r.Sym = nil
		r.Siz = 0
	}
}

// mark appends a symbol to the mark queue for flood filling.
func (d *deadcodepass) mark(psess *PackageSession, s, parent *sym.Symbol) {
	if s == nil || s.Attr.Reachable() {
		return
	}
	if s.Attr.ReflectMethod() {
		d.reflectMethod = true
	}
	if *psess.flagDumpDep {
		p := "_"
		if parent != nil {
			p = parent.Name
		}
		fmt.Printf("%s -> %s\n", p, s.Name)
	}
	s.Attr |= sym.AttrReachable
	s.Reachparent = parent
	d.markQueue = append(d.markQueue, s)
}

// markMethod marks a method as reachable.
func (d *deadcodepass) markMethod(psess *PackageSession, m methodref) {
	for _, r := range m.r {
		d.mark(psess, r.Sym, m.src)
		r.Type = objabi.R_ADDROFF
	}
}

// init marks all initial symbols as reachable.
// In a typical binary, this is *flagEntrySymbol.
func (d *deadcodepass) init(psess *PackageSession) {
	var names []string

	if d.ctxt.Arch.Family == sys.ARM {

		names = append(names, "runtime.read_tls_fallback")
	}

	if d.ctxt.BuildMode == BuildModeShared {

		for _, s := range d.ctxt.Syms.Allsym {
			if s.Type != 0 && s.Type != sym.SDYNIMPORT {
				d.mark(psess, s, nil)
			}
		}
	} else {

		if d.ctxt.linkShared && (d.ctxt.BuildMode == BuildModeExe || d.ctxt.BuildMode == BuildModePIE) {
			names = append(names, "main.main", "main.init")
		} else {

			if d.ctxt.LinkMode == LinkExternal && (d.ctxt.BuildMode == BuildModeExe || d.ctxt.BuildMode == BuildModePIE) {
				if d.ctxt.HeadType == objabi.Hwindows && d.ctxt.Arch.Family == sys.I386 {
					*psess.flagEntrySymbol = "_main"
				} else {
					*psess.flagEntrySymbol = "main"
				}
			}
			names = append(names, *psess.flagEntrySymbol)
			if d.ctxt.BuildMode == BuildModePlugin {
				names = append(names, objabi.PathToPrefix(*psess.flagPluginPath)+".init", objabi.PathToPrefix(*psess.flagPluginPath)+".main", "go.plugin.tabs")

				exports := d.ctxt.Syms.ROLookup("go.plugin.exports", 0)
				if exports != nil {
					for _, r := range exports.R {
						d.mark(psess, r.Sym, nil)
					}
				}
			}
		}
		for _, s := range psess.dynexp {
			d.mark(psess, s, nil)
		}
	}

	for _, name := range names {
		d.mark(psess, d.ctxt.Syms.ROLookup(name, 0), nil)
	}
}

// flood fills symbols reachable from the markQueue symbols.
// As it goes, it collects methodref and interface method declarations.
func (d *deadcodepass) flood(psess *PackageSession) {
	for len(d.markQueue) > 0 {
		s := d.markQueue[0]
		d.markQueue = d.markQueue[1:]
		if s.Type == sym.STEXT {
			if d.ctxt.Debugvlog > 1 {
				d.ctxt.Logf("marktext %s\n", s.Name)
			}
			if s.FuncInfo != nil {
				for _, a := range s.FuncInfo.Autom {
					d.mark(psess, a.Gotype, s)
				}
			}

		}

		if strings.HasPrefix(s.Name, "type.") && s.Name[5] != '.' {
			if len(s.P) == 0 {

				continue
			}
			if decodetypeKind(d.ctxt.Arch, s)&kindMask == kindInterface {
				for _, sig := range psess.decodeIfaceMethods(d.ctxt.Arch, s) {
					if d.ctxt.Debugvlog > 1 {
						d.ctxt.Logf("reached iface method: %s\n", sig)
					}
					d.ifaceMethod[sig] = true
				}
			}
		}

		mpos := 0
		var methods []methodref
		for i := range s.R {
			r := &s.R[i]
			if r.Sym == nil {
				continue
			}
			if r.Type == objabi.R_WEAKADDROFF {

				continue
			}
			if r.Type != objabi.R_METHODOFF {
				d.mark(psess, r.Sym, s)
				continue
			}

			if mpos == 0 {
				m := methodref{src: s}
				m.r[0] = r
				methods = append(methods, m)
			} else {
				methods[len(methods)-1].r[mpos] = r
			}
			mpos++
			if mpos == len(methodref{}.r) {
				mpos = 0
			}
		}
		if len(methods) > 0 {

			methodsigs := psess.decodetypeMethods(d.ctxt.Arch, s)
			if len(methods) != len(methodsigs) {
				panic(fmt.Sprintf("%q has %d method relocations for %d methods", s.Name, len(methods), len(methodsigs)))
			}
			for i, m := range methodsigs {
				name := string(m)
				name = name[:strings.Index(name, "(")]
				if !strings.HasSuffix(methods[i].ifn().Name, name) {
					panic(fmt.Sprintf("%q relocation for %q does not match method %q", s.Name, methods[i].ifn().Name, name))
				}
				methods[i].m = m
			}
			d.markableMethods = append(d.markableMethods, methods...)
		}

		if s.FuncInfo != nil {
			for i := range s.FuncInfo.Funcdata {
				d.mark(psess, s.FuncInfo.Funcdata[i], s)
			}
		}
		d.mark(psess, s.Gotype, s)
		d.mark(psess, s.Sub, s)
		d.mark(psess, s.Outer, s)
	}
}

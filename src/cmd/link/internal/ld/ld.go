package ld

import (
	"github.com/dave/golib/src/cmd/link/internal/sym"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func (ctxt *Link) readImportCfg(file string) {
	ctxt.PackageFile = make(map[string]string)
	ctxt.PackageShlib = make(map[string]string)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("-importcfg: %v", err)
	}

	for lineNum, line := range strings.Split(string(data), "\n") {
		lineNum++
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
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
		case "packagefile":
			if before == "" || after == "" {
				log.Fatalf("%s:%d: invalid packagefile: syntax is \"packagefile path=filename\"", file, lineNum)
			}
			ctxt.PackageFile[before] = after
		case "packageshlib":
			if before == "" || after == "" {
				log.Fatalf("%s:%d: invalid packageshlib: syntax is \"packageshlib path=filename\"", file, lineNum)
			}
			ctxt.PackageShlib[before] = after
		}
	}
}

func pkgname(ctxt *Link, lib string) string {
	name := path.Clean(lib)

	if ctxt.PackageFile != nil {
		return name
	}

	pkg := name
	if len(pkg) >= 2 && pkg[len(pkg)-2] == '.' {
		pkg = pkg[:len(pkg)-2]
	}
	return pkg
}

func findlib(ctxt *Link, lib string) (string, bool) {
	name := path.Clean(lib)

	var pname string
	isshlib := false

	if ctxt.linkShared && ctxt.PackageShlib[name] != "" {
		pname = ctxt.PackageShlib[name]
		isshlib = true
	} else if ctxt.PackageFile != nil {
		pname = ctxt.PackageFile[name]
		if pname == "" {
			ctxt.Logf("cannot find package %s (using -importcfg)\n", name)
			return "", false
		}
	} else {
		if filepath.IsAbs(name) {
			pname = name
		} else {
			pkg := pkgname(ctxt, lib)

			if !strings.HasSuffix(name, ".a") && !strings.HasSuffix(name, ".o") {
				name += ".a"
			}

			for _, dir := range ctxt.Libdir {
				if ctxt.linkShared {
					pname = dir + "/" + pkg + ".shlibname"
					if _, err := os.Stat(pname); err == nil {
						isshlib = true
						break
					}
				}
				pname = dir + "/" + name
				if _, err := os.Stat(pname); err == nil {
					break
				}
			}
		}
		pname = path.Clean(pname)
	}

	return pname, isshlib
}

func (psess *PackageSession) addlib(ctxt *Link, src string, obj string, lib string) *sym.Library {
	pkg := pkgname(ctxt, lib)

	if l := ctxt.LibraryByPkg[pkg]; l != nil {
		return l
	}

	pname, isshlib := findlib(ctxt, lib)

	if ctxt.Debugvlog > 1 {
		ctxt.Logf("%5.2f addlib: %s %s pulls in %s isshlib %v\n", psess.elapsed(), obj, src, pname, isshlib)
	}

	if isshlib {
		return psess.addlibpath(ctxt, src, obj, "", pkg, pname)
	}
	return psess.addlibpath(ctxt, src, obj, pname, pkg, "")
}

/*
 * add library to library list, return added library.
 *	srcref: src file referring to package
 *	objref: object file referring to package
 *	file: object file, e.g., /home/rsc/go/pkg/container/vector.a
 *	pkg: package import path, e.g. container/vector
 *	shlib: path to shared library, or .shlibname file holding path
 */
func (psess *PackageSession) addlibpath(ctxt *Link, srcref string, objref string, file string, pkg string, shlib string) *sym.Library {
	if l := ctxt.LibraryByPkg[pkg]; l != nil {
		return l
	}

	if ctxt.Debugvlog > 1 {
		ctxt.Logf("%5.2f addlibpath: srcref: %s objref: %s file: %s pkg: %s shlib: %s\n", psess.Cputime(), srcref, objref, file, pkg, shlib)
	}

	l := &sym.Library{}
	ctxt.LibraryByPkg[pkg] = l
	ctxt.Library = append(ctxt.Library, l)
	l.Objref = objref
	l.Srcref = srcref
	l.File = file
	l.Pkg = pkg
	if shlib != "" {
		if strings.HasSuffix(shlib, ".shlibname") {
			data, err := ioutil.ReadFile(shlib)
			if err != nil {
				psess.
					Errorf(nil, "cannot read %s: %v", shlib, err)
			}
			shlib = strings.TrimSpace(string(data))
		}
		l.Shlib = shlib
	}
	return l
}

func atolwhex(s string) int64 {
	n, _ := strconv.ParseInt(s, 0, 64)
	return n
}

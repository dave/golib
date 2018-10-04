// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/bio"
	"github.com/dave/golib/src/cmd/internal/obj"
	"github.com/dave/golib/src/cmd/internal/objabi"
	"github.com/dave/golib/src/cmd/internal/src"
	"io"
	"strconv"
)

// architecture-independent object file output
const ArhdrSize = 60

func formathdr(arhdr []byte, name string, size int64) {
	copy(arhdr[:], fmt.Sprintf("%-16s%-12d%-6d%-6d%-8o%-10d`\n", name, 0, 0, 0, 0644, size))
}

// These modes say which kind of object file to generate.
// The default use of the toolchain is to set both bits,
// generating a combined compiler+linker object, one that
// serves to describe the package to both the compiler and the linker.
// In fact the compiler and linker read nearly disjoint sections of
// that file, though, so in a distributed build setting it can be more
// efficient to split the output into two files, supplying the compiler
// object only to future compilations and the linker object only to
// future links.
//
// By default a combined object is written, but if -linkobj is specified
// on the command line then the default -o output is a compiler object
// and the -linkobj output is a linker object.
const (
	modeCompilerObj = 1 << iota
	modeLinkerObj
)

func (pstate *PackageState) dumpobj() {
	if !pstate.dolinkobj {
		pstate.dumpobj1(pstate.outfile, modeCompilerObj)
		return
	}
	if pstate.linkobj == "" {
		pstate.dumpobj1(pstate.outfile, modeCompilerObj|modeLinkerObj)
		return
	}
	pstate.dumpobj1(pstate.outfile, modeCompilerObj)
	pstate.dumpobj1(pstate.linkobj, modeLinkerObj)
}

func (pstate *PackageState) dumpobj1(outfile string, mode int) {
	bout, err := bio.Create(outfile)
	if err != nil {
		pstate.flusherrors()
		fmt.Printf("can't create %s: %v\n", outfile, err)
		pstate.errorexit()
	}
	defer bout.Close()
	bout.WriteString("!<arch>\n")

	if mode&modeCompilerObj != 0 {
		start := startArchiveEntry(bout)
		pstate.dumpCompilerObj(bout)
		finishArchiveEntry(bout, start, "__.PKGDEF")
	}
	if mode&modeLinkerObj != 0 {
		start := startArchiveEntry(bout)
		pstate.dumpLinkerObj(bout)
		finishArchiveEntry(bout, start, "_go_.o")
	}
}

func (pstate *PackageState) printObjHeader(bout *bio.Writer) {
	fmt.Fprintf(bout, "go object %s %s %s %s\n", pstate.objabi.GOOS, pstate.objabi.GOARCH, pstate.objabi.Version, pstate.objabi.Expstring())
	if pstate.buildid != "" {
		fmt.Fprintf(bout, "build id %q\n", pstate.buildid)
	}
	if pstate.localpkg.Name == "main" {
		fmt.Fprintf(bout, "main\n")
	}
	if pstate.safemode {
		fmt.Fprintf(bout, "safe\n")
	} else {
		fmt.Fprintf(bout, "----\n") // room for some other tool to write "safe"
	}
	fmt.Fprintf(bout, "\n") // header ends with blank line
}

func startArchiveEntry(bout *bio.Writer) int64 {
	var arhdr [ArhdrSize]byte
	bout.Write(arhdr[:])
	return bout.Offset()
}

func finishArchiveEntry(bout *bio.Writer, start int64, name string) {
	bout.Flush()
	size := bout.Offset() - start
	if size&1 != 0 {
		bout.WriteByte(0)
	}
	bout.Seek(start-ArhdrSize, 0)

	var arhdr [ArhdrSize]byte
	formathdr(arhdr[:], name, size)
	bout.Write(arhdr[:])
	bout.Flush()
	bout.Seek(start+size+(size&1), 0)
}

func (pstate *PackageState) dumpCompilerObj(bout *bio.Writer) {
	pstate.printObjHeader(bout)
	pstate.dumpexport(bout)
}

func (pstate *PackageState) dumpLinkerObj(bout *bio.Writer) {
	pstate.printObjHeader(bout)

	if len(pstate.pragcgobuf) != 0 {
		// write empty export section; must be before cgo section
		fmt.Fprintf(bout, "\n$$\n\n$$\n\n")
		fmt.Fprintf(bout, "\n$$  // cgo\n")
		if err := json.NewEncoder(bout).Encode(pstate.pragcgobuf); err != nil {
			pstate.Fatalf("serializing pragcgobuf: %v", err)
		}
		fmt.Fprintf(bout, "\n$$\n\n")
	}

	fmt.Fprintf(bout, "\n!\n")

	externs := len(pstate.externdcl)

	pstate.dumpglobls()
	pstate.addptabs()
	pstate.addsignats(pstate.externdcl)
	pstate.dumpsignats()
	pstate.dumptabs()
	pstate.dumpimportstrings()
	pstate.dumpbasictypes()

	// Calls to dumpsignats can generate functions,
	// like method wrappers and hash and equality routines.
	// Compile any generated functions, process any new resulting types, repeat.
	// This can't loop forever, because there is no way to generate an infinite
	// number of types in a finite amount of code.
	// In the typical case, we loop 0 or 1 times.
	// It was not until issue 24761 that we found any code that required a loop at all.
	for len(pstate.compilequeue) > 0 {
		pstate.compileFunctions()
		pstate.dumpsignats()
	}

	// Dump extra globals.
	tmp := pstate.externdcl

	if pstate.externdcl != nil {
		pstate.externdcl = pstate.externdcl[externs:]
	}
	pstate.dumpglobls()
	pstate.externdcl = tmp

	if pstate.zerosize > 0 {
		zero := pstate.mappkg.Lookup(pstate.types, "zero")
		pstate.ggloblsym(zero.Linksym(pstate.types), int32(pstate.zerosize), obj.DUPOK|obj.RODATA)
	}

	pstate.addGCLocals()

	pstate.obj.WriteObjFile(pstate.Ctxt, bout.Writer)
}

func (pstate *PackageState) addptabs() {
	if !pstate.Ctxt.Flag_dynlink || pstate.localpkg.Name != "main" {
		return
	}
	for _, exportn := range pstate.exportlist {
		s := exportn.Sym
		n := asNode(s.Def)
		if n == nil {
			continue
		}
		if n.Op != ONAME {
			continue
		}
		if !types.IsExported(s.Name) {
			continue
		}
		if s.Pkg.Name != "main" {
			continue
		}
		if n.Type.Etype == TFUNC && n.Class() == PFUNC {
			// function
			pstate.ptabs = append(pstate.ptabs, ptabEntry{s: s, t: asNode(s.Def).Type})
		} else {
			// variable
			pstate.ptabs = append(pstate.ptabs, ptabEntry{s: s, t: pstate.types.NewPtr(asNode(s.Def).Type)})
		}
	}
}

func (pstate *PackageState) dumpGlobal(n *Node) {
	if n.Type == nil {
		pstate.Fatalf("external %v nil type\n", n)
	}
	if n.Class() == PFUNC {
		return
	}
	if n.Sym.Pkg != pstate.localpkg {
		return
	}
	pstate.dowidth(n.Type)
	pstate.ggloblnod(n)
}

func (pstate *PackageState) dumpGlobalConst(n *Node) {
	// only export typed constants
	t := n.Type
	if t == nil {
		return
	}
	if n.Sym.Pkg != pstate.localpkg {
		return
	}
	// only export integer constants for now
	switch t.Etype {
	case TINT8:
	case TINT16:
	case TINT32:
	case TINT64:
	case TINT:
	case TUINT8:
	case TUINT16:
	case TUINT32:
	case TUINT64:
	case TUINT:
	case TUINTPTR:
	// ok
	case TIDEAL:
		if !pstate.Isconst(n, CTINT) {
			return
		}
		x := n.Val().U.(*Mpint)
		if x.Cmp(pstate.minintval[TINT]) < 0 || x.Cmp(pstate.maxintval[TINT]) > 0 {
			return
		}
		// Ideal integers we export as int (if they fit).
		t = pstate.types.Types[TINT]
	default:
		return
	}
	pstate.Ctxt.DwarfIntConst(pstate.obj, pstate.myimportpath, n.Sym.Name, pstate.typesymname(t), n.Int64(pstate))
}

func (pstate *PackageState) dumpglobls() {
	// add globals
	for _, n := range pstate.externdcl {
		switch n.Op {
		case ONAME:
			pstate.dumpGlobal(n)
		case OLITERAL:
			pstate.dumpGlobalConst(n)
		}
	}

	obj.SortSlice(pstate.funcsyms, func(i, j int) bool {
		return pstate.funcsyms[i].LinksymName() < pstate.funcsyms[j].LinksymName()
	})
	for _, s := range pstate.funcsyms {
		sf := s.Pkg.Lookup(pstate.types, funcsymname(s)).Linksym(pstate.types)
		pstate.dsymptr(sf, 0, s.Linksym(pstate.types), 0)
		pstate.ggloblsym(sf, int32(pstate.Widthptr), obj.DUPOK|obj.RODATA)
	}

	// Do not reprocess funcsyms on next dumpglobls call.
	pstate.funcsyms = nil
}

// addGCLocals adds gcargs and gclocals symbols to Ctxt.Data.
// It takes care not to add any duplicates.
// Though the object file format handles duplicates efficiently,
// storing only a single copy of the data,
// failure to remove these duplicates adds a few percent to object file size.
func (pstate *PackageState) addGCLocals() {
	seen := make(map[string]bool)
	for _, s := range pstate.Ctxt.Text {
		if s.Func == nil {
			continue
		}
		for _, gcsym := range []*obj.LSym{&s.Func.GCArgs, &s.Func.GCLocals, &s.Func.GCRegs} {
			if seen[gcsym.Name] {
				continue
			}
			pstate.Ctxt.Data = append(pstate.Ctxt.Data, gcsym)
			seen[gcsym.Name] = true
		}
	}
}

func (pstate *PackageState) duintxx(s *obj.LSym, off int, v uint64, wid int) int {
	if off&(wid-1) != 0 {
		pstate.Fatalf("duintxxLSym: misaligned: v=%d wid=%d off=%d", v, wid, off)
	}
	s.WriteInt(pstate.Ctxt, int64(off), wid, int64(v))
	return off + wid
}

func (pstate *PackageState) duint8(s *obj.LSym, off int, v uint8) int {
	return pstate.duintxx(s, off, uint64(v), 1)
}

func (pstate *PackageState) duint16(s *obj.LSym, off int, v uint16) int {
	return pstate.duintxx(s, off, uint64(v), 2)
}

func (pstate *PackageState) duint32(s *obj.LSym, off int, v uint32) int {
	return pstate.duintxx(s, off, uint64(v), 4)
}

func (pstate *PackageState) duintptr(s *obj.LSym, off int, v uint64) int {
	return pstate.duintxx(s, off, v, pstate.Widthptr)
}

func (pstate *PackageState) dbvec(s *obj.LSym, off int, bv bvec) int {
	// Runtime reads the bitmaps as byte arrays. Oblige.
	for j := 0; int32(j) < bv.n; j += 8 {
		word := bv.b[j/32]
		off = pstate.duint8(s, off, uint8(word>>(uint(j)%32)))
	}
	return off
}

func (pstate *PackageState) stringsym(pos src.XPos, s string) (data *obj.LSym) {
	var symname string
	if len(s) > 100 {
		// Huge strings are hashed to avoid long names in object files.
		// Indulge in some paranoia by writing the length of s, too,
		// as protection against length extension attacks.
		h := sha256.New()
		io.WriteString(h, s)
		symname = fmt.Sprintf(".gostring.%d.%x", len(s), h.Sum(nil))
	} else {
		// Small strings get named directly by their contents.
		symname = strconv.Quote(s)
	}

	const prefix = "go.string."
	symdataname := prefix + symname

	symdata := pstate.Ctxt.Lookup(symdataname)

	if !symdata.SeenGlobl() {
		// string data
		off := pstate.dsname(symdata, 0, s, pos, "string")
		pstate.ggloblsym(symdata, int32(off), obj.DUPOK|obj.RODATA|obj.LOCAL)
	}

	return symdata
}

func (pstate *PackageState) slicebytes(nam *Node, s string, len int) {
	pstate.slicebytes_gen++
	symname := fmt.Sprintf(".gobytes.%d", pstate.slicebytes_gen)
	sym := pstate.localpkg.Lookup(pstate.types, symname)
	sym.Def = asTypesNode(pstate.newname(sym))

	lsym := sym.Linksym(pstate.types)
	off := pstate.dsname(lsym, 0, s, nam.Pos, "slice")
	pstate.ggloblsym(lsym, int32(off), obj.NOPTR|obj.LOCAL)

	if nam.Op != ONAME {
		pstate.Fatalf("slicebytes %v", nam)
	}
	nsym := nam.Sym.Linksym(pstate.types)
	off = int(nam.Xoffset)
	off = pstate.dsymptr(nsym, off, lsym, 0)
	off = pstate.duintptr(nsym, off, uint64(len))
	pstate.duintptr(nsym, off, uint64(len))
}

func (pstate *PackageState) dsname(s *obj.LSym, off int, t string, pos src.XPos, what string) int {
	// Objects that are too large will cause the data section to overflow right away,
	// causing a cryptic error message by the linker. Check for oversize objects here
	// and provide a useful error message instead.
	if int64(len(t)) > 2e9 {
		pstate.yyerrorl(pos, "%v with length %v is too big", what, len(t))
		return 0
	}

	s.WriteString(pstate.Ctxt, int64(off), len(t), t)
	return off + len(t)
}

func (pstate *PackageState) dsymptr(s *obj.LSym, off int, x *obj.LSym, xoff int) int {
	off = int(pstate.Rnd(int64(off), int64(pstate.Widthptr)))
	s.WriteAddr(pstate.Ctxt, int64(off), pstate.Widthptr, x, int64(xoff))
	off += pstate.Widthptr
	return off
}

func (pstate *PackageState) dsymptrOff(s *obj.LSym, off int, x *obj.LSym) int {
	s.WriteOff(pstate.Ctxt, int64(off), x, 0)
	off += 4
	return off
}

func (pstate *PackageState) dsymptrWeakOff(s *obj.LSym, off int, x *obj.LSym) int {
	s.WriteWeakOff(pstate.Ctxt, int64(off), x, 0)
	off += 4
	return off
}

func (pstate *PackageState) gdata(nam *Node, nr *Node, wid int) {
	if nam.Op != ONAME {
		pstate.Fatalf("gdata nam op %v", nam.Op)
	}
	if nam.Sym == nil {
		pstate.Fatalf("gdata nil nam sym")
	}
	s := nam.Sym.Linksym(pstate.types)

	switch nr.Op {
	case OLITERAL:
		switch u := nr.Val().U.(type) {
		case bool:
			i := int64(obj.Bool2int(u))
			s.WriteInt(pstate.Ctxt, nam.Xoffset, wid, i)

		case *Mpint:
			s.WriteInt(pstate.Ctxt, nam.Xoffset, wid, u.Int64(pstate))

		case *Mpflt:
			f := u.Float64(pstate)
			switch nam.Type.Etype {
			case TFLOAT32:
				s.WriteFloat32(pstate.Ctxt, nam.Xoffset, float32(f))
			case TFLOAT64:
				s.WriteFloat64(pstate.Ctxt, nam.Xoffset, f)
			}

		case *Mpcplx:
			r := u.Real.Float64(pstate)
			i := u.Imag.Float64(pstate)
			switch nam.Type.Etype {
			case TCOMPLEX64:
				s.WriteFloat32(pstate.Ctxt, nam.Xoffset, float32(r))
				s.WriteFloat32(pstate.Ctxt, nam.Xoffset+4, float32(i))
			case TCOMPLEX128:
				s.WriteFloat64(pstate.Ctxt, nam.Xoffset, r)
				s.WriteFloat64(pstate.Ctxt, nam.Xoffset+8, i)
			}

		case string:
			symdata := pstate.stringsym(nam.Pos, u)
			s.WriteAddr(pstate.Ctxt, nam.Xoffset, pstate.Widthptr, symdata, 0)
			s.WriteInt(pstate.Ctxt, nam.Xoffset+int64(pstate.Widthptr), pstate.Widthptr, int64(len(u)))

		default:
			pstate.Fatalf("gdata unhandled OLITERAL %v", nr)
		}

	case OADDR:
		if nr.Left.Op != ONAME {
			pstate.Fatalf("gdata ADDR left op %v", nr.Left.Op)
		}
		to := nr.Left
		s.WriteAddr(pstate.Ctxt, nam.Xoffset, wid, to.Sym.Linksym(pstate.types), to.Xoffset)

	case ONAME:
		if nr.Class() != PFUNC {
			pstate.Fatalf("gdata NAME not PFUNC %d", nr.Class())
		}
		s.WriteAddr(pstate.Ctxt, nam.Xoffset, wid, pstate.funcsym(nr.Sym).Linksym(pstate.types), nr.Xoffset)

	default:
		pstate.Fatalf("gdata unhandled op %v %v\n", nr, nr.Op)
	}
}

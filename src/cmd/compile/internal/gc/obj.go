package gc

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/dave/golib/src/cmd/compile/internal/types"
	"github.com/dave/golib/src/cmd/internal/bio"
	"github.com/dave/golib/src/cmd/internal/obj"

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

func (psess *PackageSession) dumpobj() {
	if !psess.dolinkobj {
		psess.
			dumpobj1(psess.outfile, modeCompilerObj)
		return
	}
	if psess.linkobj == "" {
		psess.
			dumpobj1(psess.outfile, modeCompilerObj|modeLinkerObj)
		return
	}
	psess.
		dumpobj1(psess.outfile, modeCompilerObj)
	psess.
		dumpobj1(psess.linkobj, modeLinkerObj)
}

func (psess *PackageSession) dumpobj1(outfile string, mode int) {
	bout, err := bio.Create(outfile)
	if err != nil {
		psess.
			flusherrors()
		fmt.Printf("can't create %s: %v\n", outfile, err)
		psess.
			errorexit()
	}
	defer bout.Close()
	bout.WriteString("!<arch>\n")

	if mode&modeCompilerObj != 0 {
		start := startArchiveEntry(bout)
		psess.
			dumpCompilerObj(bout)
		finishArchiveEntry(bout, start, "__.PKGDEF")
	}
	if mode&modeLinkerObj != 0 {
		start := startArchiveEntry(bout)
		psess.
			dumpLinkerObj(bout)
		finishArchiveEntry(bout, start, "_go_.o")
	}
}

func (psess *PackageSession) printObjHeader(bout *bio.Writer) {
	fmt.Fprintf(bout, "go object %s %s %s %s\n", psess.objabi.GOOS, psess.objabi.GOARCH, psess.objabi.Version, psess.objabi.Expstring())
	if psess.buildid != "" {
		fmt.Fprintf(bout, "build id %q\n", psess.buildid)
	}
	if psess.localpkg.Name == "main" {
		fmt.Fprintf(bout, "main\n")
	}
	if psess.safemode {
		fmt.Fprintf(bout, "safe\n")
	} else {
		fmt.Fprintf(bout, "----\n")
	}
	fmt.Fprintf(bout, "\n")
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

func (psess *PackageSession) dumpCompilerObj(bout *bio.Writer) {
	psess.
		printObjHeader(bout)
	psess.
		dumpexport(bout)
}

func (psess *PackageSession) dumpLinkerObj(bout *bio.Writer) {
	psess.
		printObjHeader(bout)

	if len(psess.pragcgobuf) != 0 {

		fmt.Fprintf(bout, "\n$$\n\n$$\n\n")
		fmt.Fprintf(bout, "\n$$  // cgo\n")
		if err := json.NewEncoder(bout).Encode(psess.pragcgobuf); err != nil {
			psess.
				Fatalf("serializing pragcgobuf: %v", err)
		}
		fmt.Fprintf(bout, "\n$$\n\n")
	}

	fmt.Fprintf(bout, "\n!\n")

	externs := len(psess.externdcl)
	psess.
		dumpglobls()
	psess.
		addptabs()
	psess.
		addsignats(psess.externdcl)
	psess.
		dumpsignats()
	psess.
		dumptabs()
	psess.
		dumpimportstrings()
	psess.
		dumpbasictypes()

	for len(psess.compilequeue) > 0 {
		psess.
			compileFunctions()
		psess.
			dumpsignats()
	}

	tmp := psess.externdcl

	if psess.externdcl != nil {
		psess.
			externdcl = psess.externdcl[externs:]
	}
	psess.
		dumpglobls()
	psess.
		externdcl = tmp

	if psess.zerosize > 0 {
		zero := psess.mappkg.Lookup(psess.types, "zero")
		psess.
			ggloblsym(zero.Linksym(psess.types), int32(psess.zerosize), obj.DUPOK|obj.RODATA)
	}
	psess.
		addGCLocals()
	psess.obj.
		WriteObjFile(psess.Ctxt, bout.Writer)
}

func (psess *PackageSession) addptabs() {
	if !psess.Ctxt.Flag_dynlink || psess.localpkg.Name != "main" {
		return
	}
	for _, exportn := range psess.exportlist {
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
			psess.
				ptabs = append(psess.ptabs, ptabEntry{s: s, t: asNode(s.Def).Type})
		} else {
			psess.
				ptabs = append(psess.ptabs, ptabEntry{s: s, t: psess.types.NewPtr(asNode(s.Def).Type)})
		}
	}
}

func (psess *PackageSession) dumpGlobal(n *Node) {
	if n.Type == nil {
		psess.
			Fatalf("external %v nil type\n", n)
	}
	if n.Class() == PFUNC {
		return
	}
	if n.Sym.Pkg != psess.localpkg {
		return
	}
	psess.
		dowidth(n.Type)
	psess.
		ggloblnod(n)
}

func (psess *PackageSession) dumpGlobalConst(n *Node) {

	t := n.Type
	if t == nil {
		return
	}
	if n.Sym.Pkg != psess.localpkg {
		return
	}

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

	case TIDEAL:
		if !psess.Isconst(n, CTINT) {
			return
		}
		x := n.Val().U.(*Mpint)
		if x.Cmp(psess.minintval[TINT]) < 0 || x.Cmp(psess.maxintval[TINT]) > 0 {
			return
		}

		t = psess.types.Types[TINT]
	default:
		return
	}
	psess.
		Ctxt.DwarfIntConst(psess.obj, psess.myimportpath, n.Sym.Name, psess.typesymname(t), n.Int64(psess))
}

func (psess *PackageSession) dumpglobls() {

	for _, n := range psess.externdcl {
		switch n.Op {
		case ONAME:
			psess.
				dumpGlobal(n)
		case OLITERAL:
			psess.
				dumpGlobalConst(n)
		}
	}

	obj.SortSlice(psess.funcsyms, func(i, j int) bool {
		return psess.funcsyms[i].LinksymName() < psess.funcsyms[j].LinksymName()
	})
	for _, s := range psess.funcsyms {
		sf := s.Pkg.Lookup(psess.types, funcsymname(s)).Linksym(psess.types)
		psess.
			dsymptr(sf, 0, s.Linksym(psess.types), 0)
		psess.
			ggloblsym(sf, int32(psess.Widthptr), obj.DUPOK|obj.RODATA)
	}
	psess.
		funcsyms = nil
}

// addGCLocals adds gcargs and gclocals symbols to Ctxt.Data.
// It takes care not to add any duplicates.
// Though the object file format handles duplicates efficiently,
// storing only a single copy of the data,
// failure to remove these duplicates adds a few percent to object file size.
func (psess *PackageSession) addGCLocals() {
	seen := make(map[string]bool)
	for _, s := range psess.Ctxt.Text {
		if s.Func == nil {
			continue
		}
		for _, gcsym := range []*obj.LSym{&s.Func.GCArgs, &s.Func.GCLocals, &s.Func.GCRegs} {
			if seen[gcsym.Name] {
				continue
			}
			psess.
				Ctxt.Data = append(psess.Ctxt.Data, gcsym)
			seen[gcsym.Name] = true
		}
	}
}

func (psess *PackageSession) duintxx(s *obj.LSym, off int, v uint64, wid int) int {
	if off&(wid-1) != 0 {
		psess.
			Fatalf("duintxxLSym: misaligned: v=%d wid=%d off=%d", v, wid, off)
	}
	s.WriteInt(psess.Ctxt, int64(off), wid, int64(v))
	return off + wid
}

func (psess *PackageSession) duint8(s *obj.LSym, off int, v uint8) int {
	return psess.duintxx(s, off, uint64(v), 1)
}

func (psess *PackageSession) duint16(s *obj.LSym, off int, v uint16) int {
	return psess.duintxx(s, off, uint64(v), 2)
}

func (psess *PackageSession) duint32(s *obj.LSym, off int, v uint32) int {
	return psess.duintxx(s, off, uint64(v), 4)
}

func (psess *PackageSession) duintptr(s *obj.LSym, off int, v uint64) int {
	return psess.duintxx(s, off, v, psess.Widthptr)
}

func (psess *PackageSession) dbvec(s *obj.LSym, off int, bv bvec) int {

	for j := 0; int32(j) < bv.n; j += 8 {
		word := bv.b[j/32]
		off = psess.duint8(s, off, uint8(word>>(uint(j)%32)))
	}
	return off
}

func (psess *PackageSession) stringsym(pos src.XPos, s string) (data *obj.LSym) {
	var symname string
	if len(s) > 100 {

		h := sha256.New()
		io.WriteString(h, s)
		symname = fmt.Sprintf(".gostring.%d.%x", len(s), h.Sum(nil))
	} else {

		symname = strconv.Quote(s)
	}

	const prefix = "go.string."
	symdataname := prefix + symname

	symdata := psess.Ctxt.Lookup(symdataname)

	if !symdata.SeenGlobl() {

		off := psess.dsname(symdata, 0, s, pos, "string")
		psess.
			ggloblsym(symdata, int32(off), obj.DUPOK|obj.RODATA|obj.LOCAL)
	}

	return symdata
}

func (psess *PackageSession) slicebytes(nam *Node, s string, len int) {
	psess.
		slicebytes_gen++
	symname := fmt.Sprintf(".gobytes.%d", psess.slicebytes_gen)
	sym := psess.localpkg.Lookup(psess.types, symname)
	sym.Def = asTypesNode(psess.newname(sym))

	lsym := sym.Linksym(psess.types)
	off := psess.dsname(lsym, 0, s, nam.Pos, "slice")
	psess.
		ggloblsym(lsym, int32(off), obj.NOPTR|obj.LOCAL)

	if nam.Op != ONAME {
		psess.
			Fatalf("slicebytes %v", nam)
	}
	nsym := nam.Sym.Linksym(psess.types)
	off = int(nam.Xoffset)
	off = psess.dsymptr(nsym, off, lsym, 0)
	off = psess.duintptr(nsym, off, uint64(len))
	psess.
		duintptr(nsym, off, uint64(len))
}

func (psess *PackageSession) dsname(s *obj.LSym, off int, t string, pos src.XPos, what string) int {

	if int64(len(t)) > 2e9 {
		psess.
			yyerrorl(pos, "%v with length %v is too big", what, len(t))
		return 0
	}

	s.WriteString(psess.Ctxt, int64(off), len(t), t)
	return off + len(t)
}

func (psess *PackageSession) dsymptr(s *obj.LSym, off int, x *obj.LSym, xoff int) int {
	off = int(psess.Rnd(int64(off), int64(psess.Widthptr)))
	s.WriteAddr(psess.Ctxt, int64(off), psess.Widthptr, x, int64(xoff))
	off += psess.Widthptr
	return off
}

func (psess *PackageSession) dsymptrOff(s *obj.LSym, off int, x *obj.LSym) int {
	s.WriteOff(psess.Ctxt, int64(off), x, 0)
	off += 4
	return off
}

func (psess *PackageSession) dsymptrWeakOff(s *obj.LSym, off int, x *obj.LSym) int {
	s.WriteWeakOff(psess.Ctxt, int64(off), x, 0)
	off += 4
	return off
}

func (psess *PackageSession) gdata(nam *Node, nr *Node, wid int) {
	if nam.Op != ONAME {
		psess.
			Fatalf("gdata nam op %v", nam.Op)
	}
	if nam.Sym == nil {
		psess.
			Fatalf("gdata nil nam sym")
	}
	s := nam.Sym.Linksym(psess.types)

	switch nr.Op {
	case OLITERAL:
		switch u := nr.Val().U.(type) {
		case bool:
			i := int64(obj.Bool2int(u))
			s.WriteInt(psess.Ctxt, nam.Xoffset, wid, i)

		case *Mpint:
			s.WriteInt(psess.Ctxt, nam.Xoffset, wid, u.Int64(psess))

		case *Mpflt:
			f := u.Float64(psess)
			switch nam.Type.Etype {
			case TFLOAT32:
				s.WriteFloat32(psess.Ctxt, nam.Xoffset, float32(f))
			case TFLOAT64:
				s.WriteFloat64(psess.Ctxt, nam.Xoffset, f)
			}

		case *Mpcplx:
			r := u.Real.Float64(psess)
			i := u.Imag.Float64(psess)
			switch nam.Type.Etype {
			case TCOMPLEX64:
				s.WriteFloat32(psess.Ctxt, nam.Xoffset, float32(r))
				s.WriteFloat32(psess.Ctxt, nam.Xoffset+4, float32(i))
			case TCOMPLEX128:
				s.WriteFloat64(psess.Ctxt, nam.Xoffset, r)
				s.WriteFloat64(psess.Ctxt, nam.Xoffset+8, i)
			}

		case string:
			symdata := psess.stringsym(nam.Pos, u)
			s.WriteAddr(psess.Ctxt, nam.Xoffset, psess.Widthptr, symdata, 0)
			s.WriteInt(psess.Ctxt, nam.Xoffset+int64(psess.Widthptr), psess.Widthptr, int64(len(u)))

		default:
			psess.
				Fatalf("gdata unhandled OLITERAL %v", nr)
		}

	case OADDR:
		if nr.Left.Op != ONAME {
			psess.
				Fatalf("gdata ADDR left op %v", nr.Left.Op)
		}
		to := nr.Left
		s.WriteAddr(psess.Ctxt, nam.Xoffset, wid, to.Sym.Linksym(psess.types), to.Xoffset)

	case ONAME:
		if nr.Class() != PFUNC {
			psess.
				Fatalf("gdata NAME not PFUNC %d", nr.Class())
		}
		s.WriteAddr(psess.Ctxt, nam.Xoffset, wid, psess.funcsym(nr.Sym).Linksym(psess.types), nr.Xoffset)

	default:
		psess.
			Fatalf("gdata unhandled op %v %v\n", nr, nr.Op)
	}
}

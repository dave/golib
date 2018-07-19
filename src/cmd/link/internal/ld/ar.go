package ld

import (
	"encoding/binary"
	"fmt"
	"github.com/dave/golib/src/cmd/internal/bio"

	"github.com/dave/golib/src/cmd/link/internal/sym"
	"io"
	"os"
)

const (
	SARMAG  = 8
	SAR_HDR = 16 + 44
)

const (
	ARMAG = "!<arch>\n"
)

type ArHdr struct {
	name string
	date string
	uid  string
	gid  string
	mode string
	size string
	fmag string
}

// hostArchive reads an archive file holding host objects and links in
// required objects. The general format is the same as a Go archive
// file, but it has an armap listing symbols and the objects that
// define them. This is used for the compiler support library
// libgcc.a.
func (psess *PackageSession) hostArchive(ctxt *Link, name string) {
	f, err := bio.Open(name)
	if err != nil {
		if os.IsNotExist(err) {

			if ctxt.Debugvlog != 0 {
				ctxt.Logf("skipping libgcc file: %v\n", err)
			}
			return
		}
		psess.
			Exitf("cannot open file %s: %v", name, err)
	}
	defer f.Close()

	var magbuf [len(ARMAG)]byte
	if _, err := io.ReadFull(f, magbuf[:]); err != nil {
		psess.
			Exitf("file %s too short", name)
	}

	if string(magbuf[:]) != ARMAG {
		psess.
			Exitf("%s is not an archive file", name)
	}

	var arhdr ArHdr
	l := nextar(f, f.Offset(), &arhdr)
	if l <= 0 {
		psess.
			Exitf("%s missing armap", name)
	}

	var armap archiveMap
	if arhdr.name == "/" || arhdr.name == "/SYM64/" {
		armap = psess.readArmap(name, f, arhdr)
	} else {
		psess.
			Exitf("%s missing armap", name)
	}

	loaded := make(map[uint64]bool)
	any := true
	for any {
		var load []uint64
		for _, s := range ctxt.Syms.Allsym {
			for _, r := range s.R {
				if r.Sym != nil && r.Sym.Type == sym.SXREF {
					if off := armap[r.Sym.Name]; off != 0 && !loaded[off] {
						load = append(load, off)
						loaded[off] = true
					}
				}
			}
		}

		for _, off := range load {
			l := nextar(f, int64(off), &arhdr)
			if l <= 0 {
				psess.
					Exitf("%s missing archive entry at offset %d", name, off)
			}
			pname := fmt.Sprintf("%s(%s)", name, arhdr.name)
			l = atolwhex(arhdr.size)

			libgcc := sym.Library{Pkg: "libgcc"}
			h := psess.ldobj(ctxt, f, &libgcc, l, pname, name)
			f.Seek(h.off, 0)
			h.ld(ctxt, f, h.pkg, h.length, h.pn)
		}

		any = len(load) > 0
	}
}

// archiveMap is an archive symbol map: a mapping from symbol name to
// offset within the archive file.
type archiveMap map[string]uint64

// readArmap reads the archive symbol map.
func (psess *PackageSession) readArmap(filename string, f *bio.Reader, arhdr ArHdr) archiveMap {
	is64 := arhdr.name == "/SYM64/"
	wordSize := 4
	if is64 {
		wordSize = 8
	}

	contents := make([]byte, atolwhex(arhdr.size))
	if _, err := io.ReadFull(f, contents); err != nil {
		psess.
			Exitf("short read from %s", filename)
	}

	var c uint64
	if is64 {
		c = binary.BigEndian.Uint64(contents)
	} else {
		c = uint64(binary.BigEndian.Uint32(contents))
	}
	contents = contents[wordSize:]

	ret := make(archiveMap)

	names := contents[c*uint64(wordSize):]
	for i := uint64(0); i < c; i++ {
		n := 0
		for names[n] != 0 {
			n++
		}
		name := string(names[:n])
		names = names[n+1:]

		if psess.objabi.GOOS == "darwin" || (psess.objabi.GOOS == "windows" && psess.objabi.GOARCH == "386") {
			if name[0] == '_' && len(name) > 1 {
				name = name[1:]
			}
		}

		var off uint64
		if is64 {
			off = binary.BigEndian.Uint64(contents)
		} else {
			off = uint64(binary.BigEndian.Uint32(contents))
		}
		contents = contents[wordSize:]

		ret[name] = off
	}

	return ret
}

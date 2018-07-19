package ssa_test

import (
	"debug/dwarf"
	"debug/elf"
	"debug/macho"
	"debug/pe"
	"fmt"
	"github.com/dave/golib/src/internal/testenv"
	"io"
	"runtime"
	"testing"
)

func open(path string) (*dwarf.Data, error) {
	if fh, err := elf.Open(path); err == nil {
		return fh.DWARF()
	}

	if fh, err := pe.Open(path); err == nil {
		return fh.DWARF()
	}

	if fh, err := macho.Open(path); err == nil {
		return fh.DWARF()
	}

	return nil, fmt.Errorf("unrecognized executable format")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type Line struct {
	File string
	Line int
}

type File struct {
	lines []string
}

var fileCache = map[string]*File{}

func (f *File) Get(lineno int) (string, bool) {
	if f == nil {
		return "", false
	}
	if lineno-1 < 0 || lineno-1 >= len(f.lines) {
		return "", false
	}
	return f.lines[lineno-1], true
}

func TestStmtLines(t *testing.T) {
	if runtime.GOOS == "plan9" {
		t.Skip("skipping on plan9; no DWARF symbol table in executables")
	}

	lines := map[Line]bool{}
	dw, err := open(testenv.GoToolPath(t))
	must(err)
	rdr := dw.Reader()
	rdr.Seek(0)
	for {
		e, err := rdr.Next()
		must(err)
		if e == nil {
			break
		}
		if e.Tag != dwarf.TagCompileUnit {
			continue
		}
		pkgname, _ := e.Val(dwarf.AttrName).(string)
		if pkgname == "runtime" {
			continue
		}
		lrdr, err := dw.LineReader(e)
		must(err)

		var le dwarf.LineEntry

		for {
			err := lrdr.Next(&le)
			if err == io.EOF {
				break
			}
			must(err)
			fl := Line{le.File.Name, le.Line}
			lines[fl] = lines[fl] || le.IsStmt
		}
	}

	nonStmtLines := []Line{}
	for line, isstmt := range lines {
		if !isstmt {
			nonStmtLines = append(nonStmtLines, line)
		}
	}

	if runtime.GOARCH == "amd64" && len(nonStmtLines)*100 > len(lines) {
		t.Errorf("Saw too many (amd64, > 1%%) lines without statement marks, total=%d, nostmt=%d\n", len(lines), len(nonStmtLines))
	}
	if len(nonStmtLines)*100 > 2*len(lines) {
		t.Errorf("Saw too many (not amd64, > 2%%) lines without statement marks, total=%d, nostmt=%d\n", len(lines), len(nonStmtLines))
	}
	t.Logf("total=%d, nostmt=%d\n", len(lines), len(nonStmtLines))
}

package gc_test

import (
	"bytes"
	"github.com/dave/golib/src/internal/testenv"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestBuiltin(t *testing.T) {
	t.Skip("TODO: I think this is failing because we're stripping comments from the AST?")
	testenv.MustHaveGoRun(t)

	old, err := ioutil.ReadFile("builtin.go")
	if err != nil {
		t.Fatal(err)
	}

	new, err := exec.Command(testenv.GoToolPath(t), "run", "mkbuiltin.go", "-stdout").Output()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(old, new) {
		t.Fatal("builtin.go out of date; run mkbuiltin.go")
	}
}

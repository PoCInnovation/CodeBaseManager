package REPL

import (
	"fmt"
	pt "github.com/PoCFrance/CodeBaseManager/REPL/tests"
	"io/ioutil"
	"os"
	"testing"
)

type testHandleExternal struct {
	in       []string
	expected string
}

func (t *testHandleExternal) isValid(got string) error {
	if t.expected != got {
		return pt.TestError(t.expected, got)
	}
	return nil
}

func runTest(name string, test *testHandleExternal, t *testing.T) {
	save := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Skip(err)
	}

	os.Stdout = w
	handleExternal(test.in)
	w.Close()

	out, err := ioutil.ReadAll(r)
	os.Stdout = save
	fmt.Println("got", out)

	if err = test.isValid(string(out)); err != nil {
		t.Errorf("%s%s", name, err.Error())
	} else {
		t.Logf("%s: Success!", name)
	}
}

func TestHandleExternal(t *testing.T) {
	tests := map[string]testHandleExternal{
		"From PATH": {
			[]string{"echo", "hello world"},
			"hello world\n",
		},
		"From PWD": {
			[]string{"./tests/random.sh"},
			"hello world\n",
		},
		"Not found": {
			[]string{"lol"},
			"lol: command not found\n",
		},
		"Directory": {
			[]string{"./cmd"},
			"lol",
		},
	}

	for name, test := range tests {
		runTest(name, &test, t)
	}
}

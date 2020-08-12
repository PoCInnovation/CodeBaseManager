package REPL

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/REPL/builtins"
	"github.com/PoCFrance/CodeBaseManager/REPL/tests"
	"github.com/PoCFrance/CodeBaseManager/modules/funcTests"
	"reflect"
	"testing"
)

type in struct {
	input    string
	builtins Builtins
}

type exp struct {
	parsed []string
	fn     builtin
}

type testInput struct {
	name string
	in   in
	exp  exp
}

func (tIn *testInput) isValid(parsed []string, fn builtin) error {
	if len(parsed) != len(tIn.exp.parsed) {
		return pt.TestError(tIn.exp.parsed, parsed)
	}
	for idx := range tIn.exp.parsed {
		if tIn.exp.parsed[idx] != parsed[idx] {
			return pt.TestError(tIn.exp.parsed, parsed)
		}
	}
	refExpFn := reflect.ValueOf(tIn.exp.fn)
	refGotFn := reflect.ValueOf(fn)
	fmt.Println(refGotFn.String())
	if refExpFn.Pointer() != refGotFn.Pointer() {
		return pt.TestError(refExpFn.String(), refGotFn.String())
	}
	return nil
}

func TestParseInput(t *testing.T) {
	tests := []testInput{
		{
			"Basic extern command",
			in{
				"ls lol",
				Builtins{},
			},
			exp{
				[]string{"ls", "lol"},
				handleExternal,
			},
		},
		{
			"Basic common builtin",
			in{
				"cd cmd",
				Builtins{},
			},
			exp{
				[]string{"cd", "cmd"},
				builtins.CD,
			},
		},
		{
			"Basic CBM builtin",
			in{
				"run",
				Builtins{
					"run": funcTests.Run,
				},
			},
			exp{
				[]string{"run"},
				funcTests.Run,
			},
		},
		{
			"Empty input",
			in{
				"",
				Builtins{},
			},
			exp{
				[]string{"continue"},
				nil,
			},
		},
		{
			"Exit",
			in{
				"exit",
				Builtins{},
			},
			exp{
				[]string{"exit"},
				nil,
			},
		},
		{
			"Fucked up input",
			in{
				"        \t\t\t ezlaee ezae e      eza eza az",
				Builtins{},
			},
			exp{
				[]string{"ezlaee", "ezae", "e", "eza", "eza", "az"},
				handleExternal,
			},
		},
	}
	for _, test := range tests {
		parsed, fn := parseInput(test.in.input, test.in.builtins)
		if err := test.isValid(parsed, fn); err != nil {
			t.Errorf("%s%s", test.name, err.Error())
		} else {
			t.Logf("%s: Success!", test.name)
		}
	}
}

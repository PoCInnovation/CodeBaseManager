package codebase_test

import (
	"github.com/PoCFrance/CodeBaseManager/cmd/codebase"
	"log"
	"os"
	"testing"
)

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}

func TestFindCFunction(t *testing.T) {
	defer quiet()()
	tables := []struct {
		args []string
		exp  []string
	}{
		{args: []string{"signal_handler"},
			exp: []string{"Function signal_handler at line : 14"}},
		//{[]string{"-0.5"}, 0., &argv.ExitError},
	}

	for _, table := range tables {
		res := codebase.Cat(table.args)
		if res == nil {
			t.Errorf("For argument(s) [%v]), res is nil\n", table.args)
		}
		//for _, arg := range table.args {
		//	t.Log(res[arg]["lol"])
			//if contentFound, ok := res[arg]; ok {
			//
			//}
		//}
			//for key, content := range res[arg]["lol"] {
			//if res[args] != table.exp {
			//	t.Errorf(
			//		"For argument(s) [%v]), res is [%v] (Expected [%v]\n", table.args, res, table.exp)
			//}
		//}
	}
}

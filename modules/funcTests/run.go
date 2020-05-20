package funcTests

import (
	"fmt"
)

func Run(_ []string) {
	// TODO: More flexibility on path
	cfg, err := NewConfigFT(".cbm/template/cantFail.toml")
	if err != nil {
		fmt.Println(err)
		// TODO: Would you like to continue ? yes | exit
		return
	}
	// TODO: if no bin ask build module for binary
	// TODO: use exec.lookPath to make sure our bin's ok
	for _, test := range cfg.Tests {
		test.Init(&cfg.Common)
		test.Run()
		test.GetResults()
	}
	// TODO: show results
}

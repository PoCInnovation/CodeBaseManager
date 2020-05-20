package funcTests

import (
	"fmt"
)

func Run(av []string) {
	// TODO: More flexibility on path
	for _, fp := range av {
		cfg, err := NewConfigFT(fp)
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
}

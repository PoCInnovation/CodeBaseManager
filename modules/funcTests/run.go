package funcTests

import (
	"fmt"
)

func Run(_ []string) {
	cfg, err := NewConfigFT(".cbm/template/ft.toml")
	if err != nil {
		fmt.Println(err)
		// TODO: Would you like to continue ? yes | exit
		return
	}
	// TODO: if no bin ask build module for binary
	for _, test := range cfg.Tests {
		fmt.Printf("Setting up %s...\n", test.Name)
		test.Opt.SetCommon(&cfg.Common.Opt)
		test.myExec.Set(cfg.Common.Bin, test.Args...)
		if cfg.Common.RefBin != "" {
			test.refExec.Set(cfg.Common.RefBin, test.RefArgs...)
		}
		// TODO: Setup tests
		//       - Commands
		//          - their envs
		//          - pipes & stdin
		test.myExec.Run()
		fmt.Println(test.myExec.outBuf.String(), test.myExec.errBuf.String())
		// TODO: Run tests
		//       - run pre
		//       - run /!\ Options /!\
		//       - run post
		// TODO: compare with expected.
		//
	}
	// TODO: show results
}

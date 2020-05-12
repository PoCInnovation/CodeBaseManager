package funcTests

import "fmt"

const (
	noBin = ""
	noCmd = ""
)

type FT struct {
	Name    string
	Desc    string
	Args    []string
	RefArgs []string

	Exp     ftExpected     `toml:"expected"`
	Ext     ftInteractions `toml:"interactions"`
	Opt     ftOptions      `toml:"options"`
	my, ref ftExecution
	res     ftResult
}

func (test *FT) Init(basicOpt *ftCommon) {
	fmt.Printf("Setting up %s...\n", test.Name)

	test.Opt.SetCommon(&basicOpt.Opt)

	test.my.Set(&test.Ext, basicOpt.Bin, test.Args...)
	if basicOpt.RefBin != noBin {
		test.ref.Set(&test.Ext, basicOpt.RefBin, test.RefArgs...)
	}

}

func (test *FT) Run() {
	if test.Ext.Pre != noCmd {
		// TODO: Improvements? Error handling or else?
		if err := QuickRun(test.Ext.Pre); err != nil {
			fmt.Println(err)
			return
		}
	}

	test.my.Run() // TODO: care about options
	if test.ref.cmd != nil {
		test.ref.Run() // TODO: care about options
	}

	test.my.AfterPipe(test.Ext.StdoutPipe, test.Ext.StderrPipe)
	if test.ref.cmd != nil {
		test.ref.AfterPipe(test.Ext.StdoutPipe, test.Ext.StderrPipe)
	}
	if test.Ext.Post != noCmd {
		// TODO: Improvements? Error handling or else?
		if err := QuickRun(test.Ext.Post); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(test.my.outBuf.String(), test.my.errBuf.String())
}

func (test *FT) GetResults() {
	if test.ref.cmd != nil {
		test.res.CompareToRef(&test.ref, &test.my)
	} else {
		test.res.CompareToExp(&test.Exp, &test.my)
	}
}


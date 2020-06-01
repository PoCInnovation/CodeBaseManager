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
}

func (test *FT) Init(basicOpt *ftCommon) {
	test.Opt.SetCommon(&basicOpt.Opt)

	test.my.Set(&test.Ext, basicOpt.Bin, test.Args...)
	if basicOpt.RefBin != noBin {
		test.ref.Set(&test.Ext, basicOpt.RefBin, test.RefArgs...)
	}

}

func (test *FT) Run(options ftOptions) {
	if test.Ext.Pre != noCmd {
		// TODO: Improvements? Error handling or else?
		if err := QuickRun(test.Ext.Pre); err != nil {
			fmt.Println("Pre:", err)
			return
		}
	}

	test.my.Run(test.Opt) // TODO: care about options
	if test.ref.cmd != nil {
		test.ref.Run(test.Opt) // TODO: care about options
	}

	test.my.AfterPipe(test.Ext.StdoutPipe, test.Ext.StderrPipe)
	if test.ref.cmd != nil {
		test.ref.AfterPipe(test.Ext.StdoutPipe, test.Ext.StderrPipe)
	}
	if test.Ext.Post != noCmd {
		// TODO: Improvements? Error handling or else?
		if err := QuickRun(test.Ext.Post); err != nil {
			fmt.Println("Post:", err)
		}
	}
}

func (test *FT) GetResults() *ftResult {
	res := &ftResult{}
	if test.ref.cmd != nil {
		res.CompareToRef(&test.ref, &test.my)
	} else {
		res.CompareToExp(&test.Exp, &test.my)
	}
	return res
}


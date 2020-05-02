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

	Exp    ftExpected     `toml:"expected"`
	Ext    ftInteractions `toml:"interactions"`
	Opt    ftOptions      `toml:"options"`
	myExec, refExec ftExecution

}

func (test *FT) Init(basicOpt *ftCommon) {
	fmt.Printf("Setting up %s...\n", test.Name)

	test.Opt.SetCommon(&basicOpt.Opt)

	test.myExec.Set(&test.Ext, basicOpt.Bin, test.Args...)
	if basicOpt.RefBin != noBin {
		test.refExec.Set(&test.Ext, basicOpt.RefBin, test.RefArgs...)
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

	test.myExec.Run() // TODO: care about options

	if test.Ext.Post != noCmd {
		// TODO: Improvements? Error handling or else?
		if err := QuickRun(test.Ext.Post); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(test.myExec.outBuf.String(), test.myExec.errBuf.String())
}

type ftExpected struct {
	Status     int
	Stdout     string
	Stderr     string
	StdoutFile string
	StderrFile string
}

type ftInteractions struct {
	StdoutPipe string
	StderrPipe string
	Stdin      string
	StdinFile  string

	Pre    string
	Post   string
	Env    []string
	AddEnv []string
}
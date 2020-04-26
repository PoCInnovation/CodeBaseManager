package funcTests

type ConfigFT struct {
	Common struct {
		Bin string `toml:"bin"`
		RefBin string `toml:"refBin"`
		Opt ftOptions `toml:"options"`
	}
	Tests []FT `toml:"Test"`
}

type FT struct {
	Name string
	Desc string
	Args []string
	RefArgs []string

	Exp ftExpected `toml:"expected"`
	Ext ftInteractions `toml:"interactions"`
	Opt ftOptions `toml:"options"`
}

type ftOptions struct {
	Repeat bool
	Time bool
	ShouldFail bool
	Timeout int64
	BuildWith string
}

type ftExpected struct {
	Status int
	Stdout string
	Stderr string
	StdoutFile string
	StderrFile string
}

type ftInteractions struct {
	StdoutPipe string
	StderrPipe string
	Stdin string
	StdinFile string

	Pre string
	Post string
	Env string
	AddEnv string
}
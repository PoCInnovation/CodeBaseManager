package funcTests

type FT struct {
	Name    string
	Desc    string
	Args    []string
	RefArgs []string

	Exp ftExpected     `toml:"expected"`
	Ext ftInteractions `toml:"interactions"`
	Opt ftOptions      `toml:"options"`
	exec [2]ftExecution
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
	Env    string
	AddEnv string
}
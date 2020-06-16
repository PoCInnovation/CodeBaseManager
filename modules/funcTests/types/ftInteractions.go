package ft_types

type ftInteractions struct {
	StdoutPipe string `toml:"stdoutPipe"`
	StderrPipe string `toml:"stderrPipe"`
	StdinPipe  string `toml:"stdinPipe"`
	StdinFile  string `toml:"stdinFile"`
	Stdin      string `toml:"stdin"`

	Pre    string `toml:"pre"`
	Post   string `toml:"post"`
	Env    []string `toml:"env"`
	AddEnv []string `toml:"addEnv"`
}

func (test *ftInteractions) ApplyDefault(reference ftInteractions) {
	if len(test.AddEnv) == 0 {
		test.AddEnv = reference.AddEnv
	}
	if len(test.Env) == 0 {
		test.Env = reference.Env
	}
	if len(test.Post) == 0 {
		test.Post = reference.Post
	}
	if len(test.Pre) == 0 {
		test.Pre = reference.Pre
	}
	if len(test.StderrPipe) == 0 {
		test.StderrPipe = reference.StderrPipe
	}
	if len(test.StdinPipe) == 0 {
		test.StdinPipe = reference.StdinPipe
	}
	if len(test.StdoutPipe) == 0 {
		test.StdoutPipe = reference.StdoutPipe
	}
	if len(test.Stdin) == 0 {
		test.Stdin = reference.Stdin
	}
	if len(test.StdinFile) == 0 {
		test.StdinFile = reference.StdinFile
	}
}
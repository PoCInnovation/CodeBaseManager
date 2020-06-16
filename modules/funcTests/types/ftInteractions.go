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

func ApplyIfEmpty(current, ref string) {
	if len(current) == 0 {
		
	}
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
	if len(test.Stdin) == 0 {
		test.Stdin = reference.
	}
	if len() == 0 {
		test. = reference.
	}
	if len() == 0 {
		test. = reference.
	}
	if len() == 0 {
		test. = reference.
	}
	if len() == 0 {
		test. = reference.
	}
	if len() == 0 {
		test. = reference.
	}
}
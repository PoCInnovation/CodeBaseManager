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

func (test *ftInteractions) ApplyVars(vars map[string]string) {
	test.AddEnv = ApplyVarsTab(test.AddEnv, vars)
	test.Env = ApplyVarsTab(test.Env, vars)
	test.Post = ApplyVarsString(test.Post, vars)
	test.Pre = ApplyVarsString(test.Pre, vars)
	test.StderrPipe = ApplyVarsString(test.StderrPipe, vars)
	test.StdinPipe = ApplyVarsString(test.StdinPipe, vars)
	test.StdoutPipe = ApplyVarsString(test.StdoutPipe, vars)
	test.Stdin = ApplyVarsString(test.Stdin, vars)
	test.StdinFile = ApplyVarsString(test.StdinFile, vars)
}

func (test *ftInteractions) ApplyDefault(reference ftInteractions, vars map[string]string) {
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
	test.ApplyVars(vars)
}
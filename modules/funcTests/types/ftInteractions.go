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
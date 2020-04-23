package REPL

type Shell struct {
	*prompt
}

func NewShell(module string) *Shell {
	sh := &Shell{NewPrompt(module)}

	return sh
}

func (sh *Shell) Run(acceptedBuiltins []string) {
	defer sh.Close()

	for {
		sh.Display()
		input := sh.GetInput()
		parsed, todo := parseInput(input, acceptedBuiltins)

		switch todo {
		case Builtin:
			handleBuiltin(parsed)
		case ExternalBin:
			handleExternal(parsed)
		case Exit:
			return
		}
	}
}
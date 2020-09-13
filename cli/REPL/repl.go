package REPL

type Shell struct {
	*prompt
}

func NewShell(module string) *Shell {
	sh := &Shell{NewPrompt(module)}

	return sh
}

func (sh *Shell) Run(cbmBuiltins Builtins) {
	defer sh.Close()

	for {
		sh.Display()
		input := sh.GetInput()
		parsed, execFn := parseInput(input, cbmBuiltins)

		switch {
		case parsed[0] == "continue":
			continue
		case parsed[0] == "exit":
			return
		default:
			execFn(parsed)
		}
	}
}

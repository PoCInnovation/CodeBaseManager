package REPL

func CodebaseShell() {
	codebaseCmd := []string{"cat", "find"}
	p := newPrompt("CodeBase")
	defer p.Close()

	for {
		p.Display()
		in := p.GetInput()
		parsed, todo := parseInput(in, codebaseCmd)

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

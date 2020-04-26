package REPL

import (
	"strings"
)

const (
	CommonBuiltin = iota
	CBMBuiltin
	ExternalBin
	Exit
	Continue
)

func isExit(in string) bool {
	return len(in) == len("exit") && in == "exit"
}

func parseInput(in string, cbmBuiltins Builtins) ([]string, builtin) {
	in = strings.TrimSuffix(in, "\n")
	if in == "" {
		return []string{"continue"}, nil
	} else if isExit(in) {
		return []string{in}, nil
	}

	parsed := strings.Fields(in)
	if fn := isBuiltin(parsed[0], commonBuiltins); fn != nil {
		return parsed, fn
	}
	if fn := isBuiltin(parsed[0], cbmBuiltins); fn != nil {
		return parsed, fn
	}
	return parsed, handleExternal
}

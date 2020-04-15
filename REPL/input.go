package REPL

import "strings"

const (
	Builtin = iota
	ExternalBin
	Exit
	Continue
)

func isExit(in string) bool {
	return len(in) == len("exit") && in == "exit"
}

func parseInput(in string, builtins []string) ([]string, int) {
	in = strings.TrimSuffix(in, "\n")
	if in == "" {
		return nil, Continue
	} else if isExit(in) {
		return nil, Exit
	}
	//TODO: Check what's av0 according to the consts above
	return strings.Fields(in), Continue
}

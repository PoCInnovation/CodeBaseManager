package REPL

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func execute(bin string, av []string) {
	cmd := exec.Command(bin, av...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func locateBinary(given string) string {
	const X_OK = 1
	if err := syscall.Access(given, X_OK); err != nil {
		fmt.Println(err)
		return ""
	}
	return given
}

func handleExternal(av []string) {
	bin := locateBinary(av[0])
	if bin == "" {
		return
	}
	execute(bin, av)
}

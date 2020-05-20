package REPL

import (
	"fmt"
	"log"
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

func LocateBinary(given string) string {
	const X_OK = 1
	if err := syscall.Access("./" + given, X_OK); err == nil {
		return given
	}
	found, err := exec.LookPath(given)
	if err != nil {
		log.Printf("%s: command not found\n", given)
		return ""
	}
	return found
}

func handleExternal(av []string) {
	bin := LocateBinary(av[0])
	if bin == "" {
		return
	}
	execute(bin, av[1:])
}

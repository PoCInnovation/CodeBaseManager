package repository

import (
	"fmt"
	"github.com/PoCInnovation/CodeBaseManager/cli/modules/server"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	url = 0
	dir = 1
)

func CreateRepository(args []string) {
	var repo string

	if len(args) == 1 {
		fmt.Printf("Creating repo based on: %s\n", args[url])
		repo = strings.TrimSuffix(path.Base(args[url]), ".git")
	} else {
		fmt.Printf("Creating repo based on: %s in %s\n", args[url], args[dir])
		repo = args[dir]
	}

	if err := cloneRepo(args); err != nil {
		return
	}

	server.Add(repo)
}

func cloneRepo(cmd []string) error {
	args := append([]string{"clone"}, cmd...)
	execCmd := exec.Command("git", args...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	if err := execCmd.Run(); err != nil {
		return err
	}
	return nil
}

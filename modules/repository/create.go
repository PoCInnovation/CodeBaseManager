package repository

import (
	"fmt"
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
	args = append([]string{"clone"}, args...)
	if err := cloneRepo(args); err != nil {
		return
	}
	watchRepo(repo)
}

func cloneRepo(cmd []string) error {
	execCmd := exec.Command("git", cmd...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	if err := execCmd.Run(); err != nil {
		return err
	}
	return nil
}

func watchRepo(repo string) {
	//actualPath, err := os.Getwd()
	//if err := os.Chdir("./" + repo); err != nil {
	//	log.Println(err)
	//	return
	//}
	//defer func() {
	//	if err := os.Chdir("./" + repo); err != nil {
	//		log.Println(err)
	//		return
	//	}
	//}()
	WatchRepository(repo)
}

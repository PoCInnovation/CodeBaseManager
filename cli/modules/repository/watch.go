package repository

import (
	"github.com/PoCInnovation/CodeBaseManager/cli/modules/server"
	"log"
	"os"
	"strings"
)

const (
	GitDir = ".git"
)

func WatchRepository(args []string) {
	var path string
	var err error
	if len(args) == 0 {
		path, err = os.Getwd()
		if err != nil {
			log.Println(err)
		}
		WatchPath(path)
	} else {
		for _, arg := range args {
			WatchPath(arg)
		}
	}
}

func WatchPath(path string) {
	if !CheckGit(path) {
		log.Println(path, ": is not a git repository")
		return
	}
	server.Add(path)
}

func CheckGit(path string) bool {
	dir, err := os.Open(path)
	if err != nil {
		log.Printf("Error when opening %s, %v\n", path, err)
	}
	defer func() {
		if err := dir.Close(); err != nil {
			log.Printf("Cannot close :%s, %v\n", path, err)
		}
	}()

	targetList, err := dir.Readdir(0)
	if err != nil {
		log.Printf("Error when Readdir %s, %v\n", path, err)
		return false
	}

	for _, target := range targetList {
		log.Println(target.Name())
		if target.IsDir() && strings.HasSuffix(target.Name(), GitDir) {
			return true
		}
	}

	return false
}

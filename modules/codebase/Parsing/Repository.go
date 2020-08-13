package codebase

import (
	"errors"
	Components "github.com/PoCFrance/CodeBaseManager/modules/codebase/Components"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ParseRepositoryV2(path string) (*Components.Repository, error) {
	path = strings.TrimSuffix(path, "/")
	if !isDir(path) {
		return nil, errors.New(path + " is not a valid target")
	}
	module := &Components.Module{
		Path: path,
		Name: filepath.Base(path),
	}
	parser := &Components.Repository{
		Modules: []*Components.Module{module},
	}
	ProcessModules(parser, module, path)

	return parser, nil
}

func listTargets(path string) ([]os.FileInfo, func()) {
	dir, err := os.Open(path)
	if err != nil {
		log.Printf("Error when opening %s, %v\n", path, err)
	}
	fn := func() {
		if err := dir.Close(); err != nil {
			log.Printf("Cannot close :%s, %v\n", path, err)
		}
	}

	targetList, err := dir.Readdir(0)
	if err != nil {
		log.Printf("Error when Readdir %s, %v\n", path, err)
		return nil, nil
	}
	return targetList, fn
}

func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Printf("Error when stating %s, %v\n", path, err)
		return false
	}
	if !fileInfo.IsDir() {
		return false
	}
	return true
}

func isNotReadable(name string) bool {
	info, err := os.Stat(name)

	if err != nil {
		return true
	}

	perm := info.Mode()
	return info.Size() > 5000 && perm&0111 == 0111
	//return info.Size() < 10000 && perm&0111 != 0111
}

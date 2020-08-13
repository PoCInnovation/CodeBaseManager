package codebase

import (
	"log"
	"os"
)

type FileParser struct {
	Path, Name string
}

func (f FileParser) isTarget(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Printf("Error when stating %s, %v\n", path, err)
		return false
	}
	if fileInfo.IsDir() {
		return false
	}
	return true
}

func (f FileParser) String() string {
	str := "File name: " + f.Name + ", " + "File path: " + f.Path
	return str
}

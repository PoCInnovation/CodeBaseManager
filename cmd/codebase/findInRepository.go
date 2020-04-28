package codebase

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type fctParser func(string, parsingRepo)

type contentFound map[string]map[string]string

type fileFct func(map[string]string, string) (map[string]string, error)

type parsingRepo struct {
	manageFile fileFct
	parser     fctParser
	content    contentFound
	args       []string
}

func RepoParser(module string, control parsingRepo) {
	// TODO: Manage hidden folder ?
	splitName := strings.Split(module, "/")
	splitLen := len(splitName)
	if (module != ".") && (splitLen == 0 || strings.HasPrefix(splitName[splitLen-1], ".")) {
		return
	}

	dir, err := os.Open(module)
	if err != nil {
		log.Printf("Error when opening module %s, %v\n", module, err)
		return
	}
	defer func() {
		err := dir.Close()
		if err != nil {
			log.Printf("Cannot close module :%s, %v\n", module, err)
		}
	}()

	files, err := dir.Readdir(0)
	if err != nil {
		log.Printf("Error when Readdir module %s, %v\n", module, err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			control.parser(module+"/"+file.Name(), control)
			//argParser(module+"/"+file.Name(), control)
		} else {
			RepoParser(module+"/"+file.Name(), control)
		}
	}
}

func argParser(name string, control parsingRepo) {
	for _, arg := range control.args {
		splitName := strings.Split(name, "/")
		splitLen := len(splitName)
		if splitLen == 0 {
			log.Printf("Cannot Split %s\n", name)
		}

		if arg == splitName[splitLen-1] {
			// TODO: refacto parsing to use fctPtr -> common ground for cat and find
			control.content[arg], _ = control.manageFile(control.content[arg], name)
		} else {
			//read content to find function
		}
	}
}

func PrintResult(args []string, parser parsingRepo) {
	for _, arg := range args {
		fmt.Println(strings.Repeat("==", 50))
		fmt.Printf("ARG: %s\n", arg)
		if contentFound, ok := parser.content[arg]; ok {
			for key, content := range contentFound {
				fmt.Println(strings.Repeat("#", 50))
				fmt.Printf("FOUND ===> %s\n", key)
				fmt.Println(content)
			}
		}
	}
}

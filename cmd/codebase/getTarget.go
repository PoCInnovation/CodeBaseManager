package codebase

import (
	"log"
	"os"
	"strings"
)

type fctParser func(string, parsingRepo)

type contentFound map[string]map[string]string

//type fileFct func(map[string]string, string) (map[string]string, error)

type parsingRepo struct {
	parser fctParser
	//fileFct fileFct
	content contentFound
	args    []string
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
			//fmt.Println(module + "/" + file.Name())
			control.parser(module+"/"+file.Name(), control)
			// manage file
		} else {
			RepoParser(module+"/"+file.Name(), control)
		}
	}
}

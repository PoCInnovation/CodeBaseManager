package codebase

import (
	"github.com/PoCFrance/CodeBaseManager/modules/codebase"
	"log"
	"strings"
)

func CatParser(name string, control parsingRepo) {
	for _, arg := range control.args {
		splitName := strings.Split(name, "/")
		splitLen := len(splitName)
		if splitLen == 0 {
			log.Printf("Cannot Split %s\n", name)
		}

		if arg == splitName[splitLen-1] {
			// TODO: refacto parsing to use fctPtr -> common ground for cat and find
			//var err error
			control.content[arg], _ = CatFile(control.content[arg], name)
			//content, err := codebase.GetFile(name)
			//if err == nil {
			//	if control.content[arg] != nil {
			//		control.content[arg][name] = *content
			//	} else {
			//		control.content[arg] = map[string]string{}
			//		control.content[arg][name] = *content
			//	}
			//}
		} else {
			//read content to find function
		}

	}
}

func CatFile(controlContent map[string]string, name string) (map[string]string, error) {
	content, err := codebase.GetFile(name)
	if err != nil {
		return controlContent, err
	}

	if controlContent != nil {
		controlContent[name] = *content
	} else {
		controlContent = map[string]string{}
		controlContent[name] = *content
	}

	return controlContent, nil
}

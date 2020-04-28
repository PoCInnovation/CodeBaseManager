package codebase

import (
	"log"
	"strings"
)

func FindParser(name string, control parsingRepo) {
	for _, arg := range control.args {
		splitName := strings.Split(name, "/")
		splitLen := len(splitName)
		if splitLen == 0 {
			log.Printf("Cannot Split %s\n", name)
		}

		if arg == splitName[splitLen-1] {
			// TODO: refacto parsing to use fctPtr -> common ground for cat and find
			control.content[arg], _ = FindFile(control.content[arg], name)
			//if control.content[arg] != nil {
			//	control.content[arg][name] = name
			//} else {
			//	control.content[arg] = map[string]string{}
			//	control.content[arg][name] = name
			//}
		} else {
			//read content to find function
		}
	}
}

func FindFile(controlContent map[string]string, name string) (map[string]string, error) {
	if controlContent != nil {
		controlContent[name] = name
	} else {
		controlContent = map[string]string{}
		controlContent[name] = name
	}
	return controlContent, nil
}

//func CatFile() {
//
//}

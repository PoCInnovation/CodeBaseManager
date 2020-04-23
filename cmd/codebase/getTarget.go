package codebase

import (
	"fmt"
)

type findFctArray struct {
	language string
	fct      targetFctParser
}

type targetFctParser func(fileContent, toFind string) (content *string, err error)

type contentFound map[string]*string

type parsingRepo struct {
	target  targetFctParser
	content contentFound
	args    []string
	found   []bool
}

func ParseRepo(parser *parsingRepo, repo []string) {
	// TODO: change type repo by parsed type
	// TODO: if parsed by modules, open_dir management (change filepath by Dir)
	for _, filePath := range repo {
		fileContent, err := GetFile(filePath)
		if err != nil {
			continue
		}
		fmt.Println(*fileContent)
		//findTargetFromArgs(*fileContent, parser)
		if foundAllArgs(parser.found) {
			return
		}
	}
	//return print content
}

func foundAllArgs(found []bool) bool {
	for _, value := range found {
		if !value {
			return false
		}
	}
	return true
}

func findTargetFromArgs(fileContent string, parser *parsingRepo) {
	for idx, toFind := range parser.args {
		content, err := parser.target(fileContent, toFind)
		if err != nil {
			parser.content[toFind] = content
			parser.found[idx] = true
		}
	}
}

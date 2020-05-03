package codebase

import (
	"fmt"
	"regexp"
	"strings"
)

type findFctArray struct {
	language   string
	extensions []string
	fct        func(string, string) *string
	regex      string
}

//var cppExtensions = []string{".cpp", ".hpp", ".cc", ".hh"}
//var pythonExtension = []string{".py", ".pyc"}

var catTargetFcts = []findFctArray{
	{language: "C", fct: catCFunction,
		extensions: []string{".c", ".h"},
		regex:      "((?m)^(\\w+(\\s+)?){1,3})%s((\\((.*?)\\))(\\s*)\\{(\\s*?.*?)*?\n\\})",
	},
	{language: "Go", fct: catGoFunction,
		extensions: []string{".go"},
		regex:      "(?m)^((\\t| )*?)func %s\\((.+)\\{(\\s*?.*?)*?\n\\}\n",
	},
}

var findTargetFcts = []findFctArray{
	{language: "C", fct: findCFunction,
		extensions: []string{".c", ".h"},
		regex:      "((?m)^(\\w+(\\s+)?){1,3})%s((\\((.*?)\\))(\\s*)\\{(\\s*?.*?)*?\n\\})",
	},
	{language: "Go", fct: findGoFunction,
		extensions: []string{".go"},
		regex:      "(?m)^((\\t| )*?)func %s\\((.+)\\{(\\s*?.*?)*?\n\\}\n",
	},
}

func catTargetFunction(reg, fileContent string) *string {
	r, err := regexp.Compile(reg)
	if err != nil {
		return nil
	}
	if found := r.FindString(fileContent); found != "" {
		return &found
	}
	return nil
}

func findTargetFunction(reg, fileContent, fctName string) *string {
	r, err := regexp.Compile(reg)
	if err != nil {
		return nil
	}
	foundIndex := r.FindStringIndex(fileContent)
	if foundIndex == nil {
		return nil
	}
	line := strings.Count(fileContent[0:foundIndex[0]], "\n") + 1
	found := fmt.Sprintf("Function %v at line : %v", fctName, line)
	return &found
}

func catGoFunction(fileContent, fctName string) *string {
	reg := "(?m)^((\\t| )*?)func " + fctName + "\\((.+)\\{(\\s*?.*?)*?\n\\}\n"
	return catTargetFunction(reg, fileContent)
}

func catCFunction(fileContent, fctName string) *string {
	reg := "((?m)^(\\w+(\\s+)?){1,3})" + fctName + "((\\((.*?)\\))(\\s*)\\{(\\s*?.*?)*?\n\\})"
	return catTargetFunction(reg, fileContent)
}

func findGoFunction(fileContent, fctName string) *string {
	reg := "(?m)^((\\t| )*?)func " + fctName + "\\((.+)\\{(\\s*?.*?)*?\n\\}\n"
	return findTargetFunction(reg, fileContent, fctName)
}

func findCFunction(fileContent, fctName string) *string {
	reg := "((?m)^(\\w+(\\s+)?){1,3})" + fctName + "((\\((.*?)\\))(\\s*)\\{(\\s*?.*?)*?\n\\})"
	return findTargetFunction(reg, fileContent, fctName)
}

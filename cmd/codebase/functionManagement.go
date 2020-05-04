package codebase

import (
	"fmt"
	"regexp"
	"strings"
)

//TODO: transform fct into 2 fields of catFct and findFct
const (
	FIND = 0
	CAT  = 1
)

//TODO: Manage several functions/methods in same file
type findFctArray struct {
	language   string
	extensions []string
	fct        func(string, string) *string
	regex      string
}

//var cppExtensions = []string{".cpp", ".hpp", ".cc", ".hh"}
//var pythonExtension = []string{".py", ".pyc"}

var catTargetFcts = []findFctArray{
	{
		language: "C", fct: catCFunction,
		extensions: []string{".c", ".h"},
		regex:      "((?m)^(\\w+(\\s+)?){1,3})%s((\\((.*?)\\))(\\s*)\\{(\\s*?.*?)*?\n\\})",
	},
	{
		language: "Go", fct: catGoFunction,
		extensions: []string{".go"},
		regex:      "(?m)^((\\t| )*?)func %s\\((.+)\\{(\\s*?.*?)*?\n\\}\n",
	},
	{
		language: "Python", fct: catPythonFunction,
		extensions: []string{".py"},
		regex:      "(((?:^[ \t]*)*@(.*)\n)(?:^[ \t]*)def %s\\(.*\\):(\\s)*((?=.*?[^ \t\n]).*\r?\n?)*\n)",
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
	{
		language: "Python", fct: findPythonFunction,
		extensions: []string{".py"},
		regex:      "(((?:^[ \t]*)*@(.*)\n)(?:^[ \t]*)def %s\\(.*\\):(\\s)*((?=.*?[^ \t\n]).*\r?\n?)*\n)",
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
	reg := fmt.Sprintf("(?m)^((\\t| )*?)func %s\\((.+)\\{(\\s*?.*?)*?\n\\}\n", fctName)
	return catTargetFunction(reg, fileContent)
}

func catCFunction(fileContent, fctName string) *string {
	reg := fmt.Sprintf("((?m)^(\\w+(\\s+)?){1,3})%s((\\((.*?)\\))(\\s*)\\{(\\s*?.*?)*?\n\\})", fctName)
	return catTargetFunction(reg, fileContent)
}

func catPythonFunction(fileContent, fctName string) *string {
	//fmt.Println(fileContent, "\n" + fctName)
	reg := fmt.Sprintf("((((?m)^[ \t]*)@(.*)\n)*((?m)^[ \t]*)def %s\\(.*\\):(\\s)*((.*?[^ \t\n]).*\r?\n?)*)", fctName)
	// TODO: DOING FUCKING CLASS.
	//reg := fmt.Sprintf("((?m)^[ \t]*)class %s(\\(([^\r\n\f\v])*\\))*:\n((.*\\s*?)*)\n*", fctName)
	return catTargetFunction(reg, fileContent)
}

func findGoFunction(fileContent, fctName string) *string {
	reg := "(?m)^((\\t| )*?)func " + fctName + "\\((.+)\\{(\\s*?.*?)*?\n\\}\n"
	return findTargetFunction(reg, fileContent, fctName)
}

func findCFunction(fileContent, fctName string) *string {
	reg := fmt.Sprintf("(?m)^((\\t| )*?)func %s\\((.+)\\{(\\s*?.*?)*?\n\\}\n", fctName)
	return findTargetFunction(reg, fileContent, fctName)
}

func findPythonFunction(fileContent, fctName string) *string {
	//fmt.Println(fileContent, "\n" + fctName)
	reg := fmt.Sprintf("((((?m)^[ \t]*)@(.*)\n)*((?m)^[ \t]*)def %s\\(.*\\):(\\s)*((.*?[^ \t\n]).*\r?\n?)*)", fctName)
	// TODO: DOING FUCKING CLASS.
	//reg := fmt.Sprintf("((?m)^[ \t]*)class %s(\\(([^\r\n\f\v])*\\))*:\n((.*\\s*?)*)\n*", fctName)
	return findTargetFunction(reg, fileContent, fctName)
}

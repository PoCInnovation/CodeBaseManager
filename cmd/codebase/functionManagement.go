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
}

var cExtensions = []string{".c", ".h"}
var cppExtensions = []string{".cpp", ".hpp", ".cc", ".hh"}
var goExtensions = []string{".go"}
var pythonExtension = []string{".py", ".pyc"}

// TODO: manage file extension
// TODO: make const ?
var catTargetFcts = [...]findFctArray{
	{language: "C", extensions: cExtensions, fct: catGoFunction},
	{language: "Go", extensions: goExtensions, fct: catCFunction},
}

// TODO: manage file extension
// TODO: make const ?
var findTargetFcts = [...]findFctArray{
	{language: "C", extensions: cExtensions, fct: findGoFunction},
	{language: "Go", extensions: goExtensions, fct: findCFunction},
}

func catFunction(reg, fileContent string) *string {
	r, err := regexp.Compile(reg)
	if err != nil {
		return nil
	}
	if found := r.FindString(fileContent); found != "" {
		return &found
	}
	return nil
}

func findFunction(reg, fileContent, fctName string) *string {
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
	return catFunction(reg, fileContent)
}

func catCFunction(fileContent, fctName string) *string {
	reg := "((?m)^(\\w+(\\s+)?){1,3})" + fctName + "((\\((.*?)\\))(\\s*)\\{(\\s*?.*?)*?\n\\})"
	return catFunction(reg, fileContent)
}

func findGoFunction(fileContent, fctName string) *string {
	reg := "(?m)^((\\t| )*?)func " + fctName + "\\((.+)\\{(\\s*?.*?)*?\n\\}\n"
	return findFunction(reg, fileContent, fctName)
}

func findCFunction(fileContent, fctName string) *string {
	reg := "((?m)^(\\w+(\\s+)?){1,3})" + fctName + "((\\((.*?)\\))(\\s*)\\{(\\s*?.*?)*?\n\\})"
	return findFunction(reg, fileContent, fctName)
}

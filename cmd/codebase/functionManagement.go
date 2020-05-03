package codebase

import (
	"fmt"
	"regexp"
)

type languageFct func(string, string) *string

type findFctArray struct {
	language   string
	extensions []string
	fct        languageFct
}

var cExtensions = []string{".c", ".h"}
var cppExtensions = []string{".cpp", ".hpp", ".cc", ".hh"}
var goExtensions = []string{".go"}
var pythonExtension = []string{".py", ".pyc"}

// TODO: manage file extension
// TODO: make const ?
var catTargetFcts = []findFctArray{
	{language: "C", extensions: cExtensions, fct: catGoFunction},
	{language: "Go", extensions: goExtensions, fct: catCFunction},
}

// TODO: manage file extension
// TODO: make const ?
var findTargetFcts = []findFctArray{
	{language: "C", extensions: cExtensions, fct: findGoFunction},
	{language: "Go", extensions: goExtensions, fct: findCFunction},
}

func catFunction(reg, fileContent string) *string {
	r, _ := regexp.Compile(reg)
	if !r.MatchString(fileContent) {
		return nil
	}
	found := r.FindString(fileContent)
	return &found
}

func catGoFunction(fileContent, fctName string) *string {
	//TODO: Add begin line management
	reg := "func " + fctName + "\\((.+)\\{(\\s*?.*?)*?\n\\}\n"
	return catFunction(reg, fileContent)
}

func catCFunction(fileContent, fctName string) *string {
	fmt.Println(fctName)
	//TODO: Add begin line management
	reg := "(\\w+(\\s+)?)" + fctName + "\\([^!@#$+%^]+?\\)(\\s*)\\{(\\s*?.*?)*?\n\\}"
	return catFunction(reg, fileContent)
}

func findGoFunction(fileContent, fctName string) *string {
	fmt.Println(fctName)
	r, _ := regexp.Compile("func " + fctName + "\\((.+)\\{(\\s*?.*?)*?\n\\}\n")
	if !r.MatchString(fileContent) {
		return nil
	}
	found := r.FindString(fileContent)
	return &found
}

func findCFunction(fileContent, fctName string) *string {
	fmt.Println(fctName)
	r, _ := regexp.Compile("^(\\w+(\\s+)?)" + fctName + "\\([^!@#$+%^]+?\\)(\\s*)\\{(\\s*?.*?)*?\n\\}")
	if !r.MatchString(fileContent) {
		return nil
	}
	found := r.FindString(fileContent)
	return &found
}

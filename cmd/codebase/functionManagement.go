package codebase

import (
	"fmt"
	"regexp"
)

type languageFct func(string, string) *string

//const catTargetFcts = []findFctArray{
//	{"C", catCFct},
//	{"Go", catGoFct},
//}

//const findTargetFcts = []findFctArray{
//	{"C", findCFct},
//	{"Go", findGoFct},
//}

func catFunction(reg, fileContent string) *string {
	r, _ := regexp.Compile(reg)
	if !r.MatchString(fileContent) {
		return nil
	}
	found := r.FindString(fileContent)
	return &found
}

func catGoFunction(fileContent, fctName string) *string {
	reg := "func " + fctName + "\\((.+)\\{(\\s*?.*?)*?\n\\}\n"
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

package ft_types

import (
	"fmt"
	"io/ioutil"
)

type ftExpected struct {
	Status     int
	Stdout     string
	Stderr     string
	StdoutFile string
	StderrFile string
}

func getFile(fp string) string {
	if fp == "" {
		return ""
	}
	buf, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println("Expected:", err.Error())
		return ""
	}
	return string(buf)
}

func (exp *ftExpected) getExpOut() string {
	if exp.Stdout != "" {
		return exp.Stdout
	}
	return getFile(exp.StdoutFile)
}

func (exp *ftExpected) getExpErr() string {
	if exp.Stderr != "" {
		return exp.Stderr
	}
	return getFile(exp.StderrFile)
}

func (Test *ftExpected) ApplyVars(vars map[string]string) {
	Test.Stderr = ApplyVarsString(Test.Stderr, vars)
	Test.Stdout = ApplyVarsString(Test.Stdout, vars)
	Test.StderrFile = ApplyVarsString(Test.StderrFile, vars)
	Test.StdoutFile = ApplyVarsString(Test.StdoutFile, vars)
}

func (Test *ftExpected) ApplyDefault(reference ftExpected, vars map[string]string) {
	if len(Test.Stderr) == 0 {
		Test.Stderr = reference.Stderr
	}
	if len(Test.Stdout) == 0 {
		Test.Stdout = reference.Stdout
	}
	if Test.Status == 0 && reference.Status != 0 {
		Test.Status = reference.Status
	}
	if len(Test.StderrFile) == 0 {
		Test.StderrFile = reference.StderrFile
	}
	if len(Test.StdoutFile) == 0 {
		Test.StdoutFile = reference.StdoutFile
	}
	Test.ApplyVars(vars)
}
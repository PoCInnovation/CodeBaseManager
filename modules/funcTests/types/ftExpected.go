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

func (test *ftExpected) ApplyDefault(reference ftExpected) {
	if len(test.Stderr) == 0 {
		test.Stderr = reference.Stderr
	}
	if len(test.Stdout) == 0 {
		test.Stdout = reference.Stdout
	}
	if test.Status == 0 && reference.Status != 0 {
		test.Status = reference.Status
	}
	if len(test.StderrFile) == 0 {
		test.StderrFile = reference.StderrFile
	}
	if len(test.StdoutFile) == 0 {
		test.StdoutFile = reference.StdoutFile
	}
}
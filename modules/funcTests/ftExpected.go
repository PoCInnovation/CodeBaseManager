package funcTests

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
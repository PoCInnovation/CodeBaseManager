package funcTests

import (
    "bytes"
    "fmt"
    "os/exec"
)

type ftExecution struct {
    cmd *exec.Cmd
    outBuf bytes.Buffer
    errBuf bytes.Buffer
}

func (e *ftExecution) Set(bin string, args ...string) {
    e.cmd = exec.Command(bin, args...)
    e.cmd.Stdout = &e.outBuf
    e.cmd.Stderr = &e.errBuf
}

func (e *ftExecution) Run() {
    if err := e.cmd.Run(); err != nil {
        fmt.Println(err)
    }
}
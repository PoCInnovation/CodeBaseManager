package funcTests

import (
    "bytes"
    "fmt"
    "os"
    "os/exec"
)

type ftExecution struct {
    cmd *exec.Cmd
    outBuf bytes.Buffer
    errBuf bytes.Buffer
}

func (e *ftExecution) Set(inter *ftInteractions, bin string, args ...string) {
    e.cmd = exec.Command(bin, args...)

    e.cmd.Stdout = &e.outBuf
    e.cmd.Stderr = &e.errBuf

    if inter.Env != nil {
        e.cmd.Env = inter.Env
    } else {
        e.cmd.Env = os.Environ()
    }
    e.cmd.Env = append(e.cmd.Env, inter.AddEnv...)

    // TODO: - stdin
    //       - pipes

}

func (e *ftExecution) Run() {
    if err := e.cmd.Run(); err != nil {
        fmt.Println(err)
    }
}

func QuickRun(cmdLine string) error {
    cmd := exec.Command("bash", "-c", cmdLine)
    _, err := cmd.CombinedOutput()
    return err
}
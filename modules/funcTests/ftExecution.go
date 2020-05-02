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

func (e *ftExecution) setEnv(env, addEnv []string) {
    if env != nil {
        e.cmd.Env = env
    } else {
        e.cmd.Env = os.Environ()
    }
    //TODO: See for variable replacement
    e.cmd.Env = append(e.cmd.Env, addEnv...)
}

func (e *ftExecution) setStdin(stdin, stdinFile string) {
    if stdin != noCmd {
        e.cmd.Stdin = bytes.NewReader([]byte(stdin))
    } else if stdinFile != noCmd {
        file, err := os.Open(stdinFile)
        if err != nil {
            fmt.Println(err)
            // TODO: Handle that way better.
            return
        }
        e.cmd.Stdin = file
    }
}

func (e *ftExecution) Set(inter *ftInteractions, bin string, args ...string) {
    e.cmd = exec.Command(bin, args...)

    //TODO: Handle stdoutPipe
    e.cmd.Stdout = &e.outBuf
    //TODO: Handle stderrPipe
    e.cmd.Stderr = &e.errBuf
    e.setStdin(inter.Stdin, inter.StdinFile)

    e.setEnv(inter.Env, inter.AddEnv)

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
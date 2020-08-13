package funcTests

import (
	"bytes"
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/REPL"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type ftInteractions struct {
	StdoutPipe string
	StderrPipe string
	Stdin      string
	StdinFile  string

	Pre    string
	Post   string
	Env    []string
	AddEnv []string
}

type ftExecution struct {
	cmd      *exec.Cmd
	outBuf   bytes.Buffer
	errBuf   bytes.Buffer
	execTime time.Time
	status   int
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
			fmt.Println("setStdin:", err)
			// TODO: Handle that way better.
			return
		}
		e.cmd.Stdin = file
	}
}

func (e *ftExecution) Set(inter *ftInteractions, bin string, args ...string) {
	bin = REPL.LocateBinary(bin)
	e.cmd = exec.Command(bin, args...)

	//TODO: Handle stdoutPipe
	e.cmd.Stdout = &e.outBuf
	//TODO: Handle stderrPipe
	e.cmd.Stderr = &e.errBuf
	e.setStdin(inter.Stdin, inter.StdinFile)

	e.setEnv(inter.Env, inter.AddEnv)
}

func (e *ftExecution) Run(options ftOptions) {
	e.execTime = time.Now()
	if err := e.cmd.Run(); err != nil {
		if !strings.Contains(err.Error(), "exit status") {
			fmt.Println("Run:", err)
		}
		if exitError, ok := err.(*exec.ExitError); ok {
			e.status = exitError.Sys().(syscall.WaitStatus).ExitStatus()
		}
	} else {
		e.status = e.cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	}
	if options.Time {
		fmt.Println(time.Since(e.execTime))
	}
}

func (e *ftExecution) AfterPipe(outPipeCmd, errPipeCmd string) {
	if outPipeCmd != "" {
		tmp := exec.Command("bash", "-c", outPipeCmd)
		tmp.Stdin = bytes.NewReader(e.outBuf.Bytes())
		e.outBuf.Reset()
		tmp.Stdout = &e.outBuf
		if err := tmp.Run(); err != nil {
			fmt.Println("stdoutPipe:", err)
			return
		}
	}
	if errPipeCmd != "" {
		tmp := exec.Command("bash", "-c", errPipeCmd)
		tmp.Stdin = bytes.NewReader(e.errBuf.Bytes())
		e.errBuf.Reset()
		tmp.Stdout = &e.errBuf
		if err := tmp.Run(); err != nil {
			fmt.Println("stderrPipe:", err)
			return
		}
	}
}

func QuickRun(cmdLine string) error {
	cmd := exec.Command("bash", "-c", cmdLine)
	_, err := cmd.CombinedOutput()
	return err
}

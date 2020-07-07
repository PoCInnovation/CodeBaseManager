package ft_types

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/REPL"
	"github.com/PoCFrance/CodeBaseManager/modules/logs"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type ftExecution struct {
	cmd *exec.Cmd
	outBuf bytes.Buffer
	errBuf bytes.Buffer
	timeoutChan <-chan time.Time
	execChan chan error
	status int
	timeout bool
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

func (e *ftExecution) setStdin(stdin, stdinFile, stdinPipe string) {
	if stdin != "" {
		e.cmd.Stdin = bytes.NewReader([]byte(stdin))
	} else if stdinFile != "" {
		file, err := os.Open(stdinFile)
		if err != nil {
			fmt.Println("setStdin:", err)
			// TODO: Handle that way better.
			return
		}
		e.cmd.Stdin = file
	} else if stdinPipe != "" {
		tmp := exec.Command("bash", "-c", stdinPipe)
		output, err := tmp.Output()
		if err != nil {
			fmt.Println("stdinPipe:", err)
			return
		}
		e.cmd.Stdin = bytes.NewReader(output)
	}
}

func (e *ftExecution) Set(inter ftInteractions, bin string, timeString string, args ...string, ) {
	bin = REPL.LocateBinary(bin)
	e.cmd = exec.Command(bin, args...)

	e.cmd.Stdout = &e.outBuf
	e.cmd.Stderr = &e.errBuf

	//Setting channels
	if timeout, err := time.ParseDuration(timeString); err != nil {
		logs.CBMLogs.Error(err)
	} else {
		if timeString != "0.000000s" {
			e.timeoutChan = time.After(timeout)
		} else {
			e.timeoutChan = nil
		}
	}
	e.execChan = make(chan error)

	e.setStdin(inter.Stdin, inter.StdinFile, inter.StdinPipe)
	e.setEnv(inter.Env, inter.AddEnv)
}

func QuickRun(cmdLine string) error {
	cmd := exec.Command("bash", "-c", cmdLine)
	_, err := cmd.CombinedOutput()
	return err
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

func (Test *ftDescription) Run() (ftExecution, error) {
	execInfo := ftExecution{}

	execInfo.Set(Test.Interactions, Test.Bin, fmt.Sprintf("%fs", Test.Options.Timeout), Test.Args...)

	if err := execInfo.cmd.Start(); err != nil {
		return ftExecution{}, errors.New("Failed to start the cmd")
	}

	go func() { execInfo.execChan <- execInfo.cmd.Wait() }()

	select {
	case <- execInfo.timeoutChan:
		execInfo.cmd.Process.Kill()
		return ftExecution{timeout: true}, nil
	case err := <- execInfo.execChan:
		if err == nil {
			execInfo.status = execInfo.cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
		} else {
			if !strings.Contains(err.Error(), "exit status") {
				return ftExecution{}, err
			}
			if exitError, ok := err.(*exec.ExitError); ok {
				execInfo.status = exitError.Sys().(syscall.WaitStatus).ExitStatus()
			}
		}
	}
	return execInfo, nil
}
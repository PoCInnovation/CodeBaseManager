package REPL

import (
	"bufio"
	"fmt"
	"github.com/logrusorgru/aurora"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type mux struct {
	interrupt chan os.Signal
	msg       chan string
}

func (c *mux) init() {
	c.interrupt = make(chan os.Signal)
	c.msg = make(chan string)
	signal.Notify(c.interrupt, syscall.SIGINT)
}

func (c *mux) Close() {
	close(c.interrupt)
	close(c.msg)
	signal.Reset()
}

type prompt struct {
	comm   mux
	module string
	reader *bufio.Reader
}

func checkErr(err error) string {
	if err == io.EOF {
		fmt.Println("Exit.")
		return "exit"
	} else {
		return ""
	}
}

func (p *prompt) readInput() {
	for {
		if in, err := p.reader.ReadString('\n'); err != nil {
			p.comm.msg <- checkErr(err)
		} else {
			p.comm.msg <- in
		}
	}
}

func newPrompt(module string) *prompt {
	p := &prompt{module: module}
	p.reader = bufio.NewReader(os.Stdin)

	go p.readInput()
	p.comm.init()
	return p
}

func getCurrentDir() string {
	pwd := os.Getenv("PWD")
	return pwd[strings.LastIndex(pwd, "/")+1:]
}

func (p *prompt) Display() {
	sep := aurora.Green("→")
	cwd := aurora.BrightBlue(getCurrentDir())
	module := aurora.BrightBlue(p.module)

	fmt.Printf("%s %s %s %s ", cwd, sep, module, sep)
}

func (p *prompt) GetInput() string {
	select {
	case in := <-p.comm.msg:
		return in
	case <-p.comm.interrupt:
		fmt.Println("")
		return ""
	}
}

func (p *prompt) Close() {
	p.comm.Close()
}

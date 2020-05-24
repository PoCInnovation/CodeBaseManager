package REPL

import (
	"os"
	"os/signal"
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

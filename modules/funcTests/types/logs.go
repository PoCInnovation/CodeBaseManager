package ft_types

import (
    "fmt"
    "github.com/logrusorgru/aurora"
    "os"
    "reflect"
)

type FTLog uint8

const (
    InfoLog FTLog = iota
    WarnLog
    ErrLog
)

var levelNames = []aurora.Value{
    InfoLog: aurora.Blue("[Info]"),
    WarnLog: aurora.Yellow("[Warn]"),
    ErrLog:  aurora.Red("[Error]"),
}

func NewLogger(level FTLog) FTLog {
    return level
}

func (l *FTLog) Set(level FTLog) {
    *l = level
}

func (l FTLog) log(args []interface{}, req FTLog) {
    if l > req {
        return
    }
    fmt.Fprintf(os.Stderr, "%s ", levelNames[req])
    format := reflect.ValueOf(args[0])
    if len(args) > 1 {
        fmt.Fprintf(os.Stderr, format.String() + "\n", args[1:]...)
    } else {
        fmt.Fprintf(os.Stderr, format.String() + "\n")
    }
}

func (l FTLog) Info(args ...interface{}) {
    l.log(args, InfoLog)
}

func (l FTLog) Warn(args ...interface{}) {
    l.log(args, WarnLog)
}

func (l FTLog) Error(args ...interface{}) {
    l.log(args, ErrLog)
}
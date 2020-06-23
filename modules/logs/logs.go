package logs

import (
    "fmt"
    "github.com/logrusorgru/aurora"
    "os"
    "reflect"
)

type logLevel int

type FTLog struct {
    level logLevel
    writer *os.File
    init bool
}

const (
    DebugLog logLevel = iota
    WarnLog
    ErrLog
)


var levelNames = map[string]logLevel{
    "debug": DebugLog,
    "warn": WarnLog,
    "error": ErrLog,
}

var levelDisplay = []aurora.Value{
    DebugLog: aurora.Blue("[Debug]"),
    WarnLog:  aurora.Yellow("[Warn]"),
    ErrLog:   aurora.Red("[Error]"),
}

var Verbosity string
var LogsFP string
var CBMLogs FTLog

func InitCBMLogs(verbosity, fpLogs string) {
    if CBMLogs.init == true {
        return
    }
    level, ok := levelNames[verbosity]
    if ok {
        CBMLogs.level = level
    }
    if fpLogs != "stderr" {
        CBMLogs.SetWriter(fpLogs)
    }
    CBMLogs.init = true
}

func (l *FTLog) Close() {
    if l.writer != os.Stderr {
        l.writer.Close()
    }
}

func (l *FTLog) SetLevel(level int) {
    l.level = logLevel(level)
}

func (l *FTLog) SetWriter(fp string) {
    file, err := os.Open(fp)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Logs:", err)
    } else {
        CBMLogs.writer = file
    }
}

func (l FTLog) log(args []interface{}, req logLevel) {
    if l.level > req {
        return
    }
    fmt.Fprintf(os.Stderr, "%s ", levelDisplay[req])
    format := reflect.ValueOf(args[0])
    if len(args) > 1 {
        fmt.Fprintf(os.Stderr, format.String() + "\n", args[1:]...)
    } else {
        fmt.Fprintf(os.Stderr, format.String() + "\n")
    }
}

func (l FTLog) Debug(args ...interface{}) {
    l.log(args, DebugLog)
}

func (l FTLog) Warn(args ...interface{}) {
    l.log(args, WarnLog)
}

func (l FTLog) Error(args ...interface{}) {
    l.log(args, ErrLog)
}
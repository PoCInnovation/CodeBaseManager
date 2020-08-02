package watcher

import (
    "fmt"
    "log"
    "time"
)

func Run(stop chan struct{}) {
    for {
        select {
        case <-stop:
            log.Printf("Watcher's shutting down\n")
            return
        default:
            fmt.Println("hello i'm the watcher, I do shit atm")
            time.Sleep(5 * time.Second)
        }
    }
}
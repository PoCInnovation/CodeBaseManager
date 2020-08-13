package watcher

import (
	"github.com/rjeczalik/notify"
	"log"
	"time"
)

// Documentation on notify: https://godoc.org/github.com/rjeczalik/notify#example-Watch--LinuxMove

const watchedEvents = notify.InCreate | notify.InDelete | notify.InDeleteSelf | notify.InModify | notify.InMovedFrom | notify.InMovedTo | notify.InMoveSelf

func start() chan notify.EventInfo {
	log.Println("Initializing watcher")
	// TODO: Load config file?
	// TODO: Get repository list (based on config?)
	watcher := make(chan notify.EventInfo, 2)
	// TODO: Set watchpoint for everything (except .git, ...) (or based on config?)
	return watcher
}

func Run(stop chan struct{}) {
	watcher := start()
	defer notify.Stop(watcher)
	for {
		select {
		case <-stop:
			log.Println("Watcher's shutting down")
			return
		default:
			for events := range watcher {
				log.Println("received:", events)
			}
			time.Sleep(5 * time.Second)
		}
	}
}

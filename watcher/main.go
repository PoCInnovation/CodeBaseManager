package main

import (
	"github.com/PoCInnovation/CodeBaseManager/watcher/watcher"
	"log"
	"time"
)

func main() {
	if err := watcher.NewWatcher(); err != nil {
		log.Fatalln(err)
	}
	for {
		if err := watcher.GetProjectList(); err != nil {
			log.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}

// Documentation on notify: https://godoc.org/github.com/rjeczalik/notify#example-Watch--LinuxMove
//
//const (
//    watchedEvents = notify.InCreate | notify.InDelete | notify.InDeleteSelf | notify.InModify | notify.InMovedFrom | notify.InMovedTo | notify.InMoveSelf
//    configFile = "watcher.toml"
//    configGlobal = "/" + configFile
//    configLocal = ".cbm/" + configFile
//)
//
//func isIgnored(name string, in []string) bool {
//    for _, v := range in {
//        if v == name {
//            return true
//        }
//    }
//    return false
//}
//
//func isBanned(name string) bool {
//    return name == ".git" || name == ".vscode" || name == ".idea"
//}
//
//func loadProjects(projects []model.Project, cfg *config, c chan notify.EventInfo) {
//    for _, proj := range projects {
//
//        log.Println("Loading " + proj.Name)
//        if isIgnored(proj.Name, cfg.Ignored) || isBanned(proj.Name){
//            continue
//        }
//
//        if err := filepath.Walk(proj.Path, func(name string, info os.FileInfo, err error) error {
//            log.Println("Im here " + name)
//            if isIgnored(name, cfg.Ignored) || isBanned(name) {
//                return fmt.Errorf("%s: ignored\n%v", name, err)
//            }
//            if tmpErr := notify.Watch(name, c, watchedEvents); tmpErr != nil {
//                return fmt.Errorf("failed to add watchpoint for '%s': %v\n%v", name, tmpErr, err)
//            }
//            return err
//        }); err != nil {
//            log.Println(err)
//        }
//    }
//}
//
//func start() (chan notify.EventInfo, error) {
//    log.Println("Initializing Watcher")
//
//    cfg, err := loadConfig()
//    if err != nil {
//        return nil, fmt.Errorf("Watcher failed to start: %v", err)
//    }
//
//    projects, err := controllers.ListProjects()
//    ch := make(chan notify.EventInfo, 2)
//    if len(projects) == 0 {
//        return ch, nil
//    } else if err != nil && len(projects) != 0 {
//        return nil, fmt.Errorf("I ate my life: %v", err)
//    }
//
//    loadProjects(projects, cfg, ch)
//    return ch, nil
//}
//
//func Run(stop chan struct{}) {
//    ch, err := start()
//    if err != nil {
//        log.Fatalln(err)
//    }
//    defer notify.Stop(ch)
//    for {
//        select {
//        case <-stop:
//            log.Println("Watcher's shutting down")
//            return
//        default:
//            for events := range ch {
//                log.Println("received:", events)
//            }
//            time.Sleep(5 * time.Second)
//        }
//    }
//}

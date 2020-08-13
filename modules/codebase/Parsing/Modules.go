package codebase

import (
	Components "github.com/PoCFrance/CodeBaseManager/modules/codebase/Components"
	"log"
	"strings"
	"sync"
)

func ProcessModules(moduleList *Components.Repository, module *Components.Module, path string) {
	targetList, closure := listTargets(path)
	if targetList == nil {
		return
	}
	defer closure()

	var wg sync.WaitGroup

	for _, target := range targetList {
		//targetName := target.Name()
		//newPath := path + "/" + targetName
		if target.IsDir() {
			targetName := target.Name()
			newPath := path + "/" + targetName
			if strings.HasPrefix(targetName, ".") {
				continue
			}
			newModule := moduleList.Append(newPath, targetName)
			wg.Add(1)
			go func() {
				ProcessModules(moduleList, newModule, newPath)
				log.Println("module", targetName, "processed")
				wg.Done()
			}()
		} else {
			targetName := target.Name()
			newPath := path + "/" + targetName
			if isNotReadable(newPath) {
				continue
			}
			newFile := module.Append(newPath, targetName)
			wg.Add(1)
			go func() {
				parseFile(newFile, newPath)
				wg.Done()
			}()
		}
	}

	wg.Wait()
}

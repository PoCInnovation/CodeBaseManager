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
		targetName := target.Name()
		newPath := path + "/" + targetName
		if target.IsDir() {
			if strings.HasPrefix(targetName, ".") {
				continue
			}
			targetIsModule(moduleList, newPath, targetName, &wg)
		} else {
			if isNotReadable(newPath) {
				continue
			}
			targetIsFile(module, newPath, targetName, &wg)
		}
	}
	wg.Wait()
}

func targetIsModule(moduleList *Components.Repository, path, name string, wg *sync.WaitGroup) {
	newModule := moduleList.Append(path, name)
	wg.Add(1)
	go func() {
		ProcessModules(moduleList, newModule, path)
		log.Println("module", name, "processed")
		wg.Done()
	}()
}

func targetIsFile(module *Components.Module, path, name string, wg *sync.WaitGroup) {
	newFile := module.Append(path, name)
	wg.Add(1)
	go func() {
		parseFile(newFile, path)
		wg.Done()
	}()
}

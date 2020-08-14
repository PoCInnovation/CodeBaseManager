package codebase

import (
	"log"
	"strings"
	"sync"
)

func ProcessModules(moduleList *Repository, module *Module, path string) {
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
			if isNotReadable(newPath) || !isSupportedLanguage(newPath) {
				continue
			}
			targetIsFile(module, newPath, targetName, &wg)
		}
	}
	wg.Wait()
}

func targetIsModule(moduleList *Repository, path, name string, wg *sync.WaitGroup) {
	newModule := moduleList.Append(path, name)
	wg.Add(1)
	go func() {
		ProcessModules(moduleList, newModule, path)
		log.Println("module", name, "processed")
		wg.Done()
	}()
}

func targetIsFile(module *Module, path, name string, wg *sync.WaitGroup) {

	newFile := module.Append(path, name)
	wg.Add(1)
	go func() {
		parseFile(newFile, path)
		wg.Done()
	}()
}

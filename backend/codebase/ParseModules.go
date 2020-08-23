package codebase

import (
	"log"
	"strings"
	"sync"
)

func ProcessModules(moduleList *Repository, module *Module, path string) {
	targetList, closure := listTargets(path)
	defer closure()
	if targetList == nil {
		return
	}

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
	newModule := &Module{
		Path:  path,
		Name:  name,
		Files: nil,
	}
	//newModule := moduleList.Append(path, name)
	wg.Add(1)
	go func() {
		ProcessModules(moduleList, newModule, path)
		if !newModule.IsEmpty() {
			moduleList.Modules = append(moduleList.Modules, newModule)
			log.Println("module", name, "processed and added")
		} else {
			log.Println("module", name, "processed but not added")
		}
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

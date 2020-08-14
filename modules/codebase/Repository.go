package codebase

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Repository struct {
	Modules []*Module
}

func (r Repository) isTarget(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Printf("Error when stating %s, %v\n", path, err)
		return false
	}
	name := fileInfo.Name()
	if !fileInfo.IsDir() || strings.HasPrefix(name, ".") {
		return false
	}
	return true
}

func (r *Repository) Append(path, name string) *Module {
	newModule := &Module{
		Path:  path,
		Name:  name,
		Files: nil,
	}
	r.Modules = append(r.Modules, newModule)
	return newModule
}

func (r *Repository) remove(idx, len int) {
	r.Modules[idx] = r.Modules[len-1]
	r.Modules = r.Modules[:len-1]
}

func (r *Repository) Clean() {
	stop := false
main:
	for !stop {
		modulesLen := len(r.Modules)
		for idx, module := range r.Modules {
			if module.IsEmpty() {
				r.remove(idx, modulesLen)
				modulesLen -= 1
				continue main
			}
		}
		stop = true
	}
}

func (r Repository) String() string {
	var str string
	for idx, module := range r.Modules {
		str += fmt.Sprintf("%d: %s\n", idx+1, module.String())
	}
	return str
}

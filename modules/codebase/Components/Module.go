package codebase

import (
	"fmt"
)

type Module struct {
	Path, Name string
	Files      []*FileParser
}

func (m *Module) Append(path, name string) *FileParser {
	newFile := &FileParser{
		Path: path,
		Name: name,
	}
	m.Files = append(m.Files, newFile)
	return newFile
}

func (m Module) String() string {
	str := "Module name: " + m.Name + ", " + "Module path: " + m.Path + "\n"
	for _, file := range m.Files {
		str += fmt.Sprintf("\t%s\n", file.String())
	}
	return str
}

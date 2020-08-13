package codebase

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type TargetChecker interface {
	isTarget(string) bool
}

type ProcessFuncGetter interface {
	getProcessFunc() func(datas RepositoryParser, path string)
}

type Appender interface {
	Append(string) RepositoryParser
}

type RepositoryParser interface {
	TargetChecker
	Appender
	ProcessFuncGetter
}

type FileParser struct {
	Path string
}

func (f FileParser) isTarget(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Printf("Error when stating %s, %v\n", path, err)
		return false
	}
	if fileInfo.IsDir() {
		return false
	}
	return true
}

func (f *FileParser) Append(path string) RepositoryParser {
	//f.Path = append(f.Path, path)
	return f
}

func (f *FileParser) getProcessFunc() func(datas RepositoryParser, path string) {
	return ParseFiles
}

func (f FileParser) String() string {
	str := "FileName: " + f.Path
	return str
}

type ModuleParser struct {
	Path  string
	Files []FileParser
}

func (m ModuleParser) isTarget(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Printf("Error when stating %s, %v\n", path, err)
		return false
	}
	if fileInfo.IsDir() {
		return false
	}
	return true
}

func (m *ModuleParser) getProcessFunc() func(datas RepositoryParser, path string) {
	return ParseFiles
}

func (m *ModuleParser) Append(path string) RepositoryParser {
	newFile := FileParser{
		Path: path,
	}
	m.Files = append(m.Files, newFile)
	return &newFile
}

func (m ModuleParser) String() string {
	str := fmt.Sprintf("ModuleName: %s", m.Path)
	for _, file := range m.Files {
		str += fmt.Sprintf("\t%s\n", file.String())
		//log.Println(file.String())
		//str = str + file.String() + "\n"
	}
	return str
}

type RepositoryContent struct {
	Modules []ModuleParser
}

func (m RepositoryContent) isTarget(path string) bool {
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

func (m *RepositoryContent) getProcessFunc() func(datas RepositoryParser, path string) {
	return repositoryParser
}

func (m *RepositoryContent) Append(path string) RepositoryParser {
	newModule := ModuleParser{
		Path:  path,
		Files: nil,
	}
	m.Modules = append(m.Modules, newModule)
	return &newModule
}

func (m RepositoryContent) String() string {
	//line := strings.Repeat("---", 5)
	//str := line + "\n"

	var str string
	for idx, module := range m.Modules {
		str += fmt.Sprintf("%d: %s\n", idx+1, module.String())
	}

	//str += line
	return str
}

func ParseRepositoryV2(path string) (*RepositoryContent, error) {
	parser := &RepositoryContent{
		Modules: []ModuleParser{},
	}
	if !parser.isTarget(path) {
		return nil, errors.New(path + " is not a valid target")
	}
	mod := parser.Append(path)
	log.Println(mod)
	repositoryParser(mod, path)
	//repositoryParser(parser, path)

	return parser, nil
}

func repositoryParser(datas RepositoryParser, path string) {
	dir, err := os.Open(path)
	if err != nil {
		log.Printf("Error when opening %s, %v\n", path, err)
		return
	}
	defer func() {
		err := dir.Close()
		if err != nil {
			log.Printf("Cannot close :%s, %v\n", path, err)
		}
	}()

	newPaths, err := dir.Readdir(0)
	if err != nil {
		log.Printf("Error when Readdir %s, %v\n", path, err)
		return
	}
	var wg sync.WaitGroup
	processFunc := datas.getProcessFunc()
	for _, newPath := range newPaths {
		newName := path + "/" + newPath.Name()
		if datas.isTarget(newName) {
			log.Println(newName)
			newData := datas.Append(newName)
			wg.Add(1)
			go func() {
				processFunc(newData, newName)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	log.Println(datas)
}

func ParseFiles(datas RepositoryParser, path string) {
	if !datas.isTarget(path) {
		log.Println("file ", path, " is not a target")
		return
	}
	log.Println("file", path, "processed")
	//datas.Append(path)

	_, err := getFile(path)
	if err != nil {
		return
	}

	//content, err := getFile(path)
	//if err != nil {
	//	return
	//}
	//log.Println(content)
}

func getFile(fileName string) (string, error) {
	fh, err := os.Open(fileName)

	defer func() {
		if err = fh.Close(); err != nil {
			log.Printf("Error on closing file %s, %v\n", fileName, err)
		}
	}()

	if err != nil {
		log.Printf("Error when opening file %s, %v\n", fileName, err)
		return "", err
	}

	scanner := bufio.NewScanner(fh)
	var content string

	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error on closing file %s %v\n", fileName, err)
		return "", err
	}

	return content, nil
}

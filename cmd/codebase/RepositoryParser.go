package codebase

import (
	"bufio"
	"errors"
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
	Append(string)
}

type RepositoryParser interface {
	TargetChecker
	Appender
	ProcessFuncGetter
}

type FileParser struct {
	Files []string
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

func (f *FileParser) Append(path string) {
	f.Files = append(f.Files, path)
}

func (f *FileParser) getProcessFunc() func(datas RepositoryParser, path string) {
	return ParseFiles
}

type ModuleParser struct {
	Modules []string
	Files   []FileParser
}

func (m ModuleParser) isTarget(path string) bool {
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

func (m *ModuleParser) getProcessFunc() func(datas RepositoryParser, path string) {
	return repositoryParser
}

func (m *ModuleParser) Append(path string) {
	m.Modules = append(m.Modules, path)
}

type RepositoryContent []ModuleParser

func ParseRepositoryv2(path string) (*ModuleParser, error) {
	parser := &ModuleParser{
		Modules: []string{},
		Files:   []FileParser{},
	}
	if !parser.isTarget(path) {
		return nil, errors.New(path + " is not a valid target")
	}
	repositoryParser(parser, path)

	for _, module := range parser.Modules {
		newFileParser := FileParser{Files: []string{}}
		repositoryParser(&newFileParser, module)
		parser.Files = append(parser.Files, newFileParser)
	}
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
			datas.Append(newName)
			wg.Add(1)
			go func() {
				processFunc(datas, newName)
				wg.Done()
			}()
		}
	}
	wg.Wait()
}

func ParseFiles(datas RepositoryParser, path string) {
	if !datas.isTarget(path) {
		log.Println(path, " is not a target")
		return
	}
	content, err := getFile(path)
	if err != nil {
		return
	}
	log.Println(content)
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

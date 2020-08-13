package codebase

//import (
//	"bufio"
//	"errors"
//	"fmt"
//	"log"
//	"os"
//	"path/filepath"
//	"strings"
//	"sync"
//)
//
//type TargetChecker interface {
//	isTarget(string) bool
//}
//
//type ProcessFuncGetter interface {
//	getProcessFunc() func(datas RepositoryParser, path string)
//}
//
//type Appender interface {
//	Append(path, name string) RepositoryParser
//}
//
//type RepositoryParser interface {
//	TargetChecker
//	Appender
//	ProcessFuncGetter
//}
//
//type TargetAppender interface {
//	AppendTarget(path, name string) (bool, RepositoryParser)
//}
//
////type RepositoryParser interface {
////	TargetAppender
////	ProcessFuncGetter
////}
//
//type FileParser struct {
//	Path, Name string
//}
//
//func (f FileParser) isTarget(path string) bool {
//	fileInfo, err := os.Stat(path)
//	if err != nil {
//		log.Printf("Error when stating %s, %v\n", path, err)
//		return false
//	}
//	if fileInfo.IsDir() {
//		return false
//	}
//	return true
//}
//
//func (f *FileParser) Append(path, name string) RepositoryParser {
//	//f.Path = append(f.Path, path)
//	return f
//}
//
//func (f *FileParser) getProcessFunc() func(datas RepositoryParser, path string) {
//	return ParseFiles
//}
//
//func (f FileParser) String() string {
//	str := "File name: " + f.Name + ", " + "File path: " + f.Path
//	return str
//}
//
//type Module struct {
//	Path, Name  string
//	Files []FileParser
//}
//
//func (m Module) isTarget(path string) bool {
//	fileInfo, err := os.Stat(path)
//	if err != nil {
//		log.Printf("Error when stating %s, %v\n", path, err)
//		return false
//	}
//	if fileInfo.IsDir() {
//		return false
//	}
//	return true
//}
//
//func (m *Module) getProcessFunc() func(datas RepositoryParser, path string) {
//	return ParseFiles
//}
//
//func (m *Module) Append(path, name string) RepositoryParser {
//	newFile := FileParser{
//		Path: path,
//		Name: name,
//	}
//	m.Files = append(m.Files, newFile)
//	return &newFile
//}
//
//func (m Module) String() string {
//	str := "Module name: " + m.Name + ", " + "Module path: " + m.Path + "\n"
//	for _, file := range m.Files {
//		str += fmt.Sprintf("\t%s\n", file.String())
//	}
//	return str
//}
//
//type Repository struct {
//	Modules []*Module
//}
//
//func (m Repository) isTarget(path string) bool {
//	fileInfo, err := os.Stat(path)
//	if err != nil {
//		log.Printf("Error when stating %s, %v\n", path, err)
//		return false
//	}
//	name := fileInfo.Name()
//	if !fileInfo.IsDir() || strings.HasPrefix(name, ".") {
//		return false
//	}
//	return true
//}
//
//func (m *Repository) getProcessFunc() func(datas RepositoryParser, path string) {
//	return repositoryParser
//}
//
//func (m *Repository) Append(path, name string) RepositoryParser {
//	newModule := &Module{
//		Path:  path,
//		Name: name,
//		Files: nil,
//	}
//	m.Modules = append(m.Modules, newModule)
//	return newModule
//}
//
//func (m Repository) String() string {
//	//line := strings.Repeat("---", 5)
//	//str := line + "\n"
//
//	var str string
//	for idx, module := range m.Modules {
//		str += fmt.Sprintf("%d: %s\n", idx+1, module.String())
//	}
//
//	//str += line
//	return str
//}
//
//func (m Repository) AppendTarget(path string) (bool, RepositoryParser) {
//	fileInfo, err := os.Stat(path)
//	if err != nil {
//		log.Printf("Error when stating %s, %v\n", path, err)
//		return false, nil
//	}
//	name := fileInfo.Name()
//	if fileInfo.IsDir() {
//		if strings.HasPrefix(name, ".") {
//			return false, nil
//		}
//		newModule := &Module{
//			Path:  path,
//			Name: name,
//			Files: nil,
//		}
//		m.Modules = append(m.Modules, newModule)
//		return true, newModule
//	}
//	newFile := FileParser{
//		Path: path,
//		Name: name,
//	}
//	m.Files = append(m.Files, newFile)
//	return &newFile
//	return true, nil
//}
//
//func ParseRepositoryV2(path string) (*Repository, error) {
//	parser := &Repository{
//		Modules: []*Module{},
//	}
//	path = strings.TrimSuffix(path, "/")
//	if !parser.isTarget(path) {
//		return nil, errors.New(path + " is not a valid target")
//	}
//	mod := parser.Append(path, filepath.Base(path))
//	repositoryParser(mod, path)
//	repositoryParser(parser, path)
//
//	return parser, nil
//}
//
//func repositoryParser(datas RepositoryParser, path string) {
//	dir, err := os.Open(path)
//	if err != nil {
//		log.Printf("Error when opening %s, %v\n", path, err)
//		return
//	}
//	defer func() {
//		if err := dir.Close(); err != nil {
//			log.Printf("Cannot close :%s, %v\n", path, err)
//		}
//	}()
//
//	targetList, err := dir.Readdir(0)
//	if err != nil {
//		log.Printf("Error when Readdir %s, %v\n", path, err)
//		return
//	}
//
//	var wg sync.WaitGroup
//	processFunc := datas.getProcessFunc()
//	for _, target := range targetList {
//		targetName := target.Name()
//		newPath := path + "/" + targetName
//		if datas.isTarget(newPath) {
//			//log.Println(newPath)
//			newData := datas.Append(newPath, targetName)
//			wg.Add(1)
//			go func() {
//				processFunc(newData, newPath)
//				wg.Done()
//			}()
//		}
//	}
//
//	wg.Wait()
//}
////func repositoryParser(datas RepositoryParser, path string) {
////	dir, err := os.Open(path)
////	if err != nil {
////		log.Printf("Error when opening %s, %v\n", path, err)
////		return
////	}
////	defer func() {
////		if err := dir.Close(); err != nil {
////			log.Printf("Cannot close :%s, %v\n", path, err)
////		}
////	}()
////
////	targetList, err := dir.Readdir(0)
////	if err != nil {
////		log.Printf("Error when Readdir %s, %v\n", path, err)
////		return
////	}
////
////	var wg sync.WaitGroup
////	processFunc := datas.getProcessFunc()
////	for _, target := range targetList {
////		targetName := target.Name()
////		newPath := path + "/" + targetName
////		if datas.isTarget(newPath) {
////			//log.Println(newPath)
////			newData := datas.Append(newPath, targetName)
////			wg.Add(1)
////			go func() {
////				processFunc(newData, newPath)
////				wg.Done()
////			}()
////		}
////	}
////
////	wg.Wait()
////}
//
//func ParseFiles(datas RepositoryParser, path string) {
//	if !datas.isTarget(path) {
//		log.Println("file ", path, " is not a target")
//		return
//	}
//	log.Println("file", path, "processed")
//	//datas.Append(path)
//
//	_, err := getFile(path)
//	if err != nil {
//		return
//	}
//
//	//content, err := getFile(path)
//	//if err != nil {
//	//	return
//	//}
//	//log.Println(content)
//}
//
//func getFile(fileName string) (string, error) {
//	fh, err := os.Open(fileName)
//
//	defer func() {
//		if err = fh.Close(); err != nil {
//			log.Printf("Error on closing file %s, %v\n", fileName, err)
//		}
//	}()
//
//	if err != nil {
//		log.Printf("Error when opening file %s, %v\n", fileName, err)
//		return "", err
//	}
//
//	scanner := bufio.NewScanner(fh)
//	var content string
//
//	for scanner.Scan() {
//		content += scanner.Text() + "\n"
//	}
//
//	if err := scanner.Err(); err != nil {
//		log.Printf("Error on closing file %s %v\n", fileName, err)
//		return "", err
//	}
//
//	return content, nil
//}

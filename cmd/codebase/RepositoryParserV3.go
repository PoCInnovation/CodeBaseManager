package codebase

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
//func (f FileParser) String() string {
//	str := "File name: " + f.Name + ", " + "File path: " + f.Path
//	return str
//}

//type Module struct {
//	Path, Name string
//	Files      []*FileParser
//}
//
//func (m *Module) Append(path, name string) *FileParser {
//	newFile := &FileParser{
//		Path: path,
//		Name: name,
//	}
//	m.Files = append(m.Files, newFile)
//	return newFile
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
//func (r Repository) isTarget(path string) bool {
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
//func (r *Repository) Append(path, name string) *Module {
//	newModule := &Module{
//		Path:  path,
//		Name:  name,
//		Files: nil,
//	}
//	r.Modules = append(r.Modules, newModule)
//	return newModule
//}
//
//func (r Repository) String() string {
//	var str string
//	for idx, module := range r.Modules {
//		str += fmt.Sprintf("%d: %s\n", idx+1, module.String())
//	}
//	return str
//}

//func ParseRepositoryV2(path string) (*Repository, error) {
//	path = strings.TrimSuffix(path, "/")
//	if !isDir(path) {
//		return nil, errors.New(path + " is not a valid target")
//	}
//	module := &Module{
//		Path: path,
//		Name: filepath.Base(path),
//	}
//	parser := &Repository{
//		Modules: []*Module{module},
//	}
//	processModules(parser, module, path)
//
//	return parser, nil
//}
//
//func isDir(path string) bool {
//	fileInfo, err := os.Stat(path)
//	if err != nil {
//		log.Printf("Error when stating %s, %v\n", path, err)
//		return false
//	}
//	if !fileInfo.IsDir() {
//		return false
//	}
//	return true
//}
//
//func listTargets(path string) ([]os.FileInfo, func()) {
//	dir, err := os.Open(path)
//	if err != nil {
//		log.Printf("Error when opening %s, %v\n", path, err)
//	}
//	fn := func() {
//		if err := dir.Close(); err != nil {
//			log.Printf("Cannot close :%s, %v\n", path, err)
//		}
//	}
//
//	targetList, err := dir.Readdir(0)
//	if err != nil {
//		log.Printf("Error when Readdir %s, %v\n", path, err)
//		return nil, nil
//	}
//	return targetList, fn
//}
//
//func isNotReadable(name string) bool {
//	info, err := os.Stat(name)
//
//	if err != nil {
//		return true
//	}
//
//	perm := info.Mode()
//	return info.Size() > 10000 && perm&0111 == 0111
//	//return info.Size() < 10000 && perm&0111 != 0111
//}

//func ProcessModules(moduleList *Repository, module *Module, path string) {
//	targetList, closure := listTargets(path)
//	if targetList == nil {
//		return
//	}
//	defer closure()
//
//	var wg sync.WaitGroup
//
//	for _, target := range targetList {
//		//targetName := target.Name()
//		//newPath := path + "/" + targetName
//		if target.IsDir() {
//			targetName := target.Name()
//			newPath := path + "/" + targetName
//			if strings.HasPrefix(targetName, ".") {
//				continue
//			}
//			newModule := moduleList.Append(newPath, targetName)
//			wg.Add(1)
//			go func() {
//				ProcessModules(moduleList, newModule, newPath)
//				log.Println("module", targetName, "processed")
//				wg.Done()
//			}()
//		} else {
//			targetName := target.Name()
//			newPath := path + "/" + targetName
//			if isNotReadable(newPath) {
//				continue
//			}
//			newFile := module.Append(newPath, targetName)
//			wg.Add(1)
//			go func() {
//				parseFile(newFile, newPath)
//				wg.Done()
//			}()
//		}
//	}
//
//	wg.Wait()
//}
//
//func parseFile(file *FileParser, path string) {
//	//log.Println("file", path, "processed")
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

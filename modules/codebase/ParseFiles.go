package codebase

func parseFile(file *FileParser, path string) {
	//log.Println("file", path, "processed")
	//datas.Append(path)

	_, err := GetFile(path)
	if err != nil {
		return
	}

	//content, err := getFile(path)
	//if err != nil {
	//	return
	//}
	//log.Println(content)
}

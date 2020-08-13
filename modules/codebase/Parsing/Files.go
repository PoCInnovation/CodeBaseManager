package codebase

import (
	"github.com/PoCFrance/CodeBaseManager/modules/codebase"
	Components "github.com/PoCFrance/CodeBaseManager/modules/codebase/Components"
)

func parseFile(file *Components.FileParser, path string) {
	//log.Println("file", path, "processed")
	//datas.Append(path)

	_, err := codebase.GetFile(path)
	if err != nil {
		return
	}

	//content, err := getFile(path)
	//if err != nil {
	//	return
	//}
	//log.Println(content)
}

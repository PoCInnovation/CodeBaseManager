package codebase

import (
	"path/filepath"
)

type Manager struct {
	Language   string
	Extensions []string
}

var CManager = Manager{
	Language:   "C",
	Extensions: []string{".c", ".h"},
	//languageManager []findFctArray
}

var GoManager = Manager{
	Language:   "Go",
	Extensions: []string{".go"},
	//languageManager []findFctArray
}

var ManagedLangage = map[string]Manager{
	"C":  CManager,
	"Go": GoManager,
}

func isSupportedLanguage(path string) bool {
	ext := filepath.Ext(path)

	for _, val := range ManagedLangage {
		for _, suf := range val.Extensions {
			if ext == suf {
				return true
			}
		}
	}
	return false
}

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

package codebase

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

const FILEPATH_REPOSITORY = "./.cbm/template/repository.toml"

type RepoTemplate struct {
	Language []string `toml:"Language"`
	Modules []string `toml:Modules`
	Tests   []string `toml:Tests`
}

func GetTomlRepo() RepoTemplate {
	var template RepoTemplate

	_, err := toml.DecodeFile(FILEPATH_REPOSITORY, &template)
	if err != nil {
		fmt.Println(err.Error())
	}
	return template
}
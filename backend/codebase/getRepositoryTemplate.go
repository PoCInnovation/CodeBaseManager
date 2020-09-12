package codebase

import (
	"github.com/BurntSushi/toml"
)

const FILEPATH_REPOSITORY = "./.cbm/template/test.toml"

type RepoTemplate struct {
	Language []string `toml:"Language"`
	Modules  []string `toml:Modules`
	Tests    []string `toml:Tests`
}

func GetTomlRepo(filepath string) (*RepoTemplate, error) {
	var template RepoTemplate

	_, err := toml.DecodeFile(filepath, &template)
	if err != nil {
		return nil, err
	}
	return &template, nil
}

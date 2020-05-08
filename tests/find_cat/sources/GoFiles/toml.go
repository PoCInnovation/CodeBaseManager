package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type RepoTemplate struct {
	Language string `toml:"Language"`
	Sources  struct {
		Modules []string
		Tests   []string
	} `toml:"Sources"`
}

func printTomlInfo(filepath string) {
	var template RepoTemplate

	_, err := toml.DecodeFile(filepath, &template)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(template)
	}
}

func main() {
	for _, filepath := range os.Args[1:] {
		printTomlInfo(filepath)
		fmt.Println("Viper\n")
		printViperInfo(filepath)
	}
}

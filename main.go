package main

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/cmd/codebase"
	"os"
)

func main() {
	//cmd.Execute()
	res, err := codebase.ParseRepositoryv2("/work/Projects/Personnal/CodeBaseManager/cmd")
	if err != nil {
		os.Exit(84)
	}
	fmt.Println(res)
}

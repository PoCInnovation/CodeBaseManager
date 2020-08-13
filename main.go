package main

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/modules/codebase/Parsing"
	"log"
	"os"
)

func main() {
	//cmd.Execute()
	res, err := codebase.ParseRepositoryV2("/work/Projects/Personnal/CodeBaseManager/")
	if err != nil {
		log.Println(err)
		os.Exit(84)
	}
	fmt.Println(res)
}

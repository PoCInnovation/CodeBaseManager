package codebase

import "fmt"

func PrintResult(args []string, parser parsingRepo) {
	for _, arg := range args {
		fmt.Printf("ARG: %s\n", arg)
		if contentFound, ok := parser.content[arg]; ok {
			for key, content := range contentFound {
				fmt.Printf("FILE: %s\n", key)
				fmt.Println(content)
			}
		}
	}
}

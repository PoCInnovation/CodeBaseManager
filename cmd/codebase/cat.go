package codebase

import (
	"fmt"
	"github.com/spf13/cobra"
)

func registerCat(parentCmd *cobra.Command) {
	var catCmd = &cobra.Command{
		Use:   "cat elem...",
		Short: "Prints the requested elements of the codebase.",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("Printing: ", args)
			cat(args)
		},
	}

	catCmd.Args = cobra.MinimumNArgs(1)

	parentCmd.AddCommand(catCmd)
}

func cat(args []string) {
	// TODO: Change repo parsing and evaluate repo language
	repo := []string{"."}
	parser := parsingRepo{
		args:    args,
		content: contentFound{},
		parser:  CatParser,
	}
	for _, module := range repo {
		RepoParser(module, parser)
	}
	PrintResult(args, parser)
}

//func printResult(args []string, parser parsingRepo) {
//	for _, arg := range args {
//		fmt.Printf("ARG: %s\n", arg)
//		if contentFound, ok := parser.content[arg]; ok {
//			for key, content := range contentFound {
//				fmt.Printf("FILE: %s\n", key)
//				fmt.Println(content)
//			}
//		}
//	}
//}

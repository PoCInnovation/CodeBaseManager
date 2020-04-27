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
	// TODO: Change repo parsing
	repo := []string{"."}
	parser := parsingRepo{
		//found:   make([]bool, len(args)),
		args:    args,
		content: contentFound{},
		parser:  CatParser,
	}
	for _, module := range repo {
		RepoParser(module, parser)
	}
	for _, arg := range args {
		if contentFound, ok := parser.content[arg]; ok {
			for _, content := range contentFound {
				fmt.Println(content)
			}
		}
	}
}

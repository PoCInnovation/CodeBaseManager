package codebase

import (
	"fmt"
	"github.com/spf13/cobra"
)

func registerFind(parentCmd *cobra.Command) {
	var findCmd = &cobra.Command{
		Use:   "find elem...",
		Short: "Tells you where the requested elements of the codebase are located.",
		Run: func(_ *cobra.Command, args []string) {
			// TODO: Add its real behavior.
			fmt.Println("Looking for: ", args)
			find(args)
		},
	}

	findCmd.Args = cobra.MinimumNArgs(1)

	parentCmd.AddCommand(findCmd)
}

func find(args []string) {
	// TODO: Change repo parsing
	repo := []string{"."}
	parser := parsingRepo{
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

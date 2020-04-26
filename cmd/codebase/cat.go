package codebase

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Blalal(args []string) {
	repo := []string{"main.go"}
	found := make([]bool, len(args))
	parser := parsingRepo{nil, contentFound{}, args, found}
	ParseRepo(&parser, repo)
}

func registerCat(parentCmd *cobra.Command) {
	var catCmd = &cobra.Command{
		Use:   "cat elem...",
		Short: "Prints the requested elements of the codebase.",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("Printing: ", args)
			Blalal(args)
		},
	}

	catCmd.Args = cobra.MinimumNArgs(1)

	parentCmd.AddCommand(catCmd)
}

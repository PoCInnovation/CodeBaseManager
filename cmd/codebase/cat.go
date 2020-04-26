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
	// TODO
	CatParseRepo(parsingRepo{
		found: make([]bool, len(args)),
		repo:  []string{"cmd"},
	})
}

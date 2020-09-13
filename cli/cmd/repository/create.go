package repository

import (
	"github.com/PoCFrance/CodeBaseManager/cli/modules/repository"
	"github.com/spf13/cobra"
)

func registerCreate(parentCmd *cobra.Command) {
	var createCmd = &cobra.Command{
		Use:   "create <github link>",
		Short: "Create your project repository based on the given template.",
		Run: func(_ *cobra.Command, args []string) {
			repository.CreateRepository(args)
		},
	}

	createCmd.Args = cobra.RangeArgs(1, 2)

	parentCmd.AddCommand(createCmd)
}

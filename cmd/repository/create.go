package repository

import (
	"github.com/PoCFrance/CodeBaseManager/modules/repository"
	"github.com/spf13/cobra"
)

func registerCreate(parentCmd *cobra.Command) {
	var createCmd = &cobra.Command{
		Use:   "create link.to.template",
		Short: "Create your project repository based on the given template.",
		Run: func(_ *cobra.Command, args []string) {
			repository.CreateRepository(args[0])
		},
	}

	createCmd.Args = cobra.ExactArgs(1)

	parentCmd.AddCommand(createCmd)
}

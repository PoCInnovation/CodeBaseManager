package repository

import (
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var repoCmd = &cobra.Command{
		Use:                        "repository",
		Short:                      "Helps you manage your repository",
	}

	repoCmd.Args = cobra.ExactArgs(0)

	registerCreate(repoCmd)
	parentCmd.AddCommand(repoCmd)
}

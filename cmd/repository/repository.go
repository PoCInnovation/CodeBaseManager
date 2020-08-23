package repository

import (
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var repoCmd = &cobra.Command{
		Use:   "repository",
		Short: "Helps you manage your repository",
	}

	registerCreate(repoCmd)
	registerTodos(repoCmd)
	registerWatch(repoCmd)
	parentCmd.AddCommand(repoCmd)
}

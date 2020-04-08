package repository

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var repoCmd = &cobra.Command{
		Use:                        "repository",
		Short:                      "Helps you handle your repository",
		Long:                       "",
		Example:                    "",
		Run:                        func(cmd *cobra.Command, args []string) {
			fmt.Println("repository called")
		},
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	parentCmd.AddCommand(repoCmd)
}

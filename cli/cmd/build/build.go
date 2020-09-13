package build

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Helps you build your project regardless of language.",
		Run: func(cmd *cobra.Command, args []string) {
			//TODO: Add its real behavior
			fmt.Println("Build called")
		},
	}

	parentCmd.AddCommand(buildCmd)
}

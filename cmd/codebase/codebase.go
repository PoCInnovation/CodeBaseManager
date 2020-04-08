package codebase

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var codebaseCmd = &cobra.Command{
		Use:   "codebase",
		Short: "Helps you naviguate through your codebase.",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("codebase called")
		},
	}

	initCat()
	codebaseCmd.AddCommand(catCmd)
	initFind()
	codebaseCmd.AddCommand(findCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	parentCmd.AddCommand(codebaseCmd)
}

package unitTests

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var utCmd = &cobra.Command{
		Use:                        "unit-tests",
		Short:                      "Helps you handle unit tests of your projects",
		Long:                       "",
		Example:                    "",
		Run:                        func(cmd *cobra.Command, args []string) {
			fmt.Println("unit tests called")
		},
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	initCreate()

	utCmd.AddCommand(createCmd)
	parentCmd.AddCommand(utCmd)
}
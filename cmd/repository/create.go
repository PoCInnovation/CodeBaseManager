package repository

import (
	"fmt"
	"github.com/spf13/cobra"
)


func registerCreate(parentCmd *cobra.Command) {
	var createCmd = &cobra.Command{
		Use:   "create link.to.template",
		Short: "Create your project repository based on the given template.",
		Run: func(_ *cobra.Command, args []string) {
			//TODO: Add its real behavior
			fmt.Println("Creating repo based on:", args[0])
		},
	}

	createCmd.Args = cobra.ExactArgs(1)

	parentCmd.AddCommand(createCmd)
}

package codebase

import (
	"github.com/PoCFrance/CodeBaseManager/REPLs"
	"github.com/PoCFrance/CodeBaseManager/cmd/utils"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var codebaseCmd = &cobra.Command{
		Use:   "codebase",
		Short: "Simple shell to navigate through your codebase.",
		Args:  utils.IsCBMRepository,
		Run: func(_ *cobra.Command, _ []string) {
			REPLs.CodebaseShell()
		},
	}

	codebaseCmd.Args = cobra.ExactArgs(0)

	registerFind(codebaseCmd)
	registerCat(codebaseCmd)
	parentCmd.AddCommand(codebaseCmd)
}


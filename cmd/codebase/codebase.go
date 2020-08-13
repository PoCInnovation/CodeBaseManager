package codebase

import (
	"github.com/PoCFrance/CodeBaseManager/REPL"
	"github.com/PoCFrance/CodeBaseManager/cmd/common"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var codebaseCmd = &cobra.Command{
		Use:   "codebase",
		Short: "Simple shell to navigate through your codebase.",
		Args:  common.IsCBMRepository,
		Run: func(cmd *cobra.Command, _ []string) {
			sh := REPL.NewShell("CodeBase")
			accepted := REPL.Builtins{
				"cat":  Cat,
				"find": Find,
			}
			sh.Run(accepted)
		},
	}

	codebaseCmd.Args = cobra.ExactArgs(0)

	registerFind(codebaseCmd)
	registerCat(codebaseCmd)
	parentCmd.AddCommand(codebaseCmd)
}

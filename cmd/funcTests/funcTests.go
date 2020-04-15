package funcTests

import (
	"github.com/PoCFrance/CodeBaseManager/REPL"
	"github.com/PoCFrance/CodeBaseManager/cmd/common"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var ftCmd = &cobra.Command{
		Use:     "functional-tests",
		Short:   "Helps you deal with your functional tests.",
		Aliases: []string{"ft"},
		Args:    common.IsCBMRepository,
		Run: func(cmd *cobra.Command, _ []string) {
			sh := REPL.NewShell("Functional Tests")
			acceptedBuiltins := common.RetrieveSubCommandsNames(cmd)

			sh.Run(acceptedBuiltins)
		},
	}

	ftCmd.Args = cobra.ExactArgs(0)

	parentCmd.AddCommand(ftCmd)
}

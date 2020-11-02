package funcTests

import (
	"github.com/PoCInnovation/CodeBaseManager/cli/REPL"
	"github.com/PoCInnovation/CodeBaseManager/cli/cmd/common"
	"github.com/PoCInnovation/CodeBaseManager/cli/modules/funcTests"
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
			accepted := REPL.Builtins{
				"run": funcTests.Run,
			}

			sh.Run(accepted)
		},
	}

	ftCmd.Args = cobra.ExactArgs(0)
	registerRun(ftCmd)
	parentCmd.AddCommand(ftCmd)
}

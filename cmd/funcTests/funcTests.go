package funcTests

import (
	"github.com/PoCFrance/CodeBaseManager/REPLs"
	"github.com/PoCFrance/CodeBaseManager/cmd/utils"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var ftCmd = &cobra.Command{
		Use:     "functional-tests",
		Short:   "Helps you deal with your functional tests.",
		Aliases: []string{"ft"},
		Args:    utils.IsCBMRepository,
		Run:                        func(_ *cobra.Command, _ []string) {
			REPLs.FunctionalTestsShell()
		},
	}

	ftCmd.Args = cobra.ExactArgs(0)

	parentCmd.AddCommand(ftCmd)
}

package debug

import (
	"github.com/PoCFrance/CodeBaseManager/REPL"
	"github.com/PoCFrance/CodeBaseManager/cmd/common"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var debugCmd = &cobra.Command{
		Use:   "debug",
		Short: "Runs the debug prompt",
		Run: func(cmd *cobra.Command, _ []string) {
			sh := REPL.NewShell("Debug")
			acceptedBuiltins := common.RetrieveSubCommandsNames(cmd)

			sh.Run(acceptedBuiltins)
		},
	}

	debugCmd.Args = cobra.ExactArgs(0)

	parentCmd.AddCommand(debugCmd)
}

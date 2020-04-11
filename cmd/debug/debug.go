package debug

import (
	"github.com/PoCFrance/CodeBaseManager/REPLs"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var debugCmd = &cobra.Command{
		Use:                        "debug",
		Short:                      "Runs the debug prompt",
		Run:                        func(cmd *cobra.Command, args []string) {
			REPLs.DebugShell()
		},
	}

	debugCmd.Args = cobra.ExactArgs(0)

	parentCmd.AddCommand(debugCmd)
}
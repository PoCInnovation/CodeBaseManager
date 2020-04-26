package debug

import (
	"github.com/PoCFrance/CodeBaseManager/REPL"
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var debugCmd = &cobra.Command{
		Use:   "debug",
		Short: "Runs the debug prompt",
		Run: func(cmd *cobra.Command, _ []string) {
			sh := REPL.NewShell("Debug")
			accepted := REPL.Builtins{}

			sh.Run(accepted)
		},
	}

	debugCmd.Args = cobra.ExactArgs(0)

	parentCmd.AddCommand(debugCmd)
}

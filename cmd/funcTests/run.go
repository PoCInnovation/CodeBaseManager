package funcTests

import (
	"github.com/PoCFrance/CodeBaseManager/modules/funcTests"
	"github.com/PoCFrance/CodeBaseManager/modules/logs"
	"github.com/spf13/cobra"
)

func registerRun(parentCmd *cobra.Command) {
	var catCmd = &cobra.Command{
		Use:   "run [tests...]",
		Short: "Either run all tests or only the specified ones.",
		Run: func(_ *cobra.Command, args []string) {
			logs.InitCBMLogs(logs.Verbosity, logs.LogsFP)
			funcTests.Run(args)
			logs.CBMLogs.Close()
		},
	}

	parentCmd.AddCommand(catCmd)
}

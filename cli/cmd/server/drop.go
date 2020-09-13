package server

import (
	"github.com/PoCFrance/CodeBaseManager/cli/modules/server"
	"github.com/spf13/cobra"
)

func registerDrop(parent *cobra.Command) {
	var cmd = &cobra.Command{
		Use:   "drop repo",
		Short: "Tells CBM to stop watching the given repository",
		Run: func(_ *cobra.Command, args []string) {
			server.Drop(args[0])
		},
	}

	cmd.Args = cobra.ExactArgs(1)
	parent.AddCommand(cmd)
}

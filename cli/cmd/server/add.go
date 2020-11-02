package server

import (
	"github.com/PoCInnovation/CodeBaseManager/cli/modules/server"
	"github.com/spf13/cobra"
)

func registerAdd(parent *cobra.Command) {
	var cmd = &cobra.Command{
		Use:   "add path/to/repo",
		Short: "Tells CBM to watch the given repository",
		Run: func(_ *cobra.Command, args []string) {
			server.Add(args[0])
		},
	}

	cmd.Args = cobra.ExactArgs(1)
	parent.AddCommand(cmd)
}

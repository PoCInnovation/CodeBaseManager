package server

import (
	"github.com/PoCInnovation/CodeBaseManager/cli/modules/server"
	"github.com/spf13/cobra"
)

func registerStart(parent *cobra.Command) {
	var cmd = &cobra.Command{
		Use:   "start",
		Short: "Starts CBM's server",
		Run: func(_ *cobra.Command, args []string) {
			server.Start()
		},
	}

	cmd.Args = cobra.ExactArgs(0)
	parent.AddCommand(cmd)
}

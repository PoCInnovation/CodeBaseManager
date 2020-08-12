package server

import "github.com/spf13/cobra"

func RegisterCmd(parentCmd *cobra.Command) {
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Allows you to interact with CBM's server",
	}

	serverCmd.Args = cobra.ExactArgs(0)

	registerStart(serverCmd)
	registerStop(serverCmd)
	registerAdd(serverCmd)
	registerDrop(serverCmd)
	registerList(serverCmd)

	parentCmd.AddCommand(serverCmd)
}

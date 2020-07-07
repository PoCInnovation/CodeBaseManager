package server

import (
    "github.com/PoCFrance/CodeBaseManager/modules/server"
    "github.com/spf13/cobra"
)

func registerStop(parent *cobra.Command) {
    var cmd = &cobra.Command{
        Use: "stop",
        Short: "Stops CBM's server",
        Run: func(_ *cobra.Command, args[]string){
            server.Stop()
        },
    }

    cmd.Args = cobra.ExactArgs(0)
    parent.AddCommand(cmd)
}


package server

import (
    "github.com/PoCFrance/CodeBaseManager/modules/server"
    "github.com/spf13/cobra"
)

func registerList(parent *cobra.Command) {
    var cmd = &cobra.Command{
        Use: "list",
        Short: "Displays all repository watched by CBM",
        Run: func(_ *cobra.Command, args[]string){
            server.List()
        },
    }

    cmd.Args = cobra.ExactArgs(0)
    parent.AddCommand(cmd)

}

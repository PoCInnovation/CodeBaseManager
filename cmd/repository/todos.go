package repository

import (
    "github.com/PoCFrance/CodeBaseManager/modules/repository"
    "github.com/spf13/cobra"
)

func registerTodos(parentCmd *cobra.Command) {
    var todosCmd = &cobra.Command{
        Use: "todos",
        Short: "Displays all TODO in your repository",
        Run: func(_ *cobra.Command, _ []string) {
            repository.DisplayTodos()
        },
    }

    todosCmd.Args = cobra.ExactArgs(0)
    parentCmd.AddCommand(todosCmd)
}

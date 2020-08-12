package repository

import (
	"github.com/PoCFrance/CodeBaseManager/modules/repository"
	"github.com/spf13/cobra"
)

func registerTodos(parentCmd *cobra.Command) {
	var ctxDepth int
	var todosCmd = &cobra.Command{
		Use:   "todos",
		Short: "Displays all TODO in your repository",
		Run: func(_ *cobra.Command, _ []string) {
			repository.DisplayTodos(ctxDepth)
		},
	}

	todosCmd.Args = cobra.ExactArgs(0)
	todosCmd.Flags().IntVarP(&ctxDepth, "context-depth", "c", 0,
		"Number of lines above & below todo to be shown")
	parentCmd.AddCommand(todosCmd)
}

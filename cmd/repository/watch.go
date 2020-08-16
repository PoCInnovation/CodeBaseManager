package repository

import (
	"github.com/PoCFrance/CodeBaseManager/modules/repository"
	"github.com/spf13/cobra"
)

func registerWatch(parentCmd *cobra.Command) {
	var watchCmd = &cobra.Command{
		Use:   "watch",
		Short: "Tell CodeBaseManager to silently watch this repository.",
		Run: func(_ *cobra.Command, args []string) {
			repository.WatchRepository()
		},
	}

	parentCmd.AddCommand(watchCmd)
}

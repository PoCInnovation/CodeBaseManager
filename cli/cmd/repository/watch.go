package repository

import (
	"github.com/PoCInnovation/CodeBaseManager/cli/modules/repository"
	"github.com/spf13/cobra"
)

func registerWatch(parentCmd *cobra.Command) {
	var watchCmd = &cobra.Command{
		Use:   "watch (path path ...)",
		Short: "Watch repositories in parameter or WD if no parameter.",
		Run: func(_ *cobra.Command, args []string) {
			repository.WatchRepository(args)
		},
	}

	watchCmd.Args = cobra.RangeArgs(0, 1)

	parentCmd.AddCommand(watchCmd)
}

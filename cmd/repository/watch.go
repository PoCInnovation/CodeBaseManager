package repository

import (
	"github.com/PoCFrance/CodeBaseManager/modules/repository"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func registerWatch(parentCmd *cobra.Command) {
	actualPath, err := os.Getwd()
	if err != nil {
		log.Println("HERE", err)
		return
	}
	log.Println(actualPath)
	var watchCmd = &cobra.Command{
		Use:   "watch",
		Short: "Tell CodeBaseManager to silently watch this repository.",
		Run: func(_ *cobra.Command, args []string) {
			repository.WatchRepository(actualPath)
		},
	}

	//createCmd.Args = cobra.ExactArgs(1)

	parentCmd.AddCommand(watchCmd)
}

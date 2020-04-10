package unitTests

import (
	"github.com/spf13/cobra"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var utCmd = &cobra.Command{
		Use:                        "unit-tests",
		Short:                      "Helps you deal with your unit tests.",
	}

	registerCreate(utCmd)
	parentCmd.AddCommand(utCmd)
}
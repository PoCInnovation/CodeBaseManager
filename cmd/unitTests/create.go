package unitTests

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	utType string
)

func registerCreate(parentCmd *cobra.Command) {
	var createCmd = &cobra.Command{
		Use:   "create x func",
		Short: "Creates a given number of unit-tests for a given function.",
		Args:  checkArgs,
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("unit-tests create:", args)
		},
	}

	createCmd.Flags().StringVarP(&utType, "type", "t", "basic", "Type of unit-tests")

	parentCmd.AddCommand(createCmd)
}

func checkArgs(cmd *cobra.Command, av []string) error {
	if len(av) != 2 {
		return errors.New("Exactly 2 arguments required.")
	}
	if _, err := strconv.Atoi(av[0]); err != nil {
		return errors.New("The first argument must be an integer.")
	}
	//TODO: check if av[1] is a known function
	//TODO: check if type has legit value ("basic", "display", ... based on ut.toml")
	//      cmd -> flags-> lookup -> value
	return nil
}

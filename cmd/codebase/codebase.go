package codebase

import (
	"errors"
	"github.com/PoCFrance/CodeBaseManager/REPLs"
	"github.com/spf13/cobra"
	"os"
)

func RegisterCmd(parentCmd *cobra.Command) {
	var codebaseCmd = &cobra.Command{
		Use:   "codebase",
		Short: "Simple shell to navigate through your codebase.",
		Args:  isCBMRepository,
		Run: func(_ *cobra.Command, _ []string) {
			REPLs.CodebaseShell()
		},
	}

	codebaseCmd.Args = cobra.ExactArgs(0)

	registerFind(codebaseCmd)
	registerCat(codebaseCmd)
	parentCmd.AddCommand(codebaseCmd)
}

func isCBMRepository(_ *cobra.Command, av []string) error {
	const (
		currentDir = 0
		goToDir = 1
	)

	switch len(av) {
	case currentDir:
		st, err := os.Stat(".cbm/")
		if err != nil || !st.IsDir() {
			return errors.New("Not in a CBM Repository.")
		}
	case goToDir:
		st, err := os.Stat(av[0])
		if err != nil || !st.IsDir() {
			return errors.New("Invalid filepath")
		}
		st, err = os.Stat(av[0] + "/.cbm/")
		if err != nil || !st.IsDir() {
			return errors.New("Not a CBM Repository")
		}
		if err = os.Chdir(av[0]); err != nil {
			return errors.Unwrap(err)
		}
	}
	return nil
}

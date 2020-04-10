package cmd

import (
	"errors"
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/REPLs"
	"github.com/PoCFrance/CodeBaseManager/cmd/build"
	"github.com/PoCFrance/CodeBaseManager/cmd/codebase"
	"github.com/PoCFrance/CodeBaseManager/cmd/debug"
	"github.com/PoCFrance/CodeBaseManager/cmd/funcTests"
	"github.com/PoCFrance/CodeBaseManager/cmd/repository"
	"github.com/PoCFrance/CodeBaseManager/cmd/unitTests"
	"github.com/spf13/cobra"
	"os"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "CodeBaseManager [Path/To/Repository]",
		Short: "Multi-langage CLI tool to manage your code base.",
		Args:  isCBMRepository,
		Run: func(_ *cobra.Command, _ []string) {
			REPLs.CBMShell()
		},
	}

	registerSubCmds(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

func registerSubCmds(rootCmd *cobra.Command) {
	codebase.RegisterCmd(rootCmd)
	repository.RegisterCmd(rootCmd)
	unitTests.RegisterCmd(rootCmd)
	funcTests.RegisterCmd(rootCmd)
	build.RegisterCmd(rootCmd)
	debug.RegisterCmd(rootCmd)
}
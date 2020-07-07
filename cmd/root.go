package cmd

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/cmd/codebase"
	"github.com/PoCFrance/CodeBaseManager/cmd/funcTests"
	"github.com/PoCFrance/CodeBaseManager/cmd/repository"
	"github.com/PoCFrance/CodeBaseManager/cmd/server"
	"github.com/spf13/cobra"
	"os"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "CodeBaseManager",
		Short: "Multi-langage CLI tool to manage your code base.",
	}

	registerSubCmds(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func registerSubCmds(rootCmd *cobra.Command) {
	codebase.RegisterCmd(rootCmd)
	repository.RegisterCmd(rootCmd)
	funcTests.RegisterCmd(rootCmd)
	server.RegisterCmd(rootCmd)
}

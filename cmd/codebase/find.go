package codebase

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

func registerFind(parentCmd *cobra.Command) {
	var findCmd = &cobra.Command{
		Use:   "find elem...",
		Short: "Tells you where the requested elements of the codebase are located.",
		Run: func(_ *cobra.Command, args []string) {
			// TODO: Add its real behavior.
			fmt.Println("Looking for: ", args)
			find(args)
		},
	}

	findCmd.Args = cobra.MinimumNArgs(1)

	parentCmd.AddCommand(findCmd)
}

func find(args []string) {
	// TODO: Change repo parsing and evaluate repo language
	repo := []string{"."}
	parser := parsingRepo{
		args:       args,
		content:    contentFound{},
		parser:     FindParser,
		manageFile: FindFile,
	}
	for _, module := range repo {
		RepoParser(module, parser)
	}
	PrintResult(args, parser)
}

func FindParser(name string, control parsingRepo) {
	for _, arg := range control.args {
		splitName := strings.Split(name, "/")
		splitLen := len(splitName)
		if splitLen == 0 {
			log.Printf("Cannot Split %s\n", name)
		}

		if arg == splitName[splitLen-1] {
			// TODO: refacto parsing to use fctPtr -> common ground for cat and find
			control.content[arg], _ = control.manageFile(control.content[arg], name)
		} else {
			//read content to find function
		}
	}
}

func FindFile(controlContent map[string]string, name string) (map[string]string, error) {
	if controlContent != nil {
		controlContent[name] = name
	} else {
		controlContent = map[string]string{}
		controlContent[name] = name
	}
	return controlContent, nil
}

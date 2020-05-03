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
	repo := []string{"cmd", "modules", "REPL", "test_viper", "tests"}
	// TODO: Manage Panic when reading binary (regexp)
	//repo := []string{"."}
	parser := parsingRepo{
		args:    args,
		content: contentFound{},
		//parser:          FindParser,
		fileManager:     FindFile,
		functionManager: FindFunction,
	}
	for _, module := range repo {
		RepoParser(module, parser)
	}
	PrintResult(args, parser)
}

// TODO: Delete function => common ground for cat and find (argParser in FindInRepository)
func FindParser(name string, control parsingRepo) {
	for _, arg := range control.args {
		splitName := strings.Split(name, "/")
		splitLen := len(splitName)
		if splitLen == 0 {
			log.Printf("Cannot Split %s\n", name)
		}

		if arg == splitName[splitLen-1] {
			// TODO: refacto parsing to use fctPtr -> common ground for cat and find
			control.content[arg], _ = control.fileManager(control.content[arg], name)
		} else {
			// TODO: refacto parsing to use fctPtr -> common ground for cat and find
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

func FindFunction(controlContent map[string]string, name, arg string) (map[string]string, error) {
	fmt.Println(name)
	return controlContent, nil
}

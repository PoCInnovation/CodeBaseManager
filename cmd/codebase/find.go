package codebase

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/modules/codebase"
	"github.com/spf13/cobra"
)

func registerFind(parentCmd *cobra.Command) {
	var findCmd = &cobra.Command{
		Use:   "find elem...",
		Short: "Tells you where the requested elements of the codebase are located.",
		Run: func(_ *cobra.Command, args []string) {
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
		args:            args,
		content:         contentFound{},
		fileManager:     FindFile,
		functionManager: FindFunction,
	}
	for _, module := range repo {
		RepoParser(module, parser)
	}
	PrintResult(args, parser)
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
	fmt.Println(name + "\tFIND")
	content, err := codebase.GetFile(name)
	if err != nil {
		return controlContent, err
	}

	// TODO: Manage several language ? array of function pointer given repository language
	if found := findGoFunction(*content, arg); found != nil {
		if controlContent != nil {
			controlContent[name] = *found
		} else {
			controlContent = map[string]string{}
			controlContent[name] = *found
		}
	}
	//if found := findCFunction(*content, arg); found != nil {
	//	if controlContent != nil {
	//		controlContent[name] = *found
	//	} else {
	//		controlContent = map[string]string{}
	//		controlContent[name] = *found
	//	}
	//}
	return controlContent, nil
}

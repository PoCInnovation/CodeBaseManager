package codebase

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/modules/codebase"
	"github.com/spf13/cobra"
)

func registerCat(parentCmd *cobra.Command) {
	var catCmd = &cobra.Command{
		Use:   "cat elem...",
		Short: "Prints the requested elements of the codebase.",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("Printing: ", args)
			cat(args)
		},
	}

	catCmd.Args = cobra.MinimumNArgs(1)

	parentCmd.AddCommand(catCmd)
}

func cat(args []string) {
	// TODO: Change repo parsing and evaluate repo language
	repo := []string{"cmd", "modules", "REPL", "test_viper", "tests"}
	// TODO: Manage Panic when reading binary (regexp)
	//repo := []string{"."}
	parser := parsingRepo{
		args:            args,
		content:         contentFound{},
		fileManager:     CatFile,
		functionManager: CatFunction,
	}
	for _, module := range repo {
		RepoParser(module, parser)
	}
	PrintResult(args, parser)
}

func CatFile(controlContent map[string]string, name string) (map[string]string, error) {
	content, err := codebase.GetFile(name)
	if err != nil {
		return controlContent, err
	}

	if controlContent != nil {
		controlContent[name] = *content
	} else {
		controlContent = map[string]string{}
		controlContent[name] = *content
	}

	return controlContent, nil
}

func CatFunction(controlContent map[string]string, name, arg string) (map[string]string, error) {
	content, err := codebase.GetFile(name)
	if err != nil {
		return controlContent, err
	}

	// TODO: Manage several language ? array of function pointer given repository language
	if found := catGoFunction(*content, arg); found != nil {
		if controlContent != nil {
			controlContent[name] = *found
		} else {
			controlContent = map[string]string{}
			controlContent[name] = *found
		}
	}
	//if found := catCFunction(*content, arg); found != nil {
	//	if controlContent != nil {
	//		controlContent[name] = *found
	//	} else {
	//		controlContent = map[string]string{}
	//		controlContent[name] = *found
	//	}
	//}
	return controlContent, nil
}

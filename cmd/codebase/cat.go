package codebase

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/modules/codebase"
	"github.com/spf13/cobra"
	"log"
	"strings"
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
	repo := []string{"."}

	supportedLanguage := setupTargetFunctions(catTargetFcts)
	if len(supportedLanguage) == 0 {
		log.Println("No supported Language in user configuration.")
		return
	}

	parser := parsingRepo{
		args:            args,
		content:         contentFound{},
		fileManager:     catFile,
		functionManager: catFunction,
		languageManager: supportedLanguage,
	}
	for _, module := range repo {
		RepoParser(module, parser)
	}
	PrintResult(args, parser)
}

func catFile(controlContent map[string]string, name string) (map[string]string, error) {
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

func catFunction(controlContent map[string]string, name, arg string, supportedLanguages []findFctArray) (map[string]string, error) {
	for _, supportedLang := range supportedLanguages {
		for _, extension := range supportedLang.extensions {
			if strings.HasSuffix(name, extension) {
				fmt.Println(name, extension)
				content, err := codebase.GetFile(name)
				if err != nil {
					return controlContent, err
				}
				if found := supportedLang.fct(*content, arg); found != nil {
					if controlContent != nil {
						controlContent[name] = *found
					} else {
						controlContent = map[string]string{}
						controlContent[name] = *found
					}
					return controlContent, nil
				}
			}
		}
	}
	return controlContent, nil
}

//func catFunction(controlContent map[string]string, name, arg string, supportedLanguages []findFctArray) (map[string]string, error) {
//	content, err := codebase.GetFile(name)
//	if err != nil {
//		return controlContent, err
//	}
//
//	// TODO: Manage several language ? array of function pointer given repository language
//	//if found := catGoFunction(*content, arg); found != nil {
//	//	if controlContent != nil {
//	//		controlContent[name] = *found
//	//	} else {
//	//		controlContent = map[string]string{}
//	//		controlContent[name] = *found
//	//	}
//	//}
//	if found := catCFunction(*content, arg); found != nil {
//		if controlContent != nil {
//			controlContent[name] = *found
//		} else {
//			controlContent = map[string]string{}
//			controlContent[name] = *found
//		}
//	}
//	return controlContent, nil
//}

package codebase

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/modules/codebase"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

func registerFind(parentCmd *cobra.Command) {
	var findCmd = &cobra.Command{
		Use:   "find elem...",
		Short: "Tells you where the requested elements of the codebase are located.",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("Looking for: ", args)
			Find(args)
		},
	}

	findCmd.Args = cobra.MinimumNArgs(1)

	parentCmd.AddCommand(findCmd)
}

func Find(args []string) {
	// TODO: Change repo parsing and evaluate repo language
	// Repo allan
	repo := []string{"."}

	supportedLanguage, err := setupTargetFunctions(TargetFcts)
	if err != nil {
		log.Println(err)
		return
	}
	if supportedLanguage == nil {
		log.Println("No supported Language in user configuration.")
		return
	}

	parser := parsingRepo{
		args:            args,
		content:         contentFound{},
		fileManager:     findFile,
		functionManager: findFunction,
		languageManager: supportedLanguage,
	}
	for _, module := range repo {
		RepoParser(module, parser)
	}
	PrintResult(args, parser)
}

func findFile(controlContent map[string]string, name string) (map[string]string, error) {
	if controlContent != nil {
		controlContent[name] = name
	} else {
		controlContent = map[string]string{}
		controlContent[name] = name
	}
	return controlContent, nil
}

func findFunction(controlContent map[string]string, name, arg string, supportedLanguages []findFctArray) (map[string]string, error) {
	// TODO: move Supported languages management in args parser ?
	for _, supportedLang := range supportedLanguages {
		for _, extension := range supportedLang.extensions {
			if strings.HasSuffix(name, extension) {
				//fmt.Println(name, extension)
				content, err := codebase.GetFile(name)
				if err != nil {
					// TODO: continue ? (if supported
					return controlContent, err
				}
				if found := supportedLang.fct[FIND](*content, arg); found != nil {
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

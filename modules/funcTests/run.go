package funcTests

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
)

func errorPrompt(err error) bool {
	var answer bool
	prompt := &survey.Confirm{
		Message: err.Error() + "\nWould you like to continue ?",
	}
	survey.AskOne(prompt, &answer, survey.WithValidator(survey.Required))
	return answer
}

func Run(av []string) {
	failed := 0
	for _, fp := range av {
		failPerConf := 0
		cfg, err := NewConfigFT(fp)
		if err != nil {
			fmt.Println(fp, err)
			//errorPrompt(err)
			return
		}
		// TODO: if no bin ask build module for binary
		for _, test := range cfg.Tests {
			test.Init(&cfg.Common)
			test.Run()
			failPerConf += test.GetResults().Show(test.Name)
		}
		if failPerConf != 0 {
			fmt.Printf("Failed %d tests on [%s]\n", failPerConf, fp)
		}
		failed += failPerConf
	}
	fmt.Printf("Failed %d tests in total\n", failed)
	if failed != 0 {
		os.Exit(1)
	}
}

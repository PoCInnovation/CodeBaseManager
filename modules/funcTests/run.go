package funcTests

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
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
	for _, fp := range av {
		cfg, err := NewConfigFT(fp)
		if err != nil {
			fmt.Println(err)
			//errorPrompt(err)
			return
		}
		// TODO: if no bin ask build module for binary
		for _, test := range cfg.Tests {
			test.Init(&cfg.Common)
			test.Run()
			test.GetResults().Show(test.Name)
		}
	}
}

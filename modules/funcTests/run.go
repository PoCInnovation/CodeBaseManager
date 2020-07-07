package funcTests

import (
	"github.com/AlecAivazis/survey/v2"
	ft_types "github.com/PoCFrance/CodeBaseManager/modules/funcTests/types"
	"os"
)

func errorPrompt(err error) {
	var answer bool
	prompt := &survey.Confirm{
		Message: err.Error() + "\nWould you like to continue ?",
	}
	survey.AskOne(prompt, &answer, survey.WithValidator(survey.Required))
	if answer == false {
		os.Exit(1)
	}
}

func Run(av []string) {
	for _, fp := range av {
		cfg, err := ft_types.NewTestSuite(fp)
		if err != nil {
			errorPrompt(err)
			continue
		}
		cfg.Exec()
	}
}
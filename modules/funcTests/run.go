package funcTests

import (
	"github.com/PoCFrance/CodeBaseManager/modules/funcTests/types"
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
	if answer == false {
		os.Exit(1)
	}
	return answer
}

func Run(av []string) {
	for _, fp := range av {
		cfg, err := ft_types.NewTestSuite(fp)
		if err != nil {
			errorPrompt(err)
			continue
		}
		//TODO Remove this horrible printing thing
		fmt.Println("Vars:")
		fmt.Println(cfg.Vars, "\n")
		fmt.Println("Tests:")
		fmt.Println(cfg.Tests, "\n")
		fmt.Println("Default:")
		fmt.Println(cfg.Default.Name)
		fmt.Println(cfg.Default.Desc)
		fmt.Println(cfg.Default.Bin)
		fmt.Println(cfg.Default.RefBin)
		fmt.Println(cfg.Default.Args)
		fmt.Println(cfg.Default.RefArgs)
		fmt.Println(cfg.Default.Expected)
		fmt.Println(cfg.Default.Interactions)
		fmt.Println(cfg.Default.Options)
		// TODO: if no bin ask build module for binary
		//for _, test := range cfg.Tests {
		//	//TODO: now that config is load, launch fts
		//	/*test.Init(&cfg.Common)
		//	test.ref.execTime = time.Now()
		//	test.Run(test.Opt)
		//	test.GetResults().Show(test.Name)*/
		//}
	}
}
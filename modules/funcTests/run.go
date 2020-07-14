package funcTests

import (
	ft_types "github.com/PoCFrance/CodeBaseManager/modules/funcTests/types"
	"github.com/PoCFrance/CodeBaseManager/modules/logs"
)

func Run(av []string) {
	if av[0] == "run" {
		av = av[1:]
	}
	if len(av) == 0 {
		logs.CBMLogs.Error("No files given in input")
		return
	}
	for _, fp := range av {
		cfg, err := ft_types.NewTestSuite(fp)
		if err != nil {
			logs.CBMLogs.Error(err)
			return
		}
		cfg.Exec()
	}
}
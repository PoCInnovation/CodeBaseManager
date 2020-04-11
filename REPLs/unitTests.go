package REPLs

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/REPLs/utils"
	"os"
)

func FunctionalTestsShell() {
	for {
		fmt.Printf("%s > Unit Tests :> ", os.Getenv("PWD"))
		_ = utils.GetLine()
	}
}

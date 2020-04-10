package REPLs

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/REPLs/utils"
	"os"
)

func CodebaseShell() {
	for {
		fmt.Printf("%s > CodeBase :> ", os.Getenv("PWD"))
		_ = utils.GetLine()
	}
}

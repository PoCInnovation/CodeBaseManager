package REPLs

import (
	"fmt"
	"github.com/PoCFrance/CodeBaseManager/REPLs/utils"
	"os"
)

func DebugShell() {
	for {
		fmt.Printf("%s > Debug :> ", os.Getenv("PWD"))
		_ = utils.GetLine()
	}
}

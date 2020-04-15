package REPLs

import (
	"github.com/PoCFrance/CodeBaseManager/REPLs/common"
)

func CodebaseShell() {
	codebaseCmd := []string{"cat", "find"}
	p := common.NewPrompt("CodeBase")
	defer p.Close()

	for {
		p.Display()
		in := p.GetInput()
		parsed, todo := common.ParseInput(in, codebaseCmd)

		switch todo {
		case common.Builtin:
			common.HandleBuiltin(parsed)
		case common.ExternalBin:
			common.HandleExternal(parsed)
		case common.Exit:
			return
		}
	}
}

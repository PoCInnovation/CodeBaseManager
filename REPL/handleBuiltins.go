package REPL

import "github.com/PoCFrance/CodeBaseManager/REPL/builtins"

type builtin func([]string)
type Builtins map[string]builtin
//TODO: in cmd, lookup for cbm builtin via CBM codebase

var commonBuiltins = Builtins{
    "cd": builtins.CD,
}

func isBuiltin(bin string, searched Builtins) builtin {
    for k, fn := range searched {
        if k == bin {
            return fn
        }
    }
    return nil
}
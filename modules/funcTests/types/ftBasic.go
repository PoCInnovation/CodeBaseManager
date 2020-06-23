package ft_types

type ftBasic struct {
    Name string `toml:"name"`
    Desc string `toml:"desc"`
    Bin string `toml:"bin"`
    RefBin string `toml:"refBin"`
    Args []string `toml:"args"`
    RefArgs []string `toml:"refArgs"`
}

func (test *ftBasic) ApplyVars(vars map[string]string) {
    test.RefArgs = ApplyVarsTab(test.RefArgs, vars)
    test.Args = ApplyVarsTab(test.Args, vars)
    test.Bin = ApplyVarsString(test.Bin, vars)
    test.RefBin = ApplyVarsString(test.RefBin, vars)
}


func (test *ftBasic) ApplyDefault(reference ftBasic, vars map[string]string) {
    if len(test.RefArgs) == 0 {
        test.RefArgs = reference.RefArgs
    }
    if len(test.Args) == 0 {
        test.Args = reference.Args
    }
    if len(test.Bin) == 0 {
        test.Bin = reference.Bin
    }
    if len(test.RefBin) == 0 {
        test.RefBin = reference.RefBin
    }
    test.ApplyVars(vars)
}
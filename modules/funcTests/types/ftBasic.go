package ft_types

type ftBasic struct {
    Name string `toml:"name"`
    Desc string `toml:"desc"`
    Bin string `toml:"bin"`
    RefBin string `toml:"refBin"`
    Args []string `toml:"args"`
    RefArgs []string `toml:"refArgs"`
}

func (test *ftBasic) ApplyDefault(reference ftBasic) {
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
}
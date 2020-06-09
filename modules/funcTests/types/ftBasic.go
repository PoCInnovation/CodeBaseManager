package ft_types

type ftBasic struct {
    Name string `toml:"name"`
    Desc string `toml:"desc"`
    Bin string `toml:"bin"`
    RefBin string `toml:"refBin"`
    Args []string `toml:"args"`
    RefArgs []string `toml:"refArgs"`
}
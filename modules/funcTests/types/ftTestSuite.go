package ft_types

type ftTestSuite struct {
    Vars map[string]string `toml:"vars"`
    Default ftDescription `toml:"default"`
    Tests []ftDescription `toml:"Test"`
}
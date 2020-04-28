package funcTests

import (
    "errors"
    "fmt"
    "github.com/BurntSushi/toml"
)

type ConfigFT struct {
    Common struct {
        Bin    string    `toml:"bin"`
        RefBin string    `toml:"refBin"`
        Opt    ftOptions `toml:"options"`
    }
    Tests []FT `toml:"Test"`
}

func NewConfigFT(cfgPath string) (*ConfigFT, error) {
    cfg := &ConfigFT{}

    md, err := toml.DecodeFile(cfgPath, cfg)
    if err != nil {
        return nil, err
    }
    // Check if any key in the cfg were ignored
    ignored := md.Undecoded()
    if len(ignored) == 0 {
        fmt.Println("Successfully loaded config")
        return cfg, nil
    }
    // We're handling error in the TOML file.
    // TODO: Parse file to find where it fucked up :)
    errMsg := "Failed to load config, the following ignored were ignored:\n"
    for _, k := range ignored {
        errMsg = fmt.Sprintf("%s%v\n", errMsg, k)
    }
    return nil, errors.New(errMsg)
}


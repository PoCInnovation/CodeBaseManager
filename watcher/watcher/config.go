package watcher

import (
    "errors"
    "fmt"
    "os"
    "github.com/BurntSushi/toml"
    )

type Config struct {
    Repos []string `toml:"repositories"`
    Files []string `toml:"files"`
    Events []string `toml:"events"`
    Ignored []string `toml:"ignore"`
}

func NewConfig() (*Config, error) {
    cfg := &Config{}
    md, err := toml.DecodeFile(os.Getenv("HOME") + "/.cbm/watcher.toml", cfg)
    if err != nil {
        return nil, fmt.Errorf("failed to load Config: %v", err)
    }
    if len(md.Undecoded()) != 0 {
        return nil, errors.New("")
    }
    return cfg, nil
}

package watcher

import (
    "fmt"
)

func NewWatcher() error {
    cfg, err := NewConfig()
    if err != nil {
        return fmt.Errorf("failed to create watcher: %v", err)
    }
    _ = cfg
    return nil
}

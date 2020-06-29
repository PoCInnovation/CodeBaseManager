package repository

import (
    "os"
    "os/exec"
)

func DisplayTodos() {
    cmd := exec.Command("grep", "-rn", "TODO")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    _ = cmd.Run()
}

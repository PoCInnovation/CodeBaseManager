package repository

import (
    "os"
    "os/exec"
)

func DisplayTodos() {
    cmd := exec.Command("grep", "--color=auto", "--exclude-dir={.bzr,CVS,.git,.hg,.svn,.idea,.tox}", "-rn", "TODO")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    _ = cmd.Run()
}

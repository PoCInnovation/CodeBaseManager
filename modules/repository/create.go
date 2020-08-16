package repository

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateRepository(url string) {
	//cmd := fmt.Sprintf("git clone %s", url)
	fmt.Println("Creating repo based on:", url)
}

func execCmd(cmd string) {
	execCmd := exec.Command("bash", "-c", cmd)
	execCmd.Stdout = os.Stdout
	_ = execCmd.Run()
}

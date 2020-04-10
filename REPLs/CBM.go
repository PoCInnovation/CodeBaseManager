package REPLs

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func getLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil || input == "\n" {
		if err == io.EOF {
			fmt.Println("Exit.")
			os.Exit(0)
		}
		return ""
	}
	return input
}

func CBMShell() {
	for {
		fmt.Printf("%s > CBM :> ", os.Getenv("PWD"))
		_ = getLine()
	}
}

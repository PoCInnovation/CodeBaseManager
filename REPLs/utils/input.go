package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func GetLine() string {
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

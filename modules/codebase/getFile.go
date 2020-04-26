package codebase

import (
	"bufio"
	"log"
	"os"
)

func GetFile(fileName string) (*string, error) {
	fh, err := os.Open(fileName)

	defer func() {
		if err = fh.Close(); err != nil {
			log.Fatal("Error on closing file", err)
		}
	}()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	scanner := bufio.NewScanner(fh)
	var content string

	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &content, nil
}

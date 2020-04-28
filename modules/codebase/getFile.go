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
			log.Printf("Error on closing file %s, %v\n", fileName, err)
		}
	}()

	if err != nil {
		log.Printf("Error when opening file %s, %v\n", fileName, err)
		return nil, err
	}

	scanner := bufio.NewScanner(fh)
	var content string

	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error on closing file %s %v\n", fileName, err)
		return nil, err
	}

	return &content, nil
}

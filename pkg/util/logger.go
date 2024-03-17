package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WriteToFile(filePath string, text string) error {
	// Open or create the file for appending with user read/write permissions
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the text to the file
	_, err = file.WriteString(text + "\n")
	if err != nil {
		return err
	}

	return nil
}

func FindLatestValueForKey(filePath, key string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var value string
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+",") {
			split := strings.SplitN(line, ",", 2)
			if len(split) == 2 {
				value = split[1]
				found = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if !found {
		return "", fmt.Errorf("key not found")
	}

	return value, nil
}

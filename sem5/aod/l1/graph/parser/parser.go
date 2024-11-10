package parser

import (
	"bufio"
	"errors"
	"os"
)

func isFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return true, err
}

func Parse(path string) error {
	ok, err := isFileExist(path)
	if !ok {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Prinln("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	scanner.Scan()
}

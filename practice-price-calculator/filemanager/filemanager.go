package filemanager

import (
	"bufio"
	"errors"
	"os"
)

// ReadLines opens the file at the given path and returns its content as a slice
// of strings, one per line. It returns an error if the file cannot be opened or read.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("error reading file")
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("error scan from file")
	}
	file.Close()
	return lines, nil
}

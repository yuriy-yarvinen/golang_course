package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// ReadLines opens the file at the given FileManager InputFilePath and returns its content as a slice
// of strings, one per line. It returns an error if the file cannot be opened or read.
func (fileManager FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fileManager.InputFilePath)
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

func (fileManager FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fileManager.OutputFilePath)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return err
	}

	file.Close()
	return nil
}

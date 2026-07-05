package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"
)

// writeMutex сериализует запись в файл: пока одна горутина держит лок,
// остальные ждут Unlock и не могут писать одновременно.
var writeMutex sync.Mutex

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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return nil, errors.New("error scan from file")
	}
	return lines, nil
}

func (fileManager FileManager) WriteResult(data interface{}) error {
	time.Sleep(3 * time.Second) // Wait for 3 seconds before writing to the file
	writeMutex.Lock()
	defer writeMutex.Unlock()

	// Читаем текущее содержимое файла как объект (ключ = ставка).
	merged := map[string]json.RawMessage{}
	existing, err := os.ReadFile(fileManager.OutputFilePath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if len(existing) > 0 {
		if err := json.Unmarshal(existing, &merged); err != nil {
			return err
		}
	}

	// Результат горутины — это { "<ставка>": {цены...} }. Вливаем его ключи в merged.
	encoded, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var incoming map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &incoming); err != nil {
		return err
	}
	for k, v := range incoming {
		merged[k] = v
	}

	// Пишем весь объект обратно — файл всегда остаётся валидным JSON.
	out, err := json.MarshalIndent(merged, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileManager.OutputFilePath, out, 0644)
}

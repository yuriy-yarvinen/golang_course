package filefunc

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetNumberFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value value.")
	}

	return value, nil
}

func WriteNumberToFile(fileName string, number float64) {
	numberText := fmt.Sprint(number)
	os.WriteFile(fileName, []byte(numberText), 0644)
}

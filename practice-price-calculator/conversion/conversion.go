package conversion

import (
	"errors"
	"strconv"
)

// ConvertStringsToFloat64 converts a slice of strings to a slice of float64
// values. It returns an error if any of the strings cannot be parsed as a number.
func ConvertStringsToFloat64(stringsToConvert []string) ([]float64, error) {

	prices := make([]float64, len(stringsToConvert))
	for index, numberAsString := range stringsToConvert {
		price, err := strconv.ParseFloat(numberAsString, 64)
		if err != nil {
			return nil, errors.New("error strconv")
		}
		prices[index] = price
	}
	return prices, nil
}

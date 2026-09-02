package cmdmanager

import (
	"fmt"
)

type CMDmanager struct {
}

// ReadLines opens the file at the given FileManager InputFilePath and returns its content as a slice
// of strings, one per line. It returns an error if the file cannot be opened or read.
func (cmdManager CMDmanager) ReadLines() ([]string, error) {

	prices := []string{}

	fmt.Println("Write prices:")
	for {
		var price string
		fmt.Scan(&price)
		if price == "" || price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmdManager CMDmanager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

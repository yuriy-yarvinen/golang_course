package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64
	var err error

	revenue, err = getuserInput("Enter total revenue: ")
	if err != nil {
		println(err.Error())
		return
	}

	expenses, err = getuserInput("Enter total expenses: ")
	if err != nil {
		println(err.Error())
		return
	}

	taxRate, err = getuserInput("Enter tax rate (as a percentage, e.g., 20 for 20%): ")
	if err != nil {
		println(err.Error())
		return
	}
	netEarnings, taxAmount, earningsBeforeTax := calculateNetEarnings(revenue, expenses, taxRate)
	fmt.Printf("Earnings Before Tax: %.1f\n", earningsBeforeTax)
	fmt.Printf("Tax Amount: %.1f\n", taxAmount)
	fmt.Printf("Net Earnings: %.1f\n", netEarnings)

	writeToFile(fmt.Sprintf("Earnings Before Tax: %.1f\nTax Amount: %.1f\nNet Earnings: %.1f\n", earningsBeforeTax, taxAmount, netEarnings))

	text, err := readFromFile()
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(text)

}

func getuserInput(prompt string) (float64, error) {
	var input float64
	println(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %v", err)
	}
	return input, nil
}

func calculateNetEarnings(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	taxAmount := earningsBeforeTax * (taxRate / 100)
	netEarnings := earningsBeforeTax - taxAmount
	return netEarnings, taxAmount, earningsBeforeTax
}

func writeToFile(content string) error {
	text := fmt.Sprint(content)
	err := os.WriteFile("file.txt", []byte(text), 0644)
	return err
}

func readFromFile() (string, error) {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		return "", errors.New("error reading from file")
	}
	return string(data), nil
}

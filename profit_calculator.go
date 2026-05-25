package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	println("Enter total revenue: ")
	_, err := fmt.Scanln(&revenue)
	if err != nil {
		println("Invalid input for revenue. Please enter a valid number.")
		return
	}

	println("Enter total expenses: ")
	_, err = fmt.Scanln(&expenses)
	if err != nil {
		println("Invalid input for expenses. Please enter a valid number.")
		return
	}

	println("Enter tax rate (as a decimal, e.g., 0.2 for 20%): ")
	_, err = fmt.Scanln(&taxRate)
	if err != nil {
		println("Invalid input for tax rate. Please enter a valid number.")
		return
	}

	earningsBeforeTax := revenue - expenses
	taxAmount := earningsBeforeTax * taxRate
	netEarnings := earningsBeforeTax - taxAmount

	fmt.Printf("Net Earnings: %.2f\n", netEarnings)
}

package main

import (
	"fmt"
	"yarvinen-price-calculator/prices"
)

func main() {
	pricesSlice := []float64{10.0, 20.0, 30.0, 40.0, 50.0}
	taxRates := []float64{0, 0.07, 0.1, 0.15, 0.22}

	result := prices.CalculatePricesWithTax(pricesSlice, taxRates)

	fmt.Println(result)
}

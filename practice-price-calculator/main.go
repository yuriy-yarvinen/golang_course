package main

import (
	"yarvinen-price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15, 0.22}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}

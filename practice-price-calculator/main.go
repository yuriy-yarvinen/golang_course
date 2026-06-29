package main

import (
	"fmt"
	"yarvinen-price-calculator/filemanager"
	"yarvinen-price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15, 0.22}
	fileManager := filemanager.FileManager{
		OutputFilePath: "result.json",
		InputFilePath:  "prices.txt",
	}
	// cmdmanager := cmdmanager.CMDmanager{}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fileManager)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("error")
			fmt.Println(err)
		}
	}
}

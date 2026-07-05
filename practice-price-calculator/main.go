package main

import (
	"time"
	"yarvinen-price-calculator/filemanager"
	"yarvinen-price-calculator/prices"
)

func main() {
	start := time.Now()
	taxRates := []float64{0, 0.07, 0.1, 0.15, 0.22}
	fileManager := filemanager.FileManager{
		OutputFilePath: "result" + time.Now().Format("2006-01-02_15-04-05") + ".json",
		InputFilePath:  "prices.txt",
	}
	// cmdmanager := cmdmanager.CMDmanager{}
	priceJobChannels := make([]chan bool, len(taxRates))

	for i, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fileManager)
		priceJobChannels[i] = make(chan bool)

		go priceJob.Process(priceJobChannels[i])

		// if err != nil {
		// 	fmt.Println("error")
		// 	fmt.Println(err)
		// }
	}
	for _, priceJobChannel := range priceJobChannels {
		status := <-priceJobChannel
		if !status {
			println("Error occurred in one of the price jobs.")
		}
	}

	end := time.Now()
	elapsed := end.Sub(start)
	println("Elapsed time:", elapsed.String())
}

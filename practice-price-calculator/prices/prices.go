package prices

import (
	"fmt"
	"yarvinen-price-calculator/conversion"
	"yarvinen-price-calculator/filemanager"
)

// TaxIncludedPriceJob holds the data for a single tax-included price calculation:
// the tax rate to apply, the input prices loaded from a file, and the resulting
// tax-included prices keyed by their original value.
type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// Process loads the input prices, applies the tax rate to each of them and
// prints the resulting tax-included prices.
func (job TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}
	fmt.Println(result)

}

// LoadData reads the prices from the "prices.txt" file, converts them to
// float64 values and stores them in the job's InputPrices field.
func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")

	prices, err := conversion.ConvertStringsToFloat64(lines)
	if err != nil {
		fmt.Println("error scan from file")
		fmt.Println(err)
		return

	}
	job.InputPrices = prices

}

// NewTaxIncludedPriceJob creates a new TaxIncludedPriceJob initialized with the
// given tax rate.
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

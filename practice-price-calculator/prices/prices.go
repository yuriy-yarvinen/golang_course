package prices

import (
	"fmt"
	"yarvinen-price-calculator/conversion"
	"yarvinen-price-calculator/iomanager"
)

// TaxIncludedPriceJob holds the data for a single tax-included price calculation:
// the tax rate to apply, the input prices loaded from a file, and the resulting
// tax-included prices keyed by their original value.
type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
	IOManager         iomanager.IOmanager `json:"-"`
}

type PriceInfo struct {
	TaxRate       string `json:"tax_rate"`
	Price         string `json:"price"`
	PriceAfterTax string `json:"price_after_tax"`
}

// Process loads the input prices, applies the tax rate to each of them and
// prints the resulting tax-included prices.
func (job TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]PriceInfo)
	for _, price := range job.InputPrices {
		PriceInfo := PriceInfo{
			TaxRate:       fmt.Sprintf("%.2f", job.TaxRate*100),
			Price:         fmt.Sprintf("%.2f", price),
			PriceAfterTax: fmt.Sprintf("%.2f", price*(1+job.TaxRate)),
		}
		result[fmt.Sprintf("%.2f", price)] = PriceInfo
	}

	return job.IOManager.WriteResult(result)
}

// LoadData reads the prices from the "prices.txt" file, converts them to
// float64 values and stores them in the job's InputPrices field.
func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.ConvertStringsToFloat64(lines)
	if err != nil {
		return err

	}
	job.InputPrices = prices
	return nil
}

// NewTaxIncludedPriceJob creates a new TaxIncludedPriceJob initialized with the
// given tax rate.
func NewTaxIncludedPriceJob(taxRate float64, iomanager iomanager.IOmanager) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: iomanager,
	}
}

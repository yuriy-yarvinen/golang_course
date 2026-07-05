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
func (job TaxIncludedPriceJob) Process(channel chan bool, errorChannel chan error) {
	defer close(channel)
	defer close(errorChannel)

	err := job.LoadData()
	if err != nil {
		errorChannel <- err
		channel <- false
		return
	}
	taxKey := fmt.Sprintf("%.2f", job.TaxRate*100)
	result := make(map[string]map[string]PriceInfo)
	result[taxKey] = make(map[string]PriceInfo)
	for _, price := range job.InputPrices {
		PriceInfo := PriceInfo{
			TaxRate:       taxKey,
			Price:         fmt.Sprintf("%.2f", price),
			PriceAfterTax: fmt.Sprintf("%.2f", price*(1+job.TaxRate)),
		}

		result[taxKey][fmt.Sprintf("%.2f", price)] = PriceInfo
	}
	errWrite := job.IOManager.WriteResult(result)
	if errWrite != nil {
		errorChannel <- errWrite
		channel <- false
		return
	}
	channel <- true
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

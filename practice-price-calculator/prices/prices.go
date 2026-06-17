package prices

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64, inputPrices []float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:           taxRate,
		InputPrices:       inputPrices,
		TaxIncludedPrices: make(map[string]float64, len(inputPrices)),
	}
}

func CalculatePricesWithTax(prices []float64, taxRates []float64) map[float64][]float64 {
	result := make(map[float64][]float64, len(taxRates))
	for _, taxRate := range taxRates {
		for _, price := range prices {
			result[taxRate] = append(result[taxRate], price*(1+taxRate))
		}
	}

	return result
}

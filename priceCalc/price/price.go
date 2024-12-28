package price

import (
	"fmt"
	"project/price-calculator/conversion"
	"project/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager iomanager.IOManager `json:"-"`
	TaxRate     float64 `json:"tax_rate"`
	InputPrice  []float64 `json:"input_price"`
	OutputPrice map[string]string `json:"output_price"`
}

func New(iom iomanager.IOManager, tax float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:    tax,
		IOManager: iom,
	}
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.ReadPrices()
	if(err != nil){
		return err
	}
	result := make(map[string]string)
	for _, price := range job.InputPrice {
		result[fmt.Sprintf("%.2f" , price)] = fmt.Sprintf("%.4f" , price * (1 + job.TaxRate))
	}
	job.OutputPrice = result
	return job.IOManager.WriteJsonData(job)
}

func (job *TaxIncludedPriceJob) ReadPrices() error {
  lines, err := job.IOManager.ReadData()
	if(err != nil){
		return err
	}
	price, err := conversion.StringsToFloats(lines)
	if(err != nil){
		return err
	}
	job.InputPrice = price
	return nil
}
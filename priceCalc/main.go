package main

import (
	"fmt"
	"project/price-calculator/filemanager"
	"project/price-calculator/price"
)

func main() {
	// prices := []float64{10, 15, 25, 30}
	// rates := []float64{0, 0.07, 0.1, 0.15}
	// result := make(map[float64][]float64)
	// for _, rate := range rates {
	// 	taxPrice := make([]float64, len(prices))
	// 	for idx, price := range prices {
	// 		taxPrice[idx] = price * (1 + rate)
	// 	}
	// 	result[rate] = taxPrice
	// }
	// fmt.Println(result)
	rates := []float64{0, 0.07, 0.1, 0.15}
	for _, rate := range rates {
		fm := filemanager.New("price.txt", fmt.Sprintf("result_%.0f.json", rate*100))
		//cmd := cmdmanager.New()
		priceJob := price.New(fm, rate)
		err := priceJob.Process()
		if(err != nil){
			fmt.Println("Could not process the job...!")
			fmt.Println(err)
		}
	}
}
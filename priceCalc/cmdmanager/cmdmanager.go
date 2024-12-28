package cmdmanager

import (
	"fmt"
)

type CMDManager struct{}

func New() CMDManager {
	return CMDManager{}
}

func (fm CMDManager) ReadData() ([]string, error) {
	fmt.Println("Enter the price. Please press Enter to confirm..!")
	var prices []string

	for{
    var price string
		fmt.Scanln(&price); 
     
		if(price == "0"){
			break
		}
		prices = append(prices, price)
	}
   return prices, nil
}

func (fm CMDManager) WriteJsonData(data any) error {
	fmt.Println(data)
	return nil
}
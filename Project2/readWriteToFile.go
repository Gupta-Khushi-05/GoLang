package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenueAmount, err1 := OutputText("Enter the Revenue Amount: ")
	if(err1 != nil){
		panic(err1)
	}

	expensesAmount, err2 := OutputText("Enter the Expense Amount: ")
	if(err2 != nil){
		panic(err2)
	}

	taxRate, err3 := OutputText("Enter the Tax Rate: ")
	if(err3 != nil){
		panic(err3)
	}

	earningBeforeTax, earningAfterTax, ratio := CalculateValues(revenueAmount, expensesAmount, taxRate)
	fmt.Printf("Earning Before Tax [EBT]: %v\n", earningBeforeTax)
	fmt.Printf("Earning After Tax [Profit]: %v\n", earningAfterTax)
	fmt.Println("Ratio: ", ratio)

	// result := fmt.Sprintf("EBT: %.2f \n EFT: %.2f \n Ratio: %.4f \n", earningBeforeTax, earningAfterTax, ratio)
	// fmt.Println(result)	
	// ResultValues(earningBeforeTax, earningAfterTax, ratio)
}

func OutputText(text string) (float64, error){
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	if(userInput <= 0){
		return 0, errors.New("Please enter a correct value.")
	}
  return userInput, nil
}

func CalculateValues(revenue, expenses, tax float64) (ebt, eft, ratio float64) {
	ebt = revenue - expenses
	eft = ebt * (1 - tax/100)
	ratio = ebt / eft
  return
}

func ResultValues(ebt, eft, ratio float64){
	result := fmt.Sprintf("EBT: %.2f \n EFT: %.2f \n Ratio: %.4f \n", ebt, eft, ratio)
	fmt.Println(result)	
	os.WriteFile("balance.txt",  []byte(result), 0644)
}
package main

import "fmt"

func main() {
	var revenueAmount, expensesAmount, taxRate float64
	OutputText("Enter the Revenue Amount: ")
	fmt.Scan(&revenueAmount)
	OutputText("Enter the Expenses Amount: ")
	fmt.Scan(&expensesAmount)
	OutputText("Enter the Tax Rate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax, earningAfterTax, ratio := CalculateValues(revenueAmount, expensesAmount, taxRate)
	fmt.Printf("Earning Before Tax [EBT]: %v\n", earningBeforeTax)
	fmt.Printf("Earning After Tax [Profit]: %v\n", earningAfterTax)
	fmt.Println("Ratio: ", ratio)
}

func OutputText(text string){
	fmt.Print(text)
}

func CalculateValues(revenue, expenses, tax float64) (ebt, eft, ratio float64) {
	ebt = revenue - expenses
	eft = ebt * (1 - tax/100)
	ratio = ebt / eft
  return
}
package main

import "fmt"

func main() {
	var revenueAmount, expensesamount, taxRate float64
	fmt.Print("Enter the Revenue Amount: ")
	fmt.Scan(&revenueAmount)
	fmt.Print("Enter the Expenses Amount: ")
	fmt.Scan(&expensesamount)
	fmt.Print("Enter the Tax Rate: ")
	fmt.Scan(&taxRate)
	//Earning before tax [EBT]
	earningBeforeTax  := revenueAmount - expensesamount 
  //Earning after tax [profit]
	earningAfterTax := earningBeforeTax * (1 - taxRate/100) 
	//Ratio  
	ratio := earningBeforeTax / earningAfterTax
	fmt.Println("Earning Before Tax[EBT]: ", earningBeforeTax)
	fmt.Println("Earning After Tax[Profit]: ", earningAfterTax)
	fmt.Println("Ratio: ", ratio)
}


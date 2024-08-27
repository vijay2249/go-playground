package functions

import (
	"fmt"

	"wedonttrack.org/learn/go/basics/Basics/files"
)

// input - revenue, expenses, tax-rate
// calculate - earnings before & after tax
// calculate ratio - EBT/profit
// output - EBT, profit and the ratio

//Task 2
/*
1. Validations
	- Show error message and exit if invalid inut is provided
	- No Negative numbers as input
	- Not 0 as input
2. Store calculated results into file
*/

func CalculateRevenue(){
	var dataFile = "profit-calculations.txt"
	var revenue float64
	var expenses float64
	var taxRate float64

	getUserInput("Enter revenue: ", &revenue)
	// fmt.Print("Enter revenue: ")
	// fmt.Scan(&revenue)
	if (revenue <= 0) {
		panic("wrong input, value cannot be <= 0")
	}

	getUserInput("Enter expenses: ", &expenses)
	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)
	if expenses <= 0 {
		panic("wrong input, value cannot be <= 0")
	}

	getUserInput("Exter taxRate: ", &taxRate)
	// fmt.Print("Enter taxRate: ")
	// fmt.Scan(&taxRate)

	if (taxRate <= 0 ){
		panic("wrong input, value cannot be <= 0")
	}

	// ebt := revenue - expenses
	// profit := ebt *(1-taxRate/100)
	// ratio := ebt/profit
	ebt, profit, ratio := calculateEBT_Profit_Ratio(revenue, expenses, taxRate)

	// fmt.Println("EBT: ", ebt)
	// fmt.Println("Profit: ", profit)
	// fmt.Printf("Ratio: %.2f", ratio)

	outputTest := fmt.Sprintf("EBT: %v\nProfit: %v\nRatio: %.2f\n", ebt, profit, ratio)
	fmt.Println(outputTest)
	files.WriteDataToFile(dataFile, outputTest)
}

func calculateEBT_Profit_Ratio(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt *(1-taxRate/100)
	ratio = ebt/profit

	return ebt, profit, ratio
}

func getUserInput(outputText string, varName *float64) {
	fmt.Print(outputText)
	_, err := fmt.Scan(varName)
	if err != nil {
		panic("wrong input format - exiting")
	}
}

func GetUserInputV2(outputText string, varName float64) float64 {
	fmt.Print(outputText)
	_, err := fmt.Scan(&varName)
	if err != nil {
		panic("wrong input format - exiting")
	}
	return varName
}
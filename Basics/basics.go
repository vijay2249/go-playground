package basics

import (
	"fmt"
	"math"
)

const inflation = 4.5 //global scope

func calculateFutureValues(investAmount float64, expectedReturn float64, years int) (float64, float64) {
	futureValue := investAmount * math.Pow(1+expectedReturn/100, float64(years))
	futureRealValue := futureValue/(math.Pow(1+inflation/100, float64(years)))

	return futureValue, futureRealValue
}

func calculateFutureValuesNew(investAmount float64, expectedReturn float64, years int) (futureValue float64, futureRealValue float64) {
	futureValue = investAmount * math.Pow(1+expectedReturn/100, float64(years))
	futureRealValue = futureValue/(math.Pow(1+inflation/100, float64(years)))
	// return futureValue, futureRealValue
	return 
}

func Basics(){
	var investAmount float64= 100 //function scope = local variable
	var expectedReturn = 5.5
	var years = 10

	fmt.Print("Enter investmentAmount: ")
	fmt.Scan(&investAmount)

	fmt.Print("Enter expectedReturn: ")
	fmt.Scan(&expectedReturn)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	// futureValue := investAmount * math.Pow(1+expectedReturn/100, float64(years))
	// futureRealValue := futureValue/(math.Pow(1+inflation/100, float64(years)))

	futureValue, futureRealValue := calculateFutureValues(investAmount, expectedReturn, years)
	// futureValue, futureRealValue := calculateFutureValuesNew(investAmount, expectedReturn, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Println(formattedFV, formattedRFV)

	// fomatted string - multi-line string
	// fmt.Printf(`
	// 	Future value: %.2f
	// 	Future value (adjusted for inflation): %.2f
	// `, futureValue, futureRealValue)
}
package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Years of Investment: ")
	outputText("Years of Investment: ")
	fmt.Scan(&years)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)


	formattedFV, formattedRFV := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)
	

	// formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	// fmt.Printf(`Future Value: %.1f\n
	
	// Future Value (adjusted for inflation): %.1f`, futureValue, futureRealValue)

	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for inflation): %.1f", futureValue, futureRealValue)
	// println("Future Value: ",futureValue)
	// println("Future Value (adjusted for inflation): ",futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

// func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64){
// 	fv := investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
// 	rfv := fv / math.Pow(1 + inflationRate/100, years)

// 	return fv, rfv
// }

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64){
	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, years)
	rfv = fv / math.Pow(1 + inflationRate/100, years)

	// return fv, rfv
	return
}
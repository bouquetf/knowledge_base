package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)
	outputText("Years: ")
	fmt.Scan(&years)
	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, years, expectedReturnRate, inflationRate)

	formmattedFV := fmt.Sprintf("Future Value: %v\n", futureValue)
	formmattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %1f", futureRealValue)

	fmt.Print(formmattedFV, formmattedFRV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, years, expectedReturnRate, inflationRate float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}

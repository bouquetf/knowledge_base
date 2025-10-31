package main

import (
	"fmt"
)

func main() {
	revenue := getUserInput("Revenue:")
	expenses := getUserInput("Expenses:")
	taxRate := getUserInput("tax rate:")

	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println(earningsBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func getUserInput(text string) float64 {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func calculateFinancials(revenue, expences, taxrate float64) (earningsBeforeTax float64, profit float64, ratio float64) {
	earningsBeforeTax = revenue - expences
	profit = earningsBeforeTax * (1 - taxrate/100)
	ratio = profit / earningsBeforeTax

	return earningsBeforeTax, profit, ratio
}

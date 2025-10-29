package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue:")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	expenses, err := getUserInput("Expenses:")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	taxRate, err := getUserInput("tax rate:")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println(earningsBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)
	storeResults(earningsBeforeTax, profit, ratio)
}

func getUserInput(text string) (float64, error) {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return value, nil
}

func storeResults(ebt, profit, ration float64) {
	results := fmt.Sprintf("EBT: %f\nProfit: %f\nRatio: %f\n", ebt, profit, ration)
	os.WriteFile("results.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expences, taxrate float64) (earningsBeforeTax float64, profit float64, ratio float64) {
	earningsBeforeTax = revenue - expences
	profit = earningsBeforeTax * (1 - taxrate/100)
	ratio = profit / earningsBeforeTax

	return earningsBeforeTax, profit, ratio
}

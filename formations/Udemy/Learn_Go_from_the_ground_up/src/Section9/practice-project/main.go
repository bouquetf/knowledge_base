package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRate {
		iom := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate))
		//iom := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(iom, taxRate*100)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
			return
		}
	}
}

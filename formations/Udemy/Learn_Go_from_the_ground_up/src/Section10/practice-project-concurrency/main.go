package main

import (
	"fmt"

	"example.com/price-calculator-concurrency/filemanager"
	"example.com/price-calculator-concurrency/prices"
)

func main() {
	taxRate := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRate))
	errorChans := make([]chan error, len(taxRate))

	for index, taxRate := range taxRate {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		iom := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//iom := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(iom, taxRate*100)
		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRate {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done")
		}
	}
}

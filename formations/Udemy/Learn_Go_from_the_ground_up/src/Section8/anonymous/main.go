package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(numbers int) int {
	return numbers * 2
}

func triple(numbers int) int {
	return numbers * 3
}
